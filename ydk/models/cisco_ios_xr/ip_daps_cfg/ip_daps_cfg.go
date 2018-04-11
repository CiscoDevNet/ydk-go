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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enter VRF specific config mode.
    Vrfs AddressPoolService_Vrfs
}

func (addressPoolService *AddressPoolService) GetEntityData() *types.CommonEntityData {
    addressPoolService.EntityData.YFilter = addressPoolService.YFilter
    addressPoolService.EntityData.YangName = "address-pool-service"
    addressPoolService.EntityData.BundleName = "cisco_ios_xr"
    addressPoolService.EntityData.ParentYangName = "Cisco-IOS-XR-ip-daps-cfg"
    addressPoolService.EntityData.SegmentPath = "Cisco-IOS-XR-ip-daps-cfg:address-pool-service"
    addressPoolService.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressPoolService.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressPoolService.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressPoolService.EntityData.Children = make(map[string]types.YChild)
    addressPoolService.EntityData.Children["vrfs"] = types.YChild{"Vrfs", &addressPoolService.Vrfs}
    addressPoolService.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(addressPoolService.EntityData)
}

// AddressPoolService_Vrfs
// Enter VRF specific config mode
type AddressPoolService_Vrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify VRF Name. The type is slice of AddressPoolService_Vrfs_Vrf.
    Vrf []AddressPoolService_Vrfs_Vrf
}

func (vrfs *AddressPoolService_Vrfs) GetEntityData() *types.CommonEntityData {
    vrfs.EntityData.YFilter = vrfs.YFilter
    vrfs.EntityData.YangName = "vrfs"
    vrfs.EntityData.BundleName = "cisco_ios_xr"
    vrfs.EntityData.ParentYangName = "address-pool-service"
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

// AddressPoolService_Vrfs_Vrf
// Specify VRF Name
type AddressPoolService_Vrfs_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. none. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    VrfName interface{}

    // Enter IPv6 specific mode.
    Ipv6 AddressPoolService_Vrfs_Vrf_Ipv6

    // Enter IPv4 specific configuration.
    Ipv4 AddressPoolService_Vrfs_Vrf_Ipv4
}

func (vrf *AddressPoolService_Vrfs_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "vrfs"
    vrf.EntityData.SegmentPath = "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = make(map[string]types.YChild)
    vrf.EntityData.Children["ipv6"] = types.YChild{"Ipv6", &vrf.Ipv6}
    vrf.EntityData.Children["ipv4"] = types.YChild{"Ipv4", &vrf.Ipv4}
    vrf.EntityData.Leafs = make(map[string]types.YLeaf)
    vrf.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", vrf.VrfName}
    return &(vrf.EntityData)
}

// AddressPoolService_Vrfs_Vrf_Ipv6
// Enter IPv6 specific mode
type AddressPoolService_Vrfs_Vrf_Ipv6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 Pool Name.
    Pools AddressPoolService_Vrfs_Vrf_Ipv6_Pools
}

func (ipv6 *AddressPoolService_Vrfs_Vrf_Ipv6) GetEntityData() *types.CommonEntityData {
    ipv6.EntityData.YFilter = ipv6.YFilter
    ipv6.EntityData.YangName = "ipv6"
    ipv6.EntityData.BundleName = "cisco_ios_xr"
    ipv6.EntityData.ParentYangName = "vrf"
    ipv6.EntityData.SegmentPath = "ipv6"
    ipv6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6.EntityData.Children = make(map[string]types.YChild)
    ipv6.EntityData.Children["pools"] = types.YChild{"Pools", &ipv6.Pools}
    ipv6.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipv6.EntityData)
}

// AddressPoolService_Vrfs_Vrf_Ipv6_Pools
// IPv6 Pool Name
type AddressPoolService_Vrfs_Vrf_Ipv6_Pools struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enter the IPv6 Pool name. The type is slice of
    // AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool.
    Pool []AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool
}

func (pools *AddressPoolService_Vrfs_Vrf_Ipv6_Pools) GetEntityData() *types.CommonEntityData {
    pools.EntityData.YFilter = pools.YFilter
    pools.EntityData.YangName = "pools"
    pools.EntityData.BundleName = "cisco_ios_xr"
    pools.EntityData.ParentYangName = "ipv6"
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

// AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool
// Enter the IPv6 Pool name
type AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Enter the IPv6 Pool name. The type is string with
    // pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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

func (pool *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool) GetEntityData() *types.CommonEntityData {
    pool.EntityData.YFilter = pool.YFilter
    pool.EntityData.YangName = "pool"
    pool.EntityData.BundleName = "cisco_ios_xr"
    pool.EntityData.ParentYangName = "pools"
    pool.EntityData.SegmentPath = "pool" + "[ipv6-pool-name='" + fmt.Sprintf("%v", pool.Ipv6PoolName) + "']"
    pool.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pool.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pool.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pool.EntityData.Children = make(map[string]types.YChild)
    pool.EntityData.Children["address-ranges"] = types.YChild{"AddressRanges", &pool.AddressRanges}
    pool.EntityData.Children["excludes"] = types.YChild{"Excludes", &pool.Excludes}
    pool.EntityData.Children["utilization-mark"] = types.YChild{"UtilizationMark", &pool.UtilizationMark}
    pool.EntityData.Children["prefix-ranges"] = types.YChild{"PrefixRanges", &pool.PrefixRanges}
    pool.EntityData.Children["networks"] = types.YChild{"Networks", &pool.Networks}
    pool.EntityData.Leafs = make(map[string]types.YLeaf)
    pool.EntityData.Leafs["ipv6-pool-name"] = types.YLeaf{"Ipv6PoolName", pool.Ipv6PoolName}
    pool.EntityData.Leafs["prefix-length"] = types.YLeaf{"PrefixLength", pool.PrefixLength}
    return &(pool.EntityData)
}

// AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_AddressRanges
// Specify address range for allocation
type AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_AddressRanges struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // None. The type is slice of
    // AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_AddressRanges_AddressRange.
    AddressRange []AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_AddressRanges_AddressRange
}

func (addressRanges *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_AddressRanges) GetEntityData() *types.CommonEntityData {
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

// AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_AddressRanges_AddressRange
// None
type AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_AddressRanges_AddressRange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Start address of the range. The type is one of the
    // following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    StartAddress interface{}

    // Blocked flag. The type is interface{} with range: -2147483648..2147483647.
    Blocked interface{}

    // End Address of the range. The type is one of the following types: string
    // with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'
    // This attribute is mandatory., or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'
    // This attribute is mandatory..
    EndAddress interface{}
}

func (addressRange *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_AddressRanges_AddressRange) GetEntityData() *types.CommonEntityData {
    addressRange.EntityData.YFilter = addressRange.YFilter
    addressRange.EntityData.YangName = "address-range"
    addressRange.EntityData.BundleName = "cisco_ios_xr"
    addressRange.EntityData.ParentYangName = "address-ranges"
    addressRange.EntityData.SegmentPath = "address-range" + "[start-address='" + fmt.Sprintf("%v", addressRange.StartAddress) + "']"
    addressRange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressRange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressRange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressRange.EntityData.Children = make(map[string]types.YChild)
    addressRange.EntityData.Leafs = make(map[string]types.YLeaf)
    addressRange.EntityData.Leafs["start-address"] = types.YLeaf{"StartAddress", addressRange.StartAddress}
    addressRange.EntityData.Leafs["blocked"] = types.YLeaf{"Blocked", addressRange.Blocked}
    addressRange.EntityData.Leafs["end-address"] = types.YLeaf{"EndAddress", addressRange.EndAddress}
    return &(addressRange.EntityData)
}

// AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Excludes
// Exclude IPv6 addresses / prefixes
type AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Excludes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // None. The type is slice of
    // AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Excludes_Exclude.
    Exclude []AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Excludes_Exclude
}

func (excludes *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Excludes) GetEntityData() *types.CommonEntityData {
    excludes.EntityData.YFilter = excludes.YFilter
    excludes.EntityData.YangName = "excludes"
    excludes.EntityData.BundleName = "cisco_ios_xr"
    excludes.EntityData.ParentYangName = "pool"
    excludes.EntityData.SegmentPath = "excludes"
    excludes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    excludes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    excludes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    excludes.EntityData.Children = make(map[string]types.YChild)
    excludes.EntityData.Children["exclude"] = types.YChild{"Exclude", nil}
    for i := range excludes.Exclude {
        excludes.EntityData.Children[types.GetSegmentPath(&excludes.Exclude[i])] = types.YChild{"Exclude", &excludes.Exclude[i]}
    }
    excludes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(excludes.EntityData)
}

// AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Excludes_Exclude
// None
type AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Excludes_Exclude struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. First Address in IPv6 exclude range. The type is
    // one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    StartAddress interface{}

    // Last address in exclude Range. The type is one of the following types:
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'
    // This attribute is mandatory., or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'
    // This attribute is mandatory..
    EndAddress interface{}
}

func (exclude *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Excludes_Exclude) GetEntityData() *types.CommonEntityData {
    exclude.EntityData.YFilter = exclude.YFilter
    exclude.EntityData.YangName = "exclude"
    exclude.EntityData.BundleName = "cisco_ios_xr"
    exclude.EntityData.ParentYangName = "excludes"
    exclude.EntityData.SegmentPath = "exclude" + "[start-address='" + fmt.Sprintf("%v", exclude.StartAddress) + "']"
    exclude.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    exclude.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    exclude.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    exclude.EntityData.Children = make(map[string]types.YChild)
    exclude.EntityData.Leafs = make(map[string]types.YLeaf)
    exclude.EntityData.Leafs["start-address"] = types.YLeaf{"StartAddress", exclude.StartAddress}
    exclude.EntityData.Leafs["end-address"] = types.YLeaf{"EndAddress", exclude.EndAddress}
    return &(exclude.EntityData)
}

// AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_UtilizationMark
// Specify utilization mark
type AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_UtilizationMark struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify numerical value as percentage. The type is interface{} with range:
    // 0..100. Units are percentage.
    HighMark interface{}

    // Specify numerical value as percentage. The type is interface{} with range:
    // 0..100. Units are percentage.
    LowMark interface{}
}

func (utilizationMark *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_UtilizationMark) GetEntityData() *types.CommonEntityData {
    utilizationMark.EntityData.YFilter = utilizationMark.YFilter
    utilizationMark.EntityData.YangName = "utilization-mark"
    utilizationMark.EntityData.BundleName = "cisco_ios_xr"
    utilizationMark.EntityData.ParentYangName = "pool"
    utilizationMark.EntityData.SegmentPath = "utilization-mark"
    utilizationMark.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    utilizationMark.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    utilizationMark.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    utilizationMark.EntityData.Children = make(map[string]types.YChild)
    utilizationMark.EntityData.Leafs = make(map[string]types.YLeaf)
    utilizationMark.EntityData.Leafs["high-mark"] = types.YLeaf{"HighMark", utilizationMark.HighMark}
    utilizationMark.EntityData.Leafs["low-mark"] = types.YLeaf{"LowMark", utilizationMark.LowMark}
    return &(utilizationMark.EntityData)
}

// AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_PrefixRanges
// Specify prefix range for allocation
type AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_PrefixRanges struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // None. The type is slice of
    // AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_PrefixRanges_PrefixRange.
    PrefixRange []AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_PrefixRanges_PrefixRange
}

func (prefixRanges *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_PrefixRanges) GetEntityData() *types.CommonEntityData {
    prefixRanges.EntityData.YFilter = prefixRanges.YFilter
    prefixRanges.EntityData.YangName = "prefix-ranges"
    prefixRanges.EntityData.BundleName = "cisco_ios_xr"
    prefixRanges.EntityData.ParentYangName = "pool"
    prefixRanges.EntityData.SegmentPath = "prefix-ranges"
    prefixRanges.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixRanges.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixRanges.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixRanges.EntityData.Children = make(map[string]types.YChild)
    prefixRanges.EntityData.Children["prefix-range"] = types.YChild{"PrefixRange", nil}
    for i := range prefixRanges.PrefixRange {
        prefixRanges.EntityData.Children[types.GetSegmentPath(&prefixRanges.PrefixRange[i])] = types.YChild{"PrefixRange", &prefixRanges.PrefixRange[i]}
    }
    prefixRanges.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(prefixRanges.EntityData)
}

// AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_PrefixRanges_PrefixRange
// None
type AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_PrefixRanges_PrefixRange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. First prefix of range. The type is one of the
    // following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    StartPrefix interface{}

    // Blocked flag. The type is interface{} with range: -2147483648..2147483647.
    Blocked interface{}

    // Last prefix of range. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'
    // This attribute is mandatory., or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'
    // This attribute is mandatory..
    EndPrefix interface{}
}

func (prefixRange *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_PrefixRanges_PrefixRange) GetEntityData() *types.CommonEntityData {
    prefixRange.EntityData.YFilter = prefixRange.YFilter
    prefixRange.EntityData.YangName = "prefix-range"
    prefixRange.EntityData.BundleName = "cisco_ios_xr"
    prefixRange.EntityData.ParentYangName = "prefix-ranges"
    prefixRange.EntityData.SegmentPath = "prefix-range" + "[start-prefix='" + fmt.Sprintf("%v", prefixRange.StartPrefix) + "']"
    prefixRange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixRange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixRange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixRange.EntityData.Children = make(map[string]types.YChild)
    prefixRange.EntityData.Leafs = make(map[string]types.YLeaf)
    prefixRange.EntityData.Leafs["start-prefix"] = types.YLeaf{"StartPrefix", prefixRange.StartPrefix}
    prefixRange.EntityData.Leafs["blocked"] = types.YLeaf{"Blocked", prefixRange.Blocked}
    prefixRange.EntityData.Leafs["end-prefix"] = types.YLeaf{"EndPrefix", prefixRange.EndPrefix}
    return &(prefixRange.EntityData)
}

// AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Networks
// Specify network for allocation
type AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Networks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // None. The type is slice of
    // AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Networks_Network.
    Network []AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Networks_Network
}

func (networks *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Networks) GetEntityData() *types.CommonEntityData {
    networks.EntityData.YFilter = networks.YFilter
    networks.EntityData.YangName = "networks"
    networks.EntityData.BundleName = "cisco_ios_xr"
    networks.EntityData.ParentYangName = "pool"
    networks.EntityData.SegmentPath = "networks"
    networks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    networks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    networks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    networks.EntityData.Children = make(map[string]types.YChild)
    networks.EntityData.Children["network"] = types.YChild{"Network", nil}
    for i := range networks.Network {
        networks.EntityData.Children[types.GetSegmentPath(&networks.Network[i])] = types.YChild{"Network", &networks.Network[i]}
    }
    networks.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(networks.EntityData)
}

// AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Networks_Network
// None
type AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Networks_Network struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. None. The type is one of the following types:
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Prefix interface{}

    // Blocked flag. The type is interface{} with range: -2147483648..2147483647.
    Blocked interface{}

    // Prefix length for the IPv6 Prefix. The type is interface{} with range:
    // 1..128. This attribute is mandatory.
    PrefixLength interface{}
}

func (network *AddressPoolService_Vrfs_Vrf_Ipv6_Pools_Pool_Networks_Network) GetEntityData() *types.CommonEntityData {
    network.EntityData.YFilter = network.YFilter
    network.EntityData.YangName = "network"
    network.EntityData.BundleName = "cisco_ios_xr"
    network.EntityData.ParentYangName = "networks"
    network.EntityData.SegmentPath = "network" + "[prefix='" + fmt.Sprintf("%v", network.Prefix) + "']"
    network.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    network.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    network.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    network.EntityData.Children = make(map[string]types.YChild)
    network.EntityData.Leafs = make(map[string]types.YLeaf)
    network.EntityData.Leafs["prefix"] = types.YLeaf{"Prefix", network.Prefix}
    network.EntityData.Leafs["blocked"] = types.YLeaf{"Blocked", network.Blocked}
    network.EntityData.Leafs["prefix-length"] = types.YLeaf{"PrefixLength", network.PrefixLength}
    return &(network.EntityData)
}

// AddressPoolService_Vrfs_Vrf_Ipv4
// Enter IPv4 specific configuration
type AddressPoolService_Vrfs_Vrf_Ipv4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 Pool Table.
    Pools AddressPoolService_Vrfs_Vrf_Ipv4_Pools
}

func (ipv4 *AddressPoolService_Vrfs_Vrf_Ipv4) GetEntityData() *types.CommonEntityData {
    ipv4.EntityData.YFilter = ipv4.YFilter
    ipv4.EntityData.YangName = "ipv4"
    ipv4.EntityData.BundleName = "cisco_ios_xr"
    ipv4.EntityData.ParentYangName = "vrf"
    ipv4.EntityData.SegmentPath = "ipv4"
    ipv4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4.EntityData.Children = make(map[string]types.YChild)
    ipv4.EntityData.Children["pools"] = types.YChild{"Pools", &ipv4.Pools}
    ipv4.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipv4.EntityData)
}

// AddressPoolService_Vrfs_Vrf_Ipv4_Pools
// IPv4 Pool Table
type AddressPoolService_Vrfs_Vrf_Ipv4_Pools struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 Pool name. The type is slice of
    // AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool.
    Pool []AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool
}

func (pools *AddressPoolService_Vrfs_Vrf_Ipv4_Pools) GetEntityData() *types.CommonEntityData {
    pools.EntityData.YFilter = pools.YFilter
    pools.EntityData.YangName = "pools"
    pools.EntityData.BundleName = "cisco_ios_xr"
    pools.EntityData.ParentYangName = "ipv4"
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

// AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool
// IPv4 Pool name
type AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Enter the IPv4 Pool name. The type is string with
    // pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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

func (pool *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool) GetEntityData() *types.CommonEntityData {
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
    pool.EntityData.Children["excludes"] = types.YChild{"Excludes", &pool.Excludes}
    pool.EntityData.Children["utilization-mark"] = types.YChild{"UtilizationMark", &pool.UtilizationMark}
    pool.EntityData.Children["networks"] = types.YChild{"Networks", &pool.Networks}
    pool.EntityData.Leafs = make(map[string]types.YLeaf)
    pool.EntityData.Leafs["pool-name"] = types.YLeaf{"PoolName", pool.PoolName}
    return &(pool.EntityData)
}

// AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_AddressRanges
// address range for allocation
type AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_AddressRanges struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify first address in range. The type is slice of
    // AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_AddressRanges_AddressRange.
    AddressRange []AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_AddressRanges_AddressRange
}

func (addressRanges *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_AddressRanges) GetEntityData() *types.CommonEntityData {
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

// AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_AddressRanges_AddressRange
// Specify first address in range
type AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_AddressRanges_AddressRange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Specify first address of the range. The type is
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    StartAddress interface{}

    // Last address of the range. The type is one of the following types: string
    // with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'
    // This attribute is mandatory., or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'
    // This attribute is mandatory..
    EndAddress interface{}

    // Blocked flag. The type is interface{} with range: -2147483648..2147483647.
    Blocked interface{}
}

func (addressRange *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_AddressRanges_AddressRange) GetEntityData() *types.CommonEntityData {
    addressRange.EntityData.YFilter = addressRange.YFilter
    addressRange.EntityData.YangName = "address-range"
    addressRange.EntityData.BundleName = "cisco_ios_xr"
    addressRange.EntityData.ParentYangName = "address-ranges"
    addressRange.EntityData.SegmentPath = "address-range" + "[start-address='" + fmt.Sprintf("%v", addressRange.StartAddress) + "']"
    addressRange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressRange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressRange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressRange.EntityData.Children = make(map[string]types.YChild)
    addressRange.EntityData.Leafs = make(map[string]types.YLeaf)
    addressRange.EntityData.Leafs["start-address"] = types.YLeaf{"StartAddress", addressRange.StartAddress}
    addressRange.EntityData.Leafs["end-address"] = types.YLeaf{"EndAddress", addressRange.EndAddress}
    addressRange.EntityData.Leafs["blocked"] = types.YLeaf{"Blocked", addressRange.Blocked}
    return &(addressRange.EntityData)
}

// AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Excludes
// Exclude addresses
type AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Excludes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First address in range. The type is slice of
    // AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Excludes_Exclude.
    Exclude []AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Excludes_Exclude
}

func (excludes *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Excludes) GetEntityData() *types.CommonEntityData {
    excludes.EntityData.YFilter = excludes.YFilter
    excludes.EntityData.YangName = "excludes"
    excludes.EntityData.BundleName = "cisco_ios_xr"
    excludes.EntityData.ParentYangName = "pool"
    excludes.EntityData.SegmentPath = "excludes"
    excludes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    excludes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    excludes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    excludes.EntityData.Children = make(map[string]types.YChild)
    excludes.EntityData.Children["exclude"] = types.YChild{"Exclude", nil}
    for i := range excludes.Exclude {
        excludes.EntityData.Children[types.GetSegmentPath(&excludes.Exclude[i])] = types.YChild{"Exclude", &excludes.Exclude[i]}
    }
    excludes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(excludes.EntityData)
}

// AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Excludes_Exclude
// First address in range
type AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Excludes_Exclude struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. First address in exclude range. The type is string
    // with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    StartAddress interface{}

    // Last address in excluded range. The type is one of the following types:
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'
    // This attribute is mandatory., or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'
    // This attribute is mandatory..
    EndAddress interface{}
}

func (exclude *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Excludes_Exclude) GetEntityData() *types.CommonEntityData {
    exclude.EntityData.YFilter = exclude.YFilter
    exclude.EntityData.YangName = "exclude"
    exclude.EntityData.BundleName = "cisco_ios_xr"
    exclude.EntityData.ParentYangName = "excludes"
    exclude.EntityData.SegmentPath = "exclude" + "[start-address='" + fmt.Sprintf("%v", exclude.StartAddress) + "']"
    exclude.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    exclude.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    exclude.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    exclude.EntityData.Children = make(map[string]types.YChild)
    exclude.EntityData.Leafs = make(map[string]types.YLeaf)
    exclude.EntityData.Leafs["start-address"] = types.YLeaf{"StartAddress", exclude.StartAddress}
    exclude.EntityData.Leafs["end-address"] = types.YLeaf{"EndAddress", exclude.EndAddress}
    return &(exclude.EntityData)
}

// AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_UtilizationMark
// Specify utilization mark
type AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_UtilizationMark struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify numerical value as percentage. The type is interface{} with range:
    // 0..100. Units are percentage.
    High interface{}

    // Specify numerical value as percentage. The type is interface{} with range:
    // 0..100. Units are percentage.
    Low interface{}
}

func (utilizationMark *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_UtilizationMark) GetEntityData() *types.CommonEntityData {
    utilizationMark.EntityData.YFilter = utilizationMark.YFilter
    utilizationMark.EntityData.YangName = "utilization-mark"
    utilizationMark.EntityData.BundleName = "cisco_ios_xr"
    utilizationMark.EntityData.ParentYangName = "pool"
    utilizationMark.EntityData.SegmentPath = "utilization-mark"
    utilizationMark.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    utilizationMark.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    utilizationMark.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    utilizationMark.EntityData.Children = make(map[string]types.YChild)
    utilizationMark.EntityData.Leafs = make(map[string]types.YLeaf)
    utilizationMark.EntityData.Leafs["high"] = types.YLeaf{"High", utilizationMark.High}
    utilizationMark.EntityData.Leafs["low"] = types.YLeaf{"Low", utilizationMark.Low}
    return &(utilizationMark.EntityData)
}

// AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Networks
// Specify network for allocation
type AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Networks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Network Prefix. The type is slice of
    // AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Networks_Network.
    Network []AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Networks_Network
}

func (networks *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Networks) GetEntityData() *types.CommonEntityData {
    networks.EntityData.YFilter = networks.YFilter
    networks.EntityData.YangName = "networks"
    networks.EntityData.BundleName = "cisco_ios_xr"
    networks.EntityData.ParentYangName = "pool"
    networks.EntityData.SegmentPath = "networks"
    networks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    networks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    networks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    networks.EntityData.Children = make(map[string]types.YChild)
    networks.EntityData.Children["network"] = types.YChild{"Network", nil}
    for i := range networks.Network {
        networks.EntityData.Children[types.GetSegmentPath(&networks.Network[i])] = types.YChild{"Network", &networks.Network[i]}
    }
    networks.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(networks.EntityData)
}

// AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Networks_Network
// Network Prefix
type AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Networks_Network struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. None. The type is one of the following types:
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv4Prefix interface{}

    // Blocked flag. The type is interface{} with range: -2147483648..2147483647.
    Blocked interface{}

    // Subnet Length for IPv4 subnet. The type is interface{} with range:
    // -2147483648..2147483647. This attribute is mandatory.
    PrefixLength interface{}

    // Default Gateway for IPv4 subnet. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    DefaultRouter interface{}
}

func (network *AddressPoolService_Vrfs_Vrf_Ipv4_Pools_Pool_Networks_Network) GetEntityData() *types.CommonEntityData {
    network.EntityData.YFilter = network.YFilter
    network.EntityData.YangName = "network"
    network.EntityData.BundleName = "cisco_ios_xr"
    network.EntityData.ParentYangName = "networks"
    network.EntityData.SegmentPath = "network" + "[ipv4-prefix='" + fmt.Sprintf("%v", network.Ipv4Prefix) + "']"
    network.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    network.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    network.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    network.EntityData.Children = make(map[string]types.YChild)
    network.EntityData.Leafs = make(map[string]types.YLeaf)
    network.EntityData.Leafs["ipv4-prefix"] = types.YLeaf{"Ipv4Prefix", network.Ipv4Prefix}
    network.EntityData.Leafs["blocked"] = types.YLeaf{"Blocked", network.Blocked}
    network.EntityData.Leafs["prefix-length"] = types.YLeaf{"PrefixLength", network.PrefixLength}
    network.EntityData.Leafs["default-router"] = types.YLeaf{"DefaultRouter", network.DefaultRouter}
    return &(network.EntityData)
}

