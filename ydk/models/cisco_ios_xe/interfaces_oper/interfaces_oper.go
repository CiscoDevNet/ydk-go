// This module contains a collection of YANG definitions for
// monitoring the interfaces in a Network Element.
// Copyright (c) 2016-2018 by Cisco Systems, Inc.
// All rights reserved.
package interfaces_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package interfaces_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-interfaces-oper interfaces}", reflect.TypeOf(Interfaces{}))
    ydk.RegisterEntity("Cisco-IOS-XE-interfaces-oper:interfaces", reflect.TypeOf(Interfaces{}))
}

// QosMatchType represents QOS match type
type QosMatchType string

const (
    QosMatchType_qos_match_dscp QosMatchType = "qos-match-dscp"

    QosMatchType_qos_match_src_ip QosMatchType = "qos-match-src-ip"

    QosMatchType_qos_match_dst_ip QosMatchType = "qos-match-dst-ip"

    QosMatchType_qos_match_src_port QosMatchType = "qos-match-src-port"

    QosMatchType_qos_match_dst_port QosMatchType = "qos-match-dst-port"

    QosMatchType_qos_match_proto QosMatchType = "qos-match-proto"
)

// EtherDuplex represents The duplex setting of the interface
type EtherDuplex string

const (
    EtherDuplex_full_duplex EtherDuplex = "full-duplex"

    EtherDuplex_half_duplex EtherDuplex = "half-duplex"

    EtherDuplex_auto_duplex EtherDuplex = "auto-duplex"

    EtherDuplex_unknown_duplex EtherDuplex = "unknown-duplex"
)

// SerialCrc represents The Cyclic Redundancy Code type
type SerialCrc string

const (
    // 32-bit Cyclic Redundancy Code
    SerialCrc_serial_crc32 SerialCrc = "serial-crc32"

    // 16 bit Cyclic Redundancy Code
    SerialCrc_serial_crc16 SerialCrc = "serial-crc16"
)

// ThreshUnit represents Units of threshold
type ThreshUnit string

const (
    ThreshUnit_thresh_units_default ThreshUnit = "thresh-units-default"

    ThreshUnit_thresh_units_bytes ThreshUnit = "thresh-units-bytes"

    ThreshUnit_thresh_units_sec ThreshUnit = "thresh-units-sec"

    ThreshUnit_thresh_units_packets ThreshUnit = "thresh-units-packets"

    ThreshUnit_thresh_units_cells ThreshUnit = "thresh-units-cells"

    ThreshUnit_thresh_units_percent ThreshUnit = "thresh-units-percent"
)

// T1e1LoopbackMode represents Loopback mode type
type T1e1LoopbackMode string

const (
    // No loopback mode
    T1e1LoopbackMode_t1e1_no_loopback T1e1LoopbackMode = "t1e1-no-loopback"

    // Command line interface enforced local loopback
    T1e1LoopbackMode_t1e1_cli_local_loopback T1e1LoopbackMode = "t1e1-cli-local-loopback"

    // Command line interface enforced line local loopback
    T1e1LoopbackMode_t1e1_line_cli_local_loopback T1e1LoopbackMode = "t1e1-line-cli-local-loopback"

    // Command line interface enforced payload local loopback
    T1e1LoopbackMode_t1e1_payload_cli_local_loopback T1e1LoopbackMode = "t1e1-payload-cli-local-loopback"

    // Local line loopback
    T1e1LoopbackMode_t1e1_local_line_loopback T1e1LoopbackMode = "t1e1-local-line-loopback"

    // Local payload loopback
    T1e1LoopbackMode_t1e1_local_payload_loopback T1e1LoopbackMode = "t1e1-local-payload-loopback"

    // Line ANSI FDL remote loopback
    T1e1LoopbackMode_t1e1_local_ansi_fdl_remote_loopback T1e1LoopbackMode = "t1e1-local-ansi-fdl-remote-loopback"

    // Line ATT FDL remote loopback
    T1e1LoopbackMode_t1e1_line_att_fdl_remote_loopback T1e1LoopbackMode = "t1e1-line-att-fdl-remote-loopback"

    // Payload ANSI FDL remote loopback
    T1e1LoopbackMode_t1e1_payload_ansi_fdl_remote_loopback T1e1LoopbackMode = "t1e1-payload-ansi-fdl-remote-loopback"

    // Payload ATT FDL remote loopback
    T1e1LoopbackMode_t1e1_payload_att_fdl_remote_loopback T1e1LoopbackMode = "t1e1-payload-att-fdl-remote-loopback"

    // Line IBOC remote loopback
    T1e1LoopbackMode_t1e1_line_iboc_remote_loopback T1e1LoopbackMode = "t1e1-line-iboc-remote-loopback"

    // Line ANSI FDL local loopback
    T1e1LoopbackMode_t1e1_line_ansi_fdl_local_loopback T1e1LoopbackMode = "t1e1-line-ansi-fdl-local-loopback"

    // Line ATT FDL local loopback
    T1e1LoopbackMode_t1e1_line_att_fdl_local_loopback T1e1LoopbackMode = "t1e1-line-att-fdl-local-loopback"

    // Payload ANSI FDL local loopback
    T1e1LoopbackMode_t1e1_payload_ansi_fdl_local_loopback T1e1LoopbackMode = "t1e1-payload-ansi-fdl-local-loopback"

    // Payload ATT FDL local loopback
    T1e1LoopbackMode_t1e1_payload_att_fdl_local_loopback T1e1LoopbackMode = "t1e1-payload-att-fdl-local-loopback"

    // Line IBOC local loopback
    T1e1LoopbackMode_t1e1_line_iboc_local_loopback T1e1LoopbackMode = "t1e1-line-iboc-local-loopback"
)

// IntfState represents RFC 2863: The Interfaces Group MIB - ifAdminStatus
type IntfState string

const (
    IntfState_if_state_unknown IntfState = "if-state-unknown"

    IntfState_if_state_up IntfState = "if-state-up"

    IntfState_if_state_down IntfState = "if-state-down"

    IntfState_if_state_test IntfState = "if-state-test"
)

// OperState represents RFC 2863: The Interfaces Group MIB - ifOperStatus
type OperState string

const (
    OperState_if_oper_state_invalid OperState = "if-oper-state-invalid"

    OperState_if_oper_state_ready OperState = "if-oper-state-ready"

    OperState_if_oper_state_no_pass OperState = "if-oper-state-no-pass"

    OperState_if_oper_state_test OperState = "if-oper-state-test"

    OperState_if_oper_state_unknown OperState = "if-oper-state-unknown"

    OperState_if_oper_state_dormant OperState = "if-oper-state-dormant"

    OperState_if_oper_state_not_present OperState = "if-oper-state-not-present"

    OperState_if_oper_state_lower_layer_down OperState = "if-oper-state-lower-layer-down"
)

// AggregationType represents defined and managed
type AggregationType string

const (
    // LAG mode is off
    AggregationType_lag_off AggregationType = "lag-off"

    // LAG mode is auto
    AggregationType_lag_auto AggregationType = "lag-auto"

    // LAG mode is active
    AggregationType_lag_active AggregationType = "lag-active"

    // LAG mode is passive
    AggregationType_lag_passive AggregationType = "lag-passive"
)

// SubrateSpeed represents The subrate on a serial interface
type SubrateSpeed string

const (
    // 56 kilobits per second subrate
    SubrateSpeed_dsx1_subrate_56kbps SubrateSpeed = "dsx1-subrate-56kbps"

    // 64 kilobits per second subrate
    SubrateSpeed_dsx1_subrate_64kbps SubrateSpeed = "dsx1-subrate-64kbps"
)

// IetfIntfType represents object in the (updated) definition of MIB-II's ifTable
type IetfIntfType string

const (
    IetfIntfType_iana_iftype_other IetfIntfType = "iana-iftype-other"

    IetfIntfType_iana_iftype_regular1822 IetfIntfType = "iana-iftype-regular1822"

    IetfIntfType_iana_iftype_hdh1822 IetfIntfType = "iana-iftype-hdh1822"

    IetfIntfType_iana_iftype_ddnx25 IetfIntfType = "iana-iftype-ddnx25"

    IetfIntfType_iana_iftype_rfc877x25 IetfIntfType = "iana-iftype-rfc877x25"

    IetfIntfType_iana_iftype_ethernet_csmacd IetfIntfType = "iana-iftype-ethernet-csmacd"

    IetfIntfType_iana_iftype_iso88023_csmacd IetfIntfType = "iana-iftype-iso88023-csmacd"

    IetfIntfType_iana_iftype_iso88024_tokenbus IetfIntfType = "iana-iftype-iso88024-tokenbus"

    IetfIntfType_iana_iftype_iso88025_tokenring IetfIntfType = "iana-iftype-iso88025-tokenring"

    IetfIntfType_iana_iftype_iso88026_man IetfIntfType = "iana-iftype-iso88026-man"

    IetfIntfType_iana_iftype_starlan IetfIntfType = "iana-iftype-starlan"

    IetfIntfType_iana_iftype_proteon10mbit IetfIntfType = "iana-iftype-proteon10mbit"

    IetfIntfType_iana_iftype_proteon80mbit IetfIntfType = "iana-iftype-proteon80mbit"

    IetfIntfType_iana_iftype_hyperchannel IetfIntfType = "iana-iftype-hyperchannel"

    IetfIntfType_iana_iftype_fddi IetfIntfType = "iana-iftype-fddi"

    IetfIntfType_iana_iftype_lapb IetfIntfType = "iana-iftype-lapb"

    IetfIntfType_iana_iftype_sdlc IetfIntfType = "iana-iftype-sdlc"

    IetfIntfType_iana_iftype_ds1 IetfIntfType = "iana-iftype-ds1"

    IetfIntfType_iana_iftype_e1 IetfIntfType = "iana-iftype-e1"

    IetfIntfType_iana_iftype_basicisdn IetfIntfType = "iana-iftype-basicisdn"

    IetfIntfType_iana_iftype_primaryisdn IetfIntfType = "iana-iftype-primaryisdn"

    IetfIntfType_iana_iftype_prop_p2p_serial IetfIntfType = "iana-iftype-prop-p2p-serial"

    IetfIntfType_iana_iftype_ppp IetfIntfType = "iana-iftype-ppp"

    IetfIntfType_iana_iftype_sw_loopback IetfIntfType = "iana-iftype-sw-loopback"

    IetfIntfType_iana_iftype_eon IetfIntfType = "iana-iftype-eon"

    IetfIntfType_iana_iftype_ethernet3mbit IetfIntfType = "iana-iftype-ethernet3mbit"

    IetfIntfType_iana_iftype_nsip IetfIntfType = "iana-iftype-nsip"

    IetfIntfType_iana_iftype_slip IetfIntfType = "iana-iftype-slip"

    IetfIntfType_iana_iftype_ultra IetfIntfType = "iana-iftype-ultra"

    IetfIntfType_iana_iftype_ds3 IetfIntfType = "iana-iftype-ds3"

    IetfIntfType_iana_iftype_sip IetfIntfType = "iana-iftype-sip"

    IetfIntfType_iana_iftype_framerelay IetfIntfType = "iana-iftype-framerelay"

    IetfIntfType_iana_iftype_rs232 IetfIntfType = "iana-iftype-rs232"

    IetfIntfType_iana_iftype_para IetfIntfType = "iana-iftype-para"

    IetfIntfType_iana_iftype_arcnet IetfIntfType = "iana-iftype-arcnet"

    IetfIntfType_iana_iftype_arcnetplus IetfIntfType = "iana-iftype-arcnetplus"

    IetfIntfType_iana_iftype_atm IetfIntfType = "iana-iftype-atm"

    IetfIntfType_iana_iftype_miox25 IetfIntfType = "iana-iftype-miox25"

    IetfIntfType_iana_iftype_sonet IetfIntfType = "iana-iftype-sonet"

    IetfIntfType_iana_iftype_x25ple IetfIntfType = "iana-iftype-x25ple"

    IetfIntfType_iana_iftype_iso88022_llc IetfIntfType = "iana-iftype-iso88022-llc"

    IetfIntfType_iana_iftype_localtalk IetfIntfType = "iana-iftype-localtalk"

    IetfIntfType_iana_iftype_smdsdxi IetfIntfType = "iana-iftype-smdsdxi"

    IetfIntfType_iana_iftype_framerelay_service IetfIntfType = "iana-iftype-framerelay-service"

    IetfIntfType_iana_iftype_v35 IetfIntfType = "iana-iftype-v35"

    IetfIntfType_iana_iftype_hssi IetfIntfType = "iana-iftype-hssi"

    IetfIntfType_iana_iftype_hippi IetfIntfType = "iana-iftype-hippi"

    IetfIntfType_iana_iftype_modem IetfIntfType = "iana-iftype-modem"

    IetfIntfType_iana_iftype_aal5 IetfIntfType = "iana-iftype-aal5"

    IetfIntfType_iana_iftype_sonetpath IetfIntfType = "iana-iftype-sonetpath"

    IetfIntfType_iana_iftype_sonetvt IetfIntfType = "iana-iftype-sonetvt"

    IetfIntfType_iana_iftype_smdsicip IetfIntfType = "iana-iftype-smdsicip"

    IetfIntfType_iana_iftype_propvirtual IetfIntfType = "iana-iftype-propvirtual"

    IetfIntfType_iana_iftype_propmultiplexor IetfIntfType = "iana-iftype-propmultiplexor"

    IetfIntfType_iana_iftype_ieee80212 IetfIntfType = "iana-iftype-ieee80212"

    IetfIntfType_iana_iftype_fiberchannel IetfIntfType = "iana-iftype-fiberchannel"

    IetfIntfType_iana_iftype_hippi_interface IetfIntfType = "iana-iftype-hippi-interface"

    IetfIntfType_iana_iftype_framerelay_interconnect IetfIntfType = "iana-iftype-framerelay-interconnect"

    IetfIntfType_iana_iftype_aflane8023 IetfIntfType = "iana-iftype-aflane8023"

    IetfIntfType_iana_iftype_aflane8025 IetfIntfType = "iana-iftype-aflane8025"

    IetfIntfType_iana_iftype_cctemul IetfIntfType = "iana-iftype-cctemul"

    IetfIntfType_iana_iftype_fastether IetfIntfType = "iana-iftype-fastether"

    IetfIntfType_iana_iftype_isdn IetfIntfType = "iana-iftype-isdn"

    IetfIntfType_iana_iftype_v11 IetfIntfType = "iana-iftype-v11"

    IetfIntfType_iana_iftype_v36 IetfIntfType = "iana-iftype-v36"

    IetfIntfType_iana_iftype_g703at64k IetfIntfType = "iana-iftype-g703at64k"

    IetfIntfType_iana_iftype_g703at2mb IetfIntfType = "iana-iftype-g703at2mb"

    IetfIntfType_iana_iftype_qllc IetfIntfType = "iana-iftype-qllc"

    IetfIntfType_iana_iftype_fastetherfx IetfIntfType = "iana-iftype-fastetherfx"

    IetfIntfType_iana_iftype_channel IetfIntfType = "iana-iftype-channel"

    IetfIntfType_iana_iftype_ieee80211 IetfIntfType = "iana-iftype-ieee80211"

    IetfIntfType_iana_iftype_ibm370parchan IetfIntfType = "iana-iftype-ibm370parchan"

    IetfIntfType_iana_iftype_escon IetfIntfType = "iana-iftype-escon"

    IetfIntfType_iana_iftype_dlsw IetfIntfType = "iana-iftype-dlsw"

    IetfIntfType_iana_iftype_isdns IetfIntfType = "iana-iftype-isdns"

    IetfIntfType_iana_iftype_isdnu IetfIntfType = "iana-iftype-isdnu"

    IetfIntfType_iana_iftype_lapd IetfIntfType = "iana-iftype-lapd"

    IetfIntfType_iana_iftype_ipswitch IetfIntfType = "iana-iftype-ipswitch"

    IetfIntfType_iana_iftype_rsrb IetfIntfType = "iana-iftype-rsrb"

    IetfIntfType_iana_iftype_atmlogical IetfIntfType = "iana-iftype-atmlogical"

    IetfIntfType_iana_iftype_ds0 IetfIntfType = "iana-iftype-ds0"

    IetfIntfType_iana_iftype_ds0bundle IetfIntfType = "iana-iftype-ds0bundle"

    IetfIntfType_iana_iftype_bsc IetfIntfType = "iana-iftype-bsc"

    IetfIntfType_iana_iftype_async IetfIntfType = "iana-iftype-async"

    IetfIntfType_iana_iftype_cnr IetfIntfType = "iana-iftype-cnr"

    IetfIntfType_iana_iftype_iso88025_dtr IetfIntfType = "iana-iftype-iso88025-dtr"

    IetfIntfType_iana_iftype_eplrs IetfIntfType = "iana-iftype-eplrs"

    IetfIntfType_iana_iftype_arap IetfIntfType = "iana-iftype-arap"

    IetfIntfType_iana_iftype_propcnls IetfIntfType = "iana-iftype-propcnls"

    IetfIntfType_iana_iftype_hostpad IetfIntfType = "iana-iftype-hostpad"

    IetfIntfType_iana_iftype_termpad IetfIntfType = "iana-iftype-termpad"

    IetfIntfType_iana_iftype_framerelay_mpi IetfIntfType = "iana-iftype-framerelay-mpi"

    IetfIntfType_iana_iftype_x213 IetfIntfType = "iana-iftype-x213"

    IetfIntfType_iana_iftype_adsl IetfIntfType = "iana-iftype-adsl"

    IetfIntfType_iana_iftype_radsl IetfIntfType = "iana-iftype-radsl"

    IetfIntfType_iana_iftype_sdsl IetfIntfType = "iana-iftype-sdsl"

    IetfIntfType_iana_iftype_vdsl IetfIntfType = "iana-iftype-vdsl"

    IetfIntfType_iana_iftype_iso88025_crfpint IetfIntfType = "iana-iftype-iso88025-crfpint"

    IetfIntfType_iana_iftype_myrinet IetfIntfType = "iana-iftype-myrinet"

    IetfIntfType_iana_iftype_voiceem IetfIntfType = "iana-iftype-voiceem"

    IetfIntfType_iana_iftype_voicefxo IetfIntfType = "iana-iftype-voicefxo"

    IetfIntfType_iana_iftype_voicefxs IetfIntfType = "iana-iftype-voicefxs"

    IetfIntfType_iana_iftype_voiceencap IetfIntfType = "iana-iftype-voiceencap"

    IetfIntfType_iana_iftype_voip IetfIntfType = "iana-iftype-voip"

    IetfIntfType_iana_iftype_atmdxi IetfIntfType = "iana-iftype-atmdxi"

    IetfIntfType_iana_iftype_atmfuni IetfIntfType = "iana-iftype-atmfuni"

    IetfIntfType_iana_iftype_atmima IetfIntfType = "iana-iftype-atmima"

    IetfIntfType_iana_iftype_ppp_multilinkbundle IetfIntfType = "iana-iftype-ppp-multilinkbundle"

    IetfIntfType_iana_iftype_ipovercdlc IetfIntfType = "iana-iftype-ipovercdlc"

    IetfIntfType_iana_iftype_ipoverclaw IetfIntfType = "iana-iftype-ipoverclaw"

    IetfIntfType_iana_iftype_stack2stack IetfIntfType = "iana-iftype-stack2stack"

    IetfIntfType_iana_iftype_virtualipaddress IetfIntfType = "iana-iftype-virtualipaddress"

    IetfIntfType_iana_iftype_mpc IetfIntfType = "iana-iftype-mpc"

    IetfIntfType_iana_iftype_ipoveratm IetfIntfType = "iana-iftype-ipoveratm"

    IetfIntfType_iana_iftype_iso88025_fiber IetfIntfType = "iana-iftype-iso88025-fiber"

    IetfIntfType_iana_iftype_tdlc IetfIntfType = "iana-iftype-tdlc"

    IetfIntfType_iana_iftype_gige IetfIntfType = "iana-iftype-gige"

    IetfIntfType_iana_iftype_hdlc IetfIntfType = "iana-iftype-hdlc"

    IetfIntfType_iana_iftype_lapf IetfIntfType = "iana-iftype-lapf"

    IetfIntfType_iana_iftype_v37 IetfIntfType = "iana-iftype-v37"

    IetfIntfType_iana_iftype_x25mlp IetfIntfType = "iana-iftype-x25mlp"

    IetfIntfType_iana_iftype_x25huntgroup IetfIntfType = "iana-iftype-x25huntgroup"

    IetfIntfType_iana_iftype_transphdlc IetfIntfType = "iana-iftype-transphdlc"

    IetfIntfType_iana_iftype_interleave IetfIntfType = "iana-iftype-interleave"

    IetfIntfType_iana_iftype_fast IetfIntfType = "iana-iftype-fast"

    IetfIntfType_iana_iftype_ip IetfIntfType = "iana-iftype-ip"

    IetfIntfType_iana_iftype_docs_cable_maclayer IetfIntfType = "iana-iftype-docs-cable-maclayer"

    IetfIntfType_iana_iftype_docs_cable_downstream IetfIntfType = "iana-iftype-docs-cable-downstream"

    IetfIntfType_iana_iftype_docs_cable_upstream IetfIntfType = "iana-iftype-docs-cable-upstream"

    IetfIntfType_iana_iftype_a12mppswitch IetfIntfType = "iana-iftype-a12mppswitch"

    IetfIntfType_iana_iftype_tunnel IetfIntfType = "iana-iftype-tunnel"

    IetfIntfType_iana_iftype_coffee IetfIntfType = "iana-iftype-coffee"

    IetfIntfType_iana_iftype_ces IetfIntfType = "iana-iftype-ces"

    IetfIntfType_iana_iftype_atmsubinterface IetfIntfType = "iana-iftype-atmsubinterface"

    IetfIntfType_iana_iftype_l2vlan IetfIntfType = "iana-iftype-l2vlan"

    IetfIntfType_iana_iftype_l3ipvlan IetfIntfType = "iana-iftype-l3ipvlan"

    IetfIntfType_iana_iftype_l3ipxvlan IetfIntfType = "iana-iftype-l3ipxvlan"

    IetfIntfType_iana_iftype_digital_powerline IetfIntfType = "iana-iftype-digital-powerline"

    IetfIntfType_iana_iftype_media_mailoverip IetfIntfType = "iana-iftype-media-mailoverip"

    IetfIntfType_iana_iftype_dtm IetfIntfType = "iana-iftype-dtm"

    IetfIntfType_iana_iftype_dcn IetfIntfType = "iana-iftype-dcn"

    IetfIntfType_iana_iftype_ipforward IetfIntfType = "iana-iftype-ipforward"

    IetfIntfType_iana_iftype_msdsl IetfIntfType = "iana-iftype-msdsl"

    IetfIntfType_iana_iftype_ieee1394 IetfIntfType = "iana-iftype-ieee1394"

    IetfIntfType_iana_iftype_gsn IetfIntfType = "iana-iftype-gsn"

    IetfIntfType_iana_iftype_dvbrcc_maclayer IetfIntfType = "iana-iftype-dvbrcc-maclayer"

    IetfIntfType_iana_iftype_dvbrcc_downstream IetfIntfType = "iana-iftype-dvbrcc-downstream"

    IetfIntfType_iana_iftype_dvbrcc_upstream IetfIntfType = "iana-iftype-dvbrcc-upstream"

    IetfIntfType_iana_iftype_atmvirtual IetfIntfType = "iana-iftype-atmvirtual"

    IetfIntfType_iana_iftype_mplstunnel IetfIntfType = "iana-iftype-mplstunnel"

    IetfIntfType_iana_iftype_srp IetfIntfType = "iana-iftype-srp"

    IetfIntfType_iana_iftype_voiceoveratm IetfIntfType = "iana-iftype-voiceoveratm"

    IetfIntfType_iana_iftype_voiceoverframerelay IetfIntfType = "iana-iftype-voiceoverframerelay"

    IetfIntfType_iana_iftype_idsl IetfIntfType = "iana-iftype-idsl"

    IetfIntfType_iana_iftype_compositelink IetfIntfType = "iana-iftype-compositelink"

    IetfIntfType_iana_iftype_ss7siglink IetfIntfType = "iana-iftype-ss7siglink"

    IetfIntfType_iana_iftype_propwireless_p2p IetfIntfType = "iana-iftype-propwireless-p2p"

    IetfIntfType_iana_iftype_frforward IetfIntfType = "iana-iftype-frforward"

    IetfIntfType_iana_iftype_rfc1483 IetfIntfType = "iana-iftype-rfc1483"

    IetfIntfType_iana_iftype_usb IetfIntfType = "iana-iftype-usb"

    IetfIntfType_iana_iftype_ieee8023_adlag IetfIntfType = "iana-iftype-ieee8023-adlag"

    IetfIntfType_iana_iftype_bgppolicy_accounting IetfIntfType = "iana-iftype-bgppolicy-accounting"

    IetfIntfType_iana_iftype_frf16mfrbundle IetfIntfType = "iana-iftype-frf16mfrbundle"

    IetfIntfType_iana_iftype_h323gatekeeper IetfIntfType = "iana-iftype-h323gatekeeper"

    IetfIntfType_iana_iftype_h323proxy IetfIntfType = "iana-iftype-h323proxy"

    IetfIntfType_iana_iftype_mpls IetfIntfType = "iana-iftype-mpls"

    IetfIntfType_iana_iftype_mfsiglink IetfIntfType = "iana-iftype-mfsiglink"

    IetfIntfType_iana_iftype_hdsl2 IetfIntfType = "iana-iftype-hdsl2"

    IetfIntfType_iana_iftype_shdsl IetfIntfType = "iana-iftype-shdsl"

    IetfIntfType_iana_iftype_ds1fdl IetfIntfType = "iana-iftype-ds1fdl"

    IetfIntfType_iana_iftype_pos IetfIntfType = "iana-iftype-pos"

    IetfIntfType_iana_iftype_dvbasiin IetfIntfType = "iana-iftype-dvbasiin"

    IetfIntfType_iana_iftype_dvbasiout IetfIntfType = "iana-iftype-dvbasiout"

    IetfIntfType_iana_iftype_plc IetfIntfType = "iana-iftype-plc"

    IetfIntfType_iana_iftype_nfas IetfIntfType = "iana-iftype-nfas"

    IetfIntfType_iana_iftype_tr008 IetfIntfType = "iana-iftype-tr008"

    IetfIntfType_iana_iftype_gr303rdt IetfIntfType = "iana-iftype-gr303rdt"

    IetfIntfType_iana_iftype_gr303idt IetfIntfType = "iana-iftype-gr303idt"

    IetfIntfType_iana_iftype_isup IetfIntfType = "iana-iftype-isup"

    IetfIntfType_iana_iftype_prop_docs_wireless_maclayer IetfIntfType = "iana-iftype-prop-docs-wireless-maclayer"

    IetfIntfType_iana_iftype_prop_docs_wireless_downstream IetfIntfType = "iana-iftype-prop-docs-wireless-downstream"

    IetfIntfType_iana_iftype_prop_docs_wireless_upstream IetfIntfType = "iana-iftype-prop-docs-wireless-upstream"

    IetfIntfType_iana_iftype_hiperlan2 IetfIntfType = "iana-iftype-hiperlan2"

    IetfIntfType_iana_iftype_prop_bwap2mp IetfIntfType = "iana-iftype-prop-bwap2mp"

    IetfIntfType_iana_iftype_sonetoverheadchannel IetfIntfType = "iana-iftype-sonetoverheadchannel"

    IetfIntfType_iana_iftype_digital_wrapperoverheadchannel IetfIntfType = "iana-iftype-digital-wrapperoverheadchannel"

    IetfIntfType_iana_iftype_aal2 IetfIntfType = "iana-iftype-aal2"

    IetfIntfType_iana_iftype_radiomac IetfIntfType = "iana-iftype-radiomac"

    IetfIntfType_iana_iftype_atmradio IetfIntfType = "iana-iftype-atmradio"

    IetfIntfType_iana_iftype_imt IetfIntfType = "iana-iftype-imt"

    IetfIntfType_iana_iftype_mvl IetfIntfType = "iana-iftype-mvl"

    IetfIntfType_iana_iftype_reachhdsl IetfIntfType = "iana-iftype-reachhdsl"

    IetfIntfType_iana_iftype_frdlciendpt IetfIntfType = "iana-iftype-frdlciendpt"

    IetfIntfType_iana_iftype_atmvciendpt IetfIntfType = "iana-iftype-atmvciendpt"

    IetfIntfType_iana_iftype_opticalchannel IetfIntfType = "iana-iftype-opticalchannel"

    IetfIntfType_iana_iftype_opticaltransport IetfIntfType = "iana-iftype-opticaltransport"

    IetfIntfType_iana_iftype_propatm IetfIntfType = "iana-iftype-propatm"

    IetfIntfType_iana_iftype_voiceovercable IetfIntfType = "iana-iftype-voiceovercable"

    IetfIntfType_iana_iftype_infiniband IetfIntfType = "iana-iftype-infiniband"

    IetfIntfType_iana_iftype_telink IetfIntfType = "iana-iftype-telink"

    IetfIntfType_iana_iftype_q2931 IetfIntfType = "iana-iftype-q2931"

    IetfIntfType_iana_iftype_virtualatg IetfIntfType = "iana-iftype-virtualatg"

    IetfIntfType_iana_iftype_siptg IetfIntfType = "iana-iftype-siptg"

    IetfIntfType_iana_iftype_sipsig IetfIntfType = "iana-iftype-sipsig"

    IetfIntfType_iana_iftype_docs_cable_upstreamchannel IetfIntfType = "iana-iftype-docs-cable-upstreamchannel"

    IetfIntfType_iana_iftype_econet IetfIntfType = "iana-iftype-econet"

    IetfIntfType_iana_iftype_pon155 IetfIntfType = "iana-iftype-pon155"

    IetfIntfType_iana_iftype_pon622 IetfIntfType = "iana-iftype-pon622"

    IetfIntfType_iana_iftype_bridge_if IetfIntfType = "iana-iftype-bridge-if"

    IetfIntfType_iana_iftype_linegroup IetfIntfType = "iana-iftype-linegroup"

    IetfIntfType_iana_iftype_voiceemfgd IetfIntfType = "iana-iftype-voiceemfgd"

    IetfIntfType_iana_iftype_voiceefgdeana IetfIntfType = "iana-iftype-voiceefgdeana"

    IetfIntfType_iana_iftype_voicedid IetfIntfType = "iana-iftype-voicedid"

    IetfIntfType_iana_iftype_mpegtransport IetfIntfType = "iana-iftype-mpegtransport"

    IetfIntfType_iana_iftype_sixtofour IetfIntfType = "iana-iftype-sixtofour"

    IetfIntfType_iana_iftype_gtp IetfIntfType = "iana-iftype-gtp"

    IetfIntfType_iana_iftype_pdnetherloop1 IetfIntfType = "iana-iftype-pdnetherloop1"

    IetfIntfType_iana_iftype_pdnetherloop2 IetfIntfType = "iana-iftype-pdnetherloop2"

    IetfIntfType_iana_iftype_opticalchannel_group IetfIntfType = "iana-iftype-opticalchannel-group"

    IetfIntfType_iana_iftype_homepna IetfIntfType = "iana-iftype-homepna"

    IetfIntfType_iana_iftype_gfp IetfIntfType = "iana-iftype-gfp"

    IetfIntfType_iana_iftype_ciscoislvlan IetfIntfType = "iana-iftype-ciscoislvlan"

    IetfIntfType_iana_iftype_actelismetaloop IetfIntfType = "iana-iftype-actelismetaloop"

    IetfIntfType_iana_iftype_fciplink IetfIntfType = "iana-iftype-fciplink"

    IetfIntfType_iana_iftype_rpr IetfIntfType = "iana-iftype-rpr"

    IetfIntfType_iana_iftype_qam IetfIntfType = "iana-iftype-qam"

    IetfIntfType_iana_iftype_lmp IetfIntfType = "iana-iftype-lmp"

    IetfIntfType_iana_iftype_cblvectastar IetfIntfType = "iana-iftype-cblvectastar"

    IetfIntfType_iana_iftype_docs_cable_mcmts_downtream IetfIntfType = "iana-iftype-docs-cable-mcmts-downtream"

    IetfIntfType_iana_iftype_adsl2 IetfIntfType = "iana-iftype-adsl2"

    IetfIntfType_iana_iftype_macseccontrolledif IetfIntfType = "iana-iftype-macseccontrolledif"

    IetfIntfType_iana_iftype_macsecuncontrolledif IetfIntfType = "iana-iftype-macsecuncontrolledif"

    IetfIntfType_iana_iftype_aviciopticalether IetfIntfType = "iana-iftype-aviciopticalether"

    IetfIntfType_iana_iftype_atmbond IetfIntfType = "iana-iftype-atmbond"

    IetfIntfType_iana_iftype_voicefgdos IetfIntfType = "iana-iftype-voicefgdos"

    IetfIntfType_iana_iftype_mocaversion1 IetfIntfType = "iana-iftype-mocaversion1"

    IetfIntfType_iana_iftype_ieee80216_wman IetfIntfType = "iana-iftype-ieee80216-wman"

    IetfIntfType_iana_iftype_adsl2plus IetfIntfType = "iana-iftype-adsl2plus"

    IetfIntfType_iana_iftype_dvbrcsmaclayer IetfIntfType = "iana-iftype-dvbrcsmaclayer"

    IetfIntfType_iana_iftype_dvbtdm IetfIntfType = "iana-iftype-dvbtdm"

    IetfIntfType_iana_iftype_dvbrcstdma IetfIntfType = "iana-iftype-dvbrcstdma"

    IetfIntfType_iana_iftype_x86laps IetfIntfType = "iana-iftype-x86laps"

    IetfIntfType_iana_iftype_wwanpp IetfIntfType = "iana-iftype-wwanpp"

    IetfIntfType_iana_iftype_wwanpp2 IetfIntfType = "iana-iftype-wwanpp2"

    IetfIntfType_iana_iftype_voiceebs IetfIntfType = "iana-iftype-voiceebs"

    IetfIntfType_iana_iftype_ifpwtype IetfIntfType = "iana-iftype-ifpwtype"

    IetfIntfType_iana_iftype_ilan IetfIntfType = "iana-iftype-ilan"

    IetfIntfType_iana_iftype_pip IetfIntfType = "iana-iftype-pip"

    IetfIntfType_iana_iftype_aluelp IetfIntfType = "iana-iftype-aluelp"

    IetfIntfType_iana_iftype_gpon IetfIntfType = "iana-iftype-gpon"

    IetfIntfType_iana_iftype_vdsl2 IetfIntfType = "iana-iftype-vdsl2"

    IetfIntfType_iana_iftype_capwapdot11_profile IetfIntfType = "iana-iftype-capwapdot11-profile"

    IetfIntfType_iana_iftype_capwapdot11_bss IetfIntfType = "iana-iftype-capwapdot11-bss"

    IetfIntfType_iana_iftype_capwapwtp_virtualradio IetfIntfType = "iana-iftype-capwapwtp-virtualradio"

    IetfIntfType_iana_iftype_bits IetfIntfType = "iana-iftype-bits"

    IetfIntfType_iana_iftype_docs_cable_upstreamrfport IetfIntfType = "iana-iftype-docs-cable-upstreamrfport"

    IetfIntfType_iana_iftype_cable_downstreamrfport IetfIntfType = "iana-iftype-cable-downstreamrfport"

    IetfIntfType_iana_iftype_vmware_virtualnic IetfIntfType = "iana-iftype-vmware-virtualnic"

    IetfIntfType_iana_iftype_ieee802154 IetfIntfType = "iana-iftype-ieee802154"

    IetfIntfType_iana_iftype_otnodu IetfIntfType = "iana-iftype-otnodu"

    IetfIntfType_iana_iftype_otnotu IetfIntfType = "iana-iftype-otnotu"

    IetfIntfType_iana_iftype_ifvfitype IetfIntfType = "iana-iftype-ifvfitype"

    IetfIntfType_iana_iftype_g9981 IetfIntfType = "iana-iftype-g9981"

    IetfIntfType_iana_iftype_g9982 IetfIntfType = "iana-iftype-g9982"

    IetfIntfType_iana_iftype_g9983 IetfIntfType = "iana-iftype-g9983"

    IetfIntfType_iana_iftype_aluepon IetfIntfType = "iana-iftype-aluepon"

    IetfIntfType_iana_iftype_aluepon_onu IetfIntfType = "iana-iftype-aluepon-onu"

    IetfIntfType_iana_iftype_aluepon_physicaluni IetfIntfType = "iana-iftype-aluepon-physicaluni"

    IetfIntfType_iana_iftype_aluepon_logicalink IetfIntfType = "iana-iftype-aluepon-logicalink"

    IetfIntfType_iana_iftype_alugpon_onu IetfIntfType = "iana-iftype-alugpon-onu"

    IetfIntfType_iana_iftype_alugpon_physicaluni IetfIntfType = "iana-iftype-alugpon-physicaluni"

    IetfIntfType_iana_iftype_vmwarenicteam IetfIntfType = "iana-iftype-vmwarenicteam"

    IetfIntfType_iana_iftype_docs_ofdm_downstream IetfIntfType = "iana-iftype-docs-ofdm-downstream"

    IetfIntfType_iana_iftype_docs_ofdma_upstream IetfIntfType = "iana-iftype-docs-ofdma-upstream"

    IetfIntfType_iana_iftype_gfast IetfIntfType = "iana-iftype-gfast"

    IetfIntfType_iana_iftype_sdci IetfIntfType = "iana-iftype-sdci"

    IetfIntfType_iana_iftype_xbox_wireless IetfIntfType = "iana-iftype-xbox-wireless"

    IetfIntfType_iana_iftype_fastdsl IetfIntfType = "iana-iftype-fastdsl"
)

// QosDirection represents QoS direction indication
type QosDirection string

const (
    // Direction of traffic coming into the network entry
    QosDirection_qos_inbound QosDirection = "qos-inbound"

    // Direction of traffic going out of the network entry
    QosDirection_qos_outbound QosDirection = "qos-outbound"
)

// EtherSpeed represents The speed setting of the interface
type EtherSpeed string

const (
    EtherSpeed_speed_10mb EtherSpeed = "speed-10mb"

    EtherSpeed_speed_100mb EtherSpeed = "speed-100mb"

    EtherSpeed_speed_1gb EtherSpeed = "speed-1gb"

    EtherSpeed_speed_10gb EtherSpeed = "speed-10gb"

    EtherSpeed_speed_25gb EtherSpeed = "speed-25gb"

    EtherSpeed_speed_40gb EtherSpeed = "speed-40gb"

    EtherSpeed_speed_50gb EtherSpeed = "speed-50gb"

    EtherSpeed_speed_100gb EtherSpeed = "speed-100gb"

    EtherSpeed_speed_unknown EtherSpeed = "speed-unknown"

    EtherSpeed_speed_auto EtherSpeed = "speed-auto"
)

// Interfaces
// Operational state of interfaces
type Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of interfaces. The type is slice of Interfaces_Interface.
    Interface []*Interfaces_Interface
}

func (interfaces *Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xe"
    interfaces.EntityData.ParentYangName = "Cisco-IOS-XE-interfaces-oper"
    interfaces.EntityData.SegmentPath = "Cisco-IOS-XE-interfaces-oper:interfaces"
    interfaces.EntityData.AbsolutePath = interfaces.EntityData.SegmentPath
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    interfaces.EntityData.Children = types.NewOrderedMap()
    interfaces.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range interfaces.Interface {
        interfaces.EntityData.Children.Append(types.GetSegmentPath(interfaces.Interface[i]), types.YChild{"Interface", interfaces.Interface[i]})
    }
    interfaces.EntityData.Leafs = types.NewOrderedMap()

    interfaces.EntityData.YListKeys = []string {}

    return &(interfaces.EntityData)
}

// Interfaces_Interface
// List of interfaces
type Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The name of the interface. A server implementation
    // MAY map this leaf to the ifName MIB object.  Such an implementation needs
    // to use some mechanism to handle the differences in size and characters
    // allowed between this leaf and ifName.  The definition of such a mechanism
    // is outside the scope of this document. The type is string.
    Name interface{}

    // When an interface entry is created, a server MAY initialize the type leaf
    // with a valid value, e.g., if it is possible to derive the type from the
    // name of the interface. If a client tries to set the type of an interface to
    // a value that can never be used by the system, e.g., if the type is not
    // supported or if the type does not match the name of the interface, the
    // server MUST reject the request. A NETCONF server MUST reply with an
    // rpc-error with the error-tag 'invalid-value' in this case. The type is
    // IetfIntfType.
    InterfaceType interface{}

    // The desired state of the interface. This leaf has the same read semantics
    // as ifAdminStatus. The type is IntfState.
    AdminStatus interface{}

    // The current operational state of the interface. This leaf has the same
    // semantics as ifOperStatus. The type is OperState.
    OperStatus interface{}

    // The time the interface entered its current operational state. If the
    // current state was entered prior to the last re-initialization of the local
    // network management subsystem, then this node is not present. The type is
    // string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    LastChange interface{}

    // The ifIndex value for the ifEntry represented by this interface. The type
    // is interface{} with range: -2147483648..2147483647.
    IfIndex interface{}

    // The interface's address at its protocol sub-layer.  For example, for an
    // 802.x interface, this object normally contains a Media Access Control (MAC)
    // address.  The interface's media-specific modules must define the bit and
    // byte ordering and the format of the value of this object.  For interfaces
    // that do not have such an address (e.g., a serial line), this node is not
    // present. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    PhysAddress interface{}

    // A list of references to interfaces layered on top of this interface. The
    // type is slice of string.
    HigherLayerIf []interface{}

    // A list of references to interfaces layered underneath this interface. The
    // type is slice of string.
    LowerLayerIf []interface{}

    // An estimate of the interface's current bandwidth in bits per second.  For
    // interfaces that do not vary in bandwidth or for those where no accurate
    // estimation can be made, this node should contain the nominal bandwidth. For
    // interfaces that have no concept of bandwidth, this node is not present. The
    // type is interface{} with range: 0..18446744073709551615.
    Speed interface{}

    // VRF to which this interface belongs to. If the  interface is not in a VRF
    // then it is 'Global'. The type is string.
    Vrf interface{}

    // IPv4 address configured on interface. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv4 Subnet Mask. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv4SubnetMask interface{}

    // Interface description. The type is string.
    Description interface{}

    // Maximum transmission unit. The type is interface{} with range:
    // 0..4294967295.
    Mtu interface{}

    // Input Security ACL. The type is string.
    InputSecurityAcl interface{}

    // Output Security ACL. The type is string.
    OutputSecurityAcl interface{}

    // The burnt-in mac address that was associated with this interface from
    // manufacturing. This is only relevant for interfaces that have the concept
    // of burnt in ethernet  addresses, otherwise it is zero. The type is string
    // with pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    BiaAddress interface{}

    // A list of the IPv6 addresses associated with the interface. This conatins
    // all the IPv6 addresses, including the link local addresses, assigned to the
    // interface. The type is one of the following types: slice of string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or slice of string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Addrs []interface{}

    // When ip tcp adjust-mss is configured, this vlaue shows the tcp mss, or the
    // value is zero. The type is interface{} with range: 0..65535.
    Ipv4TcpAdjustMss interface{}

    // When ipv6 tcp adjust-mss is configured, this value shows the tcp mss, or
    // the value is zero. The type is interface{} with range: 0..65535.
    Ipv6TcpAdjustMss interface{}

    // No specific interface class information. The type is bool.
    IntfClassUnspecified interface{}

    // A collection of interface-related statistics objects.
    Statistics Interfaces_Interface_Statistics

    // diffserv related details. The type is slice of
    // Interfaces_Interface_DiffservInfo.
    DiffservInfo []*Interfaces_Interface_DiffservInfo

    // IPv4 traffic statistics for this interface.
    V4ProtocolStats Interfaces_Interface_V4ProtocolStats

    // IPv6 traffic statistics for this interface.
    V6ProtocolStats Interfaces_Interface_V6ProtocolStats

    // Operational state variables for logical aggregate / LAG interfaces. The
    // type is slice of Interfaces_Interface_LagAggregateState.
    LagAggregateState []*Interfaces_Interface_LagAggregateState

    // The Ethernet state information.
    EtherState Interfaces_Interface_EtherState

    // The Ethernet statistics.
    EtherStats Interfaces_Interface_EtherStats

    // The T1E1 serial state information.
    SerialState Interfaces_Interface_SerialState

    // The T1E1 statistics.
    SerialStats Interfaces_Interface_SerialStats
}

func (self *Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xe"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.Name, "name")
    self.EntityData.AbsolutePath = "Cisco-IOS-XE-interfaces-oper:interfaces/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("statistics", types.YChild{"Statistics", &self.Statistics})
    self.EntityData.Children.Append("diffserv-info", types.YChild{"DiffservInfo", nil})
    for i := range self.DiffservInfo {
        self.EntityData.Children.Append(types.GetSegmentPath(self.DiffservInfo[i]), types.YChild{"DiffservInfo", self.DiffservInfo[i]})
    }
    self.EntityData.Children.Append("v4-protocol-stats", types.YChild{"V4ProtocolStats", &self.V4ProtocolStats})
    self.EntityData.Children.Append("v6-protocol-stats", types.YChild{"V6ProtocolStats", &self.V6ProtocolStats})
    self.EntityData.Children.Append("lag-aggregate-state", types.YChild{"LagAggregateState", nil})
    for i := range self.LagAggregateState {
        self.EntityData.Children.Append(types.GetSegmentPath(self.LagAggregateState[i]), types.YChild{"LagAggregateState", self.LagAggregateState[i]})
    }
    self.EntityData.Children.Append("ether-state", types.YChild{"EtherState", &self.EtherState})
    self.EntityData.Children.Append("ether-stats", types.YChild{"EtherStats", &self.EtherStats})
    self.EntityData.Children.Append("serial-state", types.YChild{"SerialState", &self.SerialState})
    self.EntityData.Children.Append("serial-stats", types.YChild{"SerialStats", &self.SerialStats})
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("name", types.YLeaf{"Name", self.Name})
    self.EntityData.Leafs.Append("interface-type", types.YLeaf{"InterfaceType", self.InterfaceType})
    self.EntityData.Leafs.Append("admin-status", types.YLeaf{"AdminStatus", self.AdminStatus})
    self.EntityData.Leafs.Append("oper-status", types.YLeaf{"OperStatus", self.OperStatus})
    self.EntityData.Leafs.Append("last-change", types.YLeaf{"LastChange", self.LastChange})
    self.EntityData.Leafs.Append("if-index", types.YLeaf{"IfIndex", self.IfIndex})
    self.EntityData.Leafs.Append("phys-address", types.YLeaf{"PhysAddress", self.PhysAddress})
    self.EntityData.Leafs.Append("higher-layer-if", types.YLeaf{"HigherLayerIf", self.HigherLayerIf})
    self.EntityData.Leafs.Append("lower-layer-if", types.YLeaf{"LowerLayerIf", self.LowerLayerIf})
    self.EntityData.Leafs.Append("speed", types.YLeaf{"Speed", self.Speed})
    self.EntityData.Leafs.Append("vrf", types.YLeaf{"Vrf", self.Vrf})
    self.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", self.Ipv4})
    self.EntityData.Leafs.Append("ipv4-subnet-mask", types.YLeaf{"Ipv4SubnetMask", self.Ipv4SubnetMask})
    self.EntityData.Leafs.Append("description", types.YLeaf{"Description", self.Description})
    self.EntityData.Leafs.Append("mtu", types.YLeaf{"Mtu", self.Mtu})
    self.EntityData.Leafs.Append("input-security-acl", types.YLeaf{"InputSecurityAcl", self.InputSecurityAcl})
    self.EntityData.Leafs.Append("output-security-acl", types.YLeaf{"OutputSecurityAcl", self.OutputSecurityAcl})
    self.EntityData.Leafs.Append("bia-address", types.YLeaf{"BiaAddress", self.BiaAddress})
    self.EntityData.Leafs.Append("ipv6-addrs", types.YLeaf{"Ipv6Addrs", self.Ipv6Addrs})
    self.EntityData.Leafs.Append("ipv4-tcp-adjust-mss", types.YLeaf{"Ipv4TcpAdjustMss", self.Ipv4TcpAdjustMss})
    self.EntityData.Leafs.Append("ipv6-tcp-adjust-mss", types.YLeaf{"Ipv6TcpAdjustMss", self.Ipv6TcpAdjustMss})
    self.EntityData.Leafs.Append("intf-class-unspecified", types.YLeaf{"IntfClassUnspecified", self.IntfClassUnspecified})

    self.EntityData.YListKeys = []string {"Name"}

    return &(self.EntityData)
}

// Interfaces_Interface_Statistics
// A collection of interface-related statistics objects
type Interfaces_Interface_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The time on the most recent occasion at which any one or more of this
    // interface's counters suffered a discontinuity.  If no such discontinuities
    // have occurred since the last re-initialization of the local management
    // subsystem, then this node contains the time the local management subsystem
    // re-initialized itself. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    DiscontinuityTime interface{}

    // The total number of octets received on the interface, including framing
    // characters. Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of discontinuity-time. The type is interface{} with range:
    // 0..18446744073709551615.
    InOctets interface{}

    // The number of packets, delivered by this sub-layer to a higher (sub-)layer,
    // that were not addressed to a multicast or broadcast address at this
    // sub-layer. Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of 'discontinuity-time'. The type is interface{} with range:
    // 0..18446744073709551615.
    InUnicastPkts interface{}

    // The number of packets, delivered by this sub-layer to a higher (sub-)layer,
    // that were addressed to a broadcast address at this sub-layer.
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // 'discontinuity-time'. The type is interface{} with range:
    // 0..18446744073709551615.
    InBroadcastPkts interface{}

    // The number of packets, delivered by this sub-layer to a higher (sub-)layer,
    // that were addressed to a multicast address at this sub-layer.  For a
    // MAC-layer protocol, this includes both Group and Functional addresses.
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // 'discontinuity-time'. The type is interface{} with range:
    // 0..18446744073709551615.
    InMulticastPkts interface{}

    // The number of inbound packets that were chosen to be discarded even though
    // no errors had been detected to prevent their being deliverable to a
    // higher-layer protocol.  One possible reason for discarding such a packet
    // could be to free up buffer space. Discontinuities in the value of this
    // counter can occur at re-initialization of the management system, and at
    // other times as indicated by the value of 'discontinuity-time'. The type is
    // interface{} with range: 0..4294967295.
    InDiscards interface{}

    // For packet-oriented interfaces, the number of inbound packets that
    // contained errors preventing them from being deliverable to a higher-layer
    // protocol.  For character- oriented or fixed-length interfaces, the number
    // of inbound transmission units that contained errors preventing them from
    // being deliverable to a higher-layer protocol. Discontinuities in the value
    // of this counter can occur at re-initialization of the management system,
    // and at other times as indicated by the value of 'discontinuity-time'. The
    // type is interface{} with range: 0..4294967295.
    InErrors interface{}

    // For packet-oriented interfaces, the number of packets received via the
    // interface that were discarded because of an unknown or unsupported
    // protocol.  For character-oriented or fixed-length interfaces that support
    // protocol multiplexing, the number of transmission units received via the
    // interface that were discarded because of an unknown or unsupported
    // protocol. For any interface that does not support protocol multiplexing,
    // this counter is not present. Discontinuities in the value of this counter
    // can occur at re-initialization of the management system, and at other times
    // as indicated by the value of 'discontinuity-time'. The type is interface{}
    // with range: 0..4294967295.
    InUnknownProtos interface{}

    // The total number of octets transmitted out of the interface, including
    // framing characters. Discontinuities in the value of this counter can occur
    // at re-initialization of the management system, and at other times as
    // indicated by the value of 'discontinuity-time'. The type is interface{}
    // with range: 0..4294967295.
    OutOctets interface{}

    // The total number of packets that higher-level protocols requested be
    // transmitted, and that were not addressed to a multicast or broadcast
    // address at this sub-layer, including those that were discarded or not sent.
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // 'discontinuity-time'. The type is interface{} with range:
    // 0..18446744073709551615.
    OutUnicastPkts interface{}

    // The total number of packets that higher-level protocols requested be
    // transmitted, and that were addressed to a broadcast address at this
    // sub-layer, including those that were discarded or not sent. Discontinuities
    // in the value of this counter can occur at re-initialization of the
    // management system, and at other times as indicated by the value of
    // 'discontinuity-time'. The type is interface{} with range:
    // 0..18446744073709551615.
    OutBroadcastPkts interface{}

    // The total number of packets that higher-level protocols requested be
    // transmitted, and that were addressed to a multicast address at this
    // sub-layer, including those that were discarded or not sent.  For a
    // MAC-layer protocol, this includes both Group and Functional addresses.
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // 'discontinuity-time'. The type is interface{} with range:
    // 0..18446744073709551615.
    OutMulticastPkts interface{}

    // The number of outbound packets that were chosen to be discarded even though
    // no errors had been detected to prevent their being transmitted.  One
    // possible reason for discarding such a packet could be to free up buffer
    // space. Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of 'discontinuity-time'. The type is interface{} with range:
    // 0..18446744073709551615.
    OutDiscards interface{}

    // For packet-oriented interfaces, the number of outbound packets that could
    // not be transmitted because of errors. For character-oriented or
    // fixed-length interfaces, the number of outbound transmission units that
    // could not be transmitted because of errors. Discontinuities in the value of
    // this counter can occur at re-initialization of the management system, and
    // at other times as indicated by the value of 'discontinuity-time'. The type
    // is interface{} with range: 0..18446744073709551615.
    OutErrors interface{}

    // The receive packet per second rate on this interface. The type is
    // interface{} with range: 0..18446744073709551615.
    RxPps interface{}

    // The receive kilobits per second rate on this interface. The type is
    // interface{} with range: 0..18446744073709551615.
    RxKbps interface{}

    // The transmit packet per second rate on this interface. The type is
    // interface{} with range: 0..18446744073709551615.
    TxPps interface{}

    // The transmit kilobits per second rate on this interface. The type is
    // interface{} with range: 0..18446744073709551615.
    TxKbps interface{}

    // The number of times the interface state transitioned between up and down.
    // The type is interface{} with range: 0..18446744073709551615.
    NumFlaps interface{}

    // Number of receive error events due to FCS/CRC check failure. The type is
    // interface{} with range: 0..18446744073709551615.
    InCrcErrors interface{}
}

func (statistics *Interfaces_Interface_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xe"
    statistics.EntityData.ParentYangName = "interface"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.AbsolutePath = "Cisco-IOS-XE-interfaces-oper:interfaces/interface/" + statistics.EntityData.SegmentPath
    statistics.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    statistics.EntityData.Children = types.NewOrderedMap()
    statistics.EntityData.Leafs = types.NewOrderedMap()
    statistics.EntityData.Leafs.Append("discontinuity-time", types.YLeaf{"DiscontinuityTime", statistics.DiscontinuityTime})
    statistics.EntityData.Leafs.Append("in-octets", types.YLeaf{"InOctets", statistics.InOctets})
    statistics.EntityData.Leafs.Append("in-unicast-pkts", types.YLeaf{"InUnicastPkts", statistics.InUnicastPkts})
    statistics.EntityData.Leafs.Append("in-broadcast-pkts", types.YLeaf{"InBroadcastPkts", statistics.InBroadcastPkts})
    statistics.EntityData.Leafs.Append("in-multicast-pkts", types.YLeaf{"InMulticastPkts", statistics.InMulticastPkts})
    statistics.EntityData.Leafs.Append("in-discards", types.YLeaf{"InDiscards", statistics.InDiscards})
    statistics.EntityData.Leafs.Append("in-errors", types.YLeaf{"InErrors", statistics.InErrors})
    statistics.EntityData.Leafs.Append("in-unknown-protos", types.YLeaf{"InUnknownProtos", statistics.InUnknownProtos})
    statistics.EntityData.Leafs.Append("out-octets", types.YLeaf{"OutOctets", statistics.OutOctets})
    statistics.EntityData.Leafs.Append("out-unicast-pkts", types.YLeaf{"OutUnicastPkts", statistics.OutUnicastPkts})
    statistics.EntityData.Leafs.Append("out-broadcast-pkts", types.YLeaf{"OutBroadcastPkts", statistics.OutBroadcastPkts})
    statistics.EntityData.Leafs.Append("out-multicast-pkts", types.YLeaf{"OutMulticastPkts", statistics.OutMulticastPkts})
    statistics.EntityData.Leafs.Append("out-discards", types.YLeaf{"OutDiscards", statistics.OutDiscards})
    statistics.EntityData.Leafs.Append("out-errors", types.YLeaf{"OutErrors", statistics.OutErrors})
    statistics.EntityData.Leafs.Append("rx-pps", types.YLeaf{"RxPps", statistics.RxPps})
    statistics.EntityData.Leafs.Append("rx-kbps", types.YLeaf{"RxKbps", statistics.RxKbps})
    statistics.EntityData.Leafs.Append("tx-pps", types.YLeaf{"TxPps", statistics.TxPps})
    statistics.EntityData.Leafs.Append("tx-kbps", types.YLeaf{"TxKbps", statistics.TxKbps})
    statistics.EntityData.Leafs.Append("num-flaps", types.YLeaf{"NumFlaps", statistics.NumFlaps})
    statistics.EntityData.Leafs.Append("in-crc-errors", types.YLeaf{"InCrcErrors", statistics.InCrcErrors})

    statistics.EntityData.YListKeys = []string {}

    return &(statistics.EntityData)
}

// Interfaces_Interface_DiffservInfo
// diffserv related details
type Interfaces_Interface_DiffservInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Direction fo the traffic flow either inbound or
    // outbound. The type is QosDirection.
    Direction interface{}

    // This attribute is a key. Policy entry name. The type is string.
    PolicyName interface{}

    // Statistics for each Classifier Entry in a Policy. The type is slice of
    // Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats.
    DiffservTargetClassifierStats []*Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats

    // Statistics for aggregrate priority per policy instance. The type is slice
    // of Interfaces_Interface_DiffservInfo_PriorityOperList.
    PriorityOperList []*Interfaces_Interface_DiffservInfo_PriorityOperList
}

func (diffservInfo *Interfaces_Interface_DiffservInfo) GetEntityData() *types.CommonEntityData {
    diffservInfo.EntityData.YFilter = diffservInfo.YFilter
    diffservInfo.EntityData.YangName = "diffserv-info"
    diffservInfo.EntityData.BundleName = "cisco_ios_xe"
    diffservInfo.EntityData.ParentYangName = "interface"
    diffservInfo.EntityData.SegmentPath = "diffserv-info" + types.AddKeyToken(diffservInfo.Direction, "direction") + types.AddKeyToken(diffservInfo.PolicyName, "policy-name")
    diffservInfo.EntityData.AbsolutePath = "Cisco-IOS-XE-interfaces-oper:interfaces/interface/" + diffservInfo.EntityData.SegmentPath
    diffservInfo.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffservInfo.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffservInfo.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffservInfo.EntityData.Children = types.NewOrderedMap()
    diffservInfo.EntityData.Children.Append("diffserv-target-classifier-stats", types.YChild{"DiffservTargetClassifierStats", nil})
    for i := range diffservInfo.DiffservTargetClassifierStats {
        diffservInfo.EntityData.Children.Append(types.GetSegmentPath(diffservInfo.DiffservTargetClassifierStats[i]), types.YChild{"DiffservTargetClassifierStats", diffservInfo.DiffservTargetClassifierStats[i]})
    }
    diffservInfo.EntityData.Children.Append("priority-oper-list", types.YChild{"PriorityOperList", nil})
    for i := range diffservInfo.PriorityOperList {
        diffservInfo.EntityData.Children.Append(types.GetSegmentPath(diffservInfo.PriorityOperList[i]), types.YChild{"PriorityOperList", diffservInfo.PriorityOperList[i]})
    }
    diffservInfo.EntityData.Leafs = types.NewOrderedMap()
    diffservInfo.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", diffservInfo.Direction})
    diffservInfo.EntityData.Leafs.Append("policy-name", types.YLeaf{"PolicyName", diffservInfo.PolicyName})

    diffservInfo.EntityData.YListKeys = []string {"Direction", "PolicyName"}

    return &(diffservInfo.EntityData)
}

// Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats
// Statistics for each Classifier Entry in a Policy
type Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Classifier Entry Name. The type is string.
    ClassifierEntryName interface{}

    // This attribute is a key. Path of the Classifier Entry in a hierarchial
    // policy. The type is string.
    ParentPath interface{}

    // Classifier Counters.
    ClassifierEntryStats Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_ClassifierEntryStats

    // Meter statistics. The type is slice of
    // Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MeterStats.
    MeterStats []*Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MeterStats

    // Queuing Counters.
    QueuingStats Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_QueuingStats

    // List of statistics for random-detect based on subclass type and value pair
    // Technically these are a field in the queuing statistics -> wred statisitcs,
    // but GREEN EI does not allow that nesting structure. The type is slice of
    // Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList.
    SubclassList []*Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList

    // Statistics for marking actions.
    MarkingStats Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats
}

func (diffservTargetClassifierStats *Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats) GetEntityData() *types.CommonEntityData {
    diffservTargetClassifierStats.EntityData.YFilter = diffservTargetClassifierStats.YFilter
    diffservTargetClassifierStats.EntityData.YangName = "diffserv-target-classifier-stats"
    diffservTargetClassifierStats.EntityData.BundleName = "cisco_ios_xe"
    diffservTargetClassifierStats.EntityData.ParentYangName = "diffserv-info"
    diffservTargetClassifierStats.EntityData.SegmentPath = "diffserv-target-classifier-stats" + types.AddKeyToken(diffservTargetClassifierStats.ClassifierEntryName, "classifier-entry-name") + types.AddKeyToken(diffservTargetClassifierStats.ParentPath, "parent-path")
    diffservTargetClassifierStats.EntityData.AbsolutePath = "Cisco-IOS-XE-interfaces-oper:interfaces/interface/diffserv-info/" + diffservTargetClassifierStats.EntityData.SegmentPath
    diffservTargetClassifierStats.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffservTargetClassifierStats.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffservTargetClassifierStats.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffservTargetClassifierStats.EntityData.Children = types.NewOrderedMap()
    diffservTargetClassifierStats.EntityData.Children.Append("classifier-entry-stats", types.YChild{"ClassifierEntryStats", &diffservTargetClassifierStats.ClassifierEntryStats})
    diffservTargetClassifierStats.EntityData.Children.Append("meter-stats", types.YChild{"MeterStats", nil})
    for i := range diffservTargetClassifierStats.MeterStats {
        diffservTargetClassifierStats.EntityData.Children.Append(types.GetSegmentPath(diffservTargetClassifierStats.MeterStats[i]), types.YChild{"MeterStats", diffservTargetClassifierStats.MeterStats[i]})
    }
    diffservTargetClassifierStats.EntityData.Children.Append("queuing-stats", types.YChild{"QueuingStats", &diffservTargetClassifierStats.QueuingStats})
    diffservTargetClassifierStats.EntityData.Children.Append("subclass-list", types.YChild{"SubclassList", nil})
    for i := range diffservTargetClassifierStats.SubclassList {
        diffservTargetClassifierStats.EntityData.Children.Append(types.GetSegmentPath(diffservTargetClassifierStats.SubclassList[i]), types.YChild{"SubclassList", diffservTargetClassifierStats.SubclassList[i]})
    }
    diffservTargetClassifierStats.EntityData.Children.Append("marking-stats", types.YChild{"MarkingStats", &diffservTargetClassifierStats.MarkingStats})
    diffservTargetClassifierStats.EntityData.Leafs = types.NewOrderedMap()
    diffservTargetClassifierStats.EntityData.Leafs.Append("classifier-entry-name", types.YLeaf{"ClassifierEntryName", diffservTargetClassifierStats.ClassifierEntryName})
    diffservTargetClassifierStats.EntityData.Leafs.Append("parent-path", types.YLeaf{"ParentPath", diffservTargetClassifierStats.ParentPath})

    diffservTargetClassifierStats.EntityData.YListKeys = []string {"ClassifierEntryName", "ParentPath"}

    return &(diffservTargetClassifierStats.EntityData)
}

// Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_ClassifierEntryStats
// Classifier Counters
type Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_ClassifierEntryStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of total packets which filtered to the classifier-entry. The type is
    // interface{} with range: 0..18446744073709551615.
    ClassifiedPkts interface{}

    // Number of total bytes which filtered to the classifier-entry. The type is
    // interface{} with range: 0..18446744073709551615.
    ClassifiedBytes interface{}

    // Rate of average data flow through the classifier-entry. The type is
    // interface{} with range: 0..18446744073709551615.
    ClassifiedRate interface{}
}

func (classifierEntryStats *Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_ClassifierEntryStats) GetEntityData() *types.CommonEntityData {
    classifierEntryStats.EntityData.YFilter = classifierEntryStats.YFilter
    classifierEntryStats.EntityData.YangName = "classifier-entry-stats"
    classifierEntryStats.EntityData.BundleName = "cisco_ios_xe"
    classifierEntryStats.EntityData.ParentYangName = "diffserv-target-classifier-stats"
    classifierEntryStats.EntityData.SegmentPath = "classifier-entry-stats"
    classifierEntryStats.EntityData.AbsolutePath = "Cisco-IOS-XE-interfaces-oper:interfaces/interface/diffserv-info/diffserv-target-classifier-stats/" + classifierEntryStats.EntityData.SegmentPath
    classifierEntryStats.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    classifierEntryStats.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    classifierEntryStats.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    classifierEntryStats.EntityData.Children = types.NewOrderedMap()
    classifierEntryStats.EntityData.Leafs = types.NewOrderedMap()
    classifierEntryStats.EntityData.Leafs.Append("classified-pkts", types.YLeaf{"ClassifiedPkts", classifierEntryStats.ClassifiedPkts})
    classifierEntryStats.EntityData.Leafs.Append("classified-bytes", types.YLeaf{"ClassifiedBytes", classifierEntryStats.ClassifiedBytes})
    classifierEntryStats.EntityData.Leafs.Append("classified-rate", types.YLeaf{"ClassifiedRate", classifierEntryStats.ClassifiedRate})

    classifierEntryStats.EntityData.YListKeys = []string {}

    return &(classifierEntryStats.EntityData)
}

// Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MeterStats
// Meter statistics
type Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MeterStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Meter Identifier. The type is interface{} with
    // range: 0..65535.
    MeterId interface{}

    // Number of packets which succeed the meter. The type is interface{} with
    // range: 0..18446744073709551615.
    MeterSucceedPkts interface{}

    // Bytes of packets which succeed the meter. The type is interface{} with
    // range: 0..18446744073709551615.
    MeterSucceedBytes interface{}

    // Number of packets which failed the meter. The type is interface{} with
    // range: 0..18446744073709551615.
    MeterFailedPkts interface{}

    // Bytes of packets which failed the meter. The type is interface{} with
    // range: 0..18446744073709551615.
    MeterFailedBytes interface{}
}

func (meterStats *Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MeterStats) GetEntityData() *types.CommonEntityData {
    meterStats.EntityData.YFilter = meterStats.YFilter
    meterStats.EntityData.YangName = "meter-stats"
    meterStats.EntityData.BundleName = "cisco_ios_xe"
    meterStats.EntityData.ParentYangName = "diffserv-target-classifier-stats"
    meterStats.EntityData.SegmentPath = "meter-stats" + types.AddKeyToken(meterStats.MeterId, "meter-id")
    meterStats.EntityData.AbsolutePath = "Cisco-IOS-XE-interfaces-oper:interfaces/interface/diffserv-info/diffserv-target-classifier-stats/" + meterStats.EntityData.SegmentPath
    meterStats.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    meterStats.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    meterStats.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    meterStats.EntityData.Children = types.NewOrderedMap()
    meterStats.EntityData.Leafs = types.NewOrderedMap()
    meterStats.EntityData.Leafs.Append("meter-id", types.YLeaf{"MeterId", meterStats.MeterId})
    meterStats.EntityData.Leafs.Append("meter-succeed-pkts", types.YLeaf{"MeterSucceedPkts", meterStats.MeterSucceedPkts})
    meterStats.EntityData.Leafs.Append("meter-succeed-bytes", types.YLeaf{"MeterSucceedBytes", meterStats.MeterSucceedBytes})
    meterStats.EntityData.Leafs.Append("meter-failed-pkts", types.YLeaf{"MeterFailedPkts", meterStats.MeterFailedPkts})
    meterStats.EntityData.Leafs.Append("meter-failed-bytes", types.YLeaf{"MeterFailedBytes", meterStats.MeterFailedBytes})

    meterStats.EntityData.YListKeys = []string {"MeterId"}

    return &(meterStats.EntityData)
}

// Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_QueuingStats
// Queuing Counters
type Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_QueuingStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of packets transmitted from queue. The type is interface{} with
    // range: 0..18446744073709551615.
    OutputPkts interface{}

    // Number of bytes transmitted from queue. The type is interface{} with range:
    // 0..18446744073709551615.
    OutputBytes interface{}

    // Number of packets currently buffered . The type is interface{} with range:
    // 0..18446744073709551615.
    QueueSizePkts interface{}

    // Number of bytes currently buffered. The type is interface{} with range:
    // 0..18446744073709551615.
    QueueSizeBytes interface{}

    // Total number of packets dropped. The type is interface{} with range:
    // 0..18446744073709551615.
    DropPkts interface{}

    // Total number of bytes dropped. The type is interface{} with range:
    // 0..18446744073709551615.
    DropBytes interface{}

    // WRED Counters.
    WredStats Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_QueuingStats_WredStats

    // CAC statistics.
    CacStats Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_QueuingStats_CacStats
}

func (queuingStats *Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_QueuingStats) GetEntityData() *types.CommonEntityData {
    queuingStats.EntityData.YFilter = queuingStats.YFilter
    queuingStats.EntityData.YangName = "queuing-stats"
    queuingStats.EntityData.BundleName = "cisco_ios_xe"
    queuingStats.EntityData.ParentYangName = "diffserv-target-classifier-stats"
    queuingStats.EntityData.SegmentPath = "queuing-stats"
    queuingStats.EntityData.AbsolutePath = "Cisco-IOS-XE-interfaces-oper:interfaces/interface/diffserv-info/diffserv-target-classifier-stats/" + queuingStats.EntityData.SegmentPath
    queuingStats.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    queuingStats.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    queuingStats.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    queuingStats.EntityData.Children = types.NewOrderedMap()
    queuingStats.EntityData.Children.Append("wred-stats", types.YChild{"WredStats", &queuingStats.WredStats})
    queuingStats.EntityData.Children.Append("cac-stats", types.YChild{"CacStats", &queuingStats.CacStats})
    queuingStats.EntityData.Leafs = types.NewOrderedMap()
    queuingStats.EntityData.Leafs.Append("output-pkts", types.YLeaf{"OutputPkts", queuingStats.OutputPkts})
    queuingStats.EntityData.Leafs.Append("output-bytes", types.YLeaf{"OutputBytes", queuingStats.OutputBytes})
    queuingStats.EntityData.Leafs.Append("queue-size-pkts", types.YLeaf{"QueueSizePkts", queuingStats.QueueSizePkts})
    queuingStats.EntityData.Leafs.Append("queue-size-bytes", types.YLeaf{"QueueSizeBytes", queuingStats.QueueSizeBytes})
    queuingStats.EntityData.Leafs.Append("drop-pkts", types.YLeaf{"DropPkts", queuingStats.DropPkts})
    queuingStats.EntityData.Leafs.Append("drop-bytes", types.YLeaf{"DropBytes", queuingStats.DropBytes})

    queuingStats.EntityData.YListKeys = []string {}

    return &(queuingStats.EntityData)
}

// Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_QueuingStats_WredStats
// WRED Counters
type Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_QueuingStats_WredStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Early drop packets. The type is interface{} with range:
    // 0..18446744073709551615.
    EarlyDropPkts interface{}

    // Early drop bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    EarlyDropBytes interface{}

    // Current mean queue depth. The type is interface{} with range: 0..65535.
    MeanQueueDepth interface{}

    // Transmitted packtes. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedPkts interface{}

    // Transmitted bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedBytes interface{}

    // Total number of packets dropped. The type is interface{} with range:
    // 0..18446744073709551615.
    TailDropPkts interface{}

    // Total number of bytes dropped. The type is interface{} with range:
    // 0..18446744073709551615.
    TailDropBytes interface{}

    // Total number of packets dropped. The type is interface{} with range:
    // 0..18446744073709551615.
    DropPktsFlow interface{}

    // Number of packets dropped due to buffers being unavailable system-wide or
    // at the associated interface. This is a sub-set of drop-pkts. The type is
    // interface{} with range: 0..18446744073709551615.
    DropPktsNoBuffer interface{}

    // Queue max que depth Packets. The type is interface{} with range:
    // 0..18446744073709551615.
    QueuePeakSizePkts interface{}

    // Queue max que depth Bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    QueuePeakSizeBytes interface{}

    // Priority stats. Bandwidth exceed drops. The type is interface{} with range:
    // 0..18446744073709551615.
    BandwidthExceedDrops interface{}
}

func (wredStats *Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_QueuingStats_WredStats) GetEntityData() *types.CommonEntityData {
    wredStats.EntityData.YFilter = wredStats.YFilter
    wredStats.EntityData.YangName = "wred-stats"
    wredStats.EntityData.BundleName = "cisco_ios_xe"
    wredStats.EntityData.ParentYangName = "queuing-stats"
    wredStats.EntityData.SegmentPath = "wred-stats"
    wredStats.EntityData.AbsolutePath = "Cisco-IOS-XE-interfaces-oper:interfaces/interface/diffserv-info/diffserv-target-classifier-stats/queuing-stats/" + wredStats.EntityData.SegmentPath
    wredStats.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    wredStats.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    wredStats.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    wredStats.EntityData.Children = types.NewOrderedMap()
    wredStats.EntityData.Leafs = types.NewOrderedMap()
    wredStats.EntityData.Leafs.Append("early-drop-pkts", types.YLeaf{"EarlyDropPkts", wredStats.EarlyDropPkts})
    wredStats.EntityData.Leafs.Append("early-drop-bytes", types.YLeaf{"EarlyDropBytes", wredStats.EarlyDropBytes})
    wredStats.EntityData.Leafs.Append("mean-queue-depth", types.YLeaf{"MeanQueueDepth", wredStats.MeanQueueDepth})
    wredStats.EntityData.Leafs.Append("transmitted-pkts", types.YLeaf{"TransmittedPkts", wredStats.TransmittedPkts})
    wredStats.EntityData.Leafs.Append("transmitted-bytes", types.YLeaf{"TransmittedBytes", wredStats.TransmittedBytes})
    wredStats.EntityData.Leafs.Append("tail-drop-pkts", types.YLeaf{"TailDropPkts", wredStats.TailDropPkts})
    wredStats.EntityData.Leafs.Append("tail-drop-bytes", types.YLeaf{"TailDropBytes", wredStats.TailDropBytes})
    wredStats.EntityData.Leafs.Append("drop-pkts-flow", types.YLeaf{"DropPktsFlow", wredStats.DropPktsFlow})
    wredStats.EntityData.Leafs.Append("drop-pkts-no-buffer", types.YLeaf{"DropPktsNoBuffer", wredStats.DropPktsNoBuffer})
    wredStats.EntityData.Leafs.Append("queue-peak-size-pkts", types.YLeaf{"QueuePeakSizePkts", wredStats.QueuePeakSizePkts})
    wredStats.EntityData.Leafs.Append("queue-peak-size-bytes", types.YLeaf{"QueuePeakSizeBytes", wredStats.QueuePeakSizeBytes})
    wredStats.EntityData.Leafs.Append("bandwidth-exceed-drops", types.YLeaf{"BandwidthExceedDrops", wredStats.BandwidthExceedDrops})

    wredStats.EntityData.YListKeys = []string {}

    return &(wredStats.EntityData)
}

// Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_QueuingStats_CacStats
// CAC statistics
type Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_QueuingStats_CacStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of admitted flows. The type is interface{} with range:
    // 0..4294967295.
    NumAdmittedFlows interface{}

    // Number of non-admitted flows. The type is interface{} with range:
    // 0..4294967295.
    NumNonAdmittedFlows interface{}
}

func (cacStats *Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_QueuingStats_CacStats) GetEntityData() *types.CommonEntityData {
    cacStats.EntityData.YFilter = cacStats.YFilter
    cacStats.EntityData.YangName = "cac-stats"
    cacStats.EntityData.BundleName = "cisco_ios_xe"
    cacStats.EntityData.ParentYangName = "queuing-stats"
    cacStats.EntityData.SegmentPath = "cac-stats"
    cacStats.EntityData.AbsolutePath = "Cisco-IOS-XE-interfaces-oper:interfaces/interface/diffserv-info/diffserv-target-classifier-stats/queuing-stats/" + cacStats.EntityData.SegmentPath
    cacStats.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cacStats.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cacStats.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cacStats.EntityData.Children = types.NewOrderedMap()
    cacStats.EntityData.Leafs = types.NewOrderedMap()
    cacStats.EntityData.Leafs.Append("num-admitted-flows", types.YLeaf{"NumAdmittedFlows", cacStats.NumAdmittedFlows})
    cacStats.EntityData.Leafs.Append("num-non-admitted-flows", types.YLeaf{"NumNonAdmittedFlows", cacStats.NumNonAdmittedFlows})

    cacStats.EntityData.YListKeys = []string {}

    return &(cacStats.EntityData)
}

// Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList
// List of statistics for random-detect
// based on subclass type and value pair
// Technically these are a field in the queuing
// statistics -> wred statisitcs, but GREEN EI
// does not allow that nesting structure
type Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Subclass match type. The type is QosMatchType.
    MatchType interface{}

    // Counters for sub-class matching a range of Class-of-Service (COS) value
    // (and, optionally, additional COS range. The type is slice of
    // Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList_CosCounters.
    CosCounters []*Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList_CosCounters

    // statistics for cos default.
    CosDefault Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList_CosDefault

    // List for statistics based on dscp value range. The type is slice of
    // Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList_DscpCounters.
    DscpCounters []*Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList_DscpCounters

    // Statistics for dscp default.
    DscpDefault Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList_DscpDefault

    // Composed multiple discard class ranges. The type is slice of
    // Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList_DiscardClassCounters.
    DiscardClassCounters []*Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList_DiscardClassCounters

    // Statistics for discard class default.
    DiscClassDefault Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList_DiscClassDefault

    // List for statistics based on precedence value range. The type is slice of
    // Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList_PrecedenceCounters.
    PrecedenceCounters []*Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList_PrecedenceCounters

    // Precedence default.
    PrecDefault Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList_PrecDefault

    // List for statistics based on mpls exp value range. The type is slice of
    // Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList_MplsExpCounters.
    MplsExpCounters []*Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList_MplsExpCounters

    // Statistics for mpls-exp default.
    MplsExpDefault Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList_MplsExpDefault

    // Composed by multiple dei ranges. The type is slice of
    // Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList_DeiCounters.
    DeiCounters []*Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList_DeiCounters

    // Statistics for dei default.
    DeiCountsDefault Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList_DeiCountsDefault

    // Statistics for each value range for a specifc subclass type. The type is
    // slice of
    // Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList_ClpCounters.
    ClpCounters []*Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList_ClpCounters

    // Statistic for atm clp default.
    ClpDefault Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList_ClpDefault
}

func (subclassList *Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList) GetEntityData() *types.CommonEntityData {
    subclassList.EntityData.YFilter = subclassList.YFilter
    subclassList.EntityData.YangName = "subclass-list"
    subclassList.EntityData.BundleName = "cisco_ios_xe"
    subclassList.EntityData.ParentYangName = "diffserv-target-classifier-stats"
    subclassList.EntityData.SegmentPath = "subclass-list" + types.AddKeyToken(subclassList.MatchType, "match-type")
    subclassList.EntityData.AbsolutePath = "Cisco-IOS-XE-interfaces-oper:interfaces/interface/diffserv-info/diffserv-target-classifier-stats/" + subclassList.EntityData.SegmentPath
    subclassList.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    subclassList.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    subclassList.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    subclassList.EntityData.Children = types.NewOrderedMap()
    subclassList.EntityData.Children.Append("cos-counters", types.YChild{"CosCounters", nil})
    for i := range subclassList.CosCounters {
        subclassList.EntityData.Children.Append(types.GetSegmentPath(subclassList.CosCounters[i]), types.YChild{"CosCounters", subclassList.CosCounters[i]})
    }
    subclassList.EntityData.Children.Append("cos-default", types.YChild{"CosDefault", &subclassList.CosDefault})
    subclassList.EntityData.Children.Append("dscp-counters", types.YChild{"DscpCounters", nil})
    for i := range subclassList.DscpCounters {
        subclassList.EntityData.Children.Append(types.GetSegmentPath(subclassList.DscpCounters[i]), types.YChild{"DscpCounters", subclassList.DscpCounters[i]})
    }
    subclassList.EntityData.Children.Append("dscp-default", types.YChild{"DscpDefault", &subclassList.DscpDefault})
    subclassList.EntityData.Children.Append("discard-class-counters", types.YChild{"DiscardClassCounters", nil})
    for i := range subclassList.DiscardClassCounters {
        subclassList.EntityData.Children.Append(types.GetSegmentPath(subclassList.DiscardClassCounters[i]), types.YChild{"DiscardClassCounters", subclassList.DiscardClassCounters[i]})
    }
    subclassList.EntityData.Children.Append("disc-class-default", types.YChild{"DiscClassDefault", &subclassList.DiscClassDefault})
    subclassList.EntityData.Children.Append("precedence-counters", types.YChild{"PrecedenceCounters", nil})
    for i := range subclassList.PrecedenceCounters {
        subclassList.EntityData.Children.Append(types.GetSegmentPath(subclassList.PrecedenceCounters[i]), types.YChild{"PrecedenceCounters", subclassList.PrecedenceCounters[i]})
    }
    subclassList.EntityData.Children.Append("prec-default", types.YChild{"PrecDefault", &subclassList.PrecDefault})
    subclassList.EntityData.Children.Append("mpls-exp-counters", types.YChild{"MplsExpCounters", nil})
    for i := range subclassList.MplsExpCounters {
        subclassList.EntityData.Children.Append(types.GetSegmentPath(subclassList.MplsExpCounters[i]), types.YChild{"MplsExpCounters", subclassList.MplsExpCounters[i]})
    }
    subclassList.EntityData.Children.Append("mpls-exp-default", types.YChild{"MplsExpDefault", &subclassList.MplsExpDefault})
    subclassList.EntityData.Children.Append("dei-counters", types.YChild{"DeiCounters", nil})
    for i := range subclassList.DeiCounters {
        subclassList.EntityData.Children.Append(types.GetSegmentPath(subclassList.DeiCounters[i]), types.YChild{"DeiCounters", subclassList.DeiCounters[i]})
    }
    subclassList.EntityData.Children.Append("dei-counts-default", types.YChild{"DeiCountsDefault", &subclassList.DeiCountsDefault})
    subclassList.EntityData.Children.Append("clp-counters", types.YChild{"ClpCounters", nil})
    for i := range subclassList.ClpCounters {
        subclassList.EntityData.Children.Append(types.GetSegmentPath(subclassList.ClpCounters[i]), types.YChild{"ClpCounters", subclassList.ClpCounters[i]})
    }
    subclassList.EntityData.Children.Append("clp-default", types.YChild{"ClpDefault", &subclassList.ClpDefault})
    subclassList.EntityData.Leafs = types.NewOrderedMap()
    subclassList.EntityData.Leafs.Append("match-type", types.YLeaf{"MatchType", subclassList.MatchType})

    subclassList.EntityData.YListKeys = []string {"MatchType"}

    return &(subclassList.EntityData)
}

// Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList_CosCounters
// Counters for sub-class matching a range of
// Class-of-Service (COS) value (and, optionally, additional COS range
type Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList_CosCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Min COS value. The type is interface{} with range:
    // 0..4294967295.
    CosMin interface{}

    // This attribute is a key. Max COS value. The type is interface{} with range:
    // 0..4294967295.
    CosMax interface{}

    // Transmitted packtes. The type is interface{} with range:
    // 0..18446744073709551615.
    WredTxPkts interface{}

    // Transmitted bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    WredTxBytes interface{}

    // Tail drop packets. The type is interface{} with range:
    // 0..18446744073709551615.
    WredTailDropPkts interface{}

    // Tail drop bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    WredTailDropBytes interface{}

    // Early drop packets. The type is interface{} with range:
    // 0..18446744073709551615.
    WredEarlyDropPkts interface{}

    // Early drop bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    WredEarlyDropBytes interface{}
}

func (cosCounters *Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList_CosCounters) GetEntityData() *types.CommonEntityData {
    cosCounters.EntityData.YFilter = cosCounters.YFilter
    cosCounters.EntityData.YangName = "cos-counters"
    cosCounters.EntityData.BundleName = "cisco_ios_xe"
    cosCounters.EntityData.ParentYangName = "subclass-list"
    cosCounters.EntityData.SegmentPath = "cos-counters" + types.AddKeyToken(cosCounters.CosMin, "cos-min") + types.AddKeyToken(cosCounters.CosMax, "cos-max")
    cosCounters.EntityData.AbsolutePath = "Cisco-IOS-XE-interfaces-oper:interfaces/interface/diffserv-info/diffserv-target-classifier-stats/subclass-list/" + cosCounters.EntityData.SegmentPath
    cosCounters.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cosCounters.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cosCounters.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cosCounters.EntityData.Children = types.NewOrderedMap()
    cosCounters.EntityData.Leafs = types.NewOrderedMap()
    cosCounters.EntityData.Leafs.Append("cos-min", types.YLeaf{"CosMin", cosCounters.CosMin})
    cosCounters.EntityData.Leafs.Append("cos-max", types.YLeaf{"CosMax", cosCounters.CosMax})
    cosCounters.EntityData.Leafs.Append("wred-tx-pkts", types.YLeaf{"WredTxPkts", cosCounters.WredTxPkts})
    cosCounters.EntityData.Leafs.Append("wred-tx-bytes", types.YLeaf{"WredTxBytes", cosCounters.WredTxBytes})
    cosCounters.EntityData.Leafs.Append("wred-tail-drop-pkts", types.YLeaf{"WredTailDropPkts", cosCounters.WredTailDropPkts})
    cosCounters.EntityData.Leafs.Append("wred-tail-drop-bytes", types.YLeaf{"WredTailDropBytes", cosCounters.WredTailDropBytes})
    cosCounters.EntityData.Leafs.Append("wred-early-drop-pkts", types.YLeaf{"WredEarlyDropPkts", cosCounters.WredEarlyDropPkts})
    cosCounters.EntityData.Leafs.Append("wred-early-drop-bytes", types.YLeaf{"WredEarlyDropBytes", cosCounters.WredEarlyDropBytes})

    cosCounters.EntityData.YListKeys = []string {"CosMin", "CosMax"}

    return &(cosCounters.EntityData)
}

// Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList_CosDefault
// statistics for cos default
type Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList_CosDefault struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Transmitted packtes. The type is interface{} with range:
    // 0..18446744073709551615.
    WredTxPkts interface{}

    // Transmitted bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    WredTxBytes interface{}

    // Tail drop packets. The type is interface{} with range:
    // 0..18446744073709551615.
    WredTailDropPkts interface{}

    // Tail drop bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    WredTailDropBytes interface{}

    // Early drop packets. The type is interface{} with range:
    // 0..18446744073709551615.
    WredEarlyDropPkts interface{}

    // Early drop bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    WredEarlyDropBytes interface{}
}

func (cosDefault *Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList_CosDefault) GetEntityData() *types.CommonEntityData {
    cosDefault.EntityData.YFilter = cosDefault.YFilter
    cosDefault.EntityData.YangName = "cos-default"
    cosDefault.EntityData.BundleName = "cisco_ios_xe"
    cosDefault.EntityData.ParentYangName = "subclass-list"
    cosDefault.EntityData.SegmentPath = "cos-default"
    cosDefault.EntityData.AbsolutePath = "Cisco-IOS-XE-interfaces-oper:interfaces/interface/diffserv-info/diffserv-target-classifier-stats/subclass-list/" + cosDefault.EntityData.SegmentPath
    cosDefault.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cosDefault.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cosDefault.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cosDefault.EntityData.Children = types.NewOrderedMap()
    cosDefault.EntityData.Leafs = types.NewOrderedMap()
    cosDefault.EntityData.Leafs.Append("wred-tx-pkts", types.YLeaf{"WredTxPkts", cosDefault.WredTxPkts})
    cosDefault.EntityData.Leafs.Append("wred-tx-bytes", types.YLeaf{"WredTxBytes", cosDefault.WredTxBytes})
    cosDefault.EntityData.Leafs.Append("wred-tail-drop-pkts", types.YLeaf{"WredTailDropPkts", cosDefault.WredTailDropPkts})
    cosDefault.EntityData.Leafs.Append("wred-tail-drop-bytes", types.YLeaf{"WredTailDropBytes", cosDefault.WredTailDropBytes})
    cosDefault.EntityData.Leafs.Append("wred-early-drop-pkts", types.YLeaf{"WredEarlyDropPkts", cosDefault.WredEarlyDropPkts})
    cosDefault.EntityData.Leafs.Append("wred-early-drop-bytes", types.YLeaf{"WredEarlyDropBytes", cosDefault.WredEarlyDropBytes})

    cosDefault.EntityData.YListKeys = []string {}

    return &(cosDefault.EntityData)
}

// Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList_DscpCounters
// List for statistics based on dscp value range
type Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList_DscpCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Minimum of dscp range. The type is interface{}
    // with range: 0..4294967295.
    DscpMin interface{}

    // This attribute is a key. Maximum of dscp range. The type is interface{}
    // with range: 0..4294967295.
    DscpMax interface{}

    // Transmitted packtes. The type is interface{} with range:
    // 0..18446744073709551615.
    WredTxPkts interface{}

    // Transmitted bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    WredTxBytes interface{}

    // Tail drop packets. The type is interface{} with range:
    // 0..18446744073709551615.
    WredTailDropPkts interface{}

    // Tail drop bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    WredTailDropBytes interface{}

    // Early drop packets. The type is interface{} with range:
    // 0..18446744073709551615.
    WredEarlyDropPkts interface{}

    // Early drop bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    WredEarlyDropBytes interface{}
}

func (dscpCounters *Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList_DscpCounters) GetEntityData() *types.CommonEntityData {
    dscpCounters.EntityData.YFilter = dscpCounters.YFilter
    dscpCounters.EntityData.YangName = "dscp-counters"
    dscpCounters.EntityData.BundleName = "cisco_ios_xe"
    dscpCounters.EntityData.ParentYangName = "subclass-list"
    dscpCounters.EntityData.SegmentPath = "dscp-counters" + types.AddKeyToken(dscpCounters.DscpMin, "dscp-min") + types.AddKeyToken(dscpCounters.DscpMax, "dscp-max")
    dscpCounters.EntityData.AbsolutePath = "Cisco-IOS-XE-interfaces-oper:interfaces/interface/diffserv-info/diffserv-target-classifier-stats/subclass-list/" + dscpCounters.EntityData.SegmentPath
    dscpCounters.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dscpCounters.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dscpCounters.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dscpCounters.EntityData.Children = types.NewOrderedMap()
    dscpCounters.EntityData.Leafs = types.NewOrderedMap()
    dscpCounters.EntityData.Leafs.Append("dscp-min", types.YLeaf{"DscpMin", dscpCounters.DscpMin})
    dscpCounters.EntityData.Leafs.Append("dscp-max", types.YLeaf{"DscpMax", dscpCounters.DscpMax})
    dscpCounters.EntityData.Leafs.Append("wred-tx-pkts", types.YLeaf{"WredTxPkts", dscpCounters.WredTxPkts})
    dscpCounters.EntityData.Leafs.Append("wred-tx-bytes", types.YLeaf{"WredTxBytes", dscpCounters.WredTxBytes})
    dscpCounters.EntityData.Leafs.Append("wred-tail-drop-pkts", types.YLeaf{"WredTailDropPkts", dscpCounters.WredTailDropPkts})
    dscpCounters.EntityData.Leafs.Append("wred-tail-drop-bytes", types.YLeaf{"WredTailDropBytes", dscpCounters.WredTailDropBytes})
    dscpCounters.EntityData.Leafs.Append("wred-early-drop-pkts", types.YLeaf{"WredEarlyDropPkts", dscpCounters.WredEarlyDropPkts})
    dscpCounters.EntityData.Leafs.Append("wred-early-drop-bytes", types.YLeaf{"WredEarlyDropBytes", dscpCounters.WredEarlyDropBytes})

    dscpCounters.EntityData.YListKeys = []string {"DscpMin", "DscpMax"}

    return &(dscpCounters.EntityData)
}

// Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList_DscpDefault
// Statistics for dscp default
type Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList_DscpDefault struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Transmitted packtes. The type is interface{} with range:
    // 0..18446744073709551615.
    WredTxPkts interface{}

    // Transmitted bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    WredTxBytes interface{}

    // Tail drop packets. The type is interface{} with range:
    // 0..18446744073709551615.
    WredTailDropPkts interface{}

    // Tail drop bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    WredTailDropBytes interface{}

    // Early drop packets. The type is interface{} with range:
    // 0..18446744073709551615.
    WredEarlyDropPkts interface{}

    // Early drop bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    WredEarlyDropBytes interface{}
}

func (dscpDefault *Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList_DscpDefault) GetEntityData() *types.CommonEntityData {
    dscpDefault.EntityData.YFilter = dscpDefault.YFilter
    dscpDefault.EntityData.YangName = "dscp-default"
    dscpDefault.EntityData.BundleName = "cisco_ios_xe"
    dscpDefault.EntityData.ParentYangName = "subclass-list"
    dscpDefault.EntityData.SegmentPath = "dscp-default"
    dscpDefault.EntityData.AbsolutePath = "Cisco-IOS-XE-interfaces-oper:interfaces/interface/diffserv-info/diffserv-target-classifier-stats/subclass-list/" + dscpDefault.EntityData.SegmentPath
    dscpDefault.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dscpDefault.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dscpDefault.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dscpDefault.EntityData.Children = types.NewOrderedMap()
    dscpDefault.EntityData.Leafs = types.NewOrderedMap()
    dscpDefault.EntityData.Leafs.Append("wred-tx-pkts", types.YLeaf{"WredTxPkts", dscpDefault.WredTxPkts})
    dscpDefault.EntityData.Leafs.Append("wred-tx-bytes", types.YLeaf{"WredTxBytes", dscpDefault.WredTxBytes})
    dscpDefault.EntityData.Leafs.Append("wred-tail-drop-pkts", types.YLeaf{"WredTailDropPkts", dscpDefault.WredTailDropPkts})
    dscpDefault.EntityData.Leafs.Append("wred-tail-drop-bytes", types.YLeaf{"WredTailDropBytes", dscpDefault.WredTailDropBytes})
    dscpDefault.EntityData.Leafs.Append("wred-early-drop-pkts", types.YLeaf{"WredEarlyDropPkts", dscpDefault.WredEarlyDropPkts})
    dscpDefault.EntityData.Leafs.Append("wred-early-drop-bytes", types.YLeaf{"WredEarlyDropBytes", dscpDefault.WredEarlyDropBytes})

    dscpDefault.EntityData.YListKeys = []string {}

    return &(dscpDefault.EntityData)
}

// Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList_DiscardClassCounters
// Composed multiple discard class ranges
type Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList_DiscardClassCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Minimum value for discard class in the range. The
    // type is interface{} with range: 0..4294967295.
    DiscClassMin interface{}

    // This attribute is a key. Maximum value for discard class in the range. The
    // type is interface{} with range: 0..4294967295.
    DiscClassMax interface{}

    // Transmitted packtes. The type is interface{} with range:
    // 0..18446744073709551615.
    WredTxPkts interface{}

    // Transmitted bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    WredTxBytes interface{}

    // Tail drop packets. The type is interface{} with range:
    // 0..18446744073709551615.
    WredTailDropPkts interface{}

    // Tail drop bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    WredTailDropBytes interface{}

    // Early drop packets. The type is interface{} with range:
    // 0..18446744073709551615.
    WredEarlyDropPkts interface{}

    // Early drop bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    WredEarlyDropBytes interface{}
}

func (discardClassCounters *Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList_DiscardClassCounters) GetEntityData() *types.CommonEntityData {
    discardClassCounters.EntityData.YFilter = discardClassCounters.YFilter
    discardClassCounters.EntityData.YangName = "discard-class-counters"
    discardClassCounters.EntityData.BundleName = "cisco_ios_xe"
    discardClassCounters.EntityData.ParentYangName = "subclass-list"
    discardClassCounters.EntityData.SegmentPath = "discard-class-counters" + types.AddKeyToken(discardClassCounters.DiscClassMin, "disc-class-min") + types.AddKeyToken(discardClassCounters.DiscClassMax, "disc-class-max")
    discardClassCounters.EntityData.AbsolutePath = "Cisco-IOS-XE-interfaces-oper:interfaces/interface/diffserv-info/diffserv-target-classifier-stats/subclass-list/" + discardClassCounters.EntityData.SegmentPath
    discardClassCounters.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    discardClassCounters.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    discardClassCounters.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    discardClassCounters.EntityData.Children = types.NewOrderedMap()
    discardClassCounters.EntityData.Leafs = types.NewOrderedMap()
    discardClassCounters.EntityData.Leafs.Append("disc-class-min", types.YLeaf{"DiscClassMin", discardClassCounters.DiscClassMin})
    discardClassCounters.EntityData.Leafs.Append("disc-class-max", types.YLeaf{"DiscClassMax", discardClassCounters.DiscClassMax})
    discardClassCounters.EntityData.Leafs.Append("wred-tx-pkts", types.YLeaf{"WredTxPkts", discardClassCounters.WredTxPkts})
    discardClassCounters.EntityData.Leafs.Append("wred-tx-bytes", types.YLeaf{"WredTxBytes", discardClassCounters.WredTxBytes})
    discardClassCounters.EntityData.Leafs.Append("wred-tail-drop-pkts", types.YLeaf{"WredTailDropPkts", discardClassCounters.WredTailDropPkts})
    discardClassCounters.EntityData.Leafs.Append("wred-tail-drop-bytes", types.YLeaf{"WredTailDropBytes", discardClassCounters.WredTailDropBytes})
    discardClassCounters.EntityData.Leafs.Append("wred-early-drop-pkts", types.YLeaf{"WredEarlyDropPkts", discardClassCounters.WredEarlyDropPkts})
    discardClassCounters.EntityData.Leafs.Append("wred-early-drop-bytes", types.YLeaf{"WredEarlyDropBytes", discardClassCounters.WredEarlyDropBytes})

    discardClassCounters.EntityData.YListKeys = []string {"DiscClassMin", "DiscClassMax"}

    return &(discardClassCounters.EntityData)
}

// Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList_DiscClassDefault
// Statistics for discard class default
type Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList_DiscClassDefault struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Transmitted packtes. The type is interface{} with range:
    // 0..18446744073709551615.
    WredTxPkts interface{}

    // Transmitted bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    WredTxBytes interface{}

    // Tail drop packets. The type is interface{} with range:
    // 0..18446744073709551615.
    WredTailDropPkts interface{}

    // Tail drop bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    WredTailDropBytes interface{}

    // Early drop packets. The type is interface{} with range:
    // 0..18446744073709551615.
    WredEarlyDropPkts interface{}

    // Early drop bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    WredEarlyDropBytes interface{}
}

func (discClassDefault *Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList_DiscClassDefault) GetEntityData() *types.CommonEntityData {
    discClassDefault.EntityData.YFilter = discClassDefault.YFilter
    discClassDefault.EntityData.YangName = "disc-class-default"
    discClassDefault.EntityData.BundleName = "cisco_ios_xe"
    discClassDefault.EntityData.ParentYangName = "subclass-list"
    discClassDefault.EntityData.SegmentPath = "disc-class-default"
    discClassDefault.EntityData.AbsolutePath = "Cisco-IOS-XE-interfaces-oper:interfaces/interface/diffserv-info/diffserv-target-classifier-stats/subclass-list/" + discClassDefault.EntityData.SegmentPath
    discClassDefault.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    discClassDefault.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    discClassDefault.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    discClassDefault.EntityData.Children = types.NewOrderedMap()
    discClassDefault.EntityData.Leafs = types.NewOrderedMap()
    discClassDefault.EntityData.Leafs.Append("wred-tx-pkts", types.YLeaf{"WredTxPkts", discClassDefault.WredTxPkts})
    discClassDefault.EntityData.Leafs.Append("wred-tx-bytes", types.YLeaf{"WredTxBytes", discClassDefault.WredTxBytes})
    discClassDefault.EntityData.Leafs.Append("wred-tail-drop-pkts", types.YLeaf{"WredTailDropPkts", discClassDefault.WredTailDropPkts})
    discClassDefault.EntityData.Leafs.Append("wred-tail-drop-bytes", types.YLeaf{"WredTailDropBytes", discClassDefault.WredTailDropBytes})
    discClassDefault.EntityData.Leafs.Append("wred-early-drop-pkts", types.YLeaf{"WredEarlyDropPkts", discClassDefault.WredEarlyDropPkts})
    discClassDefault.EntityData.Leafs.Append("wred-early-drop-bytes", types.YLeaf{"WredEarlyDropBytes", discClassDefault.WredEarlyDropBytes})

    discClassDefault.EntityData.YListKeys = []string {}

    return &(discClassDefault.EntityData)
}

// Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList_PrecedenceCounters
// List for statistics based on precedence value range
type Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList_PrecedenceCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Precedence min. The type is interface{} with
    // range: 0..4294967295.
    PrecMin interface{}

    // This attribute is a key. Precedence max. The type is interface{} with
    // range: 0..4294967295.
    PrecMax interface{}

    // Transmitted packtes. The type is interface{} with range:
    // 0..18446744073709551615.
    WredTxPkts interface{}

    // Transmitted bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    WredTxBytes interface{}

    // Tail drop packets. The type is interface{} with range:
    // 0..18446744073709551615.
    WredTailDropPkts interface{}

    // Tail drop bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    WredTailDropBytes interface{}

    // Early drop packets. The type is interface{} with range:
    // 0..18446744073709551615.
    WredEarlyDropPkts interface{}

    // Early drop bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    WredEarlyDropBytes interface{}
}

func (precedenceCounters *Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList_PrecedenceCounters) GetEntityData() *types.CommonEntityData {
    precedenceCounters.EntityData.YFilter = precedenceCounters.YFilter
    precedenceCounters.EntityData.YangName = "precedence-counters"
    precedenceCounters.EntityData.BundleName = "cisco_ios_xe"
    precedenceCounters.EntityData.ParentYangName = "subclass-list"
    precedenceCounters.EntityData.SegmentPath = "precedence-counters" + types.AddKeyToken(precedenceCounters.PrecMin, "prec-min") + types.AddKeyToken(precedenceCounters.PrecMax, "prec-max")
    precedenceCounters.EntityData.AbsolutePath = "Cisco-IOS-XE-interfaces-oper:interfaces/interface/diffserv-info/diffserv-target-classifier-stats/subclass-list/" + precedenceCounters.EntityData.SegmentPath
    precedenceCounters.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    precedenceCounters.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    precedenceCounters.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    precedenceCounters.EntityData.Children = types.NewOrderedMap()
    precedenceCounters.EntityData.Leafs = types.NewOrderedMap()
    precedenceCounters.EntityData.Leafs.Append("prec-min", types.YLeaf{"PrecMin", precedenceCounters.PrecMin})
    precedenceCounters.EntityData.Leafs.Append("prec-max", types.YLeaf{"PrecMax", precedenceCounters.PrecMax})
    precedenceCounters.EntityData.Leafs.Append("wred-tx-pkts", types.YLeaf{"WredTxPkts", precedenceCounters.WredTxPkts})
    precedenceCounters.EntityData.Leafs.Append("wred-tx-bytes", types.YLeaf{"WredTxBytes", precedenceCounters.WredTxBytes})
    precedenceCounters.EntityData.Leafs.Append("wred-tail-drop-pkts", types.YLeaf{"WredTailDropPkts", precedenceCounters.WredTailDropPkts})
    precedenceCounters.EntityData.Leafs.Append("wred-tail-drop-bytes", types.YLeaf{"WredTailDropBytes", precedenceCounters.WredTailDropBytes})
    precedenceCounters.EntityData.Leafs.Append("wred-early-drop-pkts", types.YLeaf{"WredEarlyDropPkts", precedenceCounters.WredEarlyDropPkts})
    precedenceCounters.EntityData.Leafs.Append("wred-early-drop-bytes", types.YLeaf{"WredEarlyDropBytes", precedenceCounters.WredEarlyDropBytes})

    precedenceCounters.EntityData.YListKeys = []string {"PrecMin", "PrecMax"}

    return &(precedenceCounters.EntityData)
}

// Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList_PrecDefault
// Precedence default
type Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList_PrecDefault struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Transmitted packtes. The type is interface{} with range:
    // 0..18446744073709551615.
    WredTxPkts interface{}

    // Transmitted bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    WredTxBytes interface{}

    // Tail drop packets. The type is interface{} with range:
    // 0..18446744073709551615.
    WredTailDropPkts interface{}

    // Tail drop bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    WredTailDropBytes interface{}

    // Early drop packets. The type is interface{} with range:
    // 0..18446744073709551615.
    WredEarlyDropPkts interface{}

    // Early drop bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    WredEarlyDropBytes interface{}
}

func (precDefault *Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList_PrecDefault) GetEntityData() *types.CommonEntityData {
    precDefault.EntityData.YFilter = precDefault.YFilter
    precDefault.EntityData.YangName = "prec-default"
    precDefault.EntityData.BundleName = "cisco_ios_xe"
    precDefault.EntityData.ParentYangName = "subclass-list"
    precDefault.EntityData.SegmentPath = "prec-default"
    precDefault.EntityData.AbsolutePath = "Cisco-IOS-XE-interfaces-oper:interfaces/interface/diffserv-info/diffserv-target-classifier-stats/subclass-list/" + precDefault.EntityData.SegmentPath
    precDefault.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    precDefault.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    precDefault.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    precDefault.EntityData.Children = types.NewOrderedMap()
    precDefault.EntityData.Leafs = types.NewOrderedMap()
    precDefault.EntityData.Leafs.Append("wred-tx-pkts", types.YLeaf{"WredTxPkts", precDefault.WredTxPkts})
    precDefault.EntityData.Leafs.Append("wred-tx-bytes", types.YLeaf{"WredTxBytes", precDefault.WredTxBytes})
    precDefault.EntityData.Leafs.Append("wred-tail-drop-pkts", types.YLeaf{"WredTailDropPkts", precDefault.WredTailDropPkts})
    precDefault.EntityData.Leafs.Append("wred-tail-drop-bytes", types.YLeaf{"WredTailDropBytes", precDefault.WredTailDropBytes})
    precDefault.EntityData.Leafs.Append("wred-early-drop-pkts", types.YLeaf{"WredEarlyDropPkts", precDefault.WredEarlyDropPkts})
    precDefault.EntityData.Leafs.Append("wred-early-drop-bytes", types.YLeaf{"WredEarlyDropBytes", precDefault.WredEarlyDropBytes})

    precDefault.EntityData.YListKeys = []string {}

    return &(precDefault.EntityData)
}

// Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList_MplsExpCounters
// List for statistics based on mpls exp value range
type Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList_MplsExpCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The minimum EXP field value to be used as match
    // criteria. Any number from 0 to 7. The type is interface{} with range:
    // 0..4294967295.
    ExpMin interface{}

    // This attribute is a key. The maximum EXP field value to be used as match
    // criteria. Any number from 0 to 7. The type is interface{} with range:
    // 0..4294967295.
    ExpMax interface{}

    // Transmitted packtes. The type is interface{} with range:
    // 0..18446744073709551615.
    WredTxPkts interface{}

    // Transmitted bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    WredTxBytes interface{}

    // Tail drop packets. The type is interface{} with range:
    // 0..18446744073709551615.
    WredTailDropPkts interface{}

    // Tail drop bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    WredTailDropBytes interface{}

    // Early drop packets. The type is interface{} with range:
    // 0..18446744073709551615.
    WredEarlyDropPkts interface{}

    // Early drop bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    WredEarlyDropBytes interface{}
}

func (mplsExpCounters *Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList_MplsExpCounters) GetEntityData() *types.CommonEntityData {
    mplsExpCounters.EntityData.YFilter = mplsExpCounters.YFilter
    mplsExpCounters.EntityData.YangName = "mpls-exp-counters"
    mplsExpCounters.EntityData.BundleName = "cisco_ios_xe"
    mplsExpCounters.EntityData.ParentYangName = "subclass-list"
    mplsExpCounters.EntityData.SegmentPath = "mpls-exp-counters" + types.AddKeyToken(mplsExpCounters.ExpMin, "exp-min") + types.AddKeyToken(mplsExpCounters.ExpMax, "exp-max")
    mplsExpCounters.EntityData.AbsolutePath = "Cisco-IOS-XE-interfaces-oper:interfaces/interface/diffserv-info/diffserv-target-classifier-stats/subclass-list/" + mplsExpCounters.EntityData.SegmentPath
    mplsExpCounters.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsExpCounters.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsExpCounters.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsExpCounters.EntityData.Children = types.NewOrderedMap()
    mplsExpCounters.EntityData.Leafs = types.NewOrderedMap()
    mplsExpCounters.EntityData.Leafs.Append("exp-min", types.YLeaf{"ExpMin", mplsExpCounters.ExpMin})
    mplsExpCounters.EntityData.Leafs.Append("exp-max", types.YLeaf{"ExpMax", mplsExpCounters.ExpMax})
    mplsExpCounters.EntityData.Leafs.Append("wred-tx-pkts", types.YLeaf{"WredTxPkts", mplsExpCounters.WredTxPkts})
    mplsExpCounters.EntityData.Leafs.Append("wred-tx-bytes", types.YLeaf{"WredTxBytes", mplsExpCounters.WredTxBytes})
    mplsExpCounters.EntityData.Leafs.Append("wred-tail-drop-pkts", types.YLeaf{"WredTailDropPkts", mplsExpCounters.WredTailDropPkts})
    mplsExpCounters.EntityData.Leafs.Append("wred-tail-drop-bytes", types.YLeaf{"WredTailDropBytes", mplsExpCounters.WredTailDropBytes})
    mplsExpCounters.EntityData.Leafs.Append("wred-early-drop-pkts", types.YLeaf{"WredEarlyDropPkts", mplsExpCounters.WredEarlyDropPkts})
    mplsExpCounters.EntityData.Leafs.Append("wred-early-drop-bytes", types.YLeaf{"WredEarlyDropBytes", mplsExpCounters.WredEarlyDropBytes})

    mplsExpCounters.EntityData.YListKeys = []string {"ExpMin", "ExpMax"}

    return &(mplsExpCounters.EntityData)
}

// Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList_MplsExpDefault
// Statistics for mpls-exp default
type Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList_MplsExpDefault struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Transmitted packtes. The type is interface{} with range:
    // 0..18446744073709551615.
    WredTxPkts interface{}

    // Transmitted bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    WredTxBytes interface{}

    // Tail drop packets. The type is interface{} with range:
    // 0..18446744073709551615.
    WredTailDropPkts interface{}

    // Tail drop bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    WredTailDropBytes interface{}

    // Early drop packets. The type is interface{} with range:
    // 0..18446744073709551615.
    WredEarlyDropPkts interface{}

    // Early drop bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    WredEarlyDropBytes interface{}
}

func (mplsExpDefault *Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList_MplsExpDefault) GetEntityData() *types.CommonEntityData {
    mplsExpDefault.EntityData.YFilter = mplsExpDefault.YFilter
    mplsExpDefault.EntityData.YangName = "mpls-exp-default"
    mplsExpDefault.EntityData.BundleName = "cisco_ios_xe"
    mplsExpDefault.EntityData.ParentYangName = "subclass-list"
    mplsExpDefault.EntityData.SegmentPath = "mpls-exp-default"
    mplsExpDefault.EntityData.AbsolutePath = "Cisco-IOS-XE-interfaces-oper:interfaces/interface/diffserv-info/diffserv-target-classifier-stats/subclass-list/" + mplsExpDefault.EntityData.SegmentPath
    mplsExpDefault.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsExpDefault.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsExpDefault.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsExpDefault.EntityData.Children = types.NewOrderedMap()
    mplsExpDefault.EntityData.Leafs = types.NewOrderedMap()
    mplsExpDefault.EntityData.Leafs.Append("wred-tx-pkts", types.YLeaf{"WredTxPkts", mplsExpDefault.WredTxPkts})
    mplsExpDefault.EntityData.Leafs.Append("wred-tx-bytes", types.YLeaf{"WredTxBytes", mplsExpDefault.WredTxBytes})
    mplsExpDefault.EntityData.Leafs.Append("wred-tail-drop-pkts", types.YLeaf{"WredTailDropPkts", mplsExpDefault.WredTailDropPkts})
    mplsExpDefault.EntityData.Leafs.Append("wred-tail-drop-bytes", types.YLeaf{"WredTailDropBytes", mplsExpDefault.WredTailDropBytes})
    mplsExpDefault.EntityData.Leafs.Append("wred-early-drop-pkts", types.YLeaf{"WredEarlyDropPkts", mplsExpDefault.WredEarlyDropPkts})
    mplsExpDefault.EntityData.Leafs.Append("wred-early-drop-bytes", types.YLeaf{"WredEarlyDropBytes", mplsExpDefault.WredEarlyDropBytes})

    mplsExpDefault.EntityData.YListKeys = []string {}

    return &(mplsExpDefault.EntityData)
}

// Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList_DeiCounters
// Composed by multiple dei ranges
type Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList_DeiCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Dei min. The type is interface{} with range:
    // 0..4294967295.
    DeiMin interface{}

    // This attribute is a key. Dei max. The type is interface{} with range:
    // 0..4294967295.
    DeiMax interface{}

    // Transmitted packtes. The type is interface{} with range:
    // 0..18446744073709551615.
    WredTxPkts interface{}

    // Transmitted bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    WredTxBytes interface{}

    // Tail drop packets. The type is interface{} with range:
    // 0..18446744073709551615.
    WredTailDropPkts interface{}

    // Tail drop bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    WredTailDropBytes interface{}

    // Early drop packets. The type is interface{} with range:
    // 0..18446744073709551615.
    WredEarlyDropPkts interface{}

    // Early drop bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    WredEarlyDropBytes interface{}
}

func (deiCounters *Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList_DeiCounters) GetEntityData() *types.CommonEntityData {
    deiCounters.EntityData.YFilter = deiCounters.YFilter
    deiCounters.EntityData.YangName = "dei-counters"
    deiCounters.EntityData.BundleName = "cisco_ios_xe"
    deiCounters.EntityData.ParentYangName = "subclass-list"
    deiCounters.EntityData.SegmentPath = "dei-counters" + types.AddKeyToken(deiCounters.DeiMin, "dei-min") + types.AddKeyToken(deiCounters.DeiMax, "dei-max")
    deiCounters.EntityData.AbsolutePath = "Cisco-IOS-XE-interfaces-oper:interfaces/interface/diffserv-info/diffserv-target-classifier-stats/subclass-list/" + deiCounters.EntityData.SegmentPath
    deiCounters.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    deiCounters.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    deiCounters.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    deiCounters.EntityData.Children = types.NewOrderedMap()
    deiCounters.EntityData.Leafs = types.NewOrderedMap()
    deiCounters.EntityData.Leafs.Append("dei-min", types.YLeaf{"DeiMin", deiCounters.DeiMin})
    deiCounters.EntityData.Leafs.Append("dei-max", types.YLeaf{"DeiMax", deiCounters.DeiMax})
    deiCounters.EntityData.Leafs.Append("wred-tx-pkts", types.YLeaf{"WredTxPkts", deiCounters.WredTxPkts})
    deiCounters.EntityData.Leafs.Append("wred-tx-bytes", types.YLeaf{"WredTxBytes", deiCounters.WredTxBytes})
    deiCounters.EntityData.Leafs.Append("wred-tail-drop-pkts", types.YLeaf{"WredTailDropPkts", deiCounters.WredTailDropPkts})
    deiCounters.EntityData.Leafs.Append("wred-tail-drop-bytes", types.YLeaf{"WredTailDropBytes", deiCounters.WredTailDropBytes})
    deiCounters.EntityData.Leafs.Append("wred-early-drop-pkts", types.YLeaf{"WredEarlyDropPkts", deiCounters.WredEarlyDropPkts})
    deiCounters.EntityData.Leafs.Append("wred-early-drop-bytes", types.YLeaf{"WredEarlyDropBytes", deiCounters.WredEarlyDropBytes})

    deiCounters.EntityData.YListKeys = []string {"DeiMin", "DeiMax"}

    return &(deiCounters.EntityData)
}

// Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList_DeiCountsDefault
// Statistics for dei default
type Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList_DeiCountsDefault struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Transmitted packtes. The type is interface{} with range:
    // 0..18446744073709551615.
    WredTxPkts interface{}

    // Transmitted bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    WredTxBytes interface{}

    // Tail drop packets. The type is interface{} with range:
    // 0..18446744073709551615.
    WredTailDropPkts interface{}

    // Tail drop bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    WredTailDropBytes interface{}

    // Early drop packets. The type is interface{} with range:
    // 0..18446744073709551615.
    WredEarlyDropPkts interface{}

    // Early drop bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    WredEarlyDropBytes interface{}
}

func (deiCountsDefault *Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList_DeiCountsDefault) GetEntityData() *types.CommonEntityData {
    deiCountsDefault.EntityData.YFilter = deiCountsDefault.YFilter
    deiCountsDefault.EntityData.YangName = "dei-counts-default"
    deiCountsDefault.EntityData.BundleName = "cisco_ios_xe"
    deiCountsDefault.EntityData.ParentYangName = "subclass-list"
    deiCountsDefault.EntityData.SegmentPath = "dei-counts-default"
    deiCountsDefault.EntityData.AbsolutePath = "Cisco-IOS-XE-interfaces-oper:interfaces/interface/diffserv-info/diffserv-target-classifier-stats/subclass-list/" + deiCountsDefault.EntityData.SegmentPath
    deiCountsDefault.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    deiCountsDefault.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    deiCountsDefault.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    deiCountsDefault.EntityData.Children = types.NewOrderedMap()
    deiCountsDefault.EntityData.Leafs = types.NewOrderedMap()
    deiCountsDefault.EntityData.Leafs.Append("wred-tx-pkts", types.YLeaf{"WredTxPkts", deiCountsDefault.WredTxPkts})
    deiCountsDefault.EntityData.Leafs.Append("wred-tx-bytes", types.YLeaf{"WredTxBytes", deiCountsDefault.WredTxBytes})
    deiCountsDefault.EntityData.Leafs.Append("wred-tail-drop-pkts", types.YLeaf{"WredTailDropPkts", deiCountsDefault.WredTailDropPkts})
    deiCountsDefault.EntityData.Leafs.Append("wred-tail-drop-bytes", types.YLeaf{"WredTailDropBytes", deiCountsDefault.WredTailDropBytes})
    deiCountsDefault.EntityData.Leafs.Append("wred-early-drop-pkts", types.YLeaf{"WredEarlyDropPkts", deiCountsDefault.WredEarlyDropPkts})
    deiCountsDefault.EntityData.Leafs.Append("wred-early-drop-bytes", types.YLeaf{"WredEarlyDropBytes", deiCountsDefault.WredEarlyDropBytes})

    deiCountsDefault.EntityData.YListKeys = []string {}

    return &(deiCountsDefault.EntityData)
}

// Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList_ClpCounters
// Statistics for each value range for a specifc subclass type
type Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList_ClpCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Composed by multiple atm clp values. The type is
    // interface{} with range: 0..4294967295.
    ClpVal interface{}

    // Transmitted packtes. The type is interface{} with range:
    // 0..18446744073709551615.
    WredTxPkts interface{}

    // Transmitted bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    WredTxBytes interface{}

    // Tail drop packets. The type is interface{} with range:
    // 0..18446744073709551615.
    WredTailDropPkts interface{}

    // Tail drop bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    WredTailDropBytes interface{}

    // Early drop packets. The type is interface{} with range:
    // 0..18446744073709551615.
    WredEarlyDropPkts interface{}

    // Early drop bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    WredEarlyDropBytes interface{}
}

func (clpCounters *Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList_ClpCounters) GetEntityData() *types.CommonEntityData {
    clpCounters.EntityData.YFilter = clpCounters.YFilter
    clpCounters.EntityData.YangName = "clp-counters"
    clpCounters.EntityData.BundleName = "cisco_ios_xe"
    clpCounters.EntityData.ParentYangName = "subclass-list"
    clpCounters.EntityData.SegmentPath = "clp-counters" + types.AddKeyToken(clpCounters.ClpVal, "clp-val")
    clpCounters.EntityData.AbsolutePath = "Cisco-IOS-XE-interfaces-oper:interfaces/interface/diffserv-info/diffserv-target-classifier-stats/subclass-list/" + clpCounters.EntityData.SegmentPath
    clpCounters.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    clpCounters.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    clpCounters.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    clpCounters.EntityData.Children = types.NewOrderedMap()
    clpCounters.EntityData.Leafs = types.NewOrderedMap()
    clpCounters.EntityData.Leafs.Append("clp-val", types.YLeaf{"ClpVal", clpCounters.ClpVal})
    clpCounters.EntityData.Leafs.Append("wred-tx-pkts", types.YLeaf{"WredTxPkts", clpCounters.WredTxPkts})
    clpCounters.EntityData.Leafs.Append("wred-tx-bytes", types.YLeaf{"WredTxBytes", clpCounters.WredTxBytes})
    clpCounters.EntityData.Leafs.Append("wred-tail-drop-pkts", types.YLeaf{"WredTailDropPkts", clpCounters.WredTailDropPkts})
    clpCounters.EntityData.Leafs.Append("wred-tail-drop-bytes", types.YLeaf{"WredTailDropBytes", clpCounters.WredTailDropBytes})
    clpCounters.EntityData.Leafs.Append("wred-early-drop-pkts", types.YLeaf{"WredEarlyDropPkts", clpCounters.WredEarlyDropPkts})
    clpCounters.EntityData.Leafs.Append("wred-early-drop-bytes", types.YLeaf{"WredEarlyDropBytes", clpCounters.WredEarlyDropBytes})

    clpCounters.EntityData.YListKeys = []string {"ClpVal"}

    return &(clpCounters.EntityData)
}

// Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList_ClpDefault
// Statistic for atm clp default
type Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList_ClpDefault struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Transmitted packtes. The type is interface{} with range:
    // 0..18446744073709551615.
    WredTxPkts interface{}

    // Transmitted bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    WredTxBytes interface{}

    // Tail drop packets. The type is interface{} with range:
    // 0..18446744073709551615.
    WredTailDropPkts interface{}

    // Tail drop bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    WredTailDropBytes interface{}

    // Early drop packets. The type is interface{} with range:
    // 0..18446744073709551615.
    WredEarlyDropPkts interface{}

    // Early drop bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    WredEarlyDropBytes interface{}
}

func (clpDefault *Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_SubclassList_ClpDefault) GetEntityData() *types.CommonEntityData {
    clpDefault.EntityData.YFilter = clpDefault.YFilter
    clpDefault.EntityData.YangName = "clp-default"
    clpDefault.EntityData.BundleName = "cisco_ios_xe"
    clpDefault.EntityData.ParentYangName = "subclass-list"
    clpDefault.EntityData.SegmentPath = "clp-default"
    clpDefault.EntityData.AbsolutePath = "Cisco-IOS-XE-interfaces-oper:interfaces/interface/diffserv-info/diffserv-target-classifier-stats/subclass-list/" + clpDefault.EntityData.SegmentPath
    clpDefault.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    clpDefault.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    clpDefault.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    clpDefault.EntityData.Children = types.NewOrderedMap()
    clpDefault.EntityData.Leafs = types.NewOrderedMap()
    clpDefault.EntityData.Leafs.Append("wred-tx-pkts", types.YLeaf{"WredTxPkts", clpDefault.WredTxPkts})
    clpDefault.EntityData.Leafs.Append("wred-tx-bytes", types.YLeaf{"WredTxBytes", clpDefault.WredTxBytes})
    clpDefault.EntityData.Leafs.Append("wred-tail-drop-pkts", types.YLeaf{"WredTailDropPkts", clpDefault.WredTailDropPkts})
    clpDefault.EntityData.Leafs.Append("wred-tail-drop-bytes", types.YLeaf{"WredTailDropBytes", clpDefault.WredTailDropBytes})
    clpDefault.EntityData.Leafs.Append("wred-early-drop-pkts", types.YLeaf{"WredEarlyDropPkts", clpDefault.WredEarlyDropPkts})
    clpDefault.EntityData.Leafs.Append("wred-early-drop-bytes", types.YLeaf{"WredEarlyDropBytes", clpDefault.WredEarlyDropBytes})

    clpDefault.EntityData.YListKeys = []string {}

    return &(clpDefault.EntityData)
}

// Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats
// Statistics for marking actions
type Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Statistics for set dscp.
    MarkingDscpStatsVal Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingDscpStatsVal

    // Statistics for set dscp tunnel.
    MarkingDscpTunnelStatsVal Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingDscpTunnelStatsVal

    // Statistics for set cos.
    MarkingCosStatsVal Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingCosStatsVal

    // Statistics for set cos-inner.
    MarkingCosInnerStatsVal Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingCosInnerStatsVal

    // Statistics for set discard-class.
    MarkingDiscardClassStatsVal Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingDiscardClassStatsVal

    // Statistics for set qos-group.
    MarkingQosGrpStatsVal Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingQosGrpStatsVal

    // Statistics for set precedence.
    MarkingPrecStatsVal Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingPrecStatsVal

    // Statistics for set precedence tunnel.
    MarkingPrecTunnelStatsVal Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingPrecTunnelStatsVal

    // Statistics for set mpls exp imposition.
    MarkingMplsExpImpStatsVal Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingMplsExpImpStatsVal

    // Statistics for set mpls exp topmost.
    MarkingMplsExpTopStatsVal Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingMplsExpTopStatsVal

    // Statistics for set fr-de.
    MarkingFrDeStatsVal Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingFrDeStatsVal

    // Statistics for set fr-fecn-becn.
    MarkingFrFecnBecnStatsVal Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingFrFecnBecnStatsVal

    // Statistics for set atm-clp.
    MarkingAtmClpStatsVal Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingAtmClpStatsVal

    // Statistics for set vlan-inner.
    MarkingVlanInnerStatsVal Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingVlanInnerStatsVal

    // Statistics for set dei.
    MarkingDeiStatsVal Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingDeiStatsVal

    // Statistics for set dei-imposition.
    MarkingDeiImpStatsVal Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingDeiImpStatsVal

    // Statistics for set srp-priority.
    MarkingSrpPriorityStatsVal Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingSrpPriorityStatsVal

    // Statistics for set wlan-user-priority.
    MarkingWlanUserPriorityStatsVal Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingWlanUserPriorityStatsVal
}

func (markingStats *Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats) GetEntityData() *types.CommonEntityData {
    markingStats.EntityData.YFilter = markingStats.YFilter
    markingStats.EntityData.YangName = "marking-stats"
    markingStats.EntityData.BundleName = "cisco_ios_xe"
    markingStats.EntityData.ParentYangName = "diffserv-target-classifier-stats"
    markingStats.EntityData.SegmentPath = "marking-stats"
    markingStats.EntityData.AbsolutePath = "Cisco-IOS-XE-interfaces-oper:interfaces/interface/diffserv-info/diffserv-target-classifier-stats/" + markingStats.EntityData.SegmentPath
    markingStats.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    markingStats.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    markingStats.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    markingStats.EntityData.Children = types.NewOrderedMap()
    markingStats.EntityData.Children.Append("marking-dscp-stats-val", types.YChild{"MarkingDscpStatsVal", &markingStats.MarkingDscpStatsVal})
    markingStats.EntityData.Children.Append("marking-dscp-tunnel-stats-val", types.YChild{"MarkingDscpTunnelStatsVal", &markingStats.MarkingDscpTunnelStatsVal})
    markingStats.EntityData.Children.Append("marking-cos-stats-val", types.YChild{"MarkingCosStatsVal", &markingStats.MarkingCosStatsVal})
    markingStats.EntityData.Children.Append("marking-cos-inner-stats-val", types.YChild{"MarkingCosInnerStatsVal", &markingStats.MarkingCosInnerStatsVal})
    markingStats.EntityData.Children.Append("marking-discard-class-stats-val", types.YChild{"MarkingDiscardClassStatsVal", &markingStats.MarkingDiscardClassStatsVal})
    markingStats.EntityData.Children.Append("marking-qos-grp-stats-val", types.YChild{"MarkingQosGrpStatsVal", &markingStats.MarkingQosGrpStatsVal})
    markingStats.EntityData.Children.Append("marking-prec-stats-val", types.YChild{"MarkingPrecStatsVal", &markingStats.MarkingPrecStatsVal})
    markingStats.EntityData.Children.Append("marking-prec-tunnel-stats-val", types.YChild{"MarkingPrecTunnelStatsVal", &markingStats.MarkingPrecTunnelStatsVal})
    markingStats.EntityData.Children.Append("marking-mpls-exp-imp-stats-val", types.YChild{"MarkingMplsExpImpStatsVal", &markingStats.MarkingMplsExpImpStatsVal})
    markingStats.EntityData.Children.Append("marking-mpls-exp-top-stats-val", types.YChild{"MarkingMplsExpTopStatsVal", &markingStats.MarkingMplsExpTopStatsVal})
    markingStats.EntityData.Children.Append("marking-fr-de-stats-val", types.YChild{"MarkingFrDeStatsVal", &markingStats.MarkingFrDeStatsVal})
    markingStats.EntityData.Children.Append("marking-fr-fecn-becn-stats-val", types.YChild{"MarkingFrFecnBecnStatsVal", &markingStats.MarkingFrFecnBecnStatsVal})
    markingStats.EntityData.Children.Append("marking-atm-clp-stats-val", types.YChild{"MarkingAtmClpStatsVal", &markingStats.MarkingAtmClpStatsVal})
    markingStats.EntityData.Children.Append("marking-vlan-inner-stats-val", types.YChild{"MarkingVlanInnerStatsVal", &markingStats.MarkingVlanInnerStatsVal})
    markingStats.EntityData.Children.Append("marking-dei-stats-val", types.YChild{"MarkingDeiStatsVal", &markingStats.MarkingDeiStatsVal})
    markingStats.EntityData.Children.Append("marking-dei-imp-stats-val", types.YChild{"MarkingDeiImpStatsVal", &markingStats.MarkingDeiImpStatsVal})
    markingStats.EntityData.Children.Append("marking-srp-priority-stats-val", types.YChild{"MarkingSrpPriorityStatsVal", &markingStats.MarkingSrpPriorityStatsVal})
    markingStats.EntityData.Children.Append("marking-wlan-user-priority-stats-val", types.YChild{"MarkingWlanUserPriorityStatsVal", &markingStats.MarkingWlanUserPriorityStatsVal})
    markingStats.EntityData.Leafs = types.NewOrderedMap()

    markingStats.EntityData.YListKeys = []string {}

    return &(markingStats.EntityData)
}

// Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingDscpStatsVal
// Statistics for set dscp
type Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingDscpStatsVal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // dscp marking. The type is interface{} with range: 0..4294967295.
    Dscp interface{}

    // Number of packets been marked. The type is interface{} with range:
    // 0..18446744073709551615.
    MarkedPkts interface{}
}

func (markingDscpStatsVal *Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingDscpStatsVal) GetEntityData() *types.CommonEntityData {
    markingDscpStatsVal.EntityData.YFilter = markingDscpStatsVal.YFilter
    markingDscpStatsVal.EntityData.YangName = "marking-dscp-stats-val"
    markingDscpStatsVal.EntityData.BundleName = "cisco_ios_xe"
    markingDscpStatsVal.EntityData.ParentYangName = "marking-stats"
    markingDscpStatsVal.EntityData.SegmentPath = "marking-dscp-stats-val"
    markingDscpStatsVal.EntityData.AbsolutePath = "Cisco-IOS-XE-interfaces-oper:interfaces/interface/diffserv-info/diffserv-target-classifier-stats/marking-stats/" + markingDscpStatsVal.EntityData.SegmentPath
    markingDscpStatsVal.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    markingDscpStatsVal.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    markingDscpStatsVal.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    markingDscpStatsVal.EntityData.Children = types.NewOrderedMap()
    markingDscpStatsVal.EntityData.Leafs = types.NewOrderedMap()
    markingDscpStatsVal.EntityData.Leafs.Append("dscp", types.YLeaf{"Dscp", markingDscpStatsVal.Dscp})
    markingDscpStatsVal.EntityData.Leafs.Append("marked-pkts", types.YLeaf{"MarkedPkts", markingDscpStatsVal.MarkedPkts})

    markingDscpStatsVal.EntityData.YListKeys = []string {}

    return &(markingDscpStatsVal.EntityData)
}

// Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingDscpTunnelStatsVal
// Statistics for set dscp tunnel
type Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingDscpTunnelStatsVal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // dscp value. The type is interface{} with range: 0..4294967295.
    DscpVal interface{}

    // Number of packets been marked. The type is interface{} with range:
    // 0..18446744073709551615.
    MarkedPkts interface{}
}

func (markingDscpTunnelStatsVal *Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingDscpTunnelStatsVal) GetEntityData() *types.CommonEntityData {
    markingDscpTunnelStatsVal.EntityData.YFilter = markingDscpTunnelStatsVal.YFilter
    markingDscpTunnelStatsVal.EntityData.YangName = "marking-dscp-tunnel-stats-val"
    markingDscpTunnelStatsVal.EntityData.BundleName = "cisco_ios_xe"
    markingDscpTunnelStatsVal.EntityData.ParentYangName = "marking-stats"
    markingDscpTunnelStatsVal.EntityData.SegmentPath = "marking-dscp-tunnel-stats-val"
    markingDscpTunnelStatsVal.EntityData.AbsolutePath = "Cisco-IOS-XE-interfaces-oper:interfaces/interface/diffserv-info/diffserv-target-classifier-stats/marking-stats/" + markingDscpTunnelStatsVal.EntityData.SegmentPath
    markingDscpTunnelStatsVal.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    markingDscpTunnelStatsVal.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    markingDscpTunnelStatsVal.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    markingDscpTunnelStatsVal.EntityData.Children = types.NewOrderedMap()
    markingDscpTunnelStatsVal.EntityData.Leafs = types.NewOrderedMap()
    markingDscpTunnelStatsVal.EntityData.Leafs.Append("dscp-val", types.YLeaf{"DscpVal", markingDscpTunnelStatsVal.DscpVal})
    markingDscpTunnelStatsVal.EntityData.Leafs.Append("marked-pkts", types.YLeaf{"MarkedPkts", markingDscpTunnelStatsVal.MarkedPkts})

    markingDscpTunnelStatsVal.EntityData.YListKeys = []string {}

    return &(markingDscpTunnelStatsVal.EntityData)
}

// Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingCosStatsVal
// Statistics for set cos
type Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingCosStatsVal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // cos value. The type is interface{} with range: 0..4294967295.
    CosVal interface{}

    // Number of packets been marked. The type is interface{} with range:
    // 0..18446744073709551615.
    MarkedPkts interface{}
}

func (markingCosStatsVal *Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingCosStatsVal) GetEntityData() *types.CommonEntityData {
    markingCosStatsVal.EntityData.YFilter = markingCosStatsVal.YFilter
    markingCosStatsVal.EntityData.YangName = "marking-cos-stats-val"
    markingCosStatsVal.EntityData.BundleName = "cisco_ios_xe"
    markingCosStatsVal.EntityData.ParentYangName = "marking-stats"
    markingCosStatsVal.EntityData.SegmentPath = "marking-cos-stats-val"
    markingCosStatsVal.EntityData.AbsolutePath = "Cisco-IOS-XE-interfaces-oper:interfaces/interface/diffserv-info/diffserv-target-classifier-stats/marking-stats/" + markingCosStatsVal.EntityData.SegmentPath
    markingCosStatsVal.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    markingCosStatsVal.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    markingCosStatsVal.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    markingCosStatsVal.EntityData.Children = types.NewOrderedMap()
    markingCosStatsVal.EntityData.Leafs = types.NewOrderedMap()
    markingCosStatsVal.EntityData.Leafs.Append("cos-val", types.YLeaf{"CosVal", markingCosStatsVal.CosVal})
    markingCosStatsVal.EntityData.Leafs.Append("marked-pkts", types.YLeaf{"MarkedPkts", markingCosStatsVal.MarkedPkts})

    markingCosStatsVal.EntityData.YListKeys = []string {}

    return &(markingCosStatsVal.EntityData)
}

// Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingCosInnerStatsVal
// Statistics for set cos-inner
type Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingCosInnerStatsVal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // cos inner value. The type is interface{} with range: 0..4294967295.
    CosInnerVal interface{}

    // Number of packets been marked. The type is interface{} with range:
    // 0..18446744073709551615.
    MarkedPkts interface{}
}

func (markingCosInnerStatsVal *Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingCosInnerStatsVal) GetEntityData() *types.CommonEntityData {
    markingCosInnerStatsVal.EntityData.YFilter = markingCosInnerStatsVal.YFilter
    markingCosInnerStatsVal.EntityData.YangName = "marking-cos-inner-stats-val"
    markingCosInnerStatsVal.EntityData.BundleName = "cisco_ios_xe"
    markingCosInnerStatsVal.EntityData.ParentYangName = "marking-stats"
    markingCosInnerStatsVal.EntityData.SegmentPath = "marking-cos-inner-stats-val"
    markingCosInnerStatsVal.EntityData.AbsolutePath = "Cisco-IOS-XE-interfaces-oper:interfaces/interface/diffserv-info/diffserv-target-classifier-stats/marking-stats/" + markingCosInnerStatsVal.EntityData.SegmentPath
    markingCosInnerStatsVal.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    markingCosInnerStatsVal.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    markingCosInnerStatsVal.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    markingCosInnerStatsVal.EntityData.Children = types.NewOrderedMap()
    markingCosInnerStatsVal.EntityData.Leafs = types.NewOrderedMap()
    markingCosInnerStatsVal.EntityData.Leafs.Append("cos-inner-val", types.YLeaf{"CosInnerVal", markingCosInnerStatsVal.CosInnerVal})
    markingCosInnerStatsVal.EntityData.Leafs.Append("marked-pkts", types.YLeaf{"MarkedPkts", markingCosInnerStatsVal.MarkedPkts})

    markingCosInnerStatsVal.EntityData.YListKeys = []string {}

    return &(markingCosInnerStatsVal.EntityData)
}

// Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingDiscardClassStatsVal
// Statistics for set discard-class
type Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingDiscardClassStatsVal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // discard-class value. The type is interface{} with range: 0..4294967295.
    DiscClassVal interface{}

    // Number of packets been marked. The type is interface{} with range:
    // 0..18446744073709551615.
    MarkedPkts interface{}
}

func (markingDiscardClassStatsVal *Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingDiscardClassStatsVal) GetEntityData() *types.CommonEntityData {
    markingDiscardClassStatsVal.EntityData.YFilter = markingDiscardClassStatsVal.YFilter
    markingDiscardClassStatsVal.EntityData.YangName = "marking-discard-class-stats-val"
    markingDiscardClassStatsVal.EntityData.BundleName = "cisco_ios_xe"
    markingDiscardClassStatsVal.EntityData.ParentYangName = "marking-stats"
    markingDiscardClassStatsVal.EntityData.SegmentPath = "marking-discard-class-stats-val"
    markingDiscardClassStatsVal.EntityData.AbsolutePath = "Cisco-IOS-XE-interfaces-oper:interfaces/interface/diffserv-info/diffserv-target-classifier-stats/marking-stats/" + markingDiscardClassStatsVal.EntityData.SegmentPath
    markingDiscardClassStatsVal.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    markingDiscardClassStatsVal.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    markingDiscardClassStatsVal.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    markingDiscardClassStatsVal.EntityData.Children = types.NewOrderedMap()
    markingDiscardClassStatsVal.EntityData.Leafs = types.NewOrderedMap()
    markingDiscardClassStatsVal.EntityData.Leafs.Append("disc-class-val", types.YLeaf{"DiscClassVal", markingDiscardClassStatsVal.DiscClassVal})
    markingDiscardClassStatsVal.EntityData.Leafs.Append("marked-pkts", types.YLeaf{"MarkedPkts", markingDiscardClassStatsVal.MarkedPkts})

    markingDiscardClassStatsVal.EntityData.YListKeys = []string {}

    return &(markingDiscardClassStatsVal.EntityData)
}

// Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingQosGrpStatsVal
// Statistics for set qos-group
type Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingQosGrpStatsVal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // qos group value. The type is interface{} with range: 0..4294967295.
    QosGrpVal interface{}

    // Number of packets been marked. The type is interface{} with range:
    // 0..18446744073709551615.
    MarkedPkts interface{}
}

func (markingQosGrpStatsVal *Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingQosGrpStatsVal) GetEntityData() *types.CommonEntityData {
    markingQosGrpStatsVal.EntityData.YFilter = markingQosGrpStatsVal.YFilter
    markingQosGrpStatsVal.EntityData.YangName = "marking-qos-grp-stats-val"
    markingQosGrpStatsVal.EntityData.BundleName = "cisco_ios_xe"
    markingQosGrpStatsVal.EntityData.ParentYangName = "marking-stats"
    markingQosGrpStatsVal.EntityData.SegmentPath = "marking-qos-grp-stats-val"
    markingQosGrpStatsVal.EntityData.AbsolutePath = "Cisco-IOS-XE-interfaces-oper:interfaces/interface/diffserv-info/diffserv-target-classifier-stats/marking-stats/" + markingQosGrpStatsVal.EntityData.SegmentPath
    markingQosGrpStatsVal.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    markingQosGrpStatsVal.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    markingQosGrpStatsVal.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    markingQosGrpStatsVal.EntityData.Children = types.NewOrderedMap()
    markingQosGrpStatsVal.EntityData.Leafs = types.NewOrderedMap()
    markingQosGrpStatsVal.EntityData.Leafs.Append("qos-grp-val", types.YLeaf{"QosGrpVal", markingQosGrpStatsVal.QosGrpVal})
    markingQosGrpStatsVal.EntityData.Leafs.Append("marked-pkts", types.YLeaf{"MarkedPkts", markingQosGrpStatsVal.MarkedPkts})

    markingQosGrpStatsVal.EntityData.YListKeys = []string {}

    return &(markingQosGrpStatsVal.EntityData)
}

// Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingPrecStatsVal
// Statistics for set precedence
type Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingPrecStatsVal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // precedence value. The type is interface{} with range: 0..4294967295.
    PrecVal interface{}

    // Number of packets been marked. The type is interface{} with range:
    // 0..18446744073709551615.
    MarkedPkts interface{}
}

func (markingPrecStatsVal *Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingPrecStatsVal) GetEntityData() *types.CommonEntityData {
    markingPrecStatsVal.EntityData.YFilter = markingPrecStatsVal.YFilter
    markingPrecStatsVal.EntityData.YangName = "marking-prec-stats-val"
    markingPrecStatsVal.EntityData.BundleName = "cisco_ios_xe"
    markingPrecStatsVal.EntityData.ParentYangName = "marking-stats"
    markingPrecStatsVal.EntityData.SegmentPath = "marking-prec-stats-val"
    markingPrecStatsVal.EntityData.AbsolutePath = "Cisco-IOS-XE-interfaces-oper:interfaces/interface/diffserv-info/diffserv-target-classifier-stats/marking-stats/" + markingPrecStatsVal.EntityData.SegmentPath
    markingPrecStatsVal.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    markingPrecStatsVal.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    markingPrecStatsVal.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    markingPrecStatsVal.EntityData.Children = types.NewOrderedMap()
    markingPrecStatsVal.EntityData.Leafs = types.NewOrderedMap()
    markingPrecStatsVal.EntityData.Leafs.Append("prec-val", types.YLeaf{"PrecVal", markingPrecStatsVal.PrecVal})
    markingPrecStatsVal.EntityData.Leafs.Append("marked-pkts", types.YLeaf{"MarkedPkts", markingPrecStatsVal.MarkedPkts})

    markingPrecStatsVal.EntityData.YListKeys = []string {}

    return &(markingPrecStatsVal.EntityData)
}

// Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingPrecTunnelStatsVal
// Statistics for set precedence tunnel
type Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingPrecTunnelStatsVal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // precedence value. The type is interface{} with range: 0..4294967295.
    PrecVal interface{}

    // Number of packets been marked. The type is interface{} with range:
    // 0..18446744073709551615.
    MarkedPkts interface{}
}

func (markingPrecTunnelStatsVal *Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingPrecTunnelStatsVal) GetEntityData() *types.CommonEntityData {
    markingPrecTunnelStatsVal.EntityData.YFilter = markingPrecTunnelStatsVal.YFilter
    markingPrecTunnelStatsVal.EntityData.YangName = "marking-prec-tunnel-stats-val"
    markingPrecTunnelStatsVal.EntityData.BundleName = "cisco_ios_xe"
    markingPrecTunnelStatsVal.EntityData.ParentYangName = "marking-stats"
    markingPrecTunnelStatsVal.EntityData.SegmentPath = "marking-prec-tunnel-stats-val"
    markingPrecTunnelStatsVal.EntityData.AbsolutePath = "Cisco-IOS-XE-interfaces-oper:interfaces/interface/diffserv-info/diffserv-target-classifier-stats/marking-stats/" + markingPrecTunnelStatsVal.EntityData.SegmentPath
    markingPrecTunnelStatsVal.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    markingPrecTunnelStatsVal.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    markingPrecTunnelStatsVal.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    markingPrecTunnelStatsVal.EntityData.Children = types.NewOrderedMap()
    markingPrecTunnelStatsVal.EntityData.Leafs = types.NewOrderedMap()
    markingPrecTunnelStatsVal.EntityData.Leafs.Append("prec-val", types.YLeaf{"PrecVal", markingPrecTunnelStatsVal.PrecVal})
    markingPrecTunnelStatsVal.EntityData.Leafs.Append("marked-pkts", types.YLeaf{"MarkedPkts", markingPrecTunnelStatsVal.MarkedPkts})

    markingPrecTunnelStatsVal.EntityData.YListKeys = []string {}

    return &(markingPrecTunnelStatsVal.EntityData)
}

// Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingMplsExpImpStatsVal
// Statistics for set mpls exp imposition
type Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingMplsExpImpStatsVal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // mpls exp value. The type is interface{} with range: 0..4294967295.
    MplsExpImpVal interface{}

    // Number of packets been marked. The type is interface{} with range:
    // 0..18446744073709551615.
    MarkedPkts interface{}
}

func (markingMplsExpImpStatsVal *Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingMplsExpImpStatsVal) GetEntityData() *types.CommonEntityData {
    markingMplsExpImpStatsVal.EntityData.YFilter = markingMplsExpImpStatsVal.YFilter
    markingMplsExpImpStatsVal.EntityData.YangName = "marking-mpls-exp-imp-stats-val"
    markingMplsExpImpStatsVal.EntityData.BundleName = "cisco_ios_xe"
    markingMplsExpImpStatsVal.EntityData.ParentYangName = "marking-stats"
    markingMplsExpImpStatsVal.EntityData.SegmentPath = "marking-mpls-exp-imp-stats-val"
    markingMplsExpImpStatsVal.EntityData.AbsolutePath = "Cisco-IOS-XE-interfaces-oper:interfaces/interface/diffserv-info/diffserv-target-classifier-stats/marking-stats/" + markingMplsExpImpStatsVal.EntityData.SegmentPath
    markingMplsExpImpStatsVal.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    markingMplsExpImpStatsVal.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    markingMplsExpImpStatsVal.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    markingMplsExpImpStatsVal.EntityData.Children = types.NewOrderedMap()
    markingMplsExpImpStatsVal.EntityData.Leafs = types.NewOrderedMap()
    markingMplsExpImpStatsVal.EntityData.Leafs.Append("mpls-exp-imp-val", types.YLeaf{"MplsExpImpVal", markingMplsExpImpStatsVal.MplsExpImpVal})
    markingMplsExpImpStatsVal.EntityData.Leafs.Append("marked-pkts", types.YLeaf{"MarkedPkts", markingMplsExpImpStatsVal.MarkedPkts})

    markingMplsExpImpStatsVal.EntityData.YListKeys = []string {}

    return &(markingMplsExpImpStatsVal.EntityData)
}

// Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingMplsExpTopStatsVal
// Statistics for set mpls exp topmost
type Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingMplsExpTopStatsVal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // mpls exp value. The type is interface{} with range: 0..4294967295.
    MplsExpTopVal interface{}

    // Number of packets been marked. The type is interface{} with range:
    // 0..18446744073709551615.
    MarkedPkts interface{}
}

func (markingMplsExpTopStatsVal *Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingMplsExpTopStatsVal) GetEntityData() *types.CommonEntityData {
    markingMplsExpTopStatsVal.EntityData.YFilter = markingMplsExpTopStatsVal.YFilter
    markingMplsExpTopStatsVal.EntityData.YangName = "marking-mpls-exp-top-stats-val"
    markingMplsExpTopStatsVal.EntityData.BundleName = "cisco_ios_xe"
    markingMplsExpTopStatsVal.EntityData.ParentYangName = "marking-stats"
    markingMplsExpTopStatsVal.EntityData.SegmentPath = "marking-mpls-exp-top-stats-val"
    markingMplsExpTopStatsVal.EntityData.AbsolutePath = "Cisco-IOS-XE-interfaces-oper:interfaces/interface/diffserv-info/diffserv-target-classifier-stats/marking-stats/" + markingMplsExpTopStatsVal.EntityData.SegmentPath
    markingMplsExpTopStatsVal.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    markingMplsExpTopStatsVal.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    markingMplsExpTopStatsVal.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    markingMplsExpTopStatsVal.EntityData.Children = types.NewOrderedMap()
    markingMplsExpTopStatsVal.EntityData.Leafs = types.NewOrderedMap()
    markingMplsExpTopStatsVal.EntityData.Leafs.Append("mpls-exp-top-val", types.YLeaf{"MplsExpTopVal", markingMplsExpTopStatsVal.MplsExpTopVal})
    markingMplsExpTopStatsVal.EntityData.Leafs.Append("marked-pkts", types.YLeaf{"MarkedPkts", markingMplsExpTopStatsVal.MarkedPkts})

    markingMplsExpTopStatsVal.EntityData.YListKeys = []string {}

    return &(markingMplsExpTopStatsVal.EntityData)
}

// Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingFrDeStatsVal
// Statistics for set fr-de
type Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingFrDeStatsVal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // fr de set or not. The type is bool.
    FrDe interface{}

    // Number of packets been marked. The type is interface{} with range:
    // 0..18446744073709551615.
    MarkedPkts interface{}
}

func (markingFrDeStatsVal *Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingFrDeStatsVal) GetEntityData() *types.CommonEntityData {
    markingFrDeStatsVal.EntityData.YFilter = markingFrDeStatsVal.YFilter
    markingFrDeStatsVal.EntityData.YangName = "marking-fr-de-stats-val"
    markingFrDeStatsVal.EntityData.BundleName = "cisco_ios_xe"
    markingFrDeStatsVal.EntityData.ParentYangName = "marking-stats"
    markingFrDeStatsVal.EntityData.SegmentPath = "marking-fr-de-stats-val"
    markingFrDeStatsVal.EntityData.AbsolutePath = "Cisco-IOS-XE-interfaces-oper:interfaces/interface/diffserv-info/diffserv-target-classifier-stats/marking-stats/" + markingFrDeStatsVal.EntityData.SegmentPath
    markingFrDeStatsVal.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    markingFrDeStatsVal.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    markingFrDeStatsVal.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    markingFrDeStatsVal.EntityData.Children = types.NewOrderedMap()
    markingFrDeStatsVal.EntityData.Leafs = types.NewOrderedMap()
    markingFrDeStatsVal.EntityData.Leafs.Append("fr-de", types.YLeaf{"FrDe", markingFrDeStatsVal.FrDe})
    markingFrDeStatsVal.EntityData.Leafs.Append("marked-pkts", types.YLeaf{"MarkedPkts", markingFrDeStatsVal.MarkedPkts})

    markingFrDeStatsVal.EntityData.YListKeys = []string {}

    return &(markingFrDeStatsVal.EntityData)
}

// Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingFrFecnBecnStatsVal
// Statistics for set fr-fecn-becn
type Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingFrFecnBecnStatsVal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // fecn becn value. qos:percent-value-1to100. The type is interface{} with
    // range: 0..4294967295.
    FecnBecnVal interface{}

    // Number of packets been marked. The type is interface{} with range:
    // 0..18446744073709551615.
    MarkedPkts interface{}
}

func (markingFrFecnBecnStatsVal *Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingFrFecnBecnStatsVal) GetEntityData() *types.CommonEntityData {
    markingFrFecnBecnStatsVal.EntityData.YFilter = markingFrFecnBecnStatsVal.YFilter
    markingFrFecnBecnStatsVal.EntityData.YangName = "marking-fr-fecn-becn-stats-val"
    markingFrFecnBecnStatsVal.EntityData.BundleName = "cisco_ios_xe"
    markingFrFecnBecnStatsVal.EntityData.ParentYangName = "marking-stats"
    markingFrFecnBecnStatsVal.EntityData.SegmentPath = "marking-fr-fecn-becn-stats-val"
    markingFrFecnBecnStatsVal.EntityData.AbsolutePath = "Cisco-IOS-XE-interfaces-oper:interfaces/interface/diffserv-info/diffserv-target-classifier-stats/marking-stats/" + markingFrFecnBecnStatsVal.EntityData.SegmentPath
    markingFrFecnBecnStatsVal.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    markingFrFecnBecnStatsVal.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    markingFrFecnBecnStatsVal.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    markingFrFecnBecnStatsVal.EntityData.Children = types.NewOrderedMap()
    markingFrFecnBecnStatsVal.EntityData.Leafs = types.NewOrderedMap()
    markingFrFecnBecnStatsVal.EntityData.Leafs.Append("fecn-becn-val", types.YLeaf{"FecnBecnVal", markingFrFecnBecnStatsVal.FecnBecnVal})
    markingFrFecnBecnStatsVal.EntityData.Leafs.Append("marked-pkts", types.YLeaf{"MarkedPkts", markingFrFecnBecnStatsVal.MarkedPkts})

    markingFrFecnBecnStatsVal.EntityData.YListKeys = []string {}

    return &(markingFrFecnBecnStatsVal.EntityData)
}

// Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingAtmClpStatsVal
// Statistics for set atm-clp
type Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingAtmClpStatsVal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // atm clp value. The type is interface{} with range: 0..255.
    AtmClpVal interface{}

    // Number of packets been marked. The type is interface{} with range:
    // 0..18446744073709551615.
    MarkedPkts interface{}
}

func (markingAtmClpStatsVal *Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingAtmClpStatsVal) GetEntityData() *types.CommonEntityData {
    markingAtmClpStatsVal.EntityData.YFilter = markingAtmClpStatsVal.YFilter
    markingAtmClpStatsVal.EntityData.YangName = "marking-atm-clp-stats-val"
    markingAtmClpStatsVal.EntityData.BundleName = "cisco_ios_xe"
    markingAtmClpStatsVal.EntityData.ParentYangName = "marking-stats"
    markingAtmClpStatsVal.EntityData.SegmentPath = "marking-atm-clp-stats-val"
    markingAtmClpStatsVal.EntityData.AbsolutePath = "Cisco-IOS-XE-interfaces-oper:interfaces/interface/diffserv-info/diffserv-target-classifier-stats/marking-stats/" + markingAtmClpStatsVal.EntityData.SegmentPath
    markingAtmClpStatsVal.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    markingAtmClpStatsVal.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    markingAtmClpStatsVal.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    markingAtmClpStatsVal.EntityData.Children = types.NewOrderedMap()
    markingAtmClpStatsVal.EntityData.Leafs = types.NewOrderedMap()
    markingAtmClpStatsVal.EntityData.Leafs.Append("atm-clp-val", types.YLeaf{"AtmClpVal", markingAtmClpStatsVal.AtmClpVal})
    markingAtmClpStatsVal.EntityData.Leafs.Append("marked-pkts", types.YLeaf{"MarkedPkts", markingAtmClpStatsVal.MarkedPkts})

    markingAtmClpStatsVal.EntityData.YListKeys = []string {}

    return &(markingAtmClpStatsVal.EntityData)
}

// Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingVlanInnerStatsVal
// Statistics for set vlan-inner
type Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingVlanInnerStatsVal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // vlan value. The type is interface{} with range: 0..4294967295.
    VlanInnerVal interface{}

    // Number of packets been marked. The type is interface{} with range:
    // 0..18446744073709551615.
    MarkedPkts interface{}
}

func (markingVlanInnerStatsVal *Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingVlanInnerStatsVal) GetEntityData() *types.CommonEntityData {
    markingVlanInnerStatsVal.EntityData.YFilter = markingVlanInnerStatsVal.YFilter
    markingVlanInnerStatsVal.EntityData.YangName = "marking-vlan-inner-stats-val"
    markingVlanInnerStatsVal.EntityData.BundleName = "cisco_ios_xe"
    markingVlanInnerStatsVal.EntityData.ParentYangName = "marking-stats"
    markingVlanInnerStatsVal.EntityData.SegmentPath = "marking-vlan-inner-stats-val"
    markingVlanInnerStatsVal.EntityData.AbsolutePath = "Cisco-IOS-XE-interfaces-oper:interfaces/interface/diffserv-info/diffserv-target-classifier-stats/marking-stats/" + markingVlanInnerStatsVal.EntityData.SegmentPath
    markingVlanInnerStatsVal.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    markingVlanInnerStatsVal.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    markingVlanInnerStatsVal.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    markingVlanInnerStatsVal.EntityData.Children = types.NewOrderedMap()
    markingVlanInnerStatsVal.EntityData.Leafs = types.NewOrderedMap()
    markingVlanInnerStatsVal.EntityData.Leafs.Append("vlan-inner-val", types.YLeaf{"VlanInnerVal", markingVlanInnerStatsVal.VlanInnerVal})
    markingVlanInnerStatsVal.EntityData.Leafs.Append("marked-pkts", types.YLeaf{"MarkedPkts", markingVlanInnerStatsVal.MarkedPkts})

    markingVlanInnerStatsVal.EntityData.YListKeys = []string {}

    return &(markingVlanInnerStatsVal.EntityData)
}

// Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingDeiStatsVal
// Statistics for set dei
type Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingDeiStatsVal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // dei value. The type is interface{} with range: 0..4294967295.
    DeiImpValue interface{}

    // Number of packets been marked. The type is interface{} with range:
    // 0..18446744073709551615.
    MarkedPkts interface{}
}

func (markingDeiStatsVal *Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingDeiStatsVal) GetEntityData() *types.CommonEntityData {
    markingDeiStatsVal.EntityData.YFilter = markingDeiStatsVal.YFilter
    markingDeiStatsVal.EntityData.YangName = "marking-dei-stats-val"
    markingDeiStatsVal.EntityData.BundleName = "cisco_ios_xe"
    markingDeiStatsVal.EntityData.ParentYangName = "marking-stats"
    markingDeiStatsVal.EntityData.SegmentPath = "marking-dei-stats-val"
    markingDeiStatsVal.EntityData.AbsolutePath = "Cisco-IOS-XE-interfaces-oper:interfaces/interface/diffserv-info/diffserv-target-classifier-stats/marking-stats/" + markingDeiStatsVal.EntityData.SegmentPath
    markingDeiStatsVal.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    markingDeiStatsVal.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    markingDeiStatsVal.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    markingDeiStatsVal.EntityData.Children = types.NewOrderedMap()
    markingDeiStatsVal.EntityData.Leafs = types.NewOrderedMap()
    markingDeiStatsVal.EntityData.Leafs.Append("dei-imp-value", types.YLeaf{"DeiImpValue", markingDeiStatsVal.DeiImpValue})
    markingDeiStatsVal.EntityData.Leafs.Append("marked-pkts", types.YLeaf{"MarkedPkts", markingDeiStatsVal.MarkedPkts})

    markingDeiStatsVal.EntityData.YListKeys = []string {}

    return &(markingDeiStatsVal.EntityData)
}

// Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingDeiImpStatsVal
// Statistics for set dei-imposition
type Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingDeiImpStatsVal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // dei value. The type is interface{} with range: 0..4294967295.
    DeiImpValue interface{}

    // Number of packets been marked. The type is interface{} with range:
    // 0..18446744073709551615.
    MarkedPkts interface{}
}

func (markingDeiImpStatsVal *Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingDeiImpStatsVal) GetEntityData() *types.CommonEntityData {
    markingDeiImpStatsVal.EntityData.YFilter = markingDeiImpStatsVal.YFilter
    markingDeiImpStatsVal.EntityData.YangName = "marking-dei-imp-stats-val"
    markingDeiImpStatsVal.EntityData.BundleName = "cisco_ios_xe"
    markingDeiImpStatsVal.EntityData.ParentYangName = "marking-stats"
    markingDeiImpStatsVal.EntityData.SegmentPath = "marking-dei-imp-stats-val"
    markingDeiImpStatsVal.EntityData.AbsolutePath = "Cisco-IOS-XE-interfaces-oper:interfaces/interface/diffserv-info/diffserv-target-classifier-stats/marking-stats/" + markingDeiImpStatsVal.EntityData.SegmentPath
    markingDeiImpStatsVal.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    markingDeiImpStatsVal.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    markingDeiImpStatsVal.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    markingDeiImpStatsVal.EntityData.Children = types.NewOrderedMap()
    markingDeiImpStatsVal.EntityData.Leafs = types.NewOrderedMap()
    markingDeiImpStatsVal.EntityData.Leafs.Append("dei-imp-value", types.YLeaf{"DeiImpValue", markingDeiImpStatsVal.DeiImpValue})
    markingDeiImpStatsVal.EntityData.Leafs.Append("marked-pkts", types.YLeaf{"MarkedPkts", markingDeiImpStatsVal.MarkedPkts})

    markingDeiImpStatsVal.EntityData.YListKeys = []string {}

    return &(markingDeiImpStatsVal.EntityData)
}

// Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingSrpPriorityStatsVal
// Statistics for set srp-priority
type Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingSrpPriorityStatsVal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // srp priority value. The type is interface{} with range: 0..255.
    SrpPriorityValue interface{}

    // Number of packets been marked. The type is interface{} with range:
    // 0..18446744073709551615.
    MarkedPkts interface{}
}

func (markingSrpPriorityStatsVal *Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingSrpPriorityStatsVal) GetEntityData() *types.CommonEntityData {
    markingSrpPriorityStatsVal.EntityData.YFilter = markingSrpPriorityStatsVal.YFilter
    markingSrpPriorityStatsVal.EntityData.YangName = "marking-srp-priority-stats-val"
    markingSrpPriorityStatsVal.EntityData.BundleName = "cisco_ios_xe"
    markingSrpPriorityStatsVal.EntityData.ParentYangName = "marking-stats"
    markingSrpPriorityStatsVal.EntityData.SegmentPath = "marking-srp-priority-stats-val"
    markingSrpPriorityStatsVal.EntityData.AbsolutePath = "Cisco-IOS-XE-interfaces-oper:interfaces/interface/diffserv-info/diffserv-target-classifier-stats/marking-stats/" + markingSrpPriorityStatsVal.EntityData.SegmentPath
    markingSrpPriorityStatsVal.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    markingSrpPriorityStatsVal.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    markingSrpPriorityStatsVal.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    markingSrpPriorityStatsVal.EntityData.Children = types.NewOrderedMap()
    markingSrpPriorityStatsVal.EntityData.Leafs = types.NewOrderedMap()
    markingSrpPriorityStatsVal.EntityData.Leafs.Append("srp-priority-value", types.YLeaf{"SrpPriorityValue", markingSrpPriorityStatsVal.SrpPriorityValue})
    markingSrpPriorityStatsVal.EntityData.Leafs.Append("marked-pkts", types.YLeaf{"MarkedPkts", markingSrpPriorityStatsVal.MarkedPkts})

    markingSrpPriorityStatsVal.EntityData.YListKeys = []string {}

    return &(markingSrpPriorityStatsVal.EntityData)
}

// Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingWlanUserPriorityStatsVal
// Statistics for set wlan-user-priority
type Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingWlanUserPriorityStatsVal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // wlan user priority value. The type is interface{} with range: 0..255.
    WlanUserPriorityValue interface{}

    // Number of packets been marked. The type is interface{} with range:
    // 0..18446744073709551615.
    MarkedPkts interface{}
}

func (markingWlanUserPriorityStatsVal *Interfaces_Interface_DiffservInfo_DiffservTargetClassifierStats_MarkingStats_MarkingWlanUserPriorityStatsVal) GetEntityData() *types.CommonEntityData {
    markingWlanUserPriorityStatsVal.EntityData.YFilter = markingWlanUserPriorityStatsVal.YFilter
    markingWlanUserPriorityStatsVal.EntityData.YangName = "marking-wlan-user-priority-stats-val"
    markingWlanUserPriorityStatsVal.EntityData.BundleName = "cisco_ios_xe"
    markingWlanUserPriorityStatsVal.EntityData.ParentYangName = "marking-stats"
    markingWlanUserPriorityStatsVal.EntityData.SegmentPath = "marking-wlan-user-priority-stats-val"
    markingWlanUserPriorityStatsVal.EntityData.AbsolutePath = "Cisco-IOS-XE-interfaces-oper:interfaces/interface/diffserv-info/diffserv-target-classifier-stats/marking-stats/" + markingWlanUserPriorityStatsVal.EntityData.SegmentPath
    markingWlanUserPriorityStatsVal.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    markingWlanUserPriorityStatsVal.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    markingWlanUserPriorityStatsVal.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    markingWlanUserPriorityStatsVal.EntityData.Children = types.NewOrderedMap()
    markingWlanUserPriorityStatsVal.EntityData.Leafs = types.NewOrderedMap()
    markingWlanUserPriorityStatsVal.EntityData.Leafs.Append("wlan-user-priority-value", types.YLeaf{"WlanUserPriorityValue", markingWlanUserPriorityStatsVal.WlanUserPriorityValue})
    markingWlanUserPriorityStatsVal.EntityData.Leafs.Append("marked-pkts", types.YLeaf{"MarkedPkts", markingWlanUserPriorityStatsVal.MarkedPkts})

    markingWlanUserPriorityStatsVal.EntityData.YListKeys = []string {}

    return &(markingWlanUserPriorityStatsVal.EntityData)
}

// Interfaces_Interface_DiffservInfo_PriorityOperList
// Statistics for aggregrate priority per policy instance
type Interfaces_Interface_DiffservInfo_PriorityOperList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Priority Level, 0 means no priority level set. The
    // type is interface{} with range: 0..65535.
    PriorityLevel interface{}

    // Counters in aggregate priority.
    AggPriorityStats Interfaces_Interface_DiffservInfo_PriorityOperList_AggPriorityStats

    // qlimit default threshold.
    QlimitDefaultThresh Interfaces_Interface_DiffservInfo_PriorityOperList_QlimitDefaultThresh

    // cos-based queue limit data. The type is slice of
    // Interfaces_Interface_DiffservInfo_PriorityOperList_QlimitCosThreshList.
    QlimitCosThreshList []*Interfaces_Interface_DiffservInfo_PriorityOperList_QlimitCosThreshList

    // discard-class-based queue limit data. The type is slice of
    // Interfaces_Interface_DiffservInfo_PriorityOperList_QlimitDiscClassThreshList.
    QlimitDiscClassThreshList []*Interfaces_Interface_DiffservInfo_PriorityOperList_QlimitDiscClassThreshList

    // qos-group-based queue limit data. The type is slice of
    // Interfaces_Interface_DiffservInfo_PriorityOperList_QlimitQosGrpThreshList.
    QlimitQosGrpThreshList []*Interfaces_Interface_DiffservInfo_PriorityOperList_QlimitQosGrpThreshList

    // mpls-exp-based queue limit data. The type is slice of
    // Interfaces_Interface_DiffservInfo_PriorityOperList_QlimitMplsExpThreshList.
    QlimitMplsExpThreshList []*Interfaces_Interface_DiffservInfo_PriorityOperList_QlimitMplsExpThreshList

    // queue limit per dscp range. The type is slice of
    // Interfaces_Interface_DiffservInfo_PriorityOperList_QlimitDscpThreshList.
    QlimitDscpThreshList []*Interfaces_Interface_DiffservInfo_PriorityOperList_QlimitDscpThreshList
}

func (priorityOperList *Interfaces_Interface_DiffservInfo_PriorityOperList) GetEntityData() *types.CommonEntityData {
    priorityOperList.EntityData.YFilter = priorityOperList.YFilter
    priorityOperList.EntityData.YangName = "priority-oper-list"
    priorityOperList.EntityData.BundleName = "cisco_ios_xe"
    priorityOperList.EntityData.ParentYangName = "diffserv-info"
    priorityOperList.EntityData.SegmentPath = "priority-oper-list" + types.AddKeyToken(priorityOperList.PriorityLevel, "priority-level")
    priorityOperList.EntityData.AbsolutePath = "Cisco-IOS-XE-interfaces-oper:interfaces/interface/diffserv-info/" + priorityOperList.EntityData.SegmentPath
    priorityOperList.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    priorityOperList.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    priorityOperList.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    priorityOperList.EntityData.Children = types.NewOrderedMap()
    priorityOperList.EntityData.Children.Append("agg-priority-stats", types.YChild{"AggPriorityStats", &priorityOperList.AggPriorityStats})
    priorityOperList.EntityData.Children.Append("qlimit-default-thresh", types.YChild{"QlimitDefaultThresh", &priorityOperList.QlimitDefaultThresh})
    priorityOperList.EntityData.Children.Append("qlimit-cos-thresh-list", types.YChild{"QlimitCosThreshList", nil})
    for i := range priorityOperList.QlimitCosThreshList {
        priorityOperList.EntityData.Children.Append(types.GetSegmentPath(priorityOperList.QlimitCosThreshList[i]), types.YChild{"QlimitCosThreshList", priorityOperList.QlimitCosThreshList[i]})
    }
    priorityOperList.EntityData.Children.Append("qlimit-disc-class-thresh-list", types.YChild{"QlimitDiscClassThreshList", nil})
    for i := range priorityOperList.QlimitDiscClassThreshList {
        priorityOperList.EntityData.Children.Append(types.GetSegmentPath(priorityOperList.QlimitDiscClassThreshList[i]), types.YChild{"QlimitDiscClassThreshList", priorityOperList.QlimitDiscClassThreshList[i]})
    }
    priorityOperList.EntityData.Children.Append("qlimit-qos-grp-thresh-list", types.YChild{"QlimitQosGrpThreshList", nil})
    for i := range priorityOperList.QlimitQosGrpThreshList {
        priorityOperList.EntityData.Children.Append(types.GetSegmentPath(priorityOperList.QlimitQosGrpThreshList[i]), types.YChild{"QlimitQosGrpThreshList", priorityOperList.QlimitQosGrpThreshList[i]})
    }
    priorityOperList.EntityData.Children.Append("qlimit-mpls-exp-thresh-list", types.YChild{"QlimitMplsExpThreshList", nil})
    for i := range priorityOperList.QlimitMplsExpThreshList {
        priorityOperList.EntityData.Children.Append(types.GetSegmentPath(priorityOperList.QlimitMplsExpThreshList[i]), types.YChild{"QlimitMplsExpThreshList", priorityOperList.QlimitMplsExpThreshList[i]})
    }
    priorityOperList.EntityData.Children.Append("qlimit-dscp-thresh-list", types.YChild{"QlimitDscpThreshList", nil})
    for i := range priorityOperList.QlimitDscpThreshList {
        priorityOperList.EntityData.Children.Append(types.GetSegmentPath(priorityOperList.QlimitDscpThreshList[i]), types.YChild{"QlimitDscpThreshList", priorityOperList.QlimitDscpThreshList[i]})
    }
    priorityOperList.EntityData.Leafs = types.NewOrderedMap()
    priorityOperList.EntityData.Leafs.Append("priority-level", types.YLeaf{"PriorityLevel", priorityOperList.PriorityLevel})

    priorityOperList.EntityData.YListKeys = []string {"PriorityLevel"}

    return &(priorityOperList.EntityData)
}

// Interfaces_Interface_DiffservInfo_PriorityOperList_AggPriorityStats
// Counters in aggregate priority
type Interfaces_Interface_DiffservInfo_PriorityOperList_AggPriorityStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of packets transmitted from queue. The type is interface{} with
    // range: 0..18446744073709551615.
    OutputPkts interface{}

    // Number of bytes transmitted from queue. The type is interface{} with range:
    // 0..18446744073709551615.
    OutputBytes interface{}

    // Number of packets currently buffered. The type is interface{} with range:
    // 0..18446744073709551615.
    QueueSizePkts interface{}

    // Number of bytes currently buffered. The type is interface{} with range:
    // 0..18446744073709551615.
    QueueSizeBytes interface{}

    // Total number of packets dropped. The type is interface{} with range:
    // 0..18446744073709551615.
    DropPkts interface{}

    // Total number of bytes dropped. The type is interface{} with range:
    // 0..18446744073709551615.
    DropBytes interface{}

    // Number of packets that were dropped by flow-based fair-queuing
    // (fair-queue). This is a sub-set of drop-pkts. The type is interface{} with
    // range: 0..18446744073709551615.
    DropPktsFlow interface{}

    // Number of packets dropped due to buffers being unavailable system-wide or
    // at the associated interface. This is a sub-set of drop-pkts. The type is
    // interface{} with range: 0..18446744073709551615.
    DropPktsNoBuffer interface{}
}

func (aggPriorityStats *Interfaces_Interface_DiffservInfo_PriorityOperList_AggPriorityStats) GetEntityData() *types.CommonEntityData {
    aggPriorityStats.EntityData.YFilter = aggPriorityStats.YFilter
    aggPriorityStats.EntityData.YangName = "agg-priority-stats"
    aggPriorityStats.EntityData.BundleName = "cisco_ios_xe"
    aggPriorityStats.EntityData.ParentYangName = "priority-oper-list"
    aggPriorityStats.EntityData.SegmentPath = "agg-priority-stats"
    aggPriorityStats.EntityData.AbsolutePath = "Cisco-IOS-XE-interfaces-oper:interfaces/interface/diffserv-info/priority-oper-list/" + aggPriorityStats.EntityData.SegmentPath
    aggPriorityStats.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    aggPriorityStats.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    aggPriorityStats.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    aggPriorityStats.EntityData.Children = types.NewOrderedMap()
    aggPriorityStats.EntityData.Leafs = types.NewOrderedMap()
    aggPriorityStats.EntityData.Leafs.Append("output-pkts", types.YLeaf{"OutputPkts", aggPriorityStats.OutputPkts})
    aggPriorityStats.EntityData.Leafs.Append("output-bytes", types.YLeaf{"OutputBytes", aggPriorityStats.OutputBytes})
    aggPriorityStats.EntityData.Leafs.Append("queue-size-pkts", types.YLeaf{"QueueSizePkts", aggPriorityStats.QueueSizePkts})
    aggPriorityStats.EntityData.Leafs.Append("queue-size-bytes", types.YLeaf{"QueueSizeBytes", aggPriorityStats.QueueSizeBytes})
    aggPriorityStats.EntityData.Leafs.Append("drop-pkts", types.YLeaf{"DropPkts", aggPriorityStats.DropPkts})
    aggPriorityStats.EntityData.Leafs.Append("drop-bytes", types.YLeaf{"DropBytes", aggPriorityStats.DropBytes})
    aggPriorityStats.EntityData.Leafs.Append("drop-pkts-flow", types.YLeaf{"DropPktsFlow", aggPriorityStats.DropPktsFlow})
    aggPriorityStats.EntityData.Leafs.Append("drop-pkts-no-buffer", types.YLeaf{"DropPktsNoBuffer", aggPriorityStats.DropPktsNoBuffer})

    aggPriorityStats.EntityData.YListKeys = []string {}

    return &(aggPriorityStats.EntityData)
}

// Interfaces_Interface_DiffservInfo_PriorityOperList_QlimitDefaultThresh
// qlimit default threshold
type Interfaces_Interface_DiffservInfo_PriorityOperList_QlimitDefaultThresh struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Threshold bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    Bytes interface{}

    // Threshold size unit. The type is interface{} with range: 0..4294967295.
    ThreshSizeMetric interface{}

    // Threshold size basic units. The type is ThreshUnit.
    UnitVal interface{}

    // Threshold interval. The type is interface{} with range:
    // 0..18446744073709551615.
    ThresholdInterval interface{}

    // Threshold units metric. The type is interface{} with range: 0..4294967295.
    ThreshIntervalMetric interface{}

    // Threshold intveral basic units. The type is ThreshUnit.
    IntervalUnitVal interface{}
}

func (qlimitDefaultThresh *Interfaces_Interface_DiffservInfo_PriorityOperList_QlimitDefaultThresh) GetEntityData() *types.CommonEntityData {
    qlimitDefaultThresh.EntityData.YFilter = qlimitDefaultThresh.YFilter
    qlimitDefaultThresh.EntityData.YangName = "qlimit-default-thresh"
    qlimitDefaultThresh.EntityData.BundleName = "cisco_ios_xe"
    qlimitDefaultThresh.EntityData.ParentYangName = "priority-oper-list"
    qlimitDefaultThresh.EntityData.SegmentPath = "qlimit-default-thresh"
    qlimitDefaultThresh.EntityData.AbsolutePath = "Cisco-IOS-XE-interfaces-oper:interfaces/interface/diffserv-info/priority-oper-list/" + qlimitDefaultThresh.EntityData.SegmentPath
    qlimitDefaultThresh.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qlimitDefaultThresh.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qlimitDefaultThresh.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qlimitDefaultThresh.EntityData.Children = types.NewOrderedMap()
    qlimitDefaultThresh.EntityData.Leafs = types.NewOrderedMap()
    qlimitDefaultThresh.EntityData.Leafs.Append("bytes", types.YLeaf{"Bytes", qlimitDefaultThresh.Bytes})
    qlimitDefaultThresh.EntityData.Leafs.Append("thresh-size-metric", types.YLeaf{"ThreshSizeMetric", qlimitDefaultThresh.ThreshSizeMetric})
    qlimitDefaultThresh.EntityData.Leafs.Append("unit-val", types.YLeaf{"UnitVal", qlimitDefaultThresh.UnitVal})
    qlimitDefaultThresh.EntityData.Leafs.Append("threshold-interval", types.YLeaf{"ThresholdInterval", qlimitDefaultThresh.ThresholdInterval})
    qlimitDefaultThresh.EntityData.Leafs.Append("thresh-interval-metric", types.YLeaf{"ThreshIntervalMetric", qlimitDefaultThresh.ThreshIntervalMetric})
    qlimitDefaultThresh.EntityData.Leafs.Append("interval-unit-val", types.YLeaf{"IntervalUnitVal", qlimitDefaultThresh.IntervalUnitVal})

    qlimitDefaultThresh.EntityData.YListKeys = []string {}

    return &(qlimitDefaultThresh.EntityData)
}

// Interfaces_Interface_DiffservInfo_PriorityOperList_QlimitCosThreshList
// cos-based queue limit data
type Interfaces_Interface_DiffservInfo_PriorityOperList_QlimitCosThreshList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Min COS value. The type is interface{} with range:
    // 0..4294967295.
    CosMin interface{}

    // This attribute is a key. Max COS value. The type is interface{} with range:
    // 0..4294967295.
    CosMax interface{}

    // Threshold bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    Bytes interface{}

    // Threshold size unit. The type is interface{} with range: 0..4294967295.
    ThreshSizeMetric interface{}

    // Threshold size basic units. The type is ThreshUnit.
    UnitVal interface{}

    // Threshold interval. The type is interface{} with range:
    // 0..18446744073709551615.
    ThresholdInterval interface{}

    // Threshold units metric. The type is interface{} with range: 0..4294967295.
    ThreshIntervalMetric interface{}

    // Threshold intveral basic units. The type is ThreshUnit.
    IntervalUnitVal interface{}
}

func (qlimitCosThreshList *Interfaces_Interface_DiffservInfo_PriorityOperList_QlimitCosThreshList) GetEntityData() *types.CommonEntityData {
    qlimitCosThreshList.EntityData.YFilter = qlimitCosThreshList.YFilter
    qlimitCosThreshList.EntityData.YangName = "qlimit-cos-thresh-list"
    qlimitCosThreshList.EntityData.BundleName = "cisco_ios_xe"
    qlimitCosThreshList.EntityData.ParentYangName = "priority-oper-list"
    qlimitCosThreshList.EntityData.SegmentPath = "qlimit-cos-thresh-list" + types.AddKeyToken(qlimitCosThreshList.CosMin, "cos-min") + types.AddKeyToken(qlimitCosThreshList.CosMax, "cos-max")
    qlimitCosThreshList.EntityData.AbsolutePath = "Cisco-IOS-XE-interfaces-oper:interfaces/interface/diffserv-info/priority-oper-list/" + qlimitCosThreshList.EntityData.SegmentPath
    qlimitCosThreshList.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qlimitCosThreshList.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qlimitCosThreshList.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qlimitCosThreshList.EntityData.Children = types.NewOrderedMap()
    qlimitCosThreshList.EntityData.Leafs = types.NewOrderedMap()
    qlimitCosThreshList.EntityData.Leafs.Append("cos-min", types.YLeaf{"CosMin", qlimitCosThreshList.CosMin})
    qlimitCosThreshList.EntityData.Leafs.Append("cos-max", types.YLeaf{"CosMax", qlimitCosThreshList.CosMax})
    qlimitCosThreshList.EntityData.Leafs.Append("bytes", types.YLeaf{"Bytes", qlimitCosThreshList.Bytes})
    qlimitCosThreshList.EntityData.Leafs.Append("thresh-size-metric", types.YLeaf{"ThreshSizeMetric", qlimitCosThreshList.ThreshSizeMetric})
    qlimitCosThreshList.EntityData.Leafs.Append("unit-val", types.YLeaf{"UnitVal", qlimitCosThreshList.UnitVal})
    qlimitCosThreshList.EntityData.Leafs.Append("threshold-interval", types.YLeaf{"ThresholdInterval", qlimitCosThreshList.ThresholdInterval})
    qlimitCosThreshList.EntityData.Leafs.Append("thresh-interval-metric", types.YLeaf{"ThreshIntervalMetric", qlimitCosThreshList.ThreshIntervalMetric})
    qlimitCosThreshList.EntityData.Leafs.Append("interval-unit-val", types.YLeaf{"IntervalUnitVal", qlimitCosThreshList.IntervalUnitVal})

    qlimitCosThreshList.EntityData.YListKeys = []string {"CosMin", "CosMax"}

    return &(qlimitCosThreshList.EntityData)
}

// Interfaces_Interface_DiffservInfo_PriorityOperList_QlimitDiscClassThreshList
// discard-class-based queue limit data
type Interfaces_Interface_DiffservInfo_PriorityOperList_QlimitDiscClassThreshList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Minimum value for discard class in the range. The
    // type is interface{} with range: 0..4294967295.
    DiscClassMin interface{}

    // This attribute is a key. Maximum value for discard class in the range. The
    // type is interface{} with range: 0..4294967295.
    DiscClassMax interface{}

    // Threshold bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    Bytes interface{}

    // Threshold size unit. The type is interface{} with range: 0..4294967295.
    ThreshSizeMetric interface{}

    // Threshold size basic units. The type is ThreshUnit.
    UnitVal interface{}

    // Threshold interval. The type is interface{} with range:
    // 0..18446744073709551615.
    ThresholdInterval interface{}

    // Threshold units metric. The type is interface{} with range: 0..4294967295.
    ThreshIntervalMetric interface{}

    // Threshold intveral basic units. The type is ThreshUnit.
    IntervalUnitVal interface{}
}

func (qlimitDiscClassThreshList *Interfaces_Interface_DiffservInfo_PriorityOperList_QlimitDiscClassThreshList) GetEntityData() *types.CommonEntityData {
    qlimitDiscClassThreshList.EntityData.YFilter = qlimitDiscClassThreshList.YFilter
    qlimitDiscClassThreshList.EntityData.YangName = "qlimit-disc-class-thresh-list"
    qlimitDiscClassThreshList.EntityData.BundleName = "cisco_ios_xe"
    qlimitDiscClassThreshList.EntityData.ParentYangName = "priority-oper-list"
    qlimitDiscClassThreshList.EntityData.SegmentPath = "qlimit-disc-class-thresh-list" + types.AddKeyToken(qlimitDiscClassThreshList.DiscClassMin, "disc-class-min") + types.AddKeyToken(qlimitDiscClassThreshList.DiscClassMax, "disc-class-max")
    qlimitDiscClassThreshList.EntityData.AbsolutePath = "Cisco-IOS-XE-interfaces-oper:interfaces/interface/diffserv-info/priority-oper-list/" + qlimitDiscClassThreshList.EntityData.SegmentPath
    qlimitDiscClassThreshList.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qlimitDiscClassThreshList.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qlimitDiscClassThreshList.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qlimitDiscClassThreshList.EntityData.Children = types.NewOrderedMap()
    qlimitDiscClassThreshList.EntityData.Leafs = types.NewOrderedMap()
    qlimitDiscClassThreshList.EntityData.Leafs.Append("disc-class-min", types.YLeaf{"DiscClassMin", qlimitDiscClassThreshList.DiscClassMin})
    qlimitDiscClassThreshList.EntityData.Leafs.Append("disc-class-max", types.YLeaf{"DiscClassMax", qlimitDiscClassThreshList.DiscClassMax})
    qlimitDiscClassThreshList.EntityData.Leafs.Append("bytes", types.YLeaf{"Bytes", qlimitDiscClassThreshList.Bytes})
    qlimitDiscClassThreshList.EntityData.Leafs.Append("thresh-size-metric", types.YLeaf{"ThreshSizeMetric", qlimitDiscClassThreshList.ThreshSizeMetric})
    qlimitDiscClassThreshList.EntityData.Leafs.Append("unit-val", types.YLeaf{"UnitVal", qlimitDiscClassThreshList.UnitVal})
    qlimitDiscClassThreshList.EntityData.Leafs.Append("threshold-interval", types.YLeaf{"ThresholdInterval", qlimitDiscClassThreshList.ThresholdInterval})
    qlimitDiscClassThreshList.EntityData.Leafs.Append("thresh-interval-metric", types.YLeaf{"ThreshIntervalMetric", qlimitDiscClassThreshList.ThreshIntervalMetric})
    qlimitDiscClassThreshList.EntityData.Leafs.Append("interval-unit-val", types.YLeaf{"IntervalUnitVal", qlimitDiscClassThreshList.IntervalUnitVal})

    qlimitDiscClassThreshList.EntityData.YListKeys = []string {"DiscClassMin", "DiscClassMax"}

    return &(qlimitDiscClassThreshList.EntityData)
}

// Interfaces_Interface_DiffservInfo_PriorityOperList_QlimitQosGrpThreshList
// qos-group-based queue limit data
type Interfaces_Interface_DiffservInfo_PriorityOperList_QlimitQosGrpThreshList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Specifies the minimum value range from 0 to used
    // to identify a QoS group value. The type is interface{} with range:
    // 0..4294967295.
    QosGroupMin interface{}

    // This attribute is a key. Specifies the maximum value range from 0 to used
    // to identify a QoS group value. The type is interface{} with range:
    // 0..4294967295.
    QosGroupMax interface{}

    // Threshold bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    Bytes interface{}

    // Threshold size unit. The type is interface{} with range: 0..4294967295.
    ThreshSizeMetric interface{}

    // Threshold size basic units. The type is ThreshUnit.
    UnitVal interface{}

    // Threshold interval. The type is interface{} with range:
    // 0..18446744073709551615.
    ThresholdInterval interface{}

    // Threshold units metric. The type is interface{} with range: 0..4294967295.
    ThreshIntervalMetric interface{}

    // Threshold intveral basic units. The type is ThreshUnit.
    IntervalUnitVal interface{}
}

func (qlimitQosGrpThreshList *Interfaces_Interface_DiffservInfo_PriorityOperList_QlimitQosGrpThreshList) GetEntityData() *types.CommonEntityData {
    qlimitQosGrpThreshList.EntityData.YFilter = qlimitQosGrpThreshList.YFilter
    qlimitQosGrpThreshList.EntityData.YangName = "qlimit-qos-grp-thresh-list"
    qlimitQosGrpThreshList.EntityData.BundleName = "cisco_ios_xe"
    qlimitQosGrpThreshList.EntityData.ParentYangName = "priority-oper-list"
    qlimitQosGrpThreshList.EntityData.SegmentPath = "qlimit-qos-grp-thresh-list" + types.AddKeyToken(qlimitQosGrpThreshList.QosGroupMin, "qos-group-min") + types.AddKeyToken(qlimitQosGrpThreshList.QosGroupMax, "qos-group-max")
    qlimitQosGrpThreshList.EntityData.AbsolutePath = "Cisco-IOS-XE-interfaces-oper:interfaces/interface/diffserv-info/priority-oper-list/" + qlimitQosGrpThreshList.EntityData.SegmentPath
    qlimitQosGrpThreshList.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qlimitQosGrpThreshList.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qlimitQosGrpThreshList.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qlimitQosGrpThreshList.EntityData.Children = types.NewOrderedMap()
    qlimitQosGrpThreshList.EntityData.Leafs = types.NewOrderedMap()
    qlimitQosGrpThreshList.EntityData.Leafs.Append("qos-group-min", types.YLeaf{"QosGroupMin", qlimitQosGrpThreshList.QosGroupMin})
    qlimitQosGrpThreshList.EntityData.Leafs.Append("qos-group-max", types.YLeaf{"QosGroupMax", qlimitQosGrpThreshList.QosGroupMax})
    qlimitQosGrpThreshList.EntityData.Leafs.Append("bytes", types.YLeaf{"Bytes", qlimitQosGrpThreshList.Bytes})
    qlimitQosGrpThreshList.EntityData.Leafs.Append("thresh-size-metric", types.YLeaf{"ThreshSizeMetric", qlimitQosGrpThreshList.ThreshSizeMetric})
    qlimitQosGrpThreshList.EntityData.Leafs.Append("unit-val", types.YLeaf{"UnitVal", qlimitQosGrpThreshList.UnitVal})
    qlimitQosGrpThreshList.EntityData.Leafs.Append("threshold-interval", types.YLeaf{"ThresholdInterval", qlimitQosGrpThreshList.ThresholdInterval})
    qlimitQosGrpThreshList.EntityData.Leafs.Append("thresh-interval-metric", types.YLeaf{"ThreshIntervalMetric", qlimitQosGrpThreshList.ThreshIntervalMetric})
    qlimitQosGrpThreshList.EntityData.Leafs.Append("interval-unit-val", types.YLeaf{"IntervalUnitVal", qlimitQosGrpThreshList.IntervalUnitVal})

    qlimitQosGrpThreshList.EntityData.YListKeys = []string {"QosGroupMin", "QosGroupMax"}

    return &(qlimitQosGrpThreshList.EntityData)
}

// Interfaces_Interface_DiffservInfo_PriorityOperList_QlimitMplsExpThreshList
// mpls-exp-based queue limit data
type Interfaces_Interface_DiffservInfo_PriorityOperList_QlimitMplsExpThreshList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The minimum EXP field value to be used as match
    // criteria. Any number from 0 to 7. The type is interface{} with range:
    // 0..4294967295.
    ExpMin interface{}

    // This attribute is a key. The maximum EXP field value to be used as match
    // criteria. Any number from 0 to 7. The type is interface{} with range:
    // 0..4294967295.
    ExpMax interface{}

    // Threshold bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    Bytes interface{}

    // Threshold size unit. The type is interface{} with range: 0..4294967295.
    ThreshSizeMetric interface{}

    // Threshold size basic units. The type is ThreshUnit.
    UnitVal interface{}

    // Threshold interval. The type is interface{} with range:
    // 0..18446744073709551615.
    ThresholdInterval interface{}

    // Threshold units metric. The type is interface{} with range: 0..4294967295.
    ThreshIntervalMetric interface{}

    // Threshold intveral basic units. The type is ThreshUnit.
    IntervalUnitVal interface{}
}

func (qlimitMplsExpThreshList *Interfaces_Interface_DiffservInfo_PriorityOperList_QlimitMplsExpThreshList) GetEntityData() *types.CommonEntityData {
    qlimitMplsExpThreshList.EntityData.YFilter = qlimitMplsExpThreshList.YFilter
    qlimitMplsExpThreshList.EntityData.YangName = "qlimit-mpls-exp-thresh-list"
    qlimitMplsExpThreshList.EntityData.BundleName = "cisco_ios_xe"
    qlimitMplsExpThreshList.EntityData.ParentYangName = "priority-oper-list"
    qlimitMplsExpThreshList.EntityData.SegmentPath = "qlimit-mpls-exp-thresh-list" + types.AddKeyToken(qlimitMplsExpThreshList.ExpMin, "exp-min") + types.AddKeyToken(qlimitMplsExpThreshList.ExpMax, "exp-max")
    qlimitMplsExpThreshList.EntityData.AbsolutePath = "Cisco-IOS-XE-interfaces-oper:interfaces/interface/diffserv-info/priority-oper-list/" + qlimitMplsExpThreshList.EntityData.SegmentPath
    qlimitMplsExpThreshList.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qlimitMplsExpThreshList.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qlimitMplsExpThreshList.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qlimitMplsExpThreshList.EntityData.Children = types.NewOrderedMap()
    qlimitMplsExpThreshList.EntityData.Leafs = types.NewOrderedMap()
    qlimitMplsExpThreshList.EntityData.Leafs.Append("exp-min", types.YLeaf{"ExpMin", qlimitMplsExpThreshList.ExpMin})
    qlimitMplsExpThreshList.EntityData.Leafs.Append("exp-max", types.YLeaf{"ExpMax", qlimitMplsExpThreshList.ExpMax})
    qlimitMplsExpThreshList.EntityData.Leafs.Append("bytes", types.YLeaf{"Bytes", qlimitMplsExpThreshList.Bytes})
    qlimitMplsExpThreshList.EntityData.Leafs.Append("thresh-size-metric", types.YLeaf{"ThreshSizeMetric", qlimitMplsExpThreshList.ThreshSizeMetric})
    qlimitMplsExpThreshList.EntityData.Leafs.Append("unit-val", types.YLeaf{"UnitVal", qlimitMplsExpThreshList.UnitVal})
    qlimitMplsExpThreshList.EntityData.Leafs.Append("threshold-interval", types.YLeaf{"ThresholdInterval", qlimitMplsExpThreshList.ThresholdInterval})
    qlimitMplsExpThreshList.EntityData.Leafs.Append("thresh-interval-metric", types.YLeaf{"ThreshIntervalMetric", qlimitMplsExpThreshList.ThreshIntervalMetric})
    qlimitMplsExpThreshList.EntityData.Leafs.Append("interval-unit-val", types.YLeaf{"IntervalUnitVal", qlimitMplsExpThreshList.IntervalUnitVal})

    qlimitMplsExpThreshList.EntityData.YListKeys = []string {"ExpMin", "ExpMax"}

    return &(qlimitMplsExpThreshList.EntityData)
}

// Interfaces_Interface_DiffservInfo_PriorityOperList_QlimitDscpThreshList
// queue limit per dscp range
type Interfaces_Interface_DiffservInfo_PriorityOperList_QlimitDscpThreshList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Minimum of dscp range. The type is interface{}
    // with range: 0..4294967295.
    DscpMin interface{}

    // This attribute is a key. Maximum of dscp range. The type is interface{}
    // with range: 0..4294967295.
    DscpMax interface{}

    // Threshold bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    Bytes interface{}

    // Threshold size unit. The type is interface{} with range: 0..4294967295.
    ThreshSizeMetric interface{}

    // Threshold size basic units. The type is ThreshUnit.
    UnitVal interface{}

    // Threshold interval. The type is interface{} with range:
    // 0..18446744073709551615.
    ThresholdInterval interface{}

    // Threshold units metric. The type is interface{} with range: 0..4294967295.
    ThreshIntervalMetric interface{}

    // Threshold intveral basic units. The type is ThreshUnit.
    IntervalUnitVal interface{}
}

func (qlimitDscpThreshList *Interfaces_Interface_DiffservInfo_PriorityOperList_QlimitDscpThreshList) GetEntityData() *types.CommonEntityData {
    qlimitDscpThreshList.EntityData.YFilter = qlimitDscpThreshList.YFilter
    qlimitDscpThreshList.EntityData.YangName = "qlimit-dscp-thresh-list"
    qlimitDscpThreshList.EntityData.BundleName = "cisco_ios_xe"
    qlimitDscpThreshList.EntityData.ParentYangName = "priority-oper-list"
    qlimitDscpThreshList.EntityData.SegmentPath = "qlimit-dscp-thresh-list" + types.AddKeyToken(qlimitDscpThreshList.DscpMin, "dscp-min") + types.AddKeyToken(qlimitDscpThreshList.DscpMax, "dscp-max")
    qlimitDscpThreshList.EntityData.AbsolutePath = "Cisco-IOS-XE-interfaces-oper:interfaces/interface/diffserv-info/priority-oper-list/" + qlimitDscpThreshList.EntityData.SegmentPath
    qlimitDscpThreshList.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qlimitDscpThreshList.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qlimitDscpThreshList.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qlimitDscpThreshList.EntityData.Children = types.NewOrderedMap()
    qlimitDscpThreshList.EntityData.Leafs = types.NewOrderedMap()
    qlimitDscpThreshList.EntityData.Leafs.Append("dscp-min", types.YLeaf{"DscpMin", qlimitDscpThreshList.DscpMin})
    qlimitDscpThreshList.EntityData.Leafs.Append("dscp-max", types.YLeaf{"DscpMax", qlimitDscpThreshList.DscpMax})
    qlimitDscpThreshList.EntityData.Leafs.Append("bytes", types.YLeaf{"Bytes", qlimitDscpThreshList.Bytes})
    qlimitDscpThreshList.EntityData.Leafs.Append("thresh-size-metric", types.YLeaf{"ThreshSizeMetric", qlimitDscpThreshList.ThreshSizeMetric})
    qlimitDscpThreshList.EntityData.Leafs.Append("unit-val", types.YLeaf{"UnitVal", qlimitDscpThreshList.UnitVal})
    qlimitDscpThreshList.EntityData.Leafs.Append("threshold-interval", types.YLeaf{"ThresholdInterval", qlimitDscpThreshList.ThresholdInterval})
    qlimitDscpThreshList.EntityData.Leafs.Append("thresh-interval-metric", types.YLeaf{"ThreshIntervalMetric", qlimitDscpThreshList.ThreshIntervalMetric})
    qlimitDscpThreshList.EntityData.Leafs.Append("interval-unit-val", types.YLeaf{"IntervalUnitVal", qlimitDscpThreshList.IntervalUnitVal})

    qlimitDscpThreshList.EntityData.YListKeys = []string {"DscpMin", "DscpMax"}

    return &(qlimitDscpThreshList.EntityData)
}

// Interfaces_Interface_V4ProtocolStats
// IPv4 traffic statistics for this interface
type Interfaces_Interface_V4ProtocolStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The total number of packets received for the specified address family,
    // including those received in error. The type is interface{} with range:
    // 0..18446744073709551615.
    InPkts interface{}

    // The total number of octets received in input packets for the specified
    // address family, including those received in error. The type is interface{}
    // with range: 0..18446744073709551615.
    InOctets interface{}

    // Number of packets discarded due to errors for the specified address family,
    // including errors in the header, no route found to the destination, invalid
    // address, unknown protocol, etc. The type is interface{} with range:
    // 0..18446744073709551615.
    InErrorPkts interface{}

    // The number of input packets for which the device was not their final
    // destination and for which the device attempted to find a route to forward
    // them to that final destination. The type is interface{} with range:
    // 0..18446744073709551615.
    InForwardedPkts interface{}

    // The number of octets received in input packets for the specified address
    // family for which the device was not their final lodestination and for which
    // the device attempted to find a route to forward them to that final
    // destination. The type is interface{} with range: 0..18446744073709551615.
    InForwardedOctets interface{}

    // The number of input IP packets for the specified address family, for which
    // no problems were encountered to prevent their continued processing, but
    // were discarded (e.g., for lack of buffer space). The type is interface{}
    // with range: 0..18446744073709551615.
    InDiscardedPkts interface{}

    // The total number of IP packets for the specified address family that the
    // device supplied to the lower layers for transmission.  This includes
    // packets generated locally and those forwarded by the device. The type is
    // interface{} with range: 0..18446744073709551615.
    OutPkts interface{}

    // The total number of octets in IP packets for the specified address family
    // that the device supplied to the lower layers for transmission.  This
    // includes packets generated locally and those forwarded by the device. The
    // type is interface{} with range: 0..18446744073709551615.
    OutOctets interface{}

    // Number of IP packets for the specified address family locally generated and
    // discarded due to errors, including no route found to the IP destination.
    // The type is interface{} with range: 0..18446744073709551615.
    OutErrorPkts interface{}

    // The number of packets for which this entity was not their final IP
    // destination and for which it was successful in finding a path to their
    // final destination.text. The type is interface{} with range:
    // 0..18446744073709551615.
    OutForwardedPkts interface{}

    // The number of octets in packets for which this entity was not their final
    // IP destination and for which it was successful in finding a path to their
    // final destination. The type is interface{} with range:
    // 0..18446744073709551615.
    OutForwardedOctets interface{}

    // The number of output IP packets for the specified address family for which
    // no problem was encountered to prevent their transmission to their
    // destination, but were discarded (e.g., for lack of buffer space). The type
    // is interface{} with range: 0..18446744073709551615.
    OutDiscardedPkts interface{}
}

func (v4ProtocolStats *Interfaces_Interface_V4ProtocolStats) GetEntityData() *types.CommonEntityData {
    v4ProtocolStats.EntityData.YFilter = v4ProtocolStats.YFilter
    v4ProtocolStats.EntityData.YangName = "v4-protocol-stats"
    v4ProtocolStats.EntityData.BundleName = "cisco_ios_xe"
    v4ProtocolStats.EntityData.ParentYangName = "interface"
    v4ProtocolStats.EntityData.SegmentPath = "v4-protocol-stats"
    v4ProtocolStats.EntityData.AbsolutePath = "Cisco-IOS-XE-interfaces-oper:interfaces/interface/" + v4ProtocolStats.EntityData.SegmentPath
    v4ProtocolStats.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    v4ProtocolStats.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    v4ProtocolStats.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    v4ProtocolStats.EntityData.Children = types.NewOrderedMap()
    v4ProtocolStats.EntityData.Leafs = types.NewOrderedMap()
    v4ProtocolStats.EntityData.Leafs.Append("in-pkts", types.YLeaf{"InPkts", v4ProtocolStats.InPkts})
    v4ProtocolStats.EntityData.Leafs.Append("in-octets", types.YLeaf{"InOctets", v4ProtocolStats.InOctets})
    v4ProtocolStats.EntityData.Leafs.Append("in-error-pkts", types.YLeaf{"InErrorPkts", v4ProtocolStats.InErrorPkts})
    v4ProtocolStats.EntityData.Leafs.Append("in-forwarded-pkts", types.YLeaf{"InForwardedPkts", v4ProtocolStats.InForwardedPkts})
    v4ProtocolStats.EntityData.Leafs.Append("in-forwarded-octets", types.YLeaf{"InForwardedOctets", v4ProtocolStats.InForwardedOctets})
    v4ProtocolStats.EntityData.Leafs.Append("in-discarded-pkts", types.YLeaf{"InDiscardedPkts", v4ProtocolStats.InDiscardedPkts})
    v4ProtocolStats.EntityData.Leafs.Append("out-pkts", types.YLeaf{"OutPkts", v4ProtocolStats.OutPkts})
    v4ProtocolStats.EntityData.Leafs.Append("out-octets", types.YLeaf{"OutOctets", v4ProtocolStats.OutOctets})
    v4ProtocolStats.EntityData.Leafs.Append("out-error-pkts", types.YLeaf{"OutErrorPkts", v4ProtocolStats.OutErrorPkts})
    v4ProtocolStats.EntityData.Leafs.Append("out-forwarded-pkts", types.YLeaf{"OutForwardedPkts", v4ProtocolStats.OutForwardedPkts})
    v4ProtocolStats.EntityData.Leafs.Append("out-forwarded-octets", types.YLeaf{"OutForwardedOctets", v4ProtocolStats.OutForwardedOctets})
    v4ProtocolStats.EntityData.Leafs.Append("out-discarded-pkts", types.YLeaf{"OutDiscardedPkts", v4ProtocolStats.OutDiscardedPkts})

    v4ProtocolStats.EntityData.YListKeys = []string {}

    return &(v4ProtocolStats.EntityData)
}

// Interfaces_Interface_V6ProtocolStats
// IPv6 traffic statistics for this interface
type Interfaces_Interface_V6ProtocolStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The total number of packets received for the specified address family,
    // including those received in error. The type is interface{} with range:
    // 0..18446744073709551615.
    InPkts interface{}

    // The total number of octets received in input packets for the specified
    // address family, including those received in error. The type is interface{}
    // with range: 0..18446744073709551615.
    InOctets interface{}

    // Number of packets discarded due to errors for the specified address family,
    // including errors in the header, no route found to the destination, invalid
    // address, unknown protocol, etc. The type is interface{} with range:
    // 0..18446744073709551615.
    InErrorPkts interface{}

    // The number of input packets for which the device was not their final
    // destination and for which the device attempted to find a route to forward
    // them to that final destination. The type is interface{} with range:
    // 0..18446744073709551615.
    InForwardedPkts interface{}

    // The number of octets received in input packets for the specified address
    // family for which the device was not their final lodestination and for which
    // the device attempted to find a route to forward them to that final
    // destination. The type is interface{} with range: 0..18446744073709551615.
    InForwardedOctets interface{}

    // The number of input IP packets for the specified address family, for which
    // no problems were encountered to prevent their continued processing, but
    // were discarded (e.g., for lack of buffer space). The type is interface{}
    // with range: 0..18446744073709551615.
    InDiscardedPkts interface{}

    // The total number of IP packets for the specified address family that the
    // device supplied to the lower layers for transmission.  This includes
    // packets generated locally and those forwarded by the device. The type is
    // interface{} with range: 0..18446744073709551615.
    OutPkts interface{}

    // The total number of octets in IP packets for the specified address family
    // that the device supplied to the lower layers for transmission.  This
    // includes packets generated locally and those forwarded by the device. The
    // type is interface{} with range: 0..18446744073709551615.
    OutOctets interface{}

    // Number of IP packets for the specified address family locally generated and
    // discarded due to errors, including no route found to the IP destination.
    // The type is interface{} with range: 0..18446744073709551615.
    OutErrorPkts interface{}

    // The number of packets for which this entity was not their final IP
    // destination and for which it was successful in finding a path to their
    // final destination.text. The type is interface{} with range:
    // 0..18446744073709551615.
    OutForwardedPkts interface{}

    // The number of octets in packets for which this entity was not their final
    // IP destination and for which it was successful in finding a path to their
    // final destination. The type is interface{} with range:
    // 0..18446744073709551615.
    OutForwardedOctets interface{}

    // The number of output IP packets for the specified address family for which
    // no problem was encountered to prevent their transmission to their
    // destination, but were discarded (e.g., for lack of buffer space). The type
    // is interface{} with range: 0..18446744073709551615.
    OutDiscardedPkts interface{}
}

func (v6ProtocolStats *Interfaces_Interface_V6ProtocolStats) GetEntityData() *types.CommonEntityData {
    v6ProtocolStats.EntityData.YFilter = v6ProtocolStats.YFilter
    v6ProtocolStats.EntityData.YangName = "v6-protocol-stats"
    v6ProtocolStats.EntityData.BundleName = "cisco_ios_xe"
    v6ProtocolStats.EntityData.ParentYangName = "interface"
    v6ProtocolStats.EntityData.SegmentPath = "v6-protocol-stats"
    v6ProtocolStats.EntityData.AbsolutePath = "Cisco-IOS-XE-interfaces-oper:interfaces/interface/" + v6ProtocolStats.EntityData.SegmentPath
    v6ProtocolStats.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    v6ProtocolStats.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    v6ProtocolStats.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    v6ProtocolStats.EntityData.Children = types.NewOrderedMap()
    v6ProtocolStats.EntityData.Leafs = types.NewOrderedMap()
    v6ProtocolStats.EntityData.Leafs.Append("in-pkts", types.YLeaf{"InPkts", v6ProtocolStats.InPkts})
    v6ProtocolStats.EntityData.Leafs.Append("in-octets", types.YLeaf{"InOctets", v6ProtocolStats.InOctets})
    v6ProtocolStats.EntityData.Leafs.Append("in-error-pkts", types.YLeaf{"InErrorPkts", v6ProtocolStats.InErrorPkts})
    v6ProtocolStats.EntityData.Leafs.Append("in-forwarded-pkts", types.YLeaf{"InForwardedPkts", v6ProtocolStats.InForwardedPkts})
    v6ProtocolStats.EntityData.Leafs.Append("in-forwarded-octets", types.YLeaf{"InForwardedOctets", v6ProtocolStats.InForwardedOctets})
    v6ProtocolStats.EntityData.Leafs.Append("in-discarded-pkts", types.YLeaf{"InDiscardedPkts", v6ProtocolStats.InDiscardedPkts})
    v6ProtocolStats.EntityData.Leafs.Append("out-pkts", types.YLeaf{"OutPkts", v6ProtocolStats.OutPkts})
    v6ProtocolStats.EntityData.Leafs.Append("out-octets", types.YLeaf{"OutOctets", v6ProtocolStats.OutOctets})
    v6ProtocolStats.EntityData.Leafs.Append("out-error-pkts", types.YLeaf{"OutErrorPkts", v6ProtocolStats.OutErrorPkts})
    v6ProtocolStats.EntityData.Leafs.Append("out-forwarded-pkts", types.YLeaf{"OutForwardedPkts", v6ProtocolStats.OutForwardedPkts})
    v6ProtocolStats.EntityData.Leafs.Append("out-forwarded-octets", types.YLeaf{"OutForwardedOctets", v6ProtocolStats.OutForwardedOctets})
    v6ProtocolStats.EntityData.Leafs.Append("out-discarded-pkts", types.YLeaf{"OutDiscardedPkts", v6ProtocolStats.OutDiscardedPkts})

    v6ProtocolStats.EntityData.YListKeys = []string {}

    return &(v6ProtocolStats.EntityData)
}

// Interfaces_Interface_LagAggregateState
// Operational state variables for logical
// aggregate / LAG interfaces
type Interfaces_Interface_LagAggregateState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Specify the logical aggregate interface to which
    // this id belongs. The type is string.
    AggregateId interface{}

    // Type to define the lag-type, i.e., how the LAG is defined and managed. The
    // type is AggregationType.
    LagType interface{}

    // Specifies the minimum number of member interfaces that must be active for
    // the aggregate interface to be available. The type is interface{} with
    // range: 0..65535.
    MinLinks interface{}

    // Reports effective speed of the aggregate interface, based on speed of
    // active member interfaces. The type is interface{} with range:
    // 0..4294967295.
    LagSpeed interface{}

    // List of current member interfaces for the aggregate, expressed as
    // references to existing interfaces. The type is slice of string.
    Members []interface{}
}

func (lagAggregateState *Interfaces_Interface_LagAggregateState) GetEntityData() *types.CommonEntityData {
    lagAggregateState.EntityData.YFilter = lagAggregateState.YFilter
    lagAggregateState.EntityData.YangName = "lag-aggregate-state"
    lagAggregateState.EntityData.BundleName = "cisco_ios_xe"
    lagAggregateState.EntityData.ParentYangName = "interface"
    lagAggregateState.EntityData.SegmentPath = "lag-aggregate-state" + types.AddKeyToken(lagAggregateState.AggregateId, "aggregate-id")
    lagAggregateState.EntityData.AbsolutePath = "Cisco-IOS-XE-interfaces-oper:interfaces/interface/" + lagAggregateState.EntityData.SegmentPath
    lagAggregateState.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lagAggregateState.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lagAggregateState.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lagAggregateState.EntityData.Children = types.NewOrderedMap()
    lagAggregateState.EntityData.Leafs = types.NewOrderedMap()
    lagAggregateState.EntityData.Leafs.Append("aggregate-id", types.YLeaf{"AggregateId", lagAggregateState.AggregateId})
    lagAggregateState.EntityData.Leafs.Append("lag-type", types.YLeaf{"LagType", lagAggregateState.LagType})
    lagAggregateState.EntityData.Leafs.Append("min-links", types.YLeaf{"MinLinks", lagAggregateState.MinLinks})
    lagAggregateState.EntityData.Leafs.Append("lag-speed", types.YLeaf{"LagSpeed", lagAggregateState.LagSpeed})
    lagAggregateState.EntityData.Leafs.Append("members", types.YLeaf{"Members", lagAggregateState.Members})

    lagAggregateState.EntityData.YListKeys = []string {"AggregateId"}

    return &(lagAggregateState.EntityData)
}

// Interfaces_Interface_EtherState
// The Ethernet state information
type Interfaces_Interface_EtherState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // When auto-negotiate is set to TRUE, and the  interface has completed
    // auto-negotiation with the  remote peer, this value shows the duplex mode
    // that  has been negotiated. The type is EtherDuplex.
    NegotiatedDuplexMode interface{}

    // When auto-negotiate is set to TRUE, and  the interface has completed
    // auto-negotiation with  the remote peer, this value shows the interface 
    // speed that has been negotiated. The type is EtherSpeed.
    NegotiatedPortSpeed interface{}

    // Set to TRUE to request the interface to  auto-negotiate transmission
    // parameters with its  peer interface. When set to FALSE, the transmission
    // parameters are specified manually. The type is bool.
    AutoNegotiate interface{}

    // Enable or disable flow control for this  interface. Ethernet flow control
    // is a mechanism by  which a receiver may send PAUSE frames to a sender to
    // stop transmission for a specified time.  This setting should override
    // auto-negotiated flow  control settings. If left unspecified, and 
    // auto-negotiate is TRUE, flow control mode is  negotiated with the peer
    // interface. The type is bool.
    EnableFlowControl interface{}
}

func (etherState *Interfaces_Interface_EtherState) GetEntityData() *types.CommonEntityData {
    etherState.EntityData.YFilter = etherState.YFilter
    etherState.EntityData.YangName = "ether-state"
    etherState.EntityData.BundleName = "cisco_ios_xe"
    etherState.EntityData.ParentYangName = "interface"
    etherState.EntityData.SegmentPath = "ether-state"
    etherState.EntityData.AbsolutePath = "Cisco-IOS-XE-interfaces-oper:interfaces/interface/" + etherState.EntityData.SegmentPath
    etherState.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    etherState.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    etherState.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    etherState.EntityData.Children = types.NewOrderedMap()
    etherState.EntityData.Leafs = types.NewOrderedMap()
    etherState.EntityData.Leafs.Append("negotiated-duplex-mode", types.YLeaf{"NegotiatedDuplexMode", etherState.NegotiatedDuplexMode})
    etherState.EntityData.Leafs.Append("negotiated-port-speed", types.YLeaf{"NegotiatedPortSpeed", etherState.NegotiatedPortSpeed})
    etherState.EntityData.Leafs.Append("auto-negotiate", types.YLeaf{"AutoNegotiate", etherState.AutoNegotiate})
    etherState.EntityData.Leafs.Append("enable-flow-control", types.YLeaf{"EnableFlowControl", etherState.EnableFlowControl})

    etherState.EntityData.YListKeys = []string {}

    return &(etherState.EntityData)
}

// Interfaces_Interface_EtherStats
// The Ethernet statistics
type Interfaces_Interface_EtherStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MAC layer control frames received on the interface. The type is interface{}
    // with range: 0..18446744073709551615.
    InMacControlFrames interface{}

    // MAC layer PAUSE frames received on the interface. The type is interface{}
    // with range: 0..18446744073709551615.
    InMacPauseFrames interface{}

    // Number of oversize frames received on the interface. The type is
    // interface{} with range: 0..18446744073709551615.
    InOversizeFrames interface{}

    // Number of jabber frames received on the interface.  Jabber frames are
    // typically defined as oversize frames which also have a bad CRC. 
    // Implementations may use slightly different definitions of what constitutes
    // a jabber frame.  Often indicative of a NIC hardware problem. The type is
    // interface{} with range: 0..18446744073709551615.
    InJabberFrames interface{}

    // Number of fragment frames received on the interface. The type is
    // interface{} with range: 0..18446744073709551615.
    InFragmentFrames interface{}

    // Number of 802.1q tagged frames received on the interface. The type is
    // interface{} with range: 0..18446744073709551615.
    In8021qFrames interface{}

    // MAC layer control frames sent on the interface. The type is interface{}
    // with range: 0..18446744073709551615.
    OutMacControlFrames interface{}

    // MAC layer PAUSE frames sent on the interface. The type is interface{} with
    // range: 0..18446744073709551615.
    OutMacPauseFrames interface{}

    // Number of 802.1q tagged frames sent on the interface. The type is
    // interface{} with range: 0..18446744073709551615.
    Out8021qFrames interface{}
}

func (etherStats *Interfaces_Interface_EtherStats) GetEntityData() *types.CommonEntityData {
    etherStats.EntityData.YFilter = etherStats.YFilter
    etherStats.EntityData.YangName = "ether-stats"
    etherStats.EntityData.BundleName = "cisco_ios_xe"
    etherStats.EntityData.ParentYangName = "interface"
    etherStats.EntityData.SegmentPath = "ether-stats"
    etherStats.EntityData.AbsolutePath = "Cisco-IOS-XE-interfaces-oper:interfaces/interface/" + etherStats.EntityData.SegmentPath
    etherStats.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    etherStats.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    etherStats.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    etherStats.EntityData.Children = types.NewOrderedMap()
    etherStats.EntityData.Leafs = types.NewOrderedMap()
    etherStats.EntityData.Leafs.Append("in-mac-control-frames", types.YLeaf{"InMacControlFrames", etherStats.InMacControlFrames})
    etherStats.EntityData.Leafs.Append("in-mac-pause-frames", types.YLeaf{"InMacPauseFrames", etherStats.InMacPauseFrames})
    etherStats.EntityData.Leafs.Append("in-oversize-frames", types.YLeaf{"InOversizeFrames", etherStats.InOversizeFrames})
    etherStats.EntityData.Leafs.Append("in-jabber-frames", types.YLeaf{"InJabberFrames", etherStats.InJabberFrames})
    etherStats.EntityData.Leafs.Append("in-fragment-frames", types.YLeaf{"InFragmentFrames", etherStats.InFragmentFrames})
    etherStats.EntityData.Leafs.Append("in-8021q-frames", types.YLeaf{"In8021qFrames", etherStats.In8021qFrames})
    etherStats.EntityData.Leafs.Append("out-mac-control-frames", types.YLeaf{"OutMacControlFrames", etherStats.OutMacControlFrames})
    etherStats.EntityData.Leafs.Append("out-mac-pause-frames", types.YLeaf{"OutMacPauseFrames", etherStats.OutMacPauseFrames})
    etherStats.EntityData.Leafs.Append("out-8021q-frames", types.YLeaf{"Out8021qFrames", etherStats.Out8021qFrames})

    etherStats.EntityData.YListKeys = []string {}

    return &(etherStats.EntityData)
}

// Interfaces_Interface_SerialState
// The T1E1 serial state information
type Interfaces_Interface_SerialState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Cyclic Redundancy Code type configured on the interface. The type is
    // SerialCrc.
    CrcType interface{}

    // Loopback mode the interface is operating in. The type is T1e1LoopbackMode.
    Loopback interface{}

    // Keeplive interval in seconds. The type is interface{} with range:
    // 0..4294967295.
    Keeplive interface{}

    // Time slots bitmap occupied by this serial interface. The type is
    // interface{} with range: 0..4294967295.
    Timeslot interface{}

    // Subrate operating per slot. The type is SubrateSpeed.
    Subrate interface{}
}

func (serialState *Interfaces_Interface_SerialState) GetEntityData() *types.CommonEntityData {
    serialState.EntityData.YFilter = serialState.YFilter
    serialState.EntityData.YangName = "serial-state"
    serialState.EntityData.BundleName = "cisco_ios_xe"
    serialState.EntityData.ParentYangName = "interface"
    serialState.EntityData.SegmentPath = "serial-state"
    serialState.EntityData.AbsolutePath = "Cisco-IOS-XE-interfaces-oper:interfaces/interface/" + serialState.EntityData.SegmentPath
    serialState.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    serialState.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    serialState.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    serialState.EntityData.Children = types.NewOrderedMap()
    serialState.EntityData.Leafs = types.NewOrderedMap()
    serialState.EntityData.Leafs.Append("crc-type", types.YLeaf{"CrcType", serialState.CrcType})
    serialState.EntityData.Leafs.Append("loopback", types.YLeaf{"Loopback", serialState.Loopback})
    serialState.EntityData.Leafs.Append("keeplive", types.YLeaf{"Keeplive", serialState.Keeplive})
    serialState.EntityData.Leafs.Append("timeslot", types.YLeaf{"Timeslot", serialState.Timeslot})
    serialState.EntityData.Leafs.Append("subrate", types.YLeaf{"Subrate", serialState.Subrate})

    serialState.EntityData.YListKeys = []string {}

    return &(serialState.EntityData)
}

// Interfaces_Interface_SerialStats
// The T1E1 statistics
type Interfaces_Interface_SerialStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of receive abort packets due to clock slides. The type is
    // interface{} with range: 0..4294967295.
    InAbortClockError interface{}
}

func (serialStats *Interfaces_Interface_SerialStats) GetEntityData() *types.CommonEntityData {
    serialStats.EntityData.YFilter = serialStats.YFilter
    serialStats.EntityData.YangName = "serial-stats"
    serialStats.EntityData.BundleName = "cisco_ios_xe"
    serialStats.EntityData.ParentYangName = "interface"
    serialStats.EntityData.SegmentPath = "serial-stats"
    serialStats.EntityData.AbsolutePath = "Cisco-IOS-XE-interfaces-oper:interfaces/interface/" + serialStats.EntityData.SegmentPath
    serialStats.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    serialStats.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    serialStats.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    serialStats.EntityData.Children = types.NewOrderedMap()
    serialStats.EntityData.Leafs = types.NewOrderedMap()
    serialStats.EntityData.Leafs.Append("in-abort-clock-error", types.YLeaf{"InAbortClockError", serialStats.InAbortClockError})

    serialStats.EntityData.YListKeys = []string {}

    return &(serialStats.EntityData)
}

