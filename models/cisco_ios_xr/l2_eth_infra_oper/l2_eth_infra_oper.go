// This module contains a collection of YANG definitions
// for Cisco IOS-XR l2-eth-infra package operational data.
// 
// This module contains definitions
// for the following management objects:
//   mac-accounting: MAC accounting operational data
//   vlan: vlan
//   ethernet-encapsulation: ethernet encapsulation
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package l2_eth_infra_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package l2_eth_infra_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-l2-eth-infra-oper mac-accounting}", reflect.TypeOf(MacAccounting{}))
    ydk.RegisterEntity("Cisco-IOS-XR-l2-eth-infra-oper:mac-accounting", reflect.TypeOf(MacAccounting{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-l2-eth-infra-oper vlan}", reflect.TypeOf(Vlan{}))
    ydk.RegisterEntity("Cisco-IOS-XR-l2-eth-infra-oper:vlan", reflect.TypeOf(Vlan{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-l2-eth-infra-oper ethernet-encapsulation}", reflect.TypeOf(EthernetEncapsulation{}))
    ydk.RegisterEntity("Cisco-IOS-XR-l2-eth-infra-oper:ethernet-encapsulation", reflect.TypeOf(EthernetEncapsulation{}))
}

// VlanEncaps represents VLAN encapsulation
type VlanEncaps string

const (
    // No encapsulation
    VlanEncaps_no_encapsulation VlanEncaps = "no-encapsulation"

    // IEEE 802.1Q encapsulation
    VlanEncaps_dot1q VlanEncaps = "dot1q"

    // Double 802.1Q encapsulation
    VlanEncaps_qinq VlanEncaps = "qinq"

    // Double 802.1Q wildcarded encapsulation
    VlanEncaps_qin_any VlanEncaps = "qin-any"

    // IEEE 802.1Q native VLAN encapsulation
    VlanEncaps_dot1q_native VlanEncaps = "dot1q-native"

    // IEEE 802.1ad encapsulation
    VlanEncaps_dot1ad VlanEncaps = "dot1ad"

    // IEEE 802.1ad native VLAN encapsulation
    VlanEncaps_dot1ad_native VlanEncaps = "dot1ad-native"

    // Ethernet Service Instance
    VlanEncaps_service_instance VlanEncaps = "service-instance"

    // IEEE 802.1ad 802.1Q encapsulation
    VlanEncaps_dot1ad_dot1q VlanEncaps = "dot1ad-dot1q"

    // IEEE 802.1ad wildcard 802.1Q encapsulation
    VlanEncaps_dot1ad_any VlanEncaps = "dot1ad-any"
)

// VlanSwitchedMode represents VLAN-Switched mode
type VlanSwitchedMode string

const (
    // Disabled
    VlanSwitchedMode_none VlanSwitchedMode = "none"

    // Trunk port
    VlanSwitchedMode_trunk_port VlanSwitchedMode = "trunk-port"

    // Access port
    VlanSwitchedMode_access_port VlanSwitchedMode = "access-port"
)

// ImStateEnum represents Im state enum
type ImStateEnum string

const (
    // im state not ready
    ImStateEnum_im_state_not_ready ImStateEnum = "im-state-not-ready"

    // im state admin down
    ImStateEnum_im_state_admin_down ImStateEnum = "im-state-admin-down"

    // im state down
    ImStateEnum_im_state_down ImStateEnum = "im-state-down"

    // im state up
    ImStateEnum_im_state_up ImStateEnum = "im-state-up"

    // im state shutdown
    ImStateEnum_im_state_shutdown ImStateEnum = "im-state-shutdown"

    // im state err disable
    ImStateEnum_im_state_err_disable ImStateEnum = "im-state-err-disable"

    // im state down immediate
    ImStateEnum_im_state_down_immediate ImStateEnum = "im-state-down-immediate"

    // im state down immediate admin
    ImStateEnum_im_state_down_immediate_admin ImStateEnum = "im-state-down-immediate-admin"

    // im state down graceful
    ImStateEnum_im_state_down_graceful ImStateEnum = "im-state-down-graceful"

    // im state begin shutdown
    ImStateEnum_im_state_begin_shutdown ImStateEnum = "im-state-begin-shutdown"

    // im state end shutdown
    ImStateEnum_im_state_end_shutdown ImStateEnum = "im-state-end-shutdown"

    // im state begin error disable
    ImStateEnum_im_state_begin_error_disable ImStateEnum = "im-state-begin-error-disable"

    // im state end error disable
    ImStateEnum_im_state_end_error_disable ImStateEnum = "im-state-end-error-disable"

    // im state begin down graceful
    ImStateEnum_im_state_begin_down_graceful ImStateEnum = "im-state-begin-down-graceful"

    // im state reset
    ImStateEnum_im_state_reset ImStateEnum = "im-state-reset"

    // im state operational
    ImStateEnum_im_state_operational ImStateEnum = "im-state-operational"

    // im state not operational
    ImStateEnum_im_state_not_operational ImStateEnum = "im-state-not-operational"

    // im state unknown
    ImStateEnum_im_state_unknown ImStateEnum = "im-state-unknown"

    // im state last
    ImStateEnum_im_state_last ImStateEnum = "im-state-last"
)

// EfpTagPriority represents Priority
type EfpTagPriority string

const (
    // Priority 0
    EfpTagPriority_priority0 EfpTagPriority = "priority0"

    // Priority 1
    EfpTagPriority_priority1 EfpTagPriority = "priority1"

    // Priority 2
    EfpTagPriority_priority2 EfpTagPriority = "priority2"

    // Priority 3
    EfpTagPriority_priority3 EfpTagPriority = "priority3"

    // Priority 4
    EfpTagPriority_priority4 EfpTagPriority = "priority4"

    // Priority 5
    EfpTagPriority_priority5 EfpTagPriority = "priority5"

    // Priority 6
    EfpTagPriority_priority6 EfpTagPriority = "priority6"

    // Priority 7
    EfpTagPriority_priority7 EfpTagPriority = "priority7"

    // Any priority
    EfpTagPriority_priority_any EfpTagPriority = "priority-any"
)

// EfpTagEtype represents Tag ethertype
type EfpTagEtype string

const (
    // Untagged
    EfpTagEtype_untagged EfpTagEtype = "untagged"

    // Dot1Q
    EfpTagEtype_dot1q EfpTagEtype = "dot1q"

    // Dot1ad
    EfpTagEtype_dot1ad EfpTagEtype = "dot1ad"
)

// VlanService represents Layer 2 vs. Layer 3 Terminated Service
type VlanService string

const (
    // Layer 2 Transport Service
    VlanService_vlan_service_l2 VlanService = "vlan-service-l2"

    // Layer 3 Terminated Service
    VlanService_vlan_service_l3 VlanService = "vlan-service-l3"
)

// EfpPayloadEtype represents Payload ethertype match
type EfpPayloadEtype string

const (
    // Any
    EfpPayloadEtype_payload_ethertype_any EfpPayloadEtype = "payload-ethertype-any"

    // IP
    EfpPayloadEtype_payload_ethertype_ip EfpPayloadEtype = "payload-ethertype-ip"

    // PPPoE
    EfpPayloadEtype_payload_ethertype_pppoe EfpPayloadEtype = "payload-ethertype-pppoe"
)

// VlanQinqOuterEtype represents QinQ Outer Tag Ethertype
type VlanQinqOuterEtype string

const (
    // Dot1Q (0x8100)
    VlanQinqOuterEtype_ether_type8100 VlanQinqOuterEtype = "ether-type8100"

    // 0x9100
    VlanQinqOuterEtype_ether_type9100 VlanQinqOuterEtype = "ether-type9100"

    // 0x9200
    VlanQinqOuterEtype_ether_type9200 VlanQinqOuterEtype = "ether-type9200"
)

// EthCapsUcastMacMode represents Eth caps ucast mac mode
type EthCapsUcastMacMode string

const (
    // Reserved
    EthCapsUcastMacMode_reserved EthCapsUcastMacMode = "reserved"

    // Permit
    EthCapsUcastMacMode_permit EthCapsUcastMacMode = "permit"
)

// EthFiltering represents Ethernet frame filtering
type EthFiltering string

const (
    // No IEEE 802.1Q/802.1ad/MAC relay multicast MAC
    // address filtering
    EthFiltering_no_filtering EthFiltering = "no-filtering"

    // IEEE 802.1q C-VLAN filtering
    EthFiltering_dot1q_filtering EthFiltering = "dot1q-filtering"

    // IEEE 802.1ad S-VLAN filtering
    EthFiltering_dot1ad_filtering EthFiltering = "dot1ad-filtering"

    // IEEE 802.1aj 2-Port MAC relay filtering
    EthFiltering_two_port_mac_relay_filtering EthFiltering = "two-port-mac-relay-filtering"
)

// MacAccounting
// MAC accounting operational data
type MacAccounting struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MAC accounting interface table in MIB lexicographic order.
    Interfaces MacAccounting_Interfaces
}

func (macAccounting *MacAccounting) GetEntityData() *types.CommonEntityData {
    macAccounting.EntityData.YFilter = macAccounting.YFilter
    macAccounting.EntityData.YangName = "mac-accounting"
    macAccounting.EntityData.BundleName = "cisco_ios_xr"
    macAccounting.EntityData.ParentYangName = "Cisco-IOS-XR-l2-eth-infra-oper"
    macAccounting.EntityData.SegmentPath = "Cisco-IOS-XR-l2-eth-infra-oper:mac-accounting"
    macAccounting.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macAccounting.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macAccounting.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macAccounting.EntityData.Children = types.NewOrderedMap()
    macAccounting.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &macAccounting.Interfaces})
    macAccounting.EntityData.Leafs = types.NewOrderedMap()

    macAccounting.EntityData.YListKeys = []string {}

    return &(macAccounting.EntityData)
}

// MacAccounting_Interfaces
// MAC accounting interface table in MIB
// lexicographic order
type MacAccounting_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational data and statistics for an interface configured with MAC
    // accounting enabled. The type is slice of
    // MacAccounting_Interfaces_Interface.
    Interface []*MacAccounting_Interfaces_Interface
}

func (interfaces *MacAccounting_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "mac-accounting"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = types.NewOrderedMap()
    interfaces.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range interfaces.Interface {
        interfaces.EntityData.Children.Append(types.GetSegmentPath(interfaces.Interface[i]), types.YChild{"Interface", interfaces.Interface[i]})
    }
    interfaces.EntityData.Leafs = types.NewOrderedMap()

    interfaces.EntityData.YListKeys = []string {}

    return &(interfaces.EntityData)
}

// MacAccounting_Interfaces_Interface
// Operational data and statistics for an
// interface configured with MAC accounting
// enabled
type MacAccounting_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The interface name. The type is string with
    // pattern: [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // MAC accounting state for the interface.
    State MacAccounting_Interfaces_Interface_State

    // Ingress MAC accounting statistics. The type is slice of
    // MacAccounting_Interfaces_Interface_IngressStatistic.
    IngressStatistic []*MacAccounting_Interfaces_Interface_IngressStatistic

    // Egress MAC accounting statistics. The type is slice of
    // MacAccounting_Interfaces_Interface_EgressStatistic.
    EgressStatistic []*MacAccounting_Interfaces_Interface_EgressStatistic
}

func (self *MacAccounting_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceName, "interface-name")
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("state", types.YChild{"State", &self.State})
    self.EntityData.Children.Append("ingress-statistic", types.YChild{"IngressStatistic", nil})
    for i := range self.IngressStatistic {
        self.EntityData.Children.Append(types.GetSegmentPath(self.IngressStatistic[i]), types.YChild{"IngressStatistic", self.IngressStatistic[i]})
    }
    self.EntityData.Children.Append("egress-statistic", types.YChild{"EgressStatistic", nil})
    for i := range self.EgressStatistic {
        self.EntityData.Children.Append(types.GetSegmentPath(self.EgressStatistic[i]), types.YChild{"EgressStatistic", self.EgressStatistic[i]})
    }
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})

    self.EntityData.YListKeys = []string {"InterfaceName"}

    return &(self.EntityData)
}

// MacAccounting_Interfaces_Interface_State
// MAC accounting state for the interface
type MacAccounting_Interfaces_Interface_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MAC accounting on on ingress. The type is bool.
    IsIngressEnabled interface{}

    // MAC accounting on on egress. The type is bool.
    IsEgressEnabled interface{}

    // MAC accounting entries available on ingress. The type is interface{} with
    // range: 0..4294967295.
    NumberAvailableIngress interface{}

    // MAC accounting entries available on egress. The type is interface{} with
    // range: 0..4294967295.
    NumberAvailableEgress interface{}

    // MAC accountng entries available across the node. The type is interface{}
    // with range: 0..4294967295.
    NumberAvailableOnNode interface{}
}

func (state *MacAccounting_Interfaces_Interface_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "cisco_ios_xr"
    state.EntityData.ParentYangName = "interface"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    state.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("is-ingress-enabled", types.YLeaf{"IsIngressEnabled", state.IsIngressEnabled})
    state.EntityData.Leafs.Append("is-egress-enabled", types.YLeaf{"IsEgressEnabled", state.IsEgressEnabled})
    state.EntityData.Leafs.Append("number-available-ingress", types.YLeaf{"NumberAvailableIngress", state.NumberAvailableIngress})
    state.EntityData.Leafs.Append("number-available-egress", types.YLeaf{"NumberAvailableEgress", state.NumberAvailableEgress})
    state.EntityData.Leafs.Append("number-available-on-node", types.YLeaf{"NumberAvailableOnNode", state.NumberAvailableOnNode})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// MacAccounting_Interfaces_Interface_IngressStatistic
// Ingress MAC accounting statistics
type MacAccounting_Interfaces_Interface_IngressStatistic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // 48bit MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}

    // Number of packets counted. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Number of bytes counted. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    Bytes interface{}
}

func (ingressStatistic *MacAccounting_Interfaces_Interface_IngressStatistic) GetEntityData() *types.CommonEntityData {
    ingressStatistic.EntityData.YFilter = ingressStatistic.YFilter
    ingressStatistic.EntityData.YangName = "ingress-statistic"
    ingressStatistic.EntityData.BundleName = "cisco_ios_xr"
    ingressStatistic.EntityData.ParentYangName = "interface"
    ingressStatistic.EntityData.SegmentPath = "ingress-statistic"
    ingressStatistic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ingressStatistic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ingressStatistic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ingressStatistic.EntityData.Children = types.NewOrderedMap()
    ingressStatistic.EntityData.Leafs = types.NewOrderedMap()
    ingressStatistic.EntityData.Leafs.Append("mac-address", types.YLeaf{"MacAddress", ingressStatistic.MacAddress})
    ingressStatistic.EntityData.Leafs.Append("packets", types.YLeaf{"Packets", ingressStatistic.Packets})
    ingressStatistic.EntityData.Leafs.Append("bytes", types.YLeaf{"Bytes", ingressStatistic.Bytes})

    ingressStatistic.EntityData.YListKeys = []string {}

    return &(ingressStatistic.EntityData)
}

// MacAccounting_Interfaces_Interface_EgressStatistic
// Egress MAC accounting statistics
type MacAccounting_Interfaces_Interface_EgressStatistic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // 48bit MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}

    // Number of packets counted. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Number of bytes counted. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    Bytes interface{}
}

func (egressStatistic *MacAccounting_Interfaces_Interface_EgressStatistic) GetEntityData() *types.CommonEntityData {
    egressStatistic.EntityData.YFilter = egressStatistic.YFilter
    egressStatistic.EntityData.YangName = "egress-statistic"
    egressStatistic.EntityData.BundleName = "cisco_ios_xr"
    egressStatistic.EntityData.ParentYangName = "interface"
    egressStatistic.EntityData.SegmentPath = "egress-statistic"
    egressStatistic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    egressStatistic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    egressStatistic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    egressStatistic.EntityData.Children = types.NewOrderedMap()
    egressStatistic.EntityData.Leafs = types.NewOrderedMap()
    egressStatistic.EntityData.Leafs.Append("mac-address", types.YLeaf{"MacAddress", egressStatistic.MacAddress})
    egressStatistic.EntityData.Leafs.Append("packets", types.YLeaf{"Packets", egressStatistic.Packets})
    egressStatistic.EntityData.Leafs.Append("bytes", types.YLeaf{"Bytes", egressStatistic.Bytes})

    egressStatistic.EntityData.YListKeys = []string {}

    return &(egressStatistic.EntityData)
}

// Vlan
// vlan
type Vlan struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Per node VLAN operational data.
    Nodes Vlan_Nodes
}

func (vlan *Vlan) GetEntityData() *types.CommonEntityData {
    vlan.EntityData.YFilter = vlan.YFilter
    vlan.EntityData.YangName = "vlan"
    vlan.EntityData.BundleName = "cisco_ios_xr"
    vlan.EntityData.ParentYangName = "Cisco-IOS-XR-l2-eth-infra-oper"
    vlan.EntityData.SegmentPath = "Cisco-IOS-XR-l2-eth-infra-oper:vlan"
    vlan.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vlan.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vlan.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vlan.EntityData.Children = types.NewOrderedMap()
    vlan.EntityData.Children.Append("nodes", types.YChild{"Nodes", &vlan.Nodes})
    vlan.EntityData.Leafs = types.NewOrderedMap()

    vlan.EntityData.YListKeys = []string {}

    return &(vlan.EntityData)
}

// Vlan_Nodes
// Per node VLAN operational data
type Vlan_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The VLAN operational data for a particular node. The type is slice of
    // Vlan_Nodes_Node.
    Node []*Vlan_Nodes_Node
}

func (nodes *Vlan_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "vlan"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodes.EntityData.Children = types.NewOrderedMap()
    nodes.EntityData.Children.Append("node", types.YChild{"Node", nil})
    for i := range nodes.Node {
        nodes.EntityData.Children.Append(types.GetSegmentPath(nodes.Node[i]), types.YChild{"Node", nodes.Node[i]})
    }
    nodes.EntityData.Leafs = types.NewOrderedMap()

    nodes.EntityData.YListKeys = []string {}

    return &(nodes.EntityData)
}

// Vlan_Nodes_Node
// The VLAN operational data for a particular node
type Vlan_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The identifier for the node. The type is string
    // with pattern: ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeId interface{}

    // VLAN trunk table (specific to this node).
    Trunks Vlan_Nodes_Node_Trunks

    // VLAN interface table (specific to this node).
    Interfaces Vlan_Nodes_Node_Interfaces

    // VLAN tag allocation table (specific to this node).
    TagAllocations Vlan_Nodes_Node_TagAllocations
}

func (node *Vlan_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeId, "node-id")
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("trunks", types.YChild{"Trunks", &node.Trunks})
    node.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &node.Interfaces})
    node.EntityData.Children.Append("tag-allocations", types.YChild{"TagAllocations", &node.TagAllocations})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-id", types.YLeaf{"NodeId", node.NodeId})

    node.EntityData.YListKeys = []string {"NodeId"}

    return &(node.EntityData)
}

// Vlan_Nodes_Node_Trunks
// VLAN trunk table (specific to this node)
type Vlan_Nodes_Node_Trunks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational data for trunk interfaces configured with VLANs. The type is
    // slice of Vlan_Nodes_Node_Trunks_Trunk.
    Trunk []*Vlan_Nodes_Node_Trunks_Trunk
}

func (trunks *Vlan_Nodes_Node_Trunks) GetEntityData() *types.CommonEntityData {
    trunks.EntityData.YFilter = trunks.YFilter
    trunks.EntityData.YangName = "trunks"
    trunks.EntityData.BundleName = "cisco_ios_xr"
    trunks.EntityData.ParentYangName = "node"
    trunks.EntityData.SegmentPath = "trunks"
    trunks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trunks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trunks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trunks.EntityData.Children = types.NewOrderedMap()
    trunks.EntityData.Children.Append("trunk", types.YChild{"Trunk", nil})
    for i := range trunks.Trunk {
        trunks.EntityData.Children.Append(types.GetSegmentPath(trunks.Trunk[i]), types.YChild{"Trunk", trunks.Trunk[i]})
    }
    trunks.EntityData.Leafs = types.NewOrderedMap()

    trunks.EntityData.YListKeys = []string {}

    return &(trunks.EntityData)
}

// Vlan_Nodes_Node_Trunks_Trunk
// Operational data for trunk interfaces
// configured with VLANs
type Vlan_Nodes_Node_Trunks_Trunk struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The interface name. The type is string with
    // pattern: [a-zA-Z0-9._/-]+.
    Interface interface{}

    // Interface name. The type is string with pattern: [a-zA-Z0-9._/-]+.
    InterfaceXr interface{}

    // Interface state. The type is ImStateEnum.
    State interface{}

    // L2 MTU. The type is interface{} with range: 0..65535.
    Mtu interface{}

    // QinQ Outer Tag Ether Type. The type is VlanQinqOuterEtype.
    QinqOuterEtherType interface{}

    // Number of subinterfaces with 802.1ad outer tag. The type is interface{}
    // with range: 0..4294967295.
    Dot1adCount interface{}

    // Interface/Sub-interface handling untagged frames. The type is string with
    // pattern: [a-zA-Z0-9._/-]+.
    UntaggedInterface interface{}

    // IEEE 802.1Q/802.1ad multicast MAC address filtering. The type is
    // EthFiltering.
    MacFiltering interface{}

    // Layer 2 Transport Subinterfaces.
    Layer2SubInterfaces Vlan_Nodes_Node_Trunks_Trunk_Layer2SubInterfaces

    // Layer 3 Terminated Subinterfaces.
    Layer3SubInterfaces Vlan_Nodes_Node_Trunks_Trunk_Layer3SubInterfaces

    // VLAN-Switched information.
    VlanSwitched Vlan_Nodes_Node_Trunks_Trunk_VlanSwitched
}

func (trunk *Vlan_Nodes_Node_Trunks_Trunk) GetEntityData() *types.CommonEntityData {
    trunk.EntityData.YFilter = trunk.YFilter
    trunk.EntityData.YangName = "trunk"
    trunk.EntityData.BundleName = "cisco_ios_xr"
    trunk.EntityData.ParentYangName = "trunks"
    trunk.EntityData.SegmentPath = "trunk" + types.AddKeyToken(trunk.Interface, "interface")
    trunk.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trunk.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trunk.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trunk.EntityData.Children = types.NewOrderedMap()
    trunk.EntityData.Children.Append("layer2-sub-interfaces", types.YChild{"Layer2SubInterfaces", &trunk.Layer2SubInterfaces})
    trunk.EntityData.Children.Append("layer3-sub-interfaces", types.YChild{"Layer3SubInterfaces", &trunk.Layer3SubInterfaces})
    trunk.EntityData.Children.Append("vlan-switched", types.YChild{"VlanSwitched", &trunk.VlanSwitched})
    trunk.EntityData.Leafs = types.NewOrderedMap()
    trunk.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", trunk.Interface})
    trunk.EntityData.Leafs.Append("interface-xr", types.YLeaf{"InterfaceXr", trunk.InterfaceXr})
    trunk.EntityData.Leafs.Append("state", types.YLeaf{"State", trunk.State})
    trunk.EntityData.Leafs.Append("mtu", types.YLeaf{"Mtu", trunk.Mtu})
    trunk.EntityData.Leafs.Append("qinq-outer-ether-type", types.YLeaf{"QinqOuterEtherType", trunk.QinqOuterEtherType})
    trunk.EntityData.Leafs.Append("dot1ad-count", types.YLeaf{"Dot1adCount", trunk.Dot1adCount})
    trunk.EntityData.Leafs.Append("untagged-interface", types.YLeaf{"UntaggedInterface", trunk.UntaggedInterface})
    trunk.EntityData.Leafs.Append("mac-filtering", types.YLeaf{"MacFiltering", trunk.MacFiltering})

    trunk.EntityData.YListKeys = []string {"Interface"}

    return &(trunk.EntityData)
}

// Vlan_Nodes_Node_Trunks_Trunk_Layer2SubInterfaces
// Layer 2 Transport Subinterfaces
type Vlan_Nodes_Node_Trunks_Trunk_Layer2SubInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Total number of Layer 2 subinterfaces configured. The type is interface{}
    // with range: 0..4294967295.
    TotalCount interface{}

    // Number of single tagged subinterfaces. The type is interface{} with range:
    // 0..4294967295.
    Dot1qCount interface{}

    // Number of double tagged subinterfaces with explicit inner tag. The type is
    // interface{} with range: 0..4294967295.
    QinQCount interface{}

    // Number of double tagged subinterfaces with wildcarded inner tag. The type
    // is interface{} with range: 0..4294967295.
    QinAnyCount interface{}

    // Number of subinterfaces without VLAN tag configuration. The type is
    // interface{} with range: 0..4294967295.
    UntaggedCount interface{}

    // Numbers of subinterfaces up, down or administratively shut down.
    StateCounters Vlan_Nodes_Node_Trunks_Trunk_Layer2SubInterfaces_StateCounters
}

func (layer2SubInterfaces *Vlan_Nodes_Node_Trunks_Trunk_Layer2SubInterfaces) GetEntityData() *types.CommonEntityData {
    layer2SubInterfaces.EntityData.YFilter = layer2SubInterfaces.YFilter
    layer2SubInterfaces.EntityData.YangName = "layer2-sub-interfaces"
    layer2SubInterfaces.EntityData.BundleName = "cisco_ios_xr"
    layer2SubInterfaces.EntityData.ParentYangName = "trunk"
    layer2SubInterfaces.EntityData.SegmentPath = "layer2-sub-interfaces"
    layer2SubInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    layer2SubInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    layer2SubInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    layer2SubInterfaces.EntityData.Children = types.NewOrderedMap()
    layer2SubInterfaces.EntityData.Children.Append("state-counters", types.YChild{"StateCounters", &layer2SubInterfaces.StateCounters})
    layer2SubInterfaces.EntityData.Leafs = types.NewOrderedMap()
    layer2SubInterfaces.EntityData.Leafs.Append("total-count", types.YLeaf{"TotalCount", layer2SubInterfaces.TotalCount})
    layer2SubInterfaces.EntityData.Leafs.Append("dot1q-count", types.YLeaf{"Dot1qCount", layer2SubInterfaces.Dot1qCount})
    layer2SubInterfaces.EntityData.Leafs.Append("qin-q-count", types.YLeaf{"QinQCount", layer2SubInterfaces.QinQCount})
    layer2SubInterfaces.EntityData.Leafs.Append("qin-any-count", types.YLeaf{"QinAnyCount", layer2SubInterfaces.QinAnyCount})
    layer2SubInterfaces.EntityData.Leafs.Append("untagged-count", types.YLeaf{"UntaggedCount", layer2SubInterfaces.UntaggedCount})

    layer2SubInterfaces.EntityData.YListKeys = []string {}

    return &(layer2SubInterfaces.EntityData)
}

// Vlan_Nodes_Node_Trunks_Trunk_Layer2SubInterfaces_StateCounters
// Numbers of subinterfaces up, down or
// administratively shut down
type Vlan_Nodes_Node_Trunks_Trunk_Layer2SubInterfaces_StateCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of subinterfaces which are up. The type is interface{} with range:
    // 0..4294967295.
    Up interface{}

    // Number of subinterfaces which are down. The type is interface{} with range:
    // 0..4294967295.
    Down interface{}

    // Number of subinterfaces which are administrativelyshutdown. The type is
    // interface{} with range: 0..4294967295.
    AdminDown interface{}
}

func (stateCounters *Vlan_Nodes_Node_Trunks_Trunk_Layer2SubInterfaces_StateCounters) GetEntityData() *types.CommonEntityData {
    stateCounters.EntityData.YFilter = stateCounters.YFilter
    stateCounters.EntityData.YangName = "state-counters"
    stateCounters.EntityData.BundleName = "cisco_ios_xr"
    stateCounters.EntityData.ParentYangName = "layer2-sub-interfaces"
    stateCounters.EntityData.SegmentPath = "state-counters"
    stateCounters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stateCounters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stateCounters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stateCounters.EntityData.Children = types.NewOrderedMap()
    stateCounters.EntityData.Leafs = types.NewOrderedMap()
    stateCounters.EntityData.Leafs.Append("up", types.YLeaf{"Up", stateCounters.Up})
    stateCounters.EntityData.Leafs.Append("down", types.YLeaf{"Down", stateCounters.Down})
    stateCounters.EntityData.Leafs.Append("admin-down", types.YLeaf{"AdminDown", stateCounters.AdminDown})

    stateCounters.EntityData.YListKeys = []string {}

    return &(stateCounters.EntityData)
}

// Vlan_Nodes_Node_Trunks_Trunk_Layer3SubInterfaces
// Layer 3 Terminated Subinterfaces
type Vlan_Nodes_Node_Trunks_Trunk_Layer3SubInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Total number of Layer 3 subinterfaces configured. The type is interface{}
    // with range: 0..4294967295.
    TotalCount interface{}

    // Number of single tagged subinterfaces. The type is interface{} with range:
    // 0..4294967295.
    Dot1qCount interface{}

    // Number of double tagged subinterfaces. The type is interface{} with range:
    // 0..4294967295.
    QinQCount interface{}

    // Number of subinterfaces without VLAN tag configuration. The type is
    // interface{} with range: 0..4294967295.
    UntaggedCount interface{}

    // Native VLAN ID configured on trunk. The type is interface{} with range:
    // 0..65535.
    NativeVlan interface{}

    // Numbers of subinterfaces up, down or administratively shut down.
    StateCounters Vlan_Nodes_Node_Trunks_Trunk_Layer3SubInterfaces_StateCounters
}

func (layer3SubInterfaces *Vlan_Nodes_Node_Trunks_Trunk_Layer3SubInterfaces) GetEntityData() *types.CommonEntityData {
    layer3SubInterfaces.EntityData.YFilter = layer3SubInterfaces.YFilter
    layer3SubInterfaces.EntityData.YangName = "layer3-sub-interfaces"
    layer3SubInterfaces.EntityData.BundleName = "cisco_ios_xr"
    layer3SubInterfaces.EntityData.ParentYangName = "trunk"
    layer3SubInterfaces.EntityData.SegmentPath = "layer3-sub-interfaces"
    layer3SubInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    layer3SubInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    layer3SubInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    layer3SubInterfaces.EntityData.Children = types.NewOrderedMap()
    layer3SubInterfaces.EntityData.Children.Append("state-counters", types.YChild{"StateCounters", &layer3SubInterfaces.StateCounters})
    layer3SubInterfaces.EntityData.Leafs = types.NewOrderedMap()
    layer3SubInterfaces.EntityData.Leafs.Append("total-count", types.YLeaf{"TotalCount", layer3SubInterfaces.TotalCount})
    layer3SubInterfaces.EntityData.Leafs.Append("dot1q-count", types.YLeaf{"Dot1qCount", layer3SubInterfaces.Dot1qCount})
    layer3SubInterfaces.EntityData.Leafs.Append("qin-q-count", types.YLeaf{"QinQCount", layer3SubInterfaces.QinQCount})
    layer3SubInterfaces.EntityData.Leafs.Append("untagged-count", types.YLeaf{"UntaggedCount", layer3SubInterfaces.UntaggedCount})
    layer3SubInterfaces.EntityData.Leafs.Append("native-vlan", types.YLeaf{"NativeVlan", layer3SubInterfaces.NativeVlan})

    layer3SubInterfaces.EntityData.YListKeys = []string {}

    return &(layer3SubInterfaces.EntityData)
}

// Vlan_Nodes_Node_Trunks_Trunk_Layer3SubInterfaces_StateCounters
// Numbers of subinterfaces up, down or
// administratively shut down
type Vlan_Nodes_Node_Trunks_Trunk_Layer3SubInterfaces_StateCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of subinterfaces which are up. The type is interface{} with range:
    // 0..4294967295.
    Up interface{}

    // Number of subinterfaces which are down. The type is interface{} with range:
    // 0..4294967295.
    Down interface{}

    // Number of subinterfaces which are administrativelyshutdown. The type is
    // interface{} with range: 0..4294967295.
    AdminDown interface{}
}

func (stateCounters *Vlan_Nodes_Node_Trunks_Trunk_Layer3SubInterfaces_StateCounters) GetEntityData() *types.CommonEntityData {
    stateCounters.EntityData.YFilter = stateCounters.YFilter
    stateCounters.EntityData.YangName = "state-counters"
    stateCounters.EntityData.BundleName = "cisco_ios_xr"
    stateCounters.EntityData.ParentYangName = "layer3-sub-interfaces"
    stateCounters.EntityData.SegmentPath = "state-counters"
    stateCounters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stateCounters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stateCounters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stateCounters.EntityData.Children = types.NewOrderedMap()
    stateCounters.EntityData.Leafs = types.NewOrderedMap()
    stateCounters.EntityData.Leafs.Append("up", types.YLeaf{"Up", stateCounters.Up})
    stateCounters.EntityData.Leafs.Append("down", types.YLeaf{"Down", stateCounters.Down})
    stateCounters.EntityData.Leafs.Append("admin-down", types.YLeaf{"AdminDown", stateCounters.AdminDown})

    stateCounters.EntityData.YListKeys = []string {}

    return &(stateCounters.EntityData)
}

// Vlan_Nodes_Node_Trunks_Trunk_VlanSwitched
// VLAN-Switched information
type Vlan_Nodes_Node_Trunks_Trunk_VlanSwitched struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VLAN-Switched mode. The type is VlanSwitchedMode.
    Mode interface{}

    // VLAN-Switched Access VLAN. The type is interface{} with range: 0..65535.
    AccessVlan interface{}

    // VLAN-Switched Trunk VLAN ranges.
    TrunkVlanRanges Vlan_Nodes_Node_Trunks_Trunk_VlanSwitched_TrunkVlanRanges
}

func (vlanSwitched *Vlan_Nodes_Node_Trunks_Trunk_VlanSwitched) GetEntityData() *types.CommonEntityData {
    vlanSwitched.EntityData.YFilter = vlanSwitched.YFilter
    vlanSwitched.EntityData.YangName = "vlan-switched"
    vlanSwitched.EntityData.BundleName = "cisco_ios_xr"
    vlanSwitched.EntityData.ParentYangName = "trunk"
    vlanSwitched.EntityData.SegmentPath = "vlan-switched"
    vlanSwitched.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vlanSwitched.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vlanSwitched.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vlanSwitched.EntityData.Children = types.NewOrderedMap()
    vlanSwitched.EntityData.Children.Append("trunk-vlan-ranges", types.YChild{"TrunkVlanRanges", &vlanSwitched.TrunkVlanRanges})
    vlanSwitched.EntityData.Leafs = types.NewOrderedMap()
    vlanSwitched.EntityData.Leafs.Append("mode", types.YLeaf{"Mode", vlanSwitched.Mode})
    vlanSwitched.EntityData.Leafs.Append("access-vlan", types.YLeaf{"AccessVlan", vlanSwitched.AccessVlan})

    vlanSwitched.EntityData.YListKeys = []string {}

    return &(vlanSwitched.EntityData)
}

// Vlan_Nodes_Node_Trunks_Trunk_VlanSwitched_TrunkVlanRanges
// VLAN-Switched Trunk VLAN ranges
type Vlan_Nodes_Node_Trunks_Trunk_VlanSwitched_TrunkVlanRanges struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Payload Ethertype to match. The type is EfpPayloadEtype.
    PayloadEthertype interface{}

    // Number of tags popped on ingress. The type is interface{} with range:
    // 0..65535.
    TagsPopped interface{}

    // Whether the packet must match the encapsulation exactly, with no further
    // inner tags. The type is bool.
    IsExactMatch interface{}

    // Whether this represents the native VLAN on the port. The type is bool.
    IsNativeVlan interface{}

    // Whether the native VLAN is customer-tag preserving. The type is bool.
    IsNativePreserving interface{}

    // The source MAC address to match on ingress. The type is string with
    // pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    SourceMacMatch interface{}

    // The destination MAC address to match on ingress. The type is string with
    // pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    DestinationMacMatch interface{}

    // VLAN tags for locally-sourced traffic.
    LocalTrafficStack Vlan_Nodes_Node_Trunks_Trunk_VlanSwitched_TrunkVlanRanges_LocalTrafficStack

    // Tags to match on ingress packets. The type is slice of
    // Vlan_Nodes_Node_Trunks_Trunk_VlanSwitched_TrunkVlanRanges_TagsToMatch.
    TagsToMatch []*Vlan_Nodes_Node_Trunks_Trunk_VlanSwitched_TrunkVlanRanges_TagsToMatch

    // VLAN tags pushed on egress. The type is slice of
    // Vlan_Nodes_Node_Trunks_Trunk_VlanSwitched_TrunkVlanRanges_Pushe.
    Pushe []*Vlan_Nodes_Node_Trunks_Trunk_VlanSwitched_TrunkVlanRanges_Pushe
}

func (trunkVlanRanges *Vlan_Nodes_Node_Trunks_Trunk_VlanSwitched_TrunkVlanRanges) GetEntityData() *types.CommonEntityData {
    trunkVlanRanges.EntityData.YFilter = trunkVlanRanges.YFilter
    trunkVlanRanges.EntityData.YangName = "trunk-vlan-ranges"
    trunkVlanRanges.EntityData.BundleName = "cisco_ios_xr"
    trunkVlanRanges.EntityData.ParentYangName = "vlan-switched"
    trunkVlanRanges.EntityData.SegmentPath = "trunk-vlan-ranges"
    trunkVlanRanges.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trunkVlanRanges.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trunkVlanRanges.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trunkVlanRanges.EntityData.Children = types.NewOrderedMap()
    trunkVlanRanges.EntityData.Children.Append("local-traffic-stack", types.YChild{"LocalTrafficStack", &trunkVlanRanges.LocalTrafficStack})
    trunkVlanRanges.EntityData.Children.Append("tags-to-match", types.YChild{"TagsToMatch", nil})
    for i := range trunkVlanRanges.TagsToMatch {
        trunkVlanRanges.EntityData.Children.Append(types.GetSegmentPath(trunkVlanRanges.TagsToMatch[i]), types.YChild{"TagsToMatch", trunkVlanRanges.TagsToMatch[i]})
    }
    trunkVlanRanges.EntityData.Children.Append("pushe", types.YChild{"Pushe", nil})
    for i := range trunkVlanRanges.Pushe {
        trunkVlanRanges.EntityData.Children.Append(types.GetSegmentPath(trunkVlanRanges.Pushe[i]), types.YChild{"Pushe", trunkVlanRanges.Pushe[i]})
    }
    trunkVlanRanges.EntityData.Leafs = types.NewOrderedMap()
    trunkVlanRanges.EntityData.Leafs.Append("payload-ethertype", types.YLeaf{"PayloadEthertype", trunkVlanRanges.PayloadEthertype})
    trunkVlanRanges.EntityData.Leafs.Append("tags-popped", types.YLeaf{"TagsPopped", trunkVlanRanges.TagsPopped})
    trunkVlanRanges.EntityData.Leafs.Append("is-exact-match", types.YLeaf{"IsExactMatch", trunkVlanRanges.IsExactMatch})
    trunkVlanRanges.EntityData.Leafs.Append("is-native-vlan", types.YLeaf{"IsNativeVlan", trunkVlanRanges.IsNativeVlan})
    trunkVlanRanges.EntityData.Leafs.Append("is-native-preserving", types.YLeaf{"IsNativePreserving", trunkVlanRanges.IsNativePreserving})
    trunkVlanRanges.EntityData.Leafs.Append("source-mac-match", types.YLeaf{"SourceMacMatch", trunkVlanRanges.SourceMacMatch})
    trunkVlanRanges.EntityData.Leafs.Append("destination-mac-match", types.YLeaf{"DestinationMacMatch", trunkVlanRanges.DestinationMacMatch})

    trunkVlanRanges.EntityData.YListKeys = []string {}

    return &(trunkVlanRanges.EntityData)
}

// Vlan_Nodes_Node_Trunks_Trunk_VlanSwitched_TrunkVlanRanges_LocalTrafficStack
// VLAN tags for locally-sourced traffic
type Vlan_Nodes_Node_Trunks_Trunk_VlanSwitched_TrunkVlanRanges_LocalTrafficStack struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VLAN tags for locally-sourced traffic. The type is slice of
    // Vlan_Nodes_Node_Trunks_Trunk_VlanSwitched_TrunkVlanRanges_LocalTrafficStack_LocalTrafficTag.
    LocalTrafficTag []*Vlan_Nodes_Node_Trunks_Trunk_VlanSwitched_TrunkVlanRanges_LocalTrafficStack_LocalTrafficTag
}

func (localTrafficStack *Vlan_Nodes_Node_Trunks_Trunk_VlanSwitched_TrunkVlanRanges_LocalTrafficStack) GetEntityData() *types.CommonEntityData {
    localTrafficStack.EntityData.YFilter = localTrafficStack.YFilter
    localTrafficStack.EntityData.YangName = "local-traffic-stack"
    localTrafficStack.EntityData.BundleName = "cisco_ios_xr"
    localTrafficStack.EntityData.ParentYangName = "trunk-vlan-ranges"
    localTrafficStack.EntityData.SegmentPath = "local-traffic-stack"
    localTrafficStack.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localTrafficStack.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localTrafficStack.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localTrafficStack.EntityData.Children = types.NewOrderedMap()
    localTrafficStack.EntityData.Children.Append("local-traffic-tag", types.YChild{"LocalTrafficTag", nil})
    for i := range localTrafficStack.LocalTrafficTag {
        localTrafficStack.EntityData.Children.Append(types.GetSegmentPath(localTrafficStack.LocalTrafficTag[i]), types.YChild{"LocalTrafficTag", localTrafficStack.LocalTrafficTag[i]})
    }
    localTrafficStack.EntityData.Leafs = types.NewOrderedMap()

    localTrafficStack.EntityData.YListKeys = []string {}

    return &(localTrafficStack.EntityData)
}

// Vlan_Nodes_Node_Trunks_Trunk_VlanSwitched_TrunkVlanRanges_LocalTrafficStack_LocalTrafficTag
// VLAN tags for locally-sourced traffic
type Vlan_Nodes_Node_Trunks_Trunk_VlanSwitched_TrunkVlanRanges_LocalTrafficStack_LocalTrafficTag struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ethertype of tag. The type is EfpTagEtype.
    Ethertype interface{}

    // VLAN Id. The type is interface{} with range: 0..65535.
    VlanId interface{}
}

func (localTrafficTag *Vlan_Nodes_Node_Trunks_Trunk_VlanSwitched_TrunkVlanRanges_LocalTrafficStack_LocalTrafficTag) GetEntityData() *types.CommonEntityData {
    localTrafficTag.EntityData.YFilter = localTrafficTag.YFilter
    localTrafficTag.EntityData.YangName = "local-traffic-tag"
    localTrafficTag.EntityData.BundleName = "cisco_ios_xr"
    localTrafficTag.EntityData.ParentYangName = "local-traffic-stack"
    localTrafficTag.EntityData.SegmentPath = "local-traffic-tag"
    localTrafficTag.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localTrafficTag.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localTrafficTag.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localTrafficTag.EntityData.Children = types.NewOrderedMap()
    localTrafficTag.EntityData.Leafs = types.NewOrderedMap()
    localTrafficTag.EntityData.Leafs.Append("ethertype", types.YLeaf{"Ethertype", localTrafficTag.Ethertype})
    localTrafficTag.EntityData.Leafs.Append("vlan-id", types.YLeaf{"VlanId", localTrafficTag.VlanId})

    localTrafficTag.EntityData.YListKeys = []string {}

    return &(localTrafficTag.EntityData)
}

// Vlan_Nodes_Node_Trunks_Trunk_VlanSwitched_TrunkVlanRanges_TagsToMatch
// Tags to match on ingress packets
type Vlan_Nodes_Node_Trunks_Trunk_VlanSwitched_TrunkVlanRanges_TagsToMatch struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ethertype of tag to match. The type is EfpTagEtype.
    Ethertype interface{}

    // Priority to match. The type is EfpTagPriority.
    Priority interface{}

    // VLAN Ids to match. The type is slice of
    // Vlan_Nodes_Node_Trunks_Trunk_VlanSwitched_TrunkVlanRanges_TagsToMatch_VlanRange.
    VlanRange []*Vlan_Nodes_Node_Trunks_Trunk_VlanSwitched_TrunkVlanRanges_TagsToMatch_VlanRange
}

func (tagsToMatch *Vlan_Nodes_Node_Trunks_Trunk_VlanSwitched_TrunkVlanRanges_TagsToMatch) GetEntityData() *types.CommonEntityData {
    tagsToMatch.EntityData.YFilter = tagsToMatch.YFilter
    tagsToMatch.EntityData.YangName = "tags-to-match"
    tagsToMatch.EntityData.BundleName = "cisco_ios_xr"
    tagsToMatch.EntityData.ParentYangName = "trunk-vlan-ranges"
    tagsToMatch.EntityData.SegmentPath = "tags-to-match"
    tagsToMatch.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tagsToMatch.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tagsToMatch.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tagsToMatch.EntityData.Children = types.NewOrderedMap()
    tagsToMatch.EntityData.Children.Append("vlan-range", types.YChild{"VlanRange", nil})
    for i := range tagsToMatch.VlanRange {
        tagsToMatch.EntityData.Children.Append(types.GetSegmentPath(tagsToMatch.VlanRange[i]), types.YChild{"VlanRange", tagsToMatch.VlanRange[i]})
    }
    tagsToMatch.EntityData.Leafs = types.NewOrderedMap()
    tagsToMatch.EntityData.Leafs.Append("ethertype", types.YLeaf{"Ethertype", tagsToMatch.Ethertype})
    tagsToMatch.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", tagsToMatch.Priority})

    tagsToMatch.EntityData.YListKeys = []string {}

    return &(tagsToMatch.EntityData)
}

// Vlan_Nodes_Node_Trunks_Trunk_VlanSwitched_TrunkVlanRanges_TagsToMatch_VlanRange
// VLAN Ids to match
type Vlan_Nodes_Node_Trunks_Trunk_VlanSwitched_TrunkVlanRanges_TagsToMatch_VlanRange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VLAN ID Low. The type is interface{} with range: 0..65535.
    VlanIdLow interface{}

    // VLAN ID High. The type is interface{} with range: 0..65535.
    VlanIdHigh interface{}
}

func (vlanRange *Vlan_Nodes_Node_Trunks_Trunk_VlanSwitched_TrunkVlanRanges_TagsToMatch_VlanRange) GetEntityData() *types.CommonEntityData {
    vlanRange.EntityData.YFilter = vlanRange.YFilter
    vlanRange.EntityData.YangName = "vlan-range"
    vlanRange.EntityData.BundleName = "cisco_ios_xr"
    vlanRange.EntityData.ParentYangName = "tags-to-match"
    vlanRange.EntityData.SegmentPath = "vlan-range"
    vlanRange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vlanRange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vlanRange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vlanRange.EntityData.Children = types.NewOrderedMap()
    vlanRange.EntityData.Leafs = types.NewOrderedMap()
    vlanRange.EntityData.Leafs.Append("vlan-id-low", types.YLeaf{"VlanIdLow", vlanRange.VlanIdLow})
    vlanRange.EntityData.Leafs.Append("vlan-id-high", types.YLeaf{"VlanIdHigh", vlanRange.VlanIdHigh})

    vlanRange.EntityData.YListKeys = []string {}

    return &(vlanRange.EntityData)
}

// Vlan_Nodes_Node_Trunks_Trunk_VlanSwitched_TrunkVlanRanges_Pushe
// VLAN tags pushed on egress
type Vlan_Nodes_Node_Trunks_Trunk_VlanSwitched_TrunkVlanRanges_Pushe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ethertype of tag. The type is EfpTagEtype.
    Ethertype interface{}

    // VLAN Id. The type is interface{} with range: 0..65535.
    VlanId interface{}
}

func (pushe *Vlan_Nodes_Node_Trunks_Trunk_VlanSwitched_TrunkVlanRanges_Pushe) GetEntityData() *types.CommonEntityData {
    pushe.EntityData.YFilter = pushe.YFilter
    pushe.EntityData.YangName = "pushe"
    pushe.EntityData.BundleName = "cisco_ios_xr"
    pushe.EntityData.ParentYangName = "trunk-vlan-ranges"
    pushe.EntityData.SegmentPath = "pushe"
    pushe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pushe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pushe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pushe.EntityData.Children = types.NewOrderedMap()
    pushe.EntityData.Leafs = types.NewOrderedMap()
    pushe.EntityData.Leafs.Append("ethertype", types.YLeaf{"Ethertype", pushe.Ethertype})
    pushe.EntityData.Leafs.Append("vlan-id", types.YLeaf{"VlanId", pushe.VlanId})

    pushe.EntityData.YListKeys = []string {}

    return &(pushe.EntityData)
}

// Vlan_Nodes_Node_Interfaces
// VLAN interface table (specific to this node)
type Vlan_Nodes_Node_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational data for a sub-interface configured with VLANs. The type is
    // slice of Vlan_Nodes_Node_Interfaces_Interface.
    Interface []*Vlan_Nodes_Node_Interfaces_Interface
}

func (interfaces *Vlan_Nodes_Node_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "node"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = types.NewOrderedMap()
    interfaces.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range interfaces.Interface {
        interfaces.EntityData.Children.Append(types.GetSegmentPath(interfaces.Interface[i]), types.YChild{"Interface", interfaces.Interface[i]})
    }
    interfaces.EntityData.Leafs = types.NewOrderedMap()

    interfaces.EntityData.YListKeys = []string {}

    return &(interfaces.EntityData)
}

// Vlan_Nodes_Node_Interfaces_Interface
// Operational data for a sub-interface
// configured with VLANs
type Vlan_Nodes_Node_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The interface name. The type is string with
    // pattern: [a-zA-Z0-9._/-]+.
    Interface interface{}

    // Interface. The type is string with pattern: [a-zA-Z0-9._/-]+.
    InterfaceXr interface{}

    // Parent interface. The type is string with pattern: [a-zA-Z0-9._/-]+.
    ParentInterface interface{}

    // Service type. The type is VlanService.
    Service interface{}

    // Interface state. The type is ImStateEnum.
    State interface{}

    // L2 MTU. The type is interface{} with range: 0..65535.
    Mtu interface{}

    // L2 switched MTU. The type is interface{} with range: 0..65535.
    SwitchedMtu interface{}

    // Encapsulation type and tag stack.
    EncapsulationDetails Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails
}

func (self *Vlan_Nodes_Node_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.Interface, "interface")
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("encapsulation-details", types.YChild{"EncapsulationDetails", &self.EncapsulationDetails})
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", self.Interface})
    self.EntityData.Leafs.Append("interface-xr", types.YLeaf{"InterfaceXr", self.InterfaceXr})
    self.EntityData.Leafs.Append("parent-interface", types.YLeaf{"ParentInterface", self.ParentInterface})
    self.EntityData.Leafs.Append("service", types.YLeaf{"Service", self.Service})
    self.EntityData.Leafs.Append("state", types.YLeaf{"State", self.State})
    self.EntityData.Leafs.Append("mtu", types.YLeaf{"Mtu", self.Mtu})
    self.EntityData.Leafs.Append("switched-mtu", types.YLeaf{"SwitchedMtu", self.SwitchedMtu})

    self.EntityData.YListKeys = []string {"Interface"}

    return &(self.EntityData)
}

// Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails
// Encapsulation type and tag stack
type Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VLANEncapsulation. The type is VlanEncaps.
    VlanEncapsulation interface{}

    // Tag value. The type is interface{} with range: 0..65535.
    Tag interface{}

    // Outer tag value. The type is interface{} with range: 0..65535.
    OuterTag interface{}

    // Native tag value. The type is interface{} with range: 0..65535.
    NativeTag interface{}

    // 802.1ad tag value. The type is interface{} with range: 0..65535.
    Dot1adTag interface{}

    // 802.1ad native tag value. The type is interface{} with range: 0..65535.
    Dot1adNativeTag interface{}

    // 802.1ad Outer tag value. The type is interface{} with range: 0..65535.
    Dot1adOuterTag interface{}

    // Stack value.
    Stack Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_Stack

    // Service Instance encapsulation.
    ServiceInstanceDetails Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails

    // 802.1ad 802.1Q stack value.
    Dot1adDot1qStack Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_Dot1adDot1qStack
}

func (encapsulationDetails *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails) GetEntityData() *types.CommonEntityData {
    encapsulationDetails.EntityData.YFilter = encapsulationDetails.YFilter
    encapsulationDetails.EntityData.YangName = "encapsulation-details"
    encapsulationDetails.EntityData.BundleName = "cisco_ios_xr"
    encapsulationDetails.EntityData.ParentYangName = "interface"
    encapsulationDetails.EntityData.SegmentPath = "encapsulation-details"
    encapsulationDetails.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    encapsulationDetails.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    encapsulationDetails.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    encapsulationDetails.EntityData.Children = types.NewOrderedMap()
    encapsulationDetails.EntityData.Children.Append("stack", types.YChild{"Stack", &encapsulationDetails.Stack})
    encapsulationDetails.EntityData.Children.Append("service-instance-details", types.YChild{"ServiceInstanceDetails", &encapsulationDetails.ServiceInstanceDetails})
    encapsulationDetails.EntityData.Children.Append("dot1ad-dot1q-stack", types.YChild{"Dot1adDot1qStack", &encapsulationDetails.Dot1adDot1qStack})
    encapsulationDetails.EntityData.Leafs = types.NewOrderedMap()
    encapsulationDetails.EntityData.Leafs.Append("vlan-encapsulation", types.YLeaf{"VlanEncapsulation", encapsulationDetails.VlanEncapsulation})
    encapsulationDetails.EntityData.Leafs.Append("tag", types.YLeaf{"Tag", encapsulationDetails.Tag})
    encapsulationDetails.EntityData.Leafs.Append("outer-tag", types.YLeaf{"OuterTag", encapsulationDetails.OuterTag})
    encapsulationDetails.EntityData.Leafs.Append("native-tag", types.YLeaf{"NativeTag", encapsulationDetails.NativeTag})
    encapsulationDetails.EntityData.Leafs.Append("dot1ad-tag", types.YLeaf{"Dot1adTag", encapsulationDetails.Dot1adTag})
    encapsulationDetails.EntityData.Leafs.Append("dot1ad-native-tag", types.YLeaf{"Dot1adNativeTag", encapsulationDetails.Dot1adNativeTag})
    encapsulationDetails.EntityData.Leafs.Append("dot1ad-outer-tag", types.YLeaf{"Dot1adOuterTag", encapsulationDetails.Dot1adOuterTag})

    encapsulationDetails.EntityData.YListKeys = []string {}

    return &(encapsulationDetails.EntityData)
}

// Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_Stack
// Stack value
type Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_Stack struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Outer tag value. The type is interface{} with range: 0..65535.
    OuterTag interface{}

    // Second tag value. The type is interface{} with range: 0..65535.
    SecondTag interface{}
}

func (stack *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_Stack) GetEntityData() *types.CommonEntityData {
    stack.EntityData.YFilter = stack.YFilter
    stack.EntityData.YangName = "stack"
    stack.EntityData.BundleName = "cisco_ios_xr"
    stack.EntityData.ParentYangName = "encapsulation-details"
    stack.EntityData.SegmentPath = "stack"
    stack.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stack.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stack.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stack.EntityData.Children = types.NewOrderedMap()
    stack.EntityData.Leafs = types.NewOrderedMap()
    stack.EntityData.Leafs.Append("outer-tag", types.YLeaf{"OuterTag", stack.OuterTag})
    stack.EntityData.Leafs.Append("second-tag", types.YLeaf{"SecondTag", stack.SecondTag})

    stack.EntityData.YListKeys = []string {}

    return &(stack.EntityData)
}

// Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails
// Service Instance encapsulation
type Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Payload Ethertype to match. The type is EfpPayloadEtype.
    PayloadEthertype interface{}

    // Number of tags popped on ingress. The type is interface{} with range:
    // 0..65535.
    TagsPopped interface{}

    // Whether the packet must match the encapsulation exactly, with no further
    // inner tags. The type is bool.
    IsExactMatch interface{}

    // Whether this represents the native VLAN on the port. The type is bool.
    IsNativeVlan interface{}

    // Whether the native VLAN is customer-tag preserving. The type is bool.
    IsNativePreserving interface{}

    // The source MAC address to match on ingress. The type is string with
    // pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    SourceMacMatch interface{}

    // The destination MAC address to match on ingress. The type is string with
    // pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    DestinationMacMatch interface{}

    // VLAN tags for locally-sourced traffic.
    LocalTrafficStack Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack

    // Tags to match on ingress packets. The type is slice of
    // Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch.
    TagsToMatch []*Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch

    // VLAN tags pushed on egress. The type is slice of
    // Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_Pushe.
    Pushe []*Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_Pushe
}

func (serviceInstanceDetails *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails) GetEntityData() *types.CommonEntityData {
    serviceInstanceDetails.EntityData.YFilter = serviceInstanceDetails.YFilter
    serviceInstanceDetails.EntityData.YangName = "service-instance-details"
    serviceInstanceDetails.EntityData.BundleName = "cisco_ios_xr"
    serviceInstanceDetails.EntityData.ParentYangName = "encapsulation-details"
    serviceInstanceDetails.EntityData.SegmentPath = "service-instance-details"
    serviceInstanceDetails.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serviceInstanceDetails.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serviceInstanceDetails.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serviceInstanceDetails.EntityData.Children = types.NewOrderedMap()
    serviceInstanceDetails.EntityData.Children.Append("local-traffic-stack", types.YChild{"LocalTrafficStack", &serviceInstanceDetails.LocalTrafficStack})
    serviceInstanceDetails.EntityData.Children.Append("tags-to-match", types.YChild{"TagsToMatch", nil})
    for i := range serviceInstanceDetails.TagsToMatch {
        serviceInstanceDetails.EntityData.Children.Append(types.GetSegmentPath(serviceInstanceDetails.TagsToMatch[i]), types.YChild{"TagsToMatch", serviceInstanceDetails.TagsToMatch[i]})
    }
    serviceInstanceDetails.EntityData.Children.Append("pushe", types.YChild{"Pushe", nil})
    for i := range serviceInstanceDetails.Pushe {
        serviceInstanceDetails.EntityData.Children.Append(types.GetSegmentPath(serviceInstanceDetails.Pushe[i]), types.YChild{"Pushe", serviceInstanceDetails.Pushe[i]})
    }
    serviceInstanceDetails.EntityData.Leafs = types.NewOrderedMap()
    serviceInstanceDetails.EntityData.Leafs.Append("payload-ethertype", types.YLeaf{"PayloadEthertype", serviceInstanceDetails.PayloadEthertype})
    serviceInstanceDetails.EntityData.Leafs.Append("tags-popped", types.YLeaf{"TagsPopped", serviceInstanceDetails.TagsPopped})
    serviceInstanceDetails.EntityData.Leafs.Append("is-exact-match", types.YLeaf{"IsExactMatch", serviceInstanceDetails.IsExactMatch})
    serviceInstanceDetails.EntityData.Leafs.Append("is-native-vlan", types.YLeaf{"IsNativeVlan", serviceInstanceDetails.IsNativeVlan})
    serviceInstanceDetails.EntityData.Leafs.Append("is-native-preserving", types.YLeaf{"IsNativePreserving", serviceInstanceDetails.IsNativePreserving})
    serviceInstanceDetails.EntityData.Leafs.Append("source-mac-match", types.YLeaf{"SourceMacMatch", serviceInstanceDetails.SourceMacMatch})
    serviceInstanceDetails.EntityData.Leafs.Append("destination-mac-match", types.YLeaf{"DestinationMacMatch", serviceInstanceDetails.DestinationMacMatch})

    serviceInstanceDetails.EntityData.YListKeys = []string {}

    return &(serviceInstanceDetails.EntityData)
}

// Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack
// VLAN tags for locally-sourced traffic
type Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VLAN tags for locally-sourced traffic. The type is slice of
    // Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag.
    LocalTrafficTag []*Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag
}

func (localTrafficStack *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack) GetEntityData() *types.CommonEntityData {
    localTrafficStack.EntityData.YFilter = localTrafficStack.YFilter
    localTrafficStack.EntityData.YangName = "local-traffic-stack"
    localTrafficStack.EntityData.BundleName = "cisco_ios_xr"
    localTrafficStack.EntityData.ParentYangName = "service-instance-details"
    localTrafficStack.EntityData.SegmentPath = "local-traffic-stack"
    localTrafficStack.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localTrafficStack.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localTrafficStack.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localTrafficStack.EntityData.Children = types.NewOrderedMap()
    localTrafficStack.EntityData.Children.Append("local-traffic-tag", types.YChild{"LocalTrafficTag", nil})
    for i := range localTrafficStack.LocalTrafficTag {
        localTrafficStack.EntityData.Children.Append(types.GetSegmentPath(localTrafficStack.LocalTrafficTag[i]), types.YChild{"LocalTrafficTag", localTrafficStack.LocalTrafficTag[i]})
    }
    localTrafficStack.EntityData.Leafs = types.NewOrderedMap()

    localTrafficStack.EntityData.YListKeys = []string {}

    return &(localTrafficStack.EntityData)
}

// Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag
// VLAN tags for locally-sourced traffic
type Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ethertype of tag. The type is EfpTagEtype.
    Ethertype interface{}

    // VLAN Id. The type is interface{} with range: 0..65535.
    VlanId interface{}
}

func (localTrafficTag *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag) GetEntityData() *types.CommonEntityData {
    localTrafficTag.EntityData.YFilter = localTrafficTag.YFilter
    localTrafficTag.EntityData.YangName = "local-traffic-tag"
    localTrafficTag.EntityData.BundleName = "cisco_ios_xr"
    localTrafficTag.EntityData.ParentYangName = "local-traffic-stack"
    localTrafficTag.EntityData.SegmentPath = "local-traffic-tag"
    localTrafficTag.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localTrafficTag.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localTrafficTag.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localTrafficTag.EntityData.Children = types.NewOrderedMap()
    localTrafficTag.EntityData.Leafs = types.NewOrderedMap()
    localTrafficTag.EntityData.Leafs.Append("ethertype", types.YLeaf{"Ethertype", localTrafficTag.Ethertype})
    localTrafficTag.EntityData.Leafs.Append("vlan-id", types.YLeaf{"VlanId", localTrafficTag.VlanId})

    localTrafficTag.EntityData.YListKeys = []string {}

    return &(localTrafficTag.EntityData)
}

// Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch
// Tags to match on ingress packets
type Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ethertype of tag to match. The type is EfpTagEtype.
    Ethertype interface{}

    // Priority to match. The type is EfpTagPriority.
    Priority interface{}

    // VLAN Ids to match. The type is slice of
    // Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange.
    VlanRange []*Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange
}

func (tagsToMatch *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch) GetEntityData() *types.CommonEntityData {
    tagsToMatch.EntityData.YFilter = tagsToMatch.YFilter
    tagsToMatch.EntityData.YangName = "tags-to-match"
    tagsToMatch.EntityData.BundleName = "cisco_ios_xr"
    tagsToMatch.EntityData.ParentYangName = "service-instance-details"
    tagsToMatch.EntityData.SegmentPath = "tags-to-match"
    tagsToMatch.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tagsToMatch.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tagsToMatch.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tagsToMatch.EntityData.Children = types.NewOrderedMap()
    tagsToMatch.EntityData.Children.Append("vlan-range", types.YChild{"VlanRange", nil})
    for i := range tagsToMatch.VlanRange {
        tagsToMatch.EntityData.Children.Append(types.GetSegmentPath(tagsToMatch.VlanRange[i]), types.YChild{"VlanRange", tagsToMatch.VlanRange[i]})
    }
    tagsToMatch.EntityData.Leafs = types.NewOrderedMap()
    tagsToMatch.EntityData.Leafs.Append("ethertype", types.YLeaf{"Ethertype", tagsToMatch.Ethertype})
    tagsToMatch.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", tagsToMatch.Priority})

    tagsToMatch.EntityData.YListKeys = []string {}

    return &(tagsToMatch.EntityData)
}

// Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange
// VLAN Ids to match
type Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VLAN ID Low. The type is interface{} with range: 0..65535.
    VlanIdLow interface{}

    // VLAN ID High. The type is interface{} with range: 0..65535.
    VlanIdHigh interface{}
}

func (vlanRange *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange) GetEntityData() *types.CommonEntityData {
    vlanRange.EntityData.YFilter = vlanRange.YFilter
    vlanRange.EntityData.YangName = "vlan-range"
    vlanRange.EntityData.BundleName = "cisco_ios_xr"
    vlanRange.EntityData.ParentYangName = "tags-to-match"
    vlanRange.EntityData.SegmentPath = "vlan-range"
    vlanRange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vlanRange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vlanRange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vlanRange.EntityData.Children = types.NewOrderedMap()
    vlanRange.EntityData.Leafs = types.NewOrderedMap()
    vlanRange.EntityData.Leafs.Append("vlan-id-low", types.YLeaf{"VlanIdLow", vlanRange.VlanIdLow})
    vlanRange.EntityData.Leafs.Append("vlan-id-high", types.YLeaf{"VlanIdHigh", vlanRange.VlanIdHigh})

    vlanRange.EntityData.YListKeys = []string {}

    return &(vlanRange.EntityData)
}

// Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_Pushe
// VLAN tags pushed on egress
type Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_Pushe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ethertype of tag. The type is EfpTagEtype.
    Ethertype interface{}

    // VLAN Id. The type is interface{} with range: 0..65535.
    VlanId interface{}
}

func (pushe *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_Pushe) GetEntityData() *types.CommonEntityData {
    pushe.EntityData.YFilter = pushe.YFilter
    pushe.EntityData.YangName = "pushe"
    pushe.EntityData.BundleName = "cisco_ios_xr"
    pushe.EntityData.ParentYangName = "service-instance-details"
    pushe.EntityData.SegmentPath = "pushe"
    pushe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pushe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pushe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pushe.EntityData.Children = types.NewOrderedMap()
    pushe.EntityData.Leafs = types.NewOrderedMap()
    pushe.EntityData.Leafs.Append("ethertype", types.YLeaf{"Ethertype", pushe.Ethertype})
    pushe.EntityData.Leafs.Append("vlan-id", types.YLeaf{"VlanId", pushe.VlanId})

    pushe.EntityData.YListKeys = []string {}

    return &(pushe.EntityData)
}

// Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_Dot1adDot1qStack
// 802.1ad 802.1Q stack value
type Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_Dot1adDot1qStack struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Outer tag value. The type is interface{} with range: 0..65535.
    OuterTag interface{}

    // Second tag value. The type is interface{} with range: 0..65535.
    SecondTag interface{}
}

func (dot1adDot1qStack *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_Dot1adDot1qStack) GetEntityData() *types.CommonEntityData {
    dot1adDot1qStack.EntityData.YFilter = dot1adDot1qStack.YFilter
    dot1adDot1qStack.EntityData.YangName = "dot1ad-dot1q-stack"
    dot1adDot1qStack.EntityData.BundleName = "cisco_ios_xr"
    dot1adDot1qStack.EntityData.ParentYangName = "encapsulation-details"
    dot1adDot1qStack.EntityData.SegmentPath = "dot1ad-dot1q-stack"
    dot1adDot1qStack.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dot1adDot1qStack.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dot1adDot1qStack.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dot1adDot1qStack.EntityData.Children = types.NewOrderedMap()
    dot1adDot1qStack.EntityData.Leafs = types.NewOrderedMap()
    dot1adDot1qStack.EntityData.Leafs.Append("outer-tag", types.YLeaf{"OuterTag", dot1adDot1qStack.OuterTag})
    dot1adDot1qStack.EntityData.Leafs.Append("second-tag", types.YLeaf{"SecondTag", dot1adDot1qStack.SecondTag})

    dot1adDot1qStack.EntityData.YListKeys = []string {}

    return &(dot1adDot1qStack.EntityData)
}

// Vlan_Nodes_Node_TagAllocations
// VLAN tag allocation table (specific to this
// node)
type Vlan_Nodes_Node_TagAllocations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational data for a sub-interface configured with VLANs. The type is
    // slice of Vlan_Nodes_Node_TagAllocations_TagAllocation.
    TagAllocation []*Vlan_Nodes_Node_TagAllocations_TagAllocation
}

func (tagAllocations *Vlan_Nodes_Node_TagAllocations) GetEntityData() *types.CommonEntityData {
    tagAllocations.EntityData.YFilter = tagAllocations.YFilter
    tagAllocations.EntityData.YangName = "tag-allocations"
    tagAllocations.EntityData.BundleName = "cisco_ios_xr"
    tagAllocations.EntityData.ParentYangName = "node"
    tagAllocations.EntityData.SegmentPath = "tag-allocations"
    tagAllocations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tagAllocations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tagAllocations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tagAllocations.EntityData.Children = types.NewOrderedMap()
    tagAllocations.EntityData.Children.Append("tag-allocation", types.YChild{"TagAllocation", nil})
    for i := range tagAllocations.TagAllocation {
        tagAllocations.EntityData.Children.Append(types.GetSegmentPath(tagAllocations.TagAllocation[i]), types.YChild{"TagAllocation", tagAllocations.TagAllocation[i]})
    }
    tagAllocations.EntityData.Leafs = types.NewOrderedMap()

    tagAllocations.EntityData.YListKeys = []string {}

    return &(tagAllocations.EntityData)
}

// Vlan_Nodes_Node_TagAllocations_TagAllocation
// Operational data for a sub-interface
// configured with VLANs
type Vlan_Nodes_Node_TagAllocations_TagAllocation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The interface name. The type is string with pattern: [a-zA-Z0-9._/-]+.
    Interface interface{}

    // The first (outermost) tag. The type is interface{} with range: 1..4094.
    FirstTag interface{}

    // The second tag. The type is one of the following types: enumeration
    // VlanTagOrAny, or int with range: 1..4096.
    SecondTag interface{}

    // Interface. The type is string with pattern: [a-zA-Z0-9._/-]+.
    InterfaceXr interface{}

    // Parent interface. The type is string with pattern: [a-zA-Z0-9._/-]+.
    ParentInterface interface{}

    // Service type. The type is VlanService.
    Service interface{}

    // Interface state. The type is ImStateEnum.
    State interface{}

    // L2 MTU. The type is interface{} with range: 0..65535.
    Mtu interface{}

    // L2 switched MTU. The type is interface{} with range: 0..65535.
    SwitchedMtu interface{}

    // Encapsulation type and tag stack.
    EncapsulationDetails Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails
}

func (tagAllocation *Vlan_Nodes_Node_TagAllocations_TagAllocation) GetEntityData() *types.CommonEntityData {
    tagAllocation.EntityData.YFilter = tagAllocation.YFilter
    tagAllocation.EntityData.YangName = "tag-allocation"
    tagAllocation.EntityData.BundleName = "cisco_ios_xr"
    tagAllocation.EntityData.ParentYangName = "tag-allocations"
    tagAllocation.EntityData.SegmentPath = "tag-allocation"
    tagAllocation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tagAllocation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tagAllocation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tagAllocation.EntityData.Children = types.NewOrderedMap()
    tagAllocation.EntityData.Children.Append("encapsulation-details", types.YChild{"EncapsulationDetails", &tagAllocation.EncapsulationDetails})
    tagAllocation.EntityData.Leafs = types.NewOrderedMap()
    tagAllocation.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", tagAllocation.Interface})
    tagAllocation.EntityData.Leafs.Append("first-tag", types.YLeaf{"FirstTag", tagAllocation.FirstTag})
    tagAllocation.EntityData.Leafs.Append("second-tag", types.YLeaf{"SecondTag", tagAllocation.SecondTag})
    tagAllocation.EntityData.Leafs.Append("interface-xr", types.YLeaf{"InterfaceXr", tagAllocation.InterfaceXr})
    tagAllocation.EntityData.Leafs.Append("parent-interface", types.YLeaf{"ParentInterface", tagAllocation.ParentInterface})
    tagAllocation.EntityData.Leafs.Append("service", types.YLeaf{"Service", tagAllocation.Service})
    tagAllocation.EntityData.Leafs.Append("state", types.YLeaf{"State", tagAllocation.State})
    tagAllocation.EntityData.Leafs.Append("mtu", types.YLeaf{"Mtu", tagAllocation.Mtu})
    tagAllocation.EntityData.Leafs.Append("switched-mtu", types.YLeaf{"SwitchedMtu", tagAllocation.SwitchedMtu})

    tagAllocation.EntityData.YListKeys = []string {}

    return &(tagAllocation.EntityData)
}

// Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails
// Encapsulation type and tag stack
type Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VLANEncapsulation. The type is VlanEncaps.
    VlanEncapsulation interface{}

    // Tag value. The type is interface{} with range: 0..65535.
    Tag interface{}

    // Outer tag value. The type is interface{} with range: 0..65535.
    OuterTag interface{}

    // Native tag value. The type is interface{} with range: 0..65535.
    NativeTag interface{}

    // 802.1ad tag value. The type is interface{} with range: 0..65535.
    Dot1adTag interface{}

    // 802.1ad native tag value. The type is interface{} with range: 0..65535.
    Dot1adNativeTag interface{}

    // 802.1ad Outer tag value. The type is interface{} with range: 0..65535.
    Dot1adOuterTag interface{}

    // Stack value.
    Stack Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_Stack

    // Service Instance encapsulation.
    ServiceInstanceDetails Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails

    // 802.1ad 802.1Q stack value.
    Dot1adDot1qStack Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_Dot1adDot1qStack
}

func (encapsulationDetails *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails) GetEntityData() *types.CommonEntityData {
    encapsulationDetails.EntityData.YFilter = encapsulationDetails.YFilter
    encapsulationDetails.EntityData.YangName = "encapsulation-details"
    encapsulationDetails.EntityData.BundleName = "cisco_ios_xr"
    encapsulationDetails.EntityData.ParentYangName = "tag-allocation"
    encapsulationDetails.EntityData.SegmentPath = "encapsulation-details"
    encapsulationDetails.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    encapsulationDetails.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    encapsulationDetails.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    encapsulationDetails.EntityData.Children = types.NewOrderedMap()
    encapsulationDetails.EntityData.Children.Append("stack", types.YChild{"Stack", &encapsulationDetails.Stack})
    encapsulationDetails.EntityData.Children.Append("service-instance-details", types.YChild{"ServiceInstanceDetails", &encapsulationDetails.ServiceInstanceDetails})
    encapsulationDetails.EntityData.Children.Append("dot1ad-dot1q-stack", types.YChild{"Dot1adDot1qStack", &encapsulationDetails.Dot1adDot1qStack})
    encapsulationDetails.EntityData.Leafs = types.NewOrderedMap()
    encapsulationDetails.EntityData.Leafs.Append("vlan-encapsulation", types.YLeaf{"VlanEncapsulation", encapsulationDetails.VlanEncapsulation})
    encapsulationDetails.EntityData.Leafs.Append("tag", types.YLeaf{"Tag", encapsulationDetails.Tag})
    encapsulationDetails.EntityData.Leafs.Append("outer-tag", types.YLeaf{"OuterTag", encapsulationDetails.OuterTag})
    encapsulationDetails.EntityData.Leafs.Append("native-tag", types.YLeaf{"NativeTag", encapsulationDetails.NativeTag})
    encapsulationDetails.EntityData.Leafs.Append("dot1ad-tag", types.YLeaf{"Dot1adTag", encapsulationDetails.Dot1adTag})
    encapsulationDetails.EntityData.Leafs.Append("dot1ad-native-tag", types.YLeaf{"Dot1adNativeTag", encapsulationDetails.Dot1adNativeTag})
    encapsulationDetails.EntityData.Leafs.Append("dot1ad-outer-tag", types.YLeaf{"Dot1adOuterTag", encapsulationDetails.Dot1adOuterTag})

    encapsulationDetails.EntityData.YListKeys = []string {}

    return &(encapsulationDetails.EntityData)
}

// Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_Stack
// Stack value
type Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_Stack struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Outer tag value. The type is interface{} with range: 0..65535.
    OuterTag interface{}

    // Second tag value. The type is interface{} with range: 0..65535.
    SecondTag interface{}
}

func (stack *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_Stack) GetEntityData() *types.CommonEntityData {
    stack.EntityData.YFilter = stack.YFilter
    stack.EntityData.YangName = "stack"
    stack.EntityData.BundleName = "cisco_ios_xr"
    stack.EntityData.ParentYangName = "encapsulation-details"
    stack.EntityData.SegmentPath = "stack"
    stack.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stack.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stack.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stack.EntityData.Children = types.NewOrderedMap()
    stack.EntityData.Leafs = types.NewOrderedMap()
    stack.EntityData.Leafs.Append("outer-tag", types.YLeaf{"OuterTag", stack.OuterTag})
    stack.EntityData.Leafs.Append("second-tag", types.YLeaf{"SecondTag", stack.SecondTag})

    stack.EntityData.YListKeys = []string {}

    return &(stack.EntityData)
}

// Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails
// Service Instance encapsulation
type Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Payload Ethertype to match. The type is EfpPayloadEtype.
    PayloadEthertype interface{}

    // Number of tags popped on ingress. The type is interface{} with range:
    // 0..65535.
    TagsPopped interface{}

    // Whether the packet must match the encapsulation exactly, with no further
    // inner tags. The type is bool.
    IsExactMatch interface{}

    // Whether this represents the native VLAN on the port. The type is bool.
    IsNativeVlan interface{}

    // Whether the native VLAN is customer-tag preserving. The type is bool.
    IsNativePreserving interface{}

    // The source MAC address to match on ingress. The type is string with
    // pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    SourceMacMatch interface{}

    // The destination MAC address to match on ingress. The type is string with
    // pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    DestinationMacMatch interface{}

    // VLAN tags for locally-sourced traffic.
    LocalTrafficStack Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack

    // Tags to match on ingress packets. The type is slice of
    // Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch.
    TagsToMatch []*Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch

    // VLAN tags pushed on egress. The type is slice of
    // Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_Pushe.
    Pushe []*Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_Pushe
}

func (serviceInstanceDetails *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails) GetEntityData() *types.CommonEntityData {
    serviceInstanceDetails.EntityData.YFilter = serviceInstanceDetails.YFilter
    serviceInstanceDetails.EntityData.YangName = "service-instance-details"
    serviceInstanceDetails.EntityData.BundleName = "cisco_ios_xr"
    serviceInstanceDetails.EntityData.ParentYangName = "encapsulation-details"
    serviceInstanceDetails.EntityData.SegmentPath = "service-instance-details"
    serviceInstanceDetails.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serviceInstanceDetails.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serviceInstanceDetails.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serviceInstanceDetails.EntityData.Children = types.NewOrderedMap()
    serviceInstanceDetails.EntityData.Children.Append("local-traffic-stack", types.YChild{"LocalTrafficStack", &serviceInstanceDetails.LocalTrafficStack})
    serviceInstanceDetails.EntityData.Children.Append("tags-to-match", types.YChild{"TagsToMatch", nil})
    for i := range serviceInstanceDetails.TagsToMatch {
        serviceInstanceDetails.EntityData.Children.Append(types.GetSegmentPath(serviceInstanceDetails.TagsToMatch[i]), types.YChild{"TagsToMatch", serviceInstanceDetails.TagsToMatch[i]})
    }
    serviceInstanceDetails.EntityData.Children.Append("pushe", types.YChild{"Pushe", nil})
    for i := range serviceInstanceDetails.Pushe {
        serviceInstanceDetails.EntityData.Children.Append(types.GetSegmentPath(serviceInstanceDetails.Pushe[i]), types.YChild{"Pushe", serviceInstanceDetails.Pushe[i]})
    }
    serviceInstanceDetails.EntityData.Leafs = types.NewOrderedMap()
    serviceInstanceDetails.EntityData.Leafs.Append("payload-ethertype", types.YLeaf{"PayloadEthertype", serviceInstanceDetails.PayloadEthertype})
    serviceInstanceDetails.EntityData.Leafs.Append("tags-popped", types.YLeaf{"TagsPopped", serviceInstanceDetails.TagsPopped})
    serviceInstanceDetails.EntityData.Leafs.Append("is-exact-match", types.YLeaf{"IsExactMatch", serviceInstanceDetails.IsExactMatch})
    serviceInstanceDetails.EntityData.Leafs.Append("is-native-vlan", types.YLeaf{"IsNativeVlan", serviceInstanceDetails.IsNativeVlan})
    serviceInstanceDetails.EntityData.Leafs.Append("is-native-preserving", types.YLeaf{"IsNativePreserving", serviceInstanceDetails.IsNativePreserving})
    serviceInstanceDetails.EntityData.Leafs.Append("source-mac-match", types.YLeaf{"SourceMacMatch", serviceInstanceDetails.SourceMacMatch})
    serviceInstanceDetails.EntityData.Leafs.Append("destination-mac-match", types.YLeaf{"DestinationMacMatch", serviceInstanceDetails.DestinationMacMatch})

    serviceInstanceDetails.EntityData.YListKeys = []string {}

    return &(serviceInstanceDetails.EntityData)
}

// Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack
// VLAN tags for locally-sourced traffic
type Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VLAN tags for locally-sourced traffic. The type is slice of
    // Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag.
    LocalTrafficTag []*Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag
}

func (localTrafficStack *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack) GetEntityData() *types.CommonEntityData {
    localTrafficStack.EntityData.YFilter = localTrafficStack.YFilter
    localTrafficStack.EntityData.YangName = "local-traffic-stack"
    localTrafficStack.EntityData.BundleName = "cisco_ios_xr"
    localTrafficStack.EntityData.ParentYangName = "service-instance-details"
    localTrafficStack.EntityData.SegmentPath = "local-traffic-stack"
    localTrafficStack.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localTrafficStack.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localTrafficStack.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localTrafficStack.EntityData.Children = types.NewOrderedMap()
    localTrafficStack.EntityData.Children.Append("local-traffic-tag", types.YChild{"LocalTrafficTag", nil})
    for i := range localTrafficStack.LocalTrafficTag {
        localTrafficStack.EntityData.Children.Append(types.GetSegmentPath(localTrafficStack.LocalTrafficTag[i]), types.YChild{"LocalTrafficTag", localTrafficStack.LocalTrafficTag[i]})
    }
    localTrafficStack.EntityData.Leafs = types.NewOrderedMap()

    localTrafficStack.EntityData.YListKeys = []string {}

    return &(localTrafficStack.EntityData)
}

// Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag
// VLAN tags for locally-sourced traffic
type Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ethertype of tag. The type is EfpTagEtype.
    Ethertype interface{}

    // VLAN Id. The type is interface{} with range: 0..65535.
    VlanId interface{}
}

func (localTrafficTag *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag) GetEntityData() *types.CommonEntityData {
    localTrafficTag.EntityData.YFilter = localTrafficTag.YFilter
    localTrafficTag.EntityData.YangName = "local-traffic-tag"
    localTrafficTag.EntityData.BundleName = "cisco_ios_xr"
    localTrafficTag.EntityData.ParentYangName = "local-traffic-stack"
    localTrafficTag.EntityData.SegmentPath = "local-traffic-tag"
    localTrafficTag.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localTrafficTag.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localTrafficTag.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localTrafficTag.EntityData.Children = types.NewOrderedMap()
    localTrafficTag.EntityData.Leafs = types.NewOrderedMap()
    localTrafficTag.EntityData.Leafs.Append("ethertype", types.YLeaf{"Ethertype", localTrafficTag.Ethertype})
    localTrafficTag.EntityData.Leafs.Append("vlan-id", types.YLeaf{"VlanId", localTrafficTag.VlanId})

    localTrafficTag.EntityData.YListKeys = []string {}

    return &(localTrafficTag.EntityData)
}

// Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch
// Tags to match on ingress packets
type Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ethertype of tag to match. The type is EfpTagEtype.
    Ethertype interface{}

    // Priority to match. The type is EfpTagPriority.
    Priority interface{}

    // VLAN Ids to match. The type is slice of
    // Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange.
    VlanRange []*Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange
}

func (tagsToMatch *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch) GetEntityData() *types.CommonEntityData {
    tagsToMatch.EntityData.YFilter = tagsToMatch.YFilter
    tagsToMatch.EntityData.YangName = "tags-to-match"
    tagsToMatch.EntityData.BundleName = "cisco_ios_xr"
    tagsToMatch.EntityData.ParentYangName = "service-instance-details"
    tagsToMatch.EntityData.SegmentPath = "tags-to-match"
    tagsToMatch.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tagsToMatch.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tagsToMatch.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tagsToMatch.EntityData.Children = types.NewOrderedMap()
    tagsToMatch.EntityData.Children.Append("vlan-range", types.YChild{"VlanRange", nil})
    for i := range tagsToMatch.VlanRange {
        tagsToMatch.EntityData.Children.Append(types.GetSegmentPath(tagsToMatch.VlanRange[i]), types.YChild{"VlanRange", tagsToMatch.VlanRange[i]})
    }
    tagsToMatch.EntityData.Leafs = types.NewOrderedMap()
    tagsToMatch.EntityData.Leafs.Append("ethertype", types.YLeaf{"Ethertype", tagsToMatch.Ethertype})
    tagsToMatch.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", tagsToMatch.Priority})

    tagsToMatch.EntityData.YListKeys = []string {}

    return &(tagsToMatch.EntityData)
}

// Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange
// VLAN Ids to match
type Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VLAN ID Low. The type is interface{} with range: 0..65535.
    VlanIdLow interface{}

    // VLAN ID High. The type is interface{} with range: 0..65535.
    VlanIdHigh interface{}
}

func (vlanRange *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange) GetEntityData() *types.CommonEntityData {
    vlanRange.EntityData.YFilter = vlanRange.YFilter
    vlanRange.EntityData.YangName = "vlan-range"
    vlanRange.EntityData.BundleName = "cisco_ios_xr"
    vlanRange.EntityData.ParentYangName = "tags-to-match"
    vlanRange.EntityData.SegmentPath = "vlan-range"
    vlanRange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vlanRange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vlanRange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vlanRange.EntityData.Children = types.NewOrderedMap()
    vlanRange.EntityData.Leafs = types.NewOrderedMap()
    vlanRange.EntityData.Leafs.Append("vlan-id-low", types.YLeaf{"VlanIdLow", vlanRange.VlanIdLow})
    vlanRange.EntityData.Leafs.Append("vlan-id-high", types.YLeaf{"VlanIdHigh", vlanRange.VlanIdHigh})

    vlanRange.EntityData.YListKeys = []string {}

    return &(vlanRange.EntityData)
}

// Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_Pushe
// VLAN tags pushed on egress
type Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_Pushe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ethertype of tag. The type is EfpTagEtype.
    Ethertype interface{}

    // VLAN Id. The type is interface{} with range: 0..65535.
    VlanId interface{}
}

func (pushe *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_Pushe) GetEntityData() *types.CommonEntityData {
    pushe.EntityData.YFilter = pushe.YFilter
    pushe.EntityData.YangName = "pushe"
    pushe.EntityData.BundleName = "cisco_ios_xr"
    pushe.EntityData.ParentYangName = "service-instance-details"
    pushe.EntityData.SegmentPath = "pushe"
    pushe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pushe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pushe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pushe.EntityData.Children = types.NewOrderedMap()
    pushe.EntityData.Leafs = types.NewOrderedMap()
    pushe.EntityData.Leafs.Append("ethertype", types.YLeaf{"Ethertype", pushe.Ethertype})
    pushe.EntityData.Leafs.Append("vlan-id", types.YLeaf{"VlanId", pushe.VlanId})

    pushe.EntityData.YListKeys = []string {}

    return &(pushe.EntityData)
}

// Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_Dot1adDot1qStack
// 802.1ad 802.1Q stack value
type Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_Dot1adDot1qStack struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Outer tag value. The type is interface{} with range: 0..65535.
    OuterTag interface{}

    // Second tag value. The type is interface{} with range: 0..65535.
    SecondTag interface{}
}

func (dot1adDot1qStack *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_Dot1adDot1qStack) GetEntityData() *types.CommonEntityData {
    dot1adDot1qStack.EntityData.YFilter = dot1adDot1qStack.YFilter
    dot1adDot1qStack.EntityData.YangName = "dot1ad-dot1q-stack"
    dot1adDot1qStack.EntityData.BundleName = "cisco_ios_xr"
    dot1adDot1qStack.EntityData.ParentYangName = "encapsulation-details"
    dot1adDot1qStack.EntityData.SegmentPath = "dot1ad-dot1q-stack"
    dot1adDot1qStack.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dot1adDot1qStack.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dot1adDot1qStack.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dot1adDot1qStack.EntityData.Children = types.NewOrderedMap()
    dot1adDot1qStack.EntityData.Leafs = types.NewOrderedMap()
    dot1adDot1qStack.EntityData.Leafs.Append("outer-tag", types.YLeaf{"OuterTag", dot1adDot1qStack.OuterTag})
    dot1adDot1qStack.EntityData.Leafs.Append("second-tag", types.YLeaf{"SecondTag", dot1adDot1qStack.SecondTag})

    dot1adDot1qStack.EntityData.YListKeys = []string {}

    return &(dot1adDot1qStack.EntityData)
}

// EthernetEncapsulation
// ethernet encapsulation
type EthernetEncapsulation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Per node Ethernet encapsulation operational data.
    Nodes EthernetEncapsulation_Nodes
}

func (ethernetEncapsulation *EthernetEncapsulation) GetEntityData() *types.CommonEntityData {
    ethernetEncapsulation.EntityData.YFilter = ethernetEncapsulation.YFilter
    ethernetEncapsulation.EntityData.YangName = "ethernet-encapsulation"
    ethernetEncapsulation.EntityData.BundleName = "cisco_ios_xr"
    ethernetEncapsulation.EntityData.ParentYangName = "Cisco-IOS-XR-l2-eth-infra-oper"
    ethernetEncapsulation.EntityData.SegmentPath = "Cisco-IOS-XR-l2-eth-infra-oper:ethernet-encapsulation"
    ethernetEncapsulation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ethernetEncapsulation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ethernetEncapsulation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ethernetEncapsulation.EntityData.Children = types.NewOrderedMap()
    ethernetEncapsulation.EntityData.Children.Append("nodes", types.YChild{"Nodes", &ethernetEncapsulation.Nodes})
    ethernetEncapsulation.EntityData.Leafs = types.NewOrderedMap()

    ethernetEncapsulation.EntityData.YListKeys = []string {}

    return &(ethernetEncapsulation.EntityData)
}

// EthernetEncapsulation_Nodes
// Per node Ethernet encapsulation operational data
type EthernetEncapsulation_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The Ethernet encaps operational data for a particular node. The type is
    // slice of EthernetEncapsulation_Nodes_Node.
    Node []*EthernetEncapsulation_Nodes_Node
}

func (nodes *EthernetEncapsulation_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "ethernet-encapsulation"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodes.EntityData.Children = types.NewOrderedMap()
    nodes.EntityData.Children.Append("node", types.YChild{"Node", nil})
    for i := range nodes.Node {
        nodes.EntityData.Children.Append(types.GetSegmentPath(nodes.Node[i]), types.YChild{"Node", nodes.Node[i]})
    }
    nodes.EntityData.Leafs = types.NewOrderedMap()

    nodes.EntityData.YListKeys = []string {}

    return &(nodes.EntityData)
}

// EthernetEncapsulation_Nodes_Node
// The Ethernet encaps operational data for a
// particular node
type EthernetEncapsulation_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The identifier for the node. The type is string
    // with pattern: ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Unicast MAC filter table (specific to this node).
    UnicastMacFilters EthernetEncapsulation_Nodes_Node_UnicastMacFilters
}

func (node *EthernetEncapsulation_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeName, "node-name")
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("unicast-mac-filters", types.YChild{"UnicastMacFilters", &node.UnicastMacFilters})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", node.NodeName})

    node.EntityData.YListKeys = []string {"NodeName"}

    return &(node.EntityData)
}

// EthernetEncapsulation_Nodes_Node_UnicastMacFilters
// Unicast MAC filter table (specific to this
// node)
type EthernetEncapsulation_Nodes_Node_UnicastMacFilters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational data for interface with MAC filters configured. The type is
    // slice of
    // EthernetEncapsulation_Nodes_Node_UnicastMacFilters_UnicastMacFilter.
    UnicastMacFilter []*EthernetEncapsulation_Nodes_Node_UnicastMacFilters_UnicastMacFilter
}

func (unicastMacFilters *EthernetEncapsulation_Nodes_Node_UnicastMacFilters) GetEntityData() *types.CommonEntityData {
    unicastMacFilters.EntityData.YFilter = unicastMacFilters.YFilter
    unicastMacFilters.EntityData.YangName = "unicast-mac-filters"
    unicastMacFilters.EntityData.BundleName = "cisco_ios_xr"
    unicastMacFilters.EntityData.ParentYangName = "node"
    unicastMacFilters.EntityData.SegmentPath = "unicast-mac-filters"
    unicastMacFilters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unicastMacFilters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unicastMacFilters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unicastMacFilters.EntityData.Children = types.NewOrderedMap()
    unicastMacFilters.EntityData.Children.Append("unicast-mac-filter", types.YChild{"UnicastMacFilter", nil})
    for i := range unicastMacFilters.UnicastMacFilter {
        unicastMacFilters.EntityData.Children.Append(types.GetSegmentPath(unicastMacFilters.UnicastMacFilter[i]), types.YChild{"UnicastMacFilter", unicastMacFilters.UnicastMacFilter[i]})
    }
    unicastMacFilters.EntityData.Leafs = types.NewOrderedMap()

    unicastMacFilters.EntityData.YListKeys = []string {}

    return &(unicastMacFilters.EntityData)
}

// EthernetEncapsulation_Nodes_Node_UnicastMacFilters_UnicastMacFilter
// Operational data for interface with MAC
// filters configured
type EthernetEncapsulation_Nodes_Node_UnicastMacFilters_UnicastMacFilter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The interface name. The type is string with
    // pattern: [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // Unicast MAC filter information. The type is slice of
    // EthernetEncapsulation_Nodes_Node_UnicastMacFilters_UnicastMacFilter_UnicastFilter.
    UnicastFilter []*EthernetEncapsulation_Nodes_Node_UnicastMacFilters_UnicastMacFilter_UnicastFilter
}

func (unicastMacFilter *EthernetEncapsulation_Nodes_Node_UnicastMacFilters_UnicastMacFilter) GetEntityData() *types.CommonEntityData {
    unicastMacFilter.EntityData.YFilter = unicastMacFilter.YFilter
    unicastMacFilter.EntityData.YangName = "unicast-mac-filter"
    unicastMacFilter.EntityData.BundleName = "cisco_ios_xr"
    unicastMacFilter.EntityData.ParentYangName = "unicast-mac-filters"
    unicastMacFilter.EntityData.SegmentPath = "unicast-mac-filter" + types.AddKeyToken(unicastMacFilter.InterfaceName, "interface-name")
    unicastMacFilter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unicastMacFilter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unicastMacFilter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unicastMacFilter.EntityData.Children = types.NewOrderedMap()
    unicastMacFilter.EntityData.Children.Append("unicast-filter", types.YChild{"UnicastFilter", nil})
    for i := range unicastMacFilter.UnicastFilter {
        unicastMacFilter.EntityData.Children.Append(types.GetSegmentPath(unicastMacFilter.UnicastFilter[i]), types.YChild{"UnicastFilter", unicastMacFilter.UnicastFilter[i]})
    }
    unicastMacFilter.EntityData.Leafs = types.NewOrderedMap()
    unicastMacFilter.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", unicastMacFilter.InterfaceName})

    unicastMacFilter.EntityData.YListKeys = []string {"InterfaceName"}

    return &(unicastMacFilter.EntityData)
}

// EthernetEncapsulation_Nodes_Node_UnicastMacFilters_UnicastMacFilter_UnicastFilter
// Unicast MAC filter information
type EthernetEncapsulation_Nodes_Node_UnicastMacFilters_UnicastMacFilter_UnicastFilter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}

    // Unicast MAC mode. The type is EthCapsUcastMacMode.
    Mode interface{}
}

func (unicastFilter *EthernetEncapsulation_Nodes_Node_UnicastMacFilters_UnicastMacFilter_UnicastFilter) GetEntityData() *types.CommonEntityData {
    unicastFilter.EntityData.YFilter = unicastFilter.YFilter
    unicastFilter.EntityData.YangName = "unicast-filter"
    unicastFilter.EntityData.BundleName = "cisco_ios_xr"
    unicastFilter.EntityData.ParentYangName = "unicast-mac-filter"
    unicastFilter.EntityData.SegmentPath = "unicast-filter"
    unicastFilter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unicastFilter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unicastFilter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unicastFilter.EntityData.Children = types.NewOrderedMap()
    unicastFilter.EntityData.Leafs = types.NewOrderedMap()
    unicastFilter.EntityData.Leafs.Append("mac-address", types.YLeaf{"MacAddress", unicastFilter.MacAddress})
    unicastFilter.EntityData.Leafs.Append("mode", types.YLeaf{"Mode", unicastFilter.Mode})

    unicastFilter.EntityData.YListKeys = []string {}

    return &(unicastFilter.EntityData)
}

