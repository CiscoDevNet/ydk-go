// This module contains a collection of YANG definitions
// for Cisco IOS-XR pmengine package configuration.
// 
// This YANG module augments the
//   Cisco-IOS-XR-ifmgr-cfg
// module with configuration data.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package pmengine_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package pmengine_cfg"))
}

// SecyifReport represents Secyif report
type SecyifReport string

const (
    // PM SECYIF if InPktsUntagged report
    SecyifReport_report_if_inpkts_untagged SecyifReport = "report-if-inpkts-untagged"

    // PM SECYIF if InPktsNoTag report
    SecyifReport_report_if_inpkts_not_ag SecyifReport = "report-if-inpkts-not-ag"

    // PM SECYIF if InPktsBadTag report
    SecyifReport_report_if_inpktsbadtag SecyifReport = "report-if-inpktsbadtag"

    // PM SECYIF if InPktsUnknownSCI report
    SecyifReport_report_if_inpkts_unknown_sci SecyifReport = "report-if-inpkts-unknown-sci"

    // PM SECYIF if InPktsNoSCI report
    SecyifReport_report_if_inpktsnosci SecyifReport = "report-if-inpktsnosci"

    // PM SECYIF if InPktsOverrun report
    SecyifReport_report_if_inpkts_overrun SecyifReport = "report-if-inpkts-overrun"

    // PM SECYIF if InOctetsValidated report
    SecyifReport_report_if_inoctets_validate_d SecyifReport = "report-if-inoctets-validate-d"

    // PM SECYIF if InOctetsDecrypted report
    SecyifReport_report_if_inoctetsdecrypted SecyifReport = "report-if-inoctetsdecrypted"

    // PM SECYIF if OutPktsUntagged report
    SecyifReport_report_if_outpkts_untagged SecyifReport = "report-if-outpkts-untagged"

    // PM SECYIF if OutPktsTooLong report
    SecyifReport_report_if_outpkts_too_long SecyifReport = "report-if-outpkts-too-long"

    // PM SECYIF if OutOctetsProtected report
    SecyifReport_report_if_outoctetsprotected SecyifReport = "report-if-outoctetsprotected"

    // PM SECYIF if OutOctetsEncrypted report
    SecyifReport_report_if_outoctetsencrypted SecyifReport = "report-if-outoctetsencrypted"
)

// GfpReport represents Gfp report
type GfpReport string

const (
    // PM GFP rx bit err report
    GfpReport_report_rx_bit_err GfpReport = "report-rx-bit-err"

    // PM GFP rx inv type report
    GfpReport_report_rx_inv_typ GfpReport = "report-rx-inv-typ"

    // PM GFP rx crc report
    GfpReport_report_rx_crc GfpReport = "report-rx-crc"

    // PM GFP rx lfd report
    GfpReport_report_rx_lfd GfpReport = "report-rx-lfd"

    // PM GFP rx csf report
    GfpReport_report_rx_csf GfpReport = "report-rx-csf"
)

// HoVcThreshold represents Ho vc threshold
type HoVcThreshold string

const (
    // PM EB threshold
    HoVcThreshold_thresh_eb HoVcThreshold = "thresh-eb"

    // PM ES threshold
    HoVcThreshold_thresh_es HoVcThreshold = "thresh-es"

    // PM ESR threshold
    HoVcThreshold_thresh_esr HoVcThreshold = "thresh-esr"

    // PM SES threshold
    HoVcThreshold_thresh_ses HoVcThreshold = "thresh-ses"

    // PM SESR threshold
    HoVcThreshold_thresh_sesr HoVcThreshold = "thresh-sesr"

    // PM BBE threshold
    HoVcThreshold_thresh_bbe HoVcThreshold = "thresh-bbe"

    // PM BBER threshold
    HoVcThreshold_thresh_bber HoVcThreshold = "thresh-bber"

    // PM UASS threshold
    HoVcThreshold_thresh_uass HoVcThreshold = "thresh-uass"
)

// EtherThreshold represents Ether threshold
type EtherThreshold string

const (
    // PM Ether rx pkt threshold
    EtherThreshold_thresh_rx_pkt EtherThreshold = "thresh-rx-pkt"

    // PM Ether rx util threshold
    EtherThreshold_thresh_rx_util EtherThreshold = "thresh-rx-util"

    // PM Ether tx util threshold
    EtherThreshold_thresh_tx_util EtherThreshold = "thresh-tx-util"

    // PM ether stat pkt threshold
    EtherThreshold_thresh_stat_pkt EtherThreshold = "thresh-stat-pkt"

    // PM Ether octet stat threshold
    EtherThreshold_thresh_octet_stat EtherThreshold = "thresh-octet-stat"

    // PM Ether oversize pkt threshold
    EtherThreshold_thresh_over_size_pkt EtherThreshold = "thresh-over-size-pkt"

    // PMEther fcs error threshold
    EtherThreshold_thresh_fcs_err EtherThreshold = "thresh-fcs-err"

    // PM Ether long frames threshold
    EtherThreshold_thresh_long_frame_s EtherThreshold = "thresh-long-frame-s"

    // PM Ether jabber stats threshold
    EtherThreshold_thresh_jabber_stats EtherThreshold = "thresh-jabber-stats"

    // PM Ether 64 octet threshold
    EtherThreshold_thresh_64_octet EtherThreshold = "thresh-64-octet"

    // PM Ether 65 to 127 octet threshold
    EtherThreshold_thresh_65_127_octet EtherThreshold = "thresh-65-127-octet"

    // PM Ether 128 to 255 octet threshold
    EtherThreshold_thresh_128_255_octet EtherThreshold = "thresh-128-255-octet"

    // PM Ether 256 to 511 octet threshold
    EtherThreshold_thresh_256_511_octet EtherThreshold = "thresh-256-511-octet"

    // PM Ether 512 to 1023 octet threshold
    EtherThreshold_thresh_512_1023_octet EtherThreshold = "thresh-512-1023-octet"

    // PM Ether 1024 to 1518 threshold
    EtherThreshold_thresh_1024_1518_octet EtherThreshold = "thresh-1024-1518-octet"

    // PM Ether rx ucast threshold
    EtherThreshold_thresh_in_ucast EtherThreshold = "thresh-in-ucast"

    // PM Ether rx mcast threshold
    EtherThreshold_thresh_in_mcast EtherThreshold = "thresh-in-mcast"

    // PM Ether rx bcast threshold
    EtherThreshold_thresh_in_bcast EtherThreshold = "thresh-in-bcast"

    // PM Ether tx ucast threshold
    EtherThreshold_thresh_out_ucast EtherThreshold = "thresh-out-ucast"

    // PM Ether tx mcast threshold
    EtherThreshold_thresh_out_mcast EtherThreshold = "thresh-out-mcast"

    // PM Ether tx bcast threshold
    EtherThreshold_thresh_out_bcast EtherThreshold = "thresh-out-bcast"

    // PM Ether tx pkt threshold
    EtherThreshold_thresh_tx_pkt EtherThreshold = "thresh-tx-pkt"

    // PM ether ifIn errors threshold
    EtherThreshold_thresh_ifin_error_s EtherThreshold = "thresh-ifin-error-s"

    // PM ether ifInOctets threshold
    EtherThreshold_thresh_ifin_octets EtherThreshold = "thresh-ifin-octets"

    // PM ether stat multicast pkt threshold
    EtherThreshold_thresh_ether_stat_multicast_pkt EtherThreshold = "thresh-ether-stat-multicast-pkt"

    // PM ether stat broadcast pkt threshold
    EtherThreshold_thresh_ether_stat_broadcast_pkt EtherThreshold = "thresh-ether-stat-broadcast-pkt"

    // PM ether stat undersized pkt threshold
    EtherThreshold_thresh_ether_stat_under_size_d_pkt EtherThreshold = "thresh-ether-stat-under-size-d-pkt"

    // PM ether out octets threshold
    EtherThreshold_thresh_out_octet EtherThreshold = "thresh-out-octet"

    // PM in pause frame threshold
    EtherThreshold_thresh_in_pause_frame EtherThreshold = "thresh-in-pause-frame"

    // PM in good bytes threshold
    EtherThreshold_thresh_in_go_od_bytes EtherThreshold = "thresh-in-go-od-bytes"

    // PM in 802_1 Q threshold
    EtherThreshold_thresh_in_802_1q_frame_s EtherThreshold = "thresh-in-802-1q-frame-s"

    // PM in pkts 1519 max octets threshold
    EtherThreshold_thresh_in_pkts_1519_max_octets EtherThreshold = "thresh-in-pkts-1519-max-octets"

    // PM in good pkts threshold
    EtherThreshold_thresh_in_go_od_pkts EtherThreshold = "thresh-in-go-od-pkts"

    // PM in drop overrun threshold
    EtherThreshold_thresh_in_drop_overrun EtherThreshold = "thresh-in-drop-overrun"

    // PM in drop abort threshold
    EtherThreshold_thresh_in_drop_abort EtherThreshold = "thresh-in-drop-abort"

    // PM in drop invalid vlan threshold
    EtherThreshold_thresh_in_drop_invalid_vlan EtherThreshold = "thresh-in-drop-invalid-vlan"

    // PM in drop invalid DMAC threshold
    EtherThreshold_thresh_in_drop_invalid_dmac EtherThreshold = "thresh-in-drop-invalid-dmac"

    // PM in drop invalid encap threshold
    EtherThreshold_thresh_in_drop_invalid_encap EtherThreshold = "thresh-in-drop-invalid-encap"

    // PM in drop other threshold
    EtherThreshold_thresh_in_drop_other EtherThreshold = "thresh-in-drop-other"

    // PM in MIB giant threshold
    EtherThreshold_thresh_in_mib_giant EtherThreshold = "thresh-in-mib-giant"

    // PM in MIB jabber threshold
    EtherThreshold_thresh_in_mib_jabber EtherThreshold = "thresh-in-mib-jabber"

    // PM in MIB CRC threshold
    EtherThreshold_thresh_in_mib_crc EtherThreshold = "thresh-in-mib-crc"

    // PM in error collisions threshold
    EtherThreshold_thresh_in_error_collision_s EtherThreshold = "thresh-in-error-collision-s"

    // PM in error symbol threshold
    EtherThreshold_thresh_in_error_symbol EtherThreshold = "thresh-in-error-symbol"

    // PM out good bytes threshold
    EtherThreshold_thresh_out_go_od_bytes EtherThreshold = "thresh-out-go-od-bytes"

    // PM out 802_1 Q threshold
    EtherThreshold_thresh_out_802_1q_frame_s EtherThreshold = "thresh-out-802-1q-frame-s"

    // PM out pause frame threshold
    EtherThreshold_thresh_out_pause_frame_s EtherThreshold = "thresh-out-pause-frame-s"

    // PM out pkts 1519 max octets threshold
    EtherThreshold_thresh_out_pkts_1519_max_octets EtherThreshold = "thresh-out-pkts-1519-max-octets"

    // PM out good pkts threshold
    EtherThreshold_thresh_out_go_od_pkts EtherThreshold = "thresh-out-go-od-pkts"

    // PM out drop underrun threshold
    EtherThreshold_thresh_out_drop_under_run EtherThreshold = "thresh-out-drop-under-run"

    // PM out drop abort threshold
    EtherThreshold_thresh_out_drop_abort EtherThreshold = "thresh-out-drop-abort"

    // PM out drop other threshold
    EtherThreshold_thresh_out_drop_other EtherThreshold = "thresh-out-drop-other"

    // PM out error other threshold
    EtherThreshold_thresh_out_error_other EtherThreshold = "thresh-out-error-other"

    // PM in error giant threshold
    EtherThreshold_thresh_in_error_giant EtherThreshold = "thresh-in-error-giant"

    // PM in error runt threshold
    EtherThreshold_thresh_in_error_runt EtherThreshold = "thresh-in-error-runt"

    // PM in error jabber threshold
    EtherThreshold_thresh_in_error_jabbers EtherThreshold = "thresh-in-error-jabbers"

    // PM in error fragments threshold
    EtherThreshold_thresh_in_error_fragments EtherThreshold = "thresh-in-error-fragments"

    // PM in error other threshold
    EtherThreshold_thresh_in_error_other EtherThreshold = "thresh-in-error-other"

    // PM in pkt 64 octet threshold
    EtherThreshold_thresh_in_pkt_64_octet EtherThreshold = "thresh-in-pkt-64-octet"

    // PM in pkts 65_127 octets threshold
    EtherThreshold_thresh_in_pkts_65_127octets EtherThreshold = "thresh-in-pkts-65-127octets"

    // PM in pkts 128_255 octets threshold
    EtherThreshold_thresh_in_pkts_128_255_octets EtherThreshold = "thresh-in-pkts-128-255-octets"

    // PM in pkts 256_511 octets threshold
    EtherThreshold_thresh_in_pkts_256_511_octets EtherThreshold = "thresh-in-pkts-256-511-octets"

    // PM in pkts 512_1023 octets threshold
    EtherThreshold_thresh_in_pkts_512_1023_octets EtherThreshold = "thresh-in-pkts-512-1023-octets"

    // PM in pkts 1024_1058 octets threshold
    EtherThreshold_thresh_in_pkts_1024_1518_octets EtherThreshold = "thresh-in-pkts-1024-1518-octets"

    // PM out pkt 64 octet threshold
    EtherThreshold_thresh_out_pkt_64_octet EtherThreshold = "thresh-out-pkt-64-octet"

    // PM out pkts 65_127 octets threshold
    EtherThreshold_thresh_out_pkts_65_127octets EtherThreshold = "thresh-out-pkts-65-127octets"

    // PM out pkts 128_255 octets threshold
    EtherThreshold_thresh_out_pkts_128_255_octets EtherThreshold = "thresh-out-pkts-128-255-octets"

    // PM out pkts 256_511 octets threshold
    EtherThreshold_thresh_out_pkts_256_511_octets EtherThreshold = "thresh-out-pkts-256-511-octets"

    // PM out pkts 512_1023 octets threshold
    EtherThreshold_thresh_out_pkts_512_1023_octets EtherThreshold = "thresh-out-pkts-512-1023-octets"

    // PM out pkts 1024_1058 octets threshold
    EtherThreshold_thresh_out_pkts_1024_1518_octets EtherThreshold = "thresh-out-pkts-1024-1518-octets"

    // PM tx undersized pkt threshold
    EtherThreshold_thresh_tx_under_size_d_pkt EtherThreshold = "thresh-tx-under-size-d-pkt"

    // PM tx oversized pkt threshold
    EtherThreshold_thresh_tx_over_size_d_pkt EtherThreshold = "thresh-tx-over-size-d-pkt"

    // PM tx fragments threshold
    EtherThreshold_thresh_tx_fragments EtherThreshold = "thresh-tx-fragments"

    // PM tx jabber threshold
    EtherThreshold_thresh_tx_jabber EtherThreshold = "thresh-tx-jabber"

    // PM tx bad fcs threshold
    EtherThreshold_thresh_tx_bad_fcs EtherThreshold = "thresh-tx-bad-fcs"
)

// PathThreshold represents Path threshold
type PathThreshold string

const (
    // PM CV threshold
    PathThreshold_thresh_cv PathThreshold = "thresh-cv"

    // PM ES threshold
    PathThreshold_thresh_es PathThreshold = "thresh-es"

    // PM SES threshold
    PathThreshold_thresh_ses PathThreshold = "thresh-ses"

    // PM UAS threshold
    PathThreshold_thresh_uas PathThreshold = "thresh-uas"
)

// StsThreshold represents Sts threshold
type StsThreshold string

const (
    // PM CV threshold
    StsThreshold_thresh_cv StsThreshold = "thresh-cv"

    // PM ES threshold
    StsThreshold_thresh_es StsThreshold = "thresh-es"

    // PM SES threshold
    StsThreshold_thresh_ses StsThreshold = "thresh-ses"

    // PM UAS threshold
    StsThreshold_thresh_uas StsThreshold = "thresh-uas"
)

// SecytxReport represents Secytx report
type SecytxReport string

const (
    // PM SECYTX tx OutPktsProtected report
    SecytxReport_report_tx_outpktsprotected SecytxReport = "report-tx-outpktsprotected"

    // PM SECYTX tx OutPktsEncrypted report
    SecytxReport_report_tx_outpktsencrypted SecytxReport = "report-tx-outpktsencrypted"

    // PM SECYTX tx OutOctetsProtected report
    SecytxReport_report_tx_outoctetsprotected SecytxReport = "report-tx-outoctetsprotected"

    // PM SECYTX tx OutOctetsEncrypted report
    SecytxReport_report_tx_outoctetsencrypted SecytxReport = "report-tx-outoctetsencrypted"

    // PM SECYTX tx OutPktsTooLong report
    SecytxReport_report_tx_outpkts_too_long SecytxReport = "report-tx-outpkts-too-long"
)

// OcnReport represents Ocn report
type OcnReport string

const (
    // PM SEFSS threshold
    OcnReport_report_sefss OcnReport = "report-sefss"

    // PM CVS threshold
    OcnReport_report_cvs OcnReport = "report-cvs"

    // PM ESS threshold
    OcnReport_report_ess OcnReport = "report-ess"

    // PM SESS threshold
    OcnReport_report_sess OcnReport = "report-sess"

    // PM CVL-NE threshold
    OcnReport_report_cvl_ne OcnReport = "report-cvl-ne"

    // PM ESL-NE threshold
    OcnReport_report_esl_ne OcnReport = "report-esl-ne"

    // PM SESL-NE threshold
    OcnReport_report_sesl_ne OcnReport = "report-sesl-ne"

    // PM UASL-NE threshold
    OcnReport_report_uasl_ne OcnReport = "report-uasl-ne"

    // PM FCL-NE threshold
    OcnReport_report_fcl_ne OcnReport = "report-fcl-ne"

    // PM FCL_FE threshold
    OcnReport_report_fcl_fe OcnReport = "report-fcl-fe"

    // PM CVL-FE threshold
    OcnReport_report_cvl_fe OcnReport = "report-cvl-fe"

    // PM ESL_FE threshold
    OcnReport_report_esl_fe OcnReport = "report-esl-fe"

    // PM SESL_FE threshold
    OcnReport_report_sesl_fe OcnReport = "report-sesl-fe"

    // PM UASL_FEthreshold
    OcnReport_report_uasl_fe OcnReport = "report-uasl-fe"
)

// HoVcReport represents Ho vc report
type HoVcReport string

const (
    // PM EB report
    HoVcReport_report_eb HoVcReport = "report-eb"

    // PM ES report
    HoVcReport_report_es HoVcReport = "report-es"

    // PM ESR report
    HoVcReport_report_esr HoVcReport = "report-esr"

    // PM SES report
    HoVcReport_report_ses HoVcReport = "report-ses"

    // PM SESR report
    HoVcReport_report_sesr HoVcReport = "report-sesr"

    // PM BBE report
    HoVcReport_report_bbe HoVcReport = "report-bbe"

    // PM BBER report
    HoVcReport_report_bber HoVcReport = "report-bber"

    // PM UASS report
    HoVcReport_report_uass HoVcReport = "report-uass"
)

// OtnThreshold represents Otn threshold
type OtnThreshold string

const (
    // PM Otn es sm ne threshold
    OtnThreshold_thresh_es_sm_ne OtnThreshold = "thresh-es-sm-ne"

    // PM Otn ses sm ne threshold
    OtnThreshold_thresh_ses_sm_ne OtnThreshold = "thresh-ses-sm-ne"

    // PM Otn uas sm ne threshold
    OtnThreshold_thresh_uas_sm_ne OtnThreshold = "thresh-uas-sm-ne"

    // PM Otn bbe sm ne threshold
    OtnThreshold_thresh_bbe_sm_ne OtnThreshold = "thresh-bbe-sm-ne"

    // PM Otn fc sm ne threshold
    OtnThreshold_thresh_fc_sm_ne OtnThreshold = "thresh-fc-sm-ne"

    // PM Otn esr sm ne threshold
    OtnThreshold_thresh_esr_sm_ne OtnThreshold = "thresh-esr-sm-ne"

    // PM Otn sesr sm ne threshold
    OtnThreshold_thresh_sesr_sm_ne OtnThreshold = "thresh-sesr-sm-ne"

    // PM Otn bber sm ne threshold
    OtnThreshold_thresh_bber_sm_ne OtnThreshold = "thresh-bber-sm-ne"

    // PM Otn es pm ne threshold
    OtnThreshold_thresh_es_pm_ne OtnThreshold = "thresh-es-pm-ne"

    // PM Otn ses pm ne threshold
    OtnThreshold_thresh_ses_pm_ne OtnThreshold = "thresh-ses-pm-ne"

    // PM Otn uas pm ne threshold
    OtnThreshold_thresh_uas_pm_ne OtnThreshold = "thresh-uas-pm-ne"

    // PM Otn bbe pm ne threshold
    OtnThreshold_thresh_bbe_pm_ne OtnThreshold = "thresh-bbe-pm-ne"

    // PM Otn fc pm ne threshold
    OtnThreshold_thresh_fc_pm_ne OtnThreshold = "thresh-fc-pm-ne"

    // PM Otn esr pm ne threshold
    OtnThreshold_thresh_esr_pm_ne OtnThreshold = "thresh-esr-pm-ne"

    // PM Otn sesr pm ne threshold
    OtnThreshold_thresh_sesr_pm_ne OtnThreshold = "thresh-sesr-pm-ne"

    // PM Otn bber pm ne threshold
    OtnThreshold_thresh_bber_pm_ne OtnThreshold = "thresh-bber-pm-ne"

    // PM Otn es sm fe threshold
    OtnThreshold_thresh_es_sm_fe OtnThreshold = "thresh-es-sm-fe"

    // PM Otn ses sm fe threshold
    OtnThreshold_thresh_ses_sm_fe OtnThreshold = "thresh-ses-sm-fe"

    // PM Otn uas sm fe threshold
    OtnThreshold_thresh_uas_sm_fe OtnThreshold = "thresh-uas-sm-fe"

    // PM Otn bbe sm fe threshold
    OtnThreshold_thresh_bbe_sm_fe OtnThreshold = "thresh-bbe-sm-fe"

    // PM Otn fc sm fe threshold
    OtnThreshold_thresh_fc_sm_fe OtnThreshold = "thresh-fc-sm-fe"

    // PM Otn esr sm fe threshold
    OtnThreshold_thresh_esr_sm_fe OtnThreshold = "thresh-esr-sm-fe"

    // PM Otn sesr sm fe threshold
    OtnThreshold_thresh_sesr_sm_fe OtnThreshold = "thresh-sesr-sm-fe"

    // PM Otn bber sm fe threshold
    OtnThreshold_thresh_bber_sm_fe OtnThreshold = "thresh-bber-sm-fe"

    // PM Otn es pm fe threshold
    OtnThreshold_thresh_es_pm_fe OtnThreshold = "thresh-es-pm-fe"

    // PM Otn ses pm fe threshold
    OtnThreshold_thresh_ses_pm_fe OtnThreshold = "thresh-ses-pm-fe"

    // PM Otn uas pm fe threshold
    OtnThreshold_thresh_uas_pm_fe OtnThreshold = "thresh-uas-pm-fe"

    // PM Otn bbe pm fe threshold
    OtnThreshold_thresh_bbe_pm_fe OtnThreshold = "thresh-bbe-pm-fe"

    // PM Otn fc pm fe threshold
    OtnThreshold_thresh_fc_pm_fe OtnThreshold = "thresh-fc-pm-fe"

    // PM Otn esr pm fe threshold
    OtnThreshold_thresh_esr_pm_fe OtnThreshold = "thresh-esr-pm-fe"

    // PM Otn sesr pm fe threshold
    OtnThreshold_thresh_sesr_pm_fe OtnThreshold = "thresh-sesr-pm-fe"

    // PM Otn bber pm fe threshold
    OtnThreshold_thresh_bber_pm_fe OtnThreshold = "thresh-bber-pm-fe"
)

// OpticsThreshold represents Optics threshold
type OpticsThreshold string

const (
    // PM Optics opt min threshold in dbm or uW
    OpticsThreshold_thresh_opt_min OpticsThreshold = "thresh-opt-min"

    // PM Optics opr min threshold in dbm or uW
    OpticsThreshold_thresh_opr_min OpticsThreshold = "thresh-opr-min"

    // PM Optics lbc min threshold
    OpticsThreshold_thresh_lbc_min OpticsThreshold = "thresh-lbc-min"

    // PM Optics lbcpc min threshold
    OpticsThreshold_thresh_lbc_pc_min OpticsThreshold = "thresh-lbc-pc-min"

    // PM Optics cd min threshold
    OpticsThreshold_thresh_cd_min OpticsThreshold = "thresh-cd-min"

    // PM Optics dgd min threshold
    OpticsThreshold_thresh_dgd_min OpticsThreshold = "thresh-dgd-min"

    // PM Optics sopmd min threshold
    OpticsThreshold_thresh_pmd_min OpticsThreshold = "thresh-pmd-min"

    // PM Optics osnr min threshold
    OpticsThreshold_thresh_osnr_min OpticsThreshold = "thresh-osnr-min"

    // PM Optics pdl min threshold
    OpticsThreshold_thresh_pdl_min OpticsThreshold = "thresh-pdl-min"

    // PM Optics pcr min threshold
    OpticsThreshold_thresh_pcr_min OpticsThreshold = "thresh-pcr-min"

    // PM Optics pn min threshold
    OpticsThreshold_thresh_pn_min OpticsThreshold = "thresh-pn-min"

    // PM Optics rx sig pow min threshold
    OpticsThreshold_thresh_rx_sig_pow_min OpticsThreshold = "thresh-rx-sig-pow-min"

    // PM Optics low sig freq off min threshold
    OpticsThreshold_thresh_low_sig_freq_off_min OpticsThreshold = "thresh-low-sig-freq-off-min"

    // PM Optics ampli gain min threshold
    OpticsThreshold_thresh_ampli_gain_min OpticsThreshold = "thresh-ampli-gain-min"

    // PM Optics ampli gain tilt min threshold
    OpticsThreshold_thresh_ampli_gain_tilt_min OpticsThreshold = "thresh-ampli-gain-tilt-min"

    // PM Optics opt max threshold in dbm or uW
    OpticsThreshold_thresh_opt_max OpticsThreshold = "thresh-opt-max"

    // PM Optics opr max threshold in dbm or uW
    OpticsThreshold_thresh_opr_max OpticsThreshold = "thresh-opr-max"

    // PM Optics lbc max threshold
    OpticsThreshold_thresh_lbc_max OpticsThreshold = "thresh-lbc-max"

    // PM Optics lbcpc max threshold
    OpticsThreshold_thresh_lbc_pc_max OpticsThreshold = "thresh-lbc-pc-max"

    // PM Optics cd max threshold
    OpticsThreshold_thresh_cd_max OpticsThreshold = "thresh-cd-max"

    // PM Optics dgd max threshold
    OpticsThreshold_thresh_dgd_max OpticsThreshold = "thresh-dgd-max"

    // PM Optics sopmd max threshold
    OpticsThreshold_thresh_pmd_max OpticsThreshold = "thresh-pmd-max"

    // PM Optics osnr max threshold
    OpticsThreshold_thresh_osnr_max OpticsThreshold = "thresh-osnr-max"

    // PM Optics pdl max threshold
    OpticsThreshold_thresh_pdl_max OpticsThreshold = "thresh-pdl-max"

    // PM Optics pcr max threshold
    OpticsThreshold_thresh_pcr_max OpticsThreshold = "thresh-pcr-max"

    // PM Optics pn max threshold
    OpticsThreshold_thresh_pn_max OpticsThreshold = "thresh-pn-max"

    // PM Optics rx sig pow max threshold
    OpticsThreshold_thresh_rx_sig_pow_max OpticsThreshold = "thresh-rx-sig-pow-max"

    // PM Optics low sig freq off max threshold
    OpticsThreshold_thresh_low_sig_freq_off_max OpticsThreshold = "thresh-low-sig-freq-off-max"

    // PM Optics ampli gain max threshold
    OpticsThreshold_thresh_ampli_gain_max OpticsThreshold = "thresh-ampli-gain-max"

    // PM Optics ampli gain tilt max threshold
    OpticsThreshold_thresh_ampli_gain_tilt_max OpticsThreshold = "thresh-ampli-gain-tilt-max"
)

// OcnThreshold represents Ocn threshold
type OcnThreshold string

const (
    // PM SEFSS threshold
    OcnThreshold_thresh_sefss OcnThreshold = "thresh-sefss"

    // PM CVS threshold
    OcnThreshold_thresh_cvs OcnThreshold = "thresh-cvs"

    // PM ESS threshold
    OcnThreshold_thresh_ess OcnThreshold = "thresh-ess"

    // PM SESS threshold
    OcnThreshold_thresh_sess OcnThreshold = "thresh-sess"

    // PM CVL-NE threshold
    OcnThreshold_thresh_cvl_ne OcnThreshold = "thresh-cvl-ne"

    // PM ESL-NE threshold
    OcnThreshold_thresh_esl_ne OcnThreshold = "thresh-esl-ne"

    // PM SESL-NE threshold
    OcnThreshold_thresh_sesl_ne OcnThreshold = "thresh-sesl-ne"

    // PM UASL-NE threshold
    OcnThreshold_thresh_uasl_ne OcnThreshold = "thresh-uasl-ne"

    // PM FCL-NE threshold
    OcnThreshold_thresh_fcl_ne OcnThreshold = "thresh-fcl-ne"

    // PM FCL_FE threshold
    OcnThreshold_thresh_fcl_fe OcnThreshold = "thresh-fcl-fe"

    // PM CVL-FE threshold
    OcnThreshold_thresh_cvl_fe OcnThreshold = "thresh-cvl-fe"

    // PM ESL_FE threshold
    OcnThreshold_thresh_esl_fe OcnThreshold = "thresh-esl-fe"

    // PM SESL_FE threshold
    OcnThreshold_thresh_sesl_fe OcnThreshold = "thresh-sesl-fe"

    // PM UASL_FEthreshold
    OcnThreshold_thresh_uasl_fe OcnThreshold = "thresh-uasl-fe"
)

// OpticsReport represents Optics report
type OpticsReport string

const (
    // PM Optics opt min report
    OpticsReport_report_opt_min OpticsReport = "report-opt-min"

    // PM Optics opr min report
    OpticsReport_report_opr_min OpticsReport = "report-opr-min"

    // PM Optics lbc min report
    OpticsReport_report_lbc_min OpticsReport = "report-lbc-min"

    // PM Optics lbcpc min report
    OpticsReport_report_lbc_pc_min OpticsReport = "report-lbc-pc-min"

    // PM Optics cd min report
    OpticsReport_report_cd_min OpticsReport = "report-cd-min"

    // PM Optics dgd min report
    OpticsReport_report_dgd_min OpticsReport = "report-dgd-min"

    // PM Optics sopmd min report
    OpticsReport_report_pmd_min OpticsReport = "report-pmd-min"

    // PM Optics osnr min report
    OpticsReport_report_osnr_min OpticsReport = "report-osnr-min"

    // PM Optics pdl min report
    OpticsReport_report_pdl_min OpticsReport = "report-pdl-min"

    // PM Optics pcr min report
    OpticsReport_report_pcr_min OpticsReport = "report-pcr-min"

    // PM Optics pn min report
    OpticsReport_report_pn_min OpticsReport = "report-pn-min"

    // PM Optics rx sig pow min report
    OpticsReport_report_rx_sig_pow_min OpticsReport = "report-rx-sig-pow-min"

    // PM Optics low sig freq off min report
    OpticsReport_report_low_sig_freq_off_min OpticsReport = "report-low-sig-freq-off-min"

    // PM Optics ampli gain min report
    OpticsReport_report_ampli_gain_min OpticsReport = "report-ampli-gain-min"

    // PM Optics ampli gain tilt min report
    OpticsReport_report_ampli_gain_tilt_min OpticsReport = "report-ampli-gain-tilt-min"

    // PM Optics opt max report
    OpticsReport_report_opt_max OpticsReport = "report-opt-max"

    // PM Optics opr max report
    OpticsReport_report_opr_max OpticsReport = "report-opr-max"

    // PM Optics lbc max report
    OpticsReport_report_lbc_max OpticsReport = "report-lbc-max"

    // PM Optics lbcpc max report
    OpticsReport_report_lbc_pc_max OpticsReport = "report-lbc-pc-max"

    // PM Optics cd max report
    OpticsReport_report_cd_max OpticsReport = "report-cd-max"

    // PM Optics dgd max report
    OpticsReport_report_dgd_max OpticsReport = "report-dgd-max"

    // PM Optics sopmd max report
    OpticsReport_report_pmd_max OpticsReport = "report-pmd-max"

    // PM Optics osnr max report
    OpticsReport_report_osnr_max OpticsReport = "report-osnr-max"

    // PM Optics pdl max report
    OpticsReport_report_pdl_max OpticsReport = "report-pdl-max"

    // PM Optics pcr max report
    OpticsReport_report_pcr_max OpticsReport = "report-pcr-max"

    // PM Optics pn max report
    OpticsReport_report_pn_max OpticsReport = "report-pn-max"

    // PM Optics rx sig pow max report
    OpticsReport_report_rx_sig_pow_max OpticsReport = "report-rx-sig-pow-max"

    // PM Optics low sig freq off max report
    OpticsReport_report_low_sig_freq_off_max OpticsReport = "report-low-sig-freq-off-max"

    // PM Optics ampli gain max report
    OpticsReport_report_ampli_gain_max OpticsReport = "report-ampli-gain-max"

    // PM Optics ampli gain tilt max report
    OpticsReport_report_ampli_gain_tilt_max OpticsReport = "report-ampli-gain-tilt-max"
)

// EtherReport represents Ether report
type EtherReport string

const (
    // PM Ether rx pkt report
    EtherReport_report_rx_pkt EtherReport = "report-rx-pkt"

    // PM Ether rx util report
    EtherReport_report_rx_util EtherReport = "report-rx-util"

    // PM Ether tx util report
    EtherReport_report_tx_util EtherReport = "report-tx-util"

    // PM ether stat pkt threshold
    EtherReport_report_stat_pkt EtherReport = "report-stat-pkt"

    // PM Ether octet stat report
    EtherReport_report_octet_stat EtherReport = "report-octet-stat"

    // PM Ether oversize pkt report
    EtherReport_report_over_size_pkt EtherReport = "report-over-size-pkt"

    // PM Ether fcs error report
    EtherReport_report_fcs_err EtherReport = "report-fcs-err"

    // PM Ether long frames report
    EtherReport_report_long_frame_s EtherReport = "report-long-frame-s"

    // PM Ether jabber stats report
    EtherReport_report_jabber_stats EtherReport = "report-jabber-stats"

    // PM Ether 64 octet report
    EtherReport_report_64_octet EtherReport = "report-64-octet"

    // PM Ether 65 to 127 octet report
    EtherReport_report_65_127_octet EtherReport = "report-65-127-octet"

    // PM Ether 128 to 255 octet report
    EtherReport_report_128_255_octet EtherReport = "report-128-255-octet"

    // PM Ether 256 to 511 octet report
    EtherReport_report_256_511_octet EtherReport = "report-256-511-octet"

    // PM Ether 512 to 1023 octet report
    EtherReport_report_512_1023_octet EtherReport = "report-512-1023-octet"

    // PM Ether 1024 to 1518 report
    EtherReport_report_1024_1518_octet EtherReport = "report-1024-1518-octet"

    // PM Ether rx ucast report
    EtherReport_report_in_ucast EtherReport = "report-in-ucast"

    // PM Ether rx mcast report
    EtherReport_report_in_mcast EtherReport = "report-in-mcast"

    // PM Ether rx bcast report
    EtherReport_report_in_bcast EtherReport = "report-in-bcast"

    // PM Ether tx ucast report
    EtherReport_report_out_ucast EtherReport = "report-out-ucast"

    // PM Ether tx mcast report
    EtherReport_report_out_mcast EtherReport = "report-out-mcast"

    // PM Ether tx bcast report
    EtherReport_report_out_bcast EtherReport = "report-out-bcast"

    // PM Ether tx pkt threshold
    EtherReport_report_tx_pkt EtherReport = "report-tx-pkt"

    // PM ether ifIn errors threshold
    EtherReport_report_ifin_error_s EtherReport = "report-ifin-error-s"

    // PM ether ifInOctets threshold
    EtherReport_report_ifin_octets EtherReport = "report-ifin-octets"

    // PM ether stat multicast pkt threshold
    EtherReport_report_ether_stat_multicast_pkt EtherReport = "report-ether-stat-multicast-pkt"

    // PM ether stat broadcast pkt threshold
    EtherReport_report_ether_stat_broadcast_pkt EtherReport = "report-ether-stat-broadcast-pkt"

    // PM ether stat undersized pkt threshold
    EtherReport_report_ether_stat_under_size_d_pkt EtherReport = "report-ether-stat-under-size-d-pkt"

    // PM ether out octets threshold
    EtherReport_report_out_octet EtherReport = "report-out-octet"

    // PM ether in pause frame report
    EtherReport_report_in_pause_frame EtherReport = "report-in-pause-frame"

    // PM in good bytes report
    EtherReport_report_in_go_od_bytes EtherReport = "report-in-go-od-bytes"

    // PM in 802_1 Q report
    EtherReport_report_in_802_1q_frame_s EtherReport = "report-in-802-1q-frame-s"

    // PM in pkts 1519 max octets report
    EtherReport_report_in_pkts_1519_max_octets EtherReport = "report-in-pkts-1519-max-octets"

    // PM in good pkts report
    EtherReport_report_in_go_od_pkts EtherReport = "report-in-go-od-pkts"

    // PM in drop overrun report
    EtherReport_report_in_drop_overrun EtherReport = "report-in-drop-overrun"

    // PM in drop abort report
    EtherReport_report_in_drop_abort EtherReport = "report-in-drop-abort"

    // PM in drop invalid vlan report
    EtherReport_report_in_drop_invalid_vlan EtherReport = "report-in-drop-invalid-vlan"

    // PM in drop invalid DMAC report
    EtherReport_report_in_drop_invalid_dmac EtherReport = "report-in-drop-invalid-dmac"

    // PM in drop invalid encap report
    EtherReport_report_in_drop_invalid_encap EtherReport = "report-in-drop-invalid-encap"

    // PM in drop other report
    EtherReport_report_in_drop_other EtherReport = "report-in-drop-other"

    // PM in MIB giant report
    EtherReport_report_in_mib_giant EtherReport = "report-in-mib-giant"

    // PM in MIB jabber report
    EtherReport_report_in_mib_jabber EtherReport = "report-in-mib-jabber"

    // PM in MIB CRC report
    EtherReport_report_in_mib_crc EtherReport = "report-in-mib-crc"

    // PM in error collisions report
    EtherReport_report_in_error_collision_s EtherReport = "report-in-error-collision-s"

    // PM in error symbol report
    EtherReport_report_in_error_symbol EtherReport = "report-in-error-symbol"

    // PM out good bytes report
    EtherReport_report_out_go_od_bytes EtherReport = "report-out-go-od-bytes"

    // PM out 802_1 Q report
    EtherReport_report_out_802_1q_frame_s EtherReport = "report-out-802-1q-frame-s"

    // PM out pause frame report
    EtherReport_report_out_pause_frame_s EtherReport = "report-out-pause-frame-s"

    // PM out pkts 1519 max octets report
    EtherReport_report_out_pkts_1519_max_octets EtherReport = "report-out-pkts-1519-max-octets"

    // PM out good pkts report
    EtherReport_report_out_go_od_pkts EtherReport = "report-out-go-od-pkts"

    // PM out drop underrun report
    EtherReport_report_out_drop_under_run EtherReport = "report-out-drop-under-run"

    // PM out drop abort report
    EtherReport_report_out_drop_abort EtherReport = "report-out-drop-abort"

    // PM out drop other report
    EtherReport_report_out_drop_other EtherReport = "report-out-drop-other"

    // PM out error other report
    EtherReport_report_out_error_other EtherReport = "report-out-error-other"

    // PM in error giant report
    EtherReport_report_in_error_giant EtherReport = "report-in-error-giant"

    // PM in error runt report
    EtherReport_report_in_error_runt EtherReport = "report-in-error-runt"

    // PM in error jabber report
    EtherReport_report_in_error_jabbers EtherReport = "report-in-error-jabbers"

    // PM in error fragments report
    EtherReport_report_in_error_fragments EtherReport = "report-in-error-fragments"

    // PM in error other report
    EtherReport_report_in_error_other EtherReport = "report-in-error-other"

    // PM in pkt 64 octet report
    EtherReport_report_in_pkt_64_octet EtherReport = "report-in-pkt-64-octet"

    // PM in pkts 65_127 octets report
    EtherReport_report_in_pkts_65_127octets EtherReport = "report-in-pkts-65-127octets"

    // PM in pkts 128_255 octets report
    EtherReport_report_in_pkts_128_255_octets EtherReport = "report-in-pkts-128-255-octets"

    // PM in pkts 256_511 octets report
    EtherReport_report_in_pkts_256_511_octets EtherReport = "report-in-pkts-256-511-octets"

    // PM in pkts 512_1023 octets report
    EtherReport_report_in_pkts_512_1023_octets EtherReport = "report-in-pkts-512-1023-octets"

    // PM in pkts 1024_1058 octets report
    EtherReport_report_in_pkts_1024_1518_octets EtherReport = "report-in-pkts-1024-1518-octets"

    // PM out pkt 64 octet report
    EtherReport_report_out_pkt_64_octet EtherReport = "report-out-pkt-64-octet"

    // PM out pkts 65_127 octets report
    EtherReport_report_out_pkts_65_127octets EtherReport = "report-out-pkts-65-127octets"

    // PM out pkts 128_255 octets report
    EtherReport_report_out_pkts_128_255_octets EtherReport = "report-out-pkts-128-255-octets"

    // PM out pkts 256_511 octets report
    EtherReport_report_out_pkts_256_511_octets EtherReport = "report-out-pkts-256-511-octets"

    // PM out pkts 512_1023 octets report
    EtherReport_report_out_pkts_512_1023_octets EtherReport = "report-out-pkts-512-1023-octets"

    // PM out pkts 1024_1058 octets report
    EtherReport_report_out_pkts_1024_1518_octets EtherReport = "report-out-pkts-1024-1518-octets"

    // PM tx undersized pkt report
    EtherReport_report_tx_under_size_d_pkt EtherReport = "report-tx-under-size-d-pkt"

    // PM tx oversized pkt report
    EtherReport_report_tx_over_size_d_pkt EtherReport = "report-tx-over-size-d-pkt"

    // PM tx fragments report
    EtherReport_report_tx_fragments EtherReport = "report-tx-fragments"

    // PM tx jabber report
    EtherReport_report_tx_jabber EtherReport = "report-tx-jabber"

    // PM tx bad fcs report
    EtherReport_report_tx_bad_fcs EtherReport = "report-tx-bad-fcs"
)

// OtnTcmReport represents Otn tcm report
type OtnTcmReport string

const (
    // PM Otn es TCM fe report
    OtnTcmReport_report_es_tcm_fe OtnTcmReport = "report-es-tcm-fe"

    // PM Otn ses TCM fe report
    OtnTcmReport_report_ses_tcm_fe OtnTcmReport = "report-ses-tcm-fe"

    // PM Otn uas TCM fe report
    OtnTcmReport_report_uas_tcm_fe OtnTcmReport = "report-uas-tcm-fe"

    // PM Otn bbe TCM fe report
    OtnTcmReport_report_bbe_tcm_fe OtnTcmReport = "report-bbe-tcm-fe"

    // PM Otn fc TCM fe report
    OtnTcmReport_report_fc_tcm_fe OtnTcmReport = "report-fc-tcm-fe"

    // PM Otn esr TCM fe report
    OtnTcmReport_report_esr_tcm_fe OtnTcmReport = "report-esr-tcm-fe"

    // PM Otn sesr TCM fe report
    OtnTcmReport_report_sesr_tcm_fe OtnTcmReport = "report-sesr-tcm-fe"

    // PM Otn bber TCM fe report
    OtnTcmReport_report_bber_tcm_fe OtnTcmReport = "report-bber-tcm-fe"

    // PM Otn es TCM ne report
    OtnTcmReport_report_es_tcm_ne OtnTcmReport = "report-es-tcm-ne"

    // PM Otn ses TCM ne report
    OtnTcmReport_report_ses_tcm_ne OtnTcmReport = "report-ses-tcm-ne"

    // PM Otn uas TCM ne report
    OtnTcmReport_report_uas_tcm_ne OtnTcmReport = "report-uas-tcm-ne"

    // PM Otn bbe TCM ne report
    OtnTcmReport_report_bbe_tcm_ne OtnTcmReport = "report-bbe-tcm-ne"

    // PM Otn fc TCM ne report
    OtnTcmReport_report_fc_tcm_ne OtnTcmReport = "report-fc-tcm-ne"

    // PM Otn esr TCM ne report
    OtnTcmReport_report_esr_tcm_ne OtnTcmReport = "report-esr-tcm-ne"

    // PM Otn sesr TCM ne report
    OtnTcmReport_report_sesr_tcm_ne OtnTcmReport = "report-sesr-tcm-ne"

    // PM Otn bber TCM ne report
    OtnTcmReport_report_bber_tcm_ne OtnTcmReport = "report-bber-tcm-ne"
)

// FecThreshold represents Fec threshold
type FecThreshold string

const (
    // PM Fec ec bits threshold
    FecThreshold_thresh_ec_bits FecThreshold = "thresh-ec-bits"

    // PM Fec uc words threshold
    FecThreshold_thresh_uc_words FecThreshold = "thresh-uc-words"

    // PM Fec pre-fe-ber max threshold
    FecThreshold_thresh_pre_fec_ber_max FecThreshold = "thresh-pre-fec-ber-max"

    // PM Fec post-fec-ber max threshold
    FecThreshold_thresh_post_fec_ber_max FecThreshold = "thresh-post-fec-ber-max"

    // PM Fec Q max threshold
    FecThreshold_thresh_q_max FecThreshold = "thresh-q-max"

    // PM Fec uc words max threshold
    FecThreshold_thresh_q_margin_max FecThreshold = "thresh-q-margin-max"

    // PM Fec pre-fe-ber min threshold
    FecThreshold_thresh_pre_fec_ber_min FecThreshold = "thresh-pre-fec-ber-min"

    // PM Fec post-fec-ber min threshold
    FecThreshold_thresh_post_fec_ber_min FecThreshold = "thresh-post-fec-ber-min"

    // PM Fec Q min threshold
    FecThreshold_thresh_q_min FecThreshold = "thresh-q-min"

    // PM Fec uc words min threshold
    FecThreshold_thresh_q_margin_min FecThreshold = "thresh-q-margin-min"
)

// SecyifThreshold represents Secyif threshold
type SecyifThreshold string

const (
    // PM SECYIF if InPktsUntagged thresh
    SecyifThreshold_thresh_if_inpkts_untagged SecyifThreshold = "thresh-if-inpkts-untagged"

    // PM SECYIF if InPktsNoTag thresh
    SecyifThreshold_thresh_if_inpkts_not_ag SecyifThreshold = "thresh-if-inpkts-not-ag"

    // PM SECYIF if InPktsBadTag thresh
    SecyifThreshold_thresh_if_inpktsbadtag SecyifThreshold = "thresh-if-inpktsbadtag"

    // PM SECYIF if InPktsUnknownSCI thresh
    SecyifThreshold_thresh_if_inpktsunkownsci SecyifThreshold = "thresh-if-inpktsunkownsci"

    // PM SECYIF if InPktsNoSCI thresh
    SecyifThreshold_thresh_if_inpktsnosci SecyifThreshold = "thresh-if-inpktsnosci"

    // PM SECYIF if InPktsOverrun thresh
    SecyifThreshold_thresh_if_inpkts_overrun SecyifThreshold = "thresh-if-inpkts-overrun"

    // PM SECYIF if InOctetsValidated thresh
    SecyifThreshold_thresh_if_inoctets_validate_d SecyifThreshold = "thresh-if-inoctets-validate-d"

    // PM SECYIF if InOctetsDecrypted thresh
    SecyifThreshold_thresh_if_inoctetsdecrypted SecyifThreshold = "thresh-if-inoctetsdecrypted"

    // PM SECYIF if OutPktsUntagged thresh
    SecyifThreshold_thresh_if_outpkts_untagged SecyifThreshold = "thresh-if-outpkts-untagged"

    // PM SECYIF if OutPktsTooLong thresh
    SecyifThreshold_thresh_if_thresh_outpkts_too_long SecyifThreshold = "thresh-if-thresh-outpkts-too-long"

    // PM SECYIF if OutOctetsProtected thresh
    SecyifThreshold_thresh_if_outoctetsprotected SecyifThreshold = "thresh-if-outoctetsprotected"

    // PM SECYIF if OutOctetsEncrypted thresh
    SecyifThreshold_thresh_if_outoctetsencrypted SecyifThreshold = "thresh-if-outoctetsencrypted"
)

// StsReport represents Sts report
type StsReport string

const (
    // PM CV threshold
    StsReport_report_cv StsReport = "report-cv"

    // PM ES threshold
    StsReport_report_es StsReport = "report-es"

    // PM SES threshold
    StsReport_report_ses StsReport = "report-ses"

    // PM UAS threshold
    StsReport_report_uas StsReport = "report-uas"
)

// StmThreshold represents Stm threshold
type StmThreshold string

const (
    // PM EBS threshold
    StmThreshold_thresh_ebs StmThreshold = "thresh-ebs"

    // PM ESS threshold
    StmThreshold_thresh_ess StmThreshold = "thresh-ess"

    // PM ESRS threshold
    StmThreshold_thresh_esrs StmThreshold = "thresh-esrs"

    // PM SES threshold
    StmThreshold_thresh_sess StmThreshold = "thresh-sess"

    // PM SESR threshold
    StmThreshold_thresh_sesrs StmThreshold = "thresh-sesrs"

    // PM BBES threshold
    StmThreshold_thresh_bbes StmThreshold = "thresh-bbes"

    // PM BBESR threshold
    StmThreshold_thresh_bbesr StmThreshold = "thresh-bbesr"

    // PM UASS threshold
    StmThreshold_thresh_uass StmThreshold = "thresh-uass"

    // PM EBLNE threshold
    StmThreshold_thresh_ebl_ne StmThreshold = "thresh-ebl-ne"

    // PM ESLNE threshold
    StmThreshold_thresh_esl_ne StmThreshold = "thresh-esl-ne"

    // PM ESLRNE threshold
    StmThreshold_thresh_eslr_ne StmThreshold = "thresh-eslr-ne"

    // PM SESL threshold
    StmThreshold_thresh_sesl_ne StmThreshold = "thresh-sesl-ne"

    // PM SESRL threshold
    StmThreshold_thresh_sesrl_ne StmThreshold = "thresh-sesrl-ne"

    // PM BBERLNE threshold
    StmThreshold_thresh_bbel_ne StmThreshold = "thresh-bbel-ne"

    // PM BBERLNE threshold
    StmThreshold_thresh_bberl_ne StmThreshold = "thresh-bberl-ne"

    // PM UASNE threshold
    StmThreshold_thresh_uasl_ne StmThreshold = "thresh-uasl-ne"

    // PM EBFE threshold
    StmThreshold_thresh_ebl_fe StmThreshold = "thresh-ebl-fe"

    // PM ESFE threshold
    StmThreshold_thresh_esl_fe StmThreshold = "thresh-esl-fe"

    // PM EBFE threshold
    StmThreshold_thresh_esrl_fe StmThreshold = "thresh-esrl-fe"

    // PM SESFE threshold
    StmThreshold_thresh_sesl_fe StmThreshold = "thresh-sesl-fe"

    // PM SESRLFE threshold
    StmThreshold_thresh_sesrl_fe StmThreshold = "thresh-sesrl-fe"

    // PM BBEL threshold
    StmThreshold_thresh_bbel_fe StmThreshold = "thresh-bbel-fe"

    // PM BBELFE threshold
    StmThreshold_thresh_bberl_fe StmThreshold = "thresh-bberl-fe"

    // PM UASLFE threshold
    StmThreshold_thresh_uasl_fe StmThreshold = "thresh-uasl-fe"
)

// OtnTcmThreshold represents Otn tcm threshold
type OtnTcmThreshold string

const (
    // PM Otn es TCM fe threshold
    OtnTcmThreshold_thresh_es_tcm_fe OtnTcmThreshold = "thresh-es-tcm-fe"

    // PM Otn ses TCM fe threshold
    OtnTcmThreshold_thresh_ses_tcm_fe OtnTcmThreshold = "thresh-ses-tcm-fe"

    // PM Otn uas TCM fe threshold
    OtnTcmThreshold_thresh_uas_tcm_fe OtnTcmThreshold = "thresh-uas-tcm-fe"

    // PM Otn bbe TCM fe threshold
    OtnTcmThreshold_thresh_bbe_tcm_fe OtnTcmThreshold = "thresh-bbe-tcm-fe"

    // PM Otn fc TCM fe threshold
    OtnTcmThreshold_thresh_fc_tcm_fe OtnTcmThreshold = "thresh-fc-tcm-fe"

    // PM Otn esr TCM fe threshold
    OtnTcmThreshold_thresh_esr_tcm_fe OtnTcmThreshold = "thresh-esr-tcm-fe"

    // PM Otn sesr TCM fe threshold
    OtnTcmThreshold_thresh_sesr_tcm_fe OtnTcmThreshold = "thresh-sesr-tcm-fe"

    // PM Otn bber TCM fe threshold
    OtnTcmThreshold_thresh_bber_tcm_fe OtnTcmThreshold = "thresh-bber-tcm-fe"

    // PM Otn es TCM ne threshold
    OtnTcmThreshold_thresh_es_tcm_ne OtnTcmThreshold = "thresh-es-tcm-ne"

    // PM Otn ses TCM ne threshold
    OtnTcmThreshold_thresh_ses_tcm_ne OtnTcmThreshold = "thresh-ses-tcm-ne"

    // PM Otn uas TCM ne threshold
    OtnTcmThreshold_thresh_uas_tcm_ne OtnTcmThreshold = "thresh-uas-tcm-ne"

    // PM Otn bbe TCM ne threshold
    OtnTcmThreshold_thresh_bbe_tcm_ne OtnTcmThreshold = "thresh-bbe-tcm-ne"

    // PM Otn fc TCM ne threshold
    OtnTcmThreshold_thresh_fc_tcm_ne OtnTcmThreshold = "thresh-fc-tcm-ne"

    // PM Otn esr TCM ne threshold
    OtnTcmThreshold_thresh_esr_tcm_ne OtnTcmThreshold = "thresh-esr-tcm-ne"

    // PM Otn sesr TCM ne threshold
    OtnTcmThreshold_thresh_sesr_tcm_ne OtnTcmThreshold = "thresh-sesr-tcm-ne"

    // PM Otn bber TCM ne threshold
    OtnTcmThreshold_thresh_bber_tcm_ne OtnTcmThreshold = "thresh-bber-tcm-ne"
)

// Report represents Report
type Report string

const (
    // Performance Monitoring Disabled
    Report_false Report = "false"

    // Performance Monitoring Enabled
    Report_true Report = "true"
)

// OtnReport represents Otn report
type OtnReport string

const (
    // PM Otn es sm ne report
    OtnReport_report_es_sm_ne OtnReport = "report-es-sm-ne"

    // PM Otn ses sm ne report
    OtnReport_report_ses_sm_ne OtnReport = "report-ses-sm-ne"

    // PM Otn uas sm ne report
    OtnReport_report_uas_sm_ne OtnReport = "report-uas-sm-ne"

    // PM Otn bbe sm ne report
    OtnReport_report_bbe_sm_ne OtnReport = "report-bbe-sm-ne"

    // PM Otn fc sm ne report
    OtnReport_report_fc_sm_ne OtnReport = "report-fc-sm-ne"

    // PM Otn esr sm ne report
    OtnReport_report_esr_sm_ne OtnReport = "report-esr-sm-ne"

    // PM Otn sesr sm ne report
    OtnReport_report_sesr_sm_ne OtnReport = "report-sesr-sm-ne"

    // PM Otn bber sm ne report
    OtnReport_report_bber_sm_ne OtnReport = "report-bber-sm-ne"

    // PM Otn es pm ne report
    OtnReport_report_es_pm_ne OtnReport = "report-es-pm-ne"

    // PM Otn ses pm ne report
    OtnReport_report_ses_pm_ne OtnReport = "report-ses-pm-ne"

    // PM Otn uas pm ne report
    OtnReport_report_uas_pm_ne OtnReport = "report-uas-pm-ne"

    // PM Otn bbe pm ne report
    OtnReport_report_bbe_pm_ne OtnReport = "report-bbe-pm-ne"

    // PM Otn fc pm ne report
    OtnReport_report_fc_pm_ne OtnReport = "report-fc-pm-ne"

    // PM Otn esr pm ne report
    OtnReport_report_esr_pm_ne OtnReport = "report-esr-pm-ne"

    // PM Otn sesr pm ne report
    OtnReport_report_sesr_pm_ne OtnReport = "report-sesr-pm-ne"

    // PM Otn bber pm ne report
    OtnReport_report_bber_pm_ne OtnReport = "report-bber-pm-ne"

    // PM Otn es sm fe report
    OtnReport_report_es_sm_fe OtnReport = "report-es-sm-fe"

    // PM Otn ses sm fe report
    OtnReport_report_ses_sm_fe OtnReport = "report-ses-sm-fe"

    // PM Otn uas sm fe report
    OtnReport_report_uas_sm_fe OtnReport = "report-uas-sm-fe"

    // PM Otn bbe sm fe report
    OtnReport_report_bbe_sm_fe OtnReport = "report-bbe-sm-fe"

    // PM Otn fc sm fe report
    OtnReport_report_fc_sm_fe OtnReport = "report-fc-sm-fe"

    // PM Otn esr sm fe report
    OtnReport_report_esr_sm_fe OtnReport = "report-esr-sm-fe"

    // PM Otn sesr sm fe report
    OtnReport_report_sesr_sm_fe OtnReport = "report-sesr-sm-fe"

    // PM Otn bber sm fe report
    OtnReport_report_bber_sm_fe OtnReport = "report-bber-sm-fe"

    // PM Otn es pm fe report
    OtnReport_report_es_pm_fe OtnReport = "report-es-pm-fe"

    // PM Otn ses pm fe report
    OtnReport_report_ses_pm_fe OtnReport = "report-ses-pm-fe"

    // PM Otn uas pm fe report
    OtnReport_report_uas_pm_fe OtnReport = "report-uas-pm-fe"

    // PM Otn bbe pm fe report
    OtnReport_report_bbe_pm_fe OtnReport = "report-bbe-pm-fe"

    // PM Otn fc pm fe report
    OtnReport_report_fc_pm_fe OtnReport = "report-fc-pm-fe"

    // PM Otn esr pm fe report
    OtnReport_report_esr_pm_fe OtnReport = "report-esr-pm-fe"

    // PM Otn sesr pm fe report
    OtnReport_report_sesr_pm_fe OtnReport = "report-sesr-pm-fe"

    // PM Otn bber pm fe report
    OtnReport_report_bber_pm_fe OtnReport = "report-bber-pm-fe"
)

// SecytxThreshold represents Secytx threshold
type SecytxThreshold string

const (
    // PM SECYTX tx OutPktsProtected thresh
    SecytxThreshold_thresh_tx_outpktsprotected SecytxThreshold = "thresh-tx-outpktsprotected"

    // PM SECYTX tx OutPktsEncrypted thresh
    SecytxThreshold_thresh_tx_outpktsencrypted SecytxThreshold = "thresh-tx-outpktsencrypted"

    // PM SECYTX tx OutOctetsProtected thresh
    SecytxThreshold_thresh_tx_outoctetsprotected SecytxThreshold = "thresh-tx-outoctetsprotected"

    // PM SECYTX tx OutOctetsEncrypted thresh
    SecytxThreshold_thresh_tx_outoctetsencrypted SecytxThreshold = "thresh-tx-outoctetsencrypted"

    // PM SECYTX tx OutPktsTooLong thresh
    SecytxThreshold_thresh_tx_outpkts_too_long SecytxThreshold = "thresh-tx-outpkts-too-long"
)

// FecReport represents Fec report
type FecReport string

const (
    // PM Fec ec bits report
    FecReport_report_ec_bits FecReport = "report-ec-bits"

    // PM Fec uc words report
    FecReport_report_uc_words FecReport = "report-uc-words"

    // PM Fec pre fec ber max report
    FecReport_report_pre_fec_ber_max FecReport = "report-pre-fec-ber-max"

    // PM Fec post fec ber max report
    FecReport_report_post_fec_ber_max FecReport = "report-post-fec-ber-max"

    // PM Fec Q max report
    FecReport_report_q_max FecReport = "report-q-max"

    // PM Fec Q_margin max report
    FecReport_report_q_margin_max FecReport = "report-q-margin-max"

    // PM Fec pre fec ber min report
    FecReport_report_pre_fec_ber_min FecReport = "report-pre-fec-ber-min"

    // PM Fec post fec ber min report
    FecReport_report_post_fec_ber_min FecReport = "report-post-fec-ber-min"

    // PM Fec Q min report
    FecReport_report_q_min FecReport = "report-q-min"

    // PM Fec Q_margin min report
    FecReport_report_q_margin_min FecReport = "report-q-margin-min"
)

// PathReport represents Path report
type PathReport string

const (
    // PM CV threshold
    PathReport_report_cv PathReport = "report-cv"

    // PM ES threshold
    PathReport_report_es PathReport = "report-es"

    // PM SES threshold
    PathReport_report_ses PathReport = "report-ses"

    // PM UAS threshold
    PathReport_report_uas PathReport = "report-uas"
)

// SecyrxReport represents Secyrx report
type SecyrxReport string

const (
    // PM SECYRX rx InPktsUnchecked report
    SecyrxReport_report_rx_inpktsun_check_ed SecyrxReport = "report-rx-inpktsun-check-ed"

    // PM SECYRX rx InPktsDelayed report
    SecyrxReport_report_rx_inpkts_delayed SecyrxReport = "report-rx-inpkts-delayed"

    // PM SECYRX rx InPktsLate report
    SecyrxReport_report_rx_inpktslate SecyrxReport = "report-rx-inpktslate"

    // PM SECYRX rx InPktsOK report
    SecyrxReport_report_rx_inpktsok SecyrxReport = "report-rx-inpktsok"

    // PM SECYRX rx InPktsInvalid report
    SecyrxReport_report_rx_inpkts_invalid SecyrxReport = "report-rx-inpkts-invalid"

    // PM SECYRX rx InPktsNotValid report
    SecyrxReport_report_rx_inpkts_not_valid SecyrxReport = "report-rx-inpkts-not-valid"

    // PM SECYRX rx InPktsNotUsingSA sa report
    SecyrxReport_report_rx_inpkts_not_usingsa SecyrxReport = "report-rx-inpkts-not-usingsa"

    // PM SECYRX rx InPktsUnusedSA report
    SecyrxReport_report_rx_inpktsunusedsa SecyrxReport = "report-rx-inpktsunusedsa"

    // PM SECYRX rx InPktsUntaggedHit report
    SecyrxReport_report_rx_inpkts_untagged_hit SecyrxReport = "report-rx-inpkts-untagged-hit"

    // PM SECYRX rx InOctetsValidated report
    SecyrxReport_report_rx_inoctets_validate_d SecyrxReport = "report-rx-inoctets-validate-d"

    // PM SECYRX rx InOctetsDecrypted report
    SecyrxReport_report_rx_inoctetsdecrypted SecyrxReport = "report-rx-inoctetsdecrypted"
)

// SecyrxThreshold represents Secyrx threshold
type SecyrxThreshold string

const (
    // PM SECYRX rx InPktsUnchecked thresh
    SecyrxThreshold_thresh_rx_inpktsun_check_ed SecyrxThreshold = "thresh-rx-inpktsun-check-ed"

    // PM SECYRX rx InPktsDelayed thresh
    SecyrxThreshold_thresh_rx_inpkts_delayed SecyrxThreshold = "thresh-rx-inpkts-delayed"

    // PM SECYRX rx InPktsLate thresh
    SecyrxThreshold_thresh_rx_inpktslate SecyrxThreshold = "thresh-rx-inpktslate"

    // PM SECYRX rx InPktsOK thresh
    SecyrxThreshold_thresh_rx_inpktsok SecyrxThreshold = "thresh-rx-inpktsok"

    // PM SECYRX rx InPktsInvalid thresh
    SecyrxThreshold_thresh_rx_inpkts_invalid SecyrxThreshold = "thresh-rx-inpkts-invalid"

    // PM SECYRX rx InPktsNotValid thresh
    SecyrxThreshold_thresh_rx_inpkts_not_valid SecyrxThreshold = "thresh-rx-inpkts-not-valid"

    // PM SECYRX rx InPktsNotUsingSA thresh
    SecyrxThreshold_thresh_rx_inpkts_not_usingsa SecyrxThreshold = "thresh-rx-inpkts-not-usingsa"

    // PM SECYRX rx InPktsUnusedSA thresh
    SecyrxThreshold_thresh_rx_inpktsunusedsa SecyrxThreshold = "thresh-rx-inpktsunusedsa"

    // PM SECYRX rx InPktsUntaggedHit thresh
    SecyrxThreshold_thresh_rx_inpkts_untagged_hit SecyrxThreshold = "thresh-rx-inpkts-untagged-hit"

    // PM SECYRX rx InOctetsValidated thresh
    SecyrxThreshold_thresh_rx_inoctets_validate_d SecyrxThreshold = "thresh-rx-inoctets-validate-d"

    // PM SECYRX rx InOctetsDecrypted thresh
    SecyrxThreshold_thresh_rx_inoctetsdecrypted SecyrxThreshold = "thresh-rx-inoctetsdecrypted"
)

// GfpThreshold represents Gfp threshold
type GfpThreshold string

const (
    // PM GFP rx bit err threshold
    GfpThreshold_thresh_rx_bit_err GfpThreshold = "thresh-rx-bit-err"

    // PM GFP rx inv type threshold
    GfpThreshold_thresh_rx_inv_typ GfpThreshold = "thresh-rx-inv-typ"

    // PM GFP rx crc threshold
    GfpThreshold_thresh_rx_crc GfpThreshold = "thresh-rx-crc"

    // PM GFP rx lfd threshold
    GfpThreshold_thresh_rx_lfd GfpThreshold = "thresh-rx-lfd"

    // PM GFP rx csf threshold
    GfpThreshold_thresh_rx_csf GfpThreshold = "thresh-rx-csf"
)

// StmReport represents Stm report
type StmReport string

const (
    // PM EBS REPORT
    StmReport_report_ebs StmReport = "report-ebs"

    // PM ESS REPORT
    StmReport_report_ess StmReport = "report-ess"

    // PM ESRS REPORT
    StmReport_report_esrs StmReport = "report-esrs"

    // PM SES REPORT
    StmReport_report_sess StmReport = "report-sess"

    // PM SESR REPORT
    StmReport_report_sesrs StmReport = "report-sesrs"

    // PM BBES REPORT
    StmReport_report_bbes StmReport = "report-bbes"

    // PM BBESR REPORT
    StmReport_report_bbesr StmReport = "report-bbesr"

    // PM UASS REPORT
    StmReport_report_uass StmReport = "report-uass"

    // PM EBLNE REPORT
    StmReport_report_ebl_ne StmReport = "report-ebl-ne"

    // PM ESLNE REPORT
    StmReport_report_esl_ne StmReport = "report-esl-ne"

    // PM ESLRNE REPORT
    StmReport_report_eslr_ne StmReport = "report-eslr-ne"

    // PM SESL REPORT
    StmReport_report_sesl_ne StmReport = "report-sesl-ne"

    // PM SESRL REPORT
    StmReport_report_sesrl_ne StmReport = "report-sesrl-ne"

    // PM BBELNE REPORT
    StmReport_report_bbel_ne StmReport = "report-bbel-ne"

    // PM BBERLNE REPORT
    StmReport_report_bberl_ne StmReport = "report-bberl-ne"

    // PM UASNE REPORT
    StmReport_report_uasl_ne StmReport = "report-uasl-ne"

    // PM EBFE REPORT
    StmReport_report_ebl_fe StmReport = "report-ebl-fe"

    // PM ESFE REPORT
    StmReport_report_esl_fe StmReport = "report-esl-fe"

    // PM EBFE REPORT
    StmReport_report_esrl_fe StmReport = "report-esrl-fe"

    // PM SESFE REPORT
    StmReport_report_sesl_fe StmReport = "report-sesl-fe"

    // PM SESRLFE REPORT
    StmReport_report_sesrl_fe StmReport = "report-sesrl-fe"

    // PM BBELFE REPORT
    StmReport_report_bbel_fe StmReport = "report-bbel-fe"

    // PM ESFE REPORT
    StmReport_report_bberl_fe StmReport = "report-bberl-fe"

    // PM UASLFE REPORT
    StmReport_report_uasl_fe StmReport = "report-uasl-fe"
)

