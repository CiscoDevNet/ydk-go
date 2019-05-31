// This module contains a collection of YANG definitions for
// monitoring the operation of ospf protocol in a Network Element.
// Copyright (c) 2016-2018 by Cisco Systems, Inc.
// All rights reserved.
package ospf_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ospf_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-ospf-oper ospf-oper-data}", reflect.TypeOf(OspfOperData{}))
    ydk.RegisterEntity("Cisco-IOS-XE-ospf-oper:ospf-oper-data", reflect.TypeOf(OspfOperData{}))
}

// AddressFamily represents Address family type
type AddressFamily string

const (
    AddressFamily_address_family_ipv4 AddressFamily = "address-family-ipv4"

    AddressFamily_address_family_ipv6 AddressFamily = "address-family-ipv6"
)

// OspfOperationMode represents OSPF operational mode
type OspfOperationMode string

const (
    // Ships-in-the-night operation mode in which each OSPF instance carries only one address family
    OspfOperationMode_ospf_ships_in_the_night OspfOperationMode = "ospf-ships-in-the-night"
)

// OspfNetworkType represents OSPF network type
type OspfNetworkType string

const (
    // OSPF broadcast multi-access network
    OspfNetworkType_ospf_broadcast OspfNetworkType = "ospf-broadcast"

    // OSPF Non-Broadcast Multi-Access (NBMA) network
    OspfNetworkType_ospf_non_broadcast OspfNetworkType = "ospf-non-broadcast"

    // OSPF point-to-multipoint network
    OspfNetworkType_ospf_point_to_multipoint OspfNetworkType = "ospf-point-to-multipoint"

    // OSPF point-to-point network
    OspfNetworkType_ospf_point_to_point OspfNetworkType = "ospf-point-to-point"
)

// OspfAuthType represents OSPF Authentication type
type OspfAuthType string

const (
    OspfAuthType_ospf_auth_ipsec OspfAuthType = "ospf-auth-ipsec"

    OspfAuthType_ospf_auth_trailer_keychain OspfAuthType = "ospf-auth-trailer-keychain"

    OspfAuthType_ospf_auth_trailer_key OspfAuthType = "ospf-auth-trailer-key"

    OspfAuthType_ospf_auth_type_none OspfAuthType = "ospf-auth-type-none"
)

// NbrStateType represents OSPF neighbor state type
type NbrStateType string

const (
    // Neighbor state down
    NbrStateType_ospf_nbr_down NbrStateType = "ospf-nbr-down"

    // Neighbor attempt state
    NbrStateType_ospf_nbr_attempt NbrStateType = "ospf-nbr-attempt"

    // Neighbor init state
    NbrStateType_ospf_nbr_init NbrStateType = "ospf-nbr-init"

    // Neighbor 2-way state
    NbrStateType_ospf_nbr_two_way NbrStateType = "ospf-nbr-two-way"

    // Neighbor exchange start state
    NbrStateType_ospf_nbr_exchange_start NbrStateType = "ospf-nbr-exchange-start"

    // Neighbor exchange state
    NbrStateType_ospf_nbr_exchange NbrStateType = "ospf-nbr-exchange"

    // Neighbor loading state
    NbrStateType_ospf_nbr_loading NbrStateType = "ospf-nbr-loading"

    // Neighbor full state
    NbrStateType_ospf_nbr_full NbrStateType = "ospf-nbr-full"
)

// Ospfv2LsaType represents Link State Advertisement type
type Ospfv2LsaType string

const (
    Ospfv2LsaType_ospfv2_lsa_type_unsupported_lsa_type Ospfv2LsaType = "ospfv2-lsa-type-unsupported-lsa-type"

    Ospfv2LsaType_ospfv2_lsa_type_router Ospfv2LsaType = "ospfv2-lsa-type-router"

    Ospfv2LsaType_ospfv2_lsa_type_network Ospfv2LsaType = "ospfv2-lsa-type-network"

    Ospfv2LsaType_ospfv2_lsa_type_summary_net Ospfv2LsaType = "ospfv2-lsa-type-summary-net"

    Ospfv2LsaType_ospfv2_lsa_type_summary_router Ospfv2LsaType = "ospfv2-lsa-type-summary-router"

    Ospfv2LsaType_ospfv2_lsa_type_as_external Ospfv2LsaType = "ospfv2-lsa-type-as-external"

    Ospfv2LsaType_ospfv2_lsa_type_nssa Ospfv2LsaType = "ospfv2-lsa-type-nssa"

    Ospfv2LsaType_ospfv2_lsa_type_link_scope_opaque Ospfv2LsaType = "ospfv2-lsa-type-link-scope-opaque"

    Ospfv2LsaType_ospfv2_lsa_type_area_scope_opaque Ospfv2LsaType = "ospfv2-lsa-type-area-scope-opaque"

    Ospfv2LsaType_ospfv2_lsa_type_as_scope_opaque Ospfv2LsaType = "ospfv2-lsa-type-as-scope-opaque"
)

// OspfExternalMetricType represents External metric type
type OspfExternalMetricType string

const (
    OspfExternalMetricType_ospf_ext_metric_type_1 OspfExternalMetricType = "ospf-ext-metric-type-1"

    OspfExternalMetricType_ospf_ext_metric_type_2 OspfExternalMetricType = "ospf-ext-metric-type-2"
)

// Ospfv2IntfState represents The possible states that an interface can be in
type Ospfv2IntfState string

const (
    // The interface is in the down state
    Ospfv2IntfState_ospfv2_interface_state_down Ospfv2IntfState = "ospfv2-interface-state-down"

    // The interface is in loopback state
    Ospfv2IntfState_ospfv2_interface_state_loopback Ospfv2IntfState = "ospfv2-interface-state-loopback"

    // The interface is in waiting state
    Ospfv2IntfState_ospfv2_interface_state_waiting Ospfv2IntfState = "ospfv2-interface-state-waiting"

    // The interface is in point-to-multipoint state
    Ospfv2IntfState_ospfv2_interface_state_point_to_mpoint Ospfv2IntfState = "ospfv2-interface-state-point-to-mpoint"

    // The interface is in point-to-point state
    Ospfv2IntfState_ospfv2_interface_state_point_to_point Ospfv2IntfState = "ospfv2-interface-state-point-to-point"

    // The interface is in the designated router state
    Ospfv2IntfState_ospfv2_interface_state_dr Ospfv2IntfState = "ospfv2-interface-state-dr"

    // The interface is providing backup for another 
    // interface
    Ospfv2IntfState_ospfv2_interface_state_backup Ospfv2IntfState = "ospfv2-interface-state-backup"

    // The interface is in a state other than the ones
    // nummerated in this list
    Ospfv2IntfState_ospfv2_interface_state_other Ospfv2IntfState = "ospfv2-interface-state-other"
)

// Ospfv2AuthTypeSelection represents The authentication type
type Ospfv2AuthTypeSelection string

const (
    // No authentication configured
    Ospfv2AuthTypeSelection_ospfv2_auth_none Ospfv2AuthTypeSelection = "ospfv2-auth-none"

    // Authentication uses the trailer key
    Ospfv2AuthTypeSelection_ospfv2_auth_trailer_key Ospfv2AuthTypeSelection = "ospfv2-auth-trailer-key"

    // Authentication uses a trailer key chain
    Ospfv2AuthTypeSelection_ospfv2_auth_trailer_key_chain Ospfv2AuthTypeSelection = "ospfv2-auth-trailer-key-chain"
)

// Ospfv2CryptoAlgorithm represents The algorithm in use
type Ospfv2CryptoAlgorithm string

const (
    // The OSPFv2 authentication is sent as cleartext
    Ospfv2CryptoAlgorithm_ospfv2_crypto_cleartest Ospfv2CryptoAlgorithm = "ospfv2-crypto-cleartest"

    // The OSPFv2 authentication is encrypted using 
    // Message Digest 5
    Ospfv2CryptoAlgorithm_ospfv2_crypto_md5 Ospfv2CryptoAlgorithm = "ospfv2-crypto-md5"
)

// OspfOperData
// Operational state of ospf
type OspfOperData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF operational state.
    OspfState OspfOperData_OspfState

    // The OSPF instance. The type is slice of OspfOperData_Ospfv2Instance.
    Ospfv2Instance []*OspfOperData_Ospfv2Instance
}

func (ospfOperData *OspfOperData) GetEntityData() *types.CommonEntityData {
    ospfOperData.EntityData.YFilter = ospfOperData.YFilter
    ospfOperData.EntityData.YangName = "ospf-oper-data"
    ospfOperData.EntityData.BundleName = "cisco_ios_xe"
    ospfOperData.EntityData.ParentYangName = "Cisco-IOS-XE-ospf-oper"
    ospfOperData.EntityData.SegmentPath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data"
    ospfOperData.EntityData.AbsolutePath = ospfOperData.EntityData.SegmentPath
    ospfOperData.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfOperData.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfOperData.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfOperData.EntityData.Children = types.NewOrderedMap()
    ospfOperData.EntityData.Children.Append("ospf-state", types.YChild{"OspfState", &ospfOperData.OspfState})
    ospfOperData.EntityData.Children.Append("ospfv2-instance", types.YChild{"Ospfv2Instance", nil})
    for i := range ospfOperData.Ospfv2Instance {
        ospfOperData.EntityData.Children.Append(types.GetSegmentPath(ospfOperData.Ospfv2Instance[i]), types.YChild{"Ospfv2Instance", ospfOperData.Ospfv2Instance[i]})
    }
    ospfOperData.EntityData.Leafs = types.NewOrderedMap()

    ospfOperData.EntityData.YListKeys = []string {}

    return &(ospfOperData.EntityData)
}

// OspfOperData_OspfState
// OSPF operational state
// This type is a presence type.
type OspfOperData_OspfState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // OSPF operation mode. The type is OspfOperationMode.
    OpMode interface{}

    // OSPF routing protocol instance. The type is slice of
    // OspfOperData_OspfState_OspfInstance.
    OspfInstance []*OspfOperData_OspfState_OspfInstance
}

func (ospfState *OspfOperData_OspfState) GetEntityData() *types.CommonEntityData {
    ospfState.EntityData.YFilter = ospfState.YFilter
    ospfState.EntityData.YangName = "ospf-state"
    ospfState.EntityData.BundleName = "cisco_ios_xe"
    ospfState.EntityData.ParentYangName = "ospf-oper-data"
    ospfState.EntityData.SegmentPath = "ospf-state"
    ospfState.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/" + ospfState.EntityData.SegmentPath
    ospfState.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfState.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfState.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfState.EntityData.Children = types.NewOrderedMap()
    ospfState.EntityData.Children.Append("ospf-instance", types.YChild{"OspfInstance", nil})
    for i := range ospfState.OspfInstance {
        ospfState.EntityData.Children.Append(types.GetSegmentPath(ospfState.OspfInstance[i]), types.YChild{"OspfInstance", ospfState.OspfInstance[i]})
    }
    ospfState.EntityData.Leafs = types.NewOrderedMap()
    ospfState.EntityData.Leafs.Append("op-mode", types.YLeaf{"OpMode", ospfState.OpMode})

    ospfState.EntityData.YListKeys = []string {}

    return &(ospfState.EntityData)
}

// OspfOperData_OspfState_OspfInstance
// OSPF routing protocol instance
type OspfOperData_OspfState_OspfInstance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Address-family of the instance. The type is
    // AddressFamily.
    Af interface{}

    // This attribute is a key. Defined in RFC 2328. A 32-bit number that uniquely
    // identifies the router. The type is interface{} with range: 0..4294967295.
    RouterId interface{}

    // The process identifier used to refer to this instance. The type is
    // interface{} with range: 0..65535.
    ProcessId interface{}

    // List of ospf areas. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea.
    OspfArea []*OspfOperData_OspfState_OspfInstance_OspfArea

    // List OSPF link scope LSA. The type is slice of
    // OspfOperData_OspfState_OspfInstance_LinkScopeLsas.
    LinkScopeLsas []*OspfOperData_OspfState_OspfInstance_LinkScopeLsas

    // OSPF multi-topology interface augmentation. The type is slice of
    // OspfOperData_OspfState_OspfInstance_MultiTopology.
    MultiTopology []*OspfOperData_OspfState_OspfInstance_MultiTopology
}

func (ospfInstance *OspfOperData_OspfState_OspfInstance) GetEntityData() *types.CommonEntityData {
    ospfInstance.EntityData.YFilter = ospfInstance.YFilter
    ospfInstance.EntityData.YangName = "ospf-instance"
    ospfInstance.EntityData.BundleName = "cisco_ios_xe"
    ospfInstance.EntityData.ParentYangName = "ospf-state"
    ospfInstance.EntityData.SegmentPath = "ospf-instance" + types.AddKeyToken(ospfInstance.Af, "af") + types.AddKeyToken(ospfInstance.RouterId, "router-id")
    ospfInstance.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/" + ospfInstance.EntityData.SegmentPath
    ospfInstance.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfInstance.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfInstance.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfInstance.EntityData.Children = types.NewOrderedMap()
    ospfInstance.EntityData.Children.Append("ospf-area", types.YChild{"OspfArea", nil})
    for i := range ospfInstance.OspfArea {
        ospfInstance.EntityData.Children.Append(types.GetSegmentPath(ospfInstance.OspfArea[i]), types.YChild{"OspfArea", ospfInstance.OspfArea[i]})
    }
    ospfInstance.EntityData.Children.Append("link-scope-lsas", types.YChild{"LinkScopeLsas", nil})
    for i := range ospfInstance.LinkScopeLsas {
        ospfInstance.EntityData.Children.Append(types.GetSegmentPath(ospfInstance.LinkScopeLsas[i]), types.YChild{"LinkScopeLsas", ospfInstance.LinkScopeLsas[i]})
    }
    ospfInstance.EntityData.Children.Append("multi-topology", types.YChild{"MultiTopology", nil})
    for i := range ospfInstance.MultiTopology {
        ospfInstance.EntityData.Children.Append(types.GetSegmentPath(ospfInstance.MultiTopology[i]), types.YChild{"MultiTopology", ospfInstance.MultiTopology[i]})
    }
    ospfInstance.EntityData.Leafs = types.NewOrderedMap()
    ospfInstance.EntityData.Leafs.Append("af", types.YLeaf{"Af", ospfInstance.Af})
    ospfInstance.EntityData.Leafs.Append("router-id", types.YLeaf{"RouterId", ospfInstance.RouterId})
    ospfInstance.EntityData.Leafs.Append("process-id", types.YLeaf{"ProcessId", ospfInstance.ProcessId})

    ospfInstance.EntityData.YListKeys = []string {"Af", "RouterId"}

    return &(ospfInstance.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea
// List of ospf areas
type OspfOperData_OspfState_OspfInstance_OspfArea struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. OSPF area ID. The type is interface{} with range:
    // 0..4294967295.
    AreaId interface{}

    // List of OSPF interfaces. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface.
    OspfInterface []*OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface

    // List of OSPF area scope LSA. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa.
    AreaScopeLsa []*OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa
}

func (ospfArea *OspfOperData_OspfState_OspfInstance_OspfArea) GetEntityData() *types.CommonEntityData {
    ospfArea.EntityData.YFilter = ospfArea.YFilter
    ospfArea.EntityData.YangName = "ospf-area"
    ospfArea.EntityData.BundleName = "cisco_ios_xe"
    ospfArea.EntityData.ParentYangName = "ospf-instance"
    ospfArea.EntityData.SegmentPath = "ospf-area" + types.AddKeyToken(ospfArea.AreaId, "area-id")
    ospfArea.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/" + ospfArea.EntityData.SegmentPath
    ospfArea.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfArea.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfArea.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfArea.EntityData.Children = types.NewOrderedMap()
    ospfArea.EntityData.Children.Append("ospf-interface", types.YChild{"OspfInterface", nil})
    for i := range ospfArea.OspfInterface {
        ospfArea.EntityData.Children.Append(types.GetSegmentPath(ospfArea.OspfInterface[i]), types.YChild{"OspfInterface", ospfArea.OspfInterface[i]})
    }
    ospfArea.EntityData.Children.Append("area-scope-lsa", types.YChild{"AreaScopeLsa", nil})
    for i := range ospfArea.AreaScopeLsa {
        ospfArea.EntityData.Children.Append(types.GetSegmentPath(ospfArea.AreaScopeLsa[i]), types.YChild{"AreaScopeLsa", ospfArea.AreaScopeLsa[i]})
    }
    ospfArea.EntityData.Leafs = types.NewOrderedMap()
    ospfArea.EntityData.Leafs.Append("area-id", types.YLeaf{"AreaId", ospfArea.AreaId})

    ospfArea.EntityData.YListKeys = []string {"AreaId"}

    return &(ospfArea.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface
// List of OSPF interfaces
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface name. The type is string.
    Name interface{}

    // Network type. The type is OspfNetworkType.
    NetworkType interface{}

    // Enable/Disable passive. The type is bool.
    Passive interface{}

    // Enable/Disable demand circuit. The type is bool.
    DemandCircuit interface{}

    // Set prefix as a node representative prefix. The type is bool.
    NodeFlag interface{}

    // Interface cost. The type is interface{} with range: 0..65535.
    Cost interface{}

    // Time between hello packets. The type is interface{} with range: 0..65535.
    HelloInterval interface{}

    // Interval after which a neighbor is declared dead. The type is interface{}
    // with range: 0..65535.
    DeadInterval interface{}

    // Time between retransmitting unacknowledged Link State Advertisements
    // (LSAs). The type is interface{} with range: 0..65535.
    RetransmitInterval interface{}

    // Estimated time needed to send link-state update. The type is interface{}
    // with range: 0..65535.
    TransmitDelay interface{}

    // Enable/Disable ignoring of MTU in DBD packets. The type is bool.
    MtuIgnore interface{}

    // Enable/Disable link-local signaling (LLS) support. The type is bool.
    Lls interface{}

    // Suppress advertisement of the prefixes. The type is bool.
    PrefixSuppression interface{}

    // Enable/disable bfd. The type is bool.
    Bfd interface{}

    // Enable/disable protocol on the interface. The type is bool.
    Enable interface{}

    // Interface state. The type is string.
    State interface{}

    // Hello timer. The type is interface{} with range: 0..4294967295.
    HelloTimer interface{}

    // Wait timer. The type is interface{} with range: 0..4294967295.
    WaitTimer interface{}

    // Designated Router. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Dr interface{}

    // Backup Designated Router. The type is one of the following types: string
    // with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Bdr interface{}

    // Configure OSPF router priority. The type is interface{} with range: 0..255.
    Priority interface{}

    // Multi Area.
    MultiArea OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_MultiArea

    // Staticly configured neighbors. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_StaticNeighbor.
    StaticNeighbor []*OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_StaticNeighbor

    // Fast reroute config.
    FastReroute OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_FastReroute

    // TTL security.
    TtlSecurity OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_TtlSecurity

    // Authentication configuration.
    Authentication OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_Authentication

    // List of OSPF neighbors. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_OspfNeighbor.
    OspfNeighbor []*OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_OspfNeighbor

    // List OSPF link scope LSAs. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas.
    IntfLinkScopeLsas []*OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas

    // OSPF interface topology. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfMultiTopology.
    IntfMultiTopology []*OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfMultiTopology
}

func (ospfInterface *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface) GetEntityData() *types.CommonEntityData {
    ospfInterface.EntityData.YFilter = ospfInterface.YFilter
    ospfInterface.EntityData.YangName = "ospf-interface"
    ospfInterface.EntityData.BundleName = "cisco_ios_xe"
    ospfInterface.EntityData.ParentYangName = "ospf-area"
    ospfInterface.EntityData.SegmentPath = "ospf-interface" + types.AddKeyToken(ospfInterface.Name, "name")
    ospfInterface.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/" + ospfInterface.EntityData.SegmentPath
    ospfInterface.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfInterface.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfInterface.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfInterface.EntityData.Children = types.NewOrderedMap()
    ospfInterface.EntityData.Children.Append("multi-area", types.YChild{"MultiArea", &ospfInterface.MultiArea})
    ospfInterface.EntityData.Children.Append("static-neighbor", types.YChild{"StaticNeighbor", nil})
    for i := range ospfInterface.StaticNeighbor {
        ospfInterface.EntityData.Children.Append(types.GetSegmentPath(ospfInterface.StaticNeighbor[i]), types.YChild{"StaticNeighbor", ospfInterface.StaticNeighbor[i]})
    }
    ospfInterface.EntityData.Children.Append("fast-reroute", types.YChild{"FastReroute", &ospfInterface.FastReroute})
    ospfInterface.EntityData.Children.Append("ttl-security", types.YChild{"TtlSecurity", &ospfInterface.TtlSecurity})
    ospfInterface.EntityData.Children.Append("authentication", types.YChild{"Authentication", &ospfInterface.Authentication})
    ospfInterface.EntityData.Children.Append("ospf-neighbor", types.YChild{"OspfNeighbor", nil})
    for i := range ospfInterface.OspfNeighbor {
        ospfInterface.EntityData.Children.Append(types.GetSegmentPath(ospfInterface.OspfNeighbor[i]), types.YChild{"OspfNeighbor", ospfInterface.OspfNeighbor[i]})
    }
    ospfInterface.EntityData.Children.Append("intf-link-scope-lsas", types.YChild{"IntfLinkScopeLsas", nil})
    for i := range ospfInterface.IntfLinkScopeLsas {
        ospfInterface.EntityData.Children.Append(types.GetSegmentPath(ospfInterface.IntfLinkScopeLsas[i]), types.YChild{"IntfLinkScopeLsas", ospfInterface.IntfLinkScopeLsas[i]})
    }
    ospfInterface.EntityData.Children.Append("intf-multi-topology", types.YChild{"IntfMultiTopology", nil})
    for i := range ospfInterface.IntfMultiTopology {
        ospfInterface.EntityData.Children.Append(types.GetSegmentPath(ospfInterface.IntfMultiTopology[i]), types.YChild{"IntfMultiTopology", ospfInterface.IntfMultiTopology[i]})
    }
    ospfInterface.EntityData.Leafs = types.NewOrderedMap()
    ospfInterface.EntityData.Leafs.Append("name", types.YLeaf{"Name", ospfInterface.Name})
    ospfInterface.EntityData.Leafs.Append("network-type", types.YLeaf{"NetworkType", ospfInterface.NetworkType})
    ospfInterface.EntityData.Leafs.Append("passive", types.YLeaf{"Passive", ospfInterface.Passive})
    ospfInterface.EntityData.Leafs.Append("demand-circuit", types.YLeaf{"DemandCircuit", ospfInterface.DemandCircuit})
    ospfInterface.EntityData.Leafs.Append("node-flag", types.YLeaf{"NodeFlag", ospfInterface.NodeFlag})
    ospfInterface.EntityData.Leafs.Append("cost", types.YLeaf{"Cost", ospfInterface.Cost})
    ospfInterface.EntityData.Leafs.Append("hello-interval", types.YLeaf{"HelloInterval", ospfInterface.HelloInterval})
    ospfInterface.EntityData.Leafs.Append("dead-interval", types.YLeaf{"DeadInterval", ospfInterface.DeadInterval})
    ospfInterface.EntityData.Leafs.Append("retransmit-interval", types.YLeaf{"RetransmitInterval", ospfInterface.RetransmitInterval})
    ospfInterface.EntityData.Leafs.Append("transmit-delay", types.YLeaf{"TransmitDelay", ospfInterface.TransmitDelay})
    ospfInterface.EntityData.Leafs.Append("mtu-ignore", types.YLeaf{"MtuIgnore", ospfInterface.MtuIgnore})
    ospfInterface.EntityData.Leafs.Append("lls", types.YLeaf{"Lls", ospfInterface.Lls})
    ospfInterface.EntityData.Leafs.Append("prefix-suppression", types.YLeaf{"PrefixSuppression", ospfInterface.PrefixSuppression})
    ospfInterface.EntityData.Leafs.Append("bfd", types.YLeaf{"Bfd", ospfInterface.Bfd})
    ospfInterface.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", ospfInterface.Enable})
    ospfInterface.EntityData.Leafs.Append("state", types.YLeaf{"State", ospfInterface.State})
    ospfInterface.EntityData.Leafs.Append("hello-timer", types.YLeaf{"HelloTimer", ospfInterface.HelloTimer})
    ospfInterface.EntityData.Leafs.Append("wait-timer", types.YLeaf{"WaitTimer", ospfInterface.WaitTimer})
    ospfInterface.EntityData.Leafs.Append("dr", types.YLeaf{"Dr", ospfInterface.Dr})
    ospfInterface.EntityData.Leafs.Append("bdr", types.YLeaf{"Bdr", ospfInterface.Bdr})
    ospfInterface.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", ospfInterface.Priority})

    ospfInterface.EntityData.YListKeys = []string {"Name"}

    return &(ospfInterface.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_MultiArea
// Multi Area
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_MultiArea struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Multi-area ID. The type is interface{} with range: 0..4294967295.
    MultiAreaId interface{}

    // Interface cost for multi-area. The type is interface{} with range:
    // 0..65535.
    Cost interface{}
}

func (multiArea *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_MultiArea) GetEntityData() *types.CommonEntityData {
    multiArea.EntityData.YFilter = multiArea.YFilter
    multiArea.EntityData.YangName = "multi-area"
    multiArea.EntityData.BundleName = "cisco_ios_xe"
    multiArea.EntityData.ParentYangName = "ospf-interface"
    multiArea.EntityData.SegmentPath = "multi-area"
    multiArea.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/" + multiArea.EntityData.SegmentPath
    multiArea.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    multiArea.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    multiArea.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    multiArea.EntityData.Children = types.NewOrderedMap()
    multiArea.EntityData.Leafs = types.NewOrderedMap()
    multiArea.EntityData.Leafs.Append("multi-area-id", types.YLeaf{"MultiAreaId", multiArea.MultiAreaId})
    multiArea.EntityData.Leafs.Append("cost", types.YLeaf{"Cost", multiArea.Cost})

    multiArea.EntityData.YListKeys = []string {}

    return &(multiArea.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_StaticNeighbor
// Staticly configured neighbors
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_StaticNeighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Neighbor IP address. The type is one of the
    // following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // Neighbor cost. The type is interface{} with range: 0..65535.
    Cost interface{}

    // Neighbor polling intervali in seconds. The type is interface{} with range:
    // 0..65535. Units are seconds.
    PollInterval interface{}
}

func (staticNeighbor *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_StaticNeighbor) GetEntityData() *types.CommonEntityData {
    staticNeighbor.EntityData.YFilter = staticNeighbor.YFilter
    staticNeighbor.EntityData.YangName = "static-neighbor"
    staticNeighbor.EntityData.BundleName = "cisco_ios_xe"
    staticNeighbor.EntityData.ParentYangName = "ospf-interface"
    staticNeighbor.EntityData.SegmentPath = "static-neighbor" + types.AddKeyToken(staticNeighbor.Address, "address")
    staticNeighbor.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/" + staticNeighbor.EntityData.SegmentPath
    staticNeighbor.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    staticNeighbor.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    staticNeighbor.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    staticNeighbor.EntityData.Children = types.NewOrderedMap()
    staticNeighbor.EntityData.Leafs = types.NewOrderedMap()
    staticNeighbor.EntityData.Leafs.Append("address", types.YLeaf{"Address", staticNeighbor.Address})
    staticNeighbor.EntityData.Leafs.Append("cost", types.YLeaf{"Cost", staticNeighbor.Cost})
    staticNeighbor.EntityData.Leafs.Append("poll-interval", types.YLeaf{"PollInterval", staticNeighbor.PollInterval})

    staticNeighbor.EntityData.YListKeys = []string {"Address"}

    return &(staticNeighbor.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_FastReroute
// Fast reroute config
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_FastReroute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Prevent the interface to be used as backup. The type is bool.
    CandidateDisabled interface{}

    // Activates LFA. This model assumes activation of per-prefix LFA. The type is
    // bool.
    Enabled interface{}

    // Activates remote LFA. The type is bool.
    RemoteLfaEnabled interface{}
}

func (fastReroute *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_FastReroute) GetEntityData() *types.CommonEntityData {
    fastReroute.EntityData.YFilter = fastReroute.YFilter
    fastReroute.EntityData.YangName = "fast-reroute"
    fastReroute.EntityData.BundleName = "cisco_ios_xe"
    fastReroute.EntityData.ParentYangName = "ospf-interface"
    fastReroute.EntityData.SegmentPath = "fast-reroute"
    fastReroute.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/" + fastReroute.EntityData.SegmentPath
    fastReroute.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    fastReroute.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    fastReroute.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    fastReroute.EntityData.Children = types.NewOrderedMap()
    fastReroute.EntityData.Leafs = types.NewOrderedMap()
    fastReroute.EntityData.Leafs.Append("candidate-disabled", types.YLeaf{"CandidateDisabled", fastReroute.CandidateDisabled})
    fastReroute.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", fastReroute.Enabled})
    fastReroute.EntityData.Leafs.Append("remote-lfa-enabled", types.YLeaf{"RemoteLfaEnabled", fastReroute.RemoteLfaEnabled})

    fastReroute.EntityData.YListKeys = []string {}

    return &(fastReroute.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_TtlSecurity
// TTL security
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_TtlSecurity struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable/Disable TTL security check. The type is bool.
    Enabled interface{}

    // Maximum number of hops that a OSPF packet may have traveled. The type is
    // interface{} with range: 0..255.
    Hops interface{}
}

func (ttlSecurity *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_TtlSecurity) GetEntityData() *types.CommonEntityData {
    ttlSecurity.EntityData.YFilter = ttlSecurity.YFilter
    ttlSecurity.EntityData.YangName = "ttl-security"
    ttlSecurity.EntityData.BundleName = "cisco_ios_xe"
    ttlSecurity.EntityData.ParentYangName = "ospf-interface"
    ttlSecurity.EntityData.SegmentPath = "ttl-security"
    ttlSecurity.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/" + ttlSecurity.EntityData.SegmentPath
    ttlSecurity.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ttlSecurity.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ttlSecurity.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ttlSecurity.EntityData.Children = types.NewOrderedMap()
    ttlSecurity.EntityData.Leafs = types.NewOrderedMap()
    ttlSecurity.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", ttlSecurity.Enabled})
    ttlSecurity.EntityData.Leafs.Append("hops", types.YLeaf{"Hops", ttlSecurity.Hops})

    ttlSecurity.EntityData.YListKeys = []string {}

    return &(ttlSecurity.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_Authentication
// Authentication configuration
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_Authentication struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SA name. The type is string.
    Sa interface{}

    // key-chain name. The type is string.
    KeyChain interface{}

    // Key string in ASCII format. The type is string.
    KeyString interface{}

    // No authentication enabled. The type is interface{} with range:
    // 0..4294967295.
    NoAuth interface{}

    // Crypto algorithm.
    CryptoAlgorithmVal OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_Authentication_CryptoAlgorithmVal
}

func (authentication *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_Authentication) GetEntityData() *types.CommonEntityData {
    authentication.EntityData.YFilter = authentication.YFilter
    authentication.EntityData.YangName = "authentication"
    authentication.EntityData.BundleName = "cisco_ios_xe"
    authentication.EntityData.ParentYangName = "ospf-interface"
    authentication.EntityData.SegmentPath = "authentication"
    authentication.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/" + authentication.EntityData.SegmentPath
    authentication.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    authentication.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    authentication.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    authentication.EntityData.Children = types.NewOrderedMap()
    authentication.EntityData.Children.Append("crypto-algorithm-val", types.YChild{"CryptoAlgorithmVal", &authentication.CryptoAlgorithmVal})
    authentication.EntityData.Leafs = types.NewOrderedMap()
    authentication.EntityData.Leafs.Append("sa", types.YLeaf{"Sa", authentication.Sa})
    authentication.EntityData.Leafs.Append("key-chain", types.YLeaf{"KeyChain", authentication.KeyChain})
    authentication.EntityData.Leafs.Append("key-string", types.YLeaf{"KeyString", authentication.KeyString})
    authentication.EntityData.Leafs.Append("no-auth", types.YLeaf{"NoAuth", authentication.NoAuth})

    authentication.EntityData.YListKeys = []string {}

    return &(authentication.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_Authentication_CryptoAlgorithmVal
// Crypto algorithm
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_Authentication_CryptoAlgorithmVal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // HMAC-SHA1-12 algorithm. The type is interface{}.
    HmacSha112 interface{}

    // HMAC-SHA1-20 algorithm. The type is interface{}.
    HmacSha120 interface{}

    // MD5 algorithm. The type is interface{}.
    Md5 interface{}

    // SHA-1 algorithm. The type is interface{}.
    Sha1 interface{}

    // HMAC-SHA-1 authentication algorithm. The type is interface{}.
    HmacSha1 interface{}

    // HMAC-SHA-256 authentication algorithm. The type is interface{}.
    HmacSha256 interface{}

    // HMAC-SHA-384 authentication algorithm. The type is interface{}.
    HmacSha384 interface{}

    // HMAC-SHA-512 authentication algorithm. The type is interface{}.
    HmacSha512 interface{}
}

func (cryptoAlgorithmVal *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_Authentication_CryptoAlgorithmVal) GetEntityData() *types.CommonEntityData {
    cryptoAlgorithmVal.EntityData.YFilter = cryptoAlgorithmVal.YFilter
    cryptoAlgorithmVal.EntityData.YangName = "crypto-algorithm-val"
    cryptoAlgorithmVal.EntityData.BundleName = "cisco_ios_xe"
    cryptoAlgorithmVal.EntityData.ParentYangName = "authentication"
    cryptoAlgorithmVal.EntityData.SegmentPath = "crypto-algorithm-val"
    cryptoAlgorithmVal.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/authentication/" + cryptoAlgorithmVal.EntityData.SegmentPath
    cryptoAlgorithmVal.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cryptoAlgorithmVal.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cryptoAlgorithmVal.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cryptoAlgorithmVal.EntityData.Children = types.NewOrderedMap()
    cryptoAlgorithmVal.EntityData.Leafs = types.NewOrderedMap()
    cryptoAlgorithmVal.EntityData.Leafs.Append("hmac-sha1-12", types.YLeaf{"HmacSha112", cryptoAlgorithmVal.HmacSha112})
    cryptoAlgorithmVal.EntityData.Leafs.Append("hmac-sha1-20", types.YLeaf{"HmacSha120", cryptoAlgorithmVal.HmacSha120})
    cryptoAlgorithmVal.EntityData.Leafs.Append("md5", types.YLeaf{"Md5", cryptoAlgorithmVal.Md5})
    cryptoAlgorithmVal.EntityData.Leafs.Append("sha-1", types.YLeaf{"Sha1", cryptoAlgorithmVal.Sha1})
    cryptoAlgorithmVal.EntityData.Leafs.Append("hmac-sha-1", types.YLeaf{"HmacSha1", cryptoAlgorithmVal.HmacSha1})
    cryptoAlgorithmVal.EntityData.Leafs.Append("hmac-sha-256", types.YLeaf{"HmacSha256", cryptoAlgorithmVal.HmacSha256})
    cryptoAlgorithmVal.EntityData.Leafs.Append("hmac-sha-384", types.YLeaf{"HmacSha384", cryptoAlgorithmVal.HmacSha384})
    cryptoAlgorithmVal.EntityData.Leafs.Append("hmac-sha-512", types.YLeaf{"HmacSha512", cryptoAlgorithmVal.HmacSha512})

    cryptoAlgorithmVal.EntityData.YListKeys = []string {}

    return &(cryptoAlgorithmVal.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_OspfNeighbor
// List of OSPF neighbors
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_OspfNeighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. OSPF neighbor ID. The type is one of the following
    // types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NeighborId interface{}

    // Neighbor address. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // Designated Router. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Dr interface{}

    // Backup Designated Router. The type is one of the following types: string
    // with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Bdr interface{}

    // OSPF neighbor state. The type is NbrStateType.
    State interface{}

    // Per-neighbor statistics.
    Stats OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_OspfNeighbor_Stats
}

func (ospfNeighbor *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_OspfNeighbor) GetEntityData() *types.CommonEntityData {
    ospfNeighbor.EntityData.YFilter = ospfNeighbor.YFilter
    ospfNeighbor.EntityData.YangName = "ospf-neighbor"
    ospfNeighbor.EntityData.BundleName = "cisco_ios_xe"
    ospfNeighbor.EntityData.ParentYangName = "ospf-interface"
    ospfNeighbor.EntityData.SegmentPath = "ospf-neighbor" + types.AddKeyToken(ospfNeighbor.NeighborId, "neighbor-id")
    ospfNeighbor.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/" + ospfNeighbor.EntityData.SegmentPath
    ospfNeighbor.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfNeighbor.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfNeighbor.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfNeighbor.EntityData.Children = types.NewOrderedMap()
    ospfNeighbor.EntityData.Children.Append("stats", types.YChild{"Stats", &ospfNeighbor.Stats})
    ospfNeighbor.EntityData.Leafs = types.NewOrderedMap()
    ospfNeighbor.EntityData.Leafs.Append("neighbor-id", types.YLeaf{"NeighborId", ospfNeighbor.NeighborId})
    ospfNeighbor.EntityData.Leafs.Append("address", types.YLeaf{"Address", ospfNeighbor.Address})
    ospfNeighbor.EntityData.Leafs.Append("dr", types.YLeaf{"Dr", ospfNeighbor.Dr})
    ospfNeighbor.EntityData.Leafs.Append("bdr", types.YLeaf{"Bdr", ospfNeighbor.Bdr})
    ospfNeighbor.EntityData.Leafs.Append("state", types.YLeaf{"State", ospfNeighbor.State})

    ospfNeighbor.EntityData.YListKeys = []string {"NeighborId"}

    return &(ospfNeighbor.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_OspfNeighbor_Stats
// Per-neighbor statistics
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_OspfNeighbor_Stats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The number of time this neighbor has changed state or an error has
    // occurred. The type is interface{} with range: 0..4294967295.
    NbrEventCount interface{}

    // The current length of the retransmission queue. The type is interface{}
    // with range: 0..4294967295.
    NbrRetransQlen interface{}
}

func (stats *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_OspfNeighbor_Stats) GetEntityData() *types.CommonEntityData {
    stats.EntityData.YFilter = stats.YFilter
    stats.EntityData.YangName = "stats"
    stats.EntityData.BundleName = "cisco_ios_xe"
    stats.EntityData.ParentYangName = "ospf-neighbor"
    stats.EntityData.SegmentPath = "stats"
    stats.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/ospf-neighbor/" + stats.EntityData.SegmentPath
    stats.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stats.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stats.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stats.EntityData.Children = types.NewOrderedMap()
    stats.EntityData.Leafs = types.NewOrderedMap()
    stats.EntityData.Leafs.Append("nbr-event-count", types.YLeaf{"NbrEventCount", stats.NbrEventCount})
    stats.EntityData.Leafs.Append("nbr-retrans-qlen", types.YLeaf{"NbrRetransQlen", stats.NbrRetransQlen})

    stats.EntityData.YListKeys = []string {}

    return &(stats.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas
// List OSPF link scope LSAs
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. OSPF link scope LSA type. The type is interface{}
    // with range: 0..4294967295.
    LsaType interface{}

    // List of OSPF link scope LSAs. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa.
    LinkScopeLsa []*OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa

    // List OSPF area scope LSA databases. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa.
    AreaScopeLsa []*OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa
}

func (intfLinkScopeLsas *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas) GetEntityData() *types.CommonEntityData {
    intfLinkScopeLsas.EntityData.YFilter = intfLinkScopeLsas.YFilter
    intfLinkScopeLsas.EntityData.YangName = "intf-link-scope-lsas"
    intfLinkScopeLsas.EntityData.BundleName = "cisco_ios_xe"
    intfLinkScopeLsas.EntityData.ParentYangName = "ospf-interface"
    intfLinkScopeLsas.EntityData.SegmentPath = "intf-link-scope-lsas" + types.AddKeyToken(intfLinkScopeLsas.LsaType, "lsa-type")
    intfLinkScopeLsas.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/" + intfLinkScopeLsas.EntityData.SegmentPath
    intfLinkScopeLsas.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    intfLinkScopeLsas.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    intfLinkScopeLsas.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    intfLinkScopeLsas.EntityData.Children = types.NewOrderedMap()
    intfLinkScopeLsas.EntityData.Children.Append("link-scope-lsa", types.YChild{"LinkScopeLsa", nil})
    for i := range intfLinkScopeLsas.LinkScopeLsa {
        intfLinkScopeLsas.EntityData.Children.Append(types.GetSegmentPath(intfLinkScopeLsas.LinkScopeLsa[i]), types.YChild{"LinkScopeLsa", intfLinkScopeLsas.LinkScopeLsa[i]})
    }
    intfLinkScopeLsas.EntityData.Children.Append("area-scope-lsa", types.YChild{"AreaScopeLsa", nil})
    for i := range intfLinkScopeLsas.AreaScopeLsa {
        intfLinkScopeLsas.EntityData.Children.Append(types.GetSegmentPath(intfLinkScopeLsas.AreaScopeLsa[i]), types.YChild{"AreaScopeLsa", intfLinkScopeLsas.AreaScopeLsa[i]})
    }
    intfLinkScopeLsas.EntityData.Leafs = types.NewOrderedMap()
    intfLinkScopeLsas.EntityData.Leafs.Append("lsa-type", types.YLeaf{"LsaType", intfLinkScopeLsas.LsaType})

    intfLinkScopeLsas.EntityData.YListKeys = []string {"LsaType"}

    return &(intfLinkScopeLsas.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa
// List of OSPF link scope LSAs
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. LSA ID. The type is interface{} with range:
    // 0..4294967295.
    LsaId interface{}

    // This attribute is a key. Advertising router. The type is one of the
    // following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    AdvRouter interface{}

    // The OSPF LSA body is fully decoded. The type is bool.
    DecodedCompleted interface{}

    // The complete LSA in network byte order as received/sent over the wire. The
    // type is slice of interface{} with range: 0..255.
    RawData []interface{}

    // Version. The type is interface{} with range: 0..4294967295.
    Version interface{}

    // Router address. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    RouterAddress interface{}

    // OSPFv2 LSA.
    Ospfv2Lsa OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2Lsa

    // OSPFv2 LSA link. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2Link.
    Ospfv2Link []*OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2Link

    // Summary LSA. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2Topology.
    Ospfv2Topology []*OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2Topology

    // External LSA. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2External.
    Ospfv2External []*OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2External

    // OSPFv2 Unknown TLV. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2UnknownTlv.
    Ospfv2UnknownTlv []*OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2UnknownTlv

    // OSPFv3 LSA.
    Ospfv3LsaVal OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal

    // OSPFv3 links. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3Link.
    Ospfv3Link []*OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3Link

    // OSPFv3 prefix-list. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3PrefixList.
    Ospfv3PrefixList []*OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3PrefixList

    // OSPFv3 intra-area prefix-list. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3IaPrefix.
    Ospfv3IaPrefix []*OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3IaPrefix

    // OSPF multi-topology interface augmentation. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_MultiTopology.
    MultiTopology []*OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_MultiTopology

    // Link TLV.
    Tlv OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Tlv

    // OSPFv2 Unknown sub TLV. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_UnknownSubTlv.
    UnknownSubTlv []*OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_UnknownSubTlv
}

func (linkScopeLsa *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa) GetEntityData() *types.CommonEntityData {
    linkScopeLsa.EntityData.YFilter = linkScopeLsa.YFilter
    linkScopeLsa.EntityData.YangName = "link-scope-lsa"
    linkScopeLsa.EntityData.BundleName = "cisco_ios_xe"
    linkScopeLsa.EntityData.ParentYangName = "intf-link-scope-lsas"
    linkScopeLsa.EntityData.SegmentPath = "link-scope-lsa" + types.AddKeyToken(linkScopeLsa.LsaId, "lsa-id") + types.AddKeyToken(linkScopeLsa.AdvRouter, "adv-router")
    linkScopeLsa.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/intf-link-scope-lsas/" + linkScopeLsa.EntityData.SegmentPath
    linkScopeLsa.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    linkScopeLsa.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    linkScopeLsa.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    linkScopeLsa.EntityData.Children = types.NewOrderedMap()
    linkScopeLsa.EntityData.Children.Append("ospfv2-lsa", types.YChild{"Ospfv2Lsa", &linkScopeLsa.Ospfv2Lsa})
    linkScopeLsa.EntityData.Children.Append("ospfv2-link", types.YChild{"Ospfv2Link", nil})
    for i := range linkScopeLsa.Ospfv2Link {
        linkScopeLsa.EntityData.Children.Append(types.GetSegmentPath(linkScopeLsa.Ospfv2Link[i]), types.YChild{"Ospfv2Link", linkScopeLsa.Ospfv2Link[i]})
    }
    linkScopeLsa.EntityData.Children.Append("ospfv2-topology", types.YChild{"Ospfv2Topology", nil})
    for i := range linkScopeLsa.Ospfv2Topology {
        linkScopeLsa.EntityData.Children.Append(types.GetSegmentPath(linkScopeLsa.Ospfv2Topology[i]), types.YChild{"Ospfv2Topology", linkScopeLsa.Ospfv2Topology[i]})
    }
    linkScopeLsa.EntityData.Children.Append("ospfv2-external", types.YChild{"Ospfv2External", nil})
    for i := range linkScopeLsa.Ospfv2External {
        linkScopeLsa.EntityData.Children.Append(types.GetSegmentPath(linkScopeLsa.Ospfv2External[i]), types.YChild{"Ospfv2External", linkScopeLsa.Ospfv2External[i]})
    }
    linkScopeLsa.EntityData.Children.Append("ospfv2-unknown-tlv", types.YChild{"Ospfv2UnknownTlv", nil})
    for i := range linkScopeLsa.Ospfv2UnknownTlv {
        linkScopeLsa.EntityData.Children.Append(types.GetSegmentPath(linkScopeLsa.Ospfv2UnknownTlv[i]), types.YChild{"Ospfv2UnknownTlv", linkScopeLsa.Ospfv2UnknownTlv[i]})
    }
    linkScopeLsa.EntityData.Children.Append("ospfv3-lsa-val", types.YChild{"Ospfv3LsaVal", &linkScopeLsa.Ospfv3LsaVal})
    linkScopeLsa.EntityData.Children.Append("ospfv3-link", types.YChild{"Ospfv3Link", nil})
    for i := range linkScopeLsa.Ospfv3Link {
        linkScopeLsa.EntityData.Children.Append(types.GetSegmentPath(linkScopeLsa.Ospfv3Link[i]), types.YChild{"Ospfv3Link", linkScopeLsa.Ospfv3Link[i]})
    }
    linkScopeLsa.EntityData.Children.Append("ospfv3-prefix-list", types.YChild{"Ospfv3PrefixList", nil})
    for i := range linkScopeLsa.Ospfv3PrefixList {
        linkScopeLsa.EntityData.Children.Append(types.GetSegmentPath(linkScopeLsa.Ospfv3PrefixList[i]), types.YChild{"Ospfv3PrefixList", linkScopeLsa.Ospfv3PrefixList[i]})
    }
    linkScopeLsa.EntityData.Children.Append("ospfv3-ia-prefix", types.YChild{"Ospfv3IaPrefix", nil})
    for i := range linkScopeLsa.Ospfv3IaPrefix {
        linkScopeLsa.EntityData.Children.Append(types.GetSegmentPath(linkScopeLsa.Ospfv3IaPrefix[i]), types.YChild{"Ospfv3IaPrefix", linkScopeLsa.Ospfv3IaPrefix[i]})
    }
    linkScopeLsa.EntityData.Children.Append("multi-topology", types.YChild{"MultiTopology", nil})
    for i := range linkScopeLsa.MultiTopology {
        linkScopeLsa.EntityData.Children.Append(types.GetSegmentPath(linkScopeLsa.MultiTopology[i]), types.YChild{"MultiTopology", linkScopeLsa.MultiTopology[i]})
    }
    linkScopeLsa.EntityData.Children.Append("tlv", types.YChild{"Tlv", &linkScopeLsa.Tlv})
    linkScopeLsa.EntityData.Children.Append("unknown-sub-tlv", types.YChild{"UnknownSubTlv", nil})
    for i := range linkScopeLsa.UnknownSubTlv {
        linkScopeLsa.EntityData.Children.Append(types.GetSegmentPath(linkScopeLsa.UnknownSubTlv[i]), types.YChild{"UnknownSubTlv", linkScopeLsa.UnknownSubTlv[i]})
    }
    linkScopeLsa.EntityData.Leafs = types.NewOrderedMap()
    linkScopeLsa.EntityData.Leafs.Append("lsa-id", types.YLeaf{"LsaId", linkScopeLsa.LsaId})
    linkScopeLsa.EntityData.Leafs.Append("adv-router", types.YLeaf{"AdvRouter", linkScopeLsa.AdvRouter})
    linkScopeLsa.EntityData.Leafs.Append("decoded-completed", types.YLeaf{"DecodedCompleted", linkScopeLsa.DecodedCompleted})
    linkScopeLsa.EntityData.Leafs.Append("raw-data", types.YLeaf{"RawData", linkScopeLsa.RawData})
    linkScopeLsa.EntityData.Leafs.Append("version", types.YLeaf{"Version", linkScopeLsa.Version})
    linkScopeLsa.EntityData.Leafs.Append("router-address", types.YLeaf{"RouterAddress", linkScopeLsa.RouterAddress})

    linkScopeLsa.EntityData.YListKeys = []string {"LsaId", "AdvRouter"}

    return &(linkScopeLsa.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2Lsa
// OSPFv2 LSA
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2Lsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Decoded OSPFv2 LSA header data.
    Header OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2Lsa_Header

    // Decoded OSPFv2 LSA body data.
    LsaBody OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2Lsa_LsaBody
}

func (ospfv2Lsa *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2Lsa) GetEntityData() *types.CommonEntityData {
    ospfv2Lsa.EntityData.YFilter = ospfv2Lsa.YFilter
    ospfv2Lsa.EntityData.YangName = "ospfv2-lsa"
    ospfv2Lsa.EntityData.BundleName = "cisco_ios_xe"
    ospfv2Lsa.EntityData.ParentYangName = "link-scope-lsa"
    ospfv2Lsa.EntityData.SegmentPath = "ospfv2-lsa"
    ospfv2Lsa.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/intf-link-scope-lsas/link-scope-lsa/" + ospfv2Lsa.EntityData.SegmentPath
    ospfv2Lsa.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv2Lsa.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv2Lsa.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv2Lsa.EntityData.Children = types.NewOrderedMap()
    ospfv2Lsa.EntityData.Children.Append("header", types.YChild{"Header", &ospfv2Lsa.Header})
    ospfv2Lsa.EntityData.Children.Append("lsa-body", types.YChild{"LsaBody", &ospfv2Lsa.LsaBody})
    ospfv2Lsa.EntityData.Leafs = types.NewOrderedMap()

    ospfv2Lsa.EntityData.YListKeys = []string {}

    return &(ospfv2Lsa.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2Lsa_Header
// Decoded OSPFv2 LSA header data
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2Lsa_Header struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA ID. The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    LsaId interface{}

    // Opaque type. The type is interface{} with range: 0..255.
    OpaqueType interface{}

    // Opaque ID. The type is interface{} with range: 0..4294967295.
    OpaqueId interface{}

    // LSA age. The type is interface{} with range: 0..65535.
    Age interface{}

    // LSA type. The type is interface{} with range: 0..65535.
    Type interface{}

    // LSA advertising router. The type is interface{} with range: 0..4294967295.
    AdvRouter interface{}

    // LSA sequence number. The type is string.
    SeqNum interface{}

    // LSA checksum. The type is string.
    Checksum interface{}

    // LSA length. The type is interface{} with range: 0..65535.
    Length interface{}

    // LSA options. The type is map[string]bool.
    FlagOptions interface{}
}

func (header *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2Lsa_Header) GetEntityData() *types.CommonEntityData {
    header.EntityData.YFilter = header.YFilter
    header.EntityData.YangName = "header"
    header.EntityData.BundleName = "cisco_ios_xe"
    header.EntityData.ParentYangName = "ospfv2-lsa"
    header.EntityData.SegmentPath = "header"
    header.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/intf-link-scope-lsas/link-scope-lsa/ospfv2-lsa/" + header.EntityData.SegmentPath
    header.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    header.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    header.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    header.EntityData.Children = types.NewOrderedMap()
    header.EntityData.Leafs = types.NewOrderedMap()
    header.EntityData.Leafs.Append("lsa-id", types.YLeaf{"LsaId", header.LsaId})
    header.EntityData.Leafs.Append("opaque-type", types.YLeaf{"OpaqueType", header.OpaqueType})
    header.EntityData.Leafs.Append("opaque-id", types.YLeaf{"OpaqueId", header.OpaqueId})
    header.EntityData.Leafs.Append("age", types.YLeaf{"Age", header.Age})
    header.EntityData.Leafs.Append("type", types.YLeaf{"Type", header.Type})
    header.EntityData.Leafs.Append("adv-router", types.YLeaf{"AdvRouter", header.AdvRouter})
    header.EntityData.Leafs.Append("seq-num", types.YLeaf{"SeqNum", header.SeqNum})
    header.EntityData.Leafs.Append("checksum", types.YLeaf{"Checksum", header.Checksum})
    header.EntityData.Leafs.Append("length", types.YLeaf{"Length", header.Length})
    header.EntityData.Leafs.Append("flag-options", types.YLeaf{"FlagOptions", header.FlagOptions})

    header.EntityData.YListKeys = []string {}

    return &(header.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2Lsa_LsaBody
// Decoded OSPFv2 LSA body data
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2Lsa_LsaBody struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of links. The type is interface{} with range: 0..65535.
    NumOfLinks interface{}

    // Summary mask. The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    SummaryMask interface{}

    // External mask. The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ExternalMask interface{}

    // LSA body flags. The type is map[string]bool.
    BodyFlagOptions interface{}

    // Network details.
    Network OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2Lsa_LsaBody_Network
}

func (lsaBody *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2Lsa_LsaBody) GetEntityData() *types.CommonEntityData {
    lsaBody.EntityData.YFilter = lsaBody.YFilter
    lsaBody.EntityData.YangName = "lsa-body"
    lsaBody.EntityData.BundleName = "cisco_ios_xe"
    lsaBody.EntityData.ParentYangName = "ospfv2-lsa"
    lsaBody.EntityData.SegmentPath = "lsa-body"
    lsaBody.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/intf-link-scope-lsas/link-scope-lsa/ospfv2-lsa/" + lsaBody.EntityData.SegmentPath
    lsaBody.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lsaBody.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lsaBody.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lsaBody.EntityData.Children = types.NewOrderedMap()
    lsaBody.EntityData.Children.Append("network", types.YChild{"Network", &lsaBody.Network})
    lsaBody.EntityData.Leafs = types.NewOrderedMap()
    lsaBody.EntityData.Leafs.Append("num-of-links", types.YLeaf{"NumOfLinks", lsaBody.NumOfLinks})
    lsaBody.EntityData.Leafs.Append("summary-mask", types.YLeaf{"SummaryMask", lsaBody.SummaryMask})
    lsaBody.EntityData.Leafs.Append("external-mask", types.YLeaf{"ExternalMask", lsaBody.ExternalMask})
    lsaBody.EntityData.Leafs.Append("body-flag-options", types.YLeaf{"BodyFlagOptions", lsaBody.BodyFlagOptions})

    lsaBody.EntityData.YListKeys = []string {}

    return &(lsaBody.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2Lsa_LsaBody_Network
// Network details
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2Lsa_LsaBody_Network struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP network mask. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NetworkMask interface{}

    // List of the routers attached to the network. The type is slice of
    // interface{} with range: 0..4294967295.
    AttachedRouter []interface{}
}

func (network *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2Lsa_LsaBody_Network) GetEntityData() *types.CommonEntityData {
    network.EntityData.YFilter = network.YFilter
    network.EntityData.YangName = "network"
    network.EntityData.BundleName = "cisco_ios_xe"
    network.EntityData.ParentYangName = "lsa-body"
    network.EntityData.SegmentPath = "network"
    network.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/intf-link-scope-lsas/link-scope-lsa/ospfv2-lsa/lsa-body/" + network.EntityData.SegmentPath
    network.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    network.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    network.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    network.EntityData.Children = types.NewOrderedMap()
    network.EntityData.Leafs = types.NewOrderedMap()
    network.EntityData.Leafs.Append("network-mask", types.YLeaf{"NetworkMask", network.NetworkMask})
    network.EntityData.Leafs.Append("attached-router", types.YLeaf{"AttachedRouter", network.AttachedRouter})

    network.EntityData.YListKeys = []string {}

    return &(network.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2Link
// OSPFv2 LSA link
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2Link struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Link ID. The type is interface{} with range:
    // 0..4294967295.
    LinkId interface{}

    // This attribute is a key. Link data. The type is interface{} with range:
    // 0..4294967295.
    LinkData interface{}

    // Link type. The type is interface{} with range: 0..255.
    Type interface{}

    // Topology specific information. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2Link_Ospfv2Topology.
    Ospfv2Topology []*OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2Link_Ospfv2Topology
}

func (ospfv2Link *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2Link) GetEntityData() *types.CommonEntityData {
    ospfv2Link.EntityData.YFilter = ospfv2Link.YFilter
    ospfv2Link.EntityData.YangName = "ospfv2-link"
    ospfv2Link.EntityData.BundleName = "cisco_ios_xe"
    ospfv2Link.EntityData.ParentYangName = "link-scope-lsa"
    ospfv2Link.EntityData.SegmentPath = "ospfv2-link" + types.AddKeyToken(ospfv2Link.LinkId, "link-id") + types.AddKeyToken(ospfv2Link.LinkData, "link-data")
    ospfv2Link.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/intf-link-scope-lsas/link-scope-lsa/" + ospfv2Link.EntityData.SegmentPath
    ospfv2Link.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv2Link.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv2Link.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv2Link.EntityData.Children = types.NewOrderedMap()
    ospfv2Link.EntityData.Children.Append("ospfv2-topology", types.YChild{"Ospfv2Topology", nil})
    for i := range ospfv2Link.Ospfv2Topology {
        ospfv2Link.EntityData.Children.Append(types.GetSegmentPath(ospfv2Link.Ospfv2Topology[i]), types.YChild{"Ospfv2Topology", ospfv2Link.Ospfv2Topology[i]})
    }
    ospfv2Link.EntityData.Leafs = types.NewOrderedMap()
    ospfv2Link.EntityData.Leafs.Append("link-id", types.YLeaf{"LinkId", ospfv2Link.LinkId})
    ospfv2Link.EntityData.Leafs.Append("link-data", types.YLeaf{"LinkData", ospfv2Link.LinkData})
    ospfv2Link.EntityData.Leafs.Append("type", types.YLeaf{"Type", ospfv2Link.Type})

    ospfv2Link.EntityData.YListKeys = []string {"LinkId", "LinkData"}

    return &(ospfv2Link.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2Link_Ospfv2Topology
// Topology specific information
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2Link_Ospfv2Topology struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. MT-ID for topology enabled link. The type is
    // interface{} with range: 0..4294967295.
    MtId interface{}

    // Metric for the topology. The type is interface{} with range: 0..65535.
    Metric interface{}
}

func (ospfv2Topology *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2Link_Ospfv2Topology) GetEntityData() *types.CommonEntityData {
    ospfv2Topology.EntityData.YFilter = ospfv2Topology.YFilter
    ospfv2Topology.EntityData.YangName = "ospfv2-topology"
    ospfv2Topology.EntityData.BundleName = "cisco_ios_xe"
    ospfv2Topology.EntityData.ParentYangName = "ospfv2-link"
    ospfv2Topology.EntityData.SegmentPath = "ospfv2-topology" + types.AddKeyToken(ospfv2Topology.MtId, "mt-id")
    ospfv2Topology.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/intf-link-scope-lsas/link-scope-lsa/ospfv2-link/" + ospfv2Topology.EntityData.SegmentPath
    ospfv2Topology.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv2Topology.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv2Topology.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv2Topology.EntityData.Children = types.NewOrderedMap()
    ospfv2Topology.EntityData.Leafs = types.NewOrderedMap()
    ospfv2Topology.EntityData.Leafs.Append("mt-id", types.YLeaf{"MtId", ospfv2Topology.MtId})
    ospfv2Topology.EntityData.Leafs.Append("metric", types.YLeaf{"Metric", ospfv2Topology.Metric})

    ospfv2Topology.EntityData.YListKeys = []string {"MtId"}

    return &(ospfv2Topology.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2Topology
// Summary LSA
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2Topology struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. MT-ID for topology enabled link. The type is
    // interface{} with range: 0..4294967295.
    MtId interface{}

    // Metric for the topology. The type is interface{} with range: 0..65535.
    Metric interface{}
}

func (ospfv2Topology *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2Topology) GetEntityData() *types.CommonEntityData {
    ospfv2Topology.EntityData.YFilter = ospfv2Topology.YFilter
    ospfv2Topology.EntityData.YangName = "ospfv2-topology"
    ospfv2Topology.EntityData.BundleName = "cisco_ios_xe"
    ospfv2Topology.EntityData.ParentYangName = "link-scope-lsa"
    ospfv2Topology.EntityData.SegmentPath = "ospfv2-topology" + types.AddKeyToken(ospfv2Topology.MtId, "mt-id")
    ospfv2Topology.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/intf-link-scope-lsas/link-scope-lsa/" + ospfv2Topology.EntityData.SegmentPath
    ospfv2Topology.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv2Topology.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv2Topology.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv2Topology.EntityData.Children = types.NewOrderedMap()
    ospfv2Topology.EntityData.Leafs = types.NewOrderedMap()
    ospfv2Topology.EntityData.Leafs.Append("mt-id", types.YLeaf{"MtId", ospfv2Topology.MtId})
    ospfv2Topology.EntityData.Leafs.Append("metric", types.YLeaf{"Metric", ospfv2Topology.Metric})

    ospfv2Topology.EntityData.YListKeys = []string {"MtId"}

    return &(ospfv2Topology.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2External
// External LSA
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2External struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. MT-ID for topology enabled on the link. The type
    // is interface{} with range: 0..4294967295.
    MtId interface{}

    // Metric for the topology. The type is interface{} with range: 0..4294967295.
    Metric interface{}

    // Forwarding address. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ForwardingAddress interface{}

    // Route tag. The type is interface{} with range: 0..4294967295.
    ExternalRouteTag interface{}
}

func (ospfv2External *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2External) GetEntityData() *types.CommonEntityData {
    ospfv2External.EntityData.YFilter = ospfv2External.YFilter
    ospfv2External.EntityData.YangName = "ospfv2-external"
    ospfv2External.EntityData.BundleName = "cisco_ios_xe"
    ospfv2External.EntityData.ParentYangName = "link-scope-lsa"
    ospfv2External.EntityData.SegmentPath = "ospfv2-external" + types.AddKeyToken(ospfv2External.MtId, "mt-id")
    ospfv2External.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/intf-link-scope-lsas/link-scope-lsa/" + ospfv2External.EntityData.SegmentPath
    ospfv2External.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv2External.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv2External.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv2External.EntityData.Children = types.NewOrderedMap()
    ospfv2External.EntityData.Leafs = types.NewOrderedMap()
    ospfv2External.EntityData.Leafs.Append("mt-id", types.YLeaf{"MtId", ospfv2External.MtId})
    ospfv2External.EntityData.Leafs.Append("metric", types.YLeaf{"Metric", ospfv2External.Metric})
    ospfv2External.EntityData.Leafs.Append("forwarding-address", types.YLeaf{"ForwardingAddress", ospfv2External.ForwardingAddress})
    ospfv2External.EntityData.Leafs.Append("external-route-tag", types.YLeaf{"ExternalRouteTag", ospfv2External.ExternalRouteTag})

    ospfv2External.EntityData.YListKeys = []string {"MtId"}

    return &(ospfv2External.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2UnknownTlv
// OSPFv2 Unknown TLV
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2UnknownTlv struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. TLV type. The type is interface{} with range:
    // 0..65535.
    Type interface{}

    // TLV length. The type is interface{} with range: 0..65535.
    Length interface{}

    // TLV value. The type is slice of interface{} with range: 0..255.
    Value []interface{}
}

func (ospfv2UnknownTlv *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2UnknownTlv) GetEntityData() *types.CommonEntityData {
    ospfv2UnknownTlv.EntityData.YFilter = ospfv2UnknownTlv.YFilter
    ospfv2UnknownTlv.EntityData.YangName = "ospfv2-unknown-tlv"
    ospfv2UnknownTlv.EntityData.BundleName = "cisco_ios_xe"
    ospfv2UnknownTlv.EntityData.ParentYangName = "link-scope-lsa"
    ospfv2UnknownTlv.EntityData.SegmentPath = "ospfv2-unknown-tlv" + types.AddKeyToken(ospfv2UnknownTlv.Type, "type")
    ospfv2UnknownTlv.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/intf-link-scope-lsas/link-scope-lsa/" + ospfv2UnknownTlv.EntityData.SegmentPath
    ospfv2UnknownTlv.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv2UnknownTlv.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv2UnknownTlv.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv2UnknownTlv.EntityData.Children = types.NewOrderedMap()
    ospfv2UnknownTlv.EntityData.Leafs = types.NewOrderedMap()
    ospfv2UnknownTlv.EntityData.Leafs.Append("type", types.YLeaf{"Type", ospfv2UnknownTlv.Type})
    ospfv2UnknownTlv.EntityData.Leafs.Append("length", types.YLeaf{"Length", ospfv2UnknownTlv.Length})
    ospfv2UnknownTlv.EntityData.Leafs.Append("value", types.YLeaf{"Value", ospfv2UnknownTlv.Value})

    ospfv2UnknownTlv.EntityData.YListKeys = []string {"Type"}

    return &(ospfv2UnknownTlv.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal
// OSPFv3 LSA
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Decoded OSPFv3 LSA header.
    Header OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_Header

    // Decoded OSPFv3 LSA body.
    LsaBody OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody
}

func (ospfv3LsaVal *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal) GetEntityData() *types.CommonEntityData {
    ospfv3LsaVal.EntityData.YFilter = ospfv3LsaVal.YFilter
    ospfv3LsaVal.EntityData.YangName = "ospfv3-lsa-val"
    ospfv3LsaVal.EntityData.BundleName = "cisco_ios_xe"
    ospfv3LsaVal.EntityData.ParentYangName = "link-scope-lsa"
    ospfv3LsaVal.EntityData.SegmentPath = "ospfv3-lsa-val"
    ospfv3LsaVal.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/intf-link-scope-lsas/link-scope-lsa/" + ospfv3LsaVal.EntityData.SegmentPath
    ospfv3LsaVal.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv3LsaVal.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv3LsaVal.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv3LsaVal.EntityData.Children = types.NewOrderedMap()
    ospfv3LsaVal.EntityData.Children.Append("header", types.YChild{"Header", &ospfv3LsaVal.Header})
    ospfv3LsaVal.EntityData.Children.Append("lsa-body", types.YChild{"LsaBody", &ospfv3LsaVal.LsaBody})
    ospfv3LsaVal.EntityData.Leafs = types.NewOrderedMap()

    ospfv3LsaVal.EntityData.YListKeys = []string {}

    return &(ospfv3LsaVal.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_Header
// Decoded OSPFv3 LSA header
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_Header struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA ID. The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    LsaId interface{}

    // OSPFv3 LSA options. The type is map[string]bool.
    LsaHdrOptions interface{}

    // LSA header.
    LsaHeader OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_Header_LsaHeader
}

func (header *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_Header) GetEntityData() *types.CommonEntityData {
    header.EntityData.YFilter = header.YFilter
    header.EntityData.YangName = "header"
    header.EntityData.BundleName = "cisco_ios_xe"
    header.EntityData.ParentYangName = "ospfv3-lsa-val"
    header.EntityData.SegmentPath = "header"
    header.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/intf-link-scope-lsas/link-scope-lsa/ospfv3-lsa-val/" + header.EntityData.SegmentPath
    header.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    header.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    header.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    header.EntityData.Children = types.NewOrderedMap()
    header.EntityData.Children.Append("lsa-header", types.YChild{"LsaHeader", &header.LsaHeader})
    header.EntityData.Leafs = types.NewOrderedMap()
    header.EntityData.Leafs.Append("lsa-id", types.YLeaf{"LsaId", header.LsaId})
    header.EntityData.Leafs.Append("lsa-hdr-options", types.YLeaf{"LsaHdrOptions", header.LsaHdrOptions})

    header.EntityData.YListKeys = []string {}

    return &(header.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_Header_LsaHeader
// LSA header
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_Header_LsaHeader struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA age. The type is interface{} with range: 0..65535.
    Age interface{}

    // LSA type. The type is interface{} with range: 0..65535.
    Type interface{}

    // LSA advertising router. The type is interface{} with range: 0..4294967295.
    AdvRouter interface{}

    // LSA sequence number. The type is string.
    SeqNum interface{}

    // LSA checksum. The type is string.
    Checksum interface{}

    // LSA length. The type is interface{} with range: 0..65535.
    Length interface{}
}

func (lsaHeader *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_Header_LsaHeader) GetEntityData() *types.CommonEntityData {
    lsaHeader.EntityData.YFilter = lsaHeader.YFilter
    lsaHeader.EntityData.YangName = "lsa-header"
    lsaHeader.EntityData.BundleName = "cisco_ios_xe"
    lsaHeader.EntityData.ParentYangName = "header"
    lsaHeader.EntityData.SegmentPath = "lsa-header"
    lsaHeader.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/intf-link-scope-lsas/link-scope-lsa/ospfv3-lsa-val/header/" + lsaHeader.EntityData.SegmentPath
    lsaHeader.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lsaHeader.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lsaHeader.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lsaHeader.EntityData.Children = types.NewOrderedMap()
    lsaHeader.EntityData.Leafs = types.NewOrderedMap()
    lsaHeader.EntityData.Leafs.Append("age", types.YLeaf{"Age", lsaHeader.Age})
    lsaHeader.EntityData.Leafs.Append("type", types.YLeaf{"Type", lsaHeader.Type})
    lsaHeader.EntityData.Leafs.Append("adv-router", types.YLeaf{"AdvRouter", lsaHeader.AdvRouter})
    lsaHeader.EntityData.Leafs.Append("seq-num", types.YLeaf{"SeqNum", lsaHeader.SeqNum})
    lsaHeader.EntityData.Leafs.Append("checksum", types.YLeaf{"Checksum", lsaHeader.Checksum})
    lsaHeader.EntityData.Leafs.Append("length", types.YLeaf{"Length", lsaHeader.Length})

    lsaHeader.EntityData.YListKeys = []string {}

    return &(lsaHeader.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody
// Decoded OSPFv3 LSA body
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA options. The type is map[string]bool.
    LsaFlagOptions interface{}

    // LSA Body Flags. The type is map[string]bool.
    LsaBodyFlags interface{}

    // OSPFv3 network.
    Network OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Network

    // OSPFv3 inter area prefix.
    Prefix OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Prefix

    // OSPFv3 inter area router.
    IaRouter OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_IaRouter

    // OSPFv3 LSA external.
    LsaExternal OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_LsaExternal

    // OSPFv3 NSSA.
    Nssa OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Nssa

    // OSPFv3 Link data.
    LinkData OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_LinkData

    // OSPFv3 Intra area prefixes.
    IaPrefix OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_IaPrefix
}

func (lsaBody *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody) GetEntityData() *types.CommonEntityData {
    lsaBody.EntityData.YFilter = lsaBody.YFilter
    lsaBody.EntityData.YangName = "lsa-body"
    lsaBody.EntityData.BundleName = "cisco_ios_xe"
    lsaBody.EntityData.ParentYangName = "ospfv3-lsa-val"
    lsaBody.EntityData.SegmentPath = "lsa-body"
    lsaBody.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/intf-link-scope-lsas/link-scope-lsa/ospfv3-lsa-val/" + lsaBody.EntityData.SegmentPath
    lsaBody.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lsaBody.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lsaBody.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lsaBody.EntityData.Children = types.NewOrderedMap()
    lsaBody.EntityData.Children.Append("network", types.YChild{"Network", &lsaBody.Network})
    lsaBody.EntityData.Children.Append("prefix", types.YChild{"Prefix", &lsaBody.Prefix})
    lsaBody.EntityData.Children.Append("ia-router", types.YChild{"IaRouter", &lsaBody.IaRouter})
    lsaBody.EntityData.Children.Append("lsa-external", types.YChild{"LsaExternal", &lsaBody.LsaExternal})
    lsaBody.EntityData.Children.Append("nssa", types.YChild{"Nssa", &lsaBody.Nssa})
    lsaBody.EntityData.Children.Append("link-data", types.YChild{"LinkData", &lsaBody.LinkData})
    lsaBody.EntityData.Children.Append("ia-prefix", types.YChild{"IaPrefix", &lsaBody.IaPrefix})
    lsaBody.EntityData.Leafs = types.NewOrderedMap()
    lsaBody.EntityData.Leafs.Append("lsa-flag-options", types.YLeaf{"LsaFlagOptions", lsaBody.LsaFlagOptions})
    lsaBody.EntityData.Leafs.Append("lsa-body-flags", types.YLeaf{"LsaBodyFlags", lsaBody.LsaBodyFlags})

    lsaBody.EntityData.YListKeys = []string {}

    return &(lsaBody.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Network
// OSPFv3 network
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Network struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of the routers attached to the network. The type is slice of
    // interface{} with range: 0..4294967295.
    AttachedRouter []interface{}

    // Network LSA options. The type is map[string]bool.
    LsaNetOptions interface{}
}

func (network *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Network) GetEntityData() *types.CommonEntityData {
    network.EntityData.YFilter = network.YFilter
    network.EntityData.YangName = "network"
    network.EntityData.BundleName = "cisco_ios_xe"
    network.EntityData.ParentYangName = "lsa-body"
    network.EntityData.SegmentPath = "network"
    network.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/intf-link-scope-lsas/link-scope-lsa/ospfv3-lsa-val/lsa-body/" + network.EntityData.SegmentPath
    network.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    network.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    network.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    network.EntityData.Children = types.NewOrderedMap()
    network.EntityData.Leafs = types.NewOrderedMap()
    network.EntityData.Leafs.Append("attached-router", types.YLeaf{"AttachedRouter", network.AttachedRouter})
    network.EntityData.Leafs.Append("lsa-net-options", types.YLeaf{"LsaNetOptions", network.LsaNetOptions})

    network.EntityData.YListKeys = []string {}

    return &(network.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Prefix
// OSPFv3 inter area prefix
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Prefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Metric. The type is interface{} with range: 0..4294967295.
    Metric interface{}

    // Inter area Prefix. The type is string.
    IaPrefix interface{}

    // Inter area prefix options. The type is string.
    IaPrefixOptions interface{}
}

func (prefix *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Prefix) GetEntityData() *types.CommonEntityData {
    prefix.EntityData.YFilter = prefix.YFilter
    prefix.EntityData.YangName = "prefix"
    prefix.EntityData.BundleName = "cisco_ios_xe"
    prefix.EntityData.ParentYangName = "lsa-body"
    prefix.EntityData.SegmentPath = "prefix"
    prefix.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/intf-link-scope-lsas/link-scope-lsa/ospfv3-lsa-val/lsa-body/" + prefix.EntityData.SegmentPath
    prefix.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    prefix.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    prefix.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    prefix.EntityData.Children = types.NewOrderedMap()
    prefix.EntityData.Leafs = types.NewOrderedMap()
    prefix.EntityData.Leafs.Append("metric", types.YLeaf{"Metric", prefix.Metric})
    prefix.EntityData.Leafs.Append("ia-prefix", types.YLeaf{"IaPrefix", prefix.IaPrefix})
    prefix.EntityData.Leafs.Append("ia-prefix-options", types.YLeaf{"IaPrefixOptions", prefix.IaPrefixOptions})

    prefix.EntityData.YListKeys = []string {}

    return &(prefix.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_IaRouter
// OSPFv3 inter area router
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_IaRouter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Metric. The type is interface{} with range: 0..4294967295.
    Metric interface{}

    // Router ID of the router being described by the LSA. The type is interface{}
    // with range: 0..4294967295.
    DestinationRouterId interface{}

    // Inter area LSA options. The type is map[string]bool.
    LsaIaOptions interface{}
}

func (iaRouter *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_IaRouter) GetEntityData() *types.CommonEntityData {
    iaRouter.EntityData.YFilter = iaRouter.YFilter
    iaRouter.EntityData.YangName = "ia-router"
    iaRouter.EntityData.BundleName = "cisco_ios_xe"
    iaRouter.EntityData.ParentYangName = "lsa-body"
    iaRouter.EntityData.SegmentPath = "ia-router"
    iaRouter.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/intf-link-scope-lsas/link-scope-lsa/ospfv3-lsa-val/lsa-body/" + iaRouter.EntityData.SegmentPath
    iaRouter.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    iaRouter.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    iaRouter.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    iaRouter.EntityData.Children = types.NewOrderedMap()
    iaRouter.EntityData.Leafs = types.NewOrderedMap()
    iaRouter.EntityData.Leafs.Append("metric", types.YLeaf{"Metric", iaRouter.Metric})
    iaRouter.EntityData.Leafs.Append("destination-router-id", types.YLeaf{"DestinationRouterId", iaRouter.DestinationRouterId})
    iaRouter.EntityData.Leafs.Append("lsa-ia-options", types.YLeaf{"LsaIaOptions", iaRouter.LsaIaOptions})

    iaRouter.EntityData.YListKeys = []string {}

    return &(iaRouter.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_LsaExternal
// OSPFv3 LSA external
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_LsaExternal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Metric. The type is interface{} with range: 0..4294967295.
    Metric interface{}

    // Referenced Link State type. The type is interface{} with range: 0..65535.
    ReferencedLsType interface{}

    // Prefix. The type is string.
    ExternalPrefix interface{}

    // Prefix options. The type is string.
    ExternalPrefixOptions interface{}

    // Forwarding address. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ForwardingAddress interface{}

    // Route tag. The type is interface{} with range: 0..4294967295.
    ExternalRouteTag interface{}

    // Referenced Link State ID. The type is interface{} with range:
    // 0..4294967295.
    ReferencedLinkStateId interface{}

    // LSA Flags.
    Flags OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_LsaExternal_Flags
}

func (lsaExternal *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_LsaExternal) GetEntityData() *types.CommonEntityData {
    lsaExternal.EntityData.YFilter = lsaExternal.YFilter
    lsaExternal.EntityData.YangName = "lsa-external"
    lsaExternal.EntityData.BundleName = "cisco_ios_xe"
    lsaExternal.EntityData.ParentYangName = "lsa-body"
    lsaExternal.EntityData.SegmentPath = "lsa-external"
    lsaExternal.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/intf-link-scope-lsas/link-scope-lsa/ospfv3-lsa-val/lsa-body/" + lsaExternal.EntityData.SegmentPath
    lsaExternal.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lsaExternal.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lsaExternal.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lsaExternal.EntityData.Children = types.NewOrderedMap()
    lsaExternal.EntityData.Children.Append("flags", types.YChild{"Flags", &lsaExternal.Flags})
    lsaExternal.EntityData.Leafs = types.NewOrderedMap()
    lsaExternal.EntityData.Leafs.Append("metric", types.YLeaf{"Metric", lsaExternal.Metric})
    lsaExternal.EntityData.Leafs.Append("referenced-ls-type", types.YLeaf{"ReferencedLsType", lsaExternal.ReferencedLsType})
    lsaExternal.EntityData.Leafs.Append("external-prefix", types.YLeaf{"ExternalPrefix", lsaExternal.ExternalPrefix})
    lsaExternal.EntityData.Leafs.Append("external-prefix-options", types.YLeaf{"ExternalPrefixOptions", lsaExternal.ExternalPrefixOptions})
    lsaExternal.EntityData.Leafs.Append("forwarding-address", types.YLeaf{"ForwardingAddress", lsaExternal.ForwardingAddress})
    lsaExternal.EntityData.Leafs.Append("external-route-tag", types.YLeaf{"ExternalRouteTag", lsaExternal.ExternalRouteTag})
    lsaExternal.EntityData.Leafs.Append("referenced-link-state-id", types.YLeaf{"ReferencedLinkStateId", lsaExternal.ReferencedLinkStateId})

    lsaExternal.EntityData.YListKeys = []string {}

    return &(lsaExternal.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_LsaExternal_Flags
// LSA Flags
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_LsaExternal_Flags struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // When set, the metric specified is a Type 2 external metric. The type is
    // bool.
    EFlag interface{}
}

func (flags *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_LsaExternal_Flags) GetEntityData() *types.CommonEntityData {
    flags.EntityData.YFilter = flags.YFilter
    flags.EntityData.YangName = "flags"
    flags.EntityData.BundleName = "cisco_ios_xe"
    flags.EntityData.ParentYangName = "lsa-external"
    flags.EntityData.SegmentPath = "flags"
    flags.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/intf-link-scope-lsas/link-scope-lsa/ospfv3-lsa-val/lsa-body/lsa-external/" + flags.EntityData.SegmentPath
    flags.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    flags.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    flags.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    flags.EntityData.Children = types.NewOrderedMap()
    flags.EntityData.Leafs = types.NewOrderedMap()
    flags.EntityData.Leafs.Append("e-flag", types.YLeaf{"EFlag", flags.EFlag})

    flags.EntityData.YListKeys = []string {}

    return &(flags.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Nssa
// OSPFv3 NSSA
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Nssa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // NSSA LSA.
    LsaNssaExternal OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Nssa_LsaNssaExternal
}

func (nssa *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Nssa) GetEntityData() *types.CommonEntityData {
    nssa.EntityData.YFilter = nssa.YFilter
    nssa.EntityData.YangName = "nssa"
    nssa.EntityData.BundleName = "cisco_ios_xe"
    nssa.EntityData.ParentYangName = "lsa-body"
    nssa.EntityData.SegmentPath = "nssa"
    nssa.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/intf-link-scope-lsas/link-scope-lsa/ospfv3-lsa-val/lsa-body/" + nssa.EntityData.SegmentPath
    nssa.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nssa.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nssa.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nssa.EntityData.Children = types.NewOrderedMap()
    nssa.EntityData.Children.Append("lsa-nssa-external", types.YChild{"LsaNssaExternal", &nssa.LsaNssaExternal})
    nssa.EntityData.Leafs = types.NewOrderedMap()

    nssa.EntityData.YListKeys = []string {}

    return &(nssa.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Nssa_LsaNssaExternal
// NSSA LSA
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Nssa_LsaNssaExternal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Metric. The type is interface{} with range: 0..4294967295.
    Metric interface{}

    // Referenced Link State type. The type is interface{} with range: 0..65535.
    ReferencedLsType interface{}

    // Prefix. The type is string.
    ExternalPrefix interface{}

    // Prefix options. The type is string.
    ExternalPrefixOptions interface{}

    // Forwarding address. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ForwardingAddress interface{}

    // Route tag. The type is interface{} with range: 0..4294967295.
    ExternalRouteTag interface{}

    // Referenced Link State ID. The type is interface{} with range:
    // 0..4294967295.
    ReferencedLinkStateId interface{}

    // LSA Flags.
    Flags OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Nssa_LsaNssaExternal_Flags
}

func (lsaNssaExternal *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Nssa_LsaNssaExternal) GetEntityData() *types.CommonEntityData {
    lsaNssaExternal.EntityData.YFilter = lsaNssaExternal.YFilter
    lsaNssaExternal.EntityData.YangName = "lsa-nssa-external"
    lsaNssaExternal.EntityData.BundleName = "cisco_ios_xe"
    lsaNssaExternal.EntityData.ParentYangName = "nssa"
    lsaNssaExternal.EntityData.SegmentPath = "lsa-nssa-external"
    lsaNssaExternal.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/intf-link-scope-lsas/link-scope-lsa/ospfv3-lsa-val/lsa-body/nssa/" + lsaNssaExternal.EntityData.SegmentPath
    lsaNssaExternal.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lsaNssaExternal.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lsaNssaExternal.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lsaNssaExternal.EntityData.Children = types.NewOrderedMap()
    lsaNssaExternal.EntityData.Children.Append("flags", types.YChild{"Flags", &lsaNssaExternal.Flags})
    lsaNssaExternal.EntityData.Leafs = types.NewOrderedMap()
    lsaNssaExternal.EntityData.Leafs.Append("metric", types.YLeaf{"Metric", lsaNssaExternal.Metric})
    lsaNssaExternal.EntityData.Leafs.Append("referenced-ls-type", types.YLeaf{"ReferencedLsType", lsaNssaExternal.ReferencedLsType})
    lsaNssaExternal.EntityData.Leafs.Append("external-prefix", types.YLeaf{"ExternalPrefix", lsaNssaExternal.ExternalPrefix})
    lsaNssaExternal.EntityData.Leafs.Append("external-prefix-options", types.YLeaf{"ExternalPrefixOptions", lsaNssaExternal.ExternalPrefixOptions})
    lsaNssaExternal.EntityData.Leafs.Append("forwarding-address", types.YLeaf{"ForwardingAddress", lsaNssaExternal.ForwardingAddress})
    lsaNssaExternal.EntityData.Leafs.Append("external-route-tag", types.YLeaf{"ExternalRouteTag", lsaNssaExternal.ExternalRouteTag})
    lsaNssaExternal.EntityData.Leafs.Append("referenced-link-state-id", types.YLeaf{"ReferencedLinkStateId", lsaNssaExternal.ReferencedLinkStateId})

    lsaNssaExternal.EntityData.YListKeys = []string {}

    return &(lsaNssaExternal.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Nssa_LsaNssaExternal_Flags
// LSA Flags
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Nssa_LsaNssaExternal_Flags struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // When set, the metric specified is a Type 2 external metric. The type is
    // bool.
    EFlag interface{}
}

func (flags *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Nssa_LsaNssaExternal_Flags) GetEntityData() *types.CommonEntityData {
    flags.EntityData.YFilter = flags.YFilter
    flags.EntityData.YangName = "flags"
    flags.EntityData.BundleName = "cisco_ios_xe"
    flags.EntityData.ParentYangName = "lsa-nssa-external"
    flags.EntityData.SegmentPath = "flags"
    flags.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/intf-link-scope-lsas/link-scope-lsa/ospfv3-lsa-val/lsa-body/nssa/lsa-nssa-external/" + flags.EntityData.SegmentPath
    flags.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    flags.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    flags.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    flags.EntityData.Children = types.NewOrderedMap()
    flags.EntityData.Leafs = types.NewOrderedMap()
    flags.EntityData.Leafs.Append("e-flag", types.YLeaf{"EFlag", flags.EFlag})

    flags.EntityData.YListKeys = []string {}

    return &(flags.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_LinkData
// OSPFv3 Link data
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_LinkData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Router priority of the interce. The type is interface{} with range: 0..255.
    RtrPriority interface{}

    // The originating router's link-local interface address on the link. The type
    // is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    LinkLocalInterfaceAddress interface{}

    // Number of prefixes. The type is interface{} with range: 0..4294967295.
    NumOfPrefixes interface{}

    // Link data LSA options. The type is map[string]bool.
    LsaIdOptions interface{}
}

func (linkData *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_LinkData) GetEntityData() *types.CommonEntityData {
    linkData.EntityData.YFilter = linkData.YFilter
    linkData.EntityData.YangName = "link-data"
    linkData.EntityData.BundleName = "cisco_ios_xe"
    linkData.EntityData.ParentYangName = "lsa-body"
    linkData.EntityData.SegmentPath = "link-data"
    linkData.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/intf-link-scope-lsas/link-scope-lsa/ospfv3-lsa-val/lsa-body/" + linkData.EntityData.SegmentPath
    linkData.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    linkData.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    linkData.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    linkData.EntityData.Children = types.NewOrderedMap()
    linkData.EntityData.Leafs = types.NewOrderedMap()
    linkData.EntityData.Leafs.Append("rtr-priority", types.YLeaf{"RtrPriority", linkData.RtrPriority})
    linkData.EntityData.Leafs.Append("link-local-interface-address", types.YLeaf{"LinkLocalInterfaceAddress", linkData.LinkLocalInterfaceAddress})
    linkData.EntityData.Leafs.Append("num-of-prefixes", types.YLeaf{"NumOfPrefixes", linkData.NumOfPrefixes})
    linkData.EntityData.Leafs.Append("lsa-id-options", types.YLeaf{"LsaIdOptions", linkData.LsaIdOptions})

    linkData.EntityData.YListKeys = []string {}

    return &(linkData.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_IaPrefix
// OSPFv3 Intra area prefixes
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_IaPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Referenced Link State type. The type is interface{} with range: 0..65535.
    ReferencedLsType interface{}

    // Referenced Link State ID. The type is interface{} with range:
    // 0..4294967295.
    ReferencedLinkStateId interface{}

    // Referenced Advertising Router. The type is one of the following types:
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ReferencedAdvRouter interface{}

    // Number of prefixes. The type is interface{} with range: 0..65535.
    NumOfPrefixes interface{}
}

func (iaPrefix *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_IaPrefix) GetEntityData() *types.CommonEntityData {
    iaPrefix.EntityData.YFilter = iaPrefix.YFilter
    iaPrefix.EntityData.YangName = "ia-prefix"
    iaPrefix.EntityData.BundleName = "cisco_ios_xe"
    iaPrefix.EntityData.ParentYangName = "lsa-body"
    iaPrefix.EntityData.SegmentPath = "ia-prefix"
    iaPrefix.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/intf-link-scope-lsas/link-scope-lsa/ospfv3-lsa-val/lsa-body/" + iaPrefix.EntityData.SegmentPath
    iaPrefix.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    iaPrefix.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    iaPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    iaPrefix.EntityData.Children = types.NewOrderedMap()
    iaPrefix.EntityData.Leafs = types.NewOrderedMap()
    iaPrefix.EntityData.Leafs.Append("referenced-ls-type", types.YLeaf{"ReferencedLsType", iaPrefix.ReferencedLsType})
    iaPrefix.EntityData.Leafs.Append("referenced-link-state-id", types.YLeaf{"ReferencedLinkStateId", iaPrefix.ReferencedLinkStateId})
    iaPrefix.EntityData.Leafs.Append("referenced-adv-router", types.YLeaf{"ReferencedAdvRouter", iaPrefix.ReferencedAdvRouter})
    iaPrefix.EntityData.Leafs.Append("num-of-prefixes", types.YLeaf{"NumOfPrefixes", iaPrefix.NumOfPrefixes})

    iaPrefix.EntityData.YListKeys = []string {}

    return &(iaPrefix.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3Link
// OSPFv3 links
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3Link struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface ID. The type is interface{} with range:
    // 0..4294967295.
    InterfaceId interface{}

    // This attribute is a key. Neighbor interface ID. The type is interface{}
    // with range: 0..4294967295.
    NeighborInterfaceId interface{}

    // This attribute is a key. Neighbor router ID. The type is interface{} with
    // range: 0..4294967295.
    NeighborRouterId interface{}

    // Link type. The type is interface{} with range: 0..255.
    Type interface{}

    // Metric. The type is interface{} with range: 0..65535.
    Metric interface{}
}

func (ospfv3Link *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3Link) GetEntityData() *types.CommonEntityData {
    ospfv3Link.EntityData.YFilter = ospfv3Link.YFilter
    ospfv3Link.EntityData.YangName = "ospfv3-link"
    ospfv3Link.EntityData.BundleName = "cisco_ios_xe"
    ospfv3Link.EntityData.ParentYangName = "link-scope-lsa"
    ospfv3Link.EntityData.SegmentPath = "ospfv3-link" + types.AddKeyToken(ospfv3Link.InterfaceId, "interface-id") + types.AddKeyToken(ospfv3Link.NeighborInterfaceId, "neighbor-interface-id") + types.AddKeyToken(ospfv3Link.NeighborRouterId, "neighbor-router-id")
    ospfv3Link.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/intf-link-scope-lsas/link-scope-lsa/" + ospfv3Link.EntityData.SegmentPath
    ospfv3Link.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv3Link.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv3Link.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv3Link.EntityData.Children = types.NewOrderedMap()
    ospfv3Link.EntityData.Leafs = types.NewOrderedMap()
    ospfv3Link.EntityData.Leafs.Append("interface-id", types.YLeaf{"InterfaceId", ospfv3Link.InterfaceId})
    ospfv3Link.EntityData.Leafs.Append("neighbor-interface-id", types.YLeaf{"NeighborInterfaceId", ospfv3Link.NeighborInterfaceId})
    ospfv3Link.EntityData.Leafs.Append("neighbor-router-id", types.YLeaf{"NeighborRouterId", ospfv3Link.NeighborRouterId})
    ospfv3Link.EntityData.Leafs.Append("type", types.YLeaf{"Type", ospfv3Link.Type})
    ospfv3Link.EntityData.Leafs.Append("metric", types.YLeaf{"Metric", ospfv3Link.Metric})

    ospfv3Link.EntityData.YListKeys = []string {"InterfaceId", "NeighborInterfaceId", "NeighborRouterId"}

    return &(ospfv3Link.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3PrefixList
// OSPFv3 prefix-list
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3PrefixList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Prefix. The type is string.
    Prefix interface{}

    // Prefix options. The type is string.
    PrefixOptions interface{}
}

func (ospfv3PrefixList *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3PrefixList) GetEntityData() *types.CommonEntityData {
    ospfv3PrefixList.EntityData.YFilter = ospfv3PrefixList.YFilter
    ospfv3PrefixList.EntityData.YangName = "ospfv3-prefix-list"
    ospfv3PrefixList.EntityData.BundleName = "cisco_ios_xe"
    ospfv3PrefixList.EntityData.ParentYangName = "link-scope-lsa"
    ospfv3PrefixList.EntityData.SegmentPath = "ospfv3-prefix-list" + types.AddKeyToken(ospfv3PrefixList.Prefix, "prefix")
    ospfv3PrefixList.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/intf-link-scope-lsas/link-scope-lsa/" + ospfv3PrefixList.EntityData.SegmentPath
    ospfv3PrefixList.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv3PrefixList.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv3PrefixList.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv3PrefixList.EntityData.Children = types.NewOrderedMap()
    ospfv3PrefixList.EntityData.Leafs = types.NewOrderedMap()
    ospfv3PrefixList.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", ospfv3PrefixList.Prefix})
    ospfv3PrefixList.EntityData.Leafs.Append("prefix-options", types.YLeaf{"PrefixOptions", ospfv3PrefixList.PrefixOptions})

    ospfv3PrefixList.EntityData.YListKeys = []string {"Prefix"}

    return &(ospfv3PrefixList.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3IaPrefix
// OSPFv3 intra-area prefix-list
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3IaPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Prefix. The type is string.
    Prefix interface{}

    // Prefix options. The type is string.
    PrefixOptions interface{}
}

func (ospfv3IaPrefix *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3IaPrefix) GetEntityData() *types.CommonEntityData {
    ospfv3IaPrefix.EntityData.YFilter = ospfv3IaPrefix.YFilter
    ospfv3IaPrefix.EntityData.YangName = "ospfv3-ia-prefix"
    ospfv3IaPrefix.EntityData.BundleName = "cisco_ios_xe"
    ospfv3IaPrefix.EntityData.ParentYangName = "link-scope-lsa"
    ospfv3IaPrefix.EntityData.SegmentPath = "ospfv3-ia-prefix" + types.AddKeyToken(ospfv3IaPrefix.Prefix, "prefix")
    ospfv3IaPrefix.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/intf-link-scope-lsas/link-scope-lsa/" + ospfv3IaPrefix.EntityData.SegmentPath
    ospfv3IaPrefix.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv3IaPrefix.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv3IaPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv3IaPrefix.EntityData.Children = types.NewOrderedMap()
    ospfv3IaPrefix.EntityData.Leafs = types.NewOrderedMap()
    ospfv3IaPrefix.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", ospfv3IaPrefix.Prefix})
    ospfv3IaPrefix.EntityData.Leafs.Append("prefix-options", types.YLeaf{"PrefixOptions", ospfv3IaPrefix.PrefixOptions})

    ospfv3IaPrefix.EntityData.YListKeys = []string {"Prefix"}

    return &(ospfv3IaPrefix.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_MultiTopology
// OSPF multi-topology interface augmentation
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_MultiTopology struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. One of the topology enabled on this interface. The
    // type is string.
    Name interface{}
}

func (multiTopology *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_MultiTopology) GetEntityData() *types.CommonEntityData {
    multiTopology.EntityData.YFilter = multiTopology.YFilter
    multiTopology.EntityData.YangName = "multi-topology"
    multiTopology.EntityData.BundleName = "cisco_ios_xe"
    multiTopology.EntityData.ParentYangName = "link-scope-lsa"
    multiTopology.EntityData.SegmentPath = "multi-topology" + types.AddKeyToken(multiTopology.Name, "name")
    multiTopology.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/intf-link-scope-lsas/link-scope-lsa/" + multiTopology.EntityData.SegmentPath
    multiTopology.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    multiTopology.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    multiTopology.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    multiTopology.EntityData.Children = types.NewOrderedMap()
    multiTopology.EntityData.Leafs = types.NewOrderedMap()
    multiTopology.EntityData.Leafs.Append("name", types.YLeaf{"Name", multiTopology.Name})

    multiTopology.EntityData.YListKeys = []string {"Name"}

    return &(multiTopology.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Tlv
// Link TLV
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Tlv struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Link type. The type is interface{} with range: 0..255.
    LinkType interface{}

    // Link ID. The type is interface{} with range: 0..4294967295.
    LinkId interface{}

    // List of local interface IPv4 addresses. The type is one of the following
    // types: slice of string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or slice of string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    LocalIfIpv4Addr []interface{}

    // List of remote interface IPv4 addresses. The type is one of the following
    // types: slice of string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or slice of string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    LocalRemoteIpv4Addr []interface{}

    // TE metric. The type is interface{} with range: 0..4294967295.
    TeMetric interface{}

    // Maximum bandwidth. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    MaxBandwidth interface{}

    // Maximum reservable bandwidth. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    MaxReservableBandwidth interface{}

    // Unrseerved bandwidth. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    UnreservedBandwidth interface{}

    // Administrative group/Resource class/Color. The type is interface{} with
    // range: 0..4294967295.
    AdminGroup interface{}
}

func (tlv *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Tlv) GetEntityData() *types.CommonEntityData {
    tlv.EntityData.YFilter = tlv.YFilter
    tlv.EntityData.YangName = "tlv"
    tlv.EntityData.BundleName = "cisco_ios_xe"
    tlv.EntityData.ParentYangName = "link-scope-lsa"
    tlv.EntityData.SegmentPath = "tlv"
    tlv.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/intf-link-scope-lsas/link-scope-lsa/" + tlv.EntityData.SegmentPath
    tlv.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    tlv.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    tlv.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    tlv.EntityData.Children = types.NewOrderedMap()
    tlv.EntityData.Leafs = types.NewOrderedMap()
    tlv.EntityData.Leafs.Append("link-type", types.YLeaf{"LinkType", tlv.LinkType})
    tlv.EntityData.Leafs.Append("link-id", types.YLeaf{"LinkId", tlv.LinkId})
    tlv.EntityData.Leafs.Append("local-if-ipv4-addr", types.YLeaf{"LocalIfIpv4Addr", tlv.LocalIfIpv4Addr})
    tlv.EntityData.Leafs.Append("local-remote-ipv4-addr", types.YLeaf{"LocalRemoteIpv4Addr", tlv.LocalRemoteIpv4Addr})
    tlv.EntityData.Leafs.Append("te-metric", types.YLeaf{"TeMetric", tlv.TeMetric})
    tlv.EntityData.Leafs.Append("max-bandwidth", types.YLeaf{"MaxBandwidth", tlv.MaxBandwidth})
    tlv.EntityData.Leafs.Append("max-reservable-bandwidth", types.YLeaf{"MaxReservableBandwidth", tlv.MaxReservableBandwidth})
    tlv.EntityData.Leafs.Append("unreserved-bandwidth", types.YLeaf{"UnreservedBandwidth", tlv.UnreservedBandwidth})
    tlv.EntityData.Leafs.Append("admin-group", types.YLeaf{"AdminGroup", tlv.AdminGroup})

    tlv.EntityData.YListKeys = []string {}

    return &(tlv.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_UnknownSubTlv
// OSPFv2 Unknown sub TLV
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_UnknownSubTlv struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. TLV type. The type is interface{} with range:
    // 0..65535.
    Type interface{}

    // TLV length. The type is interface{} with range: 0..65535.
    Length interface{}

    // TLV value. The type is slice of interface{} with range: 0..255.
    Value []interface{}
}

func (unknownSubTlv *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_UnknownSubTlv) GetEntityData() *types.CommonEntityData {
    unknownSubTlv.EntityData.YFilter = unknownSubTlv.YFilter
    unknownSubTlv.EntityData.YangName = "unknown-sub-tlv"
    unknownSubTlv.EntityData.BundleName = "cisco_ios_xe"
    unknownSubTlv.EntityData.ParentYangName = "link-scope-lsa"
    unknownSubTlv.EntityData.SegmentPath = "unknown-sub-tlv" + types.AddKeyToken(unknownSubTlv.Type, "type")
    unknownSubTlv.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/intf-link-scope-lsas/link-scope-lsa/" + unknownSubTlv.EntityData.SegmentPath
    unknownSubTlv.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    unknownSubTlv.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    unknownSubTlv.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    unknownSubTlv.EntityData.Children = types.NewOrderedMap()
    unknownSubTlv.EntityData.Leafs = types.NewOrderedMap()
    unknownSubTlv.EntityData.Leafs.Append("type", types.YLeaf{"Type", unknownSubTlv.Type})
    unknownSubTlv.EntityData.Leafs.Append("length", types.YLeaf{"Length", unknownSubTlv.Length})
    unknownSubTlv.EntityData.Leafs.Append("value", types.YLeaf{"Value", unknownSubTlv.Value})

    unknownSubTlv.EntityData.YListKeys = []string {"Type"}

    return &(unknownSubTlv.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa
// List OSPF area scope LSA databases
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. LSA Type. The type is interface{} with range:
    // 0..4294967295.
    LsaType interface{}

    // This attribute is a key. Advertising router. The type is one of the
    // following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    AdvRouter interface{}

    // The OSPF LSA body is fully decoded. The type is bool.
    DecodedCompleted interface{}

    // The complete LSA in network byte order as received/sent over the wire. The
    // type is slice of interface{} with range: 0..255.
    RawData []interface{}

    // OSPFv2 LSA.
    Ospfv2Lsa OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2Lsa

    // Router LSA link. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2Link.
    Ospfv2Link []*OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2Link

    // Summary LSA. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2Topology.
    Ospfv2Topology []*OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2Topology

    // External LSA. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2External.
    Ospfv2External []*OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2External

    // OSPFv3 LSA.
    Ospfv3Lsa OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa

    // OSPFv3 links. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Link.
    Ospfv3Link []*OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Link

    // OSPFv3 prefix-list. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Prefix.
    Ospfv3Prefix []*OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Prefix

    // OSPFv3 intra-area prefix-list. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3IaPrefix.
    Ospfv3IaPrefix []*OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3IaPrefix
}

func (areaScopeLsa *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa) GetEntityData() *types.CommonEntityData {
    areaScopeLsa.EntityData.YFilter = areaScopeLsa.YFilter
    areaScopeLsa.EntityData.YangName = "area-scope-lsa"
    areaScopeLsa.EntityData.BundleName = "cisco_ios_xe"
    areaScopeLsa.EntityData.ParentYangName = "intf-link-scope-lsas"
    areaScopeLsa.EntityData.SegmentPath = "area-scope-lsa" + types.AddKeyToken(areaScopeLsa.LsaType, "lsa-type") + types.AddKeyToken(areaScopeLsa.AdvRouter, "adv-router")
    areaScopeLsa.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/intf-link-scope-lsas/" + areaScopeLsa.EntityData.SegmentPath
    areaScopeLsa.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    areaScopeLsa.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    areaScopeLsa.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    areaScopeLsa.EntityData.Children = types.NewOrderedMap()
    areaScopeLsa.EntityData.Children.Append("ospfv2-lsa", types.YChild{"Ospfv2Lsa", &areaScopeLsa.Ospfv2Lsa})
    areaScopeLsa.EntityData.Children.Append("ospfv2-link", types.YChild{"Ospfv2Link", nil})
    for i := range areaScopeLsa.Ospfv2Link {
        areaScopeLsa.EntityData.Children.Append(types.GetSegmentPath(areaScopeLsa.Ospfv2Link[i]), types.YChild{"Ospfv2Link", areaScopeLsa.Ospfv2Link[i]})
    }
    areaScopeLsa.EntityData.Children.Append("ospfv2-topology", types.YChild{"Ospfv2Topology", nil})
    for i := range areaScopeLsa.Ospfv2Topology {
        areaScopeLsa.EntityData.Children.Append(types.GetSegmentPath(areaScopeLsa.Ospfv2Topology[i]), types.YChild{"Ospfv2Topology", areaScopeLsa.Ospfv2Topology[i]})
    }
    areaScopeLsa.EntityData.Children.Append("ospfv2-external", types.YChild{"Ospfv2External", nil})
    for i := range areaScopeLsa.Ospfv2External {
        areaScopeLsa.EntityData.Children.Append(types.GetSegmentPath(areaScopeLsa.Ospfv2External[i]), types.YChild{"Ospfv2External", areaScopeLsa.Ospfv2External[i]})
    }
    areaScopeLsa.EntityData.Children.Append("ospfv3-lsa", types.YChild{"Ospfv3Lsa", &areaScopeLsa.Ospfv3Lsa})
    areaScopeLsa.EntityData.Children.Append("ospfv3-link", types.YChild{"Ospfv3Link", nil})
    for i := range areaScopeLsa.Ospfv3Link {
        areaScopeLsa.EntityData.Children.Append(types.GetSegmentPath(areaScopeLsa.Ospfv3Link[i]), types.YChild{"Ospfv3Link", areaScopeLsa.Ospfv3Link[i]})
    }
    areaScopeLsa.EntityData.Children.Append("ospfv3-prefix", types.YChild{"Ospfv3Prefix", nil})
    for i := range areaScopeLsa.Ospfv3Prefix {
        areaScopeLsa.EntityData.Children.Append(types.GetSegmentPath(areaScopeLsa.Ospfv3Prefix[i]), types.YChild{"Ospfv3Prefix", areaScopeLsa.Ospfv3Prefix[i]})
    }
    areaScopeLsa.EntityData.Children.Append("ospfv3-ia-prefix", types.YChild{"Ospfv3IaPrefix", nil})
    for i := range areaScopeLsa.Ospfv3IaPrefix {
        areaScopeLsa.EntityData.Children.Append(types.GetSegmentPath(areaScopeLsa.Ospfv3IaPrefix[i]), types.YChild{"Ospfv3IaPrefix", areaScopeLsa.Ospfv3IaPrefix[i]})
    }
    areaScopeLsa.EntityData.Leafs = types.NewOrderedMap()
    areaScopeLsa.EntityData.Leafs.Append("lsa-type", types.YLeaf{"LsaType", areaScopeLsa.LsaType})
    areaScopeLsa.EntityData.Leafs.Append("adv-router", types.YLeaf{"AdvRouter", areaScopeLsa.AdvRouter})
    areaScopeLsa.EntityData.Leafs.Append("decoded-completed", types.YLeaf{"DecodedCompleted", areaScopeLsa.DecodedCompleted})
    areaScopeLsa.EntityData.Leafs.Append("raw-data", types.YLeaf{"RawData", areaScopeLsa.RawData})

    areaScopeLsa.EntityData.YListKeys = []string {"LsaType", "AdvRouter"}

    return &(areaScopeLsa.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2Lsa
// OSPFv2 LSA
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2Lsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Decoded OSPFv2 LSA header data.
    Header OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2Lsa_Header

    // Decoded OSPFv2 LSA body data.
    LsaBody OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2Lsa_LsaBody
}

func (ospfv2Lsa *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2Lsa) GetEntityData() *types.CommonEntityData {
    ospfv2Lsa.EntityData.YFilter = ospfv2Lsa.YFilter
    ospfv2Lsa.EntityData.YangName = "ospfv2-lsa"
    ospfv2Lsa.EntityData.BundleName = "cisco_ios_xe"
    ospfv2Lsa.EntityData.ParentYangName = "area-scope-lsa"
    ospfv2Lsa.EntityData.SegmentPath = "ospfv2-lsa"
    ospfv2Lsa.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/intf-link-scope-lsas/area-scope-lsa/" + ospfv2Lsa.EntityData.SegmentPath
    ospfv2Lsa.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv2Lsa.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv2Lsa.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv2Lsa.EntityData.Children = types.NewOrderedMap()
    ospfv2Lsa.EntityData.Children.Append("header", types.YChild{"Header", &ospfv2Lsa.Header})
    ospfv2Lsa.EntityData.Children.Append("lsa-body", types.YChild{"LsaBody", &ospfv2Lsa.LsaBody})
    ospfv2Lsa.EntityData.Leafs = types.NewOrderedMap()

    ospfv2Lsa.EntityData.YListKeys = []string {}

    return &(ospfv2Lsa.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2Lsa_Header
// Decoded OSPFv2 LSA header data
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2Lsa_Header struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA ID. The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    LsaId interface{}

    // Opaque type. The type is interface{} with range: 0..255.
    OpaqueType interface{}

    // Opaque ID. The type is interface{} with range: 0..4294967295.
    OpaqueId interface{}

    // LSA age. The type is interface{} with range: 0..65535.
    Age interface{}

    // LSA type. The type is interface{} with range: 0..65535.
    Type interface{}

    // LSA advertising router. The type is interface{} with range: 0..4294967295.
    AdvRouter interface{}

    // LSA sequence number. The type is string.
    SeqNum interface{}

    // LSA checksum. The type is string.
    Checksum interface{}

    // LSA length. The type is interface{} with range: 0..65535.
    Length interface{}

    // LSA options. The type is map[string]bool.
    FlagOptions interface{}
}

func (header *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2Lsa_Header) GetEntityData() *types.CommonEntityData {
    header.EntityData.YFilter = header.YFilter
    header.EntityData.YangName = "header"
    header.EntityData.BundleName = "cisco_ios_xe"
    header.EntityData.ParentYangName = "ospfv2-lsa"
    header.EntityData.SegmentPath = "header"
    header.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/intf-link-scope-lsas/area-scope-lsa/ospfv2-lsa/" + header.EntityData.SegmentPath
    header.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    header.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    header.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    header.EntityData.Children = types.NewOrderedMap()
    header.EntityData.Leafs = types.NewOrderedMap()
    header.EntityData.Leafs.Append("lsa-id", types.YLeaf{"LsaId", header.LsaId})
    header.EntityData.Leafs.Append("opaque-type", types.YLeaf{"OpaqueType", header.OpaqueType})
    header.EntityData.Leafs.Append("opaque-id", types.YLeaf{"OpaqueId", header.OpaqueId})
    header.EntityData.Leafs.Append("age", types.YLeaf{"Age", header.Age})
    header.EntityData.Leafs.Append("type", types.YLeaf{"Type", header.Type})
    header.EntityData.Leafs.Append("adv-router", types.YLeaf{"AdvRouter", header.AdvRouter})
    header.EntityData.Leafs.Append("seq-num", types.YLeaf{"SeqNum", header.SeqNum})
    header.EntityData.Leafs.Append("checksum", types.YLeaf{"Checksum", header.Checksum})
    header.EntityData.Leafs.Append("length", types.YLeaf{"Length", header.Length})
    header.EntityData.Leafs.Append("flag-options", types.YLeaf{"FlagOptions", header.FlagOptions})

    header.EntityData.YListKeys = []string {}

    return &(header.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2Lsa_LsaBody
// Decoded OSPFv2 LSA body data
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2Lsa_LsaBody struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of links. The type is interface{} with range: 0..65535.
    NumOfLinks interface{}

    // Summary mask. The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    SummaryMask interface{}

    // External mask. The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ExternalMask interface{}

    // LSA body flags. The type is map[string]bool.
    BodyFlagOptions interface{}

    // Network details.
    Network OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2Lsa_LsaBody_Network
}

func (lsaBody *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2Lsa_LsaBody) GetEntityData() *types.CommonEntityData {
    lsaBody.EntityData.YFilter = lsaBody.YFilter
    lsaBody.EntityData.YangName = "lsa-body"
    lsaBody.EntityData.BundleName = "cisco_ios_xe"
    lsaBody.EntityData.ParentYangName = "ospfv2-lsa"
    lsaBody.EntityData.SegmentPath = "lsa-body"
    lsaBody.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/intf-link-scope-lsas/area-scope-lsa/ospfv2-lsa/" + lsaBody.EntityData.SegmentPath
    lsaBody.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lsaBody.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lsaBody.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lsaBody.EntityData.Children = types.NewOrderedMap()
    lsaBody.EntityData.Children.Append("network", types.YChild{"Network", &lsaBody.Network})
    lsaBody.EntityData.Leafs = types.NewOrderedMap()
    lsaBody.EntityData.Leafs.Append("num-of-links", types.YLeaf{"NumOfLinks", lsaBody.NumOfLinks})
    lsaBody.EntityData.Leafs.Append("summary-mask", types.YLeaf{"SummaryMask", lsaBody.SummaryMask})
    lsaBody.EntityData.Leafs.Append("external-mask", types.YLeaf{"ExternalMask", lsaBody.ExternalMask})
    lsaBody.EntityData.Leafs.Append("body-flag-options", types.YLeaf{"BodyFlagOptions", lsaBody.BodyFlagOptions})

    lsaBody.EntityData.YListKeys = []string {}

    return &(lsaBody.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2Lsa_LsaBody_Network
// Network details
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2Lsa_LsaBody_Network struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP network mask. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NetworkMask interface{}

    // List of the routers attached to the network. The type is slice of
    // interface{} with range: 0..4294967295.
    AttachedRouter []interface{}
}

func (network *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2Lsa_LsaBody_Network) GetEntityData() *types.CommonEntityData {
    network.EntityData.YFilter = network.YFilter
    network.EntityData.YangName = "network"
    network.EntityData.BundleName = "cisco_ios_xe"
    network.EntityData.ParentYangName = "lsa-body"
    network.EntityData.SegmentPath = "network"
    network.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/intf-link-scope-lsas/area-scope-lsa/ospfv2-lsa/lsa-body/" + network.EntityData.SegmentPath
    network.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    network.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    network.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    network.EntityData.Children = types.NewOrderedMap()
    network.EntityData.Leafs = types.NewOrderedMap()
    network.EntityData.Leafs.Append("network-mask", types.YLeaf{"NetworkMask", network.NetworkMask})
    network.EntityData.Leafs.Append("attached-router", types.YLeaf{"AttachedRouter", network.AttachedRouter})

    network.EntityData.YListKeys = []string {}

    return &(network.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2Link
// Router LSA link
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2Link struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Link ID. The type is interface{} with range:
    // 0..4294967295.
    LinkId interface{}

    // This attribute is a key. Link data. The type is interface{} with range:
    // 0..4294967295.
    LinkData interface{}

    // Link type. The type is interface{} with range: 0..255.
    Type interface{}

    // Topology specific information. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2Link_Ospfv2Topology.
    Ospfv2Topology []*OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2Link_Ospfv2Topology
}

func (ospfv2Link *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2Link) GetEntityData() *types.CommonEntityData {
    ospfv2Link.EntityData.YFilter = ospfv2Link.YFilter
    ospfv2Link.EntityData.YangName = "ospfv2-link"
    ospfv2Link.EntityData.BundleName = "cisco_ios_xe"
    ospfv2Link.EntityData.ParentYangName = "area-scope-lsa"
    ospfv2Link.EntityData.SegmentPath = "ospfv2-link" + types.AddKeyToken(ospfv2Link.LinkId, "link-id") + types.AddKeyToken(ospfv2Link.LinkData, "link-data")
    ospfv2Link.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/intf-link-scope-lsas/area-scope-lsa/" + ospfv2Link.EntityData.SegmentPath
    ospfv2Link.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv2Link.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv2Link.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv2Link.EntityData.Children = types.NewOrderedMap()
    ospfv2Link.EntityData.Children.Append("ospfv2-topology", types.YChild{"Ospfv2Topology", nil})
    for i := range ospfv2Link.Ospfv2Topology {
        ospfv2Link.EntityData.Children.Append(types.GetSegmentPath(ospfv2Link.Ospfv2Topology[i]), types.YChild{"Ospfv2Topology", ospfv2Link.Ospfv2Topology[i]})
    }
    ospfv2Link.EntityData.Leafs = types.NewOrderedMap()
    ospfv2Link.EntityData.Leafs.Append("link-id", types.YLeaf{"LinkId", ospfv2Link.LinkId})
    ospfv2Link.EntityData.Leafs.Append("link-data", types.YLeaf{"LinkData", ospfv2Link.LinkData})
    ospfv2Link.EntityData.Leafs.Append("type", types.YLeaf{"Type", ospfv2Link.Type})

    ospfv2Link.EntityData.YListKeys = []string {"LinkId", "LinkData"}

    return &(ospfv2Link.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2Link_Ospfv2Topology
// Topology specific information
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2Link_Ospfv2Topology struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. MT-ID for topology enabled link. The type is
    // interface{} with range: 0..4294967295.
    MtId interface{}

    // Metric for the topology. The type is interface{} with range: 0..65535.
    Metric interface{}
}

func (ospfv2Topology *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2Link_Ospfv2Topology) GetEntityData() *types.CommonEntityData {
    ospfv2Topology.EntityData.YFilter = ospfv2Topology.YFilter
    ospfv2Topology.EntityData.YangName = "ospfv2-topology"
    ospfv2Topology.EntityData.BundleName = "cisco_ios_xe"
    ospfv2Topology.EntityData.ParentYangName = "ospfv2-link"
    ospfv2Topology.EntityData.SegmentPath = "ospfv2-topology" + types.AddKeyToken(ospfv2Topology.MtId, "mt-id")
    ospfv2Topology.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/intf-link-scope-lsas/area-scope-lsa/ospfv2-link/" + ospfv2Topology.EntityData.SegmentPath
    ospfv2Topology.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv2Topology.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv2Topology.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv2Topology.EntityData.Children = types.NewOrderedMap()
    ospfv2Topology.EntityData.Leafs = types.NewOrderedMap()
    ospfv2Topology.EntityData.Leafs.Append("mt-id", types.YLeaf{"MtId", ospfv2Topology.MtId})
    ospfv2Topology.EntityData.Leafs.Append("metric", types.YLeaf{"Metric", ospfv2Topology.Metric})

    ospfv2Topology.EntityData.YListKeys = []string {"MtId"}

    return &(ospfv2Topology.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2Topology
// Summary LSA
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2Topology struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. MT-ID for topology enabled link. The type is
    // interface{} with range: 0..4294967295.
    MtId interface{}

    // Metric for the topology. The type is interface{} with range: 0..65535.
    Metric interface{}
}

func (ospfv2Topology *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2Topology) GetEntityData() *types.CommonEntityData {
    ospfv2Topology.EntityData.YFilter = ospfv2Topology.YFilter
    ospfv2Topology.EntityData.YangName = "ospfv2-topology"
    ospfv2Topology.EntityData.BundleName = "cisco_ios_xe"
    ospfv2Topology.EntityData.ParentYangName = "area-scope-lsa"
    ospfv2Topology.EntityData.SegmentPath = "ospfv2-topology" + types.AddKeyToken(ospfv2Topology.MtId, "mt-id")
    ospfv2Topology.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/intf-link-scope-lsas/area-scope-lsa/" + ospfv2Topology.EntityData.SegmentPath
    ospfv2Topology.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv2Topology.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv2Topology.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv2Topology.EntityData.Children = types.NewOrderedMap()
    ospfv2Topology.EntityData.Leafs = types.NewOrderedMap()
    ospfv2Topology.EntityData.Leafs.Append("mt-id", types.YLeaf{"MtId", ospfv2Topology.MtId})
    ospfv2Topology.EntityData.Leafs.Append("metric", types.YLeaf{"Metric", ospfv2Topology.Metric})

    ospfv2Topology.EntityData.YListKeys = []string {"MtId"}

    return &(ospfv2Topology.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2External
// External LSA
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2External struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. MT-ID for topology enabled on the link. The type
    // is interface{} with range: 0..4294967295.
    MtId interface{}

    // Metric for the topology. The type is interface{} with range: 0..4294967295.
    Metric interface{}

    // Forwarding address. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ForwardingAddress interface{}

    // Route tag. The type is interface{} with range: 0..4294967295.
    ExternalRouteTag interface{}
}

func (ospfv2External *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2External) GetEntityData() *types.CommonEntityData {
    ospfv2External.EntityData.YFilter = ospfv2External.YFilter
    ospfv2External.EntityData.YangName = "ospfv2-external"
    ospfv2External.EntityData.BundleName = "cisco_ios_xe"
    ospfv2External.EntityData.ParentYangName = "area-scope-lsa"
    ospfv2External.EntityData.SegmentPath = "ospfv2-external" + types.AddKeyToken(ospfv2External.MtId, "mt-id")
    ospfv2External.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/intf-link-scope-lsas/area-scope-lsa/" + ospfv2External.EntityData.SegmentPath
    ospfv2External.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv2External.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv2External.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv2External.EntityData.Children = types.NewOrderedMap()
    ospfv2External.EntityData.Leafs = types.NewOrderedMap()
    ospfv2External.EntityData.Leafs.Append("mt-id", types.YLeaf{"MtId", ospfv2External.MtId})
    ospfv2External.EntityData.Leafs.Append("metric", types.YLeaf{"Metric", ospfv2External.Metric})
    ospfv2External.EntityData.Leafs.Append("forwarding-address", types.YLeaf{"ForwardingAddress", ospfv2External.ForwardingAddress})
    ospfv2External.EntityData.Leafs.Append("external-route-tag", types.YLeaf{"ExternalRouteTag", ospfv2External.ExternalRouteTag})

    ospfv2External.EntityData.YListKeys = []string {"MtId"}

    return &(ospfv2External.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa
// OSPFv3 LSA
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Decoded OSPFv3 LSA header.
    Header OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_Header

    // Decoded OSPFv3 LSA body.
    LsaBody OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody
}

func (ospfv3Lsa *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa) GetEntityData() *types.CommonEntityData {
    ospfv3Lsa.EntityData.YFilter = ospfv3Lsa.YFilter
    ospfv3Lsa.EntityData.YangName = "ospfv3-lsa"
    ospfv3Lsa.EntityData.BundleName = "cisco_ios_xe"
    ospfv3Lsa.EntityData.ParentYangName = "area-scope-lsa"
    ospfv3Lsa.EntityData.SegmentPath = "ospfv3-lsa"
    ospfv3Lsa.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/intf-link-scope-lsas/area-scope-lsa/" + ospfv3Lsa.EntityData.SegmentPath
    ospfv3Lsa.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv3Lsa.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv3Lsa.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv3Lsa.EntityData.Children = types.NewOrderedMap()
    ospfv3Lsa.EntityData.Children.Append("header", types.YChild{"Header", &ospfv3Lsa.Header})
    ospfv3Lsa.EntityData.Children.Append("lsa-body", types.YChild{"LsaBody", &ospfv3Lsa.LsaBody})
    ospfv3Lsa.EntityData.Leafs = types.NewOrderedMap()

    ospfv3Lsa.EntityData.YListKeys = []string {}

    return &(ospfv3Lsa.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_Header
// Decoded OSPFv3 LSA header
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_Header struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA ID. The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    LsaId interface{}

    // OSPFv3 LSA options. The type is map[string]bool.
    LsaHdrOptions interface{}

    // LSA header.
    LsaHeader OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_Header_LsaHeader
}

func (header *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_Header) GetEntityData() *types.CommonEntityData {
    header.EntityData.YFilter = header.YFilter
    header.EntityData.YangName = "header"
    header.EntityData.BundleName = "cisco_ios_xe"
    header.EntityData.ParentYangName = "ospfv3-lsa"
    header.EntityData.SegmentPath = "header"
    header.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/intf-link-scope-lsas/area-scope-lsa/ospfv3-lsa/" + header.EntityData.SegmentPath
    header.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    header.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    header.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    header.EntityData.Children = types.NewOrderedMap()
    header.EntityData.Children.Append("lsa-header", types.YChild{"LsaHeader", &header.LsaHeader})
    header.EntityData.Leafs = types.NewOrderedMap()
    header.EntityData.Leafs.Append("lsa-id", types.YLeaf{"LsaId", header.LsaId})
    header.EntityData.Leafs.Append("lsa-hdr-options", types.YLeaf{"LsaHdrOptions", header.LsaHdrOptions})

    header.EntityData.YListKeys = []string {}

    return &(header.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_Header_LsaHeader
// LSA header
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_Header_LsaHeader struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA age. The type is interface{} with range: 0..65535.
    Age interface{}

    // LSA type. The type is interface{} with range: 0..65535.
    Type interface{}

    // LSA advertising router. The type is interface{} with range: 0..4294967295.
    AdvRouter interface{}

    // LSA sequence number. The type is string.
    SeqNum interface{}

    // LSA checksum. The type is string.
    Checksum interface{}

    // LSA length. The type is interface{} with range: 0..65535.
    Length interface{}
}

func (lsaHeader *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_Header_LsaHeader) GetEntityData() *types.CommonEntityData {
    lsaHeader.EntityData.YFilter = lsaHeader.YFilter
    lsaHeader.EntityData.YangName = "lsa-header"
    lsaHeader.EntityData.BundleName = "cisco_ios_xe"
    lsaHeader.EntityData.ParentYangName = "header"
    lsaHeader.EntityData.SegmentPath = "lsa-header"
    lsaHeader.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/intf-link-scope-lsas/area-scope-lsa/ospfv3-lsa/header/" + lsaHeader.EntityData.SegmentPath
    lsaHeader.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lsaHeader.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lsaHeader.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lsaHeader.EntityData.Children = types.NewOrderedMap()
    lsaHeader.EntityData.Leafs = types.NewOrderedMap()
    lsaHeader.EntityData.Leafs.Append("age", types.YLeaf{"Age", lsaHeader.Age})
    lsaHeader.EntityData.Leafs.Append("type", types.YLeaf{"Type", lsaHeader.Type})
    lsaHeader.EntityData.Leafs.Append("adv-router", types.YLeaf{"AdvRouter", lsaHeader.AdvRouter})
    lsaHeader.EntityData.Leafs.Append("seq-num", types.YLeaf{"SeqNum", lsaHeader.SeqNum})
    lsaHeader.EntityData.Leafs.Append("checksum", types.YLeaf{"Checksum", lsaHeader.Checksum})
    lsaHeader.EntityData.Leafs.Append("length", types.YLeaf{"Length", lsaHeader.Length})

    lsaHeader.EntityData.YListKeys = []string {}

    return &(lsaHeader.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody
// Decoded OSPFv3 LSA body
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA options. The type is map[string]bool.
    LsaFlagOptions interface{}

    // LSA Body Flags. The type is map[string]bool.
    LsaBodyFlags interface{}

    // OSPFv3 network.
    Network OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Network

    // OSPFv3 inter area prefix.
    Prefix OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Prefix

    // OSPFv3 inter area router.
    IaRouter OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_IaRouter

    // OSPFv3 LSA external.
    LsaExternal OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_LsaExternal

    // OSPFv3 NSSA.
    Nssa OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Nssa

    // OSPFv3 Link data.
    LinkData OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_LinkData

    // OSPFv3 Intra area prefixes.
    IaPrefix OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_IaPrefix
}

func (lsaBody *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody) GetEntityData() *types.CommonEntityData {
    lsaBody.EntityData.YFilter = lsaBody.YFilter
    lsaBody.EntityData.YangName = "lsa-body"
    lsaBody.EntityData.BundleName = "cisco_ios_xe"
    lsaBody.EntityData.ParentYangName = "ospfv3-lsa"
    lsaBody.EntityData.SegmentPath = "lsa-body"
    lsaBody.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/intf-link-scope-lsas/area-scope-lsa/ospfv3-lsa/" + lsaBody.EntityData.SegmentPath
    lsaBody.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lsaBody.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lsaBody.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lsaBody.EntityData.Children = types.NewOrderedMap()
    lsaBody.EntityData.Children.Append("network", types.YChild{"Network", &lsaBody.Network})
    lsaBody.EntityData.Children.Append("prefix", types.YChild{"Prefix", &lsaBody.Prefix})
    lsaBody.EntityData.Children.Append("ia-router", types.YChild{"IaRouter", &lsaBody.IaRouter})
    lsaBody.EntityData.Children.Append("lsa-external", types.YChild{"LsaExternal", &lsaBody.LsaExternal})
    lsaBody.EntityData.Children.Append("nssa", types.YChild{"Nssa", &lsaBody.Nssa})
    lsaBody.EntityData.Children.Append("link-data", types.YChild{"LinkData", &lsaBody.LinkData})
    lsaBody.EntityData.Children.Append("ia-prefix", types.YChild{"IaPrefix", &lsaBody.IaPrefix})
    lsaBody.EntityData.Leafs = types.NewOrderedMap()
    lsaBody.EntityData.Leafs.Append("lsa-flag-options", types.YLeaf{"LsaFlagOptions", lsaBody.LsaFlagOptions})
    lsaBody.EntityData.Leafs.Append("lsa-body-flags", types.YLeaf{"LsaBodyFlags", lsaBody.LsaBodyFlags})

    lsaBody.EntityData.YListKeys = []string {}

    return &(lsaBody.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Network
// OSPFv3 network
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Network struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of the routers attached to the network. The type is slice of
    // interface{} with range: 0..4294967295.
    AttachedRouter []interface{}

    // Network LSA options. The type is map[string]bool.
    LsaNetOptions interface{}
}

func (network *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Network) GetEntityData() *types.CommonEntityData {
    network.EntityData.YFilter = network.YFilter
    network.EntityData.YangName = "network"
    network.EntityData.BundleName = "cisco_ios_xe"
    network.EntityData.ParentYangName = "lsa-body"
    network.EntityData.SegmentPath = "network"
    network.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/intf-link-scope-lsas/area-scope-lsa/ospfv3-lsa/lsa-body/" + network.EntityData.SegmentPath
    network.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    network.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    network.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    network.EntityData.Children = types.NewOrderedMap()
    network.EntityData.Leafs = types.NewOrderedMap()
    network.EntityData.Leafs.Append("attached-router", types.YLeaf{"AttachedRouter", network.AttachedRouter})
    network.EntityData.Leafs.Append("lsa-net-options", types.YLeaf{"LsaNetOptions", network.LsaNetOptions})

    network.EntityData.YListKeys = []string {}

    return &(network.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Prefix
// OSPFv3 inter area prefix
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Prefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Metric. The type is interface{} with range: 0..4294967295.
    Metric interface{}

    // Inter area Prefix. The type is string.
    IaPrefix interface{}

    // Inter area prefix options. The type is string.
    IaPrefixOptions interface{}
}

func (prefix *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Prefix) GetEntityData() *types.CommonEntityData {
    prefix.EntityData.YFilter = prefix.YFilter
    prefix.EntityData.YangName = "prefix"
    prefix.EntityData.BundleName = "cisco_ios_xe"
    prefix.EntityData.ParentYangName = "lsa-body"
    prefix.EntityData.SegmentPath = "prefix"
    prefix.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/intf-link-scope-lsas/area-scope-lsa/ospfv3-lsa/lsa-body/" + prefix.EntityData.SegmentPath
    prefix.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    prefix.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    prefix.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    prefix.EntityData.Children = types.NewOrderedMap()
    prefix.EntityData.Leafs = types.NewOrderedMap()
    prefix.EntityData.Leafs.Append("metric", types.YLeaf{"Metric", prefix.Metric})
    prefix.EntityData.Leafs.Append("ia-prefix", types.YLeaf{"IaPrefix", prefix.IaPrefix})
    prefix.EntityData.Leafs.Append("ia-prefix-options", types.YLeaf{"IaPrefixOptions", prefix.IaPrefixOptions})

    prefix.EntityData.YListKeys = []string {}

    return &(prefix.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_IaRouter
// OSPFv3 inter area router
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_IaRouter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Metric. The type is interface{} with range: 0..4294967295.
    Metric interface{}

    // Router ID of the router being described by the LSA. The type is interface{}
    // with range: 0..4294967295.
    DestinationRouterId interface{}

    // Inter area LSA options. The type is map[string]bool.
    LsaIaOptions interface{}
}

func (iaRouter *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_IaRouter) GetEntityData() *types.CommonEntityData {
    iaRouter.EntityData.YFilter = iaRouter.YFilter
    iaRouter.EntityData.YangName = "ia-router"
    iaRouter.EntityData.BundleName = "cisco_ios_xe"
    iaRouter.EntityData.ParentYangName = "lsa-body"
    iaRouter.EntityData.SegmentPath = "ia-router"
    iaRouter.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/intf-link-scope-lsas/area-scope-lsa/ospfv3-lsa/lsa-body/" + iaRouter.EntityData.SegmentPath
    iaRouter.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    iaRouter.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    iaRouter.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    iaRouter.EntityData.Children = types.NewOrderedMap()
    iaRouter.EntityData.Leafs = types.NewOrderedMap()
    iaRouter.EntityData.Leafs.Append("metric", types.YLeaf{"Metric", iaRouter.Metric})
    iaRouter.EntityData.Leafs.Append("destination-router-id", types.YLeaf{"DestinationRouterId", iaRouter.DestinationRouterId})
    iaRouter.EntityData.Leafs.Append("lsa-ia-options", types.YLeaf{"LsaIaOptions", iaRouter.LsaIaOptions})

    iaRouter.EntityData.YListKeys = []string {}

    return &(iaRouter.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_LsaExternal
// OSPFv3 LSA external
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_LsaExternal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Metric. The type is interface{} with range: 0..4294967295.
    Metric interface{}

    // Referenced Link State type. The type is interface{} with range: 0..65535.
    ReferencedLsType interface{}

    // Prefix. The type is string.
    ExternalPrefix interface{}

    // Prefix options. The type is string.
    ExternalPrefixOptions interface{}

    // Forwarding address. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ForwardingAddress interface{}

    // Route tag. The type is interface{} with range: 0..4294967295.
    ExternalRouteTag interface{}

    // Referenced Link State ID. The type is interface{} with range:
    // 0..4294967295.
    ReferencedLinkStateId interface{}

    // LSA Flags.
    Flags OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_LsaExternal_Flags
}

func (lsaExternal *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_LsaExternal) GetEntityData() *types.CommonEntityData {
    lsaExternal.EntityData.YFilter = lsaExternal.YFilter
    lsaExternal.EntityData.YangName = "lsa-external"
    lsaExternal.EntityData.BundleName = "cisco_ios_xe"
    lsaExternal.EntityData.ParentYangName = "lsa-body"
    lsaExternal.EntityData.SegmentPath = "lsa-external"
    lsaExternal.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/intf-link-scope-lsas/area-scope-lsa/ospfv3-lsa/lsa-body/" + lsaExternal.EntityData.SegmentPath
    lsaExternal.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lsaExternal.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lsaExternal.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lsaExternal.EntityData.Children = types.NewOrderedMap()
    lsaExternal.EntityData.Children.Append("flags", types.YChild{"Flags", &lsaExternal.Flags})
    lsaExternal.EntityData.Leafs = types.NewOrderedMap()
    lsaExternal.EntityData.Leafs.Append("metric", types.YLeaf{"Metric", lsaExternal.Metric})
    lsaExternal.EntityData.Leafs.Append("referenced-ls-type", types.YLeaf{"ReferencedLsType", lsaExternal.ReferencedLsType})
    lsaExternal.EntityData.Leafs.Append("external-prefix", types.YLeaf{"ExternalPrefix", lsaExternal.ExternalPrefix})
    lsaExternal.EntityData.Leafs.Append("external-prefix-options", types.YLeaf{"ExternalPrefixOptions", lsaExternal.ExternalPrefixOptions})
    lsaExternal.EntityData.Leafs.Append("forwarding-address", types.YLeaf{"ForwardingAddress", lsaExternal.ForwardingAddress})
    lsaExternal.EntityData.Leafs.Append("external-route-tag", types.YLeaf{"ExternalRouteTag", lsaExternal.ExternalRouteTag})
    lsaExternal.EntityData.Leafs.Append("referenced-link-state-id", types.YLeaf{"ReferencedLinkStateId", lsaExternal.ReferencedLinkStateId})

    lsaExternal.EntityData.YListKeys = []string {}

    return &(lsaExternal.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_LsaExternal_Flags
// LSA Flags
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_LsaExternal_Flags struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // When set, the metric specified is a Type 2 external metric. The type is
    // bool.
    EFlag interface{}
}

func (flags *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_LsaExternal_Flags) GetEntityData() *types.CommonEntityData {
    flags.EntityData.YFilter = flags.YFilter
    flags.EntityData.YangName = "flags"
    flags.EntityData.BundleName = "cisco_ios_xe"
    flags.EntityData.ParentYangName = "lsa-external"
    flags.EntityData.SegmentPath = "flags"
    flags.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/intf-link-scope-lsas/area-scope-lsa/ospfv3-lsa/lsa-body/lsa-external/" + flags.EntityData.SegmentPath
    flags.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    flags.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    flags.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    flags.EntityData.Children = types.NewOrderedMap()
    flags.EntityData.Leafs = types.NewOrderedMap()
    flags.EntityData.Leafs.Append("e-flag", types.YLeaf{"EFlag", flags.EFlag})

    flags.EntityData.YListKeys = []string {}

    return &(flags.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Nssa
// OSPFv3 NSSA
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Nssa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // NSSA LSA.
    LsaNssaExternal OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Nssa_LsaNssaExternal
}

func (nssa *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Nssa) GetEntityData() *types.CommonEntityData {
    nssa.EntityData.YFilter = nssa.YFilter
    nssa.EntityData.YangName = "nssa"
    nssa.EntityData.BundleName = "cisco_ios_xe"
    nssa.EntityData.ParentYangName = "lsa-body"
    nssa.EntityData.SegmentPath = "nssa"
    nssa.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/intf-link-scope-lsas/area-scope-lsa/ospfv3-lsa/lsa-body/" + nssa.EntityData.SegmentPath
    nssa.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nssa.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nssa.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nssa.EntityData.Children = types.NewOrderedMap()
    nssa.EntityData.Children.Append("lsa-nssa-external", types.YChild{"LsaNssaExternal", &nssa.LsaNssaExternal})
    nssa.EntityData.Leafs = types.NewOrderedMap()

    nssa.EntityData.YListKeys = []string {}

    return &(nssa.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Nssa_LsaNssaExternal
// NSSA LSA
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Nssa_LsaNssaExternal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Metric. The type is interface{} with range: 0..4294967295.
    Metric interface{}

    // Referenced Link State type. The type is interface{} with range: 0..65535.
    ReferencedLsType interface{}

    // Prefix. The type is string.
    ExternalPrefix interface{}

    // Prefix options. The type is string.
    ExternalPrefixOptions interface{}

    // Forwarding address. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ForwardingAddress interface{}

    // Route tag. The type is interface{} with range: 0..4294967295.
    ExternalRouteTag interface{}

    // Referenced Link State ID. The type is interface{} with range:
    // 0..4294967295.
    ReferencedLinkStateId interface{}

    // LSA Flags.
    Flags OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Nssa_LsaNssaExternal_Flags
}

func (lsaNssaExternal *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Nssa_LsaNssaExternal) GetEntityData() *types.CommonEntityData {
    lsaNssaExternal.EntityData.YFilter = lsaNssaExternal.YFilter
    lsaNssaExternal.EntityData.YangName = "lsa-nssa-external"
    lsaNssaExternal.EntityData.BundleName = "cisco_ios_xe"
    lsaNssaExternal.EntityData.ParentYangName = "nssa"
    lsaNssaExternal.EntityData.SegmentPath = "lsa-nssa-external"
    lsaNssaExternal.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/intf-link-scope-lsas/area-scope-lsa/ospfv3-lsa/lsa-body/nssa/" + lsaNssaExternal.EntityData.SegmentPath
    lsaNssaExternal.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lsaNssaExternal.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lsaNssaExternal.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lsaNssaExternal.EntityData.Children = types.NewOrderedMap()
    lsaNssaExternal.EntityData.Children.Append("flags", types.YChild{"Flags", &lsaNssaExternal.Flags})
    lsaNssaExternal.EntityData.Leafs = types.NewOrderedMap()
    lsaNssaExternal.EntityData.Leafs.Append("metric", types.YLeaf{"Metric", lsaNssaExternal.Metric})
    lsaNssaExternal.EntityData.Leafs.Append("referenced-ls-type", types.YLeaf{"ReferencedLsType", lsaNssaExternal.ReferencedLsType})
    lsaNssaExternal.EntityData.Leafs.Append("external-prefix", types.YLeaf{"ExternalPrefix", lsaNssaExternal.ExternalPrefix})
    lsaNssaExternal.EntityData.Leafs.Append("external-prefix-options", types.YLeaf{"ExternalPrefixOptions", lsaNssaExternal.ExternalPrefixOptions})
    lsaNssaExternal.EntityData.Leafs.Append("forwarding-address", types.YLeaf{"ForwardingAddress", lsaNssaExternal.ForwardingAddress})
    lsaNssaExternal.EntityData.Leafs.Append("external-route-tag", types.YLeaf{"ExternalRouteTag", lsaNssaExternal.ExternalRouteTag})
    lsaNssaExternal.EntityData.Leafs.Append("referenced-link-state-id", types.YLeaf{"ReferencedLinkStateId", lsaNssaExternal.ReferencedLinkStateId})

    lsaNssaExternal.EntityData.YListKeys = []string {}

    return &(lsaNssaExternal.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Nssa_LsaNssaExternal_Flags
// LSA Flags
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Nssa_LsaNssaExternal_Flags struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // When set, the metric specified is a Type 2 external metric. The type is
    // bool.
    EFlag interface{}
}

func (flags *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Nssa_LsaNssaExternal_Flags) GetEntityData() *types.CommonEntityData {
    flags.EntityData.YFilter = flags.YFilter
    flags.EntityData.YangName = "flags"
    flags.EntityData.BundleName = "cisco_ios_xe"
    flags.EntityData.ParentYangName = "lsa-nssa-external"
    flags.EntityData.SegmentPath = "flags"
    flags.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/intf-link-scope-lsas/area-scope-lsa/ospfv3-lsa/lsa-body/nssa/lsa-nssa-external/" + flags.EntityData.SegmentPath
    flags.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    flags.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    flags.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    flags.EntityData.Children = types.NewOrderedMap()
    flags.EntityData.Leafs = types.NewOrderedMap()
    flags.EntityData.Leafs.Append("e-flag", types.YLeaf{"EFlag", flags.EFlag})

    flags.EntityData.YListKeys = []string {}

    return &(flags.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_LinkData
// OSPFv3 Link data
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_LinkData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Router priority of the interce. The type is interface{} with range: 0..255.
    RtrPriority interface{}

    // The originating router's link-local interface address on the link. The type
    // is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    LinkLocalInterfaceAddress interface{}

    // Number of prefixes. The type is interface{} with range: 0..4294967295.
    NumOfPrefixes interface{}

    // Link data LSA options. The type is map[string]bool.
    LsaIdOptions interface{}
}

func (linkData *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_LinkData) GetEntityData() *types.CommonEntityData {
    linkData.EntityData.YFilter = linkData.YFilter
    linkData.EntityData.YangName = "link-data"
    linkData.EntityData.BundleName = "cisco_ios_xe"
    linkData.EntityData.ParentYangName = "lsa-body"
    linkData.EntityData.SegmentPath = "link-data"
    linkData.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/intf-link-scope-lsas/area-scope-lsa/ospfv3-lsa/lsa-body/" + linkData.EntityData.SegmentPath
    linkData.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    linkData.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    linkData.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    linkData.EntityData.Children = types.NewOrderedMap()
    linkData.EntityData.Leafs = types.NewOrderedMap()
    linkData.EntityData.Leafs.Append("rtr-priority", types.YLeaf{"RtrPriority", linkData.RtrPriority})
    linkData.EntityData.Leafs.Append("link-local-interface-address", types.YLeaf{"LinkLocalInterfaceAddress", linkData.LinkLocalInterfaceAddress})
    linkData.EntityData.Leafs.Append("num-of-prefixes", types.YLeaf{"NumOfPrefixes", linkData.NumOfPrefixes})
    linkData.EntityData.Leafs.Append("lsa-id-options", types.YLeaf{"LsaIdOptions", linkData.LsaIdOptions})

    linkData.EntityData.YListKeys = []string {}

    return &(linkData.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_IaPrefix
// OSPFv3 Intra area prefixes
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_IaPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Referenced Link State type. The type is interface{} with range: 0..65535.
    ReferencedLsType interface{}

    // Referenced Link State ID. The type is interface{} with range:
    // 0..4294967295.
    ReferencedLinkStateId interface{}

    // Referenced Advertising Router. The type is one of the following types:
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ReferencedAdvRouter interface{}

    // Number of prefixes. The type is interface{} with range: 0..65535.
    NumOfPrefixes interface{}
}

func (iaPrefix *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_IaPrefix) GetEntityData() *types.CommonEntityData {
    iaPrefix.EntityData.YFilter = iaPrefix.YFilter
    iaPrefix.EntityData.YangName = "ia-prefix"
    iaPrefix.EntityData.BundleName = "cisco_ios_xe"
    iaPrefix.EntityData.ParentYangName = "lsa-body"
    iaPrefix.EntityData.SegmentPath = "ia-prefix"
    iaPrefix.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/intf-link-scope-lsas/area-scope-lsa/ospfv3-lsa/lsa-body/" + iaPrefix.EntityData.SegmentPath
    iaPrefix.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    iaPrefix.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    iaPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    iaPrefix.EntityData.Children = types.NewOrderedMap()
    iaPrefix.EntityData.Leafs = types.NewOrderedMap()
    iaPrefix.EntityData.Leafs.Append("referenced-ls-type", types.YLeaf{"ReferencedLsType", iaPrefix.ReferencedLsType})
    iaPrefix.EntityData.Leafs.Append("referenced-link-state-id", types.YLeaf{"ReferencedLinkStateId", iaPrefix.ReferencedLinkStateId})
    iaPrefix.EntityData.Leafs.Append("referenced-adv-router", types.YLeaf{"ReferencedAdvRouter", iaPrefix.ReferencedAdvRouter})
    iaPrefix.EntityData.Leafs.Append("num-of-prefixes", types.YLeaf{"NumOfPrefixes", iaPrefix.NumOfPrefixes})

    iaPrefix.EntityData.YListKeys = []string {}

    return &(iaPrefix.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Link
// OSPFv3 links
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Link struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface ID. The type is interface{} with range:
    // 0..4294967295.
    InterfaceId interface{}

    // This attribute is a key. Neighbor interface ID. The type is interface{}
    // with range: 0..4294967295.
    NeighborInterfaceId interface{}

    // This attribute is a key. Neighbor router ID. The type is interface{} with
    // range: 0..4294967295.
    NeighborRouterId interface{}

    // Link type. The type is interface{} with range: 0..255.
    Type interface{}

    // Metric. The type is interface{} with range: 0..65535.
    Metric interface{}
}

func (ospfv3Link *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Link) GetEntityData() *types.CommonEntityData {
    ospfv3Link.EntityData.YFilter = ospfv3Link.YFilter
    ospfv3Link.EntityData.YangName = "ospfv3-link"
    ospfv3Link.EntityData.BundleName = "cisco_ios_xe"
    ospfv3Link.EntityData.ParentYangName = "area-scope-lsa"
    ospfv3Link.EntityData.SegmentPath = "ospfv3-link" + types.AddKeyToken(ospfv3Link.InterfaceId, "interface-id") + types.AddKeyToken(ospfv3Link.NeighborInterfaceId, "neighbor-interface-id") + types.AddKeyToken(ospfv3Link.NeighborRouterId, "neighbor-router-id")
    ospfv3Link.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/intf-link-scope-lsas/area-scope-lsa/" + ospfv3Link.EntityData.SegmentPath
    ospfv3Link.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv3Link.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv3Link.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv3Link.EntityData.Children = types.NewOrderedMap()
    ospfv3Link.EntityData.Leafs = types.NewOrderedMap()
    ospfv3Link.EntityData.Leafs.Append("interface-id", types.YLeaf{"InterfaceId", ospfv3Link.InterfaceId})
    ospfv3Link.EntityData.Leafs.Append("neighbor-interface-id", types.YLeaf{"NeighborInterfaceId", ospfv3Link.NeighborInterfaceId})
    ospfv3Link.EntityData.Leafs.Append("neighbor-router-id", types.YLeaf{"NeighborRouterId", ospfv3Link.NeighborRouterId})
    ospfv3Link.EntityData.Leafs.Append("type", types.YLeaf{"Type", ospfv3Link.Type})
    ospfv3Link.EntityData.Leafs.Append("metric", types.YLeaf{"Metric", ospfv3Link.Metric})

    ospfv3Link.EntityData.YListKeys = []string {"InterfaceId", "NeighborInterfaceId", "NeighborRouterId"}

    return &(ospfv3Link.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Prefix
// OSPFv3 prefix-list
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Prefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Prefix. The type is string.
    Prefix interface{}

    // Prefix options. The type is string.
    PrefixOptions interface{}
}

func (ospfv3Prefix *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Prefix) GetEntityData() *types.CommonEntityData {
    ospfv3Prefix.EntityData.YFilter = ospfv3Prefix.YFilter
    ospfv3Prefix.EntityData.YangName = "ospfv3-prefix"
    ospfv3Prefix.EntityData.BundleName = "cisco_ios_xe"
    ospfv3Prefix.EntityData.ParentYangName = "area-scope-lsa"
    ospfv3Prefix.EntityData.SegmentPath = "ospfv3-prefix" + types.AddKeyToken(ospfv3Prefix.Prefix, "prefix")
    ospfv3Prefix.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/intf-link-scope-lsas/area-scope-lsa/" + ospfv3Prefix.EntityData.SegmentPath
    ospfv3Prefix.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv3Prefix.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv3Prefix.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv3Prefix.EntityData.Children = types.NewOrderedMap()
    ospfv3Prefix.EntityData.Leafs = types.NewOrderedMap()
    ospfv3Prefix.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", ospfv3Prefix.Prefix})
    ospfv3Prefix.EntityData.Leafs.Append("prefix-options", types.YLeaf{"PrefixOptions", ospfv3Prefix.PrefixOptions})

    ospfv3Prefix.EntityData.YListKeys = []string {"Prefix"}

    return &(ospfv3Prefix.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3IaPrefix
// OSPFv3 intra-area prefix-list
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3IaPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Prefix. The type is string.
    Prefix interface{}

    // Prefix options. The type is string.
    PrefixOptions interface{}
}

func (ospfv3IaPrefix *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3IaPrefix) GetEntityData() *types.CommonEntityData {
    ospfv3IaPrefix.EntityData.YFilter = ospfv3IaPrefix.YFilter
    ospfv3IaPrefix.EntityData.YangName = "ospfv3-ia-prefix"
    ospfv3IaPrefix.EntityData.BundleName = "cisco_ios_xe"
    ospfv3IaPrefix.EntityData.ParentYangName = "area-scope-lsa"
    ospfv3IaPrefix.EntityData.SegmentPath = "ospfv3-ia-prefix" + types.AddKeyToken(ospfv3IaPrefix.Prefix, "prefix")
    ospfv3IaPrefix.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/intf-link-scope-lsas/area-scope-lsa/" + ospfv3IaPrefix.EntityData.SegmentPath
    ospfv3IaPrefix.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv3IaPrefix.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv3IaPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv3IaPrefix.EntityData.Children = types.NewOrderedMap()
    ospfv3IaPrefix.EntityData.Leafs = types.NewOrderedMap()
    ospfv3IaPrefix.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", ospfv3IaPrefix.Prefix})
    ospfv3IaPrefix.EntityData.Leafs.Append("prefix-options", types.YLeaf{"PrefixOptions", ospfv3IaPrefix.PrefixOptions})

    ospfv3IaPrefix.EntityData.YListKeys = []string {"Prefix"}

    return &(ospfv3IaPrefix.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfMultiTopology
// OSPF interface topology
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfMultiTopology struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. One of the topology enabled on this interface. The
    // type is string.
    Name interface{}
}

func (intfMultiTopology *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfMultiTopology) GetEntityData() *types.CommonEntityData {
    intfMultiTopology.EntityData.YFilter = intfMultiTopology.YFilter
    intfMultiTopology.EntityData.YangName = "intf-multi-topology"
    intfMultiTopology.EntityData.BundleName = "cisco_ios_xe"
    intfMultiTopology.EntityData.ParentYangName = "ospf-interface"
    intfMultiTopology.EntityData.SegmentPath = "intf-multi-topology" + types.AddKeyToken(intfMultiTopology.Name, "name")
    intfMultiTopology.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/ospf-interface/" + intfMultiTopology.EntityData.SegmentPath
    intfMultiTopology.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    intfMultiTopology.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    intfMultiTopology.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    intfMultiTopology.EntityData.Children = types.NewOrderedMap()
    intfMultiTopology.EntityData.Leafs = types.NewOrderedMap()
    intfMultiTopology.EntityData.Leafs.Append("name", types.YLeaf{"Name", intfMultiTopology.Name})

    intfMultiTopology.EntityData.YListKeys = []string {"Name"}

    return &(intfMultiTopology.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa
// List of OSPF area scope LSA
type OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. OSPF link scope LSA type. The type is interface{}
    // with range: 0..4294967295.
    LsaType interface{}

    // List of OSPF link scope LSAs. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa.
    AreaScopeLsa []*OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa
}

func (areaScopeLsa *OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa) GetEntityData() *types.CommonEntityData {
    areaScopeLsa.EntityData.YFilter = areaScopeLsa.YFilter
    areaScopeLsa.EntityData.YangName = "area-scope-lsa"
    areaScopeLsa.EntityData.BundleName = "cisco_ios_xe"
    areaScopeLsa.EntityData.ParentYangName = "ospf-area"
    areaScopeLsa.EntityData.SegmentPath = "area-scope-lsa" + types.AddKeyToken(areaScopeLsa.LsaType, "lsa-type")
    areaScopeLsa.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/" + areaScopeLsa.EntityData.SegmentPath
    areaScopeLsa.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    areaScopeLsa.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    areaScopeLsa.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    areaScopeLsa.EntityData.Children = types.NewOrderedMap()
    areaScopeLsa.EntityData.Children.Append("area-scope-lsa", types.YChild{"AreaScopeLsa", nil})
    for i := range areaScopeLsa.AreaScopeLsa {
        areaScopeLsa.EntityData.Children.Append(types.GetSegmentPath(areaScopeLsa.AreaScopeLsa[i]), types.YChild{"AreaScopeLsa", areaScopeLsa.AreaScopeLsa[i]})
    }
    areaScopeLsa.EntityData.Leafs = types.NewOrderedMap()
    areaScopeLsa.EntityData.Leafs.Append("lsa-type", types.YLeaf{"LsaType", areaScopeLsa.LsaType})

    areaScopeLsa.EntityData.YListKeys = []string {"LsaType"}

    return &(areaScopeLsa.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa
// List of OSPF link scope LSAs
type OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. LSA Type. The type is interface{} with range:
    // 0..4294967295.
    LsaType interface{}

    // This attribute is a key. Advertising router. The type is one of the
    // following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    AdvRouter interface{}

    // The OSPF LSA body is fully decoded. The type is bool.
    DecodedCompleted interface{}

    // The complete LSA in network byte order as received/sent over the wire. The
    // type is slice of interface{} with range: 0..255.
    RawData []interface{}

    // OSPFv2 LSA.
    Ospfv2Lsa OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv2Lsa

    // Router LSA link. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv2Link.
    Ospfv2Link []*OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv2Link

    // Summary LSA. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv2Topology.
    Ospfv2Topology []*OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv2Topology

    // External LSA. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv2External.
    Ospfv2External []*OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv2External

    // OSPFv3 LSA.
    Ospfv3Lsa OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Lsa

    // OSPFv3 links. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Link.
    Ospfv3Link []*OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Link

    // OSPFv3 prefix-list. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Prefix.
    Ospfv3Prefix []*OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Prefix

    // OSPFv3 intra-area prefix-list. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3IaPrefix.
    Ospfv3IaPrefix []*OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3IaPrefix
}

func (areaScopeLsa *OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa) GetEntityData() *types.CommonEntityData {
    areaScopeLsa.EntityData.YFilter = areaScopeLsa.YFilter
    areaScopeLsa.EntityData.YangName = "area-scope-lsa"
    areaScopeLsa.EntityData.BundleName = "cisco_ios_xe"
    areaScopeLsa.EntityData.ParentYangName = "area-scope-lsa"
    areaScopeLsa.EntityData.SegmentPath = "area-scope-lsa" + types.AddKeyToken(areaScopeLsa.LsaType, "lsa-type") + types.AddKeyToken(areaScopeLsa.AdvRouter, "adv-router")
    areaScopeLsa.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/area-scope-lsa/" + areaScopeLsa.EntityData.SegmentPath
    areaScopeLsa.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    areaScopeLsa.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    areaScopeLsa.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    areaScopeLsa.EntityData.Children = types.NewOrderedMap()
    areaScopeLsa.EntityData.Children.Append("ospfv2-lsa", types.YChild{"Ospfv2Lsa", &areaScopeLsa.Ospfv2Lsa})
    areaScopeLsa.EntityData.Children.Append("ospfv2-link", types.YChild{"Ospfv2Link", nil})
    for i := range areaScopeLsa.Ospfv2Link {
        areaScopeLsa.EntityData.Children.Append(types.GetSegmentPath(areaScopeLsa.Ospfv2Link[i]), types.YChild{"Ospfv2Link", areaScopeLsa.Ospfv2Link[i]})
    }
    areaScopeLsa.EntityData.Children.Append("ospfv2-topology", types.YChild{"Ospfv2Topology", nil})
    for i := range areaScopeLsa.Ospfv2Topology {
        areaScopeLsa.EntityData.Children.Append(types.GetSegmentPath(areaScopeLsa.Ospfv2Topology[i]), types.YChild{"Ospfv2Topology", areaScopeLsa.Ospfv2Topology[i]})
    }
    areaScopeLsa.EntityData.Children.Append("ospfv2-external", types.YChild{"Ospfv2External", nil})
    for i := range areaScopeLsa.Ospfv2External {
        areaScopeLsa.EntityData.Children.Append(types.GetSegmentPath(areaScopeLsa.Ospfv2External[i]), types.YChild{"Ospfv2External", areaScopeLsa.Ospfv2External[i]})
    }
    areaScopeLsa.EntityData.Children.Append("ospfv3-lsa", types.YChild{"Ospfv3Lsa", &areaScopeLsa.Ospfv3Lsa})
    areaScopeLsa.EntityData.Children.Append("ospfv3-link", types.YChild{"Ospfv3Link", nil})
    for i := range areaScopeLsa.Ospfv3Link {
        areaScopeLsa.EntityData.Children.Append(types.GetSegmentPath(areaScopeLsa.Ospfv3Link[i]), types.YChild{"Ospfv3Link", areaScopeLsa.Ospfv3Link[i]})
    }
    areaScopeLsa.EntityData.Children.Append("ospfv3-prefix", types.YChild{"Ospfv3Prefix", nil})
    for i := range areaScopeLsa.Ospfv3Prefix {
        areaScopeLsa.EntityData.Children.Append(types.GetSegmentPath(areaScopeLsa.Ospfv3Prefix[i]), types.YChild{"Ospfv3Prefix", areaScopeLsa.Ospfv3Prefix[i]})
    }
    areaScopeLsa.EntityData.Children.Append("ospfv3-ia-prefix", types.YChild{"Ospfv3IaPrefix", nil})
    for i := range areaScopeLsa.Ospfv3IaPrefix {
        areaScopeLsa.EntityData.Children.Append(types.GetSegmentPath(areaScopeLsa.Ospfv3IaPrefix[i]), types.YChild{"Ospfv3IaPrefix", areaScopeLsa.Ospfv3IaPrefix[i]})
    }
    areaScopeLsa.EntityData.Leafs = types.NewOrderedMap()
    areaScopeLsa.EntityData.Leafs.Append("lsa-type", types.YLeaf{"LsaType", areaScopeLsa.LsaType})
    areaScopeLsa.EntityData.Leafs.Append("adv-router", types.YLeaf{"AdvRouter", areaScopeLsa.AdvRouter})
    areaScopeLsa.EntityData.Leafs.Append("decoded-completed", types.YLeaf{"DecodedCompleted", areaScopeLsa.DecodedCompleted})
    areaScopeLsa.EntityData.Leafs.Append("raw-data", types.YLeaf{"RawData", areaScopeLsa.RawData})

    areaScopeLsa.EntityData.YListKeys = []string {"LsaType", "AdvRouter"}

    return &(areaScopeLsa.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv2Lsa
// OSPFv2 LSA
type OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv2Lsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Decoded OSPFv2 LSA header data.
    Header OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv2Lsa_Header

    // Decoded OSPFv2 LSA body data.
    LsaBody OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv2Lsa_LsaBody
}

func (ospfv2Lsa *OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv2Lsa) GetEntityData() *types.CommonEntityData {
    ospfv2Lsa.EntityData.YFilter = ospfv2Lsa.YFilter
    ospfv2Lsa.EntityData.YangName = "ospfv2-lsa"
    ospfv2Lsa.EntityData.BundleName = "cisco_ios_xe"
    ospfv2Lsa.EntityData.ParentYangName = "area-scope-lsa"
    ospfv2Lsa.EntityData.SegmentPath = "ospfv2-lsa"
    ospfv2Lsa.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/area-scope-lsa/area-scope-lsa/" + ospfv2Lsa.EntityData.SegmentPath
    ospfv2Lsa.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv2Lsa.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv2Lsa.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv2Lsa.EntityData.Children = types.NewOrderedMap()
    ospfv2Lsa.EntityData.Children.Append("header", types.YChild{"Header", &ospfv2Lsa.Header})
    ospfv2Lsa.EntityData.Children.Append("lsa-body", types.YChild{"LsaBody", &ospfv2Lsa.LsaBody})
    ospfv2Lsa.EntityData.Leafs = types.NewOrderedMap()

    ospfv2Lsa.EntityData.YListKeys = []string {}

    return &(ospfv2Lsa.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv2Lsa_Header
// Decoded OSPFv2 LSA header data
type OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv2Lsa_Header struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA ID. The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    LsaId interface{}

    // Opaque type. The type is interface{} with range: 0..255.
    OpaqueType interface{}

    // Opaque ID. The type is interface{} with range: 0..4294967295.
    OpaqueId interface{}

    // LSA age. The type is interface{} with range: 0..65535.
    Age interface{}

    // LSA type. The type is interface{} with range: 0..65535.
    Type interface{}

    // LSA advertising router. The type is interface{} with range: 0..4294967295.
    AdvRouter interface{}

    // LSA sequence number. The type is string.
    SeqNum interface{}

    // LSA checksum. The type is string.
    Checksum interface{}

    // LSA length. The type is interface{} with range: 0..65535.
    Length interface{}

    // LSA options. The type is map[string]bool.
    FlagOptions interface{}
}

func (header *OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv2Lsa_Header) GetEntityData() *types.CommonEntityData {
    header.EntityData.YFilter = header.YFilter
    header.EntityData.YangName = "header"
    header.EntityData.BundleName = "cisco_ios_xe"
    header.EntityData.ParentYangName = "ospfv2-lsa"
    header.EntityData.SegmentPath = "header"
    header.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/area-scope-lsa/area-scope-lsa/ospfv2-lsa/" + header.EntityData.SegmentPath
    header.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    header.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    header.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    header.EntityData.Children = types.NewOrderedMap()
    header.EntityData.Leafs = types.NewOrderedMap()
    header.EntityData.Leafs.Append("lsa-id", types.YLeaf{"LsaId", header.LsaId})
    header.EntityData.Leafs.Append("opaque-type", types.YLeaf{"OpaqueType", header.OpaqueType})
    header.EntityData.Leafs.Append("opaque-id", types.YLeaf{"OpaqueId", header.OpaqueId})
    header.EntityData.Leafs.Append("age", types.YLeaf{"Age", header.Age})
    header.EntityData.Leafs.Append("type", types.YLeaf{"Type", header.Type})
    header.EntityData.Leafs.Append("adv-router", types.YLeaf{"AdvRouter", header.AdvRouter})
    header.EntityData.Leafs.Append("seq-num", types.YLeaf{"SeqNum", header.SeqNum})
    header.EntityData.Leafs.Append("checksum", types.YLeaf{"Checksum", header.Checksum})
    header.EntityData.Leafs.Append("length", types.YLeaf{"Length", header.Length})
    header.EntityData.Leafs.Append("flag-options", types.YLeaf{"FlagOptions", header.FlagOptions})

    header.EntityData.YListKeys = []string {}

    return &(header.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv2Lsa_LsaBody
// Decoded OSPFv2 LSA body data
type OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv2Lsa_LsaBody struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of links. The type is interface{} with range: 0..65535.
    NumOfLinks interface{}

    // Summary mask. The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    SummaryMask interface{}

    // External mask. The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ExternalMask interface{}

    // LSA body flags. The type is map[string]bool.
    BodyFlagOptions interface{}

    // Network details.
    Network OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv2Lsa_LsaBody_Network
}

func (lsaBody *OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv2Lsa_LsaBody) GetEntityData() *types.CommonEntityData {
    lsaBody.EntityData.YFilter = lsaBody.YFilter
    lsaBody.EntityData.YangName = "lsa-body"
    lsaBody.EntityData.BundleName = "cisco_ios_xe"
    lsaBody.EntityData.ParentYangName = "ospfv2-lsa"
    lsaBody.EntityData.SegmentPath = "lsa-body"
    lsaBody.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/area-scope-lsa/area-scope-lsa/ospfv2-lsa/" + lsaBody.EntityData.SegmentPath
    lsaBody.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lsaBody.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lsaBody.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lsaBody.EntityData.Children = types.NewOrderedMap()
    lsaBody.EntityData.Children.Append("network", types.YChild{"Network", &lsaBody.Network})
    lsaBody.EntityData.Leafs = types.NewOrderedMap()
    lsaBody.EntityData.Leafs.Append("num-of-links", types.YLeaf{"NumOfLinks", lsaBody.NumOfLinks})
    lsaBody.EntityData.Leafs.Append("summary-mask", types.YLeaf{"SummaryMask", lsaBody.SummaryMask})
    lsaBody.EntityData.Leafs.Append("external-mask", types.YLeaf{"ExternalMask", lsaBody.ExternalMask})
    lsaBody.EntityData.Leafs.Append("body-flag-options", types.YLeaf{"BodyFlagOptions", lsaBody.BodyFlagOptions})

    lsaBody.EntityData.YListKeys = []string {}

    return &(lsaBody.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv2Lsa_LsaBody_Network
// Network details
type OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv2Lsa_LsaBody_Network struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP network mask. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NetworkMask interface{}

    // List of the routers attached to the network. The type is slice of
    // interface{} with range: 0..4294967295.
    AttachedRouter []interface{}
}

func (network *OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv2Lsa_LsaBody_Network) GetEntityData() *types.CommonEntityData {
    network.EntityData.YFilter = network.YFilter
    network.EntityData.YangName = "network"
    network.EntityData.BundleName = "cisco_ios_xe"
    network.EntityData.ParentYangName = "lsa-body"
    network.EntityData.SegmentPath = "network"
    network.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/area-scope-lsa/area-scope-lsa/ospfv2-lsa/lsa-body/" + network.EntityData.SegmentPath
    network.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    network.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    network.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    network.EntityData.Children = types.NewOrderedMap()
    network.EntityData.Leafs = types.NewOrderedMap()
    network.EntityData.Leafs.Append("network-mask", types.YLeaf{"NetworkMask", network.NetworkMask})
    network.EntityData.Leafs.Append("attached-router", types.YLeaf{"AttachedRouter", network.AttachedRouter})

    network.EntityData.YListKeys = []string {}

    return &(network.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv2Link
// Router LSA link
type OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv2Link struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Link ID. The type is interface{} with range:
    // 0..4294967295.
    LinkId interface{}

    // This attribute is a key. Link data. The type is interface{} with range:
    // 0..4294967295.
    LinkData interface{}

    // Link type. The type is interface{} with range: 0..255.
    Type interface{}

    // Topology specific information. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv2Link_Ospfv2Topology.
    Ospfv2Topology []*OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv2Link_Ospfv2Topology
}

func (ospfv2Link *OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv2Link) GetEntityData() *types.CommonEntityData {
    ospfv2Link.EntityData.YFilter = ospfv2Link.YFilter
    ospfv2Link.EntityData.YangName = "ospfv2-link"
    ospfv2Link.EntityData.BundleName = "cisco_ios_xe"
    ospfv2Link.EntityData.ParentYangName = "area-scope-lsa"
    ospfv2Link.EntityData.SegmentPath = "ospfv2-link" + types.AddKeyToken(ospfv2Link.LinkId, "link-id") + types.AddKeyToken(ospfv2Link.LinkData, "link-data")
    ospfv2Link.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/area-scope-lsa/area-scope-lsa/" + ospfv2Link.EntityData.SegmentPath
    ospfv2Link.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv2Link.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv2Link.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv2Link.EntityData.Children = types.NewOrderedMap()
    ospfv2Link.EntityData.Children.Append("ospfv2-topology", types.YChild{"Ospfv2Topology", nil})
    for i := range ospfv2Link.Ospfv2Topology {
        ospfv2Link.EntityData.Children.Append(types.GetSegmentPath(ospfv2Link.Ospfv2Topology[i]), types.YChild{"Ospfv2Topology", ospfv2Link.Ospfv2Topology[i]})
    }
    ospfv2Link.EntityData.Leafs = types.NewOrderedMap()
    ospfv2Link.EntityData.Leafs.Append("link-id", types.YLeaf{"LinkId", ospfv2Link.LinkId})
    ospfv2Link.EntityData.Leafs.Append("link-data", types.YLeaf{"LinkData", ospfv2Link.LinkData})
    ospfv2Link.EntityData.Leafs.Append("type", types.YLeaf{"Type", ospfv2Link.Type})

    ospfv2Link.EntityData.YListKeys = []string {"LinkId", "LinkData"}

    return &(ospfv2Link.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv2Link_Ospfv2Topology
// Topology specific information
type OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv2Link_Ospfv2Topology struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. MT-ID for topology enabled link. The type is
    // interface{} with range: 0..4294967295.
    MtId interface{}

    // Metric for the topology. The type is interface{} with range: 0..65535.
    Metric interface{}
}

func (ospfv2Topology *OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv2Link_Ospfv2Topology) GetEntityData() *types.CommonEntityData {
    ospfv2Topology.EntityData.YFilter = ospfv2Topology.YFilter
    ospfv2Topology.EntityData.YangName = "ospfv2-topology"
    ospfv2Topology.EntityData.BundleName = "cisco_ios_xe"
    ospfv2Topology.EntityData.ParentYangName = "ospfv2-link"
    ospfv2Topology.EntityData.SegmentPath = "ospfv2-topology" + types.AddKeyToken(ospfv2Topology.MtId, "mt-id")
    ospfv2Topology.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/area-scope-lsa/area-scope-lsa/ospfv2-link/" + ospfv2Topology.EntityData.SegmentPath
    ospfv2Topology.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv2Topology.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv2Topology.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv2Topology.EntityData.Children = types.NewOrderedMap()
    ospfv2Topology.EntityData.Leafs = types.NewOrderedMap()
    ospfv2Topology.EntityData.Leafs.Append("mt-id", types.YLeaf{"MtId", ospfv2Topology.MtId})
    ospfv2Topology.EntityData.Leafs.Append("metric", types.YLeaf{"Metric", ospfv2Topology.Metric})

    ospfv2Topology.EntityData.YListKeys = []string {"MtId"}

    return &(ospfv2Topology.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv2Topology
// Summary LSA
type OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv2Topology struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. MT-ID for topology enabled link. The type is
    // interface{} with range: 0..4294967295.
    MtId interface{}

    // Metric for the topology. The type is interface{} with range: 0..65535.
    Metric interface{}
}

func (ospfv2Topology *OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv2Topology) GetEntityData() *types.CommonEntityData {
    ospfv2Topology.EntityData.YFilter = ospfv2Topology.YFilter
    ospfv2Topology.EntityData.YangName = "ospfv2-topology"
    ospfv2Topology.EntityData.BundleName = "cisco_ios_xe"
    ospfv2Topology.EntityData.ParentYangName = "area-scope-lsa"
    ospfv2Topology.EntityData.SegmentPath = "ospfv2-topology" + types.AddKeyToken(ospfv2Topology.MtId, "mt-id")
    ospfv2Topology.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/area-scope-lsa/area-scope-lsa/" + ospfv2Topology.EntityData.SegmentPath
    ospfv2Topology.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv2Topology.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv2Topology.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv2Topology.EntityData.Children = types.NewOrderedMap()
    ospfv2Topology.EntityData.Leafs = types.NewOrderedMap()
    ospfv2Topology.EntityData.Leafs.Append("mt-id", types.YLeaf{"MtId", ospfv2Topology.MtId})
    ospfv2Topology.EntityData.Leafs.Append("metric", types.YLeaf{"Metric", ospfv2Topology.Metric})

    ospfv2Topology.EntityData.YListKeys = []string {"MtId"}

    return &(ospfv2Topology.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv2External
// External LSA
type OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv2External struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. MT-ID for topology enabled on the link. The type
    // is interface{} with range: 0..4294967295.
    MtId interface{}

    // Metric for the topology. The type is interface{} with range: 0..4294967295.
    Metric interface{}

    // Forwarding address. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ForwardingAddress interface{}

    // Route tag. The type is interface{} with range: 0..4294967295.
    ExternalRouteTag interface{}
}

func (ospfv2External *OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv2External) GetEntityData() *types.CommonEntityData {
    ospfv2External.EntityData.YFilter = ospfv2External.YFilter
    ospfv2External.EntityData.YangName = "ospfv2-external"
    ospfv2External.EntityData.BundleName = "cisco_ios_xe"
    ospfv2External.EntityData.ParentYangName = "area-scope-lsa"
    ospfv2External.EntityData.SegmentPath = "ospfv2-external" + types.AddKeyToken(ospfv2External.MtId, "mt-id")
    ospfv2External.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/area-scope-lsa/area-scope-lsa/" + ospfv2External.EntityData.SegmentPath
    ospfv2External.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv2External.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv2External.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv2External.EntityData.Children = types.NewOrderedMap()
    ospfv2External.EntityData.Leafs = types.NewOrderedMap()
    ospfv2External.EntityData.Leafs.Append("mt-id", types.YLeaf{"MtId", ospfv2External.MtId})
    ospfv2External.EntityData.Leafs.Append("metric", types.YLeaf{"Metric", ospfv2External.Metric})
    ospfv2External.EntityData.Leafs.Append("forwarding-address", types.YLeaf{"ForwardingAddress", ospfv2External.ForwardingAddress})
    ospfv2External.EntityData.Leafs.Append("external-route-tag", types.YLeaf{"ExternalRouteTag", ospfv2External.ExternalRouteTag})

    ospfv2External.EntityData.YListKeys = []string {"MtId"}

    return &(ospfv2External.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Lsa
// OSPFv3 LSA
type OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Lsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Decoded OSPFv3 LSA header.
    Header OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Lsa_Header

    // Decoded OSPFv3 LSA body.
    LsaBody OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Lsa_LsaBody
}

func (ospfv3Lsa *OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Lsa) GetEntityData() *types.CommonEntityData {
    ospfv3Lsa.EntityData.YFilter = ospfv3Lsa.YFilter
    ospfv3Lsa.EntityData.YangName = "ospfv3-lsa"
    ospfv3Lsa.EntityData.BundleName = "cisco_ios_xe"
    ospfv3Lsa.EntityData.ParentYangName = "area-scope-lsa"
    ospfv3Lsa.EntityData.SegmentPath = "ospfv3-lsa"
    ospfv3Lsa.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/area-scope-lsa/area-scope-lsa/" + ospfv3Lsa.EntityData.SegmentPath
    ospfv3Lsa.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv3Lsa.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv3Lsa.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv3Lsa.EntityData.Children = types.NewOrderedMap()
    ospfv3Lsa.EntityData.Children.Append("header", types.YChild{"Header", &ospfv3Lsa.Header})
    ospfv3Lsa.EntityData.Children.Append("lsa-body", types.YChild{"LsaBody", &ospfv3Lsa.LsaBody})
    ospfv3Lsa.EntityData.Leafs = types.NewOrderedMap()

    ospfv3Lsa.EntityData.YListKeys = []string {}

    return &(ospfv3Lsa.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Lsa_Header
// Decoded OSPFv3 LSA header
type OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Lsa_Header struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA ID. The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    LsaId interface{}

    // OSPFv3 LSA options. The type is map[string]bool.
    LsaHdrOptions interface{}

    // LSA header.
    LsaHeader OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Lsa_Header_LsaHeader
}

func (header *OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Lsa_Header) GetEntityData() *types.CommonEntityData {
    header.EntityData.YFilter = header.YFilter
    header.EntityData.YangName = "header"
    header.EntityData.BundleName = "cisco_ios_xe"
    header.EntityData.ParentYangName = "ospfv3-lsa"
    header.EntityData.SegmentPath = "header"
    header.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/area-scope-lsa/area-scope-lsa/ospfv3-lsa/" + header.EntityData.SegmentPath
    header.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    header.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    header.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    header.EntityData.Children = types.NewOrderedMap()
    header.EntityData.Children.Append("lsa-header", types.YChild{"LsaHeader", &header.LsaHeader})
    header.EntityData.Leafs = types.NewOrderedMap()
    header.EntityData.Leafs.Append("lsa-id", types.YLeaf{"LsaId", header.LsaId})
    header.EntityData.Leafs.Append("lsa-hdr-options", types.YLeaf{"LsaHdrOptions", header.LsaHdrOptions})

    header.EntityData.YListKeys = []string {}

    return &(header.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Lsa_Header_LsaHeader
// LSA header
type OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Lsa_Header_LsaHeader struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA age. The type is interface{} with range: 0..65535.
    Age interface{}

    // LSA type. The type is interface{} with range: 0..65535.
    Type interface{}

    // LSA advertising router. The type is interface{} with range: 0..4294967295.
    AdvRouter interface{}

    // LSA sequence number. The type is string.
    SeqNum interface{}

    // LSA checksum. The type is string.
    Checksum interface{}

    // LSA length. The type is interface{} with range: 0..65535.
    Length interface{}
}

func (lsaHeader *OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Lsa_Header_LsaHeader) GetEntityData() *types.CommonEntityData {
    lsaHeader.EntityData.YFilter = lsaHeader.YFilter
    lsaHeader.EntityData.YangName = "lsa-header"
    lsaHeader.EntityData.BundleName = "cisco_ios_xe"
    lsaHeader.EntityData.ParentYangName = "header"
    lsaHeader.EntityData.SegmentPath = "lsa-header"
    lsaHeader.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/area-scope-lsa/area-scope-lsa/ospfv3-lsa/header/" + lsaHeader.EntityData.SegmentPath
    lsaHeader.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lsaHeader.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lsaHeader.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lsaHeader.EntityData.Children = types.NewOrderedMap()
    lsaHeader.EntityData.Leafs = types.NewOrderedMap()
    lsaHeader.EntityData.Leafs.Append("age", types.YLeaf{"Age", lsaHeader.Age})
    lsaHeader.EntityData.Leafs.Append("type", types.YLeaf{"Type", lsaHeader.Type})
    lsaHeader.EntityData.Leafs.Append("adv-router", types.YLeaf{"AdvRouter", lsaHeader.AdvRouter})
    lsaHeader.EntityData.Leafs.Append("seq-num", types.YLeaf{"SeqNum", lsaHeader.SeqNum})
    lsaHeader.EntityData.Leafs.Append("checksum", types.YLeaf{"Checksum", lsaHeader.Checksum})
    lsaHeader.EntityData.Leafs.Append("length", types.YLeaf{"Length", lsaHeader.Length})

    lsaHeader.EntityData.YListKeys = []string {}

    return &(lsaHeader.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Lsa_LsaBody
// Decoded OSPFv3 LSA body
type OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Lsa_LsaBody struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA options. The type is map[string]bool.
    LsaFlagOptions interface{}

    // LSA Body Flags. The type is map[string]bool.
    LsaBodyFlags interface{}

    // OSPFv3 network.
    Network OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Lsa_LsaBody_Network

    // OSPFv3 inter area prefix.
    Prefix OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Lsa_LsaBody_Prefix

    // OSPFv3 inter area router.
    IaRouter OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Lsa_LsaBody_IaRouter

    // OSPFv3 LSA external.
    LsaExternal OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Lsa_LsaBody_LsaExternal

    // OSPFv3 NSSA.
    Nssa OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Lsa_LsaBody_Nssa

    // OSPFv3 Link data.
    LinkData OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Lsa_LsaBody_LinkData

    // OSPFv3 Intra area prefixes.
    IaPrefix OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Lsa_LsaBody_IaPrefix
}

func (lsaBody *OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Lsa_LsaBody) GetEntityData() *types.CommonEntityData {
    lsaBody.EntityData.YFilter = lsaBody.YFilter
    lsaBody.EntityData.YangName = "lsa-body"
    lsaBody.EntityData.BundleName = "cisco_ios_xe"
    lsaBody.EntityData.ParentYangName = "ospfv3-lsa"
    lsaBody.EntityData.SegmentPath = "lsa-body"
    lsaBody.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/area-scope-lsa/area-scope-lsa/ospfv3-lsa/" + lsaBody.EntityData.SegmentPath
    lsaBody.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lsaBody.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lsaBody.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lsaBody.EntityData.Children = types.NewOrderedMap()
    lsaBody.EntityData.Children.Append("network", types.YChild{"Network", &lsaBody.Network})
    lsaBody.EntityData.Children.Append("prefix", types.YChild{"Prefix", &lsaBody.Prefix})
    lsaBody.EntityData.Children.Append("ia-router", types.YChild{"IaRouter", &lsaBody.IaRouter})
    lsaBody.EntityData.Children.Append("lsa-external", types.YChild{"LsaExternal", &lsaBody.LsaExternal})
    lsaBody.EntityData.Children.Append("nssa", types.YChild{"Nssa", &lsaBody.Nssa})
    lsaBody.EntityData.Children.Append("link-data", types.YChild{"LinkData", &lsaBody.LinkData})
    lsaBody.EntityData.Children.Append("ia-prefix", types.YChild{"IaPrefix", &lsaBody.IaPrefix})
    lsaBody.EntityData.Leafs = types.NewOrderedMap()
    lsaBody.EntityData.Leafs.Append("lsa-flag-options", types.YLeaf{"LsaFlagOptions", lsaBody.LsaFlagOptions})
    lsaBody.EntityData.Leafs.Append("lsa-body-flags", types.YLeaf{"LsaBodyFlags", lsaBody.LsaBodyFlags})

    lsaBody.EntityData.YListKeys = []string {}

    return &(lsaBody.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Lsa_LsaBody_Network
// OSPFv3 network
type OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Lsa_LsaBody_Network struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of the routers attached to the network. The type is slice of
    // interface{} with range: 0..4294967295.
    AttachedRouter []interface{}

    // Network LSA options. The type is map[string]bool.
    LsaNetOptions interface{}
}

func (network *OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Lsa_LsaBody_Network) GetEntityData() *types.CommonEntityData {
    network.EntityData.YFilter = network.YFilter
    network.EntityData.YangName = "network"
    network.EntityData.BundleName = "cisco_ios_xe"
    network.EntityData.ParentYangName = "lsa-body"
    network.EntityData.SegmentPath = "network"
    network.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/area-scope-lsa/area-scope-lsa/ospfv3-lsa/lsa-body/" + network.EntityData.SegmentPath
    network.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    network.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    network.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    network.EntityData.Children = types.NewOrderedMap()
    network.EntityData.Leafs = types.NewOrderedMap()
    network.EntityData.Leafs.Append("attached-router", types.YLeaf{"AttachedRouter", network.AttachedRouter})
    network.EntityData.Leafs.Append("lsa-net-options", types.YLeaf{"LsaNetOptions", network.LsaNetOptions})

    network.EntityData.YListKeys = []string {}

    return &(network.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Lsa_LsaBody_Prefix
// OSPFv3 inter area prefix
type OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Lsa_LsaBody_Prefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Metric. The type is interface{} with range: 0..4294967295.
    Metric interface{}

    // Inter area Prefix. The type is string.
    IaPrefix interface{}

    // Inter area prefix options. The type is string.
    IaPrefixOptions interface{}
}

func (prefix *OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Lsa_LsaBody_Prefix) GetEntityData() *types.CommonEntityData {
    prefix.EntityData.YFilter = prefix.YFilter
    prefix.EntityData.YangName = "prefix"
    prefix.EntityData.BundleName = "cisco_ios_xe"
    prefix.EntityData.ParentYangName = "lsa-body"
    prefix.EntityData.SegmentPath = "prefix"
    prefix.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/area-scope-lsa/area-scope-lsa/ospfv3-lsa/lsa-body/" + prefix.EntityData.SegmentPath
    prefix.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    prefix.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    prefix.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    prefix.EntityData.Children = types.NewOrderedMap()
    prefix.EntityData.Leafs = types.NewOrderedMap()
    prefix.EntityData.Leafs.Append("metric", types.YLeaf{"Metric", prefix.Metric})
    prefix.EntityData.Leafs.Append("ia-prefix", types.YLeaf{"IaPrefix", prefix.IaPrefix})
    prefix.EntityData.Leafs.Append("ia-prefix-options", types.YLeaf{"IaPrefixOptions", prefix.IaPrefixOptions})

    prefix.EntityData.YListKeys = []string {}

    return &(prefix.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Lsa_LsaBody_IaRouter
// OSPFv3 inter area router
type OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Lsa_LsaBody_IaRouter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Metric. The type is interface{} with range: 0..4294967295.
    Metric interface{}

    // Router ID of the router being described by the LSA. The type is interface{}
    // with range: 0..4294967295.
    DestinationRouterId interface{}

    // Inter area LSA options. The type is map[string]bool.
    LsaIaOptions interface{}
}

func (iaRouter *OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Lsa_LsaBody_IaRouter) GetEntityData() *types.CommonEntityData {
    iaRouter.EntityData.YFilter = iaRouter.YFilter
    iaRouter.EntityData.YangName = "ia-router"
    iaRouter.EntityData.BundleName = "cisco_ios_xe"
    iaRouter.EntityData.ParentYangName = "lsa-body"
    iaRouter.EntityData.SegmentPath = "ia-router"
    iaRouter.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/area-scope-lsa/area-scope-lsa/ospfv3-lsa/lsa-body/" + iaRouter.EntityData.SegmentPath
    iaRouter.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    iaRouter.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    iaRouter.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    iaRouter.EntityData.Children = types.NewOrderedMap()
    iaRouter.EntityData.Leafs = types.NewOrderedMap()
    iaRouter.EntityData.Leafs.Append("metric", types.YLeaf{"Metric", iaRouter.Metric})
    iaRouter.EntityData.Leafs.Append("destination-router-id", types.YLeaf{"DestinationRouterId", iaRouter.DestinationRouterId})
    iaRouter.EntityData.Leafs.Append("lsa-ia-options", types.YLeaf{"LsaIaOptions", iaRouter.LsaIaOptions})

    iaRouter.EntityData.YListKeys = []string {}

    return &(iaRouter.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Lsa_LsaBody_LsaExternal
// OSPFv3 LSA external
type OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Lsa_LsaBody_LsaExternal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Metric. The type is interface{} with range: 0..4294967295.
    Metric interface{}

    // Referenced Link State type. The type is interface{} with range: 0..65535.
    ReferencedLsType interface{}

    // Prefix. The type is string.
    ExternalPrefix interface{}

    // Prefix options. The type is string.
    ExternalPrefixOptions interface{}

    // Forwarding address. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ForwardingAddress interface{}

    // Route tag. The type is interface{} with range: 0..4294967295.
    ExternalRouteTag interface{}

    // Referenced Link State ID. The type is interface{} with range:
    // 0..4294967295.
    ReferencedLinkStateId interface{}

    // LSA Flags.
    Flags OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Lsa_LsaBody_LsaExternal_Flags
}

func (lsaExternal *OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Lsa_LsaBody_LsaExternal) GetEntityData() *types.CommonEntityData {
    lsaExternal.EntityData.YFilter = lsaExternal.YFilter
    lsaExternal.EntityData.YangName = "lsa-external"
    lsaExternal.EntityData.BundleName = "cisco_ios_xe"
    lsaExternal.EntityData.ParentYangName = "lsa-body"
    lsaExternal.EntityData.SegmentPath = "lsa-external"
    lsaExternal.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/area-scope-lsa/area-scope-lsa/ospfv3-lsa/lsa-body/" + lsaExternal.EntityData.SegmentPath
    lsaExternal.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lsaExternal.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lsaExternal.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lsaExternal.EntityData.Children = types.NewOrderedMap()
    lsaExternal.EntityData.Children.Append("flags", types.YChild{"Flags", &lsaExternal.Flags})
    lsaExternal.EntityData.Leafs = types.NewOrderedMap()
    lsaExternal.EntityData.Leafs.Append("metric", types.YLeaf{"Metric", lsaExternal.Metric})
    lsaExternal.EntityData.Leafs.Append("referenced-ls-type", types.YLeaf{"ReferencedLsType", lsaExternal.ReferencedLsType})
    lsaExternal.EntityData.Leafs.Append("external-prefix", types.YLeaf{"ExternalPrefix", lsaExternal.ExternalPrefix})
    lsaExternal.EntityData.Leafs.Append("external-prefix-options", types.YLeaf{"ExternalPrefixOptions", lsaExternal.ExternalPrefixOptions})
    lsaExternal.EntityData.Leafs.Append("forwarding-address", types.YLeaf{"ForwardingAddress", lsaExternal.ForwardingAddress})
    lsaExternal.EntityData.Leafs.Append("external-route-tag", types.YLeaf{"ExternalRouteTag", lsaExternal.ExternalRouteTag})
    lsaExternal.EntityData.Leafs.Append("referenced-link-state-id", types.YLeaf{"ReferencedLinkStateId", lsaExternal.ReferencedLinkStateId})

    lsaExternal.EntityData.YListKeys = []string {}

    return &(lsaExternal.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Lsa_LsaBody_LsaExternal_Flags
// LSA Flags
type OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Lsa_LsaBody_LsaExternal_Flags struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // When set, the metric specified is a Type 2 external metric. The type is
    // bool.
    EFlag interface{}
}

func (flags *OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Lsa_LsaBody_LsaExternal_Flags) GetEntityData() *types.CommonEntityData {
    flags.EntityData.YFilter = flags.YFilter
    flags.EntityData.YangName = "flags"
    flags.EntityData.BundleName = "cisco_ios_xe"
    flags.EntityData.ParentYangName = "lsa-external"
    flags.EntityData.SegmentPath = "flags"
    flags.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/area-scope-lsa/area-scope-lsa/ospfv3-lsa/lsa-body/lsa-external/" + flags.EntityData.SegmentPath
    flags.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    flags.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    flags.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    flags.EntityData.Children = types.NewOrderedMap()
    flags.EntityData.Leafs = types.NewOrderedMap()
    flags.EntityData.Leafs.Append("e-flag", types.YLeaf{"EFlag", flags.EFlag})

    flags.EntityData.YListKeys = []string {}

    return &(flags.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Lsa_LsaBody_Nssa
// OSPFv3 NSSA
type OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Lsa_LsaBody_Nssa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // NSSA LSA.
    LsaNssaExternal OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Lsa_LsaBody_Nssa_LsaNssaExternal
}

func (nssa *OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Lsa_LsaBody_Nssa) GetEntityData() *types.CommonEntityData {
    nssa.EntityData.YFilter = nssa.YFilter
    nssa.EntityData.YangName = "nssa"
    nssa.EntityData.BundleName = "cisco_ios_xe"
    nssa.EntityData.ParentYangName = "lsa-body"
    nssa.EntityData.SegmentPath = "nssa"
    nssa.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/area-scope-lsa/area-scope-lsa/ospfv3-lsa/lsa-body/" + nssa.EntityData.SegmentPath
    nssa.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nssa.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nssa.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nssa.EntityData.Children = types.NewOrderedMap()
    nssa.EntityData.Children.Append("lsa-nssa-external", types.YChild{"LsaNssaExternal", &nssa.LsaNssaExternal})
    nssa.EntityData.Leafs = types.NewOrderedMap()

    nssa.EntityData.YListKeys = []string {}

    return &(nssa.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Lsa_LsaBody_Nssa_LsaNssaExternal
// NSSA LSA
type OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Lsa_LsaBody_Nssa_LsaNssaExternal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Metric. The type is interface{} with range: 0..4294967295.
    Metric interface{}

    // Referenced Link State type. The type is interface{} with range: 0..65535.
    ReferencedLsType interface{}

    // Prefix. The type is string.
    ExternalPrefix interface{}

    // Prefix options. The type is string.
    ExternalPrefixOptions interface{}

    // Forwarding address. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ForwardingAddress interface{}

    // Route tag. The type is interface{} with range: 0..4294967295.
    ExternalRouteTag interface{}

    // Referenced Link State ID. The type is interface{} with range:
    // 0..4294967295.
    ReferencedLinkStateId interface{}

    // LSA Flags.
    Flags OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Lsa_LsaBody_Nssa_LsaNssaExternal_Flags
}

func (lsaNssaExternal *OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Lsa_LsaBody_Nssa_LsaNssaExternal) GetEntityData() *types.CommonEntityData {
    lsaNssaExternal.EntityData.YFilter = lsaNssaExternal.YFilter
    lsaNssaExternal.EntityData.YangName = "lsa-nssa-external"
    lsaNssaExternal.EntityData.BundleName = "cisco_ios_xe"
    lsaNssaExternal.EntityData.ParentYangName = "nssa"
    lsaNssaExternal.EntityData.SegmentPath = "lsa-nssa-external"
    lsaNssaExternal.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/area-scope-lsa/area-scope-lsa/ospfv3-lsa/lsa-body/nssa/" + lsaNssaExternal.EntityData.SegmentPath
    lsaNssaExternal.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lsaNssaExternal.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lsaNssaExternal.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lsaNssaExternal.EntityData.Children = types.NewOrderedMap()
    lsaNssaExternal.EntityData.Children.Append("flags", types.YChild{"Flags", &lsaNssaExternal.Flags})
    lsaNssaExternal.EntityData.Leafs = types.NewOrderedMap()
    lsaNssaExternal.EntityData.Leafs.Append("metric", types.YLeaf{"Metric", lsaNssaExternal.Metric})
    lsaNssaExternal.EntityData.Leafs.Append("referenced-ls-type", types.YLeaf{"ReferencedLsType", lsaNssaExternal.ReferencedLsType})
    lsaNssaExternal.EntityData.Leafs.Append("external-prefix", types.YLeaf{"ExternalPrefix", lsaNssaExternal.ExternalPrefix})
    lsaNssaExternal.EntityData.Leafs.Append("external-prefix-options", types.YLeaf{"ExternalPrefixOptions", lsaNssaExternal.ExternalPrefixOptions})
    lsaNssaExternal.EntityData.Leafs.Append("forwarding-address", types.YLeaf{"ForwardingAddress", lsaNssaExternal.ForwardingAddress})
    lsaNssaExternal.EntityData.Leafs.Append("external-route-tag", types.YLeaf{"ExternalRouteTag", lsaNssaExternal.ExternalRouteTag})
    lsaNssaExternal.EntityData.Leafs.Append("referenced-link-state-id", types.YLeaf{"ReferencedLinkStateId", lsaNssaExternal.ReferencedLinkStateId})

    lsaNssaExternal.EntityData.YListKeys = []string {}

    return &(lsaNssaExternal.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Lsa_LsaBody_Nssa_LsaNssaExternal_Flags
// LSA Flags
type OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Lsa_LsaBody_Nssa_LsaNssaExternal_Flags struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // When set, the metric specified is a Type 2 external metric. The type is
    // bool.
    EFlag interface{}
}

func (flags *OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Lsa_LsaBody_Nssa_LsaNssaExternal_Flags) GetEntityData() *types.CommonEntityData {
    flags.EntityData.YFilter = flags.YFilter
    flags.EntityData.YangName = "flags"
    flags.EntityData.BundleName = "cisco_ios_xe"
    flags.EntityData.ParentYangName = "lsa-nssa-external"
    flags.EntityData.SegmentPath = "flags"
    flags.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/area-scope-lsa/area-scope-lsa/ospfv3-lsa/lsa-body/nssa/lsa-nssa-external/" + flags.EntityData.SegmentPath
    flags.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    flags.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    flags.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    flags.EntityData.Children = types.NewOrderedMap()
    flags.EntityData.Leafs = types.NewOrderedMap()
    flags.EntityData.Leafs.Append("e-flag", types.YLeaf{"EFlag", flags.EFlag})

    flags.EntityData.YListKeys = []string {}

    return &(flags.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Lsa_LsaBody_LinkData
// OSPFv3 Link data
type OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Lsa_LsaBody_LinkData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Router priority of the interce. The type is interface{} with range: 0..255.
    RtrPriority interface{}

    // The originating router's link-local interface address on the link. The type
    // is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    LinkLocalInterfaceAddress interface{}

    // Number of prefixes. The type is interface{} with range: 0..4294967295.
    NumOfPrefixes interface{}

    // Link data LSA options. The type is map[string]bool.
    LsaIdOptions interface{}
}

func (linkData *OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Lsa_LsaBody_LinkData) GetEntityData() *types.CommonEntityData {
    linkData.EntityData.YFilter = linkData.YFilter
    linkData.EntityData.YangName = "link-data"
    linkData.EntityData.BundleName = "cisco_ios_xe"
    linkData.EntityData.ParentYangName = "lsa-body"
    linkData.EntityData.SegmentPath = "link-data"
    linkData.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/area-scope-lsa/area-scope-lsa/ospfv3-lsa/lsa-body/" + linkData.EntityData.SegmentPath
    linkData.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    linkData.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    linkData.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    linkData.EntityData.Children = types.NewOrderedMap()
    linkData.EntityData.Leafs = types.NewOrderedMap()
    linkData.EntityData.Leafs.Append("rtr-priority", types.YLeaf{"RtrPriority", linkData.RtrPriority})
    linkData.EntityData.Leafs.Append("link-local-interface-address", types.YLeaf{"LinkLocalInterfaceAddress", linkData.LinkLocalInterfaceAddress})
    linkData.EntityData.Leafs.Append("num-of-prefixes", types.YLeaf{"NumOfPrefixes", linkData.NumOfPrefixes})
    linkData.EntityData.Leafs.Append("lsa-id-options", types.YLeaf{"LsaIdOptions", linkData.LsaIdOptions})

    linkData.EntityData.YListKeys = []string {}

    return &(linkData.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Lsa_LsaBody_IaPrefix
// OSPFv3 Intra area prefixes
type OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Lsa_LsaBody_IaPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Referenced Link State type. The type is interface{} with range: 0..65535.
    ReferencedLsType interface{}

    // Referenced Link State ID. The type is interface{} with range:
    // 0..4294967295.
    ReferencedLinkStateId interface{}

    // Referenced Advertising Router. The type is one of the following types:
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ReferencedAdvRouter interface{}

    // Number of prefixes. The type is interface{} with range: 0..65535.
    NumOfPrefixes interface{}
}

func (iaPrefix *OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Lsa_LsaBody_IaPrefix) GetEntityData() *types.CommonEntityData {
    iaPrefix.EntityData.YFilter = iaPrefix.YFilter
    iaPrefix.EntityData.YangName = "ia-prefix"
    iaPrefix.EntityData.BundleName = "cisco_ios_xe"
    iaPrefix.EntityData.ParentYangName = "lsa-body"
    iaPrefix.EntityData.SegmentPath = "ia-prefix"
    iaPrefix.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/area-scope-lsa/area-scope-lsa/ospfv3-lsa/lsa-body/" + iaPrefix.EntityData.SegmentPath
    iaPrefix.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    iaPrefix.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    iaPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    iaPrefix.EntityData.Children = types.NewOrderedMap()
    iaPrefix.EntityData.Leafs = types.NewOrderedMap()
    iaPrefix.EntityData.Leafs.Append("referenced-ls-type", types.YLeaf{"ReferencedLsType", iaPrefix.ReferencedLsType})
    iaPrefix.EntityData.Leafs.Append("referenced-link-state-id", types.YLeaf{"ReferencedLinkStateId", iaPrefix.ReferencedLinkStateId})
    iaPrefix.EntityData.Leafs.Append("referenced-adv-router", types.YLeaf{"ReferencedAdvRouter", iaPrefix.ReferencedAdvRouter})
    iaPrefix.EntityData.Leafs.Append("num-of-prefixes", types.YLeaf{"NumOfPrefixes", iaPrefix.NumOfPrefixes})

    iaPrefix.EntityData.YListKeys = []string {}

    return &(iaPrefix.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Link
// OSPFv3 links
type OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Link struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface ID. The type is interface{} with range:
    // 0..4294967295.
    InterfaceId interface{}

    // This attribute is a key. Neighbor interface ID. The type is interface{}
    // with range: 0..4294967295.
    NeighborInterfaceId interface{}

    // This attribute is a key. Neighbor router ID. The type is interface{} with
    // range: 0..4294967295.
    NeighborRouterId interface{}

    // Link type. The type is interface{} with range: 0..255.
    Type interface{}

    // Metric. The type is interface{} with range: 0..65535.
    Metric interface{}
}

func (ospfv3Link *OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Link) GetEntityData() *types.CommonEntityData {
    ospfv3Link.EntityData.YFilter = ospfv3Link.YFilter
    ospfv3Link.EntityData.YangName = "ospfv3-link"
    ospfv3Link.EntityData.BundleName = "cisco_ios_xe"
    ospfv3Link.EntityData.ParentYangName = "area-scope-lsa"
    ospfv3Link.EntityData.SegmentPath = "ospfv3-link" + types.AddKeyToken(ospfv3Link.InterfaceId, "interface-id") + types.AddKeyToken(ospfv3Link.NeighborInterfaceId, "neighbor-interface-id") + types.AddKeyToken(ospfv3Link.NeighborRouterId, "neighbor-router-id")
    ospfv3Link.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/area-scope-lsa/area-scope-lsa/" + ospfv3Link.EntityData.SegmentPath
    ospfv3Link.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv3Link.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv3Link.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv3Link.EntityData.Children = types.NewOrderedMap()
    ospfv3Link.EntityData.Leafs = types.NewOrderedMap()
    ospfv3Link.EntityData.Leafs.Append("interface-id", types.YLeaf{"InterfaceId", ospfv3Link.InterfaceId})
    ospfv3Link.EntityData.Leafs.Append("neighbor-interface-id", types.YLeaf{"NeighborInterfaceId", ospfv3Link.NeighborInterfaceId})
    ospfv3Link.EntityData.Leafs.Append("neighbor-router-id", types.YLeaf{"NeighborRouterId", ospfv3Link.NeighborRouterId})
    ospfv3Link.EntityData.Leafs.Append("type", types.YLeaf{"Type", ospfv3Link.Type})
    ospfv3Link.EntityData.Leafs.Append("metric", types.YLeaf{"Metric", ospfv3Link.Metric})

    ospfv3Link.EntityData.YListKeys = []string {"InterfaceId", "NeighborInterfaceId", "NeighborRouterId"}

    return &(ospfv3Link.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Prefix
// OSPFv3 prefix-list
type OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Prefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Prefix. The type is string.
    Prefix interface{}

    // Prefix options. The type is string.
    PrefixOptions interface{}
}

func (ospfv3Prefix *OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3Prefix) GetEntityData() *types.CommonEntityData {
    ospfv3Prefix.EntityData.YFilter = ospfv3Prefix.YFilter
    ospfv3Prefix.EntityData.YangName = "ospfv3-prefix"
    ospfv3Prefix.EntityData.BundleName = "cisco_ios_xe"
    ospfv3Prefix.EntityData.ParentYangName = "area-scope-lsa"
    ospfv3Prefix.EntityData.SegmentPath = "ospfv3-prefix" + types.AddKeyToken(ospfv3Prefix.Prefix, "prefix")
    ospfv3Prefix.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/area-scope-lsa/area-scope-lsa/" + ospfv3Prefix.EntityData.SegmentPath
    ospfv3Prefix.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv3Prefix.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv3Prefix.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv3Prefix.EntityData.Children = types.NewOrderedMap()
    ospfv3Prefix.EntityData.Leafs = types.NewOrderedMap()
    ospfv3Prefix.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", ospfv3Prefix.Prefix})
    ospfv3Prefix.EntityData.Leafs.Append("prefix-options", types.YLeaf{"PrefixOptions", ospfv3Prefix.PrefixOptions})

    ospfv3Prefix.EntityData.YListKeys = []string {"Prefix"}

    return &(ospfv3Prefix.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3IaPrefix
// OSPFv3 intra-area prefix-list
type OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3IaPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Prefix. The type is string.
    Prefix interface{}

    // Prefix options. The type is string.
    PrefixOptions interface{}
}

func (ospfv3IaPrefix *OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_Ospfv3IaPrefix) GetEntityData() *types.CommonEntityData {
    ospfv3IaPrefix.EntityData.YFilter = ospfv3IaPrefix.YFilter
    ospfv3IaPrefix.EntityData.YangName = "ospfv3-ia-prefix"
    ospfv3IaPrefix.EntityData.BundleName = "cisco_ios_xe"
    ospfv3IaPrefix.EntityData.ParentYangName = "area-scope-lsa"
    ospfv3IaPrefix.EntityData.SegmentPath = "ospfv3-ia-prefix" + types.AddKeyToken(ospfv3IaPrefix.Prefix, "prefix")
    ospfv3IaPrefix.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/ospf-area/area-scope-lsa/area-scope-lsa/" + ospfv3IaPrefix.EntityData.SegmentPath
    ospfv3IaPrefix.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv3IaPrefix.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv3IaPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv3IaPrefix.EntityData.Children = types.NewOrderedMap()
    ospfv3IaPrefix.EntityData.Leafs = types.NewOrderedMap()
    ospfv3IaPrefix.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", ospfv3IaPrefix.Prefix})
    ospfv3IaPrefix.EntityData.Leafs.Append("prefix-options", types.YLeaf{"PrefixOptions", ospfv3IaPrefix.PrefixOptions})

    ospfv3IaPrefix.EntityData.YListKeys = []string {"Prefix"}

    return &(ospfv3IaPrefix.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas
// List OSPF link scope LSA
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. OSPF link scope LSA type. The type is interface{}
    // with range: 0..4294967295.
    LsaType interface{}

    // List of OSPF link scope LSAs. The type is slice of
    // OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa.
    LinkScopeLsa []*OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa

    // List OSPF area scope LSA databases. The type is slice of
    // OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa.
    AreaScopeLsa []*OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa
}

func (linkScopeLsas *OspfOperData_OspfState_OspfInstance_LinkScopeLsas) GetEntityData() *types.CommonEntityData {
    linkScopeLsas.EntityData.YFilter = linkScopeLsas.YFilter
    linkScopeLsas.EntityData.YangName = "link-scope-lsas"
    linkScopeLsas.EntityData.BundleName = "cisco_ios_xe"
    linkScopeLsas.EntityData.ParentYangName = "ospf-instance"
    linkScopeLsas.EntityData.SegmentPath = "link-scope-lsas" + types.AddKeyToken(linkScopeLsas.LsaType, "lsa-type")
    linkScopeLsas.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/" + linkScopeLsas.EntityData.SegmentPath
    linkScopeLsas.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    linkScopeLsas.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    linkScopeLsas.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    linkScopeLsas.EntityData.Children = types.NewOrderedMap()
    linkScopeLsas.EntityData.Children.Append("link-scope-lsa", types.YChild{"LinkScopeLsa", nil})
    for i := range linkScopeLsas.LinkScopeLsa {
        linkScopeLsas.EntityData.Children.Append(types.GetSegmentPath(linkScopeLsas.LinkScopeLsa[i]), types.YChild{"LinkScopeLsa", linkScopeLsas.LinkScopeLsa[i]})
    }
    linkScopeLsas.EntityData.Children.Append("area-scope-lsa", types.YChild{"AreaScopeLsa", nil})
    for i := range linkScopeLsas.AreaScopeLsa {
        linkScopeLsas.EntityData.Children.Append(types.GetSegmentPath(linkScopeLsas.AreaScopeLsa[i]), types.YChild{"AreaScopeLsa", linkScopeLsas.AreaScopeLsa[i]})
    }
    linkScopeLsas.EntityData.Leafs = types.NewOrderedMap()
    linkScopeLsas.EntityData.Leafs.Append("lsa-type", types.YLeaf{"LsaType", linkScopeLsas.LsaType})

    linkScopeLsas.EntityData.YListKeys = []string {"LsaType"}

    return &(linkScopeLsas.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa
// List of OSPF link scope LSAs
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. LSA ID. The type is interface{} with range:
    // 0..4294967295.
    LsaId interface{}

    // This attribute is a key. Advertising router. The type is one of the
    // following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    AdvRouter interface{}

    // The OSPF LSA body is fully decoded. The type is bool.
    DecodedCompleted interface{}

    // The complete LSA in network byte order as received/sent over the wire. The
    // type is slice of interface{} with range: 0..255.
    RawData []interface{}

    // Version. The type is interface{} with range: 0..4294967295.
    Version interface{}

    // Router address. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    RouterAddress interface{}

    // OSPFv2 LSA.
    Ospfv2Lsa OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2Lsa

    // OSPFv2 LSA link. The type is slice of
    // OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2Link.
    Ospfv2Link []*OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2Link

    // Summary LSA. The type is slice of
    // OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2Topology.
    Ospfv2Topology []*OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2Topology

    // External LSA. The type is slice of
    // OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2External.
    Ospfv2External []*OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2External

    // OSPFv2 Unknown TLV. The type is slice of
    // OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2UnknownTlv.
    Ospfv2UnknownTlv []*OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2UnknownTlv

    // OSPFv3 LSA.
    Ospfv3LsaVal OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal

    // OSPFv3 links. The type is slice of
    // OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3Link.
    Ospfv3Link []*OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3Link

    // OSPFv3 prefix-list. The type is slice of
    // OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3PrefixList.
    Ospfv3PrefixList []*OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3PrefixList

    // OSPFv3 intra-area prefix-list. The type is slice of
    // OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3IaPrefix.
    Ospfv3IaPrefix []*OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3IaPrefix

    // OSPF multi-topology interface augmentation. The type is slice of
    // OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_MultiTopology.
    MultiTopology []*OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_MultiTopology

    // Link TLV.
    Tlv OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Tlv

    // OSPFv2 Unknown sub TLV. The type is slice of
    // OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_UnknownSubTlv.
    UnknownSubTlv []*OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_UnknownSubTlv
}

func (linkScopeLsa *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa) GetEntityData() *types.CommonEntityData {
    linkScopeLsa.EntityData.YFilter = linkScopeLsa.YFilter
    linkScopeLsa.EntityData.YangName = "link-scope-lsa"
    linkScopeLsa.EntityData.BundleName = "cisco_ios_xe"
    linkScopeLsa.EntityData.ParentYangName = "link-scope-lsas"
    linkScopeLsa.EntityData.SegmentPath = "link-scope-lsa" + types.AddKeyToken(linkScopeLsa.LsaId, "lsa-id") + types.AddKeyToken(linkScopeLsa.AdvRouter, "adv-router")
    linkScopeLsa.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/link-scope-lsas/" + linkScopeLsa.EntityData.SegmentPath
    linkScopeLsa.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    linkScopeLsa.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    linkScopeLsa.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    linkScopeLsa.EntityData.Children = types.NewOrderedMap()
    linkScopeLsa.EntityData.Children.Append("ospfv2-lsa", types.YChild{"Ospfv2Lsa", &linkScopeLsa.Ospfv2Lsa})
    linkScopeLsa.EntityData.Children.Append("ospfv2-link", types.YChild{"Ospfv2Link", nil})
    for i := range linkScopeLsa.Ospfv2Link {
        linkScopeLsa.EntityData.Children.Append(types.GetSegmentPath(linkScopeLsa.Ospfv2Link[i]), types.YChild{"Ospfv2Link", linkScopeLsa.Ospfv2Link[i]})
    }
    linkScopeLsa.EntityData.Children.Append("ospfv2-topology", types.YChild{"Ospfv2Topology", nil})
    for i := range linkScopeLsa.Ospfv2Topology {
        linkScopeLsa.EntityData.Children.Append(types.GetSegmentPath(linkScopeLsa.Ospfv2Topology[i]), types.YChild{"Ospfv2Topology", linkScopeLsa.Ospfv2Topology[i]})
    }
    linkScopeLsa.EntityData.Children.Append("ospfv2-external", types.YChild{"Ospfv2External", nil})
    for i := range linkScopeLsa.Ospfv2External {
        linkScopeLsa.EntityData.Children.Append(types.GetSegmentPath(linkScopeLsa.Ospfv2External[i]), types.YChild{"Ospfv2External", linkScopeLsa.Ospfv2External[i]})
    }
    linkScopeLsa.EntityData.Children.Append("ospfv2-unknown-tlv", types.YChild{"Ospfv2UnknownTlv", nil})
    for i := range linkScopeLsa.Ospfv2UnknownTlv {
        linkScopeLsa.EntityData.Children.Append(types.GetSegmentPath(linkScopeLsa.Ospfv2UnknownTlv[i]), types.YChild{"Ospfv2UnknownTlv", linkScopeLsa.Ospfv2UnknownTlv[i]})
    }
    linkScopeLsa.EntityData.Children.Append("ospfv3-lsa-val", types.YChild{"Ospfv3LsaVal", &linkScopeLsa.Ospfv3LsaVal})
    linkScopeLsa.EntityData.Children.Append("ospfv3-link", types.YChild{"Ospfv3Link", nil})
    for i := range linkScopeLsa.Ospfv3Link {
        linkScopeLsa.EntityData.Children.Append(types.GetSegmentPath(linkScopeLsa.Ospfv3Link[i]), types.YChild{"Ospfv3Link", linkScopeLsa.Ospfv3Link[i]})
    }
    linkScopeLsa.EntityData.Children.Append("ospfv3-prefix-list", types.YChild{"Ospfv3PrefixList", nil})
    for i := range linkScopeLsa.Ospfv3PrefixList {
        linkScopeLsa.EntityData.Children.Append(types.GetSegmentPath(linkScopeLsa.Ospfv3PrefixList[i]), types.YChild{"Ospfv3PrefixList", linkScopeLsa.Ospfv3PrefixList[i]})
    }
    linkScopeLsa.EntityData.Children.Append("ospfv3-ia-prefix", types.YChild{"Ospfv3IaPrefix", nil})
    for i := range linkScopeLsa.Ospfv3IaPrefix {
        linkScopeLsa.EntityData.Children.Append(types.GetSegmentPath(linkScopeLsa.Ospfv3IaPrefix[i]), types.YChild{"Ospfv3IaPrefix", linkScopeLsa.Ospfv3IaPrefix[i]})
    }
    linkScopeLsa.EntityData.Children.Append("multi-topology", types.YChild{"MultiTopology", nil})
    for i := range linkScopeLsa.MultiTopology {
        linkScopeLsa.EntityData.Children.Append(types.GetSegmentPath(linkScopeLsa.MultiTopology[i]), types.YChild{"MultiTopology", linkScopeLsa.MultiTopology[i]})
    }
    linkScopeLsa.EntityData.Children.Append("tlv", types.YChild{"Tlv", &linkScopeLsa.Tlv})
    linkScopeLsa.EntityData.Children.Append("unknown-sub-tlv", types.YChild{"UnknownSubTlv", nil})
    for i := range linkScopeLsa.UnknownSubTlv {
        linkScopeLsa.EntityData.Children.Append(types.GetSegmentPath(linkScopeLsa.UnknownSubTlv[i]), types.YChild{"UnknownSubTlv", linkScopeLsa.UnknownSubTlv[i]})
    }
    linkScopeLsa.EntityData.Leafs = types.NewOrderedMap()
    linkScopeLsa.EntityData.Leafs.Append("lsa-id", types.YLeaf{"LsaId", linkScopeLsa.LsaId})
    linkScopeLsa.EntityData.Leafs.Append("adv-router", types.YLeaf{"AdvRouter", linkScopeLsa.AdvRouter})
    linkScopeLsa.EntityData.Leafs.Append("decoded-completed", types.YLeaf{"DecodedCompleted", linkScopeLsa.DecodedCompleted})
    linkScopeLsa.EntityData.Leafs.Append("raw-data", types.YLeaf{"RawData", linkScopeLsa.RawData})
    linkScopeLsa.EntityData.Leafs.Append("version", types.YLeaf{"Version", linkScopeLsa.Version})
    linkScopeLsa.EntityData.Leafs.Append("router-address", types.YLeaf{"RouterAddress", linkScopeLsa.RouterAddress})

    linkScopeLsa.EntityData.YListKeys = []string {"LsaId", "AdvRouter"}

    return &(linkScopeLsa.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2Lsa
// OSPFv2 LSA
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2Lsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Decoded OSPFv2 LSA header data.
    Header OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2Lsa_Header

    // Decoded OSPFv2 LSA body data.
    LsaBody OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2Lsa_LsaBody
}

func (ospfv2Lsa *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2Lsa) GetEntityData() *types.CommonEntityData {
    ospfv2Lsa.EntityData.YFilter = ospfv2Lsa.YFilter
    ospfv2Lsa.EntityData.YangName = "ospfv2-lsa"
    ospfv2Lsa.EntityData.BundleName = "cisco_ios_xe"
    ospfv2Lsa.EntityData.ParentYangName = "link-scope-lsa"
    ospfv2Lsa.EntityData.SegmentPath = "ospfv2-lsa"
    ospfv2Lsa.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/link-scope-lsas/link-scope-lsa/" + ospfv2Lsa.EntityData.SegmentPath
    ospfv2Lsa.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv2Lsa.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv2Lsa.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv2Lsa.EntityData.Children = types.NewOrderedMap()
    ospfv2Lsa.EntityData.Children.Append("header", types.YChild{"Header", &ospfv2Lsa.Header})
    ospfv2Lsa.EntityData.Children.Append("lsa-body", types.YChild{"LsaBody", &ospfv2Lsa.LsaBody})
    ospfv2Lsa.EntityData.Leafs = types.NewOrderedMap()

    ospfv2Lsa.EntityData.YListKeys = []string {}

    return &(ospfv2Lsa.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2Lsa_Header
// Decoded OSPFv2 LSA header data
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2Lsa_Header struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA ID. The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    LsaId interface{}

    // Opaque type. The type is interface{} with range: 0..255.
    OpaqueType interface{}

    // Opaque ID. The type is interface{} with range: 0..4294967295.
    OpaqueId interface{}

    // LSA age. The type is interface{} with range: 0..65535.
    Age interface{}

    // LSA type. The type is interface{} with range: 0..65535.
    Type interface{}

    // LSA advertising router. The type is interface{} with range: 0..4294967295.
    AdvRouter interface{}

    // LSA sequence number. The type is string.
    SeqNum interface{}

    // LSA checksum. The type is string.
    Checksum interface{}

    // LSA length. The type is interface{} with range: 0..65535.
    Length interface{}

    // LSA options. The type is map[string]bool.
    FlagOptions interface{}
}

func (header *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2Lsa_Header) GetEntityData() *types.CommonEntityData {
    header.EntityData.YFilter = header.YFilter
    header.EntityData.YangName = "header"
    header.EntityData.BundleName = "cisco_ios_xe"
    header.EntityData.ParentYangName = "ospfv2-lsa"
    header.EntityData.SegmentPath = "header"
    header.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/link-scope-lsas/link-scope-lsa/ospfv2-lsa/" + header.EntityData.SegmentPath
    header.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    header.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    header.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    header.EntityData.Children = types.NewOrderedMap()
    header.EntityData.Leafs = types.NewOrderedMap()
    header.EntityData.Leafs.Append("lsa-id", types.YLeaf{"LsaId", header.LsaId})
    header.EntityData.Leafs.Append("opaque-type", types.YLeaf{"OpaqueType", header.OpaqueType})
    header.EntityData.Leafs.Append("opaque-id", types.YLeaf{"OpaqueId", header.OpaqueId})
    header.EntityData.Leafs.Append("age", types.YLeaf{"Age", header.Age})
    header.EntityData.Leafs.Append("type", types.YLeaf{"Type", header.Type})
    header.EntityData.Leafs.Append("adv-router", types.YLeaf{"AdvRouter", header.AdvRouter})
    header.EntityData.Leafs.Append("seq-num", types.YLeaf{"SeqNum", header.SeqNum})
    header.EntityData.Leafs.Append("checksum", types.YLeaf{"Checksum", header.Checksum})
    header.EntityData.Leafs.Append("length", types.YLeaf{"Length", header.Length})
    header.EntityData.Leafs.Append("flag-options", types.YLeaf{"FlagOptions", header.FlagOptions})

    header.EntityData.YListKeys = []string {}

    return &(header.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2Lsa_LsaBody
// Decoded OSPFv2 LSA body data
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2Lsa_LsaBody struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of links. The type is interface{} with range: 0..65535.
    NumOfLinks interface{}

    // Summary mask. The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    SummaryMask interface{}

    // External mask. The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ExternalMask interface{}

    // LSA body flags. The type is map[string]bool.
    BodyFlagOptions interface{}

    // Network details.
    Network OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2Lsa_LsaBody_Network
}

func (lsaBody *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2Lsa_LsaBody) GetEntityData() *types.CommonEntityData {
    lsaBody.EntityData.YFilter = lsaBody.YFilter
    lsaBody.EntityData.YangName = "lsa-body"
    lsaBody.EntityData.BundleName = "cisco_ios_xe"
    lsaBody.EntityData.ParentYangName = "ospfv2-lsa"
    lsaBody.EntityData.SegmentPath = "lsa-body"
    lsaBody.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/link-scope-lsas/link-scope-lsa/ospfv2-lsa/" + lsaBody.EntityData.SegmentPath
    lsaBody.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lsaBody.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lsaBody.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lsaBody.EntityData.Children = types.NewOrderedMap()
    lsaBody.EntityData.Children.Append("network", types.YChild{"Network", &lsaBody.Network})
    lsaBody.EntityData.Leafs = types.NewOrderedMap()
    lsaBody.EntityData.Leafs.Append("num-of-links", types.YLeaf{"NumOfLinks", lsaBody.NumOfLinks})
    lsaBody.EntityData.Leafs.Append("summary-mask", types.YLeaf{"SummaryMask", lsaBody.SummaryMask})
    lsaBody.EntityData.Leafs.Append("external-mask", types.YLeaf{"ExternalMask", lsaBody.ExternalMask})
    lsaBody.EntityData.Leafs.Append("body-flag-options", types.YLeaf{"BodyFlagOptions", lsaBody.BodyFlagOptions})

    lsaBody.EntityData.YListKeys = []string {}

    return &(lsaBody.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2Lsa_LsaBody_Network
// Network details
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2Lsa_LsaBody_Network struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP network mask. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NetworkMask interface{}

    // List of the routers attached to the network. The type is slice of
    // interface{} with range: 0..4294967295.
    AttachedRouter []interface{}
}

func (network *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2Lsa_LsaBody_Network) GetEntityData() *types.CommonEntityData {
    network.EntityData.YFilter = network.YFilter
    network.EntityData.YangName = "network"
    network.EntityData.BundleName = "cisco_ios_xe"
    network.EntityData.ParentYangName = "lsa-body"
    network.EntityData.SegmentPath = "network"
    network.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/link-scope-lsas/link-scope-lsa/ospfv2-lsa/lsa-body/" + network.EntityData.SegmentPath
    network.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    network.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    network.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    network.EntityData.Children = types.NewOrderedMap()
    network.EntityData.Leafs = types.NewOrderedMap()
    network.EntityData.Leafs.Append("network-mask", types.YLeaf{"NetworkMask", network.NetworkMask})
    network.EntityData.Leafs.Append("attached-router", types.YLeaf{"AttachedRouter", network.AttachedRouter})

    network.EntityData.YListKeys = []string {}

    return &(network.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2Link
// OSPFv2 LSA link
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2Link struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Link ID. The type is interface{} with range:
    // 0..4294967295.
    LinkId interface{}

    // This attribute is a key. Link data. The type is interface{} with range:
    // 0..4294967295.
    LinkData interface{}

    // Link type. The type is interface{} with range: 0..255.
    Type interface{}

    // Topology specific information. The type is slice of
    // OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2Link_Ospfv2Topology.
    Ospfv2Topology []*OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2Link_Ospfv2Topology
}

func (ospfv2Link *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2Link) GetEntityData() *types.CommonEntityData {
    ospfv2Link.EntityData.YFilter = ospfv2Link.YFilter
    ospfv2Link.EntityData.YangName = "ospfv2-link"
    ospfv2Link.EntityData.BundleName = "cisco_ios_xe"
    ospfv2Link.EntityData.ParentYangName = "link-scope-lsa"
    ospfv2Link.EntityData.SegmentPath = "ospfv2-link" + types.AddKeyToken(ospfv2Link.LinkId, "link-id") + types.AddKeyToken(ospfv2Link.LinkData, "link-data")
    ospfv2Link.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/link-scope-lsas/link-scope-lsa/" + ospfv2Link.EntityData.SegmentPath
    ospfv2Link.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv2Link.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv2Link.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv2Link.EntityData.Children = types.NewOrderedMap()
    ospfv2Link.EntityData.Children.Append("ospfv2-topology", types.YChild{"Ospfv2Topology", nil})
    for i := range ospfv2Link.Ospfv2Topology {
        ospfv2Link.EntityData.Children.Append(types.GetSegmentPath(ospfv2Link.Ospfv2Topology[i]), types.YChild{"Ospfv2Topology", ospfv2Link.Ospfv2Topology[i]})
    }
    ospfv2Link.EntityData.Leafs = types.NewOrderedMap()
    ospfv2Link.EntityData.Leafs.Append("link-id", types.YLeaf{"LinkId", ospfv2Link.LinkId})
    ospfv2Link.EntityData.Leafs.Append("link-data", types.YLeaf{"LinkData", ospfv2Link.LinkData})
    ospfv2Link.EntityData.Leafs.Append("type", types.YLeaf{"Type", ospfv2Link.Type})

    ospfv2Link.EntityData.YListKeys = []string {"LinkId", "LinkData"}

    return &(ospfv2Link.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2Link_Ospfv2Topology
// Topology specific information
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2Link_Ospfv2Topology struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. MT-ID for topology enabled link. The type is
    // interface{} with range: 0..4294967295.
    MtId interface{}

    // Metric for the topology. The type is interface{} with range: 0..65535.
    Metric interface{}
}

func (ospfv2Topology *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2Link_Ospfv2Topology) GetEntityData() *types.CommonEntityData {
    ospfv2Topology.EntityData.YFilter = ospfv2Topology.YFilter
    ospfv2Topology.EntityData.YangName = "ospfv2-topology"
    ospfv2Topology.EntityData.BundleName = "cisco_ios_xe"
    ospfv2Topology.EntityData.ParentYangName = "ospfv2-link"
    ospfv2Topology.EntityData.SegmentPath = "ospfv2-topology" + types.AddKeyToken(ospfv2Topology.MtId, "mt-id")
    ospfv2Topology.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/link-scope-lsas/link-scope-lsa/ospfv2-link/" + ospfv2Topology.EntityData.SegmentPath
    ospfv2Topology.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv2Topology.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv2Topology.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv2Topology.EntityData.Children = types.NewOrderedMap()
    ospfv2Topology.EntityData.Leafs = types.NewOrderedMap()
    ospfv2Topology.EntityData.Leafs.Append("mt-id", types.YLeaf{"MtId", ospfv2Topology.MtId})
    ospfv2Topology.EntityData.Leafs.Append("metric", types.YLeaf{"Metric", ospfv2Topology.Metric})

    ospfv2Topology.EntityData.YListKeys = []string {"MtId"}

    return &(ospfv2Topology.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2Topology
// Summary LSA
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2Topology struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. MT-ID for topology enabled link. The type is
    // interface{} with range: 0..4294967295.
    MtId interface{}

    // Metric for the topology. The type is interface{} with range: 0..65535.
    Metric interface{}
}

func (ospfv2Topology *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2Topology) GetEntityData() *types.CommonEntityData {
    ospfv2Topology.EntityData.YFilter = ospfv2Topology.YFilter
    ospfv2Topology.EntityData.YangName = "ospfv2-topology"
    ospfv2Topology.EntityData.BundleName = "cisco_ios_xe"
    ospfv2Topology.EntityData.ParentYangName = "link-scope-lsa"
    ospfv2Topology.EntityData.SegmentPath = "ospfv2-topology" + types.AddKeyToken(ospfv2Topology.MtId, "mt-id")
    ospfv2Topology.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/link-scope-lsas/link-scope-lsa/" + ospfv2Topology.EntityData.SegmentPath
    ospfv2Topology.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv2Topology.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv2Topology.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv2Topology.EntityData.Children = types.NewOrderedMap()
    ospfv2Topology.EntityData.Leafs = types.NewOrderedMap()
    ospfv2Topology.EntityData.Leafs.Append("mt-id", types.YLeaf{"MtId", ospfv2Topology.MtId})
    ospfv2Topology.EntityData.Leafs.Append("metric", types.YLeaf{"Metric", ospfv2Topology.Metric})

    ospfv2Topology.EntityData.YListKeys = []string {"MtId"}

    return &(ospfv2Topology.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2External
// External LSA
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2External struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. MT-ID for topology enabled on the link. The type
    // is interface{} with range: 0..4294967295.
    MtId interface{}

    // Metric for the topology. The type is interface{} with range: 0..4294967295.
    Metric interface{}

    // Forwarding address. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ForwardingAddress interface{}

    // Route tag. The type is interface{} with range: 0..4294967295.
    ExternalRouteTag interface{}
}

func (ospfv2External *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2External) GetEntityData() *types.CommonEntityData {
    ospfv2External.EntityData.YFilter = ospfv2External.YFilter
    ospfv2External.EntityData.YangName = "ospfv2-external"
    ospfv2External.EntityData.BundleName = "cisco_ios_xe"
    ospfv2External.EntityData.ParentYangName = "link-scope-lsa"
    ospfv2External.EntityData.SegmentPath = "ospfv2-external" + types.AddKeyToken(ospfv2External.MtId, "mt-id")
    ospfv2External.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/link-scope-lsas/link-scope-lsa/" + ospfv2External.EntityData.SegmentPath
    ospfv2External.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv2External.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv2External.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv2External.EntityData.Children = types.NewOrderedMap()
    ospfv2External.EntityData.Leafs = types.NewOrderedMap()
    ospfv2External.EntityData.Leafs.Append("mt-id", types.YLeaf{"MtId", ospfv2External.MtId})
    ospfv2External.EntityData.Leafs.Append("metric", types.YLeaf{"Metric", ospfv2External.Metric})
    ospfv2External.EntityData.Leafs.Append("forwarding-address", types.YLeaf{"ForwardingAddress", ospfv2External.ForwardingAddress})
    ospfv2External.EntityData.Leafs.Append("external-route-tag", types.YLeaf{"ExternalRouteTag", ospfv2External.ExternalRouteTag})

    ospfv2External.EntityData.YListKeys = []string {"MtId"}

    return &(ospfv2External.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2UnknownTlv
// OSPFv2 Unknown TLV
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2UnknownTlv struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. TLV type. The type is interface{} with range:
    // 0..65535.
    Type interface{}

    // TLV length. The type is interface{} with range: 0..65535.
    Length interface{}

    // TLV value. The type is slice of interface{} with range: 0..255.
    Value []interface{}
}

func (ospfv2UnknownTlv *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2UnknownTlv) GetEntityData() *types.CommonEntityData {
    ospfv2UnknownTlv.EntityData.YFilter = ospfv2UnknownTlv.YFilter
    ospfv2UnknownTlv.EntityData.YangName = "ospfv2-unknown-tlv"
    ospfv2UnknownTlv.EntityData.BundleName = "cisco_ios_xe"
    ospfv2UnknownTlv.EntityData.ParentYangName = "link-scope-lsa"
    ospfv2UnknownTlv.EntityData.SegmentPath = "ospfv2-unknown-tlv" + types.AddKeyToken(ospfv2UnknownTlv.Type, "type")
    ospfv2UnknownTlv.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/link-scope-lsas/link-scope-lsa/" + ospfv2UnknownTlv.EntityData.SegmentPath
    ospfv2UnknownTlv.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv2UnknownTlv.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv2UnknownTlv.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv2UnknownTlv.EntityData.Children = types.NewOrderedMap()
    ospfv2UnknownTlv.EntityData.Leafs = types.NewOrderedMap()
    ospfv2UnknownTlv.EntityData.Leafs.Append("type", types.YLeaf{"Type", ospfv2UnknownTlv.Type})
    ospfv2UnknownTlv.EntityData.Leafs.Append("length", types.YLeaf{"Length", ospfv2UnknownTlv.Length})
    ospfv2UnknownTlv.EntityData.Leafs.Append("value", types.YLeaf{"Value", ospfv2UnknownTlv.Value})

    ospfv2UnknownTlv.EntityData.YListKeys = []string {"Type"}

    return &(ospfv2UnknownTlv.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal
// OSPFv3 LSA
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Decoded OSPFv3 LSA header.
    Header OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_Header

    // Decoded OSPFv3 LSA body.
    LsaBody OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody
}

func (ospfv3LsaVal *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal) GetEntityData() *types.CommonEntityData {
    ospfv3LsaVal.EntityData.YFilter = ospfv3LsaVal.YFilter
    ospfv3LsaVal.EntityData.YangName = "ospfv3-lsa-val"
    ospfv3LsaVal.EntityData.BundleName = "cisco_ios_xe"
    ospfv3LsaVal.EntityData.ParentYangName = "link-scope-lsa"
    ospfv3LsaVal.EntityData.SegmentPath = "ospfv3-lsa-val"
    ospfv3LsaVal.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/link-scope-lsas/link-scope-lsa/" + ospfv3LsaVal.EntityData.SegmentPath
    ospfv3LsaVal.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv3LsaVal.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv3LsaVal.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv3LsaVal.EntityData.Children = types.NewOrderedMap()
    ospfv3LsaVal.EntityData.Children.Append("header", types.YChild{"Header", &ospfv3LsaVal.Header})
    ospfv3LsaVal.EntityData.Children.Append("lsa-body", types.YChild{"LsaBody", &ospfv3LsaVal.LsaBody})
    ospfv3LsaVal.EntityData.Leafs = types.NewOrderedMap()

    ospfv3LsaVal.EntityData.YListKeys = []string {}

    return &(ospfv3LsaVal.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_Header
// Decoded OSPFv3 LSA header
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_Header struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA ID. The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    LsaId interface{}

    // OSPFv3 LSA options. The type is map[string]bool.
    LsaHdrOptions interface{}

    // LSA header.
    LsaHeader OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_Header_LsaHeader
}

func (header *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_Header) GetEntityData() *types.CommonEntityData {
    header.EntityData.YFilter = header.YFilter
    header.EntityData.YangName = "header"
    header.EntityData.BundleName = "cisco_ios_xe"
    header.EntityData.ParentYangName = "ospfv3-lsa-val"
    header.EntityData.SegmentPath = "header"
    header.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/link-scope-lsas/link-scope-lsa/ospfv3-lsa-val/" + header.EntityData.SegmentPath
    header.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    header.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    header.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    header.EntityData.Children = types.NewOrderedMap()
    header.EntityData.Children.Append("lsa-header", types.YChild{"LsaHeader", &header.LsaHeader})
    header.EntityData.Leafs = types.NewOrderedMap()
    header.EntityData.Leafs.Append("lsa-id", types.YLeaf{"LsaId", header.LsaId})
    header.EntityData.Leafs.Append("lsa-hdr-options", types.YLeaf{"LsaHdrOptions", header.LsaHdrOptions})

    header.EntityData.YListKeys = []string {}

    return &(header.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_Header_LsaHeader
// LSA header
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_Header_LsaHeader struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA age. The type is interface{} with range: 0..65535.
    Age interface{}

    // LSA type. The type is interface{} with range: 0..65535.
    Type interface{}

    // LSA advertising router. The type is interface{} with range: 0..4294967295.
    AdvRouter interface{}

    // LSA sequence number. The type is string.
    SeqNum interface{}

    // LSA checksum. The type is string.
    Checksum interface{}

    // LSA length. The type is interface{} with range: 0..65535.
    Length interface{}
}

func (lsaHeader *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_Header_LsaHeader) GetEntityData() *types.CommonEntityData {
    lsaHeader.EntityData.YFilter = lsaHeader.YFilter
    lsaHeader.EntityData.YangName = "lsa-header"
    lsaHeader.EntityData.BundleName = "cisco_ios_xe"
    lsaHeader.EntityData.ParentYangName = "header"
    lsaHeader.EntityData.SegmentPath = "lsa-header"
    lsaHeader.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/link-scope-lsas/link-scope-lsa/ospfv3-lsa-val/header/" + lsaHeader.EntityData.SegmentPath
    lsaHeader.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lsaHeader.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lsaHeader.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lsaHeader.EntityData.Children = types.NewOrderedMap()
    lsaHeader.EntityData.Leafs = types.NewOrderedMap()
    lsaHeader.EntityData.Leafs.Append("age", types.YLeaf{"Age", lsaHeader.Age})
    lsaHeader.EntityData.Leafs.Append("type", types.YLeaf{"Type", lsaHeader.Type})
    lsaHeader.EntityData.Leafs.Append("adv-router", types.YLeaf{"AdvRouter", lsaHeader.AdvRouter})
    lsaHeader.EntityData.Leafs.Append("seq-num", types.YLeaf{"SeqNum", lsaHeader.SeqNum})
    lsaHeader.EntityData.Leafs.Append("checksum", types.YLeaf{"Checksum", lsaHeader.Checksum})
    lsaHeader.EntityData.Leafs.Append("length", types.YLeaf{"Length", lsaHeader.Length})

    lsaHeader.EntityData.YListKeys = []string {}

    return &(lsaHeader.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody
// Decoded OSPFv3 LSA body
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA options. The type is map[string]bool.
    LsaFlagOptions interface{}

    // LSA Body Flags. The type is map[string]bool.
    LsaBodyFlags interface{}

    // OSPFv3 network.
    Network OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Network

    // OSPFv3 inter area prefix.
    Prefix OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Prefix

    // OSPFv3 inter area router.
    IaRouter OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_IaRouter

    // OSPFv3 LSA external.
    LsaExternal OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_LsaExternal

    // OSPFv3 NSSA.
    Nssa OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Nssa

    // OSPFv3 Link data.
    LinkData OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_LinkData

    // OSPFv3 Intra area prefixes.
    IaPrefix OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_IaPrefix
}

func (lsaBody *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody) GetEntityData() *types.CommonEntityData {
    lsaBody.EntityData.YFilter = lsaBody.YFilter
    lsaBody.EntityData.YangName = "lsa-body"
    lsaBody.EntityData.BundleName = "cisco_ios_xe"
    lsaBody.EntityData.ParentYangName = "ospfv3-lsa-val"
    lsaBody.EntityData.SegmentPath = "lsa-body"
    lsaBody.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/link-scope-lsas/link-scope-lsa/ospfv3-lsa-val/" + lsaBody.EntityData.SegmentPath
    lsaBody.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lsaBody.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lsaBody.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lsaBody.EntityData.Children = types.NewOrderedMap()
    lsaBody.EntityData.Children.Append("network", types.YChild{"Network", &lsaBody.Network})
    lsaBody.EntityData.Children.Append("prefix", types.YChild{"Prefix", &lsaBody.Prefix})
    lsaBody.EntityData.Children.Append("ia-router", types.YChild{"IaRouter", &lsaBody.IaRouter})
    lsaBody.EntityData.Children.Append("lsa-external", types.YChild{"LsaExternal", &lsaBody.LsaExternal})
    lsaBody.EntityData.Children.Append("nssa", types.YChild{"Nssa", &lsaBody.Nssa})
    lsaBody.EntityData.Children.Append("link-data", types.YChild{"LinkData", &lsaBody.LinkData})
    lsaBody.EntityData.Children.Append("ia-prefix", types.YChild{"IaPrefix", &lsaBody.IaPrefix})
    lsaBody.EntityData.Leafs = types.NewOrderedMap()
    lsaBody.EntityData.Leafs.Append("lsa-flag-options", types.YLeaf{"LsaFlagOptions", lsaBody.LsaFlagOptions})
    lsaBody.EntityData.Leafs.Append("lsa-body-flags", types.YLeaf{"LsaBodyFlags", lsaBody.LsaBodyFlags})

    lsaBody.EntityData.YListKeys = []string {}

    return &(lsaBody.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Network
// OSPFv3 network
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Network struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of the routers attached to the network. The type is slice of
    // interface{} with range: 0..4294967295.
    AttachedRouter []interface{}

    // Network LSA options. The type is map[string]bool.
    LsaNetOptions interface{}
}

func (network *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Network) GetEntityData() *types.CommonEntityData {
    network.EntityData.YFilter = network.YFilter
    network.EntityData.YangName = "network"
    network.EntityData.BundleName = "cisco_ios_xe"
    network.EntityData.ParentYangName = "lsa-body"
    network.EntityData.SegmentPath = "network"
    network.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/link-scope-lsas/link-scope-lsa/ospfv3-lsa-val/lsa-body/" + network.EntityData.SegmentPath
    network.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    network.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    network.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    network.EntityData.Children = types.NewOrderedMap()
    network.EntityData.Leafs = types.NewOrderedMap()
    network.EntityData.Leafs.Append("attached-router", types.YLeaf{"AttachedRouter", network.AttachedRouter})
    network.EntityData.Leafs.Append("lsa-net-options", types.YLeaf{"LsaNetOptions", network.LsaNetOptions})

    network.EntityData.YListKeys = []string {}

    return &(network.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Prefix
// OSPFv3 inter area prefix
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Prefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Metric. The type is interface{} with range: 0..4294967295.
    Metric interface{}

    // Inter area Prefix. The type is string.
    IaPrefix interface{}

    // Inter area prefix options. The type is string.
    IaPrefixOptions interface{}
}

func (prefix *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Prefix) GetEntityData() *types.CommonEntityData {
    prefix.EntityData.YFilter = prefix.YFilter
    prefix.EntityData.YangName = "prefix"
    prefix.EntityData.BundleName = "cisco_ios_xe"
    prefix.EntityData.ParentYangName = "lsa-body"
    prefix.EntityData.SegmentPath = "prefix"
    prefix.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/link-scope-lsas/link-scope-lsa/ospfv3-lsa-val/lsa-body/" + prefix.EntityData.SegmentPath
    prefix.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    prefix.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    prefix.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    prefix.EntityData.Children = types.NewOrderedMap()
    prefix.EntityData.Leafs = types.NewOrderedMap()
    prefix.EntityData.Leafs.Append("metric", types.YLeaf{"Metric", prefix.Metric})
    prefix.EntityData.Leafs.Append("ia-prefix", types.YLeaf{"IaPrefix", prefix.IaPrefix})
    prefix.EntityData.Leafs.Append("ia-prefix-options", types.YLeaf{"IaPrefixOptions", prefix.IaPrefixOptions})

    prefix.EntityData.YListKeys = []string {}

    return &(prefix.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_IaRouter
// OSPFv3 inter area router
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_IaRouter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Metric. The type is interface{} with range: 0..4294967295.
    Metric interface{}

    // Router ID of the router being described by the LSA. The type is interface{}
    // with range: 0..4294967295.
    DestinationRouterId interface{}

    // Inter area LSA options. The type is map[string]bool.
    LsaIaOptions interface{}
}

func (iaRouter *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_IaRouter) GetEntityData() *types.CommonEntityData {
    iaRouter.EntityData.YFilter = iaRouter.YFilter
    iaRouter.EntityData.YangName = "ia-router"
    iaRouter.EntityData.BundleName = "cisco_ios_xe"
    iaRouter.EntityData.ParentYangName = "lsa-body"
    iaRouter.EntityData.SegmentPath = "ia-router"
    iaRouter.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/link-scope-lsas/link-scope-lsa/ospfv3-lsa-val/lsa-body/" + iaRouter.EntityData.SegmentPath
    iaRouter.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    iaRouter.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    iaRouter.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    iaRouter.EntityData.Children = types.NewOrderedMap()
    iaRouter.EntityData.Leafs = types.NewOrderedMap()
    iaRouter.EntityData.Leafs.Append("metric", types.YLeaf{"Metric", iaRouter.Metric})
    iaRouter.EntityData.Leafs.Append("destination-router-id", types.YLeaf{"DestinationRouterId", iaRouter.DestinationRouterId})
    iaRouter.EntityData.Leafs.Append("lsa-ia-options", types.YLeaf{"LsaIaOptions", iaRouter.LsaIaOptions})

    iaRouter.EntityData.YListKeys = []string {}

    return &(iaRouter.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_LsaExternal
// OSPFv3 LSA external
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_LsaExternal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Metric. The type is interface{} with range: 0..4294967295.
    Metric interface{}

    // Referenced Link State type. The type is interface{} with range: 0..65535.
    ReferencedLsType interface{}

    // Prefix. The type is string.
    ExternalPrefix interface{}

    // Prefix options. The type is string.
    ExternalPrefixOptions interface{}

    // Forwarding address. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ForwardingAddress interface{}

    // Route tag. The type is interface{} with range: 0..4294967295.
    ExternalRouteTag interface{}

    // Referenced Link State ID. The type is interface{} with range:
    // 0..4294967295.
    ReferencedLinkStateId interface{}

    // LSA Flags.
    Flags OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_LsaExternal_Flags
}

func (lsaExternal *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_LsaExternal) GetEntityData() *types.CommonEntityData {
    lsaExternal.EntityData.YFilter = lsaExternal.YFilter
    lsaExternal.EntityData.YangName = "lsa-external"
    lsaExternal.EntityData.BundleName = "cisco_ios_xe"
    lsaExternal.EntityData.ParentYangName = "lsa-body"
    lsaExternal.EntityData.SegmentPath = "lsa-external"
    lsaExternal.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/link-scope-lsas/link-scope-lsa/ospfv3-lsa-val/lsa-body/" + lsaExternal.EntityData.SegmentPath
    lsaExternal.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lsaExternal.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lsaExternal.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lsaExternal.EntityData.Children = types.NewOrderedMap()
    lsaExternal.EntityData.Children.Append("flags", types.YChild{"Flags", &lsaExternal.Flags})
    lsaExternal.EntityData.Leafs = types.NewOrderedMap()
    lsaExternal.EntityData.Leafs.Append("metric", types.YLeaf{"Metric", lsaExternal.Metric})
    lsaExternal.EntityData.Leafs.Append("referenced-ls-type", types.YLeaf{"ReferencedLsType", lsaExternal.ReferencedLsType})
    lsaExternal.EntityData.Leafs.Append("external-prefix", types.YLeaf{"ExternalPrefix", lsaExternal.ExternalPrefix})
    lsaExternal.EntityData.Leafs.Append("external-prefix-options", types.YLeaf{"ExternalPrefixOptions", lsaExternal.ExternalPrefixOptions})
    lsaExternal.EntityData.Leafs.Append("forwarding-address", types.YLeaf{"ForwardingAddress", lsaExternal.ForwardingAddress})
    lsaExternal.EntityData.Leafs.Append("external-route-tag", types.YLeaf{"ExternalRouteTag", lsaExternal.ExternalRouteTag})
    lsaExternal.EntityData.Leafs.Append("referenced-link-state-id", types.YLeaf{"ReferencedLinkStateId", lsaExternal.ReferencedLinkStateId})

    lsaExternal.EntityData.YListKeys = []string {}

    return &(lsaExternal.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_LsaExternal_Flags
// LSA Flags
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_LsaExternal_Flags struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // When set, the metric specified is a Type 2 external metric. The type is
    // bool.
    EFlag interface{}
}

func (flags *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_LsaExternal_Flags) GetEntityData() *types.CommonEntityData {
    flags.EntityData.YFilter = flags.YFilter
    flags.EntityData.YangName = "flags"
    flags.EntityData.BundleName = "cisco_ios_xe"
    flags.EntityData.ParentYangName = "lsa-external"
    flags.EntityData.SegmentPath = "flags"
    flags.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/link-scope-lsas/link-scope-lsa/ospfv3-lsa-val/lsa-body/lsa-external/" + flags.EntityData.SegmentPath
    flags.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    flags.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    flags.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    flags.EntityData.Children = types.NewOrderedMap()
    flags.EntityData.Leafs = types.NewOrderedMap()
    flags.EntityData.Leafs.Append("e-flag", types.YLeaf{"EFlag", flags.EFlag})

    flags.EntityData.YListKeys = []string {}

    return &(flags.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Nssa
// OSPFv3 NSSA
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Nssa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // NSSA LSA.
    LsaNssaExternal OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Nssa_LsaNssaExternal
}

func (nssa *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Nssa) GetEntityData() *types.CommonEntityData {
    nssa.EntityData.YFilter = nssa.YFilter
    nssa.EntityData.YangName = "nssa"
    nssa.EntityData.BundleName = "cisco_ios_xe"
    nssa.EntityData.ParentYangName = "lsa-body"
    nssa.EntityData.SegmentPath = "nssa"
    nssa.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/link-scope-lsas/link-scope-lsa/ospfv3-lsa-val/lsa-body/" + nssa.EntityData.SegmentPath
    nssa.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nssa.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nssa.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nssa.EntityData.Children = types.NewOrderedMap()
    nssa.EntityData.Children.Append("lsa-nssa-external", types.YChild{"LsaNssaExternal", &nssa.LsaNssaExternal})
    nssa.EntityData.Leafs = types.NewOrderedMap()

    nssa.EntityData.YListKeys = []string {}

    return &(nssa.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Nssa_LsaNssaExternal
// NSSA LSA
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Nssa_LsaNssaExternal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Metric. The type is interface{} with range: 0..4294967295.
    Metric interface{}

    // Referenced Link State type. The type is interface{} with range: 0..65535.
    ReferencedLsType interface{}

    // Prefix. The type is string.
    ExternalPrefix interface{}

    // Prefix options. The type is string.
    ExternalPrefixOptions interface{}

    // Forwarding address. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ForwardingAddress interface{}

    // Route tag. The type is interface{} with range: 0..4294967295.
    ExternalRouteTag interface{}

    // Referenced Link State ID. The type is interface{} with range:
    // 0..4294967295.
    ReferencedLinkStateId interface{}

    // LSA Flags.
    Flags OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Nssa_LsaNssaExternal_Flags
}

func (lsaNssaExternal *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Nssa_LsaNssaExternal) GetEntityData() *types.CommonEntityData {
    lsaNssaExternal.EntityData.YFilter = lsaNssaExternal.YFilter
    lsaNssaExternal.EntityData.YangName = "lsa-nssa-external"
    lsaNssaExternal.EntityData.BundleName = "cisco_ios_xe"
    lsaNssaExternal.EntityData.ParentYangName = "nssa"
    lsaNssaExternal.EntityData.SegmentPath = "lsa-nssa-external"
    lsaNssaExternal.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/link-scope-lsas/link-scope-lsa/ospfv3-lsa-val/lsa-body/nssa/" + lsaNssaExternal.EntityData.SegmentPath
    lsaNssaExternal.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lsaNssaExternal.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lsaNssaExternal.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lsaNssaExternal.EntityData.Children = types.NewOrderedMap()
    lsaNssaExternal.EntityData.Children.Append("flags", types.YChild{"Flags", &lsaNssaExternal.Flags})
    lsaNssaExternal.EntityData.Leafs = types.NewOrderedMap()
    lsaNssaExternal.EntityData.Leafs.Append("metric", types.YLeaf{"Metric", lsaNssaExternal.Metric})
    lsaNssaExternal.EntityData.Leafs.Append("referenced-ls-type", types.YLeaf{"ReferencedLsType", lsaNssaExternal.ReferencedLsType})
    lsaNssaExternal.EntityData.Leafs.Append("external-prefix", types.YLeaf{"ExternalPrefix", lsaNssaExternal.ExternalPrefix})
    lsaNssaExternal.EntityData.Leafs.Append("external-prefix-options", types.YLeaf{"ExternalPrefixOptions", lsaNssaExternal.ExternalPrefixOptions})
    lsaNssaExternal.EntityData.Leafs.Append("forwarding-address", types.YLeaf{"ForwardingAddress", lsaNssaExternal.ForwardingAddress})
    lsaNssaExternal.EntityData.Leafs.Append("external-route-tag", types.YLeaf{"ExternalRouteTag", lsaNssaExternal.ExternalRouteTag})
    lsaNssaExternal.EntityData.Leafs.Append("referenced-link-state-id", types.YLeaf{"ReferencedLinkStateId", lsaNssaExternal.ReferencedLinkStateId})

    lsaNssaExternal.EntityData.YListKeys = []string {}

    return &(lsaNssaExternal.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Nssa_LsaNssaExternal_Flags
// LSA Flags
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Nssa_LsaNssaExternal_Flags struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // When set, the metric specified is a Type 2 external metric. The type is
    // bool.
    EFlag interface{}
}

func (flags *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Nssa_LsaNssaExternal_Flags) GetEntityData() *types.CommonEntityData {
    flags.EntityData.YFilter = flags.YFilter
    flags.EntityData.YangName = "flags"
    flags.EntityData.BundleName = "cisco_ios_xe"
    flags.EntityData.ParentYangName = "lsa-nssa-external"
    flags.EntityData.SegmentPath = "flags"
    flags.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/link-scope-lsas/link-scope-lsa/ospfv3-lsa-val/lsa-body/nssa/lsa-nssa-external/" + flags.EntityData.SegmentPath
    flags.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    flags.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    flags.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    flags.EntityData.Children = types.NewOrderedMap()
    flags.EntityData.Leafs = types.NewOrderedMap()
    flags.EntityData.Leafs.Append("e-flag", types.YLeaf{"EFlag", flags.EFlag})

    flags.EntityData.YListKeys = []string {}

    return &(flags.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_LinkData
// OSPFv3 Link data
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_LinkData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Router priority of the interce. The type is interface{} with range: 0..255.
    RtrPriority interface{}

    // The originating router's link-local interface address on the link. The type
    // is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    LinkLocalInterfaceAddress interface{}

    // Number of prefixes. The type is interface{} with range: 0..4294967295.
    NumOfPrefixes interface{}

    // Link data LSA options. The type is map[string]bool.
    LsaIdOptions interface{}
}

func (linkData *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_LinkData) GetEntityData() *types.CommonEntityData {
    linkData.EntityData.YFilter = linkData.YFilter
    linkData.EntityData.YangName = "link-data"
    linkData.EntityData.BundleName = "cisco_ios_xe"
    linkData.EntityData.ParentYangName = "lsa-body"
    linkData.EntityData.SegmentPath = "link-data"
    linkData.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/link-scope-lsas/link-scope-lsa/ospfv3-lsa-val/lsa-body/" + linkData.EntityData.SegmentPath
    linkData.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    linkData.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    linkData.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    linkData.EntityData.Children = types.NewOrderedMap()
    linkData.EntityData.Leafs = types.NewOrderedMap()
    linkData.EntityData.Leafs.Append("rtr-priority", types.YLeaf{"RtrPriority", linkData.RtrPriority})
    linkData.EntityData.Leafs.Append("link-local-interface-address", types.YLeaf{"LinkLocalInterfaceAddress", linkData.LinkLocalInterfaceAddress})
    linkData.EntityData.Leafs.Append("num-of-prefixes", types.YLeaf{"NumOfPrefixes", linkData.NumOfPrefixes})
    linkData.EntityData.Leafs.Append("lsa-id-options", types.YLeaf{"LsaIdOptions", linkData.LsaIdOptions})

    linkData.EntityData.YListKeys = []string {}

    return &(linkData.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_IaPrefix
// OSPFv3 Intra area prefixes
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_IaPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Referenced Link State type. The type is interface{} with range: 0..65535.
    ReferencedLsType interface{}

    // Referenced Link State ID. The type is interface{} with range:
    // 0..4294967295.
    ReferencedLinkStateId interface{}

    // Referenced Advertising Router. The type is one of the following types:
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ReferencedAdvRouter interface{}

    // Number of prefixes. The type is interface{} with range: 0..65535.
    NumOfPrefixes interface{}
}

func (iaPrefix *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_IaPrefix) GetEntityData() *types.CommonEntityData {
    iaPrefix.EntityData.YFilter = iaPrefix.YFilter
    iaPrefix.EntityData.YangName = "ia-prefix"
    iaPrefix.EntityData.BundleName = "cisco_ios_xe"
    iaPrefix.EntityData.ParentYangName = "lsa-body"
    iaPrefix.EntityData.SegmentPath = "ia-prefix"
    iaPrefix.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/link-scope-lsas/link-scope-lsa/ospfv3-lsa-val/lsa-body/" + iaPrefix.EntityData.SegmentPath
    iaPrefix.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    iaPrefix.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    iaPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    iaPrefix.EntityData.Children = types.NewOrderedMap()
    iaPrefix.EntityData.Leafs = types.NewOrderedMap()
    iaPrefix.EntityData.Leafs.Append("referenced-ls-type", types.YLeaf{"ReferencedLsType", iaPrefix.ReferencedLsType})
    iaPrefix.EntityData.Leafs.Append("referenced-link-state-id", types.YLeaf{"ReferencedLinkStateId", iaPrefix.ReferencedLinkStateId})
    iaPrefix.EntityData.Leafs.Append("referenced-adv-router", types.YLeaf{"ReferencedAdvRouter", iaPrefix.ReferencedAdvRouter})
    iaPrefix.EntityData.Leafs.Append("num-of-prefixes", types.YLeaf{"NumOfPrefixes", iaPrefix.NumOfPrefixes})

    iaPrefix.EntityData.YListKeys = []string {}

    return &(iaPrefix.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3Link
// OSPFv3 links
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3Link struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface ID. The type is interface{} with range:
    // 0..4294967295.
    InterfaceId interface{}

    // This attribute is a key. Neighbor interface ID. The type is interface{}
    // with range: 0..4294967295.
    NeighborInterfaceId interface{}

    // This attribute is a key. Neighbor router ID. The type is interface{} with
    // range: 0..4294967295.
    NeighborRouterId interface{}

    // Link type. The type is interface{} with range: 0..255.
    Type interface{}

    // Metric. The type is interface{} with range: 0..65535.
    Metric interface{}
}

func (ospfv3Link *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3Link) GetEntityData() *types.CommonEntityData {
    ospfv3Link.EntityData.YFilter = ospfv3Link.YFilter
    ospfv3Link.EntityData.YangName = "ospfv3-link"
    ospfv3Link.EntityData.BundleName = "cisco_ios_xe"
    ospfv3Link.EntityData.ParentYangName = "link-scope-lsa"
    ospfv3Link.EntityData.SegmentPath = "ospfv3-link" + types.AddKeyToken(ospfv3Link.InterfaceId, "interface-id") + types.AddKeyToken(ospfv3Link.NeighborInterfaceId, "neighbor-interface-id") + types.AddKeyToken(ospfv3Link.NeighborRouterId, "neighbor-router-id")
    ospfv3Link.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/link-scope-lsas/link-scope-lsa/" + ospfv3Link.EntityData.SegmentPath
    ospfv3Link.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv3Link.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv3Link.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv3Link.EntityData.Children = types.NewOrderedMap()
    ospfv3Link.EntityData.Leafs = types.NewOrderedMap()
    ospfv3Link.EntityData.Leafs.Append("interface-id", types.YLeaf{"InterfaceId", ospfv3Link.InterfaceId})
    ospfv3Link.EntityData.Leafs.Append("neighbor-interface-id", types.YLeaf{"NeighborInterfaceId", ospfv3Link.NeighborInterfaceId})
    ospfv3Link.EntityData.Leafs.Append("neighbor-router-id", types.YLeaf{"NeighborRouterId", ospfv3Link.NeighborRouterId})
    ospfv3Link.EntityData.Leafs.Append("type", types.YLeaf{"Type", ospfv3Link.Type})
    ospfv3Link.EntityData.Leafs.Append("metric", types.YLeaf{"Metric", ospfv3Link.Metric})

    ospfv3Link.EntityData.YListKeys = []string {"InterfaceId", "NeighborInterfaceId", "NeighborRouterId"}

    return &(ospfv3Link.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3PrefixList
// OSPFv3 prefix-list
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3PrefixList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Prefix. The type is string.
    Prefix interface{}

    // Prefix options. The type is string.
    PrefixOptions interface{}
}

func (ospfv3PrefixList *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3PrefixList) GetEntityData() *types.CommonEntityData {
    ospfv3PrefixList.EntityData.YFilter = ospfv3PrefixList.YFilter
    ospfv3PrefixList.EntityData.YangName = "ospfv3-prefix-list"
    ospfv3PrefixList.EntityData.BundleName = "cisco_ios_xe"
    ospfv3PrefixList.EntityData.ParentYangName = "link-scope-lsa"
    ospfv3PrefixList.EntityData.SegmentPath = "ospfv3-prefix-list" + types.AddKeyToken(ospfv3PrefixList.Prefix, "prefix")
    ospfv3PrefixList.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/link-scope-lsas/link-scope-lsa/" + ospfv3PrefixList.EntityData.SegmentPath
    ospfv3PrefixList.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv3PrefixList.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv3PrefixList.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv3PrefixList.EntityData.Children = types.NewOrderedMap()
    ospfv3PrefixList.EntityData.Leafs = types.NewOrderedMap()
    ospfv3PrefixList.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", ospfv3PrefixList.Prefix})
    ospfv3PrefixList.EntityData.Leafs.Append("prefix-options", types.YLeaf{"PrefixOptions", ospfv3PrefixList.PrefixOptions})

    ospfv3PrefixList.EntityData.YListKeys = []string {"Prefix"}

    return &(ospfv3PrefixList.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3IaPrefix
// OSPFv3 intra-area prefix-list
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3IaPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Prefix. The type is string.
    Prefix interface{}

    // Prefix options. The type is string.
    PrefixOptions interface{}
}

func (ospfv3IaPrefix *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3IaPrefix) GetEntityData() *types.CommonEntityData {
    ospfv3IaPrefix.EntityData.YFilter = ospfv3IaPrefix.YFilter
    ospfv3IaPrefix.EntityData.YangName = "ospfv3-ia-prefix"
    ospfv3IaPrefix.EntityData.BundleName = "cisco_ios_xe"
    ospfv3IaPrefix.EntityData.ParentYangName = "link-scope-lsa"
    ospfv3IaPrefix.EntityData.SegmentPath = "ospfv3-ia-prefix" + types.AddKeyToken(ospfv3IaPrefix.Prefix, "prefix")
    ospfv3IaPrefix.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/link-scope-lsas/link-scope-lsa/" + ospfv3IaPrefix.EntityData.SegmentPath
    ospfv3IaPrefix.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv3IaPrefix.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv3IaPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv3IaPrefix.EntityData.Children = types.NewOrderedMap()
    ospfv3IaPrefix.EntityData.Leafs = types.NewOrderedMap()
    ospfv3IaPrefix.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", ospfv3IaPrefix.Prefix})
    ospfv3IaPrefix.EntityData.Leafs.Append("prefix-options", types.YLeaf{"PrefixOptions", ospfv3IaPrefix.PrefixOptions})

    ospfv3IaPrefix.EntityData.YListKeys = []string {"Prefix"}

    return &(ospfv3IaPrefix.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_MultiTopology
// OSPF multi-topology interface augmentation
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_MultiTopology struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. One of the topology enabled on this interface. The
    // type is string.
    Name interface{}
}

func (multiTopology *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_MultiTopology) GetEntityData() *types.CommonEntityData {
    multiTopology.EntityData.YFilter = multiTopology.YFilter
    multiTopology.EntityData.YangName = "multi-topology"
    multiTopology.EntityData.BundleName = "cisco_ios_xe"
    multiTopology.EntityData.ParentYangName = "link-scope-lsa"
    multiTopology.EntityData.SegmentPath = "multi-topology" + types.AddKeyToken(multiTopology.Name, "name")
    multiTopology.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/link-scope-lsas/link-scope-lsa/" + multiTopology.EntityData.SegmentPath
    multiTopology.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    multiTopology.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    multiTopology.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    multiTopology.EntityData.Children = types.NewOrderedMap()
    multiTopology.EntityData.Leafs = types.NewOrderedMap()
    multiTopology.EntityData.Leafs.Append("name", types.YLeaf{"Name", multiTopology.Name})

    multiTopology.EntityData.YListKeys = []string {"Name"}

    return &(multiTopology.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Tlv
// Link TLV
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Tlv struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Link type. The type is interface{} with range: 0..255.
    LinkType interface{}

    // Link ID. The type is interface{} with range: 0..4294967295.
    LinkId interface{}

    // List of local interface IPv4 addresses. The type is one of the following
    // types: slice of string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or slice of string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    LocalIfIpv4Addr []interface{}

    // List of remote interface IPv4 addresses. The type is one of the following
    // types: slice of string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or slice of string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    LocalRemoteIpv4Addr []interface{}

    // TE metric. The type is interface{} with range: 0..4294967295.
    TeMetric interface{}

    // Maximum bandwidth. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    MaxBandwidth interface{}

    // Maximum reservable bandwidth. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    MaxReservableBandwidth interface{}

    // Unrseerved bandwidth. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    UnreservedBandwidth interface{}

    // Administrative group/Resource class/Color. The type is interface{} with
    // range: 0..4294967295.
    AdminGroup interface{}
}

func (tlv *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Tlv) GetEntityData() *types.CommonEntityData {
    tlv.EntityData.YFilter = tlv.YFilter
    tlv.EntityData.YangName = "tlv"
    tlv.EntityData.BundleName = "cisco_ios_xe"
    tlv.EntityData.ParentYangName = "link-scope-lsa"
    tlv.EntityData.SegmentPath = "tlv"
    tlv.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/link-scope-lsas/link-scope-lsa/" + tlv.EntityData.SegmentPath
    tlv.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    tlv.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    tlv.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    tlv.EntityData.Children = types.NewOrderedMap()
    tlv.EntityData.Leafs = types.NewOrderedMap()
    tlv.EntityData.Leafs.Append("link-type", types.YLeaf{"LinkType", tlv.LinkType})
    tlv.EntityData.Leafs.Append("link-id", types.YLeaf{"LinkId", tlv.LinkId})
    tlv.EntityData.Leafs.Append("local-if-ipv4-addr", types.YLeaf{"LocalIfIpv4Addr", tlv.LocalIfIpv4Addr})
    tlv.EntityData.Leafs.Append("local-remote-ipv4-addr", types.YLeaf{"LocalRemoteIpv4Addr", tlv.LocalRemoteIpv4Addr})
    tlv.EntityData.Leafs.Append("te-metric", types.YLeaf{"TeMetric", tlv.TeMetric})
    tlv.EntityData.Leafs.Append("max-bandwidth", types.YLeaf{"MaxBandwidth", tlv.MaxBandwidth})
    tlv.EntityData.Leafs.Append("max-reservable-bandwidth", types.YLeaf{"MaxReservableBandwidth", tlv.MaxReservableBandwidth})
    tlv.EntityData.Leafs.Append("unreserved-bandwidth", types.YLeaf{"UnreservedBandwidth", tlv.UnreservedBandwidth})
    tlv.EntityData.Leafs.Append("admin-group", types.YLeaf{"AdminGroup", tlv.AdminGroup})

    tlv.EntityData.YListKeys = []string {}

    return &(tlv.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_UnknownSubTlv
// OSPFv2 Unknown sub TLV
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_UnknownSubTlv struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. TLV type. The type is interface{} with range:
    // 0..65535.
    Type interface{}

    // TLV length. The type is interface{} with range: 0..65535.
    Length interface{}

    // TLV value. The type is slice of interface{} with range: 0..255.
    Value []interface{}
}

func (unknownSubTlv *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_UnknownSubTlv) GetEntityData() *types.CommonEntityData {
    unknownSubTlv.EntityData.YFilter = unknownSubTlv.YFilter
    unknownSubTlv.EntityData.YangName = "unknown-sub-tlv"
    unknownSubTlv.EntityData.BundleName = "cisco_ios_xe"
    unknownSubTlv.EntityData.ParentYangName = "link-scope-lsa"
    unknownSubTlv.EntityData.SegmentPath = "unknown-sub-tlv" + types.AddKeyToken(unknownSubTlv.Type, "type")
    unknownSubTlv.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/link-scope-lsas/link-scope-lsa/" + unknownSubTlv.EntityData.SegmentPath
    unknownSubTlv.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    unknownSubTlv.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    unknownSubTlv.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    unknownSubTlv.EntityData.Children = types.NewOrderedMap()
    unknownSubTlv.EntityData.Leafs = types.NewOrderedMap()
    unknownSubTlv.EntityData.Leafs.Append("type", types.YLeaf{"Type", unknownSubTlv.Type})
    unknownSubTlv.EntityData.Leafs.Append("length", types.YLeaf{"Length", unknownSubTlv.Length})
    unknownSubTlv.EntityData.Leafs.Append("value", types.YLeaf{"Value", unknownSubTlv.Value})

    unknownSubTlv.EntityData.YListKeys = []string {"Type"}

    return &(unknownSubTlv.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa
// List OSPF area scope LSA databases
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. LSA Type. The type is interface{} with range:
    // 0..4294967295.
    LsaType interface{}

    // This attribute is a key. Advertising router. The type is one of the
    // following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    AdvRouter interface{}

    // The OSPF LSA body is fully decoded. The type is bool.
    DecodedCompleted interface{}

    // The complete LSA in network byte order as received/sent over the wire. The
    // type is slice of interface{} with range: 0..255.
    RawData []interface{}

    // OSPFv2 LSA.
    Ospfv2Lsa OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2Lsa

    // Router LSA link. The type is slice of
    // OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2Link.
    Ospfv2Link []*OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2Link

    // Summary LSA. The type is slice of
    // OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2Topology.
    Ospfv2Topology []*OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2Topology

    // External LSA. The type is slice of
    // OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2External.
    Ospfv2External []*OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2External

    // OSPFv3 LSA.
    Ospfv3Lsa OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa

    // OSPFv3 links. The type is slice of
    // OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Link.
    Ospfv3Link []*OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Link

    // OSPFv3 prefix-list. The type is slice of
    // OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Prefix.
    Ospfv3Prefix []*OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Prefix

    // OSPFv3 intra-area prefix-list. The type is slice of
    // OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3IaPrefix.
    Ospfv3IaPrefix []*OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3IaPrefix
}

func (areaScopeLsa *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa) GetEntityData() *types.CommonEntityData {
    areaScopeLsa.EntityData.YFilter = areaScopeLsa.YFilter
    areaScopeLsa.EntityData.YangName = "area-scope-lsa"
    areaScopeLsa.EntityData.BundleName = "cisco_ios_xe"
    areaScopeLsa.EntityData.ParentYangName = "link-scope-lsas"
    areaScopeLsa.EntityData.SegmentPath = "area-scope-lsa" + types.AddKeyToken(areaScopeLsa.LsaType, "lsa-type") + types.AddKeyToken(areaScopeLsa.AdvRouter, "adv-router")
    areaScopeLsa.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/link-scope-lsas/" + areaScopeLsa.EntityData.SegmentPath
    areaScopeLsa.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    areaScopeLsa.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    areaScopeLsa.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    areaScopeLsa.EntityData.Children = types.NewOrderedMap()
    areaScopeLsa.EntityData.Children.Append("ospfv2-lsa", types.YChild{"Ospfv2Lsa", &areaScopeLsa.Ospfv2Lsa})
    areaScopeLsa.EntityData.Children.Append("ospfv2-link", types.YChild{"Ospfv2Link", nil})
    for i := range areaScopeLsa.Ospfv2Link {
        areaScopeLsa.EntityData.Children.Append(types.GetSegmentPath(areaScopeLsa.Ospfv2Link[i]), types.YChild{"Ospfv2Link", areaScopeLsa.Ospfv2Link[i]})
    }
    areaScopeLsa.EntityData.Children.Append("ospfv2-topology", types.YChild{"Ospfv2Topology", nil})
    for i := range areaScopeLsa.Ospfv2Topology {
        areaScopeLsa.EntityData.Children.Append(types.GetSegmentPath(areaScopeLsa.Ospfv2Topology[i]), types.YChild{"Ospfv2Topology", areaScopeLsa.Ospfv2Topology[i]})
    }
    areaScopeLsa.EntityData.Children.Append("ospfv2-external", types.YChild{"Ospfv2External", nil})
    for i := range areaScopeLsa.Ospfv2External {
        areaScopeLsa.EntityData.Children.Append(types.GetSegmentPath(areaScopeLsa.Ospfv2External[i]), types.YChild{"Ospfv2External", areaScopeLsa.Ospfv2External[i]})
    }
    areaScopeLsa.EntityData.Children.Append("ospfv3-lsa", types.YChild{"Ospfv3Lsa", &areaScopeLsa.Ospfv3Lsa})
    areaScopeLsa.EntityData.Children.Append("ospfv3-link", types.YChild{"Ospfv3Link", nil})
    for i := range areaScopeLsa.Ospfv3Link {
        areaScopeLsa.EntityData.Children.Append(types.GetSegmentPath(areaScopeLsa.Ospfv3Link[i]), types.YChild{"Ospfv3Link", areaScopeLsa.Ospfv3Link[i]})
    }
    areaScopeLsa.EntityData.Children.Append("ospfv3-prefix", types.YChild{"Ospfv3Prefix", nil})
    for i := range areaScopeLsa.Ospfv3Prefix {
        areaScopeLsa.EntityData.Children.Append(types.GetSegmentPath(areaScopeLsa.Ospfv3Prefix[i]), types.YChild{"Ospfv3Prefix", areaScopeLsa.Ospfv3Prefix[i]})
    }
    areaScopeLsa.EntityData.Children.Append("ospfv3-ia-prefix", types.YChild{"Ospfv3IaPrefix", nil})
    for i := range areaScopeLsa.Ospfv3IaPrefix {
        areaScopeLsa.EntityData.Children.Append(types.GetSegmentPath(areaScopeLsa.Ospfv3IaPrefix[i]), types.YChild{"Ospfv3IaPrefix", areaScopeLsa.Ospfv3IaPrefix[i]})
    }
    areaScopeLsa.EntityData.Leafs = types.NewOrderedMap()
    areaScopeLsa.EntityData.Leafs.Append("lsa-type", types.YLeaf{"LsaType", areaScopeLsa.LsaType})
    areaScopeLsa.EntityData.Leafs.Append("adv-router", types.YLeaf{"AdvRouter", areaScopeLsa.AdvRouter})
    areaScopeLsa.EntityData.Leafs.Append("decoded-completed", types.YLeaf{"DecodedCompleted", areaScopeLsa.DecodedCompleted})
    areaScopeLsa.EntityData.Leafs.Append("raw-data", types.YLeaf{"RawData", areaScopeLsa.RawData})

    areaScopeLsa.EntityData.YListKeys = []string {"LsaType", "AdvRouter"}

    return &(areaScopeLsa.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2Lsa
// OSPFv2 LSA
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2Lsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Decoded OSPFv2 LSA header data.
    Header OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2Lsa_Header

    // Decoded OSPFv2 LSA body data.
    LsaBody OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2Lsa_LsaBody
}

func (ospfv2Lsa *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2Lsa) GetEntityData() *types.CommonEntityData {
    ospfv2Lsa.EntityData.YFilter = ospfv2Lsa.YFilter
    ospfv2Lsa.EntityData.YangName = "ospfv2-lsa"
    ospfv2Lsa.EntityData.BundleName = "cisco_ios_xe"
    ospfv2Lsa.EntityData.ParentYangName = "area-scope-lsa"
    ospfv2Lsa.EntityData.SegmentPath = "ospfv2-lsa"
    ospfv2Lsa.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/link-scope-lsas/area-scope-lsa/" + ospfv2Lsa.EntityData.SegmentPath
    ospfv2Lsa.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv2Lsa.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv2Lsa.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv2Lsa.EntityData.Children = types.NewOrderedMap()
    ospfv2Lsa.EntityData.Children.Append("header", types.YChild{"Header", &ospfv2Lsa.Header})
    ospfv2Lsa.EntityData.Children.Append("lsa-body", types.YChild{"LsaBody", &ospfv2Lsa.LsaBody})
    ospfv2Lsa.EntityData.Leafs = types.NewOrderedMap()

    ospfv2Lsa.EntityData.YListKeys = []string {}

    return &(ospfv2Lsa.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2Lsa_Header
// Decoded OSPFv2 LSA header data
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2Lsa_Header struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA ID. The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    LsaId interface{}

    // Opaque type. The type is interface{} with range: 0..255.
    OpaqueType interface{}

    // Opaque ID. The type is interface{} with range: 0..4294967295.
    OpaqueId interface{}

    // LSA age. The type is interface{} with range: 0..65535.
    Age interface{}

    // LSA type. The type is interface{} with range: 0..65535.
    Type interface{}

    // LSA advertising router. The type is interface{} with range: 0..4294967295.
    AdvRouter interface{}

    // LSA sequence number. The type is string.
    SeqNum interface{}

    // LSA checksum. The type is string.
    Checksum interface{}

    // LSA length. The type is interface{} with range: 0..65535.
    Length interface{}

    // LSA options. The type is map[string]bool.
    FlagOptions interface{}
}

func (header *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2Lsa_Header) GetEntityData() *types.CommonEntityData {
    header.EntityData.YFilter = header.YFilter
    header.EntityData.YangName = "header"
    header.EntityData.BundleName = "cisco_ios_xe"
    header.EntityData.ParentYangName = "ospfv2-lsa"
    header.EntityData.SegmentPath = "header"
    header.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/link-scope-lsas/area-scope-lsa/ospfv2-lsa/" + header.EntityData.SegmentPath
    header.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    header.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    header.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    header.EntityData.Children = types.NewOrderedMap()
    header.EntityData.Leafs = types.NewOrderedMap()
    header.EntityData.Leafs.Append("lsa-id", types.YLeaf{"LsaId", header.LsaId})
    header.EntityData.Leafs.Append("opaque-type", types.YLeaf{"OpaqueType", header.OpaqueType})
    header.EntityData.Leafs.Append("opaque-id", types.YLeaf{"OpaqueId", header.OpaqueId})
    header.EntityData.Leafs.Append("age", types.YLeaf{"Age", header.Age})
    header.EntityData.Leafs.Append("type", types.YLeaf{"Type", header.Type})
    header.EntityData.Leafs.Append("adv-router", types.YLeaf{"AdvRouter", header.AdvRouter})
    header.EntityData.Leafs.Append("seq-num", types.YLeaf{"SeqNum", header.SeqNum})
    header.EntityData.Leafs.Append("checksum", types.YLeaf{"Checksum", header.Checksum})
    header.EntityData.Leafs.Append("length", types.YLeaf{"Length", header.Length})
    header.EntityData.Leafs.Append("flag-options", types.YLeaf{"FlagOptions", header.FlagOptions})

    header.EntityData.YListKeys = []string {}

    return &(header.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2Lsa_LsaBody
// Decoded OSPFv2 LSA body data
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2Lsa_LsaBody struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of links. The type is interface{} with range: 0..65535.
    NumOfLinks interface{}

    // Summary mask. The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    SummaryMask interface{}

    // External mask. The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ExternalMask interface{}

    // LSA body flags. The type is map[string]bool.
    BodyFlagOptions interface{}

    // Network details.
    Network OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2Lsa_LsaBody_Network
}

func (lsaBody *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2Lsa_LsaBody) GetEntityData() *types.CommonEntityData {
    lsaBody.EntityData.YFilter = lsaBody.YFilter
    lsaBody.EntityData.YangName = "lsa-body"
    lsaBody.EntityData.BundleName = "cisco_ios_xe"
    lsaBody.EntityData.ParentYangName = "ospfv2-lsa"
    lsaBody.EntityData.SegmentPath = "lsa-body"
    lsaBody.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/link-scope-lsas/area-scope-lsa/ospfv2-lsa/" + lsaBody.EntityData.SegmentPath
    lsaBody.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lsaBody.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lsaBody.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lsaBody.EntityData.Children = types.NewOrderedMap()
    lsaBody.EntityData.Children.Append("network", types.YChild{"Network", &lsaBody.Network})
    lsaBody.EntityData.Leafs = types.NewOrderedMap()
    lsaBody.EntityData.Leafs.Append("num-of-links", types.YLeaf{"NumOfLinks", lsaBody.NumOfLinks})
    lsaBody.EntityData.Leafs.Append("summary-mask", types.YLeaf{"SummaryMask", lsaBody.SummaryMask})
    lsaBody.EntityData.Leafs.Append("external-mask", types.YLeaf{"ExternalMask", lsaBody.ExternalMask})
    lsaBody.EntityData.Leafs.Append("body-flag-options", types.YLeaf{"BodyFlagOptions", lsaBody.BodyFlagOptions})

    lsaBody.EntityData.YListKeys = []string {}

    return &(lsaBody.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2Lsa_LsaBody_Network
// Network details
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2Lsa_LsaBody_Network struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP network mask. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NetworkMask interface{}

    // List of the routers attached to the network. The type is slice of
    // interface{} with range: 0..4294967295.
    AttachedRouter []interface{}
}

func (network *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2Lsa_LsaBody_Network) GetEntityData() *types.CommonEntityData {
    network.EntityData.YFilter = network.YFilter
    network.EntityData.YangName = "network"
    network.EntityData.BundleName = "cisco_ios_xe"
    network.EntityData.ParentYangName = "lsa-body"
    network.EntityData.SegmentPath = "network"
    network.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/link-scope-lsas/area-scope-lsa/ospfv2-lsa/lsa-body/" + network.EntityData.SegmentPath
    network.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    network.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    network.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    network.EntityData.Children = types.NewOrderedMap()
    network.EntityData.Leafs = types.NewOrderedMap()
    network.EntityData.Leafs.Append("network-mask", types.YLeaf{"NetworkMask", network.NetworkMask})
    network.EntityData.Leafs.Append("attached-router", types.YLeaf{"AttachedRouter", network.AttachedRouter})

    network.EntityData.YListKeys = []string {}

    return &(network.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2Link
// Router LSA link
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2Link struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Link ID. The type is interface{} with range:
    // 0..4294967295.
    LinkId interface{}

    // This attribute is a key. Link data. The type is interface{} with range:
    // 0..4294967295.
    LinkData interface{}

    // Link type. The type is interface{} with range: 0..255.
    Type interface{}

    // Topology specific information. The type is slice of
    // OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2Link_Ospfv2Topology.
    Ospfv2Topology []*OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2Link_Ospfv2Topology
}

func (ospfv2Link *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2Link) GetEntityData() *types.CommonEntityData {
    ospfv2Link.EntityData.YFilter = ospfv2Link.YFilter
    ospfv2Link.EntityData.YangName = "ospfv2-link"
    ospfv2Link.EntityData.BundleName = "cisco_ios_xe"
    ospfv2Link.EntityData.ParentYangName = "area-scope-lsa"
    ospfv2Link.EntityData.SegmentPath = "ospfv2-link" + types.AddKeyToken(ospfv2Link.LinkId, "link-id") + types.AddKeyToken(ospfv2Link.LinkData, "link-data")
    ospfv2Link.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/link-scope-lsas/area-scope-lsa/" + ospfv2Link.EntityData.SegmentPath
    ospfv2Link.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv2Link.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv2Link.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv2Link.EntityData.Children = types.NewOrderedMap()
    ospfv2Link.EntityData.Children.Append("ospfv2-topology", types.YChild{"Ospfv2Topology", nil})
    for i := range ospfv2Link.Ospfv2Topology {
        ospfv2Link.EntityData.Children.Append(types.GetSegmentPath(ospfv2Link.Ospfv2Topology[i]), types.YChild{"Ospfv2Topology", ospfv2Link.Ospfv2Topology[i]})
    }
    ospfv2Link.EntityData.Leafs = types.NewOrderedMap()
    ospfv2Link.EntityData.Leafs.Append("link-id", types.YLeaf{"LinkId", ospfv2Link.LinkId})
    ospfv2Link.EntityData.Leafs.Append("link-data", types.YLeaf{"LinkData", ospfv2Link.LinkData})
    ospfv2Link.EntityData.Leafs.Append("type", types.YLeaf{"Type", ospfv2Link.Type})

    ospfv2Link.EntityData.YListKeys = []string {"LinkId", "LinkData"}

    return &(ospfv2Link.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2Link_Ospfv2Topology
// Topology specific information
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2Link_Ospfv2Topology struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. MT-ID for topology enabled link. The type is
    // interface{} with range: 0..4294967295.
    MtId interface{}

    // Metric for the topology. The type is interface{} with range: 0..65535.
    Metric interface{}
}

func (ospfv2Topology *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2Link_Ospfv2Topology) GetEntityData() *types.CommonEntityData {
    ospfv2Topology.EntityData.YFilter = ospfv2Topology.YFilter
    ospfv2Topology.EntityData.YangName = "ospfv2-topology"
    ospfv2Topology.EntityData.BundleName = "cisco_ios_xe"
    ospfv2Topology.EntityData.ParentYangName = "ospfv2-link"
    ospfv2Topology.EntityData.SegmentPath = "ospfv2-topology" + types.AddKeyToken(ospfv2Topology.MtId, "mt-id")
    ospfv2Topology.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/link-scope-lsas/area-scope-lsa/ospfv2-link/" + ospfv2Topology.EntityData.SegmentPath
    ospfv2Topology.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv2Topology.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv2Topology.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv2Topology.EntityData.Children = types.NewOrderedMap()
    ospfv2Topology.EntityData.Leafs = types.NewOrderedMap()
    ospfv2Topology.EntityData.Leafs.Append("mt-id", types.YLeaf{"MtId", ospfv2Topology.MtId})
    ospfv2Topology.EntityData.Leafs.Append("metric", types.YLeaf{"Metric", ospfv2Topology.Metric})

    ospfv2Topology.EntityData.YListKeys = []string {"MtId"}

    return &(ospfv2Topology.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2Topology
// Summary LSA
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2Topology struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. MT-ID for topology enabled link. The type is
    // interface{} with range: 0..4294967295.
    MtId interface{}

    // Metric for the topology. The type is interface{} with range: 0..65535.
    Metric interface{}
}

func (ospfv2Topology *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2Topology) GetEntityData() *types.CommonEntityData {
    ospfv2Topology.EntityData.YFilter = ospfv2Topology.YFilter
    ospfv2Topology.EntityData.YangName = "ospfv2-topology"
    ospfv2Topology.EntityData.BundleName = "cisco_ios_xe"
    ospfv2Topology.EntityData.ParentYangName = "area-scope-lsa"
    ospfv2Topology.EntityData.SegmentPath = "ospfv2-topology" + types.AddKeyToken(ospfv2Topology.MtId, "mt-id")
    ospfv2Topology.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/link-scope-lsas/area-scope-lsa/" + ospfv2Topology.EntityData.SegmentPath
    ospfv2Topology.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv2Topology.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv2Topology.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv2Topology.EntityData.Children = types.NewOrderedMap()
    ospfv2Topology.EntityData.Leafs = types.NewOrderedMap()
    ospfv2Topology.EntityData.Leafs.Append("mt-id", types.YLeaf{"MtId", ospfv2Topology.MtId})
    ospfv2Topology.EntityData.Leafs.Append("metric", types.YLeaf{"Metric", ospfv2Topology.Metric})

    ospfv2Topology.EntityData.YListKeys = []string {"MtId"}

    return &(ospfv2Topology.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2External
// External LSA
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2External struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. MT-ID for topology enabled on the link. The type
    // is interface{} with range: 0..4294967295.
    MtId interface{}

    // Metric for the topology. The type is interface{} with range: 0..4294967295.
    Metric interface{}

    // Forwarding address. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ForwardingAddress interface{}

    // Route tag. The type is interface{} with range: 0..4294967295.
    ExternalRouteTag interface{}
}

func (ospfv2External *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2External) GetEntityData() *types.CommonEntityData {
    ospfv2External.EntityData.YFilter = ospfv2External.YFilter
    ospfv2External.EntityData.YangName = "ospfv2-external"
    ospfv2External.EntityData.BundleName = "cisco_ios_xe"
    ospfv2External.EntityData.ParentYangName = "area-scope-lsa"
    ospfv2External.EntityData.SegmentPath = "ospfv2-external" + types.AddKeyToken(ospfv2External.MtId, "mt-id")
    ospfv2External.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/link-scope-lsas/area-scope-lsa/" + ospfv2External.EntityData.SegmentPath
    ospfv2External.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv2External.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv2External.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv2External.EntityData.Children = types.NewOrderedMap()
    ospfv2External.EntityData.Leafs = types.NewOrderedMap()
    ospfv2External.EntityData.Leafs.Append("mt-id", types.YLeaf{"MtId", ospfv2External.MtId})
    ospfv2External.EntityData.Leafs.Append("metric", types.YLeaf{"Metric", ospfv2External.Metric})
    ospfv2External.EntityData.Leafs.Append("forwarding-address", types.YLeaf{"ForwardingAddress", ospfv2External.ForwardingAddress})
    ospfv2External.EntityData.Leafs.Append("external-route-tag", types.YLeaf{"ExternalRouteTag", ospfv2External.ExternalRouteTag})

    ospfv2External.EntityData.YListKeys = []string {"MtId"}

    return &(ospfv2External.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa
// OSPFv3 LSA
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Decoded OSPFv3 LSA header.
    Header OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_Header

    // Decoded OSPFv3 LSA body.
    LsaBody OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody
}

func (ospfv3Lsa *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa) GetEntityData() *types.CommonEntityData {
    ospfv3Lsa.EntityData.YFilter = ospfv3Lsa.YFilter
    ospfv3Lsa.EntityData.YangName = "ospfv3-lsa"
    ospfv3Lsa.EntityData.BundleName = "cisco_ios_xe"
    ospfv3Lsa.EntityData.ParentYangName = "area-scope-lsa"
    ospfv3Lsa.EntityData.SegmentPath = "ospfv3-lsa"
    ospfv3Lsa.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/link-scope-lsas/area-scope-lsa/" + ospfv3Lsa.EntityData.SegmentPath
    ospfv3Lsa.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv3Lsa.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv3Lsa.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv3Lsa.EntityData.Children = types.NewOrderedMap()
    ospfv3Lsa.EntityData.Children.Append("header", types.YChild{"Header", &ospfv3Lsa.Header})
    ospfv3Lsa.EntityData.Children.Append("lsa-body", types.YChild{"LsaBody", &ospfv3Lsa.LsaBody})
    ospfv3Lsa.EntityData.Leafs = types.NewOrderedMap()

    ospfv3Lsa.EntityData.YListKeys = []string {}

    return &(ospfv3Lsa.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_Header
// Decoded OSPFv3 LSA header
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_Header struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA ID. The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    LsaId interface{}

    // OSPFv3 LSA options. The type is map[string]bool.
    LsaHdrOptions interface{}

    // LSA header.
    LsaHeader OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_Header_LsaHeader
}

func (header *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_Header) GetEntityData() *types.CommonEntityData {
    header.EntityData.YFilter = header.YFilter
    header.EntityData.YangName = "header"
    header.EntityData.BundleName = "cisco_ios_xe"
    header.EntityData.ParentYangName = "ospfv3-lsa"
    header.EntityData.SegmentPath = "header"
    header.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/link-scope-lsas/area-scope-lsa/ospfv3-lsa/" + header.EntityData.SegmentPath
    header.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    header.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    header.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    header.EntityData.Children = types.NewOrderedMap()
    header.EntityData.Children.Append("lsa-header", types.YChild{"LsaHeader", &header.LsaHeader})
    header.EntityData.Leafs = types.NewOrderedMap()
    header.EntityData.Leafs.Append("lsa-id", types.YLeaf{"LsaId", header.LsaId})
    header.EntityData.Leafs.Append("lsa-hdr-options", types.YLeaf{"LsaHdrOptions", header.LsaHdrOptions})

    header.EntityData.YListKeys = []string {}

    return &(header.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_Header_LsaHeader
// LSA header
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_Header_LsaHeader struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA age. The type is interface{} with range: 0..65535.
    Age interface{}

    // LSA type. The type is interface{} with range: 0..65535.
    Type interface{}

    // LSA advertising router. The type is interface{} with range: 0..4294967295.
    AdvRouter interface{}

    // LSA sequence number. The type is string.
    SeqNum interface{}

    // LSA checksum. The type is string.
    Checksum interface{}

    // LSA length. The type is interface{} with range: 0..65535.
    Length interface{}
}

func (lsaHeader *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_Header_LsaHeader) GetEntityData() *types.CommonEntityData {
    lsaHeader.EntityData.YFilter = lsaHeader.YFilter
    lsaHeader.EntityData.YangName = "lsa-header"
    lsaHeader.EntityData.BundleName = "cisco_ios_xe"
    lsaHeader.EntityData.ParentYangName = "header"
    lsaHeader.EntityData.SegmentPath = "lsa-header"
    lsaHeader.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/link-scope-lsas/area-scope-lsa/ospfv3-lsa/header/" + lsaHeader.EntityData.SegmentPath
    lsaHeader.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lsaHeader.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lsaHeader.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lsaHeader.EntityData.Children = types.NewOrderedMap()
    lsaHeader.EntityData.Leafs = types.NewOrderedMap()
    lsaHeader.EntityData.Leafs.Append("age", types.YLeaf{"Age", lsaHeader.Age})
    lsaHeader.EntityData.Leafs.Append("type", types.YLeaf{"Type", lsaHeader.Type})
    lsaHeader.EntityData.Leafs.Append("adv-router", types.YLeaf{"AdvRouter", lsaHeader.AdvRouter})
    lsaHeader.EntityData.Leafs.Append("seq-num", types.YLeaf{"SeqNum", lsaHeader.SeqNum})
    lsaHeader.EntityData.Leafs.Append("checksum", types.YLeaf{"Checksum", lsaHeader.Checksum})
    lsaHeader.EntityData.Leafs.Append("length", types.YLeaf{"Length", lsaHeader.Length})

    lsaHeader.EntityData.YListKeys = []string {}

    return &(lsaHeader.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody
// Decoded OSPFv3 LSA body
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA options. The type is map[string]bool.
    LsaFlagOptions interface{}

    // LSA Body Flags. The type is map[string]bool.
    LsaBodyFlags interface{}

    // OSPFv3 network.
    Network OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Network

    // OSPFv3 inter area prefix.
    Prefix OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Prefix

    // OSPFv3 inter area router.
    IaRouter OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_IaRouter

    // OSPFv3 LSA external.
    LsaExternal OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_LsaExternal

    // OSPFv3 NSSA.
    Nssa OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Nssa

    // OSPFv3 Link data.
    LinkData OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_LinkData

    // OSPFv3 Intra area prefixes.
    IaPrefix OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_IaPrefix
}

func (lsaBody *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody) GetEntityData() *types.CommonEntityData {
    lsaBody.EntityData.YFilter = lsaBody.YFilter
    lsaBody.EntityData.YangName = "lsa-body"
    lsaBody.EntityData.BundleName = "cisco_ios_xe"
    lsaBody.EntityData.ParentYangName = "ospfv3-lsa"
    lsaBody.EntityData.SegmentPath = "lsa-body"
    lsaBody.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/link-scope-lsas/area-scope-lsa/ospfv3-lsa/" + lsaBody.EntityData.SegmentPath
    lsaBody.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lsaBody.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lsaBody.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lsaBody.EntityData.Children = types.NewOrderedMap()
    lsaBody.EntityData.Children.Append("network", types.YChild{"Network", &lsaBody.Network})
    lsaBody.EntityData.Children.Append("prefix", types.YChild{"Prefix", &lsaBody.Prefix})
    lsaBody.EntityData.Children.Append("ia-router", types.YChild{"IaRouter", &lsaBody.IaRouter})
    lsaBody.EntityData.Children.Append("lsa-external", types.YChild{"LsaExternal", &lsaBody.LsaExternal})
    lsaBody.EntityData.Children.Append("nssa", types.YChild{"Nssa", &lsaBody.Nssa})
    lsaBody.EntityData.Children.Append("link-data", types.YChild{"LinkData", &lsaBody.LinkData})
    lsaBody.EntityData.Children.Append("ia-prefix", types.YChild{"IaPrefix", &lsaBody.IaPrefix})
    lsaBody.EntityData.Leafs = types.NewOrderedMap()
    lsaBody.EntityData.Leafs.Append("lsa-flag-options", types.YLeaf{"LsaFlagOptions", lsaBody.LsaFlagOptions})
    lsaBody.EntityData.Leafs.Append("lsa-body-flags", types.YLeaf{"LsaBodyFlags", lsaBody.LsaBodyFlags})

    lsaBody.EntityData.YListKeys = []string {}

    return &(lsaBody.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Network
// OSPFv3 network
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Network struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of the routers attached to the network. The type is slice of
    // interface{} with range: 0..4294967295.
    AttachedRouter []interface{}

    // Network LSA options. The type is map[string]bool.
    LsaNetOptions interface{}
}

func (network *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Network) GetEntityData() *types.CommonEntityData {
    network.EntityData.YFilter = network.YFilter
    network.EntityData.YangName = "network"
    network.EntityData.BundleName = "cisco_ios_xe"
    network.EntityData.ParentYangName = "lsa-body"
    network.EntityData.SegmentPath = "network"
    network.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/link-scope-lsas/area-scope-lsa/ospfv3-lsa/lsa-body/" + network.EntityData.SegmentPath
    network.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    network.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    network.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    network.EntityData.Children = types.NewOrderedMap()
    network.EntityData.Leafs = types.NewOrderedMap()
    network.EntityData.Leafs.Append("attached-router", types.YLeaf{"AttachedRouter", network.AttachedRouter})
    network.EntityData.Leafs.Append("lsa-net-options", types.YLeaf{"LsaNetOptions", network.LsaNetOptions})

    network.EntityData.YListKeys = []string {}

    return &(network.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Prefix
// OSPFv3 inter area prefix
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Prefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Metric. The type is interface{} with range: 0..4294967295.
    Metric interface{}

    // Inter area Prefix. The type is string.
    IaPrefix interface{}

    // Inter area prefix options. The type is string.
    IaPrefixOptions interface{}
}

func (prefix *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Prefix) GetEntityData() *types.CommonEntityData {
    prefix.EntityData.YFilter = prefix.YFilter
    prefix.EntityData.YangName = "prefix"
    prefix.EntityData.BundleName = "cisco_ios_xe"
    prefix.EntityData.ParentYangName = "lsa-body"
    prefix.EntityData.SegmentPath = "prefix"
    prefix.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/link-scope-lsas/area-scope-lsa/ospfv3-lsa/lsa-body/" + prefix.EntityData.SegmentPath
    prefix.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    prefix.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    prefix.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    prefix.EntityData.Children = types.NewOrderedMap()
    prefix.EntityData.Leafs = types.NewOrderedMap()
    prefix.EntityData.Leafs.Append("metric", types.YLeaf{"Metric", prefix.Metric})
    prefix.EntityData.Leafs.Append("ia-prefix", types.YLeaf{"IaPrefix", prefix.IaPrefix})
    prefix.EntityData.Leafs.Append("ia-prefix-options", types.YLeaf{"IaPrefixOptions", prefix.IaPrefixOptions})

    prefix.EntityData.YListKeys = []string {}

    return &(prefix.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_IaRouter
// OSPFv3 inter area router
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_IaRouter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Metric. The type is interface{} with range: 0..4294967295.
    Metric interface{}

    // Router ID of the router being described by the LSA. The type is interface{}
    // with range: 0..4294967295.
    DestinationRouterId interface{}

    // Inter area LSA options. The type is map[string]bool.
    LsaIaOptions interface{}
}

func (iaRouter *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_IaRouter) GetEntityData() *types.CommonEntityData {
    iaRouter.EntityData.YFilter = iaRouter.YFilter
    iaRouter.EntityData.YangName = "ia-router"
    iaRouter.EntityData.BundleName = "cisco_ios_xe"
    iaRouter.EntityData.ParentYangName = "lsa-body"
    iaRouter.EntityData.SegmentPath = "ia-router"
    iaRouter.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/link-scope-lsas/area-scope-lsa/ospfv3-lsa/lsa-body/" + iaRouter.EntityData.SegmentPath
    iaRouter.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    iaRouter.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    iaRouter.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    iaRouter.EntityData.Children = types.NewOrderedMap()
    iaRouter.EntityData.Leafs = types.NewOrderedMap()
    iaRouter.EntityData.Leafs.Append("metric", types.YLeaf{"Metric", iaRouter.Metric})
    iaRouter.EntityData.Leafs.Append("destination-router-id", types.YLeaf{"DestinationRouterId", iaRouter.DestinationRouterId})
    iaRouter.EntityData.Leafs.Append("lsa-ia-options", types.YLeaf{"LsaIaOptions", iaRouter.LsaIaOptions})

    iaRouter.EntityData.YListKeys = []string {}

    return &(iaRouter.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_LsaExternal
// OSPFv3 LSA external
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_LsaExternal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Metric. The type is interface{} with range: 0..4294967295.
    Metric interface{}

    // Referenced Link State type. The type is interface{} with range: 0..65535.
    ReferencedLsType interface{}

    // Prefix. The type is string.
    ExternalPrefix interface{}

    // Prefix options. The type is string.
    ExternalPrefixOptions interface{}

    // Forwarding address. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ForwardingAddress interface{}

    // Route tag. The type is interface{} with range: 0..4294967295.
    ExternalRouteTag interface{}

    // Referenced Link State ID. The type is interface{} with range:
    // 0..4294967295.
    ReferencedLinkStateId interface{}

    // LSA Flags.
    Flags OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_LsaExternal_Flags
}

func (lsaExternal *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_LsaExternal) GetEntityData() *types.CommonEntityData {
    lsaExternal.EntityData.YFilter = lsaExternal.YFilter
    lsaExternal.EntityData.YangName = "lsa-external"
    lsaExternal.EntityData.BundleName = "cisco_ios_xe"
    lsaExternal.EntityData.ParentYangName = "lsa-body"
    lsaExternal.EntityData.SegmentPath = "lsa-external"
    lsaExternal.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/link-scope-lsas/area-scope-lsa/ospfv3-lsa/lsa-body/" + lsaExternal.EntityData.SegmentPath
    lsaExternal.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lsaExternal.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lsaExternal.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lsaExternal.EntityData.Children = types.NewOrderedMap()
    lsaExternal.EntityData.Children.Append("flags", types.YChild{"Flags", &lsaExternal.Flags})
    lsaExternal.EntityData.Leafs = types.NewOrderedMap()
    lsaExternal.EntityData.Leafs.Append("metric", types.YLeaf{"Metric", lsaExternal.Metric})
    lsaExternal.EntityData.Leafs.Append("referenced-ls-type", types.YLeaf{"ReferencedLsType", lsaExternal.ReferencedLsType})
    lsaExternal.EntityData.Leafs.Append("external-prefix", types.YLeaf{"ExternalPrefix", lsaExternal.ExternalPrefix})
    lsaExternal.EntityData.Leafs.Append("external-prefix-options", types.YLeaf{"ExternalPrefixOptions", lsaExternal.ExternalPrefixOptions})
    lsaExternal.EntityData.Leafs.Append("forwarding-address", types.YLeaf{"ForwardingAddress", lsaExternal.ForwardingAddress})
    lsaExternal.EntityData.Leafs.Append("external-route-tag", types.YLeaf{"ExternalRouteTag", lsaExternal.ExternalRouteTag})
    lsaExternal.EntityData.Leafs.Append("referenced-link-state-id", types.YLeaf{"ReferencedLinkStateId", lsaExternal.ReferencedLinkStateId})

    lsaExternal.EntityData.YListKeys = []string {}

    return &(lsaExternal.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_LsaExternal_Flags
// LSA Flags
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_LsaExternal_Flags struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // When set, the metric specified is a Type 2 external metric. The type is
    // bool.
    EFlag interface{}
}

func (flags *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_LsaExternal_Flags) GetEntityData() *types.CommonEntityData {
    flags.EntityData.YFilter = flags.YFilter
    flags.EntityData.YangName = "flags"
    flags.EntityData.BundleName = "cisco_ios_xe"
    flags.EntityData.ParentYangName = "lsa-external"
    flags.EntityData.SegmentPath = "flags"
    flags.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/link-scope-lsas/area-scope-lsa/ospfv3-lsa/lsa-body/lsa-external/" + flags.EntityData.SegmentPath
    flags.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    flags.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    flags.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    flags.EntityData.Children = types.NewOrderedMap()
    flags.EntityData.Leafs = types.NewOrderedMap()
    flags.EntityData.Leafs.Append("e-flag", types.YLeaf{"EFlag", flags.EFlag})

    flags.EntityData.YListKeys = []string {}

    return &(flags.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Nssa
// OSPFv3 NSSA
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Nssa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // NSSA LSA.
    LsaNssaExternal OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Nssa_LsaNssaExternal
}

func (nssa *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Nssa) GetEntityData() *types.CommonEntityData {
    nssa.EntityData.YFilter = nssa.YFilter
    nssa.EntityData.YangName = "nssa"
    nssa.EntityData.BundleName = "cisco_ios_xe"
    nssa.EntityData.ParentYangName = "lsa-body"
    nssa.EntityData.SegmentPath = "nssa"
    nssa.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/link-scope-lsas/area-scope-lsa/ospfv3-lsa/lsa-body/" + nssa.EntityData.SegmentPath
    nssa.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nssa.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nssa.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nssa.EntityData.Children = types.NewOrderedMap()
    nssa.EntityData.Children.Append("lsa-nssa-external", types.YChild{"LsaNssaExternal", &nssa.LsaNssaExternal})
    nssa.EntityData.Leafs = types.NewOrderedMap()

    nssa.EntityData.YListKeys = []string {}

    return &(nssa.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Nssa_LsaNssaExternal
// NSSA LSA
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Nssa_LsaNssaExternal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Metric. The type is interface{} with range: 0..4294967295.
    Metric interface{}

    // Referenced Link State type. The type is interface{} with range: 0..65535.
    ReferencedLsType interface{}

    // Prefix. The type is string.
    ExternalPrefix interface{}

    // Prefix options. The type is string.
    ExternalPrefixOptions interface{}

    // Forwarding address. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ForwardingAddress interface{}

    // Route tag. The type is interface{} with range: 0..4294967295.
    ExternalRouteTag interface{}

    // Referenced Link State ID. The type is interface{} with range:
    // 0..4294967295.
    ReferencedLinkStateId interface{}

    // LSA Flags.
    Flags OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Nssa_LsaNssaExternal_Flags
}

func (lsaNssaExternal *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Nssa_LsaNssaExternal) GetEntityData() *types.CommonEntityData {
    lsaNssaExternal.EntityData.YFilter = lsaNssaExternal.YFilter
    lsaNssaExternal.EntityData.YangName = "lsa-nssa-external"
    lsaNssaExternal.EntityData.BundleName = "cisco_ios_xe"
    lsaNssaExternal.EntityData.ParentYangName = "nssa"
    lsaNssaExternal.EntityData.SegmentPath = "lsa-nssa-external"
    lsaNssaExternal.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/link-scope-lsas/area-scope-lsa/ospfv3-lsa/lsa-body/nssa/" + lsaNssaExternal.EntityData.SegmentPath
    lsaNssaExternal.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lsaNssaExternal.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lsaNssaExternal.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lsaNssaExternal.EntityData.Children = types.NewOrderedMap()
    lsaNssaExternal.EntityData.Children.Append("flags", types.YChild{"Flags", &lsaNssaExternal.Flags})
    lsaNssaExternal.EntityData.Leafs = types.NewOrderedMap()
    lsaNssaExternal.EntityData.Leafs.Append("metric", types.YLeaf{"Metric", lsaNssaExternal.Metric})
    lsaNssaExternal.EntityData.Leafs.Append("referenced-ls-type", types.YLeaf{"ReferencedLsType", lsaNssaExternal.ReferencedLsType})
    lsaNssaExternal.EntityData.Leafs.Append("external-prefix", types.YLeaf{"ExternalPrefix", lsaNssaExternal.ExternalPrefix})
    lsaNssaExternal.EntityData.Leafs.Append("external-prefix-options", types.YLeaf{"ExternalPrefixOptions", lsaNssaExternal.ExternalPrefixOptions})
    lsaNssaExternal.EntityData.Leafs.Append("forwarding-address", types.YLeaf{"ForwardingAddress", lsaNssaExternal.ForwardingAddress})
    lsaNssaExternal.EntityData.Leafs.Append("external-route-tag", types.YLeaf{"ExternalRouteTag", lsaNssaExternal.ExternalRouteTag})
    lsaNssaExternal.EntityData.Leafs.Append("referenced-link-state-id", types.YLeaf{"ReferencedLinkStateId", lsaNssaExternal.ReferencedLinkStateId})

    lsaNssaExternal.EntityData.YListKeys = []string {}

    return &(lsaNssaExternal.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Nssa_LsaNssaExternal_Flags
// LSA Flags
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Nssa_LsaNssaExternal_Flags struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // When set, the metric specified is a Type 2 external metric. The type is
    // bool.
    EFlag interface{}
}

func (flags *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Nssa_LsaNssaExternal_Flags) GetEntityData() *types.CommonEntityData {
    flags.EntityData.YFilter = flags.YFilter
    flags.EntityData.YangName = "flags"
    flags.EntityData.BundleName = "cisco_ios_xe"
    flags.EntityData.ParentYangName = "lsa-nssa-external"
    flags.EntityData.SegmentPath = "flags"
    flags.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/link-scope-lsas/area-scope-lsa/ospfv3-lsa/lsa-body/nssa/lsa-nssa-external/" + flags.EntityData.SegmentPath
    flags.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    flags.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    flags.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    flags.EntityData.Children = types.NewOrderedMap()
    flags.EntityData.Leafs = types.NewOrderedMap()
    flags.EntityData.Leafs.Append("e-flag", types.YLeaf{"EFlag", flags.EFlag})

    flags.EntityData.YListKeys = []string {}

    return &(flags.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_LinkData
// OSPFv3 Link data
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_LinkData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Router priority of the interce. The type is interface{} with range: 0..255.
    RtrPriority interface{}

    // The originating router's link-local interface address on the link. The type
    // is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    LinkLocalInterfaceAddress interface{}

    // Number of prefixes. The type is interface{} with range: 0..4294967295.
    NumOfPrefixes interface{}

    // Link data LSA options. The type is map[string]bool.
    LsaIdOptions interface{}
}

func (linkData *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_LinkData) GetEntityData() *types.CommonEntityData {
    linkData.EntityData.YFilter = linkData.YFilter
    linkData.EntityData.YangName = "link-data"
    linkData.EntityData.BundleName = "cisco_ios_xe"
    linkData.EntityData.ParentYangName = "lsa-body"
    linkData.EntityData.SegmentPath = "link-data"
    linkData.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/link-scope-lsas/area-scope-lsa/ospfv3-lsa/lsa-body/" + linkData.EntityData.SegmentPath
    linkData.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    linkData.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    linkData.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    linkData.EntityData.Children = types.NewOrderedMap()
    linkData.EntityData.Leafs = types.NewOrderedMap()
    linkData.EntityData.Leafs.Append("rtr-priority", types.YLeaf{"RtrPriority", linkData.RtrPriority})
    linkData.EntityData.Leafs.Append("link-local-interface-address", types.YLeaf{"LinkLocalInterfaceAddress", linkData.LinkLocalInterfaceAddress})
    linkData.EntityData.Leafs.Append("num-of-prefixes", types.YLeaf{"NumOfPrefixes", linkData.NumOfPrefixes})
    linkData.EntityData.Leafs.Append("lsa-id-options", types.YLeaf{"LsaIdOptions", linkData.LsaIdOptions})

    linkData.EntityData.YListKeys = []string {}

    return &(linkData.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_IaPrefix
// OSPFv3 Intra area prefixes
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_IaPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Referenced Link State type. The type is interface{} with range: 0..65535.
    ReferencedLsType interface{}

    // Referenced Link State ID. The type is interface{} with range:
    // 0..4294967295.
    ReferencedLinkStateId interface{}

    // Referenced Advertising Router. The type is one of the following types:
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ReferencedAdvRouter interface{}

    // Number of prefixes. The type is interface{} with range: 0..65535.
    NumOfPrefixes interface{}
}

func (iaPrefix *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_IaPrefix) GetEntityData() *types.CommonEntityData {
    iaPrefix.EntityData.YFilter = iaPrefix.YFilter
    iaPrefix.EntityData.YangName = "ia-prefix"
    iaPrefix.EntityData.BundleName = "cisco_ios_xe"
    iaPrefix.EntityData.ParentYangName = "lsa-body"
    iaPrefix.EntityData.SegmentPath = "ia-prefix"
    iaPrefix.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/link-scope-lsas/area-scope-lsa/ospfv3-lsa/lsa-body/" + iaPrefix.EntityData.SegmentPath
    iaPrefix.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    iaPrefix.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    iaPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    iaPrefix.EntityData.Children = types.NewOrderedMap()
    iaPrefix.EntityData.Leafs = types.NewOrderedMap()
    iaPrefix.EntityData.Leafs.Append("referenced-ls-type", types.YLeaf{"ReferencedLsType", iaPrefix.ReferencedLsType})
    iaPrefix.EntityData.Leafs.Append("referenced-link-state-id", types.YLeaf{"ReferencedLinkStateId", iaPrefix.ReferencedLinkStateId})
    iaPrefix.EntityData.Leafs.Append("referenced-adv-router", types.YLeaf{"ReferencedAdvRouter", iaPrefix.ReferencedAdvRouter})
    iaPrefix.EntityData.Leafs.Append("num-of-prefixes", types.YLeaf{"NumOfPrefixes", iaPrefix.NumOfPrefixes})

    iaPrefix.EntityData.YListKeys = []string {}

    return &(iaPrefix.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Link
// OSPFv3 links
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Link struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface ID. The type is interface{} with range:
    // 0..4294967295.
    InterfaceId interface{}

    // This attribute is a key. Neighbor interface ID. The type is interface{}
    // with range: 0..4294967295.
    NeighborInterfaceId interface{}

    // This attribute is a key. Neighbor router ID. The type is interface{} with
    // range: 0..4294967295.
    NeighborRouterId interface{}

    // Link type. The type is interface{} with range: 0..255.
    Type interface{}

    // Metric. The type is interface{} with range: 0..65535.
    Metric interface{}
}

func (ospfv3Link *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Link) GetEntityData() *types.CommonEntityData {
    ospfv3Link.EntityData.YFilter = ospfv3Link.YFilter
    ospfv3Link.EntityData.YangName = "ospfv3-link"
    ospfv3Link.EntityData.BundleName = "cisco_ios_xe"
    ospfv3Link.EntityData.ParentYangName = "area-scope-lsa"
    ospfv3Link.EntityData.SegmentPath = "ospfv3-link" + types.AddKeyToken(ospfv3Link.InterfaceId, "interface-id") + types.AddKeyToken(ospfv3Link.NeighborInterfaceId, "neighbor-interface-id") + types.AddKeyToken(ospfv3Link.NeighborRouterId, "neighbor-router-id")
    ospfv3Link.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/link-scope-lsas/area-scope-lsa/" + ospfv3Link.EntityData.SegmentPath
    ospfv3Link.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv3Link.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv3Link.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv3Link.EntityData.Children = types.NewOrderedMap()
    ospfv3Link.EntityData.Leafs = types.NewOrderedMap()
    ospfv3Link.EntityData.Leafs.Append("interface-id", types.YLeaf{"InterfaceId", ospfv3Link.InterfaceId})
    ospfv3Link.EntityData.Leafs.Append("neighbor-interface-id", types.YLeaf{"NeighborInterfaceId", ospfv3Link.NeighborInterfaceId})
    ospfv3Link.EntityData.Leafs.Append("neighbor-router-id", types.YLeaf{"NeighborRouterId", ospfv3Link.NeighborRouterId})
    ospfv3Link.EntityData.Leafs.Append("type", types.YLeaf{"Type", ospfv3Link.Type})
    ospfv3Link.EntityData.Leafs.Append("metric", types.YLeaf{"Metric", ospfv3Link.Metric})

    ospfv3Link.EntityData.YListKeys = []string {"InterfaceId", "NeighborInterfaceId", "NeighborRouterId"}

    return &(ospfv3Link.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Prefix
// OSPFv3 prefix-list
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Prefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Prefix. The type is string.
    Prefix interface{}

    // Prefix options. The type is string.
    PrefixOptions interface{}
}

func (ospfv3Prefix *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Prefix) GetEntityData() *types.CommonEntityData {
    ospfv3Prefix.EntityData.YFilter = ospfv3Prefix.YFilter
    ospfv3Prefix.EntityData.YangName = "ospfv3-prefix"
    ospfv3Prefix.EntityData.BundleName = "cisco_ios_xe"
    ospfv3Prefix.EntityData.ParentYangName = "area-scope-lsa"
    ospfv3Prefix.EntityData.SegmentPath = "ospfv3-prefix" + types.AddKeyToken(ospfv3Prefix.Prefix, "prefix")
    ospfv3Prefix.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/link-scope-lsas/area-scope-lsa/" + ospfv3Prefix.EntityData.SegmentPath
    ospfv3Prefix.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv3Prefix.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv3Prefix.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv3Prefix.EntityData.Children = types.NewOrderedMap()
    ospfv3Prefix.EntityData.Leafs = types.NewOrderedMap()
    ospfv3Prefix.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", ospfv3Prefix.Prefix})
    ospfv3Prefix.EntityData.Leafs.Append("prefix-options", types.YLeaf{"PrefixOptions", ospfv3Prefix.PrefixOptions})

    ospfv3Prefix.EntityData.YListKeys = []string {"Prefix"}

    return &(ospfv3Prefix.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3IaPrefix
// OSPFv3 intra-area prefix-list
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3IaPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Prefix. The type is string.
    Prefix interface{}

    // Prefix options. The type is string.
    PrefixOptions interface{}
}

func (ospfv3IaPrefix *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3IaPrefix) GetEntityData() *types.CommonEntityData {
    ospfv3IaPrefix.EntityData.YFilter = ospfv3IaPrefix.YFilter
    ospfv3IaPrefix.EntityData.YangName = "ospfv3-ia-prefix"
    ospfv3IaPrefix.EntityData.BundleName = "cisco_ios_xe"
    ospfv3IaPrefix.EntityData.ParentYangName = "area-scope-lsa"
    ospfv3IaPrefix.EntityData.SegmentPath = "ospfv3-ia-prefix" + types.AddKeyToken(ospfv3IaPrefix.Prefix, "prefix")
    ospfv3IaPrefix.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/link-scope-lsas/area-scope-lsa/" + ospfv3IaPrefix.EntityData.SegmentPath
    ospfv3IaPrefix.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv3IaPrefix.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv3IaPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv3IaPrefix.EntityData.Children = types.NewOrderedMap()
    ospfv3IaPrefix.EntityData.Leafs = types.NewOrderedMap()
    ospfv3IaPrefix.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", ospfv3IaPrefix.Prefix})
    ospfv3IaPrefix.EntityData.Leafs.Append("prefix-options", types.YLeaf{"PrefixOptions", ospfv3IaPrefix.PrefixOptions})

    ospfv3IaPrefix.EntityData.YListKeys = []string {"Prefix"}

    return &(ospfv3IaPrefix.EntityData)
}

// OspfOperData_OspfState_OspfInstance_MultiTopology
// OSPF multi-topology interface augmentation
type OspfOperData_OspfState_OspfInstance_MultiTopology struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. One of the topology enabled on this interface. The
    // type is string.
    Name interface{}
}

func (multiTopology *OspfOperData_OspfState_OspfInstance_MultiTopology) GetEntityData() *types.CommonEntityData {
    multiTopology.EntityData.YFilter = multiTopology.YFilter
    multiTopology.EntityData.YangName = "multi-topology"
    multiTopology.EntityData.BundleName = "cisco_ios_xe"
    multiTopology.EntityData.ParentYangName = "ospf-instance"
    multiTopology.EntityData.SegmentPath = "multi-topology" + types.AddKeyToken(multiTopology.Name, "name")
    multiTopology.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospf-state/ospf-instance/" + multiTopology.EntityData.SegmentPath
    multiTopology.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    multiTopology.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    multiTopology.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    multiTopology.EntityData.Children = types.NewOrderedMap()
    multiTopology.EntityData.Leafs = types.NewOrderedMap()
    multiTopology.EntityData.Leafs.Append("name", types.YLeaf{"Name", multiTopology.Name})

    multiTopology.EntityData.YListKeys = []string {"Name"}

    return &(multiTopology.EntityData)
}

// OspfOperData_Ospfv2Instance
// The OSPF instance
type OspfOperData_Ospfv2Instance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The routing instance identifier assigned to the
    // OSPF instance. The type is interface{} with range: 0..4294967295.
    InstanceId interface{}

    // The name of the Virtual Routing and Forwarding instance that the OSPF
    // instance is operating within. The type is string.
    VrfName interface{}

    // The router identifer assigned to the OSPF instance. The type is interface{}
    // with range: 0..4294967295.
    RouterId interface{}

    // The OSPF area information. The type is slice of
    // OspfOperData_Ospfv2Instance_Ospfv2Area.
    Ospfv2Area []*OspfOperData_Ospfv2Instance_Ospfv2Area

    // The external LSDB information. The type is slice of
    // OspfOperData_Ospfv2Instance_Ospfv2LsdbExternal.
    Ospfv2LsdbExternal []*OspfOperData_Ospfv2Instance_Ospfv2LsdbExternal
}

func (ospfv2Instance *OspfOperData_Ospfv2Instance) GetEntityData() *types.CommonEntityData {
    ospfv2Instance.EntityData.YFilter = ospfv2Instance.YFilter
    ospfv2Instance.EntityData.YangName = "ospfv2-instance"
    ospfv2Instance.EntityData.BundleName = "cisco_ios_xe"
    ospfv2Instance.EntityData.ParentYangName = "ospf-oper-data"
    ospfv2Instance.EntityData.SegmentPath = "ospfv2-instance" + types.AddKeyToken(ospfv2Instance.InstanceId, "instance-id")
    ospfv2Instance.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/" + ospfv2Instance.EntityData.SegmentPath
    ospfv2Instance.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv2Instance.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv2Instance.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv2Instance.EntityData.Children = types.NewOrderedMap()
    ospfv2Instance.EntityData.Children.Append("ospfv2-area", types.YChild{"Ospfv2Area", nil})
    for i := range ospfv2Instance.Ospfv2Area {
        ospfv2Instance.EntityData.Children.Append(types.GetSegmentPath(ospfv2Instance.Ospfv2Area[i]), types.YChild{"Ospfv2Area", ospfv2Instance.Ospfv2Area[i]})
    }
    ospfv2Instance.EntityData.Children.Append("ospfv2-lsdb-external", types.YChild{"Ospfv2LsdbExternal", nil})
    for i := range ospfv2Instance.Ospfv2LsdbExternal {
        ospfv2Instance.EntityData.Children.Append(types.GetSegmentPath(ospfv2Instance.Ospfv2LsdbExternal[i]), types.YChild{"Ospfv2LsdbExternal", ospfv2Instance.Ospfv2LsdbExternal[i]})
    }
    ospfv2Instance.EntityData.Leafs = types.NewOrderedMap()
    ospfv2Instance.EntityData.Leafs.Append("instance-id", types.YLeaf{"InstanceId", ospfv2Instance.InstanceId})
    ospfv2Instance.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", ospfv2Instance.VrfName})
    ospfv2Instance.EntityData.Leafs.Append("router-id", types.YLeaf{"RouterId", ospfv2Instance.RouterId})

    ospfv2Instance.EntityData.YListKeys = []string {"InstanceId"}

    return &(ospfv2Instance.EntityData)
}

// OspfOperData_Ospfv2Instance_Ospfv2Area
// The OSPF area information
type OspfOperData_Ospfv2Instance_Ospfv2Area struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The area identifier. The type is interface{} with
    // range: 0..4294967295.
    AreaId interface{}

    // The OSPF Link State Database information for this area. The type is slice
    // of OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2LsdbArea.
    Ospfv2LsdbArea []*OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2LsdbArea

    // A list of interfaces that belong to the area. The type is slice of
    // OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2Interface.
    Ospfv2Interface []*OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2Interface
}

func (ospfv2Area *OspfOperData_Ospfv2Instance_Ospfv2Area) GetEntityData() *types.CommonEntityData {
    ospfv2Area.EntityData.YFilter = ospfv2Area.YFilter
    ospfv2Area.EntityData.YangName = "ospfv2-area"
    ospfv2Area.EntityData.BundleName = "cisco_ios_xe"
    ospfv2Area.EntityData.ParentYangName = "ospfv2-instance"
    ospfv2Area.EntityData.SegmentPath = "ospfv2-area" + types.AddKeyToken(ospfv2Area.AreaId, "area-id")
    ospfv2Area.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospfv2-instance/" + ospfv2Area.EntityData.SegmentPath
    ospfv2Area.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv2Area.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv2Area.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv2Area.EntityData.Children = types.NewOrderedMap()
    ospfv2Area.EntityData.Children.Append("ospfv2-lsdb-area", types.YChild{"Ospfv2LsdbArea", nil})
    for i := range ospfv2Area.Ospfv2LsdbArea {
        ospfv2Area.EntityData.Children.Append(types.GetSegmentPath(ospfv2Area.Ospfv2LsdbArea[i]), types.YChild{"Ospfv2LsdbArea", ospfv2Area.Ospfv2LsdbArea[i]})
    }
    ospfv2Area.EntityData.Children.Append("ospfv2-interface", types.YChild{"Ospfv2Interface", nil})
    for i := range ospfv2Area.Ospfv2Interface {
        ospfv2Area.EntityData.Children.Append(types.GetSegmentPath(ospfv2Area.Ospfv2Interface[i]), types.YChild{"Ospfv2Interface", ospfv2Area.Ospfv2Interface[i]})
    }
    ospfv2Area.EntityData.Leafs = types.NewOrderedMap()
    ospfv2Area.EntityData.Leafs.Append("area-id", types.YLeaf{"AreaId", ospfv2Area.AreaId})

    ospfv2Area.EntityData.YListKeys = []string {"AreaId"}

    return &(ospfv2Area.EntityData)
}

// OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2LsdbArea
// The OSPF Link State Database information for this area
type OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2LsdbArea struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Link State Advertisement type. The type is
    // interface{} with range: 0..255.
    LsaType interface{}

    // This attribute is a key. Link State Advertisement Identifer. The type is
    // interface{} with range: 0..4294967295.
    LsaId interface{}

    // This attribute is a key. Advertising router. The type is interface{} with
    // range: 0..4294967295.
    AdvertisingRouter interface{}

    // The age of the Link State Advertisement. The type is interface{} with
    // range: 0..65535.
    LsaAge interface{}

    // The options of the Link State Advertisement. The type is map[string]bool.
    LsaOptions interface{}

    // The sequence number for the Link State Advertisement. The type is
    // interface{} with range: 0..4294967295.
    LsaSeqNumber interface{}

    // The checksum of the Link State Advertisement. The type is interface{} with
    // range: 0..65535.
    LsaChecksum interface{}

    // The length, in bytes, of the Link State Advertisement. The type is
    // interface{} with range: 0..65535.
    LsaLength interface{}

    // The router Link State Advertisement links. The type is slice of
    // OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2LsdbArea_Ospfv2RouterLsaLinks.
    Ospfv2RouterLsaLinks []*OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2LsdbArea_Ospfv2RouterLsaLinks

    // The unsupported Link State Advertisements.
    UnsupportedLsa OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2LsdbArea_UnsupportedLsa

    // The router Link State Advertisements.
    RouterLsa OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2LsdbArea_RouterLsa

    // The network Link State Advertisements.
    NetworkLsa OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2LsdbArea_NetworkLsa

    // The network summary Link State Advertisements.
    NetworkSummaryLsa OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2LsdbArea_NetworkSummaryLsa

    // The router summary Link State Advertisements.
    RouterSummaryLsa OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2LsdbArea_RouterSummaryLsa

    // The external Link State Advertisements.
    ExternalLsa OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2LsdbArea_ExternalLsa

    // The Not So Stubby Area Link state advertisements.
    NssaLsa OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2LsdbArea_NssaLsa
}

func (ospfv2LsdbArea *OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2LsdbArea) GetEntityData() *types.CommonEntityData {
    ospfv2LsdbArea.EntityData.YFilter = ospfv2LsdbArea.YFilter
    ospfv2LsdbArea.EntityData.YangName = "ospfv2-lsdb-area"
    ospfv2LsdbArea.EntityData.BundleName = "cisco_ios_xe"
    ospfv2LsdbArea.EntityData.ParentYangName = "ospfv2-area"
    ospfv2LsdbArea.EntityData.SegmentPath = "ospfv2-lsdb-area" + types.AddKeyToken(ospfv2LsdbArea.LsaType, "lsa-type") + types.AddKeyToken(ospfv2LsdbArea.LsaId, "lsa-id") + types.AddKeyToken(ospfv2LsdbArea.AdvertisingRouter, "advertising-router")
    ospfv2LsdbArea.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospfv2-instance/ospfv2-area/" + ospfv2LsdbArea.EntityData.SegmentPath
    ospfv2LsdbArea.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv2LsdbArea.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv2LsdbArea.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv2LsdbArea.EntityData.Children = types.NewOrderedMap()
    ospfv2LsdbArea.EntityData.Children.Append("ospfv2-router-lsa-links", types.YChild{"Ospfv2RouterLsaLinks", nil})
    for i := range ospfv2LsdbArea.Ospfv2RouterLsaLinks {
        ospfv2LsdbArea.EntityData.Children.Append(types.GetSegmentPath(ospfv2LsdbArea.Ospfv2RouterLsaLinks[i]), types.YChild{"Ospfv2RouterLsaLinks", ospfv2LsdbArea.Ospfv2RouterLsaLinks[i]})
    }
    ospfv2LsdbArea.EntityData.Children.Append("unsupported-lsa", types.YChild{"UnsupportedLsa", &ospfv2LsdbArea.UnsupportedLsa})
    ospfv2LsdbArea.EntityData.Children.Append("router-lsa", types.YChild{"RouterLsa", &ospfv2LsdbArea.RouterLsa})
    ospfv2LsdbArea.EntityData.Children.Append("network-lsa", types.YChild{"NetworkLsa", &ospfv2LsdbArea.NetworkLsa})
    ospfv2LsdbArea.EntityData.Children.Append("network-summary-lsa", types.YChild{"NetworkSummaryLsa", &ospfv2LsdbArea.NetworkSummaryLsa})
    ospfv2LsdbArea.EntityData.Children.Append("router-summary-lsa", types.YChild{"RouterSummaryLsa", &ospfv2LsdbArea.RouterSummaryLsa})
    ospfv2LsdbArea.EntityData.Children.Append("external-lsa", types.YChild{"ExternalLsa", &ospfv2LsdbArea.ExternalLsa})
    ospfv2LsdbArea.EntityData.Children.Append("nssa-lsa", types.YChild{"NssaLsa", &ospfv2LsdbArea.NssaLsa})
    ospfv2LsdbArea.EntityData.Leafs = types.NewOrderedMap()
    ospfv2LsdbArea.EntityData.Leafs.Append("lsa-type", types.YLeaf{"LsaType", ospfv2LsdbArea.LsaType})
    ospfv2LsdbArea.EntityData.Leafs.Append("lsa-id", types.YLeaf{"LsaId", ospfv2LsdbArea.LsaId})
    ospfv2LsdbArea.EntityData.Leafs.Append("advertising-router", types.YLeaf{"AdvertisingRouter", ospfv2LsdbArea.AdvertisingRouter})
    ospfv2LsdbArea.EntityData.Leafs.Append("lsa-age", types.YLeaf{"LsaAge", ospfv2LsdbArea.LsaAge})
    ospfv2LsdbArea.EntityData.Leafs.Append("lsa-options", types.YLeaf{"LsaOptions", ospfv2LsdbArea.LsaOptions})
    ospfv2LsdbArea.EntityData.Leafs.Append("lsa-seq-number", types.YLeaf{"LsaSeqNumber", ospfv2LsdbArea.LsaSeqNumber})
    ospfv2LsdbArea.EntityData.Leafs.Append("lsa-checksum", types.YLeaf{"LsaChecksum", ospfv2LsdbArea.LsaChecksum})
    ospfv2LsdbArea.EntityData.Leafs.Append("lsa-length", types.YLeaf{"LsaLength", ospfv2LsdbArea.LsaLength})

    ospfv2LsdbArea.EntityData.YListKeys = []string {"LsaType", "LsaId", "AdvertisingRouter"}

    return &(ospfv2LsdbArea.EntityData)
}

// OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2LsdbArea_Ospfv2RouterLsaLinks
// The router Link State Advertisement links
type OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2LsdbArea_Ospfv2RouterLsaLinks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Link Type. The type is interface{} with range:
    // 0..255.
    LinkType interface{}

    // This attribute is a key. link Identifier. The type is interface{} with
    // range: 0..4294967295.
    LinkId interface{}

    // This attribute is a key. link data. The type is interface{} with range:
    // 0..4294967295.
    LinkData interface{}

    // Link topology. The type is slice of
    // OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2LsdbArea_Ospfv2RouterLsaLinks_LinkTopo.
    LinkTopo []*OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2LsdbArea_Ospfv2RouterLsaLinks_LinkTopo
}

func (ospfv2RouterLsaLinks *OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2LsdbArea_Ospfv2RouterLsaLinks) GetEntityData() *types.CommonEntityData {
    ospfv2RouterLsaLinks.EntityData.YFilter = ospfv2RouterLsaLinks.YFilter
    ospfv2RouterLsaLinks.EntityData.YangName = "ospfv2-router-lsa-links"
    ospfv2RouterLsaLinks.EntityData.BundleName = "cisco_ios_xe"
    ospfv2RouterLsaLinks.EntityData.ParentYangName = "ospfv2-lsdb-area"
    ospfv2RouterLsaLinks.EntityData.SegmentPath = "ospfv2-router-lsa-links" + types.AddKeyToken(ospfv2RouterLsaLinks.LinkType, "link-type") + types.AddKeyToken(ospfv2RouterLsaLinks.LinkId, "link-id") + types.AddKeyToken(ospfv2RouterLsaLinks.LinkData, "link-data")
    ospfv2RouterLsaLinks.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospfv2-instance/ospfv2-area/ospfv2-lsdb-area/" + ospfv2RouterLsaLinks.EntityData.SegmentPath
    ospfv2RouterLsaLinks.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv2RouterLsaLinks.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv2RouterLsaLinks.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv2RouterLsaLinks.EntityData.Children = types.NewOrderedMap()
    ospfv2RouterLsaLinks.EntityData.Children.Append("link-topo", types.YChild{"LinkTopo", nil})
    for i := range ospfv2RouterLsaLinks.LinkTopo {
        types.SetYListKey(ospfv2RouterLsaLinks.LinkTopo[i], i)
        ospfv2RouterLsaLinks.EntityData.Children.Append(types.GetSegmentPath(ospfv2RouterLsaLinks.LinkTopo[i]), types.YChild{"LinkTopo", ospfv2RouterLsaLinks.LinkTopo[i]})
    }
    ospfv2RouterLsaLinks.EntityData.Leafs = types.NewOrderedMap()
    ospfv2RouterLsaLinks.EntityData.Leafs.Append("link-type", types.YLeaf{"LinkType", ospfv2RouterLsaLinks.LinkType})
    ospfv2RouterLsaLinks.EntityData.Leafs.Append("link-id", types.YLeaf{"LinkId", ospfv2RouterLsaLinks.LinkId})
    ospfv2RouterLsaLinks.EntityData.Leafs.Append("link-data", types.YLeaf{"LinkData", ospfv2RouterLsaLinks.LinkData})

    ospfv2RouterLsaLinks.EntityData.YListKeys = []string {"LinkType", "LinkId", "LinkData"}

    return &(ospfv2RouterLsaLinks.EntityData)
}

// OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2LsdbArea_Ospfv2RouterLsaLinks_LinkTopo
// Link topology
type OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2LsdbArea_Ospfv2RouterLsaLinks_LinkTopo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Multi topology identifier. The type is interface{} with range: 0..255.
    MtId interface{}

    // Topology metric. The type is interface{} with range: 0..65535.
    TopoMetric interface{}
}

func (linkTopo *OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2LsdbArea_Ospfv2RouterLsaLinks_LinkTopo) GetEntityData() *types.CommonEntityData {
    linkTopo.EntityData.YFilter = linkTopo.YFilter
    linkTopo.EntityData.YangName = "link-topo"
    linkTopo.EntityData.BundleName = "cisco_ios_xe"
    linkTopo.EntityData.ParentYangName = "ospfv2-router-lsa-links"
    linkTopo.EntityData.SegmentPath = "link-topo" + types.AddNoKeyToken(linkTopo)
    linkTopo.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospfv2-instance/ospfv2-area/ospfv2-lsdb-area/ospfv2-router-lsa-links/" + linkTopo.EntityData.SegmentPath
    linkTopo.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    linkTopo.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    linkTopo.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    linkTopo.EntityData.Children = types.NewOrderedMap()
    linkTopo.EntityData.Leafs = types.NewOrderedMap()
    linkTopo.EntityData.Leafs.Append("mt-id", types.YLeaf{"MtId", linkTopo.MtId})
    linkTopo.EntityData.Leafs.Append("topo-metric", types.YLeaf{"TopoMetric", linkTopo.TopoMetric})

    linkTopo.EntityData.YListKeys = []string {}

    return &(linkTopo.EntityData)
}

// OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2LsdbArea_UnsupportedLsa
// The unsupported Link State Advertisements
type OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2LsdbArea_UnsupportedLsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Link State Advertisement data. The type is slice of interface{} with range:
    // 0..255.
    LsaData []interface{}
}

func (unsupportedLsa *OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2LsdbArea_UnsupportedLsa) GetEntityData() *types.CommonEntityData {
    unsupportedLsa.EntityData.YFilter = unsupportedLsa.YFilter
    unsupportedLsa.EntityData.YangName = "unsupported-lsa"
    unsupportedLsa.EntityData.BundleName = "cisco_ios_xe"
    unsupportedLsa.EntityData.ParentYangName = "ospfv2-lsdb-area"
    unsupportedLsa.EntityData.SegmentPath = "unsupported-lsa"
    unsupportedLsa.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospfv2-instance/ospfv2-area/ospfv2-lsdb-area/" + unsupportedLsa.EntityData.SegmentPath
    unsupportedLsa.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    unsupportedLsa.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    unsupportedLsa.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    unsupportedLsa.EntityData.Children = types.NewOrderedMap()
    unsupportedLsa.EntityData.Leafs = types.NewOrderedMap()
    unsupportedLsa.EntityData.Leafs.Append("lsa-data", types.YLeaf{"LsaData", unsupportedLsa.LsaData})

    unsupportedLsa.EntityData.YListKeys = []string {}

    return &(unsupportedLsa.EntityData)
}

// OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2LsdbArea_RouterLsa
// The router Link State Advertisements
type OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2LsdbArea_RouterLsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Router Link State Advertisement bits. The type is map[string]bool.
    RouterLsaBits interface{}

    // Router Link State Advertisement number of links. The type is interface{}
    // with range: 0..65535.
    RouterLsaNumberLinks interface{}
}

func (routerLsa *OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2LsdbArea_RouterLsa) GetEntityData() *types.CommonEntityData {
    routerLsa.EntityData.YFilter = routerLsa.YFilter
    routerLsa.EntityData.YangName = "router-lsa"
    routerLsa.EntityData.BundleName = "cisco_ios_xe"
    routerLsa.EntityData.ParentYangName = "ospfv2-lsdb-area"
    routerLsa.EntityData.SegmentPath = "router-lsa"
    routerLsa.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospfv2-instance/ospfv2-area/ospfv2-lsdb-area/" + routerLsa.EntityData.SegmentPath
    routerLsa.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    routerLsa.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    routerLsa.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    routerLsa.EntityData.Children = types.NewOrderedMap()
    routerLsa.EntityData.Leafs = types.NewOrderedMap()
    routerLsa.EntityData.Leafs.Append("router-lsa-bits", types.YLeaf{"RouterLsaBits", routerLsa.RouterLsaBits})
    routerLsa.EntityData.Leafs.Append("router-lsa-number-links", types.YLeaf{"RouterLsaNumberLinks", routerLsa.RouterLsaNumberLinks})

    routerLsa.EntityData.YListKeys = []string {}

    return &(routerLsa.EntityData)
}

// OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2LsdbArea_NetworkLsa
// The network Link State Advertisements
type OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2LsdbArea_NetworkLsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Network Link State Advertisement mask. The type is interface{} with range:
    // 0..4294967295.
    NetworkLsaMask interface{}

    // Network attached routers. The type is slice of interface{} with range:
    // 0..4294967295.
    NetworkAttachedRouters []interface{}
}

func (networkLsa *OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2LsdbArea_NetworkLsa) GetEntityData() *types.CommonEntityData {
    networkLsa.EntityData.YFilter = networkLsa.YFilter
    networkLsa.EntityData.YangName = "network-lsa"
    networkLsa.EntityData.BundleName = "cisco_ios_xe"
    networkLsa.EntityData.ParentYangName = "ospfv2-lsdb-area"
    networkLsa.EntityData.SegmentPath = "network-lsa"
    networkLsa.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospfv2-instance/ospfv2-area/ospfv2-lsdb-area/" + networkLsa.EntityData.SegmentPath
    networkLsa.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    networkLsa.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    networkLsa.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    networkLsa.EntityData.Children = types.NewOrderedMap()
    networkLsa.EntityData.Leafs = types.NewOrderedMap()
    networkLsa.EntityData.Leafs.Append("network-lsa-mask", types.YLeaf{"NetworkLsaMask", networkLsa.NetworkLsaMask})
    networkLsa.EntityData.Leafs.Append("network-attached-routers", types.YLeaf{"NetworkAttachedRouters", networkLsa.NetworkAttachedRouters})

    networkLsa.EntityData.YListKeys = []string {}

    return &(networkLsa.EntityData)
}

// OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2LsdbArea_NetworkSummaryLsa
// The network summary Link State Advertisements
type OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2LsdbArea_NetworkSummaryLsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The summary Link State Advertisement mask. The type is interface{} with
    // range: 0..4294967295.
    SummaryLsaMask interface{}

    // The summary topology. The type is slice of
    // OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2LsdbArea_NetworkSummaryLsa_SummaryTopo.
    SummaryTopo []*OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2LsdbArea_NetworkSummaryLsa_SummaryTopo
}

func (networkSummaryLsa *OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2LsdbArea_NetworkSummaryLsa) GetEntityData() *types.CommonEntityData {
    networkSummaryLsa.EntityData.YFilter = networkSummaryLsa.YFilter
    networkSummaryLsa.EntityData.YangName = "network-summary-lsa"
    networkSummaryLsa.EntityData.BundleName = "cisco_ios_xe"
    networkSummaryLsa.EntityData.ParentYangName = "ospfv2-lsdb-area"
    networkSummaryLsa.EntityData.SegmentPath = "network-summary-lsa"
    networkSummaryLsa.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospfv2-instance/ospfv2-area/ospfv2-lsdb-area/" + networkSummaryLsa.EntityData.SegmentPath
    networkSummaryLsa.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    networkSummaryLsa.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    networkSummaryLsa.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    networkSummaryLsa.EntityData.Children = types.NewOrderedMap()
    networkSummaryLsa.EntityData.Children.Append("summary-topo", types.YChild{"SummaryTopo", nil})
    for i := range networkSummaryLsa.SummaryTopo {
        types.SetYListKey(networkSummaryLsa.SummaryTopo[i], i)
        networkSummaryLsa.EntityData.Children.Append(types.GetSegmentPath(networkSummaryLsa.SummaryTopo[i]), types.YChild{"SummaryTopo", networkSummaryLsa.SummaryTopo[i]})
    }
    networkSummaryLsa.EntityData.Leafs = types.NewOrderedMap()
    networkSummaryLsa.EntityData.Leafs.Append("summary-lsa-mask", types.YLeaf{"SummaryLsaMask", networkSummaryLsa.SummaryLsaMask})

    networkSummaryLsa.EntityData.YListKeys = []string {}

    return &(networkSummaryLsa.EntityData)
}

// OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2LsdbArea_NetworkSummaryLsa_SummaryTopo
// The summary topology
type OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2LsdbArea_NetworkSummaryLsa_SummaryTopo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Multi topology identifier. The type is interface{} with range: 0..255.
    MtId interface{}

    // Topology Metric. The type is interface{} with range: 0..4294967295.
    TopoMetric interface{}
}

func (summaryTopo *OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2LsdbArea_NetworkSummaryLsa_SummaryTopo) GetEntityData() *types.CommonEntityData {
    summaryTopo.EntityData.YFilter = summaryTopo.YFilter
    summaryTopo.EntityData.YangName = "summary-topo"
    summaryTopo.EntityData.BundleName = "cisco_ios_xe"
    summaryTopo.EntityData.ParentYangName = "network-summary-lsa"
    summaryTopo.EntityData.SegmentPath = "summary-topo" + types.AddNoKeyToken(summaryTopo)
    summaryTopo.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospfv2-instance/ospfv2-area/ospfv2-lsdb-area/network-summary-lsa/" + summaryTopo.EntityData.SegmentPath
    summaryTopo.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    summaryTopo.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    summaryTopo.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    summaryTopo.EntityData.Children = types.NewOrderedMap()
    summaryTopo.EntityData.Leafs = types.NewOrderedMap()
    summaryTopo.EntityData.Leafs.Append("mt-id", types.YLeaf{"MtId", summaryTopo.MtId})
    summaryTopo.EntityData.Leafs.Append("topo-metric", types.YLeaf{"TopoMetric", summaryTopo.TopoMetric})

    summaryTopo.EntityData.YListKeys = []string {}

    return &(summaryTopo.EntityData)
}

// OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2LsdbArea_RouterSummaryLsa
// The router summary Link State Advertisements
type OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2LsdbArea_RouterSummaryLsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The summary Link State Advertisement mask. The type is interface{} with
    // range: 0..4294967295.
    SummaryLsaMask interface{}

    // The summary topology. The type is slice of
    // OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2LsdbArea_RouterSummaryLsa_SummaryTopo.
    SummaryTopo []*OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2LsdbArea_RouterSummaryLsa_SummaryTopo
}

func (routerSummaryLsa *OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2LsdbArea_RouterSummaryLsa) GetEntityData() *types.CommonEntityData {
    routerSummaryLsa.EntityData.YFilter = routerSummaryLsa.YFilter
    routerSummaryLsa.EntityData.YangName = "router-summary-lsa"
    routerSummaryLsa.EntityData.BundleName = "cisco_ios_xe"
    routerSummaryLsa.EntityData.ParentYangName = "ospfv2-lsdb-area"
    routerSummaryLsa.EntityData.SegmentPath = "router-summary-lsa"
    routerSummaryLsa.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospfv2-instance/ospfv2-area/ospfv2-lsdb-area/" + routerSummaryLsa.EntityData.SegmentPath
    routerSummaryLsa.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    routerSummaryLsa.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    routerSummaryLsa.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    routerSummaryLsa.EntityData.Children = types.NewOrderedMap()
    routerSummaryLsa.EntityData.Children.Append("summary-topo", types.YChild{"SummaryTopo", nil})
    for i := range routerSummaryLsa.SummaryTopo {
        types.SetYListKey(routerSummaryLsa.SummaryTopo[i], i)
        routerSummaryLsa.EntityData.Children.Append(types.GetSegmentPath(routerSummaryLsa.SummaryTopo[i]), types.YChild{"SummaryTopo", routerSummaryLsa.SummaryTopo[i]})
    }
    routerSummaryLsa.EntityData.Leafs = types.NewOrderedMap()
    routerSummaryLsa.EntityData.Leafs.Append("summary-lsa-mask", types.YLeaf{"SummaryLsaMask", routerSummaryLsa.SummaryLsaMask})

    routerSummaryLsa.EntityData.YListKeys = []string {}

    return &(routerSummaryLsa.EntityData)
}

// OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2LsdbArea_RouterSummaryLsa_SummaryTopo
// The summary topology
type OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2LsdbArea_RouterSummaryLsa_SummaryTopo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Multi topology identifier. The type is interface{} with range: 0..255.
    MtId interface{}

    // Topology Metric. The type is interface{} with range: 0..4294967295.
    TopoMetric interface{}
}

func (summaryTopo *OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2LsdbArea_RouterSummaryLsa_SummaryTopo) GetEntityData() *types.CommonEntityData {
    summaryTopo.EntityData.YFilter = summaryTopo.YFilter
    summaryTopo.EntityData.YangName = "summary-topo"
    summaryTopo.EntityData.BundleName = "cisco_ios_xe"
    summaryTopo.EntityData.ParentYangName = "router-summary-lsa"
    summaryTopo.EntityData.SegmentPath = "summary-topo" + types.AddNoKeyToken(summaryTopo)
    summaryTopo.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospfv2-instance/ospfv2-area/ospfv2-lsdb-area/router-summary-lsa/" + summaryTopo.EntityData.SegmentPath
    summaryTopo.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    summaryTopo.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    summaryTopo.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    summaryTopo.EntityData.Children = types.NewOrderedMap()
    summaryTopo.EntityData.Leafs = types.NewOrderedMap()
    summaryTopo.EntityData.Leafs.Append("mt-id", types.YLeaf{"MtId", summaryTopo.MtId})
    summaryTopo.EntityData.Leafs.Append("topo-metric", types.YLeaf{"TopoMetric", summaryTopo.TopoMetric})

    summaryTopo.EntityData.YListKeys = []string {}

    return &(summaryTopo.EntityData)
}

// OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2LsdbArea_ExternalLsa
// The external Link State Advertisements
type OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2LsdbArea_ExternalLsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The mask for the external Link State Advertisement. The type is interface{}
    // with range: 0..4294967295.
    ExternalLsaMask interface{}

    // The external topology Link State Advertisement. The type is slice of
    // OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2LsdbArea_ExternalLsa_ExternalTopo.
    ExternalTopo []*OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2LsdbArea_ExternalLsa_ExternalTopo
}

func (externalLsa *OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2LsdbArea_ExternalLsa) GetEntityData() *types.CommonEntityData {
    externalLsa.EntityData.YFilter = externalLsa.YFilter
    externalLsa.EntityData.YangName = "external-lsa"
    externalLsa.EntityData.BundleName = "cisco_ios_xe"
    externalLsa.EntityData.ParentYangName = "ospfv2-lsdb-area"
    externalLsa.EntityData.SegmentPath = "external-lsa"
    externalLsa.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospfv2-instance/ospfv2-area/ospfv2-lsdb-area/" + externalLsa.EntityData.SegmentPath
    externalLsa.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    externalLsa.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    externalLsa.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    externalLsa.EntityData.Children = types.NewOrderedMap()
    externalLsa.EntityData.Children.Append("external-topo", types.YChild{"ExternalTopo", nil})
    for i := range externalLsa.ExternalTopo {
        types.SetYListKey(externalLsa.ExternalTopo[i], i)
        externalLsa.EntityData.Children.Append(types.GetSegmentPath(externalLsa.ExternalTopo[i]), types.YChild{"ExternalTopo", externalLsa.ExternalTopo[i]})
    }
    externalLsa.EntityData.Leafs = types.NewOrderedMap()
    externalLsa.EntityData.Leafs.Append("external-lsa-mask", types.YLeaf{"ExternalLsaMask", externalLsa.ExternalLsaMask})

    externalLsa.EntityData.YListKeys = []string {}

    return &(externalLsa.EntityData)
}

// OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2LsdbArea_ExternalLsa_ExternalTopo
// The external topology Link State Advertisement
type OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2LsdbArea_ExternalLsa_ExternalTopo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // The multi topology identifier. The type is interface{} with range: 0..255.
    MtId interface{}

    // The topoligy metric type associated with the  Link State Advertisement. The
    // type is OspfExternalMetricType.
    TopoMetricType interface{}

    // The topology metric. The type is interface{} with range: 0..4294967295.
    TopoMetric interface{}

    // The topology forwarding address. The type is one of the following types:
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    TopoForwardingAddress interface{}

    // The topology route tag. The type is interface{} with range: 0..4294967295.
    TopoRouteTag interface{}
}

func (externalTopo *OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2LsdbArea_ExternalLsa_ExternalTopo) GetEntityData() *types.CommonEntityData {
    externalTopo.EntityData.YFilter = externalTopo.YFilter
    externalTopo.EntityData.YangName = "external-topo"
    externalTopo.EntityData.BundleName = "cisco_ios_xe"
    externalTopo.EntityData.ParentYangName = "external-lsa"
    externalTopo.EntityData.SegmentPath = "external-topo" + types.AddNoKeyToken(externalTopo)
    externalTopo.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospfv2-instance/ospfv2-area/ospfv2-lsdb-area/external-lsa/" + externalTopo.EntityData.SegmentPath
    externalTopo.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    externalTopo.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    externalTopo.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    externalTopo.EntityData.Children = types.NewOrderedMap()
    externalTopo.EntityData.Leafs = types.NewOrderedMap()
    externalTopo.EntityData.Leafs.Append("mt-id", types.YLeaf{"MtId", externalTopo.MtId})
    externalTopo.EntityData.Leafs.Append("topo-metric-type", types.YLeaf{"TopoMetricType", externalTopo.TopoMetricType})
    externalTopo.EntityData.Leafs.Append("topo-metric", types.YLeaf{"TopoMetric", externalTopo.TopoMetric})
    externalTopo.EntityData.Leafs.Append("topo-forwarding-address", types.YLeaf{"TopoForwardingAddress", externalTopo.TopoForwardingAddress})
    externalTopo.EntityData.Leafs.Append("topo-route-tag", types.YLeaf{"TopoRouteTag", externalTopo.TopoRouteTag})

    externalTopo.EntityData.YListKeys = []string {}

    return &(externalTopo.EntityData)
}

// OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2LsdbArea_NssaLsa
// The Not So Stubby Area Link state advertisements
type OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2LsdbArea_NssaLsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The mask for the external Link State Advertisement. The type is interface{}
    // with range: 0..4294967295.
    ExternalLsaMask interface{}

    // The external topology Link State Advertisement. The type is slice of
    // OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2LsdbArea_NssaLsa_ExternalTopo.
    ExternalTopo []*OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2LsdbArea_NssaLsa_ExternalTopo
}

func (nssaLsa *OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2LsdbArea_NssaLsa) GetEntityData() *types.CommonEntityData {
    nssaLsa.EntityData.YFilter = nssaLsa.YFilter
    nssaLsa.EntityData.YangName = "nssa-lsa"
    nssaLsa.EntityData.BundleName = "cisco_ios_xe"
    nssaLsa.EntityData.ParentYangName = "ospfv2-lsdb-area"
    nssaLsa.EntityData.SegmentPath = "nssa-lsa"
    nssaLsa.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospfv2-instance/ospfv2-area/ospfv2-lsdb-area/" + nssaLsa.EntityData.SegmentPath
    nssaLsa.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nssaLsa.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nssaLsa.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nssaLsa.EntityData.Children = types.NewOrderedMap()
    nssaLsa.EntityData.Children.Append("external-topo", types.YChild{"ExternalTopo", nil})
    for i := range nssaLsa.ExternalTopo {
        types.SetYListKey(nssaLsa.ExternalTopo[i], i)
        nssaLsa.EntityData.Children.Append(types.GetSegmentPath(nssaLsa.ExternalTopo[i]), types.YChild{"ExternalTopo", nssaLsa.ExternalTopo[i]})
    }
    nssaLsa.EntityData.Leafs = types.NewOrderedMap()
    nssaLsa.EntityData.Leafs.Append("external-lsa-mask", types.YLeaf{"ExternalLsaMask", nssaLsa.ExternalLsaMask})

    nssaLsa.EntityData.YListKeys = []string {}

    return &(nssaLsa.EntityData)
}

// OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2LsdbArea_NssaLsa_ExternalTopo
// The external topology Link State Advertisement
type OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2LsdbArea_NssaLsa_ExternalTopo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // The multi topology identifier. The type is interface{} with range: 0..255.
    MtId interface{}

    // The topoligy metric type associated with the  Link State Advertisement. The
    // type is OspfExternalMetricType.
    TopoMetricType interface{}

    // The topology metric. The type is interface{} with range: 0..4294967295.
    TopoMetric interface{}

    // The topology forwarding address. The type is one of the following types:
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    TopoForwardingAddress interface{}

    // The topology route tag. The type is interface{} with range: 0..4294967295.
    TopoRouteTag interface{}
}

func (externalTopo *OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2LsdbArea_NssaLsa_ExternalTopo) GetEntityData() *types.CommonEntityData {
    externalTopo.EntityData.YFilter = externalTopo.YFilter
    externalTopo.EntityData.YangName = "external-topo"
    externalTopo.EntityData.BundleName = "cisco_ios_xe"
    externalTopo.EntityData.ParentYangName = "nssa-lsa"
    externalTopo.EntityData.SegmentPath = "external-topo" + types.AddNoKeyToken(externalTopo)
    externalTopo.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospfv2-instance/ospfv2-area/ospfv2-lsdb-area/nssa-lsa/" + externalTopo.EntityData.SegmentPath
    externalTopo.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    externalTopo.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    externalTopo.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    externalTopo.EntityData.Children = types.NewOrderedMap()
    externalTopo.EntityData.Leafs = types.NewOrderedMap()
    externalTopo.EntityData.Leafs.Append("mt-id", types.YLeaf{"MtId", externalTopo.MtId})
    externalTopo.EntityData.Leafs.Append("topo-metric-type", types.YLeaf{"TopoMetricType", externalTopo.TopoMetricType})
    externalTopo.EntityData.Leafs.Append("topo-metric", types.YLeaf{"TopoMetric", externalTopo.TopoMetric})
    externalTopo.EntityData.Leafs.Append("topo-forwarding-address", types.YLeaf{"TopoForwardingAddress", externalTopo.TopoForwardingAddress})
    externalTopo.EntityData.Leafs.Append("topo-route-tag", types.YLeaf{"TopoRouteTag", externalTopo.TopoRouteTag})

    externalTopo.EntityData.YListKeys = []string {}

    return &(externalTopo.EntityData)
}

// OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2Interface
// A list of interfaces that belong to the area
type OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Name of the interface. The type is string.
    Name interface{}

    // Network type. The type is OspfNetworkType.
    NetworkType interface{}

    // If the interface is enabled. The type is bool.
    Enable interface{}

    // If the interface is in passive mode. The type is bool.
    Passive interface{}

    // If this is a demand circuit. The type is bool.
    DemandCircuit interface{}

    // If the MTU is being ignored. The type is bool.
    MtuIgnore interface{}

    // If prefix suppression is enabled. The type is bool.
    PrefixSuppresion interface{}

    // The OSPFv2 cost. The type is interface{} with range: 0..65535.
    Cost interface{}

    // The hello interval in seconds. The type is interface{} with range:
    // 0..65535.
    HelloInterval interface{}

    // The dead interval in seconds. The type is interface{} with range: 0..65535.
    DeadInterval interface{}

    // The retransmit interval in seconds. The type is interface{} with range:
    // 0..65535.
    RetransmitInterval interface{}

    // The delay before transmitting a keepalive in seconds. The type is
    // interface{} with range: 0..65535.
    TransmitDelay interface{}

    // The current hello timer in seconds. The type is interface{} with range:
    // 0..4294967295.
    HelloTimer interface{}

    // The wait timer in seconds. The type is interface{} with range:
    // 0..4294967295.
    WaitTimer interface{}

    // The designated router identifier. The type is interface{} with range:
    // 0..4294967295.
    Dr interface{}

    // The backup designated router identifier. The type is interface{} with
    // range: 0..4294967295.
    Bdr interface{}

    // The address of the designated router. The type is one of the following
    // types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    DrIp interface{}

    // The address of the backup designated router. The type is one of the
    // following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    BdrIp interface{}

    // The current state of the interface. The type is Ospfv2IntfState.
    State interface{}

    // The TTL security information.
    TtlSecurityVal OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2Interface_TtlSecurityVal

    // The authentication information.
    AuthVal OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2Interface_AuthVal

    // All the neighbors on the interface. The type is slice of
    // OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2Interface_Ospfv2Neighbor.
    Ospfv2Neighbor []*OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2Interface_Ospfv2Neighbor
}

func (ospfv2Interface *OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2Interface) GetEntityData() *types.CommonEntityData {
    ospfv2Interface.EntityData.YFilter = ospfv2Interface.YFilter
    ospfv2Interface.EntityData.YangName = "ospfv2-interface"
    ospfv2Interface.EntityData.BundleName = "cisco_ios_xe"
    ospfv2Interface.EntityData.ParentYangName = "ospfv2-area"
    ospfv2Interface.EntityData.SegmentPath = "ospfv2-interface" + types.AddKeyToken(ospfv2Interface.Name, "name")
    ospfv2Interface.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospfv2-instance/ospfv2-area/" + ospfv2Interface.EntityData.SegmentPath
    ospfv2Interface.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv2Interface.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv2Interface.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv2Interface.EntityData.Children = types.NewOrderedMap()
    ospfv2Interface.EntityData.Children.Append("ttl-security-val", types.YChild{"TtlSecurityVal", &ospfv2Interface.TtlSecurityVal})
    ospfv2Interface.EntityData.Children.Append("auth-val", types.YChild{"AuthVal", &ospfv2Interface.AuthVal})
    ospfv2Interface.EntityData.Children.Append("ospfv2-neighbor", types.YChild{"Ospfv2Neighbor", nil})
    for i := range ospfv2Interface.Ospfv2Neighbor {
        ospfv2Interface.EntityData.Children.Append(types.GetSegmentPath(ospfv2Interface.Ospfv2Neighbor[i]), types.YChild{"Ospfv2Neighbor", ospfv2Interface.Ospfv2Neighbor[i]})
    }
    ospfv2Interface.EntityData.Leafs = types.NewOrderedMap()
    ospfv2Interface.EntityData.Leafs.Append("name", types.YLeaf{"Name", ospfv2Interface.Name})
    ospfv2Interface.EntityData.Leafs.Append("network-type", types.YLeaf{"NetworkType", ospfv2Interface.NetworkType})
    ospfv2Interface.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", ospfv2Interface.Enable})
    ospfv2Interface.EntityData.Leafs.Append("passive", types.YLeaf{"Passive", ospfv2Interface.Passive})
    ospfv2Interface.EntityData.Leafs.Append("demand-circuit", types.YLeaf{"DemandCircuit", ospfv2Interface.DemandCircuit})
    ospfv2Interface.EntityData.Leafs.Append("mtu-ignore", types.YLeaf{"MtuIgnore", ospfv2Interface.MtuIgnore})
    ospfv2Interface.EntityData.Leafs.Append("prefix-suppresion", types.YLeaf{"PrefixSuppresion", ospfv2Interface.PrefixSuppresion})
    ospfv2Interface.EntityData.Leafs.Append("cost", types.YLeaf{"Cost", ospfv2Interface.Cost})
    ospfv2Interface.EntityData.Leafs.Append("hello-interval", types.YLeaf{"HelloInterval", ospfv2Interface.HelloInterval})
    ospfv2Interface.EntityData.Leafs.Append("dead-interval", types.YLeaf{"DeadInterval", ospfv2Interface.DeadInterval})
    ospfv2Interface.EntityData.Leafs.Append("retransmit-interval", types.YLeaf{"RetransmitInterval", ospfv2Interface.RetransmitInterval})
    ospfv2Interface.EntityData.Leafs.Append("transmit-delay", types.YLeaf{"TransmitDelay", ospfv2Interface.TransmitDelay})
    ospfv2Interface.EntityData.Leafs.Append("hello-timer", types.YLeaf{"HelloTimer", ospfv2Interface.HelloTimer})
    ospfv2Interface.EntityData.Leafs.Append("wait-timer", types.YLeaf{"WaitTimer", ospfv2Interface.WaitTimer})
    ospfv2Interface.EntityData.Leafs.Append("dr", types.YLeaf{"Dr", ospfv2Interface.Dr})
    ospfv2Interface.EntityData.Leafs.Append("bdr", types.YLeaf{"Bdr", ospfv2Interface.Bdr})
    ospfv2Interface.EntityData.Leafs.Append("dr-ip", types.YLeaf{"DrIp", ospfv2Interface.DrIp})
    ospfv2Interface.EntityData.Leafs.Append("bdr-ip", types.YLeaf{"BdrIp", ospfv2Interface.BdrIp})
    ospfv2Interface.EntityData.Leafs.Append("state", types.YLeaf{"State", ospfv2Interface.State})

    ospfv2Interface.EntityData.YListKeys = []string {"Name"}

    return &(ospfv2Interface.EntityData)
}

// OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2Interface_TtlSecurityVal
// The TTL security information
type OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2Interface_TtlSecurityVal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Indicates whether time to live security is enabled. The type is bool.
    Enable interface{}

    // Number of hops for time to live security. The type is interface{} with
    // range: -2147483648..2147483647.
    Hops interface{}
}

func (ttlSecurityVal *OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2Interface_TtlSecurityVal) GetEntityData() *types.CommonEntityData {
    ttlSecurityVal.EntityData.YFilter = ttlSecurityVal.YFilter
    ttlSecurityVal.EntityData.YangName = "ttl-security-val"
    ttlSecurityVal.EntityData.BundleName = "cisco_ios_xe"
    ttlSecurityVal.EntityData.ParentYangName = "ospfv2-interface"
    ttlSecurityVal.EntityData.SegmentPath = "ttl-security-val"
    ttlSecurityVal.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospfv2-instance/ospfv2-area/ospfv2-interface/" + ttlSecurityVal.EntityData.SegmentPath
    ttlSecurityVal.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ttlSecurityVal.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ttlSecurityVal.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ttlSecurityVal.EntityData.Children = types.NewOrderedMap()
    ttlSecurityVal.EntityData.Leafs = types.NewOrderedMap()
    ttlSecurityVal.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", ttlSecurityVal.Enable})
    ttlSecurityVal.EntityData.Leafs.Append("hops", types.YLeaf{"Hops", ttlSecurityVal.Hops})

    ttlSecurityVal.EntityData.YListKeys = []string {}

    return &(ttlSecurityVal.EntityData)
}

// OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2Interface_AuthVal
// The authentication information
type OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2Interface_AuthVal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // No authentication in use. The type is interface{} with range:
    // 0..4294967295.
    NoAuth interface{}

    // Trailer key chain information.
    AuthKey OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2Interface_AuthVal_AuthKey

    // Trailer key information.
    KeyChain OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2Interface_AuthVal_KeyChain
}

func (authVal *OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2Interface_AuthVal) GetEntityData() *types.CommonEntityData {
    authVal.EntityData.YFilter = authVal.YFilter
    authVal.EntityData.YangName = "auth-val"
    authVal.EntityData.BundleName = "cisco_ios_xe"
    authVal.EntityData.ParentYangName = "ospfv2-interface"
    authVal.EntityData.SegmentPath = "auth-val"
    authVal.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospfv2-instance/ospfv2-area/ospfv2-interface/" + authVal.EntityData.SegmentPath
    authVal.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    authVal.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    authVal.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    authVal.EntityData.Children = types.NewOrderedMap()
    authVal.EntityData.Children.Append("auth-key", types.YChild{"AuthKey", &authVal.AuthKey})
    authVal.EntityData.Children.Append("key-chain", types.YChild{"KeyChain", &authVal.KeyChain})
    authVal.EntityData.Leafs = types.NewOrderedMap()
    authVal.EntityData.Leafs.Append("no-auth", types.YLeaf{"NoAuth", authVal.NoAuth})

    authVal.EntityData.YListKeys = []string {}

    return &(authVal.EntityData)
}

// OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2Interface_AuthVal_AuthKey
// Trailer key chain information
type OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2Interface_AuthVal_AuthKey struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The key identifier. The type is interface{} with range: 0..255.
    KeyId interface{}

    // The key string. The type is slice of interface{} with range: 0..255.
    KeyString []interface{}

    // The algorithm in use. The type is Ospfv2CryptoAlgorithm.
    CryptoAlgo interface{}
}

func (authKey *OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2Interface_AuthVal_AuthKey) GetEntityData() *types.CommonEntityData {
    authKey.EntityData.YFilter = authKey.YFilter
    authKey.EntityData.YangName = "auth-key"
    authKey.EntityData.BundleName = "cisco_ios_xe"
    authKey.EntityData.ParentYangName = "auth-val"
    authKey.EntityData.SegmentPath = "auth-key"
    authKey.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospfv2-instance/ospfv2-area/ospfv2-interface/auth-val/" + authKey.EntityData.SegmentPath
    authKey.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    authKey.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    authKey.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    authKey.EntityData.Children = types.NewOrderedMap()
    authKey.EntityData.Leafs = types.NewOrderedMap()
    authKey.EntityData.Leafs.Append("key-id", types.YLeaf{"KeyId", authKey.KeyId})
    authKey.EntityData.Leafs.Append("key-string", types.YLeaf{"KeyString", authKey.KeyString})
    authKey.EntityData.Leafs.Append("crypto-algo", types.YLeaf{"CryptoAlgo", authKey.CryptoAlgo})

    authKey.EntityData.YListKeys = []string {}

    return &(authKey.EntityData)
}

// OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2Interface_AuthVal_KeyChain
// Trailer key information
type OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2Interface_AuthVal_KeyChain struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The key chain. The type is slice of interface{} with range: 0..255.
    KeyChain []interface{}
}

func (keyChain *OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2Interface_AuthVal_KeyChain) GetEntityData() *types.CommonEntityData {
    keyChain.EntityData.YFilter = keyChain.YFilter
    keyChain.EntityData.YangName = "key-chain"
    keyChain.EntityData.BundleName = "cisco_ios_xe"
    keyChain.EntityData.ParentYangName = "auth-val"
    keyChain.EntityData.SegmentPath = "key-chain"
    keyChain.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospfv2-instance/ospfv2-area/ospfv2-interface/auth-val/" + keyChain.EntityData.SegmentPath
    keyChain.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    keyChain.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    keyChain.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    keyChain.EntityData.Children = types.NewOrderedMap()
    keyChain.EntityData.Leafs = types.NewOrderedMap()
    keyChain.EntityData.Leafs.Append("key-chain", types.YLeaf{"KeyChain", keyChain.KeyChain})

    keyChain.EntityData.YListKeys = []string {}

    return &(keyChain.EntityData)
}

// OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2Interface_Ospfv2Neighbor
// All the neighbors on the interface
type OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2Interface_Ospfv2Neighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The neighbor identifier. The type is interface{}
    // with range: 0..4294967295.
    NbrId interface{}

    // Neighbor address. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // The neighbor's Designated Router indentifier . The type is interface{} with
    // range: 0..4294967295.
    Dr interface{}

    // The neighbor's Backup Designated Router identifier. The type is interface{}
    // with range: 0..4294967295.
    Bdr interface{}

    // The designated routers' IP address. The type is one of the following types:
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    DrIp interface{}

    // The backup designated routers' IP address. The type is one of the following
    // types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    BdrIp interface{}

    // A count of neighbor events. The type is interface{} with range:
    // 0..4294967295.
    EventCount interface{}

    // A count of the retransmission events. The type is interface{} with range:
    // 0..4294967295.
    RetransCount interface{}

    // The current neighbor state. The type is NbrStateType.
    State interface{}

    // The dead timer in seconds. The type is interface{} with range:
    // 0..4294967295.
    DeadTimer interface{}
}

func (ospfv2Neighbor *OspfOperData_Ospfv2Instance_Ospfv2Area_Ospfv2Interface_Ospfv2Neighbor) GetEntityData() *types.CommonEntityData {
    ospfv2Neighbor.EntityData.YFilter = ospfv2Neighbor.YFilter
    ospfv2Neighbor.EntityData.YangName = "ospfv2-neighbor"
    ospfv2Neighbor.EntityData.BundleName = "cisco_ios_xe"
    ospfv2Neighbor.EntityData.ParentYangName = "ospfv2-interface"
    ospfv2Neighbor.EntityData.SegmentPath = "ospfv2-neighbor" + types.AddKeyToken(ospfv2Neighbor.NbrId, "nbr-id")
    ospfv2Neighbor.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospfv2-instance/ospfv2-area/ospfv2-interface/" + ospfv2Neighbor.EntityData.SegmentPath
    ospfv2Neighbor.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv2Neighbor.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv2Neighbor.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv2Neighbor.EntityData.Children = types.NewOrderedMap()
    ospfv2Neighbor.EntityData.Leafs = types.NewOrderedMap()
    ospfv2Neighbor.EntityData.Leafs.Append("nbr-id", types.YLeaf{"NbrId", ospfv2Neighbor.NbrId})
    ospfv2Neighbor.EntityData.Leafs.Append("address", types.YLeaf{"Address", ospfv2Neighbor.Address})
    ospfv2Neighbor.EntityData.Leafs.Append("dr", types.YLeaf{"Dr", ospfv2Neighbor.Dr})
    ospfv2Neighbor.EntityData.Leafs.Append("bdr", types.YLeaf{"Bdr", ospfv2Neighbor.Bdr})
    ospfv2Neighbor.EntityData.Leafs.Append("dr-ip", types.YLeaf{"DrIp", ospfv2Neighbor.DrIp})
    ospfv2Neighbor.EntityData.Leafs.Append("bdr-ip", types.YLeaf{"BdrIp", ospfv2Neighbor.BdrIp})
    ospfv2Neighbor.EntityData.Leafs.Append("event-count", types.YLeaf{"EventCount", ospfv2Neighbor.EventCount})
    ospfv2Neighbor.EntityData.Leafs.Append("retrans-count", types.YLeaf{"RetransCount", ospfv2Neighbor.RetransCount})
    ospfv2Neighbor.EntityData.Leafs.Append("state", types.YLeaf{"State", ospfv2Neighbor.State})
    ospfv2Neighbor.EntityData.Leafs.Append("dead-timer", types.YLeaf{"DeadTimer", ospfv2Neighbor.DeadTimer})

    ospfv2Neighbor.EntityData.YListKeys = []string {"NbrId"}

    return &(ospfv2Neighbor.EntityData)
}

// OspfOperData_Ospfv2Instance_Ospfv2LsdbExternal
// The external LSDB information
type OspfOperData_Ospfv2Instance_Ospfv2LsdbExternal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Link State Advertisement type. The type is
    // interface{} with range: 0..255.
    LsaType interface{}

    // This attribute is a key. Link State Advertisement Identifer. The type is
    // interface{} with range: 0..4294967295.
    LsaId interface{}

    // This attribute is a key. Advertising router. The type is interface{} with
    // range: 0..4294967295.
    AdvertisingRouter interface{}

    // The age of the Link State Advertisement. The type is interface{} with
    // range: 0..65535.
    LsaAge interface{}

    // The options of the Link State Advertisement. The type is map[string]bool.
    LsaOptions interface{}

    // The sequence number for the Link State Advertisement. The type is
    // interface{} with range: 0..4294967295.
    LsaSeqNumber interface{}

    // The checksum of the Link State Advertisement. The type is interface{} with
    // range: 0..65535.
    LsaChecksum interface{}

    // The length, in bytes, of the Link State Advertisement. The type is
    // interface{} with range: 0..65535.
    LsaLength interface{}

    // The router Link State Advertisement links. The type is slice of
    // OspfOperData_Ospfv2Instance_Ospfv2LsdbExternal_Ospfv2RouterLsaLinks.
    Ospfv2RouterLsaLinks []*OspfOperData_Ospfv2Instance_Ospfv2LsdbExternal_Ospfv2RouterLsaLinks

    // The unsupported Link State Advertisements.
    UnsupportedLsa OspfOperData_Ospfv2Instance_Ospfv2LsdbExternal_UnsupportedLsa

    // The router Link State Advertisements.
    RouterLsa OspfOperData_Ospfv2Instance_Ospfv2LsdbExternal_RouterLsa

    // The network Link State Advertisements.
    NetworkLsa OspfOperData_Ospfv2Instance_Ospfv2LsdbExternal_NetworkLsa

    // The network summary Link State Advertisements.
    NetworkSummaryLsa OspfOperData_Ospfv2Instance_Ospfv2LsdbExternal_NetworkSummaryLsa

    // The router summary Link State Advertisements.
    RouterSummaryLsa OspfOperData_Ospfv2Instance_Ospfv2LsdbExternal_RouterSummaryLsa

    // The external Link State Advertisements.
    ExternalLsa OspfOperData_Ospfv2Instance_Ospfv2LsdbExternal_ExternalLsa

    // The Not So Stubby Area Link state advertisements.
    NssaLsa OspfOperData_Ospfv2Instance_Ospfv2LsdbExternal_NssaLsa
}

func (ospfv2LsdbExternal *OspfOperData_Ospfv2Instance_Ospfv2LsdbExternal) GetEntityData() *types.CommonEntityData {
    ospfv2LsdbExternal.EntityData.YFilter = ospfv2LsdbExternal.YFilter
    ospfv2LsdbExternal.EntityData.YangName = "ospfv2-lsdb-external"
    ospfv2LsdbExternal.EntityData.BundleName = "cisco_ios_xe"
    ospfv2LsdbExternal.EntityData.ParentYangName = "ospfv2-instance"
    ospfv2LsdbExternal.EntityData.SegmentPath = "ospfv2-lsdb-external" + types.AddKeyToken(ospfv2LsdbExternal.LsaType, "lsa-type") + types.AddKeyToken(ospfv2LsdbExternal.LsaId, "lsa-id") + types.AddKeyToken(ospfv2LsdbExternal.AdvertisingRouter, "advertising-router")
    ospfv2LsdbExternal.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospfv2-instance/" + ospfv2LsdbExternal.EntityData.SegmentPath
    ospfv2LsdbExternal.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv2LsdbExternal.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv2LsdbExternal.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv2LsdbExternal.EntityData.Children = types.NewOrderedMap()
    ospfv2LsdbExternal.EntityData.Children.Append("ospfv2-router-lsa-links", types.YChild{"Ospfv2RouterLsaLinks", nil})
    for i := range ospfv2LsdbExternal.Ospfv2RouterLsaLinks {
        ospfv2LsdbExternal.EntityData.Children.Append(types.GetSegmentPath(ospfv2LsdbExternal.Ospfv2RouterLsaLinks[i]), types.YChild{"Ospfv2RouterLsaLinks", ospfv2LsdbExternal.Ospfv2RouterLsaLinks[i]})
    }
    ospfv2LsdbExternal.EntityData.Children.Append("unsupported-lsa", types.YChild{"UnsupportedLsa", &ospfv2LsdbExternal.UnsupportedLsa})
    ospfv2LsdbExternal.EntityData.Children.Append("router-lsa", types.YChild{"RouterLsa", &ospfv2LsdbExternal.RouterLsa})
    ospfv2LsdbExternal.EntityData.Children.Append("network-lsa", types.YChild{"NetworkLsa", &ospfv2LsdbExternal.NetworkLsa})
    ospfv2LsdbExternal.EntityData.Children.Append("network-summary-lsa", types.YChild{"NetworkSummaryLsa", &ospfv2LsdbExternal.NetworkSummaryLsa})
    ospfv2LsdbExternal.EntityData.Children.Append("router-summary-lsa", types.YChild{"RouterSummaryLsa", &ospfv2LsdbExternal.RouterSummaryLsa})
    ospfv2LsdbExternal.EntityData.Children.Append("external-lsa", types.YChild{"ExternalLsa", &ospfv2LsdbExternal.ExternalLsa})
    ospfv2LsdbExternal.EntityData.Children.Append("nssa-lsa", types.YChild{"NssaLsa", &ospfv2LsdbExternal.NssaLsa})
    ospfv2LsdbExternal.EntityData.Leafs = types.NewOrderedMap()
    ospfv2LsdbExternal.EntityData.Leafs.Append("lsa-type", types.YLeaf{"LsaType", ospfv2LsdbExternal.LsaType})
    ospfv2LsdbExternal.EntityData.Leafs.Append("lsa-id", types.YLeaf{"LsaId", ospfv2LsdbExternal.LsaId})
    ospfv2LsdbExternal.EntityData.Leafs.Append("advertising-router", types.YLeaf{"AdvertisingRouter", ospfv2LsdbExternal.AdvertisingRouter})
    ospfv2LsdbExternal.EntityData.Leafs.Append("lsa-age", types.YLeaf{"LsaAge", ospfv2LsdbExternal.LsaAge})
    ospfv2LsdbExternal.EntityData.Leafs.Append("lsa-options", types.YLeaf{"LsaOptions", ospfv2LsdbExternal.LsaOptions})
    ospfv2LsdbExternal.EntityData.Leafs.Append("lsa-seq-number", types.YLeaf{"LsaSeqNumber", ospfv2LsdbExternal.LsaSeqNumber})
    ospfv2LsdbExternal.EntityData.Leafs.Append("lsa-checksum", types.YLeaf{"LsaChecksum", ospfv2LsdbExternal.LsaChecksum})
    ospfv2LsdbExternal.EntityData.Leafs.Append("lsa-length", types.YLeaf{"LsaLength", ospfv2LsdbExternal.LsaLength})

    ospfv2LsdbExternal.EntityData.YListKeys = []string {"LsaType", "LsaId", "AdvertisingRouter"}

    return &(ospfv2LsdbExternal.EntityData)
}

// OspfOperData_Ospfv2Instance_Ospfv2LsdbExternal_Ospfv2RouterLsaLinks
// The router Link State Advertisement links
type OspfOperData_Ospfv2Instance_Ospfv2LsdbExternal_Ospfv2RouterLsaLinks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Link Type. The type is interface{} with range:
    // 0..255.
    LinkType interface{}

    // This attribute is a key. link Identifier. The type is interface{} with
    // range: 0..4294967295.
    LinkId interface{}

    // This attribute is a key. link data. The type is interface{} with range:
    // 0..4294967295.
    LinkData interface{}

    // Link topology. The type is slice of
    // OspfOperData_Ospfv2Instance_Ospfv2LsdbExternal_Ospfv2RouterLsaLinks_LinkTopo.
    LinkTopo []*OspfOperData_Ospfv2Instance_Ospfv2LsdbExternal_Ospfv2RouterLsaLinks_LinkTopo
}

func (ospfv2RouterLsaLinks *OspfOperData_Ospfv2Instance_Ospfv2LsdbExternal_Ospfv2RouterLsaLinks) GetEntityData() *types.CommonEntityData {
    ospfv2RouterLsaLinks.EntityData.YFilter = ospfv2RouterLsaLinks.YFilter
    ospfv2RouterLsaLinks.EntityData.YangName = "ospfv2-router-lsa-links"
    ospfv2RouterLsaLinks.EntityData.BundleName = "cisco_ios_xe"
    ospfv2RouterLsaLinks.EntityData.ParentYangName = "ospfv2-lsdb-external"
    ospfv2RouterLsaLinks.EntityData.SegmentPath = "ospfv2-router-lsa-links" + types.AddKeyToken(ospfv2RouterLsaLinks.LinkType, "link-type") + types.AddKeyToken(ospfv2RouterLsaLinks.LinkId, "link-id") + types.AddKeyToken(ospfv2RouterLsaLinks.LinkData, "link-data")
    ospfv2RouterLsaLinks.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospfv2-instance/ospfv2-lsdb-external/" + ospfv2RouterLsaLinks.EntityData.SegmentPath
    ospfv2RouterLsaLinks.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv2RouterLsaLinks.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv2RouterLsaLinks.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv2RouterLsaLinks.EntityData.Children = types.NewOrderedMap()
    ospfv2RouterLsaLinks.EntityData.Children.Append("link-topo", types.YChild{"LinkTopo", nil})
    for i := range ospfv2RouterLsaLinks.LinkTopo {
        types.SetYListKey(ospfv2RouterLsaLinks.LinkTopo[i], i)
        ospfv2RouterLsaLinks.EntityData.Children.Append(types.GetSegmentPath(ospfv2RouterLsaLinks.LinkTopo[i]), types.YChild{"LinkTopo", ospfv2RouterLsaLinks.LinkTopo[i]})
    }
    ospfv2RouterLsaLinks.EntityData.Leafs = types.NewOrderedMap()
    ospfv2RouterLsaLinks.EntityData.Leafs.Append("link-type", types.YLeaf{"LinkType", ospfv2RouterLsaLinks.LinkType})
    ospfv2RouterLsaLinks.EntityData.Leafs.Append("link-id", types.YLeaf{"LinkId", ospfv2RouterLsaLinks.LinkId})
    ospfv2RouterLsaLinks.EntityData.Leafs.Append("link-data", types.YLeaf{"LinkData", ospfv2RouterLsaLinks.LinkData})

    ospfv2RouterLsaLinks.EntityData.YListKeys = []string {"LinkType", "LinkId", "LinkData"}

    return &(ospfv2RouterLsaLinks.EntityData)
}

// OspfOperData_Ospfv2Instance_Ospfv2LsdbExternal_Ospfv2RouterLsaLinks_LinkTopo
// Link topology
type OspfOperData_Ospfv2Instance_Ospfv2LsdbExternal_Ospfv2RouterLsaLinks_LinkTopo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Multi topology identifier. The type is interface{} with range: 0..255.
    MtId interface{}

    // Topology metric. The type is interface{} with range: 0..65535.
    TopoMetric interface{}
}

func (linkTopo *OspfOperData_Ospfv2Instance_Ospfv2LsdbExternal_Ospfv2RouterLsaLinks_LinkTopo) GetEntityData() *types.CommonEntityData {
    linkTopo.EntityData.YFilter = linkTopo.YFilter
    linkTopo.EntityData.YangName = "link-topo"
    linkTopo.EntityData.BundleName = "cisco_ios_xe"
    linkTopo.EntityData.ParentYangName = "ospfv2-router-lsa-links"
    linkTopo.EntityData.SegmentPath = "link-topo" + types.AddNoKeyToken(linkTopo)
    linkTopo.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospfv2-instance/ospfv2-lsdb-external/ospfv2-router-lsa-links/" + linkTopo.EntityData.SegmentPath
    linkTopo.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    linkTopo.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    linkTopo.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    linkTopo.EntityData.Children = types.NewOrderedMap()
    linkTopo.EntityData.Leafs = types.NewOrderedMap()
    linkTopo.EntityData.Leafs.Append("mt-id", types.YLeaf{"MtId", linkTopo.MtId})
    linkTopo.EntityData.Leafs.Append("topo-metric", types.YLeaf{"TopoMetric", linkTopo.TopoMetric})

    linkTopo.EntityData.YListKeys = []string {}

    return &(linkTopo.EntityData)
}

// OspfOperData_Ospfv2Instance_Ospfv2LsdbExternal_UnsupportedLsa
// The unsupported Link State Advertisements
type OspfOperData_Ospfv2Instance_Ospfv2LsdbExternal_UnsupportedLsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Link State Advertisement data. The type is slice of interface{} with range:
    // 0..255.
    LsaData []interface{}
}

func (unsupportedLsa *OspfOperData_Ospfv2Instance_Ospfv2LsdbExternal_UnsupportedLsa) GetEntityData() *types.CommonEntityData {
    unsupportedLsa.EntityData.YFilter = unsupportedLsa.YFilter
    unsupportedLsa.EntityData.YangName = "unsupported-lsa"
    unsupportedLsa.EntityData.BundleName = "cisco_ios_xe"
    unsupportedLsa.EntityData.ParentYangName = "ospfv2-lsdb-external"
    unsupportedLsa.EntityData.SegmentPath = "unsupported-lsa"
    unsupportedLsa.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospfv2-instance/ospfv2-lsdb-external/" + unsupportedLsa.EntityData.SegmentPath
    unsupportedLsa.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    unsupportedLsa.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    unsupportedLsa.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    unsupportedLsa.EntityData.Children = types.NewOrderedMap()
    unsupportedLsa.EntityData.Leafs = types.NewOrderedMap()
    unsupportedLsa.EntityData.Leafs.Append("lsa-data", types.YLeaf{"LsaData", unsupportedLsa.LsaData})

    unsupportedLsa.EntityData.YListKeys = []string {}

    return &(unsupportedLsa.EntityData)
}

// OspfOperData_Ospfv2Instance_Ospfv2LsdbExternal_RouterLsa
// The router Link State Advertisements
type OspfOperData_Ospfv2Instance_Ospfv2LsdbExternal_RouterLsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Router Link State Advertisement bits. The type is map[string]bool.
    RouterLsaBits interface{}

    // Router Link State Advertisement number of links. The type is interface{}
    // with range: 0..65535.
    RouterLsaNumberLinks interface{}
}

func (routerLsa *OspfOperData_Ospfv2Instance_Ospfv2LsdbExternal_RouterLsa) GetEntityData() *types.CommonEntityData {
    routerLsa.EntityData.YFilter = routerLsa.YFilter
    routerLsa.EntityData.YangName = "router-lsa"
    routerLsa.EntityData.BundleName = "cisco_ios_xe"
    routerLsa.EntityData.ParentYangName = "ospfv2-lsdb-external"
    routerLsa.EntityData.SegmentPath = "router-lsa"
    routerLsa.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospfv2-instance/ospfv2-lsdb-external/" + routerLsa.EntityData.SegmentPath
    routerLsa.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    routerLsa.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    routerLsa.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    routerLsa.EntityData.Children = types.NewOrderedMap()
    routerLsa.EntityData.Leafs = types.NewOrderedMap()
    routerLsa.EntityData.Leafs.Append("router-lsa-bits", types.YLeaf{"RouterLsaBits", routerLsa.RouterLsaBits})
    routerLsa.EntityData.Leafs.Append("router-lsa-number-links", types.YLeaf{"RouterLsaNumberLinks", routerLsa.RouterLsaNumberLinks})

    routerLsa.EntityData.YListKeys = []string {}

    return &(routerLsa.EntityData)
}

// OspfOperData_Ospfv2Instance_Ospfv2LsdbExternal_NetworkLsa
// The network Link State Advertisements
type OspfOperData_Ospfv2Instance_Ospfv2LsdbExternal_NetworkLsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Network Link State Advertisement mask. The type is interface{} with range:
    // 0..4294967295.
    NetworkLsaMask interface{}

    // Network attached routers. The type is slice of interface{} with range:
    // 0..4294967295.
    NetworkAttachedRouters []interface{}
}

func (networkLsa *OspfOperData_Ospfv2Instance_Ospfv2LsdbExternal_NetworkLsa) GetEntityData() *types.CommonEntityData {
    networkLsa.EntityData.YFilter = networkLsa.YFilter
    networkLsa.EntityData.YangName = "network-lsa"
    networkLsa.EntityData.BundleName = "cisco_ios_xe"
    networkLsa.EntityData.ParentYangName = "ospfv2-lsdb-external"
    networkLsa.EntityData.SegmentPath = "network-lsa"
    networkLsa.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospfv2-instance/ospfv2-lsdb-external/" + networkLsa.EntityData.SegmentPath
    networkLsa.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    networkLsa.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    networkLsa.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    networkLsa.EntityData.Children = types.NewOrderedMap()
    networkLsa.EntityData.Leafs = types.NewOrderedMap()
    networkLsa.EntityData.Leafs.Append("network-lsa-mask", types.YLeaf{"NetworkLsaMask", networkLsa.NetworkLsaMask})
    networkLsa.EntityData.Leafs.Append("network-attached-routers", types.YLeaf{"NetworkAttachedRouters", networkLsa.NetworkAttachedRouters})

    networkLsa.EntityData.YListKeys = []string {}

    return &(networkLsa.EntityData)
}

// OspfOperData_Ospfv2Instance_Ospfv2LsdbExternal_NetworkSummaryLsa
// The network summary Link State Advertisements
type OspfOperData_Ospfv2Instance_Ospfv2LsdbExternal_NetworkSummaryLsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The summary Link State Advertisement mask. The type is interface{} with
    // range: 0..4294967295.
    SummaryLsaMask interface{}

    // The summary topology. The type is slice of
    // OspfOperData_Ospfv2Instance_Ospfv2LsdbExternal_NetworkSummaryLsa_SummaryTopo.
    SummaryTopo []*OspfOperData_Ospfv2Instance_Ospfv2LsdbExternal_NetworkSummaryLsa_SummaryTopo
}

func (networkSummaryLsa *OspfOperData_Ospfv2Instance_Ospfv2LsdbExternal_NetworkSummaryLsa) GetEntityData() *types.CommonEntityData {
    networkSummaryLsa.EntityData.YFilter = networkSummaryLsa.YFilter
    networkSummaryLsa.EntityData.YangName = "network-summary-lsa"
    networkSummaryLsa.EntityData.BundleName = "cisco_ios_xe"
    networkSummaryLsa.EntityData.ParentYangName = "ospfv2-lsdb-external"
    networkSummaryLsa.EntityData.SegmentPath = "network-summary-lsa"
    networkSummaryLsa.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospfv2-instance/ospfv2-lsdb-external/" + networkSummaryLsa.EntityData.SegmentPath
    networkSummaryLsa.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    networkSummaryLsa.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    networkSummaryLsa.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    networkSummaryLsa.EntityData.Children = types.NewOrderedMap()
    networkSummaryLsa.EntityData.Children.Append("summary-topo", types.YChild{"SummaryTopo", nil})
    for i := range networkSummaryLsa.SummaryTopo {
        types.SetYListKey(networkSummaryLsa.SummaryTopo[i], i)
        networkSummaryLsa.EntityData.Children.Append(types.GetSegmentPath(networkSummaryLsa.SummaryTopo[i]), types.YChild{"SummaryTopo", networkSummaryLsa.SummaryTopo[i]})
    }
    networkSummaryLsa.EntityData.Leafs = types.NewOrderedMap()
    networkSummaryLsa.EntityData.Leafs.Append("summary-lsa-mask", types.YLeaf{"SummaryLsaMask", networkSummaryLsa.SummaryLsaMask})

    networkSummaryLsa.EntityData.YListKeys = []string {}

    return &(networkSummaryLsa.EntityData)
}

// OspfOperData_Ospfv2Instance_Ospfv2LsdbExternal_NetworkSummaryLsa_SummaryTopo
// The summary topology
type OspfOperData_Ospfv2Instance_Ospfv2LsdbExternal_NetworkSummaryLsa_SummaryTopo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Multi topology identifier. The type is interface{} with range: 0..255.
    MtId interface{}

    // Topology Metric. The type is interface{} with range: 0..4294967295.
    TopoMetric interface{}
}

func (summaryTopo *OspfOperData_Ospfv2Instance_Ospfv2LsdbExternal_NetworkSummaryLsa_SummaryTopo) GetEntityData() *types.CommonEntityData {
    summaryTopo.EntityData.YFilter = summaryTopo.YFilter
    summaryTopo.EntityData.YangName = "summary-topo"
    summaryTopo.EntityData.BundleName = "cisco_ios_xe"
    summaryTopo.EntityData.ParentYangName = "network-summary-lsa"
    summaryTopo.EntityData.SegmentPath = "summary-topo" + types.AddNoKeyToken(summaryTopo)
    summaryTopo.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospfv2-instance/ospfv2-lsdb-external/network-summary-lsa/" + summaryTopo.EntityData.SegmentPath
    summaryTopo.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    summaryTopo.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    summaryTopo.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    summaryTopo.EntityData.Children = types.NewOrderedMap()
    summaryTopo.EntityData.Leafs = types.NewOrderedMap()
    summaryTopo.EntityData.Leafs.Append("mt-id", types.YLeaf{"MtId", summaryTopo.MtId})
    summaryTopo.EntityData.Leafs.Append("topo-metric", types.YLeaf{"TopoMetric", summaryTopo.TopoMetric})

    summaryTopo.EntityData.YListKeys = []string {}

    return &(summaryTopo.EntityData)
}

// OspfOperData_Ospfv2Instance_Ospfv2LsdbExternal_RouterSummaryLsa
// The router summary Link State Advertisements
type OspfOperData_Ospfv2Instance_Ospfv2LsdbExternal_RouterSummaryLsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The summary Link State Advertisement mask. The type is interface{} with
    // range: 0..4294967295.
    SummaryLsaMask interface{}

    // The summary topology. The type is slice of
    // OspfOperData_Ospfv2Instance_Ospfv2LsdbExternal_RouterSummaryLsa_SummaryTopo.
    SummaryTopo []*OspfOperData_Ospfv2Instance_Ospfv2LsdbExternal_RouterSummaryLsa_SummaryTopo
}

func (routerSummaryLsa *OspfOperData_Ospfv2Instance_Ospfv2LsdbExternal_RouterSummaryLsa) GetEntityData() *types.CommonEntityData {
    routerSummaryLsa.EntityData.YFilter = routerSummaryLsa.YFilter
    routerSummaryLsa.EntityData.YangName = "router-summary-lsa"
    routerSummaryLsa.EntityData.BundleName = "cisco_ios_xe"
    routerSummaryLsa.EntityData.ParentYangName = "ospfv2-lsdb-external"
    routerSummaryLsa.EntityData.SegmentPath = "router-summary-lsa"
    routerSummaryLsa.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospfv2-instance/ospfv2-lsdb-external/" + routerSummaryLsa.EntityData.SegmentPath
    routerSummaryLsa.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    routerSummaryLsa.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    routerSummaryLsa.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    routerSummaryLsa.EntityData.Children = types.NewOrderedMap()
    routerSummaryLsa.EntityData.Children.Append("summary-topo", types.YChild{"SummaryTopo", nil})
    for i := range routerSummaryLsa.SummaryTopo {
        types.SetYListKey(routerSummaryLsa.SummaryTopo[i], i)
        routerSummaryLsa.EntityData.Children.Append(types.GetSegmentPath(routerSummaryLsa.SummaryTopo[i]), types.YChild{"SummaryTopo", routerSummaryLsa.SummaryTopo[i]})
    }
    routerSummaryLsa.EntityData.Leafs = types.NewOrderedMap()
    routerSummaryLsa.EntityData.Leafs.Append("summary-lsa-mask", types.YLeaf{"SummaryLsaMask", routerSummaryLsa.SummaryLsaMask})

    routerSummaryLsa.EntityData.YListKeys = []string {}

    return &(routerSummaryLsa.EntityData)
}

// OspfOperData_Ospfv2Instance_Ospfv2LsdbExternal_RouterSummaryLsa_SummaryTopo
// The summary topology
type OspfOperData_Ospfv2Instance_Ospfv2LsdbExternal_RouterSummaryLsa_SummaryTopo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Multi topology identifier. The type is interface{} with range: 0..255.
    MtId interface{}

    // Topology Metric. The type is interface{} with range: 0..4294967295.
    TopoMetric interface{}
}

func (summaryTopo *OspfOperData_Ospfv2Instance_Ospfv2LsdbExternal_RouterSummaryLsa_SummaryTopo) GetEntityData() *types.CommonEntityData {
    summaryTopo.EntityData.YFilter = summaryTopo.YFilter
    summaryTopo.EntityData.YangName = "summary-topo"
    summaryTopo.EntityData.BundleName = "cisco_ios_xe"
    summaryTopo.EntityData.ParentYangName = "router-summary-lsa"
    summaryTopo.EntityData.SegmentPath = "summary-topo" + types.AddNoKeyToken(summaryTopo)
    summaryTopo.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospfv2-instance/ospfv2-lsdb-external/router-summary-lsa/" + summaryTopo.EntityData.SegmentPath
    summaryTopo.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    summaryTopo.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    summaryTopo.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    summaryTopo.EntityData.Children = types.NewOrderedMap()
    summaryTopo.EntityData.Leafs = types.NewOrderedMap()
    summaryTopo.EntityData.Leafs.Append("mt-id", types.YLeaf{"MtId", summaryTopo.MtId})
    summaryTopo.EntityData.Leafs.Append("topo-metric", types.YLeaf{"TopoMetric", summaryTopo.TopoMetric})

    summaryTopo.EntityData.YListKeys = []string {}

    return &(summaryTopo.EntityData)
}

// OspfOperData_Ospfv2Instance_Ospfv2LsdbExternal_ExternalLsa
// The external Link State Advertisements
type OspfOperData_Ospfv2Instance_Ospfv2LsdbExternal_ExternalLsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The mask for the external Link State Advertisement. The type is interface{}
    // with range: 0..4294967295.
    ExternalLsaMask interface{}

    // The external topology Link State Advertisement. The type is slice of
    // OspfOperData_Ospfv2Instance_Ospfv2LsdbExternal_ExternalLsa_ExternalTopo.
    ExternalTopo []*OspfOperData_Ospfv2Instance_Ospfv2LsdbExternal_ExternalLsa_ExternalTopo
}

func (externalLsa *OspfOperData_Ospfv2Instance_Ospfv2LsdbExternal_ExternalLsa) GetEntityData() *types.CommonEntityData {
    externalLsa.EntityData.YFilter = externalLsa.YFilter
    externalLsa.EntityData.YangName = "external-lsa"
    externalLsa.EntityData.BundleName = "cisco_ios_xe"
    externalLsa.EntityData.ParentYangName = "ospfv2-lsdb-external"
    externalLsa.EntityData.SegmentPath = "external-lsa"
    externalLsa.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospfv2-instance/ospfv2-lsdb-external/" + externalLsa.EntityData.SegmentPath
    externalLsa.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    externalLsa.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    externalLsa.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    externalLsa.EntityData.Children = types.NewOrderedMap()
    externalLsa.EntityData.Children.Append("external-topo", types.YChild{"ExternalTopo", nil})
    for i := range externalLsa.ExternalTopo {
        types.SetYListKey(externalLsa.ExternalTopo[i], i)
        externalLsa.EntityData.Children.Append(types.GetSegmentPath(externalLsa.ExternalTopo[i]), types.YChild{"ExternalTopo", externalLsa.ExternalTopo[i]})
    }
    externalLsa.EntityData.Leafs = types.NewOrderedMap()
    externalLsa.EntityData.Leafs.Append("external-lsa-mask", types.YLeaf{"ExternalLsaMask", externalLsa.ExternalLsaMask})

    externalLsa.EntityData.YListKeys = []string {}

    return &(externalLsa.EntityData)
}

// OspfOperData_Ospfv2Instance_Ospfv2LsdbExternal_ExternalLsa_ExternalTopo
// The external topology Link State Advertisement
type OspfOperData_Ospfv2Instance_Ospfv2LsdbExternal_ExternalLsa_ExternalTopo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // The multi topology identifier. The type is interface{} with range: 0..255.
    MtId interface{}

    // The topoligy metric type associated with the  Link State Advertisement. The
    // type is OspfExternalMetricType.
    TopoMetricType interface{}

    // The topology metric. The type is interface{} with range: 0..4294967295.
    TopoMetric interface{}

    // The topology forwarding address. The type is one of the following types:
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    TopoForwardingAddress interface{}

    // The topology route tag. The type is interface{} with range: 0..4294967295.
    TopoRouteTag interface{}
}

func (externalTopo *OspfOperData_Ospfv2Instance_Ospfv2LsdbExternal_ExternalLsa_ExternalTopo) GetEntityData() *types.CommonEntityData {
    externalTopo.EntityData.YFilter = externalTopo.YFilter
    externalTopo.EntityData.YangName = "external-topo"
    externalTopo.EntityData.BundleName = "cisco_ios_xe"
    externalTopo.EntityData.ParentYangName = "external-lsa"
    externalTopo.EntityData.SegmentPath = "external-topo" + types.AddNoKeyToken(externalTopo)
    externalTopo.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospfv2-instance/ospfv2-lsdb-external/external-lsa/" + externalTopo.EntityData.SegmentPath
    externalTopo.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    externalTopo.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    externalTopo.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    externalTopo.EntityData.Children = types.NewOrderedMap()
    externalTopo.EntityData.Leafs = types.NewOrderedMap()
    externalTopo.EntityData.Leafs.Append("mt-id", types.YLeaf{"MtId", externalTopo.MtId})
    externalTopo.EntityData.Leafs.Append("topo-metric-type", types.YLeaf{"TopoMetricType", externalTopo.TopoMetricType})
    externalTopo.EntityData.Leafs.Append("topo-metric", types.YLeaf{"TopoMetric", externalTopo.TopoMetric})
    externalTopo.EntityData.Leafs.Append("topo-forwarding-address", types.YLeaf{"TopoForwardingAddress", externalTopo.TopoForwardingAddress})
    externalTopo.EntityData.Leafs.Append("topo-route-tag", types.YLeaf{"TopoRouteTag", externalTopo.TopoRouteTag})

    externalTopo.EntityData.YListKeys = []string {}

    return &(externalTopo.EntityData)
}

// OspfOperData_Ospfv2Instance_Ospfv2LsdbExternal_NssaLsa
// The Not So Stubby Area Link state advertisements
type OspfOperData_Ospfv2Instance_Ospfv2LsdbExternal_NssaLsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The mask for the external Link State Advertisement. The type is interface{}
    // with range: 0..4294967295.
    ExternalLsaMask interface{}

    // The external topology Link State Advertisement. The type is slice of
    // OspfOperData_Ospfv2Instance_Ospfv2LsdbExternal_NssaLsa_ExternalTopo.
    ExternalTopo []*OspfOperData_Ospfv2Instance_Ospfv2LsdbExternal_NssaLsa_ExternalTopo
}

func (nssaLsa *OspfOperData_Ospfv2Instance_Ospfv2LsdbExternal_NssaLsa) GetEntityData() *types.CommonEntityData {
    nssaLsa.EntityData.YFilter = nssaLsa.YFilter
    nssaLsa.EntityData.YangName = "nssa-lsa"
    nssaLsa.EntityData.BundleName = "cisco_ios_xe"
    nssaLsa.EntityData.ParentYangName = "ospfv2-lsdb-external"
    nssaLsa.EntityData.SegmentPath = "nssa-lsa"
    nssaLsa.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospfv2-instance/ospfv2-lsdb-external/" + nssaLsa.EntityData.SegmentPath
    nssaLsa.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nssaLsa.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nssaLsa.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nssaLsa.EntityData.Children = types.NewOrderedMap()
    nssaLsa.EntityData.Children.Append("external-topo", types.YChild{"ExternalTopo", nil})
    for i := range nssaLsa.ExternalTopo {
        types.SetYListKey(nssaLsa.ExternalTopo[i], i)
        nssaLsa.EntityData.Children.Append(types.GetSegmentPath(nssaLsa.ExternalTopo[i]), types.YChild{"ExternalTopo", nssaLsa.ExternalTopo[i]})
    }
    nssaLsa.EntityData.Leafs = types.NewOrderedMap()
    nssaLsa.EntityData.Leafs.Append("external-lsa-mask", types.YLeaf{"ExternalLsaMask", nssaLsa.ExternalLsaMask})

    nssaLsa.EntityData.YListKeys = []string {}

    return &(nssaLsa.EntityData)
}

// OspfOperData_Ospfv2Instance_Ospfv2LsdbExternal_NssaLsa_ExternalTopo
// The external topology Link State Advertisement
type OspfOperData_Ospfv2Instance_Ospfv2LsdbExternal_NssaLsa_ExternalTopo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // The multi topology identifier. The type is interface{} with range: 0..255.
    MtId interface{}

    // The topoligy metric type associated with the  Link State Advertisement. The
    // type is OspfExternalMetricType.
    TopoMetricType interface{}

    // The topology metric. The type is interface{} with range: 0..4294967295.
    TopoMetric interface{}

    // The topology forwarding address. The type is one of the following types:
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    TopoForwardingAddress interface{}

    // The topology route tag. The type is interface{} with range: 0..4294967295.
    TopoRouteTag interface{}
}

func (externalTopo *OspfOperData_Ospfv2Instance_Ospfv2LsdbExternal_NssaLsa_ExternalTopo) GetEntityData() *types.CommonEntityData {
    externalTopo.EntityData.YFilter = externalTopo.YFilter
    externalTopo.EntityData.YangName = "external-topo"
    externalTopo.EntityData.BundleName = "cisco_ios_xe"
    externalTopo.EntityData.ParentYangName = "nssa-lsa"
    externalTopo.EntityData.SegmentPath = "external-topo" + types.AddNoKeyToken(externalTopo)
    externalTopo.EntityData.AbsolutePath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data/ospfv2-instance/ospfv2-lsdb-external/nssa-lsa/" + externalTopo.EntityData.SegmentPath
    externalTopo.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    externalTopo.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    externalTopo.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    externalTopo.EntityData.Children = types.NewOrderedMap()
    externalTopo.EntityData.Leafs = types.NewOrderedMap()
    externalTopo.EntityData.Leafs.Append("mt-id", types.YLeaf{"MtId", externalTopo.MtId})
    externalTopo.EntityData.Leafs.Append("topo-metric-type", types.YLeaf{"TopoMetricType", externalTopo.TopoMetricType})
    externalTopo.EntityData.Leafs.Append("topo-metric", types.YLeaf{"TopoMetric", externalTopo.TopoMetric})
    externalTopo.EntityData.Leafs.Append("topo-forwarding-address", types.YLeaf{"TopoForwardingAddress", externalTopo.TopoForwardingAddress})
    externalTopo.EntityData.Leafs.Append("topo-route-tag", types.YLeaf{"TopoRouteTag", externalTopo.TopoRouteTag})

    externalTopo.EntityData.YListKeys = []string {}

    return &(externalTopo.EntityData)
}

