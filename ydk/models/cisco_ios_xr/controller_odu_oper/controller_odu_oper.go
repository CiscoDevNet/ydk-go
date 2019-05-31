// This module contains a collection of YANG definitions
// for Cisco IOS-XR controller-odu package operational data.
// 
// This module contains definitions
// for the following management objects:
//   odu: ODU operational data
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package controller_odu_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package controller_odu_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-controller-odu-oper odu}", reflect.TypeOf(Odu{}))
    ydk.RegisterEntity("Cisco-IOS-XR-controller-odu-oper:odu", reflect.TypeOf(Odu{}))
}

// OduAinsStateEt represents Odu ains state et
type OduAinsStateEt string

const (
    // None
    OduAinsStateEt_none OduAinsStateEt = "none"

    // Running
    OduAinsStateEt_active_running OduAinsStateEt = "active-running"

    // Pending
    OduAinsStateEt_active_pending OduAinsStateEt = "active-pending"
)

// OduEtherMapPingEt represents Odu ether map ping et
type OduEtherMapPingEt string

const (
    // None
    OduEtherMapPingEt_none OduEtherMapPingEt = "none"

    // GfpF
    OduEtherMapPingEt_gfp OduEtherMapPingEt = "gfp"

    // Amp
    OduEtherMapPingEt_amp OduEtherMapPingEt = "amp"

    // Bmp
    OduEtherMapPingEt_bmp OduEtherMapPingEt = "bmp"

    // Gmp
    OduEtherMapPingEt_gmp OduEtherMapPingEt = "gmp"

    // Wis
    OduEtherMapPingEt_wis OduEtherMapPingEt = "wis"

    // GfpF Ext
    OduEtherMapPingEt_gfp_ext OduEtherMapPingEt = "gfp-ext"
)

// OduTcmMode represents Odu tcm mode
type OduTcmMode string

const (
    // Transparent
    OduTcmMode_odu_tcm_mode_trans_parent OduTcmMode = "odu-tcm-mode-trans-parent"

    // Non-Intrusive Monitor
    OduTcmMode_nim OduTcmMode = "nim"

    // Operational
    OduTcmMode_oper OduTcmMode = "oper"
)

// OduTcmPerMon represents Odu tcm per mon
type OduTcmPerMon string

const (
    // Disable
    OduTcmPerMon_disable OduTcmPerMon = "disable"

    // Enable
    OduTcmPerMon_enable OduTcmPerMon = "enable"
)

// OduTcmStateEt represents Odu tcm state et
type OduTcmStateEt string

const (
    // Disable
    OduTcmStateEt_disable OduTcmStateEt = "disable"

    // Enable
    OduTcmStateEt_enable OduTcmStateEt = "enable"
)

// GmplsTtiMode represents Gmpls tti mode
type GmplsTtiMode string

const (
    // Not Set
    GmplsTtiMode_gmpls_odu_tti_mode_none GmplsTtiMode = "gmpls-odu-tti-mode-none"

    // Section Monitoring
    GmplsTtiMode_gmpls_odu_tti_mode_sm GmplsTtiMode = "gmpls-odu-tti-mode-sm"

    // Path Monitoring
    GmplsTtiMode_gmpls_odu_tti_mode_pm GmplsTtiMode = "gmpls-odu-tti-mode-pm"

    // Tandem Connection Monitoring
    GmplsTtiMode_gmpls_odu_tti_mode_tcm GmplsTtiMode = "gmpls-odu-tti-mode-tcm"
)

// OduPmMode represents Odu pm mode
type OduPmMode string

const (
    // Non-Intrusive Monitor
    OduPmMode_nim OduPmMode = "nim"

    // Operational
    OduPmMode_oper OduPmMode = "oper"
)

// OduPmCaEt represents Odu pm ca et
type OduPmCaEt string

const (
    // Disable
    OduPmCaEt_disable OduPmCaEt = "disable"

    // Enable
    OduPmCaEt_enable OduPmCaEt = "enable"
)

// OduPerMon represents Odu per mon
type OduPerMon string

const (
    // Disable
    OduPerMon_disable OduPerMon = "disable"

    // Enable
    OduPerMon_enable OduPerMon = "enable"
)

// DpProgrammed represents Dp programmed
type DpProgrammed string

const (
    // DP not programmed
    DpProgrammed_dp_not_programmed DpProgrammed = "dp-not-programmed"

    // DP programmed
    DpProgrammed_dp_programmed_success DpProgrammed = "dp-programmed-success"

    // ENDPT FIRST CHANNELIZED
    DpProgrammed_end_pt_first_channel_ized DpProgrammed = "end-pt-first-channel-ized"

    // ENDPT SECOND CHANNELIZED
    DpProgrammed_end_pt_se_cond_channel_ized DpProgrammed = "end-pt-se-cond-channel-ized"

    // ENDPT FIRST CROSS CONNECTED
    DpProgrammed_end_pt_first_cross_connected DpProgrammed = "end-pt-first-cross-connected"

    // ENDPT SECOND CROSS CONNECTED
    DpProgrammed_end_pt_se_cond_cross_connected DpProgrammed = "end-pt-se-cond-cross-connected"

    // ENDPT FIRST OPEN CONNECTED
    DpProgrammed_end_pt_first_open_connected DpProgrammed = "end-pt-first-open-connected"

    // ENDPT SECOND OPEN CONNECTED
    DpProgrammed_end_pt_se_cond_open_connected DpProgrammed = "end-pt-se-cond-open-connected"

    // ENDPT FIRST LOOPBACKED
    DpProgrammed_end_pt_first_loop_back_ed DpProgrammed = "end-pt-first-loop-back-ed"

    // ENDPT SECOND LOOPBACKED
    DpProgrammed_end_pt_se_cond_loop_back_ed DpProgrammed = "end-pt-se-cond-loop-back-ed"

    // ENDPT ODU TYPE MISMATCH
    DpProgrammed_end_pt_odu_type_mismatch DpProgrammed = "end-pt-odu-type-mismatch"

    // XCONNECT NOT SET
    DpProgrammed_xc_not_set DpProgrammed = "xc-not-set"
)

// OtmMplsLibC represents Otm mpls lib c
type OtmMplsLibC string

const (
    // NULL
    OtmMplsLibC_otm_mpls_lib_c_type_null OtmMplsLibC = "otm-mpls-lib-c-type-null"

    // IPV4
    OtmMplsLibC_otm_mpls_lib_c_type_ipv4 OtmMplsLibC = "otm-mpls-lib-c-type-ipv4"

    // IPV4 P2P TUNNEL
    OtmMplsLibC_otm_mpls_lib_c_type_ipv4_p2p_tunnel OtmMplsLibC = "otm-mpls-lib-c-type-ipv4-p2p-tunnel"

    // IPV6 P2P TUNNEL
    OtmMplsLibC_otm_mpls_lib_c_type_ipv6_p2p_tunnel OtmMplsLibC = "otm-mpls-lib-c-type-ipv6-p2p-tunnel"

    // IPV4 UNI
    OtmMplsLibC_otm_mpls_lib_c_type_ipv4_uni OtmMplsLibC = "otm-mpls-lib-c-type-ipv4-uni"

    // IPV4 P2MP TUNNEL
    OtmMplsLibC_otm_mpls_lib_c_type_ipv4_p2mp_tunnel OtmMplsLibC = "otm-mpls-lib-c-type-ipv4-p2mp-tunnel"

    // IPV6 P2MP TUNNEL
    OtmMplsLibC_otm_mpls_lib_c_type_ipv6_p2mp_tunnel OtmMplsLibC = "otm-mpls-lib-c-type-ipv6-p2mp-tunnel"

    // IPV4 TP TUNNEL
    OtmMplsLibC_otm_mpls_lib_c_type_ipv4_tp_tunnel OtmMplsLibC = "otm-mpls-lib-c-type-ipv4-tp-tunnel"

    // IPV6 TP TUNNEL
    OtmMplsLibC_otm_mpls_lib_c_type_ipv6_tp_tunnel OtmMplsLibC = "otm-mpls-lib-c-type-ipv6-tp-tunnel"
)

// OtmTeTunnelInfo represents Otm te tunnel info
type OtmTeTunnelInfo string

const (
    // NONE
    OtmTeTunnelInfo_otm_te_info_none OtmTeTunnelInfo = "otm-te-info-none"

    // S2L
    OtmTeTunnelInfo_otm_te_info_s2l OtmTeTunnelInfo = "otm-te-info-s2l"

    // ID
    OtmTeTunnelInfo_otm_te_info_tunnel_id OtmTeTunnelInfo = "otm-te-info-tunnel-id"

    // MAT
    OtmTeTunnelInfo_otm_te_info_passive_match OtmTeTunnelInfo = "otm-te-info-passive-match"
)

// OtmOpticalRmCtxtRm represents Otm optical rm ctxt rm
type OtmOpticalRmCtxtRm string

const (
    // NONE
    OtmOpticalRmCtxtRm_otm_opt_rm_ctxt_rm_none OtmOpticalRmCtxtRm = "otm-opt-rm-ctxt-rm-none"

    // WDM
    OtmOpticalRmCtxtRm_otm_opt_rm_ctxt_rm_wdm OtmOpticalRmCtxtRm = "otm-opt-rm-ctxt-rm-wdm"

    // FSC
    OtmOpticalRmCtxtRm_otm_opt_rm_ctxt_rm_fsc OtmOpticalRmCtxtRm = "otm-opt-rm-ctxt-rm-fsc"

    // TDM
    OtmOpticalRmCtxtRm_otm_opt_rm_ctxt_rm_tdm OtmOpticalRmCtxtRm = "otm-opt-rm-ctxt-rm-tdm"

    // G709 OTN
    OtmOpticalRmCtxtRm_otm_opt_rm_ctxt_rm_g709_otn OtmOpticalRmCtxtRm = "otm-opt-rm-ctxt-rm-g709-otn"

    // LAST
    OtmOpticalRmCtxtRm_otm_optical_rm_ctxt_rm_last OtmOpticalRmCtxtRm = "otm-optical-rm-ctxt-rm-last"
)

// OtmOpticalRmCtxt represents Otm optical rm ctxt
type OtmOpticalRmCtxt string

const (
    // NONE
    OtmOpticalRmCtxt_otm_opt_rm_ctxt_type_none OtmOpticalRmCtxt = "otm-opt-rm-ctxt-type-none"

    // DOWNSTREAM RW ADD
    OtmOpticalRmCtxt_otm_opt_rm_ctxt_type_down_stream_rw_add OtmOpticalRmCtxt = "otm-opt-rm-ctxt-type-down-stream-rw-add"

    // UPSTREAM RW ADD
    OtmOpticalRmCtxt_otm_opt_rm_ctxt_type_up_stream_rw_add OtmOpticalRmCtxt = "otm-opt-rm-ctxt-type-up-stream-rw-add"

    // DOWNSTREAM RW DEL
    OtmOpticalRmCtxt_otm_opt_rm_ctxt_type_down_stream_rw_del OtmOpticalRmCtxt = "otm-opt-rm-ctxt-type-down-stream-rw-del"

    // UPSTREAM RW DEL
    OtmOpticalRmCtxt_otm_opt_rm_ctxt_type_up_stream_rw_del OtmOpticalRmCtxt = "otm-opt-rm-ctxt-type-up-stream-rw-del"

    // DOWNSTREAM LBL GET
    OtmOpticalRmCtxt_otm_opt_rm_ctxt_type_down_stream_lbl_get OtmOpticalRmCtxt = "otm-opt-rm-ctxt-type-down-stream-lbl-get"

    // UPSTREAM LBL GET
    OtmOpticalRmCtxt_otm_opt_rm_ctxt_type_up_stream_lbl_get OtmOpticalRmCtxt = "otm-opt-rm-ctxt-type-up-stream-lbl-get"

    // DOWNSTREAM LBL REL
    OtmOpticalRmCtxt_otm_opt_rm_ctxt_type_down_stream_lbl_rel OtmOpticalRmCtxt = "otm-opt-rm-ctxt-type-down-stream-lbl-rel"

    // UPSTREAM LBL REL
    OtmOpticalRmCtxt_otm_opt_rm_ctxt_type_up_stream_lbl_rel OtmOpticalRmCtxt = "otm-opt-rm-ctxt-type-up-stream-lbl-rel"

    // ENDPOINT RW ADD
    OtmOpticalRmCtxt_otm_opt_rm_ctxt_type_end_point_rw_add OtmOpticalRmCtxt = "otm-opt-rm-ctxt-type-end-point-rw-add"

    // ENDPOINT RW DEL
    OtmOpticalRmCtxt_otm_opt_rm_ctxt_type_end_point_rw_del OtmOpticalRmCtxt = "otm-opt-rm-ctxt-type-end-point-rw-del"

    // ODU GROUP ADD
    OtmOpticalRmCtxt_otm_opt_rm_ctxt_type_odu_group_add OtmOpticalRmCtxt = "otm-opt-rm-ctxt-type-odu-group-add"

    // ODU GROUP DEL
    OtmOpticalRmCtxt_otm_opt_rm_ctxt_type_odu_group_del OtmOpticalRmCtxt = "otm-opt-rm-ctxt-type-odu-group-del"

    // LAST
    OtmOpticalRmCtxt_otm_optical_rm_ctxt_type_last OtmOpticalRmCtxt = "otm-optical-rm-ctxt-type-last"
)

// OduTsGEt represents Odu ts g et
type OduTsGEt string

const (
    // 1.25G
    OduTsGEt_one_dot_two_five_g OduTsGEt = "one-dot-two-five-g"

    // 2.5G
    OduTsGEt_two_dot_five_g OduTsGEt = "two-dot-five-g"

    // NA
    OduTsGEt_tsg_not_applicable OduTsGEt = "tsg-not-applicable"
)

// OduFlexTypeEt represents Odu flex type et
type OduFlexTypeEt string

const (
    // NA
    OduFlexTypeEt_na OduFlexTypeEt = "na"

    // ODU Flex Type CBR
    OduFlexTypeEt_cbr OduFlexTypeEt = "cbr"

    // ODU Flex Type GFP resizable
    OduFlexTypeEt_gfp_resizable OduFlexTypeEt = "gfp-resizable"

    // ODU Flex Type GFP fix
    OduFlexTypeEt_gfp_fix OduFlexTypeEt = "gfp-fix"
)

// OduPtTypeEt represents Odu pt type et
type OduPtTypeEt string

const (
    // NA
    OduPtTypeEt_na OduPtTypeEt = "na"

    // 02 (Asynchronous CBR mapping)
    OduPtTypeEt_two_asynchronous_cbr_mapping OduPtTypeEt = "two-asynchronous-cbr-mapping"

    // 03 (Bit synchronous CBR mapping)
    OduPtTypeEt_three_bit_synchronous_cbr_mapping OduPtTypeEt = "three-bit-synchronous-cbr-mapping"

    // 05 (GFP mapping)
    OduPtTypeEt_five_gfp_mapping OduPtTypeEt = "five-gfp-mapping"

    // 06 (Virtual Concatenated Signal)
    OduPtTypeEt_six_virtual_concatenated_signal OduPtTypeEt = "six-virtual-concatenated-signal"

    // 07 (PCS codeword transparent Ethernet mapping)
    OduPtTypeEt_seven_pcs_codeword_transparent_ethernet_mapping OduPtTypeEt = "seven-pcs-codeword-transparent-ethernet-mapping"

    // 09 (GFP mapping into OPU2)
    OduPtTypeEt_nine_gfp_mapping_into_opu2 OduPtTypeEt = "nine-gfp-mapping-into-opu2"

    // 0A (STM1 mapping into OPU0)
    OduPtTypeEt_zero_astm1_mapping_into_opu0 OduPtTypeEt = "zero-astm1-mapping-into-opu0"

    // 0B (STM4 mapping into OPU0)
    OduPtTypeEt_zero_bstm4_mapping_into_opu0 OduPtTypeEt = "zero-bstm4-mapping-into-opu0"

    // 20 (ODU multiplex structure supporting ODTUjk)
    OduPtTypeEt_twenty_odu_multiplex_structure_supporting_odt_ujk OduPtTypeEt = "twenty-odu-multiplex-structure-supporting-odt-ujk"

    // 21 (ODU multiplex structure supporting ODTUjk
    // and ODTUk.ts)
    OduPtTypeEt_twenty_one_odu_multiplex_structure_supporting_odt_ujk_and_odt_ukts OduPtTypeEt = "twenty-one-odu-multiplex-structure-supporting-odt-ujk-and-odt-ukts"
)

// OduResourceEt represents Odu resource et
type OduResourceEt string

const (
    // ODU Resource Free
    OduResourceEt_resource_free OduResourceEt = "resource-free"

    // ODU Open Connection
    OduResourceEt_open_connection OduResourceEt = "open-connection"

    // ODU Cross Connection
    OduResourceEt_cross_connection OduResourceEt = "cross-connection"

    // ODU Channelized
    OduResourceEt_channelized OduResourceEt = "channelized"

    // ODU Loopbacked
    OduResourceEt_loopbacked OduResourceEt = "loopbacked"

    // ODU Cross Connection and Loopbacked
    OduResourceEt_cross_connected_and_loopbacked OduResourceEt = "cross-connected-and-loopbacked"

    // ODU Terminated
    OduResourceEt_terminated OduResourceEt = "terminated"

    // ODU Invalid
    OduResourceEt_invalid OduResourceEt = "invalid"
)

// OduUserEt represents Odu user et
type OduUserEt string

const (
    // MP
    OduUserEt_mp OduUserEt = "mp"

    // GMPLS
    OduUserEt_gmpls OduUserEt = "gmpls"

    // All
    OduUserEt_all OduUserEt = "all"
)

// OduSecState represents Odu sec state
type OduSecState string

const (
    // Normal
    OduSecState_normal OduSecState = "normal"

    // Maintenance
    OduSecState_maintenance OduSecState = "maintenance"

    // Automatic In Service
    OduSecState_ains OduSecState = "ains"
)

// OduDerState represents Odu der state
type OduDerState string

const (
    // Out Of Service
    OduDerState_out_of_service OduDerState = "out-of-service"

    // In Service
    OduDerState_in_service OduDerState = "in-service"

    // Maintenance
    OduDerState_maintenance OduDerState = "maintenance"

    // Automatic In Service
    OduDerState_ains OduDerState = "ains"
)

// OduTtiEt represents Odu tti et
type OduTtiEt string

const (
    // ASCII
    OduTtiEt_ascii OduTtiEt = "ascii"

    // HEX
    OduTtiEt_hex OduTtiEt = "hex"

    // FULL ASCII
    OduTtiEt_full_ascii OduTtiEt = "full-ascii"

    // FULL HEX
    OduTtiEt_full_hex OduTtiEt = "full-hex"
)

// OduLoopBackMode represents Odu loop back mode
type OduLoopBackMode string

const (
    // None
    OduLoopBackMode_none OduLoopBackMode = "none"

    // Line
    OduLoopBackMode_line OduLoopBackMode = "line"

    // Internal
    OduLoopBackMode_internal OduLoopBackMode = "internal"
)

// OduStateEt represents Odu state et
type OduStateEt string

const (
    // Not Ready
    OduStateEt_not_ready OduStateEt = "not-ready"

    // Admin Down
    OduStateEt_admin_down OduStateEt = "admin-down"

    // Down
    OduStateEt_down OduStateEt = "down"

    // Up
    OduStateEt_up OduStateEt = "up"

    // Shutdown
    OduStateEt_shutdown OduStateEt = "shutdown"

    // Error Disable
    OduStateEt_error_disable OduStateEt = "error-disable"

    // Down Immediate
    OduStateEt_down_immediate OduStateEt = "down-immediate"

    // Down Immediate Admin
    OduStateEt_down_immediate_admin OduStateEt = "down-immediate-admin"

    // Down Graceful
    OduStateEt_down_graceful OduStateEt = "down-graceful"

    // Begin Shutdown
    OduStateEt_begin_shutdown OduStateEt = "begin-shutdown"

    // End Shutdown
    OduStateEt_end_shutdown OduStateEt = "end-shutdown"

    // Begin Error Disable
    OduStateEt_begin_error_disable OduStateEt = "begin-error-disable"

    // End Error Disable
    OduStateEt_end_error_disable OduStateEt = "end-error-disable"

    // Begin Down Graceful
    OduStateEt_begin_down_graceful OduStateEt = "begin-down-graceful"

    // Reset
    OduStateEt_reset OduStateEt = "reset"

    // Operational
    OduStateEt_operational OduStateEt = "operational"

    // Not Operational
    OduStateEt_not_operational OduStateEt = "not-operational"

    // Unknown
    OduStateEt_unknown OduStateEt = "unknown"

    // Last
    OduStateEt_last OduStateEt = "last"
)

// OduPrbsStatus represents Odu prbs status
type OduPrbsStatus string

const (
    // Locked
    OduPrbsStatus_locked OduPrbsStatus = "locked"

    // Unlocked
    OduPrbsStatus_unlocked OduPrbsStatus = "unlocked"

    // Not Applicable
    OduPrbsStatus_not_applicable OduPrbsStatus = "not-applicable"
)

// OduPrbsPattern represents Odu prbs pattern
type OduPrbsPattern string

const (
    // PNNONE
    OduPrbsPattern_pn_none OduPrbsPattern = "pn-none"

    // PN31
    OduPrbsPattern_pn31 OduPrbsPattern = "pn31"

    // PN23
    OduPrbsPattern_pn23 OduPrbsPattern = "pn23"

    // PN11
    OduPrbsPattern_pn11 OduPrbsPattern = "pn11"

    // INVERTED PN31
    OduPrbsPattern_inverted_pn31 OduPrbsPattern = "inverted-pn31"

    // INVERTED PN11
    OduPrbsPattern_inverted_pn11 OduPrbsPattern = "inverted-pn11"
)

// OduPrbsMode represents Odu prbs mode
type OduPrbsMode string

const (
    // invalid
    OduPrbsMode_invalid OduPrbsMode = "invalid"

    // Source
    OduPrbsMode_source OduPrbsMode = "source"

    // Sink
    OduPrbsMode_sink OduPrbsMode = "sink"

    // Source Sink
    OduPrbsMode_source_sink OduPrbsMode = "source-sink"
)

// OduPrbsTest represents Odu prbs test
type OduPrbsTest string

const (
    // Disable
    OduPrbsTest_disable OduPrbsTest = "disable"

    // Enable
    OduPrbsTest_enable OduPrbsTest = "enable"
)

// Odu
// ODU operational data
type Odu struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // All ODU Port operational data.
    Controllers Odu_Controllers
}

func (odu *Odu) GetEntityData() *types.CommonEntityData {
    odu.EntityData.YFilter = odu.YFilter
    odu.EntityData.YangName = "odu"
    odu.EntityData.BundleName = "cisco_ios_xr"
    odu.EntityData.ParentYangName = "Cisco-IOS-XR-controller-odu-oper"
    odu.EntityData.SegmentPath = "Cisco-IOS-XR-controller-odu-oper:odu"
    odu.EntityData.AbsolutePath = odu.EntityData.SegmentPath
    odu.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    odu.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    odu.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    odu.EntityData.Children = types.NewOrderedMap()
    odu.EntityData.Children.Append("controllers", types.YChild{"Controllers", &odu.Controllers})
    odu.EntityData.Leafs = types.NewOrderedMap()

    odu.EntityData.YListKeys = []string {}

    return &(odu.EntityData)
}

// Odu_Controllers
// All ODU Port operational data
type Odu_Controllers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ODU Port operational data. The type is slice of Odu_Controllers_Controller.
    Controller []*Odu_Controllers_Controller
}

func (controllers *Odu_Controllers) GetEntityData() *types.CommonEntityData {
    controllers.EntityData.YFilter = controllers.YFilter
    controllers.EntityData.YangName = "controllers"
    controllers.EntityData.BundleName = "cisco_ios_xr"
    controllers.EntityData.ParentYangName = "odu"
    controllers.EntityData.SegmentPath = "controllers"
    controllers.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/" + controllers.EntityData.SegmentPath
    controllers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    controllers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    controllers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    controllers.EntityData.Children = types.NewOrderedMap()
    controllers.EntityData.Children.Append("controller", types.YChild{"Controller", nil})
    for i := range controllers.Controller {
        controllers.EntityData.Children.Append(types.GetSegmentPath(controllers.Controller[i]), types.YChild{"Controller", controllers.Controller[i]})
    }
    controllers.EntityData.Leafs = types.NewOrderedMap()

    controllers.EntityData.YListKeys = []string {}

    return &(controllers.EntityData)
}

// Odu_Controllers_Controller
// ODU Port operational data
type Odu_Controllers_Controller struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Port name. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    ControllerName interface{}

    // ODU port operational data.
    Prbs Odu_Controllers_Controller_Prbs

    // ODU port operational data.
    Info Odu_Controllers_Controller_Info
}

func (controller *Odu_Controllers_Controller) GetEntityData() *types.CommonEntityData {
    controller.EntityData.YFilter = controller.YFilter
    controller.EntityData.YangName = "controller"
    controller.EntityData.BundleName = "cisco_ios_xr"
    controller.EntityData.ParentYangName = "controllers"
    controller.EntityData.SegmentPath = "controller" + types.AddKeyToken(controller.ControllerName, "controller-name")
    controller.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/" + controller.EntityData.SegmentPath
    controller.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    controller.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    controller.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    controller.EntityData.Children = types.NewOrderedMap()
    controller.EntityData.Children.Append("prbs", types.YChild{"Prbs", &controller.Prbs})
    controller.EntityData.Children.Append("info", types.YChild{"Info", &controller.Info})
    controller.EntityData.Leafs = types.NewOrderedMap()
    controller.EntityData.Leafs.Append("controller-name", types.YLeaf{"ControllerName", controller.ControllerName})

    controller.EntityData.YListKeys = []string {"ControllerName"}

    return &(controller.EntityData)
}

// Odu_Controllers_Controller_Prbs
// ODU port operational data
type Odu_Controllers_Controller_Prbs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // odu prbs test. The type is OduPrbsTest.
    OduPrbsTest interface{}

    // odu prbs mode. The type is OduPrbsMode.
    OduPrbsMode interface{}

    // odu prbs pattern. The type is OduPrbsPattern.
    OduPrbsPattern interface{}

    // odu prbs status. The type is OduPrbsStatus.
    OduPrbsStatus interface{}
}

func (prbs *Odu_Controllers_Controller_Prbs) GetEntityData() *types.CommonEntityData {
    prbs.EntityData.YFilter = prbs.YFilter
    prbs.EntityData.YangName = "prbs"
    prbs.EntityData.BundleName = "cisco_ios_xr"
    prbs.EntityData.ParentYangName = "controller"
    prbs.EntityData.SegmentPath = "prbs"
    prbs.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/" + prbs.EntityData.SegmentPath
    prbs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prbs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prbs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prbs.EntityData.Children = types.NewOrderedMap()
    prbs.EntityData.Leafs = types.NewOrderedMap()
    prbs.EntityData.Leafs.Append("odu-prbs-test", types.YLeaf{"OduPrbsTest", prbs.OduPrbsTest})
    prbs.EntityData.Leafs.Append("odu-prbs-mode", types.YLeaf{"OduPrbsMode", prbs.OduPrbsMode})
    prbs.EntityData.Leafs.Append("odu-prbs-pattern", types.YLeaf{"OduPrbsPattern", prbs.OduPrbsPattern})
    prbs.EntityData.Leafs.Append("odu-prbs-status", types.YLeaf{"OduPrbsStatus", prbs.OduPrbsStatus})

    prbs.EntityData.YListKeys = []string {}

    return &(prbs.EntityData)
}

// Odu_Controllers_Controller_Info
// ODU port operational data
type Odu_Controllers_Controller_Info struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Admin State. The type is OduStateEt.
    State interface{}

    // SF in the form of 1.0E - <SF>. The type is interface{} with range: 0..255.
    Sf interface{}

    // SD in the form of 1.0E - <SD>. The type is interface{} with range: 0..255.
    Sd interface{}

    // Loopback. The type is OduLoopBackMode.
    LoopbackMode interface{}

    // Derived State. The type is OduDerState.
    DerivedMode interface{}

    // Sec State . The type is OduSecState.
    InheritSecState interface{}

    // Sec State . The type is OduSecState.
    ConfigSecState interface{}

    // ODU GCC. The type is bool.
    GccMode interface{}

    // Child Name. The type is string with length: 0..64.
    ChildName interface{}

    // ODU maximum no of children. The type is interface{} with range: 0..255.
    MaxOduChild interface{}

    // ODU User. The type is OduUserEt.
    OdUuser interface{}

    // Resource State. The type is OduResourceEt.
    ResourceState interface{}

    // PT type. The type is OduPtTypeEt.
    PtType interface{}

    // FLEX type. The type is OduFlexTypeEt.
    FlexType interface{}

    // FLEX bw. The type is interface{} with range: 0..4294967295.
    FlexBw interface{}

    // FLEX tolerence. The type is interface{} with range: 0..65535.
    FlexTolerence interface{}

    // Option. The type is interface{} with range: 0..255.
    Option interface{}

    // TPN. The type is interface{} with range: 0..255.
    TpnValue interface{}

    // Number of TS. The type is interface{} with range: 0..255.
    NumTs interface{}

    // TS Granuality. The type is OduTsGEt.
    TsG interface{}

    // child ts bitmap. The type is string with length: 0..256.
    TsB interface{}

    // tpn bitmap. The type is string with length: 0..256.
    TpnB interface{}

    // ts bitmap. The type is string with length: 0..256.
    PtsB interface{}

    // fwd ref. The type is string with length: 0..64.
    FwdRef interface{}

    // Xconnect ID. The type is interface{} with range: 0..4294967295.
    XcId interface{}

    // Xconnect Name. The type is string.
    XconnectName interface{}

    // fwd_ref ifhandle. The type is interface{} with range: 0..4294967295.
    FwdRefIfhandle interface{}

    // Number of parent slot. The type is interface{} with range: 0..4294967295.
    NoParentSlot interface{}

    // Odu Xconnect Response code. The type is DpProgrammed.
    XcRespCode interface{}

    // Performance Monitoring. The type is OduPerMon.
    PerformanceMonitoring interface{}

    // PM TIM-CA state. The type is OduPmCaEt.
    Pmtimca interface{}

    // ODU PM Mode. The type is OduPmMode.
    PmMode interface{}

    // NV Optical support. The type is bool.
    NvOpticalSupport interface{}

    // tti mode. The type is GmplsTtiMode.
    GmplsTtiMode interface{}

    // tcm id. The type is interface{} with range: 0..255.
    GmplsTcmId interface{}

    // TTI.
    Local Odu_Controllers_Controller_Info_Local

    // Remote.
    Remote Odu_Controllers_Controller_Info_Remote

    // TTI.
    TtiMode Odu_Controllers_Controller_Info_TtiMode

    // ODU fwd_ref .
    OduFwdRef Odu_Controllers_Controller_Info_OduFwdRef

    // Alarm.
    Alarm Odu_Controllers_Controller_Info_Alarm

    // Label Get Data.
    TeCtxData Odu_Controllers_Controller_Info_TeCtxData

    // Xconnect Add Data.
    XcAddCtxData Odu_Controllers_Controller_Info_XcAddCtxData

    // Xconnect Remove Data.
    XcRemCtxData Odu_Controllers_Controller_Info_XcRemCtxData

    // ODU Delay.
    OduDelay Odu_Controllers_Controller_Info_OduDelay

    // odu terminate ether.
    OduTerminateEther Odu_Controllers_Controller_Info_OduTerminateEther

    // AINS information.
    AinsInfo Odu_Controllers_Controller_Info_AinsInfo

    // Child Ts. The type is slice of Odu_Controllers_Controller_Info_Odu.
    Odu []*Odu_Controllers_Controller_Info_Odu

    // ODU TCM. The type is slice of Odu_Controllers_Controller_Info_Odutcm.
    Odutcm []*Odu_Controllers_Controller_Info_Odutcm
}

func (info *Odu_Controllers_Controller_Info) GetEntityData() *types.CommonEntityData {
    info.EntityData.YFilter = info.YFilter
    info.EntityData.YangName = "info"
    info.EntityData.BundleName = "cisco_ios_xr"
    info.EntityData.ParentYangName = "controller"
    info.EntityData.SegmentPath = "info"
    info.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/" + info.EntityData.SegmentPath
    info.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    info.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    info.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    info.EntityData.Children = types.NewOrderedMap()
    info.EntityData.Children.Append("local", types.YChild{"Local", &info.Local})
    info.EntityData.Children.Append("remote", types.YChild{"Remote", &info.Remote})
    info.EntityData.Children.Append("tti-mode", types.YChild{"TtiMode", &info.TtiMode})
    info.EntityData.Children.Append("odu-fwd-ref", types.YChild{"OduFwdRef", &info.OduFwdRef})
    info.EntityData.Children.Append("alarm", types.YChild{"Alarm", &info.Alarm})
    info.EntityData.Children.Append("te-ctx-data", types.YChild{"TeCtxData", &info.TeCtxData})
    info.EntityData.Children.Append("xc-add-ctx-data", types.YChild{"XcAddCtxData", &info.XcAddCtxData})
    info.EntityData.Children.Append("xc-rem-ctx-data", types.YChild{"XcRemCtxData", &info.XcRemCtxData})
    info.EntityData.Children.Append("odu-delay", types.YChild{"OduDelay", &info.OduDelay})
    info.EntityData.Children.Append("odu-terminate-ether", types.YChild{"OduTerminateEther", &info.OduTerminateEther})
    info.EntityData.Children.Append("ains-info", types.YChild{"AinsInfo", &info.AinsInfo})
    info.EntityData.Children.Append("odu", types.YChild{"Odu", nil})
    for i := range info.Odu {
        types.SetYListKey(info.Odu[i], i)
        info.EntityData.Children.Append(types.GetSegmentPath(info.Odu[i]), types.YChild{"Odu", info.Odu[i]})
    }
    info.EntityData.Children.Append("odutcm", types.YChild{"Odutcm", nil})
    for i := range info.Odutcm {
        types.SetYListKey(info.Odutcm[i], i)
        info.EntityData.Children.Append(types.GetSegmentPath(info.Odutcm[i]), types.YChild{"Odutcm", info.Odutcm[i]})
    }
    info.EntityData.Leafs = types.NewOrderedMap()
    info.EntityData.Leafs.Append("state", types.YLeaf{"State", info.State})
    info.EntityData.Leafs.Append("sf", types.YLeaf{"Sf", info.Sf})
    info.EntityData.Leafs.Append("sd", types.YLeaf{"Sd", info.Sd})
    info.EntityData.Leafs.Append("loopback-mode", types.YLeaf{"LoopbackMode", info.LoopbackMode})
    info.EntityData.Leafs.Append("derived-mode", types.YLeaf{"DerivedMode", info.DerivedMode})
    info.EntityData.Leafs.Append("inherit-sec-state", types.YLeaf{"InheritSecState", info.InheritSecState})
    info.EntityData.Leafs.Append("config-sec-state", types.YLeaf{"ConfigSecState", info.ConfigSecState})
    info.EntityData.Leafs.Append("gcc-mode", types.YLeaf{"GccMode", info.GccMode})
    info.EntityData.Leafs.Append("child-name", types.YLeaf{"ChildName", info.ChildName})
    info.EntityData.Leafs.Append("max-odu-child", types.YLeaf{"MaxOduChild", info.MaxOduChild})
    info.EntityData.Leafs.Append("od-uuser", types.YLeaf{"OdUuser", info.OdUuser})
    info.EntityData.Leafs.Append("resource-state", types.YLeaf{"ResourceState", info.ResourceState})
    info.EntityData.Leafs.Append("pt-type", types.YLeaf{"PtType", info.PtType})
    info.EntityData.Leafs.Append("flex-type", types.YLeaf{"FlexType", info.FlexType})
    info.EntityData.Leafs.Append("flex-bw", types.YLeaf{"FlexBw", info.FlexBw})
    info.EntityData.Leafs.Append("flex-tolerence", types.YLeaf{"FlexTolerence", info.FlexTolerence})
    info.EntityData.Leafs.Append("option", types.YLeaf{"Option", info.Option})
    info.EntityData.Leafs.Append("tpn-value", types.YLeaf{"TpnValue", info.TpnValue})
    info.EntityData.Leafs.Append("num-ts", types.YLeaf{"NumTs", info.NumTs})
    info.EntityData.Leafs.Append("ts-g", types.YLeaf{"TsG", info.TsG})
    info.EntityData.Leafs.Append("ts-b", types.YLeaf{"TsB", info.TsB})
    info.EntityData.Leafs.Append("tpn-b", types.YLeaf{"TpnB", info.TpnB})
    info.EntityData.Leafs.Append("pts-b", types.YLeaf{"PtsB", info.PtsB})
    info.EntityData.Leafs.Append("fwd-ref", types.YLeaf{"FwdRef", info.FwdRef})
    info.EntityData.Leafs.Append("xc-id", types.YLeaf{"XcId", info.XcId})
    info.EntityData.Leafs.Append("xconnect-name", types.YLeaf{"XconnectName", info.XconnectName})
    info.EntityData.Leafs.Append("fwd-ref-ifhandle", types.YLeaf{"FwdRefIfhandle", info.FwdRefIfhandle})
    info.EntityData.Leafs.Append("no-parent-slot", types.YLeaf{"NoParentSlot", info.NoParentSlot})
    info.EntityData.Leafs.Append("xc-resp-code", types.YLeaf{"XcRespCode", info.XcRespCode})
    info.EntityData.Leafs.Append("performance-monitoring", types.YLeaf{"PerformanceMonitoring", info.PerformanceMonitoring})
    info.EntityData.Leafs.Append("pmtimca", types.YLeaf{"Pmtimca", info.Pmtimca})
    info.EntityData.Leafs.Append("pm-mode", types.YLeaf{"PmMode", info.PmMode})
    info.EntityData.Leafs.Append("nv-optical-support", types.YLeaf{"NvOpticalSupport", info.NvOpticalSupport})
    info.EntityData.Leafs.Append("gmpls-tti-mode", types.YLeaf{"GmplsTtiMode", info.GmplsTtiMode})
    info.EntityData.Leafs.Append("gmpls-tcm-id", types.YLeaf{"GmplsTcmId", info.GmplsTcmId})

    info.EntityData.YListKeys = []string {}

    return &(info.EntityData)
}

// Odu_Controllers_Controller_Info_Local
// TTI
type Odu_Controllers_Controller_Info_Local struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Router ID. The type is interface{} with range: 0..4294967295.
    RouterId interface{}

    // IfIndex. The type is interface{} with range: 0..4294967295.
    IfIndex interface{}
}

func (local *Odu_Controllers_Controller_Info_Local) GetEntityData() *types.CommonEntityData {
    local.EntityData.YFilter = local.YFilter
    local.EntityData.YangName = "local"
    local.EntityData.BundleName = "cisco_ios_xr"
    local.EntityData.ParentYangName = "info"
    local.EntityData.SegmentPath = "local"
    local.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/" + local.EntityData.SegmentPath
    local.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    local.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    local.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    local.EntityData.Children = types.NewOrderedMap()
    local.EntityData.Leafs = types.NewOrderedMap()
    local.EntityData.Leafs.Append("router-id", types.YLeaf{"RouterId", local.RouterId})
    local.EntityData.Leafs.Append("if-index", types.YLeaf{"IfIndex", local.IfIndex})

    local.EntityData.YListKeys = []string {}

    return &(local.EntityData)
}

// Odu_Controllers_Controller_Info_Remote
// Remote
type Odu_Controllers_Controller_Info_Remote struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Router ID. The type is interface{} with range: 0..4294967295.
    RouterId interface{}

    // IfIndex. The type is interface{} with range: 0..4294967295.
    IfIndex interface{}
}

func (remote *Odu_Controllers_Controller_Info_Remote) GetEntityData() *types.CommonEntityData {
    remote.EntityData.YFilter = remote.YFilter
    remote.EntityData.YangName = "remote"
    remote.EntityData.BundleName = "cisco_ios_xr"
    remote.EntityData.ParentYangName = "info"
    remote.EntityData.SegmentPath = "remote"
    remote.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/" + remote.EntityData.SegmentPath
    remote.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remote.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remote.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remote.EntityData.Children = types.NewOrderedMap()
    remote.EntityData.Leafs = types.NewOrderedMap()
    remote.EntityData.Leafs.Append("router-id", types.YLeaf{"RouterId", remote.RouterId})
    remote.EntityData.Leafs.Append("if-index", types.YLeaf{"IfIndex", remote.IfIndex})

    remote.EntityData.YListKeys = []string {}

    return &(remote.EntityData)
}

// Odu_Controllers_Controller_Info_TtiMode
// TTI
type Odu_Controllers_Controller_Info_TtiMode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // G709TTI Sent. The type is OduTtiEt.
    G709ttiSentMode interface{}

    // G709TTI Expected. The type is OduTtiEt.
    G709ttiExpMode interface{}

    // G709TTI Recieved. The type is OduTtiEt.
    G709ttiRecMode interface{}

    // String Sent.
    Tx Odu_Controllers_Controller_Info_TtiMode_Tx

    // String Expected.
    Exp Odu_Controllers_Controller_Info_TtiMode_Exp

    // String Received.
    Rec Odu_Controllers_Controller_Info_TtiMode_Rec
}

func (ttiMode *Odu_Controllers_Controller_Info_TtiMode) GetEntityData() *types.CommonEntityData {
    ttiMode.EntityData.YFilter = ttiMode.YFilter
    ttiMode.EntityData.YangName = "tti-mode"
    ttiMode.EntityData.BundleName = "cisco_ios_xr"
    ttiMode.EntityData.ParentYangName = "info"
    ttiMode.EntityData.SegmentPath = "tti-mode"
    ttiMode.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/" + ttiMode.EntityData.SegmentPath
    ttiMode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ttiMode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ttiMode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ttiMode.EntityData.Children = types.NewOrderedMap()
    ttiMode.EntityData.Children.Append("tx", types.YChild{"Tx", &ttiMode.Tx})
    ttiMode.EntityData.Children.Append("exp", types.YChild{"Exp", &ttiMode.Exp})
    ttiMode.EntityData.Children.Append("rec", types.YChild{"Rec", &ttiMode.Rec})
    ttiMode.EntityData.Leafs = types.NewOrderedMap()
    ttiMode.EntityData.Leafs.Append("g709tti-sent-mode", types.YLeaf{"G709ttiSentMode", ttiMode.G709ttiSentMode})
    ttiMode.EntityData.Leafs.Append("g709tti-exp-mode", types.YLeaf{"G709ttiExpMode", ttiMode.G709ttiExpMode})
    ttiMode.EntityData.Leafs.Append("g709tti-rec-mode", types.YLeaf{"G709ttiRecMode", ttiMode.G709ttiRecMode})

    ttiMode.EntityData.YListKeys = []string {}

    return &(ttiMode.EntityData)
}

// Odu_Controllers_Controller_Info_TtiMode_Tx
// String Sent
type Odu_Controllers_Controller_Info_TtiMode_Tx struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // tx String. The type is slice of interface{} with range: 0..255.
    Sapi []interface{}

    // exp String. The type is slice of interface{} with range: 0..255.
    Dapi []interface{}

    // rec String. The type is slice of interface{} with range: 0..255.
    OperatorSpecific []interface{}
}

func (tx *Odu_Controllers_Controller_Info_TtiMode_Tx) GetEntityData() *types.CommonEntityData {
    tx.EntityData.YFilter = tx.YFilter
    tx.EntityData.YangName = "tx"
    tx.EntityData.BundleName = "cisco_ios_xr"
    tx.EntityData.ParentYangName = "tti-mode"
    tx.EntityData.SegmentPath = "tx"
    tx.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/tti-mode/" + tx.EntityData.SegmentPath
    tx.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tx.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tx.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tx.EntityData.Children = types.NewOrderedMap()
    tx.EntityData.Leafs = types.NewOrderedMap()
    tx.EntityData.Leafs.Append("sapi", types.YLeaf{"Sapi", tx.Sapi})
    tx.EntityData.Leafs.Append("dapi", types.YLeaf{"Dapi", tx.Dapi})
    tx.EntityData.Leafs.Append("operator-specific", types.YLeaf{"OperatorSpecific", tx.OperatorSpecific})

    tx.EntityData.YListKeys = []string {}

    return &(tx.EntityData)
}

// Odu_Controllers_Controller_Info_TtiMode_Exp
// String Expected
type Odu_Controllers_Controller_Info_TtiMode_Exp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // tx String. The type is slice of interface{} with range: 0..255.
    Sapi []interface{}

    // exp String. The type is slice of interface{} with range: 0..255.
    Dapi []interface{}

    // rec String. The type is slice of interface{} with range: 0..255.
    OperatorSpecific []interface{}
}

func (exp *Odu_Controllers_Controller_Info_TtiMode_Exp) GetEntityData() *types.CommonEntityData {
    exp.EntityData.YFilter = exp.YFilter
    exp.EntityData.YangName = "exp"
    exp.EntityData.BundleName = "cisco_ios_xr"
    exp.EntityData.ParentYangName = "tti-mode"
    exp.EntityData.SegmentPath = "exp"
    exp.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/tti-mode/" + exp.EntityData.SegmentPath
    exp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    exp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    exp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    exp.EntityData.Children = types.NewOrderedMap()
    exp.EntityData.Leafs = types.NewOrderedMap()
    exp.EntityData.Leafs.Append("sapi", types.YLeaf{"Sapi", exp.Sapi})
    exp.EntityData.Leafs.Append("dapi", types.YLeaf{"Dapi", exp.Dapi})
    exp.EntityData.Leafs.Append("operator-specific", types.YLeaf{"OperatorSpecific", exp.OperatorSpecific})

    exp.EntityData.YListKeys = []string {}

    return &(exp.EntityData)
}

// Odu_Controllers_Controller_Info_TtiMode_Rec
// String Received
type Odu_Controllers_Controller_Info_TtiMode_Rec struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // tx String. The type is slice of interface{} with range: 0..255.
    Sapi []interface{}

    // exp String. The type is slice of interface{} with range: 0..255.
    Dapi []interface{}

    // rec String. The type is slice of interface{} with range: 0..255.
    OperatorSpecific []interface{}
}

func (rec *Odu_Controllers_Controller_Info_TtiMode_Rec) GetEntityData() *types.CommonEntityData {
    rec.EntityData.YFilter = rec.YFilter
    rec.EntityData.YangName = "rec"
    rec.EntityData.BundleName = "cisco_ios_xr"
    rec.EntityData.ParentYangName = "tti-mode"
    rec.EntityData.SegmentPath = "rec"
    rec.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/tti-mode/" + rec.EntityData.SegmentPath
    rec.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rec.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rec.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rec.EntityData.Children = types.NewOrderedMap()
    rec.EntityData.Leafs = types.NewOrderedMap()
    rec.EntityData.Leafs.Append("sapi", types.YLeaf{"Sapi", rec.Sapi})
    rec.EntityData.Leafs.Append("dapi", types.YLeaf{"Dapi", rec.Dapi})
    rec.EntityData.Leafs.Append("operator-specific", types.YLeaf{"OperatorSpecific", rec.OperatorSpecific})

    rec.EntityData.YListKeys = []string {}

    return &(rec.EntityData)
}

// Odu_Controllers_Controller_Info_OduFwdRef
// ODU fwd_ref 
type Odu_Controllers_Controller_Info_OduFwdRef struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ODU User. The type is OduUserEt.
    OdUuser interface{}

    // Resource State. The type is OduResourceEt.
    ResourceState interface{}
}

func (oduFwdRef *Odu_Controllers_Controller_Info_OduFwdRef) GetEntityData() *types.CommonEntityData {
    oduFwdRef.EntityData.YFilter = oduFwdRef.YFilter
    oduFwdRef.EntityData.YangName = "odu-fwd-ref"
    oduFwdRef.EntityData.BundleName = "cisco_ios_xr"
    oduFwdRef.EntityData.ParentYangName = "info"
    oduFwdRef.EntityData.SegmentPath = "odu-fwd-ref"
    oduFwdRef.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/" + oduFwdRef.EntityData.SegmentPath
    oduFwdRef.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    oduFwdRef.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    oduFwdRef.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    oduFwdRef.EntityData.Children = types.NewOrderedMap()
    oduFwdRef.EntityData.Leafs = types.NewOrderedMap()
    oduFwdRef.EntityData.Leafs.Append("od-uuser", types.YLeaf{"OdUuser", oduFwdRef.OdUuser})
    oduFwdRef.EntityData.Leafs.Append("resource-state", types.YLeaf{"ResourceState", oduFwdRef.ResourceState})

    oduFwdRef.EntityData.YListKeys = []string {}

    return &(oduFwdRef.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm
// Alarm
type Odu_Controllers_Controller_Info_Alarm struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Open Connection Indiction.
    Oci Odu_Controllers_Controller_Info_Alarm_Oci

    // Alarm Indication Signal.
    Ais Odu_Controllers_Controller_Info_Alarm_Ais

    // Upstream Connection Locked.
    Lck Odu_Controllers_Controller_Info_Alarm_Lck

    // Backward Defect Indication.
    Bdi Odu_Controllers_Controller_Info_Alarm_Bdi

    // GCC End of Channel.
    Eoc Odu_Controllers_Controller_Info_Alarm_Eoc

    // Payload Type Identifier Mismatch.
    Ptim Odu_Controllers_Controller_Info_Alarm_Ptim

    // Trace Identifier Mismatch information.
    Tim Odu_Controllers_Controller_Info_Alarm_Tim

    // Incoming Alignment Error.
    Iae Odu_Controllers_Controller_Info_Alarm_Iae

    // Backward Incoming Alignment Error.
    Biae Odu_Controllers_Controller_Info_Alarm_Biae

    // SF BER alarm.
    SfBer Odu_Controllers_Controller_Info_Alarm_SfBer

    // SD BER alarm.
    SdBer Odu_Controllers_Controller_Info_Alarm_SdBer

    // Client Signal Failure.
    Csf Odu_Controllers_Controller_Info_Alarm_Csf

    // TCM1 Alarm Indication Signal.
    Tcm1Ais Odu_Controllers_Controller_Info_Alarm_Tcm1Ais

    // TCM1 Loss of Tandem connection Monitoring.
    Tcm1Ltc Odu_Controllers_Controller_Info_Alarm_Tcm1Ltc

    // TCM1 Open Connection Indiction.
    Tcm1Oci Odu_Controllers_Controller_Info_Alarm_Tcm1Oci

    // TCM1  Upstream Connection Locked.
    Tcm1Lck Odu_Controllers_Controller_Info_Alarm_Tcm1Lck

    // TCM1 Incoming Alignment Error.
    Tcm1Iae Odu_Controllers_Controller_Info_Alarm_Tcm1Iae

    // TCM1 Backward Incoming Alignment Error.
    Tcm1Biae Odu_Controllers_Controller_Info_Alarm_Tcm1Biae

    // TCM1 Backward Defect Monitoring.
    Tcm1Bdi Odu_Controllers_Controller_Info_Alarm_Tcm1Bdi

    // TCM1 Trail Trace Identifier Mismatch .
    Tcm1Tim Odu_Controllers_Controller_Info_Alarm_Tcm1Tim

    // TCM1 SF BER alarm.
    Tcm1SfBer Odu_Controllers_Controller_Info_Alarm_Tcm1SfBer

    // TCM1 SD BER alarm.
    Tcm1SdBer Odu_Controllers_Controller_Info_Alarm_Tcm1SdBer

    // TCM2 Alarm Indication Signal.
    Tcm2Ais Odu_Controllers_Controller_Info_Alarm_Tcm2Ais

    // TCM2 Loss of Tandem connection Monitoring.
    Tcm2Ltc Odu_Controllers_Controller_Info_Alarm_Tcm2Ltc

    // TCM2 Open Connection Indiction.
    Tcm2Oci Odu_Controllers_Controller_Info_Alarm_Tcm2Oci

    // TCM2  Upstream Connection Locked.
    Tcm2Lck Odu_Controllers_Controller_Info_Alarm_Tcm2Lck

    // TCM2 Incoming Alignment Error.
    Tcm2Iae Odu_Controllers_Controller_Info_Alarm_Tcm2Iae

    // TCM2 Backward Incoming Alignment Error.
    Tcm2Biae Odu_Controllers_Controller_Info_Alarm_Tcm2Biae

    // TCM2 Backward Defect Monitoring.
    Tcm2Bdi Odu_Controllers_Controller_Info_Alarm_Tcm2Bdi

    // TCM2 Trail Trace Identifier Mismatch .
    Tcm2Tim Odu_Controllers_Controller_Info_Alarm_Tcm2Tim

    // TCM2 SF BER alarm.
    Tcm2SfBer Odu_Controllers_Controller_Info_Alarm_Tcm2SfBer

    // TCM2 SD BER alarm.
    Tcm2SdBer Odu_Controllers_Controller_Info_Alarm_Tcm2SdBer

    // TCM3 Alarm Indication Signal.
    Tcm3Ais Odu_Controllers_Controller_Info_Alarm_Tcm3Ais

    // TCM3 Loss of Tandem connection Monitoring.
    Tcm3Ltc Odu_Controllers_Controller_Info_Alarm_Tcm3Ltc

    // TCM3 Open Connection Indiction.
    Tcm3Oci Odu_Controllers_Controller_Info_Alarm_Tcm3Oci

    // TCM3  Upstream Connection Locked.
    Tcm3Lck Odu_Controllers_Controller_Info_Alarm_Tcm3Lck

    // TCM3 Incoming Alignment Error.
    Tcm3Iae Odu_Controllers_Controller_Info_Alarm_Tcm3Iae

    // TCM3 Backward Incoming Alignment Error.
    Tcm3Biae Odu_Controllers_Controller_Info_Alarm_Tcm3Biae

    // TCM3 Backward Defect Monitoring.
    Tcm3Bdi Odu_Controllers_Controller_Info_Alarm_Tcm3Bdi

    // TCM3 Trail Trace Identifier Mismatch .
    Tcm3Tim Odu_Controllers_Controller_Info_Alarm_Tcm3Tim

    // TCM3 SF BER alarm.
    Tcm3SfBer Odu_Controllers_Controller_Info_Alarm_Tcm3SfBer

    // TCM3 SD BER alarm.
    Tcm3SdBer Odu_Controllers_Controller_Info_Alarm_Tcm3SdBer

    // TCM4 Alarm Indication Signal.
    Tcm4Ais Odu_Controllers_Controller_Info_Alarm_Tcm4Ais

    // TCM4 Loss of Tandem connection Monitoring.
    Tcm4Ltc Odu_Controllers_Controller_Info_Alarm_Tcm4Ltc

    // TCM4 Open Connection Indiction.
    Tcm4Oci Odu_Controllers_Controller_Info_Alarm_Tcm4Oci

    // TCM4  Upstream Connection Locked.
    Tcm4Lck Odu_Controllers_Controller_Info_Alarm_Tcm4Lck

    // TCM4 Incoming Alignment Error.
    Tcm4Iae Odu_Controllers_Controller_Info_Alarm_Tcm4Iae

    // TCM4 Backward Incoming Alignment Error.
    Tcm4Biae Odu_Controllers_Controller_Info_Alarm_Tcm4Biae

    // TCM4 Backward Defect Monitoring.
    Tcm4Bdi Odu_Controllers_Controller_Info_Alarm_Tcm4Bdi

    // TCM4 Trail Trace Identifier Mismatch .
    Tcm4Tim Odu_Controllers_Controller_Info_Alarm_Tcm4Tim

    // TCM4 SF BER alarm.
    Tcm4SfBer Odu_Controllers_Controller_Info_Alarm_Tcm4SfBer

    // TCM4 SD BER alarm.
    Tcm4SdBer Odu_Controllers_Controller_Info_Alarm_Tcm4SdBer

    // TCM5 Alarm Indication Signal.
    Tcm5Ais Odu_Controllers_Controller_Info_Alarm_Tcm5Ais

    // TCM5 Loss of Tandem connection Monitoring.
    Tcm5Ltc Odu_Controllers_Controller_Info_Alarm_Tcm5Ltc

    // TCM5 Open Connection Indiction.
    Tcm5Oci Odu_Controllers_Controller_Info_Alarm_Tcm5Oci

    // TCM5  Upstream Connection Locked.
    Tcm5Lck Odu_Controllers_Controller_Info_Alarm_Tcm5Lck

    // TCM5 Incoming Alignment Error.
    Tcm5Iae Odu_Controllers_Controller_Info_Alarm_Tcm5Iae

    // TCM5 Backward Incoming Alignment Error.
    Tcm5Biae Odu_Controllers_Controller_Info_Alarm_Tcm5Biae

    // TCM5 Backward Defect Monitoring.
    Tcm5Bdi Odu_Controllers_Controller_Info_Alarm_Tcm5Bdi

    // TCM5 Trail Trace Identifier Mismatch .
    Tcm5Tim Odu_Controllers_Controller_Info_Alarm_Tcm5Tim

    // TCM5 SF BER alarm.
    Tcm5SfBer Odu_Controllers_Controller_Info_Alarm_Tcm5SfBer

    // TCM5 SD BER alarm.
    Tcm5SdBer Odu_Controllers_Controller_Info_Alarm_Tcm5SdBer

    // TCM6 Alarm Indication Signal.
    Tcm6Ais Odu_Controllers_Controller_Info_Alarm_Tcm6Ais

    // TCM6 Loss of Tandem connection Monitoring.
    Tcm6Ltc Odu_Controllers_Controller_Info_Alarm_Tcm6Ltc

    // TCM6 Open Connection Indiction.
    Tcm6Oci Odu_Controllers_Controller_Info_Alarm_Tcm6Oci

    // TCM6  Upstream Connection Locked.
    Tcm6Lck Odu_Controllers_Controller_Info_Alarm_Tcm6Lck

    // TCM6 Incoming Alignment Error.
    Tcm6Iae Odu_Controllers_Controller_Info_Alarm_Tcm6Iae

    // TCM6 Backward Incoming Alignment Error.
    Tcm6Biae Odu_Controllers_Controller_Info_Alarm_Tcm6Biae

    // TCM6 Backward Defect Monitoring.
    Tcm6Bdi Odu_Controllers_Controller_Info_Alarm_Tcm6Bdi

    // TCM6 Trail Trace Identifier Mismatch .
    Tcm6Tim Odu_Controllers_Controller_Info_Alarm_Tcm6Tim

    // TCM6 SF BER alarm.
    Tcm6SfBer Odu_Controllers_Controller_Info_Alarm_Tcm6SfBer

    // TCM6 SD BER alarm.
    Tcm6SdBer Odu_Controllers_Controller_Info_Alarm_Tcm6SdBer

    // Loss Of Frame Delineation.
    GfpLfd Odu_Controllers_Controller_Info_Alarm_GfpLfd

    // Loss Of Client Signal.
    GfpLocs Odu_Controllers_Controller_Info_Alarm_GfpLocs

    // Loss Of Character Synchronization.
    GfpLoccs Odu_Controllers_Controller_Info_Alarm_GfpLoccs

    // User Payload Mismatch.
    GfpUpm Odu_Controllers_Controller_Info_Alarm_GfpUpm
}

func (alarm *Odu_Controllers_Controller_Info_Alarm) GetEntityData() *types.CommonEntityData {
    alarm.EntityData.YFilter = alarm.YFilter
    alarm.EntityData.YangName = "alarm"
    alarm.EntityData.BundleName = "cisco_ios_xr"
    alarm.EntityData.ParentYangName = "info"
    alarm.EntityData.SegmentPath = "alarm"
    alarm.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/" + alarm.EntityData.SegmentPath
    alarm.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    alarm.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    alarm.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    alarm.EntityData.Children = types.NewOrderedMap()
    alarm.EntityData.Children.Append("oci", types.YChild{"Oci", &alarm.Oci})
    alarm.EntityData.Children.Append("ais", types.YChild{"Ais", &alarm.Ais})
    alarm.EntityData.Children.Append("lck", types.YChild{"Lck", &alarm.Lck})
    alarm.EntityData.Children.Append("bdi", types.YChild{"Bdi", &alarm.Bdi})
    alarm.EntityData.Children.Append("eoc", types.YChild{"Eoc", &alarm.Eoc})
    alarm.EntityData.Children.Append("ptim", types.YChild{"Ptim", &alarm.Ptim})
    alarm.EntityData.Children.Append("tim", types.YChild{"Tim", &alarm.Tim})
    alarm.EntityData.Children.Append("iae", types.YChild{"Iae", &alarm.Iae})
    alarm.EntityData.Children.Append("biae", types.YChild{"Biae", &alarm.Biae})
    alarm.EntityData.Children.Append("sf-ber", types.YChild{"SfBer", &alarm.SfBer})
    alarm.EntityData.Children.Append("sd-ber", types.YChild{"SdBer", &alarm.SdBer})
    alarm.EntityData.Children.Append("csf", types.YChild{"Csf", &alarm.Csf})
    alarm.EntityData.Children.Append("tcm1-ais", types.YChild{"Tcm1Ais", &alarm.Tcm1Ais})
    alarm.EntityData.Children.Append("tcm1-ltc", types.YChild{"Tcm1Ltc", &alarm.Tcm1Ltc})
    alarm.EntityData.Children.Append("tcm1-oci", types.YChild{"Tcm1Oci", &alarm.Tcm1Oci})
    alarm.EntityData.Children.Append("tcm1-lck", types.YChild{"Tcm1Lck", &alarm.Tcm1Lck})
    alarm.EntityData.Children.Append("tcm1-iae", types.YChild{"Tcm1Iae", &alarm.Tcm1Iae})
    alarm.EntityData.Children.Append("tcm1-biae", types.YChild{"Tcm1Biae", &alarm.Tcm1Biae})
    alarm.EntityData.Children.Append("tcm1-bdi", types.YChild{"Tcm1Bdi", &alarm.Tcm1Bdi})
    alarm.EntityData.Children.Append("tcm1-tim", types.YChild{"Tcm1Tim", &alarm.Tcm1Tim})
    alarm.EntityData.Children.Append("tcm1-sf-ber", types.YChild{"Tcm1SfBer", &alarm.Tcm1SfBer})
    alarm.EntityData.Children.Append("tcm1-sd-ber", types.YChild{"Tcm1SdBer", &alarm.Tcm1SdBer})
    alarm.EntityData.Children.Append("tcm2-ais", types.YChild{"Tcm2Ais", &alarm.Tcm2Ais})
    alarm.EntityData.Children.Append("tcm2-ltc", types.YChild{"Tcm2Ltc", &alarm.Tcm2Ltc})
    alarm.EntityData.Children.Append("tcm2-oci", types.YChild{"Tcm2Oci", &alarm.Tcm2Oci})
    alarm.EntityData.Children.Append("tcm2-lck", types.YChild{"Tcm2Lck", &alarm.Tcm2Lck})
    alarm.EntityData.Children.Append("tcm2-iae", types.YChild{"Tcm2Iae", &alarm.Tcm2Iae})
    alarm.EntityData.Children.Append("tcm2-biae", types.YChild{"Tcm2Biae", &alarm.Tcm2Biae})
    alarm.EntityData.Children.Append("tcm2-bdi", types.YChild{"Tcm2Bdi", &alarm.Tcm2Bdi})
    alarm.EntityData.Children.Append("tcm2-tim", types.YChild{"Tcm2Tim", &alarm.Tcm2Tim})
    alarm.EntityData.Children.Append("tcm2-sf-ber", types.YChild{"Tcm2SfBer", &alarm.Tcm2SfBer})
    alarm.EntityData.Children.Append("tcm2-sd-ber", types.YChild{"Tcm2SdBer", &alarm.Tcm2SdBer})
    alarm.EntityData.Children.Append("tcm3-ais", types.YChild{"Tcm3Ais", &alarm.Tcm3Ais})
    alarm.EntityData.Children.Append("tcm3-ltc", types.YChild{"Tcm3Ltc", &alarm.Tcm3Ltc})
    alarm.EntityData.Children.Append("tcm3-oci", types.YChild{"Tcm3Oci", &alarm.Tcm3Oci})
    alarm.EntityData.Children.Append("tcm3-lck", types.YChild{"Tcm3Lck", &alarm.Tcm3Lck})
    alarm.EntityData.Children.Append("tcm3-iae", types.YChild{"Tcm3Iae", &alarm.Tcm3Iae})
    alarm.EntityData.Children.Append("tcm3-biae", types.YChild{"Tcm3Biae", &alarm.Tcm3Biae})
    alarm.EntityData.Children.Append("tcm3-bdi", types.YChild{"Tcm3Bdi", &alarm.Tcm3Bdi})
    alarm.EntityData.Children.Append("tcm3-tim", types.YChild{"Tcm3Tim", &alarm.Tcm3Tim})
    alarm.EntityData.Children.Append("tcm3-sf-ber", types.YChild{"Tcm3SfBer", &alarm.Tcm3SfBer})
    alarm.EntityData.Children.Append("tcm3-sd-ber", types.YChild{"Tcm3SdBer", &alarm.Tcm3SdBer})
    alarm.EntityData.Children.Append("tcm4-ais", types.YChild{"Tcm4Ais", &alarm.Tcm4Ais})
    alarm.EntityData.Children.Append("tcm4-ltc", types.YChild{"Tcm4Ltc", &alarm.Tcm4Ltc})
    alarm.EntityData.Children.Append("tcm4-oci", types.YChild{"Tcm4Oci", &alarm.Tcm4Oci})
    alarm.EntityData.Children.Append("tcm4-lck", types.YChild{"Tcm4Lck", &alarm.Tcm4Lck})
    alarm.EntityData.Children.Append("tcm4-iae", types.YChild{"Tcm4Iae", &alarm.Tcm4Iae})
    alarm.EntityData.Children.Append("tcm4-biae", types.YChild{"Tcm4Biae", &alarm.Tcm4Biae})
    alarm.EntityData.Children.Append("tcm4-bdi", types.YChild{"Tcm4Bdi", &alarm.Tcm4Bdi})
    alarm.EntityData.Children.Append("tcm4-tim", types.YChild{"Tcm4Tim", &alarm.Tcm4Tim})
    alarm.EntityData.Children.Append("tcm4-sf-ber", types.YChild{"Tcm4SfBer", &alarm.Tcm4SfBer})
    alarm.EntityData.Children.Append("tcm4-sd-ber", types.YChild{"Tcm4SdBer", &alarm.Tcm4SdBer})
    alarm.EntityData.Children.Append("tcm5-ais", types.YChild{"Tcm5Ais", &alarm.Tcm5Ais})
    alarm.EntityData.Children.Append("tcm5-ltc", types.YChild{"Tcm5Ltc", &alarm.Tcm5Ltc})
    alarm.EntityData.Children.Append("tcm5-oci", types.YChild{"Tcm5Oci", &alarm.Tcm5Oci})
    alarm.EntityData.Children.Append("tcm5-lck", types.YChild{"Tcm5Lck", &alarm.Tcm5Lck})
    alarm.EntityData.Children.Append("tcm5-iae", types.YChild{"Tcm5Iae", &alarm.Tcm5Iae})
    alarm.EntityData.Children.Append("tcm5-biae", types.YChild{"Tcm5Biae", &alarm.Tcm5Biae})
    alarm.EntityData.Children.Append("tcm5-bdi", types.YChild{"Tcm5Bdi", &alarm.Tcm5Bdi})
    alarm.EntityData.Children.Append("tcm5-tim", types.YChild{"Tcm5Tim", &alarm.Tcm5Tim})
    alarm.EntityData.Children.Append("tcm5-sf-ber", types.YChild{"Tcm5SfBer", &alarm.Tcm5SfBer})
    alarm.EntityData.Children.Append("tcm5-sd-ber", types.YChild{"Tcm5SdBer", &alarm.Tcm5SdBer})
    alarm.EntityData.Children.Append("tcm6-ais", types.YChild{"Tcm6Ais", &alarm.Tcm6Ais})
    alarm.EntityData.Children.Append("tcm6-ltc", types.YChild{"Tcm6Ltc", &alarm.Tcm6Ltc})
    alarm.EntityData.Children.Append("tcm6-oci", types.YChild{"Tcm6Oci", &alarm.Tcm6Oci})
    alarm.EntityData.Children.Append("tcm6-lck", types.YChild{"Tcm6Lck", &alarm.Tcm6Lck})
    alarm.EntityData.Children.Append("tcm6-iae", types.YChild{"Tcm6Iae", &alarm.Tcm6Iae})
    alarm.EntityData.Children.Append("tcm6-biae", types.YChild{"Tcm6Biae", &alarm.Tcm6Biae})
    alarm.EntityData.Children.Append("tcm6-bdi", types.YChild{"Tcm6Bdi", &alarm.Tcm6Bdi})
    alarm.EntityData.Children.Append("tcm6-tim", types.YChild{"Tcm6Tim", &alarm.Tcm6Tim})
    alarm.EntityData.Children.Append("tcm6-sf-ber", types.YChild{"Tcm6SfBer", &alarm.Tcm6SfBer})
    alarm.EntityData.Children.Append("tcm6-sd-ber", types.YChild{"Tcm6SdBer", &alarm.Tcm6SdBer})
    alarm.EntityData.Children.Append("gfp-lfd", types.YChild{"GfpLfd", &alarm.GfpLfd})
    alarm.EntityData.Children.Append("gfp-locs", types.YChild{"GfpLocs", &alarm.GfpLocs})
    alarm.EntityData.Children.Append("gfp-loccs", types.YChild{"GfpLoccs", &alarm.GfpLoccs})
    alarm.EntityData.Children.Append("gfp-upm", types.YChild{"GfpUpm", &alarm.GfpUpm})
    alarm.EntityData.Leafs = types.NewOrderedMap()

    alarm.EntityData.YListKeys = []string {}

    return &(alarm.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Oci
// Open Connection Indiction
type Odu_Controllers_Controller_Info_Alarm_Oci struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (oci *Odu_Controllers_Controller_Info_Alarm_Oci) GetEntityData() *types.CommonEntityData {
    oci.EntityData.YFilter = oci.YFilter
    oci.EntityData.YangName = "oci"
    oci.EntityData.BundleName = "cisco_ios_xr"
    oci.EntityData.ParentYangName = "alarm"
    oci.EntityData.SegmentPath = "oci"
    oci.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + oci.EntityData.SegmentPath
    oci.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    oci.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    oci.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    oci.EntityData.Children = types.NewOrderedMap()
    oci.EntityData.Leafs = types.NewOrderedMap()
    oci.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", oci.ReportingEnabled})
    oci.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", oci.IsDetected})
    oci.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", oci.IsAsserted})
    oci.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", oci.Counter})

    oci.EntityData.YListKeys = []string {}

    return &(oci.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Ais
// Alarm Indication Signal
type Odu_Controllers_Controller_Info_Alarm_Ais struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (ais *Odu_Controllers_Controller_Info_Alarm_Ais) GetEntityData() *types.CommonEntityData {
    ais.EntityData.YFilter = ais.YFilter
    ais.EntityData.YangName = "ais"
    ais.EntityData.BundleName = "cisco_ios_xr"
    ais.EntityData.ParentYangName = "alarm"
    ais.EntityData.SegmentPath = "ais"
    ais.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + ais.EntityData.SegmentPath
    ais.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ais.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ais.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ais.EntityData.Children = types.NewOrderedMap()
    ais.EntityData.Leafs = types.NewOrderedMap()
    ais.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", ais.ReportingEnabled})
    ais.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", ais.IsDetected})
    ais.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", ais.IsAsserted})
    ais.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", ais.Counter})

    ais.EntityData.YListKeys = []string {}

    return &(ais.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Lck
// Upstream Connection Locked
type Odu_Controllers_Controller_Info_Alarm_Lck struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (lck *Odu_Controllers_Controller_Info_Alarm_Lck) GetEntityData() *types.CommonEntityData {
    lck.EntityData.YFilter = lck.YFilter
    lck.EntityData.YangName = "lck"
    lck.EntityData.BundleName = "cisco_ios_xr"
    lck.EntityData.ParentYangName = "alarm"
    lck.EntityData.SegmentPath = "lck"
    lck.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + lck.EntityData.SegmentPath
    lck.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lck.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lck.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lck.EntityData.Children = types.NewOrderedMap()
    lck.EntityData.Leafs = types.NewOrderedMap()
    lck.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", lck.ReportingEnabled})
    lck.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", lck.IsDetected})
    lck.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", lck.IsAsserted})
    lck.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", lck.Counter})

    lck.EntityData.YListKeys = []string {}

    return &(lck.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Bdi
// Backward Defect Indication
type Odu_Controllers_Controller_Info_Alarm_Bdi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (bdi *Odu_Controllers_Controller_Info_Alarm_Bdi) GetEntityData() *types.CommonEntityData {
    bdi.EntityData.YFilter = bdi.YFilter
    bdi.EntityData.YangName = "bdi"
    bdi.EntityData.BundleName = "cisco_ios_xr"
    bdi.EntityData.ParentYangName = "alarm"
    bdi.EntityData.SegmentPath = "bdi"
    bdi.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + bdi.EntityData.SegmentPath
    bdi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdi.EntityData.Children = types.NewOrderedMap()
    bdi.EntityData.Leafs = types.NewOrderedMap()
    bdi.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", bdi.ReportingEnabled})
    bdi.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", bdi.IsDetected})
    bdi.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", bdi.IsAsserted})
    bdi.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", bdi.Counter})

    bdi.EntityData.YListKeys = []string {}

    return &(bdi.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Eoc
// GCC End of Channel
type Odu_Controllers_Controller_Info_Alarm_Eoc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (eoc *Odu_Controllers_Controller_Info_Alarm_Eoc) GetEntityData() *types.CommonEntityData {
    eoc.EntityData.YFilter = eoc.YFilter
    eoc.EntityData.YangName = "eoc"
    eoc.EntityData.BundleName = "cisco_ios_xr"
    eoc.EntityData.ParentYangName = "alarm"
    eoc.EntityData.SegmentPath = "eoc"
    eoc.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + eoc.EntityData.SegmentPath
    eoc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eoc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eoc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eoc.EntityData.Children = types.NewOrderedMap()
    eoc.EntityData.Leafs = types.NewOrderedMap()
    eoc.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", eoc.ReportingEnabled})
    eoc.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", eoc.IsDetected})
    eoc.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", eoc.IsAsserted})
    eoc.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", eoc.Counter})

    eoc.EntityData.YListKeys = []string {}

    return &(eoc.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Ptim
// Payload Type Identifier Mismatch
type Odu_Controllers_Controller_Info_Alarm_Ptim struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (ptim *Odu_Controllers_Controller_Info_Alarm_Ptim) GetEntityData() *types.CommonEntityData {
    ptim.EntityData.YFilter = ptim.YFilter
    ptim.EntityData.YangName = "ptim"
    ptim.EntityData.BundleName = "cisco_ios_xr"
    ptim.EntityData.ParentYangName = "alarm"
    ptim.EntityData.SegmentPath = "ptim"
    ptim.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + ptim.EntityData.SegmentPath
    ptim.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ptim.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ptim.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ptim.EntityData.Children = types.NewOrderedMap()
    ptim.EntityData.Leafs = types.NewOrderedMap()
    ptim.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", ptim.ReportingEnabled})
    ptim.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", ptim.IsDetected})
    ptim.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", ptim.IsAsserted})
    ptim.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", ptim.Counter})

    ptim.EntityData.YListKeys = []string {}

    return &(ptim.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Tim
// Trace Identifier Mismatch information
type Odu_Controllers_Controller_Info_Alarm_Tim struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (tim *Odu_Controllers_Controller_Info_Alarm_Tim) GetEntityData() *types.CommonEntityData {
    tim.EntityData.YFilter = tim.YFilter
    tim.EntityData.YangName = "tim"
    tim.EntityData.BundleName = "cisco_ios_xr"
    tim.EntityData.ParentYangName = "alarm"
    tim.EntityData.SegmentPath = "tim"
    tim.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + tim.EntityData.SegmentPath
    tim.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tim.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tim.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tim.EntityData.Children = types.NewOrderedMap()
    tim.EntityData.Leafs = types.NewOrderedMap()
    tim.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", tim.ReportingEnabled})
    tim.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", tim.IsDetected})
    tim.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", tim.IsAsserted})
    tim.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", tim.Counter})

    tim.EntityData.YListKeys = []string {}

    return &(tim.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Iae
// Incoming Alignment Error
type Odu_Controllers_Controller_Info_Alarm_Iae struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (iae *Odu_Controllers_Controller_Info_Alarm_Iae) GetEntityData() *types.CommonEntityData {
    iae.EntityData.YFilter = iae.YFilter
    iae.EntityData.YangName = "iae"
    iae.EntityData.BundleName = "cisco_ios_xr"
    iae.EntityData.ParentYangName = "alarm"
    iae.EntityData.SegmentPath = "iae"
    iae.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + iae.EntityData.SegmentPath
    iae.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    iae.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    iae.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    iae.EntityData.Children = types.NewOrderedMap()
    iae.EntityData.Leafs = types.NewOrderedMap()
    iae.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", iae.ReportingEnabled})
    iae.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", iae.IsDetected})
    iae.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", iae.IsAsserted})
    iae.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", iae.Counter})

    iae.EntityData.YListKeys = []string {}

    return &(iae.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Biae
// Backward Incoming Alignment Error
type Odu_Controllers_Controller_Info_Alarm_Biae struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (biae *Odu_Controllers_Controller_Info_Alarm_Biae) GetEntityData() *types.CommonEntityData {
    biae.EntityData.YFilter = biae.YFilter
    biae.EntityData.YangName = "biae"
    biae.EntityData.BundleName = "cisco_ios_xr"
    biae.EntityData.ParentYangName = "alarm"
    biae.EntityData.SegmentPath = "biae"
    biae.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + biae.EntityData.SegmentPath
    biae.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    biae.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    biae.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    biae.EntityData.Children = types.NewOrderedMap()
    biae.EntityData.Leafs = types.NewOrderedMap()
    biae.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", biae.ReportingEnabled})
    biae.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", biae.IsDetected})
    biae.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", biae.IsAsserted})
    biae.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", biae.Counter})

    biae.EntityData.YListKeys = []string {}

    return &(biae.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_SfBer
// SF BER alarm
type Odu_Controllers_Controller_Info_Alarm_SfBer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (sfBer *Odu_Controllers_Controller_Info_Alarm_SfBer) GetEntityData() *types.CommonEntityData {
    sfBer.EntityData.YFilter = sfBer.YFilter
    sfBer.EntityData.YangName = "sf-ber"
    sfBer.EntityData.BundleName = "cisco_ios_xr"
    sfBer.EntityData.ParentYangName = "alarm"
    sfBer.EntityData.SegmentPath = "sf-ber"
    sfBer.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + sfBer.EntityData.SegmentPath
    sfBer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sfBer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sfBer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sfBer.EntityData.Children = types.NewOrderedMap()
    sfBer.EntityData.Leafs = types.NewOrderedMap()
    sfBer.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", sfBer.ReportingEnabled})
    sfBer.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", sfBer.IsDetected})
    sfBer.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", sfBer.IsAsserted})
    sfBer.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", sfBer.Counter})

    sfBer.EntityData.YListKeys = []string {}

    return &(sfBer.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_SdBer
// SD BER alarm
type Odu_Controllers_Controller_Info_Alarm_SdBer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (sdBer *Odu_Controllers_Controller_Info_Alarm_SdBer) GetEntityData() *types.CommonEntityData {
    sdBer.EntityData.YFilter = sdBer.YFilter
    sdBer.EntityData.YangName = "sd-ber"
    sdBer.EntityData.BundleName = "cisco_ios_xr"
    sdBer.EntityData.ParentYangName = "alarm"
    sdBer.EntityData.SegmentPath = "sd-ber"
    sdBer.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + sdBer.EntityData.SegmentPath
    sdBer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sdBer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sdBer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sdBer.EntityData.Children = types.NewOrderedMap()
    sdBer.EntityData.Leafs = types.NewOrderedMap()
    sdBer.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", sdBer.ReportingEnabled})
    sdBer.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", sdBer.IsDetected})
    sdBer.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", sdBer.IsAsserted})
    sdBer.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", sdBer.Counter})

    sdBer.EntityData.YListKeys = []string {}

    return &(sdBer.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Csf
// Client Signal Failure
type Odu_Controllers_Controller_Info_Alarm_Csf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (csf *Odu_Controllers_Controller_Info_Alarm_Csf) GetEntityData() *types.CommonEntityData {
    csf.EntityData.YFilter = csf.YFilter
    csf.EntityData.YangName = "csf"
    csf.EntityData.BundleName = "cisco_ios_xr"
    csf.EntityData.ParentYangName = "alarm"
    csf.EntityData.SegmentPath = "csf"
    csf.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + csf.EntityData.SegmentPath
    csf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    csf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    csf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    csf.EntityData.Children = types.NewOrderedMap()
    csf.EntityData.Leafs = types.NewOrderedMap()
    csf.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", csf.ReportingEnabled})
    csf.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", csf.IsDetected})
    csf.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", csf.IsAsserted})
    csf.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", csf.Counter})

    csf.EntityData.YListKeys = []string {}

    return &(csf.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Tcm1Ais
// TCM1 Alarm Indication Signal
type Odu_Controllers_Controller_Info_Alarm_Tcm1Ais struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (tcm1Ais *Odu_Controllers_Controller_Info_Alarm_Tcm1Ais) GetEntityData() *types.CommonEntityData {
    tcm1Ais.EntityData.YFilter = tcm1Ais.YFilter
    tcm1Ais.EntityData.YangName = "tcm1-ais"
    tcm1Ais.EntityData.BundleName = "cisco_ios_xr"
    tcm1Ais.EntityData.ParentYangName = "alarm"
    tcm1Ais.EntityData.SegmentPath = "tcm1-ais"
    tcm1Ais.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + tcm1Ais.EntityData.SegmentPath
    tcm1Ais.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcm1Ais.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcm1Ais.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcm1Ais.EntityData.Children = types.NewOrderedMap()
    tcm1Ais.EntityData.Leafs = types.NewOrderedMap()
    tcm1Ais.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", tcm1Ais.ReportingEnabled})
    tcm1Ais.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", tcm1Ais.IsDetected})
    tcm1Ais.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", tcm1Ais.IsAsserted})
    tcm1Ais.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", tcm1Ais.Counter})

    tcm1Ais.EntityData.YListKeys = []string {}

    return &(tcm1Ais.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Tcm1Ltc
// TCM1 Loss of Tandem connection Monitoring
type Odu_Controllers_Controller_Info_Alarm_Tcm1Ltc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (tcm1Ltc *Odu_Controllers_Controller_Info_Alarm_Tcm1Ltc) GetEntityData() *types.CommonEntityData {
    tcm1Ltc.EntityData.YFilter = tcm1Ltc.YFilter
    tcm1Ltc.EntityData.YangName = "tcm1-ltc"
    tcm1Ltc.EntityData.BundleName = "cisco_ios_xr"
    tcm1Ltc.EntityData.ParentYangName = "alarm"
    tcm1Ltc.EntityData.SegmentPath = "tcm1-ltc"
    tcm1Ltc.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + tcm1Ltc.EntityData.SegmentPath
    tcm1Ltc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcm1Ltc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcm1Ltc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcm1Ltc.EntityData.Children = types.NewOrderedMap()
    tcm1Ltc.EntityData.Leafs = types.NewOrderedMap()
    tcm1Ltc.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", tcm1Ltc.ReportingEnabled})
    tcm1Ltc.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", tcm1Ltc.IsDetected})
    tcm1Ltc.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", tcm1Ltc.IsAsserted})
    tcm1Ltc.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", tcm1Ltc.Counter})

    tcm1Ltc.EntityData.YListKeys = []string {}

    return &(tcm1Ltc.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Tcm1Oci
// TCM1 Open Connection Indiction
type Odu_Controllers_Controller_Info_Alarm_Tcm1Oci struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (tcm1Oci *Odu_Controllers_Controller_Info_Alarm_Tcm1Oci) GetEntityData() *types.CommonEntityData {
    tcm1Oci.EntityData.YFilter = tcm1Oci.YFilter
    tcm1Oci.EntityData.YangName = "tcm1-oci"
    tcm1Oci.EntityData.BundleName = "cisco_ios_xr"
    tcm1Oci.EntityData.ParentYangName = "alarm"
    tcm1Oci.EntityData.SegmentPath = "tcm1-oci"
    tcm1Oci.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + tcm1Oci.EntityData.SegmentPath
    tcm1Oci.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcm1Oci.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcm1Oci.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcm1Oci.EntityData.Children = types.NewOrderedMap()
    tcm1Oci.EntityData.Leafs = types.NewOrderedMap()
    tcm1Oci.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", tcm1Oci.ReportingEnabled})
    tcm1Oci.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", tcm1Oci.IsDetected})
    tcm1Oci.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", tcm1Oci.IsAsserted})
    tcm1Oci.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", tcm1Oci.Counter})

    tcm1Oci.EntityData.YListKeys = []string {}

    return &(tcm1Oci.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Tcm1Lck
// TCM1  Upstream Connection Locked
type Odu_Controllers_Controller_Info_Alarm_Tcm1Lck struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (tcm1Lck *Odu_Controllers_Controller_Info_Alarm_Tcm1Lck) GetEntityData() *types.CommonEntityData {
    tcm1Lck.EntityData.YFilter = tcm1Lck.YFilter
    tcm1Lck.EntityData.YangName = "tcm1-lck"
    tcm1Lck.EntityData.BundleName = "cisco_ios_xr"
    tcm1Lck.EntityData.ParentYangName = "alarm"
    tcm1Lck.EntityData.SegmentPath = "tcm1-lck"
    tcm1Lck.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + tcm1Lck.EntityData.SegmentPath
    tcm1Lck.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcm1Lck.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcm1Lck.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcm1Lck.EntityData.Children = types.NewOrderedMap()
    tcm1Lck.EntityData.Leafs = types.NewOrderedMap()
    tcm1Lck.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", tcm1Lck.ReportingEnabled})
    tcm1Lck.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", tcm1Lck.IsDetected})
    tcm1Lck.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", tcm1Lck.IsAsserted})
    tcm1Lck.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", tcm1Lck.Counter})

    tcm1Lck.EntityData.YListKeys = []string {}

    return &(tcm1Lck.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Tcm1Iae
// TCM1 Incoming Alignment Error
type Odu_Controllers_Controller_Info_Alarm_Tcm1Iae struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (tcm1Iae *Odu_Controllers_Controller_Info_Alarm_Tcm1Iae) GetEntityData() *types.CommonEntityData {
    tcm1Iae.EntityData.YFilter = tcm1Iae.YFilter
    tcm1Iae.EntityData.YangName = "tcm1-iae"
    tcm1Iae.EntityData.BundleName = "cisco_ios_xr"
    tcm1Iae.EntityData.ParentYangName = "alarm"
    tcm1Iae.EntityData.SegmentPath = "tcm1-iae"
    tcm1Iae.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + tcm1Iae.EntityData.SegmentPath
    tcm1Iae.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcm1Iae.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcm1Iae.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcm1Iae.EntityData.Children = types.NewOrderedMap()
    tcm1Iae.EntityData.Leafs = types.NewOrderedMap()
    tcm1Iae.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", tcm1Iae.ReportingEnabled})
    tcm1Iae.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", tcm1Iae.IsDetected})
    tcm1Iae.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", tcm1Iae.IsAsserted})
    tcm1Iae.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", tcm1Iae.Counter})

    tcm1Iae.EntityData.YListKeys = []string {}

    return &(tcm1Iae.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Tcm1Biae
// TCM1 Backward Incoming Alignment Error
type Odu_Controllers_Controller_Info_Alarm_Tcm1Biae struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (tcm1Biae *Odu_Controllers_Controller_Info_Alarm_Tcm1Biae) GetEntityData() *types.CommonEntityData {
    tcm1Biae.EntityData.YFilter = tcm1Biae.YFilter
    tcm1Biae.EntityData.YangName = "tcm1-biae"
    tcm1Biae.EntityData.BundleName = "cisco_ios_xr"
    tcm1Biae.EntityData.ParentYangName = "alarm"
    tcm1Biae.EntityData.SegmentPath = "tcm1-biae"
    tcm1Biae.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + tcm1Biae.EntityData.SegmentPath
    tcm1Biae.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcm1Biae.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcm1Biae.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcm1Biae.EntityData.Children = types.NewOrderedMap()
    tcm1Biae.EntityData.Leafs = types.NewOrderedMap()
    tcm1Biae.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", tcm1Biae.ReportingEnabled})
    tcm1Biae.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", tcm1Biae.IsDetected})
    tcm1Biae.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", tcm1Biae.IsAsserted})
    tcm1Biae.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", tcm1Biae.Counter})

    tcm1Biae.EntityData.YListKeys = []string {}

    return &(tcm1Biae.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Tcm1Bdi
// TCM1 Backward Defect Monitoring
type Odu_Controllers_Controller_Info_Alarm_Tcm1Bdi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (tcm1Bdi *Odu_Controllers_Controller_Info_Alarm_Tcm1Bdi) GetEntityData() *types.CommonEntityData {
    tcm1Bdi.EntityData.YFilter = tcm1Bdi.YFilter
    tcm1Bdi.EntityData.YangName = "tcm1-bdi"
    tcm1Bdi.EntityData.BundleName = "cisco_ios_xr"
    tcm1Bdi.EntityData.ParentYangName = "alarm"
    tcm1Bdi.EntityData.SegmentPath = "tcm1-bdi"
    tcm1Bdi.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + tcm1Bdi.EntityData.SegmentPath
    tcm1Bdi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcm1Bdi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcm1Bdi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcm1Bdi.EntityData.Children = types.NewOrderedMap()
    tcm1Bdi.EntityData.Leafs = types.NewOrderedMap()
    tcm1Bdi.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", tcm1Bdi.ReportingEnabled})
    tcm1Bdi.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", tcm1Bdi.IsDetected})
    tcm1Bdi.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", tcm1Bdi.IsAsserted})
    tcm1Bdi.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", tcm1Bdi.Counter})

    tcm1Bdi.EntityData.YListKeys = []string {}

    return &(tcm1Bdi.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Tcm1Tim
// TCM1 Trail Trace Identifier Mismatch 
type Odu_Controllers_Controller_Info_Alarm_Tcm1Tim struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (tcm1Tim *Odu_Controllers_Controller_Info_Alarm_Tcm1Tim) GetEntityData() *types.CommonEntityData {
    tcm1Tim.EntityData.YFilter = tcm1Tim.YFilter
    tcm1Tim.EntityData.YangName = "tcm1-tim"
    tcm1Tim.EntityData.BundleName = "cisco_ios_xr"
    tcm1Tim.EntityData.ParentYangName = "alarm"
    tcm1Tim.EntityData.SegmentPath = "tcm1-tim"
    tcm1Tim.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + tcm1Tim.EntityData.SegmentPath
    tcm1Tim.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcm1Tim.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcm1Tim.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcm1Tim.EntityData.Children = types.NewOrderedMap()
    tcm1Tim.EntityData.Leafs = types.NewOrderedMap()
    tcm1Tim.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", tcm1Tim.ReportingEnabled})
    tcm1Tim.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", tcm1Tim.IsDetected})
    tcm1Tim.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", tcm1Tim.IsAsserted})
    tcm1Tim.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", tcm1Tim.Counter})

    tcm1Tim.EntityData.YListKeys = []string {}

    return &(tcm1Tim.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Tcm1SfBer
// TCM1 SF BER alarm
type Odu_Controllers_Controller_Info_Alarm_Tcm1SfBer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (tcm1SfBer *Odu_Controllers_Controller_Info_Alarm_Tcm1SfBer) GetEntityData() *types.CommonEntityData {
    tcm1SfBer.EntityData.YFilter = tcm1SfBer.YFilter
    tcm1SfBer.EntityData.YangName = "tcm1-sf-ber"
    tcm1SfBer.EntityData.BundleName = "cisco_ios_xr"
    tcm1SfBer.EntityData.ParentYangName = "alarm"
    tcm1SfBer.EntityData.SegmentPath = "tcm1-sf-ber"
    tcm1SfBer.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + tcm1SfBer.EntityData.SegmentPath
    tcm1SfBer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcm1SfBer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcm1SfBer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcm1SfBer.EntityData.Children = types.NewOrderedMap()
    tcm1SfBer.EntityData.Leafs = types.NewOrderedMap()
    tcm1SfBer.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", tcm1SfBer.ReportingEnabled})
    tcm1SfBer.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", tcm1SfBer.IsDetected})
    tcm1SfBer.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", tcm1SfBer.IsAsserted})
    tcm1SfBer.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", tcm1SfBer.Counter})

    tcm1SfBer.EntityData.YListKeys = []string {}

    return &(tcm1SfBer.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Tcm1SdBer
// TCM1 SD BER alarm
type Odu_Controllers_Controller_Info_Alarm_Tcm1SdBer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (tcm1SdBer *Odu_Controllers_Controller_Info_Alarm_Tcm1SdBer) GetEntityData() *types.CommonEntityData {
    tcm1SdBer.EntityData.YFilter = tcm1SdBer.YFilter
    tcm1SdBer.EntityData.YangName = "tcm1-sd-ber"
    tcm1SdBer.EntityData.BundleName = "cisco_ios_xr"
    tcm1SdBer.EntityData.ParentYangName = "alarm"
    tcm1SdBer.EntityData.SegmentPath = "tcm1-sd-ber"
    tcm1SdBer.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + tcm1SdBer.EntityData.SegmentPath
    tcm1SdBer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcm1SdBer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcm1SdBer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcm1SdBer.EntityData.Children = types.NewOrderedMap()
    tcm1SdBer.EntityData.Leafs = types.NewOrderedMap()
    tcm1SdBer.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", tcm1SdBer.ReportingEnabled})
    tcm1SdBer.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", tcm1SdBer.IsDetected})
    tcm1SdBer.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", tcm1SdBer.IsAsserted})
    tcm1SdBer.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", tcm1SdBer.Counter})

    tcm1SdBer.EntityData.YListKeys = []string {}

    return &(tcm1SdBer.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Tcm2Ais
// TCM2 Alarm Indication Signal
type Odu_Controllers_Controller_Info_Alarm_Tcm2Ais struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (tcm2Ais *Odu_Controllers_Controller_Info_Alarm_Tcm2Ais) GetEntityData() *types.CommonEntityData {
    tcm2Ais.EntityData.YFilter = tcm2Ais.YFilter
    tcm2Ais.EntityData.YangName = "tcm2-ais"
    tcm2Ais.EntityData.BundleName = "cisco_ios_xr"
    tcm2Ais.EntityData.ParentYangName = "alarm"
    tcm2Ais.EntityData.SegmentPath = "tcm2-ais"
    tcm2Ais.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + tcm2Ais.EntityData.SegmentPath
    tcm2Ais.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcm2Ais.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcm2Ais.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcm2Ais.EntityData.Children = types.NewOrderedMap()
    tcm2Ais.EntityData.Leafs = types.NewOrderedMap()
    tcm2Ais.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", tcm2Ais.ReportingEnabled})
    tcm2Ais.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", tcm2Ais.IsDetected})
    tcm2Ais.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", tcm2Ais.IsAsserted})
    tcm2Ais.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", tcm2Ais.Counter})

    tcm2Ais.EntityData.YListKeys = []string {}

    return &(tcm2Ais.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Tcm2Ltc
// TCM2 Loss of Tandem connection Monitoring
type Odu_Controllers_Controller_Info_Alarm_Tcm2Ltc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (tcm2Ltc *Odu_Controllers_Controller_Info_Alarm_Tcm2Ltc) GetEntityData() *types.CommonEntityData {
    tcm2Ltc.EntityData.YFilter = tcm2Ltc.YFilter
    tcm2Ltc.EntityData.YangName = "tcm2-ltc"
    tcm2Ltc.EntityData.BundleName = "cisco_ios_xr"
    tcm2Ltc.EntityData.ParentYangName = "alarm"
    tcm2Ltc.EntityData.SegmentPath = "tcm2-ltc"
    tcm2Ltc.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + tcm2Ltc.EntityData.SegmentPath
    tcm2Ltc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcm2Ltc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcm2Ltc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcm2Ltc.EntityData.Children = types.NewOrderedMap()
    tcm2Ltc.EntityData.Leafs = types.NewOrderedMap()
    tcm2Ltc.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", tcm2Ltc.ReportingEnabled})
    tcm2Ltc.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", tcm2Ltc.IsDetected})
    tcm2Ltc.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", tcm2Ltc.IsAsserted})
    tcm2Ltc.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", tcm2Ltc.Counter})

    tcm2Ltc.EntityData.YListKeys = []string {}

    return &(tcm2Ltc.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Tcm2Oci
// TCM2 Open Connection Indiction
type Odu_Controllers_Controller_Info_Alarm_Tcm2Oci struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (tcm2Oci *Odu_Controllers_Controller_Info_Alarm_Tcm2Oci) GetEntityData() *types.CommonEntityData {
    tcm2Oci.EntityData.YFilter = tcm2Oci.YFilter
    tcm2Oci.EntityData.YangName = "tcm2-oci"
    tcm2Oci.EntityData.BundleName = "cisco_ios_xr"
    tcm2Oci.EntityData.ParentYangName = "alarm"
    tcm2Oci.EntityData.SegmentPath = "tcm2-oci"
    tcm2Oci.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + tcm2Oci.EntityData.SegmentPath
    tcm2Oci.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcm2Oci.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcm2Oci.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcm2Oci.EntityData.Children = types.NewOrderedMap()
    tcm2Oci.EntityData.Leafs = types.NewOrderedMap()
    tcm2Oci.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", tcm2Oci.ReportingEnabled})
    tcm2Oci.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", tcm2Oci.IsDetected})
    tcm2Oci.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", tcm2Oci.IsAsserted})
    tcm2Oci.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", tcm2Oci.Counter})

    tcm2Oci.EntityData.YListKeys = []string {}

    return &(tcm2Oci.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Tcm2Lck
// TCM2  Upstream Connection Locked
type Odu_Controllers_Controller_Info_Alarm_Tcm2Lck struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (tcm2Lck *Odu_Controllers_Controller_Info_Alarm_Tcm2Lck) GetEntityData() *types.CommonEntityData {
    tcm2Lck.EntityData.YFilter = tcm2Lck.YFilter
    tcm2Lck.EntityData.YangName = "tcm2-lck"
    tcm2Lck.EntityData.BundleName = "cisco_ios_xr"
    tcm2Lck.EntityData.ParentYangName = "alarm"
    tcm2Lck.EntityData.SegmentPath = "tcm2-lck"
    tcm2Lck.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + tcm2Lck.EntityData.SegmentPath
    tcm2Lck.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcm2Lck.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcm2Lck.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcm2Lck.EntityData.Children = types.NewOrderedMap()
    tcm2Lck.EntityData.Leafs = types.NewOrderedMap()
    tcm2Lck.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", tcm2Lck.ReportingEnabled})
    tcm2Lck.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", tcm2Lck.IsDetected})
    tcm2Lck.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", tcm2Lck.IsAsserted})
    tcm2Lck.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", tcm2Lck.Counter})

    tcm2Lck.EntityData.YListKeys = []string {}

    return &(tcm2Lck.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Tcm2Iae
// TCM2 Incoming Alignment Error
type Odu_Controllers_Controller_Info_Alarm_Tcm2Iae struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (tcm2Iae *Odu_Controllers_Controller_Info_Alarm_Tcm2Iae) GetEntityData() *types.CommonEntityData {
    tcm2Iae.EntityData.YFilter = tcm2Iae.YFilter
    tcm2Iae.EntityData.YangName = "tcm2-iae"
    tcm2Iae.EntityData.BundleName = "cisco_ios_xr"
    tcm2Iae.EntityData.ParentYangName = "alarm"
    tcm2Iae.EntityData.SegmentPath = "tcm2-iae"
    tcm2Iae.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + tcm2Iae.EntityData.SegmentPath
    tcm2Iae.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcm2Iae.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcm2Iae.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcm2Iae.EntityData.Children = types.NewOrderedMap()
    tcm2Iae.EntityData.Leafs = types.NewOrderedMap()
    tcm2Iae.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", tcm2Iae.ReportingEnabled})
    tcm2Iae.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", tcm2Iae.IsDetected})
    tcm2Iae.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", tcm2Iae.IsAsserted})
    tcm2Iae.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", tcm2Iae.Counter})

    tcm2Iae.EntityData.YListKeys = []string {}

    return &(tcm2Iae.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Tcm2Biae
// TCM2 Backward Incoming Alignment Error
type Odu_Controllers_Controller_Info_Alarm_Tcm2Biae struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (tcm2Biae *Odu_Controllers_Controller_Info_Alarm_Tcm2Biae) GetEntityData() *types.CommonEntityData {
    tcm2Biae.EntityData.YFilter = tcm2Biae.YFilter
    tcm2Biae.EntityData.YangName = "tcm2-biae"
    tcm2Biae.EntityData.BundleName = "cisco_ios_xr"
    tcm2Biae.EntityData.ParentYangName = "alarm"
    tcm2Biae.EntityData.SegmentPath = "tcm2-biae"
    tcm2Biae.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + tcm2Biae.EntityData.SegmentPath
    tcm2Biae.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcm2Biae.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcm2Biae.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcm2Biae.EntityData.Children = types.NewOrderedMap()
    tcm2Biae.EntityData.Leafs = types.NewOrderedMap()
    tcm2Biae.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", tcm2Biae.ReportingEnabled})
    tcm2Biae.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", tcm2Biae.IsDetected})
    tcm2Biae.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", tcm2Biae.IsAsserted})
    tcm2Biae.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", tcm2Biae.Counter})

    tcm2Biae.EntityData.YListKeys = []string {}

    return &(tcm2Biae.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Tcm2Bdi
// TCM2 Backward Defect Monitoring
type Odu_Controllers_Controller_Info_Alarm_Tcm2Bdi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (tcm2Bdi *Odu_Controllers_Controller_Info_Alarm_Tcm2Bdi) GetEntityData() *types.CommonEntityData {
    tcm2Bdi.EntityData.YFilter = tcm2Bdi.YFilter
    tcm2Bdi.EntityData.YangName = "tcm2-bdi"
    tcm2Bdi.EntityData.BundleName = "cisco_ios_xr"
    tcm2Bdi.EntityData.ParentYangName = "alarm"
    tcm2Bdi.EntityData.SegmentPath = "tcm2-bdi"
    tcm2Bdi.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + tcm2Bdi.EntityData.SegmentPath
    tcm2Bdi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcm2Bdi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcm2Bdi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcm2Bdi.EntityData.Children = types.NewOrderedMap()
    tcm2Bdi.EntityData.Leafs = types.NewOrderedMap()
    tcm2Bdi.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", tcm2Bdi.ReportingEnabled})
    tcm2Bdi.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", tcm2Bdi.IsDetected})
    tcm2Bdi.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", tcm2Bdi.IsAsserted})
    tcm2Bdi.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", tcm2Bdi.Counter})

    tcm2Bdi.EntityData.YListKeys = []string {}

    return &(tcm2Bdi.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Tcm2Tim
// TCM2 Trail Trace Identifier Mismatch 
type Odu_Controllers_Controller_Info_Alarm_Tcm2Tim struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (tcm2Tim *Odu_Controllers_Controller_Info_Alarm_Tcm2Tim) GetEntityData() *types.CommonEntityData {
    tcm2Tim.EntityData.YFilter = tcm2Tim.YFilter
    tcm2Tim.EntityData.YangName = "tcm2-tim"
    tcm2Tim.EntityData.BundleName = "cisco_ios_xr"
    tcm2Tim.EntityData.ParentYangName = "alarm"
    tcm2Tim.EntityData.SegmentPath = "tcm2-tim"
    tcm2Tim.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + tcm2Tim.EntityData.SegmentPath
    tcm2Tim.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcm2Tim.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcm2Tim.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcm2Tim.EntityData.Children = types.NewOrderedMap()
    tcm2Tim.EntityData.Leafs = types.NewOrderedMap()
    tcm2Tim.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", tcm2Tim.ReportingEnabled})
    tcm2Tim.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", tcm2Tim.IsDetected})
    tcm2Tim.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", tcm2Tim.IsAsserted})
    tcm2Tim.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", tcm2Tim.Counter})

    tcm2Tim.EntityData.YListKeys = []string {}

    return &(tcm2Tim.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Tcm2SfBer
// TCM2 SF BER alarm
type Odu_Controllers_Controller_Info_Alarm_Tcm2SfBer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (tcm2SfBer *Odu_Controllers_Controller_Info_Alarm_Tcm2SfBer) GetEntityData() *types.CommonEntityData {
    tcm2SfBer.EntityData.YFilter = tcm2SfBer.YFilter
    tcm2SfBer.EntityData.YangName = "tcm2-sf-ber"
    tcm2SfBer.EntityData.BundleName = "cisco_ios_xr"
    tcm2SfBer.EntityData.ParentYangName = "alarm"
    tcm2SfBer.EntityData.SegmentPath = "tcm2-sf-ber"
    tcm2SfBer.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + tcm2SfBer.EntityData.SegmentPath
    tcm2SfBer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcm2SfBer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcm2SfBer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcm2SfBer.EntityData.Children = types.NewOrderedMap()
    tcm2SfBer.EntityData.Leafs = types.NewOrderedMap()
    tcm2SfBer.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", tcm2SfBer.ReportingEnabled})
    tcm2SfBer.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", tcm2SfBer.IsDetected})
    tcm2SfBer.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", tcm2SfBer.IsAsserted})
    tcm2SfBer.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", tcm2SfBer.Counter})

    tcm2SfBer.EntityData.YListKeys = []string {}

    return &(tcm2SfBer.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Tcm2SdBer
// TCM2 SD BER alarm
type Odu_Controllers_Controller_Info_Alarm_Tcm2SdBer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (tcm2SdBer *Odu_Controllers_Controller_Info_Alarm_Tcm2SdBer) GetEntityData() *types.CommonEntityData {
    tcm2SdBer.EntityData.YFilter = tcm2SdBer.YFilter
    tcm2SdBer.EntityData.YangName = "tcm2-sd-ber"
    tcm2SdBer.EntityData.BundleName = "cisco_ios_xr"
    tcm2SdBer.EntityData.ParentYangName = "alarm"
    tcm2SdBer.EntityData.SegmentPath = "tcm2-sd-ber"
    tcm2SdBer.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + tcm2SdBer.EntityData.SegmentPath
    tcm2SdBer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcm2SdBer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcm2SdBer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcm2SdBer.EntityData.Children = types.NewOrderedMap()
    tcm2SdBer.EntityData.Leafs = types.NewOrderedMap()
    tcm2SdBer.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", tcm2SdBer.ReportingEnabled})
    tcm2SdBer.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", tcm2SdBer.IsDetected})
    tcm2SdBer.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", tcm2SdBer.IsAsserted})
    tcm2SdBer.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", tcm2SdBer.Counter})

    tcm2SdBer.EntityData.YListKeys = []string {}

    return &(tcm2SdBer.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Tcm3Ais
// TCM3 Alarm Indication Signal
type Odu_Controllers_Controller_Info_Alarm_Tcm3Ais struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (tcm3Ais *Odu_Controllers_Controller_Info_Alarm_Tcm3Ais) GetEntityData() *types.CommonEntityData {
    tcm3Ais.EntityData.YFilter = tcm3Ais.YFilter
    tcm3Ais.EntityData.YangName = "tcm3-ais"
    tcm3Ais.EntityData.BundleName = "cisco_ios_xr"
    tcm3Ais.EntityData.ParentYangName = "alarm"
    tcm3Ais.EntityData.SegmentPath = "tcm3-ais"
    tcm3Ais.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + tcm3Ais.EntityData.SegmentPath
    tcm3Ais.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcm3Ais.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcm3Ais.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcm3Ais.EntityData.Children = types.NewOrderedMap()
    tcm3Ais.EntityData.Leafs = types.NewOrderedMap()
    tcm3Ais.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", tcm3Ais.ReportingEnabled})
    tcm3Ais.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", tcm3Ais.IsDetected})
    tcm3Ais.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", tcm3Ais.IsAsserted})
    tcm3Ais.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", tcm3Ais.Counter})

    tcm3Ais.EntityData.YListKeys = []string {}

    return &(tcm3Ais.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Tcm3Ltc
// TCM3 Loss of Tandem connection Monitoring
type Odu_Controllers_Controller_Info_Alarm_Tcm3Ltc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (tcm3Ltc *Odu_Controllers_Controller_Info_Alarm_Tcm3Ltc) GetEntityData() *types.CommonEntityData {
    tcm3Ltc.EntityData.YFilter = tcm3Ltc.YFilter
    tcm3Ltc.EntityData.YangName = "tcm3-ltc"
    tcm3Ltc.EntityData.BundleName = "cisco_ios_xr"
    tcm3Ltc.EntityData.ParentYangName = "alarm"
    tcm3Ltc.EntityData.SegmentPath = "tcm3-ltc"
    tcm3Ltc.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + tcm3Ltc.EntityData.SegmentPath
    tcm3Ltc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcm3Ltc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcm3Ltc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcm3Ltc.EntityData.Children = types.NewOrderedMap()
    tcm3Ltc.EntityData.Leafs = types.NewOrderedMap()
    tcm3Ltc.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", tcm3Ltc.ReportingEnabled})
    tcm3Ltc.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", tcm3Ltc.IsDetected})
    tcm3Ltc.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", tcm3Ltc.IsAsserted})
    tcm3Ltc.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", tcm3Ltc.Counter})

    tcm3Ltc.EntityData.YListKeys = []string {}

    return &(tcm3Ltc.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Tcm3Oci
// TCM3 Open Connection Indiction
type Odu_Controllers_Controller_Info_Alarm_Tcm3Oci struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (tcm3Oci *Odu_Controllers_Controller_Info_Alarm_Tcm3Oci) GetEntityData() *types.CommonEntityData {
    tcm3Oci.EntityData.YFilter = tcm3Oci.YFilter
    tcm3Oci.EntityData.YangName = "tcm3-oci"
    tcm3Oci.EntityData.BundleName = "cisco_ios_xr"
    tcm3Oci.EntityData.ParentYangName = "alarm"
    tcm3Oci.EntityData.SegmentPath = "tcm3-oci"
    tcm3Oci.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + tcm3Oci.EntityData.SegmentPath
    tcm3Oci.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcm3Oci.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcm3Oci.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcm3Oci.EntityData.Children = types.NewOrderedMap()
    tcm3Oci.EntityData.Leafs = types.NewOrderedMap()
    tcm3Oci.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", tcm3Oci.ReportingEnabled})
    tcm3Oci.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", tcm3Oci.IsDetected})
    tcm3Oci.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", tcm3Oci.IsAsserted})
    tcm3Oci.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", tcm3Oci.Counter})

    tcm3Oci.EntityData.YListKeys = []string {}

    return &(tcm3Oci.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Tcm3Lck
// TCM3  Upstream Connection Locked
type Odu_Controllers_Controller_Info_Alarm_Tcm3Lck struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (tcm3Lck *Odu_Controllers_Controller_Info_Alarm_Tcm3Lck) GetEntityData() *types.CommonEntityData {
    tcm3Lck.EntityData.YFilter = tcm3Lck.YFilter
    tcm3Lck.EntityData.YangName = "tcm3-lck"
    tcm3Lck.EntityData.BundleName = "cisco_ios_xr"
    tcm3Lck.EntityData.ParentYangName = "alarm"
    tcm3Lck.EntityData.SegmentPath = "tcm3-lck"
    tcm3Lck.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + tcm3Lck.EntityData.SegmentPath
    tcm3Lck.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcm3Lck.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcm3Lck.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcm3Lck.EntityData.Children = types.NewOrderedMap()
    tcm3Lck.EntityData.Leafs = types.NewOrderedMap()
    tcm3Lck.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", tcm3Lck.ReportingEnabled})
    tcm3Lck.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", tcm3Lck.IsDetected})
    tcm3Lck.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", tcm3Lck.IsAsserted})
    tcm3Lck.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", tcm3Lck.Counter})

    tcm3Lck.EntityData.YListKeys = []string {}

    return &(tcm3Lck.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Tcm3Iae
// TCM3 Incoming Alignment Error
type Odu_Controllers_Controller_Info_Alarm_Tcm3Iae struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (tcm3Iae *Odu_Controllers_Controller_Info_Alarm_Tcm3Iae) GetEntityData() *types.CommonEntityData {
    tcm3Iae.EntityData.YFilter = tcm3Iae.YFilter
    tcm3Iae.EntityData.YangName = "tcm3-iae"
    tcm3Iae.EntityData.BundleName = "cisco_ios_xr"
    tcm3Iae.EntityData.ParentYangName = "alarm"
    tcm3Iae.EntityData.SegmentPath = "tcm3-iae"
    tcm3Iae.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + tcm3Iae.EntityData.SegmentPath
    tcm3Iae.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcm3Iae.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcm3Iae.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcm3Iae.EntityData.Children = types.NewOrderedMap()
    tcm3Iae.EntityData.Leafs = types.NewOrderedMap()
    tcm3Iae.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", tcm3Iae.ReportingEnabled})
    tcm3Iae.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", tcm3Iae.IsDetected})
    tcm3Iae.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", tcm3Iae.IsAsserted})
    tcm3Iae.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", tcm3Iae.Counter})

    tcm3Iae.EntityData.YListKeys = []string {}

    return &(tcm3Iae.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Tcm3Biae
// TCM3 Backward Incoming Alignment Error
type Odu_Controllers_Controller_Info_Alarm_Tcm3Biae struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (tcm3Biae *Odu_Controllers_Controller_Info_Alarm_Tcm3Biae) GetEntityData() *types.CommonEntityData {
    tcm3Biae.EntityData.YFilter = tcm3Biae.YFilter
    tcm3Biae.EntityData.YangName = "tcm3-biae"
    tcm3Biae.EntityData.BundleName = "cisco_ios_xr"
    tcm3Biae.EntityData.ParentYangName = "alarm"
    tcm3Biae.EntityData.SegmentPath = "tcm3-biae"
    tcm3Biae.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + tcm3Biae.EntityData.SegmentPath
    tcm3Biae.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcm3Biae.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcm3Biae.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcm3Biae.EntityData.Children = types.NewOrderedMap()
    tcm3Biae.EntityData.Leafs = types.NewOrderedMap()
    tcm3Biae.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", tcm3Biae.ReportingEnabled})
    tcm3Biae.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", tcm3Biae.IsDetected})
    tcm3Biae.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", tcm3Biae.IsAsserted})
    tcm3Biae.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", tcm3Biae.Counter})

    tcm3Biae.EntityData.YListKeys = []string {}

    return &(tcm3Biae.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Tcm3Bdi
// TCM3 Backward Defect Monitoring
type Odu_Controllers_Controller_Info_Alarm_Tcm3Bdi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (tcm3Bdi *Odu_Controllers_Controller_Info_Alarm_Tcm3Bdi) GetEntityData() *types.CommonEntityData {
    tcm3Bdi.EntityData.YFilter = tcm3Bdi.YFilter
    tcm3Bdi.EntityData.YangName = "tcm3-bdi"
    tcm3Bdi.EntityData.BundleName = "cisco_ios_xr"
    tcm3Bdi.EntityData.ParentYangName = "alarm"
    tcm3Bdi.EntityData.SegmentPath = "tcm3-bdi"
    tcm3Bdi.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + tcm3Bdi.EntityData.SegmentPath
    tcm3Bdi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcm3Bdi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcm3Bdi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcm3Bdi.EntityData.Children = types.NewOrderedMap()
    tcm3Bdi.EntityData.Leafs = types.NewOrderedMap()
    tcm3Bdi.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", tcm3Bdi.ReportingEnabled})
    tcm3Bdi.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", tcm3Bdi.IsDetected})
    tcm3Bdi.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", tcm3Bdi.IsAsserted})
    tcm3Bdi.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", tcm3Bdi.Counter})

    tcm3Bdi.EntityData.YListKeys = []string {}

    return &(tcm3Bdi.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Tcm3Tim
// TCM3 Trail Trace Identifier Mismatch 
type Odu_Controllers_Controller_Info_Alarm_Tcm3Tim struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (tcm3Tim *Odu_Controllers_Controller_Info_Alarm_Tcm3Tim) GetEntityData() *types.CommonEntityData {
    tcm3Tim.EntityData.YFilter = tcm3Tim.YFilter
    tcm3Tim.EntityData.YangName = "tcm3-tim"
    tcm3Tim.EntityData.BundleName = "cisco_ios_xr"
    tcm3Tim.EntityData.ParentYangName = "alarm"
    tcm3Tim.EntityData.SegmentPath = "tcm3-tim"
    tcm3Tim.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + tcm3Tim.EntityData.SegmentPath
    tcm3Tim.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcm3Tim.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcm3Tim.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcm3Tim.EntityData.Children = types.NewOrderedMap()
    tcm3Tim.EntityData.Leafs = types.NewOrderedMap()
    tcm3Tim.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", tcm3Tim.ReportingEnabled})
    tcm3Tim.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", tcm3Tim.IsDetected})
    tcm3Tim.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", tcm3Tim.IsAsserted})
    tcm3Tim.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", tcm3Tim.Counter})

    tcm3Tim.EntityData.YListKeys = []string {}

    return &(tcm3Tim.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Tcm3SfBer
// TCM3 SF BER alarm
type Odu_Controllers_Controller_Info_Alarm_Tcm3SfBer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (tcm3SfBer *Odu_Controllers_Controller_Info_Alarm_Tcm3SfBer) GetEntityData() *types.CommonEntityData {
    tcm3SfBer.EntityData.YFilter = tcm3SfBer.YFilter
    tcm3SfBer.EntityData.YangName = "tcm3-sf-ber"
    tcm3SfBer.EntityData.BundleName = "cisco_ios_xr"
    tcm3SfBer.EntityData.ParentYangName = "alarm"
    tcm3SfBer.EntityData.SegmentPath = "tcm3-sf-ber"
    tcm3SfBer.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + tcm3SfBer.EntityData.SegmentPath
    tcm3SfBer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcm3SfBer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcm3SfBer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcm3SfBer.EntityData.Children = types.NewOrderedMap()
    tcm3SfBer.EntityData.Leafs = types.NewOrderedMap()
    tcm3SfBer.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", tcm3SfBer.ReportingEnabled})
    tcm3SfBer.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", tcm3SfBer.IsDetected})
    tcm3SfBer.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", tcm3SfBer.IsAsserted})
    tcm3SfBer.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", tcm3SfBer.Counter})

    tcm3SfBer.EntityData.YListKeys = []string {}

    return &(tcm3SfBer.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Tcm3SdBer
// TCM3 SD BER alarm
type Odu_Controllers_Controller_Info_Alarm_Tcm3SdBer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (tcm3SdBer *Odu_Controllers_Controller_Info_Alarm_Tcm3SdBer) GetEntityData() *types.CommonEntityData {
    tcm3SdBer.EntityData.YFilter = tcm3SdBer.YFilter
    tcm3SdBer.EntityData.YangName = "tcm3-sd-ber"
    tcm3SdBer.EntityData.BundleName = "cisco_ios_xr"
    tcm3SdBer.EntityData.ParentYangName = "alarm"
    tcm3SdBer.EntityData.SegmentPath = "tcm3-sd-ber"
    tcm3SdBer.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + tcm3SdBer.EntityData.SegmentPath
    tcm3SdBer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcm3SdBer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcm3SdBer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcm3SdBer.EntityData.Children = types.NewOrderedMap()
    tcm3SdBer.EntityData.Leafs = types.NewOrderedMap()
    tcm3SdBer.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", tcm3SdBer.ReportingEnabled})
    tcm3SdBer.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", tcm3SdBer.IsDetected})
    tcm3SdBer.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", tcm3SdBer.IsAsserted})
    tcm3SdBer.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", tcm3SdBer.Counter})

    tcm3SdBer.EntityData.YListKeys = []string {}

    return &(tcm3SdBer.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Tcm4Ais
// TCM4 Alarm Indication Signal
type Odu_Controllers_Controller_Info_Alarm_Tcm4Ais struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (tcm4Ais *Odu_Controllers_Controller_Info_Alarm_Tcm4Ais) GetEntityData() *types.CommonEntityData {
    tcm4Ais.EntityData.YFilter = tcm4Ais.YFilter
    tcm4Ais.EntityData.YangName = "tcm4-ais"
    tcm4Ais.EntityData.BundleName = "cisco_ios_xr"
    tcm4Ais.EntityData.ParentYangName = "alarm"
    tcm4Ais.EntityData.SegmentPath = "tcm4-ais"
    tcm4Ais.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + tcm4Ais.EntityData.SegmentPath
    tcm4Ais.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcm4Ais.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcm4Ais.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcm4Ais.EntityData.Children = types.NewOrderedMap()
    tcm4Ais.EntityData.Leafs = types.NewOrderedMap()
    tcm4Ais.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", tcm4Ais.ReportingEnabled})
    tcm4Ais.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", tcm4Ais.IsDetected})
    tcm4Ais.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", tcm4Ais.IsAsserted})
    tcm4Ais.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", tcm4Ais.Counter})

    tcm4Ais.EntityData.YListKeys = []string {}

    return &(tcm4Ais.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Tcm4Ltc
// TCM4 Loss of Tandem connection Monitoring
type Odu_Controllers_Controller_Info_Alarm_Tcm4Ltc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (tcm4Ltc *Odu_Controllers_Controller_Info_Alarm_Tcm4Ltc) GetEntityData() *types.CommonEntityData {
    tcm4Ltc.EntityData.YFilter = tcm4Ltc.YFilter
    tcm4Ltc.EntityData.YangName = "tcm4-ltc"
    tcm4Ltc.EntityData.BundleName = "cisco_ios_xr"
    tcm4Ltc.EntityData.ParentYangName = "alarm"
    tcm4Ltc.EntityData.SegmentPath = "tcm4-ltc"
    tcm4Ltc.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + tcm4Ltc.EntityData.SegmentPath
    tcm4Ltc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcm4Ltc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcm4Ltc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcm4Ltc.EntityData.Children = types.NewOrderedMap()
    tcm4Ltc.EntityData.Leafs = types.NewOrderedMap()
    tcm4Ltc.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", tcm4Ltc.ReportingEnabled})
    tcm4Ltc.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", tcm4Ltc.IsDetected})
    tcm4Ltc.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", tcm4Ltc.IsAsserted})
    tcm4Ltc.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", tcm4Ltc.Counter})

    tcm4Ltc.EntityData.YListKeys = []string {}

    return &(tcm4Ltc.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Tcm4Oci
// TCM4 Open Connection Indiction
type Odu_Controllers_Controller_Info_Alarm_Tcm4Oci struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (tcm4Oci *Odu_Controllers_Controller_Info_Alarm_Tcm4Oci) GetEntityData() *types.CommonEntityData {
    tcm4Oci.EntityData.YFilter = tcm4Oci.YFilter
    tcm4Oci.EntityData.YangName = "tcm4-oci"
    tcm4Oci.EntityData.BundleName = "cisco_ios_xr"
    tcm4Oci.EntityData.ParentYangName = "alarm"
    tcm4Oci.EntityData.SegmentPath = "tcm4-oci"
    tcm4Oci.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + tcm4Oci.EntityData.SegmentPath
    tcm4Oci.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcm4Oci.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcm4Oci.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcm4Oci.EntityData.Children = types.NewOrderedMap()
    tcm4Oci.EntityData.Leafs = types.NewOrderedMap()
    tcm4Oci.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", tcm4Oci.ReportingEnabled})
    tcm4Oci.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", tcm4Oci.IsDetected})
    tcm4Oci.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", tcm4Oci.IsAsserted})
    tcm4Oci.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", tcm4Oci.Counter})

    tcm4Oci.EntityData.YListKeys = []string {}

    return &(tcm4Oci.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Tcm4Lck
// TCM4  Upstream Connection Locked
type Odu_Controllers_Controller_Info_Alarm_Tcm4Lck struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (tcm4Lck *Odu_Controllers_Controller_Info_Alarm_Tcm4Lck) GetEntityData() *types.CommonEntityData {
    tcm4Lck.EntityData.YFilter = tcm4Lck.YFilter
    tcm4Lck.EntityData.YangName = "tcm4-lck"
    tcm4Lck.EntityData.BundleName = "cisco_ios_xr"
    tcm4Lck.EntityData.ParentYangName = "alarm"
    tcm4Lck.EntityData.SegmentPath = "tcm4-lck"
    tcm4Lck.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + tcm4Lck.EntityData.SegmentPath
    tcm4Lck.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcm4Lck.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcm4Lck.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcm4Lck.EntityData.Children = types.NewOrderedMap()
    tcm4Lck.EntityData.Leafs = types.NewOrderedMap()
    tcm4Lck.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", tcm4Lck.ReportingEnabled})
    tcm4Lck.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", tcm4Lck.IsDetected})
    tcm4Lck.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", tcm4Lck.IsAsserted})
    tcm4Lck.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", tcm4Lck.Counter})

    tcm4Lck.EntityData.YListKeys = []string {}

    return &(tcm4Lck.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Tcm4Iae
// TCM4 Incoming Alignment Error
type Odu_Controllers_Controller_Info_Alarm_Tcm4Iae struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (tcm4Iae *Odu_Controllers_Controller_Info_Alarm_Tcm4Iae) GetEntityData() *types.CommonEntityData {
    tcm4Iae.EntityData.YFilter = tcm4Iae.YFilter
    tcm4Iae.EntityData.YangName = "tcm4-iae"
    tcm4Iae.EntityData.BundleName = "cisco_ios_xr"
    tcm4Iae.EntityData.ParentYangName = "alarm"
    tcm4Iae.EntityData.SegmentPath = "tcm4-iae"
    tcm4Iae.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + tcm4Iae.EntityData.SegmentPath
    tcm4Iae.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcm4Iae.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcm4Iae.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcm4Iae.EntityData.Children = types.NewOrderedMap()
    tcm4Iae.EntityData.Leafs = types.NewOrderedMap()
    tcm4Iae.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", tcm4Iae.ReportingEnabled})
    tcm4Iae.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", tcm4Iae.IsDetected})
    tcm4Iae.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", tcm4Iae.IsAsserted})
    tcm4Iae.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", tcm4Iae.Counter})

    tcm4Iae.EntityData.YListKeys = []string {}

    return &(tcm4Iae.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Tcm4Biae
// TCM4 Backward Incoming Alignment Error
type Odu_Controllers_Controller_Info_Alarm_Tcm4Biae struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (tcm4Biae *Odu_Controllers_Controller_Info_Alarm_Tcm4Biae) GetEntityData() *types.CommonEntityData {
    tcm4Biae.EntityData.YFilter = tcm4Biae.YFilter
    tcm4Biae.EntityData.YangName = "tcm4-biae"
    tcm4Biae.EntityData.BundleName = "cisco_ios_xr"
    tcm4Biae.EntityData.ParentYangName = "alarm"
    tcm4Biae.EntityData.SegmentPath = "tcm4-biae"
    tcm4Biae.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + tcm4Biae.EntityData.SegmentPath
    tcm4Biae.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcm4Biae.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcm4Biae.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcm4Biae.EntityData.Children = types.NewOrderedMap()
    tcm4Biae.EntityData.Leafs = types.NewOrderedMap()
    tcm4Biae.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", tcm4Biae.ReportingEnabled})
    tcm4Biae.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", tcm4Biae.IsDetected})
    tcm4Biae.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", tcm4Biae.IsAsserted})
    tcm4Biae.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", tcm4Biae.Counter})

    tcm4Biae.EntityData.YListKeys = []string {}

    return &(tcm4Biae.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Tcm4Bdi
// TCM4 Backward Defect Monitoring
type Odu_Controllers_Controller_Info_Alarm_Tcm4Bdi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (tcm4Bdi *Odu_Controllers_Controller_Info_Alarm_Tcm4Bdi) GetEntityData() *types.CommonEntityData {
    tcm4Bdi.EntityData.YFilter = tcm4Bdi.YFilter
    tcm4Bdi.EntityData.YangName = "tcm4-bdi"
    tcm4Bdi.EntityData.BundleName = "cisco_ios_xr"
    tcm4Bdi.EntityData.ParentYangName = "alarm"
    tcm4Bdi.EntityData.SegmentPath = "tcm4-bdi"
    tcm4Bdi.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + tcm4Bdi.EntityData.SegmentPath
    tcm4Bdi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcm4Bdi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcm4Bdi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcm4Bdi.EntityData.Children = types.NewOrderedMap()
    tcm4Bdi.EntityData.Leafs = types.NewOrderedMap()
    tcm4Bdi.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", tcm4Bdi.ReportingEnabled})
    tcm4Bdi.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", tcm4Bdi.IsDetected})
    tcm4Bdi.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", tcm4Bdi.IsAsserted})
    tcm4Bdi.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", tcm4Bdi.Counter})

    tcm4Bdi.EntityData.YListKeys = []string {}

    return &(tcm4Bdi.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Tcm4Tim
// TCM4 Trail Trace Identifier Mismatch 
type Odu_Controllers_Controller_Info_Alarm_Tcm4Tim struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (tcm4Tim *Odu_Controllers_Controller_Info_Alarm_Tcm4Tim) GetEntityData() *types.CommonEntityData {
    tcm4Tim.EntityData.YFilter = tcm4Tim.YFilter
    tcm4Tim.EntityData.YangName = "tcm4-tim"
    tcm4Tim.EntityData.BundleName = "cisco_ios_xr"
    tcm4Tim.EntityData.ParentYangName = "alarm"
    tcm4Tim.EntityData.SegmentPath = "tcm4-tim"
    tcm4Tim.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + tcm4Tim.EntityData.SegmentPath
    tcm4Tim.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcm4Tim.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcm4Tim.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcm4Tim.EntityData.Children = types.NewOrderedMap()
    tcm4Tim.EntityData.Leafs = types.NewOrderedMap()
    tcm4Tim.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", tcm4Tim.ReportingEnabled})
    tcm4Tim.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", tcm4Tim.IsDetected})
    tcm4Tim.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", tcm4Tim.IsAsserted})
    tcm4Tim.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", tcm4Tim.Counter})

    tcm4Tim.EntityData.YListKeys = []string {}

    return &(tcm4Tim.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Tcm4SfBer
// TCM4 SF BER alarm
type Odu_Controllers_Controller_Info_Alarm_Tcm4SfBer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (tcm4SfBer *Odu_Controllers_Controller_Info_Alarm_Tcm4SfBer) GetEntityData() *types.CommonEntityData {
    tcm4SfBer.EntityData.YFilter = tcm4SfBer.YFilter
    tcm4SfBer.EntityData.YangName = "tcm4-sf-ber"
    tcm4SfBer.EntityData.BundleName = "cisco_ios_xr"
    tcm4SfBer.EntityData.ParentYangName = "alarm"
    tcm4SfBer.EntityData.SegmentPath = "tcm4-sf-ber"
    tcm4SfBer.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + tcm4SfBer.EntityData.SegmentPath
    tcm4SfBer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcm4SfBer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcm4SfBer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcm4SfBer.EntityData.Children = types.NewOrderedMap()
    tcm4SfBer.EntityData.Leafs = types.NewOrderedMap()
    tcm4SfBer.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", tcm4SfBer.ReportingEnabled})
    tcm4SfBer.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", tcm4SfBer.IsDetected})
    tcm4SfBer.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", tcm4SfBer.IsAsserted})
    tcm4SfBer.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", tcm4SfBer.Counter})

    tcm4SfBer.EntityData.YListKeys = []string {}

    return &(tcm4SfBer.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Tcm4SdBer
// TCM4 SD BER alarm
type Odu_Controllers_Controller_Info_Alarm_Tcm4SdBer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (tcm4SdBer *Odu_Controllers_Controller_Info_Alarm_Tcm4SdBer) GetEntityData() *types.CommonEntityData {
    tcm4SdBer.EntityData.YFilter = tcm4SdBer.YFilter
    tcm4SdBer.EntityData.YangName = "tcm4-sd-ber"
    tcm4SdBer.EntityData.BundleName = "cisco_ios_xr"
    tcm4SdBer.EntityData.ParentYangName = "alarm"
    tcm4SdBer.EntityData.SegmentPath = "tcm4-sd-ber"
    tcm4SdBer.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + tcm4SdBer.EntityData.SegmentPath
    tcm4SdBer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcm4SdBer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcm4SdBer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcm4SdBer.EntityData.Children = types.NewOrderedMap()
    tcm4SdBer.EntityData.Leafs = types.NewOrderedMap()
    tcm4SdBer.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", tcm4SdBer.ReportingEnabled})
    tcm4SdBer.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", tcm4SdBer.IsDetected})
    tcm4SdBer.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", tcm4SdBer.IsAsserted})
    tcm4SdBer.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", tcm4SdBer.Counter})

    tcm4SdBer.EntityData.YListKeys = []string {}

    return &(tcm4SdBer.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Tcm5Ais
// TCM5 Alarm Indication Signal
type Odu_Controllers_Controller_Info_Alarm_Tcm5Ais struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (tcm5Ais *Odu_Controllers_Controller_Info_Alarm_Tcm5Ais) GetEntityData() *types.CommonEntityData {
    tcm5Ais.EntityData.YFilter = tcm5Ais.YFilter
    tcm5Ais.EntityData.YangName = "tcm5-ais"
    tcm5Ais.EntityData.BundleName = "cisco_ios_xr"
    tcm5Ais.EntityData.ParentYangName = "alarm"
    tcm5Ais.EntityData.SegmentPath = "tcm5-ais"
    tcm5Ais.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + tcm5Ais.EntityData.SegmentPath
    tcm5Ais.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcm5Ais.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcm5Ais.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcm5Ais.EntityData.Children = types.NewOrderedMap()
    tcm5Ais.EntityData.Leafs = types.NewOrderedMap()
    tcm5Ais.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", tcm5Ais.ReportingEnabled})
    tcm5Ais.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", tcm5Ais.IsDetected})
    tcm5Ais.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", tcm5Ais.IsAsserted})
    tcm5Ais.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", tcm5Ais.Counter})

    tcm5Ais.EntityData.YListKeys = []string {}

    return &(tcm5Ais.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Tcm5Ltc
// TCM5 Loss of Tandem connection Monitoring
type Odu_Controllers_Controller_Info_Alarm_Tcm5Ltc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (tcm5Ltc *Odu_Controllers_Controller_Info_Alarm_Tcm5Ltc) GetEntityData() *types.CommonEntityData {
    tcm5Ltc.EntityData.YFilter = tcm5Ltc.YFilter
    tcm5Ltc.EntityData.YangName = "tcm5-ltc"
    tcm5Ltc.EntityData.BundleName = "cisco_ios_xr"
    tcm5Ltc.EntityData.ParentYangName = "alarm"
    tcm5Ltc.EntityData.SegmentPath = "tcm5-ltc"
    tcm5Ltc.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + tcm5Ltc.EntityData.SegmentPath
    tcm5Ltc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcm5Ltc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcm5Ltc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcm5Ltc.EntityData.Children = types.NewOrderedMap()
    tcm5Ltc.EntityData.Leafs = types.NewOrderedMap()
    tcm5Ltc.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", tcm5Ltc.ReportingEnabled})
    tcm5Ltc.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", tcm5Ltc.IsDetected})
    tcm5Ltc.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", tcm5Ltc.IsAsserted})
    tcm5Ltc.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", tcm5Ltc.Counter})

    tcm5Ltc.EntityData.YListKeys = []string {}

    return &(tcm5Ltc.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Tcm5Oci
// TCM5 Open Connection Indiction
type Odu_Controllers_Controller_Info_Alarm_Tcm5Oci struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (tcm5Oci *Odu_Controllers_Controller_Info_Alarm_Tcm5Oci) GetEntityData() *types.CommonEntityData {
    tcm5Oci.EntityData.YFilter = tcm5Oci.YFilter
    tcm5Oci.EntityData.YangName = "tcm5-oci"
    tcm5Oci.EntityData.BundleName = "cisco_ios_xr"
    tcm5Oci.EntityData.ParentYangName = "alarm"
    tcm5Oci.EntityData.SegmentPath = "tcm5-oci"
    tcm5Oci.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + tcm5Oci.EntityData.SegmentPath
    tcm5Oci.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcm5Oci.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcm5Oci.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcm5Oci.EntityData.Children = types.NewOrderedMap()
    tcm5Oci.EntityData.Leafs = types.NewOrderedMap()
    tcm5Oci.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", tcm5Oci.ReportingEnabled})
    tcm5Oci.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", tcm5Oci.IsDetected})
    tcm5Oci.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", tcm5Oci.IsAsserted})
    tcm5Oci.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", tcm5Oci.Counter})

    tcm5Oci.EntityData.YListKeys = []string {}

    return &(tcm5Oci.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Tcm5Lck
// TCM5  Upstream Connection Locked
type Odu_Controllers_Controller_Info_Alarm_Tcm5Lck struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (tcm5Lck *Odu_Controllers_Controller_Info_Alarm_Tcm5Lck) GetEntityData() *types.CommonEntityData {
    tcm5Lck.EntityData.YFilter = tcm5Lck.YFilter
    tcm5Lck.EntityData.YangName = "tcm5-lck"
    tcm5Lck.EntityData.BundleName = "cisco_ios_xr"
    tcm5Lck.EntityData.ParentYangName = "alarm"
    tcm5Lck.EntityData.SegmentPath = "tcm5-lck"
    tcm5Lck.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + tcm5Lck.EntityData.SegmentPath
    tcm5Lck.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcm5Lck.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcm5Lck.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcm5Lck.EntityData.Children = types.NewOrderedMap()
    tcm5Lck.EntityData.Leafs = types.NewOrderedMap()
    tcm5Lck.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", tcm5Lck.ReportingEnabled})
    tcm5Lck.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", tcm5Lck.IsDetected})
    tcm5Lck.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", tcm5Lck.IsAsserted})
    tcm5Lck.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", tcm5Lck.Counter})

    tcm5Lck.EntityData.YListKeys = []string {}

    return &(tcm5Lck.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Tcm5Iae
// TCM5 Incoming Alignment Error
type Odu_Controllers_Controller_Info_Alarm_Tcm5Iae struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (tcm5Iae *Odu_Controllers_Controller_Info_Alarm_Tcm5Iae) GetEntityData() *types.CommonEntityData {
    tcm5Iae.EntityData.YFilter = tcm5Iae.YFilter
    tcm5Iae.EntityData.YangName = "tcm5-iae"
    tcm5Iae.EntityData.BundleName = "cisco_ios_xr"
    tcm5Iae.EntityData.ParentYangName = "alarm"
    tcm5Iae.EntityData.SegmentPath = "tcm5-iae"
    tcm5Iae.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + tcm5Iae.EntityData.SegmentPath
    tcm5Iae.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcm5Iae.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcm5Iae.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcm5Iae.EntityData.Children = types.NewOrderedMap()
    tcm5Iae.EntityData.Leafs = types.NewOrderedMap()
    tcm5Iae.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", tcm5Iae.ReportingEnabled})
    tcm5Iae.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", tcm5Iae.IsDetected})
    tcm5Iae.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", tcm5Iae.IsAsserted})
    tcm5Iae.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", tcm5Iae.Counter})

    tcm5Iae.EntityData.YListKeys = []string {}

    return &(tcm5Iae.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Tcm5Biae
// TCM5 Backward Incoming Alignment Error
type Odu_Controllers_Controller_Info_Alarm_Tcm5Biae struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (tcm5Biae *Odu_Controllers_Controller_Info_Alarm_Tcm5Biae) GetEntityData() *types.CommonEntityData {
    tcm5Biae.EntityData.YFilter = tcm5Biae.YFilter
    tcm5Biae.EntityData.YangName = "tcm5-biae"
    tcm5Biae.EntityData.BundleName = "cisco_ios_xr"
    tcm5Biae.EntityData.ParentYangName = "alarm"
    tcm5Biae.EntityData.SegmentPath = "tcm5-biae"
    tcm5Biae.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + tcm5Biae.EntityData.SegmentPath
    tcm5Biae.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcm5Biae.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcm5Biae.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcm5Biae.EntityData.Children = types.NewOrderedMap()
    tcm5Biae.EntityData.Leafs = types.NewOrderedMap()
    tcm5Biae.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", tcm5Biae.ReportingEnabled})
    tcm5Biae.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", tcm5Biae.IsDetected})
    tcm5Biae.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", tcm5Biae.IsAsserted})
    tcm5Biae.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", tcm5Biae.Counter})

    tcm5Biae.EntityData.YListKeys = []string {}

    return &(tcm5Biae.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Tcm5Bdi
// TCM5 Backward Defect Monitoring
type Odu_Controllers_Controller_Info_Alarm_Tcm5Bdi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (tcm5Bdi *Odu_Controllers_Controller_Info_Alarm_Tcm5Bdi) GetEntityData() *types.CommonEntityData {
    tcm5Bdi.EntityData.YFilter = tcm5Bdi.YFilter
    tcm5Bdi.EntityData.YangName = "tcm5-bdi"
    tcm5Bdi.EntityData.BundleName = "cisco_ios_xr"
    tcm5Bdi.EntityData.ParentYangName = "alarm"
    tcm5Bdi.EntityData.SegmentPath = "tcm5-bdi"
    tcm5Bdi.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + tcm5Bdi.EntityData.SegmentPath
    tcm5Bdi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcm5Bdi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcm5Bdi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcm5Bdi.EntityData.Children = types.NewOrderedMap()
    tcm5Bdi.EntityData.Leafs = types.NewOrderedMap()
    tcm5Bdi.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", tcm5Bdi.ReportingEnabled})
    tcm5Bdi.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", tcm5Bdi.IsDetected})
    tcm5Bdi.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", tcm5Bdi.IsAsserted})
    tcm5Bdi.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", tcm5Bdi.Counter})

    tcm5Bdi.EntityData.YListKeys = []string {}

    return &(tcm5Bdi.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Tcm5Tim
// TCM5 Trail Trace Identifier Mismatch 
type Odu_Controllers_Controller_Info_Alarm_Tcm5Tim struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (tcm5Tim *Odu_Controllers_Controller_Info_Alarm_Tcm5Tim) GetEntityData() *types.CommonEntityData {
    tcm5Tim.EntityData.YFilter = tcm5Tim.YFilter
    tcm5Tim.EntityData.YangName = "tcm5-tim"
    tcm5Tim.EntityData.BundleName = "cisco_ios_xr"
    tcm5Tim.EntityData.ParentYangName = "alarm"
    tcm5Tim.EntityData.SegmentPath = "tcm5-tim"
    tcm5Tim.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + tcm5Tim.EntityData.SegmentPath
    tcm5Tim.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcm5Tim.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcm5Tim.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcm5Tim.EntityData.Children = types.NewOrderedMap()
    tcm5Tim.EntityData.Leafs = types.NewOrderedMap()
    tcm5Tim.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", tcm5Tim.ReportingEnabled})
    tcm5Tim.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", tcm5Tim.IsDetected})
    tcm5Tim.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", tcm5Tim.IsAsserted})
    tcm5Tim.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", tcm5Tim.Counter})

    tcm5Tim.EntityData.YListKeys = []string {}

    return &(tcm5Tim.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Tcm5SfBer
// TCM5 SF BER alarm
type Odu_Controllers_Controller_Info_Alarm_Tcm5SfBer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (tcm5SfBer *Odu_Controllers_Controller_Info_Alarm_Tcm5SfBer) GetEntityData() *types.CommonEntityData {
    tcm5SfBer.EntityData.YFilter = tcm5SfBer.YFilter
    tcm5SfBer.EntityData.YangName = "tcm5-sf-ber"
    tcm5SfBer.EntityData.BundleName = "cisco_ios_xr"
    tcm5SfBer.EntityData.ParentYangName = "alarm"
    tcm5SfBer.EntityData.SegmentPath = "tcm5-sf-ber"
    tcm5SfBer.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + tcm5SfBer.EntityData.SegmentPath
    tcm5SfBer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcm5SfBer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcm5SfBer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcm5SfBer.EntityData.Children = types.NewOrderedMap()
    tcm5SfBer.EntityData.Leafs = types.NewOrderedMap()
    tcm5SfBer.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", tcm5SfBer.ReportingEnabled})
    tcm5SfBer.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", tcm5SfBer.IsDetected})
    tcm5SfBer.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", tcm5SfBer.IsAsserted})
    tcm5SfBer.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", tcm5SfBer.Counter})

    tcm5SfBer.EntityData.YListKeys = []string {}

    return &(tcm5SfBer.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Tcm5SdBer
// TCM5 SD BER alarm
type Odu_Controllers_Controller_Info_Alarm_Tcm5SdBer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (tcm5SdBer *Odu_Controllers_Controller_Info_Alarm_Tcm5SdBer) GetEntityData() *types.CommonEntityData {
    tcm5SdBer.EntityData.YFilter = tcm5SdBer.YFilter
    tcm5SdBer.EntityData.YangName = "tcm5-sd-ber"
    tcm5SdBer.EntityData.BundleName = "cisco_ios_xr"
    tcm5SdBer.EntityData.ParentYangName = "alarm"
    tcm5SdBer.EntityData.SegmentPath = "tcm5-sd-ber"
    tcm5SdBer.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + tcm5SdBer.EntityData.SegmentPath
    tcm5SdBer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcm5SdBer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcm5SdBer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcm5SdBer.EntityData.Children = types.NewOrderedMap()
    tcm5SdBer.EntityData.Leafs = types.NewOrderedMap()
    tcm5SdBer.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", tcm5SdBer.ReportingEnabled})
    tcm5SdBer.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", tcm5SdBer.IsDetected})
    tcm5SdBer.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", tcm5SdBer.IsAsserted})
    tcm5SdBer.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", tcm5SdBer.Counter})

    tcm5SdBer.EntityData.YListKeys = []string {}

    return &(tcm5SdBer.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Tcm6Ais
// TCM6 Alarm Indication Signal
type Odu_Controllers_Controller_Info_Alarm_Tcm6Ais struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (tcm6Ais *Odu_Controllers_Controller_Info_Alarm_Tcm6Ais) GetEntityData() *types.CommonEntityData {
    tcm6Ais.EntityData.YFilter = tcm6Ais.YFilter
    tcm6Ais.EntityData.YangName = "tcm6-ais"
    tcm6Ais.EntityData.BundleName = "cisco_ios_xr"
    tcm6Ais.EntityData.ParentYangName = "alarm"
    tcm6Ais.EntityData.SegmentPath = "tcm6-ais"
    tcm6Ais.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + tcm6Ais.EntityData.SegmentPath
    tcm6Ais.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcm6Ais.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcm6Ais.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcm6Ais.EntityData.Children = types.NewOrderedMap()
    tcm6Ais.EntityData.Leafs = types.NewOrderedMap()
    tcm6Ais.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", tcm6Ais.ReportingEnabled})
    tcm6Ais.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", tcm6Ais.IsDetected})
    tcm6Ais.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", tcm6Ais.IsAsserted})
    tcm6Ais.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", tcm6Ais.Counter})

    tcm6Ais.EntityData.YListKeys = []string {}

    return &(tcm6Ais.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Tcm6Ltc
// TCM6 Loss of Tandem connection Monitoring
type Odu_Controllers_Controller_Info_Alarm_Tcm6Ltc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (tcm6Ltc *Odu_Controllers_Controller_Info_Alarm_Tcm6Ltc) GetEntityData() *types.CommonEntityData {
    tcm6Ltc.EntityData.YFilter = tcm6Ltc.YFilter
    tcm6Ltc.EntityData.YangName = "tcm6-ltc"
    tcm6Ltc.EntityData.BundleName = "cisco_ios_xr"
    tcm6Ltc.EntityData.ParentYangName = "alarm"
    tcm6Ltc.EntityData.SegmentPath = "tcm6-ltc"
    tcm6Ltc.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + tcm6Ltc.EntityData.SegmentPath
    tcm6Ltc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcm6Ltc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcm6Ltc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcm6Ltc.EntityData.Children = types.NewOrderedMap()
    tcm6Ltc.EntityData.Leafs = types.NewOrderedMap()
    tcm6Ltc.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", tcm6Ltc.ReportingEnabled})
    tcm6Ltc.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", tcm6Ltc.IsDetected})
    tcm6Ltc.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", tcm6Ltc.IsAsserted})
    tcm6Ltc.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", tcm6Ltc.Counter})

    tcm6Ltc.EntityData.YListKeys = []string {}

    return &(tcm6Ltc.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Tcm6Oci
// TCM6 Open Connection Indiction
type Odu_Controllers_Controller_Info_Alarm_Tcm6Oci struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (tcm6Oci *Odu_Controllers_Controller_Info_Alarm_Tcm6Oci) GetEntityData() *types.CommonEntityData {
    tcm6Oci.EntityData.YFilter = tcm6Oci.YFilter
    tcm6Oci.EntityData.YangName = "tcm6-oci"
    tcm6Oci.EntityData.BundleName = "cisco_ios_xr"
    tcm6Oci.EntityData.ParentYangName = "alarm"
    tcm6Oci.EntityData.SegmentPath = "tcm6-oci"
    tcm6Oci.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + tcm6Oci.EntityData.SegmentPath
    tcm6Oci.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcm6Oci.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcm6Oci.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcm6Oci.EntityData.Children = types.NewOrderedMap()
    tcm6Oci.EntityData.Leafs = types.NewOrderedMap()
    tcm6Oci.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", tcm6Oci.ReportingEnabled})
    tcm6Oci.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", tcm6Oci.IsDetected})
    tcm6Oci.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", tcm6Oci.IsAsserted})
    tcm6Oci.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", tcm6Oci.Counter})

    tcm6Oci.EntityData.YListKeys = []string {}

    return &(tcm6Oci.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Tcm6Lck
// TCM6  Upstream Connection Locked
type Odu_Controllers_Controller_Info_Alarm_Tcm6Lck struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (tcm6Lck *Odu_Controllers_Controller_Info_Alarm_Tcm6Lck) GetEntityData() *types.CommonEntityData {
    tcm6Lck.EntityData.YFilter = tcm6Lck.YFilter
    tcm6Lck.EntityData.YangName = "tcm6-lck"
    tcm6Lck.EntityData.BundleName = "cisco_ios_xr"
    tcm6Lck.EntityData.ParentYangName = "alarm"
    tcm6Lck.EntityData.SegmentPath = "tcm6-lck"
    tcm6Lck.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + tcm6Lck.EntityData.SegmentPath
    tcm6Lck.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcm6Lck.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcm6Lck.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcm6Lck.EntityData.Children = types.NewOrderedMap()
    tcm6Lck.EntityData.Leafs = types.NewOrderedMap()
    tcm6Lck.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", tcm6Lck.ReportingEnabled})
    tcm6Lck.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", tcm6Lck.IsDetected})
    tcm6Lck.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", tcm6Lck.IsAsserted})
    tcm6Lck.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", tcm6Lck.Counter})

    tcm6Lck.EntityData.YListKeys = []string {}

    return &(tcm6Lck.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Tcm6Iae
// TCM6 Incoming Alignment Error
type Odu_Controllers_Controller_Info_Alarm_Tcm6Iae struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (tcm6Iae *Odu_Controllers_Controller_Info_Alarm_Tcm6Iae) GetEntityData() *types.CommonEntityData {
    tcm6Iae.EntityData.YFilter = tcm6Iae.YFilter
    tcm6Iae.EntityData.YangName = "tcm6-iae"
    tcm6Iae.EntityData.BundleName = "cisco_ios_xr"
    tcm6Iae.EntityData.ParentYangName = "alarm"
    tcm6Iae.EntityData.SegmentPath = "tcm6-iae"
    tcm6Iae.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + tcm6Iae.EntityData.SegmentPath
    tcm6Iae.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcm6Iae.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcm6Iae.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcm6Iae.EntityData.Children = types.NewOrderedMap()
    tcm6Iae.EntityData.Leafs = types.NewOrderedMap()
    tcm6Iae.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", tcm6Iae.ReportingEnabled})
    tcm6Iae.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", tcm6Iae.IsDetected})
    tcm6Iae.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", tcm6Iae.IsAsserted})
    tcm6Iae.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", tcm6Iae.Counter})

    tcm6Iae.EntityData.YListKeys = []string {}

    return &(tcm6Iae.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Tcm6Biae
// TCM6 Backward Incoming Alignment Error
type Odu_Controllers_Controller_Info_Alarm_Tcm6Biae struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (tcm6Biae *Odu_Controllers_Controller_Info_Alarm_Tcm6Biae) GetEntityData() *types.CommonEntityData {
    tcm6Biae.EntityData.YFilter = tcm6Biae.YFilter
    tcm6Biae.EntityData.YangName = "tcm6-biae"
    tcm6Biae.EntityData.BundleName = "cisco_ios_xr"
    tcm6Biae.EntityData.ParentYangName = "alarm"
    tcm6Biae.EntityData.SegmentPath = "tcm6-biae"
    tcm6Biae.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + tcm6Biae.EntityData.SegmentPath
    tcm6Biae.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcm6Biae.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcm6Biae.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcm6Biae.EntityData.Children = types.NewOrderedMap()
    tcm6Biae.EntityData.Leafs = types.NewOrderedMap()
    tcm6Biae.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", tcm6Biae.ReportingEnabled})
    tcm6Biae.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", tcm6Biae.IsDetected})
    tcm6Biae.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", tcm6Biae.IsAsserted})
    tcm6Biae.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", tcm6Biae.Counter})

    tcm6Biae.EntityData.YListKeys = []string {}

    return &(tcm6Biae.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Tcm6Bdi
// TCM6 Backward Defect Monitoring
type Odu_Controllers_Controller_Info_Alarm_Tcm6Bdi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (tcm6Bdi *Odu_Controllers_Controller_Info_Alarm_Tcm6Bdi) GetEntityData() *types.CommonEntityData {
    tcm6Bdi.EntityData.YFilter = tcm6Bdi.YFilter
    tcm6Bdi.EntityData.YangName = "tcm6-bdi"
    tcm6Bdi.EntityData.BundleName = "cisco_ios_xr"
    tcm6Bdi.EntityData.ParentYangName = "alarm"
    tcm6Bdi.EntityData.SegmentPath = "tcm6-bdi"
    tcm6Bdi.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + tcm6Bdi.EntityData.SegmentPath
    tcm6Bdi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcm6Bdi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcm6Bdi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcm6Bdi.EntityData.Children = types.NewOrderedMap()
    tcm6Bdi.EntityData.Leafs = types.NewOrderedMap()
    tcm6Bdi.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", tcm6Bdi.ReportingEnabled})
    tcm6Bdi.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", tcm6Bdi.IsDetected})
    tcm6Bdi.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", tcm6Bdi.IsAsserted})
    tcm6Bdi.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", tcm6Bdi.Counter})

    tcm6Bdi.EntityData.YListKeys = []string {}

    return &(tcm6Bdi.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Tcm6Tim
// TCM6 Trail Trace Identifier Mismatch 
type Odu_Controllers_Controller_Info_Alarm_Tcm6Tim struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (tcm6Tim *Odu_Controllers_Controller_Info_Alarm_Tcm6Tim) GetEntityData() *types.CommonEntityData {
    tcm6Tim.EntityData.YFilter = tcm6Tim.YFilter
    tcm6Tim.EntityData.YangName = "tcm6-tim"
    tcm6Tim.EntityData.BundleName = "cisco_ios_xr"
    tcm6Tim.EntityData.ParentYangName = "alarm"
    tcm6Tim.EntityData.SegmentPath = "tcm6-tim"
    tcm6Tim.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + tcm6Tim.EntityData.SegmentPath
    tcm6Tim.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcm6Tim.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcm6Tim.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcm6Tim.EntityData.Children = types.NewOrderedMap()
    tcm6Tim.EntityData.Leafs = types.NewOrderedMap()
    tcm6Tim.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", tcm6Tim.ReportingEnabled})
    tcm6Tim.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", tcm6Tim.IsDetected})
    tcm6Tim.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", tcm6Tim.IsAsserted})
    tcm6Tim.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", tcm6Tim.Counter})

    tcm6Tim.EntityData.YListKeys = []string {}

    return &(tcm6Tim.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Tcm6SfBer
// TCM6 SF BER alarm
type Odu_Controllers_Controller_Info_Alarm_Tcm6SfBer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (tcm6SfBer *Odu_Controllers_Controller_Info_Alarm_Tcm6SfBer) GetEntityData() *types.CommonEntityData {
    tcm6SfBer.EntityData.YFilter = tcm6SfBer.YFilter
    tcm6SfBer.EntityData.YangName = "tcm6-sf-ber"
    tcm6SfBer.EntityData.BundleName = "cisco_ios_xr"
    tcm6SfBer.EntityData.ParentYangName = "alarm"
    tcm6SfBer.EntityData.SegmentPath = "tcm6-sf-ber"
    tcm6SfBer.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + tcm6SfBer.EntityData.SegmentPath
    tcm6SfBer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcm6SfBer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcm6SfBer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcm6SfBer.EntityData.Children = types.NewOrderedMap()
    tcm6SfBer.EntityData.Leafs = types.NewOrderedMap()
    tcm6SfBer.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", tcm6SfBer.ReportingEnabled})
    tcm6SfBer.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", tcm6SfBer.IsDetected})
    tcm6SfBer.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", tcm6SfBer.IsAsserted})
    tcm6SfBer.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", tcm6SfBer.Counter})

    tcm6SfBer.EntityData.YListKeys = []string {}

    return &(tcm6SfBer.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_Tcm6SdBer
// TCM6 SD BER alarm
type Odu_Controllers_Controller_Info_Alarm_Tcm6SdBer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (tcm6SdBer *Odu_Controllers_Controller_Info_Alarm_Tcm6SdBer) GetEntityData() *types.CommonEntityData {
    tcm6SdBer.EntityData.YFilter = tcm6SdBer.YFilter
    tcm6SdBer.EntityData.YangName = "tcm6-sd-ber"
    tcm6SdBer.EntityData.BundleName = "cisco_ios_xr"
    tcm6SdBer.EntityData.ParentYangName = "alarm"
    tcm6SdBer.EntityData.SegmentPath = "tcm6-sd-ber"
    tcm6SdBer.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + tcm6SdBer.EntityData.SegmentPath
    tcm6SdBer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcm6SdBer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcm6SdBer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcm6SdBer.EntityData.Children = types.NewOrderedMap()
    tcm6SdBer.EntityData.Leafs = types.NewOrderedMap()
    tcm6SdBer.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", tcm6SdBer.ReportingEnabled})
    tcm6SdBer.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", tcm6SdBer.IsDetected})
    tcm6SdBer.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", tcm6SdBer.IsAsserted})
    tcm6SdBer.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", tcm6SdBer.Counter})

    tcm6SdBer.EntityData.YListKeys = []string {}

    return &(tcm6SdBer.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_GfpLfd
// Loss Of Frame Delineation
type Odu_Controllers_Controller_Info_Alarm_GfpLfd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (gfpLfd *Odu_Controllers_Controller_Info_Alarm_GfpLfd) GetEntityData() *types.CommonEntityData {
    gfpLfd.EntityData.YFilter = gfpLfd.YFilter
    gfpLfd.EntityData.YangName = "gfp-lfd"
    gfpLfd.EntityData.BundleName = "cisco_ios_xr"
    gfpLfd.EntityData.ParentYangName = "alarm"
    gfpLfd.EntityData.SegmentPath = "gfp-lfd"
    gfpLfd.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + gfpLfd.EntityData.SegmentPath
    gfpLfd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    gfpLfd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    gfpLfd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    gfpLfd.EntityData.Children = types.NewOrderedMap()
    gfpLfd.EntityData.Leafs = types.NewOrderedMap()
    gfpLfd.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", gfpLfd.ReportingEnabled})
    gfpLfd.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", gfpLfd.IsDetected})
    gfpLfd.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", gfpLfd.IsAsserted})
    gfpLfd.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", gfpLfd.Counter})

    gfpLfd.EntityData.YListKeys = []string {}

    return &(gfpLfd.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_GfpLocs
// Loss Of Client Signal
type Odu_Controllers_Controller_Info_Alarm_GfpLocs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (gfpLocs *Odu_Controllers_Controller_Info_Alarm_GfpLocs) GetEntityData() *types.CommonEntityData {
    gfpLocs.EntityData.YFilter = gfpLocs.YFilter
    gfpLocs.EntityData.YangName = "gfp-locs"
    gfpLocs.EntityData.BundleName = "cisco_ios_xr"
    gfpLocs.EntityData.ParentYangName = "alarm"
    gfpLocs.EntityData.SegmentPath = "gfp-locs"
    gfpLocs.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + gfpLocs.EntityData.SegmentPath
    gfpLocs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    gfpLocs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    gfpLocs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    gfpLocs.EntityData.Children = types.NewOrderedMap()
    gfpLocs.EntityData.Leafs = types.NewOrderedMap()
    gfpLocs.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", gfpLocs.ReportingEnabled})
    gfpLocs.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", gfpLocs.IsDetected})
    gfpLocs.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", gfpLocs.IsAsserted})
    gfpLocs.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", gfpLocs.Counter})

    gfpLocs.EntityData.YListKeys = []string {}

    return &(gfpLocs.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_GfpLoccs
// Loss Of Character Synchronization
type Odu_Controllers_Controller_Info_Alarm_GfpLoccs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (gfpLoccs *Odu_Controllers_Controller_Info_Alarm_GfpLoccs) GetEntityData() *types.CommonEntityData {
    gfpLoccs.EntityData.YFilter = gfpLoccs.YFilter
    gfpLoccs.EntityData.YangName = "gfp-loccs"
    gfpLoccs.EntityData.BundleName = "cisco_ios_xr"
    gfpLoccs.EntityData.ParentYangName = "alarm"
    gfpLoccs.EntityData.SegmentPath = "gfp-loccs"
    gfpLoccs.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + gfpLoccs.EntityData.SegmentPath
    gfpLoccs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    gfpLoccs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    gfpLoccs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    gfpLoccs.EntityData.Children = types.NewOrderedMap()
    gfpLoccs.EntityData.Leafs = types.NewOrderedMap()
    gfpLoccs.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", gfpLoccs.ReportingEnabled})
    gfpLoccs.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", gfpLoccs.IsDetected})
    gfpLoccs.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", gfpLoccs.IsAsserted})
    gfpLoccs.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", gfpLoccs.Counter})

    gfpLoccs.EntityData.YListKeys = []string {}

    return &(gfpLoccs.EntityData)
}

// Odu_Controllers_Controller_Info_Alarm_GfpUpm
// User Payload Mismatch
type Odu_Controllers_Controller_Info_Alarm_GfpUpm struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (gfpUpm *Odu_Controllers_Controller_Info_Alarm_GfpUpm) GetEntityData() *types.CommonEntityData {
    gfpUpm.EntityData.YFilter = gfpUpm.YFilter
    gfpUpm.EntityData.YangName = "gfp-upm"
    gfpUpm.EntityData.BundleName = "cisco_ios_xr"
    gfpUpm.EntityData.ParentYangName = "alarm"
    gfpUpm.EntityData.SegmentPath = "gfp-upm"
    gfpUpm.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/alarm/" + gfpUpm.EntityData.SegmentPath
    gfpUpm.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    gfpUpm.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    gfpUpm.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    gfpUpm.EntityData.Children = types.NewOrderedMap()
    gfpUpm.EntityData.Leafs = types.NewOrderedMap()
    gfpUpm.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", gfpUpm.ReportingEnabled})
    gfpUpm.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", gfpUpm.IsDetected})
    gfpUpm.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", gfpUpm.IsAsserted})
    gfpUpm.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", gfpUpm.Counter})

    gfpUpm.EntityData.YListKeys = []string {}

    return &(gfpUpm.EntityData)
}

// Odu_Controllers_Controller_Info_TeCtxData
// Label Get Data
type Odu_Controllers_Controller_Info_TeCtxData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Req Time. The type is interface{} with range: 0..4294967295.
    GmplsReqTime interface{}

    // Ctxt Type. The type is OtmOpticalRmCtxt.
    CtxtType interface{}

    // Rm Type. The type is OtmOpticalRmCtxtRm.
    RmType interface{}

    // Tunnel Information.
    TeTunnelInfo Odu_Controllers_Controller_Info_TeCtxData_TeTunnelInfo
}

func (teCtxData *Odu_Controllers_Controller_Info_TeCtxData) GetEntityData() *types.CommonEntityData {
    teCtxData.EntityData.YFilter = teCtxData.YFilter
    teCtxData.EntityData.YangName = "te-ctx-data"
    teCtxData.EntityData.BundleName = "cisco_ios_xr"
    teCtxData.EntityData.ParentYangName = "info"
    teCtxData.EntityData.SegmentPath = "te-ctx-data"
    teCtxData.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/" + teCtxData.EntityData.SegmentPath
    teCtxData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    teCtxData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    teCtxData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    teCtxData.EntityData.Children = types.NewOrderedMap()
    teCtxData.EntityData.Children.Append("te-tunnel-info", types.YChild{"TeTunnelInfo", &teCtxData.TeTunnelInfo})
    teCtxData.EntityData.Leafs = types.NewOrderedMap()
    teCtxData.EntityData.Leafs.Append("gmpls-req-time", types.YLeaf{"GmplsReqTime", teCtxData.GmplsReqTime})
    teCtxData.EntityData.Leafs.Append("ctxt-type", types.YLeaf{"CtxtType", teCtxData.CtxtType})
    teCtxData.EntityData.Leafs.Append("rm-type", types.YLeaf{"RmType", teCtxData.RmType})

    teCtxData.EntityData.YListKeys = []string {}

    return &(teCtxData.EntityData)
}

// Odu_Controllers_Controller_Info_TeCtxData_TeTunnelInfo
// Tunnel Information
type Odu_Controllers_Controller_Info_TeCtxData_TeTunnelInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // INFO TYPE. The type is OtmTeTunnelInfo.
    InfoType interface{}

    // Tunnel Id. The type is interface{} with range: 0..4294967295.
    TunnelId interface{}

    // Lbl Ctxt.
    LbCtxt Odu_Controllers_Controller_Info_TeCtxData_TeTunnelInfo_LbCtxt

    // Passive Match.
    PassiveMatch Odu_Controllers_Controller_Info_TeCtxData_TeTunnelInfo_PassiveMatch
}

func (teTunnelInfo *Odu_Controllers_Controller_Info_TeCtxData_TeTunnelInfo) GetEntityData() *types.CommonEntityData {
    teTunnelInfo.EntityData.YFilter = teTunnelInfo.YFilter
    teTunnelInfo.EntityData.YangName = "te-tunnel-info"
    teTunnelInfo.EntityData.BundleName = "cisco_ios_xr"
    teTunnelInfo.EntityData.ParentYangName = "te-ctx-data"
    teTunnelInfo.EntityData.SegmentPath = "te-tunnel-info"
    teTunnelInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/te-ctx-data/" + teTunnelInfo.EntityData.SegmentPath
    teTunnelInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    teTunnelInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    teTunnelInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    teTunnelInfo.EntityData.Children = types.NewOrderedMap()
    teTunnelInfo.EntityData.Children.Append("lb-ctxt", types.YChild{"LbCtxt", &teTunnelInfo.LbCtxt})
    teTunnelInfo.EntityData.Children.Append("passive-match", types.YChild{"PassiveMatch", &teTunnelInfo.PassiveMatch})
    teTunnelInfo.EntityData.Leafs = types.NewOrderedMap()
    teTunnelInfo.EntityData.Leafs.Append("info-type", types.YLeaf{"InfoType", teTunnelInfo.InfoType})
    teTunnelInfo.EntityData.Leafs.Append("tunnel-id", types.YLeaf{"TunnelId", teTunnelInfo.TunnelId})

    teTunnelInfo.EntityData.YListKeys = []string {}

    return &(teTunnelInfo.EntityData)
}

// Odu_Controllers_Controller_Info_TeCtxData_TeTunnelInfo_LbCtxt
// Lbl Ctxt
type Odu_Controllers_Controller_Info_TeCtxData_TeTunnelInfo_LbCtxt struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SubGroup Id. The type is interface{} with range: 0..65535.
    S2lFecSubGroupId interface{}

    // Lsp Id. The type is interface{} with range: 0..65535.
    S2lFecLspId interface{}

    // Tunnel Id. The type is interface{} with range: 0..65535.
    S2lFecTunnelId interface{}

    // Ext Tunnel Id. The type is interface{} with range: 0..4294967295.
    ExtTunnelId interface{}

    // FEC Source. The type is interface{} with range: 0..4294967295.
    FecSource interface{}

    // FEC Dest. The type is interface{} with range: 0..4294967295.
    FecDest interface{}

    // P2MP Id. The type is interface{} with range: 0..4294967295.
    S2lFecP2mpId interface{}

    // SubGroup Originator. The type is interface{} with range: 0..4294967295.
    SubGroupOriginAtor interface{}

    // Ctype. The type is OtmMplsLibC.
    FecCType interface{}
}

func (lbCtxt *Odu_Controllers_Controller_Info_TeCtxData_TeTunnelInfo_LbCtxt) GetEntityData() *types.CommonEntityData {
    lbCtxt.EntityData.YFilter = lbCtxt.YFilter
    lbCtxt.EntityData.YangName = "lb-ctxt"
    lbCtxt.EntityData.BundleName = "cisco_ios_xr"
    lbCtxt.EntityData.ParentYangName = "te-tunnel-info"
    lbCtxt.EntityData.SegmentPath = "lb-ctxt"
    lbCtxt.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/te-ctx-data/te-tunnel-info/" + lbCtxt.EntityData.SegmentPath
    lbCtxt.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lbCtxt.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lbCtxt.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lbCtxt.EntityData.Children = types.NewOrderedMap()
    lbCtxt.EntityData.Leafs = types.NewOrderedMap()
    lbCtxt.EntityData.Leafs.Append("s2l-fec-sub-group-id", types.YLeaf{"S2lFecSubGroupId", lbCtxt.S2lFecSubGroupId})
    lbCtxt.EntityData.Leafs.Append("s2l-fec-lsp-id", types.YLeaf{"S2lFecLspId", lbCtxt.S2lFecLspId})
    lbCtxt.EntityData.Leafs.Append("s2l-fec-tunnel-id", types.YLeaf{"S2lFecTunnelId", lbCtxt.S2lFecTunnelId})
    lbCtxt.EntityData.Leafs.Append("ext-tunnel-id", types.YLeaf{"ExtTunnelId", lbCtxt.ExtTunnelId})
    lbCtxt.EntityData.Leafs.Append("fec-source", types.YLeaf{"FecSource", lbCtxt.FecSource})
    lbCtxt.EntityData.Leafs.Append("fec-dest", types.YLeaf{"FecDest", lbCtxt.FecDest})
    lbCtxt.EntityData.Leafs.Append("s2l-fec-p2mp-id", types.YLeaf{"S2lFecP2mpId", lbCtxt.S2lFecP2mpId})
    lbCtxt.EntityData.Leafs.Append("sub-group-origin-ator", types.YLeaf{"SubGroupOriginAtor", lbCtxt.SubGroupOriginAtor})
    lbCtxt.EntityData.Leafs.Append("fec-c-type", types.YLeaf{"FecCType", lbCtxt.FecCType})

    lbCtxt.EntityData.YListKeys = []string {}

    return &(lbCtxt.EntityData)
}

// Odu_Controllers_Controller_Info_TeCtxData_TeTunnelInfo_PassiveMatch
// Passive Match
type Odu_Controllers_Controller_Info_TeCtxData_TeTunnelInfo_PassiveMatch struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Src TId. The type is interface{} with range: 0..65535.
    SrcTid interface{}

    // Src RId. The type is interface{} with range: 0..4294967295.
    SrcRid interface{}
}

func (passiveMatch *Odu_Controllers_Controller_Info_TeCtxData_TeTunnelInfo_PassiveMatch) GetEntityData() *types.CommonEntityData {
    passiveMatch.EntityData.YFilter = passiveMatch.YFilter
    passiveMatch.EntityData.YangName = "passive-match"
    passiveMatch.EntityData.BundleName = "cisco_ios_xr"
    passiveMatch.EntityData.ParentYangName = "te-tunnel-info"
    passiveMatch.EntityData.SegmentPath = "passive-match"
    passiveMatch.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/te-ctx-data/te-tunnel-info/" + passiveMatch.EntityData.SegmentPath
    passiveMatch.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    passiveMatch.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    passiveMatch.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    passiveMatch.EntityData.Children = types.NewOrderedMap()
    passiveMatch.EntityData.Leafs = types.NewOrderedMap()
    passiveMatch.EntityData.Leafs.Append("src-tid", types.YLeaf{"SrcTid", passiveMatch.SrcTid})
    passiveMatch.EntityData.Leafs.Append("src-rid", types.YLeaf{"SrcRid", passiveMatch.SrcRid})

    passiveMatch.EntityData.YListKeys = []string {}

    return &(passiveMatch.EntityData)
}

// Odu_Controllers_Controller_Info_XcAddCtxData
// Xconnect Add Data
type Odu_Controllers_Controller_Info_XcAddCtxData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Req Time. The type is interface{} with range: 0..4294967295.
    GmplsReqTime interface{}

    // Ctxt Type. The type is OtmOpticalRmCtxt.
    CtxtType interface{}

    // Rm Type. The type is OtmOpticalRmCtxtRm.
    RmType interface{}

    // Tunnel Information.
    TeTunnelInfo Odu_Controllers_Controller_Info_XcAddCtxData_TeTunnelInfo
}

func (xcAddCtxData *Odu_Controllers_Controller_Info_XcAddCtxData) GetEntityData() *types.CommonEntityData {
    xcAddCtxData.EntityData.YFilter = xcAddCtxData.YFilter
    xcAddCtxData.EntityData.YangName = "xc-add-ctx-data"
    xcAddCtxData.EntityData.BundleName = "cisco_ios_xr"
    xcAddCtxData.EntityData.ParentYangName = "info"
    xcAddCtxData.EntityData.SegmentPath = "xc-add-ctx-data"
    xcAddCtxData.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/" + xcAddCtxData.EntityData.SegmentPath
    xcAddCtxData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    xcAddCtxData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    xcAddCtxData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    xcAddCtxData.EntityData.Children = types.NewOrderedMap()
    xcAddCtxData.EntityData.Children.Append("te-tunnel-info", types.YChild{"TeTunnelInfo", &xcAddCtxData.TeTunnelInfo})
    xcAddCtxData.EntityData.Leafs = types.NewOrderedMap()
    xcAddCtxData.EntityData.Leafs.Append("gmpls-req-time", types.YLeaf{"GmplsReqTime", xcAddCtxData.GmplsReqTime})
    xcAddCtxData.EntityData.Leafs.Append("ctxt-type", types.YLeaf{"CtxtType", xcAddCtxData.CtxtType})
    xcAddCtxData.EntityData.Leafs.Append("rm-type", types.YLeaf{"RmType", xcAddCtxData.RmType})

    xcAddCtxData.EntityData.YListKeys = []string {}

    return &(xcAddCtxData.EntityData)
}

// Odu_Controllers_Controller_Info_XcAddCtxData_TeTunnelInfo
// Tunnel Information
type Odu_Controllers_Controller_Info_XcAddCtxData_TeTunnelInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // INFO TYPE. The type is OtmTeTunnelInfo.
    InfoType interface{}

    // Tunnel Id. The type is interface{} with range: 0..4294967295.
    TunnelId interface{}

    // Lbl Ctxt.
    LbCtxt Odu_Controllers_Controller_Info_XcAddCtxData_TeTunnelInfo_LbCtxt

    // Passive Match.
    PassiveMatch Odu_Controllers_Controller_Info_XcAddCtxData_TeTunnelInfo_PassiveMatch
}

func (teTunnelInfo *Odu_Controllers_Controller_Info_XcAddCtxData_TeTunnelInfo) GetEntityData() *types.CommonEntityData {
    teTunnelInfo.EntityData.YFilter = teTunnelInfo.YFilter
    teTunnelInfo.EntityData.YangName = "te-tunnel-info"
    teTunnelInfo.EntityData.BundleName = "cisco_ios_xr"
    teTunnelInfo.EntityData.ParentYangName = "xc-add-ctx-data"
    teTunnelInfo.EntityData.SegmentPath = "te-tunnel-info"
    teTunnelInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/xc-add-ctx-data/" + teTunnelInfo.EntityData.SegmentPath
    teTunnelInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    teTunnelInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    teTunnelInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    teTunnelInfo.EntityData.Children = types.NewOrderedMap()
    teTunnelInfo.EntityData.Children.Append("lb-ctxt", types.YChild{"LbCtxt", &teTunnelInfo.LbCtxt})
    teTunnelInfo.EntityData.Children.Append("passive-match", types.YChild{"PassiveMatch", &teTunnelInfo.PassiveMatch})
    teTunnelInfo.EntityData.Leafs = types.NewOrderedMap()
    teTunnelInfo.EntityData.Leafs.Append("info-type", types.YLeaf{"InfoType", teTunnelInfo.InfoType})
    teTunnelInfo.EntityData.Leafs.Append("tunnel-id", types.YLeaf{"TunnelId", teTunnelInfo.TunnelId})

    teTunnelInfo.EntityData.YListKeys = []string {}

    return &(teTunnelInfo.EntityData)
}

// Odu_Controllers_Controller_Info_XcAddCtxData_TeTunnelInfo_LbCtxt
// Lbl Ctxt
type Odu_Controllers_Controller_Info_XcAddCtxData_TeTunnelInfo_LbCtxt struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SubGroup Id. The type is interface{} with range: 0..65535.
    S2lFecSubGroupId interface{}

    // Lsp Id. The type is interface{} with range: 0..65535.
    S2lFecLspId interface{}

    // Tunnel Id. The type is interface{} with range: 0..65535.
    S2lFecTunnelId interface{}

    // Ext Tunnel Id. The type is interface{} with range: 0..4294967295.
    ExtTunnelId interface{}

    // FEC Source. The type is interface{} with range: 0..4294967295.
    FecSource interface{}

    // FEC Dest. The type is interface{} with range: 0..4294967295.
    FecDest interface{}

    // P2MP Id. The type is interface{} with range: 0..4294967295.
    S2lFecP2mpId interface{}

    // SubGroup Originator. The type is interface{} with range: 0..4294967295.
    SubGroupOriginAtor interface{}

    // Ctype. The type is OtmMplsLibC.
    FecCType interface{}
}

func (lbCtxt *Odu_Controllers_Controller_Info_XcAddCtxData_TeTunnelInfo_LbCtxt) GetEntityData() *types.CommonEntityData {
    lbCtxt.EntityData.YFilter = lbCtxt.YFilter
    lbCtxt.EntityData.YangName = "lb-ctxt"
    lbCtxt.EntityData.BundleName = "cisco_ios_xr"
    lbCtxt.EntityData.ParentYangName = "te-tunnel-info"
    lbCtxt.EntityData.SegmentPath = "lb-ctxt"
    lbCtxt.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/xc-add-ctx-data/te-tunnel-info/" + lbCtxt.EntityData.SegmentPath
    lbCtxt.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lbCtxt.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lbCtxt.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lbCtxt.EntityData.Children = types.NewOrderedMap()
    lbCtxt.EntityData.Leafs = types.NewOrderedMap()
    lbCtxt.EntityData.Leafs.Append("s2l-fec-sub-group-id", types.YLeaf{"S2lFecSubGroupId", lbCtxt.S2lFecSubGroupId})
    lbCtxt.EntityData.Leafs.Append("s2l-fec-lsp-id", types.YLeaf{"S2lFecLspId", lbCtxt.S2lFecLspId})
    lbCtxt.EntityData.Leafs.Append("s2l-fec-tunnel-id", types.YLeaf{"S2lFecTunnelId", lbCtxt.S2lFecTunnelId})
    lbCtxt.EntityData.Leafs.Append("ext-tunnel-id", types.YLeaf{"ExtTunnelId", lbCtxt.ExtTunnelId})
    lbCtxt.EntityData.Leafs.Append("fec-source", types.YLeaf{"FecSource", lbCtxt.FecSource})
    lbCtxt.EntityData.Leafs.Append("fec-dest", types.YLeaf{"FecDest", lbCtxt.FecDest})
    lbCtxt.EntityData.Leafs.Append("s2l-fec-p2mp-id", types.YLeaf{"S2lFecP2mpId", lbCtxt.S2lFecP2mpId})
    lbCtxt.EntityData.Leafs.Append("sub-group-origin-ator", types.YLeaf{"SubGroupOriginAtor", lbCtxt.SubGroupOriginAtor})
    lbCtxt.EntityData.Leafs.Append("fec-c-type", types.YLeaf{"FecCType", lbCtxt.FecCType})

    lbCtxt.EntityData.YListKeys = []string {}

    return &(lbCtxt.EntityData)
}

// Odu_Controllers_Controller_Info_XcAddCtxData_TeTunnelInfo_PassiveMatch
// Passive Match
type Odu_Controllers_Controller_Info_XcAddCtxData_TeTunnelInfo_PassiveMatch struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Src TId. The type is interface{} with range: 0..65535.
    SrcTid interface{}

    // Src RId. The type is interface{} with range: 0..4294967295.
    SrcRid interface{}
}

func (passiveMatch *Odu_Controllers_Controller_Info_XcAddCtxData_TeTunnelInfo_PassiveMatch) GetEntityData() *types.CommonEntityData {
    passiveMatch.EntityData.YFilter = passiveMatch.YFilter
    passiveMatch.EntityData.YangName = "passive-match"
    passiveMatch.EntityData.BundleName = "cisco_ios_xr"
    passiveMatch.EntityData.ParentYangName = "te-tunnel-info"
    passiveMatch.EntityData.SegmentPath = "passive-match"
    passiveMatch.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/xc-add-ctx-data/te-tunnel-info/" + passiveMatch.EntityData.SegmentPath
    passiveMatch.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    passiveMatch.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    passiveMatch.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    passiveMatch.EntityData.Children = types.NewOrderedMap()
    passiveMatch.EntityData.Leafs = types.NewOrderedMap()
    passiveMatch.EntityData.Leafs.Append("src-tid", types.YLeaf{"SrcTid", passiveMatch.SrcTid})
    passiveMatch.EntityData.Leafs.Append("src-rid", types.YLeaf{"SrcRid", passiveMatch.SrcRid})

    passiveMatch.EntityData.YListKeys = []string {}

    return &(passiveMatch.EntityData)
}

// Odu_Controllers_Controller_Info_XcRemCtxData
// Xconnect Remove Data
type Odu_Controllers_Controller_Info_XcRemCtxData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Req Time. The type is interface{} with range: 0..4294967295.
    GmplsReqTime interface{}

    // Ctxt Type. The type is OtmOpticalRmCtxt.
    CtxtType interface{}

    // Rm Type. The type is OtmOpticalRmCtxtRm.
    RmType interface{}

    // Tunnel Information.
    TeTunnelInfo Odu_Controllers_Controller_Info_XcRemCtxData_TeTunnelInfo
}

func (xcRemCtxData *Odu_Controllers_Controller_Info_XcRemCtxData) GetEntityData() *types.CommonEntityData {
    xcRemCtxData.EntityData.YFilter = xcRemCtxData.YFilter
    xcRemCtxData.EntityData.YangName = "xc-rem-ctx-data"
    xcRemCtxData.EntityData.BundleName = "cisco_ios_xr"
    xcRemCtxData.EntityData.ParentYangName = "info"
    xcRemCtxData.EntityData.SegmentPath = "xc-rem-ctx-data"
    xcRemCtxData.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/" + xcRemCtxData.EntityData.SegmentPath
    xcRemCtxData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    xcRemCtxData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    xcRemCtxData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    xcRemCtxData.EntityData.Children = types.NewOrderedMap()
    xcRemCtxData.EntityData.Children.Append("te-tunnel-info", types.YChild{"TeTunnelInfo", &xcRemCtxData.TeTunnelInfo})
    xcRemCtxData.EntityData.Leafs = types.NewOrderedMap()
    xcRemCtxData.EntityData.Leafs.Append("gmpls-req-time", types.YLeaf{"GmplsReqTime", xcRemCtxData.GmplsReqTime})
    xcRemCtxData.EntityData.Leafs.Append("ctxt-type", types.YLeaf{"CtxtType", xcRemCtxData.CtxtType})
    xcRemCtxData.EntityData.Leafs.Append("rm-type", types.YLeaf{"RmType", xcRemCtxData.RmType})

    xcRemCtxData.EntityData.YListKeys = []string {}

    return &(xcRemCtxData.EntityData)
}

// Odu_Controllers_Controller_Info_XcRemCtxData_TeTunnelInfo
// Tunnel Information
type Odu_Controllers_Controller_Info_XcRemCtxData_TeTunnelInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // INFO TYPE. The type is OtmTeTunnelInfo.
    InfoType interface{}

    // Tunnel Id. The type is interface{} with range: 0..4294967295.
    TunnelId interface{}

    // Lbl Ctxt.
    LbCtxt Odu_Controllers_Controller_Info_XcRemCtxData_TeTunnelInfo_LbCtxt

    // Passive Match.
    PassiveMatch Odu_Controllers_Controller_Info_XcRemCtxData_TeTunnelInfo_PassiveMatch
}

func (teTunnelInfo *Odu_Controllers_Controller_Info_XcRemCtxData_TeTunnelInfo) GetEntityData() *types.CommonEntityData {
    teTunnelInfo.EntityData.YFilter = teTunnelInfo.YFilter
    teTunnelInfo.EntityData.YangName = "te-tunnel-info"
    teTunnelInfo.EntityData.BundleName = "cisco_ios_xr"
    teTunnelInfo.EntityData.ParentYangName = "xc-rem-ctx-data"
    teTunnelInfo.EntityData.SegmentPath = "te-tunnel-info"
    teTunnelInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/xc-rem-ctx-data/" + teTunnelInfo.EntityData.SegmentPath
    teTunnelInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    teTunnelInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    teTunnelInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    teTunnelInfo.EntityData.Children = types.NewOrderedMap()
    teTunnelInfo.EntityData.Children.Append("lb-ctxt", types.YChild{"LbCtxt", &teTunnelInfo.LbCtxt})
    teTunnelInfo.EntityData.Children.Append("passive-match", types.YChild{"PassiveMatch", &teTunnelInfo.PassiveMatch})
    teTunnelInfo.EntityData.Leafs = types.NewOrderedMap()
    teTunnelInfo.EntityData.Leafs.Append("info-type", types.YLeaf{"InfoType", teTunnelInfo.InfoType})
    teTunnelInfo.EntityData.Leafs.Append("tunnel-id", types.YLeaf{"TunnelId", teTunnelInfo.TunnelId})

    teTunnelInfo.EntityData.YListKeys = []string {}

    return &(teTunnelInfo.EntityData)
}

// Odu_Controllers_Controller_Info_XcRemCtxData_TeTunnelInfo_LbCtxt
// Lbl Ctxt
type Odu_Controllers_Controller_Info_XcRemCtxData_TeTunnelInfo_LbCtxt struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SubGroup Id. The type is interface{} with range: 0..65535.
    S2lFecSubGroupId interface{}

    // Lsp Id. The type is interface{} with range: 0..65535.
    S2lFecLspId interface{}

    // Tunnel Id. The type is interface{} with range: 0..65535.
    S2lFecTunnelId interface{}

    // Ext Tunnel Id. The type is interface{} with range: 0..4294967295.
    ExtTunnelId interface{}

    // FEC Source. The type is interface{} with range: 0..4294967295.
    FecSource interface{}

    // FEC Dest. The type is interface{} with range: 0..4294967295.
    FecDest interface{}

    // P2MP Id. The type is interface{} with range: 0..4294967295.
    S2lFecP2mpId interface{}

    // SubGroup Originator. The type is interface{} with range: 0..4294967295.
    SubGroupOriginAtor interface{}

    // Ctype. The type is OtmMplsLibC.
    FecCType interface{}
}

func (lbCtxt *Odu_Controllers_Controller_Info_XcRemCtxData_TeTunnelInfo_LbCtxt) GetEntityData() *types.CommonEntityData {
    lbCtxt.EntityData.YFilter = lbCtxt.YFilter
    lbCtxt.EntityData.YangName = "lb-ctxt"
    lbCtxt.EntityData.BundleName = "cisco_ios_xr"
    lbCtxt.EntityData.ParentYangName = "te-tunnel-info"
    lbCtxt.EntityData.SegmentPath = "lb-ctxt"
    lbCtxt.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/xc-rem-ctx-data/te-tunnel-info/" + lbCtxt.EntityData.SegmentPath
    lbCtxt.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lbCtxt.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lbCtxt.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lbCtxt.EntityData.Children = types.NewOrderedMap()
    lbCtxt.EntityData.Leafs = types.NewOrderedMap()
    lbCtxt.EntityData.Leafs.Append("s2l-fec-sub-group-id", types.YLeaf{"S2lFecSubGroupId", lbCtxt.S2lFecSubGroupId})
    lbCtxt.EntityData.Leafs.Append("s2l-fec-lsp-id", types.YLeaf{"S2lFecLspId", lbCtxt.S2lFecLspId})
    lbCtxt.EntityData.Leafs.Append("s2l-fec-tunnel-id", types.YLeaf{"S2lFecTunnelId", lbCtxt.S2lFecTunnelId})
    lbCtxt.EntityData.Leafs.Append("ext-tunnel-id", types.YLeaf{"ExtTunnelId", lbCtxt.ExtTunnelId})
    lbCtxt.EntityData.Leafs.Append("fec-source", types.YLeaf{"FecSource", lbCtxt.FecSource})
    lbCtxt.EntityData.Leafs.Append("fec-dest", types.YLeaf{"FecDest", lbCtxt.FecDest})
    lbCtxt.EntityData.Leafs.Append("s2l-fec-p2mp-id", types.YLeaf{"S2lFecP2mpId", lbCtxt.S2lFecP2mpId})
    lbCtxt.EntityData.Leafs.Append("sub-group-origin-ator", types.YLeaf{"SubGroupOriginAtor", lbCtxt.SubGroupOriginAtor})
    lbCtxt.EntityData.Leafs.Append("fec-c-type", types.YLeaf{"FecCType", lbCtxt.FecCType})

    lbCtxt.EntityData.YListKeys = []string {}

    return &(lbCtxt.EntityData)
}

// Odu_Controllers_Controller_Info_XcRemCtxData_TeTunnelInfo_PassiveMatch
// Passive Match
type Odu_Controllers_Controller_Info_XcRemCtxData_TeTunnelInfo_PassiveMatch struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Src TId. The type is interface{} with range: 0..65535.
    SrcTid interface{}

    // Src RId. The type is interface{} with range: 0..4294967295.
    SrcRid interface{}
}

func (passiveMatch *Odu_Controllers_Controller_Info_XcRemCtxData_TeTunnelInfo_PassiveMatch) GetEntityData() *types.CommonEntityData {
    passiveMatch.EntityData.YFilter = passiveMatch.YFilter
    passiveMatch.EntityData.YangName = "passive-match"
    passiveMatch.EntityData.BundleName = "cisco_ios_xr"
    passiveMatch.EntityData.ParentYangName = "te-tunnel-info"
    passiveMatch.EntityData.SegmentPath = "passive-match"
    passiveMatch.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/xc-rem-ctx-data/te-tunnel-info/" + passiveMatch.EntityData.SegmentPath
    passiveMatch.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    passiveMatch.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    passiveMatch.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    passiveMatch.EntityData.Children = types.NewOrderedMap()
    passiveMatch.EntityData.Leafs = types.NewOrderedMap()
    passiveMatch.EntityData.Leafs.Append("src-tid", types.YLeaf{"SrcTid", passiveMatch.SrcTid})
    passiveMatch.EntityData.Leafs.Append("src-rid", types.YLeaf{"SrcRid", passiveMatch.SrcRid})

    passiveMatch.EntityData.YListKeys = []string {}

    return &(passiveMatch.EntityData)
}

// Odu_Controllers_Controller_Info_OduDelay
// ODU Delay
type Odu_Controllers_Controller_Info_OduDelay struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Latency Mode. The type is interface{} with range: 0..255.
    Mode interface{}

    // Delay Value. The type is interface{} with range: 0..4294967295.
    Delay interface{}
}

func (oduDelay *Odu_Controllers_Controller_Info_OduDelay) GetEntityData() *types.CommonEntityData {
    oduDelay.EntityData.YFilter = oduDelay.YFilter
    oduDelay.EntityData.YangName = "odu-delay"
    oduDelay.EntityData.BundleName = "cisco_ios_xr"
    oduDelay.EntityData.ParentYangName = "info"
    oduDelay.EntityData.SegmentPath = "odu-delay"
    oduDelay.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/" + oduDelay.EntityData.SegmentPath
    oduDelay.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    oduDelay.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    oduDelay.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    oduDelay.EntityData.Children = types.NewOrderedMap()
    oduDelay.EntityData.Leafs = types.NewOrderedMap()
    oduDelay.EntityData.Leafs.Append("mode", types.YLeaf{"Mode", oduDelay.Mode})
    oduDelay.EntityData.Leafs.Append("delay", types.YLeaf{"Delay", oduDelay.Delay})

    oduDelay.EntityData.YListKeys = []string {}

    return &(oduDelay.EntityData)
}

// Odu_Controllers_Controller_Info_OduTerminateEther
// odu terminate ether
type Odu_Controllers_Controller_Info_OduTerminateEther struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // interface handle. The type is interface{} with range: 0..4294967295.
    VetherIfhandle interface{}

    // ethernet mapping. The type is OduEtherMapPingEt.
    EthernetMapping interface{}

    // Ethernet interface name. The type is string.
    EthernetInterface interface{}
}

func (oduTerminateEther *Odu_Controllers_Controller_Info_OduTerminateEther) GetEntityData() *types.CommonEntityData {
    oduTerminateEther.EntityData.YFilter = oduTerminateEther.YFilter
    oduTerminateEther.EntityData.YangName = "odu-terminate-ether"
    oduTerminateEther.EntityData.BundleName = "cisco_ios_xr"
    oduTerminateEther.EntityData.ParentYangName = "info"
    oduTerminateEther.EntityData.SegmentPath = "odu-terminate-ether"
    oduTerminateEther.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/" + oduTerminateEther.EntityData.SegmentPath
    oduTerminateEther.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    oduTerminateEther.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    oduTerminateEther.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    oduTerminateEther.EntityData.Children = types.NewOrderedMap()
    oduTerminateEther.EntityData.Leafs = types.NewOrderedMap()
    oduTerminateEther.EntityData.Leafs.Append("vether-ifhandle", types.YLeaf{"VetherIfhandle", oduTerminateEther.VetherIfhandle})
    oduTerminateEther.EntityData.Leafs.Append("ethernet-mapping", types.YLeaf{"EthernetMapping", oduTerminateEther.EthernetMapping})
    oduTerminateEther.EntityData.Leafs.Append("ethernet-interface", types.YLeaf{"EthernetInterface", oduTerminateEther.EthernetInterface})

    oduTerminateEther.EntityData.YListKeys = []string {}

    return &(oduTerminateEther.EntityData)
}

// Odu_Controllers_Controller_Info_AinsInfo
// AINS information
type Odu_Controllers_Controller_Info_AinsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AINS State. The type is OduAinsStateEt.
    AinsState interface{}

    // AINS Timer in Minutes. The type is interface{} with range: 0..4294967295.
    // Units are minute.
    AinsTimerMinutes interface{}

    // AINS Remaining Seconds. The type is interface{} with range: 0..4294967295.
    // Units are second.
    AinsRemainingSecs interface{}
}

func (ainsInfo *Odu_Controllers_Controller_Info_AinsInfo) GetEntityData() *types.CommonEntityData {
    ainsInfo.EntityData.YFilter = ainsInfo.YFilter
    ainsInfo.EntityData.YangName = "ains-info"
    ainsInfo.EntityData.BundleName = "cisco_ios_xr"
    ainsInfo.EntityData.ParentYangName = "info"
    ainsInfo.EntityData.SegmentPath = "ains-info"
    ainsInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/" + ainsInfo.EntityData.SegmentPath
    ainsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ainsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ainsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ainsInfo.EntityData.Children = types.NewOrderedMap()
    ainsInfo.EntityData.Leafs = types.NewOrderedMap()
    ainsInfo.EntityData.Leafs.Append("ains-state", types.YLeaf{"AinsState", ainsInfo.AinsState})
    ainsInfo.EntityData.Leafs.Append("ains-timer-minutes", types.YLeaf{"AinsTimerMinutes", ainsInfo.AinsTimerMinutes})
    ainsInfo.EntityData.Leafs.Append("ains-remaining-secs", types.YLeaf{"AinsRemainingSecs", ainsInfo.AinsRemainingSecs})

    ainsInfo.EntityData.YListKeys = []string {}

    return &(ainsInfo.EntityData)
}

// Odu_Controllers_Controller_Info_Odu
// Child Ts
type Odu_Controllers_Controller_Info_Odu struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Child Interface Name. The type is string with length: 0..64.
    IntfName interface{}

    // Tpn Bitmap. The type is interface{} with range: 0..255.
    TpnValue interface{}

    // Ts Bitmap. The type is string with length: 0..256.
    TsBitmap interface{}
}

func (odu *Odu_Controllers_Controller_Info_Odu) GetEntityData() *types.CommonEntityData {
    odu.EntityData.YFilter = odu.YFilter
    odu.EntityData.YangName = "odu"
    odu.EntityData.BundleName = "cisco_ios_xr"
    odu.EntityData.ParentYangName = "info"
    odu.EntityData.SegmentPath = "odu" + types.AddNoKeyToken(odu)
    odu.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/" + odu.EntityData.SegmentPath
    odu.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    odu.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    odu.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    odu.EntityData.Children = types.NewOrderedMap()
    odu.EntityData.Leafs = types.NewOrderedMap()
    odu.EntityData.Leafs.Append("intf-name", types.YLeaf{"IntfName", odu.IntfName})
    odu.EntityData.Leafs.Append("tpn-value", types.YLeaf{"TpnValue", odu.TpnValue})
    odu.EntityData.Leafs.Append("ts-bitmap", types.YLeaf{"TsBitmap", odu.TsBitmap})

    odu.EntityData.YListKeys = []string {}

    return &(odu.EntityData)
}

// Odu_Controllers_Controller_Info_Odutcm
// ODU TCM
type Odu_Controllers_Controller_Info_Odutcm struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // ODU TCM SF in the form of 1.0E - <SF>. The type is interface{} with range:
    // 0..255.
    Tcmsf interface{}

    // ODU TCM SD in the form of 1.0E - <SD>. The type is interface{} with range:
    // 0..255.
    Tcmsd interface{}

    // ODU TCM state. The type is OduTcmStateEt.
    TcmState interface{}

    // Performance Monitoring. The type is OduTcmPerMon.
    TcmperMon interface{}

    // ODU TCM Mode. The type is OduTcmMode.
    TcmMode interface{}

    // TCM Mode in H/W. The type is OduTcmMode.
    ActualTcmMode interface{}

    // ODU TCM LTC CA state. The type is OduTcmStateEt.
    TcmltcState interface{}

    // ODU TCM  TIM CAstate. The type is OduTcmStateEt.
    TcmtimState interface{}

    // ODU TCM DELAY. The type is interface{} with range: 0..4294967295.
    TcmDelay interface{}

    // TTI.
    TcmttiMode Odu_Controllers_Controller_Info_Odutcm_TcmttiMode
}

func (odutcm *Odu_Controllers_Controller_Info_Odutcm) GetEntityData() *types.CommonEntityData {
    odutcm.EntityData.YFilter = odutcm.YFilter
    odutcm.EntityData.YangName = "odutcm"
    odutcm.EntityData.BundleName = "cisco_ios_xr"
    odutcm.EntityData.ParentYangName = "info"
    odutcm.EntityData.SegmentPath = "odutcm" + types.AddNoKeyToken(odutcm)
    odutcm.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/" + odutcm.EntityData.SegmentPath
    odutcm.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    odutcm.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    odutcm.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    odutcm.EntityData.Children = types.NewOrderedMap()
    odutcm.EntityData.Children.Append("tcmtti-mode", types.YChild{"TcmttiMode", &odutcm.TcmttiMode})
    odutcm.EntityData.Leafs = types.NewOrderedMap()
    odutcm.EntityData.Leafs.Append("tcmsf", types.YLeaf{"Tcmsf", odutcm.Tcmsf})
    odutcm.EntityData.Leafs.Append("tcmsd", types.YLeaf{"Tcmsd", odutcm.Tcmsd})
    odutcm.EntityData.Leafs.Append("tcm-state", types.YLeaf{"TcmState", odutcm.TcmState})
    odutcm.EntityData.Leafs.Append("tcmper-mon", types.YLeaf{"TcmperMon", odutcm.TcmperMon})
    odutcm.EntityData.Leafs.Append("tcm-mode", types.YLeaf{"TcmMode", odutcm.TcmMode})
    odutcm.EntityData.Leafs.Append("actual-tcm-mode", types.YLeaf{"ActualTcmMode", odutcm.ActualTcmMode})
    odutcm.EntityData.Leafs.Append("tcmltc-state", types.YLeaf{"TcmltcState", odutcm.TcmltcState})
    odutcm.EntityData.Leafs.Append("tcmtim-state", types.YLeaf{"TcmtimState", odutcm.TcmtimState})
    odutcm.EntityData.Leafs.Append("tcm-delay", types.YLeaf{"TcmDelay", odutcm.TcmDelay})

    odutcm.EntityData.YListKeys = []string {}

    return &(odutcm.EntityData)
}

// Odu_Controllers_Controller_Info_Odutcm_TcmttiMode
// TTI
type Odu_Controllers_Controller_Info_Odutcm_TcmttiMode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // G709TTI Sent. The type is OduTtiEt.
    G709ttiSentMode interface{}

    // G709TTI Expected. The type is OduTtiEt.
    G709ttiExpMode interface{}

    // G709TTI Recieved. The type is OduTtiEt.
    G709ttiRecMode interface{}

    // String Sent.
    Tx Odu_Controllers_Controller_Info_Odutcm_TcmttiMode_Tx

    // String Expected.
    Exp Odu_Controllers_Controller_Info_Odutcm_TcmttiMode_Exp

    // String Received.
    Rec Odu_Controllers_Controller_Info_Odutcm_TcmttiMode_Rec
}

func (tcmttiMode *Odu_Controllers_Controller_Info_Odutcm_TcmttiMode) GetEntityData() *types.CommonEntityData {
    tcmttiMode.EntityData.YFilter = tcmttiMode.YFilter
    tcmttiMode.EntityData.YangName = "tcmtti-mode"
    tcmttiMode.EntityData.BundleName = "cisco_ios_xr"
    tcmttiMode.EntityData.ParentYangName = "odutcm"
    tcmttiMode.EntityData.SegmentPath = "tcmtti-mode"
    tcmttiMode.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/odutcm/" + tcmttiMode.EntityData.SegmentPath
    tcmttiMode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcmttiMode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcmttiMode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcmttiMode.EntityData.Children = types.NewOrderedMap()
    tcmttiMode.EntityData.Children.Append("tx", types.YChild{"Tx", &tcmttiMode.Tx})
    tcmttiMode.EntityData.Children.Append("exp", types.YChild{"Exp", &tcmttiMode.Exp})
    tcmttiMode.EntityData.Children.Append("rec", types.YChild{"Rec", &tcmttiMode.Rec})
    tcmttiMode.EntityData.Leafs = types.NewOrderedMap()
    tcmttiMode.EntityData.Leafs.Append("g709tti-sent-mode", types.YLeaf{"G709ttiSentMode", tcmttiMode.G709ttiSentMode})
    tcmttiMode.EntityData.Leafs.Append("g709tti-exp-mode", types.YLeaf{"G709ttiExpMode", tcmttiMode.G709ttiExpMode})
    tcmttiMode.EntityData.Leafs.Append("g709tti-rec-mode", types.YLeaf{"G709ttiRecMode", tcmttiMode.G709ttiRecMode})

    tcmttiMode.EntityData.YListKeys = []string {}

    return &(tcmttiMode.EntityData)
}

// Odu_Controllers_Controller_Info_Odutcm_TcmttiMode_Tx
// String Sent
type Odu_Controllers_Controller_Info_Odutcm_TcmttiMode_Tx struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // tx String. The type is slice of interface{} with range: 0..255.
    Sapi []interface{}

    // exp String. The type is slice of interface{} with range: 0..255.
    Dapi []interface{}

    // rec String. The type is slice of interface{} with range: 0..255.
    OperatorSpecific []interface{}
}

func (tx *Odu_Controllers_Controller_Info_Odutcm_TcmttiMode_Tx) GetEntityData() *types.CommonEntityData {
    tx.EntityData.YFilter = tx.YFilter
    tx.EntityData.YangName = "tx"
    tx.EntityData.BundleName = "cisco_ios_xr"
    tx.EntityData.ParentYangName = "tcmtti-mode"
    tx.EntityData.SegmentPath = "tx"
    tx.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/odutcm/tcmtti-mode/" + tx.EntityData.SegmentPath
    tx.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tx.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tx.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tx.EntityData.Children = types.NewOrderedMap()
    tx.EntityData.Leafs = types.NewOrderedMap()
    tx.EntityData.Leafs.Append("sapi", types.YLeaf{"Sapi", tx.Sapi})
    tx.EntityData.Leafs.Append("dapi", types.YLeaf{"Dapi", tx.Dapi})
    tx.EntityData.Leafs.Append("operator-specific", types.YLeaf{"OperatorSpecific", tx.OperatorSpecific})

    tx.EntityData.YListKeys = []string {}

    return &(tx.EntityData)
}

// Odu_Controllers_Controller_Info_Odutcm_TcmttiMode_Exp
// String Expected
type Odu_Controllers_Controller_Info_Odutcm_TcmttiMode_Exp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // tx String. The type is slice of interface{} with range: 0..255.
    Sapi []interface{}

    // exp String. The type is slice of interface{} with range: 0..255.
    Dapi []interface{}

    // rec String. The type is slice of interface{} with range: 0..255.
    OperatorSpecific []interface{}
}

func (exp *Odu_Controllers_Controller_Info_Odutcm_TcmttiMode_Exp) GetEntityData() *types.CommonEntityData {
    exp.EntityData.YFilter = exp.YFilter
    exp.EntityData.YangName = "exp"
    exp.EntityData.BundleName = "cisco_ios_xr"
    exp.EntityData.ParentYangName = "tcmtti-mode"
    exp.EntityData.SegmentPath = "exp"
    exp.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/odutcm/tcmtti-mode/" + exp.EntityData.SegmentPath
    exp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    exp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    exp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    exp.EntityData.Children = types.NewOrderedMap()
    exp.EntityData.Leafs = types.NewOrderedMap()
    exp.EntityData.Leafs.Append("sapi", types.YLeaf{"Sapi", exp.Sapi})
    exp.EntityData.Leafs.Append("dapi", types.YLeaf{"Dapi", exp.Dapi})
    exp.EntityData.Leafs.Append("operator-specific", types.YLeaf{"OperatorSpecific", exp.OperatorSpecific})

    exp.EntityData.YListKeys = []string {}

    return &(exp.EntityData)
}

// Odu_Controllers_Controller_Info_Odutcm_TcmttiMode_Rec
// String Received
type Odu_Controllers_Controller_Info_Odutcm_TcmttiMode_Rec struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // tx String. The type is slice of interface{} with range: 0..255.
    Sapi []interface{}

    // exp String. The type is slice of interface{} with range: 0..255.
    Dapi []interface{}

    // rec String. The type is slice of interface{} with range: 0..255.
    OperatorSpecific []interface{}
}

func (rec *Odu_Controllers_Controller_Info_Odutcm_TcmttiMode_Rec) GetEntityData() *types.CommonEntityData {
    rec.EntityData.YFilter = rec.YFilter
    rec.EntityData.YangName = "rec"
    rec.EntityData.BundleName = "cisco_ios_xr"
    rec.EntityData.ParentYangName = "tcmtti-mode"
    rec.EntityData.SegmentPath = "rec"
    rec.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-odu-oper:odu/controllers/controller/info/odutcm/tcmtti-mode/" + rec.EntityData.SegmentPath
    rec.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rec.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rec.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rec.EntityData.Children = types.NewOrderedMap()
    rec.EntityData.Leafs = types.NewOrderedMap()
    rec.EntityData.Leafs.Append("sapi", types.YLeaf{"Sapi", rec.Sapi})
    rec.EntityData.Leafs.Append("dapi", types.YLeaf{"Dapi", rec.Dapi})
    rec.EntityData.Leafs.Append("operator-specific", types.YLeaf{"OperatorSpecific", rec.OperatorSpecific})

    rec.EntityData.YListKeys = []string {}

    return &(rec.EntityData)
}

