// This module contains a collection of YANG definitions
// for Cisco IOS-XR ipv6-new-dhcpv6d package operational data.
// 
// This module contains definitions
// for the following management objects:
//   dhcpv6: IPV6 DHCPD operational data
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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

// BagDhcpv6dSubMode represents Bag dhcpv6d sub mode
type BagDhcpv6dSubMode string

const (
    // DHCPv6 Base mode
    BagDhcpv6dSubMode_base BagDhcpv6dSubMode = "base"

    // DHCPv6 Server mode
    BagDhcpv6dSubMode_server BagDhcpv6dSubMode = "server"

    // DHCPv6 Proxy mode
    BagDhcpv6dSubMode_proxy BagDhcpv6dSubMode = "proxy"
)

// BagDhcpv6dFsmState represents Bag dhcpv6d fsm state
type BagDhcpv6dFsmState string

const (
    // Server initializing state for client binding
    BagDhcpv6dFsmState_server_initializing BagDhcpv6dFsmState = "server-initializing"

    // Server waiting on DPM to validate subscriber
    BagDhcpv6dFsmState_server_waiting_dpm BagDhcpv6dFsmState = "server-waiting-dpm"

    // Server waiting on DAPS to assign/free
    // addr/prefix
    BagDhcpv6dFsmState_server_waiting_daps BagDhcpv6dFsmState = "server-waiting-daps"

    // Server waiting for a request from the client
    BagDhcpv6dFsmState_server_waiting_client BagDhcpv6dFsmState = "server-waiting-client"

    // Server waiting for iedge response for the
    // session
    BagDhcpv6dFsmState_server_waiting_subscriber BagDhcpv6dFsmState = "server-waiting-subscriber"

    // Server waiting for RIB response for route add
    BagDhcpv6dFsmState_server_waiting_rib BagDhcpv6dFsmState = "server-waiting-rib"

    // Server bound state to the client
    BagDhcpv6dFsmState_server_bound_client BagDhcpv6dFsmState = "server-bound-client"

    // Proxy initializing state for client binding
    BagDhcpv6dFsmState_proxy_initializing BagDhcpv6dFsmState = "proxy-initializing"

    // Proxy waiting on DPM to validate subscriber
    BagDhcpv6dFsmState_proxy_waiting_dpm BagDhcpv6dFsmState = "proxy-waiting-dpm"

    // Proxy waiting on DAPS to assign/free prefix(ND)
    BagDhcpv6dFsmState_proxy_waiting_daps BagDhcpv6dFsmState = "proxy-waiting-daps"

    // Proxy waiting for a msg from the client/srv
    BagDhcpv6dFsmState_proxy_waiting_client_server BagDhcpv6dFsmState = "proxy-waiting-client-server"

    // Proxy waiting on iedge to sub session resp
    BagDhcpv6dFsmState_proxy_waiting_subscriber BagDhcpv6dFsmState = "proxy-waiting-subscriber"

    // Proxy waiting on RIB response
    BagDhcpv6dFsmState_proxy_waiting_rib BagDhcpv6dFsmState = "proxy-waiting-rib"

    // Proxy bound state to the client
    BagDhcpv6dFsmState_proxy_bound_client BagDhcpv6dFsmState = "proxy-bound-client"
)

// BagDhcpv6dIntfSrgRole represents Bag dhcpv6d intf srg role
type BagDhcpv6dIntfSrgRole string

const (
    // DHCPv6 Interface SRG role NONE
    BagDhcpv6dIntfSrgRole_none BagDhcpv6dIntfSrgRole = "none"

    // DHCPv6 Interface SRG role Master
    BagDhcpv6dIntfSrgRole_master BagDhcpv6dIntfSrgRole = "master"

    // DHCPv6 Interface SRG role Slave
    BagDhcpv6dIntfSrgRole_slave BagDhcpv6dIntfSrgRole = "slave"
)

// Dhcpv6IssuVersion represents Dhcpv6 issu version
type Dhcpv6IssuVersion string

const (
    // Version 1
    Dhcpv6IssuVersion_version1 Dhcpv6IssuVersion = "version1"

    // Version 2
    Dhcpv6IssuVersion_version2 Dhcpv6IssuVersion = "version2"
)

// BagDhcpv6dIaId represents Bag dhcpv6d ia id
type BagDhcpv6dIaId string

const (
    // Non-temporary Addresses assigned 
    BagDhcpv6dIaId_iana BagDhcpv6dIaId = "iana"

    // Prefix delegeated to client      
    BagDhcpv6dIaId_iapd BagDhcpv6dIaId = "iapd"

    // Temporary Addresses - not supported 
    BagDhcpv6dIaId_iata BagDhcpv6dIaId = "iata"
)

// Dhcpv6IssuRole represents Dhcpv6 issu role
type Dhcpv6IssuRole string

const (
    // Primary role
    Dhcpv6IssuRole_role_primary Dhcpv6IssuRole = "role-primary"

    // Secondary role
    Dhcpv6IssuRole_role_secondary Dhcpv6IssuRole = "role-secondary"
)

// BagDhcpv6dIntfSergRole represents Bag dhcpv6d intf serg role
type BagDhcpv6dIntfSergRole string

const (
    // DHCPv6 Interface SERG role NONE
    BagDhcpv6dIntfSergRole_none BagDhcpv6dIntfSergRole = "none"

    // DHCPv6 Interface SERG role Master
    BagDhcpv6dIntfSergRole_master BagDhcpv6dIntfSergRole = "master"

    // DHCPv6 Interface SERG role Slave
    BagDhcpv6dIntfSergRole_slave BagDhcpv6dIntfSergRole = "slave"
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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DHCP IssuStatus.
    IssuStatus Dhcpv6_IssuStatus

    // IPv6 DHCP list of nodes.
    Nodes Dhcpv6_Nodes
}

func (dhcpv6 *Dhcpv6) GetEntityData() *types.CommonEntityData {
    dhcpv6.EntityData.YFilter = dhcpv6.YFilter
    dhcpv6.EntityData.YangName = "dhcpv6"
    dhcpv6.EntityData.BundleName = "cisco_ios_xr"
    dhcpv6.EntityData.ParentYangName = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper"
    dhcpv6.EntityData.SegmentPath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6"
    dhcpv6.EntityData.AbsolutePath = dhcpv6.EntityData.SegmentPath
    dhcpv6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dhcpv6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dhcpv6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dhcpv6.EntityData.Children = types.NewOrderedMap()
    dhcpv6.EntityData.Children.Append("issu-status", types.YChild{"IssuStatus", &dhcpv6.IssuStatus})
    dhcpv6.EntityData.Children.Append("nodes", types.YChild{"Nodes", &dhcpv6.Nodes})
    dhcpv6.EntityData.Leafs = types.NewOrderedMap()

    dhcpv6.EntityData.YListKeys = []string {}

    return &(dhcpv6.EntityData)
}

// Dhcpv6_IssuStatus
// DHCP IssuStatus
type Dhcpv6_IssuStatus struct {
    EntityData types.CommonEntityData
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

func (issuStatus *Dhcpv6_IssuStatus) GetEntityData() *types.CommonEntityData {
    issuStatus.EntityData.YFilter = issuStatus.YFilter
    issuStatus.EntityData.YangName = "issu-status"
    issuStatus.EntityData.BundleName = "cisco_ios_xr"
    issuStatus.EntityData.ParentYangName = "dhcpv6"
    issuStatus.EntityData.SegmentPath = "issu-status"
    issuStatus.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/" + issuStatus.EntityData.SegmentPath
    issuStatus.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    issuStatus.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    issuStatus.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    issuStatus.EntityData.Children = types.NewOrderedMap()
    issuStatus.EntityData.Leafs = types.NewOrderedMap()
    issuStatus.EntityData.Leafs.Append("process-start-time", types.YLeaf{"ProcessStartTime", issuStatus.ProcessStartTime})
    issuStatus.EntityData.Leafs.Append("issu-sync-complete-time", types.YLeaf{"IssuSyncCompleteTime", issuStatus.IssuSyncCompleteTime})
    issuStatus.EntityData.Leafs.Append("issu-sync-start-time", types.YLeaf{"IssuSyncStartTime", issuStatus.IssuSyncStartTime})
    issuStatus.EntityData.Leafs.Append("issu-ready-time", types.YLeaf{"IssuReadyTime", issuStatus.IssuReadyTime})
    issuStatus.EntityData.Leafs.Append("big-bang-time", types.YLeaf{"BigBangTime", issuStatus.BigBangTime})
    issuStatus.EntityData.Leafs.Append("primary-role-time", types.YLeaf{"PrimaryRoleTime", issuStatus.PrimaryRoleTime})
    issuStatus.EntityData.Leafs.Append("issu-ready-issu-mgr-connection", types.YLeaf{"IssuReadyIssuMgrConnection", issuStatus.IssuReadyIssuMgrConnection})
    issuStatus.EntityData.Leafs.Append("role", types.YLeaf{"Role", issuStatus.Role})
    issuStatus.EntityData.Leafs.Append("phase", types.YLeaf{"Phase", issuStatus.Phase})
    issuStatus.EntityData.Leafs.Append("version", types.YLeaf{"Version", issuStatus.Version})

    issuStatus.EntityData.YListKeys = []string {}

    return &(issuStatus.EntityData)
}

// Dhcpv6_Nodes
// IPv6 DHCP list of nodes
type Dhcpv6_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 DHCP particular node name. The type is slice of Dhcpv6_Nodes_Node.
    Node []*Dhcpv6_Nodes_Node
}

func (nodes *Dhcpv6_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "dhcpv6"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/" + nodes.EntityData.SegmentPath
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

// Dhcpv6_Nodes_Node
// IPv6 DHCP particular node name
type Dhcpv6_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (node *Dhcpv6_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeName, "node-name")
    node.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/" + node.EntityData.SegmentPath
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("proxy", types.YChild{"Proxy", &node.Proxy})
    node.EntityData.Children.Append("base", types.YChild{"Base", &node.Base})
    node.EntityData.Children.Append("server", types.YChild{"Server", &node.Server})
    node.EntityData.Children.Append("relay", types.YChild{"Relay", &node.Relay})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", node.NodeName})

    node.EntityData.YListKeys = []string {"NodeName"}

    return &(node.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy
// IPv6 DHCP proxy operational data
type Dhcpv6_Nodes_Node_Proxy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DHCPV6 proxy list of VRF names.
    Vrfs Dhcpv6_Nodes_Node_Proxy_Vrfs

    // IPv6 DHCP proxy profile.
    Profiles Dhcpv6_Nodes_Node_Proxy_Profiles

    // DHCPv6 proxy disconnect history.
    DisconnectHistories Dhcpv6_Nodes_Node_Proxy_DisconnectHistories

    // DHCPV6 proxy interface.
    Interfaces Dhcpv6_Nodes_Node_Proxy_Interfaces

    // DHCPv6 proxy statistics.
    Statistics Dhcpv6_Nodes_Node_Proxy_Statistics

    // DHCPV6 proxy bindings.
    Binding Dhcpv6_Nodes_Node_Proxy_Binding
}

func (proxy *Dhcpv6_Nodes_Node_Proxy) GetEntityData() *types.CommonEntityData {
    proxy.EntityData.YFilter = proxy.YFilter
    proxy.EntityData.YangName = "proxy"
    proxy.EntityData.BundleName = "cisco_ios_xr"
    proxy.EntityData.ParentYangName = "node"
    proxy.EntityData.SegmentPath = "proxy"
    proxy.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/" + proxy.EntityData.SegmentPath
    proxy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    proxy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    proxy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    proxy.EntityData.Children = types.NewOrderedMap()
    proxy.EntityData.Children.Append("vrfs", types.YChild{"Vrfs", &proxy.Vrfs})
    proxy.EntityData.Children.Append("profiles", types.YChild{"Profiles", &proxy.Profiles})
    proxy.EntityData.Children.Append("disconnect-histories", types.YChild{"DisconnectHistories", &proxy.DisconnectHistories})
    proxy.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &proxy.Interfaces})
    proxy.EntityData.Children.Append("statistics", types.YChild{"Statistics", &proxy.Statistics})
    proxy.EntityData.Children.Append("binding", types.YChild{"Binding", &proxy.Binding})
    proxy.EntityData.Leafs = types.NewOrderedMap()

    proxy.EntityData.YListKeys = []string {}

    return &(proxy.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Vrfs
// DHCPV6 proxy list of VRF names
type Dhcpv6_Nodes_Node_Proxy_Vrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 DHCP proxy VRF name. The type is slice of
    // Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf.
    Vrf []*Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf
}

func (vrfs *Dhcpv6_Nodes_Node_Proxy_Vrfs) GetEntityData() *types.CommonEntityData {
    vrfs.EntityData.YFilter = vrfs.YFilter
    vrfs.EntityData.YangName = "vrfs"
    vrfs.EntityData.BundleName = "cisco_ios_xr"
    vrfs.EntityData.ParentYangName = "proxy"
    vrfs.EntityData.SegmentPath = "vrfs"
    vrfs.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/proxy/" + vrfs.EntityData.SegmentPath
    vrfs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfs.EntityData.Children = types.NewOrderedMap()
    vrfs.EntityData.Children.Append("vrf", types.YChild{"Vrf", nil})
    for i := range vrfs.Vrf {
        vrfs.EntityData.Children.Append(types.GetSegmentPath(vrfs.Vrf[i]), types.YChild{"Vrf", vrfs.Vrf[i]})
    }
    vrfs.EntityData.Leafs = types.NewOrderedMap()

    vrfs.EntityData.YListKeys = []string {}

    return &(vrfs.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf
// IPv6 DHCP proxy VRF name
type Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. VRF name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // IPv6 DHCP proxy statistics.
    Statistics Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics
}

func (vrf *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "vrfs"
    vrf.EntityData.SegmentPath = "vrf" + types.AddKeyToken(vrf.VrfName, "vrf-name")
    vrf.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/proxy/vrfs/" + vrf.EntityData.SegmentPath
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = types.NewOrderedMap()
    vrf.EntityData.Children.Append("statistics", types.YChild{"Statistics", &vrf.Statistics})
    vrf.EntityData.Leafs = types.NewOrderedMap()
    vrf.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", vrf.VrfName})

    vrf.EntityData.YListKeys = []string {"VrfName"}

    return &(vrf.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics
// IPv6 DHCP proxy statistics
type Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics struct {
    EntityData types.CommonEntityData
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

func (statistics *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "vrf"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/proxy/vrfs/vrf/" + statistics.EntityData.SegmentPath
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = types.NewOrderedMap()
    statistics.EntityData.Children.Append("solicit", types.YChild{"Solicit", &statistics.Solicit})
    statistics.EntityData.Children.Append("advertise", types.YChild{"Advertise", &statistics.Advertise})
    statistics.EntityData.Children.Append("request", types.YChild{"Request", &statistics.Request})
    statistics.EntityData.Children.Append("reply", types.YChild{"Reply", &statistics.Reply})
    statistics.EntityData.Children.Append("confirm", types.YChild{"Confirm", &statistics.Confirm})
    statistics.EntityData.Children.Append("decline", types.YChild{"Decline", &statistics.Decline})
    statistics.EntityData.Children.Append("renew", types.YChild{"Renew", &statistics.Renew})
    statistics.EntityData.Children.Append("rebind", types.YChild{"Rebind", &statistics.Rebind})
    statistics.EntityData.Children.Append("release", types.YChild{"Release", &statistics.Release})
    statistics.EntityData.Children.Append("reconfig", types.YChild{"Reconfig", &statistics.Reconfig})
    statistics.EntityData.Children.Append("inform", types.YChild{"Inform", &statistics.Inform})
    statistics.EntityData.Children.Append("relay-forward", types.YChild{"RelayForward", &statistics.RelayForward})
    statistics.EntityData.Children.Append("relay-reply", types.YChild{"RelayReply", &statistics.RelayReply})
    statistics.EntityData.Children.Append("lease-query", types.YChild{"LeaseQuery", &statistics.LeaseQuery})
    statistics.EntityData.Children.Append("lease-query-reply", types.YChild{"LeaseQueryReply", &statistics.LeaseQueryReply})
    statistics.EntityData.Children.Append("lease-query-done", types.YChild{"LeaseQueryDone", &statistics.LeaseQueryDone})
    statistics.EntityData.Children.Append("lease-query-data", types.YChild{"LeaseQueryData", &statistics.LeaseQueryData})
    statistics.EntityData.Leafs = types.NewOrderedMap()

    statistics.EntityData.YListKeys = []string {}

    return &(statistics.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Solicit
// DHCPV6 solicit packets
type Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Solicit struct {
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

func (solicit *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Solicit) GetEntityData() *types.CommonEntityData {
    solicit.EntityData.YFilter = solicit.YFilter
    solicit.EntityData.YangName = "solicit"
    solicit.EntityData.BundleName = "cisco_ios_xr"
    solicit.EntityData.ParentYangName = "statistics"
    solicit.EntityData.SegmentPath = "solicit"
    solicit.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/proxy/vrfs/vrf/statistics/" + solicit.EntityData.SegmentPath
    solicit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    solicit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    solicit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    solicit.EntityData.Children = types.NewOrderedMap()
    solicit.EntityData.Leafs = types.NewOrderedMap()
    solicit.EntityData.Leafs.Append("received-packets", types.YLeaf{"ReceivedPackets", solicit.ReceivedPackets})
    solicit.EntityData.Leafs.Append("transmitted-packets", types.YLeaf{"TransmittedPackets", solicit.TransmittedPackets})
    solicit.EntityData.Leafs.Append("dropped-packets", types.YLeaf{"DroppedPackets", solicit.DroppedPackets})

    solicit.EntityData.YListKeys = []string {}

    return &(solicit.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Advertise
// DHCPV6 advertise packets
type Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Advertise struct {
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

func (advertise *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Advertise) GetEntityData() *types.CommonEntityData {
    advertise.EntityData.YFilter = advertise.YFilter
    advertise.EntityData.YangName = "advertise"
    advertise.EntityData.BundleName = "cisco_ios_xr"
    advertise.EntityData.ParentYangName = "statistics"
    advertise.EntityData.SegmentPath = "advertise"
    advertise.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/proxy/vrfs/vrf/statistics/" + advertise.EntityData.SegmentPath
    advertise.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    advertise.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    advertise.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    advertise.EntityData.Children = types.NewOrderedMap()
    advertise.EntityData.Leafs = types.NewOrderedMap()
    advertise.EntityData.Leafs.Append("received-packets", types.YLeaf{"ReceivedPackets", advertise.ReceivedPackets})
    advertise.EntityData.Leafs.Append("transmitted-packets", types.YLeaf{"TransmittedPackets", advertise.TransmittedPackets})
    advertise.EntityData.Leafs.Append("dropped-packets", types.YLeaf{"DroppedPackets", advertise.DroppedPackets})

    advertise.EntityData.YListKeys = []string {}

    return &(advertise.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Request
// DHCPV6 request packets
type Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Request struct {
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

func (request *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Request) GetEntityData() *types.CommonEntityData {
    request.EntityData.YFilter = request.YFilter
    request.EntityData.YangName = "request"
    request.EntityData.BundleName = "cisco_ios_xr"
    request.EntityData.ParentYangName = "statistics"
    request.EntityData.SegmentPath = "request"
    request.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/proxy/vrfs/vrf/statistics/" + request.EntityData.SegmentPath
    request.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    request.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    request.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    request.EntityData.Children = types.NewOrderedMap()
    request.EntityData.Leafs = types.NewOrderedMap()
    request.EntityData.Leafs.Append("received-packets", types.YLeaf{"ReceivedPackets", request.ReceivedPackets})
    request.EntityData.Leafs.Append("transmitted-packets", types.YLeaf{"TransmittedPackets", request.TransmittedPackets})
    request.EntityData.Leafs.Append("dropped-packets", types.YLeaf{"DroppedPackets", request.DroppedPackets})

    request.EntityData.YListKeys = []string {}

    return &(request.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Reply
// DHCPV6 reply packets
type Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Reply struct {
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

func (reply *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Reply) GetEntityData() *types.CommonEntityData {
    reply.EntityData.YFilter = reply.YFilter
    reply.EntityData.YangName = "reply"
    reply.EntityData.BundleName = "cisco_ios_xr"
    reply.EntityData.ParentYangName = "statistics"
    reply.EntityData.SegmentPath = "reply"
    reply.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/proxy/vrfs/vrf/statistics/" + reply.EntityData.SegmentPath
    reply.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reply.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reply.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reply.EntityData.Children = types.NewOrderedMap()
    reply.EntityData.Leafs = types.NewOrderedMap()
    reply.EntityData.Leafs.Append("received-packets", types.YLeaf{"ReceivedPackets", reply.ReceivedPackets})
    reply.EntityData.Leafs.Append("transmitted-packets", types.YLeaf{"TransmittedPackets", reply.TransmittedPackets})
    reply.EntityData.Leafs.Append("dropped-packets", types.YLeaf{"DroppedPackets", reply.DroppedPackets})

    reply.EntityData.YListKeys = []string {}

    return &(reply.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Confirm
// DHCPV6 confirm packets
type Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Confirm struct {
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

func (confirm *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Confirm) GetEntityData() *types.CommonEntityData {
    confirm.EntityData.YFilter = confirm.YFilter
    confirm.EntityData.YangName = "confirm"
    confirm.EntityData.BundleName = "cisco_ios_xr"
    confirm.EntityData.ParentYangName = "statistics"
    confirm.EntityData.SegmentPath = "confirm"
    confirm.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/proxy/vrfs/vrf/statistics/" + confirm.EntityData.SegmentPath
    confirm.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    confirm.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    confirm.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    confirm.EntityData.Children = types.NewOrderedMap()
    confirm.EntityData.Leafs = types.NewOrderedMap()
    confirm.EntityData.Leafs.Append("received-packets", types.YLeaf{"ReceivedPackets", confirm.ReceivedPackets})
    confirm.EntityData.Leafs.Append("transmitted-packets", types.YLeaf{"TransmittedPackets", confirm.TransmittedPackets})
    confirm.EntityData.Leafs.Append("dropped-packets", types.YLeaf{"DroppedPackets", confirm.DroppedPackets})

    confirm.EntityData.YListKeys = []string {}

    return &(confirm.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Decline
// DHCPV6 decline packets
type Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Decline struct {
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

func (decline *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Decline) GetEntityData() *types.CommonEntityData {
    decline.EntityData.YFilter = decline.YFilter
    decline.EntityData.YangName = "decline"
    decline.EntityData.BundleName = "cisco_ios_xr"
    decline.EntityData.ParentYangName = "statistics"
    decline.EntityData.SegmentPath = "decline"
    decline.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/proxy/vrfs/vrf/statistics/" + decline.EntityData.SegmentPath
    decline.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    decline.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    decline.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    decline.EntityData.Children = types.NewOrderedMap()
    decline.EntityData.Leafs = types.NewOrderedMap()
    decline.EntityData.Leafs.Append("received-packets", types.YLeaf{"ReceivedPackets", decline.ReceivedPackets})
    decline.EntityData.Leafs.Append("transmitted-packets", types.YLeaf{"TransmittedPackets", decline.TransmittedPackets})
    decline.EntityData.Leafs.Append("dropped-packets", types.YLeaf{"DroppedPackets", decline.DroppedPackets})

    decline.EntityData.YListKeys = []string {}

    return &(decline.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Renew
// DHCPV6 renew packets
type Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Renew struct {
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

func (renew *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Renew) GetEntityData() *types.CommonEntityData {
    renew.EntityData.YFilter = renew.YFilter
    renew.EntityData.YangName = "renew"
    renew.EntityData.BundleName = "cisco_ios_xr"
    renew.EntityData.ParentYangName = "statistics"
    renew.EntityData.SegmentPath = "renew"
    renew.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/proxy/vrfs/vrf/statistics/" + renew.EntityData.SegmentPath
    renew.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    renew.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    renew.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    renew.EntityData.Children = types.NewOrderedMap()
    renew.EntityData.Leafs = types.NewOrderedMap()
    renew.EntityData.Leafs.Append("received-packets", types.YLeaf{"ReceivedPackets", renew.ReceivedPackets})
    renew.EntityData.Leafs.Append("transmitted-packets", types.YLeaf{"TransmittedPackets", renew.TransmittedPackets})
    renew.EntityData.Leafs.Append("dropped-packets", types.YLeaf{"DroppedPackets", renew.DroppedPackets})

    renew.EntityData.YListKeys = []string {}

    return &(renew.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Rebind
// DHCPV6 rebind packets
type Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Rebind struct {
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

func (rebind *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Rebind) GetEntityData() *types.CommonEntityData {
    rebind.EntityData.YFilter = rebind.YFilter
    rebind.EntityData.YangName = "rebind"
    rebind.EntityData.BundleName = "cisco_ios_xr"
    rebind.EntityData.ParentYangName = "statistics"
    rebind.EntityData.SegmentPath = "rebind"
    rebind.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/proxy/vrfs/vrf/statistics/" + rebind.EntityData.SegmentPath
    rebind.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rebind.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rebind.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rebind.EntityData.Children = types.NewOrderedMap()
    rebind.EntityData.Leafs = types.NewOrderedMap()
    rebind.EntityData.Leafs.Append("received-packets", types.YLeaf{"ReceivedPackets", rebind.ReceivedPackets})
    rebind.EntityData.Leafs.Append("transmitted-packets", types.YLeaf{"TransmittedPackets", rebind.TransmittedPackets})
    rebind.EntityData.Leafs.Append("dropped-packets", types.YLeaf{"DroppedPackets", rebind.DroppedPackets})

    rebind.EntityData.YListKeys = []string {}

    return &(rebind.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Release
// DHCPV6 release packets
type Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Release struct {
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

func (release *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Release) GetEntityData() *types.CommonEntityData {
    release.EntityData.YFilter = release.YFilter
    release.EntityData.YangName = "release"
    release.EntityData.BundleName = "cisco_ios_xr"
    release.EntityData.ParentYangName = "statistics"
    release.EntityData.SegmentPath = "release"
    release.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/proxy/vrfs/vrf/statistics/" + release.EntityData.SegmentPath
    release.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    release.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    release.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    release.EntityData.Children = types.NewOrderedMap()
    release.EntityData.Leafs = types.NewOrderedMap()
    release.EntityData.Leafs.Append("received-packets", types.YLeaf{"ReceivedPackets", release.ReceivedPackets})
    release.EntityData.Leafs.Append("transmitted-packets", types.YLeaf{"TransmittedPackets", release.TransmittedPackets})
    release.EntityData.Leafs.Append("dropped-packets", types.YLeaf{"DroppedPackets", release.DroppedPackets})

    release.EntityData.YListKeys = []string {}

    return &(release.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Reconfig
// DHCPV6 reconfig packets
type Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Reconfig struct {
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

func (reconfig *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Reconfig) GetEntityData() *types.CommonEntityData {
    reconfig.EntityData.YFilter = reconfig.YFilter
    reconfig.EntityData.YangName = "reconfig"
    reconfig.EntityData.BundleName = "cisco_ios_xr"
    reconfig.EntityData.ParentYangName = "statistics"
    reconfig.EntityData.SegmentPath = "reconfig"
    reconfig.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/proxy/vrfs/vrf/statistics/" + reconfig.EntityData.SegmentPath
    reconfig.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reconfig.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reconfig.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reconfig.EntityData.Children = types.NewOrderedMap()
    reconfig.EntityData.Leafs = types.NewOrderedMap()
    reconfig.EntityData.Leafs.Append("received-packets", types.YLeaf{"ReceivedPackets", reconfig.ReceivedPackets})
    reconfig.EntityData.Leafs.Append("transmitted-packets", types.YLeaf{"TransmittedPackets", reconfig.TransmittedPackets})
    reconfig.EntityData.Leafs.Append("dropped-packets", types.YLeaf{"DroppedPackets", reconfig.DroppedPackets})

    reconfig.EntityData.YListKeys = []string {}

    return &(reconfig.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Inform
// DHCPV6 inform packets
type Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Inform struct {
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

func (inform *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_Inform) GetEntityData() *types.CommonEntityData {
    inform.EntityData.YFilter = inform.YFilter
    inform.EntityData.YangName = "inform"
    inform.EntityData.BundleName = "cisco_ios_xr"
    inform.EntityData.ParentYangName = "statistics"
    inform.EntityData.SegmentPath = "inform"
    inform.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/proxy/vrfs/vrf/statistics/" + inform.EntityData.SegmentPath
    inform.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inform.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inform.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inform.EntityData.Children = types.NewOrderedMap()
    inform.EntityData.Leafs = types.NewOrderedMap()
    inform.EntityData.Leafs.Append("received-packets", types.YLeaf{"ReceivedPackets", inform.ReceivedPackets})
    inform.EntityData.Leafs.Append("transmitted-packets", types.YLeaf{"TransmittedPackets", inform.TransmittedPackets})
    inform.EntityData.Leafs.Append("dropped-packets", types.YLeaf{"DroppedPackets", inform.DroppedPackets})

    inform.EntityData.YListKeys = []string {}

    return &(inform.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_RelayForward
// DHCPV6 relay forward packets
type Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_RelayForward struct {
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

func (relayForward *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_RelayForward) GetEntityData() *types.CommonEntityData {
    relayForward.EntityData.YFilter = relayForward.YFilter
    relayForward.EntityData.YangName = "relay-forward"
    relayForward.EntityData.BundleName = "cisco_ios_xr"
    relayForward.EntityData.ParentYangName = "statistics"
    relayForward.EntityData.SegmentPath = "relay-forward"
    relayForward.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/proxy/vrfs/vrf/statistics/" + relayForward.EntityData.SegmentPath
    relayForward.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    relayForward.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    relayForward.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    relayForward.EntityData.Children = types.NewOrderedMap()
    relayForward.EntityData.Leafs = types.NewOrderedMap()
    relayForward.EntityData.Leafs.Append("received-packets", types.YLeaf{"ReceivedPackets", relayForward.ReceivedPackets})
    relayForward.EntityData.Leafs.Append("transmitted-packets", types.YLeaf{"TransmittedPackets", relayForward.TransmittedPackets})
    relayForward.EntityData.Leafs.Append("dropped-packets", types.YLeaf{"DroppedPackets", relayForward.DroppedPackets})

    relayForward.EntityData.YListKeys = []string {}

    return &(relayForward.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_RelayReply
// DHCPV6 relay reply packets
type Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_RelayReply struct {
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

func (relayReply *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_RelayReply) GetEntityData() *types.CommonEntityData {
    relayReply.EntityData.YFilter = relayReply.YFilter
    relayReply.EntityData.YangName = "relay-reply"
    relayReply.EntityData.BundleName = "cisco_ios_xr"
    relayReply.EntityData.ParentYangName = "statistics"
    relayReply.EntityData.SegmentPath = "relay-reply"
    relayReply.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/proxy/vrfs/vrf/statistics/" + relayReply.EntityData.SegmentPath
    relayReply.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    relayReply.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    relayReply.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    relayReply.EntityData.Children = types.NewOrderedMap()
    relayReply.EntityData.Leafs = types.NewOrderedMap()
    relayReply.EntityData.Leafs.Append("received-packets", types.YLeaf{"ReceivedPackets", relayReply.ReceivedPackets})
    relayReply.EntityData.Leafs.Append("transmitted-packets", types.YLeaf{"TransmittedPackets", relayReply.TransmittedPackets})
    relayReply.EntityData.Leafs.Append("dropped-packets", types.YLeaf{"DroppedPackets", relayReply.DroppedPackets})

    relayReply.EntityData.YListKeys = []string {}

    return &(relayReply.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQuery
// DHCPV6 lease query packets
type Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQuery struct {
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

func (leaseQuery *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQuery) GetEntityData() *types.CommonEntityData {
    leaseQuery.EntityData.YFilter = leaseQuery.YFilter
    leaseQuery.EntityData.YangName = "lease-query"
    leaseQuery.EntityData.BundleName = "cisco_ios_xr"
    leaseQuery.EntityData.ParentYangName = "statistics"
    leaseQuery.EntityData.SegmentPath = "lease-query"
    leaseQuery.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/proxy/vrfs/vrf/statistics/" + leaseQuery.EntityData.SegmentPath
    leaseQuery.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    leaseQuery.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    leaseQuery.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    leaseQuery.EntityData.Children = types.NewOrderedMap()
    leaseQuery.EntityData.Leafs = types.NewOrderedMap()
    leaseQuery.EntityData.Leafs.Append("received-packets", types.YLeaf{"ReceivedPackets", leaseQuery.ReceivedPackets})
    leaseQuery.EntityData.Leafs.Append("transmitted-packets", types.YLeaf{"TransmittedPackets", leaseQuery.TransmittedPackets})
    leaseQuery.EntityData.Leafs.Append("dropped-packets", types.YLeaf{"DroppedPackets", leaseQuery.DroppedPackets})

    leaseQuery.EntityData.YListKeys = []string {}

    return &(leaseQuery.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQueryReply
// DHCPV6 lease query reply packets
type Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQueryReply struct {
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

func (leaseQueryReply *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQueryReply) GetEntityData() *types.CommonEntityData {
    leaseQueryReply.EntityData.YFilter = leaseQueryReply.YFilter
    leaseQueryReply.EntityData.YangName = "lease-query-reply"
    leaseQueryReply.EntityData.BundleName = "cisco_ios_xr"
    leaseQueryReply.EntityData.ParentYangName = "statistics"
    leaseQueryReply.EntityData.SegmentPath = "lease-query-reply"
    leaseQueryReply.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/proxy/vrfs/vrf/statistics/" + leaseQueryReply.EntityData.SegmentPath
    leaseQueryReply.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    leaseQueryReply.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    leaseQueryReply.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    leaseQueryReply.EntityData.Children = types.NewOrderedMap()
    leaseQueryReply.EntityData.Leafs = types.NewOrderedMap()
    leaseQueryReply.EntityData.Leafs.Append("received-packets", types.YLeaf{"ReceivedPackets", leaseQueryReply.ReceivedPackets})
    leaseQueryReply.EntityData.Leafs.Append("transmitted-packets", types.YLeaf{"TransmittedPackets", leaseQueryReply.TransmittedPackets})
    leaseQueryReply.EntityData.Leafs.Append("dropped-packets", types.YLeaf{"DroppedPackets", leaseQueryReply.DroppedPackets})

    leaseQueryReply.EntityData.YListKeys = []string {}

    return &(leaseQueryReply.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQueryDone
// DHCPV6 lease query done packets
type Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQueryDone struct {
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

func (leaseQueryDone *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQueryDone) GetEntityData() *types.CommonEntityData {
    leaseQueryDone.EntityData.YFilter = leaseQueryDone.YFilter
    leaseQueryDone.EntityData.YangName = "lease-query-done"
    leaseQueryDone.EntityData.BundleName = "cisco_ios_xr"
    leaseQueryDone.EntityData.ParentYangName = "statistics"
    leaseQueryDone.EntityData.SegmentPath = "lease-query-done"
    leaseQueryDone.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/proxy/vrfs/vrf/statistics/" + leaseQueryDone.EntityData.SegmentPath
    leaseQueryDone.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    leaseQueryDone.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    leaseQueryDone.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    leaseQueryDone.EntityData.Children = types.NewOrderedMap()
    leaseQueryDone.EntityData.Leafs = types.NewOrderedMap()
    leaseQueryDone.EntityData.Leafs.Append("received-packets", types.YLeaf{"ReceivedPackets", leaseQueryDone.ReceivedPackets})
    leaseQueryDone.EntityData.Leafs.Append("transmitted-packets", types.YLeaf{"TransmittedPackets", leaseQueryDone.TransmittedPackets})
    leaseQueryDone.EntityData.Leafs.Append("dropped-packets", types.YLeaf{"DroppedPackets", leaseQueryDone.DroppedPackets})

    leaseQueryDone.EntityData.YListKeys = []string {}

    return &(leaseQueryDone.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQueryData
// DHCPV6 lease query data packets
type Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQueryData struct {
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

func (leaseQueryData *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics_LeaseQueryData) GetEntityData() *types.CommonEntityData {
    leaseQueryData.EntityData.YFilter = leaseQueryData.YFilter
    leaseQueryData.EntityData.YangName = "lease-query-data"
    leaseQueryData.EntityData.BundleName = "cisco_ios_xr"
    leaseQueryData.EntityData.ParentYangName = "statistics"
    leaseQueryData.EntityData.SegmentPath = "lease-query-data"
    leaseQueryData.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/proxy/vrfs/vrf/statistics/" + leaseQueryData.EntityData.SegmentPath
    leaseQueryData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    leaseQueryData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    leaseQueryData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    leaseQueryData.EntityData.Children = types.NewOrderedMap()
    leaseQueryData.EntityData.Leafs = types.NewOrderedMap()
    leaseQueryData.EntityData.Leafs.Append("received-packets", types.YLeaf{"ReceivedPackets", leaseQueryData.ReceivedPackets})
    leaseQueryData.EntityData.Leafs.Append("transmitted-packets", types.YLeaf{"TransmittedPackets", leaseQueryData.TransmittedPackets})
    leaseQueryData.EntityData.Leafs.Append("dropped-packets", types.YLeaf{"DroppedPackets", leaseQueryData.DroppedPackets})

    leaseQueryData.EntityData.YListKeys = []string {}

    return &(leaseQueryData.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Profiles
// IPv6 DHCP proxy profile
type Dhcpv6_Nodes_Node_Proxy_Profiles struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 DHCP proxy profile. The type is slice of
    // Dhcpv6_Nodes_Node_Proxy_Profiles_Profile.
    Profile []*Dhcpv6_Nodes_Node_Proxy_Profiles_Profile
}

func (profiles *Dhcpv6_Nodes_Node_Proxy_Profiles) GetEntityData() *types.CommonEntityData {
    profiles.EntityData.YFilter = profiles.YFilter
    profiles.EntityData.YangName = "profiles"
    profiles.EntityData.BundleName = "cisco_ios_xr"
    profiles.EntityData.ParentYangName = "proxy"
    profiles.EntityData.SegmentPath = "profiles"
    profiles.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/proxy/" + profiles.EntityData.SegmentPath
    profiles.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    profiles.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    profiles.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    profiles.EntityData.Children = types.NewOrderedMap()
    profiles.EntityData.Children.Append("profile", types.YChild{"Profile", nil})
    for i := range profiles.Profile {
        profiles.EntityData.Children.Append(types.GetSegmentPath(profiles.Profile[i]), types.YChild{"Profile", profiles.Profile[i]})
    }
    profiles.EntityData.Leafs = types.NewOrderedMap()

    profiles.EntityData.YListKeys = []string {}

    return &(profiles.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Profiles_Profile
// IPv6 DHCP proxy profile
type Dhcpv6_Nodes_Node_Proxy_Profiles_Profile struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Profile name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    ProfileName interface{}

    // DHCPV6 throttle table.
    ThrottleInfos Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_ThrottleInfos

    // IPv6 DHCP proxy class table.
    ProxyClasses Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_ProxyClasses

    // IPv6 DHCP proxy profile Info.
    Info Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info
}

func (profile *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile) GetEntityData() *types.CommonEntityData {
    profile.EntityData.YFilter = profile.YFilter
    profile.EntityData.YangName = "profile"
    profile.EntityData.BundleName = "cisco_ios_xr"
    profile.EntityData.ParentYangName = "profiles"
    profile.EntityData.SegmentPath = "profile" + types.AddKeyToken(profile.ProfileName, "profile-name")
    profile.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/proxy/profiles/" + profile.EntityData.SegmentPath
    profile.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    profile.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    profile.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    profile.EntityData.Children = types.NewOrderedMap()
    profile.EntityData.Children.Append("throttle-infos", types.YChild{"ThrottleInfos", &profile.ThrottleInfos})
    profile.EntityData.Children.Append("proxy-classes", types.YChild{"ProxyClasses", &profile.ProxyClasses})
    profile.EntityData.Children.Append("info", types.YChild{"Info", &profile.Info})
    profile.EntityData.Leafs = types.NewOrderedMap()
    profile.EntityData.Leafs.Append("profile-name", types.YLeaf{"ProfileName", profile.ProfileName})

    profile.EntityData.YListKeys = []string {"ProfileName"}

    return &(profile.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_ThrottleInfos
// DHCPV6 throttle table
type Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_ThrottleInfos struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 DHCP proxy profile Throttle Info. The type is slice of
    // Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_ThrottleInfos_ThrottleInfo.
    ThrottleInfo []*Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_ThrottleInfos_ThrottleInfo
}

func (throttleInfos *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_ThrottleInfos) GetEntityData() *types.CommonEntityData {
    throttleInfos.EntityData.YFilter = throttleInfos.YFilter
    throttleInfos.EntityData.YangName = "throttle-infos"
    throttleInfos.EntityData.BundleName = "cisco_ios_xr"
    throttleInfos.EntityData.ParentYangName = "profile"
    throttleInfos.EntityData.SegmentPath = "throttle-infos"
    throttleInfos.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/proxy/profiles/profile/" + throttleInfos.EntityData.SegmentPath
    throttleInfos.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    throttleInfos.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    throttleInfos.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    throttleInfos.EntityData.Children = types.NewOrderedMap()
    throttleInfos.EntityData.Children.Append("throttle-info", types.YChild{"ThrottleInfo", nil})
    for i := range throttleInfos.ThrottleInfo {
        throttleInfos.EntityData.Children.Append(types.GetSegmentPath(throttleInfos.ThrottleInfo[i]), types.YChild{"ThrottleInfo", throttleInfos.ThrottleInfo[i]})
    }
    throttleInfos.EntityData.Leafs = types.NewOrderedMap()

    throttleInfos.EntityData.YListKeys = []string {}

    return &(throttleInfos.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_ThrottleInfos_ThrottleInfo
// IPv6 DHCP proxy profile Throttle Info
type Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_ThrottleInfos_ThrottleInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (throttleInfo *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_ThrottleInfos_ThrottleInfo) GetEntityData() *types.CommonEntityData {
    throttleInfo.EntityData.YFilter = throttleInfo.YFilter
    throttleInfo.EntityData.YangName = "throttle-info"
    throttleInfo.EntityData.BundleName = "cisco_ios_xr"
    throttleInfo.EntityData.ParentYangName = "throttle-infos"
    throttleInfo.EntityData.SegmentPath = "throttle-info" + types.AddKeyToken(throttleInfo.MacAddress, "mac-address")
    throttleInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/proxy/profiles/profile/throttle-infos/" + throttleInfo.EntityData.SegmentPath
    throttleInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    throttleInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    throttleInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    throttleInfo.EntityData.Children = types.NewOrderedMap()
    throttleInfo.EntityData.Leafs = types.NewOrderedMap()
    throttleInfo.EntityData.Leafs.Append("mac-address", types.YLeaf{"MacAddress", throttleInfo.MacAddress})
    throttleInfo.EntityData.Leafs.Append("binding-chaddr", types.YLeaf{"BindingChaddr", throttleInfo.BindingChaddr})
    throttleInfo.EntityData.Leafs.Append("ifname", types.YLeaf{"Ifname", throttleInfo.Ifname})
    throttleInfo.EntityData.Leafs.Append("state", types.YLeaf{"State", throttleInfo.State})
    throttleInfo.EntityData.Leafs.Append("time-left", types.YLeaf{"TimeLeft", throttleInfo.TimeLeft})

    throttleInfo.EntityData.YListKeys = []string {"MacAddress"}

    return &(throttleInfo.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_ProxyClasses
// IPv6 DHCP proxy class table
type Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_ProxyClasses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 DHCP proxy class profile. The type is slice of
    // Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_ProxyClasses_ProxyClass.
    ProxyClass []*Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_ProxyClasses_ProxyClass
}

func (proxyClasses *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_ProxyClasses) GetEntityData() *types.CommonEntityData {
    proxyClasses.EntityData.YFilter = proxyClasses.YFilter
    proxyClasses.EntityData.YangName = "proxy-classes"
    proxyClasses.EntityData.BundleName = "cisco_ios_xr"
    proxyClasses.EntityData.ParentYangName = "profile"
    proxyClasses.EntityData.SegmentPath = "proxy-classes"
    proxyClasses.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/proxy/profiles/profile/" + proxyClasses.EntityData.SegmentPath
    proxyClasses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    proxyClasses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    proxyClasses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    proxyClasses.EntityData.Children = types.NewOrderedMap()
    proxyClasses.EntityData.Children.Append("proxy-class", types.YChild{"ProxyClass", nil})
    for i := range proxyClasses.ProxyClass {
        proxyClasses.EntityData.Children.Append(types.GetSegmentPath(proxyClasses.ProxyClass[i]), types.YChild{"ProxyClass", proxyClasses.ProxyClass[i]})
    }
    proxyClasses.EntityData.Leafs = types.NewOrderedMap()

    proxyClasses.EntityData.YListKeys = []string {}

    return &(proxyClasses.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_ProxyClasses_ProxyClass
// IPv6 DHCP proxy class profile
type Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_ProxyClasses_ProxyClass struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Class name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    ClassName interface{}

    // Class name. The type is string with length: 0..65.
    ClassNameXr interface{}

    // Helper addresses. The type is slice of string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    ProfileHelperAddress []interface{}

    // VRF names. The type is slice of string with length: 0..33.
    VrfName []interface{}
}

func (proxyClass *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_ProxyClasses_ProxyClass) GetEntityData() *types.CommonEntityData {
    proxyClass.EntityData.YFilter = proxyClass.YFilter
    proxyClass.EntityData.YangName = "proxy-class"
    proxyClass.EntityData.BundleName = "cisco_ios_xr"
    proxyClass.EntityData.ParentYangName = "proxy-classes"
    proxyClass.EntityData.SegmentPath = "proxy-class" + types.AddKeyToken(proxyClass.ClassName, "class-name")
    proxyClass.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/proxy/profiles/profile/proxy-classes/" + proxyClass.EntityData.SegmentPath
    proxyClass.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    proxyClass.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    proxyClass.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    proxyClass.EntityData.Children = types.NewOrderedMap()
    proxyClass.EntityData.Leafs = types.NewOrderedMap()
    proxyClass.EntityData.Leafs.Append("class-name", types.YLeaf{"ClassName", proxyClass.ClassName})
    proxyClass.EntityData.Leafs.Append("class-name-xr", types.YLeaf{"ClassNameXr", proxyClass.ClassNameXr})
    proxyClass.EntityData.Leafs.Append("profile-helper-address", types.YLeaf{"ProfileHelperAddress", proxyClass.ProfileHelperAddress})
    proxyClass.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", proxyClass.VrfName})

    proxyClass.EntityData.YListKeys = []string {"ClassName"}

    return &(proxyClass.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info
// IPv6 DHCP proxy profile Info
type Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info struct {
    EntityData types.CommonEntityData
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

func (info *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info) GetEntityData() *types.CommonEntityData {
    info.EntityData.YFilter = info.YFilter
    info.EntityData.YangName = "info"
    info.EntityData.BundleName = "cisco_ios_xr"
    info.EntityData.ParentYangName = "profile"
    info.EntityData.SegmentPath = "info"
    info.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/proxy/profiles/profile/" + info.EntityData.SegmentPath
    info.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    info.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    info.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    info.EntityData.Children = types.NewOrderedMap()
    info.EntityData.Children.Append("interface-id-references", types.YChild{"InterfaceIdReferences", &info.InterfaceIdReferences})
    info.EntityData.Children.Append("vrf-references", types.YChild{"VrfReferences", &info.VrfReferences})
    info.EntityData.Children.Append("interface-references", types.YChild{"InterfaceReferences", &info.InterfaceReferences})
    info.EntityData.Leafs = types.NewOrderedMap()
    info.EntityData.Leafs.Append("profile-name", types.YLeaf{"ProfileName", info.ProfileName})
    info.EntityData.Leafs.Append("remote-id", types.YLeaf{"RemoteId", info.RemoteId})
    info.EntityData.Leafs.Append("profile-link-address", types.YLeaf{"ProfileLinkAddress", info.ProfileLinkAddress})
    info.EntityData.Leafs.Append("proxy-profile-linkaddress-from-ra-enable", types.YLeaf{"ProxyProfileLinkaddressFromRaEnable", info.ProxyProfileLinkaddressFromRaEnable})
    info.EntityData.Leafs.Append("profile-helper-address", types.YLeaf{"ProfileHelperAddress", info.ProfileHelperAddress})
    info.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", info.VrfName})
    info.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", info.InterfaceName})

    info.EntityData.YListKeys = []string {}

    return &(info.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceIdReferences
// Interface id references
type Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceIdReferences struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ipv6 dhcpv6d proxy iid reference. The type is slice of
    // Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceIdReferences_Ipv6Dhcpv6dProxyIidReference.
    Ipv6Dhcpv6dProxyIidReference []*Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceIdReferences_Ipv6Dhcpv6dProxyIidReference
}

func (interfaceIdReferences *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceIdReferences) GetEntityData() *types.CommonEntityData {
    interfaceIdReferences.EntityData.YFilter = interfaceIdReferences.YFilter
    interfaceIdReferences.EntityData.YangName = "interface-id-references"
    interfaceIdReferences.EntityData.BundleName = "cisco_ios_xr"
    interfaceIdReferences.EntityData.ParentYangName = "info"
    interfaceIdReferences.EntityData.SegmentPath = "interface-id-references"
    interfaceIdReferences.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/proxy/profiles/profile/info/" + interfaceIdReferences.EntityData.SegmentPath
    interfaceIdReferences.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceIdReferences.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceIdReferences.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceIdReferences.EntityData.Children = types.NewOrderedMap()
    interfaceIdReferences.EntityData.Children.Append("ipv6-dhcpv6d-proxy-iid-reference", types.YChild{"Ipv6Dhcpv6dProxyIidReference", nil})
    for i := range interfaceIdReferences.Ipv6Dhcpv6dProxyIidReference {
        types.SetYListKey(interfaceIdReferences.Ipv6Dhcpv6dProxyIidReference[i], i)
        interfaceIdReferences.EntityData.Children.Append(types.GetSegmentPath(interfaceIdReferences.Ipv6Dhcpv6dProxyIidReference[i]), types.YChild{"Ipv6Dhcpv6dProxyIidReference", interfaceIdReferences.Ipv6Dhcpv6dProxyIidReference[i]})
    }
    interfaceIdReferences.EntityData.Leafs = types.NewOrderedMap()

    interfaceIdReferences.EntityData.YListKeys = []string {}

    return &(interfaceIdReferences.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceIdReferences_Ipv6Dhcpv6dProxyIidReference
// ipv6 dhcpv6d proxy iid reference
type Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceIdReferences_Ipv6Dhcpv6dProxyIidReference struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Interface name for interface id. The type is string with length: 0..65.
    ProxyIidInterfaceName interface{}

    // Interface id. The type is string with length: 0..257.
    ProxyInterfaceId interface{}
}

func (ipv6Dhcpv6dProxyIidReference *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceIdReferences_Ipv6Dhcpv6dProxyIidReference) GetEntityData() *types.CommonEntityData {
    ipv6Dhcpv6dProxyIidReference.EntityData.YFilter = ipv6Dhcpv6dProxyIidReference.YFilter
    ipv6Dhcpv6dProxyIidReference.EntityData.YangName = "ipv6-dhcpv6d-proxy-iid-reference"
    ipv6Dhcpv6dProxyIidReference.EntityData.BundleName = "cisco_ios_xr"
    ipv6Dhcpv6dProxyIidReference.EntityData.ParentYangName = "interface-id-references"
    ipv6Dhcpv6dProxyIidReference.EntityData.SegmentPath = "ipv6-dhcpv6d-proxy-iid-reference" + types.AddNoKeyToken(ipv6Dhcpv6dProxyIidReference)
    ipv6Dhcpv6dProxyIidReference.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/proxy/profiles/profile/info/interface-id-references/" + ipv6Dhcpv6dProxyIidReference.EntityData.SegmentPath
    ipv6Dhcpv6dProxyIidReference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Dhcpv6dProxyIidReference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Dhcpv6dProxyIidReference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Dhcpv6dProxyIidReference.EntityData.Children = types.NewOrderedMap()
    ipv6Dhcpv6dProxyIidReference.EntityData.Leafs = types.NewOrderedMap()
    ipv6Dhcpv6dProxyIidReference.EntityData.Leafs.Append("proxy-iid-interface-name", types.YLeaf{"ProxyIidInterfaceName", ipv6Dhcpv6dProxyIidReference.ProxyIidInterfaceName})
    ipv6Dhcpv6dProxyIidReference.EntityData.Leafs.Append("proxy-interface-id", types.YLeaf{"ProxyInterfaceId", ipv6Dhcpv6dProxyIidReference.ProxyInterfaceId})

    ipv6Dhcpv6dProxyIidReference.EntityData.YListKeys = []string {}

    return &(ipv6Dhcpv6dProxyIidReference.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_VrfReferences
// VRF references
type Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_VrfReferences struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ipv6 dhcpv6d proxy vrf reference. The type is slice of
    // Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_VrfReferences_Ipv6Dhcpv6dProxyVrfReference.
    Ipv6Dhcpv6dProxyVrfReference []*Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_VrfReferences_Ipv6Dhcpv6dProxyVrfReference
}

func (vrfReferences *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_VrfReferences) GetEntityData() *types.CommonEntityData {
    vrfReferences.EntityData.YFilter = vrfReferences.YFilter
    vrfReferences.EntityData.YangName = "vrf-references"
    vrfReferences.EntityData.BundleName = "cisco_ios_xr"
    vrfReferences.EntityData.ParentYangName = "info"
    vrfReferences.EntityData.SegmentPath = "vrf-references"
    vrfReferences.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/proxy/profiles/profile/info/" + vrfReferences.EntityData.SegmentPath
    vrfReferences.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfReferences.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfReferences.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfReferences.EntityData.Children = types.NewOrderedMap()
    vrfReferences.EntityData.Children.Append("ipv6-dhcpv6d-proxy-vrf-reference", types.YChild{"Ipv6Dhcpv6dProxyVrfReference", nil})
    for i := range vrfReferences.Ipv6Dhcpv6dProxyVrfReference {
        types.SetYListKey(vrfReferences.Ipv6Dhcpv6dProxyVrfReference[i], i)
        vrfReferences.EntityData.Children.Append(types.GetSegmentPath(vrfReferences.Ipv6Dhcpv6dProxyVrfReference[i]), types.YChild{"Ipv6Dhcpv6dProxyVrfReference", vrfReferences.Ipv6Dhcpv6dProxyVrfReference[i]})
    }
    vrfReferences.EntityData.Leafs = types.NewOrderedMap()

    vrfReferences.EntityData.YListKeys = []string {}

    return &(vrfReferences.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_VrfReferences_Ipv6Dhcpv6dProxyVrfReference
// ipv6 dhcpv6d proxy vrf reference
type Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_VrfReferences_Ipv6Dhcpv6dProxyVrfReference struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // VRF name. The type is string with length: 0..33.
    ProxyReferenceVrfName interface{}
}

func (ipv6Dhcpv6dProxyVrfReference *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_VrfReferences_Ipv6Dhcpv6dProxyVrfReference) GetEntityData() *types.CommonEntityData {
    ipv6Dhcpv6dProxyVrfReference.EntityData.YFilter = ipv6Dhcpv6dProxyVrfReference.YFilter
    ipv6Dhcpv6dProxyVrfReference.EntityData.YangName = "ipv6-dhcpv6d-proxy-vrf-reference"
    ipv6Dhcpv6dProxyVrfReference.EntityData.BundleName = "cisco_ios_xr"
    ipv6Dhcpv6dProxyVrfReference.EntityData.ParentYangName = "vrf-references"
    ipv6Dhcpv6dProxyVrfReference.EntityData.SegmentPath = "ipv6-dhcpv6d-proxy-vrf-reference" + types.AddNoKeyToken(ipv6Dhcpv6dProxyVrfReference)
    ipv6Dhcpv6dProxyVrfReference.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/proxy/profiles/profile/info/vrf-references/" + ipv6Dhcpv6dProxyVrfReference.EntityData.SegmentPath
    ipv6Dhcpv6dProxyVrfReference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Dhcpv6dProxyVrfReference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Dhcpv6dProxyVrfReference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Dhcpv6dProxyVrfReference.EntityData.Children = types.NewOrderedMap()
    ipv6Dhcpv6dProxyVrfReference.EntityData.Leafs = types.NewOrderedMap()
    ipv6Dhcpv6dProxyVrfReference.EntityData.Leafs.Append("proxy-reference-vrf-name", types.YLeaf{"ProxyReferenceVrfName", ipv6Dhcpv6dProxyVrfReference.ProxyReferenceVrfName})

    ipv6Dhcpv6dProxyVrfReference.EntityData.YListKeys = []string {}

    return &(ipv6Dhcpv6dProxyVrfReference.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceReferences
// Interface references
type Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceReferences struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ipv6 dhcpv6d proxy interface reference. The type is slice of
    // Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceReferences_Ipv6Dhcpv6dProxyInterfaceReference.
    Ipv6Dhcpv6dProxyInterfaceReference []*Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceReferences_Ipv6Dhcpv6dProxyInterfaceReference
}

func (interfaceReferences *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceReferences) GetEntityData() *types.CommonEntityData {
    interfaceReferences.EntityData.YFilter = interfaceReferences.YFilter
    interfaceReferences.EntityData.YangName = "interface-references"
    interfaceReferences.EntityData.BundleName = "cisco_ios_xr"
    interfaceReferences.EntityData.ParentYangName = "info"
    interfaceReferences.EntityData.SegmentPath = "interface-references"
    interfaceReferences.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/proxy/profiles/profile/info/" + interfaceReferences.EntityData.SegmentPath
    interfaceReferences.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceReferences.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceReferences.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceReferences.EntityData.Children = types.NewOrderedMap()
    interfaceReferences.EntityData.Children.Append("ipv6-dhcpv6d-proxy-interface-reference", types.YChild{"Ipv6Dhcpv6dProxyInterfaceReference", nil})
    for i := range interfaceReferences.Ipv6Dhcpv6dProxyInterfaceReference {
        types.SetYListKey(interfaceReferences.Ipv6Dhcpv6dProxyInterfaceReference[i], i)
        interfaceReferences.EntityData.Children.Append(types.GetSegmentPath(interfaceReferences.Ipv6Dhcpv6dProxyInterfaceReference[i]), types.YChild{"Ipv6Dhcpv6dProxyInterfaceReference", interfaceReferences.Ipv6Dhcpv6dProxyInterfaceReference[i]})
    }
    interfaceReferences.EntityData.Leafs = types.NewOrderedMap()

    interfaceReferences.EntityData.YListKeys = []string {}

    return &(interfaceReferences.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceReferences_Ipv6Dhcpv6dProxyInterfaceReference
// ipv6 dhcpv6d proxy interface reference
type Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceReferences_Ipv6Dhcpv6dProxyInterfaceReference struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Interface name. The type is string with length: 0..65.
    ProxyReferenceInterfaceName interface{}
}

func (ipv6Dhcpv6dProxyInterfaceReference *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceReferences_Ipv6Dhcpv6dProxyInterfaceReference) GetEntityData() *types.CommonEntityData {
    ipv6Dhcpv6dProxyInterfaceReference.EntityData.YFilter = ipv6Dhcpv6dProxyInterfaceReference.YFilter
    ipv6Dhcpv6dProxyInterfaceReference.EntityData.YangName = "ipv6-dhcpv6d-proxy-interface-reference"
    ipv6Dhcpv6dProxyInterfaceReference.EntityData.BundleName = "cisco_ios_xr"
    ipv6Dhcpv6dProxyInterfaceReference.EntityData.ParentYangName = "interface-references"
    ipv6Dhcpv6dProxyInterfaceReference.EntityData.SegmentPath = "ipv6-dhcpv6d-proxy-interface-reference" + types.AddNoKeyToken(ipv6Dhcpv6dProxyInterfaceReference)
    ipv6Dhcpv6dProxyInterfaceReference.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/proxy/profiles/profile/info/interface-references/" + ipv6Dhcpv6dProxyInterfaceReference.EntityData.SegmentPath
    ipv6Dhcpv6dProxyInterfaceReference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Dhcpv6dProxyInterfaceReference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Dhcpv6dProxyInterfaceReference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Dhcpv6dProxyInterfaceReference.EntityData.Children = types.NewOrderedMap()
    ipv6Dhcpv6dProxyInterfaceReference.EntityData.Leafs = types.NewOrderedMap()
    ipv6Dhcpv6dProxyInterfaceReference.EntityData.Leafs.Append("proxy-reference-interface-name", types.YLeaf{"ProxyReferenceInterfaceName", ipv6Dhcpv6dProxyInterfaceReference.ProxyReferenceInterfaceName})

    ipv6Dhcpv6dProxyInterfaceReference.EntityData.YListKeys = []string {}

    return &(ipv6Dhcpv6dProxyInterfaceReference.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_DisconnectHistories
// DHCPv6 proxy disconnect history
type Dhcpv6_Nodes_Node_Proxy_DisconnectHistories struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Single DHCPv6 proxy disconnect history. The type is slice of
    // Dhcpv6_Nodes_Node_Proxy_DisconnectHistories_DisconnectHistory.
    DisconnectHistory []*Dhcpv6_Nodes_Node_Proxy_DisconnectHistories_DisconnectHistory
}

func (disconnectHistories *Dhcpv6_Nodes_Node_Proxy_DisconnectHistories) GetEntityData() *types.CommonEntityData {
    disconnectHistories.EntityData.YFilter = disconnectHistories.YFilter
    disconnectHistories.EntityData.YangName = "disconnect-histories"
    disconnectHistories.EntityData.BundleName = "cisco_ios_xr"
    disconnectHistories.EntityData.ParentYangName = "proxy"
    disconnectHistories.EntityData.SegmentPath = "disconnect-histories"
    disconnectHistories.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/proxy/" + disconnectHistories.EntityData.SegmentPath
    disconnectHistories.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    disconnectHistories.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    disconnectHistories.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    disconnectHistories.EntityData.Children = types.NewOrderedMap()
    disconnectHistories.EntityData.Children.Append("disconnect-history", types.YChild{"DisconnectHistory", nil})
    for i := range disconnectHistories.DisconnectHistory {
        disconnectHistories.EntityData.Children.Append(types.GetSegmentPath(disconnectHistories.DisconnectHistory[i]), types.YChild{"DisconnectHistory", disconnectHistories.DisconnectHistory[i]})
    }
    disconnectHistories.EntityData.Leafs = types.NewOrderedMap()

    disconnectHistories.EntityData.YListKeys = []string {}

    return &(disconnectHistories.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_DisconnectHistories_DisconnectHistory
// Single DHCPv6 proxy disconnect history
type Dhcpv6_Nodes_Node_Proxy_DisconnectHistories_DisconnectHistory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Index. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Index interface{}

    // session start time epoch. The type is interface{} with range:
    // 0..18446744073709551615.
    SessionStartTimeEpoch interface{}

    // session end time epoch. The type is interface{} with range:
    // 0..18446744073709551615.
    SessionEndTimeEpoch interface{}

    // DiscReason. The type is string with length: 0..256.
    DiscReason interface{}

    // sub label. The type is interface{} with range: 0..4294967295.
    SubLabel interface{}

    // Client DUID. The type is string with length: 0..131.
    Duid interface{}

    // IAType. The type is string with length: 0..6.
    IaType interface{}

    // ia id. The type is interface{} with range: 0..4294967295.
    IaId interface{}

    // MACAddress. The type is string with length: 0..17.
    MacAddress interface{}
}

func (disconnectHistory *Dhcpv6_Nodes_Node_Proxy_DisconnectHistories_DisconnectHistory) GetEntityData() *types.CommonEntityData {
    disconnectHistory.EntityData.YFilter = disconnectHistory.YFilter
    disconnectHistory.EntityData.YangName = "disconnect-history"
    disconnectHistory.EntityData.BundleName = "cisco_ios_xr"
    disconnectHistory.EntityData.ParentYangName = "disconnect-histories"
    disconnectHistory.EntityData.SegmentPath = "disconnect-history" + types.AddKeyToken(disconnectHistory.Index, "index")
    disconnectHistory.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/proxy/disconnect-histories/" + disconnectHistory.EntityData.SegmentPath
    disconnectHistory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    disconnectHistory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    disconnectHistory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    disconnectHistory.EntityData.Children = types.NewOrderedMap()
    disconnectHistory.EntityData.Leafs = types.NewOrderedMap()
    disconnectHistory.EntityData.Leafs.Append("index", types.YLeaf{"Index", disconnectHistory.Index})
    disconnectHistory.EntityData.Leafs.Append("session-start-time-epoch", types.YLeaf{"SessionStartTimeEpoch", disconnectHistory.SessionStartTimeEpoch})
    disconnectHistory.EntityData.Leafs.Append("session-end-time-epoch", types.YLeaf{"SessionEndTimeEpoch", disconnectHistory.SessionEndTimeEpoch})
    disconnectHistory.EntityData.Leafs.Append("disc-reason", types.YLeaf{"DiscReason", disconnectHistory.DiscReason})
    disconnectHistory.EntityData.Leafs.Append("sub-label", types.YLeaf{"SubLabel", disconnectHistory.SubLabel})
    disconnectHistory.EntityData.Leafs.Append("duid", types.YLeaf{"Duid", disconnectHistory.Duid})
    disconnectHistory.EntityData.Leafs.Append("ia-type", types.YLeaf{"IaType", disconnectHistory.IaType})
    disconnectHistory.EntityData.Leafs.Append("ia-id", types.YLeaf{"IaId", disconnectHistory.IaId})
    disconnectHistory.EntityData.Leafs.Append("mac-address", types.YLeaf{"MacAddress", disconnectHistory.MacAddress})

    disconnectHistory.EntityData.YListKeys = []string {"Index"}

    return &(disconnectHistory.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Interfaces
// DHCPV6 proxy interface
type Dhcpv6_Nodes_Node_Proxy_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 DHCP proxy interface. The type is slice of
    // Dhcpv6_Nodes_Node_Proxy_Interfaces_Interface.
    Interface []*Dhcpv6_Nodes_Node_Proxy_Interfaces_Interface
}

func (interfaces *Dhcpv6_Nodes_Node_Proxy_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "proxy"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/proxy/" + interfaces.EntityData.SegmentPath
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

// Dhcpv6_Nodes_Node_Proxy_Interfaces_Interface
// IPv6 DHCP proxy interface
type Dhcpv6_Nodes_Node_Proxy_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    InterfaceName interface{}

    // VRF name. The type is string with length: 0..33.
    ProxyVrfName interface{}

    // Mode of interface. The type is BagDhcpv6dSubMode.
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

    // DHCPv6 Interface SRG role. The type is BagDhcpv6dIntfSrgRole.
    SrgRole interface{}

    // DHCPv6 Interface SERG role. The type is BagDhcpv6dIntfSergRole.
    SergRole interface{}

    // Mac Throttle Status. The type is bool.
    MacThrottle interface{}

    // SRG VRF name. The type is string with length: 0..33.
    SrgVrfName interface{}

    // SRG P2P Status. The type is bool.
    Srgp2p interface{}
}

func (self *Dhcpv6_Nodes_Node_Proxy_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceName, "interface-name")
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/proxy/interfaces/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})
    self.EntityData.Leafs.Append("proxy-vrf-name", types.YLeaf{"ProxyVrfName", self.ProxyVrfName})
    self.EntityData.Leafs.Append("proxy-interface-mode", types.YLeaf{"ProxyInterfaceMode", self.ProxyInterfaceMode})
    self.EntityData.Leafs.Append("is-proxy-interface-ambiguous", types.YLeaf{"IsProxyInterfaceAmbiguous", self.IsProxyInterfaceAmbiguous})
    self.EntityData.Leafs.Append("proxy-interface-profile-name", types.YLeaf{"ProxyInterfaceProfileName", self.ProxyInterfaceProfileName})
    self.EntityData.Leafs.Append("proxy-interface-lease-limit-type", types.YLeaf{"ProxyInterfaceLeaseLimitType", self.ProxyInterfaceLeaseLimitType})
    self.EntityData.Leafs.Append("proxy-interface-lease-limits", types.YLeaf{"ProxyInterfaceLeaseLimits", self.ProxyInterfaceLeaseLimits})
    self.EntityData.Leafs.Append("srg-role", types.YLeaf{"SrgRole", self.SrgRole})
    self.EntityData.Leafs.Append("serg-role", types.YLeaf{"SergRole", self.SergRole})
    self.EntityData.Leafs.Append("mac-throttle", types.YLeaf{"MacThrottle", self.MacThrottle})
    self.EntityData.Leafs.Append("srg-vrf-name", types.YLeaf{"SrgVrfName", self.SrgVrfName})
    self.EntityData.Leafs.Append("srgp2p", types.YLeaf{"Srgp2p", self.Srgp2p})

    self.EntityData.YListKeys = []string {"InterfaceName"}

    return &(self.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Statistics
// DHCPv6 proxy statistics
type Dhcpv6_Nodes_Node_Proxy_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ipv6 dhcpv6d proxy stat. The type is slice of
    // Dhcpv6_Nodes_Node_Proxy_Statistics_Ipv6Dhcpv6dProxyStat.
    Ipv6Dhcpv6dProxyStat []*Dhcpv6_Nodes_Node_Proxy_Statistics_Ipv6Dhcpv6dProxyStat
}

func (statistics *Dhcpv6_Nodes_Node_Proxy_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "proxy"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/proxy/" + statistics.EntityData.SegmentPath
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = types.NewOrderedMap()
    statistics.EntityData.Children.Append("ipv6-dhcpv6d-proxy-stat", types.YChild{"Ipv6Dhcpv6dProxyStat", nil})
    for i := range statistics.Ipv6Dhcpv6dProxyStat {
        types.SetYListKey(statistics.Ipv6Dhcpv6dProxyStat[i], i)
        statistics.EntityData.Children.Append(types.GetSegmentPath(statistics.Ipv6Dhcpv6dProxyStat[i]), types.YChild{"Ipv6Dhcpv6dProxyStat", statistics.Ipv6Dhcpv6dProxyStat[i]})
    }
    statistics.EntityData.Leafs = types.NewOrderedMap()

    statistics.EntityData.YListKeys = []string {}

    return &(statistics.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Statistics_Ipv6Dhcpv6dProxyStat
// ipv6 dhcpv6d proxy stat
type Dhcpv6_Nodes_Node_Proxy_Statistics_Ipv6Dhcpv6dProxyStat struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // DHCPv6 L3 VRF name. The type is string with length: 0..33.
    VrfName interface{}

    // Proxy statistics.
    Statistics Dhcpv6_Nodes_Node_Proxy_Statistics_Ipv6Dhcpv6dProxyStat_Statistics
}

func (ipv6Dhcpv6dProxyStat *Dhcpv6_Nodes_Node_Proxy_Statistics_Ipv6Dhcpv6dProxyStat) GetEntityData() *types.CommonEntityData {
    ipv6Dhcpv6dProxyStat.EntityData.YFilter = ipv6Dhcpv6dProxyStat.YFilter
    ipv6Dhcpv6dProxyStat.EntityData.YangName = "ipv6-dhcpv6d-proxy-stat"
    ipv6Dhcpv6dProxyStat.EntityData.BundleName = "cisco_ios_xr"
    ipv6Dhcpv6dProxyStat.EntityData.ParentYangName = "statistics"
    ipv6Dhcpv6dProxyStat.EntityData.SegmentPath = "ipv6-dhcpv6d-proxy-stat" + types.AddNoKeyToken(ipv6Dhcpv6dProxyStat)
    ipv6Dhcpv6dProxyStat.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/proxy/statistics/" + ipv6Dhcpv6dProxyStat.EntityData.SegmentPath
    ipv6Dhcpv6dProxyStat.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Dhcpv6dProxyStat.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Dhcpv6dProxyStat.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Dhcpv6dProxyStat.EntityData.Children = types.NewOrderedMap()
    ipv6Dhcpv6dProxyStat.EntityData.Children.Append("statistics", types.YChild{"Statistics", &ipv6Dhcpv6dProxyStat.Statistics})
    ipv6Dhcpv6dProxyStat.EntityData.Leafs = types.NewOrderedMap()
    ipv6Dhcpv6dProxyStat.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", ipv6Dhcpv6dProxyStat.VrfName})

    ipv6Dhcpv6dProxyStat.EntityData.YListKeys = []string {}

    return &(ipv6Dhcpv6dProxyStat.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Statistics_Ipv6Dhcpv6dProxyStat_Statistics
// Proxy statistics
type Dhcpv6_Nodes_Node_Proxy_Statistics_Ipv6Dhcpv6dProxyStat_Statistics struct {
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

func (statistics *Dhcpv6_Nodes_Node_Proxy_Statistics_Ipv6Dhcpv6dProxyStat_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "ipv6-dhcpv6d-proxy-stat"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/proxy/statistics/ipv6-dhcpv6d-proxy-stat/" + statistics.EntityData.SegmentPath
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = types.NewOrderedMap()
    statistics.EntityData.Leafs = types.NewOrderedMap()
    statistics.EntityData.Leafs.Append("received-packets", types.YLeaf{"ReceivedPackets", statistics.ReceivedPackets})
    statistics.EntityData.Leafs.Append("transmitted-packets", types.YLeaf{"TransmittedPackets", statistics.TransmittedPackets})
    statistics.EntityData.Leafs.Append("dropped-packets", types.YLeaf{"DroppedPackets", statistics.DroppedPackets})

    statistics.EntityData.YListKeys = []string {}

    return &(statistics.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Binding
// DHCPV6 proxy bindings
type Dhcpv6_Nodes_Node_Proxy_Binding struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DHCPV6 proxy client bindings.
    Clients Dhcpv6_Nodes_Node_Proxy_Binding_Clients

    // DHCPV6 proxy binding summary.
    Summary Dhcpv6_Nodes_Node_Proxy_Binding_Summary
}

func (binding *Dhcpv6_Nodes_Node_Proxy_Binding) GetEntityData() *types.CommonEntityData {
    binding.EntityData.YFilter = binding.YFilter
    binding.EntityData.YangName = "binding"
    binding.EntityData.BundleName = "cisco_ios_xr"
    binding.EntityData.ParentYangName = "proxy"
    binding.EntityData.SegmentPath = "binding"
    binding.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/proxy/" + binding.EntityData.SegmentPath
    binding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    binding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    binding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    binding.EntityData.Children = types.NewOrderedMap()
    binding.EntityData.Children.Append("clients", types.YChild{"Clients", &binding.Clients})
    binding.EntityData.Children.Append("summary", types.YChild{"Summary", &binding.Summary})
    binding.EntityData.Leafs = types.NewOrderedMap()

    binding.EntityData.YListKeys = []string {}

    return &(binding.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Binding_Clients
// DHCPV6 proxy client bindings
type Dhcpv6_Nodes_Node_Proxy_Binding_Clients struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Single DHCPV6 proxy binding. The type is slice of
    // Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client.
    Client []*Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client
}

func (clients *Dhcpv6_Nodes_Node_Proxy_Binding_Clients) GetEntityData() *types.CommonEntityData {
    clients.EntityData.YFilter = clients.YFilter
    clients.EntityData.YangName = "clients"
    clients.EntityData.BundleName = "cisco_ios_xr"
    clients.EntityData.ParentYangName = "binding"
    clients.EntityData.SegmentPath = "clients"
    clients.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/proxy/binding/" + clients.EntityData.SegmentPath
    clients.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clients.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clients.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clients.EntityData.Children = types.NewOrderedMap()
    clients.EntityData.Children.Append("client", types.YChild{"Client", nil})
    for i := range clients.Client {
        clients.EntityData.Children.Append(types.GetSegmentPath(clients.Client[i]), types.YChild{"Client", clients.Client[i]})
    }
    clients.EntityData.Leafs = types.NewOrderedMap()

    clients.EntityData.YListKeys = []string {}

    return &(clients.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client
// Single DHCPV6 proxy binding
type Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    // [a-zA-Z0-9._/-]+.
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
    Srgp2p interface{}

    // DHCPV6 SRG VRF NAME. The type is string with length: 0..33.
    SrgVrfName interface{}

    // DHCPV6 SERG state. The type is interface{} with range: 0..4294967295.
    SergState interface{}

    // DHCPV6 SERG Intf Role. The type is interface{} with range: 0..4294967295.
    SergIntfRole interface{}

    // List of DHCPv6 IA_ID/PDs.
    IaIdPd Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd
}

func (client *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client) GetEntityData() *types.CommonEntityData {
    client.EntityData.YFilter = client.YFilter
    client.EntityData.YangName = "client"
    client.EntityData.BundleName = "cisco_ios_xr"
    client.EntityData.ParentYangName = "clients"
    client.EntityData.SegmentPath = "client" + types.AddKeyToken(client.ClientId, "client-id")
    client.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/proxy/binding/clients/" + client.EntityData.SegmentPath
    client.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    client.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    client.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    client.EntityData.Children = types.NewOrderedMap()
    client.EntityData.Children.Append("ia-id-pd", types.YChild{"IaIdPd", &client.IaIdPd})
    client.EntityData.Leafs = types.NewOrderedMap()
    client.EntityData.Leafs.Append("client-id", types.YLeaf{"ClientId", client.ClientId})
    client.EntityData.Leafs.Append("duid", types.YLeaf{"Duid", client.Duid})
    client.EntityData.Leafs.Append("client-flag", types.YLeaf{"ClientFlag", client.ClientFlag})
    client.EntityData.Leafs.Append("subscriber-label", types.YLeaf{"SubscriberLabel", client.SubscriberLabel})
    client.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", client.VrfName})
    client.EntityData.Leafs.Append("mac-address", types.YLeaf{"MacAddress", client.MacAddress})
    client.EntityData.Leafs.Append("ia-id-p-ds", types.YLeaf{"IaIdPDs", client.IaIdPDs})
    client.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", client.InterfaceName})
    client.EntityData.Leafs.Append("access-vrf-name", types.YLeaf{"AccessVrfName", client.AccessVrfName})
    client.EntityData.Leafs.Append("proxy-binding-tags", types.YLeaf{"ProxyBindingTags", client.ProxyBindingTags})
    client.EntityData.Leafs.Append("proxy-binding-outer-tag", types.YLeaf{"ProxyBindingOuterTag", client.ProxyBindingOuterTag})
    client.EntityData.Leafs.Append("proxy-binding-inner-tag", types.YLeaf{"ProxyBindingInnerTag", client.ProxyBindingInnerTag})
    client.EntityData.Leafs.Append("class-name", types.YLeaf{"ClassName", client.ClassName})
    client.EntityData.Leafs.Append("pool-name", types.YLeaf{"PoolName", client.PoolName})
    client.EntityData.Leafs.Append("rx-remote-id", types.YLeaf{"RxRemoteId", client.RxRemoteId})
    client.EntityData.Leafs.Append("tx-remote-id", types.YLeaf{"TxRemoteId", client.TxRemoteId})
    client.EntityData.Leafs.Append("rx-interface-id", types.YLeaf{"RxInterfaceId", client.RxInterfaceId})
    client.EntityData.Leafs.Append("tx-interface-id", types.YLeaf{"TxInterfaceId", client.TxInterfaceId})
    client.EntityData.Leafs.Append("server-ipv6-address", types.YLeaf{"ServerIpv6Address", client.ServerIpv6Address})
    client.EntityData.Leafs.Append("profile-name", types.YLeaf{"ProfileName", client.ProfileName})
    client.EntityData.Leafs.Append("framed-ipv6-prefix", types.YLeaf{"FramedIpv6Prefix", client.FramedIpv6Prefix})
    client.EntityData.Leafs.Append("framed-prefix-length", types.YLeaf{"FramedPrefixLength", client.FramedPrefixLength})
    client.EntityData.Leafs.Append("is-nak-next-renew", types.YLeaf{"IsNakNextRenew", client.IsNakNextRenew})
    client.EntityData.Leafs.Append("srg-state", types.YLeaf{"SrgState", client.SrgState})
    client.EntityData.Leafs.Append("srg-intf-role", types.YLeaf{"SrgIntfRole", client.SrgIntfRole})
    client.EntityData.Leafs.Append("srgp2p", types.YLeaf{"Srgp2p", client.Srgp2p})
    client.EntityData.Leafs.Append("srg-vrf-name", types.YLeaf{"SrgVrfName", client.SrgVrfName})
    client.EntityData.Leafs.Append("serg-state", types.YLeaf{"SergState", client.SergState})
    client.EntityData.Leafs.Append("serg-intf-role", types.YLeaf{"SergIntfRole", client.SergIntfRole})

    client.EntityData.YListKeys = []string {"ClientId"}

    return &(client.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd
// List of DHCPv6 IA_ID/PDs
type Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bag dhcpv6d ia id pd info. The type is slice of
    // Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6dIaIdPdInfo.
    BagDhcpv6dIaIdPdInfo []*Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6dIaIdPdInfo
}

func (iaIdPd *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd) GetEntityData() *types.CommonEntityData {
    iaIdPd.EntityData.YFilter = iaIdPd.YFilter
    iaIdPd.EntityData.YangName = "ia-id-pd"
    iaIdPd.EntityData.BundleName = "cisco_ios_xr"
    iaIdPd.EntityData.ParentYangName = "client"
    iaIdPd.EntityData.SegmentPath = "ia-id-pd"
    iaIdPd.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/proxy/binding/clients/client/" + iaIdPd.EntityData.SegmentPath
    iaIdPd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    iaIdPd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    iaIdPd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    iaIdPd.EntityData.Children = types.NewOrderedMap()
    iaIdPd.EntityData.Children.Append("bag-dhcpv6d-ia-id-pd-info", types.YChild{"BagDhcpv6dIaIdPdInfo", nil})
    for i := range iaIdPd.BagDhcpv6dIaIdPdInfo {
        types.SetYListKey(iaIdPd.BagDhcpv6dIaIdPdInfo[i], i)
        iaIdPd.EntityData.Children.Append(types.GetSegmentPath(iaIdPd.BagDhcpv6dIaIdPdInfo[i]), types.YChild{"BagDhcpv6dIaIdPdInfo", iaIdPd.BagDhcpv6dIaIdPdInfo[i]})
    }
    iaIdPd.EntityData.Leafs = types.NewOrderedMap()

    iaIdPd.EntityData.YListKeys = []string {}

    return &(iaIdPd.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6dIaIdPdInfo
// bag dhcpv6d ia id pd info
type Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6dIaIdPdInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // IA type. The type is BagDhcpv6dIaId.
    IaType interface{}

    // IA_ID of this IA. The type is interface{} with range: 0..4294967295.
    IaId interface{}

    // FSM Flag for this IA. The type is interface{} with range: 0..4294967295.
    Flags interface{}

    // Total address in this IA. The type is interface{} with range: 0..65535.
    TotalAddress interface{}

    // State. The type is BagDhcpv6dFsmState.
    State interface{}

    // List of addresses in this IA.
    Addresses Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6dIaIdPdInfo_Addresses
}

func (bagDhcpv6dIaIdPdInfo *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6dIaIdPdInfo) GetEntityData() *types.CommonEntityData {
    bagDhcpv6dIaIdPdInfo.EntityData.YFilter = bagDhcpv6dIaIdPdInfo.YFilter
    bagDhcpv6dIaIdPdInfo.EntityData.YangName = "bag-dhcpv6d-ia-id-pd-info"
    bagDhcpv6dIaIdPdInfo.EntityData.BundleName = "cisco_ios_xr"
    bagDhcpv6dIaIdPdInfo.EntityData.ParentYangName = "ia-id-pd"
    bagDhcpv6dIaIdPdInfo.EntityData.SegmentPath = "bag-dhcpv6d-ia-id-pd-info" + types.AddNoKeyToken(bagDhcpv6dIaIdPdInfo)
    bagDhcpv6dIaIdPdInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/proxy/binding/clients/client/ia-id-pd/" + bagDhcpv6dIaIdPdInfo.EntityData.SegmentPath
    bagDhcpv6dIaIdPdInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bagDhcpv6dIaIdPdInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bagDhcpv6dIaIdPdInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bagDhcpv6dIaIdPdInfo.EntityData.Children = types.NewOrderedMap()
    bagDhcpv6dIaIdPdInfo.EntityData.Children.Append("addresses", types.YChild{"Addresses", &bagDhcpv6dIaIdPdInfo.Addresses})
    bagDhcpv6dIaIdPdInfo.EntityData.Leafs = types.NewOrderedMap()
    bagDhcpv6dIaIdPdInfo.EntityData.Leafs.Append("ia-type", types.YLeaf{"IaType", bagDhcpv6dIaIdPdInfo.IaType})
    bagDhcpv6dIaIdPdInfo.EntityData.Leafs.Append("ia-id", types.YLeaf{"IaId", bagDhcpv6dIaIdPdInfo.IaId})
    bagDhcpv6dIaIdPdInfo.EntityData.Leafs.Append("flags", types.YLeaf{"Flags", bagDhcpv6dIaIdPdInfo.Flags})
    bagDhcpv6dIaIdPdInfo.EntityData.Leafs.Append("total-address", types.YLeaf{"TotalAddress", bagDhcpv6dIaIdPdInfo.TotalAddress})
    bagDhcpv6dIaIdPdInfo.EntityData.Leafs.Append("state", types.YLeaf{"State", bagDhcpv6dIaIdPdInfo.State})

    bagDhcpv6dIaIdPdInfo.EntityData.YListKeys = []string {}

    return &(bagDhcpv6dIaIdPdInfo.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6dIaIdPdInfo_Addresses
// List of addresses in this IA
type Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6dIaIdPdInfo_Addresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bag dhcpv6d addr attrb. The type is slice of
    // Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6dIaIdPdInfo_Addresses_BagDhcpv6dAddrAttrb.
    BagDhcpv6dAddrAttrb []*Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6dIaIdPdInfo_Addresses_BagDhcpv6dAddrAttrb
}

func (addresses *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6dIaIdPdInfo_Addresses) GetEntityData() *types.CommonEntityData {
    addresses.EntityData.YFilter = addresses.YFilter
    addresses.EntityData.YangName = "addresses"
    addresses.EntityData.BundleName = "cisco_ios_xr"
    addresses.EntityData.ParentYangName = "bag-dhcpv6d-ia-id-pd-info"
    addresses.EntityData.SegmentPath = "addresses"
    addresses.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/proxy/binding/clients/client/ia-id-pd/bag-dhcpv6d-ia-id-pd-info/" + addresses.EntityData.SegmentPath
    addresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addresses.EntityData.Children = types.NewOrderedMap()
    addresses.EntityData.Children.Append("bag-dhcpv6d-addr-attrb", types.YChild{"BagDhcpv6dAddrAttrb", nil})
    for i := range addresses.BagDhcpv6dAddrAttrb {
        types.SetYListKey(addresses.BagDhcpv6dAddrAttrb[i], i)
        addresses.EntityData.Children.Append(types.GetSegmentPath(addresses.BagDhcpv6dAddrAttrb[i]), types.YChild{"BagDhcpv6dAddrAttrb", addresses.BagDhcpv6dAddrAttrb[i]})
    }
    addresses.EntityData.Leafs = types.NewOrderedMap()

    addresses.EntityData.YListKeys = []string {}

    return &(addresses.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6dIaIdPdInfo_Addresses_BagDhcpv6dAddrAttrb
// bag dhcpv6d addr attrb
type Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6dIaIdPdInfo_Addresses_BagDhcpv6dAddrAttrb struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (bagDhcpv6dAddrAttrb *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6dIaIdPdInfo_Addresses_BagDhcpv6dAddrAttrb) GetEntityData() *types.CommonEntityData {
    bagDhcpv6dAddrAttrb.EntityData.YFilter = bagDhcpv6dAddrAttrb.YFilter
    bagDhcpv6dAddrAttrb.EntityData.YangName = "bag-dhcpv6d-addr-attrb"
    bagDhcpv6dAddrAttrb.EntityData.BundleName = "cisco_ios_xr"
    bagDhcpv6dAddrAttrb.EntityData.ParentYangName = "addresses"
    bagDhcpv6dAddrAttrb.EntityData.SegmentPath = "bag-dhcpv6d-addr-attrb" + types.AddNoKeyToken(bagDhcpv6dAddrAttrb)
    bagDhcpv6dAddrAttrb.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/proxy/binding/clients/client/ia-id-pd/bag-dhcpv6d-ia-id-pd-info/addresses/" + bagDhcpv6dAddrAttrb.EntityData.SegmentPath
    bagDhcpv6dAddrAttrb.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bagDhcpv6dAddrAttrb.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bagDhcpv6dAddrAttrb.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bagDhcpv6dAddrAttrb.EntityData.Children = types.NewOrderedMap()
    bagDhcpv6dAddrAttrb.EntityData.Leafs = types.NewOrderedMap()
    bagDhcpv6dAddrAttrb.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", bagDhcpv6dAddrAttrb.Prefix})
    bagDhcpv6dAddrAttrb.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", bagDhcpv6dAddrAttrb.PrefixLength})
    bagDhcpv6dAddrAttrb.EntityData.Leafs.Append("lease-time", types.YLeaf{"LeaseTime", bagDhcpv6dAddrAttrb.LeaseTime})
    bagDhcpv6dAddrAttrb.EntityData.Leafs.Append("remaining-lease-time", types.YLeaf{"RemainingLeaseTime", bagDhcpv6dAddrAttrb.RemainingLeaseTime})

    bagDhcpv6dAddrAttrb.EntityData.YListKeys = []string {}

    return &(bagDhcpv6dAddrAttrb.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Binding_Summary
// DHCPV6 proxy binding summary
type Dhcpv6_Nodes_Node_Proxy_Binding_Summary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Total number of clients. The type is interface{} with range: 0..4294967295.
    Clients interface{}

    // IANA proxy binding summary.
    Iana Dhcpv6_Nodes_Node_Proxy_Binding_Summary_Iana

    // IAPD proxy binding summary.
    Iapd Dhcpv6_Nodes_Node_Proxy_Binding_Summary_Iapd
}

func (summary *Dhcpv6_Nodes_Node_Proxy_Binding_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "binding"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/proxy/binding/" + summary.EntityData.SegmentPath
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = types.NewOrderedMap()
    summary.EntityData.Children.Append("iana", types.YChild{"Iana", &summary.Iana})
    summary.EntityData.Children.Append("iapd", types.YChild{"Iapd", &summary.Iapd})
    summary.EntityData.Leafs = types.NewOrderedMap()
    summary.EntityData.Leafs.Append("clients", types.YLeaf{"Clients", summary.Clients})

    summary.EntityData.YListKeys = []string {}

    return &(summary.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Binding_Summary_Iana
// IANA proxy binding summary
type Dhcpv6_Nodes_Node_Proxy_Binding_Summary_Iana struct {
    EntityData types.CommonEntityData
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

func (iana *Dhcpv6_Nodes_Node_Proxy_Binding_Summary_Iana) GetEntityData() *types.CommonEntityData {
    iana.EntityData.YFilter = iana.YFilter
    iana.EntityData.YangName = "iana"
    iana.EntityData.BundleName = "cisco_ios_xr"
    iana.EntityData.ParentYangName = "summary"
    iana.EntityData.SegmentPath = "iana"
    iana.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/proxy/binding/summary/" + iana.EntityData.SegmentPath
    iana.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    iana.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    iana.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    iana.EntityData.Children = types.NewOrderedMap()
    iana.EntityData.Leafs = types.NewOrderedMap()
    iana.EntityData.Leafs.Append("initializing-clients", types.YLeaf{"InitializingClients", iana.InitializingClients})
    iana.EntityData.Leafs.Append("dpm-waiting-clients", types.YLeaf{"DpmWaitingClients", iana.DpmWaitingClients})
    iana.EntityData.Leafs.Append("daps-waiting-clients", types.YLeaf{"DapsWaitingClients", iana.DapsWaitingClients})
    iana.EntityData.Leafs.Append("msg-waiting-clients", types.YLeaf{"MsgWaitingClients", iana.MsgWaitingClients})
    iana.EntityData.Leafs.Append("iedge-waiting-clients", types.YLeaf{"IedgeWaitingClients", iana.IedgeWaitingClients})
    iana.EntityData.Leafs.Append("rib-waiting-clients", types.YLeaf{"RibWaitingClients", iana.RibWaitingClients})
    iana.EntityData.Leafs.Append("bound-clients", types.YLeaf{"BoundClients", iana.BoundClients})

    iana.EntityData.YListKeys = []string {}

    return &(iana.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Binding_Summary_Iapd
// IAPD proxy binding summary
type Dhcpv6_Nodes_Node_Proxy_Binding_Summary_Iapd struct {
    EntityData types.CommonEntityData
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

func (iapd *Dhcpv6_Nodes_Node_Proxy_Binding_Summary_Iapd) GetEntityData() *types.CommonEntityData {
    iapd.EntityData.YFilter = iapd.YFilter
    iapd.EntityData.YangName = "iapd"
    iapd.EntityData.BundleName = "cisco_ios_xr"
    iapd.EntityData.ParentYangName = "summary"
    iapd.EntityData.SegmentPath = "iapd"
    iapd.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/proxy/binding/summary/" + iapd.EntityData.SegmentPath
    iapd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    iapd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    iapd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    iapd.EntityData.Children = types.NewOrderedMap()
    iapd.EntityData.Leafs = types.NewOrderedMap()
    iapd.EntityData.Leafs.Append("initializing-clients", types.YLeaf{"InitializingClients", iapd.InitializingClients})
    iapd.EntityData.Leafs.Append("dpm-waiting-clients", types.YLeaf{"DpmWaitingClients", iapd.DpmWaitingClients})
    iapd.EntityData.Leafs.Append("daps-waiting-clients", types.YLeaf{"DapsWaitingClients", iapd.DapsWaitingClients})
    iapd.EntityData.Leafs.Append("msg-waiting-clients", types.YLeaf{"MsgWaitingClients", iapd.MsgWaitingClients})
    iapd.EntityData.Leafs.Append("iedge-waiting-clients", types.YLeaf{"IedgeWaitingClients", iapd.IedgeWaitingClients})
    iapd.EntityData.Leafs.Append("rib-waiting-clients", types.YLeaf{"RibWaitingClients", iapd.RibWaitingClients})
    iapd.EntityData.Leafs.Append("bound-clients", types.YLeaf{"BoundClients", iapd.BoundClients})

    iapd.EntityData.YListKeys = []string {}

    return &(iapd.EntityData)
}

// Dhcpv6_Nodes_Node_Base
// IPv6 DHCP Base
type Dhcpv6_Nodes_Node_Base struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DHCP database.
    Database Dhcpv6_Nodes_Node_Base_Database

    // IPv6 DHCP Base Binding.
    AddrBindings Dhcpv6_Nodes_Node_Base_AddrBindings
}

func (base *Dhcpv6_Nodes_Node_Base) GetEntityData() *types.CommonEntityData {
    base.EntityData.YFilter = base.YFilter
    base.EntityData.YangName = "base"
    base.EntityData.BundleName = "cisco_ios_xr"
    base.EntityData.ParentYangName = "node"
    base.EntityData.SegmentPath = "base"
    base.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/" + base.EntityData.SegmentPath
    base.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    base.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    base.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    base.EntityData.Children = types.NewOrderedMap()
    base.EntityData.Children.Append("database", types.YChild{"Database", &base.Database})
    base.EntityData.Children.Append("addr-bindings", types.YChild{"AddrBindings", &base.AddrBindings})
    base.EntityData.Leafs = types.NewOrderedMap()

    base.EntityData.YListKeys = []string {}

    return &(base.EntityData)
}

// Dhcpv6_Nodes_Node_Base_Database
// DHCP database
type Dhcpv6_Nodes_Node_Base_Database struct {
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

func (database *Dhcpv6_Nodes_Node_Base_Database) GetEntityData() *types.CommonEntityData {
    database.EntityData.YFilter = database.YFilter
    database.EntityData.YangName = "database"
    database.EntityData.BundleName = "cisco_ios_xr"
    database.EntityData.ParentYangName = "base"
    database.EntityData.SegmentPath = "database"
    database.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/base/" + database.EntityData.SegmentPath
    database.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    database.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    database.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    database.EntityData.Children = types.NewOrderedMap()
    database.EntityData.Leafs = types.NewOrderedMap()
    database.EntityData.Leafs.Append("configured", types.YLeaf{"Configured", database.Configured})
    database.EntityData.Leafs.Append("version", types.YLeaf{"Version", database.Version})
    database.EntityData.Leafs.Append("full-file-write-interval", types.YLeaf{"FullFileWriteInterval", database.FullFileWriteInterval})
    database.EntityData.Leafs.Append("last-full-write-file-name", types.YLeaf{"LastFullWriteFileName", database.LastFullWriteFileName})
    database.EntityData.Leafs.Append("last-full-write-time", types.YLeaf{"LastFullWriteTime", database.LastFullWriteTime})
    database.EntityData.Leafs.Append("full-file-write-count", types.YLeaf{"FullFileWriteCount", database.FullFileWriteCount})
    database.EntityData.Leafs.Append("failed-full-file-write-count", types.YLeaf{"FailedFullFileWriteCount", database.FailedFullFileWriteCount})
    database.EntityData.Leafs.Append("full-file-record-count", types.YLeaf{"FullFileRecordCount", database.FullFileRecordCount})
    database.EntityData.Leafs.Append("last-full-file-write-error-timestamp", types.YLeaf{"LastFullFileWriteErrorTimestamp", database.LastFullFileWriteErrorTimestamp})
    database.EntityData.Leafs.Append("incremental-file-write-interval", types.YLeaf{"IncrementalFileWriteInterval", database.IncrementalFileWriteInterval})
    database.EntityData.Leafs.Append("last-incremental-write-file-name", types.YLeaf{"LastIncrementalWriteFileName", database.LastIncrementalWriteFileName})
    database.EntityData.Leafs.Append("last-incremental-write-time", types.YLeaf{"LastIncrementalWriteTime", database.LastIncrementalWriteTime})
    database.EntityData.Leafs.Append("incremental-file-write-count", types.YLeaf{"IncrementalFileWriteCount", database.IncrementalFileWriteCount})
    database.EntityData.Leafs.Append("failed-incremental-file-write-count", types.YLeaf{"FailedIncrementalFileWriteCount", database.FailedIncrementalFileWriteCount})
    database.EntityData.Leafs.Append("incremental-file-record-count", types.YLeaf{"IncrementalFileRecordCount", database.IncrementalFileRecordCount})
    database.EntityData.Leafs.Append("last-incremental-file-write-error-timestamp", types.YLeaf{"LastIncrementalFileWriteErrorTimestamp", database.LastIncrementalFileWriteErrorTimestamp})

    database.EntityData.YListKeys = []string {}

    return &(database.EntityData)
}

// Dhcpv6_Nodes_Node_Base_AddrBindings
// IPv6 DHCP Base Binding
type Dhcpv6_Nodes_Node_Base_AddrBindings struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DHCPv6 base stats debug. The type is slice of
    // Dhcpv6_Nodes_Node_Base_AddrBindings_AddrBinding.
    AddrBinding []*Dhcpv6_Nodes_Node_Base_AddrBindings_AddrBinding
}

func (addrBindings *Dhcpv6_Nodes_Node_Base_AddrBindings) GetEntityData() *types.CommonEntityData {
    addrBindings.EntityData.YFilter = addrBindings.YFilter
    addrBindings.EntityData.YangName = "addr-bindings"
    addrBindings.EntityData.BundleName = "cisco_ios_xr"
    addrBindings.EntityData.ParentYangName = "base"
    addrBindings.EntityData.SegmentPath = "addr-bindings"
    addrBindings.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/base/" + addrBindings.EntityData.SegmentPath
    addrBindings.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addrBindings.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addrBindings.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addrBindings.EntityData.Children = types.NewOrderedMap()
    addrBindings.EntityData.Children.Append("addr-binding", types.YChild{"AddrBinding", nil})
    for i := range addrBindings.AddrBinding {
        addrBindings.EntityData.Children.Append(types.GetSegmentPath(addrBindings.AddrBinding[i]), types.YChild{"AddrBinding", addrBindings.AddrBinding[i]})
    }
    addrBindings.EntityData.Leafs = types.NewOrderedMap()

    addrBindings.EntityData.YListKeys = []string {}

    return &(addrBindings.EntityData)
}

// Dhcpv6_Nodes_Node_Base_AddrBindings_AddrBinding
// DHCPv6 base stats debug
type Dhcpv6_Nodes_Node_Base_AddrBindings_AddrBinding struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

    // DHCPV6 client state. The type is BagDhcpv6dFsmState.
    State interface{}

    // DHCPV6 access interface to client. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
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

func (addrBinding *Dhcpv6_Nodes_Node_Base_AddrBindings_AddrBinding) GetEntityData() *types.CommonEntityData {
    addrBinding.EntityData.YFilter = addrBinding.YFilter
    addrBinding.EntityData.YangName = "addr-binding"
    addrBinding.EntityData.BundleName = "cisco_ios_xr"
    addrBinding.EntityData.ParentYangName = "addr-bindings"
    addrBinding.EntityData.SegmentPath = "addr-binding" + types.AddKeyToken(addrBinding.AddrString, "addr-string")
    addrBinding.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/base/addr-bindings/" + addrBinding.EntityData.SegmentPath
    addrBinding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addrBinding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addrBinding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addrBinding.EntityData.Children = types.NewOrderedMap()
    addrBinding.EntityData.Leafs = types.NewOrderedMap()
    addrBinding.EntityData.Leafs.Append("addr-string", types.YLeaf{"AddrString", addrBinding.AddrString})
    addrBinding.EntityData.Leafs.Append("mac-address", types.YLeaf{"MacAddress", addrBinding.MacAddress})
    addrBinding.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", addrBinding.VrfName})
    addrBinding.EntityData.Leafs.Append("server-vrf-name", types.YLeaf{"ServerVrfName", addrBinding.ServerVrfName})
    addrBinding.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", addrBinding.Ipv6Address})
    addrBinding.EntityData.Leafs.Append("server-ipv6-address", types.YLeaf{"ServerIpv6Address", addrBinding.ServerIpv6Address})
    addrBinding.EntityData.Leafs.Append("reply-server-ipv6-address", types.YLeaf{"ReplyServerIpv6Address", addrBinding.ReplyServerIpv6Address})
    addrBinding.EntityData.Leafs.Append("lease-time", types.YLeaf{"LeaseTime", addrBinding.LeaseTime})
    addrBinding.EntityData.Leafs.Append("remaining-lease-time", types.YLeaf{"RemainingLeaseTime", addrBinding.RemainingLeaseTime})
    addrBinding.EntityData.Leafs.Append("state", types.YLeaf{"State", addrBinding.State})
    addrBinding.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", addrBinding.InterfaceName})
    addrBinding.EntityData.Leafs.Append("access-vrf-name", types.YLeaf{"AccessVrfName", addrBinding.AccessVrfName})
    addrBinding.EntityData.Leafs.Append("base-binding-tags", types.YLeaf{"BaseBindingTags", addrBinding.BaseBindingTags})
    addrBinding.EntityData.Leafs.Append("base-binding-outer-tag", types.YLeaf{"BaseBindingOuterTag", addrBinding.BaseBindingOuterTag})
    addrBinding.EntityData.Leafs.Append("base-binding-inner-tag", types.YLeaf{"BaseBindingInnerTag", addrBinding.BaseBindingInnerTag})
    addrBinding.EntityData.Leafs.Append("profile-name", types.YLeaf{"ProfileName", addrBinding.ProfileName})
    addrBinding.EntityData.Leafs.Append("is-nak-next-renew", types.YLeaf{"IsNakNextRenew", addrBinding.IsNakNextRenew})
    addrBinding.EntityData.Leafs.Append("subscriber-label", types.YLeaf{"SubscriberLabel", addrBinding.SubscriberLabel})
    addrBinding.EntityData.Leafs.Append("old-subscriber-label", types.YLeaf{"OldSubscriberLabel", addrBinding.OldSubscriberLabel})
    addrBinding.EntityData.Leafs.Append("rx-client-duid", types.YLeaf{"RxClientDuid", addrBinding.RxClientDuid})
    addrBinding.EntityData.Leafs.Append("tx-client-uid", types.YLeaf{"TxClientUid", addrBinding.TxClientUid})
    addrBinding.EntityData.Leafs.Append("rx-remote-id", types.YLeaf{"RxRemoteId", addrBinding.RxRemoteId})
    addrBinding.EntityData.Leafs.Append("tx-remote-id", types.YLeaf{"TxRemoteId", addrBinding.TxRemoteId})
    addrBinding.EntityData.Leafs.Append("rx-interface-id", types.YLeaf{"RxInterfaceId", addrBinding.RxInterfaceId})
    addrBinding.EntityData.Leafs.Append("tx-interface-id", types.YLeaf{"TxInterfaceId", addrBinding.TxInterfaceId})

    addrBinding.EntityData.YListKeys = []string {"AddrString"}

    return &(addrBinding.EntityData)
}

// Dhcpv6_Nodes_Node_Server
// IPv6 DHCP server operational data
type Dhcpv6_Nodes_Node_Server struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DHCPv6 server disconnect history.
    DisconnectHistories Dhcpv6_Nodes_Node_Server_DisconnectHistories

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

func (server *Dhcpv6_Nodes_Node_Server) GetEntityData() *types.CommonEntityData {
    server.EntityData.YFilter = server.YFilter
    server.EntityData.YangName = "server"
    server.EntityData.BundleName = "cisco_ios_xr"
    server.EntityData.ParentYangName = "node"
    server.EntityData.SegmentPath = "server"
    server.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/" + server.EntityData.SegmentPath
    server.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    server.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    server.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    server.EntityData.Children = types.NewOrderedMap()
    server.EntityData.Children.Append("disconnect-histories", types.YChild{"DisconnectHistories", &server.DisconnectHistories})
    server.EntityData.Children.Append("binding", types.YChild{"Binding", &server.Binding})
    server.EntityData.Children.Append("vrfs", types.YChild{"Vrfs", &server.Vrfs})
    server.EntityData.Children.Append("profiles", types.YChild{"Profiles", &server.Profiles})
    server.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &server.Interfaces})
    server.EntityData.Children.Append("statistics", types.YChild{"Statistics", &server.Statistics})
    server.EntityData.Children.Append("binding-options", types.YChild{"BindingOptions", &server.BindingOptions})
    server.EntityData.Leafs = types.NewOrderedMap()

    server.EntityData.YListKeys = []string {}

    return &(server.EntityData)
}

// Dhcpv6_Nodes_Node_Server_DisconnectHistories
// DHCPv6 server disconnect history
type Dhcpv6_Nodes_Node_Server_DisconnectHistories struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Single DHCPv6 server disconnect history. The type is slice of
    // Dhcpv6_Nodes_Node_Server_DisconnectHistories_DisconnectHistory.
    DisconnectHistory []*Dhcpv6_Nodes_Node_Server_DisconnectHistories_DisconnectHistory
}

func (disconnectHistories *Dhcpv6_Nodes_Node_Server_DisconnectHistories) GetEntityData() *types.CommonEntityData {
    disconnectHistories.EntityData.YFilter = disconnectHistories.YFilter
    disconnectHistories.EntityData.YangName = "disconnect-histories"
    disconnectHistories.EntityData.BundleName = "cisco_ios_xr"
    disconnectHistories.EntityData.ParentYangName = "server"
    disconnectHistories.EntityData.SegmentPath = "disconnect-histories"
    disconnectHistories.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/server/" + disconnectHistories.EntityData.SegmentPath
    disconnectHistories.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    disconnectHistories.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    disconnectHistories.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    disconnectHistories.EntityData.Children = types.NewOrderedMap()
    disconnectHistories.EntityData.Children.Append("disconnect-history", types.YChild{"DisconnectHistory", nil})
    for i := range disconnectHistories.DisconnectHistory {
        disconnectHistories.EntityData.Children.Append(types.GetSegmentPath(disconnectHistories.DisconnectHistory[i]), types.YChild{"DisconnectHistory", disconnectHistories.DisconnectHistory[i]})
    }
    disconnectHistories.EntityData.Leafs = types.NewOrderedMap()

    disconnectHistories.EntityData.YListKeys = []string {}

    return &(disconnectHistories.EntityData)
}

// Dhcpv6_Nodes_Node_Server_DisconnectHistories_DisconnectHistory
// Single DHCPv6 server disconnect history
type Dhcpv6_Nodes_Node_Server_DisconnectHistories_DisconnectHistory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Index. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Index interface{}

    // session start time epoch. The type is interface{} with range:
    // 0..18446744073709551615.
    SessionStartTimeEpoch interface{}

    // session end time epoch. The type is interface{} with range:
    // 0..18446744073709551615.
    SessionEndTimeEpoch interface{}

    // DiscReason. The type is string with length: 0..256.
    DiscReason interface{}

    // sub label. The type is interface{} with range: 0..4294967295.
    SubLabel interface{}

    // Client DUID. The type is string with length: 0..131.
    Duid interface{}

    // IAType. The type is string with length: 0..6.
    IaType interface{}

    // ia id. The type is interface{} with range: 0..4294967295.
    IaId interface{}

    // MACAddress. The type is string with length: 0..17.
    MacAddress interface{}
}

func (disconnectHistory *Dhcpv6_Nodes_Node_Server_DisconnectHistories_DisconnectHistory) GetEntityData() *types.CommonEntityData {
    disconnectHistory.EntityData.YFilter = disconnectHistory.YFilter
    disconnectHistory.EntityData.YangName = "disconnect-history"
    disconnectHistory.EntityData.BundleName = "cisco_ios_xr"
    disconnectHistory.EntityData.ParentYangName = "disconnect-histories"
    disconnectHistory.EntityData.SegmentPath = "disconnect-history" + types.AddKeyToken(disconnectHistory.Index, "index")
    disconnectHistory.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/server/disconnect-histories/" + disconnectHistory.EntityData.SegmentPath
    disconnectHistory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    disconnectHistory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    disconnectHistory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    disconnectHistory.EntityData.Children = types.NewOrderedMap()
    disconnectHistory.EntityData.Leafs = types.NewOrderedMap()
    disconnectHistory.EntityData.Leafs.Append("index", types.YLeaf{"Index", disconnectHistory.Index})
    disconnectHistory.EntityData.Leafs.Append("session-start-time-epoch", types.YLeaf{"SessionStartTimeEpoch", disconnectHistory.SessionStartTimeEpoch})
    disconnectHistory.EntityData.Leafs.Append("session-end-time-epoch", types.YLeaf{"SessionEndTimeEpoch", disconnectHistory.SessionEndTimeEpoch})
    disconnectHistory.EntityData.Leafs.Append("disc-reason", types.YLeaf{"DiscReason", disconnectHistory.DiscReason})
    disconnectHistory.EntityData.Leafs.Append("sub-label", types.YLeaf{"SubLabel", disconnectHistory.SubLabel})
    disconnectHistory.EntityData.Leafs.Append("duid", types.YLeaf{"Duid", disconnectHistory.Duid})
    disconnectHistory.EntityData.Leafs.Append("ia-type", types.YLeaf{"IaType", disconnectHistory.IaType})
    disconnectHistory.EntityData.Leafs.Append("ia-id", types.YLeaf{"IaId", disconnectHistory.IaId})
    disconnectHistory.EntityData.Leafs.Append("mac-address", types.YLeaf{"MacAddress", disconnectHistory.MacAddress})

    disconnectHistory.EntityData.YListKeys = []string {"Index"}

    return &(disconnectHistory.EntityData)
}

// Dhcpv6_Nodes_Node_Server_Binding
// DHCPV6 server bindings
type Dhcpv6_Nodes_Node_Server_Binding struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DHCPV6 server binding summary.
    Summary Dhcpv6_Nodes_Node_Server_Binding_Summary

    // DHCPV6 server client bindings.
    Clients Dhcpv6_Nodes_Node_Server_Binding_Clients
}

func (binding *Dhcpv6_Nodes_Node_Server_Binding) GetEntityData() *types.CommonEntityData {
    binding.EntityData.YFilter = binding.YFilter
    binding.EntityData.YangName = "binding"
    binding.EntityData.BundleName = "cisco_ios_xr"
    binding.EntityData.ParentYangName = "server"
    binding.EntityData.SegmentPath = "binding"
    binding.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/server/" + binding.EntityData.SegmentPath
    binding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    binding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    binding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    binding.EntityData.Children = types.NewOrderedMap()
    binding.EntityData.Children.Append("summary", types.YChild{"Summary", &binding.Summary})
    binding.EntityData.Children.Append("clients", types.YChild{"Clients", &binding.Clients})
    binding.EntityData.Leafs = types.NewOrderedMap()

    binding.EntityData.YListKeys = []string {}

    return &(binding.EntityData)
}

// Dhcpv6_Nodes_Node_Server_Binding_Summary
// DHCPV6 server binding summary
type Dhcpv6_Nodes_Node_Server_Binding_Summary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Total number of clients. The type is interface{} with range: 0..4294967295.
    Clients interface{}

    // IANA server binding summary.
    Iana Dhcpv6_Nodes_Node_Server_Binding_Summary_Iana

    // IAPD server binding summary.
    Iapd Dhcpv6_Nodes_Node_Server_Binding_Summary_Iapd
}

func (summary *Dhcpv6_Nodes_Node_Server_Binding_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "binding"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/server/binding/" + summary.EntityData.SegmentPath
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = types.NewOrderedMap()
    summary.EntityData.Children.Append("iana", types.YChild{"Iana", &summary.Iana})
    summary.EntityData.Children.Append("iapd", types.YChild{"Iapd", &summary.Iapd})
    summary.EntityData.Leafs = types.NewOrderedMap()
    summary.EntityData.Leafs.Append("clients", types.YLeaf{"Clients", summary.Clients})

    summary.EntityData.YListKeys = []string {}

    return &(summary.EntityData)
}

// Dhcpv6_Nodes_Node_Server_Binding_Summary_Iana
// IANA server binding summary
type Dhcpv6_Nodes_Node_Server_Binding_Summary_Iana struct {
    EntityData types.CommonEntityData
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

func (iana *Dhcpv6_Nodes_Node_Server_Binding_Summary_Iana) GetEntityData() *types.CommonEntityData {
    iana.EntityData.YFilter = iana.YFilter
    iana.EntityData.YangName = "iana"
    iana.EntityData.BundleName = "cisco_ios_xr"
    iana.EntityData.ParentYangName = "summary"
    iana.EntityData.SegmentPath = "iana"
    iana.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/server/binding/summary/" + iana.EntityData.SegmentPath
    iana.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    iana.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    iana.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    iana.EntityData.Children = types.NewOrderedMap()
    iana.EntityData.Leafs = types.NewOrderedMap()
    iana.EntityData.Leafs.Append("initializing-clients", types.YLeaf{"InitializingClients", iana.InitializingClients})
    iana.EntityData.Leafs.Append("dpm-waiting-clients", types.YLeaf{"DpmWaitingClients", iana.DpmWaitingClients})
    iana.EntityData.Leafs.Append("daps-waiting-clients", types.YLeaf{"DapsWaitingClients", iana.DapsWaitingClients})
    iana.EntityData.Leafs.Append("request-waiting-clients", types.YLeaf{"RequestWaitingClients", iana.RequestWaitingClients})
    iana.EntityData.Leafs.Append("iedge-waiting-clients", types.YLeaf{"IedgeWaitingClients", iana.IedgeWaitingClients})
    iana.EntityData.Leafs.Append("rib-waiting-clients", types.YLeaf{"RibWaitingClients", iana.RibWaitingClients})
    iana.EntityData.Leafs.Append("bound-clients", types.YLeaf{"BoundClients", iana.BoundClients})

    iana.EntityData.YListKeys = []string {}

    return &(iana.EntityData)
}

// Dhcpv6_Nodes_Node_Server_Binding_Summary_Iapd
// IAPD server binding summary
type Dhcpv6_Nodes_Node_Server_Binding_Summary_Iapd struct {
    EntityData types.CommonEntityData
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

func (iapd *Dhcpv6_Nodes_Node_Server_Binding_Summary_Iapd) GetEntityData() *types.CommonEntityData {
    iapd.EntityData.YFilter = iapd.YFilter
    iapd.EntityData.YangName = "iapd"
    iapd.EntityData.BundleName = "cisco_ios_xr"
    iapd.EntityData.ParentYangName = "summary"
    iapd.EntityData.SegmentPath = "iapd"
    iapd.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/server/binding/summary/" + iapd.EntityData.SegmentPath
    iapd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    iapd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    iapd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    iapd.EntityData.Children = types.NewOrderedMap()
    iapd.EntityData.Leafs = types.NewOrderedMap()
    iapd.EntityData.Leafs.Append("initializing-clients", types.YLeaf{"InitializingClients", iapd.InitializingClients})
    iapd.EntityData.Leafs.Append("dpm-waiting-clients", types.YLeaf{"DpmWaitingClients", iapd.DpmWaitingClients})
    iapd.EntityData.Leafs.Append("daps-waiting-clients", types.YLeaf{"DapsWaitingClients", iapd.DapsWaitingClients})
    iapd.EntityData.Leafs.Append("request-waiting-clients", types.YLeaf{"RequestWaitingClients", iapd.RequestWaitingClients})
    iapd.EntityData.Leafs.Append("iedge-waiting-clients", types.YLeaf{"IedgeWaitingClients", iapd.IedgeWaitingClients})
    iapd.EntityData.Leafs.Append("rib-waiting-clients", types.YLeaf{"RibWaitingClients", iapd.RibWaitingClients})
    iapd.EntityData.Leafs.Append("bound-clients", types.YLeaf{"BoundClients", iapd.BoundClients})

    iapd.EntityData.YListKeys = []string {}

    return &(iapd.EntityData)
}

// Dhcpv6_Nodes_Node_Server_Binding_Clients
// DHCPV6 server client bindings
type Dhcpv6_Nodes_Node_Server_Binding_Clients struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Single DHCPV6 server binding. The type is slice of
    // Dhcpv6_Nodes_Node_Server_Binding_Clients_Client.
    Client []*Dhcpv6_Nodes_Node_Server_Binding_Clients_Client
}

func (clients *Dhcpv6_Nodes_Node_Server_Binding_Clients) GetEntityData() *types.CommonEntityData {
    clients.EntityData.YFilter = clients.YFilter
    clients.EntityData.YangName = "clients"
    clients.EntityData.BundleName = "cisco_ios_xr"
    clients.EntityData.ParentYangName = "binding"
    clients.EntityData.SegmentPath = "clients"
    clients.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/server/binding/" + clients.EntityData.SegmentPath
    clients.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clients.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clients.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clients.EntityData.Children = types.NewOrderedMap()
    clients.EntityData.Children.Append("client", types.YChild{"Client", nil})
    for i := range clients.Client {
        clients.EntityData.Children.Append(types.GetSegmentPath(clients.Client[i]), types.YChild{"Client", clients.Client[i]})
    }
    clients.EntityData.Leafs = types.NewOrderedMap()

    clients.EntityData.YListKeys = []string {}

    return &(clients.EntityData)
}

// Dhcpv6_Nodes_Node_Server_Binding_Clients_Client
// Single DHCPV6 server binding
type Dhcpv6_Nodes_Node_Server_Binding_Clients_Client struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    // [a-zA-Z0-9._/-]+.
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
    Srgp2p interface{}

    // DHCPV6 SRG VRF NAME. The type is string with length: 0..33.
    SrgVrfName interface{}

    // DHCPV6 SERG state. The type is interface{} with range: 0..4294967295.
    SesrgState interface{}

    // DHCPV6 SERG Intf Role. The type is interface{} with range: 0..4294967295.
    SergIntfRole interface{}

    // List of DHCPv6 IA_ID/PDs.
    IaIdPd Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd
}

func (client *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client) GetEntityData() *types.CommonEntityData {
    client.EntityData.YFilter = client.YFilter
    client.EntityData.YangName = "client"
    client.EntityData.BundleName = "cisco_ios_xr"
    client.EntityData.ParentYangName = "clients"
    client.EntityData.SegmentPath = "client" + types.AddKeyToken(client.ClientId, "client-id")
    client.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/server/binding/clients/" + client.EntityData.SegmentPath
    client.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    client.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    client.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    client.EntityData.Children = types.NewOrderedMap()
    client.EntityData.Children.Append("ia-id-pd", types.YChild{"IaIdPd", &client.IaIdPd})
    client.EntityData.Leafs = types.NewOrderedMap()
    client.EntityData.Leafs.Append("client-id", types.YLeaf{"ClientId", client.ClientId})
    client.EntityData.Leafs.Append("duid", types.YLeaf{"Duid", client.Duid})
    client.EntityData.Leafs.Append("client-id-xr", types.YLeaf{"ClientIdXr", client.ClientIdXr})
    client.EntityData.Leafs.Append("client-flag", types.YLeaf{"ClientFlag", client.ClientFlag})
    client.EntityData.Leafs.Append("subscriber-label", types.YLeaf{"SubscriberLabel", client.SubscriberLabel})
    client.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", client.VrfName})
    client.EntityData.Leafs.Append("mac-address", types.YLeaf{"MacAddress", client.MacAddress})
    client.EntityData.Leafs.Append("ia-id-p-ds", types.YLeaf{"IaIdPDs", client.IaIdPDs})
    client.EntityData.Leafs.Append("link-local-address", types.YLeaf{"LinkLocalAddress", client.LinkLocalAddress})
    client.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", client.InterfaceName})
    client.EntityData.Leafs.Append("access-vrf-name", types.YLeaf{"AccessVrfName", client.AccessVrfName})
    client.EntityData.Leafs.Append("server-binding-tags", types.YLeaf{"ServerBindingTags", client.ServerBindingTags})
    client.EntityData.Leafs.Append("server-binding-outer-tag", types.YLeaf{"ServerBindingOuterTag", client.ServerBindingOuterTag})
    client.EntityData.Leafs.Append("server-binding-inner-tag", types.YLeaf{"ServerBindingInnerTag", client.ServerBindingInnerTag})
    client.EntityData.Leafs.Append("pool-name", types.YLeaf{"PoolName", client.PoolName})
    client.EntityData.Leafs.Append("profile-name", types.YLeaf{"ProfileName", client.ProfileName})
    client.EntityData.Leafs.Append("framed-ipv6-prefix", types.YLeaf{"FramedIpv6Prefix", client.FramedIpv6Prefix})
    client.EntityData.Leafs.Append("framed-prefix-length", types.YLeaf{"FramedPrefixLength", client.FramedPrefixLength})
    client.EntityData.Leafs.Append("class-name", types.YLeaf{"ClassName", client.ClassName})
    client.EntityData.Leafs.Append("rx-remote-id", types.YLeaf{"RxRemoteId", client.RxRemoteId})
    client.EntityData.Leafs.Append("rx-interface-id", types.YLeaf{"RxInterfaceId", client.RxInterfaceId})
    client.EntityData.Leafs.Append("prefix-pool-name", types.YLeaf{"PrefixPoolName", client.PrefixPoolName})
    client.EntityData.Leafs.Append("address-pool-name", types.YLeaf{"AddressPoolName", client.AddressPoolName})
    client.EntityData.Leafs.Append("dns-server-count", types.YLeaf{"DnsServerCount", client.DnsServerCount})
    client.EntityData.Leafs.Append("is-nak-next-renew", types.YLeaf{"IsNakNextRenew", client.IsNakNextRenew})
    client.EntityData.Leafs.Append("srg-state", types.YLeaf{"SrgState", client.SrgState})
    client.EntityData.Leafs.Append("srg-intf-role", types.YLeaf{"SrgIntfRole", client.SrgIntfRole})
    client.EntityData.Leafs.Append("srgp2p", types.YLeaf{"Srgp2p", client.Srgp2p})
    client.EntityData.Leafs.Append("srg-vrf-name", types.YLeaf{"SrgVrfName", client.SrgVrfName})
    client.EntityData.Leafs.Append("sesrg-state", types.YLeaf{"SesrgState", client.SesrgState})
    client.EntityData.Leafs.Append("serg-intf-role", types.YLeaf{"SergIntfRole", client.SergIntfRole})

    client.EntityData.YListKeys = []string {"ClientId"}

    return &(client.EntityData)
}

// Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd
// List of DHCPv6 IA_ID/PDs
type Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bag dhcpv6d ia id pd info. The type is slice of
    // Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6dIaIdPdInfo.
    BagDhcpv6dIaIdPdInfo []*Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6dIaIdPdInfo
}

func (iaIdPd *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd) GetEntityData() *types.CommonEntityData {
    iaIdPd.EntityData.YFilter = iaIdPd.YFilter
    iaIdPd.EntityData.YangName = "ia-id-pd"
    iaIdPd.EntityData.BundleName = "cisco_ios_xr"
    iaIdPd.EntityData.ParentYangName = "client"
    iaIdPd.EntityData.SegmentPath = "ia-id-pd"
    iaIdPd.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/server/binding/clients/client/" + iaIdPd.EntityData.SegmentPath
    iaIdPd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    iaIdPd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    iaIdPd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    iaIdPd.EntityData.Children = types.NewOrderedMap()
    iaIdPd.EntityData.Children.Append("bag-dhcpv6d-ia-id-pd-info", types.YChild{"BagDhcpv6dIaIdPdInfo", nil})
    for i := range iaIdPd.BagDhcpv6dIaIdPdInfo {
        types.SetYListKey(iaIdPd.BagDhcpv6dIaIdPdInfo[i], i)
        iaIdPd.EntityData.Children.Append(types.GetSegmentPath(iaIdPd.BagDhcpv6dIaIdPdInfo[i]), types.YChild{"BagDhcpv6dIaIdPdInfo", iaIdPd.BagDhcpv6dIaIdPdInfo[i]})
    }
    iaIdPd.EntityData.Leafs = types.NewOrderedMap()

    iaIdPd.EntityData.YListKeys = []string {}

    return &(iaIdPd.EntityData)
}

// Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6dIaIdPdInfo
// bag dhcpv6d ia id pd info
type Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6dIaIdPdInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // IA type. The type is BagDhcpv6dIaId.
    IaType interface{}

    // IA_ID of this IA. The type is interface{} with range: 0..4294967295.
    IaId interface{}

    // FSM Flag for this IA. The type is interface{} with range: 0..4294967295.
    Flags interface{}

    // Total address in this IA. The type is interface{} with range: 0..65535.
    TotalAddress interface{}

    // State. The type is BagDhcpv6dFsmState.
    State interface{}

    // List of addresses in this IA.
    Addresses Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6dIaIdPdInfo_Addresses
}

func (bagDhcpv6dIaIdPdInfo *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6dIaIdPdInfo) GetEntityData() *types.CommonEntityData {
    bagDhcpv6dIaIdPdInfo.EntityData.YFilter = bagDhcpv6dIaIdPdInfo.YFilter
    bagDhcpv6dIaIdPdInfo.EntityData.YangName = "bag-dhcpv6d-ia-id-pd-info"
    bagDhcpv6dIaIdPdInfo.EntityData.BundleName = "cisco_ios_xr"
    bagDhcpv6dIaIdPdInfo.EntityData.ParentYangName = "ia-id-pd"
    bagDhcpv6dIaIdPdInfo.EntityData.SegmentPath = "bag-dhcpv6d-ia-id-pd-info" + types.AddNoKeyToken(bagDhcpv6dIaIdPdInfo)
    bagDhcpv6dIaIdPdInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/server/binding/clients/client/ia-id-pd/" + bagDhcpv6dIaIdPdInfo.EntityData.SegmentPath
    bagDhcpv6dIaIdPdInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bagDhcpv6dIaIdPdInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bagDhcpv6dIaIdPdInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bagDhcpv6dIaIdPdInfo.EntityData.Children = types.NewOrderedMap()
    bagDhcpv6dIaIdPdInfo.EntityData.Children.Append("addresses", types.YChild{"Addresses", &bagDhcpv6dIaIdPdInfo.Addresses})
    bagDhcpv6dIaIdPdInfo.EntityData.Leafs = types.NewOrderedMap()
    bagDhcpv6dIaIdPdInfo.EntityData.Leafs.Append("ia-type", types.YLeaf{"IaType", bagDhcpv6dIaIdPdInfo.IaType})
    bagDhcpv6dIaIdPdInfo.EntityData.Leafs.Append("ia-id", types.YLeaf{"IaId", bagDhcpv6dIaIdPdInfo.IaId})
    bagDhcpv6dIaIdPdInfo.EntityData.Leafs.Append("flags", types.YLeaf{"Flags", bagDhcpv6dIaIdPdInfo.Flags})
    bagDhcpv6dIaIdPdInfo.EntityData.Leafs.Append("total-address", types.YLeaf{"TotalAddress", bagDhcpv6dIaIdPdInfo.TotalAddress})
    bagDhcpv6dIaIdPdInfo.EntityData.Leafs.Append("state", types.YLeaf{"State", bagDhcpv6dIaIdPdInfo.State})

    bagDhcpv6dIaIdPdInfo.EntityData.YListKeys = []string {}

    return &(bagDhcpv6dIaIdPdInfo.EntityData)
}

// Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6dIaIdPdInfo_Addresses
// List of addresses in this IA
type Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6dIaIdPdInfo_Addresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bag dhcpv6d addr attrb. The type is slice of
    // Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6dIaIdPdInfo_Addresses_BagDhcpv6dAddrAttrb.
    BagDhcpv6dAddrAttrb []*Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6dIaIdPdInfo_Addresses_BagDhcpv6dAddrAttrb
}

func (addresses *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6dIaIdPdInfo_Addresses) GetEntityData() *types.CommonEntityData {
    addresses.EntityData.YFilter = addresses.YFilter
    addresses.EntityData.YangName = "addresses"
    addresses.EntityData.BundleName = "cisco_ios_xr"
    addresses.EntityData.ParentYangName = "bag-dhcpv6d-ia-id-pd-info"
    addresses.EntityData.SegmentPath = "addresses"
    addresses.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/server/binding/clients/client/ia-id-pd/bag-dhcpv6d-ia-id-pd-info/" + addresses.EntityData.SegmentPath
    addresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addresses.EntityData.Children = types.NewOrderedMap()
    addresses.EntityData.Children.Append("bag-dhcpv6d-addr-attrb", types.YChild{"BagDhcpv6dAddrAttrb", nil})
    for i := range addresses.BagDhcpv6dAddrAttrb {
        types.SetYListKey(addresses.BagDhcpv6dAddrAttrb[i], i)
        addresses.EntityData.Children.Append(types.GetSegmentPath(addresses.BagDhcpv6dAddrAttrb[i]), types.YChild{"BagDhcpv6dAddrAttrb", addresses.BagDhcpv6dAddrAttrb[i]})
    }
    addresses.EntityData.Leafs = types.NewOrderedMap()

    addresses.EntityData.YListKeys = []string {}

    return &(addresses.EntityData)
}

// Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6dIaIdPdInfo_Addresses_BagDhcpv6dAddrAttrb
// bag dhcpv6d addr attrb
type Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6dIaIdPdInfo_Addresses_BagDhcpv6dAddrAttrb struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (bagDhcpv6dAddrAttrb *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6dIaIdPdInfo_Addresses_BagDhcpv6dAddrAttrb) GetEntityData() *types.CommonEntityData {
    bagDhcpv6dAddrAttrb.EntityData.YFilter = bagDhcpv6dAddrAttrb.YFilter
    bagDhcpv6dAddrAttrb.EntityData.YangName = "bag-dhcpv6d-addr-attrb"
    bagDhcpv6dAddrAttrb.EntityData.BundleName = "cisco_ios_xr"
    bagDhcpv6dAddrAttrb.EntityData.ParentYangName = "addresses"
    bagDhcpv6dAddrAttrb.EntityData.SegmentPath = "bag-dhcpv6d-addr-attrb" + types.AddNoKeyToken(bagDhcpv6dAddrAttrb)
    bagDhcpv6dAddrAttrb.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/server/binding/clients/client/ia-id-pd/bag-dhcpv6d-ia-id-pd-info/addresses/" + bagDhcpv6dAddrAttrb.EntityData.SegmentPath
    bagDhcpv6dAddrAttrb.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bagDhcpv6dAddrAttrb.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bagDhcpv6dAddrAttrb.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bagDhcpv6dAddrAttrb.EntityData.Children = types.NewOrderedMap()
    bagDhcpv6dAddrAttrb.EntityData.Leafs = types.NewOrderedMap()
    bagDhcpv6dAddrAttrb.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", bagDhcpv6dAddrAttrb.Prefix})
    bagDhcpv6dAddrAttrb.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", bagDhcpv6dAddrAttrb.PrefixLength})
    bagDhcpv6dAddrAttrb.EntityData.Leafs.Append("lease-time", types.YLeaf{"LeaseTime", bagDhcpv6dAddrAttrb.LeaseTime})
    bagDhcpv6dAddrAttrb.EntityData.Leafs.Append("remaining-lease-time", types.YLeaf{"RemainingLeaseTime", bagDhcpv6dAddrAttrb.RemainingLeaseTime})

    bagDhcpv6dAddrAttrb.EntityData.YListKeys = []string {}

    return &(bagDhcpv6dAddrAttrb.EntityData)
}

// Dhcpv6_Nodes_Node_Server_Vrfs
// DHCPV6 server list of VRF names
type Dhcpv6_Nodes_Node_Server_Vrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 DHCP server VRF name. The type is slice of
    // Dhcpv6_Nodes_Node_Server_Vrfs_Vrf.
    Vrf []*Dhcpv6_Nodes_Node_Server_Vrfs_Vrf
}

func (vrfs *Dhcpv6_Nodes_Node_Server_Vrfs) GetEntityData() *types.CommonEntityData {
    vrfs.EntityData.YFilter = vrfs.YFilter
    vrfs.EntityData.YangName = "vrfs"
    vrfs.EntityData.BundleName = "cisco_ios_xr"
    vrfs.EntityData.ParentYangName = "server"
    vrfs.EntityData.SegmentPath = "vrfs"
    vrfs.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/server/" + vrfs.EntityData.SegmentPath
    vrfs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfs.EntityData.Children = types.NewOrderedMap()
    vrfs.EntityData.Children.Append("vrf", types.YChild{"Vrf", nil})
    for i := range vrfs.Vrf {
        vrfs.EntityData.Children.Append(types.GetSegmentPath(vrfs.Vrf[i]), types.YChild{"Vrf", vrfs.Vrf[i]})
    }
    vrfs.EntityData.Leafs = types.NewOrderedMap()

    vrfs.EntityData.YListKeys = []string {}

    return &(vrfs.EntityData)
}

// Dhcpv6_Nodes_Node_Server_Vrfs_Vrf
// IPv6 DHCP server VRF name
type Dhcpv6_Nodes_Node_Server_Vrfs_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. VRF name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // IPv6 DHCP server statistics.
    Statistics Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics
}

func (vrf *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "vrfs"
    vrf.EntityData.SegmentPath = "vrf" + types.AddKeyToken(vrf.VrfName, "vrf-name")
    vrf.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/server/vrfs/" + vrf.EntityData.SegmentPath
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = types.NewOrderedMap()
    vrf.EntityData.Children.Append("statistics", types.YChild{"Statistics", &vrf.Statistics})
    vrf.EntityData.Leafs = types.NewOrderedMap()
    vrf.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", vrf.VrfName})

    vrf.EntityData.YListKeys = []string {"VrfName"}

    return &(vrf.EntityData)
}

// Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics
// IPv6 DHCP server statistics
type Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics struct {
    EntityData types.CommonEntityData
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

func (statistics *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "vrf"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/server/vrfs/vrf/" + statistics.EntityData.SegmentPath
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = types.NewOrderedMap()
    statistics.EntityData.Children.Append("solicit", types.YChild{"Solicit", &statistics.Solicit})
    statistics.EntityData.Children.Append("advertise", types.YChild{"Advertise", &statistics.Advertise})
    statistics.EntityData.Children.Append("request", types.YChild{"Request", &statistics.Request})
    statistics.EntityData.Children.Append("reply", types.YChild{"Reply", &statistics.Reply})
    statistics.EntityData.Children.Append("confirm", types.YChild{"Confirm", &statistics.Confirm})
    statistics.EntityData.Children.Append("decline", types.YChild{"Decline", &statistics.Decline})
    statistics.EntityData.Children.Append("renew", types.YChild{"Renew", &statistics.Renew})
    statistics.EntityData.Children.Append("rebind", types.YChild{"Rebind", &statistics.Rebind})
    statistics.EntityData.Children.Append("release", types.YChild{"Release", &statistics.Release})
    statistics.EntityData.Children.Append("reconfig", types.YChild{"Reconfig", &statistics.Reconfig})
    statistics.EntityData.Children.Append("inform", types.YChild{"Inform", &statistics.Inform})
    statistics.EntityData.Children.Append("relay-forward", types.YChild{"RelayForward", &statistics.RelayForward})
    statistics.EntityData.Children.Append("relay-reply", types.YChild{"RelayReply", &statistics.RelayReply})
    statistics.EntityData.Children.Append("lease-query", types.YChild{"LeaseQuery", &statistics.LeaseQuery})
    statistics.EntityData.Children.Append("lease-query-reply", types.YChild{"LeaseQueryReply", &statistics.LeaseQueryReply})
    statistics.EntityData.Children.Append("lease-query-done", types.YChild{"LeaseQueryDone", &statistics.LeaseQueryDone})
    statistics.EntityData.Children.Append("lease-query-data", types.YChild{"LeaseQueryData", &statistics.LeaseQueryData})
    statistics.EntityData.Leafs = types.NewOrderedMap()

    statistics.EntityData.YListKeys = []string {}

    return &(statistics.EntityData)
}

// Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Solicit
// DHCPV6 solicit packets
type Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Solicit struct {
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

func (solicit *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Solicit) GetEntityData() *types.CommonEntityData {
    solicit.EntityData.YFilter = solicit.YFilter
    solicit.EntityData.YangName = "solicit"
    solicit.EntityData.BundleName = "cisco_ios_xr"
    solicit.EntityData.ParentYangName = "statistics"
    solicit.EntityData.SegmentPath = "solicit"
    solicit.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/server/vrfs/vrf/statistics/" + solicit.EntityData.SegmentPath
    solicit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    solicit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    solicit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    solicit.EntityData.Children = types.NewOrderedMap()
    solicit.EntityData.Leafs = types.NewOrderedMap()
    solicit.EntityData.Leafs.Append("received-packets", types.YLeaf{"ReceivedPackets", solicit.ReceivedPackets})
    solicit.EntityData.Leafs.Append("transmitted-packets", types.YLeaf{"TransmittedPackets", solicit.TransmittedPackets})
    solicit.EntityData.Leafs.Append("dropped-packets", types.YLeaf{"DroppedPackets", solicit.DroppedPackets})

    solicit.EntityData.YListKeys = []string {}

    return &(solicit.EntityData)
}

// Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Advertise
// DHCPV6 advertise packets
type Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Advertise struct {
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

func (advertise *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Advertise) GetEntityData() *types.CommonEntityData {
    advertise.EntityData.YFilter = advertise.YFilter
    advertise.EntityData.YangName = "advertise"
    advertise.EntityData.BundleName = "cisco_ios_xr"
    advertise.EntityData.ParentYangName = "statistics"
    advertise.EntityData.SegmentPath = "advertise"
    advertise.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/server/vrfs/vrf/statistics/" + advertise.EntityData.SegmentPath
    advertise.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    advertise.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    advertise.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    advertise.EntityData.Children = types.NewOrderedMap()
    advertise.EntityData.Leafs = types.NewOrderedMap()
    advertise.EntityData.Leafs.Append("received-packets", types.YLeaf{"ReceivedPackets", advertise.ReceivedPackets})
    advertise.EntityData.Leafs.Append("transmitted-packets", types.YLeaf{"TransmittedPackets", advertise.TransmittedPackets})
    advertise.EntityData.Leafs.Append("dropped-packets", types.YLeaf{"DroppedPackets", advertise.DroppedPackets})

    advertise.EntityData.YListKeys = []string {}

    return &(advertise.EntityData)
}

// Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Request
// DHCPV6 request packets
type Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Request struct {
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

func (request *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Request) GetEntityData() *types.CommonEntityData {
    request.EntityData.YFilter = request.YFilter
    request.EntityData.YangName = "request"
    request.EntityData.BundleName = "cisco_ios_xr"
    request.EntityData.ParentYangName = "statistics"
    request.EntityData.SegmentPath = "request"
    request.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/server/vrfs/vrf/statistics/" + request.EntityData.SegmentPath
    request.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    request.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    request.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    request.EntityData.Children = types.NewOrderedMap()
    request.EntityData.Leafs = types.NewOrderedMap()
    request.EntityData.Leafs.Append("received-packets", types.YLeaf{"ReceivedPackets", request.ReceivedPackets})
    request.EntityData.Leafs.Append("transmitted-packets", types.YLeaf{"TransmittedPackets", request.TransmittedPackets})
    request.EntityData.Leafs.Append("dropped-packets", types.YLeaf{"DroppedPackets", request.DroppedPackets})

    request.EntityData.YListKeys = []string {}

    return &(request.EntityData)
}

// Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Reply
// DHCPV6 reply packets
type Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Reply struct {
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

func (reply *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Reply) GetEntityData() *types.CommonEntityData {
    reply.EntityData.YFilter = reply.YFilter
    reply.EntityData.YangName = "reply"
    reply.EntityData.BundleName = "cisco_ios_xr"
    reply.EntityData.ParentYangName = "statistics"
    reply.EntityData.SegmentPath = "reply"
    reply.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/server/vrfs/vrf/statistics/" + reply.EntityData.SegmentPath
    reply.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reply.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reply.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reply.EntityData.Children = types.NewOrderedMap()
    reply.EntityData.Leafs = types.NewOrderedMap()
    reply.EntityData.Leafs.Append("received-packets", types.YLeaf{"ReceivedPackets", reply.ReceivedPackets})
    reply.EntityData.Leafs.Append("transmitted-packets", types.YLeaf{"TransmittedPackets", reply.TransmittedPackets})
    reply.EntityData.Leafs.Append("dropped-packets", types.YLeaf{"DroppedPackets", reply.DroppedPackets})

    reply.EntityData.YListKeys = []string {}

    return &(reply.EntityData)
}

// Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Confirm
// DHCPV6 confirm packets
type Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Confirm struct {
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

func (confirm *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Confirm) GetEntityData() *types.CommonEntityData {
    confirm.EntityData.YFilter = confirm.YFilter
    confirm.EntityData.YangName = "confirm"
    confirm.EntityData.BundleName = "cisco_ios_xr"
    confirm.EntityData.ParentYangName = "statistics"
    confirm.EntityData.SegmentPath = "confirm"
    confirm.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/server/vrfs/vrf/statistics/" + confirm.EntityData.SegmentPath
    confirm.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    confirm.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    confirm.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    confirm.EntityData.Children = types.NewOrderedMap()
    confirm.EntityData.Leafs = types.NewOrderedMap()
    confirm.EntityData.Leafs.Append("received-packets", types.YLeaf{"ReceivedPackets", confirm.ReceivedPackets})
    confirm.EntityData.Leafs.Append("transmitted-packets", types.YLeaf{"TransmittedPackets", confirm.TransmittedPackets})
    confirm.EntityData.Leafs.Append("dropped-packets", types.YLeaf{"DroppedPackets", confirm.DroppedPackets})

    confirm.EntityData.YListKeys = []string {}

    return &(confirm.EntityData)
}

// Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Decline
// DHCPV6 decline packets
type Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Decline struct {
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

func (decline *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Decline) GetEntityData() *types.CommonEntityData {
    decline.EntityData.YFilter = decline.YFilter
    decline.EntityData.YangName = "decline"
    decline.EntityData.BundleName = "cisco_ios_xr"
    decline.EntityData.ParentYangName = "statistics"
    decline.EntityData.SegmentPath = "decline"
    decline.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/server/vrfs/vrf/statistics/" + decline.EntityData.SegmentPath
    decline.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    decline.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    decline.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    decline.EntityData.Children = types.NewOrderedMap()
    decline.EntityData.Leafs = types.NewOrderedMap()
    decline.EntityData.Leafs.Append("received-packets", types.YLeaf{"ReceivedPackets", decline.ReceivedPackets})
    decline.EntityData.Leafs.Append("transmitted-packets", types.YLeaf{"TransmittedPackets", decline.TransmittedPackets})
    decline.EntityData.Leafs.Append("dropped-packets", types.YLeaf{"DroppedPackets", decline.DroppedPackets})

    decline.EntityData.YListKeys = []string {}

    return &(decline.EntityData)
}

// Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Renew
// DHCPV6 renew packets
type Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Renew struct {
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

func (renew *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Renew) GetEntityData() *types.CommonEntityData {
    renew.EntityData.YFilter = renew.YFilter
    renew.EntityData.YangName = "renew"
    renew.EntityData.BundleName = "cisco_ios_xr"
    renew.EntityData.ParentYangName = "statistics"
    renew.EntityData.SegmentPath = "renew"
    renew.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/server/vrfs/vrf/statistics/" + renew.EntityData.SegmentPath
    renew.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    renew.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    renew.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    renew.EntityData.Children = types.NewOrderedMap()
    renew.EntityData.Leafs = types.NewOrderedMap()
    renew.EntityData.Leafs.Append("received-packets", types.YLeaf{"ReceivedPackets", renew.ReceivedPackets})
    renew.EntityData.Leafs.Append("transmitted-packets", types.YLeaf{"TransmittedPackets", renew.TransmittedPackets})
    renew.EntityData.Leafs.Append("dropped-packets", types.YLeaf{"DroppedPackets", renew.DroppedPackets})

    renew.EntityData.YListKeys = []string {}

    return &(renew.EntityData)
}

// Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Rebind
// DHCPV6 rebind packets
type Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Rebind struct {
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

func (rebind *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Rebind) GetEntityData() *types.CommonEntityData {
    rebind.EntityData.YFilter = rebind.YFilter
    rebind.EntityData.YangName = "rebind"
    rebind.EntityData.BundleName = "cisco_ios_xr"
    rebind.EntityData.ParentYangName = "statistics"
    rebind.EntityData.SegmentPath = "rebind"
    rebind.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/server/vrfs/vrf/statistics/" + rebind.EntityData.SegmentPath
    rebind.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rebind.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rebind.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rebind.EntityData.Children = types.NewOrderedMap()
    rebind.EntityData.Leafs = types.NewOrderedMap()
    rebind.EntityData.Leafs.Append("received-packets", types.YLeaf{"ReceivedPackets", rebind.ReceivedPackets})
    rebind.EntityData.Leafs.Append("transmitted-packets", types.YLeaf{"TransmittedPackets", rebind.TransmittedPackets})
    rebind.EntityData.Leafs.Append("dropped-packets", types.YLeaf{"DroppedPackets", rebind.DroppedPackets})

    rebind.EntityData.YListKeys = []string {}

    return &(rebind.EntityData)
}

// Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Release
// DHCPV6 release packets
type Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Release struct {
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

func (release *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Release) GetEntityData() *types.CommonEntityData {
    release.EntityData.YFilter = release.YFilter
    release.EntityData.YangName = "release"
    release.EntityData.BundleName = "cisco_ios_xr"
    release.EntityData.ParentYangName = "statistics"
    release.EntityData.SegmentPath = "release"
    release.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/server/vrfs/vrf/statistics/" + release.EntityData.SegmentPath
    release.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    release.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    release.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    release.EntityData.Children = types.NewOrderedMap()
    release.EntityData.Leafs = types.NewOrderedMap()
    release.EntityData.Leafs.Append("received-packets", types.YLeaf{"ReceivedPackets", release.ReceivedPackets})
    release.EntityData.Leafs.Append("transmitted-packets", types.YLeaf{"TransmittedPackets", release.TransmittedPackets})
    release.EntityData.Leafs.Append("dropped-packets", types.YLeaf{"DroppedPackets", release.DroppedPackets})

    release.EntityData.YListKeys = []string {}

    return &(release.EntityData)
}

// Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Reconfig
// DHCPV6 reconfig packets
type Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Reconfig struct {
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

func (reconfig *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Reconfig) GetEntityData() *types.CommonEntityData {
    reconfig.EntityData.YFilter = reconfig.YFilter
    reconfig.EntityData.YangName = "reconfig"
    reconfig.EntityData.BundleName = "cisco_ios_xr"
    reconfig.EntityData.ParentYangName = "statistics"
    reconfig.EntityData.SegmentPath = "reconfig"
    reconfig.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/server/vrfs/vrf/statistics/" + reconfig.EntityData.SegmentPath
    reconfig.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reconfig.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reconfig.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reconfig.EntityData.Children = types.NewOrderedMap()
    reconfig.EntityData.Leafs = types.NewOrderedMap()
    reconfig.EntityData.Leafs.Append("received-packets", types.YLeaf{"ReceivedPackets", reconfig.ReceivedPackets})
    reconfig.EntityData.Leafs.Append("transmitted-packets", types.YLeaf{"TransmittedPackets", reconfig.TransmittedPackets})
    reconfig.EntityData.Leafs.Append("dropped-packets", types.YLeaf{"DroppedPackets", reconfig.DroppedPackets})

    reconfig.EntityData.YListKeys = []string {}

    return &(reconfig.EntityData)
}

// Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Inform
// DHCPV6 inform packets
type Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Inform struct {
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

func (inform *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_Inform) GetEntityData() *types.CommonEntityData {
    inform.EntityData.YFilter = inform.YFilter
    inform.EntityData.YangName = "inform"
    inform.EntityData.BundleName = "cisco_ios_xr"
    inform.EntityData.ParentYangName = "statistics"
    inform.EntityData.SegmentPath = "inform"
    inform.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/server/vrfs/vrf/statistics/" + inform.EntityData.SegmentPath
    inform.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inform.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inform.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inform.EntityData.Children = types.NewOrderedMap()
    inform.EntityData.Leafs = types.NewOrderedMap()
    inform.EntityData.Leafs.Append("received-packets", types.YLeaf{"ReceivedPackets", inform.ReceivedPackets})
    inform.EntityData.Leafs.Append("transmitted-packets", types.YLeaf{"TransmittedPackets", inform.TransmittedPackets})
    inform.EntityData.Leafs.Append("dropped-packets", types.YLeaf{"DroppedPackets", inform.DroppedPackets})

    inform.EntityData.YListKeys = []string {}

    return &(inform.EntityData)
}

// Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_RelayForward
// DHCPV6 relay forward packets
type Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_RelayForward struct {
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

func (relayForward *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_RelayForward) GetEntityData() *types.CommonEntityData {
    relayForward.EntityData.YFilter = relayForward.YFilter
    relayForward.EntityData.YangName = "relay-forward"
    relayForward.EntityData.BundleName = "cisco_ios_xr"
    relayForward.EntityData.ParentYangName = "statistics"
    relayForward.EntityData.SegmentPath = "relay-forward"
    relayForward.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/server/vrfs/vrf/statistics/" + relayForward.EntityData.SegmentPath
    relayForward.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    relayForward.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    relayForward.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    relayForward.EntityData.Children = types.NewOrderedMap()
    relayForward.EntityData.Leafs = types.NewOrderedMap()
    relayForward.EntityData.Leafs.Append("received-packets", types.YLeaf{"ReceivedPackets", relayForward.ReceivedPackets})
    relayForward.EntityData.Leafs.Append("transmitted-packets", types.YLeaf{"TransmittedPackets", relayForward.TransmittedPackets})
    relayForward.EntityData.Leafs.Append("dropped-packets", types.YLeaf{"DroppedPackets", relayForward.DroppedPackets})

    relayForward.EntityData.YListKeys = []string {}

    return &(relayForward.EntityData)
}

// Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_RelayReply
// DHCPV6 relay reply packets
type Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_RelayReply struct {
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

func (relayReply *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_RelayReply) GetEntityData() *types.CommonEntityData {
    relayReply.EntityData.YFilter = relayReply.YFilter
    relayReply.EntityData.YangName = "relay-reply"
    relayReply.EntityData.BundleName = "cisco_ios_xr"
    relayReply.EntityData.ParentYangName = "statistics"
    relayReply.EntityData.SegmentPath = "relay-reply"
    relayReply.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/server/vrfs/vrf/statistics/" + relayReply.EntityData.SegmentPath
    relayReply.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    relayReply.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    relayReply.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    relayReply.EntityData.Children = types.NewOrderedMap()
    relayReply.EntityData.Leafs = types.NewOrderedMap()
    relayReply.EntityData.Leafs.Append("received-packets", types.YLeaf{"ReceivedPackets", relayReply.ReceivedPackets})
    relayReply.EntityData.Leafs.Append("transmitted-packets", types.YLeaf{"TransmittedPackets", relayReply.TransmittedPackets})
    relayReply.EntityData.Leafs.Append("dropped-packets", types.YLeaf{"DroppedPackets", relayReply.DroppedPackets})

    relayReply.EntityData.YListKeys = []string {}

    return &(relayReply.EntityData)
}

// Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQuery
// DHCPV6 lease query packets
type Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQuery struct {
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

func (leaseQuery *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQuery) GetEntityData() *types.CommonEntityData {
    leaseQuery.EntityData.YFilter = leaseQuery.YFilter
    leaseQuery.EntityData.YangName = "lease-query"
    leaseQuery.EntityData.BundleName = "cisco_ios_xr"
    leaseQuery.EntityData.ParentYangName = "statistics"
    leaseQuery.EntityData.SegmentPath = "lease-query"
    leaseQuery.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/server/vrfs/vrf/statistics/" + leaseQuery.EntityData.SegmentPath
    leaseQuery.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    leaseQuery.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    leaseQuery.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    leaseQuery.EntityData.Children = types.NewOrderedMap()
    leaseQuery.EntityData.Leafs = types.NewOrderedMap()
    leaseQuery.EntityData.Leafs.Append("received-packets", types.YLeaf{"ReceivedPackets", leaseQuery.ReceivedPackets})
    leaseQuery.EntityData.Leafs.Append("transmitted-packets", types.YLeaf{"TransmittedPackets", leaseQuery.TransmittedPackets})
    leaseQuery.EntityData.Leafs.Append("dropped-packets", types.YLeaf{"DroppedPackets", leaseQuery.DroppedPackets})

    leaseQuery.EntityData.YListKeys = []string {}

    return &(leaseQuery.EntityData)
}

// Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQueryReply
// DHCPV6 lease query reply packets
type Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQueryReply struct {
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

func (leaseQueryReply *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQueryReply) GetEntityData() *types.CommonEntityData {
    leaseQueryReply.EntityData.YFilter = leaseQueryReply.YFilter
    leaseQueryReply.EntityData.YangName = "lease-query-reply"
    leaseQueryReply.EntityData.BundleName = "cisco_ios_xr"
    leaseQueryReply.EntityData.ParentYangName = "statistics"
    leaseQueryReply.EntityData.SegmentPath = "lease-query-reply"
    leaseQueryReply.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/server/vrfs/vrf/statistics/" + leaseQueryReply.EntityData.SegmentPath
    leaseQueryReply.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    leaseQueryReply.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    leaseQueryReply.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    leaseQueryReply.EntityData.Children = types.NewOrderedMap()
    leaseQueryReply.EntityData.Leafs = types.NewOrderedMap()
    leaseQueryReply.EntityData.Leafs.Append("received-packets", types.YLeaf{"ReceivedPackets", leaseQueryReply.ReceivedPackets})
    leaseQueryReply.EntityData.Leafs.Append("transmitted-packets", types.YLeaf{"TransmittedPackets", leaseQueryReply.TransmittedPackets})
    leaseQueryReply.EntityData.Leafs.Append("dropped-packets", types.YLeaf{"DroppedPackets", leaseQueryReply.DroppedPackets})

    leaseQueryReply.EntityData.YListKeys = []string {}

    return &(leaseQueryReply.EntityData)
}

// Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQueryDone
// DHCPV6 lease query done packets
type Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQueryDone struct {
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

func (leaseQueryDone *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQueryDone) GetEntityData() *types.CommonEntityData {
    leaseQueryDone.EntityData.YFilter = leaseQueryDone.YFilter
    leaseQueryDone.EntityData.YangName = "lease-query-done"
    leaseQueryDone.EntityData.BundleName = "cisco_ios_xr"
    leaseQueryDone.EntityData.ParentYangName = "statistics"
    leaseQueryDone.EntityData.SegmentPath = "lease-query-done"
    leaseQueryDone.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/server/vrfs/vrf/statistics/" + leaseQueryDone.EntityData.SegmentPath
    leaseQueryDone.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    leaseQueryDone.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    leaseQueryDone.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    leaseQueryDone.EntityData.Children = types.NewOrderedMap()
    leaseQueryDone.EntityData.Leafs = types.NewOrderedMap()
    leaseQueryDone.EntityData.Leafs.Append("received-packets", types.YLeaf{"ReceivedPackets", leaseQueryDone.ReceivedPackets})
    leaseQueryDone.EntityData.Leafs.Append("transmitted-packets", types.YLeaf{"TransmittedPackets", leaseQueryDone.TransmittedPackets})
    leaseQueryDone.EntityData.Leafs.Append("dropped-packets", types.YLeaf{"DroppedPackets", leaseQueryDone.DroppedPackets})

    leaseQueryDone.EntityData.YListKeys = []string {}

    return &(leaseQueryDone.EntityData)
}

// Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQueryData
// DHCPV6 lease query data packets
type Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQueryData struct {
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

func (leaseQueryData *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics_LeaseQueryData) GetEntityData() *types.CommonEntityData {
    leaseQueryData.EntityData.YFilter = leaseQueryData.YFilter
    leaseQueryData.EntityData.YangName = "lease-query-data"
    leaseQueryData.EntityData.BundleName = "cisco_ios_xr"
    leaseQueryData.EntityData.ParentYangName = "statistics"
    leaseQueryData.EntityData.SegmentPath = "lease-query-data"
    leaseQueryData.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/server/vrfs/vrf/statistics/" + leaseQueryData.EntityData.SegmentPath
    leaseQueryData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    leaseQueryData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    leaseQueryData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    leaseQueryData.EntityData.Children = types.NewOrderedMap()
    leaseQueryData.EntityData.Leafs = types.NewOrderedMap()
    leaseQueryData.EntityData.Leafs.Append("received-packets", types.YLeaf{"ReceivedPackets", leaseQueryData.ReceivedPackets})
    leaseQueryData.EntityData.Leafs.Append("transmitted-packets", types.YLeaf{"TransmittedPackets", leaseQueryData.TransmittedPackets})
    leaseQueryData.EntityData.Leafs.Append("dropped-packets", types.YLeaf{"DroppedPackets", leaseQueryData.DroppedPackets})

    leaseQueryData.EntityData.YListKeys = []string {}

    return &(leaseQueryData.EntityData)
}

// Dhcpv6_Nodes_Node_Server_Profiles
// IPv6 DHCP server profile
type Dhcpv6_Nodes_Node_Server_Profiles struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 DHCP server profile. The type is slice of
    // Dhcpv6_Nodes_Node_Server_Profiles_Profile.
    Profile []*Dhcpv6_Nodes_Node_Server_Profiles_Profile
}

func (profiles *Dhcpv6_Nodes_Node_Server_Profiles) GetEntityData() *types.CommonEntityData {
    profiles.EntityData.YFilter = profiles.YFilter
    profiles.EntityData.YangName = "profiles"
    profiles.EntityData.BundleName = "cisco_ios_xr"
    profiles.EntityData.ParentYangName = "server"
    profiles.EntityData.SegmentPath = "profiles"
    profiles.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/server/" + profiles.EntityData.SegmentPath
    profiles.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    profiles.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    profiles.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    profiles.EntityData.Children = types.NewOrderedMap()
    profiles.EntityData.Children.Append("profile", types.YChild{"Profile", nil})
    for i := range profiles.Profile {
        profiles.EntityData.Children.Append(types.GetSegmentPath(profiles.Profile[i]), types.YChild{"Profile", profiles.Profile[i]})
    }
    profiles.EntityData.Leafs = types.NewOrderedMap()

    profiles.EntityData.YListKeys = []string {}

    return &(profiles.EntityData)
}

// Dhcpv6_Nodes_Node_Server_Profiles_Profile
// IPv6 DHCP server profile
type Dhcpv6_Nodes_Node_Server_Profiles_Profile struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Profile name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    ProfileName interface{}

    // IPv6 DHCP server profile Info.
    Info Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info

    // DHCPV6 throttle table.
    ThrottleInfos Dhcpv6_Nodes_Node_Server_Profiles_Profile_ThrottleInfos

    // IPv6 DHCP server class table.
    ServerClasses Dhcpv6_Nodes_Node_Server_Profiles_Profile_ServerClasses
}

func (profile *Dhcpv6_Nodes_Node_Server_Profiles_Profile) GetEntityData() *types.CommonEntityData {
    profile.EntityData.YFilter = profile.YFilter
    profile.EntityData.YangName = "profile"
    profile.EntityData.BundleName = "cisco_ios_xr"
    profile.EntityData.ParentYangName = "profiles"
    profile.EntityData.SegmentPath = "profile" + types.AddKeyToken(profile.ProfileName, "profile-name")
    profile.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/server/profiles/" + profile.EntityData.SegmentPath
    profile.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    profile.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    profile.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    profile.EntityData.Children = types.NewOrderedMap()
    profile.EntityData.Children.Append("info", types.YChild{"Info", &profile.Info})
    profile.EntityData.Children.Append("throttle-infos", types.YChild{"ThrottleInfos", &profile.ThrottleInfos})
    profile.EntityData.Children.Append("server-classes", types.YChild{"ServerClasses", &profile.ServerClasses})
    profile.EntityData.Leafs = types.NewOrderedMap()
    profile.EntityData.Leafs.Append("profile-name", types.YLeaf{"ProfileName", profile.ProfileName})

    profile.EntityData.YListKeys = []string {"ProfileName"}

    return &(profile.EntityData)
}

// Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info
// IPv6 DHCP server profile Info
type Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Server profile name. The type is string with length: 0..65.
    ProfileName interface{}

    // Allowed DUID Type. The type is interface{} with range: 0..65535.
    ProfileAllowedDuidType interface{}

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

func (info *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info) GetEntityData() *types.CommonEntityData {
    info.EntityData.YFilter = info.YFilter
    info.EntityData.YangName = "info"
    info.EntityData.BundleName = "cisco_ios_xr"
    info.EntityData.ParentYangName = "profile"
    info.EntityData.SegmentPath = "info"
    info.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/server/profiles/profile/" + info.EntityData.SegmentPath
    info.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    info.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    info.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    info.EntityData.Children = types.NewOrderedMap()
    info.EntityData.Children.Append("lease", types.YChild{"Lease", &info.Lease})
    info.EntityData.Children.Append("interface-references", types.YChild{"InterfaceReferences", &info.InterfaceReferences})
    info.EntityData.Leafs = types.NewOrderedMap()
    info.EntityData.Leafs.Append("profile-name", types.YLeaf{"ProfileName", info.ProfileName})
    info.EntityData.Leafs.Append("profile-allowed-duid-type", types.YLeaf{"ProfileAllowedDuidType", info.ProfileAllowedDuidType})
    info.EntityData.Leafs.Append("domain-name", types.YLeaf{"DomainName", info.DomainName})
    info.EntityData.Leafs.Append("profile-dns", types.YLeaf{"ProfileDns", info.ProfileDns})
    info.EntityData.Leafs.Append("aftr-name", types.YLeaf{"AftrName", info.AftrName})
    info.EntityData.Leafs.Append("framed-addr-pool-name", types.YLeaf{"FramedAddrPoolName", info.FramedAddrPoolName})
    info.EntityData.Leafs.Append("delegated-prefix-pool-name", types.YLeaf{"DelegatedPrefixPoolName", info.DelegatedPrefixPoolName})
    info.EntityData.Leafs.Append("rapid-commit", types.YLeaf{"RapidCommit", info.RapidCommit})
    info.EntityData.Leafs.Append("profile-dns-address", types.YLeaf{"ProfileDnsAddress", info.ProfileDnsAddress})

    info.EntityData.YListKeys = []string {}

    return &(info.EntityData)
}

// Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_Lease
// Server lease time
type Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_Lease struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DHCPV6 client lease in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    Seconds interface{}

    // Time in format HH:MM:SS. The type is string with length: 0..10.
    Time interface{}
}

func (lease *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_Lease) GetEntityData() *types.CommonEntityData {
    lease.EntityData.YFilter = lease.YFilter
    lease.EntityData.YangName = "lease"
    lease.EntityData.BundleName = "cisco_ios_xr"
    lease.EntityData.ParentYangName = "info"
    lease.EntityData.SegmentPath = "lease"
    lease.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/server/profiles/profile/info/" + lease.EntityData.SegmentPath
    lease.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lease.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lease.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lease.EntityData.Children = types.NewOrderedMap()
    lease.EntityData.Leafs = types.NewOrderedMap()
    lease.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", lease.Seconds})
    lease.EntityData.Leafs.Append("time", types.YLeaf{"Time", lease.Time})

    lease.EntityData.YListKeys = []string {}

    return &(lease.EntityData)
}

// Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_InterfaceReferences
// Interface references
type Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_InterfaceReferences struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ipv6 dhcpv6d server interface reference. The type is slice of
    // Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_InterfaceReferences_Ipv6Dhcpv6dServerInterfaceReference.
    Ipv6Dhcpv6dServerInterfaceReference []*Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_InterfaceReferences_Ipv6Dhcpv6dServerInterfaceReference
}

func (interfaceReferences *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_InterfaceReferences) GetEntityData() *types.CommonEntityData {
    interfaceReferences.EntityData.YFilter = interfaceReferences.YFilter
    interfaceReferences.EntityData.YangName = "interface-references"
    interfaceReferences.EntityData.BundleName = "cisco_ios_xr"
    interfaceReferences.EntityData.ParentYangName = "info"
    interfaceReferences.EntityData.SegmentPath = "interface-references"
    interfaceReferences.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/server/profiles/profile/info/" + interfaceReferences.EntityData.SegmentPath
    interfaceReferences.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceReferences.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceReferences.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceReferences.EntityData.Children = types.NewOrderedMap()
    interfaceReferences.EntityData.Children.Append("ipv6-dhcpv6d-server-interface-reference", types.YChild{"Ipv6Dhcpv6dServerInterfaceReference", nil})
    for i := range interfaceReferences.Ipv6Dhcpv6dServerInterfaceReference {
        types.SetYListKey(interfaceReferences.Ipv6Dhcpv6dServerInterfaceReference[i], i)
        interfaceReferences.EntityData.Children.Append(types.GetSegmentPath(interfaceReferences.Ipv6Dhcpv6dServerInterfaceReference[i]), types.YChild{"Ipv6Dhcpv6dServerInterfaceReference", interfaceReferences.Ipv6Dhcpv6dServerInterfaceReference[i]})
    }
    interfaceReferences.EntityData.Leafs = types.NewOrderedMap()

    interfaceReferences.EntityData.YListKeys = []string {}

    return &(interfaceReferences.EntityData)
}

// Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_InterfaceReferences_Ipv6Dhcpv6dServerInterfaceReference
// ipv6 dhcpv6d server interface reference
type Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_InterfaceReferences_Ipv6Dhcpv6dServerInterfaceReference struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Interface name. The type is string with length: 0..65.
    ServerReferenceInterfaceName interface{}
}

func (ipv6Dhcpv6dServerInterfaceReference *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_InterfaceReferences_Ipv6Dhcpv6dServerInterfaceReference) GetEntityData() *types.CommonEntityData {
    ipv6Dhcpv6dServerInterfaceReference.EntityData.YFilter = ipv6Dhcpv6dServerInterfaceReference.YFilter
    ipv6Dhcpv6dServerInterfaceReference.EntityData.YangName = "ipv6-dhcpv6d-server-interface-reference"
    ipv6Dhcpv6dServerInterfaceReference.EntityData.BundleName = "cisco_ios_xr"
    ipv6Dhcpv6dServerInterfaceReference.EntityData.ParentYangName = "interface-references"
    ipv6Dhcpv6dServerInterfaceReference.EntityData.SegmentPath = "ipv6-dhcpv6d-server-interface-reference" + types.AddNoKeyToken(ipv6Dhcpv6dServerInterfaceReference)
    ipv6Dhcpv6dServerInterfaceReference.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/server/profiles/profile/info/interface-references/" + ipv6Dhcpv6dServerInterfaceReference.EntityData.SegmentPath
    ipv6Dhcpv6dServerInterfaceReference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Dhcpv6dServerInterfaceReference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Dhcpv6dServerInterfaceReference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Dhcpv6dServerInterfaceReference.EntityData.Children = types.NewOrderedMap()
    ipv6Dhcpv6dServerInterfaceReference.EntityData.Leafs = types.NewOrderedMap()
    ipv6Dhcpv6dServerInterfaceReference.EntityData.Leafs.Append("server-reference-interface-name", types.YLeaf{"ServerReferenceInterfaceName", ipv6Dhcpv6dServerInterfaceReference.ServerReferenceInterfaceName})

    ipv6Dhcpv6dServerInterfaceReference.EntityData.YListKeys = []string {}

    return &(ipv6Dhcpv6dServerInterfaceReference.EntityData)
}

// Dhcpv6_Nodes_Node_Server_Profiles_Profile_ThrottleInfos
// DHCPV6 throttle table
type Dhcpv6_Nodes_Node_Server_Profiles_Profile_ThrottleInfos struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 DHCP server profile Throttle Info. The type is slice of
    // Dhcpv6_Nodes_Node_Server_Profiles_Profile_ThrottleInfos_ThrottleInfo.
    ThrottleInfo []*Dhcpv6_Nodes_Node_Server_Profiles_Profile_ThrottleInfos_ThrottleInfo
}

func (throttleInfos *Dhcpv6_Nodes_Node_Server_Profiles_Profile_ThrottleInfos) GetEntityData() *types.CommonEntityData {
    throttleInfos.EntityData.YFilter = throttleInfos.YFilter
    throttleInfos.EntityData.YangName = "throttle-infos"
    throttleInfos.EntityData.BundleName = "cisco_ios_xr"
    throttleInfos.EntityData.ParentYangName = "profile"
    throttleInfos.EntityData.SegmentPath = "throttle-infos"
    throttleInfos.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/server/profiles/profile/" + throttleInfos.EntityData.SegmentPath
    throttleInfos.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    throttleInfos.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    throttleInfos.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    throttleInfos.EntityData.Children = types.NewOrderedMap()
    throttleInfos.EntityData.Children.Append("throttle-info", types.YChild{"ThrottleInfo", nil})
    for i := range throttleInfos.ThrottleInfo {
        throttleInfos.EntityData.Children.Append(types.GetSegmentPath(throttleInfos.ThrottleInfo[i]), types.YChild{"ThrottleInfo", throttleInfos.ThrottleInfo[i]})
    }
    throttleInfos.EntityData.Leafs = types.NewOrderedMap()

    throttleInfos.EntityData.YListKeys = []string {}

    return &(throttleInfos.EntityData)
}

// Dhcpv6_Nodes_Node_Server_Profiles_Profile_ThrottleInfos_ThrottleInfo
// IPv6 DHCP server profile Throttle Info
type Dhcpv6_Nodes_Node_Server_Profiles_Profile_ThrottleInfos_ThrottleInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (throttleInfo *Dhcpv6_Nodes_Node_Server_Profiles_Profile_ThrottleInfos_ThrottleInfo) GetEntityData() *types.CommonEntityData {
    throttleInfo.EntityData.YFilter = throttleInfo.YFilter
    throttleInfo.EntityData.YangName = "throttle-info"
    throttleInfo.EntityData.BundleName = "cisco_ios_xr"
    throttleInfo.EntityData.ParentYangName = "throttle-infos"
    throttleInfo.EntityData.SegmentPath = "throttle-info" + types.AddKeyToken(throttleInfo.MacAddress, "mac-address")
    throttleInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/server/profiles/profile/throttle-infos/" + throttleInfo.EntityData.SegmentPath
    throttleInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    throttleInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    throttleInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    throttleInfo.EntityData.Children = types.NewOrderedMap()
    throttleInfo.EntityData.Leafs = types.NewOrderedMap()
    throttleInfo.EntityData.Leafs.Append("mac-address", types.YLeaf{"MacAddress", throttleInfo.MacAddress})
    throttleInfo.EntityData.Leafs.Append("binding-chaddr", types.YLeaf{"BindingChaddr", throttleInfo.BindingChaddr})
    throttleInfo.EntityData.Leafs.Append("ifname", types.YLeaf{"Ifname", throttleInfo.Ifname})
    throttleInfo.EntityData.Leafs.Append("state", types.YLeaf{"State", throttleInfo.State})
    throttleInfo.EntityData.Leafs.Append("time-left", types.YLeaf{"TimeLeft", throttleInfo.TimeLeft})

    throttleInfo.EntityData.YListKeys = []string {"MacAddress"}

    return &(throttleInfo.EntityData)
}

// Dhcpv6_Nodes_Node_Server_Profiles_Profile_ServerClasses
// IPv6 DHCP server class table
type Dhcpv6_Nodes_Node_Server_Profiles_Profile_ServerClasses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 DHCP server class profile. The type is slice of
    // Dhcpv6_Nodes_Node_Server_Profiles_Profile_ServerClasses_ServerClass.
    ServerClass []*Dhcpv6_Nodes_Node_Server_Profiles_Profile_ServerClasses_ServerClass
}

func (serverClasses *Dhcpv6_Nodes_Node_Server_Profiles_Profile_ServerClasses) GetEntityData() *types.CommonEntityData {
    serverClasses.EntityData.YFilter = serverClasses.YFilter
    serverClasses.EntityData.YangName = "server-classes"
    serverClasses.EntityData.BundleName = "cisco_ios_xr"
    serverClasses.EntityData.ParentYangName = "profile"
    serverClasses.EntityData.SegmentPath = "server-classes"
    serverClasses.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/server/profiles/profile/" + serverClasses.EntityData.SegmentPath
    serverClasses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serverClasses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serverClasses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serverClasses.EntityData.Children = types.NewOrderedMap()
    serverClasses.EntityData.Children.Append("server-class", types.YChild{"ServerClass", nil})
    for i := range serverClasses.ServerClass {
        serverClasses.EntityData.Children.Append(types.GetSegmentPath(serverClasses.ServerClass[i]), types.YChild{"ServerClass", serverClasses.ServerClass[i]})
    }
    serverClasses.EntityData.Leafs = types.NewOrderedMap()

    serverClasses.EntityData.YListKeys = []string {}

    return &(serverClasses.EntityData)
}

// Dhcpv6_Nodes_Node_Server_Profiles_Profile_ServerClasses_ServerClass
// IPv6 DHCP server class profile
type Dhcpv6_Nodes_Node_Server_Profiles_Profile_ServerClasses_ServerClass struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Class name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    ClassName interface{}

    // Class name. The type is string with length: 0..65.
    ClassNameXr interface{}

    // Server domain name. The type is string with length: 0..65.
    DomainName interface{}

    // DNS address count. The type is interface{} with range: 0..255.
    ProfileDns interface{}

    // Server framed address pool name. The type is string with length: 0..65.
    FramedAddrPoolName interface{}

    // Server delegated prefix pool name. The type is string with length: 0..65.
    DelegatedPrefixPoolName interface{}

    // DNS addresses. The type is slice of string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    ProfileDnsAddress []interface{}
}

func (serverClass *Dhcpv6_Nodes_Node_Server_Profiles_Profile_ServerClasses_ServerClass) GetEntityData() *types.CommonEntityData {
    serverClass.EntityData.YFilter = serverClass.YFilter
    serverClass.EntityData.YangName = "server-class"
    serverClass.EntityData.BundleName = "cisco_ios_xr"
    serverClass.EntityData.ParentYangName = "server-classes"
    serverClass.EntityData.SegmentPath = "server-class" + types.AddKeyToken(serverClass.ClassName, "class-name")
    serverClass.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/server/profiles/profile/server-classes/" + serverClass.EntityData.SegmentPath
    serverClass.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serverClass.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serverClass.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serverClass.EntityData.Children = types.NewOrderedMap()
    serverClass.EntityData.Leafs = types.NewOrderedMap()
    serverClass.EntityData.Leafs.Append("class-name", types.YLeaf{"ClassName", serverClass.ClassName})
    serverClass.EntityData.Leafs.Append("class-name-xr", types.YLeaf{"ClassNameXr", serverClass.ClassNameXr})
    serverClass.EntityData.Leafs.Append("domain-name", types.YLeaf{"DomainName", serverClass.DomainName})
    serverClass.EntityData.Leafs.Append("profile-dns", types.YLeaf{"ProfileDns", serverClass.ProfileDns})
    serverClass.EntityData.Leafs.Append("framed-addr-pool-name", types.YLeaf{"FramedAddrPoolName", serverClass.FramedAddrPoolName})
    serverClass.EntityData.Leafs.Append("delegated-prefix-pool-name", types.YLeaf{"DelegatedPrefixPoolName", serverClass.DelegatedPrefixPoolName})
    serverClass.EntityData.Leafs.Append("profile-dns-address", types.YLeaf{"ProfileDnsAddress", serverClass.ProfileDnsAddress})

    serverClass.EntityData.YListKeys = []string {"ClassName"}

    return &(serverClass.EntityData)
}

// Dhcpv6_Nodes_Node_Server_Interfaces
// DHCPV6 server interface
type Dhcpv6_Nodes_Node_Server_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 DHCP server interface. The type is slice of
    // Dhcpv6_Nodes_Node_Server_Interfaces_Interface.
    Interface []*Dhcpv6_Nodes_Node_Server_Interfaces_Interface
}

func (interfaces *Dhcpv6_Nodes_Node_Server_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "server"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/server/" + interfaces.EntityData.SegmentPath
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

// Dhcpv6_Nodes_Node_Server_Interfaces_Interface
// IPv6 DHCP server interface
type Dhcpv6_Nodes_Node_Server_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    InterfaceName interface{}

    // VRF name. The type is string with length: 0..33.
    ServerVrfName interface{}

    // Mode of interface. The type is BagDhcpv6dSubMode.
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

    // DHCPv6 Interface SRG role. The type is BagDhcpv6dIntfSrgRole.
    SrgRole interface{}

    // DHCPv6 Interface SERG role. The type is BagDhcpv6dIntfSergRole.
    SergRole interface{}

    // Mac Throttle Status. The type is bool.
    MacThrottle interface{}

    // SRG VRF name. The type is string with length: 0..33.
    SrgVrfName interface{}

    // SRG P2P Status. The type is bool.
    Srgp2p interface{}
}

func (self *Dhcpv6_Nodes_Node_Server_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceName, "interface-name")
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/server/interfaces/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})
    self.EntityData.Leafs.Append("server-vrf-name", types.YLeaf{"ServerVrfName", self.ServerVrfName})
    self.EntityData.Leafs.Append("server-interface-mode", types.YLeaf{"ServerInterfaceMode", self.ServerInterfaceMode})
    self.EntityData.Leafs.Append("is-server-interface-ambiguous", types.YLeaf{"IsServerInterfaceAmbiguous", self.IsServerInterfaceAmbiguous})
    self.EntityData.Leafs.Append("server-interface-profile-name", types.YLeaf{"ServerInterfaceProfileName", self.ServerInterfaceProfileName})
    self.EntityData.Leafs.Append("server-interface-lease-limit-type", types.YLeaf{"ServerInterfaceLeaseLimitType", self.ServerInterfaceLeaseLimitType})
    self.EntityData.Leafs.Append("server-interface-lease-limits", types.YLeaf{"ServerInterfaceLeaseLimits", self.ServerInterfaceLeaseLimits})
    self.EntityData.Leafs.Append("srg-role", types.YLeaf{"SrgRole", self.SrgRole})
    self.EntityData.Leafs.Append("serg-role", types.YLeaf{"SergRole", self.SergRole})
    self.EntityData.Leafs.Append("mac-throttle", types.YLeaf{"MacThrottle", self.MacThrottle})
    self.EntityData.Leafs.Append("srg-vrf-name", types.YLeaf{"SrgVrfName", self.SrgVrfName})
    self.EntityData.Leafs.Append("srgp2p", types.YLeaf{"Srgp2p", self.Srgp2p})

    self.EntityData.YListKeys = []string {"InterfaceName"}

    return &(self.EntityData)
}

// Dhcpv6_Nodes_Node_Server_Statistics
// DHCPv6 server statistics
type Dhcpv6_Nodes_Node_Server_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ipv6 dhcpv6d server stat. The type is slice of
    // Dhcpv6_Nodes_Node_Server_Statistics_Ipv6Dhcpv6dServerStat.
    Ipv6Dhcpv6dServerStat []*Dhcpv6_Nodes_Node_Server_Statistics_Ipv6Dhcpv6dServerStat
}

func (statistics *Dhcpv6_Nodes_Node_Server_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "server"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/server/" + statistics.EntityData.SegmentPath
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = types.NewOrderedMap()
    statistics.EntityData.Children.Append("ipv6-dhcpv6d-server-stat", types.YChild{"Ipv6Dhcpv6dServerStat", nil})
    for i := range statistics.Ipv6Dhcpv6dServerStat {
        types.SetYListKey(statistics.Ipv6Dhcpv6dServerStat[i], i)
        statistics.EntityData.Children.Append(types.GetSegmentPath(statistics.Ipv6Dhcpv6dServerStat[i]), types.YChild{"Ipv6Dhcpv6dServerStat", statistics.Ipv6Dhcpv6dServerStat[i]})
    }
    statistics.EntityData.Leafs = types.NewOrderedMap()

    statistics.EntityData.YListKeys = []string {}

    return &(statistics.EntityData)
}

// Dhcpv6_Nodes_Node_Server_Statistics_Ipv6Dhcpv6dServerStat
// ipv6 dhcpv6d server stat
type Dhcpv6_Nodes_Node_Server_Statistics_Ipv6Dhcpv6dServerStat struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // DHCPv6 L3 VRF name. The type is string with length: 0..33.
    VrfName interface{}

    // Server statistics.
    Statistics Dhcpv6_Nodes_Node_Server_Statistics_Ipv6Dhcpv6dServerStat_Statistics
}

func (ipv6Dhcpv6dServerStat *Dhcpv6_Nodes_Node_Server_Statistics_Ipv6Dhcpv6dServerStat) GetEntityData() *types.CommonEntityData {
    ipv6Dhcpv6dServerStat.EntityData.YFilter = ipv6Dhcpv6dServerStat.YFilter
    ipv6Dhcpv6dServerStat.EntityData.YangName = "ipv6-dhcpv6d-server-stat"
    ipv6Dhcpv6dServerStat.EntityData.BundleName = "cisco_ios_xr"
    ipv6Dhcpv6dServerStat.EntityData.ParentYangName = "statistics"
    ipv6Dhcpv6dServerStat.EntityData.SegmentPath = "ipv6-dhcpv6d-server-stat" + types.AddNoKeyToken(ipv6Dhcpv6dServerStat)
    ipv6Dhcpv6dServerStat.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/server/statistics/" + ipv6Dhcpv6dServerStat.EntityData.SegmentPath
    ipv6Dhcpv6dServerStat.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Dhcpv6dServerStat.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Dhcpv6dServerStat.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Dhcpv6dServerStat.EntityData.Children = types.NewOrderedMap()
    ipv6Dhcpv6dServerStat.EntityData.Children.Append("statistics", types.YChild{"Statistics", &ipv6Dhcpv6dServerStat.Statistics})
    ipv6Dhcpv6dServerStat.EntityData.Leafs = types.NewOrderedMap()
    ipv6Dhcpv6dServerStat.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", ipv6Dhcpv6dServerStat.VrfName})

    ipv6Dhcpv6dServerStat.EntityData.YListKeys = []string {}

    return &(ipv6Dhcpv6dServerStat.EntityData)
}

// Dhcpv6_Nodes_Node_Server_Statistics_Ipv6Dhcpv6dServerStat_Statistics
// Server statistics
type Dhcpv6_Nodes_Node_Server_Statistics_Ipv6Dhcpv6dServerStat_Statistics struct {
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

func (statistics *Dhcpv6_Nodes_Node_Server_Statistics_Ipv6Dhcpv6dServerStat_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "ipv6-dhcpv6d-server-stat"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/server/statistics/ipv6-dhcpv6d-server-stat/" + statistics.EntityData.SegmentPath
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = types.NewOrderedMap()
    statistics.EntityData.Leafs = types.NewOrderedMap()
    statistics.EntityData.Leafs.Append("received-packets", types.YLeaf{"ReceivedPackets", statistics.ReceivedPackets})
    statistics.EntityData.Leafs.Append("transmitted-packets", types.YLeaf{"TransmittedPackets", statistics.TransmittedPackets})
    statistics.EntityData.Leafs.Append("dropped-packets", types.YLeaf{"DroppedPackets", statistics.DroppedPackets})

    statistics.EntityData.YListKeys = []string {}

    return &(statistics.EntityData)
}

// Dhcpv6_Nodes_Node_Server_BindingOptions
// DHCPv6 server binding with options
type Dhcpv6_Nodes_Node_Server_BindingOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DHCPv6 server binding from MAC address.
    MacBindOptions Dhcpv6_Nodes_Node_Server_BindingOptions_MacBindOptions

    // DHCPv6 server binding from DUID.
    DuidBindOptions Dhcpv6_Nodes_Node_Server_BindingOptions_DuidBindOptions
}

func (bindingOptions *Dhcpv6_Nodes_Node_Server_BindingOptions) GetEntityData() *types.CommonEntityData {
    bindingOptions.EntityData.YFilter = bindingOptions.YFilter
    bindingOptions.EntityData.YangName = "binding-options"
    bindingOptions.EntityData.BundleName = "cisco_ios_xr"
    bindingOptions.EntityData.ParentYangName = "server"
    bindingOptions.EntityData.SegmentPath = "binding-options"
    bindingOptions.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/server/" + bindingOptions.EntityData.SegmentPath
    bindingOptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bindingOptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bindingOptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bindingOptions.EntityData.Children = types.NewOrderedMap()
    bindingOptions.EntityData.Children.Append("mac-bind-options", types.YChild{"MacBindOptions", &bindingOptions.MacBindOptions})
    bindingOptions.EntityData.Children.Append("duid-bind-options", types.YChild{"DuidBindOptions", &bindingOptions.DuidBindOptions})
    bindingOptions.EntityData.Leafs = types.NewOrderedMap()

    bindingOptions.EntityData.YListKeys = []string {}

    return &(bindingOptions.EntityData)
}

// Dhcpv6_Nodes_Node_Server_BindingOptions_MacBindOptions
// DHCPv6 server binding from MAC address
type Dhcpv6_Nodes_Node_Server_BindingOptions_MacBindOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DHCPv6 server binding with options. The type is slice of
    // Dhcpv6_Nodes_Node_Server_BindingOptions_MacBindOptions_MacBindOption.
    MacBindOption []*Dhcpv6_Nodes_Node_Server_BindingOptions_MacBindOptions_MacBindOption
}

func (macBindOptions *Dhcpv6_Nodes_Node_Server_BindingOptions_MacBindOptions) GetEntityData() *types.CommonEntityData {
    macBindOptions.EntityData.YFilter = macBindOptions.YFilter
    macBindOptions.EntityData.YangName = "mac-bind-options"
    macBindOptions.EntityData.BundleName = "cisco_ios_xr"
    macBindOptions.EntityData.ParentYangName = "binding-options"
    macBindOptions.EntityData.SegmentPath = "mac-bind-options"
    macBindOptions.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/server/binding-options/" + macBindOptions.EntityData.SegmentPath
    macBindOptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macBindOptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macBindOptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macBindOptions.EntityData.Children = types.NewOrderedMap()
    macBindOptions.EntityData.Children.Append("mac-bind-option", types.YChild{"MacBindOption", nil})
    for i := range macBindOptions.MacBindOption {
        macBindOptions.EntityData.Children.Append(types.GetSegmentPath(macBindOptions.MacBindOption[i]), types.YChild{"MacBindOption", macBindOptions.MacBindOption[i]})
    }
    macBindOptions.EntityData.Leafs = types.NewOrderedMap()

    macBindOptions.EntityData.YListKeys = []string {}

    return &(macBindOptions.EntityData)
}

// Dhcpv6_Nodes_Node_Server_BindingOptions_MacBindOptions_MacBindOption
// DHCPv6 server binding with options
type Dhcpv6_Nodes_Node_Server_BindingOptions_MacBindOptions_MacBindOption struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (macBindOption *Dhcpv6_Nodes_Node_Server_BindingOptions_MacBindOptions_MacBindOption) GetEntityData() *types.CommonEntityData {
    macBindOption.EntityData.YFilter = macBindOption.YFilter
    macBindOption.EntityData.YangName = "mac-bind-option"
    macBindOption.EntityData.BundleName = "cisco_ios_xr"
    macBindOption.EntityData.ParentYangName = "mac-bind-options"
    macBindOption.EntityData.SegmentPath = "mac-bind-option" + types.AddKeyToken(macBindOption.MacAddress, "mac-address")
    macBindOption.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/server/binding-options/mac-bind-options/" + macBindOption.EntityData.SegmentPath
    macBindOption.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macBindOption.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macBindOption.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macBindOption.EntityData.Children = types.NewOrderedMap()
    macBindOption.EntityData.Leafs = types.NewOrderedMap()
    macBindOption.EntityData.Leafs.Append("mac-address", types.YLeaf{"MacAddress", macBindOption.MacAddress})
    macBindOption.EntityData.Leafs.Append("mac-address-xr", types.YLeaf{"MacAddressXr", macBindOption.MacAddressXr})
    macBindOption.EntityData.Leafs.Append("duid-xr", types.YLeaf{"DuidXr", macBindOption.DuidXr})
    macBindOption.EntityData.Leafs.Append("dns-count", types.YLeaf{"DnsCount", macBindOption.DnsCount})
    macBindOption.EntityData.Leafs.Append("opt17", types.YLeaf{"Opt17", macBindOption.Opt17})
    macBindOption.EntityData.Leafs.Append("dns-address", types.YLeaf{"DnsAddress", macBindOption.DnsAddress})

    macBindOption.EntityData.YListKeys = []string {"MacAddress"}

    return &(macBindOption.EntityData)
}

// Dhcpv6_Nodes_Node_Server_BindingOptions_DuidBindOptions
// DHCPv6 server binding from DUID
type Dhcpv6_Nodes_Node_Server_BindingOptions_DuidBindOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DHCPv6 server binding with options. The type is slice of
    // Dhcpv6_Nodes_Node_Server_BindingOptions_DuidBindOptions_DuidBindOption.
    DuidBindOption []*Dhcpv6_Nodes_Node_Server_BindingOptions_DuidBindOptions_DuidBindOption
}

func (duidBindOptions *Dhcpv6_Nodes_Node_Server_BindingOptions_DuidBindOptions) GetEntityData() *types.CommonEntityData {
    duidBindOptions.EntityData.YFilter = duidBindOptions.YFilter
    duidBindOptions.EntityData.YangName = "duid-bind-options"
    duidBindOptions.EntityData.BundleName = "cisco_ios_xr"
    duidBindOptions.EntityData.ParentYangName = "binding-options"
    duidBindOptions.EntityData.SegmentPath = "duid-bind-options"
    duidBindOptions.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/server/binding-options/" + duidBindOptions.EntityData.SegmentPath
    duidBindOptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    duidBindOptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    duidBindOptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    duidBindOptions.EntityData.Children = types.NewOrderedMap()
    duidBindOptions.EntityData.Children.Append("duid-bind-option", types.YChild{"DuidBindOption", nil})
    for i := range duidBindOptions.DuidBindOption {
        duidBindOptions.EntityData.Children.Append(types.GetSegmentPath(duidBindOptions.DuidBindOption[i]), types.YChild{"DuidBindOption", duidBindOptions.DuidBindOption[i]})
    }
    duidBindOptions.EntityData.Leafs = types.NewOrderedMap()

    duidBindOptions.EntityData.YListKeys = []string {}

    return &(duidBindOptions.EntityData)
}

// Dhcpv6_Nodes_Node_Server_BindingOptions_DuidBindOptions_DuidBindOption
// DHCPv6 server binding with options
type Dhcpv6_Nodes_Node_Server_BindingOptions_DuidBindOptions_DuidBindOption struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (duidBindOption *Dhcpv6_Nodes_Node_Server_BindingOptions_DuidBindOptions_DuidBindOption) GetEntityData() *types.CommonEntityData {
    duidBindOption.EntityData.YFilter = duidBindOption.YFilter
    duidBindOption.EntityData.YangName = "duid-bind-option"
    duidBindOption.EntityData.BundleName = "cisco_ios_xr"
    duidBindOption.EntityData.ParentYangName = "duid-bind-options"
    duidBindOption.EntityData.SegmentPath = "duid-bind-option" + types.AddKeyToken(duidBindOption.Duid, "duid")
    duidBindOption.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/server/binding-options/duid-bind-options/" + duidBindOption.EntityData.SegmentPath
    duidBindOption.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    duidBindOption.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    duidBindOption.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    duidBindOption.EntityData.Children = types.NewOrderedMap()
    duidBindOption.EntityData.Leafs = types.NewOrderedMap()
    duidBindOption.EntityData.Leafs.Append("duid", types.YLeaf{"Duid", duidBindOption.Duid})
    duidBindOption.EntityData.Leafs.Append("mac-address-xr", types.YLeaf{"MacAddressXr", duidBindOption.MacAddressXr})
    duidBindOption.EntityData.Leafs.Append("duid-xr", types.YLeaf{"DuidXr", duidBindOption.DuidXr})
    duidBindOption.EntityData.Leafs.Append("dns-count", types.YLeaf{"DnsCount", duidBindOption.DnsCount})
    duidBindOption.EntityData.Leafs.Append("opt17", types.YLeaf{"Opt17", duidBindOption.Opt17})
    duidBindOption.EntityData.Leafs.Append("dns-address", types.YLeaf{"DnsAddress", duidBindOption.DnsAddress})

    duidBindOption.EntityData.YListKeys = []string {"Duid"}

    return &(duidBindOption.EntityData)
}

// Dhcpv6_Nodes_Node_Relay
// IPv6 DHCP relay operational data
type Dhcpv6_Nodes_Node_Relay struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DHCPv6 relay statistics.
    Statistics Dhcpv6_Nodes_Node_Relay_Statistics

    // DHCPV6 relay bindings.
    Binding Dhcpv6_Nodes_Node_Relay_Binding

    // DHCPV6 relay list of VRF names.
    Vrfs Dhcpv6_Nodes_Node_Relay_Vrfs
}

func (relay *Dhcpv6_Nodes_Node_Relay) GetEntityData() *types.CommonEntityData {
    relay.EntityData.YFilter = relay.YFilter
    relay.EntityData.YangName = "relay"
    relay.EntityData.BundleName = "cisco_ios_xr"
    relay.EntityData.ParentYangName = "node"
    relay.EntityData.SegmentPath = "relay"
    relay.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/" + relay.EntityData.SegmentPath
    relay.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    relay.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    relay.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    relay.EntityData.Children = types.NewOrderedMap()
    relay.EntityData.Children.Append("statistics", types.YChild{"Statistics", &relay.Statistics})
    relay.EntityData.Children.Append("binding", types.YChild{"Binding", &relay.Binding})
    relay.EntityData.Children.Append("vrfs", types.YChild{"Vrfs", &relay.Vrfs})
    relay.EntityData.Leafs = types.NewOrderedMap()

    relay.EntityData.YListKeys = []string {}

    return &(relay.EntityData)
}

// Dhcpv6_Nodes_Node_Relay_Statistics
// DHCPv6 relay statistics
type Dhcpv6_Nodes_Node_Relay_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ipv6 dhcpv6d relay stat. The type is slice of
    // Dhcpv6_Nodes_Node_Relay_Statistics_Ipv6Dhcpv6dRelayStat.
    Ipv6Dhcpv6dRelayStat []*Dhcpv6_Nodes_Node_Relay_Statistics_Ipv6Dhcpv6dRelayStat
}

func (statistics *Dhcpv6_Nodes_Node_Relay_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "relay"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/relay/" + statistics.EntityData.SegmentPath
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = types.NewOrderedMap()
    statistics.EntityData.Children.Append("ipv6-dhcpv6d-relay-stat", types.YChild{"Ipv6Dhcpv6dRelayStat", nil})
    for i := range statistics.Ipv6Dhcpv6dRelayStat {
        types.SetYListKey(statistics.Ipv6Dhcpv6dRelayStat[i], i)
        statistics.EntityData.Children.Append(types.GetSegmentPath(statistics.Ipv6Dhcpv6dRelayStat[i]), types.YChild{"Ipv6Dhcpv6dRelayStat", statistics.Ipv6Dhcpv6dRelayStat[i]})
    }
    statistics.EntityData.Leafs = types.NewOrderedMap()

    statistics.EntityData.YListKeys = []string {}

    return &(statistics.EntityData)
}

// Dhcpv6_Nodes_Node_Relay_Statistics_Ipv6Dhcpv6dRelayStat
// ipv6 dhcpv6d relay stat
type Dhcpv6_Nodes_Node_Relay_Statistics_Ipv6Dhcpv6dRelayStat struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // DHCPv6 L3 VRF name. The type is string with length: 0..33.
    VrfName interface{}

    // Relay statistics.
    Statistics Dhcpv6_Nodes_Node_Relay_Statistics_Ipv6Dhcpv6dRelayStat_Statistics
}

func (ipv6Dhcpv6dRelayStat *Dhcpv6_Nodes_Node_Relay_Statistics_Ipv6Dhcpv6dRelayStat) GetEntityData() *types.CommonEntityData {
    ipv6Dhcpv6dRelayStat.EntityData.YFilter = ipv6Dhcpv6dRelayStat.YFilter
    ipv6Dhcpv6dRelayStat.EntityData.YangName = "ipv6-dhcpv6d-relay-stat"
    ipv6Dhcpv6dRelayStat.EntityData.BundleName = "cisco_ios_xr"
    ipv6Dhcpv6dRelayStat.EntityData.ParentYangName = "statistics"
    ipv6Dhcpv6dRelayStat.EntityData.SegmentPath = "ipv6-dhcpv6d-relay-stat" + types.AddNoKeyToken(ipv6Dhcpv6dRelayStat)
    ipv6Dhcpv6dRelayStat.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/relay/statistics/" + ipv6Dhcpv6dRelayStat.EntityData.SegmentPath
    ipv6Dhcpv6dRelayStat.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Dhcpv6dRelayStat.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Dhcpv6dRelayStat.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Dhcpv6dRelayStat.EntityData.Children = types.NewOrderedMap()
    ipv6Dhcpv6dRelayStat.EntityData.Children.Append("statistics", types.YChild{"Statistics", &ipv6Dhcpv6dRelayStat.Statistics})
    ipv6Dhcpv6dRelayStat.EntityData.Leafs = types.NewOrderedMap()
    ipv6Dhcpv6dRelayStat.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", ipv6Dhcpv6dRelayStat.VrfName})

    ipv6Dhcpv6dRelayStat.EntityData.YListKeys = []string {}

    return &(ipv6Dhcpv6dRelayStat.EntityData)
}

// Dhcpv6_Nodes_Node_Relay_Statistics_Ipv6Dhcpv6dRelayStat_Statistics
// Relay statistics
type Dhcpv6_Nodes_Node_Relay_Statistics_Ipv6Dhcpv6dRelayStat_Statistics struct {
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

func (statistics *Dhcpv6_Nodes_Node_Relay_Statistics_Ipv6Dhcpv6dRelayStat_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "ipv6-dhcpv6d-relay-stat"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/relay/statistics/ipv6-dhcpv6d-relay-stat/" + statistics.EntityData.SegmentPath
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = types.NewOrderedMap()
    statistics.EntityData.Leafs = types.NewOrderedMap()
    statistics.EntityData.Leafs.Append("received-packets", types.YLeaf{"ReceivedPackets", statistics.ReceivedPackets})
    statistics.EntityData.Leafs.Append("transmitted-packets", types.YLeaf{"TransmittedPackets", statistics.TransmittedPackets})
    statistics.EntityData.Leafs.Append("dropped-packets", types.YLeaf{"DroppedPackets", statistics.DroppedPackets})

    statistics.EntityData.YListKeys = []string {}

    return &(statistics.EntityData)
}

// Dhcpv6_Nodes_Node_Relay_Binding
// DHCPV6 relay bindings
type Dhcpv6_Nodes_Node_Relay_Binding struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DHCPV6 relay binding summary.
    Summary Dhcpv6_Nodes_Node_Relay_Binding_Summary

    // DHCPV6 relay client bindings.
    Clients Dhcpv6_Nodes_Node_Relay_Binding_Clients
}

func (binding *Dhcpv6_Nodes_Node_Relay_Binding) GetEntityData() *types.CommonEntityData {
    binding.EntityData.YFilter = binding.YFilter
    binding.EntityData.YangName = "binding"
    binding.EntityData.BundleName = "cisco_ios_xr"
    binding.EntityData.ParentYangName = "relay"
    binding.EntityData.SegmentPath = "binding"
    binding.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/relay/" + binding.EntityData.SegmentPath
    binding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    binding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    binding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    binding.EntityData.Children = types.NewOrderedMap()
    binding.EntityData.Children.Append("summary", types.YChild{"Summary", &binding.Summary})
    binding.EntityData.Children.Append("clients", types.YChild{"Clients", &binding.Clients})
    binding.EntityData.Leafs = types.NewOrderedMap()

    binding.EntityData.YListKeys = []string {}

    return &(binding.EntityData)
}

// Dhcpv6_Nodes_Node_Relay_Binding_Summary
// DHCPV6 relay binding summary
type Dhcpv6_Nodes_Node_Relay_Binding_Summary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Total number of clients. The type is interface{} with range: 0..4294967295.
    Clients interface{}
}

func (summary *Dhcpv6_Nodes_Node_Relay_Binding_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "binding"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/relay/binding/" + summary.EntityData.SegmentPath
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = types.NewOrderedMap()
    summary.EntityData.Leafs = types.NewOrderedMap()
    summary.EntityData.Leafs.Append("clients", types.YLeaf{"Clients", summary.Clients})

    summary.EntityData.YListKeys = []string {}

    return &(summary.EntityData)
}

// Dhcpv6_Nodes_Node_Relay_Binding_Clients
// DHCPV6 relay client bindings
type Dhcpv6_Nodes_Node_Relay_Binding_Clients struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Single DHCPV6 relay binding. The type is slice of
    // Dhcpv6_Nodes_Node_Relay_Binding_Clients_Client.
    Client []*Dhcpv6_Nodes_Node_Relay_Binding_Clients_Client
}

func (clients *Dhcpv6_Nodes_Node_Relay_Binding_Clients) GetEntityData() *types.CommonEntityData {
    clients.EntityData.YFilter = clients.YFilter
    clients.EntityData.YangName = "clients"
    clients.EntityData.BundleName = "cisco_ios_xr"
    clients.EntityData.ParentYangName = "binding"
    clients.EntityData.SegmentPath = "clients"
    clients.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/relay/binding/" + clients.EntityData.SegmentPath
    clients.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clients.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clients.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clients.EntityData.Children = types.NewOrderedMap()
    clients.EntityData.Children.Append("client", types.YChild{"Client", nil})
    for i := range clients.Client {
        clients.EntityData.Children.Append(types.GetSegmentPath(clients.Client[i]), types.YChild{"Client", clients.Client[i]})
    }
    clients.EntityData.Leafs = types.NewOrderedMap()

    clients.EntityData.YListKeys = []string {}

    return &(clients.EntityData)
}

// Dhcpv6_Nodes_Node_Relay_Binding_Clients_Client
// Single DHCPV6 relay binding
type Dhcpv6_Nodes_Node_Relay_Binding_Clients_Client struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (client *Dhcpv6_Nodes_Node_Relay_Binding_Clients_Client) GetEntityData() *types.CommonEntityData {
    client.EntityData.YFilter = client.YFilter
    client.EntityData.YangName = "client"
    client.EntityData.BundleName = "cisco_ios_xr"
    client.EntityData.ParentYangName = "clients"
    client.EntityData.SegmentPath = "client" + types.AddKeyToken(client.ClientId, "client-id")
    client.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/relay/binding/clients/" + client.EntityData.SegmentPath
    client.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    client.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    client.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    client.EntityData.Children = types.NewOrderedMap()
    client.EntityData.Leafs = types.NewOrderedMap()
    client.EntityData.Leafs.Append("client-id", types.YLeaf{"ClientId", client.ClientId})
    client.EntityData.Leafs.Append("duid", types.YLeaf{"Duid", client.Duid})
    client.EntityData.Leafs.Append("client-id-xr", types.YLeaf{"ClientIdXr", client.ClientIdXr})
    client.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", client.PrefixLength})
    client.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", client.Prefix})
    client.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", client.VrfName})
    client.EntityData.Leafs.Append("lifetime", types.YLeaf{"Lifetime", client.Lifetime})
    client.EntityData.Leafs.Append("rem-life-time", types.YLeaf{"RemLifeTime", client.RemLifeTime})
    client.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", client.InterfaceName})
    client.EntityData.Leafs.Append("next-hop-addr", types.YLeaf{"NextHopAddr", client.NextHopAddr})
    client.EntityData.Leafs.Append("ia-id", types.YLeaf{"IaId", client.IaId})
    client.EntityData.Leafs.Append("relay-profile-name", types.YLeaf{"RelayProfileName", client.RelayProfileName})

    client.EntityData.YListKeys = []string {"ClientId"}

    return &(client.EntityData)
}

// Dhcpv6_Nodes_Node_Relay_Vrfs
// DHCPV6 relay list of VRF names
type Dhcpv6_Nodes_Node_Relay_Vrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 DHCP relay VRF name. The type is slice of
    // Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf.
    Vrf []*Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf
}

func (vrfs *Dhcpv6_Nodes_Node_Relay_Vrfs) GetEntityData() *types.CommonEntityData {
    vrfs.EntityData.YFilter = vrfs.YFilter
    vrfs.EntityData.YangName = "vrfs"
    vrfs.EntityData.BundleName = "cisco_ios_xr"
    vrfs.EntityData.ParentYangName = "relay"
    vrfs.EntityData.SegmentPath = "vrfs"
    vrfs.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/relay/" + vrfs.EntityData.SegmentPath
    vrfs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfs.EntityData.Children = types.NewOrderedMap()
    vrfs.EntityData.Children.Append("vrf", types.YChild{"Vrf", nil})
    for i := range vrfs.Vrf {
        vrfs.EntityData.Children.Append(types.GetSegmentPath(vrfs.Vrf[i]), types.YChild{"Vrf", vrfs.Vrf[i]})
    }
    vrfs.EntityData.Leafs = types.NewOrderedMap()

    vrfs.EntityData.YListKeys = []string {}

    return &(vrfs.EntityData)
}

// Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf
// IPv6 DHCP relay VRF name
type Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. VRF name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // IPv6 DHCP relay statistics.
    Statistics Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics
}

func (vrf *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "vrfs"
    vrf.EntityData.SegmentPath = "vrf" + types.AddKeyToken(vrf.VrfName, "vrf-name")
    vrf.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/relay/vrfs/" + vrf.EntityData.SegmentPath
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = types.NewOrderedMap()
    vrf.EntityData.Children.Append("statistics", types.YChild{"Statistics", &vrf.Statistics})
    vrf.EntityData.Leafs = types.NewOrderedMap()
    vrf.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", vrf.VrfName})

    vrf.EntityData.YListKeys = []string {"VrfName"}

    return &(vrf.EntityData)
}

// Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics
// IPv6 DHCP relay statistics
type Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics struct {
    EntityData types.CommonEntityData
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

func (statistics *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "vrf"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/relay/vrfs/vrf/" + statistics.EntityData.SegmentPath
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = types.NewOrderedMap()
    statistics.EntityData.Children.Append("solicit", types.YChild{"Solicit", &statistics.Solicit})
    statistics.EntityData.Children.Append("advertise", types.YChild{"Advertise", &statistics.Advertise})
    statistics.EntityData.Children.Append("request", types.YChild{"Request", &statistics.Request})
    statistics.EntityData.Children.Append("reply", types.YChild{"Reply", &statistics.Reply})
    statistics.EntityData.Children.Append("confirm", types.YChild{"Confirm", &statistics.Confirm})
    statistics.EntityData.Children.Append("decline", types.YChild{"Decline", &statistics.Decline})
    statistics.EntityData.Children.Append("renew", types.YChild{"Renew", &statistics.Renew})
    statistics.EntityData.Children.Append("rebind", types.YChild{"Rebind", &statistics.Rebind})
    statistics.EntityData.Children.Append("release", types.YChild{"Release", &statistics.Release})
    statistics.EntityData.Children.Append("reconfig", types.YChild{"Reconfig", &statistics.Reconfig})
    statistics.EntityData.Children.Append("inform", types.YChild{"Inform", &statistics.Inform})
    statistics.EntityData.Children.Append("relay-forward", types.YChild{"RelayForward", &statistics.RelayForward})
    statistics.EntityData.Children.Append("relay-reply", types.YChild{"RelayReply", &statistics.RelayReply})
    statistics.EntityData.Children.Append("lease-query", types.YChild{"LeaseQuery", &statistics.LeaseQuery})
    statistics.EntityData.Children.Append("lease-query-reply", types.YChild{"LeaseQueryReply", &statistics.LeaseQueryReply})
    statistics.EntityData.Children.Append("lease-query-done", types.YChild{"LeaseQueryDone", &statistics.LeaseQueryDone})
    statistics.EntityData.Children.Append("lease-query-data", types.YChild{"LeaseQueryData", &statistics.LeaseQueryData})
    statistics.EntityData.Leafs = types.NewOrderedMap()

    statistics.EntityData.YListKeys = []string {}

    return &(statistics.EntityData)
}

// Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Solicit
// DHCPV6 solicit packets
type Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Solicit struct {
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

func (solicit *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Solicit) GetEntityData() *types.CommonEntityData {
    solicit.EntityData.YFilter = solicit.YFilter
    solicit.EntityData.YangName = "solicit"
    solicit.EntityData.BundleName = "cisco_ios_xr"
    solicit.EntityData.ParentYangName = "statistics"
    solicit.EntityData.SegmentPath = "solicit"
    solicit.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/relay/vrfs/vrf/statistics/" + solicit.EntityData.SegmentPath
    solicit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    solicit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    solicit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    solicit.EntityData.Children = types.NewOrderedMap()
    solicit.EntityData.Leafs = types.NewOrderedMap()
    solicit.EntityData.Leafs.Append("received-packets", types.YLeaf{"ReceivedPackets", solicit.ReceivedPackets})
    solicit.EntityData.Leafs.Append("transmitted-packets", types.YLeaf{"TransmittedPackets", solicit.TransmittedPackets})
    solicit.EntityData.Leafs.Append("dropped-packets", types.YLeaf{"DroppedPackets", solicit.DroppedPackets})

    solicit.EntityData.YListKeys = []string {}

    return &(solicit.EntityData)
}

// Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Advertise
// DHCPV6 advertise packets
type Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Advertise struct {
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

func (advertise *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Advertise) GetEntityData() *types.CommonEntityData {
    advertise.EntityData.YFilter = advertise.YFilter
    advertise.EntityData.YangName = "advertise"
    advertise.EntityData.BundleName = "cisco_ios_xr"
    advertise.EntityData.ParentYangName = "statistics"
    advertise.EntityData.SegmentPath = "advertise"
    advertise.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/relay/vrfs/vrf/statistics/" + advertise.EntityData.SegmentPath
    advertise.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    advertise.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    advertise.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    advertise.EntityData.Children = types.NewOrderedMap()
    advertise.EntityData.Leafs = types.NewOrderedMap()
    advertise.EntityData.Leafs.Append("received-packets", types.YLeaf{"ReceivedPackets", advertise.ReceivedPackets})
    advertise.EntityData.Leafs.Append("transmitted-packets", types.YLeaf{"TransmittedPackets", advertise.TransmittedPackets})
    advertise.EntityData.Leafs.Append("dropped-packets", types.YLeaf{"DroppedPackets", advertise.DroppedPackets})

    advertise.EntityData.YListKeys = []string {}

    return &(advertise.EntityData)
}

// Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Request
// DHCPV6 request packets
type Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Request struct {
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

func (request *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Request) GetEntityData() *types.CommonEntityData {
    request.EntityData.YFilter = request.YFilter
    request.EntityData.YangName = "request"
    request.EntityData.BundleName = "cisco_ios_xr"
    request.EntityData.ParentYangName = "statistics"
    request.EntityData.SegmentPath = "request"
    request.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/relay/vrfs/vrf/statistics/" + request.EntityData.SegmentPath
    request.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    request.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    request.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    request.EntityData.Children = types.NewOrderedMap()
    request.EntityData.Leafs = types.NewOrderedMap()
    request.EntityData.Leafs.Append("received-packets", types.YLeaf{"ReceivedPackets", request.ReceivedPackets})
    request.EntityData.Leafs.Append("transmitted-packets", types.YLeaf{"TransmittedPackets", request.TransmittedPackets})
    request.EntityData.Leafs.Append("dropped-packets", types.YLeaf{"DroppedPackets", request.DroppedPackets})

    request.EntityData.YListKeys = []string {}

    return &(request.EntityData)
}

// Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Reply
// DHCPV6 reply packets
type Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Reply struct {
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

func (reply *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Reply) GetEntityData() *types.CommonEntityData {
    reply.EntityData.YFilter = reply.YFilter
    reply.EntityData.YangName = "reply"
    reply.EntityData.BundleName = "cisco_ios_xr"
    reply.EntityData.ParentYangName = "statistics"
    reply.EntityData.SegmentPath = "reply"
    reply.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/relay/vrfs/vrf/statistics/" + reply.EntityData.SegmentPath
    reply.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reply.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reply.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reply.EntityData.Children = types.NewOrderedMap()
    reply.EntityData.Leafs = types.NewOrderedMap()
    reply.EntityData.Leafs.Append("received-packets", types.YLeaf{"ReceivedPackets", reply.ReceivedPackets})
    reply.EntityData.Leafs.Append("transmitted-packets", types.YLeaf{"TransmittedPackets", reply.TransmittedPackets})
    reply.EntityData.Leafs.Append("dropped-packets", types.YLeaf{"DroppedPackets", reply.DroppedPackets})

    reply.EntityData.YListKeys = []string {}

    return &(reply.EntityData)
}

// Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Confirm
// DHCPV6 confirm packets
type Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Confirm struct {
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

func (confirm *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Confirm) GetEntityData() *types.CommonEntityData {
    confirm.EntityData.YFilter = confirm.YFilter
    confirm.EntityData.YangName = "confirm"
    confirm.EntityData.BundleName = "cisco_ios_xr"
    confirm.EntityData.ParentYangName = "statistics"
    confirm.EntityData.SegmentPath = "confirm"
    confirm.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/relay/vrfs/vrf/statistics/" + confirm.EntityData.SegmentPath
    confirm.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    confirm.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    confirm.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    confirm.EntityData.Children = types.NewOrderedMap()
    confirm.EntityData.Leafs = types.NewOrderedMap()
    confirm.EntityData.Leafs.Append("received-packets", types.YLeaf{"ReceivedPackets", confirm.ReceivedPackets})
    confirm.EntityData.Leafs.Append("transmitted-packets", types.YLeaf{"TransmittedPackets", confirm.TransmittedPackets})
    confirm.EntityData.Leafs.Append("dropped-packets", types.YLeaf{"DroppedPackets", confirm.DroppedPackets})

    confirm.EntityData.YListKeys = []string {}

    return &(confirm.EntityData)
}

// Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Decline
// DHCPV6 decline packets
type Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Decline struct {
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

func (decline *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Decline) GetEntityData() *types.CommonEntityData {
    decline.EntityData.YFilter = decline.YFilter
    decline.EntityData.YangName = "decline"
    decline.EntityData.BundleName = "cisco_ios_xr"
    decline.EntityData.ParentYangName = "statistics"
    decline.EntityData.SegmentPath = "decline"
    decline.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/relay/vrfs/vrf/statistics/" + decline.EntityData.SegmentPath
    decline.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    decline.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    decline.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    decline.EntityData.Children = types.NewOrderedMap()
    decline.EntityData.Leafs = types.NewOrderedMap()
    decline.EntityData.Leafs.Append("received-packets", types.YLeaf{"ReceivedPackets", decline.ReceivedPackets})
    decline.EntityData.Leafs.Append("transmitted-packets", types.YLeaf{"TransmittedPackets", decline.TransmittedPackets})
    decline.EntityData.Leafs.Append("dropped-packets", types.YLeaf{"DroppedPackets", decline.DroppedPackets})

    decline.EntityData.YListKeys = []string {}

    return &(decline.EntityData)
}

// Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Renew
// DHCPV6 renew packets
type Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Renew struct {
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

func (renew *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Renew) GetEntityData() *types.CommonEntityData {
    renew.EntityData.YFilter = renew.YFilter
    renew.EntityData.YangName = "renew"
    renew.EntityData.BundleName = "cisco_ios_xr"
    renew.EntityData.ParentYangName = "statistics"
    renew.EntityData.SegmentPath = "renew"
    renew.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/relay/vrfs/vrf/statistics/" + renew.EntityData.SegmentPath
    renew.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    renew.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    renew.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    renew.EntityData.Children = types.NewOrderedMap()
    renew.EntityData.Leafs = types.NewOrderedMap()
    renew.EntityData.Leafs.Append("received-packets", types.YLeaf{"ReceivedPackets", renew.ReceivedPackets})
    renew.EntityData.Leafs.Append("transmitted-packets", types.YLeaf{"TransmittedPackets", renew.TransmittedPackets})
    renew.EntityData.Leafs.Append("dropped-packets", types.YLeaf{"DroppedPackets", renew.DroppedPackets})

    renew.EntityData.YListKeys = []string {}

    return &(renew.EntityData)
}

// Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Rebind
// DHCPV6 rebind packets
type Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Rebind struct {
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

func (rebind *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Rebind) GetEntityData() *types.CommonEntityData {
    rebind.EntityData.YFilter = rebind.YFilter
    rebind.EntityData.YangName = "rebind"
    rebind.EntityData.BundleName = "cisco_ios_xr"
    rebind.EntityData.ParentYangName = "statistics"
    rebind.EntityData.SegmentPath = "rebind"
    rebind.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/relay/vrfs/vrf/statistics/" + rebind.EntityData.SegmentPath
    rebind.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rebind.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rebind.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rebind.EntityData.Children = types.NewOrderedMap()
    rebind.EntityData.Leafs = types.NewOrderedMap()
    rebind.EntityData.Leafs.Append("received-packets", types.YLeaf{"ReceivedPackets", rebind.ReceivedPackets})
    rebind.EntityData.Leafs.Append("transmitted-packets", types.YLeaf{"TransmittedPackets", rebind.TransmittedPackets})
    rebind.EntityData.Leafs.Append("dropped-packets", types.YLeaf{"DroppedPackets", rebind.DroppedPackets})

    rebind.EntityData.YListKeys = []string {}

    return &(rebind.EntityData)
}

// Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Release
// DHCPV6 release packets
type Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Release struct {
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

func (release *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Release) GetEntityData() *types.CommonEntityData {
    release.EntityData.YFilter = release.YFilter
    release.EntityData.YangName = "release"
    release.EntityData.BundleName = "cisco_ios_xr"
    release.EntityData.ParentYangName = "statistics"
    release.EntityData.SegmentPath = "release"
    release.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/relay/vrfs/vrf/statistics/" + release.EntityData.SegmentPath
    release.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    release.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    release.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    release.EntityData.Children = types.NewOrderedMap()
    release.EntityData.Leafs = types.NewOrderedMap()
    release.EntityData.Leafs.Append("received-packets", types.YLeaf{"ReceivedPackets", release.ReceivedPackets})
    release.EntityData.Leafs.Append("transmitted-packets", types.YLeaf{"TransmittedPackets", release.TransmittedPackets})
    release.EntityData.Leafs.Append("dropped-packets", types.YLeaf{"DroppedPackets", release.DroppedPackets})

    release.EntityData.YListKeys = []string {}

    return &(release.EntityData)
}

// Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Reconfig
// DHCPV6 reconfig packets
type Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Reconfig struct {
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

func (reconfig *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Reconfig) GetEntityData() *types.CommonEntityData {
    reconfig.EntityData.YFilter = reconfig.YFilter
    reconfig.EntityData.YangName = "reconfig"
    reconfig.EntityData.BundleName = "cisco_ios_xr"
    reconfig.EntityData.ParentYangName = "statistics"
    reconfig.EntityData.SegmentPath = "reconfig"
    reconfig.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/relay/vrfs/vrf/statistics/" + reconfig.EntityData.SegmentPath
    reconfig.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reconfig.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reconfig.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reconfig.EntityData.Children = types.NewOrderedMap()
    reconfig.EntityData.Leafs = types.NewOrderedMap()
    reconfig.EntityData.Leafs.Append("received-packets", types.YLeaf{"ReceivedPackets", reconfig.ReceivedPackets})
    reconfig.EntityData.Leafs.Append("transmitted-packets", types.YLeaf{"TransmittedPackets", reconfig.TransmittedPackets})
    reconfig.EntityData.Leafs.Append("dropped-packets", types.YLeaf{"DroppedPackets", reconfig.DroppedPackets})

    reconfig.EntityData.YListKeys = []string {}

    return &(reconfig.EntityData)
}

// Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Inform
// DHCPV6 inform packets
type Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Inform struct {
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

func (inform *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_Inform) GetEntityData() *types.CommonEntityData {
    inform.EntityData.YFilter = inform.YFilter
    inform.EntityData.YangName = "inform"
    inform.EntityData.BundleName = "cisco_ios_xr"
    inform.EntityData.ParentYangName = "statistics"
    inform.EntityData.SegmentPath = "inform"
    inform.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/relay/vrfs/vrf/statistics/" + inform.EntityData.SegmentPath
    inform.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inform.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inform.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inform.EntityData.Children = types.NewOrderedMap()
    inform.EntityData.Leafs = types.NewOrderedMap()
    inform.EntityData.Leafs.Append("received-packets", types.YLeaf{"ReceivedPackets", inform.ReceivedPackets})
    inform.EntityData.Leafs.Append("transmitted-packets", types.YLeaf{"TransmittedPackets", inform.TransmittedPackets})
    inform.EntityData.Leafs.Append("dropped-packets", types.YLeaf{"DroppedPackets", inform.DroppedPackets})

    inform.EntityData.YListKeys = []string {}

    return &(inform.EntityData)
}

// Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_RelayForward
// DHCPV6 relay forward packets
type Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_RelayForward struct {
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

func (relayForward *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_RelayForward) GetEntityData() *types.CommonEntityData {
    relayForward.EntityData.YFilter = relayForward.YFilter
    relayForward.EntityData.YangName = "relay-forward"
    relayForward.EntityData.BundleName = "cisco_ios_xr"
    relayForward.EntityData.ParentYangName = "statistics"
    relayForward.EntityData.SegmentPath = "relay-forward"
    relayForward.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/relay/vrfs/vrf/statistics/" + relayForward.EntityData.SegmentPath
    relayForward.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    relayForward.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    relayForward.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    relayForward.EntityData.Children = types.NewOrderedMap()
    relayForward.EntityData.Leafs = types.NewOrderedMap()
    relayForward.EntityData.Leafs.Append("received-packets", types.YLeaf{"ReceivedPackets", relayForward.ReceivedPackets})
    relayForward.EntityData.Leafs.Append("transmitted-packets", types.YLeaf{"TransmittedPackets", relayForward.TransmittedPackets})
    relayForward.EntityData.Leafs.Append("dropped-packets", types.YLeaf{"DroppedPackets", relayForward.DroppedPackets})

    relayForward.EntityData.YListKeys = []string {}

    return &(relayForward.EntityData)
}

// Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_RelayReply
// DHCPV6 relay reply packets
type Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_RelayReply struct {
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

func (relayReply *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_RelayReply) GetEntityData() *types.CommonEntityData {
    relayReply.EntityData.YFilter = relayReply.YFilter
    relayReply.EntityData.YangName = "relay-reply"
    relayReply.EntityData.BundleName = "cisco_ios_xr"
    relayReply.EntityData.ParentYangName = "statistics"
    relayReply.EntityData.SegmentPath = "relay-reply"
    relayReply.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/relay/vrfs/vrf/statistics/" + relayReply.EntityData.SegmentPath
    relayReply.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    relayReply.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    relayReply.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    relayReply.EntityData.Children = types.NewOrderedMap()
    relayReply.EntityData.Leafs = types.NewOrderedMap()
    relayReply.EntityData.Leafs.Append("received-packets", types.YLeaf{"ReceivedPackets", relayReply.ReceivedPackets})
    relayReply.EntityData.Leafs.Append("transmitted-packets", types.YLeaf{"TransmittedPackets", relayReply.TransmittedPackets})
    relayReply.EntityData.Leafs.Append("dropped-packets", types.YLeaf{"DroppedPackets", relayReply.DroppedPackets})

    relayReply.EntityData.YListKeys = []string {}

    return &(relayReply.EntityData)
}

// Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQuery
// DHCPV6 lease query packets
type Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQuery struct {
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

func (leaseQuery *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQuery) GetEntityData() *types.CommonEntityData {
    leaseQuery.EntityData.YFilter = leaseQuery.YFilter
    leaseQuery.EntityData.YangName = "lease-query"
    leaseQuery.EntityData.BundleName = "cisco_ios_xr"
    leaseQuery.EntityData.ParentYangName = "statistics"
    leaseQuery.EntityData.SegmentPath = "lease-query"
    leaseQuery.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/relay/vrfs/vrf/statistics/" + leaseQuery.EntityData.SegmentPath
    leaseQuery.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    leaseQuery.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    leaseQuery.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    leaseQuery.EntityData.Children = types.NewOrderedMap()
    leaseQuery.EntityData.Leafs = types.NewOrderedMap()
    leaseQuery.EntityData.Leafs.Append("received-packets", types.YLeaf{"ReceivedPackets", leaseQuery.ReceivedPackets})
    leaseQuery.EntityData.Leafs.Append("transmitted-packets", types.YLeaf{"TransmittedPackets", leaseQuery.TransmittedPackets})
    leaseQuery.EntityData.Leafs.Append("dropped-packets", types.YLeaf{"DroppedPackets", leaseQuery.DroppedPackets})

    leaseQuery.EntityData.YListKeys = []string {}

    return &(leaseQuery.EntityData)
}

// Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQueryReply
// DHCPV6 lease query reply packets
type Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQueryReply struct {
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

func (leaseQueryReply *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQueryReply) GetEntityData() *types.CommonEntityData {
    leaseQueryReply.EntityData.YFilter = leaseQueryReply.YFilter
    leaseQueryReply.EntityData.YangName = "lease-query-reply"
    leaseQueryReply.EntityData.BundleName = "cisco_ios_xr"
    leaseQueryReply.EntityData.ParentYangName = "statistics"
    leaseQueryReply.EntityData.SegmentPath = "lease-query-reply"
    leaseQueryReply.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/relay/vrfs/vrf/statistics/" + leaseQueryReply.EntityData.SegmentPath
    leaseQueryReply.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    leaseQueryReply.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    leaseQueryReply.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    leaseQueryReply.EntityData.Children = types.NewOrderedMap()
    leaseQueryReply.EntityData.Leafs = types.NewOrderedMap()
    leaseQueryReply.EntityData.Leafs.Append("received-packets", types.YLeaf{"ReceivedPackets", leaseQueryReply.ReceivedPackets})
    leaseQueryReply.EntityData.Leafs.Append("transmitted-packets", types.YLeaf{"TransmittedPackets", leaseQueryReply.TransmittedPackets})
    leaseQueryReply.EntityData.Leafs.Append("dropped-packets", types.YLeaf{"DroppedPackets", leaseQueryReply.DroppedPackets})

    leaseQueryReply.EntityData.YListKeys = []string {}

    return &(leaseQueryReply.EntityData)
}

// Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQueryDone
// DHCPV6 lease query done packets
type Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQueryDone struct {
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

func (leaseQueryDone *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQueryDone) GetEntityData() *types.CommonEntityData {
    leaseQueryDone.EntityData.YFilter = leaseQueryDone.YFilter
    leaseQueryDone.EntityData.YangName = "lease-query-done"
    leaseQueryDone.EntityData.BundleName = "cisco_ios_xr"
    leaseQueryDone.EntityData.ParentYangName = "statistics"
    leaseQueryDone.EntityData.SegmentPath = "lease-query-done"
    leaseQueryDone.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/relay/vrfs/vrf/statistics/" + leaseQueryDone.EntityData.SegmentPath
    leaseQueryDone.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    leaseQueryDone.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    leaseQueryDone.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    leaseQueryDone.EntityData.Children = types.NewOrderedMap()
    leaseQueryDone.EntityData.Leafs = types.NewOrderedMap()
    leaseQueryDone.EntityData.Leafs.Append("received-packets", types.YLeaf{"ReceivedPackets", leaseQueryDone.ReceivedPackets})
    leaseQueryDone.EntityData.Leafs.Append("transmitted-packets", types.YLeaf{"TransmittedPackets", leaseQueryDone.TransmittedPackets})
    leaseQueryDone.EntityData.Leafs.Append("dropped-packets", types.YLeaf{"DroppedPackets", leaseQueryDone.DroppedPackets})

    leaseQueryDone.EntityData.YListKeys = []string {}

    return &(leaseQueryDone.EntityData)
}

// Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQueryData
// DHCPV6 lease query data packets
type Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQueryData struct {
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

func (leaseQueryData *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics_LeaseQueryData) GetEntityData() *types.CommonEntityData {
    leaseQueryData.EntityData.YFilter = leaseQueryData.YFilter
    leaseQueryData.EntityData.YangName = "lease-query-data"
    leaseQueryData.EntityData.BundleName = "cisco_ios_xr"
    leaseQueryData.EntityData.ParentYangName = "statistics"
    leaseQueryData.EntityData.SegmentPath = "lease-query-data"
    leaseQueryData.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-oper:dhcpv6/nodes/node/relay/vrfs/vrf/statistics/" + leaseQueryData.EntityData.SegmentPath
    leaseQueryData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    leaseQueryData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    leaseQueryData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    leaseQueryData.EntityData.Children = types.NewOrderedMap()
    leaseQueryData.EntityData.Leafs = types.NewOrderedMap()
    leaseQueryData.EntityData.Leafs.Append("received-packets", types.YLeaf{"ReceivedPackets", leaseQueryData.ReceivedPackets})
    leaseQueryData.EntityData.Leafs.Append("transmitted-packets", types.YLeaf{"TransmittedPackets", leaseQueryData.TransmittedPackets})
    leaseQueryData.EntityData.Leafs.Append("dropped-packets", types.YLeaf{"DroppedPackets", leaseQueryData.DroppedPackets})

    leaseQueryData.EntityData.YListKeys = []string {}

    return &(leaseQueryData.EntityData)
}

