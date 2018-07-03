// This module contains a collection of YANG definitions
// for Cisco IOS-XR dnx-driver-fabric-plane package operational data.
// 
// This module contains definitions
// for the following management objects:
//   fabric: Admin fabric oper data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package dnx_driver_fabric_plane_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package dnx_driver_fabric_plane_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-dnx-driver-fabric-plane-oper fabric}", reflect.TypeOf(Fabric{}))
    ydk.RegisterEntity("Cisco-IOS-XR-dnx-driver-fabric-plane-oper:fabric", reflect.TypeOf(Fabric{}))
}

// FsdbPlaneMode represents FSDB plane mode enum
type FsdbPlaneMode string

const (
    // plane mode unknown
    FsdbPlaneMode_plane_mode_unknown FsdbPlaneMode = "plane-mode-unknown"

    // plane mode sc
    FsdbPlaneMode_plane_mode_sc FsdbPlaneMode = "plane-mode-sc"

    // plane mode b2b
    FsdbPlaneMode_plane_mode_b2b FsdbPlaneMode = "plane-mode-b2b"

    // plane mode mc
    FsdbPlaneMode_plane_mode_mc FsdbPlaneMode = "plane-mode-mc"

    // plane mode folded
    FsdbPlaneMode_plane_mode_folded FsdbPlaneMode = "plane-mode-folded"
)

// FsdbPlaneAdminState represents FSDB Plane admin state enum
type FsdbPlaneAdminState string

const (
    // plane state admin up
    FsdbPlaneAdminState_plane_state_admin_up FsdbPlaneAdminState = "plane-state-admin-up"

    // plane state admin down
    FsdbPlaneAdminState_plane_state_admin_down FsdbPlaneAdminState = "plane-state-admin-down"
)

// FsdbPlaneOperState represents FSDB plane oper state enum
type FsdbPlaneOperState string

const (
    // plane state oper up
    FsdbPlaneOperState_plane_state_oper_up FsdbPlaneOperState = "plane-state-oper-up"

    // plane state oper down
    FsdbPlaneOperState_plane_state_oper_down FsdbPlaneOperState = "plane-state-oper-down"

    // plane state oper mcast down
    FsdbPlaneOperState_plane_state_oper_mcast_down FsdbPlaneOperState = "plane-state-oper-mcast-down"
)

// Fabric
// Admin fabric oper data
type Fabric struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Plane state table for Fabric .
    PlaneTable Fabric_PlaneTable
}

func (fabric *Fabric) GetEntityData() *types.CommonEntityData {
    fabric.EntityData.YFilter = fabric.YFilter
    fabric.EntityData.YangName = "fabric"
    fabric.EntityData.BundleName = "cisco_ios_xr"
    fabric.EntityData.ParentYangName = "Cisco-IOS-XR-dnx-driver-fabric-plane-oper"
    fabric.EntityData.SegmentPath = "Cisco-IOS-XR-dnx-driver-fabric-plane-oper:fabric"
    fabric.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fabric.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fabric.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fabric.EntityData.Children = types.NewOrderedMap()
    fabric.EntityData.Children.Append("plane-table", types.YChild{"PlaneTable", &fabric.PlaneTable})
    fabric.EntityData.Leafs = types.NewOrderedMap()

    fabric.EntityData.YListKeys = []string {}

    return &(fabric.EntityData)
}

// Fabric_PlaneTable
// Plane state table for Fabric 
type Fabric_PlaneTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Show Calvados Plane Statistics.
    Statistics Fabric_PlaneTable_Statistics

    // Show Calvados Plane State.
    Plane Fabric_PlaneTable_Plane
}

func (planeTable *Fabric_PlaneTable) GetEntityData() *types.CommonEntityData {
    planeTable.EntityData.YFilter = planeTable.YFilter
    planeTable.EntityData.YangName = "plane-table"
    planeTable.EntityData.BundleName = "cisco_ios_xr"
    planeTable.EntityData.ParentYangName = "fabric"
    planeTable.EntityData.SegmentPath = "plane-table"
    planeTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    planeTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    planeTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    planeTable.EntityData.Children = types.NewOrderedMap()
    planeTable.EntityData.Children.Append("statistics", types.YChild{"Statistics", &planeTable.Statistics})
    planeTable.EntityData.Children.Append("plane", types.YChild{"Plane", &planeTable.Plane})
    planeTable.EntityData.Leafs = types.NewOrderedMap()

    planeTable.EntityData.YListKeys = []string {}

    return &(planeTable.EntityData)
}

// Fabric_PlaneTable_Statistics
// Show Calvados Plane Statistics
type Fabric_PlaneTable_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // plane stats info. The type is slice of
    // Fabric_PlaneTable_Statistics_PlaneStatsInfo.
    PlaneStatsInfo []*Fabric_PlaneTable_Statistics_PlaneStatsInfo
}

func (statistics *Fabric_PlaneTable_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "plane-table"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = types.NewOrderedMap()
    statistics.EntityData.Children.Append("plane-stats-info", types.YChild{"PlaneStatsInfo", nil})
    for i := range statistics.PlaneStatsInfo {
        statistics.EntityData.Children.Append(types.GetSegmentPath(statistics.PlaneStatsInfo[i]), types.YChild{"PlaneStatsInfo", statistics.PlaneStatsInfo[i]})
    }
    statistics.EntityData.Leafs = types.NewOrderedMap()

    statistics.EntityData.YListKeys = []string {}

    return &(statistics.EntityData)
}

// Fabric_PlaneTable_Statistics_PlaneStatsInfo
// plane stats info
type Fabric_PlaneTable_Statistics_PlaneStatsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // plane id. The type is interface{} with range: 0..4294967295.
    PlaneId interface{}

    // RxDataCells. The type is interface{} with range: 0..18446744073709551615.
    RxDataCells interface{}

    // TxDataCells. The type is interface{} with range: 0..18446744073709551615.
    TxDataCells interface{}

    // RxCorrectableErrorCells. The type is interface{} with range: 0..4294967295.
    RxCorrectableErrorCells interface{}

    // RxUnCorrectableErrorCells. The type is interface{} with range:
    // 0..4294967295.
    RxUnCorrectableErrorCells interface{}

    // RxParityErrorCells. The type is interface{} with range: 0..4294967295.
    RxParityErrorCells interface{}
}

func (planeStatsInfo *Fabric_PlaneTable_Statistics_PlaneStatsInfo) GetEntityData() *types.CommonEntityData {
    planeStatsInfo.EntityData.YFilter = planeStatsInfo.YFilter
    planeStatsInfo.EntityData.YangName = "plane-stats-info"
    planeStatsInfo.EntityData.BundleName = "cisco_ios_xr"
    planeStatsInfo.EntityData.ParentYangName = "statistics"
    planeStatsInfo.EntityData.SegmentPath = "plane-stats-info"
    planeStatsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    planeStatsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    planeStatsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    planeStatsInfo.EntityData.Children = types.NewOrderedMap()
    planeStatsInfo.EntityData.Leafs = types.NewOrderedMap()
    planeStatsInfo.EntityData.Leafs.Append("plane-id", types.YLeaf{"PlaneId", planeStatsInfo.PlaneId})
    planeStatsInfo.EntityData.Leafs.Append("rx-data-cells", types.YLeaf{"RxDataCells", planeStatsInfo.RxDataCells})
    planeStatsInfo.EntityData.Leafs.Append("tx-data-cells", types.YLeaf{"TxDataCells", planeStatsInfo.TxDataCells})
    planeStatsInfo.EntityData.Leafs.Append("rx-correctable-error-cells", types.YLeaf{"RxCorrectableErrorCells", planeStatsInfo.RxCorrectableErrorCells})
    planeStatsInfo.EntityData.Leafs.Append("rx-un-correctable-error-cells", types.YLeaf{"RxUnCorrectableErrorCells", planeStatsInfo.RxUnCorrectableErrorCells})
    planeStatsInfo.EntityData.Leafs.Append("rx-parity-error-cells", types.YLeaf{"RxParityErrorCells", planeStatsInfo.RxParityErrorCells})

    planeStatsInfo.EntityData.YListKeys = []string {}

    return &(planeStatsInfo.EntityData)
}

// Fabric_PlaneTable_Plane
// Show Calvados Plane State
type Fabric_PlaneTable_Plane struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // detail plane info. The type is slice of
    // Fabric_PlaneTable_Plane_DetailPlaneInfo.
    DetailPlaneInfo []*Fabric_PlaneTable_Plane_DetailPlaneInfo
}

func (plane *Fabric_PlaneTable_Plane) GetEntityData() *types.CommonEntityData {
    plane.EntityData.YFilter = plane.YFilter
    plane.EntityData.YangName = "plane"
    plane.EntityData.BundleName = "cisco_ios_xr"
    plane.EntityData.ParentYangName = "plane-table"
    plane.EntityData.SegmentPath = "plane"
    plane.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    plane.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    plane.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    plane.EntityData.Children = types.NewOrderedMap()
    plane.EntityData.Children.Append("detail-plane-info", types.YChild{"DetailPlaneInfo", nil})
    for i := range plane.DetailPlaneInfo {
        plane.EntityData.Children.Append(types.GetSegmentPath(plane.DetailPlaneInfo[i]), types.YChild{"DetailPlaneInfo", plane.DetailPlaneInfo[i]})
    }
    plane.EntityData.Leafs = types.NewOrderedMap()

    plane.EntityData.YListKeys = []string {}

    return &(plane.EntityData)
}

// Fabric_PlaneTable_Plane_DetailPlaneInfo
// detail plane info
type Fabric_PlaneTable_Plane_DetailPlaneInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // plane id. The type is interface{} with range: 0..4294967295.
    PlaneId interface{}

    // PlaneOperStatus. The type is FsdbPlaneOperState.
    PlaneOperStatus interface{}

    // PlaneAdminStatus. The type is FsdbPlaneAdminState.
    PlaneAdminStatus interface{}

    // Plane Mode Configuration. The type is FsdbPlaneMode.
    PlaneMode interface{}

    // Total number of bundles. The type is interface{} with range: 0..65535.
    Bundles interface{}

    // Total down bundles. The type is interface{} with range: 0..65535.
    DownBundles interface{}

    // Plane up down count. The type is interface{} with range: 0..4294967295.
    PlaneUpDownCount interface{}

    // Plane up multicast count. The type is interface{} with range:
    // 0..4294967295.
    UpMulticastCount interface{}

    // Plane PPU State. The type is string.
    PpuState interface{}
}

func (detailPlaneInfo *Fabric_PlaneTable_Plane_DetailPlaneInfo) GetEntityData() *types.CommonEntityData {
    detailPlaneInfo.EntityData.YFilter = detailPlaneInfo.YFilter
    detailPlaneInfo.EntityData.YangName = "detail-plane-info"
    detailPlaneInfo.EntityData.BundleName = "cisco_ios_xr"
    detailPlaneInfo.EntityData.ParentYangName = "plane"
    detailPlaneInfo.EntityData.SegmentPath = "detail-plane-info"
    detailPlaneInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detailPlaneInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detailPlaneInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detailPlaneInfo.EntityData.Children = types.NewOrderedMap()
    detailPlaneInfo.EntityData.Leafs = types.NewOrderedMap()
    detailPlaneInfo.EntityData.Leafs.Append("plane-id", types.YLeaf{"PlaneId", detailPlaneInfo.PlaneId})
    detailPlaneInfo.EntityData.Leafs.Append("plane-oper-status", types.YLeaf{"PlaneOperStatus", detailPlaneInfo.PlaneOperStatus})
    detailPlaneInfo.EntityData.Leafs.Append("plane-admin-status", types.YLeaf{"PlaneAdminStatus", detailPlaneInfo.PlaneAdminStatus})
    detailPlaneInfo.EntityData.Leafs.Append("plane-mode", types.YLeaf{"PlaneMode", detailPlaneInfo.PlaneMode})
    detailPlaneInfo.EntityData.Leafs.Append("bundles", types.YLeaf{"Bundles", detailPlaneInfo.Bundles})
    detailPlaneInfo.EntityData.Leafs.Append("down-bundles", types.YLeaf{"DownBundles", detailPlaneInfo.DownBundles})
    detailPlaneInfo.EntityData.Leafs.Append("plane-up-down-count", types.YLeaf{"PlaneUpDownCount", detailPlaneInfo.PlaneUpDownCount})
    detailPlaneInfo.EntityData.Leafs.Append("up-multicast-count", types.YLeaf{"UpMulticastCount", detailPlaneInfo.UpMulticastCount})
    detailPlaneInfo.EntityData.Leafs.Append("ppu-state", types.YLeaf{"PpuState", detailPlaneInfo.PpuState})

    detailPlaneInfo.EntityData.YListKeys = []string {}

    return &(detailPlaneInfo.EntityData)
}

