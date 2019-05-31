// This YANG module defines YANG identities for IANA-registered
// interface types.
// 
// This YANG module is maintained by IANA and reflects the
// 'ifType definitions' registry.
// 
// The latest revision of this YANG module can be obtained from
// the IANA web site.
// 
// Requests for new values should be made to IANA via
// email (iana@iana.org).
// 
// Copyright (c) 2014 IETF Trust and the persons identified as
// authors of the code.  All rights reserved.
// 
// Redistribution and use in source and binary forms, with or
// without modification, is permitted pursuant to, and subject
// to the license terms contained in, the Simplified BSD License
// set forth in Section 4.c of the IETF Trust's Legal Provisions
// Relating to IETF Documents
// (http://trustee.ietf.org/license-info).
// 
// The initial version of this YANG module is part of RFC 7224;
// see the RFC itself for full legal notices.
package iana_if_type

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package iana_if_type"))
}

type IanaInterfaceType struct {
}

func (id IanaInterfaceType) String() string {
	return "iana-if-type:iana-interface-type"
}

type Other struct {
}

func (id Other) String() string {
	return "iana-if-type:other"
}

type Regular1822 struct {
}

func (id Regular1822) String() string {
	return "iana-if-type:regular1822"
}

type Hdh1822 struct {
}

func (id Hdh1822) String() string {
	return "iana-if-type:hdh1822"
}

type DdnX25 struct {
}

func (id DdnX25) String() string {
	return "iana-if-type:ddnX25"
}

type Rfc877x25 struct {
}

func (id Rfc877x25) String() string {
	return "iana-if-type:rfc877x25"
}

type EthernetCsmacd struct {
}

func (id EthernetCsmacd) String() string {
	return "iana-if-type:ethernetCsmacd"
}

type Iso88023Csmacd struct {
}

func (id Iso88023Csmacd) String() string {
	return "iana-if-type:iso88023Csmacd"
}

type Iso88024TokenBus struct {
}

func (id Iso88024TokenBus) String() string {
	return "iana-if-type:iso88024TokenBus"
}

type Iso88025TokenRing struct {
}

func (id Iso88025TokenRing) String() string {
	return "iana-if-type:iso88025TokenRing"
}

type Iso88026Man struct {
}

func (id Iso88026Man) String() string {
	return "iana-if-type:iso88026Man"
}

type StarLan struct {
}

func (id StarLan) String() string {
	return "iana-if-type:starLan"
}

type Proteon10Mbit struct {
}

func (id Proteon10Mbit) String() string {
	return "iana-if-type:proteon10Mbit"
}

type Proteon80Mbit struct {
}

func (id Proteon80Mbit) String() string {
	return "iana-if-type:proteon80Mbit"
}

type Hyperchannel struct {
}

func (id Hyperchannel) String() string {
	return "iana-if-type:hyperchannel"
}

type Fddi struct {
}

func (id Fddi) String() string {
	return "iana-if-type:fddi"
}

type Lapb struct {
}

func (id Lapb) String() string {
	return "iana-if-type:lapb"
}

type Sdlc struct {
}

func (id Sdlc) String() string {
	return "iana-if-type:sdlc"
}

type Ds1 struct {
}

func (id Ds1) String() string {
	return "iana-if-type:ds1"
}

type E1 struct {
}

func (id E1) String() string {
	return "iana-if-type:e1"
}

type BasicISDN struct {
}

func (id BasicISDN) String() string {
	return "iana-if-type:basicISDN"
}

type PrimaryISDN struct {
}

func (id PrimaryISDN) String() string {
	return "iana-if-type:primaryISDN"
}

type PropPointToPointSerial struct {
}

func (id PropPointToPointSerial) String() string {
	return "iana-if-type:propPointToPointSerial"
}

type Ppp struct {
}

func (id Ppp) String() string {
	return "iana-if-type:ppp"
}

type SoftwareLoopback struct {
}

func (id SoftwareLoopback) String() string {
	return "iana-if-type:softwareLoopback"
}

type Eon struct {
}

func (id Eon) String() string {
	return "iana-if-type:eon"
}

type Ethernet3Mbit struct {
}

func (id Ethernet3Mbit) String() string {
	return "iana-if-type:ethernet3Mbit"
}

type Nsip struct {
}

func (id Nsip) String() string {
	return "iana-if-type:nsip"
}

type Slip struct {
}

func (id Slip) String() string {
	return "iana-if-type:slip"
}

type Ultra struct {
}

func (id Ultra) String() string {
	return "iana-if-type:ultra"
}

type Ds3 struct {
}

func (id Ds3) String() string {
	return "iana-if-type:ds3"
}

type Sip struct {
}

func (id Sip) String() string {
	return "iana-if-type:sip"
}

type FrameRelay struct {
}

func (id FrameRelay) String() string {
	return "iana-if-type:frameRelay"
}

type Rs232 struct {
}

func (id Rs232) String() string {
	return "iana-if-type:rs232"
}

type Para struct {
}

func (id Para) String() string {
	return "iana-if-type:para"
}

type Arcnet struct {
}

func (id Arcnet) String() string {
	return "iana-if-type:arcnet"
}

type ArcnetPlus struct {
}

func (id ArcnetPlus) String() string {
	return "iana-if-type:arcnetPlus"
}

type Atm struct {
}

func (id Atm) String() string {
	return "iana-if-type:atm"
}

type Miox25 struct {
}

func (id Miox25) String() string {
	return "iana-if-type:miox25"
}

type Sonet struct {
}

func (id Sonet) String() string {
	return "iana-if-type:sonet"
}

type X25ple struct {
}

func (id X25ple) String() string {
	return "iana-if-type:x25ple"
}

type Iso88022llc struct {
}

func (id Iso88022llc) String() string {
	return "iana-if-type:iso88022llc"
}

type LocalTalk struct {
}

func (id LocalTalk) String() string {
	return "iana-if-type:localTalk"
}

type SmdsDxi struct {
}

func (id SmdsDxi) String() string {
	return "iana-if-type:smdsDxi"
}

type FrameRelayService struct {
}

func (id FrameRelayService) String() string {
	return "iana-if-type:frameRelayService"
}

type V35 struct {
}

func (id V35) String() string {
	return "iana-if-type:v35"
}

type Hssi struct {
}

func (id Hssi) String() string {
	return "iana-if-type:hssi"
}

type Hippi struct {
}

func (id Hippi) String() string {
	return "iana-if-type:hippi"
}

type Modem struct {
}

func (id Modem) String() string {
	return "iana-if-type:modem"
}

type Aal5 struct {
}

func (id Aal5) String() string {
	return "iana-if-type:aal5"
}

type SonetPath struct {
}

func (id SonetPath) String() string {
	return "iana-if-type:sonetPath"
}

type SonetVT struct {
}

func (id SonetVT) String() string {
	return "iana-if-type:sonetVT"
}

type SmdsIcip struct {
}

func (id SmdsIcip) String() string {
	return "iana-if-type:smdsIcip"
}

type PropVirtual struct {
}

func (id PropVirtual) String() string {
	return "iana-if-type:propVirtual"
}

type PropMultiplexor struct {
}

func (id PropMultiplexor) String() string {
	return "iana-if-type:propMultiplexor"
}

type Ieee80212 struct {
}

func (id Ieee80212) String() string {
	return "iana-if-type:ieee80212"
}

type FibreChannel struct {
}

func (id FibreChannel) String() string {
	return "iana-if-type:fibreChannel"
}

type HippiInterface struct {
}

func (id HippiInterface) String() string {
	return "iana-if-type:hippiInterface"
}

type FrameRelayInterconnect struct {
}

func (id FrameRelayInterconnect) String() string {
	return "iana-if-type:frameRelayInterconnect"
}

type Aflane8023 struct {
}

func (id Aflane8023) String() string {
	return "iana-if-type:aflane8023"
}

type Aflane8025 struct {
}

func (id Aflane8025) String() string {
	return "iana-if-type:aflane8025"
}

type CctEmul struct {
}

func (id CctEmul) String() string {
	return "iana-if-type:cctEmul"
}

type FastEther struct {
}

func (id FastEther) String() string {
	return "iana-if-type:fastEther"
}

type Isdn struct {
}

func (id Isdn) String() string {
	return "iana-if-type:isdn"
}

type V11 struct {
}

func (id V11) String() string {
	return "iana-if-type:v11"
}

type V36 struct {
}

func (id V36) String() string {
	return "iana-if-type:v36"
}

type G703at64k struct {
}

func (id G703at64k) String() string {
	return "iana-if-type:g703at64k"
}

type G703at2mb struct {
}

func (id G703at2mb) String() string {
	return "iana-if-type:g703at2mb"
}

type Qllc struct {
}

func (id Qllc) String() string {
	return "iana-if-type:qllc"
}

type FastEtherFX struct {
}

func (id FastEtherFX) String() string {
	return "iana-if-type:fastEtherFX"
}

type Channel struct {
}

func (id Channel) String() string {
	return "iana-if-type:channel"
}

type Ieee80211 struct {
}

func (id Ieee80211) String() string {
	return "iana-if-type:ieee80211"
}

type Ibm370parChan struct {
}

func (id Ibm370parChan) String() string {
	return "iana-if-type:ibm370parChan"
}

type Escon struct {
}

func (id Escon) String() string {
	return "iana-if-type:escon"
}

type Dlsw struct {
}

func (id Dlsw) String() string {
	return "iana-if-type:dlsw"
}

type Isdns struct {
}

func (id Isdns) String() string {
	return "iana-if-type:isdns"
}

type Isdnu struct {
}

func (id Isdnu) String() string {
	return "iana-if-type:isdnu"
}

type Lapd struct {
}

func (id Lapd) String() string {
	return "iana-if-type:lapd"
}

type IpSwitch struct {
}

func (id IpSwitch) String() string {
	return "iana-if-type:ipSwitch"
}

type Rsrb struct {
}

func (id Rsrb) String() string {
	return "iana-if-type:rsrb"
}

type AtmLogical struct {
}

func (id AtmLogical) String() string {
	return "iana-if-type:atmLogical"
}

type Ds0 struct {
}

func (id Ds0) String() string {
	return "iana-if-type:ds0"
}

type Ds0Bundle struct {
}

func (id Ds0Bundle) String() string {
	return "iana-if-type:ds0Bundle"
}

type Bsc struct {
}

func (id Bsc) String() string {
	return "iana-if-type:bsc"
}

type Async struct {
}

func (id Async) String() string {
	return "iana-if-type:async"
}

type Cnr struct {
}

func (id Cnr) String() string {
	return "iana-if-type:cnr"
}

type Iso88025Dtr struct {
}

func (id Iso88025Dtr) String() string {
	return "iana-if-type:iso88025Dtr"
}

type Eplrs struct {
}

func (id Eplrs) String() string {
	return "iana-if-type:eplrs"
}

type Arap struct {
}

func (id Arap) String() string {
	return "iana-if-type:arap"
}

type PropCnls struct {
}

func (id PropCnls) String() string {
	return "iana-if-type:propCnls"
}

type HostPad struct {
}

func (id HostPad) String() string {
	return "iana-if-type:hostPad"
}

type TermPad struct {
}

func (id TermPad) String() string {
	return "iana-if-type:termPad"
}

type FrameRelayMPI struct {
}

func (id FrameRelayMPI) String() string {
	return "iana-if-type:frameRelayMPI"
}

type X213 struct {
}

func (id X213) String() string {
	return "iana-if-type:x213"
}

type Adsl struct {
}

func (id Adsl) String() string {
	return "iana-if-type:adsl"
}

type Radsl struct {
}

func (id Radsl) String() string {
	return "iana-if-type:radsl"
}

type Sdsl struct {
}

func (id Sdsl) String() string {
	return "iana-if-type:sdsl"
}

type Vdsl struct {
}

func (id Vdsl) String() string {
	return "iana-if-type:vdsl"
}

type Iso88025CRFPInt struct {
}

func (id Iso88025CRFPInt) String() string {
	return "iana-if-type:iso88025CRFPInt"
}

type Myrinet struct {
}

func (id Myrinet) String() string {
	return "iana-if-type:myrinet"
}

type VoiceEM struct {
}

func (id VoiceEM) String() string {
	return "iana-if-type:voiceEM"
}

type VoiceFXO struct {
}

func (id VoiceFXO) String() string {
	return "iana-if-type:voiceFXO"
}

type VoiceFXS struct {
}

func (id VoiceFXS) String() string {
	return "iana-if-type:voiceFXS"
}

type VoiceEncap struct {
}

func (id VoiceEncap) String() string {
	return "iana-if-type:voiceEncap"
}

type VoiceOverIp struct {
}

func (id VoiceOverIp) String() string {
	return "iana-if-type:voiceOverIp"
}

type AtmDxi struct {
}

func (id AtmDxi) String() string {
	return "iana-if-type:atmDxi"
}

type AtmFuni struct {
}

func (id AtmFuni) String() string {
	return "iana-if-type:atmFuni"
}

type AtmIma struct {
}

func (id AtmIma) String() string {
	return "iana-if-type:atmIma"
}

type PppMultilinkBundle struct {
}

func (id PppMultilinkBundle) String() string {
	return "iana-if-type:pppMultilinkBundle"
}

type IpOverCdlc struct {
}

func (id IpOverCdlc) String() string {
	return "iana-if-type:ipOverCdlc"
}

type IpOverClaw struct {
}

func (id IpOverClaw) String() string {
	return "iana-if-type:ipOverClaw"
}

type StackToStack struct {
}

func (id StackToStack) String() string {
	return "iana-if-type:stackToStack"
}

type VirtualIpAddress struct {
}

func (id VirtualIpAddress) String() string {
	return "iana-if-type:virtualIpAddress"
}

type Mpc struct {
}

func (id Mpc) String() string {
	return "iana-if-type:mpc"
}

type IpOverAtm struct {
}

func (id IpOverAtm) String() string {
	return "iana-if-type:ipOverAtm"
}

type Iso88025Fiber struct {
}

func (id Iso88025Fiber) String() string {
	return "iana-if-type:iso88025Fiber"
}

type Tdlc struct {
}

func (id Tdlc) String() string {
	return "iana-if-type:tdlc"
}

type GigabitEthernet struct {
}

func (id GigabitEthernet) String() string {
	return "iana-if-type:gigabitEthernet"
}

type Hdlc struct {
}

func (id Hdlc) String() string {
	return "iana-if-type:hdlc"
}

type Lapf struct {
}

func (id Lapf) String() string {
	return "iana-if-type:lapf"
}

type V37 struct {
}

func (id V37) String() string {
	return "iana-if-type:v37"
}

type X25mlp struct {
}

func (id X25mlp) String() string {
	return "iana-if-type:x25mlp"
}

type X25huntGroup struct {
}

func (id X25huntGroup) String() string {
	return "iana-if-type:x25huntGroup"
}

type TranspHdlc struct {
}

func (id TranspHdlc) String() string {
	return "iana-if-type:transpHdlc"
}

type Interleave struct {
}

func (id Interleave) String() string {
	return "iana-if-type:interleave"
}

type Fast struct {
}

func (id Fast) String() string {
	return "iana-if-type:fast"
}

type Ip struct {
}

func (id Ip) String() string {
	return "iana-if-type:ip"
}

type DocsCableMaclayer struct {
}

func (id DocsCableMaclayer) String() string {
	return "iana-if-type:docsCableMaclayer"
}

type DocsCableDownstream struct {
}

func (id DocsCableDownstream) String() string {
	return "iana-if-type:docsCableDownstream"
}

type DocsCableUpstream struct {
}

func (id DocsCableUpstream) String() string {
	return "iana-if-type:docsCableUpstream"
}

type A12MppSwitch struct {
}

func (id A12MppSwitch) String() string {
	return "iana-if-type:a12MppSwitch"
}

type Tunnel struct {
}

func (id Tunnel) String() string {
	return "iana-if-type:tunnel"
}

type Coffee struct {
}

func (id Coffee) String() string {
	return "iana-if-type:coffee"
}

type Ces struct {
}

func (id Ces) String() string {
	return "iana-if-type:ces"
}

type AtmSubInterface struct {
}

func (id AtmSubInterface) String() string {
	return "iana-if-type:atmSubInterface"
}

type L2vlan struct {
}

func (id L2vlan) String() string {
	return "iana-if-type:l2vlan"
}

type L3ipvlan struct {
}

func (id L3ipvlan) String() string {
	return "iana-if-type:l3ipvlan"
}

type L3ipxvlan struct {
}

func (id L3ipxvlan) String() string {
	return "iana-if-type:l3ipxvlan"
}

type DigitalPowerline struct {
}

func (id DigitalPowerline) String() string {
	return "iana-if-type:digitalPowerline"
}

type MediaMailOverIp struct {
}

func (id MediaMailOverIp) String() string {
	return "iana-if-type:mediaMailOverIp"
}

type Dtm struct {
}

func (id Dtm) String() string {
	return "iana-if-type:dtm"
}

type Dcn struct {
}

func (id Dcn) String() string {
	return "iana-if-type:dcn"
}

type IpForward struct {
}

func (id IpForward) String() string {
	return "iana-if-type:ipForward"
}

type Msdsl struct {
}

func (id Msdsl) String() string {
	return "iana-if-type:msdsl"
}

type Ieee1394 struct {
}

func (id Ieee1394) String() string {
	return "iana-if-type:ieee1394"
}

type IfGsn struct {
}

func (id IfGsn) String() string {
	return "iana-if-type:if-gsn"
}

type DvbRccMacLayer struct {
}

func (id DvbRccMacLayer) String() string {
	return "iana-if-type:dvbRccMacLayer"
}

type DvbRccDownstream struct {
}

func (id DvbRccDownstream) String() string {
	return "iana-if-type:dvbRccDownstream"
}

type DvbRccUpstream struct {
}

func (id DvbRccUpstream) String() string {
	return "iana-if-type:dvbRccUpstream"
}

type AtmVirtual struct {
}

func (id AtmVirtual) String() string {
	return "iana-if-type:atmVirtual"
}

type MplsTunnel struct {
}

func (id MplsTunnel) String() string {
	return "iana-if-type:mplsTunnel"
}

type Srp struct {
}

func (id Srp) String() string {
	return "iana-if-type:srp"
}

type VoiceOverAtm struct {
}

func (id VoiceOverAtm) String() string {
	return "iana-if-type:voiceOverAtm"
}

type VoiceOverFrameRelay struct {
}

func (id VoiceOverFrameRelay) String() string {
	return "iana-if-type:voiceOverFrameRelay"
}

type Idsl struct {
}

func (id Idsl) String() string {
	return "iana-if-type:idsl"
}

type CompositeLink struct {
}

func (id CompositeLink) String() string {
	return "iana-if-type:compositeLink"
}

type Ss7SigLink struct {
}

func (id Ss7SigLink) String() string {
	return "iana-if-type:ss7SigLink"
}

type PropWirelessP2P struct {
}

func (id PropWirelessP2P) String() string {
	return "iana-if-type:propWirelessP2P"
}

type FrForward struct {
}

func (id FrForward) String() string {
	return "iana-if-type:frForward"
}

type Rfc1483 struct {
}

func (id Rfc1483) String() string {
	return "iana-if-type:rfc1483"
}

type Usb struct {
}

func (id Usb) String() string {
	return "iana-if-type:usb"
}

type Ieee8023adLag struct {
}

func (id Ieee8023adLag) String() string {
	return "iana-if-type:ieee8023adLag"
}

type Bgppolicyaccounting struct {
}

func (id Bgppolicyaccounting) String() string {
	return "iana-if-type:bgppolicyaccounting"
}

type Frf16MfrBundle struct {
}

func (id Frf16MfrBundle) String() string {
	return "iana-if-type:frf16MfrBundle"
}

type H323Gatekeeper struct {
}

func (id H323Gatekeeper) String() string {
	return "iana-if-type:h323Gatekeeper"
}

type H323Proxy struct {
}

func (id H323Proxy) String() string {
	return "iana-if-type:h323Proxy"
}

type Mpls struct {
}

func (id Mpls) String() string {
	return "iana-if-type:mpls"
}

type MfSigLink struct {
}

func (id MfSigLink) String() string {
	return "iana-if-type:mfSigLink"
}

type Hdsl2 struct {
}

func (id Hdsl2) String() string {
	return "iana-if-type:hdsl2"
}

type Shdsl struct {
}

func (id Shdsl) String() string {
	return "iana-if-type:shdsl"
}

type Ds1FDL struct {
}

func (id Ds1FDL) String() string {
	return "iana-if-type:ds1FDL"
}

type Pos struct {
}

func (id Pos) String() string {
	return "iana-if-type:pos"
}

type DvbAsiIn struct {
}

func (id DvbAsiIn) String() string {
	return "iana-if-type:dvbAsiIn"
}

type DvbAsiOut struct {
}

func (id DvbAsiOut) String() string {
	return "iana-if-type:dvbAsiOut"
}

type Plc struct {
}

func (id Plc) String() string {
	return "iana-if-type:plc"
}

type Nfas struct {
}

func (id Nfas) String() string {
	return "iana-if-type:nfas"
}

type Tr008 struct {
}

func (id Tr008) String() string {
	return "iana-if-type:tr008"
}

type Gr303RDT struct {
}

func (id Gr303RDT) String() string {
	return "iana-if-type:gr303RDT"
}

type Gr303IDT struct {
}

func (id Gr303IDT) String() string {
	return "iana-if-type:gr303IDT"
}

type Isup struct {
}

func (id Isup) String() string {
	return "iana-if-type:isup"
}

type PropDocsWirelessMaclayer struct {
}

func (id PropDocsWirelessMaclayer) String() string {
	return "iana-if-type:propDocsWirelessMaclayer"
}

type PropDocsWirelessDownstream struct {
}

func (id PropDocsWirelessDownstream) String() string {
	return "iana-if-type:propDocsWirelessDownstream"
}

type PropDocsWirelessUpstream struct {
}

func (id PropDocsWirelessUpstream) String() string {
	return "iana-if-type:propDocsWirelessUpstream"
}

type Hiperlan2 struct {
}

func (id Hiperlan2) String() string {
	return "iana-if-type:hiperlan2"
}

type PropBWAp2Mp struct {
}

func (id PropBWAp2Mp) String() string {
	return "iana-if-type:propBWAp2Mp"
}

type SonetOverheadChannel struct {
}

func (id SonetOverheadChannel) String() string {
	return "iana-if-type:sonetOverheadChannel"
}

type DigitalWrapperOverheadChannel struct {
}

func (id DigitalWrapperOverheadChannel) String() string {
	return "iana-if-type:digitalWrapperOverheadChannel"
}

type Aal2 struct {
}

func (id Aal2) String() string {
	return "iana-if-type:aal2"
}

type RadioMAC struct {
}

func (id RadioMAC) String() string {
	return "iana-if-type:radioMAC"
}

type AtmRadio struct {
}

func (id AtmRadio) String() string {
	return "iana-if-type:atmRadio"
}

type Imt struct {
}

func (id Imt) String() string {
	return "iana-if-type:imt"
}

type Mvl struct {
}

func (id Mvl) String() string {
	return "iana-if-type:mvl"
}

type ReachDSL struct {
}

func (id ReachDSL) String() string {
	return "iana-if-type:reachDSL"
}

type FrDlciEndPt struct {
}

func (id FrDlciEndPt) String() string {
	return "iana-if-type:frDlciEndPt"
}

type AtmVciEndPt struct {
}

func (id AtmVciEndPt) String() string {
	return "iana-if-type:atmVciEndPt"
}

type OpticalChannel struct {
}

func (id OpticalChannel) String() string {
	return "iana-if-type:opticalChannel"
}

type OpticalTransport struct {
}

func (id OpticalTransport) String() string {
	return "iana-if-type:opticalTransport"
}

type PropAtm struct {
}

func (id PropAtm) String() string {
	return "iana-if-type:propAtm"
}

type VoiceOverCable struct {
}

func (id VoiceOverCable) String() string {
	return "iana-if-type:voiceOverCable"
}

type Infiniband struct {
}

func (id Infiniband) String() string {
	return "iana-if-type:infiniband"
}

type TeLink struct {
}

func (id TeLink) String() string {
	return "iana-if-type:teLink"
}

type Q2931 struct {
}

func (id Q2931) String() string {
	return "iana-if-type:q2931"
}

type VirtualTg struct {
}

func (id VirtualTg) String() string {
	return "iana-if-type:virtualTg"
}

type SipTg struct {
}

func (id SipTg) String() string {
	return "iana-if-type:sipTg"
}

type SipSig struct {
}

func (id SipSig) String() string {
	return "iana-if-type:sipSig"
}

type DocsCableUpstreamChannel struct {
}

func (id DocsCableUpstreamChannel) String() string {
	return "iana-if-type:docsCableUpstreamChannel"
}

type Econet struct {
}

func (id Econet) String() string {
	return "iana-if-type:econet"
}

type Pon155 struct {
}

func (id Pon155) String() string {
	return "iana-if-type:pon155"
}

type Pon622 struct {
}

func (id Pon622) String() string {
	return "iana-if-type:pon622"
}

type Bridge struct {
}

func (id Bridge) String() string {
	return "iana-if-type:bridge"
}

type Linegroup struct {
}

func (id Linegroup) String() string {
	return "iana-if-type:linegroup"
}

type VoiceEMFGD struct {
}

func (id VoiceEMFGD) String() string {
	return "iana-if-type:voiceEMFGD"
}

type VoiceFGDEANA struct {
}

func (id VoiceFGDEANA) String() string {
	return "iana-if-type:voiceFGDEANA"
}

type VoiceDID struct {
}

func (id VoiceDID) String() string {
	return "iana-if-type:voiceDID"
}

type MpegTransport struct {
}

func (id MpegTransport) String() string {
	return "iana-if-type:mpegTransport"
}

type SixToFour struct {
}

func (id SixToFour) String() string {
	return "iana-if-type:sixToFour"
}

type Gtp struct {
}

func (id Gtp) String() string {
	return "iana-if-type:gtp"
}

type PdnEtherLoop1 struct {
}

func (id PdnEtherLoop1) String() string {
	return "iana-if-type:pdnEtherLoop1"
}

type PdnEtherLoop2 struct {
}

func (id PdnEtherLoop2) String() string {
	return "iana-if-type:pdnEtherLoop2"
}

type OpticalChannelGroup struct {
}

func (id OpticalChannelGroup) String() string {
	return "iana-if-type:opticalChannelGroup"
}

type Homepna struct {
}

func (id Homepna) String() string {
	return "iana-if-type:homepna"
}

type Gfp struct {
}

func (id Gfp) String() string {
	return "iana-if-type:gfp"
}

type CiscoISLvlan struct {
}

func (id CiscoISLvlan) String() string {
	return "iana-if-type:ciscoISLvlan"
}

type ActelisMetaLOOP struct {
}

func (id ActelisMetaLOOP) String() string {
	return "iana-if-type:actelisMetaLOOP"
}

type FcipLink struct {
}

func (id FcipLink) String() string {
	return "iana-if-type:fcipLink"
}

type Rpr struct {
}

func (id Rpr) String() string {
	return "iana-if-type:rpr"
}

type Qam struct {
}

func (id Qam) String() string {
	return "iana-if-type:qam"
}

type Lmp struct {
}

func (id Lmp) String() string {
	return "iana-if-type:lmp"
}

type CblVectaStar struct {
}

func (id CblVectaStar) String() string {
	return "iana-if-type:cblVectaStar"
}

type DocsCableMCmtsDownstream struct {
}

func (id DocsCableMCmtsDownstream) String() string {
	return "iana-if-type:docsCableMCmtsDownstream"
}

type Adsl2 struct {
}

func (id Adsl2) String() string {
	return "iana-if-type:adsl2"
}

type MacSecControlledIF struct {
}

func (id MacSecControlledIF) String() string {
	return "iana-if-type:macSecControlledIF"
}

type MacSecUncontrolledIF struct {
}

func (id MacSecUncontrolledIF) String() string {
	return "iana-if-type:macSecUncontrolledIF"
}

type AviciOpticalEther struct {
}

func (id AviciOpticalEther) String() string {
	return "iana-if-type:aviciOpticalEther"
}

type Atmbond struct {
}

func (id Atmbond) String() string {
	return "iana-if-type:atmbond"
}

type VoiceFGDOS struct {
}

func (id VoiceFGDOS) String() string {
	return "iana-if-type:voiceFGDOS"
}

type MocaVersion1 struct {
}

func (id MocaVersion1) String() string {
	return "iana-if-type:mocaVersion1"
}

type Ieee80216WMAN struct {
}

func (id Ieee80216WMAN) String() string {
	return "iana-if-type:ieee80216WMAN"
}

type Adsl2plus struct {
}

func (id Adsl2plus) String() string {
	return "iana-if-type:adsl2plus"
}

type DvbRcsMacLayer struct {
}

func (id DvbRcsMacLayer) String() string {
	return "iana-if-type:dvbRcsMacLayer"
}

type DvbTdm struct {
}

func (id DvbTdm) String() string {
	return "iana-if-type:dvbTdm"
}

type DvbRcsTdma struct {
}

func (id DvbRcsTdma) String() string {
	return "iana-if-type:dvbRcsTdma"
}

type X86Laps struct {
}

func (id X86Laps) String() string {
	return "iana-if-type:x86Laps"
}

type WwanPP struct {
}

func (id WwanPP) String() string {
	return "iana-if-type:wwanPP"
}

type WwanPP2 struct {
}

func (id WwanPP2) String() string {
	return "iana-if-type:wwanPP2"
}

type VoiceEBS struct {
}

func (id VoiceEBS) String() string {
	return "iana-if-type:voiceEBS"
}

type IfPwType struct {
}

func (id IfPwType) String() string {
	return "iana-if-type:ifPwType"
}

type Ilan struct {
}

func (id Ilan) String() string {
	return "iana-if-type:ilan"
}

type Pip struct {
}

func (id Pip) String() string {
	return "iana-if-type:pip"
}

type AluELP struct {
}

func (id AluELP) String() string {
	return "iana-if-type:aluELP"
}

type Gpon struct {
}

func (id Gpon) String() string {
	return "iana-if-type:gpon"
}

type Vdsl2 struct {
}

func (id Vdsl2) String() string {
	return "iana-if-type:vdsl2"
}

type CapwapDot11Profile struct {
}

func (id CapwapDot11Profile) String() string {
	return "iana-if-type:capwapDot11Profile"
}

type CapwapDot11Bss struct {
}

func (id CapwapDot11Bss) String() string {
	return "iana-if-type:capwapDot11Bss"
}

type CapwapWtpVirtualRadio struct {
}

func (id CapwapWtpVirtualRadio) String() string {
	return "iana-if-type:capwapWtpVirtualRadio"
}

type Bits struct {
}

func (id Bits) String() string {
	return "iana-if-type:bits"
}

type DocsCableUpstreamRfPort struct {
}

func (id DocsCableUpstreamRfPort) String() string {
	return "iana-if-type:docsCableUpstreamRfPort"
}

type CableDownstreamRfPort struct {
}

func (id CableDownstreamRfPort) String() string {
	return "iana-if-type:cableDownstreamRfPort"
}

type VmwareVirtualNic struct {
}

func (id VmwareVirtualNic) String() string {
	return "iana-if-type:vmwareVirtualNic"
}

type Ieee802154 struct {
}

func (id Ieee802154) String() string {
	return "iana-if-type:ieee802154"
}

type OtnOdu struct {
}

func (id OtnOdu) String() string {
	return "iana-if-type:otnOdu"
}

type OtnOtu struct {
}

func (id OtnOtu) String() string {
	return "iana-if-type:otnOtu"
}

type IfVfiType struct {
}

func (id IfVfiType) String() string {
	return "iana-if-type:ifVfiType"
}

type G9981 struct {
}

func (id G9981) String() string {
	return "iana-if-type:g9981"
}

type G9982 struct {
}

func (id G9982) String() string {
	return "iana-if-type:g9982"
}

type G9983 struct {
}

func (id G9983) String() string {
	return "iana-if-type:g9983"
}

type AluEpon struct {
}

func (id AluEpon) String() string {
	return "iana-if-type:aluEpon"
}

type AluEponOnu struct {
}

func (id AluEponOnu) String() string {
	return "iana-if-type:aluEponOnu"
}

type AluEponPhysicalUni struct {
}

func (id AluEponPhysicalUni) String() string {
	return "iana-if-type:aluEponPhysicalUni"
}

type AluEponLogicalLink struct {
}

func (id AluEponLogicalLink) String() string {
	return "iana-if-type:aluEponLogicalLink"
}

type AluGponOnu struct {
}

func (id AluGponOnu) String() string {
	return "iana-if-type:aluGponOnu"
}

type AluGponPhysicalUni struct {
}

func (id AluGponPhysicalUni) String() string {
	return "iana-if-type:aluGponPhysicalUni"
}

type VmwareNicTeam struct {
}

func (id VmwareNicTeam) String() string {
	return "iana-if-type:vmwareNicTeam"
}

