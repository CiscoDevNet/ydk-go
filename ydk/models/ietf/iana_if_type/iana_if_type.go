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

type Ddnx25 struct {
}

func (id Ddnx25) String() string {
	return "iana-if-type:ddnX25"
}

type Rfc877X25 struct {
}

func (id Rfc877X25) String() string {
	return "iana-if-type:rfc877x25"
}

type Ethernetcsmacd struct {
}

func (id Ethernetcsmacd) String() string {
	return "iana-if-type:ethernetCsmacd"
}

type Iso88023Csmacd struct {
}

func (id Iso88023Csmacd) String() string {
	return "iana-if-type:iso88023Csmacd"
}

type Iso88024Tokenbus struct {
}

func (id Iso88024Tokenbus) String() string {
	return "iana-if-type:iso88024TokenBus"
}

type Iso88025Tokenring struct {
}

func (id Iso88025Tokenring) String() string {
	return "iana-if-type:iso88025TokenRing"
}

type Iso88026Man struct {
}

func (id Iso88026Man) String() string {
	return "iana-if-type:iso88026Man"
}

type Starlan struct {
}

func (id Starlan) String() string {
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

type Basicisdn struct {
}

func (id Basicisdn) String() string {
	return "iana-if-type:basicISDN"
}

type Primaryisdn struct {
}

func (id Primaryisdn) String() string {
	return "iana-if-type:primaryISDN"
}

type Proppointtopointserial struct {
}

func (id Proppointtopointserial) String() string {
	return "iana-if-type:propPointToPointSerial"
}

type Ppp struct {
}

func (id Ppp) String() string {
	return "iana-if-type:ppp"
}

type Softwareloopback struct {
}

func (id Softwareloopback) String() string {
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

type Framerelay struct {
}

func (id Framerelay) String() string {
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

type Arcnetplus struct {
}

func (id Arcnetplus) String() string {
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

type X25Ple struct {
}

func (id X25Ple) String() string {
	return "iana-if-type:x25ple"
}

type Iso88022Llc struct {
}

func (id Iso88022Llc) String() string {
	return "iana-if-type:iso88022llc"
}

type Localtalk struct {
}

func (id Localtalk) String() string {
	return "iana-if-type:localTalk"
}

type Smdsdxi struct {
}

func (id Smdsdxi) String() string {
	return "iana-if-type:smdsDxi"
}

type Framerelayservice struct {
}

func (id Framerelayservice) String() string {
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

type Sonetpath struct {
}

func (id Sonetpath) String() string {
	return "iana-if-type:sonetPath"
}

type Sonetvt struct {
}

func (id Sonetvt) String() string {
	return "iana-if-type:sonetVT"
}

type Smdsicip struct {
}

func (id Smdsicip) String() string {
	return "iana-if-type:smdsIcip"
}

type Propvirtual struct {
}

func (id Propvirtual) String() string {
	return "iana-if-type:propVirtual"
}

type Propmultiplexor struct {
}

func (id Propmultiplexor) String() string {
	return "iana-if-type:propMultiplexor"
}

type Ieee80212 struct {
}

func (id Ieee80212) String() string {
	return "iana-if-type:ieee80212"
}

type Fibrechannel struct {
}

func (id Fibrechannel) String() string {
	return "iana-if-type:fibreChannel"
}

type Hippiinterface struct {
}

func (id Hippiinterface) String() string {
	return "iana-if-type:hippiInterface"
}

type Framerelayinterconnect struct {
}

func (id Framerelayinterconnect) String() string {
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

type Cctemul struct {
}

func (id Cctemul) String() string {
	return "iana-if-type:cctEmul"
}

type Fastether struct {
}

func (id Fastether) String() string {
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

type G703At64K struct {
}

func (id G703At64K) String() string {
	return "iana-if-type:g703at64k"
}

type G703At2Mb struct {
}

func (id G703At2Mb) String() string {
	return "iana-if-type:g703at2mb"
}

type Qllc struct {
}

func (id Qllc) String() string {
	return "iana-if-type:qllc"
}

type Fastetherfx struct {
}

func (id Fastetherfx) String() string {
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

type Ibm370Parchan struct {
}

func (id Ibm370Parchan) String() string {
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

type Ipswitch struct {
}

func (id Ipswitch) String() string {
	return "iana-if-type:ipSwitch"
}

type Rsrb struct {
}

func (id Rsrb) String() string {
	return "iana-if-type:rsrb"
}

type Atmlogical struct {
}

func (id Atmlogical) String() string {
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

type Propcnls struct {
}

func (id Propcnls) String() string {
	return "iana-if-type:propCnls"
}

type Hostpad struct {
}

func (id Hostpad) String() string {
	return "iana-if-type:hostPad"
}

type Termpad struct {
}

func (id Termpad) String() string {
	return "iana-if-type:termPad"
}

type Framerelaympi struct {
}

func (id Framerelaympi) String() string {
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

type Iso88025Crfpint struct {
}

func (id Iso88025Crfpint) String() string {
	return "iana-if-type:iso88025CRFPInt"
}

type Myrinet struct {
}

func (id Myrinet) String() string {
	return "iana-if-type:myrinet"
}

type Voiceem struct {
}

func (id Voiceem) String() string {
	return "iana-if-type:voiceEM"
}

type Voicefxo struct {
}

func (id Voicefxo) String() string {
	return "iana-if-type:voiceFXO"
}

type Voicefxs struct {
}

func (id Voicefxs) String() string {
	return "iana-if-type:voiceFXS"
}

type Voiceencap struct {
}

func (id Voiceencap) String() string {
	return "iana-if-type:voiceEncap"
}

type Voiceoverip struct {
}

func (id Voiceoverip) String() string {
	return "iana-if-type:voiceOverIp"
}

type Atmdxi struct {
}

func (id Atmdxi) String() string {
	return "iana-if-type:atmDxi"
}

type Atmfuni struct {
}

func (id Atmfuni) String() string {
	return "iana-if-type:atmFuni"
}

type Atmima struct {
}

func (id Atmima) String() string {
	return "iana-if-type:atmIma"
}

type Pppmultilinkbundle struct {
}

func (id Pppmultilinkbundle) String() string {
	return "iana-if-type:pppMultilinkBundle"
}

type Ipovercdlc struct {
}

func (id Ipovercdlc) String() string {
	return "iana-if-type:ipOverCdlc"
}

type Ipoverclaw struct {
}

func (id Ipoverclaw) String() string {
	return "iana-if-type:ipOverClaw"
}

type Stacktostack struct {
}

func (id Stacktostack) String() string {
	return "iana-if-type:stackToStack"
}

type Virtualipaddress struct {
}

func (id Virtualipaddress) String() string {
	return "iana-if-type:virtualIpAddress"
}

type Mpc struct {
}

func (id Mpc) String() string {
	return "iana-if-type:mpc"
}

type Ipoveratm struct {
}

func (id Ipoveratm) String() string {
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

type Gigabitethernet struct {
}

func (id Gigabitethernet) String() string {
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

type X25Mlp struct {
}

func (id X25Mlp) String() string {
	return "iana-if-type:x25mlp"
}

type X25Huntgroup struct {
}

func (id X25Huntgroup) String() string {
	return "iana-if-type:x25huntGroup"
}

type Transphdlc struct {
}

func (id Transphdlc) String() string {
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

type Docscablemaclayer struct {
}

func (id Docscablemaclayer) String() string {
	return "iana-if-type:docsCableMaclayer"
}

type Docscabledownstream struct {
}

func (id Docscabledownstream) String() string {
	return "iana-if-type:docsCableDownstream"
}

type Docscableupstream struct {
}

func (id Docscableupstream) String() string {
	return "iana-if-type:docsCableUpstream"
}

type A12Mppswitch struct {
}

func (id A12Mppswitch) String() string {
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

type Atmsubinterface struct {
}

func (id Atmsubinterface) String() string {
	return "iana-if-type:atmSubInterface"
}

type L2Vlan struct {
}

func (id L2Vlan) String() string {
	return "iana-if-type:l2vlan"
}

type L3Ipvlan struct {
}

func (id L3Ipvlan) String() string {
	return "iana-if-type:l3ipvlan"
}

type L3Ipxvlan struct {
}

func (id L3Ipxvlan) String() string {
	return "iana-if-type:l3ipxvlan"
}

type Digitalpowerline struct {
}

func (id Digitalpowerline) String() string {
	return "iana-if-type:digitalPowerline"
}

type Mediamailoverip struct {
}

func (id Mediamailoverip) String() string {
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

type Ipforward struct {
}

func (id Ipforward) String() string {
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

type Dvbrccmaclayer struct {
}

func (id Dvbrccmaclayer) String() string {
	return "iana-if-type:dvbRccMacLayer"
}

type Dvbrccdownstream struct {
}

func (id Dvbrccdownstream) String() string {
	return "iana-if-type:dvbRccDownstream"
}

type Dvbrccupstream struct {
}

func (id Dvbrccupstream) String() string {
	return "iana-if-type:dvbRccUpstream"
}

type Atmvirtual struct {
}

func (id Atmvirtual) String() string {
	return "iana-if-type:atmVirtual"
}

type Mplstunnel struct {
}

func (id Mplstunnel) String() string {
	return "iana-if-type:mplsTunnel"
}

type Srp struct {
}

func (id Srp) String() string {
	return "iana-if-type:srp"
}

type Voiceoveratm struct {
}

func (id Voiceoveratm) String() string {
	return "iana-if-type:voiceOverAtm"
}

type Voiceoverframerelay struct {
}

func (id Voiceoverframerelay) String() string {
	return "iana-if-type:voiceOverFrameRelay"
}

type Idsl struct {
}

func (id Idsl) String() string {
	return "iana-if-type:idsl"
}

type Compositelink struct {
}

func (id Compositelink) String() string {
	return "iana-if-type:compositeLink"
}

type Ss7Siglink struct {
}

func (id Ss7Siglink) String() string {
	return "iana-if-type:ss7SigLink"
}

type Propwirelessp2P struct {
}

func (id Propwirelessp2P) String() string {
	return "iana-if-type:propWirelessP2P"
}

type Frforward struct {
}

func (id Frforward) String() string {
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

type Ieee8023Adlag struct {
}

func (id Ieee8023Adlag) String() string {
	return "iana-if-type:ieee8023adLag"
}

type Bgppolicyaccounting struct {
}

func (id Bgppolicyaccounting) String() string {
	return "iana-if-type:bgppolicyaccounting"
}

type Frf16Mfrbundle struct {
}

func (id Frf16Mfrbundle) String() string {
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

type Mfsiglink struct {
}

func (id Mfsiglink) String() string {
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

type Ds1Fdl struct {
}

func (id Ds1Fdl) String() string {
	return "iana-if-type:ds1FDL"
}

type Pos struct {
}

func (id Pos) String() string {
	return "iana-if-type:pos"
}

type Dvbasiin struct {
}

func (id Dvbasiin) String() string {
	return "iana-if-type:dvbAsiIn"
}

type Dvbasiout struct {
}

func (id Dvbasiout) String() string {
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

type Gr303Rdt struct {
}

func (id Gr303Rdt) String() string {
	return "iana-if-type:gr303RDT"
}

type Gr303Idt struct {
}

func (id Gr303Idt) String() string {
	return "iana-if-type:gr303IDT"
}

type Isup struct {
}

func (id Isup) String() string {
	return "iana-if-type:isup"
}

type Propdocswirelessmaclayer struct {
}

func (id Propdocswirelessmaclayer) String() string {
	return "iana-if-type:propDocsWirelessMaclayer"
}

type Propdocswirelessdownstream struct {
}

func (id Propdocswirelessdownstream) String() string {
	return "iana-if-type:propDocsWirelessDownstream"
}

type Propdocswirelessupstream struct {
}

func (id Propdocswirelessupstream) String() string {
	return "iana-if-type:propDocsWirelessUpstream"
}

type Hiperlan2 struct {
}

func (id Hiperlan2) String() string {
	return "iana-if-type:hiperlan2"
}

type Propbwap2Mp struct {
}

func (id Propbwap2Mp) String() string {
	return "iana-if-type:propBWAp2Mp"
}

type Sonetoverheadchannel struct {
}

func (id Sonetoverheadchannel) String() string {
	return "iana-if-type:sonetOverheadChannel"
}

type Digitalwrapperoverheadchannel struct {
}

func (id Digitalwrapperoverheadchannel) String() string {
	return "iana-if-type:digitalWrapperOverheadChannel"
}

type Aal2 struct {
}

func (id Aal2) String() string {
	return "iana-if-type:aal2"
}

type Radiomac struct {
}

func (id Radiomac) String() string {
	return "iana-if-type:radioMAC"
}

type Atmradio struct {
}

func (id Atmradio) String() string {
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

type Reachdsl struct {
}

func (id Reachdsl) String() string {
	return "iana-if-type:reachDSL"
}

type Frdlciendpt struct {
}

func (id Frdlciendpt) String() string {
	return "iana-if-type:frDlciEndPt"
}

type Atmvciendpt struct {
}

func (id Atmvciendpt) String() string {
	return "iana-if-type:atmVciEndPt"
}

type Opticalchannel struct {
}

func (id Opticalchannel) String() string {
	return "iana-if-type:opticalChannel"
}

type Opticaltransport struct {
}

func (id Opticaltransport) String() string {
	return "iana-if-type:opticalTransport"
}

type Propatm struct {
}

func (id Propatm) String() string {
	return "iana-if-type:propAtm"
}

type Voiceovercable struct {
}

func (id Voiceovercable) String() string {
	return "iana-if-type:voiceOverCable"
}

type Infiniband struct {
}

func (id Infiniband) String() string {
	return "iana-if-type:infiniband"
}

type Telink struct {
}

func (id Telink) String() string {
	return "iana-if-type:teLink"
}

type Q2931 struct {
}

func (id Q2931) String() string {
	return "iana-if-type:q2931"
}

type Virtualtg struct {
}

func (id Virtualtg) String() string {
	return "iana-if-type:virtualTg"
}

type Siptg struct {
}

func (id Siptg) String() string {
	return "iana-if-type:sipTg"
}

type Sipsig struct {
}

func (id Sipsig) String() string {
	return "iana-if-type:sipSig"
}

type Docscableupstreamchannel struct {
}

func (id Docscableupstreamchannel) String() string {
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

type Voiceemfgd struct {
}

func (id Voiceemfgd) String() string {
	return "iana-if-type:voiceEMFGD"
}

type Voicefgdeana struct {
}

func (id Voicefgdeana) String() string {
	return "iana-if-type:voiceFGDEANA"
}

type Voicedid struct {
}

func (id Voicedid) String() string {
	return "iana-if-type:voiceDID"
}

type Mpegtransport struct {
}

func (id Mpegtransport) String() string {
	return "iana-if-type:mpegTransport"
}

type Sixtofour struct {
}

func (id Sixtofour) String() string {
	return "iana-if-type:sixToFour"
}

type Gtp struct {
}

func (id Gtp) String() string {
	return "iana-if-type:gtp"
}

type Pdnetherloop1 struct {
}

func (id Pdnetherloop1) String() string {
	return "iana-if-type:pdnEtherLoop1"
}

type Pdnetherloop2 struct {
}

func (id Pdnetherloop2) String() string {
	return "iana-if-type:pdnEtherLoop2"
}

type Opticalchannelgroup struct {
}

func (id Opticalchannelgroup) String() string {
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

type Ciscoislvlan struct {
}

func (id Ciscoislvlan) String() string {
	return "iana-if-type:ciscoISLvlan"
}

type Actelismetaloop struct {
}

func (id Actelismetaloop) String() string {
	return "iana-if-type:actelisMetaLOOP"
}

type Fciplink struct {
}

func (id Fciplink) String() string {
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

type Cblvectastar struct {
}

func (id Cblvectastar) String() string {
	return "iana-if-type:cblVectaStar"
}

type Docscablemcmtsdownstream struct {
}

func (id Docscablemcmtsdownstream) String() string {
	return "iana-if-type:docsCableMCmtsDownstream"
}

type Adsl2 struct {
}

func (id Adsl2) String() string {
	return "iana-if-type:adsl2"
}

type Macseccontrolledif struct {
}

func (id Macseccontrolledif) String() string {
	return "iana-if-type:macSecControlledIF"
}

type Macsecuncontrolledif struct {
}

func (id Macsecuncontrolledif) String() string {
	return "iana-if-type:macSecUncontrolledIF"
}

type Aviciopticalether struct {
}

func (id Aviciopticalether) String() string {
	return "iana-if-type:aviciOpticalEther"
}

type Atmbond struct {
}

func (id Atmbond) String() string {
	return "iana-if-type:atmbond"
}

type Voicefgdos struct {
}

func (id Voicefgdos) String() string {
	return "iana-if-type:voiceFGDOS"
}

type Mocaversion1 struct {
}

func (id Mocaversion1) String() string {
	return "iana-if-type:mocaVersion1"
}

type Ieee80216Wman struct {
}

func (id Ieee80216Wman) String() string {
	return "iana-if-type:ieee80216WMAN"
}

type Adsl2Plus struct {
}

func (id Adsl2Plus) String() string {
	return "iana-if-type:adsl2plus"
}

type Dvbrcsmaclayer struct {
}

func (id Dvbrcsmaclayer) String() string {
	return "iana-if-type:dvbRcsMacLayer"
}

type Dvbtdm struct {
}

func (id Dvbtdm) String() string {
	return "iana-if-type:dvbTdm"
}

type Dvbrcstdma struct {
}

func (id Dvbrcstdma) String() string {
	return "iana-if-type:dvbRcsTdma"
}

type X86Laps struct {
}

func (id X86Laps) String() string {
	return "iana-if-type:x86Laps"
}

type Wwanpp struct {
}

func (id Wwanpp) String() string {
	return "iana-if-type:wwanPP"
}

type Wwanpp2 struct {
}

func (id Wwanpp2) String() string {
	return "iana-if-type:wwanPP2"
}

type Voiceebs struct {
}

func (id Voiceebs) String() string {
	return "iana-if-type:voiceEBS"
}

type Ifpwtype struct {
}

func (id Ifpwtype) String() string {
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

type Aluelp struct {
}

func (id Aluelp) String() string {
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

type Capwapdot11Profile struct {
}

func (id Capwapdot11Profile) String() string {
	return "iana-if-type:capwapDot11Profile"
}

type Capwapdot11Bss struct {
}

func (id Capwapdot11Bss) String() string {
	return "iana-if-type:capwapDot11Bss"
}

type Capwapwtpvirtualradio struct {
}

func (id Capwapwtpvirtualradio) String() string {
	return "iana-if-type:capwapWtpVirtualRadio"
}

type Bits struct {
}

func (id Bits) String() string {
	return "iana-if-type:bits"
}

type Docscableupstreamrfport struct {
}

func (id Docscableupstreamrfport) String() string {
	return "iana-if-type:docsCableUpstreamRfPort"
}

type Cabledownstreamrfport struct {
}

func (id Cabledownstreamrfport) String() string {
	return "iana-if-type:cableDownstreamRfPort"
}

type Vmwarevirtualnic struct {
}

func (id Vmwarevirtualnic) String() string {
	return "iana-if-type:vmwareVirtualNic"
}

type Ieee802154 struct {
}

func (id Ieee802154) String() string {
	return "iana-if-type:ieee802154"
}

type Otnodu struct {
}

func (id Otnodu) String() string {
	return "iana-if-type:otnOdu"
}

type Otnotu struct {
}

func (id Otnotu) String() string {
	return "iana-if-type:otnOtu"
}

type Ifvfitype struct {
}

func (id Ifvfitype) String() string {
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

type Aluepon struct {
}

func (id Aluepon) String() string {
	return "iana-if-type:aluEpon"
}

type Aluepononu struct {
}

func (id Aluepononu) String() string {
	return "iana-if-type:aluEponOnu"
}

type Alueponphysicaluni struct {
}

func (id Alueponphysicaluni) String() string {
	return "iana-if-type:aluEponPhysicalUni"
}

type Alueponlogicallink struct {
}

func (id Alueponlogicallink) String() string {
	return "iana-if-type:aluEponLogicalLink"
}

type Alugpononu struct {
}

func (id Alugpononu) String() string {
	return "iana-if-type:aluGponOnu"
}

type Alugponphysicaluni struct {
}

func (id Alugponphysicaluni) String() string {
	return "iana-if-type:aluGponPhysicalUni"
}

type Vmwarenicteam struct {
}

func (id Vmwarenicteam) String() string {
	return "iana-if-type:vmwareNicTeam"
}

