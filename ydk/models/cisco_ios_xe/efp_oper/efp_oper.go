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
    EfpStat []*EfpStats_EfpStat
}

func (efpStats *EfpStats) GetEntityData() *types.CommonEntityData {
    efpStats.EntityData.YFilter = efpStats.YFilter
    efpStats.EntityData.YangName = "efp-stats"
    efpStats.EntityData.BundleName = "cisco_ios_xe"
    efpStats.EntityData.ParentYangName = "Cisco-IOS-XE-efp-oper"
    efpStats.EntityData.SegmentPath = "Cisco-IOS-XE-efp-oper:efp-stats"
    efpStats.EntityData.AbsolutePath = efpStats.EntityData.SegmentPath
    efpStats.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    efpStats.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    efpStats.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    efpStats.EntityData.Children = types.NewOrderedMap()
    efpStats.EntityData.Children.Append("efp-stat", types.YChild{"EfpStat", nil})
    for i := range efpStats.EfpStat {
        efpStats.EntityData.Children.Append(types.GetSegmentPath(efpStats.EfpStat[i]), types.YChild{"EfpStat", efpStats.EfpStat[i]})
    }
    efpStats.EntityData.Leafs = types.NewOrderedMap()

    efpStats.EntityData.YListKeys = []string {}

    return &(efpStats.EntityData)
}

// EfpStats_EfpStat
// List of service instance stats
type EfpStats_EfpStat struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (efpStat *EfpStats_EfpStat) GetEntityData() *types.CommonEntityData {
    efpStat.EntityData.YFilter = efpStat.YFilter
    efpStat.EntityData.YangName = "efp-stat"
    efpStat.EntityData.BundleName = "cisco_ios_xe"
    efpStat.EntityData.ParentYangName = "efp-stats"
    efpStat.EntityData.SegmentPath = "efp-stat" + types.AddKeyToken(efpStat.Id, "id") + types.AddKeyToken(efpStat.Interface, "interface")
    efpStat.EntityData.AbsolutePath = "Cisco-IOS-XE-efp-oper:efp-stats/" + efpStat.EntityData.SegmentPath
    efpStat.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    efpStat.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    efpStat.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    efpStat.EntityData.Children = types.NewOrderedMap()
    efpStat.EntityData.Leafs = types.NewOrderedMap()
    efpStat.EntityData.Leafs.Append("id", types.YLeaf{"Id", efpStat.Id})
    efpStat.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", efpStat.Interface})
    efpStat.EntityData.Leafs.Append("in-pkts", types.YLeaf{"InPkts", efpStat.InPkts})
    efpStat.EntityData.Leafs.Append("in-bytes", types.YLeaf{"InBytes", efpStat.InBytes})
    efpStat.EntityData.Leafs.Append("out-pkts", types.YLeaf{"OutPkts", efpStat.OutPkts})
    efpStat.EntityData.Leafs.Append("out-bytes", types.YLeaf{"OutBytes", efpStat.OutBytes})

    efpStat.EntityData.YListKeys = []string {"Id", "Interface"}

    return &(efpStat.EntityData)
}

