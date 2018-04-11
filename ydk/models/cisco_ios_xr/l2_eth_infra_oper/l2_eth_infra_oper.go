// This module contains a collection of YANG definitions
// for Cisco IOS-XR l2-eth-infra package operational data.
// 
// This module contains definitions
// for the following management objects:
//   mac-accounting: MAC accounting operational data
//   vlan: vlan
//   ethernet-encapsulation: ethernet encapsulation
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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

// EthCapsUcastMacMode represents Eth caps ucast mac mode
type EthCapsUcastMacMode string

const (
    // Reserved
    EthCapsUcastMacMode_reserved EthCapsUcastMacMode = "reserved"

    // Permit
    EthCapsUcastMacMode_permit EthCapsUcastMacMode = "permit"
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

    macAccounting.EntityData.Children = make(map[string]types.YChild)
    macAccounting.EntityData.Children["interfaces"] = types.YChild{"Interfaces", &macAccounting.Interfaces}
    macAccounting.EntityData.Leafs = make(map[string]types.YLeaf)
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
    // MacAccounting_Interfaces_Interface_.
    Interface_ []MacAccounting_Interfaces_Interface
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

    interfaces.EntityData.Children = make(map[string]types.YChild)
    interfaces.EntityData.Children["interface"] = types.YChild{"Interface_", nil}
    for i := range interfaces.Interface_ {
        interfaces.EntityData.Children[types.GetSegmentPath(&interfaces.Interface_[i])] = types.YChild{"Interface_", &interfaces.Interface_[i]}
    }
    interfaces.EntityData.Leafs = make(map[string]types.YLeaf)
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
    // pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // MAC accounting state for the interface.
    State MacAccounting_Interfaces_Interface_State

    // Ingress MAC accounting statistics. The type is slice of
    // MacAccounting_Interfaces_Interface_IngressStatistic.
    IngressStatistic []MacAccounting_Interfaces_Interface_IngressStatistic

    // Egress MAC accounting statistics. The type is slice of
    // MacAccounting_Interfaces_Interface_EgressStatistic.
    EgressStatistic []MacAccounting_Interfaces_Interface_EgressStatistic
}

func (self *MacAccounting_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Children["state"] = types.YChild{"State", &self.State}
    self.EntityData.Children["ingress-statistic"] = types.YChild{"IngressStatistic", nil}
    for i := range self.IngressStatistic {
        self.EntityData.Children[types.GetSegmentPath(&self.IngressStatistic[i])] = types.YChild{"IngressStatistic", &self.IngressStatistic[i]}
    }
    self.EntityData.Children["egress-statistic"] = types.YChild{"EgressStatistic", nil}
    for i := range self.EgressStatistic {
        self.EntityData.Children[types.GetSegmentPath(&self.EgressStatistic[i])] = types.YChild{"EgressStatistic", &self.EgressStatistic[i]}
    }
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", self.InterfaceName}
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

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["is-ingress-enabled"] = types.YLeaf{"IsIngressEnabled", state.IsIngressEnabled}
    state.EntityData.Leafs["is-egress-enabled"] = types.YLeaf{"IsEgressEnabled", state.IsEgressEnabled}
    state.EntityData.Leafs["number-available-ingress"] = types.YLeaf{"NumberAvailableIngress", state.NumberAvailableIngress}
    state.EntityData.Leafs["number-available-egress"] = types.YLeaf{"NumberAvailableEgress", state.NumberAvailableEgress}
    state.EntityData.Leafs["number-available-on-node"] = types.YLeaf{"NumberAvailableOnNode", state.NumberAvailableOnNode}
    return &(state.EntityData)
}

// MacAccounting_Interfaces_Interface_IngressStatistic
// Ingress MAC accounting statistics
type MacAccounting_Interfaces_Interface_IngressStatistic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // 48bit MAC address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
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

    ingressStatistic.EntityData.Children = make(map[string]types.YChild)
    ingressStatistic.EntityData.Leafs = make(map[string]types.YLeaf)
    ingressStatistic.EntityData.Leafs["mac-address"] = types.YLeaf{"MacAddress", ingressStatistic.MacAddress}
    ingressStatistic.EntityData.Leafs["packets"] = types.YLeaf{"Packets", ingressStatistic.Packets}
    ingressStatistic.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", ingressStatistic.Bytes}
    return &(ingressStatistic.EntityData)
}

// MacAccounting_Interfaces_Interface_EgressStatistic
// Egress MAC accounting statistics
type MacAccounting_Interfaces_Interface_EgressStatistic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // 48bit MAC address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
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

    egressStatistic.EntityData.Children = make(map[string]types.YChild)
    egressStatistic.EntityData.Leafs = make(map[string]types.YLeaf)
    egressStatistic.EntityData.Leafs["mac-address"] = types.YLeaf{"MacAddress", egressStatistic.MacAddress}
    egressStatistic.EntityData.Leafs["packets"] = types.YLeaf{"Packets", egressStatistic.Packets}
    egressStatistic.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", egressStatistic.Bytes}
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

    vlan.EntityData.Children = make(map[string]types.YChild)
    vlan.EntityData.Children["nodes"] = types.YChild{"Nodes", &vlan.Nodes}
    vlan.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(vlan.EntityData)
}

// Vlan_Nodes
// Per node VLAN operational data
type Vlan_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The VLAN operational data for a particular node. The type is slice of
    // Vlan_Nodes_Node.
    Node []Vlan_Nodes_Node
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

    nodes.EntityData.Children = make(map[string]types.YChild)
    nodes.EntityData.Children["node"] = types.YChild{"Node", nil}
    for i := range nodes.Node {
        nodes.EntityData.Children[types.GetSegmentPath(&nodes.Node[i])] = types.YChild{"Node", &nodes.Node[i]}
    }
    nodes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nodes.EntityData)
}

// Vlan_Nodes_Node
// The VLAN operational data for a particular node
type Vlan_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The identifier for the node. The type is string
    // with pattern: b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
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
    node.EntityData.SegmentPath = "node" + "[node-id='" + fmt.Sprintf("%v", node.NodeId) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Children["trunks"] = types.YChild{"Trunks", &node.Trunks}
    node.EntityData.Children["interfaces"] = types.YChild{"Interfaces", &node.Interfaces}
    node.EntityData.Children["tag-allocations"] = types.YChild{"TagAllocations", &node.TagAllocations}
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-id"] = types.YLeaf{"NodeId", node.NodeId}
    return &(node.EntityData)
}

// Vlan_Nodes_Node_Trunks
// VLAN trunk table (specific to this node)
type Vlan_Nodes_Node_Trunks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational data for trunk interfaces configured with VLANs. The type is
    // slice of Vlan_Nodes_Node_Trunks_Trunk.
    Trunk []Vlan_Nodes_Node_Trunks_Trunk
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

    trunks.EntityData.Children = make(map[string]types.YChild)
    trunks.EntityData.Children["trunk"] = types.YChild{"Trunk", nil}
    for i := range trunks.Trunk {
        trunks.EntityData.Children[types.GetSegmentPath(&trunks.Trunk[i])] = types.YChild{"Trunk", &trunks.Trunk[i]}
    }
    trunks.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(trunks.EntityData)
}

// Vlan_Nodes_Node_Trunks_Trunk
// Operational data for trunk interfaces
// configured with VLANs
type Vlan_Nodes_Node_Trunks_Trunk struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The interface name. The type is string with
    // pattern: b'[a-zA-Z0-9./-]+'.
    Interface_ interface{}

    // Interface name. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceXr interface{}

    // Interface state. The type is ImStateEnum.
    State interface{}

    // L2 MTU. The type is interface{} with range: 0..65535.
    Mtu interface{}

    // QinQ Outer Tag Ether Type. The type is VlanQinqOuterEtype.
    QinqOuterEtherType interface{}

    // Number of subinterfaces with 802.1ad outer tag. The type is interface{}
    // with range: 0..4294967295.
    Dot1AdCount interface{}

    // Interface/Sub-interface handling untagged frames. The type is string with
    // pattern: b'[a-zA-Z0-9./-]+'.
    UntaggedInterface interface{}

    // IEEE 802.1Q/802.1ad multicast MAC address filtering. The type is
    // EthFiltering.
    MacFiltering interface{}

    // Layer 2 Transport Subinterfaces.
    Layer2SubInterfaces Vlan_Nodes_Node_Trunks_Trunk_Layer2SubInterfaces

    // Layer 3 Terminated Subinterfaces.
    Layer3SubInterfaces Vlan_Nodes_Node_Trunks_Trunk_Layer3SubInterfaces
}

func (trunk *Vlan_Nodes_Node_Trunks_Trunk) GetEntityData() *types.CommonEntityData {
    trunk.EntityData.YFilter = trunk.YFilter
    trunk.EntityData.YangName = "trunk"
    trunk.EntityData.BundleName = "cisco_ios_xr"
    trunk.EntityData.ParentYangName = "trunks"
    trunk.EntityData.SegmentPath = "trunk" + "[interface='" + fmt.Sprintf("%v", trunk.Interface_) + "']"
    trunk.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trunk.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trunk.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trunk.EntityData.Children = make(map[string]types.YChild)
    trunk.EntityData.Children["layer2-sub-interfaces"] = types.YChild{"Layer2SubInterfaces", &trunk.Layer2SubInterfaces}
    trunk.EntityData.Children["layer3-sub-interfaces"] = types.YChild{"Layer3SubInterfaces", &trunk.Layer3SubInterfaces}
    trunk.EntityData.Leafs = make(map[string]types.YLeaf)
    trunk.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", trunk.Interface_}
    trunk.EntityData.Leafs["interface-xr"] = types.YLeaf{"InterfaceXr", trunk.InterfaceXr}
    trunk.EntityData.Leafs["state"] = types.YLeaf{"State", trunk.State}
    trunk.EntityData.Leafs["mtu"] = types.YLeaf{"Mtu", trunk.Mtu}
    trunk.EntityData.Leafs["qinq-outer-ether-type"] = types.YLeaf{"QinqOuterEtherType", trunk.QinqOuterEtherType}
    trunk.EntityData.Leafs["dot1ad-count"] = types.YLeaf{"Dot1AdCount", trunk.Dot1AdCount}
    trunk.EntityData.Leafs["untagged-interface"] = types.YLeaf{"UntaggedInterface", trunk.UntaggedInterface}
    trunk.EntityData.Leafs["mac-filtering"] = types.YLeaf{"MacFiltering", trunk.MacFiltering}
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
    Dot1QCount interface{}

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

    layer2SubInterfaces.EntityData.Children = make(map[string]types.YChild)
    layer2SubInterfaces.EntityData.Children["state-counters"] = types.YChild{"StateCounters", &layer2SubInterfaces.StateCounters}
    layer2SubInterfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    layer2SubInterfaces.EntityData.Leafs["total-count"] = types.YLeaf{"TotalCount", layer2SubInterfaces.TotalCount}
    layer2SubInterfaces.EntityData.Leafs["dot1q-count"] = types.YLeaf{"Dot1QCount", layer2SubInterfaces.Dot1QCount}
    layer2SubInterfaces.EntityData.Leafs["qin-q-count"] = types.YLeaf{"QinQCount", layer2SubInterfaces.QinQCount}
    layer2SubInterfaces.EntityData.Leafs["qin-any-count"] = types.YLeaf{"QinAnyCount", layer2SubInterfaces.QinAnyCount}
    layer2SubInterfaces.EntityData.Leafs["untagged-count"] = types.YLeaf{"UntaggedCount", layer2SubInterfaces.UntaggedCount}
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

    stateCounters.EntityData.Children = make(map[string]types.YChild)
    stateCounters.EntityData.Leafs = make(map[string]types.YLeaf)
    stateCounters.EntityData.Leafs["up"] = types.YLeaf{"Up", stateCounters.Up}
    stateCounters.EntityData.Leafs["down"] = types.YLeaf{"Down", stateCounters.Down}
    stateCounters.EntityData.Leafs["admin-down"] = types.YLeaf{"AdminDown", stateCounters.AdminDown}
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
    Dot1QCount interface{}

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

    layer3SubInterfaces.EntityData.Children = make(map[string]types.YChild)
    layer3SubInterfaces.EntityData.Children["state-counters"] = types.YChild{"StateCounters", &layer3SubInterfaces.StateCounters}
    layer3SubInterfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    layer3SubInterfaces.EntityData.Leafs["total-count"] = types.YLeaf{"TotalCount", layer3SubInterfaces.TotalCount}
    layer3SubInterfaces.EntityData.Leafs["dot1q-count"] = types.YLeaf{"Dot1QCount", layer3SubInterfaces.Dot1QCount}
    layer3SubInterfaces.EntityData.Leafs["qin-q-count"] = types.YLeaf{"QinQCount", layer3SubInterfaces.QinQCount}
    layer3SubInterfaces.EntityData.Leafs["untagged-count"] = types.YLeaf{"UntaggedCount", layer3SubInterfaces.UntaggedCount}
    layer3SubInterfaces.EntityData.Leafs["native-vlan"] = types.YLeaf{"NativeVlan", layer3SubInterfaces.NativeVlan}
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

    stateCounters.EntityData.Children = make(map[string]types.YChild)
    stateCounters.EntityData.Leafs = make(map[string]types.YLeaf)
    stateCounters.EntityData.Leafs["up"] = types.YLeaf{"Up", stateCounters.Up}
    stateCounters.EntityData.Leafs["down"] = types.YLeaf{"Down", stateCounters.Down}
    stateCounters.EntityData.Leafs["admin-down"] = types.YLeaf{"AdminDown", stateCounters.AdminDown}
    return &(stateCounters.EntityData)
}

// Vlan_Nodes_Node_Interfaces
// VLAN interface table (specific to this node)
type Vlan_Nodes_Node_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational data for a sub-interface configured with VLANs. The type is
    // slice of Vlan_Nodes_Node_Interfaces_Interface_.
    Interface_ []Vlan_Nodes_Node_Interfaces_Interface
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

    interfaces.EntityData.Children = make(map[string]types.YChild)
    interfaces.EntityData.Children["interface"] = types.YChild{"Interface_", nil}
    for i := range interfaces.Interface_ {
        interfaces.EntityData.Children[types.GetSegmentPath(&interfaces.Interface_[i])] = types.YChild{"Interface_", &interfaces.Interface_[i]}
    }
    interfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaces.EntityData)
}

// Vlan_Nodes_Node_Interfaces_Interface
// Operational data for a sub-interface
// configured with VLANs
type Vlan_Nodes_Node_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The interface name. The type is string with
    // pattern: b'[a-zA-Z0-9./-]+'.
    Interface_ interface{}

    // Interface. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceXr interface{}

    // Parent interface. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
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
    self.EntityData.SegmentPath = "interface" + "[interface='" + fmt.Sprintf("%v", self.Interface_) + "']"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Children["encapsulation-details"] = types.YChild{"EncapsulationDetails", &self.EncapsulationDetails}
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", self.Interface_}
    self.EntityData.Leafs["interface-xr"] = types.YLeaf{"InterfaceXr", self.InterfaceXr}
    self.EntityData.Leafs["parent-interface"] = types.YLeaf{"ParentInterface", self.ParentInterface}
    self.EntityData.Leafs["service"] = types.YLeaf{"Service", self.Service}
    self.EntityData.Leafs["state"] = types.YLeaf{"State", self.State}
    self.EntityData.Leafs["mtu"] = types.YLeaf{"Mtu", self.Mtu}
    self.EntityData.Leafs["switched-mtu"] = types.YLeaf{"SwitchedMtu", self.SwitchedMtu}
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
    Dot1AdTag interface{}

    // 802.1ad native tag value. The type is interface{} with range: 0..65535.
    Dot1AdNativeTag interface{}

    // 802.1ad Outer tag value. The type is interface{} with range: 0..65535.
    Dot1AdOuterTag interface{}

    // Stack value.
    Stack Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_Stack

    // Service Instance encapsulation.
    ServiceInstanceDetails Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails

    // 802.1ad 802.1Q stack value.
    Dot1AdDot1QStack Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_Dot1AdDot1QStack
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

    encapsulationDetails.EntityData.Children = make(map[string]types.YChild)
    encapsulationDetails.EntityData.Children["stack"] = types.YChild{"Stack", &encapsulationDetails.Stack}
    encapsulationDetails.EntityData.Children["service-instance-details"] = types.YChild{"ServiceInstanceDetails", &encapsulationDetails.ServiceInstanceDetails}
    encapsulationDetails.EntityData.Children["dot1ad-dot1q-stack"] = types.YChild{"Dot1AdDot1QStack", &encapsulationDetails.Dot1AdDot1QStack}
    encapsulationDetails.EntityData.Leafs = make(map[string]types.YLeaf)
    encapsulationDetails.EntityData.Leafs["vlan-encapsulation"] = types.YLeaf{"VlanEncapsulation", encapsulationDetails.VlanEncapsulation}
    encapsulationDetails.EntityData.Leafs["tag"] = types.YLeaf{"Tag", encapsulationDetails.Tag}
    encapsulationDetails.EntityData.Leafs["outer-tag"] = types.YLeaf{"OuterTag", encapsulationDetails.OuterTag}
    encapsulationDetails.EntityData.Leafs["native-tag"] = types.YLeaf{"NativeTag", encapsulationDetails.NativeTag}
    encapsulationDetails.EntityData.Leafs["dot1ad-tag"] = types.YLeaf{"Dot1AdTag", encapsulationDetails.Dot1AdTag}
    encapsulationDetails.EntityData.Leafs["dot1ad-native-tag"] = types.YLeaf{"Dot1AdNativeTag", encapsulationDetails.Dot1AdNativeTag}
    encapsulationDetails.EntityData.Leafs["dot1ad-outer-tag"] = types.YLeaf{"Dot1AdOuterTag", encapsulationDetails.Dot1AdOuterTag}
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

    stack.EntityData.Children = make(map[string]types.YChild)
    stack.EntityData.Leafs = make(map[string]types.YLeaf)
    stack.EntityData.Leafs["outer-tag"] = types.YLeaf{"OuterTag", stack.OuterTag}
    stack.EntityData.Leafs["second-tag"] = types.YLeaf{"SecondTag", stack.SecondTag}
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
    // pattern: b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    SourceMacMatch interface{}

    // The destination MAC address to match on ingress. The type is string with
    // pattern: b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    DestinationMacMatch interface{}

    // VLAN tags for locally-sourced traffic.
    LocalTrafficStack Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack

    // Tags to match on ingress packets. The type is slice of
    // Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch.
    TagsToMatch []Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch

    // VLAN tags pushed on egress. The type is slice of
    // Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_Pushe.
    Pushe []Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_Pushe
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

    serviceInstanceDetails.EntityData.Children = make(map[string]types.YChild)
    serviceInstanceDetails.EntityData.Children["local-traffic-stack"] = types.YChild{"LocalTrafficStack", &serviceInstanceDetails.LocalTrafficStack}
    serviceInstanceDetails.EntityData.Children["tags-to-match"] = types.YChild{"TagsToMatch", nil}
    for i := range serviceInstanceDetails.TagsToMatch {
        serviceInstanceDetails.EntityData.Children[types.GetSegmentPath(&serviceInstanceDetails.TagsToMatch[i])] = types.YChild{"TagsToMatch", &serviceInstanceDetails.TagsToMatch[i]}
    }
    serviceInstanceDetails.EntityData.Children["pushe"] = types.YChild{"Pushe", nil}
    for i := range serviceInstanceDetails.Pushe {
        serviceInstanceDetails.EntityData.Children[types.GetSegmentPath(&serviceInstanceDetails.Pushe[i])] = types.YChild{"Pushe", &serviceInstanceDetails.Pushe[i]}
    }
    serviceInstanceDetails.EntityData.Leafs = make(map[string]types.YLeaf)
    serviceInstanceDetails.EntityData.Leafs["payload-ethertype"] = types.YLeaf{"PayloadEthertype", serviceInstanceDetails.PayloadEthertype}
    serviceInstanceDetails.EntityData.Leafs["tags-popped"] = types.YLeaf{"TagsPopped", serviceInstanceDetails.TagsPopped}
    serviceInstanceDetails.EntityData.Leafs["is-exact-match"] = types.YLeaf{"IsExactMatch", serviceInstanceDetails.IsExactMatch}
    serviceInstanceDetails.EntityData.Leafs["is-native-vlan"] = types.YLeaf{"IsNativeVlan", serviceInstanceDetails.IsNativeVlan}
    serviceInstanceDetails.EntityData.Leafs["is-native-preserving"] = types.YLeaf{"IsNativePreserving", serviceInstanceDetails.IsNativePreserving}
    serviceInstanceDetails.EntityData.Leafs["source-mac-match"] = types.YLeaf{"SourceMacMatch", serviceInstanceDetails.SourceMacMatch}
    serviceInstanceDetails.EntityData.Leafs["destination-mac-match"] = types.YLeaf{"DestinationMacMatch", serviceInstanceDetails.DestinationMacMatch}
    return &(serviceInstanceDetails.EntityData)
}

// Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack
// VLAN tags for locally-sourced traffic
type Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VLAN tags for locally-sourced traffic. The type is slice of
    // Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag.
    LocalTrafficTag []Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag
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

    localTrafficStack.EntityData.Children = make(map[string]types.YChild)
    localTrafficStack.EntityData.Children["local-traffic-tag"] = types.YChild{"LocalTrafficTag", nil}
    for i := range localTrafficStack.LocalTrafficTag {
        localTrafficStack.EntityData.Children[types.GetSegmentPath(&localTrafficStack.LocalTrafficTag[i])] = types.YChild{"LocalTrafficTag", &localTrafficStack.LocalTrafficTag[i]}
    }
    localTrafficStack.EntityData.Leafs = make(map[string]types.YLeaf)
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

    localTrafficTag.EntityData.Children = make(map[string]types.YChild)
    localTrafficTag.EntityData.Leafs = make(map[string]types.YLeaf)
    localTrafficTag.EntityData.Leafs["ethertype"] = types.YLeaf{"Ethertype", localTrafficTag.Ethertype}
    localTrafficTag.EntityData.Leafs["vlan-id"] = types.YLeaf{"VlanId", localTrafficTag.VlanId}
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
    VlanRange []Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange
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

    tagsToMatch.EntityData.Children = make(map[string]types.YChild)
    tagsToMatch.EntityData.Children["vlan-range"] = types.YChild{"VlanRange", nil}
    for i := range tagsToMatch.VlanRange {
        tagsToMatch.EntityData.Children[types.GetSegmentPath(&tagsToMatch.VlanRange[i])] = types.YChild{"VlanRange", &tagsToMatch.VlanRange[i]}
    }
    tagsToMatch.EntityData.Leafs = make(map[string]types.YLeaf)
    tagsToMatch.EntityData.Leafs["ethertype"] = types.YLeaf{"Ethertype", tagsToMatch.Ethertype}
    tagsToMatch.EntityData.Leafs["priority"] = types.YLeaf{"Priority", tagsToMatch.Priority}
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

    vlanRange.EntityData.Children = make(map[string]types.YChild)
    vlanRange.EntityData.Leafs = make(map[string]types.YLeaf)
    vlanRange.EntityData.Leafs["vlan-id-low"] = types.YLeaf{"VlanIdLow", vlanRange.VlanIdLow}
    vlanRange.EntityData.Leafs["vlan-id-high"] = types.YLeaf{"VlanIdHigh", vlanRange.VlanIdHigh}
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

    pushe.EntityData.Children = make(map[string]types.YChild)
    pushe.EntityData.Leafs = make(map[string]types.YLeaf)
    pushe.EntityData.Leafs["ethertype"] = types.YLeaf{"Ethertype", pushe.Ethertype}
    pushe.EntityData.Leafs["vlan-id"] = types.YLeaf{"VlanId", pushe.VlanId}
    return &(pushe.EntityData)
}

// Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_Dot1AdDot1QStack
// 802.1ad 802.1Q stack value
type Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_Dot1AdDot1QStack struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Outer tag value. The type is interface{} with range: 0..65535.
    OuterTag interface{}

    // Second tag value. The type is interface{} with range: 0..65535.
    SecondTag interface{}
}

func (dot1AdDot1QStack *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_Dot1AdDot1QStack) GetEntityData() *types.CommonEntityData {
    dot1AdDot1QStack.EntityData.YFilter = dot1AdDot1QStack.YFilter
    dot1AdDot1QStack.EntityData.YangName = "dot1ad-dot1q-stack"
    dot1AdDot1QStack.EntityData.BundleName = "cisco_ios_xr"
    dot1AdDot1QStack.EntityData.ParentYangName = "encapsulation-details"
    dot1AdDot1QStack.EntityData.SegmentPath = "dot1ad-dot1q-stack"
    dot1AdDot1QStack.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dot1AdDot1QStack.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dot1AdDot1QStack.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dot1AdDot1QStack.EntityData.Children = make(map[string]types.YChild)
    dot1AdDot1QStack.EntityData.Leafs = make(map[string]types.YLeaf)
    dot1AdDot1QStack.EntityData.Leafs["outer-tag"] = types.YLeaf{"OuterTag", dot1AdDot1QStack.OuterTag}
    dot1AdDot1QStack.EntityData.Leafs["second-tag"] = types.YLeaf{"SecondTag", dot1AdDot1QStack.SecondTag}
    return &(dot1AdDot1QStack.EntityData)
}

// Vlan_Nodes_Node_TagAllocations
// VLAN tag allocation table (specific to this
// node)
type Vlan_Nodes_Node_TagAllocations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational data for a sub-interface configured with VLANs. The type is
    // slice of Vlan_Nodes_Node_TagAllocations_TagAllocation.
    TagAllocation []Vlan_Nodes_Node_TagAllocations_TagAllocation
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

    tagAllocations.EntityData.Children = make(map[string]types.YChild)
    tagAllocations.EntityData.Children["tag-allocation"] = types.YChild{"TagAllocation", nil}
    for i := range tagAllocations.TagAllocation {
        tagAllocations.EntityData.Children[types.GetSegmentPath(&tagAllocations.TagAllocation[i])] = types.YChild{"TagAllocation", &tagAllocations.TagAllocation[i]}
    }
    tagAllocations.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(tagAllocations.EntityData)
}

// Vlan_Nodes_Node_TagAllocations_TagAllocation
// Operational data for a sub-interface
// configured with VLANs
type Vlan_Nodes_Node_TagAllocations_TagAllocation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The interface name. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    Interface_ interface{}

    // The first (outermost) tag. The type is interface{} with range: 1..4094.
    FirstTag interface{}

    // The second tag. The type is one of the following types: enumeration
    // VlanTagOrAny, or int with range: 1..4096.
    SecondTag interface{}

    // Interface. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceXr interface{}

    // Parent interface. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
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

    tagAllocation.EntityData.Children = make(map[string]types.YChild)
    tagAllocation.EntityData.Children["encapsulation-details"] = types.YChild{"EncapsulationDetails", &tagAllocation.EncapsulationDetails}
    tagAllocation.EntityData.Leafs = make(map[string]types.YLeaf)
    tagAllocation.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", tagAllocation.Interface_}
    tagAllocation.EntityData.Leafs["first-tag"] = types.YLeaf{"FirstTag", tagAllocation.FirstTag}
    tagAllocation.EntityData.Leafs["second-tag"] = types.YLeaf{"SecondTag", tagAllocation.SecondTag}
    tagAllocation.EntityData.Leafs["interface-xr"] = types.YLeaf{"InterfaceXr", tagAllocation.InterfaceXr}
    tagAllocation.EntityData.Leafs["parent-interface"] = types.YLeaf{"ParentInterface", tagAllocation.ParentInterface}
    tagAllocation.EntityData.Leafs["service"] = types.YLeaf{"Service", tagAllocation.Service}
    tagAllocation.EntityData.Leafs["state"] = types.YLeaf{"State", tagAllocation.State}
    tagAllocation.EntityData.Leafs["mtu"] = types.YLeaf{"Mtu", tagAllocation.Mtu}
    tagAllocation.EntityData.Leafs["switched-mtu"] = types.YLeaf{"SwitchedMtu", tagAllocation.SwitchedMtu}
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
    Dot1AdTag interface{}

    // 802.1ad native tag value. The type is interface{} with range: 0..65535.
    Dot1AdNativeTag interface{}

    // 802.1ad Outer tag value. The type is interface{} with range: 0..65535.
    Dot1AdOuterTag interface{}

    // Stack value.
    Stack Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_Stack

    // Service Instance encapsulation.
    ServiceInstanceDetails Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails

    // 802.1ad 802.1Q stack value.
    Dot1AdDot1QStack Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_Dot1AdDot1QStack
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

    encapsulationDetails.EntityData.Children = make(map[string]types.YChild)
    encapsulationDetails.EntityData.Children["stack"] = types.YChild{"Stack", &encapsulationDetails.Stack}
    encapsulationDetails.EntityData.Children["service-instance-details"] = types.YChild{"ServiceInstanceDetails", &encapsulationDetails.ServiceInstanceDetails}
    encapsulationDetails.EntityData.Children["dot1ad-dot1q-stack"] = types.YChild{"Dot1AdDot1QStack", &encapsulationDetails.Dot1AdDot1QStack}
    encapsulationDetails.EntityData.Leafs = make(map[string]types.YLeaf)
    encapsulationDetails.EntityData.Leafs["vlan-encapsulation"] = types.YLeaf{"VlanEncapsulation", encapsulationDetails.VlanEncapsulation}
    encapsulationDetails.EntityData.Leafs["tag"] = types.YLeaf{"Tag", encapsulationDetails.Tag}
    encapsulationDetails.EntityData.Leafs["outer-tag"] = types.YLeaf{"OuterTag", encapsulationDetails.OuterTag}
    encapsulationDetails.EntityData.Leafs["native-tag"] = types.YLeaf{"NativeTag", encapsulationDetails.NativeTag}
    encapsulationDetails.EntityData.Leafs["dot1ad-tag"] = types.YLeaf{"Dot1AdTag", encapsulationDetails.Dot1AdTag}
    encapsulationDetails.EntityData.Leafs["dot1ad-native-tag"] = types.YLeaf{"Dot1AdNativeTag", encapsulationDetails.Dot1AdNativeTag}
    encapsulationDetails.EntityData.Leafs["dot1ad-outer-tag"] = types.YLeaf{"Dot1AdOuterTag", encapsulationDetails.Dot1AdOuterTag}
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

    stack.EntityData.Children = make(map[string]types.YChild)
    stack.EntityData.Leafs = make(map[string]types.YLeaf)
    stack.EntityData.Leafs["outer-tag"] = types.YLeaf{"OuterTag", stack.OuterTag}
    stack.EntityData.Leafs["second-tag"] = types.YLeaf{"SecondTag", stack.SecondTag}
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
    // pattern: b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    SourceMacMatch interface{}

    // The destination MAC address to match on ingress. The type is string with
    // pattern: b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    DestinationMacMatch interface{}

    // VLAN tags for locally-sourced traffic.
    LocalTrafficStack Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack

    // Tags to match on ingress packets. The type is slice of
    // Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch.
    TagsToMatch []Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch

    // VLAN tags pushed on egress. The type is slice of
    // Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_Pushe.
    Pushe []Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_Pushe
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

    serviceInstanceDetails.EntityData.Children = make(map[string]types.YChild)
    serviceInstanceDetails.EntityData.Children["local-traffic-stack"] = types.YChild{"LocalTrafficStack", &serviceInstanceDetails.LocalTrafficStack}
    serviceInstanceDetails.EntityData.Children["tags-to-match"] = types.YChild{"TagsToMatch", nil}
    for i := range serviceInstanceDetails.TagsToMatch {
        serviceInstanceDetails.EntityData.Children[types.GetSegmentPath(&serviceInstanceDetails.TagsToMatch[i])] = types.YChild{"TagsToMatch", &serviceInstanceDetails.TagsToMatch[i]}
    }
    serviceInstanceDetails.EntityData.Children["pushe"] = types.YChild{"Pushe", nil}
    for i := range serviceInstanceDetails.Pushe {
        serviceInstanceDetails.EntityData.Children[types.GetSegmentPath(&serviceInstanceDetails.Pushe[i])] = types.YChild{"Pushe", &serviceInstanceDetails.Pushe[i]}
    }
    serviceInstanceDetails.EntityData.Leafs = make(map[string]types.YLeaf)
    serviceInstanceDetails.EntityData.Leafs["payload-ethertype"] = types.YLeaf{"PayloadEthertype", serviceInstanceDetails.PayloadEthertype}
    serviceInstanceDetails.EntityData.Leafs["tags-popped"] = types.YLeaf{"TagsPopped", serviceInstanceDetails.TagsPopped}
    serviceInstanceDetails.EntityData.Leafs["is-exact-match"] = types.YLeaf{"IsExactMatch", serviceInstanceDetails.IsExactMatch}
    serviceInstanceDetails.EntityData.Leafs["is-native-vlan"] = types.YLeaf{"IsNativeVlan", serviceInstanceDetails.IsNativeVlan}
    serviceInstanceDetails.EntityData.Leafs["is-native-preserving"] = types.YLeaf{"IsNativePreserving", serviceInstanceDetails.IsNativePreserving}
    serviceInstanceDetails.EntityData.Leafs["source-mac-match"] = types.YLeaf{"SourceMacMatch", serviceInstanceDetails.SourceMacMatch}
    serviceInstanceDetails.EntityData.Leafs["destination-mac-match"] = types.YLeaf{"DestinationMacMatch", serviceInstanceDetails.DestinationMacMatch}
    return &(serviceInstanceDetails.EntityData)
}

// Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack
// VLAN tags for locally-sourced traffic
type Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VLAN tags for locally-sourced traffic. The type is slice of
    // Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag.
    LocalTrafficTag []Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag
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

    localTrafficStack.EntityData.Children = make(map[string]types.YChild)
    localTrafficStack.EntityData.Children["local-traffic-tag"] = types.YChild{"LocalTrafficTag", nil}
    for i := range localTrafficStack.LocalTrafficTag {
        localTrafficStack.EntityData.Children[types.GetSegmentPath(&localTrafficStack.LocalTrafficTag[i])] = types.YChild{"LocalTrafficTag", &localTrafficStack.LocalTrafficTag[i]}
    }
    localTrafficStack.EntityData.Leafs = make(map[string]types.YLeaf)
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

    localTrafficTag.EntityData.Children = make(map[string]types.YChild)
    localTrafficTag.EntityData.Leafs = make(map[string]types.YLeaf)
    localTrafficTag.EntityData.Leafs["ethertype"] = types.YLeaf{"Ethertype", localTrafficTag.Ethertype}
    localTrafficTag.EntityData.Leafs["vlan-id"] = types.YLeaf{"VlanId", localTrafficTag.VlanId}
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
    VlanRange []Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange
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

    tagsToMatch.EntityData.Children = make(map[string]types.YChild)
    tagsToMatch.EntityData.Children["vlan-range"] = types.YChild{"VlanRange", nil}
    for i := range tagsToMatch.VlanRange {
        tagsToMatch.EntityData.Children[types.GetSegmentPath(&tagsToMatch.VlanRange[i])] = types.YChild{"VlanRange", &tagsToMatch.VlanRange[i]}
    }
    tagsToMatch.EntityData.Leafs = make(map[string]types.YLeaf)
    tagsToMatch.EntityData.Leafs["ethertype"] = types.YLeaf{"Ethertype", tagsToMatch.Ethertype}
    tagsToMatch.EntityData.Leafs["priority"] = types.YLeaf{"Priority", tagsToMatch.Priority}
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

    vlanRange.EntityData.Children = make(map[string]types.YChild)
    vlanRange.EntityData.Leafs = make(map[string]types.YLeaf)
    vlanRange.EntityData.Leafs["vlan-id-low"] = types.YLeaf{"VlanIdLow", vlanRange.VlanIdLow}
    vlanRange.EntityData.Leafs["vlan-id-high"] = types.YLeaf{"VlanIdHigh", vlanRange.VlanIdHigh}
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

    pushe.EntityData.Children = make(map[string]types.YChild)
    pushe.EntityData.Leafs = make(map[string]types.YLeaf)
    pushe.EntityData.Leafs["ethertype"] = types.YLeaf{"Ethertype", pushe.Ethertype}
    pushe.EntityData.Leafs["vlan-id"] = types.YLeaf{"VlanId", pushe.VlanId}
    return &(pushe.EntityData)
}

// Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_Dot1AdDot1QStack
// 802.1ad 802.1Q stack value
type Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_Dot1AdDot1QStack struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Outer tag value. The type is interface{} with range: 0..65535.
    OuterTag interface{}

    // Second tag value. The type is interface{} with range: 0..65535.
    SecondTag interface{}
}

func (dot1AdDot1QStack *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_Dot1AdDot1QStack) GetEntityData() *types.CommonEntityData {
    dot1AdDot1QStack.EntityData.YFilter = dot1AdDot1QStack.YFilter
    dot1AdDot1QStack.EntityData.YangName = "dot1ad-dot1q-stack"
    dot1AdDot1QStack.EntityData.BundleName = "cisco_ios_xr"
    dot1AdDot1QStack.EntityData.ParentYangName = "encapsulation-details"
    dot1AdDot1QStack.EntityData.SegmentPath = "dot1ad-dot1q-stack"
    dot1AdDot1QStack.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dot1AdDot1QStack.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dot1AdDot1QStack.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dot1AdDot1QStack.EntityData.Children = make(map[string]types.YChild)
    dot1AdDot1QStack.EntityData.Leafs = make(map[string]types.YLeaf)
    dot1AdDot1QStack.EntityData.Leafs["outer-tag"] = types.YLeaf{"OuterTag", dot1AdDot1QStack.OuterTag}
    dot1AdDot1QStack.EntityData.Leafs["second-tag"] = types.YLeaf{"SecondTag", dot1AdDot1QStack.SecondTag}
    return &(dot1AdDot1QStack.EntityData)
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

    ethernetEncapsulation.EntityData.Children = make(map[string]types.YChild)
    ethernetEncapsulation.EntityData.Children["nodes"] = types.YChild{"Nodes", &ethernetEncapsulation.Nodes}
    ethernetEncapsulation.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ethernetEncapsulation.EntityData)
}

// EthernetEncapsulation_Nodes
// Per node Ethernet encapsulation operational data
type EthernetEncapsulation_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The Ethernet encaps operational data for a particular node. The type is
    // slice of EthernetEncapsulation_Nodes_Node.
    Node []EthernetEncapsulation_Nodes_Node
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

    nodes.EntityData.Children = make(map[string]types.YChild)
    nodes.EntityData.Children["node"] = types.YChild{"Node", nil}
    for i := range nodes.Node {
        nodes.EntityData.Children[types.GetSegmentPath(&nodes.Node[i])] = types.YChild{"Node", &nodes.Node[i]}
    }
    nodes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nodes.EntityData)
}

// EthernetEncapsulation_Nodes_Node
// The Ethernet encaps operational data for a
// particular node
type EthernetEncapsulation_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The identifier for the node. The type is string
    // with pattern: b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeName interface{}

    // Unicast MAC filter table (specific to this node).
    UnicastMacFilters EthernetEncapsulation_Nodes_Node_UnicastMacFilters
}

func (node *EthernetEncapsulation_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Children["unicast-mac-filters"] = types.YChild{"UnicastMacFilters", &node.UnicastMacFilters}
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", node.NodeName}
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
    UnicastMacFilter []EthernetEncapsulation_Nodes_Node_UnicastMacFilters_UnicastMacFilter
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

    unicastMacFilters.EntityData.Children = make(map[string]types.YChild)
    unicastMacFilters.EntityData.Children["unicast-mac-filter"] = types.YChild{"UnicastMacFilter", nil}
    for i := range unicastMacFilters.UnicastMacFilter {
        unicastMacFilters.EntityData.Children[types.GetSegmentPath(&unicastMacFilters.UnicastMacFilter[i])] = types.YChild{"UnicastMacFilter", &unicastMacFilters.UnicastMacFilter[i]}
    }
    unicastMacFilters.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(unicastMacFilters.EntityData)
}

// EthernetEncapsulation_Nodes_Node_UnicastMacFilters_UnicastMacFilter
// Operational data for interface with MAC
// filters configured
type EthernetEncapsulation_Nodes_Node_UnicastMacFilters_UnicastMacFilter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The interface name. The type is string with
    // pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // Unicast MAC filter information. The type is slice of
    // EthernetEncapsulation_Nodes_Node_UnicastMacFilters_UnicastMacFilter_UnicastFilter.
    UnicastFilter []EthernetEncapsulation_Nodes_Node_UnicastMacFilters_UnicastMacFilter_UnicastFilter
}

func (unicastMacFilter *EthernetEncapsulation_Nodes_Node_UnicastMacFilters_UnicastMacFilter) GetEntityData() *types.CommonEntityData {
    unicastMacFilter.EntityData.YFilter = unicastMacFilter.YFilter
    unicastMacFilter.EntityData.YangName = "unicast-mac-filter"
    unicastMacFilter.EntityData.BundleName = "cisco_ios_xr"
    unicastMacFilter.EntityData.ParentYangName = "unicast-mac-filters"
    unicastMacFilter.EntityData.SegmentPath = "unicast-mac-filter" + "[interface-name='" + fmt.Sprintf("%v", unicastMacFilter.InterfaceName) + "']"
    unicastMacFilter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unicastMacFilter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unicastMacFilter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unicastMacFilter.EntityData.Children = make(map[string]types.YChild)
    unicastMacFilter.EntityData.Children["unicast-filter"] = types.YChild{"UnicastFilter", nil}
    for i := range unicastMacFilter.UnicastFilter {
        unicastMacFilter.EntityData.Children[types.GetSegmentPath(&unicastMacFilter.UnicastFilter[i])] = types.YChild{"UnicastFilter", &unicastMacFilter.UnicastFilter[i]}
    }
    unicastMacFilter.EntityData.Leafs = make(map[string]types.YLeaf)
    unicastMacFilter.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", unicastMacFilter.InterfaceName}
    return &(unicastMacFilter.EntityData)
}

// EthernetEncapsulation_Nodes_Node_UnicastMacFilters_UnicastMacFilter_UnicastFilter
// Unicast MAC filter information
type EthernetEncapsulation_Nodes_Node_UnicastMacFilters_UnicastMacFilter_UnicastFilter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MAC address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
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

    unicastFilter.EntityData.Children = make(map[string]types.YChild)
    unicastFilter.EntityData.Leafs = make(map[string]types.YLeaf)
    unicastFilter.EntityData.Leafs["mac-address"] = types.YLeaf{"MacAddress", unicastFilter.MacAddress}
    unicastFilter.EntityData.Leafs["mode"] = types.YLeaf{"Mode", unicastFilter.Mode}
    return &(unicastFilter.EntityData)
}

