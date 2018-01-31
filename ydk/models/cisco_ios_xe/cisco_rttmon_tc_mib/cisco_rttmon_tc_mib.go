// This MIB contains textual conventions used by
// CISCO-RTTMON-MIB, CISCO-RTTMON-RTP-MIB and 
// CISCO-RTTMON-ICMP-MIB, but they are not limited 
// to only these MIBs. 
// These textual conventions were originally defined in 
// CISCO-RTTMON-MIB.
// 
// Acronyms:
//   FEC: Forward Equivalence Class
//   LPD: Label Path Discovery
//   LSP: Label Switched Path
//   MPLS: Multi Protocol Label Switching
//   RTT: Round Trip Time
//   SAA: Service Assurance Agent
//   VPN: Virtual Private Network
//   CFM: Connection Fault Management
package cisco_rttmon_tc_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_rttmon_tc_mib"))
}

// RttMonRttType represents udp jitter stream analysis on a multicast network.
type RttMonRttType string

const (
    RttMonRttType_none RttMonRttType = "none"

    RttMonRttType_echo RttMonRttType = "echo"

    RttMonRttType_pathEcho RttMonRttType = "pathEcho"

    RttMonRttType_fileIO RttMonRttType = "fileIO"

    RttMonRttType_script RttMonRttType = "script"

    RttMonRttType_udpEcho RttMonRttType = "udpEcho"

    RttMonRttType_tcpConnect RttMonRttType = "tcpConnect"

    RttMonRttType_http RttMonRttType = "http"

    RttMonRttType_dns RttMonRttType = "dns"

    RttMonRttType_jitter RttMonRttType = "jitter"

    RttMonRttType_dlsw RttMonRttType = "dlsw"

    RttMonRttType_dhcp RttMonRttType = "dhcp"

    RttMonRttType_ftp RttMonRttType = "ftp"

    RttMonRttType_voip RttMonRttType = "voip"

    RttMonRttType_rtp RttMonRttType = "rtp"

    RttMonRttType_lspGroup RttMonRttType = "lspGroup"

    RttMonRttType_icmpjitter RttMonRttType = "icmpjitter"

    RttMonRttType_lspPing RttMonRttType = "lspPing"

    RttMonRttType_lspTrace RttMonRttType = "lspTrace"

    RttMonRttType_ethernetPing RttMonRttType = "ethernetPing"

    RttMonRttType_ethernetJitter RttMonRttType = "ethernetJitter"

    RttMonRttType_lspPingPseudowire RttMonRttType = "lspPingPseudowire"

    RttMonRttType_video RttMonRttType = "video"

    RttMonRttType_y1731Delay RttMonRttType = "y1731Delay"

    RttMonRttType_y1731Loss RttMonRttType = "y1731Loss"

    RttMonRttType_mcastJitter RttMonRttType = "mcastJitter"
)

// RttMonLSPPingReplyMode represents                          unreliable.
type RttMonLSPPingReplyMode string

const (
    RttMonLSPPingReplyMode_replyIpv4Udp RttMonLSPPingReplyMode = "replyIpv4Udp"

    RttMonLSPPingReplyMode_replyIpv4UdpRA RttMonLSPPingReplyMode = "replyIpv4UdpRA"
)

// RttMonProtocol represents                       performance
type RttMonProtocol string

const (
    RttMonProtocol_none RttMonProtocol = "none"

    RttMonProtocol_notApplicable RttMonProtocol = "notApplicable"

    RttMonProtocol_ipIcmpEcho RttMonProtocol = "ipIcmpEcho"

    RttMonProtocol_ipUdpEchoAppl RttMonProtocol = "ipUdpEchoAppl"

    RttMonProtocol_snaRUEcho RttMonProtocol = "snaRUEcho"

    RttMonProtocol_snaLU0EchoAppl RttMonProtocol = "snaLU0EchoAppl"

    RttMonProtocol_snaLU2EchoAppl RttMonProtocol = "snaLU2EchoAppl"

    RttMonProtocol_snaLU62Echo RttMonProtocol = "snaLU62Echo"

    RttMonProtocol_snaLU62EchoAppl RttMonProtocol = "snaLU62EchoAppl"

    RttMonProtocol_appleTalkEcho RttMonProtocol = "appleTalkEcho"

    RttMonProtocol_appleTalkEchoAppl RttMonProtocol = "appleTalkEchoAppl"

    RttMonProtocol_decNetEcho RttMonProtocol = "decNetEcho"

    RttMonProtocol_decNetEchoAppl RttMonProtocol = "decNetEchoAppl"

    RttMonProtocol_ipxEcho RttMonProtocol = "ipxEcho"

    RttMonProtocol_ipxEchoAppl RttMonProtocol = "ipxEchoAppl"

    RttMonProtocol_isoClnsEcho RttMonProtocol = "isoClnsEcho"

    RttMonProtocol_isoClnsEchoAppl RttMonProtocol = "isoClnsEchoAppl"

    RttMonProtocol_vinesEcho RttMonProtocol = "vinesEcho"

    RttMonProtocol_vinesEchoAppl RttMonProtocol = "vinesEchoAppl"

    RttMonProtocol_xnsEcho RttMonProtocol = "xnsEcho"

    RttMonProtocol_xnsEchoAppl RttMonProtocol = "xnsEchoAppl"

    RttMonProtocol_apolloEcho RttMonProtocol = "apolloEcho"

    RttMonProtocol_apolloEchoAppl RttMonProtocol = "apolloEchoAppl"

    RttMonProtocol_netbiosEchoAppl RttMonProtocol = "netbiosEchoAppl"

    RttMonProtocol_ipTcpConn RttMonProtocol = "ipTcpConn"

    RttMonProtocol_httpAppl RttMonProtocol = "httpAppl"

    RttMonProtocol_dnsAppl RttMonProtocol = "dnsAppl"

    RttMonProtocol_jitterAppl RttMonProtocol = "jitterAppl"

    RttMonProtocol_dlswAppl RttMonProtocol = "dlswAppl"

    RttMonProtocol_dhcpAppl RttMonProtocol = "dhcpAppl"

    RttMonProtocol_ftpAppl RttMonProtocol = "ftpAppl"

    RttMonProtocol_mplsLspPingAppl RttMonProtocol = "mplsLspPingAppl"

    RttMonProtocol_voipAppl RttMonProtocol = "voipAppl"

    RttMonProtocol_rtpAppl RttMonProtocol = "rtpAppl"

    RttMonProtocol_icmpJitterAppl RttMonProtocol = "icmpJitterAppl"

    RttMonProtocol_ethernetPingAppl RttMonProtocol = "ethernetPingAppl"

    RttMonProtocol_ethernetJitterAppl RttMonProtocol = "ethernetJitterAppl"

    RttMonProtocol_videoAppl RttMonProtocol = "videoAppl"

    RttMonProtocol_y1731dmm RttMonProtocol = "y1731dmm"

    RttMonProtocol_y17311dm RttMonProtocol = "y17311dm"

    RttMonProtocol_y1731lmm RttMonProtocol = "y1731lmm"

    RttMonProtocol_mcastJitterAppl RttMonProtocol = "mcastJitterAppl"

    RttMonProtocol_y1731slm RttMonProtocol = "y1731slm"

    RttMonProtocol_y1731dmmv1 RttMonProtocol = "y1731dmmv1"
)

// RttMplsVpnMonLpdGrpStatus represents              PE.
type RttMplsVpnMonLpdGrpStatus string

const (
    RttMplsVpnMonLpdGrpStatus_unknown RttMplsVpnMonLpdGrpStatus = "unknown"

    RttMplsVpnMonLpdGrpStatus_up RttMplsVpnMonLpdGrpStatus = "up"

    RttMplsVpnMonLpdGrpStatus_partial RttMplsVpnMonLpdGrpStatus = "partial"

    RttMplsVpnMonLpdGrpStatus_down RttMplsVpnMonLpdGrpStatus = "down"
)

// RttMonReactVar represents  rFactorSD(32)           - R-Factor value at Destination.
type RttMonReactVar string

const (
    RttMonReactVar_rtt RttMonReactVar = "rtt"

    RttMonReactVar_jitterSDAvg RttMonReactVar = "jitterSDAvg"

    RttMonReactVar_jitterDSAvg RttMonReactVar = "jitterDSAvg"

    RttMonReactVar_packetLossSD RttMonReactVar = "packetLossSD"

    RttMonReactVar_packetLossDS RttMonReactVar = "packetLossDS"

    RttMonReactVar_mos RttMonReactVar = "mos"

    RttMonReactVar_timeout RttMonReactVar = "timeout"

    RttMonReactVar_connectionLoss RttMonReactVar = "connectionLoss"

    RttMonReactVar_verifyError RttMonReactVar = "verifyError"

    RttMonReactVar_jitterAvg RttMonReactVar = "jitterAvg"

    RttMonReactVar_icpif RttMonReactVar = "icpif"

    RttMonReactVar_packetMIA RttMonReactVar = "packetMIA"

    RttMonReactVar_packetLateArrival RttMonReactVar = "packetLateArrival"

    RttMonReactVar_packetOutOfSequence RttMonReactVar = "packetOutOfSequence"

    RttMonReactVar_maxOfPositiveSD RttMonReactVar = "maxOfPositiveSD"

    RttMonReactVar_maxOfNegativeSD RttMonReactVar = "maxOfNegativeSD"

    RttMonReactVar_maxOfPositiveDS RttMonReactVar = "maxOfPositiveDS"

    RttMonReactVar_maxOfNegativeDS RttMonReactVar = "maxOfNegativeDS"

    RttMonReactVar_iaJitterDS RttMonReactVar = "iaJitterDS"

    RttMonReactVar_frameLossDS RttMonReactVar = "frameLossDS"

    RttMonReactVar_mosLQDS RttMonReactVar = "mosLQDS"

    RttMonReactVar_mosCQDS RttMonReactVar = "mosCQDS"

    RttMonReactVar_rFactorDS RttMonReactVar = "rFactorDS"

    RttMonReactVar_successivePacketLoss RttMonReactVar = "successivePacketLoss"

    RttMonReactVar_maxOfLatencyDS RttMonReactVar = "maxOfLatencyDS"

    RttMonReactVar_maxOfLatencySD RttMonReactVar = "maxOfLatencySD"

    RttMonReactVar_latencyDSAvg RttMonReactVar = "latencyDSAvg"

    RttMonReactVar_latencySDAvg RttMonReactVar = "latencySDAvg"

    RttMonReactVar_packetLoss RttMonReactVar = "packetLoss"

    RttMonReactVar_iaJitterSD RttMonReactVar = "iaJitterSD"

    RttMonReactVar_mosCQSD RttMonReactVar = "mosCQSD"

    RttMonReactVar_rFactorSD RttMonReactVar = "rFactorSD"
)

// RttReset represents the value is 'ready'.
type RttReset string

const (
    RttReset_ready RttReset = "ready"

    RttReset_reset RttReset = "reset"
)

// RttMplsVpnMonLpdFailureSense represents                                   Discovery.
type RttMplsVpnMonLpdFailureSense string

const (
    RttMplsVpnMonLpdFailureSense_unknown RttMplsVpnMonLpdFailureSense = "unknown"

    RttMplsVpnMonLpdFailureSense_noPath RttMplsVpnMonLpdFailureSense = "noPath"

    RttMplsVpnMonLpdFailureSense_allPathsBroken RttMplsVpnMonLpdFailureSense = "allPathsBroken"

    RttMplsVpnMonLpdFailureSense_allPathsUnexplorable RttMplsVpnMonLpdFailureSense = "allPathsUnexplorable"

    RttMplsVpnMonLpdFailureSense_allPathsBrokenOrUnexplorable RttMplsVpnMonLpdFailureSense = "allPathsBrokenOrUnexplorable"

    RttMplsVpnMonLpdFailureSense_timeout RttMplsVpnMonLpdFailureSense = "timeout"

    RttMplsVpnMonLpdFailureSense_error RttMplsVpnMonLpdFailureSense = "error"
)

// RttMonCodecType represents g729a             - uses G.729 8000 bps
type RttMonCodecType string

const (
    RttMonCodecType_notApplicable RttMonCodecType = "notApplicable"

    RttMonCodecType_g711ulaw RttMonCodecType = "g711ulaw"

    RttMonCodecType_g711alaw RttMonCodecType = "g711alaw"

    RttMonCodecType_g729a RttMonCodecType = "g729a"
)

// RttMonOperation represents                            Connect /OK
type RttMonOperation string

const (
    RttMonOperation_notApplicable RttMonOperation = "notApplicable"

    RttMonOperation_httpGet RttMonOperation = "httpGet"

    RttMonOperation_httpRaw RttMonOperation = "httpRaw"

    RttMonOperation_ftpGet RttMonOperation = "ftpGet"

    RttMonOperation_ftpPassive RttMonOperation = "ftpPassive"

    RttMonOperation_ftpActive RttMonOperation = "ftpActive"

    RttMonOperation_voipDTAlertRinging RttMonOperation = "voipDTAlertRinging"

    RttMonOperation_voipDTConnectOK RttMonOperation = "voipDTConnectOK"
)

// RttMplsVpnMonRttType represents automatically configure jitter operations.
type RttMplsVpnMonRttType string

const (
    RttMplsVpnMonRttType_jitter RttMplsVpnMonRttType = "jitter"

    RttMplsVpnMonRttType_echo RttMplsVpnMonRttType = "echo"

    RttMplsVpnMonRttType_pathEcho RttMplsVpnMonRttType = "pathEcho"
)

// RttResponseSense represents                  - Stats retrieve request fail due to port in use.
type RttResponseSense string

const (
    RttResponseSense_other RttResponseSense = "other"

    RttResponseSense_ok RttResponseSense = "ok"

    RttResponseSense_disconnected RttResponseSense = "disconnected"

    RttResponseSense_overThreshold RttResponseSense = "overThreshold"

    RttResponseSense_timeout RttResponseSense = "timeout"

    RttResponseSense_busy RttResponseSense = "busy"

    RttResponseSense_notConnected RttResponseSense = "notConnected"

    RttResponseSense_dropped RttResponseSense = "dropped"

    RttResponseSense_sequenceError RttResponseSense = "sequenceError"

    RttResponseSense_verifyError RttResponseSense = "verifyError"

    RttResponseSense_applicationSpecific RttResponseSense = "applicationSpecific"

    RttResponseSense_dnsServerTimeout RttResponseSense = "dnsServerTimeout"

    RttResponseSense_tcpConnectTimeout RttResponseSense = "tcpConnectTimeout"

    RttResponseSense_httpTransactionTimeout RttResponseSense = "httpTransactionTimeout"

    RttResponseSense_dnsQueryError RttResponseSense = "dnsQueryError"

    RttResponseSense_httpError RttResponseSense = "httpError"

    RttResponseSense_error RttResponseSense = "error"

    RttResponseSense_mplsLspEchoTxError RttResponseSense = "mplsLspEchoTxError"

    RttResponseSense_mplsLspUnreachable RttResponseSense = "mplsLspUnreachable"

    RttResponseSense_mplsLspMalformedReq RttResponseSense = "mplsLspMalformedReq"

    RttResponseSense_mplsLspReachButNotFEC RttResponseSense = "mplsLspReachButNotFEC"

    RttResponseSense_enableOk RttResponseSense = "enableOk"

    RttResponseSense_enableNoConnect RttResponseSense = "enableNoConnect"

    RttResponseSense_enableVersionFail RttResponseSense = "enableVersionFail"

    RttResponseSense_enableInternalError RttResponseSense = "enableInternalError"

    RttResponseSense_enableAbort RttResponseSense = "enableAbort"

    RttResponseSense_enableFail RttResponseSense = "enableFail"

    RttResponseSense_enableAuthFail RttResponseSense = "enableAuthFail"

    RttResponseSense_enableFormatError RttResponseSense = "enableFormatError"

    RttResponseSense_enablePortInUse RttResponseSense = "enablePortInUse"

    RttResponseSense_statsRetrieveOk RttResponseSense = "statsRetrieveOk"

    RttResponseSense_statsRetrieveNoConnect RttResponseSense = "statsRetrieveNoConnect"

    RttResponseSense_statsRetrieveVersionFail RttResponseSense = "statsRetrieveVersionFail"

    RttResponseSense_statsRetrieveInternalError RttResponseSense = "statsRetrieveInternalError"

    RttResponseSense_statsRetrieveAbort RttResponseSense = "statsRetrieveAbort"

    RttResponseSense_statsRetrieveFail RttResponseSense = "statsRetrieveFail"

    RttResponseSense_statsRetrieveAuthFail RttResponseSense = "statsRetrieveAuthFail"

    RttResponseSense_statsRetrieveFormatError RttResponseSense = "statsRetrieveFormatError"

    RttResponseSense_statsRetrievePortInUse RttResponseSense = "statsRetrievePortInUse"
)

