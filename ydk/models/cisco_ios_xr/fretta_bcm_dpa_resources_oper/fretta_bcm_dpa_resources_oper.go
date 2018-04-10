// This module contains a collection of YANG definitions
// for Cisco IOS-XR fretta-bcm-dpa-resources package operational data.
// 
// This module contains definitions
// for the following management objects:
//   dpa: Stats Data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package fretta_bcm_dpa_resources_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package fretta_bcm_dpa_resources_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-fretta-bcm-dpa-resources-oper dpa}", reflect.TypeOf(Dpa{}))
    ydk.RegisterEntity("Cisco-IOS-XR-fretta-bcm-dpa-resources-oper:dpa", reflect.TypeOf(Dpa{}))
}

// DpaTable represents Dpa table
type DpaTable string

const (
    // test if
    DpaTable_test_if DpaTable = "test-if"

    // test nhg
    DpaTable_test_nhg DpaTable = "test-nhg"

    // test nh
    DpaTable_test_nh DpaTable = "test-nh"

    // test rt
    DpaTable_test_rt DpaTable = "test-rt"

    // testdynhg
    DpaTable_test_dynhg DpaTable = "test-dynhg"

    // test ippuntpolicy
    DpaTable_test_ip_punt_policy DpaTable = "test-ip-punt-policy"

    // test punt
    DpaTable_test_punt DpaTable = "test-punt"

    // testpuntpolicystats
    DpaTable_test_punt_policy_stats DpaTable = "test-punt-policy-stats"

    // testasyncupdate
    DpaTable_test_async_update DpaTable = "test-async-update"

    // test ddel q
    DpaTable_test_ddel_q DpaTable = "test-ddel-q"

    // bdbvi
    DpaTable_bdbvi DpaTable = "bdbvi"

    // sys
    DpaTable_sys DpaTable = "sys"

    // npu
    DpaTable_npu DpaTable = "npu"

    // npuhwid
    DpaTable_npuhwid DpaTable = "npuhwid"

    // l1port
    DpaTable_l1_port DpaTable = "l1-port"

    // l2port
    DpaTable_l2_port DpaTable = "l2-port"

    // l2intf
    DpaTable_l2intf DpaTable = "l2intf"

    // mplspweport
    DpaTable_mplspwe_port DpaTable = "mplspwe-port"

    // mhpweport
    DpaTable_mhpwe_port DpaTable = "mhpwe-port"

    // l2xc
    DpaTable_l2xc DpaTable = "l2xc"

    // l2vpnstats
    DpaTable_l2vpnstats DpaTable = "l2vpnstats"

    // l1portstats
    DpaTable_l1_ports_tats DpaTable = "l1-ports-tats"

    // l3intf
    DpaTable_l3intf DpaTable = "l3intf"

    // l3intfrxstats
    DpaTable_l3intfrxstats DpaTable = "l3intfrxstats"

    // iproute
    DpaTable_ip_route DpaTable = "ip-route"

    // ip6route
    DpaTable_ip6_route DpaTable = "ip6-route"

    // puntpolicystats
    DpaTable_punt_policy_stats DpaTable = "punt-policy-stats"

    // ippuntpolicy
    DpaTable_ip_punt_policy DpaTable = "ip-punt-policy"

    // ip6puntpolicy
    DpaTable_ip6_punt_policy DpaTable = "ip6-punt-policy"

    // isispuntpolicy
    DpaTable_isis_punt_policy DpaTable = "isis-punt-policy"

    // ipnhgroup
    DpaTable_ipnh_group DpaTable = "ipnh-group"

    // ip6nhgroup
    DpaTable_ip6nh_group DpaTable = "ip6nh-group"

    // ipnh
    DpaTable_ipnh DpaTable = "ipnh"

    // ip6nh
    DpaTable_ip6nh DpaTable = "ip6nh"

    // ipvrf
    DpaTable_ipvrf DpaTable = "ipvrf"

    // mplsnh
    DpaTable_mplsnh DpaTable = "mplsnh"

    // mplslabel
    DpaTable_mpls_label DpaTable = "mpls-label"

    // dnx trap
    DpaTable_dnx_trap DpaTable = "dnx-trap"

    // punt
    DpaTable_punt DpaTable = "punt"

    // puntpolicer
    DpaTable_punt_police_r DpaTable = "punt-police-r"

    // puntlptspolicer
    DpaTable_punt_lpts_police_r DpaTable = "punt-lpts-police-r"

    // puntstats
    DpaTable_punt_stats DpaTable = "punt-stats"

    // tmport
    DpaTable_tm_port DpaTable = "tm-port"

    // spansession
    DpaTable_span_session DpaTable = "span-session"

    // policerstats
    DpaTable_police_rstats DpaTable = "police-rstats"

    // tmportstats
    DpaTable_tm_ports_tats DpaTable = "tm-ports-tats"

    // l3intftxstats
    DpaTable_l3intftxstats DpaTable = "l3intftxstats"

    // mplstetxstats
    DpaTable_mplstetxstats DpaTable = "mplstetxstats"

    // mplslblstats
    DpaTable_mplslblstats DpaTable = "mplslblstats"

    // policer
    DpaTable_police_r DpaTable = "police-r"

    // l2intfrxstats
    DpaTable_l2intfrxstats DpaTable = "l2intfrxstats"

    // l2intftxstats
    DpaTable_l2intftxstats DpaTable = "l2intftxstats"

    // dnx pbr tt ipv4
    DpaTable_dnx_pbr_tt_ipv4 DpaTable = "dnx-pbr-tt-ipv4"

    // dnx pbr tt ipv6
    DpaTable_dnx_pbr_tt_ipv6 DpaTable = "dnx-pbr-tt-ipv6"

    // bfdhwoff
    DpaTable_bfdhwoff DpaTable = "bfdhwoff"

    // global
    DpaTable_global DpaTable = "global"

    // lagport
    DpaTable_lag_port DpaTable = "lag-port"

    // qosprofile
    DpaTable_qos_profile DpaTable = "qos-profile"

    // tmrateprofile
    DpaTable_tmrate_profile DpaTable = "tmrate-profile"

    // dnx port range
    DpaTable_dnx_port_range DpaTable = "dnx-port-range"

    // ipacl
    DpaTable_ipacl DpaTable = "ipacl"

    // ip6acl
    DpaTable_ip6acl DpaTable = "ip6acl"

    // schedtree
    DpaTable_sched_tree DpaTable = "sched-tree"

    // tmcos
    DpaTable_tmcos DpaTable = "tmcos"

    // statsagg
    DpaTable_statsagg DpaTable = "statsagg"

    // nhprotect
    DpaTable_nhprotect DpaTable = "nhprotect"

    // sampler
    DpaTable_sampler DpaTable = "sampler"

    // l2qos
    DpaTable_l2qos DpaTable = "l2qos"

    // peerqos
    DpaTable_peer_qos DpaTable = "peer-qos"

    // ipqos
    DpaTable_ipqos DpaTable = "ipqos"

    // ip6qos
    DpaTable_ip6qos DpaTable = "ip6qos"

    // mplsqos
    DpaTable_mplsqos DpaTable = "mplsqos"

    // qosid
    DpaTable_qosid DpaTable = "qosid"

    // extlif
    DpaTable_extlif DpaTable = "extlif"

    // elif
    DpaTable_elif DpaTable = "elif"

    // ingaclstats
    DpaTable_ingaclstats DpaTable = "ingaclstats"

    // egraclstats
    DpaTable_egraclstats DpaTable = "egraclstats"

    // edpl
    DpaTable_edpl DpaTable = "edpl"

    // cfmmaid
    DpaTable_cfmmaid DpaTable = "cfmmaid"

    // cfmdefmps
    DpaTable_cfmdefmps DpaTable = "cfmdefmps"

    // cfmofflmep
    DpaTable_cfmofflmep DpaTable = "cfmofflmep"

    // cfmoffrmep
    DpaTable_cfmoffrmep DpaTable = "cfmoffrmep"

    // cfmnonoff
    DpaTable_cfmnonoff DpaTable = "cfmnonoff"

    // cfmhwoffrxstats
    DpaTable_cfmhwoffrxstats DpaTable = "cfmhwoffrxstats"

    // ipmcroute
    DpaTable_ipmc_route DpaTable = "ipmc-route"

    // l2ipmcroute
    DpaTable_l2ipmc_route DpaTable = "l2ipmc-route"

    // ipmcolist
    DpaTable_ipmco_list DpaTable = "ipmco-list"

    // l2mcolist
    DpaTable_l2mco_list DpaTable = "l2mco-list"

    // ipmcmergedolist
    DpaTable_ipmc_merge_do_list DpaTable = "ipmc-merge-do-list"

    // sgfgidlist
    DpaTable_sgfgid_list DpaTable = "sgfgid-list"

    // meshmc
    DpaTable_meshmc DpaTable = "meshmc"

    // l2bridge
    DpaTable_l2_bridge DpaTable = "l2-bridge"

    // l2bridgeport
    DpaTable_l2_bridge_port DpaTable = "l2-bridge-port"

    // l2bridgemac
    DpaTable_l2_bridge_mac DpaTable = "l2-bridge-mac"

    // l2brmac
    DpaTable_l2brmac DpaTable = "l2brmac"

    // iptunneldecap
    DpaTable_ip_tunnel_decap DpaTable = "ip-tunnel-decap"

    // l2vlanrange
    DpaTable_l2vlan_range DpaTable = "l2vlan-range"

    // iptunnelencap
    DpaTable_ip_tunnel_encap DpaTable = "ip-tunnel-encap"

    // rawget
    DpaTable_rawget DpaTable = "rawget"

    // ip6mcroute
    DpaTable_ip6mc_route DpaTable = "ip6mc-route"

    // l2evpnactremotepeerid
    DpaTable_l2evpnact_remote_peer_id DpaTable = "l2evpnact-remote-peer-id"

    // l2evpnactlocalshl
    DpaTable_l2evpnact_local_shl DpaTable = "l2evpnact-local-shl"

    // l2evpnactremoteshl
    DpaTable_l2evpnact_remote_shl DpaTable = "l2evpnact-remote-shl"

    // evpn imlrange
    DpaTable_evpn_iml_range DpaTable = "evpn-iml-range"

    // l2bridgeolist
    DpaTable_l2_bridge_o_list DpaTable = "l2-bridge-o-list"

    // l2acl
    DpaTable_l2acl DpaTable = "l2acl"

    // l2evpn nh
    DpaTable_l2evpn_nh DpaTable = "l2evpn-nh"

    // l2bridgeport sc
    DpaTable_l2_bridge_port_sc DpaTable = "l2-bridge-port-sc"

    // l3intfmctxstats
    DpaTable_l3intfmctxstats DpaTable = "l3intfmctxstats"

    // tidl sample
    DpaTable_tidl_sample DpaTable = "tidl-sample"

    // tidl ref sample
    DpaTable_tidl_ref_sample DpaTable = "tidl-ref-sample"

    // ipaclprefix
    DpaTable_ipacl_prefix DpaTable = "ipacl-prefix"

    // ip6aclprefix
    DpaTable_ip6acl_prefix DpaTable = "ip6acl-prefix"

    // ipaclport
    DpaTable_ipacl_port DpaTable = "ipacl-port"

    // scaleacl
    DpaTable_scaleacl DpaTable = "scaleacl"

    // ipmcfhop
    DpaTable_ipmcf_hop DpaTable = "ipmcf-hop"

    // bundle swoff
    DpaTable_bundle_swoff DpaTable = "bundle-swoff"

    // mcidswoff
    DpaTable_mcidswoff DpaTable = "mcidswoff"

    // destmap
    DpaTable_dest_map DpaTable = "dest-map"

    // l2bridgeport pw
    DpaTable_l2_bridge_port_pw DpaTable = "l2-bridge-port-pw"

    // l2evpnactlocalshlstats
    DpaTable_l2evpnact_local_shlstats DpaTable = "l2evpnact-local-shlstats"

    // testhidden
    DpaTable_test_hidden DpaTable = "test-hidden"

    // testlocal
    DpaTable_test_local DpaTable = "test-local"

    // testrepeated
    DpaTable_test_repeated DpaTable = "test-repeated"

    // limd
    DpaTable_limd DpaTable = "limd"

    // litap
    DpaTable_litap DpaTable = "litap"

    // test xtf
    DpaTable_test_xtf DpaTable = "test-xtf"

    // ipmcrxstats
    DpaTable_ipmcrxstats DpaTable = "ipmcrxstats"
)

// Dpa
// Stats Data
type Dpa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DPA Resources Data.
    Resources Dpa_Resources

    // Voq or Trap Data.
    Stats Dpa_Stats
}

func (dpa *Dpa) GetEntityData() *types.CommonEntityData {
    dpa.EntityData.YFilter = dpa.YFilter
    dpa.EntityData.YangName = "dpa"
    dpa.EntityData.BundleName = "cisco_ios_xr"
    dpa.EntityData.ParentYangName = "Cisco-IOS-XR-fretta-bcm-dpa-resources-oper"
    dpa.EntityData.SegmentPath = "Cisco-IOS-XR-fretta-bcm-dpa-resources-oper:dpa"
    dpa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dpa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dpa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dpa.EntityData.Children = make(map[string]types.YChild)
    dpa.EntityData.Children["resources"] = types.YChild{"Resources", &dpa.Resources}
    dpa.EntityData.Children["stats"] = types.YChild{"Stats", &dpa.Stats}
    dpa.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dpa.EntityData)
}

// Dpa_Resources
// DPA Resources Data
type Dpa_Resources struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DPA data for available nodes.
    Nodes Dpa_Resources_Nodes
}

func (resources *Dpa_Resources) GetEntityData() *types.CommonEntityData {
    resources.EntityData.YFilter = resources.YFilter
    resources.EntityData.YangName = "resources"
    resources.EntityData.BundleName = "cisco_ios_xr"
    resources.EntityData.ParentYangName = "dpa"
    resources.EntityData.SegmentPath = "resources"
    resources.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    resources.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    resources.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    resources.EntityData.Children = make(map[string]types.YChild)
    resources.EntityData.Children["nodes"] = types.YChild{"Nodes", &resources.Nodes}
    resources.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(resources.EntityData)
}

// Dpa_Resources_Nodes
// DPA data for available nodes
type Dpa_Resources_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DPA operational data for a particular node. The type is slice of
    // Dpa_Resources_Nodes_Node.
    Node []Dpa_Resources_Nodes_Node
}

func (nodes *Dpa_Resources_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "resources"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodes.EntityData.Children = make(map[string]types.YChild)
    nodes.EntityData.Children["node"] = types.YChild{"Node", nil}
    for i := range nodes.Node {
        nodes.EntityData.Children[types.GetSegmentPath(&nodes.Node[i])] = types.YChild{"Node", &nodes.Node[i]}
    }
    nodes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nodes.EntityData)
}

// Dpa_Resources_Nodes_Node
// DPA operational data for a particular node
type Dpa_Resources_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node ID. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeName interface{}

    // DPA Resources table.
    TableDatas Dpa_Resources_Nodes_Node_TableDatas
}

func (node *Dpa_Resources_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Children["table-datas"] = types.YChild{"TableDatas", &node.TableDatas}
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", node.NodeName}
    return &(node.EntityData)
}

// Dpa_Resources_Nodes_Node_TableDatas
// DPA Resources table
type Dpa_Resources_Nodes_Node_TableDatas struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DPA Resources table. The type is slice of
    // Dpa_Resources_Nodes_Node_TableDatas_TableData.
    TableData []Dpa_Resources_Nodes_Node_TableDatas_TableData
}

func (tableDatas *Dpa_Resources_Nodes_Node_TableDatas) GetEntityData() *types.CommonEntityData {
    tableDatas.EntityData.YFilter = tableDatas.YFilter
    tableDatas.EntityData.YangName = "table-datas"
    tableDatas.EntityData.BundleName = "cisco_ios_xr"
    tableDatas.EntityData.ParentYangName = "node"
    tableDatas.EntityData.SegmentPath = "table-datas"
    tableDatas.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tableDatas.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tableDatas.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tableDatas.EntityData.Children = make(map[string]types.YChild)
    tableDatas.EntityData.Children["table-data"] = types.YChild{"TableData", nil}
    for i := range tableDatas.TableData {
        tableDatas.EntityData.Children[types.GetSegmentPath(&tableDatas.TableData[i])] = types.YChild{"TableData", &tableDatas.TableData[i]}
    }
    tableDatas.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(tableDatas.EntityData)
}

// Dpa_Resources_Nodes_Node_TableDatas_TableData
// DPA Resources table
type Dpa_Resources_Nodes_Node_TableDatas_TableData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Resource type. The type is DpaTable.
    Resource interface{}

    // table id. The type is interface{} with range: 0..4294967295.
    TableId interface{}

    // name. The type is string with length: 0..64.
    Name interface{}

    // is global. The type is bool.
    IsGlobal interface{}

    // num npus. The type is interface{} with range: 0..4294967295.
    NumNpus interface{}

    // table specific list. The type is string.
    TableSpecificList interface{}

    // npu tblr. The type is slice of
    // Dpa_Resources_Nodes_Node_TableDatas_TableData_NpuTblr.
    NpuTblr []Dpa_Resources_Nodes_Node_TableDatas_TableData_NpuTblr
}

func (tableData *Dpa_Resources_Nodes_Node_TableDatas_TableData) GetEntityData() *types.CommonEntityData {
    tableData.EntityData.YFilter = tableData.YFilter
    tableData.EntityData.YangName = "table-data"
    tableData.EntityData.BundleName = "cisco_ios_xr"
    tableData.EntityData.ParentYangName = "table-datas"
    tableData.EntityData.SegmentPath = "table-data" + "[resource='" + fmt.Sprintf("%v", tableData.Resource) + "']"
    tableData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tableData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tableData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tableData.EntityData.Children = make(map[string]types.YChild)
    tableData.EntityData.Children["npu-tblr"] = types.YChild{"NpuTblr", nil}
    for i := range tableData.NpuTblr {
        tableData.EntityData.Children[types.GetSegmentPath(&tableData.NpuTblr[i])] = types.YChild{"NpuTblr", &tableData.NpuTblr[i]}
    }
    tableData.EntityData.Leafs = make(map[string]types.YLeaf)
    tableData.EntityData.Leafs["resource"] = types.YLeaf{"Resource", tableData.Resource}
    tableData.EntityData.Leafs["table-id"] = types.YLeaf{"TableId", tableData.TableId}
    tableData.EntityData.Leafs["name"] = types.YLeaf{"Name", tableData.Name}
    tableData.EntityData.Leafs["is-global"] = types.YLeaf{"IsGlobal", tableData.IsGlobal}
    tableData.EntityData.Leafs["num-npus"] = types.YLeaf{"NumNpus", tableData.NumNpus}
    tableData.EntityData.Leafs["table-specific-list"] = types.YLeaf{"TableSpecificList", tableData.TableSpecificList}
    return &(tableData.EntityData)
}

// Dpa_Resources_Nodes_Node_TableDatas_TableData_NpuTblr
// npu tblr
type Dpa_Resources_Nodes_Node_TableDatas_TableData_NpuTblr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // npu id. The type is interface{} with range: 0..4294967295.
    NpuId interface{}

    // num objs. The type is interface{} with range: 0..4294967295.
    NumObjs interface{}

    // create req. The type is interface{} with range: 0..4294967295.
    CreateReq interface{}

    // create ok. The type is interface{} with range: 0..4294967295.
    CreateOk interface{}

    // delete req. The type is interface{} with range: 0..4294967295.
    DeleteReq interface{}

    // delete ok. The type is interface{} with range: 0..4294967295.
    DeleteOk interface{}

    // update req. The type is interface{} with range: 0..4294967295.
    UpdateReq interface{}

    // update ok. The type is interface{} with range: 0..4294967295.
    UpdateOk interface{}

    // eod req. The type is interface{} with range: 0..4294967295.
    EodReq interface{}

    // eod ok. The type is interface{} with range: 0..4294967295.
    EodOk interface{}

    // err hw fail. The type is interface{} with range: 0..4294967295.
    ErrHwFail interface{}

    // err ref resolve. The type is interface{} with range: 0..4294967295.
    ErrRefResolve interface{}

    // err db notfound. The type is interface{} with range: 0..4294967295.
    ErrDbNotfound interface{}

    // err db exists. The type is interface{} with range: 0..4294967295.
    ErrDbExists interface{}

    // err db nomem. The type is interface{} with range: 0..4294967295.
    ErrDbNomem interface{}
}

func (npuTblr *Dpa_Resources_Nodes_Node_TableDatas_TableData_NpuTblr) GetEntityData() *types.CommonEntityData {
    npuTblr.EntityData.YFilter = npuTblr.YFilter
    npuTblr.EntityData.YangName = "npu-tblr"
    npuTblr.EntityData.BundleName = "cisco_ios_xr"
    npuTblr.EntityData.ParentYangName = "table-data"
    npuTblr.EntityData.SegmentPath = "npu-tblr"
    npuTblr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    npuTblr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    npuTblr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    npuTblr.EntityData.Children = make(map[string]types.YChild)
    npuTblr.EntityData.Leafs = make(map[string]types.YLeaf)
    npuTblr.EntityData.Leafs["npu-id"] = types.YLeaf{"NpuId", npuTblr.NpuId}
    npuTblr.EntityData.Leafs["num-objs"] = types.YLeaf{"NumObjs", npuTblr.NumObjs}
    npuTblr.EntityData.Leafs["create-req"] = types.YLeaf{"CreateReq", npuTblr.CreateReq}
    npuTblr.EntityData.Leafs["create-ok"] = types.YLeaf{"CreateOk", npuTblr.CreateOk}
    npuTblr.EntityData.Leafs["delete-req"] = types.YLeaf{"DeleteReq", npuTblr.DeleteReq}
    npuTblr.EntityData.Leafs["delete-ok"] = types.YLeaf{"DeleteOk", npuTblr.DeleteOk}
    npuTblr.EntityData.Leafs["update-req"] = types.YLeaf{"UpdateReq", npuTblr.UpdateReq}
    npuTblr.EntityData.Leafs["update-ok"] = types.YLeaf{"UpdateOk", npuTblr.UpdateOk}
    npuTblr.EntityData.Leafs["eod-req"] = types.YLeaf{"EodReq", npuTblr.EodReq}
    npuTblr.EntityData.Leafs["eod-ok"] = types.YLeaf{"EodOk", npuTblr.EodOk}
    npuTblr.EntityData.Leafs["err-hw-fail"] = types.YLeaf{"ErrHwFail", npuTblr.ErrHwFail}
    npuTblr.EntityData.Leafs["err-ref-resolve"] = types.YLeaf{"ErrRefResolve", npuTblr.ErrRefResolve}
    npuTblr.EntityData.Leafs["err-db-notfound"] = types.YLeaf{"ErrDbNotfound", npuTblr.ErrDbNotfound}
    npuTblr.EntityData.Leafs["err-db-exists"] = types.YLeaf{"ErrDbExists", npuTblr.ErrDbExists}
    npuTblr.EntityData.Leafs["err-db-nomem"] = types.YLeaf{"ErrDbNomem", npuTblr.ErrDbNomem}
    return &(npuTblr.EntityData)
}

// Dpa_Stats
// Voq or Trap Data
type Dpa_Stats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DPA data for available nodes.
    Nodes Dpa_Stats_Nodes
}

func (stats *Dpa_Stats) GetEntityData() *types.CommonEntityData {
    stats.EntityData.YFilter = stats.YFilter
    stats.EntityData.YangName = "stats"
    stats.EntityData.BundleName = "cisco_ios_xr"
    stats.EntityData.ParentYangName = "dpa"
    stats.EntityData.SegmentPath = "stats"
    stats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stats.EntityData.Children = make(map[string]types.YChild)
    stats.EntityData.Children["nodes"] = types.YChild{"Nodes", &stats.Nodes}
    stats.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(stats.EntityData)
}

// Dpa_Stats_Nodes
// DPA data for available nodes
type Dpa_Stats_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DPA operational data for a particular node. The type is slice of
    // Dpa_Stats_Nodes_Node.
    Node []Dpa_Stats_Nodes_Node
}

func (nodes *Dpa_Stats_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "stats"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodes.EntityData.Children = make(map[string]types.YChild)
    nodes.EntityData.Children["node"] = types.YChild{"Node", nil}
    for i := range nodes.Node {
        nodes.EntityData.Children[types.GetSegmentPath(&nodes.Node[i])] = types.YChild{"Node", &nodes.Node[i]}
    }
    nodes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nodes.EntityData)
}

// Dpa_Stats_Nodes_Node
// DPA operational data for a particular node
type Dpa_Stats_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node ID. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeName interface{}

    // ASIC statistics table.
    AsicStatistics Dpa_Stats_Nodes_Node_AsicStatistics

    // Ingress Stats.
    NpuNumbers Dpa_Stats_Nodes_Node_NpuNumbers
}

func (node *Dpa_Stats_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Children["asic-statistics"] = types.YChild{"AsicStatistics", &node.AsicStatistics}
    node.EntityData.Children["npu-numbers"] = types.YChild{"NpuNumbers", &node.NpuNumbers}
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", node.NodeName}
    return &(node.EntityData)
}

// Dpa_Stats_Nodes_Node_AsicStatistics
// ASIC statistics table
type Dpa_Stats_Nodes_Node_AsicStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ASIC statistics.
    AsicStatisticsForNpuIds Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds

    // Detailed ASIC statistics.
    AsicStatisticsDetailForNpuIds Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds
}

func (asicStatistics *Dpa_Stats_Nodes_Node_AsicStatistics) GetEntityData() *types.CommonEntityData {
    asicStatistics.EntityData.YFilter = asicStatistics.YFilter
    asicStatistics.EntityData.YangName = "asic-statistics"
    asicStatistics.EntityData.BundleName = "cisco_ios_xr"
    asicStatistics.EntityData.ParentYangName = "node"
    asicStatistics.EntityData.SegmentPath = "asic-statistics"
    asicStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asicStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asicStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asicStatistics.EntityData.Children = make(map[string]types.YChild)
    asicStatistics.EntityData.Children["asic-statistics-for-npu-ids"] = types.YChild{"AsicStatisticsForNpuIds", &asicStatistics.AsicStatisticsForNpuIds}
    asicStatistics.EntityData.Children["asic-statistics-detail-for-npu-ids"] = types.YChild{"AsicStatisticsDetailForNpuIds", &asicStatistics.AsicStatisticsDetailForNpuIds}
    asicStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(asicStatistics.EntityData)
}

// Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds
// ASIC statistics
type Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ASIC statistics for a particular NPU. The type is slice of
    // Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds_AsicStatisticsForNpuId.
    AsicStatisticsForNpuId []Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds_AsicStatisticsForNpuId
}

func (asicStatisticsForNpuIds *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds) GetEntityData() *types.CommonEntityData {
    asicStatisticsForNpuIds.EntityData.YFilter = asicStatisticsForNpuIds.YFilter
    asicStatisticsForNpuIds.EntityData.YangName = "asic-statistics-for-npu-ids"
    asicStatisticsForNpuIds.EntityData.BundleName = "cisco_ios_xr"
    asicStatisticsForNpuIds.EntityData.ParentYangName = "asic-statistics"
    asicStatisticsForNpuIds.EntityData.SegmentPath = "asic-statistics-for-npu-ids"
    asicStatisticsForNpuIds.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asicStatisticsForNpuIds.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asicStatisticsForNpuIds.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asicStatisticsForNpuIds.EntityData.Children = make(map[string]types.YChild)
    asicStatisticsForNpuIds.EntityData.Children["asic-statistics-for-npu-id"] = types.YChild{"AsicStatisticsForNpuId", nil}
    for i := range asicStatisticsForNpuIds.AsicStatisticsForNpuId {
        asicStatisticsForNpuIds.EntityData.Children[types.GetSegmentPath(&asicStatisticsForNpuIds.AsicStatisticsForNpuId[i])] = types.YChild{"AsicStatisticsForNpuId", &asicStatisticsForNpuIds.AsicStatisticsForNpuId[i]}
    }
    asicStatisticsForNpuIds.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(asicStatisticsForNpuIds.EntityData)
}

// Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds_AsicStatisticsForNpuId
// ASIC statistics for a particular NPU
type Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds_AsicStatisticsForNpuId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. NPU number. The type is interface{} with range:
    // -2147483648..2147483647.
    NpuId interface{}

    // Flag to indicate if data is valid. The type is bool.
    Valid interface{}

    // Rack number. The type is interface{} with range: 0..4294967295.
    RackNumber interface{}

    // Slot number. The type is interface{} with range: 0..4294967295.
    SlotNumber interface{}

    // ASIC instance. The type is interface{} with range: 0..4294967295.
    AsicInstance interface{}

    // Chip version. The type is interface{} with range: 0..65535.
    ChipVersion interface{}

    // Statistics.
    Statistics Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds_AsicStatisticsForNpuId_Statistics
}

func (asicStatisticsForNpuId *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds_AsicStatisticsForNpuId) GetEntityData() *types.CommonEntityData {
    asicStatisticsForNpuId.EntityData.YFilter = asicStatisticsForNpuId.YFilter
    asicStatisticsForNpuId.EntityData.YangName = "asic-statistics-for-npu-id"
    asicStatisticsForNpuId.EntityData.BundleName = "cisco_ios_xr"
    asicStatisticsForNpuId.EntityData.ParentYangName = "asic-statistics-for-npu-ids"
    asicStatisticsForNpuId.EntityData.SegmentPath = "asic-statistics-for-npu-id" + "[npu-id='" + fmt.Sprintf("%v", asicStatisticsForNpuId.NpuId) + "']"
    asicStatisticsForNpuId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asicStatisticsForNpuId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asicStatisticsForNpuId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asicStatisticsForNpuId.EntityData.Children = make(map[string]types.YChild)
    asicStatisticsForNpuId.EntityData.Children["statistics"] = types.YChild{"Statistics", &asicStatisticsForNpuId.Statistics}
    asicStatisticsForNpuId.EntityData.Leafs = make(map[string]types.YLeaf)
    asicStatisticsForNpuId.EntityData.Leafs["npu-id"] = types.YLeaf{"NpuId", asicStatisticsForNpuId.NpuId}
    asicStatisticsForNpuId.EntityData.Leafs["valid"] = types.YLeaf{"Valid", asicStatisticsForNpuId.Valid}
    asicStatisticsForNpuId.EntityData.Leafs["rack-number"] = types.YLeaf{"RackNumber", asicStatisticsForNpuId.RackNumber}
    asicStatisticsForNpuId.EntityData.Leafs["slot-number"] = types.YLeaf{"SlotNumber", asicStatisticsForNpuId.SlotNumber}
    asicStatisticsForNpuId.EntityData.Leafs["asic-instance"] = types.YLeaf{"AsicInstance", asicStatisticsForNpuId.AsicInstance}
    asicStatisticsForNpuId.EntityData.Leafs["chip-version"] = types.YLeaf{"ChipVersion", asicStatisticsForNpuId.ChipVersion}
    return &(asicStatisticsForNpuId.EntityData)
}

// Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds_AsicStatisticsForNpuId_Statistics
// Statistics
type Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds_AsicStatisticsForNpuId_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Total bytes sent from NIF to IRE. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    NbiRxTotalByteCnt interface{}

    // Total packets sent from NIF to IRE. The type is interface{} with range:
    // 0..18446744073709551615.
    NbiRxTotalPktCnt interface{}

    // CPU ingress received packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    IreCpuPktCnt interface{}

    // NIF received packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    IreNifPktCnt interface{}

    // OAMP ingress received packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    IreOampPktCnt interface{}

    // OLP ingress received packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    IreOlpPktCnt interface{}

    // Recycling ingress received packet counter. The type is interface{} with
    // range: 0..18446744073709551615.
    IreRcyPktCnt interface{}

    // Performance counter of the FDT interface. The type is interface{} with
    // range: 0..18446744073709551615.
    IreFdtIfCnt interface{}

    // Performance counter of the MMU interface. The type is interface{} with
    // range: 0..18446744073709551615.
    IdrMmuIfCnt interface{}

    // Performance counter of the OCB interface. The type is interface{} with
    // range: 0..18446744073709551615.
    IdrOcbIfCnt interface{}

    // Counts enqueued packets. The type is interface{} with range:
    // 0..18446744073709551615.
    IqmEnqueuePktCnt interface{}

    // Counts dequeued packets. The type is interface{} with range:
    // 0..18446744073709551615.
    IqmDequeuePktCnt interface{}

    // Counts matched packets discarded in the DEQ process. The type is
    // interface{} with range: 0..18446744073709551615.
    IqmDeletedPktCnt interface{}

    // Counts all packets discarded at the ENQ pipe. The type is interface{} with
    // range: 0..18446744073709551615.
    IqmEnqDiscardedPktCnt interface{}

    // EGQ packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    IptEgqPktCnt interface{}

    // ENQ packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    IptEnqPktCnt interface{}

    // FDT packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    IptFdtPktCnt interface{}

    // Configurable event counter. The type is interface{} with range:
    // 0..18446744073709551615.
    IptCfgEventCnt interface{}

    // Configurable bytes counter. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    IptCfgByteCnt interface{}

    // Descriptor cell counter. The type is interface{} with range:
    // 0..18446744073709551615.
    FdtIptDescCellCnt interface{}

    // IRE internal descriptor cell counter. The type is interface{} with range:
    // 0..18446744073709551615.
    FdtIreDescCellCnt interface{}

    // Counts all transmitted data cells. The type is interface{} with range:
    // 0..18446744073709551615.
    FdtTransmittedDataCellsCnt interface{}

    // FDR total incoming cell counter at pipe 1. The type is interface{} with
    // range: 0..18446744073709551615.
    FdrP1CellInCnt interface{}

    // FDR total incoming cell counter at pipe 2. The type is interface{} with
    // range: 0..18446744073709551615.
    FdrP2CellInCnt interface{}

    // FDR total incoming cell counter at pipe 3. The type is interface{} with
    // range: 0..18446744073709551615.
    FdrP3CellInCnt interface{}

    // FDR total incoming cell counter. The type is interface{} with range:
    // 0..18446744073709551615.
    FdrCellInCntTotal interface{}

    // FDA input cell counter P1. The type is interface{} with range:
    // 0..18446744073709551615.
    FdaCellsInCntP1 interface{}

    // FDA input cell counter P2. The type is interface{} with range:
    // 0..18446744073709551615.
    FdaCellsInCntP2 interface{}

    // FDA input cell counter P3. The type is interface{} with range:
    // 0..18446744073709551615.
    FdaCellsInCntP3 interface{}

    // FDA input cell counter TDM. The type is interface{} with range:
    // 0..18446744073709551615.
    FdaCellsInTdmCnt interface{}

    // FDA input cell counter MESHMC. The type is interface{} with range:
    // 0..18446744073709551615.
    FdaCellsInMeshmcCnt interface{}

    // FDA input cell counter IPT. The type is interface{} with range:
    // 0..18446744073709551615.
    FdaCellsInIptCnt interface{}

    // FDA output cell counter P1. The type is interface{} with range:
    // 0..18446744073709551615.
    FdaCellsOutCntP1 interface{}

    // FDA output cell counter P2. The type is interface{} with range:
    // 0..18446744073709551615.
    FdaCellsOutCntP2 interface{}

    // FDA output cell counter P3. The type is interface{} with range:
    // 0..18446744073709551615.
    FdaCellsOutCntP3 interface{}

    // FDA output cell counter TDM. The type is interface{} with range:
    // 0..18446744073709551615.
    FdaCellsOutTdmCnt interface{}

    // FDA output cell counter MESHMC. The type is interface{} with range:
    // 0..18446744073709551615.
    FdaCellsOutMeshmcCnt interface{}

    // FDA output cell counter IPT. The type is interface{} with range:
    // 0..18446744073709551615.
    FdaCellsOutIptCnt interface{}

    // FDA EGQ drop counter. The type is interface{} with range:
    // 0..18446744073709551615.
    FdaEgqDropCnt interface{}

    // FDA EGQ MESHMC drop counter. The type is interface{} with range:
    // 0..18446744073709551615.
    FdaEgqMeshmcDropCnt interface{}

    // FQP2EPE packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    EgqFqpPktCnt interface{}

    // PQP2FQP unicast packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    EgqPqpUcPktCnt interface{}

    // PQP discarded unicast packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    EgqPqpDiscardUcPktCnt interface{}

    // PQP2FQP unicast bytes counter. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    EgqPqpUcBytesCnt interface{}

    // PQP2FQP multicast packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    EgqPqpMcPktCnt interface{}

    // PQP discarded multicast packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    EgqPqpDiscardMcPktCnt interface{}

    // PQP2FQP multicast bytes counter. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    EgqPqpMcBytesCnt interface{}

    // EHP2PQP unicast packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    EgqEhpUcPktCnt interface{}

    // EHP2PQP multicast high packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    EgqEhpMcHighPktCnt interface{}

    // EHP2PQP multicast low packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    EgqEhpMcLowPktCnt interface{}

    // EHP2PQP discarded packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    EgqDeletedPktCnt interface{}

    // Number of multicast high packets discarded because multicast FIFO is full.
    // The type is interface{} with range: 0..18446744073709551615.
    EgqEhpMcHighDiscardCnt interface{}

    // Number of multicast low packets discarded because multicast FIFO is full.
    // The type is interface{} with range: 0..18446744073709551615.
    EgqEhpMcLowDiscardCnt interface{}

    // Number of packet descriptors discarded due to LAG multicast pruning. The
    // type is interface{} with range: 0..18446744073709551615.
    EgqErppLagPruningDiscardCnt interface{}

    // Number of packet descriptors discarded due to ERPP PMF. The type is
    // interface{} with range: 0..18446744073709551615.
    EgqErppPmfDiscardCnt interface{}

    // Number of packet descriptors discarded because of egress VLAN membership.
    // The type is interface{} with range: 0..18446744073709551615.
    EgqErppVlanMbrDiscardCnt interface{}

    // EPE2PNI bytes counter. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    EpniEpeByteCnt interface{}

    // EPE2PNI packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    EpniEpePktCnt interface{}

    // EPE discarded packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    EpniEpeDiscardCnt interface{}

    // Total bytes sent from EGQ to NIF. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    NbiTxTotalByteCnt interface{}

    // Total packets sent from EGQ to NIF. The type is interface{} with range:
    // 0..18446744073709551615.
    NbiTxTotalPktCnt interface{}
}

func (statistics *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds_AsicStatisticsForNpuId_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "asic-statistics-for-npu-id"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = make(map[string]types.YChild)
    statistics.EntityData.Leafs = make(map[string]types.YLeaf)
    statistics.EntityData.Leafs["nbi-rx-total-byte-cnt"] = types.YLeaf{"NbiRxTotalByteCnt", statistics.NbiRxTotalByteCnt}
    statistics.EntityData.Leafs["nbi-rx-total-pkt-cnt"] = types.YLeaf{"NbiRxTotalPktCnt", statistics.NbiRxTotalPktCnt}
    statistics.EntityData.Leafs["ire-cpu-pkt-cnt"] = types.YLeaf{"IreCpuPktCnt", statistics.IreCpuPktCnt}
    statistics.EntityData.Leafs["ire-nif-pkt-cnt"] = types.YLeaf{"IreNifPktCnt", statistics.IreNifPktCnt}
    statistics.EntityData.Leafs["ire-oamp-pkt-cnt"] = types.YLeaf{"IreOampPktCnt", statistics.IreOampPktCnt}
    statistics.EntityData.Leafs["ire-olp-pkt-cnt"] = types.YLeaf{"IreOlpPktCnt", statistics.IreOlpPktCnt}
    statistics.EntityData.Leafs["ire-rcy-pkt-cnt"] = types.YLeaf{"IreRcyPktCnt", statistics.IreRcyPktCnt}
    statistics.EntityData.Leafs["ire-fdt-if-cnt"] = types.YLeaf{"IreFdtIfCnt", statistics.IreFdtIfCnt}
    statistics.EntityData.Leafs["idr-mmu-if-cnt"] = types.YLeaf{"IdrMmuIfCnt", statistics.IdrMmuIfCnt}
    statistics.EntityData.Leafs["idr-ocb-if-cnt"] = types.YLeaf{"IdrOcbIfCnt", statistics.IdrOcbIfCnt}
    statistics.EntityData.Leafs["iqm-enqueue-pkt-cnt"] = types.YLeaf{"IqmEnqueuePktCnt", statistics.IqmEnqueuePktCnt}
    statistics.EntityData.Leafs["iqm-dequeue-pkt-cnt"] = types.YLeaf{"IqmDequeuePktCnt", statistics.IqmDequeuePktCnt}
    statistics.EntityData.Leafs["iqm-deleted-pkt-cnt"] = types.YLeaf{"IqmDeletedPktCnt", statistics.IqmDeletedPktCnt}
    statistics.EntityData.Leafs["iqm-enq-discarded-pkt-cnt"] = types.YLeaf{"IqmEnqDiscardedPktCnt", statistics.IqmEnqDiscardedPktCnt}
    statistics.EntityData.Leafs["ipt-egq-pkt-cnt"] = types.YLeaf{"IptEgqPktCnt", statistics.IptEgqPktCnt}
    statistics.EntityData.Leafs["ipt-enq-pkt-cnt"] = types.YLeaf{"IptEnqPktCnt", statistics.IptEnqPktCnt}
    statistics.EntityData.Leafs["ipt-fdt-pkt-cnt"] = types.YLeaf{"IptFdtPktCnt", statistics.IptFdtPktCnt}
    statistics.EntityData.Leafs["ipt-cfg-event-cnt"] = types.YLeaf{"IptCfgEventCnt", statistics.IptCfgEventCnt}
    statistics.EntityData.Leafs["ipt-cfg-byte-cnt"] = types.YLeaf{"IptCfgByteCnt", statistics.IptCfgByteCnt}
    statistics.EntityData.Leafs["fdt-ipt-desc-cell-cnt"] = types.YLeaf{"FdtIptDescCellCnt", statistics.FdtIptDescCellCnt}
    statistics.EntityData.Leafs["fdt-ire-desc-cell-cnt"] = types.YLeaf{"FdtIreDescCellCnt", statistics.FdtIreDescCellCnt}
    statistics.EntityData.Leafs["fdt-transmitted-data-cells-cnt"] = types.YLeaf{"FdtTransmittedDataCellsCnt", statistics.FdtTransmittedDataCellsCnt}
    statistics.EntityData.Leafs["fdr-p1-cell-in-cnt"] = types.YLeaf{"FdrP1CellInCnt", statistics.FdrP1CellInCnt}
    statistics.EntityData.Leafs["fdr-p2-cell-in-cnt"] = types.YLeaf{"FdrP2CellInCnt", statistics.FdrP2CellInCnt}
    statistics.EntityData.Leafs["fdr-p3-cell-in-cnt"] = types.YLeaf{"FdrP3CellInCnt", statistics.FdrP3CellInCnt}
    statistics.EntityData.Leafs["fdr-cell-in-cnt-total"] = types.YLeaf{"FdrCellInCntTotal", statistics.FdrCellInCntTotal}
    statistics.EntityData.Leafs["fda-cells-in-cnt-p1"] = types.YLeaf{"FdaCellsInCntP1", statistics.FdaCellsInCntP1}
    statistics.EntityData.Leafs["fda-cells-in-cnt-p2"] = types.YLeaf{"FdaCellsInCntP2", statistics.FdaCellsInCntP2}
    statistics.EntityData.Leafs["fda-cells-in-cnt-p3"] = types.YLeaf{"FdaCellsInCntP3", statistics.FdaCellsInCntP3}
    statistics.EntityData.Leafs["fda-cells-in-tdm-cnt"] = types.YLeaf{"FdaCellsInTdmCnt", statistics.FdaCellsInTdmCnt}
    statistics.EntityData.Leafs["fda-cells-in-meshmc-cnt"] = types.YLeaf{"FdaCellsInMeshmcCnt", statistics.FdaCellsInMeshmcCnt}
    statistics.EntityData.Leafs["fda-cells-in-ipt-cnt"] = types.YLeaf{"FdaCellsInIptCnt", statistics.FdaCellsInIptCnt}
    statistics.EntityData.Leafs["fda-cells-out-cnt-p1"] = types.YLeaf{"FdaCellsOutCntP1", statistics.FdaCellsOutCntP1}
    statistics.EntityData.Leafs["fda-cells-out-cnt-p2"] = types.YLeaf{"FdaCellsOutCntP2", statistics.FdaCellsOutCntP2}
    statistics.EntityData.Leafs["fda-cells-out-cnt-p3"] = types.YLeaf{"FdaCellsOutCntP3", statistics.FdaCellsOutCntP3}
    statistics.EntityData.Leafs["fda-cells-out-tdm-cnt"] = types.YLeaf{"FdaCellsOutTdmCnt", statistics.FdaCellsOutTdmCnt}
    statistics.EntityData.Leafs["fda-cells-out-meshmc-cnt"] = types.YLeaf{"FdaCellsOutMeshmcCnt", statistics.FdaCellsOutMeshmcCnt}
    statistics.EntityData.Leafs["fda-cells-out-ipt-cnt"] = types.YLeaf{"FdaCellsOutIptCnt", statistics.FdaCellsOutIptCnt}
    statistics.EntityData.Leafs["fda-egq-drop-cnt"] = types.YLeaf{"FdaEgqDropCnt", statistics.FdaEgqDropCnt}
    statistics.EntityData.Leafs["fda-egq-meshmc-drop-cnt"] = types.YLeaf{"FdaEgqMeshmcDropCnt", statistics.FdaEgqMeshmcDropCnt}
    statistics.EntityData.Leafs["egq-fqp-pkt-cnt"] = types.YLeaf{"EgqFqpPktCnt", statistics.EgqFqpPktCnt}
    statistics.EntityData.Leafs["egq-pqp-uc-pkt-cnt"] = types.YLeaf{"EgqPqpUcPktCnt", statistics.EgqPqpUcPktCnt}
    statistics.EntityData.Leafs["egq-pqp-discard-uc-pkt-cnt"] = types.YLeaf{"EgqPqpDiscardUcPktCnt", statistics.EgqPqpDiscardUcPktCnt}
    statistics.EntityData.Leafs["egq-pqp-uc-bytes-cnt"] = types.YLeaf{"EgqPqpUcBytesCnt", statistics.EgqPqpUcBytesCnt}
    statistics.EntityData.Leafs["egq-pqp-mc-pkt-cnt"] = types.YLeaf{"EgqPqpMcPktCnt", statistics.EgqPqpMcPktCnt}
    statistics.EntityData.Leafs["egq-pqp-discard-mc-pkt-cnt"] = types.YLeaf{"EgqPqpDiscardMcPktCnt", statistics.EgqPqpDiscardMcPktCnt}
    statistics.EntityData.Leafs["egq-pqp-mc-bytes-cnt"] = types.YLeaf{"EgqPqpMcBytesCnt", statistics.EgqPqpMcBytesCnt}
    statistics.EntityData.Leafs["egq-ehp-uc-pkt-cnt"] = types.YLeaf{"EgqEhpUcPktCnt", statistics.EgqEhpUcPktCnt}
    statistics.EntityData.Leafs["egq-ehp-mc-high-pkt-cnt"] = types.YLeaf{"EgqEhpMcHighPktCnt", statistics.EgqEhpMcHighPktCnt}
    statistics.EntityData.Leafs["egq-ehp-mc-low-pkt-cnt"] = types.YLeaf{"EgqEhpMcLowPktCnt", statistics.EgqEhpMcLowPktCnt}
    statistics.EntityData.Leafs["egq-deleted-pkt-cnt"] = types.YLeaf{"EgqDeletedPktCnt", statistics.EgqDeletedPktCnt}
    statistics.EntityData.Leafs["egq-ehp-mc-high-discard-cnt"] = types.YLeaf{"EgqEhpMcHighDiscardCnt", statistics.EgqEhpMcHighDiscardCnt}
    statistics.EntityData.Leafs["egq-ehp-mc-low-discard-cnt"] = types.YLeaf{"EgqEhpMcLowDiscardCnt", statistics.EgqEhpMcLowDiscardCnt}
    statistics.EntityData.Leafs["egq-erpp-lag-pruning-discard-cnt"] = types.YLeaf{"EgqErppLagPruningDiscardCnt", statistics.EgqErppLagPruningDiscardCnt}
    statistics.EntityData.Leafs["egq-erpp-pmf-discard-cnt"] = types.YLeaf{"EgqErppPmfDiscardCnt", statistics.EgqErppPmfDiscardCnt}
    statistics.EntityData.Leafs["egq-erpp-vlan-mbr-discard-cnt"] = types.YLeaf{"EgqErppVlanMbrDiscardCnt", statistics.EgqErppVlanMbrDiscardCnt}
    statistics.EntityData.Leafs["epni-epe-byte-cnt"] = types.YLeaf{"EpniEpeByteCnt", statistics.EpniEpeByteCnt}
    statistics.EntityData.Leafs["epni-epe-pkt-cnt"] = types.YLeaf{"EpniEpePktCnt", statistics.EpniEpePktCnt}
    statistics.EntityData.Leafs["epni-epe-discard-cnt"] = types.YLeaf{"EpniEpeDiscardCnt", statistics.EpniEpeDiscardCnt}
    statistics.EntityData.Leafs["nbi-tx-total-byte-cnt"] = types.YLeaf{"NbiTxTotalByteCnt", statistics.NbiTxTotalByteCnt}
    statistics.EntityData.Leafs["nbi-tx-total-pkt-cnt"] = types.YLeaf{"NbiTxTotalPktCnt", statistics.NbiTxTotalPktCnt}
    return &(statistics.EntityData)
}

// Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds
// Detailed ASIC statistics
type Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Detailed ASIC statistics for a particular NPU. The type is slice of
    // Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId.
    AsicStatisticsDetailForNpuId []Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId
}

func (asicStatisticsDetailForNpuIds *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds) GetEntityData() *types.CommonEntityData {
    asicStatisticsDetailForNpuIds.EntityData.YFilter = asicStatisticsDetailForNpuIds.YFilter
    asicStatisticsDetailForNpuIds.EntityData.YangName = "asic-statistics-detail-for-npu-ids"
    asicStatisticsDetailForNpuIds.EntityData.BundleName = "cisco_ios_xr"
    asicStatisticsDetailForNpuIds.EntityData.ParentYangName = "asic-statistics"
    asicStatisticsDetailForNpuIds.EntityData.SegmentPath = "asic-statistics-detail-for-npu-ids"
    asicStatisticsDetailForNpuIds.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asicStatisticsDetailForNpuIds.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asicStatisticsDetailForNpuIds.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asicStatisticsDetailForNpuIds.EntityData.Children = make(map[string]types.YChild)
    asicStatisticsDetailForNpuIds.EntityData.Children["asic-statistics-detail-for-npu-id"] = types.YChild{"AsicStatisticsDetailForNpuId", nil}
    for i := range asicStatisticsDetailForNpuIds.AsicStatisticsDetailForNpuId {
        asicStatisticsDetailForNpuIds.EntityData.Children[types.GetSegmentPath(&asicStatisticsDetailForNpuIds.AsicStatisticsDetailForNpuId[i])] = types.YChild{"AsicStatisticsDetailForNpuId", &asicStatisticsDetailForNpuIds.AsicStatisticsDetailForNpuId[i]}
    }
    asicStatisticsDetailForNpuIds.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(asicStatisticsDetailForNpuIds.EntityData)
}

// Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId
// Detailed ASIC statistics for a particular
// NPU
type Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. NPU number. The type is interface{} with range:
    // -2147483648..2147483647.
    NpuId interface{}

    // Flag to indicate if data is valid. The type is bool.
    Valid interface{}

    // Rack number. The type is interface{} with range: 0..4294967295.
    RackNumber interface{}

    // Slot number. The type is interface{} with range: 0..4294967295.
    SlotNumber interface{}

    // ASIC instance. The type is interface{} with range: 0..4294967295.
    AsicInstance interface{}

    // Chip version. The type is interface{} with range: 0..65535.
    ChipVersion interface{}

    // Statistics.
    Statistics Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics
}

func (asicStatisticsDetailForNpuId *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId) GetEntityData() *types.CommonEntityData {
    asicStatisticsDetailForNpuId.EntityData.YFilter = asicStatisticsDetailForNpuId.YFilter
    asicStatisticsDetailForNpuId.EntityData.YangName = "asic-statistics-detail-for-npu-id"
    asicStatisticsDetailForNpuId.EntityData.BundleName = "cisco_ios_xr"
    asicStatisticsDetailForNpuId.EntityData.ParentYangName = "asic-statistics-detail-for-npu-ids"
    asicStatisticsDetailForNpuId.EntityData.SegmentPath = "asic-statistics-detail-for-npu-id" + "[npu-id='" + fmt.Sprintf("%v", asicStatisticsDetailForNpuId.NpuId) + "']"
    asicStatisticsDetailForNpuId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asicStatisticsDetailForNpuId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asicStatisticsDetailForNpuId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asicStatisticsDetailForNpuId.EntityData.Children = make(map[string]types.YChild)
    asicStatisticsDetailForNpuId.EntityData.Children["statistics"] = types.YChild{"Statistics", &asicStatisticsDetailForNpuId.Statistics}
    asicStatisticsDetailForNpuId.EntityData.Leafs = make(map[string]types.YLeaf)
    asicStatisticsDetailForNpuId.EntityData.Leafs["npu-id"] = types.YLeaf{"NpuId", asicStatisticsDetailForNpuId.NpuId}
    asicStatisticsDetailForNpuId.EntityData.Leafs["valid"] = types.YLeaf{"Valid", asicStatisticsDetailForNpuId.Valid}
    asicStatisticsDetailForNpuId.EntityData.Leafs["rack-number"] = types.YLeaf{"RackNumber", asicStatisticsDetailForNpuId.RackNumber}
    asicStatisticsDetailForNpuId.EntityData.Leafs["slot-number"] = types.YLeaf{"SlotNumber", asicStatisticsDetailForNpuId.SlotNumber}
    asicStatisticsDetailForNpuId.EntityData.Leafs["asic-instance"] = types.YLeaf{"AsicInstance", asicStatisticsDetailForNpuId.AsicInstance}
    asicStatisticsDetailForNpuId.EntityData.Leafs["chip-version"] = types.YLeaf{"ChipVersion", asicStatisticsDetailForNpuId.ChipVersion}
    return &(asicStatisticsDetailForNpuId.EntityData)
}

// Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics
// Statistics
type Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of blocks. The type is interface{} with range: 0..255.
    NumBlocks interface{}

    // Block information. The type is slice of
    // Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics_BlockInfo.
    BlockInfo []Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics_BlockInfo
}

func (statistics *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "asic-statistics-detail-for-npu-id"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = make(map[string]types.YChild)
    statistics.EntityData.Children["block-info"] = types.YChild{"BlockInfo", nil}
    for i := range statistics.BlockInfo {
        statistics.EntityData.Children[types.GetSegmentPath(&statistics.BlockInfo[i])] = types.YChild{"BlockInfo", &statistics.BlockInfo[i]}
    }
    statistics.EntityData.Leafs = make(map[string]types.YLeaf)
    statistics.EntityData.Leafs["num-blocks"] = types.YLeaf{"NumBlocks", statistics.NumBlocks}
    return &(statistics.EntityData)
}

// Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics_BlockInfo
// Block information
type Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics_BlockInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Block name. The type is string with length: 0..10.
    BlockName interface{}

    // Number of fields. The type is interface{} with range: 0..255.
    NumFields interface{}

    // Field information. The type is slice of
    // Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics_BlockInfo_FieldInfo.
    FieldInfo []Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics_BlockInfo_FieldInfo
}

func (blockInfo *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics_BlockInfo) GetEntityData() *types.CommonEntityData {
    blockInfo.EntityData.YFilter = blockInfo.YFilter
    blockInfo.EntityData.YangName = "block-info"
    blockInfo.EntityData.BundleName = "cisco_ios_xr"
    blockInfo.EntityData.ParentYangName = "statistics"
    blockInfo.EntityData.SegmentPath = "block-info"
    blockInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    blockInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    blockInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    blockInfo.EntityData.Children = make(map[string]types.YChild)
    blockInfo.EntityData.Children["field-info"] = types.YChild{"FieldInfo", nil}
    for i := range blockInfo.FieldInfo {
        blockInfo.EntityData.Children[types.GetSegmentPath(&blockInfo.FieldInfo[i])] = types.YChild{"FieldInfo", &blockInfo.FieldInfo[i]}
    }
    blockInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    blockInfo.EntityData.Leafs["block-name"] = types.YLeaf{"BlockName", blockInfo.BlockName}
    blockInfo.EntityData.Leafs["num-fields"] = types.YLeaf{"NumFields", blockInfo.NumFields}
    return &(blockInfo.EntityData)
}

// Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics_BlockInfo_FieldInfo
// Field information
type Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics_BlockInfo_FieldInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Field name. The type is string with length: 0..80.
    FieldName interface{}

    // Field value. The type is interface{} with range: 0..18446744073709551615.
    FieldValue interface{}

    // Flag to indicate overflow. The type is bool.
    IsOverflow interface{}
}

func (fieldInfo *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics_BlockInfo_FieldInfo) GetEntityData() *types.CommonEntityData {
    fieldInfo.EntityData.YFilter = fieldInfo.YFilter
    fieldInfo.EntityData.YangName = "field-info"
    fieldInfo.EntityData.BundleName = "cisco_ios_xr"
    fieldInfo.EntityData.ParentYangName = "block-info"
    fieldInfo.EntityData.SegmentPath = "field-info"
    fieldInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fieldInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fieldInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fieldInfo.EntityData.Children = make(map[string]types.YChild)
    fieldInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    fieldInfo.EntityData.Leafs["field-name"] = types.YLeaf{"FieldName", fieldInfo.FieldName}
    fieldInfo.EntityData.Leafs["field-value"] = types.YLeaf{"FieldValue", fieldInfo.FieldValue}
    fieldInfo.EntityData.Leafs["is-overflow"] = types.YLeaf{"IsOverflow", fieldInfo.IsOverflow}
    return &(fieldInfo.EntityData)
}

// Dpa_Stats_Nodes_Node_NpuNumbers
// Ingress Stats
type Dpa_Stats_Nodes_Node_NpuNumbers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Stats for a particular npu. The type is slice of
    // Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber.
    NpuNumber []Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber
}

func (npuNumbers *Dpa_Stats_Nodes_Node_NpuNumbers) GetEntityData() *types.CommonEntityData {
    npuNumbers.EntityData.YFilter = npuNumbers.YFilter
    npuNumbers.EntityData.YangName = "npu-numbers"
    npuNumbers.EntityData.BundleName = "cisco_ios_xr"
    npuNumbers.EntityData.ParentYangName = "node"
    npuNumbers.EntityData.SegmentPath = "npu-numbers"
    npuNumbers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    npuNumbers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    npuNumbers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    npuNumbers.EntityData.Children = make(map[string]types.YChild)
    npuNumbers.EntityData.Children["npu-number"] = types.YChild{"NpuNumber", nil}
    for i := range npuNumbers.NpuNumber {
        npuNumbers.EntityData.Children[types.GetSegmentPath(&npuNumbers.NpuNumber[i])] = types.YChild{"NpuNumber", &npuNumbers.NpuNumber[i]}
    }
    npuNumbers.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(npuNumbers.EntityData)
}

// Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber
// Stats for a particular npu
type Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Npu number. The type is interface{} with range:
    // -2147483648..2147483647.
    NpuId interface{}

    // show npu specific voq or trap stats.
    Display Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display
}

func (npuNumber *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber) GetEntityData() *types.CommonEntityData {
    npuNumber.EntityData.YFilter = npuNumber.YFilter
    npuNumber.EntityData.YangName = "npu-number"
    npuNumber.EntityData.BundleName = "cisco_ios_xr"
    npuNumber.EntityData.ParentYangName = "npu-numbers"
    npuNumber.EntityData.SegmentPath = "npu-number" + "[npu-id='" + fmt.Sprintf("%v", npuNumber.NpuId) + "']"
    npuNumber.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    npuNumber.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    npuNumber.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    npuNumber.EntityData.Children = make(map[string]types.YChild)
    npuNumber.EntityData.Children["display"] = types.YChild{"Display", &npuNumber.Display}
    npuNumber.EntityData.Leafs = make(map[string]types.YLeaf)
    npuNumber.EntityData.Leafs["npu-id"] = types.YLeaf{"NpuId", npuNumber.NpuId}
    return &(npuNumber.EntityData)
}

// Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display
// show npu specific voq or trap stats
type Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Voq stats grouped by voq base numbers.
    BaseNumbers Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers

    // Trap stats for a particular npu.
    TrapIds Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_TrapIds

    // Voq stats grouped by interface handle.
    InterfaceHandles Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles
}

func (display *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display) GetEntityData() *types.CommonEntityData {
    display.EntityData.YFilter = display.YFilter
    display.EntityData.YangName = "display"
    display.EntityData.BundleName = "cisco_ios_xr"
    display.EntityData.ParentYangName = "npu-number"
    display.EntityData.SegmentPath = "display"
    display.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    display.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    display.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    display.EntityData.Children = make(map[string]types.YChild)
    display.EntityData.Children["base-numbers"] = types.YChild{"BaseNumbers", &display.BaseNumbers}
    display.EntityData.Children["trap-ids"] = types.YChild{"TrapIds", &display.TrapIds}
    display.EntityData.Children["interface-handles"] = types.YChild{"InterfaceHandles", &display.InterfaceHandles}
    display.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(display.EntityData)
}

// Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers
// Voq stats grouped by voq base numbers
type Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Voq Base Number for a particular voq. The type is slice of
    // Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers_BaseNumber.
    BaseNumber []Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers_BaseNumber
}

func (baseNumbers *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers) GetEntityData() *types.CommonEntityData {
    baseNumbers.EntityData.YFilter = baseNumbers.YFilter
    baseNumbers.EntityData.YangName = "base-numbers"
    baseNumbers.EntityData.BundleName = "cisco_ios_xr"
    baseNumbers.EntityData.ParentYangName = "display"
    baseNumbers.EntityData.SegmentPath = "base-numbers"
    baseNumbers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    baseNumbers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    baseNumbers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    baseNumbers.EntityData.Children = make(map[string]types.YChild)
    baseNumbers.EntityData.Children["base-number"] = types.YChild{"BaseNumber", nil}
    for i := range baseNumbers.BaseNumber {
        baseNumbers.EntityData.Children[types.GetSegmentPath(&baseNumbers.BaseNumber[i])] = types.YChild{"BaseNumber", &baseNumbers.BaseNumber[i]}
    }
    baseNumbers.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(baseNumbers.EntityData)
}

// Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers_BaseNumber
// Voq Base Number for a particular voq
type Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers_BaseNumber struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface handle. The type is interface{} with
    // range: 0..4294967295.
    BaseNumber interface{}

    // Flag to indicate if port is in use. The type is bool.
    InUse interface{}

    // Rack of port. The type is interface{} with range: 0..255.
    RackNum interface{}

    // Slot of port. The type is interface{} with range: 0..255.
    SlotNum interface{}

    // NPU of port. The type is interface{} with range: 0..255.
    NpuNum interface{}

    // NPU core of port. The type is interface{} with range: 0..255.
    NpuCore interface{}

    // Port Number of port. The type is interface{} with range: 0..255.
    PortNum interface{}

    // IfHandle of port. The type is interface{} with range: 0..4294967295.
    IfHandle interface{}

    // System port of port. The type is interface{} with range: 0..4294967295.
    SysPort interface{}

    // PP Port number of port. The type is interface{} with range: 0..4294967295.
    PpPort interface{}

    // Port speed of port. The type is interface{} with range: 0..4294967295.
    PortSpeed interface{}

    // Voq Base number of port. The type is interface{} with range: 0..4294967295.
    VoqBase interface{}

    // Connector id of port. The type is interface{} with range: 0..4294967295.
    ConnectorId interface{}

    // Flag to indicate if port is local to the node. The type is bool.
    IsLocalPort interface{}

    // Keeps a record of the received and dropped packets and bytes on the port.
    // The type is slice of
    // Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers_BaseNumber_VoqStat.
    VoqStat []Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers_BaseNumber_VoqStat
}

func (baseNumber *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers_BaseNumber) GetEntityData() *types.CommonEntityData {
    baseNumber.EntityData.YFilter = baseNumber.YFilter
    baseNumber.EntityData.YangName = "base-number"
    baseNumber.EntityData.BundleName = "cisco_ios_xr"
    baseNumber.EntityData.ParentYangName = "base-numbers"
    baseNumber.EntityData.SegmentPath = "base-number" + "[base-number='" + fmt.Sprintf("%v", baseNumber.BaseNumber) + "']"
    baseNumber.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    baseNumber.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    baseNumber.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    baseNumber.EntityData.Children = make(map[string]types.YChild)
    baseNumber.EntityData.Children["voq-stat"] = types.YChild{"VoqStat", nil}
    for i := range baseNumber.VoqStat {
        baseNumber.EntityData.Children[types.GetSegmentPath(&baseNumber.VoqStat[i])] = types.YChild{"VoqStat", &baseNumber.VoqStat[i]}
    }
    baseNumber.EntityData.Leafs = make(map[string]types.YLeaf)
    baseNumber.EntityData.Leafs["base-number"] = types.YLeaf{"BaseNumber", baseNumber.BaseNumber}
    baseNumber.EntityData.Leafs["in-use"] = types.YLeaf{"InUse", baseNumber.InUse}
    baseNumber.EntityData.Leafs["rack-num"] = types.YLeaf{"RackNum", baseNumber.RackNum}
    baseNumber.EntityData.Leafs["slot-num"] = types.YLeaf{"SlotNum", baseNumber.SlotNum}
    baseNumber.EntityData.Leafs["npu-num"] = types.YLeaf{"NpuNum", baseNumber.NpuNum}
    baseNumber.EntityData.Leafs["npu-core"] = types.YLeaf{"NpuCore", baseNumber.NpuCore}
    baseNumber.EntityData.Leafs["port-num"] = types.YLeaf{"PortNum", baseNumber.PortNum}
    baseNumber.EntityData.Leafs["if-handle"] = types.YLeaf{"IfHandle", baseNumber.IfHandle}
    baseNumber.EntityData.Leafs["sys-port"] = types.YLeaf{"SysPort", baseNumber.SysPort}
    baseNumber.EntityData.Leafs["pp-port"] = types.YLeaf{"PpPort", baseNumber.PpPort}
    baseNumber.EntityData.Leafs["port-speed"] = types.YLeaf{"PortSpeed", baseNumber.PortSpeed}
    baseNumber.EntityData.Leafs["voq-base"] = types.YLeaf{"VoqBase", baseNumber.VoqBase}
    baseNumber.EntityData.Leafs["connector-id"] = types.YLeaf{"ConnectorId", baseNumber.ConnectorId}
    baseNumber.EntityData.Leafs["is-local-port"] = types.YLeaf{"IsLocalPort", baseNumber.IsLocalPort}
    return &(baseNumber.EntityData)
}

// Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers_BaseNumber_VoqStat
// Keeps a record of the received and dropped
// packets and bytes on the port
type Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers_BaseNumber_VoqStat struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Bytes Received on the port. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ReceivedBytes interface{}

    // Packets Received on the port. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPackets interface{}

    // Bytes Dropped on the port. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    DroppedBytes interface{}

    // Packets Dropeed on the port. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}
}

func (voqStat *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers_BaseNumber_VoqStat) GetEntityData() *types.CommonEntityData {
    voqStat.EntityData.YFilter = voqStat.YFilter
    voqStat.EntityData.YangName = "voq-stat"
    voqStat.EntityData.BundleName = "cisco_ios_xr"
    voqStat.EntityData.ParentYangName = "base-number"
    voqStat.EntityData.SegmentPath = "voq-stat"
    voqStat.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    voqStat.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    voqStat.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    voqStat.EntityData.Children = make(map[string]types.YChild)
    voqStat.EntityData.Leafs = make(map[string]types.YLeaf)
    voqStat.EntityData.Leafs["received-bytes"] = types.YLeaf{"ReceivedBytes", voqStat.ReceivedBytes}
    voqStat.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", voqStat.ReceivedPackets}
    voqStat.EntityData.Leafs["dropped-bytes"] = types.YLeaf{"DroppedBytes", voqStat.DroppedBytes}
    voqStat.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", voqStat.DroppedPackets}
    return &(voqStat.EntityData)
}

// Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_TrapIds
// Trap stats for a particular npu
type Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_TrapIds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Filter by specific trap id. The type is slice of
    // Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_TrapIds_TrapId.
    TrapId []Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_TrapIds_TrapId
}

func (trapIds *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_TrapIds) GetEntityData() *types.CommonEntityData {
    trapIds.EntityData.YFilter = trapIds.YFilter
    trapIds.EntityData.YangName = "trap-ids"
    trapIds.EntityData.BundleName = "cisco_ios_xr"
    trapIds.EntityData.ParentYangName = "display"
    trapIds.EntityData.SegmentPath = "trap-ids"
    trapIds.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trapIds.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trapIds.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trapIds.EntityData.Children = make(map[string]types.YChild)
    trapIds.EntityData.Children["trap-id"] = types.YChild{"TrapId", nil}
    for i := range trapIds.TrapId {
        trapIds.EntityData.Children[types.GetSegmentPath(&trapIds.TrapId[i])] = types.YChild{"TrapId", &trapIds.TrapId[i]}
    }
    trapIds.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(trapIds.EntityData)
}

// Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_TrapIds_TrapId
// Filter by specific trap id
type Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_TrapIds_TrapId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Trap ID. The type is interface{} with range:
    // 0..4294967295.
    TrapId interface{}

    // Trap Strength of the trap. The type is interface{} with range:
    // 0..4294967295.
    TrapStrength interface{}

    // Priority of the trap. The type is interface{} with range: 0..4294967295.
    Priority interface{}

    // Id of the trap. The type is interface{} with range: 0..4294967295.
    TrapIdXr interface{}

    // Gport of the trap. The type is interface{} with range: 0..4294967295.
    Gport interface{}

    // Fec id of the trap. The type is interface{} with range: 0..4294967295.
    FecId interface{}

    // Id of the policer on the trap. The type is interface{} with range:
    // 0..4294967295.
    PolicerId interface{}

    // Stats Id of the trap. The type is interface{} with range: 0..4294967295.
    StatsId interface{}

    // Encap Id of the trap. The type is interface{} with range: 0..4294967295.
    EncapId interface{}

    // McGroup of the trap. The type is interface{} with range: 0..4294967295.
    McGroup interface{}

    // Name String of the trap. The type is string.
    TrapString interface{}

    // Id for internal use. The type is interface{} with range: 0..4294967295.
    Id interface{}

    // Offset for internal use. The type is interface{} with range:
    // 0..18446744073709551615.
    Offset interface{}

    // NpuId on which trap is enabled. The type is interface{} with range:
    // 0..18446744073709551615.
    NpuId interface{}

    // Number of packets dropped after hitting the trap. The type is interface{}
    // with range: 0..18446744073709551615.
    PacketDropped interface{}

    // Number of packets accepted after hitting the trap. The type is interface{}
    // with range: 0..18446744073709551615.
    PacketAccepted interface{}
}

func (trapId *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_TrapIds_TrapId) GetEntityData() *types.CommonEntityData {
    trapId.EntityData.YFilter = trapId.YFilter
    trapId.EntityData.YangName = "trap-id"
    trapId.EntityData.BundleName = "cisco_ios_xr"
    trapId.EntityData.ParentYangName = "trap-ids"
    trapId.EntityData.SegmentPath = "trap-id" + "[trap-id='" + fmt.Sprintf("%v", trapId.TrapId) + "']"
    trapId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trapId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trapId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trapId.EntityData.Children = make(map[string]types.YChild)
    trapId.EntityData.Leafs = make(map[string]types.YLeaf)
    trapId.EntityData.Leafs["trap-id"] = types.YLeaf{"TrapId", trapId.TrapId}
    trapId.EntityData.Leafs["trap-strength"] = types.YLeaf{"TrapStrength", trapId.TrapStrength}
    trapId.EntityData.Leafs["priority"] = types.YLeaf{"Priority", trapId.Priority}
    trapId.EntityData.Leafs["trap-id-xr"] = types.YLeaf{"TrapIdXr", trapId.TrapIdXr}
    trapId.EntityData.Leafs["gport"] = types.YLeaf{"Gport", trapId.Gport}
    trapId.EntityData.Leafs["fec-id"] = types.YLeaf{"FecId", trapId.FecId}
    trapId.EntityData.Leafs["policer-id"] = types.YLeaf{"PolicerId", trapId.PolicerId}
    trapId.EntityData.Leafs["stats-id"] = types.YLeaf{"StatsId", trapId.StatsId}
    trapId.EntityData.Leafs["encap-id"] = types.YLeaf{"EncapId", trapId.EncapId}
    trapId.EntityData.Leafs["mc-group"] = types.YLeaf{"McGroup", trapId.McGroup}
    trapId.EntityData.Leafs["trap-string"] = types.YLeaf{"TrapString", trapId.TrapString}
    trapId.EntityData.Leafs["id"] = types.YLeaf{"Id", trapId.Id}
    trapId.EntityData.Leafs["offset"] = types.YLeaf{"Offset", trapId.Offset}
    trapId.EntityData.Leafs["npu-id"] = types.YLeaf{"NpuId", trapId.NpuId}
    trapId.EntityData.Leafs["packet-dropped"] = types.YLeaf{"PacketDropped", trapId.PacketDropped}
    trapId.EntityData.Leafs["packet-accepted"] = types.YLeaf{"PacketAccepted", trapId.PacketAccepted}
    return &(trapId.EntityData)
}

// Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles
// Voq stats grouped by interface handle
type Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Voq stats for a particular interface handle. The type is slice of
    // Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles_InterfaceHandle.
    InterfaceHandle []Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles_InterfaceHandle
}

func (interfaceHandles *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles) GetEntityData() *types.CommonEntityData {
    interfaceHandles.EntityData.YFilter = interfaceHandles.YFilter
    interfaceHandles.EntityData.YangName = "interface-handles"
    interfaceHandles.EntityData.BundleName = "cisco_ios_xr"
    interfaceHandles.EntityData.ParentYangName = "display"
    interfaceHandles.EntityData.SegmentPath = "interface-handles"
    interfaceHandles.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceHandles.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceHandles.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceHandles.EntityData.Children = make(map[string]types.YChild)
    interfaceHandles.EntityData.Children["interface-handle"] = types.YChild{"InterfaceHandle", nil}
    for i := range interfaceHandles.InterfaceHandle {
        interfaceHandles.EntityData.Children[types.GetSegmentPath(&interfaceHandles.InterfaceHandle[i])] = types.YChild{"InterfaceHandle", &interfaceHandles.InterfaceHandle[i]}
    }
    interfaceHandles.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaceHandles.EntityData)
}

// Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles_InterfaceHandle
// Voq stats for a particular interface
// handle
type Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles_InterfaceHandle struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Handle. The type is interface{} with
    // range: 0..4294967295.
    InterfaceHandle interface{}

    // Flag to indicate if port is in use. The type is bool.
    InUse interface{}

    // Rack of port. The type is interface{} with range: 0..255.
    RackNum interface{}

    // Slot of port. The type is interface{} with range: 0..255.
    SlotNum interface{}

    // NPU of port. The type is interface{} with range: 0..255.
    NpuNum interface{}

    // NPU core of port. The type is interface{} with range: 0..255.
    NpuCore interface{}

    // Port Number of port. The type is interface{} with range: 0..255.
    PortNum interface{}

    // IfHandle of port. The type is interface{} with range: 0..4294967295.
    IfHandle interface{}

    // System port of port. The type is interface{} with range: 0..4294967295.
    SysPort interface{}

    // PP Port number of port. The type is interface{} with range: 0..4294967295.
    PpPort interface{}

    // Port speed of port. The type is interface{} with range: 0..4294967295.
    PortSpeed interface{}

    // Voq Base number of port. The type is interface{} with range: 0..4294967295.
    VoqBase interface{}

    // Connector id of port. The type is interface{} with range: 0..4294967295.
    ConnectorId interface{}

    // Flag to indicate if port is local to the node. The type is bool.
    IsLocalPort interface{}

    // Keeps a record of the received and dropped packets and bytes on the port.
    // The type is slice of
    // Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles_InterfaceHandle_VoqStat.
    VoqStat []Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles_InterfaceHandle_VoqStat
}

func (interfaceHandle *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles_InterfaceHandle) GetEntityData() *types.CommonEntityData {
    interfaceHandle.EntityData.YFilter = interfaceHandle.YFilter
    interfaceHandle.EntityData.YangName = "interface-handle"
    interfaceHandle.EntityData.BundleName = "cisco_ios_xr"
    interfaceHandle.EntityData.ParentYangName = "interface-handles"
    interfaceHandle.EntityData.SegmentPath = "interface-handle" + "[interface-handle='" + fmt.Sprintf("%v", interfaceHandle.InterfaceHandle) + "']"
    interfaceHandle.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceHandle.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceHandle.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceHandle.EntityData.Children = make(map[string]types.YChild)
    interfaceHandle.EntityData.Children["voq-stat"] = types.YChild{"VoqStat", nil}
    for i := range interfaceHandle.VoqStat {
        interfaceHandle.EntityData.Children[types.GetSegmentPath(&interfaceHandle.VoqStat[i])] = types.YChild{"VoqStat", &interfaceHandle.VoqStat[i]}
    }
    interfaceHandle.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceHandle.EntityData.Leafs["interface-handle"] = types.YLeaf{"InterfaceHandle", interfaceHandle.InterfaceHandle}
    interfaceHandle.EntityData.Leafs["in-use"] = types.YLeaf{"InUse", interfaceHandle.InUse}
    interfaceHandle.EntityData.Leafs["rack-num"] = types.YLeaf{"RackNum", interfaceHandle.RackNum}
    interfaceHandle.EntityData.Leafs["slot-num"] = types.YLeaf{"SlotNum", interfaceHandle.SlotNum}
    interfaceHandle.EntityData.Leafs["npu-num"] = types.YLeaf{"NpuNum", interfaceHandle.NpuNum}
    interfaceHandle.EntityData.Leafs["npu-core"] = types.YLeaf{"NpuCore", interfaceHandle.NpuCore}
    interfaceHandle.EntityData.Leafs["port-num"] = types.YLeaf{"PortNum", interfaceHandle.PortNum}
    interfaceHandle.EntityData.Leafs["if-handle"] = types.YLeaf{"IfHandle", interfaceHandle.IfHandle}
    interfaceHandle.EntityData.Leafs["sys-port"] = types.YLeaf{"SysPort", interfaceHandle.SysPort}
    interfaceHandle.EntityData.Leafs["pp-port"] = types.YLeaf{"PpPort", interfaceHandle.PpPort}
    interfaceHandle.EntityData.Leafs["port-speed"] = types.YLeaf{"PortSpeed", interfaceHandle.PortSpeed}
    interfaceHandle.EntityData.Leafs["voq-base"] = types.YLeaf{"VoqBase", interfaceHandle.VoqBase}
    interfaceHandle.EntityData.Leafs["connector-id"] = types.YLeaf{"ConnectorId", interfaceHandle.ConnectorId}
    interfaceHandle.EntityData.Leafs["is-local-port"] = types.YLeaf{"IsLocalPort", interfaceHandle.IsLocalPort}
    return &(interfaceHandle.EntityData)
}

// Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles_InterfaceHandle_VoqStat
// Keeps a record of the received and dropped
// packets and bytes on the port
type Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles_InterfaceHandle_VoqStat struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Bytes Received on the port. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ReceivedBytes interface{}

    // Packets Received on the port. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPackets interface{}

    // Bytes Dropped on the port. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    DroppedBytes interface{}

    // Packets Dropeed on the port. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}
}

func (voqStat *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles_InterfaceHandle_VoqStat) GetEntityData() *types.CommonEntityData {
    voqStat.EntityData.YFilter = voqStat.YFilter
    voqStat.EntityData.YangName = "voq-stat"
    voqStat.EntityData.BundleName = "cisco_ios_xr"
    voqStat.EntityData.ParentYangName = "interface-handle"
    voqStat.EntityData.SegmentPath = "voq-stat"
    voqStat.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    voqStat.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    voqStat.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    voqStat.EntityData.Children = make(map[string]types.YChild)
    voqStat.EntityData.Leafs = make(map[string]types.YLeaf)
    voqStat.EntityData.Leafs["received-bytes"] = types.YLeaf{"ReceivedBytes", voqStat.ReceivedBytes}
    voqStat.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", voqStat.ReceivedPackets}
    voqStat.EntityData.Leafs["dropped-bytes"] = types.YLeaf{"DroppedBytes", voqStat.DroppedBytes}
    voqStat.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", voqStat.DroppedPackets}
    return &(voqStat.EntityData)
}

