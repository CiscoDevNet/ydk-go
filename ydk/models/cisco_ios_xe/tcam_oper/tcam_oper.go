// This module contains a collection of YANG definitions
// for ASIC TCAM Memory Utilization Statistics.
// Copyright (c) 2016-2017 by Cisco Systems, Inc.
// All rights reserved.
package tcam_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package tcam_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-tcam-oper tcam-details}", reflect.TypeOf(TcamDetails{}))
    ydk.RegisterEntity("Cisco-IOS-XE-tcam-oper:tcam-details", reflect.TypeOf(TcamDetails{}))
}

// TcamDetails
// ASIC TCAM Memory Statistics
type TcamDetails struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // FED ASIC TCAM Utilization. The type is slice of TcamDetails_TcamDetail.
    TcamDetail []TcamDetails_TcamDetail
}

func (tcamDetails *TcamDetails) GetEntityData() *types.CommonEntityData {
    tcamDetails.EntityData.YFilter = tcamDetails.YFilter
    tcamDetails.EntityData.YangName = "tcam-details"
    tcamDetails.EntityData.BundleName = "cisco_ios_xe"
    tcamDetails.EntityData.ParentYangName = "Cisco-IOS-XE-tcam-oper"
    tcamDetails.EntityData.SegmentPath = "Cisco-IOS-XE-tcam-oper:tcam-details"
    tcamDetails.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    tcamDetails.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    tcamDetails.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    tcamDetails.EntityData.Children = make(map[string]types.YChild)
    tcamDetails.EntityData.Children["tcam-detail"] = types.YChild{"TcamDetail", nil}
    for i := range tcamDetails.TcamDetail {
        tcamDetails.EntityData.Children[types.GetSegmentPath(&tcamDetails.TcamDetail[i])] = types.YChild{"TcamDetail", &tcamDetails.TcamDetail[i]}
    }
    tcamDetails.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(tcamDetails.EntityData)
}

// TcamDetails_TcamDetail
// FED ASIC TCAM Utilization
type TcamDetails_TcamDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Asic Number of Switch. The type is interface{}
    // with range: 0..8.
    AsicNo interface{}

    // This attribute is a key. Protocol Name. The type is string.
    Name interface{}

    // Maximum Hash Entries. The type is interface{} with range: 0..4294967295.
    HashEntriesMax interface{}

    // Maximum Tcam Entries. The type is interface{} with range: 0..4294967295.
    TcamEntriesMax interface{}

    // Hash Entries Used. The type is interface{} with range: 0..4294967295.
    HashEntriesUsed interface{}

    // Tcam Entries Used. The type is interface{} with range: 0..4294967295.
    TcamEntriesUsed interface{}
}

func (tcamDetail *TcamDetails_TcamDetail) GetEntityData() *types.CommonEntityData {
    tcamDetail.EntityData.YFilter = tcamDetail.YFilter
    tcamDetail.EntityData.YangName = "tcam-detail"
    tcamDetail.EntityData.BundleName = "cisco_ios_xe"
    tcamDetail.EntityData.ParentYangName = "tcam-details"
    tcamDetail.EntityData.SegmentPath = "tcam-detail" + "[asic-no='" + fmt.Sprintf("%v", tcamDetail.AsicNo) + "']" + "[name='" + fmt.Sprintf("%v", tcamDetail.Name) + "']"
    tcamDetail.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    tcamDetail.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    tcamDetail.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    tcamDetail.EntityData.Children = make(map[string]types.YChild)
    tcamDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    tcamDetail.EntityData.Leafs["asic-no"] = types.YLeaf{"AsicNo", tcamDetail.AsicNo}
    tcamDetail.EntityData.Leafs["name"] = types.YLeaf{"Name", tcamDetail.Name}
    tcamDetail.EntityData.Leafs["hash-entries-max"] = types.YLeaf{"HashEntriesMax", tcamDetail.HashEntriesMax}
    tcamDetail.EntityData.Leafs["tcam-entries-max"] = types.YLeaf{"TcamEntriesMax", tcamDetail.TcamEntriesMax}
    tcamDetail.EntityData.Leafs["hash-entries-used"] = types.YLeaf{"HashEntriesUsed", tcamDetail.HashEntriesUsed}
    tcamDetail.EntityData.Leafs["tcam-entries-used"] = types.YLeaf{"TcamEntriesUsed", tcamDetail.TcamEntriesUsed}
    return &(tcamDetail.EntityData)
}

