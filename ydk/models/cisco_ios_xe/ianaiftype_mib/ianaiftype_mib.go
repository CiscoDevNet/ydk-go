// This MIB module defines the IANAifType Textual
// Convention, and thus the enumerated values of
// the ifType object defined in MIB-II's ifTable.
package ianaiftype_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ianaiftype_mib"))
}

// IANAifType represents transmission subtree OIDs.
type IANAifType string

const (
    IANAifType_other IANAifType = "other"

    IANAifType_regular1822 IANAifType = "regular1822"

    IANAifType_hdh1822 IANAifType = "hdh1822"

    IANAifType_ddnX25 IANAifType = "ddnX25"

    IANAifType_rfc877x25 IANAifType = "rfc877x25"

    IANAifType_ethernetCsmacd IANAifType = "ethernetCsmacd"

    IANAifType_iso88023Csmacd IANAifType = "iso88023Csmacd"

    IANAifType_iso88024TokenBus IANAifType = "iso88024TokenBus"

    IANAifType_iso88025TokenRing IANAifType = "iso88025TokenRing"

    IANAifType_iso88026Man IANAifType = "iso88026Man"

    IANAifType_starLan IANAifType = "starLan"

    IANAifType_proteon10Mbit IANAifType = "proteon10Mbit"

    IANAifType_proteon80Mbit IANAifType = "proteon80Mbit"

    IANAifType_hyperchannel IANAifType = "hyperchannel"

    IANAifType_fddi IANAifType = "fddi"

    IANAifType_lapb IANAifType = "lapb"

    IANAifType_sdlc IANAifType = "sdlc"

    IANAifType_ds1 IANAifType = "ds1"

    IANAifType_e1 IANAifType = "e1"

    IANAifType_basicISDN IANAifType = "basicISDN"

    IANAifType_primaryISDN IANAifType = "primaryISDN"

    IANAifType_propPointToPointSerial IANAifType = "propPointToPointSerial"

    IANAifType_ppp IANAifType = "ppp"

    IANAifType_softwareLoopback IANAifType = "softwareLoopback"

    IANAifType_eon IANAifType = "eon"

    IANAifType_ethernet3Mbit IANAifType = "ethernet3Mbit"

    IANAifType_nsip IANAifType = "nsip"

    IANAifType_slip IANAifType = "slip"

    IANAifType_ultra IANAifType = "ultra"

    IANAifType_ds3 IANAifType = "ds3"

    IANAifType_sip IANAifType = "sip"

    IANAifType_frameRelay IANAifType = "frameRelay"

    IANAifType_rs232 IANAifType = "rs232"

    IANAifType_para IANAifType = "para"

    IANAifType_arcnet IANAifType = "arcnet"

    IANAifType_arcnetPlus IANAifType = "arcnetPlus"

    IANAifType_atm IANAifType = "atm"

    IANAifType_miox25 IANAifType = "miox25"

    IANAifType_sonet IANAifType = "sonet"

    IANAifType_x25ple IANAifType = "x25ple"

    IANAifType_iso88022llc IANAifType = "iso88022llc"

    IANAifType_localTalk IANAifType = "localTalk"

    IANAifType_smdsDxi IANAifType = "smdsDxi"

    IANAifType_frameRelayService IANAifType = "frameRelayService"

    IANAifType_v35 IANAifType = "v35"

    IANAifType_hssi IANAifType = "hssi"

    IANAifType_hippi IANAifType = "hippi"

    IANAifType_modem IANAifType = "modem"

    IANAifType_aal5 IANAifType = "aal5"

    IANAifType_sonetPath IANAifType = "sonetPath"

    IANAifType_sonetVT IANAifType = "sonetVT"

    IANAifType_smdsIcip IANAifType = "smdsIcip"

    IANAifType_propVirtual IANAifType = "propVirtual"

    IANAifType_propMultiplexor IANAifType = "propMultiplexor"

    IANAifType_ieee80212 IANAifType = "ieee80212"

    IANAifType_fibreChannel IANAifType = "fibreChannel"

    IANAifType_hippiInterface IANAifType = "hippiInterface"

    IANAifType_frameRelayInterconnect IANAifType = "frameRelayInterconnect"

    IANAifType_aflane8023 IANAifType = "aflane8023"

    IANAifType_aflane8025 IANAifType = "aflane8025"

    IANAifType_cctEmul IANAifType = "cctEmul"

    IANAifType_fastEther IANAifType = "fastEther"

    IANAifType_isdn IANAifType = "isdn"

    IANAifType_v11 IANAifType = "v11"

    IANAifType_v36 IANAifType = "v36"

    IANAifType_g703at64k IANAifType = "g703at64k"

    IANAifType_g703at2mb IANAifType = "g703at2mb"

    IANAifType_qllc IANAifType = "qllc"

    IANAifType_fastEtherFX IANAifType = "fastEtherFX"

    IANAifType_channel IANAifType = "channel"

    IANAifType_ieee80211 IANAifType = "ieee80211"

    IANAifType_ibm370parChan IANAifType = "ibm370parChan"

    IANAifType_escon IANAifType = "escon"

    IANAifType_dlsw IANAifType = "dlsw"

    IANAifType_isdns IANAifType = "isdns"

    IANAifType_isdnu IANAifType = "isdnu"

    IANAifType_lapd IANAifType = "lapd"

    IANAifType_ipSwitch IANAifType = "ipSwitch"

    IANAifType_rsrb IANAifType = "rsrb"

    IANAifType_atmLogical IANAifType = "atmLogical"

    IANAifType_ds0 IANAifType = "ds0"

    IANAifType_ds0Bundle IANAifType = "ds0Bundle"

    IANAifType_bsc IANAifType = "bsc"

    IANAifType_async IANAifType = "async"

    IANAifType_cnr IANAifType = "cnr"

    IANAifType_iso88025Dtr IANAifType = "iso88025Dtr"

    IANAifType_eplrs IANAifType = "eplrs"

    IANAifType_arap IANAifType = "arap"

    IANAifType_propCnls IANAifType = "propCnls"

    IANAifType_hostPad IANAifType = "hostPad"

    IANAifType_termPad IANAifType = "termPad"

    IANAifType_frameRelayMPI IANAifType = "frameRelayMPI"

    IANAifType_x213 IANAifType = "x213"

    IANAifType_adsl IANAifType = "adsl"

    IANAifType_radsl IANAifType = "radsl"

    IANAifType_sdsl IANAifType = "sdsl"

    IANAifType_vdsl IANAifType = "vdsl"

    IANAifType_iso88025CRFPInt IANAifType = "iso88025CRFPInt"

    IANAifType_myrinet IANAifType = "myrinet"

    IANAifType_voiceEM IANAifType = "voiceEM"

    IANAifType_voiceFXO IANAifType = "voiceFXO"

    IANAifType_voiceFXS IANAifType = "voiceFXS"

    IANAifType_voiceEncap IANAifType = "voiceEncap"

    IANAifType_voiceOverIp IANAifType = "voiceOverIp"

    IANAifType_atmDxi IANAifType = "atmDxi"

    IANAifType_atmFuni IANAifType = "atmFuni"

    IANAifType_atmIma IANAifType = "atmIma"

    IANAifType_pppMultilinkBundle IANAifType = "pppMultilinkBundle"

    IANAifType_ipOverCdlc IANAifType = "ipOverCdlc"

    IANAifType_ipOverClaw IANAifType = "ipOverClaw"

    IANAifType_stackToStack IANAifType = "stackToStack"

    IANAifType_virtualIpAddress IANAifType = "virtualIpAddress"

    IANAifType_mpc IANAifType = "mpc"

    IANAifType_ipOverAtm IANAifType = "ipOverAtm"

    IANAifType_iso88025Fiber IANAifType = "iso88025Fiber"

    IANAifType_tdlc IANAifType = "tdlc"

    IANAifType_gigabitEthernet IANAifType = "gigabitEthernet"

    IANAifType_hdlc IANAifType = "hdlc"

    IANAifType_lapf IANAifType = "lapf"

    IANAifType_v37 IANAifType = "v37"

    IANAifType_x25mlp IANAifType = "x25mlp"

    IANAifType_x25huntGroup IANAifType = "x25huntGroup"

    IANAifType_trasnpHdlc IANAifType = "trasnpHdlc"

    IANAifType_interleave IANAifType = "interleave"

    IANAifType_fast IANAifType = "fast"

    IANAifType_ip IANAifType = "ip"

    IANAifType_docsCableMaclayer IANAifType = "docsCableMaclayer"

    IANAifType_docsCableDownstream IANAifType = "docsCableDownstream"

    IANAifType_docsCableUpstream IANAifType = "docsCableUpstream"

    IANAifType_a12MppSwitch IANAifType = "a12MppSwitch"

    IANAifType_tunnel IANAifType = "tunnel"

    IANAifType_coffee IANAifType = "coffee"

    IANAifType_ces IANAifType = "ces"

    IANAifType_atmSubInterface IANAifType = "atmSubInterface"

    IANAifType_l2vlan IANAifType = "l2vlan"

    IANAifType_l3ipvlan IANAifType = "l3ipvlan"

    IANAifType_l3ipxvlan IANAifType = "l3ipxvlan"

    IANAifType_digitalPowerline IANAifType = "digitalPowerline"

    IANAifType_mediaMailOverIp IANAifType = "mediaMailOverIp"

    IANAifType_dtm IANAifType = "dtm"

    IANAifType_dcn IANAifType = "dcn"

    IANAifType_ipForward IANAifType = "ipForward"

    IANAifType_msdsl IANAifType = "msdsl"

    IANAifType_ieee1394 IANAifType = "ieee1394"

    IANAifType_if_gsn IANAifType = "if-gsn"

    IANAifType_dvbRccMacLayer IANAifType = "dvbRccMacLayer"

    IANAifType_dvbRccDownstream IANAifType = "dvbRccDownstream"

    IANAifType_dvbRccUpstream IANAifType = "dvbRccUpstream"

    IANAifType_atmVirtual IANAifType = "atmVirtual"

    IANAifType_mplsTunnel IANAifType = "mplsTunnel"

    IANAifType_srp IANAifType = "srp"

    IANAifType_voiceOverAtm IANAifType = "voiceOverAtm"

    IANAifType_voiceOverFrameRelay IANAifType = "voiceOverFrameRelay"

    IANAifType_idsl IANAifType = "idsl"

    IANAifType_compositeLink IANAifType = "compositeLink"

    IANAifType_ss7SigLink IANAifType = "ss7SigLink"

    IANAifType_propWirelessP2P IANAifType = "propWirelessP2P"

    IANAifType_frForward IANAifType = "frForward"

    IANAifType_rfc1483 IANAifType = "rfc1483"

    IANAifType_usb IANAifType = "usb"

    IANAifType_ieee8023adLag IANAifType = "ieee8023adLag"

    IANAifType_bgppolicyaccounting IANAifType = "bgppolicyaccounting"

    IANAifType_frf16MfrBundle IANAifType = "frf16MfrBundle"

    IANAifType_h323Gatekeeper IANAifType = "h323Gatekeeper"

    IANAifType_h323Proxy IANAifType = "h323Proxy"

    IANAifType_mpls IANAifType = "mpls"

    IANAifType_mfSigLink IANAifType = "mfSigLink"

    IANAifType_hdsl2 IANAifType = "hdsl2"

    IANAifType_shdsl IANAifType = "shdsl"

    IANAifType_ds1FDL IANAifType = "ds1FDL"

    IANAifType_pos IANAifType = "pos"

    IANAifType_dvbAsiIn IANAifType = "dvbAsiIn"

    IANAifType_dvbAsiOut IANAifType = "dvbAsiOut"

    IANAifType_plc IANAifType = "plc"

    IANAifType_nfas IANAifType = "nfas"

    IANAifType_tr008 IANAifType = "tr008"

    IANAifType_gr303RDT IANAifType = "gr303RDT"

    IANAifType_gr303IDT IANAifType = "gr303IDT"

    IANAifType_isup IANAifType = "isup"

    IANAifType_propDocsWirelessMaclayer IANAifType = "propDocsWirelessMaclayer"

    IANAifType_propDocsWirelessDownstream IANAifType = "propDocsWirelessDownstream"

    IANAifType_propDocsWirelessUpstream IANAifType = "propDocsWirelessUpstream"

    IANAifType_hiperlan2 IANAifType = "hiperlan2"

    IANAifType_propBWAp2Mp IANAifType = "propBWAp2Mp"

    IANAifType_sonetOverheadChannel IANAifType = "sonetOverheadChannel"

    IANAifType_digitalWrapperOverheadChannel IANAifType = "digitalWrapperOverheadChannel"

    IANAifType_aal2 IANAifType = "aal2"

    IANAifType_radioMAC IANAifType = "radioMAC"

    IANAifType_atmRadio IANAifType = "atmRadio"

    IANAifType_imt IANAifType = "imt"

    IANAifType_mvl IANAifType = "mvl"

    IANAifType_reachDSL IANAifType = "reachDSL"

    IANAifType_frDlciEndPt IANAifType = "frDlciEndPt"

    IANAifType_atmVciEndPt IANAifType = "atmVciEndPt"

    IANAifType_opticalChannel IANAifType = "opticalChannel"

    IANAifType_opticalTransport IANAifType = "opticalTransport"

    IANAifType_propAtm IANAifType = "propAtm"

    IANAifType_voiceOverCable IANAifType = "voiceOverCable"

    IANAifType_infiniband IANAifType = "infiniband"

    IANAifType_teLink IANAifType = "teLink"

    IANAifType_q2931 IANAifType = "q2931"

    IANAifType_virtualTg IANAifType = "virtualTg"

    IANAifType_sipTg IANAifType = "sipTg"

    IANAifType_sipSig IANAifType = "sipSig"

    IANAifType_docsCableUpstreamChannel IANAifType = "docsCableUpstreamChannel"

    IANAifType_econet IANAifType = "econet"

    IANAifType_pon155 IANAifType = "pon155"

    IANAifType_pon622 IANAifType = "pon622"

    IANAifType_bridge IANAifType = "bridge"

    IANAifType_linegroup IANAifType = "linegroup"

    IANAifType_voiceEMFGD IANAifType = "voiceEMFGD"

    IANAifType_voiceFGDEANA IANAifType = "voiceFGDEANA"

    IANAifType_voiceDID IANAifType = "voiceDID"

    IANAifType_mpegTransport IANAifType = "mpegTransport"

    IANAifType_sixToFour IANAifType = "sixToFour"

    IANAifType_gtp IANAifType = "gtp"

    IANAifType_pdnEtherLoop1 IANAifType = "pdnEtherLoop1"

    IANAifType_pdnEtherLoop2 IANAifType = "pdnEtherLoop2"

    IANAifType_opticalChannelGroup IANAifType = "opticalChannelGroup"

    IANAifType_homepna IANAifType = "homepna"

    IANAifType_gfp IANAifType = "gfp"

    IANAifType_ciscoISLvlan IANAifType = "ciscoISLvlan"

    IANAifType_actelisMetaLOOP IANAifType = "actelisMetaLOOP"

    IANAifType_fcipLink IANAifType = "fcipLink"

    IANAifType_rpr IANAifType = "rpr"

    IANAifType_qam IANAifType = "qam"

    IANAifType_lmp IANAifType = "lmp"

    IANAifType_cblVectaStar IANAifType = "cblVectaStar"

    IANAifType_docsCableMCmtsDownstream IANAifType = "docsCableMCmtsDownstream"

    IANAifType_adsl2 IANAifType = "adsl2"

    IANAifType_macSecControlledIF IANAifType = "macSecControlledIF"

    IANAifType_macSecUncontrolledIF IANAifType = "macSecUncontrolledIF"

    IANAifType_aviciOpticalEther IANAifType = "aviciOpticalEther"

    IANAifType_atmbond IANAifType = "atmbond"

    IANAifType_voiceFGDOS IANAifType = "voiceFGDOS"

    IANAifType_mocaVersion1 IANAifType = "mocaVersion1"

    IANAifType_ieee80216WMAN IANAifType = "ieee80216WMAN"

    IANAifType_adsl2plus IANAifType = "adsl2plus"

    IANAifType_dvbRcsMacLayer IANAifType = "dvbRcsMacLayer"

    IANAifType_dvbTdm IANAifType = "dvbTdm"

    IANAifType_dvbRcsTdma IANAifType = "dvbRcsTdma"

    IANAifType_x86Laps IANAifType = "x86Laps"

    IANAifType_wwanPP IANAifType = "wwanPP"

    IANAifType_wwanPP2 IANAifType = "wwanPP2"

    IANAifType_voiceEBS IANAifType = "voiceEBS"

    IANAifType_ifPwType IANAifType = "ifPwType"

    IANAifType_ilan IANAifType = "ilan"

    IANAifType_pip IANAifType = "pip"

    IANAifType_aluELP IANAifType = "aluELP"

    IANAifType_gpon IANAifType = "gpon"

    IANAifType_vdsl2 IANAifType = "vdsl2"

    IANAifType_capwapDot11Profile IANAifType = "capwapDot11Profile"

    IANAifType_capwapDot11Bss IANAifType = "capwapDot11Bss"

    IANAifType_capwapWtpVirtualRadio IANAifType = "capwapWtpVirtualRadio"

    IANAifType_bits IANAifType = "bits"

    IANAifType_docsCableUpstreamRfPort IANAifType = "docsCableUpstreamRfPort"

    IANAifType_cableDownstreamRfPort IANAifType = "cableDownstreamRfPort"

    IANAifType_vmwareVirtualNic IANAifType = "vmwareVirtualNic"

    IANAifType_ieee802154 IANAifType = "ieee802154"

    IANAifType_otnOdu IANAifType = "otnOdu"

    IANAifType_otnOtu IANAifType = "otnOtu"

    IANAifType_ifVfiType IANAifType = "ifVfiType"

    IANAifType_g9981 IANAifType = "g9981"

    IANAifType_g9982 IANAifType = "g9982"

    IANAifType_g9983 IANAifType = "g9983"

    IANAifType_aluEpon IANAifType = "aluEpon"

    IANAifType_aluEponOnu IANAifType = "aluEponOnu"

    IANAifType_aluEponPhysicalUni IANAifType = "aluEponPhysicalUni"

    IANAifType_aluEponLogicalLink IANAifType = "aluEponLogicalLink"

    IANAifType_aluGponOnu IANAifType = "aluGponOnu"

    IANAifType_aluGponPhysicalUni IANAifType = "aluGponPhysicalUni"

    IANAifType_vmwareNicTeam IANAifType = "vmwareNicTeam"

    IANAifType_docsOfdmDownstream IANAifType = "docsOfdmDownstream"

    IANAifType_docsOfdmaUpstream IANAifType = "docsOfdmaUpstream"

    IANAifType_gfast IANAifType = "gfast"

    IANAifType_sdci IANAifType = "sdci"

    IANAifType_xboxWireless IANAifType = "xboxWireless"

    IANAifType_fastdsl IANAifType = "fastdsl"

    IANAifType_docsCableScte55d1FwdOob IANAifType = "docsCableScte55d1FwdOob"

    IANAifType_docsCableScte55d1RetOob IANAifType = "docsCableScte55d1RetOob"

    IANAifType_docsCableScte55d2DsOob IANAifType = "docsCableScte55d2DsOob"

    IANAifType_docsCableScte55d2UsOob IANAifType = "docsCableScte55d2UsOob"

    IANAifType_docsCableNdf IANAifType = "docsCableNdf"

    IANAifType_docsCableNdr IANAifType = "docsCableNdr"
)

// IANAtunnelType represents values.
type IANAtunnelType string

const (
    IANAtunnelType_other IANAtunnelType = "other"

    IANAtunnelType_direct IANAtunnelType = "direct"

    IANAtunnelType_gre IANAtunnelType = "gre"

    IANAtunnelType_minimal IANAtunnelType = "minimal"

    IANAtunnelType_l2tp IANAtunnelType = "l2tp"

    IANAtunnelType_pptp IANAtunnelType = "pptp"

    IANAtunnelType_l2f IANAtunnelType = "l2f"

    IANAtunnelType_udp IANAtunnelType = "udp"

    IANAtunnelType_atmp IANAtunnelType = "atmp"

    IANAtunnelType_msdp IANAtunnelType = "msdp"

    IANAtunnelType_sixToFour IANAtunnelType = "sixToFour"

    IANAtunnelType_sixOverFour IANAtunnelType = "sixOverFour"

    IANAtunnelType_isatap IANAtunnelType = "isatap"

    IANAtunnelType_teredo IANAtunnelType = "teredo"
)

