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

// Dhcpv6IssuVersion represents Dhcpv6 issu version
type Dhcpv6IssuVersion string

const (
    // Version 1
    Dhcpv6IssuVersion_version1 Dhcpv6IssuVersion = "version1"

    // Version 2
    Dhcpv6IssuVersion_version2 Dhcpv6IssuVersion = "version2"
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

// Dhcpv6IssuRole represents Dhcpv6 issu role
type Dhcpv6IssuRole string

const (
    // Primary role
    Dhcpv6IssuRole_role_primary Dhcpv6IssuRole = "role-primary"

    // Secondary role
    Dhcpv6IssuRole_role_secondary Dhcpv6IssuRole = "role-secondary"
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
    dhcpv6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dhcpv6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dhcpv6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dhcpv6.EntityData.Children = make(map[string]types.YChild)
    dhcpv6.EntityData.Children["issu-status"] = types.YChild{"IssuStatus", &dhcpv6.IssuStatus}
    dhcpv6.EntityData.Children["nodes"] = types.YChild{"Nodes", &dhcpv6.Nodes}
    dhcpv6.EntityData.Leafs = make(map[string]types.YLeaf)
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
    issuStatus.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    issuStatus.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    issuStatus.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    issuStatus.EntityData.Children = make(map[string]types.YChild)
    issuStatus.EntityData.Leafs = make(map[string]types.YLeaf)
    issuStatus.EntityData.Leafs["process-start-time"] = types.YLeaf{"ProcessStartTime", issuStatus.ProcessStartTime}
    issuStatus.EntityData.Leafs["issu-sync-complete-time"] = types.YLeaf{"IssuSyncCompleteTime", issuStatus.IssuSyncCompleteTime}
    issuStatus.EntityData.Leafs["issu-sync-start-time"] = types.YLeaf{"IssuSyncStartTime", issuStatus.IssuSyncStartTime}
    issuStatus.EntityData.Leafs["issu-ready-time"] = types.YLeaf{"IssuReadyTime", issuStatus.IssuReadyTime}
    issuStatus.EntityData.Leafs["big-bang-time"] = types.YLeaf{"BigBangTime", issuStatus.BigBangTime}
    issuStatus.EntityData.Leafs["primary-role-time"] = types.YLeaf{"PrimaryRoleTime", issuStatus.PrimaryRoleTime}
    issuStatus.EntityData.Leafs["issu-ready-issu-mgr-connection"] = types.YLeaf{"IssuReadyIssuMgrConnection", issuStatus.IssuReadyIssuMgrConnection}
    issuStatus.EntityData.Leafs["role"] = types.YLeaf{"Role", issuStatus.Role}
    issuStatus.EntityData.Leafs["phase"] = types.YLeaf{"Phase", issuStatus.Phase}
    issuStatus.EntityData.Leafs["version"] = types.YLeaf{"Version", issuStatus.Version}
    return &(issuStatus.EntityData)
}

// Dhcpv6_Nodes
// IPv6 DHCP list of nodes
type Dhcpv6_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 DHCP particular node name. The type is slice of Dhcpv6_Nodes_Node.
    Node []Dhcpv6_Nodes_Node
}

func (nodes *Dhcpv6_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "dhcpv6"
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

// Dhcpv6_Nodes_Node
// IPv6 DHCP particular node name
type Dhcpv6_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node name. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
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
    node.EntityData.SegmentPath = "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Children["proxy"] = types.YChild{"Proxy", &node.Proxy}
    node.EntityData.Children["base"] = types.YChild{"Base", &node.Base}
    node.EntityData.Children["server"] = types.YChild{"Server", &node.Server}
    node.EntityData.Children["relay"] = types.YChild{"Relay", &node.Relay}
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", node.NodeName}
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
    proxy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    proxy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    proxy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    proxy.EntityData.Children = make(map[string]types.YChild)
    proxy.EntityData.Children["vrfs"] = types.YChild{"Vrfs", &proxy.Vrfs}
    proxy.EntityData.Children["profiles"] = types.YChild{"Profiles", &proxy.Profiles}
    proxy.EntityData.Children["interfaces"] = types.YChild{"Interfaces", &proxy.Interfaces}
    proxy.EntityData.Children["statistics"] = types.YChild{"Statistics", &proxy.Statistics}
    proxy.EntityData.Children["binding"] = types.YChild{"Binding", &proxy.Binding}
    proxy.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(proxy.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Vrfs
// DHCPV6 proxy list of VRF names
type Dhcpv6_Nodes_Node_Proxy_Vrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 DHCP proxy VRF name. The type is slice of
    // Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf.
    Vrf []Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf
}

func (vrfs *Dhcpv6_Nodes_Node_Proxy_Vrfs) GetEntityData() *types.CommonEntityData {
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

// Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf
// IPv6 DHCP proxy VRF name
type Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    VrfName interface{}

    // IPv6 DHCP proxy statistics.
    Statistics Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf_Statistics
}

func (vrf *Dhcpv6_Nodes_Node_Proxy_Vrfs_Vrf) GetEntityData() *types.CommonEntityData {
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
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = make(map[string]types.YChild)
    statistics.EntityData.Children["solicit"] = types.YChild{"Solicit", &statistics.Solicit}
    statistics.EntityData.Children["advertise"] = types.YChild{"Advertise", &statistics.Advertise}
    statistics.EntityData.Children["request"] = types.YChild{"Request", &statistics.Request}
    statistics.EntityData.Children["reply"] = types.YChild{"Reply", &statistics.Reply}
    statistics.EntityData.Children["confirm"] = types.YChild{"Confirm", &statistics.Confirm}
    statistics.EntityData.Children["decline"] = types.YChild{"Decline", &statistics.Decline}
    statistics.EntityData.Children["renew"] = types.YChild{"Renew", &statistics.Renew}
    statistics.EntityData.Children["rebind"] = types.YChild{"Rebind", &statistics.Rebind}
    statistics.EntityData.Children["release"] = types.YChild{"Release", &statistics.Release}
    statistics.EntityData.Children["reconfig"] = types.YChild{"Reconfig", &statistics.Reconfig}
    statistics.EntityData.Children["inform"] = types.YChild{"Inform", &statistics.Inform}
    statistics.EntityData.Children["relay-forward"] = types.YChild{"RelayForward", &statistics.RelayForward}
    statistics.EntityData.Children["relay-reply"] = types.YChild{"RelayReply", &statistics.RelayReply}
    statistics.EntityData.Children["lease-query"] = types.YChild{"LeaseQuery", &statistics.LeaseQuery}
    statistics.EntityData.Children["lease-query-reply"] = types.YChild{"LeaseQueryReply", &statistics.LeaseQueryReply}
    statistics.EntityData.Children["lease-query-done"] = types.YChild{"LeaseQueryDone", &statistics.LeaseQueryDone}
    statistics.EntityData.Children["lease-query-data"] = types.YChild{"LeaseQueryData", &statistics.LeaseQueryData}
    statistics.EntityData.Leafs = make(map[string]types.YLeaf)
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
    solicit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    solicit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    solicit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    solicit.EntityData.Children = make(map[string]types.YChild)
    solicit.EntityData.Leafs = make(map[string]types.YLeaf)
    solicit.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", solicit.ReceivedPackets}
    solicit.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", solicit.TransmittedPackets}
    solicit.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", solicit.DroppedPackets}
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
    advertise.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    advertise.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    advertise.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    advertise.EntityData.Children = make(map[string]types.YChild)
    advertise.EntityData.Leafs = make(map[string]types.YLeaf)
    advertise.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", advertise.ReceivedPackets}
    advertise.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", advertise.TransmittedPackets}
    advertise.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", advertise.DroppedPackets}
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
    reply.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reply.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reply.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reply.EntityData.Children = make(map[string]types.YChild)
    reply.EntityData.Leafs = make(map[string]types.YLeaf)
    reply.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", reply.ReceivedPackets}
    reply.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", reply.TransmittedPackets}
    reply.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", reply.DroppedPackets}
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
    confirm.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    confirm.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    confirm.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    confirm.EntityData.Children = make(map[string]types.YChild)
    confirm.EntityData.Leafs = make(map[string]types.YLeaf)
    confirm.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", confirm.ReceivedPackets}
    confirm.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", confirm.TransmittedPackets}
    confirm.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", confirm.DroppedPackets}
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
    renew.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    renew.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    renew.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    renew.EntityData.Children = make(map[string]types.YChild)
    renew.EntityData.Leafs = make(map[string]types.YLeaf)
    renew.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", renew.ReceivedPackets}
    renew.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", renew.TransmittedPackets}
    renew.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", renew.DroppedPackets}
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
    rebind.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rebind.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rebind.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rebind.EntityData.Children = make(map[string]types.YChild)
    rebind.EntityData.Leafs = make(map[string]types.YLeaf)
    rebind.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", rebind.ReceivedPackets}
    rebind.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", rebind.TransmittedPackets}
    rebind.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", rebind.DroppedPackets}
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
    reconfig.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reconfig.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reconfig.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reconfig.EntityData.Children = make(map[string]types.YChild)
    reconfig.EntityData.Leafs = make(map[string]types.YLeaf)
    reconfig.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", reconfig.ReceivedPackets}
    reconfig.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", reconfig.TransmittedPackets}
    reconfig.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", reconfig.DroppedPackets}
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
    relayForward.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    relayForward.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    relayForward.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    relayForward.EntityData.Children = make(map[string]types.YChild)
    relayForward.EntityData.Leafs = make(map[string]types.YLeaf)
    relayForward.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", relayForward.ReceivedPackets}
    relayForward.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", relayForward.TransmittedPackets}
    relayForward.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", relayForward.DroppedPackets}
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
    relayReply.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    relayReply.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    relayReply.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    relayReply.EntityData.Children = make(map[string]types.YChild)
    relayReply.EntityData.Leafs = make(map[string]types.YLeaf)
    relayReply.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", relayReply.ReceivedPackets}
    relayReply.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", relayReply.TransmittedPackets}
    relayReply.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", relayReply.DroppedPackets}
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
    leaseQueryReply.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    leaseQueryReply.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    leaseQueryReply.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    leaseQueryReply.EntityData.Children = make(map[string]types.YChild)
    leaseQueryReply.EntityData.Leafs = make(map[string]types.YLeaf)
    leaseQueryReply.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", leaseQueryReply.ReceivedPackets}
    leaseQueryReply.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", leaseQueryReply.TransmittedPackets}
    leaseQueryReply.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", leaseQueryReply.DroppedPackets}
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
    leaseQueryDone.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    leaseQueryDone.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    leaseQueryDone.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    leaseQueryDone.EntityData.Children = make(map[string]types.YChild)
    leaseQueryDone.EntityData.Leafs = make(map[string]types.YLeaf)
    leaseQueryDone.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", leaseQueryDone.ReceivedPackets}
    leaseQueryDone.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", leaseQueryDone.TransmittedPackets}
    leaseQueryDone.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", leaseQueryDone.DroppedPackets}
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
    leaseQueryData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    leaseQueryData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    leaseQueryData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    leaseQueryData.EntityData.Children = make(map[string]types.YChild)
    leaseQueryData.EntityData.Leafs = make(map[string]types.YLeaf)
    leaseQueryData.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", leaseQueryData.ReceivedPackets}
    leaseQueryData.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", leaseQueryData.TransmittedPackets}
    leaseQueryData.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", leaseQueryData.DroppedPackets}
    return &(leaseQueryData.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Profiles
// IPv6 DHCP proxy profile
type Dhcpv6_Nodes_Node_Proxy_Profiles struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 DHCP proxy profile. The type is slice of
    // Dhcpv6_Nodes_Node_Proxy_Profiles_Profile.
    Profile []Dhcpv6_Nodes_Node_Proxy_Profiles_Profile
}

func (profiles *Dhcpv6_Nodes_Node_Proxy_Profiles) GetEntityData() *types.CommonEntityData {
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

// Dhcpv6_Nodes_Node_Proxy_Profiles_Profile
// IPv6 DHCP proxy profile
type Dhcpv6_Nodes_Node_Proxy_Profiles_Profile struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Profile name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    ProfileName interface{}

    // DHCPV6 throttle table.
    ThrottleInfos Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_ThrottleInfos

    // IPv6 DHCP proxy profile Info.
    Info Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info
}

func (profile *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile) GetEntityData() *types.CommonEntityData {
    profile.EntityData.YFilter = profile.YFilter
    profile.EntityData.YangName = "profile"
    profile.EntityData.BundleName = "cisco_ios_xr"
    profile.EntityData.ParentYangName = "profiles"
    profile.EntityData.SegmentPath = "profile" + "[profile-name='" + fmt.Sprintf("%v", profile.ProfileName) + "']"
    profile.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    profile.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    profile.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    profile.EntityData.Children = make(map[string]types.YChild)
    profile.EntityData.Children["throttle-infos"] = types.YChild{"ThrottleInfos", &profile.ThrottleInfos}
    profile.EntityData.Children["info"] = types.YChild{"Info", &profile.Info}
    profile.EntityData.Leafs = make(map[string]types.YLeaf)
    profile.EntityData.Leafs["profile-name"] = types.YLeaf{"ProfileName", profile.ProfileName}
    return &(profile.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_ThrottleInfos
// DHCPV6 throttle table
type Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_ThrottleInfos struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 DHCP proxy profile Throttle Info. The type is slice of
    // Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_ThrottleInfos_ThrottleInfo.
    ThrottleInfo []Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_ThrottleInfos_ThrottleInfo
}

func (throttleInfos *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_ThrottleInfos) GetEntityData() *types.CommonEntityData {
    throttleInfos.EntityData.YFilter = throttleInfos.YFilter
    throttleInfos.EntityData.YangName = "throttle-infos"
    throttleInfos.EntityData.BundleName = "cisco_ios_xr"
    throttleInfos.EntityData.ParentYangName = "profile"
    throttleInfos.EntityData.SegmentPath = "throttle-infos"
    throttleInfos.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    throttleInfos.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    throttleInfos.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    throttleInfos.EntityData.Children = make(map[string]types.YChild)
    throttleInfos.EntityData.Children["throttle-info"] = types.YChild{"ThrottleInfo", nil}
    for i := range throttleInfos.ThrottleInfo {
        throttleInfos.EntityData.Children[types.GetSegmentPath(&throttleInfos.ThrottleInfo[i])] = types.YChild{"ThrottleInfo", &throttleInfos.ThrottleInfo[i]}
    }
    throttleInfos.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(throttleInfos.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_ThrottleInfos_ThrottleInfo
// IPv6 DHCP proxy profile Throttle Info
type Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_ThrottleInfos_ThrottleInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. MAC address. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    MacAddress interface{}

    // Client MAC address. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
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
    throttleInfo.EntityData.SegmentPath = "throttle-info" + "[mac-address='" + fmt.Sprintf("%v", throttleInfo.MacAddress) + "']"
    throttleInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    throttleInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    throttleInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    throttleInfo.EntityData.Children = make(map[string]types.YChild)
    throttleInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    throttleInfo.EntityData.Leafs["mac-address"] = types.YLeaf{"MacAddress", throttleInfo.MacAddress}
    throttleInfo.EntityData.Leafs["binding-chaddr"] = types.YLeaf{"BindingChaddr", throttleInfo.BindingChaddr}
    throttleInfo.EntityData.Leafs["ifname"] = types.YLeaf{"Ifname", throttleInfo.Ifname}
    throttleInfo.EntityData.Leafs["state"] = types.YLeaf{"State", throttleInfo.State}
    throttleInfo.EntityData.Leafs["time-left"] = types.YLeaf{"TimeLeft", throttleInfo.TimeLeft}
    return &(throttleInfo.EntityData)
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
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ProfileLinkAddress interface{}

    // LinkAddress From RA mesage. The type is bool.
    ProxyProfileLinkaddressFromRaEnable interface{}

    // Helper addresses. The type is slice of string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
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
    info.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    info.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    info.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    info.EntityData.Children = make(map[string]types.YChild)
    info.EntityData.Children["interface-id-references"] = types.YChild{"InterfaceIdReferences", &info.InterfaceIdReferences}
    info.EntityData.Children["vrf-references"] = types.YChild{"VrfReferences", &info.VrfReferences}
    info.EntityData.Children["interface-references"] = types.YChild{"InterfaceReferences", &info.InterfaceReferences}
    info.EntityData.Leafs = make(map[string]types.YLeaf)
    info.EntityData.Leafs["profile-name"] = types.YLeaf{"ProfileName", info.ProfileName}
    info.EntityData.Leafs["remote-id"] = types.YLeaf{"RemoteId", info.RemoteId}
    info.EntityData.Leafs["profile-link-address"] = types.YLeaf{"ProfileLinkAddress", info.ProfileLinkAddress}
    info.EntityData.Leafs["proxy-profile-linkaddress-from-ra-enable"] = types.YLeaf{"ProxyProfileLinkaddressFromRaEnable", info.ProxyProfileLinkaddressFromRaEnable}
    info.EntityData.Leafs["profile-helper-address"] = types.YLeaf{"ProfileHelperAddress", info.ProfileHelperAddress}
    info.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", info.VrfName}
    info.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", info.InterfaceName}
    return &(info.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceIdReferences
// Interface id references
type Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceIdReferences struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ipv6 dhcpv6d proxy iid reference. The type is slice of
    // Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceIdReferences_Ipv6Dhcpv6DProxyIidReference.
    Ipv6Dhcpv6DProxyIidReference []Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceIdReferences_Ipv6Dhcpv6DProxyIidReference
}

func (interfaceIdReferences *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceIdReferences) GetEntityData() *types.CommonEntityData {
    interfaceIdReferences.EntityData.YFilter = interfaceIdReferences.YFilter
    interfaceIdReferences.EntityData.YangName = "interface-id-references"
    interfaceIdReferences.EntityData.BundleName = "cisco_ios_xr"
    interfaceIdReferences.EntityData.ParentYangName = "info"
    interfaceIdReferences.EntityData.SegmentPath = "interface-id-references"
    interfaceIdReferences.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceIdReferences.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceIdReferences.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceIdReferences.EntityData.Children = make(map[string]types.YChild)
    interfaceIdReferences.EntityData.Children["ipv6-dhcpv6d-proxy-iid-reference"] = types.YChild{"Ipv6Dhcpv6DProxyIidReference", nil}
    for i := range interfaceIdReferences.Ipv6Dhcpv6DProxyIidReference {
        interfaceIdReferences.EntityData.Children[types.GetSegmentPath(&interfaceIdReferences.Ipv6Dhcpv6DProxyIidReference[i])] = types.YChild{"Ipv6Dhcpv6DProxyIidReference", &interfaceIdReferences.Ipv6Dhcpv6DProxyIidReference[i]}
    }
    interfaceIdReferences.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaceIdReferences.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceIdReferences_Ipv6Dhcpv6DProxyIidReference
// ipv6 dhcpv6d proxy iid reference
type Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceIdReferences_Ipv6Dhcpv6DProxyIidReference struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface name for interface id. The type is string with length: 0..65.
    ProxyIidInterfaceName interface{}

    // Interface id. The type is string with length: 0..257.
    ProxyInterfaceId interface{}
}

func (ipv6Dhcpv6DProxyIidReference *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceIdReferences_Ipv6Dhcpv6DProxyIidReference) GetEntityData() *types.CommonEntityData {
    ipv6Dhcpv6DProxyIidReference.EntityData.YFilter = ipv6Dhcpv6DProxyIidReference.YFilter
    ipv6Dhcpv6DProxyIidReference.EntityData.YangName = "ipv6-dhcpv6d-proxy-iid-reference"
    ipv6Dhcpv6DProxyIidReference.EntityData.BundleName = "cisco_ios_xr"
    ipv6Dhcpv6DProxyIidReference.EntityData.ParentYangName = "interface-id-references"
    ipv6Dhcpv6DProxyIidReference.EntityData.SegmentPath = "ipv6-dhcpv6d-proxy-iid-reference"
    ipv6Dhcpv6DProxyIidReference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Dhcpv6DProxyIidReference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Dhcpv6DProxyIidReference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Dhcpv6DProxyIidReference.EntityData.Children = make(map[string]types.YChild)
    ipv6Dhcpv6DProxyIidReference.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv6Dhcpv6DProxyIidReference.EntityData.Leafs["proxy-iid-interface-name"] = types.YLeaf{"ProxyIidInterfaceName", ipv6Dhcpv6DProxyIidReference.ProxyIidInterfaceName}
    ipv6Dhcpv6DProxyIidReference.EntityData.Leafs["proxy-interface-id"] = types.YLeaf{"ProxyInterfaceId", ipv6Dhcpv6DProxyIidReference.ProxyInterfaceId}
    return &(ipv6Dhcpv6DProxyIidReference.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_VrfReferences
// VRF references
type Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_VrfReferences struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ipv6 dhcpv6d proxy vrf reference. The type is slice of
    // Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_VrfReferences_Ipv6Dhcpv6DProxyVrfReference.
    Ipv6Dhcpv6DProxyVrfReference []Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_VrfReferences_Ipv6Dhcpv6DProxyVrfReference
}

func (vrfReferences *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_VrfReferences) GetEntityData() *types.CommonEntityData {
    vrfReferences.EntityData.YFilter = vrfReferences.YFilter
    vrfReferences.EntityData.YangName = "vrf-references"
    vrfReferences.EntityData.BundleName = "cisco_ios_xr"
    vrfReferences.EntityData.ParentYangName = "info"
    vrfReferences.EntityData.SegmentPath = "vrf-references"
    vrfReferences.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfReferences.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfReferences.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfReferences.EntityData.Children = make(map[string]types.YChild)
    vrfReferences.EntityData.Children["ipv6-dhcpv6d-proxy-vrf-reference"] = types.YChild{"Ipv6Dhcpv6DProxyVrfReference", nil}
    for i := range vrfReferences.Ipv6Dhcpv6DProxyVrfReference {
        vrfReferences.EntityData.Children[types.GetSegmentPath(&vrfReferences.Ipv6Dhcpv6DProxyVrfReference[i])] = types.YChild{"Ipv6Dhcpv6DProxyVrfReference", &vrfReferences.Ipv6Dhcpv6DProxyVrfReference[i]}
    }
    vrfReferences.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(vrfReferences.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_VrfReferences_Ipv6Dhcpv6DProxyVrfReference
// ipv6 dhcpv6d proxy vrf reference
type Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_VrfReferences_Ipv6Dhcpv6DProxyVrfReference struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF name. The type is string with length: 0..33.
    ProxyReferenceVrfName interface{}
}

func (ipv6Dhcpv6DProxyVrfReference *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_VrfReferences_Ipv6Dhcpv6DProxyVrfReference) GetEntityData() *types.CommonEntityData {
    ipv6Dhcpv6DProxyVrfReference.EntityData.YFilter = ipv6Dhcpv6DProxyVrfReference.YFilter
    ipv6Dhcpv6DProxyVrfReference.EntityData.YangName = "ipv6-dhcpv6d-proxy-vrf-reference"
    ipv6Dhcpv6DProxyVrfReference.EntityData.BundleName = "cisco_ios_xr"
    ipv6Dhcpv6DProxyVrfReference.EntityData.ParentYangName = "vrf-references"
    ipv6Dhcpv6DProxyVrfReference.EntityData.SegmentPath = "ipv6-dhcpv6d-proxy-vrf-reference"
    ipv6Dhcpv6DProxyVrfReference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Dhcpv6DProxyVrfReference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Dhcpv6DProxyVrfReference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Dhcpv6DProxyVrfReference.EntityData.Children = make(map[string]types.YChild)
    ipv6Dhcpv6DProxyVrfReference.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv6Dhcpv6DProxyVrfReference.EntityData.Leafs["proxy-reference-vrf-name"] = types.YLeaf{"ProxyReferenceVrfName", ipv6Dhcpv6DProxyVrfReference.ProxyReferenceVrfName}
    return &(ipv6Dhcpv6DProxyVrfReference.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceReferences
// Interface references
type Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceReferences struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ipv6 dhcpv6d proxy interface reference. The type is slice of
    // Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceReferences_Ipv6Dhcpv6DProxyInterfaceReference.
    Ipv6Dhcpv6DProxyInterfaceReference []Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceReferences_Ipv6Dhcpv6DProxyInterfaceReference
}

func (interfaceReferences *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceReferences) GetEntityData() *types.CommonEntityData {
    interfaceReferences.EntityData.YFilter = interfaceReferences.YFilter
    interfaceReferences.EntityData.YangName = "interface-references"
    interfaceReferences.EntityData.BundleName = "cisco_ios_xr"
    interfaceReferences.EntityData.ParentYangName = "info"
    interfaceReferences.EntityData.SegmentPath = "interface-references"
    interfaceReferences.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceReferences.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceReferences.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceReferences.EntityData.Children = make(map[string]types.YChild)
    interfaceReferences.EntityData.Children["ipv6-dhcpv6d-proxy-interface-reference"] = types.YChild{"Ipv6Dhcpv6DProxyInterfaceReference", nil}
    for i := range interfaceReferences.Ipv6Dhcpv6DProxyInterfaceReference {
        interfaceReferences.EntityData.Children[types.GetSegmentPath(&interfaceReferences.Ipv6Dhcpv6DProxyInterfaceReference[i])] = types.YChild{"Ipv6Dhcpv6DProxyInterfaceReference", &interfaceReferences.Ipv6Dhcpv6DProxyInterfaceReference[i]}
    }
    interfaceReferences.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaceReferences.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceReferences_Ipv6Dhcpv6DProxyInterfaceReference
// ipv6 dhcpv6d proxy interface reference
type Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceReferences_Ipv6Dhcpv6DProxyInterfaceReference struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface name. The type is string with length: 0..65.
    ProxyReferenceInterfaceName interface{}
}

func (ipv6Dhcpv6DProxyInterfaceReference *Dhcpv6_Nodes_Node_Proxy_Profiles_Profile_Info_InterfaceReferences_Ipv6Dhcpv6DProxyInterfaceReference) GetEntityData() *types.CommonEntityData {
    ipv6Dhcpv6DProxyInterfaceReference.EntityData.YFilter = ipv6Dhcpv6DProxyInterfaceReference.YFilter
    ipv6Dhcpv6DProxyInterfaceReference.EntityData.YangName = "ipv6-dhcpv6d-proxy-interface-reference"
    ipv6Dhcpv6DProxyInterfaceReference.EntityData.BundleName = "cisco_ios_xr"
    ipv6Dhcpv6DProxyInterfaceReference.EntityData.ParentYangName = "interface-references"
    ipv6Dhcpv6DProxyInterfaceReference.EntityData.SegmentPath = "ipv6-dhcpv6d-proxy-interface-reference"
    ipv6Dhcpv6DProxyInterfaceReference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Dhcpv6DProxyInterfaceReference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Dhcpv6DProxyInterfaceReference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Dhcpv6DProxyInterfaceReference.EntityData.Children = make(map[string]types.YChild)
    ipv6Dhcpv6DProxyInterfaceReference.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv6Dhcpv6DProxyInterfaceReference.EntityData.Leafs["proxy-reference-interface-name"] = types.YLeaf{"ProxyReferenceInterfaceName", ipv6Dhcpv6DProxyInterfaceReference.ProxyReferenceInterfaceName}
    return &(ipv6Dhcpv6DProxyInterfaceReference.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Interfaces
// DHCPV6 proxy interface
type Dhcpv6_Nodes_Node_Proxy_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 DHCP proxy interface. The type is slice of
    // Dhcpv6_Nodes_Node_Proxy_Interfaces_Interface_.
    Interface_ []Dhcpv6_Nodes_Node_Proxy_Interfaces_Interface
}

func (interfaces *Dhcpv6_Nodes_Node_Proxy_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "proxy"
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

// Dhcpv6_Nodes_Node_Proxy_Interfaces_Interface
// IPv6 DHCP proxy interface
type Dhcpv6_Nodes_Node_Proxy_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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

func (self *Dhcpv6_Nodes_Node_Proxy_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
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
    self.EntityData.Leafs["proxy-vrf-name"] = types.YLeaf{"ProxyVrfName", self.ProxyVrfName}
    self.EntityData.Leafs["proxy-interface-mode"] = types.YLeaf{"ProxyInterfaceMode", self.ProxyInterfaceMode}
    self.EntityData.Leafs["is-proxy-interface-ambiguous"] = types.YLeaf{"IsProxyInterfaceAmbiguous", self.IsProxyInterfaceAmbiguous}
    self.EntityData.Leafs["proxy-interface-profile-name"] = types.YLeaf{"ProxyInterfaceProfileName", self.ProxyInterfaceProfileName}
    self.EntityData.Leafs["proxy-interface-lease-limit-type"] = types.YLeaf{"ProxyInterfaceLeaseLimitType", self.ProxyInterfaceLeaseLimitType}
    self.EntityData.Leafs["proxy-interface-lease-limits"] = types.YLeaf{"ProxyInterfaceLeaseLimits", self.ProxyInterfaceLeaseLimits}
    self.EntityData.Leafs["srg-role"] = types.YLeaf{"SrgRole", self.SrgRole}
    self.EntityData.Leafs["serg-role"] = types.YLeaf{"SergRole", self.SergRole}
    self.EntityData.Leafs["mac-throttle"] = types.YLeaf{"MacThrottle", self.MacThrottle}
    self.EntityData.Leafs["srg-vrf-name"] = types.YLeaf{"SrgVrfName", self.SrgVrfName}
    self.EntityData.Leafs["srgp2p"] = types.YLeaf{"Srgp2P", self.Srgp2P}
    return &(self.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Statistics
// DHCPv6 proxy statistics
type Dhcpv6_Nodes_Node_Proxy_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ipv6 dhcpv6d proxy stat. The type is slice of
    // Dhcpv6_Nodes_Node_Proxy_Statistics_Ipv6Dhcpv6DProxyStat.
    Ipv6Dhcpv6DProxyStat []Dhcpv6_Nodes_Node_Proxy_Statistics_Ipv6Dhcpv6DProxyStat
}

func (statistics *Dhcpv6_Nodes_Node_Proxy_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "proxy"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = make(map[string]types.YChild)
    statistics.EntityData.Children["ipv6-dhcpv6d-proxy-stat"] = types.YChild{"Ipv6Dhcpv6DProxyStat", nil}
    for i := range statistics.Ipv6Dhcpv6DProxyStat {
        statistics.EntityData.Children[types.GetSegmentPath(&statistics.Ipv6Dhcpv6DProxyStat[i])] = types.YChild{"Ipv6Dhcpv6DProxyStat", &statistics.Ipv6Dhcpv6DProxyStat[i]}
    }
    statistics.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(statistics.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Statistics_Ipv6Dhcpv6DProxyStat
// ipv6 dhcpv6d proxy stat
type Dhcpv6_Nodes_Node_Proxy_Statistics_Ipv6Dhcpv6DProxyStat struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DHCPv6 L3 VRF name. The type is string with length: 0..33.
    VrfName interface{}

    // Proxy statistics.
    Statistics Dhcpv6_Nodes_Node_Proxy_Statistics_Ipv6Dhcpv6DProxyStat_Statistics_
}

func (ipv6Dhcpv6DProxyStat *Dhcpv6_Nodes_Node_Proxy_Statistics_Ipv6Dhcpv6DProxyStat) GetEntityData() *types.CommonEntityData {
    ipv6Dhcpv6DProxyStat.EntityData.YFilter = ipv6Dhcpv6DProxyStat.YFilter
    ipv6Dhcpv6DProxyStat.EntityData.YangName = "ipv6-dhcpv6d-proxy-stat"
    ipv6Dhcpv6DProxyStat.EntityData.BundleName = "cisco_ios_xr"
    ipv6Dhcpv6DProxyStat.EntityData.ParentYangName = "statistics"
    ipv6Dhcpv6DProxyStat.EntityData.SegmentPath = "ipv6-dhcpv6d-proxy-stat"
    ipv6Dhcpv6DProxyStat.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Dhcpv6DProxyStat.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Dhcpv6DProxyStat.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Dhcpv6DProxyStat.EntityData.Children = make(map[string]types.YChild)
    ipv6Dhcpv6DProxyStat.EntityData.Children["statistics"] = types.YChild{"Statistics", &ipv6Dhcpv6DProxyStat.Statistics}
    ipv6Dhcpv6DProxyStat.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv6Dhcpv6DProxyStat.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", ipv6Dhcpv6DProxyStat.VrfName}
    return &(ipv6Dhcpv6DProxyStat.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Statistics_Ipv6Dhcpv6DProxyStat_Statistics_
// Proxy statistics
type Dhcpv6_Nodes_Node_Proxy_Statistics_Ipv6Dhcpv6DProxyStat_Statistics_ struct {
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

func (statistics_ *Dhcpv6_Nodes_Node_Proxy_Statistics_Ipv6Dhcpv6DProxyStat_Statistics_) GetEntityData() *types.CommonEntityData {
    statistics_.EntityData.YFilter = statistics_.YFilter
    statistics_.EntityData.YangName = "statistics"
    statistics_.EntityData.BundleName = "cisco_ios_xr"
    statistics_.EntityData.ParentYangName = "ipv6-dhcpv6d-proxy-stat"
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
    binding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    binding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    binding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    binding.EntityData.Children = make(map[string]types.YChild)
    binding.EntityData.Children["clients"] = types.YChild{"Clients", &binding.Clients}
    binding.EntityData.Children["summary"] = types.YChild{"Summary", &binding.Summary}
    binding.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(binding.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Binding_Clients
// DHCPV6 proxy client bindings
type Dhcpv6_Nodes_Node_Proxy_Binding_Clients struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Single DHCPV6 proxy binding. The type is slice of
    // Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client.
    Client []Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client
}

func (clients *Dhcpv6_Nodes_Node_Proxy_Binding_Clients) GetEntityData() *types.CommonEntityData {
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

// Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client
// Single DHCPV6 proxy binding
type Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Client ID. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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
    // b'[a-zA-Z0-9./-]+'.
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
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ServerIpv6Address interface{}

    // DHCPV6 profile name. The type is string with length: 0..65.
    ProfileName interface{}

    // DHCPV6 framed ipv6 addess used by ND. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
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

func (client *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client) GetEntityData() *types.CommonEntityData {
    client.EntityData.YFilter = client.YFilter
    client.EntityData.YangName = "client"
    client.EntityData.BundleName = "cisco_ios_xr"
    client.EntityData.ParentYangName = "clients"
    client.EntityData.SegmentPath = "client" + "[client-id='" + fmt.Sprintf("%v", client.ClientId) + "']"
    client.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    client.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    client.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    client.EntityData.Children = make(map[string]types.YChild)
    client.EntityData.Children["ia-id-pd"] = types.YChild{"IaIdPd", &client.IaIdPd}
    client.EntityData.Leafs = make(map[string]types.YLeaf)
    client.EntityData.Leafs["client-id"] = types.YLeaf{"ClientId", client.ClientId}
    client.EntityData.Leafs["duid"] = types.YLeaf{"Duid", client.Duid}
    client.EntityData.Leafs["client-flag"] = types.YLeaf{"ClientFlag", client.ClientFlag}
    client.EntityData.Leafs["subscriber-label"] = types.YLeaf{"SubscriberLabel", client.SubscriberLabel}
    client.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", client.VrfName}
    client.EntityData.Leafs["mac-address"] = types.YLeaf{"MacAddress", client.MacAddress}
    client.EntityData.Leafs["ia-id-p-ds"] = types.YLeaf{"IaIdPDs", client.IaIdPDs}
    client.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", client.InterfaceName}
    client.EntityData.Leafs["access-vrf-name"] = types.YLeaf{"AccessVrfName", client.AccessVrfName}
    client.EntityData.Leafs["proxy-binding-tags"] = types.YLeaf{"ProxyBindingTags", client.ProxyBindingTags}
    client.EntityData.Leafs["proxy-binding-outer-tag"] = types.YLeaf{"ProxyBindingOuterTag", client.ProxyBindingOuterTag}
    client.EntityData.Leafs["proxy-binding-inner-tag"] = types.YLeaf{"ProxyBindingInnerTag", client.ProxyBindingInnerTag}
    client.EntityData.Leafs["class-name"] = types.YLeaf{"ClassName", client.ClassName}
    client.EntityData.Leafs["pool-name"] = types.YLeaf{"PoolName", client.PoolName}
    client.EntityData.Leafs["rx-remote-id"] = types.YLeaf{"RxRemoteId", client.RxRemoteId}
    client.EntityData.Leafs["tx-remote-id"] = types.YLeaf{"TxRemoteId", client.TxRemoteId}
    client.EntityData.Leafs["rx-interface-id"] = types.YLeaf{"RxInterfaceId", client.RxInterfaceId}
    client.EntityData.Leafs["tx-interface-id"] = types.YLeaf{"TxInterfaceId", client.TxInterfaceId}
    client.EntityData.Leafs["server-ipv6-address"] = types.YLeaf{"ServerIpv6Address", client.ServerIpv6Address}
    client.EntityData.Leafs["profile-name"] = types.YLeaf{"ProfileName", client.ProfileName}
    client.EntityData.Leafs["framed-ipv6-prefix"] = types.YLeaf{"FramedIpv6Prefix", client.FramedIpv6Prefix}
    client.EntityData.Leafs["framed-prefix-length"] = types.YLeaf{"FramedPrefixLength", client.FramedPrefixLength}
    client.EntityData.Leafs["is-nak-next-renew"] = types.YLeaf{"IsNakNextRenew", client.IsNakNextRenew}
    client.EntityData.Leafs["srg-state"] = types.YLeaf{"SrgState", client.SrgState}
    client.EntityData.Leafs["srg-intf-role"] = types.YLeaf{"SrgIntfRole", client.SrgIntfRole}
    client.EntityData.Leafs["srgp2p"] = types.YLeaf{"Srgp2P", client.Srgp2P}
    client.EntityData.Leafs["srg-vrf-name"] = types.YLeaf{"SrgVrfName", client.SrgVrfName}
    client.EntityData.Leafs["serg-state"] = types.YLeaf{"SergState", client.SergState}
    client.EntityData.Leafs["serg-intf-role"] = types.YLeaf{"SergIntfRole", client.SergIntfRole}
    return &(client.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd
// List of DHCPv6 IA_ID/PDs
type Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bag dhcpv6d ia id pd info. The type is slice of
    // Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo.
    BagDhcpv6DIaIdPdInfo []Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo
}

func (iaIdPd *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd) GetEntityData() *types.CommonEntityData {
    iaIdPd.EntityData.YFilter = iaIdPd.YFilter
    iaIdPd.EntityData.YangName = "ia-id-pd"
    iaIdPd.EntityData.BundleName = "cisco_ios_xr"
    iaIdPd.EntityData.ParentYangName = "client"
    iaIdPd.EntityData.SegmentPath = "ia-id-pd"
    iaIdPd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    iaIdPd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    iaIdPd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    iaIdPd.EntityData.Children = make(map[string]types.YChild)
    iaIdPd.EntityData.Children["bag-dhcpv6d-ia-id-pd-info"] = types.YChild{"BagDhcpv6DIaIdPdInfo", nil}
    for i := range iaIdPd.BagDhcpv6DIaIdPdInfo {
        iaIdPd.EntityData.Children[types.GetSegmentPath(&iaIdPd.BagDhcpv6DIaIdPdInfo[i])] = types.YChild{"BagDhcpv6DIaIdPdInfo", &iaIdPd.BagDhcpv6DIaIdPdInfo[i]}
    }
    iaIdPd.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(iaIdPd.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo
// bag dhcpv6d ia id pd info
type Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo struct {
    EntityData types.CommonEntityData
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

func (bagDhcpv6DIaIdPdInfo *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo) GetEntityData() *types.CommonEntityData {
    bagDhcpv6DIaIdPdInfo.EntityData.YFilter = bagDhcpv6DIaIdPdInfo.YFilter
    bagDhcpv6DIaIdPdInfo.EntityData.YangName = "bag-dhcpv6d-ia-id-pd-info"
    bagDhcpv6DIaIdPdInfo.EntityData.BundleName = "cisco_ios_xr"
    bagDhcpv6DIaIdPdInfo.EntityData.ParentYangName = "ia-id-pd"
    bagDhcpv6DIaIdPdInfo.EntityData.SegmentPath = "bag-dhcpv6d-ia-id-pd-info"
    bagDhcpv6DIaIdPdInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bagDhcpv6DIaIdPdInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bagDhcpv6DIaIdPdInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bagDhcpv6DIaIdPdInfo.EntityData.Children = make(map[string]types.YChild)
    bagDhcpv6DIaIdPdInfo.EntityData.Children["addresses"] = types.YChild{"Addresses", &bagDhcpv6DIaIdPdInfo.Addresses}
    bagDhcpv6DIaIdPdInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    bagDhcpv6DIaIdPdInfo.EntityData.Leafs["ia-type"] = types.YLeaf{"IaType", bagDhcpv6DIaIdPdInfo.IaType}
    bagDhcpv6DIaIdPdInfo.EntityData.Leafs["ia-id"] = types.YLeaf{"IaId", bagDhcpv6DIaIdPdInfo.IaId}
    bagDhcpv6DIaIdPdInfo.EntityData.Leafs["flags"] = types.YLeaf{"Flags", bagDhcpv6DIaIdPdInfo.Flags}
    bagDhcpv6DIaIdPdInfo.EntityData.Leafs["total-address"] = types.YLeaf{"TotalAddress", bagDhcpv6DIaIdPdInfo.TotalAddress}
    bagDhcpv6DIaIdPdInfo.EntityData.Leafs["state"] = types.YLeaf{"State", bagDhcpv6DIaIdPdInfo.State}
    return &(bagDhcpv6DIaIdPdInfo.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses
// List of addresses in this IA
type Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bag dhcpv6d addr attrb. The type is slice of
    // Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses_BagDhcpv6DAddrAttrb.
    BagDhcpv6DAddrAttrb []Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses_BagDhcpv6DAddrAttrb
}

func (addresses *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses) GetEntityData() *types.CommonEntityData {
    addresses.EntityData.YFilter = addresses.YFilter
    addresses.EntityData.YangName = "addresses"
    addresses.EntityData.BundleName = "cisco_ios_xr"
    addresses.EntityData.ParentYangName = "bag-dhcpv6d-ia-id-pd-info"
    addresses.EntityData.SegmentPath = "addresses"
    addresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addresses.EntityData.Children = make(map[string]types.YChild)
    addresses.EntityData.Children["bag-dhcpv6d-addr-attrb"] = types.YChild{"BagDhcpv6DAddrAttrb", nil}
    for i := range addresses.BagDhcpv6DAddrAttrb {
        addresses.EntityData.Children[types.GetSegmentPath(&addresses.BagDhcpv6DAddrAttrb[i])] = types.YChild{"BagDhcpv6DAddrAttrb", &addresses.BagDhcpv6DAddrAttrb[i]}
    }
    addresses.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(addresses.EntityData)
}

// Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses_BagDhcpv6DAddrAttrb
// bag dhcpv6d addr attrb
type Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses_BagDhcpv6DAddrAttrb struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 prefix. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
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

func (bagDhcpv6DAddrAttrb *Dhcpv6_Nodes_Node_Proxy_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses_BagDhcpv6DAddrAttrb) GetEntityData() *types.CommonEntityData {
    bagDhcpv6DAddrAttrb.EntityData.YFilter = bagDhcpv6DAddrAttrb.YFilter
    bagDhcpv6DAddrAttrb.EntityData.YangName = "bag-dhcpv6d-addr-attrb"
    bagDhcpv6DAddrAttrb.EntityData.BundleName = "cisco_ios_xr"
    bagDhcpv6DAddrAttrb.EntityData.ParentYangName = "addresses"
    bagDhcpv6DAddrAttrb.EntityData.SegmentPath = "bag-dhcpv6d-addr-attrb"
    bagDhcpv6DAddrAttrb.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bagDhcpv6DAddrAttrb.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bagDhcpv6DAddrAttrb.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bagDhcpv6DAddrAttrb.EntityData.Children = make(map[string]types.YChild)
    bagDhcpv6DAddrAttrb.EntityData.Leafs = make(map[string]types.YLeaf)
    bagDhcpv6DAddrAttrb.EntityData.Leafs["prefix"] = types.YLeaf{"Prefix", bagDhcpv6DAddrAttrb.Prefix}
    bagDhcpv6DAddrAttrb.EntityData.Leafs["prefix-length"] = types.YLeaf{"PrefixLength", bagDhcpv6DAddrAttrb.PrefixLength}
    bagDhcpv6DAddrAttrb.EntityData.Leafs["lease-time"] = types.YLeaf{"LeaseTime", bagDhcpv6DAddrAttrb.LeaseTime}
    bagDhcpv6DAddrAttrb.EntityData.Leafs["remaining-lease-time"] = types.YLeaf{"RemainingLeaseTime", bagDhcpv6DAddrAttrb.RemainingLeaseTime}
    return &(bagDhcpv6DAddrAttrb.EntityData)
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
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = make(map[string]types.YChild)
    summary.EntityData.Children["iana"] = types.YChild{"Iana", &summary.Iana}
    summary.EntityData.Children["iapd"] = types.YChild{"Iapd", &summary.Iapd}
    summary.EntityData.Leafs = make(map[string]types.YLeaf)
    summary.EntityData.Leafs["clients"] = types.YLeaf{"Clients", summary.Clients}
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
    iana.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    iana.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    iana.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    iana.EntityData.Children = make(map[string]types.YChild)
    iana.EntityData.Leafs = make(map[string]types.YLeaf)
    iana.EntityData.Leafs["initializing-clients"] = types.YLeaf{"InitializingClients", iana.InitializingClients}
    iana.EntityData.Leafs["dpm-waiting-clients"] = types.YLeaf{"DpmWaitingClients", iana.DpmWaitingClients}
    iana.EntityData.Leafs["daps-waiting-clients"] = types.YLeaf{"DapsWaitingClients", iana.DapsWaitingClients}
    iana.EntityData.Leafs["msg-waiting-clients"] = types.YLeaf{"MsgWaitingClients", iana.MsgWaitingClients}
    iana.EntityData.Leafs["iedge-waiting-clients"] = types.YLeaf{"IedgeWaitingClients", iana.IedgeWaitingClients}
    iana.EntityData.Leafs["rib-waiting-clients"] = types.YLeaf{"RibWaitingClients", iana.RibWaitingClients}
    iana.EntityData.Leafs["bound-clients"] = types.YLeaf{"BoundClients", iana.BoundClients}
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
    iapd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    iapd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    iapd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    iapd.EntityData.Children = make(map[string]types.YChild)
    iapd.EntityData.Leafs = make(map[string]types.YLeaf)
    iapd.EntityData.Leafs["initializing-clients"] = types.YLeaf{"InitializingClients", iapd.InitializingClients}
    iapd.EntityData.Leafs["dpm-waiting-clients"] = types.YLeaf{"DpmWaitingClients", iapd.DpmWaitingClients}
    iapd.EntityData.Leafs["daps-waiting-clients"] = types.YLeaf{"DapsWaitingClients", iapd.DapsWaitingClients}
    iapd.EntityData.Leafs["msg-waiting-clients"] = types.YLeaf{"MsgWaitingClients", iapd.MsgWaitingClients}
    iapd.EntityData.Leafs["iedge-waiting-clients"] = types.YLeaf{"IedgeWaitingClients", iapd.IedgeWaitingClients}
    iapd.EntityData.Leafs["rib-waiting-clients"] = types.YLeaf{"RibWaitingClients", iapd.RibWaitingClients}
    iapd.EntityData.Leafs["bound-clients"] = types.YLeaf{"BoundClients", iapd.BoundClients}
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
    base.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    base.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    base.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    base.EntityData.Children = make(map[string]types.YChild)
    base.EntityData.Children["database"] = types.YChild{"Database", &base.Database}
    base.EntityData.Children["addr-bindings"] = types.YChild{"AddrBindings", &base.AddrBindings}
    base.EntityData.Leafs = make(map[string]types.YLeaf)
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

// Dhcpv6_Nodes_Node_Base_AddrBindings
// IPv6 DHCP Base Binding
type Dhcpv6_Nodes_Node_Base_AddrBindings struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DHCPv6 base stats debug. The type is slice of
    // Dhcpv6_Nodes_Node_Base_AddrBindings_AddrBinding.
    AddrBinding []Dhcpv6_Nodes_Node_Base_AddrBindings_AddrBinding
}

func (addrBindings *Dhcpv6_Nodes_Node_Base_AddrBindings) GetEntityData() *types.CommonEntityData {
    addrBindings.EntityData.YFilter = addrBindings.YFilter
    addrBindings.EntityData.YangName = "addr-bindings"
    addrBindings.EntityData.BundleName = "cisco_ios_xr"
    addrBindings.EntityData.ParentYangName = "base"
    addrBindings.EntityData.SegmentPath = "addr-bindings"
    addrBindings.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addrBindings.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addrBindings.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addrBindings.EntityData.Children = make(map[string]types.YChild)
    addrBindings.EntityData.Children["addr-binding"] = types.YChild{"AddrBinding", nil}
    for i := range addrBindings.AddrBinding {
        addrBindings.EntityData.Children[types.GetSegmentPath(&addrBindings.AddrBinding[i])] = types.YChild{"AddrBinding", &addrBindings.AddrBinding[i]}
    }
    addrBindings.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(addrBindings.EntityData)
}

// Dhcpv6_Nodes_Node_Base_AddrBindings_AddrBinding
// DHCPv6 base stats debug
type Dhcpv6_Nodes_Node_Base_AddrBindings_AddrBinding struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Address String. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    AddrString interface{}

    // DHCPV6 client MAC address. The type is string.
    MacAddress interface{}

    // DHCPV6 client/subscriber VRF name. The type is string with length: 0..33.
    VrfName interface{}

    // DHCPV6 server VRF name. The type is string with length: 0..33.
    ServerVrfName interface{}

    // DHCPV6 IPv6 address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6Address interface{}

    // DHCPV6 server IPv6 address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ServerIpv6Address interface{}

    // DHCPV6 reply server IPv6 address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
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
    // b'[a-zA-Z0-9./-]+'.
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
    addrBinding.EntityData.SegmentPath = "addr-binding" + "[addr-string='" + fmt.Sprintf("%v", addrBinding.AddrString) + "']"
    addrBinding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addrBinding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addrBinding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addrBinding.EntityData.Children = make(map[string]types.YChild)
    addrBinding.EntityData.Leafs = make(map[string]types.YLeaf)
    addrBinding.EntityData.Leafs["addr-string"] = types.YLeaf{"AddrString", addrBinding.AddrString}
    addrBinding.EntityData.Leafs["mac-address"] = types.YLeaf{"MacAddress", addrBinding.MacAddress}
    addrBinding.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", addrBinding.VrfName}
    addrBinding.EntityData.Leafs["server-vrf-name"] = types.YLeaf{"ServerVrfName", addrBinding.ServerVrfName}
    addrBinding.EntityData.Leafs["ipv6-address"] = types.YLeaf{"Ipv6Address", addrBinding.Ipv6Address}
    addrBinding.EntityData.Leafs["server-ipv6-address"] = types.YLeaf{"ServerIpv6Address", addrBinding.ServerIpv6Address}
    addrBinding.EntityData.Leafs["reply-server-ipv6-address"] = types.YLeaf{"ReplyServerIpv6Address", addrBinding.ReplyServerIpv6Address}
    addrBinding.EntityData.Leafs["lease-time"] = types.YLeaf{"LeaseTime", addrBinding.LeaseTime}
    addrBinding.EntityData.Leafs["remaining-lease-time"] = types.YLeaf{"RemainingLeaseTime", addrBinding.RemainingLeaseTime}
    addrBinding.EntityData.Leafs["state"] = types.YLeaf{"State", addrBinding.State}
    addrBinding.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", addrBinding.InterfaceName}
    addrBinding.EntityData.Leafs["access-vrf-name"] = types.YLeaf{"AccessVrfName", addrBinding.AccessVrfName}
    addrBinding.EntityData.Leafs["base-binding-tags"] = types.YLeaf{"BaseBindingTags", addrBinding.BaseBindingTags}
    addrBinding.EntityData.Leafs["base-binding-outer-tag"] = types.YLeaf{"BaseBindingOuterTag", addrBinding.BaseBindingOuterTag}
    addrBinding.EntityData.Leafs["base-binding-inner-tag"] = types.YLeaf{"BaseBindingInnerTag", addrBinding.BaseBindingInnerTag}
    addrBinding.EntityData.Leafs["profile-name"] = types.YLeaf{"ProfileName", addrBinding.ProfileName}
    addrBinding.EntityData.Leafs["is-nak-next-renew"] = types.YLeaf{"IsNakNextRenew", addrBinding.IsNakNextRenew}
    addrBinding.EntityData.Leafs["subscriber-label"] = types.YLeaf{"SubscriberLabel", addrBinding.SubscriberLabel}
    addrBinding.EntityData.Leafs["old-subscriber-label"] = types.YLeaf{"OldSubscriberLabel", addrBinding.OldSubscriberLabel}
    addrBinding.EntityData.Leafs["rx-client-duid"] = types.YLeaf{"RxClientDuid", addrBinding.RxClientDuid}
    addrBinding.EntityData.Leafs["tx-client-uid"] = types.YLeaf{"TxClientUid", addrBinding.TxClientUid}
    addrBinding.EntityData.Leafs["rx-remote-id"] = types.YLeaf{"RxRemoteId", addrBinding.RxRemoteId}
    addrBinding.EntityData.Leafs["tx-remote-id"] = types.YLeaf{"TxRemoteId", addrBinding.TxRemoteId}
    addrBinding.EntityData.Leafs["rx-interface-id"] = types.YLeaf{"RxInterfaceId", addrBinding.RxInterfaceId}
    addrBinding.EntityData.Leafs["tx-interface-id"] = types.YLeaf{"TxInterfaceId", addrBinding.TxInterfaceId}
    return &(addrBinding.EntityData)
}

// Dhcpv6_Nodes_Node_Server
// IPv6 DHCP server operational data
type Dhcpv6_Nodes_Node_Server struct {
    EntityData types.CommonEntityData
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

func (server *Dhcpv6_Nodes_Node_Server) GetEntityData() *types.CommonEntityData {
    server.EntityData.YFilter = server.YFilter
    server.EntityData.YangName = "server"
    server.EntityData.BundleName = "cisco_ios_xr"
    server.EntityData.ParentYangName = "node"
    server.EntityData.SegmentPath = "server"
    server.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    server.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    server.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    server.EntityData.Children = make(map[string]types.YChild)
    server.EntityData.Children["binding"] = types.YChild{"Binding", &server.Binding}
    server.EntityData.Children["vrfs"] = types.YChild{"Vrfs", &server.Vrfs}
    server.EntityData.Children["profiles"] = types.YChild{"Profiles", &server.Profiles}
    server.EntityData.Children["interfaces"] = types.YChild{"Interfaces", &server.Interfaces}
    server.EntityData.Children["statistics"] = types.YChild{"Statistics", &server.Statistics}
    server.EntityData.Children["binding-options"] = types.YChild{"BindingOptions", &server.BindingOptions}
    server.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(server.EntityData)
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
    binding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    binding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    binding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    binding.EntityData.Children = make(map[string]types.YChild)
    binding.EntityData.Children["summary"] = types.YChild{"Summary", &binding.Summary}
    binding.EntityData.Children["clients"] = types.YChild{"Clients", &binding.Clients}
    binding.EntityData.Leafs = make(map[string]types.YLeaf)
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
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = make(map[string]types.YChild)
    summary.EntityData.Children["iana"] = types.YChild{"Iana", &summary.Iana}
    summary.EntityData.Children["iapd"] = types.YChild{"Iapd", &summary.Iapd}
    summary.EntityData.Leafs = make(map[string]types.YLeaf)
    summary.EntityData.Leafs["clients"] = types.YLeaf{"Clients", summary.Clients}
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
    iana.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    iana.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    iana.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    iana.EntityData.Children = make(map[string]types.YChild)
    iana.EntityData.Leafs = make(map[string]types.YLeaf)
    iana.EntityData.Leafs["initializing-clients"] = types.YLeaf{"InitializingClients", iana.InitializingClients}
    iana.EntityData.Leafs["dpm-waiting-clients"] = types.YLeaf{"DpmWaitingClients", iana.DpmWaitingClients}
    iana.EntityData.Leafs["daps-waiting-clients"] = types.YLeaf{"DapsWaitingClients", iana.DapsWaitingClients}
    iana.EntityData.Leafs["request-waiting-clients"] = types.YLeaf{"RequestWaitingClients", iana.RequestWaitingClients}
    iana.EntityData.Leafs["iedge-waiting-clients"] = types.YLeaf{"IedgeWaitingClients", iana.IedgeWaitingClients}
    iana.EntityData.Leafs["rib-waiting-clients"] = types.YLeaf{"RibWaitingClients", iana.RibWaitingClients}
    iana.EntityData.Leafs["bound-clients"] = types.YLeaf{"BoundClients", iana.BoundClients}
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
    iapd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    iapd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    iapd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    iapd.EntityData.Children = make(map[string]types.YChild)
    iapd.EntityData.Leafs = make(map[string]types.YLeaf)
    iapd.EntityData.Leafs["initializing-clients"] = types.YLeaf{"InitializingClients", iapd.InitializingClients}
    iapd.EntityData.Leafs["dpm-waiting-clients"] = types.YLeaf{"DpmWaitingClients", iapd.DpmWaitingClients}
    iapd.EntityData.Leafs["daps-waiting-clients"] = types.YLeaf{"DapsWaitingClients", iapd.DapsWaitingClients}
    iapd.EntityData.Leafs["request-waiting-clients"] = types.YLeaf{"RequestWaitingClients", iapd.RequestWaitingClients}
    iapd.EntityData.Leafs["iedge-waiting-clients"] = types.YLeaf{"IedgeWaitingClients", iapd.IedgeWaitingClients}
    iapd.EntityData.Leafs["rib-waiting-clients"] = types.YLeaf{"RibWaitingClients", iapd.RibWaitingClients}
    iapd.EntityData.Leafs["bound-clients"] = types.YLeaf{"BoundClients", iapd.BoundClients}
    return &(iapd.EntityData)
}

// Dhcpv6_Nodes_Node_Server_Binding_Clients
// DHCPV6 server client bindings
type Dhcpv6_Nodes_Node_Server_Binding_Clients struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Single DHCPV6 server binding. The type is slice of
    // Dhcpv6_Nodes_Node_Server_Binding_Clients_Client.
    Client []Dhcpv6_Nodes_Node_Server_Binding_Clients_Client
}

func (clients *Dhcpv6_Nodes_Node_Server_Binding_Clients) GetEntityData() *types.CommonEntityData {
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

// Dhcpv6_Nodes_Node_Server_Binding_Clients_Client
// Single DHCPV6 server binding
type Dhcpv6_Nodes_Node_Server_Binding_Clients_Client struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Client ID. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    LinkLocalAddress interface{}

    // DHCPV6 access interface to client. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
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
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
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

func (client *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client) GetEntityData() *types.CommonEntityData {
    client.EntityData.YFilter = client.YFilter
    client.EntityData.YangName = "client"
    client.EntityData.BundleName = "cisco_ios_xr"
    client.EntityData.ParentYangName = "clients"
    client.EntityData.SegmentPath = "client" + "[client-id='" + fmt.Sprintf("%v", client.ClientId) + "']"
    client.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    client.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    client.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    client.EntityData.Children = make(map[string]types.YChild)
    client.EntityData.Children["ia-id-pd"] = types.YChild{"IaIdPd", &client.IaIdPd}
    client.EntityData.Leafs = make(map[string]types.YLeaf)
    client.EntityData.Leafs["client-id"] = types.YLeaf{"ClientId", client.ClientId}
    client.EntityData.Leafs["duid"] = types.YLeaf{"Duid", client.Duid}
    client.EntityData.Leafs["client-id-xr"] = types.YLeaf{"ClientIdXr", client.ClientIdXr}
    client.EntityData.Leafs["client-flag"] = types.YLeaf{"ClientFlag", client.ClientFlag}
    client.EntityData.Leafs["subscriber-label"] = types.YLeaf{"SubscriberLabel", client.SubscriberLabel}
    client.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", client.VrfName}
    client.EntityData.Leafs["mac-address"] = types.YLeaf{"MacAddress", client.MacAddress}
    client.EntityData.Leafs["ia-id-p-ds"] = types.YLeaf{"IaIdPDs", client.IaIdPDs}
    client.EntityData.Leafs["link-local-address"] = types.YLeaf{"LinkLocalAddress", client.LinkLocalAddress}
    client.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", client.InterfaceName}
    client.EntityData.Leafs["access-vrf-name"] = types.YLeaf{"AccessVrfName", client.AccessVrfName}
    client.EntityData.Leafs["server-binding-tags"] = types.YLeaf{"ServerBindingTags", client.ServerBindingTags}
    client.EntityData.Leafs["server-binding-outer-tag"] = types.YLeaf{"ServerBindingOuterTag", client.ServerBindingOuterTag}
    client.EntityData.Leafs["server-binding-inner-tag"] = types.YLeaf{"ServerBindingInnerTag", client.ServerBindingInnerTag}
    client.EntityData.Leafs["pool-name"] = types.YLeaf{"PoolName", client.PoolName}
    client.EntityData.Leafs["profile-name"] = types.YLeaf{"ProfileName", client.ProfileName}
    client.EntityData.Leafs["framed-ipv6-prefix"] = types.YLeaf{"FramedIpv6Prefix", client.FramedIpv6Prefix}
    client.EntityData.Leafs["framed-prefix-length"] = types.YLeaf{"FramedPrefixLength", client.FramedPrefixLength}
    client.EntityData.Leafs["class-name"] = types.YLeaf{"ClassName", client.ClassName}
    client.EntityData.Leafs["rx-remote-id"] = types.YLeaf{"RxRemoteId", client.RxRemoteId}
    client.EntityData.Leafs["rx-interface-id"] = types.YLeaf{"RxInterfaceId", client.RxInterfaceId}
    client.EntityData.Leafs["prefix-pool-name"] = types.YLeaf{"PrefixPoolName", client.PrefixPoolName}
    client.EntityData.Leafs["address-pool-name"] = types.YLeaf{"AddressPoolName", client.AddressPoolName}
    client.EntityData.Leafs["dns-server-count"] = types.YLeaf{"DnsServerCount", client.DnsServerCount}
    client.EntityData.Leafs["is-nak-next-renew"] = types.YLeaf{"IsNakNextRenew", client.IsNakNextRenew}
    client.EntityData.Leafs["srg-state"] = types.YLeaf{"SrgState", client.SrgState}
    client.EntityData.Leafs["srg-intf-role"] = types.YLeaf{"SrgIntfRole", client.SrgIntfRole}
    client.EntityData.Leafs["srgp2p"] = types.YLeaf{"Srgp2P", client.Srgp2P}
    client.EntityData.Leafs["srg-vrf-name"] = types.YLeaf{"SrgVrfName", client.SrgVrfName}
    client.EntityData.Leafs["sesrg-state"] = types.YLeaf{"SesrgState", client.SesrgState}
    client.EntityData.Leafs["serg-intf-role"] = types.YLeaf{"SergIntfRole", client.SergIntfRole}
    return &(client.EntityData)
}

// Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd
// List of DHCPv6 IA_ID/PDs
type Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bag dhcpv6d ia id pd info. The type is slice of
    // Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo.
    BagDhcpv6DIaIdPdInfo []Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo
}

func (iaIdPd *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd) GetEntityData() *types.CommonEntityData {
    iaIdPd.EntityData.YFilter = iaIdPd.YFilter
    iaIdPd.EntityData.YangName = "ia-id-pd"
    iaIdPd.EntityData.BundleName = "cisco_ios_xr"
    iaIdPd.EntityData.ParentYangName = "client"
    iaIdPd.EntityData.SegmentPath = "ia-id-pd"
    iaIdPd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    iaIdPd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    iaIdPd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    iaIdPd.EntityData.Children = make(map[string]types.YChild)
    iaIdPd.EntityData.Children["bag-dhcpv6d-ia-id-pd-info"] = types.YChild{"BagDhcpv6DIaIdPdInfo", nil}
    for i := range iaIdPd.BagDhcpv6DIaIdPdInfo {
        iaIdPd.EntityData.Children[types.GetSegmentPath(&iaIdPd.BagDhcpv6DIaIdPdInfo[i])] = types.YChild{"BagDhcpv6DIaIdPdInfo", &iaIdPd.BagDhcpv6DIaIdPdInfo[i]}
    }
    iaIdPd.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(iaIdPd.EntityData)
}

// Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo
// bag dhcpv6d ia id pd info
type Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo struct {
    EntityData types.CommonEntityData
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

func (bagDhcpv6DIaIdPdInfo *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo) GetEntityData() *types.CommonEntityData {
    bagDhcpv6DIaIdPdInfo.EntityData.YFilter = bagDhcpv6DIaIdPdInfo.YFilter
    bagDhcpv6DIaIdPdInfo.EntityData.YangName = "bag-dhcpv6d-ia-id-pd-info"
    bagDhcpv6DIaIdPdInfo.EntityData.BundleName = "cisco_ios_xr"
    bagDhcpv6DIaIdPdInfo.EntityData.ParentYangName = "ia-id-pd"
    bagDhcpv6DIaIdPdInfo.EntityData.SegmentPath = "bag-dhcpv6d-ia-id-pd-info"
    bagDhcpv6DIaIdPdInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bagDhcpv6DIaIdPdInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bagDhcpv6DIaIdPdInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bagDhcpv6DIaIdPdInfo.EntityData.Children = make(map[string]types.YChild)
    bagDhcpv6DIaIdPdInfo.EntityData.Children["addresses"] = types.YChild{"Addresses", &bagDhcpv6DIaIdPdInfo.Addresses}
    bagDhcpv6DIaIdPdInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    bagDhcpv6DIaIdPdInfo.EntityData.Leafs["ia-type"] = types.YLeaf{"IaType", bagDhcpv6DIaIdPdInfo.IaType}
    bagDhcpv6DIaIdPdInfo.EntityData.Leafs["ia-id"] = types.YLeaf{"IaId", bagDhcpv6DIaIdPdInfo.IaId}
    bagDhcpv6DIaIdPdInfo.EntityData.Leafs["flags"] = types.YLeaf{"Flags", bagDhcpv6DIaIdPdInfo.Flags}
    bagDhcpv6DIaIdPdInfo.EntityData.Leafs["total-address"] = types.YLeaf{"TotalAddress", bagDhcpv6DIaIdPdInfo.TotalAddress}
    bagDhcpv6DIaIdPdInfo.EntityData.Leafs["state"] = types.YLeaf{"State", bagDhcpv6DIaIdPdInfo.State}
    return &(bagDhcpv6DIaIdPdInfo.EntityData)
}

// Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses
// List of addresses in this IA
type Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bag dhcpv6d addr attrb. The type is slice of
    // Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses_BagDhcpv6DAddrAttrb.
    BagDhcpv6DAddrAttrb []Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses_BagDhcpv6DAddrAttrb
}

func (addresses *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses) GetEntityData() *types.CommonEntityData {
    addresses.EntityData.YFilter = addresses.YFilter
    addresses.EntityData.YangName = "addresses"
    addresses.EntityData.BundleName = "cisco_ios_xr"
    addresses.EntityData.ParentYangName = "bag-dhcpv6d-ia-id-pd-info"
    addresses.EntityData.SegmentPath = "addresses"
    addresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addresses.EntityData.Children = make(map[string]types.YChild)
    addresses.EntityData.Children["bag-dhcpv6d-addr-attrb"] = types.YChild{"BagDhcpv6DAddrAttrb", nil}
    for i := range addresses.BagDhcpv6DAddrAttrb {
        addresses.EntityData.Children[types.GetSegmentPath(&addresses.BagDhcpv6DAddrAttrb[i])] = types.YChild{"BagDhcpv6DAddrAttrb", &addresses.BagDhcpv6DAddrAttrb[i]}
    }
    addresses.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(addresses.EntityData)
}

// Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses_BagDhcpv6DAddrAttrb
// bag dhcpv6d addr attrb
type Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses_BagDhcpv6DAddrAttrb struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 prefix. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
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

func (bagDhcpv6DAddrAttrb *Dhcpv6_Nodes_Node_Server_Binding_Clients_Client_IaIdPd_BagDhcpv6DIaIdPdInfo_Addresses_BagDhcpv6DAddrAttrb) GetEntityData() *types.CommonEntityData {
    bagDhcpv6DAddrAttrb.EntityData.YFilter = bagDhcpv6DAddrAttrb.YFilter
    bagDhcpv6DAddrAttrb.EntityData.YangName = "bag-dhcpv6d-addr-attrb"
    bagDhcpv6DAddrAttrb.EntityData.BundleName = "cisco_ios_xr"
    bagDhcpv6DAddrAttrb.EntityData.ParentYangName = "addresses"
    bagDhcpv6DAddrAttrb.EntityData.SegmentPath = "bag-dhcpv6d-addr-attrb"
    bagDhcpv6DAddrAttrb.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bagDhcpv6DAddrAttrb.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bagDhcpv6DAddrAttrb.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bagDhcpv6DAddrAttrb.EntityData.Children = make(map[string]types.YChild)
    bagDhcpv6DAddrAttrb.EntityData.Leafs = make(map[string]types.YLeaf)
    bagDhcpv6DAddrAttrb.EntityData.Leafs["prefix"] = types.YLeaf{"Prefix", bagDhcpv6DAddrAttrb.Prefix}
    bagDhcpv6DAddrAttrb.EntityData.Leafs["prefix-length"] = types.YLeaf{"PrefixLength", bagDhcpv6DAddrAttrb.PrefixLength}
    bagDhcpv6DAddrAttrb.EntityData.Leafs["lease-time"] = types.YLeaf{"LeaseTime", bagDhcpv6DAddrAttrb.LeaseTime}
    bagDhcpv6DAddrAttrb.EntityData.Leafs["remaining-lease-time"] = types.YLeaf{"RemainingLeaseTime", bagDhcpv6DAddrAttrb.RemainingLeaseTime}
    return &(bagDhcpv6DAddrAttrb.EntityData)
}

// Dhcpv6_Nodes_Node_Server_Vrfs
// DHCPV6 server list of VRF names
type Dhcpv6_Nodes_Node_Server_Vrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 DHCP server VRF name. The type is slice of
    // Dhcpv6_Nodes_Node_Server_Vrfs_Vrf.
    Vrf []Dhcpv6_Nodes_Node_Server_Vrfs_Vrf
}

func (vrfs *Dhcpv6_Nodes_Node_Server_Vrfs) GetEntityData() *types.CommonEntityData {
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

// Dhcpv6_Nodes_Node_Server_Vrfs_Vrf
// IPv6 DHCP server VRF name
type Dhcpv6_Nodes_Node_Server_Vrfs_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    VrfName interface{}

    // IPv6 DHCP server statistics.
    Statistics Dhcpv6_Nodes_Node_Server_Vrfs_Vrf_Statistics
}

func (vrf *Dhcpv6_Nodes_Node_Server_Vrfs_Vrf) GetEntityData() *types.CommonEntityData {
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
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = make(map[string]types.YChild)
    statistics.EntityData.Children["solicit"] = types.YChild{"Solicit", &statistics.Solicit}
    statistics.EntityData.Children["advertise"] = types.YChild{"Advertise", &statistics.Advertise}
    statistics.EntityData.Children["request"] = types.YChild{"Request", &statistics.Request}
    statistics.EntityData.Children["reply"] = types.YChild{"Reply", &statistics.Reply}
    statistics.EntityData.Children["confirm"] = types.YChild{"Confirm", &statistics.Confirm}
    statistics.EntityData.Children["decline"] = types.YChild{"Decline", &statistics.Decline}
    statistics.EntityData.Children["renew"] = types.YChild{"Renew", &statistics.Renew}
    statistics.EntityData.Children["rebind"] = types.YChild{"Rebind", &statistics.Rebind}
    statistics.EntityData.Children["release"] = types.YChild{"Release", &statistics.Release}
    statistics.EntityData.Children["reconfig"] = types.YChild{"Reconfig", &statistics.Reconfig}
    statistics.EntityData.Children["inform"] = types.YChild{"Inform", &statistics.Inform}
    statistics.EntityData.Children["relay-forward"] = types.YChild{"RelayForward", &statistics.RelayForward}
    statistics.EntityData.Children["relay-reply"] = types.YChild{"RelayReply", &statistics.RelayReply}
    statistics.EntityData.Children["lease-query"] = types.YChild{"LeaseQuery", &statistics.LeaseQuery}
    statistics.EntityData.Children["lease-query-reply"] = types.YChild{"LeaseQueryReply", &statistics.LeaseQueryReply}
    statistics.EntityData.Children["lease-query-done"] = types.YChild{"LeaseQueryDone", &statistics.LeaseQueryDone}
    statistics.EntityData.Children["lease-query-data"] = types.YChild{"LeaseQueryData", &statistics.LeaseQueryData}
    statistics.EntityData.Leafs = make(map[string]types.YLeaf)
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
    solicit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    solicit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    solicit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    solicit.EntityData.Children = make(map[string]types.YChild)
    solicit.EntityData.Leafs = make(map[string]types.YLeaf)
    solicit.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", solicit.ReceivedPackets}
    solicit.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", solicit.TransmittedPackets}
    solicit.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", solicit.DroppedPackets}
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
    advertise.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    advertise.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    advertise.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    advertise.EntityData.Children = make(map[string]types.YChild)
    advertise.EntityData.Leafs = make(map[string]types.YLeaf)
    advertise.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", advertise.ReceivedPackets}
    advertise.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", advertise.TransmittedPackets}
    advertise.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", advertise.DroppedPackets}
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
    reply.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reply.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reply.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reply.EntityData.Children = make(map[string]types.YChild)
    reply.EntityData.Leafs = make(map[string]types.YLeaf)
    reply.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", reply.ReceivedPackets}
    reply.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", reply.TransmittedPackets}
    reply.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", reply.DroppedPackets}
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
    confirm.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    confirm.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    confirm.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    confirm.EntityData.Children = make(map[string]types.YChild)
    confirm.EntityData.Leafs = make(map[string]types.YLeaf)
    confirm.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", confirm.ReceivedPackets}
    confirm.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", confirm.TransmittedPackets}
    confirm.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", confirm.DroppedPackets}
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
    renew.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    renew.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    renew.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    renew.EntityData.Children = make(map[string]types.YChild)
    renew.EntityData.Leafs = make(map[string]types.YLeaf)
    renew.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", renew.ReceivedPackets}
    renew.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", renew.TransmittedPackets}
    renew.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", renew.DroppedPackets}
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
    rebind.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rebind.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rebind.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rebind.EntityData.Children = make(map[string]types.YChild)
    rebind.EntityData.Leafs = make(map[string]types.YLeaf)
    rebind.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", rebind.ReceivedPackets}
    rebind.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", rebind.TransmittedPackets}
    rebind.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", rebind.DroppedPackets}
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
    reconfig.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reconfig.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reconfig.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reconfig.EntityData.Children = make(map[string]types.YChild)
    reconfig.EntityData.Leafs = make(map[string]types.YLeaf)
    reconfig.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", reconfig.ReceivedPackets}
    reconfig.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", reconfig.TransmittedPackets}
    reconfig.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", reconfig.DroppedPackets}
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
    relayForward.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    relayForward.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    relayForward.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    relayForward.EntityData.Children = make(map[string]types.YChild)
    relayForward.EntityData.Leafs = make(map[string]types.YLeaf)
    relayForward.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", relayForward.ReceivedPackets}
    relayForward.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", relayForward.TransmittedPackets}
    relayForward.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", relayForward.DroppedPackets}
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
    relayReply.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    relayReply.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    relayReply.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    relayReply.EntityData.Children = make(map[string]types.YChild)
    relayReply.EntityData.Leafs = make(map[string]types.YLeaf)
    relayReply.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", relayReply.ReceivedPackets}
    relayReply.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", relayReply.TransmittedPackets}
    relayReply.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", relayReply.DroppedPackets}
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
    leaseQueryReply.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    leaseQueryReply.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    leaseQueryReply.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    leaseQueryReply.EntityData.Children = make(map[string]types.YChild)
    leaseQueryReply.EntityData.Leafs = make(map[string]types.YLeaf)
    leaseQueryReply.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", leaseQueryReply.ReceivedPackets}
    leaseQueryReply.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", leaseQueryReply.TransmittedPackets}
    leaseQueryReply.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", leaseQueryReply.DroppedPackets}
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
    leaseQueryDone.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    leaseQueryDone.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    leaseQueryDone.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    leaseQueryDone.EntityData.Children = make(map[string]types.YChild)
    leaseQueryDone.EntityData.Leafs = make(map[string]types.YLeaf)
    leaseQueryDone.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", leaseQueryDone.ReceivedPackets}
    leaseQueryDone.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", leaseQueryDone.TransmittedPackets}
    leaseQueryDone.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", leaseQueryDone.DroppedPackets}
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
    leaseQueryData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    leaseQueryData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    leaseQueryData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    leaseQueryData.EntityData.Children = make(map[string]types.YChild)
    leaseQueryData.EntityData.Leafs = make(map[string]types.YLeaf)
    leaseQueryData.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", leaseQueryData.ReceivedPackets}
    leaseQueryData.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", leaseQueryData.TransmittedPackets}
    leaseQueryData.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", leaseQueryData.DroppedPackets}
    return &(leaseQueryData.EntityData)
}

// Dhcpv6_Nodes_Node_Server_Profiles
// IPv6 DHCP server profile
type Dhcpv6_Nodes_Node_Server_Profiles struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 DHCP server profile. The type is slice of
    // Dhcpv6_Nodes_Node_Server_Profiles_Profile.
    Profile []Dhcpv6_Nodes_Node_Server_Profiles_Profile
}

func (profiles *Dhcpv6_Nodes_Node_Server_Profiles) GetEntityData() *types.CommonEntityData {
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

// Dhcpv6_Nodes_Node_Server_Profiles_Profile
// IPv6 DHCP server profile
type Dhcpv6_Nodes_Node_Server_Profiles_Profile struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Profile name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    ProfileName interface{}

    // IPv6 DHCP server profile Info.
    Info Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info

    // DHCPV6 throttle table.
    ThrottleInfos Dhcpv6_Nodes_Node_Server_Profiles_Profile_ThrottleInfos
}

func (profile *Dhcpv6_Nodes_Node_Server_Profiles_Profile) GetEntityData() *types.CommonEntityData {
    profile.EntityData.YFilter = profile.YFilter
    profile.EntityData.YangName = "profile"
    profile.EntityData.BundleName = "cisco_ios_xr"
    profile.EntityData.ParentYangName = "profiles"
    profile.EntityData.SegmentPath = "profile" + "[profile-name='" + fmt.Sprintf("%v", profile.ProfileName) + "']"
    profile.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    profile.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    profile.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    profile.EntityData.Children = make(map[string]types.YChild)
    profile.EntityData.Children["info"] = types.YChild{"Info", &profile.Info}
    profile.EntityData.Children["throttle-infos"] = types.YChild{"ThrottleInfos", &profile.ThrottleInfos}
    profile.EntityData.Leafs = make(map[string]types.YLeaf)
    profile.EntityData.Leafs["profile-name"] = types.YLeaf{"ProfileName", profile.ProfileName}
    return &(profile.EntityData)
}

// Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info
// IPv6 DHCP server profile Info
type Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info struct {
    EntityData types.CommonEntityData
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
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
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
    info.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    info.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    info.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    info.EntityData.Children = make(map[string]types.YChild)
    info.EntityData.Children["lease"] = types.YChild{"Lease", &info.Lease}
    info.EntityData.Children["interface-references"] = types.YChild{"InterfaceReferences", &info.InterfaceReferences}
    info.EntityData.Leafs = make(map[string]types.YLeaf)
    info.EntityData.Leafs["profile-name"] = types.YLeaf{"ProfileName", info.ProfileName}
    info.EntityData.Leafs["domain-name"] = types.YLeaf{"DomainName", info.DomainName}
    info.EntityData.Leafs["profile-dns"] = types.YLeaf{"ProfileDns", info.ProfileDns}
    info.EntityData.Leafs["aftr-name"] = types.YLeaf{"AftrName", info.AftrName}
    info.EntityData.Leafs["framed-addr-pool-name"] = types.YLeaf{"FramedAddrPoolName", info.FramedAddrPoolName}
    info.EntityData.Leafs["delegated-prefix-pool-name"] = types.YLeaf{"DelegatedPrefixPoolName", info.DelegatedPrefixPoolName}
    info.EntityData.Leafs["rapid-commit"] = types.YLeaf{"RapidCommit", info.RapidCommit}
    info.EntityData.Leafs["profile-dns-address"] = types.YLeaf{"ProfileDnsAddress", info.ProfileDnsAddress}
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
    lease.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lease.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lease.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lease.EntityData.Children = make(map[string]types.YChild)
    lease.EntityData.Leafs = make(map[string]types.YLeaf)
    lease.EntityData.Leafs["seconds"] = types.YLeaf{"Seconds", lease.Seconds}
    lease.EntityData.Leafs["time"] = types.YLeaf{"Time", lease.Time}
    return &(lease.EntityData)
}

// Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_InterfaceReferences
// Interface references
type Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_InterfaceReferences struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ipv6 dhcpv6d server interface reference. The type is slice of
    // Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_InterfaceReferences_Ipv6Dhcpv6DServerInterfaceReference.
    Ipv6Dhcpv6DServerInterfaceReference []Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_InterfaceReferences_Ipv6Dhcpv6DServerInterfaceReference
}

func (interfaceReferences *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_InterfaceReferences) GetEntityData() *types.CommonEntityData {
    interfaceReferences.EntityData.YFilter = interfaceReferences.YFilter
    interfaceReferences.EntityData.YangName = "interface-references"
    interfaceReferences.EntityData.BundleName = "cisco_ios_xr"
    interfaceReferences.EntityData.ParentYangName = "info"
    interfaceReferences.EntityData.SegmentPath = "interface-references"
    interfaceReferences.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceReferences.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceReferences.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceReferences.EntityData.Children = make(map[string]types.YChild)
    interfaceReferences.EntityData.Children["ipv6-dhcpv6d-server-interface-reference"] = types.YChild{"Ipv6Dhcpv6DServerInterfaceReference", nil}
    for i := range interfaceReferences.Ipv6Dhcpv6DServerInterfaceReference {
        interfaceReferences.EntityData.Children[types.GetSegmentPath(&interfaceReferences.Ipv6Dhcpv6DServerInterfaceReference[i])] = types.YChild{"Ipv6Dhcpv6DServerInterfaceReference", &interfaceReferences.Ipv6Dhcpv6DServerInterfaceReference[i]}
    }
    interfaceReferences.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaceReferences.EntityData)
}

// Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_InterfaceReferences_Ipv6Dhcpv6DServerInterfaceReference
// ipv6 dhcpv6d server interface reference
type Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_InterfaceReferences_Ipv6Dhcpv6DServerInterfaceReference struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface name. The type is string with length: 0..65.
    ServerReferenceInterfaceName interface{}
}

func (ipv6Dhcpv6DServerInterfaceReference *Dhcpv6_Nodes_Node_Server_Profiles_Profile_Info_InterfaceReferences_Ipv6Dhcpv6DServerInterfaceReference) GetEntityData() *types.CommonEntityData {
    ipv6Dhcpv6DServerInterfaceReference.EntityData.YFilter = ipv6Dhcpv6DServerInterfaceReference.YFilter
    ipv6Dhcpv6DServerInterfaceReference.EntityData.YangName = "ipv6-dhcpv6d-server-interface-reference"
    ipv6Dhcpv6DServerInterfaceReference.EntityData.BundleName = "cisco_ios_xr"
    ipv6Dhcpv6DServerInterfaceReference.EntityData.ParentYangName = "interface-references"
    ipv6Dhcpv6DServerInterfaceReference.EntityData.SegmentPath = "ipv6-dhcpv6d-server-interface-reference"
    ipv6Dhcpv6DServerInterfaceReference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Dhcpv6DServerInterfaceReference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Dhcpv6DServerInterfaceReference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Dhcpv6DServerInterfaceReference.EntityData.Children = make(map[string]types.YChild)
    ipv6Dhcpv6DServerInterfaceReference.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv6Dhcpv6DServerInterfaceReference.EntityData.Leafs["server-reference-interface-name"] = types.YLeaf{"ServerReferenceInterfaceName", ipv6Dhcpv6DServerInterfaceReference.ServerReferenceInterfaceName}
    return &(ipv6Dhcpv6DServerInterfaceReference.EntityData)
}

// Dhcpv6_Nodes_Node_Server_Profiles_Profile_ThrottleInfos
// DHCPV6 throttle table
type Dhcpv6_Nodes_Node_Server_Profiles_Profile_ThrottleInfos struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 DHCP server profile Throttle Info. The type is slice of
    // Dhcpv6_Nodes_Node_Server_Profiles_Profile_ThrottleInfos_ThrottleInfo.
    ThrottleInfo []Dhcpv6_Nodes_Node_Server_Profiles_Profile_ThrottleInfos_ThrottleInfo
}

func (throttleInfos *Dhcpv6_Nodes_Node_Server_Profiles_Profile_ThrottleInfos) GetEntityData() *types.CommonEntityData {
    throttleInfos.EntityData.YFilter = throttleInfos.YFilter
    throttleInfos.EntityData.YangName = "throttle-infos"
    throttleInfos.EntityData.BundleName = "cisco_ios_xr"
    throttleInfos.EntityData.ParentYangName = "profile"
    throttleInfos.EntityData.SegmentPath = "throttle-infos"
    throttleInfos.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    throttleInfos.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    throttleInfos.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    throttleInfos.EntityData.Children = make(map[string]types.YChild)
    throttleInfos.EntityData.Children["throttle-info"] = types.YChild{"ThrottleInfo", nil}
    for i := range throttleInfos.ThrottleInfo {
        throttleInfos.EntityData.Children[types.GetSegmentPath(&throttleInfos.ThrottleInfo[i])] = types.YChild{"ThrottleInfo", &throttleInfos.ThrottleInfo[i]}
    }
    throttleInfos.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(throttleInfos.EntityData)
}

// Dhcpv6_Nodes_Node_Server_Profiles_Profile_ThrottleInfos_ThrottleInfo
// IPv6 DHCP server profile Throttle Info
type Dhcpv6_Nodes_Node_Server_Profiles_Profile_ThrottleInfos_ThrottleInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. MAC address. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    MacAddress interface{}

    // Client MAC address. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
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
    throttleInfo.EntityData.SegmentPath = "throttle-info" + "[mac-address='" + fmt.Sprintf("%v", throttleInfo.MacAddress) + "']"
    throttleInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    throttleInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    throttleInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    throttleInfo.EntityData.Children = make(map[string]types.YChild)
    throttleInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    throttleInfo.EntityData.Leafs["mac-address"] = types.YLeaf{"MacAddress", throttleInfo.MacAddress}
    throttleInfo.EntityData.Leafs["binding-chaddr"] = types.YLeaf{"BindingChaddr", throttleInfo.BindingChaddr}
    throttleInfo.EntityData.Leafs["ifname"] = types.YLeaf{"Ifname", throttleInfo.Ifname}
    throttleInfo.EntityData.Leafs["state"] = types.YLeaf{"State", throttleInfo.State}
    throttleInfo.EntityData.Leafs["time-left"] = types.YLeaf{"TimeLeft", throttleInfo.TimeLeft}
    return &(throttleInfo.EntityData)
}

// Dhcpv6_Nodes_Node_Server_Interfaces
// DHCPV6 server interface
type Dhcpv6_Nodes_Node_Server_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 DHCP server interface. The type is slice of
    // Dhcpv6_Nodes_Node_Server_Interfaces_Interface_.
    Interface_ []Dhcpv6_Nodes_Node_Server_Interfaces_Interface
}

func (interfaces *Dhcpv6_Nodes_Node_Server_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "server"
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

// Dhcpv6_Nodes_Node_Server_Interfaces_Interface
// IPv6 DHCP server interface
type Dhcpv6_Nodes_Node_Server_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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

func (self *Dhcpv6_Nodes_Node_Server_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
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
    self.EntityData.Leafs["server-vrf-name"] = types.YLeaf{"ServerVrfName", self.ServerVrfName}
    self.EntityData.Leafs["server-interface-mode"] = types.YLeaf{"ServerInterfaceMode", self.ServerInterfaceMode}
    self.EntityData.Leafs["is-server-interface-ambiguous"] = types.YLeaf{"IsServerInterfaceAmbiguous", self.IsServerInterfaceAmbiguous}
    self.EntityData.Leafs["server-interface-profile-name"] = types.YLeaf{"ServerInterfaceProfileName", self.ServerInterfaceProfileName}
    self.EntityData.Leafs["server-interface-lease-limit-type"] = types.YLeaf{"ServerInterfaceLeaseLimitType", self.ServerInterfaceLeaseLimitType}
    self.EntityData.Leafs["server-interface-lease-limits"] = types.YLeaf{"ServerInterfaceLeaseLimits", self.ServerInterfaceLeaseLimits}
    self.EntityData.Leafs["srg-role"] = types.YLeaf{"SrgRole", self.SrgRole}
    self.EntityData.Leafs["serg-role"] = types.YLeaf{"SergRole", self.SergRole}
    self.EntityData.Leafs["mac-throttle"] = types.YLeaf{"MacThrottle", self.MacThrottle}
    self.EntityData.Leafs["srg-vrf-name"] = types.YLeaf{"SrgVrfName", self.SrgVrfName}
    self.EntityData.Leafs["srgp2p"] = types.YLeaf{"Srgp2P", self.Srgp2P}
    return &(self.EntityData)
}

// Dhcpv6_Nodes_Node_Server_Statistics
// DHCPv6 server statistics
type Dhcpv6_Nodes_Node_Server_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ipv6 dhcpv6d server stat. The type is slice of
    // Dhcpv6_Nodes_Node_Server_Statistics_Ipv6Dhcpv6DServerStat.
    Ipv6Dhcpv6DServerStat []Dhcpv6_Nodes_Node_Server_Statistics_Ipv6Dhcpv6DServerStat
}

func (statistics *Dhcpv6_Nodes_Node_Server_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "server"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = make(map[string]types.YChild)
    statistics.EntityData.Children["ipv6-dhcpv6d-server-stat"] = types.YChild{"Ipv6Dhcpv6DServerStat", nil}
    for i := range statistics.Ipv6Dhcpv6DServerStat {
        statistics.EntityData.Children[types.GetSegmentPath(&statistics.Ipv6Dhcpv6DServerStat[i])] = types.YChild{"Ipv6Dhcpv6DServerStat", &statistics.Ipv6Dhcpv6DServerStat[i]}
    }
    statistics.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(statistics.EntityData)
}

// Dhcpv6_Nodes_Node_Server_Statistics_Ipv6Dhcpv6DServerStat
// ipv6 dhcpv6d server stat
type Dhcpv6_Nodes_Node_Server_Statistics_Ipv6Dhcpv6DServerStat struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DHCPv6 L3 VRF name. The type is string with length: 0..33.
    VrfName interface{}

    // Server statistics.
    Statistics Dhcpv6_Nodes_Node_Server_Statistics_Ipv6Dhcpv6DServerStat_Statistics_
}

func (ipv6Dhcpv6DServerStat *Dhcpv6_Nodes_Node_Server_Statistics_Ipv6Dhcpv6DServerStat) GetEntityData() *types.CommonEntityData {
    ipv6Dhcpv6DServerStat.EntityData.YFilter = ipv6Dhcpv6DServerStat.YFilter
    ipv6Dhcpv6DServerStat.EntityData.YangName = "ipv6-dhcpv6d-server-stat"
    ipv6Dhcpv6DServerStat.EntityData.BundleName = "cisco_ios_xr"
    ipv6Dhcpv6DServerStat.EntityData.ParentYangName = "statistics"
    ipv6Dhcpv6DServerStat.EntityData.SegmentPath = "ipv6-dhcpv6d-server-stat"
    ipv6Dhcpv6DServerStat.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Dhcpv6DServerStat.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Dhcpv6DServerStat.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Dhcpv6DServerStat.EntityData.Children = make(map[string]types.YChild)
    ipv6Dhcpv6DServerStat.EntityData.Children["statistics"] = types.YChild{"Statistics", &ipv6Dhcpv6DServerStat.Statistics}
    ipv6Dhcpv6DServerStat.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv6Dhcpv6DServerStat.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", ipv6Dhcpv6DServerStat.VrfName}
    return &(ipv6Dhcpv6DServerStat.EntityData)
}

// Dhcpv6_Nodes_Node_Server_Statistics_Ipv6Dhcpv6DServerStat_Statistics_
// Server statistics
type Dhcpv6_Nodes_Node_Server_Statistics_Ipv6Dhcpv6DServerStat_Statistics_ struct {
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

func (statistics_ *Dhcpv6_Nodes_Node_Server_Statistics_Ipv6Dhcpv6DServerStat_Statistics_) GetEntityData() *types.CommonEntityData {
    statistics_.EntityData.YFilter = statistics_.YFilter
    statistics_.EntityData.YangName = "statistics"
    statistics_.EntityData.BundleName = "cisco_ios_xr"
    statistics_.EntityData.ParentYangName = "ipv6-dhcpv6d-server-stat"
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
    bindingOptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bindingOptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bindingOptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bindingOptions.EntityData.Children = make(map[string]types.YChild)
    bindingOptions.EntityData.Children["mac-bind-options"] = types.YChild{"MacBindOptions", &bindingOptions.MacBindOptions}
    bindingOptions.EntityData.Children["duid-bind-options"] = types.YChild{"DuidBindOptions", &bindingOptions.DuidBindOptions}
    bindingOptions.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bindingOptions.EntityData)
}

// Dhcpv6_Nodes_Node_Server_BindingOptions_MacBindOptions
// DHCPv6 server binding from MAC address
type Dhcpv6_Nodes_Node_Server_BindingOptions_MacBindOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DHCPv6 server binding with options. The type is slice of
    // Dhcpv6_Nodes_Node_Server_BindingOptions_MacBindOptions_MacBindOption.
    MacBindOption []Dhcpv6_Nodes_Node_Server_BindingOptions_MacBindOptions_MacBindOption
}

func (macBindOptions *Dhcpv6_Nodes_Node_Server_BindingOptions_MacBindOptions) GetEntityData() *types.CommonEntityData {
    macBindOptions.EntityData.YFilter = macBindOptions.YFilter
    macBindOptions.EntityData.YangName = "mac-bind-options"
    macBindOptions.EntityData.BundleName = "cisco_ios_xr"
    macBindOptions.EntityData.ParentYangName = "binding-options"
    macBindOptions.EntityData.SegmentPath = "mac-bind-options"
    macBindOptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macBindOptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macBindOptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macBindOptions.EntityData.Children = make(map[string]types.YChild)
    macBindOptions.EntityData.Children["mac-bind-option"] = types.YChild{"MacBindOption", nil}
    for i := range macBindOptions.MacBindOption {
        macBindOptions.EntityData.Children[types.GetSegmentPath(&macBindOptions.MacBindOption[i])] = types.YChild{"MacBindOption", &macBindOptions.MacBindOption[i]}
    }
    macBindOptions.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(macBindOptions.EntityData)
}

// Dhcpv6_Nodes_Node_Server_BindingOptions_MacBindOptions_MacBindOption
// DHCPv6 server binding with options
type Dhcpv6_Nodes_Node_Server_BindingOptions_MacBindOptions_MacBindOption struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. MAC address. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    DnsAddress []interface{}
}

func (macBindOption *Dhcpv6_Nodes_Node_Server_BindingOptions_MacBindOptions_MacBindOption) GetEntityData() *types.CommonEntityData {
    macBindOption.EntityData.YFilter = macBindOption.YFilter
    macBindOption.EntityData.YangName = "mac-bind-option"
    macBindOption.EntityData.BundleName = "cisco_ios_xr"
    macBindOption.EntityData.ParentYangName = "mac-bind-options"
    macBindOption.EntityData.SegmentPath = "mac-bind-option" + "[mac-address='" + fmt.Sprintf("%v", macBindOption.MacAddress) + "']"
    macBindOption.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macBindOption.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macBindOption.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macBindOption.EntityData.Children = make(map[string]types.YChild)
    macBindOption.EntityData.Leafs = make(map[string]types.YLeaf)
    macBindOption.EntityData.Leafs["mac-address"] = types.YLeaf{"MacAddress", macBindOption.MacAddress}
    macBindOption.EntityData.Leafs["mac-address-xr"] = types.YLeaf{"MacAddressXr", macBindOption.MacAddressXr}
    macBindOption.EntityData.Leafs["duid-xr"] = types.YLeaf{"DuidXr", macBindOption.DuidXr}
    macBindOption.EntityData.Leafs["dns-count"] = types.YLeaf{"DnsCount", macBindOption.DnsCount}
    macBindOption.EntityData.Leafs["opt17"] = types.YLeaf{"Opt17", macBindOption.Opt17}
    macBindOption.EntityData.Leafs["dns-address"] = types.YLeaf{"DnsAddress", macBindOption.DnsAddress}
    return &(macBindOption.EntityData)
}

// Dhcpv6_Nodes_Node_Server_BindingOptions_DuidBindOptions
// DHCPv6 server binding from DUID
type Dhcpv6_Nodes_Node_Server_BindingOptions_DuidBindOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DHCPv6 server binding with options. The type is slice of
    // Dhcpv6_Nodes_Node_Server_BindingOptions_DuidBindOptions_DuidBindOption.
    DuidBindOption []Dhcpv6_Nodes_Node_Server_BindingOptions_DuidBindOptions_DuidBindOption
}

func (duidBindOptions *Dhcpv6_Nodes_Node_Server_BindingOptions_DuidBindOptions) GetEntityData() *types.CommonEntityData {
    duidBindOptions.EntityData.YFilter = duidBindOptions.YFilter
    duidBindOptions.EntityData.YangName = "duid-bind-options"
    duidBindOptions.EntityData.BundleName = "cisco_ios_xr"
    duidBindOptions.EntityData.ParentYangName = "binding-options"
    duidBindOptions.EntityData.SegmentPath = "duid-bind-options"
    duidBindOptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    duidBindOptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    duidBindOptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    duidBindOptions.EntityData.Children = make(map[string]types.YChild)
    duidBindOptions.EntityData.Children["duid-bind-option"] = types.YChild{"DuidBindOption", nil}
    for i := range duidBindOptions.DuidBindOption {
        duidBindOptions.EntityData.Children[types.GetSegmentPath(&duidBindOptions.DuidBindOption[i])] = types.YChild{"DuidBindOption", &duidBindOptions.DuidBindOption[i]}
    }
    duidBindOptions.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(duidBindOptions.EntityData)
}

// Dhcpv6_Nodes_Node_Server_BindingOptions_DuidBindOptions_DuidBindOption
// DHCPv6 server binding with options
type Dhcpv6_Nodes_Node_Server_BindingOptions_DuidBindOptions_DuidBindOption struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. DUID of Binding. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    DnsAddress []interface{}
}

func (duidBindOption *Dhcpv6_Nodes_Node_Server_BindingOptions_DuidBindOptions_DuidBindOption) GetEntityData() *types.CommonEntityData {
    duidBindOption.EntityData.YFilter = duidBindOption.YFilter
    duidBindOption.EntityData.YangName = "duid-bind-option"
    duidBindOption.EntityData.BundleName = "cisco_ios_xr"
    duidBindOption.EntityData.ParentYangName = "duid-bind-options"
    duidBindOption.EntityData.SegmentPath = "duid-bind-option" + "[duid='" + fmt.Sprintf("%v", duidBindOption.Duid) + "']"
    duidBindOption.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    duidBindOption.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    duidBindOption.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    duidBindOption.EntityData.Children = make(map[string]types.YChild)
    duidBindOption.EntityData.Leafs = make(map[string]types.YLeaf)
    duidBindOption.EntityData.Leafs["duid"] = types.YLeaf{"Duid", duidBindOption.Duid}
    duidBindOption.EntityData.Leafs["mac-address-xr"] = types.YLeaf{"MacAddressXr", duidBindOption.MacAddressXr}
    duidBindOption.EntityData.Leafs["duid-xr"] = types.YLeaf{"DuidXr", duidBindOption.DuidXr}
    duidBindOption.EntityData.Leafs["dns-count"] = types.YLeaf{"DnsCount", duidBindOption.DnsCount}
    duidBindOption.EntityData.Leafs["opt17"] = types.YLeaf{"Opt17", duidBindOption.Opt17}
    duidBindOption.EntityData.Leafs["dns-address"] = types.YLeaf{"DnsAddress", duidBindOption.DnsAddress}
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
    relay.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    relay.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    relay.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    relay.EntityData.Children = make(map[string]types.YChild)
    relay.EntityData.Children["statistics"] = types.YChild{"Statistics", &relay.Statistics}
    relay.EntityData.Children["binding"] = types.YChild{"Binding", &relay.Binding}
    relay.EntityData.Children["vrfs"] = types.YChild{"Vrfs", &relay.Vrfs}
    relay.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(relay.EntityData)
}

// Dhcpv6_Nodes_Node_Relay_Statistics
// DHCPv6 relay statistics
type Dhcpv6_Nodes_Node_Relay_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ipv6 dhcpv6d relay stat. The type is slice of
    // Dhcpv6_Nodes_Node_Relay_Statistics_Ipv6Dhcpv6DRelayStat.
    Ipv6Dhcpv6DRelayStat []Dhcpv6_Nodes_Node_Relay_Statistics_Ipv6Dhcpv6DRelayStat
}

func (statistics *Dhcpv6_Nodes_Node_Relay_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "relay"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = make(map[string]types.YChild)
    statistics.EntityData.Children["ipv6-dhcpv6d-relay-stat"] = types.YChild{"Ipv6Dhcpv6DRelayStat", nil}
    for i := range statistics.Ipv6Dhcpv6DRelayStat {
        statistics.EntityData.Children[types.GetSegmentPath(&statistics.Ipv6Dhcpv6DRelayStat[i])] = types.YChild{"Ipv6Dhcpv6DRelayStat", &statistics.Ipv6Dhcpv6DRelayStat[i]}
    }
    statistics.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(statistics.EntityData)
}

// Dhcpv6_Nodes_Node_Relay_Statistics_Ipv6Dhcpv6DRelayStat
// ipv6 dhcpv6d relay stat
type Dhcpv6_Nodes_Node_Relay_Statistics_Ipv6Dhcpv6DRelayStat struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DHCPv6 L3 VRF name. The type is string with length: 0..33.
    VrfName interface{}

    // Relay statistics.
    Statistics Dhcpv6_Nodes_Node_Relay_Statistics_Ipv6Dhcpv6DRelayStat_Statistics_
}

func (ipv6Dhcpv6DRelayStat *Dhcpv6_Nodes_Node_Relay_Statistics_Ipv6Dhcpv6DRelayStat) GetEntityData() *types.CommonEntityData {
    ipv6Dhcpv6DRelayStat.EntityData.YFilter = ipv6Dhcpv6DRelayStat.YFilter
    ipv6Dhcpv6DRelayStat.EntityData.YangName = "ipv6-dhcpv6d-relay-stat"
    ipv6Dhcpv6DRelayStat.EntityData.BundleName = "cisco_ios_xr"
    ipv6Dhcpv6DRelayStat.EntityData.ParentYangName = "statistics"
    ipv6Dhcpv6DRelayStat.EntityData.SegmentPath = "ipv6-dhcpv6d-relay-stat"
    ipv6Dhcpv6DRelayStat.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Dhcpv6DRelayStat.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Dhcpv6DRelayStat.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Dhcpv6DRelayStat.EntityData.Children = make(map[string]types.YChild)
    ipv6Dhcpv6DRelayStat.EntityData.Children["statistics"] = types.YChild{"Statistics", &ipv6Dhcpv6DRelayStat.Statistics}
    ipv6Dhcpv6DRelayStat.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv6Dhcpv6DRelayStat.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", ipv6Dhcpv6DRelayStat.VrfName}
    return &(ipv6Dhcpv6DRelayStat.EntityData)
}

// Dhcpv6_Nodes_Node_Relay_Statistics_Ipv6Dhcpv6DRelayStat_Statistics_
// Relay statistics
type Dhcpv6_Nodes_Node_Relay_Statistics_Ipv6Dhcpv6DRelayStat_Statistics_ struct {
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

func (statistics_ *Dhcpv6_Nodes_Node_Relay_Statistics_Ipv6Dhcpv6DRelayStat_Statistics_) GetEntityData() *types.CommonEntityData {
    statistics_.EntityData.YFilter = statistics_.YFilter
    statistics_.EntityData.YangName = "statistics"
    statistics_.EntityData.BundleName = "cisco_ios_xr"
    statistics_.EntityData.ParentYangName = "ipv6-dhcpv6d-relay-stat"
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
    binding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    binding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    binding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    binding.EntityData.Children = make(map[string]types.YChild)
    binding.EntityData.Children["summary"] = types.YChild{"Summary", &binding.Summary}
    binding.EntityData.Children["clients"] = types.YChild{"Clients", &binding.Clients}
    binding.EntityData.Leafs = make(map[string]types.YLeaf)
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
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = make(map[string]types.YChild)
    summary.EntityData.Leafs = make(map[string]types.YLeaf)
    summary.EntityData.Leafs["clients"] = types.YLeaf{"Clients", summary.Clients}
    return &(summary.EntityData)
}

// Dhcpv6_Nodes_Node_Relay_Binding_Clients
// DHCPV6 relay client bindings
type Dhcpv6_Nodes_Node_Relay_Binding_Clients struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Single DHCPV6 relay binding. The type is slice of
    // Dhcpv6_Nodes_Node_Relay_Binding_Clients_Client.
    Client []Dhcpv6_Nodes_Node_Relay_Binding_Clients_Client
}

func (clients *Dhcpv6_Nodes_Node_Relay_Binding_Clients) GetEntityData() *types.CommonEntityData {
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

// Dhcpv6_Nodes_Node_Relay_Binding_Clients_Client
// Single DHCPV6 relay binding
type Dhcpv6_Nodes_Node_Relay_Binding_Clients_Client struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Client ID. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    ClientId interface{}

    // Client DUID. The type is string.
    Duid interface{}

    // Client unique identifier. The type is interface{} with range:
    // 0..4294967295.
    ClientIdXr interface{}

    // length of prefix. The type is interface{} with range: 0..255.
    PrefixLength interface{}

    // DHCPV6 IPv6 Prefix allotted to client. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
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
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
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
    client.EntityData.SegmentPath = "client" + "[client-id='" + fmt.Sprintf("%v", client.ClientId) + "']"
    client.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    client.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    client.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    client.EntityData.Children = make(map[string]types.YChild)
    client.EntityData.Leafs = make(map[string]types.YLeaf)
    client.EntityData.Leafs["client-id"] = types.YLeaf{"ClientId", client.ClientId}
    client.EntityData.Leafs["duid"] = types.YLeaf{"Duid", client.Duid}
    client.EntityData.Leafs["client-id-xr"] = types.YLeaf{"ClientIdXr", client.ClientIdXr}
    client.EntityData.Leafs["prefix-length"] = types.YLeaf{"PrefixLength", client.PrefixLength}
    client.EntityData.Leafs["prefix"] = types.YLeaf{"Prefix", client.Prefix}
    client.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", client.VrfName}
    client.EntityData.Leafs["lifetime"] = types.YLeaf{"Lifetime", client.Lifetime}
    client.EntityData.Leafs["rem-life-time"] = types.YLeaf{"RemLifeTime", client.RemLifeTime}
    client.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", client.InterfaceName}
    client.EntityData.Leafs["next-hop-addr"] = types.YLeaf{"NextHopAddr", client.NextHopAddr}
    client.EntityData.Leafs["ia-id"] = types.YLeaf{"IaId", client.IaId}
    client.EntityData.Leafs["relay-profile-name"] = types.YLeaf{"RelayProfileName", client.RelayProfileName}
    return &(client.EntityData)
}

// Dhcpv6_Nodes_Node_Relay_Vrfs
// DHCPV6 relay list of VRF names
type Dhcpv6_Nodes_Node_Relay_Vrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 DHCP relay VRF name. The type is slice of
    // Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf.
    Vrf []Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf
}

func (vrfs *Dhcpv6_Nodes_Node_Relay_Vrfs) GetEntityData() *types.CommonEntityData {
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

// Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf
// IPv6 DHCP relay VRF name
type Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    VrfName interface{}

    // IPv6 DHCP relay statistics.
    Statistics Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf_Statistics
}

func (vrf *Dhcpv6_Nodes_Node_Relay_Vrfs_Vrf) GetEntityData() *types.CommonEntityData {
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
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = make(map[string]types.YChild)
    statistics.EntityData.Children["solicit"] = types.YChild{"Solicit", &statistics.Solicit}
    statistics.EntityData.Children["advertise"] = types.YChild{"Advertise", &statistics.Advertise}
    statistics.EntityData.Children["request"] = types.YChild{"Request", &statistics.Request}
    statistics.EntityData.Children["reply"] = types.YChild{"Reply", &statistics.Reply}
    statistics.EntityData.Children["confirm"] = types.YChild{"Confirm", &statistics.Confirm}
    statistics.EntityData.Children["decline"] = types.YChild{"Decline", &statistics.Decline}
    statistics.EntityData.Children["renew"] = types.YChild{"Renew", &statistics.Renew}
    statistics.EntityData.Children["rebind"] = types.YChild{"Rebind", &statistics.Rebind}
    statistics.EntityData.Children["release"] = types.YChild{"Release", &statistics.Release}
    statistics.EntityData.Children["reconfig"] = types.YChild{"Reconfig", &statistics.Reconfig}
    statistics.EntityData.Children["inform"] = types.YChild{"Inform", &statistics.Inform}
    statistics.EntityData.Children["relay-forward"] = types.YChild{"RelayForward", &statistics.RelayForward}
    statistics.EntityData.Children["relay-reply"] = types.YChild{"RelayReply", &statistics.RelayReply}
    statistics.EntityData.Children["lease-query"] = types.YChild{"LeaseQuery", &statistics.LeaseQuery}
    statistics.EntityData.Children["lease-query-reply"] = types.YChild{"LeaseQueryReply", &statistics.LeaseQueryReply}
    statistics.EntityData.Children["lease-query-done"] = types.YChild{"LeaseQueryDone", &statistics.LeaseQueryDone}
    statistics.EntityData.Children["lease-query-data"] = types.YChild{"LeaseQueryData", &statistics.LeaseQueryData}
    statistics.EntityData.Leafs = make(map[string]types.YLeaf)
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
    solicit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    solicit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    solicit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    solicit.EntityData.Children = make(map[string]types.YChild)
    solicit.EntityData.Leafs = make(map[string]types.YLeaf)
    solicit.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", solicit.ReceivedPackets}
    solicit.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", solicit.TransmittedPackets}
    solicit.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", solicit.DroppedPackets}
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
    advertise.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    advertise.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    advertise.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    advertise.EntityData.Children = make(map[string]types.YChild)
    advertise.EntityData.Leafs = make(map[string]types.YLeaf)
    advertise.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", advertise.ReceivedPackets}
    advertise.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", advertise.TransmittedPackets}
    advertise.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", advertise.DroppedPackets}
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
    reply.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reply.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reply.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reply.EntityData.Children = make(map[string]types.YChild)
    reply.EntityData.Leafs = make(map[string]types.YLeaf)
    reply.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", reply.ReceivedPackets}
    reply.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", reply.TransmittedPackets}
    reply.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", reply.DroppedPackets}
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
    confirm.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    confirm.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    confirm.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    confirm.EntityData.Children = make(map[string]types.YChild)
    confirm.EntityData.Leafs = make(map[string]types.YLeaf)
    confirm.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", confirm.ReceivedPackets}
    confirm.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", confirm.TransmittedPackets}
    confirm.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", confirm.DroppedPackets}
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
    renew.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    renew.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    renew.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    renew.EntityData.Children = make(map[string]types.YChild)
    renew.EntityData.Leafs = make(map[string]types.YLeaf)
    renew.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", renew.ReceivedPackets}
    renew.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", renew.TransmittedPackets}
    renew.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", renew.DroppedPackets}
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
    rebind.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rebind.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rebind.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rebind.EntityData.Children = make(map[string]types.YChild)
    rebind.EntityData.Leafs = make(map[string]types.YLeaf)
    rebind.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", rebind.ReceivedPackets}
    rebind.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", rebind.TransmittedPackets}
    rebind.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", rebind.DroppedPackets}
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
    reconfig.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reconfig.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reconfig.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reconfig.EntityData.Children = make(map[string]types.YChild)
    reconfig.EntityData.Leafs = make(map[string]types.YLeaf)
    reconfig.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", reconfig.ReceivedPackets}
    reconfig.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", reconfig.TransmittedPackets}
    reconfig.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", reconfig.DroppedPackets}
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
    relayForward.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    relayForward.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    relayForward.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    relayForward.EntityData.Children = make(map[string]types.YChild)
    relayForward.EntityData.Leafs = make(map[string]types.YLeaf)
    relayForward.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", relayForward.ReceivedPackets}
    relayForward.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", relayForward.TransmittedPackets}
    relayForward.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", relayForward.DroppedPackets}
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
    relayReply.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    relayReply.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    relayReply.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    relayReply.EntityData.Children = make(map[string]types.YChild)
    relayReply.EntityData.Leafs = make(map[string]types.YLeaf)
    relayReply.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", relayReply.ReceivedPackets}
    relayReply.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", relayReply.TransmittedPackets}
    relayReply.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", relayReply.DroppedPackets}
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
    leaseQueryReply.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    leaseQueryReply.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    leaseQueryReply.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    leaseQueryReply.EntityData.Children = make(map[string]types.YChild)
    leaseQueryReply.EntityData.Leafs = make(map[string]types.YLeaf)
    leaseQueryReply.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", leaseQueryReply.ReceivedPackets}
    leaseQueryReply.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", leaseQueryReply.TransmittedPackets}
    leaseQueryReply.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", leaseQueryReply.DroppedPackets}
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
    leaseQueryDone.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    leaseQueryDone.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    leaseQueryDone.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    leaseQueryDone.EntityData.Children = make(map[string]types.YChild)
    leaseQueryDone.EntityData.Leafs = make(map[string]types.YLeaf)
    leaseQueryDone.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", leaseQueryDone.ReceivedPackets}
    leaseQueryDone.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", leaseQueryDone.TransmittedPackets}
    leaseQueryDone.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", leaseQueryDone.DroppedPackets}
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
    leaseQueryData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    leaseQueryData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    leaseQueryData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    leaseQueryData.EntityData.Children = make(map[string]types.YChild)
    leaseQueryData.EntityData.Leafs = make(map[string]types.YLeaf)
    leaseQueryData.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", leaseQueryData.ReceivedPackets}
    leaseQueryData.EntityData.Leafs["transmitted-packets"] = types.YLeaf{"TransmittedPackets", leaseQueryData.TransmittedPackets}
    leaseQueryData.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", leaseQueryData.DroppedPackets}
    return &(leaseQueryData.EntityData)
}

