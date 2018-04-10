// This module contains a collection of YANG definitions
// for Cisco IOS-XR ip-daps package operational data.
// 
// This module contains definitions
// for the following management objects:
//   address-pool-service: Pool operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ip_daps_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ip_daps_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ip-daps-oper address-pool-service}", reflect.TypeOf(AddressPoolService{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ip-daps-oper:address-pool-service", reflect.TypeOf(AddressPoolService{}))
}

// DapsClient represents DAPS client types
type DapsClient string

const (
    // Client type is None
    DapsClient_none DapsClient = "none"

    // Client type is PPP
    DapsClient_ppp DapsClient = "ppp"

    // Client type is DHCP
    DapsClient_dhcp DapsClient = "dhcp"

    // Client type is DHCPv6
    DapsClient_dhcpv6 DapsClient = "dhcpv6"

    // Client type is IPv6ND
    DapsClient_ipv6nd DapsClient = "ipv6nd"
)

// IpAddr represents Address type
type IpAddr string

const (
    // IPv4 address
    IpAddr_ipv4 IpAddr = "ipv4"

    // IPv6 address
    IpAddr_ipv6 IpAddr = "ipv6"
)

// AddressPoolService
// Pool operational data
type AddressPoolService struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Pool operational data for a particular location.
    Nodes AddressPoolService_Nodes
}

func (addressPoolService *AddressPoolService) GetEntityData() *types.CommonEntityData {
    addressPoolService.EntityData.YFilter = addressPoolService.YFilter
    addressPoolService.EntityData.YangName = "address-pool-service"
    addressPoolService.EntityData.BundleName = "cisco_ios_xr"
    addressPoolService.EntityData.ParentYangName = "Cisco-IOS-XR-ip-daps-oper"
    addressPoolService.EntityData.SegmentPath = "Cisco-IOS-XR-ip-daps-oper:address-pool-service"
    addressPoolService.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressPoolService.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressPoolService.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressPoolService.EntityData.Children = make(map[string]types.YChild)
    addressPoolService.EntityData.Children["nodes"] = types.YChild{"Nodes", &addressPoolService.Nodes}
    addressPoolService.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(addressPoolService.EntityData)
}

// AddressPoolService_Nodes
// Pool operational data for a particular location
type AddressPoolService_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Location. For eg., 0/1/CPU0. The type is slice of
    // AddressPoolService_Nodes_Node.
    Node []AddressPoolService_Nodes_Node
}

func (nodes *AddressPoolService_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "address-pool-service"
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

// AddressPoolService_Nodes_Node
// Location. For eg., 0/1/CPU0
type AddressPoolService_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node name. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeName interface{}

    // List of IPv4/IPv6 pool data.
    Pools AddressPoolService_Nodes_Node_Pools

    // Show total utilization for pool.
    TotalUtilization AddressPoolService_Nodes_Node_TotalUtilization

    // Pool VRF data.
    Vrfs AddressPoolService_Nodes_Node_Vrfs
}

func (node *AddressPoolService_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Children["pools"] = types.YChild{"Pools", &node.Pools}
    node.EntityData.Children["total-utilization"] = types.YChild{"TotalUtilization", &node.TotalUtilization}
    node.EntityData.Children["vrfs"] = types.YChild{"Vrfs", &node.Vrfs}
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", node.NodeName}
    return &(node.EntityData)
}

// AddressPoolService_Nodes_Node_Pools
// List of IPv4/IPv6 pool data
type AddressPoolService_Nodes_Node_Pools struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Pool data by pool name. The type is slice of
    // AddressPoolService_Nodes_Node_Pools_Pool.
    Pool []AddressPoolService_Nodes_Node_Pools_Pool
}

func (pools *AddressPoolService_Nodes_Node_Pools) GetEntityData() *types.CommonEntityData {
    pools.EntityData.YFilter = pools.YFilter
    pools.EntityData.YangName = "pools"
    pools.EntityData.BundleName = "cisco_ios_xr"
    pools.EntityData.ParentYangName = "node"
    pools.EntityData.SegmentPath = "pools"
    pools.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pools.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pools.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pools.EntityData.Children = make(map[string]types.YChild)
    pools.EntityData.Children["pool"] = types.YChild{"Pool", nil}
    for i := range pools.Pool {
        pools.EntityData.Children[types.GetSegmentPath(&pools.Pool[i])] = types.YChild{"Pool", &pools.Pool[i]}
    }
    pools.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(pools.EntityData)
}

// AddressPoolService_Nodes_Node_Pools_Pool
// Pool data by pool name
type AddressPoolService_Nodes_Node_Pools_Pool struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The pool name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    PoolName interface{}

    // Summary info for pool.
    AddressRanges AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges

    // Detailed info for the Pool.
    AllocatedAddresses AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses

    // Configuration info for pool.
    Configuration AddressPoolService_Nodes_Node_Pools_Pool_Configuration
}

func (pool *AddressPoolService_Nodes_Node_Pools_Pool) GetEntityData() *types.CommonEntityData {
    pool.EntityData.YFilter = pool.YFilter
    pool.EntityData.YangName = "pool"
    pool.EntityData.BundleName = "cisco_ios_xr"
    pool.EntityData.ParentYangName = "pools"
    pool.EntityData.SegmentPath = "pool" + "[pool-name='" + fmt.Sprintf("%v", pool.PoolName) + "']"
    pool.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pool.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pool.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pool.EntityData.Children = make(map[string]types.YChild)
    pool.EntityData.Children["address-ranges"] = types.YChild{"AddressRanges", &pool.AddressRanges}
    pool.EntityData.Children["allocated-addresses"] = types.YChild{"AllocatedAddresses", &pool.AllocatedAddresses}
    pool.EntityData.Children["configuration"] = types.YChild{"Configuration", &pool.Configuration}
    pool.EntityData.Leafs = make(map[string]types.YLeaf)
    pool.EntityData.Leafs["pool-name"] = types.YLeaf{"PoolName", pool.PoolName}
    return &(pool.EntityData)
}

// AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges
// Summary info for pool
type AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Start Address of the Range. The type is slice of
    // AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange.
    AddressRange []AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange
}

func (addressRanges *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges) GetEntityData() *types.CommonEntityData {
    addressRanges.EntityData.YFilter = addressRanges.YFilter
    addressRanges.EntityData.YangName = "address-ranges"
    addressRanges.EntityData.BundleName = "cisco_ios_xr"
    addressRanges.EntityData.ParentYangName = "pool"
    addressRanges.EntityData.SegmentPath = "address-ranges"
    addressRanges.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressRanges.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressRanges.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressRanges.EntityData.Children = make(map[string]types.YChild)
    addressRanges.EntityData.Children["address-range"] = types.YChild{"AddressRange", nil}
    for i := range addressRanges.AddressRange {
        addressRanges.EntityData.Children[types.GetSegmentPath(&addressRanges.AddressRange[i])] = types.YChild{"AddressRange", &addressRanges.AddressRange[i]}
    }
    addressRanges.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(addressRanges.EntityData)
}

// AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange
// Start Address of the Range
type AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. IP Address. The type is one of the following
    // types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    StartAddress interface{}

    // Pool name. The type is string with length: 0..64.
    PoolName interface{}

    // VRF name. The type is string with length: 0..64.
    VrfName interface{}

    // Pool scope. The type is string with length: 0..64.
    PoolScope interface{}

    // Number of addresses allocated. The type is interface{} with range:
    // 0..4294967295.
    AllocatedAddresses interface{}

    // Number of addresses free. The type is interface{} with range:
    // 0..4294967295.
    FreeAddresses interface{}

    // Number of addresses excluded. The type is interface{} with range:
    // 0..4294967295.
    ExcludedAddresses interface{}

    // Is network blocked. The type is interface{} with range: 0..4294967295.
    NetworkBlockedStatus interface{}

    // Is network blocked trap send. The type is interface{} with range:
    // 0..4294967295.
    NetworkBlockedStatusTrp interface{}

    // Range start.
    StartAddressXr AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_StartAddressXr

    // Range end.
    EndAddress AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_EndAddress

    // Default router.
    DefaultRouter AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_DefaultRouter
}

func (addressRange *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange) GetEntityData() *types.CommonEntityData {
    addressRange.EntityData.YFilter = addressRange.YFilter
    addressRange.EntityData.YangName = "address-range"
    addressRange.EntityData.BundleName = "cisco_ios_xr"
    addressRange.EntityData.ParentYangName = "address-ranges"
    addressRange.EntityData.SegmentPath = "address-range" + "[start-address='" + fmt.Sprintf("%v", addressRange.StartAddress) + "']"
    addressRange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressRange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressRange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressRange.EntityData.Children = make(map[string]types.YChild)
    addressRange.EntityData.Children["start-address-xr"] = types.YChild{"StartAddressXr", &addressRange.StartAddressXr}
    addressRange.EntityData.Children["end-address"] = types.YChild{"EndAddress", &addressRange.EndAddress}
    addressRange.EntityData.Children["default-router"] = types.YChild{"DefaultRouter", &addressRange.DefaultRouter}
    addressRange.EntityData.Leafs = make(map[string]types.YLeaf)
    addressRange.EntityData.Leafs["start-address"] = types.YLeaf{"StartAddress", addressRange.StartAddress}
    addressRange.EntityData.Leafs["pool-name"] = types.YLeaf{"PoolName", addressRange.PoolName}
    addressRange.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", addressRange.VrfName}
    addressRange.EntityData.Leafs["pool-scope"] = types.YLeaf{"PoolScope", addressRange.PoolScope}
    addressRange.EntityData.Leafs["allocated-addresses"] = types.YLeaf{"AllocatedAddresses", addressRange.AllocatedAddresses}
    addressRange.EntityData.Leafs["free-addresses"] = types.YLeaf{"FreeAddresses", addressRange.FreeAddresses}
    addressRange.EntityData.Leafs["excluded-addresses"] = types.YLeaf{"ExcludedAddresses", addressRange.ExcludedAddresses}
    addressRange.EntityData.Leafs["network-blocked-status"] = types.YLeaf{"NetworkBlockedStatus", addressRange.NetworkBlockedStatus}
    addressRange.EntityData.Leafs["network-blocked-status-trp"] = types.YLeaf{"NetworkBlockedStatusTrp", addressRange.NetworkBlockedStatusTrp}
    return &(addressRange.EntityData)
}

// AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_StartAddressXr
// Range start
type AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_StartAddressXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address.
    Address AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_StartAddressXr_Address
}

func (startAddressXr *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_StartAddressXr) GetEntityData() *types.CommonEntityData {
    startAddressXr.EntityData.YFilter = startAddressXr.YFilter
    startAddressXr.EntityData.YangName = "start-address-xr"
    startAddressXr.EntityData.BundleName = "cisco_ios_xr"
    startAddressXr.EntityData.ParentYangName = "address-range"
    startAddressXr.EntityData.SegmentPath = "start-address-xr"
    startAddressXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    startAddressXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    startAddressXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    startAddressXr.EntityData.Children = make(map[string]types.YChild)
    startAddressXr.EntityData.Children["address"] = types.YChild{"Address", &startAddressXr.Address}
    startAddressXr.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(startAddressXr.EntityData)
}

// AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_StartAddressXr_Address
// Address
type AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_StartAddressXr_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AddressFamily. The type is IpAddr.
    AddressFamily interface{}

    // IPv4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4Address interface{}

    // IPv6 address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6Address interface{}
}

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_StartAddressXr_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "start-address-xr"
    address.EntityData.SegmentPath = "address"
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = make(map[string]types.YChild)
    address.EntityData.Leafs = make(map[string]types.YLeaf)
    address.EntityData.Leafs["address-family"] = types.YLeaf{"AddressFamily", address.AddressFamily}
    address.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", address.Ipv4Address}
    address.EntityData.Leafs["ipv6-address"] = types.YLeaf{"Ipv6Address", address.Ipv6Address}
    return &(address.EntityData)
}

// AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_EndAddress
// Range end
type AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_EndAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address.
    Address AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_EndAddress_Address
}

func (endAddress *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_EndAddress) GetEntityData() *types.CommonEntityData {
    endAddress.EntityData.YFilter = endAddress.YFilter
    endAddress.EntityData.YangName = "end-address"
    endAddress.EntityData.BundleName = "cisco_ios_xr"
    endAddress.EntityData.ParentYangName = "address-range"
    endAddress.EntityData.SegmentPath = "end-address"
    endAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    endAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    endAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    endAddress.EntityData.Children = make(map[string]types.YChild)
    endAddress.EntityData.Children["address"] = types.YChild{"Address", &endAddress.Address}
    endAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(endAddress.EntityData)
}

// AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_EndAddress_Address
// Address
type AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_EndAddress_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AddressFamily. The type is IpAddr.
    AddressFamily interface{}

    // IPv4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4Address interface{}

    // IPv6 address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6Address interface{}
}

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_EndAddress_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "end-address"
    address.EntityData.SegmentPath = "address"
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = make(map[string]types.YChild)
    address.EntityData.Leafs = make(map[string]types.YLeaf)
    address.EntityData.Leafs["address-family"] = types.YLeaf{"AddressFamily", address.AddressFamily}
    address.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", address.Ipv4Address}
    address.EntityData.Leafs["ipv6-address"] = types.YLeaf{"Ipv6Address", address.Ipv6Address}
    return &(address.EntityData)
}

// AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_DefaultRouter
// Default router
type AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_DefaultRouter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address.
    Address AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_DefaultRouter_Address
}

func (defaultRouter *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_DefaultRouter) GetEntityData() *types.CommonEntityData {
    defaultRouter.EntityData.YFilter = defaultRouter.YFilter
    defaultRouter.EntityData.YangName = "default-router"
    defaultRouter.EntityData.BundleName = "cisco_ios_xr"
    defaultRouter.EntityData.ParentYangName = "address-range"
    defaultRouter.EntityData.SegmentPath = "default-router"
    defaultRouter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defaultRouter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defaultRouter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defaultRouter.EntityData.Children = make(map[string]types.YChild)
    defaultRouter.EntityData.Children["address"] = types.YChild{"Address", &defaultRouter.Address}
    defaultRouter.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(defaultRouter.EntityData)
}

// AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_DefaultRouter_Address
// Address
type AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_DefaultRouter_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AddressFamily. The type is IpAddr.
    AddressFamily interface{}

    // IPv4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4Address interface{}

    // IPv6 address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6Address interface{}
}

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_DefaultRouter_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "default-router"
    address.EntityData.SegmentPath = "address"
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = make(map[string]types.YChild)
    address.EntityData.Leafs = make(map[string]types.YLeaf)
    address.EntityData.Leafs["address-family"] = types.YLeaf{"AddressFamily", address.AddressFamily}
    address.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", address.Ipv4Address}
    address.EntityData.Leafs["ipv6-address"] = types.YLeaf{"Ipv6Address", address.Ipv6Address}
    return &(address.EntityData)
}

// AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses
// Detailed info for the Pool
type AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Pool allocations.
    PoolAllocations AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_PoolAllocations

    // Address ranges. The type is slice of
    // AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange.
    AddressRange []AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange

    // In-use addresses. The type is slice of
    // AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress.
    InUseAddress []AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress
}

func (allocatedAddresses *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses) GetEntityData() *types.CommonEntityData {
    allocatedAddresses.EntityData.YFilter = allocatedAddresses.YFilter
    allocatedAddresses.EntityData.YangName = "allocated-addresses"
    allocatedAddresses.EntityData.BundleName = "cisco_ios_xr"
    allocatedAddresses.EntityData.ParentYangName = "pool"
    allocatedAddresses.EntityData.SegmentPath = "allocated-addresses"
    allocatedAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allocatedAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allocatedAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allocatedAddresses.EntityData.Children = make(map[string]types.YChild)
    allocatedAddresses.EntityData.Children["pool-allocations"] = types.YChild{"PoolAllocations", &allocatedAddresses.PoolAllocations}
    allocatedAddresses.EntityData.Children["address-range"] = types.YChild{"AddressRange", nil}
    for i := range allocatedAddresses.AddressRange {
        allocatedAddresses.EntityData.Children[types.GetSegmentPath(&allocatedAddresses.AddressRange[i])] = types.YChild{"AddressRange", &allocatedAddresses.AddressRange[i]}
    }
    allocatedAddresses.EntityData.Children["in-use-address"] = types.YChild{"InUseAddress", nil}
    for i := range allocatedAddresses.InUseAddress {
        allocatedAddresses.EntityData.Children[types.GetSegmentPath(&allocatedAddresses.InUseAddress[i])] = types.YChild{"InUseAddress", &allocatedAddresses.InUseAddress[i]}
    }
    allocatedAddresses.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(allocatedAddresses.EntityData)
}

// AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_PoolAllocations
// Pool allocations
type AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_PoolAllocations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF name. The type is string with length: 0..64.
    VrfName interface{}

    // Used allocations. The type is interface{} with range: 0..4294967295.
    Used interface{}

    // Excluded allocations. The type is interface{} with range: 0..4294967295.
    Excluded interface{}

    // Free allocations. The type is interface{} with range: 0..4294967295.
    Free interface{}

    // Total allocations. The type is interface{} with range: 0..4294967295.
    Total interface{}

    // Current utilization in percentage. The type is interface{} with range:
    // 0..4294967295. Units are percentage.
    Utilization interface{}

    // High threshold data.
    HighThreshold AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_PoolAllocations_HighThreshold

    // Low threshold data.
    LowThreshold AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_PoolAllocations_LowThreshold
}

func (poolAllocations *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_PoolAllocations) GetEntityData() *types.CommonEntityData {
    poolAllocations.EntityData.YFilter = poolAllocations.YFilter
    poolAllocations.EntityData.YangName = "pool-allocations"
    poolAllocations.EntityData.BundleName = "cisco_ios_xr"
    poolAllocations.EntityData.ParentYangName = "allocated-addresses"
    poolAllocations.EntityData.SegmentPath = "pool-allocations"
    poolAllocations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    poolAllocations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    poolAllocations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    poolAllocations.EntityData.Children = make(map[string]types.YChild)
    poolAllocations.EntityData.Children["high-threshold"] = types.YChild{"HighThreshold", &poolAllocations.HighThreshold}
    poolAllocations.EntityData.Children["low-threshold"] = types.YChild{"LowThreshold", &poolAllocations.LowThreshold}
    poolAllocations.EntityData.Leafs = make(map[string]types.YLeaf)
    poolAllocations.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", poolAllocations.VrfName}
    poolAllocations.EntityData.Leafs["used"] = types.YLeaf{"Used", poolAllocations.Used}
    poolAllocations.EntityData.Leafs["excluded"] = types.YLeaf{"Excluded", poolAllocations.Excluded}
    poolAllocations.EntityData.Leafs["free"] = types.YLeaf{"Free", poolAllocations.Free}
    poolAllocations.EntityData.Leafs["total"] = types.YLeaf{"Total", poolAllocations.Total}
    poolAllocations.EntityData.Leafs["utilization"] = types.YLeaf{"Utilization", poolAllocations.Utilization}
    return &(poolAllocations.EntityData)
}

// AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_PoolAllocations_HighThreshold
// High threshold data
type AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_PoolAllocations_HighThreshold struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Threshold in percentage. The type is interface{} with range: 0..4294967295.
    // Units are percentage.
    Threshold interface{}

    // Number of Triggers. The type is interface{} with range: 0..4294967295.
    Triggers interface{}

    // Last time at which threshold crossed in DDD MMM DD HH:MM:SS YYYY format eg:
    // Tue Apr 11 21:30:47 2011. The type is string.
    TimeLastCrossed interface{}
}

func (highThreshold *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_PoolAllocations_HighThreshold) GetEntityData() *types.CommonEntityData {
    highThreshold.EntityData.YFilter = highThreshold.YFilter
    highThreshold.EntityData.YangName = "high-threshold"
    highThreshold.EntityData.BundleName = "cisco_ios_xr"
    highThreshold.EntityData.ParentYangName = "pool-allocations"
    highThreshold.EntityData.SegmentPath = "high-threshold"
    highThreshold.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    highThreshold.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    highThreshold.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    highThreshold.EntityData.Children = make(map[string]types.YChild)
    highThreshold.EntityData.Leafs = make(map[string]types.YLeaf)
    highThreshold.EntityData.Leafs["threshold"] = types.YLeaf{"Threshold", highThreshold.Threshold}
    highThreshold.EntityData.Leafs["triggers"] = types.YLeaf{"Triggers", highThreshold.Triggers}
    highThreshold.EntityData.Leafs["time-last-crossed"] = types.YLeaf{"TimeLastCrossed", highThreshold.TimeLastCrossed}
    return &(highThreshold.EntityData)
}

// AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_PoolAllocations_LowThreshold
// Low threshold data
type AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_PoolAllocations_LowThreshold struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Threshold in percentage. The type is interface{} with range: 0..4294967295.
    // Units are percentage.
    Threshold interface{}

    // Number of Triggers. The type is interface{} with range: 0..4294967295.
    Triggers interface{}

    // Last time at which threshold crossed in DDD MMM DD HH:MM:SS YYYY format eg:
    // Tue Apr 11 21:30:47 2011. The type is string.
    TimeLastCrossed interface{}
}

func (lowThreshold *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_PoolAllocations_LowThreshold) GetEntityData() *types.CommonEntityData {
    lowThreshold.EntityData.YFilter = lowThreshold.YFilter
    lowThreshold.EntityData.YangName = "low-threshold"
    lowThreshold.EntityData.BundleName = "cisco_ios_xr"
    lowThreshold.EntityData.ParentYangName = "pool-allocations"
    lowThreshold.EntityData.SegmentPath = "low-threshold"
    lowThreshold.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lowThreshold.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lowThreshold.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lowThreshold.EntityData.Children = make(map[string]types.YChild)
    lowThreshold.EntityData.Leafs = make(map[string]types.YLeaf)
    lowThreshold.EntityData.Leafs["threshold"] = types.YLeaf{"Threshold", lowThreshold.Threshold}
    lowThreshold.EntityData.Leafs["triggers"] = types.YLeaf{"Triggers", lowThreshold.Triggers}
    lowThreshold.EntityData.Leafs["time-last-crossed"] = types.YLeaf{"TimeLastCrossed", lowThreshold.TimeLastCrossed}
    return &(lowThreshold.EntityData)
}

// AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange
// Address ranges
type AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Used allocations. The type is interface{} with range: 0..4294967295.
    Used interface{}

    // Excluded allocations. The type is interface{} with range: 0..4294967295.
    Excluded interface{}

    // Free allocations. The type is interface{} with range: 0..4294967295.
    Free interface{}

    // Range start.
    StartAddress AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_StartAddress

    // Range end.
    EndAddress AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_EndAddress
}

func (addressRange *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange) GetEntityData() *types.CommonEntityData {
    addressRange.EntityData.YFilter = addressRange.YFilter
    addressRange.EntityData.YangName = "address-range"
    addressRange.EntityData.BundleName = "cisco_ios_xr"
    addressRange.EntityData.ParentYangName = "allocated-addresses"
    addressRange.EntityData.SegmentPath = "address-range"
    addressRange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressRange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressRange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressRange.EntityData.Children = make(map[string]types.YChild)
    addressRange.EntityData.Children["start-address"] = types.YChild{"StartAddress", &addressRange.StartAddress}
    addressRange.EntityData.Children["end-address"] = types.YChild{"EndAddress", &addressRange.EndAddress}
    addressRange.EntityData.Leafs = make(map[string]types.YLeaf)
    addressRange.EntityData.Leafs["used"] = types.YLeaf{"Used", addressRange.Used}
    addressRange.EntityData.Leafs["excluded"] = types.YLeaf{"Excluded", addressRange.Excluded}
    addressRange.EntityData.Leafs["free"] = types.YLeaf{"Free", addressRange.Free}
    return &(addressRange.EntityData)
}

// AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_StartAddress
// Range start
type AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_StartAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address.
    Address AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_StartAddress_Address
}

func (startAddress *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_StartAddress) GetEntityData() *types.CommonEntityData {
    startAddress.EntityData.YFilter = startAddress.YFilter
    startAddress.EntityData.YangName = "start-address"
    startAddress.EntityData.BundleName = "cisco_ios_xr"
    startAddress.EntityData.ParentYangName = "address-range"
    startAddress.EntityData.SegmentPath = "start-address"
    startAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    startAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    startAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    startAddress.EntityData.Children = make(map[string]types.YChild)
    startAddress.EntityData.Children["address"] = types.YChild{"Address", &startAddress.Address}
    startAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(startAddress.EntityData)
}

// AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_StartAddress_Address
// Address
type AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_StartAddress_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AddressFamily. The type is IpAddr.
    AddressFamily interface{}

    // IPv4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4Address interface{}

    // IPv6 address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6Address interface{}
}

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_StartAddress_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "start-address"
    address.EntityData.SegmentPath = "address"
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = make(map[string]types.YChild)
    address.EntityData.Leafs = make(map[string]types.YLeaf)
    address.EntityData.Leafs["address-family"] = types.YLeaf{"AddressFamily", address.AddressFamily}
    address.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", address.Ipv4Address}
    address.EntityData.Leafs["ipv6-address"] = types.YLeaf{"Ipv6Address", address.Ipv6Address}
    return &(address.EntityData)
}

// AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_EndAddress
// Range end
type AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_EndAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address.
    Address AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_EndAddress_Address
}

func (endAddress *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_EndAddress) GetEntityData() *types.CommonEntityData {
    endAddress.EntityData.YFilter = endAddress.YFilter
    endAddress.EntityData.YangName = "end-address"
    endAddress.EntityData.BundleName = "cisco_ios_xr"
    endAddress.EntityData.ParentYangName = "address-range"
    endAddress.EntityData.SegmentPath = "end-address"
    endAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    endAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    endAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    endAddress.EntityData.Children = make(map[string]types.YChild)
    endAddress.EntityData.Children["address"] = types.YChild{"Address", &endAddress.Address}
    endAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(endAddress.EntityData)
}

// AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_EndAddress_Address
// Address
type AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_EndAddress_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AddressFamily. The type is IpAddr.
    AddressFamily interface{}

    // IPv4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4Address interface{}

    // IPv6 address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6Address interface{}
}

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_EndAddress_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "end-address"
    address.EntityData.SegmentPath = "address"
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = make(map[string]types.YChild)
    address.EntityData.Leafs = make(map[string]types.YLeaf)
    address.EntityData.Leafs["address-family"] = types.YLeaf{"AddressFamily", address.AddressFamily}
    address.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", address.Ipv4Address}
    address.EntityData.Leafs["ipv6-address"] = types.YLeaf{"Ipv6Address", address.Ipv6Address}
    return &(address.EntityData)
}

// AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress
// In-use addresses
type AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Client type. The type is DapsClient.
    ClientType interface{}

    // Client address.
    Address AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress_Address
}

func (inUseAddress *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress) GetEntityData() *types.CommonEntityData {
    inUseAddress.EntityData.YFilter = inUseAddress.YFilter
    inUseAddress.EntityData.YangName = "in-use-address"
    inUseAddress.EntityData.BundleName = "cisco_ios_xr"
    inUseAddress.EntityData.ParentYangName = "allocated-addresses"
    inUseAddress.EntityData.SegmentPath = "in-use-address"
    inUseAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inUseAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inUseAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inUseAddress.EntityData.Children = make(map[string]types.YChild)
    inUseAddress.EntityData.Children["address"] = types.YChild{"Address", &inUseAddress.Address}
    inUseAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    inUseAddress.EntityData.Leafs["client-type"] = types.YLeaf{"ClientType", inUseAddress.ClientType}
    return &(inUseAddress.EntityData)
}

// AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress_Address
// Client address
type AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address.
    Address AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress_Address_Address_
}

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "in-use-address"
    address.EntityData.SegmentPath = "address"
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = make(map[string]types.YChild)
    address.EntityData.Children["address"] = types.YChild{"Address", &address.Address}
    address.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(address.EntityData)
}

// AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress_Address_Address_
// Address
type AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress_Address_Address_ struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AddressFamily. The type is IpAddr.
    AddressFamily interface{}

    // IPv4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4Address interface{}

    // IPv6 address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6Address interface{}
}

func (address_ *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress_Address_Address_) GetEntityData() *types.CommonEntityData {
    address_.EntityData.YFilter = address_.YFilter
    address_.EntityData.YangName = "address"
    address_.EntityData.BundleName = "cisco_ios_xr"
    address_.EntityData.ParentYangName = "address"
    address_.EntityData.SegmentPath = "address"
    address_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address_.EntityData.Children = make(map[string]types.YChild)
    address_.EntityData.Leafs = make(map[string]types.YLeaf)
    address_.EntityData.Leafs["address-family"] = types.YLeaf{"AddressFamily", address_.AddressFamily}
    address_.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", address_.Ipv4Address}
    address_.EntityData.Leafs["ipv6-address"] = types.YLeaf{"Ipv6Address", address_.Ipv6Address}
    return &(address_.EntityData)
}

// AddressPoolService_Nodes_Node_Pools_Pool_Configuration
// Configuration info for pool
type AddressPoolService_Nodes_Node_Pools_Pool_Configuration struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Pool name. The type is string with length: 0..64.
    PoolName interface{}

    // Pool ID for MIBS. The type is interface{} with range: 0..4294967295.
    PoolId interface{}

    // VRF name. The type is string with length: 0..64.
    VrfName interface{}

    // Pool Scope. The type is string with length: 0..64.
    PoolScope interface{}

    // Prefix length. The type is interface{} with range: 0..4294967295.
    PoolPrefixLength interface{}

    // High utilization mark. The type is interface{} with range: 0..255.
    HighUtilizationMark interface{}

    // Low utilization mark. The type is interface{} with range: 0..255.
    LowUtilizationMark interface{}

    // Current utilization. The type is interface{} with range: 0..255.
    CurrentUtilization interface{}

    // Number of times High utilization threshold was crossed. The type is
    // interface{} with range: 0..4294967295.
    UtilizationHighCount interface{}

    // Number of times Low utilization threshold was crossed. The type is
    // interface{} with range: 0..4294967295.
    UtilizationLowCount interface{}
}

func (configuration *AddressPoolService_Nodes_Node_Pools_Pool_Configuration) GetEntityData() *types.CommonEntityData {
    configuration.EntityData.YFilter = configuration.YFilter
    configuration.EntityData.YangName = "configuration"
    configuration.EntityData.BundleName = "cisco_ios_xr"
    configuration.EntityData.ParentYangName = "pool"
    configuration.EntityData.SegmentPath = "configuration"
    configuration.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configuration.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configuration.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configuration.EntityData.Children = make(map[string]types.YChild)
    configuration.EntityData.Leafs = make(map[string]types.YLeaf)
    configuration.EntityData.Leafs["pool-name"] = types.YLeaf{"PoolName", configuration.PoolName}
    configuration.EntityData.Leafs["pool-id"] = types.YLeaf{"PoolId", configuration.PoolId}
    configuration.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", configuration.VrfName}
    configuration.EntityData.Leafs["pool-scope"] = types.YLeaf{"PoolScope", configuration.PoolScope}
    configuration.EntityData.Leafs["pool-prefix-length"] = types.YLeaf{"PoolPrefixLength", configuration.PoolPrefixLength}
    configuration.EntityData.Leafs["high-utilization-mark"] = types.YLeaf{"HighUtilizationMark", configuration.HighUtilizationMark}
    configuration.EntityData.Leafs["low-utilization-mark"] = types.YLeaf{"LowUtilizationMark", configuration.LowUtilizationMark}
    configuration.EntityData.Leafs["current-utilization"] = types.YLeaf{"CurrentUtilization", configuration.CurrentUtilization}
    configuration.EntityData.Leafs["utilization-high-count"] = types.YLeaf{"UtilizationHighCount", configuration.UtilizationHighCount}
    configuration.EntityData.Leafs["utilization-low-count"] = types.YLeaf{"UtilizationLowCount", configuration.UtilizationLowCount}
    return &(configuration.EntityData)
}

// AddressPoolService_Nodes_Node_TotalUtilization
// Show total utilization for pool
type AddressPoolService_Nodes_Node_TotalUtilization struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // High utilization mark. The type is interface{} with range: 0..255.
    TotalUtilizationHighMark interface{}

    // Low utilization mark. The type is interface{} with range: 0..255.
    TotalUtilizationLowMark interface{}

    // Current utilization. The type is interface{} with range: 0..255.
    CurrentTotalUtilization interface{}
}

func (totalUtilization *AddressPoolService_Nodes_Node_TotalUtilization) GetEntityData() *types.CommonEntityData {
    totalUtilization.EntityData.YFilter = totalUtilization.YFilter
    totalUtilization.EntityData.YangName = "total-utilization"
    totalUtilization.EntityData.BundleName = "cisco_ios_xr"
    totalUtilization.EntityData.ParentYangName = "node"
    totalUtilization.EntityData.SegmentPath = "total-utilization"
    totalUtilization.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    totalUtilization.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    totalUtilization.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    totalUtilization.EntityData.Children = make(map[string]types.YChild)
    totalUtilization.EntityData.Leafs = make(map[string]types.YLeaf)
    totalUtilization.EntityData.Leafs["total-utilization-high-mark"] = types.YLeaf{"TotalUtilizationHighMark", totalUtilization.TotalUtilizationHighMark}
    totalUtilization.EntityData.Leafs["total-utilization-low-mark"] = types.YLeaf{"TotalUtilizationLowMark", totalUtilization.TotalUtilizationLowMark}
    totalUtilization.EntityData.Leafs["current-total-utilization"] = types.YLeaf{"CurrentTotalUtilization", totalUtilization.CurrentTotalUtilization}
    return &(totalUtilization.EntityData)
}

// AddressPoolService_Nodes_Node_Vrfs
// Pool VRF data
type AddressPoolService_Nodes_Node_Vrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF level Pool information. The type is slice of
    // AddressPoolService_Nodes_Node_Vrfs_Vrf.
    Vrf []AddressPoolService_Nodes_Node_Vrfs_Vrf
}

func (vrfs *AddressPoolService_Nodes_Node_Vrfs) GetEntityData() *types.CommonEntityData {
    vrfs.EntityData.YFilter = vrfs.YFilter
    vrfs.EntityData.YangName = "vrfs"
    vrfs.EntityData.BundleName = "cisco_ios_xr"
    vrfs.EntityData.ParentYangName = "node"
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

// AddressPoolService_Nodes_Node_Vrfs_Vrf
// VRF level Pool information
type AddressPoolService_Nodes_Node_Vrfs_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The VRF name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    VrfName interface{}

    // IPv4 pool VRF data.
    Ipv4 AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4

    // IPv6 Pool VRF data.
    Ipv6 AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6
}

func (vrf *AddressPoolService_Nodes_Node_Vrfs_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "vrfs"
    vrf.EntityData.SegmentPath = "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = make(map[string]types.YChild)
    vrf.EntityData.Children["ipv4"] = types.YChild{"Ipv4", &vrf.Ipv4}
    vrf.EntityData.Children["ipv6"] = types.YChild{"Ipv6", &vrf.Ipv6}
    vrf.EntityData.Leafs = make(map[string]types.YLeaf)
    vrf.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", vrf.VrfName}
    return &(vrf.EntityData)
}

// AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4
// IPv4 pool VRF data
type AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Allocation summary.
    AllocationSummary AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4_AllocationSummary

    // Pools data. The type is slice of
    // AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4_Pools.
    Pools []AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4_Pools
}

func (ipv4 *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4) GetEntityData() *types.CommonEntityData {
    ipv4.EntityData.YFilter = ipv4.YFilter
    ipv4.EntityData.YangName = "ipv4"
    ipv4.EntityData.BundleName = "cisco_ios_xr"
    ipv4.EntityData.ParentYangName = "vrf"
    ipv4.EntityData.SegmentPath = "ipv4"
    ipv4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4.EntityData.Children = make(map[string]types.YChild)
    ipv4.EntityData.Children["allocation-summary"] = types.YChild{"AllocationSummary", &ipv4.AllocationSummary}
    ipv4.EntityData.Children["pools"] = types.YChild{"Pools", nil}
    for i := range ipv4.Pools {
        ipv4.EntityData.Children[types.GetSegmentPath(&ipv4.Pools[i])] = types.YChild{"Pools", &ipv4.Pools[i]}
    }
    ipv4.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipv4.EntityData)
}

// AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4_AllocationSummary
// Allocation summary
type AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4_AllocationSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Used allocations. The type is interface{} with range: 0..4294967295.
    Used interface{}

    // Excluded allocations. The type is interface{} with range: 0..4294967295.
    Excluded interface{}

    // Free allocations. The type is interface{} with range: 0..4294967295.
    Free interface{}

    // Total allocations. The type is interface{} with range: 0..4294967295.
    Total interface{}

    // High utilization threshold in percentage. The type is interface{} with
    // range: 0..255. Units are percentage.
    HighUtilizationThreshold interface{}

    // Low utilization threshold in percentage. The type is interface{} with
    // range: 0..255. Units are percentage.
    LowUtilizationThreshold interface{}

    // Current utilization in percentage. The type is interface{} with range:
    // 0..255. Units are percentage.
    Utilization interface{}
}

func (allocationSummary *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4_AllocationSummary) GetEntityData() *types.CommonEntityData {
    allocationSummary.EntityData.YFilter = allocationSummary.YFilter
    allocationSummary.EntityData.YangName = "allocation-summary"
    allocationSummary.EntityData.BundleName = "cisco_ios_xr"
    allocationSummary.EntityData.ParentYangName = "ipv4"
    allocationSummary.EntityData.SegmentPath = "allocation-summary"
    allocationSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allocationSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allocationSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allocationSummary.EntityData.Children = make(map[string]types.YChild)
    allocationSummary.EntityData.Leafs = make(map[string]types.YLeaf)
    allocationSummary.EntityData.Leafs["used"] = types.YLeaf{"Used", allocationSummary.Used}
    allocationSummary.EntityData.Leafs["excluded"] = types.YLeaf{"Excluded", allocationSummary.Excluded}
    allocationSummary.EntityData.Leafs["free"] = types.YLeaf{"Free", allocationSummary.Free}
    allocationSummary.EntityData.Leafs["total"] = types.YLeaf{"Total", allocationSummary.Total}
    allocationSummary.EntityData.Leafs["high-utilization-threshold"] = types.YLeaf{"HighUtilizationThreshold", allocationSummary.HighUtilizationThreshold}
    allocationSummary.EntityData.Leafs["low-utilization-threshold"] = types.YLeaf{"LowUtilizationThreshold", allocationSummary.LowUtilizationThreshold}
    allocationSummary.EntityData.Leafs["utilization"] = types.YLeaf{"Utilization", allocationSummary.Utilization}
    return &(allocationSummary.EntityData)
}

// AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4_Pools
// Pools data
type AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4_Pools struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Pool name. The type is string with length: 0..64.
    PoolName interface{}

    // VRF name. The type is string with length: 0..64.
    VrfName interface{}

    // Used allocations. The type is interface{} with range: 0..4294967295.
    Used interface{}

    // Excluded allocations. The type is interface{} with range: 0..4294967295.
    Excluded interface{}

    // Free allocations. The type is interface{} with range: 0..4294967295.
    Free interface{}

    // Total allocations. The type is interface{} with range: 0..4294967295.
    Total interface{}
}

func (pools *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4_Pools) GetEntityData() *types.CommonEntityData {
    pools.EntityData.YFilter = pools.YFilter
    pools.EntityData.YangName = "pools"
    pools.EntityData.BundleName = "cisco_ios_xr"
    pools.EntityData.ParentYangName = "ipv4"
    pools.EntityData.SegmentPath = "pools"
    pools.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pools.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pools.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pools.EntityData.Children = make(map[string]types.YChild)
    pools.EntityData.Leafs = make(map[string]types.YLeaf)
    pools.EntityData.Leafs["pool-name"] = types.YLeaf{"PoolName", pools.PoolName}
    pools.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", pools.VrfName}
    pools.EntityData.Leafs["used"] = types.YLeaf{"Used", pools.Used}
    pools.EntityData.Leafs["excluded"] = types.YLeaf{"Excluded", pools.Excluded}
    pools.EntityData.Leafs["free"] = types.YLeaf{"Free", pools.Free}
    pools.EntityData.Leafs["total"] = types.YLeaf{"Total", pools.Total}
    return &(pools.EntityData)
}

// AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6
// IPv6 Pool VRF data
type AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Allocation summary.
    AllocationSummary AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6_AllocationSummary

    // Pools data. The type is slice of
    // AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6_Pools.
    Pools []AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6_Pools
}

func (ipv6 *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6) GetEntityData() *types.CommonEntityData {
    ipv6.EntityData.YFilter = ipv6.YFilter
    ipv6.EntityData.YangName = "ipv6"
    ipv6.EntityData.BundleName = "cisco_ios_xr"
    ipv6.EntityData.ParentYangName = "vrf"
    ipv6.EntityData.SegmentPath = "ipv6"
    ipv6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6.EntityData.Children = make(map[string]types.YChild)
    ipv6.EntityData.Children["allocation-summary"] = types.YChild{"AllocationSummary", &ipv6.AllocationSummary}
    ipv6.EntityData.Children["pools"] = types.YChild{"Pools", nil}
    for i := range ipv6.Pools {
        ipv6.EntityData.Children[types.GetSegmentPath(&ipv6.Pools[i])] = types.YChild{"Pools", &ipv6.Pools[i]}
    }
    ipv6.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipv6.EntityData)
}

// AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6_AllocationSummary
// Allocation summary
type AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6_AllocationSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Used allocations. The type is interface{} with range: 0..4294967295.
    Used interface{}

    // Excluded allocations. The type is interface{} with range: 0..4294967295.
    Excluded interface{}

    // Free allocations. The type is interface{} with range: 0..4294967295.
    Free interface{}

    // Total allocations. The type is interface{} with range: 0..4294967295.
    Total interface{}

    // High utilization threshold in percentage. The type is interface{} with
    // range: 0..255. Units are percentage.
    HighUtilizationThreshold interface{}

    // Low utilization threshold in percentage. The type is interface{} with
    // range: 0..255. Units are percentage.
    LowUtilizationThreshold interface{}

    // Current utilization in percentage. The type is interface{} with range:
    // 0..255. Units are percentage.
    Utilization interface{}
}

func (allocationSummary *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6_AllocationSummary) GetEntityData() *types.CommonEntityData {
    allocationSummary.EntityData.YFilter = allocationSummary.YFilter
    allocationSummary.EntityData.YangName = "allocation-summary"
    allocationSummary.EntityData.BundleName = "cisco_ios_xr"
    allocationSummary.EntityData.ParentYangName = "ipv6"
    allocationSummary.EntityData.SegmentPath = "allocation-summary"
    allocationSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allocationSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allocationSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allocationSummary.EntityData.Children = make(map[string]types.YChild)
    allocationSummary.EntityData.Leafs = make(map[string]types.YLeaf)
    allocationSummary.EntityData.Leafs["used"] = types.YLeaf{"Used", allocationSummary.Used}
    allocationSummary.EntityData.Leafs["excluded"] = types.YLeaf{"Excluded", allocationSummary.Excluded}
    allocationSummary.EntityData.Leafs["free"] = types.YLeaf{"Free", allocationSummary.Free}
    allocationSummary.EntityData.Leafs["total"] = types.YLeaf{"Total", allocationSummary.Total}
    allocationSummary.EntityData.Leafs["high-utilization-threshold"] = types.YLeaf{"HighUtilizationThreshold", allocationSummary.HighUtilizationThreshold}
    allocationSummary.EntityData.Leafs["low-utilization-threshold"] = types.YLeaf{"LowUtilizationThreshold", allocationSummary.LowUtilizationThreshold}
    allocationSummary.EntityData.Leafs["utilization"] = types.YLeaf{"Utilization", allocationSummary.Utilization}
    return &(allocationSummary.EntityData)
}

// AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6_Pools
// Pools data
type AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6_Pools struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Pool name. The type is string with length: 0..64.
    PoolName interface{}

    // VRF name. The type is string with length: 0..64.
    VrfName interface{}

    // Used allocations. The type is interface{} with range: 0..4294967295.
    Used interface{}

    // Excluded allocations. The type is interface{} with range: 0..4294967295.
    Excluded interface{}

    // Free allocations. The type is interface{} with range: 0..4294967295.
    Free interface{}

    // Total allocations. The type is interface{} with range: 0..4294967295.
    Total interface{}
}

func (pools *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6_Pools) GetEntityData() *types.CommonEntityData {
    pools.EntityData.YFilter = pools.YFilter
    pools.EntityData.YangName = "pools"
    pools.EntityData.BundleName = "cisco_ios_xr"
    pools.EntityData.ParentYangName = "ipv6"
    pools.EntityData.SegmentPath = "pools"
    pools.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pools.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pools.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pools.EntityData.Children = make(map[string]types.YChild)
    pools.EntityData.Leafs = make(map[string]types.YLeaf)
    pools.EntityData.Leafs["pool-name"] = types.YLeaf{"PoolName", pools.PoolName}
    pools.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", pools.VrfName}
    pools.EntityData.Leafs["used"] = types.YLeaf{"Used", pools.Used}
    pools.EntityData.Leafs["excluded"] = types.YLeaf{"Excluded", pools.Excluded}
    pools.EntityData.Leafs["free"] = types.YLeaf{"Free", pools.Free}
    pools.EntityData.Leafs["total"] = types.YLeaf{"Total", pools.Total}
    return &(pools.EntityData)
}

