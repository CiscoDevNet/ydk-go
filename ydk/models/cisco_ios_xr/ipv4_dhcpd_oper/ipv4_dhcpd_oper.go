// This module contains a collection of YANG definitions
// for Cisco IOS-XR ipv4-dhcpd package operational data.
// 
// This module contains definitions
// for the following management objects:
//   dhcp-client: DHCP client operational data
//   ipv4-dhcpd: ipv4 dhcpd
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ipv4_dhcpd_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ipv4_dhcpd_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ipv4-dhcpd-oper dhcp-client}", reflect.TypeOf(DhcpClient{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ipv4-dhcpd-oper:dhcp-client", reflect.TypeOf(DhcpClient{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ipv4-dhcpd-oper ipv4-dhcpd}", reflect.TypeOf(Ipv4Dhcpd{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ipv4-dhcpd-oper:ipv4-dhcpd", reflect.TypeOf(Ipv4Dhcpd{}))
}

// DhcpIssuVersion represents Dhcp issu version
type DhcpIssuVersion string

const (
    // Version 1
    DhcpIssuVersion_version1 DhcpIssuVersion = "version1"

    // Version 2
    DhcpIssuVersion_version2 DhcpIssuVersion = "version2"
)

// DhcpIssuRole represents Dhcp issu role
type DhcpIssuRole string

const (
    // Primary role
    DhcpIssuRole_role_primary DhcpIssuRole = "role-primary"

    // Secondary role
    DhcpIssuRole_role_secondary DhcpIssuRole = "role-secondary"
)

// ProxyLeaseLimit represents Proxy profile lease limit type
type ProxyLeaseLimit string

const (
    // Proxy lease limit type none
    ProxyLeaseLimit_none ProxyLeaseLimit = "none"

    // Proxy lease limit type interface
    ProxyLeaseLimit_interface_ ProxyLeaseLimit = "interface"

    // Proxy lease limit type circuit ID
    ProxyLeaseLimit_circuit_id ProxyLeaseLimit = "circuit-id"

    // Proxy lease limit type remote ID
    ProxyLeaseLimit_remote_id ProxyLeaseLimit = "remote-id"

    // Proxy lease limit type remote ID + circuit ID
    ProxyLeaseLimit_remote_id_circuit_id ProxyLeaseLimit = "remote-id-circuit-id"
)

// DhcpIssuPhase represents Dhcp issu phase
type DhcpIssuPhase string

const (
    // An ISSU event has not started
    DhcpIssuPhase_phase_not_started DhcpIssuPhase = "phase-not-started"

    // ISSU Load Phase
    DhcpIssuPhase_phase_load DhcpIssuPhase = "phase-load"

    // ISSU Run Phase
    DhcpIssuPhase_phase_run DhcpIssuPhase = "phase-run"

    // An ISSU event has completed successfully
    DhcpIssuPhase_phase_completed DhcpIssuPhase = "phase-completed"

    // An ISSU event has aborted
    DhcpIssuPhase_phase_aborted DhcpIssuPhase = "phase-aborted"
)

// RelayInfoPolicy represents Proxy profile relay policy
type RelayInfoPolicy string

const (
    // Relay policy replace
    RelayInfoPolicy_replace RelayInfoPolicy = "replace"

    // Relay policy keep
    RelayInfoPolicy_keep RelayInfoPolicy = "keep"

    // Relay policy drop
    RelayInfoPolicy_drop RelayInfoPolicy = "drop"

    // Relay policy encapsulate
    RelayInfoPolicy_encapsulate RelayInfoPolicy = "encapsulate"
)

// BagDhcpdIntfSrgRole represents Bag dhcpd intf srg role
type BagDhcpdIntfSrgRole string

const (
    // DHCPv4 Interface SRG role NONE
    BagDhcpdIntfSrgRole_none BagDhcpdIntfSrgRole = "none"

    // DHCPv4 Interface SRG role Master
    BagDhcpdIntfSrgRole_master BagDhcpdIntfSrgRole = "master"

    // DHCPv4 Interface SRG role Slave
    BagDhcpdIntfSrgRole_slave BagDhcpdIntfSrgRole = "slave"
)

// DhcpcIpv4State represents Dhcp Client IPv4 State
type DhcpcIpv4State string

const (
    // Init state
    DhcpcIpv4State_init DhcpcIpv4State = "init"

    // Init Reboot state
    DhcpcIpv4State_init_reboot DhcpcIpv4State = "init-reboot"

    // Rebooting state
    DhcpcIpv4State_rebooting DhcpcIpv4State = "rebooting"

    // Selecting state
    DhcpcIpv4State_selecting DhcpcIpv4State = "selecting"

    // Requesting state
    DhcpcIpv4State_requesting DhcpcIpv4State = "requesting"

    // Bound state
    DhcpcIpv4State_bound DhcpcIpv4State = "bound"

    // Renewing state
    DhcpcIpv4State_renewing DhcpcIpv4State = "renewing"

    // Rebinding state
    DhcpcIpv4State_rebinding DhcpcIpv4State = "rebinding"

    // Invalid state
    DhcpcIpv4State_invalid DhcpcIpv4State = "invalid"
)

// BagDhcpdProxyState represents Bag dhcpd proxy state
type BagDhcpdProxyState string

const (
    // Initializing
    BagDhcpdProxyState_initializing BagDhcpdProxyState = "initializing"

    // Selecting
    BagDhcpdProxyState_selecting BagDhcpdProxyState = "selecting"

    // Requesting
    BagDhcpdProxyState_requesting BagDhcpdProxyState = "requesting"

    // Bound
    BagDhcpdProxyState_bound BagDhcpdProxyState = "bound"

    // Renewing
    BagDhcpdProxyState_renewing BagDhcpdProxyState = "renewing"

    // Informing
    BagDhcpdProxyState_informing BagDhcpdProxyState = "informing"

    // Deleting
    BagDhcpdProxyState_deleting BagDhcpdProxyState = "deleting"

    // Create dpm
    BagDhcpdProxyState_create_dpm BagDhcpdProxyState = "create-dpm"

    // Offer sent
    BagDhcpdProxyState_offer_sent BagDhcpdProxyState = "offer-sent"

    // Update dpm
    BagDhcpdProxyState_update_dpm BagDhcpdProxyState = "update-dpm"

    // Route install
    BagDhcpdProxyState_route_install BagDhcpdProxyState = "route-install"

    // Disc dpm
    BagDhcpdProxyState_disc_dpm BagDhcpdProxyState = "disc-dpm"

    // Renew new intf
    BagDhcpdProxyState_renew_new_intf BagDhcpdProxyState = "renew-new-intf"

    // Other intf dpm
    BagDhcpdProxyState_other_intf_dpm BagDhcpdProxyState = "other-intf-dpm"

    // Request dpm
    BagDhcpdProxyState_request_dpm BagDhcpdProxyState = "request-dpm"

    // Change addr dpm
    BagDhcpdProxyState_change_addr_dpm BagDhcpdProxyState = "change-addr-dpm"

    // Max
    BagDhcpdProxyState_max BagDhcpdProxyState = "max"
)

// RelayInfoVpnMode represents Relay Info Vpn Mode
type RelayInfoVpnMode string

const (
    // RFC Mode
    RelayInfoVpnMode_rfc RelayInfoVpnMode = "rfc"

    // Cisco Mode
    RelayInfoVpnMode_cisco RelayInfoVpnMode = "cisco"
)

// RelayInfoAuthenticate represents Profile relay authenticate
type RelayInfoAuthenticate string

const (
    // Relay authenticate received
    RelayInfoAuthenticate_received RelayInfoAuthenticate = "received"

    // Relay authenticate inserted
    RelayInfoAuthenticate_inserted RelayInfoAuthenticate = "inserted"
)

// BroadcastFlag represents Proxy profile broadcast flag
type BroadcastFlag string

const (
    // Broadcast policy ignore
    BroadcastFlag_ignore BroadcastFlag = "ignore"

    // Broadcast policy check
    BroadcastFlag_check BroadcastFlag = "check"

    // Broadcast policy unicast always
    BroadcastFlag_unicast_always BroadcastFlag = "unicast-always"
)

// DhcpClient
// DHCP client operational data
type DhcpClient struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DHCP client list of nodes.
    Nodes DhcpClient_Nodes
}

func (dhcpClient *DhcpClient) GetFilter() yfilter.YFilter { return dhcpClient.YFilter }

func (dhcpClient *DhcpClient) SetFilter(yf yfilter.YFilter) { dhcpClient.YFilter = yf }

func (dhcpClient *DhcpClient) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (dhcpClient *DhcpClient) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv4-dhcpd-oper:dhcp-client"
}

func (dhcpClient *DhcpClient) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &dhcpClient.Nodes
    }
    return nil
}

func (dhcpClient *DhcpClient) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &dhcpClient.Nodes
    return children
}

func (dhcpClient *DhcpClient) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dhcpClient *DhcpClient) GetBundleName() string { return "cisco_ios_xr" }

func (dhcpClient *DhcpClient) GetYangName() string { return "dhcp-client" }

func (dhcpClient *DhcpClient) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dhcpClient *DhcpClient) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dhcpClient *DhcpClient) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dhcpClient *DhcpClient) SetParent(parent types.Entity) { dhcpClient.parent = parent }

func (dhcpClient *DhcpClient) GetParent() types.Entity { return dhcpClient.parent }

func (dhcpClient *DhcpClient) GetParentYangName() string { return "Cisco-IOS-XR-ipv4-dhcpd-oper" }

// DhcpClient_Nodes
// DHCP client list of nodes
type DhcpClient_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DHCP client particular node name. The type is slice of
    // DhcpClient_Nodes_Node.
    Node []DhcpClient_Nodes_Node
}

func (nodes *DhcpClient_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *DhcpClient_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *DhcpClient_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *DhcpClient_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *DhcpClient_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DhcpClient_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *DhcpClient_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *DhcpClient_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *DhcpClient_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *DhcpClient_Nodes) GetYangName() string { return "nodes" }

func (nodes *DhcpClient_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *DhcpClient_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *DhcpClient_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *DhcpClient_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *DhcpClient_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *DhcpClient_Nodes) GetParentYangName() string { return "dhcp-client" }

// DhcpClient_Nodes_Node
// DHCP client particular node name
type DhcpClient_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // IPv4 DHCP client statistics table.
    ClientStats DhcpClient_Nodes_Node_ClientStats

    // IPv4 DHCP client table.
    Clients DhcpClient_Nodes_Node_Clients
}

func (node *DhcpClient_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *DhcpClient_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *DhcpClient_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "client-stats" { return "ClientStats" }
    if yname == "clients" { return "Clients" }
    return ""
}

func (node *DhcpClient_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *DhcpClient_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "client-stats" {
        return &node.ClientStats
    }
    if childYangName == "clients" {
        return &node.Clients
    }
    return nil
}

func (node *DhcpClient_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["client-stats"] = &node.ClientStats
    children["clients"] = &node.Clients
    return children
}

func (node *DhcpClient_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    return leafs
}

func (node *DhcpClient_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *DhcpClient_Nodes_Node) GetYangName() string { return "node" }

func (node *DhcpClient_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *DhcpClient_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *DhcpClient_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *DhcpClient_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *DhcpClient_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *DhcpClient_Nodes_Node) GetParentYangName() string { return "nodes" }

// DhcpClient_Nodes_Node_ClientStats
// IPv4 DHCP client statistics table
type DhcpClient_Nodes_Node_ClientStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DHCP client binding statistics. The type is slice of
    // DhcpClient_Nodes_Node_ClientStats_ClientStat.
    ClientStat []DhcpClient_Nodes_Node_ClientStats_ClientStat
}

func (clientStats *DhcpClient_Nodes_Node_ClientStats) GetFilter() yfilter.YFilter { return clientStats.YFilter }

func (clientStats *DhcpClient_Nodes_Node_ClientStats) SetFilter(yf yfilter.YFilter) { clientStats.YFilter = yf }

func (clientStats *DhcpClient_Nodes_Node_ClientStats) GetGoName(yname string) string {
    if yname == "client-stat" { return "ClientStat" }
    return ""
}

func (clientStats *DhcpClient_Nodes_Node_ClientStats) GetSegmentPath() string {
    return "client-stats"
}

func (clientStats *DhcpClient_Nodes_Node_ClientStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "client-stat" {
        for _, c := range clientStats.ClientStat {
            if clientStats.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DhcpClient_Nodes_Node_ClientStats_ClientStat{}
        clientStats.ClientStat = append(clientStats.ClientStat, child)
        return &clientStats.ClientStat[len(clientStats.ClientStat)-1]
    }
    return nil
}

func (clientStats *DhcpClient_Nodes_Node_ClientStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range clientStats.ClientStat {
        children[clientStats.ClientStat[i].GetSegmentPath()] = &clientStats.ClientStat[i]
    }
    return children
}

func (clientStats *DhcpClient_Nodes_Node_ClientStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (clientStats *DhcpClient_Nodes_Node_ClientStats) GetBundleName() string { return "cisco_ios_xr" }

func (clientStats *DhcpClient_Nodes_Node_ClientStats) GetYangName() string { return "client-stats" }

func (clientStats *DhcpClient_Nodes_Node_ClientStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clientStats *DhcpClient_Nodes_Node_ClientStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clientStats *DhcpClient_Nodes_Node_ClientStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clientStats *DhcpClient_Nodes_Node_ClientStats) SetParent(parent types.Entity) { clientStats.parent = parent }

func (clientStats *DhcpClient_Nodes_Node_ClientStats) GetParent() types.Entity { return clientStats.parent }

func (clientStats *DhcpClient_Nodes_Node_ClientStats) GetParentYangName() string { return "node" }

// DhcpClient_Nodes_Node_ClientStats_ClientStat
// DHCP client binding statistics
type DhcpClient_Nodes_Node_ClientStats_ClientStat struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Client Ifhandle. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    ClientIfhandle interface{}

    // Dhcp Client interface name. The type is string with length: 0..65.
    InterfaceName interface{}

    // Number of events received. The type is interface{} with range:
    // 0..4294967295.
    NumEventsReceived interface{}

    // Number of create client event received. The type is interface{} with range:
    // 0..4294967295.
    NumCreateEventReceived interface{}

    // Number of delete client event received. The type is interface{} with range:
    // 0..4294967295.
    NumDeleteEventReceived interface{}

    // Number of client rebooted event received. The type is interface{} with
    // range: 0..4294967295.
    NumRebootEventReceived interface{}

    // Number of reinit client event received. The type is interface{} with range:
    // 0..4294967295.
    NumReinitEventReceived interface{}

    // Number of packet event received. The type is interface{} with range:
    // 0..4294967295.
    NumPacketEventReceived interface{}

    // Number of init timer event. The type is interface{} with range:
    // 0..4294967295.
    NumInitTimerEventi interface{}

    // Number of T1 timer event. The type is interface{} with range:
    // 0..4294967295.
    NumT1TimerEvent interface{}

    // Number of T2 timer event. The type is interface{} with range:
    // 0..4294967295.
    NumT2TimerEvent interface{}

    // Number of Lease timer event. The type is interface{} with range:
    // 0..4294967295.
    NumLeaseTimerEvent interface{}

    // Number of discovers sent successfully. The type is interface{} with range:
    // 0..4294967295.
    NumDiscoversSentSuccessfully interface{}

    // Number of requests sent successfully. The type is interface{} with range:
    // 0..4294967295.
    NumRequestsSentSuccessfully interface{}

    // Number of releases sent successfully. The type is interface{} with range:
    // 0..4294967295.
    NumReleasesSentSuccessfully interface{}

    // Number of renews sent successfully. The type is interface{} with range:
    // 0..4294967295.
    NumRenewsSentSuccessfully interface{}

    // Number of rebinds sent successfully. The type is interface{} with range:
    // 0..4294967295.
    NumRebindsSentSuccessfully interface{}

    // Number of declines sent successfully. The type is interface{} with range:
    // 0..4294967295.
    NumDeclinesSentSuccessfully interface{}

    // Number of requests sent after reboot. The type is interface{} with range:
    // 0..4294967295.
    NumRequestAfterRebootSent interface{}

    // Number of valid offers received. The type is interface{} with range:
    // 0..4294967295.
    NumValidOffersReceived interface{}

    // Number of valid acks received. The type is interface{} with range:
    // 0..4294967295.
    NumValidAcksReceived interface{}

    // Number of valid nacks received. The type is interface{} with range:
    // 0..4294967295.
    NumValidNacksReceived interface{}

    // Number of unicast packet sent successfully. The type is interface{} with
    // range: 0..4294967295.
    NumUnicastPacketSentSuccessfully interface{}

    // Number of broadcast packet sent successfully. The type is interface{} with
    // range: 0..4294967295.
    NumBroadcastPacketSentSuccess interface{}

    // Number of init timer starts. The type is interface{} with range:
    // 0..4294967295.
    NumInitTimerStart interface{}

    // Number of init timer stops. The type is interface{} with range:
    // 0..4294967295.
    NumInitTimerStop interface{}

    // Number of T1 timer starts. The type is interface{} with range:
    // 0..4294967295.
    NumT1TimerStart interface{}

    // Number of T1 timer stops. The type is interface{} with range:
    // 0..4294967295.
    NumT1TimerStop interface{}

    // Number of T2 timer starts. The type is interface{} with range:
    // 0..4294967295.
    NumT2TimerStart interface{}

    // Number of T2 timer stops. The type is interface{} with range:
    // 0..4294967295.
    NumT2TimerStop interface{}

    // Number of Lease timer starts. The type is interface{} with range:
    // 0..4294967295.
    NumLeaseTimerStart interface{}

    // Number of Lease timer stops. The type is interface{} with range:
    // 0..4294967295.
    NumLeaseTimerStop interface{}

    // Number of invalid events received. The type is interface{} with range:
    // 0..4294967295.
    NumInvalidEvents interface{}

    // Number of discover send failed. The type is interface{} with range:
    // 0..4294967295.
    NumDiscoversFailed interface{}

    // Number of request send failed. The type is interface{} with range:
    // 0..4294967295.
    NumRequestsFailed interface{}

    // Number of release send failed. The type is interface{} with range:
    // 0..4294967295.
    NumReleasesFailed interface{}

    // Number of renew send failed. The type is interface{} with range:
    // 0..4294967295.
    NumRenewsFailed interface{}

    // Number of rebind send failed. The type is interface{} with range:
    // 0..4294967295.
    NumRebindsFailed interface{}

    // Number of decline send failed. The type is interface{} with range:
    // 0..4294967295.
    NumDeclinesFailed interface{}

    // Number of requests sent after reboot failed. The type is interface{} with
    // range: 0..4294967295.
    NumRequestAfterRebootFailed interface{}

    // Number of invalid offers received. The type is interface{} with range:
    // 0..4294967295.
    NumInvalidOffers interface{}

    // Number of invalid acks received. The type is interface{} with range:
    // 0..4294967295.
    NumInvalidAcks interface{}

    // Number of invalid nacks received. The type is interface{} with range:
    // 0..4294967295.
    NumInvalidNacks interface{}

    // Number of invalid packets dropped. The type is interface{} with range:
    // 0..4294967295.
    NumInvalidPackets interface{}

    // Number of unicast packet send failed. The type is interface{} with range:
    // 0..4294967295.
    NumUnicastFailed interface{}

    // Number of broadcast packet send failed. The type is interface{} with range:
    // 0..4294967295.
    NumBroadcastFailed interface{}

    // Number of XID mismatch packets received. The type is interface{} with
    // range: 0..4294967295.
    NumXidMismatch interface{}
}

func (clientStat *DhcpClient_Nodes_Node_ClientStats_ClientStat) GetFilter() yfilter.YFilter { return clientStat.YFilter }

func (clientStat *DhcpClient_Nodes_Node_ClientStats_ClientStat) SetFilter(yf yfilter.YFilter) { clientStat.YFilter = yf }

func (clientStat *DhcpClient_Nodes_Node_ClientStats_ClientStat) GetGoName(yname string) string {
    if yname == "client-ifhandle" { return "ClientIfhandle" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "num-events-received" { return "NumEventsReceived" }
    if yname == "num-create-event-received" { return "NumCreateEventReceived" }
    if yname == "num-delete-event-received" { return "NumDeleteEventReceived" }
    if yname == "num-reboot-event-received" { return "NumRebootEventReceived" }
    if yname == "num-reinit-event-received" { return "NumReinitEventReceived" }
    if yname == "num-packet-event-received" { return "NumPacketEventReceived" }
    if yname == "num-init-timer-eventi" { return "NumInitTimerEventi" }
    if yname == "num-t1-timer-event" { return "NumT1TimerEvent" }
    if yname == "num-t2-timer-event" { return "NumT2TimerEvent" }
    if yname == "num-lease-timer-event" { return "NumLeaseTimerEvent" }
    if yname == "num-discovers-sent-successfully" { return "NumDiscoversSentSuccessfully" }
    if yname == "num-requests-sent-successfully" { return "NumRequestsSentSuccessfully" }
    if yname == "num-releases-sent-successfully" { return "NumReleasesSentSuccessfully" }
    if yname == "num-renews-sent-successfully" { return "NumRenewsSentSuccessfully" }
    if yname == "num-rebinds-sent-successfully" { return "NumRebindsSentSuccessfully" }
    if yname == "num-declines-sent-successfully" { return "NumDeclinesSentSuccessfully" }
    if yname == "num-request-after-reboot-sent" { return "NumRequestAfterRebootSent" }
    if yname == "num-valid-offers-received" { return "NumValidOffersReceived" }
    if yname == "num-valid-acks-received" { return "NumValidAcksReceived" }
    if yname == "num-valid-nacks-received" { return "NumValidNacksReceived" }
    if yname == "num-unicast-packet-sent-successfully" { return "NumUnicastPacketSentSuccessfully" }
    if yname == "num-broadcast-packet-sent-success" { return "NumBroadcastPacketSentSuccess" }
    if yname == "num-init-timer-start" { return "NumInitTimerStart" }
    if yname == "num-init-timer-stop" { return "NumInitTimerStop" }
    if yname == "num-t1-timer-start" { return "NumT1TimerStart" }
    if yname == "num-t1-timer-stop" { return "NumT1TimerStop" }
    if yname == "num-t2-timer-start" { return "NumT2TimerStart" }
    if yname == "num-t2-timer-stop" { return "NumT2TimerStop" }
    if yname == "num-lease-timer-start" { return "NumLeaseTimerStart" }
    if yname == "num-lease-timer-stop" { return "NumLeaseTimerStop" }
    if yname == "num-invalid-events" { return "NumInvalidEvents" }
    if yname == "num-discovers-failed" { return "NumDiscoversFailed" }
    if yname == "num-requests-failed" { return "NumRequestsFailed" }
    if yname == "num-releases-failed" { return "NumReleasesFailed" }
    if yname == "num-renews-failed" { return "NumRenewsFailed" }
    if yname == "num-rebinds-failed" { return "NumRebindsFailed" }
    if yname == "num-declines-failed" { return "NumDeclinesFailed" }
    if yname == "num-request-after-reboot-failed" { return "NumRequestAfterRebootFailed" }
    if yname == "num-invalid-offers" { return "NumInvalidOffers" }
    if yname == "num-invalid-acks" { return "NumInvalidAcks" }
    if yname == "num-invalid-nacks" { return "NumInvalidNacks" }
    if yname == "num-invalid-packets" { return "NumInvalidPackets" }
    if yname == "num-unicast-failed" { return "NumUnicastFailed" }
    if yname == "num-broadcast-failed" { return "NumBroadcastFailed" }
    if yname == "num-xid-mismatch" { return "NumXidMismatch" }
    return ""
}

func (clientStat *DhcpClient_Nodes_Node_ClientStats_ClientStat) GetSegmentPath() string {
    return "client-stat" + "[client-ifhandle='" + fmt.Sprintf("%v", clientStat.ClientIfhandle) + "']"
}

func (clientStat *DhcpClient_Nodes_Node_ClientStats_ClientStat) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (clientStat *DhcpClient_Nodes_Node_ClientStats_ClientStat) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (clientStat *DhcpClient_Nodes_Node_ClientStats_ClientStat) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["client-ifhandle"] = clientStat.ClientIfhandle
    leafs["interface-name"] = clientStat.InterfaceName
    leafs["num-events-received"] = clientStat.NumEventsReceived
    leafs["num-create-event-received"] = clientStat.NumCreateEventReceived
    leafs["num-delete-event-received"] = clientStat.NumDeleteEventReceived
    leafs["num-reboot-event-received"] = clientStat.NumRebootEventReceived
    leafs["num-reinit-event-received"] = clientStat.NumReinitEventReceived
    leafs["num-packet-event-received"] = clientStat.NumPacketEventReceived
    leafs["num-init-timer-eventi"] = clientStat.NumInitTimerEventi
    leafs["num-t1-timer-event"] = clientStat.NumT1TimerEvent
    leafs["num-t2-timer-event"] = clientStat.NumT2TimerEvent
    leafs["num-lease-timer-event"] = clientStat.NumLeaseTimerEvent
    leafs["num-discovers-sent-successfully"] = clientStat.NumDiscoversSentSuccessfully
    leafs["num-requests-sent-successfully"] = clientStat.NumRequestsSentSuccessfully
    leafs["num-releases-sent-successfully"] = clientStat.NumReleasesSentSuccessfully
    leafs["num-renews-sent-successfully"] = clientStat.NumRenewsSentSuccessfully
    leafs["num-rebinds-sent-successfully"] = clientStat.NumRebindsSentSuccessfully
    leafs["num-declines-sent-successfully"] = clientStat.NumDeclinesSentSuccessfully
    leafs["num-request-after-reboot-sent"] = clientStat.NumRequestAfterRebootSent
    leafs["num-valid-offers-received"] = clientStat.NumValidOffersReceived
    leafs["num-valid-acks-received"] = clientStat.NumValidAcksReceived
    leafs["num-valid-nacks-received"] = clientStat.NumValidNacksReceived
    leafs["num-unicast-packet-sent-successfully"] = clientStat.NumUnicastPacketSentSuccessfully
    leafs["num-broadcast-packet-sent-success"] = clientStat.NumBroadcastPacketSentSuccess
    leafs["num-init-timer-start"] = clientStat.NumInitTimerStart
    leafs["num-init-timer-stop"] = clientStat.NumInitTimerStop
    leafs["num-t1-timer-start"] = clientStat.NumT1TimerStart
    leafs["num-t1-timer-stop"] = clientStat.NumT1TimerStop
    leafs["num-t2-timer-start"] = clientStat.NumT2TimerStart
    leafs["num-t2-timer-stop"] = clientStat.NumT2TimerStop
    leafs["num-lease-timer-start"] = clientStat.NumLeaseTimerStart
    leafs["num-lease-timer-stop"] = clientStat.NumLeaseTimerStop
    leafs["num-invalid-events"] = clientStat.NumInvalidEvents
    leafs["num-discovers-failed"] = clientStat.NumDiscoversFailed
    leafs["num-requests-failed"] = clientStat.NumRequestsFailed
    leafs["num-releases-failed"] = clientStat.NumReleasesFailed
    leafs["num-renews-failed"] = clientStat.NumRenewsFailed
    leafs["num-rebinds-failed"] = clientStat.NumRebindsFailed
    leafs["num-declines-failed"] = clientStat.NumDeclinesFailed
    leafs["num-request-after-reboot-failed"] = clientStat.NumRequestAfterRebootFailed
    leafs["num-invalid-offers"] = clientStat.NumInvalidOffers
    leafs["num-invalid-acks"] = clientStat.NumInvalidAcks
    leafs["num-invalid-nacks"] = clientStat.NumInvalidNacks
    leafs["num-invalid-packets"] = clientStat.NumInvalidPackets
    leafs["num-unicast-failed"] = clientStat.NumUnicastFailed
    leafs["num-broadcast-failed"] = clientStat.NumBroadcastFailed
    leafs["num-xid-mismatch"] = clientStat.NumXidMismatch
    return leafs
}

func (clientStat *DhcpClient_Nodes_Node_ClientStats_ClientStat) GetBundleName() string { return "cisco_ios_xr" }

func (clientStat *DhcpClient_Nodes_Node_ClientStats_ClientStat) GetYangName() string { return "client-stat" }

func (clientStat *DhcpClient_Nodes_Node_ClientStats_ClientStat) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clientStat *DhcpClient_Nodes_Node_ClientStats_ClientStat) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clientStat *DhcpClient_Nodes_Node_ClientStats_ClientStat) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clientStat *DhcpClient_Nodes_Node_ClientStats_ClientStat) SetParent(parent types.Entity) { clientStat.parent = parent }

func (clientStat *DhcpClient_Nodes_Node_ClientStats_ClientStat) GetParent() types.Entity { return clientStat.parent }

func (clientStat *DhcpClient_Nodes_Node_ClientStats_ClientStat) GetParentYangName() string { return "client-stats" }

// DhcpClient_Nodes_Node_Clients
// IPv4 DHCP client table
type DhcpClient_Nodes_Node_Clients struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Single DHCP client binding. The type is slice of
    // DhcpClient_Nodes_Node_Clients_Client.
    Client []DhcpClient_Nodes_Node_Clients_Client
}

func (clients *DhcpClient_Nodes_Node_Clients) GetFilter() yfilter.YFilter { return clients.YFilter }

func (clients *DhcpClient_Nodes_Node_Clients) SetFilter(yf yfilter.YFilter) { clients.YFilter = yf }

func (clients *DhcpClient_Nodes_Node_Clients) GetGoName(yname string) string {
    if yname == "client" { return "Client" }
    return ""
}

func (clients *DhcpClient_Nodes_Node_Clients) GetSegmentPath() string {
    return "clients"
}

func (clients *DhcpClient_Nodes_Node_Clients) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "client" {
        for _, c := range clients.Client {
            if clients.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DhcpClient_Nodes_Node_Clients_Client{}
        clients.Client = append(clients.Client, child)
        return &clients.Client[len(clients.Client)-1]
    }
    return nil
}

func (clients *DhcpClient_Nodes_Node_Clients) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range clients.Client {
        children[clients.Client[i].GetSegmentPath()] = &clients.Client[i]
    }
    return children
}

func (clients *DhcpClient_Nodes_Node_Clients) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (clients *DhcpClient_Nodes_Node_Clients) GetBundleName() string { return "cisco_ios_xr" }

func (clients *DhcpClient_Nodes_Node_Clients) GetYangName() string { return "clients" }

func (clients *DhcpClient_Nodes_Node_Clients) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clients *DhcpClient_Nodes_Node_Clients) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clients *DhcpClient_Nodes_Node_Clients) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clients *DhcpClient_Nodes_Node_Clients) SetParent(parent types.Entity) { clients.parent = parent }

func (clients *DhcpClient_Nodes_Node_Clients) GetParent() types.Entity { return clients.parent }

func (clients *DhcpClient_Nodes_Node_Clients) GetParentYangName() string { return "node" }

// DhcpClient_Nodes_Node_Clients_Client
// Single DHCP client binding
type DhcpClient_Nodes_Node_Clients_Client struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Client Ifhandle. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    ClientIfhandle interface{}

    // Dhcp Client interface name. The type is string with length: 0..65.
    InterfaceName interface{}

    // Dhcp Client Interface MAC address. The type is string with length: 0..17.
    ClientMacAddress interface{}

    // Dhcp Client ID. The type is string with length: 0..256.
    ClientId interface{}

    // Dhcp Client State. The type is DhcpcIpv4State.
    Ipv4ClientState interface{}

    // Dhcp Client IP Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // Dhcp Client IP Address mask. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4SubnetMask interface{}

    // Dhcp Client selected server IP Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4ServerAddress interface{}

    // Dhcp Client next hop IP Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NextHopIpv4Address interface{}

    // Dhcp Client Lease time. The type is interface{} with range: 0..4294967295.
    Ipv4LeaseTime interface{}

    // Dhcp Client Renew time. The type is interface{} with range: 0..4294967295.
    Ipv4RenewTime interface{}

    // Dhcp Client Rebind time. The type is interface{} with range: 0..4294967295.
    Ipv4RebindTime interface{}

    // Dhcp Client IPV4 address configured in interface. The type is bool.
    Ipv4AddressConfigured interface{}
}

func (client *DhcpClient_Nodes_Node_Clients_Client) GetFilter() yfilter.YFilter { return client.YFilter }

func (client *DhcpClient_Nodes_Node_Clients_Client) SetFilter(yf yfilter.YFilter) { client.YFilter = yf }

func (client *DhcpClient_Nodes_Node_Clients_Client) GetGoName(yname string) string {
    if yname == "client-ifhandle" { return "ClientIfhandle" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "client-mac-address" { return "ClientMacAddress" }
    if yname == "client-id" { return "ClientId" }
    if yname == "ipv4-client-state" { return "Ipv4ClientState" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv4-subnet-mask" { return "Ipv4SubnetMask" }
    if yname == "ipv4-server-address" { return "Ipv4ServerAddress" }
    if yname == "next-hop-ipv4-address" { return "NextHopIpv4Address" }
    if yname == "ipv4-lease-time" { return "Ipv4LeaseTime" }
    if yname == "ipv4-renew-time" { return "Ipv4RenewTime" }
    if yname == "ipv4-rebind-time" { return "Ipv4RebindTime" }
    if yname == "ipv4-address-configured" { return "Ipv4AddressConfigured" }
    return ""
}

func (client *DhcpClient_Nodes_Node_Clients_Client) GetSegmentPath() string {
    return "client" + "[client-ifhandle='" + fmt.Sprintf("%v", client.ClientIfhandle) + "']"
}

func (client *DhcpClient_Nodes_Node_Clients_Client) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (client *DhcpClient_Nodes_Node_Clients_Client) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (client *DhcpClient_Nodes_Node_Clients_Client) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["client-ifhandle"] = client.ClientIfhandle
    leafs["interface-name"] = client.InterfaceName
    leafs["client-mac-address"] = client.ClientMacAddress
    leafs["client-id"] = client.ClientId
    leafs["ipv4-client-state"] = client.Ipv4ClientState
    leafs["ipv4-address"] = client.Ipv4Address
    leafs["ipv4-subnet-mask"] = client.Ipv4SubnetMask
    leafs["ipv4-server-address"] = client.Ipv4ServerAddress
    leafs["next-hop-ipv4-address"] = client.NextHopIpv4Address
    leafs["ipv4-lease-time"] = client.Ipv4LeaseTime
    leafs["ipv4-renew-time"] = client.Ipv4RenewTime
    leafs["ipv4-rebind-time"] = client.Ipv4RebindTime
    leafs["ipv4-address-configured"] = client.Ipv4AddressConfigured
    return leafs
}

func (client *DhcpClient_Nodes_Node_Clients_Client) GetBundleName() string { return "cisco_ios_xr" }

func (client *DhcpClient_Nodes_Node_Clients_Client) GetYangName() string { return "client" }

func (client *DhcpClient_Nodes_Node_Clients_Client) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (client *DhcpClient_Nodes_Node_Clients_Client) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (client *DhcpClient_Nodes_Node_Clients_Client) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (client *DhcpClient_Nodes_Node_Clients_Client) SetParent(parent types.Entity) { client.parent = parent }

func (client *DhcpClient_Nodes_Node_Clients_Client) GetParent() types.Entity { return client.parent }

func (client *DhcpClient_Nodes_Node_Clients_Client) GetParentYangName() string { return "clients" }

// Ipv4Dhcpd
// ipv4 dhcpd
type Ipv4Dhcpd struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DHCP Snoop operational data.
    Snoop Ipv4Dhcpd_Snoop

    // IPv4 DHCPD operational data for a particular location.
    Nodes Ipv4Dhcpd_Nodes
}

func (ipv4Dhcpd *Ipv4Dhcpd) GetFilter() yfilter.YFilter { return ipv4Dhcpd.YFilter }

func (ipv4Dhcpd *Ipv4Dhcpd) SetFilter(yf yfilter.YFilter) { ipv4Dhcpd.YFilter = yf }

func (ipv4Dhcpd *Ipv4Dhcpd) GetGoName(yname string) string {
    if yname == "snoop" { return "Snoop" }
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (ipv4Dhcpd *Ipv4Dhcpd) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv4-dhcpd-oper:ipv4-dhcpd"
}

func (ipv4Dhcpd *Ipv4Dhcpd) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "snoop" {
        return &ipv4Dhcpd.Snoop
    }
    if childYangName == "nodes" {
        return &ipv4Dhcpd.Nodes
    }
    return nil
}

func (ipv4Dhcpd *Ipv4Dhcpd) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["snoop"] = &ipv4Dhcpd.Snoop
    children["nodes"] = &ipv4Dhcpd.Nodes
    return children
}

func (ipv4Dhcpd *Ipv4Dhcpd) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv4Dhcpd *Ipv4Dhcpd) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4Dhcpd *Ipv4Dhcpd) GetYangName() string { return "ipv4-dhcpd" }

func (ipv4Dhcpd *Ipv4Dhcpd) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4Dhcpd *Ipv4Dhcpd) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4Dhcpd *Ipv4Dhcpd) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4Dhcpd *Ipv4Dhcpd) SetParent(parent types.Entity) { ipv4Dhcpd.parent = parent }

func (ipv4Dhcpd *Ipv4Dhcpd) GetParent() types.Entity { return ipv4Dhcpd.parent }

func (ipv4Dhcpd *Ipv4Dhcpd) GetParentYangName() string { return "Cisco-IOS-XR-ipv4-dhcpd-oper" }

// Ipv4Dhcpd_Snoop
// DHCP Snoop operational data
type Ipv4Dhcpd_Snoop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DHCP Snoop Bindings.
    Bindings Ipv4Dhcpd_Snoop_Bindings

    // DHCP snoop binding statistics.
    BindingStatistics Ipv4Dhcpd_Snoop_BindingStatistics

    // DHCP snoop statistics info.
    StatisticsInfo Ipv4Dhcpd_Snoop_StatisticsInfo

    // DHCP Snoop Profile.
    Profiles Ipv4Dhcpd_Snoop_Profiles

    // DHCP Snoop Statistics.
    Statistics Ipv4Dhcpd_Snoop_Statistics
}

func (snoop *Ipv4Dhcpd_Snoop) GetFilter() yfilter.YFilter { return snoop.YFilter }

func (snoop *Ipv4Dhcpd_Snoop) SetFilter(yf yfilter.YFilter) { snoop.YFilter = yf }

func (snoop *Ipv4Dhcpd_Snoop) GetGoName(yname string) string {
    if yname == "bindings" { return "Bindings" }
    if yname == "binding-statistics" { return "BindingStatistics" }
    if yname == "statistics-info" { return "StatisticsInfo" }
    if yname == "profiles" { return "Profiles" }
    if yname == "statistics" { return "Statistics" }
    return ""
}

func (snoop *Ipv4Dhcpd_Snoop) GetSegmentPath() string {
    return "snoop"
}

func (snoop *Ipv4Dhcpd_Snoop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bindings" {
        return &snoop.Bindings
    }
    if childYangName == "binding-statistics" {
        return &snoop.BindingStatistics
    }
    if childYangName == "statistics-info" {
        return &snoop.StatisticsInfo
    }
    if childYangName == "profiles" {
        return &snoop.Profiles
    }
    if childYangName == "statistics" {
        return &snoop.Statistics
    }
    return nil
}

func (snoop *Ipv4Dhcpd_Snoop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["bindings"] = &snoop.Bindings
    children["binding-statistics"] = &snoop.BindingStatistics
    children["statistics-info"] = &snoop.StatisticsInfo
    children["profiles"] = &snoop.Profiles
    children["statistics"] = &snoop.Statistics
    return children
}

func (snoop *Ipv4Dhcpd_Snoop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (snoop *Ipv4Dhcpd_Snoop) GetBundleName() string { return "cisco_ios_xr" }

func (snoop *Ipv4Dhcpd_Snoop) GetYangName() string { return "snoop" }

func (snoop *Ipv4Dhcpd_Snoop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (snoop *Ipv4Dhcpd_Snoop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (snoop *Ipv4Dhcpd_Snoop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (snoop *Ipv4Dhcpd_Snoop) SetParent(parent types.Entity) { snoop.parent = parent }

func (snoop *Ipv4Dhcpd_Snoop) GetParent() types.Entity { return snoop.parent }

func (snoop *Ipv4Dhcpd_Snoop) GetParentYangName() string { return "ipv4-dhcpd" }

// Ipv4Dhcpd_Snoop_Bindings
// DHCP Snoop Bindings
type Ipv4Dhcpd_Snoop_Bindings struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DHCP Snoop binding. The type is slice of Ipv4Dhcpd_Snoop_Bindings_Binding.
    Binding []Ipv4Dhcpd_Snoop_Bindings_Binding
}

func (bindings *Ipv4Dhcpd_Snoop_Bindings) GetFilter() yfilter.YFilter { return bindings.YFilter }

func (bindings *Ipv4Dhcpd_Snoop_Bindings) SetFilter(yf yfilter.YFilter) { bindings.YFilter = yf }

func (bindings *Ipv4Dhcpd_Snoop_Bindings) GetGoName(yname string) string {
    if yname == "binding" { return "Binding" }
    return ""
}

func (bindings *Ipv4Dhcpd_Snoop_Bindings) GetSegmentPath() string {
    return "bindings"
}

func (bindings *Ipv4Dhcpd_Snoop_Bindings) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "binding" {
        for _, c := range bindings.Binding {
            if bindings.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4Dhcpd_Snoop_Bindings_Binding{}
        bindings.Binding = append(bindings.Binding, child)
        return &bindings.Binding[len(bindings.Binding)-1]
    }
    return nil
}

func (bindings *Ipv4Dhcpd_Snoop_Bindings) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range bindings.Binding {
        children[bindings.Binding[i].GetSegmentPath()] = &bindings.Binding[i]
    }
    return children
}

func (bindings *Ipv4Dhcpd_Snoop_Bindings) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bindings *Ipv4Dhcpd_Snoop_Bindings) GetBundleName() string { return "cisco_ios_xr" }

func (bindings *Ipv4Dhcpd_Snoop_Bindings) GetYangName() string { return "bindings" }

func (bindings *Ipv4Dhcpd_Snoop_Bindings) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bindings *Ipv4Dhcpd_Snoop_Bindings) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bindings *Ipv4Dhcpd_Snoop_Bindings) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bindings *Ipv4Dhcpd_Snoop_Bindings) SetParent(parent types.Entity) { bindings.parent = parent }

func (bindings *Ipv4Dhcpd_Snoop_Bindings) GetParent() types.Entity { return bindings.parent }

func (bindings *Ipv4Dhcpd_Snoop_Bindings) GetParentYangName() string { return "snoop" }

// Ipv4Dhcpd_Snoop_Bindings_Binding
// DHCP Snoop binding
type Ipv4Dhcpd_Snoop_Bindings_Binding struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Client opaque handle. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    ClientUid interface{}

    // DHCP client MAC address. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    SnoopBindingChAddr interface{}

    // DHCP client MAC address length. The type is interface{} with range: 0..255.
    SnoopBindingChAddrLen interface{}

    // DHCP iaddr. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    SnoopBindingIAddr interface{}

    // DHCP client id. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    SnoopBindingClientId interface{}

    // DHCP client id len. The type is interface{} with range: 0..255.
    SnoopBindingClientIdLen interface{}

    // DHCP sm state. The type is interface{} with range: 0..255.
    SnoopBindingState interface{}

    // DHCP lease time. The type is interface{} with range: 0..4294967295.
    SnoopBindingLease interface{}

    // DHCP lease start time. The type is interface{} with range: 0..4294967295.
    SnoopBindingLeaseStartTime interface{}

    // DHCP profile name. The type is string with length: 0..65.
    SnoopBindingProfileName interface{}

    // DHCP interface to client. The type is string with length: 0..321.
    SnoopBindngInterfaceName interface{}

    // DHCP L2 bridge name. The type is string with length: 0..74.
    SnoopBindingBridgeName interface{}
}

func (binding *Ipv4Dhcpd_Snoop_Bindings_Binding) GetFilter() yfilter.YFilter { return binding.YFilter }

func (binding *Ipv4Dhcpd_Snoop_Bindings_Binding) SetFilter(yf yfilter.YFilter) { binding.YFilter = yf }

func (binding *Ipv4Dhcpd_Snoop_Bindings_Binding) GetGoName(yname string) string {
    if yname == "client-uid" { return "ClientUid" }
    if yname == "snoop-binding-ch-addr" { return "SnoopBindingChAddr" }
    if yname == "snoop-binding-ch-addr-len" { return "SnoopBindingChAddrLen" }
    if yname == "snoop-binding-i-addr" { return "SnoopBindingIAddr" }
    if yname == "snoop-binding-client-id" { return "SnoopBindingClientId" }
    if yname == "snoop-binding-client-id-len" { return "SnoopBindingClientIdLen" }
    if yname == "snoop-binding-state" { return "SnoopBindingState" }
    if yname == "snoop-binding-lease" { return "SnoopBindingLease" }
    if yname == "snoop-binding-lease-start-time" { return "SnoopBindingLeaseStartTime" }
    if yname == "snoop-binding-profile-name" { return "SnoopBindingProfileName" }
    if yname == "snoop-bindng-interface-name" { return "SnoopBindngInterfaceName" }
    if yname == "snoop-binding-bridge-name" { return "SnoopBindingBridgeName" }
    return ""
}

func (binding *Ipv4Dhcpd_Snoop_Bindings_Binding) GetSegmentPath() string {
    return "binding" + "[client-uid='" + fmt.Sprintf("%v", binding.ClientUid) + "']"
}

func (binding *Ipv4Dhcpd_Snoop_Bindings_Binding) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (binding *Ipv4Dhcpd_Snoop_Bindings_Binding) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (binding *Ipv4Dhcpd_Snoop_Bindings_Binding) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["client-uid"] = binding.ClientUid
    leafs["snoop-binding-ch-addr"] = binding.SnoopBindingChAddr
    leafs["snoop-binding-ch-addr-len"] = binding.SnoopBindingChAddrLen
    leafs["snoop-binding-i-addr"] = binding.SnoopBindingIAddr
    leafs["snoop-binding-client-id"] = binding.SnoopBindingClientId
    leafs["snoop-binding-client-id-len"] = binding.SnoopBindingClientIdLen
    leafs["snoop-binding-state"] = binding.SnoopBindingState
    leafs["snoop-binding-lease"] = binding.SnoopBindingLease
    leafs["snoop-binding-lease-start-time"] = binding.SnoopBindingLeaseStartTime
    leafs["snoop-binding-profile-name"] = binding.SnoopBindingProfileName
    leafs["snoop-bindng-interface-name"] = binding.SnoopBindngInterfaceName
    leafs["snoop-binding-bridge-name"] = binding.SnoopBindingBridgeName
    return leafs
}

func (binding *Ipv4Dhcpd_Snoop_Bindings_Binding) GetBundleName() string { return "cisco_ios_xr" }

func (binding *Ipv4Dhcpd_Snoop_Bindings_Binding) GetYangName() string { return "binding" }

func (binding *Ipv4Dhcpd_Snoop_Bindings_Binding) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (binding *Ipv4Dhcpd_Snoop_Bindings_Binding) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (binding *Ipv4Dhcpd_Snoop_Bindings_Binding) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (binding *Ipv4Dhcpd_Snoop_Bindings_Binding) SetParent(parent types.Entity) { binding.parent = parent }

func (binding *Ipv4Dhcpd_Snoop_Bindings_Binding) GetParent() types.Entity { return binding.parent }

func (binding *Ipv4Dhcpd_Snoop_Bindings_Binding) GetParentYangName() string { return "bindings" }

// Ipv4Dhcpd_Snoop_BindingStatistics
// DHCP snoop binding statistics
type Ipv4Dhcpd_Snoop_BindingStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Total number of snoop bindings. The type is interface{} with range:
    // 0..4294967295.
    SnoopBindingTotal interface{}

    // Snoop binding timestamp. The type is interface{} with range: 0..4294967295.
    SnoopBindingTimestamp interface{}
}

func (bindingStatistics *Ipv4Dhcpd_Snoop_BindingStatistics) GetFilter() yfilter.YFilter { return bindingStatistics.YFilter }

func (bindingStatistics *Ipv4Dhcpd_Snoop_BindingStatistics) SetFilter(yf yfilter.YFilter) { bindingStatistics.YFilter = yf }

func (bindingStatistics *Ipv4Dhcpd_Snoop_BindingStatistics) GetGoName(yname string) string {
    if yname == "snoop-binding-total" { return "SnoopBindingTotal" }
    if yname == "snoop-binding-timestamp" { return "SnoopBindingTimestamp" }
    return ""
}

func (bindingStatistics *Ipv4Dhcpd_Snoop_BindingStatistics) GetSegmentPath() string {
    return "binding-statistics"
}

func (bindingStatistics *Ipv4Dhcpd_Snoop_BindingStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bindingStatistics *Ipv4Dhcpd_Snoop_BindingStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bindingStatistics *Ipv4Dhcpd_Snoop_BindingStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["snoop-binding-total"] = bindingStatistics.SnoopBindingTotal
    leafs["snoop-binding-timestamp"] = bindingStatistics.SnoopBindingTimestamp
    return leafs
}

func (bindingStatistics *Ipv4Dhcpd_Snoop_BindingStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (bindingStatistics *Ipv4Dhcpd_Snoop_BindingStatistics) GetYangName() string { return "binding-statistics" }

func (bindingStatistics *Ipv4Dhcpd_Snoop_BindingStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bindingStatistics *Ipv4Dhcpd_Snoop_BindingStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bindingStatistics *Ipv4Dhcpd_Snoop_BindingStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bindingStatistics *Ipv4Dhcpd_Snoop_BindingStatistics) SetParent(parent types.Entity) { bindingStatistics.parent = parent }

func (bindingStatistics *Ipv4Dhcpd_Snoop_BindingStatistics) GetParent() types.Entity { return bindingStatistics.parent }

func (bindingStatistics *Ipv4Dhcpd_Snoop_BindingStatistics) GetParentYangName() string { return "snoop" }

// Ipv4Dhcpd_Snoop_StatisticsInfo
// DHCP snoop statistics info
type Ipv4Dhcpd_Snoop_StatisticsInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Snoop Stats timestamp. The type is interface{} with range: 0..4294967295.
    SnoopStatsTimestamp interface{}
}

func (statisticsInfo *Ipv4Dhcpd_Snoop_StatisticsInfo) GetFilter() yfilter.YFilter { return statisticsInfo.YFilter }

func (statisticsInfo *Ipv4Dhcpd_Snoop_StatisticsInfo) SetFilter(yf yfilter.YFilter) { statisticsInfo.YFilter = yf }

func (statisticsInfo *Ipv4Dhcpd_Snoop_StatisticsInfo) GetGoName(yname string) string {
    if yname == "snoop-stats-timestamp" { return "SnoopStatsTimestamp" }
    return ""
}

func (statisticsInfo *Ipv4Dhcpd_Snoop_StatisticsInfo) GetSegmentPath() string {
    return "statistics-info"
}

func (statisticsInfo *Ipv4Dhcpd_Snoop_StatisticsInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (statisticsInfo *Ipv4Dhcpd_Snoop_StatisticsInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (statisticsInfo *Ipv4Dhcpd_Snoop_StatisticsInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["snoop-stats-timestamp"] = statisticsInfo.SnoopStatsTimestamp
    return leafs
}

func (statisticsInfo *Ipv4Dhcpd_Snoop_StatisticsInfo) GetBundleName() string { return "cisco_ios_xr" }

func (statisticsInfo *Ipv4Dhcpd_Snoop_StatisticsInfo) GetYangName() string { return "statistics-info" }

func (statisticsInfo *Ipv4Dhcpd_Snoop_StatisticsInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statisticsInfo *Ipv4Dhcpd_Snoop_StatisticsInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statisticsInfo *Ipv4Dhcpd_Snoop_StatisticsInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statisticsInfo *Ipv4Dhcpd_Snoop_StatisticsInfo) SetParent(parent types.Entity) { statisticsInfo.parent = parent }

func (statisticsInfo *Ipv4Dhcpd_Snoop_StatisticsInfo) GetParent() types.Entity { return statisticsInfo.parent }

func (statisticsInfo *Ipv4Dhcpd_Snoop_StatisticsInfo) GetParentYangName() string { return "snoop" }

// Ipv4Dhcpd_Snoop_Profiles
// DHCP Snoop Profile
type Ipv4Dhcpd_Snoop_Profiles struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DHCP Snoop profile. The type is slice of Ipv4Dhcpd_Snoop_Profiles_Profile.
    Profile []Ipv4Dhcpd_Snoop_Profiles_Profile
}

func (profiles *Ipv4Dhcpd_Snoop_Profiles) GetFilter() yfilter.YFilter { return profiles.YFilter }

func (profiles *Ipv4Dhcpd_Snoop_Profiles) SetFilter(yf yfilter.YFilter) { profiles.YFilter = yf }

func (profiles *Ipv4Dhcpd_Snoop_Profiles) GetGoName(yname string) string {
    if yname == "profile" { return "Profile" }
    return ""
}

func (profiles *Ipv4Dhcpd_Snoop_Profiles) GetSegmentPath() string {
    return "profiles"
}

func (profiles *Ipv4Dhcpd_Snoop_Profiles) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "profile" {
        for _, c := range profiles.Profile {
            if profiles.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4Dhcpd_Snoop_Profiles_Profile{}
        profiles.Profile = append(profiles.Profile, child)
        return &profiles.Profile[len(profiles.Profile)-1]
    }
    return nil
}

func (profiles *Ipv4Dhcpd_Snoop_Profiles) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range profiles.Profile {
        children[profiles.Profile[i].GetSegmentPath()] = &profiles.Profile[i]
    }
    return children
}

func (profiles *Ipv4Dhcpd_Snoop_Profiles) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (profiles *Ipv4Dhcpd_Snoop_Profiles) GetBundleName() string { return "cisco_ios_xr" }

func (profiles *Ipv4Dhcpd_Snoop_Profiles) GetYangName() string { return "profiles" }

func (profiles *Ipv4Dhcpd_Snoop_Profiles) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (profiles *Ipv4Dhcpd_Snoop_Profiles) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (profiles *Ipv4Dhcpd_Snoop_Profiles) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (profiles *Ipv4Dhcpd_Snoop_Profiles) SetParent(parent types.Entity) { profiles.parent = parent }

func (profiles *Ipv4Dhcpd_Snoop_Profiles) GetParent() types.Entity { return profiles.parent }

func (profiles *Ipv4Dhcpd_Snoop_Profiles) GetParentYangName() string { return "snoop" }

// Ipv4Dhcpd_Snoop_Profiles_Profile
// DHCP Snoop profile
type Ipv4Dhcpd_Snoop_Profiles_Profile struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Profile name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    ProfileName interface{}

    // Profile Name. The type is string with length: 0..65.
    SnoopProfileName interface{}

    // Profile unique ID. The type is interface{} with range: 0..4294967295.
    SnoopProfileUid interface{}

    // Relay info option. The type is interface{} with range: 0..255.
    SnoopProfileRelayInfoOption interface{}

    // Allow untrusted relay info. The type is interface{} with range: 0..255.
    SnoopProfileRelayInfoAllowUntrusted interface{}

    // Relay info policy. The type is interface{} with range: 0..255.
    SnoopProfileRelayInfoPolicy interface{}

    // Trust. The type is interface{} with range: 0..255.
    SnoopProfileTrusted interface{}
}

func (profile *Ipv4Dhcpd_Snoop_Profiles_Profile) GetFilter() yfilter.YFilter { return profile.YFilter }

func (profile *Ipv4Dhcpd_Snoop_Profiles_Profile) SetFilter(yf yfilter.YFilter) { profile.YFilter = yf }

func (profile *Ipv4Dhcpd_Snoop_Profiles_Profile) GetGoName(yname string) string {
    if yname == "profile-name" { return "ProfileName" }
    if yname == "snoop-profile-name" { return "SnoopProfileName" }
    if yname == "snoop-profile-uid" { return "SnoopProfileUid" }
    if yname == "snoop-profile-relay-info-option" { return "SnoopProfileRelayInfoOption" }
    if yname == "snoop-profile-relay-info-allow-untrusted" { return "SnoopProfileRelayInfoAllowUntrusted" }
    if yname == "snoop-profile-relay-info-policy" { return "SnoopProfileRelayInfoPolicy" }
    if yname == "snoop-profile-trusted" { return "SnoopProfileTrusted" }
    return ""
}

func (profile *Ipv4Dhcpd_Snoop_Profiles_Profile) GetSegmentPath() string {
    return "profile" + "[profile-name='" + fmt.Sprintf("%v", profile.ProfileName) + "']"
}

func (profile *Ipv4Dhcpd_Snoop_Profiles_Profile) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (profile *Ipv4Dhcpd_Snoop_Profiles_Profile) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (profile *Ipv4Dhcpd_Snoop_Profiles_Profile) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["profile-name"] = profile.ProfileName
    leafs["snoop-profile-name"] = profile.SnoopProfileName
    leafs["snoop-profile-uid"] = profile.SnoopProfileUid
    leafs["snoop-profile-relay-info-option"] = profile.SnoopProfileRelayInfoOption
    leafs["snoop-profile-relay-info-allow-untrusted"] = profile.SnoopProfileRelayInfoAllowUntrusted
    leafs["snoop-profile-relay-info-policy"] = profile.SnoopProfileRelayInfoPolicy
    leafs["snoop-profile-trusted"] = profile.SnoopProfileTrusted
    return leafs
}

func (profile *Ipv4Dhcpd_Snoop_Profiles_Profile) GetBundleName() string { return "cisco_ios_xr" }

func (profile *Ipv4Dhcpd_Snoop_Profiles_Profile) GetYangName() string { return "profile" }

func (profile *Ipv4Dhcpd_Snoop_Profiles_Profile) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (profile *Ipv4Dhcpd_Snoop_Profiles_Profile) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (profile *Ipv4Dhcpd_Snoop_Profiles_Profile) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (profile *Ipv4Dhcpd_Snoop_Profiles_Profile) SetParent(parent types.Entity) { profile.parent = parent }

func (profile *Ipv4Dhcpd_Snoop_Profiles_Profile) GetParent() types.Entity { return profile.parent }

func (profile *Ipv4Dhcpd_Snoop_Profiles_Profile) GetParentYangName() string { return "profiles" }

// Ipv4Dhcpd_Snoop_Statistics
// DHCP Snoop Statistics
type Ipv4Dhcpd_Snoop_Statistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DHCP Snoop bridge domain statistics. The type is slice of
    // Ipv4Dhcpd_Snoop_Statistics_Statistic.
    Statistic []Ipv4Dhcpd_Snoop_Statistics_Statistic
}

func (statistics *Ipv4Dhcpd_Snoop_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *Ipv4Dhcpd_Snoop_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *Ipv4Dhcpd_Snoop_Statistics) GetGoName(yname string) string {
    if yname == "statistic" { return "Statistic" }
    return ""
}

func (statistics *Ipv4Dhcpd_Snoop_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *Ipv4Dhcpd_Snoop_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "statistic" {
        for _, c := range statistics.Statistic {
            if statistics.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4Dhcpd_Snoop_Statistics_Statistic{}
        statistics.Statistic = append(statistics.Statistic, child)
        return &statistics.Statistic[len(statistics.Statistic)-1]
    }
    return nil
}

func (statistics *Ipv4Dhcpd_Snoop_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range statistics.Statistic {
        children[statistics.Statistic[i].GetSegmentPath()] = &statistics.Statistic[i]
    }
    return children
}

func (statistics *Ipv4Dhcpd_Snoop_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (statistics *Ipv4Dhcpd_Snoop_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *Ipv4Dhcpd_Snoop_Statistics) GetYangName() string { return "statistics" }

func (statistics *Ipv4Dhcpd_Snoop_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *Ipv4Dhcpd_Snoop_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *Ipv4Dhcpd_Snoop_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *Ipv4Dhcpd_Snoop_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *Ipv4Dhcpd_Snoop_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *Ipv4Dhcpd_Snoop_Statistics) GetParentYangName() string { return "snoop" }

// Ipv4Dhcpd_Snoop_Statistics_Statistic
// DHCP Snoop bridge domain statistics
type Ipv4Dhcpd_Snoop_Statistics_Statistic struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Bridge domain name. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    BridgeName interface{}

    // DHCP L2 bridge name. The type is string with length: 0..74.
    SnoopStatisticsBridgeName interface{}

    // Public snoop statistics. The type is slice of interface{} with range:
    // 0..18446744073709551615.
    SnoopStatistic []interface{}
}

func (statistic *Ipv4Dhcpd_Snoop_Statistics_Statistic) GetFilter() yfilter.YFilter { return statistic.YFilter }

func (statistic *Ipv4Dhcpd_Snoop_Statistics_Statistic) SetFilter(yf yfilter.YFilter) { statistic.YFilter = yf }

func (statistic *Ipv4Dhcpd_Snoop_Statistics_Statistic) GetGoName(yname string) string {
    if yname == "bridge-name" { return "BridgeName" }
    if yname == "snoop-statistics-bridge-name" { return "SnoopStatisticsBridgeName" }
    if yname == "snoop-statistic" { return "SnoopStatistic" }
    return ""
}

func (statistic *Ipv4Dhcpd_Snoop_Statistics_Statistic) GetSegmentPath() string {
    return "statistic" + "[bridge-name='" + fmt.Sprintf("%v", statistic.BridgeName) + "']"
}

func (statistic *Ipv4Dhcpd_Snoop_Statistics_Statistic) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (statistic *Ipv4Dhcpd_Snoop_Statistics_Statistic) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (statistic *Ipv4Dhcpd_Snoop_Statistics_Statistic) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["bridge-name"] = statistic.BridgeName
    leafs["snoop-statistics-bridge-name"] = statistic.SnoopStatisticsBridgeName
    leafs["snoop-statistic"] = statistic.SnoopStatistic
    return leafs
}

func (statistic *Ipv4Dhcpd_Snoop_Statistics_Statistic) GetBundleName() string { return "cisco_ios_xr" }

func (statistic *Ipv4Dhcpd_Snoop_Statistics_Statistic) GetYangName() string { return "statistic" }

func (statistic *Ipv4Dhcpd_Snoop_Statistics_Statistic) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistic *Ipv4Dhcpd_Snoop_Statistics_Statistic) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistic *Ipv4Dhcpd_Snoop_Statistics_Statistic) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistic *Ipv4Dhcpd_Snoop_Statistics_Statistic) SetParent(parent types.Entity) { statistic.parent = parent }

func (statistic *Ipv4Dhcpd_Snoop_Statistics_Statistic) GetParent() types.Entity { return statistic.parent }

func (statistic *Ipv4Dhcpd_Snoop_Statistics_Statistic) GetParentYangName() string { return "statistics" }

// Ipv4Dhcpd_Nodes
// IPv4 DHCPD operational data for a particular
// location
type Ipv4Dhcpd_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Location. For eg., 0/1/CPU0. The type is slice of Ipv4Dhcpd_Nodes_Node.
    Node []Ipv4Dhcpd_Nodes_Node
}

func (nodes *Ipv4Dhcpd_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *Ipv4Dhcpd_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *Ipv4Dhcpd_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *Ipv4Dhcpd_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *Ipv4Dhcpd_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4Dhcpd_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *Ipv4Dhcpd_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *Ipv4Dhcpd_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *Ipv4Dhcpd_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *Ipv4Dhcpd_Nodes) GetYangName() string { return "nodes" }

func (nodes *Ipv4Dhcpd_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *Ipv4Dhcpd_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *Ipv4Dhcpd_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *Ipv4Dhcpd_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *Ipv4Dhcpd_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *Ipv4Dhcpd_Nodes) GetParentYangName() string { return "ipv4-dhcpd" }

// Ipv4Dhcpd_Nodes_Node
// Location. For eg., 0/1/CPU0
type Ipv4Dhcpd_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The node id to filter on. For eg., 0/1/CPU0. The
    // type is string with pattern: ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Nodeid interface{}

    // IPv4 DHCP proxy operational data.
    Proxy Ipv4Dhcpd_Nodes_Node_Proxy

    // IPv4 DHCP proxy/server Interface.
    Interfaces Ipv4Dhcpd_Nodes_Node_Interfaces

    // IPv4 DHCP base operational data.
    Base Ipv4Dhcpd_Nodes_Node_Base

    // IPv4 DHCP Server operational data.
    Server Ipv4Dhcpd_Nodes_Node_Server

    // IPv4 DHCPD Relay operational data.
    Relay Ipv4Dhcpd_Nodes_Node_Relay
}

func (node *Ipv4Dhcpd_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *Ipv4Dhcpd_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *Ipv4Dhcpd_Nodes_Node) GetGoName(yname string) string {
    if yname == "nodeid" { return "Nodeid" }
    if yname == "proxy" { return "Proxy" }
    if yname == "interfaces" { return "Interfaces" }
    if yname == "base" { return "Base" }
    if yname == "server" { return "Server" }
    if yname == "relay" { return "Relay" }
    return ""
}

func (node *Ipv4Dhcpd_Nodes_Node) GetSegmentPath() string {
    return "node" + "[nodeid='" + fmt.Sprintf("%v", node.Nodeid) + "']"
}

func (node *Ipv4Dhcpd_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "proxy" {
        return &node.Proxy
    }
    if childYangName == "interfaces" {
        return &node.Interfaces
    }
    if childYangName == "base" {
        return &node.Base
    }
    if childYangName == "server" {
        return &node.Server
    }
    if childYangName == "relay" {
        return &node.Relay
    }
    return nil
}

func (node *Ipv4Dhcpd_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["proxy"] = &node.Proxy
    children["interfaces"] = &node.Interfaces
    children["base"] = &node.Base
    children["server"] = &node.Server
    children["relay"] = &node.Relay
    return children
}

func (node *Ipv4Dhcpd_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["nodeid"] = node.Nodeid
    return leafs
}

func (node *Ipv4Dhcpd_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *Ipv4Dhcpd_Nodes_Node) GetYangName() string { return "node" }

func (node *Ipv4Dhcpd_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *Ipv4Dhcpd_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *Ipv4Dhcpd_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *Ipv4Dhcpd_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *Ipv4Dhcpd_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *Ipv4Dhcpd_Nodes_Node) GetParentYangName() string { return "nodes" }

// Ipv4Dhcpd_Nodes_Node_Proxy
// IPv4 DHCP proxy operational data
type Ipv4Dhcpd_Nodes_Node_Proxy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DHCP proxy stats info.
    StatisticsInfo Ipv4Dhcpd_Nodes_Node_Proxy_StatisticsInfo

    // DHCP proxy list of VRF names.
    Vrfs Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs

    // IPv4 DHCP proxy profile.
    Profiles Ipv4Dhcpd_Nodes_Node_Proxy_Profiles

    // DHCP proxy statistics.
    Statistics Ipv4Dhcpd_Nodes_Node_Proxy_Statistics

    // DHCP proxy bindings.
    Binding Ipv4Dhcpd_Nodes_Node_Proxy_Binding
}

func (proxy *Ipv4Dhcpd_Nodes_Node_Proxy) GetFilter() yfilter.YFilter { return proxy.YFilter }

func (proxy *Ipv4Dhcpd_Nodes_Node_Proxy) SetFilter(yf yfilter.YFilter) { proxy.YFilter = yf }

func (proxy *Ipv4Dhcpd_Nodes_Node_Proxy) GetGoName(yname string) string {
    if yname == "statistics-info" { return "StatisticsInfo" }
    if yname == "vrfs" { return "Vrfs" }
    if yname == "profiles" { return "Profiles" }
    if yname == "statistics" { return "Statistics" }
    if yname == "binding" { return "Binding" }
    return ""
}

func (proxy *Ipv4Dhcpd_Nodes_Node_Proxy) GetSegmentPath() string {
    return "proxy"
}

func (proxy *Ipv4Dhcpd_Nodes_Node_Proxy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "statistics-info" {
        return &proxy.StatisticsInfo
    }
    if childYangName == "vrfs" {
        return &proxy.Vrfs
    }
    if childYangName == "profiles" {
        return &proxy.Profiles
    }
    if childYangName == "statistics" {
        return &proxy.Statistics
    }
    if childYangName == "binding" {
        return &proxy.Binding
    }
    return nil
}

func (proxy *Ipv4Dhcpd_Nodes_Node_Proxy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["statistics-info"] = &proxy.StatisticsInfo
    children["vrfs"] = &proxy.Vrfs
    children["profiles"] = &proxy.Profiles
    children["statistics"] = &proxy.Statistics
    children["binding"] = &proxy.Binding
    return children
}

func (proxy *Ipv4Dhcpd_Nodes_Node_Proxy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (proxy *Ipv4Dhcpd_Nodes_Node_Proxy) GetBundleName() string { return "cisco_ios_xr" }

func (proxy *Ipv4Dhcpd_Nodes_Node_Proxy) GetYangName() string { return "proxy" }

func (proxy *Ipv4Dhcpd_Nodes_Node_Proxy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (proxy *Ipv4Dhcpd_Nodes_Node_Proxy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (proxy *Ipv4Dhcpd_Nodes_Node_Proxy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (proxy *Ipv4Dhcpd_Nodes_Node_Proxy) SetParent(parent types.Entity) { proxy.parent = parent }

func (proxy *Ipv4Dhcpd_Nodes_Node_Proxy) GetParent() types.Entity { return proxy.parent }

func (proxy *Ipv4Dhcpd_Nodes_Node_Proxy) GetParentYangName() string { return "node" }

// Ipv4Dhcpd_Nodes_Node_Proxy_StatisticsInfo
// DHCP proxy stats info
type Ipv4Dhcpd_Nodes_Node_Proxy_StatisticsInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Proxy Stats timestamp. The type is interface{} with range: 0..4294967295.
    ProxyStatsTimestamp interface{}
}

func (statisticsInfo *Ipv4Dhcpd_Nodes_Node_Proxy_StatisticsInfo) GetFilter() yfilter.YFilter { return statisticsInfo.YFilter }

func (statisticsInfo *Ipv4Dhcpd_Nodes_Node_Proxy_StatisticsInfo) SetFilter(yf yfilter.YFilter) { statisticsInfo.YFilter = yf }

func (statisticsInfo *Ipv4Dhcpd_Nodes_Node_Proxy_StatisticsInfo) GetGoName(yname string) string {
    if yname == "proxy-stats-timestamp" { return "ProxyStatsTimestamp" }
    return ""
}

func (statisticsInfo *Ipv4Dhcpd_Nodes_Node_Proxy_StatisticsInfo) GetSegmentPath() string {
    return "statistics-info"
}

func (statisticsInfo *Ipv4Dhcpd_Nodes_Node_Proxy_StatisticsInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (statisticsInfo *Ipv4Dhcpd_Nodes_Node_Proxy_StatisticsInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (statisticsInfo *Ipv4Dhcpd_Nodes_Node_Proxy_StatisticsInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["proxy-stats-timestamp"] = statisticsInfo.ProxyStatsTimestamp
    return leafs
}

func (statisticsInfo *Ipv4Dhcpd_Nodes_Node_Proxy_StatisticsInfo) GetBundleName() string { return "cisco_ios_xr" }

func (statisticsInfo *Ipv4Dhcpd_Nodes_Node_Proxy_StatisticsInfo) GetYangName() string { return "statistics-info" }

func (statisticsInfo *Ipv4Dhcpd_Nodes_Node_Proxy_StatisticsInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statisticsInfo *Ipv4Dhcpd_Nodes_Node_Proxy_StatisticsInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statisticsInfo *Ipv4Dhcpd_Nodes_Node_Proxy_StatisticsInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statisticsInfo *Ipv4Dhcpd_Nodes_Node_Proxy_StatisticsInfo) SetParent(parent types.Entity) { statisticsInfo.parent = parent }

func (statisticsInfo *Ipv4Dhcpd_Nodes_Node_Proxy_StatisticsInfo) GetParent() types.Entity { return statisticsInfo.parent }

func (statisticsInfo *Ipv4Dhcpd_Nodes_Node_Proxy_StatisticsInfo) GetParentYangName() string { return "proxy" }

// Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs
// DHCP proxy list of VRF names
type Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 DHCP proxy VRF name. The type is slice of
    // Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf.
    Vrf []Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf
}

func (vrfs *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs) GetFilter() yfilter.YFilter { return vrfs.YFilter }

func (vrfs *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs) SetFilter(yf yfilter.YFilter) { vrfs.YFilter = yf }

func (vrfs *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs) GetGoName(yname string) string {
    if yname == "vrf" { return "Vrf" }
    return ""
}

func (vrfs *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs) GetSegmentPath() string {
    return "vrfs"
}

func (vrfs *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf" {
        for _, c := range vrfs.Vrf {
            if vrfs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf{}
        vrfs.Vrf = append(vrfs.Vrf, child)
        return &vrfs.Vrf[len(vrfs.Vrf)-1]
    }
    return nil
}

func (vrfs *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vrfs.Vrf {
        children[vrfs.Vrf[i].GetSegmentPath()] = &vrfs.Vrf[i]
    }
    return children
}

func (vrfs *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vrfs *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs) GetBundleName() string { return "cisco_ios_xr" }

func (vrfs *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs) GetYangName() string { return "vrfs" }

func (vrfs *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfs *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfs *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfs *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs) SetParent(parent types.Entity) { vrfs.parent = parent }

func (vrfs *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs) GetParent() types.Entity { return vrfs.parent }

func (vrfs *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs) GetParentYangName() string { return "proxy" }

// Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf
// IPv4 DHCP proxy VRF name
type Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // IPv4 DHCP proxy statistics.
    Statistics Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics
}

func (vrf *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf) GetFilter() yfilter.YFilter { return vrf.YFilter }

func (vrf *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf) SetFilter(yf yfilter.YFilter) { vrf.YFilter = yf }

func (vrf *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "statistics" { return "Statistics" }
    return ""
}

func (vrf *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf) GetSegmentPath() string {
    return "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
}

func (vrf *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "statistics" {
        return &vrf.Statistics
    }
    return nil
}

func (vrf *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["statistics"] = &vrf.Statistics
    return children
}

func (vrf *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = vrf.VrfName
    return leafs
}

func (vrf *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf) GetBundleName() string { return "cisco_ios_xr" }

func (vrf *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf) GetYangName() string { return "vrf" }

func (vrf *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrf *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrf *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrf *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf) SetParent(parent types.Entity) { vrf.parent = parent }

func (vrf *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf) GetParent() types.Entity { return vrf.parent }

func (vrf *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf) GetParentYangName() string { return "vrfs" }

// Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics
// IPv4 DHCP proxy statistics
type Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DHCP discover packets.
    Discover Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Discover

    // DHCP offer packets.
    Offer Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Offer

    // DHCP request packets.
    Request Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Request

    // DHCP decline packets.
    Decline Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Decline

    // DHCP ack packets.
    Ack Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Ack

    // DHCP nak packets.
    Nak Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Nak

    // DHCP release packets.
    Release Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Release

    // DHCP inform packets.
    Inform Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Inform

    // DHCP lease query packets.
    LeaseQuery Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQuery

    // DHCP lease not assigned.
    LeaseNotAssigned Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseNotAssigned

    // DHCP lease unknown.
    LeaseUnknown Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseUnknown

    // DHCP lease active.
    LeaseActive Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseActive

    // DHCP BOOTP request.
    BootpRequest Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_BootpRequest

    // DHCP BOOTP reply.
    BootpReply Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_BootpReply
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics) GetGoName(yname string) string {
    if yname == "discover" { return "Discover" }
    if yname == "offer" { return "Offer" }
    if yname == "request" { return "Request" }
    if yname == "decline" { return "Decline" }
    if yname == "ack" { return "Ack" }
    if yname == "nak" { return "Nak" }
    if yname == "release" { return "Release" }
    if yname == "inform" { return "Inform" }
    if yname == "lease-query" { return "LeaseQuery" }
    if yname == "lease-not-assigned" { return "LeaseNotAssigned" }
    if yname == "lease-unknown" { return "LeaseUnknown" }
    if yname == "lease-active" { return "LeaseActive" }
    if yname == "bootp-request" { return "BootpRequest" }
    if yname == "bootp-reply" { return "BootpReply" }
    return ""
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "discover" {
        return &statistics.Discover
    }
    if childYangName == "offer" {
        return &statistics.Offer
    }
    if childYangName == "request" {
        return &statistics.Request
    }
    if childYangName == "decline" {
        return &statistics.Decline
    }
    if childYangName == "ack" {
        return &statistics.Ack
    }
    if childYangName == "nak" {
        return &statistics.Nak
    }
    if childYangName == "release" {
        return &statistics.Release
    }
    if childYangName == "inform" {
        return &statistics.Inform
    }
    if childYangName == "lease-query" {
        return &statistics.LeaseQuery
    }
    if childYangName == "lease-not-assigned" {
        return &statistics.LeaseNotAssigned
    }
    if childYangName == "lease-unknown" {
        return &statistics.LeaseUnknown
    }
    if childYangName == "lease-active" {
        return &statistics.LeaseActive
    }
    if childYangName == "bootp-request" {
        return &statistics.BootpRequest
    }
    if childYangName == "bootp-reply" {
        return &statistics.BootpReply
    }
    return nil
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["discover"] = &statistics.Discover
    children["offer"] = &statistics.Offer
    children["request"] = &statistics.Request
    children["decline"] = &statistics.Decline
    children["ack"] = &statistics.Ack
    children["nak"] = &statistics.Nak
    children["release"] = &statistics.Release
    children["inform"] = &statistics.Inform
    children["lease-query"] = &statistics.LeaseQuery
    children["lease-not-assigned"] = &statistics.LeaseNotAssigned
    children["lease-unknown"] = &statistics.LeaseUnknown
    children["lease-active"] = &statistics.LeaseActive
    children["bootp-request"] = &statistics.BootpRequest
    children["bootp-reply"] = &statistics.BootpReply
    return children
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics) GetYangName() string { return "statistics" }

func (statistics *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics) GetParentYangName() string { return "vrf" }

// Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Discover
// DHCP discover packets
type Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Discover struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Received packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPackets interface{}

    // Transmitted packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedPackets interface{}

    // Dropped packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}
}

func (discover *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Discover) GetFilter() yfilter.YFilter { return discover.YFilter }

func (discover *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Discover) SetFilter(yf yfilter.YFilter) { discover.YFilter = yf }

func (discover *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Discover) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (discover *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Discover) GetSegmentPath() string {
    return "discover"
}

func (discover *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Discover) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (discover *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Discover) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (discover *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Discover) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = discover.ReceivedPackets
    leafs["transmitted-packets"] = discover.TransmittedPackets
    leafs["dropped-packets"] = discover.DroppedPackets
    return leafs
}

func (discover *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Discover) GetBundleName() string { return "cisco_ios_xr" }

func (discover *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Discover) GetYangName() string { return "discover" }

func (discover *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Discover) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (discover *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Discover) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (discover *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Discover) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (discover *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Discover) SetParent(parent types.Entity) { discover.parent = parent }

func (discover *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Discover) GetParent() types.Entity { return discover.parent }

func (discover *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Discover) GetParentYangName() string { return "statistics" }

// Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Offer
// DHCP offer packets
type Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Offer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Received packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPackets interface{}

    // Transmitted packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedPackets interface{}

    // Dropped packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}
}

func (offer *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Offer) GetFilter() yfilter.YFilter { return offer.YFilter }

func (offer *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Offer) SetFilter(yf yfilter.YFilter) { offer.YFilter = yf }

func (offer *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Offer) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (offer *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Offer) GetSegmentPath() string {
    return "offer"
}

func (offer *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Offer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (offer *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Offer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (offer *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Offer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = offer.ReceivedPackets
    leafs["transmitted-packets"] = offer.TransmittedPackets
    leafs["dropped-packets"] = offer.DroppedPackets
    return leafs
}

func (offer *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Offer) GetBundleName() string { return "cisco_ios_xr" }

func (offer *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Offer) GetYangName() string { return "offer" }

func (offer *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Offer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (offer *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Offer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (offer *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Offer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (offer *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Offer) SetParent(parent types.Entity) { offer.parent = parent }

func (offer *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Offer) GetParent() types.Entity { return offer.parent }

func (offer *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Offer) GetParentYangName() string { return "statistics" }

// Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Request
// DHCP request packets
type Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Request struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Received packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPackets interface{}

    // Transmitted packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedPackets interface{}

    // Dropped packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}
}

func (request *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Request) GetFilter() yfilter.YFilter { return request.YFilter }

func (request *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Request) SetFilter(yf yfilter.YFilter) { request.YFilter = yf }

func (request *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Request) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (request *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Request) GetSegmentPath() string {
    return "request"
}

func (request *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Request) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (request *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Request) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (request *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Request) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = request.ReceivedPackets
    leafs["transmitted-packets"] = request.TransmittedPackets
    leafs["dropped-packets"] = request.DroppedPackets
    return leafs
}

func (request *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Request) GetBundleName() string { return "cisco_ios_xr" }

func (request *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Request) GetYangName() string { return "request" }

func (request *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Request) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (request *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Request) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (request *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Request) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (request *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Request) SetParent(parent types.Entity) { request.parent = parent }

func (request *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Request) GetParent() types.Entity { return request.parent }

func (request *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Request) GetParentYangName() string { return "statistics" }

// Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Decline
// DHCP decline packets
type Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Decline struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Received packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPackets interface{}

    // Transmitted packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedPackets interface{}

    // Dropped packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}
}

func (decline *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Decline) GetFilter() yfilter.YFilter { return decline.YFilter }

func (decline *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Decline) SetFilter(yf yfilter.YFilter) { decline.YFilter = yf }

func (decline *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Decline) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (decline *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Decline) GetSegmentPath() string {
    return "decline"
}

func (decline *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Decline) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (decline *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Decline) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (decline *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Decline) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = decline.ReceivedPackets
    leafs["transmitted-packets"] = decline.TransmittedPackets
    leafs["dropped-packets"] = decline.DroppedPackets
    return leafs
}

func (decline *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Decline) GetBundleName() string { return "cisco_ios_xr" }

func (decline *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Decline) GetYangName() string { return "decline" }

func (decline *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Decline) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (decline *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Decline) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (decline *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Decline) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (decline *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Decline) SetParent(parent types.Entity) { decline.parent = parent }

func (decline *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Decline) GetParent() types.Entity { return decline.parent }

func (decline *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Decline) GetParentYangName() string { return "statistics" }

// Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Ack
// DHCP ack packets
type Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Ack struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Received packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPackets interface{}

    // Transmitted packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedPackets interface{}

    // Dropped packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}
}

func (ack *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Ack) GetFilter() yfilter.YFilter { return ack.YFilter }

func (ack *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Ack) SetFilter(yf yfilter.YFilter) { ack.YFilter = yf }

func (ack *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Ack) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (ack *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Ack) GetSegmentPath() string {
    return "ack"
}

func (ack *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Ack) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ack *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Ack) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ack *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Ack) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = ack.ReceivedPackets
    leafs["transmitted-packets"] = ack.TransmittedPackets
    leafs["dropped-packets"] = ack.DroppedPackets
    return leafs
}

func (ack *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Ack) GetBundleName() string { return "cisco_ios_xr" }

func (ack *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Ack) GetYangName() string { return "ack" }

func (ack *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Ack) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ack *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Ack) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ack *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Ack) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ack *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Ack) SetParent(parent types.Entity) { ack.parent = parent }

func (ack *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Ack) GetParent() types.Entity { return ack.parent }

func (ack *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Ack) GetParentYangName() string { return "statistics" }

// Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Nak
// DHCP nak packets
type Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Nak struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Received packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPackets interface{}

    // Transmitted packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedPackets interface{}

    // Dropped packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}
}

func (nak *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Nak) GetFilter() yfilter.YFilter { return nak.YFilter }

func (nak *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Nak) SetFilter(yf yfilter.YFilter) { nak.YFilter = yf }

func (nak *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Nak) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (nak *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Nak) GetSegmentPath() string {
    return "nak"
}

func (nak *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Nak) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nak *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Nak) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nak *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Nak) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = nak.ReceivedPackets
    leafs["transmitted-packets"] = nak.TransmittedPackets
    leafs["dropped-packets"] = nak.DroppedPackets
    return leafs
}

func (nak *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Nak) GetBundleName() string { return "cisco_ios_xr" }

func (nak *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Nak) GetYangName() string { return "nak" }

func (nak *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Nak) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nak *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Nak) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nak *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Nak) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nak *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Nak) SetParent(parent types.Entity) { nak.parent = parent }

func (nak *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Nak) GetParent() types.Entity { return nak.parent }

func (nak *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Nak) GetParentYangName() string { return "statistics" }

// Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Release
// DHCP release packets
type Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Release struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Received packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPackets interface{}

    // Transmitted packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedPackets interface{}

    // Dropped packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}
}

func (release *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Release) GetFilter() yfilter.YFilter { return release.YFilter }

func (release *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Release) SetFilter(yf yfilter.YFilter) { release.YFilter = yf }

func (release *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Release) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (release *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Release) GetSegmentPath() string {
    return "release"
}

func (release *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Release) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (release *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Release) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (release *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Release) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = release.ReceivedPackets
    leafs["transmitted-packets"] = release.TransmittedPackets
    leafs["dropped-packets"] = release.DroppedPackets
    return leafs
}

func (release *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Release) GetBundleName() string { return "cisco_ios_xr" }

func (release *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Release) GetYangName() string { return "release" }

func (release *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Release) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (release *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Release) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (release *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Release) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (release *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Release) SetParent(parent types.Entity) { release.parent = parent }

func (release *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Release) GetParent() types.Entity { return release.parent }

func (release *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Release) GetParentYangName() string { return "statistics" }

// Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Inform
// DHCP inform packets
type Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Inform struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Received packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPackets interface{}

    // Transmitted packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedPackets interface{}

    // Dropped packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}
}

func (inform *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Inform) GetFilter() yfilter.YFilter { return inform.YFilter }

func (inform *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Inform) SetFilter(yf yfilter.YFilter) { inform.YFilter = yf }

func (inform *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Inform) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (inform *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Inform) GetSegmentPath() string {
    return "inform"
}

func (inform *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Inform) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inform *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Inform) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inform *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Inform) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = inform.ReceivedPackets
    leafs["transmitted-packets"] = inform.TransmittedPackets
    leafs["dropped-packets"] = inform.DroppedPackets
    return leafs
}

func (inform *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Inform) GetBundleName() string { return "cisco_ios_xr" }

func (inform *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Inform) GetYangName() string { return "inform" }

func (inform *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Inform) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inform *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Inform) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inform *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Inform) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inform *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Inform) SetParent(parent types.Entity) { inform.parent = parent }

func (inform *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Inform) GetParent() types.Entity { return inform.parent }

func (inform *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Inform) GetParentYangName() string { return "statistics" }

// Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQuery
// DHCP lease query packets
type Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQuery struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Received packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPackets interface{}

    // Transmitted packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedPackets interface{}

    // Dropped packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}
}

func (leaseQuery *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQuery) GetFilter() yfilter.YFilter { return leaseQuery.YFilter }

func (leaseQuery *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQuery) SetFilter(yf yfilter.YFilter) { leaseQuery.YFilter = yf }

func (leaseQuery *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQuery) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (leaseQuery *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQuery) GetSegmentPath() string {
    return "lease-query"
}

func (leaseQuery *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQuery) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (leaseQuery *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQuery) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (leaseQuery *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQuery) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = leaseQuery.ReceivedPackets
    leafs["transmitted-packets"] = leaseQuery.TransmittedPackets
    leafs["dropped-packets"] = leaseQuery.DroppedPackets
    return leafs
}

func (leaseQuery *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQuery) GetBundleName() string { return "cisco_ios_xr" }

func (leaseQuery *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQuery) GetYangName() string { return "lease-query" }

func (leaseQuery *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQuery) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (leaseQuery *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQuery) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (leaseQuery *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQuery) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (leaseQuery *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQuery) SetParent(parent types.Entity) { leaseQuery.parent = parent }

func (leaseQuery *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQuery) GetParent() types.Entity { return leaseQuery.parent }

func (leaseQuery *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQuery) GetParentYangName() string { return "statistics" }

// Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseNotAssigned
// DHCP lease not assigned
type Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseNotAssigned struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Received packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPackets interface{}

    // Transmitted packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedPackets interface{}

    // Dropped packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}
}

func (leaseNotAssigned *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseNotAssigned) GetFilter() yfilter.YFilter { return leaseNotAssigned.YFilter }

func (leaseNotAssigned *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseNotAssigned) SetFilter(yf yfilter.YFilter) { leaseNotAssigned.YFilter = yf }

func (leaseNotAssigned *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseNotAssigned) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (leaseNotAssigned *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseNotAssigned) GetSegmentPath() string {
    return "lease-not-assigned"
}

func (leaseNotAssigned *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseNotAssigned) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (leaseNotAssigned *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseNotAssigned) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (leaseNotAssigned *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseNotAssigned) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = leaseNotAssigned.ReceivedPackets
    leafs["transmitted-packets"] = leaseNotAssigned.TransmittedPackets
    leafs["dropped-packets"] = leaseNotAssigned.DroppedPackets
    return leafs
}

func (leaseNotAssigned *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseNotAssigned) GetBundleName() string { return "cisco_ios_xr" }

func (leaseNotAssigned *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseNotAssigned) GetYangName() string { return "lease-not-assigned" }

func (leaseNotAssigned *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseNotAssigned) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (leaseNotAssigned *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseNotAssigned) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (leaseNotAssigned *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseNotAssigned) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (leaseNotAssigned *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseNotAssigned) SetParent(parent types.Entity) { leaseNotAssigned.parent = parent }

func (leaseNotAssigned *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseNotAssigned) GetParent() types.Entity { return leaseNotAssigned.parent }

func (leaseNotAssigned *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseNotAssigned) GetParentYangName() string { return "statistics" }

// Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseUnknown
// DHCP lease unknown
type Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseUnknown struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Received packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPackets interface{}

    // Transmitted packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedPackets interface{}

    // Dropped packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}
}

func (leaseUnknown *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseUnknown) GetFilter() yfilter.YFilter { return leaseUnknown.YFilter }

func (leaseUnknown *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseUnknown) SetFilter(yf yfilter.YFilter) { leaseUnknown.YFilter = yf }

func (leaseUnknown *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseUnknown) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (leaseUnknown *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseUnknown) GetSegmentPath() string {
    return "lease-unknown"
}

func (leaseUnknown *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseUnknown) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (leaseUnknown *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseUnknown) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (leaseUnknown *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseUnknown) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = leaseUnknown.ReceivedPackets
    leafs["transmitted-packets"] = leaseUnknown.TransmittedPackets
    leafs["dropped-packets"] = leaseUnknown.DroppedPackets
    return leafs
}

func (leaseUnknown *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseUnknown) GetBundleName() string { return "cisco_ios_xr" }

func (leaseUnknown *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseUnknown) GetYangName() string { return "lease-unknown" }

func (leaseUnknown *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseUnknown) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (leaseUnknown *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseUnknown) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (leaseUnknown *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseUnknown) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (leaseUnknown *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseUnknown) SetParent(parent types.Entity) { leaseUnknown.parent = parent }

func (leaseUnknown *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseUnknown) GetParent() types.Entity { return leaseUnknown.parent }

func (leaseUnknown *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseUnknown) GetParentYangName() string { return "statistics" }

// Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseActive
// DHCP lease active
type Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseActive struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Received packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPackets interface{}

    // Transmitted packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedPackets interface{}

    // Dropped packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}
}

func (leaseActive *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseActive) GetFilter() yfilter.YFilter { return leaseActive.YFilter }

func (leaseActive *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseActive) SetFilter(yf yfilter.YFilter) { leaseActive.YFilter = yf }

func (leaseActive *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseActive) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (leaseActive *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseActive) GetSegmentPath() string {
    return "lease-active"
}

func (leaseActive *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseActive) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (leaseActive *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseActive) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (leaseActive *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseActive) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = leaseActive.ReceivedPackets
    leafs["transmitted-packets"] = leaseActive.TransmittedPackets
    leafs["dropped-packets"] = leaseActive.DroppedPackets
    return leafs
}

func (leaseActive *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseActive) GetBundleName() string { return "cisco_ios_xr" }

func (leaseActive *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseActive) GetYangName() string { return "lease-active" }

func (leaseActive *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseActive) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (leaseActive *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseActive) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (leaseActive *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseActive) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (leaseActive *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseActive) SetParent(parent types.Entity) { leaseActive.parent = parent }

func (leaseActive *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseActive) GetParent() types.Entity { return leaseActive.parent }

func (leaseActive *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseActive) GetParentYangName() string { return "statistics" }

// Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_BootpRequest
// DHCP BOOTP request
type Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_BootpRequest struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Received packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPackets interface{}

    // Transmitted packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedPackets interface{}

    // Dropped packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}
}

func (bootpRequest *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_BootpRequest) GetFilter() yfilter.YFilter { return bootpRequest.YFilter }

func (bootpRequest *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_BootpRequest) SetFilter(yf yfilter.YFilter) { bootpRequest.YFilter = yf }

func (bootpRequest *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_BootpRequest) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (bootpRequest *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_BootpRequest) GetSegmentPath() string {
    return "bootp-request"
}

func (bootpRequest *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_BootpRequest) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bootpRequest *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_BootpRequest) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bootpRequest *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_BootpRequest) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = bootpRequest.ReceivedPackets
    leafs["transmitted-packets"] = bootpRequest.TransmittedPackets
    leafs["dropped-packets"] = bootpRequest.DroppedPackets
    return leafs
}

func (bootpRequest *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_BootpRequest) GetBundleName() string { return "cisco_ios_xr" }

func (bootpRequest *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_BootpRequest) GetYangName() string { return "bootp-request" }

func (bootpRequest *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_BootpRequest) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bootpRequest *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_BootpRequest) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bootpRequest *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_BootpRequest) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bootpRequest *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_BootpRequest) SetParent(parent types.Entity) { bootpRequest.parent = parent }

func (bootpRequest *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_BootpRequest) GetParent() types.Entity { return bootpRequest.parent }

func (bootpRequest *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_BootpRequest) GetParentYangName() string { return "statistics" }

// Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_BootpReply
// DHCP BOOTP reply
type Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_BootpReply struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Received packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPackets interface{}

    // Transmitted packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedPackets interface{}

    // Dropped packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}
}

func (bootpReply *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_BootpReply) GetFilter() yfilter.YFilter { return bootpReply.YFilter }

func (bootpReply *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_BootpReply) SetFilter(yf yfilter.YFilter) { bootpReply.YFilter = yf }

func (bootpReply *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_BootpReply) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (bootpReply *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_BootpReply) GetSegmentPath() string {
    return "bootp-reply"
}

func (bootpReply *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_BootpReply) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bootpReply *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_BootpReply) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bootpReply *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_BootpReply) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = bootpReply.ReceivedPackets
    leafs["transmitted-packets"] = bootpReply.TransmittedPackets
    leafs["dropped-packets"] = bootpReply.DroppedPackets
    return leafs
}

func (bootpReply *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_BootpReply) GetBundleName() string { return "cisco_ios_xr" }

func (bootpReply *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_BootpReply) GetYangName() string { return "bootp-reply" }

func (bootpReply *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_BootpReply) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bootpReply *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_BootpReply) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bootpReply *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_BootpReply) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bootpReply *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_BootpReply) SetParent(parent types.Entity) { bootpReply.parent = parent }

func (bootpReply *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_BootpReply) GetParent() types.Entity { return bootpReply.parent }

func (bootpReply *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_BootpReply) GetParentYangName() string { return "statistics" }

// Ipv4Dhcpd_Nodes_Node_Proxy_Profiles
// IPv4 DHCP proxy profile
type Ipv4Dhcpd_Nodes_Node_Proxy_Profiles struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 DHCP proxy profile. The type is slice of
    // Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile.
    Profile []Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile
}

func (profiles *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles) GetFilter() yfilter.YFilter { return profiles.YFilter }

func (profiles *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles) SetFilter(yf yfilter.YFilter) { profiles.YFilter = yf }

func (profiles *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles) GetGoName(yname string) string {
    if yname == "profile" { return "Profile" }
    return ""
}

func (profiles *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles) GetSegmentPath() string {
    return "profiles"
}

func (profiles *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "profile" {
        for _, c := range profiles.Profile {
            if profiles.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile{}
        profiles.Profile = append(profiles.Profile, child)
        return &profiles.Profile[len(profiles.Profile)-1]
    }
    return nil
}

func (profiles *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range profiles.Profile {
        children[profiles.Profile[i].GetSegmentPath()] = &profiles.Profile[i]
    }
    return children
}

func (profiles *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (profiles *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles) GetBundleName() string { return "cisco_ios_xr" }

func (profiles *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles) GetYangName() string { return "profiles" }

func (profiles *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (profiles *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (profiles *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (profiles *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles) SetParent(parent types.Entity) { profiles.parent = parent }

func (profiles *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles) GetParent() types.Entity { return profiles.parent }

func (profiles *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles) GetParentYangName() string { return "proxy" }

// Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile
// IPv4 DHCP proxy profile
type Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Profile name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    ProfileName interface{}

    // Is true if relay option is enabled. The type is bool.
    IsRelayOptionEnabled interface{}

    // Relay policy. The type is RelayInfoPolicy.
    RelayPolicy interface{}

    // Relay authenticate. The type is RelayInfoAuthenticate.
    RelayAuthenticate interface{}

    // Is true if relay untrusted is enabled. The type is bool.
    IsRelayAllowUntrustedEnabled interface{}

    // Is true if relay VPN enabled. The type is bool.
    IsRelayOptionvpnEnabled interface{}

    // Relay VPN RFC/Cisco mode. The type is RelayInfoVpnMode.
    RelayOptionvpnEnabledMode interface{}

    // Is true if relay check enabled. The type is bool.
    IsRelayCheck interface{}

    // Is true if dhcp subscriber is allowed to move. The type is bool.
    IsMoveAllowed interface{}

    // Broadcast policy. The type is BroadcastFlag.
    ProxyBroadcastFlagPolicy interface{}

    // Client lease time in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    ProxyProfileClientLeaseTime interface{}

    // Lease limit type. The type is ProxyLeaseLimit.
    ProxyLeaseLimitType interface{}

    // Lease limit count. The type is interface{} with range: 0..4294967295.
    ProxyLeaseLimitCount interface{}

    // Helper addresses. The type is slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    ProfileHelperAddress []interface{}

    // VRF names. The type is slice of string with length: 0..33.
    VrfName []interface{}

    // Gateway addresses. The type is slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    GiAddr []interface{}

    // VRF references.
    VrfReferences Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_VrfReferences

    // Interface references.
    InterfaceReferences Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_InterfaceReferences
}

func (profile *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile) GetFilter() yfilter.YFilter { return profile.YFilter }

func (profile *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile) SetFilter(yf yfilter.YFilter) { profile.YFilter = yf }

func (profile *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile) GetGoName(yname string) string {
    if yname == "profile-name" { return "ProfileName" }
    if yname == "is-relay-option-enabled" { return "IsRelayOptionEnabled" }
    if yname == "relay-policy" { return "RelayPolicy" }
    if yname == "relay-authenticate" { return "RelayAuthenticate" }
    if yname == "is-relay-allow-untrusted-enabled" { return "IsRelayAllowUntrustedEnabled" }
    if yname == "is-relay-optionvpn-enabled" { return "IsRelayOptionvpnEnabled" }
    if yname == "relay-optionvpn-enabled-mode" { return "RelayOptionvpnEnabledMode" }
    if yname == "is-relay-check" { return "IsRelayCheck" }
    if yname == "is-move-allowed" { return "IsMoveAllowed" }
    if yname == "proxy-broadcast-flag-policy" { return "ProxyBroadcastFlagPolicy" }
    if yname == "proxy-profile-client-lease-time" { return "ProxyProfileClientLeaseTime" }
    if yname == "proxy-lease-limit-type" { return "ProxyLeaseLimitType" }
    if yname == "proxy-lease-limit-count" { return "ProxyLeaseLimitCount" }
    if yname == "profile-helper-address" { return "ProfileHelperAddress" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "gi-addr" { return "GiAddr" }
    if yname == "vrf-references" { return "VrfReferences" }
    if yname == "interface-references" { return "InterfaceReferences" }
    return ""
}

func (profile *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile) GetSegmentPath() string {
    return "profile" + "[profile-name='" + fmt.Sprintf("%v", profile.ProfileName) + "']"
}

func (profile *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf-references" {
        return &profile.VrfReferences
    }
    if childYangName == "interface-references" {
        return &profile.InterfaceReferences
    }
    return nil
}

func (profile *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["vrf-references"] = &profile.VrfReferences
    children["interface-references"] = &profile.InterfaceReferences
    return children
}

func (profile *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["profile-name"] = profile.ProfileName
    leafs["is-relay-option-enabled"] = profile.IsRelayOptionEnabled
    leafs["relay-policy"] = profile.RelayPolicy
    leafs["relay-authenticate"] = profile.RelayAuthenticate
    leafs["is-relay-allow-untrusted-enabled"] = profile.IsRelayAllowUntrustedEnabled
    leafs["is-relay-optionvpn-enabled"] = profile.IsRelayOptionvpnEnabled
    leafs["relay-optionvpn-enabled-mode"] = profile.RelayOptionvpnEnabledMode
    leafs["is-relay-check"] = profile.IsRelayCheck
    leafs["is-move-allowed"] = profile.IsMoveAllowed
    leafs["proxy-broadcast-flag-policy"] = profile.ProxyBroadcastFlagPolicy
    leafs["proxy-profile-client-lease-time"] = profile.ProxyProfileClientLeaseTime
    leafs["proxy-lease-limit-type"] = profile.ProxyLeaseLimitType
    leafs["proxy-lease-limit-count"] = profile.ProxyLeaseLimitCount
    leafs["profile-helper-address"] = profile.ProfileHelperAddress
    leafs["vrf-name"] = profile.VrfName
    leafs["gi-addr"] = profile.GiAddr
    return leafs
}

func (profile *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile) GetBundleName() string { return "cisco_ios_xr" }

func (profile *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile) GetYangName() string { return "profile" }

func (profile *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (profile *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (profile *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (profile *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile) SetParent(parent types.Entity) { profile.parent = parent }

func (profile *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile) GetParent() types.Entity { return profile.parent }

func (profile *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile) GetParentYangName() string { return "profiles" }

// Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_VrfReferences
// VRF references
type Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_VrfReferences struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ipv4 dhcpd proxy vrf reference. The type is slice of
    // Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_VrfReferences_Ipv4DhcpdProxyVrfReference.
    Ipv4DhcpdProxyVrfReference []Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_VrfReferences_Ipv4DhcpdProxyVrfReference
}

func (vrfReferences *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_VrfReferences) GetFilter() yfilter.YFilter { return vrfReferences.YFilter }

func (vrfReferences *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_VrfReferences) SetFilter(yf yfilter.YFilter) { vrfReferences.YFilter = yf }

func (vrfReferences *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_VrfReferences) GetGoName(yname string) string {
    if yname == "ipv4-dhcpd-proxy-vrf-reference" { return "Ipv4DhcpdProxyVrfReference" }
    return ""
}

func (vrfReferences *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_VrfReferences) GetSegmentPath() string {
    return "vrf-references"
}

func (vrfReferences *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_VrfReferences) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv4-dhcpd-proxy-vrf-reference" {
        for _, c := range vrfReferences.Ipv4DhcpdProxyVrfReference {
            if vrfReferences.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_VrfReferences_Ipv4DhcpdProxyVrfReference{}
        vrfReferences.Ipv4DhcpdProxyVrfReference = append(vrfReferences.Ipv4DhcpdProxyVrfReference, child)
        return &vrfReferences.Ipv4DhcpdProxyVrfReference[len(vrfReferences.Ipv4DhcpdProxyVrfReference)-1]
    }
    return nil
}

func (vrfReferences *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_VrfReferences) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vrfReferences.Ipv4DhcpdProxyVrfReference {
        children[vrfReferences.Ipv4DhcpdProxyVrfReference[i].GetSegmentPath()] = &vrfReferences.Ipv4DhcpdProxyVrfReference[i]
    }
    return children
}

func (vrfReferences *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_VrfReferences) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vrfReferences *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_VrfReferences) GetBundleName() string { return "cisco_ios_xr" }

func (vrfReferences *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_VrfReferences) GetYangName() string { return "vrf-references" }

func (vrfReferences *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_VrfReferences) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfReferences *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_VrfReferences) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfReferences *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_VrfReferences) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfReferences *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_VrfReferences) SetParent(parent types.Entity) { vrfReferences.parent = parent }

func (vrfReferences *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_VrfReferences) GetParent() types.Entity { return vrfReferences.parent }

func (vrfReferences *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_VrfReferences) GetParentYangName() string { return "profile" }

// Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_VrfReferences_Ipv4DhcpdProxyVrfReference
// ipv4 dhcpd proxy vrf reference
type Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_VrfReferences_Ipv4DhcpdProxyVrfReference struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF name. The type is string with length: 0..33.
    ProxyReferenceVrfName interface{}
}

func (ipv4DhcpdProxyVrfReference *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_VrfReferences_Ipv4DhcpdProxyVrfReference) GetFilter() yfilter.YFilter { return ipv4DhcpdProxyVrfReference.YFilter }

func (ipv4DhcpdProxyVrfReference *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_VrfReferences_Ipv4DhcpdProxyVrfReference) SetFilter(yf yfilter.YFilter) { ipv4DhcpdProxyVrfReference.YFilter = yf }

func (ipv4DhcpdProxyVrfReference *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_VrfReferences_Ipv4DhcpdProxyVrfReference) GetGoName(yname string) string {
    if yname == "proxy-reference-vrf-name" { return "ProxyReferenceVrfName" }
    return ""
}

func (ipv4DhcpdProxyVrfReference *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_VrfReferences_Ipv4DhcpdProxyVrfReference) GetSegmentPath() string {
    return "ipv4-dhcpd-proxy-vrf-reference"
}

func (ipv4DhcpdProxyVrfReference *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_VrfReferences_Ipv4DhcpdProxyVrfReference) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv4DhcpdProxyVrfReference *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_VrfReferences_Ipv4DhcpdProxyVrfReference) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv4DhcpdProxyVrfReference *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_VrfReferences_Ipv4DhcpdProxyVrfReference) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["proxy-reference-vrf-name"] = ipv4DhcpdProxyVrfReference.ProxyReferenceVrfName
    return leafs
}

func (ipv4DhcpdProxyVrfReference *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_VrfReferences_Ipv4DhcpdProxyVrfReference) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4DhcpdProxyVrfReference *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_VrfReferences_Ipv4DhcpdProxyVrfReference) GetYangName() string { return "ipv4-dhcpd-proxy-vrf-reference" }

func (ipv4DhcpdProxyVrfReference *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_VrfReferences_Ipv4DhcpdProxyVrfReference) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4DhcpdProxyVrfReference *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_VrfReferences_Ipv4DhcpdProxyVrfReference) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4DhcpdProxyVrfReference *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_VrfReferences_Ipv4DhcpdProxyVrfReference) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4DhcpdProxyVrfReference *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_VrfReferences_Ipv4DhcpdProxyVrfReference) SetParent(parent types.Entity) { ipv4DhcpdProxyVrfReference.parent = parent }

func (ipv4DhcpdProxyVrfReference *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_VrfReferences_Ipv4DhcpdProxyVrfReference) GetParent() types.Entity { return ipv4DhcpdProxyVrfReference.parent }

func (ipv4DhcpdProxyVrfReference *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_VrfReferences_Ipv4DhcpdProxyVrfReference) GetParentYangName() string { return "vrf-references" }

// Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_InterfaceReferences
// Interface references
type Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_InterfaceReferences struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ipv4 dhcpd proxy interface reference. The type is slice of
    // Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_InterfaceReferences_Ipv4DhcpdProxyInterfaceReference.
    Ipv4DhcpdProxyInterfaceReference []Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_InterfaceReferences_Ipv4DhcpdProxyInterfaceReference
}

func (interfaceReferences *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_InterfaceReferences) GetFilter() yfilter.YFilter { return interfaceReferences.YFilter }

func (interfaceReferences *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_InterfaceReferences) SetFilter(yf yfilter.YFilter) { interfaceReferences.YFilter = yf }

func (interfaceReferences *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_InterfaceReferences) GetGoName(yname string) string {
    if yname == "ipv4-dhcpd-proxy-interface-reference" { return "Ipv4DhcpdProxyInterfaceReference" }
    return ""
}

func (interfaceReferences *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_InterfaceReferences) GetSegmentPath() string {
    return "interface-references"
}

func (interfaceReferences *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_InterfaceReferences) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv4-dhcpd-proxy-interface-reference" {
        for _, c := range interfaceReferences.Ipv4DhcpdProxyInterfaceReference {
            if interfaceReferences.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_InterfaceReferences_Ipv4DhcpdProxyInterfaceReference{}
        interfaceReferences.Ipv4DhcpdProxyInterfaceReference = append(interfaceReferences.Ipv4DhcpdProxyInterfaceReference, child)
        return &interfaceReferences.Ipv4DhcpdProxyInterfaceReference[len(interfaceReferences.Ipv4DhcpdProxyInterfaceReference)-1]
    }
    return nil
}

func (interfaceReferences *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_InterfaceReferences) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaceReferences.Ipv4DhcpdProxyInterfaceReference {
        children[interfaceReferences.Ipv4DhcpdProxyInterfaceReference[i].GetSegmentPath()] = &interfaceReferences.Ipv4DhcpdProxyInterfaceReference[i]
    }
    return children
}

func (interfaceReferences *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_InterfaceReferences) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceReferences *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_InterfaceReferences) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceReferences *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_InterfaceReferences) GetYangName() string { return "interface-references" }

func (interfaceReferences *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_InterfaceReferences) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceReferences *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_InterfaceReferences) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceReferences *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_InterfaceReferences) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceReferences *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_InterfaceReferences) SetParent(parent types.Entity) { interfaceReferences.parent = parent }

func (interfaceReferences *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_InterfaceReferences) GetParent() types.Entity { return interfaceReferences.parent }

func (interfaceReferences *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_InterfaceReferences) GetParentYangName() string { return "profile" }

// Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_InterfaceReferences_Ipv4DhcpdProxyInterfaceReference
// ipv4 dhcpd proxy interface reference
type Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_InterfaceReferences_Ipv4DhcpdProxyInterfaceReference struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface name. The type is string with length: 0..65.
    ProxyReferenceInterfaceName interface{}
}

func (ipv4DhcpdProxyInterfaceReference *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_InterfaceReferences_Ipv4DhcpdProxyInterfaceReference) GetFilter() yfilter.YFilter { return ipv4DhcpdProxyInterfaceReference.YFilter }

func (ipv4DhcpdProxyInterfaceReference *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_InterfaceReferences_Ipv4DhcpdProxyInterfaceReference) SetFilter(yf yfilter.YFilter) { ipv4DhcpdProxyInterfaceReference.YFilter = yf }

func (ipv4DhcpdProxyInterfaceReference *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_InterfaceReferences_Ipv4DhcpdProxyInterfaceReference) GetGoName(yname string) string {
    if yname == "proxy-reference-interface-name" { return "ProxyReferenceInterfaceName" }
    return ""
}

func (ipv4DhcpdProxyInterfaceReference *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_InterfaceReferences_Ipv4DhcpdProxyInterfaceReference) GetSegmentPath() string {
    return "ipv4-dhcpd-proxy-interface-reference"
}

func (ipv4DhcpdProxyInterfaceReference *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_InterfaceReferences_Ipv4DhcpdProxyInterfaceReference) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv4DhcpdProxyInterfaceReference *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_InterfaceReferences_Ipv4DhcpdProxyInterfaceReference) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv4DhcpdProxyInterfaceReference *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_InterfaceReferences_Ipv4DhcpdProxyInterfaceReference) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["proxy-reference-interface-name"] = ipv4DhcpdProxyInterfaceReference.ProxyReferenceInterfaceName
    return leafs
}

func (ipv4DhcpdProxyInterfaceReference *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_InterfaceReferences_Ipv4DhcpdProxyInterfaceReference) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4DhcpdProxyInterfaceReference *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_InterfaceReferences_Ipv4DhcpdProxyInterfaceReference) GetYangName() string { return "ipv4-dhcpd-proxy-interface-reference" }

func (ipv4DhcpdProxyInterfaceReference *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_InterfaceReferences_Ipv4DhcpdProxyInterfaceReference) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4DhcpdProxyInterfaceReference *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_InterfaceReferences_Ipv4DhcpdProxyInterfaceReference) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4DhcpdProxyInterfaceReference *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_InterfaceReferences_Ipv4DhcpdProxyInterfaceReference) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4DhcpdProxyInterfaceReference *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_InterfaceReferences_Ipv4DhcpdProxyInterfaceReference) SetParent(parent types.Entity) { ipv4DhcpdProxyInterfaceReference.parent = parent }

func (ipv4DhcpdProxyInterfaceReference *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_InterfaceReferences_Ipv4DhcpdProxyInterfaceReference) GetParent() types.Entity { return ipv4DhcpdProxyInterfaceReference.parent }

func (ipv4DhcpdProxyInterfaceReference *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_InterfaceReferences_Ipv4DhcpdProxyInterfaceReference) GetParentYangName() string { return "interface-references" }

// Ipv4Dhcpd_Nodes_Node_Proxy_Statistics
// DHCP proxy statistics
type Ipv4Dhcpd_Nodes_Node_Proxy_Statistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ipv4 dhcpd proxy stat. The type is slice of
    // Ipv4Dhcpd_Nodes_Node_Proxy_Statistics_Ipv4DhcpdProxyStat.
    Ipv4DhcpdProxyStat []Ipv4Dhcpd_Nodes_Node_Proxy_Statistics_Ipv4DhcpdProxyStat
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Proxy_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *Ipv4Dhcpd_Nodes_Node_Proxy_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *Ipv4Dhcpd_Nodes_Node_Proxy_Statistics) GetGoName(yname string) string {
    if yname == "ipv4-dhcpd-proxy-stat" { return "Ipv4DhcpdProxyStat" }
    return ""
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Proxy_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Proxy_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv4-dhcpd-proxy-stat" {
        for _, c := range statistics.Ipv4DhcpdProxyStat {
            if statistics.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4Dhcpd_Nodes_Node_Proxy_Statistics_Ipv4DhcpdProxyStat{}
        statistics.Ipv4DhcpdProxyStat = append(statistics.Ipv4DhcpdProxyStat, child)
        return &statistics.Ipv4DhcpdProxyStat[len(statistics.Ipv4DhcpdProxyStat)-1]
    }
    return nil
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Proxy_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range statistics.Ipv4DhcpdProxyStat {
        children[statistics.Ipv4DhcpdProxyStat[i].GetSegmentPath()] = &statistics.Ipv4DhcpdProxyStat[i]
    }
    return children
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Proxy_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Proxy_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *Ipv4Dhcpd_Nodes_Node_Proxy_Statistics) GetYangName() string { return "statistics" }

func (statistics *Ipv4Dhcpd_Nodes_Node_Proxy_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *Ipv4Dhcpd_Nodes_Node_Proxy_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *Ipv4Dhcpd_Nodes_Node_Proxy_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *Ipv4Dhcpd_Nodes_Node_Proxy_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *Ipv4Dhcpd_Nodes_Node_Proxy_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *Ipv4Dhcpd_Nodes_Node_Proxy_Statistics) GetParentYangName() string { return "proxy" }

// Ipv4Dhcpd_Nodes_Node_Proxy_Statistics_Ipv4DhcpdProxyStat
// ipv4 dhcpd proxy stat
type Ipv4Dhcpd_Nodes_Node_Proxy_Statistics_Ipv4DhcpdProxyStat struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DHCP L3 VRF name. The type is string with length: 0..33.
    VrfName interface{}

    // Proxy statistics.
    Statistics Ipv4Dhcpd_Nodes_Node_Proxy_Statistics_Ipv4DhcpdProxyStat_Statistics
}

func (ipv4DhcpdProxyStat *Ipv4Dhcpd_Nodes_Node_Proxy_Statistics_Ipv4DhcpdProxyStat) GetFilter() yfilter.YFilter { return ipv4DhcpdProxyStat.YFilter }

func (ipv4DhcpdProxyStat *Ipv4Dhcpd_Nodes_Node_Proxy_Statistics_Ipv4DhcpdProxyStat) SetFilter(yf yfilter.YFilter) { ipv4DhcpdProxyStat.YFilter = yf }

func (ipv4DhcpdProxyStat *Ipv4Dhcpd_Nodes_Node_Proxy_Statistics_Ipv4DhcpdProxyStat) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "statistics" { return "Statistics" }
    return ""
}

func (ipv4DhcpdProxyStat *Ipv4Dhcpd_Nodes_Node_Proxy_Statistics_Ipv4DhcpdProxyStat) GetSegmentPath() string {
    return "ipv4-dhcpd-proxy-stat"
}

func (ipv4DhcpdProxyStat *Ipv4Dhcpd_Nodes_Node_Proxy_Statistics_Ipv4DhcpdProxyStat) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "statistics" {
        return &ipv4DhcpdProxyStat.Statistics
    }
    return nil
}

func (ipv4DhcpdProxyStat *Ipv4Dhcpd_Nodes_Node_Proxy_Statistics_Ipv4DhcpdProxyStat) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["statistics"] = &ipv4DhcpdProxyStat.Statistics
    return children
}

func (ipv4DhcpdProxyStat *Ipv4Dhcpd_Nodes_Node_Proxy_Statistics_Ipv4DhcpdProxyStat) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = ipv4DhcpdProxyStat.VrfName
    return leafs
}

func (ipv4DhcpdProxyStat *Ipv4Dhcpd_Nodes_Node_Proxy_Statistics_Ipv4DhcpdProxyStat) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4DhcpdProxyStat *Ipv4Dhcpd_Nodes_Node_Proxy_Statistics_Ipv4DhcpdProxyStat) GetYangName() string { return "ipv4-dhcpd-proxy-stat" }

func (ipv4DhcpdProxyStat *Ipv4Dhcpd_Nodes_Node_Proxy_Statistics_Ipv4DhcpdProxyStat) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4DhcpdProxyStat *Ipv4Dhcpd_Nodes_Node_Proxy_Statistics_Ipv4DhcpdProxyStat) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4DhcpdProxyStat *Ipv4Dhcpd_Nodes_Node_Proxy_Statistics_Ipv4DhcpdProxyStat) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4DhcpdProxyStat *Ipv4Dhcpd_Nodes_Node_Proxy_Statistics_Ipv4DhcpdProxyStat) SetParent(parent types.Entity) { ipv4DhcpdProxyStat.parent = parent }

func (ipv4DhcpdProxyStat *Ipv4Dhcpd_Nodes_Node_Proxy_Statistics_Ipv4DhcpdProxyStat) GetParent() types.Entity { return ipv4DhcpdProxyStat.parent }

func (ipv4DhcpdProxyStat *Ipv4Dhcpd_Nodes_Node_Proxy_Statistics_Ipv4DhcpdProxyStat) GetParentYangName() string { return "statistics" }

// Ipv4Dhcpd_Nodes_Node_Proxy_Statistics_Ipv4DhcpdProxyStat_Statistics
// Proxy statistics
type Ipv4Dhcpd_Nodes_Node_Proxy_Statistics_Ipv4DhcpdProxyStat_Statistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Received packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPackets interface{}

    // Transmitted packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedPackets interface{}

    // Dropped packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Proxy_Statistics_Ipv4DhcpdProxyStat_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *Ipv4Dhcpd_Nodes_Node_Proxy_Statistics_Ipv4DhcpdProxyStat_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *Ipv4Dhcpd_Nodes_Node_Proxy_Statistics_Ipv4DhcpdProxyStat_Statistics) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Proxy_Statistics_Ipv4DhcpdProxyStat_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Proxy_Statistics_Ipv4DhcpdProxyStat_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Proxy_Statistics_Ipv4DhcpdProxyStat_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Proxy_Statistics_Ipv4DhcpdProxyStat_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = statistics.ReceivedPackets
    leafs["transmitted-packets"] = statistics.TransmittedPackets
    leafs["dropped-packets"] = statistics.DroppedPackets
    return leafs
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Proxy_Statistics_Ipv4DhcpdProxyStat_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *Ipv4Dhcpd_Nodes_Node_Proxy_Statistics_Ipv4DhcpdProxyStat_Statistics) GetYangName() string { return "statistics" }

func (statistics *Ipv4Dhcpd_Nodes_Node_Proxy_Statistics_Ipv4DhcpdProxyStat_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *Ipv4Dhcpd_Nodes_Node_Proxy_Statistics_Ipv4DhcpdProxyStat_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *Ipv4Dhcpd_Nodes_Node_Proxy_Statistics_Ipv4DhcpdProxyStat_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *Ipv4Dhcpd_Nodes_Node_Proxy_Statistics_Ipv4DhcpdProxyStat_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *Ipv4Dhcpd_Nodes_Node_Proxy_Statistics_Ipv4DhcpdProxyStat_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *Ipv4Dhcpd_Nodes_Node_Proxy_Statistics_Ipv4DhcpdProxyStat_Statistics) GetParentYangName() string { return "ipv4-dhcpd-proxy-stat" }

// Ipv4Dhcpd_Nodes_Node_Proxy_Binding
// DHCP proxy bindings
type Ipv4Dhcpd_Nodes_Node_Proxy_Binding struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DHCP proxy client bindings.
    Clients Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Clients

    // DHCP proxy binding summary.
    Summary Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Summary
}

func (binding *Ipv4Dhcpd_Nodes_Node_Proxy_Binding) GetFilter() yfilter.YFilter { return binding.YFilter }

func (binding *Ipv4Dhcpd_Nodes_Node_Proxy_Binding) SetFilter(yf yfilter.YFilter) { binding.YFilter = yf }

func (binding *Ipv4Dhcpd_Nodes_Node_Proxy_Binding) GetGoName(yname string) string {
    if yname == "clients" { return "Clients" }
    if yname == "summary" { return "Summary" }
    return ""
}

func (binding *Ipv4Dhcpd_Nodes_Node_Proxy_Binding) GetSegmentPath() string {
    return "binding"
}

func (binding *Ipv4Dhcpd_Nodes_Node_Proxy_Binding) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "clients" {
        return &binding.Clients
    }
    if childYangName == "summary" {
        return &binding.Summary
    }
    return nil
}

func (binding *Ipv4Dhcpd_Nodes_Node_Proxy_Binding) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["clients"] = &binding.Clients
    children["summary"] = &binding.Summary
    return children
}

func (binding *Ipv4Dhcpd_Nodes_Node_Proxy_Binding) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (binding *Ipv4Dhcpd_Nodes_Node_Proxy_Binding) GetBundleName() string { return "cisco_ios_xr" }

func (binding *Ipv4Dhcpd_Nodes_Node_Proxy_Binding) GetYangName() string { return "binding" }

func (binding *Ipv4Dhcpd_Nodes_Node_Proxy_Binding) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (binding *Ipv4Dhcpd_Nodes_Node_Proxy_Binding) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (binding *Ipv4Dhcpd_Nodes_Node_Proxy_Binding) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (binding *Ipv4Dhcpd_Nodes_Node_Proxy_Binding) SetParent(parent types.Entity) { binding.parent = parent }

func (binding *Ipv4Dhcpd_Nodes_Node_Proxy_Binding) GetParent() types.Entity { return binding.parent }

func (binding *Ipv4Dhcpd_Nodes_Node_Proxy_Binding) GetParentYangName() string { return "proxy" }

// Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Clients
// DHCP proxy client bindings
type Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Clients struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Single DHCP proxy binding. The type is slice of
    // Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Clients_Client.
    Client []Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Clients_Client
}

func (clients *Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Clients) GetFilter() yfilter.YFilter { return clients.YFilter }

func (clients *Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Clients) SetFilter(yf yfilter.YFilter) { clients.YFilter = yf }

func (clients *Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Clients) GetGoName(yname string) string {
    if yname == "client" { return "Client" }
    return ""
}

func (clients *Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Clients) GetSegmentPath() string {
    return "clients"
}

func (clients *Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Clients) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "client" {
        for _, c := range clients.Client {
            if clients.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Clients_Client{}
        clients.Client = append(clients.Client, child)
        return &clients.Client[len(clients.Client)-1]
    }
    return nil
}

func (clients *Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Clients) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range clients.Client {
        children[clients.Client[i].GetSegmentPath()] = &clients.Client[i]
    }
    return children
}

func (clients *Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Clients) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (clients *Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Clients) GetBundleName() string { return "cisco_ios_xr" }

func (clients *Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Clients) GetYangName() string { return "clients" }

func (clients *Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Clients) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clients *Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Clients) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clients *Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Clients) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clients *Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Clients) SetParent(parent types.Entity) { clients.parent = parent }

func (clients *Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Clients) GetParent() types.Entity { return clients.parent }

func (clients *Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Clients) GetParentYangName() string { return "binding" }

// Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Clients_Client
// Single DHCP proxy binding
type Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Clients_Client struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Client ID. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    ClientId interface{}

    // DHCP client identifier. The type is string with length: 0..1275.
    ClientIdXr interface{}

    // DHCP client MAC address. The type is string.
    MacAddress interface{}

    // DHCP client/subscriber VRF name. The type is string with length: 0..33.
    VrfName interface{}

    // DHCP server VRF name. The type is string with length: 0..33.
    ServerVrfName interface{}

    // DHCP IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    IpAddress interface{}

    // DHCP client GIADDR. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    ClientGiAddr interface{}

    // DHCP to server GIADDR. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    ToServerGiAddr interface{}

    // DHCP server IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    ServerIpAddress interface{}

    // DHCP reply server IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    ReplyServerIpAddress interface{}

    // Lease time in seconds. The type is interface{} with range: 0..4294967295.
    // Units are second.
    LeaseTime interface{}

    // Remaining lease time in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    RemainingLeaseTime interface{}

    // DHCP client state. The type is BagDhcpdProxyState.
    State interface{}

    // DHCP access interface to client. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // DHCP access interface VRF name. The type is string with length: 0..33.
    AccessVrfName interface{}

    // DHCP VLAN outer VLAN. The type is interface{} with range: 0..4294967295.
    ProxyBindingOuterTag interface{}

    // DHCP VLAN inner VLAN. The type is interface{} with range: 0..4294967295.
    ProxyBindingInnerTag interface{}

    // DHCP profile name. The type is string with length: 0..65.
    ProfileName interface{}

    // Is true if DHCP next renew from client will be NAK'd. The type is bool.
    IsNakNextRenew interface{}

    // DHCP subscriber label. The type is interface{} with range: 0..4294967295.
    SubscriberLabel interface{}

    // DHCP old subscriber label. The type is interface{} with range:
    // 0..4294967295.
    OldSubscriberLabel interface{}

    // DHCP subscriber interface. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    SubscriberInterfaceName interface{}

    // DHCP received circuit ID. The type is string with length: 0..768.
    RxCircuitId interface{}

    // DHCP transmitted circuit ID. The type is string with length: 0..768.
    TxCircuitId interface{}

    // DHCP received Remote ID. The type is string with length: 0..768.
    RxRemoteId interface{}

    // DHCP transmitted Remote ID. The type is string with length: 0..768.
    TxRemoteId interface{}

    // DHCP received VSISO. The type is string with length: 0..768.
    RxVsiso interface{}

    // DHCP transmitted VSISO. The type is string with length: 0..768.
    TxVsiso interface{}

    // Is true if authentication is on received option82. The type is bool.
    IsAuthReceived interface{}

    // Is true if DHCP subscriber is Mobile. The type is bool.
    IsMblSubscriber interface{}

    // DHCP parameter request option. The type is string with length: 0..513.
    ParamRequest interface{}

    // DHCP saved options. The type is string with length: 0..2051.
    ParamResponse interface{}

    // session start time. The type is interface{} with range:
    // 0..18446744073709551615.
    SessionStartTime interface{}

    // DHCPV4 SRG state. The type is interface{} with range: 0..4294967295.
    SrgState interface{}

    // event history. The type is slice of interface{} with range: 0..4294967295.
    EventHistory []interface{}
}

func (client *Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Clients_Client) GetFilter() yfilter.YFilter { return client.YFilter }

func (client *Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Clients_Client) SetFilter(yf yfilter.YFilter) { client.YFilter = yf }

func (client *Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Clients_Client) GetGoName(yname string) string {
    if yname == "client-id" { return "ClientId" }
    if yname == "client-id-xr" { return "ClientIdXr" }
    if yname == "mac-address" { return "MacAddress" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "server-vrf-name" { return "ServerVrfName" }
    if yname == "ip-address" { return "IpAddress" }
    if yname == "client-gi-addr" { return "ClientGiAddr" }
    if yname == "to-server-gi-addr" { return "ToServerGiAddr" }
    if yname == "server-ip-address" { return "ServerIpAddress" }
    if yname == "reply-server-ip-address" { return "ReplyServerIpAddress" }
    if yname == "lease-time" { return "LeaseTime" }
    if yname == "remaining-lease-time" { return "RemainingLeaseTime" }
    if yname == "state" { return "State" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "access-vrf-name" { return "AccessVrfName" }
    if yname == "proxy-binding-outer-tag" { return "ProxyBindingOuterTag" }
    if yname == "proxy-binding-inner-tag" { return "ProxyBindingInnerTag" }
    if yname == "profile-name" { return "ProfileName" }
    if yname == "is-nak-next-renew" { return "IsNakNextRenew" }
    if yname == "subscriber-label" { return "SubscriberLabel" }
    if yname == "old-subscriber-label" { return "OldSubscriberLabel" }
    if yname == "subscriber-interface-name" { return "SubscriberInterfaceName" }
    if yname == "rx-circuit-id" { return "RxCircuitId" }
    if yname == "tx-circuit-id" { return "TxCircuitId" }
    if yname == "rx-remote-id" { return "RxRemoteId" }
    if yname == "tx-remote-id" { return "TxRemoteId" }
    if yname == "rx-vsiso" { return "RxVsiso" }
    if yname == "tx-vsiso" { return "TxVsiso" }
    if yname == "is-auth-received" { return "IsAuthReceived" }
    if yname == "is-mbl-subscriber" { return "IsMblSubscriber" }
    if yname == "param-request" { return "ParamRequest" }
    if yname == "param-response" { return "ParamResponse" }
    if yname == "session-start-time" { return "SessionStartTime" }
    if yname == "srg-state" { return "SrgState" }
    if yname == "event-history" { return "EventHistory" }
    return ""
}

func (client *Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Clients_Client) GetSegmentPath() string {
    return "client" + "[client-id='" + fmt.Sprintf("%v", client.ClientId) + "']"
}

func (client *Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Clients_Client) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (client *Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Clients_Client) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (client *Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Clients_Client) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["client-id"] = client.ClientId
    leafs["client-id-xr"] = client.ClientIdXr
    leafs["mac-address"] = client.MacAddress
    leafs["vrf-name"] = client.VrfName
    leafs["server-vrf-name"] = client.ServerVrfName
    leafs["ip-address"] = client.IpAddress
    leafs["client-gi-addr"] = client.ClientGiAddr
    leafs["to-server-gi-addr"] = client.ToServerGiAddr
    leafs["server-ip-address"] = client.ServerIpAddress
    leafs["reply-server-ip-address"] = client.ReplyServerIpAddress
    leafs["lease-time"] = client.LeaseTime
    leafs["remaining-lease-time"] = client.RemainingLeaseTime
    leafs["state"] = client.State
    leafs["interface-name"] = client.InterfaceName
    leafs["access-vrf-name"] = client.AccessVrfName
    leafs["proxy-binding-outer-tag"] = client.ProxyBindingOuterTag
    leafs["proxy-binding-inner-tag"] = client.ProxyBindingInnerTag
    leafs["profile-name"] = client.ProfileName
    leafs["is-nak-next-renew"] = client.IsNakNextRenew
    leafs["subscriber-label"] = client.SubscriberLabel
    leafs["old-subscriber-label"] = client.OldSubscriberLabel
    leafs["subscriber-interface-name"] = client.SubscriberInterfaceName
    leafs["rx-circuit-id"] = client.RxCircuitId
    leafs["tx-circuit-id"] = client.TxCircuitId
    leafs["rx-remote-id"] = client.RxRemoteId
    leafs["tx-remote-id"] = client.TxRemoteId
    leafs["rx-vsiso"] = client.RxVsiso
    leafs["tx-vsiso"] = client.TxVsiso
    leafs["is-auth-received"] = client.IsAuthReceived
    leafs["is-mbl-subscriber"] = client.IsMblSubscriber
    leafs["param-request"] = client.ParamRequest
    leafs["param-response"] = client.ParamResponse
    leafs["session-start-time"] = client.SessionStartTime
    leafs["srg-state"] = client.SrgState
    leafs["event-history"] = client.EventHistory
    return leafs
}

func (client *Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Clients_Client) GetBundleName() string { return "cisco_ios_xr" }

func (client *Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Clients_Client) GetYangName() string { return "client" }

func (client *Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Clients_Client) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (client *Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Clients_Client) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (client *Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Clients_Client) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (client *Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Clients_Client) SetParent(parent types.Entity) { client.parent = parent }

func (client *Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Clients_Client) GetParent() types.Entity { return client.parent }

func (client *Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Clients_Client) GetParentYangName() string { return "clients" }

// Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Summary
// DHCP proxy binding summary
type Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Summary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Total number of clients. The type is interface{} with range: 0..4294967295.
    Clients interface{}

    // Number of clients in init state. The type is interface{} with range:
    // 0..4294967295.
    InitializingClients interface{}

    // Number of clients in Init DPM wait state. The type is interface{} with
    // range: 0..4294967295.
    WaitingForDpmInit interface{}

    // Number of clients in Request DPM wait state. The type is interface{} with
    // range: 0..4294967295.
    WaitingForDpmRequest interface{}

    // Number of clients in Init DAPS wait state. The type is interface{} with
    // range: 0..4294967295.
    WaitingForDapsInit interface{}

    // Number of clients in selecting state. The type is interface{} with range:
    // 0..4294967295.
    SelectingClients interface{}

    // Number of clients in Offer sent state. The type is interface{} with range:
    // 0..4294967295.
    OfferSentForClient interface{}

    // Number of clients in requesting state. The type is interface{} with range:
    // 0..4294967295.
    RequestingClients interface{}

    // Number of clients in Waiting for DPM with Request. The type is interface{}
    // with range: 0..4294967295.
    RequestWaitingForDpm interface{}

    // Number of clients in Waiting for DPM with ACK. The type is interface{} with
    // range: 0..4294967295.
    AckWaitingForDpm interface{}

    // Number of clients in bound state. The type is interface{} with range:
    // 0..4294967295.
    BoundClients interface{}

    // Number of clients in renewing state. The type is interface{} with range:
    // 0..4294967295.
    RenewingClients interface{}

    // Number of clients in informing state. The type is interface{} with range:
    // 0..4294967295.
    InformingClients interface{}

    // Number of clients in reauth state. The type is interface{} with range:
    // 0..4294967295.
    ReauthorizingClients interface{}

    // Number of clients in waiting for DPM disconnect state. The type is
    // interface{} with range: 0..4294967295.
    WaitingForDpmDisconnect interface{}

    // Number of clients in Waiting for DPM after addr change. The type is
    // interface{} with range: 0..4294967295.
    WaitingForDpmAddrChange interface{}

    // Number of clients in deleting state. The type is interface{} with range:
    // 0..4294967295.
    DeletingClientsD interface{}

    // Number of clients in disconnected state. The type is interface{} with
    // range: 0..4294967295.
    DisconnectedClients interface{}

    // Number of clients in restarting state. The type is interface{} with range:
    // 0..4294967295.
    RestartingClients interface{}
}

func (summary *Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Summary) GetFilter() yfilter.YFilter { return summary.YFilter }

func (summary *Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Summary) SetFilter(yf yfilter.YFilter) { summary.YFilter = yf }

func (summary *Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Summary) GetGoName(yname string) string {
    if yname == "clients" { return "Clients" }
    if yname == "initializing-clients" { return "InitializingClients" }
    if yname == "waiting-for-dpm-init" { return "WaitingForDpmInit" }
    if yname == "waiting-for-dpm-request" { return "WaitingForDpmRequest" }
    if yname == "waiting-for-daps-init" { return "WaitingForDapsInit" }
    if yname == "selecting-clients" { return "SelectingClients" }
    if yname == "offer-sent-for-client" { return "OfferSentForClient" }
    if yname == "requesting-clients" { return "RequestingClients" }
    if yname == "request-waiting-for-dpm" { return "RequestWaitingForDpm" }
    if yname == "ack-waiting-for-dpm" { return "AckWaitingForDpm" }
    if yname == "bound-clients" { return "BoundClients" }
    if yname == "renewing-clients" { return "RenewingClients" }
    if yname == "informing-clients" { return "InformingClients" }
    if yname == "reauthorizing-clients" { return "ReauthorizingClients" }
    if yname == "waiting-for-dpm-disconnect" { return "WaitingForDpmDisconnect" }
    if yname == "waiting-for-dpm-addr-change" { return "WaitingForDpmAddrChange" }
    if yname == "deleting-clients-d" { return "DeletingClientsD" }
    if yname == "disconnected-clients" { return "DisconnectedClients" }
    if yname == "restarting-clients" { return "RestartingClients" }
    return ""
}

func (summary *Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Summary) GetSegmentPath() string {
    return "summary"
}

func (summary *Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Summary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (summary *Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Summary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (summary *Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Summary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["clients"] = summary.Clients
    leafs["initializing-clients"] = summary.InitializingClients
    leafs["waiting-for-dpm-init"] = summary.WaitingForDpmInit
    leafs["waiting-for-dpm-request"] = summary.WaitingForDpmRequest
    leafs["waiting-for-daps-init"] = summary.WaitingForDapsInit
    leafs["selecting-clients"] = summary.SelectingClients
    leafs["offer-sent-for-client"] = summary.OfferSentForClient
    leafs["requesting-clients"] = summary.RequestingClients
    leafs["request-waiting-for-dpm"] = summary.RequestWaitingForDpm
    leafs["ack-waiting-for-dpm"] = summary.AckWaitingForDpm
    leafs["bound-clients"] = summary.BoundClients
    leafs["renewing-clients"] = summary.RenewingClients
    leafs["informing-clients"] = summary.InformingClients
    leafs["reauthorizing-clients"] = summary.ReauthorizingClients
    leafs["waiting-for-dpm-disconnect"] = summary.WaitingForDpmDisconnect
    leafs["waiting-for-dpm-addr-change"] = summary.WaitingForDpmAddrChange
    leafs["deleting-clients-d"] = summary.DeletingClientsD
    leafs["disconnected-clients"] = summary.DisconnectedClients
    leafs["restarting-clients"] = summary.RestartingClients
    return leafs
}

func (summary *Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Summary) GetBundleName() string { return "cisco_ios_xr" }

func (summary *Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Summary) GetYangName() string { return "summary" }

func (summary *Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Summary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (summary *Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Summary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (summary *Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Summary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (summary *Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Summary) SetParent(parent types.Entity) { summary.parent = parent }

func (summary *Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Summary) GetParent() types.Entity { return summary.parent }

func (summary *Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Summary) GetParentYangName() string { return "binding" }

// Ipv4Dhcpd_Nodes_Node_Interfaces
// IPv4 DHCP proxy/server Interface
type Ipv4Dhcpd_Nodes_Node_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 DHCP proxy/server interface info. The type is slice of
    // Ipv4Dhcpd_Nodes_Node_Interfaces_Interface.
    Interface []Ipv4Dhcpd_Nodes_Node_Interfaces_Interface
}

func (interfaces *Ipv4Dhcpd_Nodes_Node_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *Ipv4Dhcpd_Nodes_Node_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *Ipv4Dhcpd_Nodes_Node_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *Ipv4Dhcpd_Nodes_Node_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *Ipv4Dhcpd_Nodes_Node_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4Dhcpd_Nodes_Node_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *Ipv4Dhcpd_Nodes_Node_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *Ipv4Dhcpd_Nodes_Node_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *Ipv4Dhcpd_Nodes_Node_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *Ipv4Dhcpd_Nodes_Node_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *Ipv4Dhcpd_Nodes_Node_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *Ipv4Dhcpd_Nodes_Node_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *Ipv4Dhcpd_Nodes_Node_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *Ipv4Dhcpd_Nodes_Node_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *Ipv4Dhcpd_Nodes_Node_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *Ipv4Dhcpd_Nodes_Node_Interfaces) GetParentYangName() string { return "node" }

// Ipv4Dhcpd_Nodes_Node_Interfaces_Interface
// IPv4 DHCP proxy/server interface info
type Ipv4Dhcpd_Nodes_Node_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    InterfaceName interface{}

    // Ifhandle of the interface. The type is interface{} with range:
    // 0..4294967295.
    IntfIfhandle interface{}

    // VRF name. The type is string with length: 0..33.
    VrfName interface{}

    // Mode of interface. The type is interface{} with range: 0..4294967295.
    IntfMode interface{}

    // Is interface ambiguous. The type is interface{} with range: 0..4294967295.
    IntfIsAmbiguous interface{}

    // Name of profile attached to the interface. The type is string with length:
    // 0..65.
    IntfProfileName interface{}

    // Lease limit type on interface. The type is interface{} with range:
    // 0..4294967295.
    IntfLeaseLimitType interface{}

    // Lease limit count on interface. The type is interface{} with range:
    // 0..4294967295.
    IntfLeaseLimitCount interface{}

    // DHCPv6 Interface SRG role. The type is BagDhcpdIntfSrgRole.
    SrgRole interface{}

    // Mac Throttle Status. The type is bool.
    MacThrottle interface{}
}

func (self *Ipv4Dhcpd_Nodes_Node_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Ipv4Dhcpd_Nodes_Node_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Ipv4Dhcpd_Nodes_Node_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "intf-ifhandle" { return "IntfIfhandle" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "intf-mode" { return "IntfMode" }
    if yname == "intf-is-ambiguous" { return "IntfIsAmbiguous" }
    if yname == "intf-profile-name" { return "IntfProfileName" }
    if yname == "intf-lease-limit-type" { return "IntfLeaseLimitType" }
    if yname == "intf-lease-limit-count" { return "IntfLeaseLimitCount" }
    if yname == "srg-role" { return "SrgRole" }
    if yname == "mac-throttle" { return "MacThrottle" }
    return ""
}

func (self *Ipv4Dhcpd_Nodes_Node_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *Ipv4Dhcpd_Nodes_Node_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *Ipv4Dhcpd_Nodes_Node_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *Ipv4Dhcpd_Nodes_Node_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    leafs["intf-ifhandle"] = self.IntfIfhandle
    leafs["vrf-name"] = self.VrfName
    leafs["intf-mode"] = self.IntfMode
    leafs["intf-is-ambiguous"] = self.IntfIsAmbiguous
    leafs["intf-profile-name"] = self.IntfProfileName
    leafs["intf-lease-limit-type"] = self.IntfLeaseLimitType
    leafs["intf-lease-limit-count"] = self.IntfLeaseLimitCount
    leafs["srg-role"] = self.SrgRole
    leafs["mac-throttle"] = self.MacThrottle
    return leafs
}

func (self *Ipv4Dhcpd_Nodes_Node_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *Ipv4Dhcpd_Nodes_Node_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *Ipv4Dhcpd_Nodes_Node_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Ipv4Dhcpd_Nodes_Node_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Ipv4Dhcpd_Nodes_Node_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Ipv4Dhcpd_Nodes_Node_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *Ipv4Dhcpd_Nodes_Node_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *Ipv4Dhcpd_Nodes_Node_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// Ipv4Dhcpd_Nodes_Node_Base
// IPv4 DHCP base operational data
type Ipv4Dhcpd_Nodes_Node_Base struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DHCP base statistics.
    Statistics Ipv4Dhcpd_Nodes_Node_Base_Statistics

    // IPv4 DHCP ISSU status.
    IssuStatus Ipv4Dhcpd_Nodes_Node_Base_IssuStatus

    // DHCP base list of VRF names.
    Vrfs Ipv4Dhcpd_Nodes_Node_Base_Vrfs

    // IPv4 DHCP Base profile.
    Profiles Ipv4Dhcpd_Nodes_Node_Base_Profiles

    // IPv4 DHCP database.
    Database Ipv4Dhcpd_Nodes_Node_Base_Database
}

func (base *Ipv4Dhcpd_Nodes_Node_Base) GetFilter() yfilter.YFilter { return base.YFilter }

func (base *Ipv4Dhcpd_Nodes_Node_Base) SetFilter(yf yfilter.YFilter) { base.YFilter = yf }

func (base *Ipv4Dhcpd_Nodes_Node_Base) GetGoName(yname string) string {
    if yname == "statistics" { return "Statistics" }
    if yname == "issu-status" { return "IssuStatus" }
    if yname == "vrfs" { return "Vrfs" }
    if yname == "profiles" { return "Profiles" }
    if yname == "database" { return "Database" }
    return ""
}

func (base *Ipv4Dhcpd_Nodes_Node_Base) GetSegmentPath() string {
    return "base"
}

func (base *Ipv4Dhcpd_Nodes_Node_Base) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "statistics" {
        return &base.Statistics
    }
    if childYangName == "issu-status" {
        return &base.IssuStatus
    }
    if childYangName == "vrfs" {
        return &base.Vrfs
    }
    if childYangName == "profiles" {
        return &base.Profiles
    }
    if childYangName == "database" {
        return &base.Database
    }
    return nil
}

func (base *Ipv4Dhcpd_Nodes_Node_Base) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["statistics"] = &base.Statistics
    children["issu-status"] = &base.IssuStatus
    children["vrfs"] = &base.Vrfs
    children["profiles"] = &base.Profiles
    children["database"] = &base.Database
    return children
}

func (base *Ipv4Dhcpd_Nodes_Node_Base) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (base *Ipv4Dhcpd_Nodes_Node_Base) GetBundleName() string { return "cisco_ios_xr" }

func (base *Ipv4Dhcpd_Nodes_Node_Base) GetYangName() string { return "base" }

func (base *Ipv4Dhcpd_Nodes_Node_Base) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (base *Ipv4Dhcpd_Nodes_Node_Base) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (base *Ipv4Dhcpd_Nodes_Node_Base) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (base *Ipv4Dhcpd_Nodes_Node_Base) SetParent(parent types.Entity) { base.parent = parent }

func (base *Ipv4Dhcpd_Nodes_Node_Base) GetParent() types.Entity { return base.parent }

func (base *Ipv4Dhcpd_Nodes_Node_Base) GetParentYangName() string { return "node" }

// Ipv4Dhcpd_Nodes_Node_Base_Statistics
// DHCP base statistics
type Ipv4Dhcpd_Nodes_Node_Base_Statistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ipv4 dhcpd proxy stat. The type is slice of
    // Ipv4Dhcpd_Nodes_Node_Base_Statistics_Ipv4DhcpdProxyStat.
    Ipv4DhcpdProxyStat []Ipv4Dhcpd_Nodes_Node_Base_Statistics_Ipv4DhcpdProxyStat
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Base_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *Ipv4Dhcpd_Nodes_Node_Base_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *Ipv4Dhcpd_Nodes_Node_Base_Statistics) GetGoName(yname string) string {
    if yname == "ipv4-dhcpd-proxy-stat" { return "Ipv4DhcpdProxyStat" }
    return ""
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Base_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Base_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv4-dhcpd-proxy-stat" {
        for _, c := range statistics.Ipv4DhcpdProxyStat {
            if statistics.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4Dhcpd_Nodes_Node_Base_Statistics_Ipv4DhcpdProxyStat{}
        statistics.Ipv4DhcpdProxyStat = append(statistics.Ipv4DhcpdProxyStat, child)
        return &statistics.Ipv4DhcpdProxyStat[len(statistics.Ipv4DhcpdProxyStat)-1]
    }
    return nil
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Base_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range statistics.Ipv4DhcpdProxyStat {
        children[statistics.Ipv4DhcpdProxyStat[i].GetSegmentPath()] = &statistics.Ipv4DhcpdProxyStat[i]
    }
    return children
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Base_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Base_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *Ipv4Dhcpd_Nodes_Node_Base_Statistics) GetYangName() string { return "statistics" }

func (statistics *Ipv4Dhcpd_Nodes_Node_Base_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *Ipv4Dhcpd_Nodes_Node_Base_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *Ipv4Dhcpd_Nodes_Node_Base_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *Ipv4Dhcpd_Nodes_Node_Base_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *Ipv4Dhcpd_Nodes_Node_Base_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *Ipv4Dhcpd_Nodes_Node_Base_Statistics) GetParentYangName() string { return "base" }

// Ipv4Dhcpd_Nodes_Node_Base_Statistics_Ipv4DhcpdProxyStat
// ipv4 dhcpd proxy stat
type Ipv4Dhcpd_Nodes_Node_Base_Statistics_Ipv4DhcpdProxyStat struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DHCP L3 VRF name. The type is string with length: 0..33.
    VrfName interface{}

    // Proxy statistics.
    Statistics Ipv4Dhcpd_Nodes_Node_Base_Statistics_Ipv4DhcpdProxyStat_Statistics
}

func (ipv4DhcpdProxyStat *Ipv4Dhcpd_Nodes_Node_Base_Statistics_Ipv4DhcpdProxyStat) GetFilter() yfilter.YFilter { return ipv4DhcpdProxyStat.YFilter }

func (ipv4DhcpdProxyStat *Ipv4Dhcpd_Nodes_Node_Base_Statistics_Ipv4DhcpdProxyStat) SetFilter(yf yfilter.YFilter) { ipv4DhcpdProxyStat.YFilter = yf }

func (ipv4DhcpdProxyStat *Ipv4Dhcpd_Nodes_Node_Base_Statistics_Ipv4DhcpdProxyStat) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "statistics" { return "Statistics" }
    return ""
}

func (ipv4DhcpdProxyStat *Ipv4Dhcpd_Nodes_Node_Base_Statistics_Ipv4DhcpdProxyStat) GetSegmentPath() string {
    return "ipv4-dhcpd-proxy-stat"
}

func (ipv4DhcpdProxyStat *Ipv4Dhcpd_Nodes_Node_Base_Statistics_Ipv4DhcpdProxyStat) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "statistics" {
        return &ipv4DhcpdProxyStat.Statistics
    }
    return nil
}

func (ipv4DhcpdProxyStat *Ipv4Dhcpd_Nodes_Node_Base_Statistics_Ipv4DhcpdProxyStat) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["statistics"] = &ipv4DhcpdProxyStat.Statistics
    return children
}

func (ipv4DhcpdProxyStat *Ipv4Dhcpd_Nodes_Node_Base_Statistics_Ipv4DhcpdProxyStat) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = ipv4DhcpdProxyStat.VrfName
    return leafs
}

func (ipv4DhcpdProxyStat *Ipv4Dhcpd_Nodes_Node_Base_Statistics_Ipv4DhcpdProxyStat) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4DhcpdProxyStat *Ipv4Dhcpd_Nodes_Node_Base_Statistics_Ipv4DhcpdProxyStat) GetYangName() string { return "ipv4-dhcpd-proxy-stat" }

func (ipv4DhcpdProxyStat *Ipv4Dhcpd_Nodes_Node_Base_Statistics_Ipv4DhcpdProxyStat) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4DhcpdProxyStat *Ipv4Dhcpd_Nodes_Node_Base_Statistics_Ipv4DhcpdProxyStat) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4DhcpdProxyStat *Ipv4Dhcpd_Nodes_Node_Base_Statistics_Ipv4DhcpdProxyStat) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4DhcpdProxyStat *Ipv4Dhcpd_Nodes_Node_Base_Statistics_Ipv4DhcpdProxyStat) SetParent(parent types.Entity) { ipv4DhcpdProxyStat.parent = parent }

func (ipv4DhcpdProxyStat *Ipv4Dhcpd_Nodes_Node_Base_Statistics_Ipv4DhcpdProxyStat) GetParent() types.Entity { return ipv4DhcpdProxyStat.parent }

func (ipv4DhcpdProxyStat *Ipv4Dhcpd_Nodes_Node_Base_Statistics_Ipv4DhcpdProxyStat) GetParentYangName() string { return "statistics" }

// Ipv4Dhcpd_Nodes_Node_Base_Statistics_Ipv4DhcpdProxyStat_Statistics
// Proxy statistics
type Ipv4Dhcpd_Nodes_Node_Base_Statistics_Ipv4DhcpdProxyStat_Statistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Received packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPackets interface{}

    // Transmitted packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedPackets interface{}

    // Dropped packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Base_Statistics_Ipv4DhcpdProxyStat_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *Ipv4Dhcpd_Nodes_Node_Base_Statistics_Ipv4DhcpdProxyStat_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *Ipv4Dhcpd_Nodes_Node_Base_Statistics_Ipv4DhcpdProxyStat_Statistics) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Base_Statistics_Ipv4DhcpdProxyStat_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Base_Statistics_Ipv4DhcpdProxyStat_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Base_Statistics_Ipv4DhcpdProxyStat_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Base_Statistics_Ipv4DhcpdProxyStat_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = statistics.ReceivedPackets
    leafs["transmitted-packets"] = statistics.TransmittedPackets
    leafs["dropped-packets"] = statistics.DroppedPackets
    return leafs
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Base_Statistics_Ipv4DhcpdProxyStat_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *Ipv4Dhcpd_Nodes_Node_Base_Statistics_Ipv4DhcpdProxyStat_Statistics) GetYangName() string { return "statistics" }

func (statistics *Ipv4Dhcpd_Nodes_Node_Base_Statistics_Ipv4DhcpdProxyStat_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *Ipv4Dhcpd_Nodes_Node_Base_Statistics_Ipv4DhcpdProxyStat_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *Ipv4Dhcpd_Nodes_Node_Base_Statistics_Ipv4DhcpdProxyStat_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *Ipv4Dhcpd_Nodes_Node_Base_Statistics_Ipv4DhcpdProxyStat_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *Ipv4Dhcpd_Nodes_Node_Base_Statistics_Ipv4DhcpdProxyStat_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *Ipv4Dhcpd_Nodes_Node_Base_Statistics_Ipv4DhcpdProxyStat_Statistics) GetParentYangName() string { return "ipv4-dhcpd-proxy-stat" }

// Ipv4Dhcpd_Nodes_Node_Base_IssuStatus
// IPv4 DHCP ISSU status
type Ipv4Dhcpd_Nodes_Node_Base_IssuStatus struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Timestamp for the ISSU sync complete in nanoseconds since Epoch, i.e. since
    // 00:00:00 UTC , January 1, 1970. The type is interface{} with range:
    // 0..18446744073709551615. Units are nanosecond.
    IssuSyncCompleteTime interface{}

    // Timestamp for the ISSU sync start in nanoseconds since Epoch, i.e. since
    // 00:00:00 UTC, January 1, 1970. The type is interface{} with range:
    // 0..18446744073709551615. Units are nanosecond.
    IssuSyncStartTime interface{}

    // Timestamp for the ISSU ready declaration in nanoseconds since Epoch, i.e.
    // since 00:00:00 UTC , January 1, 1970. The type is interface{} with range:
    // 0..18446744073709551615. Units are nanosecond.
    IssuReadyTime interface{}

    // Timestamp for the Big Bang notification time in nanoseconds since Epoch,
    // i.e. since 00:00:00 UTC , January 1, 1970. The type is interface{} with
    // range: 0..18446744073709551615. Units are nanosecond.
    BigBangTime interface{}

    // Timestamp for the change to Primary role notification time in nanoseconds
    // since Epoch, i .e. since 00:00:00 UTC, January 1, 1970. The type is
    // interface{} with range: 0..18446744073709551615. Units are nanosecond.
    PrimaryRoleTime interface{}

    // The current role of the DHCP process. The type is DhcpIssuRole.
    Role interface{}

    // The current ISSU phase of the DHCP process. The type is DhcpIssuPhase.
    Phase interface{}

    // The current version of the DHCP process in the context of an ISSU. The type
    // is DhcpIssuVersion.
    Version interface{}

    // Whether or not DHCP is currently connected to ISSU Manager during the ISSU
    // Load Phase. The type is bool.
    IssuReadyIssuMgrConnection interface{}

    // Whether or not DHCP has received all replicated entries during the ISSU
    // Load Phase. The type is bool.
    IssuReadyEntriesReplicate interface{}
}

func (issuStatus *Ipv4Dhcpd_Nodes_Node_Base_IssuStatus) GetFilter() yfilter.YFilter { return issuStatus.YFilter }

func (issuStatus *Ipv4Dhcpd_Nodes_Node_Base_IssuStatus) SetFilter(yf yfilter.YFilter) { issuStatus.YFilter = yf }

func (issuStatus *Ipv4Dhcpd_Nodes_Node_Base_IssuStatus) GetGoName(yname string) string {
    if yname == "issu-sync-complete-time" { return "IssuSyncCompleteTime" }
    if yname == "issu-sync-start-time" { return "IssuSyncStartTime" }
    if yname == "issu-ready-time" { return "IssuReadyTime" }
    if yname == "big-bang-time" { return "BigBangTime" }
    if yname == "primary-role-time" { return "PrimaryRoleTime" }
    if yname == "role" { return "Role" }
    if yname == "phase" { return "Phase" }
    if yname == "version" { return "Version" }
    if yname == "issu-ready-issu-mgr-connection" { return "IssuReadyIssuMgrConnection" }
    if yname == "issu-ready-entries-replicate" { return "IssuReadyEntriesReplicate" }
    return ""
}

func (issuStatus *Ipv4Dhcpd_Nodes_Node_Base_IssuStatus) GetSegmentPath() string {
    return "issu-status"
}

func (issuStatus *Ipv4Dhcpd_Nodes_Node_Base_IssuStatus) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (issuStatus *Ipv4Dhcpd_Nodes_Node_Base_IssuStatus) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (issuStatus *Ipv4Dhcpd_Nodes_Node_Base_IssuStatus) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["issu-sync-complete-time"] = issuStatus.IssuSyncCompleteTime
    leafs["issu-sync-start-time"] = issuStatus.IssuSyncStartTime
    leafs["issu-ready-time"] = issuStatus.IssuReadyTime
    leafs["big-bang-time"] = issuStatus.BigBangTime
    leafs["primary-role-time"] = issuStatus.PrimaryRoleTime
    leafs["role"] = issuStatus.Role
    leafs["phase"] = issuStatus.Phase
    leafs["version"] = issuStatus.Version
    leafs["issu-ready-issu-mgr-connection"] = issuStatus.IssuReadyIssuMgrConnection
    leafs["issu-ready-entries-replicate"] = issuStatus.IssuReadyEntriesReplicate
    return leafs
}

func (issuStatus *Ipv4Dhcpd_Nodes_Node_Base_IssuStatus) GetBundleName() string { return "cisco_ios_xr" }

func (issuStatus *Ipv4Dhcpd_Nodes_Node_Base_IssuStatus) GetYangName() string { return "issu-status" }

func (issuStatus *Ipv4Dhcpd_Nodes_Node_Base_IssuStatus) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (issuStatus *Ipv4Dhcpd_Nodes_Node_Base_IssuStatus) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (issuStatus *Ipv4Dhcpd_Nodes_Node_Base_IssuStatus) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (issuStatus *Ipv4Dhcpd_Nodes_Node_Base_IssuStatus) SetParent(parent types.Entity) { issuStatus.parent = parent }

func (issuStatus *Ipv4Dhcpd_Nodes_Node_Base_IssuStatus) GetParent() types.Entity { return issuStatus.parent }

func (issuStatus *Ipv4Dhcpd_Nodes_Node_Base_IssuStatus) GetParentYangName() string { return "base" }

// Ipv4Dhcpd_Nodes_Node_Base_Vrfs
// DHCP base list of VRF names
type Ipv4Dhcpd_Nodes_Node_Base_Vrfs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 DHCP base VRF name. The type is slice of
    // Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf.
    Vrf []Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf
}

func (vrfs *Ipv4Dhcpd_Nodes_Node_Base_Vrfs) GetFilter() yfilter.YFilter { return vrfs.YFilter }

func (vrfs *Ipv4Dhcpd_Nodes_Node_Base_Vrfs) SetFilter(yf yfilter.YFilter) { vrfs.YFilter = yf }

func (vrfs *Ipv4Dhcpd_Nodes_Node_Base_Vrfs) GetGoName(yname string) string {
    if yname == "vrf" { return "Vrf" }
    return ""
}

func (vrfs *Ipv4Dhcpd_Nodes_Node_Base_Vrfs) GetSegmentPath() string {
    return "vrfs"
}

func (vrfs *Ipv4Dhcpd_Nodes_Node_Base_Vrfs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf" {
        for _, c := range vrfs.Vrf {
            if vrfs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf{}
        vrfs.Vrf = append(vrfs.Vrf, child)
        return &vrfs.Vrf[len(vrfs.Vrf)-1]
    }
    return nil
}

func (vrfs *Ipv4Dhcpd_Nodes_Node_Base_Vrfs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vrfs.Vrf {
        children[vrfs.Vrf[i].GetSegmentPath()] = &vrfs.Vrf[i]
    }
    return children
}

func (vrfs *Ipv4Dhcpd_Nodes_Node_Base_Vrfs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vrfs *Ipv4Dhcpd_Nodes_Node_Base_Vrfs) GetBundleName() string { return "cisco_ios_xr" }

func (vrfs *Ipv4Dhcpd_Nodes_Node_Base_Vrfs) GetYangName() string { return "vrfs" }

func (vrfs *Ipv4Dhcpd_Nodes_Node_Base_Vrfs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfs *Ipv4Dhcpd_Nodes_Node_Base_Vrfs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfs *Ipv4Dhcpd_Nodes_Node_Base_Vrfs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfs *Ipv4Dhcpd_Nodes_Node_Base_Vrfs) SetParent(parent types.Entity) { vrfs.parent = parent }

func (vrfs *Ipv4Dhcpd_Nodes_Node_Base_Vrfs) GetParent() types.Entity { return vrfs.parent }

func (vrfs *Ipv4Dhcpd_Nodes_Node_Base_Vrfs) GetParentYangName() string { return "base" }

// Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf
// IPv4 DHCP base VRF name
type Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // IPv4 DHCP base statistics.
    Statistics Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics
}

func (vrf *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf) GetFilter() yfilter.YFilter { return vrf.YFilter }

func (vrf *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf) SetFilter(yf yfilter.YFilter) { vrf.YFilter = yf }

func (vrf *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "statistics" { return "Statistics" }
    return ""
}

func (vrf *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf) GetSegmentPath() string {
    return "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
}

func (vrf *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "statistics" {
        return &vrf.Statistics
    }
    return nil
}

func (vrf *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["statistics"] = &vrf.Statistics
    return children
}

func (vrf *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = vrf.VrfName
    return leafs
}

func (vrf *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf) GetBundleName() string { return "cisco_ios_xr" }

func (vrf *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf) GetYangName() string { return "vrf" }

func (vrf *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrf *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrf *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrf *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf) SetParent(parent types.Entity) { vrf.parent = parent }

func (vrf *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf) GetParent() types.Entity { return vrf.parent }

func (vrf *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf) GetParentYangName() string { return "vrfs" }

// Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics
// IPv4 DHCP base statistics
type Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DHCP discover packets.
    Discover Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Discover

    // DHCP offer packets.
    Offer Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Offer

    // DHCP request packets.
    Request Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Request

    // DHCP decline packets.
    Decline Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Decline

    // DHCP ack packets.
    Ack Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Ack

    // DHCP nak packets.
    Nak Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Nak

    // DHCP release packets.
    Release Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Release

    // DHCP inform packets.
    Inform Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Inform

    // DHCP lease query packets.
    LeaseQuery Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseQuery

    // DHCP lease not assigned.
    LeaseNotAssigned Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseNotAssigned

    // DHCP lease unknown.
    LeaseUnknown Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseUnknown

    // DHCP lease active.
    LeaseActive Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseActive

    // DHCP BOOTP request.
    BootpRequest Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_BootpRequest

    // DHCP BOOTP reply.
    BootpReply Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_BootpReply
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics) GetGoName(yname string) string {
    if yname == "discover" { return "Discover" }
    if yname == "offer" { return "Offer" }
    if yname == "request" { return "Request" }
    if yname == "decline" { return "Decline" }
    if yname == "ack" { return "Ack" }
    if yname == "nak" { return "Nak" }
    if yname == "release" { return "Release" }
    if yname == "inform" { return "Inform" }
    if yname == "lease-query" { return "LeaseQuery" }
    if yname == "lease-not-assigned" { return "LeaseNotAssigned" }
    if yname == "lease-unknown" { return "LeaseUnknown" }
    if yname == "lease-active" { return "LeaseActive" }
    if yname == "bootp-request" { return "BootpRequest" }
    if yname == "bootp-reply" { return "BootpReply" }
    return ""
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "discover" {
        return &statistics.Discover
    }
    if childYangName == "offer" {
        return &statistics.Offer
    }
    if childYangName == "request" {
        return &statistics.Request
    }
    if childYangName == "decline" {
        return &statistics.Decline
    }
    if childYangName == "ack" {
        return &statistics.Ack
    }
    if childYangName == "nak" {
        return &statistics.Nak
    }
    if childYangName == "release" {
        return &statistics.Release
    }
    if childYangName == "inform" {
        return &statistics.Inform
    }
    if childYangName == "lease-query" {
        return &statistics.LeaseQuery
    }
    if childYangName == "lease-not-assigned" {
        return &statistics.LeaseNotAssigned
    }
    if childYangName == "lease-unknown" {
        return &statistics.LeaseUnknown
    }
    if childYangName == "lease-active" {
        return &statistics.LeaseActive
    }
    if childYangName == "bootp-request" {
        return &statistics.BootpRequest
    }
    if childYangName == "bootp-reply" {
        return &statistics.BootpReply
    }
    return nil
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["discover"] = &statistics.Discover
    children["offer"] = &statistics.Offer
    children["request"] = &statistics.Request
    children["decline"] = &statistics.Decline
    children["ack"] = &statistics.Ack
    children["nak"] = &statistics.Nak
    children["release"] = &statistics.Release
    children["inform"] = &statistics.Inform
    children["lease-query"] = &statistics.LeaseQuery
    children["lease-not-assigned"] = &statistics.LeaseNotAssigned
    children["lease-unknown"] = &statistics.LeaseUnknown
    children["lease-active"] = &statistics.LeaseActive
    children["bootp-request"] = &statistics.BootpRequest
    children["bootp-reply"] = &statistics.BootpReply
    return children
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics) GetYangName() string { return "statistics" }

func (statistics *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics) GetParentYangName() string { return "vrf" }

// Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Discover
// DHCP discover packets
type Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Discover struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Received packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPackets interface{}

    // Transmitted packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedPackets interface{}

    // Dropped packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}
}

func (discover *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Discover) GetFilter() yfilter.YFilter { return discover.YFilter }

func (discover *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Discover) SetFilter(yf yfilter.YFilter) { discover.YFilter = yf }

func (discover *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Discover) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (discover *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Discover) GetSegmentPath() string {
    return "discover"
}

func (discover *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Discover) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (discover *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Discover) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (discover *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Discover) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = discover.ReceivedPackets
    leafs["transmitted-packets"] = discover.TransmittedPackets
    leafs["dropped-packets"] = discover.DroppedPackets
    return leafs
}

func (discover *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Discover) GetBundleName() string { return "cisco_ios_xr" }

func (discover *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Discover) GetYangName() string { return "discover" }

func (discover *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Discover) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (discover *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Discover) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (discover *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Discover) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (discover *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Discover) SetParent(parent types.Entity) { discover.parent = parent }

func (discover *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Discover) GetParent() types.Entity { return discover.parent }

func (discover *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Discover) GetParentYangName() string { return "statistics" }

// Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Offer
// DHCP offer packets
type Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Offer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Received packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPackets interface{}

    // Transmitted packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedPackets interface{}

    // Dropped packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}
}

func (offer *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Offer) GetFilter() yfilter.YFilter { return offer.YFilter }

func (offer *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Offer) SetFilter(yf yfilter.YFilter) { offer.YFilter = yf }

func (offer *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Offer) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (offer *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Offer) GetSegmentPath() string {
    return "offer"
}

func (offer *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Offer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (offer *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Offer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (offer *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Offer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = offer.ReceivedPackets
    leafs["transmitted-packets"] = offer.TransmittedPackets
    leafs["dropped-packets"] = offer.DroppedPackets
    return leafs
}

func (offer *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Offer) GetBundleName() string { return "cisco_ios_xr" }

func (offer *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Offer) GetYangName() string { return "offer" }

func (offer *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Offer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (offer *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Offer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (offer *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Offer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (offer *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Offer) SetParent(parent types.Entity) { offer.parent = parent }

func (offer *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Offer) GetParent() types.Entity { return offer.parent }

func (offer *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Offer) GetParentYangName() string { return "statistics" }

// Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Request
// DHCP request packets
type Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Request struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Received packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPackets interface{}

    // Transmitted packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedPackets interface{}

    // Dropped packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}
}

func (request *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Request) GetFilter() yfilter.YFilter { return request.YFilter }

func (request *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Request) SetFilter(yf yfilter.YFilter) { request.YFilter = yf }

func (request *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Request) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (request *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Request) GetSegmentPath() string {
    return "request"
}

func (request *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Request) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (request *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Request) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (request *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Request) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = request.ReceivedPackets
    leafs["transmitted-packets"] = request.TransmittedPackets
    leafs["dropped-packets"] = request.DroppedPackets
    return leafs
}

func (request *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Request) GetBundleName() string { return "cisco_ios_xr" }

func (request *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Request) GetYangName() string { return "request" }

func (request *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Request) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (request *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Request) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (request *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Request) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (request *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Request) SetParent(parent types.Entity) { request.parent = parent }

func (request *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Request) GetParent() types.Entity { return request.parent }

func (request *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Request) GetParentYangName() string { return "statistics" }

// Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Decline
// DHCP decline packets
type Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Decline struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Received packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPackets interface{}

    // Transmitted packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedPackets interface{}

    // Dropped packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}
}

func (decline *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Decline) GetFilter() yfilter.YFilter { return decline.YFilter }

func (decline *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Decline) SetFilter(yf yfilter.YFilter) { decline.YFilter = yf }

func (decline *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Decline) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (decline *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Decline) GetSegmentPath() string {
    return "decline"
}

func (decline *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Decline) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (decline *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Decline) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (decline *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Decline) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = decline.ReceivedPackets
    leafs["transmitted-packets"] = decline.TransmittedPackets
    leafs["dropped-packets"] = decline.DroppedPackets
    return leafs
}

func (decline *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Decline) GetBundleName() string { return "cisco_ios_xr" }

func (decline *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Decline) GetYangName() string { return "decline" }

func (decline *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Decline) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (decline *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Decline) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (decline *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Decline) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (decline *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Decline) SetParent(parent types.Entity) { decline.parent = parent }

func (decline *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Decline) GetParent() types.Entity { return decline.parent }

func (decline *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Decline) GetParentYangName() string { return "statistics" }

// Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Ack
// DHCP ack packets
type Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Ack struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Received packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPackets interface{}

    // Transmitted packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedPackets interface{}

    // Dropped packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}
}

func (ack *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Ack) GetFilter() yfilter.YFilter { return ack.YFilter }

func (ack *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Ack) SetFilter(yf yfilter.YFilter) { ack.YFilter = yf }

func (ack *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Ack) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (ack *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Ack) GetSegmentPath() string {
    return "ack"
}

func (ack *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Ack) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ack *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Ack) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ack *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Ack) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = ack.ReceivedPackets
    leafs["transmitted-packets"] = ack.TransmittedPackets
    leafs["dropped-packets"] = ack.DroppedPackets
    return leafs
}

func (ack *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Ack) GetBundleName() string { return "cisco_ios_xr" }

func (ack *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Ack) GetYangName() string { return "ack" }

func (ack *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Ack) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ack *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Ack) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ack *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Ack) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ack *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Ack) SetParent(parent types.Entity) { ack.parent = parent }

func (ack *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Ack) GetParent() types.Entity { return ack.parent }

func (ack *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Ack) GetParentYangName() string { return "statistics" }

// Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Nak
// DHCP nak packets
type Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Nak struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Received packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPackets interface{}

    // Transmitted packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedPackets interface{}

    // Dropped packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}
}

func (nak *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Nak) GetFilter() yfilter.YFilter { return nak.YFilter }

func (nak *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Nak) SetFilter(yf yfilter.YFilter) { nak.YFilter = yf }

func (nak *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Nak) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (nak *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Nak) GetSegmentPath() string {
    return "nak"
}

func (nak *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Nak) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nak *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Nak) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nak *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Nak) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = nak.ReceivedPackets
    leafs["transmitted-packets"] = nak.TransmittedPackets
    leafs["dropped-packets"] = nak.DroppedPackets
    return leafs
}

func (nak *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Nak) GetBundleName() string { return "cisco_ios_xr" }

func (nak *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Nak) GetYangName() string { return "nak" }

func (nak *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Nak) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nak *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Nak) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nak *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Nak) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nak *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Nak) SetParent(parent types.Entity) { nak.parent = parent }

func (nak *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Nak) GetParent() types.Entity { return nak.parent }

func (nak *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Nak) GetParentYangName() string { return "statistics" }

// Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Release
// DHCP release packets
type Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Release struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Received packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPackets interface{}

    // Transmitted packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedPackets interface{}

    // Dropped packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}
}

func (release *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Release) GetFilter() yfilter.YFilter { return release.YFilter }

func (release *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Release) SetFilter(yf yfilter.YFilter) { release.YFilter = yf }

func (release *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Release) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (release *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Release) GetSegmentPath() string {
    return "release"
}

func (release *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Release) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (release *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Release) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (release *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Release) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = release.ReceivedPackets
    leafs["transmitted-packets"] = release.TransmittedPackets
    leafs["dropped-packets"] = release.DroppedPackets
    return leafs
}

func (release *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Release) GetBundleName() string { return "cisco_ios_xr" }

func (release *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Release) GetYangName() string { return "release" }

func (release *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Release) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (release *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Release) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (release *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Release) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (release *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Release) SetParent(parent types.Entity) { release.parent = parent }

func (release *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Release) GetParent() types.Entity { return release.parent }

func (release *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Release) GetParentYangName() string { return "statistics" }

// Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Inform
// DHCP inform packets
type Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Inform struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Received packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPackets interface{}

    // Transmitted packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedPackets interface{}

    // Dropped packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}
}

func (inform *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Inform) GetFilter() yfilter.YFilter { return inform.YFilter }

func (inform *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Inform) SetFilter(yf yfilter.YFilter) { inform.YFilter = yf }

func (inform *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Inform) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (inform *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Inform) GetSegmentPath() string {
    return "inform"
}

func (inform *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Inform) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inform *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Inform) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inform *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Inform) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = inform.ReceivedPackets
    leafs["transmitted-packets"] = inform.TransmittedPackets
    leafs["dropped-packets"] = inform.DroppedPackets
    return leafs
}

func (inform *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Inform) GetBundleName() string { return "cisco_ios_xr" }

func (inform *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Inform) GetYangName() string { return "inform" }

func (inform *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Inform) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inform *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Inform) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inform *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Inform) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inform *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Inform) SetParent(parent types.Entity) { inform.parent = parent }

func (inform *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Inform) GetParent() types.Entity { return inform.parent }

func (inform *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Inform) GetParentYangName() string { return "statistics" }

// Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseQuery
// DHCP lease query packets
type Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseQuery struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Received packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPackets interface{}

    // Transmitted packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedPackets interface{}

    // Dropped packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}
}

func (leaseQuery *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseQuery) GetFilter() yfilter.YFilter { return leaseQuery.YFilter }

func (leaseQuery *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseQuery) SetFilter(yf yfilter.YFilter) { leaseQuery.YFilter = yf }

func (leaseQuery *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseQuery) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (leaseQuery *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseQuery) GetSegmentPath() string {
    return "lease-query"
}

func (leaseQuery *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseQuery) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (leaseQuery *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseQuery) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (leaseQuery *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseQuery) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = leaseQuery.ReceivedPackets
    leafs["transmitted-packets"] = leaseQuery.TransmittedPackets
    leafs["dropped-packets"] = leaseQuery.DroppedPackets
    return leafs
}

func (leaseQuery *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseQuery) GetBundleName() string { return "cisco_ios_xr" }

func (leaseQuery *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseQuery) GetYangName() string { return "lease-query" }

func (leaseQuery *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseQuery) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (leaseQuery *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseQuery) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (leaseQuery *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseQuery) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (leaseQuery *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseQuery) SetParent(parent types.Entity) { leaseQuery.parent = parent }

func (leaseQuery *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseQuery) GetParent() types.Entity { return leaseQuery.parent }

func (leaseQuery *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseQuery) GetParentYangName() string { return "statistics" }

// Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseNotAssigned
// DHCP lease not assigned
type Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseNotAssigned struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Received packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPackets interface{}

    // Transmitted packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedPackets interface{}

    // Dropped packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}
}

func (leaseNotAssigned *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseNotAssigned) GetFilter() yfilter.YFilter { return leaseNotAssigned.YFilter }

func (leaseNotAssigned *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseNotAssigned) SetFilter(yf yfilter.YFilter) { leaseNotAssigned.YFilter = yf }

func (leaseNotAssigned *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseNotAssigned) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (leaseNotAssigned *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseNotAssigned) GetSegmentPath() string {
    return "lease-not-assigned"
}

func (leaseNotAssigned *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseNotAssigned) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (leaseNotAssigned *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseNotAssigned) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (leaseNotAssigned *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseNotAssigned) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = leaseNotAssigned.ReceivedPackets
    leafs["transmitted-packets"] = leaseNotAssigned.TransmittedPackets
    leafs["dropped-packets"] = leaseNotAssigned.DroppedPackets
    return leafs
}

func (leaseNotAssigned *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseNotAssigned) GetBundleName() string { return "cisco_ios_xr" }

func (leaseNotAssigned *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseNotAssigned) GetYangName() string { return "lease-not-assigned" }

func (leaseNotAssigned *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseNotAssigned) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (leaseNotAssigned *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseNotAssigned) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (leaseNotAssigned *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseNotAssigned) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (leaseNotAssigned *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseNotAssigned) SetParent(parent types.Entity) { leaseNotAssigned.parent = parent }

func (leaseNotAssigned *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseNotAssigned) GetParent() types.Entity { return leaseNotAssigned.parent }

func (leaseNotAssigned *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseNotAssigned) GetParentYangName() string { return "statistics" }

// Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseUnknown
// DHCP lease unknown
type Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseUnknown struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Received packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPackets interface{}

    // Transmitted packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedPackets interface{}

    // Dropped packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}
}

func (leaseUnknown *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseUnknown) GetFilter() yfilter.YFilter { return leaseUnknown.YFilter }

func (leaseUnknown *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseUnknown) SetFilter(yf yfilter.YFilter) { leaseUnknown.YFilter = yf }

func (leaseUnknown *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseUnknown) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (leaseUnknown *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseUnknown) GetSegmentPath() string {
    return "lease-unknown"
}

func (leaseUnknown *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseUnknown) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (leaseUnknown *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseUnknown) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (leaseUnknown *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseUnknown) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = leaseUnknown.ReceivedPackets
    leafs["transmitted-packets"] = leaseUnknown.TransmittedPackets
    leafs["dropped-packets"] = leaseUnknown.DroppedPackets
    return leafs
}

func (leaseUnknown *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseUnknown) GetBundleName() string { return "cisco_ios_xr" }

func (leaseUnknown *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseUnknown) GetYangName() string { return "lease-unknown" }

func (leaseUnknown *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseUnknown) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (leaseUnknown *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseUnknown) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (leaseUnknown *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseUnknown) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (leaseUnknown *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseUnknown) SetParent(parent types.Entity) { leaseUnknown.parent = parent }

func (leaseUnknown *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseUnknown) GetParent() types.Entity { return leaseUnknown.parent }

func (leaseUnknown *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseUnknown) GetParentYangName() string { return "statistics" }

// Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseActive
// DHCP lease active
type Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseActive struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Received packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPackets interface{}

    // Transmitted packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedPackets interface{}

    // Dropped packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}
}

func (leaseActive *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseActive) GetFilter() yfilter.YFilter { return leaseActive.YFilter }

func (leaseActive *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseActive) SetFilter(yf yfilter.YFilter) { leaseActive.YFilter = yf }

func (leaseActive *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseActive) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (leaseActive *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseActive) GetSegmentPath() string {
    return "lease-active"
}

func (leaseActive *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseActive) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (leaseActive *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseActive) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (leaseActive *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseActive) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = leaseActive.ReceivedPackets
    leafs["transmitted-packets"] = leaseActive.TransmittedPackets
    leafs["dropped-packets"] = leaseActive.DroppedPackets
    return leafs
}

func (leaseActive *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseActive) GetBundleName() string { return "cisco_ios_xr" }

func (leaseActive *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseActive) GetYangName() string { return "lease-active" }

func (leaseActive *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseActive) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (leaseActive *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseActive) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (leaseActive *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseActive) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (leaseActive *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseActive) SetParent(parent types.Entity) { leaseActive.parent = parent }

func (leaseActive *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseActive) GetParent() types.Entity { return leaseActive.parent }

func (leaseActive *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseActive) GetParentYangName() string { return "statistics" }

// Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_BootpRequest
// DHCP BOOTP request
type Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_BootpRequest struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Received packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPackets interface{}

    // Transmitted packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedPackets interface{}

    // Dropped packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}
}

func (bootpRequest *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_BootpRequest) GetFilter() yfilter.YFilter { return bootpRequest.YFilter }

func (bootpRequest *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_BootpRequest) SetFilter(yf yfilter.YFilter) { bootpRequest.YFilter = yf }

func (bootpRequest *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_BootpRequest) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (bootpRequest *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_BootpRequest) GetSegmentPath() string {
    return "bootp-request"
}

func (bootpRequest *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_BootpRequest) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bootpRequest *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_BootpRequest) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bootpRequest *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_BootpRequest) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = bootpRequest.ReceivedPackets
    leafs["transmitted-packets"] = bootpRequest.TransmittedPackets
    leafs["dropped-packets"] = bootpRequest.DroppedPackets
    return leafs
}

func (bootpRequest *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_BootpRequest) GetBundleName() string { return "cisco_ios_xr" }

func (bootpRequest *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_BootpRequest) GetYangName() string { return "bootp-request" }

func (bootpRequest *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_BootpRequest) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bootpRequest *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_BootpRequest) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bootpRequest *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_BootpRequest) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bootpRequest *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_BootpRequest) SetParent(parent types.Entity) { bootpRequest.parent = parent }

func (bootpRequest *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_BootpRequest) GetParent() types.Entity { return bootpRequest.parent }

func (bootpRequest *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_BootpRequest) GetParentYangName() string { return "statistics" }

// Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_BootpReply
// DHCP BOOTP reply
type Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_BootpReply struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Received packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPackets interface{}

    // Transmitted packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedPackets interface{}

    // Dropped packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}
}

func (bootpReply *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_BootpReply) GetFilter() yfilter.YFilter { return bootpReply.YFilter }

func (bootpReply *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_BootpReply) SetFilter(yf yfilter.YFilter) { bootpReply.YFilter = yf }

func (bootpReply *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_BootpReply) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (bootpReply *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_BootpReply) GetSegmentPath() string {
    return "bootp-reply"
}

func (bootpReply *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_BootpReply) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bootpReply *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_BootpReply) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bootpReply *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_BootpReply) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = bootpReply.ReceivedPackets
    leafs["transmitted-packets"] = bootpReply.TransmittedPackets
    leafs["dropped-packets"] = bootpReply.DroppedPackets
    return leafs
}

func (bootpReply *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_BootpReply) GetBundleName() string { return "cisco_ios_xr" }

func (bootpReply *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_BootpReply) GetYangName() string { return "bootp-reply" }

func (bootpReply *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_BootpReply) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bootpReply *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_BootpReply) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bootpReply *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_BootpReply) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bootpReply *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_BootpReply) SetParent(parent types.Entity) { bootpReply.parent = parent }

func (bootpReply *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_BootpReply) GetParent() types.Entity { return bootpReply.parent }

func (bootpReply *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_BootpReply) GetParentYangName() string { return "statistics" }

// Ipv4Dhcpd_Nodes_Node_Base_Profiles
// IPv4 DHCP Base profile
type Ipv4Dhcpd_Nodes_Node_Base_Profiles struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 DHCP base profile. The type is slice of
    // Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile.
    Profile []Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile
}

func (profiles *Ipv4Dhcpd_Nodes_Node_Base_Profiles) GetFilter() yfilter.YFilter { return profiles.YFilter }

func (profiles *Ipv4Dhcpd_Nodes_Node_Base_Profiles) SetFilter(yf yfilter.YFilter) { profiles.YFilter = yf }

func (profiles *Ipv4Dhcpd_Nodes_Node_Base_Profiles) GetGoName(yname string) string {
    if yname == "profile" { return "Profile" }
    return ""
}

func (profiles *Ipv4Dhcpd_Nodes_Node_Base_Profiles) GetSegmentPath() string {
    return "profiles"
}

func (profiles *Ipv4Dhcpd_Nodes_Node_Base_Profiles) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "profile" {
        for _, c := range profiles.Profile {
            if profiles.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile{}
        profiles.Profile = append(profiles.Profile, child)
        return &profiles.Profile[len(profiles.Profile)-1]
    }
    return nil
}

func (profiles *Ipv4Dhcpd_Nodes_Node_Base_Profiles) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range profiles.Profile {
        children[profiles.Profile[i].GetSegmentPath()] = &profiles.Profile[i]
    }
    return children
}

func (profiles *Ipv4Dhcpd_Nodes_Node_Base_Profiles) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (profiles *Ipv4Dhcpd_Nodes_Node_Base_Profiles) GetBundleName() string { return "cisco_ios_xr" }

func (profiles *Ipv4Dhcpd_Nodes_Node_Base_Profiles) GetYangName() string { return "profiles" }

func (profiles *Ipv4Dhcpd_Nodes_Node_Base_Profiles) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (profiles *Ipv4Dhcpd_Nodes_Node_Base_Profiles) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (profiles *Ipv4Dhcpd_Nodes_Node_Base_Profiles) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (profiles *Ipv4Dhcpd_Nodes_Node_Base_Profiles) SetParent(parent types.Entity) { profiles.parent = parent }

func (profiles *Ipv4Dhcpd_Nodes_Node_Base_Profiles) GetParent() types.Entity { return profiles.parent }

func (profiles *Ipv4Dhcpd_Nodes_Node_Base_Profiles) GetParentYangName() string { return "base" }

// Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile
// IPv4 DHCP base profile
type Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Profile name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    ProfileName interface{}

    // Base Default Profile name. The type is string with length: 0..65.
    BaseDefaultProfileName interface{}

    // Default Profile mode. The type is interface{} with range: 0..255.
    DefaultProfileMode interface{}

    // Relay authenticate. The type is RelayInfoAuthenticate.
    RelayAuthenticate interface{}

    // DHCP configured Remote ID. The type is string with length: 0..768.
    RemoteId interface{}

    // Child profile count. The type is interface{} with range: 0..255.
    ChildProfileCount interface{}

    // Interface reference count. The type is interface{} with range: 0..255.
    IntfRefCount interface{}

    // Interface references.
    InterfaceReferences Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_InterfaceReferences

    // child profile info.
    ChildProfileInfo Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_ChildProfileInfo
}

func (profile *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile) GetFilter() yfilter.YFilter { return profile.YFilter }

func (profile *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile) SetFilter(yf yfilter.YFilter) { profile.YFilter = yf }

func (profile *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile) GetGoName(yname string) string {
    if yname == "profile-name" { return "ProfileName" }
    if yname == "base-default-profile-name" { return "BaseDefaultProfileName" }
    if yname == "default-profile-mode" { return "DefaultProfileMode" }
    if yname == "relay-authenticate" { return "RelayAuthenticate" }
    if yname == "remote-id" { return "RemoteId" }
    if yname == "child-profile-count" { return "ChildProfileCount" }
    if yname == "intf-ref-count" { return "IntfRefCount" }
    if yname == "interface-references" { return "InterfaceReferences" }
    if yname == "child-profile-info" { return "ChildProfileInfo" }
    return ""
}

func (profile *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile) GetSegmentPath() string {
    return "profile" + "[profile-name='" + fmt.Sprintf("%v", profile.ProfileName) + "']"
}

func (profile *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-references" {
        return &profile.InterfaceReferences
    }
    if childYangName == "child-profile-info" {
        return &profile.ChildProfileInfo
    }
    return nil
}

func (profile *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interface-references"] = &profile.InterfaceReferences
    children["child-profile-info"] = &profile.ChildProfileInfo
    return children
}

func (profile *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["profile-name"] = profile.ProfileName
    leafs["base-default-profile-name"] = profile.BaseDefaultProfileName
    leafs["default-profile-mode"] = profile.DefaultProfileMode
    leafs["relay-authenticate"] = profile.RelayAuthenticate
    leafs["remote-id"] = profile.RemoteId
    leafs["child-profile-count"] = profile.ChildProfileCount
    leafs["intf-ref-count"] = profile.IntfRefCount
    return leafs
}

func (profile *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile) GetBundleName() string { return "cisco_ios_xr" }

func (profile *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile) GetYangName() string { return "profile" }

func (profile *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (profile *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (profile *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (profile *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile) SetParent(parent types.Entity) { profile.parent = parent }

func (profile *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile) GetParent() types.Entity { return profile.parent }

func (profile *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile) GetParentYangName() string { return "profiles" }

// Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_InterfaceReferences
// Interface references
type Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_InterfaceReferences struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ipv4 dhcpd base interface reference. The type is slice of
    // Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_InterfaceReferences_Ipv4DhcpdBaseInterfaceReference.
    Ipv4DhcpdBaseInterfaceReference []Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_InterfaceReferences_Ipv4DhcpdBaseInterfaceReference
}

func (interfaceReferences *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_InterfaceReferences) GetFilter() yfilter.YFilter { return interfaceReferences.YFilter }

func (interfaceReferences *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_InterfaceReferences) SetFilter(yf yfilter.YFilter) { interfaceReferences.YFilter = yf }

func (interfaceReferences *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_InterfaceReferences) GetGoName(yname string) string {
    if yname == "ipv4-dhcpd-base-interface-reference" { return "Ipv4DhcpdBaseInterfaceReference" }
    return ""
}

func (interfaceReferences *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_InterfaceReferences) GetSegmentPath() string {
    return "interface-references"
}

func (interfaceReferences *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_InterfaceReferences) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv4-dhcpd-base-interface-reference" {
        for _, c := range interfaceReferences.Ipv4DhcpdBaseInterfaceReference {
            if interfaceReferences.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_InterfaceReferences_Ipv4DhcpdBaseInterfaceReference{}
        interfaceReferences.Ipv4DhcpdBaseInterfaceReference = append(interfaceReferences.Ipv4DhcpdBaseInterfaceReference, child)
        return &interfaceReferences.Ipv4DhcpdBaseInterfaceReference[len(interfaceReferences.Ipv4DhcpdBaseInterfaceReference)-1]
    }
    return nil
}

func (interfaceReferences *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_InterfaceReferences) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaceReferences.Ipv4DhcpdBaseInterfaceReference {
        children[interfaceReferences.Ipv4DhcpdBaseInterfaceReference[i].GetSegmentPath()] = &interfaceReferences.Ipv4DhcpdBaseInterfaceReference[i]
    }
    return children
}

func (interfaceReferences *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_InterfaceReferences) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceReferences *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_InterfaceReferences) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceReferences *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_InterfaceReferences) GetYangName() string { return "interface-references" }

func (interfaceReferences *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_InterfaceReferences) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceReferences *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_InterfaceReferences) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceReferences *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_InterfaceReferences) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceReferences *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_InterfaceReferences) SetParent(parent types.Entity) { interfaceReferences.parent = parent }

func (interfaceReferences *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_InterfaceReferences) GetParent() types.Entity { return interfaceReferences.parent }

func (interfaceReferences *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_InterfaceReferences) GetParentYangName() string { return "profile" }

// Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_InterfaceReferences_Ipv4DhcpdBaseInterfaceReference
// ipv4 dhcpd base interface reference
type Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_InterfaceReferences_Ipv4DhcpdBaseInterfaceReference struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface name. The type is string with length: 0..65.
    BaseReferenceInterfaceName interface{}
}

func (ipv4DhcpdBaseInterfaceReference *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_InterfaceReferences_Ipv4DhcpdBaseInterfaceReference) GetFilter() yfilter.YFilter { return ipv4DhcpdBaseInterfaceReference.YFilter }

func (ipv4DhcpdBaseInterfaceReference *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_InterfaceReferences_Ipv4DhcpdBaseInterfaceReference) SetFilter(yf yfilter.YFilter) { ipv4DhcpdBaseInterfaceReference.YFilter = yf }

func (ipv4DhcpdBaseInterfaceReference *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_InterfaceReferences_Ipv4DhcpdBaseInterfaceReference) GetGoName(yname string) string {
    if yname == "base-reference-interface-name" { return "BaseReferenceInterfaceName" }
    return ""
}

func (ipv4DhcpdBaseInterfaceReference *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_InterfaceReferences_Ipv4DhcpdBaseInterfaceReference) GetSegmentPath() string {
    return "ipv4-dhcpd-base-interface-reference"
}

func (ipv4DhcpdBaseInterfaceReference *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_InterfaceReferences_Ipv4DhcpdBaseInterfaceReference) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv4DhcpdBaseInterfaceReference *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_InterfaceReferences_Ipv4DhcpdBaseInterfaceReference) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv4DhcpdBaseInterfaceReference *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_InterfaceReferences_Ipv4DhcpdBaseInterfaceReference) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["base-reference-interface-name"] = ipv4DhcpdBaseInterfaceReference.BaseReferenceInterfaceName
    return leafs
}

func (ipv4DhcpdBaseInterfaceReference *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_InterfaceReferences_Ipv4DhcpdBaseInterfaceReference) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4DhcpdBaseInterfaceReference *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_InterfaceReferences_Ipv4DhcpdBaseInterfaceReference) GetYangName() string { return "ipv4-dhcpd-base-interface-reference" }

func (ipv4DhcpdBaseInterfaceReference *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_InterfaceReferences_Ipv4DhcpdBaseInterfaceReference) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4DhcpdBaseInterfaceReference *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_InterfaceReferences_Ipv4DhcpdBaseInterfaceReference) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4DhcpdBaseInterfaceReference *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_InterfaceReferences_Ipv4DhcpdBaseInterfaceReference) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4DhcpdBaseInterfaceReference *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_InterfaceReferences_Ipv4DhcpdBaseInterfaceReference) SetParent(parent types.Entity) { ipv4DhcpdBaseInterfaceReference.parent = parent }

func (ipv4DhcpdBaseInterfaceReference *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_InterfaceReferences_Ipv4DhcpdBaseInterfaceReference) GetParent() types.Entity { return ipv4DhcpdBaseInterfaceReference.parent }

func (ipv4DhcpdBaseInterfaceReference *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_InterfaceReferences_Ipv4DhcpdBaseInterfaceReference) GetParentYangName() string { return "interface-references" }

// Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_ChildProfileInfo
// child profile info
type Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_ChildProfileInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ipv4 dhcpd base child profile info. The type is slice of
    // Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_ChildProfileInfo_Ipv4DhcpdBaseChildProfileInfo.
    Ipv4DhcpdBaseChildProfileInfo []Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_ChildProfileInfo_Ipv4DhcpdBaseChildProfileInfo
}

func (childProfileInfo *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_ChildProfileInfo) GetFilter() yfilter.YFilter { return childProfileInfo.YFilter }

func (childProfileInfo *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_ChildProfileInfo) SetFilter(yf yfilter.YFilter) { childProfileInfo.YFilter = yf }

func (childProfileInfo *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_ChildProfileInfo) GetGoName(yname string) string {
    if yname == "ipv4-dhcpd-base-child-profile-info" { return "Ipv4DhcpdBaseChildProfileInfo" }
    return ""
}

func (childProfileInfo *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_ChildProfileInfo) GetSegmentPath() string {
    return "child-profile-info"
}

func (childProfileInfo *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_ChildProfileInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv4-dhcpd-base-child-profile-info" {
        for _, c := range childProfileInfo.Ipv4DhcpdBaseChildProfileInfo {
            if childProfileInfo.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_ChildProfileInfo_Ipv4DhcpdBaseChildProfileInfo{}
        childProfileInfo.Ipv4DhcpdBaseChildProfileInfo = append(childProfileInfo.Ipv4DhcpdBaseChildProfileInfo, child)
        return &childProfileInfo.Ipv4DhcpdBaseChildProfileInfo[len(childProfileInfo.Ipv4DhcpdBaseChildProfileInfo)-1]
    }
    return nil
}

func (childProfileInfo *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_ChildProfileInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range childProfileInfo.Ipv4DhcpdBaseChildProfileInfo {
        children[childProfileInfo.Ipv4DhcpdBaseChildProfileInfo[i].GetSegmentPath()] = &childProfileInfo.Ipv4DhcpdBaseChildProfileInfo[i]
    }
    return children
}

func (childProfileInfo *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_ChildProfileInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (childProfileInfo *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_ChildProfileInfo) GetBundleName() string { return "cisco_ios_xr" }

func (childProfileInfo *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_ChildProfileInfo) GetYangName() string { return "child-profile-info" }

func (childProfileInfo *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_ChildProfileInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (childProfileInfo *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_ChildProfileInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (childProfileInfo *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_ChildProfileInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (childProfileInfo *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_ChildProfileInfo) SetParent(parent types.Entity) { childProfileInfo.parent = parent }

func (childProfileInfo *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_ChildProfileInfo) GetParent() types.Entity { return childProfileInfo.parent }

func (childProfileInfo *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_ChildProfileInfo) GetParentYangName() string { return "profile" }

// Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_ChildProfileInfo_Ipv4DhcpdBaseChildProfileInfo
// ipv4 dhcpd base child profile info
type Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_ChildProfileInfo_Ipv4DhcpdBaseChildProfileInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Base Child Profile name. The type is string with length: 0..65.
    BaseChildProfileName interface{}

    // Profile mode. The type is interface{} with range: 0..255.
    Mode interface{}

    // Matched option code. The type is interface{} with range: 0..255.
    MatchedOptionCode interface{}

    // Matched option len. The type is interface{} with range: 0..255.
    MatchedOptionLen interface{}

    // Matched option data. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    OptionData interface{}
}

func (ipv4DhcpdBaseChildProfileInfo *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_ChildProfileInfo_Ipv4DhcpdBaseChildProfileInfo) GetFilter() yfilter.YFilter { return ipv4DhcpdBaseChildProfileInfo.YFilter }

func (ipv4DhcpdBaseChildProfileInfo *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_ChildProfileInfo_Ipv4DhcpdBaseChildProfileInfo) SetFilter(yf yfilter.YFilter) { ipv4DhcpdBaseChildProfileInfo.YFilter = yf }

func (ipv4DhcpdBaseChildProfileInfo *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_ChildProfileInfo_Ipv4DhcpdBaseChildProfileInfo) GetGoName(yname string) string {
    if yname == "base-child-profile-name" { return "BaseChildProfileName" }
    if yname == "mode" { return "Mode" }
    if yname == "matched-option-code" { return "MatchedOptionCode" }
    if yname == "matched-option-len" { return "MatchedOptionLen" }
    if yname == "option-data" { return "OptionData" }
    return ""
}

func (ipv4DhcpdBaseChildProfileInfo *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_ChildProfileInfo_Ipv4DhcpdBaseChildProfileInfo) GetSegmentPath() string {
    return "ipv4-dhcpd-base-child-profile-info"
}

func (ipv4DhcpdBaseChildProfileInfo *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_ChildProfileInfo_Ipv4DhcpdBaseChildProfileInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv4DhcpdBaseChildProfileInfo *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_ChildProfileInfo_Ipv4DhcpdBaseChildProfileInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv4DhcpdBaseChildProfileInfo *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_ChildProfileInfo_Ipv4DhcpdBaseChildProfileInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["base-child-profile-name"] = ipv4DhcpdBaseChildProfileInfo.BaseChildProfileName
    leafs["mode"] = ipv4DhcpdBaseChildProfileInfo.Mode
    leafs["matched-option-code"] = ipv4DhcpdBaseChildProfileInfo.MatchedOptionCode
    leafs["matched-option-len"] = ipv4DhcpdBaseChildProfileInfo.MatchedOptionLen
    leafs["option-data"] = ipv4DhcpdBaseChildProfileInfo.OptionData
    return leafs
}

func (ipv4DhcpdBaseChildProfileInfo *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_ChildProfileInfo_Ipv4DhcpdBaseChildProfileInfo) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4DhcpdBaseChildProfileInfo *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_ChildProfileInfo_Ipv4DhcpdBaseChildProfileInfo) GetYangName() string { return "ipv4-dhcpd-base-child-profile-info" }

func (ipv4DhcpdBaseChildProfileInfo *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_ChildProfileInfo_Ipv4DhcpdBaseChildProfileInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4DhcpdBaseChildProfileInfo *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_ChildProfileInfo_Ipv4DhcpdBaseChildProfileInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4DhcpdBaseChildProfileInfo *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_ChildProfileInfo_Ipv4DhcpdBaseChildProfileInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4DhcpdBaseChildProfileInfo *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_ChildProfileInfo_Ipv4DhcpdBaseChildProfileInfo) SetParent(parent types.Entity) { ipv4DhcpdBaseChildProfileInfo.parent = parent }

func (ipv4DhcpdBaseChildProfileInfo *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_ChildProfileInfo_Ipv4DhcpdBaseChildProfileInfo) GetParent() types.Entity { return ipv4DhcpdBaseChildProfileInfo.parent }

func (ipv4DhcpdBaseChildProfileInfo *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_ChildProfileInfo_Ipv4DhcpdBaseChildProfileInfo) GetParentYangName() string { return "child-profile-info" }

// Ipv4Dhcpd_Nodes_Node_Base_Database
// IPv4 DHCP database
type Ipv4Dhcpd_Nodes_Node_Base_Database struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Database feature configured. The type is bool.
    Configured interface{}

    // Current file version. The type is interface{} with range: 0..4294967295.
    Version interface{}

    // Full file write interval in minutes. The type is interface{} with range:
    // 0..4294967295. Units are minute.
    FullFileWriteInterval interface{}

    // Last full write file name. The type is string with length: 0..64.
    LastFullWriteFileName interface{}

    // Last full write time since epoch. The type is interface{} with range:
    // 0..4294967295.
    LastFullWriteTime interface{}

    // Full file write count. The type is interface{} with range: 0..4294967295.
    FullFileWriteCount interface{}

    // Failed full file write count. The type is interface{} with range:
    // 0..4294967295.
    FailedFullFileWriteCount interface{}

    // Full file record count. The type is interface{} with range: 0..4294967295.
    FullFileRecordCount interface{}

    // Last full file write error timestamp since epoch. The type is interface{}
    // with range: 0..4294967295.
    LastFullFileWriteErrorTimestamp interface{}

    // Incremental file write interval in minutes. The type is interface{} with
    // range: 0..4294967295. Units are minute.
    IncrementalFileWriteInterval interface{}

    // Last incremental write file name. The type is string with length: 0..64.
    LastIncrementalWriteFileName interface{}

    // Last incremental write time since epoch. The type is interface{} with
    // range: 0..4294967295.
    LastIncrementalWriteTime interface{}

    // Incremental file write count. The type is interface{} with range:
    // 0..4294967295.
    IncrementalFileWriteCount interface{}

    // Failed incremental file write count. The type is interface{} with range:
    // 0..4294967295.
    FailedIncrementalFileWriteCount interface{}

    // Incremental file record count. The type is interface{} with range:
    // 0..4294967295.
    IncrementalFileRecordCount interface{}

    // Last incremental file write error timestamp since epoch. The type is
    // interface{} with range: 0..4294967295.
    LastIncrementalFileWriteErrorTimestamp interface{}
}

func (database *Ipv4Dhcpd_Nodes_Node_Base_Database) GetFilter() yfilter.YFilter { return database.YFilter }

func (database *Ipv4Dhcpd_Nodes_Node_Base_Database) SetFilter(yf yfilter.YFilter) { database.YFilter = yf }

func (database *Ipv4Dhcpd_Nodes_Node_Base_Database) GetGoName(yname string) string {
    if yname == "configured" { return "Configured" }
    if yname == "version" { return "Version" }
    if yname == "full-file-write-interval" { return "FullFileWriteInterval" }
    if yname == "last-full-write-file-name" { return "LastFullWriteFileName" }
    if yname == "last-full-write-time" { return "LastFullWriteTime" }
    if yname == "full-file-write-count" { return "FullFileWriteCount" }
    if yname == "failed-full-file-write-count" { return "FailedFullFileWriteCount" }
    if yname == "full-file-record-count" { return "FullFileRecordCount" }
    if yname == "last-full-file-write-error-timestamp" { return "LastFullFileWriteErrorTimestamp" }
    if yname == "incremental-file-write-interval" { return "IncrementalFileWriteInterval" }
    if yname == "last-incremental-write-file-name" { return "LastIncrementalWriteFileName" }
    if yname == "last-incremental-write-time" { return "LastIncrementalWriteTime" }
    if yname == "incremental-file-write-count" { return "IncrementalFileWriteCount" }
    if yname == "failed-incremental-file-write-count" { return "FailedIncrementalFileWriteCount" }
    if yname == "incremental-file-record-count" { return "IncrementalFileRecordCount" }
    if yname == "last-incremental-file-write-error-timestamp" { return "LastIncrementalFileWriteErrorTimestamp" }
    return ""
}

func (database *Ipv4Dhcpd_Nodes_Node_Base_Database) GetSegmentPath() string {
    return "database"
}

func (database *Ipv4Dhcpd_Nodes_Node_Base_Database) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (database *Ipv4Dhcpd_Nodes_Node_Base_Database) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (database *Ipv4Dhcpd_Nodes_Node_Base_Database) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["configured"] = database.Configured
    leafs["version"] = database.Version
    leafs["full-file-write-interval"] = database.FullFileWriteInterval
    leafs["last-full-write-file-name"] = database.LastFullWriteFileName
    leafs["last-full-write-time"] = database.LastFullWriteTime
    leafs["full-file-write-count"] = database.FullFileWriteCount
    leafs["failed-full-file-write-count"] = database.FailedFullFileWriteCount
    leafs["full-file-record-count"] = database.FullFileRecordCount
    leafs["last-full-file-write-error-timestamp"] = database.LastFullFileWriteErrorTimestamp
    leafs["incremental-file-write-interval"] = database.IncrementalFileWriteInterval
    leafs["last-incremental-write-file-name"] = database.LastIncrementalWriteFileName
    leafs["last-incremental-write-time"] = database.LastIncrementalWriteTime
    leafs["incremental-file-write-count"] = database.IncrementalFileWriteCount
    leafs["failed-incremental-file-write-count"] = database.FailedIncrementalFileWriteCount
    leafs["incremental-file-record-count"] = database.IncrementalFileRecordCount
    leafs["last-incremental-file-write-error-timestamp"] = database.LastIncrementalFileWriteErrorTimestamp
    return leafs
}

func (database *Ipv4Dhcpd_Nodes_Node_Base_Database) GetBundleName() string { return "cisco_ios_xr" }

func (database *Ipv4Dhcpd_Nodes_Node_Base_Database) GetYangName() string { return "database" }

func (database *Ipv4Dhcpd_Nodes_Node_Base_Database) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (database *Ipv4Dhcpd_Nodes_Node_Base_Database) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (database *Ipv4Dhcpd_Nodes_Node_Base_Database) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (database *Ipv4Dhcpd_Nodes_Node_Base_Database) SetParent(parent types.Entity) { database.parent = parent }

func (database *Ipv4Dhcpd_Nodes_Node_Base_Database) GetParent() types.Entity { return database.parent }

func (database *Ipv4Dhcpd_Nodes_Node_Base_Database) GetParentYangName() string { return "base" }

// Ipv4Dhcpd_Nodes_Node_Server
// IPv4 DHCP Server operational data
type Ipv4Dhcpd_Nodes_Node_Server struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 DHCP Server profile.
    Profiles Ipv4Dhcpd_Nodes_Node_Server_Profiles

    // DHCP Server statistics.
    Statistics Ipv4Dhcpd_Nodes_Node_Server_Statistics

    // DHCP server bindings.
    Binding Ipv4Dhcpd_Nodes_Node_Server_Binding

    // DHCP proxy stats info.
    StatisticsInfo Ipv4Dhcpd_Nodes_Node_Server_StatisticsInfo

    // DHCP Server list of VRF names.
    Vrfs Ipv4Dhcpd_Nodes_Node_Server_Vrfs
}

func (server *Ipv4Dhcpd_Nodes_Node_Server) GetFilter() yfilter.YFilter { return server.YFilter }

func (server *Ipv4Dhcpd_Nodes_Node_Server) SetFilter(yf yfilter.YFilter) { server.YFilter = yf }

func (server *Ipv4Dhcpd_Nodes_Node_Server) GetGoName(yname string) string {
    if yname == "profiles" { return "Profiles" }
    if yname == "statistics" { return "Statistics" }
    if yname == "binding" { return "Binding" }
    if yname == "statistics-info" { return "StatisticsInfo" }
    if yname == "vrfs" { return "Vrfs" }
    return ""
}

func (server *Ipv4Dhcpd_Nodes_Node_Server) GetSegmentPath() string {
    return "server"
}

func (server *Ipv4Dhcpd_Nodes_Node_Server) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "profiles" {
        return &server.Profiles
    }
    if childYangName == "statistics" {
        return &server.Statistics
    }
    if childYangName == "binding" {
        return &server.Binding
    }
    if childYangName == "statistics-info" {
        return &server.StatisticsInfo
    }
    if childYangName == "vrfs" {
        return &server.Vrfs
    }
    return nil
}

func (server *Ipv4Dhcpd_Nodes_Node_Server) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["profiles"] = &server.Profiles
    children["statistics"] = &server.Statistics
    children["binding"] = &server.Binding
    children["statistics-info"] = &server.StatisticsInfo
    children["vrfs"] = &server.Vrfs
    return children
}

func (server *Ipv4Dhcpd_Nodes_Node_Server) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (server *Ipv4Dhcpd_Nodes_Node_Server) GetBundleName() string { return "cisco_ios_xr" }

func (server *Ipv4Dhcpd_Nodes_Node_Server) GetYangName() string { return "server" }

func (server *Ipv4Dhcpd_Nodes_Node_Server) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (server *Ipv4Dhcpd_Nodes_Node_Server) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (server *Ipv4Dhcpd_Nodes_Node_Server) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (server *Ipv4Dhcpd_Nodes_Node_Server) SetParent(parent types.Entity) { server.parent = parent }

func (server *Ipv4Dhcpd_Nodes_Node_Server) GetParent() types.Entity { return server.parent }

func (server *Ipv4Dhcpd_Nodes_Node_Server) GetParentYangName() string { return "node" }

// Ipv4Dhcpd_Nodes_Node_Server_Profiles
// IPv4 DHCP Server profile
type Ipv4Dhcpd_Nodes_Node_Server_Profiles struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 DHCP server profile. The type is slice of
    // Ipv4Dhcpd_Nodes_Node_Server_Profiles_Profile.
    Profile []Ipv4Dhcpd_Nodes_Node_Server_Profiles_Profile
}

func (profiles *Ipv4Dhcpd_Nodes_Node_Server_Profiles) GetFilter() yfilter.YFilter { return profiles.YFilter }

func (profiles *Ipv4Dhcpd_Nodes_Node_Server_Profiles) SetFilter(yf yfilter.YFilter) { profiles.YFilter = yf }

func (profiles *Ipv4Dhcpd_Nodes_Node_Server_Profiles) GetGoName(yname string) string {
    if yname == "profile" { return "Profile" }
    return ""
}

func (profiles *Ipv4Dhcpd_Nodes_Node_Server_Profiles) GetSegmentPath() string {
    return "profiles"
}

func (profiles *Ipv4Dhcpd_Nodes_Node_Server_Profiles) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "profile" {
        for _, c := range profiles.Profile {
            if profiles.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4Dhcpd_Nodes_Node_Server_Profiles_Profile{}
        profiles.Profile = append(profiles.Profile, child)
        return &profiles.Profile[len(profiles.Profile)-1]
    }
    return nil
}

func (profiles *Ipv4Dhcpd_Nodes_Node_Server_Profiles) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range profiles.Profile {
        children[profiles.Profile[i].GetSegmentPath()] = &profiles.Profile[i]
    }
    return children
}

func (profiles *Ipv4Dhcpd_Nodes_Node_Server_Profiles) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (profiles *Ipv4Dhcpd_Nodes_Node_Server_Profiles) GetBundleName() string { return "cisco_ios_xr" }

func (profiles *Ipv4Dhcpd_Nodes_Node_Server_Profiles) GetYangName() string { return "profiles" }

func (profiles *Ipv4Dhcpd_Nodes_Node_Server_Profiles) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (profiles *Ipv4Dhcpd_Nodes_Node_Server_Profiles) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (profiles *Ipv4Dhcpd_Nodes_Node_Server_Profiles) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (profiles *Ipv4Dhcpd_Nodes_Node_Server_Profiles) SetParent(parent types.Entity) { profiles.parent = parent }

func (profiles *Ipv4Dhcpd_Nodes_Node_Server_Profiles) GetParent() types.Entity { return profiles.parent }

func (profiles *Ipv4Dhcpd_Nodes_Node_Server_Profiles) GetParentYangName() string { return "server" }

// Ipv4Dhcpd_Nodes_Node_Server_Profiles_Profile
// IPv4 DHCP server profile
type Ipv4Dhcpd_Nodes_Node_Server_Profiles_Profile struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Profile name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    ServerProfileName interface{}

    // Profile Name. The type is string with length: 0..65.
    ServerProfileNameXr interface{}

    // Secure ARP. The type is bool.
    SecureArp interface{}

    // Requested Address Check. The type is bool.
    RequestedAddressCheck interface{}

    // Server ID Check. The type is bool.
    ServerIdCheck interface{}

    // Duplicate MAC Address Check. The type is bool.
    DuplicateMacAddressCheck interface{}

    // Duplicate IP Address Check. The type is bool.
    DuplicateIpAddressCheck interface{}

    // Is true if dhcp subscriber is allowed to move. The type is bool.
    IsMoveAllowed interface{}

    // Bcast Policy. The type is interface{} with range: 0..255.
    BcastPolicy interface{}

    // Giaddr Policy. The type is interface{} with range: 0..255.
    GiaddrPolicy interface{}

    // Subnet Mask. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    SubnetMask interface{}

    // Pool Name. The type is string with length: 0..65.
    ServerPoolName interface{}

    // Lease. The type is interface{} with range: 0..4294967295.
    ServerProfileLease interface{}

    // Server netbios node type. The type is interface{} with range: 0..255.
    ServerProfileNetbiosNodeType interface{}

    // Server Bootfile name. The type is string with length: 0..256.
    ServerBootfileName interface{}

    // Server Domain name. The type is string with length: 0..256.
    ServerDomainName interface{}

    // Server iEdge Check. The type is interface{} with range: 0..255.
    ServerProfileiedgeCheck interface{}

    // Server DNS Count. The type is interface{} with range: 0..255.
    ServerProfileServerDnsCount interface{}

    // Server default count . The type is interface{} with range: 0..255.
    ServerProfiledefaultRouterCount interface{}

    // Server netbios svr count . The type is interface{} with range: 0..255.
    ServerProfileNetbiosNameSvrCount interface{}

    // Server time svr count . The type is interface{} with range: 0..255.
    ServerProfileTimeSvrCount interface{}

    // Lease Limit Type. The type is interface{} with range: 0..255.
    LeaseLimitType interface{}

    // Lease Limit Count. The type is interface{} with range: 0..4294967295.
    LeaseLimitCount interface{}

    // Server DNS addresses. The type is slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    ServerProfileDns []interface{}

    // Server default addresses. The type is slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    ServerProfileDefaultRouter []interface{}

    // Server netbios addresses. The type is slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    ServerProfileNetbiousNameServer []interface{}

    // Server Time addresses. The type is slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    ServerProfileTimeServer []interface{}
}

func (profile *Ipv4Dhcpd_Nodes_Node_Server_Profiles_Profile) GetFilter() yfilter.YFilter { return profile.YFilter }

func (profile *Ipv4Dhcpd_Nodes_Node_Server_Profiles_Profile) SetFilter(yf yfilter.YFilter) { profile.YFilter = yf }

func (profile *Ipv4Dhcpd_Nodes_Node_Server_Profiles_Profile) GetGoName(yname string) string {
    if yname == "server-profile-name" { return "ServerProfileName" }
    if yname == "server-profile-name-xr" { return "ServerProfileNameXr" }
    if yname == "secure-arp" { return "SecureArp" }
    if yname == "requested-address-check" { return "RequestedAddressCheck" }
    if yname == "server-id-check" { return "ServerIdCheck" }
    if yname == "duplicate-mac-address-check" { return "DuplicateMacAddressCheck" }
    if yname == "duplicate-ip-address-check" { return "DuplicateIpAddressCheck" }
    if yname == "is-move-allowed" { return "IsMoveAllowed" }
    if yname == "bcast-policy" { return "BcastPolicy" }
    if yname == "giaddr-policy" { return "GiaddrPolicy" }
    if yname == "subnet-mask" { return "SubnetMask" }
    if yname == "server-pool-name" { return "ServerPoolName" }
    if yname == "server-profile-lease" { return "ServerProfileLease" }
    if yname == "server-profile-netbios-node-type" { return "ServerProfileNetbiosNodeType" }
    if yname == "server-bootfile-name" { return "ServerBootfileName" }
    if yname == "server-domain-name" { return "ServerDomainName" }
    if yname == "server-profileiedge-check" { return "ServerProfileiedgeCheck" }
    if yname == "server-profile-server-dns-count" { return "ServerProfileServerDnsCount" }
    if yname == "server-profiledefault-router-count" { return "ServerProfiledefaultRouterCount" }
    if yname == "server-profile-netbios-name-svr-count" { return "ServerProfileNetbiosNameSvrCount" }
    if yname == "server-profile-time-svr-count" { return "ServerProfileTimeSvrCount" }
    if yname == "lease-limit-type" { return "LeaseLimitType" }
    if yname == "lease-limit-count" { return "LeaseLimitCount" }
    if yname == "server-profile-dns" { return "ServerProfileDns" }
    if yname == "server-profile-default-router" { return "ServerProfileDefaultRouter" }
    if yname == "server-profile-netbious-name-server" { return "ServerProfileNetbiousNameServer" }
    if yname == "server-profile-time-server" { return "ServerProfileTimeServer" }
    return ""
}

func (profile *Ipv4Dhcpd_Nodes_Node_Server_Profiles_Profile) GetSegmentPath() string {
    return "profile" + "[server-profile-name='" + fmt.Sprintf("%v", profile.ServerProfileName) + "']"
}

func (profile *Ipv4Dhcpd_Nodes_Node_Server_Profiles_Profile) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (profile *Ipv4Dhcpd_Nodes_Node_Server_Profiles_Profile) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (profile *Ipv4Dhcpd_Nodes_Node_Server_Profiles_Profile) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["server-profile-name"] = profile.ServerProfileName
    leafs["server-profile-name-xr"] = profile.ServerProfileNameXr
    leafs["secure-arp"] = profile.SecureArp
    leafs["requested-address-check"] = profile.RequestedAddressCheck
    leafs["server-id-check"] = profile.ServerIdCheck
    leafs["duplicate-mac-address-check"] = profile.DuplicateMacAddressCheck
    leafs["duplicate-ip-address-check"] = profile.DuplicateIpAddressCheck
    leafs["is-move-allowed"] = profile.IsMoveAllowed
    leafs["bcast-policy"] = profile.BcastPolicy
    leafs["giaddr-policy"] = profile.GiaddrPolicy
    leafs["subnet-mask"] = profile.SubnetMask
    leafs["server-pool-name"] = profile.ServerPoolName
    leafs["server-profile-lease"] = profile.ServerProfileLease
    leafs["server-profile-netbios-node-type"] = profile.ServerProfileNetbiosNodeType
    leafs["server-bootfile-name"] = profile.ServerBootfileName
    leafs["server-domain-name"] = profile.ServerDomainName
    leafs["server-profileiedge-check"] = profile.ServerProfileiedgeCheck
    leafs["server-profile-server-dns-count"] = profile.ServerProfileServerDnsCount
    leafs["server-profiledefault-router-count"] = profile.ServerProfiledefaultRouterCount
    leafs["server-profile-netbios-name-svr-count"] = profile.ServerProfileNetbiosNameSvrCount
    leafs["server-profile-time-svr-count"] = profile.ServerProfileTimeSvrCount
    leafs["lease-limit-type"] = profile.LeaseLimitType
    leafs["lease-limit-count"] = profile.LeaseLimitCount
    leafs["server-profile-dns"] = profile.ServerProfileDns
    leafs["server-profile-default-router"] = profile.ServerProfileDefaultRouter
    leafs["server-profile-netbious-name-server"] = profile.ServerProfileNetbiousNameServer
    leafs["server-profile-time-server"] = profile.ServerProfileTimeServer
    return leafs
}

func (profile *Ipv4Dhcpd_Nodes_Node_Server_Profiles_Profile) GetBundleName() string { return "cisco_ios_xr" }

func (profile *Ipv4Dhcpd_Nodes_Node_Server_Profiles_Profile) GetYangName() string { return "profile" }

func (profile *Ipv4Dhcpd_Nodes_Node_Server_Profiles_Profile) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (profile *Ipv4Dhcpd_Nodes_Node_Server_Profiles_Profile) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (profile *Ipv4Dhcpd_Nodes_Node_Server_Profiles_Profile) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (profile *Ipv4Dhcpd_Nodes_Node_Server_Profiles_Profile) SetParent(parent types.Entity) { profile.parent = parent }

func (profile *Ipv4Dhcpd_Nodes_Node_Server_Profiles_Profile) GetParent() types.Entity { return profile.parent }

func (profile *Ipv4Dhcpd_Nodes_Node_Server_Profiles_Profile) GetParentYangName() string { return "profiles" }

// Ipv4Dhcpd_Nodes_Node_Server_Statistics
// DHCP Server statistics
type Ipv4Dhcpd_Nodes_Node_Server_Statistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ipv4 dhcpd proxy stat. The type is slice of
    // Ipv4Dhcpd_Nodes_Node_Server_Statistics_Ipv4DhcpdProxyStat.
    Ipv4DhcpdProxyStat []Ipv4Dhcpd_Nodes_Node_Server_Statistics_Ipv4DhcpdProxyStat
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Server_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *Ipv4Dhcpd_Nodes_Node_Server_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *Ipv4Dhcpd_Nodes_Node_Server_Statistics) GetGoName(yname string) string {
    if yname == "ipv4-dhcpd-proxy-stat" { return "Ipv4DhcpdProxyStat" }
    return ""
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Server_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Server_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv4-dhcpd-proxy-stat" {
        for _, c := range statistics.Ipv4DhcpdProxyStat {
            if statistics.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4Dhcpd_Nodes_Node_Server_Statistics_Ipv4DhcpdProxyStat{}
        statistics.Ipv4DhcpdProxyStat = append(statistics.Ipv4DhcpdProxyStat, child)
        return &statistics.Ipv4DhcpdProxyStat[len(statistics.Ipv4DhcpdProxyStat)-1]
    }
    return nil
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Server_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range statistics.Ipv4DhcpdProxyStat {
        children[statistics.Ipv4DhcpdProxyStat[i].GetSegmentPath()] = &statistics.Ipv4DhcpdProxyStat[i]
    }
    return children
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Server_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Server_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *Ipv4Dhcpd_Nodes_Node_Server_Statistics) GetYangName() string { return "statistics" }

func (statistics *Ipv4Dhcpd_Nodes_Node_Server_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *Ipv4Dhcpd_Nodes_Node_Server_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *Ipv4Dhcpd_Nodes_Node_Server_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *Ipv4Dhcpd_Nodes_Node_Server_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *Ipv4Dhcpd_Nodes_Node_Server_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *Ipv4Dhcpd_Nodes_Node_Server_Statistics) GetParentYangName() string { return "server" }

// Ipv4Dhcpd_Nodes_Node_Server_Statistics_Ipv4DhcpdProxyStat
// ipv4 dhcpd proxy stat
type Ipv4Dhcpd_Nodes_Node_Server_Statistics_Ipv4DhcpdProxyStat struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DHCP L3 VRF name. The type is string with length: 0..33.
    VrfName interface{}

    // Proxy statistics.
    Statistics Ipv4Dhcpd_Nodes_Node_Server_Statistics_Ipv4DhcpdProxyStat_Statistics
}

func (ipv4DhcpdProxyStat *Ipv4Dhcpd_Nodes_Node_Server_Statistics_Ipv4DhcpdProxyStat) GetFilter() yfilter.YFilter { return ipv4DhcpdProxyStat.YFilter }

func (ipv4DhcpdProxyStat *Ipv4Dhcpd_Nodes_Node_Server_Statistics_Ipv4DhcpdProxyStat) SetFilter(yf yfilter.YFilter) { ipv4DhcpdProxyStat.YFilter = yf }

func (ipv4DhcpdProxyStat *Ipv4Dhcpd_Nodes_Node_Server_Statistics_Ipv4DhcpdProxyStat) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "statistics" { return "Statistics" }
    return ""
}

func (ipv4DhcpdProxyStat *Ipv4Dhcpd_Nodes_Node_Server_Statistics_Ipv4DhcpdProxyStat) GetSegmentPath() string {
    return "ipv4-dhcpd-proxy-stat"
}

func (ipv4DhcpdProxyStat *Ipv4Dhcpd_Nodes_Node_Server_Statistics_Ipv4DhcpdProxyStat) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "statistics" {
        return &ipv4DhcpdProxyStat.Statistics
    }
    return nil
}

func (ipv4DhcpdProxyStat *Ipv4Dhcpd_Nodes_Node_Server_Statistics_Ipv4DhcpdProxyStat) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["statistics"] = &ipv4DhcpdProxyStat.Statistics
    return children
}

func (ipv4DhcpdProxyStat *Ipv4Dhcpd_Nodes_Node_Server_Statistics_Ipv4DhcpdProxyStat) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = ipv4DhcpdProxyStat.VrfName
    return leafs
}

func (ipv4DhcpdProxyStat *Ipv4Dhcpd_Nodes_Node_Server_Statistics_Ipv4DhcpdProxyStat) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4DhcpdProxyStat *Ipv4Dhcpd_Nodes_Node_Server_Statistics_Ipv4DhcpdProxyStat) GetYangName() string { return "ipv4-dhcpd-proxy-stat" }

func (ipv4DhcpdProxyStat *Ipv4Dhcpd_Nodes_Node_Server_Statistics_Ipv4DhcpdProxyStat) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4DhcpdProxyStat *Ipv4Dhcpd_Nodes_Node_Server_Statistics_Ipv4DhcpdProxyStat) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4DhcpdProxyStat *Ipv4Dhcpd_Nodes_Node_Server_Statistics_Ipv4DhcpdProxyStat) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4DhcpdProxyStat *Ipv4Dhcpd_Nodes_Node_Server_Statistics_Ipv4DhcpdProxyStat) SetParent(parent types.Entity) { ipv4DhcpdProxyStat.parent = parent }

func (ipv4DhcpdProxyStat *Ipv4Dhcpd_Nodes_Node_Server_Statistics_Ipv4DhcpdProxyStat) GetParent() types.Entity { return ipv4DhcpdProxyStat.parent }

func (ipv4DhcpdProxyStat *Ipv4Dhcpd_Nodes_Node_Server_Statistics_Ipv4DhcpdProxyStat) GetParentYangName() string { return "statistics" }

// Ipv4Dhcpd_Nodes_Node_Server_Statistics_Ipv4DhcpdProxyStat_Statistics
// Proxy statistics
type Ipv4Dhcpd_Nodes_Node_Server_Statistics_Ipv4DhcpdProxyStat_Statistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Received packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPackets interface{}

    // Transmitted packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedPackets interface{}

    // Dropped packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Server_Statistics_Ipv4DhcpdProxyStat_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *Ipv4Dhcpd_Nodes_Node_Server_Statistics_Ipv4DhcpdProxyStat_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *Ipv4Dhcpd_Nodes_Node_Server_Statistics_Ipv4DhcpdProxyStat_Statistics) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Server_Statistics_Ipv4DhcpdProxyStat_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Server_Statistics_Ipv4DhcpdProxyStat_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Server_Statistics_Ipv4DhcpdProxyStat_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Server_Statistics_Ipv4DhcpdProxyStat_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = statistics.ReceivedPackets
    leafs["transmitted-packets"] = statistics.TransmittedPackets
    leafs["dropped-packets"] = statistics.DroppedPackets
    return leafs
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Server_Statistics_Ipv4DhcpdProxyStat_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *Ipv4Dhcpd_Nodes_Node_Server_Statistics_Ipv4DhcpdProxyStat_Statistics) GetYangName() string { return "statistics" }

func (statistics *Ipv4Dhcpd_Nodes_Node_Server_Statistics_Ipv4DhcpdProxyStat_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *Ipv4Dhcpd_Nodes_Node_Server_Statistics_Ipv4DhcpdProxyStat_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *Ipv4Dhcpd_Nodes_Node_Server_Statistics_Ipv4DhcpdProxyStat_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *Ipv4Dhcpd_Nodes_Node_Server_Statistics_Ipv4DhcpdProxyStat_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *Ipv4Dhcpd_Nodes_Node_Server_Statistics_Ipv4DhcpdProxyStat_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *Ipv4Dhcpd_Nodes_Node_Server_Statistics_Ipv4DhcpdProxyStat_Statistics) GetParentYangName() string { return "ipv4-dhcpd-proxy-stat" }

// Ipv4Dhcpd_Nodes_Node_Server_Binding
// DHCP server bindings
type Ipv4Dhcpd_Nodes_Node_Server_Binding struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DHCP server binding summary.
    Summary Ipv4Dhcpd_Nodes_Node_Server_Binding_Summary

    // DHCP server client bindings.
    Clients Ipv4Dhcpd_Nodes_Node_Server_Binding_Clients
}

func (binding *Ipv4Dhcpd_Nodes_Node_Server_Binding) GetFilter() yfilter.YFilter { return binding.YFilter }

func (binding *Ipv4Dhcpd_Nodes_Node_Server_Binding) SetFilter(yf yfilter.YFilter) { binding.YFilter = yf }

func (binding *Ipv4Dhcpd_Nodes_Node_Server_Binding) GetGoName(yname string) string {
    if yname == "summary" { return "Summary" }
    if yname == "clients" { return "Clients" }
    return ""
}

func (binding *Ipv4Dhcpd_Nodes_Node_Server_Binding) GetSegmentPath() string {
    return "binding"
}

func (binding *Ipv4Dhcpd_Nodes_Node_Server_Binding) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "summary" {
        return &binding.Summary
    }
    if childYangName == "clients" {
        return &binding.Clients
    }
    return nil
}

func (binding *Ipv4Dhcpd_Nodes_Node_Server_Binding) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["summary"] = &binding.Summary
    children["clients"] = &binding.Clients
    return children
}

func (binding *Ipv4Dhcpd_Nodes_Node_Server_Binding) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (binding *Ipv4Dhcpd_Nodes_Node_Server_Binding) GetBundleName() string { return "cisco_ios_xr" }

func (binding *Ipv4Dhcpd_Nodes_Node_Server_Binding) GetYangName() string { return "binding" }

func (binding *Ipv4Dhcpd_Nodes_Node_Server_Binding) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (binding *Ipv4Dhcpd_Nodes_Node_Server_Binding) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (binding *Ipv4Dhcpd_Nodes_Node_Server_Binding) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (binding *Ipv4Dhcpd_Nodes_Node_Server_Binding) SetParent(parent types.Entity) { binding.parent = parent }

func (binding *Ipv4Dhcpd_Nodes_Node_Server_Binding) GetParent() types.Entity { return binding.parent }

func (binding *Ipv4Dhcpd_Nodes_Node_Server_Binding) GetParentYangName() string { return "server" }

// Ipv4Dhcpd_Nodes_Node_Server_Binding_Summary
// DHCP server binding summary
type Ipv4Dhcpd_Nodes_Node_Server_Binding_Summary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Total number of clients. The type is interface{} with range: 0..4294967295.
    Clients interface{}

    // Number of clients in init state. The type is interface{} with range:
    // 0..4294967295.
    InitializingClients interface{}

    // Number of clients in Init DPM wait state. The type is interface{} with
    // range: 0..4294967295.
    WaitingForDpmInit interface{}

    // Number of clients in Request DPM wait state. The type is interface{} with
    // range: 0..4294967295.
    WaitingForDpmRequest interface{}

    // Number of clients in Init DAPS wait state. The type is interface{} with
    // range: 0..4294967295.
    WaitingForDapsInit interface{}

    // Number of clients in selecting state. The type is interface{} with range:
    // 0..4294967295.
    SelectingClients interface{}

    // Number of clients in Offer sent state. The type is interface{} with range:
    // 0..4294967295.
    OfferSentForClient interface{}

    // Number of clients in requesting state. The type is interface{} with range:
    // 0..4294967295.
    RequestingClients interface{}

    // Number of clients in Waiting for DPM with Request. The type is interface{}
    // with range: 0..4294967295.
    RequestWaitingForDpm interface{}

    // Number of clients in Waiting for DPM with ACK. The type is interface{} with
    // range: 0..4294967295.
    AckWaitingForDpm interface{}

    // Number of clients in bound state. The type is interface{} with range:
    // 0..4294967295.
    BoundClients interface{}

    // Number of clients in renewing state. The type is interface{} with range:
    // 0..4294967295.
    RenewingClients interface{}

    // Number of clients in informing state. The type is interface{} with range:
    // 0..4294967295.
    InformingClients interface{}

    // Number of clients in reauth state. The type is interface{} with range:
    // 0..4294967295.
    ReauthorizingClients interface{}

    // Number of clients in waiting for DPM disconnect state. The type is
    // interface{} with range: 0..4294967295.
    WaitingForDpmDisconnect interface{}

    // Number of clients in Waiting for DPM after addr change. The type is
    // interface{} with range: 0..4294967295.
    WaitingForDpmAddrChange interface{}

    // Number of clients in deleting state. The type is interface{} with range:
    // 0..4294967295.
    DeletingClientsD interface{}

    // Number of clients in disconnected state. The type is interface{} with
    // range: 0..4294967295.
    DisconnectedClients interface{}

    // Number of clients in restarting state. The type is interface{} with range:
    // 0..4294967295.
    RestartingClients interface{}
}

func (summary *Ipv4Dhcpd_Nodes_Node_Server_Binding_Summary) GetFilter() yfilter.YFilter { return summary.YFilter }

func (summary *Ipv4Dhcpd_Nodes_Node_Server_Binding_Summary) SetFilter(yf yfilter.YFilter) { summary.YFilter = yf }

func (summary *Ipv4Dhcpd_Nodes_Node_Server_Binding_Summary) GetGoName(yname string) string {
    if yname == "clients" { return "Clients" }
    if yname == "initializing-clients" { return "InitializingClients" }
    if yname == "waiting-for-dpm-init" { return "WaitingForDpmInit" }
    if yname == "waiting-for-dpm-request" { return "WaitingForDpmRequest" }
    if yname == "waiting-for-daps-init" { return "WaitingForDapsInit" }
    if yname == "selecting-clients" { return "SelectingClients" }
    if yname == "offer-sent-for-client" { return "OfferSentForClient" }
    if yname == "requesting-clients" { return "RequestingClients" }
    if yname == "request-waiting-for-dpm" { return "RequestWaitingForDpm" }
    if yname == "ack-waiting-for-dpm" { return "AckWaitingForDpm" }
    if yname == "bound-clients" { return "BoundClients" }
    if yname == "renewing-clients" { return "RenewingClients" }
    if yname == "informing-clients" { return "InformingClients" }
    if yname == "reauthorizing-clients" { return "ReauthorizingClients" }
    if yname == "waiting-for-dpm-disconnect" { return "WaitingForDpmDisconnect" }
    if yname == "waiting-for-dpm-addr-change" { return "WaitingForDpmAddrChange" }
    if yname == "deleting-clients-d" { return "DeletingClientsD" }
    if yname == "disconnected-clients" { return "DisconnectedClients" }
    if yname == "restarting-clients" { return "RestartingClients" }
    return ""
}

func (summary *Ipv4Dhcpd_Nodes_Node_Server_Binding_Summary) GetSegmentPath() string {
    return "summary"
}

func (summary *Ipv4Dhcpd_Nodes_Node_Server_Binding_Summary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (summary *Ipv4Dhcpd_Nodes_Node_Server_Binding_Summary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (summary *Ipv4Dhcpd_Nodes_Node_Server_Binding_Summary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["clients"] = summary.Clients
    leafs["initializing-clients"] = summary.InitializingClients
    leafs["waiting-for-dpm-init"] = summary.WaitingForDpmInit
    leafs["waiting-for-dpm-request"] = summary.WaitingForDpmRequest
    leafs["waiting-for-daps-init"] = summary.WaitingForDapsInit
    leafs["selecting-clients"] = summary.SelectingClients
    leafs["offer-sent-for-client"] = summary.OfferSentForClient
    leafs["requesting-clients"] = summary.RequestingClients
    leafs["request-waiting-for-dpm"] = summary.RequestWaitingForDpm
    leafs["ack-waiting-for-dpm"] = summary.AckWaitingForDpm
    leafs["bound-clients"] = summary.BoundClients
    leafs["renewing-clients"] = summary.RenewingClients
    leafs["informing-clients"] = summary.InformingClients
    leafs["reauthorizing-clients"] = summary.ReauthorizingClients
    leafs["waiting-for-dpm-disconnect"] = summary.WaitingForDpmDisconnect
    leafs["waiting-for-dpm-addr-change"] = summary.WaitingForDpmAddrChange
    leafs["deleting-clients-d"] = summary.DeletingClientsD
    leafs["disconnected-clients"] = summary.DisconnectedClients
    leafs["restarting-clients"] = summary.RestartingClients
    return leafs
}

func (summary *Ipv4Dhcpd_Nodes_Node_Server_Binding_Summary) GetBundleName() string { return "cisco_ios_xr" }

func (summary *Ipv4Dhcpd_Nodes_Node_Server_Binding_Summary) GetYangName() string { return "summary" }

func (summary *Ipv4Dhcpd_Nodes_Node_Server_Binding_Summary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (summary *Ipv4Dhcpd_Nodes_Node_Server_Binding_Summary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (summary *Ipv4Dhcpd_Nodes_Node_Server_Binding_Summary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (summary *Ipv4Dhcpd_Nodes_Node_Server_Binding_Summary) SetParent(parent types.Entity) { summary.parent = parent }

func (summary *Ipv4Dhcpd_Nodes_Node_Server_Binding_Summary) GetParent() types.Entity { return summary.parent }

func (summary *Ipv4Dhcpd_Nodes_Node_Server_Binding_Summary) GetParentYangName() string { return "binding" }

// Ipv4Dhcpd_Nodes_Node_Server_Binding_Clients
// DHCP server client bindings
type Ipv4Dhcpd_Nodes_Node_Server_Binding_Clients struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Single DHCP Server binding. The type is slice of
    // Ipv4Dhcpd_Nodes_Node_Server_Binding_Clients_Client.
    Client []Ipv4Dhcpd_Nodes_Node_Server_Binding_Clients_Client
}

func (clients *Ipv4Dhcpd_Nodes_Node_Server_Binding_Clients) GetFilter() yfilter.YFilter { return clients.YFilter }

func (clients *Ipv4Dhcpd_Nodes_Node_Server_Binding_Clients) SetFilter(yf yfilter.YFilter) { clients.YFilter = yf }

func (clients *Ipv4Dhcpd_Nodes_Node_Server_Binding_Clients) GetGoName(yname string) string {
    if yname == "client" { return "Client" }
    return ""
}

func (clients *Ipv4Dhcpd_Nodes_Node_Server_Binding_Clients) GetSegmentPath() string {
    return "clients"
}

func (clients *Ipv4Dhcpd_Nodes_Node_Server_Binding_Clients) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "client" {
        for _, c := range clients.Client {
            if clients.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4Dhcpd_Nodes_Node_Server_Binding_Clients_Client{}
        clients.Client = append(clients.Client, child)
        return &clients.Client[len(clients.Client)-1]
    }
    return nil
}

func (clients *Ipv4Dhcpd_Nodes_Node_Server_Binding_Clients) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range clients.Client {
        children[clients.Client[i].GetSegmentPath()] = &clients.Client[i]
    }
    return children
}

func (clients *Ipv4Dhcpd_Nodes_Node_Server_Binding_Clients) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (clients *Ipv4Dhcpd_Nodes_Node_Server_Binding_Clients) GetBundleName() string { return "cisco_ios_xr" }

func (clients *Ipv4Dhcpd_Nodes_Node_Server_Binding_Clients) GetYangName() string { return "clients" }

func (clients *Ipv4Dhcpd_Nodes_Node_Server_Binding_Clients) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clients *Ipv4Dhcpd_Nodes_Node_Server_Binding_Clients) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clients *Ipv4Dhcpd_Nodes_Node_Server_Binding_Clients) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clients *Ipv4Dhcpd_Nodes_Node_Server_Binding_Clients) SetParent(parent types.Entity) { clients.parent = parent }

func (clients *Ipv4Dhcpd_Nodes_Node_Server_Binding_Clients) GetParent() types.Entity { return clients.parent }

func (clients *Ipv4Dhcpd_Nodes_Node_Server_Binding_Clients) GetParentYangName() string { return "binding" }

// Ipv4Dhcpd_Nodes_Node_Server_Binding_Clients_Client
// Single DHCP Server binding
type Ipv4Dhcpd_Nodes_Node_Server_Binding_Clients_Client struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Client ID. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    ClientId interface{}

    // DHCP client identifier. The type is string with length: 0..1275.
    ClientIdXr interface{}

    // DHCP client MAC address. The type is string.
    MacAddress interface{}

    // DHCP client/subscriber VRF name. The type is string with length: 0..33.
    VrfName interface{}

    // DHCP server VRF name. The type is string with length: 0..33.
    ServerVrfName interface{}

    // DHCP IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    IpAddress interface{}

    // DHCP client GIADDR. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    ClientGiAddr interface{}

    // DHCP to server GIADDR. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    ToServerGiAddr interface{}

    // DHCP server IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    ServerIpAddress interface{}

    // DHCP reply server IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    ReplyServerIpAddress interface{}

    // Lease time in seconds. The type is interface{} with range: 0..4294967295.
    // Units are second.
    LeaseTime interface{}

    // Remaining lease time in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    RemainingLeaseTime interface{}

    // DHCP client state. The type is BagDhcpdProxyState.
    State interface{}

    // DHCP access interface to client. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // DHCP access interface VRF name. The type is string with length: 0..33.
    AccessVrfName interface{}

    // DHCP VLAN outer VLAN. The type is interface{} with range: 0..4294967295.
    ProxyBindingOuterTag interface{}

    // DHCP VLAN inner VLAN. The type is interface{} with range: 0..4294967295.
    ProxyBindingInnerTag interface{}

    // DHCP profile name. The type is string with length: 0..65.
    ProfileName interface{}

    // Is true if DHCP next renew from client will be NAK'd. The type is bool.
    IsNakNextRenew interface{}

    // DHCP subscriber label. The type is interface{} with range: 0..4294967295.
    SubscriberLabel interface{}

    // DHCP old subscriber label. The type is interface{} with range:
    // 0..4294967295.
    OldSubscriberLabel interface{}

    // DHCP subscriber interface. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    SubscriberInterfaceName interface{}

    // DHCP received circuit ID. The type is string with length: 0..768.
    RxCircuitId interface{}

    // DHCP transmitted circuit ID. The type is string with length: 0..768.
    TxCircuitId interface{}

    // DHCP received Remote ID. The type is string with length: 0..768.
    RxRemoteId interface{}

    // DHCP transmitted Remote ID. The type is string with length: 0..768.
    TxRemoteId interface{}

    // DHCP received VSISO. The type is string with length: 0..768.
    RxVsiso interface{}

    // DHCP transmitted VSISO. The type is string with length: 0..768.
    TxVsiso interface{}

    // Is true if authentication is on received option82. The type is bool.
    IsAuthReceived interface{}

    // Is true if DHCP subscriber is Mobile. The type is bool.
    IsMblSubscriber interface{}

    // DHCP parameter request option. The type is string with length: 0..513.
    ParamRequest interface{}

    // DHCP saved options. The type is string with length: 0..2051.
    ParamResponse interface{}

    // session start time. The type is interface{} with range:
    // 0..18446744073709551615.
    SessionStartTime interface{}

    // DHCPV4 SRG state. The type is interface{} with range: 0..4294967295.
    SrgState interface{}

    // event history. The type is slice of interface{} with range: 0..4294967295.
    EventHistory []interface{}
}

func (client *Ipv4Dhcpd_Nodes_Node_Server_Binding_Clients_Client) GetFilter() yfilter.YFilter { return client.YFilter }

func (client *Ipv4Dhcpd_Nodes_Node_Server_Binding_Clients_Client) SetFilter(yf yfilter.YFilter) { client.YFilter = yf }

func (client *Ipv4Dhcpd_Nodes_Node_Server_Binding_Clients_Client) GetGoName(yname string) string {
    if yname == "client-id" { return "ClientId" }
    if yname == "client-id-xr" { return "ClientIdXr" }
    if yname == "mac-address" { return "MacAddress" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "server-vrf-name" { return "ServerVrfName" }
    if yname == "ip-address" { return "IpAddress" }
    if yname == "client-gi-addr" { return "ClientGiAddr" }
    if yname == "to-server-gi-addr" { return "ToServerGiAddr" }
    if yname == "server-ip-address" { return "ServerIpAddress" }
    if yname == "reply-server-ip-address" { return "ReplyServerIpAddress" }
    if yname == "lease-time" { return "LeaseTime" }
    if yname == "remaining-lease-time" { return "RemainingLeaseTime" }
    if yname == "state" { return "State" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "access-vrf-name" { return "AccessVrfName" }
    if yname == "proxy-binding-outer-tag" { return "ProxyBindingOuterTag" }
    if yname == "proxy-binding-inner-tag" { return "ProxyBindingInnerTag" }
    if yname == "profile-name" { return "ProfileName" }
    if yname == "is-nak-next-renew" { return "IsNakNextRenew" }
    if yname == "subscriber-label" { return "SubscriberLabel" }
    if yname == "old-subscriber-label" { return "OldSubscriberLabel" }
    if yname == "subscriber-interface-name" { return "SubscriberInterfaceName" }
    if yname == "rx-circuit-id" { return "RxCircuitId" }
    if yname == "tx-circuit-id" { return "TxCircuitId" }
    if yname == "rx-remote-id" { return "RxRemoteId" }
    if yname == "tx-remote-id" { return "TxRemoteId" }
    if yname == "rx-vsiso" { return "RxVsiso" }
    if yname == "tx-vsiso" { return "TxVsiso" }
    if yname == "is-auth-received" { return "IsAuthReceived" }
    if yname == "is-mbl-subscriber" { return "IsMblSubscriber" }
    if yname == "param-request" { return "ParamRequest" }
    if yname == "param-response" { return "ParamResponse" }
    if yname == "session-start-time" { return "SessionStartTime" }
    if yname == "srg-state" { return "SrgState" }
    if yname == "event-history" { return "EventHistory" }
    return ""
}

func (client *Ipv4Dhcpd_Nodes_Node_Server_Binding_Clients_Client) GetSegmentPath() string {
    return "client" + "[client-id='" + fmt.Sprintf("%v", client.ClientId) + "']"
}

func (client *Ipv4Dhcpd_Nodes_Node_Server_Binding_Clients_Client) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (client *Ipv4Dhcpd_Nodes_Node_Server_Binding_Clients_Client) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (client *Ipv4Dhcpd_Nodes_Node_Server_Binding_Clients_Client) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["client-id"] = client.ClientId
    leafs["client-id-xr"] = client.ClientIdXr
    leafs["mac-address"] = client.MacAddress
    leafs["vrf-name"] = client.VrfName
    leafs["server-vrf-name"] = client.ServerVrfName
    leafs["ip-address"] = client.IpAddress
    leafs["client-gi-addr"] = client.ClientGiAddr
    leafs["to-server-gi-addr"] = client.ToServerGiAddr
    leafs["server-ip-address"] = client.ServerIpAddress
    leafs["reply-server-ip-address"] = client.ReplyServerIpAddress
    leafs["lease-time"] = client.LeaseTime
    leafs["remaining-lease-time"] = client.RemainingLeaseTime
    leafs["state"] = client.State
    leafs["interface-name"] = client.InterfaceName
    leafs["access-vrf-name"] = client.AccessVrfName
    leafs["proxy-binding-outer-tag"] = client.ProxyBindingOuterTag
    leafs["proxy-binding-inner-tag"] = client.ProxyBindingInnerTag
    leafs["profile-name"] = client.ProfileName
    leafs["is-nak-next-renew"] = client.IsNakNextRenew
    leafs["subscriber-label"] = client.SubscriberLabel
    leafs["old-subscriber-label"] = client.OldSubscriberLabel
    leafs["subscriber-interface-name"] = client.SubscriberInterfaceName
    leafs["rx-circuit-id"] = client.RxCircuitId
    leafs["tx-circuit-id"] = client.TxCircuitId
    leafs["rx-remote-id"] = client.RxRemoteId
    leafs["tx-remote-id"] = client.TxRemoteId
    leafs["rx-vsiso"] = client.RxVsiso
    leafs["tx-vsiso"] = client.TxVsiso
    leafs["is-auth-received"] = client.IsAuthReceived
    leafs["is-mbl-subscriber"] = client.IsMblSubscriber
    leafs["param-request"] = client.ParamRequest
    leafs["param-response"] = client.ParamResponse
    leafs["session-start-time"] = client.SessionStartTime
    leafs["srg-state"] = client.SrgState
    leafs["event-history"] = client.EventHistory
    return leafs
}

func (client *Ipv4Dhcpd_Nodes_Node_Server_Binding_Clients_Client) GetBundleName() string { return "cisco_ios_xr" }

func (client *Ipv4Dhcpd_Nodes_Node_Server_Binding_Clients_Client) GetYangName() string { return "client" }

func (client *Ipv4Dhcpd_Nodes_Node_Server_Binding_Clients_Client) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (client *Ipv4Dhcpd_Nodes_Node_Server_Binding_Clients_Client) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (client *Ipv4Dhcpd_Nodes_Node_Server_Binding_Clients_Client) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (client *Ipv4Dhcpd_Nodes_Node_Server_Binding_Clients_Client) SetParent(parent types.Entity) { client.parent = parent }

func (client *Ipv4Dhcpd_Nodes_Node_Server_Binding_Clients_Client) GetParent() types.Entity { return client.parent }

func (client *Ipv4Dhcpd_Nodes_Node_Server_Binding_Clients_Client) GetParentYangName() string { return "clients" }

// Ipv4Dhcpd_Nodes_Node_Server_StatisticsInfo
// DHCP proxy stats info
type Ipv4Dhcpd_Nodes_Node_Server_StatisticsInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Proxy Stats timestamp. The type is interface{} with range: 0..4294967295.
    ProxyStatsTimestamp interface{}
}

func (statisticsInfo *Ipv4Dhcpd_Nodes_Node_Server_StatisticsInfo) GetFilter() yfilter.YFilter { return statisticsInfo.YFilter }

func (statisticsInfo *Ipv4Dhcpd_Nodes_Node_Server_StatisticsInfo) SetFilter(yf yfilter.YFilter) { statisticsInfo.YFilter = yf }

func (statisticsInfo *Ipv4Dhcpd_Nodes_Node_Server_StatisticsInfo) GetGoName(yname string) string {
    if yname == "proxy-stats-timestamp" { return "ProxyStatsTimestamp" }
    return ""
}

func (statisticsInfo *Ipv4Dhcpd_Nodes_Node_Server_StatisticsInfo) GetSegmentPath() string {
    return "statistics-info"
}

func (statisticsInfo *Ipv4Dhcpd_Nodes_Node_Server_StatisticsInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (statisticsInfo *Ipv4Dhcpd_Nodes_Node_Server_StatisticsInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (statisticsInfo *Ipv4Dhcpd_Nodes_Node_Server_StatisticsInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["proxy-stats-timestamp"] = statisticsInfo.ProxyStatsTimestamp
    return leafs
}

func (statisticsInfo *Ipv4Dhcpd_Nodes_Node_Server_StatisticsInfo) GetBundleName() string { return "cisco_ios_xr" }

func (statisticsInfo *Ipv4Dhcpd_Nodes_Node_Server_StatisticsInfo) GetYangName() string { return "statistics-info" }

func (statisticsInfo *Ipv4Dhcpd_Nodes_Node_Server_StatisticsInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statisticsInfo *Ipv4Dhcpd_Nodes_Node_Server_StatisticsInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statisticsInfo *Ipv4Dhcpd_Nodes_Node_Server_StatisticsInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statisticsInfo *Ipv4Dhcpd_Nodes_Node_Server_StatisticsInfo) SetParent(parent types.Entity) { statisticsInfo.parent = parent }

func (statisticsInfo *Ipv4Dhcpd_Nodes_Node_Server_StatisticsInfo) GetParent() types.Entity { return statisticsInfo.parent }

func (statisticsInfo *Ipv4Dhcpd_Nodes_Node_Server_StatisticsInfo) GetParentYangName() string { return "server" }

// Ipv4Dhcpd_Nodes_Node_Server_Vrfs
// DHCP Server list of VRF names
type Ipv4Dhcpd_Nodes_Node_Server_Vrfs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 DHCP server VRF name. The type is slice of
    // Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf.
    Vrf []Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf
}

func (vrfs *Ipv4Dhcpd_Nodes_Node_Server_Vrfs) GetFilter() yfilter.YFilter { return vrfs.YFilter }

func (vrfs *Ipv4Dhcpd_Nodes_Node_Server_Vrfs) SetFilter(yf yfilter.YFilter) { vrfs.YFilter = yf }

func (vrfs *Ipv4Dhcpd_Nodes_Node_Server_Vrfs) GetGoName(yname string) string {
    if yname == "vrf" { return "Vrf" }
    return ""
}

func (vrfs *Ipv4Dhcpd_Nodes_Node_Server_Vrfs) GetSegmentPath() string {
    return "vrfs"
}

func (vrfs *Ipv4Dhcpd_Nodes_Node_Server_Vrfs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf" {
        for _, c := range vrfs.Vrf {
            if vrfs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf{}
        vrfs.Vrf = append(vrfs.Vrf, child)
        return &vrfs.Vrf[len(vrfs.Vrf)-1]
    }
    return nil
}

func (vrfs *Ipv4Dhcpd_Nodes_Node_Server_Vrfs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vrfs.Vrf {
        children[vrfs.Vrf[i].GetSegmentPath()] = &vrfs.Vrf[i]
    }
    return children
}

func (vrfs *Ipv4Dhcpd_Nodes_Node_Server_Vrfs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vrfs *Ipv4Dhcpd_Nodes_Node_Server_Vrfs) GetBundleName() string { return "cisco_ios_xr" }

func (vrfs *Ipv4Dhcpd_Nodes_Node_Server_Vrfs) GetYangName() string { return "vrfs" }

func (vrfs *Ipv4Dhcpd_Nodes_Node_Server_Vrfs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfs *Ipv4Dhcpd_Nodes_Node_Server_Vrfs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfs *Ipv4Dhcpd_Nodes_Node_Server_Vrfs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfs *Ipv4Dhcpd_Nodes_Node_Server_Vrfs) SetParent(parent types.Entity) { vrfs.parent = parent }

func (vrfs *Ipv4Dhcpd_Nodes_Node_Server_Vrfs) GetParent() types.Entity { return vrfs.parent }

func (vrfs *Ipv4Dhcpd_Nodes_Node_Server_Vrfs) GetParentYangName() string { return "server" }

// Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf
// IPv4 DHCP server VRF name
type Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // IPv4 DHCP server statistics.
    Statistics Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics
}

func (vrf *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf) GetFilter() yfilter.YFilter { return vrf.YFilter }

func (vrf *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf) SetFilter(yf yfilter.YFilter) { vrf.YFilter = yf }

func (vrf *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "statistics" { return "Statistics" }
    return ""
}

func (vrf *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf) GetSegmentPath() string {
    return "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
}

func (vrf *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "statistics" {
        return &vrf.Statistics
    }
    return nil
}

func (vrf *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["statistics"] = &vrf.Statistics
    return children
}

func (vrf *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = vrf.VrfName
    return leafs
}

func (vrf *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf) GetBundleName() string { return "cisco_ios_xr" }

func (vrf *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf) GetYangName() string { return "vrf" }

func (vrf *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrf *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrf *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrf *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf) SetParent(parent types.Entity) { vrf.parent = parent }

func (vrf *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf) GetParent() types.Entity { return vrf.parent }

func (vrf *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf) GetParentYangName() string { return "vrfs" }

// Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics
// IPv4 DHCP server statistics
type Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DHCP discover packets.
    Discover Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Discover

    // DHCP offer packets.
    Offer Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Offer

    // DHCP request packets.
    Request Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Request

    // DHCP decline packets.
    Decline Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Decline

    // DHCP ack packets.
    Ack Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Ack

    // DHCP nak packets.
    Nak Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Nak

    // DHCP release packets.
    Release Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Release

    // DHCP inform packets.
    Inform Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Inform

    // DHCP lease query packets.
    LeaseQuery Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQuery

    // DHCP lease not assigned.
    LeaseNotAssigned Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseNotAssigned

    // DHCP lease unknown.
    LeaseUnknown Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseUnknown

    // DHCP lease active.
    LeaseActive Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseActive

    // DHCP BOOTP request.
    BootpRequest Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_BootpRequest

    // DHCP BOOTP reply.
    BootpReply Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_BootpReply
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics) GetGoName(yname string) string {
    if yname == "discover" { return "Discover" }
    if yname == "offer" { return "Offer" }
    if yname == "request" { return "Request" }
    if yname == "decline" { return "Decline" }
    if yname == "ack" { return "Ack" }
    if yname == "nak" { return "Nak" }
    if yname == "release" { return "Release" }
    if yname == "inform" { return "Inform" }
    if yname == "lease-query" { return "LeaseQuery" }
    if yname == "lease-not-assigned" { return "LeaseNotAssigned" }
    if yname == "lease-unknown" { return "LeaseUnknown" }
    if yname == "lease-active" { return "LeaseActive" }
    if yname == "bootp-request" { return "BootpRequest" }
    if yname == "bootp-reply" { return "BootpReply" }
    return ""
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "discover" {
        return &statistics.Discover
    }
    if childYangName == "offer" {
        return &statistics.Offer
    }
    if childYangName == "request" {
        return &statistics.Request
    }
    if childYangName == "decline" {
        return &statistics.Decline
    }
    if childYangName == "ack" {
        return &statistics.Ack
    }
    if childYangName == "nak" {
        return &statistics.Nak
    }
    if childYangName == "release" {
        return &statistics.Release
    }
    if childYangName == "inform" {
        return &statistics.Inform
    }
    if childYangName == "lease-query" {
        return &statistics.LeaseQuery
    }
    if childYangName == "lease-not-assigned" {
        return &statistics.LeaseNotAssigned
    }
    if childYangName == "lease-unknown" {
        return &statistics.LeaseUnknown
    }
    if childYangName == "lease-active" {
        return &statistics.LeaseActive
    }
    if childYangName == "bootp-request" {
        return &statistics.BootpRequest
    }
    if childYangName == "bootp-reply" {
        return &statistics.BootpReply
    }
    return nil
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["discover"] = &statistics.Discover
    children["offer"] = &statistics.Offer
    children["request"] = &statistics.Request
    children["decline"] = &statistics.Decline
    children["ack"] = &statistics.Ack
    children["nak"] = &statistics.Nak
    children["release"] = &statistics.Release
    children["inform"] = &statistics.Inform
    children["lease-query"] = &statistics.LeaseQuery
    children["lease-not-assigned"] = &statistics.LeaseNotAssigned
    children["lease-unknown"] = &statistics.LeaseUnknown
    children["lease-active"] = &statistics.LeaseActive
    children["bootp-request"] = &statistics.BootpRequest
    children["bootp-reply"] = &statistics.BootpReply
    return children
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics) GetYangName() string { return "statistics" }

func (statistics *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics) GetParentYangName() string { return "vrf" }

// Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Discover
// DHCP discover packets
type Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Discover struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Received packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPackets interface{}

    // Transmitted packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedPackets interface{}

    // Dropped packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}
}

func (discover *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Discover) GetFilter() yfilter.YFilter { return discover.YFilter }

func (discover *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Discover) SetFilter(yf yfilter.YFilter) { discover.YFilter = yf }

func (discover *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Discover) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (discover *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Discover) GetSegmentPath() string {
    return "discover"
}

func (discover *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Discover) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (discover *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Discover) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (discover *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Discover) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = discover.ReceivedPackets
    leafs["transmitted-packets"] = discover.TransmittedPackets
    leafs["dropped-packets"] = discover.DroppedPackets
    return leafs
}

func (discover *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Discover) GetBundleName() string { return "cisco_ios_xr" }

func (discover *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Discover) GetYangName() string { return "discover" }

func (discover *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Discover) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (discover *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Discover) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (discover *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Discover) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (discover *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Discover) SetParent(parent types.Entity) { discover.parent = parent }

func (discover *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Discover) GetParent() types.Entity { return discover.parent }

func (discover *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Discover) GetParentYangName() string { return "statistics" }

// Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Offer
// DHCP offer packets
type Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Offer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Received packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPackets interface{}

    // Transmitted packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedPackets interface{}

    // Dropped packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}
}

func (offer *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Offer) GetFilter() yfilter.YFilter { return offer.YFilter }

func (offer *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Offer) SetFilter(yf yfilter.YFilter) { offer.YFilter = yf }

func (offer *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Offer) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (offer *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Offer) GetSegmentPath() string {
    return "offer"
}

func (offer *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Offer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (offer *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Offer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (offer *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Offer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = offer.ReceivedPackets
    leafs["transmitted-packets"] = offer.TransmittedPackets
    leafs["dropped-packets"] = offer.DroppedPackets
    return leafs
}

func (offer *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Offer) GetBundleName() string { return "cisco_ios_xr" }

func (offer *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Offer) GetYangName() string { return "offer" }

func (offer *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Offer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (offer *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Offer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (offer *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Offer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (offer *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Offer) SetParent(parent types.Entity) { offer.parent = parent }

func (offer *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Offer) GetParent() types.Entity { return offer.parent }

func (offer *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Offer) GetParentYangName() string { return "statistics" }

// Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Request
// DHCP request packets
type Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Request struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Received packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPackets interface{}

    // Transmitted packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedPackets interface{}

    // Dropped packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}
}

func (request *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Request) GetFilter() yfilter.YFilter { return request.YFilter }

func (request *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Request) SetFilter(yf yfilter.YFilter) { request.YFilter = yf }

func (request *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Request) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (request *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Request) GetSegmentPath() string {
    return "request"
}

func (request *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Request) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (request *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Request) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (request *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Request) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = request.ReceivedPackets
    leafs["transmitted-packets"] = request.TransmittedPackets
    leafs["dropped-packets"] = request.DroppedPackets
    return leafs
}

func (request *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Request) GetBundleName() string { return "cisco_ios_xr" }

func (request *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Request) GetYangName() string { return "request" }

func (request *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Request) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (request *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Request) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (request *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Request) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (request *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Request) SetParent(parent types.Entity) { request.parent = parent }

func (request *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Request) GetParent() types.Entity { return request.parent }

func (request *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Request) GetParentYangName() string { return "statistics" }

// Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Decline
// DHCP decline packets
type Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Decline struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Received packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPackets interface{}

    // Transmitted packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedPackets interface{}

    // Dropped packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}
}

func (decline *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Decline) GetFilter() yfilter.YFilter { return decline.YFilter }

func (decline *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Decline) SetFilter(yf yfilter.YFilter) { decline.YFilter = yf }

func (decline *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Decline) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (decline *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Decline) GetSegmentPath() string {
    return "decline"
}

func (decline *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Decline) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (decline *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Decline) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (decline *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Decline) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = decline.ReceivedPackets
    leafs["transmitted-packets"] = decline.TransmittedPackets
    leafs["dropped-packets"] = decline.DroppedPackets
    return leafs
}

func (decline *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Decline) GetBundleName() string { return "cisco_ios_xr" }

func (decline *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Decline) GetYangName() string { return "decline" }

func (decline *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Decline) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (decline *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Decline) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (decline *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Decline) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (decline *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Decline) SetParent(parent types.Entity) { decline.parent = parent }

func (decline *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Decline) GetParent() types.Entity { return decline.parent }

func (decline *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Decline) GetParentYangName() string { return "statistics" }

// Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Ack
// DHCP ack packets
type Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Ack struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Received packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPackets interface{}

    // Transmitted packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedPackets interface{}

    // Dropped packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}
}

func (ack *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Ack) GetFilter() yfilter.YFilter { return ack.YFilter }

func (ack *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Ack) SetFilter(yf yfilter.YFilter) { ack.YFilter = yf }

func (ack *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Ack) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (ack *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Ack) GetSegmentPath() string {
    return "ack"
}

func (ack *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Ack) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ack *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Ack) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ack *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Ack) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = ack.ReceivedPackets
    leafs["transmitted-packets"] = ack.TransmittedPackets
    leafs["dropped-packets"] = ack.DroppedPackets
    return leafs
}

func (ack *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Ack) GetBundleName() string { return "cisco_ios_xr" }

func (ack *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Ack) GetYangName() string { return "ack" }

func (ack *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Ack) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ack *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Ack) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ack *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Ack) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ack *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Ack) SetParent(parent types.Entity) { ack.parent = parent }

func (ack *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Ack) GetParent() types.Entity { return ack.parent }

func (ack *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Ack) GetParentYangName() string { return "statistics" }

// Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Nak
// DHCP nak packets
type Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Nak struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Received packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPackets interface{}

    // Transmitted packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedPackets interface{}

    // Dropped packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}
}

func (nak *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Nak) GetFilter() yfilter.YFilter { return nak.YFilter }

func (nak *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Nak) SetFilter(yf yfilter.YFilter) { nak.YFilter = yf }

func (nak *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Nak) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (nak *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Nak) GetSegmentPath() string {
    return "nak"
}

func (nak *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Nak) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nak *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Nak) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nak *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Nak) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = nak.ReceivedPackets
    leafs["transmitted-packets"] = nak.TransmittedPackets
    leafs["dropped-packets"] = nak.DroppedPackets
    return leafs
}

func (nak *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Nak) GetBundleName() string { return "cisco_ios_xr" }

func (nak *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Nak) GetYangName() string { return "nak" }

func (nak *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Nak) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nak *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Nak) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nak *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Nak) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nak *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Nak) SetParent(parent types.Entity) { nak.parent = parent }

func (nak *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Nak) GetParent() types.Entity { return nak.parent }

func (nak *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Nak) GetParentYangName() string { return "statistics" }

// Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Release
// DHCP release packets
type Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Release struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Received packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPackets interface{}

    // Transmitted packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedPackets interface{}

    // Dropped packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}
}

func (release *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Release) GetFilter() yfilter.YFilter { return release.YFilter }

func (release *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Release) SetFilter(yf yfilter.YFilter) { release.YFilter = yf }

func (release *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Release) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (release *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Release) GetSegmentPath() string {
    return "release"
}

func (release *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Release) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (release *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Release) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (release *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Release) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = release.ReceivedPackets
    leafs["transmitted-packets"] = release.TransmittedPackets
    leafs["dropped-packets"] = release.DroppedPackets
    return leafs
}

func (release *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Release) GetBundleName() string { return "cisco_ios_xr" }

func (release *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Release) GetYangName() string { return "release" }

func (release *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Release) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (release *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Release) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (release *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Release) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (release *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Release) SetParent(parent types.Entity) { release.parent = parent }

func (release *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Release) GetParent() types.Entity { return release.parent }

func (release *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Release) GetParentYangName() string { return "statistics" }

// Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Inform
// DHCP inform packets
type Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Inform struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Received packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPackets interface{}

    // Transmitted packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedPackets interface{}

    // Dropped packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}
}

func (inform *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Inform) GetFilter() yfilter.YFilter { return inform.YFilter }

func (inform *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Inform) SetFilter(yf yfilter.YFilter) { inform.YFilter = yf }

func (inform *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Inform) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (inform *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Inform) GetSegmentPath() string {
    return "inform"
}

func (inform *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Inform) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inform *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Inform) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inform *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Inform) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = inform.ReceivedPackets
    leafs["transmitted-packets"] = inform.TransmittedPackets
    leafs["dropped-packets"] = inform.DroppedPackets
    return leafs
}

func (inform *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Inform) GetBundleName() string { return "cisco_ios_xr" }

func (inform *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Inform) GetYangName() string { return "inform" }

func (inform *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Inform) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inform *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Inform) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inform *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Inform) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inform *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Inform) SetParent(parent types.Entity) { inform.parent = parent }

func (inform *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Inform) GetParent() types.Entity { return inform.parent }

func (inform *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Inform) GetParentYangName() string { return "statistics" }

// Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQuery
// DHCP lease query packets
type Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQuery struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Received packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPackets interface{}

    // Transmitted packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedPackets interface{}

    // Dropped packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}
}

func (leaseQuery *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQuery) GetFilter() yfilter.YFilter { return leaseQuery.YFilter }

func (leaseQuery *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQuery) SetFilter(yf yfilter.YFilter) { leaseQuery.YFilter = yf }

func (leaseQuery *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQuery) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (leaseQuery *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQuery) GetSegmentPath() string {
    return "lease-query"
}

func (leaseQuery *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQuery) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (leaseQuery *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQuery) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (leaseQuery *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQuery) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = leaseQuery.ReceivedPackets
    leafs["transmitted-packets"] = leaseQuery.TransmittedPackets
    leafs["dropped-packets"] = leaseQuery.DroppedPackets
    return leafs
}

func (leaseQuery *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQuery) GetBundleName() string { return "cisco_ios_xr" }

func (leaseQuery *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQuery) GetYangName() string { return "lease-query" }

func (leaseQuery *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQuery) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (leaseQuery *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQuery) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (leaseQuery *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQuery) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (leaseQuery *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQuery) SetParent(parent types.Entity) { leaseQuery.parent = parent }

func (leaseQuery *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQuery) GetParent() types.Entity { return leaseQuery.parent }

func (leaseQuery *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQuery) GetParentYangName() string { return "statistics" }

// Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseNotAssigned
// DHCP lease not assigned
type Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseNotAssigned struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Received packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPackets interface{}

    // Transmitted packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedPackets interface{}

    // Dropped packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}
}

func (leaseNotAssigned *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseNotAssigned) GetFilter() yfilter.YFilter { return leaseNotAssigned.YFilter }

func (leaseNotAssigned *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseNotAssigned) SetFilter(yf yfilter.YFilter) { leaseNotAssigned.YFilter = yf }

func (leaseNotAssigned *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseNotAssigned) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (leaseNotAssigned *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseNotAssigned) GetSegmentPath() string {
    return "lease-not-assigned"
}

func (leaseNotAssigned *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseNotAssigned) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (leaseNotAssigned *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseNotAssigned) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (leaseNotAssigned *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseNotAssigned) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = leaseNotAssigned.ReceivedPackets
    leafs["transmitted-packets"] = leaseNotAssigned.TransmittedPackets
    leafs["dropped-packets"] = leaseNotAssigned.DroppedPackets
    return leafs
}

func (leaseNotAssigned *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseNotAssigned) GetBundleName() string { return "cisco_ios_xr" }

func (leaseNotAssigned *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseNotAssigned) GetYangName() string { return "lease-not-assigned" }

func (leaseNotAssigned *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseNotAssigned) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (leaseNotAssigned *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseNotAssigned) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (leaseNotAssigned *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseNotAssigned) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (leaseNotAssigned *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseNotAssigned) SetParent(parent types.Entity) { leaseNotAssigned.parent = parent }

func (leaseNotAssigned *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseNotAssigned) GetParent() types.Entity { return leaseNotAssigned.parent }

func (leaseNotAssigned *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseNotAssigned) GetParentYangName() string { return "statistics" }

// Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseUnknown
// DHCP lease unknown
type Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseUnknown struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Received packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPackets interface{}

    // Transmitted packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedPackets interface{}

    // Dropped packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}
}

func (leaseUnknown *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseUnknown) GetFilter() yfilter.YFilter { return leaseUnknown.YFilter }

func (leaseUnknown *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseUnknown) SetFilter(yf yfilter.YFilter) { leaseUnknown.YFilter = yf }

func (leaseUnknown *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseUnknown) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (leaseUnknown *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseUnknown) GetSegmentPath() string {
    return "lease-unknown"
}

func (leaseUnknown *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseUnknown) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (leaseUnknown *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseUnknown) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (leaseUnknown *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseUnknown) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = leaseUnknown.ReceivedPackets
    leafs["transmitted-packets"] = leaseUnknown.TransmittedPackets
    leafs["dropped-packets"] = leaseUnknown.DroppedPackets
    return leafs
}

func (leaseUnknown *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseUnknown) GetBundleName() string { return "cisco_ios_xr" }

func (leaseUnknown *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseUnknown) GetYangName() string { return "lease-unknown" }

func (leaseUnknown *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseUnknown) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (leaseUnknown *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseUnknown) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (leaseUnknown *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseUnknown) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (leaseUnknown *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseUnknown) SetParent(parent types.Entity) { leaseUnknown.parent = parent }

func (leaseUnknown *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseUnknown) GetParent() types.Entity { return leaseUnknown.parent }

func (leaseUnknown *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseUnknown) GetParentYangName() string { return "statistics" }

// Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseActive
// DHCP lease active
type Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseActive struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Received packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPackets interface{}

    // Transmitted packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedPackets interface{}

    // Dropped packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}
}

func (leaseActive *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseActive) GetFilter() yfilter.YFilter { return leaseActive.YFilter }

func (leaseActive *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseActive) SetFilter(yf yfilter.YFilter) { leaseActive.YFilter = yf }

func (leaseActive *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseActive) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (leaseActive *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseActive) GetSegmentPath() string {
    return "lease-active"
}

func (leaseActive *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseActive) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (leaseActive *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseActive) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (leaseActive *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseActive) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = leaseActive.ReceivedPackets
    leafs["transmitted-packets"] = leaseActive.TransmittedPackets
    leafs["dropped-packets"] = leaseActive.DroppedPackets
    return leafs
}

func (leaseActive *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseActive) GetBundleName() string { return "cisco_ios_xr" }

func (leaseActive *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseActive) GetYangName() string { return "lease-active" }

func (leaseActive *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseActive) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (leaseActive *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseActive) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (leaseActive *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseActive) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (leaseActive *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseActive) SetParent(parent types.Entity) { leaseActive.parent = parent }

func (leaseActive *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseActive) GetParent() types.Entity { return leaseActive.parent }

func (leaseActive *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseActive) GetParentYangName() string { return "statistics" }

// Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_BootpRequest
// DHCP BOOTP request
type Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_BootpRequest struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Received packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPackets interface{}

    // Transmitted packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedPackets interface{}

    // Dropped packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}
}

func (bootpRequest *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_BootpRequest) GetFilter() yfilter.YFilter { return bootpRequest.YFilter }

func (bootpRequest *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_BootpRequest) SetFilter(yf yfilter.YFilter) { bootpRequest.YFilter = yf }

func (bootpRequest *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_BootpRequest) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (bootpRequest *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_BootpRequest) GetSegmentPath() string {
    return "bootp-request"
}

func (bootpRequest *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_BootpRequest) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bootpRequest *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_BootpRequest) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bootpRequest *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_BootpRequest) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = bootpRequest.ReceivedPackets
    leafs["transmitted-packets"] = bootpRequest.TransmittedPackets
    leafs["dropped-packets"] = bootpRequest.DroppedPackets
    return leafs
}

func (bootpRequest *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_BootpRequest) GetBundleName() string { return "cisco_ios_xr" }

func (bootpRequest *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_BootpRequest) GetYangName() string { return "bootp-request" }

func (bootpRequest *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_BootpRequest) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bootpRequest *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_BootpRequest) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bootpRequest *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_BootpRequest) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bootpRequest *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_BootpRequest) SetParent(parent types.Entity) { bootpRequest.parent = parent }

func (bootpRequest *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_BootpRequest) GetParent() types.Entity { return bootpRequest.parent }

func (bootpRequest *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_BootpRequest) GetParentYangName() string { return "statistics" }

// Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_BootpReply
// DHCP BOOTP reply
type Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_BootpReply struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Received packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPackets interface{}

    // Transmitted packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedPackets interface{}

    // Dropped packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}
}

func (bootpReply *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_BootpReply) GetFilter() yfilter.YFilter { return bootpReply.YFilter }

func (bootpReply *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_BootpReply) SetFilter(yf yfilter.YFilter) { bootpReply.YFilter = yf }

func (bootpReply *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_BootpReply) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (bootpReply *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_BootpReply) GetSegmentPath() string {
    return "bootp-reply"
}

func (bootpReply *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_BootpReply) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bootpReply *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_BootpReply) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bootpReply *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_BootpReply) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = bootpReply.ReceivedPackets
    leafs["transmitted-packets"] = bootpReply.TransmittedPackets
    leafs["dropped-packets"] = bootpReply.DroppedPackets
    return leafs
}

func (bootpReply *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_BootpReply) GetBundleName() string { return "cisco_ios_xr" }

func (bootpReply *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_BootpReply) GetYangName() string { return "bootp-reply" }

func (bootpReply *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_BootpReply) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bootpReply *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_BootpReply) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bootpReply *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_BootpReply) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bootpReply *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_BootpReply) SetParent(parent types.Entity) { bootpReply.parent = parent }

func (bootpReply *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_BootpReply) GetParent() types.Entity { return bootpReply.parent }

func (bootpReply *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_BootpReply) GetParentYangName() string { return "statistics" }

// Ipv4Dhcpd_Nodes_Node_Relay
// IPv4 DHCPD Relay operational data
type Ipv4Dhcpd_Nodes_Node_Relay struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DHCP Relay Profiles.
    Profiles Ipv4Dhcpd_Nodes_Node_Relay_Profiles

    // DHCP relay statistics info.
    StatisticsInfo Ipv4Dhcpd_Nodes_Node_Relay_StatisticsInfo

    // DHCP Relay VRF statistics.
    Statistics Ipv4Dhcpd_Nodes_Node_Relay_Statistics

    // DHCP relay list of VRF names.
    Vrfs Ipv4Dhcpd_Nodes_Node_Relay_Vrfs
}

func (relay *Ipv4Dhcpd_Nodes_Node_Relay) GetFilter() yfilter.YFilter { return relay.YFilter }

func (relay *Ipv4Dhcpd_Nodes_Node_Relay) SetFilter(yf yfilter.YFilter) { relay.YFilter = yf }

func (relay *Ipv4Dhcpd_Nodes_Node_Relay) GetGoName(yname string) string {
    if yname == "profiles" { return "Profiles" }
    if yname == "statistics-info" { return "StatisticsInfo" }
    if yname == "statistics" { return "Statistics" }
    if yname == "vrfs" { return "Vrfs" }
    return ""
}

func (relay *Ipv4Dhcpd_Nodes_Node_Relay) GetSegmentPath() string {
    return "relay"
}

func (relay *Ipv4Dhcpd_Nodes_Node_Relay) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "profiles" {
        return &relay.Profiles
    }
    if childYangName == "statistics-info" {
        return &relay.StatisticsInfo
    }
    if childYangName == "statistics" {
        return &relay.Statistics
    }
    if childYangName == "vrfs" {
        return &relay.Vrfs
    }
    return nil
}

func (relay *Ipv4Dhcpd_Nodes_Node_Relay) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["profiles"] = &relay.Profiles
    children["statistics-info"] = &relay.StatisticsInfo
    children["statistics"] = &relay.Statistics
    children["vrfs"] = &relay.Vrfs
    return children
}

func (relay *Ipv4Dhcpd_Nodes_Node_Relay) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (relay *Ipv4Dhcpd_Nodes_Node_Relay) GetBundleName() string { return "cisco_ios_xr" }

func (relay *Ipv4Dhcpd_Nodes_Node_Relay) GetYangName() string { return "relay" }

func (relay *Ipv4Dhcpd_Nodes_Node_Relay) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (relay *Ipv4Dhcpd_Nodes_Node_Relay) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (relay *Ipv4Dhcpd_Nodes_Node_Relay) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (relay *Ipv4Dhcpd_Nodes_Node_Relay) SetParent(parent types.Entity) { relay.parent = parent }

func (relay *Ipv4Dhcpd_Nodes_Node_Relay) GetParent() types.Entity { return relay.parent }

func (relay *Ipv4Dhcpd_Nodes_Node_Relay) GetParentYangName() string { return "node" }

// Ipv4Dhcpd_Nodes_Node_Relay_Profiles
// DHCP Relay Profiles
type Ipv4Dhcpd_Nodes_Node_Relay_Profiles struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DHCP Relay profile. The type is slice of
    // Ipv4Dhcpd_Nodes_Node_Relay_Profiles_Profile.
    Profile []Ipv4Dhcpd_Nodes_Node_Relay_Profiles_Profile
}

func (profiles *Ipv4Dhcpd_Nodes_Node_Relay_Profiles) GetFilter() yfilter.YFilter { return profiles.YFilter }

func (profiles *Ipv4Dhcpd_Nodes_Node_Relay_Profiles) SetFilter(yf yfilter.YFilter) { profiles.YFilter = yf }

func (profiles *Ipv4Dhcpd_Nodes_Node_Relay_Profiles) GetGoName(yname string) string {
    if yname == "profile" { return "Profile" }
    return ""
}

func (profiles *Ipv4Dhcpd_Nodes_Node_Relay_Profiles) GetSegmentPath() string {
    return "profiles"
}

func (profiles *Ipv4Dhcpd_Nodes_Node_Relay_Profiles) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "profile" {
        for _, c := range profiles.Profile {
            if profiles.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4Dhcpd_Nodes_Node_Relay_Profiles_Profile{}
        profiles.Profile = append(profiles.Profile, child)
        return &profiles.Profile[len(profiles.Profile)-1]
    }
    return nil
}

func (profiles *Ipv4Dhcpd_Nodes_Node_Relay_Profiles) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range profiles.Profile {
        children[profiles.Profile[i].GetSegmentPath()] = &profiles.Profile[i]
    }
    return children
}

func (profiles *Ipv4Dhcpd_Nodes_Node_Relay_Profiles) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (profiles *Ipv4Dhcpd_Nodes_Node_Relay_Profiles) GetBundleName() string { return "cisco_ios_xr" }

func (profiles *Ipv4Dhcpd_Nodes_Node_Relay_Profiles) GetYangName() string { return "profiles" }

func (profiles *Ipv4Dhcpd_Nodes_Node_Relay_Profiles) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (profiles *Ipv4Dhcpd_Nodes_Node_Relay_Profiles) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (profiles *Ipv4Dhcpd_Nodes_Node_Relay_Profiles) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (profiles *Ipv4Dhcpd_Nodes_Node_Relay_Profiles) SetParent(parent types.Entity) { profiles.parent = parent }

func (profiles *Ipv4Dhcpd_Nodes_Node_Relay_Profiles) GetParent() types.Entity { return profiles.parent }

func (profiles *Ipv4Dhcpd_Nodes_Node_Relay_Profiles) GetParentYangName() string { return "relay" }

// Ipv4Dhcpd_Nodes_Node_Relay_Profiles_Profile
// DHCP Relay profile
type Ipv4Dhcpd_Nodes_Node_Relay_Profiles_Profile struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Profile name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    ProfileName interface{}

    // Profile Name. The type is string with length: 0..65.
    RelayProfileName interface{}

    // Profile UID. The type is interface{} with range: 0..4294967295.
    RelayProfileUid interface{}

    // Helper address count. The type is interface{} with range: 0..255.
    RelayProfileHelperCount interface{}

    // Relay info option. The type is interface{} with range: 0..255.
    RelayProfileRelayInfoOption interface{}

    // Relay info policy. The type is interface{} with range: 0..255.
    RelayProfileRelayInfoPolicy interface{}

    // Relay info untrusted. The type is interface{} with range: 0..255.
    RelayProfileRelayInfoAllowUntrusted interface{}

    // Relay info option vpn. The type is interface{} with range: 0..255.
    RelayProfileRelayInfoOptionvpn interface{}

    // Relay info option vpn-mode. The type is RelayInfoVpnMode.
    RelayProfileRelayInfoOptionvpnMode interface{}

    // Relay info check. The type is interface{} with range: 0..255.
    RelayProfileRelayInfoCheck interface{}

    // GIADDR policy. The type is interface{} with range: 0..255.
    RelayProfileGiAddrPolicy interface{}

    // Broadcast policy. The type is interface{} with range: 0..255.
    RelayProfileBroadcastFlagPolicy interface{}

    // Mac Mismatch Action. The type is interface{} with range: 0..255.
    RelayProfileMacMismatchAction interface{}

    // Helper addresses. The type is slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RelayProfileHelperAddress []interface{}

    // Helper address vrfs. The type is slice of string with length: 0..33.
    RelayProfileHelperVrf []interface{}

    // Gateway addresses. The type is slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RelayProfileGiAddr []interface{}
}

func (profile *Ipv4Dhcpd_Nodes_Node_Relay_Profiles_Profile) GetFilter() yfilter.YFilter { return profile.YFilter }

func (profile *Ipv4Dhcpd_Nodes_Node_Relay_Profiles_Profile) SetFilter(yf yfilter.YFilter) { profile.YFilter = yf }

func (profile *Ipv4Dhcpd_Nodes_Node_Relay_Profiles_Profile) GetGoName(yname string) string {
    if yname == "profile-name" { return "ProfileName" }
    if yname == "relay-profile-name" { return "RelayProfileName" }
    if yname == "relay-profile-uid" { return "RelayProfileUid" }
    if yname == "relay-profile-helper-count" { return "RelayProfileHelperCount" }
    if yname == "relay-profile-relay-info-option" { return "RelayProfileRelayInfoOption" }
    if yname == "relay-profile-relay-info-policy" { return "RelayProfileRelayInfoPolicy" }
    if yname == "relay-profile-relay-info-allow-untrusted" { return "RelayProfileRelayInfoAllowUntrusted" }
    if yname == "relay-profile-relay-info-optionvpn" { return "RelayProfileRelayInfoOptionvpn" }
    if yname == "relay-profile-relay-info-optionvpn-mode" { return "RelayProfileRelayInfoOptionvpnMode" }
    if yname == "relay-profile-relay-info-check" { return "RelayProfileRelayInfoCheck" }
    if yname == "relay-profile-gi-addr-policy" { return "RelayProfileGiAddrPolicy" }
    if yname == "relay-profile-broadcast-flag-policy" { return "RelayProfileBroadcastFlagPolicy" }
    if yname == "relay-profile-mac-mismatch-action" { return "RelayProfileMacMismatchAction" }
    if yname == "relay-profile-helper-address" { return "RelayProfileHelperAddress" }
    if yname == "relay-profile-helper-vrf" { return "RelayProfileHelperVrf" }
    if yname == "relay-profile-gi-addr" { return "RelayProfileGiAddr" }
    return ""
}

func (profile *Ipv4Dhcpd_Nodes_Node_Relay_Profiles_Profile) GetSegmentPath() string {
    return "profile" + "[profile-name='" + fmt.Sprintf("%v", profile.ProfileName) + "']"
}

func (profile *Ipv4Dhcpd_Nodes_Node_Relay_Profiles_Profile) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (profile *Ipv4Dhcpd_Nodes_Node_Relay_Profiles_Profile) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (profile *Ipv4Dhcpd_Nodes_Node_Relay_Profiles_Profile) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["profile-name"] = profile.ProfileName
    leafs["relay-profile-name"] = profile.RelayProfileName
    leafs["relay-profile-uid"] = profile.RelayProfileUid
    leafs["relay-profile-helper-count"] = profile.RelayProfileHelperCount
    leafs["relay-profile-relay-info-option"] = profile.RelayProfileRelayInfoOption
    leafs["relay-profile-relay-info-policy"] = profile.RelayProfileRelayInfoPolicy
    leafs["relay-profile-relay-info-allow-untrusted"] = profile.RelayProfileRelayInfoAllowUntrusted
    leafs["relay-profile-relay-info-optionvpn"] = profile.RelayProfileRelayInfoOptionvpn
    leafs["relay-profile-relay-info-optionvpn-mode"] = profile.RelayProfileRelayInfoOptionvpnMode
    leafs["relay-profile-relay-info-check"] = profile.RelayProfileRelayInfoCheck
    leafs["relay-profile-gi-addr-policy"] = profile.RelayProfileGiAddrPolicy
    leafs["relay-profile-broadcast-flag-policy"] = profile.RelayProfileBroadcastFlagPolicy
    leafs["relay-profile-mac-mismatch-action"] = profile.RelayProfileMacMismatchAction
    leafs["relay-profile-helper-address"] = profile.RelayProfileHelperAddress
    leafs["relay-profile-helper-vrf"] = profile.RelayProfileHelperVrf
    leafs["relay-profile-gi-addr"] = profile.RelayProfileGiAddr
    return leafs
}

func (profile *Ipv4Dhcpd_Nodes_Node_Relay_Profiles_Profile) GetBundleName() string { return "cisco_ios_xr" }

func (profile *Ipv4Dhcpd_Nodes_Node_Relay_Profiles_Profile) GetYangName() string { return "profile" }

func (profile *Ipv4Dhcpd_Nodes_Node_Relay_Profiles_Profile) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (profile *Ipv4Dhcpd_Nodes_Node_Relay_Profiles_Profile) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (profile *Ipv4Dhcpd_Nodes_Node_Relay_Profiles_Profile) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (profile *Ipv4Dhcpd_Nodes_Node_Relay_Profiles_Profile) SetParent(parent types.Entity) { profile.parent = parent }

func (profile *Ipv4Dhcpd_Nodes_Node_Relay_Profiles_Profile) GetParent() types.Entity { return profile.parent }

func (profile *Ipv4Dhcpd_Nodes_Node_Relay_Profiles_Profile) GetParentYangName() string { return "profiles" }

// Ipv4Dhcpd_Nodes_Node_Relay_StatisticsInfo
// DHCP relay statistics info
type Ipv4Dhcpd_Nodes_Node_Relay_StatisticsInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Relay Stats timestamp. The type is interface{} with range: 0..4294967295.
    RelayStatsTimestamp interface{}
}

func (statisticsInfo *Ipv4Dhcpd_Nodes_Node_Relay_StatisticsInfo) GetFilter() yfilter.YFilter { return statisticsInfo.YFilter }

func (statisticsInfo *Ipv4Dhcpd_Nodes_Node_Relay_StatisticsInfo) SetFilter(yf yfilter.YFilter) { statisticsInfo.YFilter = yf }

func (statisticsInfo *Ipv4Dhcpd_Nodes_Node_Relay_StatisticsInfo) GetGoName(yname string) string {
    if yname == "relay-stats-timestamp" { return "RelayStatsTimestamp" }
    return ""
}

func (statisticsInfo *Ipv4Dhcpd_Nodes_Node_Relay_StatisticsInfo) GetSegmentPath() string {
    return "statistics-info"
}

func (statisticsInfo *Ipv4Dhcpd_Nodes_Node_Relay_StatisticsInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (statisticsInfo *Ipv4Dhcpd_Nodes_Node_Relay_StatisticsInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (statisticsInfo *Ipv4Dhcpd_Nodes_Node_Relay_StatisticsInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["relay-stats-timestamp"] = statisticsInfo.RelayStatsTimestamp
    return leafs
}

func (statisticsInfo *Ipv4Dhcpd_Nodes_Node_Relay_StatisticsInfo) GetBundleName() string { return "cisco_ios_xr" }

func (statisticsInfo *Ipv4Dhcpd_Nodes_Node_Relay_StatisticsInfo) GetYangName() string { return "statistics-info" }

func (statisticsInfo *Ipv4Dhcpd_Nodes_Node_Relay_StatisticsInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statisticsInfo *Ipv4Dhcpd_Nodes_Node_Relay_StatisticsInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statisticsInfo *Ipv4Dhcpd_Nodes_Node_Relay_StatisticsInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statisticsInfo *Ipv4Dhcpd_Nodes_Node_Relay_StatisticsInfo) SetParent(parent types.Entity) { statisticsInfo.parent = parent }

func (statisticsInfo *Ipv4Dhcpd_Nodes_Node_Relay_StatisticsInfo) GetParent() types.Entity { return statisticsInfo.parent }

func (statisticsInfo *Ipv4Dhcpd_Nodes_Node_Relay_StatisticsInfo) GetParentYangName() string { return "relay" }

// Ipv4Dhcpd_Nodes_Node_Relay_Statistics
// DHCP Relay VRF statistics
type Ipv4Dhcpd_Nodes_Node_Relay_Statistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ipv4 dhcpd relay stat. The type is slice of
    // Ipv4Dhcpd_Nodes_Node_Relay_Statistics_Ipv4DhcpdRelayStat.
    Ipv4DhcpdRelayStat []Ipv4Dhcpd_Nodes_Node_Relay_Statistics_Ipv4DhcpdRelayStat
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Relay_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *Ipv4Dhcpd_Nodes_Node_Relay_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *Ipv4Dhcpd_Nodes_Node_Relay_Statistics) GetGoName(yname string) string {
    if yname == "ipv4-dhcpd-relay-stat" { return "Ipv4DhcpdRelayStat" }
    return ""
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Relay_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Relay_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv4-dhcpd-relay-stat" {
        for _, c := range statistics.Ipv4DhcpdRelayStat {
            if statistics.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4Dhcpd_Nodes_Node_Relay_Statistics_Ipv4DhcpdRelayStat{}
        statistics.Ipv4DhcpdRelayStat = append(statistics.Ipv4DhcpdRelayStat, child)
        return &statistics.Ipv4DhcpdRelayStat[len(statistics.Ipv4DhcpdRelayStat)-1]
    }
    return nil
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Relay_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range statistics.Ipv4DhcpdRelayStat {
        children[statistics.Ipv4DhcpdRelayStat[i].GetSegmentPath()] = &statistics.Ipv4DhcpdRelayStat[i]
    }
    return children
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Relay_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Relay_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *Ipv4Dhcpd_Nodes_Node_Relay_Statistics) GetYangName() string { return "statistics" }

func (statistics *Ipv4Dhcpd_Nodes_Node_Relay_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *Ipv4Dhcpd_Nodes_Node_Relay_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *Ipv4Dhcpd_Nodes_Node_Relay_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *Ipv4Dhcpd_Nodes_Node_Relay_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *Ipv4Dhcpd_Nodes_Node_Relay_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *Ipv4Dhcpd_Nodes_Node_Relay_Statistics) GetParentYangName() string { return "relay" }

// Ipv4Dhcpd_Nodes_Node_Relay_Statistics_Ipv4DhcpdRelayStat
// ipv4 dhcpd relay stat
type Ipv4Dhcpd_Nodes_Node_Relay_Statistics_Ipv4DhcpdRelayStat struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DHCP L3 VRF Name. The type is string with length: 0..33.
    RelayStatisticsVrfName interface{}

    // Public relay statistics.
    Statistics Ipv4Dhcpd_Nodes_Node_Relay_Statistics_Ipv4DhcpdRelayStat_Statistics
}

func (ipv4DhcpdRelayStat *Ipv4Dhcpd_Nodes_Node_Relay_Statistics_Ipv4DhcpdRelayStat) GetFilter() yfilter.YFilter { return ipv4DhcpdRelayStat.YFilter }

func (ipv4DhcpdRelayStat *Ipv4Dhcpd_Nodes_Node_Relay_Statistics_Ipv4DhcpdRelayStat) SetFilter(yf yfilter.YFilter) { ipv4DhcpdRelayStat.YFilter = yf }

func (ipv4DhcpdRelayStat *Ipv4Dhcpd_Nodes_Node_Relay_Statistics_Ipv4DhcpdRelayStat) GetGoName(yname string) string {
    if yname == "relay-statistics-vrf-name" { return "RelayStatisticsVrfName" }
    if yname == "statistics" { return "Statistics" }
    return ""
}

func (ipv4DhcpdRelayStat *Ipv4Dhcpd_Nodes_Node_Relay_Statistics_Ipv4DhcpdRelayStat) GetSegmentPath() string {
    return "ipv4-dhcpd-relay-stat"
}

func (ipv4DhcpdRelayStat *Ipv4Dhcpd_Nodes_Node_Relay_Statistics_Ipv4DhcpdRelayStat) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "statistics" {
        return &ipv4DhcpdRelayStat.Statistics
    }
    return nil
}

func (ipv4DhcpdRelayStat *Ipv4Dhcpd_Nodes_Node_Relay_Statistics_Ipv4DhcpdRelayStat) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["statistics"] = &ipv4DhcpdRelayStat.Statistics
    return children
}

func (ipv4DhcpdRelayStat *Ipv4Dhcpd_Nodes_Node_Relay_Statistics_Ipv4DhcpdRelayStat) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["relay-statistics-vrf-name"] = ipv4DhcpdRelayStat.RelayStatisticsVrfName
    return leafs
}

func (ipv4DhcpdRelayStat *Ipv4Dhcpd_Nodes_Node_Relay_Statistics_Ipv4DhcpdRelayStat) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4DhcpdRelayStat *Ipv4Dhcpd_Nodes_Node_Relay_Statistics_Ipv4DhcpdRelayStat) GetYangName() string { return "ipv4-dhcpd-relay-stat" }

func (ipv4DhcpdRelayStat *Ipv4Dhcpd_Nodes_Node_Relay_Statistics_Ipv4DhcpdRelayStat) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4DhcpdRelayStat *Ipv4Dhcpd_Nodes_Node_Relay_Statistics_Ipv4DhcpdRelayStat) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4DhcpdRelayStat *Ipv4Dhcpd_Nodes_Node_Relay_Statistics_Ipv4DhcpdRelayStat) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4DhcpdRelayStat *Ipv4Dhcpd_Nodes_Node_Relay_Statistics_Ipv4DhcpdRelayStat) SetParent(parent types.Entity) { ipv4DhcpdRelayStat.parent = parent }

func (ipv4DhcpdRelayStat *Ipv4Dhcpd_Nodes_Node_Relay_Statistics_Ipv4DhcpdRelayStat) GetParent() types.Entity { return ipv4DhcpdRelayStat.parent }

func (ipv4DhcpdRelayStat *Ipv4Dhcpd_Nodes_Node_Relay_Statistics_Ipv4DhcpdRelayStat) GetParentYangName() string { return "statistics" }

// Ipv4Dhcpd_Nodes_Node_Relay_Statistics_Ipv4DhcpdRelayStat_Statistics
// Public relay statistics
type Ipv4Dhcpd_Nodes_Node_Relay_Statistics_Ipv4DhcpdRelayStat_Statistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Received packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPackets interface{}

    // Transmitted packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedPackets interface{}

    // Dropped packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Relay_Statistics_Ipv4DhcpdRelayStat_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *Ipv4Dhcpd_Nodes_Node_Relay_Statistics_Ipv4DhcpdRelayStat_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *Ipv4Dhcpd_Nodes_Node_Relay_Statistics_Ipv4DhcpdRelayStat_Statistics) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Relay_Statistics_Ipv4DhcpdRelayStat_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Relay_Statistics_Ipv4DhcpdRelayStat_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Relay_Statistics_Ipv4DhcpdRelayStat_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Relay_Statistics_Ipv4DhcpdRelayStat_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = statistics.ReceivedPackets
    leafs["transmitted-packets"] = statistics.TransmittedPackets
    leafs["dropped-packets"] = statistics.DroppedPackets
    return leafs
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Relay_Statistics_Ipv4DhcpdRelayStat_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *Ipv4Dhcpd_Nodes_Node_Relay_Statistics_Ipv4DhcpdRelayStat_Statistics) GetYangName() string { return "statistics" }

func (statistics *Ipv4Dhcpd_Nodes_Node_Relay_Statistics_Ipv4DhcpdRelayStat_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *Ipv4Dhcpd_Nodes_Node_Relay_Statistics_Ipv4DhcpdRelayStat_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *Ipv4Dhcpd_Nodes_Node_Relay_Statistics_Ipv4DhcpdRelayStat_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *Ipv4Dhcpd_Nodes_Node_Relay_Statistics_Ipv4DhcpdRelayStat_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *Ipv4Dhcpd_Nodes_Node_Relay_Statistics_Ipv4DhcpdRelayStat_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *Ipv4Dhcpd_Nodes_Node_Relay_Statistics_Ipv4DhcpdRelayStat_Statistics) GetParentYangName() string { return "ipv4-dhcpd-relay-stat" }

// Ipv4Dhcpd_Nodes_Node_Relay_Vrfs
// DHCP relay list of VRF names
type Ipv4Dhcpd_Nodes_Node_Relay_Vrfs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 DHCP relay VRF name. The type is slice of
    // Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf.
    Vrf []Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf
}

func (vrfs *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs) GetFilter() yfilter.YFilter { return vrfs.YFilter }

func (vrfs *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs) SetFilter(yf yfilter.YFilter) { vrfs.YFilter = yf }

func (vrfs *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs) GetGoName(yname string) string {
    if yname == "vrf" { return "Vrf" }
    return ""
}

func (vrfs *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs) GetSegmentPath() string {
    return "vrfs"
}

func (vrfs *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf" {
        for _, c := range vrfs.Vrf {
            if vrfs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf{}
        vrfs.Vrf = append(vrfs.Vrf, child)
        return &vrfs.Vrf[len(vrfs.Vrf)-1]
    }
    return nil
}

func (vrfs *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vrfs.Vrf {
        children[vrfs.Vrf[i].GetSegmentPath()] = &vrfs.Vrf[i]
    }
    return children
}

func (vrfs *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vrfs *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs) GetBundleName() string { return "cisco_ios_xr" }

func (vrfs *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs) GetYangName() string { return "vrfs" }

func (vrfs *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfs *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfs *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfs *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs) SetParent(parent types.Entity) { vrfs.parent = parent }

func (vrfs *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs) GetParent() types.Entity { return vrfs.parent }

func (vrfs *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs) GetParentYangName() string { return "relay" }

// Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf
// IPv4 DHCP relay VRF name
type Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // IPv4 DHCP relay statistics.
    VrfStatistics Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics
}

func (vrf *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf) GetFilter() yfilter.YFilter { return vrf.YFilter }

func (vrf *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf) SetFilter(yf yfilter.YFilter) { vrf.YFilter = yf }

func (vrf *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "vrf-statistics" { return "VrfStatistics" }
    return ""
}

func (vrf *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf) GetSegmentPath() string {
    return "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
}

func (vrf *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf-statistics" {
        return &vrf.VrfStatistics
    }
    return nil
}

func (vrf *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["vrf-statistics"] = &vrf.VrfStatistics
    return children
}

func (vrf *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = vrf.VrfName
    return leafs
}

func (vrf *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf) GetBundleName() string { return "cisco_ios_xr" }

func (vrf *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf) GetYangName() string { return "vrf" }

func (vrf *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrf *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrf *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrf *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf) SetParent(parent types.Entity) { vrf.parent = parent }

func (vrf *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf) GetParent() types.Entity { return vrf.parent }

func (vrf *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf) GetParentYangName() string { return "vrfs" }

// Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics
// IPv4 DHCP relay statistics
type Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DHCP discover packets.
    Discover Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Discover

    // DHCP offer packets.
    Offer Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Offer

    // DHCP request packets.
    Request Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Request

    // DHCP decline packets.
    Decline Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Decline

    // DHCP ack packets.
    Ack Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Ack

    // DHCP nak packets.
    Nak Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Nak

    // DHCP release packets.
    Release Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Release

    // DHCP inform packets.
    Inform Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Inform

    // DHCP lease query packets.
    LeaseQuery Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseQuery

    // DHCP lease not assigned.
    LeaseNotAssigned Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseNotAssigned

    // DHCP lease unknown.
    LeaseUnknown Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseUnknown

    // DHCP lease active.
    LeaseActive Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseActive

    // DHCP BOOTP request.
    BootpRequest Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_BootpRequest

    // DHCP BOOTP reply.
    BootpReply Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_BootpReply
}

func (vrfStatistics *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics) GetFilter() yfilter.YFilter { return vrfStatistics.YFilter }

func (vrfStatistics *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics) SetFilter(yf yfilter.YFilter) { vrfStatistics.YFilter = yf }

func (vrfStatistics *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics) GetGoName(yname string) string {
    if yname == "discover" { return "Discover" }
    if yname == "offer" { return "Offer" }
    if yname == "request" { return "Request" }
    if yname == "decline" { return "Decline" }
    if yname == "ack" { return "Ack" }
    if yname == "nak" { return "Nak" }
    if yname == "release" { return "Release" }
    if yname == "inform" { return "Inform" }
    if yname == "lease-query" { return "LeaseQuery" }
    if yname == "lease-not-assigned" { return "LeaseNotAssigned" }
    if yname == "lease-unknown" { return "LeaseUnknown" }
    if yname == "lease-active" { return "LeaseActive" }
    if yname == "bootp-request" { return "BootpRequest" }
    if yname == "bootp-reply" { return "BootpReply" }
    return ""
}

func (vrfStatistics *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics) GetSegmentPath() string {
    return "vrf-statistics"
}

func (vrfStatistics *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "discover" {
        return &vrfStatistics.Discover
    }
    if childYangName == "offer" {
        return &vrfStatistics.Offer
    }
    if childYangName == "request" {
        return &vrfStatistics.Request
    }
    if childYangName == "decline" {
        return &vrfStatistics.Decline
    }
    if childYangName == "ack" {
        return &vrfStatistics.Ack
    }
    if childYangName == "nak" {
        return &vrfStatistics.Nak
    }
    if childYangName == "release" {
        return &vrfStatistics.Release
    }
    if childYangName == "inform" {
        return &vrfStatistics.Inform
    }
    if childYangName == "lease-query" {
        return &vrfStatistics.LeaseQuery
    }
    if childYangName == "lease-not-assigned" {
        return &vrfStatistics.LeaseNotAssigned
    }
    if childYangName == "lease-unknown" {
        return &vrfStatistics.LeaseUnknown
    }
    if childYangName == "lease-active" {
        return &vrfStatistics.LeaseActive
    }
    if childYangName == "bootp-request" {
        return &vrfStatistics.BootpRequest
    }
    if childYangName == "bootp-reply" {
        return &vrfStatistics.BootpReply
    }
    return nil
}

func (vrfStatistics *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["discover"] = &vrfStatistics.Discover
    children["offer"] = &vrfStatistics.Offer
    children["request"] = &vrfStatistics.Request
    children["decline"] = &vrfStatistics.Decline
    children["ack"] = &vrfStatistics.Ack
    children["nak"] = &vrfStatistics.Nak
    children["release"] = &vrfStatistics.Release
    children["inform"] = &vrfStatistics.Inform
    children["lease-query"] = &vrfStatistics.LeaseQuery
    children["lease-not-assigned"] = &vrfStatistics.LeaseNotAssigned
    children["lease-unknown"] = &vrfStatistics.LeaseUnknown
    children["lease-active"] = &vrfStatistics.LeaseActive
    children["bootp-request"] = &vrfStatistics.BootpRequest
    children["bootp-reply"] = &vrfStatistics.BootpReply
    return children
}

func (vrfStatistics *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vrfStatistics *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (vrfStatistics *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics) GetYangName() string { return "vrf-statistics" }

func (vrfStatistics *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfStatistics *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfStatistics *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfStatistics *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics) SetParent(parent types.Entity) { vrfStatistics.parent = parent }

func (vrfStatistics *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics) GetParent() types.Entity { return vrfStatistics.parent }

func (vrfStatistics *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics) GetParentYangName() string { return "vrf" }

// Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Discover
// DHCP discover packets
type Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Discover struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Received packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPackets interface{}

    // Transmitted packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedPackets interface{}

    // Dropped packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}
}

func (discover *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Discover) GetFilter() yfilter.YFilter { return discover.YFilter }

func (discover *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Discover) SetFilter(yf yfilter.YFilter) { discover.YFilter = yf }

func (discover *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Discover) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (discover *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Discover) GetSegmentPath() string {
    return "discover"
}

func (discover *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Discover) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (discover *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Discover) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (discover *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Discover) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = discover.ReceivedPackets
    leafs["transmitted-packets"] = discover.TransmittedPackets
    leafs["dropped-packets"] = discover.DroppedPackets
    return leafs
}

func (discover *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Discover) GetBundleName() string { return "cisco_ios_xr" }

func (discover *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Discover) GetYangName() string { return "discover" }

func (discover *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Discover) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (discover *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Discover) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (discover *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Discover) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (discover *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Discover) SetParent(parent types.Entity) { discover.parent = parent }

func (discover *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Discover) GetParent() types.Entity { return discover.parent }

func (discover *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Discover) GetParentYangName() string { return "vrf-statistics" }

// Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Offer
// DHCP offer packets
type Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Offer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Received packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPackets interface{}

    // Transmitted packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedPackets interface{}

    // Dropped packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}
}

func (offer *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Offer) GetFilter() yfilter.YFilter { return offer.YFilter }

func (offer *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Offer) SetFilter(yf yfilter.YFilter) { offer.YFilter = yf }

func (offer *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Offer) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (offer *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Offer) GetSegmentPath() string {
    return "offer"
}

func (offer *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Offer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (offer *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Offer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (offer *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Offer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = offer.ReceivedPackets
    leafs["transmitted-packets"] = offer.TransmittedPackets
    leafs["dropped-packets"] = offer.DroppedPackets
    return leafs
}

func (offer *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Offer) GetBundleName() string { return "cisco_ios_xr" }

func (offer *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Offer) GetYangName() string { return "offer" }

func (offer *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Offer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (offer *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Offer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (offer *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Offer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (offer *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Offer) SetParent(parent types.Entity) { offer.parent = parent }

func (offer *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Offer) GetParent() types.Entity { return offer.parent }

func (offer *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Offer) GetParentYangName() string { return "vrf-statistics" }

// Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Request
// DHCP request packets
type Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Request struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Received packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPackets interface{}

    // Transmitted packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedPackets interface{}

    // Dropped packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}
}

func (request *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Request) GetFilter() yfilter.YFilter { return request.YFilter }

func (request *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Request) SetFilter(yf yfilter.YFilter) { request.YFilter = yf }

func (request *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Request) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (request *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Request) GetSegmentPath() string {
    return "request"
}

func (request *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Request) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (request *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Request) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (request *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Request) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = request.ReceivedPackets
    leafs["transmitted-packets"] = request.TransmittedPackets
    leafs["dropped-packets"] = request.DroppedPackets
    return leafs
}

func (request *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Request) GetBundleName() string { return "cisco_ios_xr" }

func (request *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Request) GetYangName() string { return "request" }

func (request *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Request) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (request *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Request) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (request *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Request) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (request *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Request) SetParent(parent types.Entity) { request.parent = parent }

func (request *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Request) GetParent() types.Entity { return request.parent }

func (request *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Request) GetParentYangName() string { return "vrf-statistics" }

// Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Decline
// DHCP decline packets
type Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Decline struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Received packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPackets interface{}

    // Transmitted packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedPackets interface{}

    // Dropped packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}
}

func (decline *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Decline) GetFilter() yfilter.YFilter { return decline.YFilter }

func (decline *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Decline) SetFilter(yf yfilter.YFilter) { decline.YFilter = yf }

func (decline *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Decline) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (decline *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Decline) GetSegmentPath() string {
    return "decline"
}

func (decline *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Decline) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (decline *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Decline) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (decline *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Decline) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = decline.ReceivedPackets
    leafs["transmitted-packets"] = decline.TransmittedPackets
    leafs["dropped-packets"] = decline.DroppedPackets
    return leafs
}

func (decline *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Decline) GetBundleName() string { return "cisco_ios_xr" }

func (decline *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Decline) GetYangName() string { return "decline" }

func (decline *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Decline) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (decline *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Decline) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (decline *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Decline) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (decline *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Decline) SetParent(parent types.Entity) { decline.parent = parent }

func (decline *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Decline) GetParent() types.Entity { return decline.parent }

func (decline *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Decline) GetParentYangName() string { return "vrf-statistics" }

// Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Ack
// DHCP ack packets
type Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Ack struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Received packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPackets interface{}

    // Transmitted packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedPackets interface{}

    // Dropped packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}
}

func (ack *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Ack) GetFilter() yfilter.YFilter { return ack.YFilter }

func (ack *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Ack) SetFilter(yf yfilter.YFilter) { ack.YFilter = yf }

func (ack *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Ack) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (ack *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Ack) GetSegmentPath() string {
    return "ack"
}

func (ack *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Ack) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ack *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Ack) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ack *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Ack) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = ack.ReceivedPackets
    leafs["transmitted-packets"] = ack.TransmittedPackets
    leafs["dropped-packets"] = ack.DroppedPackets
    return leafs
}

func (ack *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Ack) GetBundleName() string { return "cisco_ios_xr" }

func (ack *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Ack) GetYangName() string { return "ack" }

func (ack *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Ack) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ack *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Ack) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ack *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Ack) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ack *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Ack) SetParent(parent types.Entity) { ack.parent = parent }

func (ack *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Ack) GetParent() types.Entity { return ack.parent }

func (ack *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Ack) GetParentYangName() string { return "vrf-statistics" }

// Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Nak
// DHCP nak packets
type Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Nak struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Received packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPackets interface{}

    // Transmitted packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedPackets interface{}

    // Dropped packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}
}

func (nak *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Nak) GetFilter() yfilter.YFilter { return nak.YFilter }

func (nak *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Nak) SetFilter(yf yfilter.YFilter) { nak.YFilter = yf }

func (nak *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Nak) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (nak *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Nak) GetSegmentPath() string {
    return "nak"
}

func (nak *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Nak) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nak *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Nak) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nak *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Nak) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = nak.ReceivedPackets
    leafs["transmitted-packets"] = nak.TransmittedPackets
    leafs["dropped-packets"] = nak.DroppedPackets
    return leafs
}

func (nak *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Nak) GetBundleName() string { return "cisco_ios_xr" }

func (nak *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Nak) GetYangName() string { return "nak" }

func (nak *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Nak) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nak *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Nak) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nak *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Nak) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nak *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Nak) SetParent(parent types.Entity) { nak.parent = parent }

func (nak *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Nak) GetParent() types.Entity { return nak.parent }

func (nak *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Nak) GetParentYangName() string { return "vrf-statistics" }

// Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Release
// DHCP release packets
type Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Release struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Received packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPackets interface{}

    // Transmitted packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedPackets interface{}

    // Dropped packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}
}

func (release *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Release) GetFilter() yfilter.YFilter { return release.YFilter }

func (release *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Release) SetFilter(yf yfilter.YFilter) { release.YFilter = yf }

func (release *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Release) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (release *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Release) GetSegmentPath() string {
    return "release"
}

func (release *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Release) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (release *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Release) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (release *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Release) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = release.ReceivedPackets
    leafs["transmitted-packets"] = release.TransmittedPackets
    leafs["dropped-packets"] = release.DroppedPackets
    return leafs
}

func (release *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Release) GetBundleName() string { return "cisco_ios_xr" }

func (release *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Release) GetYangName() string { return "release" }

func (release *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Release) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (release *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Release) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (release *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Release) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (release *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Release) SetParent(parent types.Entity) { release.parent = parent }

func (release *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Release) GetParent() types.Entity { return release.parent }

func (release *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Release) GetParentYangName() string { return "vrf-statistics" }

// Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Inform
// DHCP inform packets
type Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Inform struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Received packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPackets interface{}

    // Transmitted packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedPackets interface{}

    // Dropped packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}
}

func (inform *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Inform) GetFilter() yfilter.YFilter { return inform.YFilter }

func (inform *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Inform) SetFilter(yf yfilter.YFilter) { inform.YFilter = yf }

func (inform *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Inform) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (inform *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Inform) GetSegmentPath() string {
    return "inform"
}

func (inform *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Inform) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inform *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Inform) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inform *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Inform) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = inform.ReceivedPackets
    leafs["transmitted-packets"] = inform.TransmittedPackets
    leafs["dropped-packets"] = inform.DroppedPackets
    return leafs
}

func (inform *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Inform) GetBundleName() string { return "cisco_ios_xr" }

func (inform *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Inform) GetYangName() string { return "inform" }

func (inform *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Inform) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inform *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Inform) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inform *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Inform) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inform *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Inform) SetParent(parent types.Entity) { inform.parent = parent }

func (inform *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Inform) GetParent() types.Entity { return inform.parent }

func (inform *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Inform) GetParentYangName() string { return "vrf-statistics" }

// Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseQuery
// DHCP lease query packets
type Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseQuery struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Received packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPackets interface{}

    // Transmitted packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedPackets interface{}

    // Dropped packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}
}

func (leaseQuery *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseQuery) GetFilter() yfilter.YFilter { return leaseQuery.YFilter }

func (leaseQuery *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseQuery) SetFilter(yf yfilter.YFilter) { leaseQuery.YFilter = yf }

func (leaseQuery *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseQuery) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (leaseQuery *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseQuery) GetSegmentPath() string {
    return "lease-query"
}

func (leaseQuery *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseQuery) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (leaseQuery *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseQuery) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (leaseQuery *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseQuery) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = leaseQuery.ReceivedPackets
    leafs["transmitted-packets"] = leaseQuery.TransmittedPackets
    leafs["dropped-packets"] = leaseQuery.DroppedPackets
    return leafs
}

func (leaseQuery *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseQuery) GetBundleName() string { return "cisco_ios_xr" }

func (leaseQuery *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseQuery) GetYangName() string { return "lease-query" }

func (leaseQuery *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseQuery) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (leaseQuery *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseQuery) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (leaseQuery *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseQuery) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (leaseQuery *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseQuery) SetParent(parent types.Entity) { leaseQuery.parent = parent }

func (leaseQuery *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseQuery) GetParent() types.Entity { return leaseQuery.parent }

func (leaseQuery *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseQuery) GetParentYangName() string { return "vrf-statistics" }

// Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseNotAssigned
// DHCP lease not assigned
type Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseNotAssigned struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Received packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPackets interface{}

    // Transmitted packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedPackets interface{}

    // Dropped packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}
}

func (leaseNotAssigned *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseNotAssigned) GetFilter() yfilter.YFilter { return leaseNotAssigned.YFilter }

func (leaseNotAssigned *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseNotAssigned) SetFilter(yf yfilter.YFilter) { leaseNotAssigned.YFilter = yf }

func (leaseNotAssigned *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseNotAssigned) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (leaseNotAssigned *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseNotAssigned) GetSegmentPath() string {
    return "lease-not-assigned"
}

func (leaseNotAssigned *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseNotAssigned) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (leaseNotAssigned *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseNotAssigned) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (leaseNotAssigned *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseNotAssigned) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = leaseNotAssigned.ReceivedPackets
    leafs["transmitted-packets"] = leaseNotAssigned.TransmittedPackets
    leafs["dropped-packets"] = leaseNotAssigned.DroppedPackets
    return leafs
}

func (leaseNotAssigned *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseNotAssigned) GetBundleName() string { return "cisco_ios_xr" }

func (leaseNotAssigned *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseNotAssigned) GetYangName() string { return "lease-not-assigned" }

func (leaseNotAssigned *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseNotAssigned) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (leaseNotAssigned *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseNotAssigned) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (leaseNotAssigned *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseNotAssigned) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (leaseNotAssigned *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseNotAssigned) SetParent(parent types.Entity) { leaseNotAssigned.parent = parent }

func (leaseNotAssigned *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseNotAssigned) GetParent() types.Entity { return leaseNotAssigned.parent }

func (leaseNotAssigned *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseNotAssigned) GetParentYangName() string { return "vrf-statistics" }

// Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseUnknown
// DHCP lease unknown
type Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseUnknown struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Received packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPackets interface{}

    // Transmitted packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedPackets interface{}

    // Dropped packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}
}

func (leaseUnknown *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseUnknown) GetFilter() yfilter.YFilter { return leaseUnknown.YFilter }

func (leaseUnknown *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseUnknown) SetFilter(yf yfilter.YFilter) { leaseUnknown.YFilter = yf }

func (leaseUnknown *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseUnknown) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (leaseUnknown *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseUnknown) GetSegmentPath() string {
    return "lease-unknown"
}

func (leaseUnknown *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseUnknown) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (leaseUnknown *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseUnknown) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (leaseUnknown *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseUnknown) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = leaseUnknown.ReceivedPackets
    leafs["transmitted-packets"] = leaseUnknown.TransmittedPackets
    leafs["dropped-packets"] = leaseUnknown.DroppedPackets
    return leafs
}

func (leaseUnknown *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseUnknown) GetBundleName() string { return "cisco_ios_xr" }

func (leaseUnknown *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseUnknown) GetYangName() string { return "lease-unknown" }

func (leaseUnknown *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseUnknown) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (leaseUnknown *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseUnknown) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (leaseUnknown *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseUnknown) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (leaseUnknown *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseUnknown) SetParent(parent types.Entity) { leaseUnknown.parent = parent }

func (leaseUnknown *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseUnknown) GetParent() types.Entity { return leaseUnknown.parent }

func (leaseUnknown *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseUnknown) GetParentYangName() string { return "vrf-statistics" }

// Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseActive
// DHCP lease active
type Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseActive struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Received packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPackets interface{}

    // Transmitted packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedPackets interface{}

    // Dropped packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}
}

func (leaseActive *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseActive) GetFilter() yfilter.YFilter { return leaseActive.YFilter }

func (leaseActive *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseActive) SetFilter(yf yfilter.YFilter) { leaseActive.YFilter = yf }

func (leaseActive *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseActive) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (leaseActive *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseActive) GetSegmentPath() string {
    return "lease-active"
}

func (leaseActive *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseActive) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (leaseActive *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseActive) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (leaseActive *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseActive) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = leaseActive.ReceivedPackets
    leafs["transmitted-packets"] = leaseActive.TransmittedPackets
    leafs["dropped-packets"] = leaseActive.DroppedPackets
    return leafs
}

func (leaseActive *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseActive) GetBundleName() string { return "cisco_ios_xr" }

func (leaseActive *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseActive) GetYangName() string { return "lease-active" }

func (leaseActive *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseActive) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (leaseActive *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseActive) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (leaseActive *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseActive) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (leaseActive *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseActive) SetParent(parent types.Entity) { leaseActive.parent = parent }

func (leaseActive *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseActive) GetParent() types.Entity { return leaseActive.parent }

func (leaseActive *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseActive) GetParentYangName() string { return "vrf-statistics" }

// Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_BootpRequest
// DHCP BOOTP request
type Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_BootpRequest struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Received packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPackets interface{}

    // Transmitted packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedPackets interface{}

    // Dropped packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}
}

func (bootpRequest *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_BootpRequest) GetFilter() yfilter.YFilter { return bootpRequest.YFilter }

func (bootpRequest *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_BootpRequest) SetFilter(yf yfilter.YFilter) { bootpRequest.YFilter = yf }

func (bootpRequest *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_BootpRequest) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (bootpRequest *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_BootpRequest) GetSegmentPath() string {
    return "bootp-request"
}

func (bootpRequest *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_BootpRequest) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bootpRequest *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_BootpRequest) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bootpRequest *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_BootpRequest) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = bootpRequest.ReceivedPackets
    leafs["transmitted-packets"] = bootpRequest.TransmittedPackets
    leafs["dropped-packets"] = bootpRequest.DroppedPackets
    return leafs
}

func (bootpRequest *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_BootpRequest) GetBundleName() string { return "cisco_ios_xr" }

func (bootpRequest *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_BootpRequest) GetYangName() string { return "bootp-request" }

func (bootpRequest *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_BootpRequest) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bootpRequest *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_BootpRequest) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bootpRequest *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_BootpRequest) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bootpRequest *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_BootpRequest) SetParent(parent types.Entity) { bootpRequest.parent = parent }

func (bootpRequest *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_BootpRequest) GetParent() types.Entity { return bootpRequest.parent }

func (bootpRequest *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_BootpRequest) GetParentYangName() string { return "vrf-statistics" }

// Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_BootpReply
// DHCP BOOTP reply
type Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_BootpReply struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Received packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPackets interface{}

    // Transmitted packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedPackets interface{}

    // Dropped packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}
}

func (bootpReply *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_BootpReply) GetFilter() yfilter.YFilter { return bootpReply.YFilter }

func (bootpReply *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_BootpReply) SetFilter(yf yfilter.YFilter) { bootpReply.YFilter = yf }

func (bootpReply *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_BootpReply) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (bootpReply *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_BootpReply) GetSegmentPath() string {
    return "bootp-reply"
}

func (bootpReply *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_BootpReply) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bootpReply *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_BootpReply) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bootpReply *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_BootpReply) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = bootpReply.ReceivedPackets
    leafs["transmitted-packets"] = bootpReply.TransmittedPackets
    leafs["dropped-packets"] = bootpReply.DroppedPackets
    return leafs
}

func (bootpReply *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_BootpReply) GetBundleName() string { return "cisco_ios_xr" }

func (bootpReply *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_BootpReply) GetYangName() string { return "bootp-reply" }

func (bootpReply *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_BootpReply) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bootpReply *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_BootpReply) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bootpReply *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_BootpReply) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bootpReply *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_BootpReply) SetParent(parent types.Entity) { bootpReply.parent = parent }

func (bootpReply *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_BootpReply) GetParent() types.Entity { return bootpReply.parent }

func (bootpReply *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_BootpReply) GetParentYangName() string { return "vrf-statistics" }

