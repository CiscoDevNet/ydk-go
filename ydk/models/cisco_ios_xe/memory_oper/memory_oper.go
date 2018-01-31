// This module contains a collection of YANG definitions for
// monitoring memory in a Network Element.
// Copyright (c) 2016-2017 by Cisco Systems, Inc.
// All rights reserved.
package memory_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package memory_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-memory-oper memory-statistics}", reflect.TypeOf(MemoryStatistics{}))
    ydk.RegisterEntity("Cisco-IOS-XE-memory-oper:memory-statistics", reflect.TypeOf(MemoryStatistics{}))
}

// MemoryStatistics
// Data nodes for All Memory Pool Statistics.
type MemoryStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The list of software memory pools in the system. The type is slice of
    // MemoryStatistics_MemoryStatistic.
    MemoryStatistic []MemoryStatistics_MemoryStatistic
}

func (memoryStatistics *MemoryStatistics) GetFilter() yfilter.YFilter { return memoryStatistics.YFilter }

func (memoryStatistics *MemoryStatistics) SetFilter(yf yfilter.YFilter) { memoryStatistics.YFilter = yf }

func (memoryStatistics *MemoryStatistics) GetGoName(yname string) string {
    if yname == "memory-statistic" { return "MemoryStatistic" }
    return ""
}

func (memoryStatistics *MemoryStatistics) GetSegmentPath() string {
    return "Cisco-IOS-XE-memory-oper:memory-statistics"
}

func (memoryStatistics *MemoryStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "memory-statistic" {
        for _, c := range memoryStatistics.MemoryStatistic {
            if memoryStatistics.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MemoryStatistics_MemoryStatistic{}
        memoryStatistics.MemoryStatistic = append(memoryStatistics.MemoryStatistic, child)
        return &memoryStatistics.MemoryStatistic[len(memoryStatistics.MemoryStatistic)-1]
    }
    return nil
}

func (memoryStatistics *MemoryStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range memoryStatistics.MemoryStatistic {
        children[memoryStatistics.MemoryStatistic[i].GetSegmentPath()] = &memoryStatistics.MemoryStatistic[i]
    }
    return children
}

func (memoryStatistics *MemoryStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (memoryStatistics *MemoryStatistics) GetBundleName() string { return "cisco_ios_xe" }

func (memoryStatistics *MemoryStatistics) GetYangName() string { return "memory-statistics" }

func (memoryStatistics *MemoryStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (memoryStatistics *MemoryStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (memoryStatistics *MemoryStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (memoryStatistics *MemoryStatistics) SetParent(parent types.Entity) { memoryStatistics.parent = parent }

func (memoryStatistics *MemoryStatistics) GetParent() types.Entity { return memoryStatistics.parent }

func (memoryStatistics *MemoryStatistics) GetParentYangName() string { return "Cisco-IOS-XE-memory-oper" }

// MemoryStatistics_MemoryStatistic
// The list of software memory pools in the system.
type MemoryStatistics_MemoryStatistic struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the memory pool. The type is string.
    Name interface{}

    // Total memory in the pool (bytes). The type is interface{} with range:
    // 0..18446744073709551615. Units are bytes.
    TotalMemory interface{}

    // Total used memory in the pool (bytes). The type is interface{} with range:
    // 0..18446744073709551615. Units are bytes.
    UsedMemory interface{}

    // Total free memory in the pool (bytes). The type is interface{} with range:
    // 0..18446744073709551615. Units are bytes.
    FreeMemory interface{}

    // Historical lowest memory usage (bytes). The type is interface{} with range:
    // 0..18446744073709551615. Units are bytes.
    LowestUsage interface{}

    // Historical highest memory usage (bytes). The type is interface{} with
    // range: 0..18446744073709551615. Units are bytes.
    HighestUsage interface{}
}

func (memoryStatistic *MemoryStatistics_MemoryStatistic) GetFilter() yfilter.YFilter { return memoryStatistic.YFilter }

func (memoryStatistic *MemoryStatistics_MemoryStatistic) SetFilter(yf yfilter.YFilter) { memoryStatistic.YFilter = yf }

func (memoryStatistic *MemoryStatistics_MemoryStatistic) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "total-memory" { return "TotalMemory" }
    if yname == "used-memory" { return "UsedMemory" }
    if yname == "free-memory" { return "FreeMemory" }
    if yname == "lowest-usage" { return "LowestUsage" }
    if yname == "highest-usage" { return "HighestUsage" }
    return ""
}

func (memoryStatistic *MemoryStatistics_MemoryStatistic) GetSegmentPath() string {
    return "memory-statistic" + "[name='" + fmt.Sprintf("%v", memoryStatistic.Name) + "']"
}

func (memoryStatistic *MemoryStatistics_MemoryStatistic) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (memoryStatistic *MemoryStatistics_MemoryStatistic) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (memoryStatistic *MemoryStatistics_MemoryStatistic) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = memoryStatistic.Name
    leafs["total-memory"] = memoryStatistic.TotalMemory
    leafs["used-memory"] = memoryStatistic.UsedMemory
    leafs["free-memory"] = memoryStatistic.FreeMemory
    leafs["lowest-usage"] = memoryStatistic.LowestUsage
    leafs["highest-usage"] = memoryStatistic.HighestUsage
    return leafs
}

func (memoryStatistic *MemoryStatistics_MemoryStatistic) GetBundleName() string { return "cisco_ios_xe" }

func (memoryStatistic *MemoryStatistics_MemoryStatistic) GetYangName() string { return "memory-statistic" }

func (memoryStatistic *MemoryStatistics_MemoryStatistic) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (memoryStatistic *MemoryStatistics_MemoryStatistic) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (memoryStatistic *MemoryStatistics_MemoryStatistic) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (memoryStatistic *MemoryStatistics_MemoryStatistic) SetParent(parent types.Entity) { memoryStatistic.parent = parent }

func (memoryStatistic *MemoryStatistics_MemoryStatistic) GetParent() types.Entity { return memoryStatistic.parent }

func (memoryStatistic *MemoryStatistics_MemoryStatistic) GetParentYangName() string { return "memory-statistics" }

