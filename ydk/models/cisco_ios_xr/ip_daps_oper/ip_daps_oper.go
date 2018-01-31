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
    parent types.Entity
    YFilter yfilter.YFilter

    // Pool operational data for a particular location.
    Nodes AddressPoolService_Nodes
}

func (addressPoolService *AddressPoolService) GetFilter() yfilter.YFilter { return addressPoolService.YFilter }

func (addressPoolService *AddressPoolService) SetFilter(yf yfilter.YFilter) { addressPoolService.YFilter = yf }

func (addressPoolService *AddressPoolService) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (addressPoolService *AddressPoolService) GetSegmentPath() string {
    return "Cisco-IOS-XR-ip-daps-oper:address-pool-service"
}

func (addressPoolService *AddressPoolService) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &addressPoolService.Nodes
    }
    return nil
}

func (addressPoolService *AddressPoolService) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &addressPoolService.Nodes
    return children
}

func (addressPoolService *AddressPoolService) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (addressPoolService *AddressPoolService) GetBundleName() string { return "cisco_ios_xr" }

func (addressPoolService *AddressPoolService) GetYangName() string { return "address-pool-service" }

func (addressPoolService *AddressPoolService) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (addressPoolService *AddressPoolService) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (addressPoolService *AddressPoolService) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (addressPoolService *AddressPoolService) SetParent(parent types.Entity) { addressPoolService.parent = parent }

func (addressPoolService *AddressPoolService) GetParent() types.Entity { return addressPoolService.parent }

func (addressPoolService *AddressPoolService) GetParentYangName() string { return "Cisco-IOS-XR-ip-daps-oper" }

// AddressPoolService_Nodes
// Pool operational data for a particular location
type AddressPoolService_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Location. For eg., 0/1/CPU0. The type is slice of
    // AddressPoolService_Nodes_Node.
    Node []AddressPoolService_Nodes_Node
}

func (nodes *AddressPoolService_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *AddressPoolService_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *AddressPoolService_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *AddressPoolService_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *AddressPoolService_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AddressPoolService_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *AddressPoolService_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *AddressPoolService_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *AddressPoolService_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *AddressPoolService_Nodes) GetYangName() string { return "nodes" }

func (nodes *AddressPoolService_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *AddressPoolService_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *AddressPoolService_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *AddressPoolService_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *AddressPoolService_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *AddressPoolService_Nodes) GetParentYangName() string { return "address-pool-service" }

// AddressPoolService_Nodes_Node
// Location. For eg., 0/1/CPU0
type AddressPoolService_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // List of IPv4/IPv6 pool data.
    Pools AddressPoolService_Nodes_Node_Pools

    // Show total utilization for pool.
    TotalUtilization AddressPoolService_Nodes_Node_TotalUtilization

    // Pool VRF data.
    Vrfs AddressPoolService_Nodes_Node_Vrfs
}

func (node *AddressPoolService_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *AddressPoolService_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *AddressPoolService_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "pools" { return "Pools" }
    if yname == "total-utilization" { return "TotalUtilization" }
    if yname == "vrfs" { return "Vrfs" }
    return ""
}

func (node *AddressPoolService_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *AddressPoolService_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "pools" {
        return &node.Pools
    }
    if childYangName == "total-utilization" {
        return &node.TotalUtilization
    }
    if childYangName == "vrfs" {
        return &node.Vrfs
    }
    return nil
}

func (node *AddressPoolService_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["pools"] = &node.Pools
    children["total-utilization"] = &node.TotalUtilization
    children["vrfs"] = &node.Vrfs
    return children
}

func (node *AddressPoolService_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    return leafs
}

func (node *AddressPoolService_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *AddressPoolService_Nodes_Node) GetYangName() string { return "node" }

func (node *AddressPoolService_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *AddressPoolService_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *AddressPoolService_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *AddressPoolService_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *AddressPoolService_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *AddressPoolService_Nodes_Node) GetParentYangName() string { return "nodes" }

// AddressPoolService_Nodes_Node_Pools
// List of IPv4/IPv6 pool data
type AddressPoolService_Nodes_Node_Pools struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Pool data by pool name. The type is slice of
    // AddressPoolService_Nodes_Node_Pools_Pool.
    Pool []AddressPoolService_Nodes_Node_Pools_Pool
}

func (pools *AddressPoolService_Nodes_Node_Pools) GetFilter() yfilter.YFilter { return pools.YFilter }

func (pools *AddressPoolService_Nodes_Node_Pools) SetFilter(yf yfilter.YFilter) { pools.YFilter = yf }

func (pools *AddressPoolService_Nodes_Node_Pools) GetGoName(yname string) string {
    if yname == "pool" { return "Pool" }
    return ""
}

func (pools *AddressPoolService_Nodes_Node_Pools) GetSegmentPath() string {
    return "pools"
}

func (pools *AddressPoolService_Nodes_Node_Pools) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "pool" {
        for _, c := range pools.Pool {
            if pools.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AddressPoolService_Nodes_Node_Pools_Pool{}
        pools.Pool = append(pools.Pool, child)
        return &pools.Pool[len(pools.Pool)-1]
    }
    return nil
}

func (pools *AddressPoolService_Nodes_Node_Pools) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range pools.Pool {
        children[pools.Pool[i].GetSegmentPath()] = &pools.Pool[i]
    }
    return children
}

func (pools *AddressPoolService_Nodes_Node_Pools) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (pools *AddressPoolService_Nodes_Node_Pools) GetBundleName() string { return "cisco_ios_xr" }

func (pools *AddressPoolService_Nodes_Node_Pools) GetYangName() string { return "pools" }

func (pools *AddressPoolService_Nodes_Node_Pools) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pools *AddressPoolService_Nodes_Node_Pools) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pools *AddressPoolService_Nodes_Node_Pools) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pools *AddressPoolService_Nodes_Node_Pools) SetParent(parent types.Entity) { pools.parent = parent }

func (pools *AddressPoolService_Nodes_Node_Pools) GetParent() types.Entity { return pools.parent }

func (pools *AddressPoolService_Nodes_Node_Pools) GetParentYangName() string { return "node" }

// AddressPoolService_Nodes_Node_Pools_Pool
// Pool data by pool name
type AddressPoolService_Nodes_Node_Pools_Pool struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The pool name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    PoolName interface{}

    // Summary info for pool.
    AddressRanges AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges

    // Detailed info for the Pool.
    AllocatedAddresses AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses

    // Configuration info for pool.
    Configuration AddressPoolService_Nodes_Node_Pools_Pool_Configuration
}

func (pool *AddressPoolService_Nodes_Node_Pools_Pool) GetFilter() yfilter.YFilter { return pool.YFilter }

func (pool *AddressPoolService_Nodes_Node_Pools_Pool) SetFilter(yf yfilter.YFilter) { pool.YFilter = yf }

func (pool *AddressPoolService_Nodes_Node_Pools_Pool) GetGoName(yname string) string {
    if yname == "pool-name" { return "PoolName" }
    if yname == "address-ranges" { return "AddressRanges" }
    if yname == "allocated-addresses" { return "AllocatedAddresses" }
    if yname == "configuration" { return "Configuration" }
    return ""
}

func (pool *AddressPoolService_Nodes_Node_Pools_Pool) GetSegmentPath() string {
    return "pool" + "[pool-name='" + fmt.Sprintf("%v", pool.PoolName) + "']"
}

func (pool *AddressPoolService_Nodes_Node_Pools_Pool) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "address-ranges" {
        return &pool.AddressRanges
    }
    if childYangName == "allocated-addresses" {
        return &pool.AllocatedAddresses
    }
    if childYangName == "configuration" {
        return &pool.Configuration
    }
    return nil
}

func (pool *AddressPoolService_Nodes_Node_Pools_Pool) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["address-ranges"] = &pool.AddressRanges
    children["allocated-addresses"] = &pool.AllocatedAddresses
    children["configuration"] = &pool.Configuration
    return children
}

func (pool *AddressPoolService_Nodes_Node_Pools_Pool) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pool-name"] = pool.PoolName
    return leafs
}

func (pool *AddressPoolService_Nodes_Node_Pools_Pool) GetBundleName() string { return "cisco_ios_xr" }

func (pool *AddressPoolService_Nodes_Node_Pools_Pool) GetYangName() string { return "pool" }

func (pool *AddressPoolService_Nodes_Node_Pools_Pool) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pool *AddressPoolService_Nodes_Node_Pools_Pool) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pool *AddressPoolService_Nodes_Node_Pools_Pool) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pool *AddressPoolService_Nodes_Node_Pools_Pool) SetParent(parent types.Entity) { pool.parent = parent }

func (pool *AddressPoolService_Nodes_Node_Pools_Pool) GetParent() types.Entity { return pool.parent }

func (pool *AddressPoolService_Nodes_Node_Pools_Pool) GetParentYangName() string { return "pools" }

// AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges
// Summary info for pool
type AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Start Address of the Range. The type is slice of
    // AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange.
    AddressRange []AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange
}

func (addressRanges *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges) GetFilter() yfilter.YFilter { return addressRanges.YFilter }

func (addressRanges *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges) SetFilter(yf yfilter.YFilter) { addressRanges.YFilter = yf }

func (addressRanges *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges) GetGoName(yname string) string {
    if yname == "address-range" { return "AddressRange" }
    return ""
}

func (addressRanges *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges) GetSegmentPath() string {
    return "address-ranges"
}

func (addressRanges *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "address-range" {
        for _, c := range addressRanges.AddressRange {
            if addressRanges.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange{}
        addressRanges.AddressRange = append(addressRanges.AddressRange, child)
        return &addressRanges.AddressRange[len(addressRanges.AddressRange)-1]
    }
    return nil
}

func (addressRanges *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range addressRanges.AddressRange {
        children[addressRanges.AddressRange[i].GetSegmentPath()] = &addressRanges.AddressRange[i]
    }
    return children
}

func (addressRanges *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (addressRanges *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges) GetBundleName() string { return "cisco_ios_xr" }

func (addressRanges *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges) GetYangName() string { return "address-ranges" }

func (addressRanges *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (addressRanges *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (addressRanges *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (addressRanges *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges) SetParent(parent types.Entity) { addressRanges.parent = parent }

func (addressRanges *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges) GetParent() types.Entity { return addressRanges.parent }

func (addressRanges *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges) GetParentYangName() string { return "pool" }

// AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange
// Start Address of the Range
type AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP Address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
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

func (addressRange *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange) GetFilter() yfilter.YFilter { return addressRange.YFilter }

func (addressRange *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange) SetFilter(yf yfilter.YFilter) { addressRange.YFilter = yf }

func (addressRange *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange) GetGoName(yname string) string {
    if yname == "start-address" { return "StartAddress" }
    if yname == "pool-name" { return "PoolName" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "pool-scope" { return "PoolScope" }
    if yname == "allocated-addresses" { return "AllocatedAddresses" }
    if yname == "free-addresses" { return "FreeAddresses" }
    if yname == "excluded-addresses" { return "ExcludedAddresses" }
    if yname == "network-blocked-status" { return "NetworkBlockedStatus" }
    if yname == "network-blocked-status-trp" { return "NetworkBlockedStatusTrp" }
    if yname == "start-address-xr" { return "StartAddressXr" }
    if yname == "end-address" { return "EndAddress" }
    if yname == "default-router" { return "DefaultRouter" }
    return ""
}

func (addressRange *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange) GetSegmentPath() string {
    return "address-range" + "[start-address='" + fmt.Sprintf("%v", addressRange.StartAddress) + "']"
}

func (addressRange *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "start-address-xr" {
        return &addressRange.StartAddressXr
    }
    if childYangName == "end-address" {
        return &addressRange.EndAddress
    }
    if childYangName == "default-router" {
        return &addressRange.DefaultRouter
    }
    return nil
}

func (addressRange *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["start-address-xr"] = &addressRange.StartAddressXr
    children["end-address"] = &addressRange.EndAddress
    children["default-router"] = &addressRange.DefaultRouter
    return children
}

func (addressRange *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["start-address"] = addressRange.StartAddress
    leafs["pool-name"] = addressRange.PoolName
    leafs["vrf-name"] = addressRange.VrfName
    leafs["pool-scope"] = addressRange.PoolScope
    leafs["allocated-addresses"] = addressRange.AllocatedAddresses
    leafs["free-addresses"] = addressRange.FreeAddresses
    leafs["excluded-addresses"] = addressRange.ExcludedAddresses
    leafs["network-blocked-status"] = addressRange.NetworkBlockedStatus
    leafs["network-blocked-status-trp"] = addressRange.NetworkBlockedStatusTrp
    return leafs
}

func (addressRange *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange) GetBundleName() string { return "cisco_ios_xr" }

func (addressRange *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange) GetYangName() string { return "address-range" }

func (addressRange *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (addressRange *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (addressRange *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (addressRange *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange) SetParent(parent types.Entity) { addressRange.parent = parent }

func (addressRange *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange) GetParent() types.Entity { return addressRange.parent }

func (addressRange *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange) GetParentYangName() string { return "address-ranges" }

// AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_StartAddressXr
// Range start
type AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_StartAddressXr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Address.
    Address AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_StartAddressXr_Address
}

func (startAddressXr *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_StartAddressXr) GetFilter() yfilter.YFilter { return startAddressXr.YFilter }

func (startAddressXr *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_StartAddressXr) SetFilter(yf yfilter.YFilter) { startAddressXr.YFilter = yf }

func (startAddressXr *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_StartAddressXr) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    return ""
}

func (startAddressXr *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_StartAddressXr) GetSegmentPath() string {
    return "start-address-xr"
}

func (startAddressXr *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_StartAddressXr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "address" {
        return &startAddressXr.Address
    }
    return nil
}

func (startAddressXr *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_StartAddressXr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["address"] = &startAddressXr.Address
    return children
}

func (startAddressXr *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_StartAddressXr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (startAddressXr *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_StartAddressXr) GetBundleName() string { return "cisco_ios_xr" }

func (startAddressXr *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_StartAddressXr) GetYangName() string { return "start-address-xr" }

func (startAddressXr *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_StartAddressXr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (startAddressXr *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_StartAddressXr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (startAddressXr *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_StartAddressXr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (startAddressXr *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_StartAddressXr) SetParent(parent types.Entity) { startAddressXr.parent = parent }

func (startAddressXr *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_StartAddressXr) GetParent() types.Entity { return startAddressXr.parent }

func (startAddressXr *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_StartAddressXr) GetParentYangName() string { return "address-range" }

// AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_StartAddressXr_Address
// Address
type AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_StartAddressXr_Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AddressFamily. The type is IpAddr.
    AddressFamily interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_StartAddressXr_Address) GetFilter() yfilter.YFilter { return address.YFilter }

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_StartAddressXr_Address) SetFilter(yf yfilter.YFilter) { address.YFilter = yf }

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_StartAddressXr_Address) GetGoName(yname string) string {
    if yname == "address-family" { return "AddressFamily" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_StartAddressXr_Address) GetSegmentPath() string {
    return "address"
}

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_StartAddressXr_Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_StartAddressXr_Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_StartAddressXr_Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address-family"] = address.AddressFamily
    leafs["ipv4-address"] = address.Ipv4Address
    leafs["ipv6-address"] = address.Ipv6Address
    return leafs
}

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_StartAddressXr_Address) GetBundleName() string { return "cisco_ios_xr" }

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_StartAddressXr_Address) GetYangName() string { return "address" }

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_StartAddressXr_Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_StartAddressXr_Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_StartAddressXr_Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_StartAddressXr_Address) SetParent(parent types.Entity) { address.parent = parent }

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_StartAddressXr_Address) GetParent() types.Entity { return address.parent }

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_StartAddressXr_Address) GetParentYangName() string { return "start-address-xr" }

// AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_EndAddress
// Range end
type AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_EndAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Address.
    Address AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_EndAddress_Address
}

func (endAddress *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_EndAddress) GetFilter() yfilter.YFilter { return endAddress.YFilter }

func (endAddress *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_EndAddress) SetFilter(yf yfilter.YFilter) { endAddress.YFilter = yf }

func (endAddress *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_EndAddress) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    return ""
}

func (endAddress *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_EndAddress) GetSegmentPath() string {
    return "end-address"
}

func (endAddress *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_EndAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "address" {
        return &endAddress.Address
    }
    return nil
}

func (endAddress *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_EndAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["address"] = &endAddress.Address
    return children
}

func (endAddress *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_EndAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (endAddress *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_EndAddress) GetBundleName() string { return "cisco_ios_xr" }

func (endAddress *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_EndAddress) GetYangName() string { return "end-address" }

func (endAddress *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_EndAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (endAddress *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_EndAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (endAddress *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_EndAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (endAddress *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_EndAddress) SetParent(parent types.Entity) { endAddress.parent = parent }

func (endAddress *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_EndAddress) GetParent() types.Entity { return endAddress.parent }

func (endAddress *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_EndAddress) GetParentYangName() string { return "address-range" }

// AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_EndAddress_Address
// Address
type AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_EndAddress_Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AddressFamily. The type is IpAddr.
    AddressFamily interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_EndAddress_Address) GetFilter() yfilter.YFilter { return address.YFilter }

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_EndAddress_Address) SetFilter(yf yfilter.YFilter) { address.YFilter = yf }

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_EndAddress_Address) GetGoName(yname string) string {
    if yname == "address-family" { return "AddressFamily" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_EndAddress_Address) GetSegmentPath() string {
    return "address"
}

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_EndAddress_Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_EndAddress_Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_EndAddress_Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address-family"] = address.AddressFamily
    leafs["ipv4-address"] = address.Ipv4Address
    leafs["ipv6-address"] = address.Ipv6Address
    return leafs
}

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_EndAddress_Address) GetBundleName() string { return "cisco_ios_xr" }

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_EndAddress_Address) GetYangName() string { return "address" }

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_EndAddress_Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_EndAddress_Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_EndAddress_Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_EndAddress_Address) SetParent(parent types.Entity) { address.parent = parent }

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_EndAddress_Address) GetParent() types.Entity { return address.parent }

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_EndAddress_Address) GetParentYangName() string { return "end-address" }

// AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_DefaultRouter
// Default router
type AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_DefaultRouter struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Address.
    Address AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_DefaultRouter_Address
}

func (defaultRouter *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_DefaultRouter) GetFilter() yfilter.YFilter { return defaultRouter.YFilter }

func (defaultRouter *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_DefaultRouter) SetFilter(yf yfilter.YFilter) { defaultRouter.YFilter = yf }

func (defaultRouter *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_DefaultRouter) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    return ""
}

func (defaultRouter *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_DefaultRouter) GetSegmentPath() string {
    return "default-router"
}

func (defaultRouter *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_DefaultRouter) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "address" {
        return &defaultRouter.Address
    }
    return nil
}

func (defaultRouter *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_DefaultRouter) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["address"] = &defaultRouter.Address
    return children
}

func (defaultRouter *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_DefaultRouter) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (defaultRouter *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_DefaultRouter) GetBundleName() string { return "cisco_ios_xr" }

func (defaultRouter *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_DefaultRouter) GetYangName() string { return "default-router" }

func (defaultRouter *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_DefaultRouter) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (defaultRouter *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_DefaultRouter) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (defaultRouter *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_DefaultRouter) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (defaultRouter *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_DefaultRouter) SetParent(parent types.Entity) { defaultRouter.parent = parent }

func (defaultRouter *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_DefaultRouter) GetParent() types.Entity { return defaultRouter.parent }

func (defaultRouter *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_DefaultRouter) GetParentYangName() string { return "address-range" }

// AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_DefaultRouter_Address
// Address
type AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_DefaultRouter_Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AddressFamily. The type is IpAddr.
    AddressFamily interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_DefaultRouter_Address) GetFilter() yfilter.YFilter { return address.YFilter }

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_DefaultRouter_Address) SetFilter(yf yfilter.YFilter) { address.YFilter = yf }

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_DefaultRouter_Address) GetGoName(yname string) string {
    if yname == "address-family" { return "AddressFamily" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_DefaultRouter_Address) GetSegmentPath() string {
    return "address"
}

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_DefaultRouter_Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_DefaultRouter_Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_DefaultRouter_Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address-family"] = address.AddressFamily
    leafs["ipv4-address"] = address.Ipv4Address
    leafs["ipv6-address"] = address.Ipv6Address
    return leafs
}

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_DefaultRouter_Address) GetBundleName() string { return "cisco_ios_xr" }

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_DefaultRouter_Address) GetYangName() string { return "address" }

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_DefaultRouter_Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_DefaultRouter_Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_DefaultRouter_Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_DefaultRouter_Address) SetParent(parent types.Entity) { address.parent = parent }

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_DefaultRouter_Address) GetParent() types.Entity { return address.parent }

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AddressRanges_AddressRange_DefaultRouter_Address) GetParentYangName() string { return "default-router" }

// AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses
// Detailed info for the Pool
type AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses struct {
    parent types.Entity
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

func (allocatedAddresses *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses) GetFilter() yfilter.YFilter { return allocatedAddresses.YFilter }

func (allocatedAddresses *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses) SetFilter(yf yfilter.YFilter) { allocatedAddresses.YFilter = yf }

func (allocatedAddresses *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses) GetGoName(yname string) string {
    if yname == "pool-allocations" { return "PoolAllocations" }
    if yname == "address-range" { return "AddressRange" }
    if yname == "in-use-address" { return "InUseAddress" }
    return ""
}

func (allocatedAddresses *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses) GetSegmentPath() string {
    return "allocated-addresses"
}

func (allocatedAddresses *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "pool-allocations" {
        return &allocatedAddresses.PoolAllocations
    }
    if childYangName == "address-range" {
        for _, c := range allocatedAddresses.AddressRange {
            if allocatedAddresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange{}
        allocatedAddresses.AddressRange = append(allocatedAddresses.AddressRange, child)
        return &allocatedAddresses.AddressRange[len(allocatedAddresses.AddressRange)-1]
    }
    if childYangName == "in-use-address" {
        for _, c := range allocatedAddresses.InUseAddress {
            if allocatedAddresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress{}
        allocatedAddresses.InUseAddress = append(allocatedAddresses.InUseAddress, child)
        return &allocatedAddresses.InUseAddress[len(allocatedAddresses.InUseAddress)-1]
    }
    return nil
}

func (allocatedAddresses *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["pool-allocations"] = &allocatedAddresses.PoolAllocations
    for i := range allocatedAddresses.AddressRange {
        children[allocatedAddresses.AddressRange[i].GetSegmentPath()] = &allocatedAddresses.AddressRange[i]
    }
    for i := range allocatedAddresses.InUseAddress {
        children[allocatedAddresses.InUseAddress[i].GetSegmentPath()] = &allocatedAddresses.InUseAddress[i]
    }
    return children
}

func (allocatedAddresses *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (allocatedAddresses *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses) GetBundleName() string { return "cisco_ios_xr" }

func (allocatedAddresses *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses) GetYangName() string { return "allocated-addresses" }

func (allocatedAddresses *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (allocatedAddresses *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (allocatedAddresses *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (allocatedAddresses *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses) SetParent(parent types.Entity) { allocatedAddresses.parent = parent }

func (allocatedAddresses *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses) GetParent() types.Entity { return allocatedAddresses.parent }

func (allocatedAddresses *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses) GetParentYangName() string { return "pool" }

// AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_PoolAllocations
// Pool allocations
type AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_PoolAllocations struct {
    parent types.Entity
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

func (poolAllocations *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_PoolAllocations) GetFilter() yfilter.YFilter { return poolAllocations.YFilter }

func (poolAllocations *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_PoolAllocations) SetFilter(yf yfilter.YFilter) { poolAllocations.YFilter = yf }

func (poolAllocations *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_PoolAllocations) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "used" { return "Used" }
    if yname == "excluded" { return "Excluded" }
    if yname == "free" { return "Free" }
    if yname == "total" { return "Total" }
    if yname == "utilization" { return "Utilization" }
    if yname == "high-threshold" { return "HighThreshold" }
    if yname == "low-threshold" { return "LowThreshold" }
    return ""
}

func (poolAllocations *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_PoolAllocations) GetSegmentPath() string {
    return "pool-allocations"
}

func (poolAllocations *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_PoolAllocations) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "high-threshold" {
        return &poolAllocations.HighThreshold
    }
    if childYangName == "low-threshold" {
        return &poolAllocations.LowThreshold
    }
    return nil
}

func (poolAllocations *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_PoolAllocations) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["high-threshold"] = &poolAllocations.HighThreshold
    children["low-threshold"] = &poolAllocations.LowThreshold
    return children
}

func (poolAllocations *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_PoolAllocations) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = poolAllocations.VrfName
    leafs["used"] = poolAllocations.Used
    leafs["excluded"] = poolAllocations.Excluded
    leafs["free"] = poolAllocations.Free
    leafs["total"] = poolAllocations.Total
    leafs["utilization"] = poolAllocations.Utilization
    return leafs
}

func (poolAllocations *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_PoolAllocations) GetBundleName() string { return "cisco_ios_xr" }

func (poolAllocations *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_PoolAllocations) GetYangName() string { return "pool-allocations" }

func (poolAllocations *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_PoolAllocations) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (poolAllocations *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_PoolAllocations) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (poolAllocations *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_PoolAllocations) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (poolAllocations *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_PoolAllocations) SetParent(parent types.Entity) { poolAllocations.parent = parent }

func (poolAllocations *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_PoolAllocations) GetParent() types.Entity { return poolAllocations.parent }

func (poolAllocations *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_PoolAllocations) GetParentYangName() string { return "allocated-addresses" }

// AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_PoolAllocations_HighThreshold
// High threshold data
type AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_PoolAllocations_HighThreshold struct {
    parent types.Entity
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

func (highThreshold *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_PoolAllocations_HighThreshold) GetFilter() yfilter.YFilter { return highThreshold.YFilter }

func (highThreshold *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_PoolAllocations_HighThreshold) SetFilter(yf yfilter.YFilter) { highThreshold.YFilter = yf }

func (highThreshold *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_PoolAllocations_HighThreshold) GetGoName(yname string) string {
    if yname == "threshold" { return "Threshold" }
    if yname == "triggers" { return "Triggers" }
    if yname == "time-last-crossed" { return "TimeLastCrossed" }
    return ""
}

func (highThreshold *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_PoolAllocations_HighThreshold) GetSegmentPath() string {
    return "high-threshold"
}

func (highThreshold *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_PoolAllocations_HighThreshold) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (highThreshold *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_PoolAllocations_HighThreshold) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (highThreshold *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_PoolAllocations_HighThreshold) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["threshold"] = highThreshold.Threshold
    leafs["triggers"] = highThreshold.Triggers
    leafs["time-last-crossed"] = highThreshold.TimeLastCrossed
    return leafs
}

func (highThreshold *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_PoolAllocations_HighThreshold) GetBundleName() string { return "cisco_ios_xr" }

func (highThreshold *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_PoolAllocations_HighThreshold) GetYangName() string { return "high-threshold" }

func (highThreshold *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_PoolAllocations_HighThreshold) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (highThreshold *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_PoolAllocations_HighThreshold) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (highThreshold *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_PoolAllocations_HighThreshold) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (highThreshold *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_PoolAllocations_HighThreshold) SetParent(parent types.Entity) { highThreshold.parent = parent }

func (highThreshold *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_PoolAllocations_HighThreshold) GetParent() types.Entity { return highThreshold.parent }

func (highThreshold *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_PoolAllocations_HighThreshold) GetParentYangName() string { return "pool-allocations" }

// AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_PoolAllocations_LowThreshold
// Low threshold data
type AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_PoolAllocations_LowThreshold struct {
    parent types.Entity
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

func (lowThreshold *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_PoolAllocations_LowThreshold) GetFilter() yfilter.YFilter { return lowThreshold.YFilter }

func (lowThreshold *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_PoolAllocations_LowThreshold) SetFilter(yf yfilter.YFilter) { lowThreshold.YFilter = yf }

func (lowThreshold *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_PoolAllocations_LowThreshold) GetGoName(yname string) string {
    if yname == "threshold" { return "Threshold" }
    if yname == "triggers" { return "Triggers" }
    if yname == "time-last-crossed" { return "TimeLastCrossed" }
    return ""
}

func (lowThreshold *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_PoolAllocations_LowThreshold) GetSegmentPath() string {
    return "low-threshold"
}

func (lowThreshold *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_PoolAllocations_LowThreshold) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lowThreshold *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_PoolAllocations_LowThreshold) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lowThreshold *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_PoolAllocations_LowThreshold) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["threshold"] = lowThreshold.Threshold
    leafs["triggers"] = lowThreshold.Triggers
    leafs["time-last-crossed"] = lowThreshold.TimeLastCrossed
    return leafs
}

func (lowThreshold *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_PoolAllocations_LowThreshold) GetBundleName() string { return "cisco_ios_xr" }

func (lowThreshold *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_PoolAllocations_LowThreshold) GetYangName() string { return "low-threshold" }

func (lowThreshold *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_PoolAllocations_LowThreshold) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lowThreshold *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_PoolAllocations_LowThreshold) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lowThreshold *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_PoolAllocations_LowThreshold) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lowThreshold *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_PoolAllocations_LowThreshold) SetParent(parent types.Entity) { lowThreshold.parent = parent }

func (lowThreshold *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_PoolAllocations_LowThreshold) GetParent() types.Entity { return lowThreshold.parent }

func (lowThreshold *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_PoolAllocations_LowThreshold) GetParentYangName() string { return "pool-allocations" }

// AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange
// Address ranges
type AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange struct {
    parent types.Entity
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

func (addressRange *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange) GetFilter() yfilter.YFilter { return addressRange.YFilter }

func (addressRange *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange) SetFilter(yf yfilter.YFilter) { addressRange.YFilter = yf }

func (addressRange *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange) GetGoName(yname string) string {
    if yname == "used" { return "Used" }
    if yname == "excluded" { return "Excluded" }
    if yname == "free" { return "Free" }
    if yname == "start-address" { return "StartAddress" }
    if yname == "end-address" { return "EndAddress" }
    return ""
}

func (addressRange *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange) GetSegmentPath() string {
    return "address-range"
}

func (addressRange *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "start-address" {
        return &addressRange.StartAddress
    }
    if childYangName == "end-address" {
        return &addressRange.EndAddress
    }
    return nil
}

func (addressRange *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["start-address"] = &addressRange.StartAddress
    children["end-address"] = &addressRange.EndAddress
    return children
}

func (addressRange *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["used"] = addressRange.Used
    leafs["excluded"] = addressRange.Excluded
    leafs["free"] = addressRange.Free
    return leafs
}

func (addressRange *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange) GetBundleName() string { return "cisco_ios_xr" }

func (addressRange *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange) GetYangName() string { return "address-range" }

func (addressRange *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (addressRange *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (addressRange *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (addressRange *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange) SetParent(parent types.Entity) { addressRange.parent = parent }

func (addressRange *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange) GetParent() types.Entity { return addressRange.parent }

func (addressRange *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange) GetParentYangName() string { return "allocated-addresses" }

// AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_StartAddress
// Range start
type AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_StartAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Address.
    Address AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_StartAddress_Address
}

func (startAddress *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_StartAddress) GetFilter() yfilter.YFilter { return startAddress.YFilter }

func (startAddress *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_StartAddress) SetFilter(yf yfilter.YFilter) { startAddress.YFilter = yf }

func (startAddress *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_StartAddress) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    return ""
}

func (startAddress *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_StartAddress) GetSegmentPath() string {
    return "start-address"
}

func (startAddress *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_StartAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "address" {
        return &startAddress.Address
    }
    return nil
}

func (startAddress *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_StartAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["address"] = &startAddress.Address
    return children
}

func (startAddress *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_StartAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (startAddress *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_StartAddress) GetBundleName() string { return "cisco_ios_xr" }

func (startAddress *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_StartAddress) GetYangName() string { return "start-address" }

func (startAddress *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_StartAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (startAddress *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_StartAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (startAddress *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_StartAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (startAddress *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_StartAddress) SetParent(parent types.Entity) { startAddress.parent = parent }

func (startAddress *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_StartAddress) GetParent() types.Entity { return startAddress.parent }

func (startAddress *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_StartAddress) GetParentYangName() string { return "address-range" }

// AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_StartAddress_Address
// Address
type AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_StartAddress_Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AddressFamily. The type is IpAddr.
    AddressFamily interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_StartAddress_Address) GetFilter() yfilter.YFilter { return address.YFilter }

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_StartAddress_Address) SetFilter(yf yfilter.YFilter) { address.YFilter = yf }

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_StartAddress_Address) GetGoName(yname string) string {
    if yname == "address-family" { return "AddressFamily" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_StartAddress_Address) GetSegmentPath() string {
    return "address"
}

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_StartAddress_Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_StartAddress_Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_StartAddress_Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address-family"] = address.AddressFamily
    leafs["ipv4-address"] = address.Ipv4Address
    leafs["ipv6-address"] = address.Ipv6Address
    return leafs
}

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_StartAddress_Address) GetBundleName() string { return "cisco_ios_xr" }

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_StartAddress_Address) GetYangName() string { return "address" }

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_StartAddress_Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_StartAddress_Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_StartAddress_Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_StartAddress_Address) SetParent(parent types.Entity) { address.parent = parent }

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_StartAddress_Address) GetParent() types.Entity { return address.parent }

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_StartAddress_Address) GetParentYangName() string { return "start-address" }

// AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_EndAddress
// Range end
type AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_EndAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Address.
    Address AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_EndAddress_Address
}

func (endAddress *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_EndAddress) GetFilter() yfilter.YFilter { return endAddress.YFilter }

func (endAddress *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_EndAddress) SetFilter(yf yfilter.YFilter) { endAddress.YFilter = yf }

func (endAddress *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_EndAddress) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    return ""
}

func (endAddress *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_EndAddress) GetSegmentPath() string {
    return "end-address"
}

func (endAddress *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_EndAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "address" {
        return &endAddress.Address
    }
    return nil
}

func (endAddress *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_EndAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["address"] = &endAddress.Address
    return children
}

func (endAddress *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_EndAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (endAddress *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_EndAddress) GetBundleName() string { return "cisco_ios_xr" }

func (endAddress *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_EndAddress) GetYangName() string { return "end-address" }

func (endAddress *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_EndAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (endAddress *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_EndAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (endAddress *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_EndAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (endAddress *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_EndAddress) SetParent(parent types.Entity) { endAddress.parent = parent }

func (endAddress *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_EndAddress) GetParent() types.Entity { return endAddress.parent }

func (endAddress *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_EndAddress) GetParentYangName() string { return "address-range" }

// AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_EndAddress_Address
// Address
type AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_EndAddress_Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AddressFamily. The type is IpAddr.
    AddressFamily interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_EndAddress_Address) GetFilter() yfilter.YFilter { return address.YFilter }

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_EndAddress_Address) SetFilter(yf yfilter.YFilter) { address.YFilter = yf }

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_EndAddress_Address) GetGoName(yname string) string {
    if yname == "address-family" { return "AddressFamily" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_EndAddress_Address) GetSegmentPath() string {
    return "address"
}

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_EndAddress_Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_EndAddress_Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_EndAddress_Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address-family"] = address.AddressFamily
    leafs["ipv4-address"] = address.Ipv4Address
    leafs["ipv6-address"] = address.Ipv6Address
    return leafs
}

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_EndAddress_Address) GetBundleName() string { return "cisco_ios_xr" }

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_EndAddress_Address) GetYangName() string { return "address" }

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_EndAddress_Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_EndAddress_Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_EndAddress_Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_EndAddress_Address) SetParent(parent types.Entity) { address.parent = parent }

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_EndAddress_Address) GetParent() types.Entity { return address.parent }

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_AddressRange_EndAddress_Address) GetParentYangName() string { return "end-address" }

// AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress
// In-use addresses
type AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Client type. The type is DapsClient.
    ClientType interface{}

    // Client address.
    Address AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress_Address
}

func (inUseAddress *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress) GetFilter() yfilter.YFilter { return inUseAddress.YFilter }

func (inUseAddress *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress) SetFilter(yf yfilter.YFilter) { inUseAddress.YFilter = yf }

func (inUseAddress *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress) GetGoName(yname string) string {
    if yname == "client-type" { return "ClientType" }
    if yname == "address" { return "Address" }
    return ""
}

func (inUseAddress *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress) GetSegmentPath() string {
    return "in-use-address"
}

func (inUseAddress *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "address" {
        return &inUseAddress.Address
    }
    return nil
}

func (inUseAddress *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["address"] = &inUseAddress.Address
    return children
}

func (inUseAddress *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["client-type"] = inUseAddress.ClientType
    return leafs
}

func (inUseAddress *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress) GetBundleName() string { return "cisco_ios_xr" }

func (inUseAddress *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress) GetYangName() string { return "in-use-address" }

func (inUseAddress *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inUseAddress *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inUseAddress *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inUseAddress *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress) SetParent(parent types.Entity) { inUseAddress.parent = parent }

func (inUseAddress *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress) GetParent() types.Entity { return inUseAddress.parent }

func (inUseAddress *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress) GetParentYangName() string { return "allocated-addresses" }

// AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress_Address
// Client address
type AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress_Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Address.
    Address AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress_Address_Address
}

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress_Address) GetFilter() yfilter.YFilter { return address.YFilter }

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress_Address) SetFilter(yf yfilter.YFilter) { address.YFilter = yf }

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress_Address) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    return ""
}

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress_Address) GetSegmentPath() string {
    return "address"
}

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress_Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "address" {
        return &address.Address
    }
    return nil
}

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress_Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["address"] = &address.Address
    return children
}

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress_Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress_Address) GetBundleName() string { return "cisco_ios_xr" }

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress_Address) GetYangName() string { return "address" }

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress_Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress_Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress_Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress_Address) SetParent(parent types.Entity) { address.parent = parent }

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress_Address) GetParent() types.Entity { return address.parent }

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress_Address) GetParentYangName() string { return "in-use-address" }

// AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress_Address_Address
// Address
type AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress_Address_Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AddressFamily. The type is IpAddr.
    AddressFamily interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress_Address_Address) GetFilter() yfilter.YFilter { return address.YFilter }

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress_Address_Address) SetFilter(yf yfilter.YFilter) { address.YFilter = yf }

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress_Address_Address) GetGoName(yname string) string {
    if yname == "address-family" { return "AddressFamily" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress_Address_Address) GetSegmentPath() string {
    return "address"
}

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress_Address_Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress_Address_Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress_Address_Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address-family"] = address.AddressFamily
    leafs["ipv4-address"] = address.Ipv4Address
    leafs["ipv6-address"] = address.Ipv6Address
    return leafs
}

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress_Address_Address) GetBundleName() string { return "cisco_ios_xr" }

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress_Address_Address) GetYangName() string { return "address" }

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress_Address_Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress_Address_Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress_Address_Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress_Address_Address) SetParent(parent types.Entity) { address.parent = parent }

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress_Address_Address) GetParent() types.Entity { return address.parent }

func (address *AddressPoolService_Nodes_Node_Pools_Pool_AllocatedAddresses_InUseAddress_Address_Address) GetParentYangName() string { return "address" }

// AddressPoolService_Nodes_Node_Pools_Pool_Configuration
// Configuration info for pool
type AddressPoolService_Nodes_Node_Pools_Pool_Configuration struct {
    parent types.Entity
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

func (configuration *AddressPoolService_Nodes_Node_Pools_Pool_Configuration) GetFilter() yfilter.YFilter { return configuration.YFilter }

func (configuration *AddressPoolService_Nodes_Node_Pools_Pool_Configuration) SetFilter(yf yfilter.YFilter) { configuration.YFilter = yf }

func (configuration *AddressPoolService_Nodes_Node_Pools_Pool_Configuration) GetGoName(yname string) string {
    if yname == "pool-name" { return "PoolName" }
    if yname == "pool-id" { return "PoolId" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "pool-scope" { return "PoolScope" }
    if yname == "pool-prefix-length" { return "PoolPrefixLength" }
    if yname == "high-utilization-mark" { return "HighUtilizationMark" }
    if yname == "low-utilization-mark" { return "LowUtilizationMark" }
    if yname == "current-utilization" { return "CurrentUtilization" }
    if yname == "utilization-high-count" { return "UtilizationHighCount" }
    if yname == "utilization-low-count" { return "UtilizationLowCount" }
    return ""
}

func (configuration *AddressPoolService_Nodes_Node_Pools_Pool_Configuration) GetSegmentPath() string {
    return "configuration"
}

func (configuration *AddressPoolService_Nodes_Node_Pools_Pool_Configuration) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (configuration *AddressPoolService_Nodes_Node_Pools_Pool_Configuration) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (configuration *AddressPoolService_Nodes_Node_Pools_Pool_Configuration) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pool-name"] = configuration.PoolName
    leafs["pool-id"] = configuration.PoolId
    leafs["vrf-name"] = configuration.VrfName
    leafs["pool-scope"] = configuration.PoolScope
    leafs["pool-prefix-length"] = configuration.PoolPrefixLength
    leafs["high-utilization-mark"] = configuration.HighUtilizationMark
    leafs["low-utilization-mark"] = configuration.LowUtilizationMark
    leafs["current-utilization"] = configuration.CurrentUtilization
    leafs["utilization-high-count"] = configuration.UtilizationHighCount
    leafs["utilization-low-count"] = configuration.UtilizationLowCount
    return leafs
}

func (configuration *AddressPoolService_Nodes_Node_Pools_Pool_Configuration) GetBundleName() string { return "cisco_ios_xr" }

func (configuration *AddressPoolService_Nodes_Node_Pools_Pool_Configuration) GetYangName() string { return "configuration" }

func (configuration *AddressPoolService_Nodes_Node_Pools_Pool_Configuration) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configuration *AddressPoolService_Nodes_Node_Pools_Pool_Configuration) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configuration *AddressPoolService_Nodes_Node_Pools_Pool_Configuration) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configuration *AddressPoolService_Nodes_Node_Pools_Pool_Configuration) SetParent(parent types.Entity) { configuration.parent = parent }

func (configuration *AddressPoolService_Nodes_Node_Pools_Pool_Configuration) GetParent() types.Entity { return configuration.parent }

func (configuration *AddressPoolService_Nodes_Node_Pools_Pool_Configuration) GetParentYangName() string { return "pool" }

// AddressPoolService_Nodes_Node_TotalUtilization
// Show total utilization for pool
type AddressPoolService_Nodes_Node_TotalUtilization struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // High utilization mark. The type is interface{} with range: 0..255.
    TotalUtilizationHighMark interface{}

    // Low utilization mark. The type is interface{} with range: 0..255.
    TotalUtilizationLowMark interface{}

    // Current utilization. The type is interface{} with range: 0..255.
    CurrentTotalUtilization interface{}
}

func (totalUtilization *AddressPoolService_Nodes_Node_TotalUtilization) GetFilter() yfilter.YFilter { return totalUtilization.YFilter }

func (totalUtilization *AddressPoolService_Nodes_Node_TotalUtilization) SetFilter(yf yfilter.YFilter) { totalUtilization.YFilter = yf }

func (totalUtilization *AddressPoolService_Nodes_Node_TotalUtilization) GetGoName(yname string) string {
    if yname == "total-utilization-high-mark" { return "TotalUtilizationHighMark" }
    if yname == "total-utilization-low-mark" { return "TotalUtilizationLowMark" }
    if yname == "current-total-utilization" { return "CurrentTotalUtilization" }
    return ""
}

func (totalUtilization *AddressPoolService_Nodes_Node_TotalUtilization) GetSegmentPath() string {
    return "total-utilization"
}

func (totalUtilization *AddressPoolService_Nodes_Node_TotalUtilization) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (totalUtilization *AddressPoolService_Nodes_Node_TotalUtilization) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (totalUtilization *AddressPoolService_Nodes_Node_TotalUtilization) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["total-utilization-high-mark"] = totalUtilization.TotalUtilizationHighMark
    leafs["total-utilization-low-mark"] = totalUtilization.TotalUtilizationLowMark
    leafs["current-total-utilization"] = totalUtilization.CurrentTotalUtilization
    return leafs
}

func (totalUtilization *AddressPoolService_Nodes_Node_TotalUtilization) GetBundleName() string { return "cisco_ios_xr" }

func (totalUtilization *AddressPoolService_Nodes_Node_TotalUtilization) GetYangName() string { return "total-utilization" }

func (totalUtilization *AddressPoolService_Nodes_Node_TotalUtilization) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (totalUtilization *AddressPoolService_Nodes_Node_TotalUtilization) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (totalUtilization *AddressPoolService_Nodes_Node_TotalUtilization) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (totalUtilization *AddressPoolService_Nodes_Node_TotalUtilization) SetParent(parent types.Entity) { totalUtilization.parent = parent }

func (totalUtilization *AddressPoolService_Nodes_Node_TotalUtilization) GetParent() types.Entity { return totalUtilization.parent }

func (totalUtilization *AddressPoolService_Nodes_Node_TotalUtilization) GetParentYangName() string { return "node" }

// AddressPoolService_Nodes_Node_Vrfs
// Pool VRF data
type AddressPoolService_Nodes_Node_Vrfs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF level Pool information. The type is slice of
    // AddressPoolService_Nodes_Node_Vrfs_Vrf.
    Vrf []AddressPoolService_Nodes_Node_Vrfs_Vrf
}

func (vrfs *AddressPoolService_Nodes_Node_Vrfs) GetFilter() yfilter.YFilter { return vrfs.YFilter }

func (vrfs *AddressPoolService_Nodes_Node_Vrfs) SetFilter(yf yfilter.YFilter) { vrfs.YFilter = yf }

func (vrfs *AddressPoolService_Nodes_Node_Vrfs) GetGoName(yname string) string {
    if yname == "vrf" { return "Vrf" }
    return ""
}

func (vrfs *AddressPoolService_Nodes_Node_Vrfs) GetSegmentPath() string {
    return "vrfs"
}

func (vrfs *AddressPoolService_Nodes_Node_Vrfs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf" {
        for _, c := range vrfs.Vrf {
            if vrfs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AddressPoolService_Nodes_Node_Vrfs_Vrf{}
        vrfs.Vrf = append(vrfs.Vrf, child)
        return &vrfs.Vrf[len(vrfs.Vrf)-1]
    }
    return nil
}

func (vrfs *AddressPoolService_Nodes_Node_Vrfs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vrfs.Vrf {
        children[vrfs.Vrf[i].GetSegmentPath()] = &vrfs.Vrf[i]
    }
    return children
}

func (vrfs *AddressPoolService_Nodes_Node_Vrfs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vrfs *AddressPoolService_Nodes_Node_Vrfs) GetBundleName() string { return "cisco_ios_xr" }

func (vrfs *AddressPoolService_Nodes_Node_Vrfs) GetYangName() string { return "vrfs" }

func (vrfs *AddressPoolService_Nodes_Node_Vrfs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfs *AddressPoolService_Nodes_Node_Vrfs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfs *AddressPoolService_Nodes_Node_Vrfs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfs *AddressPoolService_Nodes_Node_Vrfs) SetParent(parent types.Entity) { vrfs.parent = parent }

func (vrfs *AddressPoolService_Nodes_Node_Vrfs) GetParent() types.Entity { return vrfs.parent }

func (vrfs *AddressPoolService_Nodes_Node_Vrfs) GetParentYangName() string { return "node" }

// AddressPoolService_Nodes_Node_Vrfs_Vrf
// VRF level Pool information
type AddressPoolService_Nodes_Node_Vrfs_Vrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The VRF name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // IPv4 pool VRF data.
    Ipv4 AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4

    // IPv6 Pool VRF data.
    Ipv6 AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6
}

func (vrf *AddressPoolService_Nodes_Node_Vrfs_Vrf) GetFilter() yfilter.YFilter { return vrf.YFilter }

func (vrf *AddressPoolService_Nodes_Node_Vrfs_Vrf) SetFilter(yf yfilter.YFilter) { vrf.YFilter = yf }

func (vrf *AddressPoolService_Nodes_Node_Vrfs_Vrf) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "ipv4" { return "Ipv4" }
    if yname == "ipv6" { return "Ipv6" }
    return ""
}

func (vrf *AddressPoolService_Nodes_Node_Vrfs_Vrf) GetSegmentPath() string {
    return "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
}

func (vrf *AddressPoolService_Nodes_Node_Vrfs_Vrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv4" {
        return &vrf.Ipv4
    }
    if childYangName == "ipv6" {
        return &vrf.Ipv6
    }
    return nil
}

func (vrf *AddressPoolService_Nodes_Node_Vrfs_Vrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ipv4"] = &vrf.Ipv4
    children["ipv6"] = &vrf.Ipv6
    return children
}

func (vrf *AddressPoolService_Nodes_Node_Vrfs_Vrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = vrf.VrfName
    return leafs
}

func (vrf *AddressPoolService_Nodes_Node_Vrfs_Vrf) GetBundleName() string { return "cisco_ios_xr" }

func (vrf *AddressPoolService_Nodes_Node_Vrfs_Vrf) GetYangName() string { return "vrf" }

func (vrf *AddressPoolService_Nodes_Node_Vrfs_Vrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrf *AddressPoolService_Nodes_Node_Vrfs_Vrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrf *AddressPoolService_Nodes_Node_Vrfs_Vrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrf *AddressPoolService_Nodes_Node_Vrfs_Vrf) SetParent(parent types.Entity) { vrf.parent = parent }

func (vrf *AddressPoolService_Nodes_Node_Vrfs_Vrf) GetParent() types.Entity { return vrf.parent }

func (vrf *AddressPoolService_Nodes_Node_Vrfs_Vrf) GetParentYangName() string { return "vrfs" }

// AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4
// IPv4 pool VRF data
type AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Allocation summary.
    AllocationSummary AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4_AllocationSummary

    // Pools data. The type is slice of
    // AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4_Pools.
    Pools []AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4_Pools
}

func (ipv4 *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4) GetFilter() yfilter.YFilter { return ipv4.YFilter }

func (ipv4 *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4) SetFilter(yf yfilter.YFilter) { ipv4.YFilter = yf }

func (ipv4 *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4) GetGoName(yname string) string {
    if yname == "allocation-summary" { return "AllocationSummary" }
    if yname == "pools" { return "Pools" }
    return ""
}

func (ipv4 *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4) GetSegmentPath() string {
    return "ipv4"
}

func (ipv4 *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "allocation-summary" {
        return &ipv4.AllocationSummary
    }
    if childYangName == "pools" {
        for _, c := range ipv4.Pools {
            if ipv4.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4_Pools{}
        ipv4.Pools = append(ipv4.Pools, child)
        return &ipv4.Pools[len(ipv4.Pools)-1]
    }
    return nil
}

func (ipv4 *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["allocation-summary"] = &ipv4.AllocationSummary
    for i := range ipv4.Pools {
        children[ipv4.Pools[i].GetSegmentPath()] = &ipv4.Pools[i]
    }
    return children
}

func (ipv4 *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv4 *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4 *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4) GetYangName() string { return "ipv4" }

func (ipv4 *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4 *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4 *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4 *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4) SetParent(parent types.Entity) { ipv4.parent = parent }

func (ipv4 *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4) GetParent() types.Entity { return ipv4.parent }

func (ipv4 *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4) GetParentYangName() string { return "vrf" }

// AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4_AllocationSummary
// Allocation summary
type AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4_AllocationSummary struct {
    parent types.Entity
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

func (allocationSummary *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4_AllocationSummary) GetFilter() yfilter.YFilter { return allocationSummary.YFilter }

func (allocationSummary *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4_AllocationSummary) SetFilter(yf yfilter.YFilter) { allocationSummary.YFilter = yf }

func (allocationSummary *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4_AllocationSummary) GetGoName(yname string) string {
    if yname == "used" { return "Used" }
    if yname == "excluded" { return "Excluded" }
    if yname == "free" { return "Free" }
    if yname == "total" { return "Total" }
    if yname == "high-utilization-threshold" { return "HighUtilizationThreshold" }
    if yname == "low-utilization-threshold" { return "LowUtilizationThreshold" }
    if yname == "utilization" { return "Utilization" }
    return ""
}

func (allocationSummary *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4_AllocationSummary) GetSegmentPath() string {
    return "allocation-summary"
}

func (allocationSummary *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4_AllocationSummary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (allocationSummary *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4_AllocationSummary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (allocationSummary *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4_AllocationSummary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["used"] = allocationSummary.Used
    leafs["excluded"] = allocationSummary.Excluded
    leafs["free"] = allocationSummary.Free
    leafs["total"] = allocationSummary.Total
    leafs["high-utilization-threshold"] = allocationSummary.HighUtilizationThreshold
    leafs["low-utilization-threshold"] = allocationSummary.LowUtilizationThreshold
    leafs["utilization"] = allocationSummary.Utilization
    return leafs
}

func (allocationSummary *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4_AllocationSummary) GetBundleName() string { return "cisco_ios_xr" }

func (allocationSummary *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4_AllocationSummary) GetYangName() string { return "allocation-summary" }

func (allocationSummary *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4_AllocationSummary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (allocationSummary *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4_AllocationSummary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (allocationSummary *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4_AllocationSummary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (allocationSummary *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4_AllocationSummary) SetParent(parent types.Entity) { allocationSummary.parent = parent }

func (allocationSummary *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4_AllocationSummary) GetParent() types.Entity { return allocationSummary.parent }

func (allocationSummary *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4_AllocationSummary) GetParentYangName() string { return "ipv4" }

// AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4_Pools
// Pools data
type AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4_Pools struct {
    parent types.Entity
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

func (pools *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4_Pools) GetFilter() yfilter.YFilter { return pools.YFilter }

func (pools *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4_Pools) SetFilter(yf yfilter.YFilter) { pools.YFilter = yf }

func (pools *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4_Pools) GetGoName(yname string) string {
    if yname == "pool-name" { return "PoolName" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "used" { return "Used" }
    if yname == "excluded" { return "Excluded" }
    if yname == "free" { return "Free" }
    if yname == "total" { return "Total" }
    return ""
}

func (pools *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4_Pools) GetSegmentPath() string {
    return "pools"
}

func (pools *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4_Pools) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pools *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4_Pools) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pools *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4_Pools) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pool-name"] = pools.PoolName
    leafs["vrf-name"] = pools.VrfName
    leafs["used"] = pools.Used
    leafs["excluded"] = pools.Excluded
    leafs["free"] = pools.Free
    leafs["total"] = pools.Total
    return leafs
}

func (pools *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4_Pools) GetBundleName() string { return "cisco_ios_xr" }

func (pools *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4_Pools) GetYangName() string { return "pools" }

func (pools *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4_Pools) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pools *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4_Pools) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pools *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4_Pools) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pools *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4_Pools) SetParent(parent types.Entity) { pools.parent = parent }

func (pools *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4_Pools) GetParent() types.Entity { return pools.parent }

func (pools *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv4_Pools) GetParentYangName() string { return "ipv4" }

// AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6
// IPv6 Pool VRF data
type AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Allocation summary.
    AllocationSummary AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6_AllocationSummary

    // Pools data. The type is slice of
    // AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6_Pools.
    Pools []AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6_Pools
}

func (ipv6 *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6) GetFilter() yfilter.YFilter { return ipv6.YFilter }

func (ipv6 *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6) SetFilter(yf yfilter.YFilter) { ipv6.YFilter = yf }

func (ipv6 *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6) GetGoName(yname string) string {
    if yname == "allocation-summary" { return "AllocationSummary" }
    if yname == "pools" { return "Pools" }
    return ""
}

func (ipv6 *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6) GetSegmentPath() string {
    return "ipv6"
}

func (ipv6 *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "allocation-summary" {
        return &ipv6.AllocationSummary
    }
    if childYangName == "pools" {
        for _, c := range ipv6.Pools {
            if ipv6.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6_Pools{}
        ipv6.Pools = append(ipv6.Pools, child)
        return &ipv6.Pools[len(ipv6.Pools)-1]
    }
    return nil
}

func (ipv6 *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["allocation-summary"] = &ipv6.AllocationSummary
    for i := range ipv6.Pools {
        children[ipv6.Pools[i].GetSegmentPath()] = &ipv6.Pools[i]
    }
    return children
}

func (ipv6 *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv6 *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6 *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6) GetYangName() string { return "ipv6" }

func (ipv6 *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6 *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6 *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6 *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6) SetParent(parent types.Entity) { ipv6.parent = parent }

func (ipv6 *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6) GetParent() types.Entity { return ipv6.parent }

func (ipv6 *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6) GetParentYangName() string { return "vrf" }

// AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6_AllocationSummary
// Allocation summary
type AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6_AllocationSummary struct {
    parent types.Entity
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

func (allocationSummary *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6_AllocationSummary) GetFilter() yfilter.YFilter { return allocationSummary.YFilter }

func (allocationSummary *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6_AllocationSummary) SetFilter(yf yfilter.YFilter) { allocationSummary.YFilter = yf }

func (allocationSummary *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6_AllocationSummary) GetGoName(yname string) string {
    if yname == "used" { return "Used" }
    if yname == "excluded" { return "Excluded" }
    if yname == "free" { return "Free" }
    if yname == "total" { return "Total" }
    if yname == "high-utilization-threshold" { return "HighUtilizationThreshold" }
    if yname == "low-utilization-threshold" { return "LowUtilizationThreshold" }
    if yname == "utilization" { return "Utilization" }
    return ""
}

func (allocationSummary *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6_AllocationSummary) GetSegmentPath() string {
    return "allocation-summary"
}

func (allocationSummary *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6_AllocationSummary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (allocationSummary *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6_AllocationSummary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (allocationSummary *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6_AllocationSummary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["used"] = allocationSummary.Used
    leafs["excluded"] = allocationSummary.Excluded
    leafs["free"] = allocationSummary.Free
    leafs["total"] = allocationSummary.Total
    leafs["high-utilization-threshold"] = allocationSummary.HighUtilizationThreshold
    leafs["low-utilization-threshold"] = allocationSummary.LowUtilizationThreshold
    leafs["utilization"] = allocationSummary.Utilization
    return leafs
}

func (allocationSummary *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6_AllocationSummary) GetBundleName() string { return "cisco_ios_xr" }

func (allocationSummary *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6_AllocationSummary) GetYangName() string { return "allocation-summary" }

func (allocationSummary *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6_AllocationSummary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (allocationSummary *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6_AllocationSummary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (allocationSummary *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6_AllocationSummary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (allocationSummary *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6_AllocationSummary) SetParent(parent types.Entity) { allocationSummary.parent = parent }

func (allocationSummary *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6_AllocationSummary) GetParent() types.Entity { return allocationSummary.parent }

func (allocationSummary *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6_AllocationSummary) GetParentYangName() string { return "ipv6" }

// AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6_Pools
// Pools data
type AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6_Pools struct {
    parent types.Entity
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

func (pools *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6_Pools) GetFilter() yfilter.YFilter { return pools.YFilter }

func (pools *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6_Pools) SetFilter(yf yfilter.YFilter) { pools.YFilter = yf }

func (pools *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6_Pools) GetGoName(yname string) string {
    if yname == "pool-name" { return "PoolName" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "used" { return "Used" }
    if yname == "excluded" { return "Excluded" }
    if yname == "free" { return "Free" }
    if yname == "total" { return "Total" }
    return ""
}

func (pools *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6_Pools) GetSegmentPath() string {
    return "pools"
}

func (pools *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6_Pools) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pools *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6_Pools) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pools *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6_Pools) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pool-name"] = pools.PoolName
    leafs["vrf-name"] = pools.VrfName
    leafs["used"] = pools.Used
    leafs["excluded"] = pools.Excluded
    leafs["free"] = pools.Free
    leafs["total"] = pools.Total
    return leafs
}

func (pools *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6_Pools) GetBundleName() string { return "cisco_ios_xr" }

func (pools *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6_Pools) GetYangName() string { return "pools" }

func (pools *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6_Pools) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pools *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6_Pools) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pools *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6_Pools) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pools *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6_Pools) SetParent(parent types.Entity) { pools.parent = parent }

func (pools *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6_Pools) GetParent() types.Entity { return pools.parent }

func (pools *AddressPoolService_Nodes_Node_Vrfs_Vrf_Ipv6_Pools) GetParentYangName() string { return "ipv6" }

