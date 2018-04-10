// This module contains a collection of generally useful
// derived YANG data types.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package fretta_bcm_dpa_resources_datatypes

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package fretta_bcm_dpa_resources_datatypes"))
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

