// This module contains a collection of YANG definitions
// for DHCP Server and Client operational data.
// Copyright (c) 2017-2018 by Cisco Systems, Inc.
// All rights reserved.
package dhcp_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package dhcp_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-dhcp-oper dhcp-oper-data}", reflect.TypeOf(DhcpOperData{}))
    ydk.RegisterEntity("Cisco-IOS-XE-dhcp-oper:dhcp-oper-data", reflect.TypeOf(DhcpOperData{}))
}

// DhcpServerBindingType represents DHCP server binding type
type DhcpServerBindingType string

const (
    // Server binding Type Manual
    DhcpServerBindingType_dhcp_server_binding_type_manual DhcpServerBindingType = "dhcp-server-binding-type-manual"

    // Sever binding type Static
    DhcpServerBindingType_dhcp_server_binding_type_static DhcpServerBindingType = "dhcp-server-binding-type-static"

    // Server binding type relay
    DhcpServerBindingType_dhcp_server_binding_type_relay DhcpServerBindingType = "dhcp-server-binding-type-relay"

    // Server binding type automatic
    DhcpServerBindingType_dhcp_server_binding_type_automatic DhcpServerBindingType = "dhcp-server-binding-type-automatic"

    // Server binding Type ODAP
    DhcpServerBindingType_dhcp_server_binding_type_odap DhcpServerBindingType = "dhcp-server-binding-type-odap"

    // Sever binding type from AAA
    DhcpServerBindingType_dhcp_server_binding_type_from_aaa DhcpServerBindingType = "dhcp-server-binding-type-from-aaa"

    // Server binding type remembered
    DhcpServerBindingType_dhcp_server_binding_type_remembered DhcpServerBindingType = "dhcp-server-binding-type-remembered"
)

// DhcpExpiryOption represents DHCP expiration option 
type DhcpExpiryOption string

const (
    // Expiration option time
    DhcpExpiryOption_dhcp_expiration_time DhcpExpiryOption = "dhcp-expiration-time"

    // Expiration option infinite
    DhcpExpiryOption_dhcp_expiration_infinite DhcpExpiryOption = "dhcp-expiration-infinite"
)

// DhcpClientIdType represents DHCP Client id hardware types 
type DhcpClientIdType string

const (
    // DHCP Client hardware type Ethernet
    DhcpClientIdType_dhcp_htype_ethernet DhcpClientIdType = "dhcp-htype-ethernet"

    // DHCP Client hardware type 802
    DhcpClientIdType_dhcp_htype_ieee802 DhcpClientIdType = "dhcp-htype-ieee802"

    // DHCP Client hardware type RFCLIMIT
    DhcpClientIdType_dhcp_htype_rfclimit DhcpClientIdType = "dhcp-htype-rfclimit"

    // DHCP Client hardware type CLIENTID
    DhcpClientIdType_dhcp_htype_clientid DhcpClientIdType = "dhcp-htype-clientid"
)

// DhcpClientState represents DHCP Client state
type DhcpClientState string

const (
    // Client state address/sync from other client
    DhcpClientState_dhcp_client_state_temp_from_client DhcpClientState = "dhcp-client-state-temp-from-client"

    // Client state Sync
    DhcpClientState_dhcp_client_state_temp_from_sync DhcpClientState = "dhcp-client-state-temp-from-sync"

    // Client state Initialing
    DhcpClientState_dhcp_client_state_initial DhcpClientState = "dhcp-client-state-initial"

    // Client state Selecting
    DhcpClientState_dhcp_client_state_selecting DhcpClientState = "dhcp-client-state-selecting"

    // Client state Requesting
    DhcpClientState_dhcp_client_state_requesting DhcpClientState = "dhcp-client-state-requesting"

    // Client state bound
    DhcpClientState_dhcp_client_state_bound DhcpClientState = "dhcp-client-state-bound"

    // Client state rebinding
    DhcpClientState_dhcp_client_state_rebinding DhcpClientState = "dhcp-client-state-rebinding"

    // Client state renewing
    DhcpClientState_dhcp_client_state_renewing DhcpClientState = "dhcp-client-state-renewing"

    // Client state holdtime
    DhcpClientState_dhcp_client_state_holdtime DhcpClientState = "dhcp-client-state-holdtime"

    // Client state DDNS wait
    DhcpClientState_dhcp_client_state_ddns_wait DhcpClientState = "dhcp-client-state-ddns-wait"

    // Client state releasing
    DhcpClientState_dhcp_client_state_releasing DhcpClientState = "dhcp-client-state-releasing"

    // Client state purging
    DhcpClientState_dhcp_client_state_purging DhcpClientState = "dhcp-client-state-purging"

    // Client state leasequery
    DhcpClientState_dhcp_client_state_leasequery DhcpClientState = "dhcp-client-state-leasequery"

    // Client state unknown
    DhcpClientState_dhcp_client_state_unknown DhcpClientState = "dhcp-client-state-unknown"
)

// DhcpServerBindingState represents DHCP server binding states 
type DhcpServerBindingState string

const (
    // Server state is in Selecting mode
    DhcpServerBindingState_dhcp_server_binding_state_selecting DhcpServerBindingState = "dhcp-server-binding-state-selecting"

    // Server state Active new address provided
    DhcpServerBindingState_dhcp_server_binding_state_active DhcpServerBindingState = "dhcp-server-binding-state-active"

    // Server terminated the connection with a client
    DhcpServerBindingState_dhcp_server_binding_state_terminated DhcpServerBindingState = "dhcp-server-binding-state-terminated"

    // Server state unknown
    DhcpServerBindingState_dhcp_server_binding_state_unknown DhcpServerBindingState = "dhcp-server-binding-state-unknown"
)

// DhcpOperData
// Operational state of DHCP
type DhcpOperData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of DHCP server bidning. The type is slice of
    // DhcpOperData_Dhcpv4ServerOper.
    Dhcpv4ServerOper []*DhcpOperData_Dhcpv4ServerOper

    // List of DHCP clients. The type is slice of DhcpOperData_Dhcpv4ClientOper.
    Dhcpv4ClientOper []*DhcpOperData_Dhcpv4ClientOper
}

func (dhcpOperData *DhcpOperData) GetEntityData() *types.CommonEntityData {
    dhcpOperData.EntityData.YFilter = dhcpOperData.YFilter
    dhcpOperData.EntityData.YangName = "dhcp-oper-data"
    dhcpOperData.EntityData.BundleName = "cisco_ios_xe"
    dhcpOperData.EntityData.ParentYangName = "Cisco-IOS-XE-dhcp-oper"
    dhcpOperData.EntityData.SegmentPath = "Cisco-IOS-XE-dhcp-oper:dhcp-oper-data"
    dhcpOperData.EntityData.AbsolutePath = dhcpOperData.EntityData.SegmentPath
    dhcpOperData.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dhcpOperData.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dhcpOperData.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dhcpOperData.EntityData.Children = types.NewOrderedMap()
    dhcpOperData.EntityData.Children.Append("dhcpv4-server-oper", types.YChild{"Dhcpv4ServerOper", nil})
    for i := range dhcpOperData.Dhcpv4ServerOper {
        dhcpOperData.EntityData.Children.Append(types.GetSegmentPath(dhcpOperData.Dhcpv4ServerOper[i]), types.YChild{"Dhcpv4ServerOper", dhcpOperData.Dhcpv4ServerOper[i]})
    }
    dhcpOperData.EntityData.Children.Append("dhcpv4-client-oper", types.YChild{"Dhcpv4ClientOper", nil})
    for i := range dhcpOperData.Dhcpv4ClientOper {
        dhcpOperData.EntityData.Children.Append(types.GetSegmentPath(dhcpOperData.Dhcpv4ClientOper[i]), types.YChild{"Dhcpv4ClientOper", dhcpOperData.Dhcpv4ClientOper[i]})
    }
    dhcpOperData.EntityData.Leafs = types.NewOrderedMap()

    dhcpOperData.EntityData.YListKeys = []string {}

    return &(dhcpOperData.EntityData)
}

// DhcpOperData_Dhcpv4ServerOper
// List of DHCP server bidning
type DhcpOperData_Dhcpv4ServerOper struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Server Pool name from where the Client  ip is
    // provided. The type is string.
    PoolName interface{}

    // This attribute is a key. ipaddress released for a speicfic Client  from
    // Server. The type is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    ClientIp interface{}

    // This attribute is a key. Query based on the vrfname speicfic to that pool
    // and Client ip address as key. The type is string.
    VrfName interface{}

    // Client identification Hardware types. The type is DhcpClientIdType.
    ClientIdType interface{}

    // Client identification can be based on Hardware types/Mac. The type is
    // string.
    ClientId interface{}

    // Server binding type. The type is DhcpServerBindingType.
    Type interface{}

    // Server binding states. The type is DhcpServerBindingState.
    State interface{}

    // interface name of the pool. The type is string.
    Interface interface{}

    // Expiration time infomation.
    Expiration DhcpOperData_Dhcpv4ServerOper_Expiration
}

func (dhcpv4ServerOper *DhcpOperData_Dhcpv4ServerOper) GetEntityData() *types.CommonEntityData {
    dhcpv4ServerOper.EntityData.YFilter = dhcpv4ServerOper.YFilter
    dhcpv4ServerOper.EntityData.YangName = "dhcpv4-server-oper"
    dhcpv4ServerOper.EntityData.BundleName = "cisco_ios_xe"
    dhcpv4ServerOper.EntityData.ParentYangName = "dhcp-oper-data"
    dhcpv4ServerOper.EntityData.SegmentPath = "dhcpv4-server-oper" + types.AddKeyToken(dhcpv4ServerOper.PoolName, "pool-name") + types.AddKeyToken(dhcpv4ServerOper.ClientIp, "client-ip") + types.AddKeyToken(dhcpv4ServerOper.VrfName, "vrf-name")
    dhcpv4ServerOper.EntityData.AbsolutePath = "Cisco-IOS-XE-dhcp-oper:dhcp-oper-data/" + dhcpv4ServerOper.EntityData.SegmentPath
    dhcpv4ServerOper.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dhcpv4ServerOper.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dhcpv4ServerOper.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dhcpv4ServerOper.EntityData.Children = types.NewOrderedMap()
    dhcpv4ServerOper.EntityData.Children.Append("expiration", types.YChild{"Expiration", &dhcpv4ServerOper.Expiration})
    dhcpv4ServerOper.EntityData.Leafs = types.NewOrderedMap()
    dhcpv4ServerOper.EntityData.Leafs.Append("pool-name", types.YLeaf{"PoolName", dhcpv4ServerOper.PoolName})
    dhcpv4ServerOper.EntityData.Leafs.Append("client-ip", types.YLeaf{"ClientIp", dhcpv4ServerOper.ClientIp})
    dhcpv4ServerOper.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", dhcpv4ServerOper.VrfName})
    dhcpv4ServerOper.EntityData.Leafs.Append("client-id-type", types.YLeaf{"ClientIdType", dhcpv4ServerOper.ClientIdType})
    dhcpv4ServerOper.EntityData.Leafs.Append("client-id", types.YLeaf{"ClientId", dhcpv4ServerOper.ClientId})
    dhcpv4ServerOper.EntityData.Leafs.Append("type", types.YLeaf{"Type", dhcpv4ServerOper.Type})
    dhcpv4ServerOper.EntityData.Leafs.Append("state", types.YLeaf{"State", dhcpv4ServerOper.State})
    dhcpv4ServerOper.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", dhcpv4ServerOper.Interface})

    dhcpv4ServerOper.EntityData.YListKeys = []string {"PoolName", "ClientIp", "VrfName"}

    return &(dhcpv4ServerOper.EntityData)
}

// DhcpOperData_Dhcpv4ServerOper_Expiration
// Expiration time infomation
type DhcpOperData_Dhcpv4ServerOper_Expiration struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Date and time of expiry . The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    Time interface{}

    // Expiry time infinite. The type is interface{}.
    Infinite interface{}
}

func (expiration *DhcpOperData_Dhcpv4ServerOper_Expiration) GetEntityData() *types.CommonEntityData {
    expiration.EntityData.YFilter = expiration.YFilter
    expiration.EntityData.YangName = "expiration"
    expiration.EntityData.BundleName = "cisco_ios_xe"
    expiration.EntityData.ParentYangName = "dhcpv4-server-oper"
    expiration.EntityData.SegmentPath = "expiration"
    expiration.EntityData.AbsolutePath = "Cisco-IOS-XE-dhcp-oper:dhcp-oper-data/dhcpv4-server-oper/" + expiration.EntityData.SegmentPath
    expiration.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    expiration.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    expiration.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    expiration.EntityData.Children = types.NewOrderedMap()
    expiration.EntityData.Leafs = types.NewOrderedMap()
    expiration.EntityData.Leafs.Append("time", types.YLeaf{"Time", expiration.Time})
    expiration.EntityData.Leafs.Append("infinite", types.YLeaf{"Infinite", expiration.Infinite})

    expiration.EntityData.YListKeys = []string {}

    return &(expiration.EntityData)
}

// DhcpOperData_Dhcpv4ClientOper
// List of DHCP clients
type DhcpOperData_Dhcpv4ClientOper struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface infomation where dhcp Client is
    // configured. The type is string.
    IfName interface{}

    // This attribute is a key. Client_addr address Allocated from Server. The
    // type is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    ClientAddr interface{}

    // This attribute is a key. Vrfname infomation related to Client. The type is
    // string.
    VrfName interface{}

    // DHCP Client States . The type is DhcpClientState.
    State interface{}

    // IP address of Server from where we got IP. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    LeaseServerAddr interface{}

    // Gateway Address we got from Server. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    GatewayAddr interface{}

    // Total Lease Time in Seconds. The type is interface{} with range:
    // 0..4294967295.
    LeaseTime interface{}

    // Lease remaining time for the IP address. The type is interface{} with
    // range: 0..4294967295.
    LeaseRemaining interface{}

    // First DNS address. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    DnsAddress interface{}

    // Secondary DNS address. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    DnsAddressSecondary interface{}

    // Subnet mask address. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SubnetMask interface{}

    // Lease Expiry time for the IP address we got.
    LeaseExpiry DhcpOperData_Dhcpv4ClientOper_LeaseExpiry
}

func (dhcpv4ClientOper *DhcpOperData_Dhcpv4ClientOper) GetEntityData() *types.CommonEntityData {
    dhcpv4ClientOper.EntityData.YFilter = dhcpv4ClientOper.YFilter
    dhcpv4ClientOper.EntityData.YangName = "dhcpv4-client-oper"
    dhcpv4ClientOper.EntityData.BundleName = "cisco_ios_xe"
    dhcpv4ClientOper.EntityData.ParentYangName = "dhcp-oper-data"
    dhcpv4ClientOper.EntityData.SegmentPath = "dhcpv4-client-oper" + types.AddKeyToken(dhcpv4ClientOper.IfName, "if-name") + types.AddKeyToken(dhcpv4ClientOper.ClientAddr, "client-addr") + types.AddKeyToken(dhcpv4ClientOper.VrfName, "vrf-name")
    dhcpv4ClientOper.EntityData.AbsolutePath = "Cisco-IOS-XE-dhcp-oper:dhcp-oper-data/" + dhcpv4ClientOper.EntityData.SegmentPath
    dhcpv4ClientOper.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dhcpv4ClientOper.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dhcpv4ClientOper.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dhcpv4ClientOper.EntityData.Children = types.NewOrderedMap()
    dhcpv4ClientOper.EntityData.Children.Append("lease-expiry", types.YChild{"LeaseExpiry", &dhcpv4ClientOper.LeaseExpiry})
    dhcpv4ClientOper.EntityData.Leafs = types.NewOrderedMap()
    dhcpv4ClientOper.EntityData.Leafs.Append("if-name", types.YLeaf{"IfName", dhcpv4ClientOper.IfName})
    dhcpv4ClientOper.EntityData.Leafs.Append("client-addr", types.YLeaf{"ClientAddr", dhcpv4ClientOper.ClientAddr})
    dhcpv4ClientOper.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", dhcpv4ClientOper.VrfName})
    dhcpv4ClientOper.EntityData.Leafs.Append("state", types.YLeaf{"State", dhcpv4ClientOper.State})
    dhcpv4ClientOper.EntityData.Leafs.Append("lease-server-addr", types.YLeaf{"LeaseServerAddr", dhcpv4ClientOper.LeaseServerAddr})
    dhcpv4ClientOper.EntityData.Leafs.Append("gateway-addr", types.YLeaf{"GatewayAddr", dhcpv4ClientOper.GatewayAddr})
    dhcpv4ClientOper.EntityData.Leafs.Append("lease-time", types.YLeaf{"LeaseTime", dhcpv4ClientOper.LeaseTime})
    dhcpv4ClientOper.EntityData.Leafs.Append("lease-remaining", types.YLeaf{"LeaseRemaining", dhcpv4ClientOper.LeaseRemaining})
    dhcpv4ClientOper.EntityData.Leafs.Append("dns-address", types.YLeaf{"DnsAddress", dhcpv4ClientOper.DnsAddress})
    dhcpv4ClientOper.EntityData.Leafs.Append("dns-address-secondary", types.YLeaf{"DnsAddressSecondary", dhcpv4ClientOper.DnsAddressSecondary})
    dhcpv4ClientOper.EntityData.Leafs.Append("subnet-mask", types.YLeaf{"SubnetMask", dhcpv4ClientOper.SubnetMask})

    dhcpv4ClientOper.EntityData.YListKeys = []string {"IfName", "ClientAddr", "VrfName"}

    return &(dhcpv4ClientOper.EntityData)
}

// DhcpOperData_Dhcpv4ClientOper_LeaseExpiry
// Lease Expiry time for the IP address we got
type DhcpOperData_Dhcpv4ClientOper_LeaseExpiry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Date and time of expiry . The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    Time interface{}

    // Expiry time infinite. The type is interface{}.
    Infinite interface{}
}

func (leaseExpiry *DhcpOperData_Dhcpv4ClientOper_LeaseExpiry) GetEntityData() *types.CommonEntityData {
    leaseExpiry.EntityData.YFilter = leaseExpiry.YFilter
    leaseExpiry.EntityData.YangName = "lease-expiry"
    leaseExpiry.EntityData.BundleName = "cisco_ios_xe"
    leaseExpiry.EntityData.ParentYangName = "dhcpv4-client-oper"
    leaseExpiry.EntityData.SegmentPath = "lease-expiry"
    leaseExpiry.EntityData.AbsolutePath = "Cisco-IOS-XE-dhcp-oper:dhcp-oper-data/dhcpv4-client-oper/" + leaseExpiry.EntityData.SegmentPath
    leaseExpiry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    leaseExpiry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    leaseExpiry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    leaseExpiry.EntityData.Children = types.NewOrderedMap()
    leaseExpiry.EntityData.Leafs = types.NewOrderedMap()
    leaseExpiry.EntityData.Leafs.Append("time", types.YLeaf{"Time", leaseExpiry.Time})
    leaseExpiry.EntityData.Leafs.Append("infinite", types.YLeaf{"Infinite", leaseExpiry.Infinite})

    leaseExpiry.EntityData.YListKeys = []string {}

    return &(leaseExpiry.EntityData)
}

