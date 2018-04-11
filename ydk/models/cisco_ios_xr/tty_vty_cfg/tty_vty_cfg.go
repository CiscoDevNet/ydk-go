// This module contains a collection of YANG definitions
// for Cisco IOS-XR tty-vty package configuration.
// 
// This module contains definitions
// for the following management objects:
//   vty: VTY Pools configuration
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package tty_vty_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package tty_vty_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-tty-vty-cfg vty}", reflect.TypeOf(Vty{}))
    ydk.RegisterEntity("Cisco-IOS-XR-tty-vty-cfg:vty", reflect.TypeOf(Vty{}))
}

// Vty
// VTY Pools configuration
type Vty struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of VTY Pools.
    VtyPools Vty_VtyPools
}

func (vty *Vty) GetEntityData() *types.CommonEntityData {
    vty.EntityData.YFilter = vty.YFilter
    vty.EntityData.YangName = "vty"
    vty.EntityData.BundleName = "cisco_ios_xr"
    vty.EntityData.ParentYangName = "Cisco-IOS-XR-tty-vty-cfg"
    vty.EntityData.SegmentPath = "Cisco-IOS-XR-tty-vty-cfg:vty"
    vty.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vty.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vty.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vty.EntityData.Children = make(map[string]types.YChild)
    vty.EntityData.Children["vty-pools"] = types.YChild{"VtyPools", &vty.VtyPools}
    vty.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(vty.EntityData)
}

// Vty_VtyPools
// List of VTY Pools
type Vty_VtyPools struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VTY Pool. The type is slice of Vty_VtyPools_VtyPool.
    VtyPool []Vty_VtyPools_VtyPool
}

func (vtyPools *Vty_VtyPools) GetEntityData() *types.CommonEntityData {
    vtyPools.EntityData.YFilter = vtyPools.YFilter
    vtyPools.EntityData.YangName = "vty-pools"
    vtyPools.EntityData.BundleName = "cisco_ios_xr"
    vtyPools.EntityData.ParentYangName = "vty"
    vtyPools.EntityData.SegmentPath = "vty-pools"
    vtyPools.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vtyPools.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vtyPools.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vtyPools.EntityData.Children = make(map[string]types.YChild)
    vtyPools.EntityData.Children["vty-pool"] = types.YChild{"VtyPool", nil}
    for i := range vtyPools.VtyPool {
        vtyPools.EntityData.Children[types.GetSegmentPath(&vtyPools.VtyPool[i])] = types.YChild{"VtyPool", &vtyPools.VtyPool[i]}
    }
    vtyPools.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(vtyPools.EntityData)
}

// Vty_VtyPools_VtyPool
// VTY Pool
type Vty_VtyPools_VtyPool struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. For configuring range for default pool use
    // 'default',For configuring range for fault-manager pool use 'fm',For
    // configuring range for any user defined pool use any other string. The type
    // is string with pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    PoolName interface{}

    // First VTY number,For default VTY use 0,For user-defined use 5,For
    // fault-manager use 100. The type is interface{} with range:
    // -2147483648..2147483647. This attribute is mandatory.
    FirstVty interface{}

    // Last VTY number,For default configure between 0-99,For user-defined
    // configure between 5-99 ,For fault-manager configure between 100-199. The
    // type is interface{} with range: -2147483648..2147483647. This attribute is
    // mandatory.
    LastVty interface{}

    // Name of line template. The type is string.
    LineTemplate interface{}

    // Empty Option. The type is string.
    None interface{}
}

func (vtyPool *Vty_VtyPools_VtyPool) GetEntityData() *types.CommonEntityData {
    vtyPool.EntityData.YFilter = vtyPool.YFilter
    vtyPool.EntityData.YangName = "vty-pool"
    vtyPool.EntityData.BundleName = "cisco_ios_xr"
    vtyPool.EntityData.ParentYangName = "vty-pools"
    vtyPool.EntityData.SegmentPath = "vty-pool" + "[pool-name='" + fmt.Sprintf("%v", vtyPool.PoolName) + "']"
    vtyPool.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vtyPool.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vtyPool.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vtyPool.EntityData.Children = make(map[string]types.YChild)
    vtyPool.EntityData.Leafs = make(map[string]types.YLeaf)
    vtyPool.EntityData.Leafs["pool-name"] = types.YLeaf{"PoolName", vtyPool.PoolName}
    vtyPool.EntityData.Leafs["first-vty"] = types.YLeaf{"FirstVty", vtyPool.FirstVty}
    vtyPool.EntityData.Leafs["last-vty"] = types.YLeaf{"LastVty", vtyPool.LastVty}
    vtyPool.EntityData.Leafs["line-template"] = types.YLeaf{"LineTemplate", vtyPool.LineTemplate}
    vtyPool.EntityData.Leafs["none"] = types.YLeaf{"None", vtyPool.None}
    return &(vtyPool.EntityData)
}

