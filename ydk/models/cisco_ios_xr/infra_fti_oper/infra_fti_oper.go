// This module contains a collection of YANG definitions
// for Cisco IOS-XR infra-fti package operational data.
// 
// This module contains definitions
// for the following management objects:
//   dci-fabric-interconnect: Display FTI operational data
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package infra_fti_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package infra_fti_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-infra-fti-oper dci-fabric-interconnect}", reflect.TypeOf(DciFabricInterconnect{}))
    ydk.RegisterEntity("Cisco-IOS-XR-infra-fti-oper:dci-fabric-interconnect", reflect.TypeOf(DciFabricInterconnect{}))
}

// FtiBagFabricPeerState represents FTI Fabric Peer States
type FtiBagFabricPeerState string

const (
    // Disconnected
    FtiBagFabricPeerState_fti_bag_fabric_peer_status_disconnected FtiBagFabricPeerState = "fti-bag-fabric-peer-status-disconnected"

    // Connecting
    FtiBagFabricPeerState_fti_bag_fabric_peer_status_connecting FtiBagFabricPeerState = "fti-bag-fabric-peer-status-connecting"

    // Connected
    FtiBagFabricPeerState_fti_bag_fabric_peer_status_connected FtiBagFabricPeerState = "fti-bag-fabric-peer-status-connected"

    // Ready
    FtiBagFabricPeerState_fti_bag_fabric_peer_status_ready FtiBagFabricPeerState = "fti-bag-fabric-peer-status-ready"

    // Closing
    FtiBagFabricPeerState_fti_bag_fabric_peer_status_closing FtiBagFabricPeerState = "fti-bag-fabric-peer-status-closing"
)

// FtiBagFabricState represents FTI Fabric States
type FtiBagFabricState string

const (
    // Active (Down)
    FtiBagFabricState_fti_bag_fabric_state_active_down FtiBagFabricState = "fti-bag-fabric-state-active-down"

    // Active (Degraded)
    FtiBagFabricState_fti_bag_fabric_state_active_degraded FtiBagFabricState = "fti-bag-fabric-state-active-degraded"

    // Active (Healthy)
    FtiBagFabricState_fti_bag_fabric_state_active_healthy FtiBagFabricState = "fti-bag-fabric-state-active-healthy"

    // Inactive
    FtiBagFabricState_fti_bag_fabric_state_inactive FtiBagFabricState = "fti-bag-fabric-state-inactive"
)

// FtiBagFabricConfigState represents FTI Fabric Config States
type FtiBagFabricConfigState string

const (
    // Config Complete
    FtiBagFabricConfigState_fti_bag_config_complete FtiBagFabricConfigState = "fti-bag-config-complete"

    // Config Incomplete
    FtiBagFabricConfigState_fti_bag_config_incomplete FtiBagFabricConfigState = "fti-bag-config-incomplete"
)

// DciFabricInterconnect
// Display FTI operational data
type DciFabricInterconnect struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // FTI Opflex Session Info for all fabrics.
    OpflexSessionInfos DciFabricInterconnect_OpflexSessionInfos

    // FTI Fabric-VRF DB for all fabrics.
    FabricVrfDbs DciFabricInterconnect_FabricVrfDbs

    // FTI DCI-VRF DB for all VRFs.
    DciVrfs DciFabricInterconnect_DciVrfs

    // Auto Config Pool Info.
    Acp DciFabricInterconnect_Acp
}

func (dciFabricInterconnect *DciFabricInterconnect) GetEntityData() *types.CommonEntityData {
    dciFabricInterconnect.EntityData.YFilter = dciFabricInterconnect.YFilter
    dciFabricInterconnect.EntityData.YangName = "dci-fabric-interconnect"
    dciFabricInterconnect.EntityData.BundleName = "cisco_ios_xr"
    dciFabricInterconnect.EntityData.ParentYangName = "Cisco-IOS-XR-infra-fti-oper"
    dciFabricInterconnect.EntityData.SegmentPath = "Cisco-IOS-XR-infra-fti-oper:dci-fabric-interconnect"
    dciFabricInterconnect.EntityData.AbsolutePath = dciFabricInterconnect.EntityData.SegmentPath
    dciFabricInterconnect.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dciFabricInterconnect.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dciFabricInterconnect.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dciFabricInterconnect.EntityData.Children = types.NewOrderedMap()
    dciFabricInterconnect.EntityData.Children.Append("opflex-session-infos", types.YChild{"OpflexSessionInfos", &dciFabricInterconnect.OpflexSessionInfos})
    dciFabricInterconnect.EntityData.Children.Append("fabric-vrf-dbs", types.YChild{"FabricVrfDbs", &dciFabricInterconnect.FabricVrfDbs})
    dciFabricInterconnect.EntityData.Children.Append("dci-vrfs", types.YChild{"DciVrfs", &dciFabricInterconnect.DciVrfs})
    dciFabricInterconnect.EntityData.Children.Append("acp", types.YChild{"Acp", &dciFabricInterconnect.Acp})
    dciFabricInterconnect.EntityData.Leafs = types.NewOrderedMap()

    dciFabricInterconnect.EntityData.YListKeys = []string {}

    return &(dciFabricInterconnect.EntityData)
}

// DciFabricInterconnect_OpflexSessionInfos
// FTI Opflex Session Info for all fabrics
type DciFabricInterconnect_OpflexSessionInfos struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // FTI Opflex Session Info for a particular fabric. The type is slice of
    // DciFabricInterconnect_OpflexSessionInfos_OpflexSessionInfo.
    OpflexSessionInfo []*DciFabricInterconnect_OpflexSessionInfos_OpflexSessionInfo
}

func (opflexSessionInfos *DciFabricInterconnect_OpflexSessionInfos) GetEntityData() *types.CommonEntityData {
    opflexSessionInfos.EntityData.YFilter = opflexSessionInfos.YFilter
    opflexSessionInfos.EntityData.YangName = "opflex-session-infos"
    opflexSessionInfos.EntityData.BundleName = "cisco_ios_xr"
    opflexSessionInfos.EntityData.ParentYangName = "dci-fabric-interconnect"
    opflexSessionInfos.EntityData.SegmentPath = "opflex-session-infos"
    opflexSessionInfos.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-fti-oper:dci-fabric-interconnect/" + opflexSessionInfos.EntityData.SegmentPath
    opflexSessionInfos.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    opflexSessionInfos.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    opflexSessionInfos.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    opflexSessionInfos.EntityData.Children = types.NewOrderedMap()
    opflexSessionInfos.EntityData.Children.Append("opflex-session-info", types.YChild{"OpflexSessionInfo", nil})
    for i := range opflexSessionInfos.OpflexSessionInfo {
        opflexSessionInfos.EntityData.Children.Append(types.GetSegmentPath(opflexSessionInfos.OpflexSessionInfo[i]), types.YChild{"OpflexSessionInfo", opflexSessionInfos.OpflexSessionInfo[i]})
    }
    opflexSessionInfos.EntityData.Leafs = types.NewOrderedMap()

    opflexSessionInfos.EntityData.YListKeys = []string {}

    return &(opflexSessionInfos.EntityData)
}

// DciFabricInterconnect_OpflexSessionInfos_OpflexSessionInfo
// FTI Opflex Session Info for a particular fabric
type DciFabricInterconnect_OpflexSessionInfos_OpflexSessionInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. fabric identifier. The type is interface{} with
    // range: 1000..9999.
    Id1 interface{}

    // Fabric Id. The type is interface{} with range: 0..4294967295.
    FabricId interface{}

    // Config State. The type is FtiBagFabricConfigState.
    ConfigState interface{}

    // Config Last Update Timestamp. The type is interface{} with range:
    // 0..18446744073709551615.
    LastUpdTsConfig interface{}

    // Fabric State. The type is FtiBagFabricState.
    FabricState interface{}

    // Fabric Last Update Timestamp. The type is interface{} with range:
    // 0..18446744073709551615.
    LastUpdTsFabric interface{}

    // Fabric Per Peer Info. The type is slice of
    // DciFabricInterconnect_OpflexSessionInfos_OpflexSessionInfo_PeerInfo.
    PeerInfo []*DciFabricInterconnect_OpflexSessionInfos_OpflexSessionInfo_PeerInfo
}

func (opflexSessionInfo *DciFabricInterconnect_OpflexSessionInfos_OpflexSessionInfo) GetEntityData() *types.CommonEntityData {
    opflexSessionInfo.EntityData.YFilter = opflexSessionInfo.YFilter
    opflexSessionInfo.EntityData.YangName = "opflex-session-info"
    opflexSessionInfo.EntityData.BundleName = "cisco_ios_xr"
    opflexSessionInfo.EntityData.ParentYangName = "opflex-session-infos"
    opflexSessionInfo.EntityData.SegmentPath = "opflex-session-info" + types.AddKeyToken(opflexSessionInfo.Id1, "id1")
    opflexSessionInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-fti-oper:dci-fabric-interconnect/opflex-session-infos/" + opflexSessionInfo.EntityData.SegmentPath
    opflexSessionInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    opflexSessionInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    opflexSessionInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    opflexSessionInfo.EntityData.Children = types.NewOrderedMap()
    opflexSessionInfo.EntityData.Children.Append("peer-info", types.YChild{"PeerInfo", nil})
    for i := range opflexSessionInfo.PeerInfo {
        types.SetYListKey(opflexSessionInfo.PeerInfo[i], i)
        opflexSessionInfo.EntityData.Children.Append(types.GetSegmentPath(opflexSessionInfo.PeerInfo[i]), types.YChild{"PeerInfo", opflexSessionInfo.PeerInfo[i]})
    }
    opflexSessionInfo.EntityData.Leafs = types.NewOrderedMap()
    opflexSessionInfo.EntityData.Leafs.Append("id1", types.YLeaf{"Id1", opflexSessionInfo.Id1})
    opflexSessionInfo.EntityData.Leafs.Append("fabric-id", types.YLeaf{"FabricId", opflexSessionInfo.FabricId})
    opflexSessionInfo.EntityData.Leafs.Append("config-state", types.YLeaf{"ConfigState", opflexSessionInfo.ConfigState})
    opflexSessionInfo.EntityData.Leafs.Append("last-upd-ts-config", types.YLeaf{"LastUpdTsConfig", opflexSessionInfo.LastUpdTsConfig})
    opflexSessionInfo.EntityData.Leafs.Append("fabric-state", types.YLeaf{"FabricState", opflexSessionInfo.FabricState})
    opflexSessionInfo.EntityData.Leafs.Append("last-upd-ts-fabric", types.YLeaf{"LastUpdTsFabric", opflexSessionInfo.LastUpdTsFabric})

    opflexSessionInfo.EntityData.YListKeys = []string {"Id1"}

    return &(opflexSessionInfo.EntityData)
}

// DciFabricInterconnect_OpflexSessionInfos_OpflexSessionInfo_PeerInfo
// Fabric Per Peer Info
type DciFabricInterconnect_OpflexSessionInfos_OpflexSessionInfo_PeerInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Peer IP. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    PeerIp interface{}

    // Peer Port. The type is interface{} with range: 0..4294967295.
    PeerPort interface{}

    // Peer State. The type is FtiBagFabricPeerState.
    PeerState interface{}

    // Peer Status Last Update Timestamp. The type is interface{} with range:
    // 0..18446744073709551615.
    LastUpdTsPeerStatus interface{}
}

func (peerInfo *DciFabricInterconnect_OpflexSessionInfos_OpflexSessionInfo_PeerInfo) GetEntityData() *types.CommonEntityData {
    peerInfo.EntityData.YFilter = peerInfo.YFilter
    peerInfo.EntityData.YangName = "peer-info"
    peerInfo.EntityData.BundleName = "cisco_ios_xr"
    peerInfo.EntityData.ParentYangName = "opflex-session-info"
    peerInfo.EntityData.SegmentPath = "peer-info" + types.AddNoKeyToken(peerInfo)
    peerInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-fti-oper:dci-fabric-interconnect/opflex-session-infos/opflex-session-info/" + peerInfo.EntityData.SegmentPath
    peerInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerInfo.EntityData.Children = types.NewOrderedMap()
    peerInfo.EntityData.Leafs = types.NewOrderedMap()
    peerInfo.EntityData.Leafs.Append("peer-ip", types.YLeaf{"PeerIp", peerInfo.PeerIp})
    peerInfo.EntityData.Leafs.Append("peer-port", types.YLeaf{"PeerPort", peerInfo.PeerPort})
    peerInfo.EntityData.Leafs.Append("peer-state", types.YLeaf{"PeerState", peerInfo.PeerState})
    peerInfo.EntityData.Leafs.Append("last-upd-ts-peer-status", types.YLeaf{"LastUpdTsPeerStatus", peerInfo.LastUpdTsPeerStatus})

    peerInfo.EntityData.YListKeys = []string {}

    return &(peerInfo.EntityData)
}

// DciFabricInterconnect_FabricVrfDbs
// FTI Fabric-VRF DB for all fabrics
type DciFabricInterconnect_FabricVrfDbs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // FTI Fabric-VRF DB for a particular fabric. The type is slice of
    // DciFabricInterconnect_FabricVrfDbs_FabricVrfDb.
    FabricVrfDb []*DciFabricInterconnect_FabricVrfDbs_FabricVrfDb
}

func (fabricVrfDbs *DciFabricInterconnect_FabricVrfDbs) GetEntityData() *types.CommonEntityData {
    fabricVrfDbs.EntityData.YFilter = fabricVrfDbs.YFilter
    fabricVrfDbs.EntityData.YangName = "fabric-vrf-dbs"
    fabricVrfDbs.EntityData.BundleName = "cisco_ios_xr"
    fabricVrfDbs.EntityData.ParentYangName = "dci-fabric-interconnect"
    fabricVrfDbs.EntityData.SegmentPath = "fabric-vrf-dbs"
    fabricVrfDbs.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-fti-oper:dci-fabric-interconnect/" + fabricVrfDbs.EntityData.SegmentPath
    fabricVrfDbs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fabricVrfDbs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fabricVrfDbs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fabricVrfDbs.EntityData.Children = types.NewOrderedMap()
    fabricVrfDbs.EntityData.Children.Append("fabric-vrf-db", types.YChild{"FabricVrfDb", nil})
    for i := range fabricVrfDbs.FabricVrfDb {
        fabricVrfDbs.EntityData.Children.Append(types.GetSegmentPath(fabricVrfDbs.FabricVrfDb[i]), types.YChild{"FabricVrfDb", fabricVrfDbs.FabricVrfDb[i]})
    }
    fabricVrfDbs.EntityData.Leafs = types.NewOrderedMap()

    fabricVrfDbs.EntityData.YListKeys = []string {}

    return &(fabricVrfDbs.EntityData)
}

// DciFabricInterconnect_FabricVrfDbs_FabricVrfDb
// FTI Fabric-VRF DB for a particular fabric
type DciFabricInterconnect_FabricVrfDbs_FabricVrfDb struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. fabric identifier. The type is interface{} with
    // range: 1000..9999.
    Id1 interface{}

    // Fabric Id. The type is interface{} with range: 0..4294967295.
    FabricId interface{}

    // Fabric VRFs. The type is slice of
    // DciFabricInterconnect_FabricVrfDbs_FabricVrfDb_FabricVrf.
    FabricVrf []*DciFabricInterconnect_FabricVrfDbs_FabricVrfDb_FabricVrf
}

func (fabricVrfDb *DciFabricInterconnect_FabricVrfDbs_FabricVrfDb) GetEntityData() *types.CommonEntityData {
    fabricVrfDb.EntityData.YFilter = fabricVrfDb.YFilter
    fabricVrfDb.EntityData.YangName = "fabric-vrf-db"
    fabricVrfDb.EntityData.BundleName = "cisco_ios_xr"
    fabricVrfDb.EntityData.ParentYangName = "fabric-vrf-dbs"
    fabricVrfDb.EntityData.SegmentPath = "fabric-vrf-db" + types.AddKeyToken(fabricVrfDb.Id1, "id1")
    fabricVrfDb.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-fti-oper:dci-fabric-interconnect/fabric-vrf-dbs/" + fabricVrfDb.EntityData.SegmentPath
    fabricVrfDb.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fabricVrfDb.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fabricVrfDb.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fabricVrfDb.EntityData.Children = types.NewOrderedMap()
    fabricVrfDb.EntityData.Children.Append("fabric-vrf", types.YChild{"FabricVrf", nil})
    for i := range fabricVrfDb.FabricVrf {
        types.SetYListKey(fabricVrfDb.FabricVrf[i], i)
        fabricVrfDb.EntityData.Children.Append(types.GetSegmentPath(fabricVrfDb.FabricVrf[i]), types.YChild{"FabricVrf", fabricVrfDb.FabricVrf[i]})
    }
    fabricVrfDb.EntityData.Leafs = types.NewOrderedMap()
    fabricVrfDb.EntityData.Leafs.Append("id1", types.YLeaf{"Id1", fabricVrfDb.Id1})
    fabricVrfDb.EntityData.Leafs.Append("fabric-id", types.YLeaf{"FabricId", fabricVrfDb.FabricId})

    fabricVrfDb.EntityData.YListKeys = []string {"Id1"}

    return &(fabricVrfDb.EntityData)
}

// DciFabricInterconnect_FabricVrfDbs_FabricVrfDb_FabricVrf
// Fabric VRFs
type DciFabricInterconnect_FabricVrfDbs_FabricVrfDb_FabricVrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Fabric VRF. The type is string.
    FabricVrf interface{}

    // Dci VRF. The type is string.
    DciVrf interface{}

    // Fabric VRF Flags. The type is string.
    FabricVrfFlags interface{}

    // V4 Import Route Targets. The type is slice of
    // DciFabricInterconnect_FabricVrfDbs_FabricVrfDb_FabricVrf_V4ImportRt.
    V4ImportRt []*DciFabricInterconnect_FabricVrfDbs_FabricVrfDb_FabricVrf_V4ImportRt

    // V4 Export Route Targets. The type is slice of
    // DciFabricInterconnect_FabricVrfDbs_FabricVrfDb_FabricVrf_V4ExportRt.
    V4ExportRt []*DciFabricInterconnect_FabricVrfDbs_FabricVrfDb_FabricVrf_V4ExportRt

    // V6 Import Route Targets. The type is slice of
    // DciFabricInterconnect_FabricVrfDbs_FabricVrfDb_FabricVrf_V6ImportRt.
    V6ImportRt []*DciFabricInterconnect_FabricVrfDbs_FabricVrfDb_FabricVrf_V6ImportRt

    // V6 Export Route Targets. The type is slice of
    // DciFabricInterconnect_FabricVrfDbs_FabricVrfDb_FabricVrf_V6ExportRt.
    V6ExportRt []*DciFabricInterconnect_FabricVrfDbs_FabricVrfDb_FabricVrf_V6ExportRt
}

func (fabricVrf *DciFabricInterconnect_FabricVrfDbs_FabricVrfDb_FabricVrf) GetEntityData() *types.CommonEntityData {
    fabricVrf.EntityData.YFilter = fabricVrf.YFilter
    fabricVrf.EntityData.YangName = "fabric-vrf"
    fabricVrf.EntityData.BundleName = "cisco_ios_xr"
    fabricVrf.EntityData.ParentYangName = "fabric-vrf-db"
    fabricVrf.EntityData.SegmentPath = "fabric-vrf" + types.AddNoKeyToken(fabricVrf)
    fabricVrf.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-fti-oper:dci-fabric-interconnect/fabric-vrf-dbs/fabric-vrf-db/" + fabricVrf.EntityData.SegmentPath
    fabricVrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fabricVrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fabricVrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fabricVrf.EntityData.Children = types.NewOrderedMap()
    fabricVrf.EntityData.Children.Append("v4-import-rt", types.YChild{"V4ImportRt", nil})
    for i := range fabricVrf.V4ImportRt {
        types.SetYListKey(fabricVrf.V4ImportRt[i], i)
        fabricVrf.EntityData.Children.Append(types.GetSegmentPath(fabricVrf.V4ImportRt[i]), types.YChild{"V4ImportRt", fabricVrf.V4ImportRt[i]})
    }
    fabricVrf.EntityData.Children.Append("v4-export-rt", types.YChild{"V4ExportRt", nil})
    for i := range fabricVrf.V4ExportRt {
        types.SetYListKey(fabricVrf.V4ExportRt[i], i)
        fabricVrf.EntityData.Children.Append(types.GetSegmentPath(fabricVrf.V4ExportRt[i]), types.YChild{"V4ExportRt", fabricVrf.V4ExportRt[i]})
    }
    fabricVrf.EntityData.Children.Append("v6-import-rt", types.YChild{"V6ImportRt", nil})
    for i := range fabricVrf.V6ImportRt {
        types.SetYListKey(fabricVrf.V6ImportRt[i], i)
        fabricVrf.EntityData.Children.Append(types.GetSegmentPath(fabricVrf.V6ImportRt[i]), types.YChild{"V6ImportRt", fabricVrf.V6ImportRt[i]})
    }
    fabricVrf.EntityData.Children.Append("v6-export-rt", types.YChild{"V6ExportRt", nil})
    for i := range fabricVrf.V6ExportRt {
        types.SetYListKey(fabricVrf.V6ExportRt[i], i)
        fabricVrf.EntityData.Children.Append(types.GetSegmentPath(fabricVrf.V6ExportRt[i]), types.YChild{"V6ExportRt", fabricVrf.V6ExportRt[i]})
    }
    fabricVrf.EntityData.Leafs = types.NewOrderedMap()
    fabricVrf.EntityData.Leafs.Append("fabric-vrf", types.YLeaf{"FabricVrf", fabricVrf.FabricVrf})
    fabricVrf.EntityData.Leafs.Append("dci-vrf", types.YLeaf{"DciVrf", fabricVrf.DciVrf})
    fabricVrf.EntityData.Leafs.Append("fabric-vrf-flags", types.YLeaf{"FabricVrfFlags", fabricVrf.FabricVrfFlags})

    fabricVrf.EntityData.YListKeys = []string {}

    return &(fabricVrf.EntityData)
}

// DciFabricInterconnect_FabricVrfDbs_FabricVrfDb_FabricVrf_V4ImportRt
// V4 Import Route Targets
type DciFabricInterconnect_FabricVrfDbs_FabricVrfDb_FabricVrf_V4ImportRt struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // RT Value. The type is string.
    RtValue interface{}

    // RT Flags. The type is string.
    RtFlags interface{}
}

func (v4ImportRt *DciFabricInterconnect_FabricVrfDbs_FabricVrfDb_FabricVrf_V4ImportRt) GetEntityData() *types.CommonEntityData {
    v4ImportRt.EntityData.YFilter = v4ImportRt.YFilter
    v4ImportRt.EntityData.YangName = "v4-import-rt"
    v4ImportRt.EntityData.BundleName = "cisco_ios_xr"
    v4ImportRt.EntityData.ParentYangName = "fabric-vrf"
    v4ImportRt.EntityData.SegmentPath = "v4-import-rt" + types.AddNoKeyToken(v4ImportRt)
    v4ImportRt.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-fti-oper:dci-fabric-interconnect/fabric-vrf-dbs/fabric-vrf-db/fabric-vrf/" + v4ImportRt.EntityData.SegmentPath
    v4ImportRt.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    v4ImportRt.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    v4ImportRt.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    v4ImportRt.EntityData.Children = types.NewOrderedMap()
    v4ImportRt.EntityData.Leafs = types.NewOrderedMap()
    v4ImportRt.EntityData.Leafs.Append("rt-value", types.YLeaf{"RtValue", v4ImportRt.RtValue})
    v4ImportRt.EntityData.Leafs.Append("rt-flags", types.YLeaf{"RtFlags", v4ImportRt.RtFlags})

    v4ImportRt.EntityData.YListKeys = []string {}

    return &(v4ImportRt.EntityData)
}

// DciFabricInterconnect_FabricVrfDbs_FabricVrfDb_FabricVrf_V4ExportRt
// V4 Export Route Targets
type DciFabricInterconnect_FabricVrfDbs_FabricVrfDb_FabricVrf_V4ExportRt struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // RT Value. The type is string.
    RtValue interface{}

    // RT Flags. The type is string.
    RtFlags interface{}
}

func (v4ExportRt *DciFabricInterconnect_FabricVrfDbs_FabricVrfDb_FabricVrf_V4ExportRt) GetEntityData() *types.CommonEntityData {
    v4ExportRt.EntityData.YFilter = v4ExportRt.YFilter
    v4ExportRt.EntityData.YangName = "v4-export-rt"
    v4ExportRt.EntityData.BundleName = "cisco_ios_xr"
    v4ExportRt.EntityData.ParentYangName = "fabric-vrf"
    v4ExportRt.EntityData.SegmentPath = "v4-export-rt" + types.AddNoKeyToken(v4ExportRt)
    v4ExportRt.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-fti-oper:dci-fabric-interconnect/fabric-vrf-dbs/fabric-vrf-db/fabric-vrf/" + v4ExportRt.EntityData.SegmentPath
    v4ExportRt.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    v4ExportRt.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    v4ExportRt.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    v4ExportRt.EntityData.Children = types.NewOrderedMap()
    v4ExportRt.EntityData.Leafs = types.NewOrderedMap()
    v4ExportRt.EntityData.Leafs.Append("rt-value", types.YLeaf{"RtValue", v4ExportRt.RtValue})
    v4ExportRt.EntityData.Leafs.Append("rt-flags", types.YLeaf{"RtFlags", v4ExportRt.RtFlags})

    v4ExportRt.EntityData.YListKeys = []string {}

    return &(v4ExportRt.EntityData)
}

// DciFabricInterconnect_FabricVrfDbs_FabricVrfDb_FabricVrf_V6ImportRt
// V6 Import Route Targets
type DciFabricInterconnect_FabricVrfDbs_FabricVrfDb_FabricVrf_V6ImportRt struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // RT Value. The type is string.
    RtValue interface{}

    // RT Flags. The type is string.
    RtFlags interface{}
}

func (v6ImportRt *DciFabricInterconnect_FabricVrfDbs_FabricVrfDb_FabricVrf_V6ImportRt) GetEntityData() *types.CommonEntityData {
    v6ImportRt.EntityData.YFilter = v6ImportRt.YFilter
    v6ImportRt.EntityData.YangName = "v6-import-rt"
    v6ImportRt.EntityData.BundleName = "cisco_ios_xr"
    v6ImportRt.EntityData.ParentYangName = "fabric-vrf"
    v6ImportRt.EntityData.SegmentPath = "v6-import-rt" + types.AddNoKeyToken(v6ImportRt)
    v6ImportRt.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-fti-oper:dci-fabric-interconnect/fabric-vrf-dbs/fabric-vrf-db/fabric-vrf/" + v6ImportRt.EntityData.SegmentPath
    v6ImportRt.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    v6ImportRt.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    v6ImportRt.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    v6ImportRt.EntityData.Children = types.NewOrderedMap()
    v6ImportRt.EntityData.Leafs = types.NewOrderedMap()
    v6ImportRt.EntityData.Leafs.Append("rt-value", types.YLeaf{"RtValue", v6ImportRt.RtValue})
    v6ImportRt.EntityData.Leafs.Append("rt-flags", types.YLeaf{"RtFlags", v6ImportRt.RtFlags})

    v6ImportRt.EntityData.YListKeys = []string {}

    return &(v6ImportRt.EntityData)
}

// DciFabricInterconnect_FabricVrfDbs_FabricVrfDb_FabricVrf_V6ExportRt
// V6 Export Route Targets
type DciFabricInterconnect_FabricVrfDbs_FabricVrfDb_FabricVrf_V6ExportRt struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // RT Value. The type is string.
    RtValue interface{}

    // RT Flags. The type is string.
    RtFlags interface{}
}

func (v6ExportRt *DciFabricInterconnect_FabricVrfDbs_FabricVrfDb_FabricVrf_V6ExportRt) GetEntityData() *types.CommonEntityData {
    v6ExportRt.EntityData.YFilter = v6ExportRt.YFilter
    v6ExportRt.EntityData.YangName = "v6-export-rt"
    v6ExportRt.EntityData.BundleName = "cisco_ios_xr"
    v6ExportRt.EntityData.ParentYangName = "fabric-vrf"
    v6ExportRt.EntityData.SegmentPath = "v6-export-rt" + types.AddNoKeyToken(v6ExportRt)
    v6ExportRt.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-fti-oper:dci-fabric-interconnect/fabric-vrf-dbs/fabric-vrf-db/fabric-vrf/" + v6ExportRt.EntityData.SegmentPath
    v6ExportRt.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    v6ExportRt.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    v6ExportRt.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    v6ExportRt.EntityData.Children = types.NewOrderedMap()
    v6ExportRt.EntityData.Leafs = types.NewOrderedMap()
    v6ExportRt.EntityData.Leafs.Append("rt-value", types.YLeaf{"RtValue", v6ExportRt.RtValue})
    v6ExportRt.EntityData.Leafs.Append("rt-flags", types.YLeaf{"RtFlags", v6ExportRt.RtFlags})

    v6ExportRt.EntityData.YListKeys = []string {}

    return &(v6ExportRt.EntityData)
}

// DciFabricInterconnect_DciVrfs
// FTI DCI-VRF DB for all VRFs
type DciFabricInterconnect_DciVrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // FTI DCI-VRF Info for a particular VRF. The type is slice of
    // DciFabricInterconnect_DciVrfs_DciVrf.
    DciVrf []*DciFabricInterconnect_DciVrfs_DciVrf
}

func (dciVrfs *DciFabricInterconnect_DciVrfs) GetEntityData() *types.CommonEntityData {
    dciVrfs.EntityData.YFilter = dciVrfs.YFilter
    dciVrfs.EntityData.YangName = "dci-vrfs"
    dciVrfs.EntityData.BundleName = "cisco_ios_xr"
    dciVrfs.EntityData.ParentYangName = "dci-fabric-interconnect"
    dciVrfs.EntityData.SegmentPath = "dci-vrfs"
    dciVrfs.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-fti-oper:dci-fabric-interconnect/" + dciVrfs.EntityData.SegmentPath
    dciVrfs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dciVrfs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dciVrfs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dciVrfs.EntityData.Children = types.NewOrderedMap()
    dciVrfs.EntityData.Children.Append("dci-vrf", types.YChild{"DciVrf", nil})
    for i := range dciVrfs.DciVrf {
        dciVrfs.EntityData.Children.Append(types.GetSegmentPath(dciVrfs.DciVrf[i]), types.YChild{"DciVrf", dciVrfs.DciVrf[i]})
    }
    dciVrfs.EntityData.Leafs = types.NewOrderedMap()

    dciVrfs.EntityData.YListKeys = []string {}

    return &(dciVrfs.EntityData)
}

// DciFabricInterconnect_DciVrfs_DciVrf
// FTI DCI-VRF Info for a particular VRF
type DciFabricInterconnect_DciVrfs_DciVrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. vrf name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Vrf1 interface{}

    // DCI VRF. The type is string.
    DciVrf interface{}

    // Number of Fabric VRFs. The type is interface{} with range: 0..4294967295.
    NumFabricVrf interface{}

    // Associated Fabric Vrfs Info. The type is string.
    FaricVrfsIdName interface{}

    // VNI Id. The type is interface{} with range: 0..4294967295.
    VniId interface{}

    // BD Name. The type is string.
    BdName interface{}

    // BVI Id. The type is interface{} with range: 0..4294967295.
    BviId interface{}

    // BVI Override IP. The type is string.
    BviIp interface{}

    // BVI IPV6. False is disabled, True is enabled. The type is bool.
    BviIpV6 interface{}

    // DCI VRF Flags. The type is string.
    DciVrfFlags interface{}

    // V4 Import Route Targets. The type is slice of
    // DciFabricInterconnect_DciVrfs_DciVrf_V4ImportRt.
    V4ImportRt []*DciFabricInterconnect_DciVrfs_DciVrf_V4ImportRt

    // V4 Export Route Targets. The type is slice of
    // DciFabricInterconnect_DciVrfs_DciVrf_V4ExportRt.
    V4ExportRt []*DciFabricInterconnect_DciVrfs_DciVrf_V4ExportRt

    // V6 Import Route Targets. The type is slice of
    // DciFabricInterconnect_DciVrfs_DciVrf_V6ImportRt.
    V6ImportRt []*DciFabricInterconnect_DciVrfs_DciVrf_V6ImportRt

    // V6 Export Route Targets. The type is slice of
    // DciFabricInterconnect_DciVrfs_DciVrf_V6ExportRt.
    V6ExportRt []*DciFabricInterconnect_DciVrfs_DciVrf_V6ExportRt
}

func (dciVrf *DciFabricInterconnect_DciVrfs_DciVrf) GetEntityData() *types.CommonEntityData {
    dciVrf.EntityData.YFilter = dciVrf.YFilter
    dciVrf.EntityData.YangName = "dci-vrf"
    dciVrf.EntityData.BundleName = "cisco_ios_xr"
    dciVrf.EntityData.ParentYangName = "dci-vrfs"
    dciVrf.EntityData.SegmentPath = "dci-vrf" + types.AddKeyToken(dciVrf.Vrf1, "vrf1")
    dciVrf.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-fti-oper:dci-fabric-interconnect/dci-vrfs/" + dciVrf.EntityData.SegmentPath
    dciVrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dciVrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dciVrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dciVrf.EntityData.Children = types.NewOrderedMap()
    dciVrf.EntityData.Children.Append("v4-import-rt", types.YChild{"V4ImportRt", nil})
    for i := range dciVrf.V4ImportRt {
        types.SetYListKey(dciVrf.V4ImportRt[i], i)
        dciVrf.EntityData.Children.Append(types.GetSegmentPath(dciVrf.V4ImportRt[i]), types.YChild{"V4ImportRt", dciVrf.V4ImportRt[i]})
    }
    dciVrf.EntityData.Children.Append("v4-export-rt", types.YChild{"V4ExportRt", nil})
    for i := range dciVrf.V4ExportRt {
        types.SetYListKey(dciVrf.V4ExportRt[i], i)
        dciVrf.EntityData.Children.Append(types.GetSegmentPath(dciVrf.V4ExportRt[i]), types.YChild{"V4ExportRt", dciVrf.V4ExportRt[i]})
    }
    dciVrf.EntityData.Children.Append("v6-import-rt", types.YChild{"V6ImportRt", nil})
    for i := range dciVrf.V6ImportRt {
        types.SetYListKey(dciVrf.V6ImportRt[i], i)
        dciVrf.EntityData.Children.Append(types.GetSegmentPath(dciVrf.V6ImportRt[i]), types.YChild{"V6ImportRt", dciVrf.V6ImportRt[i]})
    }
    dciVrf.EntityData.Children.Append("v6-export-rt", types.YChild{"V6ExportRt", nil})
    for i := range dciVrf.V6ExportRt {
        types.SetYListKey(dciVrf.V6ExportRt[i], i)
        dciVrf.EntityData.Children.Append(types.GetSegmentPath(dciVrf.V6ExportRt[i]), types.YChild{"V6ExportRt", dciVrf.V6ExportRt[i]})
    }
    dciVrf.EntityData.Leafs = types.NewOrderedMap()
    dciVrf.EntityData.Leafs.Append("vrf1", types.YLeaf{"Vrf1", dciVrf.Vrf1})
    dciVrf.EntityData.Leafs.Append("dci-vrf", types.YLeaf{"DciVrf", dciVrf.DciVrf})
    dciVrf.EntityData.Leafs.Append("num-fabric-vrf", types.YLeaf{"NumFabricVrf", dciVrf.NumFabricVrf})
    dciVrf.EntityData.Leafs.Append("faric-vrfs-id-name", types.YLeaf{"FaricVrfsIdName", dciVrf.FaricVrfsIdName})
    dciVrf.EntityData.Leafs.Append("vni-id", types.YLeaf{"VniId", dciVrf.VniId})
    dciVrf.EntityData.Leafs.Append("bd-name", types.YLeaf{"BdName", dciVrf.BdName})
    dciVrf.EntityData.Leafs.Append("bvi-id", types.YLeaf{"BviId", dciVrf.BviId})
    dciVrf.EntityData.Leafs.Append("bvi-ip", types.YLeaf{"BviIp", dciVrf.BviIp})
    dciVrf.EntityData.Leafs.Append("bvi-ip-v6", types.YLeaf{"BviIpV6", dciVrf.BviIpV6})
    dciVrf.EntityData.Leafs.Append("dci-vrf-flags", types.YLeaf{"DciVrfFlags", dciVrf.DciVrfFlags})

    dciVrf.EntityData.YListKeys = []string {"Vrf1"}

    return &(dciVrf.EntityData)
}

// DciFabricInterconnect_DciVrfs_DciVrf_V4ImportRt
// V4 Import Route Targets
type DciFabricInterconnect_DciVrfs_DciVrf_V4ImportRt struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // RT Value. The type is string.
    RtValue interface{}

    // RT Flags. The type is string.
    RtFlags interface{}
}

func (v4ImportRt *DciFabricInterconnect_DciVrfs_DciVrf_V4ImportRt) GetEntityData() *types.CommonEntityData {
    v4ImportRt.EntityData.YFilter = v4ImportRt.YFilter
    v4ImportRt.EntityData.YangName = "v4-import-rt"
    v4ImportRt.EntityData.BundleName = "cisco_ios_xr"
    v4ImportRt.EntityData.ParentYangName = "dci-vrf"
    v4ImportRt.EntityData.SegmentPath = "v4-import-rt" + types.AddNoKeyToken(v4ImportRt)
    v4ImportRt.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-fti-oper:dci-fabric-interconnect/dci-vrfs/dci-vrf/" + v4ImportRt.EntityData.SegmentPath
    v4ImportRt.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    v4ImportRt.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    v4ImportRt.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    v4ImportRt.EntityData.Children = types.NewOrderedMap()
    v4ImportRt.EntityData.Leafs = types.NewOrderedMap()
    v4ImportRt.EntityData.Leafs.Append("rt-value", types.YLeaf{"RtValue", v4ImportRt.RtValue})
    v4ImportRt.EntityData.Leafs.Append("rt-flags", types.YLeaf{"RtFlags", v4ImportRt.RtFlags})

    v4ImportRt.EntityData.YListKeys = []string {}

    return &(v4ImportRt.EntityData)
}

// DciFabricInterconnect_DciVrfs_DciVrf_V4ExportRt
// V4 Export Route Targets
type DciFabricInterconnect_DciVrfs_DciVrf_V4ExportRt struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // RT Value. The type is string.
    RtValue interface{}

    // RT Flags. The type is string.
    RtFlags interface{}
}

func (v4ExportRt *DciFabricInterconnect_DciVrfs_DciVrf_V4ExportRt) GetEntityData() *types.CommonEntityData {
    v4ExportRt.EntityData.YFilter = v4ExportRt.YFilter
    v4ExportRt.EntityData.YangName = "v4-export-rt"
    v4ExportRt.EntityData.BundleName = "cisco_ios_xr"
    v4ExportRt.EntityData.ParentYangName = "dci-vrf"
    v4ExportRt.EntityData.SegmentPath = "v4-export-rt" + types.AddNoKeyToken(v4ExportRt)
    v4ExportRt.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-fti-oper:dci-fabric-interconnect/dci-vrfs/dci-vrf/" + v4ExportRt.EntityData.SegmentPath
    v4ExportRt.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    v4ExportRt.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    v4ExportRt.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    v4ExportRt.EntityData.Children = types.NewOrderedMap()
    v4ExportRt.EntityData.Leafs = types.NewOrderedMap()
    v4ExportRt.EntityData.Leafs.Append("rt-value", types.YLeaf{"RtValue", v4ExportRt.RtValue})
    v4ExportRt.EntityData.Leafs.Append("rt-flags", types.YLeaf{"RtFlags", v4ExportRt.RtFlags})

    v4ExportRt.EntityData.YListKeys = []string {}

    return &(v4ExportRt.EntityData)
}

// DciFabricInterconnect_DciVrfs_DciVrf_V6ImportRt
// V6 Import Route Targets
type DciFabricInterconnect_DciVrfs_DciVrf_V6ImportRt struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // RT Value. The type is string.
    RtValue interface{}

    // RT Flags. The type is string.
    RtFlags interface{}
}

func (v6ImportRt *DciFabricInterconnect_DciVrfs_DciVrf_V6ImportRt) GetEntityData() *types.CommonEntityData {
    v6ImportRt.EntityData.YFilter = v6ImportRt.YFilter
    v6ImportRt.EntityData.YangName = "v6-import-rt"
    v6ImportRt.EntityData.BundleName = "cisco_ios_xr"
    v6ImportRt.EntityData.ParentYangName = "dci-vrf"
    v6ImportRt.EntityData.SegmentPath = "v6-import-rt" + types.AddNoKeyToken(v6ImportRt)
    v6ImportRt.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-fti-oper:dci-fabric-interconnect/dci-vrfs/dci-vrf/" + v6ImportRt.EntityData.SegmentPath
    v6ImportRt.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    v6ImportRt.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    v6ImportRt.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    v6ImportRt.EntityData.Children = types.NewOrderedMap()
    v6ImportRt.EntityData.Leafs = types.NewOrderedMap()
    v6ImportRt.EntityData.Leafs.Append("rt-value", types.YLeaf{"RtValue", v6ImportRt.RtValue})
    v6ImportRt.EntityData.Leafs.Append("rt-flags", types.YLeaf{"RtFlags", v6ImportRt.RtFlags})

    v6ImportRt.EntityData.YListKeys = []string {}

    return &(v6ImportRt.EntityData)
}

// DciFabricInterconnect_DciVrfs_DciVrf_V6ExportRt
// V6 Export Route Targets
type DciFabricInterconnect_DciVrfs_DciVrf_V6ExportRt struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // RT Value. The type is string.
    RtValue interface{}

    // RT Flags. The type is string.
    RtFlags interface{}
}

func (v6ExportRt *DciFabricInterconnect_DciVrfs_DciVrf_V6ExportRt) GetEntityData() *types.CommonEntityData {
    v6ExportRt.EntityData.YFilter = v6ExportRt.YFilter
    v6ExportRt.EntityData.YangName = "v6-export-rt"
    v6ExportRt.EntityData.BundleName = "cisco_ios_xr"
    v6ExportRt.EntityData.ParentYangName = "dci-vrf"
    v6ExportRt.EntityData.SegmentPath = "v6-export-rt" + types.AddNoKeyToken(v6ExportRt)
    v6ExportRt.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-fti-oper:dci-fabric-interconnect/dci-vrfs/dci-vrf/" + v6ExportRt.EntityData.SegmentPath
    v6ExportRt.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    v6ExportRt.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    v6ExportRt.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    v6ExportRt.EntityData.Children = types.NewOrderedMap()
    v6ExportRt.EntityData.Leafs = types.NewOrderedMap()
    v6ExportRt.EntityData.Leafs.Append("rt-value", types.YLeaf{"RtValue", v6ExportRt.RtValue})
    v6ExportRt.EntityData.Leafs.Append("rt-flags", types.YLeaf{"RtFlags", v6ExportRt.RtFlags})

    v6ExportRt.EntityData.YListKeys = []string {}

    return &(v6ExportRt.EntityData)
}

// DciFabricInterconnect_Acp
// Auto Config Pool Info
type DciFabricInterconnect_Acp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VNI MIN VALUE. The type is interface{} with range: 0..4294967295.
    VniMin interface{}

    // VNI MAX VALUE. The type is interface{} with range: 0..4294967295.
    VniMax interface{}

    // Num of VNI used bits. The type is interface{} with range: 0..4294967295.
    VniBits interface{}

    // BVI MIN VALUE. The type is interface{} with range: 0..4294967295.
    BviMin interface{}

    // BVI MAX VALUE. The type is interface{} with range: 0..4294967295.
    BviMax interface{}

    // Num of BVI used bits. The type is interface{} with range: 0..4294967295.
    BviBits interface{}

    // BD MIN VALUE. The type is interface{} with range: 0..4294967295.
    BdMin interface{}

    // BD MAX VALUE. The type is interface{} with range: 0..4294967295.
    BdMax interface{}

    // Num of BD used bits. The type is interface{} with range: 0..4294967295.
    BdBits interface{}

    // Used VNI Range. The type is string.
    VniusedRange interface{}

    // Used BVI Range. The type is string.
    BviusedRange interface{}

    // Used BD Range. The type is string.
    BdusedRange interface{}
}

func (acp *DciFabricInterconnect_Acp) GetEntityData() *types.CommonEntityData {
    acp.EntityData.YFilter = acp.YFilter
    acp.EntityData.YangName = "acp"
    acp.EntityData.BundleName = "cisco_ios_xr"
    acp.EntityData.ParentYangName = "dci-fabric-interconnect"
    acp.EntityData.SegmentPath = "acp"
    acp.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-fti-oper:dci-fabric-interconnect/" + acp.EntityData.SegmentPath
    acp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    acp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    acp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    acp.EntityData.Children = types.NewOrderedMap()
    acp.EntityData.Leafs = types.NewOrderedMap()
    acp.EntityData.Leafs.Append("vni-min", types.YLeaf{"VniMin", acp.VniMin})
    acp.EntityData.Leafs.Append("vni-max", types.YLeaf{"VniMax", acp.VniMax})
    acp.EntityData.Leafs.Append("vni-bits", types.YLeaf{"VniBits", acp.VniBits})
    acp.EntityData.Leafs.Append("bvi-min", types.YLeaf{"BviMin", acp.BviMin})
    acp.EntityData.Leafs.Append("bvi-max", types.YLeaf{"BviMax", acp.BviMax})
    acp.EntityData.Leafs.Append("bvi-bits", types.YLeaf{"BviBits", acp.BviBits})
    acp.EntityData.Leafs.Append("bd-min", types.YLeaf{"BdMin", acp.BdMin})
    acp.EntityData.Leafs.Append("bd-max", types.YLeaf{"BdMax", acp.BdMax})
    acp.EntityData.Leafs.Append("bd-bits", types.YLeaf{"BdBits", acp.BdBits})
    acp.EntityData.Leafs.Append("vniused-range", types.YLeaf{"VniusedRange", acp.VniusedRange})
    acp.EntityData.Leafs.Append("bviused-range", types.YLeaf{"BviusedRange", acp.BviusedRange})
    acp.EntityData.Leafs.Append("bdused-range", types.YLeaf{"BdusedRange", acp.BdusedRange})

    acp.EntityData.YListKeys = []string {}

    return &(acp.EntityData)
}

