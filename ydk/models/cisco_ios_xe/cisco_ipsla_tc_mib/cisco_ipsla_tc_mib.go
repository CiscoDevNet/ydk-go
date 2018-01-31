// This MIB contains textual conventions used by
// CISCO IPSLA MIBs.
// 
// Acronyms:
//  FEC: Forward Equivalence Class
//  LPD: Label Path Discovery
//  LSP: Label Switched Path
//  MPLS: Multi Protocol Label Switching
//  RTT: Round Trip Time
//  SAA: Service Assurance Agent
//  SLA: Service Level Agreement
//  VPN: Virtual Private Network
//  ICPIF: Calculated Planning Impairment Factor
package cisco_ipsla_tc_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_ipsla_tc_mib"))
}

// IpSlaCodecType represents g729a(3)             - uses G.729 8000 bps
type IpSlaCodecType string

const (
    IpSlaCodecType_notApplicable IpSlaCodecType = "notApplicable"

    IpSlaCodecType_g711ulaw IpSlaCodecType = "g711ulaw"

    IpSlaCodecType_g711alaw IpSlaCodecType = "g711alaw"

    IpSlaCodecType_g729a IpSlaCodecType = "g729a"
)

// IpSlaOperType represents                  analysis using ICMP timestamp packets.
type IpSlaOperType string

const (
    IpSlaOperType_icmpEcho IpSlaOperType = "icmpEcho"

    IpSlaOperType_udpEcho IpSlaOperType = "udpEcho"

    IpSlaOperType_tcpConnect IpSlaOperType = "tcpConnect"

    IpSlaOperType_udpJitter IpSlaOperType = "udpJitter"

    IpSlaOperType_icmpJitter IpSlaOperType = "icmpJitter"
)

// IpSlaReactVar represents  packetLoss(24)          - Packets loss in both directions
type IpSlaReactVar string

const (
    IpSlaReactVar_rtt IpSlaReactVar = "rtt"

    IpSlaReactVar_jitterSDAvg IpSlaReactVar = "jitterSDAvg"

    IpSlaReactVar_jitterDSAvg IpSlaReactVar = "jitterDSAvg"

    IpSlaReactVar_packetLossSD IpSlaReactVar = "packetLossSD"

    IpSlaReactVar_packetLossDS IpSlaReactVar = "packetLossDS"

    IpSlaReactVar_mos IpSlaReactVar = "mos"

    IpSlaReactVar_timeout IpSlaReactVar = "timeout"

    IpSlaReactVar_connectionLoss IpSlaReactVar = "connectionLoss"

    IpSlaReactVar_verifyError IpSlaReactVar = "verifyError"

    IpSlaReactVar_jitterAvg IpSlaReactVar = "jitterAvg"

    IpSlaReactVar_icpif IpSlaReactVar = "icpif"

    IpSlaReactVar_packetMIA IpSlaReactVar = "packetMIA"

    IpSlaReactVar_packetLateArrival IpSlaReactVar = "packetLateArrival"

    IpSlaReactVar_packetOutOfSequence IpSlaReactVar = "packetOutOfSequence"

    IpSlaReactVar_maxOfPositiveSD IpSlaReactVar = "maxOfPositiveSD"

    IpSlaReactVar_maxOfNegativeSD IpSlaReactVar = "maxOfNegativeSD"

    IpSlaReactVar_maxOfPositiveDS IpSlaReactVar = "maxOfPositiveDS"

    IpSlaReactVar_maxOfNegativeDS IpSlaReactVar = "maxOfNegativeDS"

    IpSlaReactVar_successivePacketLoss IpSlaReactVar = "successivePacketLoss"

    IpSlaReactVar_maxOfLatencyDS IpSlaReactVar = "maxOfLatencyDS"

    IpSlaReactVar_maxOfLatencySD IpSlaReactVar = "maxOfLatencySD"

    IpSlaReactVar_latencyDSAvg IpSlaReactVar = "latencyDSAvg"

    IpSlaReactVar_latencySDAvg IpSlaReactVar = "latencySDAvg"

    IpSlaReactVar_packetLoss IpSlaReactVar = "packetLoss"
)

