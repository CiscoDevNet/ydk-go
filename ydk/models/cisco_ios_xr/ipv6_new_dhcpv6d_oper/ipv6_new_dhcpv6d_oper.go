// This module contains a collection of YANG definitions
// for Cisco IOS-XR ipv6-new-dhcpv6d package operational data.
// 
// This module contains definitions
// for the following management objects:
//   dhcpv6: IPV6 DHCPD operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ipv6_new_dhcpv6d_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ipv6_new_dhcpv6d_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ipv6-new-dhcpv6d-oper dhcpv6}", reflect.TypeOf(Dhcpv6{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6", reflect.TypeOf(Dhcpv6{}))
}

// LeaseLimit represents Profile lease limit type
type LeaseLimit string

const (
    // Lease limit type none
    LeaseLimit_none LeaseLimit = "none"

    // Lease limit type interface
    LeaseLimit_interface_ LeaseLimit = "interface"

    // Lease limit type circuit ID
    LeaseLimit_circuit_id LeaseLimit = "circuit-id"

    // Lease limit type remote ID
    LeaseLimit_remote_id LeaseLimit = "remote-id"
)

// BagDhcpv6DSubMode represents Bag dhcpv6d sub mode
type BagDhcpv6DSubMode string

const (
    // DHCPv6 Base mode
    BagDhcpv6DSubMode_base BagDhcpv6DSubMode = "base"

    // DHCPv6 Server mode
    BagDhcpv6DSubMode_server BagDhcpv6DSubMode = "server"

    // DHCPv6 Proxy mode
    BagDhcpv6DSubMode_proxy BagDhcpv6DSubMode = "proxy"
)

// BagDhcpv6DFsmState represents Bag dhcpv6d fsm state
type BagDhcpv6DFsmState string

const (
    // Server initializing state for client binding
    BagDhcpv6DFsmState_server_initializing BagDhcpv6DFsmState = "server-initializing"

    // Server waiting on DPM to validate subscriber
    BagDhcpv6DFsmState_server_waiting_dpm BagDhcpv6DFsmState = "server-waiting-dpm"

    // Server waiting on DAPS to assign/free
    // addr/prefix
    BagDhcpv6DFsmState_server_waiting_daps BagDhcpv6DFsmState = "server-waiting-daps"

    // Server waiting for a request from the client
    BagDhcpv6DFsmState_server_waiting_client BagDhcpv6DFsmState = "server-waiting-client"

    // Server waiting for iedge response for the
    // session
    BagDhcpv6DFsmState_server_waiting_subscriber BagDhcpv6DFsmState = "server-waiting-subscriber"

    // Server waiting for RIB response for route add
    BagDhcpv6DFsmState_server_waiting_rib BagDhcpv6DFsmState = "server-waiting-rib"

    // Server bound state to the client
    BagDhcpv6DFsmState_server_bound_client BagDhcpv6DFsmState = "server-bound-client"

    // Proxy initializing state for client binding
    BagDhcpv6DFsmState_proxy_initializing BagDhcpv6DFsmState = "proxy-initializing"

    // Proxy waiting on DPM to validate subscriber
    BagDhcpv6DFsmState_proxy_waiting_dpm BagDhcpv6DFsmState = "proxy-waiting-dpm"

    // Proxy waiting on DAPS to assign/free prefix(ND)
    BagDhcpv6DFsmState_proxy_waiting_daps BagDhcpv6DFsmState = "proxy-waiting-daps"

    // Proxy waiting for a msg from the client/srv
    BagDhcpv6DFsmState_proxy_waiting_client_server BagDhcpv6DFsmState = "proxy-waiting-client-server"

    // Proxy waiting on iedge to sub session resp
    BagDhcpv6DFsmState_proxy_waiting_subscriber BagDhcpv6DFsmState = "proxy-waiting-subscriber"

    // Proxy waiting on RIB response
    BagDhcpv6DFsmState_proxy_waiting_rib BagDhcpv6DFsmState = "proxy-waiting-rib"

    // Proxy bound state to the client
    BagDhcpv6DFsmState_proxy_bound_client BagDhcpv6DFsmState = "proxy-bound-client"
)

// BagDhcpv6DIntfSrgRole represents Bag dhcpv6d intf srg role
type BagDhcpv6DIntfSrgRole string

const (
    // DHCPv6 Interface SRG role NONE
    BagDhcpv6DIntfSrgRole_none BagDhcpv6DIntfSrgRole = "none"

    // DHCPv6 Interface SRG role Master
    BagDhcpv6DIntfSrgRole_master BagDhcpv6DIntfSrgRole = "master"

    // DHCPv6 Interface SRG role Slave
    BagDhcpv6DIntfSrgRole_slave BagDhcpv6DIntfSrgRole = "slave"
)

// Dhcpv6IssuVersion represents Dhcpv6 issu version
type Dhcpv6IssuVersion string

const (
    // Version 1
    Dhcpv6IssuVersion_version1 Dhcpv6IssuVersion = "version1"

    // Version 2
    Dhcpv6IssuVersion_version2 Dhcpv6IssuVersion = "version2"
)

// BagDhcpv6DIaId represents Bag dhcpv6d ia id
type BagDhcpv6DIaId string

const (
    // Non-temporary Addresses assigned 
    BagDhcpv6DIaId_iana BagDhcpv6DIaId = "iana"

    // Prefix delegeated to client      
    BagDhcpv6DIaId_iapd BagDhcpv6DIaId = "iapd"

    // Temporary Addresses - not supported 
    BagDhcpv6DIaId_iata BagDhcpv6DIaId = "iata"
)

// Dhcpv6IssuRole represents Dhcpv6 issu role
type Dhcpv6IssuRole string

const (
    // Primary role
    Dhcpv6IssuRole_role_primary Dhcpv6IssuRole = "role-primary"

    // Secondary role
    Dhcpv6IssuRole_role_secondary Dhcpv6IssuRole = "role-secondary"
)

// BagDhcpv6DIntfSergRole represents Bag dhcpv6d intf serg role
type BagDhcpv6DIntfSergRole string

const (
    // DHCPv6 Interface SERG role NONE
    BagDhcpv6DIntfSergRole_none BagDhcpv6DIntfSergRole = "none"

    // DHCPv6 Interface SERG role Master
    BagDhcpv6DIntfSergRole_master BagDhcpv6DIntfSergRole = "master"

    // DHCPv6 Interface SERG role Slave
    BagDhcpv6DIntfSergRole_slave BagDhcpv6DIntfSergRole = "slave"
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

// Dhcpv6
// IPV6 DHCPD operational data
type Dhcpv6 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DHCP IssuStatus.
    IssuStatus Dhcpv6_IssuStatus

    // IPv6 DHCP list of nodes.
    Nodes Dhcpv6_Nodes
}

func (dhcpv6 *Dhcpv6) GetFilter() yfilter.YFilter { return dhcpv6.YFilter }

func (dhcpv6 *Dhcpv6) SetFilter(yf yfilter.YFilter) { dhcpv6.YFilter = yf }

func (dhcpv6 *Dhcpv6) GetGoName(yname string) string {
    if yname == "issu-status" { return "IssuStatus" }
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (dhcpv6 *Dhcpv6) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6"
}

func (dhcpv6 *Dhcpv6) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "issu-status" {
        return &dhcpv6.IssuStatus
    }
    if childYangName == "nodes" {
        return &dhcpv6.Nodes
    }
    return nil
}

func (dhcpv6 *Dhcpv6) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["issu-status"] = &dhcpv6.IssuStatus
    children["nodes"] = &dhcpv6.Nodes
    return children
}

func (dhcpv6 *Dhcpv6) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dhcpv6 *Dhcpv6) GetBundleName() string { return "cisco_ios_xr" }

func (dhcpv6 *Dhcpv6) GetYangName() string { return "dhcpv6" }

func (dhcpv6 *Dhcpv6) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dhcpv6 *Dhcpv6) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dhcpv6 *Dhcpv6) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dhcpv6 *Dhcpv6) SetParent(parent types.Entity) { dhcpv6.parent = parent }

func (dhcpv6 *Dhcpv6) GetParent() types.Entity { return dhcpv6.parent }

func (dhcpv6 *Dhcpv6) GetParentYangName() string { return "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper" }

// Dhcpv6_IssuStatus
// DHCP IssuStatus
type Dhcpv6_IssuStatus struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Timestamp for the process start time in nanoseconds since Epoch, i.e. since
    // 00:00:00 UTC , January 1, 1970. The type is interface{} with range:
    // 0..18446744073709551615. Units are nanosecond.
    ProcessStartTime interface{}

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

    // Whether or not DHCP is currently connected to ISSU Manager during the ISSU
    // Load Phase. The type is bool.
    IssuReadyIssuMgrConnection interface{}

    // The current role of the DHCP process. The type is Dhcpv6IssuRole.
    Role interface{}

    // The current ISSU phase of the DHCP process. The type is DhcpIssuPhase.
    Phase interface{}

    // The current version of the DHCP process in the context of an ISSU. The type
    // is Dhcpv6IssuVersion.
    Version interface{}
}

func (issuStatus *Dhcpv6_IssuStatus) GetFilter() yfilter.YFilter { return issuStatus.YFilter }

func (issuStatus *Dhcpv6_IssuStatus) SetFilter(yf yfilter.YFilter) { issuStatus.YFilter = yf }

func (issuStatus *Dhcpv6_IssuStatus) GetGoName(yname string) string {
    if yname == "process-start-time" { return "ProcessStartTime" }
    if yname == "issu-sync-complete-time" { return "IssuSyncCompleteTime" }
    if yname == "issu-sync-start-time" { return "IssuSyncStartTime" }
    if yname == "issu-ready-time" { return "IssuReadyTime" }
    if yname == "big-bang-time" { return "BigBangTime" }
    if yname == "primary-role-time" { return "PrimaryRoleTime" }
    if yname == "issu-ready-issu-mgr-connection" { return "IssuReadyIssuMgrConnection" }
    if yname == "role" { return "Role" }
    if yname == "phase" { return "Phase" }
    if yname == "version" { return "Version" }
    return ""
}

func (issuStatus *Dhcpv6_IssuStatus) GetSegmentPath() string {
    return "issu-status"
}

func (issuStatus *Dhcpv6_IssuStatus) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (issuStatus *Dhcpv6_IssuStatus) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (issuStatus *Dhcpv6_IssuStatus) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["process-start-time"] = issuStatus.ProcessStartTime
    leafs["issu-sync-complete-time"] = issuStatus.IssuSyncCompleteTime
    leafs["issu-sync-start-time"] = issuStatus.IssuSyncStartTime
    leafs["issu-ready-time"] = issuStatus.IssuReadyTime
    leafs["big-bang-time"] = issuStatus.BigBangTime
    leafs["primary-role-time"] = issuStatus.PrimaryRoleTime
    leafs["issu-ready-issu-mgr-connection"] = issuStatus.IssuReadyIssuMgrConnection
    leafs["role"] = issuStatus.Role
    leafs["phase"] = issuStatus.Phase
    leafs["version"] = issuStatus.Version
    return leafs
}

func (issuStatus *Dhcpv6_IssuStatus) GetBundleName() string { return "cisco_ios_xr" }

func (issuStatus *Dhcpv6_IssuStatus) GetYangName() string { return "issu-status" }

func (issuStatus *Dhcpv6_IssuStatus) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (issuStatus *Dhcpv6_IssuStatus) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (issuStatus *Dhcpv6_IssuStatus) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (issuStatus *Dhcpv6_IssuStatus) SetParent(parent types.Entity) { issuStatus.parent = parent }

func (issuStatus *Dhcpv6_IssuStatus) GetParent() types.Entity { return issuStatus.parent }

func (issuStatus *Dhcpv6_IssuStatus) GetParentYangName() string { return "dhcpv6" }

// Dhcpv6_Nodes
// IPv6 DHCP list of nodes
type Dhcpv6_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 DHCP particular node name. The type is slice of Dhcpv6_Nodes_Node.
    Node []Dhcpv6_Nodes_Node
}

func (nodes *Dhcpv6_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *Dhcpv6_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *Dhcpv6_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *Dhcpv6_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *Dhcpv6_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dhcpv6_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *Dhcpv6_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *Dhcpv6_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *Dhcpv6_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *Dhcpv6_Nodes) GetYangName() string { return "nodes" }

func (nodes *Dhcpv6_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *Dhcpv6_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *Dhcpv6_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *Dhcpv6_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *Dhcpv6_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *Dhcpv6_Nodes) GetParentYangName() string { return "dhcpv6" }

// Dhcpv6_Nodes_Node
// IPv6 DHCP particular node name
type Dhcpv6_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // IPv6 DHCP proxy operational data.
    Proxy Dhcpv6_Nodes_Node_Proxy

    // IPv6 DHCP Base.
    Base Dhcpv6_Nodes_Node_Base

    // IPv6 DHCP server operational data.
    Server Dhcpv6_Nodes_Node_Server

    // IPv6 DHCP relay operational data.
    Relay Dhcpv6_Nodes_Node_Relay
}

func (node *Dhcpv6_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *Dhcpv6_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *Dhcpv6_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "proxy" { return "Proxy" }
    if yname == "base" { return "Base" }
    if yname == "server" { return "Server" }
    if yname == "relay" { return "Relay" }
    return ""
}

func (node *Dhcpv6_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *Dhcpv6_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "proxy" {
        return &node.Proxy
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

func (node *Dhcpv6_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["proxy"] = &node.Proxy
    children["base"] = &node.Base
    children["server"] = &node.Server
    children["relay"] = &node.Relay
    return children
}

func (node *Dhcpv6_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    return leafs
}

func (node *Dhcpv6_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *Dhcpv6_Nodes_Node) GetYangName() string { return "node" }

func (node *Dhcpv6_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *Dhcpv6_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *Dhcpv6_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *Dhcpv6_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *Dhcpv6_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *Dhcpv6_Nodes_Node) GetParentYangName() string { return "nodes" }

// Dhcpv6_Nodes_Node_Proxy
// IPv6 DHCP proxy operational data
type Dhcpv6_Nodes_Node_Proxy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DHCPV6 proxy list of VRF names.
    Vrfs Dhcpv6_Nodes_Node_Proxy_Vrfs

    // IPv6 DHCP proxy profile.
    Profiles Dhcpv6_Nodes_Node_Proxy_Profiles

    // DHCPV6 proxy interface.
    Interfaces Dhcpv6_Nodes_Node_Proxy_Interfaces

    // DHCPv6 proxy statistics.
    Statistics Dhcpv6_Nodes_Node_Proxy_Statistics

    // DHCPV6 proxy bindings.
    Binding Dhcpv6_Nodes_Node_Proxy_Binding
}

func (proxy *Dhcpv6_Nodes_Node_Proxy) GetFilter() yfilter.YFilter { return proxy.YFilter }

func (proxy *Dhcpv6_Nodes_Node_Proxy) SetFilter(yf yfilter.YFilter) { proxy.YFilter = yf }

func (proxy *Dhcpv6_Nodes_Node_Proxy) GetGoName(yname string) string {
    if yname == "vrfs" { return "Vrfs" }
    if yname == "profiles" { return "Profiles" }
    if yname == "interfaces" { return "Interfaces" }
    if yname == "statistics" { return "Statistics" }
    if yname == "binding" { return "Binding" }
    return ""
}

func (proxy *Dhcpv6_Nodes_Node_Proxy) GetSegmentPath() string {
    return "proxy"
}

func (proxy *Dhcpv6_Nodes_Node_Proxy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrfs" {
        return &proxy.Vrfs
    }
    if childYangName == "profiles" {
        return &proxy.Profiles
    }
    if childYangName == "interfaces" {
        return &proxy.Interfaces
    }
    if childYangName == "statistics" {
        return &proxy.Statistics
    }
    if childYangName == "binding" {
        return &proxy.Binding
    }
    return nil
}

func (proxy *Dhcpv6_Nodes_Node_Proxy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["vrfs"] = &proxy.Vrfs
    children["profiles"] = &proxy.Profiles
    children["interfaces"] = &proxy.Interfaces
    children["statistics"] = &proxy.Statistics
    children["binding"] = &proxy.Binding
    return children
}

func (proxy *Dhcpv6_Nodes_Node_Proxy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (proxy *Dhcpv6_Nodes_Node_Proxy) GetBundleName() string { return "cisco_ios_xr" }

func (proxy *Dhcpv6_Nodes_Node_Proxy) GetYangName() string { return "proxy" }

func (proxy *Dhcpv6_Nodes_Node_Proxy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (proxy *Dhcpv6_Nodes_Node_Proxy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (proxy *Dhcpv6_Nodes_Node_Proxy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (proxy *Dhcpv6_Nodes_Node_Proxy) SetParent(parent types.Entity) { proxy.parent = parent }

func (proxy *Dhcpv6_Nodes_Node_Proxy) GetParent() types.Entity { return proxy.parent }

func (proxy *Dhcpv6_Nodes_Node_Proxy) GetParentYangName() string { return "node" }

// Dhcpv6_Nodes_Node_Proxy_Vrfs
// DHCPV6 proxy list of VRF names
type Dhcpv6_Nodes_Node_Proxy_Vrfs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 DHCP proxy VRF name. The type is slice of
    // Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf.
    Vrf []Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf
}

func (vrfs *Dhcpv6_Nodes_Node_Proxy_Vrfs) GetFilter() yfilter.YFilter { return vrfs.YFilter }

func (vrfs *Dhcpv6_Nodes_Node_Proxy_Vrfs) SetFilter(yf yfilter.YFilter) { vrfs.YFilter = yf }

func (vrfs *Dhcpv6_Nodes_Node_Proxy_Vrfs) GetGoName(yname string) string {
    if yname == "vrf" { return "Vrf" }
    return ""
}

func (vrfs *Dhcpv6_Nodes_Node_Proxy_Vrfs) GetSegmentPath() string {
    return "vrfs"
}

func (vrfs *Dhcpv6_Nodes_Node_Proxy_Vrfs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf" {
        for _, c := range vrfs.Vrf {
            if vrfs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf{}
        vrfs.Vrf = append(vrfs.Vrf, child)
        return &vrfs.Vrf[len(vrfs.Vrf)-1]
    }
    return nil
}

func (vrfs *Dhcpv6_Nodes_Node_Proxy_Vrfs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vrfs.Vrf {
        children[vrfs.Vrf[i].GetSegmentPath()] = &vrfs.Vrf[i]
    }
    return children
}

func (vrfs *Dhcpv6_Nodes_Node_Proxy_Vrfs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vrfs *Dhcpv6_Nodes_Node_Proxy_Vrfs) GetBundleName() string { return "cisco_ios_xr" }

func (vrfs *Dhcpv6_Nodes_Node_Proxy_Vrfs) GetYangName() string { return "vrfs" }

func (vrfs *Dhcpv6_Nodes_Node_Proxy_Vrfs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfs *Dhcpv6_Nodes_Node_Proxy_Vrfs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfs *Dhcpv6_Nodes_Node_Proxy_Vrfs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfs *Dhcpv6_Nodes_Node_Proxy_Vrfs) SetParent(parent types.Entity) { vrfs.parent = parent }

func (vrfs *Dhcpv6_Nodes_Node_Proxy_Vrfs) GetParent() types.Entity { return vrfs.parent }

func (vrfs *Dhcpv6_Nodes_Node_Proxy_Vrfs) GetParentYangName() string { return "proxy" }

// Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf
// IPv6 DHCP proxy VRF name
type Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // IPv6 DHCP proxy statistics.
    Statistics Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics
}

func (vrf *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf) GetFilter() yfilter.YFilter { return vrf.YFilter }

func (vrf *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf) SetFilter(yf yfilter.YFilter) { vrf.YFilter = yf }

func (vrf *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "statistics" { return "Statistics" }
    return ""
}

func (vrf *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf) GetSegmentPath() string {
    return "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
}

func (vrf *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "statistics" {
        return &vrf.Statistics
    }
    return nil
}

func (vrf *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["statistics"] = &vrf.Statistics
    return children
}

func (vrf *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = vrf.VrfName
    return leafs
}

func (vrf *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf) GetBundleName() string { return "cisco_ios_xr" }

func (vrf *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf) GetYangName() string { return "vrf" }

func (vrf *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrf *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrf *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrf *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf) SetParent(parent types.Entity) { vrf.parent = parent }

func (vrf *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf) GetParent() types.Entity { return vrf.parent }

func (vrf *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf) GetParentYangName() string { return "vrfs" }

// Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics
// IPv6 DHCP proxy statistics
type Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DHCPV6 solicit packets.
    Solicit Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Solicit

    // DHCPV6 advertise packets.
    Advertise Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Advertise

    // DHCPV6 request packets.
    Request Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Request

    // DHCPV6 reply packets.
    Reply Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Reply

    // DHCPV6 confirm packets.
    Confirm Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Confirm

    // DHCPV6 decline packets.
    Decline Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Decline

    // DHCPV6 renew packets.
    Renew Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Renew

    // DHCPV6 rebind packets.
    Rebind Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Rebind

    // DHCPV6 release packets.
    Release Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Release

    // DHCPV6 reconfig packets.
    Reconfig Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Reconfig

    // DHCPV6 inform packets.
    Inform Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Inform

    // DHCPV6 relay forward packets.
    RelayForward Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_RelayForward

    // DHCPV6 relay reply packets.
    RelayReply Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_RelayReply

    // DHCPV6 lease query packets.
    LeaseQuery Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQuery

    // DHCPV6 lease query reply packets.
    LeaseQueryReply Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQueryReply

    // DHCPV6 lease query done packets.
    LeaseQueryDone Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQueryDone

    // DHCPV6 lease query data packets.
    LeaseQueryData Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQueryData
}

func (statistics *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics) GetGoName(yname string) string {
    if yname == "solicit" { return "Solicit" }
    if yname == "advertise" { return "Advertise" }
    if yname == "request" { return "Request" }
    if yname == "reply" { return "Reply" }
    if yname == "confirm" { return "Confirm" }
    if yname == "decline" { return "Decline" }
    if yname == "renew" { return "Renew" }
    if yname == "rebind" { return "Rebind" }
    if yname == "release" { return "Release" }
    if yname == "reconfig" { return "Reconfig" }
    if yname == "inform" { return "Inform" }
    if yname == "relay-forward" { return "RelayForward" }
    if yname == "relay-reply" { return "RelayReply" }
    if yname == "lease-query" { return "LeaseQuery" }
    if yname == "lease-query-reply" { return "LeaseQueryReply" }
    if yname == "lease-query-done" { return "LeaseQueryDone" }
    if yname == "lease-query-data" { return "LeaseQueryData" }
    return ""
}

func (statistics *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "solicit" {
        return &statistics.Solicit
    }
    if childYangName == "advertise" {
        return &statistics.Advertise
    }
    if childYangName == "request" {
        return &statistics.Request
    }
    if childYangName == "reply" {
        return &statistics.Reply
    }
    if childYangName == "confirm" {
        return &statistics.Confirm
    }
    if childYangName == "decline" {
        return &statistics.Decline
    }
    if childYangName == "renew" {
        return &statistics.Renew
    }
    if childYangName == "rebind" {
        return &statistics.Rebind
    }
    if childYangName == "release" {
        return &statistics.Release
    }
    if childYangName == "reconfig" {
        return &statistics.Reconfig
    }
    if childYangName == "inform" {
        return &statistics.Inform
    }
    if childYangName == "relay-forward" {
        return &statistics.RelayForward
    }
    if childYangName == "relay-reply" {
        return &statistics.RelayReply
    }
    if childYangName == "lease-query" {
        return &statistics.LeaseQuery
    }
    if childYangName == "lease-query-reply" {
        return &statistics.LeaseQueryReply
    }
    if childYangName == "lease-query-done" {
        return &statistics.LeaseQueryDone
    }
    if childYangName == "lease-query-data" {
        return &statistics.LeaseQueryData
    }
    return nil
}

func (statistics *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["solicit"] = &statistics.Solicit
    children["advertise"] = &statistics.Advertise
    children["request"] = &statistics.Request
    children["reply"] = &statistics.Reply
    children["confirm"] = &statistics.Confirm
    children["decline"] = &statistics.Decline
    children["renew"] = &statistics.Renew
    children["rebind"] = &statistics.Rebind
    children["release"] = &statistics.Release
    children["reconfig"] = &statistics.Reconfig
    children["inform"] = &statistics.Inform
    children["relay-forward"] = &statistics.RelayForward
    children["relay-reply"] = &statistics.RelayReply
    children["lease-query"] = &statistics.LeaseQuery
    children["lease-query-reply"] = &statistics.LeaseQueryReply
    children["lease-query-done"] = &statistics.LeaseQueryDone
    children["lease-query-data"] = &statistics.LeaseQueryData
    return children
}

func (statistics *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (statistics *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics) GetYangName() string { return "statistics" }

func (statistics *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics) GetParentYangName() string { return "vrf" }

// Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Solicit
// DHCPV6 solicit packets
type Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Solicit struct {
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

func (solicit *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Solicit) GetFilter() yfilter.YFilter { return solicit.YFilter }

func (solicit *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Solicit) SetFilter(yf yfilter.YFilter) { solicit.YFilter = yf }

func (solicit *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Solicit) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (solicit *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Solicit) GetSegmentPath() string {
    return "solicit"
}

func (solicit *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Solicit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (solicit *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Solicit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (solicit *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Solicit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = solicit.ReceivedPackets
    leafs["transmitted-packets"] = solicit.TransmittedPackets
    leafs["dropped-packets"] = solicit.DroppedPackets
    return leafs
}

func (solicit *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Solicit) GetBundleName() string { return "cisco_ios_xr" }

func (solicit *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Solicit) GetYangName() string { return "solicit" }

func (solicit *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Solicit) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (solicit *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Solicit) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (solicit *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Solicit) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (solicit *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Solicit) SetParent(parent types.Entity) { solicit.parent = parent }

func (solicit *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Solicit) GetParent() types.Entity { return solicit.parent }

func (solicit *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Solicit) GetParentYangName() string { return "statistics" }

// Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Advertise
// DHCPV6 advertise packets
type Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Advertise struct {
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

func (advertise *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Advertise) GetFilter() yfilter.YFilter { return advertise.YFilter }

func (advertise *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Advertise) SetFilter(yf yfilter.YFilter) { advertise.YFilter = yf }

func (advertise *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Advertise) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (advertise *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Advertise) GetSegmentPath() string {
    return "advertise"
}

func (advertise *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Advertise) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (advertise *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Advertise) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (advertise *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Advertise) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = advertise.ReceivedPackets
    leafs["transmitted-packets"] = advertise.TransmittedPackets
    leafs["dropped-packets"] = advertise.DroppedPackets
    return leafs
}

func (advertise *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Advertise) GetBundleName() string { return "cisco_ios_xr" }

func (advertise *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Advertise) GetYangName() string { return "advertise" }

func (advertise *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Advertise) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (advertise *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Advertise) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (advertise *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Advertise) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (advertise *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Advertise) SetParent(parent types.Entity) { advertise.parent = parent }

func (advertise *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Advertise) GetParent() types.Entity { return advertise.parent }

func (advertise *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Advertise) GetParentYangName() string { return "statistics" }

// Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Request
// DHCPV6 request packets
type Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Request struct {
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

func (request *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Request) GetFilter() yfilter.YFilter { return request.YFilter }

func (request *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Request) SetFilter(yf yfilter.YFilter) { request.YFilter = yf }

func (request *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Request) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (request *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Request) GetSegmentPath() string {
    return "request"
}

func (request *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Request) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (request *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Request) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (request *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Request) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = request.ReceivedPackets
    leafs["transmitted-packets"] = request.TransmittedPackets
    leafs["dropped-packets"] = request.DroppedPackets
    return leafs
}

func (request *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Request) GetBundleName() string { return "cisco_ios_xr" }

func (request *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Request) GetYangName() string { return "request" }

func (request *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Request) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (request *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Request) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (request *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Request) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (request *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Request) SetParent(parent types.Entity) { request.parent = parent }

func (request *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Request) GetParent() types.Entity { return request.parent }

func (request *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Request) GetParentYangName() string { return "statistics" }

// Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Reply
// DHCPV6 reply packets
type Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Reply struct {
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

func (reply *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Reply) GetFilter() yfilter.YFilter { return reply.YFilter }

func (reply *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Reply) SetFilter(yf yfilter.YFilter) { reply.YFilter = yf }

func (reply *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Reply) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (reply *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Reply) GetSegmentPath() string {
    return "reply"
}

func (reply *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Reply) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (reply *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Reply) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (reply *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Reply) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = reply.ReceivedPackets
    leafs["transmitted-packets"] = reply.TransmittedPackets
    leafs["dropped-packets"] = reply.DroppedPackets
    return leafs
}

func (reply *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Reply) GetBundleName() string { return "cisco_ios_xr" }

func (reply *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Reply) GetYangName() string { return "reply" }

func (reply *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Reply) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (reply *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Reply) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (reply *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Reply) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (reply *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Reply) SetParent(parent types.Entity) { reply.parent = parent }

func (reply *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Reply) GetParent() types.Entity { return reply.parent }

func (reply *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Reply) GetParentYangName() string { return "statistics" }

// Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Confirm
// DHCPV6 confirm packets
type Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Confirm struct {
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

func (confirm *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Confirm) GetFilter() yfilter.YFilter { return confirm.YFilter }

func (confirm *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Confirm) SetFilter(yf yfilter.YFilter) { confirm.YFilter = yf }

func (confirm *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Confirm) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (confirm *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Confirm) GetSegmentPath() string {
    return "confirm"
}

func (confirm *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Confirm) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (confirm *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Confirm) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (confirm *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Confirm) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = confirm.ReceivedPackets
    leafs["transmitted-packets"] = confirm.TransmittedPackets
    leafs["dropped-packets"] = confirm.DroppedPackets
    return leafs
}

func (confirm *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Confirm) GetBundleName() string { return "cisco_ios_xr" }

func (confirm *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Confirm) GetYangName() string { return "confirm" }

func (confirm *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Confirm) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (confirm *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Confirm) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (confirm *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Confirm) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (confirm *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Confirm) SetParent(parent types.Entity) { confirm.parent = parent }

func (confirm *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Confirm) GetParent() types.Entity { return confirm.parent }

func (confirm *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Confirm) GetParentYangName() string { return "statistics" }

// Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Decline
// DHCPV6 decline packets
type Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Decline struct {
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

func (decline *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Decline) GetFilter() yfilter.YFilter { return decline.YFilter }

func (decline *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Decline) SetFilter(yf yfilter.YFilter) { decline.YFilter = yf }

func (decline *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Decline) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (decline *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Decline) GetSegmentPath() string {
    return "decline"
}

func (decline *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Decline) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (decline *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Decline) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (decline *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Decline) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = decline.ReceivedPackets
    leafs["transmitted-packets"] = decline.TransmittedPackets
    leafs["dropped-packets"] = decline.DroppedPackets
    return leafs
}

func (decline *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Decline) GetBundleName() string { return "cisco_ios_xr" }

func (decline *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Decline) GetYangName() string { return "decline" }

func (decline *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Decline) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (decline *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Decline) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (decline *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Decline) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (decline *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Decline) SetParent(parent types.Entity) { decline.parent = parent }

func (decline *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Decline) GetParent() types.Entity { return decline.parent }

func (decline *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Decline) GetParentYangName() string { return "statistics" }

// Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Renew
// DHCPV6 renew packets
type Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Renew struct {
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

func (renew *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Renew) GetFilter() yfilter.YFilter { return renew.YFilter }

func (renew *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Renew) SetFilter(yf yfilter.YFilter) { renew.YFilter = yf }

func (renew *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Renew) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (renew *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Renew) GetSegmentPath() string {
    return "renew"
}

func (renew *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Renew) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (renew *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Renew) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (renew *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Renew) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = renew.ReceivedPackets
    leafs["transmitted-packets"] = renew.TransmittedPackets
    leafs["dropped-packets"] = renew.DroppedPackets
    return leafs
}

func (renew *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Renew) GetBundleName() string { return "cisco_ios_xr" }

func (renew *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Renew) GetYangName() string { return "renew" }

func (renew *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Renew) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (renew *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Renew) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (renew *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Renew) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (renew *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Renew) SetParent(parent types.Entity) { renew.parent = parent }

func (renew *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Renew) GetParent() types.Entity { return renew.parent }

func (renew *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Renew) GetParentYangName() string { return "statistics" }

// Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Rebind
// DHCPV6 rebind packets
type Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Rebind struct {
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

func (rebind *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Rebind) GetFilter() yfilter.YFilter { return rebind.YFilter }

func (rebind *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Rebind) SetFilter(yf yfilter.YFilter) { rebind.YFilter = yf }

func (rebind *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Rebind) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (rebind *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Rebind) GetSegmentPath() string {
    return "rebind"
}

func (rebind *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Rebind) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rebind *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Rebind) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rebind *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Rebind) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = rebind.ReceivedPackets
    leafs["transmitted-packets"] = rebind.TransmittedPackets
    leafs["dropped-packets"] = rebind.DroppedPackets
    return leafs
}

func (rebind *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Rebind) GetBundleName() string { return "cisco_ios_xr" }

func (rebind *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Rebind) GetYangName() string { return "rebind" }

func (rebind *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Rebind) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rebind *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Rebind) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rebind *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Rebind) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rebind *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Rebind) SetParent(parent types.Entity) { rebind.parent = parent }

func (rebind *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Rebind) GetParent() types.Entity { return rebind.parent }

func (rebind *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Rebind) GetParentYangName() string { return "statistics" }

// Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Release
// DHCPV6 release packets
type Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Release struct {
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

func (release *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Release) GetFilter() yfilter.YFilter { return release.YFilter }

func (release *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Release) SetFilter(yf yfilter.YFilter) { release.YFilter = yf }

func (release *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Release) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (release *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Release) GetSegmentPath() string {
    return "release"
}

func (release *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Release) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (release *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Release) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (release *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Release) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = release.ReceivedPackets
    leafs["transmitted-packets"] = release.TransmittedPackets
    leafs["dropped-packets"] = release.DroppedPackets
    return leafs
}

func (release *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Release) GetBundleName() string { return "cisco_ios_xr" }

func (release *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Release) GetYangName() string { return "release" }

func (release *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Release) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (release *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Release) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (release *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Release) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (release *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Release) SetParent(parent types.Entity) { release.parent = parent }

func (release *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Release) GetParent() types.Entity { return release.parent }

func (release *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Release) GetParentYangName() string { return "statistics" }

// Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Reconfig
// DHCPV6 reconfig packets
type Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Reconfig struct {
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

func (reconfig *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Reconfig) GetFilter() yfilter.YFilter { return reconfig.YFilter }

func (reconfig *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Reconfig) SetFilter(yf yfilter.YFilter) { reconfig.YFilter = yf }

func (reconfig *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Reconfig) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (reconfig *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Reconfig) GetSegmentPath() string {
    return "reconfig"
}

func (reconfig *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Reconfig) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (reconfig *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Reconfig) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (reconfig *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Reconfig) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = reconfig.ReceivedPackets
    leafs["transmitted-packets"] = reconfig.TransmittedPackets
    leafs["dropped-packets"] = reconfig.DroppedPackets
    return leafs
}

func (reconfig *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Reconfig) GetBundleName() string { return "cisco_ios_xr" }

func (reconfig *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Reconfig) GetYangName() string { return "reconfig" }

func (reconfig *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Reconfig) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (reconfig *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Reconfig) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (reconfig *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Reconfig) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (reconfig *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Reconfig) SetParent(parent types.Entity) { reconfig.parent = parent }

func (reconfig *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Reconfig) GetParent() types.Entity { return reconfig.parent }

func (reconfig *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Reconfig) GetParentYangName() string { return "statistics" }

// Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Inform
// DHCPV6 inform packets
type Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Inform struct {
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

func (inform *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Inform) GetFilter() yfilter.YFilter { return inform.YFilter }

func (inform *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Inform) SetFilter(yf yfilter.YFilter) { inform.YFilter = yf }

func (inform *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Inform) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (inform *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Inform) GetSegmentPath() string {
    return "inform"
}

func (inform *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Inform) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inform *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Inform) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inform *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Inform) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = inform.ReceivedPackets
    leafs["transmitted-packets"] = inform.TransmittedPackets
    leafs["dropped-packets"] = inform.DroppedPackets
    return leafs
}

func (inform *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Inform) GetBundleName() string { return "cisco_ios_xr" }

func (inform *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Inform) GetYangName() string { return "inform" }

func (inform *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Inform) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inform *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Inform) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inform *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Inform) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inform *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Inform) SetParent(parent types.Entity) { inform.parent = parent }

func (inform *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Inform) GetParent() types.Entity { return inform.parent }

func (inform *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Inform) GetParentYangName() string { return "statistics" }

// Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_RelayForward
// DHCPV6 relay forward packets
type Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_RelayForward struct {
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

func (relayForward *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_RelayForward) GetFilter() yfilter.YFilter { return relayForward.YFilter }

func (relayForward *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_RelayForward) SetFilter(yf yfilter.YFilter) { relayForward.YFilter = yf }

func (relayForward *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_RelayForward) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (relayForward *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_RelayForward) GetSegmentPath() string {
    return "relay-forward"
}

func (relayForward *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_RelayForward) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (relayForward *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_RelayForward) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (relayForward *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_RelayForward) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = relayForward.ReceivedPackets
    leafs["transmitted-packets"] = relayForward.TransmittedPackets
    leafs["dropped-packets"] = relayForward.DroppedPackets
    return leafs
}

func (relayForward *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_RelayForward) GetBundleName() string { return "cisco_ios_xr" }

func (relayForward *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_RelayForward) GetYangName() string { return "relay-forward" }

func (relayForward *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_RelayForward) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (relayForward *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_RelayForward) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (relayForward *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_RelayForward) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (relayForward *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_RelayForward) SetParent(parent types.Entity) { relayForward.parent = parent }

func (relayForward *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_RelayForward) GetParent() types.Entity { return relayForward.parent }

func (relayForward *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_RelayForward) GetParentYangName() string { return "statistics" }

// Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_RelayReply
// DHCPV6 relay reply packets
type Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_RelayReply struct {
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

func (relayReply *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_RelayReply) GetFilter() yfilter.YFilter { return relayReply.YFilter }

func (relayReply *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_RelayReply) SetFilter(yf yfilter.YFilter) { relayReply.YFilter = yf }

func (relayReply *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_RelayReply) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (relayReply *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_RelayReply) GetSegmentPath() string {
    return "relay-reply"
}

func (relayReply *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_RelayReply) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (relayReply *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_RelayReply) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (relayReply *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_RelayReply) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = relayReply.ReceivedPackets
    leafs["transmitted-packets"] = relayReply.TransmittedPackets
    leafs["dropped-packets"] = relayReply.DroppedPackets
    return leafs
}

func (relayReply *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_RelayReply) GetBundleName() string { return "cisco_ios_xr" }

func (relayReply *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_RelayReply) GetYangName() string { return "relay-reply" }

func (relayReply *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_RelayReply) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (relayReply *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_RelayReply) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (relayReply *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_RelayReply) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (relayReply *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_RelayReply) SetParent(parent types.Entity) { relayReply.parent = parent }

func (relayReply *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_RelayReply) GetParent() types.Entity { return relayReply.parent }

func (relayReply *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_RelayReply) GetParentYangName() string { return "statistics" }

// Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQuery
// DHCPV6 lease query packets
type Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQuery struct {
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

func (leaseQuery *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQuery) GetFilter() yfilter.YFilter { return leaseQuery.YFilter }

func (leaseQuery *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQuery) SetFilter(yf yfilter.YFilter) { leaseQuery.YFilter = yf }

func (leaseQuery *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQuery) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (leaseQuery *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQuery) GetSegmentPath() string {
    return "lease-query"
}

func (leaseQuery *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQuery) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (leaseQuery *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQuery) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (leaseQuery *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQuery) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = leaseQuery.ReceivedPackets
    leafs["transmitted-packets"] = leaseQuery.TransmittedPackets
    leafs["dropped-packets"] = leaseQuery.DroppedPackets
    return leafs
}

func (leaseQuery *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQuery) GetBundleName() string { return "cisco_ios_xr" }

func (leaseQuery *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQuery) GetYangName() string { return "lease-query" }

func (leaseQuery *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQuery) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (leaseQuery *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQuery) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (leaseQuery *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQuery) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (leaseQuery *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQuery) SetParent(parent types.Entity) { leaseQuery.parent = parent }

func (leaseQuery *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQuery) GetParent() types.Entity { return leaseQuery.parent }

func (leaseQuery *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQuery) GetParentYangName() string { return "statistics" }

// Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQueryReply
// DHCPV6 lease query reply packets
type Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQueryReply struct {
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

func (leaseQueryReply *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQueryReply) GetFilter() yfilter.YFilter { return leaseQueryReply.YFilter }

func (leaseQueryReply *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQueryReply) SetFilter(yf yfilter.YFilter) { leaseQueryReply.YFilter = yf }

func (leaseQueryReply *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQueryReply) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (leaseQueryReply *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQueryReply) GetSegmentPath() string {
    return "lease-query-reply"
}

func (leaseQueryReply *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQueryReply) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (leaseQueryReply *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQueryReply) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (leaseQueryReply *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQueryReply) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = leaseQueryReply.ReceivedPackets
    leafs["transmitted-packets"] = leaseQueryReply.TransmittedPackets
    leafs["dropped-packets"] = leaseQueryReply.DroppedPackets
    return leafs
}

func (leaseQueryReply *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQueryReply) GetBundleName() string { return "cisco_ios_xr" }

func (leaseQueryReply *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQueryReply) GetYangName() string { return "lease-query-reply" }

func (leaseQueryReply *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQueryReply) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (leaseQueryReply *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQueryReply) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (leaseQueryReply *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQueryReply) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (leaseQueryReply *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQueryReply) SetParent(parent types.Entity) { leaseQueryReply.parent = parent }

func (leaseQueryReply *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQueryReply) GetParent() types.Entity { return leaseQueryReply.parent }

func (leaseQueryReply *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQueryReply) GetParentYangName() string { return "statistics" }

// Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQueryDone
// DHCPV6 lease query done packets
type Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQueryDone struct {
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

func (leaseQueryDone *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQueryDone) GetFilter() yfilter.YFilter { return leaseQueryDone.YFilter }

func (leaseQueryDone *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQueryDone) SetFilter(yf yfilter.YFilter) { leaseQueryDone.YFilter = yf }

func (leaseQueryDone *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQueryDone) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (leaseQueryDone *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQueryDone) GetSegmentPath() string {
    return "lease-query-done"
}

func (leaseQueryDone *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQueryDone) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (leaseQueryDone *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQueryDone) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (leaseQueryDone *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQueryDone) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = leaseQueryDone.ReceivedPackets
    leafs["transmitted-packets"] = leaseQueryDone.TransmittedPackets
    leafs["dropped-packets"] = leaseQueryDone.DroppedPackets
    return leafs
}

func (leaseQueryDone *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQueryDone) GetBundleName() string { return "cisco_ios_xr" }

func (leaseQueryDone *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQueryDone) GetYangName() string { return "lease-query-done" }

func (leaseQueryDone *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQueryDone) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (leaseQueryDone *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQueryDone) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (leaseQueryDone *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQueryDone) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (leaseQueryDone *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQueryDone) SetParent(parent types.Entity) { leaseQueryDone.parent = parent }

func (leaseQueryDone *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQueryDone) GetParent() types.Entity { return leaseQueryDone.parent }

func (leaseQueryDone *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQueryDone) GetParentYangName() string { return "statistics" }

// Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQueryData
// DHCPV6 lease query data packets
type Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQueryData struct {
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

func (leaseQueryData *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQueryData) GetFilter() yfilter.YFilter { return leaseQueryData.YFilter }

func (leaseQueryData *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQueryData) SetFilter(yf yfilter.YFilter) { leaseQueryData.YFilter = yf }

func (leaseQueryData *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQueryData) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (leaseQueryData *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQueryData) GetSegmentPath() string {
    return "lease-query-data"
}

func (leaseQueryData *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQueryData) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (leaseQueryData *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQueryData) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (leaseQueryData *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQueryData) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = leaseQueryData.ReceivedPackets
    leafs["transmitted-packets"] = leaseQueryData.TransmittedPackets
    leafs["dropped-packets"] = leaseQueryData.DroppedPackets
    return leafs
}

func (leaseQueryData *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQueryData) GetBundleName() string { return "cisco_ios_xr" }

func (leaseQueryData *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQueryData) GetYangName() string { return "lease-query-data" }

func (leaseQueryData *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQueryData) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (leaseQueryData *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQueryData) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (leaseQueryData *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQueryData) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (leaseQueryData *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQueryData) SetParent(parent types.Entity) { leaseQueryData.parent = parent }

func (leaseQueryData *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQueryData) GetParent() types.Entity { return leaseQueryData.parent }

func (leaseQueryData *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQueryData) GetParentYangName() string { return "statistics" }

// Dhcpv6_Nodes_Node_Proxy_Profiles
// IPv6 DHCP proxy profile
type Dhcpv6_Nodes_Node_Proxy_Profiles struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 DHCP proxy profile. The type is slice of
    // Dhcpv6_Nodes_Node_Proxy_Profiles_Profile.
    Profile []Dhcpv6_Nodes_Node_Proxy_Profiles_Profile
}

func (profiles *Dhcpv6_Nodes_Node_Proxy_Profiles) GetFilter() yfilter.YFilter { return profiles.YFilter }

func (profiles *Dhcpv6_Nodes_Node_Proxy_Profiles) SetFilter(yf yfilter.YFilter) { profiles.YFilter = yf }

func (profiles *Dhcpv6_Nodes_Node_Proxy_Profiles) GetGoName(yname string) string {
    if yname == "profile" { return "Profile" }
    return ""
}

func (profiles *Dhcpv6_Nodes_Node_Proxy_Profiles) GetSegmentPath() string {
    return "profiles"
}

func (profiles *Dhcpv6_Nodes_Node_Proxy_Profiles) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "profile" {
        for _, c := range profiles.Profile {
            if profiles.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dhcpv6_Nodes_Node_Proxy_Profiles_Profile{}
        profiles.Profile = append(profiles.Profile, child)
        return &profiles.Profile[len(profiles.Profile)-1]
    }
    return nil
}

func (profiles *Dhcpv6_Nodes_Node_Proxy_Profiles) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range profiles.Profile {
        children[profiles.Profile[i].GetSegmentPath()] = &profiles.Profile[i]
    }
    return children
}

func (profiles *Dhcpv6_Nodes_Node_Proxy_Profiles) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (profiles *Dhcpv6_Nodes_Node_Proxy_Profiles) GetBundleName() string { return "cisco_ios_xr" }

func (profiles *Dhcpv6_Nodes_Node_Proxy_Profiles) GetYangName() string { return "profiles" }

func (profiles *Dhcpv6_Nodes_Node_Proxy_Profiles) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (profiles *Dhcpv6_Nodes_Node_Proxy_Profiles) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (profiles *Dhcpv6_Nodes_Node_Proxy_Profiles) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (profiles *Dhcpv6_Nodes_Node_Proxy_Profiles) SetParent(parent types.Entity) { profiles.parent = parent }

func (profiles *Dhcpv6_Nodes_Node_Proxy_Profiles) GetParent() types.Entity { return profiles.parent }

func (profiles *Dhcpv6_Nodes_Node_Proxy_Profiles) GetParentYangName() string { return "proxy" }

// Dhcpv6_Nodes_Node_Proxy_Profiles_Profile
// IPv6 DHCP proxy profile
type Dhcpv6_Nodes_Node_Proxy_Profiles_Profile struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Profile name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    ProfileName interface{}

    // DHCPV6 throttle table.
    ThrottleInfos Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_ThrottleInfos

    // IPv6 DHCP proxy profile Info.
    Info Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info
}

func (profile *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile) GetFilter() yfilter.YFilter { return profile.YFilter }

func (profile *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile) SetFilter(yf yfilter.YFilter) { profile.YFilter = yf }

func (profile *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile) GetGoName(yname string) string {
    if yname == "profile-name" { return "ProfileName" }
    if yname == "throttle-infos" { return "ThrottleInfos" }
    if yname == "info" { return "Info" }
    return ""
}

func (profile *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile) GetSegmentPath() string {
    return "profile" + "[profile-name='" + fmt.Sprintf("%v", profile.ProfileName) + "']"
}

func (profile *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "throttle-infos" {
        return &profile.ThrottleInfos
    }
    if childYangName == "info" {
        return &profile.Info
    }
    return nil
}

func (profile *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["throttle-infos"] = &profile.ThrottleInfos
    children["info"] = &profile.Info
    return children
}

func (profile *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["profile-name"] = profile.ProfileName
    return leafs
}

func (profile *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile) GetBundleName() string { return "cisco_ios_xr" }

func (profile *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile) GetYangName() string { return "profile" }

func (profile *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (profile *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (profile *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (profile *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile) SetParent(parent types.Entity) { profile.parent = parent }

func (profile *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile) GetParent() types.Entity { return profile.parent }

func (profile *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile) GetParentYangName() string { return "profiles" }

// Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_ThrottleInfos
// DHCPV6 throttle table
type Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_ThrottleInfos struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 DHCP proxy profile Throttle Info. The type is slice of
    // Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_ThrottleInfos_ThrottleInfo.
    ThrottleInfo []Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_ThrottleInfos_ThrottleInfo
}

func (throttleInfos *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_ThrottleInfos) GetFilter() yfilter.YFilter { return throttleInfos.YFilter }

func (throttleInfos *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_ThrottleInfos) SetFilter(yf yfilter.YFilter) { throttleInfos.YFilter = yf }

func (throttleInfos *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_ThrottleInfos) GetGoName(yname string) string {
    if yname == "throttle-info" { return "ThrottleInfo" }
    return ""
}

func (throttleInfos *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_ThrottleInfos) GetSegmentPath() string {
    return "throttle-infos"
}

func (throttleInfos *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_ThrottleInfos) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "throttle-info" {
        for _, c := range throttleInfos.ThrottleInfo {
            if throttleInfos.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_ThrottleInfos_ThrottleInfo{}
        throttleInfos.ThrottleInfo = append(throttleInfos.ThrottleInfo, child)
        return &throttleInfos.ThrottleInfo[len(throttleInfos.ThrottleInfo)-1]
    }
    return nil
}

func (throttleInfos *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_ThrottleInfos) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range throttleInfos.ThrottleInfo {
        children[throttleInfos.ThrottleInfo[i].GetSegmentPath()] = &throttleInfos.ThrottleInfo[i]
    }
    return children
}

func (throttleInfos *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_ThrottleInfos) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (throttleInfos *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_ThrottleInfos) GetBundleName() string { return "cisco_ios_xr" }

func (throttleInfos *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_ThrottleInfos) GetYangName() string { return "throttle-infos" }

func (throttleInfos *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_ThrottleInfos) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (throttleInfos *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_ThrottleInfos) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (throttleInfos *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_ThrottleInfos) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (throttleInfos *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_ThrottleInfos) SetParent(parent types.Entity) { throttleInfos.parent = parent }

func (throttleInfos *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_ThrottleInfos) GetParent() types.Entity { return throttleInfos.parent }

func (throttleInfos *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_ThrottleInfos) GetParentYangName() string { return "profile" }

// Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_ThrottleInfos_ThrottleInfo
// IPv6 DHCP proxy profile Throttle Info
type Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_ThrottleInfos_ThrottleInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. MAC address. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    MacAddress interface{}

    // Client MAC address. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    BindingChaddr interface{}

    // DHCP access interface. The type is string with length: 0..65.
    Ifname interface{}

    // State of entry. The type is interface{} with range: 0..4294967295.
    State interface{}

    // Time Left in secs. The type is interface{} with range: 0..4294967295. Units
    // are second.
    TimeLeft interface{}
}

func (throttleInfo *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_ThrottleInfos_ThrottleInfo) GetFilter() yfilter.YFilter { return throttleInfo.YFilter }

func (throttleInfo *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_ThrottleInfos_ThrottleInfo) SetFilter(yf yfilter.YFilter) { throttleInfo.YFilter = yf }

func (throttleInfo *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_ThrottleInfos_ThrottleInfo) GetGoName(yname string) string {
    if yname == "mac-address" { return "MacAddress" }
    if yname == "binding-chaddr" { return "BindingChaddr" }
    if yname == "ifname" { return "Ifname" }
    if yname == "state" { return "State" }
    if yname == "time-left" { return "TimeLeft" }
    return ""
}

func (throttleInfo *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_ThrottleInfos_ThrottleInfo) GetSegmentPath() string {
    return "throttle-info" + "[mac-address='" + fmt.Sprintf("%v", throttleInfo.MacAddress) + "']"
}

func (throttleInfo *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_ThrottleInfos_ThrottleInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (throttleInfo *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_ThrottleInfos_ThrottleInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (throttleInfo *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_ThrottleInfos_ThrottleInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mac-address"] = throttleInfo.MacAddress
    leafs["binding-chaddr"] = throttleInfo.BindingChaddr
    leafs["ifname"] = throttleInfo.Ifname
    leafs["state"] = throttleInfo.State
    leafs["time-left"] = throttleInfo.TimeLeft
    return leafs
}

func (throttleInfo *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_ThrottleInfos_ThrottleInfo) GetBundleName() string { return "cisco_ios_xr" }

func (throttleInfo *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_ThrottleInfos_ThrottleInfo) GetYangName() string { return "throttle-info" }

func (throttleInfo *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_ThrottleInfos_ThrottleInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (throttleInfo *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_ThrottleInfos_ThrottleInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (throttleInfo *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_ThrottleInfos_ThrottleInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (throttleInfo *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_ThrottleInfos_ThrottleInfo) SetParent(parent types.Entity) { throttleInfo.parent = parent }

func (throttleInfo *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_ThrottleInfos_ThrottleInfo) GetParent() types.Entity { return throttleInfo.parent }

func (throttleInfo *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_ThrottleInfos_ThrottleInfo) GetParentYangName() string { return "throttle-infos" }

// Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info
// IPv6 DHCP proxy profile Info
type Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Proxy profile name. The type is string with length: 0..65.
    ProfileName interface{}

    // Remote id. The type is string with length: 0..257.
    RemoteId interface{}

    // Link address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    ProfileLinkAddress interface{}

    // LinkAddress From RA mesage. The type is bool.
    ProxyProfileLinkaddressFromRaEnable interface{}

    // Helper addresses. The type is slice of string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    ProfileHelperAddress []interface{}

    // VRF names. The type is slice of string with length: 0..33.
    VrfName []interface{}

    // Interface names. The type is slice of string with length: 0..65.
    InterfaceName []interface{}

    // Interface id references.
    InterfaceIdReferences Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceIdReferences

    // VRF references.
    VrfReferences Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_VrfReferences

    // Interface references.
    InterfaceReferences Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceReferences
}

func (info *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info) GetFilter() yfilter.YFilter { return info.YFilter }

func (info *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info) SetFilter(yf yfilter.YFilter) { info.YFilter = yf }

func (info *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info) GetGoName(yname string) string {
    if yname == "profile-name" { return "ProfileName" }
    if yname == "remote-id" { return "RemoteId" }
    if yname == "profile-link-address" { return "ProfileLinkAddress" }
    if yname == "proxy-profile-linkaddress-from-ra-enable" { return "ProxyProfileLinkaddressFromRaEnable" }
    if yname == "profile-helper-address" { return "ProfileHelperAddress" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "interface-id-references" { return "InterfaceIdReferences" }
    if yname == "vrf-references" { return "VrfReferences" }
    if yname == "interface-references" { return "InterfaceReferences" }
    return ""
}

func (info *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info) GetSegmentPath() string {
    return "info"
}

func (info *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-id-references" {
        return &info.InterfaceIdReferences
    }
    if childYangName == "vrf-references" {
        return &info.VrfReferences
    }
    if childYangName == "interface-references" {
        return &info.InterfaceReferences
    }
    return nil
}

func (info *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interface-id-references"] = &info.InterfaceIdReferences
    children["vrf-references"] = &info.VrfReferences
    children["interface-references"] = &info.InterfaceReferences
    return children
}

func (info *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["profile-name"] = info.ProfileName
    leafs["remote-id"] = info.RemoteId
    leafs["profile-link-address"] = info.ProfileLinkAddress
    leafs["proxy-profile-linkaddress-from-ra-enable"] = info.ProxyProfileLinkaddressFromRaEnable
    leafs["profile-helper-address"] = info.ProfileHelperAddress
    leafs["vrf-name"] = info.VrfName
    leafs["interface-name"] = info.InterfaceName
    return leafs
}

func (info *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info) GetBundleName() string { return "cisco_ios_xr" }

func (info *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info) GetYangName() string { return "info" }

func (info *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (info *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (info *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (info *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info) SetParent(parent types.Entity) { info.parent = parent }

func (info *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info) GetParent() types.Entity { return info.parent }

func (info *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info) GetParentYangName() string { return "profile" }

// Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceIdReferences
// Interface id references
type Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceIdReferences struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ipv6 dhcpv6d proxy iid reference. The type is slice of
    // Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceIdReferences_Ipv6Dhcpv6DProxyIidReference.
    Ipv6Dhcpv6DProxyIidReference []Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceIdReferences_Ipv6Dhcpv6DProxyIidReference
}

func (interfaceIdReferences *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceIdReferences) GetFilter() yfilter.YFilter { return interfaceIdReferences.YFilter }

func (interfaceIdReferences *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceIdReferences) SetFilter(yf yfilter.YFilter) { interfaceIdReferences.YFilter = yf }

func (interfaceIdReferences *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceIdReferences) GetGoName(yname string) string {
    if yname == "ipv6-dhcpv6d-proxy-iid-reference" { return "Ipv6Dhcpv6DProxyIidReference" }
    return ""
}

func (interfaceIdReferences *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceIdReferences) GetSegmentPath() string {
    return "interface-id-references"
}

func (interfaceIdReferences *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceIdReferences) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv6-dhcpv6d-proxy-iid-reference" {
        for _, c := range interfaceIdReferences.Ipv6Dhcpv6DProxyIidReference {
            if interfaceIdReferences.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceIdReferences_Ipv6Dhcpv6DProxyIidReference{}
        interfaceIdReferences.Ipv6Dhcpv6DProxyIidReference = append(interfaceIdReferences.Ipv6Dhcpv6DProxyIidReference, child)
        return &interfaceIdReferences.Ipv6Dhcpv6DProxyIidReference[len(interfaceIdReferences.Ipv6Dhcpv6DProxyIidReference)-1]
    }
    return nil
}

func (interfaceIdReferences *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceIdReferences) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaceIdReferences.Ipv6Dhcpv6DProxyIidReference {
        children[interfaceIdReferences.Ipv6Dhcpv6DProxyIidReference[i].GetSegmentPath()] = &interfaceIdReferences.Ipv6Dhcpv6DProxyIidReference[i]
    }
    return children
}

func (interfaceIdReferences *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceIdReferences) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceIdReferences *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceIdReferences) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceIdReferences *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceIdReferences) GetYangName() string { return "interface-id-references" }

func (interfaceIdReferences *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceIdReferences) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceIdReferences *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceIdReferences) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceIdReferences *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceIdReferences) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceIdReferences *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceIdReferences) SetParent(parent types.Entity) { interfaceIdReferences.parent = parent }

func (interfaceIdReferences *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceIdReferences) GetParent() types.Entity { return interfaceIdReferences.parent }

func (interfaceIdReferences *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceIdReferences) GetParentYangName() string { return "info" }

// Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceIdReferences_Ipv6Dhcpv6DProxyIidReference
// ipv6 dhcpv6d proxy iid reference
type Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceIdReferences_Ipv6Dhcpv6DProxyIidReference struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface name for interface id. The type is string with length: 0..65.
    ProxyIidInterfaceName interface{}

    // Interface id. The type is string with length: 0..257.
    ProxyInterfaceId interface{}
}

func (ipv6Dhcpv6DProxyIidReference *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceIdReferences_Ipv6Dhcpv6DProxyIidReference) GetFilter() yfilter.YFilter { return ipv6Dhcpv6DProxyIidReference.YFilter }

func (ipv6Dhcpv6DProxyIidReference *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceIdReferences_Ipv6Dhcpv6DProxyIidReference) SetFilter(yf yfilter.YFilter) { ipv6Dhcpv6DProxyIidReference.YFilter = yf }

func (ipv6Dhcpv6DProxyIidReference *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceIdReferences_Ipv6Dhcpv6DProxyIidReference) GetGoName(yname string) string {
    if yname == "proxy-iid-interface-name" { return "ProxyIidInterfaceName" }
    if yname == "proxy-interface-id" { return "ProxyInterfaceId" }
    return ""
}

func (ipv6Dhcpv6DProxyIidReference *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceIdReferences_Ipv6Dhcpv6DProxyIidReference) GetSegmentPath() string {
    return "ipv6-dhcpv6d-proxy-iid-reference"
}

func (ipv6Dhcpv6DProxyIidReference *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceIdReferences_Ipv6Dhcpv6DProxyIidReference) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv6Dhcpv6DProxyIidReference *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceIdReferences_Ipv6Dhcpv6DProxyIidReference) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv6Dhcpv6DProxyIidReference *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceIdReferences_Ipv6Dhcpv6DProxyIidReference) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["proxy-iid-interface-name"] = ipv6Dhcpv6DProxyIidReference.ProxyIidInterfaceName
    leafs["proxy-interface-id"] = ipv6Dhcpv6DProxyIidReference.ProxyInterfaceId
    return leafs
}

func (ipv6Dhcpv6DProxyIidReference *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceIdReferences_Ipv6Dhcpv6DProxyIidReference) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6Dhcpv6DProxyIidReference *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceIdReferences_Ipv6Dhcpv6DProxyIidReference) GetYangName() string { return "ipv6-dhcpv6d-proxy-iid-reference" }

func (ipv6Dhcpv6DProxyIidReference *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceIdReferences_Ipv6Dhcpv6DProxyIidReference) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6Dhcpv6DProxyIidReference *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceIdReferences_Ipv6Dhcpv6DProxyIidReference) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6Dhcpv6DProxyIidReference *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceIdReferences_Ipv6Dhcpv6DProxyIidReference) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6Dhcpv6DProxyIidReference *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceIdReferences_Ipv6Dhcpv6DProxyIidReference) SetParent(parent types.Entity) { ipv6Dhcpv6DProxyIidReference.parent = parent }

func (ipv6Dhcpv6DProxyIidReference *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceIdReferences_Ipv6Dhcpv6DProxyIidReference) GetParent() types.Entity { return ipv6Dhcpv6DProxyIidReference.parent }

func (ipv6Dhcpv6DProxyIidReference *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceIdReferences_Ipv6Dhcpv6DProxyIidReference) GetParentYangName() string { return "interface-id-references" }

// Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_VrfReferences
// VRF references
type Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_VrfReferences struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ipv6 dhcpv6d proxy vrf reference. The type is slice of
    // Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_VrfReferences_Ipv6Dhcpv6DProxyVrfReference.
    Ipv6Dhcpv6DProxyVrfReference []Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_VrfReferences_Ipv6Dhcpv6DProxyVrfReference
}

func (vrfReferences *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_VrfReferences) GetFilter() yfilter.YFilter { return vrfReferences.YFilter }

func (vrfReferences *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_VrfReferences) SetFilter(yf yfilter.YFilter) { vrfReferences.YFilter = yf }

func (vrfReferences *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_VrfReferences) GetGoName(yname string) string {
    if yname == "ipv6-dhcpv6d-proxy-vrf-reference" { return "Ipv6Dhcpv6DProxyVrfReference" }
    return ""
}

func (vrfReferences *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_VrfReferences) GetSegmentPath() string {
    return "vrf-references"
}

func (vrfReferences *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_VrfReferences) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv6-dhcpv6d-proxy-vrf-reference" {
        for _, c := range vrfReferences.Ipv6Dhcpv6DProxyVrfReference {
            if vrfReferences.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_VrfReferences_Ipv6Dhcpv6DProxyVrfReference{}
        vrfReferences.Ipv6Dhcpv6DProxyVrfReference = append(vrfReferences.Ipv6Dhcpv6DProxyVrfReference, child)
        return &vrfReferences.Ipv6Dhcpv6DProxyVrfReference[len(vrfReferences.Ipv6Dhcpv6DProxyVrfReference)-1]
    }
    return nil
}

func (vrfReferences *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_VrfReferences) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vrfReferences.Ipv6Dhcpv6DProxyVrfReference {
        children[vrfReferences.Ipv6Dhcpv6DProxyVrfReference[i].GetSegmentPath()] = &vrfReferences.Ipv6Dhcpv6DProxyVrfReference[i]
    }
    return children
}

func (vrfReferences *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_VrfReferences) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vrfReferences *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_VrfReferences) GetBundleName() string { return "cisco_ios_xr" }

func (vrfReferences *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_VrfReferences) GetYangName() string { return "vrf-references" }

func (vrfReferences *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_VrfReferences) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfReferences *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_VrfReferences) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfReferences *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_VrfReferences) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfReferences *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_VrfReferences) SetParent(parent types.Entity) { vrfReferences.parent = parent }

func (vrfReferences *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_VrfReferences) GetParent() types.Entity { return vrfReferences.parent }

func (vrfReferences *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_VrfReferences) GetParentYangName() string { return "info" }

// Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_VrfReferences_Ipv6Dhcpv6DProxyVrfReference
// ipv6 dhcpv6d proxy vrf reference
type Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_VrfReferences_Ipv6Dhcpv6DProxyVrfReference struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF name. The type is string with length: 0..33.
    ProxyReferenceVrfName interface{}
}

func (ipv6Dhcpv6DProxyVrfReference *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_VrfReferences_Ipv6Dhcpv6DProxyVrfReference) GetFilter() yfilter.YFilter { return ipv6Dhcpv6DProxyVrfReference.YFilter }

func (ipv6Dhcpv6DProxyVrfReference *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_VrfReferences_Ipv6Dhcpv6DProxyVrfReference) SetFilter(yf yfilter.YFilter) { ipv6Dhcpv6DProxyVrfReference.YFilter = yf }

func (ipv6Dhcpv6DProxyVrfReference *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_VrfReferences_Ipv6Dhcpv6DProxyVrfReference) GetGoName(yname string) string {
    if yname == "proxy-reference-vrf-name" { return "ProxyReferenceVrfName" }
    return ""
}

func (ipv6Dhcpv6DProxyVrfReference *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_VrfReferences_Ipv6Dhcpv6DProxyVrfReference) GetSegmentPath() string {
    return "ipv6-dhcpv6d-proxy-vrf-reference"
}

func (ipv6Dhcpv6DProxyVrfReference *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_VrfReferences_Ipv6Dhcpv6DProxyVrfReference) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv6Dhcpv6DProxyVrfReference *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_VrfReferences_Ipv6Dhcpv6DProxyVrfReference) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv6Dhcpv6DProxyVrfReference *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_VrfReferences_Ipv6Dhcpv6DProxyVrfReference) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["proxy-reference-vrf-name"] = ipv6Dhcpv6DProxyVrfReference.ProxyReferenceVrfName
    return leafs
}

func (ipv6Dhcpv6DProxyVrfReference *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_VrfReferences_Ipv6Dhcpv6DProxyVrfReference) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6Dhcpv6DProxyVrfReference *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_VrfReferences_Ipv6Dhcpv6DProxyVrfReference) GetYangName() string { return "ipv6-dhcpv6d-proxy-vrf-reference" }

func (ipv6Dhcpv6DProxyVrfReference *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_VrfReferences_Ipv6Dhcpv6DProxyVrfReference) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6Dhcpv6DProxyVrfReference *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_VrfReferences_Ipv6Dhcpv6DProxyVrfReference) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6Dhcpv6DProxyVrfReference *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_VrfReferences_Ipv6Dhcpv6DProxyVrfReference) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6Dhcpv6DProxyVrfReference *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_VrfReferences_Ipv6Dhcpv6DProxyVrfReference) SetParent(parent types.Entity) { ipv6Dhcpv6DProxyVrfReference.parent = parent }

func (ipv6Dhcpv6DProxyVrfReference *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_VrfReferences_Ipv6Dhcpv6DProxyVrfReference) GetParent() types.Entity { return ipv6Dhcpv6DProxyVrfReference.parent }

func (ipv6Dhcpv6DProxyVrfReference *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_VrfReferences_Ipv6Dhcpv6DProxyVrfReference) GetParentYangName() string { return "vrf-references" }

// Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceReferences
// Interface references
type Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceReferences struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ipv6 dhcpv6d proxy interface reference. The type is slice of
    // Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceReferences_Ipv6Dhcpv6DProxyInterfaceReference.
    Ipv6Dhcpv6DProxyInterfaceReference []Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceReferences_Ipv6Dhcpv6DProxyInterfaceReference
}

func (interfaceReferences *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceReferences) GetFilter() yfilter.YFilter { return interfaceReferences.YFilter }

func (interfaceReferences *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceReferences) SetFilter(yf yfilter.YFilter) { interfaceReferences.YFilter = yf }

func (interfaceReferences *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceReferences) GetGoName(yname string) string {
    if yname == "ipv6-dhcpv6d-proxy-interface-reference" { return "Ipv6Dhcpv6DProxyInterfaceReference" }
    return ""
}

func (interfaceReferences *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceReferences) GetSegmentPath() string {
    return "interface-references"
}

func (interfaceReferences *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceReferences) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv6-dhcpv6d-proxy-interface-reference" {
        for _, c := range interfaceReferences.Ipv6Dhcpv6DProxyInterfaceReference {
            if interfaceReferences.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceReferences_Ipv6Dhcpv6DProxyInterfaceReference{}
        interfaceReferences.Ipv6Dhcpv6DProxyInterfaceReference = append(interfaceReferences.Ipv6Dhcpv6DProxyInterfaceReference, child)
        return &interfaceReferences.Ipv6Dhcpv6DProxyInterfaceReference[len(interfaceReferences.Ipv6Dhcpv6DProxyInterfaceReference)-1]
    }
    return nil
}

func (interfaceReferences *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceReferences) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaceReferences.Ipv6Dhcpv6DProxyInterfaceReference {
        children[interfaceReferences.Ipv6Dhcpv6DProxyInterfaceReference[i].GetSegmentPath()] = &interfaceReferences.Ipv6Dhcpv6DProxyInterfaceReference[i]
    }
    return children
}

func (interfaceReferences *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceReferences) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceReferences *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceReferences) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceReferences *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceReferences) GetYangName() string { return "interface-references" }

func (interfaceReferences *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceReferences) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceReferences *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceReferences) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceReferences *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceReferences) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceReferences *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceReferences) SetParent(parent types.Entity) { interfaceReferences.parent = parent }

func (interfaceReferences *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceReferences) GetParent() types.Entity { return interfaceReferences.parent }

func (interfaceReferences *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceReferences) GetParentYangName() string { return "info" }

// Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceReferences_Ipv6Dhcpv6DProxyInterfaceReference
// ipv6 dhcpv6d proxy interface reference
type Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceReferences_Ipv6Dhcpv6DProxyInterfaceReference struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface name. The type is string with length: 0..65.
    ProxyReferenceInterfaceName interface{}
}

func (ipv6Dhcpv6DProxyInterfaceReference *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceReferences_Ipv6Dhcpv6DProxyInterfaceReference) GetFilter() yfilter.YFilter { return ipv6Dhcpv6DProxyInterfaceReference.YFilter }

func (ipv6Dhcpv6DProxyInterfaceReference *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceReferences_Ipv6Dhcpv6DProxyInterfaceReference) SetFilter(yf yfilter.YFilter) { ipv6Dhcpv6DProxyInterfaceReference.YFilter = yf }

func (ipv6Dhcpv6DProxyInterfaceReference *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceReferences_Ipv6Dhcpv6DProxyInterfaceReference) GetGoName(yname string) string {
    if yname == "proxy-reference-interface-name" { return "ProxyReferenceInterfaceName" }
    return ""
}

func (ipv6Dhcpv6DProxyInterfaceReference *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceReferences_Ipv6Dhcpv6DProxyInterfaceReference) GetSegmentPath() string {
    return "ipv6-dhcpv6d-proxy-interface-reference"
}

func (ipv6Dhcpv6DProxyInterfaceReference *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceReferences_Ipv6Dhcpv6DProxyInterfaceReference) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv6Dhcpv6DProxyInterfaceReference *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceReferences_Ipv6Dhcpv6DProxyInterfaceReference) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv6Dhcpv6DProxyInterfaceReference *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceReferences_Ipv6Dhcpv6DProxyInterfaceReference) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["proxy-reference-interface-name"] = ipv6Dhcpv6DProxyInterfaceReference.ProxyReferenceInterfaceName
    return leafs
}

func (ipv6Dhcpv6DProxyInterfaceReference *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceReferences_Ipv6Dhcpv6DProxyInterfaceReference) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6Dhcpv6DProxyInterfaceReference *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceReferences_Ipv6Dhcpv6DProxyInterfaceReference) GetYangName() string { return "ipv6-dhcpv6d-proxy-interface-reference" }

func (ipv6Dhcpv6DProxyInterfaceReference *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceReferences_Ipv6Dhcpv6DProxyInterfaceReference) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6Dhcpv6DProxyInterfaceReference *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceReferences_Ipv6Dhcpv6DProxyInterfaceReference) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6Dhcpv6DProxyInterfaceReference *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceReferences_Ipv6Dhcpv6DProxyInterfaceReference) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6Dhcpv6DProxyInterfaceReference *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceReferences_Ipv6Dhcpv6DProxyInterfaceReference) SetParent(parent types.Entity) { ipv6Dhcpv6DProxyInterfaceReference.parent = parent }

func (ipv6Dhcpv6DProxyInterfaceReference *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceReferences_Ipv6Dhcpv6DProxyInterfaceReference) GetParent() types.Entity { return ipv6Dhcpv6DProxyInterfaceReference.parent }

func (ipv6Dhcpv6DProxyInterfaceReference *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceReferences_Ipv6Dhcpv6DProxyInterfaceReference) GetParentYangName() string { return "interface-references" }

// Dhcpv6_Nodes_Node_Proxy_Interfaces
// DHCPV6 proxy interface
type Dhcpv6_Nodes_Node_Proxy_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 DHCP proxy interface. The type is slice of
    // Dhcpv6_Nodes_Node_Proxy_Interfaces_Interface.
    Interface []Dhcpv6_Nodes_Node_Proxy_Interfaces_Interface
}

func (interfaces *Dhcpv6_Nodes_Node_Proxy_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *Dhcpv6_Nodes_Node_Proxy_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *Dhcpv6_Nodes_Node_Proxy_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *Dhcpv6_Nodes_Node_Proxy_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *Dhcpv6_Nodes_Node_Proxy_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dhcpv6_Nodes_Node_Proxy_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *Dhcpv6_Nodes_Node_Proxy_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *Dhcpv6_Nodes_Node_Proxy_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *Dhcpv6_Nodes_Node_Proxy_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *Dhcpv6_Nodes_Node_Proxy_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *Dhcpv6_Nodes_Node_Proxy_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *Dhcpv6_Nodes_Node_Proxy_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *Dhcpv6_Nodes_Node_Proxy_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *Dhcpv6_Nodes_Node_Proxy_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *Dhcpv6_Nodes_Node_Proxy_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *Dhcpv6_Nodes_Node_Proxy_Interfaces) GetParentYangName() string { return "proxy" }

// Dhcpv6_Nodes_Node_Proxy_Interfaces_Interface
// IPv6 DHCP proxy interface
type Dhcpv6_Nodes_Node_Proxy_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    InterfaceName interface{}

    // VRF name. The type is string with length: 0..33.
    ProxyVrfName interface{}

    // Mode of interface. The type is BagDhcpv6DSubMode.
    ProxyInterfaceMode interface{}

    // Is interface ambiguous. The type is interface{} with range: 0..4294967295.
    IsProxyInterfaceAmbiguous interface{}

    // Name of profile attached to the interface. The type is string with length:
    // 0..65.
    ProxyInterfaceProfileName interface{}

    // Lease limit type on interface. The type is LeaseLimit.
    ProxyInterfaceLeaseLimitType interface{}

    // Lease limit count on interface. The type is interface{} with range:
    // 0..4294967295.
    ProxyInterfaceLeaseLimits interface{}

    // DHCPv6 Interface SRG role. The type is BagDhcpv6DIntfSrgRole.
    SrgRole interface{}

    // DHCPv6 Interface SERG role. The type is BagDhcpv6DIntfSergRole.
    SergRole interface{}

    // Mac Throttle Status. The type is bool.
    MacThrottle interface{}

    // SRG VRF name. The type is string with length: 0..33.
    SrgVrfName interface{}

    // SRG P2P Status. The type is bool.
    Srgp2P interface{}
}

func (self *Dhcpv6_Nodes_Node_Proxy_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Dhcpv6_Nodes_Node_Proxy_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Dhcpv6_Nodes_Node_Proxy_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "proxy-vrf-name" { return "ProxyVrfName" }
    if yname == "proxy-interface-mode" { return "ProxyInterfaceMode" }
    if yname == "is-proxy-interface-ambiguous" { return "IsProxyInterfaceAmbiguous" }
    if yname == "proxy-interface-profile-name" { return "ProxyInterfaceProfileName" }
    if yname == "proxy-interface-lease-limit-type" { return "ProxyInterfaceLeaseLimitType" }
    if yname == "proxy-interface-lease-limits" { return "ProxyInterfaceLeaseLimits" }
    if yname == "srg-role" { return "SrgRole" }
    if yname == "serg-role" { return "SergRole" }
    if yname == "mac-throttle" { return "MacThrottle" }
    if yname == "srg-vrf-name" { return "SrgVrfName" }
    if yname == "srgp2p" { return "Srgp2P" }
    return ""
}

func (self *Dhcpv6_Nodes_Node_Proxy_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *Dhcpv6_Nodes_Node_Proxy_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *Dhcpv6_Nodes_Node_Proxy_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *Dhcpv6_Nodes_Node_Proxy_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    leafs["proxy-vrf-name"] = self.ProxyVrfName
    leafs["proxy-interface-mode"] = self.ProxyInterfaceMode
    leafs["is-proxy-interface-ambiguous"] = self.IsProxyInterfaceAmbiguous
    leafs["proxy-interface-profile-name"] = self.ProxyInterfaceProfileName
    leafs["proxy-interface-lease-limit-type"] = self.ProxyInterfaceLeaseLimitType
    leafs["proxy-interface-lease-limits"] = self.ProxyInterfaceLeaseLimits
    leafs["srg-role"] = self.SrgRole
    leafs["serg-role"] = self.SergRole
    leafs["mac-throttle"] = self.MacThrottle
    leafs["srg-vrf-name"] = self.SrgVrfName
    leafs["srgp2p"] = self.Srgp2P
    return leafs
}

func (self *Dhcpv6_Nodes_Node_Proxy_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *Dhcpv6_Nodes_Node_Proxy_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *Dhcpv6_Nodes_Node_Proxy_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Dhcpv6_Nodes_Node_Proxy_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Dhcpv6_Nodes_Node_Proxy_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Dhcpv6_Nodes_Node_Proxy_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *Dhcpv6_Nodes_Node_Proxy_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *Dhcpv6_Nodes_Node_Proxy_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// Dhcpv6_Nodes_Node_Proxy_Statistics
// DHCPv6 proxy statistics
type Dhcpv6_Nodes_Node_Proxy_Statistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ipv6 dhcpv6d proxy stat. The type is slice of
    // Dhcpv6_Nodes_Node_Proxy_Statistics_Ipv6Dhcpv6DProxyStat.
    Ipv6Dhcpv6DProxyStat []Dhcpv6_Nodes_Node_Proxy_Statistics_Ipv6Dhcpv6DProxyStat
}

func (statistics *Dhcpv6_Nodes_Node_Proxy_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *Dhcpv6_Nodes_Node_Proxy_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *Dhcpv6_Nodes_Node_Proxy_Statistics) GetGoName(yname string) string {
    if yname == "ipv6-dhcpv6d-proxy-stat" { return "Ipv6Dhcpv6DProxyStat" }
    return ""
}

func (statistics *Dhcpv6_Nodes_Node_Proxy_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *Dhcpv6_Nodes_Node_Proxy_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv6-dhcpv6d-proxy-stat" {
        for _, c := range statistics.Ipv6Dhcpv6DProxyStat {
            if statistics.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dhcpv6_Nodes_Node_Proxy_Statistics_Ipv6Dhcpv6DProxyStat{}
        statistics.Ipv6Dhcpv6DProxyStat = append(statistics.Ipv6Dhcpv6DProxyStat, child)
        return &statistics.Ipv6Dhcpv6DProxyStat[len(statistics.Ipv6Dhcpv6DProxyStat)-1]
    }
    return nil
}

func (statistics *Dhcpv6_Nodes_Node_Proxy_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range statistics.Ipv6Dhcpv6DProxyStat {
        children[statistics.Ipv6Dhcpv6DProxyStat[i].GetSegmentPath()] = &statistics.Ipv6Dhcpv6DProxyStat[i]
    }
    return children
}

func (statistics *Dhcpv6_Nodes_Node_Proxy_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (statistics *Dhcpv6_Nodes_Node_Proxy_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *Dhcpv6_Nodes_Node_Proxy_Statistics) GetYangName() string { return "statistics" }

func (statistics *Dhcpv6_Nodes_Node_Proxy_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *Dhcpv6_Nodes_Node_Proxy_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *Dhcpv6_Nodes_Node_Proxy_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *Dhcpv6_Nodes_Node_Proxy_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *Dhcpv6_Nodes_Node_Proxy_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *Dhcpv6_Nodes_Node_Proxy_Statistics) GetParentYangName() string { return "proxy" }

// Dhcpv6_Nodes_Node_Proxy_Statistics_Ipv6Dhcpv6DProxyStat
// ipv6 dhcpv6d proxy stat
type Dhcpv6_Nodes_Node_Proxy_Statistics_Ipv6Dhcpv6DProxyStat struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DHCPv6 L3 VRF name. The type is string with length: 0..33.
    VrfName interface{}

    // Proxy statistics.
    Statistics Dhcpv6_Nodes_Node_Proxy_Statistics_Ipv6Dhcpv6DProxyStat_Statistics
}

func (ipv6Dhcpv6DProxyStat *Dhcpv6_Nodes_Node_Proxy_Statistics_Ipv6Dhcpv6DProxyStat) GetFilter() yfilter.YFilter { return ipv6Dhcpv6DProxyStat.YFilter }

func (ipv6Dhcpv6DProxyStat *Dhcpv6_Nodes_Node_Proxy_Statistics_Ipv6Dhcpv6DProxyStat) SetFilter(yf yfilter.YFilter) { ipv6Dhcpv6DProxyStat.YFilter = yf }

func (ipv6Dhcpv6DProxyStat *Dhcpv6_Nodes_Node_Proxy_Statistics_Ipv6Dhcpv6DProxyStat) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "statistics" { return "Statistics" }
    return ""
}

func (ipv6Dhcpv6DProxyStat *Dhcpv6_Nodes_Node_Proxy_Statistics_Ipv6Dhcpv6DProxyStat) GetSegmentPath() string {
    return "ipv6-dhcpv6d-proxy-stat"
}

func (ipv6Dhcpv6DProxyStat *Dhcpv6_Nodes_Node_Proxy_Statistics_Ipv6Dhcpv6DProxyStat) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "statistics" {
        return &ipv6Dhcpv6DProxyStat.Statistics
    }
    return nil
}

func (ipv6Dhcpv6DProxyStat *Dhcpv6_Nodes_Node_Proxy_Statistics_Ipv6Dhcpv6DProxyStat) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["statistics"] = &ipv6Dhcpv6DProxyStat.Statistics
    return children
}

func (ipv6Dhcpv6DProxyStat *Dhcpv6_Nodes_Node_Proxy_Statistics_Ipv6Dhcpv6DProxyStat) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = ipv6Dhcpv6DProxyStat.VrfName
    return leafs
}

func (ipv6Dhcpv6DProxyStat *Dhcpv6_Nodes_Node_Proxy_Statistics_Ipv6Dhcpv6DProxyStat) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6Dhcpv6DProxyStat *Dhcpv6_Nodes_Node_Proxy_Statistics_Ipv6Dhcpv6DProxyStat) GetYangName() string { return "ipv6-dhcpv6d-proxy-stat" }

func (ipv6Dhcpv6DProxyStat *Dhcpv6_Nodes_Node_Proxy_Statistics_Ipv6Dhcpv6DProxyStat) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6Dhcpv6DProxyStat *Dhcpv6_Nodes_Node_Proxy_Statistics_Ipv6Dhcpv6DProxyStat) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6Dhcpv6DProxyStat *Dhcpv6_Nodes_Node_Proxy_Statistics_Ipv6Dhcpv6DProxyStat) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6Dhcpv6DProxyStat *Dhcpv6_Nodes_Node_Proxy_Statistics_Ipv6Dhcpv6DProxyStat) SetParent(parent types.Entity) { ipv6Dhcpv6DProxyStat.parent = parent }

func (ipv6Dhcpv6DProxyStat *Dhcpv6_Nodes_Node_Proxy_Statistics_Ipv6Dhcpv6DProxyStat) GetParent() types.Entity { return ipv6Dhcpv6DProxyStat.parent }

func (ipv6Dhcpv6DProxyStat *Dhcpv6_Nodes_Node_Proxy_Statistics_Ipv6Dhcpv6DProxyStat) GetParentYangName() string { return "statistics" }

// Dhcpv6_Nodes_Node_Proxy_Statistics_Ipv6Dhcpv6DProxyStat_Statistics
// Proxy statistics
type Dhcpv6_Nodes_Node_Proxy_Statistics_Ipv6Dhcpv6DProxyStat_Statistics struct {
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

func (statistics *Dhcpv6_Nodes_Node_Proxy_Statistics_Ipv6Dhcpv6DProxyStat_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *Dhcpv6_Nodes_Node_Proxy_Statistics_Ipv6Dhcpv6DProxyStat_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *Dhcpv6_Nodes_Node_Proxy_Statistics_Ipv6Dhcpv6DProxyStat_Statistics) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (statistics *Dhcpv6_Nodes_Node_Proxy_Statistics_Ipv6Dhcpv6DProxyStat_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *Dhcpv6_Nodes_Node_Proxy_Statistics_Ipv6Dhcpv6DProxyStat_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (statistics *Dhcpv6_Nodes_Node_Proxy_Statistics_Ipv6Dhcpv6DProxyStat_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (statistics *Dhcpv6_Nodes_Node_Proxy_Statistics_Ipv6Dhcpv6DProxyStat_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = statistics.ReceivedPackets
    leafs["transmitted-packets"] = statistics.TransmittedPackets
    leafs["dropped-packets"] = statistics.DroppedPackets
    return leafs
}

func (statistics *Dhcpv6_Nodes_Node_Proxy_Statistics_Ipv6Dhcpv6DProxyStat_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *Dhcpv6_Nodes_Node_Proxy_Statistics_Ipv6Dhcpv6DProxyStat_Statistics) GetYangName() string { return "statistics" }

func (statistics *Dhcpv6_Nodes_Node_Proxy_Statistics_Ipv6Dhcpv6DProxyStat_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *Dhcpv6_Nodes_Node_Proxy_Statistics_Ipv6Dhcpv6DProxyStat_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *Dhcpv6_Nodes_Node_Proxy_Statistics_Ipv6Dhcpv6DProxyStat_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *Dhcpv6_Nodes_Node_Proxy_Statistics_Ipv6Dhcpv6DProxyStat_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *Dhcpv6_Nodes_Node_Proxy_Statistics_Ipv6Dhcpv6DProxyStat_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *Dhcpv6_Nodes_Node_Proxy_Statistics_Ipv6Dhcpv6DProxyStat_Statistics) GetParentYangName() string { return "ipv6-dhcpv6d-proxy-stat" }

// Dhcpv6_Nodes_Node_Proxy_Binding
// DHCPV6 proxy bindings
type Dhcpv6_Nodes_Node_Proxy_Binding struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DHCPV6 proxy client bindings.
    Clients Dhcpv6_Nodes_Node_Proxy_Binding_Clients

    // DHCPV6 proxy binding summary.
    Summary Dhcpv6_Nodes_Node_Proxy_Binding_Summary
}

func (binding *Dhcpv6_Nodes_Node_Proxy_Binding) GetFilter() yfilter.YFilter { return binding.YFilter }

func (binding *Dhcpv6_Nodes_Node_Proxy_Binding) SetFilter(yf yfilter.YFilter) { binding.YFilter = yf }

func (binding *Dhcpv6_Nodes_Node_Proxy_Binding) GetGoName(yname string) string {
    if yname == "clients" { return "Clients" }
    if yname == "summary" { return "Summary" }
    return ""
}

func (binding *Dhcpv6_Nodes_Node_Proxy_Binding) GetSegmentPath() string {
    return "binding"
}

func (binding *Dhcpv6_Nodes_Node_Proxy_Binding) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "clients" {
        return &binding.Clients
    }
    if childYangName == "summary" {
        return &binding.Summary
    }
    return nil
}

func (binding *Dhcpv6_Nodes_Node_Proxy_Binding) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["clients"] = &binding.Clients
    children["summary"] = &binding.Summary
    return children
}

func (binding *Dhcpv6_Nodes_Node_Proxy_Binding) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (binding *Dhcpv6_Nodes_Node_Proxy_Binding) GetBundleName() string { return "cisco_ios_xr" }

func (binding *Dhcpv6_Nodes_Node_Proxy_Binding) GetYangName() string { return "binding" }

func (binding *Dhcpv6_Nodes_Node_Proxy_Binding) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (binding *Dhcpv6_Nodes_Node_Proxy_Binding) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (binding *Dhcpv6_Nodes_Node_Proxy_Binding) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (binding *Dhcpv6_Nodes_Node_Proxy_Binding) SetParent(parent types.Entity) { binding.parent = parent }

func (binding *Dhcpv6_Nodes_Node_Proxy_Binding) GetParent() types.Entity { return binding.parent }

func (binding *Dhcpv6_Nodes_Node_Proxy_Binding) GetParentYangName() string { return "proxy" }

// Dhcpv6_Nodes_Node_Proxy_Binding_Clients
// DHCPV6 proxy client bindings
type Dhcpv6_Nodes_Node_Proxy_Binding_Clients struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Single DHCPV6 proxy binding. The type is slice of
    // Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client.
    Client []Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client
}

func (clients *Dhcpv6_Nodes_Node_Proxy_Binding_Clients) GetFilter() yfilter.YFilter { return clients.YFilter }

func (clients *Dhcpv6_Nodes_Node_Proxy_Binding_Clients) SetFilter(yf yfilter.YFilter) { clients.YFilter = yf }

func (clients *Dhcpv6_Nodes_Node_Proxy_Binding_Clients) GetGoName(yname string) string {
    if yname == "client" { return "Client" }
    return ""
}

func (clients *Dhcpv6_Nodes_Node_Proxy_Binding_Clients) GetSegmentPath() string {
    return "clients"
}

func (clients *Dhcpv6_Nodes_Node_Proxy_Binding_Clients) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "client" {
        for _, c := range clients.Client {
            if clients.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client{}
        clients.Client = append(clients.Client, child)
        return &clients.Client[len(clients.Client)-1]
    }
    return nil
}

func (clients *Dhcpv6_Nodes_Node_Proxy_Binding_Clients) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range clients.Client {
        children[clients.Client[i].GetSegmentPath()] = &clients.Client[i]
    }
    return children
}

func (clients *Dhcpv6_Nodes_Node_Proxy_Binding_Clients) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (clients *Dhcpv6_Nodes_Node_Proxy_Binding_Clients) GetBundleName() string { return "cisco_ios_xr" }

func (clients *Dhcpv6_Nodes_Node_Proxy_Binding_Clients) GetYangName() string { return "clients" }

func (clients *Dhcpv6_Nodes_Node_Proxy_Binding_Clients) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clients *Dhcpv6_Nodes_Node_Proxy_Binding_Clients) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clients *Dhcpv6_Nodes_Node_Proxy_Binding_Clients) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clients *Dhcpv6_Nodes_Node_Proxy_Binding_Clients) SetParent(parent types.Entity) { clients.parent = parent }

func (clients *Dhcpv6_Nodes_Node_Proxy_Binding_Clients) GetParent() types.Entity { return clients.parent }

func (clients *Dhcpv6_Nodes_Node_Proxy_Binding_Clients) GetParentYangName() string { return "binding" }

// Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client
// Single DHCPV6 proxy binding
type Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Client ID. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    ClientId interface{}

    // Client DUID. The type is string.
    Duid interface{}

    // DHCPV6 client flag. The type is interface{} with range: 0..4294967295.
    ClientFlag interface{}

    // DHCPV6 subscriber label. The type is interface{} with range: 0..4294967295.
    SubscriberLabel interface{}

    // DHCPVV6 client/subscriber VRF name. The type is string with length: 0..33.
    VrfName interface{}

    // Client MAC address. The type is string.
    MacAddress interface{}

    // Number of ia_id/pd. The type is interface{} with range: 0..4294967295.
    IaIdPDs interface{}

    // DHCPV6 access interface to client. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // DHCPV6 access VRF name to client. The type is string with length: 0..33.
    AccessVrfName interface{}

    // DHCPV6 VLAN tag count. The type is interface{} with range: 0..255.
    ProxyBindingTags interface{}

    // DHCPV6 VLAN Outer VLAN. The type is interface{} with range: 0..4294967295.
    ProxyBindingOuterTag interface{}

    // DHCPV6 VLAN Inner VLAN. The type is interface{} with range: 0..4294967295.
    ProxyBindingInnerTag interface{}

    // DHCPV6 class name. The type is string with length: 0..64.
    ClassName interface{}

    // DHCPV6 pool name. The type is string with length: 0..64.
    PoolName interface{}

    // DHCPV6 received Remote ID. The type is string with length: 0..771.
    RxRemoteId interface{}

    // DHCPV6 transmitted Remote ID. The type is string with length: 0..771.
    TxRemoteId interface{}

    // DHCPV6 received Interface ID. The type is string with length: 0..771.
    RxInterfaceId interface{}

    // DHCPV6 transmitted Interface ID. The type is string with length: 0..771.
    TxInterfaceId interface{}

    // DHCPV6 server IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    ServerIpv6Address interface{}

    // DHCPV6 profile name. The type is string with length: 0..65.
    ProfileName interface{}

    // DHCPV6 framed ipv6 addess used by ND. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    FramedIpv6Prefix interface{}

    // DHCPV6 framed ipv6 prefix length used by ND. The type is interface{} with
    // range: 0..255.
    FramedPrefixLength interface{}

    // Is true if DHCP next renew from client will be NAK'd. The type is bool.
    IsNakNextRenew interface{}

    // DHCPV6 SRG state. The type is interface{} with range: 0..4294967295.
    SrgState interface{}

    // DHCPV6 SRG Intf Role. The type is interface{} with range: 0..4294967295.
    SrgIntfRole interface{}

    // SRG P2P Status. The type is bool.
    Srgp2P interface{}

    // DHCPV6 SRG VRF NAME. The type is string with length: 0..33.
    SrgVrfName interface{}

    // DHCPV6 SERG state. The type is interface{} with range: 0..4294967295.
    SergState interface{}

    // DHCPV6 SERG Intf Role. The type is interface{} with range: 0..4294967295.
    SergIntfRole interface{}

    // List of DHCPv6 IA_ID/PDs.
    IaIdPd Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd
}

func (client *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client) GetFilter() yfilter.YFilter { return client.YFilter }

func (client *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client) SetFilter(yf yfilter.YFilter) { client.YFilter = yf }

func (client *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client) GetGoName(yname string) string {
    if yname == "client-id" { return "ClientId" }
    if yname == "duid" { return "Duid" }
    if yname == "client-flag" { return "ClientFlag" }
    if yname == "subscriber-label" { return "SubscriberLabel" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "mac-address" { return "MacAddress" }
    if yname == "ia-id-p-ds" { return "IaIdPDs" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "access-vrf-name" { return "AccessVrfName" }
    if yname == "proxy-binding-tags" { return "ProxyBindingTags" }
    if yname == "proxy-binding-outer-tag" { return "ProxyBindingOuterTag" }
    if yname == "proxy-binding-inner-tag" { return "ProxyBindingInnerTag" }
    if yname == "class-name" { return "ClassName" }
    if yname == "pool-name" { return "PoolName" }
    if yname == "rx-remote-id" { return "RxRemoteId" }
    if yname == "tx-remote-id" { return "TxRemoteId" }
    if yname == "rx-interface-id" { return "RxInterfaceId" }
    if yname == "tx-interface-id" { return "TxInterfaceId" }
    if yname == "server-ipv6-address" { return "ServerIpv6Address" }
    if yname == "profile-name" { return "ProfileName" }
    if yname == "framed-ipv6-prefix" { return "FramedIpv6Prefix" }
    if yname == "framed-prefix-length" { return "FramedPrefixLength" }
    if yname == "is-nak-next-renew" { return "IsNakNextRenew" }
    if yname == "srg-state" { return "SrgState" }
    if yname == "srg-intf-role" { return "SrgIntfRole" }
    if yname == "srgp2p" { return "Srgp2P" }
    if yname == "srg-vrf-name" { return "SrgVrfName" }
    if yname == "serg-state" { return "SergState" }
    if yname == "serg-intf-role" { return "SergIntfRole" }
    if yname == "ia-id-pd" { return "IaIdPd" }
    return ""
}

func (client *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client) GetSegmentPath() string {
    return "client" + "[client-id='" + fmt.Sprintf("%v", client.ClientId) + "']"
}

func (client *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ia-id-pd" {
        return &client.IaIdPd
    }
    return nil
}

func (client *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ia-id-pd"] = &client.IaIdPd
    return children
}

func (client *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["client-id"] = client.ClientId
    leafs["duid"] = client.Duid
    leafs["client-flag"] = client.ClientFlag
    leafs["subscriber-label"] = client.SubscriberLabel
    leafs["vrf-name"] = client.VrfName
    leafs["mac-address"] = client.MacAddress
    leafs["ia-id-p-ds"] = client.IaIdPDs
    leafs["interface-name"] = client.InterfaceName
    leafs["access-vrf-name"] = client.AccessVrfName
    leafs["proxy-binding-tags"] = client.ProxyBindingTags
    leafs["proxy-binding-outer-tag"] = client.ProxyBindingOuterTag
    leafs["proxy-binding-inner-tag"] = client.ProxyBindingInnerTag
    leafs["class-name"] = client.ClassName
    leafs["pool-name"] = client.PoolName
    leafs["rx-remote-id"] = client.RxRemoteId
    leafs["tx-remote-id"] = client.TxRemoteId
    leafs["rx-interface-id"] = client.RxInterfaceId
    leafs["tx-interface-id"] = client.TxInterfaceId
    leafs["server-ipv6-address"] = client.ServerIpv6Address
    leafs["profile-name"] = client.ProfileName
    leafs["framed-ipv6-prefix"] = client.FramedIpv6Prefix
    leafs["framed-prefix-length"] = client.FramedPrefixLength
    leafs["is-nak-next-renew"] = client.IsNakNextRenew
    leafs["srg-state"] = client.SrgState
    leafs["srg-intf-role"] = client.SrgIntfRole
    leafs["srgp2p"] = client.Srgp2P
    leafs["srg-vrf-name"] = client.SrgVrfName
    leafs["serg-state"] = client.SergState
    leafs["serg-intf-role"] = client.SergIntfRole
    return leafs
}

func (client *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client) GetBundleName() string { return "cisco_ios_xr" }

func (client *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client) GetYangName() string { return "client" }

func (client *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (client *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (client *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (client *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client) SetParent(parent types.Entity) { client.parent = parent }

func (client *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client) GetParent() types.Entity { return client.parent }

func (client *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client) GetParentYangName() string { return "clients" }

// Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd
// List of DHCPv6 IA_ID/PDs
type Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // bag dhcpv6d ia id pd info. The type is slice of
    // Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo.
    BagDhcpv6DIaIdPdInfo []Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo
}

func (iaIdPd *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd) GetFilter() yfilter.YFilter { return iaIdPd.YFilter }

func (iaIdPd *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd) SetFilter(yf yfilter.YFilter) { iaIdPd.YFilter = yf }

func (iaIdPd *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd) GetGoName(yname string) string {
    if yname == "bag-dhcpv6d-ia-id-pd-info" { return "BagDhcpv6DIaIdPdInfo" }
    return ""
}

func (iaIdPd *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd) GetSegmentPath() string {
    return "ia-id-pd"
}

func (iaIdPd *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bag-dhcpv6d-ia-id-pd-info" {
        for _, c := range iaIdPd.BagDhcpv6DIaIdPdInfo {
            if iaIdPd.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo{}
        iaIdPd.BagDhcpv6DIaIdPdInfo = append(iaIdPd.BagDhcpv6DIaIdPdInfo, child)
        return &iaIdPd.BagDhcpv6DIaIdPdInfo[len(iaIdPd.BagDhcpv6DIaIdPdInfo)-1]
    }
    return nil
}

func (iaIdPd *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range iaIdPd.BagDhcpv6DIaIdPdInfo {
        children[iaIdPd.BagDhcpv6DIaIdPdInfo[i].GetSegmentPath()] = &iaIdPd.BagDhcpv6DIaIdPdInfo[i]
    }
    return children
}

func (iaIdPd *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (iaIdPd *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd) GetBundleName() string { return "cisco_ios_xr" }

func (iaIdPd *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd) GetYangName() string { return "ia-id-pd" }

func (iaIdPd *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (iaIdPd *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (iaIdPd *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (iaIdPd *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd) SetParent(parent types.Entity) { iaIdPd.parent = parent }

func (iaIdPd *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd) GetParent() types.Entity { return iaIdPd.parent }

func (iaIdPd *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd) GetParentYangName() string { return "client" }

// Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo
// bag dhcpv6d ia id pd info
type Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IA type. The type is BagDhcpv6DIaId.
    IaType interface{}

    // IA_ID of this IA. The type is interface{} with range: 0..4294967295.
    IaId interface{}

    // FSM Flag for this IA. The type is interface{} with range: 0..4294967295.
    Flags interface{}

    // Total address in this IA. The type is interface{} with range: 0..65535.
    TotalAddress interface{}

    // State. The type is BagDhcpv6DFsmState.
    State interface{}

    // List of addresses in this IA.
    Addresses Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses
}

func (bagDhcpv6DIaIdPdInfo *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo) GetFilter() yfilter.YFilter { return bagDhcpv6DIaIdPdInfo.YFilter }

func (bagDhcpv6DIaIdPdInfo *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo) SetFilter(yf yfilter.YFilter) { bagDhcpv6DIaIdPdInfo.YFilter = yf }

func (bagDhcpv6DIaIdPdInfo *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo) GetGoName(yname string) string {
    if yname == "ia-type" { return "IaType" }
    if yname == "ia-id" { return "IaId" }
    if yname == "flags" { return "Flags" }
    if yname == "total-address" { return "TotalAddress" }
    if yname == "state" { return "State" }
    if yname == "addresses" { return "Addresses" }
    return ""
}

func (bagDhcpv6DIaIdPdInfo *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo) GetSegmentPath() string {
    return "bag-dhcpv6d-ia-id-pd-info"
}

func (bagDhcpv6DIaIdPdInfo *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "addresses" {
        return &bagDhcpv6DIaIdPdInfo.Addresses
    }
    return nil
}

func (bagDhcpv6DIaIdPdInfo *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["addresses"] = &bagDhcpv6DIaIdPdInfo.Addresses
    return children
}

func (bagDhcpv6DIaIdPdInfo *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ia-type"] = bagDhcpv6DIaIdPdInfo.IaType
    leafs["ia-id"] = bagDhcpv6DIaIdPdInfo.IaId
    leafs["flags"] = bagDhcpv6DIaIdPdInfo.Flags
    leafs["total-address"] = bagDhcpv6DIaIdPdInfo.TotalAddress
    leafs["state"] = bagDhcpv6DIaIdPdInfo.State
    return leafs
}

func (bagDhcpv6DIaIdPdInfo *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo) GetBundleName() string { return "cisco_ios_xr" }

func (bagDhcpv6DIaIdPdInfo *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo) GetYangName() string { return "bag-dhcpv6d-ia-id-pd-info" }

func (bagDhcpv6DIaIdPdInfo *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bagDhcpv6DIaIdPdInfo *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bagDhcpv6DIaIdPdInfo *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bagDhcpv6DIaIdPdInfo *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo) SetParent(parent types.Entity) { bagDhcpv6DIaIdPdInfo.parent = parent }

func (bagDhcpv6DIaIdPdInfo *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo) GetParent() types.Entity { return bagDhcpv6DIaIdPdInfo.parent }

func (bagDhcpv6DIaIdPdInfo *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo) GetParentYangName() string { return "ia-id-pd" }

// Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses
// List of addresses in this IA
type Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // bag dhcpv6d addr attrb. The type is slice of
    // Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses_BagDhcpv6DAddrAttrb.
    BagDhcpv6DAddrAttrb []Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses_BagDhcpv6DAddrAttrb
}

func (addresses *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses) GetFilter() yfilter.YFilter { return addresses.YFilter }

func (addresses *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses) SetFilter(yf yfilter.YFilter) { addresses.YFilter = yf }

func (addresses *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses) GetGoName(yname string) string {
    if yname == "bag-dhcpv6d-addr-attrb" { return "BagDhcpv6DAddrAttrb" }
    return ""
}

func (addresses *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses) GetSegmentPath() string {
    return "addresses"
}

func (addresses *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bag-dhcpv6d-addr-attrb" {
        for _, c := range addresses.BagDhcpv6DAddrAttrb {
            if addresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses_BagDhcpv6DAddrAttrb{}
        addresses.BagDhcpv6DAddrAttrb = append(addresses.BagDhcpv6DAddrAttrb, child)
        return &addresses.BagDhcpv6DAddrAttrb[len(addresses.BagDhcpv6DAddrAttrb)-1]
    }
    return nil
}

func (addresses *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range addresses.BagDhcpv6DAddrAttrb {
        children[addresses.BagDhcpv6DAddrAttrb[i].GetSegmentPath()] = &addresses.BagDhcpv6DAddrAttrb[i]
    }
    return children
}

func (addresses *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (addresses *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses) GetBundleName() string { return "cisco_ios_xr" }

func (addresses *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses) GetYangName() string { return "addresses" }

func (addresses *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (addresses *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (addresses *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (addresses *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses) SetParent(parent types.Entity) { addresses.parent = parent }

func (addresses *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses) GetParent() types.Entity { return addresses.parent }

func (addresses *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses) GetParentYangName() string { return "bag-dhcpv6d-ia-id-pd-info" }

// Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses_BagDhcpv6DAddrAttrb
// bag dhcpv6d addr attrb
type Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses_BagDhcpv6DAddrAttrb struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 prefix. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Prefix interface{}

    // Prefix length. The type is interface{} with range: 0..255.
    PrefixLength interface{}

    // Lease time in seconds. The type is interface{} with range: 0..4294967295.
    // Units are second.
    LeaseTime interface{}

    // Remaining lease time in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    RemainingLeaseTime interface{}
}

func (bagDhcpv6DAddrAttrb *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses_BagDhcpv6DAddrAttrb) GetFilter() yfilter.YFilter { return bagDhcpv6DAddrAttrb.YFilter }

func (bagDhcpv6DAddrAttrb *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses_BagDhcpv6DAddrAttrb) SetFilter(yf yfilter.YFilter) { bagDhcpv6DAddrAttrb.YFilter = yf }

func (bagDhcpv6DAddrAttrb *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses_BagDhcpv6DAddrAttrb) GetGoName(yname string) string {
    if yname == "prefix" { return "Prefix" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "lease-time" { return "LeaseTime" }
    if yname == "remaining-lease-time" { return "RemainingLeaseTime" }
    return ""
}

func (bagDhcpv6DAddrAttrb *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses_BagDhcpv6DAddrAttrb) GetSegmentPath() string {
    return "bag-dhcpv6d-addr-attrb"
}

func (bagDhcpv6DAddrAttrb *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses_BagDhcpv6DAddrAttrb) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bagDhcpv6DAddrAttrb *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses_BagDhcpv6DAddrAttrb) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bagDhcpv6DAddrAttrb *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses_BagDhcpv6DAddrAttrb) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix"] = bagDhcpv6DAddrAttrb.Prefix
    leafs["prefix-length"] = bagDhcpv6DAddrAttrb.PrefixLength
    leafs["lease-time"] = bagDhcpv6DAddrAttrb.LeaseTime
    leafs["remaining-lease-time"] = bagDhcpv6DAddrAttrb.RemainingLeaseTime
    return leafs
}

func (bagDhcpv6DAddrAttrb *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses_BagDhcpv6DAddrAttrb) GetBundleName() string { return "cisco_ios_xr" }

func (bagDhcpv6DAddrAttrb *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses_BagDhcpv6DAddrAttrb) GetYangName() string { return "bag-dhcpv6d-addr-attrb" }

func (bagDhcpv6DAddrAttrb *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses_BagDhcpv6DAddrAttrb) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bagDhcpv6DAddrAttrb *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses_BagDhcpv6DAddrAttrb) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bagDhcpv6DAddrAttrb *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses_BagDhcpv6DAddrAttrb) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bagDhcpv6DAddrAttrb *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses_BagDhcpv6DAddrAttrb) SetParent(parent types.Entity) { bagDhcpv6DAddrAttrb.parent = parent }

func (bagDhcpv6DAddrAttrb *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses_BagDhcpv6DAddrAttrb) GetParent() types.Entity { return bagDhcpv6DAddrAttrb.parent }

func (bagDhcpv6DAddrAttrb *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses_BagDhcpv6DAddrAttrb) GetParentYangName() string { return "addresses" }

// Dhcpv6_Nodes_Node_Proxy_Binding_Summary
// DHCPV6 proxy binding summary
type Dhcpv6_Nodes_Node_Proxy_Binding_Summary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Total number of clients. The type is interface{} with range: 0..4294967295.
    Clients interface{}

    // IANA proxy binding summary.
    Iana Dhcpv6_Nodes_Node_Proxy_Binding_Summary_Iana

    // IAPD proxy binding summary.
    Iapd Dhcpv6_Nodes_Node_Proxy_Binding_Summary_Iapd
}

func (summary *Dhcpv6_Nodes_Node_Proxy_Binding_Summary) GetFilter() yfilter.YFilter { return summary.YFilter }

func (summary *Dhcpv6_Nodes_Node_Proxy_Binding_Summary) SetFilter(yf yfilter.YFilter) { summary.YFilter = yf }

func (summary *Dhcpv6_Nodes_Node_Proxy_Binding_Summary) GetGoName(yname string) string {
    if yname == "clients" { return "Clients" }
    if yname == "iana" { return "Iana" }
    if yname == "iapd" { return "Iapd" }
    return ""
}

func (summary *Dhcpv6_Nodes_Node_Proxy_Binding_Summary) GetSegmentPath() string {
    return "summary"
}

func (summary *Dhcpv6_Nodes_Node_Proxy_Binding_Summary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "iana" {
        return &summary.Iana
    }
    if childYangName == "iapd" {
        return &summary.Iapd
    }
    return nil
}

func (summary *Dhcpv6_Nodes_Node_Proxy_Binding_Summary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["iana"] = &summary.Iana
    children["iapd"] = &summary.Iapd
    return children
}

func (summary *Dhcpv6_Nodes_Node_Proxy_Binding_Summary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["clients"] = summary.Clients
    return leafs
}

func (summary *Dhcpv6_Nodes_Node_Proxy_Binding_Summary) GetBundleName() string { return "cisco_ios_xr" }

func (summary *Dhcpv6_Nodes_Node_Proxy_Binding_Summary) GetYangName() string { return "summary" }

func (summary *Dhcpv6_Nodes_Node_Proxy_Binding_Summary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (summary *Dhcpv6_Nodes_Node_Proxy_Binding_Summary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (summary *Dhcpv6_Nodes_Node_Proxy_Binding_Summary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (summary *Dhcpv6_Nodes_Node_Proxy_Binding_Summary) SetParent(parent types.Entity) { summary.parent = parent }

func (summary *Dhcpv6_Nodes_Node_Proxy_Binding_Summary) GetParent() types.Entity { return summary.parent }

func (summary *Dhcpv6_Nodes_Node_Proxy_Binding_Summary) GetParentYangName() string { return "binding" }

// Dhcpv6_Nodes_Node_Proxy_Binding_Summary_Iana
// IANA proxy binding summary
type Dhcpv6_Nodes_Node_Proxy_Binding_Summary_Iana struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of clients in init state. The type is interface{} with range:
    // 0..4294967295.
    InitializingClients interface{}

    // Number of clients waiting on DPM to validate subscriber. The type is
    // interface{} with range: 0..4294967295.
    DpmWaitingClients interface{}

    // Number of clients waiting on DAPS to assign/free prefix(ND). The type is
    // interface{} with range: 0..4294967295.
    DapsWaitingClients interface{}

    // Number of clients waiting for a message from the client/server. The type is
    // interface{} with range: 0..4294967295.
    MsgWaitingClients interface{}

    // Number of clients waiting on iedge to subscriber session. The type is
    // interface{} with range: 0..4294967295.
    IedgeWaitingClients interface{}

    // Number of clients in waiting on RIB response. The type is interface{} with
    // range: 0..4294967295.
    RibWaitingClients interface{}

    // Number of clients in bound state. The type is interface{} with range:
    // 0..4294967295.
    BoundClients interface{}
}

func (iana *Dhcpv6_Nodes_Node_Proxy_Binding_Summary_Iana) GetFilter() yfilter.YFilter { return iana.YFilter }

func (iana *Dhcpv6_Nodes_Node_Proxy_Binding_Summary_Iana) SetFilter(yf yfilter.YFilter) { iana.YFilter = yf }

func (iana *Dhcpv6_Nodes_Node_Proxy_Binding_Summary_Iana) GetGoName(yname string) string {
    if yname == "initializing-clients" { return "InitializingClients" }
    if yname == "dpm-waiting-clients" { return "DpmWaitingClients" }
    if yname == "daps-waiting-clients" { return "DapsWaitingClients" }
    if yname == "msg-waiting-clients" { return "MsgWaitingClients" }
    if yname == "iedge-waiting-clients" { return "IedgeWaitingClients" }
    if yname == "rib-waiting-clients" { return "RibWaitingClients" }
    if yname == "bound-clients" { return "BoundClients" }
    return ""
}

func (iana *Dhcpv6_Nodes_Node_Proxy_Binding_Summary_Iana) GetSegmentPath() string {
    return "iana"
}

func (iana *Dhcpv6_Nodes_Node_Proxy_Binding_Summary_Iana) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (iana *Dhcpv6_Nodes_Node_Proxy_Binding_Summary_Iana) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (iana *Dhcpv6_Nodes_Node_Proxy_Binding_Summary_Iana) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["initializing-clients"] = iana.InitializingClients
    leafs["dpm-waiting-clients"] = iana.DpmWaitingClients
    leafs["daps-waiting-clients"] = iana.DapsWaitingClients
    leafs["msg-waiting-clients"] = iana.MsgWaitingClients
    leafs["iedge-waiting-clients"] = iana.IedgeWaitingClients
    leafs["rib-waiting-clients"] = iana.RibWaitingClients
    leafs["bound-clients"] = iana.BoundClients
    return leafs
}

func (iana *Dhcpv6_Nodes_Node_Proxy_Binding_Summary_Iana) GetBundleName() string { return "cisco_ios_xr" }

func (iana *Dhcpv6_Nodes_Node_Proxy_Binding_Summary_Iana) GetYangName() string { return "iana" }

func (iana *Dhcpv6_Nodes_Node_Proxy_Binding_Summary_Iana) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (iana *Dhcpv6_Nodes_Node_Proxy_Binding_Summary_Iana) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (iana *Dhcpv6_Nodes_Node_Proxy_Binding_Summary_Iana) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (iana *Dhcpv6_Nodes_Node_Proxy_Binding_Summary_Iana) SetParent(parent types.Entity) { iana.parent = parent }

func (iana *Dhcpv6_Nodes_Node_Proxy_Binding_Summary_Iana) GetParent() types.Entity { return iana.parent }

func (iana *Dhcpv6_Nodes_Node_Proxy_Binding_Summary_Iana) GetParentYangName() string { return "summary" }

// Dhcpv6_Nodes_Node_Proxy_Binding_Summary_Iapd
// IAPD proxy binding summary
type Dhcpv6_Nodes_Node_Proxy_Binding_Summary_Iapd struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of clients in init state. The type is interface{} with range:
    // 0..4294967295.
    InitializingClients interface{}

    // Number of clients waiting on DPM to validate subscriber. The type is
    // interface{} with range: 0..4294967295.
    DpmWaitingClients interface{}

    // Number of clients waiting on DAPS to assign/free prefix(ND). The type is
    // interface{} with range: 0..4294967295.
    DapsWaitingClients interface{}

    // Number of clients waiting for a message from the client/server. The type is
    // interface{} with range: 0..4294967295.
    MsgWaitingClients interface{}

    // Number of clients waiting on iedge to subscriber session. The type is
    // interface{} with range: 0..4294967295.
    IedgeWaitingClients interface{}

    // Number of clients in waiting on RIB response. The type is interface{} with
    // range: 0..4294967295.
    RibWaitingClients interface{}

    // Number of clients in bound state. The type is interface{} with range:
    // 0..4294967295.
    BoundClients interface{}
}

func (iapd *Dhcpv6_Nodes_Node_Proxy_Binding_Summary_Iapd) GetFilter() yfilter.YFilter { return iapd.YFilter }

func (iapd *Dhcpv6_Nodes_Node_Proxy_Binding_Summary_Iapd) SetFilter(yf yfilter.YFilter) { iapd.YFilter = yf }

func (iapd *Dhcpv6_Nodes_Node_Proxy_Binding_Summary_Iapd) GetGoName(yname string) string {
    if yname == "initializing-clients" { return "InitializingClients" }
    if yname == "dpm-waiting-clients" { return "DpmWaitingClients" }
    if yname == "daps-waiting-clients" { return "DapsWaitingClients" }
    if yname == "msg-waiting-clients" { return "MsgWaitingClients" }
    if yname == "iedge-waiting-clients" { return "IedgeWaitingClients" }
    if yname == "rib-waiting-clients" { return "RibWaitingClients" }
    if yname == "bound-clients" { return "BoundClients" }
    return ""
}

func (iapd *Dhcpv6_Nodes_Node_Proxy_Binding_Summary_Iapd) GetSegmentPath() string {
    return "iapd"
}

func (iapd *Dhcpv6_Nodes_Node_Proxy_Binding_Summary_Iapd) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (iapd *Dhcpv6_Nodes_Node_Proxy_Binding_Summary_Iapd) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (iapd *Dhcpv6_Nodes_Node_Proxy_Binding_Summary_Iapd) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["initializing-clients"] = iapd.InitializingClients
    leafs["dpm-waiting-clients"] = iapd.DpmWaitingClients
    leafs["daps-waiting-clients"] = iapd.DapsWaitingClients
    leafs["msg-waiting-clients"] = iapd.MsgWaitingClients
    leafs["iedge-waiting-clients"] = iapd.IedgeWaitingClients
    leafs["rib-waiting-clients"] = iapd.RibWaitingClients
    leafs["bound-clients"] = iapd.BoundClients
    return leafs
}

func (iapd *Dhcpv6_Nodes_Node_Proxy_Binding_Summary_Iapd) GetBundleName() string { return "cisco_ios_xr" }

func (iapd *Dhcpv6_Nodes_Node_Proxy_Binding_Summary_Iapd) GetYangName() string { return "iapd" }

func (iapd *Dhcpv6_Nodes_Node_Proxy_Binding_Summary_Iapd) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (iapd *Dhcpv6_Nodes_Node_Proxy_Binding_Summary_Iapd) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (iapd *Dhcpv6_Nodes_Node_Proxy_Binding_Summary_Iapd) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (iapd *Dhcpv6_Nodes_Node_Proxy_Binding_Summary_Iapd) SetParent(parent types.Entity) { iapd.parent = parent }

func (iapd *Dhcpv6_Nodes_Node_Proxy_Binding_Summary_Iapd) GetParent() types.Entity { return iapd.parent }

func (iapd *Dhcpv6_Nodes_Node_Proxy_Binding_Summary_Iapd) GetParentYangName() string { return "summary" }

// Dhcpv6_Nodes_Node_Base
// IPv6 DHCP Base
type Dhcpv6_Nodes_Node_Base struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DHCP database.
    Database Dhcpv6_Nodes_Node_Base_Database

    // IPv6 DHCP Base Binding.
    AddrBindings Dhcpv6_Nodes_Node_Base_AddrBindings
}

func (base *Dhcpv6_Nodes_Node_Base) GetFilter() yfilter.YFilter { return base.YFilter }

func (base *Dhcpv6_Nodes_Node_Base) SetFilter(yf yfilter.YFilter) { base.YFilter = yf }

func (base *Dhcpv6_Nodes_Node_Base) GetGoName(yname string) string {
    if yname == "database" { return "Database" }
    if yname == "addr-bindings" { return "AddrBindings" }
    return ""
}

func (base *Dhcpv6_Nodes_Node_Base) GetSegmentPath() string {
    return "base"
}

func (base *Dhcpv6_Nodes_Node_Base) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "database" {
        return &base.Database
    }
    if childYangName == "addr-bindings" {
        return &base.AddrBindings
    }
    return nil
}

func (base *Dhcpv6_Nodes_Node_Base) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["database"] = &base.Database
    children["addr-bindings"] = &base.AddrBindings
    return children
}

func (base *Dhcpv6_Nodes_Node_Base) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (base *Dhcpv6_Nodes_Node_Base) GetBundleName() string { return "cisco_ios_xr" }

func (base *Dhcpv6_Nodes_Node_Base) GetYangName() string { return "base" }

func (base *Dhcpv6_Nodes_Node_Base) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (base *Dhcpv6_Nodes_Node_Base) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (base *Dhcpv6_Nodes_Node_Base) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (base *Dhcpv6_Nodes_Node_Base) SetParent(parent types.Entity) { base.parent = parent }

func (base *Dhcpv6_Nodes_Node_Base) GetParent() types.Entity { return base.parent }

func (base *Dhcpv6_Nodes_Node_Base) GetParentYangName() string { return "node" }

// Dhcpv6_Nodes_Node_Base_Database
// DHCP database
type Dhcpv6_Nodes_Node_Base_Database struct {
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

func (database *Dhcpv6_Nodes_Node_Base_Database) GetFilter() yfilter.YFilter { return database.YFilter }

func (database *Dhcpv6_Nodes_Node_Base_Database) SetFilter(yf yfilter.YFilter) { database.YFilter = yf }

func (database *Dhcpv6_Nodes_Node_Base_Database) GetGoName(yname string) string {
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

func (database *Dhcpv6_Nodes_Node_Base_Database) GetSegmentPath() string {
    return "database"
}

func (database *Dhcpv6_Nodes_Node_Base_Database) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (database *Dhcpv6_Nodes_Node_Base_Database) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (database *Dhcpv6_Nodes_Node_Base_Database) GetLeafs() map[string]interface{} {
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

func (database *Dhcpv6_Nodes_Node_Base_Database) GetBundleName() string { return "cisco_ios_xr" }

func (database *Dhcpv6_Nodes_Node_Base_Database) GetYangName() string { return "database" }

func (database *Dhcpv6_Nodes_Node_Base_Database) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (database *Dhcpv6_Nodes_Node_Base_Database) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (database *Dhcpv6_Nodes_Node_Base_Database) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (database *Dhcpv6_Nodes_Node_Base_Database) SetParent(parent types.Entity) { database.parent = parent }

func (database *Dhcpv6_Nodes_Node_Base_Database) GetParent() types.Entity { return database.parent }

func (database *Dhcpv6_Nodes_Node_Base_Database) GetParentYangName() string { return "base" }

// Dhcpv6_Nodes_Node_Base_AddrBindings
// IPv6 DHCP Base Binding
type Dhcpv6_Nodes_Node_Base_AddrBindings struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DHCPv6 base stats debug. The type is slice of
    // Dhcpv6_Nodes_Node_Base_AddrBindings_AddrBinding.
    AddrBinding []Dhcpv6_Nodes_Node_Base_AddrBindings_AddrBinding
}

func (addrBindings *Dhcpv6_Nodes_Node_Base_AddrBindings) GetFilter() yfilter.YFilter { return addrBindings.YFilter }

func (addrBindings *Dhcpv6_Nodes_Node_Base_AddrBindings) SetFilter(yf yfilter.YFilter) { addrBindings.YFilter = yf }

func (addrBindings *Dhcpv6_Nodes_Node_Base_AddrBindings) GetGoName(yname string) string {
    if yname == "addr-binding" { return "AddrBinding" }
    return ""
}

func (addrBindings *Dhcpv6_Nodes_Node_Base_AddrBindings) GetSegmentPath() string {
    return "addr-bindings"
}

func (addrBindings *Dhcpv6_Nodes_Node_Base_AddrBindings) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "addr-binding" {
        for _, c := range addrBindings.AddrBinding {
            if addrBindings.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dhcpv6_Nodes_Node_Base_AddrBindings_AddrBinding{}
        addrBindings.AddrBinding = append(addrBindings.AddrBinding, child)
        return &addrBindings.AddrBinding[len(addrBindings.AddrBinding)-1]
    }
    return nil
}

func (addrBindings *Dhcpv6_Nodes_Node_Base_AddrBindings) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range addrBindings.AddrBinding {
        children[addrBindings.AddrBinding[i].GetSegmentPath()] = &addrBindings.AddrBinding[i]
    }
    return children
}

func (addrBindings *Dhcpv6_Nodes_Node_Base_AddrBindings) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (addrBindings *Dhcpv6_Nodes_Node_Base_AddrBindings) GetBundleName() string { return "cisco_ios_xr" }

func (addrBindings *Dhcpv6_Nodes_Node_Base_AddrBindings) GetYangName() string { return "addr-bindings" }

func (addrBindings *Dhcpv6_Nodes_Node_Base_AddrBindings) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (addrBindings *Dhcpv6_Nodes_Node_Base_AddrBindings) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (addrBindings *Dhcpv6_Nodes_Node_Base_AddrBindings) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (addrBindings *Dhcpv6_Nodes_Node_Base_AddrBindings) SetParent(parent types.Entity) { addrBindings.parent = parent }

func (addrBindings *Dhcpv6_Nodes_Node_Base_AddrBindings) GetParent() types.Entity { return addrBindings.parent }

func (addrBindings *Dhcpv6_Nodes_Node_Base_AddrBindings) GetParentYangName() string { return "base" }

// Dhcpv6_Nodes_Node_Base_AddrBindings_AddrBinding
// DHCPv6 base stats debug
type Dhcpv6_Nodes_Node_Base_AddrBindings_AddrBinding struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Address String. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    AddrString interface{}

    // DHCPV6 client MAC address. The type is string.
    MacAddress interface{}

    // DHCPV6 client/subscriber VRF name. The type is string with length: 0..33.
    VrfName interface{}

    // DHCPV6 server VRF name. The type is string with length: 0..33.
    ServerVrfName interface{}

    // DHCPV6 IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}

    // DHCPV6 server IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    ServerIpv6Address interface{}

    // DHCPV6 reply server IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    ReplyServerIpv6Address interface{}

    // Lease time in seconds. The type is interface{} with range: 0..4294967295.
    // Units are second.
    LeaseTime interface{}

    // Remaining lease time in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    RemainingLeaseTime interface{}

    // DHCPV6 client state. The type is BagDhcpv6DFsmState.
    State interface{}

    // DHCPV6 access interface to client. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // DHCPV6 access interface VRF name. The type is string with length: 0..33.
    AccessVrfName interface{}

    // DHCPV6 VLAN tag count. The type is interface{} with range: 0..255.
    BaseBindingTags interface{}

    // DHCPV6 VLAN Outer VLAN. The type is interface{} with range: 0..4294967295.
    BaseBindingOuterTag interface{}

    // DHCPV6 VLAN Inner VLAN. The type is interface{} with range: 0..4294967295.
    BaseBindingInnerTag interface{}

    // DHCPV6 profile name. The type is string with length: 0..64.
    ProfileName interface{}

    // Is true if DHCPV6 next renew from client will be NAK'd. The type is bool.
    IsNakNextRenew interface{}

    // DHCPV6 subscriber label. The type is interface{} with range: 0..4294967295.
    SubscriberLabel interface{}

    // DHCPV6 old subscriber label. The type is interface{} with range:
    // 0..4294967295.
    OldSubscriberLabel interface{}

    // DHCPV6 received client DUID. The type is string with length: 0..771.
    RxClientDuid interface{}

    // DHCPV6 transmitted client DUID. The type is string with length: 0..771.
    TxClientUid interface{}

    // DHCPV6 received Remote ID. The type is string with length: 0..771.
    RxRemoteId interface{}

    // DHCPV6 transmitted Remote ID. The type is string with length: 0..771.
    TxRemoteId interface{}

    // DHCPV6 received Interface ID. The type is string with length: 0..771.
    RxInterfaceId interface{}

    // DHCPV6 transmitted Interface ID. The type is string with length: 0..771.
    TxInterfaceId interface{}
}

func (addrBinding *Dhcpv6_Nodes_Node_Base_AddrBindings_AddrBinding) GetFilter() yfilter.YFilter { return addrBinding.YFilter }

func (addrBinding *Dhcpv6_Nodes_Node_Base_AddrBindings_AddrBinding) SetFilter(yf yfilter.YFilter) { addrBinding.YFilter = yf }

func (addrBinding *Dhcpv6_Nodes_Node_Base_AddrBindings_AddrBinding) GetGoName(yname string) string {
    if yname == "addr-string" { return "AddrString" }
    if yname == "mac-address" { return "MacAddress" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "server-vrf-name" { return "ServerVrfName" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    if yname == "server-ipv6-address" { return "ServerIpv6Address" }
    if yname == "reply-server-ipv6-address" { return "ReplyServerIpv6Address" }
    if yname == "lease-time" { return "LeaseTime" }
    if yname == "remaining-lease-time" { return "RemainingLeaseTime" }
    if yname == "state" { return "State" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "access-vrf-name" { return "AccessVrfName" }
    if yname == "base-binding-tags" { return "BaseBindingTags" }
    if yname == "base-binding-outer-tag" { return "BaseBindingOuterTag" }
    if yname == "base-binding-inner-tag" { return "BaseBindingInnerTag" }
    if yname == "profile-name" { return "ProfileName" }
    if yname == "is-nak-next-renew" { return "IsNakNextRenew" }
    if yname == "subscriber-label" { return "SubscriberLabel" }
    if yname == "old-subscriber-label" { return "OldSubscriberLabel" }
    if yname == "rx-client-duid" { return "RxClientDuid" }
    if yname == "tx-client-uid" { return "TxClientUid" }
    if yname == "rx-remote-id" { return "RxRemoteId" }
    if yname == "tx-remote-id" { return "TxRemoteId" }
    if yname == "rx-interface-id" { return "RxInterfaceId" }
    if yname == "tx-interface-id" { return "TxInterfaceId" }
    return ""
}

func (addrBinding *Dhcpv6_Nodes_Node_Base_AddrBindings_AddrBinding) GetSegmentPath() string {
    return "addr-binding" + "[addr-string='" + fmt.Sprintf("%v", addrBinding.AddrString) + "']"
}

func (addrBinding *Dhcpv6_Nodes_Node_Base_AddrBindings_AddrBinding) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (addrBinding *Dhcpv6_Nodes_Node_Base_AddrBindings_AddrBinding) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (addrBinding *Dhcpv6_Nodes_Node_Base_AddrBindings_AddrBinding) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["addr-string"] = addrBinding.AddrString
    leafs["mac-address"] = addrBinding.MacAddress
    leafs["vrf-name"] = addrBinding.VrfName
    leafs["server-vrf-name"] = addrBinding.ServerVrfName
    leafs["ipv6-address"] = addrBinding.Ipv6Address
    leafs["server-ipv6-address"] = addrBinding.ServerIpv6Address
    leafs["reply-server-ipv6-address"] = addrBinding.ReplyServerIpv6Address
    leafs["lease-time"] = addrBinding.LeaseTime
    leafs["remaining-lease-time"] = addrBinding.RemainingLeaseTime
    leafs["state"] = addrBinding.State
    leafs["interface-name"] = addrBinding.InterfaceName
    leafs["access-vrf-name"] = addrBinding.AccessVrfName
    leafs["base-binding-tags"] = addrBinding.BaseBindingTags
    leafs["base-binding-outer-tag"] = addrBinding.BaseBindingOuterTag
    leafs["base-binding-inner-tag"] = addrBinding.BaseBindingInnerTag
    leafs["profile-name"] = addrBinding.ProfileName
    leafs["is-nak-next-renew"] = addrBinding.IsNakNextRenew
    leafs["subscriber-label"] = addrBinding.SubscriberLabel
    leafs["old-subscriber-label"] = addrBinding.OldSubscriberLabel
    leafs["rx-client-duid"] = addrBinding.RxClientDuid
    leafs["tx-client-uid"] = addrBinding.TxClientUid
    leafs["rx-remote-id"] = addrBinding.RxRemoteId
    leafs["tx-remote-id"] = addrBinding.TxRemoteId
    leafs["rx-interface-id"] = addrBinding.RxInterfaceId
    leafs["tx-interface-id"] = addrBinding.TxInterfaceId
    return leafs
}

func (addrBinding *Dhcpv6_Nodes_Node_Base_AddrBindings_AddrBinding) GetBundleName() string { return "cisco_ios_xr" }

func (addrBinding *Dhcpv6_Nodes_Node_Base_AddrBindings_AddrBinding) GetYangName() string { return "addr-binding" }

func (addrBinding *Dhcpv6_Nodes_Node_Base_AddrBindings_AddrBinding) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (addrBinding *Dhcpv6_Nodes_Node_Base_AddrBindings_AddrBinding) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (addrBinding *Dhcpv6_Nodes_Node_Base_AddrBindings_AddrBinding) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (addrBinding *Dhcpv6_Nodes_Node_Base_AddrBindings_AddrBinding) SetParent(parent types.Entity) { addrBinding.parent = parent }

func (addrBinding *Dhcpv6_Nodes_Node_Base_AddrBindings_AddrBinding) GetParent() types.Entity { return addrBinding.parent }

func (addrBinding *Dhcpv6_Nodes_Node_Base_AddrBindings_AddrBinding) GetParentYangName() string { return "addr-bindings" }

// Dhcpv6_Nodes_Node_Server
// IPv6 DHCP server operational data
type Dhcpv6_Nodes_Node_Server struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DHCPV6 server bindings.
    Binding Dhcpv6_Nodes_Node_Server_Binding

    // DHCPV6 server list of VRF names.
    Vrfs Dhcpv6_Nodes_Node_Server_Vrfs

    // IPv6 DHCP server profile.
    Profiles Dhcpv6_Nodes_Node_Server_Profiles

    // DHCPV6 server interface.
    Interfaces Dhcpv6_Nodes_Node_Server_Interfaces

    // DHCPv6 server statistics.
    Statistics Dhcpv6_Nodes_Node_Server_Statistics

    // DHCPv6 server binding with options.
    BindingOptions Dhcpv6_Nodes_Node_Server_BindingOptions
}

func (server *Dhcpv6_Nodes_Node_Server) GetFilter() yfilter.YFilter { return server.YFilter }

func (server *Dhcpv6_Nodes_Node_Server) SetFilter(yf yfilter.YFilter) { server.YFilter = yf }

func (server *Dhcpv6_Nodes_Node_Server) GetGoName(yname string) string {
    if yname == "binding" { return "Binding" }
    if yname == "vrfs" { return "Vrfs" }
    if yname == "profiles" { return "Profiles" }
    if yname == "interfaces" { return "Interfaces" }
    if yname == "statistics" { return "Statistics" }
    if yname == "binding-options" { return "BindingOptions" }
    return ""
}

func (server *Dhcpv6_Nodes_Node_Server) GetSegmentPath() string {
    return "server"
}

func (server *Dhcpv6_Nodes_Node_Server) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "binding" {
        return &server.Binding
    }
    if childYangName == "vrfs" {
        return &server.Vrfs
    }
    if childYangName == "profiles" {
        return &server.Profiles
    }
    if childYangName == "interfaces" {
        return &server.Interfaces
    }
    if childYangName == "statistics" {
        return &server.Statistics
    }
    if childYangName == "binding-options" {
        return &server.BindingOptions
    }
    return nil
}

func (server *Dhcpv6_Nodes_Node_Server) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["binding"] = &server.Binding
    children["vrfs"] = &server.Vrfs
    children["profiles"] = &server.Profiles
    children["interfaces"] = &server.Interfaces
    children["statistics"] = &server.Statistics
    children["binding-options"] = &server.BindingOptions
    return children
}

func (server *Dhcpv6_Nodes_Node_Server) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (server *Dhcpv6_Nodes_Node_Server) GetBundleName() string { return "cisco_ios_xr" }

func (server *Dhcpv6_Nodes_Node_Server) GetYangName() string { return "server" }

func (server *Dhcpv6_Nodes_Node_Server) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (server *Dhcpv6_Nodes_Node_Server) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (server *Dhcpv6_Nodes_Node_Server) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (server *Dhcpv6_Nodes_Node_Server) SetParent(parent types.Entity) { server.parent = parent }

func (server *Dhcpv6_Nodes_Node_Server) GetParent() types.Entity { return server.parent }

func (server *Dhcpv6_Nodes_Node_Server) GetParentYangName() string { return "node" }

// Dhcpv6_Nodes_Node_Server_Binding
// DHCPV6 server bindings
type Dhcpv6_Nodes_Node_Server_Binding struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DHCPV6 server binding summary.
    Summary Dhcpv6_Nodes_Node_Server_Binding_Summary

    // DHCPV6 server client bindings.
    Clients Dhcpv6_Nodes_Node_Server_Binding_Clients
}

func (binding *Dhcpv6_Nodes_Node_Server_Binding) GetFilter() yfilter.YFilter { return binding.YFilter }

func (binding *Dhcpv6_Nodes_Node_Server_Binding) SetFilter(yf yfilter.YFilter) { binding.YFilter = yf }

func (binding *Dhcpv6_Nodes_Node_Server_Binding) GetGoName(yname string) string {
    if yname == "summary" { return "Summary" }
    if yname == "clients" { return "Clients" }
    return ""
}

func (binding *Dhcpv6_Nodes_Node_Server_Binding) GetSegmentPath() string {
    return "binding"
}

func (binding *Dhcpv6_Nodes_Node_Server_Binding) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "summary" {
        return &binding.Summary
    }
    if childYangName == "clients" {
        return &binding.Clients
    }
    return nil
}

func (binding *Dhcpv6_Nodes_Node_Server_Binding) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["summary"] = &binding.Summary
    children["clients"] = &binding.Clients
    return children
}

func (binding *Dhcpv6_Nodes_Node_Server_Binding) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (binding *Dhcpv6_Nodes_Node_Server_Binding) GetBundleName() string { return "cisco_ios_xr" }

func (binding *Dhcpv6_Nodes_Node_Server_Binding) GetYangName() string { return "binding" }

func (binding *Dhcpv6_Nodes_Node_Server_Binding) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (binding *Dhcpv6_Nodes_Node_Server_Binding) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (binding *Dhcpv6_Nodes_Node_Server_Binding) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (binding *Dhcpv6_Nodes_Node_Server_Binding) SetParent(parent types.Entity) { binding.parent = parent }

func (binding *Dhcpv6_Nodes_Node_Server_Binding) GetParent() types.Entity { return binding.parent }

func (binding *Dhcpv6_Nodes_Node_Server_Binding) GetParentYangName() string { return "server" }

// Dhcpv6_Nodes_Node_Server_Binding_Summary
// DHCPV6 server binding summary
type Dhcpv6_Nodes_Node_Server_Binding_Summary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Total number of clients. The type is interface{} with range: 0..4294967295.
    Clients interface{}

    // IANA server binding summary.
    Iana Dhcpv6_Nodes_Node_Server_Binding_Summary_Iana

    // IAPD server binding summary.
    Iapd Dhcpv6_Nodes_Node_Server_Binding_Summary_Iapd
}

func (summary *Dhcpv6_Nodes_Node_Server_Binding_Summary) GetFilter() yfilter.YFilter { return summary.YFilter }

func (summary *Dhcpv6_Nodes_Node_Server_Binding_Summary) SetFilter(yf yfilter.YFilter) { summary.YFilter = yf }

func (summary *Dhcpv6_Nodes_Node_Server_Binding_Summary) GetGoName(yname string) string {
    if yname == "clients" { return "Clients" }
    if yname == "iana" { return "Iana" }
    if yname == "iapd" { return "Iapd" }
    return ""
}

func (summary *Dhcpv6_Nodes_Node_Server_Binding_Summary) GetSegmentPath() string {
    return "summary"
}

func (summary *Dhcpv6_Nodes_Node_Server_Binding_Summary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "iana" {
        return &summary.Iana
    }
    if childYangName == "iapd" {
        return &summary.Iapd
    }
    return nil
}

func (summary *Dhcpv6_Nodes_Node_Server_Binding_Summary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["iana"] = &summary.Iana
    children["iapd"] = &summary.Iapd
    return children
}

func (summary *Dhcpv6_Nodes_Node_Server_Binding_Summary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["clients"] = summary.Clients
    return leafs
}

func (summary *Dhcpv6_Nodes_Node_Server_Binding_Summary) GetBundleName() string { return "cisco_ios_xr" }

func (summary *Dhcpv6_Nodes_Node_Server_Binding_Summary) GetYangName() string { return "summary" }

func (summary *Dhcpv6_Nodes_Node_Server_Binding_Summary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (summary *Dhcpv6_Nodes_Node_Server_Binding_Summary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (summary *Dhcpv6_Nodes_Node_Server_Binding_Summary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (summary *Dhcpv6_Nodes_Node_Server_Binding_Summary) SetParent(parent types.Entity) { summary.parent = parent }

func (summary *Dhcpv6_Nodes_Node_Server_Binding_Summary) GetParent() types.Entity { return summary.parent }

func (summary *Dhcpv6_Nodes_Node_Server_Binding_Summary) GetParentYangName() string { return "binding" }

// Dhcpv6_Nodes_Node_Server_Binding_Summary_Iana
// IANA server binding summary
type Dhcpv6_Nodes_Node_Server_Binding_Summary_Iana struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of clients in init state. The type is interface{} with range:
    // 0..4294967295.
    InitializingClients interface{}

    // Number of clients waiting on DPM to validate subscriber. The type is
    // interface{} with range: 0..4294967295.
    DpmWaitingClients interface{}

    // Number of clients waiting on DAPS to assign/free addr/prefix. The type is
    // interface{} with range: 0..4294967295.
    DapsWaitingClients interface{}

    // Number of clients waiting for a request from the client. The type is
    // interface{} with range: 0..4294967295.
    RequestWaitingClients interface{}

    // Number of clients waiting for iedge for the session. The type is
    // interface{} with range: 0..4294967295.
    IedgeWaitingClients interface{}

    // Number of clients in waiting for RIB response. The type is interface{} with
    // range: 0..4294967295.
    RibWaitingClients interface{}

    // Number of clients in bound state. The type is interface{} with range:
    // 0..4294967295.
    BoundClients interface{}
}

func (iana *Dhcpv6_Nodes_Node_Server_Binding_Summary_Iana) GetFilter() yfilter.YFilter { return iana.YFilter }

func (iana *Dhcpv6_Nodes_Node_Server_Binding_Summary_Iana) SetFilter(yf yfilter.YFilter) { iana.YFilter = yf }

func (iana *Dhcpv6_Nodes_Node_Server_Binding_Summary_Iana) GetGoName(yname string) string {
    if yname == "initializing-clients" { return "InitializingClients" }
    if yname == "dpm-waiting-clients" { return "DpmWaitingClients" }
    if yname == "daps-waiting-clients" { return "DapsWaitingClients" }
    if yname == "request-waiting-clients" { return "RequestWaitingClients" }
    if yname == "iedge-waiting-clients" { return "IedgeWaitingClients" }
    if yname == "rib-waiting-clients" { return "RibWaitingClients" }
    if yname == "bound-clients" { return "BoundClients" }
    return ""
}

func (iana *Dhcpv6_Nodes_Node_Server_Binding_Summary_Iana) GetSegmentPath() string {
    return "iana"
}

func (iana *Dhcpv6_Nodes_Node_Server_Binding_Summary_Iana) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (iana *Dhcpv6_Nodes_Node_Server_Binding_Summary_Iana) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (iana *Dhcpv6_Nodes_Node_Server_Binding_Summary_Iana) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["initializing-clients"] = iana.InitializingClients
    leafs["dpm-waiting-clients"] = iana.DpmWaitingClients
    leafs["daps-waiting-clients"] = iana.DapsWaitingClients
    leafs["request-waiting-clients"] = iana.RequestWaitingClients
    leafs["iedge-waiting-clients"] = iana.IedgeWaitingClients
    leafs["rib-waiting-clients"] = iana.RibWaitingClients
    leafs["bound-clients"] = iana.BoundClients
    return leafs
}

func (iana *Dhcpv6_Nodes_Node_Server_Binding_Summary_Iana) GetBundleName() string { return "cisco_ios_xr" }

func (iana *Dhcpv6_Nodes_Node_Server_Binding_Summary_Iana) GetYangName() string { return "iana" }

func (iana *Dhcpv6_Nodes_Node_Server_Binding_Summary_Iana) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (iana *Dhcpv6_Nodes_Node_Server_Binding_Summary_Iana) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (iana *Dhcpv6_Nodes_Node_Server_Binding_Summary_Iana) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (iana *Dhcpv6_Nodes_Node_Server_Binding_Summary_Iana) SetParent(parent types.Entity) { iana.parent = parent }

func (iana *Dhcpv6_Nodes_Node_Server_Binding_Summary_Iana) GetParent() types.Entity { return iana.parent }

func (iana *Dhcpv6_Nodes_Node_Server_Binding_Summary_Iana) GetParentYangName() string { return "summary" }

// Dhcpv6_Nodes_Node_Server_Binding_Summary_Iapd
// IAPD server binding summary
type Dhcpv6_Nodes_Node_Server_Binding_Summary_Iapd struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of clients in init state. The type is interface{} with range:
    // 0..4294967295.
    InitializingClients interface{}

    // Number of clients waiting on DPM to validate subscriber. The type is
    // interface{} with range: 0..4294967295.
    DpmWaitingClients interface{}

    // Number of clients waiting on DAPS to assign/free addr/prefix. The type is
    // interface{} with range: 0..4294967295.
    DapsWaitingClients interface{}

    // Number of clients waiting for a request from the client. The type is
    // interface{} with range: 0..4294967295.
    RequestWaitingClients interface{}

    // Number of clients waiting for iedge for the session. The type is
    // interface{} with range: 0..4294967295.
    IedgeWaitingClients interface{}

    // Number of clients in waiting for RIB response. The type is interface{} with
    // range: 0..4294967295.
    RibWaitingClients interface{}

    // Number of clients in bound state. The type is interface{} with range:
    // 0..4294967295.
    BoundClients interface{}
}

func (iapd *Dhcpv6_Nodes_Node_Server_Binding_Summary_Iapd) GetFilter() yfilter.YFilter { return iapd.YFilter }

func (iapd *Dhcpv6_Nodes_Node_Server_Binding_Summary_Iapd) SetFilter(yf yfilter.YFilter) { iapd.YFilter = yf }

func (iapd *Dhcpv6_Nodes_Node_Server_Binding_Summary_Iapd) GetGoName(yname string) string {
    if yname == "initializing-clients" { return "InitializingClients" }
    if yname == "dpm-waiting-clients" { return "DpmWaitingClients" }
    if yname == "daps-waiting-clients" { return "DapsWaitingClients" }
    if yname == "request-waiting-clients" { return "RequestWaitingClients" }
    if yname == "iedge-waiting-clients" { return "IedgeWaitingClients" }
    if yname == "rib-waiting-clients" { return "RibWaitingClients" }
    if yname == "bound-clients" { return "BoundClients" }
    return ""
}

func (iapd *Dhcpv6_Nodes_Node_Server_Binding_Summary_Iapd) GetSegmentPath() string {
    return "iapd"
}

func (iapd *Dhcpv6_Nodes_Node_Server_Binding_Summary_Iapd) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (iapd *Dhcpv6_Nodes_Node_Server_Binding_Summary_Iapd) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (iapd *Dhcpv6_Nodes_Node_Server_Binding_Summary_Iapd) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["initializing-clients"] = iapd.InitializingClients
    leafs["dpm-waiting-clients"] = iapd.DpmWaitingClients
    leafs["daps-waiting-clients"] = iapd.DapsWaitingClients
    leafs["request-waiting-clients"] = iapd.RequestWaitingClients
    leafs["iedge-waiting-clients"] = iapd.IedgeWaitingClients
    leafs["rib-waiting-clients"] = iapd.RibWaitingClients
    leafs["bound-clients"] = iapd.BoundClients
    return leafs
}

func (iapd *Dhcpv6_Nodes_Node_Server_Binding_Summary_Iapd) GetBundleName() string { return "cisco_ios_xr" }

func (iapd *Dhcpv6_Nodes_Node_Server_Binding_Summary_Iapd) GetYangName() string { return "iapd" }

func (iapd *Dhcpv6_Nodes_Node_Server_Binding_Summary_Iapd) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (iapd *Dhcpv6_Nodes_Node_Server_Binding_Summary_Iapd) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (iapd *Dhcpv6_Nodes_Node_Server_Binding_Summary_Iapd) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (iapd *Dhcpv6_Nodes_Node_Server_Binding_Summary_Iapd) SetParent(parent types.Entity) { iapd.parent = parent }

func (iapd *Dhcpv6_Nodes_Node_Server_Binding_Summary_Iapd) GetParent() types.Entity { return iapd.parent }

func (iapd *Dhcpv6_Nodes_Node_Server_Binding_Summary_Iapd) GetParentYangName() string { return "summary" }

// Dhcpv6_Nodes_Node_Server_Binding_Clients
// DHCPV6 server client bindings
type Dhcpv6_Nodes_Node_Server_Binding_Clients struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Single DHCPV6 server binding. The type is slice of
    // Dhcpv6_Nodes_Node_Server_Binding_Clients_Client.
    Client []Dhcpv6_Nodes_Node_Server_Binding_Clients_Client
}

func (clients *Dhcpv6_Nodes_Node_Server_Binding_Clients) GetFilter() yfilter.YFilter { return clients.YFilter }

func (clients *Dhcpv6_Nodes_Node_Server_Binding_Clients) SetFilter(yf yfilter.YFilter) { clients.YFilter = yf }

func (clients *Dhcpv6_Nodes_Node_Server_Binding_Clients) GetGoName(yname string) string {
    if yname == "client" { return "Client" }
    return ""
}

func (clients *Dhcpv6_Nodes_Node_Server_Binding_Clients) GetSegmentPath() string {
    return "clients"
}

func (clients *Dhcpv6_Nodes_Node_Server_Binding_Clients) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "client" {
        for _, c := range clients.Client {
            if clients.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dhcpv6_Nodes_Node_Server_Binding_Clients_Client{}
        clients.Client = append(clients.Client, child)
        return &clients.Client[len(clients.Client)-1]
    }
    return nil
}

func (clients *Dhcpv6_Nodes_Node_Server_Binding_Clients) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range clients.Client {
        children[clients.Client[i].GetSegmentPath()] = &clients.Client[i]
    }
    return children
}

func (clients *Dhcpv6_Nodes_Node_Server_Binding_Clients) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (clients *Dhcpv6_Nodes_Node_Server_Binding_Clients) GetBundleName() string { return "cisco_ios_xr" }

func (clients *Dhcpv6_Nodes_Node_Server_Binding_Clients) GetYangName() string { return "clients" }

func (clients *Dhcpv6_Nodes_Node_Server_Binding_Clients) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clients *Dhcpv6_Nodes_Node_Server_Binding_Clients) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clients *Dhcpv6_Nodes_Node_Server_Binding_Clients) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clients *Dhcpv6_Nodes_Node_Server_Binding_Clients) SetParent(parent types.Entity) { clients.parent = parent }

func (clients *Dhcpv6_Nodes_Node_Server_Binding_Clients) GetParent() types.Entity { return clients.parent }

func (clients *Dhcpv6_Nodes_Node_Server_Binding_Clients) GetParentYangName() string { return "binding" }

// Dhcpv6_Nodes_Node_Server_Binding_Clients_Client
// Single DHCPV6 server binding
type Dhcpv6_Nodes_Node_Server_Binding_Clients_Client struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Client ID. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    ClientId interface{}

    // Client DUID. The type is string.
    Duid interface{}

    // Client unique identifier. The type is interface{} with range:
    // 0..4294967295.
    ClientIdXr interface{}

    // DHCPV6 client flag. The type is interface{} with range: 0..4294967295.
    ClientFlag interface{}

    // DHCPV6 subscriber label. The type is interface{} with range: 0..4294967295.
    SubscriberLabel interface{}

    // DHCPVV6 client/subscriber VRF name. The type is string with length: 0..33.
    VrfName interface{}

    // Client MAC address. The type is string.
    MacAddress interface{}

    // Number of ia_id/pd. The type is interface{} with range: 0..4294967295.
    IaIdPDs interface{}

    // DHCPV6 IPv6 client link local address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    LinkLocalAddress interface{}

    // DHCPV6 access interface to client. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // DHCPV6 access VRF name to client. The type is string with length: 0..33.
    AccessVrfName interface{}

    // DHCPV6 VLAN tag count. The type is interface{} with range: 0..255.
    ServerBindingTags interface{}

    // DHCPV6 VLAN Outer VLAN. The type is interface{} with range: 0..4294967295.
    ServerBindingOuterTag interface{}

    // DHCPV6 VLAN Inner VLAN. The type is interface{} with range: 0..4294967295.
    ServerBindingInnerTag interface{}

    // DHCPV6 pool name. The type is string with length: 0..64.
    PoolName interface{}

    // DHCPV6 profile name. The type is string with length: 0..64.
    ProfileName interface{}

    // DHCPV6 framed ipv6 addess used by ND. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    FramedIpv6Prefix interface{}

    // DHCPV6 framed ipv6 prefix length used by ND. The type is interface{} with
    // range: 0..255.
    FramedPrefixLength interface{}

    // DHCPV6 class name. The type is string with length: 0..64.
    ClassName interface{}

    // DHCPV6 received Remote ID. The type is string with length: 0..771.
    RxRemoteId interface{}

    // DHCPV6 received Interface ID. The type is string with length: 0..771.
    RxInterfaceId interface{}

    // DHCPV6 server prefix pool name. The type is string with length: 0..64.
    PrefixPoolName interface{}

    // DHCPV6 server address pool name. The type is string with length: 0..64.
    AddressPoolName interface{}

    // DNS server count. The type is interface{} with range: 0..4294967295.
    DnsServerCount interface{}

    // Is true if DHCPv6 next renew from client will be NAK'd. The type is bool.
    IsNakNextRenew interface{}

    // DHCPV6 SRG state. The type is interface{} with range: 0..4294967295.
    SrgState interface{}

    // DHCPV6 SRG Intf Role. The type is interface{} with range: 0..4294967295.
    SrgIntfRole interface{}

    // SRG P2P Status. The type is bool.
    Srgp2P interface{}

    // DHCPV6 SRG VRF NAME. The type is string with length: 0..33.
    SrgVrfName interface{}

    // DHCPV6 SERG state. The type is interface{} with range: 0..4294967295.
    SesrgState interface{}

    // DHCPV6 SERG Intf Role. The type is interface{} with range: 0..4294967295.
    SergIntfRole interface{}

    // List of DHCPv6 IA_ID/PDs.
    IaIdPd Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd
}

func (client *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client) GetFilter() yfilter.YFilter { return client.YFilter }

func (client *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client) SetFilter(yf yfilter.YFilter) { client.YFilter = yf }

func (client *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client) GetGoName(yname string) string {
    if yname == "client-id" { return "ClientId" }
    if yname == "duid" { return "Duid" }
    if yname == "client-id-xr" { return "ClientIdXr" }
    if yname == "client-flag" { return "ClientFlag" }
    if yname == "subscriber-label" { return "SubscriberLabel" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "mac-address" { return "MacAddress" }
    if yname == "ia-id-p-ds" { return "IaIdPDs" }
    if yname == "link-local-address" { return "LinkLocalAddress" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "access-vrf-name" { return "AccessVrfName" }
    if yname == "server-binding-tags" { return "ServerBindingTags" }
    if yname == "server-binding-outer-tag" { return "ServerBindingOuterTag" }
    if yname == "server-binding-inner-tag" { return "ServerBindingInnerTag" }
    if yname == "pool-name" { return "PoolName" }
    if yname == "profile-name" { return "ProfileName" }
    if yname == "framed-ipv6-prefix" { return "FramedIpv6Prefix" }
    if yname == "framed-prefix-length" { return "FramedPrefixLength" }
    if yname == "class-name" { return "ClassName" }
    if yname == "rx-remote-id" { return "RxRemoteId" }
    if yname == "rx-interface-id" { return "RxInterfaceId" }
    if yname == "prefix-pool-name" { return "PrefixPoolName" }
    if yname == "address-pool-name" { return "AddressPoolName" }
    if yname == "dns-server-count" { return "DnsServerCount" }
    if yname == "is-nak-next-renew" { return "IsNakNextRenew" }
    if yname == "srg-state" { return "SrgState" }
    if yname == "srg-intf-role" { return "SrgIntfRole" }
    if yname == "srgp2p" { return "Srgp2P" }
    if yname == "srg-vrf-name" { return "SrgVrfName" }
    if yname == "sesrg-state" { return "SesrgState" }
    if yname == "serg-intf-role" { return "SergIntfRole" }
    if yname == "ia-id-pd" { return "IaIdPd" }
    return ""
}

func (client *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client) GetSegmentPath() string {
    return "client" + "[client-id='" + fmt.Sprintf("%v", client.ClientId) + "']"
}

func (client *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ia-id-pd" {
        return &client.IaIdPd
    }
    return nil
}

func (client *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ia-id-pd"] = &client.IaIdPd
    return children
}

func (client *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["client-id"] = client.ClientId
    leafs["duid"] = client.Duid
    leafs["client-id-xr"] = client.ClientIdXr
    leafs["client-flag"] = client.ClientFlag
    leafs["subscriber-label"] = client.SubscriberLabel
    leafs["vrf-name"] = client.VrfName
    leafs["mac-address"] = client.MacAddress
    leafs["ia-id-p-ds"] = client.IaIdPDs
    leafs["link-local-address"] = client.LinkLocalAddress
    leafs["interface-name"] = client.InterfaceName
    leafs["access-vrf-name"] = client.AccessVrfName
    leafs["server-binding-tags"] = client.ServerBindingTags
    leafs["server-binding-outer-tag"] = client.ServerBindingOuterTag
    leafs["server-binding-inner-tag"] = client.ServerBindingInnerTag
    leafs["pool-name"] = client.PoolName
    leafs["profile-name"] = client.ProfileName
    leafs["framed-ipv6-prefix"] = client.FramedIpv6Prefix
    leafs["framed-prefix-length"] = client.FramedPrefixLength
    leafs["class-name"] = client.ClassName
    leafs["rx-remote-id"] = client.RxRemoteId
    leafs["rx-interface-id"] = client.RxInterfaceId
    leafs["prefix-pool-name"] = client.PrefixPoolName
    leafs["address-pool-name"] = client.AddressPoolName
    leafs["dns-server-count"] = client.DnsServerCount
    leafs["is-nak-next-renew"] = client.IsNakNextRenew
    leafs["srg-state"] = client.SrgState
    leafs["srg-intf-role"] = client.SrgIntfRole
    leafs["srgp2p"] = client.Srgp2P
    leafs["srg-vrf-name"] = client.SrgVrfName
    leafs["sesrg-state"] = client.SesrgState
    leafs["serg-intf-role"] = client.SergIntfRole
    return leafs
}

func (client *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client) GetBundleName() string { return "cisco_ios_xr" }

func (client *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client) GetYangName() string { return "client" }

func (client *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (client *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (client *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (client *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client) SetParent(parent types.Entity) { client.parent = parent }

func (client *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client) GetParent() types.Entity { return client.parent }

func (client *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client) GetParentYangName() string { return "clients" }

// Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd
// List of DHCPv6 IA_ID/PDs
type Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // bag dhcpv6d ia id pd info. The type is slice of
    // Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo.
    BagDhcpv6DIaIdPdInfo []Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo
}

func (iaIdPd *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd) GetFilter() yfilter.YFilter { return iaIdPd.YFilter }

func (iaIdPd *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd) SetFilter(yf yfilter.YFilter) { iaIdPd.YFilter = yf }

func (iaIdPd *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd) GetGoName(yname string) string {
    if yname == "bag-dhcpv6d-ia-id-pd-info" { return "BagDhcpv6DIaIdPdInfo" }
    return ""
}

func (iaIdPd *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd) GetSegmentPath() string {
    return "ia-id-pd"
}

func (iaIdPd *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bag-dhcpv6d-ia-id-pd-info" {
        for _, c := range iaIdPd.BagDhcpv6DIaIdPdInfo {
            if iaIdPd.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo{}
        iaIdPd.BagDhcpv6DIaIdPdInfo = append(iaIdPd.BagDhcpv6DIaIdPdInfo, child)
        return &iaIdPd.BagDhcpv6DIaIdPdInfo[len(iaIdPd.BagDhcpv6DIaIdPdInfo)-1]
    }
    return nil
}

func (iaIdPd *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range iaIdPd.BagDhcpv6DIaIdPdInfo {
        children[iaIdPd.BagDhcpv6DIaIdPdInfo[i].GetSegmentPath()] = &iaIdPd.BagDhcpv6DIaIdPdInfo[i]
    }
    return children
}

func (iaIdPd *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (iaIdPd *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd) GetBundleName() string { return "cisco_ios_xr" }

func (iaIdPd *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd) GetYangName() string { return "ia-id-pd" }

func (iaIdPd *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (iaIdPd *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (iaIdPd *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (iaIdPd *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd) SetParent(parent types.Entity) { iaIdPd.parent = parent }

func (iaIdPd *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd) GetParent() types.Entity { return iaIdPd.parent }

func (iaIdPd *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd) GetParentYangName() string { return "client" }

// Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo
// bag dhcpv6d ia id pd info
type Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IA type. The type is BagDhcpv6DIaId.
    IaType interface{}

    // IA_ID of this IA. The type is interface{} with range: 0..4294967295.
    IaId interface{}

    // FSM Flag for this IA. The type is interface{} with range: 0..4294967295.
    Flags interface{}

    // Total address in this IA. The type is interface{} with range: 0..65535.
    TotalAddress interface{}

    // State. The type is BagDhcpv6DFsmState.
    State interface{}

    // List of addresses in this IA.
    Addresses Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses
}

func (bagDhcpv6DIaIdPdInfo *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo) GetFilter() yfilter.YFilter { return bagDhcpv6DIaIdPdInfo.YFilter }

func (bagDhcpv6DIaIdPdInfo *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo) SetFilter(yf yfilter.YFilter) { bagDhcpv6DIaIdPdInfo.YFilter = yf }

func (bagDhcpv6DIaIdPdInfo *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo) GetGoName(yname string) string {
    if yname == "ia-type" { return "IaType" }
    if yname == "ia-id" { return "IaId" }
    if yname == "flags" { return "Flags" }
    if yname == "total-address" { return "TotalAddress" }
    if yname == "state" { return "State" }
    if yname == "addresses" { return "Addresses" }
    return ""
}

func (bagDhcpv6DIaIdPdInfo *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo) GetSegmentPath() string {
    return "bag-dhcpv6d-ia-id-pd-info"
}

func (bagDhcpv6DIaIdPdInfo *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "addresses" {
        return &bagDhcpv6DIaIdPdInfo.Addresses
    }
    return nil
}

func (bagDhcpv6DIaIdPdInfo *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["addresses"] = &bagDhcpv6DIaIdPdInfo.Addresses
    return children
}

func (bagDhcpv6DIaIdPdInfo *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ia-type"] = bagDhcpv6DIaIdPdInfo.IaType
    leafs["ia-id"] = bagDhcpv6DIaIdPdInfo.IaId
    leafs["flags"] = bagDhcpv6DIaIdPdInfo.Flags
    leafs["total-address"] = bagDhcpv6DIaIdPdInfo.TotalAddress
    leafs["state"] = bagDhcpv6DIaIdPdInfo.State
    return leafs
}

func (bagDhcpv6DIaIdPdInfo *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo) GetBundleName() string { return "cisco_ios_xr" }

func (bagDhcpv6DIaIdPdInfo *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo) GetYangName() string { return "bag-dhcpv6d-ia-id-pd-info" }

func (bagDhcpv6DIaIdPdInfo *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bagDhcpv6DIaIdPdInfo *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bagDhcpv6DIaIdPdInfo *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bagDhcpv6DIaIdPdInfo *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo) SetParent(parent types.Entity) { bagDhcpv6DIaIdPdInfo.parent = parent }

func (bagDhcpv6DIaIdPdInfo *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo) GetParent() types.Entity { return bagDhcpv6DIaIdPdInfo.parent }

func (bagDhcpv6DIaIdPdInfo *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo) GetParentYangName() string { return "ia-id-pd" }

// Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses
// List of addresses in this IA
type Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // bag dhcpv6d addr attrb. The type is slice of
    // Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses_BagDhcpv6DAddrAttrb.
    BagDhcpv6DAddrAttrb []Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses_BagDhcpv6DAddrAttrb
}

func (addresses *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses) GetFilter() yfilter.YFilter { return addresses.YFilter }

func (addresses *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses) SetFilter(yf yfilter.YFilter) { addresses.YFilter = yf }

func (addresses *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses) GetGoName(yname string) string {
    if yname == "bag-dhcpv6d-addr-attrb" { return "BagDhcpv6DAddrAttrb" }
    return ""
}

func (addresses *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses) GetSegmentPath() string {
    return "addresses"
}

func (addresses *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bag-dhcpv6d-addr-attrb" {
        for _, c := range addresses.BagDhcpv6DAddrAttrb {
            if addresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses_BagDhcpv6DAddrAttrb{}
        addresses.BagDhcpv6DAddrAttrb = append(addresses.BagDhcpv6DAddrAttrb, child)
        return &addresses.BagDhcpv6DAddrAttrb[len(addresses.BagDhcpv6DAddrAttrb)-1]
    }
    return nil
}

func (addresses *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range addresses.BagDhcpv6DAddrAttrb {
        children[addresses.BagDhcpv6DAddrAttrb[i].GetSegmentPath()] = &addresses.BagDhcpv6DAddrAttrb[i]
    }
    return children
}

func (addresses *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (addresses *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses) GetBundleName() string { return "cisco_ios_xr" }

func (addresses *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses) GetYangName() string { return "addresses" }

func (addresses *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (addresses *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (addresses *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (addresses *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses) SetParent(parent types.Entity) { addresses.parent = parent }

func (addresses *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses) GetParent() types.Entity { return addresses.parent }

func (addresses *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses) GetParentYangName() string { return "bag-dhcpv6d-ia-id-pd-info" }

// Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses_BagDhcpv6DAddrAttrb
// bag dhcpv6d addr attrb
type Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses_BagDhcpv6DAddrAttrb struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 prefix. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Prefix interface{}

    // Prefix length. The type is interface{} with range: 0..255.
    PrefixLength interface{}

    // Lease time in seconds. The type is interface{} with range: 0..4294967295.
    // Units are second.
    LeaseTime interface{}

    // Remaining lease time in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    RemainingLeaseTime interface{}
}

func (bagDhcpv6DAddrAttrb *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses_BagDhcpv6DAddrAttrb) GetFilter() yfilter.YFilter { return bagDhcpv6DAddrAttrb.YFilter }

func (bagDhcpv6DAddrAttrb *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses_BagDhcpv6DAddrAttrb) SetFilter(yf yfilter.YFilter) { bagDhcpv6DAddrAttrb.YFilter = yf }

func (bagDhcpv6DAddrAttrb *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses_BagDhcpv6DAddrAttrb) GetGoName(yname string) string {
    if yname == "prefix" { return "Prefix" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "lease-time" { return "LeaseTime" }
    if yname == "remaining-lease-time" { return "RemainingLeaseTime" }
    return ""
}

func (bagDhcpv6DAddrAttrb *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses_BagDhcpv6DAddrAttrb) GetSegmentPath() string {
    return "bag-dhcpv6d-addr-attrb"
}

func (bagDhcpv6DAddrAttrb *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses_BagDhcpv6DAddrAttrb) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bagDhcpv6DAddrAttrb *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses_BagDhcpv6DAddrAttrb) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bagDhcpv6DAddrAttrb *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses_BagDhcpv6DAddrAttrb) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix"] = bagDhcpv6DAddrAttrb.Prefix
    leafs["prefix-length"] = bagDhcpv6DAddrAttrb.PrefixLength
    leafs["lease-time"] = bagDhcpv6DAddrAttrb.LeaseTime
    leafs["remaining-lease-time"] = bagDhcpv6DAddrAttrb.RemainingLeaseTime
    return leafs
}

func (bagDhcpv6DAddrAttrb *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses_BagDhcpv6DAddrAttrb) GetBundleName() string { return "cisco_ios_xr" }

func (bagDhcpv6DAddrAttrb *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses_BagDhcpv6DAddrAttrb) GetYangName() string { return "bag-dhcpv6d-addr-attrb" }

func (bagDhcpv6DAddrAttrb *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses_BagDhcpv6DAddrAttrb) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bagDhcpv6DAddrAttrb *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses_BagDhcpv6DAddrAttrb) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bagDhcpv6DAddrAttrb *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses_BagDhcpv6DAddrAttrb) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bagDhcpv6DAddrAttrb *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses_BagDhcpv6DAddrAttrb) SetParent(parent types.Entity) { bagDhcpv6DAddrAttrb.parent = parent }

func (bagDhcpv6DAddrAttrb *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses_BagDhcpv6DAddrAttrb) GetParent() types.Entity { return bagDhcpv6DAddrAttrb.parent }

func (bagDhcpv6DAddrAttrb *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses_BagDhcpv6DAddrAttrb) GetParentYangName() string { return "addresses" }

// Dhcpv6_Nodes_Node_Server_Vrfs
// DHCPV6 server list of VRF names
type Dhcpv6_Nodes_Node_Server_Vrfs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 DHCP server VRF name. The type is slice of
    // Dhcpv6_Nodes_Node_Server_Vrfs_Vrf.
    Vrf []Dhcpv6_Nodes_Node_Server_Vrfs_Vrf
}

func (vrfs *Dhcpv6_Nodes_Node_Server_Vrfs) GetFilter() yfilter.YFilter { return vrfs.YFilter }

func (vrfs *Dhcpv6_Nodes_Node_Server_Vrfs) SetFilter(yf yfilter.YFilter) { vrfs.YFilter = yf }

func (vrfs *Dhcpv6_Nodes_Node_Server_Vrfs) GetGoName(yname string) string {
    if yname == "vrf" { return "Vrf" }
    return ""
}

func (vrfs *Dhcpv6_Nodes_Node_Server_Vrfs) GetSegmentPath() string {
    return "vrfs"
}

func (vrfs *Dhcpv6_Nodes_Node_Server_Vrfs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf" {
        for _, c := range vrfs.Vrf {
            if vrfs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dhcpv6_Nodes_Node_Server_Vrfs_Vrf{}
        vrfs.Vrf = append(vrfs.Vrf, child)
        return &vrfs.Vrf[len(vrfs.Vrf)-1]
    }
    return nil
}

func (vrfs *Dhcpv6_Nodes_Node_Server_Vrfs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vrfs.Vrf {
        children[vrfs.Vrf[i].GetSegmentPath()] = &vrfs.Vrf[i]
    }
    return children
}

func (vrfs *Dhcpv6_Nodes_Node_Server_Vrfs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vrfs *Dhcpv6_Nodes_Node_Server_Vrfs) GetBundleName() string { return "cisco_ios_xr" }

func (vrfs *Dhcpv6_Nodes_Node_Server_Vrfs) GetYangName() string { return "vrfs" }

func (vrfs *Dhcpv6_Nodes_Node_Server_Vrfs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfs *Dhcpv6_Nodes_Node_Server_Vrfs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfs *Dhcpv6_Nodes_Node_Server_Vrfs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfs *Dhcpv6_Nodes_Node_Server_Vrfs) SetParent(parent types.Entity) { vrfs.parent = parent }

func (vrfs *Dhcpv6_Nodes_Node_Server_Vrfs) GetParent() types.Entity { return vrfs.parent }

func (vrfs *Dhcpv6_Nodes_Node_Server_Vrfs) GetParentYangName() string { return "server" }

// Dhcpv6_Nodes_Node_Server_Vrfs_Vrf
// IPv6 DHCP server VRF name
type Dhcpv6_Nodes_Node_Server_Vrfs_Vrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // IPv6 DHCP server statistics.
    Statistics Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics
}

func (vrf *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf) GetFilter() yfilter.YFilter { return vrf.YFilter }

func (vrf *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf) SetFilter(yf yfilter.YFilter) { vrf.YFilter = yf }

func (vrf *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "statistics" { return "Statistics" }
    return ""
}

func (vrf *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf) GetSegmentPath() string {
    return "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
}

func (vrf *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "statistics" {
        return &vrf.Statistics
    }
    return nil
}

func (vrf *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["statistics"] = &vrf.Statistics
    return children
}

func (vrf *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = vrf.VrfName
    return leafs
}

func (vrf *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf) GetBundleName() string { return "cisco_ios_xr" }

func (vrf *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf) GetYangName() string { return "vrf" }

func (vrf *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrf *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrf *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrf *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf) SetParent(parent types.Entity) { vrf.parent = parent }

func (vrf *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf) GetParent() types.Entity { return vrf.parent }

func (vrf *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf) GetParentYangName() string { return "vrfs" }

// Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics
// IPv6 DHCP server statistics
type Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DHCPV6 solicit packets.
    Solicit Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Solicit

    // DHCPV6 advertise packets.
    Advertise Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Advertise

    // DHCPV6 request packets.
    Request Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Request

    // DHCPV6 reply packets.
    Reply Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Reply

    // DHCPV6 confirm packets.
    Confirm Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Confirm

    // DHCPV6 decline packets.
    Decline Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Decline

    // DHCPV6 renew packets.
    Renew Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Renew

    // DHCPV6 rebind packets.
    Rebind Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Rebind

    // DHCPV6 release packets.
    Release Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Release

    // DHCPV6 reconfig packets.
    Reconfig Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Reconfig

    // DHCPV6 inform packets.
    Inform Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Inform

    // DHCPV6 relay forward packets.
    RelayForward Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_RelayForward

    // DHCPV6 relay reply packets.
    RelayReply Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_RelayReply

    // DHCPV6 lease query packets.
    LeaseQuery Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQuery

    // DHCPV6 lease query reply packets.
    LeaseQueryReply Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQueryReply

    // DHCPV6 lease query done packets.
    LeaseQueryDone Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQueryDone

    // DHCPV6 lease query data packets.
    LeaseQueryData Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQueryData
}

func (statistics *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics) GetGoName(yname string) string {
    if yname == "solicit" { return "Solicit" }
    if yname == "advertise" { return "Advertise" }
    if yname == "request" { return "Request" }
    if yname == "reply" { return "Reply" }
    if yname == "confirm" { return "Confirm" }
    if yname == "decline" { return "Decline" }
    if yname == "renew" { return "Renew" }
    if yname == "rebind" { return "Rebind" }
    if yname == "release" { return "Release" }
    if yname == "reconfig" { return "Reconfig" }
    if yname == "inform" { return "Inform" }
    if yname == "relay-forward" { return "RelayForward" }
    if yname == "relay-reply" { return "RelayReply" }
    if yname == "lease-query" { return "LeaseQuery" }
    if yname == "lease-query-reply" { return "LeaseQueryReply" }
    if yname == "lease-query-done" { return "LeaseQueryDone" }
    if yname == "lease-query-data" { return "LeaseQueryData" }
    return ""
}

func (statistics *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "solicit" {
        return &statistics.Solicit
    }
    if childYangName == "advertise" {
        return &statistics.Advertise
    }
    if childYangName == "request" {
        return &statistics.Request
    }
    if childYangName == "reply" {
        return &statistics.Reply
    }
    if childYangName == "confirm" {
        return &statistics.Confirm
    }
    if childYangName == "decline" {
        return &statistics.Decline
    }
    if childYangName == "renew" {
        return &statistics.Renew
    }
    if childYangName == "rebind" {
        return &statistics.Rebind
    }
    if childYangName == "release" {
        return &statistics.Release
    }
    if childYangName == "reconfig" {
        return &statistics.Reconfig
    }
    if childYangName == "inform" {
        return &statistics.Inform
    }
    if childYangName == "relay-forward" {
        return &statistics.RelayForward
    }
    if childYangName == "relay-reply" {
        return &statistics.RelayReply
    }
    if childYangName == "lease-query" {
        return &statistics.LeaseQuery
    }
    if childYangName == "lease-query-reply" {
        return &statistics.LeaseQueryReply
    }
    if childYangName == "lease-query-done" {
        return &statistics.LeaseQueryDone
    }
    if childYangName == "lease-query-data" {
        return &statistics.LeaseQueryData
    }
    return nil
}

func (statistics *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["solicit"] = &statistics.Solicit
    children["advertise"] = &statistics.Advertise
    children["request"] = &statistics.Request
    children["reply"] = &statistics.Reply
    children["confirm"] = &statistics.Confirm
    children["decline"] = &statistics.Decline
    children["renew"] = &statistics.Renew
    children["rebind"] = &statistics.Rebind
    children["release"] = &statistics.Release
    children["reconfig"] = &statistics.Reconfig
    children["inform"] = &statistics.Inform
    children["relay-forward"] = &statistics.RelayForward
    children["relay-reply"] = &statistics.RelayReply
    children["lease-query"] = &statistics.LeaseQuery
    children["lease-query-reply"] = &statistics.LeaseQueryReply
    children["lease-query-done"] = &statistics.LeaseQueryDone
    children["lease-query-data"] = &statistics.LeaseQueryData
    return children
}

func (statistics *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (statistics *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics) GetYangName() string { return "statistics" }

func (statistics *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics) GetParentYangName() string { return "vrf" }

// Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Solicit
// DHCPV6 solicit packets
type Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Solicit struct {
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

func (solicit *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Solicit) GetFilter() yfilter.YFilter { return solicit.YFilter }

func (solicit *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Solicit) SetFilter(yf yfilter.YFilter) { solicit.YFilter = yf }

func (solicit *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Solicit) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (solicit *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Solicit) GetSegmentPath() string {
    return "solicit"
}

func (solicit *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Solicit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (solicit *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Solicit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (solicit *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Solicit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = solicit.ReceivedPackets
    leafs["transmitted-packets"] = solicit.TransmittedPackets
    leafs["dropped-packets"] = solicit.DroppedPackets
    return leafs
}

func (solicit *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Solicit) GetBundleName() string { return "cisco_ios_xr" }

func (solicit *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Solicit) GetYangName() string { return "solicit" }

func (solicit *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Solicit) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (solicit *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Solicit) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (solicit *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Solicit) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (solicit *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Solicit) SetParent(parent types.Entity) { solicit.parent = parent }

func (solicit *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Solicit) GetParent() types.Entity { return solicit.parent }

func (solicit *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Solicit) GetParentYangName() string { return "statistics" }

// Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Advertise
// DHCPV6 advertise packets
type Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Advertise struct {
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

func (advertise *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Advertise) GetFilter() yfilter.YFilter { return advertise.YFilter }

func (advertise *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Advertise) SetFilter(yf yfilter.YFilter) { advertise.YFilter = yf }

func (advertise *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Advertise) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (advertise *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Advertise) GetSegmentPath() string {
    return "advertise"
}

func (advertise *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Advertise) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (advertise *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Advertise) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (advertise *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Advertise) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = advertise.ReceivedPackets
    leafs["transmitted-packets"] = advertise.TransmittedPackets
    leafs["dropped-packets"] = advertise.DroppedPackets
    return leafs
}

func (advertise *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Advertise) GetBundleName() string { return "cisco_ios_xr" }

func (advertise *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Advertise) GetYangName() string { return "advertise" }

func (advertise *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Advertise) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (advertise *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Advertise) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (advertise *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Advertise) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (advertise *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Advertise) SetParent(parent types.Entity) { advertise.parent = parent }

func (advertise *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Advertise) GetParent() types.Entity { return advertise.parent }

func (advertise *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Advertise) GetParentYangName() string { return "statistics" }

// Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Request
// DHCPV6 request packets
type Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Request struct {
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

func (request *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Request) GetFilter() yfilter.YFilter { return request.YFilter }

func (request *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Request) SetFilter(yf yfilter.YFilter) { request.YFilter = yf }

func (request *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Request) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (request *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Request) GetSegmentPath() string {
    return "request"
}

func (request *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Request) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (request *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Request) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (request *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Request) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = request.ReceivedPackets
    leafs["transmitted-packets"] = request.TransmittedPackets
    leafs["dropped-packets"] = request.DroppedPackets
    return leafs
}

func (request *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Request) GetBundleName() string { return "cisco_ios_xr" }

func (request *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Request) GetYangName() string { return "request" }

func (request *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Request) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (request *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Request) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (request *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Request) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (request *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Request) SetParent(parent types.Entity) { request.parent = parent }

func (request *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Request) GetParent() types.Entity { return request.parent }

func (request *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Request) GetParentYangName() string { return "statistics" }

// Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Reply
// DHCPV6 reply packets
type Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Reply struct {
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

func (reply *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Reply) GetFilter() yfilter.YFilter { return reply.YFilter }

func (reply *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Reply) SetFilter(yf yfilter.YFilter) { reply.YFilter = yf }

func (reply *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Reply) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (reply *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Reply) GetSegmentPath() string {
    return "reply"
}

func (reply *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Reply) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (reply *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Reply) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (reply *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Reply) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = reply.ReceivedPackets
    leafs["transmitted-packets"] = reply.TransmittedPackets
    leafs["dropped-packets"] = reply.DroppedPackets
    return leafs
}

func (reply *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Reply) GetBundleName() string { return "cisco_ios_xr" }

func (reply *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Reply) GetYangName() string { return "reply" }

func (reply *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Reply) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (reply *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Reply) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (reply *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Reply) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (reply *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Reply) SetParent(parent types.Entity) { reply.parent = parent }

func (reply *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Reply) GetParent() types.Entity { return reply.parent }

func (reply *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Reply) GetParentYangName() string { return "statistics" }

// Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Confirm
// DHCPV6 confirm packets
type Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Confirm struct {
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

func (confirm *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Confirm) GetFilter() yfilter.YFilter { return confirm.YFilter }

func (confirm *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Confirm) SetFilter(yf yfilter.YFilter) { confirm.YFilter = yf }

func (confirm *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Confirm) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (confirm *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Confirm) GetSegmentPath() string {
    return "confirm"
}

func (confirm *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Confirm) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (confirm *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Confirm) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (confirm *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Confirm) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = confirm.ReceivedPackets
    leafs["transmitted-packets"] = confirm.TransmittedPackets
    leafs["dropped-packets"] = confirm.DroppedPackets
    return leafs
}

func (confirm *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Confirm) GetBundleName() string { return "cisco_ios_xr" }

func (confirm *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Confirm) GetYangName() string { return "confirm" }

func (confirm *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Confirm) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (confirm *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Confirm) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (confirm *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Confirm) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (confirm *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Confirm) SetParent(parent types.Entity) { confirm.parent = parent }

func (confirm *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Confirm) GetParent() types.Entity { return confirm.parent }

func (confirm *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Confirm) GetParentYangName() string { return "statistics" }

// Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Decline
// DHCPV6 decline packets
type Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Decline struct {
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

func (decline *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Decline) GetFilter() yfilter.YFilter { return decline.YFilter }

func (decline *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Decline) SetFilter(yf yfilter.YFilter) { decline.YFilter = yf }

func (decline *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Decline) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (decline *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Decline) GetSegmentPath() string {
    return "decline"
}

func (decline *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Decline) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (decline *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Decline) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (decline *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Decline) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = decline.ReceivedPackets
    leafs["transmitted-packets"] = decline.TransmittedPackets
    leafs["dropped-packets"] = decline.DroppedPackets
    return leafs
}

func (decline *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Decline) GetBundleName() string { return "cisco_ios_xr" }

func (decline *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Decline) GetYangName() string { return "decline" }

func (decline *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Decline) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (decline *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Decline) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (decline *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Decline) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (decline *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Decline) SetParent(parent types.Entity) { decline.parent = parent }

func (decline *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Decline) GetParent() types.Entity { return decline.parent }

func (decline *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Decline) GetParentYangName() string { return "statistics" }

// Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Renew
// DHCPV6 renew packets
type Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Renew struct {
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

func (renew *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Renew) GetFilter() yfilter.YFilter { return renew.YFilter }

func (renew *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Renew) SetFilter(yf yfilter.YFilter) { renew.YFilter = yf }

func (renew *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Renew) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (renew *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Renew) GetSegmentPath() string {
    return "renew"
}

func (renew *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Renew) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (renew *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Renew) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (renew *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Renew) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = renew.ReceivedPackets
    leafs["transmitted-packets"] = renew.TransmittedPackets
    leafs["dropped-packets"] = renew.DroppedPackets
    return leafs
}

func (renew *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Renew) GetBundleName() string { return "cisco_ios_xr" }

func (renew *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Renew) GetYangName() string { return "renew" }

func (renew *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Renew) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (renew *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Renew) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (renew *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Renew) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (renew *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Renew) SetParent(parent types.Entity) { renew.parent = parent }

func (renew *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Renew) GetParent() types.Entity { return renew.parent }

func (renew *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Renew) GetParentYangName() string { return "statistics" }

// Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Rebind
// DHCPV6 rebind packets
type Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Rebind struct {
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

func (rebind *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Rebind) GetFilter() yfilter.YFilter { return rebind.YFilter }

func (rebind *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Rebind) SetFilter(yf yfilter.YFilter) { rebind.YFilter = yf }

func (rebind *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Rebind) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (rebind *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Rebind) GetSegmentPath() string {
    return "rebind"
}

func (rebind *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Rebind) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rebind *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Rebind) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rebind *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Rebind) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = rebind.ReceivedPackets
    leafs["transmitted-packets"] = rebind.TransmittedPackets
    leafs["dropped-packets"] = rebind.DroppedPackets
    return leafs
}

func (rebind *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Rebind) GetBundleName() string { return "cisco_ios_xr" }

func (rebind *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Rebind) GetYangName() string { return "rebind" }

func (rebind *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Rebind) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rebind *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Rebind) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rebind *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Rebind) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rebind *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Rebind) SetParent(parent types.Entity) { rebind.parent = parent }

func (rebind *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Rebind) GetParent() types.Entity { return rebind.parent }

func (rebind *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Rebind) GetParentYangName() string { return "statistics" }

// Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Release
// DHCPV6 release packets
type Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Release struct {
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

func (release *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Release) GetFilter() yfilter.YFilter { return release.YFilter }

func (release *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Release) SetFilter(yf yfilter.YFilter) { release.YFilter = yf }

func (release *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Release) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (release *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Release) GetSegmentPath() string {
    return "release"
}

func (release *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Release) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (release *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Release) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (release *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Release) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = release.ReceivedPackets
    leafs["transmitted-packets"] = release.TransmittedPackets
    leafs["dropped-packets"] = release.DroppedPackets
    return leafs
}

func (release *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Release) GetBundleName() string { return "cisco_ios_xr" }

func (release *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Release) GetYangName() string { return "release" }

func (release *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Release) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (release *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Release) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (release *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Release) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (release *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Release) SetParent(parent types.Entity) { release.parent = parent }

func (release *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Release) GetParent() types.Entity { return release.parent }

func (release *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Release) GetParentYangName() string { return "statistics" }

// Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Reconfig
// DHCPV6 reconfig packets
type Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Reconfig struct {
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

func (reconfig *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Reconfig) GetFilter() yfilter.YFilter { return reconfig.YFilter }

func (reconfig *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Reconfig) SetFilter(yf yfilter.YFilter) { reconfig.YFilter = yf }

func (reconfig *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Reconfig) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (reconfig *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Reconfig) GetSegmentPath() string {
    return "reconfig"
}

func (reconfig *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Reconfig) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (reconfig *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Reconfig) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (reconfig *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Reconfig) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = reconfig.ReceivedPackets
    leafs["transmitted-packets"] = reconfig.TransmittedPackets
    leafs["dropped-packets"] = reconfig.DroppedPackets
    return leafs
}

func (reconfig *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Reconfig) GetBundleName() string { return "cisco_ios_xr" }

func (reconfig *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Reconfig) GetYangName() string { return "reconfig" }

func (reconfig *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Reconfig) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (reconfig *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Reconfig) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (reconfig *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Reconfig) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (reconfig *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Reconfig) SetParent(parent types.Entity) { reconfig.parent = parent }

func (reconfig *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Reconfig) GetParent() types.Entity { return reconfig.parent }

func (reconfig *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Reconfig) GetParentYangName() string { return "statistics" }

// Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Inform
// DHCPV6 inform packets
type Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Inform struct {
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

func (inform *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Inform) GetFilter() yfilter.YFilter { return inform.YFilter }

func (inform *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Inform) SetFilter(yf yfilter.YFilter) { inform.YFilter = yf }

func (inform *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Inform) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (inform *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Inform) GetSegmentPath() string {
    return "inform"
}

func (inform *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Inform) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inform *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Inform) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inform *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Inform) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = inform.ReceivedPackets
    leafs["transmitted-packets"] = inform.TransmittedPackets
    leafs["dropped-packets"] = inform.DroppedPackets
    return leafs
}

func (inform *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Inform) GetBundleName() string { return "cisco_ios_xr" }

func (inform *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Inform) GetYangName() string { return "inform" }

func (inform *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Inform) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inform *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Inform) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inform *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Inform) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inform *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Inform) SetParent(parent types.Entity) { inform.parent = parent }

func (inform *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Inform) GetParent() types.Entity { return inform.parent }

func (inform *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Inform) GetParentYangName() string { return "statistics" }

// Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_RelayForward
// DHCPV6 relay forward packets
type Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_RelayForward struct {
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

func (relayForward *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_RelayForward) GetFilter() yfilter.YFilter { return relayForward.YFilter }

func (relayForward *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_RelayForward) SetFilter(yf yfilter.YFilter) { relayForward.YFilter = yf }

func (relayForward *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_RelayForward) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (relayForward *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_RelayForward) GetSegmentPath() string {
    return "relay-forward"
}

func (relayForward *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_RelayForward) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (relayForward *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_RelayForward) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (relayForward *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_RelayForward) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = relayForward.ReceivedPackets
    leafs["transmitted-packets"] = relayForward.TransmittedPackets
    leafs["dropped-packets"] = relayForward.DroppedPackets
    return leafs
}

func (relayForward *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_RelayForward) GetBundleName() string { return "cisco_ios_xr" }

func (relayForward *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_RelayForward) GetYangName() string { return "relay-forward" }

func (relayForward *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_RelayForward) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (relayForward *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_RelayForward) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (relayForward *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_RelayForward) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (relayForward *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_RelayForward) SetParent(parent types.Entity) { relayForward.parent = parent }

func (relayForward *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_RelayForward) GetParent() types.Entity { return relayForward.parent }

func (relayForward *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_RelayForward) GetParentYangName() string { return "statistics" }

// Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_RelayReply
// DHCPV6 relay reply packets
type Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_RelayReply struct {
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

func (relayReply *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_RelayReply) GetFilter() yfilter.YFilter { return relayReply.YFilter }

func (relayReply *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_RelayReply) SetFilter(yf yfilter.YFilter) { relayReply.YFilter = yf }

func (relayReply *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_RelayReply) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (relayReply *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_RelayReply) GetSegmentPath() string {
    return "relay-reply"
}

func (relayReply *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_RelayReply) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (relayReply *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_RelayReply) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (relayReply *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_RelayReply) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = relayReply.ReceivedPackets
    leafs["transmitted-packets"] = relayReply.TransmittedPackets
    leafs["dropped-packets"] = relayReply.DroppedPackets
    return leafs
}

func (relayReply *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_RelayReply) GetBundleName() string { return "cisco_ios_xr" }

func (relayReply *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_RelayReply) GetYangName() string { return "relay-reply" }

func (relayReply *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_RelayReply) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (relayReply *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_RelayReply) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (relayReply *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_RelayReply) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (relayReply *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_RelayReply) SetParent(parent types.Entity) { relayReply.parent = parent }

func (relayReply *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_RelayReply) GetParent() types.Entity { return relayReply.parent }

func (relayReply *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_RelayReply) GetParentYangName() string { return "statistics" }

// Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQuery
// DHCPV6 lease query packets
type Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQuery struct {
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

func (leaseQuery *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQuery) GetFilter() yfilter.YFilter { return leaseQuery.YFilter }

func (leaseQuery *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQuery) SetFilter(yf yfilter.YFilter) { leaseQuery.YFilter = yf }

func (leaseQuery *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQuery) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (leaseQuery *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQuery) GetSegmentPath() string {
    return "lease-query"
}

func (leaseQuery *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQuery) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (leaseQuery *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQuery) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (leaseQuery *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQuery) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = leaseQuery.ReceivedPackets
    leafs["transmitted-packets"] = leaseQuery.TransmittedPackets
    leafs["dropped-packets"] = leaseQuery.DroppedPackets
    return leafs
}

func (leaseQuery *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQuery) GetBundleName() string { return "cisco_ios_xr" }

func (leaseQuery *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQuery) GetYangName() string { return "lease-query" }

func (leaseQuery *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQuery) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (leaseQuery *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQuery) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (leaseQuery *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQuery) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (leaseQuery *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQuery) SetParent(parent types.Entity) { leaseQuery.parent = parent }

func (leaseQuery *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQuery) GetParent() types.Entity { return leaseQuery.parent }

func (leaseQuery *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQuery) GetParentYangName() string { return "statistics" }

// Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQueryReply
// DHCPV6 lease query reply packets
type Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQueryReply struct {
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

func (leaseQueryReply *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQueryReply) GetFilter() yfilter.YFilter { return leaseQueryReply.YFilter }

func (leaseQueryReply *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQueryReply) SetFilter(yf yfilter.YFilter) { leaseQueryReply.YFilter = yf }

func (leaseQueryReply *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQueryReply) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (leaseQueryReply *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQueryReply) GetSegmentPath() string {
    return "lease-query-reply"
}

func (leaseQueryReply *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQueryReply) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (leaseQueryReply *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQueryReply) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (leaseQueryReply *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQueryReply) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = leaseQueryReply.ReceivedPackets
    leafs["transmitted-packets"] = leaseQueryReply.TransmittedPackets
    leafs["dropped-packets"] = leaseQueryReply.DroppedPackets
    return leafs
}

func (leaseQueryReply *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQueryReply) GetBundleName() string { return "cisco_ios_xr" }

func (leaseQueryReply *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQueryReply) GetYangName() string { return "lease-query-reply" }

func (leaseQueryReply *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQueryReply) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (leaseQueryReply *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQueryReply) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (leaseQueryReply *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQueryReply) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (leaseQueryReply *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQueryReply) SetParent(parent types.Entity) { leaseQueryReply.parent = parent }

func (leaseQueryReply *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQueryReply) GetParent() types.Entity { return leaseQueryReply.parent }

func (leaseQueryReply *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQueryReply) GetParentYangName() string { return "statistics" }

// Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQueryDone
// DHCPV6 lease query done packets
type Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQueryDone struct {
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

func (leaseQueryDone *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQueryDone) GetFilter() yfilter.YFilter { return leaseQueryDone.YFilter }

func (leaseQueryDone *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQueryDone) SetFilter(yf yfilter.YFilter) { leaseQueryDone.YFilter = yf }

func (leaseQueryDone *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQueryDone) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (leaseQueryDone *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQueryDone) GetSegmentPath() string {
    return "lease-query-done"
}

func (leaseQueryDone *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQueryDone) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (leaseQueryDone *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQueryDone) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (leaseQueryDone *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQueryDone) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = leaseQueryDone.ReceivedPackets
    leafs["transmitted-packets"] = leaseQueryDone.TransmittedPackets
    leafs["dropped-packets"] = leaseQueryDone.DroppedPackets
    return leafs
}

func (leaseQueryDone *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQueryDone) GetBundleName() string { return "cisco_ios_xr" }

func (leaseQueryDone *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQueryDone) GetYangName() string { return "lease-query-done" }

func (leaseQueryDone *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQueryDone) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (leaseQueryDone *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQueryDone) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (leaseQueryDone *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQueryDone) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (leaseQueryDone *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQueryDone) SetParent(parent types.Entity) { leaseQueryDone.parent = parent }

func (leaseQueryDone *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQueryDone) GetParent() types.Entity { return leaseQueryDone.parent }

func (leaseQueryDone *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQueryDone) GetParentYangName() string { return "statistics" }

// Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQueryData
// DHCPV6 lease query data packets
type Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQueryData struct {
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

func (leaseQueryData *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQueryData) GetFilter() yfilter.YFilter { return leaseQueryData.YFilter }

func (leaseQueryData *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQueryData) SetFilter(yf yfilter.YFilter) { leaseQueryData.YFilter = yf }

func (leaseQueryData *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQueryData) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (leaseQueryData *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQueryData) GetSegmentPath() string {
    return "lease-query-data"
}

func (leaseQueryData *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQueryData) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (leaseQueryData *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQueryData) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (leaseQueryData *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQueryData) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = leaseQueryData.ReceivedPackets
    leafs["transmitted-packets"] = leaseQueryData.TransmittedPackets
    leafs["dropped-packets"] = leaseQueryData.DroppedPackets
    return leafs
}

func (leaseQueryData *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQueryData) GetBundleName() string { return "cisco_ios_xr" }

func (leaseQueryData *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQueryData) GetYangName() string { return "lease-query-data" }

func (leaseQueryData *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQueryData) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (leaseQueryData *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQueryData) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (leaseQueryData *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQueryData) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (leaseQueryData *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQueryData) SetParent(parent types.Entity) { leaseQueryData.parent = parent }

func (leaseQueryData *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQueryData) GetParent() types.Entity { return leaseQueryData.parent }

func (leaseQueryData *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQueryData) GetParentYangName() string { return "statistics" }

// Dhcpv6_Nodes_Node_Server_Profiles
// IPv6 DHCP server profile
type Dhcpv6_Nodes_Node_Server_Profiles struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 DHCP server profile. The type is slice of
    // Dhcpv6_Nodes_Node_Server_Profiles_Profile.
    Profile []Dhcpv6_Nodes_Node_Server_Profiles_Profile
}

func (profiles *Dhcpv6_Nodes_Node_Server_Profiles) GetFilter() yfilter.YFilter { return profiles.YFilter }

func (profiles *Dhcpv6_Nodes_Node_Server_Profiles) SetFilter(yf yfilter.YFilter) { profiles.YFilter = yf }

func (profiles *Dhcpv6_Nodes_Node_Server_Profiles) GetGoName(yname string) string {
    if yname == "profile" { return "Profile" }
    return ""
}

func (profiles *Dhcpv6_Nodes_Node_Server_Profiles) GetSegmentPath() string {
    return "profiles"
}

func (profiles *Dhcpv6_Nodes_Node_Server_Profiles) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "profile" {
        for _, c := range profiles.Profile {
            if profiles.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dhcpv6_Nodes_Node_Server_Profiles_Profile{}
        profiles.Profile = append(profiles.Profile, child)
        return &profiles.Profile[len(profiles.Profile)-1]
    }
    return nil
}

func (profiles *Dhcpv6_Nodes_Node_Server_Profiles) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range profiles.Profile {
        children[profiles.Profile[i].GetSegmentPath()] = &profiles.Profile[i]
    }
    return children
}

func (profiles *Dhcpv6_Nodes_Node_Server_Profiles) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (profiles *Dhcpv6_Nodes_Node_Server_Profiles) GetBundleName() string { return "cisco_ios_xr" }

func (profiles *Dhcpv6_Nodes_Node_Server_Profiles) GetYangName() string { return "profiles" }

func (profiles *Dhcpv6_Nodes_Node_Server_Profiles) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (profiles *Dhcpv6_Nodes_Node_Server_Profiles) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (profiles *Dhcpv6_Nodes_Node_Server_Profiles) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (profiles *Dhcpv6_Nodes_Node_Server_Profiles) SetParent(parent types.Entity) { profiles.parent = parent }

func (profiles *Dhcpv6_Nodes_Node_Server_Profiles) GetParent() types.Entity { return profiles.parent }

func (profiles *Dhcpv6_Nodes_Node_Server_Profiles) GetParentYangName() string { return "server" }

// Dhcpv6_Nodes_Node_Server_Profiles_Profile
// IPv6 DHCP server profile
type Dhcpv6_Nodes_Node_Server_Profiles_Profile struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Profile name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    ProfileName interface{}

    // IPv6 DHCP server profile Info.
    Info Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info

    // DHCPV6 throttle table.
    ThrottleInfos Dhcpv6_Nodes_Node_Server_Profiles_Profile_ThrottleInfos
}

func (profile *Dhcpv6_Nodes_Node_Server_Profiles_Profile) GetFilter() yfilter.YFilter { return profile.YFilter }

func (profile *Dhcpv6_Nodes_Node_Server_Profiles_Profile) SetFilter(yf yfilter.YFilter) { profile.YFilter = yf }

func (profile *Dhcpv6_Nodes_Node_Server_Profiles_Profile) GetGoName(yname string) string {
    if yname == "profile-name" { return "ProfileName" }
    if yname == "info" { return "Info" }
    if yname == "throttle-infos" { return "ThrottleInfos" }
    return ""
}

func (profile *Dhcpv6_Nodes_Node_Server_Profiles_Profile) GetSegmentPath() string {
    return "profile" + "[profile-name='" + fmt.Sprintf("%v", profile.ProfileName) + "']"
}

func (profile *Dhcpv6_Nodes_Node_Server_Profiles_Profile) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "info" {
        return &profile.Info
    }
    if childYangName == "throttle-infos" {
        return &profile.ThrottleInfos
    }
    return nil
}

func (profile *Dhcpv6_Nodes_Node_Server_Profiles_Profile) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["info"] = &profile.Info
    children["throttle-infos"] = &profile.ThrottleInfos
    return children
}

func (profile *Dhcpv6_Nodes_Node_Server_Profiles_Profile) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["profile-name"] = profile.ProfileName
    return leafs
}

func (profile *Dhcpv6_Nodes_Node_Server_Profiles_Profile) GetBundleName() string { return "cisco_ios_xr" }

func (profile *Dhcpv6_Nodes_Node_Server_Profiles_Profile) GetYangName() string { return "profile" }

func (profile *Dhcpv6_Nodes_Node_Server_Profiles_Profile) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (profile *Dhcpv6_Nodes_Node_Server_Profiles_Profile) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (profile *Dhcpv6_Nodes_Node_Server_Profiles_Profile) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (profile *Dhcpv6_Nodes_Node_Server_Profiles_Profile) SetParent(parent types.Entity) { profile.parent = parent }

func (profile *Dhcpv6_Nodes_Node_Server_Profiles_Profile) GetParent() types.Entity { return profile.parent }

func (profile *Dhcpv6_Nodes_Node_Server_Profiles_Profile) GetParentYangName() string { return "profiles" }

// Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info
// IPv6 DHCP server profile Info
type Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Server profile name. The type is string with length: 0..65.
    ProfileName interface{}

    // Server domain name. The type is string with length: 0..65.
    DomainName interface{}

    // DNS address count. The type is interface{} with range: 0..255.
    ProfileDns interface{}

    // Server aftr name. The type is string with length: 0..65.
    AftrName interface{}

    // Server framed address pool name. The type is string with length: 0..65.
    FramedAddrPoolName interface{}

    // Server delegated prefix pool name. The type is string with length: 0..65.
    DelegatedPrefixPoolName interface{}

    // Rapid Commit. The type is bool.
    RapidCommit interface{}

    // DNS addresses. The type is slice of string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    ProfileDnsAddress []interface{}

    // Server lease time.
    Lease Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_Lease

    // Interface references.
    InterfaceReferences Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_InterfaceReferences
}

func (info *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info) GetFilter() yfilter.YFilter { return info.YFilter }

func (info *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info) SetFilter(yf yfilter.YFilter) { info.YFilter = yf }

func (info *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info) GetGoName(yname string) string {
    if yname == "profile-name" { return "ProfileName" }
    if yname == "domain-name" { return "DomainName" }
    if yname == "profile-dns" { return "ProfileDns" }
    if yname == "aftr-name" { return "AftrName" }
    if yname == "framed-addr-pool-name" { return "FramedAddrPoolName" }
    if yname == "delegated-prefix-pool-name" { return "DelegatedPrefixPoolName" }
    if yname == "rapid-commit" { return "RapidCommit" }
    if yname == "profile-dns-address" { return "ProfileDnsAddress" }
    if yname == "lease" { return "Lease" }
    if yname == "interface-references" { return "InterfaceReferences" }
    return ""
}

func (info *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info) GetSegmentPath() string {
    return "info"
}

func (info *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lease" {
        return &info.Lease
    }
    if childYangName == "interface-references" {
        return &info.InterfaceReferences
    }
    return nil
}

func (info *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["lease"] = &info.Lease
    children["interface-references"] = &info.InterfaceReferences
    return children
}

func (info *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["profile-name"] = info.ProfileName
    leafs["domain-name"] = info.DomainName
    leafs["profile-dns"] = info.ProfileDns
    leafs["aftr-name"] = info.AftrName
    leafs["framed-addr-pool-name"] = info.FramedAddrPoolName
    leafs["delegated-prefix-pool-name"] = info.DelegatedPrefixPoolName
    leafs["rapid-commit"] = info.RapidCommit
    leafs["profile-dns-address"] = info.ProfileDnsAddress
    return leafs
}

func (info *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info) GetBundleName() string { return "cisco_ios_xr" }

func (info *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info) GetYangName() string { return "info" }

func (info *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (info *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (info *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (info *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info) SetParent(parent types.Entity) { info.parent = parent }

func (info *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info) GetParent() types.Entity { return info.parent }

func (info *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info) GetParentYangName() string { return "profile" }

// Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_Lease
// Server lease time
type Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_Lease struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DHCPV6 client lease in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    Seconds interface{}

    // Time in format HH:MM:SS. The type is string with length: 0..10.
    Time interface{}
}

func (lease *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_Lease) GetFilter() yfilter.YFilter { return lease.YFilter }

func (lease *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_Lease) SetFilter(yf yfilter.YFilter) { lease.YFilter = yf }

func (lease *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_Lease) GetGoName(yname string) string {
    if yname == "seconds" { return "Seconds" }
    if yname == "time" { return "Time" }
    return ""
}

func (lease *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_Lease) GetSegmentPath() string {
    return "lease"
}

func (lease *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_Lease) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lease *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_Lease) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lease *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_Lease) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["seconds"] = lease.Seconds
    leafs["time"] = lease.Time
    return leafs
}

func (lease *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_Lease) GetBundleName() string { return "cisco_ios_xr" }

func (lease *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_Lease) GetYangName() string { return "lease" }

func (lease *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_Lease) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lease *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_Lease) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lease *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_Lease) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lease *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_Lease) SetParent(parent types.Entity) { lease.parent = parent }

func (lease *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_Lease) GetParent() types.Entity { return lease.parent }

func (lease *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_Lease) GetParentYangName() string { return "info" }

// Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_InterfaceReferences
// Interface references
type Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_InterfaceReferences struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ipv6 dhcpv6d server interface reference. The type is slice of
    // Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_InterfaceReferences_Ipv6Dhcpv6DServerInterfaceReference.
    Ipv6Dhcpv6DServerInterfaceReference []Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_InterfaceReferences_Ipv6Dhcpv6DServerInterfaceReference
}

func (interfaceReferences *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_InterfaceReferences) GetFilter() yfilter.YFilter { return interfaceReferences.YFilter }

func (interfaceReferences *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_InterfaceReferences) SetFilter(yf yfilter.YFilter) { interfaceReferences.YFilter = yf }

func (interfaceReferences *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_InterfaceReferences) GetGoName(yname string) string {
    if yname == "ipv6-dhcpv6d-server-interface-reference" { return "Ipv6Dhcpv6DServerInterfaceReference" }
    return ""
}

func (interfaceReferences *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_InterfaceReferences) GetSegmentPath() string {
    return "interface-references"
}

func (interfaceReferences *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_InterfaceReferences) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv6-dhcpv6d-server-interface-reference" {
        for _, c := range interfaceReferences.Ipv6Dhcpv6DServerInterfaceReference {
            if interfaceReferences.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_InterfaceReferences_Ipv6Dhcpv6DServerInterfaceReference{}
        interfaceReferences.Ipv6Dhcpv6DServerInterfaceReference = append(interfaceReferences.Ipv6Dhcpv6DServerInterfaceReference, child)
        return &interfaceReferences.Ipv6Dhcpv6DServerInterfaceReference[len(interfaceReferences.Ipv6Dhcpv6DServerInterfaceReference)-1]
    }
    return nil
}

func (interfaceReferences *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_InterfaceReferences) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaceReferences.Ipv6Dhcpv6DServerInterfaceReference {
        children[interfaceReferences.Ipv6Dhcpv6DServerInterfaceReference[i].GetSegmentPath()] = &interfaceReferences.Ipv6Dhcpv6DServerInterfaceReference[i]
    }
    return children
}

func (interfaceReferences *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_InterfaceReferences) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceReferences *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_InterfaceReferences) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceReferences *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_InterfaceReferences) GetYangName() string { return "interface-references" }

func (interfaceReferences *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_InterfaceReferences) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceReferences *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_InterfaceReferences) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceReferences *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_InterfaceReferences) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceReferences *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_InterfaceReferences) SetParent(parent types.Entity) { interfaceReferences.parent = parent }

func (interfaceReferences *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_InterfaceReferences) GetParent() types.Entity { return interfaceReferences.parent }

func (interfaceReferences *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_InterfaceReferences) GetParentYangName() string { return "info" }

// Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_InterfaceReferences_Ipv6Dhcpv6DServerInterfaceReference
// ipv6 dhcpv6d server interface reference
type Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_InterfaceReferences_Ipv6Dhcpv6DServerInterfaceReference struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface name. The type is string with length: 0..65.
    ServerReferenceInterfaceName interface{}
}

func (ipv6Dhcpv6DServerInterfaceReference *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_InterfaceReferences_Ipv6Dhcpv6DServerInterfaceReference) GetFilter() yfilter.YFilter { return ipv6Dhcpv6DServerInterfaceReference.YFilter }

func (ipv6Dhcpv6DServerInterfaceReference *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_InterfaceReferences_Ipv6Dhcpv6DServerInterfaceReference) SetFilter(yf yfilter.YFilter) { ipv6Dhcpv6DServerInterfaceReference.YFilter = yf }

func (ipv6Dhcpv6DServerInterfaceReference *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_InterfaceReferences_Ipv6Dhcpv6DServerInterfaceReference) GetGoName(yname string) string {
    if yname == "server-reference-interface-name" { return "ServerReferenceInterfaceName" }
    return ""
}

func (ipv6Dhcpv6DServerInterfaceReference *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_InterfaceReferences_Ipv6Dhcpv6DServerInterfaceReference) GetSegmentPath() string {
    return "ipv6-dhcpv6d-server-interface-reference"
}

func (ipv6Dhcpv6DServerInterfaceReference *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_InterfaceReferences_Ipv6Dhcpv6DServerInterfaceReference) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv6Dhcpv6DServerInterfaceReference *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_InterfaceReferences_Ipv6Dhcpv6DServerInterfaceReference) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv6Dhcpv6DServerInterfaceReference *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_InterfaceReferences_Ipv6Dhcpv6DServerInterfaceReference) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["server-reference-interface-name"] = ipv6Dhcpv6DServerInterfaceReference.ServerReferenceInterfaceName
    return leafs
}

func (ipv6Dhcpv6DServerInterfaceReference *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_InterfaceReferences_Ipv6Dhcpv6DServerInterfaceReference) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6Dhcpv6DServerInterfaceReference *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_InterfaceReferences_Ipv6Dhcpv6DServerInterfaceReference) GetYangName() string { return "ipv6-dhcpv6d-server-interface-reference" }

func (ipv6Dhcpv6DServerInterfaceReference *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_InterfaceReferences_Ipv6Dhcpv6DServerInterfaceReference) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6Dhcpv6DServerInterfaceReference *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_InterfaceReferences_Ipv6Dhcpv6DServerInterfaceReference) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6Dhcpv6DServerInterfaceReference *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_InterfaceReferences_Ipv6Dhcpv6DServerInterfaceReference) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6Dhcpv6DServerInterfaceReference *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_InterfaceReferences_Ipv6Dhcpv6DServerInterfaceReference) SetParent(parent types.Entity) { ipv6Dhcpv6DServerInterfaceReference.parent = parent }

func (ipv6Dhcpv6DServerInterfaceReference *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_InterfaceReferences_Ipv6Dhcpv6DServerInterfaceReference) GetParent() types.Entity { return ipv6Dhcpv6DServerInterfaceReference.parent }

func (ipv6Dhcpv6DServerInterfaceReference *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_InterfaceReferences_Ipv6Dhcpv6DServerInterfaceReference) GetParentYangName() string { return "interface-references" }

// Dhcpv6_Nodes_Node_Server_Profiles_Profile_ThrottleInfos
// DHCPV6 throttle table
type Dhcpv6_Nodes_Node_Server_Profiles_Profile_ThrottleInfos struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 DHCP server profile Throttle Info. The type is slice of
    // Dhcpv6_Nodes_Node_Server_Profiles_Profile_ThrottleInfos_ThrottleInfo.
    ThrottleInfo []Dhcpv6_Nodes_Node_Server_Profiles_Profile_ThrottleInfos_ThrottleInfo
}

func (throttleInfos *Dhcpv6_Nodes_Node_Server_Profiles_Profile_ThrottleInfos) GetFilter() yfilter.YFilter { return throttleInfos.YFilter }

func (throttleInfos *Dhcpv6_Nodes_Node_Server_Profiles_Profile_ThrottleInfos) SetFilter(yf yfilter.YFilter) { throttleInfos.YFilter = yf }

func (throttleInfos *Dhcpv6_Nodes_Node_Server_Profiles_Profile_ThrottleInfos) GetGoName(yname string) string {
    if yname == "throttle-info" { return "ThrottleInfo" }
    return ""
}

func (throttleInfos *Dhcpv6_Nodes_Node_Server_Profiles_Profile_ThrottleInfos) GetSegmentPath() string {
    return "throttle-infos"
}

func (throttleInfos *Dhcpv6_Nodes_Node_Server_Profiles_Profile_ThrottleInfos) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "throttle-info" {
        for _, c := range throttleInfos.ThrottleInfo {
            if throttleInfos.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dhcpv6_Nodes_Node_Server_Profiles_Profile_ThrottleInfos_ThrottleInfo{}
        throttleInfos.ThrottleInfo = append(throttleInfos.ThrottleInfo, child)
        return &throttleInfos.ThrottleInfo[len(throttleInfos.ThrottleInfo)-1]
    }
    return nil
}

func (throttleInfos *Dhcpv6_Nodes_Node_Server_Profiles_Profile_ThrottleInfos) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range throttleInfos.ThrottleInfo {
        children[throttleInfos.ThrottleInfo[i].GetSegmentPath()] = &throttleInfos.ThrottleInfo[i]
    }
    return children
}

func (throttleInfos *Dhcpv6_Nodes_Node_Server_Profiles_Profile_ThrottleInfos) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (throttleInfos *Dhcpv6_Nodes_Node_Server_Profiles_Profile_ThrottleInfos) GetBundleName() string { return "cisco_ios_xr" }

func (throttleInfos *Dhcpv6_Nodes_Node_Server_Profiles_Profile_ThrottleInfos) GetYangName() string { return "throttle-infos" }

func (throttleInfos *Dhcpv6_Nodes_Node_Server_Profiles_Profile_ThrottleInfos) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (throttleInfos *Dhcpv6_Nodes_Node_Server_Profiles_Profile_ThrottleInfos) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (throttleInfos *Dhcpv6_Nodes_Node_Server_Profiles_Profile_ThrottleInfos) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (throttleInfos *Dhcpv6_Nodes_Node_Server_Profiles_Profile_ThrottleInfos) SetParent(parent types.Entity) { throttleInfos.parent = parent }

func (throttleInfos *Dhcpv6_Nodes_Node_Server_Profiles_Profile_ThrottleInfos) GetParent() types.Entity { return throttleInfos.parent }

func (throttleInfos *Dhcpv6_Nodes_Node_Server_Profiles_Profile_ThrottleInfos) GetParentYangName() string { return "profile" }

// Dhcpv6_Nodes_Node_Server_Profiles_Profile_ThrottleInfos_ThrottleInfo
// IPv6 DHCP server profile Throttle Info
type Dhcpv6_Nodes_Node_Server_Profiles_Profile_ThrottleInfos_ThrottleInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. MAC address. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    MacAddress interface{}

    // Client MAC address. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    BindingChaddr interface{}

    // DHCP access interface. The type is string with length: 0..65.
    Ifname interface{}

    // State of entry. The type is interface{} with range: 0..4294967295.
    State interface{}

    // Time Left in secs. The type is interface{} with range: 0..4294967295. Units
    // are second.
    TimeLeft interface{}
}

func (throttleInfo *Dhcpv6_Nodes_Node_Server_Profiles_Profile_ThrottleInfos_ThrottleInfo) GetFilter() yfilter.YFilter { return throttleInfo.YFilter }

func (throttleInfo *Dhcpv6_Nodes_Node_Server_Profiles_Profile_ThrottleInfos_ThrottleInfo) SetFilter(yf yfilter.YFilter) { throttleInfo.YFilter = yf }

func (throttleInfo *Dhcpv6_Nodes_Node_Server_Profiles_Profile_ThrottleInfos_ThrottleInfo) GetGoName(yname string) string {
    if yname == "mac-address" { return "MacAddress" }
    if yname == "binding-chaddr" { return "BindingChaddr" }
    if yname == "ifname" { return "Ifname" }
    if yname == "state" { return "State" }
    if yname == "time-left" { return "TimeLeft" }
    return ""
}

func (throttleInfo *Dhcpv6_Nodes_Node_Server_Profiles_Profile_ThrottleInfos_ThrottleInfo) GetSegmentPath() string {
    return "throttle-info" + "[mac-address='" + fmt.Sprintf("%v", throttleInfo.MacAddress) + "']"
}

func (throttleInfo *Dhcpv6_Nodes_Node_Server_Profiles_Profile_ThrottleInfos_ThrottleInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (throttleInfo *Dhcpv6_Nodes_Node_Server_Profiles_Profile_ThrottleInfos_ThrottleInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (throttleInfo *Dhcpv6_Nodes_Node_Server_Profiles_Profile_ThrottleInfos_ThrottleInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mac-address"] = throttleInfo.MacAddress
    leafs["binding-chaddr"] = throttleInfo.BindingChaddr
    leafs["ifname"] = throttleInfo.Ifname
    leafs["state"] = throttleInfo.State
    leafs["time-left"] = throttleInfo.TimeLeft
    return leafs
}

func (throttleInfo *Dhcpv6_Nodes_Node_Server_Profiles_Profile_ThrottleInfos_ThrottleInfo) GetBundleName() string { return "cisco_ios_xr" }

func (throttleInfo *Dhcpv6_Nodes_Node_Server_Profiles_Profile_ThrottleInfos_ThrottleInfo) GetYangName() string { return "throttle-info" }

func (throttleInfo *Dhcpv6_Nodes_Node_Server_Profiles_Profile_ThrottleInfos_ThrottleInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (throttleInfo *Dhcpv6_Nodes_Node_Server_Profiles_Profile_ThrottleInfos_ThrottleInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (throttleInfo *Dhcpv6_Nodes_Node_Server_Profiles_Profile_ThrottleInfos_ThrottleInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (throttleInfo *Dhcpv6_Nodes_Node_Server_Profiles_Profile_ThrottleInfos_ThrottleInfo) SetParent(parent types.Entity) { throttleInfo.parent = parent }

func (throttleInfo *Dhcpv6_Nodes_Node_Server_Profiles_Profile_ThrottleInfos_ThrottleInfo) GetParent() types.Entity { return throttleInfo.parent }

func (throttleInfo *Dhcpv6_Nodes_Node_Server_Profiles_Profile_ThrottleInfos_ThrottleInfo) GetParentYangName() string { return "throttle-infos" }

// Dhcpv6_Nodes_Node_Server_Interfaces
// DHCPV6 server interface
type Dhcpv6_Nodes_Node_Server_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 DHCP server interface. The type is slice of
    // Dhcpv6_Nodes_Node_Server_Interfaces_Interface.
    Interface []Dhcpv6_Nodes_Node_Server_Interfaces_Interface
}

func (interfaces *Dhcpv6_Nodes_Node_Server_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *Dhcpv6_Nodes_Node_Server_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *Dhcpv6_Nodes_Node_Server_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *Dhcpv6_Nodes_Node_Server_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *Dhcpv6_Nodes_Node_Server_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dhcpv6_Nodes_Node_Server_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *Dhcpv6_Nodes_Node_Server_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *Dhcpv6_Nodes_Node_Server_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *Dhcpv6_Nodes_Node_Server_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *Dhcpv6_Nodes_Node_Server_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *Dhcpv6_Nodes_Node_Server_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *Dhcpv6_Nodes_Node_Server_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *Dhcpv6_Nodes_Node_Server_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *Dhcpv6_Nodes_Node_Server_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *Dhcpv6_Nodes_Node_Server_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *Dhcpv6_Nodes_Node_Server_Interfaces) GetParentYangName() string { return "server" }

// Dhcpv6_Nodes_Node_Server_Interfaces_Interface
// IPv6 DHCP server interface
type Dhcpv6_Nodes_Node_Server_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    InterfaceName interface{}

    // VRF name. The type is string with length: 0..33.
    ServerVrfName interface{}

    // Mode of interface. The type is BagDhcpv6DSubMode.
    ServerInterfaceMode interface{}

    // Is interface ambiguous. The type is interface{} with range: 0..4294967295.
    IsServerInterfaceAmbiguous interface{}

    // Name of profile attached to the interface. The type is string with length:
    // 0..65.
    ServerInterfaceProfileName interface{}

    // Lease limit type on interface. The type is LeaseLimit.
    ServerInterfaceLeaseLimitType interface{}

    // Lease limit count on interface. The type is interface{} with range:
    // 0..4294967295.
    ServerInterfaceLeaseLimits interface{}

    // DHCPv6 Interface SRG role. The type is BagDhcpv6DIntfSrgRole.
    SrgRole interface{}

    // DHCPv6 Interface SERG role. The type is BagDhcpv6DIntfSergRole.
    SergRole interface{}

    // Mac Throttle Status. The type is bool.
    MacThrottle interface{}

    // SRG VRF name. The type is string with length: 0..33.
    SrgVrfName interface{}

    // SRG P2P Status. The type is bool.
    Srgp2P interface{}
}

func (self *Dhcpv6_Nodes_Node_Server_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Dhcpv6_Nodes_Node_Server_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Dhcpv6_Nodes_Node_Server_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "server-vrf-name" { return "ServerVrfName" }
    if yname == "server-interface-mode" { return "ServerInterfaceMode" }
    if yname == "is-server-interface-ambiguous" { return "IsServerInterfaceAmbiguous" }
    if yname == "server-interface-profile-name" { return "ServerInterfaceProfileName" }
    if yname == "server-interface-lease-limit-type" { return "ServerInterfaceLeaseLimitType" }
    if yname == "server-interface-lease-limits" { return "ServerInterfaceLeaseLimits" }
    if yname == "srg-role" { return "SrgRole" }
    if yname == "serg-role" { return "SergRole" }
    if yname == "mac-throttle" { return "MacThrottle" }
    if yname == "srg-vrf-name" { return "SrgVrfName" }
    if yname == "srgp2p" { return "Srgp2P" }
    return ""
}

func (self *Dhcpv6_Nodes_Node_Server_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *Dhcpv6_Nodes_Node_Server_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *Dhcpv6_Nodes_Node_Server_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *Dhcpv6_Nodes_Node_Server_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    leafs["server-vrf-name"] = self.ServerVrfName
    leafs["server-interface-mode"] = self.ServerInterfaceMode
    leafs["is-server-interface-ambiguous"] = self.IsServerInterfaceAmbiguous
    leafs["server-interface-profile-name"] = self.ServerInterfaceProfileName
    leafs["server-interface-lease-limit-type"] = self.ServerInterfaceLeaseLimitType
    leafs["server-interface-lease-limits"] = self.ServerInterfaceLeaseLimits
    leafs["srg-role"] = self.SrgRole
    leafs["serg-role"] = self.SergRole
    leafs["mac-throttle"] = self.MacThrottle
    leafs["srg-vrf-name"] = self.SrgVrfName
    leafs["srgp2p"] = self.Srgp2P
    return leafs
}

func (self *Dhcpv6_Nodes_Node_Server_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *Dhcpv6_Nodes_Node_Server_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *Dhcpv6_Nodes_Node_Server_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Dhcpv6_Nodes_Node_Server_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Dhcpv6_Nodes_Node_Server_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Dhcpv6_Nodes_Node_Server_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *Dhcpv6_Nodes_Node_Server_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *Dhcpv6_Nodes_Node_Server_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// Dhcpv6_Nodes_Node_Server_Statistics
// DHCPv6 server statistics
type Dhcpv6_Nodes_Node_Server_Statistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ipv6 dhcpv6d server stat. The type is slice of
    // Dhcpv6_Nodes_Node_Server_Statistics_Ipv6Dhcpv6DServerStat.
    Ipv6Dhcpv6DServerStat []Dhcpv6_Nodes_Node_Server_Statistics_Ipv6Dhcpv6DServerStat
}

func (statistics *Dhcpv6_Nodes_Node_Server_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *Dhcpv6_Nodes_Node_Server_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *Dhcpv6_Nodes_Node_Server_Statistics) GetGoName(yname string) string {
    if yname == "ipv6-dhcpv6d-server-stat" { return "Ipv6Dhcpv6DServerStat" }
    return ""
}

func (statistics *Dhcpv6_Nodes_Node_Server_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *Dhcpv6_Nodes_Node_Server_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv6-dhcpv6d-server-stat" {
        for _, c := range statistics.Ipv6Dhcpv6DServerStat {
            if statistics.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dhcpv6_Nodes_Node_Server_Statistics_Ipv6Dhcpv6DServerStat{}
        statistics.Ipv6Dhcpv6DServerStat = append(statistics.Ipv6Dhcpv6DServerStat, child)
        return &statistics.Ipv6Dhcpv6DServerStat[len(statistics.Ipv6Dhcpv6DServerStat)-1]
    }
    return nil
}

func (statistics *Dhcpv6_Nodes_Node_Server_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range statistics.Ipv6Dhcpv6DServerStat {
        children[statistics.Ipv6Dhcpv6DServerStat[i].GetSegmentPath()] = &statistics.Ipv6Dhcpv6DServerStat[i]
    }
    return children
}

func (statistics *Dhcpv6_Nodes_Node_Server_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (statistics *Dhcpv6_Nodes_Node_Server_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *Dhcpv6_Nodes_Node_Server_Statistics) GetYangName() string { return "statistics" }

func (statistics *Dhcpv6_Nodes_Node_Server_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *Dhcpv6_Nodes_Node_Server_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *Dhcpv6_Nodes_Node_Server_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *Dhcpv6_Nodes_Node_Server_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *Dhcpv6_Nodes_Node_Server_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *Dhcpv6_Nodes_Node_Server_Statistics) GetParentYangName() string { return "server" }

// Dhcpv6_Nodes_Node_Server_Statistics_Ipv6Dhcpv6DServerStat
// ipv6 dhcpv6d server stat
type Dhcpv6_Nodes_Node_Server_Statistics_Ipv6Dhcpv6DServerStat struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DHCPv6 L3 VRF name. The type is string with length: 0..33.
    VrfName interface{}

    // Server statistics.
    Statistics Dhcpv6_Nodes_Node_Server_Statistics_Ipv6Dhcpv6DServerStat_Statistics
}

func (ipv6Dhcpv6DServerStat *Dhcpv6_Nodes_Node_Server_Statistics_Ipv6Dhcpv6DServerStat) GetFilter() yfilter.YFilter { return ipv6Dhcpv6DServerStat.YFilter }

func (ipv6Dhcpv6DServerStat *Dhcpv6_Nodes_Node_Server_Statistics_Ipv6Dhcpv6DServerStat) SetFilter(yf yfilter.YFilter) { ipv6Dhcpv6DServerStat.YFilter = yf }

func (ipv6Dhcpv6DServerStat *Dhcpv6_Nodes_Node_Server_Statistics_Ipv6Dhcpv6DServerStat) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "statistics" { return "Statistics" }
    return ""
}

func (ipv6Dhcpv6DServerStat *Dhcpv6_Nodes_Node_Server_Statistics_Ipv6Dhcpv6DServerStat) GetSegmentPath() string {
    return "ipv6-dhcpv6d-server-stat"
}

func (ipv6Dhcpv6DServerStat *Dhcpv6_Nodes_Node_Server_Statistics_Ipv6Dhcpv6DServerStat) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "statistics" {
        return &ipv6Dhcpv6DServerStat.Statistics
    }
    return nil
}

func (ipv6Dhcpv6DServerStat *Dhcpv6_Nodes_Node_Server_Statistics_Ipv6Dhcpv6DServerStat) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["statistics"] = &ipv6Dhcpv6DServerStat.Statistics
    return children
}

func (ipv6Dhcpv6DServerStat *Dhcpv6_Nodes_Node_Server_Statistics_Ipv6Dhcpv6DServerStat) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = ipv6Dhcpv6DServerStat.VrfName
    return leafs
}

func (ipv6Dhcpv6DServerStat *Dhcpv6_Nodes_Node_Server_Statistics_Ipv6Dhcpv6DServerStat) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6Dhcpv6DServerStat *Dhcpv6_Nodes_Node_Server_Statistics_Ipv6Dhcpv6DServerStat) GetYangName() string { return "ipv6-dhcpv6d-server-stat" }

func (ipv6Dhcpv6DServerStat *Dhcpv6_Nodes_Node_Server_Statistics_Ipv6Dhcpv6DServerStat) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6Dhcpv6DServerStat *Dhcpv6_Nodes_Node_Server_Statistics_Ipv6Dhcpv6DServerStat) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6Dhcpv6DServerStat *Dhcpv6_Nodes_Node_Server_Statistics_Ipv6Dhcpv6DServerStat) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6Dhcpv6DServerStat *Dhcpv6_Nodes_Node_Server_Statistics_Ipv6Dhcpv6DServerStat) SetParent(parent types.Entity) { ipv6Dhcpv6DServerStat.parent = parent }

func (ipv6Dhcpv6DServerStat *Dhcpv6_Nodes_Node_Server_Statistics_Ipv6Dhcpv6DServerStat) GetParent() types.Entity { return ipv6Dhcpv6DServerStat.parent }

func (ipv6Dhcpv6DServerStat *Dhcpv6_Nodes_Node_Server_Statistics_Ipv6Dhcpv6DServerStat) GetParentYangName() string { return "statistics" }

// Dhcpv6_Nodes_Node_Server_Statistics_Ipv6Dhcpv6DServerStat_Statistics
// Server statistics
type Dhcpv6_Nodes_Node_Server_Statistics_Ipv6Dhcpv6DServerStat_Statistics struct {
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

func (statistics *Dhcpv6_Nodes_Node_Server_Statistics_Ipv6Dhcpv6DServerStat_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *Dhcpv6_Nodes_Node_Server_Statistics_Ipv6Dhcpv6DServerStat_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *Dhcpv6_Nodes_Node_Server_Statistics_Ipv6Dhcpv6DServerStat_Statistics) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (statistics *Dhcpv6_Nodes_Node_Server_Statistics_Ipv6Dhcpv6DServerStat_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *Dhcpv6_Nodes_Node_Server_Statistics_Ipv6Dhcpv6DServerStat_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (statistics *Dhcpv6_Nodes_Node_Server_Statistics_Ipv6Dhcpv6DServerStat_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (statistics *Dhcpv6_Nodes_Node_Server_Statistics_Ipv6Dhcpv6DServerStat_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = statistics.ReceivedPackets
    leafs["transmitted-packets"] = statistics.TransmittedPackets
    leafs["dropped-packets"] = statistics.DroppedPackets
    return leafs
}

func (statistics *Dhcpv6_Nodes_Node_Server_Statistics_Ipv6Dhcpv6DServerStat_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *Dhcpv6_Nodes_Node_Server_Statistics_Ipv6Dhcpv6DServerStat_Statistics) GetYangName() string { return "statistics" }

func (statistics *Dhcpv6_Nodes_Node_Server_Statistics_Ipv6Dhcpv6DServerStat_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *Dhcpv6_Nodes_Node_Server_Statistics_Ipv6Dhcpv6DServerStat_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *Dhcpv6_Nodes_Node_Server_Statistics_Ipv6Dhcpv6DServerStat_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *Dhcpv6_Nodes_Node_Server_Statistics_Ipv6Dhcpv6DServerStat_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *Dhcpv6_Nodes_Node_Server_Statistics_Ipv6Dhcpv6DServerStat_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *Dhcpv6_Nodes_Node_Server_Statistics_Ipv6Dhcpv6DServerStat_Statistics) GetParentYangName() string { return "ipv6-dhcpv6d-server-stat" }

// Dhcpv6_Nodes_Node_Server_BindingOptions
// DHCPv6 server binding with options
type Dhcpv6_Nodes_Node_Server_BindingOptions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DHCPv6 server binding from MAC address.
    MacBindOptions Dhcpv6_Nodes_Node_Server_BindingOptions_MacBindOptions

    // DHCPv6 server binding from DUID.
    DuidBindOptions Dhcpv6_Nodes_Node_Server_BindingOptions_DuidBindOptions
}

func (bindingOptions *Dhcpv6_Nodes_Node_Server_BindingOptions) GetFilter() yfilter.YFilter { return bindingOptions.YFilter }

func (bindingOptions *Dhcpv6_Nodes_Node_Server_BindingOptions) SetFilter(yf yfilter.YFilter) { bindingOptions.YFilter = yf }

func (bindingOptions *Dhcpv6_Nodes_Node_Server_BindingOptions) GetGoName(yname string) string {
    if yname == "mac-bind-options" { return "MacBindOptions" }
    if yname == "duid-bind-options" { return "DuidBindOptions" }
    return ""
}

func (bindingOptions *Dhcpv6_Nodes_Node_Server_BindingOptions) GetSegmentPath() string {
    return "binding-options"
}

func (bindingOptions *Dhcpv6_Nodes_Node_Server_BindingOptions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mac-bind-options" {
        return &bindingOptions.MacBindOptions
    }
    if childYangName == "duid-bind-options" {
        return &bindingOptions.DuidBindOptions
    }
    return nil
}

func (bindingOptions *Dhcpv6_Nodes_Node_Server_BindingOptions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["mac-bind-options"] = &bindingOptions.MacBindOptions
    children["duid-bind-options"] = &bindingOptions.DuidBindOptions
    return children
}

func (bindingOptions *Dhcpv6_Nodes_Node_Server_BindingOptions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bindingOptions *Dhcpv6_Nodes_Node_Server_BindingOptions) GetBundleName() string { return "cisco_ios_xr" }

func (bindingOptions *Dhcpv6_Nodes_Node_Server_BindingOptions) GetYangName() string { return "binding-options" }

func (bindingOptions *Dhcpv6_Nodes_Node_Server_BindingOptions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bindingOptions *Dhcpv6_Nodes_Node_Server_BindingOptions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bindingOptions *Dhcpv6_Nodes_Node_Server_BindingOptions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bindingOptions *Dhcpv6_Nodes_Node_Server_BindingOptions) SetParent(parent types.Entity) { bindingOptions.parent = parent }

func (bindingOptions *Dhcpv6_Nodes_Node_Server_BindingOptions) GetParent() types.Entity { return bindingOptions.parent }

func (bindingOptions *Dhcpv6_Nodes_Node_Server_BindingOptions) GetParentYangName() string { return "server" }

// Dhcpv6_Nodes_Node_Server_BindingOptions_MacBindOptions
// DHCPv6 server binding from MAC address
type Dhcpv6_Nodes_Node_Server_BindingOptions_MacBindOptions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DHCPv6 server binding with options. The type is slice of
    // Dhcpv6_Nodes_Node_Server_BindingOptions_MacBindOptions_MacBindOption.
    MacBindOption []Dhcpv6_Nodes_Node_Server_BindingOptions_MacBindOptions_MacBindOption
}

func (macBindOptions *Dhcpv6_Nodes_Node_Server_BindingOptions_MacBindOptions) GetFilter() yfilter.YFilter { return macBindOptions.YFilter }

func (macBindOptions *Dhcpv6_Nodes_Node_Server_BindingOptions_MacBindOptions) SetFilter(yf yfilter.YFilter) { macBindOptions.YFilter = yf }

func (macBindOptions *Dhcpv6_Nodes_Node_Server_BindingOptions_MacBindOptions) GetGoName(yname string) string {
    if yname == "mac-bind-option" { return "MacBindOption" }
    return ""
}

func (macBindOptions *Dhcpv6_Nodes_Node_Server_BindingOptions_MacBindOptions) GetSegmentPath() string {
    return "mac-bind-options"
}

func (macBindOptions *Dhcpv6_Nodes_Node_Server_BindingOptions_MacBindOptions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mac-bind-option" {
        for _, c := range macBindOptions.MacBindOption {
            if macBindOptions.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dhcpv6_Nodes_Node_Server_BindingOptions_MacBindOptions_MacBindOption{}
        macBindOptions.MacBindOption = append(macBindOptions.MacBindOption, child)
        return &macBindOptions.MacBindOption[len(macBindOptions.MacBindOption)-1]
    }
    return nil
}

func (macBindOptions *Dhcpv6_Nodes_Node_Server_BindingOptions_MacBindOptions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range macBindOptions.MacBindOption {
        children[macBindOptions.MacBindOption[i].GetSegmentPath()] = &macBindOptions.MacBindOption[i]
    }
    return children
}

func (macBindOptions *Dhcpv6_Nodes_Node_Server_BindingOptions_MacBindOptions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (macBindOptions *Dhcpv6_Nodes_Node_Server_BindingOptions_MacBindOptions) GetBundleName() string { return "cisco_ios_xr" }

func (macBindOptions *Dhcpv6_Nodes_Node_Server_BindingOptions_MacBindOptions) GetYangName() string { return "mac-bind-options" }

func (macBindOptions *Dhcpv6_Nodes_Node_Server_BindingOptions_MacBindOptions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (macBindOptions *Dhcpv6_Nodes_Node_Server_BindingOptions_MacBindOptions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (macBindOptions *Dhcpv6_Nodes_Node_Server_BindingOptions_MacBindOptions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (macBindOptions *Dhcpv6_Nodes_Node_Server_BindingOptions_MacBindOptions) SetParent(parent types.Entity) { macBindOptions.parent = parent }

func (macBindOptions *Dhcpv6_Nodes_Node_Server_BindingOptions_MacBindOptions) GetParent() types.Entity { return macBindOptions.parent }

func (macBindOptions *Dhcpv6_Nodes_Node_Server_BindingOptions_MacBindOptions) GetParentYangName() string { return "binding-options" }

// Dhcpv6_Nodes_Node_Server_BindingOptions_MacBindOptions_MacBindOption
// DHCPv6 server binding with options
type Dhcpv6_Nodes_Node_Server_BindingOptions_MacBindOptions_MacBindOption struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. MAC address. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    MacAddress interface{}

    // Client MAC address. The type is string.
    MacAddressXr interface{}

    // Client DUID. The type is string.
    DuidXr interface{}

    // DNS address count. The type is interface{} with range: 0..255.
    DnsCount interface{}

    // Client Option 17 value. The type is string.
    Opt17 interface{}

    // DNS addresses. The type is slice of string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    DnsAddress []interface{}
}

func (macBindOption *Dhcpv6_Nodes_Node_Server_BindingOptions_MacBindOptions_MacBindOption) GetFilter() yfilter.YFilter { return macBindOption.YFilter }

func (macBindOption *Dhcpv6_Nodes_Node_Server_BindingOptions_MacBindOptions_MacBindOption) SetFilter(yf yfilter.YFilter) { macBindOption.YFilter = yf }

func (macBindOption *Dhcpv6_Nodes_Node_Server_BindingOptions_MacBindOptions_MacBindOption) GetGoName(yname string) string {
    if yname == "mac-address" { return "MacAddress" }
    if yname == "mac-address-xr" { return "MacAddressXr" }
    if yname == "duid-xr" { return "DuidXr" }
    if yname == "dns-count" { return "DnsCount" }
    if yname == "opt17" { return "Opt17" }
    if yname == "dns-address" { return "DnsAddress" }
    return ""
}

func (macBindOption *Dhcpv6_Nodes_Node_Server_BindingOptions_MacBindOptions_MacBindOption) GetSegmentPath() string {
    return "mac-bind-option" + "[mac-address='" + fmt.Sprintf("%v", macBindOption.MacAddress) + "']"
}

func (macBindOption *Dhcpv6_Nodes_Node_Server_BindingOptions_MacBindOptions_MacBindOption) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (macBindOption *Dhcpv6_Nodes_Node_Server_BindingOptions_MacBindOptions_MacBindOption) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (macBindOption *Dhcpv6_Nodes_Node_Server_BindingOptions_MacBindOptions_MacBindOption) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mac-address"] = macBindOption.MacAddress
    leafs["mac-address-xr"] = macBindOption.MacAddressXr
    leafs["duid-xr"] = macBindOption.DuidXr
    leafs["dns-count"] = macBindOption.DnsCount
    leafs["opt17"] = macBindOption.Opt17
    leafs["dns-address"] = macBindOption.DnsAddress
    return leafs
}

func (macBindOption *Dhcpv6_Nodes_Node_Server_BindingOptions_MacBindOptions_MacBindOption) GetBundleName() string { return "cisco_ios_xr" }

func (macBindOption *Dhcpv6_Nodes_Node_Server_BindingOptions_MacBindOptions_MacBindOption) GetYangName() string { return "mac-bind-option" }

func (macBindOption *Dhcpv6_Nodes_Node_Server_BindingOptions_MacBindOptions_MacBindOption) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (macBindOption *Dhcpv6_Nodes_Node_Server_BindingOptions_MacBindOptions_MacBindOption) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (macBindOption *Dhcpv6_Nodes_Node_Server_BindingOptions_MacBindOptions_MacBindOption) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (macBindOption *Dhcpv6_Nodes_Node_Server_BindingOptions_MacBindOptions_MacBindOption) SetParent(parent types.Entity) { macBindOption.parent = parent }

func (macBindOption *Dhcpv6_Nodes_Node_Server_BindingOptions_MacBindOptions_MacBindOption) GetParent() types.Entity { return macBindOption.parent }

func (macBindOption *Dhcpv6_Nodes_Node_Server_BindingOptions_MacBindOptions_MacBindOption) GetParentYangName() string { return "mac-bind-options" }

// Dhcpv6_Nodes_Node_Server_BindingOptions_DuidBindOptions
// DHCPv6 server binding from DUID
type Dhcpv6_Nodes_Node_Server_BindingOptions_DuidBindOptions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DHCPv6 server binding with options. The type is slice of
    // Dhcpv6_Nodes_Node_Server_BindingOptions_DuidBindOptions_DuidBindOption.
    DuidBindOption []Dhcpv6_Nodes_Node_Server_BindingOptions_DuidBindOptions_DuidBindOption
}

func (duidBindOptions *Dhcpv6_Nodes_Node_Server_BindingOptions_DuidBindOptions) GetFilter() yfilter.YFilter { return duidBindOptions.YFilter }

func (duidBindOptions *Dhcpv6_Nodes_Node_Server_BindingOptions_DuidBindOptions) SetFilter(yf yfilter.YFilter) { duidBindOptions.YFilter = yf }

func (duidBindOptions *Dhcpv6_Nodes_Node_Server_BindingOptions_DuidBindOptions) GetGoName(yname string) string {
    if yname == "duid-bind-option" { return "DuidBindOption" }
    return ""
}

func (duidBindOptions *Dhcpv6_Nodes_Node_Server_BindingOptions_DuidBindOptions) GetSegmentPath() string {
    return "duid-bind-options"
}

func (duidBindOptions *Dhcpv6_Nodes_Node_Server_BindingOptions_DuidBindOptions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "duid-bind-option" {
        for _, c := range duidBindOptions.DuidBindOption {
            if duidBindOptions.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dhcpv6_Nodes_Node_Server_BindingOptions_DuidBindOptions_DuidBindOption{}
        duidBindOptions.DuidBindOption = append(duidBindOptions.DuidBindOption, child)
        return &duidBindOptions.DuidBindOption[len(duidBindOptions.DuidBindOption)-1]
    }
    return nil
}

func (duidBindOptions *Dhcpv6_Nodes_Node_Server_BindingOptions_DuidBindOptions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range duidBindOptions.DuidBindOption {
        children[duidBindOptions.DuidBindOption[i].GetSegmentPath()] = &duidBindOptions.DuidBindOption[i]
    }
    return children
}

func (duidBindOptions *Dhcpv6_Nodes_Node_Server_BindingOptions_DuidBindOptions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (duidBindOptions *Dhcpv6_Nodes_Node_Server_BindingOptions_DuidBindOptions) GetBundleName() string { return "cisco_ios_xr" }

func (duidBindOptions *Dhcpv6_Nodes_Node_Server_BindingOptions_DuidBindOptions) GetYangName() string { return "duid-bind-options" }

func (duidBindOptions *Dhcpv6_Nodes_Node_Server_BindingOptions_DuidBindOptions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (duidBindOptions *Dhcpv6_Nodes_Node_Server_BindingOptions_DuidBindOptions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (duidBindOptions *Dhcpv6_Nodes_Node_Server_BindingOptions_DuidBindOptions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (duidBindOptions *Dhcpv6_Nodes_Node_Server_BindingOptions_DuidBindOptions) SetParent(parent types.Entity) { duidBindOptions.parent = parent }

func (duidBindOptions *Dhcpv6_Nodes_Node_Server_BindingOptions_DuidBindOptions) GetParent() types.Entity { return duidBindOptions.parent }

func (duidBindOptions *Dhcpv6_Nodes_Node_Server_BindingOptions_DuidBindOptions) GetParentYangName() string { return "binding-options" }

// Dhcpv6_Nodes_Node_Server_BindingOptions_DuidBindOptions_DuidBindOption
// DHCPv6 server binding with options
type Dhcpv6_Nodes_Node_Server_BindingOptions_DuidBindOptions_DuidBindOption struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. DUID of Binding. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Duid interface{}

    // Client MAC address. The type is string.
    MacAddressXr interface{}

    // Client DUID. The type is string.
    DuidXr interface{}

    // DNS address count. The type is interface{} with range: 0..255.
    DnsCount interface{}

    // Client Option 17 value. The type is string.
    Opt17 interface{}

    // DNS addresses. The type is slice of string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    DnsAddress []interface{}
}

func (duidBindOption *Dhcpv6_Nodes_Node_Server_BindingOptions_DuidBindOptions_DuidBindOption) GetFilter() yfilter.YFilter { return duidBindOption.YFilter }

func (duidBindOption *Dhcpv6_Nodes_Node_Server_BindingOptions_DuidBindOptions_DuidBindOption) SetFilter(yf yfilter.YFilter) { duidBindOption.YFilter = yf }

func (duidBindOption *Dhcpv6_Nodes_Node_Server_BindingOptions_DuidBindOptions_DuidBindOption) GetGoName(yname string) string {
    if yname == "duid" { return "Duid" }
    if yname == "mac-address-xr" { return "MacAddressXr" }
    if yname == "duid-xr" { return "DuidXr" }
    if yname == "dns-count" { return "DnsCount" }
    if yname == "opt17" { return "Opt17" }
    if yname == "dns-address" { return "DnsAddress" }
    return ""
}

func (duidBindOption *Dhcpv6_Nodes_Node_Server_BindingOptions_DuidBindOptions_DuidBindOption) GetSegmentPath() string {
    return "duid-bind-option" + "[duid='" + fmt.Sprintf("%v", duidBindOption.Duid) + "']"
}

func (duidBindOption *Dhcpv6_Nodes_Node_Server_BindingOptions_DuidBindOptions_DuidBindOption) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (duidBindOption *Dhcpv6_Nodes_Node_Server_BindingOptions_DuidBindOptions_DuidBindOption) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (duidBindOption *Dhcpv6_Nodes_Node_Server_BindingOptions_DuidBindOptions_DuidBindOption) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["duid"] = duidBindOption.Duid
    leafs["mac-address-xr"] = duidBindOption.MacAddressXr
    leafs["duid-xr"] = duidBindOption.DuidXr
    leafs["dns-count"] = duidBindOption.DnsCount
    leafs["opt17"] = duidBindOption.Opt17
    leafs["dns-address"] = duidBindOption.DnsAddress
    return leafs
}

func (duidBindOption *Dhcpv6_Nodes_Node_Server_BindingOptions_DuidBindOptions_DuidBindOption) GetBundleName() string { return "cisco_ios_xr" }

func (duidBindOption *Dhcpv6_Nodes_Node_Server_BindingOptions_DuidBindOptions_DuidBindOption) GetYangName() string { return "duid-bind-option" }

func (duidBindOption *Dhcpv6_Nodes_Node_Server_BindingOptions_DuidBindOptions_DuidBindOption) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (duidBindOption *Dhcpv6_Nodes_Node_Server_BindingOptions_DuidBindOptions_DuidBindOption) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (duidBindOption *Dhcpv6_Nodes_Node_Server_BindingOptions_DuidBindOptions_DuidBindOption) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (duidBindOption *Dhcpv6_Nodes_Node_Server_BindingOptions_DuidBindOptions_DuidBindOption) SetParent(parent types.Entity) { duidBindOption.parent = parent }

func (duidBindOption *Dhcpv6_Nodes_Node_Server_BindingOptions_DuidBindOptions_DuidBindOption) GetParent() types.Entity { return duidBindOption.parent }

func (duidBindOption *Dhcpv6_Nodes_Node_Server_BindingOptions_DuidBindOptions_DuidBindOption) GetParentYangName() string { return "duid-bind-options" }

// Dhcpv6_Nodes_Node_Relay
// IPv6 DHCP relay operational data
type Dhcpv6_Nodes_Node_Relay struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DHCPv6 relay statistics.
    Statistics Dhcpv6_Nodes_Node_Relay_Statistics

    // DHCPV6 relay bindings.
    Binding Dhcpv6_Nodes_Node_Relay_Binding

    // DHCPV6 relay list of VRF names.
    Vrfs Dhcpv6_Nodes_Node_Relay_Vrfs
}

func (relay *Dhcpv6_Nodes_Node_Relay) GetFilter() yfilter.YFilter { return relay.YFilter }

func (relay *Dhcpv6_Nodes_Node_Relay) SetFilter(yf yfilter.YFilter) { relay.YFilter = yf }

func (relay *Dhcpv6_Nodes_Node_Relay) GetGoName(yname string) string {
    if yname == "statistics" { return "Statistics" }
    if yname == "binding" { return "Binding" }
    if yname == "vrfs" { return "Vrfs" }
    return ""
}

func (relay *Dhcpv6_Nodes_Node_Relay) GetSegmentPath() string {
    return "relay"
}

func (relay *Dhcpv6_Nodes_Node_Relay) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "statistics" {
        return &relay.Statistics
    }
    if childYangName == "binding" {
        return &relay.Binding
    }
    if childYangName == "vrfs" {
        return &relay.Vrfs
    }
    return nil
}

func (relay *Dhcpv6_Nodes_Node_Relay) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["statistics"] = &relay.Statistics
    children["binding"] = &relay.Binding
    children["vrfs"] = &relay.Vrfs
    return children
}

func (relay *Dhcpv6_Nodes_Node_Relay) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (relay *Dhcpv6_Nodes_Node_Relay) GetBundleName() string { return "cisco_ios_xr" }

func (relay *Dhcpv6_Nodes_Node_Relay) GetYangName() string { return "relay" }

func (relay *Dhcpv6_Nodes_Node_Relay) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (relay *Dhcpv6_Nodes_Node_Relay) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (relay *Dhcpv6_Nodes_Node_Relay) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (relay *Dhcpv6_Nodes_Node_Relay) SetParent(parent types.Entity) { relay.parent = parent }

func (relay *Dhcpv6_Nodes_Node_Relay) GetParent() types.Entity { return relay.parent }

func (relay *Dhcpv6_Nodes_Node_Relay) GetParentYangName() string { return "node" }

// Dhcpv6_Nodes_Node_Relay_Statistics
// DHCPv6 relay statistics
type Dhcpv6_Nodes_Node_Relay_Statistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ipv6 dhcpv6d relay stat. The type is slice of
    // Dhcpv6_Nodes_Node_Relay_Statistics_Ipv6Dhcpv6DRelayStat.
    Ipv6Dhcpv6DRelayStat []Dhcpv6_Nodes_Node_Relay_Statistics_Ipv6Dhcpv6DRelayStat
}

func (statistics *Dhcpv6_Nodes_Node_Relay_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *Dhcpv6_Nodes_Node_Relay_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *Dhcpv6_Nodes_Node_Relay_Statistics) GetGoName(yname string) string {
    if yname == "ipv6-dhcpv6d-relay-stat" { return "Ipv6Dhcpv6DRelayStat" }
    return ""
}

func (statistics *Dhcpv6_Nodes_Node_Relay_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *Dhcpv6_Nodes_Node_Relay_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv6-dhcpv6d-relay-stat" {
        for _, c := range statistics.Ipv6Dhcpv6DRelayStat {
            if statistics.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dhcpv6_Nodes_Node_Relay_Statistics_Ipv6Dhcpv6DRelayStat{}
        statistics.Ipv6Dhcpv6DRelayStat = append(statistics.Ipv6Dhcpv6DRelayStat, child)
        return &statistics.Ipv6Dhcpv6DRelayStat[len(statistics.Ipv6Dhcpv6DRelayStat)-1]
    }
    return nil
}

func (statistics *Dhcpv6_Nodes_Node_Relay_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range statistics.Ipv6Dhcpv6DRelayStat {
        children[statistics.Ipv6Dhcpv6DRelayStat[i].GetSegmentPath()] = &statistics.Ipv6Dhcpv6DRelayStat[i]
    }
    return children
}

func (statistics *Dhcpv6_Nodes_Node_Relay_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (statistics *Dhcpv6_Nodes_Node_Relay_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *Dhcpv6_Nodes_Node_Relay_Statistics) GetYangName() string { return "statistics" }

func (statistics *Dhcpv6_Nodes_Node_Relay_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *Dhcpv6_Nodes_Node_Relay_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *Dhcpv6_Nodes_Node_Relay_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *Dhcpv6_Nodes_Node_Relay_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *Dhcpv6_Nodes_Node_Relay_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *Dhcpv6_Nodes_Node_Relay_Statistics) GetParentYangName() string { return "relay" }

// Dhcpv6_Nodes_Node_Relay_Statistics_Ipv6Dhcpv6DRelayStat
// ipv6 dhcpv6d relay stat
type Dhcpv6_Nodes_Node_Relay_Statistics_Ipv6Dhcpv6DRelayStat struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DHCPv6 L3 VRF name. The type is string with length: 0..33.
    VrfName interface{}

    // Relay statistics.
    Statistics Dhcpv6_Nodes_Node_Relay_Statistics_Ipv6Dhcpv6DRelayStat_Statistics
}

func (ipv6Dhcpv6DRelayStat *Dhcpv6_Nodes_Node_Relay_Statistics_Ipv6Dhcpv6DRelayStat) GetFilter() yfilter.YFilter { return ipv6Dhcpv6DRelayStat.YFilter }

func (ipv6Dhcpv6DRelayStat *Dhcpv6_Nodes_Node_Relay_Statistics_Ipv6Dhcpv6DRelayStat) SetFilter(yf yfilter.YFilter) { ipv6Dhcpv6DRelayStat.YFilter = yf }

func (ipv6Dhcpv6DRelayStat *Dhcpv6_Nodes_Node_Relay_Statistics_Ipv6Dhcpv6DRelayStat) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "statistics" { return "Statistics" }
    return ""
}

func (ipv6Dhcpv6DRelayStat *Dhcpv6_Nodes_Node_Relay_Statistics_Ipv6Dhcpv6DRelayStat) GetSegmentPath() string {
    return "ipv6-dhcpv6d-relay-stat"
}

func (ipv6Dhcpv6DRelayStat *Dhcpv6_Nodes_Node_Relay_Statistics_Ipv6Dhcpv6DRelayStat) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "statistics" {
        return &ipv6Dhcpv6DRelayStat.Statistics
    }
    return nil
}

func (ipv6Dhcpv6DRelayStat *Dhcpv6_Nodes_Node_Relay_Statistics_Ipv6Dhcpv6DRelayStat) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["statistics"] = &ipv6Dhcpv6DRelayStat.Statistics
    return children
}

func (ipv6Dhcpv6DRelayStat *Dhcpv6_Nodes_Node_Relay_Statistics_Ipv6Dhcpv6DRelayStat) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = ipv6Dhcpv6DRelayStat.VrfName
    return leafs
}

func (ipv6Dhcpv6DRelayStat *Dhcpv6_Nodes_Node_Relay_Statistics_Ipv6Dhcpv6DRelayStat) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6Dhcpv6DRelayStat *Dhcpv6_Nodes_Node_Relay_Statistics_Ipv6Dhcpv6DRelayStat) GetYangName() string { return "ipv6-dhcpv6d-relay-stat" }

func (ipv6Dhcpv6DRelayStat *Dhcpv6_Nodes_Node_Relay_Statistics_Ipv6Dhcpv6DRelayStat) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6Dhcpv6DRelayStat *Dhcpv6_Nodes_Node_Relay_Statistics_Ipv6Dhcpv6DRelayStat) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6Dhcpv6DRelayStat *Dhcpv6_Nodes_Node_Relay_Statistics_Ipv6Dhcpv6DRelayStat) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6Dhcpv6DRelayStat *Dhcpv6_Nodes_Node_Relay_Statistics_Ipv6Dhcpv6DRelayStat) SetParent(parent types.Entity) { ipv6Dhcpv6DRelayStat.parent = parent }

func (ipv6Dhcpv6DRelayStat *Dhcpv6_Nodes_Node_Relay_Statistics_Ipv6Dhcpv6DRelayStat) GetParent() types.Entity { return ipv6Dhcpv6DRelayStat.parent }

func (ipv6Dhcpv6DRelayStat *Dhcpv6_Nodes_Node_Relay_Statistics_Ipv6Dhcpv6DRelayStat) GetParentYangName() string { return "statistics" }

// Dhcpv6_Nodes_Node_Relay_Statistics_Ipv6Dhcpv6DRelayStat_Statistics
// Relay statistics
type Dhcpv6_Nodes_Node_Relay_Statistics_Ipv6Dhcpv6DRelayStat_Statistics struct {
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

func (statistics *Dhcpv6_Nodes_Node_Relay_Statistics_Ipv6Dhcpv6DRelayStat_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *Dhcpv6_Nodes_Node_Relay_Statistics_Ipv6Dhcpv6DRelayStat_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *Dhcpv6_Nodes_Node_Relay_Statistics_Ipv6Dhcpv6DRelayStat_Statistics) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (statistics *Dhcpv6_Nodes_Node_Relay_Statistics_Ipv6Dhcpv6DRelayStat_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *Dhcpv6_Nodes_Node_Relay_Statistics_Ipv6Dhcpv6DRelayStat_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (statistics *Dhcpv6_Nodes_Node_Relay_Statistics_Ipv6Dhcpv6DRelayStat_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (statistics *Dhcpv6_Nodes_Node_Relay_Statistics_Ipv6Dhcpv6DRelayStat_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = statistics.ReceivedPackets
    leafs["transmitted-packets"] = statistics.TransmittedPackets
    leafs["dropped-packets"] = statistics.DroppedPackets
    return leafs
}

func (statistics *Dhcpv6_Nodes_Node_Relay_Statistics_Ipv6Dhcpv6DRelayStat_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *Dhcpv6_Nodes_Node_Relay_Statistics_Ipv6Dhcpv6DRelayStat_Statistics) GetYangName() string { return "statistics" }

func (statistics *Dhcpv6_Nodes_Node_Relay_Statistics_Ipv6Dhcpv6DRelayStat_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *Dhcpv6_Nodes_Node_Relay_Statistics_Ipv6Dhcpv6DRelayStat_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *Dhcpv6_Nodes_Node_Relay_Statistics_Ipv6Dhcpv6DRelayStat_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *Dhcpv6_Nodes_Node_Relay_Statistics_Ipv6Dhcpv6DRelayStat_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *Dhcpv6_Nodes_Node_Relay_Statistics_Ipv6Dhcpv6DRelayStat_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *Dhcpv6_Nodes_Node_Relay_Statistics_Ipv6Dhcpv6DRelayStat_Statistics) GetParentYangName() string { return "ipv6-dhcpv6d-relay-stat" }

// Dhcpv6_Nodes_Node_Relay_Binding
// DHCPV6 relay bindings
type Dhcpv6_Nodes_Node_Relay_Binding struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DHCPV6 relay binding summary.
    Summary Dhcpv6_Nodes_Node_Relay_Binding_Summary

    // DHCPV6 relay client bindings.
    Clients Dhcpv6_Nodes_Node_Relay_Binding_Clients
}

func (binding *Dhcpv6_Nodes_Node_Relay_Binding) GetFilter() yfilter.YFilter { return binding.YFilter }

func (binding *Dhcpv6_Nodes_Node_Relay_Binding) SetFilter(yf yfilter.YFilter) { binding.YFilter = yf }

func (binding *Dhcpv6_Nodes_Node_Relay_Binding) GetGoName(yname string) string {
    if yname == "summary" { return "Summary" }
    if yname == "clients" { return "Clients" }
    return ""
}

func (binding *Dhcpv6_Nodes_Node_Relay_Binding) GetSegmentPath() string {
    return "binding"
}

func (binding *Dhcpv6_Nodes_Node_Relay_Binding) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "summary" {
        return &binding.Summary
    }
    if childYangName == "clients" {
        return &binding.Clients
    }
    return nil
}

func (binding *Dhcpv6_Nodes_Node_Relay_Binding) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["summary"] = &binding.Summary
    children["clients"] = &binding.Clients
    return children
}

func (binding *Dhcpv6_Nodes_Node_Relay_Binding) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (binding *Dhcpv6_Nodes_Node_Relay_Binding) GetBundleName() string { return "cisco_ios_xr" }

func (binding *Dhcpv6_Nodes_Node_Relay_Binding) GetYangName() string { return "binding" }

func (binding *Dhcpv6_Nodes_Node_Relay_Binding) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (binding *Dhcpv6_Nodes_Node_Relay_Binding) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (binding *Dhcpv6_Nodes_Node_Relay_Binding) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (binding *Dhcpv6_Nodes_Node_Relay_Binding) SetParent(parent types.Entity) { binding.parent = parent }

func (binding *Dhcpv6_Nodes_Node_Relay_Binding) GetParent() types.Entity { return binding.parent }

func (binding *Dhcpv6_Nodes_Node_Relay_Binding) GetParentYangName() string { return "relay" }

// Dhcpv6_Nodes_Node_Relay_Binding_Summary
// DHCPV6 relay binding summary
type Dhcpv6_Nodes_Node_Relay_Binding_Summary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Total number of clients. The type is interface{} with range: 0..4294967295.
    Clients interface{}
}

func (summary *Dhcpv6_Nodes_Node_Relay_Binding_Summary) GetFilter() yfilter.YFilter { return summary.YFilter }

func (summary *Dhcpv6_Nodes_Node_Relay_Binding_Summary) SetFilter(yf yfilter.YFilter) { summary.YFilter = yf }

func (summary *Dhcpv6_Nodes_Node_Relay_Binding_Summary) GetGoName(yname string) string {
    if yname == "clients" { return "Clients" }
    return ""
}

func (summary *Dhcpv6_Nodes_Node_Relay_Binding_Summary) GetSegmentPath() string {
    return "summary"
}

func (summary *Dhcpv6_Nodes_Node_Relay_Binding_Summary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (summary *Dhcpv6_Nodes_Node_Relay_Binding_Summary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (summary *Dhcpv6_Nodes_Node_Relay_Binding_Summary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["clients"] = summary.Clients
    return leafs
}

func (summary *Dhcpv6_Nodes_Node_Relay_Binding_Summary) GetBundleName() string { return "cisco_ios_xr" }

func (summary *Dhcpv6_Nodes_Node_Relay_Binding_Summary) GetYangName() string { return "summary" }

func (summary *Dhcpv6_Nodes_Node_Relay_Binding_Summary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (summary *Dhcpv6_Nodes_Node_Relay_Binding_Summary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (summary *Dhcpv6_Nodes_Node_Relay_Binding_Summary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (summary *Dhcpv6_Nodes_Node_Relay_Binding_Summary) SetParent(parent types.Entity) { summary.parent = parent }

func (summary *Dhcpv6_Nodes_Node_Relay_Binding_Summary) GetParent() types.Entity { return summary.parent }

func (summary *Dhcpv6_Nodes_Node_Relay_Binding_Summary) GetParentYangName() string { return "binding" }

// Dhcpv6_Nodes_Node_Relay_Binding_Clients
// DHCPV6 relay client bindings
type Dhcpv6_Nodes_Node_Relay_Binding_Clients struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Single DHCPV6 relay binding. The type is slice of
    // Dhcpv6_Nodes_Node_Relay_Binding_Clients_Client.
    Client []Dhcpv6_Nodes_Node_Relay_Binding_Clients_Client
}

func (clients *Dhcpv6_Nodes_Node_Relay_Binding_Clients) GetFilter() yfilter.YFilter { return clients.YFilter }

func (clients *Dhcpv6_Nodes_Node_Relay_Binding_Clients) SetFilter(yf yfilter.YFilter) { clients.YFilter = yf }

func (clients *Dhcpv6_Nodes_Node_Relay_Binding_Clients) GetGoName(yname string) string {
    if yname == "client" { return "Client" }
    return ""
}

func (clients *Dhcpv6_Nodes_Node_Relay_Binding_Clients) GetSegmentPath() string {
    return "clients"
}

func (clients *Dhcpv6_Nodes_Node_Relay_Binding_Clients) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "client" {
        for _, c := range clients.Client {
            if clients.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dhcpv6_Nodes_Node_Relay_Binding_Clients_Client{}
        clients.Client = append(clients.Client, child)
        return &clients.Client[len(clients.Client)-1]
    }
    return nil
}

func (clients *Dhcpv6_Nodes_Node_Relay_Binding_Clients) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range clients.Client {
        children[clients.Client[i].GetSegmentPath()] = &clients.Client[i]
    }
    return children
}

func (clients *Dhcpv6_Nodes_Node_Relay_Binding_Clients) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (clients *Dhcpv6_Nodes_Node_Relay_Binding_Clients) GetBundleName() string { return "cisco_ios_xr" }

func (clients *Dhcpv6_Nodes_Node_Relay_Binding_Clients) GetYangName() string { return "clients" }

func (clients *Dhcpv6_Nodes_Node_Relay_Binding_Clients) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clients *Dhcpv6_Nodes_Node_Relay_Binding_Clients) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clients *Dhcpv6_Nodes_Node_Relay_Binding_Clients) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clients *Dhcpv6_Nodes_Node_Relay_Binding_Clients) SetParent(parent types.Entity) { clients.parent = parent }

func (clients *Dhcpv6_Nodes_Node_Relay_Binding_Clients) GetParent() types.Entity { return clients.parent }

func (clients *Dhcpv6_Nodes_Node_Relay_Binding_Clients) GetParentYangName() string { return "binding" }

// Dhcpv6_Nodes_Node_Relay_Binding_Clients_Client
// Single DHCPV6 relay binding
type Dhcpv6_Nodes_Node_Relay_Binding_Clients_Client struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Client ID. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    ClientId interface{}

    // Client DUID. The type is string.
    Duid interface{}

    // Client unique identifier. The type is interface{} with range:
    // 0..4294967295.
    ClientIdXr interface{}

    // length of prefix. The type is interface{} with range: 0..255.
    PrefixLength interface{}

    // DHCPV6 IPv6 Prefix allotted to client. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Prefix interface{}

    // DHCPv6 client/subscriber Vrf name. The type is string with length: 0..33.
    VrfName interface{}

    // Client route lifetime. The type is interface{} with range: 0..4294967295.
    Lifetime interface{}

    // Client route remaining lifetime. The type is interface{} with range:
    // 0..4294967295.
    RemLifeTime interface{}

    // Interface name. The type is string with length: 0..65.
    InterfaceName interface{}

    // Next hop is our address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHopAddr interface{}

    // IA_ID of this IA. The type is interface{} with range: 0..4294967295.
    IaId interface{}

    // Relay Profile name. The type is string with length: 0..65.
    RelayProfileName interface{}
}

func (client *Dhcpv6_Nodes_Node_Relay_Binding_Clients_Client) GetFilter() yfilter.YFilter { return client.YFilter }

func (client *Dhcpv6_Nodes_Node_Relay_Binding_Clients_Client) SetFilter(yf yfilter.YFilter) { client.YFilter = yf }

func (client *Dhcpv6_Nodes_Node_Relay_Binding_Clients_Client) GetGoName(yname string) string {
    if yname == "client-id" { return "ClientId" }
    if yname == "duid" { return "Duid" }
    if yname == "client-id-xr" { return "ClientIdXr" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "prefix" { return "Prefix" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "lifetime" { return "Lifetime" }
    if yname == "rem-life-time" { return "RemLifeTime" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "next-hop-addr" { return "NextHopAddr" }
    if yname == "ia-id" { return "IaId" }
    if yname == "relay-profile-name" { return "RelayProfileName" }
    return ""
}

func (client *Dhcpv6_Nodes_Node_Relay_Binding_Clients_Client) GetSegmentPath() string {
    return "client" + "[client-id='" + fmt.Sprintf("%v", client.ClientId) + "']"
}

func (client *Dhcpv6_Nodes_Node_Relay_Binding_Clients_Client) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (client *Dhcpv6_Nodes_Node_Relay_Binding_Clients_Client) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (client *Dhcpv6_Nodes_Node_Relay_Binding_Clients_Client) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["client-id"] = client.ClientId
    leafs["duid"] = client.Duid
    leafs["client-id-xr"] = client.ClientIdXr
    leafs["prefix-length"] = client.PrefixLength
    leafs["prefix"] = client.Prefix
    leafs["vrf-name"] = client.VrfName
    leafs["lifetime"] = client.Lifetime
    leafs["rem-life-time"] = client.RemLifeTime
    leafs["interface-name"] = client.InterfaceName
    leafs["next-hop-addr"] = client.NextHopAddr
    leafs["ia-id"] = client.IaId
    leafs["relay-profile-name"] = client.RelayProfileName
    return leafs
}

func (client *Dhcpv6_Nodes_Node_Relay_Binding_Clients_Client) GetBundleName() string { return "cisco_ios_xr" }

func (client *Dhcpv6_Nodes_Node_Relay_Binding_Clients_Client) GetYangName() string { return "client" }

func (client *Dhcpv6_Nodes_Node_Relay_Binding_Clients_Client) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (client *Dhcpv6_Nodes_Node_Relay_Binding_Clients_Client) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (client *Dhcpv6_Nodes_Node_Relay_Binding_Clients_Client) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (client *Dhcpv6_Nodes_Node_Relay_Binding_Clients_Client) SetParent(parent types.Entity) { client.parent = parent }

func (client *Dhcpv6_Nodes_Node_Relay_Binding_Clients_Client) GetParent() types.Entity { return client.parent }

func (client *Dhcpv6_Nodes_Node_Relay_Binding_Clients_Client) GetParentYangName() string { return "clients" }

// Dhcpv6_Nodes_Node_Relay_Vrfs
// DHCPV6 relay list of VRF names
type Dhcpv6_Nodes_Node_Relay_Vrfs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 DHCP relay VRF name. The type is slice of
    // Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf.
    Vrf []Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf
}

func (vrfs *Dhcpv6_Nodes_Node_Relay_Vrfs) GetFilter() yfilter.YFilter { return vrfs.YFilter }

func (vrfs *Dhcpv6_Nodes_Node_Relay_Vrfs) SetFilter(yf yfilter.YFilter) { vrfs.YFilter = yf }

func (vrfs *Dhcpv6_Nodes_Node_Relay_Vrfs) GetGoName(yname string) string {
    if yname == "vrf" { return "Vrf" }
    return ""
}

func (vrfs *Dhcpv6_Nodes_Node_Relay_Vrfs) GetSegmentPath() string {
    return "vrfs"
}

func (vrfs *Dhcpv6_Nodes_Node_Relay_Vrfs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf" {
        for _, c := range vrfs.Vrf {
            if vrfs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf{}
        vrfs.Vrf = append(vrfs.Vrf, child)
        return &vrfs.Vrf[len(vrfs.Vrf)-1]
    }
    return nil
}

func (vrfs *Dhcpv6_Nodes_Node_Relay_Vrfs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vrfs.Vrf {
        children[vrfs.Vrf[i].GetSegmentPath()] = &vrfs.Vrf[i]
    }
    return children
}

func (vrfs *Dhcpv6_Nodes_Node_Relay_Vrfs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vrfs *Dhcpv6_Nodes_Node_Relay_Vrfs) GetBundleName() string { return "cisco_ios_xr" }

func (vrfs *Dhcpv6_Nodes_Node_Relay_Vrfs) GetYangName() string { return "vrfs" }

func (vrfs *Dhcpv6_Nodes_Node_Relay_Vrfs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfs *Dhcpv6_Nodes_Node_Relay_Vrfs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfs *Dhcpv6_Nodes_Node_Relay_Vrfs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfs *Dhcpv6_Nodes_Node_Relay_Vrfs) SetParent(parent types.Entity) { vrfs.parent = parent }

func (vrfs *Dhcpv6_Nodes_Node_Relay_Vrfs) GetParent() types.Entity { return vrfs.parent }

func (vrfs *Dhcpv6_Nodes_Node_Relay_Vrfs) GetParentYangName() string { return "relay" }

// Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf
// IPv6 DHCP relay VRF name
type Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // IPv6 DHCP relay statistics.
    Statistics Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics
}

func (vrf *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf) GetFilter() yfilter.YFilter { return vrf.YFilter }

func (vrf *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf) SetFilter(yf yfilter.YFilter) { vrf.YFilter = yf }

func (vrf *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "statistics" { return "Statistics" }
    return ""
}

func (vrf *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf) GetSegmentPath() string {
    return "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
}

func (vrf *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "statistics" {
        return &vrf.Statistics
    }
    return nil
}

func (vrf *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["statistics"] = &vrf.Statistics
    return children
}

func (vrf *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = vrf.VrfName
    return leafs
}

func (vrf *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf) GetBundleName() string { return "cisco_ios_xr" }

func (vrf *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf) GetYangName() string { return "vrf" }

func (vrf *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrf *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrf *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrf *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf) SetParent(parent types.Entity) { vrf.parent = parent }

func (vrf *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf) GetParent() types.Entity { return vrf.parent }

func (vrf *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf) GetParentYangName() string { return "vrfs" }

// Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics
// IPv6 DHCP relay statistics
type Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DHCPV6 solicit packets.
    Solicit Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Solicit

    // DHCPV6 advertise packets.
    Advertise Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Advertise

    // DHCPV6 request packets.
    Request Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Request

    // DHCPV6 reply packets.
    Reply Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Reply

    // DHCPV6 confirm packets.
    Confirm Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Confirm

    // DHCPV6 decline packets.
    Decline Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Decline

    // DHCPV6 renew packets.
    Renew Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Renew

    // DHCPV6 rebind packets.
    Rebind Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Rebind

    // DHCPV6 release packets.
    Release Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Release

    // DHCPV6 reconfig packets.
    Reconfig Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Reconfig

    // DHCPV6 inform packets.
    Inform Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Inform

    // DHCPV6 relay forward packets.
    RelayForward Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_RelayForward

    // DHCPV6 relay reply packets.
    RelayReply Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_RelayReply

    // DHCPV6 lease query packets.
    LeaseQuery Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQuery

    // DHCPV6 lease query reply packets.
    LeaseQueryReply Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQueryReply

    // DHCPV6 lease query done packets.
    LeaseQueryDone Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQueryDone

    // DHCPV6 lease query data packets.
    LeaseQueryData Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQueryData
}

func (statistics *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics) GetGoName(yname string) string {
    if yname == "solicit" { return "Solicit" }
    if yname == "advertise" { return "Advertise" }
    if yname == "request" { return "Request" }
    if yname == "reply" { return "Reply" }
    if yname == "confirm" { return "Confirm" }
    if yname == "decline" { return "Decline" }
    if yname == "renew" { return "Renew" }
    if yname == "rebind" { return "Rebind" }
    if yname == "release" { return "Release" }
    if yname == "reconfig" { return "Reconfig" }
    if yname == "inform" { return "Inform" }
    if yname == "relay-forward" { return "RelayForward" }
    if yname == "relay-reply" { return "RelayReply" }
    if yname == "lease-query" { return "LeaseQuery" }
    if yname == "lease-query-reply" { return "LeaseQueryReply" }
    if yname == "lease-query-done" { return "LeaseQueryDone" }
    if yname == "lease-query-data" { return "LeaseQueryData" }
    return ""
}

func (statistics *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "solicit" {
        return &statistics.Solicit
    }
    if childYangName == "advertise" {
        return &statistics.Advertise
    }
    if childYangName == "request" {
        return &statistics.Request
    }
    if childYangName == "reply" {
        return &statistics.Reply
    }
    if childYangName == "confirm" {
        return &statistics.Confirm
    }
    if childYangName == "decline" {
        return &statistics.Decline
    }
    if childYangName == "renew" {
        return &statistics.Renew
    }
    if childYangName == "rebind" {
        return &statistics.Rebind
    }
    if childYangName == "release" {
        return &statistics.Release
    }
    if childYangName == "reconfig" {
        return &statistics.Reconfig
    }
    if childYangName == "inform" {
        return &statistics.Inform
    }
    if childYangName == "relay-forward" {
        return &statistics.RelayForward
    }
    if childYangName == "relay-reply" {
        return &statistics.RelayReply
    }
    if childYangName == "lease-query" {
        return &statistics.LeaseQuery
    }
    if childYangName == "lease-query-reply" {
        return &statistics.LeaseQueryReply
    }
    if childYangName == "lease-query-done" {
        return &statistics.LeaseQueryDone
    }
    if childYangName == "lease-query-data" {
        return &statistics.LeaseQueryData
    }
    return nil
}

func (statistics *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["solicit"] = &statistics.Solicit
    children["advertise"] = &statistics.Advertise
    children["request"] = &statistics.Request
    children["reply"] = &statistics.Reply
    children["confirm"] = &statistics.Confirm
    children["decline"] = &statistics.Decline
    children["renew"] = &statistics.Renew
    children["rebind"] = &statistics.Rebind
    children["release"] = &statistics.Release
    children["reconfig"] = &statistics.Reconfig
    children["inform"] = &statistics.Inform
    children["relay-forward"] = &statistics.RelayForward
    children["relay-reply"] = &statistics.RelayReply
    children["lease-query"] = &statistics.LeaseQuery
    children["lease-query-reply"] = &statistics.LeaseQueryReply
    children["lease-query-done"] = &statistics.LeaseQueryDone
    children["lease-query-data"] = &statistics.LeaseQueryData
    return children
}

func (statistics *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (statistics *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics) GetYangName() string { return "statistics" }

func (statistics *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics) GetParentYangName() string { return "vrf" }

// Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Solicit
// DHCPV6 solicit packets
type Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Solicit struct {
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

func (solicit *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Solicit) GetFilter() yfilter.YFilter { return solicit.YFilter }

func (solicit *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Solicit) SetFilter(yf yfilter.YFilter) { solicit.YFilter = yf }

func (solicit *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Solicit) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (solicit *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Solicit) GetSegmentPath() string {
    return "solicit"
}

func (solicit *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Solicit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (solicit *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Solicit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (solicit *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Solicit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = solicit.ReceivedPackets
    leafs["transmitted-packets"] = solicit.TransmittedPackets
    leafs["dropped-packets"] = solicit.DroppedPackets
    return leafs
}

func (solicit *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Solicit) GetBundleName() string { return "cisco_ios_xr" }

func (solicit *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Solicit) GetYangName() string { return "solicit" }

func (solicit *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Solicit) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (solicit *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Solicit) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (solicit *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Solicit) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (solicit *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Solicit) SetParent(parent types.Entity) { solicit.parent = parent }

func (solicit *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Solicit) GetParent() types.Entity { return solicit.parent }

func (solicit *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Solicit) GetParentYangName() string { return "statistics" }

// Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Advertise
// DHCPV6 advertise packets
type Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Advertise struct {
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

func (advertise *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Advertise) GetFilter() yfilter.YFilter { return advertise.YFilter }

func (advertise *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Advertise) SetFilter(yf yfilter.YFilter) { advertise.YFilter = yf }

func (advertise *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Advertise) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (advertise *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Advertise) GetSegmentPath() string {
    return "advertise"
}

func (advertise *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Advertise) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (advertise *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Advertise) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (advertise *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Advertise) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = advertise.ReceivedPackets
    leafs["transmitted-packets"] = advertise.TransmittedPackets
    leafs["dropped-packets"] = advertise.DroppedPackets
    return leafs
}

func (advertise *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Advertise) GetBundleName() string { return "cisco_ios_xr" }

func (advertise *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Advertise) GetYangName() string { return "advertise" }

func (advertise *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Advertise) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (advertise *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Advertise) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (advertise *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Advertise) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (advertise *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Advertise) SetParent(parent types.Entity) { advertise.parent = parent }

func (advertise *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Advertise) GetParent() types.Entity { return advertise.parent }

func (advertise *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Advertise) GetParentYangName() string { return "statistics" }

// Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Request
// DHCPV6 request packets
type Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Request struct {
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

func (request *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Request) GetFilter() yfilter.YFilter { return request.YFilter }

func (request *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Request) SetFilter(yf yfilter.YFilter) { request.YFilter = yf }

func (request *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Request) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (request *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Request) GetSegmentPath() string {
    return "request"
}

func (request *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Request) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (request *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Request) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (request *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Request) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = request.ReceivedPackets
    leafs["transmitted-packets"] = request.TransmittedPackets
    leafs["dropped-packets"] = request.DroppedPackets
    return leafs
}

func (request *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Request) GetBundleName() string { return "cisco_ios_xr" }

func (request *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Request) GetYangName() string { return "request" }

func (request *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Request) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (request *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Request) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (request *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Request) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (request *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Request) SetParent(parent types.Entity) { request.parent = parent }

func (request *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Request) GetParent() types.Entity { return request.parent }

func (request *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Request) GetParentYangName() string { return "statistics" }

// Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Reply
// DHCPV6 reply packets
type Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Reply struct {
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

func (reply *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Reply) GetFilter() yfilter.YFilter { return reply.YFilter }

func (reply *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Reply) SetFilter(yf yfilter.YFilter) { reply.YFilter = yf }

func (reply *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Reply) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (reply *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Reply) GetSegmentPath() string {
    return "reply"
}

func (reply *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Reply) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (reply *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Reply) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (reply *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Reply) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = reply.ReceivedPackets
    leafs["transmitted-packets"] = reply.TransmittedPackets
    leafs["dropped-packets"] = reply.DroppedPackets
    return leafs
}

func (reply *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Reply) GetBundleName() string { return "cisco_ios_xr" }

func (reply *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Reply) GetYangName() string { return "reply" }

func (reply *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Reply) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (reply *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Reply) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (reply *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Reply) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (reply *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Reply) SetParent(parent types.Entity) { reply.parent = parent }

func (reply *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Reply) GetParent() types.Entity { return reply.parent }

func (reply *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Reply) GetParentYangName() string { return "statistics" }

// Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Confirm
// DHCPV6 confirm packets
type Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Confirm struct {
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

func (confirm *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Confirm) GetFilter() yfilter.YFilter { return confirm.YFilter }

func (confirm *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Confirm) SetFilter(yf yfilter.YFilter) { confirm.YFilter = yf }

func (confirm *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Confirm) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (confirm *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Confirm) GetSegmentPath() string {
    return "confirm"
}

func (confirm *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Confirm) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (confirm *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Confirm) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (confirm *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Confirm) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = confirm.ReceivedPackets
    leafs["transmitted-packets"] = confirm.TransmittedPackets
    leafs["dropped-packets"] = confirm.DroppedPackets
    return leafs
}

func (confirm *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Confirm) GetBundleName() string { return "cisco_ios_xr" }

func (confirm *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Confirm) GetYangName() string { return "confirm" }

func (confirm *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Confirm) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (confirm *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Confirm) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (confirm *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Confirm) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (confirm *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Confirm) SetParent(parent types.Entity) { confirm.parent = parent }

func (confirm *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Confirm) GetParent() types.Entity { return confirm.parent }

func (confirm *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Confirm) GetParentYangName() string { return "statistics" }

// Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Decline
// DHCPV6 decline packets
type Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Decline struct {
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

func (decline *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Decline) GetFilter() yfilter.YFilter { return decline.YFilter }

func (decline *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Decline) SetFilter(yf yfilter.YFilter) { decline.YFilter = yf }

func (decline *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Decline) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (decline *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Decline) GetSegmentPath() string {
    return "decline"
}

func (decline *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Decline) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (decline *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Decline) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (decline *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Decline) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = decline.ReceivedPackets
    leafs["transmitted-packets"] = decline.TransmittedPackets
    leafs["dropped-packets"] = decline.DroppedPackets
    return leafs
}

func (decline *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Decline) GetBundleName() string { return "cisco_ios_xr" }

func (decline *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Decline) GetYangName() string { return "decline" }

func (decline *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Decline) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (decline *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Decline) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (decline *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Decline) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (decline *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Decline) SetParent(parent types.Entity) { decline.parent = parent }

func (decline *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Decline) GetParent() types.Entity { return decline.parent }

func (decline *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Decline) GetParentYangName() string { return "statistics" }

// Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Renew
// DHCPV6 renew packets
type Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Renew struct {
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

func (renew *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Renew) GetFilter() yfilter.YFilter { return renew.YFilter }

func (renew *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Renew) SetFilter(yf yfilter.YFilter) { renew.YFilter = yf }

func (renew *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Renew) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (renew *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Renew) GetSegmentPath() string {
    return "renew"
}

func (renew *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Renew) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (renew *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Renew) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (renew *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Renew) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = renew.ReceivedPackets
    leafs["transmitted-packets"] = renew.TransmittedPackets
    leafs["dropped-packets"] = renew.DroppedPackets
    return leafs
}

func (renew *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Renew) GetBundleName() string { return "cisco_ios_xr" }

func (renew *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Renew) GetYangName() string { return "renew" }

func (renew *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Renew) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (renew *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Renew) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (renew *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Renew) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (renew *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Renew) SetParent(parent types.Entity) { renew.parent = parent }

func (renew *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Renew) GetParent() types.Entity { return renew.parent }

func (renew *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Renew) GetParentYangName() string { return "statistics" }

// Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Rebind
// DHCPV6 rebind packets
type Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Rebind struct {
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

func (rebind *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Rebind) GetFilter() yfilter.YFilter { return rebind.YFilter }

func (rebind *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Rebind) SetFilter(yf yfilter.YFilter) { rebind.YFilter = yf }

func (rebind *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Rebind) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (rebind *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Rebind) GetSegmentPath() string {
    return "rebind"
}

func (rebind *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Rebind) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rebind *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Rebind) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rebind *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Rebind) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = rebind.ReceivedPackets
    leafs["transmitted-packets"] = rebind.TransmittedPackets
    leafs["dropped-packets"] = rebind.DroppedPackets
    return leafs
}

func (rebind *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Rebind) GetBundleName() string { return "cisco_ios_xr" }

func (rebind *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Rebind) GetYangName() string { return "rebind" }

func (rebind *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Rebind) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rebind *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Rebind) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rebind *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Rebind) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rebind *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Rebind) SetParent(parent types.Entity) { rebind.parent = parent }

func (rebind *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Rebind) GetParent() types.Entity { return rebind.parent }

func (rebind *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Rebind) GetParentYangName() string { return "statistics" }

// Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Release
// DHCPV6 release packets
type Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Release struct {
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

func (release *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Release) GetFilter() yfilter.YFilter { return release.YFilter }

func (release *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Release) SetFilter(yf yfilter.YFilter) { release.YFilter = yf }

func (release *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Release) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (release *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Release) GetSegmentPath() string {
    return "release"
}

func (release *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Release) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (release *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Release) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (release *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Release) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = release.ReceivedPackets
    leafs["transmitted-packets"] = release.TransmittedPackets
    leafs["dropped-packets"] = release.DroppedPackets
    return leafs
}

func (release *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Release) GetBundleName() string { return "cisco_ios_xr" }

func (release *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Release) GetYangName() string { return "release" }

func (release *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Release) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (release *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Release) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (release *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Release) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (release *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Release) SetParent(parent types.Entity) { release.parent = parent }

func (release *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Release) GetParent() types.Entity { return release.parent }

func (release *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Release) GetParentYangName() string { return "statistics" }

// Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Reconfig
// DHCPV6 reconfig packets
type Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Reconfig struct {
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

func (reconfig *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Reconfig) GetFilter() yfilter.YFilter { return reconfig.YFilter }

func (reconfig *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Reconfig) SetFilter(yf yfilter.YFilter) { reconfig.YFilter = yf }

func (reconfig *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Reconfig) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (reconfig *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Reconfig) GetSegmentPath() string {
    return "reconfig"
}

func (reconfig *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Reconfig) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (reconfig *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Reconfig) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (reconfig *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Reconfig) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = reconfig.ReceivedPackets
    leafs["transmitted-packets"] = reconfig.TransmittedPackets
    leafs["dropped-packets"] = reconfig.DroppedPackets
    return leafs
}

func (reconfig *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Reconfig) GetBundleName() string { return "cisco_ios_xr" }

func (reconfig *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Reconfig) GetYangName() string { return "reconfig" }

func (reconfig *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Reconfig) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (reconfig *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Reconfig) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (reconfig *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Reconfig) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (reconfig *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Reconfig) SetParent(parent types.Entity) { reconfig.parent = parent }

func (reconfig *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Reconfig) GetParent() types.Entity { return reconfig.parent }

func (reconfig *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Reconfig) GetParentYangName() string { return "statistics" }

// Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Inform
// DHCPV6 inform packets
type Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Inform struct {
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

func (inform *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Inform) GetFilter() yfilter.YFilter { return inform.YFilter }

func (inform *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Inform) SetFilter(yf yfilter.YFilter) { inform.YFilter = yf }

func (inform *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Inform) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (inform *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Inform) GetSegmentPath() string {
    return "inform"
}

func (inform *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Inform) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inform *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Inform) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inform *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Inform) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = inform.ReceivedPackets
    leafs["transmitted-packets"] = inform.TransmittedPackets
    leafs["dropped-packets"] = inform.DroppedPackets
    return leafs
}

func (inform *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Inform) GetBundleName() string { return "cisco_ios_xr" }

func (inform *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Inform) GetYangName() string { return "inform" }

func (inform *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Inform) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inform *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Inform) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inform *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Inform) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inform *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Inform) SetParent(parent types.Entity) { inform.parent = parent }

func (inform *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Inform) GetParent() types.Entity { return inform.parent }

func (inform *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Inform) GetParentYangName() string { return "statistics" }

// Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_RelayForward
// DHCPV6 relay forward packets
type Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_RelayForward struct {
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

func (relayForward *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_RelayForward) GetFilter() yfilter.YFilter { return relayForward.YFilter }

func (relayForward *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_RelayForward) SetFilter(yf yfilter.YFilter) { relayForward.YFilter = yf }

func (relayForward *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_RelayForward) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (relayForward *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_RelayForward) GetSegmentPath() string {
    return "relay-forward"
}

func (relayForward *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_RelayForward) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (relayForward *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_RelayForward) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (relayForward *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_RelayForward) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = relayForward.ReceivedPackets
    leafs["transmitted-packets"] = relayForward.TransmittedPackets
    leafs["dropped-packets"] = relayForward.DroppedPackets
    return leafs
}

func (relayForward *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_RelayForward) GetBundleName() string { return "cisco_ios_xr" }

func (relayForward *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_RelayForward) GetYangName() string { return "relay-forward" }

func (relayForward *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_RelayForward) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (relayForward *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_RelayForward) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (relayForward *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_RelayForward) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (relayForward *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_RelayForward) SetParent(parent types.Entity) { relayForward.parent = parent }

func (relayForward *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_RelayForward) GetParent() types.Entity { return relayForward.parent }

func (relayForward *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_RelayForward) GetParentYangName() string { return "statistics" }

// Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_RelayReply
// DHCPV6 relay reply packets
type Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_RelayReply struct {
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

func (relayReply *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_RelayReply) GetFilter() yfilter.YFilter { return relayReply.YFilter }

func (relayReply *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_RelayReply) SetFilter(yf yfilter.YFilter) { relayReply.YFilter = yf }

func (relayReply *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_RelayReply) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (relayReply *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_RelayReply) GetSegmentPath() string {
    return "relay-reply"
}

func (relayReply *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_RelayReply) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (relayReply *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_RelayReply) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (relayReply *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_RelayReply) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = relayReply.ReceivedPackets
    leafs["transmitted-packets"] = relayReply.TransmittedPackets
    leafs["dropped-packets"] = relayReply.DroppedPackets
    return leafs
}

func (relayReply *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_RelayReply) GetBundleName() string { return "cisco_ios_xr" }

func (relayReply *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_RelayReply) GetYangName() string { return "relay-reply" }

func (relayReply *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_RelayReply) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (relayReply *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_RelayReply) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (relayReply *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_RelayReply) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (relayReply *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_RelayReply) SetParent(parent types.Entity) { relayReply.parent = parent }

func (relayReply *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_RelayReply) GetParent() types.Entity { return relayReply.parent }

func (relayReply *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_RelayReply) GetParentYangName() string { return "statistics" }

// Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQuery
// DHCPV6 lease query packets
type Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQuery struct {
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

func (leaseQuery *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQuery) GetFilter() yfilter.YFilter { return leaseQuery.YFilter }

func (leaseQuery *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQuery) SetFilter(yf yfilter.YFilter) { leaseQuery.YFilter = yf }

func (leaseQuery *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQuery) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (leaseQuery *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQuery) GetSegmentPath() string {
    return "lease-query"
}

func (leaseQuery *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQuery) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (leaseQuery *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQuery) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (leaseQuery *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQuery) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = leaseQuery.ReceivedPackets
    leafs["transmitted-packets"] = leaseQuery.TransmittedPackets
    leafs["dropped-packets"] = leaseQuery.DroppedPackets
    return leafs
}

func (leaseQuery *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQuery) GetBundleName() string { return "cisco_ios_xr" }

func (leaseQuery *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQuery) GetYangName() string { return "lease-query" }

func (leaseQuery *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQuery) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (leaseQuery *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQuery) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (leaseQuery *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQuery) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (leaseQuery *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQuery) SetParent(parent types.Entity) { leaseQuery.parent = parent }

func (leaseQuery *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQuery) GetParent() types.Entity { return leaseQuery.parent }

func (leaseQuery *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQuery) GetParentYangName() string { return "statistics" }

// Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQueryReply
// DHCPV6 lease query reply packets
type Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQueryReply struct {
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

func (leaseQueryReply *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQueryReply) GetFilter() yfilter.YFilter { return leaseQueryReply.YFilter }

func (leaseQueryReply *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQueryReply) SetFilter(yf yfilter.YFilter) { leaseQueryReply.YFilter = yf }

func (leaseQueryReply *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQueryReply) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (leaseQueryReply *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQueryReply) GetSegmentPath() string {
    return "lease-query-reply"
}

func (leaseQueryReply *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQueryReply) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (leaseQueryReply *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQueryReply) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (leaseQueryReply *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQueryReply) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = leaseQueryReply.ReceivedPackets
    leafs["transmitted-packets"] = leaseQueryReply.TransmittedPackets
    leafs["dropped-packets"] = leaseQueryReply.DroppedPackets
    return leafs
}

func (leaseQueryReply *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQueryReply) GetBundleName() string { return "cisco_ios_xr" }

func (leaseQueryReply *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQueryReply) GetYangName() string { return "lease-query-reply" }

func (leaseQueryReply *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQueryReply) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (leaseQueryReply *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQueryReply) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (leaseQueryReply *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQueryReply) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (leaseQueryReply *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQueryReply) SetParent(parent types.Entity) { leaseQueryReply.parent = parent }

func (leaseQueryReply *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQueryReply) GetParent() types.Entity { return leaseQueryReply.parent }

func (leaseQueryReply *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQueryReply) GetParentYangName() string { return "statistics" }

// Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQueryDone
// DHCPV6 lease query done packets
type Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQueryDone struct {
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

func (leaseQueryDone *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQueryDone) GetFilter() yfilter.YFilter { return leaseQueryDone.YFilter }

func (leaseQueryDone *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQueryDone) SetFilter(yf yfilter.YFilter) { leaseQueryDone.YFilter = yf }

func (leaseQueryDone *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQueryDone) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (leaseQueryDone *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQueryDone) GetSegmentPath() string {
    return "lease-query-done"
}

func (leaseQueryDone *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQueryDone) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (leaseQueryDone *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQueryDone) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (leaseQueryDone *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQueryDone) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = leaseQueryDone.ReceivedPackets
    leafs["transmitted-packets"] = leaseQueryDone.TransmittedPackets
    leafs["dropped-packets"] = leaseQueryDone.DroppedPackets
    return leafs
}

func (leaseQueryDone *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQueryDone) GetBundleName() string { return "cisco_ios_xr" }

func (leaseQueryDone *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQueryDone) GetYangName() string { return "lease-query-done" }

func (leaseQueryDone *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQueryDone) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (leaseQueryDone *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQueryDone) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (leaseQueryDone *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQueryDone) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (leaseQueryDone *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQueryDone) SetParent(parent types.Entity) { leaseQueryDone.parent = parent }

func (leaseQueryDone *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQueryDone) GetParent() types.Entity { return leaseQueryDone.parent }

func (leaseQueryDone *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQueryDone) GetParentYangName() string { return "statistics" }

// Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQueryData
// DHCPV6 lease query data packets
type Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQueryData struct {
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

func (leaseQueryData *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQueryData) GetFilter() yfilter.YFilter { return leaseQueryData.YFilter }

func (leaseQueryData *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQueryData) SetFilter(yf yfilter.YFilter) { leaseQueryData.YFilter = yf }

func (leaseQueryData *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQueryData) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (leaseQueryData *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQueryData) GetSegmentPath() string {
    return "lease-query-data"
}

func (leaseQueryData *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQueryData) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (leaseQueryData *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQueryData) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (leaseQueryData *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQueryData) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = leaseQueryData.ReceivedPackets
    leafs["transmitted-packets"] = leaseQueryData.TransmittedPackets
    leafs["dropped-packets"] = leaseQueryData.DroppedPackets
    return leafs
}

func (leaseQueryData *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQueryData) GetBundleName() string { return "cisco_ios_xr" }

func (leaseQueryData *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQueryData) GetYangName() string { return "lease-query-data" }

func (leaseQueryData *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQueryData) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (leaseQueryData *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQueryData) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (leaseQueryData *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQueryData) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (leaseQueryData *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQueryData) SetParent(parent types.Entity) { leaseQueryData.parent = parent }

func (leaseQueryData *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQueryData) GetParent() types.Entity { return leaseQueryData.parent }

func (leaseQueryData *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQueryData) GetParentYangName() string { return "statistics" }

