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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of service instance stats. The type is slice of EfpStats_EfpStat.
    EfpStat []EfpStats_EfpStat
}

func (efpStats *EfpStats) GetEntityData() *types.CommonEntityData {
    efpStats.EntityData.YFilter = efpStats.YFilter
    efpStats.EntityData.YangName = "efp-stats"
    efpStats.EntityData.BundleName = "cisco_ios_xe"
    efpStats.EntityData.ParentYangName = "Cisco-IOS-XE-efp-oper"
    efpStats.EntityData.SegmentPath = "Cisco-IOS-XE-efp-oper:efp-stats"
    efpStats.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    efpStats.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    efpStats.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    efpStats.EntityData.Children = make(map[string]types.YChild)
    efpStats.EntityData.Children["efp-stat"] = types.YChild{"EfpStat", nil}
    for i := range efpStats.EfpStat {
        efpStats.EntityData.Children[types.GetSegmentPath(&efpStats.EfpStat[i])] = types.YChild{"EfpStat", &efpStats.EfpStat[i]}
    }
    efpStats.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(efpStats.EntityData)
}

// EfpStats_EfpStat
// List of service instance stats
type EfpStats_EfpStat struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. EFP id. The type is interface{} with range:
    // 0..4294967295.
    Id interface{}

    // This attribute is a key. Interface name. The type is string.
    Interface_ interface{}

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

func (efpStat *EfpStats_EfpStat) GetEntityData() *types.CommonEntityData {
    efpStat.EntityData.YFilter = efpStat.YFilter
    efpStat.EntityData.YangName = "efp-stat"
    efpStat.EntityData.BundleName = "cisco_ios_xe"
    efpStat.EntityData.ParentYangName = "efp-stats"
    efpStat.EntityData.SegmentPath = "efp-stat" + "[id='" + fmt.Sprintf("%v", efpStat.Id) + "']" + "[interface='" + fmt.Sprintf("%v", efpStat.Interface_) + "']"
    efpStat.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    efpStat.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    efpStat.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    efpStat.EntityData.Children = make(map[string]types.YChild)
    efpStat.EntityData.Leafs = make(map[string]types.YLeaf)
    efpStat.EntityData.Leafs["id"] = types.YLeaf{"Id", efpStat.Id}
    efpStat.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", efpStat.Interface_}
    efpStat.EntityData.Leafs["in-pkts"] = types.YLeaf{"InPkts", efpStat.InPkts}
    efpStat.EntityData.Leafs["in-bytes"] = types.YLeaf{"InBytes", efpStat.InBytes}
    efpStat.EntityData.Leafs["out-pkts"] = types.YLeaf{"OutPkts", efpStat.OutPkts}
    efpStat.EntityData.Leafs["out-bytes"] = types.YLeaf{"OutBytes", efpStat.OutBytes}
    return &(efpStat.EntityData)
}

