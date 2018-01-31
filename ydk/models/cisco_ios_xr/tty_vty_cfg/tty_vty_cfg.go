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
    parent types.Entity
    YFilter yfilter.YFilter

    // List of VTY Pools.
    VtyPools Vty_VtyPools
}

func (vty *Vty) GetFilter() yfilter.YFilter { return vty.YFilter }

func (vty *Vty) SetFilter(yf yfilter.YFilter) { vty.YFilter = yf }

func (vty *Vty) GetGoName(yname string) string {
    if yname == "vty-pools" { return "VtyPools" }
    return ""
}

func (vty *Vty) GetSegmentPath() string {
    return "Cisco-IOS-XR-tty-vty-cfg:vty"
}

func (vty *Vty) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vty-pools" {
        return &vty.VtyPools
    }
    return nil
}

func (vty *Vty) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["vty-pools"] = &vty.VtyPools
    return children
}

func (vty *Vty) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vty *Vty) GetBundleName() string { return "cisco_ios_xr" }

func (vty *Vty) GetYangName() string { return "vty" }

func (vty *Vty) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vty *Vty) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vty *Vty) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vty *Vty) SetParent(parent types.Entity) { vty.parent = parent }

func (vty *Vty) GetParent() types.Entity { return vty.parent }

func (vty *Vty) GetParentYangName() string { return "Cisco-IOS-XR-tty-vty-cfg" }

// Vty_VtyPools
// List of VTY Pools
type Vty_VtyPools struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VTY Pool. The type is slice of Vty_VtyPools_VtyPool.
    VtyPool []Vty_VtyPools_VtyPool
}

func (vtyPools *Vty_VtyPools) GetFilter() yfilter.YFilter { return vtyPools.YFilter }

func (vtyPools *Vty_VtyPools) SetFilter(yf yfilter.YFilter) { vtyPools.YFilter = yf }

func (vtyPools *Vty_VtyPools) GetGoName(yname string) string {
    if yname == "vty-pool" { return "VtyPool" }
    return ""
}

func (vtyPools *Vty_VtyPools) GetSegmentPath() string {
    return "vty-pools"
}

func (vtyPools *Vty_VtyPools) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vty-pool" {
        for _, c := range vtyPools.VtyPool {
            if vtyPools.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vty_VtyPools_VtyPool{}
        vtyPools.VtyPool = append(vtyPools.VtyPool, child)
        return &vtyPools.VtyPool[len(vtyPools.VtyPool)-1]
    }
    return nil
}

func (vtyPools *Vty_VtyPools) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vtyPools.VtyPool {
        children[vtyPools.VtyPool[i].GetSegmentPath()] = &vtyPools.VtyPool[i]
    }
    return children
}

func (vtyPools *Vty_VtyPools) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vtyPools *Vty_VtyPools) GetBundleName() string { return "cisco_ios_xr" }

func (vtyPools *Vty_VtyPools) GetYangName() string { return "vty-pools" }

func (vtyPools *Vty_VtyPools) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vtyPools *Vty_VtyPools) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vtyPools *Vty_VtyPools) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vtyPools *Vty_VtyPools) SetParent(parent types.Entity) { vtyPools.parent = parent }

func (vtyPools *Vty_VtyPools) GetParent() types.Entity { return vtyPools.parent }

func (vtyPools *Vty_VtyPools) GetParentYangName() string { return "vty" }

// Vty_VtyPools_VtyPool
// VTY Pool
type Vty_VtyPools_VtyPool struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. For configuring range for default pool use
    // 'default',For configuring range for fault-manager pool use 'fm',For
    // configuring range for any user defined pool use any other string. The type
    // is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
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

func (vtyPool *Vty_VtyPools_VtyPool) GetFilter() yfilter.YFilter { return vtyPool.YFilter }

func (vtyPool *Vty_VtyPools_VtyPool) SetFilter(yf yfilter.YFilter) { vtyPool.YFilter = yf }

func (vtyPool *Vty_VtyPools_VtyPool) GetGoName(yname string) string {
    if yname == "pool-name" { return "PoolName" }
    if yname == "first-vty" { return "FirstVty" }
    if yname == "last-vty" { return "LastVty" }
    if yname == "line-template" { return "LineTemplate" }
    if yname == "none" { return "None" }
    return ""
}

func (vtyPool *Vty_VtyPools_VtyPool) GetSegmentPath() string {
    return "vty-pool" + "[pool-name='" + fmt.Sprintf("%v", vtyPool.PoolName) + "']"
}

func (vtyPool *Vty_VtyPools_VtyPool) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vtyPool *Vty_VtyPools_VtyPool) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vtyPool *Vty_VtyPools_VtyPool) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pool-name"] = vtyPool.PoolName
    leafs["first-vty"] = vtyPool.FirstVty
    leafs["last-vty"] = vtyPool.LastVty
    leafs["line-template"] = vtyPool.LineTemplate
    leafs["none"] = vtyPool.None
    return leafs
}

func (vtyPool *Vty_VtyPools_VtyPool) GetBundleName() string { return "cisco_ios_xr" }

func (vtyPool *Vty_VtyPools_VtyPool) GetYangName() string { return "vty-pool" }

func (vtyPool *Vty_VtyPools_VtyPool) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vtyPool *Vty_VtyPools_VtyPool) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vtyPool *Vty_VtyPools_VtyPool) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vtyPool *Vty_VtyPools_VtyPool) SetParent(parent types.Entity) { vtyPool.parent = parent }

func (vtyPool *Vty_VtyPools_VtyPool) GetParent() types.Entity { return vtyPool.parent }

func (vtyPool *Vty_VtyPools_VtyPool) GetParentYangName() string { return "vty-pools" }

