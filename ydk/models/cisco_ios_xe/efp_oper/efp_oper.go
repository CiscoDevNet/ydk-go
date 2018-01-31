// This module contains a collection of YANG definitions
// for service instance (efp) stats.
// Copyright (c) 2016-2017 by Cisco Systems, Inc.
// All rights reserved.
package efp_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package efp_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-efp-oper efp-stats}", reflect.TypeOf(EfpStats{}))
    ydk.RegisterEntity("Cisco-IOS-XE-efp-oper:efp-stats", reflect.TypeOf(EfpStats{}))
}

// EfpStats
// Service instance stats
type EfpStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of service instance stats. The type is slice of EfpStats_EfpStat.
    EfpStat []EfpStats_EfpStat
}

func (efpStats *EfpStats) GetFilter() yfilter.YFilter { return efpStats.YFilter }

func (efpStats *EfpStats) SetFilter(yf yfilter.YFilter) { efpStats.YFilter = yf }

func (efpStats *EfpStats) GetGoName(yname string) string {
    if yname == "efp-stat" { return "EfpStat" }
    return ""
}

func (efpStats *EfpStats) GetSegmentPath() string {
    return "Cisco-IOS-XE-efp-oper:efp-stats"
}

func (efpStats *EfpStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "efp-stat" {
        for _, c := range efpStats.EfpStat {
            if efpStats.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := EfpStats_EfpStat{}
        efpStats.EfpStat = append(efpStats.EfpStat, child)
        return &efpStats.EfpStat[len(efpStats.EfpStat)-1]
    }
    return nil
}

func (efpStats *EfpStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range efpStats.EfpStat {
        children[efpStats.EfpStat[i].GetSegmentPath()] = &efpStats.EfpStat[i]
    }
    return children
}

func (efpStats *EfpStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (efpStats *EfpStats) GetBundleName() string { return "cisco_ios_xe" }

func (efpStats *EfpStats) GetYangName() string { return "efp-stats" }

func (efpStats *EfpStats) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (efpStats *EfpStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (efpStats *EfpStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (efpStats *EfpStats) SetParent(parent types.Entity) { efpStats.parent = parent }

func (efpStats *EfpStats) GetParent() types.Entity { return efpStats.parent }

func (efpStats *EfpStats) GetParentYangName() string { return "Cisco-IOS-XE-efp-oper" }

// EfpStats_EfpStat
// List of service instance stats
type EfpStats_EfpStat struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. EFP id. The type is interface{} with range:
    // 0..4294967295.
    Id interface{}

    // This attribute is a key. Interface name. The type is string.
    Interface interface{}

    // Incoming packets. The type is interface{} with range:
    // 0..18446744073709551615.
    InPkts interface{}

    // Incoming bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    InBytes interface{}

    // Outgoing packets. The type is interface{} with range:
    // 0..18446744073709551615.
    OutPkts interface{}

    // Outgoing bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    OutBytes interface{}
}

func (efpStat *EfpStats_EfpStat) GetFilter() yfilter.YFilter { return efpStat.YFilter }

func (efpStat *EfpStats_EfpStat) SetFilter(yf yfilter.YFilter) { efpStat.YFilter = yf }

func (efpStat *EfpStats_EfpStat) GetGoName(yname string) string {
    if yname == "id" { return "Id" }
    if yname == "interface" { return "Interface" }
    if yname == "in-pkts" { return "InPkts" }
    if yname == "in-bytes" { return "InBytes" }
    if yname == "out-pkts" { return "OutPkts" }
    if yname == "out-bytes" { return "OutBytes" }
    return ""
}

func (efpStat *EfpStats_EfpStat) GetSegmentPath() string {
    return "efp-stat" + "[id='" + fmt.Sprintf("%v", efpStat.Id) + "']" + "[interface='" + fmt.Sprintf("%v", efpStat.Interface) + "']"
}

func (efpStat *EfpStats_EfpStat) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (efpStat *EfpStats_EfpStat) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (efpStat *EfpStats_EfpStat) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["id"] = efpStat.Id
    leafs["interface"] = efpStat.Interface
    leafs["in-pkts"] = efpStat.InPkts
    leafs["in-bytes"] = efpStat.InBytes
    leafs["out-pkts"] = efpStat.OutPkts
    leafs["out-bytes"] = efpStat.OutBytes
    return leafs
}

func (efpStat *EfpStats_EfpStat) GetBundleName() string { return "cisco_ios_xe" }

func (efpStat *EfpStats_EfpStat) GetYangName() string { return "efp-stat" }

func (efpStat *EfpStats_EfpStat) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (efpStat *EfpStats_EfpStat) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (efpStat *EfpStats_EfpStat) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (efpStat *EfpStats_EfpStat) SetParent(parent types.Entity) { efpStat.parent = parent }

func (efpStat *EfpStats_EfpStat) GetParent() types.Entity { return efpStat.parent }

func (efpStat *EfpStats_EfpStat) GetParentYangName() string { return "efp-stats" }

