// This module contains a collection of YANG definitions
// for Cisco IOS-XR ip-daps package configuration.
// 
// This module contains definitions
// for the following management objects:
//   address-pool-service: Address Pool configuration data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ip_daps_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ip_daps_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ip-daps-cfg address-pool-service}", reflect.TypeOf(AddressPoolService{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ip-daps-cfg:address-pool-service", reflect.TypeOf(AddressPoolService{}))
}

// AddressPoolService
// Address Pool configuration data
type AddressPoolService struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enter VRF specific config mode.
    Vrfs AddressPoolService_Vrfs
}

func (addressPoolService *AddressPoolService) GetFilter() yfilter.YFilter { return addressPoolService.YFilter }

func (addressPoolService *AddressPoolService) SetFilter(yf yfilter.YFilter) { addressPoolService.YFilter = yf }

func (addressPoolService *AddressPoolService) GetGoName(yname string) string {
    if yname == "vrfs" { return "Vrfs" }
    return ""
}

func (addressPoolService *AddressPoolService) GetSegmentPath() string {
    return "Cisco-IOS-XR-ip-daps-cfg:address-pool-service"
}

func (addressPoolService *AddressPoolService) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrfs" {
        return &addressPoolService.Vrfs
    }
    return nil
}

func (addressPoolService *AddressPoolService) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["vrfs"] = &addressPoolService.Vrfs
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

func (addressPoolService *AddressPoolService) GetParentYangName() string { return "Cisco-IOS-XR-ip-daps-cfg" }

// AddressPoolService_Vrfs
// Enter VRF specific config mode
type AddressPoolService_Vrfs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify VRF Name. The type is slice of AddressPoolService_Vrfs_Vrf.
    Vrf []AddressPoolService_Vrfs_Vrf
}

func (vrfs *AddressPoolService_Vrfs) GetFilter() yfilter.YFilter { return vrfs.YFilter }

func (vrfs *AddressPoolService_Vrfs) SetFilter(yf yfilter.YFilter) { vrfs.YFilter = yf }

func (vrfs *AddressPoolService_Vrfs) GetGoName(yname string) string {
    if yname == "vrf" { return "Vrf" }
    return ""
}

func (vrfs *AddressPoolService_Vrfs) GetSegmentPath() string {
    return "vrfs"
}

func (vrfs *AddressPoolService_Vrfs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf" {
        for _, c := range vrfs.Vrf {
            if vrfs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AddressPoolService_Vrfs_Vrf{}
        vrfs.Vrf = append(vrfs.Vrf, child)
        return &vrfs.Vrf[len(vrfs.Vrf)-1]
    }
    return nil
}

func (vrfs *AddressPoolService_Vrfs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vrfs.Vrf {
        children[vrfs.Vrf[i].GetSegmentPath()] = &vrfs.Vrf[i]
    }
    return children
}

func (vrfs *AddressPoolService_Vrfs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vrfs *AddressPoolService_Vrfs) GetBundleName() string { return "cisco_ios_xr" }

func (vrfs *AddressPoolService_Vrfs) GetYangName() string { return "vrfs" }

func (vrfs *AddressPoolService_Vrfs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfs *AddressPoolService_Vrfs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfs *AddressPoolService_Vrfs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfs *AddressPoolService_Vrfs) SetParent(parent types.Entity) { vrfs.parent = parent }

func (vrfs *AddressPoolService_Vrfs) GetParent() types.Entity { return vrfs.parent }

func (vrfs *AddressPoolService_Vrfs) GetParentYangName() string { return "address-pool-service" }

// AddressPoolService_Vrfs_Vrf
// Specify VRF Name
type AddressPoolService_Vrfs_Vrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. none. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // Enter IPv6 specific mode.
    Ipv6 AddressPoolService_Vrfs_Vrf_Ipv6

    // Enter IPv4 specific configuration.
    Ipv4 AddressPoolService_Vrfs_Vrf_Ipv4
}

func (vrf *AddressPoolService_Vrfs_Vrf) GetFilter() yfilter.YFilter { return vrf.YFilter }

func (vrf *AddressPoolService_Vrfs_Vrf) SetFilter(yf yfilter.YFilter) { vrf.YFilter = yf }

func (vrf *AddressPoolService_Vrfs_Vrf) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "ipv6" { return "Ipv6" }
    if yname == "ipv4" { return "Ipv4" }
    return ""
}

func (vrf *AddressPoolService_Vrfs_Vrf) GetSegmentPath() string {
    return "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
}

func (vrf *AddressPoolService_Vrfs_Vrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv6" {
        return &vrf.Ipv6
    }
    if childYangName == "ipv4" {
        return &vrf.Ipv4
    }
    return nil
}

func (vrf *AddressPoolService_Vrfs_Vrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ipv6"] = &vrf.Ipv6
    children["ipv4"] = &vrf.Ipv4
    return children
}

func (vrf *AddressPoolService_Vrfs_Vrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = vrf.VrfName
    return leafs
}

func (vrf *AddressPoolService_Vrfs_Vrf) GetBundleName() string { return "cisco_ios_xr" }

func (vrf *AddressPoolService_Vrfs_Vrf) GetYangName() string { return "vrf" }

func (vrf *AddressPoolService_Vrfs_Vrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrf *AddressPoolService_Vrfs_Vrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrf *AddressPoolService_Vrfs_Vrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrf *AddressPoolService_Vrfs_Vrf) SetParent(parent types.Entity) { vrf.parent = parent }

func (vrf *AddressPoolService_Vrfs_Vrf) GetParent() types.Entity { return vrf.parent }

func (vrf *AddressPoolService_Vrfs_Vrf) GetParentYangName() string { return "vrfs" }

// AddressPoolService_Vrfs_Vrf_Ipv6
// Enter IPv6 specific mode
type AddressPoolService_Vrfs_Vrf_Ipv6 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 Pool Name.
    Pools AddressPoolService_Vrfs_Vrf_Ipv6_Pools
}

func (ipv6 *AddressPoolService_Vrfs_Vrf_Ipv6) GetFilter() yfilter.YFilter { return ipv6.YFilter }

func (ipv6 *AddressPoolService_Vrfs_Vrf_Ipv6) SetFilter(yf yfilter.YFilter) { ipv6.YFilter = yf }

func (ipv6 *AddressPoolService_Vrfs_Vrf_Ipv6) GetGoName(yname string) string {
    if yname == "pools" { return "Pools" }
    return ""
}

func (ipv6 *AddressPoolService_Vrfs_Vrf_Ipv6) GetSegmentPath() string {
    return "ipv6"
}

func (ipv6 *AddressPoolService_Vrfs_Vrf_Ipv6) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "pools" {
        return &ipv6.Pools
    }
    return nil
}

func (ipv6 *AddressPoolService_Vrfs_Vrf_Ipv6) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["pools"] = &ipv6.Pools
    return children
}

func (ipv6 *AddressPoolService_Vrfs_Vrf_Ipv6) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv6 *AddressPoolService_Vrfs_Vrf_Ipv6) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6 *AddressPoolService_Vrfs_Vrf_Ipv6) GetYangName() string { return "ipv6" }

func (ipv6 *AddressPoolService_Vrfs_Vrf_Ipv6) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6 *AddressPoolService_Vrfs_Vrf_Ipv6) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6 *AddressPoolService_Vrfs_Vrf_Ipv6) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6 *AddressPoolService_Vrfs_Vrf_Ipv6) SetParent(parent types.Entity) { ipv6.parent = parent }

func (ipv6 *AddressPoolService_Vrfs_Vrf_Ipv6) GetParent() types.Entity { return ipv6.parent }

func (ipv6 *AddressPoolService_Vrfs_Vrf_Ipv6) GetParentYangName() string { return "vrf" }

// AddressPoolService_Vrfs_Vrf_Ipv6_Pools
// IPv6 Pool Name
type AddressPoolService_Vrfs_Vrf_Ipv6_Pools struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enter the IPv6 Pool name. The type is slice of
    // AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool.
    Pool []AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool
}

func (pools *AddressPoolService_Vrfs_Vrf_Ipv6_Pools) GetFilter() yfilter.YFilter { return pools.YFilter }

func (pools *AddressPoolService_Vrfs_Vrf_Ipv6_Pools) SetFilter(yf yfilter.YFilter) { pools.YFilter = yf }

func (pools *AddressPoolService_Vrfs_Vrf_Ipv6_Pools) GetGoName(yname string) string {
    if yname == "pool" { return "Pool" }
    return ""
}

func (pools *AddressPoolService_Vrfs_Vrf_Ipv6_Pools) GetSegmentPath() string {
    return "pools"
}

func (pools *AddressPoolService_Vrfs_Vrf_Ipv6_Pools) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "pool" {
        for _, c := range pools.Pool {
            if pools.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool{}
        pools.Pool = append(pools.Pool, child)
        return &pools.Pool[len(pools.Pool)-1]
    }
    return nil
}

func (pools *AddressPoolService_Vrfs_Vrf_Ipv6_Pools) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range pools.Pool {
        children[pools.Pool[i].GetSegmentPath()] = &pools.Pool[i]
    }
    return children
}

func (pools *AddressPoolService_Vrfs_Vrf_Ipv6_Pools) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (pools *AddressPoolService_Vrfs_Vrf_Ipv6_Pools) GetBundleName() string { return "cisco_ios_xr" }

func (pools *AddressPoolService_Vrfs_Vrf_Ipv6_Pools) GetYangName() string { return "pools" }

func (pools *AddressPoolService_Vrfs_Vrf_Ipv6_Pools) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pools *AddressPoolService_Vrfs_Vrf_Ipv6_Pools) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pools *AddressPoolService_Vrfs_Vrf_Ipv6_Pools) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pools *AddressPoolService_Vrfs_Vrf_Ipv6_Pools) SetParent(parent types.Entity) { pools.parent = parent }

func (pools *AddressPoolService_Vrfs_Vrf_Ipv6_Pools) GetParent() types.Entity { return pools.parent }

func (pools *AddressPoolService_Vrfs_Vrf_Ipv6_Pools) GetParentYangName() string { return "ipv6" }

// AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool
// Enter the IPv6 Pool name
type AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Enter the IPv6 Pool name. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Ipv6PoolName interface{}

    // Enter the prefix-length for the Pool. The type is interface{} with range:
    // 1..128.
    PrefixLength interface{}

    // Specify address range for allocation.
    AddressRanges AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_AddressRanges

    // Exclude IPv6 addresses / prefixes.
    Excludes AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Excludes

    // Specify utilization mark.
    UtilizationMark AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_UtilizationMark

    // Specify prefix range for allocation.
    PrefixRanges AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_PrefixRanges

    // Specify network for allocation.
    Networks AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Networks
}

func (pool *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool) GetFilter() yfilter.YFilter { return pool.YFilter }

func (pool *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool) SetFilter(yf yfilter.YFilter) { pool.YFilter = yf }

func (pool *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool) GetGoName(yname string) string {
    if yname == "ipv6-pool-name" { return "Ipv6PoolName" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "address-ranges" { return "AddressRanges" }
    if yname == "excludes" { return "Excludes" }
    if yname == "utilization-mark" { return "UtilizationMark" }
    if yname == "prefix-ranges" { return "PrefixRanges" }
    if yname == "networks" { return "Networks" }
    return ""
}

func (pool *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool) GetSegmentPath() string {
    return "pool" + "[ipv6-pool-name='" + fmt.Sprintf("%v", pool.Ipv6PoolName) + "']"
}

func (pool *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "address-ranges" {
        return &pool.AddressRanges
    }
    if childYangName == "excludes" {
        return &pool.Excludes
    }
    if childYangName == "utilization-mark" {
        return &pool.UtilizationMark
    }
    if childYangName == "prefix-ranges" {
        return &pool.PrefixRanges
    }
    if childYangName == "networks" {
        return &pool.Networks
    }
    return nil
}

func (pool *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["address-ranges"] = &pool.AddressRanges
    children["excludes"] = &pool.Excludes
    children["utilization-mark"] = &pool.UtilizationMark
    children["prefix-ranges"] = &pool.PrefixRanges
    children["networks"] = &pool.Networks
    return children
}

func (pool *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv6-pool-name"] = pool.Ipv6PoolName
    leafs["prefix-length"] = pool.PrefixLength
    return leafs
}

func (pool *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool) GetBundleName() string { return "cisco_ios_xr" }

func (pool *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool) GetYangName() string { return "pool" }

func (pool *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pool *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pool *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pool *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool) SetParent(parent types.Entity) { pool.parent = parent }

func (pool *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool) GetParent() types.Entity { return pool.parent }

func (pool *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool) GetParentYangName() string { return "pools" }

// AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_AddressRanges
// Specify address range for allocation
type AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_AddressRanges struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // None. The type is slice of
    // AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_AddressRanges_AddressRange.
    AddressRange []AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_AddressRanges_AddressRange
}

func (addressRanges *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_AddressRanges) GetFilter() yfilter.YFilter { return addressRanges.YFilter }

func (addressRanges *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_AddressRanges) SetFilter(yf yfilter.YFilter) { addressRanges.YFilter = yf }

func (addressRanges *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_AddressRanges) GetGoName(yname string) string {
    if yname == "address-range" { return "AddressRange" }
    return ""
}

func (addressRanges *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_AddressRanges) GetSegmentPath() string {
    return "address-ranges"
}

func (addressRanges *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_AddressRanges) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "address-range" {
        for _, c := range addressRanges.AddressRange {
            if addressRanges.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_AddressRanges_AddressRange{}
        addressRanges.AddressRange = append(addressRanges.AddressRange, child)
        return &addressRanges.AddressRange[len(addressRanges.AddressRange)-1]
    }
    return nil
}

func (addressRanges *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_AddressRanges) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range addressRanges.AddressRange {
        children[addressRanges.AddressRange[i].GetSegmentPath()] = &addressRanges.AddressRange[i]
    }
    return children
}

func (addressRanges *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_AddressRanges) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (addressRanges *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_AddressRanges) GetBundleName() string { return "cisco_ios_xr" }

func (addressRanges *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_AddressRanges) GetYangName() string { return "address-ranges" }

func (addressRanges *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_AddressRanges) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (addressRanges *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_AddressRanges) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (addressRanges *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_AddressRanges) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (addressRanges *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_AddressRanges) SetParent(parent types.Entity) { addressRanges.parent = parent }

func (addressRanges *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_AddressRanges) GetParent() types.Entity { return addressRanges.parent }

func (addressRanges *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_AddressRanges) GetParentYangName() string { return "pool" }

// AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_AddressRanges_AddressRange
// None
type AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_AddressRanges_AddressRange struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Start address of the range. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    StartAddress interface{}

    // Blocked flag. The type is interface{} with range: -2147483648..2147483647.
    Blocked interface{}

    // End Address of the range. The type is one of the following types: string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?
    // This attribute is mandatory., or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?
    // This attribute is mandatory..
    EndAddress interface{}
}

func (addressRange *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_AddressRanges_AddressRange) GetFilter() yfilter.YFilter { return addressRange.YFilter }

func (addressRange *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_AddressRanges_AddressRange) SetFilter(yf yfilter.YFilter) { addressRange.YFilter = yf }

func (addressRange *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_AddressRanges_AddressRange) GetGoName(yname string) string {
    if yname == "start-address" { return "StartAddress" }
    if yname == "blocked" { return "Blocked" }
    if yname == "end-address" { return "EndAddress" }
    return ""
}

func (addressRange *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_AddressRanges_AddressRange) GetSegmentPath() string {
    return "address-range" + "[start-address='" + fmt.Sprintf("%v", addressRange.StartAddress) + "']"
}

func (addressRange *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_AddressRanges_AddressRange) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (addressRange *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_AddressRanges_AddressRange) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (addressRange *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_AddressRanges_AddressRange) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["start-address"] = addressRange.StartAddress
    leafs["blocked"] = addressRange.Blocked
    leafs["end-address"] = addressRange.EndAddress
    return leafs
}

func (addressRange *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_AddressRanges_AddressRange) GetBundleName() string { return "cisco_ios_xr" }

func (addressRange *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_AddressRanges_AddressRange) GetYangName() string { return "address-range" }

func (addressRange *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_AddressRanges_AddressRange) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (addressRange *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_AddressRanges_AddressRange) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (addressRange *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_AddressRanges_AddressRange) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (addressRange *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_AddressRanges_AddressRange) SetParent(parent types.Entity) { addressRange.parent = parent }

func (addressRange *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_AddressRanges_AddressRange) GetParent() types.Entity { return addressRange.parent }

func (addressRange *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_AddressRanges_AddressRange) GetParentYangName() string { return "address-ranges" }

// AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Excludes
// Exclude IPv6 addresses / prefixes
type AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Excludes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // None. The type is slice of
    // AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Excludes_Exclude.
    Exclude []AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Excludes_Exclude
}

func (excludes *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Excludes) GetFilter() yfilter.YFilter { return excludes.YFilter }

func (excludes *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Excludes) SetFilter(yf yfilter.YFilter) { excludes.YFilter = yf }

func (excludes *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Excludes) GetGoName(yname string) string {
    if yname == "exclude" { return "Exclude" }
    return ""
}

func (excludes *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Excludes) GetSegmentPath() string {
    return "excludes"
}

func (excludes *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Excludes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "exclude" {
        for _, c := range excludes.Exclude {
            if excludes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Excludes_Exclude{}
        excludes.Exclude = append(excludes.Exclude, child)
        return &excludes.Exclude[len(excludes.Exclude)-1]
    }
    return nil
}

func (excludes *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Excludes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range excludes.Exclude {
        children[excludes.Exclude[i].GetSegmentPath()] = &excludes.Exclude[i]
    }
    return children
}

func (excludes *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Excludes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (excludes *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Excludes) GetBundleName() string { return "cisco_ios_xr" }

func (excludes *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Excludes) GetYangName() string { return "excludes" }

func (excludes *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Excludes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (excludes *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Excludes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (excludes *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Excludes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (excludes *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Excludes) SetParent(parent types.Entity) { excludes.parent = parent }

func (excludes *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Excludes) GetParent() types.Entity { return excludes.parent }

func (excludes *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Excludes) GetParentYangName() string { return "pool" }

// AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Excludes_Exclude
// None
type AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Excludes_Exclude struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. First Address in IPv6 exclude range. The type is
    // one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    StartAddress interface{}

    // Last address in exclude Range. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?
    // This attribute is mandatory., or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?
    // This attribute is mandatory..
    EndAddress interface{}
}

func (exclude *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Excludes_Exclude) GetFilter() yfilter.YFilter { return exclude.YFilter }

func (exclude *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Excludes_Exclude) SetFilter(yf yfilter.YFilter) { exclude.YFilter = yf }

func (exclude *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Excludes_Exclude) GetGoName(yname string) string {
    if yname == "start-address" { return "StartAddress" }
    if yname == "end-address" { return "EndAddress" }
    return ""
}

func (exclude *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Excludes_Exclude) GetSegmentPath() string {
    return "exclude" + "[start-address='" + fmt.Sprintf("%v", exclude.StartAddress) + "']"
}

func (exclude *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Excludes_Exclude) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (exclude *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Excludes_Exclude) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (exclude *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Excludes_Exclude) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["start-address"] = exclude.StartAddress
    leafs["end-address"] = exclude.EndAddress
    return leafs
}

func (exclude *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Excludes_Exclude) GetBundleName() string { return "cisco_ios_xr" }

func (exclude *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Excludes_Exclude) GetYangName() string { return "exclude" }

func (exclude *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Excludes_Exclude) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (exclude *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Excludes_Exclude) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (exclude *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Excludes_Exclude) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (exclude *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Excludes_Exclude) SetParent(parent types.Entity) { exclude.parent = parent }

func (exclude *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Excludes_Exclude) GetParent() types.Entity { return exclude.parent }

func (exclude *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Excludes_Exclude) GetParentYangName() string { return "excludes" }

// AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_UtilizationMark
// Specify utilization mark
type AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_UtilizationMark struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify numerical value as percentage. The type is interface{} with range:
    // 0..100. Units are percentage.
    HighMark interface{}

    // Specify numerical value as percentage. The type is interface{} with range:
    // 0..100. Units are percentage.
    LowMark interface{}
}

func (utilizationMark *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_UtilizationMark) GetFilter() yfilter.YFilter { return utilizationMark.YFilter }

func (utilizationMark *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_UtilizationMark) SetFilter(yf yfilter.YFilter) { utilizationMark.YFilter = yf }

func (utilizationMark *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_UtilizationMark) GetGoName(yname string) string {
    if yname == "high-mark" { return "HighMark" }
    if yname == "low-mark" { return "LowMark" }
    return ""
}

func (utilizationMark *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_UtilizationMark) GetSegmentPath() string {
    return "utilization-mark"
}

func (utilizationMark *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_UtilizationMark) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (utilizationMark *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_UtilizationMark) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (utilizationMark *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_UtilizationMark) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["high-mark"] = utilizationMark.HighMark
    leafs["low-mark"] = utilizationMark.LowMark
    return leafs
}

func (utilizationMark *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_UtilizationMark) GetBundleName() string { return "cisco_ios_xr" }

func (utilizationMark *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_UtilizationMark) GetYangName() string { return "utilization-mark" }

func (utilizationMark *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_UtilizationMark) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (utilizationMark *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_UtilizationMark) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (utilizationMark *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_UtilizationMark) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (utilizationMark *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_UtilizationMark) SetParent(parent types.Entity) { utilizationMark.parent = parent }

func (utilizationMark *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_UtilizationMark) GetParent() types.Entity { return utilizationMark.parent }

func (utilizationMark *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_UtilizationMark) GetParentYangName() string { return "pool" }

// AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_PrefixRanges
// Specify prefix range for allocation
type AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_PrefixRanges struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // None. The type is slice of
    // AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_PrefixRanges_PrefixRange.
    PrefixRange []AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_PrefixRanges_PrefixRange
}

func (prefixRanges *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_PrefixRanges) GetFilter() yfilter.YFilter { return prefixRanges.YFilter }

func (prefixRanges *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_PrefixRanges) SetFilter(yf yfilter.YFilter) { prefixRanges.YFilter = yf }

func (prefixRanges *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_PrefixRanges) GetGoName(yname string) string {
    if yname == "prefix-range" { return "PrefixRange" }
    return ""
}

func (prefixRanges *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_PrefixRanges) GetSegmentPath() string {
    return "prefix-ranges"
}

func (prefixRanges *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_PrefixRanges) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-range" {
        for _, c := range prefixRanges.PrefixRange {
            if prefixRanges.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_PrefixRanges_PrefixRange{}
        prefixRanges.PrefixRange = append(prefixRanges.PrefixRange, child)
        return &prefixRanges.PrefixRange[len(prefixRanges.PrefixRange)-1]
    }
    return nil
}

func (prefixRanges *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_PrefixRanges) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range prefixRanges.PrefixRange {
        children[prefixRanges.PrefixRange[i].GetSegmentPath()] = &prefixRanges.PrefixRange[i]
    }
    return children
}

func (prefixRanges *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_PrefixRanges) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (prefixRanges *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_PrefixRanges) GetBundleName() string { return "cisco_ios_xr" }

func (prefixRanges *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_PrefixRanges) GetYangName() string { return "prefix-ranges" }

func (prefixRanges *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_PrefixRanges) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefixRanges *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_PrefixRanges) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefixRanges *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_PrefixRanges) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefixRanges *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_PrefixRanges) SetParent(parent types.Entity) { prefixRanges.parent = parent }

func (prefixRanges *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_PrefixRanges) GetParent() types.Entity { return prefixRanges.parent }

func (prefixRanges *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_PrefixRanges) GetParentYangName() string { return "pool" }

// AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_PrefixRanges_PrefixRange
// None
type AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_PrefixRanges_PrefixRange struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. First prefix of range. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    StartPrefix interface{}

    // Blocked flag. The type is interface{} with range: -2147483648..2147483647.
    Blocked interface{}

    // Last prefix of range. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?
    // This attribute is mandatory., or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?
    // This attribute is mandatory..
    EndPrefix interface{}
}

func (prefixRange *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_PrefixRanges_PrefixRange) GetFilter() yfilter.YFilter { return prefixRange.YFilter }

func (prefixRange *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_PrefixRanges_PrefixRange) SetFilter(yf yfilter.YFilter) { prefixRange.YFilter = yf }

func (prefixRange *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_PrefixRanges_PrefixRange) GetGoName(yname string) string {
    if yname == "start-prefix" { return "StartPrefix" }
    if yname == "blocked" { return "Blocked" }
    if yname == "end-prefix" { return "EndPrefix" }
    return ""
}

func (prefixRange *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_PrefixRanges_PrefixRange) GetSegmentPath() string {
    return "prefix-range" + "[start-prefix='" + fmt.Sprintf("%v", prefixRange.StartPrefix) + "']"
}

func (prefixRange *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_PrefixRanges_PrefixRange) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (prefixRange *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_PrefixRanges_PrefixRange) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (prefixRange *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_PrefixRanges_PrefixRange) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["start-prefix"] = prefixRange.StartPrefix
    leafs["blocked"] = prefixRange.Blocked
    leafs["end-prefix"] = prefixRange.EndPrefix
    return leafs
}

func (prefixRange *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_PrefixRanges_PrefixRange) GetBundleName() string { return "cisco_ios_xr" }

func (prefixRange *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_PrefixRanges_PrefixRange) GetYangName() string { return "prefix-range" }

func (prefixRange *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_PrefixRanges_PrefixRange) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefixRange *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_PrefixRanges_PrefixRange) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefixRange *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_PrefixRanges_PrefixRange) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefixRange *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_PrefixRanges_PrefixRange) SetParent(parent types.Entity) { prefixRange.parent = parent }

func (prefixRange *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_PrefixRanges_PrefixRange) GetParent() types.Entity { return prefixRange.parent }

func (prefixRange *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_PrefixRanges_PrefixRange) GetParentYangName() string { return "prefix-ranges" }

// AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Networks
// Specify network for allocation
type AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Networks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // None. The type is slice of
    // AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Networks_Network.
    Network []AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Networks_Network
}

func (networks *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Networks) GetFilter() yfilter.YFilter { return networks.YFilter }

func (networks *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Networks) SetFilter(yf yfilter.YFilter) { networks.YFilter = yf }

func (networks *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Networks) GetGoName(yname string) string {
    if yname == "network" { return "Network" }
    return ""
}

func (networks *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Networks) GetSegmentPath() string {
    return "networks"
}

func (networks *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Networks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "network" {
        for _, c := range networks.Network {
            if networks.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Networks_Network{}
        networks.Network = append(networks.Network, child)
        return &networks.Network[len(networks.Network)-1]
    }
    return nil
}

func (networks *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Networks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range networks.Network {
        children[networks.Network[i].GetSegmentPath()] = &networks.Network[i]
    }
    return children
}

func (networks *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Networks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (networks *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Networks) GetBundleName() string { return "cisco_ios_xr" }

func (networks *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Networks) GetYangName() string { return "networks" }

func (networks *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Networks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (networks *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Networks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (networks *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Networks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (networks *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Networks) SetParent(parent types.Entity) { networks.parent = parent }

func (networks *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Networks) GetParent() types.Entity { return networks.parent }

func (networks *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Networks) GetParentYangName() string { return "pool" }

// AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Networks_Network
// None
type AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Networks_Network struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. None. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Prefix interface{}

    // Blocked flag. The type is interface{} with range: -2147483648..2147483647.
    Blocked interface{}

    // Prefix length for the IPv6 Prefix. The type is interface{} with range:
    // 1..128. This attribute is mandatory.
    PrefixLength interface{}
}

func (network *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Networks_Network) GetFilter() yfilter.YFilter { return network.YFilter }

func (network *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Networks_Network) SetFilter(yf yfilter.YFilter) { network.YFilter = yf }

func (network *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Networks_Network) GetGoName(yname string) string {
    if yname == "prefix" { return "Prefix" }
    if yname == "blocked" { return "Blocked" }
    if yname == "prefix-length" { return "PrefixLength" }
    return ""
}

func (network *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Networks_Network) GetSegmentPath() string {
    return "network" + "[prefix='" + fmt.Sprintf("%v", network.Prefix) + "']"
}

func (network *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Networks_Network) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (network *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Networks_Network) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (network *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Networks_Network) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix"] = network.Prefix
    leafs["blocked"] = network.Blocked
    leafs["prefix-length"] = network.PrefixLength
    return leafs
}

func (network *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Networks_Network) GetBundleName() string { return "cisco_ios_xr" }

func (network *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Networks_Network) GetYangName() string { return "network" }

func (network *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Networks_Network) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (network *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Networks_Network) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (network *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Networks_Network) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (network *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Networks_Network) SetParent(parent types.Entity) { network.parent = parent }

func (network *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Networks_Network) GetParent() types.Entity { return network.parent }

func (network *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Networks_Network) GetParentYangName() string { return "networks" }

// AddressPoolService_Vrfs_Vrf_Ipv4
// Enter IPv4 specific configuration
type AddressPoolService_Vrfs_Vrf_Ipv4 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 Pool Table.
    Pools AddressPoolService_Vrfs_Vrf_Ipv4_Pools
}

func (ipv4 *AddressPoolService_Vrfs_Vrf_Ipv4) GetFilter() yfilter.YFilter { return ipv4.YFilter }

func (ipv4 *AddressPoolService_Vrfs_Vrf_Ipv4) SetFilter(yf yfilter.YFilter) { ipv4.YFilter = yf }

func (ipv4 *AddressPoolService_Vrfs_Vrf_Ipv4) GetGoName(yname string) string {
    if yname == "pools" { return "Pools" }
    return ""
}

func (ipv4 *AddressPoolService_Vrfs_Vrf_Ipv4) GetSegmentPath() string {
    return "ipv4"
}

func (ipv4 *AddressPoolService_Vrfs_Vrf_Ipv4) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "pools" {
        return &ipv4.Pools
    }
    return nil
}

func (ipv4 *AddressPoolService_Vrfs_Vrf_Ipv4) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["pools"] = &ipv4.Pools
    return children
}

func (ipv4 *AddressPoolService_Vrfs_Vrf_Ipv4) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv4 *AddressPoolService_Vrfs_Vrf_Ipv4) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4 *AddressPoolService_Vrfs_Vrf_Ipv4) GetYangName() string { return "ipv4" }

func (ipv4 *AddressPoolService_Vrfs_Vrf_Ipv4) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4 *AddressPoolService_Vrfs_Vrf_Ipv4) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4 *AddressPoolService_Vrfs_Vrf_Ipv4) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4 *AddressPoolService_Vrfs_Vrf_Ipv4) SetParent(parent types.Entity) { ipv4.parent = parent }

func (ipv4 *AddressPoolService_Vrfs_Vrf_Ipv4) GetParent() types.Entity { return ipv4.parent }

func (ipv4 *AddressPoolService_Vrfs_Vrf_Ipv4) GetParentYangName() string { return "vrf" }

// AddressPoolService_Vrfs_Vrf_Ipv4_Pools
// IPv4 Pool Table
type AddressPoolService_Vrfs_Vrf_Ipv4_Pools struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 Pool name. The type is slice of
    // AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool.
    Pool []AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool
}

func (pools *AddressPoolService_Vrfs_Vrf_Ipv4_Pools) GetFilter() yfilter.YFilter { return pools.YFilter }

func (pools *AddressPoolService_Vrfs_Vrf_Ipv4_Pools) SetFilter(yf yfilter.YFilter) { pools.YFilter = yf }

func (pools *AddressPoolService_Vrfs_Vrf_Ipv4_Pools) GetGoName(yname string) string {
    if yname == "pool" { return "Pool" }
    return ""
}

func (pools *AddressPoolService_Vrfs_Vrf_Ipv4_Pools) GetSegmentPath() string {
    return "pools"
}

func (pools *AddressPoolService_Vrfs_Vrf_Ipv4_Pools) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "pool" {
        for _, c := range pools.Pool {
            if pools.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool{}
        pools.Pool = append(pools.Pool, child)
        return &pools.Pool[len(pools.Pool)-1]
    }
    return nil
}

func (pools *AddressPoolService_Vrfs_Vrf_Ipv4_Pools) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range pools.Pool {
        children[pools.Pool[i].GetSegmentPath()] = &pools.Pool[i]
    }
    return children
}

func (pools *AddressPoolService_Vrfs_Vrf_Ipv4_Pools) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (pools *AddressPoolService_Vrfs_Vrf_Ipv4_Pools) GetBundleName() string { return "cisco_ios_xr" }

func (pools *AddressPoolService_Vrfs_Vrf_Ipv4_Pools) GetYangName() string { return "pools" }

func (pools *AddressPoolService_Vrfs_Vrf_Ipv4_Pools) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pools *AddressPoolService_Vrfs_Vrf_Ipv4_Pools) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pools *AddressPoolService_Vrfs_Vrf_Ipv4_Pools) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pools *AddressPoolService_Vrfs_Vrf_Ipv4_Pools) SetParent(parent types.Entity) { pools.parent = parent }

func (pools *AddressPoolService_Vrfs_Vrf_Ipv4_Pools) GetParent() types.Entity { return pools.parent }

func (pools *AddressPoolService_Vrfs_Vrf_Ipv4_Pools) GetParentYangName() string { return "ipv4" }

// AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool
// IPv4 Pool name
type AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Enter the IPv4 Pool name. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    PoolName interface{}

    // address range for allocation.
    AddressRanges AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_AddressRanges

    // Exclude addresses.
    Excludes AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Excludes

    // Specify utilization mark.
    UtilizationMark AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_UtilizationMark

    // Specify network for allocation.
    Networks AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Networks
}

func (pool *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool) GetFilter() yfilter.YFilter { return pool.YFilter }

func (pool *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool) SetFilter(yf yfilter.YFilter) { pool.YFilter = yf }

func (pool *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool) GetGoName(yname string) string {
    if yname == "pool-name" { return "PoolName" }
    if yname == "address-ranges" { return "AddressRanges" }
    if yname == "excludes" { return "Excludes" }
    if yname == "utilization-mark" { return "UtilizationMark" }
    if yname == "networks" { return "Networks" }
    return ""
}

func (pool *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool) GetSegmentPath() string {
    return "pool" + "[pool-name='" + fmt.Sprintf("%v", pool.PoolName) + "']"
}

func (pool *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "address-ranges" {
        return &pool.AddressRanges
    }
    if childYangName == "excludes" {
        return &pool.Excludes
    }
    if childYangName == "utilization-mark" {
        return &pool.UtilizationMark
    }
    if childYangName == "networks" {
        return &pool.Networks
    }
    return nil
}

func (pool *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["address-ranges"] = &pool.AddressRanges
    children["excludes"] = &pool.Excludes
    children["utilization-mark"] = &pool.UtilizationMark
    children["networks"] = &pool.Networks
    return children
}

func (pool *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pool-name"] = pool.PoolName
    return leafs
}

func (pool *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool) GetBundleName() string { return "cisco_ios_xr" }

func (pool *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool) GetYangName() string { return "pool" }

func (pool *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pool *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pool *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pool *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool) SetParent(parent types.Entity) { pool.parent = parent }

func (pool *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool) GetParent() types.Entity { return pool.parent }

func (pool *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool) GetParentYangName() string { return "pools" }

// AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_AddressRanges
// address range for allocation
type AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_AddressRanges struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify first address in range. The type is slice of
    // AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_AddressRanges_AddressRange.
    AddressRange []AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_AddressRanges_AddressRange
}

func (addressRanges *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_AddressRanges) GetFilter() yfilter.YFilter { return addressRanges.YFilter }

func (addressRanges *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_AddressRanges) SetFilter(yf yfilter.YFilter) { addressRanges.YFilter = yf }

func (addressRanges *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_AddressRanges) GetGoName(yname string) string {
    if yname == "address-range" { return "AddressRange" }
    return ""
}

func (addressRanges *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_AddressRanges) GetSegmentPath() string {
    return "address-ranges"
}

func (addressRanges *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_AddressRanges) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "address-range" {
        for _, c := range addressRanges.AddressRange {
            if addressRanges.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_AddressRanges_AddressRange{}
        addressRanges.AddressRange = append(addressRanges.AddressRange, child)
        return &addressRanges.AddressRange[len(addressRanges.AddressRange)-1]
    }
    return nil
}

func (addressRanges *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_AddressRanges) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range addressRanges.AddressRange {
        children[addressRanges.AddressRange[i].GetSegmentPath()] = &addressRanges.AddressRange[i]
    }
    return children
}

func (addressRanges *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_AddressRanges) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (addressRanges *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_AddressRanges) GetBundleName() string { return "cisco_ios_xr" }

func (addressRanges *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_AddressRanges) GetYangName() string { return "address-ranges" }

func (addressRanges *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_AddressRanges) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (addressRanges *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_AddressRanges) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (addressRanges *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_AddressRanges) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (addressRanges *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_AddressRanges) SetParent(parent types.Entity) { addressRanges.parent = parent }

func (addressRanges *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_AddressRanges) GetParent() types.Entity { return addressRanges.parent }

func (addressRanges *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_AddressRanges) GetParentYangName() string { return "pool" }

// AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_AddressRanges_AddressRange
// Specify first address in range
type AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_AddressRanges_AddressRange struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Specify first address of the range. The type is
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    StartAddress interface{}

    // Last address of the range. The type is one of the following types: string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?
    // This attribute is mandatory., or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?
    // This attribute is mandatory..
    EndAddress interface{}

    // Blocked flag. The type is interface{} with range: -2147483648..2147483647.
    Blocked interface{}
}

func (addressRange *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_AddressRanges_AddressRange) GetFilter() yfilter.YFilter { return addressRange.YFilter }

func (addressRange *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_AddressRanges_AddressRange) SetFilter(yf yfilter.YFilter) { addressRange.YFilter = yf }

func (addressRange *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_AddressRanges_AddressRange) GetGoName(yname string) string {
    if yname == "start-address" { return "StartAddress" }
    if yname == "end-address" { return "EndAddress" }
    if yname == "blocked" { return "Blocked" }
    return ""
}

func (addressRange *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_AddressRanges_AddressRange) GetSegmentPath() string {
    return "address-range" + "[start-address='" + fmt.Sprintf("%v", addressRange.StartAddress) + "']"
}

func (addressRange *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_AddressRanges_AddressRange) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (addressRange *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_AddressRanges_AddressRange) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (addressRange *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_AddressRanges_AddressRange) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["start-address"] = addressRange.StartAddress
    leafs["end-address"] = addressRange.EndAddress
    leafs["blocked"] = addressRange.Blocked
    return leafs
}

func (addressRange *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_AddressRanges_AddressRange) GetBundleName() string { return "cisco_ios_xr" }

func (addressRange *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_AddressRanges_AddressRange) GetYangName() string { return "address-range" }

func (addressRange *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_AddressRanges_AddressRange) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (addressRange *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_AddressRanges_AddressRange) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (addressRange *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_AddressRanges_AddressRange) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (addressRange *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_AddressRanges_AddressRange) SetParent(parent types.Entity) { addressRange.parent = parent }

func (addressRange *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_AddressRanges_AddressRange) GetParent() types.Entity { return addressRange.parent }

func (addressRange *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_AddressRanges_AddressRange) GetParentYangName() string { return "address-ranges" }

// AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Excludes
// Exclude addresses
type AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Excludes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // First address in range. The type is slice of
    // AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Excludes_Exclude.
    Exclude []AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Excludes_Exclude
}

func (excludes *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Excludes) GetFilter() yfilter.YFilter { return excludes.YFilter }

func (excludes *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Excludes) SetFilter(yf yfilter.YFilter) { excludes.YFilter = yf }

func (excludes *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Excludes) GetGoName(yname string) string {
    if yname == "exclude" { return "Exclude" }
    return ""
}

func (excludes *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Excludes) GetSegmentPath() string {
    return "excludes"
}

func (excludes *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Excludes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "exclude" {
        for _, c := range excludes.Exclude {
            if excludes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Excludes_Exclude{}
        excludes.Exclude = append(excludes.Exclude, child)
        return &excludes.Exclude[len(excludes.Exclude)-1]
    }
    return nil
}

func (excludes *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Excludes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range excludes.Exclude {
        children[excludes.Exclude[i].GetSegmentPath()] = &excludes.Exclude[i]
    }
    return children
}

func (excludes *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Excludes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (excludes *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Excludes) GetBundleName() string { return "cisco_ios_xr" }

func (excludes *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Excludes) GetYangName() string { return "excludes" }

func (excludes *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Excludes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (excludes *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Excludes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (excludes *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Excludes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (excludes *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Excludes) SetParent(parent types.Entity) { excludes.parent = parent }

func (excludes *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Excludes) GetParent() types.Entity { return excludes.parent }

func (excludes *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Excludes) GetParentYangName() string { return "pool" }

// AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Excludes_Exclude
// First address in range
type AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Excludes_Exclude struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. First address in exclude range. The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    StartAddress interface{}

    // Last address in excluded range. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?
    // This attribute is mandatory., or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?
    // This attribute is mandatory..
    EndAddress interface{}
}

func (exclude *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Excludes_Exclude) GetFilter() yfilter.YFilter { return exclude.YFilter }

func (exclude *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Excludes_Exclude) SetFilter(yf yfilter.YFilter) { exclude.YFilter = yf }

func (exclude *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Excludes_Exclude) GetGoName(yname string) string {
    if yname == "start-address" { return "StartAddress" }
    if yname == "end-address" { return "EndAddress" }
    return ""
}

func (exclude *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Excludes_Exclude) GetSegmentPath() string {
    return "exclude" + "[start-address='" + fmt.Sprintf("%v", exclude.StartAddress) + "']"
}

func (exclude *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Excludes_Exclude) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (exclude *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Excludes_Exclude) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (exclude *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Excludes_Exclude) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["start-address"] = exclude.StartAddress
    leafs["end-address"] = exclude.EndAddress
    return leafs
}

func (exclude *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Excludes_Exclude) GetBundleName() string { return "cisco_ios_xr" }

func (exclude *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Excludes_Exclude) GetYangName() string { return "exclude" }

func (exclude *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Excludes_Exclude) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (exclude *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Excludes_Exclude) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (exclude *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Excludes_Exclude) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (exclude *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Excludes_Exclude) SetParent(parent types.Entity) { exclude.parent = parent }

func (exclude *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Excludes_Exclude) GetParent() types.Entity { return exclude.parent }

func (exclude *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Excludes_Exclude) GetParentYangName() string { return "excludes" }

// AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_UtilizationMark
// Specify utilization mark
type AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_UtilizationMark struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify numerical value as percentage. The type is interface{} with range:
    // 0..100. Units are percentage.
    High interface{}

    // Specify numerical value as percentage. The type is interface{} with range:
    // 0..100. Units are percentage.
    Low interface{}
}

func (utilizationMark *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_UtilizationMark) GetFilter() yfilter.YFilter { return utilizationMark.YFilter }

func (utilizationMark *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_UtilizationMark) SetFilter(yf yfilter.YFilter) { utilizationMark.YFilter = yf }

func (utilizationMark *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_UtilizationMark) GetGoName(yname string) string {
    if yname == "high" { return "High" }
    if yname == "low" { return "Low" }
    return ""
}

func (utilizationMark *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_UtilizationMark) GetSegmentPath() string {
    return "utilization-mark"
}

func (utilizationMark *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_UtilizationMark) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (utilizationMark *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_UtilizationMark) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (utilizationMark *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_UtilizationMark) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["high"] = utilizationMark.High
    leafs["low"] = utilizationMark.Low
    return leafs
}

func (utilizationMark *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_UtilizationMark) GetBundleName() string { return "cisco_ios_xr" }

func (utilizationMark *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_UtilizationMark) GetYangName() string { return "utilization-mark" }

func (utilizationMark *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_UtilizationMark) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (utilizationMark *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_UtilizationMark) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (utilizationMark *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_UtilizationMark) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (utilizationMark *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_UtilizationMark) SetParent(parent types.Entity) { utilizationMark.parent = parent }

func (utilizationMark *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_UtilizationMark) GetParent() types.Entity { return utilizationMark.parent }

func (utilizationMark *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_UtilizationMark) GetParentYangName() string { return "pool" }

// AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Networks
// Specify network for allocation
type AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Networks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Network Prefix. The type is slice of
    // AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Networks_Network.
    Network []AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Networks_Network
}

func (networks *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Networks) GetFilter() yfilter.YFilter { return networks.YFilter }

func (networks *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Networks) SetFilter(yf yfilter.YFilter) { networks.YFilter = yf }

func (networks *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Networks) GetGoName(yname string) string {
    if yname == "network" { return "Network" }
    return ""
}

func (networks *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Networks) GetSegmentPath() string {
    return "networks"
}

func (networks *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Networks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "network" {
        for _, c := range networks.Network {
            if networks.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Networks_Network{}
        networks.Network = append(networks.Network, child)
        return &networks.Network[len(networks.Network)-1]
    }
    return nil
}

func (networks *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Networks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range networks.Network {
        children[networks.Network[i].GetSegmentPath()] = &networks.Network[i]
    }
    return children
}

func (networks *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Networks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (networks *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Networks) GetBundleName() string { return "cisco_ios_xr" }

func (networks *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Networks) GetYangName() string { return "networks" }

func (networks *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Networks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (networks *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Networks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (networks *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Networks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (networks *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Networks) SetParent(parent types.Entity) { networks.parent = parent }

func (networks *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Networks) GetParent() types.Entity { return networks.parent }

func (networks *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Networks) GetParentYangName() string { return "pool" }

// AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Networks_Network
// Network Prefix
type AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Networks_Network struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. None. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv4Prefix interface{}

    // Blocked flag. The type is interface{} with range: -2147483648..2147483647.
    Blocked interface{}

    // Subnet Length for IPv4 subnet. The type is interface{} with range:
    // -2147483648..2147483647. This attribute is mandatory.
    PrefixLength interface{}

    // Default Gateway for IPv4 subnet. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DefaultRouter interface{}
}

func (network *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Networks_Network) GetFilter() yfilter.YFilter { return network.YFilter }

func (network *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Networks_Network) SetFilter(yf yfilter.YFilter) { network.YFilter = yf }

func (network *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Networks_Network) GetGoName(yname string) string {
    if yname == "ipv4-prefix" { return "Ipv4Prefix" }
    if yname == "blocked" { return "Blocked" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "default-router" { return "DefaultRouter" }
    return ""
}

func (network *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Networks_Network) GetSegmentPath() string {
    return "network" + "[ipv4-prefix='" + fmt.Sprintf("%v", network.Ipv4Prefix) + "']"
}

func (network *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Networks_Network) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (network *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Networks_Network) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (network *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Networks_Network) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv4-prefix"] = network.Ipv4Prefix
    leafs["blocked"] = network.Blocked
    leafs["prefix-length"] = network.PrefixLength
    leafs["default-router"] = network.DefaultRouter
    return leafs
}

func (network *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Networks_Network) GetBundleName() string { return "cisco_ios_xr" }

func (network *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Networks_Network) GetYangName() string { return "network" }

func (network *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Networks_Network) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (network *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Networks_Network) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (network *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Networks_Network) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (network *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Networks_Network) SetParent(parent types.Entity) { network.parent = parent }

func (network *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Networks_Network) GetParent() types.Entity { return network.parent }

func (network *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Networks_Network) GetParentYangName() string { return "networks" }

