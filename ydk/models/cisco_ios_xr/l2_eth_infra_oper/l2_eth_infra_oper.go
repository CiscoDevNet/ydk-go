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

// EthCapsUcastMacMode represents Eth caps ucast mac mode
type EthCapsUcastMacMode string

const (
    // Reserved
    EthCapsUcastMacMode_reserved EthCapsUcastMacMode = "reserved"

    // Permit
    EthCapsUcastMacMode_permit EthCapsUcastMacMode = "permit"
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
    parent types.Entity
    YFilter yfilter.YFilter

    // MAC accounting interface table in MIB lexicographic order.
    Interfaces MacAccounting_Interfaces
}

func (macAccounting *MacAccounting) GetFilter() yfilter.YFilter { return macAccounting.YFilter }

func (macAccounting *MacAccounting) SetFilter(yf yfilter.YFilter) { macAccounting.YFilter = yf }

func (macAccounting *MacAccounting) GetGoName(yname string) string {
    if yname == "interfaces" { return "Interfaces" }
    return ""
}

func (macAccounting *MacAccounting) GetSegmentPath() string {
    return "Cisco-IOS-XR-l2-eth-infra-oper:mac-accounting"
}

func (macAccounting *MacAccounting) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interfaces" {
        return &macAccounting.Interfaces
    }
    return nil
}

func (macAccounting *MacAccounting) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interfaces"] = &macAccounting.Interfaces
    return children
}

func (macAccounting *MacAccounting) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (macAccounting *MacAccounting) GetBundleName() string { return "cisco_ios_xr" }

func (macAccounting *MacAccounting) GetYangName() string { return "mac-accounting" }

func (macAccounting *MacAccounting) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (macAccounting *MacAccounting) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (macAccounting *MacAccounting) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (macAccounting *MacAccounting) SetParent(parent types.Entity) { macAccounting.parent = parent }

func (macAccounting *MacAccounting) GetParent() types.Entity { return macAccounting.parent }

func (macAccounting *MacAccounting) GetParentYangName() string { return "Cisco-IOS-XR-l2-eth-infra-oper" }

// MacAccounting_Interfaces
// MAC accounting interface table in MIB
// lexicographic order
type MacAccounting_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operational data and statistics for an interface configured with MAC
    // accounting enabled. The type is slice of
    // MacAccounting_Interfaces_Interface.
    Interface []MacAccounting_Interfaces_Interface
}

func (interfaces *MacAccounting_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *MacAccounting_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *MacAccounting_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *MacAccounting_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *MacAccounting_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MacAccounting_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *MacAccounting_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *MacAccounting_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *MacAccounting_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *MacAccounting_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *MacAccounting_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *MacAccounting_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *MacAccounting_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *MacAccounting_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *MacAccounting_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *MacAccounting_Interfaces) GetParentYangName() string { return "mac-accounting" }

// MacAccounting_Interfaces_Interface
// Operational data and statistics for an
// interface configured with MAC accounting
// enabled
type MacAccounting_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The interface name. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
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

func (self *MacAccounting_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *MacAccounting_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *MacAccounting_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "state" { return "State" }
    if yname == "ingress-statistic" { return "IngressStatistic" }
    if yname == "egress-statistic" { return "EgressStatistic" }
    return ""
}

func (self *MacAccounting_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *MacAccounting_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "state" {
        return &self.State
    }
    if childYangName == "ingress-statistic" {
        for _, c := range self.IngressStatistic {
            if self.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MacAccounting_Interfaces_Interface_IngressStatistic{}
        self.IngressStatistic = append(self.IngressStatistic, child)
        return &self.IngressStatistic[len(self.IngressStatistic)-1]
    }
    if childYangName == "egress-statistic" {
        for _, c := range self.EgressStatistic {
            if self.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MacAccounting_Interfaces_Interface_EgressStatistic{}
        self.EgressStatistic = append(self.EgressStatistic, child)
        return &self.EgressStatistic[len(self.EgressStatistic)-1]
    }
    return nil
}

func (self *MacAccounting_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["state"] = &self.State
    for i := range self.IngressStatistic {
        children[self.IngressStatistic[i].GetSegmentPath()] = &self.IngressStatistic[i]
    }
    for i := range self.EgressStatistic {
        children[self.EgressStatistic[i].GetSegmentPath()] = &self.EgressStatistic[i]
    }
    return children
}

func (self *MacAccounting_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    return leafs
}

func (self *MacAccounting_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *MacAccounting_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *MacAccounting_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *MacAccounting_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *MacAccounting_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *MacAccounting_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *MacAccounting_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *MacAccounting_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// MacAccounting_Interfaces_Interface_State
// MAC accounting state for the interface
type MacAccounting_Interfaces_Interface_State struct {
    parent types.Entity
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

func (state *MacAccounting_Interfaces_Interface_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *MacAccounting_Interfaces_Interface_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *MacAccounting_Interfaces_Interface_State) GetGoName(yname string) string {
    if yname == "is-ingress-enabled" { return "IsIngressEnabled" }
    if yname == "is-egress-enabled" { return "IsEgressEnabled" }
    if yname == "number-available-ingress" { return "NumberAvailableIngress" }
    if yname == "number-available-egress" { return "NumberAvailableEgress" }
    if yname == "number-available-on-node" { return "NumberAvailableOnNode" }
    return ""
}

func (state *MacAccounting_Interfaces_Interface_State) GetSegmentPath() string {
    return "state"
}

func (state *MacAccounting_Interfaces_Interface_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *MacAccounting_Interfaces_Interface_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *MacAccounting_Interfaces_Interface_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-ingress-enabled"] = state.IsIngressEnabled
    leafs["is-egress-enabled"] = state.IsEgressEnabled
    leafs["number-available-ingress"] = state.NumberAvailableIngress
    leafs["number-available-egress"] = state.NumberAvailableEgress
    leafs["number-available-on-node"] = state.NumberAvailableOnNode
    return leafs
}

func (state *MacAccounting_Interfaces_Interface_State) GetBundleName() string { return "cisco_ios_xr" }

func (state *MacAccounting_Interfaces_Interface_State) GetYangName() string { return "state" }

func (state *MacAccounting_Interfaces_Interface_State) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (state *MacAccounting_Interfaces_Interface_State) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (state *MacAccounting_Interfaces_Interface_State) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (state *MacAccounting_Interfaces_Interface_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *MacAccounting_Interfaces_Interface_State) GetParent() types.Entity { return state.parent }

func (state *MacAccounting_Interfaces_Interface_State) GetParentYangName() string { return "interface" }

// MacAccounting_Interfaces_Interface_IngressStatistic
// Ingress MAC accounting statistics
type MacAccounting_Interfaces_Interface_IngressStatistic struct {
    parent types.Entity
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

func (ingressStatistic *MacAccounting_Interfaces_Interface_IngressStatistic) GetFilter() yfilter.YFilter { return ingressStatistic.YFilter }

func (ingressStatistic *MacAccounting_Interfaces_Interface_IngressStatistic) SetFilter(yf yfilter.YFilter) { ingressStatistic.YFilter = yf }

func (ingressStatistic *MacAccounting_Interfaces_Interface_IngressStatistic) GetGoName(yname string) string {
    if yname == "mac-address" { return "MacAddress" }
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (ingressStatistic *MacAccounting_Interfaces_Interface_IngressStatistic) GetSegmentPath() string {
    return "ingress-statistic"
}

func (ingressStatistic *MacAccounting_Interfaces_Interface_IngressStatistic) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ingressStatistic *MacAccounting_Interfaces_Interface_IngressStatistic) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ingressStatistic *MacAccounting_Interfaces_Interface_IngressStatistic) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mac-address"] = ingressStatistic.MacAddress
    leafs["packets"] = ingressStatistic.Packets
    leafs["bytes"] = ingressStatistic.Bytes
    return leafs
}

func (ingressStatistic *MacAccounting_Interfaces_Interface_IngressStatistic) GetBundleName() string { return "cisco_ios_xr" }

func (ingressStatistic *MacAccounting_Interfaces_Interface_IngressStatistic) GetYangName() string { return "ingress-statistic" }

func (ingressStatistic *MacAccounting_Interfaces_Interface_IngressStatistic) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ingressStatistic *MacAccounting_Interfaces_Interface_IngressStatistic) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ingressStatistic *MacAccounting_Interfaces_Interface_IngressStatistic) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ingressStatistic *MacAccounting_Interfaces_Interface_IngressStatistic) SetParent(parent types.Entity) { ingressStatistic.parent = parent }

func (ingressStatistic *MacAccounting_Interfaces_Interface_IngressStatistic) GetParent() types.Entity { return ingressStatistic.parent }

func (ingressStatistic *MacAccounting_Interfaces_Interface_IngressStatistic) GetParentYangName() string { return "interface" }

// MacAccounting_Interfaces_Interface_EgressStatistic
// Egress MAC accounting statistics
type MacAccounting_Interfaces_Interface_EgressStatistic struct {
    parent types.Entity
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

func (egressStatistic *MacAccounting_Interfaces_Interface_EgressStatistic) GetFilter() yfilter.YFilter { return egressStatistic.YFilter }

func (egressStatistic *MacAccounting_Interfaces_Interface_EgressStatistic) SetFilter(yf yfilter.YFilter) { egressStatistic.YFilter = yf }

func (egressStatistic *MacAccounting_Interfaces_Interface_EgressStatistic) GetGoName(yname string) string {
    if yname == "mac-address" { return "MacAddress" }
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (egressStatistic *MacAccounting_Interfaces_Interface_EgressStatistic) GetSegmentPath() string {
    return "egress-statistic"
}

func (egressStatistic *MacAccounting_Interfaces_Interface_EgressStatistic) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (egressStatistic *MacAccounting_Interfaces_Interface_EgressStatistic) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (egressStatistic *MacAccounting_Interfaces_Interface_EgressStatistic) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mac-address"] = egressStatistic.MacAddress
    leafs["packets"] = egressStatistic.Packets
    leafs["bytes"] = egressStatistic.Bytes
    return leafs
}

func (egressStatistic *MacAccounting_Interfaces_Interface_EgressStatistic) GetBundleName() string { return "cisco_ios_xr" }

func (egressStatistic *MacAccounting_Interfaces_Interface_EgressStatistic) GetYangName() string { return "egress-statistic" }

func (egressStatistic *MacAccounting_Interfaces_Interface_EgressStatistic) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (egressStatistic *MacAccounting_Interfaces_Interface_EgressStatistic) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (egressStatistic *MacAccounting_Interfaces_Interface_EgressStatistic) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (egressStatistic *MacAccounting_Interfaces_Interface_EgressStatistic) SetParent(parent types.Entity) { egressStatistic.parent = parent }

func (egressStatistic *MacAccounting_Interfaces_Interface_EgressStatistic) GetParent() types.Entity { return egressStatistic.parent }

func (egressStatistic *MacAccounting_Interfaces_Interface_EgressStatistic) GetParentYangName() string { return "interface" }

// Vlan
// vlan
type Vlan struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Per node VLAN operational data.
    Nodes Vlan_Nodes
}

func (vlan *Vlan) GetFilter() yfilter.YFilter { return vlan.YFilter }

func (vlan *Vlan) SetFilter(yf yfilter.YFilter) { vlan.YFilter = yf }

func (vlan *Vlan) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (vlan *Vlan) GetSegmentPath() string {
    return "Cisco-IOS-XR-l2-eth-infra-oper:vlan"
}

func (vlan *Vlan) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &vlan.Nodes
    }
    return nil
}

func (vlan *Vlan) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &vlan.Nodes
    return children
}

func (vlan *Vlan) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vlan *Vlan) GetBundleName() string { return "cisco_ios_xr" }

func (vlan *Vlan) GetYangName() string { return "vlan" }

func (vlan *Vlan) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vlan *Vlan) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vlan *Vlan) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vlan *Vlan) SetParent(parent types.Entity) { vlan.parent = parent }

func (vlan *Vlan) GetParent() types.Entity { return vlan.parent }

func (vlan *Vlan) GetParentYangName() string { return "Cisco-IOS-XR-l2-eth-infra-oper" }

// Vlan_Nodes
// Per node VLAN operational data
type Vlan_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The VLAN operational data for a particular node. The type is slice of
    // Vlan_Nodes_Node.
    Node []Vlan_Nodes_Node
}

func (nodes *Vlan_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *Vlan_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *Vlan_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *Vlan_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *Vlan_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vlan_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *Vlan_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *Vlan_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *Vlan_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *Vlan_Nodes) GetYangName() string { return "nodes" }

func (nodes *Vlan_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *Vlan_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *Vlan_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *Vlan_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *Vlan_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *Vlan_Nodes) GetParentYangName() string { return "vlan" }

// Vlan_Nodes_Node
// The VLAN operational data for a particular node
type Vlan_Nodes_Node struct {
    parent types.Entity
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

func (node *Vlan_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *Vlan_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *Vlan_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-id" { return "NodeId" }
    if yname == "trunks" { return "Trunks" }
    if yname == "interfaces" { return "Interfaces" }
    if yname == "tag-allocations" { return "TagAllocations" }
    return ""
}

func (node *Vlan_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-id='" + fmt.Sprintf("%v", node.NodeId) + "']"
}

func (node *Vlan_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "trunks" {
        return &node.Trunks
    }
    if childYangName == "interfaces" {
        return &node.Interfaces
    }
    if childYangName == "tag-allocations" {
        return &node.TagAllocations
    }
    return nil
}

func (node *Vlan_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["trunks"] = &node.Trunks
    children["interfaces"] = &node.Interfaces
    children["tag-allocations"] = &node.TagAllocations
    return children
}

func (node *Vlan_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-id"] = node.NodeId
    return leafs
}

func (node *Vlan_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *Vlan_Nodes_Node) GetYangName() string { return "node" }

func (node *Vlan_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *Vlan_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *Vlan_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *Vlan_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *Vlan_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *Vlan_Nodes_Node) GetParentYangName() string { return "nodes" }

// Vlan_Nodes_Node_Trunks
// VLAN trunk table (specific to this node)
type Vlan_Nodes_Node_Trunks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operational data for trunk interfaces configured with VLANs. The type is
    // slice of Vlan_Nodes_Node_Trunks_Trunk.
    Trunk []Vlan_Nodes_Node_Trunks_Trunk
}

func (trunks *Vlan_Nodes_Node_Trunks) GetFilter() yfilter.YFilter { return trunks.YFilter }

func (trunks *Vlan_Nodes_Node_Trunks) SetFilter(yf yfilter.YFilter) { trunks.YFilter = yf }

func (trunks *Vlan_Nodes_Node_Trunks) GetGoName(yname string) string {
    if yname == "trunk" { return "Trunk" }
    return ""
}

func (trunks *Vlan_Nodes_Node_Trunks) GetSegmentPath() string {
    return "trunks"
}

func (trunks *Vlan_Nodes_Node_Trunks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "trunk" {
        for _, c := range trunks.Trunk {
            if trunks.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vlan_Nodes_Node_Trunks_Trunk{}
        trunks.Trunk = append(trunks.Trunk, child)
        return &trunks.Trunk[len(trunks.Trunk)-1]
    }
    return nil
}

func (trunks *Vlan_Nodes_Node_Trunks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range trunks.Trunk {
        children[trunks.Trunk[i].GetSegmentPath()] = &trunks.Trunk[i]
    }
    return children
}

func (trunks *Vlan_Nodes_Node_Trunks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (trunks *Vlan_Nodes_Node_Trunks) GetBundleName() string { return "cisco_ios_xr" }

func (trunks *Vlan_Nodes_Node_Trunks) GetYangName() string { return "trunks" }

func (trunks *Vlan_Nodes_Node_Trunks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trunks *Vlan_Nodes_Node_Trunks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trunks *Vlan_Nodes_Node_Trunks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trunks *Vlan_Nodes_Node_Trunks) SetParent(parent types.Entity) { trunks.parent = parent }

func (trunks *Vlan_Nodes_Node_Trunks) GetParent() types.Entity { return trunks.parent }

func (trunks *Vlan_Nodes_Node_Trunks) GetParentYangName() string { return "node" }

// Vlan_Nodes_Node_Trunks_Trunk
// Operational data for trunk interfaces
// configured with VLANs
type Vlan_Nodes_Node_Trunks_Trunk struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The interface name. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    Interface interface{}

    // Interface name. The type is string with pattern: [a-zA-Z0-9./-]+.
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
    // pattern: [a-zA-Z0-9./-]+.
    UntaggedInterface interface{}

    // IEEE 802.1Q/802.1ad multicast MAC address filtering. The type is
    // EthFiltering.
    MacFiltering interface{}

    // Layer 2 Transport Subinterfaces.
    Layer2SubInterfaces Vlan_Nodes_Node_Trunks_Trunk_Layer2SubInterfaces

    // Layer 3 Terminated Subinterfaces.
    Layer3SubInterfaces Vlan_Nodes_Node_Trunks_Trunk_Layer3SubInterfaces
}

func (trunk *Vlan_Nodes_Node_Trunks_Trunk) GetFilter() yfilter.YFilter { return trunk.YFilter }

func (trunk *Vlan_Nodes_Node_Trunks_Trunk) SetFilter(yf yfilter.YFilter) { trunk.YFilter = yf }

func (trunk *Vlan_Nodes_Node_Trunks_Trunk) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    if yname == "interface-xr" { return "InterfaceXr" }
    if yname == "state" { return "State" }
    if yname == "mtu" { return "Mtu" }
    if yname == "qinq-outer-ether-type" { return "QinqOuterEtherType" }
    if yname == "dot1ad-count" { return "Dot1AdCount" }
    if yname == "untagged-interface" { return "UntaggedInterface" }
    if yname == "mac-filtering" { return "MacFiltering" }
    if yname == "layer2-sub-interfaces" { return "Layer2SubInterfaces" }
    if yname == "layer3-sub-interfaces" { return "Layer3SubInterfaces" }
    return ""
}

func (trunk *Vlan_Nodes_Node_Trunks_Trunk) GetSegmentPath() string {
    return "trunk" + "[interface='" + fmt.Sprintf("%v", trunk.Interface) + "']"
}

func (trunk *Vlan_Nodes_Node_Trunks_Trunk) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "layer2-sub-interfaces" {
        return &trunk.Layer2SubInterfaces
    }
    if childYangName == "layer3-sub-interfaces" {
        return &trunk.Layer3SubInterfaces
    }
    return nil
}

func (trunk *Vlan_Nodes_Node_Trunks_Trunk) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["layer2-sub-interfaces"] = &trunk.Layer2SubInterfaces
    children["layer3-sub-interfaces"] = &trunk.Layer3SubInterfaces
    return children
}

func (trunk *Vlan_Nodes_Node_Trunks_Trunk) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface"] = trunk.Interface
    leafs["interface-xr"] = trunk.InterfaceXr
    leafs["state"] = trunk.State
    leafs["mtu"] = trunk.Mtu
    leafs["qinq-outer-ether-type"] = trunk.QinqOuterEtherType
    leafs["dot1ad-count"] = trunk.Dot1AdCount
    leafs["untagged-interface"] = trunk.UntaggedInterface
    leafs["mac-filtering"] = trunk.MacFiltering
    return leafs
}

func (trunk *Vlan_Nodes_Node_Trunks_Trunk) GetBundleName() string { return "cisco_ios_xr" }

func (trunk *Vlan_Nodes_Node_Trunks_Trunk) GetYangName() string { return "trunk" }

func (trunk *Vlan_Nodes_Node_Trunks_Trunk) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trunk *Vlan_Nodes_Node_Trunks_Trunk) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trunk *Vlan_Nodes_Node_Trunks_Trunk) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trunk *Vlan_Nodes_Node_Trunks_Trunk) SetParent(parent types.Entity) { trunk.parent = parent }

func (trunk *Vlan_Nodes_Node_Trunks_Trunk) GetParent() types.Entity { return trunk.parent }

func (trunk *Vlan_Nodes_Node_Trunks_Trunk) GetParentYangName() string { return "trunks" }

// Vlan_Nodes_Node_Trunks_Trunk_Layer2SubInterfaces
// Layer 2 Transport Subinterfaces
type Vlan_Nodes_Node_Trunks_Trunk_Layer2SubInterfaces struct {
    parent types.Entity
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

func (layer2SubInterfaces *Vlan_Nodes_Node_Trunks_Trunk_Layer2SubInterfaces) GetFilter() yfilter.YFilter { return layer2SubInterfaces.YFilter }

func (layer2SubInterfaces *Vlan_Nodes_Node_Trunks_Trunk_Layer2SubInterfaces) SetFilter(yf yfilter.YFilter) { layer2SubInterfaces.YFilter = yf }

func (layer2SubInterfaces *Vlan_Nodes_Node_Trunks_Trunk_Layer2SubInterfaces) GetGoName(yname string) string {
    if yname == "total-count" { return "TotalCount" }
    if yname == "dot1q-count" { return "Dot1QCount" }
    if yname == "qin-q-count" { return "QinQCount" }
    if yname == "qin-any-count" { return "QinAnyCount" }
    if yname == "untagged-count" { return "UntaggedCount" }
    if yname == "state-counters" { return "StateCounters" }
    return ""
}

func (layer2SubInterfaces *Vlan_Nodes_Node_Trunks_Trunk_Layer2SubInterfaces) GetSegmentPath() string {
    return "layer2-sub-interfaces"
}

func (layer2SubInterfaces *Vlan_Nodes_Node_Trunks_Trunk_Layer2SubInterfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "state-counters" {
        return &layer2SubInterfaces.StateCounters
    }
    return nil
}

func (layer2SubInterfaces *Vlan_Nodes_Node_Trunks_Trunk_Layer2SubInterfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["state-counters"] = &layer2SubInterfaces.StateCounters
    return children
}

func (layer2SubInterfaces *Vlan_Nodes_Node_Trunks_Trunk_Layer2SubInterfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["total-count"] = layer2SubInterfaces.TotalCount
    leafs["dot1q-count"] = layer2SubInterfaces.Dot1QCount
    leafs["qin-q-count"] = layer2SubInterfaces.QinQCount
    leafs["qin-any-count"] = layer2SubInterfaces.QinAnyCount
    leafs["untagged-count"] = layer2SubInterfaces.UntaggedCount
    return leafs
}

func (layer2SubInterfaces *Vlan_Nodes_Node_Trunks_Trunk_Layer2SubInterfaces) GetBundleName() string { return "cisco_ios_xr" }

func (layer2SubInterfaces *Vlan_Nodes_Node_Trunks_Trunk_Layer2SubInterfaces) GetYangName() string { return "layer2-sub-interfaces" }

func (layer2SubInterfaces *Vlan_Nodes_Node_Trunks_Trunk_Layer2SubInterfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (layer2SubInterfaces *Vlan_Nodes_Node_Trunks_Trunk_Layer2SubInterfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (layer2SubInterfaces *Vlan_Nodes_Node_Trunks_Trunk_Layer2SubInterfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (layer2SubInterfaces *Vlan_Nodes_Node_Trunks_Trunk_Layer2SubInterfaces) SetParent(parent types.Entity) { layer2SubInterfaces.parent = parent }

func (layer2SubInterfaces *Vlan_Nodes_Node_Trunks_Trunk_Layer2SubInterfaces) GetParent() types.Entity { return layer2SubInterfaces.parent }

func (layer2SubInterfaces *Vlan_Nodes_Node_Trunks_Trunk_Layer2SubInterfaces) GetParentYangName() string { return "trunk" }

// Vlan_Nodes_Node_Trunks_Trunk_Layer2SubInterfaces_StateCounters
// Numbers of subinterfaces up, down or
// administratively shut down
type Vlan_Nodes_Node_Trunks_Trunk_Layer2SubInterfaces_StateCounters struct {
    parent types.Entity
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

func (stateCounters *Vlan_Nodes_Node_Trunks_Trunk_Layer2SubInterfaces_StateCounters) GetFilter() yfilter.YFilter { return stateCounters.YFilter }

func (stateCounters *Vlan_Nodes_Node_Trunks_Trunk_Layer2SubInterfaces_StateCounters) SetFilter(yf yfilter.YFilter) { stateCounters.YFilter = yf }

func (stateCounters *Vlan_Nodes_Node_Trunks_Trunk_Layer2SubInterfaces_StateCounters) GetGoName(yname string) string {
    if yname == "up" { return "Up" }
    if yname == "down" { return "Down" }
    if yname == "admin-down" { return "AdminDown" }
    return ""
}

func (stateCounters *Vlan_Nodes_Node_Trunks_Trunk_Layer2SubInterfaces_StateCounters) GetSegmentPath() string {
    return "state-counters"
}

func (stateCounters *Vlan_Nodes_Node_Trunks_Trunk_Layer2SubInterfaces_StateCounters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stateCounters *Vlan_Nodes_Node_Trunks_Trunk_Layer2SubInterfaces_StateCounters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stateCounters *Vlan_Nodes_Node_Trunks_Trunk_Layer2SubInterfaces_StateCounters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["up"] = stateCounters.Up
    leafs["down"] = stateCounters.Down
    leafs["admin-down"] = stateCounters.AdminDown
    return leafs
}

func (stateCounters *Vlan_Nodes_Node_Trunks_Trunk_Layer2SubInterfaces_StateCounters) GetBundleName() string { return "cisco_ios_xr" }

func (stateCounters *Vlan_Nodes_Node_Trunks_Trunk_Layer2SubInterfaces_StateCounters) GetYangName() string { return "state-counters" }

func (stateCounters *Vlan_Nodes_Node_Trunks_Trunk_Layer2SubInterfaces_StateCounters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (stateCounters *Vlan_Nodes_Node_Trunks_Trunk_Layer2SubInterfaces_StateCounters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (stateCounters *Vlan_Nodes_Node_Trunks_Trunk_Layer2SubInterfaces_StateCounters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (stateCounters *Vlan_Nodes_Node_Trunks_Trunk_Layer2SubInterfaces_StateCounters) SetParent(parent types.Entity) { stateCounters.parent = parent }

func (stateCounters *Vlan_Nodes_Node_Trunks_Trunk_Layer2SubInterfaces_StateCounters) GetParent() types.Entity { return stateCounters.parent }

func (stateCounters *Vlan_Nodes_Node_Trunks_Trunk_Layer2SubInterfaces_StateCounters) GetParentYangName() string { return "layer2-sub-interfaces" }

// Vlan_Nodes_Node_Trunks_Trunk_Layer3SubInterfaces
// Layer 3 Terminated Subinterfaces
type Vlan_Nodes_Node_Trunks_Trunk_Layer3SubInterfaces struct {
    parent types.Entity
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

func (layer3SubInterfaces *Vlan_Nodes_Node_Trunks_Trunk_Layer3SubInterfaces) GetFilter() yfilter.YFilter { return layer3SubInterfaces.YFilter }

func (layer3SubInterfaces *Vlan_Nodes_Node_Trunks_Trunk_Layer3SubInterfaces) SetFilter(yf yfilter.YFilter) { layer3SubInterfaces.YFilter = yf }

func (layer3SubInterfaces *Vlan_Nodes_Node_Trunks_Trunk_Layer3SubInterfaces) GetGoName(yname string) string {
    if yname == "total-count" { return "TotalCount" }
    if yname == "dot1q-count" { return "Dot1QCount" }
    if yname == "qin-q-count" { return "QinQCount" }
    if yname == "untagged-count" { return "UntaggedCount" }
    if yname == "native-vlan" { return "NativeVlan" }
    if yname == "state-counters" { return "StateCounters" }
    return ""
}

func (layer3SubInterfaces *Vlan_Nodes_Node_Trunks_Trunk_Layer3SubInterfaces) GetSegmentPath() string {
    return "layer3-sub-interfaces"
}

func (layer3SubInterfaces *Vlan_Nodes_Node_Trunks_Trunk_Layer3SubInterfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "state-counters" {
        return &layer3SubInterfaces.StateCounters
    }
    return nil
}

func (layer3SubInterfaces *Vlan_Nodes_Node_Trunks_Trunk_Layer3SubInterfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["state-counters"] = &layer3SubInterfaces.StateCounters
    return children
}

func (layer3SubInterfaces *Vlan_Nodes_Node_Trunks_Trunk_Layer3SubInterfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["total-count"] = layer3SubInterfaces.TotalCount
    leafs["dot1q-count"] = layer3SubInterfaces.Dot1QCount
    leafs["qin-q-count"] = layer3SubInterfaces.QinQCount
    leafs["untagged-count"] = layer3SubInterfaces.UntaggedCount
    leafs["native-vlan"] = layer3SubInterfaces.NativeVlan
    return leafs
}

func (layer3SubInterfaces *Vlan_Nodes_Node_Trunks_Trunk_Layer3SubInterfaces) GetBundleName() string { return "cisco_ios_xr" }

func (layer3SubInterfaces *Vlan_Nodes_Node_Trunks_Trunk_Layer3SubInterfaces) GetYangName() string { return "layer3-sub-interfaces" }

func (layer3SubInterfaces *Vlan_Nodes_Node_Trunks_Trunk_Layer3SubInterfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (layer3SubInterfaces *Vlan_Nodes_Node_Trunks_Trunk_Layer3SubInterfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (layer3SubInterfaces *Vlan_Nodes_Node_Trunks_Trunk_Layer3SubInterfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (layer3SubInterfaces *Vlan_Nodes_Node_Trunks_Trunk_Layer3SubInterfaces) SetParent(parent types.Entity) { layer3SubInterfaces.parent = parent }

func (layer3SubInterfaces *Vlan_Nodes_Node_Trunks_Trunk_Layer3SubInterfaces) GetParent() types.Entity { return layer3SubInterfaces.parent }

func (layer3SubInterfaces *Vlan_Nodes_Node_Trunks_Trunk_Layer3SubInterfaces) GetParentYangName() string { return "trunk" }

// Vlan_Nodes_Node_Trunks_Trunk_Layer3SubInterfaces_StateCounters
// Numbers of subinterfaces up, down or
// administratively shut down
type Vlan_Nodes_Node_Trunks_Trunk_Layer3SubInterfaces_StateCounters struct {
    parent types.Entity
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

func (stateCounters *Vlan_Nodes_Node_Trunks_Trunk_Layer3SubInterfaces_StateCounters) GetFilter() yfilter.YFilter { return stateCounters.YFilter }

func (stateCounters *Vlan_Nodes_Node_Trunks_Trunk_Layer3SubInterfaces_StateCounters) SetFilter(yf yfilter.YFilter) { stateCounters.YFilter = yf }

func (stateCounters *Vlan_Nodes_Node_Trunks_Trunk_Layer3SubInterfaces_StateCounters) GetGoName(yname string) string {
    if yname == "up" { return "Up" }
    if yname == "down" { return "Down" }
    if yname == "admin-down" { return "AdminDown" }
    return ""
}

func (stateCounters *Vlan_Nodes_Node_Trunks_Trunk_Layer3SubInterfaces_StateCounters) GetSegmentPath() string {
    return "state-counters"
}

func (stateCounters *Vlan_Nodes_Node_Trunks_Trunk_Layer3SubInterfaces_StateCounters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stateCounters *Vlan_Nodes_Node_Trunks_Trunk_Layer3SubInterfaces_StateCounters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stateCounters *Vlan_Nodes_Node_Trunks_Trunk_Layer3SubInterfaces_StateCounters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["up"] = stateCounters.Up
    leafs["down"] = stateCounters.Down
    leafs["admin-down"] = stateCounters.AdminDown
    return leafs
}

func (stateCounters *Vlan_Nodes_Node_Trunks_Trunk_Layer3SubInterfaces_StateCounters) GetBundleName() string { return "cisco_ios_xr" }

func (stateCounters *Vlan_Nodes_Node_Trunks_Trunk_Layer3SubInterfaces_StateCounters) GetYangName() string { return "state-counters" }

func (stateCounters *Vlan_Nodes_Node_Trunks_Trunk_Layer3SubInterfaces_StateCounters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (stateCounters *Vlan_Nodes_Node_Trunks_Trunk_Layer3SubInterfaces_StateCounters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (stateCounters *Vlan_Nodes_Node_Trunks_Trunk_Layer3SubInterfaces_StateCounters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (stateCounters *Vlan_Nodes_Node_Trunks_Trunk_Layer3SubInterfaces_StateCounters) SetParent(parent types.Entity) { stateCounters.parent = parent }

func (stateCounters *Vlan_Nodes_Node_Trunks_Trunk_Layer3SubInterfaces_StateCounters) GetParent() types.Entity { return stateCounters.parent }

func (stateCounters *Vlan_Nodes_Node_Trunks_Trunk_Layer3SubInterfaces_StateCounters) GetParentYangName() string { return "layer3-sub-interfaces" }

// Vlan_Nodes_Node_Interfaces
// VLAN interface table (specific to this node)
type Vlan_Nodes_Node_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operational data for a sub-interface configured with VLANs. The type is
    // slice of Vlan_Nodes_Node_Interfaces_Interface.
    Interface []Vlan_Nodes_Node_Interfaces_Interface
}

func (interfaces *Vlan_Nodes_Node_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *Vlan_Nodes_Node_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *Vlan_Nodes_Node_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *Vlan_Nodes_Node_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *Vlan_Nodes_Node_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vlan_Nodes_Node_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *Vlan_Nodes_Node_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *Vlan_Nodes_Node_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *Vlan_Nodes_Node_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *Vlan_Nodes_Node_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *Vlan_Nodes_Node_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *Vlan_Nodes_Node_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *Vlan_Nodes_Node_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *Vlan_Nodes_Node_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *Vlan_Nodes_Node_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *Vlan_Nodes_Node_Interfaces) GetParentYangName() string { return "node" }

// Vlan_Nodes_Node_Interfaces_Interface
// Operational data for a sub-interface
// configured with VLANs
type Vlan_Nodes_Node_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The interface name. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    Interface interface{}

    // Interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    InterfaceXr interface{}

    // Parent interface. The type is string with pattern: [a-zA-Z0-9./-]+.
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

func (self *Vlan_Nodes_Node_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Vlan_Nodes_Node_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Vlan_Nodes_Node_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    if yname == "interface-xr" { return "InterfaceXr" }
    if yname == "parent-interface" { return "ParentInterface" }
    if yname == "service" { return "Service" }
    if yname == "state" { return "State" }
    if yname == "mtu" { return "Mtu" }
    if yname == "switched-mtu" { return "SwitchedMtu" }
    if yname == "encapsulation-details" { return "EncapsulationDetails" }
    return ""
}

func (self *Vlan_Nodes_Node_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface='" + fmt.Sprintf("%v", self.Interface) + "']"
}

func (self *Vlan_Nodes_Node_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "encapsulation-details" {
        return &self.EncapsulationDetails
    }
    return nil
}

func (self *Vlan_Nodes_Node_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["encapsulation-details"] = &self.EncapsulationDetails
    return children
}

func (self *Vlan_Nodes_Node_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface"] = self.Interface
    leafs["interface-xr"] = self.InterfaceXr
    leafs["parent-interface"] = self.ParentInterface
    leafs["service"] = self.Service
    leafs["state"] = self.State
    leafs["mtu"] = self.Mtu
    leafs["switched-mtu"] = self.SwitchedMtu
    return leafs
}

func (self *Vlan_Nodes_Node_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *Vlan_Nodes_Node_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *Vlan_Nodes_Node_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Vlan_Nodes_Node_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Vlan_Nodes_Node_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Vlan_Nodes_Node_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *Vlan_Nodes_Node_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *Vlan_Nodes_Node_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails
// Encapsulation type and tag stack
type Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails struct {
    parent types.Entity
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

func (encapsulationDetails *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails) GetFilter() yfilter.YFilter { return encapsulationDetails.YFilter }

func (encapsulationDetails *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails) SetFilter(yf yfilter.YFilter) { encapsulationDetails.YFilter = yf }

func (encapsulationDetails *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails) GetGoName(yname string) string {
    if yname == "vlan-encapsulation" { return "VlanEncapsulation" }
    if yname == "tag" { return "Tag" }
    if yname == "outer-tag" { return "OuterTag" }
    if yname == "native-tag" { return "NativeTag" }
    if yname == "dot1ad-tag" { return "Dot1AdTag" }
    if yname == "dot1ad-native-tag" { return "Dot1AdNativeTag" }
    if yname == "dot1ad-outer-tag" { return "Dot1AdOuterTag" }
    if yname == "stack" { return "Stack" }
    if yname == "service-instance-details" { return "ServiceInstanceDetails" }
    if yname == "dot1ad-dot1q-stack" { return "Dot1AdDot1QStack" }
    return ""
}

func (encapsulationDetails *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails) GetSegmentPath() string {
    return "encapsulation-details"
}

func (encapsulationDetails *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stack" {
        return &encapsulationDetails.Stack
    }
    if childYangName == "service-instance-details" {
        return &encapsulationDetails.ServiceInstanceDetails
    }
    if childYangName == "dot1ad-dot1q-stack" {
        return &encapsulationDetails.Dot1AdDot1QStack
    }
    return nil
}

func (encapsulationDetails *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["stack"] = &encapsulationDetails.Stack
    children["service-instance-details"] = &encapsulationDetails.ServiceInstanceDetails
    children["dot1ad-dot1q-stack"] = &encapsulationDetails.Dot1AdDot1QStack
    return children
}

func (encapsulationDetails *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vlan-encapsulation"] = encapsulationDetails.VlanEncapsulation
    leafs["tag"] = encapsulationDetails.Tag
    leafs["outer-tag"] = encapsulationDetails.OuterTag
    leafs["native-tag"] = encapsulationDetails.NativeTag
    leafs["dot1ad-tag"] = encapsulationDetails.Dot1AdTag
    leafs["dot1ad-native-tag"] = encapsulationDetails.Dot1AdNativeTag
    leafs["dot1ad-outer-tag"] = encapsulationDetails.Dot1AdOuterTag
    return leafs
}

func (encapsulationDetails *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails) GetBundleName() string { return "cisco_ios_xr" }

func (encapsulationDetails *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails) GetYangName() string { return "encapsulation-details" }

func (encapsulationDetails *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (encapsulationDetails *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (encapsulationDetails *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (encapsulationDetails *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails) SetParent(parent types.Entity) { encapsulationDetails.parent = parent }

func (encapsulationDetails *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails) GetParent() types.Entity { return encapsulationDetails.parent }

func (encapsulationDetails *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails) GetParentYangName() string { return "interface" }

// Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_Stack
// Stack value
type Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_Stack struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Outer tag value. The type is interface{} with range: 0..65535.
    OuterTag interface{}

    // Second tag value. The type is interface{} with range: 0..65535.
    SecondTag interface{}
}

func (stack *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_Stack) GetFilter() yfilter.YFilter { return stack.YFilter }

func (stack *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_Stack) SetFilter(yf yfilter.YFilter) { stack.YFilter = yf }

func (stack *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_Stack) GetGoName(yname string) string {
    if yname == "outer-tag" { return "OuterTag" }
    if yname == "second-tag" { return "SecondTag" }
    return ""
}

func (stack *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_Stack) GetSegmentPath() string {
    return "stack"
}

func (stack *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_Stack) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stack *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_Stack) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stack *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_Stack) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["outer-tag"] = stack.OuterTag
    leafs["second-tag"] = stack.SecondTag
    return leafs
}

func (stack *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_Stack) GetBundleName() string { return "cisco_ios_xr" }

func (stack *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_Stack) GetYangName() string { return "stack" }

func (stack *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_Stack) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (stack *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_Stack) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (stack *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_Stack) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (stack *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_Stack) SetParent(parent types.Entity) { stack.parent = parent }

func (stack *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_Stack) GetParent() types.Entity { return stack.parent }

func (stack *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_Stack) GetParentYangName() string { return "encapsulation-details" }

// Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails
// Service Instance encapsulation
type Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails struct {
    parent types.Entity
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
    TagsToMatch []Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch

    // VLAN tags pushed on egress. The type is slice of
    // Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_Pushe.
    Pushe []Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_Pushe
}

func (serviceInstanceDetails *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails) GetFilter() yfilter.YFilter { return serviceInstanceDetails.YFilter }

func (serviceInstanceDetails *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails) SetFilter(yf yfilter.YFilter) { serviceInstanceDetails.YFilter = yf }

func (serviceInstanceDetails *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails) GetGoName(yname string) string {
    if yname == "payload-ethertype" { return "PayloadEthertype" }
    if yname == "tags-popped" { return "TagsPopped" }
    if yname == "is-exact-match" { return "IsExactMatch" }
    if yname == "is-native-vlan" { return "IsNativeVlan" }
    if yname == "is-native-preserving" { return "IsNativePreserving" }
    if yname == "source-mac-match" { return "SourceMacMatch" }
    if yname == "destination-mac-match" { return "DestinationMacMatch" }
    if yname == "local-traffic-stack" { return "LocalTrafficStack" }
    if yname == "tags-to-match" { return "TagsToMatch" }
    if yname == "pushe" { return "Pushe" }
    return ""
}

func (serviceInstanceDetails *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails) GetSegmentPath() string {
    return "service-instance-details"
}

func (serviceInstanceDetails *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "local-traffic-stack" {
        return &serviceInstanceDetails.LocalTrafficStack
    }
    if childYangName == "tags-to-match" {
        for _, c := range serviceInstanceDetails.TagsToMatch {
            if serviceInstanceDetails.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch{}
        serviceInstanceDetails.TagsToMatch = append(serviceInstanceDetails.TagsToMatch, child)
        return &serviceInstanceDetails.TagsToMatch[len(serviceInstanceDetails.TagsToMatch)-1]
    }
    if childYangName == "pushe" {
        for _, c := range serviceInstanceDetails.Pushe {
            if serviceInstanceDetails.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_Pushe{}
        serviceInstanceDetails.Pushe = append(serviceInstanceDetails.Pushe, child)
        return &serviceInstanceDetails.Pushe[len(serviceInstanceDetails.Pushe)-1]
    }
    return nil
}

func (serviceInstanceDetails *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["local-traffic-stack"] = &serviceInstanceDetails.LocalTrafficStack
    for i := range serviceInstanceDetails.TagsToMatch {
        children[serviceInstanceDetails.TagsToMatch[i].GetSegmentPath()] = &serviceInstanceDetails.TagsToMatch[i]
    }
    for i := range serviceInstanceDetails.Pushe {
        children[serviceInstanceDetails.Pushe[i].GetSegmentPath()] = &serviceInstanceDetails.Pushe[i]
    }
    return children
}

func (serviceInstanceDetails *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["payload-ethertype"] = serviceInstanceDetails.PayloadEthertype
    leafs["tags-popped"] = serviceInstanceDetails.TagsPopped
    leafs["is-exact-match"] = serviceInstanceDetails.IsExactMatch
    leafs["is-native-vlan"] = serviceInstanceDetails.IsNativeVlan
    leafs["is-native-preserving"] = serviceInstanceDetails.IsNativePreserving
    leafs["source-mac-match"] = serviceInstanceDetails.SourceMacMatch
    leafs["destination-mac-match"] = serviceInstanceDetails.DestinationMacMatch
    return leafs
}

func (serviceInstanceDetails *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails) GetBundleName() string { return "cisco_ios_xr" }

func (serviceInstanceDetails *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails) GetYangName() string { return "service-instance-details" }

func (serviceInstanceDetails *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (serviceInstanceDetails *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (serviceInstanceDetails *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (serviceInstanceDetails *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails) SetParent(parent types.Entity) { serviceInstanceDetails.parent = parent }

func (serviceInstanceDetails *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails) GetParent() types.Entity { return serviceInstanceDetails.parent }

func (serviceInstanceDetails *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails) GetParentYangName() string { return "encapsulation-details" }

// Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack
// VLAN tags for locally-sourced traffic
type Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VLAN tags for locally-sourced traffic. The type is slice of
    // Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag.
    LocalTrafficTag []Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag
}

func (localTrafficStack *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack) GetFilter() yfilter.YFilter { return localTrafficStack.YFilter }

func (localTrafficStack *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack) SetFilter(yf yfilter.YFilter) { localTrafficStack.YFilter = yf }

func (localTrafficStack *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack) GetGoName(yname string) string {
    if yname == "local-traffic-tag" { return "LocalTrafficTag" }
    return ""
}

func (localTrafficStack *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack) GetSegmentPath() string {
    return "local-traffic-stack"
}

func (localTrafficStack *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "local-traffic-tag" {
        for _, c := range localTrafficStack.LocalTrafficTag {
            if localTrafficStack.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag{}
        localTrafficStack.LocalTrafficTag = append(localTrafficStack.LocalTrafficTag, child)
        return &localTrafficStack.LocalTrafficTag[len(localTrafficStack.LocalTrafficTag)-1]
    }
    return nil
}

func (localTrafficStack *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range localTrafficStack.LocalTrafficTag {
        children[localTrafficStack.LocalTrafficTag[i].GetSegmentPath()] = &localTrafficStack.LocalTrafficTag[i]
    }
    return children
}

func (localTrafficStack *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (localTrafficStack *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack) GetBundleName() string { return "cisco_ios_xr" }

func (localTrafficStack *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack) GetYangName() string { return "local-traffic-stack" }

func (localTrafficStack *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (localTrafficStack *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (localTrafficStack *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (localTrafficStack *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack) SetParent(parent types.Entity) { localTrafficStack.parent = parent }

func (localTrafficStack *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack) GetParent() types.Entity { return localTrafficStack.parent }

func (localTrafficStack *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack) GetParentYangName() string { return "service-instance-details" }

// Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag
// VLAN tags for locally-sourced traffic
type Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Ethertype of tag. The type is EfpTagEtype.
    Ethertype interface{}

    // VLAN Id. The type is interface{} with range: 0..65535.
    VlanId interface{}
}

func (localTrafficTag *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag) GetFilter() yfilter.YFilter { return localTrafficTag.YFilter }

func (localTrafficTag *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag) SetFilter(yf yfilter.YFilter) { localTrafficTag.YFilter = yf }

func (localTrafficTag *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag) GetGoName(yname string) string {
    if yname == "ethertype" { return "Ethertype" }
    if yname == "vlan-id" { return "VlanId" }
    return ""
}

func (localTrafficTag *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag) GetSegmentPath() string {
    return "local-traffic-tag"
}

func (localTrafficTag *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (localTrafficTag *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (localTrafficTag *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ethertype"] = localTrafficTag.Ethertype
    leafs["vlan-id"] = localTrafficTag.VlanId
    return leafs
}

func (localTrafficTag *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag) GetBundleName() string { return "cisco_ios_xr" }

func (localTrafficTag *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag) GetYangName() string { return "local-traffic-tag" }

func (localTrafficTag *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (localTrafficTag *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (localTrafficTag *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (localTrafficTag *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag) SetParent(parent types.Entity) { localTrafficTag.parent = parent }

func (localTrafficTag *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag) GetParent() types.Entity { return localTrafficTag.parent }

func (localTrafficTag *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag) GetParentYangName() string { return "local-traffic-stack" }

// Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch
// Tags to match on ingress packets
type Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Ethertype of tag to match. The type is EfpTagEtype.
    Ethertype interface{}

    // Priority to match. The type is EfpTagPriority.
    Priority interface{}

    // VLAN Ids to match. The type is slice of
    // Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange.
    VlanRange []Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange
}

func (tagsToMatch *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch) GetFilter() yfilter.YFilter { return tagsToMatch.YFilter }

func (tagsToMatch *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch) SetFilter(yf yfilter.YFilter) { tagsToMatch.YFilter = yf }

func (tagsToMatch *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch) GetGoName(yname string) string {
    if yname == "ethertype" { return "Ethertype" }
    if yname == "priority" { return "Priority" }
    if yname == "vlan-range" { return "VlanRange" }
    return ""
}

func (tagsToMatch *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch) GetSegmentPath() string {
    return "tags-to-match"
}

func (tagsToMatch *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vlan-range" {
        for _, c := range tagsToMatch.VlanRange {
            if tagsToMatch.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange{}
        tagsToMatch.VlanRange = append(tagsToMatch.VlanRange, child)
        return &tagsToMatch.VlanRange[len(tagsToMatch.VlanRange)-1]
    }
    return nil
}

func (tagsToMatch *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range tagsToMatch.VlanRange {
        children[tagsToMatch.VlanRange[i].GetSegmentPath()] = &tagsToMatch.VlanRange[i]
    }
    return children
}

func (tagsToMatch *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ethertype"] = tagsToMatch.Ethertype
    leafs["priority"] = tagsToMatch.Priority
    return leafs
}

func (tagsToMatch *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch) GetBundleName() string { return "cisco_ios_xr" }

func (tagsToMatch *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch) GetYangName() string { return "tags-to-match" }

func (tagsToMatch *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tagsToMatch *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tagsToMatch *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tagsToMatch *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch) SetParent(parent types.Entity) { tagsToMatch.parent = parent }

func (tagsToMatch *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch) GetParent() types.Entity { return tagsToMatch.parent }

func (tagsToMatch *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch) GetParentYangName() string { return "service-instance-details" }

// Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange
// VLAN Ids to match
type Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VLAN ID Low. The type is interface{} with range: 0..65535.
    VlanIdLow interface{}

    // VLAN ID High. The type is interface{} with range: 0..65535.
    VlanIdHigh interface{}
}

func (vlanRange *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange) GetFilter() yfilter.YFilter { return vlanRange.YFilter }

func (vlanRange *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange) SetFilter(yf yfilter.YFilter) { vlanRange.YFilter = yf }

func (vlanRange *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange) GetGoName(yname string) string {
    if yname == "vlan-id-low" { return "VlanIdLow" }
    if yname == "vlan-id-high" { return "VlanIdHigh" }
    return ""
}

func (vlanRange *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange) GetSegmentPath() string {
    return "vlan-range"
}

func (vlanRange *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vlanRange *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vlanRange *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vlan-id-low"] = vlanRange.VlanIdLow
    leafs["vlan-id-high"] = vlanRange.VlanIdHigh
    return leafs
}

func (vlanRange *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange) GetBundleName() string { return "cisco_ios_xr" }

func (vlanRange *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange) GetYangName() string { return "vlan-range" }

func (vlanRange *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vlanRange *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vlanRange *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vlanRange *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange) SetParent(parent types.Entity) { vlanRange.parent = parent }

func (vlanRange *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange) GetParent() types.Entity { return vlanRange.parent }

func (vlanRange *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange) GetParentYangName() string { return "tags-to-match" }

// Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_Pushe
// VLAN tags pushed on egress
type Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_Pushe struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Ethertype of tag. The type is EfpTagEtype.
    Ethertype interface{}

    // VLAN Id. The type is interface{} with range: 0..65535.
    VlanId interface{}
}

func (pushe *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_Pushe) GetFilter() yfilter.YFilter { return pushe.YFilter }

func (pushe *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_Pushe) SetFilter(yf yfilter.YFilter) { pushe.YFilter = yf }

func (pushe *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_Pushe) GetGoName(yname string) string {
    if yname == "ethertype" { return "Ethertype" }
    if yname == "vlan-id" { return "VlanId" }
    return ""
}

func (pushe *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_Pushe) GetSegmentPath() string {
    return "pushe"
}

func (pushe *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_Pushe) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pushe *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_Pushe) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pushe *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_Pushe) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ethertype"] = pushe.Ethertype
    leafs["vlan-id"] = pushe.VlanId
    return leafs
}

func (pushe *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_Pushe) GetBundleName() string { return "cisco_ios_xr" }

func (pushe *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_Pushe) GetYangName() string { return "pushe" }

func (pushe *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_Pushe) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pushe *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_Pushe) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pushe *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_Pushe) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pushe *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_Pushe) SetParent(parent types.Entity) { pushe.parent = parent }

func (pushe *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_Pushe) GetParent() types.Entity { return pushe.parent }

func (pushe *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_ServiceInstanceDetails_Pushe) GetParentYangName() string { return "service-instance-details" }

// Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_Dot1AdDot1QStack
// 802.1ad 802.1Q stack value
type Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_Dot1AdDot1QStack struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Outer tag value. The type is interface{} with range: 0..65535.
    OuterTag interface{}

    // Second tag value. The type is interface{} with range: 0..65535.
    SecondTag interface{}
}

func (dot1AdDot1QStack *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_Dot1AdDot1QStack) GetFilter() yfilter.YFilter { return dot1AdDot1QStack.YFilter }

func (dot1AdDot1QStack *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_Dot1AdDot1QStack) SetFilter(yf yfilter.YFilter) { dot1AdDot1QStack.YFilter = yf }

func (dot1AdDot1QStack *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_Dot1AdDot1QStack) GetGoName(yname string) string {
    if yname == "outer-tag" { return "OuterTag" }
    if yname == "second-tag" { return "SecondTag" }
    return ""
}

func (dot1AdDot1QStack *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_Dot1AdDot1QStack) GetSegmentPath() string {
    return "dot1ad-dot1q-stack"
}

func (dot1AdDot1QStack *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_Dot1AdDot1QStack) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dot1AdDot1QStack *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_Dot1AdDot1QStack) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dot1AdDot1QStack *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_Dot1AdDot1QStack) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["outer-tag"] = dot1AdDot1QStack.OuterTag
    leafs["second-tag"] = dot1AdDot1QStack.SecondTag
    return leafs
}

func (dot1AdDot1QStack *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_Dot1AdDot1QStack) GetBundleName() string { return "cisco_ios_xr" }

func (dot1AdDot1QStack *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_Dot1AdDot1QStack) GetYangName() string { return "dot1ad-dot1q-stack" }

func (dot1AdDot1QStack *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_Dot1AdDot1QStack) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dot1AdDot1QStack *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_Dot1AdDot1QStack) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dot1AdDot1QStack *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_Dot1AdDot1QStack) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dot1AdDot1QStack *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_Dot1AdDot1QStack) SetParent(parent types.Entity) { dot1AdDot1QStack.parent = parent }

func (dot1AdDot1QStack *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_Dot1AdDot1QStack) GetParent() types.Entity { return dot1AdDot1QStack.parent }

func (dot1AdDot1QStack *Vlan_Nodes_Node_Interfaces_Interface_EncapsulationDetails_Dot1AdDot1QStack) GetParentYangName() string { return "encapsulation-details" }

// Vlan_Nodes_Node_TagAllocations
// VLAN tag allocation table (specific to this
// node)
type Vlan_Nodes_Node_TagAllocations struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operational data for a sub-interface configured with VLANs. The type is
    // slice of Vlan_Nodes_Node_TagAllocations_TagAllocation.
    TagAllocation []Vlan_Nodes_Node_TagAllocations_TagAllocation
}

func (tagAllocations *Vlan_Nodes_Node_TagAllocations) GetFilter() yfilter.YFilter { return tagAllocations.YFilter }

func (tagAllocations *Vlan_Nodes_Node_TagAllocations) SetFilter(yf yfilter.YFilter) { tagAllocations.YFilter = yf }

func (tagAllocations *Vlan_Nodes_Node_TagAllocations) GetGoName(yname string) string {
    if yname == "tag-allocation" { return "TagAllocation" }
    return ""
}

func (tagAllocations *Vlan_Nodes_Node_TagAllocations) GetSegmentPath() string {
    return "tag-allocations"
}

func (tagAllocations *Vlan_Nodes_Node_TagAllocations) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tag-allocation" {
        for _, c := range tagAllocations.TagAllocation {
            if tagAllocations.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vlan_Nodes_Node_TagAllocations_TagAllocation{}
        tagAllocations.TagAllocation = append(tagAllocations.TagAllocation, child)
        return &tagAllocations.TagAllocation[len(tagAllocations.TagAllocation)-1]
    }
    return nil
}

func (tagAllocations *Vlan_Nodes_Node_TagAllocations) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range tagAllocations.TagAllocation {
        children[tagAllocations.TagAllocation[i].GetSegmentPath()] = &tagAllocations.TagAllocation[i]
    }
    return children
}

func (tagAllocations *Vlan_Nodes_Node_TagAllocations) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (tagAllocations *Vlan_Nodes_Node_TagAllocations) GetBundleName() string { return "cisco_ios_xr" }

func (tagAllocations *Vlan_Nodes_Node_TagAllocations) GetYangName() string { return "tag-allocations" }

func (tagAllocations *Vlan_Nodes_Node_TagAllocations) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tagAllocations *Vlan_Nodes_Node_TagAllocations) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tagAllocations *Vlan_Nodes_Node_TagAllocations) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tagAllocations *Vlan_Nodes_Node_TagAllocations) SetParent(parent types.Entity) { tagAllocations.parent = parent }

func (tagAllocations *Vlan_Nodes_Node_TagAllocations) GetParent() types.Entity { return tagAllocations.parent }

func (tagAllocations *Vlan_Nodes_Node_TagAllocations) GetParentYangName() string { return "node" }

// Vlan_Nodes_Node_TagAllocations_TagAllocation
// Operational data for a sub-interface
// configured with VLANs
type Vlan_Nodes_Node_TagAllocations_TagAllocation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The interface name. The type is string with pattern: [a-zA-Z0-9./-]+.
    Interface interface{}

    // The first (outermost) tag. The type is interface{} with range: 1..4094.
    FirstTag interface{}

    // The second tag. The type is one of the following types: enumeration
    // VlanTagOrAny, or int with range: 1..4096.
    SecondTag interface{}

    // Interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    InterfaceXr interface{}

    // Parent interface. The type is string with pattern: [a-zA-Z0-9./-]+.
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

func (tagAllocation *Vlan_Nodes_Node_TagAllocations_TagAllocation) GetFilter() yfilter.YFilter { return tagAllocation.YFilter }

func (tagAllocation *Vlan_Nodes_Node_TagAllocations_TagAllocation) SetFilter(yf yfilter.YFilter) { tagAllocation.YFilter = yf }

func (tagAllocation *Vlan_Nodes_Node_TagAllocations_TagAllocation) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    if yname == "first-tag" { return "FirstTag" }
    if yname == "second-tag" { return "SecondTag" }
    if yname == "interface-xr" { return "InterfaceXr" }
    if yname == "parent-interface" { return "ParentInterface" }
    if yname == "service" { return "Service" }
    if yname == "state" { return "State" }
    if yname == "mtu" { return "Mtu" }
    if yname == "switched-mtu" { return "SwitchedMtu" }
    if yname == "encapsulation-details" { return "EncapsulationDetails" }
    return ""
}

func (tagAllocation *Vlan_Nodes_Node_TagAllocations_TagAllocation) GetSegmentPath() string {
    return "tag-allocation"
}

func (tagAllocation *Vlan_Nodes_Node_TagAllocations_TagAllocation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "encapsulation-details" {
        return &tagAllocation.EncapsulationDetails
    }
    return nil
}

func (tagAllocation *Vlan_Nodes_Node_TagAllocations_TagAllocation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["encapsulation-details"] = &tagAllocation.EncapsulationDetails
    return children
}

func (tagAllocation *Vlan_Nodes_Node_TagAllocations_TagAllocation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface"] = tagAllocation.Interface
    leafs["first-tag"] = tagAllocation.FirstTag
    leafs["second-tag"] = tagAllocation.SecondTag
    leafs["interface-xr"] = tagAllocation.InterfaceXr
    leafs["parent-interface"] = tagAllocation.ParentInterface
    leafs["service"] = tagAllocation.Service
    leafs["state"] = tagAllocation.State
    leafs["mtu"] = tagAllocation.Mtu
    leafs["switched-mtu"] = tagAllocation.SwitchedMtu
    return leafs
}

func (tagAllocation *Vlan_Nodes_Node_TagAllocations_TagAllocation) GetBundleName() string { return "cisco_ios_xr" }

func (tagAllocation *Vlan_Nodes_Node_TagAllocations_TagAllocation) GetYangName() string { return "tag-allocation" }

func (tagAllocation *Vlan_Nodes_Node_TagAllocations_TagAllocation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tagAllocation *Vlan_Nodes_Node_TagAllocations_TagAllocation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tagAllocation *Vlan_Nodes_Node_TagAllocations_TagAllocation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tagAllocation *Vlan_Nodes_Node_TagAllocations_TagAllocation) SetParent(parent types.Entity) { tagAllocation.parent = parent }

func (tagAllocation *Vlan_Nodes_Node_TagAllocations_TagAllocation) GetParent() types.Entity { return tagAllocation.parent }

func (tagAllocation *Vlan_Nodes_Node_TagAllocations_TagAllocation) GetParentYangName() string { return "tag-allocations" }

// Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails
// Encapsulation type and tag stack
type Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails struct {
    parent types.Entity
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

func (encapsulationDetails *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails) GetFilter() yfilter.YFilter { return encapsulationDetails.YFilter }

func (encapsulationDetails *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails) SetFilter(yf yfilter.YFilter) { encapsulationDetails.YFilter = yf }

func (encapsulationDetails *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails) GetGoName(yname string) string {
    if yname == "vlan-encapsulation" { return "VlanEncapsulation" }
    if yname == "tag" { return "Tag" }
    if yname == "outer-tag" { return "OuterTag" }
    if yname == "native-tag" { return "NativeTag" }
    if yname == "dot1ad-tag" { return "Dot1AdTag" }
    if yname == "dot1ad-native-tag" { return "Dot1AdNativeTag" }
    if yname == "dot1ad-outer-tag" { return "Dot1AdOuterTag" }
    if yname == "stack" { return "Stack" }
    if yname == "service-instance-details" { return "ServiceInstanceDetails" }
    if yname == "dot1ad-dot1q-stack" { return "Dot1AdDot1QStack" }
    return ""
}

func (encapsulationDetails *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails) GetSegmentPath() string {
    return "encapsulation-details"
}

func (encapsulationDetails *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stack" {
        return &encapsulationDetails.Stack
    }
    if childYangName == "service-instance-details" {
        return &encapsulationDetails.ServiceInstanceDetails
    }
    if childYangName == "dot1ad-dot1q-stack" {
        return &encapsulationDetails.Dot1AdDot1QStack
    }
    return nil
}

func (encapsulationDetails *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["stack"] = &encapsulationDetails.Stack
    children["service-instance-details"] = &encapsulationDetails.ServiceInstanceDetails
    children["dot1ad-dot1q-stack"] = &encapsulationDetails.Dot1AdDot1QStack
    return children
}

func (encapsulationDetails *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vlan-encapsulation"] = encapsulationDetails.VlanEncapsulation
    leafs["tag"] = encapsulationDetails.Tag
    leafs["outer-tag"] = encapsulationDetails.OuterTag
    leafs["native-tag"] = encapsulationDetails.NativeTag
    leafs["dot1ad-tag"] = encapsulationDetails.Dot1AdTag
    leafs["dot1ad-native-tag"] = encapsulationDetails.Dot1AdNativeTag
    leafs["dot1ad-outer-tag"] = encapsulationDetails.Dot1AdOuterTag
    return leafs
}

func (encapsulationDetails *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails) GetBundleName() string { return "cisco_ios_xr" }

func (encapsulationDetails *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails) GetYangName() string { return "encapsulation-details" }

func (encapsulationDetails *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (encapsulationDetails *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (encapsulationDetails *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (encapsulationDetails *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails) SetParent(parent types.Entity) { encapsulationDetails.parent = parent }

func (encapsulationDetails *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails) GetParent() types.Entity { return encapsulationDetails.parent }

func (encapsulationDetails *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails) GetParentYangName() string { return "tag-allocation" }

// Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_Stack
// Stack value
type Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_Stack struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Outer tag value. The type is interface{} with range: 0..65535.
    OuterTag interface{}

    // Second tag value. The type is interface{} with range: 0..65535.
    SecondTag interface{}
}

func (stack *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_Stack) GetFilter() yfilter.YFilter { return stack.YFilter }

func (stack *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_Stack) SetFilter(yf yfilter.YFilter) { stack.YFilter = yf }

func (stack *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_Stack) GetGoName(yname string) string {
    if yname == "outer-tag" { return "OuterTag" }
    if yname == "second-tag" { return "SecondTag" }
    return ""
}

func (stack *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_Stack) GetSegmentPath() string {
    return "stack"
}

func (stack *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_Stack) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stack *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_Stack) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stack *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_Stack) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["outer-tag"] = stack.OuterTag
    leafs["second-tag"] = stack.SecondTag
    return leafs
}

func (stack *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_Stack) GetBundleName() string { return "cisco_ios_xr" }

func (stack *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_Stack) GetYangName() string { return "stack" }

func (stack *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_Stack) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (stack *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_Stack) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (stack *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_Stack) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (stack *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_Stack) SetParent(parent types.Entity) { stack.parent = parent }

func (stack *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_Stack) GetParent() types.Entity { return stack.parent }

func (stack *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_Stack) GetParentYangName() string { return "encapsulation-details" }

// Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails
// Service Instance encapsulation
type Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails struct {
    parent types.Entity
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
    TagsToMatch []Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch

    // VLAN tags pushed on egress. The type is slice of
    // Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_Pushe.
    Pushe []Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_Pushe
}

func (serviceInstanceDetails *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails) GetFilter() yfilter.YFilter { return serviceInstanceDetails.YFilter }

func (serviceInstanceDetails *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails) SetFilter(yf yfilter.YFilter) { serviceInstanceDetails.YFilter = yf }

func (serviceInstanceDetails *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails) GetGoName(yname string) string {
    if yname == "payload-ethertype" { return "PayloadEthertype" }
    if yname == "tags-popped" { return "TagsPopped" }
    if yname == "is-exact-match" { return "IsExactMatch" }
    if yname == "is-native-vlan" { return "IsNativeVlan" }
    if yname == "is-native-preserving" { return "IsNativePreserving" }
    if yname == "source-mac-match" { return "SourceMacMatch" }
    if yname == "destination-mac-match" { return "DestinationMacMatch" }
    if yname == "local-traffic-stack" { return "LocalTrafficStack" }
    if yname == "tags-to-match" { return "TagsToMatch" }
    if yname == "pushe" { return "Pushe" }
    return ""
}

func (serviceInstanceDetails *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails) GetSegmentPath() string {
    return "service-instance-details"
}

func (serviceInstanceDetails *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "local-traffic-stack" {
        return &serviceInstanceDetails.LocalTrafficStack
    }
    if childYangName == "tags-to-match" {
        for _, c := range serviceInstanceDetails.TagsToMatch {
            if serviceInstanceDetails.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch{}
        serviceInstanceDetails.TagsToMatch = append(serviceInstanceDetails.TagsToMatch, child)
        return &serviceInstanceDetails.TagsToMatch[len(serviceInstanceDetails.TagsToMatch)-1]
    }
    if childYangName == "pushe" {
        for _, c := range serviceInstanceDetails.Pushe {
            if serviceInstanceDetails.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_Pushe{}
        serviceInstanceDetails.Pushe = append(serviceInstanceDetails.Pushe, child)
        return &serviceInstanceDetails.Pushe[len(serviceInstanceDetails.Pushe)-1]
    }
    return nil
}

func (serviceInstanceDetails *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["local-traffic-stack"] = &serviceInstanceDetails.LocalTrafficStack
    for i := range serviceInstanceDetails.TagsToMatch {
        children[serviceInstanceDetails.TagsToMatch[i].GetSegmentPath()] = &serviceInstanceDetails.TagsToMatch[i]
    }
    for i := range serviceInstanceDetails.Pushe {
        children[serviceInstanceDetails.Pushe[i].GetSegmentPath()] = &serviceInstanceDetails.Pushe[i]
    }
    return children
}

func (serviceInstanceDetails *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["payload-ethertype"] = serviceInstanceDetails.PayloadEthertype
    leafs["tags-popped"] = serviceInstanceDetails.TagsPopped
    leafs["is-exact-match"] = serviceInstanceDetails.IsExactMatch
    leafs["is-native-vlan"] = serviceInstanceDetails.IsNativeVlan
    leafs["is-native-preserving"] = serviceInstanceDetails.IsNativePreserving
    leafs["source-mac-match"] = serviceInstanceDetails.SourceMacMatch
    leafs["destination-mac-match"] = serviceInstanceDetails.DestinationMacMatch
    return leafs
}

func (serviceInstanceDetails *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails) GetBundleName() string { return "cisco_ios_xr" }

func (serviceInstanceDetails *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails) GetYangName() string { return "service-instance-details" }

func (serviceInstanceDetails *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (serviceInstanceDetails *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (serviceInstanceDetails *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (serviceInstanceDetails *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails) SetParent(parent types.Entity) { serviceInstanceDetails.parent = parent }

func (serviceInstanceDetails *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails) GetParent() types.Entity { return serviceInstanceDetails.parent }

func (serviceInstanceDetails *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails) GetParentYangName() string { return "encapsulation-details" }

// Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack
// VLAN tags for locally-sourced traffic
type Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VLAN tags for locally-sourced traffic. The type is slice of
    // Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag.
    LocalTrafficTag []Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag
}

func (localTrafficStack *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack) GetFilter() yfilter.YFilter { return localTrafficStack.YFilter }

func (localTrafficStack *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack) SetFilter(yf yfilter.YFilter) { localTrafficStack.YFilter = yf }

func (localTrafficStack *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack) GetGoName(yname string) string {
    if yname == "local-traffic-tag" { return "LocalTrafficTag" }
    return ""
}

func (localTrafficStack *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack) GetSegmentPath() string {
    return "local-traffic-stack"
}

func (localTrafficStack *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "local-traffic-tag" {
        for _, c := range localTrafficStack.LocalTrafficTag {
            if localTrafficStack.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag{}
        localTrafficStack.LocalTrafficTag = append(localTrafficStack.LocalTrafficTag, child)
        return &localTrafficStack.LocalTrafficTag[len(localTrafficStack.LocalTrafficTag)-1]
    }
    return nil
}

func (localTrafficStack *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range localTrafficStack.LocalTrafficTag {
        children[localTrafficStack.LocalTrafficTag[i].GetSegmentPath()] = &localTrafficStack.LocalTrafficTag[i]
    }
    return children
}

func (localTrafficStack *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (localTrafficStack *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack) GetBundleName() string { return "cisco_ios_xr" }

func (localTrafficStack *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack) GetYangName() string { return "local-traffic-stack" }

func (localTrafficStack *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (localTrafficStack *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (localTrafficStack *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (localTrafficStack *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack) SetParent(parent types.Entity) { localTrafficStack.parent = parent }

func (localTrafficStack *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack) GetParent() types.Entity { return localTrafficStack.parent }

func (localTrafficStack *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack) GetParentYangName() string { return "service-instance-details" }

// Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag
// VLAN tags for locally-sourced traffic
type Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Ethertype of tag. The type is EfpTagEtype.
    Ethertype interface{}

    // VLAN Id. The type is interface{} with range: 0..65535.
    VlanId interface{}
}

func (localTrafficTag *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag) GetFilter() yfilter.YFilter { return localTrafficTag.YFilter }

func (localTrafficTag *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag) SetFilter(yf yfilter.YFilter) { localTrafficTag.YFilter = yf }

func (localTrafficTag *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag) GetGoName(yname string) string {
    if yname == "ethertype" { return "Ethertype" }
    if yname == "vlan-id" { return "VlanId" }
    return ""
}

func (localTrafficTag *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag) GetSegmentPath() string {
    return "local-traffic-tag"
}

func (localTrafficTag *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (localTrafficTag *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (localTrafficTag *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ethertype"] = localTrafficTag.Ethertype
    leafs["vlan-id"] = localTrafficTag.VlanId
    return leafs
}

func (localTrafficTag *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag) GetBundleName() string { return "cisco_ios_xr" }

func (localTrafficTag *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag) GetYangName() string { return "local-traffic-tag" }

func (localTrafficTag *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (localTrafficTag *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (localTrafficTag *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (localTrafficTag *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag) SetParent(parent types.Entity) { localTrafficTag.parent = parent }

func (localTrafficTag *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag) GetParent() types.Entity { return localTrafficTag.parent }

func (localTrafficTag *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag) GetParentYangName() string { return "local-traffic-stack" }

// Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch
// Tags to match on ingress packets
type Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Ethertype of tag to match. The type is EfpTagEtype.
    Ethertype interface{}

    // Priority to match. The type is EfpTagPriority.
    Priority interface{}

    // VLAN Ids to match. The type is slice of
    // Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange.
    VlanRange []Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange
}

func (tagsToMatch *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch) GetFilter() yfilter.YFilter { return tagsToMatch.YFilter }

func (tagsToMatch *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch) SetFilter(yf yfilter.YFilter) { tagsToMatch.YFilter = yf }

func (tagsToMatch *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch) GetGoName(yname string) string {
    if yname == "ethertype" { return "Ethertype" }
    if yname == "priority" { return "Priority" }
    if yname == "vlan-range" { return "VlanRange" }
    return ""
}

func (tagsToMatch *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch) GetSegmentPath() string {
    return "tags-to-match"
}

func (tagsToMatch *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vlan-range" {
        for _, c := range tagsToMatch.VlanRange {
            if tagsToMatch.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange{}
        tagsToMatch.VlanRange = append(tagsToMatch.VlanRange, child)
        return &tagsToMatch.VlanRange[len(tagsToMatch.VlanRange)-1]
    }
    return nil
}

func (tagsToMatch *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range tagsToMatch.VlanRange {
        children[tagsToMatch.VlanRange[i].GetSegmentPath()] = &tagsToMatch.VlanRange[i]
    }
    return children
}

func (tagsToMatch *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ethertype"] = tagsToMatch.Ethertype
    leafs["priority"] = tagsToMatch.Priority
    return leafs
}

func (tagsToMatch *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch) GetBundleName() string { return "cisco_ios_xr" }

func (tagsToMatch *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch) GetYangName() string { return "tags-to-match" }

func (tagsToMatch *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tagsToMatch *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tagsToMatch *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tagsToMatch *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch) SetParent(parent types.Entity) { tagsToMatch.parent = parent }

func (tagsToMatch *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch) GetParent() types.Entity { return tagsToMatch.parent }

func (tagsToMatch *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch) GetParentYangName() string { return "service-instance-details" }

// Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange
// VLAN Ids to match
type Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VLAN ID Low. The type is interface{} with range: 0..65535.
    VlanIdLow interface{}

    // VLAN ID High. The type is interface{} with range: 0..65535.
    VlanIdHigh interface{}
}

func (vlanRange *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange) GetFilter() yfilter.YFilter { return vlanRange.YFilter }

func (vlanRange *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange) SetFilter(yf yfilter.YFilter) { vlanRange.YFilter = yf }

func (vlanRange *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange) GetGoName(yname string) string {
    if yname == "vlan-id-low" { return "VlanIdLow" }
    if yname == "vlan-id-high" { return "VlanIdHigh" }
    return ""
}

func (vlanRange *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange) GetSegmentPath() string {
    return "vlan-range"
}

func (vlanRange *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vlanRange *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vlanRange *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vlan-id-low"] = vlanRange.VlanIdLow
    leafs["vlan-id-high"] = vlanRange.VlanIdHigh
    return leafs
}

func (vlanRange *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange) GetBundleName() string { return "cisco_ios_xr" }

func (vlanRange *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange) GetYangName() string { return "vlan-range" }

func (vlanRange *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vlanRange *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vlanRange *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vlanRange *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange) SetParent(parent types.Entity) { vlanRange.parent = parent }

func (vlanRange *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange) GetParent() types.Entity { return vlanRange.parent }

func (vlanRange *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange) GetParentYangName() string { return "tags-to-match" }

// Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_Pushe
// VLAN tags pushed on egress
type Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_Pushe struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Ethertype of tag. The type is EfpTagEtype.
    Ethertype interface{}

    // VLAN Id. The type is interface{} with range: 0..65535.
    VlanId interface{}
}

func (pushe *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_Pushe) GetFilter() yfilter.YFilter { return pushe.YFilter }

func (pushe *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_Pushe) SetFilter(yf yfilter.YFilter) { pushe.YFilter = yf }

func (pushe *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_Pushe) GetGoName(yname string) string {
    if yname == "ethertype" { return "Ethertype" }
    if yname == "vlan-id" { return "VlanId" }
    return ""
}

func (pushe *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_Pushe) GetSegmentPath() string {
    return "pushe"
}

func (pushe *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_Pushe) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pushe *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_Pushe) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pushe *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_Pushe) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ethertype"] = pushe.Ethertype
    leafs["vlan-id"] = pushe.VlanId
    return leafs
}

func (pushe *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_Pushe) GetBundleName() string { return "cisco_ios_xr" }

func (pushe *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_Pushe) GetYangName() string { return "pushe" }

func (pushe *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_Pushe) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pushe *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_Pushe) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pushe *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_Pushe) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pushe *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_Pushe) SetParent(parent types.Entity) { pushe.parent = parent }

func (pushe *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_Pushe) GetParent() types.Entity { return pushe.parent }

func (pushe *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_ServiceInstanceDetails_Pushe) GetParentYangName() string { return "service-instance-details" }

// Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_Dot1AdDot1QStack
// 802.1ad 802.1Q stack value
type Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_Dot1AdDot1QStack struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Outer tag value. The type is interface{} with range: 0..65535.
    OuterTag interface{}

    // Second tag value. The type is interface{} with range: 0..65535.
    SecondTag interface{}
}

func (dot1AdDot1QStack *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_Dot1AdDot1QStack) GetFilter() yfilter.YFilter { return dot1AdDot1QStack.YFilter }

func (dot1AdDot1QStack *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_Dot1AdDot1QStack) SetFilter(yf yfilter.YFilter) { dot1AdDot1QStack.YFilter = yf }

func (dot1AdDot1QStack *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_Dot1AdDot1QStack) GetGoName(yname string) string {
    if yname == "outer-tag" { return "OuterTag" }
    if yname == "second-tag" { return "SecondTag" }
    return ""
}

func (dot1AdDot1QStack *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_Dot1AdDot1QStack) GetSegmentPath() string {
    return "dot1ad-dot1q-stack"
}

func (dot1AdDot1QStack *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_Dot1AdDot1QStack) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dot1AdDot1QStack *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_Dot1AdDot1QStack) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dot1AdDot1QStack *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_Dot1AdDot1QStack) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["outer-tag"] = dot1AdDot1QStack.OuterTag
    leafs["second-tag"] = dot1AdDot1QStack.SecondTag
    return leafs
}

func (dot1AdDot1QStack *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_Dot1AdDot1QStack) GetBundleName() string { return "cisco_ios_xr" }

func (dot1AdDot1QStack *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_Dot1AdDot1QStack) GetYangName() string { return "dot1ad-dot1q-stack" }

func (dot1AdDot1QStack *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_Dot1AdDot1QStack) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dot1AdDot1QStack *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_Dot1AdDot1QStack) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dot1AdDot1QStack *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_Dot1AdDot1QStack) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dot1AdDot1QStack *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_Dot1AdDot1QStack) SetParent(parent types.Entity) { dot1AdDot1QStack.parent = parent }

func (dot1AdDot1QStack *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_Dot1AdDot1QStack) GetParent() types.Entity { return dot1AdDot1QStack.parent }

func (dot1AdDot1QStack *Vlan_Nodes_Node_TagAllocations_TagAllocation_EncapsulationDetails_Dot1AdDot1QStack) GetParentYangName() string { return "encapsulation-details" }

// EthernetEncapsulation
// ethernet encapsulation
type EthernetEncapsulation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Per node Ethernet encapsulation operational data.
    Nodes EthernetEncapsulation_Nodes
}

func (ethernetEncapsulation *EthernetEncapsulation) GetFilter() yfilter.YFilter { return ethernetEncapsulation.YFilter }

func (ethernetEncapsulation *EthernetEncapsulation) SetFilter(yf yfilter.YFilter) { ethernetEncapsulation.YFilter = yf }

func (ethernetEncapsulation *EthernetEncapsulation) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (ethernetEncapsulation *EthernetEncapsulation) GetSegmentPath() string {
    return "Cisco-IOS-XR-l2-eth-infra-oper:ethernet-encapsulation"
}

func (ethernetEncapsulation *EthernetEncapsulation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &ethernetEncapsulation.Nodes
    }
    return nil
}

func (ethernetEncapsulation *EthernetEncapsulation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &ethernetEncapsulation.Nodes
    return children
}

func (ethernetEncapsulation *EthernetEncapsulation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ethernetEncapsulation *EthernetEncapsulation) GetBundleName() string { return "cisco_ios_xr" }

func (ethernetEncapsulation *EthernetEncapsulation) GetYangName() string { return "ethernet-encapsulation" }

func (ethernetEncapsulation *EthernetEncapsulation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ethernetEncapsulation *EthernetEncapsulation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ethernetEncapsulation *EthernetEncapsulation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ethernetEncapsulation *EthernetEncapsulation) SetParent(parent types.Entity) { ethernetEncapsulation.parent = parent }

func (ethernetEncapsulation *EthernetEncapsulation) GetParent() types.Entity { return ethernetEncapsulation.parent }

func (ethernetEncapsulation *EthernetEncapsulation) GetParentYangName() string { return "Cisco-IOS-XR-l2-eth-infra-oper" }

// EthernetEncapsulation_Nodes
// Per node Ethernet encapsulation operational data
type EthernetEncapsulation_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The Ethernet encaps operational data for a particular node. The type is
    // slice of EthernetEncapsulation_Nodes_Node.
    Node []EthernetEncapsulation_Nodes_Node
}

func (nodes *EthernetEncapsulation_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *EthernetEncapsulation_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *EthernetEncapsulation_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *EthernetEncapsulation_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *EthernetEncapsulation_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := EthernetEncapsulation_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *EthernetEncapsulation_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *EthernetEncapsulation_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *EthernetEncapsulation_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *EthernetEncapsulation_Nodes) GetYangName() string { return "nodes" }

func (nodes *EthernetEncapsulation_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *EthernetEncapsulation_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *EthernetEncapsulation_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *EthernetEncapsulation_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *EthernetEncapsulation_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *EthernetEncapsulation_Nodes) GetParentYangName() string { return "ethernet-encapsulation" }

// EthernetEncapsulation_Nodes_Node
// The Ethernet encaps operational data for a
// particular node
type EthernetEncapsulation_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The identifier for the node. The type is string
    // with pattern: ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Unicast MAC filter table (specific to this node).
    UnicastMacFilters EthernetEncapsulation_Nodes_Node_UnicastMacFilters
}

func (node *EthernetEncapsulation_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *EthernetEncapsulation_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *EthernetEncapsulation_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "unicast-mac-filters" { return "UnicastMacFilters" }
    return ""
}

func (node *EthernetEncapsulation_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *EthernetEncapsulation_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "unicast-mac-filters" {
        return &node.UnicastMacFilters
    }
    return nil
}

func (node *EthernetEncapsulation_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["unicast-mac-filters"] = &node.UnicastMacFilters
    return children
}

func (node *EthernetEncapsulation_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    return leafs
}

func (node *EthernetEncapsulation_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *EthernetEncapsulation_Nodes_Node) GetYangName() string { return "node" }

func (node *EthernetEncapsulation_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *EthernetEncapsulation_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *EthernetEncapsulation_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *EthernetEncapsulation_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *EthernetEncapsulation_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *EthernetEncapsulation_Nodes_Node) GetParentYangName() string { return "nodes" }

// EthernetEncapsulation_Nodes_Node_UnicastMacFilters
// Unicast MAC filter table (specific to this
// node)
type EthernetEncapsulation_Nodes_Node_UnicastMacFilters struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operational data for interface with MAC filters configured. The type is
    // slice of
    // EthernetEncapsulation_Nodes_Node_UnicastMacFilters_UnicastMacFilter.
    UnicastMacFilter []EthernetEncapsulation_Nodes_Node_UnicastMacFilters_UnicastMacFilter
}

func (unicastMacFilters *EthernetEncapsulation_Nodes_Node_UnicastMacFilters) GetFilter() yfilter.YFilter { return unicastMacFilters.YFilter }

func (unicastMacFilters *EthernetEncapsulation_Nodes_Node_UnicastMacFilters) SetFilter(yf yfilter.YFilter) { unicastMacFilters.YFilter = yf }

func (unicastMacFilters *EthernetEncapsulation_Nodes_Node_UnicastMacFilters) GetGoName(yname string) string {
    if yname == "unicast-mac-filter" { return "UnicastMacFilter" }
    return ""
}

func (unicastMacFilters *EthernetEncapsulation_Nodes_Node_UnicastMacFilters) GetSegmentPath() string {
    return "unicast-mac-filters"
}

func (unicastMacFilters *EthernetEncapsulation_Nodes_Node_UnicastMacFilters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "unicast-mac-filter" {
        for _, c := range unicastMacFilters.UnicastMacFilter {
            if unicastMacFilters.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := EthernetEncapsulation_Nodes_Node_UnicastMacFilters_UnicastMacFilter{}
        unicastMacFilters.UnicastMacFilter = append(unicastMacFilters.UnicastMacFilter, child)
        return &unicastMacFilters.UnicastMacFilter[len(unicastMacFilters.UnicastMacFilter)-1]
    }
    return nil
}

func (unicastMacFilters *EthernetEncapsulation_Nodes_Node_UnicastMacFilters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range unicastMacFilters.UnicastMacFilter {
        children[unicastMacFilters.UnicastMacFilter[i].GetSegmentPath()] = &unicastMacFilters.UnicastMacFilter[i]
    }
    return children
}

func (unicastMacFilters *EthernetEncapsulation_Nodes_Node_UnicastMacFilters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (unicastMacFilters *EthernetEncapsulation_Nodes_Node_UnicastMacFilters) GetBundleName() string { return "cisco_ios_xr" }

func (unicastMacFilters *EthernetEncapsulation_Nodes_Node_UnicastMacFilters) GetYangName() string { return "unicast-mac-filters" }

func (unicastMacFilters *EthernetEncapsulation_Nodes_Node_UnicastMacFilters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (unicastMacFilters *EthernetEncapsulation_Nodes_Node_UnicastMacFilters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (unicastMacFilters *EthernetEncapsulation_Nodes_Node_UnicastMacFilters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (unicastMacFilters *EthernetEncapsulation_Nodes_Node_UnicastMacFilters) SetParent(parent types.Entity) { unicastMacFilters.parent = parent }

func (unicastMacFilters *EthernetEncapsulation_Nodes_Node_UnicastMacFilters) GetParent() types.Entity { return unicastMacFilters.parent }

func (unicastMacFilters *EthernetEncapsulation_Nodes_Node_UnicastMacFilters) GetParentYangName() string { return "node" }

// EthernetEncapsulation_Nodes_Node_UnicastMacFilters_UnicastMacFilter
// Operational data for interface with MAC
// filters configured
type EthernetEncapsulation_Nodes_Node_UnicastMacFilters_UnicastMacFilter struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The interface name. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Unicast MAC filter information. The type is slice of
    // EthernetEncapsulation_Nodes_Node_UnicastMacFilters_UnicastMacFilter_UnicastFilter.
    UnicastFilter []EthernetEncapsulation_Nodes_Node_UnicastMacFilters_UnicastMacFilter_UnicastFilter
}

func (unicastMacFilter *EthernetEncapsulation_Nodes_Node_UnicastMacFilters_UnicastMacFilter) GetFilter() yfilter.YFilter { return unicastMacFilter.YFilter }

func (unicastMacFilter *EthernetEncapsulation_Nodes_Node_UnicastMacFilters_UnicastMacFilter) SetFilter(yf yfilter.YFilter) { unicastMacFilter.YFilter = yf }

func (unicastMacFilter *EthernetEncapsulation_Nodes_Node_UnicastMacFilters_UnicastMacFilter) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "unicast-filter" { return "UnicastFilter" }
    return ""
}

func (unicastMacFilter *EthernetEncapsulation_Nodes_Node_UnicastMacFilters_UnicastMacFilter) GetSegmentPath() string {
    return "unicast-mac-filter" + "[interface-name='" + fmt.Sprintf("%v", unicastMacFilter.InterfaceName) + "']"
}

func (unicastMacFilter *EthernetEncapsulation_Nodes_Node_UnicastMacFilters_UnicastMacFilter) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "unicast-filter" {
        for _, c := range unicastMacFilter.UnicastFilter {
            if unicastMacFilter.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := EthernetEncapsulation_Nodes_Node_UnicastMacFilters_UnicastMacFilter_UnicastFilter{}
        unicastMacFilter.UnicastFilter = append(unicastMacFilter.UnicastFilter, child)
        return &unicastMacFilter.UnicastFilter[len(unicastMacFilter.UnicastFilter)-1]
    }
    return nil
}

func (unicastMacFilter *EthernetEncapsulation_Nodes_Node_UnicastMacFilters_UnicastMacFilter) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range unicastMacFilter.UnicastFilter {
        children[unicastMacFilter.UnicastFilter[i].GetSegmentPath()] = &unicastMacFilter.UnicastFilter[i]
    }
    return children
}

func (unicastMacFilter *EthernetEncapsulation_Nodes_Node_UnicastMacFilters_UnicastMacFilter) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = unicastMacFilter.InterfaceName
    return leafs
}

func (unicastMacFilter *EthernetEncapsulation_Nodes_Node_UnicastMacFilters_UnicastMacFilter) GetBundleName() string { return "cisco_ios_xr" }

func (unicastMacFilter *EthernetEncapsulation_Nodes_Node_UnicastMacFilters_UnicastMacFilter) GetYangName() string { return "unicast-mac-filter" }

func (unicastMacFilter *EthernetEncapsulation_Nodes_Node_UnicastMacFilters_UnicastMacFilter) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (unicastMacFilter *EthernetEncapsulation_Nodes_Node_UnicastMacFilters_UnicastMacFilter) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (unicastMacFilter *EthernetEncapsulation_Nodes_Node_UnicastMacFilters_UnicastMacFilter) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (unicastMacFilter *EthernetEncapsulation_Nodes_Node_UnicastMacFilters_UnicastMacFilter) SetParent(parent types.Entity) { unicastMacFilter.parent = parent }

func (unicastMacFilter *EthernetEncapsulation_Nodes_Node_UnicastMacFilters_UnicastMacFilter) GetParent() types.Entity { return unicastMacFilter.parent }

func (unicastMacFilter *EthernetEncapsulation_Nodes_Node_UnicastMacFilters_UnicastMacFilter) GetParentYangName() string { return "unicast-mac-filters" }

// EthernetEncapsulation_Nodes_Node_UnicastMacFilters_UnicastMacFilter_UnicastFilter
// Unicast MAC filter information
type EthernetEncapsulation_Nodes_Node_UnicastMacFilters_UnicastMacFilter_UnicastFilter struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}

    // Unicast MAC mode. The type is EthCapsUcastMacMode.
    Mode interface{}
}

func (unicastFilter *EthernetEncapsulation_Nodes_Node_UnicastMacFilters_UnicastMacFilter_UnicastFilter) GetFilter() yfilter.YFilter { return unicastFilter.YFilter }

func (unicastFilter *EthernetEncapsulation_Nodes_Node_UnicastMacFilters_UnicastMacFilter_UnicastFilter) SetFilter(yf yfilter.YFilter) { unicastFilter.YFilter = yf }

func (unicastFilter *EthernetEncapsulation_Nodes_Node_UnicastMacFilters_UnicastMacFilter_UnicastFilter) GetGoName(yname string) string {
    if yname == "mac-address" { return "MacAddress" }
    if yname == "mode" { return "Mode" }
    return ""
}

func (unicastFilter *EthernetEncapsulation_Nodes_Node_UnicastMacFilters_UnicastMacFilter_UnicastFilter) GetSegmentPath() string {
    return "unicast-filter"
}

func (unicastFilter *EthernetEncapsulation_Nodes_Node_UnicastMacFilters_UnicastMacFilter_UnicastFilter) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (unicastFilter *EthernetEncapsulation_Nodes_Node_UnicastMacFilters_UnicastMacFilter_UnicastFilter) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (unicastFilter *EthernetEncapsulation_Nodes_Node_UnicastMacFilters_UnicastMacFilter_UnicastFilter) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mac-address"] = unicastFilter.MacAddress
    leafs["mode"] = unicastFilter.Mode
    return leafs
}

func (unicastFilter *EthernetEncapsulation_Nodes_Node_UnicastMacFilters_UnicastMacFilter_UnicastFilter) GetBundleName() string { return "cisco_ios_xr" }

func (unicastFilter *EthernetEncapsulation_Nodes_Node_UnicastMacFilters_UnicastMacFilter_UnicastFilter) GetYangName() string { return "unicast-filter" }

func (unicastFilter *EthernetEncapsulation_Nodes_Node_UnicastMacFilters_UnicastMacFilter_UnicastFilter) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (unicastFilter *EthernetEncapsulation_Nodes_Node_UnicastMacFilters_UnicastMacFilter_UnicastFilter) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (unicastFilter *EthernetEncapsulation_Nodes_Node_UnicastMacFilters_UnicastMacFilter_UnicastFilter) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (unicastFilter *EthernetEncapsulation_Nodes_Node_UnicastMacFilters_UnicastMacFilter_UnicastFilter) SetParent(parent types.Entity) { unicastFilter.parent = parent }

func (unicastFilter *EthernetEncapsulation_Nodes_Node_UnicastMacFilters_UnicastMacFilter_UnicastFilter) GetParent() types.Entity { return unicastFilter.parent }

func (unicastFilter *EthernetEncapsulation_Nodes_Node_UnicastMacFilters_UnicastMacFilter_UnicastFilter) GetParentYangName() string { return "unicast-mac-filter" }

