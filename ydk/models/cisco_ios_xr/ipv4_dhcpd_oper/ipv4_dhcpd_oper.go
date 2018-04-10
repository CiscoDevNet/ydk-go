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

// DhcpIssuVersion represents Dhcp issu version
type DhcpIssuVersion string

const (
    // Version 1
    DhcpIssuVersion_version1 DhcpIssuVersion = "version1"

    // Version 2
    DhcpIssuVersion_version2 DhcpIssuVersion = "version2"
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

// DhcpIssuRole represents Dhcp issu role
type DhcpIssuRole string

const (
    // Primary role
    DhcpIssuRole_role_primary DhcpIssuRole = "role-primary"

    // Secondary role
    DhcpIssuRole_role_secondary DhcpIssuRole = "role-secondary"
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

// DhcpClient
// DHCP client operational data
type DhcpClient struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DHCP client list of nodes.
    Nodes DhcpClient_Nodes
}

func (dhcpClient *DhcpClient) GetEntityData() *types.CommonEntityData {
    dhcpClient.EntityData.YFilter = dhcpClient.YFilter
    dhcpClient.EntityData.YangName = "dhcp-client"
    dhcpClient.EntityData.BundleName = "cisco_ios_xr"
    dhcpClient.EntityData.ParentYangName = "Cisco-IOS-XR-ipv4-dhcpd-oper"
    dhcpClient.EntityData.SegmentPath = "Cisco-IOS-XR-ipv4-dhcpd-oper:dhcp-client"
    dhcpClient.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dhcpClient.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dhcpClient.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dhcpClient.EntityData.Children = make(map[string]types.YChild)
    dhcpClient.EntityData.Children["nodes"] = types.YChild{"Nodes", &dhcpClient.Nodes}
    dhcpClient.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dhcpClient.EntityData)
}

// DhcpClient_Nodes
// DHCP client list of nodes
type DhcpClient_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DHCP client particular node name. The type is slice of
    // DhcpClient_Nodes_Node.
    Node []DhcpClient_Nodes_Node
}

func (nodes *DhcpClient_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "dhcp-client"
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

// DhcpClient_Nodes_Node
// DHCP client particular node name
type DhcpClient_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node name. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeName interface{}

    // IPv4 DHCP client statistics table.
    ClientStats DhcpClient_Nodes_Node_ClientStats

    // IPv4 DHCP client table.
    Clients DhcpClient_Nodes_Node_Clients
}

func (node *DhcpClient_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Children["client-stats"] = types.YChild{"ClientStats", &node.ClientStats}
    node.EntityData.Children["clients"] = types.YChild{"Clients", &node.Clients}
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", node.NodeName}
    return &(node.EntityData)
}

// DhcpClient_Nodes_Node_ClientStats
// IPv4 DHCP client statistics table
type DhcpClient_Nodes_Node_ClientStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DHCP client binding statistics. The type is slice of
    // DhcpClient_Nodes_Node_ClientStats_ClientStat.
    ClientStat []DhcpClient_Nodes_Node_ClientStats_ClientStat
}

func (clientStats *DhcpClient_Nodes_Node_ClientStats) GetEntityData() *types.CommonEntityData {
    clientStats.EntityData.YFilter = clientStats.YFilter
    clientStats.EntityData.YangName = "client-stats"
    clientStats.EntityData.BundleName = "cisco_ios_xr"
    clientStats.EntityData.ParentYangName = "node"
    clientStats.EntityData.SegmentPath = "client-stats"
    clientStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clientStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clientStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clientStats.EntityData.Children = make(map[string]types.YChild)
    clientStats.EntityData.Children["client-stat"] = types.YChild{"ClientStat", nil}
    for i := range clientStats.ClientStat {
        clientStats.EntityData.Children[types.GetSegmentPath(&clientStats.ClientStat[i])] = types.YChild{"ClientStat", &clientStats.ClientStat[i]}
    }
    clientStats.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(clientStats.EntityData)
}

// DhcpClient_Nodes_Node_ClientStats_ClientStat
// DHCP client binding statistics
type DhcpClient_Nodes_Node_ClientStats_ClientStat struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Client Ifhandle. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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

    // Number of vbind timer event. The type is interface{} with range:
    // 0..4294967295.
    NumVbindTimerEvent interface{}

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

    // Number of vbind timer starts. The type is interface{} with range:
    // 0..4294967295.
    NumVbindTimerStart interface{}

    // Number of vbind timer stops. The type is interface{} with range:
    // 0..4294967295.
    NumVbindTimerStop interface{}

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

func (clientStat *DhcpClient_Nodes_Node_ClientStats_ClientStat) GetEntityData() *types.CommonEntityData {
    clientStat.EntityData.YFilter = clientStat.YFilter
    clientStat.EntityData.YangName = "client-stat"
    clientStat.EntityData.BundleName = "cisco_ios_xr"
    clientStat.EntityData.ParentYangName = "client-stats"
    clientStat.EntityData.SegmentPath = "client-stat" + "[client-ifhandle='" + fmt.Sprintf("%v", clientStat.ClientIfhandle) + "']"
    clientStat.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clientStat.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clientStat.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clientStat.EntityData.Children = make(map[string]types.YChild)
    clientStat.EntityData.Leafs = make(map[string]types.YLeaf)
    clientStat.EntityData.Leafs["client-ifhandle"] = types.YLeaf{"ClientIfhandle", clientStat.ClientIfhandle}
    clientStat.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", clientStat.InterfaceName}
    clientStat.EntityData.Leafs["num-events-received"] = types.YLeaf{"NumEventsReceived", clientStat.NumEventsReceived}
    clientStat.EntityData.Leafs["num-create-event-received"] = types.YLeaf{"NumCreateEventReceived", clientStat.NumCreateEventReceived}
    clientStat.EntityData.Leafs["num-delete-event-received"] = types.YLeaf{"NumDeleteEventReceived", clientStat.NumDeleteEventReceived}
    clientStat.EntityData.Leafs["num-reboot-event-received"] = types.YLeaf{"NumRebootEventReceived", clientStat.NumRebootEventReceived}
    clientStat.EntityData.Leafs["num-reinit-event-received"] = types.YLeaf{"NumReinitEventReceived", clientStat.NumReinitEventReceived}
    clientStat.EntityData.Leafs["num-packet-event-received"] = types.YLeaf{"NumPacketEventReceived", clientStat.NumPacketEventReceived}
    clientStat.EntityData.Leafs["num-init-timer-eventi"] = types.YLeaf{"NumInitTimerEventi", clientStat.NumInitTimerEventi}
    clientStat.EntityData.Leafs["num-t1-timer-event"] = types.YLeaf{"NumT1TimerEvent", clientStat.NumT1TimerEvent}
    clientStat.EntityData.Leafs["num-t2-timer-event"] = types.YLeaf{"NumT2TimerEvent", clientStat.NumT2TimerEvent}
    clientStat.EntityData.Leafs["num-lease-timer-event"] = types.YLeaf{"NumLeaseTimerEvent", clientStat.NumLeaseTimerEvent}
    clientStat.EntityData.Leafs["num-vbind-timer-event"] = types.YLeaf{"NumVbindTimerEvent", clientStat.NumVbindTimerEvent}
    clientStat.EntityData.Leafs["num-discovers-sent-successfully"] = types.YLeaf{"NumDiscoversSentSuccessfully", clientStat.NumDiscoversSentSuccessfully}
    clientStat.EntityData.Leafs["num-requests-sent-successfully"] = types.YLeaf{"NumRequestsSentSuccessfully", clientStat.NumRequestsSentSuccessfully}
    clientStat.EntityData.Leafs["num-releases-sent-successfully"] = types.YLeaf{"NumReleasesSentSuccessfully", clientStat.NumReleasesSentSuccessfully}
    clientStat.EntityData.Leafs["num-renews-sent-successfully"] = types.YLeaf{"NumRenewsSentSuccessfully", clientStat.NumRenewsSentSuccessfully}
    clientStat.EntityData.Leafs["num-rebinds-sent-successfully"] = types.YLeaf{"NumRebindsSentSuccessfully", clientStat.NumRebindsSentSuccessfully}
    clientStat.EntityData.Leafs["num-declines-sent-successfully"] = types.YLeaf{"NumDeclinesSentSuccessfully", clientStat.NumDeclinesSentSuccessfully}
    clientStat.EntityData.Leafs["num-request-after-reboot-sent"] = types.YLeaf{"NumRequestAfterRebootSent", clientStat.NumRequestAfterRebootSent}
    clientStat.EntityData.Leafs["num-valid-offers-received"] = types.YLeaf{"NumValidOffersReceived", clientStat.NumValidOffersReceived}
    clientStat.EntityData.Leafs["num-valid-acks-received"] = types.YLeaf{"NumValidAcksReceived", clientStat.NumValidAcksReceived}
    clientStat.EntityData.Leafs["num-valid-nacks-received"] = types.YLeaf{"NumValidNacksReceived", clientStat.NumValidNacksReceived}
    clientStat.EntityData.Leafs["num-unicast-packet-sent-successfully"] = types.YLeaf{"NumUnicastPacketSentSuccessfully", clientStat.NumUnicastPacketSentSuccessfully}
    clientStat.EntityData.Leafs["num-broadcast-packet-sent-success"] = types.YLeaf{"NumBroadcastPacketSentSuccess", clientStat.NumBroadcastPacketSentSuccess}
    clientStat.EntityData.Leafs["num-init-timer-start"] = types.YLeaf{"NumInitTimerStart", clientStat.NumInitTimerStart}
    clientStat.EntityData.Leafs["num-init-timer-stop"] = types.YLeaf{"NumInitTimerStop", clientStat.NumInitTimerStop}
    clientStat.EntityData.Leafs["num-t1-timer-start"] = types.YLeaf{"NumT1TimerStart", clientStat.NumT1TimerStart}
    clientStat.EntityData.Leafs["num-t1-timer-stop"] = types.YLeaf{"NumT1TimerStop", clientStat.NumT1TimerStop}
    clientStat.EntityData.Leafs["num-t2-timer-start"] = types.YLeaf{"NumT2TimerStart", clientStat.NumT2TimerStart}
    clientStat.EntityData.Leafs["num-t2-timer-stop"] = types.YLeaf{"NumT2TimerStop", clientStat.NumT2TimerStop}
    clientStat.EntityData.Leafs["num-lease-timer-start"] = types.YLeaf{"NumLeaseTimerStart", clientStat.NumLeaseTimerStart}
    clientStat.EntityData.Leafs["num-lease-timer-stop"] = types.YLeaf{"NumLeaseTimerStop", clientStat.NumLeaseTimerStop}
    clientStat.EntityData.Leafs["num-vbind-timer-start"] = types.YLeaf{"NumVbindTimerStart", clientStat.NumVbindTimerStart}
    clientStat.EntityData.Leafs["num-vbind-timer-stop"] = types.YLeaf{"NumVbindTimerStop", clientStat.NumVbindTimerStop}
    clientStat.EntityData.Leafs["num-invalid-events"] = types.YLeaf{"NumInvalidEvents", clientStat.NumInvalidEvents}
    clientStat.EntityData.Leafs["num-discovers-failed"] = types.YLeaf{"NumDiscoversFailed", clientStat.NumDiscoversFailed}
    clientStat.EntityData.Leafs["num-requests-failed"] = types.YLeaf{"NumRequestsFailed", clientStat.NumRequestsFailed}
    clientStat.EntityData.Leafs["num-releases-failed"] = types.YLeaf{"NumReleasesFailed", clientStat.NumReleasesFailed}
    clientStat.EntityData.Leafs["num-renews-failed"] = types.YLeaf{"NumRenewsFailed", clientStat.NumRenewsFailed}
    clientStat.EntityData.Leafs["num-rebinds-failed"] = types.YLeaf{"NumRebindsFailed", clientStat.NumRebindsFailed}
    clientStat.EntityData.Leafs["num-declines-failed"] = types.YLeaf{"NumDeclinesFailed", clientStat.NumDeclinesFailed}
    clientStat.EntityData.Leafs["num-request-after-reboot-failed"] = types.YLeaf{"NumRequestAfterRebootFailed", clientStat.NumRequestAfterRebootFailed}
    clientStat.EntityData.Leafs["num-invalid-offers"] = types.YLeaf{"NumInvalidOffers", clientStat.NumInvalidOffers}
    clientStat.EntityData.Leafs["num-invalid-acks"] = types.YLeaf{"NumInvalidAcks", clientStat.NumInvalidAcks}
    clientStat.EntityData.Leafs["num-invalid-nacks"] = types.YLeaf{"NumInvalidNacks", clientStat.NumInvalidNacks}
    clientStat.EntityData.Leafs["num-invalid-packets"] = types.YLeaf{"NumInvalidPackets", clientStat.NumInvalidPackets}
    clientStat.EntityData.Leafs["num-unicast-failed"] = types.YLeaf{"NumUnicastFailed", clientStat.NumUnicastFailed}
    clientStat.EntityData.Leafs["num-broadcast-failed"] = types.YLeaf{"NumBroadcastFailed", clientStat.NumBroadcastFailed}
    clientStat.EntityData.Leafs["num-xid-mismatch"] = types.YLeaf{"NumXidMismatch", clientStat.NumXidMismatch}
    return &(clientStat.EntityData)
}

// DhcpClient_Nodes_Node_Clients
// IPv4 DHCP client table
type DhcpClient_Nodes_Node_Clients struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Single DHCP client binding. The type is slice of
    // DhcpClient_Nodes_Node_Clients_Client.
    Client []DhcpClient_Nodes_Node_Clients_Client
}

func (clients *DhcpClient_Nodes_Node_Clients) GetEntityData() *types.CommonEntityData {
    clients.EntityData.YFilter = clients.YFilter
    clients.EntityData.YangName = "clients"
    clients.EntityData.BundleName = "cisco_ios_xr"
    clients.EntityData.ParentYangName = "node"
    clients.EntityData.SegmentPath = "clients"
    clients.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clients.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clients.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clients.EntityData.Children = make(map[string]types.YChild)
    clients.EntityData.Children["client"] = types.YChild{"Client", nil}
    for i := range clients.Client {
        clients.EntityData.Children[types.GetSegmentPath(&clients.Client[i])] = types.YChild{"Client", &clients.Client[i]}
    }
    clients.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(clients.EntityData)
}

// DhcpClient_Nodes_Node_Clients_Client
// Single DHCP client binding
type DhcpClient_Nodes_Node_Clients_Client struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Client Ifhandle. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4Address interface{}

    // Dhcp Client IP Address mask. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4SubnetMask interface{}

    // Dhcp Client selected server IP Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4ServerAddress interface{}

    // Dhcp Client next hop IP Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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

func (client *DhcpClient_Nodes_Node_Clients_Client) GetEntityData() *types.CommonEntityData {
    client.EntityData.YFilter = client.YFilter
    client.EntityData.YangName = "client"
    client.EntityData.BundleName = "cisco_ios_xr"
    client.EntityData.ParentYangName = "clients"
    client.EntityData.SegmentPath = "client" + "[client-ifhandle='" + fmt.Sprintf("%v", client.ClientIfhandle) + "']"
    client.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    client.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    client.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    client.EntityData.Children = make(map[string]types.YChild)
    client.EntityData.Leafs = make(map[string]types.YLeaf)
    client.EntityData.Leafs["client-ifhandle"] = types.YLeaf{"ClientIfhandle", client.ClientIfhandle}
    client.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", client.InterfaceName}
    client.EntityData.Leafs["client-mac-address"] = types.YLeaf{"ClientMacAddress", client.ClientMacAddress}
    client.EntityData.Leafs["client-id"] = types.YLeaf{"ClientId", client.ClientId}
    client.EntityData.Leafs["ipv4-client-state"] = types.YLeaf{"Ipv4ClientState", client.Ipv4ClientState}
    client.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", client.Ipv4Address}
    client.EntityData.Leafs["ipv4-subnet-mask"] = types.YLeaf{"Ipv4SubnetMask", client.Ipv4SubnetMask}
    client.EntityData.Leafs["ipv4-server-address"] = types.YLeaf{"Ipv4ServerAddress", client.Ipv4ServerAddress}
    client.EntityData.Leafs["next-hop-ipv4-address"] = types.YLeaf{"NextHopIpv4Address", client.NextHopIpv4Address}
    client.EntityData.Leafs["ipv4-lease-time"] = types.YLeaf{"Ipv4LeaseTime", client.Ipv4LeaseTime}
    client.EntityData.Leafs["ipv4-renew-time"] = types.YLeaf{"Ipv4RenewTime", client.Ipv4RenewTime}
    client.EntityData.Leafs["ipv4-rebind-time"] = types.YLeaf{"Ipv4RebindTime", client.Ipv4RebindTime}
    client.EntityData.Leafs["ipv4-address-configured"] = types.YLeaf{"Ipv4AddressConfigured", client.Ipv4AddressConfigured}
    return &(client.EntityData)
}

// Ipv4Dhcpd
// ipv4 dhcpd
type Ipv4Dhcpd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DHCP Snoop operational data.
    Snoop Ipv4Dhcpd_Snoop

    // IPv4 DHCPD operational data for a particular location.
    Nodes Ipv4Dhcpd_Nodes
}

func (ipv4Dhcpd *Ipv4Dhcpd) GetEntityData() *types.CommonEntityData {
    ipv4Dhcpd.EntityData.YFilter = ipv4Dhcpd.YFilter
    ipv4Dhcpd.EntityData.YangName = "ipv4-dhcpd"
    ipv4Dhcpd.EntityData.BundleName = "cisco_ios_xr"
    ipv4Dhcpd.EntityData.ParentYangName = "Cisco-IOS-XR-ipv4-dhcpd-oper"
    ipv4Dhcpd.EntityData.SegmentPath = "Cisco-IOS-XR-ipv4-dhcpd-oper:ipv4-dhcpd"
    ipv4Dhcpd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4Dhcpd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4Dhcpd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4Dhcpd.EntityData.Children = make(map[string]types.YChild)
    ipv4Dhcpd.EntityData.Children["snoop"] = types.YChild{"Snoop", &ipv4Dhcpd.Snoop}
    ipv4Dhcpd.EntityData.Children["nodes"] = types.YChild{"Nodes", &ipv4Dhcpd.Nodes}
    ipv4Dhcpd.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipv4Dhcpd.EntityData)
}

// Ipv4Dhcpd_Snoop
// DHCP Snoop operational data
type Ipv4Dhcpd_Snoop struct {
    EntityData types.CommonEntityData
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

func (snoop *Ipv4Dhcpd_Snoop) GetEntityData() *types.CommonEntityData {
    snoop.EntityData.YFilter = snoop.YFilter
    snoop.EntityData.YangName = "snoop"
    snoop.EntityData.BundleName = "cisco_ios_xr"
    snoop.EntityData.ParentYangName = "ipv4-dhcpd"
    snoop.EntityData.SegmentPath = "snoop"
    snoop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    snoop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    snoop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    snoop.EntityData.Children = make(map[string]types.YChild)
    snoop.EntityData.Children["bindings"] = types.YChild{"Bindings", &snoop.Bindings}
    snoop.EntityData.Children["binding-statistics"] = types.YChild{"BindingStatistics", &snoop.BindingStatistics}
    snoop.EntityData.Children["statistics-info"] = types.YChild{"StatisticsInfo", &snoop.StatisticsInfo}
    snoop.EntityData.Children["profiles"] = types.YChild{"Profiles", &snoop.Profiles}
    snoop.EntityData.Children["statistics"] = types.YChild{"Statistics", &snoop.Statistics}
    snoop.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(snoop.EntityData)
}

// Ipv4Dhcpd_Snoop_Bindings
// DHCP Snoop Bindings
type Ipv4Dhcpd_Snoop_Bindings struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DHCP Snoop binding. The type is slice of Ipv4Dhcpd_Snoop_Bindings_Binding.
    Binding []Ipv4Dhcpd_Snoop_Bindings_Binding
}

func (bindings *Ipv4Dhcpd_Snoop_Bindings) GetEntityData() *types.CommonEntityData {
    bindings.EntityData.YFilter = bindings.YFilter
    bindings.EntityData.YangName = "bindings"
    bindings.EntityData.BundleName = "cisco_ios_xr"
    bindings.EntityData.ParentYangName = "snoop"
    bindings.EntityData.SegmentPath = "bindings"
    bindings.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bindings.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bindings.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bindings.EntityData.Children = make(map[string]types.YChild)
    bindings.EntityData.Children["binding"] = types.YChild{"Binding", nil}
    for i := range bindings.Binding {
        bindings.EntityData.Children[types.GetSegmentPath(&bindings.Binding[i])] = types.YChild{"Binding", &bindings.Binding[i]}
    }
    bindings.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bindings.EntityData)
}

// Ipv4Dhcpd_Snoop_Bindings_Binding
// DHCP Snoop binding
type Ipv4Dhcpd_Snoop_Bindings_Binding struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Client opaque handle. The type is string with
    // pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    ClientUid interface{}

    // DHCP client MAC address. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    SnoopBindingChAddr interface{}

    // DHCP client MAC address length. The type is interface{} with range: 0..255.
    SnoopBindingChAddrLen interface{}

    // DHCP iaddr. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    SnoopBindingIAddr interface{}

    // DHCP client id. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
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

func (binding *Ipv4Dhcpd_Snoop_Bindings_Binding) GetEntityData() *types.CommonEntityData {
    binding.EntityData.YFilter = binding.YFilter
    binding.EntityData.YangName = "binding"
    binding.EntityData.BundleName = "cisco_ios_xr"
    binding.EntityData.ParentYangName = "bindings"
    binding.EntityData.SegmentPath = "binding" + "[client-uid='" + fmt.Sprintf("%v", binding.ClientUid) + "']"
    binding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    binding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    binding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    binding.EntityData.Children = make(map[string]types.YChild)
    binding.EntityData.Leafs = make(map[string]types.YLeaf)
    binding.EntityData.Leafs["client-uid"] = types.YLeaf{"ClientUid", binding.ClientUid}
    binding.EntityData.Leafs["snoop-binding-ch-addr"] = types.YLeaf{"SnoopBindingChAddr", binding.SnoopBindingChAddr}
    binding.EntityData.Leafs["snoop-binding-ch-addr-len"] = types.YLeaf{"SnoopBindingChAddrLen", binding.SnoopBindingChAddrLen}
    binding.EntityData.Leafs["snoop-binding-i-addr"] = types.YLeaf{"SnoopBindingIAddr", binding.SnoopBindingIAddr}
    binding.EntityData.Leafs["snoop-binding-client-id"] = types.YLeaf{"SnoopBindingClientId", binding.SnoopBindingClientId}
    binding.EntityData.Leafs["snoop-binding-client-id-len"] = types.YLeaf{"SnoopBindingClientIdLen", binding.SnoopBindingClientIdLen}
    binding.EntityData.Leafs["snoop-binding-state"] = types.YLeaf{"SnoopBindingState", binding.SnoopBindingState}
    binding.EntityData.Leafs["snoop-binding-lease"] = types.YLeaf{"SnoopBindingLease", binding.SnoopBindingLease}
    binding.EntityData.Leafs["snoop-binding-lease-start-time"] = types.YLeaf{"SnoopBindingLeaseStartTime", binding.SnoopBindingLeaseStartTime}
    binding.EntityData.Leafs["snoop-binding-profile-name"] = types.YLeaf{"SnoopBindingProfileName", binding.SnoopBindingProfileName}
    binding.EntityData.Leafs["snoop-bindng-interface-name"] = types.YLeaf{"SnoopBindngInterfaceName", binding.SnoopBindngInterfaceName}
    binding.EntityData.Leafs["snoop-binding-bridge-name"] = types.YLeaf{"SnoopBindingBridgeName", binding.SnoopBindingBridgeName}
    return &(binding.EntityData)
}

// Ipv4Dhcpd_Snoop_BindingStatistics
// DHCP snoop binding statistics
type Ipv4Dhcpd_Snoop_BindingStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Total number of snoop bindings. The type is interface{} with range:
    // 0..4294967295.
    SnoopBindingTotal interface{}

    // Snoop binding timestamp. The type is interface{} with range: 0..4294967295.
    SnoopBindingTimestamp interface{}
}

func (bindingStatistics *Ipv4Dhcpd_Snoop_BindingStatistics) GetEntityData() *types.CommonEntityData {
    bindingStatistics.EntityData.YFilter = bindingStatistics.YFilter
    bindingStatistics.EntityData.YangName = "binding-statistics"
    bindingStatistics.EntityData.BundleName = "cisco_ios_xr"
    bindingStatistics.EntityData.ParentYangName = "snoop"
    bindingStatistics.EntityData.SegmentPath = "binding-statistics"
    bindingStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bindingStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bindingStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bindingStatistics.EntityData.Children = make(map[string]types.YChild)
    bindingStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    bindingStatistics.EntityData.Leafs["snoop-binding-total"] = types.YLeaf{"SnoopBindingTotal", bindingStatistics.SnoopBindingTotal}
    bindingStatistics.EntityData.Leafs["snoop-binding-timestamp"] = types.YLeaf{"SnoopBindingTimestamp", bindingStatistics.SnoopBindingTimestamp}
    return &(bindingStatistics.EntityData)
}

// Ipv4Dhcpd_Snoop_StatisticsInfo
// DHCP snoop statistics info
type Ipv4Dhcpd_Snoop_StatisticsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Snoop Stats timestamp. The type is interface{} with range: 0..4294967295.
    SnoopStatsTimestamp interface{}
}

func (statisticsInfo *Ipv4Dhcpd_Snoop_StatisticsInfo) GetEntityData() *types.CommonEntityData {
    statisticsInfo.EntityData.YFilter = statisticsInfo.YFilter
    statisticsInfo.EntityData.YangName = "statistics-info"
    statisticsInfo.EntityData.BundleName = "cisco_ios_xr"
    statisticsInfo.EntityData.ParentYangName = "snoop"
    statisticsInfo.EntityData.SegmentPath = "statistics-info"
    statisticsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statisticsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statisticsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statisticsInfo.EntityData.Children = make(map[string]types.YChild)
    statisticsInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    statisticsInfo.EntityData.Leafs["snoop-stats-timestamp"] = types.YLeaf{"SnoopStatsTimestamp", statisticsInfo.SnoopStatsTimestamp}
    return &(statisticsInfo.EntityData)
}

// Ipv4Dhcpd_Snoop_Profiles
// DHCP Snoop Profile
type Ipv4Dhcpd_Snoop_Profiles struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DHCP Snoop profile. The type is slice of Ipv4Dhcpd_Snoop_Profiles_Profile.
    Profile []Ipv4Dhcpd_Snoop_Profiles_Profile
}

func (profiles *Ipv4Dhcpd_Snoop_Profiles) GetEntityData() *types.CommonEntityData {
    profiles.EntityData.YFilter = profiles.YFilter
    profiles.EntityData.YangName = "profiles"
    profiles.EntityData.BundleName = "cisco_ios_xr"
    profiles.EntityData.ParentYangName = "snoop"
    profiles.EntityData.SegmentPath = "profiles"
    profiles.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    profiles.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    profiles.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    profiles.EntityData.Children = make(map[string]types.YChild)
    profiles.EntityData.Children["profile"] = types.YChild{"Profile", nil}
    for i := range profiles.Profile {
        profiles.EntityData.Children[types.GetSegmentPath(&profiles.Profile[i])] = types.YChild{"Profile", &profiles.Profile[i]}
    }
    profiles.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(profiles.EntityData)
}

// Ipv4Dhcpd_Snoop_Profiles_Profile
// DHCP Snoop profile
type Ipv4Dhcpd_Snoop_Profiles_Profile struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Profile name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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

func (profile *Ipv4Dhcpd_Snoop_Profiles_Profile) GetEntityData() *types.CommonEntityData {
    profile.EntityData.YFilter = profile.YFilter
    profile.EntityData.YangName = "profile"
    profile.EntityData.BundleName = "cisco_ios_xr"
    profile.EntityData.ParentYangName = "profiles"
    profile.EntityData.SegmentPath = "profile" + "[profile-name='" + fmt.Sprintf("%v", profile.ProfileName) + "']"
    profile.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    profile.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    profile.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    profile.EntityData.Children = make(map[string]types.YChild)
    profile.EntityData.Leafs = make(map[string]types.YLeaf)
    profile.EntityData.Leafs["profile-name"] = types.YLeaf{"ProfileName", profile.ProfileName}
    profile.EntityData.Leafs["snoop-profile-name"] = types.YLeaf{"SnoopProfileName", profile.SnoopProfileName}
    profile.EntityData.Leafs["snoop-profile-uid"] = types.YLeaf{"SnoopProfileUid", profile.SnoopProfileUid}
    profile.EntityData.Leafs["snoop-profile-relay-info-option"] = types.YLeaf{"SnoopProfileRelayInfoOption", profile.SnoopProfileRelayInfoOption}
    profile.EntityData.Leafs["snoop-profile-relay-info-allow-untrusted"] = types.YLeaf{"SnoopProfileRelayInfoAllowUntrusted", profile.SnoopProfileRelayInfoAllowUntrusted}
    profile.EntityData.Leafs["snoop-profile-relay-info-policy"] = types.YLeaf{"SnoopProfileRelayInfoPolicy", profile.SnoopProfileRelayInfoPolicy}
    profile.EntityData.Leafs["snoop-profile-trusted"] = types.YLeaf{"SnoopProfileTrusted", profile.SnoopProfileTrusted}
    return &(profile.EntityData)
}

// Ipv4Dhcpd_Snoop_Statistics
// DHCP Snoop Statistics
type Ipv4Dhcpd_Snoop_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DHCP Snoop bridge domain statistics. The type is slice of
    // Ipv4Dhcpd_Snoop_Statistics_Statistic.
    Statistic []Ipv4Dhcpd_Snoop_Statistics_Statistic
}

func (statistics *Ipv4Dhcpd_Snoop_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "snoop"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = make(map[string]types.YChild)
    statistics.EntityData.Children["statistic"] = types.YChild{"Statistic", nil}
    for i := range statistics.Statistic {
        statistics.EntityData.Children[types.GetSegmentPath(&statistics.Statistic[i])] = types.YChild{"Statistic", &statistics.Statistic[i]}
    }
    statistics.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(statistics.EntityData)
}

// Ipv4Dhcpd_Snoop_Statistics_Statistic
// DHCP Snoop bridge domain statistics
type Ipv4Dhcpd_Snoop_Statistics_Statistic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Bridge domain name. The type is string with
    // pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    BridgeName interface{}

    // DHCP L2 bridge name. The type is string with length: 0..74.
    SnoopStatisticsBridgeName interface{}

    // Public snoop statistics. The type is slice of interface{} with range:
    // 0..18446744073709551615.
    SnoopStatistic []interface{}
}

func (statistic *Ipv4Dhcpd_Snoop_Statistics_Statistic) GetEntityData() *types.CommonEntityData {
    statistic.EntityData.YFilter = statistic.YFilter
    statistic.EntityData.YangName = "statistic"
    statistic.EntityData.BundleName = "cisco_ios_xr"
    statistic.EntityData.ParentYangName = "statistics"
    statistic.EntityData.SegmentPath = "statistic" + "[bridge-name='" + fmt.Sprintf("%v", statistic.BridgeName) + "']"
    statistic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistic.EntityData.Children = make(map[string]types.YChild)
    statistic.EntityData.Leafs = make(map[string]types.YLeaf)
    statistic.EntityData.Leafs["bridge-name"] = types.YLeaf{"BridgeName", statistic.BridgeName}
    statistic.EntityData.Leafs["snoop-statistics-bridge-name"] = types.YLeaf{"SnoopStatisticsBridgeName", statistic.SnoopStatisticsBridgeName}
    statistic.EntityData.Leafs["snoop-statistic"] = types.YLeaf{"SnoopStatistic", statistic.SnoopStatistic}
    return &(statistic.EntityData)
}

// Ipv4Dhcpd_Nodes
// IPv4 DHCPD operational data for a particular
// location
type Ipv4Dhcpd_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Location. For eg., 0/1/CPU0. The type is slice of Ipv4Dhcpd_Nodes_Node.
    Node []Ipv4Dhcpd_Nodes_Node
}

func (nodes *Ipv4Dhcpd_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "ipv4-dhcpd"
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

// Ipv4Dhcpd_Nodes_Node
// Location. For eg., 0/1/CPU0
type Ipv4Dhcpd_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The node id to filter on. For eg., 0/1/CPU0. The
    // type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
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

func (node *Ipv4Dhcpd_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + "[nodeid='" + fmt.Sprintf("%v", node.Nodeid) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Children["proxy"] = types.YChild{"Proxy", &node.Proxy}
    node.EntityData.Children["interfaces"] = types.YChild{"Interfaces", &node.Interfaces}
    node.EntityData.Children["base"] = types.YChild{"Base", &node.Base}
    node.EntityData.Children["server"] = types.YChild{"Server", &node.Server}
    node.EntityData.Children["relay"] = types.YChild{"Relay", &node.Relay}
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["nodeid"] = types.YLeaf{"Nodeid", node.Nodeid}
    return &(node.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Proxy
// IPv4 DHCP proxy operational data
type Ipv4Dhcpd_Nodes_Node_Proxy struct {
    EntityData types.CommonEntityData
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

func (proxy *Ipv4Dhcpd_Nodes_Node_Proxy) GetEntityData() *types.CommonEntityData {
    proxy.EntityData.YFilter = proxy.YFilter
    proxy.EntityData.YangName = "proxy"
    proxy.EntityData.BundleName = "cisco_ios_xr"
    proxy.EntityData.ParentYangName = "node"
    proxy.EntityData.SegmentPath = "proxy"
    proxy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    proxy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    proxy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    proxy.EntityData.Children = make(map[string]types.YChild)
    proxy.EntityData.Children["statistics-info"] = types.YChild{"StatisticsInfo", &proxy.StatisticsInfo}
    proxy.EntityData.Children["vrfs"] = types.YChild{"Vrfs", &proxy.Vrfs}
    proxy.EntityData.Children["profiles"] = types.YChild{"Profiles", &proxy.Profiles}
    proxy.EntityData.Children["statistics"] = types.YChild{"Statistics", &proxy.Statistics}
    proxy.EntityData.Children["binding"] = types.YChild{"Binding", &proxy.Binding}
    proxy.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(proxy.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Proxy_StatisticsInfo
// DHCP proxy stats info
type Ipv4Dhcpd_Nodes_Node_Proxy_StatisticsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Proxy Stats timestamp. The type is interface{} with range: 0..4294967295.
    ProxyStatsTimestamp interface{}
}

func (statisticsInfo *Ipv4Dhcpd_Nodes_Node_Proxy_StatisticsInfo) GetEntityData() *types.CommonEntityData {
    statisticsInfo.EntityData.YFilter = statisticsInfo.YFilter
    statisticsInfo.EntityData.YangName = "statistics-info"
    statisticsInfo.EntityData.BundleName = "cisco_ios_xr"
    statisticsInfo.EntityData.ParentYangName = "proxy"
    statisticsInfo.EntityData.SegmentPath = "statistics-info"
    statisticsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statisticsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statisticsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statisticsInfo.EntityData.Children = make(map[string]types.YChild)
    statisticsInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    statisticsInfo.EntityData.Leafs["proxy-stats-timestamp"] = types.YLeaf{"ProxyStatsTimestamp", statisticsInfo.ProxyStatsTimestamp}
    return &(statisticsInfo.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs
// DHCP proxy list of VRF names
type Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 DHCP proxy VRF name. The type is slice of
    // Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf.
    Vrf []Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf
}

func (vrfs *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs) GetEntityData() *types.CommonEntityData {
    vrfs.EntityData.YFilter = vrfs.YFilter
    vrfs.EntityData.YangName = "vrfs"
    vrfs.EntityData.BundleName = "cisco_ios_xr"
    vrfs.EntityData.ParentYangName = "proxy"
    vrfs.EntityData.SegmentPath = "vrfs"
    vrfs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfs.EntityData.Children = make(map[string]types.YChild)
    vrfs.EntityData.Children["vrf"] = types.YChild{"Vrf", nil}
    for i := range vrfs.Vrf {
        vrfs.EntityData.Children[types.GetSegmentPath(&vrfs.Vrf[i])] = types.YChild{"Vrf", &vrfs.Vrf[i]}
    }
    vrfs.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(vrfs.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf
// IPv4 DHCP proxy VRF name
type Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    VrfName interface{}

    // IPv4 DHCP proxy statistics.
    Statistics Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics
}

func (vrf *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "vrfs"
    vrf.EntityData.SegmentPath = "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = make(map[string]types.YChild)
    vrf.EntityData.Children["statistics"] = types.YChild{"Statistics", &vrf.Statistics}
    vrf.EntityData.Leafs = make(map[string]types.YLeaf)
    vrf.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", vrf.VrfName}
    return &(vrf.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics
// IPv4 DHCP proxy statistics
type Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics struct {
    EntityData types.CommonEntityData
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

func (statistics *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "vrf"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = make(map[string]types.YChild)
    statistics.EntityData.Children["discover"] = types.YChild{"Discover", &statistics.Discover}
    statistics.EntityData.Children["offer"] = types.YChild{"Offer", &statistics.Offer}
    statistics.EntityData.Children["request"] = types.YChild{"Request", &statistics.Request}
    statistics.EntityData.Children["decline"] = types.YChild{"Decline", &statistics.Decline}
    statistics.EntityData.Children["ack"] = types.YChild{"Ack", &statistics.Ack}
    statistics.EntityData.Children["nak"] = types.YChild{"Nak", &statistics.Nak}
    statistics.EntityData.Children["release"] = types.YChild{"Release", &statistics.Release}
    statistics.EntityData.Children["inform"] = types.YChild{"Inform", &statistics.Inform}
    statistics.EntityData.Children["lease-query"] = types.YChild{"LeaseQuery", &statistics.LeaseQuery}
    statistics.EntityData.Children["lease-not-assigned"] = types.YChild{"LeaseNotAssigned", &statistics.LeaseNotAssigned}
    statistics.EntityData.Children["lease-unknown"] = types.YChild{"LeaseUnknown", &statistics.LeaseUnknown}
    statistics.EntityData.Children["lease-active"] = types.YChild{"LeaseActive", &statistics.LeaseActive}
    statistics.EntityData.Children["bootp-request"] = types.YChild{"BootpRequest", &statistics.BootpRequest}
    statistics.EntityData.Children["bootp-reply"] = types.YChild{"BootpReply", &statistics.BootpReply}
    statistics.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(statistics.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Discover
// DHCP discover packets
type Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Discover struct {
    EntityData types.CommonEntityData
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

func (discover *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Discover) GetEntityData() *types.CommonEntityData {
    discover.EntityData.YFilter = discover.YFilter
    discover.EntityData.YangName = "discover"
    discover.EntityData.BundleName = "cisco_ios_xr"
    discover.EntityData.ParentYangName = "statistics"
    discover.EntityData.SegmentPath = "discover"
    discover.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    discover.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    discover.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    discover.EntityData.Children = make(map[string]types.YChild)
    discover.EntityData.Leafs = make(map[string]types.YLeaf)
    discover.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", discover.ReceivedPackets}
    discover.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", discover.TransmittedPackets}
    discover.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", discover.DroppedPackets}
    return &(discover.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Offer
// DHCP offer packets
type Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Offer struct {
    EntityData types.CommonEntityData
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

func (offer *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Offer) GetEntityData() *types.CommonEntityData {
    offer.EntityData.YFilter = offer.YFilter
    offer.EntityData.YangName = "offer"
    offer.EntityData.BundleName = "cisco_ios_xr"
    offer.EntityData.ParentYangName = "statistics"
    offer.EntityData.SegmentPath = "offer"
    offer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    offer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    offer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    offer.EntityData.Children = make(map[string]types.YChild)
    offer.EntityData.Leafs = make(map[string]types.YLeaf)
    offer.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", offer.ReceivedPackets}
    offer.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", offer.TransmittedPackets}
    offer.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", offer.DroppedPackets}
    return &(offer.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Request
// DHCP request packets
type Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Request struct {
    EntityData types.CommonEntityData
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

func (request *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Request) GetEntityData() *types.CommonEntityData {
    request.EntityData.YFilter = request.YFilter
    request.EntityData.YangName = "request"
    request.EntityData.BundleName = "cisco_ios_xr"
    request.EntityData.ParentYangName = "statistics"
    request.EntityData.SegmentPath = "request"
    request.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    request.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    request.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    request.EntityData.Children = make(map[string]types.YChild)
    request.EntityData.Leafs = make(map[string]types.YLeaf)
    request.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", request.ReceivedPackets}
    request.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", request.TransmittedPackets}
    request.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", request.DroppedPackets}
    return &(request.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Decline
// DHCP decline packets
type Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Decline struct {
    EntityData types.CommonEntityData
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

func (decline *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Decline) GetEntityData() *types.CommonEntityData {
    decline.EntityData.YFilter = decline.YFilter
    decline.EntityData.YangName = "decline"
    decline.EntityData.BundleName = "cisco_ios_xr"
    decline.EntityData.ParentYangName = "statistics"
    decline.EntityData.SegmentPath = "decline"
    decline.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    decline.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    decline.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    decline.EntityData.Children = make(map[string]types.YChild)
    decline.EntityData.Leafs = make(map[string]types.YLeaf)
    decline.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", decline.ReceivedPackets}
    decline.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", decline.TransmittedPackets}
    decline.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", decline.DroppedPackets}
    return &(decline.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Ack
// DHCP ack packets
type Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Ack struct {
    EntityData types.CommonEntityData
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

func (ack *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Ack) GetEntityData() *types.CommonEntityData {
    ack.EntityData.YFilter = ack.YFilter
    ack.EntityData.YangName = "ack"
    ack.EntityData.BundleName = "cisco_ios_xr"
    ack.EntityData.ParentYangName = "statistics"
    ack.EntityData.SegmentPath = "ack"
    ack.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ack.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ack.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ack.EntityData.Children = make(map[string]types.YChild)
    ack.EntityData.Leafs = make(map[string]types.YLeaf)
    ack.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", ack.ReceivedPackets}
    ack.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", ack.TransmittedPackets}
    ack.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", ack.DroppedPackets}
    return &(ack.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Nak
// DHCP nak packets
type Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Nak struct {
    EntityData types.CommonEntityData
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

func (nak *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Nak) GetEntityData() *types.CommonEntityData {
    nak.EntityData.YFilter = nak.YFilter
    nak.EntityData.YangName = "nak"
    nak.EntityData.BundleName = "cisco_ios_xr"
    nak.EntityData.ParentYangName = "statistics"
    nak.EntityData.SegmentPath = "nak"
    nak.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nak.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nak.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nak.EntityData.Children = make(map[string]types.YChild)
    nak.EntityData.Leafs = make(map[string]types.YLeaf)
    nak.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", nak.ReceivedPackets}
    nak.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", nak.TransmittedPackets}
    nak.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", nak.DroppedPackets}
    return &(nak.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Release
// DHCP release packets
type Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Release struct {
    EntityData types.CommonEntityData
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

func (release *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Release) GetEntityData() *types.CommonEntityData {
    release.EntityData.YFilter = release.YFilter
    release.EntityData.YangName = "release"
    release.EntityData.BundleName = "cisco_ios_xr"
    release.EntityData.ParentYangName = "statistics"
    release.EntityData.SegmentPath = "release"
    release.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    release.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    release.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    release.EntityData.Children = make(map[string]types.YChild)
    release.EntityData.Leafs = make(map[string]types.YLeaf)
    release.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", release.ReceivedPackets}
    release.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", release.TransmittedPackets}
    release.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", release.DroppedPackets}
    return &(release.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Inform
// DHCP inform packets
type Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Inform struct {
    EntityData types.CommonEntityData
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

func (inform *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Inform) GetEntityData() *types.CommonEntityData {
    inform.EntityData.YFilter = inform.YFilter
    inform.EntityData.YangName = "inform"
    inform.EntityData.BundleName = "cisco_ios_xr"
    inform.EntityData.ParentYangName = "statistics"
    inform.EntityData.SegmentPath = "inform"
    inform.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inform.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inform.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inform.EntityData.Children = make(map[string]types.YChild)
    inform.EntityData.Leafs = make(map[string]types.YLeaf)
    inform.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", inform.ReceivedPackets}
    inform.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", inform.TransmittedPackets}
    inform.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", inform.DroppedPackets}
    return &(inform.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQuery
// DHCP lease query packets
type Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQuery struct {
    EntityData types.CommonEntityData
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

func (leaseQuery *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQuery) GetEntityData() *types.CommonEntityData {
    leaseQuery.EntityData.YFilter = leaseQuery.YFilter
    leaseQuery.EntityData.YangName = "lease-query"
    leaseQuery.EntityData.BundleName = "cisco_ios_xr"
    leaseQuery.EntityData.ParentYangName = "statistics"
    leaseQuery.EntityData.SegmentPath = "lease-query"
    leaseQuery.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    leaseQuery.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    leaseQuery.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    leaseQuery.EntityData.Children = make(map[string]types.YChild)
    leaseQuery.EntityData.Leafs = make(map[string]types.YLeaf)
    leaseQuery.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", leaseQuery.ReceivedPackets}
    leaseQuery.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", leaseQuery.TransmittedPackets}
    leaseQuery.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", leaseQuery.DroppedPackets}
    return &(leaseQuery.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseNotAssigned
// DHCP lease not assigned
type Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseNotAssigned struct {
    EntityData types.CommonEntityData
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

func (leaseNotAssigned *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseNotAssigned) GetEntityData() *types.CommonEntityData {
    leaseNotAssigned.EntityData.YFilter = leaseNotAssigned.YFilter
    leaseNotAssigned.EntityData.YangName = "lease-not-assigned"
    leaseNotAssigned.EntityData.BundleName = "cisco_ios_xr"
    leaseNotAssigned.EntityData.ParentYangName = "statistics"
    leaseNotAssigned.EntityData.SegmentPath = "lease-not-assigned"
    leaseNotAssigned.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    leaseNotAssigned.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    leaseNotAssigned.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    leaseNotAssigned.EntityData.Children = make(map[string]types.YChild)
    leaseNotAssigned.EntityData.Leafs = make(map[string]types.YLeaf)
    leaseNotAssigned.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", leaseNotAssigned.ReceivedPackets}
    leaseNotAssigned.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", leaseNotAssigned.TransmittedPackets}
    leaseNotAssigned.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", leaseNotAssigned.DroppedPackets}
    return &(leaseNotAssigned.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseUnknown
// DHCP lease unknown
type Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseUnknown struct {
    EntityData types.CommonEntityData
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

func (leaseUnknown *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseUnknown) GetEntityData() *types.CommonEntityData {
    leaseUnknown.EntityData.YFilter = leaseUnknown.YFilter
    leaseUnknown.EntityData.YangName = "lease-unknown"
    leaseUnknown.EntityData.BundleName = "cisco_ios_xr"
    leaseUnknown.EntityData.ParentYangName = "statistics"
    leaseUnknown.EntityData.SegmentPath = "lease-unknown"
    leaseUnknown.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    leaseUnknown.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    leaseUnknown.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    leaseUnknown.EntityData.Children = make(map[string]types.YChild)
    leaseUnknown.EntityData.Leafs = make(map[string]types.YLeaf)
    leaseUnknown.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", leaseUnknown.ReceivedPackets}
    leaseUnknown.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", leaseUnknown.TransmittedPackets}
    leaseUnknown.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", leaseUnknown.DroppedPackets}
    return &(leaseUnknown.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseActive
// DHCP lease active
type Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseActive struct {
    EntityData types.CommonEntityData
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

func (leaseActive *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseActive) GetEntityData() *types.CommonEntityData {
    leaseActive.EntityData.YFilter = leaseActive.YFilter
    leaseActive.EntityData.YangName = "lease-active"
    leaseActive.EntityData.BundleName = "cisco_ios_xr"
    leaseActive.EntityData.ParentYangName = "statistics"
    leaseActive.EntityData.SegmentPath = "lease-active"
    leaseActive.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    leaseActive.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    leaseActive.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    leaseActive.EntityData.Children = make(map[string]types.YChild)
    leaseActive.EntityData.Leafs = make(map[string]types.YLeaf)
    leaseActive.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", leaseActive.ReceivedPackets}
    leaseActive.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", leaseActive.TransmittedPackets}
    leaseActive.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", leaseActive.DroppedPackets}
    return &(leaseActive.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_BootpRequest
// DHCP BOOTP request
type Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_BootpRequest struct {
    EntityData types.CommonEntityData
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

func (bootpRequest *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_BootpRequest) GetEntityData() *types.CommonEntityData {
    bootpRequest.EntityData.YFilter = bootpRequest.YFilter
    bootpRequest.EntityData.YangName = "bootp-request"
    bootpRequest.EntityData.BundleName = "cisco_ios_xr"
    bootpRequest.EntityData.ParentYangName = "statistics"
    bootpRequest.EntityData.SegmentPath = "bootp-request"
    bootpRequest.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bootpRequest.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bootpRequest.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bootpRequest.EntityData.Children = make(map[string]types.YChild)
    bootpRequest.EntityData.Leafs = make(map[string]types.YLeaf)
    bootpRequest.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", bootpRequest.ReceivedPackets}
    bootpRequest.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", bootpRequest.TransmittedPackets}
    bootpRequest.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", bootpRequest.DroppedPackets}
    return &(bootpRequest.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_BootpReply
// DHCP BOOTP reply
type Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_BootpReply struct {
    EntityData types.CommonEntityData
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

func (bootpReply *Ipv4Dhcpd_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_BootpReply) GetEntityData() *types.CommonEntityData {
    bootpReply.EntityData.YFilter = bootpReply.YFilter
    bootpReply.EntityData.YangName = "bootp-reply"
    bootpReply.EntityData.BundleName = "cisco_ios_xr"
    bootpReply.EntityData.ParentYangName = "statistics"
    bootpReply.EntityData.SegmentPath = "bootp-reply"
    bootpReply.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bootpReply.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bootpReply.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bootpReply.EntityData.Children = make(map[string]types.YChild)
    bootpReply.EntityData.Leafs = make(map[string]types.YLeaf)
    bootpReply.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", bootpReply.ReceivedPackets}
    bootpReply.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", bootpReply.TransmittedPackets}
    bootpReply.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", bootpReply.DroppedPackets}
    return &(bootpReply.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Proxy_Profiles
// IPv4 DHCP proxy profile
type Ipv4Dhcpd_Nodes_Node_Proxy_Profiles struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 DHCP proxy profile. The type is slice of
    // Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile.
    Profile []Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile
}

func (profiles *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles) GetEntityData() *types.CommonEntityData {
    profiles.EntityData.YFilter = profiles.YFilter
    profiles.EntityData.YangName = "profiles"
    profiles.EntityData.BundleName = "cisco_ios_xr"
    profiles.EntityData.ParentYangName = "proxy"
    profiles.EntityData.SegmentPath = "profiles"
    profiles.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    profiles.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    profiles.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    profiles.EntityData.Children = make(map[string]types.YChild)
    profiles.EntityData.Children["profile"] = types.YChild{"Profile", nil}
    for i := range profiles.Profile {
        profiles.EntityData.Children[types.GetSegmentPath(&profiles.Profile[i])] = types.YChild{"Profile", &profiles.Profile[i]}
    }
    profiles.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(profiles.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile
// IPv4 DHCP proxy profile
type Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Profile name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    ProfileHelperAddress []interface{}

    // VRF names. The type is slice of string with length: 0..33.
    VrfName []interface{}

    // Gateway addresses. The type is slice of string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    GiAddr []interface{}

    // VRF references.
    VrfReferences Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_VrfReferences

    // Interface references.
    InterfaceReferences Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_InterfaceReferences
}

func (profile *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile) GetEntityData() *types.CommonEntityData {
    profile.EntityData.YFilter = profile.YFilter
    profile.EntityData.YangName = "profile"
    profile.EntityData.BundleName = "cisco_ios_xr"
    profile.EntityData.ParentYangName = "profiles"
    profile.EntityData.SegmentPath = "profile" + "[profile-name='" + fmt.Sprintf("%v", profile.ProfileName) + "']"
    profile.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    profile.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    profile.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    profile.EntityData.Children = make(map[string]types.YChild)
    profile.EntityData.Children["vrf-references"] = types.YChild{"VrfReferences", &profile.VrfReferences}
    profile.EntityData.Children["interface-references"] = types.YChild{"InterfaceReferences", &profile.InterfaceReferences}
    profile.EntityData.Leafs = make(map[string]types.YLeaf)
    profile.EntityData.Leafs["profile-name"] = types.YLeaf{"ProfileName", profile.ProfileName}
    profile.EntityData.Leafs["is-relay-option-enabled"] = types.YLeaf{"IsRelayOptionEnabled", profile.IsRelayOptionEnabled}
    profile.EntityData.Leafs["relay-policy"] = types.YLeaf{"RelayPolicy", profile.RelayPolicy}
    profile.EntityData.Leafs["relay-authenticate"] = types.YLeaf{"RelayAuthenticate", profile.RelayAuthenticate}
    profile.EntityData.Leafs["is-relay-allow-untrusted-enabled"] = types.YLeaf{"IsRelayAllowUntrustedEnabled", profile.IsRelayAllowUntrustedEnabled}
    profile.EntityData.Leafs["is-relay-optionvpn-enabled"] = types.YLeaf{"IsRelayOptionvpnEnabled", profile.IsRelayOptionvpnEnabled}
    profile.EntityData.Leafs["relay-optionvpn-enabled-mode"] = types.YLeaf{"RelayOptionvpnEnabledMode", profile.RelayOptionvpnEnabledMode}
    profile.EntityData.Leafs["is-relay-check"] = types.YLeaf{"IsRelayCheck", profile.IsRelayCheck}
    profile.EntityData.Leafs["is-move-allowed"] = types.YLeaf{"IsMoveAllowed", profile.IsMoveAllowed}
    profile.EntityData.Leafs["proxy-broadcast-flag-policy"] = types.YLeaf{"ProxyBroadcastFlagPolicy", profile.ProxyBroadcastFlagPolicy}
    profile.EntityData.Leafs["proxy-profile-client-lease-time"] = types.YLeaf{"ProxyProfileClientLeaseTime", profile.ProxyProfileClientLeaseTime}
    profile.EntityData.Leafs["proxy-lease-limit-type"] = types.YLeaf{"ProxyLeaseLimitType", profile.ProxyLeaseLimitType}
    profile.EntityData.Leafs["proxy-lease-limit-count"] = types.YLeaf{"ProxyLeaseLimitCount", profile.ProxyLeaseLimitCount}
    profile.EntityData.Leafs["profile-helper-address"] = types.YLeaf{"ProfileHelperAddress", profile.ProfileHelperAddress}
    profile.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", profile.VrfName}
    profile.EntityData.Leafs["gi-addr"] = types.YLeaf{"GiAddr", profile.GiAddr}
    return &(profile.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_VrfReferences
// VRF references
type Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_VrfReferences struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ipv4 dhcpd proxy vrf reference. The type is slice of
    // Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_VrfReferences_Ipv4DhcpdProxyVrfReference.
    Ipv4DhcpdProxyVrfReference []Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_VrfReferences_Ipv4DhcpdProxyVrfReference
}

func (vrfReferences *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_VrfReferences) GetEntityData() *types.CommonEntityData {
    vrfReferences.EntityData.YFilter = vrfReferences.YFilter
    vrfReferences.EntityData.YangName = "vrf-references"
    vrfReferences.EntityData.BundleName = "cisco_ios_xr"
    vrfReferences.EntityData.ParentYangName = "profile"
    vrfReferences.EntityData.SegmentPath = "vrf-references"
    vrfReferences.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfReferences.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfReferences.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfReferences.EntityData.Children = make(map[string]types.YChild)
    vrfReferences.EntityData.Children["ipv4-dhcpd-proxy-vrf-reference"] = types.YChild{"Ipv4DhcpdProxyVrfReference", nil}
    for i := range vrfReferences.Ipv4DhcpdProxyVrfReference {
        vrfReferences.EntityData.Children[types.GetSegmentPath(&vrfReferences.Ipv4DhcpdProxyVrfReference[i])] = types.YChild{"Ipv4DhcpdProxyVrfReference", &vrfReferences.Ipv4DhcpdProxyVrfReference[i]}
    }
    vrfReferences.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(vrfReferences.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_VrfReferences_Ipv4DhcpdProxyVrfReference
// ipv4 dhcpd proxy vrf reference
type Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_VrfReferences_Ipv4DhcpdProxyVrfReference struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF name. The type is string with length: 0..33.
    ProxyReferenceVrfName interface{}
}

func (ipv4DhcpdProxyVrfReference *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_VrfReferences_Ipv4DhcpdProxyVrfReference) GetEntityData() *types.CommonEntityData {
    ipv4DhcpdProxyVrfReference.EntityData.YFilter = ipv4DhcpdProxyVrfReference.YFilter
    ipv4DhcpdProxyVrfReference.EntityData.YangName = "ipv4-dhcpd-proxy-vrf-reference"
    ipv4DhcpdProxyVrfReference.EntityData.BundleName = "cisco_ios_xr"
    ipv4DhcpdProxyVrfReference.EntityData.ParentYangName = "vrf-references"
    ipv4DhcpdProxyVrfReference.EntityData.SegmentPath = "ipv4-dhcpd-proxy-vrf-reference"
    ipv4DhcpdProxyVrfReference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4DhcpdProxyVrfReference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4DhcpdProxyVrfReference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4DhcpdProxyVrfReference.EntityData.Children = make(map[string]types.YChild)
    ipv4DhcpdProxyVrfReference.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv4DhcpdProxyVrfReference.EntityData.Leafs["proxy-reference-vrf-name"] = types.YLeaf{"ProxyReferenceVrfName", ipv4DhcpdProxyVrfReference.ProxyReferenceVrfName}
    return &(ipv4DhcpdProxyVrfReference.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_InterfaceReferences
// Interface references
type Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_InterfaceReferences struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ipv4 dhcpd proxy interface reference. The type is slice of
    // Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_InterfaceReferences_Ipv4DhcpdProxyInterfaceReference.
    Ipv4DhcpdProxyInterfaceReference []Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_InterfaceReferences_Ipv4DhcpdProxyInterfaceReference
}

func (interfaceReferences *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_InterfaceReferences) GetEntityData() *types.CommonEntityData {
    interfaceReferences.EntityData.YFilter = interfaceReferences.YFilter
    interfaceReferences.EntityData.YangName = "interface-references"
    interfaceReferences.EntityData.BundleName = "cisco_ios_xr"
    interfaceReferences.EntityData.ParentYangName = "profile"
    interfaceReferences.EntityData.SegmentPath = "interface-references"
    interfaceReferences.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceReferences.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceReferences.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceReferences.EntityData.Children = make(map[string]types.YChild)
    interfaceReferences.EntityData.Children["ipv4-dhcpd-proxy-interface-reference"] = types.YChild{"Ipv4DhcpdProxyInterfaceReference", nil}
    for i := range interfaceReferences.Ipv4DhcpdProxyInterfaceReference {
        interfaceReferences.EntityData.Children[types.GetSegmentPath(&interfaceReferences.Ipv4DhcpdProxyInterfaceReference[i])] = types.YChild{"Ipv4DhcpdProxyInterfaceReference", &interfaceReferences.Ipv4DhcpdProxyInterfaceReference[i]}
    }
    interfaceReferences.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaceReferences.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_InterfaceReferences_Ipv4DhcpdProxyInterfaceReference
// ipv4 dhcpd proxy interface reference
type Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_InterfaceReferences_Ipv4DhcpdProxyInterfaceReference struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface name. The type is string with length: 0..65.
    ProxyReferenceInterfaceName interface{}
}

func (ipv4DhcpdProxyInterfaceReference *Ipv4Dhcpd_Nodes_Node_Proxy_Profiles_Profile_InterfaceReferences_Ipv4DhcpdProxyInterfaceReference) GetEntityData() *types.CommonEntityData {
    ipv4DhcpdProxyInterfaceReference.EntityData.YFilter = ipv4DhcpdProxyInterfaceReference.YFilter
    ipv4DhcpdProxyInterfaceReference.EntityData.YangName = "ipv4-dhcpd-proxy-interface-reference"
    ipv4DhcpdProxyInterfaceReference.EntityData.BundleName = "cisco_ios_xr"
    ipv4DhcpdProxyInterfaceReference.EntityData.ParentYangName = "interface-references"
    ipv4DhcpdProxyInterfaceReference.EntityData.SegmentPath = "ipv4-dhcpd-proxy-interface-reference"
    ipv4DhcpdProxyInterfaceReference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4DhcpdProxyInterfaceReference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4DhcpdProxyInterfaceReference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4DhcpdProxyInterfaceReference.EntityData.Children = make(map[string]types.YChild)
    ipv4DhcpdProxyInterfaceReference.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv4DhcpdProxyInterfaceReference.EntityData.Leafs["proxy-reference-interface-name"] = types.YLeaf{"ProxyReferenceInterfaceName", ipv4DhcpdProxyInterfaceReference.ProxyReferenceInterfaceName}
    return &(ipv4DhcpdProxyInterfaceReference.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Proxy_Statistics
// DHCP proxy statistics
type Ipv4Dhcpd_Nodes_Node_Proxy_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ipv4 dhcpd proxy stat. The type is slice of
    // Ipv4Dhcpd_Nodes_Node_Proxy_Statistics_Ipv4DhcpdProxyStat.
    Ipv4DhcpdProxyStat []Ipv4Dhcpd_Nodes_Node_Proxy_Statistics_Ipv4DhcpdProxyStat
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Proxy_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "proxy"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = make(map[string]types.YChild)
    statistics.EntityData.Children["ipv4-dhcpd-proxy-stat"] = types.YChild{"Ipv4DhcpdProxyStat", nil}
    for i := range statistics.Ipv4DhcpdProxyStat {
        statistics.EntityData.Children[types.GetSegmentPath(&statistics.Ipv4DhcpdProxyStat[i])] = types.YChild{"Ipv4DhcpdProxyStat", &statistics.Ipv4DhcpdProxyStat[i]}
    }
    statistics.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(statistics.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Proxy_Statistics_Ipv4DhcpdProxyStat
// ipv4 dhcpd proxy stat
type Ipv4Dhcpd_Nodes_Node_Proxy_Statistics_Ipv4DhcpdProxyStat struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DHCP L3 VRF name. The type is string with length: 0..33.
    VrfName interface{}

    // Proxy statistics.
    Statistics Ipv4Dhcpd_Nodes_Node_Proxy_Statistics_Ipv4DhcpdProxyStat_Statistics_
}

func (ipv4DhcpdProxyStat *Ipv4Dhcpd_Nodes_Node_Proxy_Statistics_Ipv4DhcpdProxyStat) GetEntityData() *types.CommonEntityData {
    ipv4DhcpdProxyStat.EntityData.YFilter = ipv4DhcpdProxyStat.YFilter
    ipv4DhcpdProxyStat.EntityData.YangName = "ipv4-dhcpd-proxy-stat"
    ipv4DhcpdProxyStat.EntityData.BundleName = "cisco_ios_xr"
    ipv4DhcpdProxyStat.EntityData.ParentYangName = "statistics"
    ipv4DhcpdProxyStat.EntityData.SegmentPath = "ipv4-dhcpd-proxy-stat"
    ipv4DhcpdProxyStat.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4DhcpdProxyStat.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4DhcpdProxyStat.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4DhcpdProxyStat.EntityData.Children = make(map[string]types.YChild)
    ipv4DhcpdProxyStat.EntityData.Children["statistics"] = types.YChild{"Statistics", &ipv4DhcpdProxyStat.Statistics}
    ipv4DhcpdProxyStat.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv4DhcpdProxyStat.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", ipv4DhcpdProxyStat.VrfName}
    return &(ipv4DhcpdProxyStat.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Proxy_Statistics_Ipv4DhcpdProxyStat_Statistics_
// Proxy statistics
type Ipv4Dhcpd_Nodes_Node_Proxy_Statistics_Ipv4DhcpdProxyStat_Statistics_ struct {
    EntityData types.CommonEntityData
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

func (statistics_ *Ipv4Dhcpd_Nodes_Node_Proxy_Statistics_Ipv4DhcpdProxyStat_Statistics_) GetEntityData() *types.CommonEntityData {
    statistics_.EntityData.YFilter = statistics_.YFilter
    statistics_.EntityData.YangName = "statistics"
    statistics_.EntityData.BundleName = "cisco_ios_xr"
    statistics_.EntityData.ParentYangName = "ipv4-dhcpd-proxy-stat"
    statistics_.EntityData.SegmentPath = "statistics"
    statistics_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics_.EntityData.Children = make(map[string]types.YChild)
    statistics_.EntityData.Leafs = make(map[string]types.YLeaf)
    statistics_.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", statistics_.ReceivedPackets}
    statistics_.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", statistics_.TransmittedPackets}
    statistics_.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", statistics_.DroppedPackets}
    return &(statistics_.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Proxy_Binding
// DHCP proxy bindings
type Ipv4Dhcpd_Nodes_Node_Proxy_Binding struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DHCP proxy client bindings.
    Clients Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Clients

    // DHCP proxy binding summary.
    Summary Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Summary
}

func (binding *Ipv4Dhcpd_Nodes_Node_Proxy_Binding) GetEntityData() *types.CommonEntityData {
    binding.EntityData.YFilter = binding.YFilter
    binding.EntityData.YangName = "binding"
    binding.EntityData.BundleName = "cisco_ios_xr"
    binding.EntityData.ParentYangName = "proxy"
    binding.EntityData.SegmentPath = "binding"
    binding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    binding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    binding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    binding.EntityData.Children = make(map[string]types.YChild)
    binding.EntityData.Children["clients"] = types.YChild{"Clients", &binding.Clients}
    binding.EntityData.Children["summary"] = types.YChild{"Summary", &binding.Summary}
    binding.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(binding.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Clients
// DHCP proxy client bindings
type Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Clients struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Single DHCP proxy binding. The type is slice of
    // Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Clients_Client.
    Client []Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Clients_Client
}

func (clients *Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Clients) GetEntityData() *types.CommonEntityData {
    clients.EntityData.YFilter = clients.YFilter
    clients.EntityData.YangName = "clients"
    clients.EntityData.BundleName = "cisco_ios_xr"
    clients.EntityData.ParentYangName = "binding"
    clients.EntityData.SegmentPath = "clients"
    clients.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clients.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clients.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clients.EntityData.Children = make(map[string]types.YChild)
    clients.EntityData.Children["client"] = types.YChild{"Client", nil}
    for i := range clients.Client {
        clients.EntityData.Children[types.GetSegmentPath(&clients.Client[i])] = types.YChild{"Client", &clients.Client[i]}
    }
    clients.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(clients.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Clients_Client
// Single DHCP proxy binding
type Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Clients_Client struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Client ID. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    IpAddress interface{}

    // DHCP client GIADDR. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    ClientGiAddr interface{}

    // DHCP to server GIADDR. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    ToServerGiAddr interface{}

    // DHCP server IP address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    ServerIpAddress interface{}

    // DHCP reply server IP address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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
    // b'[a-zA-Z0-9./-]+'.
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
    // b'[a-zA-Z0-9./-]+'.
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

func (client *Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Clients_Client) GetEntityData() *types.CommonEntityData {
    client.EntityData.YFilter = client.YFilter
    client.EntityData.YangName = "client"
    client.EntityData.BundleName = "cisco_ios_xr"
    client.EntityData.ParentYangName = "clients"
    client.EntityData.SegmentPath = "client" + "[client-id='" + fmt.Sprintf("%v", client.ClientId) + "']"
    client.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    client.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    client.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    client.EntityData.Children = make(map[string]types.YChild)
    client.EntityData.Leafs = make(map[string]types.YLeaf)
    client.EntityData.Leafs["client-id"] = types.YLeaf{"ClientId", client.ClientId}
    client.EntityData.Leafs["client-id-xr"] = types.YLeaf{"ClientIdXr", client.ClientIdXr}
    client.EntityData.Leafs["mac-address"] = types.YLeaf{"MacAddress", client.MacAddress}
    client.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", client.VrfName}
    client.EntityData.Leafs["server-vrf-name"] = types.YLeaf{"ServerVrfName", client.ServerVrfName}
    client.EntityData.Leafs["ip-address"] = types.YLeaf{"IpAddress", client.IpAddress}
    client.EntityData.Leafs["client-gi-addr"] = types.YLeaf{"ClientGiAddr", client.ClientGiAddr}
    client.EntityData.Leafs["to-server-gi-addr"] = types.YLeaf{"ToServerGiAddr", client.ToServerGiAddr}
    client.EntityData.Leafs["server-ip-address"] = types.YLeaf{"ServerIpAddress", client.ServerIpAddress}
    client.EntityData.Leafs["reply-server-ip-address"] = types.YLeaf{"ReplyServerIpAddress", client.ReplyServerIpAddress}
    client.EntityData.Leafs["lease-time"] = types.YLeaf{"LeaseTime", client.LeaseTime}
    client.EntityData.Leafs["remaining-lease-time"] = types.YLeaf{"RemainingLeaseTime", client.RemainingLeaseTime}
    client.EntityData.Leafs["state"] = types.YLeaf{"State", client.State}
    client.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", client.InterfaceName}
    client.EntityData.Leafs["access-vrf-name"] = types.YLeaf{"AccessVrfName", client.AccessVrfName}
    client.EntityData.Leafs["proxy-binding-outer-tag"] = types.YLeaf{"ProxyBindingOuterTag", client.ProxyBindingOuterTag}
    client.EntityData.Leafs["proxy-binding-inner-tag"] = types.YLeaf{"ProxyBindingInnerTag", client.ProxyBindingInnerTag}
    client.EntityData.Leafs["profile-name"] = types.YLeaf{"ProfileName", client.ProfileName}
    client.EntityData.Leafs["is-nak-next-renew"] = types.YLeaf{"IsNakNextRenew", client.IsNakNextRenew}
    client.EntityData.Leafs["subscriber-label"] = types.YLeaf{"SubscriberLabel", client.SubscriberLabel}
    client.EntityData.Leafs["old-subscriber-label"] = types.YLeaf{"OldSubscriberLabel", client.OldSubscriberLabel}
    client.EntityData.Leafs["subscriber-interface-name"] = types.YLeaf{"SubscriberInterfaceName", client.SubscriberInterfaceName}
    client.EntityData.Leafs["rx-circuit-id"] = types.YLeaf{"RxCircuitId", client.RxCircuitId}
    client.EntityData.Leafs["tx-circuit-id"] = types.YLeaf{"TxCircuitId", client.TxCircuitId}
    client.EntityData.Leafs["rx-remote-id"] = types.YLeaf{"RxRemoteId", client.RxRemoteId}
    client.EntityData.Leafs["tx-remote-id"] = types.YLeaf{"TxRemoteId", client.TxRemoteId}
    client.EntityData.Leafs["rx-vsiso"] = types.YLeaf{"RxVsiso", client.RxVsiso}
    client.EntityData.Leafs["tx-vsiso"] = types.YLeaf{"TxVsiso", client.TxVsiso}
    client.EntityData.Leafs["is-auth-received"] = types.YLeaf{"IsAuthReceived", client.IsAuthReceived}
    client.EntityData.Leafs["is-mbl-subscriber"] = types.YLeaf{"IsMblSubscriber", client.IsMblSubscriber}
    client.EntityData.Leafs["param-request"] = types.YLeaf{"ParamRequest", client.ParamRequest}
    client.EntityData.Leafs["param-response"] = types.YLeaf{"ParamResponse", client.ParamResponse}
    client.EntityData.Leafs["session-start-time"] = types.YLeaf{"SessionStartTime", client.SessionStartTime}
    client.EntityData.Leafs["srg-state"] = types.YLeaf{"SrgState", client.SrgState}
    client.EntityData.Leafs["event-history"] = types.YLeaf{"EventHistory", client.EventHistory}
    return &(client.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Summary
// DHCP proxy binding summary
type Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Summary struct {
    EntityData types.CommonEntityData
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

func (summary *Ipv4Dhcpd_Nodes_Node_Proxy_Binding_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "binding"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = make(map[string]types.YChild)
    summary.EntityData.Leafs = make(map[string]types.YLeaf)
    summary.EntityData.Leafs["clients"] = types.YLeaf{"Clients", summary.Clients}
    summary.EntityData.Leafs["initializing-clients"] = types.YLeaf{"InitializingClients", summary.InitializingClients}
    summary.EntityData.Leafs["waiting-for-dpm-init"] = types.YLeaf{"WaitingForDpmInit", summary.WaitingForDpmInit}
    summary.EntityData.Leafs["waiting-for-dpm-request"] = types.YLeaf{"WaitingForDpmRequest", summary.WaitingForDpmRequest}
    summary.EntityData.Leafs["waiting-for-daps-init"] = types.YLeaf{"WaitingForDapsInit", summary.WaitingForDapsInit}
    summary.EntityData.Leafs["selecting-clients"] = types.YLeaf{"SelectingClients", summary.SelectingClients}
    summary.EntityData.Leafs["offer-sent-for-client"] = types.YLeaf{"OfferSentForClient", summary.OfferSentForClient}
    summary.EntityData.Leafs["requesting-clients"] = types.YLeaf{"RequestingClients", summary.RequestingClients}
    summary.EntityData.Leafs["request-waiting-for-dpm"] = types.YLeaf{"RequestWaitingForDpm", summary.RequestWaitingForDpm}
    summary.EntityData.Leafs["ack-waiting-for-dpm"] = types.YLeaf{"AckWaitingForDpm", summary.AckWaitingForDpm}
    summary.EntityData.Leafs["bound-clients"] = types.YLeaf{"BoundClients", summary.BoundClients}
    summary.EntityData.Leafs["renewing-clients"] = types.YLeaf{"RenewingClients", summary.RenewingClients}
    summary.EntityData.Leafs["informing-clients"] = types.YLeaf{"InformingClients", summary.InformingClients}
    summary.EntityData.Leafs["reauthorizing-clients"] = types.YLeaf{"ReauthorizingClients", summary.ReauthorizingClients}
    summary.EntityData.Leafs["waiting-for-dpm-disconnect"] = types.YLeaf{"WaitingForDpmDisconnect", summary.WaitingForDpmDisconnect}
    summary.EntityData.Leafs["waiting-for-dpm-addr-change"] = types.YLeaf{"WaitingForDpmAddrChange", summary.WaitingForDpmAddrChange}
    summary.EntityData.Leafs["deleting-clients-d"] = types.YLeaf{"DeletingClientsD", summary.DeletingClientsD}
    summary.EntityData.Leafs["disconnected-clients"] = types.YLeaf{"DisconnectedClients", summary.DisconnectedClients}
    summary.EntityData.Leafs["restarting-clients"] = types.YLeaf{"RestartingClients", summary.RestartingClients}
    return &(summary.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Interfaces
// IPv4 DHCP proxy/server Interface
type Ipv4Dhcpd_Nodes_Node_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 DHCP proxy/server interface info. The type is slice of
    // Ipv4Dhcpd_Nodes_Node_Interfaces_Interface_.
    Interface_ []Ipv4Dhcpd_Nodes_Node_Interfaces_Interface
}

func (interfaces *Ipv4Dhcpd_Nodes_Node_Interfaces) GetEntityData() *types.CommonEntityData {
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

// Ipv4Dhcpd_Nodes_Node_Interfaces_Interface
// IPv4 DHCP proxy/server interface info
type Ipv4Dhcpd_Nodes_Node_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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

func (self *Ipv4Dhcpd_Nodes_Node_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", self.InterfaceName}
    self.EntityData.Leafs["intf-ifhandle"] = types.YLeaf{"IntfIfhandle", self.IntfIfhandle}
    self.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", self.VrfName}
    self.EntityData.Leafs["intf-mode"] = types.YLeaf{"IntfMode", self.IntfMode}
    self.EntityData.Leafs["intf-is-ambiguous"] = types.YLeaf{"IntfIsAmbiguous", self.IntfIsAmbiguous}
    self.EntityData.Leafs["intf-profile-name"] = types.YLeaf{"IntfProfileName", self.IntfProfileName}
    self.EntityData.Leafs["intf-lease-limit-type"] = types.YLeaf{"IntfLeaseLimitType", self.IntfLeaseLimitType}
    self.EntityData.Leafs["intf-lease-limit-count"] = types.YLeaf{"IntfLeaseLimitCount", self.IntfLeaseLimitCount}
    self.EntityData.Leafs["srg-role"] = types.YLeaf{"SrgRole", self.SrgRole}
    self.EntityData.Leafs["mac-throttle"] = types.YLeaf{"MacThrottle", self.MacThrottle}
    return &(self.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Base
// IPv4 DHCP base operational data
type Ipv4Dhcpd_Nodes_Node_Base struct {
    EntityData types.CommonEntityData
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

func (base *Ipv4Dhcpd_Nodes_Node_Base) GetEntityData() *types.CommonEntityData {
    base.EntityData.YFilter = base.YFilter
    base.EntityData.YangName = "base"
    base.EntityData.BundleName = "cisco_ios_xr"
    base.EntityData.ParentYangName = "node"
    base.EntityData.SegmentPath = "base"
    base.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    base.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    base.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    base.EntityData.Children = make(map[string]types.YChild)
    base.EntityData.Children["statistics"] = types.YChild{"Statistics", &base.Statistics}
    base.EntityData.Children["issu-status"] = types.YChild{"IssuStatus", &base.IssuStatus}
    base.EntityData.Children["vrfs"] = types.YChild{"Vrfs", &base.Vrfs}
    base.EntityData.Children["profiles"] = types.YChild{"Profiles", &base.Profiles}
    base.EntityData.Children["database"] = types.YChild{"Database", &base.Database}
    base.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(base.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Base_Statistics
// DHCP base statistics
type Ipv4Dhcpd_Nodes_Node_Base_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ipv4 dhcpd proxy stat. The type is slice of
    // Ipv4Dhcpd_Nodes_Node_Base_Statistics_Ipv4DhcpdProxyStat.
    Ipv4DhcpdProxyStat []Ipv4Dhcpd_Nodes_Node_Base_Statistics_Ipv4DhcpdProxyStat
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Base_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "base"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = make(map[string]types.YChild)
    statistics.EntityData.Children["ipv4-dhcpd-proxy-stat"] = types.YChild{"Ipv4DhcpdProxyStat", nil}
    for i := range statistics.Ipv4DhcpdProxyStat {
        statistics.EntityData.Children[types.GetSegmentPath(&statistics.Ipv4DhcpdProxyStat[i])] = types.YChild{"Ipv4DhcpdProxyStat", &statistics.Ipv4DhcpdProxyStat[i]}
    }
    statistics.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(statistics.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Base_Statistics_Ipv4DhcpdProxyStat
// ipv4 dhcpd proxy stat
type Ipv4Dhcpd_Nodes_Node_Base_Statistics_Ipv4DhcpdProxyStat struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DHCP L3 VRF name. The type is string with length: 0..33.
    VrfName interface{}

    // Proxy statistics.
    Statistics Ipv4Dhcpd_Nodes_Node_Base_Statistics_Ipv4DhcpdProxyStat_Statistics_
}

func (ipv4DhcpdProxyStat *Ipv4Dhcpd_Nodes_Node_Base_Statistics_Ipv4DhcpdProxyStat) GetEntityData() *types.CommonEntityData {
    ipv4DhcpdProxyStat.EntityData.YFilter = ipv4DhcpdProxyStat.YFilter
    ipv4DhcpdProxyStat.EntityData.YangName = "ipv4-dhcpd-proxy-stat"
    ipv4DhcpdProxyStat.EntityData.BundleName = "cisco_ios_xr"
    ipv4DhcpdProxyStat.EntityData.ParentYangName = "statistics"
    ipv4DhcpdProxyStat.EntityData.SegmentPath = "ipv4-dhcpd-proxy-stat"
    ipv4DhcpdProxyStat.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4DhcpdProxyStat.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4DhcpdProxyStat.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4DhcpdProxyStat.EntityData.Children = make(map[string]types.YChild)
    ipv4DhcpdProxyStat.EntityData.Children["statistics"] = types.YChild{"Statistics", &ipv4DhcpdProxyStat.Statistics}
    ipv4DhcpdProxyStat.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv4DhcpdProxyStat.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", ipv4DhcpdProxyStat.VrfName}
    return &(ipv4DhcpdProxyStat.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Base_Statistics_Ipv4DhcpdProxyStat_Statistics_
// Proxy statistics
type Ipv4Dhcpd_Nodes_Node_Base_Statistics_Ipv4DhcpdProxyStat_Statistics_ struct {
    EntityData types.CommonEntityData
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

func (statistics_ *Ipv4Dhcpd_Nodes_Node_Base_Statistics_Ipv4DhcpdProxyStat_Statistics_) GetEntityData() *types.CommonEntityData {
    statistics_.EntityData.YFilter = statistics_.YFilter
    statistics_.EntityData.YangName = "statistics"
    statistics_.EntityData.BundleName = "cisco_ios_xr"
    statistics_.EntityData.ParentYangName = "ipv4-dhcpd-proxy-stat"
    statistics_.EntityData.SegmentPath = "statistics"
    statistics_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics_.EntityData.Children = make(map[string]types.YChild)
    statistics_.EntityData.Leafs = make(map[string]types.YLeaf)
    statistics_.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", statistics_.ReceivedPackets}
    statistics_.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", statistics_.TransmittedPackets}
    statistics_.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", statistics_.DroppedPackets}
    return &(statistics_.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Base_IssuStatus
// IPv4 DHCP ISSU status
type Ipv4Dhcpd_Nodes_Node_Base_IssuStatus struct {
    EntityData types.CommonEntityData
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

func (issuStatus *Ipv4Dhcpd_Nodes_Node_Base_IssuStatus) GetEntityData() *types.CommonEntityData {
    issuStatus.EntityData.YFilter = issuStatus.YFilter
    issuStatus.EntityData.YangName = "issu-status"
    issuStatus.EntityData.BundleName = "cisco_ios_xr"
    issuStatus.EntityData.ParentYangName = "base"
    issuStatus.EntityData.SegmentPath = "issu-status"
    issuStatus.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    issuStatus.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    issuStatus.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    issuStatus.EntityData.Children = make(map[string]types.YChild)
    issuStatus.EntityData.Leafs = make(map[string]types.YLeaf)
    issuStatus.EntityData.Leafs["issu-sync-complete-time"] = types.YLeaf{"IssuSyncCompleteTime", issuStatus.IssuSyncCompleteTime}
    issuStatus.EntityData.Leafs["issu-sync-start-time"] = types.YLeaf{"IssuSyncStartTime", issuStatus.IssuSyncStartTime}
    issuStatus.EntityData.Leafs["issu-ready-time"] = types.YLeaf{"IssuReadyTime", issuStatus.IssuReadyTime}
    issuStatus.EntityData.Leafs["big-bang-time"] = types.YLeaf{"BigBangTime", issuStatus.BigBangTime}
    issuStatus.EntityData.Leafs["primary-role-time"] = types.YLeaf{"PrimaryRoleTime", issuStatus.PrimaryRoleTime}
    issuStatus.EntityData.Leafs["role"] = types.YLeaf{"Role", issuStatus.Role}
    issuStatus.EntityData.Leafs["phase"] = types.YLeaf{"Phase", issuStatus.Phase}
    issuStatus.EntityData.Leafs["version"] = types.YLeaf{"Version", issuStatus.Version}
    issuStatus.EntityData.Leafs["issu-ready-issu-mgr-connection"] = types.YLeaf{"IssuReadyIssuMgrConnection", issuStatus.IssuReadyIssuMgrConnection}
    issuStatus.EntityData.Leafs["issu-ready-entries-replicate"] = types.YLeaf{"IssuReadyEntriesReplicate", issuStatus.IssuReadyEntriesReplicate}
    return &(issuStatus.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Base_Vrfs
// DHCP base list of VRF names
type Ipv4Dhcpd_Nodes_Node_Base_Vrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 DHCP base VRF name. The type is slice of
    // Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf.
    Vrf []Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf
}

func (vrfs *Ipv4Dhcpd_Nodes_Node_Base_Vrfs) GetEntityData() *types.CommonEntityData {
    vrfs.EntityData.YFilter = vrfs.YFilter
    vrfs.EntityData.YangName = "vrfs"
    vrfs.EntityData.BundleName = "cisco_ios_xr"
    vrfs.EntityData.ParentYangName = "base"
    vrfs.EntityData.SegmentPath = "vrfs"
    vrfs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfs.EntityData.Children = make(map[string]types.YChild)
    vrfs.EntityData.Children["vrf"] = types.YChild{"Vrf", nil}
    for i := range vrfs.Vrf {
        vrfs.EntityData.Children[types.GetSegmentPath(&vrfs.Vrf[i])] = types.YChild{"Vrf", &vrfs.Vrf[i]}
    }
    vrfs.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(vrfs.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf
// IPv4 DHCP base VRF name
type Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    VrfName interface{}

    // IPv4 DHCP base statistics.
    Statistics Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics
}

func (vrf *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "vrfs"
    vrf.EntityData.SegmentPath = "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = make(map[string]types.YChild)
    vrf.EntityData.Children["statistics"] = types.YChild{"Statistics", &vrf.Statistics}
    vrf.EntityData.Leafs = make(map[string]types.YLeaf)
    vrf.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", vrf.VrfName}
    return &(vrf.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics
// IPv4 DHCP base statistics
type Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics struct {
    EntityData types.CommonEntityData
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

func (statistics *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "vrf"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = make(map[string]types.YChild)
    statistics.EntityData.Children["discover"] = types.YChild{"Discover", &statistics.Discover}
    statistics.EntityData.Children["offer"] = types.YChild{"Offer", &statistics.Offer}
    statistics.EntityData.Children["request"] = types.YChild{"Request", &statistics.Request}
    statistics.EntityData.Children["decline"] = types.YChild{"Decline", &statistics.Decline}
    statistics.EntityData.Children["ack"] = types.YChild{"Ack", &statistics.Ack}
    statistics.EntityData.Children["nak"] = types.YChild{"Nak", &statistics.Nak}
    statistics.EntityData.Children["release"] = types.YChild{"Release", &statistics.Release}
    statistics.EntityData.Children["inform"] = types.YChild{"Inform", &statistics.Inform}
    statistics.EntityData.Children["lease-query"] = types.YChild{"LeaseQuery", &statistics.LeaseQuery}
    statistics.EntityData.Children["lease-not-assigned"] = types.YChild{"LeaseNotAssigned", &statistics.LeaseNotAssigned}
    statistics.EntityData.Children["lease-unknown"] = types.YChild{"LeaseUnknown", &statistics.LeaseUnknown}
    statistics.EntityData.Children["lease-active"] = types.YChild{"LeaseActive", &statistics.LeaseActive}
    statistics.EntityData.Children["bootp-request"] = types.YChild{"BootpRequest", &statistics.BootpRequest}
    statistics.EntityData.Children["bootp-reply"] = types.YChild{"BootpReply", &statistics.BootpReply}
    statistics.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(statistics.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Discover
// DHCP discover packets
type Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Discover struct {
    EntityData types.CommonEntityData
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

func (discover *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Discover) GetEntityData() *types.CommonEntityData {
    discover.EntityData.YFilter = discover.YFilter
    discover.EntityData.YangName = "discover"
    discover.EntityData.BundleName = "cisco_ios_xr"
    discover.EntityData.ParentYangName = "statistics"
    discover.EntityData.SegmentPath = "discover"
    discover.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    discover.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    discover.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    discover.EntityData.Children = make(map[string]types.YChild)
    discover.EntityData.Leafs = make(map[string]types.YLeaf)
    discover.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", discover.ReceivedPackets}
    discover.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", discover.TransmittedPackets}
    discover.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", discover.DroppedPackets}
    return &(discover.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Offer
// DHCP offer packets
type Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Offer struct {
    EntityData types.CommonEntityData
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

func (offer *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Offer) GetEntityData() *types.CommonEntityData {
    offer.EntityData.YFilter = offer.YFilter
    offer.EntityData.YangName = "offer"
    offer.EntityData.BundleName = "cisco_ios_xr"
    offer.EntityData.ParentYangName = "statistics"
    offer.EntityData.SegmentPath = "offer"
    offer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    offer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    offer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    offer.EntityData.Children = make(map[string]types.YChild)
    offer.EntityData.Leafs = make(map[string]types.YLeaf)
    offer.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", offer.ReceivedPackets}
    offer.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", offer.TransmittedPackets}
    offer.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", offer.DroppedPackets}
    return &(offer.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Request
// DHCP request packets
type Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Request struct {
    EntityData types.CommonEntityData
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

func (request *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Request) GetEntityData() *types.CommonEntityData {
    request.EntityData.YFilter = request.YFilter
    request.EntityData.YangName = "request"
    request.EntityData.BundleName = "cisco_ios_xr"
    request.EntityData.ParentYangName = "statistics"
    request.EntityData.SegmentPath = "request"
    request.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    request.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    request.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    request.EntityData.Children = make(map[string]types.YChild)
    request.EntityData.Leafs = make(map[string]types.YLeaf)
    request.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", request.ReceivedPackets}
    request.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", request.TransmittedPackets}
    request.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", request.DroppedPackets}
    return &(request.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Decline
// DHCP decline packets
type Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Decline struct {
    EntityData types.CommonEntityData
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

func (decline *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Decline) GetEntityData() *types.CommonEntityData {
    decline.EntityData.YFilter = decline.YFilter
    decline.EntityData.YangName = "decline"
    decline.EntityData.BundleName = "cisco_ios_xr"
    decline.EntityData.ParentYangName = "statistics"
    decline.EntityData.SegmentPath = "decline"
    decline.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    decline.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    decline.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    decline.EntityData.Children = make(map[string]types.YChild)
    decline.EntityData.Leafs = make(map[string]types.YLeaf)
    decline.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", decline.ReceivedPackets}
    decline.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", decline.TransmittedPackets}
    decline.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", decline.DroppedPackets}
    return &(decline.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Ack
// DHCP ack packets
type Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Ack struct {
    EntityData types.CommonEntityData
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

func (ack *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Ack) GetEntityData() *types.CommonEntityData {
    ack.EntityData.YFilter = ack.YFilter
    ack.EntityData.YangName = "ack"
    ack.EntityData.BundleName = "cisco_ios_xr"
    ack.EntityData.ParentYangName = "statistics"
    ack.EntityData.SegmentPath = "ack"
    ack.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ack.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ack.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ack.EntityData.Children = make(map[string]types.YChild)
    ack.EntityData.Leafs = make(map[string]types.YLeaf)
    ack.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", ack.ReceivedPackets}
    ack.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", ack.TransmittedPackets}
    ack.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", ack.DroppedPackets}
    return &(ack.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Nak
// DHCP nak packets
type Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Nak struct {
    EntityData types.CommonEntityData
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

func (nak *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Nak) GetEntityData() *types.CommonEntityData {
    nak.EntityData.YFilter = nak.YFilter
    nak.EntityData.YangName = "nak"
    nak.EntityData.BundleName = "cisco_ios_xr"
    nak.EntityData.ParentYangName = "statistics"
    nak.EntityData.SegmentPath = "nak"
    nak.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nak.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nak.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nak.EntityData.Children = make(map[string]types.YChild)
    nak.EntityData.Leafs = make(map[string]types.YLeaf)
    nak.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", nak.ReceivedPackets}
    nak.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", nak.TransmittedPackets}
    nak.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", nak.DroppedPackets}
    return &(nak.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Release
// DHCP release packets
type Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Release struct {
    EntityData types.CommonEntityData
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

func (release *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Release) GetEntityData() *types.CommonEntityData {
    release.EntityData.YFilter = release.YFilter
    release.EntityData.YangName = "release"
    release.EntityData.BundleName = "cisco_ios_xr"
    release.EntityData.ParentYangName = "statistics"
    release.EntityData.SegmentPath = "release"
    release.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    release.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    release.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    release.EntityData.Children = make(map[string]types.YChild)
    release.EntityData.Leafs = make(map[string]types.YLeaf)
    release.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", release.ReceivedPackets}
    release.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", release.TransmittedPackets}
    release.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", release.DroppedPackets}
    return &(release.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Inform
// DHCP inform packets
type Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Inform struct {
    EntityData types.CommonEntityData
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

func (inform *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_Inform) GetEntityData() *types.CommonEntityData {
    inform.EntityData.YFilter = inform.YFilter
    inform.EntityData.YangName = "inform"
    inform.EntityData.BundleName = "cisco_ios_xr"
    inform.EntityData.ParentYangName = "statistics"
    inform.EntityData.SegmentPath = "inform"
    inform.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inform.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inform.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inform.EntityData.Children = make(map[string]types.YChild)
    inform.EntityData.Leafs = make(map[string]types.YLeaf)
    inform.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", inform.ReceivedPackets}
    inform.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", inform.TransmittedPackets}
    inform.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", inform.DroppedPackets}
    return &(inform.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseQuery
// DHCP lease query packets
type Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseQuery struct {
    EntityData types.CommonEntityData
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

func (leaseQuery *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseQuery) GetEntityData() *types.CommonEntityData {
    leaseQuery.EntityData.YFilter = leaseQuery.YFilter
    leaseQuery.EntityData.YangName = "lease-query"
    leaseQuery.EntityData.BundleName = "cisco_ios_xr"
    leaseQuery.EntityData.ParentYangName = "statistics"
    leaseQuery.EntityData.SegmentPath = "lease-query"
    leaseQuery.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    leaseQuery.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    leaseQuery.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    leaseQuery.EntityData.Children = make(map[string]types.YChild)
    leaseQuery.EntityData.Leafs = make(map[string]types.YLeaf)
    leaseQuery.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", leaseQuery.ReceivedPackets}
    leaseQuery.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", leaseQuery.TransmittedPackets}
    leaseQuery.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", leaseQuery.DroppedPackets}
    return &(leaseQuery.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseNotAssigned
// DHCP lease not assigned
type Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseNotAssigned struct {
    EntityData types.CommonEntityData
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

func (leaseNotAssigned *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseNotAssigned) GetEntityData() *types.CommonEntityData {
    leaseNotAssigned.EntityData.YFilter = leaseNotAssigned.YFilter
    leaseNotAssigned.EntityData.YangName = "lease-not-assigned"
    leaseNotAssigned.EntityData.BundleName = "cisco_ios_xr"
    leaseNotAssigned.EntityData.ParentYangName = "statistics"
    leaseNotAssigned.EntityData.SegmentPath = "lease-not-assigned"
    leaseNotAssigned.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    leaseNotAssigned.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    leaseNotAssigned.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    leaseNotAssigned.EntityData.Children = make(map[string]types.YChild)
    leaseNotAssigned.EntityData.Leafs = make(map[string]types.YLeaf)
    leaseNotAssigned.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", leaseNotAssigned.ReceivedPackets}
    leaseNotAssigned.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", leaseNotAssigned.TransmittedPackets}
    leaseNotAssigned.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", leaseNotAssigned.DroppedPackets}
    return &(leaseNotAssigned.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseUnknown
// DHCP lease unknown
type Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseUnknown struct {
    EntityData types.CommonEntityData
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

func (leaseUnknown *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseUnknown) GetEntityData() *types.CommonEntityData {
    leaseUnknown.EntityData.YFilter = leaseUnknown.YFilter
    leaseUnknown.EntityData.YangName = "lease-unknown"
    leaseUnknown.EntityData.BundleName = "cisco_ios_xr"
    leaseUnknown.EntityData.ParentYangName = "statistics"
    leaseUnknown.EntityData.SegmentPath = "lease-unknown"
    leaseUnknown.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    leaseUnknown.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    leaseUnknown.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    leaseUnknown.EntityData.Children = make(map[string]types.YChild)
    leaseUnknown.EntityData.Leafs = make(map[string]types.YLeaf)
    leaseUnknown.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", leaseUnknown.ReceivedPackets}
    leaseUnknown.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", leaseUnknown.TransmittedPackets}
    leaseUnknown.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", leaseUnknown.DroppedPackets}
    return &(leaseUnknown.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseActive
// DHCP lease active
type Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseActive struct {
    EntityData types.CommonEntityData
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

func (leaseActive *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_LeaseActive) GetEntityData() *types.CommonEntityData {
    leaseActive.EntityData.YFilter = leaseActive.YFilter
    leaseActive.EntityData.YangName = "lease-active"
    leaseActive.EntityData.BundleName = "cisco_ios_xr"
    leaseActive.EntityData.ParentYangName = "statistics"
    leaseActive.EntityData.SegmentPath = "lease-active"
    leaseActive.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    leaseActive.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    leaseActive.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    leaseActive.EntityData.Children = make(map[string]types.YChild)
    leaseActive.EntityData.Leafs = make(map[string]types.YLeaf)
    leaseActive.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", leaseActive.ReceivedPackets}
    leaseActive.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", leaseActive.TransmittedPackets}
    leaseActive.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", leaseActive.DroppedPackets}
    return &(leaseActive.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_BootpRequest
// DHCP BOOTP request
type Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_BootpRequest struct {
    EntityData types.CommonEntityData
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

func (bootpRequest *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_BootpRequest) GetEntityData() *types.CommonEntityData {
    bootpRequest.EntityData.YFilter = bootpRequest.YFilter
    bootpRequest.EntityData.YangName = "bootp-request"
    bootpRequest.EntityData.BundleName = "cisco_ios_xr"
    bootpRequest.EntityData.ParentYangName = "statistics"
    bootpRequest.EntityData.SegmentPath = "bootp-request"
    bootpRequest.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bootpRequest.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bootpRequest.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bootpRequest.EntityData.Children = make(map[string]types.YChild)
    bootpRequest.EntityData.Leafs = make(map[string]types.YLeaf)
    bootpRequest.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", bootpRequest.ReceivedPackets}
    bootpRequest.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", bootpRequest.TransmittedPackets}
    bootpRequest.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", bootpRequest.DroppedPackets}
    return &(bootpRequest.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_BootpReply
// DHCP BOOTP reply
type Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_BootpReply struct {
    EntityData types.CommonEntityData
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

func (bootpReply *Ipv4Dhcpd_Nodes_Node_Base_Vrfs_Vrf_Statistics_BootpReply) GetEntityData() *types.CommonEntityData {
    bootpReply.EntityData.YFilter = bootpReply.YFilter
    bootpReply.EntityData.YangName = "bootp-reply"
    bootpReply.EntityData.BundleName = "cisco_ios_xr"
    bootpReply.EntityData.ParentYangName = "statistics"
    bootpReply.EntityData.SegmentPath = "bootp-reply"
    bootpReply.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bootpReply.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bootpReply.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bootpReply.EntityData.Children = make(map[string]types.YChild)
    bootpReply.EntityData.Leafs = make(map[string]types.YLeaf)
    bootpReply.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", bootpReply.ReceivedPackets}
    bootpReply.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", bootpReply.TransmittedPackets}
    bootpReply.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", bootpReply.DroppedPackets}
    return &(bootpReply.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Base_Profiles
// IPv4 DHCP Base profile
type Ipv4Dhcpd_Nodes_Node_Base_Profiles struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 DHCP base profile. The type is slice of
    // Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile.
    Profile []Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile
}

func (profiles *Ipv4Dhcpd_Nodes_Node_Base_Profiles) GetEntityData() *types.CommonEntityData {
    profiles.EntityData.YFilter = profiles.YFilter
    profiles.EntityData.YangName = "profiles"
    profiles.EntityData.BundleName = "cisco_ios_xr"
    profiles.EntityData.ParentYangName = "base"
    profiles.EntityData.SegmentPath = "profiles"
    profiles.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    profiles.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    profiles.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    profiles.EntityData.Children = make(map[string]types.YChild)
    profiles.EntityData.Children["profile"] = types.YChild{"Profile", nil}
    for i := range profiles.Profile {
        profiles.EntityData.Children[types.GetSegmentPath(&profiles.Profile[i])] = types.YChild{"Profile", &profiles.Profile[i]}
    }
    profiles.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(profiles.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile
// IPv4 DHCP base profile
type Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Profile name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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

func (profile *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile) GetEntityData() *types.CommonEntityData {
    profile.EntityData.YFilter = profile.YFilter
    profile.EntityData.YangName = "profile"
    profile.EntityData.BundleName = "cisco_ios_xr"
    profile.EntityData.ParentYangName = "profiles"
    profile.EntityData.SegmentPath = "profile" + "[profile-name='" + fmt.Sprintf("%v", profile.ProfileName) + "']"
    profile.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    profile.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    profile.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    profile.EntityData.Children = make(map[string]types.YChild)
    profile.EntityData.Children["interface-references"] = types.YChild{"InterfaceReferences", &profile.InterfaceReferences}
    profile.EntityData.Children["child-profile-info"] = types.YChild{"ChildProfileInfo", &profile.ChildProfileInfo}
    profile.EntityData.Leafs = make(map[string]types.YLeaf)
    profile.EntityData.Leafs["profile-name"] = types.YLeaf{"ProfileName", profile.ProfileName}
    profile.EntityData.Leafs["base-default-profile-name"] = types.YLeaf{"BaseDefaultProfileName", profile.BaseDefaultProfileName}
    profile.EntityData.Leafs["default-profile-mode"] = types.YLeaf{"DefaultProfileMode", profile.DefaultProfileMode}
    profile.EntityData.Leafs["relay-authenticate"] = types.YLeaf{"RelayAuthenticate", profile.RelayAuthenticate}
    profile.EntityData.Leafs["remote-id"] = types.YLeaf{"RemoteId", profile.RemoteId}
    profile.EntityData.Leafs["child-profile-count"] = types.YLeaf{"ChildProfileCount", profile.ChildProfileCount}
    profile.EntityData.Leafs["intf-ref-count"] = types.YLeaf{"IntfRefCount", profile.IntfRefCount}
    return &(profile.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_InterfaceReferences
// Interface references
type Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_InterfaceReferences struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ipv4 dhcpd base interface reference. The type is slice of
    // Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_InterfaceReferences_Ipv4DhcpdBaseInterfaceReference.
    Ipv4DhcpdBaseInterfaceReference []Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_InterfaceReferences_Ipv4DhcpdBaseInterfaceReference
}

func (interfaceReferences *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_InterfaceReferences) GetEntityData() *types.CommonEntityData {
    interfaceReferences.EntityData.YFilter = interfaceReferences.YFilter
    interfaceReferences.EntityData.YangName = "interface-references"
    interfaceReferences.EntityData.BundleName = "cisco_ios_xr"
    interfaceReferences.EntityData.ParentYangName = "profile"
    interfaceReferences.EntityData.SegmentPath = "interface-references"
    interfaceReferences.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceReferences.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceReferences.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceReferences.EntityData.Children = make(map[string]types.YChild)
    interfaceReferences.EntityData.Children["ipv4-dhcpd-base-interface-reference"] = types.YChild{"Ipv4DhcpdBaseInterfaceReference", nil}
    for i := range interfaceReferences.Ipv4DhcpdBaseInterfaceReference {
        interfaceReferences.EntityData.Children[types.GetSegmentPath(&interfaceReferences.Ipv4DhcpdBaseInterfaceReference[i])] = types.YChild{"Ipv4DhcpdBaseInterfaceReference", &interfaceReferences.Ipv4DhcpdBaseInterfaceReference[i]}
    }
    interfaceReferences.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaceReferences.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_InterfaceReferences_Ipv4DhcpdBaseInterfaceReference
// ipv4 dhcpd base interface reference
type Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_InterfaceReferences_Ipv4DhcpdBaseInterfaceReference struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface name. The type is string with length: 0..65.
    BaseReferenceInterfaceName interface{}
}

func (ipv4DhcpdBaseInterfaceReference *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_InterfaceReferences_Ipv4DhcpdBaseInterfaceReference) GetEntityData() *types.CommonEntityData {
    ipv4DhcpdBaseInterfaceReference.EntityData.YFilter = ipv4DhcpdBaseInterfaceReference.YFilter
    ipv4DhcpdBaseInterfaceReference.EntityData.YangName = "ipv4-dhcpd-base-interface-reference"
    ipv4DhcpdBaseInterfaceReference.EntityData.BundleName = "cisco_ios_xr"
    ipv4DhcpdBaseInterfaceReference.EntityData.ParentYangName = "interface-references"
    ipv4DhcpdBaseInterfaceReference.EntityData.SegmentPath = "ipv4-dhcpd-base-interface-reference"
    ipv4DhcpdBaseInterfaceReference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4DhcpdBaseInterfaceReference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4DhcpdBaseInterfaceReference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4DhcpdBaseInterfaceReference.EntityData.Children = make(map[string]types.YChild)
    ipv4DhcpdBaseInterfaceReference.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv4DhcpdBaseInterfaceReference.EntityData.Leafs["base-reference-interface-name"] = types.YLeaf{"BaseReferenceInterfaceName", ipv4DhcpdBaseInterfaceReference.BaseReferenceInterfaceName}
    return &(ipv4DhcpdBaseInterfaceReference.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_ChildProfileInfo
// child profile info
type Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_ChildProfileInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ipv4 dhcpd base child profile info. The type is slice of
    // Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_ChildProfileInfo_Ipv4DhcpdBaseChildProfileInfo.
    Ipv4DhcpdBaseChildProfileInfo []Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_ChildProfileInfo_Ipv4DhcpdBaseChildProfileInfo
}

func (childProfileInfo *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_ChildProfileInfo) GetEntityData() *types.CommonEntityData {
    childProfileInfo.EntityData.YFilter = childProfileInfo.YFilter
    childProfileInfo.EntityData.YangName = "child-profile-info"
    childProfileInfo.EntityData.BundleName = "cisco_ios_xr"
    childProfileInfo.EntityData.ParentYangName = "profile"
    childProfileInfo.EntityData.SegmentPath = "child-profile-info"
    childProfileInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    childProfileInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    childProfileInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    childProfileInfo.EntityData.Children = make(map[string]types.YChild)
    childProfileInfo.EntityData.Children["ipv4-dhcpd-base-child-profile-info"] = types.YChild{"Ipv4DhcpdBaseChildProfileInfo", nil}
    for i := range childProfileInfo.Ipv4DhcpdBaseChildProfileInfo {
        childProfileInfo.EntityData.Children[types.GetSegmentPath(&childProfileInfo.Ipv4DhcpdBaseChildProfileInfo[i])] = types.YChild{"Ipv4DhcpdBaseChildProfileInfo", &childProfileInfo.Ipv4DhcpdBaseChildProfileInfo[i]}
    }
    childProfileInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(childProfileInfo.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_ChildProfileInfo_Ipv4DhcpdBaseChildProfileInfo
// ipv4 dhcpd base child profile info
type Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_ChildProfileInfo_Ipv4DhcpdBaseChildProfileInfo struct {
    EntityData types.CommonEntityData
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
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    OptionData interface{}
}

func (ipv4DhcpdBaseChildProfileInfo *Ipv4Dhcpd_Nodes_Node_Base_Profiles_Profile_ChildProfileInfo_Ipv4DhcpdBaseChildProfileInfo) GetEntityData() *types.CommonEntityData {
    ipv4DhcpdBaseChildProfileInfo.EntityData.YFilter = ipv4DhcpdBaseChildProfileInfo.YFilter
    ipv4DhcpdBaseChildProfileInfo.EntityData.YangName = "ipv4-dhcpd-base-child-profile-info"
    ipv4DhcpdBaseChildProfileInfo.EntityData.BundleName = "cisco_ios_xr"
    ipv4DhcpdBaseChildProfileInfo.EntityData.ParentYangName = "child-profile-info"
    ipv4DhcpdBaseChildProfileInfo.EntityData.SegmentPath = "ipv4-dhcpd-base-child-profile-info"
    ipv4DhcpdBaseChildProfileInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4DhcpdBaseChildProfileInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4DhcpdBaseChildProfileInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4DhcpdBaseChildProfileInfo.EntityData.Children = make(map[string]types.YChild)
    ipv4DhcpdBaseChildProfileInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv4DhcpdBaseChildProfileInfo.EntityData.Leafs["base-child-profile-name"] = types.YLeaf{"BaseChildProfileName", ipv4DhcpdBaseChildProfileInfo.BaseChildProfileName}
    ipv4DhcpdBaseChildProfileInfo.EntityData.Leafs["mode"] = types.YLeaf{"Mode", ipv4DhcpdBaseChildProfileInfo.Mode}
    ipv4DhcpdBaseChildProfileInfo.EntityData.Leafs["matched-option-code"] = types.YLeaf{"MatchedOptionCode", ipv4DhcpdBaseChildProfileInfo.MatchedOptionCode}
    ipv4DhcpdBaseChildProfileInfo.EntityData.Leafs["matched-option-len"] = types.YLeaf{"MatchedOptionLen", ipv4DhcpdBaseChildProfileInfo.MatchedOptionLen}
    ipv4DhcpdBaseChildProfileInfo.EntityData.Leafs["option-data"] = types.YLeaf{"OptionData", ipv4DhcpdBaseChildProfileInfo.OptionData}
    return &(ipv4DhcpdBaseChildProfileInfo.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Base_Database
// IPv4 DHCP database
type Ipv4Dhcpd_Nodes_Node_Base_Database struct {
    EntityData types.CommonEntityData
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

func (database *Ipv4Dhcpd_Nodes_Node_Base_Database) GetEntityData() *types.CommonEntityData {
    database.EntityData.YFilter = database.YFilter
    database.EntityData.YangName = "database"
    database.EntityData.BundleName = "cisco_ios_xr"
    database.EntityData.ParentYangName = "base"
    database.EntityData.SegmentPath = "database"
    database.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    database.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    database.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    database.EntityData.Children = make(map[string]types.YChild)
    database.EntityData.Leafs = make(map[string]types.YLeaf)
    database.EntityData.Leafs["configured"] = types.YLeaf{"Configured", database.Configured}
    database.EntityData.Leafs["version"] = types.YLeaf{"Version", database.Version}
    database.EntityData.Leafs["full-file-write-interval"] = types.YLeaf{"FullFileWriteInterval", database.FullFileWriteInterval}
    database.EntityData.Leafs["last-full-write-file-name"] = types.YLeaf{"LastFullWriteFileName", database.LastFullWriteFileName}
    database.EntityData.Leafs["last-full-write-time"] = types.YLeaf{"LastFullWriteTime", database.LastFullWriteTime}
    database.EntityData.Leafs["full-file-write-count"] = types.YLeaf{"FullFileWriteCount", database.FullFileWriteCount}
    database.EntityData.Leafs["failed-full-file-write-count"] = types.YLeaf{"FailedFullFileWriteCount", database.FailedFullFileWriteCount}
    database.EntityData.Leafs["full-file-record-count"] = types.YLeaf{"FullFileRecordCount", database.FullFileRecordCount}
    database.EntityData.Leafs["last-full-file-write-error-timestamp"] = types.YLeaf{"LastFullFileWriteErrorTimestamp", database.LastFullFileWriteErrorTimestamp}
    database.EntityData.Leafs["incremental-file-write-interval"] = types.YLeaf{"IncrementalFileWriteInterval", database.IncrementalFileWriteInterval}
    database.EntityData.Leafs["last-incremental-write-file-name"] = types.YLeaf{"LastIncrementalWriteFileName", database.LastIncrementalWriteFileName}
    database.EntityData.Leafs["last-incremental-write-time"] = types.YLeaf{"LastIncrementalWriteTime", database.LastIncrementalWriteTime}
    database.EntityData.Leafs["incremental-file-write-count"] = types.YLeaf{"IncrementalFileWriteCount", database.IncrementalFileWriteCount}
    database.EntityData.Leafs["failed-incremental-file-write-count"] = types.YLeaf{"FailedIncrementalFileWriteCount", database.FailedIncrementalFileWriteCount}
    database.EntityData.Leafs["incremental-file-record-count"] = types.YLeaf{"IncrementalFileRecordCount", database.IncrementalFileRecordCount}
    database.EntityData.Leafs["last-incremental-file-write-error-timestamp"] = types.YLeaf{"LastIncrementalFileWriteErrorTimestamp", database.LastIncrementalFileWriteErrorTimestamp}
    return &(database.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Server
// IPv4 DHCP Server operational data
type Ipv4Dhcpd_Nodes_Node_Server struct {
    EntityData types.CommonEntityData
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

func (server *Ipv4Dhcpd_Nodes_Node_Server) GetEntityData() *types.CommonEntityData {
    server.EntityData.YFilter = server.YFilter
    server.EntityData.YangName = "server"
    server.EntityData.BundleName = "cisco_ios_xr"
    server.EntityData.ParentYangName = "node"
    server.EntityData.SegmentPath = "server"
    server.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    server.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    server.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    server.EntityData.Children = make(map[string]types.YChild)
    server.EntityData.Children["profiles"] = types.YChild{"Profiles", &server.Profiles}
    server.EntityData.Children["statistics"] = types.YChild{"Statistics", &server.Statistics}
    server.EntityData.Children["binding"] = types.YChild{"Binding", &server.Binding}
    server.EntityData.Children["statistics-info"] = types.YChild{"StatisticsInfo", &server.StatisticsInfo}
    server.EntityData.Children["vrfs"] = types.YChild{"Vrfs", &server.Vrfs}
    server.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(server.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Server_Profiles
// IPv4 DHCP Server profile
type Ipv4Dhcpd_Nodes_Node_Server_Profiles struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 DHCP server profile. The type is slice of
    // Ipv4Dhcpd_Nodes_Node_Server_Profiles_Profile.
    Profile []Ipv4Dhcpd_Nodes_Node_Server_Profiles_Profile
}

func (profiles *Ipv4Dhcpd_Nodes_Node_Server_Profiles) GetEntityData() *types.CommonEntityData {
    profiles.EntityData.YFilter = profiles.YFilter
    profiles.EntityData.YangName = "profiles"
    profiles.EntityData.BundleName = "cisco_ios_xr"
    profiles.EntityData.ParentYangName = "server"
    profiles.EntityData.SegmentPath = "profiles"
    profiles.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    profiles.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    profiles.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    profiles.EntityData.Children = make(map[string]types.YChild)
    profiles.EntityData.Children["profile"] = types.YChild{"Profile", nil}
    for i := range profiles.Profile {
        profiles.EntityData.Children[types.GetSegmentPath(&profiles.Profile[i])] = types.YChild{"Profile", &profiles.Profile[i]}
    }
    profiles.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(profiles.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Server_Profiles_Profile
// IPv4 DHCP server profile
type Ipv4Dhcpd_Nodes_Node_Server_Profiles_Profile struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Profile name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    ServerProfileDns []interface{}

    // Server default addresses. The type is slice of string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    ServerProfileDefaultRouter []interface{}

    // Server netbios addresses. The type is slice of string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    ServerProfileNetbiousNameServer []interface{}

    // Server Time addresses. The type is slice of string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    ServerProfileTimeServer []interface{}
}

func (profile *Ipv4Dhcpd_Nodes_Node_Server_Profiles_Profile) GetEntityData() *types.CommonEntityData {
    profile.EntityData.YFilter = profile.YFilter
    profile.EntityData.YangName = "profile"
    profile.EntityData.BundleName = "cisco_ios_xr"
    profile.EntityData.ParentYangName = "profiles"
    profile.EntityData.SegmentPath = "profile" + "[server-profile-name='" + fmt.Sprintf("%v", profile.ServerProfileName) + "']"
    profile.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    profile.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    profile.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    profile.EntityData.Children = make(map[string]types.YChild)
    profile.EntityData.Leafs = make(map[string]types.YLeaf)
    profile.EntityData.Leafs["server-profile-name"] = types.YLeaf{"ServerProfileName", profile.ServerProfileName}
    profile.EntityData.Leafs["server-profile-name-xr"] = types.YLeaf{"ServerProfileNameXr", profile.ServerProfileNameXr}
    profile.EntityData.Leafs["secure-arp"] = types.YLeaf{"SecureArp", profile.SecureArp}
    profile.EntityData.Leafs["requested-address-check"] = types.YLeaf{"RequestedAddressCheck", profile.RequestedAddressCheck}
    profile.EntityData.Leafs["server-id-check"] = types.YLeaf{"ServerIdCheck", profile.ServerIdCheck}
    profile.EntityData.Leafs["duplicate-mac-address-check"] = types.YLeaf{"DuplicateMacAddressCheck", profile.DuplicateMacAddressCheck}
    profile.EntityData.Leafs["duplicate-ip-address-check"] = types.YLeaf{"DuplicateIpAddressCheck", profile.DuplicateIpAddressCheck}
    profile.EntityData.Leafs["is-move-allowed"] = types.YLeaf{"IsMoveAllowed", profile.IsMoveAllowed}
    profile.EntityData.Leafs["bcast-policy"] = types.YLeaf{"BcastPolicy", profile.BcastPolicy}
    profile.EntityData.Leafs["giaddr-policy"] = types.YLeaf{"GiaddrPolicy", profile.GiaddrPolicy}
    profile.EntityData.Leafs["subnet-mask"] = types.YLeaf{"SubnetMask", profile.SubnetMask}
    profile.EntityData.Leafs["server-pool-name"] = types.YLeaf{"ServerPoolName", profile.ServerPoolName}
    profile.EntityData.Leafs["server-profile-lease"] = types.YLeaf{"ServerProfileLease", profile.ServerProfileLease}
    profile.EntityData.Leafs["server-profile-netbios-node-type"] = types.YLeaf{"ServerProfileNetbiosNodeType", profile.ServerProfileNetbiosNodeType}
    profile.EntityData.Leafs["server-bootfile-name"] = types.YLeaf{"ServerBootfileName", profile.ServerBootfileName}
    profile.EntityData.Leafs["server-domain-name"] = types.YLeaf{"ServerDomainName", profile.ServerDomainName}
    profile.EntityData.Leafs["server-profileiedge-check"] = types.YLeaf{"ServerProfileiedgeCheck", profile.ServerProfileiedgeCheck}
    profile.EntityData.Leafs["server-profile-server-dns-count"] = types.YLeaf{"ServerProfileServerDnsCount", profile.ServerProfileServerDnsCount}
    profile.EntityData.Leafs["server-profiledefault-router-count"] = types.YLeaf{"ServerProfiledefaultRouterCount", profile.ServerProfiledefaultRouterCount}
    profile.EntityData.Leafs["server-profile-netbios-name-svr-count"] = types.YLeaf{"ServerProfileNetbiosNameSvrCount", profile.ServerProfileNetbiosNameSvrCount}
    profile.EntityData.Leafs["server-profile-time-svr-count"] = types.YLeaf{"ServerProfileTimeSvrCount", profile.ServerProfileTimeSvrCount}
    profile.EntityData.Leafs["lease-limit-type"] = types.YLeaf{"LeaseLimitType", profile.LeaseLimitType}
    profile.EntityData.Leafs["lease-limit-count"] = types.YLeaf{"LeaseLimitCount", profile.LeaseLimitCount}
    profile.EntityData.Leafs["server-profile-dns"] = types.YLeaf{"ServerProfileDns", profile.ServerProfileDns}
    profile.EntityData.Leafs["server-profile-default-router"] = types.YLeaf{"ServerProfileDefaultRouter", profile.ServerProfileDefaultRouter}
    profile.EntityData.Leafs["server-profile-netbious-name-server"] = types.YLeaf{"ServerProfileNetbiousNameServer", profile.ServerProfileNetbiousNameServer}
    profile.EntityData.Leafs["server-profile-time-server"] = types.YLeaf{"ServerProfileTimeServer", profile.ServerProfileTimeServer}
    return &(profile.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Server_Statistics
// DHCP Server statistics
type Ipv4Dhcpd_Nodes_Node_Server_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ipv4 dhcpd proxy stat. The type is slice of
    // Ipv4Dhcpd_Nodes_Node_Server_Statistics_Ipv4DhcpdProxyStat.
    Ipv4DhcpdProxyStat []Ipv4Dhcpd_Nodes_Node_Server_Statistics_Ipv4DhcpdProxyStat
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Server_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "server"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = make(map[string]types.YChild)
    statistics.EntityData.Children["ipv4-dhcpd-proxy-stat"] = types.YChild{"Ipv4DhcpdProxyStat", nil}
    for i := range statistics.Ipv4DhcpdProxyStat {
        statistics.EntityData.Children[types.GetSegmentPath(&statistics.Ipv4DhcpdProxyStat[i])] = types.YChild{"Ipv4DhcpdProxyStat", &statistics.Ipv4DhcpdProxyStat[i]}
    }
    statistics.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(statistics.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Server_Statistics_Ipv4DhcpdProxyStat
// ipv4 dhcpd proxy stat
type Ipv4Dhcpd_Nodes_Node_Server_Statistics_Ipv4DhcpdProxyStat struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DHCP L3 VRF name. The type is string with length: 0..33.
    VrfName interface{}

    // Proxy statistics.
    Statistics Ipv4Dhcpd_Nodes_Node_Server_Statistics_Ipv4DhcpdProxyStat_Statistics_
}

func (ipv4DhcpdProxyStat *Ipv4Dhcpd_Nodes_Node_Server_Statistics_Ipv4DhcpdProxyStat) GetEntityData() *types.CommonEntityData {
    ipv4DhcpdProxyStat.EntityData.YFilter = ipv4DhcpdProxyStat.YFilter
    ipv4DhcpdProxyStat.EntityData.YangName = "ipv4-dhcpd-proxy-stat"
    ipv4DhcpdProxyStat.EntityData.BundleName = "cisco_ios_xr"
    ipv4DhcpdProxyStat.EntityData.ParentYangName = "statistics"
    ipv4DhcpdProxyStat.EntityData.SegmentPath = "ipv4-dhcpd-proxy-stat"
    ipv4DhcpdProxyStat.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4DhcpdProxyStat.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4DhcpdProxyStat.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4DhcpdProxyStat.EntityData.Children = make(map[string]types.YChild)
    ipv4DhcpdProxyStat.EntityData.Children["statistics"] = types.YChild{"Statistics", &ipv4DhcpdProxyStat.Statistics}
    ipv4DhcpdProxyStat.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv4DhcpdProxyStat.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", ipv4DhcpdProxyStat.VrfName}
    return &(ipv4DhcpdProxyStat.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Server_Statistics_Ipv4DhcpdProxyStat_Statistics_
// Proxy statistics
type Ipv4Dhcpd_Nodes_Node_Server_Statistics_Ipv4DhcpdProxyStat_Statistics_ struct {
    EntityData types.CommonEntityData
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

func (statistics_ *Ipv4Dhcpd_Nodes_Node_Server_Statistics_Ipv4DhcpdProxyStat_Statistics_) GetEntityData() *types.CommonEntityData {
    statistics_.EntityData.YFilter = statistics_.YFilter
    statistics_.EntityData.YangName = "statistics"
    statistics_.EntityData.BundleName = "cisco_ios_xr"
    statistics_.EntityData.ParentYangName = "ipv4-dhcpd-proxy-stat"
    statistics_.EntityData.SegmentPath = "statistics"
    statistics_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics_.EntityData.Children = make(map[string]types.YChild)
    statistics_.EntityData.Leafs = make(map[string]types.YLeaf)
    statistics_.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", statistics_.ReceivedPackets}
    statistics_.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", statistics_.TransmittedPackets}
    statistics_.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", statistics_.DroppedPackets}
    return &(statistics_.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Server_Binding
// DHCP server bindings
type Ipv4Dhcpd_Nodes_Node_Server_Binding struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DHCP server binding summary.
    Summary Ipv4Dhcpd_Nodes_Node_Server_Binding_Summary

    // DHCP server client bindings.
    Clients Ipv4Dhcpd_Nodes_Node_Server_Binding_Clients
}

func (binding *Ipv4Dhcpd_Nodes_Node_Server_Binding) GetEntityData() *types.CommonEntityData {
    binding.EntityData.YFilter = binding.YFilter
    binding.EntityData.YangName = "binding"
    binding.EntityData.BundleName = "cisco_ios_xr"
    binding.EntityData.ParentYangName = "server"
    binding.EntityData.SegmentPath = "binding"
    binding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    binding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    binding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    binding.EntityData.Children = make(map[string]types.YChild)
    binding.EntityData.Children["summary"] = types.YChild{"Summary", &binding.Summary}
    binding.EntityData.Children["clients"] = types.YChild{"Clients", &binding.Clients}
    binding.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(binding.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Server_Binding_Summary
// DHCP server binding summary
type Ipv4Dhcpd_Nodes_Node_Server_Binding_Summary struct {
    EntityData types.CommonEntityData
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

func (summary *Ipv4Dhcpd_Nodes_Node_Server_Binding_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "binding"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = make(map[string]types.YChild)
    summary.EntityData.Leafs = make(map[string]types.YLeaf)
    summary.EntityData.Leafs["clients"] = types.YLeaf{"Clients", summary.Clients}
    summary.EntityData.Leafs["initializing-clients"] = types.YLeaf{"InitializingClients", summary.InitializingClients}
    summary.EntityData.Leafs["waiting-for-dpm-init"] = types.YLeaf{"WaitingForDpmInit", summary.WaitingForDpmInit}
    summary.EntityData.Leafs["waiting-for-dpm-request"] = types.YLeaf{"WaitingForDpmRequest", summary.WaitingForDpmRequest}
    summary.EntityData.Leafs["waiting-for-daps-init"] = types.YLeaf{"WaitingForDapsInit", summary.WaitingForDapsInit}
    summary.EntityData.Leafs["selecting-clients"] = types.YLeaf{"SelectingClients", summary.SelectingClients}
    summary.EntityData.Leafs["offer-sent-for-client"] = types.YLeaf{"OfferSentForClient", summary.OfferSentForClient}
    summary.EntityData.Leafs["requesting-clients"] = types.YLeaf{"RequestingClients", summary.RequestingClients}
    summary.EntityData.Leafs["request-waiting-for-dpm"] = types.YLeaf{"RequestWaitingForDpm", summary.RequestWaitingForDpm}
    summary.EntityData.Leafs["ack-waiting-for-dpm"] = types.YLeaf{"AckWaitingForDpm", summary.AckWaitingForDpm}
    summary.EntityData.Leafs["bound-clients"] = types.YLeaf{"BoundClients", summary.BoundClients}
    summary.EntityData.Leafs["renewing-clients"] = types.YLeaf{"RenewingClients", summary.RenewingClients}
    summary.EntityData.Leafs["informing-clients"] = types.YLeaf{"InformingClients", summary.InformingClients}
    summary.EntityData.Leafs["reauthorizing-clients"] = types.YLeaf{"ReauthorizingClients", summary.ReauthorizingClients}
    summary.EntityData.Leafs["waiting-for-dpm-disconnect"] = types.YLeaf{"WaitingForDpmDisconnect", summary.WaitingForDpmDisconnect}
    summary.EntityData.Leafs["waiting-for-dpm-addr-change"] = types.YLeaf{"WaitingForDpmAddrChange", summary.WaitingForDpmAddrChange}
    summary.EntityData.Leafs["deleting-clients-d"] = types.YLeaf{"DeletingClientsD", summary.DeletingClientsD}
    summary.EntityData.Leafs["disconnected-clients"] = types.YLeaf{"DisconnectedClients", summary.DisconnectedClients}
    summary.EntityData.Leafs["restarting-clients"] = types.YLeaf{"RestartingClients", summary.RestartingClients}
    return &(summary.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Server_Binding_Clients
// DHCP server client bindings
type Ipv4Dhcpd_Nodes_Node_Server_Binding_Clients struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Single DHCP Server binding. The type is slice of
    // Ipv4Dhcpd_Nodes_Node_Server_Binding_Clients_Client.
    Client []Ipv4Dhcpd_Nodes_Node_Server_Binding_Clients_Client
}

func (clients *Ipv4Dhcpd_Nodes_Node_Server_Binding_Clients) GetEntityData() *types.CommonEntityData {
    clients.EntityData.YFilter = clients.YFilter
    clients.EntityData.YangName = "clients"
    clients.EntityData.BundleName = "cisco_ios_xr"
    clients.EntityData.ParentYangName = "binding"
    clients.EntityData.SegmentPath = "clients"
    clients.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clients.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clients.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clients.EntityData.Children = make(map[string]types.YChild)
    clients.EntityData.Children["client"] = types.YChild{"Client", nil}
    for i := range clients.Client {
        clients.EntityData.Children[types.GetSegmentPath(&clients.Client[i])] = types.YChild{"Client", &clients.Client[i]}
    }
    clients.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(clients.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Server_Binding_Clients_Client
// Single DHCP Server binding
type Ipv4Dhcpd_Nodes_Node_Server_Binding_Clients_Client struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Client ID. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    IpAddress interface{}

    // DHCP client GIADDR. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    ClientGiAddr interface{}

    // DHCP to server GIADDR. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    ToServerGiAddr interface{}

    // DHCP server IP address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    ServerIpAddress interface{}

    // DHCP reply server IP address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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
    // b'[a-zA-Z0-9./-]+'.
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
    // b'[a-zA-Z0-9./-]+'.
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

func (client *Ipv4Dhcpd_Nodes_Node_Server_Binding_Clients_Client) GetEntityData() *types.CommonEntityData {
    client.EntityData.YFilter = client.YFilter
    client.EntityData.YangName = "client"
    client.EntityData.BundleName = "cisco_ios_xr"
    client.EntityData.ParentYangName = "clients"
    client.EntityData.SegmentPath = "client" + "[client-id='" + fmt.Sprintf("%v", client.ClientId) + "']"
    client.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    client.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    client.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    client.EntityData.Children = make(map[string]types.YChild)
    client.EntityData.Leafs = make(map[string]types.YLeaf)
    client.EntityData.Leafs["client-id"] = types.YLeaf{"ClientId", client.ClientId}
    client.EntityData.Leafs["client-id-xr"] = types.YLeaf{"ClientIdXr", client.ClientIdXr}
    client.EntityData.Leafs["mac-address"] = types.YLeaf{"MacAddress", client.MacAddress}
    client.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", client.VrfName}
    client.EntityData.Leafs["server-vrf-name"] = types.YLeaf{"ServerVrfName", client.ServerVrfName}
    client.EntityData.Leafs["ip-address"] = types.YLeaf{"IpAddress", client.IpAddress}
    client.EntityData.Leafs["client-gi-addr"] = types.YLeaf{"ClientGiAddr", client.ClientGiAddr}
    client.EntityData.Leafs["to-server-gi-addr"] = types.YLeaf{"ToServerGiAddr", client.ToServerGiAddr}
    client.EntityData.Leafs["server-ip-address"] = types.YLeaf{"ServerIpAddress", client.ServerIpAddress}
    client.EntityData.Leafs["reply-server-ip-address"] = types.YLeaf{"ReplyServerIpAddress", client.ReplyServerIpAddress}
    client.EntityData.Leafs["lease-time"] = types.YLeaf{"LeaseTime", client.LeaseTime}
    client.EntityData.Leafs["remaining-lease-time"] = types.YLeaf{"RemainingLeaseTime", client.RemainingLeaseTime}
    client.EntityData.Leafs["state"] = types.YLeaf{"State", client.State}
    client.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", client.InterfaceName}
    client.EntityData.Leafs["access-vrf-name"] = types.YLeaf{"AccessVrfName", client.AccessVrfName}
    client.EntityData.Leafs["proxy-binding-outer-tag"] = types.YLeaf{"ProxyBindingOuterTag", client.ProxyBindingOuterTag}
    client.EntityData.Leafs["proxy-binding-inner-tag"] = types.YLeaf{"ProxyBindingInnerTag", client.ProxyBindingInnerTag}
    client.EntityData.Leafs["profile-name"] = types.YLeaf{"ProfileName", client.ProfileName}
    client.EntityData.Leafs["is-nak-next-renew"] = types.YLeaf{"IsNakNextRenew", client.IsNakNextRenew}
    client.EntityData.Leafs["subscriber-label"] = types.YLeaf{"SubscriberLabel", client.SubscriberLabel}
    client.EntityData.Leafs["old-subscriber-label"] = types.YLeaf{"OldSubscriberLabel", client.OldSubscriberLabel}
    client.EntityData.Leafs["subscriber-interface-name"] = types.YLeaf{"SubscriberInterfaceName", client.SubscriberInterfaceName}
    client.EntityData.Leafs["rx-circuit-id"] = types.YLeaf{"RxCircuitId", client.RxCircuitId}
    client.EntityData.Leafs["tx-circuit-id"] = types.YLeaf{"TxCircuitId", client.TxCircuitId}
    client.EntityData.Leafs["rx-remote-id"] = types.YLeaf{"RxRemoteId", client.RxRemoteId}
    client.EntityData.Leafs["tx-remote-id"] = types.YLeaf{"TxRemoteId", client.TxRemoteId}
    client.EntityData.Leafs["rx-vsiso"] = types.YLeaf{"RxVsiso", client.RxVsiso}
    client.EntityData.Leafs["tx-vsiso"] = types.YLeaf{"TxVsiso", client.TxVsiso}
    client.EntityData.Leafs["is-auth-received"] = types.YLeaf{"IsAuthReceived", client.IsAuthReceived}
    client.EntityData.Leafs["is-mbl-subscriber"] = types.YLeaf{"IsMblSubscriber", client.IsMblSubscriber}
    client.EntityData.Leafs["param-request"] = types.YLeaf{"ParamRequest", client.ParamRequest}
    client.EntityData.Leafs["param-response"] = types.YLeaf{"ParamResponse", client.ParamResponse}
    client.EntityData.Leafs["session-start-time"] = types.YLeaf{"SessionStartTime", client.SessionStartTime}
    client.EntityData.Leafs["srg-state"] = types.YLeaf{"SrgState", client.SrgState}
    client.EntityData.Leafs["event-history"] = types.YLeaf{"EventHistory", client.EventHistory}
    return &(client.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Server_StatisticsInfo
// DHCP proxy stats info
type Ipv4Dhcpd_Nodes_Node_Server_StatisticsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Proxy Stats timestamp. The type is interface{} with range: 0..4294967295.
    ProxyStatsTimestamp interface{}
}

func (statisticsInfo *Ipv4Dhcpd_Nodes_Node_Server_StatisticsInfo) GetEntityData() *types.CommonEntityData {
    statisticsInfo.EntityData.YFilter = statisticsInfo.YFilter
    statisticsInfo.EntityData.YangName = "statistics-info"
    statisticsInfo.EntityData.BundleName = "cisco_ios_xr"
    statisticsInfo.EntityData.ParentYangName = "server"
    statisticsInfo.EntityData.SegmentPath = "statistics-info"
    statisticsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statisticsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statisticsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statisticsInfo.EntityData.Children = make(map[string]types.YChild)
    statisticsInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    statisticsInfo.EntityData.Leafs["proxy-stats-timestamp"] = types.YLeaf{"ProxyStatsTimestamp", statisticsInfo.ProxyStatsTimestamp}
    return &(statisticsInfo.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Server_Vrfs
// DHCP Server list of VRF names
type Ipv4Dhcpd_Nodes_Node_Server_Vrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 DHCP server VRF name. The type is slice of
    // Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf.
    Vrf []Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf
}

func (vrfs *Ipv4Dhcpd_Nodes_Node_Server_Vrfs) GetEntityData() *types.CommonEntityData {
    vrfs.EntityData.YFilter = vrfs.YFilter
    vrfs.EntityData.YangName = "vrfs"
    vrfs.EntityData.BundleName = "cisco_ios_xr"
    vrfs.EntityData.ParentYangName = "server"
    vrfs.EntityData.SegmentPath = "vrfs"
    vrfs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfs.EntityData.Children = make(map[string]types.YChild)
    vrfs.EntityData.Children["vrf"] = types.YChild{"Vrf", nil}
    for i := range vrfs.Vrf {
        vrfs.EntityData.Children[types.GetSegmentPath(&vrfs.Vrf[i])] = types.YChild{"Vrf", &vrfs.Vrf[i]}
    }
    vrfs.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(vrfs.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf
// IPv4 DHCP server VRF name
type Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    VrfName interface{}

    // IPv4 DHCP server statistics.
    Statistics Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics
}

func (vrf *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "vrfs"
    vrf.EntityData.SegmentPath = "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = make(map[string]types.YChild)
    vrf.EntityData.Children["statistics"] = types.YChild{"Statistics", &vrf.Statistics}
    vrf.EntityData.Leafs = make(map[string]types.YLeaf)
    vrf.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", vrf.VrfName}
    return &(vrf.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics
// IPv4 DHCP server statistics
type Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics struct {
    EntityData types.CommonEntityData
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

func (statistics *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "vrf"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = make(map[string]types.YChild)
    statistics.EntityData.Children["discover"] = types.YChild{"Discover", &statistics.Discover}
    statistics.EntityData.Children["offer"] = types.YChild{"Offer", &statistics.Offer}
    statistics.EntityData.Children["request"] = types.YChild{"Request", &statistics.Request}
    statistics.EntityData.Children["decline"] = types.YChild{"Decline", &statistics.Decline}
    statistics.EntityData.Children["ack"] = types.YChild{"Ack", &statistics.Ack}
    statistics.EntityData.Children["nak"] = types.YChild{"Nak", &statistics.Nak}
    statistics.EntityData.Children["release"] = types.YChild{"Release", &statistics.Release}
    statistics.EntityData.Children["inform"] = types.YChild{"Inform", &statistics.Inform}
    statistics.EntityData.Children["lease-query"] = types.YChild{"LeaseQuery", &statistics.LeaseQuery}
    statistics.EntityData.Children["lease-not-assigned"] = types.YChild{"LeaseNotAssigned", &statistics.LeaseNotAssigned}
    statistics.EntityData.Children["lease-unknown"] = types.YChild{"LeaseUnknown", &statistics.LeaseUnknown}
    statistics.EntityData.Children["lease-active"] = types.YChild{"LeaseActive", &statistics.LeaseActive}
    statistics.EntityData.Children["bootp-request"] = types.YChild{"BootpRequest", &statistics.BootpRequest}
    statistics.EntityData.Children["bootp-reply"] = types.YChild{"BootpReply", &statistics.BootpReply}
    statistics.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(statistics.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Discover
// DHCP discover packets
type Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Discover struct {
    EntityData types.CommonEntityData
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

func (discover *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Discover) GetEntityData() *types.CommonEntityData {
    discover.EntityData.YFilter = discover.YFilter
    discover.EntityData.YangName = "discover"
    discover.EntityData.BundleName = "cisco_ios_xr"
    discover.EntityData.ParentYangName = "statistics"
    discover.EntityData.SegmentPath = "discover"
    discover.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    discover.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    discover.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    discover.EntityData.Children = make(map[string]types.YChild)
    discover.EntityData.Leafs = make(map[string]types.YLeaf)
    discover.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", discover.ReceivedPackets}
    discover.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", discover.TransmittedPackets}
    discover.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", discover.DroppedPackets}
    return &(discover.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Offer
// DHCP offer packets
type Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Offer struct {
    EntityData types.CommonEntityData
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

func (offer *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Offer) GetEntityData() *types.CommonEntityData {
    offer.EntityData.YFilter = offer.YFilter
    offer.EntityData.YangName = "offer"
    offer.EntityData.BundleName = "cisco_ios_xr"
    offer.EntityData.ParentYangName = "statistics"
    offer.EntityData.SegmentPath = "offer"
    offer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    offer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    offer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    offer.EntityData.Children = make(map[string]types.YChild)
    offer.EntityData.Leafs = make(map[string]types.YLeaf)
    offer.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", offer.ReceivedPackets}
    offer.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", offer.TransmittedPackets}
    offer.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", offer.DroppedPackets}
    return &(offer.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Request
// DHCP request packets
type Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Request struct {
    EntityData types.CommonEntityData
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

func (request *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Request) GetEntityData() *types.CommonEntityData {
    request.EntityData.YFilter = request.YFilter
    request.EntityData.YangName = "request"
    request.EntityData.BundleName = "cisco_ios_xr"
    request.EntityData.ParentYangName = "statistics"
    request.EntityData.SegmentPath = "request"
    request.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    request.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    request.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    request.EntityData.Children = make(map[string]types.YChild)
    request.EntityData.Leafs = make(map[string]types.YLeaf)
    request.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", request.ReceivedPackets}
    request.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", request.TransmittedPackets}
    request.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", request.DroppedPackets}
    return &(request.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Decline
// DHCP decline packets
type Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Decline struct {
    EntityData types.CommonEntityData
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

func (decline *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Decline) GetEntityData() *types.CommonEntityData {
    decline.EntityData.YFilter = decline.YFilter
    decline.EntityData.YangName = "decline"
    decline.EntityData.BundleName = "cisco_ios_xr"
    decline.EntityData.ParentYangName = "statistics"
    decline.EntityData.SegmentPath = "decline"
    decline.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    decline.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    decline.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    decline.EntityData.Children = make(map[string]types.YChild)
    decline.EntityData.Leafs = make(map[string]types.YLeaf)
    decline.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", decline.ReceivedPackets}
    decline.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", decline.TransmittedPackets}
    decline.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", decline.DroppedPackets}
    return &(decline.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Ack
// DHCP ack packets
type Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Ack struct {
    EntityData types.CommonEntityData
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

func (ack *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Ack) GetEntityData() *types.CommonEntityData {
    ack.EntityData.YFilter = ack.YFilter
    ack.EntityData.YangName = "ack"
    ack.EntityData.BundleName = "cisco_ios_xr"
    ack.EntityData.ParentYangName = "statistics"
    ack.EntityData.SegmentPath = "ack"
    ack.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ack.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ack.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ack.EntityData.Children = make(map[string]types.YChild)
    ack.EntityData.Leafs = make(map[string]types.YLeaf)
    ack.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", ack.ReceivedPackets}
    ack.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", ack.TransmittedPackets}
    ack.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", ack.DroppedPackets}
    return &(ack.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Nak
// DHCP nak packets
type Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Nak struct {
    EntityData types.CommonEntityData
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

func (nak *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Nak) GetEntityData() *types.CommonEntityData {
    nak.EntityData.YFilter = nak.YFilter
    nak.EntityData.YangName = "nak"
    nak.EntityData.BundleName = "cisco_ios_xr"
    nak.EntityData.ParentYangName = "statistics"
    nak.EntityData.SegmentPath = "nak"
    nak.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nak.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nak.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nak.EntityData.Children = make(map[string]types.YChild)
    nak.EntityData.Leafs = make(map[string]types.YLeaf)
    nak.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", nak.ReceivedPackets}
    nak.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", nak.TransmittedPackets}
    nak.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", nak.DroppedPackets}
    return &(nak.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Release
// DHCP release packets
type Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Release struct {
    EntityData types.CommonEntityData
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

func (release *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Release) GetEntityData() *types.CommonEntityData {
    release.EntityData.YFilter = release.YFilter
    release.EntityData.YangName = "release"
    release.EntityData.BundleName = "cisco_ios_xr"
    release.EntityData.ParentYangName = "statistics"
    release.EntityData.SegmentPath = "release"
    release.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    release.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    release.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    release.EntityData.Children = make(map[string]types.YChild)
    release.EntityData.Leafs = make(map[string]types.YLeaf)
    release.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", release.ReceivedPackets}
    release.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", release.TransmittedPackets}
    release.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", release.DroppedPackets}
    return &(release.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Inform
// DHCP inform packets
type Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Inform struct {
    EntityData types.CommonEntityData
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

func (inform *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_Inform) GetEntityData() *types.CommonEntityData {
    inform.EntityData.YFilter = inform.YFilter
    inform.EntityData.YangName = "inform"
    inform.EntityData.BundleName = "cisco_ios_xr"
    inform.EntityData.ParentYangName = "statistics"
    inform.EntityData.SegmentPath = "inform"
    inform.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inform.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inform.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inform.EntityData.Children = make(map[string]types.YChild)
    inform.EntityData.Leafs = make(map[string]types.YLeaf)
    inform.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", inform.ReceivedPackets}
    inform.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", inform.TransmittedPackets}
    inform.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", inform.DroppedPackets}
    return &(inform.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQuery
// DHCP lease query packets
type Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQuery struct {
    EntityData types.CommonEntityData
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

func (leaseQuery *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQuery) GetEntityData() *types.CommonEntityData {
    leaseQuery.EntityData.YFilter = leaseQuery.YFilter
    leaseQuery.EntityData.YangName = "lease-query"
    leaseQuery.EntityData.BundleName = "cisco_ios_xr"
    leaseQuery.EntityData.ParentYangName = "statistics"
    leaseQuery.EntityData.SegmentPath = "lease-query"
    leaseQuery.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    leaseQuery.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    leaseQuery.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    leaseQuery.EntityData.Children = make(map[string]types.YChild)
    leaseQuery.EntityData.Leafs = make(map[string]types.YLeaf)
    leaseQuery.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", leaseQuery.ReceivedPackets}
    leaseQuery.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", leaseQuery.TransmittedPackets}
    leaseQuery.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", leaseQuery.DroppedPackets}
    return &(leaseQuery.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseNotAssigned
// DHCP lease not assigned
type Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseNotAssigned struct {
    EntityData types.CommonEntityData
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

func (leaseNotAssigned *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseNotAssigned) GetEntityData() *types.CommonEntityData {
    leaseNotAssigned.EntityData.YFilter = leaseNotAssigned.YFilter
    leaseNotAssigned.EntityData.YangName = "lease-not-assigned"
    leaseNotAssigned.EntityData.BundleName = "cisco_ios_xr"
    leaseNotAssigned.EntityData.ParentYangName = "statistics"
    leaseNotAssigned.EntityData.SegmentPath = "lease-not-assigned"
    leaseNotAssigned.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    leaseNotAssigned.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    leaseNotAssigned.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    leaseNotAssigned.EntityData.Children = make(map[string]types.YChild)
    leaseNotAssigned.EntityData.Leafs = make(map[string]types.YLeaf)
    leaseNotAssigned.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", leaseNotAssigned.ReceivedPackets}
    leaseNotAssigned.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", leaseNotAssigned.TransmittedPackets}
    leaseNotAssigned.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", leaseNotAssigned.DroppedPackets}
    return &(leaseNotAssigned.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseUnknown
// DHCP lease unknown
type Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseUnknown struct {
    EntityData types.CommonEntityData
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

func (leaseUnknown *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseUnknown) GetEntityData() *types.CommonEntityData {
    leaseUnknown.EntityData.YFilter = leaseUnknown.YFilter
    leaseUnknown.EntityData.YangName = "lease-unknown"
    leaseUnknown.EntityData.BundleName = "cisco_ios_xr"
    leaseUnknown.EntityData.ParentYangName = "statistics"
    leaseUnknown.EntityData.SegmentPath = "lease-unknown"
    leaseUnknown.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    leaseUnknown.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    leaseUnknown.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    leaseUnknown.EntityData.Children = make(map[string]types.YChild)
    leaseUnknown.EntityData.Leafs = make(map[string]types.YLeaf)
    leaseUnknown.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", leaseUnknown.ReceivedPackets}
    leaseUnknown.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", leaseUnknown.TransmittedPackets}
    leaseUnknown.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", leaseUnknown.DroppedPackets}
    return &(leaseUnknown.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseActive
// DHCP lease active
type Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseActive struct {
    EntityData types.CommonEntityData
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

func (leaseActive *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseActive) GetEntityData() *types.CommonEntityData {
    leaseActive.EntityData.YFilter = leaseActive.YFilter
    leaseActive.EntityData.YangName = "lease-active"
    leaseActive.EntityData.BundleName = "cisco_ios_xr"
    leaseActive.EntityData.ParentYangName = "statistics"
    leaseActive.EntityData.SegmentPath = "lease-active"
    leaseActive.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    leaseActive.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    leaseActive.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    leaseActive.EntityData.Children = make(map[string]types.YChild)
    leaseActive.EntityData.Leafs = make(map[string]types.YLeaf)
    leaseActive.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", leaseActive.ReceivedPackets}
    leaseActive.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", leaseActive.TransmittedPackets}
    leaseActive.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", leaseActive.DroppedPackets}
    return &(leaseActive.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_BootpRequest
// DHCP BOOTP request
type Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_BootpRequest struct {
    EntityData types.CommonEntityData
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

func (bootpRequest *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_BootpRequest) GetEntityData() *types.CommonEntityData {
    bootpRequest.EntityData.YFilter = bootpRequest.YFilter
    bootpRequest.EntityData.YangName = "bootp-request"
    bootpRequest.EntityData.BundleName = "cisco_ios_xr"
    bootpRequest.EntityData.ParentYangName = "statistics"
    bootpRequest.EntityData.SegmentPath = "bootp-request"
    bootpRequest.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bootpRequest.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bootpRequest.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bootpRequest.EntityData.Children = make(map[string]types.YChild)
    bootpRequest.EntityData.Leafs = make(map[string]types.YLeaf)
    bootpRequest.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", bootpRequest.ReceivedPackets}
    bootpRequest.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", bootpRequest.TransmittedPackets}
    bootpRequest.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", bootpRequest.DroppedPackets}
    return &(bootpRequest.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_BootpReply
// DHCP BOOTP reply
type Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_BootpReply struct {
    EntityData types.CommonEntityData
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

func (bootpReply *Ipv4Dhcpd_Nodes_Node_Server_Vrfs_Vrf_Statistics_BootpReply) GetEntityData() *types.CommonEntityData {
    bootpReply.EntityData.YFilter = bootpReply.YFilter
    bootpReply.EntityData.YangName = "bootp-reply"
    bootpReply.EntityData.BundleName = "cisco_ios_xr"
    bootpReply.EntityData.ParentYangName = "statistics"
    bootpReply.EntityData.SegmentPath = "bootp-reply"
    bootpReply.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bootpReply.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bootpReply.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bootpReply.EntityData.Children = make(map[string]types.YChild)
    bootpReply.EntityData.Leafs = make(map[string]types.YLeaf)
    bootpReply.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", bootpReply.ReceivedPackets}
    bootpReply.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", bootpReply.TransmittedPackets}
    bootpReply.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", bootpReply.DroppedPackets}
    return &(bootpReply.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Relay
// IPv4 DHCPD Relay operational data
type Ipv4Dhcpd_Nodes_Node_Relay struct {
    EntityData types.CommonEntityData
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

func (relay *Ipv4Dhcpd_Nodes_Node_Relay) GetEntityData() *types.CommonEntityData {
    relay.EntityData.YFilter = relay.YFilter
    relay.EntityData.YangName = "relay"
    relay.EntityData.BundleName = "cisco_ios_xr"
    relay.EntityData.ParentYangName = "node"
    relay.EntityData.SegmentPath = "relay"
    relay.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    relay.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    relay.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    relay.EntityData.Children = make(map[string]types.YChild)
    relay.EntityData.Children["profiles"] = types.YChild{"Profiles", &relay.Profiles}
    relay.EntityData.Children["statistics-info"] = types.YChild{"StatisticsInfo", &relay.StatisticsInfo}
    relay.EntityData.Children["statistics"] = types.YChild{"Statistics", &relay.Statistics}
    relay.EntityData.Children["vrfs"] = types.YChild{"Vrfs", &relay.Vrfs}
    relay.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(relay.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Relay_Profiles
// DHCP Relay Profiles
type Ipv4Dhcpd_Nodes_Node_Relay_Profiles struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DHCP Relay profile. The type is slice of
    // Ipv4Dhcpd_Nodes_Node_Relay_Profiles_Profile.
    Profile []Ipv4Dhcpd_Nodes_Node_Relay_Profiles_Profile
}

func (profiles *Ipv4Dhcpd_Nodes_Node_Relay_Profiles) GetEntityData() *types.CommonEntityData {
    profiles.EntityData.YFilter = profiles.YFilter
    profiles.EntityData.YangName = "profiles"
    profiles.EntityData.BundleName = "cisco_ios_xr"
    profiles.EntityData.ParentYangName = "relay"
    profiles.EntityData.SegmentPath = "profiles"
    profiles.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    profiles.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    profiles.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    profiles.EntityData.Children = make(map[string]types.YChild)
    profiles.EntityData.Children["profile"] = types.YChild{"Profile", nil}
    for i := range profiles.Profile {
        profiles.EntityData.Children[types.GetSegmentPath(&profiles.Profile[i])] = types.YChild{"Profile", &profiles.Profile[i]}
    }
    profiles.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(profiles.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Relay_Profiles_Profile
// DHCP Relay profile
type Ipv4Dhcpd_Nodes_Node_Relay_Profiles_Profile struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Profile name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RelayProfileHelperAddress []interface{}

    // Helper address vrfs. The type is slice of string with length: 0..33.
    RelayProfileHelperVrf []interface{}

    // Gateway addresses. The type is slice of string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RelayProfileGiAddr []interface{}
}

func (profile *Ipv4Dhcpd_Nodes_Node_Relay_Profiles_Profile) GetEntityData() *types.CommonEntityData {
    profile.EntityData.YFilter = profile.YFilter
    profile.EntityData.YangName = "profile"
    profile.EntityData.BundleName = "cisco_ios_xr"
    profile.EntityData.ParentYangName = "profiles"
    profile.EntityData.SegmentPath = "profile" + "[profile-name='" + fmt.Sprintf("%v", profile.ProfileName) + "']"
    profile.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    profile.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    profile.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    profile.EntityData.Children = make(map[string]types.YChild)
    profile.EntityData.Leafs = make(map[string]types.YLeaf)
    profile.EntityData.Leafs["profile-name"] = types.YLeaf{"ProfileName", profile.ProfileName}
    profile.EntityData.Leafs["relay-profile-name"] = types.YLeaf{"RelayProfileName", profile.RelayProfileName}
    profile.EntityData.Leafs["relay-profile-uid"] = types.YLeaf{"RelayProfileUid", profile.RelayProfileUid}
    profile.EntityData.Leafs["relay-profile-helper-count"] = types.YLeaf{"RelayProfileHelperCount", profile.RelayProfileHelperCount}
    profile.EntityData.Leafs["relay-profile-relay-info-option"] = types.YLeaf{"RelayProfileRelayInfoOption", profile.RelayProfileRelayInfoOption}
    profile.EntityData.Leafs["relay-profile-relay-info-policy"] = types.YLeaf{"RelayProfileRelayInfoPolicy", profile.RelayProfileRelayInfoPolicy}
    profile.EntityData.Leafs["relay-profile-relay-info-allow-untrusted"] = types.YLeaf{"RelayProfileRelayInfoAllowUntrusted", profile.RelayProfileRelayInfoAllowUntrusted}
    profile.EntityData.Leafs["relay-profile-relay-info-optionvpn"] = types.YLeaf{"RelayProfileRelayInfoOptionvpn", profile.RelayProfileRelayInfoOptionvpn}
    profile.EntityData.Leafs["relay-profile-relay-info-optionvpn-mode"] = types.YLeaf{"RelayProfileRelayInfoOptionvpnMode", profile.RelayProfileRelayInfoOptionvpnMode}
    profile.EntityData.Leafs["relay-profile-relay-info-check"] = types.YLeaf{"RelayProfileRelayInfoCheck", profile.RelayProfileRelayInfoCheck}
    profile.EntityData.Leafs["relay-profile-gi-addr-policy"] = types.YLeaf{"RelayProfileGiAddrPolicy", profile.RelayProfileGiAddrPolicy}
    profile.EntityData.Leafs["relay-profile-broadcast-flag-policy"] = types.YLeaf{"RelayProfileBroadcastFlagPolicy", profile.RelayProfileBroadcastFlagPolicy}
    profile.EntityData.Leafs["relay-profile-mac-mismatch-action"] = types.YLeaf{"RelayProfileMacMismatchAction", profile.RelayProfileMacMismatchAction}
    profile.EntityData.Leafs["relay-profile-helper-address"] = types.YLeaf{"RelayProfileHelperAddress", profile.RelayProfileHelperAddress}
    profile.EntityData.Leafs["relay-profile-helper-vrf"] = types.YLeaf{"RelayProfileHelperVrf", profile.RelayProfileHelperVrf}
    profile.EntityData.Leafs["relay-profile-gi-addr"] = types.YLeaf{"RelayProfileGiAddr", profile.RelayProfileGiAddr}
    return &(profile.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Relay_StatisticsInfo
// DHCP relay statistics info
type Ipv4Dhcpd_Nodes_Node_Relay_StatisticsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Relay Stats timestamp. The type is interface{} with range: 0..4294967295.
    RelayStatsTimestamp interface{}
}

func (statisticsInfo *Ipv4Dhcpd_Nodes_Node_Relay_StatisticsInfo) GetEntityData() *types.CommonEntityData {
    statisticsInfo.EntityData.YFilter = statisticsInfo.YFilter
    statisticsInfo.EntityData.YangName = "statistics-info"
    statisticsInfo.EntityData.BundleName = "cisco_ios_xr"
    statisticsInfo.EntityData.ParentYangName = "relay"
    statisticsInfo.EntityData.SegmentPath = "statistics-info"
    statisticsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statisticsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statisticsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statisticsInfo.EntityData.Children = make(map[string]types.YChild)
    statisticsInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    statisticsInfo.EntityData.Leafs["relay-stats-timestamp"] = types.YLeaf{"RelayStatsTimestamp", statisticsInfo.RelayStatsTimestamp}
    return &(statisticsInfo.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Relay_Statistics
// DHCP Relay VRF statistics
type Ipv4Dhcpd_Nodes_Node_Relay_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ipv4 dhcpd relay stat. The type is slice of
    // Ipv4Dhcpd_Nodes_Node_Relay_Statistics_Ipv4DhcpdRelayStat.
    Ipv4DhcpdRelayStat []Ipv4Dhcpd_Nodes_Node_Relay_Statistics_Ipv4DhcpdRelayStat
}

func (statistics *Ipv4Dhcpd_Nodes_Node_Relay_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "relay"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = make(map[string]types.YChild)
    statistics.EntityData.Children["ipv4-dhcpd-relay-stat"] = types.YChild{"Ipv4DhcpdRelayStat", nil}
    for i := range statistics.Ipv4DhcpdRelayStat {
        statistics.EntityData.Children[types.GetSegmentPath(&statistics.Ipv4DhcpdRelayStat[i])] = types.YChild{"Ipv4DhcpdRelayStat", &statistics.Ipv4DhcpdRelayStat[i]}
    }
    statistics.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(statistics.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Relay_Statistics_Ipv4DhcpdRelayStat
// ipv4 dhcpd relay stat
type Ipv4Dhcpd_Nodes_Node_Relay_Statistics_Ipv4DhcpdRelayStat struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DHCP L3 VRF Name. The type is string with length: 0..33.
    RelayStatisticsVrfName interface{}

    // Public relay statistics.
    Statistics Ipv4Dhcpd_Nodes_Node_Relay_Statistics_Ipv4DhcpdRelayStat_Statistics_
}

func (ipv4DhcpdRelayStat *Ipv4Dhcpd_Nodes_Node_Relay_Statistics_Ipv4DhcpdRelayStat) GetEntityData() *types.CommonEntityData {
    ipv4DhcpdRelayStat.EntityData.YFilter = ipv4DhcpdRelayStat.YFilter
    ipv4DhcpdRelayStat.EntityData.YangName = "ipv4-dhcpd-relay-stat"
    ipv4DhcpdRelayStat.EntityData.BundleName = "cisco_ios_xr"
    ipv4DhcpdRelayStat.EntityData.ParentYangName = "statistics"
    ipv4DhcpdRelayStat.EntityData.SegmentPath = "ipv4-dhcpd-relay-stat"
    ipv4DhcpdRelayStat.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4DhcpdRelayStat.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4DhcpdRelayStat.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4DhcpdRelayStat.EntityData.Children = make(map[string]types.YChild)
    ipv4DhcpdRelayStat.EntityData.Children["statistics"] = types.YChild{"Statistics", &ipv4DhcpdRelayStat.Statistics}
    ipv4DhcpdRelayStat.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv4DhcpdRelayStat.EntityData.Leafs["relay-statistics-vrf-name"] = types.YLeaf{"RelayStatisticsVrfName", ipv4DhcpdRelayStat.RelayStatisticsVrfName}
    return &(ipv4DhcpdRelayStat.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Relay_Statistics_Ipv4DhcpdRelayStat_Statistics_
// Public relay statistics
type Ipv4Dhcpd_Nodes_Node_Relay_Statistics_Ipv4DhcpdRelayStat_Statistics_ struct {
    EntityData types.CommonEntityData
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

func (statistics_ *Ipv4Dhcpd_Nodes_Node_Relay_Statistics_Ipv4DhcpdRelayStat_Statistics_) GetEntityData() *types.CommonEntityData {
    statistics_.EntityData.YFilter = statistics_.YFilter
    statistics_.EntityData.YangName = "statistics"
    statistics_.EntityData.BundleName = "cisco_ios_xr"
    statistics_.EntityData.ParentYangName = "ipv4-dhcpd-relay-stat"
    statistics_.EntityData.SegmentPath = "statistics"
    statistics_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics_.EntityData.Children = make(map[string]types.YChild)
    statistics_.EntityData.Leafs = make(map[string]types.YLeaf)
    statistics_.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", statistics_.ReceivedPackets}
    statistics_.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", statistics_.TransmittedPackets}
    statistics_.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", statistics_.DroppedPackets}
    return &(statistics_.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Relay_Vrfs
// DHCP relay list of VRF names
type Ipv4Dhcpd_Nodes_Node_Relay_Vrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 DHCP relay VRF name. The type is slice of
    // Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf.
    Vrf []Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf
}

func (vrfs *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs) GetEntityData() *types.CommonEntityData {
    vrfs.EntityData.YFilter = vrfs.YFilter
    vrfs.EntityData.YangName = "vrfs"
    vrfs.EntityData.BundleName = "cisco_ios_xr"
    vrfs.EntityData.ParentYangName = "relay"
    vrfs.EntityData.SegmentPath = "vrfs"
    vrfs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfs.EntityData.Children = make(map[string]types.YChild)
    vrfs.EntityData.Children["vrf"] = types.YChild{"Vrf", nil}
    for i := range vrfs.Vrf {
        vrfs.EntityData.Children[types.GetSegmentPath(&vrfs.Vrf[i])] = types.YChild{"Vrf", &vrfs.Vrf[i]}
    }
    vrfs.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(vrfs.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf
// IPv4 DHCP relay VRF name
type Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    VrfName interface{}

    // IPv4 DHCP relay statistics.
    VrfStatistics Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics
}

func (vrf *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "vrfs"
    vrf.EntityData.SegmentPath = "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = make(map[string]types.YChild)
    vrf.EntityData.Children["vrf-statistics"] = types.YChild{"VrfStatistics", &vrf.VrfStatistics}
    vrf.EntityData.Leafs = make(map[string]types.YLeaf)
    vrf.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", vrf.VrfName}
    return &(vrf.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics
// IPv4 DHCP relay statistics
type Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics struct {
    EntityData types.CommonEntityData
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

func (vrfStatistics *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics) GetEntityData() *types.CommonEntityData {
    vrfStatistics.EntityData.YFilter = vrfStatistics.YFilter
    vrfStatistics.EntityData.YangName = "vrf-statistics"
    vrfStatistics.EntityData.BundleName = "cisco_ios_xr"
    vrfStatistics.EntityData.ParentYangName = "vrf"
    vrfStatistics.EntityData.SegmentPath = "vrf-statistics"
    vrfStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfStatistics.EntityData.Children = make(map[string]types.YChild)
    vrfStatistics.EntityData.Children["discover"] = types.YChild{"Discover", &vrfStatistics.Discover}
    vrfStatistics.EntityData.Children["offer"] = types.YChild{"Offer", &vrfStatistics.Offer}
    vrfStatistics.EntityData.Children["request"] = types.YChild{"Request", &vrfStatistics.Request}
    vrfStatistics.EntityData.Children["decline"] = types.YChild{"Decline", &vrfStatistics.Decline}
    vrfStatistics.EntityData.Children["ack"] = types.YChild{"Ack", &vrfStatistics.Ack}
    vrfStatistics.EntityData.Children["nak"] = types.YChild{"Nak", &vrfStatistics.Nak}
    vrfStatistics.EntityData.Children["release"] = types.YChild{"Release", &vrfStatistics.Release}
    vrfStatistics.EntityData.Children["inform"] = types.YChild{"Inform", &vrfStatistics.Inform}
    vrfStatistics.EntityData.Children["lease-query"] = types.YChild{"LeaseQuery", &vrfStatistics.LeaseQuery}
    vrfStatistics.EntityData.Children["lease-not-assigned"] = types.YChild{"LeaseNotAssigned", &vrfStatistics.LeaseNotAssigned}
    vrfStatistics.EntityData.Children["lease-unknown"] = types.YChild{"LeaseUnknown", &vrfStatistics.LeaseUnknown}
    vrfStatistics.EntityData.Children["lease-active"] = types.YChild{"LeaseActive", &vrfStatistics.LeaseActive}
    vrfStatistics.EntityData.Children["bootp-request"] = types.YChild{"BootpRequest", &vrfStatistics.BootpRequest}
    vrfStatistics.EntityData.Children["bootp-reply"] = types.YChild{"BootpReply", &vrfStatistics.BootpReply}
    vrfStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(vrfStatistics.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Discover
// DHCP discover packets
type Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Discover struct {
    EntityData types.CommonEntityData
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

func (discover *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Discover) GetEntityData() *types.CommonEntityData {
    discover.EntityData.YFilter = discover.YFilter
    discover.EntityData.YangName = "discover"
    discover.EntityData.BundleName = "cisco_ios_xr"
    discover.EntityData.ParentYangName = "vrf-statistics"
    discover.EntityData.SegmentPath = "discover"
    discover.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    discover.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    discover.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    discover.EntityData.Children = make(map[string]types.YChild)
    discover.EntityData.Leafs = make(map[string]types.YLeaf)
    discover.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", discover.ReceivedPackets}
    discover.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", discover.TransmittedPackets}
    discover.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", discover.DroppedPackets}
    return &(discover.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Offer
// DHCP offer packets
type Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Offer struct {
    EntityData types.CommonEntityData
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

func (offer *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Offer) GetEntityData() *types.CommonEntityData {
    offer.EntityData.YFilter = offer.YFilter
    offer.EntityData.YangName = "offer"
    offer.EntityData.BundleName = "cisco_ios_xr"
    offer.EntityData.ParentYangName = "vrf-statistics"
    offer.EntityData.SegmentPath = "offer"
    offer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    offer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    offer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    offer.EntityData.Children = make(map[string]types.YChild)
    offer.EntityData.Leafs = make(map[string]types.YLeaf)
    offer.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", offer.ReceivedPackets}
    offer.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", offer.TransmittedPackets}
    offer.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", offer.DroppedPackets}
    return &(offer.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Request
// DHCP request packets
type Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Request struct {
    EntityData types.CommonEntityData
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

func (request *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Request) GetEntityData() *types.CommonEntityData {
    request.EntityData.YFilter = request.YFilter
    request.EntityData.YangName = "request"
    request.EntityData.BundleName = "cisco_ios_xr"
    request.EntityData.ParentYangName = "vrf-statistics"
    request.EntityData.SegmentPath = "request"
    request.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    request.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    request.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    request.EntityData.Children = make(map[string]types.YChild)
    request.EntityData.Leafs = make(map[string]types.YLeaf)
    request.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", request.ReceivedPackets}
    request.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", request.TransmittedPackets}
    request.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", request.DroppedPackets}
    return &(request.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Decline
// DHCP decline packets
type Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Decline struct {
    EntityData types.CommonEntityData
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

func (decline *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Decline) GetEntityData() *types.CommonEntityData {
    decline.EntityData.YFilter = decline.YFilter
    decline.EntityData.YangName = "decline"
    decline.EntityData.BundleName = "cisco_ios_xr"
    decline.EntityData.ParentYangName = "vrf-statistics"
    decline.EntityData.SegmentPath = "decline"
    decline.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    decline.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    decline.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    decline.EntityData.Children = make(map[string]types.YChild)
    decline.EntityData.Leafs = make(map[string]types.YLeaf)
    decline.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", decline.ReceivedPackets}
    decline.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", decline.TransmittedPackets}
    decline.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", decline.DroppedPackets}
    return &(decline.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Ack
// DHCP ack packets
type Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Ack struct {
    EntityData types.CommonEntityData
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

func (ack *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Ack) GetEntityData() *types.CommonEntityData {
    ack.EntityData.YFilter = ack.YFilter
    ack.EntityData.YangName = "ack"
    ack.EntityData.BundleName = "cisco_ios_xr"
    ack.EntityData.ParentYangName = "vrf-statistics"
    ack.EntityData.SegmentPath = "ack"
    ack.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ack.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ack.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ack.EntityData.Children = make(map[string]types.YChild)
    ack.EntityData.Leafs = make(map[string]types.YLeaf)
    ack.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", ack.ReceivedPackets}
    ack.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", ack.TransmittedPackets}
    ack.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", ack.DroppedPackets}
    return &(ack.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Nak
// DHCP nak packets
type Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Nak struct {
    EntityData types.CommonEntityData
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

func (nak *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Nak) GetEntityData() *types.CommonEntityData {
    nak.EntityData.YFilter = nak.YFilter
    nak.EntityData.YangName = "nak"
    nak.EntityData.BundleName = "cisco_ios_xr"
    nak.EntityData.ParentYangName = "vrf-statistics"
    nak.EntityData.SegmentPath = "nak"
    nak.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nak.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nak.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nak.EntityData.Children = make(map[string]types.YChild)
    nak.EntityData.Leafs = make(map[string]types.YLeaf)
    nak.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", nak.ReceivedPackets}
    nak.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", nak.TransmittedPackets}
    nak.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", nak.DroppedPackets}
    return &(nak.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Release
// DHCP release packets
type Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Release struct {
    EntityData types.CommonEntityData
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

func (release *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Release) GetEntityData() *types.CommonEntityData {
    release.EntityData.YFilter = release.YFilter
    release.EntityData.YangName = "release"
    release.EntityData.BundleName = "cisco_ios_xr"
    release.EntityData.ParentYangName = "vrf-statistics"
    release.EntityData.SegmentPath = "release"
    release.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    release.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    release.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    release.EntityData.Children = make(map[string]types.YChild)
    release.EntityData.Leafs = make(map[string]types.YLeaf)
    release.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", release.ReceivedPackets}
    release.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", release.TransmittedPackets}
    release.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", release.DroppedPackets}
    return &(release.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Inform
// DHCP inform packets
type Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Inform struct {
    EntityData types.CommonEntityData
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

func (inform *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_Inform) GetEntityData() *types.CommonEntityData {
    inform.EntityData.YFilter = inform.YFilter
    inform.EntityData.YangName = "inform"
    inform.EntityData.BundleName = "cisco_ios_xr"
    inform.EntityData.ParentYangName = "vrf-statistics"
    inform.EntityData.SegmentPath = "inform"
    inform.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inform.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inform.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inform.EntityData.Children = make(map[string]types.YChild)
    inform.EntityData.Leafs = make(map[string]types.YLeaf)
    inform.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", inform.ReceivedPackets}
    inform.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", inform.TransmittedPackets}
    inform.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", inform.DroppedPackets}
    return &(inform.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseQuery
// DHCP lease query packets
type Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseQuery struct {
    EntityData types.CommonEntityData
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

func (leaseQuery *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseQuery) GetEntityData() *types.CommonEntityData {
    leaseQuery.EntityData.YFilter = leaseQuery.YFilter
    leaseQuery.EntityData.YangName = "lease-query"
    leaseQuery.EntityData.BundleName = "cisco_ios_xr"
    leaseQuery.EntityData.ParentYangName = "vrf-statistics"
    leaseQuery.EntityData.SegmentPath = "lease-query"
    leaseQuery.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    leaseQuery.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    leaseQuery.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    leaseQuery.EntityData.Children = make(map[string]types.YChild)
    leaseQuery.EntityData.Leafs = make(map[string]types.YLeaf)
    leaseQuery.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", leaseQuery.ReceivedPackets}
    leaseQuery.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", leaseQuery.TransmittedPackets}
    leaseQuery.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", leaseQuery.DroppedPackets}
    return &(leaseQuery.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseNotAssigned
// DHCP lease not assigned
type Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseNotAssigned struct {
    EntityData types.CommonEntityData
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

func (leaseNotAssigned *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseNotAssigned) GetEntityData() *types.CommonEntityData {
    leaseNotAssigned.EntityData.YFilter = leaseNotAssigned.YFilter
    leaseNotAssigned.EntityData.YangName = "lease-not-assigned"
    leaseNotAssigned.EntityData.BundleName = "cisco_ios_xr"
    leaseNotAssigned.EntityData.ParentYangName = "vrf-statistics"
    leaseNotAssigned.EntityData.SegmentPath = "lease-not-assigned"
    leaseNotAssigned.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    leaseNotAssigned.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    leaseNotAssigned.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    leaseNotAssigned.EntityData.Children = make(map[string]types.YChild)
    leaseNotAssigned.EntityData.Leafs = make(map[string]types.YLeaf)
    leaseNotAssigned.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", leaseNotAssigned.ReceivedPackets}
    leaseNotAssigned.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", leaseNotAssigned.TransmittedPackets}
    leaseNotAssigned.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", leaseNotAssigned.DroppedPackets}
    return &(leaseNotAssigned.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseUnknown
// DHCP lease unknown
type Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseUnknown struct {
    EntityData types.CommonEntityData
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

func (leaseUnknown *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseUnknown) GetEntityData() *types.CommonEntityData {
    leaseUnknown.EntityData.YFilter = leaseUnknown.YFilter
    leaseUnknown.EntityData.YangName = "lease-unknown"
    leaseUnknown.EntityData.BundleName = "cisco_ios_xr"
    leaseUnknown.EntityData.ParentYangName = "vrf-statistics"
    leaseUnknown.EntityData.SegmentPath = "lease-unknown"
    leaseUnknown.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    leaseUnknown.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    leaseUnknown.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    leaseUnknown.EntityData.Children = make(map[string]types.YChild)
    leaseUnknown.EntityData.Leafs = make(map[string]types.YLeaf)
    leaseUnknown.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", leaseUnknown.ReceivedPackets}
    leaseUnknown.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", leaseUnknown.TransmittedPackets}
    leaseUnknown.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", leaseUnknown.DroppedPackets}
    return &(leaseUnknown.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseActive
// DHCP lease active
type Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseActive struct {
    EntityData types.CommonEntityData
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

func (leaseActive *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_LeaseActive) GetEntityData() *types.CommonEntityData {
    leaseActive.EntityData.YFilter = leaseActive.YFilter
    leaseActive.EntityData.YangName = "lease-active"
    leaseActive.EntityData.BundleName = "cisco_ios_xr"
    leaseActive.EntityData.ParentYangName = "vrf-statistics"
    leaseActive.EntityData.SegmentPath = "lease-active"
    leaseActive.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    leaseActive.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    leaseActive.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    leaseActive.EntityData.Children = make(map[string]types.YChild)
    leaseActive.EntityData.Leafs = make(map[string]types.YLeaf)
    leaseActive.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", leaseActive.ReceivedPackets}
    leaseActive.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", leaseActive.TransmittedPackets}
    leaseActive.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", leaseActive.DroppedPackets}
    return &(leaseActive.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_BootpRequest
// DHCP BOOTP request
type Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_BootpRequest struct {
    EntityData types.CommonEntityData
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

func (bootpRequest *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_BootpRequest) GetEntityData() *types.CommonEntityData {
    bootpRequest.EntityData.YFilter = bootpRequest.YFilter
    bootpRequest.EntityData.YangName = "bootp-request"
    bootpRequest.EntityData.BundleName = "cisco_ios_xr"
    bootpRequest.EntityData.ParentYangName = "vrf-statistics"
    bootpRequest.EntityData.SegmentPath = "bootp-request"
    bootpRequest.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bootpRequest.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bootpRequest.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bootpRequest.EntityData.Children = make(map[string]types.YChild)
    bootpRequest.EntityData.Leafs = make(map[string]types.YLeaf)
    bootpRequest.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", bootpRequest.ReceivedPackets}
    bootpRequest.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", bootpRequest.TransmittedPackets}
    bootpRequest.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", bootpRequest.DroppedPackets}
    return &(bootpRequest.EntityData)
}

// Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_BootpReply
// DHCP BOOTP reply
type Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_BootpReply struct {
    EntityData types.CommonEntityData
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

func (bootpReply *Ipv4Dhcpd_Nodes_Node_Relay_Vrfs_Vrf_VrfStatistics_BootpReply) GetEntityData() *types.CommonEntityData {
    bootpReply.EntityData.YFilter = bootpReply.YFilter
    bootpReply.EntityData.YangName = "bootp-reply"
    bootpReply.EntityData.BundleName = "cisco_ios_xr"
    bootpReply.EntityData.ParentYangName = "vrf-statistics"
    bootpReply.EntityData.SegmentPath = "bootp-reply"
    bootpReply.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bootpReply.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bootpReply.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bootpReply.EntityData.Children = make(map[string]types.YChild)
    bootpReply.EntityData.Leafs = make(map[string]types.YLeaf)
    bootpReply.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", bootpReply.ReceivedPackets}
    bootpReply.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", bootpReply.TransmittedPackets}
    bootpReply.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", bootpReply.DroppedPackets}
    return &(bootpReply.EntityData)
}

