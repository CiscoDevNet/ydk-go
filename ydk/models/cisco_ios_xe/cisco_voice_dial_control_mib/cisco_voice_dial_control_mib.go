// This MIB module enhances the IETF Dial Control MIB
// (RFC2128) by providing management of voice telephony
// peers on both a circuit-switched telephony network,
// and an IP data network.
// 
// *** ABBREVIATIONS, ACRONYMS AND SYMBOLS ***
// 
// GSM    - Global System for Mobile Communication
// 
// AMR-NB - Adaptive Multi Rate - Narrow Band 
// 
// iLBC   - internet Low Bitrate Codec 
// 
// KPML   - Key Press Markup Language
// 
// MGCP   - Media Gateway Control Protocol.
// 
// SIP    - Session Initiation Protocol.
// 
// H323   - One of the voip signalling protocol.
// 
// SCCP   - Skinny Client Control Protocol.
// 
// dialpeer - This represents a configured interface that 
//            carries the dial map.  A dialpeer maps the 
//            called and calling numbers with the port or 
//            IP interface to be used.
// 
// License call capacity - This represents the licensed 
//                         maximum call capacity of 
//                         the gateway.
// 
// iSAC    -  Internet Speech Audio Codec
// 
// RPH    - Resource Priority Header
// 
// DSCP   - Diffserv Code Points
package cisco_voice_dial_control_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_voice_dial_control_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-VOICE-DIAL-CONTROL-MIB CISCO-VOICE-DIAL-CONTROL-MIB}", reflect.TypeOf(CISCOVOICEDIALCONTROLMIB{}))
    ydk.RegisterEntity("CISCO-VOICE-DIAL-CONTROL-MIB:CISCO-VOICE-DIAL-CONTROL-MIB", reflect.TypeOf(CISCOVOICEDIALCONTROLMIB{}))
}

// CvCallVolumeWMIntvlType represents 4 : Uptime Table: Here the entries are from last reload/reboot
type CvCallVolumeWMIntvlType string

const (
    CvCallVolumeWMIntvlType_secondStats CvCallVolumeWMIntvlType = "secondStats"

    CvCallVolumeWMIntvlType_minuteStats CvCallVolumeWMIntvlType = "minuteStats"

    CvCallVolumeWMIntvlType_hourStats CvCallVolumeWMIntvlType = "hourStats"

    CvCallVolumeWMIntvlType_fromReloadStats CvCallVolumeWMIntvlType = "fromReloadStats"
)

// CvIlbcFrameMode represents              frame is packetized in 50 bytes.
type CvIlbcFrameMode string

const (
    CvIlbcFrameMode_frameMode20 CvIlbcFrameMode = "frameMode20"

    CvIlbcFrameMode_frameMode30 CvIlbcFrameMode = "frameMode30"
)

// CvAmrNbRtpEncap represents Represents GSM AMR-NB codec RTP encapsulation type.
type CvAmrNbRtpEncap string

const (
    CvAmrNbRtpEncap_rfc3267 CvAmrNbRtpEncap = "rfc3267"
)

// CvSessionProtocol represents sccp  - Skinny Call Control Protocol.
type CvSessionProtocol string

const (
    CvSessionProtocol_other CvSessionProtocol = "other"

    CvSessionProtocol_cisco CvSessionProtocol = "cisco"

    CvSessionProtocol_sdp CvSessionProtocol = "sdp"

    CvSessionProtocol_sip CvSessionProtocol = "sip"

    CvSessionProtocol_multicast CvSessionProtocol = "multicast"

    CvSessionProtocol_sccp CvSessionProtocol = "sccp"
)

// CvCallConnectionType represents telephony  - telephony signal call connections.
type CvCallConnectionType string

const (
    CvCallConnectionType_h323 CvCallConnectionType = "h323"

    CvCallConnectionType_sip CvCallConnectionType = "sip"

    CvCallConnectionType_mgcp CvCallConnectionType = "mgcp"

    CvCallConnectionType_sccp CvCallConnectionType = "sccp"

    CvCallConnectionType_multicast CvCallConnectionType = "multicast"

    CvCallConnectionType_cacontrol CvCallConnectionType = "cacontrol"

    CvCallConnectionType_telephony CvCallConnectionType = "telephony"
)

// CvCallVolumeStatsIntvlType represents 3 : Hours Table: Here each entry corresponds to an hour
type CvCallVolumeStatsIntvlType string

const (
    CvCallVolumeStatsIntvlType_secondStats CvCallVolumeStatsIntvlType = "secondStats"

    CvCallVolumeStatsIntvlType_minuteStats CvCallVolumeStatsIntvlType = "minuteStats"

    CvCallVolumeStatsIntvlType_hourStats CvCallVolumeStatsIntvlType = "hourStats"
)

// CISCOVOICEDIALCONTROLMIB
type CISCOVOICEDIALCONTROLMIB struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Cvgeneralconfiguration CISCOVOICEDIALCONTROLMIB_Cvgeneralconfiguration

    
    Cvgatewaycallactive CISCOVOICEDIALCONTROLMIB_Cvgatewaycallactive

    
    Cvcallvolume CISCOVOICEDIALCONTROLMIB_Cvcallvolume

    
    Cvcallratemonitor CISCOVOICEDIALCONTROLMIB_Cvcallratemonitor

    
    Cvcallvolumestatshistory CISCOVOICEDIALCONTROLMIB_Cvcallvolumestatshistory

    // The table contains the Voice Generic Peer information that is used to
    // create an ifIndexed row with an appropriate ifType that is associated with
    // the cvPeerCfgType and cvPeerCfgPeerType objects.
    Cvpeercfgtable CISCOVOICEDIALCONTROLMIB_Cvpeercfgtable

    // The table contains the Voice over Telephony peer specific information that
    // is required to accept voice calls or to which it will place them or perform
    // various loopback tests via interface.
    Cvvoicepeercfgtable CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable

    // The table contains the Voice over IP (VoIP) peer specific information that
    // is required to accept voice calls or to which it will place them via IP
    // backbone with the specified session protocol in
    // cvVoIPPeerCfgSessionProtocol.
    Cvvoippeercfgtable CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable

    // The table contains the Voice specific peer common configuration information
    // that is required to accept voice calls or to which it will place them or
    // process the incoming calls.
    Cvpeercommoncfgtable CISCOVOICEDIALCONTROLMIB_Cvpeercommoncfgtable

    // This table is the voice extension to the call active table of IETF Dial
    // Control MIB. It contains voice encapsulation call leg information that is
    // derived from the statistics of lower layer telephony interface.
    Cvcallactivetable CISCOVOICEDIALCONTROLMIB_Cvcallactivetable

    // This table is the VoIP extension to the call active table of IETF Dial
    // Control MIB. It contains VoIP call leg information about specific VoIP call
    // destination and the selected QoS for the call leg.
    Cvvoipcallactivetable CISCOVOICEDIALCONTROLMIB_Cvvoipcallactivetable

    // This table represents the number of active call connections for each call
    // connection type in the voice gateway.
    Cvcallvolconntable CISCOVOICEDIALCONTROLMIB_Cvcallvolconntable

    // This table represents the information about the usage of an IP interface in
    // a voice gateway for voice media calls. This table has a sparse-dependent
    // relationship with   ifTable. There exists an entry in this table,  for each
    // of the  entries in ifTable where ifType  is one of 'ethernetCsmacd' and
    // 'softwareLoopback'.
    Cvcallvoliftable CISCOVOICEDIALCONTROLMIB_Cvcallvoliftable

    // This table is the voice extension to the call history table of IETF Dial
    // Control MIB. It contains voice encapsulation call leg information such as
    // voice packet statistics, coder usage and end to end bandwidth of the call
    // leg.
    Cvcallhistorytable CISCOVOICEDIALCONTROLMIB_Cvcallhistorytable

    // This table is the VoIP extension to the call history table of IETF Dial
    // Control MIB. It contains VoIP call leg information about specific VoIP call
    // destination and the selected QoS for the call leg.
    Cvvoipcallhistorytable CISCOVOICEDIALCONTROLMIB_Cvvoipcallhistorytable

    // This table represents voice call rate measurement in various interval
    // lengths defined by the  CvCallVolumeStatsIntvlType object.  Each interval
    // may contain one or more entries to allow for detailed measurement to be
    // collected.
    Cvcallratestatstable CISCOVOICEDIALCONTROLMIB_Cvcallratestatstable

    // cvCallLegRateStatsTable table represents voice call leg rate measurement in
    // various interval lengths defined by  the CvCallVolumeStatsIntvlType object.
    // Each interval may contain one or more entries to allow for detailed
    // measurement to be collected.
    Cvcalllegratestatstable CISCOVOICEDIALCONTROLMIB_Cvcalllegratestatstable

    // This table represents the active voice calls in various interval lengths
    // defined by the  CvCallVolumeStatsIntvlType object.  Each interval may
    // contain one or more entries to allow for detailed measurement to be
    // collected.
    Cvactivecallstatstable CISCOVOICEDIALCONTROLMIB_Cvactivecallstatstable

    // This table represents the number of calls below a specific duration in
    // various interval length defined by  the CvCallVolumeStatsIntvlType object. 
    // The specific duration is configurable value of  
    // cvCallDurationStatsThreshold object.  Each interval may contain one or more
    // entries to allow for  detailed measurement to be collected.
    Cvcalldurationstatstable CISCOVOICEDIALCONTROLMIB_Cvcalldurationstatstable

    // This table represents the SIP message rate measurement in various interval
    // length defined by the  CvCallVolumeStatsIntvlType object.  Each interval
    // may contain one or more entries to allow for detailed measurement to be
    // collected.
    Cvsipmsgratestatstable CISCOVOICEDIALCONTROLMIB_Cvsipmsgratestatstable

    // This table represents high watermarks achieved by call rate in various
    // interval length defined  by CvCallVolumeWMIntvlType.   Each interval may
    // contain one or more entries to allow for  detailed measurement to be
    // collected.
    Cvcallratewmtable CISCOVOICEDIALCONTROLMIB_Cvcallratewmtable

    // cvCallLegRateWMTable table represents high watermarks achieved by call-leg
    // rate in various interval length defined  by CvCallVolumeWMIntvlType.   Each
    // interval may contain one or more entries to allow for  detailed measurement
    // to be collected.
    Cvcalllegratewmtable CISCOVOICEDIALCONTROLMIB_Cvcalllegratewmtable

    // This table represents high watermarks achieved by active calls in various
    // interval length defined  by CvCallVolumeWMIntvlType.   Each interval may
    // contain one or more entries to allow  for detailed measurement to be
    // collected.
    Cvactivecallwmtable CISCOVOICEDIALCONTROLMIB_Cvactivecallwmtable

    // This table represents of high watermarks achieved by SIP message rate in
    // various interval length defined  by CvCallVolumeWMIntvlType.   Each
    // interval may contain one or more entries to allow for detailed measurement
    // to be collected.
    Cvsipmsgratewmtable CISCOVOICEDIALCONTROLMIB_Cvsipmsgratewmtable
}

func (cISCOVOICEDIALCONTROLMIB *CISCOVOICEDIALCONTROLMIB) GetFilter() yfilter.YFilter { return cISCOVOICEDIALCONTROLMIB.YFilter }

func (cISCOVOICEDIALCONTROLMIB *CISCOVOICEDIALCONTROLMIB) SetFilter(yf yfilter.YFilter) { cISCOVOICEDIALCONTROLMIB.YFilter = yf }

func (cISCOVOICEDIALCONTROLMIB *CISCOVOICEDIALCONTROLMIB) GetGoName(yname string) string {
    if yname == "cvGeneralConfiguration" { return "Cvgeneralconfiguration" }
    if yname == "cvGatewayCallActive" { return "Cvgatewaycallactive" }
    if yname == "cvCallVolume" { return "Cvcallvolume" }
    if yname == "cvCallRateMonitor" { return "Cvcallratemonitor" }
    if yname == "cvCallVolumeStatsHistory" { return "Cvcallvolumestatshistory" }
    if yname == "cvPeerCfgTable" { return "Cvpeercfgtable" }
    if yname == "cvVoicePeerCfgTable" { return "Cvvoicepeercfgtable" }
    if yname == "cvVoIPPeerCfgTable" { return "Cvvoippeercfgtable" }
    if yname == "cvPeerCommonCfgTable" { return "Cvpeercommoncfgtable" }
    if yname == "cvCallActiveTable" { return "Cvcallactivetable" }
    if yname == "cvVoIPCallActiveTable" { return "Cvvoipcallactivetable" }
    if yname == "cvCallVolConnTable" { return "Cvcallvolconntable" }
    if yname == "cvCallVolIfTable" { return "Cvcallvoliftable" }
    if yname == "cvCallHistoryTable" { return "Cvcallhistorytable" }
    if yname == "cvVoIPCallHistoryTable" { return "Cvvoipcallhistorytable" }
    if yname == "cvCallRateStatsTable" { return "Cvcallratestatstable" }
    if yname == "cvCallLegRateStatsTable" { return "Cvcalllegratestatstable" }
    if yname == "cvActiveCallStatsTable" { return "Cvactivecallstatstable" }
    if yname == "cvCallDurationStatsTable" { return "Cvcalldurationstatstable" }
    if yname == "cvSipMsgRateStatsTable" { return "Cvsipmsgratestatstable" }
    if yname == "cvCallRateWMTable" { return "Cvcallratewmtable" }
    if yname == "cvCallLegRateWMTable" { return "Cvcalllegratewmtable" }
    if yname == "cvActiveCallWMTable" { return "Cvactivecallwmtable" }
    if yname == "cvSipMsgRateWMTable" { return "Cvsipmsgratewmtable" }
    return ""
}

func (cISCOVOICEDIALCONTROLMIB *CISCOVOICEDIALCONTROLMIB) GetSegmentPath() string {
    return "CISCO-VOICE-DIAL-CONTROL-MIB:CISCO-VOICE-DIAL-CONTROL-MIB"
}

func (cISCOVOICEDIALCONTROLMIB *CISCOVOICEDIALCONTROLMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cvGeneralConfiguration" {
        return &cISCOVOICEDIALCONTROLMIB.Cvgeneralconfiguration
    }
    if childYangName == "cvGatewayCallActive" {
        return &cISCOVOICEDIALCONTROLMIB.Cvgatewaycallactive
    }
    if childYangName == "cvCallVolume" {
        return &cISCOVOICEDIALCONTROLMIB.Cvcallvolume
    }
    if childYangName == "cvCallRateMonitor" {
        return &cISCOVOICEDIALCONTROLMIB.Cvcallratemonitor
    }
    if childYangName == "cvCallVolumeStatsHistory" {
        return &cISCOVOICEDIALCONTROLMIB.Cvcallvolumestatshistory
    }
    if childYangName == "cvPeerCfgTable" {
        return &cISCOVOICEDIALCONTROLMIB.Cvpeercfgtable
    }
    if childYangName == "cvVoicePeerCfgTable" {
        return &cISCOVOICEDIALCONTROLMIB.Cvvoicepeercfgtable
    }
    if childYangName == "cvVoIPPeerCfgTable" {
        return &cISCOVOICEDIALCONTROLMIB.Cvvoippeercfgtable
    }
    if childYangName == "cvPeerCommonCfgTable" {
        return &cISCOVOICEDIALCONTROLMIB.Cvpeercommoncfgtable
    }
    if childYangName == "cvCallActiveTable" {
        return &cISCOVOICEDIALCONTROLMIB.Cvcallactivetable
    }
    if childYangName == "cvVoIPCallActiveTable" {
        return &cISCOVOICEDIALCONTROLMIB.Cvvoipcallactivetable
    }
    if childYangName == "cvCallVolConnTable" {
        return &cISCOVOICEDIALCONTROLMIB.Cvcallvolconntable
    }
    if childYangName == "cvCallVolIfTable" {
        return &cISCOVOICEDIALCONTROLMIB.Cvcallvoliftable
    }
    if childYangName == "cvCallHistoryTable" {
        return &cISCOVOICEDIALCONTROLMIB.Cvcallhistorytable
    }
    if childYangName == "cvVoIPCallHistoryTable" {
        return &cISCOVOICEDIALCONTROLMIB.Cvvoipcallhistorytable
    }
    if childYangName == "cvCallRateStatsTable" {
        return &cISCOVOICEDIALCONTROLMIB.Cvcallratestatstable
    }
    if childYangName == "cvCallLegRateStatsTable" {
        return &cISCOVOICEDIALCONTROLMIB.Cvcalllegratestatstable
    }
    if childYangName == "cvActiveCallStatsTable" {
        return &cISCOVOICEDIALCONTROLMIB.Cvactivecallstatstable
    }
    if childYangName == "cvCallDurationStatsTable" {
        return &cISCOVOICEDIALCONTROLMIB.Cvcalldurationstatstable
    }
    if childYangName == "cvSipMsgRateStatsTable" {
        return &cISCOVOICEDIALCONTROLMIB.Cvsipmsgratestatstable
    }
    if childYangName == "cvCallRateWMTable" {
        return &cISCOVOICEDIALCONTROLMIB.Cvcallratewmtable
    }
    if childYangName == "cvCallLegRateWMTable" {
        return &cISCOVOICEDIALCONTROLMIB.Cvcalllegratewmtable
    }
    if childYangName == "cvActiveCallWMTable" {
        return &cISCOVOICEDIALCONTROLMIB.Cvactivecallwmtable
    }
    if childYangName == "cvSipMsgRateWMTable" {
        return &cISCOVOICEDIALCONTROLMIB.Cvsipmsgratewmtable
    }
    return nil
}

func (cISCOVOICEDIALCONTROLMIB *CISCOVOICEDIALCONTROLMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cvGeneralConfiguration"] = &cISCOVOICEDIALCONTROLMIB.Cvgeneralconfiguration
    children["cvGatewayCallActive"] = &cISCOVOICEDIALCONTROLMIB.Cvgatewaycallactive
    children["cvCallVolume"] = &cISCOVOICEDIALCONTROLMIB.Cvcallvolume
    children["cvCallRateMonitor"] = &cISCOVOICEDIALCONTROLMIB.Cvcallratemonitor
    children["cvCallVolumeStatsHistory"] = &cISCOVOICEDIALCONTROLMIB.Cvcallvolumestatshistory
    children["cvPeerCfgTable"] = &cISCOVOICEDIALCONTROLMIB.Cvpeercfgtable
    children["cvVoicePeerCfgTable"] = &cISCOVOICEDIALCONTROLMIB.Cvvoicepeercfgtable
    children["cvVoIPPeerCfgTable"] = &cISCOVOICEDIALCONTROLMIB.Cvvoippeercfgtable
    children["cvPeerCommonCfgTable"] = &cISCOVOICEDIALCONTROLMIB.Cvpeercommoncfgtable
    children["cvCallActiveTable"] = &cISCOVOICEDIALCONTROLMIB.Cvcallactivetable
    children["cvVoIPCallActiveTable"] = &cISCOVOICEDIALCONTROLMIB.Cvvoipcallactivetable
    children["cvCallVolConnTable"] = &cISCOVOICEDIALCONTROLMIB.Cvcallvolconntable
    children["cvCallVolIfTable"] = &cISCOVOICEDIALCONTROLMIB.Cvcallvoliftable
    children["cvCallHistoryTable"] = &cISCOVOICEDIALCONTROLMIB.Cvcallhistorytable
    children["cvVoIPCallHistoryTable"] = &cISCOVOICEDIALCONTROLMIB.Cvvoipcallhistorytable
    children["cvCallRateStatsTable"] = &cISCOVOICEDIALCONTROLMIB.Cvcallratestatstable
    children["cvCallLegRateStatsTable"] = &cISCOVOICEDIALCONTROLMIB.Cvcalllegratestatstable
    children["cvActiveCallStatsTable"] = &cISCOVOICEDIALCONTROLMIB.Cvactivecallstatstable
    children["cvCallDurationStatsTable"] = &cISCOVOICEDIALCONTROLMIB.Cvcalldurationstatstable
    children["cvSipMsgRateStatsTable"] = &cISCOVOICEDIALCONTROLMIB.Cvsipmsgratestatstable
    children["cvCallRateWMTable"] = &cISCOVOICEDIALCONTROLMIB.Cvcallratewmtable
    children["cvCallLegRateWMTable"] = &cISCOVOICEDIALCONTROLMIB.Cvcalllegratewmtable
    children["cvActiveCallWMTable"] = &cISCOVOICEDIALCONTROLMIB.Cvactivecallwmtable
    children["cvSipMsgRateWMTable"] = &cISCOVOICEDIALCONTROLMIB.Cvsipmsgratewmtable
    return children
}

func (cISCOVOICEDIALCONTROLMIB *CISCOVOICEDIALCONTROLMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOVOICEDIALCONTROLMIB *CISCOVOICEDIALCONTROLMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOVOICEDIALCONTROLMIB *CISCOVOICEDIALCONTROLMIB) GetYangName() string { return "CISCO-VOICE-DIAL-CONTROL-MIB" }

func (cISCOVOICEDIALCONTROLMIB *CISCOVOICEDIALCONTROLMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOVOICEDIALCONTROLMIB *CISCOVOICEDIALCONTROLMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOVOICEDIALCONTROLMIB *CISCOVOICEDIALCONTROLMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOVOICEDIALCONTROLMIB *CISCOVOICEDIALCONTROLMIB) SetParent(parent types.Entity) { cISCOVOICEDIALCONTROLMIB.parent = parent }

func (cISCOVOICEDIALCONTROLMIB *CISCOVOICEDIALCONTROLMIB) GetParent() types.Entity { return cISCOVOICEDIALCONTROLMIB.parent }

func (cISCOVOICEDIALCONTROLMIB *CISCOVOICEDIALCONTROLMIB) GetParentYangName() string { return "CISCO-VOICE-DIAL-CONTROL-MIB" }

// CISCOVOICEDIALCONTROLMIB_Cvgeneralconfiguration
type CISCOVOICEDIALCONTROLMIB_Cvgeneralconfiguration struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This object indicates whether cvdcPoorQoVNotification (or the newer
    // cvdcPoorQoVNotificationRev1) traps should be generated for a poor quality
    // of voice calls.  If the value of this object is 'true',
    // cvdcPoorQoVNotification (or the newer cvdcPoorQoVNotificationRev1) traps
    // will be generated for all voice over IP peers when a poor quality of voice
    // call condition is detected after the voice gateway call disconnection.  If
    // the value of this object is 'false', cvdcPoorQoVNotification (or the newer
    // cvdcPoorQoVNotificationRev1) traps will be generated only for calls for
    // which the cvVoIPPeerCfgPoorQoVNotificationEnable object of voice over IP
    // peers having set to 'true'. The type is bool.
    Cvgeneralpoorqovnotificationenable interface{}

    // This object indicates whether cvdcFallbackNotification traps should be
    // generated for fallback. If the value of this object is 'true',
    // cvdcFallbackNotification traps will be generated for all voice over IP
    // peers. The type is bool.
    Cvgeneralfallbacknotificationenable interface{}

    // This object indicates whether cvdcPolicyViolationNotification traps should
    // be generated for a RPH to DSCP mapping violation for SIP voice calls.  If
    // the value of this object is 'true', cvdcPolicyViolationNotification traps
    // will be generated for SIP voice over IP peers when a RPH to DSCP violation
    // condition is detected .  If the value of this object is 'false',
    // cvdcPolicyViolationNotification traps will be generated only for calls for
    // which the  cvVoIPPeerCfgDSCPPolicyNotificationEnable object of voice over
    // IP peers having set to 'true'. The type is bool.
    Cvgeneraldscppolicynotificationenable interface{}

    // This object indicates whether cvdcPolicyViolationNotification traps should
    // be generated for Media violation for SIP voice calls.  If the value of this
    // object is 'true', cvdcPolicyViolationNotification traps will be generated
    // for SIP voice over IP peers when media violation condition is detected . 
    // If the value of this object is 'false', cvdcPolicyViolationNotification
    // traps will be generated only for calls for which the 
    // cvVoIPPeerCfgMediaPolicyNotificationEnable object of voice over IP peers
    // having set to 'true'. The type is bool.
    Cvgeneralmediapolicynotificationenable interface{}
}

func (cvgeneralconfiguration *CISCOVOICEDIALCONTROLMIB_Cvgeneralconfiguration) GetFilter() yfilter.YFilter { return cvgeneralconfiguration.YFilter }

func (cvgeneralconfiguration *CISCOVOICEDIALCONTROLMIB_Cvgeneralconfiguration) SetFilter(yf yfilter.YFilter) { cvgeneralconfiguration.YFilter = yf }

func (cvgeneralconfiguration *CISCOVOICEDIALCONTROLMIB_Cvgeneralconfiguration) GetGoName(yname string) string {
    if yname == "cvGeneralPoorQoVNotificationEnable" { return "Cvgeneralpoorqovnotificationenable" }
    if yname == "cvGeneralFallbackNotificationEnable" { return "Cvgeneralfallbacknotificationenable" }
    if yname == "cvGeneralDSCPPolicyNotificationEnable" { return "Cvgeneraldscppolicynotificationenable" }
    if yname == "cvGeneralMediaPolicyNotificationEnable" { return "Cvgeneralmediapolicynotificationenable" }
    return ""
}

func (cvgeneralconfiguration *CISCOVOICEDIALCONTROLMIB_Cvgeneralconfiguration) GetSegmentPath() string {
    return "cvGeneralConfiguration"
}

func (cvgeneralconfiguration *CISCOVOICEDIALCONTROLMIB_Cvgeneralconfiguration) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cvgeneralconfiguration *CISCOVOICEDIALCONTROLMIB_Cvgeneralconfiguration) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cvgeneralconfiguration *CISCOVOICEDIALCONTROLMIB_Cvgeneralconfiguration) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cvGeneralPoorQoVNotificationEnable"] = cvgeneralconfiguration.Cvgeneralpoorqovnotificationenable
    leafs["cvGeneralFallbackNotificationEnable"] = cvgeneralconfiguration.Cvgeneralfallbacknotificationenable
    leafs["cvGeneralDSCPPolicyNotificationEnable"] = cvgeneralconfiguration.Cvgeneraldscppolicynotificationenable
    leafs["cvGeneralMediaPolicyNotificationEnable"] = cvgeneralconfiguration.Cvgeneralmediapolicynotificationenable
    return leafs
}

func (cvgeneralconfiguration *CISCOVOICEDIALCONTROLMIB_Cvgeneralconfiguration) GetBundleName() string { return "cisco_ios_xe" }

func (cvgeneralconfiguration *CISCOVOICEDIALCONTROLMIB_Cvgeneralconfiguration) GetYangName() string { return "cvGeneralConfiguration" }

func (cvgeneralconfiguration *CISCOVOICEDIALCONTROLMIB_Cvgeneralconfiguration) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvgeneralconfiguration *CISCOVOICEDIALCONTROLMIB_Cvgeneralconfiguration) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvgeneralconfiguration *CISCOVOICEDIALCONTROLMIB_Cvgeneralconfiguration) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvgeneralconfiguration *CISCOVOICEDIALCONTROLMIB_Cvgeneralconfiguration) SetParent(parent types.Entity) { cvgeneralconfiguration.parent = parent }

func (cvgeneralconfiguration *CISCOVOICEDIALCONTROLMIB_Cvgeneralconfiguration) GetParent() types.Entity { return cvgeneralconfiguration.parent }

func (cvgeneralconfiguration *CISCOVOICEDIALCONTROLMIB_Cvgeneralconfiguration) GetParentYangName() string { return "CISCO-VOICE-DIAL-CONTROL-MIB" }

// CISCOVOICEDIALCONTROLMIB_Cvgatewaycallactive
type CISCOVOICEDIALCONTROLMIB_Cvgatewaycallactive struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The current number of DS0 interfaces used for the active calls. The type is
    // interface{} with range: 0..4294967295. Units are interfaces.
    Cvcallactiveds0S interface{}

    // A high threshold used to determine when to generate the
    // cvdcActiveDS0sHighNotification. This object  represents the percentage of
    // active DS0s in total number  of DS0s. The type is interface{} with range:
    // 0..100. Units are percent.
    Cvcallactiveds0Shighthreshold interface{}

    // A low threshold used to determine when to generate the
    // cvdcActiveDS0sLowNotification notification. This object  represents the
    // percentage of active DS0s in total number  of DS0s. The type is interface{}
    // with range: 0..100. Units are percent.
    Cvcallactiveds0Slowthreshold interface{}

    // Specifies whether or not cvdcActiveDS0sHighNotification should be
    // generated.  'true' : Indicates that the cvdcActiveDS0sHighNotification     
    // generation is enabled.  'false': Indicates that
    // cvdcActiveDS0sHighNotification          generation is disabled. The type is
    // bool.
    Cvcallactiveds0Shighnotifyenable interface{}

    // Specifies whether or not cvdcActiveDS0sLowNotification should be generated.
    // 'true' : Indicates that the cvdcActiveDS0sLowNotification         
    // generation is enabled.  'false': Indicates that
    // cvdcActiveDS0sLowNotification          generation is disabled. The type is
    // bool.
    Cvcallactiveds0Slownotifyenable interface{}
}

func (cvgatewaycallactive *CISCOVOICEDIALCONTROLMIB_Cvgatewaycallactive) GetFilter() yfilter.YFilter { return cvgatewaycallactive.YFilter }

func (cvgatewaycallactive *CISCOVOICEDIALCONTROLMIB_Cvgatewaycallactive) SetFilter(yf yfilter.YFilter) { cvgatewaycallactive.YFilter = yf }

func (cvgatewaycallactive *CISCOVOICEDIALCONTROLMIB_Cvgatewaycallactive) GetGoName(yname string) string {
    if yname == "cvCallActiveDS0s" { return "Cvcallactiveds0S" }
    if yname == "cvCallActiveDS0sHighThreshold" { return "Cvcallactiveds0Shighthreshold" }
    if yname == "cvCallActiveDS0sLowThreshold" { return "Cvcallactiveds0Slowthreshold" }
    if yname == "cvCallActiveDS0sHighNotifyEnable" { return "Cvcallactiveds0Shighnotifyenable" }
    if yname == "cvCallActiveDS0sLowNotifyEnable" { return "Cvcallactiveds0Slownotifyenable" }
    return ""
}

func (cvgatewaycallactive *CISCOVOICEDIALCONTROLMIB_Cvgatewaycallactive) GetSegmentPath() string {
    return "cvGatewayCallActive"
}

func (cvgatewaycallactive *CISCOVOICEDIALCONTROLMIB_Cvgatewaycallactive) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cvgatewaycallactive *CISCOVOICEDIALCONTROLMIB_Cvgatewaycallactive) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cvgatewaycallactive *CISCOVOICEDIALCONTROLMIB_Cvgatewaycallactive) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cvCallActiveDS0s"] = cvgatewaycallactive.Cvcallactiveds0S
    leafs["cvCallActiveDS0sHighThreshold"] = cvgatewaycallactive.Cvcallactiveds0Shighthreshold
    leafs["cvCallActiveDS0sLowThreshold"] = cvgatewaycallactive.Cvcallactiveds0Slowthreshold
    leafs["cvCallActiveDS0sHighNotifyEnable"] = cvgatewaycallactive.Cvcallactiveds0Shighnotifyenable
    leafs["cvCallActiveDS0sLowNotifyEnable"] = cvgatewaycallactive.Cvcallactiveds0Slownotifyenable
    return leafs
}

func (cvgatewaycallactive *CISCOVOICEDIALCONTROLMIB_Cvgatewaycallactive) GetBundleName() string { return "cisco_ios_xe" }

func (cvgatewaycallactive *CISCOVOICEDIALCONTROLMIB_Cvgatewaycallactive) GetYangName() string { return "cvGatewayCallActive" }

func (cvgatewaycallactive *CISCOVOICEDIALCONTROLMIB_Cvgatewaycallactive) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvgatewaycallactive *CISCOVOICEDIALCONTROLMIB_Cvgatewaycallactive) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvgatewaycallactive *CISCOVOICEDIALCONTROLMIB_Cvgatewaycallactive) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvgatewaycallactive *CISCOVOICEDIALCONTROLMIB_Cvgatewaycallactive) SetParent(parent types.Entity) { cvgatewaycallactive.parent = parent }

func (cvgatewaycallactive *CISCOVOICEDIALCONTROLMIB_Cvgatewaycallactive) GetParent() types.Entity { return cvgatewaycallactive.parent }

func (cvgatewaycallactive *CISCOVOICEDIALCONTROLMIB_Cvgatewaycallactive) GetParentYangName() string { return "CISCO-VOICE-DIAL-CONTROL-MIB" }

// CISCOVOICEDIALCONTROLMIB_Cvcallvolume
type CISCOVOICEDIALCONTROLMIB_Cvcallvolume struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This object represents the total number of active call legs in the voice
    // gateway. The type is interface{} with range: 0..65535.
    Cvcallvolconntotalactiveconnections interface{}

    // This object represents the licensed call capacity for a voice gateway.  If
    // the value is 0, no  licensing is done and the gateway can be  accomodate as
    // many calls depending on its capability. The type is interface{} with range:
    // 0..65535.
    Cvcallvolconnmaxcallconnectionlicenese interface{}
}

func (cvcallvolume *CISCOVOICEDIALCONTROLMIB_Cvcallvolume) GetFilter() yfilter.YFilter { return cvcallvolume.YFilter }

func (cvcallvolume *CISCOVOICEDIALCONTROLMIB_Cvcallvolume) SetFilter(yf yfilter.YFilter) { cvcallvolume.YFilter = yf }

func (cvcallvolume *CISCOVOICEDIALCONTROLMIB_Cvcallvolume) GetGoName(yname string) string {
    if yname == "cvCallVolConnTotalActiveConnections" { return "Cvcallvolconntotalactiveconnections" }
    if yname == "cvCallVolConnMaxCallConnectionLicenese" { return "Cvcallvolconnmaxcallconnectionlicenese" }
    return ""
}

func (cvcallvolume *CISCOVOICEDIALCONTROLMIB_Cvcallvolume) GetSegmentPath() string {
    return "cvCallVolume"
}

func (cvcallvolume *CISCOVOICEDIALCONTROLMIB_Cvcallvolume) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cvcallvolume *CISCOVOICEDIALCONTROLMIB_Cvcallvolume) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cvcallvolume *CISCOVOICEDIALCONTROLMIB_Cvcallvolume) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cvCallVolConnTotalActiveConnections"] = cvcallvolume.Cvcallvolconntotalactiveconnections
    leafs["cvCallVolConnMaxCallConnectionLicenese"] = cvcallvolume.Cvcallvolconnmaxcallconnectionlicenese
    return leafs
}

func (cvcallvolume *CISCOVOICEDIALCONTROLMIB_Cvcallvolume) GetBundleName() string { return "cisco_ios_xe" }

func (cvcallvolume *CISCOVOICEDIALCONTROLMIB_Cvcallvolume) GetYangName() string { return "cvCallVolume" }

func (cvcallvolume *CISCOVOICEDIALCONTROLMIB_Cvcallvolume) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvcallvolume *CISCOVOICEDIALCONTROLMIB_Cvcallvolume) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvcallvolume *CISCOVOICEDIALCONTROLMIB_Cvcallvolume) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvcallvolume *CISCOVOICEDIALCONTROLMIB_Cvcallvolume) SetParent(parent types.Entity) { cvcallvolume.parent = parent }

func (cvcallvolume *CISCOVOICEDIALCONTROLMIB_Cvcallvolume) GetParent() types.Entity { return cvcallvolume.parent }

func (cvcallvolume *CISCOVOICEDIALCONTROLMIB_Cvcallvolume) GetParentYangName() string { return "CISCO-VOICE-DIAL-CONTROL-MIB" }

// CISCOVOICEDIALCONTROLMIB_Cvcallratemonitor
type CISCOVOICEDIALCONTROLMIB_Cvcallratemonitor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This object represents the state of call-monitoring. A value of 'true'
    // indicates that call-monitoring  is enabled.  A value of 'false' indicates
    // that  call-monitoring is disabled. The type is bool.
    Cvcallratemonitorenable interface{}

    // This object represents the interval for which the gateway monitors the
    // call-rate. The type is interface{} with range: 1..12. Units are five
    // seconds.
    Cvcallratemonitortime interface{}

    // This object represents the total number of calls handled by the gateway
    // during the  monitored time. The type is interface{} with range: 0..65535.
    Cvcallrate interface{}

    // This object represents the high water mark for the number of calls handled
    // by the  gateway in an unit interval of  cvCallRateMonitorTime, from the
    // time  the call-monitoring is enabled. The type is interface{} with range:
    // 0..65535.
    Cvcallratehiwatermark interface{}
}

func (cvcallratemonitor *CISCOVOICEDIALCONTROLMIB_Cvcallratemonitor) GetFilter() yfilter.YFilter { return cvcallratemonitor.YFilter }

func (cvcallratemonitor *CISCOVOICEDIALCONTROLMIB_Cvcallratemonitor) SetFilter(yf yfilter.YFilter) { cvcallratemonitor.YFilter = yf }

func (cvcallratemonitor *CISCOVOICEDIALCONTROLMIB_Cvcallratemonitor) GetGoName(yname string) string {
    if yname == "cvCallRateMonitorEnable" { return "Cvcallratemonitorenable" }
    if yname == "cvCallRateMonitorTime" { return "Cvcallratemonitortime" }
    if yname == "cvCallRate" { return "Cvcallrate" }
    if yname == "cvCallRateHiWaterMark" { return "Cvcallratehiwatermark" }
    return ""
}

func (cvcallratemonitor *CISCOVOICEDIALCONTROLMIB_Cvcallratemonitor) GetSegmentPath() string {
    return "cvCallRateMonitor"
}

func (cvcallratemonitor *CISCOVOICEDIALCONTROLMIB_Cvcallratemonitor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cvcallratemonitor *CISCOVOICEDIALCONTROLMIB_Cvcallratemonitor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cvcallratemonitor *CISCOVOICEDIALCONTROLMIB_Cvcallratemonitor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cvCallRateMonitorEnable"] = cvcallratemonitor.Cvcallratemonitorenable
    leafs["cvCallRateMonitorTime"] = cvcallratemonitor.Cvcallratemonitortime
    leafs["cvCallRate"] = cvcallratemonitor.Cvcallrate
    leafs["cvCallRateHiWaterMark"] = cvcallratemonitor.Cvcallratehiwatermark
    return leafs
}

func (cvcallratemonitor *CISCOVOICEDIALCONTROLMIB_Cvcallratemonitor) GetBundleName() string { return "cisco_ios_xe" }

func (cvcallratemonitor *CISCOVOICEDIALCONTROLMIB_Cvcallratemonitor) GetYangName() string { return "cvCallRateMonitor" }

func (cvcallratemonitor *CISCOVOICEDIALCONTROLMIB_Cvcallratemonitor) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvcallratemonitor *CISCOVOICEDIALCONTROLMIB_Cvcallratemonitor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvcallratemonitor *CISCOVOICEDIALCONTROLMIB_Cvcallratemonitor) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvcallratemonitor *CISCOVOICEDIALCONTROLMIB_Cvcallratemonitor) SetParent(parent types.Entity) { cvcallratemonitor.parent = parent }

func (cvcallratemonitor *CISCOVOICEDIALCONTROLMIB_Cvcallratemonitor) GetParent() types.Entity { return cvcallratemonitor.parent }

func (cvcallratemonitor *CISCOVOICEDIALCONTROLMIB_Cvcallratemonitor) GetParentYangName() string { return "CISCO-VOICE-DIAL-CONTROL-MIB" }

// CISCOVOICEDIALCONTROLMIB_Cvcallvolumestatshistory
type CISCOVOICEDIALCONTROLMIB_Cvcallvolumestatshistory struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This Object specifies the thresold duration in seconds.
    // cvCallDurationStatsTable will monitor all the calls below this  threshold. 
    // Decresing the value of the threshold will reset this table. The type is
    // interface{} with range: 1..3600. Units are seconds.
    Cvcalldurationstatsthreshold interface{}

    // This Object specifies the number of entries the watermark table will
    // maintain.  This value will decide the number of elements in
    // cvCallRateWMTable, cvCallLegRateWMTable, cvActiveCallWMTable and
    // cvSipMsgRateWMTable. The type is interface{} with range: 3..10.
    Cvcallvolumewmtablesize interface{}
}

func (cvcallvolumestatshistory *CISCOVOICEDIALCONTROLMIB_Cvcallvolumestatshistory) GetFilter() yfilter.YFilter { return cvcallvolumestatshistory.YFilter }

func (cvcallvolumestatshistory *CISCOVOICEDIALCONTROLMIB_Cvcallvolumestatshistory) SetFilter(yf yfilter.YFilter) { cvcallvolumestatshistory.YFilter = yf }

func (cvcallvolumestatshistory *CISCOVOICEDIALCONTROLMIB_Cvcallvolumestatshistory) GetGoName(yname string) string {
    if yname == "cvCallDurationStatsThreshold" { return "Cvcalldurationstatsthreshold" }
    if yname == "cvCallVolumeWMTableSize" { return "Cvcallvolumewmtablesize" }
    return ""
}

func (cvcallvolumestatshistory *CISCOVOICEDIALCONTROLMIB_Cvcallvolumestatshistory) GetSegmentPath() string {
    return "cvCallVolumeStatsHistory"
}

func (cvcallvolumestatshistory *CISCOVOICEDIALCONTROLMIB_Cvcallvolumestatshistory) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cvcallvolumestatshistory *CISCOVOICEDIALCONTROLMIB_Cvcallvolumestatshistory) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cvcallvolumestatshistory *CISCOVOICEDIALCONTROLMIB_Cvcallvolumestatshistory) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cvCallDurationStatsThreshold"] = cvcallvolumestatshistory.Cvcalldurationstatsthreshold
    leafs["cvCallVolumeWMTableSize"] = cvcallvolumestatshistory.Cvcallvolumewmtablesize
    return leafs
}

func (cvcallvolumestatshistory *CISCOVOICEDIALCONTROLMIB_Cvcallvolumestatshistory) GetBundleName() string { return "cisco_ios_xe" }

func (cvcallvolumestatshistory *CISCOVOICEDIALCONTROLMIB_Cvcallvolumestatshistory) GetYangName() string { return "cvCallVolumeStatsHistory" }

func (cvcallvolumestatshistory *CISCOVOICEDIALCONTROLMIB_Cvcallvolumestatshistory) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvcallvolumestatshistory *CISCOVOICEDIALCONTROLMIB_Cvcallvolumestatshistory) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvcallvolumestatshistory *CISCOVOICEDIALCONTROLMIB_Cvcallvolumestatshistory) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvcallvolumestatshistory *CISCOVOICEDIALCONTROLMIB_Cvcallvolumestatshistory) SetParent(parent types.Entity) { cvcallvolumestatshistory.parent = parent }

func (cvcallvolumestatshistory *CISCOVOICEDIALCONTROLMIB_Cvcallvolumestatshistory) GetParent() types.Entity { return cvcallvolumestatshistory.parent }

func (cvcallvolumestatshistory *CISCOVOICEDIALCONTROLMIB_Cvcallvolumestatshistory) GetParentYangName() string { return "CISCO-VOICE-DIAL-CONTROL-MIB" }

// CISCOVOICEDIALCONTROLMIB_Cvpeercfgtable
// The table contains the Voice Generic Peer information that
// is used to create an ifIndexed row with an appropriate
// ifType that is associated with the cvPeerCfgType and
// cvPeerCfgPeerType objects.
type CISCOVOICEDIALCONTROLMIB_Cvpeercfgtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A single voice generic Peer. The creation of this entry will create an
    // associated ifEntry with an ifType that is associated with cvPeerCfgType,
    // i.e., for 'voiceEncap' encapsulation, an ifEntry will contain an ifType
    // voiceEncap(103); for 'voiceOverIp' encapsulation, an ifEntry will contain
    // an ifType voiceOverIp(104). The ifAdminStatus of the newly created ifEntry
    // is set to 'up' and ifOperStatus is set to 'down'. In addition, an
    // associated voiceEncap/voiceOverIp Peer configuration entry is created after
    // the successful ifEntry creation. Then ifIndex of the newly created ifEntry
    // must be used by the network manager to create a peer configuration entry of
    // IETF Dial Control MIB (Refer to RFC 2128 section 2.2.3.1 and the
    // description of dialCtlPeerCfgEntry for the detailed information). In
    // summary, the voice dial peer creation steps are as follows: [1] create this
    // entry (voice/data generic peer entry). [2] read the cvPeerCfgIfIndex of
    // this entry for the     ifIndex of newly created voice/data generic peer.
    // [3] create the dialCtlPeerCfgEntry of RFC 2128 with the     indices of
    // dialCtlPeerCfgId and the ifIndex of newly     created voice generic peer. 
    // For each VoIP peer, it uses IP address and UDP port with RTP protocol to
    // transfer voice packet. Therefore, it does not have its lower layer
    // interface. The dialCtlPeerCfgIfType object of IETF Dial Control MIB must
    // set to 'other' and the dialCtlPeerCfgLowerIf must set to '0'.  After the
    // successful creation of peer configuration entry of IETF Dial Control MIB,
    // the dial plan software in managed device will set the ifOperStatus of the
    // newly created voiceEncap/voiceOverIp ifEntry to 'up' for enabling the peer
    // function if the peer configuration is completed. When this entry is
    // deleted, its associated ifEntry, voiceEncap/voiceOverIp specific peer entry
    // and the peer entry of IETF Dial Control MIB are deleted. The type is slice
    // of CISCOVOICEDIALCONTROLMIB_Cvpeercfgtable_Cvpeercfgentry.
    Cvpeercfgentry []CISCOVOICEDIALCONTROLMIB_Cvpeercfgtable_Cvpeercfgentry
}

func (cvpeercfgtable *CISCOVOICEDIALCONTROLMIB_Cvpeercfgtable) GetFilter() yfilter.YFilter { return cvpeercfgtable.YFilter }

func (cvpeercfgtable *CISCOVOICEDIALCONTROLMIB_Cvpeercfgtable) SetFilter(yf yfilter.YFilter) { cvpeercfgtable.YFilter = yf }

func (cvpeercfgtable *CISCOVOICEDIALCONTROLMIB_Cvpeercfgtable) GetGoName(yname string) string {
    if yname == "cvPeerCfgEntry" { return "Cvpeercfgentry" }
    return ""
}

func (cvpeercfgtable *CISCOVOICEDIALCONTROLMIB_Cvpeercfgtable) GetSegmentPath() string {
    return "cvPeerCfgTable"
}

func (cvpeercfgtable *CISCOVOICEDIALCONTROLMIB_Cvpeercfgtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cvPeerCfgEntry" {
        for _, c := range cvpeercfgtable.Cvpeercfgentry {
            if cvpeercfgtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOVOICEDIALCONTROLMIB_Cvpeercfgtable_Cvpeercfgentry{}
        cvpeercfgtable.Cvpeercfgentry = append(cvpeercfgtable.Cvpeercfgentry, child)
        return &cvpeercfgtable.Cvpeercfgentry[len(cvpeercfgtable.Cvpeercfgentry)-1]
    }
    return nil
}

func (cvpeercfgtable *CISCOVOICEDIALCONTROLMIB_Cvpeercfgtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cvpeercfgtable.Cvpeercfgentry {
        children[cvpeercfgtable.Cvpeercfgentry[i].GetSegmentPath()] = &cvpeercfgtable.Cvpeercfgentry[i]
    }
    return children
}

func (cvpeercfgtable *CISCOVOICEDIALCONTROLMIB_Cvpeercfgtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cvpeercfgtable *CISCOVOICEDIALCONTROLMIB_Cvpeercfgtable) GetBundleName() string { return "cisco_ios_xe" }

func (cvpeercfgtable *CISCOVOICEDIALCONTROLMIB_Cvpeercfgtable) GetYangName() string { return "cvPeerCfgTable" }

func (cvpeercfgtable *CISCOVOICEDIALCONTROLMIB_Cvpeercfgtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvpeercfgtable *CISCOVOICEDIALCONTROLMIB_Cvpeercfgtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvpeercfgtable *CISCOVOICEDIALCONTROLMIB_Cvpeercfgtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvpeercfgtable *CISCOVOICEDIALCONTROLMIB_Cvpeercfgtable) SetParent(parent types.Entity) { cvpeercfgtable.parent = parent }

func (cvpeercfgtable *CISCOVOICEDIALCONTROLMIB_Cvpeercfgtable) GetParent() types.Entity { return cvpeercfgtable.parent }

func (cvpeercfgtable *CISCOVOICEDIALCONTROLMIB_Cvpeercfgtable) GetParentYangName() string { return "CISCO-VOICE-DIAL-CONTROL-MIB" }

// CISCOVOICEDIALCONTROLMIB_Cvpeercfgtable_Cvpeercfgentry
// A single voice generic Peer. The creation of this
// entry will create an associated ifEntry with an ifType
// that is associated with cvPeerCfgType, i.e., for
// 'voiceEncap' encapsulation, an ifEntry will contain an
// ifType voiceEncap(103); for 'voiceOverIp' encapsulation,
// an ifEntry will contain an ifType voiceOverIp(104). The
// ifAdminStatus of the newly created ifEntry is set to 'up'
// and ifOperStatus is set to 'down'. In addition, an
// associated voiceEncap/voiceOverIp Peer configuration
// entry is created after the successful ifEntry creation.
// Then ifIndex of the newly created ifEntry must be used by
// the network manager to create a peer configuration entry
// of IETF Dial Control MIB (Refer to RFC 2128 section
// 2.2.3.1 and the description of dialCtlPeerCfgEntry for the
// detailed information).
// In summary, the voice dial peer creation steps are as
// follows:
// [1] create this entry (voice/data generic peer entry).
// [2] read the cvPeerCfgIfIndex of this entry for the
//     ifIndex of newly created voice/data generic peer.
// [3] create the dialCtlPeerCfgEntry of RFC 2128 with the
//     indices of dialCtlPeerCfgId and the ifIndex of newly
//     created voice generic peer.
// 
// For each VoIP peer, it uses IP address and UDP port with
// RTP protocol to transfer voice packet. Therefore, it does
// not have its lower layer interface. The
// dialCtlPeerCfgIfType object of IETF Dial Control MIB must
// set to 'other' and the dialCtlPeerCfgLowerIf must set to
// '0'.
// 
// After the successful creation of peer configuration entry
// of IETF Dial Control MIB, the dial plan software in
// managed device will set the ifOperStatus of the newly
// created voiceEncap/voiceOverIp ifEntry to 'up' for
// enabling the peer function if the peer configuration is
// completed.
// When this entry is deleted, its associated ifEntry,
// voiceEncap/voiceOverIp specific peer entry and the peer
// entry of IETF Dial Control MIB are deleted.
type CISCOVOICEDIALCONTROLMIB_Cvpeercfgtable_Cvpeercfgentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. An arbitrary index that uniquely identifies a
    // generic voice peer. The type is interface{} with range: 1..2147483647.
    Cvpeercfgindex interface{}

    // The ifIndex of the peer associated ifEntry. The ifIndex appears after the
    // associated ifEntry is created successfully. This ifIndex will be used to
    // access the objects in the Voice over Telephony or Voice over IP peer
    // specific table. In addition, the ifIndex is also used to access the
    // associated peer configuration entry of the IETF Dial Control MIB. If the
    // peer associated ifEntry had not been created, then this object has a value
    // of zero. The type is interface{} with range: 0..2147483647.
    Cvpeercfgifindex interface{}

    // Specifies the type of voice related encapsulation. voice - voice
    // encapsulation (voiceEncap ifType) on the         telephony network. voip  -
    // VoIP encapsulation (voiceOverIp ifType) on the IP         network. mmail -
    // Media Mail over IP encapsulation (mediaMailOverIp         ifType) on the IP
    // network. voatm - VoATM encapsulation (voiceOverATM ifType) on the        
    // ATM network. vofr  - VoFR encapsulation (voiceOverFR ifType) on the        
    // Frame Relay network. The type is Cvpeercfgtype.
    Cvpeercfgtype interface{}

    // This object is used to create a new row or modify or delete an existing row
    // in this table. The type is RowStatus.
    Cvpeercfgrowstatus interface{}

    // Specifies the type of a peer. voice - peer in voice type to be defined in a
    // voice         gateway for voice calls.  data  - peer in data type to be
    // defined in gateway         that supports universal ports for modem/data    
    // calls and integrated ports for data calls. The type is Cvpeercfgpeertype.
    Cvpeercfgpeertype interface{}

    // This object represents the total number of active calls that has selected
    // the dialpeer as an incoming dialpeer. The type is interface{} with range:
    // 0..65535.
    Cvcallvolpeerincomingcalls interface{}

    // This object represents the total number of active calls that has selected
    // the dialpeer as an outgoing dialpeer. The type is interface{} with range:
    // 0..65535.
    Cvcallvolpeeroutgoingcalls interface{}
}

func (cvpeercfgentry *CISCOVOICEDIALCONTROLMIB_Cvpeercfgtable_Cvpeercfgentry) GetFilter() yfilter.YFilter { return cvpeercfgentry.YFilter }

func (cvpeercfgentry *CISCOVOICEDIALCONTROLMIB_Cvpeercfgtable_Cvpeercfgentry) SetFilter(yf yfilter.YFilter) { cvpeercfgentry.YFilter = yf }

func (cvpeercfgentry *CISCOVOICEDIALCONTROLMIB_Cvpeercfgtable_Cvpeercfgentry) GetGoName(yname string) string {
    if yname == "cvPeerCfgIndex" { return "Cvpeercfgindex" }
    if yname == "cvPeerCfgIfIndex" { return "Cvpeercfgifindex" }
    if yname == "cvPeerCfgType" { return "Cvpeercfgtype" }
    if yname == "cvPeerCfgRowStatus" { return "Cvpeercfgrowstatus" }
    if yname == "cvPeerCfgPeerType" { return "Cvpeercfgpeertype" }
    if yname == "cvCallVolPeerIncomingCalls" { return "Cvcallvolpeerincomingcalls" }
    if yname == "cvCallVolPeerOutgoingCalls" { return "Cvcallvolpeeroutgoingcalls" }
    return ""
}

func (cvpeercfgentry *CISCOVOICEDIALCONTROLMIB_Cvpeercfgtable_Cvpeercfgentry) GetSegmentPath() string {
    return "cvPeerCfgEntry" + "[cvPeerCfgIndex='" + fmt.Sprintf("%v", cvpeercfgentry.Cvpeercfgindex) + "']"
}

func (cvpeercfgentry *CISCOVOICEDIALCONTROLMIB_Cvpeercfgtable_Cvpeercfgentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cvpeercfgentry *CISCOVOICEDIALCONTROLMIB_Cvpeercfgtable_Cvpeercfgentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cvpeercfgentry *CISCOVOICEDIALCONTROLMIB_Cvpeercfgtable_Cvpeercfgentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cvPeerCfgIndex"] = cvpeercfgentry.Cvpeercfgindex
    leafs["cvPeerCfgIfIndex"] = cvpeercfgentry.Cvpeercfgifindex
    leafs["cvPeerCfgType"] = cvpeercfgentry.Cvpeercfgtype
    leafs["cvPeerCfgRowStatus"] = cvpeercfgentry.Cvpeercfgrowstatus
    leafs["cvPeerCfgPeerType"] = cvpeercfgentry.Cvpeercfgpeertype
    leafs["cvCallVolPeerIncomingCalls"] = cvpeercfgentry.Cvcallvolpeerincomingcalls
    leafs["cvCallVolPeerOutgoingCalls"] = cvpeercfgentry.Cvcallvolpeeroutgoingcalls
    return leafs
}

func (cvpeercfgentry *CISCOVOICEDIALCONTROLMIB_Cvpeercfgtable_Cvpeercfgentry) GetBundleName() string { return "cisco_ios_xe" }

func (cvpeercfgentry *CISCOVOICEDIALCONTROLMIB_Cvpeercfgtable_Cvpeercfgentry) GetYangName() string { return "cvPeerCfgEntry" }

func (cvpeercfgentry *CISCOVOICEDIALCONTROLMIB_Cvpeercfgtable_Cvpeercfgentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvpeercfgentry *CISCOVOICEDIALCONTROLMIB_Cvpeercfgtable_Cvpeercfgentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvpeercfgentry *CISCOVOICEDIALCONTROLMIB_Cvpeercfgtable_Cvpeercfgentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvpeercfgentry *CISCOVOICEDIALCONTROLMIB_Cvpeercfgtable_Cvpeercfgentry) SetParent(parent types.Entity) { cvpeercfgentry.parent = parent }

func (cvpeercfgentry *CISCOVOICEDIALCONTROLMIB_Cvpeercfgtable_Cvpeercfgentry) GetParent() types.Entity { return cvpeercfgentry.parent }

func (cvpeercfgentry *CISCOVOICEDIALCONTROLMIB_Cvpeercfgtable_Cvpeercfgentry) GetParentYangName() string { return "cvPeerCfgTable" }

// CISCOVOICEDIALCONTROLMIB_Cvpeercfgtable_Cvpeercfgentry_Cvpeercfgpeertype represents         calls and integrated ports for data calls.
type CISCOVOICEDIALCONTROLMIB_Cvpeercfgtable_Cvpeercfgentry_Cvpeercfgpeertype string

const (
    CISCOVOICEDIALCONTROLMIB_Cvpeercfgtable_Cvpeercfgentry_Cvpeercfgpeertype_voice CISCOVOICEDIALCONTROLMIB_Cvpeercfgtable_Cvpeercfgentry_Cvpeercfgpeertype = "voice"

    CISCOVOICEDIALCONTROLMIB_Cvpeercfgtable_Cvpeercfgentry_Cvpeercfgpeertype_data CISCOVOICEDIALCONTROLMIB_Cvpeercfgtable_Cvpeercfgentry_Cvpeercfgpeertype = "data"
)

// CISCOVOICEDIALCONTROLMIB_Cvpeercfgtable_Cvpeercfgentry_Cvpeercfgtype represents         Frame Relay network.
type CISCOVOICEDIALCONTROLMIB_Cvpeercfgtable_Cvpeercfgentry_Cvpeercfgtype string

const (
    CISCOVOICEDIALCONTROLMIB_Cvpeercfgtable_Cvpeercfgentry_Cvpeercfgtype_voice CISCOVOICEDIALCONTROLMIB_Cvpeercfgtable_Cvpeercfgentry_Cvpeercfgtype = "voice"

    CISCOVOICEDIALCONTROLMIB_Cvpeercfgtable_Cvpeercfgentry_Cvpeercfgtype_voip CISCOVOICEDIALCONTROLMIB_Cvpeercfgtable_Cvpeercfgentry_Cvpeercfgtype = "voip"

    CISCOVOICEDIALCONTROLMIB_Cvpeercfgtable_Cvpeercfgentry_Cvpeercfgtype_mmail CISCOVOICEDIALCONTROLMIB_Cvpeercfgtable_Cvpeercfgentry_Cvpeercfgtype = "mmail"

    CISCOVOICEDIALCONTROLMIB_Cvpeercfgtable_Cvpeercfgentry_Cvpeercfgtype_voatm CISCOVOICEDIALCONTROLMIB_Cvpeercfgtable_Cvpeercfgentry_Cvpeercfgtype = "voatm"

    CISCOVOICEDIALCONTROLMIB_Cvpeercfgtable_Cvpeercfgentry_Cvpeercfgtype_vofr CISCOVOICEDIALCONTROLMIB_Cvpeercfgtable_Cvpeercfgentry_Cvpeercfgtype = "vofr"
)

// CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable
// The table contains the Voice over Telephony peer specific
// information that is required to accept voice calls or to
// which it will place them or perform various loopback tests
// via interface.
type CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A single Voice specific Peer. One entry per voice encapsulation. The entry
    // is created when its associated 'voiceEncap(103)' encapsulation ifEntry is
    // created. This entry is deleted when its associated ifEntry is deleted. The
    // type is slice of
    // CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable_Cvvoicepeercfgentry.
    Cvvoicepeercfgentry []CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable_Cvvoicepeercfgentry
}

func (cvvoicepeercfgtable *CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable) GetFilter() yfilter.YFilter { return cvvoicepeercfgtable.YFilter }

func (cvvoicepeercfgtable *CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable) SetFilter(yf yfilter.YFilter) { cvvoicepeercfgtable.YFilter = yf }

func (cvvoicepeercfgtable *CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable) GetGoName(yname string) string {
    if yname == "cvVoicePeerCfgEntry" { return "Cvvoicepeercfgentry" }
    return ""
}

func (cvvoicepeercfgtable *CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable) GetSegmentPath() string {
    return "cvVoicePeerCfgTable"
}

func (cvvoicepeercfgtable *CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cvVoicePeerCfgEntry" {
        for _, c := range cvvoicepeercfgtable.Cvvoicepeercfgentry {
            if cvvoicepeercfgtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable_Cvvoicepeercfgentry{}
        cvvoicepeercfgtable.Cvvoicepeercfgentry = append(cvvoicepeercfgtable.Cvvoicepeercfgentry, child)
        return &cvvoicepeercfgtable.Cvvoicepeercfgentry[len(cvvoicepeercfgtable.Cvvoicepeercfgentry)-1]
    }
    return nil
}

func (cvvoicepeercfgtable *CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cvvoicepeercfgtable.Cvvoicepeercfgentry {
        children[cvvoicepeercfgtable.Cvvoicepeercfgentry[i].GetSegmentPath()] = &cvvoicepeercfgtable.Cvvoicepeercfgentry[i]
    }
    return children
}

func (cvvoicepeercfgtable *CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cvvoicepeercfgtable *CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable) GetBundleName() string { return "cisco_ios_xe" }

func (cvvoicepeercfgtable *CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable) GetYangName() string { return "cvVoicePeerCfgTable" }

func (cvvoicepeercfgtable *CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvvoicepeercfgtable *CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvvoicepeercfgtable *CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvvoicepeercfgtable *CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable) SetParent(parent types.Entity) { cvvoicepeercfgtable.parent = parent }

func (cvvoicepeercfgtable *CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable) GetParent() types.Entity { return cvvoicepeercfgtable.parent }

func (cvvoicepeercfgtable *CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable) GetParentYangName() string { return "CISCO-VOICE-DIAL-CONTROL-MIB" }

// CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable_Cvvoicepeercfgentry
// A single Voice specific Peer. One entry per voice
// encapsulation.
// The entry is created when its associated 'voiceEncap(103)'
// encapsulation ifEntry is created.
// This entry is deleted when its associated ifEntry is
// deleted.
type CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable_Cvvoicepeercfgentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // The object specifies the session target of the peer. Session Targets
    // definitions: The session target has the syntax used by the IETF service
    // location protocol. The syntax is as follows: 
    // mapping-type:type-specific-syntax  the mapping-type specifies a scheme for
    // mapping the matching dial string to a session target.  The valid Mapping
    // type definitions for the peer are as follows: loopback - Syntax:
    // loopback:where    'where' string is defined as follows:    compressed -
    // loopback is performed on compressed voice                 as close to the
    // CODEC which processes the                 data as possible.    uncompressed
    // - loopback is performed on the PCM or                 analog voice as close
    // to the telephony                 endpoint as possible.  Local loopback
    // case: uncompressed - the PCM voice coming into the DSP is simply     turned
    // around and sent back out, allowing testing of     the transmit--> receive
    // paths in the telephony     endpoint. compressed - the compressed voice
    // coming out of the CODEC is     turned around on the DSP module and fed back
    // into the     decompressor through the jitter buffer. In addition to     the
    // telephony endpoint, this tests both the encode and     decode paths without
    // involving data paths or packet     handling on the host router.  Remote
    // loopback case: compressed - RTP packets received from the network are    
    // decapsulated and passed to the DSP board. Instead of     feeding these into
    // the CODEC for decompression, they     are immediately sent back to the
    // session application     as if they had originated locally and been
    // compressed. uncompressed - In addition to the above, the voice samples    
    // are sent all the way through the CODEC and then turned     around instead
    // of being sent to the telephony     endpoint. The type is string with
    // length: 0..64.
    Cvvoicepeercfgsessiontarget interface{}

    // The object specifies the prefix of the dial digits for the peer. The dial
    // digits prefix is sent to telephony interface before the real phone number
    // when the system places the outgoing call to the voice encapsulation peer
    // over telephony interface. The type is string with length: 0..32.
    Cvvoicepeercfgdialdigitsprefix interface{}

    // The object enables/disables the DID call treatment for incoming DNIS
    // digits. true  - treat the incoming DNIS digits as if the digits         are
    // received from DID trunk. false - Disable DID call treatment for incoming
    // DNIS         digits. The type is bool.
    Cvvoicepeercfgdidcallenable interface{}

    // The object specifies the CAS group number of a CAS capable T1/E1  that is
    // in the dialCtlPeerCfgLowerIf object of RFC2128. This object can be set to a
    // valid CAS group number only if the dialCtlPeerCfgLowerIf contains a valid
    // ifIndex for a CAS capable T1/E1. The object must set to -1 before the value
    // of the  dialCtlPeerCfgLowerIf object of RFC2128 can be changed. The type is
    // interface{} with range: -1..30.
    Cvvoicepeercfgcasgroup interface{}

    // This object specifies that the E.164 number specified in the
    // dialCtlPeerCfgOriginateAddress field of the associated dialCtlPeerCfgTable
    // entry be registered as an extension  phone number of this gateway for H323
    // gatekeeper and/or  SIP registrar. The type is bool.
    Cvvoicepeercfgregistere164 interface{}

    // This object specifies the number of dialed digits to forward to the remote
    // destination in the setup message. The object can take the value:     0-32
    // number of right justified digits to forward     -1 default     -2 forward
    // extra digits i.e those over and above        those needed to match to the
    // destination pattern     -3 forward all digits. The type is interface{} with
    // range: -3..32.
    Cvvoicepeercfgforwarddigits interface{}

    // This object specifies which, if any, test to run in the echo canceller when
    // a call from the network is connected. echoCancellerTestNone    - do not run
    // a test. echoCancellerG168Test2A  - run ITU-T G.168 Test 2A.
    // echoCancellerG168Test2B  - run ITU-T G.168 Test 2B.
    // echoCancellerG168Test2Ca - run ITU-T G.168 Test 2C(a).
    // echoCancellerG168Test2Cb - run ITU-T G.168 Test 2C(b).
    // echoCancellerG168Test3A  - run ITU-T G.168 Test 3A. echoCancellerG168Test3B
    // - run ITU-T G.168 Test 3B. echoCancellerG168Test3C  - run ITU-T G.168 Test
    // 3C. echoCancellerG168Test4   - run ITU-T G.168 Test 4.
    // echoCancellerG168Test5   - run ITU-T G.168 Test 5. echoCancellerG168Test6  
    // - run ITU-T G.168 Test 6. echoCancellerG168Test7   - run ITU-T G.168 Test
    // 7. echoCancellerG168Test9   - run ITU-T G.168 Test 9. The type is
    // Cvvoicepeercfgechocancellertest.
    Cvvoicepeercfgechocancellertest interface{}
}

func (cvvoicepeercfgentry *CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable_Cvvoicepeercfgentry) GetFilter() yfilter.YFilter { return cvvoicepeercfgentry.YFilter }

func (cvvoicepeercfgentry *CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable_Cvvoicepeercfgentry) SetFilter(yf yfilter.YFilter) { cvvoicepeercfgentry.YFilter = yf }

func (cvvoicepeercfgentry *CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable_Cvvoicepeercfgentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "cvVoicePeerCfgSessionTarget" { return "Cvvoicepeercfgsessiontarget" }
    if yname == "cvVoicePeerCfgDialDigitsPrefix" { return "Cvvoicepeercfgdialdigitsprefix" }
    if yname == "cvVoicePeerCfgDIDCallEnable" { return "Cvvoicepeercfgdidcallenable" }
    if yname == "cvVoicePeerCfgCasGroup" { return "Cvvoicepeercfgcasgroup" }
    if yname == "cvVoicePeerCfgRegisterE164" { return "Cvvoicepeercfgregistere164" }
    if yname == "cvVoicePeerCfgForwardDigits" { return "Cvvoicepeercfgforwarddigits" }
    if yname == "cvVoicePeerCfgEchoCancellerTest" { return "Cvvoicepeercfgechocancellertest" }
    return ""
}

func (cvvoicepeercfgentry *CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable_Cvvoicepeercfgentry) GetSegmentPath() string {
    return "cvVoicePeerCfgEntry" + "[ifIndex='" + fmt.Sprintf("%v", cvvoicepeercfgentry.Ifindex) + "']"
}

func (cvvoicepeercfgentry *CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable_Cvvoicepeercfgentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cvvoicepeercfgentry *CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable_Cvvoicepeercfgentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cvvoicepeercfgentry *CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable_Cvvoicepeercfgentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = cvvoicepeercfgentry.Ifindex
    leafs["cvVoicePeerCfgSessionTarget"] = cvvoicepeercfgentry.Cvvoicepeercfgsessiontarget
    leafs["cvVoicePeerCfgDialDigitsPrefix"] = cvvoicepeercfgentry.Cvvoicepeercfgdialdigitsprefix
    leafs["cvVoicePeerCfgDIDCallEnable"] = cvvoicepeercfgentry.Cvvoicepeercfgdidcallenable
    leafs["cvVoicePeerCfgCasGroup"] = cvvoicepeercfgentry.Cvvoicepeercfgcasgroup
    leafs["cvVoicePeerCfgRegisterE164"] = cvvoicepeercfgentry.Cvvoicepeercfgregistere164
    leafs["cvVoicePeerCfgForwardDigits"] = cvvoicepeercfgentry.Cvvoicepeercfgforwarddigits
    leafs["cvVoicePeerCfgEchoCancellerTest"] = cvvoicepeercfgentry.Cvvoicepeercfgechocancellertest
    return leafs
}

func (cvvoicepeercfgentry *CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable_Cvvoicepeercfgentry) GetBundleName() string { return "cisco_ios_xe" }

func (cvvoicepeercfgentry *CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable_Cvvoicepeercfgentry) GetYangName() string { return "cvVoicePeerCfgEntry" }

func (cvvoicepeercfgentry *CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable_Cvvoicepeercfgentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvvoicepeercfgentry *CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable_Cvvoicepeercfgentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvvoicepeercfgentry *CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable_Cvvoicepeercfgentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvvoicepeercfgentry *CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable_Cvvoicepeercfgentry) SetParent(parent types.Entity) { cvvoicepeercfgentry.parent = parent }

func (cvvoicepeercfgentry *CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable_Cvvoicepeercfgentry) GetParent() types.Entity { return cvvoicepeercfgentry.parent }

func (cvvoicepeercfgentry *CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable_Cvvoicepeercfgentry) GetParentYangName() string { return "cvVoicePeerCfgTable" }

// CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable_Cvvoicepeercfgentry_Cvvoicepeercfgechocancellertest represents echoCancellerG168Test9   - run ITU-T G.168 Test 9.
type CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable_Cvvoicepeercfgentry_Cvvoicepeercfgechocancellertest string

const (
    CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable_Cvvoicepeercfgentry_Cvvoicepeercfgechocancellertest_echoCancellerTestNone CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable_Cvvoicepeercfgentry_Cvvoicepeercfgechocancellertest = "echoCancellerTestNone"

    CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable_Cvvoicepeercfgentry_Cvvoicepeercfgechocancellertest_echoCancellerG168Test2A CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable_Cvvoicepeercfgentry_Cvvoicepeercfgechocancellertest = "echoCancellerG168Test2A"

    CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable_Cvvoicepeercfgentry_Cvvoicepeercfgechocancellertest_echoCancellerG168Test2B CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable_Cvvoicepeercfgentry_Cvvoicepeercfgechocancellertest = "echoCancellerG168Test2B"

    CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable_Cvvoicepeercfgentry_Cvvoicepeercfgechocancellertest_echoCancellerG168Test2Ca CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable_Cvvoicepeercfgentry_Cvvoicepeercfgechocancellertest = "echoCancellerG168Test2Ca"

    CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable_Cvvoicepeercfgentry_Cvvoicepeercfgechocancellertest_echoCancellerG168Test2Cb CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable_Cvvoicepeercfgentry_Cvvoicepeercfgechocancellertest = "echoCancellerG168Test2Cb"

    CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable_Cvvoicepeercfgentry_Cvvoicepeercfgechocancellertest_echoCancellerG168Test3A CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable_Cvvoicepeercfgentry_Cvvoicepeercfgechocancellertest = "echoCancellerG168Test3A"

    CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable_Cvvoicepeercfgentry_Cvvoicepeercfgechocancellertest_echoCancellerG168Test3B CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable_Cvvoicepeercfgentry_Cvvoicepeercfgechocancellertest = "echoCancellerG168Test3B"

    CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable_Cvvoicepeercfgentry_Cvvoicepeercfgechocancellertest_echoCancellerG168Test3C CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable_Cvvoicepeercfgentry_Cvvoicepeercfgechocancellertest = "echoCancellerG168Test3C"

    CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable_Cvvoicepeercfgentry_Cvvoicepeercfgechocancellertest_echoCancellerG168Test4 CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable_Cvvoicepeercfgentry_Cvvoicepeercfgechocancellertest = "echoCancellerG168Test4"

    CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable_Cvvoicepeercfgentry_Cvvoicepeercfgechocancellertest_echoCancellerG168Test6 CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable_Cvvoicepeercfgentry_Cvvoicepeercfgechocancellertest = "echoCancellerG168Test6"

    CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable_Cvvoicepeercfgentry_Cvvoicepeercfgechocancellertest_echoCancellerG168Test9 CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable_Cvvoicepeercfgentry_Cvvoicepeercfgechocancellertest = "echoCancellerG168Test9"

    CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable_Cvvoicepeercfgentry_Cvvoicepeercfgechocancellertest_echoCancellerG168Test5 CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable_Cvvoicepeercfgentry_Cvvoicepeercfgechocancellertest = "echoCancellerG168Test5"

    CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable_Cvvoicepeercfgentry_Cvvoicepeercfgechocancellertest_echoCancellerG168Test7 CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable_Cvvoicepeercfgentry_Cvvoicepeercfgechocancellertest = "echoCancellerG168Test7"
)

// CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable
// The table contains the Voice over IP (VoIP) peer specific
// information that is required to accept voice calls or to
// which it will place them via IP backbone with the
// specified session protocol in cvVoIPPeerCfgSessionProtocol.
type CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A single VoIP specific Peer. One entry per VoIP encapsulation. The entry is
    // created when its associated 'voiceOverIp(104)' encapsulation ifEntry is
    // created. This entry is deleted when its associated ifEntry is deleted. The
    // type is slice of
    // CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable_Cvvoippeercfgentry.
    Cvvoippeercfgentry []CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable_Cvvoippeercfgentry
}

func (cvvoippeercfgtable *CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable) GetFilter() yfilter.YFilter { return cvvoippeercfgtable.YFilter }

func (cvvoippeercfgtable *CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable) SetFilter(yf yfilter.YFilter) { cvvoippeercfgtable.YFilter = yf }

func (cvvoippeercfgtable *CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable) GetGoName(yname string) string {
    if yname == "cvVoIPPeerCfgEntry" { return "Cvvoippeercfgentry" }
    return ""
}

func (cvvoippeercfgtable *CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable) GetSegmentPath() string {
    return "cvVoIPPeerCfgTable"
}

func (cvvoippeercfgtable *CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cvVoIPPeerCfgEntry" {
        for _, c := range cvvoippeercfgtable.Cvvoippeercfgentry {
            if cvvoippeercfgtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable_Cvvoippeercfgentry{}
        cvvoippeercfgtable.Cvvoippeercfgentry = append(cvvoippeercfgtable.Cvvoippeercfgentry, child)
        return &cvvoippeercfgtable.Cvvoippeercfgentry[len(cvvoippeercfgtable.Cvvoippeercfgentry)-1]
    }
    return nil
}

func (cvvoippeercfgtable *CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cvvoippeercfgtable.Cvvoippeercfgentry {
        children[cvvoippeercfgtable.Cvvoippeercfgentry[i].GetSegmentPath()] = &cvvoippeercfgtable.Cvvoippeercfgentry[i]
    }
    return children
}

func (cvvoippeercfgtable *CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cvvoippeercfgtable *CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable) GetBundleName() string { return "cisco_ios_xe" }

func (cvvoippeercfgtable *CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable) GetYangName() string { return "cvVoIPPeerCfgTable" }

func (cvvoippeercfgtable *CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvvoippeercfgtable *CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvvoippeercfgtable *CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvvoippeercfgtable *CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable) SetParent(parent types.Entity) { cvvoippeercfgtable.parent = parent }

func (cvvoippeercfgtable *CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable) GetParent() types.Entity { return cvvoippeercfgtable.parent }

func (cvvoippeercfgtable *CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable) GetParentYangName() string { return "CISCO-VOICE-DIAL-CONTROL-MIB" }

// CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable_Cvvoippeercfgentry
// A single VoIP specific Peer. One entry per VoIP
// encapsulation.
// The entry is created when its associated 'voiceOverIp(104)'
// encapsulation ifEntry is created.
// This entry is deleted when its associated ifEntry is
// deleted.
type CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable_Cvvoippeercfgentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // The object specifies the session protocol to be used for Internet call
    // between local and remote router via IP backbone. The type is
    // CvSessionProtocol.
    Cvvoippeercfgsessionprotocol interface{}

    // The object specifies the user requested Quality of Service for the call.
    // The type is QosService.
    Cvvoippeercfgdesiredqos interface{}

    // The object specifies the minimally acceptable Quality of Service for the
    // call. The type is QosService.
    Cvvoippeercfgminacceptableqos interface{}

    // The object specifies the session target of the peer. Session Targets
    // definitions: The session target has the syntax used by the IETF service
    // location protocol. The syntax is as follows: 
    // mapping-type:type-specific-syntax  the mapping-type specifies a scheme for
    // mapping the matching dial string to a session target. The
    // type-specific-syntax is exactly that, something that the particular mapping
    // scheme can understand. For example, Session target           Meaning
    // ipv4:171.68.13.55:1006   The session target is the IP                      
    // version 4 address of 171.68.13.55                          and port 1006.
    // dns:pots.cisco.com:1661  The session target is the IP host                 
    // with dns name pots.cisco.com, and                          port 1661. ras  
    // The session target is the                          gatekeeper with RAS
    // (Registration                          , Admission,  Status protocol).
    // settlement               The session target is the                         
    // settlement server. enum:1                   The session target is the enum 
    // profile match table 1.  The valid Mapping type definitions for the peer are
    // as follows: ipv4       - Syntax: ipv4:w.x.y.z:port or  ipv4:w.x.y.z dns    
    // - Syntax: dns:host.domain:port or                      dns:host.domain ras 
    // - Syntax: ras settlement - Syntax: settlement enum       - Syntax: enum: 
    // loopback - Syntax: loopback:where    'where' string is defined as follows: 
    // rtp - loopback is performed at the transport protocol          level. 
    // Local loopback case: rtp - the session application sets up an RTP stream to
    // itself (i.e. actually allocates a port pair and opens     the appropriate
    // UDP sockets). It then does the full     RTP encapsulation, sends the
    // packets to the loopback     IP address, receives the RTP packets, and hands
    // the     compressed voice back to the CODEC. This tests the     entire local
    // processing path, both transmit and     receive, in the router, as well as
    // all of the above     paths.  Remote loopback case: rtp: RTP packets
    // received from the network are decapsulated and      immediately
    // re-encapsulated in the outbound RTP      stream, using the same media clock
    // (i.e. timestamp)      as the received packet. They are then sent back to   
    // the remote source router as if the voice had      originated on a telephony
    // port on the local router. The type is string.
    Cvvoippeercfgsessiontarget interface{}

    // This object specifies the most desirable codec of speech for the VoIP peer.
    // The type is CvcSpeechCoderRate.
    Cvvoippeercfgcoderrate interface{}

    // This object specifies the default transmit rate of FAX the VoIP peer. If
    // the value of this object is 'none', then the Fax relay feature is disabled;
    // otherwise the Fax relay feature is enabled. The type is CvcFaxTransmitRate.
    Cvvoippeercfgfaxrate interface{}

    // This object specifies whether or not the VAD (Voice Activity Detection)
    // voice data is continuously transmitted to IP backbone. The type is bool.
    Cvvoippeercfgvadenable interface{}

    // This object specifies the user requested Expectation Factor of voice
    // quality for the call via this peer. The type is interface{} with range:
    // 0..20. Units are equipment impairment factor (eif).
    Cvvoippeercfgexpectfactor interface{}

    // This object specifies the user requested Calculated Planning Impairment
    // Factor (Icpif) for the call via this peer. The type is interface{} with
    // range: 0..55. Units are equipment impairment factor (eif).
    Cvvoippeercfgicpif interface{}

    // This object specifies whether cvdcPoorQoVNotification (or the newer
    // cvdcPoorQoVNotificationRev1) traps should be generated for the call that is
    // associated with this peer. The type is bool.
    Cvvoippeercfgpoorqovnotificationenable interface{}

    // This object specifies whether the outgoing voice related UDP packet
    // contains a valid checksum value. true  - enable the checksum of outgoing
    // voice UDP packets false - disable the checksum of outgoing voice UDP
    // packets. The type is bool.
    Cvvoippeercfgudpchecksumenable interface{}

    // This object specifies the value to be stored in the IP Precedence field of
    // voice packets, with values ranging from 0 (normal precedence) through 7
    // (network control), allowing the managed system to set the importance of
    // each voice packet for delivering them to the destination peer. The type is
    // interface{} with range: 0..7.
    Cvvoippeercfgipprecedence interface{}

    // This object specifies the technology prefix of the peer, The technology
    // prefix and the called party address are passed in Admission Request (ARQ)
    // to gatekeeper for the called party address resolution during call setup.
    // The type is string with length: 0..128.
    Cvvoippeercfgtechprefix interface{}

    // This object specifies the methods to transmit dial digits (DTMF or MF
    // digits) via IP network. rtpCisco       - Enable capability to transmit dial
    // digits                  with Cisco proprietary RTP payload type. h245Signal
    // - Enable capability to transmit dtmf digits                  across the
    // H.245 channel, via the signal                  field of the
    // UserInputIndication message h245Alphanumeric - Enable capability to
    // transmit dtmf                  digit across the H.245 channel, via the     
    // string or alphanumeric fields of the                  UserInputIndication
    // message rtpNte         - Enable capability to transmit dial digits         
    // using Named Telephony Event per RFC 2833                  section 3.
    // sipNotify      - Enable capability to transmit dtmf                  digits
    // using unsolicited SIP NOTIFY                  messages. This mechanism is
    // only available                  for SIP dialpeers. sipKpml        - Enable
    // capability to transmit dtmf                  digits using KPML over SIP
    // SUBSCRIBE                  and NOTIFY messages. This mechanism is          
    // only available for SIP dialpeers.   Modifying the value of
    // cvVoIPPeerCfgSessionProtocol can reset the digit-relay method associated
    // bits value in this object if the modified session protocol does not support
    // these digit-relay methods. The type is map[string]bool.
    Cvvoippeercfgdigitrelay interface{}

    // This object specifies the size of the voice payload sample to be produced
    // by the coder specified in cvVoIPPeerCfgCoderRate. Each coder sample
    // produces 10 bytes of voice payload. The specified value will be rounded
    // down to the nearest valid size.  A value of 0, specifies that the coder
    // defined by cvVoIPPeerCfgCoderRate should produce its default payload size.
    // The type is interface{} with range: 0..None | 10..240. Units are bytes.
    Cvvoippeercfgcoderbytes interface{}

    // This object specifies the payload size to be produced by the coder when it
    // is generating fax data. A value of 0, specifies that the coder specified in
    // cvVoIPCfgPeerCoderRate should produce its default fax payload size. The
    // type is interface{} with range: 0..None | 10..255. Units are bytes.
    Cvvoippeercfgfaxbytes interface{}

    // This object specifies the type of in-band signaling that will be used
    // between the end points of the call. It is used by the router to determine
    // how to interpret ABCD signaling bits sent as part of voice payload data.
    // The type is CvcInBandSignaling.
    Cvvoippeercfginbandsignaling interface{}

    // This object specifies how the media is to be setup on an IP-IP Gateway. Two
    // choices are valid: flow-through and flow-around. When in flow-through mode,
    // which is the default setting, the IP-IP Gateway will terminate and  then
    // re-originate the media stream. When flow-around is configured the Gateway
    // will not be involved with the media, since it will flow-around the Gateway
    // and will be established directly between the endpoints. The type is
    // Cvvoippeercfgmediasetting.
    Cvvoippeercfgmediasetting interface{}

    // The object specifies the user requested Quality of Service for the video
    // portion of the call. The type is QosService.
    Cvvoippeercfgdesiredqosvideo interface{}

    // The object specifies the minimally acceptable Quality of Service for the
    // video portion of the call. The type is QosService.
    Cvvoippeercfgminacceptableqosvideo interface{}

    // This object specifies the Inbound VoIP calls that match an outbound VoIP
    // dialpeer will be responded with a SIP  redirect(for inbound SIP) or H.450.3
    // call-forward(for  inbound H.323). The type is bool.
    Cvvoippeercfgredirectip2Ip interface{}

    // If the object has a value true(1) octet align operation is used, and if the
    // value is false(2), bandwidth efficient operation is used. This object is
    // not instantiated when the object cvVoIPPeerCfgCoderRate is not equal to
    // gsmAmrNb enum. The type is bool.
    Cvvoippeercfgoctetaligned interface{}

    // This object indicates modes of Bit rates. One or more upto four modes can
    // be configured at the same time as bit rates can be changed dynamically for
    // AMR-NB codec. This object is not instantiated when the object
    // cvVoIPPeerCfgCoderRate is not equal to gsmAmrNb enum. The type is
    // map[string]bool.
    Cvvoippeercfgbitrates interface{}

    // If the object has a value of true(1), frame CRC will be included in the
    // payload and if the value is false(2), frame CRC will not be included in the
    // payload. This object is applicable only when RTP frame type is octet
    // aligned. This object is not instantiated when the object
    // cvVoIPPeerCfgCoderRate is not equal to gsmAmrNb enum. The type is bool.
    Cvvoippeercfgcrc interface{}

    // This object indicates the iLBC codec mode to be used. The value of this
    // object is valid only if  cvVoIPPeerCfgCoderRate is equal to 'iLBC'. The
    // type is CvIlbcFrameMode.
    Cvvoippeercfgcodermode interface{}

    // This object specifies the coding mode to be used. The object is
    // instantiated only if cvVoIPPeerCfgCoderRate is 'iSAC'. Following coding
    // modes are supported: adaptive    (1) - adaptive mode where iSAC performs
    // bandwidth                     estimation and adapts to the available
    // channel                    bandwidth. independent (2) - independent mode in
    // which no bandwidth estimation                    is performed. The type is
    // Cvvoippeercfgcodingmode.
    Cvvoippeercfgcodingmode interface{}

    // This object specifies the target bit rate. The object is instantiated only
    // if cvVoIPPeerCfgCoderRate is 'iSAC'. The type is interface{} with range:
    // 10000..32000.
    Cvvoippeercfgbitrate interface{}

    // This object specifies the frame size used. The object is instantiated only
    // if cvVoIPPeerCfgCoderRate is 'iSAC'. The frame size can be 30 ms or 60 ms,
    // and it can be fixed for all packets or vary depending on the configuration
    // and bandwidth estimation. Thus it can have the following values:
    // frameSize30      - initial frame size of 30 ms frameSize60      - initial
    // frame size of 60 ms frameSize30fixed - fixed frame size 30 ms
    // frameSize60fixed - fixed frame size 60 ms. The type is
    // Cvvoippeercfgframesize.
    Cvvoippeercfgframesize interface{}

    // This object specifies whether cvdcPolicyViolationNotification traps should
    // be generated for the call that is associated with this peer for RPH to DSCP
    // mapping and policing feature. The type is bool.
    Cvvoippeercfgdscppolicynotificationenable interface{}

    // This object specifies whether cvdcPolicyViolationNotification traps should
    // be generated for the call that is associated with this peer for Media
    // policing feature.. The type is bool.
    Cvvoippeercfgmediapolicynotificationenable interface{}
}

func (cvvoippeercfgentry *CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable_Cvvoippeercfgentry) GetFilter() yfilter.YFilter { return cvvoippeercfgentry.YFilter }

func (cvvoippeercfgentry *CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable_Cvvoippeercfgentry) SetFilter(yf yfilter.YFilter) { cvvoippeercfgentry.YFilter = yf }

func (cvvoippeercfgentry *CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable_Cvvoippeercfgentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "cvVoIPPeerCfgSessionProtocol" { return "Cvvoippeercfgsessionprotocol" }
    if yname == "cvVoIPPeerCfgDesiredQoS" { return "Cvvoippeercfgdesiredqos" }
    if yname == "cvVoIPPeerCfgMinAcceptableQoS" { return "Cvvoippeercfgminacceptableqos" }
    if yname == "cvVoIPPeerCfgSessionTarget" { return "Cvvoippeercfgsessiontarget" }
    if yname == "cvVoIPPeerCfgCoderRate" { return "Cvvoippeercfgcoderrate" }
    if yname == "cvVoIPPeerCfgFaxRate" { return "Cvvoippeercfgfaxrate" }
    if yname == "cvVoIPPeerCfgVADEnable" { return "Cvvoippeercfgvadenable" }
    if yname == "cvVoIPPeerCfgExpectFactor" { return "Cvvoippeercfgexpectfactor" }
    if yname == "cvVoIPPeerCfgIcpif" { return "Cvvoippeercfgicpif" }
    if yname == "cvVoIPPeerCfgPoorQoVNotificationEnable" { return "Cvvoippeercfgpoorqovnotificationenable" }
    if yname == "cvVoIPPeerCfgUDPChecksumEnable" { return "Cvvoippeercfgudpchecksumenable" }
    if yname == "cvVoIPPeerCfgIPPrecedence" { return "Cvvoippeercfgipprecedence" }
    if yname == "cvVoIPPeerCfgTechPrefix" { return "Cvvoippeercfgtechprefix" }
    if yname == "cvVoIPPeerCfgDigitRelay" { return "Cvvoippeercfgdigitrelay" }
    if yname == "cvVoIPPeerCfgCoderBytes" { return "Cvvoippeercfgcoderbytes" }
    if yname == "cvVoIPPeerCfgFaxBytes" { return "Cvvoippeercfgfaxbytes" }
    if yname == "cvVoIPPeerCfgInBandSignaling" { return "Cvvoippeercfginbandsignaling" }
    if yname == "cvVoIPPeerCfgMediaSetting" { return "Cvvoippeercfgmediasetting" }
    if yname == "cvVoIPPeerCfgDesiredQoSVideo" { return "Cvvoippeercfgdesiredqosvideo" }
    if yname == "cvVoIPPeerCfgMinAcceptableQoSVideo" { return "Cvvoippeercfgminacceptableqosvideo" }
    if yname == "cvVoIPPeerCfgRedirectip2ip" { return "Cvvoippeercfgredirectip2Ip" }
    if yname == "cvVoIPPeerCfgOctetAligned" { return "Cvvoippeercfgoctetaligned" }
    if yname == "cvVoIPPeerCfgBitRates" { return "Cvvoippeercfgbitrates" }
    if yname == "cvVoIPPeerCfgCRC" { return "Cvvoippeercfgcrc" }
    if yname == "cvVoIPPeerCfgCoderMode" { return "Cvvoippeercfgcodermode" }
    if yname == "cvVoIPPeerCfgCodingMode" { return "Cvvoippeercfgcodingmode" }
    if yname == "cvVoIPPeerCfgBitRate" { return "Cvvoippeercfgbitrate" }
    if yname == "cvVoIPPeerCfgFrameSize" { return "Cvvoippeercfgframesize" }
    if yname == "cvVoIPPeerCfgDSCPPolicyNotificationEnable" { return "Cvvoippeercfgdscppolicynotificationenable" }
    if yname == "cvVoIPPeerCfgMediaPolicyNotificationEnable" { return "Cvvoippeercfgmediapolicynotificationenable" }
    return ""
}

func (cvvoippeercfgentry *CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable_Cvvoippeercfgentry) GetSegmentPath() string {
    return "cvVoIPPeerCfgEntry" + "[ifIndex='" + fmt.Sprintf("%v", cvvoippeercfgentry.Ifindex) + "']"
}

func (cvvoippeercfgentry *CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable_Cvvoippeercfgentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cvvoippeercfgentry *CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable_Cvvoippeercfgentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cvvoippeercfgentry *CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable_Cvvoippeercfgentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = cvvoippeercfgentry.Ifindex
    leafs["cvVoIPPeerCfgSessionProtocol"] = cvvoippeercfgentry.Cvvoippeercfgsessionprotocol
    leafs["cvVoIPPeerCfgDesiredQoS"] = cvvoippeercfgentry.Cvvoippeercfgdesiredqos
    leafs["cvVoIPPeerCfgMinAcceptableQoS"] = cvvoippeercfgentry.Cvvoippeercfgminacceptableqos
    leafs["cvVoIPPeerCfgSessionTarget"] = cvvoippeercfgentry.Cvvoippeercfgsessiontarget
    leafs["cvVoIPPeerCfgCoderRate"] = cvvoippeercfgentry.Cvvoippeercfgcoderrate
    leafs["cvVoIPPeerCfgFaxRate"] = cvvoippeercfgentry.Cvvoippeercfgfaxrate
    leafs["cvVoIPPeerCfgVADEnable"] = cvvoippeercfgentry.Cvvoippeercfgvadenable
    leafs["cvVoIPPeerCfgExpectFactor"] = cvvoippeercfgentry.Cvvoippeercfgexpectfactor
    leafs["cvVoIPPeerCfgIcpif"] = cvvoippeercfgentry.Cvvoippeercfgicpif
    leafs["cvVoIPPeerCfgPoorQoVNotificationEnable"] = cvvoippeercfgentry.Cvvoippeercfgpoorqovnotificationenable
    leafs["cvVoIPPeerCfgUDPChecksumEnable"] = cvvoippeercfgentry.Cvvoippeercfgudpchecksumenable
    leafs["cvVoIPPeerCfgIPPrecedence"] = cvvoippeercfgentry.Cvvoippeercfgipprecedence
    leafs["cvVoIPPeerCfgTechPrefix"] = cvvoippeercfgentry.Cvvoippeercfgtechprefix
    leafs["cvVoIPPeerCfgDigitRelay"] = cvvoippeercfgentry.Cvvoippeercfgdigitrelay
    leafs["cvVoIPPeerCfgCoderBytes"] = cvvoippeercfgentry.Cvvoippeercfgcoderbytes
    leafs["cvVoIPPeerCfgFaxBytes"] = cvvoippeercfgentry.Cvvoippeercfgfaxbytes
    leafs["cvVoIPPeerCfgInBandSignaling"] = cvvoippeercfgentry.Cvvoippeercfginbandsignaling
    leafs["cvVoIPPeerCfgMediaSetting"] = cvvoippeercfgentry.Cvvoippeercfgmediasetting
    leafs["cvVoIPPeerCfgDesiredQoSVideo"] = cvvoippeercfgentry.Cvvoippeercfgdesiredqosvideo
    leafs["cvVoIPPeerCfgMinAcceptableQoSVideo"] = cvvoippeercfgentry.Cvvoippeercfgminacceptableqosvideo
    leafs["cvVoIPPeerCfgRedirectip2ip"] = cvvoippeercfgentry.Cvvoippeercfgredirectip2Ip
    leafs["cvVoIPPeerCfgOctetAligned"] = cvvoippeercfgentry.Cvvoippeercfgoctetaligned
    leafs["cvVoIPPeerCfgBitRates"] = cvvoippeercfgentry.Cvvoippeercfgbitrates
    leafs["cvVoIPPeerCfgCRC"] = cvvoippeercfgentry.Cvvoippeercfgcrc
    leafs["cvVoIPPeerCfgCoderMode"] = cvvoippeercfgentry.Cvvoippeercfgcodermode
    leafs["cvVoIPPeerCfgCodingMode"] = cvvoippeercfgentry.Cvvoippeercfgcodingmode
    leafs["cvVoIPPeerCfgBitRate"] = cvvoippeercfgentry.Cvvoippeercfgbitrate
    leafs["cvVoIPPeerCfgFrameSize"] = cvvoippeercfgentry.Cvvoippeercfgframesize
    leafs["cvVoIPPeerCfgDSCPPolicyNotificationEnable"] = cvvoippeercfgentry.Cvvoippeercfgdscppolicynotificationenable
    leafs["cvVoIPPeerCfgMediaPolicyNotificationEnable"] = cvvoippeercfgentry.Cvvoippeercfgmediapolicynotificationenable
    return leafs
}

func (cvvoippeercfgentry *CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable_Cvvoippeercfgentry) GetBundleName() string { return "cisco_ios_xe" }

func (cvvoippeercfgentry *CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable_Cvvoippeercfgentry) GetYangName() string { return "cvVoIPPeerCfgEntry" }

func (cvvoippeercfgentry *CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable_Cvvoippeercfgentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvvoippeercfgentry *CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable_Cvvoippeercfgentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvvoippeercfgentry *CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable_Cvvoippeercfgentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvvoippeercfgentry *CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable_Cvvoippeercfgentry) SetParent(parent types.Entity) { cvvoippeercfgentry.parent = parent }

func (cvvoippeercfgentry *CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable_Cvvoippeercfgentry) GetParent() types.Entity { return cvvoippeercfgentry.parent }

func (cvvoippeercfgentry *CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable_Cvvoippeercfgentry) GetParentYangName() string { return "cvVoIPPeerCfgTable" }

// CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable_Cvvoippeercfgentry_Cvvoippeercfgcodingmode represents                   is performed.
type CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable_Cvvoippeercfgentry_Cvvoippeercfgcodingmode string

const (
    CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable_Cvvoippeercfgentry_Cvvoippeercfgcodingmode_adaptive CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable_Cvvoippeercfgentry_Cvvoippeercfgcodingmode = "adaptive"

    CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable_Cvvoippeercfgentry_Cvvoippeercfgcodingmode_independent CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable_Cvvoippeercfgentry_Cvvoippeercfgcodingmode = "independent"
)

// CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable_Cvvoippeercfgentry_Cvvoippeercfgframesize represents frameSize60fixed - fixed frame size 60 ms
type CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable_Cvvoippeercfgentry_Cvvoippeercfgframesize string

const (
    CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable_Cvvoippeercfgentry_Cvvoippeercfgframesize_frameSize30 CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable_Cvvoippeercfgentry_Cvvoippeercfgframesize = "frameSize30"

    CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable_Cvvoippeercfgentry_Cvvoippeercfgframesize_frameSize60 CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable_Cvvoippeercfgentry_Cvvoippeercfgframesize = "frameSize60"

    CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable_Cvvoippeercfgentry_Cvvoippeercfgframesize_frameSize30fixed CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable_Cvvoippeercfgentry_Cvvoippeercfgframesize = "frameSize30fixed"

    CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable_Cvvoippeercfgentry_Cvvoippeercfgframesize_frameSize60fixed CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable_Cvvoippeercfgentry_Cvvoippeercfgframesize = "frameSize60fixed"
)

// CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable_Cvvoippeercfgentry_Cvvoippeercfgmediasetting represents be established directly between the endpoints.
type CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable_Cvvoippeercfgentry_Cvvoippeercfgmediasetting string

const (
    CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable_Cvvoippeercfgentry_Cvvoippeercfgmediasetting_flowThrough CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable_Cvvoippeercfgentry_Cvvoippeercfgmediasetting = "flowThrough"

    CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable_Cvvoippeercfgentry_Cvvoippeercfgmediasetting_flowAround CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable_Cvvoippeercfgentry_Cvvoippeercfgmediasetting = "flowAround"
)

// CISCOVOICEDIALCONTROLMIB_Cvpeercommoncfgtable
// The table contains the Voice specific peer common
// configuration information that is required to accept voice
// calls or to which it will place them or process the
// incoming calls.
type CISCOVOICEDIALCONTROLMIB_Cvpeercommoncfgtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A single Voice specific Peer. One entry per voice related encapsulation.
    // The entry is created when a voice related encapsulation ifEntry is created.
    // This entry is deleted when its associated ifEntry is deleted. The type is
    // slice of
    // CISCOVOICEDIALCONTROLMIB_Cvpeercommoncfgtable_Cvpeercommoncfgentry.
    Cvpeercommoncfgentry []CISCOVOICEDIALCONTROLMIB_Cvpeercommoncfgtable_Cvpeercommoncfgentry
}

func (cvpeercommoncfgtable *CISCOVOICEDIALCONTROLMIB_Cvpeercommoncfgtable) GetFilter() yfilter.YFilter { return cvpeercommoncfgtable.YFilter }

func (cvpeercommoncfgtable *CISCOVOICEDIALCONTROLMIB_Cvpeercommoncfgtable) SetFilter(yf yfilter.YFilter) { cvpeercommoncfgtable.YFilter = yf }

func (cvpeercommoncfgtable *CISCOVOICEDIALCONTROLMIB_Cvpeercommoncfgtable) GetGoName(yname string) string {
    if yname == "cvPeerCommonCfgEntry" { return "Cvpeercommoncfgentry" }
    return ""
}

func (cvpeercommoncfgtable *CISCOVOICEDIALCONTROLMIB_Cvpeercommoncfgtable) GetSegmentPath() string {
    return "cvPeerCommonCfgTable"
}

func (cvpeercommoncfgtable *CISCOVOICEDIALCONTROLMIB_Cvpeercommoncfgtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cvPeerCommonCfgEntry" {
        for _, c := range cvpeercommoncfgtable.Cvpeercommoncfgentry {
            if cvpeercommoncfgtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOVOICEDIALCONTROLMIB_Cvpeercommoncfgtable_Cvpeercommoncfgentry{}
        cvpeercommoncfgtable.Cvpeercommoncfgentry = append(cvpeercommoncfgtable.Cvpeercommoncfgentry, child)
        return &cvpeercommoncfgtable.Cvpeercommoncfgentry[len(cvpeercommoncfgtable.Cvpeercommoncfgentry)-1]
    }
    return nil
}

func (cvpeercommoncfgtable *CISCOVOICEDIALCONTROLMIB_Cvpeercommoncfgtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cvpeercommoncfgtable.Cvpeercommoncfgentry {
        children[cvpeercommoncfgtable.Cvpeercommoncfgentry[i].GetSegmentPath()] = &cvpeercommoncfgtable.Cvpeercommoncfgentry[i]
    }
    return children
}

func (cvpeercommoncfgtable *CISCOVOICEDIALCONTROLMIB_Cvpeercommoncfgtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cvpeercommoncfgtable *CISCOVOICEDIALCONTROLMIB_Cvpeercommoncfgtable) GetBundleName() string { return "cisco_ios_xe" }

func (cvpeercommoncfgtable *CISCOVOICEDIALCONTROLMIB_Cvpeercommoncfgtable) GetYangName() string { return "cvPeerCommonCfgTable" }

func (cvpeercommoncfgtable *CISCOVOICEDIALCONTROLMIB_Cvpeercommoncfgtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvpeercommoncfgtable *CISCOVOICEDIALCONTROLMIB_Cvpeercommoncfgtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvpeercommoncfgtable *CISCOVOICEDIALCONTROLMIB_Cvpeercommoncfgtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvpeercommoncfgtable *CISCOVOICEDIALCONTROLMIB_Cvpeercommoncfgtable) SetParent(parent types.Entity) { cvpeercommoncfgtable.parent = parent }

func (cvpeercommoncfgtable *CISCOVOICEDIALCONTROLMIB_Cvpeercommoncfgtable) GetParent() types.Entity { return cvpeercommoncfgtable.parent }

func (cvpeercommoncfgtable *CISCOVOICEDIALCONTROLMIB_Cvpeercommoncfgtable) GetParentYangName() string { return "CISCO-VOICE-DIAL-CONTROL-MIB" }

// CISCOVOICEDIALCONTROLMIB_Cvpeercommoncfgtable_Cvpeercommoncfgentry
// A single Voice specific Peer. One entry per voice related
// encapsulation.
// The entry is created when a voice related encapsulation
// ifEntry is created.
// This entry is deleted when its associated ifEntry is
// deleted.
type CISCOVOICEDIALCONTROLMIB_Cvpeercommoncfgtable_Cvpeercommoncfgentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // The object specifies the prefix of the incoming Dialed Number
    // Identification Service (DNIS) digits for the peer. The DNIS digits prefix
    // is used to match with the incoming DNIS number for incoming call
    // discrimination. If the digits in this object are matched with incoming DNIS
    // number, the  associated dialCtlPeerCfgInfoType in RFC 2128 will be used as
    // a call discriminator for differentiating speech, data, fax, video or modem
    // calls. The type is string with length: 0..32.
    Cvpeercommoncfgincomingdnisdigits interface{}

    // The object specifies the maximum allowed connection to/from the peer. A
    // value of -1 disables the limit of maximum connections. The type is
    // interface{} with range: -1..None | 1..2147483647. Units are connections.
    Cvpeercommoncfgmaxconnections interface{}

    // The object specifies the application to handle the incoming call after the
    // peer is selected. If no application name is specified, then the default
    // session application will take care of the incoming call. The type is
    // string.
    Cvpeercommoncfgapplicationname interface{}

    // This object specifies the selection preference of a peer when multiple
    // peers are matched to the selection criteria. The value of 0 has the lowest
    // preference for peer selection. The type is interface{} with range: 0..10.
    Cvpeercommoncfgpreference interface{}

    // This object specifies whether dialpeer hunting should stop when this peer
    // is reached. The type is bool.
    Cvpeercommoncfghuntstop interface{}

    // The object specifies a Dialer Number Identification Service (DNIS) map name
    // for the Voice specific peer entry specified in this row. A DNIS is a called
    // party number and they can be grouped and identified by DNIS map. The type
    // is string.
    Cvpeercommoncfgdnismappingname interface{}

    // The object specifies the Source Carrier Id for the peer. The Source Carrier
    // Id is used to match with the Source Carrier Id of a call. If the Source
    // Carrier Id in this object is matched with the Source Carrier Id of a call,
    // then the associated peer will be used to handle the call.  Only
    // alphanumeric characters, '-' and '_' are allowed in the string. The type is
    // string.
    Cvpeercommoncfgsourcecarrierid interface{}

    // The object specifies the Target Carrier Id for the peer. The Target Carrier
    // Id is used to match with the Target Carrier Id of a call. If the Target
    // Carrier Id in this object is matched with the Target Carrier Id of a call,
    // then the associated peer will be used to handle the call. Only alphanumeric
    // characters, '-' and '_' are allowed in the string. The type is string.
    Cvpeercommoncfgtargetcarrierid interface{}

    // The object specifies the Source Trunk Group Label for the peer. The Source
    // Trunk Group Label is used to match with the Source Trunk Group Label of a
    // call. If the Source Trunk Group Label in this object is matched with the
    // Source Trunk Group Label of a call, then the associated peer will be used
    // to handle the call.  Only alphanumeric characters, '-' and '_' are allowed
    // in the string. The type is string.
    Cvpeercommoncfgsourcetrunkgrplabel interface{}

    // The object specifies the Target Trunk Group Label for the peer. The Target
    // Trunk Group Label is used to match with the Target Trunk Group Label of a
    // call. If the Target Trunk Group Label in this object is matched with the
    // Target Trunk Group Label of a call, then the associated peer will be used
    // to handle the call. Only alphanumeric characters, '-' and '_' are allowed
    // in the string. The type is string.
    Cvpeercommoncfgtargettrunkgrplabel interface{}
}

func (cvpeercommoncfgentry *CISCOVOICEDIALCONTROLMIB_Cvpeercommoncfgtable_Cvpeercommoncfgentry) GetFilter() yfilter.YFilter { return cvpeercommoncfgentry.YFilter }

func (cvpeercommoncfgentry *CISCOVOICEDIALCONTROLMIB_Cvpeercommoncfgtable_Cvpeercommoncfgentry) SetFilter(yf yfilter.YFilter) { cvpeercommoncfgentry.YFilter = yf }

func (cvpeercommoncfgentry *CISCOVOICEDIALCONTROLMIB_Cvpeercommoncfgtable_Cvpeercommoncfgentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "cvPeerCommonCfgIncomingDnisDigits" { return "Cvpeercommoncfgincomingdnisdigits" }
    if yname == "cvPeerCommonCfgMaxConnections" { return "Cvpeercommoncfgmaxconnections" }
    if yname == "cvPeerCommonCfgApplicationName" { return "Cvpeercommoncfgapplicationname" }
    if yname == "cvPeerCommonCfgPreference" { return "Cvpeercommoncfgpreference" }
    if yname == "cvPeerCommonCfgHuntStop" { return "Cvpeercommoncfghuntstop" }
    if yname == "cvPeerCommonCfgDnisMappingName" { return "Cvpeercommoncfgdnismappingname" }
    if yname == "cvPeerCommonCfgSourceCarrierId" { return "Cvpeercommoncfgsourcecarrierid" }
    if yname == "cvPeerCommonCfgTargetCarrierId" { return "Cvpeercommoncfgtargetcarrierid" }
    if yname == "cvPeerCommonCfgSourceTrunkGrpLabel" { return "Cvpeercommoncfgsourcetrunkgrplabel" }
    if yname == "cvPeerCommonCfgTargetTrunkGrpLabel" { return "Cvpeercommoncfgtargettrunkgrplabel" }
    return ""
}

func (cvpeercommoncfgentry *CISCOVOICEDIALCONTROLMIB_Cvpeercommoncfgtable_Cvpeercommoncfgentry) GetSegmentPath() string {
    return "cvPeerCommonCfgEntry" + "[ifIndex='" + fmt.Sprintf("%v", cvpeercommoncfgentry.Ifindex) + "']"
}

func (cvpeercommoncfgentry *CISCOVOICEDIALCONTROLMIB_Cvpeercommoncfgtable_Cvpeercommoncfgentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cvpeercommoncfgentry *CISCOVOICEDIALCONTROLMIB_Cvpeercommoncfgtable_Cvpeercommoncfgentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cvpeercommoncfgentry *CISCOVOICEDIALCONTROLMIB_Cvpeercommoncfgtable_Cvpeercommoncfgentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = cvpeercommoncfgentry.Ifindex
    leafs["cvPeerCommonCfgIncomingDnisDigits"] = cvpeercommoncfgentry.Cvpeercommoncfgincomingdnisdigits
    leafs["cvPeerCommonCfgMaxConnections"] = cvpeercommoncfgentry.Cvpeercommoncfgmaxconnections
    leafs["cvPeerCommonCfgApplicationName"] = cvpeercommoncfgentry.Cvpeercommoncfgapplicationname
    leafs["cvPeerCommonCfgPreference"] = cvpeercommoncfgentry.Cvpeercommoncfgpreference
    leafs["cvPeerCommonCfgHuntStop"] = cvpeercommoncfgentry.Cvpeercommoncfghuntstop
    leafs["cvPeerCommonCfgDnisMappingName"] = cvpeercommoncfgentry.Cvpeercommoncfgdnismappingname
    leafs["cvPeerCommonCfgSourceCarrierId"] = cvpeercommoncfgentry.Cvpeercommoncfgsourcecarrierid
    leafs["cvPeerCommonCfgTargetCarrierId"] = cvpeercommoncfgentry.Cvpeercommoncfgtargetcarrierid
    leafs["cvPeerCommonCfgSourceTrunkGrpLabel"] = cvpeercommoncfgentry.Cvpeercommoncfgsourcetrunkgrplabel
    leafs["cvPeerCommonCfgTargetTrunkGrpLabel"] = cvpeercommoncfgentry.Cvpeercommoncfgtargettrunkgrplabel
    return leafs
}

func (cvpeercommoncfgentry *CISCOVOICEDIALCONTROLMIB_Cvpeercommoncfgtable_Cvpeercommoncfgentry) GetBundleName() string { return "cisco_ios_xe" }

func (cvpeercommoncfgentry *CISCOVOICEDIALCONTROLMIB_Cvpeercommoncfgtable_Cvpeercommoncfgentry) GetYangName() string { return "cvPeerCommonCfgEntry" }

func (cvpeercommoncfgentry *CISCOVOICEDIALCONTROLMIB_Cvpeercommoncfgtable_Cvpeercommoncfgentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvpeercommoncfgentry *CISCOVOICEDIALCONTROLMIB_Cvpeercommoncfgtable_Cvpeercommoncfgentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvpeercommoncfgentry *CISCOVOICEDIALCONTROLMIB_Cvpeercommoncfgtable_Cvpeercommoncfgentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvpeercommoncfgentry *CISCOVOICEDIALCONTROLMIB_Cvpeercommoncfgtable_Cvpeercommoncfgentry) SetParent(parent types.Entity) { cvpeercommoncfgentry.parent = parent }

func (cvpeercommoncfgentry *CISCOVOICEDIALCONTROLMIB_Cvpeercommoncfgtable_Cvpeercommoncfgentry) GetParent() types.Entity { return cvpeercommoncfgentry.parent }

func (cvpeercommoncfgentry *CISCOVOICEDIALCONTROLMIB_Cvpeercommoncfgtable_Cvpeercommoncfgentry) GetParentYangName() string { return "cvPeerCommonCfgTable" }

// CISCOVOICEDIALCONTROLMIB_Cvcallactivetable
// This table is the voice extension to the call active table
// of IETF Dial Control MIB. It contains voice encapsulation
// call leg information that is derived from the statistics
// of lower layer telephony interface.
type CISCOVOICEDIALCONTROLMIB_Cvcallactivetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The information regarding a single voice encapsulation call leg. The call
    // leg entry is identified by using the same index objects that are used by
    // Call Active table of IETF Dial Control MIB to identify the call. An entry
    // of this table is created when its associated call active entry in the IETF
    // Dial Control MIB is created and call active entry contains the call
    // establishment to a voice over telephony network peer. The entry is deleted
    // when its associated call active entry in the IETF Dial Control MIB is
    // deleted. The type is slice of
    // CISCOVOICEDIALCONTROLMIB_Cvcallactivetable_Cvcallactiveentry.
    Cvcallactiveentry []CISCOVOICEDIALCONTROLMIB_Cvcallactivetable_Cvcallactiveentry
}

func (cvcallactivetable *CISCOVOICEDIALCONTROLMIB_Cvcallactivetable) GetFilter() yfilter.YFilter { return cvcallactivetable.YFilter }

func (cvcallactivetable *CISCOVOICEDIALCONTROLMIB_Cvcallactivetable) SetFilter(yf yfilter.YFilter) { cvcallactivetable.YFilter = yf }

func (cvcallactivetable *CISCOVOICEDIALCONTROLMIB_Cvcallactivetable) GetGoName(yname string) string {
    if yname == "cvCallActiveEntry" { return "Cvcallactiveentry" }
    return ""
}

func (cvcallactivetable *CISCOVOICEDIALCONTROLMIB_Cvcallactivetable) GetSegmentPath() string {
    return "cvCallActiveTable"
}

func (cvcallactivetable *CISCOVOICEDIALCONTROLMIB_Cvcallactivetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cvCallActiveEntry" {
        for _, c := range cvcallactivetable.Cvcallactiveentry {
            if cvcallactivetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOVOICEDIALCONTROLMIB_Cvcallactivetable_Cvcallactiveentry{}
        cvcallactivetable.Cvcallactiveentry = append(cvcallactivetable.Cvcallactiveentry, child)
        return &cvcallactivetable.Cvcallactiveentry[len(cvcallactivetable.Cvcallactiveentry)-1]
    }
    return nil
}

func (cvcallactivetable *CISCOVOICEDIALCONTROLMIB_Cvcallactivetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cvcallactivetable.Cvcallactiveentry {
        children[cvcallactivetable.Cvcallactiveentry[i].GetSegmentPath()] = &cvcallactivetable.Cvcallactiveentry[i]
    }
    return children
}

func (cvcallactivetable *CISCOVOICEDIALCONTROLMIB_Cvcallactivetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cvcallactivetable *CISCOVOICEDIALCONTROLMIB_Cvcallactivetable) GetBundleName() string { return "cisco_ios_xe" }

func (cvcallactivetable *CISCOVOICEDIALCONTROLMIB_Cvcallactivetable) GetYangName() string { return "cvCallActiveTable" }

func (cvcallactivetable *CISCOVOICEDIALCONTROLMIB_Cvcallactivetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvcallactivetable *CISCOVOICEDIALCONTROLMIB_Cvcallactivetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvcallactivetable *CISCOVOICEDIALCONTROLMIB_Cvcallactivetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvcallactivetable *CISCOVOICEDIALCONTROLMIB_Cvcallactivetable) SetParent(parent types.Entity) { cvcallactivetable.parent = parent }

func (cvcallactivetable *CISCOVOICEDIALCONTROLMIB_Cvcallactivetable) GetParent() types.Entity { return cvcallactivetable.parent }

func (cvcallactivetable *CISCOVOICEDIALCONTROLMIB_Cvcallactivetable) GetParentYangName() string { return "CISCO-VOICE-DIAL-CONTROL-MIB" }

// CISCOVOICEDIALCONTROLMIB_Cvcallactivetable_Cvcallactiveentry
// The information regarding a single voice encapsulation
// call leg.
// The call leg entry is identified by using the same index
// objects that are used by Call Active table of IETF Dial
// Control MIB to identify the call.
// An entry of this table is created when its associated call
// active entry in the IETF Dial Control MIB is created and
// call active entry contains the call establishment to a
// voice over telephony network peer.
// The entry is deleted when its associated call active entry
// in the IETF Dial Control MIB is deleted.
type CISCOVOICEDIALCONTROLMIB_Cvcallactivetable_Cvcallactiveentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // dial_control_mib.DIALCONTROLMIB_Callactivetable_Callactiveentry_Callactivesetuptime
    Callactivesetuptime interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // dial_control_mib.DIALCONTROLMIB_Callactivetable_Callactiveentry_Callactiveindex
    Callactiveindex interface{}

    // The global connection identifier for the active telephony leg of the call.
    // The type is string with length: 0..16.
    Cvcallactiveconnectionid interface{}

    // Duration of Transmit path open from this peer to the voice gateway for the
    // call leg. This counter object will lock at the maximum value which is
    // approximately two days. The type is interface{} with range: 0..4294967295.
    // Units are milliseconds.
    Cvcallactivetxduration interface{}

    // Duration of voice transmitted from this peer to voice gateway for this call
    // leg. The Voice Utilization Rate can be obtained by dividing this by
    // cvCallActiveTXDuration object. This counter object will lock at the maximum
    // value which is approximately two days. The type is interface{} with range:
    // 0..4294967295. Units are milliseconds.
    Cvcallactivevoicetxduration interface{}

    // Duration of fax transmitted from this peer to voice gateway for this call
    // leg. The FAX Utilization Rate can be obtained by dividing this by
    // cvCallActiveTXDuration object. This counter object will lock at the maximum
    // value which is approximately two days. The type is interface{} with range:
    // 0..4294967295. Units are milliseconds.
    Cvcallactivefaxtxduration interface{}

    // The negotiated coder rate. It specifies the transmit rate of voice/fax
    // compression to its associated call leg for the call. The type is
    // CvcCoderTypeRate.
    Cvcallactivecodertyperate interface{}

    // The object contains the active noise level for the call leg. The type is
    // interface{} with range: -128..8. Units are dBm.
    Cvcallactivenoiselevel interface{}

    // The object contains the sum of Echo Return Loss (ERL), cancellation loss
    // (Echo Return Loss Enhancement) and nonlinear processing loss for the call
    // leg. The value -1 indicates the level can not be determined or level
    // detection is disabled. The type is interface{} with range: -1..127. Units
    // are dB.
    Cvcallactiveacomlevel interface{}

    // The object contains the active output signal level to telephony interface
    // that is used by the call leg. The type is interface{} with range: -128..8.
    // Units are dBm.
    Cvcallactiveoutsignallevel interface{}

    // The object contains the active input signal level from telephony interface
    // that is used by the call leg. The type is interface{} with range: -128..8.
    // Units are dBm.
    Cvcallactiveinsignallevel interface{}

    // The object contains the current Echo Return Loss (ERL) level for the call
    // leg. The value -1 indicates the level can not be determined or level
    // detection is disabled. The type is interface{} with range: -1..45. Units
    // are dB.
    Cvcallactiveerllevel interface{}

    // The object specifies the session target of the peer that is used for the
    // call leg. This object is set with the information in the call associated
    // cvVoicePeerCfgSessionTarget object when the call is connected. The type is
    // string with length: 0..64.
    Cvcallactivesessiontarget interface{}

    // The number of FAX related image pages are received or transmitted via the
    // peer for the call leg. The type is interface{} with range: 0..4294967295.
    // Units are pages.
    Cvcallactiveimgpagecount interface{}

    // The calling party name of the call. If the name is not available, then it
    // will have a length of zero. The type is string.
    Cvcallactivecallingname interface{}

    // The object indicates whether or not the caller ID feature is blocked for
    // this call. The type is bool.
    Cvcallactivecalleridblock interface{}

    // The location in milliseconds of the largest amplitude reflector detected by
    // the echo canceller for this call.  The value 0 indicates there is no
    // reflector or the  information is not available. The type is interface{}
    // with range: 0..128.
    Cvcallactiveecanreflectorlocation interface{}

    // The object indicates the account code input to the call. It could be used
    // for call screen or by down stream server for billing purpose. The value of
    // empty string indicates no account code input. The type is string with
    // length: 0..50.
    Cvcallactiveaccountcode interface{}

    // The object contains the current Echo Return Loss (ERL) level for the call
    // leg. The value -1 indicates the level can not be determined or level
    // detection is disabled. The type is interface{} with range: -1..200. Units
    // are dB.
    Cvcallactiveerllevelrev1 interface{}

    // This object represents the call identifier for the active telephony leg of
    // the call. The type is interface{} with range: 1..4294967295.
    Cvcallactivecallid interface{}
}

func (cvcallactiveentry *CISCOVOICEDIALCONTROLMIB_Cvcallactivetable_Cvcallactiveentry) GetFilter() yfilter.YFilter { return cvcallactiveentry.YFilter }

func (cvcallactiveentry *CISCOVOICEDIALCONTROLMIB_Cvcallactivetable_Cvcallactiveentry) SetFilter(yf yfilter.YFilter) { cvcallactiveentry.YFilter = yf }

func (cvcallactiveentry *CISCOVOICEDIALCONTROLMIB_Cvcallactivetable_Cvcallactiveentry) GetGoName(yname string) string {
    if yname == "callActiveSetupTime" { return "Callactivesetuptime" }
    if yname == "callActiveIndex" { return "Callactiveindex" }
    if yname == "cvCallActiveConnectionId" { return "Cvcallactiveconnectionid" }
    if yname == "cvCallActiveTxDuration" { return "Cvcallactivetxduration" }
    if yname == "cvCallActiveVoiceTxDuration" { return "Cvcallactivevoicetxduration" }
    if yname == "cvCallActiveFaxTxDuration" { return "Cvcallactivefaxtxduration" }
    if yname == "cvCallActiveCoderTypeRate" { return "Cvcallactivecodertyperate" }
    if yname == "cvCallActiveNoiseLevel" { return "Cvcallactivenoiselevel" }
    if yname == "cvCallActiveACOMLevel" { return "Cvcallactiveacomlevel" }
    if yname == "cvCallActiveOutSignalLevel" { return "Cvcallactiveoutsignallevel" }
    if yname == "cvCallActiveInSignalLevel" { return "Cvcallactiveinsignallevel" }
    if yname == "cvCallActiveERLLevel" { return "Cvcallactiveerllevel" }
    if yname == "cvCallActiveSessionTarget" { return "Cvcallactivesessiontarget" }
    if yname == "cvCallActiveImgPageCount" { return "Cvcallactiveimgpagecount" }
    if yname == "cvCallActiveCallingName" { return "Cvcallactivecallingname" }
    if yname == "cvCallActiveCallerIDBlock" { return "Cvcallactivecalleridblock" }
    if yname == "cvCallActiveEcanReflectorLocation" { return "Cvcallactiveecanreflectorlocation" }
    if yname == "cvCallActiveAccountCode" { return "Cvcallactiveaccountcode" }
    if yname == "cvCallActiveERLLevelRev1" { return "Cvcallactiveerllevelrev1" }
    if yname == "cvCallActiveCallId" { return "Cvcallactivecallid" }
    return ""
}

func (cvcallactiveentry *CISCOVOICEDIALCONTROLMIB_Cvcallactivetable_Cvcallactiveentry) GetSegmentPath() string {
    return "cvCallActiveEntry" + "[callActiveSetupTime='" + fmt.Sprintf("%v", cvcallactiveentry.Callactivesetuptime) + "']" + "[callActiveIndex='" + fmt.Sprintf("%v", cvcallactiveentry.Callactiveindex) + "']"
}

func (cvcallactiveentry *CISCOVOICEDIALCONTROLMIB_Cvcallactivetable_Cvcallactiveentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cvcallactiveentry *CISCOVOICEDIALCONTROLMIB_Cvcallactivetable_Cvcallactiveentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cvcallactiveentry *CISCOVOICEDIALCONTROLMIB_Cvcallactivetable_Cvcallactiveentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["callActiveSetupTime"] = cvcallactiveentry.Callactivesetuptime
    leafs["callActiveIndex"] = cvcallactiveentry.Callactiveindex
    leafs["cvCallActiveConnectionId"] = cvcallactiveentry.Cvcallactiveconnectionid
    leafs["cvCallActiveTxDuration"] = cvcallactiveentry.Cvcallactivetxduration
    leafs["cvCallActiveVoiceTxDuration"] = cvcallactiveentry.Cvcallactivevoicetxduration
    leafs["cvCallActiveFaxTxDuration"] = cvcallactiveentry.Cvcallactivefaxtxduration
    leafs["cvCallActiveCoderTypeRate"] = cvcallactiveentry.Cvcallactivecodertyperate
    leafs["cvCallActiveNoiseLevel"] = cvcallactiveentry.Cvcallactivenoiselevel
    leafs["cvCallActiveACOMLevel"] = cvcallactiveentry.Cvcallactiveacomlevel
    leafs["cvCallActiveOutSignalLevel"] = cvcallactiveentry.Cvcallactiveoutsignallevel
    leafs["cvCallActiveInSignalLevel"] = cvcallactiveentry.Cvcallactiveinsignallevel
    leafs["cvCallActiveERLLevel"] = cvcallactiveentry.Cvcallactiveerllevel
    leafs["cvCallActiveSessionTarget"] = cvcallactiveentry.Cvcallactivesessiontarget
    leafs["cvCallActiveImgPageCount"] = cvcallactiveentry.Cvcallactiveimgpagecount
    leafs["cvCallActiveCallingName"] = cvcallactiveentry.Cvcallactivecallingname
    leafs["cvCallActiveCallerIDBlock"] = cvcallactiveentry.Cvcallactivecalleridblock
    leafs["cvCallActiveEcanReflectorLocation"] = cvcallactiveentry.Cvcallactiveecanreflectorlocation
    leafs["cvCallActiveAccountCode"] = cvcallactiveentry.Cvcallactiveaccountcode
    leafs["cvCallActiveERLLevelRev1"] = cvcallactiveentry.Cvcallactiveerllevelrev1
    leafs["cvCallActiveCallId"] = cvcallactiveentry.Cvcallactivecallid
    return leafs
}

func (cvcallactiveentry *CISCOVOICEDIALCONTROLMIB_Cvcallactivetable_Cvcallactiveentry) GetBundleName() string { return "cisco_ios_xe" }

func (cvcallactiveentry *CISCOVOICEDIALCONTROLMIB_Cvcallactivetable_Cvcallactiveentry) GetYangName() string { return "cvCallActiveEntry" }

func (cvcallactiveentry *CISCOVOICEDIALCONTROLMIB_Cvcallactivetable_Cvcallactiveentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvcallactiveentry *CISCOVOICEDIALCONTROLMIB_Cvcallactivetable_Cvcallactiveentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvcallactiveentry *CISCOVOICEDIALCONTROLMIB_Cvcallactivetable_Cvcallactiveentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvcallactiveentry *CISCOVOICEDIALCONTROLMIB_Cvcallactivetable_Cvcallactiveentry) SetParent(parent types.Entity) { cvcallactiveentry.parent = parent }

func (cvcallactiveentry *CISCOVOICEDIALCONTROLMIB_Cvcallactivetable_Cvcallactiveentry) GetParent() types.Entity { return cvcallactiveentry.parent }

func (cvcallactiveentry *CISCOVOICEDIALCONTROLMIB_Cvcallactivetable_Cvcallactiveentry) GetParentYangName() string { return "cvCallActiveTable" }

// CISCOVOICEDIALCONTROLMIB_Cvvoipcallactivetable
// This table is the VoIP extension to the call active table of
// IETF Dial Control MIB. It contains VoIP call leg
// information about specific VoIP call destination and the
// selected QoS for the call leg.
type CISCOVOICEDIALCONTROLMIB_Cvvoipcallactivetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The information regarding a single VoIP call leg. The call leg entry is
    // identified by using the same index objects that are used by Call Active
    // table of IETF Dial Control MIB to identify the call. An entry of this table
    // is created when its associated call active entry in the IETF Dial Control
    // MIB is created and the call active entry contains information for the call
    // establishment to the peer on the IP backbone via a voice over  IP peer. The
    // entry is deleted when its associated call active entry in the IETF Dial
    // Control MIB is deleted. The type is slice of
    // CISCOVOICEDIALCONTROLMIB_Cvvoipcallactivetable_Cvvoipcallactiveentry.
    Cvvoipcallactiveentry []CISCOVOICEDIALCONTROLMIB_Cvvoipcallactivetable_Cvvoipcallactiveentry
}

func (cvvoipcallactivetable *CISCOVOICEDIALCONTROLMIB_Cvvoipcallactivetable) GetFilter() yfilter.YFilter { return cvvoipcallactivetable.YFilter }

func (cvvoipcallactivetable *CISCOVOICEDIALCONTROLMIB_Cvvoipcallactivetable) SetFilter(yf yfilter.YFilter) { cvvoipcallactivetable.YFilter = yf }

func (cvvoipcallactivetable *CISCOVOICEDIALCONTROLMIB_Cvvoipcallactivetable) GetGoName(yname string) string {
    if yname == "cvVoIPCallActiveEntry" { return "Cvvoipcallactiveentry" }
    return ""
}

func (cvvoipcallactivetable *CISCOVOICEDIALCONTROLMIB_Cvvoipcallactivetable) GetSegmentPath() string {
    return "cvVoIPCallActiveTable"
}

func (cvvoipcallactivetable *CISCOVOICEDIALCONTROLMIB_Cvvoipcallactivetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cvVoIPCallActiveEntry" {
        for _, c := range cvvoipcallactivetable.Cvvoipcallactiveentry {
            if cvvoipcallactivetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOVOICEDIALCONTROLMIB_Cvvoipcallactivetable_Cvvoipcallactiveentry{}
        cvvoipcallactivetable.Cvvoipcallactiveentry = append(cvvoipcallactivetable.Cvvoipcallactiveentry, child)
        return &cvvoipcallactivetable.Cvvoipcallactiveentry[len(cvvoipcallactivetable.Cvvoipcallactiveentry)-1]
    }
    return nil
}

func (cvvoipcallactivetable *CISCOVOICEDIALCONTROLMIB_Cvvoipcallactivetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cvvoipcallactivetable.Cvvoipcallactiveentry {
        children[cvvoipcallactivetable.Cvvoipcallactiveentry[i].GetSegmentPath()] = &cvvoipcallactivetable.Cvvoipcallactiveentry[i]
    }
    return children
}

func (cvvoipcallactivetable *CISCOVOICEDIALCONTROLMIB_Cvvoipcallactivetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cvvoipcallactivetable *CISCOVOICEDIALCONTROLMIB_Cvvoipcallactivetable) GetBundleName() string { return "cisco_ios_xe" }

func (cvvoipcallactivetable *CISCOVOICEDIALCONTROLMIB_Cvvoipcallactivetable) GetYangName() string { return "cvVoIPCallActiveTable" }

func (cvvoipcallactivetable *CISCOVOICEDIALCONTROLMIB_Cvvoipcallactivetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvvoipcallactivetable *CISCOVOICEDIALCONTROLMIB_Cvvoipcallactivetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvvoipcallactivetable *CISCOVOICEDIALCONTROLMIB_Cvvoipcallactivetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvvoipcallactivetable *CISCOVOICEDIALCONTROLMIB_Cvvoipcallactivetable) SetParent(parent types.Entity) { cvvoipcallactivetable.parent = parent }

func (cvvoipcallactivetable *CISCOVOICEDIALCONTROLMIB_Cvvoipcallactivetable) GetParent() types.Entity { return cvvoipcallactivetable.parent }

func (cvvoipcallactivetable *CISCOVOICEDIALCONTROLMIB_Cvvoipcallactivetable) GetParentYangName() string { return "CISCO-VOICE-DIAL-CONTROL-MIB" }

// CISCOVOICEDIALCONTROLMIB_Cvvoipcallactivetable_Cvvoipcallactiveentry
// The information regarding a single VoIP call leg.
// The call leg entry is identified by using the same index
// objects that are used by Call Active table of IETF Dial
// Control MIB to identify the call.
// An entry of this table is created when its associated call
// active entry in the IETF Dial Control MIB is created and
// the call active entry contains information for the call
// establishment to the peer on the IP backbone via a voice
// over  IP peer.
// The entry is deleted when its associated call active entry
// in the IETF Dial Control MIB is deleted.
type CISCOVOICEDIALCONTROLMIB_Cvvoipcallactivetable_Cvvoipcallactiveentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // dial_control_mib.DIALCONTROLMIB_Callactivetable_Callactiveentry_Callactivesetuptime
    Callactivesetuptime interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // dial_control_mib.DIALCONTROLMIB_Callactivetable_Callactiveentry_Callactiveindex
    Callactiveindex interface{}

    // The global connection identifier for the active VoIP leg of the call. The
    // type is string with length: 0..16.
    Cvvoipcallactiveconnectionid interface{}

    // Remote system IP address for the VoIP call. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Cvvoipcallactiveremoteipaddress interface{}

    // Remote system UDP listener port to which to transmit voice packets. The
    // type is interface{} with range: 0..65535.
    Cvvoipcallactiveremoteudpport interface{}

    // The voice packet round trip delay between local and the remote system on
    // the IP backbone during the call. The type is interface{} with range:
    // 0..4294967295. Units are milliseconds.
    Cvvoipcallactiveroundtripdelay interface{}

    // The selected RSVP QoS for the voice call. The type is QosService.
    Cvvoipcallactiveselectedqos interface{}

    // The object specifies the session protocol to be used for Internet call
    // between local and remote router via IP backbone. The type is
    // CvSessionProtocol.
    Cvvoipcallactivesessionprotocol interface{}

    // The object specifies the session target of the peer that is used for the
    // call. This object is set with the information in the call associated
    // cvVoIPPeerCfgSessionTarget object when the voice over IP call is connected.
    // The type is string.
    Cvvoipcallactivesessiontarget interface{}

    // Duration of voice playout from data received on time for this call. This
    // plus the durations for the GapFills in the following entries gives the
    // Total Voice Playout Duration for Active Voice. This does not include
    // duration for which no data was sent by the Transmit end as voice signal,
    // e.g., silence suppression and fax signal. The On Time Playout Rate can be
    // computed by dividing this entry by the Total Voice Playout Duration. This
    // counter object will lock at the maximum value which is approximately two
    // days. The type is interface{} with range: 0..4294967295. Units are
    // milliseconds.
    Cvvoipcallactiveontimervplayout interface{}

    // Duration of voice signal replaced with signal played out during silence due
    // to voice data not received on time (or lost) from voice gateway this call.
    // This counter object will lock at the maximum value which is approximately
    // two days. The type is interface{} with range: 0..4294967295. Units are
    // milliseconds.
    Cvvoipcallactivegapfillwithsilence interface{}

    // Duration of voice signal played out with signal synthesized from parameters
    // or samples of data preceding in time due to voice data not received on time
    // (or lost) from voice gateway for this call. An example of such playout is
    // frame-erasure or frame-concealment strategies in G.729 and G.723.1
    // compression algorithms. This counter object will lock at the maximum value
    // which is approximately two days. The type is interface{} with range:
    // 0..4294967295. Units are milliseconds.
    Cvvoipcallactivegapfillwithprediction interface{}

    // Duration of voice signal played out with signal synthesized from parameters
    // or samples of data preceding and following in time due to voice data not
    // received on time (or lost) from voice gateway for this call. This counter
    // object will lock at the maximum value which is approximately two days. The
    // type is interface{} with range: 0..4294967295. Units are milliseconds.
    Cvvoipcallactivegapfillwithinterpolation interface{}

    // Duration of voice signal played out with signal synthesized from redundancy
    // parameters available due to voice data not received on time (or lost) from
    // voice gateway for this call. This counter object will lock at the maximum
    // value which is approximately two days. The type is interface{} with range:
    // 0..4294967295. Units are milliseconds.
    Cvvoipcallactivegapfillwithredundancy interface{}

    // The high water mark Voice Playout FIFO Delay during the voice call. This
    // counter object will lock at the maximum value which is approximately two
    // days. The type is interface{} with range: 0..4294967295. Units are
    // milliseconds.
    Cvvoipcallactivehiwaterplayoutdelay interface{}

    // The low water mark Voice Playout FIFO Delay during the voice call. The type
    // is interface{} with range: 0..4294967295. Units are milliseconds.
    Cvvoipcallactivelowaterplayoutdelay interface{}

    // The average Playout FIFO Delay plus the decoder delay during the voice
    // call. The type is interface{} with range: 0..4294967295.
    Cvvoipcallactivereceivedelay interface{}

    // The object indicates whether or not the VAD (Voice Activity Detection) was
    // enabled for the voice call. The type is bool.
    Cvvoipcallactivevadenable interface{}

    // The negotiated coder rate. It specifies the transmit rate of voice/fax
    // compression to its associated call leg for the call. This rate is different
    // from the configuration rate because of rate negotiation during the call.
    // The type is CvcCoderTypeRate.
    Cvvoipcallactivecodertyperate interface{}

    // The number of lost voice packets during the call. The type is interface{}
    // with range: 0..4294967295. Units are packets.
    Cvvoipcallactivelostpackets interface{}

    // The number of received voice packets that arrived too early to store in
    // jitter buffer during the call. The type is interface{} with range:
    // 0..4294967295. Units are packets.
    Cvvoipcallactiveearlypackets interface{}

    // The number of received voice packets that arrived too late to playout with
    // CODEC (Coder/Decoder) during the call. The type is interface{} with range:
    // 0..4294967295. Units are packets.
    Cvvoipcallactivelatepackets interface{}

    // The textual identifier of the calling party (user) of the call. If the
    // username is not available, then the value of this object will have a length
    // of zero. The type is string with length: 0..16.
    Cvvoipcallactiveusername interface{}

    // The protocol-specific call identifier for the VoIP call. The type is string
    // with length: 0..255.
    Cvvoipcallactiveprotocolcallid interface{}

    // This object specifies the type of address contained in the associated
    // instance of cvVoIPCallActiveRemSigIPAddr. The type is InetAddressType.
    Cvvoipcallactiveremsigipaddrt interface{}

    // Remote signalling IP address for the VoIP call. The type is string with
    // length: 0..255.
    Cvvoipcallactiveremsigipaddr interface{}

    // Remote signalling listener port to which to transmit voice packets. The
    // type is interface{} with range: 0..65535.
    Cvvoipcallactiveremsigport interface{}

    // This object specifies the type of address contained in the associated
    // instance of cvVoIPCallActiveRemMediaIPAddr. The type is InetAddressType.
    Cvvoipcallactiveremmediaipaddrt interface{}

    // Remote media end point IP address for the VoIP call. The type is string
    // with length: 0..255.
    Cvvoipcallactiveremmediaipaddr interface{}

    // Remote media end point listener port to which to transmit voice packets.
    // The type is interface{} with range: 0..65535.
    Cvvoipcallactiveremmediaport interface{}

    // The object indicates whether or not the SRTP (Secured RTP) was enabled for
    // the voice call. The type is bool.
    Cvvoipcallactivesrtpenable interface{}

    // If the object has a value true(1) octet align operation is used, and if the
    // value is false(2), bandwidth efficient operation is used. This object is
    // not instantiated when the object cvVoIPCallActiveCoderTypeRate is not equal
    // to gsmAmrNb enum. The type is bool.
    Cvvoipcallactiveoctetaligned interface{}

    // This object indicates modes of Bit rates. This object is not instantiated
    // when the object cvVoIPCallActiveCoderTypeRate is not equal to gsmAmrNb
    // enum. The type is map[string]bool.
    Cvvoipcallactivebitrates interface{}

    // The object indicates the interval (N frame-blocks) at which codec mode
    // changes are allowed. This object is not instantiated when the object
    // cvVoIPCallActiveCoderTypeRate is not equal to gsmAmrNb enum. The type is
    // interface{} with range: 1..100. Units are frame-blocks.
    Cvvoipcallactivemodechgperiod interface{}

    // If the object has a value of true(1), mode changes will be made to only
    // neighboring modes set to cvVoIPCallActiveBitRates object. If the value is
    // false(2), mode changes will be allowed to any modes set to
    // cvVoIPCallActiveBitRates object. This object is not instantiated when the
    // object cvVoIPCallActiveCoderTypeRate is not equal to gsmAmrNb enum. The
    // type is bool.
    Cvvoipcallactivemodechgneighbor interface{}

    // The object indicates the maximum amount of media that can be encapsulated
    // in a payload. Supported value is 20 msec. This object is not instantiated
    // when the object cvVoIPCallActiveCoderTypeRate is not equal to gsmAmrNb
    // enum. The type is interface{} with range: 20..100. Units are milliseconds.
    Cvvoipcallactivemaxptime interface{}

    // If the object has a value of true(1), frame CRC will be included in the
    // payload and if the value is false(2), frame CRC will not be included in the
    // payload. This object is applicable only when RTP frame type is octet
    // aligned. This object is not instantiated when the object
    // cvVoIPCallActiveCoderTypeRate is not equal to gsmAmrNb enum. The type is
    // bool.
    Cvvoipcallactivecrc interface{}

    // If the object has a value of true(1), payload employs robust sorting and if
    // the value is false(2), payload does not employ robust sorting. This object
    // is applicable only when RTP frame type is octet aligned. This object is not
    // instantiated when the object cvVoIPCallActiveCoderTypeRate is not equal to
    // gsmAmrNb enum. The type is bool.
    Cvvoipcallactiverobustsorting interface{}

    // The object indicates the RTP encapsulation type. Supported RTP
    // encapsulation type is RFC3267. This object is not instantiated when the
    // object cvVoIPCallActiveCoderTypeRate is not equal to gsmAmrNb enum. The
    // type is CvAmrNbRtpEncap.
    Cvvoipcallactiveencap interface{}

    // The object indicates the maximum number of frame-blocks allowed in an
    // interleaving group. It indicates that frame-block level interleaving will
    // be used for that session. If this object is not set, interleaving is not
    // used. This object is applicable only when RTP frame type is octet aligned.
    // This object is not instantiated when the object
    // cvVoIPCallActiveCoderTypeRate is not equal to gsmAmrNb enum. The type is
    // interface{} with range: 1..50. Units are frame-blocks.
    Cvvoipcallactiveinterleaving interface{}

    // The object indicates the length of the time in milliseconds represented by
    // the media of the packet. Supported value is 20 milliseconds. This object is
    // not instantiated when the object cvVoIPCallActiveCoderTypeRate is not equal
    // to gsmAmrNb enum. The type is interface{} with range: 20..100. Units are
    // milliseconds.
    Cvvoipcallactiveptime interface{}

    // The object indicates the number of audio channels. Supported value is 1.
    // This object is not instantiated when the object
    // cvVoIPCallActiveCoderTypeRate is not equal to gsmAmrNb enum. The type is
    // interface{} with range: 1..6. Units are channels.
    Cvvoipcallactivechannels interface{}

    // The object indicates the iLBC codec mode. The value of this object is valid
    // only if  cvVoIPCallActiveCoderTypeRate is equal to  'iLBC'. The type is
    // CvIlbcFrameMode. Units are milliseconds.
    Cvvoipcallactivecodermode interface{}

    // This object represents the call identifier for the active VoIP leg of the
    // call. The type is interface{} with range: 1..4294967295.
    Cvvoipcallactivecallid interface{}

    // The call reference ID associates the video call entry and voice call entry
    // of the same endpoint.  An audio-only call may or may not have a valid call
    // reference ID (i.e. value greater than zero), but in both cases, there will
    // not be a video call entry associated with it.    For a video call, the
    // video-specific information  is stored in a call entry in
    // cVideoSessionActive of CISCO-VIDEO-SESSION-MIB, in which the call reference
    // ID is also identified. The type is interface{} with range: 0..4294967295.
    Cvvoipcallactivecallreferenceid interface{}

    // This object holds the policy name. It could be media policy, DSCP policy
    // etc. The type is string.
    Ccvoipcallactivepolicyname interface{}

    // This object store the reversed direction peer address  If the address is
    // not available, then it will have a length of zero.  If the call is ingress
    // then it contains called number and if the call is egress then it contains
    // calling number. The type is string.
    Cvvoipcallactivereverseddirectionpeeraddress interface{}

    // This object indicates the active session ID assigned by the call manager to
    // identify call legs that belong to the same call session. The type is
    // interface{} with range: 0..4294967295.
    Cvvoipcallactivesessionid interface{}
}

func (cvvoipcallactiveentry *CISCOVOICEDIALCONTROLMIB_Cvvoipcallactivetable_Cvvoipcallactiveentry) GetFilter() yfilter.YFilter { return cvvoipcallactiveentry.YFilter }

func (cvvoipcallactiveentry *CISCOVOICEDIALCONTROLMIB_Cvvoipcallactivetable_Cvvoipcallactiveentry) SetFilter(yf yfilter.YFilter) { cvvoipcallactiveentry.YFilter = yf }

func (cvvoipcallactiveentry *CISCOVOICEDIALCONTROLMIB_Cvvoipcallactivetable_Cvvoipcallactiveentry) GetGoName(yname string) string {
    if yname == "callActiveSetupTime" { return "Callactivesetuptime" }
    if yname == "callActiveIndex" { return "Callactiveindex" }
    if yname == "cvVoIPCallActiveConnectionId" { return "Cvvoipcallactiveconnectionid" }
    if yname == "cvVoIPCallActiveRemoteIPAddress" { return "Cvvoipcallactiveremoteipaddress" }
    if yname == "cvVoIPCallActiveRemoteUDPPort" { return "Cvvoipcallactiveremoteudpport" }
    if yname == "cvVoIPCallActiveRoundTripDelay" { return "Cvvoipcallactiveroundtripdelay" }
    if yname == "cvVoIPCallActiveSelectedQoS" { return "Cvvoipcallactiveselectedqos" }
    if yname == "cvVoIPCallActiveSessionProtocol" { return "Cvvoipcallactivesessionprotocol" }
    if yname == "cvVoIPCallActiveSessionTarget" { return "Cvvoipcallactivesessiontarget" }
    if yname == "cvVoIPCallActiveOnTimeRvPlayout" { return "Cvvoipcallactiveontimervplayout" }
    if yname == "cvVoIPCallActiveGapFillWithSilence" { return "Cvvoipcallactivegapfillwithsilence" }
    if yname == "cvVoIPCallActiveGapFillWithPrediction" { return "Cvvoipcallactivegapfillwithprediction" }
    if yname == "cvVoIPCallActiveGapFillWithInterpolation" { return "Cvvoipcallactivegapfillwithinterpolation" }
    if yname == "cvVoIPCallActiveGapFillWithRedundancy" { return "Cvvoipcallactivegapfillwithredundancy" }
    if yname == "cvVoIPCallActiveHiWaterPlayoutDelay" { return "Cvvoipcallactivehiwaterplayoutdelay" }
    if yname == "cvVoIPCallActiveLoWaterPlayoutDelay" { return "Cvvoipcallactivelowaterplayoutdelay" }
    if yname == "cvVoIPCallActiveReceiveDelay" { return "Cvvoipcallactivereceivedelay" }
    if yname == "cvVoIPCallActiveVADEnable" { return "Cvvoipcallactivevadenable" }
    if yname == "cvVoIPCallActiveCoderTypeRate" { return "Cvvoipcallactivecodertyperate" }
    if yname == "cvVoIPCallActiveLostPackets" { return "Cvvoipcallactivelostpackets" }
    if yname == "cvVoIPCallActiveEarlyPackets" { return "Cvvoipcallactiveearlypackets" }
    if yname == "cvVoIPCallActiveLatePackets" { return "Cvvoipcallactivelatepackets" }
    if yname == "cvVoIPCallActiveUsername" { return "Cvvoipcallactiveusername" }
    if yname == "cvVoIPCallActiveProtocolCallId" { return "Cvvoipcallactiveprotocolcallid" }
    if yname == "cvVoIPCallActiveRemSigIPAddrT" { return "Cvvoipcallactiveremsigipaddrt" }
    if yname == "cvVoIPCallActiveRemSigIPAddr" { return "Cvvoipcallactiveremsigipaddr" }
    if yname == "cvVoIPCallActiveRemSigPort" { return "Cvvoipcallactiveremsigport" }
    if yname == "cvVoIPCallActiveRemMediaIPAddrT" { return "Cvvoipcallactiveremmediaipaddrt" }
    if yname == "cvVoIPCallActiveRemMediaIPAddr" { return "Cvvoipcallactiveremmediaipaddr" }
    if yname == "cvVoIPCallActiveRemMediaPort" { return "Cvvoipcallactiveremmediaport" }
    if yname == "cvVoIPCallActiveSRTPEnable" { return "Cvvoipcallactivesrtpenable" }
    if yname == "cvVoIPCallActiveOctetAligned" { return "Cvvoipcallactiveoctetaligned" }
    if yname == "cvVoIPCallActiveBitRates" { return "Cvvoipcallactivebitrates" }
    if yname == "cvVoIPCallActiveModeChgPeriod" { return "Cvvoipcallactivemodechgperiod" }
    if yname == "cvVoIPCallActiveModeChgNeighbor" { return "Cvvoipcallactivemodechgneighbor" }
    if yname == "cvVoIPCallActiveMaxPtime" { return "Cvvoipcallactivemaxptime" }
    if yname == "cvVoIPCallActiveCRC" { return "Cvvoipcallactivecrc" }
    if yname == "cvVoIPCallActiveRobustSorting" { return "Cvvoipcallactiverobustsorting" }
    if yname == "cvVoIPCallActiveEncap" { return "Cvvoipcallactiveencap" }
    if yname == "cvVoIPCallActiveInterleaving" { return "Cvvoipcallactiveinterleaving" }
    if yname == "cvVoIPCallActivePtime" { return "Cvvoipcallactiveptime" }
    if yname == "cvVoIPCallActiveChannels" { return "Cvvoipcallactivechannels" }
    if yname == "cvVoIPCallActiveCoderMode" { return "Cvvoipcallactivecodermode" }
    if yname == "cvVoIPCallActiveCallId" { return "Cvvoipcallactivecallid" }
    if yname == "cvVoIPCallActiveCallReferenceId" { return "Cvvoipcallactivecallreferenceid" }
    if yname == "ccVoIPCallActivePolicyName" { return "Ccvoipcallactivepolicyname" }
    if yname == "cvVoIPCallActiveReversedDirectionPeerAddress" { return "Cvvoipcallactivereverseddirectionpeeraddress" }
    if yname == "cvVoIPCallActiveSessionId" { return "Cvvoipcallactivesessionid" }
    return ""
}

func (cvvoipcallactiveentry *CISCOVOICEDIALCONTROLMIB_Cvvoipcallactivetable_Cvvoipcallactiveentry) GetSegmentPath() string {
    return "cvVoIPCallActiveEntry" + "[callActiveSetupTime='" + fmt.Sprintf("%v", cvvoipcallactiveentry.Callactivesetuptime) + "']" + "[callActiveIndex='" + fmt.Sprintf("%v", cvvoipcallactiveentry.Callactiveindex) + "']"
}

func (cvvoipcallactiveentry *CISCOVOICEDIALCONTROLMIB_Cvvoipcallactivetable_Cvvoipcallactiveentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cvvoipcallactiveentry *CISCOVOICEDIALCONTROLMIB_Cvvoipcallactivetable_Cvvoipcallactiveentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cvvoipcallactiveentry *CISCOVOICEDIALCONTROLMIB_Cvvoipcallactivetable_Cvvoipcallactiveentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["callActiveSetupTime"] = cvvoipcallactiveentry.Callactivesetuptime
    leafs["callActiveIndex"] = cvvoipcallactiveentry.Callactiveindex
    leafs["cvVoIPCallActiveConnectionId"] = cvvoipcallactiveentry.Cvvoipcallactiveconnectionid
    leafs["cvVoIPCallActiveRemoteIPAddress"] = cvvoipcallactiveentry.Cvvoipcallactiveremoteipaddress
    leafs["cvVoIPCallActiveRemoteUDPPort"] = cvvoipcallactiveentry.Cvvoipcallactiveremoteudpport
    leafs["cvVoIPCallActiveRoundTripDelay"] = cvvoipcallactiveentry.Cvvoipcallactiveroundtripdelay
    leafs["cvVoIPCallActiveSelectedQoS"] = cvvoipcallactiveentry.Cvvoipcallactiveselectedqos
    leafs["cvVoIPCallActiveSessionProtocol"] = cvvoipcallactiveentry.Cvvoipcallactivesessionprotocol
    leafs["cvVoIPCallActiveSessionTarget"] = cvvoipcallactiveentry.Cvvoipcallactivesessiontarget
    leafs["cvVoIPCallActiveOnTimeRvPlayout"] = cvvoipcallactiveentry.Cvvoipcallactiveontimervplayout
    leafs["cvVoIPCallActiveGapFillWithSilence"] = cvvoipcallactiveentry.Cvvoipcallactivegapfillwithsilence
    leafs["cvVoIPCallActiveGapFillWithPrediction"] = cvvoipcallactiveentry.Cvvoipcallactivegapfillwithprediction
    leafs["cvVoIPCallActiveGapFillWithInterpolation"] = cvvoipcallactiveentry.Cvvoipcallactivegapfillwithinterpolation
    leafs["cvVoIPCallActiveGapFillWithRedundancy"] = cvvoipcallactiveentry.Cvvoipcallactivegapfillwithredundancy
    leafs["cvVoIPCallActiveHiWaterPlayoutDelay"] = cvvoipcallactiveentry.Cvvoipcallactivehiwaterplayoutdelay
    leafs["cvVoIPCallActiveLoWaterPlayoutDelay"] = cvvoipcallactiveentry.Cvvoipcallactivelowaterplayoutdelay
    leafs["cvVoIPCallActiveReceiveDelay"] = cvvoipcallactiveentry.Cvvoipcallactivereceivedelay
    leafs["cvVoIPCallActiveVADEnable"] = cvvoipcallactiveentry.Cvvoipcallactivevadenable
    leafs["cvVoIPCallActiveCoderTypeRate"] = cvvoipcallactiveentry.Cvvoipcallactivecodertyperate
    leafs["cvVoIPCallActiveLostPackets"] = cvvoipcallactiveentry.Cvvoipcallactivelostpackets
    leafs["cvVoIPCallActiveEarlyPackets"] = cvvoipcallactiveentry.Cvvoipcallactiveearlypackets
    leafs["cvVoIPCallActiveLatePackets"] = cvvoipcallactiveentry.Cvvoipcallactivelatepackets
    leafs["cvVoIPCallActiveUsername"] = cvvoipcallactiveentry.Cvvoipcallactiveusername
    leafs["cvVoIPCallActiveProtocolCallId"] = cvvoipcallactiveentry.Cvvoipcallactiveprotocolcallid
    leafs["cvVoIPCallActiveRemSigIPAddrT"] = cvvoipcallactiveentry.Cvvoipcallactiveremsigipaddrt
    leafs["cvVoIPCallActiveRemSigIPAddr"] = cvvoipcallactiveentry.Cvvoipcallactiveremsigipaddr
    leafs["cvVoIPCallActiveRemSigPort"] = cvvoipcallactiveentry.Cvvoipcallactiveremsigport
    leafs["cvVoIPCallActiveRemMediaIPAddrT"] = cvvoipcallactiveentry.Cvvoipcallactiveremmediaipaddrt
    leafs["cvVoIPCallActiveRemMediaIPAddr"] = cvvoipcallactiveentry.Cvvoipcallactiveremmediaipaddr
    leafs["cvVoIPCallActiveRemMediaPort"] = cvvoipcallactiveentry.Cvvoipcallactiveremmediaport
    leafs["cvVoIPCallActiveSRTPEnable"] = cvvoipcallactiveentry.Cvvoipcallactivesrtpenable
    leafs["cvVoIPCallActiveOctetAligned"] = cvvoipcallactiveentry.Cvvoipcallactiveoctetaligned
    leafs["cvVoIPCallActiveBitRates"] = cvvoipcallactiveentry.Cvvoipcallactivebitrates
    leafs["cvVoIPCallActiveModeChgPeriod"] = cvvoipcallactiveentry.Cvvoipcallactivemodechgperiod
    leafs["cvVoIPCallActiveModeChgNeighbor"] = cvvoipcallactiveentry.Cvvoipcallactivemodechgneighbor
    leafs["cvVoIPCallActiveMaxPtime"] = cvvoipcallactiveentry.Cvvoipcallactivemaxptime
    leafs["cvVoIPCallActiveCRC"] = cvvoipcallactiveentry.Cvvoipcallactivecrc
    leafs["cvVoIPCallActiveRobustSorting"] = cvvoipcallactiveentry.Cvvoipcallactiverobustsorting
    leafs["cvVoIPCallActiveEncap"] = cvvoipcallactiveentry.Cvvoipcallactiveencap
    leafs["cvVoIPCallActiveInterleaving"] = cvvoipcallactiveentry.Cvvoipcallactiveinterleaving
    leafs["cvVoIPCallActivePtime"] = cvvoipcallactiveentry.Cvvoipcallactiveptime
    leafs["cvVoIPCallActiveChannels"] = cvvoipcallactiveentry.Cvvoipcallactivechannels
    leafs["cvVoIPCallActiveCoderMode"] = cvvoipcallactiveentry.Cvvoipcallactivecodermode
    leafs["cvVoIPCallActiveCallId"] = cvvoipcallactiveentry.Cvvoipcallactivecallid
    leafs["cvVoIPCallActiveCallReferenceId"] = cvvoipcallactiveentry.Cvvoipcallactivecallreferenceid
    leafs["ccVoIPCallActivePolicyName"] = cvvoipcallactiveentry.Ccvoipcallactivepolicyname
    leafs["cvVoIPCallActiveReversedDirectionPeerAddress"] = cvvoipcallactiveentry.Cvvoipcallactivereverseddirectionpeeraddress
    leafs["cvVoIPCallActiveSessionId"] = cvvoipcallactiveentry.Cvvoipcallactivesessionid
    return leafs
}

func (cvvoipcallactiveentry *CISCOVOICEDIALCONTROLMIB_Cvvoipcallactivetable_Cvvoipcallactiveentry) GetBundleName() string { return "cisco_ios_xe" }

func (cvvoipcallactiveentry *CISCOVOICEDIALCONTROLMIB_Cvvoipcallactivetable_Cvvoipcallactiveentry) GetYangName() string { return "cvVoIPCallActiveEntry" }

func (cvvoipcallactiveentry *CISCOVOICEDIALCONTROLMIB_Cvvoipcallactivetable_Cvvoipcallactiveentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvvoipcallactiveentry *CISCOVOICEDIALCONTROLMIB_Cvvoipcallactivetable_Cvvoipcallactiveentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvvoipcallactiveentry *CISCOVOICEDIALCONTROLMIB_Cvvoipcallactivetable_Cvvoipcallactiveentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvvoipcallactiveentry *CISCOVOICEDIALCONTROLMIB_Cvvoipcallactivetable_Cvvoipcallactiveentry) SetParent(parent types.Entity) { cvvoipcallactiveentry.parent = parent }

func (cvvoipcallactiveentry *CISCOVOICEDIALCONTROLMIB_Cvvoipcallactivetable_Cvvoipcallactiveentry) GetParent() types.Entity { return cvvoipcallactiveentry.parent }

func (cvvoipcallactiveentry *CISCOVOICEDIALCONTROLMIB_Cvvoipcallactivetable_Cvvoipcallactiveentry) GetParentYangName() string { return "cvVoIPCallActiveTable" }

// CISCOVOICEDIALCONTROLMIB_Cvcallvolconntable
// This table represents the number of active
// call connections for each call connection type
// in the voice gateway.
type CISCOVOICEDIALCONTROLMIB_Cvcallvolconntable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the cvCallVolConnTable indicates number of active calls for a
    // call connection type in the voice gateway. The type is slice of
    // CISCOVOICEDIALCONTROLMIB_Cvcallvolconntable_Cvcallvolconnentry.
    Cvcallvolconnentry []CISCOVOICEDIALCONTROLMIB_Cvcallvolconntable_Cvcallvolconnentry
}

func (cvcallvolconntable *CISCOVOICEDIALCONTROLMIB_Cvcallvolconntable) GetFilter() yfilter.YFilter { return cvcallvolconntable.YFilter }

func (cvcallvolconntable *CISCOVOICEDIALCONTROLMIB_Cvcallvolconntable) SetFilter(yf yfilter.YFilter) { cvcallvolconntable.YFilter = yf }

func (cvcallvolconntable *CISCOVOICEDIALCONTROLMIB_Cvcallvolconntable) GetGoName(yname string) string {
    if yname == "cvCallVolConnEntry" { return "Cvcallvolconnentry" }
    return ""
}

func (cvcallvolconntable *CISCOVOICEDIALCONTROLMIB_Cvcallvolconntable) GetSegmentPath() string {
    return "cvCallVolConnTable"
}

func (cvcallvolconntable *CISCOVOICEDIALCONTROLMIB_Cvcallvolconntable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cvCallVolConnEntry" {
        for _, c := range cvcallvolconntable.Cvcallvolconnentry {
            if cvcallvolconntable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOVOICEDIALCONTROLMIB_Cvcallvolconntable_Cvcallvolconnentry{}
        cvcallvolconntable.Cvcallvolconnentry = append(cvcallvolconntable.Cvcallvolconnentry, child)
        return &cvcallvolconntable.Cvcallvolconnentry[len(cvcallvolconntable.Cvcallvolconnentry)-1]
    }
    return nil
}

func (cvcallvolconntable *CISCOVOICEDIALCONTROLMIB_Cvcallvolconntable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cvcallvolconntable.Cvcallvolconnentry {
        children[cvcallvolconntable.Cvcallvolconnentry[i].GetSegmentPath()] = &cvcallvolconntable.Cvcallvolconnentry[i]
    }
    return children
}

func (cvcallvolconntable *CISCOVOICEDIALCONTROLMIB_Cvcallvolconntable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cvcallvolconntable *CISCOVOICEDIALCONTROLMIB_Cvcallvolconntable) GetBundleName() string { return "cisco_ios_xe" }

func (cvcallvolconntable *CISCOVOICEDIALCONTROLMIB_Cvcallvolconntable) GetYangName() string { return "cvCallVolConnTable" }

func (cvcallvolconntable *CISCOVOICEDIALCONTROLMIB_Cvcallvolconntable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvcallvolconntable *CISCOVOICEDIALCONTROLMIB_Cvcallvolconntable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvcallvolconntable *CISCOVOICEDIALCONTROLMIB_Cvcallvolconntable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvcallvolconntable *CISCOVOICEDIALCONTROLMIB_Cvcallvolconntable) SetParent(parent types.Entity) { cvcallvolconntable.parent = parent }

func (cvcallvolconntable *CISCOVOICEDIALCONTROLMIB_Cvcallvolconntable) GetParent() types.Entity { return cvcallvolconntable.parent }

func (cvcallvolconntable *CISCOVOICEDIALCONTROLMIB_Cvcallvolconntable) GetParentYangName() string { return "CISCO-VOICE-DIAL-CONTROL-MIB" }

// CISCOVOICEDIALCONTROLMIB_Cvcallvolconntable_Cvcallvolconnentry
// An entry in the cvCallVolConnTable indicates
// number of active calls for a call connection type
// in the voice gateway.
type CISCOVOICEDIALCONTROLMIB_Cvcallvolconntable_Cvcallvolconnentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. This object represents index to the
    // cvCallVolConnTable. The type is CvCallConnectionType.
    Cvcallvolconnindex interface{}

    // This object represents the total number of active calls for a connection
    // type  in the voice gateway. The type is interface{} with range: 0..65535.
    Cvcallvolconnactiveconnection interface{}
}

func (cvcallvolconnentry *CISCOVOICEDIALCONTROLMIB_Cvcallvolconntable_Cvcallvolconnentry) GetFilter() yfilter.YFilter { return cvcallvolconnentry.YFilter }

func (cvcallvolconnentry *CISCOVOICEDIALCONTROLMIB_Cvcallvolconntable_Cvcallvolconnentry) SetFilter(yf yfilter.YFilter) { cvcallvolconnentry.YFilter = yf }

func (cvcallvolconnentry *CISCOVOICEDIALCONTROLMIB_Cvcallvolconntable_Cvcallvolconnentry) GetGoName(yname string) string {
    if yname == "cvCallVolConnIndex" { return "Cvcallvolconnindex" }
    if yname == "cvCallVolConnActiveConnection" { return "Cvcallvolconnactiveconnection" }
    return ""
}

func (cvcallvolconnentry *CISCOVOICEDIALCONTROLMIB_Cvcallvolconntable_Cvcallvolconnentry) GetSegmentPath() string {
    return "cvCallVolConnEntry" + "[cvCallVolConnIndex='" + fmt.Sprintf("%v", cvcallvolconnentry.Cvcallvolconnindex) + "']"
}

func (cvcallvolconnentry *CISCOVOICEDIALCONTROLMIB_Cvcallvolconntable_Cvcallvolconnentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cvcallvolconnentry *CISCOVOICEDIALCONTROLMIB_Cvcallvolconntable_Cvcallvolconnentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cvcallvolconnentry *CISCOVOICEDIALCONTROLMIB_Cvcallvolconntable_Cvcallvolconnentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cvCallVolConnIndex"] = cvcallvolconnentry.Cvcallvolconnindex
    leafs["cvCallVolConnActiveConnection"] = cvcallvolconnentry.Cvcallvolconnactiveconnection
    return leafs
}

func (cvcallvolconnentry *CISCOVOICEDIALCONTROLMIB_Cvcallvolconntable_Cvcallvolconnentry) GetBundleName() string { return "cisco_ios_xe" }

func (cvcallvolconnentry *CISCOVOICEDIALCONTROLMIB_Cvcallvolconntable_Cvcallvolconnentry) GetYangName() string { return "cvCallVolConnEntry" }

func (cvcallvolconnentry *CISCOVOICEDIALCONTROLMIB_Cvcallvolconntable_Cvcallvolconnentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvcallvolconnentry *CISCOVOICEDIALCONTROLMIB_Cvcallvolconntable_Cvcallvolconnentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvcallvolconnentry *CISCOVOICEDIALCONTROLMIB_Cvcallvolconntable_Cvcallvolconnentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvcallvolconnentry *CISCOVOICEDIALCONTROLMIB_Cvcallvolconntable_Cvcallvolconnentry) SetParent(parent types.Entity) { cvcallvolconnentry.parent = parent }

func (cvcallvolconnentry *CISCOVOICEDIALCONTROLMIB_Cvcallvolconntable_Cvcallvolconnentry) GetParent() types.Entity { return cvcallvolconnentry.parent }

func (cvcallvolconnentry *CISCOVOICEDIALCONTROLMIB_Cvcallvolconntable_Cvcallvolconnentry) GetParentYangName() string { return "cvCallVolConnTable" }

// CISCOVOICEDIALCONTROLMIB_Cvcallvoliftable
// This table represents the information about
// the usage of an IP interface in a voice
// gateway for voice media calls. This table
// has a sparse-dependent relationship with  
// ifTable. There exists an entry in this table, 
// for each of the  entries in ifTable where ifType 
// is one of 'ethernetCsmacd' and 'softwareLoopback'.
type CISCOVOICEDIALCONTROLMIB_Cvcallvoliftable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry represents a row in cvCallVolIfTable and corresponds to the
    // information about an IP  interface in the voice gateway. The type is slice
    // of CISCOVOICEDIALCONTROLMIB_Cvcallvoliftable_Cvcallvolifentry.
    Cvcallvolifentry []CISCOVOICEDIALCONTROLMIB_Cvcallvoliftable_Cvcallvolifentry
}

func (cvcallvoliftable *CISCOVOICEDIALCONTROLMIB_Cvcallvoliftable) GetFilter() yfilter.YFilter { return cvcallvoliftable.YFilter }

func (cvcallvoliftable *CISCOVOICEDIALCONTROLMIB_Cvcallvoliftable) SetFilter(yf yfilter.YFilter) { cvcallvoliftable.YFilter = yf }

func (cvcallvoliftable *CISCOVOICEDIALCONTROLMIB_Cvcallvoliftable) GetGoName(yname string) string {
    if yname == "cvCallVolIfEntry" { return "Cvcallvolifentry" }
    return ""
}

func (cvcallvoliftable *CISCOVOICEDIALCONTROLMIB_Cvcallvoliftable) GetSegmentPath() string {
    return "cvCallVolIfTable"
}

func (cvcallvoliftable *CISCOVOICEDIALCONTROLMIB_Cvcallvoliftable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cvCallVolIfEntry" {
        for _, c := range cvcallvoliftable.Cvcallvolifentry {
            if cvcallvoliftable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOVOICEDIALCONTROLMIB_Cvcallvoliftable_Cvcallvolifentry{}
        cvcallvoliftable.Cvcallvolifentry = append(cvcallvoliftable.Cvcallvolifentry, child)
        return &cvcallvoliftable.Cvcallvolifentry[len(cvcallvoliftable.Cvcallvolifentry)-1]
    }
    return nil
}

func (cvcallvoliftable *CISCOVOICEDIALCONTROLMIB_Cvcallvoliftable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cvcallvoliftable.Cvcallvolifentry {
        children[cvcallvoliftable.Cvcallvolifentry[i].GetSegmentPath()] = &cvcallvoliftable.Cvcallvolifentry[i]
    }
    return children
}

func (cvcallvoliftable *CISCOVOICEDIALCONTROLMIB_Cvcallvoliftable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cvcallvoliftable *CISCOVOICEDIALCONTROLMIB_Cvcallvoliftable) GetBundleName() string { return "cisco_ios_xe" }

func (cvcallvoliftable *CISCOVOICEDIALCONTROLMIB_Cvcallvoliftable) GetYangName() string { return "cvCallVolIfTable" }

func (cvcallvoliftable *CISCOVOICEDIALCONTROLMIB_Cvcallvoliftable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvcallvoliftable *CISCOVOICEDIALCONTROLMIB_Cvcallvoliftable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvcallvoliftable *CISCOVOICEDIALCONTROLMIB_Cvcallvoliftable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvcallvoliftable *CISCOVOICEDIALCONTROLMIB_Cvcallvoliftable) SetParent(parent types.Entity) { cvcallvoliftable.parent = parent }

func (cvcallvoliftable *CISCOVOICEDIALCONTROLMIB_Cvcallvoliftable) GetParent() types.Entity { return cvcallvoliftable.parent }

func (cvcallvoliftable *CISCOVOICEDIALCONTROLMIB_Cvcallvoliftable) GetParentYangName() string { return "CISCO-VOICE-DIAL-CONTROL-MIB" }

// CISCOVOICEDIALCONTROLMIB_Cvcallvoliftable_Cvcallvolifentry
// Each entry represents a row in cvCallVolIfTable
// and corresponds to the information about an IP 
// interface in the voice gateway.
type CISCOVOICEDIALCONTROLMIB_Cvcallvoliftable_Cvcallvolifentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This object represents the total number of inbound active media calls
    // through this IP  interface. The type is interface{} with range: 0..65535.
    Cvcallvolmediaincomingcalls interface{}

    // This object represents the total number of outbound active media calls
    // through the IP  interface. The type is interface{} with range: 0..65535.
    Cvcallvolmediaoutgoingcalls interface{}
}

func (cvcallvolifentry *CISCOVOICEDIALCONTROLMIB_Cvcallvoliftable_Cvcallvolifentry) GetFilter() yfilter.YFilter { return cvcallvolifentry.YFilter }

func (cvcallvolifentry *CISCOVOICEDIALCONTROLMIB_Cvcallvoliftable_Cvcallvolifentry) SetFilter(yf yfilter.YFilter) { cvcallvolifentry.YFilter = yf }

func (cvcallvolifentry *CISCOVOICEDIALCONTROLMIB_Cvcallvoliftable_Cvcallvolifentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "cvCallVolMediaIncomingCalls" { return "Cvcallvolmediaincomingcalls" }
    if yname == "cvCallVolMediaOutgoingCalls" { return "Cvcallvolmediaoutgoingcalls" }
    return ""
}

func (cvcallvolifentry *CISCOVOICEDIALCONTROLMIB_Cvcallvoliftable_Cvcallvolifentry) GetSegmentPath() string {
    return "cvCallVolIfEntry" + "[ifIndex='" + fmt.Sprintf("%v", cvcallvolifentry.Ifindex) + "']"
}

func (cvcallvolifentry *CISCOVOICEDIALCONTROLMIB_Cvcallvoliftable_Cvcallvolifentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cvcallvolifentry *CISCOVOICEDIALCONTROLMIB_Cvcallvoliftable_Cvcallvolifentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cvcallvolifentry *CISCOVOICEDIALCONTROLMIB_Cvcallvoliftable_Cvcallvolifentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = cvcallvolifentry.Ifindex
    leafs["cvCallVolMediaIncomingCalls"] = cvcallvolifentry.Cvcallvolmediaincomingcalls
    leafs["cvCallVolMediaOutgoingCalls"] = cvcallvolifentry.Cvcallvolmediaoutgoingcalls
    return leafs
}

func (cvcallvolifentry *CISCOVOICEDIALCONTROLMIB_Cvcallvoliftable_Cvcallvolifentry) GetBundleName() string { return "cisco_ios_xe" }

func (cvcallvolifentry *CISCOVOICEDIALCONTROLMIB_Cvcallvoliftable_Cvcallvolifentry) GetYangName() string { return "cvCallVolIfEntry" }

func (cvcallvolifentry *CISCOVOICEDIALCONTROLMIB_Cvcallvoliftable_Cvcallvolifentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvcallvolifentry *CISCOVOICEDIALCONTROLMIB_Cvcallvoliftable_Cvcallvolifentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvcallvolifentry *CISCOVOICEDIALCONTROLMIB_Cvcallvoliftable_Cvcallvolifentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvcallvolifentry *CISCOVOICEDIALCONTROLMIB_Cvcallvoliftable_Cvcallvolifentry) SetParent(parent types.Entity) { cvcallvolifentry.parent = parent }

func (cvcallvolifentry *CISCOVOICEDIALCONTROLMIB_Cvcallvoliftable_Cvcallvolifentry) GetParent() types.Entity { return cvcallvolifentry.parent }

func (cvcallvolifentry *CISCOVOICEDIALCONTROLMIB_Cvcallvoliftable_Cvcallvolifentry) GetParentYangName() string { return "cvCallVolIfTable" }

// CISCOVOICEDIALCONTROLMIB_Cvcallhistorytable
// This table is the voice extension to the call history table
// of IETF Dial Control MIB. It contains voice encapsulation
// call leg information such as voice packet statistics,
// coder usage and end to end bandwidth of the call leg.
type CISCOVOICEDIALCONTROLMIB_Cvcallhistorytable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The information regarding a single voice encapsulation call leg. The call
    // leg entry is identified by using the same index objects that are used by
    // Call Active table of IETF Dial Control MIB to identify the call. An entry
    // of this table is created when its associated call history entry in the IETF
    // Dial Control MIB is created and the call history entry contains the call
    // establishment to a voice encapsulation peer. The entry is deleted when its
    // associated call active entry in the IETF Dial Control MIB is deleted. The
    // type is slice of
    // CISCOVOICEDIALCONTROLMIB_Cvcallhistorytable_Cvcallhistoryentry.
    Cvcallhistoryentry []CISCOVOICEDIALCONTROLMIB_Cvcallhistorytable_Cvcallhistoryentry
}

func (cvcallhistorytable *CISCOVOICEDIALCONTROLMIB_Cvcallhistorytable) GetFilter() yfilter.YFilter { return cvcallhistorytable.YFilter }

func (cvcallhistorytable *CISCOVOICEDIALCONTROLMIB_Cvcallhistorytable) SetFilter(yf yfilter.YFilter) { cvcallhistorytable.YFilter = yf }

func (cvcallhistorytable *CISCOVOICEDIALCONTROLMIB_Cvcallhistorytable) GetGoName(yname string) string {
    if yname == "cvCallHistoryEntry" { return "Cvcallhistoryentry" }
    return ""
}

func (cvcallhistorytable *CISCOVOICEDIALCONTROLMIB_Cvcallhistorytable) GetSegmentPath() string {
    return "cvCallHistoryTable"
}

func (cvcallhistorytable *CISCOVOICEDIALCONTROLMIB_Cvcallhistorytable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cvCallHistoryEntry" {
        for _, c := range cvcallhistorytable.Cvcallhistoryentry {
            if cvcallhistorytable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOVOICEDIALCONTROLMIB_Cvcallhistorytable_Cvcallhistoryentry{}
        cvcallhistorytable.Cvcallhistoryentry = append(cvcallhistorytable.Cvcallhistoryentry, child)
        return &cvcallhistorytable.Cvcallhistoryentry[len(cvcallhistorytable.Cvcallhistoryentry)-1]
    }
    return nil
}

func (cvcallhistorytable *CISCOVOICEDIALCONTROLMIB_Cvcallhistorytable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cvcallhistorytable.Cvcallhistoryentry {
        children[cvcallhistorytable.Cvcallhistoryentry[i].GetSegmentPath()] = &cvcallhistorytable.Cvcallhistoryentry[i]
    }
    return children
}

func (cvcallhistorytable *CISCOVOICEDIALCONTROLMIB_Cvcallhistorytable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cvcallhistorytable *CISCOVOICEDIALCONTROLMIB_Cvcallhistorytable) GetBundleName() string { return "cisco_ios_xe" }

func (cvcallhistorytable *CISCOVOICEDIALCONTROLMIB_Cvcallhistorytable) GetYangName() string { return "cvCallHistoryTable" }

func (cvcallhistorytable *CISCOVOICEDIALCONTROLMIB_Cvcallhistorytable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvcallhistorytable *CISCOVOICEDIALCONTROLMIB_Cvcallhistorytable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvcallhistorytable *CISCOVOICEDIALCONTROLMIB_Cvcallhistorytable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvcallhistorytable *CISCOVOICEDIALCONTROLMIB_Cvcallhistorytable) SetParent(parent types.Entity) { cvcallhistorytable.parent = parent }

func (cvcallhistorytable *CISCOVOICEDIALCONTROLMIB_Cvcallhistorytable) GetParent() types.Entity { return cvcallhistorytable.parent }

func (cvcallhistorytable *CISCOVOICEDIALCONTROLMIB_Cvcallhistorytable) GetParentYangName() string { return "CISCO-VOICE-DIAL-CONTROL-MIB" }

// CISCOVOICEDIALCONTROLMIB_Cvcallhistorytable_Cvcallhistoryentry
// The information regarding a single voice encapsulation
// call leg.
// The call leg entry is identified by using the same index
// objects that are used by Call Active table of IETF Dial
// Control MIB to identify the call.
// An entry of this table is created when its associated call
// history entry in the IETF Dial Control MIB is created and
// the call history entry contains the call establishment to
// a voice encapsulation peer.
// The entry is deleted when its associated call active entry
// in the IETF Dial Control MIB is deleted.
type CISCOVOICEDIALCONTROLMIB_Cvcallhistorytable_Cvcallhistoryentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // cisco_dial_control_mib.CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryindex
    Ccallhistoryindex interface{}

    // The global connection identifier for the telephony leg, which was assigned
    // to the call. The type is string with length: 0..16.
    Cvcallhistoryconnectionid interface{}

    // Duration of Transmit path open from this peer to the voice gateway for the
    // call leg. This counter object will lock at the maximum value which is
    // approximately two days. The type is interface{} with range: 0..4294967295.
    // Units are milliseconds.
    Cvcallhistorytxduration interface{}

    // Duration for this call leg. The Voice Utilization Rate can be obtained by
    // dividing this by cvCallHistoryTXDuration object. This counter object will
    // lock at the maximum value which is approximately two days. The type is
    // interface{} with range: 0..4294967295. Units are milliseconds.
    Cvcallhistoryvoicetxduration interface{}

    // Duration of fax transmitted from this peer to voice gateway for this call
    // leg. The FAX Utilization Rate can be obtained by dividing this by
    // cvCallHistoryTXDuration object. This counter object will lock at the
    // maximum value which is approximately two days. The type is interface{} with
    // range: 0..4294967295. Units are milliseconds.
    Cvcallhistoryfaxtxduration interface{}

    // The negotiated coder rate. It specifies the transmit rate of voice/fax
    // compression to its associated call leg for the call. The type is
    // CvcCoderTypeRate.
    Cvcallhistorycodertyperate interface{}

    // The object contains the average noise level for the call leg. The type is
    // interface{} with range: -128..8. Units are dBm.
    Cvcallhistorynoiselevel interface{}

    // The object contains the average ACOM level for the call leg. The value -1
    // indicates the level can not be determined or level detection is disabled.
    // The type is interface{} with range: -1..127. Units are dB.
    Cvcallhistoryacomlevel interface{}

    // The object specifies the session target of the peer that is used for the
    // call leg via telephony interface. The type is string with length: 0..64.
    Cvcallhistorysessiontarget interface{}

    // The number of FAX related image pages are received or transmitted via the
    // peer for the call leg. The type is interface{} with range: 0..4294967295.
    // Units are pages.
    Cvcallhistoryimgpagecount interface{}

    // The calling party name of the call. If the name is not available, then it
    // will have a length of zero. The type is string.
    Cvcallhistorycallingname interface{}

    // The object indicates whether or not the caller ID feature is blocked for
    // this call. The type is bool.
    Cvcallhistorycalleridblock interface{}

    // The object indicates the account code input to the call. It could be used
    // by down stream billing server. The value of empty string indicates no
    // account code input. The type is string with length: 0..50.
    Cvcallhistoryaccountcode interface{}

    // This object represents the call identifier for the telephony leg, which was
    // assigned to the call. The type is interface{} with range: 1..4294967295.
    Cvcallhistorycallid interface{}
}

func (cvcallhistoryentry *CISCOVOICEDIALCONTROLMIB_Cvcallhistorytable_Cvcallhistoryentry) GetFilter() yfilter.YFilter { return cvcallhistoryentry.YFilter }

func (cvcallhistoryentry *CISCOVOICEDIALCONTROLMIB_Cvcallhistorytable_Cvcallhistoryentry) SetFilter(yf yfilter.YFilter) { cvcallhistoryentry.YFilter = yf }

func (cvcallhistoryentry *CISCOVOICEDIALCONTROLMIB_Cvcallhistorytable_Cvcallhistoryentry) GetGoName(yname string) string {
    if yname == "cCallHistoryIndex" { return "Ccallhistoryindex" }
    if yname == "cvCallHistoryConnectionId" { return "Cvcallhistoryconnectionid" }
    if yname == "cvCallHistoryTxDuration" { return "Cvcallhistorytxduration" }
    if yname == "cvCallHistoryVoiceTxDuration" { return "Cvcallhistoryvoicetxduration" }
    if yname == "cvCallHistoryFaxTxDuration" { return "Cvcallhistoryfaxtxduration" }
    if yname == "cvCallHistoryCoderTypeRate" { return "Cvcallhistorycodertyperate" }
    if yname == "cvCallHistoryNoiseLevel" { return "Cvcallhistorynoiselevel" }
    if yname == "cvCallHistoryACOMLevel" { return "Cvcallhistoryacomlevel" }
    if yname == "cvCallHistorySessionTarget" { return "Cvcallhistorysessiontarget" }
    if yname == "cvCallHistoryImgPageCount" { return "Cvcallhistoryimgpagecount" }
    if yname == "cvCallHistoryCallingName" { return "Cvcallhistorycallingname" }
    if yname == "cvCallHistoryCallerIDBlock" { return "Cvcallhistorycalleridblock" }
    if yname == "cvCallHistoryAccountCode" { return "Cvcallhistoryaccountcode" }
    if yname == "cvCallHistoryCallId" { return "Cvcallhistorycallid" }
    return ""
}

func (cvcallhistoryentry *CISCOVOICEDIALCONTROLMIB_Cvcallhistorytable_Cvcallhistoryentry) GetSegmentPath() string {
    return "cvCallHistoryEntry" + "[cCallHistoryIndex='" + fmt.Sprintf("%v", cvcallhistoryentry.Ccallhistoryindex) + "']"
}

func (cvcallhistoryentry *CISCOVOICEDIALCONTROLMIB_Cvcallhistorytable_Cvcallhistoryentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cvcallhistoryentry *CISCOVOICEDIALCONTROLMIB_Cvcallhistorytable_Cvcallhistoryentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cvcallhistoryentry *CISCOVOICEDIALCONTROLMIB_Cvcallhistorytable_Cvcallhistoryentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cCallHistoryIndex"] = cvcallhistoryentry.Ccallhistoryindex
    leafs["cvCallHistoryConnectionId"] = cvcallhistoryentry.Cvcallhistoryconnectionid
    leafs["cvCallHistoryTxDuration"] = cvcallhistoryentry.Cvcallhistorytxduration
    leafs["cvCallHistoryVoiceTxDuration"] = cvcallhistoryentry.Cvcallhistoryvoicetxduration
    leafs["cvCallHistoryFaxTxDuration"] = cvcallhistoryentry.Cvcallhistoryfaxtxduration
    leafs["cvCallHistoryCoderTypeRate"] = cvcallhistoryentry.Cvcallhistorycodertyperate
    leafs["cvCallHistoryNoiseLevel"] = cvcallhistoryentry.Cvcallhistorynoiselevel
    leafs["cvCallHistoryACOMLevel"] = cvcallhistoryentry.Cvcallhistoryacomlevel
    leafs["cvCallHistorySessionTarget"] = cvcallhistoryentry.Cvcallhistorysessiontarget
    leafs["cvCallHistoryImgPageCount"] = cvcallhistoryentry.Cvcallhistoryimgpagecount
    leafs["cvCallHistoryCallingName"] = cvcallhistoryentry.Cvcallhistorycallingname
    leafs["cvCallHistoryCallerIDBlock"] = cvcallhistoryentry.Cvcallhistorycalleridblock
    leafs["cvCallHistoryAccountCode"] = cvcallhistoryentry.Cvcallhistoryaccountcode
    leafs["cvCallHistoryCallId"] = cvcallhistoryentry.Cvcallhistorycallid
    return leafs
}

func (cvcallhistoryentry *CISCOVOICEDIALCONTROLMIB_Cvcallhistorytable_Cvcallhistoryentry) GetBundleName() string { return "cisco_ios_xe" }

func (cvcallhistoryentry *CISCOVOICEDIALCONTROLMIB_Cvcallhistorytable_Cvcallhistoryentry) GetYangName() string { return "cvCallHistoryEntry" }

func (cvcallhistoryentry *CISCOVOICEDIALCONTROLMIB_Cvcallhistorytable_Cvcallhistoryentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvcallhistoryentry *CISCOVOICEDIALCONTROLMIB_Cvcallhistorytable_Cvcallhistoryentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvcallhistoryentry *CISCOVOICEDIALCONTROLMIB_Cvcallhistorytable_Cvcallhistoryentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvcallhistoryentry *CISCOVOICEDIALCONTROLMIB_Cvcallhistorytable_Cvcallhistoryentry) SetParent(parent types.Entity) { cvcallhistoryentry.parent = parent }

func (cvcallhistoryentry *CISCOVOICEDIALCONTROLMIB_Cvcallhistorytable_Cvcallhistoryentry) GetParent() types.Entity { return cvcallhistoryentry.parent }

func (cvcallhistoryentry *CISCOVOICEDIALCONTROLMIB_Cvcallhistorytable_Cvcallhistoryentry) GetParentYangName() string { return "cvCallHistoryTable" }

// CISCOVOICEDIALCONTROLMIB_Cvvoipcallhistorytable
// This table is the VoIP extension to the call history table
// of IETF Dial Control MIB. It contains VoIP call leg
// information about specific VoIP call destination and the
// selected QoS for the call leg.
type CISCOVOICEDIALCONTROLMIB_Cvvoipcallhistorytable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The information regarding a single VoIP call leg. The call leg entry is
    // identified by using the same index objects that are used by Call Active
    // table of IETF Dial Control MIB to identify the call. An entry of this table
    // is created when its associated call history entry in the IETF Dial Control
    // MIB is created and the call history entry contains information for the call
    // establishment to the peer on the IP backbone via a voice over IP peer. The
    // entry is deleted when its associated call history entry in the IETF Dial
    // Control MIB is deleted. The type is slice of
    // CISCOVOICEDIALCONTROLMIB_Cvvoipcallhistorytable_Cvvoipcallhistoryentry.
    Cvvoipcallhistoryentry []CISCOVOICEDIALCONTROLMIB_Cvvoipcallhistorytable_Cvvoipcallhistoryentry
}

func (cvvoipcallhistorytable *CISCOVOICEDIALCONTROLMIB_Cvvoipcallhistorytable) GetFilter() yfilter.YFilter { return cvvoipcallhistorytable.YFilter }

func (cvvoipcallhistorytable *CISCOVOICEDIALCONTROLMIB_Cvvoipcallhistorytable) SetFilter(yf yfilter.YFilter) { cvvoipcallhistorytable.YFilter = yf }

func (cvvoipcallhistorytable *CISCOVOICEDIALCONTROLMIB_Cvvoipcallhistorytable) GetGoName(yname string) string {
    if yname == "cvVoIPCallHistoryEntry" { return "Cvvoipcallhistoryentry" }
    return ""
}

func (cvvoipcallhistorytable *CISCOVOICEDIALCONTROLMIB_Cvvoipcallhistorytable) GetSegmentPath() string {
    return "cvVoIPCallHistoryTable"
}

func (cvvoipcallhistorytable *CISCOVOICEDIALCONTROLMIB_Cvvoipcallhistorytable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cvVoIPCallHistoryEntry" {
        for _, c := range cvvoipcallhistorytable.Cvvoipcallhistoryentry {
            if cvvoipcallhistorytable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOVOICEDIALCONTROLMIB_Cvvoipcallhistorytable_Cvvoipcallhistoryentry{}
        cvvoipcallhistorytable.Cvvoipcallhistoryentry = append(cvvoipcallhistorytable.Cvvoipcallhistoryentry, child)
        return &cvvoipcallhistorytable.Cvvoipcallhistoryentry[len(cvvoipcallhistorytable.Cvvoipcallhistoryentry)-1]
    }
    return nil
}

func (cvvoipcallhistorytable *CISCOVOICEDIALCONTROLMIB_Cvvoipcallhistorytable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cvvoipcallhistorytable.Cvvoipcallhistoryentry {
        children[cvvoipcallhistorytable.Cvvoipcallhistoryentry[i].GetSegmentPath()] = &cvvoipcallhistorytable.Cvvoipcallhistoryentry[i]
    }
    return children
}

func (cvvoipcallhistorytable *CISCOVOICEDIALCONTROLMIB_Cvvoipcallhistorytable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cvvoipcallhistorytable *CISCOVOICEDIALCONTROLMIB_Cvvoipcallhistorytable) GetBundleName() string { return "cisco_ios_xe" }

func (cvvoipcallhistorytable *CISCOVOICEDIALCONTROLMIB_Cvvoipcallhistorytable) GetYangName() string { return "cvVoIPCallHistoryTable" }

func (cvvoipcallhistorytable *CISCOVOICEDIALCONTROLMIB_Cvvoipcallhistorytable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvvoipcallhistorytable *CISCOVOICEDIALCONTROLMIB_Cvvoipcallhistorytable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvvoipcallhistorytable *CISCOVOICEDIALCONTROLMIB_Cvvoipcallhistorytable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvvoipcallhistorytable *CISCOVOICEDIALCONTROLMIB_Cvvoipcallhistorytable) SetParent(parent types.Entity) { cvvoipcallhistorytable.parent = parent }

func (cvvoipcallhistorytable *CISCOVOICEDIALCONTROLMIB_Cvvoipcallhistorytable) GetParent() types.Entity { return cvvoipcallhistorytable.parent }

func (cvvoipcallhistorytable *CISCOVOICEDIALCONTROLMIB_Cvvoipcallhistorytable) GetParentYangName() string { return "CISCO-VOICE-DIAL-CONTROL-MIB" }

// CISCOVOICEDIALCONTROLMIB_Cvvoipcallhistorytable_Cvvoipcallhistoryentry
// The information regarding a single VoIP call leg.
// The call leg entry is identified by using the same index
// objects that are used by Call Active table of IETF Dial
// Control MIB to identify the call.
// An entry of this table is created when its associated call
// history entry in the IETF Dial Control MIB is created and
// the call history entry contains information for the call
// establishment to the peer on the IP backbone via a voice
// over IP peer.
// The entry is deleted when its associated call history
// entry in the IETF Dial Control MIB is deleted.
type CISCOVOICEDIALCONTROLMIB_Cvvoipcallhistorytable_Cvvoipcallhistoryentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // cisco_dial_control_mib.CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryindex
    Ccallhistoryindex interface{}

    // The global connection identifier for the VoIP leg, which was assigned to
    // the call. The type is string with length: 0..16.
    Cvvoipcallhistoryconnectionid interface{}

    // Remote system IP address for the call. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Cvvoipcallhistoryremoteipaddress interface{}

    // Remote system UDP listener port to which to transmit voice packets for the
    // call. The type is interface{} with range: 0..65535.
    Cvvoipcallhistoryremoteudpport interface{}

    // The voice packet round trip delay between local and the remote system on
    // the IP backbone during the call. The type is interface{} with range:
    // 0..4294967295. Units are milliseconds.
    Cvvoipcallhistoryroundtripdelay interface{}

    // The selected RSVP QoS for the call. The type is QosService.
    Cvvoipcallhistoryselectedqos interface{}

    // The object specifies the session protocol to be used for Internet call
    // between local and remote router via IP backbone. The type is
    // CvSessionProtocol.
    Cvvoipcallhistorysessionprotocol interface{}

    // The object specifies the session target of the peer that is used for the
    // Voice over IP call. The type is string.
    Cvvoipcallhistorysessiontarget interface{}

    // Duration of voice playout from data received on time for this call. This
    // plus the durations for the GapFills in the following entries gives the
    // Total Voice Playout Duration for Active Voice. This does not include
    // duration for which no data was sent by the Transmit end as voice signal,
    // e.g., silence suppression and fax signal. The On Time Playout Rate can be
    // computed by dividing this entry by the Total Voice Playout Duration. This
    // counter object will lock at the maximum value which is approximately two
    // days. The type is interface{} with range: 0..4294967295. Units are
    // milliseconds.
    Cvvoipcallhistoryontimervplayout interface{}

    // Duration of voice signal replaced with signal played out during silence due
    // to voice data not received on time (or lost) from voice gateway this call.
    // This counter object will lock at the maximum value which is approximately
    // two days. The type is interface{} with range: 0..4294967295. Units are
    // milliseconds.
    Cvvoipcallhistorygapfillwithsilence interface{}

    // Duration of voice signal played out with signal synthesized from parameters
    // or samples of data preceding in time due to voice data not received on time
    // (or lost) from voice gateway for this call. An example of such playout is
    // frame-erasure or  frame-concealment strategies in G.729 and G.723.1
    // compression algorithms. This counter object will lock at the maximum value
    // which is approximately two days. The type is interface{} with range:
    // 0..4294967295. Units are milliseconds.
    Cvvoipcallhistorygapfillwithprediction interface{}

    // Duration of voice signal played out with signal synthesized from parameters
    // or samples of data preceding and following in time due to voice data not
    // received on time (or lost) from voice gateway for this call. This counter
    // object will lock at the maximum value which is approximately two days. The
    // type is interface{} with range: 0..4294967295. Units are milliseconds.
    Cvvoipcallhistorygapfillwithinterpolation interface{}

    // Duration of voice signal played out with signal synthesized from redundancy
    // parameters available due to voice data not received on time (or lost) from
    // voice gateway for this call. This counter object will lock at the maximum
    // value which is approximately two days. The type is interface{} with range:
    // 0..4294967295. Units are milliseconds.
    Cvvoipcallhistorygapfillwithredundancy interface{}

    // The high water mark Voice Playout FIFO Delay during the voice call. This
    // counter object will lock at the maximum value which is approximately two
    // days. The type is interface{} with range: 0..4294967295. Units are
    // milliseconds.
    Cvvoipcallhistoryhiwaterplayoutdelay interface{}

    // The low water mark Voice Playout FIFO Delay during the voice call. The type
    // is interface{} with range: 0..4294967295. Units are milliseconds.
    Cvvoipcallhistorylowaterplayoutdelay interface{}

    // The average Playout FIFO Delay plus the decoder delay during the voice
    // call. The type is interface{} with range: 0..4294967295.
    Cvvoipcallhistoryreceivedelay interface{}

    // The object indicates whether or not the VAD (Voice Activity Detection) was
    // enabled for the voice call. The type is bool.
    Cvvoipcallhistoryvadenable interface{}

    // The negotiated coder rate. It specifies the transmit rate of voice/fax
    // compression to its associated call leg for the call. This rate is different
    // from the configuration rate because of rate negotiation during the call.
    // The type is CvcCoderTypeRate.
    Cvvoipcallhistorycodertyperate interface{}

    // The Calculated Planning Impairment Factor (Icpif) of the call  that is
    // associated to this call leg. The value in this object is computed by the
    // following equation. Icpif of the call = Itotal (total impairment value) of
    // the call - A (Expectation Factor) in the cvVoIPPeerCfgExpectFactor of the
    // call leg associated peer. A value of -1 implies that Icpif was not
    // calculated and is meaningless for this call. The type is interface{} with
    // range: -1..55. Units are equipment impairment factor (eif).
    Cvvoipcallhistoryicpif interface{}

    // The number of lost voice packets during the call. The type is interface{}
    // with range: 0..4294967295. Units are packets.
    Cvvoipcallhistorylostpackets interface{}

    // The number of received voice packets that are arrived too early to store in
    // jitter buffer during the call. The type is interface{} with range:
    // 0..4294967295. Units are packets.
    Cvvoipcallhistoryearlypackets interface{}

    // The number of received voice packets that are arrived too late to playout
    // with CODEC (Coder/Decoder) during the call. The type is interface{} with
    // range: 0..4294967295. Units are packets.
    Cvvoipcallhistorylatepackets interface{}

    // The textual identifier of the calling party (user) of the call. If the
    // username is not available, then the value of this object will have a length
    // of zero. The type is string with length: 0..16.
    Cvvoipcallhistoryusername interface{}

    // The protocol-specific call identifier for the VoIP call. The type is string
    // with length: 0..255.
    Cvvoipcallhistoryprotocolcallid interface{}

    // This object specifies the type of address contained in the associated
    // instance of cvVoIPCallHistoryRemSigIPAddr. The type is InetAddressType.
    Cvvoipcallhistoryremsigipaddrt interface{}

    // Remote signalling IP address for the VoIP call. The type is string with
    // length: 0..255.
    Cvvoipcallhistoryremsigipaddr interface{}

    // Remote signalling listener port to which to transmit voice packets. The
    // type is interface{} with range: 0..65535.
    Cvvoipcallhistoryremsigport interface{}

    // This object specifies the type of address contained in the associated
    // instance of cvVoIPCallHistoryRemMediaIPAddr. The type is InetAddressType.
    Cvvoipcallhistoryremmediaipaddrt interface{}

    // Remote media end point IP address for the VoIP call. The type is string
    // with length: 0..255.
    Cvvoipcallhistoryremmediaipaddr interface{}

    // Remote media end point listener port to which to transmit voice packets.
    // The type is interface{} with range: 0..65535.
    Cvvoipcallhistoryremmediaport interface{}

    // The object indicates whether or not the SRTP (Secured RTP) was enabled for
    // the voice call. The type is bool.
    Cvvoipcallhistorysrtpenable interface{}

    // The Calculated Planning Impairment Factor (Icpif) of the call  that is
    // associated to this call leg. The value in this object is computed by the
    // following equation. Icpif of the fallback probe = Itotal (total impairment
    // value)  - configured fallback (Expectation Factor). A value of 0 implies
    // that Icpif was not calculated and is meaningless for this attempt. The type
    // is interface{} with range: -2147483648..2147483647.
    Cvvoipcallhistoryfallbackicpif interface{}

    // FallbackLoss is the percentage of loss packets based on the total packets
    // sent. The type is interface{} with range: 0..4294967295.
    Cvvoipcallhistoryfallbackloss interface{}

    // The FallbackDelay is calculated as follows - Take the sum of the round
    // trips for all the probes,  divide by the number of probes,  and divide by
    // two to get the one-way delay.   Then add in jitter_in or jiter_out, which
    // ever is higher. The type is interface{} with range: 0..4294967295.
    Cvvoipcallhistoryfallbackdelay interface{}

    // If the object has a value true(1) octet align operation is used, and if the
    // value is false(2), bandwidth efficient operation is used. This object is
    // not instantiated when the object cvVoIPCallHistoryCoderTypeRate is not
    // equal to gsmAmrNb enum. The type is bool.
    Cvvoipcallhistoryoctetaligned interface{}

    // This object indicates modes of Bit rates. This object is not instantiated
    // when the object cvVoIPCallHistoryCoderTypeRate is not equal to gsmAmrNb
    // enum. The type is map[string]bool.
    Cvvoipcallhistorybitrates interface{}

    // The object indicates the interval (N frame-blocks) at which codec mode
    // changes are allowed. This object is not instantiated when the object
    // cvVoIPCallHistoryCoderTypeRate is not equal to gsmAmrNb enum. The type is
    // interface{} with range: 1..100. Units are frame-blocks.
    Cvvoipcallhistorymodechgperiod interface{}

    // If the object has a value of true(1), mode changes will be made to only
    // neighboring modes set to cvVoIPCallHistoryBitRates object. If the value is
    // false(2), mode changes will be allowed to any modes set to
    // cvVoIPCallHistoryBitRates object. This object is not instantiated when the
    // object cvVoIPCallHistoryCoderTypeRate is not equal to gsmAmrNb enum. The
    // type is bool.
    Cvvoipcallhistorymodechgneighbor interface{}

    // The object indicates the maximum amount of media that can be encapsulated
    // in a payload. Supported value is 20 msec. This object is not instantiated
    // when the object cvVoIPCallHistoryCoderTypeRate is not equal to gsmAmrNb
    // enum. The type is interface{} with range: 20..100. Units are milliseconds.
    Cvvoipcallhistorymaxptime interface{}

    // If the object has a value of true(1), frame CRC will be included in the
    // payload and if the value is false(2), frame CRC will not be included in the
    // payload. This object is applicable only when RTP frame type is octet
    // aligned. This object is not instantiated when the object
    // cvVoIPCallHistoryCoderTypeRate is not equal to gsmAmrNb enum. The type is
    // bool.
    Cvvoipcallhistorycrc interface{}

    // If the object has a value of true(1), payload employs robust sorting and if
    // the value is false(2), payload does not employ robust sorting. This object
    // is applicable only when RTP frame type is octet aligned. This object is not
    // instantiated when the object cvVoIPCallHistoryCoderTypeRate is not equal to
    // gsmAmrNb enum. The type is bool.
    Cvvoipcallhistoryrobustsorting interface{}

    // The object indicates the RTP encapsulation type. Supported RTP
    // encapsulation type is RFC3267. This object is not instantiated when the
    // object cvVoIPCallHistoryCoderTypeRate is not equal to gsmAmrNb enum. The
    // type is CvAmrNbRtpEncap.
    Cvvoipcallhistoryencap interface{}

    // The object indicates the maximum number of frame-blocks allowed in an
    // interleaving group. It indicates that frame-block level interleaving will
    // be used for that session. If this object is not set, interleaving is not
    // used. This object is applicable only when RTP frame type is octet aligned.
    // This object is not instantiated when the object
    // cvVoIPCallHistoryCoderTypeRate is not equal to gsmAmrNb enum. The type is
    // interface{} with range: 1..50. Units are frame-blocks.
    Cvvoipcallhistoryinterleaving interface{}

    // The object indicates the length of the time in milliseconds represented by
    // the media of the packet. Supported value is 20 milliseconds. This object is
    // not instantiated when the object cvVoIPCallHistoryCoderTypeRate is not
    // equal to gsmAmrNb enum. The type is interface{} with range: 20..100. Units
    // are milliseconds.
    Cvvoipcallhistoryptime interface{}

    // The object indicates the number of audio channels. Supported value is 1.
    // This object is not instantiated when the object
    // cvVoIPCallHistoryCoderTypeRate is not equal to gsmAmrNb enum. The type is
    // interface{} with range: 1..6. Units are channels.
    Cvvoipcallhistorychannels interface{}

    // The object indicates the iLBC mode. The value of this object is valid only
    // if  cvVoIPCallHistoryCoderTypeRate is equal to  'iLBC'. The type is
    // CvIlbcFrameMode. Units are milliseconds.
    Cvvoipcallhistorycodermode interface{}

    // This object represents the call identifier for the VoIP leg, which was
    // assigned to the call. The type is interface{} with range: 1..4294967295.
    Cvvoipcallhistorycallid interface{}

    // The call reference ID associates the video call entry and voice call entry
    // of the same endpoint.  An audio-only call may or may not have a valid call
    // reference ID (i.e. value greater than zero), but in both cases, there will
    // not be a video call entry associated with it.   For a video call, the
    // video-specific information  is stored in a call entry in
    // cVideoSessionActive of CISCO-VIDEO-SESSION-MIB, in which the call reference
    // ID is also identified. The type is interface{} with range: 0..4294967295.
    Cvvoipcallhistorycallreferenceid interface{}

    // This object indicates the session ID assigned by the call manager to
    // identify call legs that belong to the same call session.  This session ID
    // (history) represents a completed call session, whereas the active session
    // ID (cvVoIPCallActiveSessionId) represents an ongoing session. The type is
    // interface{} with range: 0..4294967295.
    Cvvoipcallhistorysessionid interface{}
}

func (cvvoipcallhistoryentry *CISCOVOICEDIALCONTROLMIB_Cvvoipcallhistorytable_Cvvoipcallhistoryentry) GetFilter() yfilter.YFilter { return cvvoipcallhistoryentry.YFilter }

func (cvvoipcallhistoryentry *CISCOVOICEDIALCONTROLMIB_Cvvoipcallhistorytable_Cvvoipcallhistoryentry) SetFilter(yf yfilter.YFilter) { cvvoipcallhistoryentry.YFilter = yf }

func (cvvoipcallhistoryentry *CISCOVOICEDIALCONTROLMIB_Cvvoipcallhistorytable_Cvvoipcallhistoryentry) GetGoName(yname string) string {
    if yname == "cCallHistoryIndex" { return "Ccallhistoryindex" }
    if yname == "cvVoIPCallHistoryConnectionId" { return "Cvvoipcallhistoryconnectionid" }
    if yname == "cvVoIPCallHistoryRemoteIPAddress" { return "Cvvoipcallhistoryremoteipaddress" }
    if yname == "cvVoIPCallHistoryRemoteUDPPort" { return "Cvvoipcallhistoryremoteudpport" }
    if yname == "cvVoIPCallHistoryRoundTripDelay" { return "Cvvoipcallhistoryroundtripdelay" }
    if yname == "cvVoIPCallHistorySelectedQoS" { return "Cvvoipcallhistoryselectedqos" }
    if yname == "cvVoIPCallHistorySessionProtocol" { return "Cvvoipcallhistorysessionprotocol" }
    if yname == "cvVoIPCallHistorySessionTarget" { return "Cvvoipcallhistorysessiontarget" }
    if yname == "cvVoIPCallHistoryOnTimeRvPlayout" { return "Cvvoipcallhistoryontimervplayout" }
    if yname == "cvVoIPCallHistoryGapFillWithSilence" { return "Cvvoipcallhistorygapfillwithsilence" }
    if yname == "cvVoIPCallHistoryGapFillWithPrediction" { return "Cvvoipcallhistorygapfillwithprediction" }
    if yname == "cvVoIPCallHistoryGapFillWithInterpolation" { return "Cvvoipcallhistorygapfillwithinterpolation" }
    if yname == "cvVoIPCallHistoryGapFillWithRedundancy" { return "Cvvoipcallhistorygapfillwithredundancy" }
    if yname == "cvVoIPCallHistoryHiWaterPlayoutDelay" { return "Cvvoipcallhistoryhiwaterplayoutdelay" }
    if yname == "cvVoIPCallHistoryLoWaterPlayoutDelay" { return "Cvvoipcallhistorylowaterplayoutdelay" }
    if yname == "cvVoIPCallHistoryReceiveDelay" { return "Cvvoipcallhistoryreceivedelay" }
    if yname == "cvVoIPCallHistoryVADEnable" { return "Cvvoipcallhistoryvadenable" }
    if yname == "cvVoIPCallHistoryCoderTypeRate" { return "Cvvoipcallhistorycodertyperate" }
    if yname == "cvVoIPCallHistoryIcpif" { return "Cvvoipcallhistoryicpif" }
    if yname == "cvVoIPCallHistoryLostPackets" { return "Cvvoipcallhistorylostpackets" }
    if yname == "cvVoIPCallHistoryEarlyPackets" { return "Cvvoipcallhistoryearlypackets" }
    if yname == "cvVoIPCallHistoryLatePackets" { return "Cvvoipcallhistorylatepackets" }
    if yname == "cvVoIPCallHistoryUsername" { return "Cvvoipcallhistoryusername" }
    if yname == "cvVoIPCallHistoryProtocolCallId" { return "Cvvoipcallhistoryprotocolcallid" }
    if yname == "cvVoIPCallHistoryRemSigIPAddrT" { return "Cvvoipcallhistoryremsigipaddrt" }
    if yname == "cvVoIPCallHistoryRemSigIPAddr" { return "Cvvoipcallhistoryremsigipaddr" }
    if yname == "cvVoIPCallHistoryRemSigPort" { return "Cvvoipcallhistoryremsigport" }
    if yname == "cvVoIPCallHistoryRemMediaIPAddrT" { return "Cvvoipcallhistoryremmediaipaddrt" }
    if yname == "cvVoIPCallHistoryRemMediaIPAddr" { return "Cvvoipcallhistoryremmediaipaddr" }
    if yname == "cvVoIPCallHistoryRemMediaPort" { return "Cvvoipcallhistoryremmediaport" }
    if yname == "cvVoIPCallHistorySRTPEnable" { return "Cvvoipcallhistorysrtpenable" }
    if yname == "cvVoIPCallHistoryFallbackIcpif" { return "Cvvoipcallhistoryfallbackicpif" }
    if yname == "cvVoIPCallHistoryFallbackLoss" { return "Cvvoipcallhistoryfallbackloss" }
    if yname == "cvVoIPCallHistoryFallbackDelay" { return "Cvvoipcallhistoryfallbackdelay" }
    if yname == "cvVoIPCallHistoryOctetAligned" { return "Cvvoipcallhistoryoctetaligned" }
    if yname == "cvVoIPCallHistoryBitRates" { return "Cvvoipcallhistorybitrates" }
    if yname == "cvVoIPCallHistoryModeChgPeriod" { return "Cvvoipcallhistorymodechgperiod" }
    if yname == "cvVoIPCallHistoryModeChgNeighbor" { return "Cvvoipcallhistorymodechgneighbor" }
    if yname == "cvVoIPCallHistoryMaxPtime" { return "Cvvoipcallhistorymaxptime" }
    if yname == "cvVoIPCallHistoryCRC" { return "Cvvoipcallhistorycrc" }
    if yname == "cvVoIPCallHistoryRobustSorting" { return "Cvvoipcallhistoryrobustsorting" }
    if yname == "cvVoIPCallHistoryEncap" { return "Cvvoipcallhistoryencap" }
    if yname == "cvVoIPCallHistoryInterleaving" { return "Cvvoipcallhistoryinterleaving" }
    if yname == "cvVoIPCallHistoryPtime" { return "Cvvoipcallhistoryptime" }
    if yname == "cvVoIPCallHistoryChannels" { return "Cvvoipcallhistorychannels" }
    if yname == "cvVoIPCallHistoryCoderMode" { return "Cvvoipcallhistorycodermode" }
    if yname == "cvVoIPCallHistoryCallId" { return "Cvvoipcallhistorycallid" }
    if yname == "cvVoIPCallHistoryCallReferenceId" { return "Cvvoipcallhistorycallreferenceid" }
    if yname == "cvVoIPCallHistorySessionId" { return "Cvvoipcallhistorysessionid" }
    return ""
}

func (cvvoipcallhistoryentry *CISCOVOICEDIALCONTROLMIB_Cvvoipcallhistorytable_Cvvoipcallhistoryentry) GetSegmentPath() string {
    return "cvVoIPCallHistoryEntry" + "[cCallHistoryIndex='" + fmt.Sprintf("%v", cvvoipcallhistoryentry.Ccallhistoryindex) + "']"
}

func (cvvoipcallhistoryentry *CISCOVOICEDIALCONTROLMIB_Cvvoipcallhistorytable_Cvvoipcallhistoryentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cvvoipcallhistoryentry *CISCOVOICEDIALCONTROLMIB_Cvvoipcallhistorytable_Cvvoipcallhistoryentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cvvoipcallhistoryentry *CISCOVOICEDIALCONTROLMIB_Cvvoipcallhistorytable_Cvvoipcallhistoryentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cCallHistoryIndex"] = cvvoipcallhistoryentry.Ccallhistoryindex
    leafs["cvVoIPCallHistoryConnectionId"] = cvvoipcallhistoryentry.Cvvoipcallhistoryconnectionid
    leafs["cvVoIPCallHistoryRemoteIPAddress"] = cvvoipcallhistoryentry.Cvvoipcallhistoryremoteipaddress
    leafs["cvVoIPCallHistoryRemoteUDPPort"] = cvvoipcallhistoryentry.Cvvoipcallhistoryremoteudpport
    leafs["cvVoIPCallHistoryRoundTripDelay"] = cvvoipcallhistoryentry.Cvvoipcallhistoryroundtripdelay
    leafs["cvVoIPCallHistorySelectedQoS"] = cvvoipcallhistoryentry.Cvvoipcallhistoryselectedqos
    leafs["cvVoIPCallHistorySessionProtocol"] = cvvoipcallhistoryentry.Cvvoipcallhistorysessionprotocol
    leafs["cvVoIPCallHistorySessionTarget"] = cvvoipcallhistoryentry.Cvvoipcallhistorysessiontarget
    leafs["cvVoIPCallHistoryOnTimeRvPlayout"] = cvvoipcallhistoryentry.Cvvoipcallhistoryontimervplayout
    leafs["cvVoIPCallHistoryGapFillWithSilence"] = cvvoipcallhistoryentry.Cvvoipcallhistorygapfillwithsilence
    leafs["cvVoIPCallHistoryGapFillWithPrediction"] = cvvoipcallhistoryentry.Cvvoipcallhistorygapfillwithprediction
    leafs["cvVoIPCallHistoryGapFillWithInterpolation"] = cvvoipcallhistoryentry.Cvvoipcallhistorygapfillwithinterpolation
    leafs["cvVoIPCallHistoryGapFillWithRedundancy"] = cvvoipcallhistoryentry.Cvvoipcallhistorygapfillwithredundancy
    leafs["cvVoIPCallHistoryHiWaterPlayoutDelay"] = cvvoipcallhistoryentry.Cvvoipcallhistoryhiwaterplayoutdelay
    leafs["cvVoIPCallHistoryLoWaterPlayoutDelay"] = cvvoipcallhistoryentry.Cvvoipcallhistorylowaterplayoutdelay
    leafs["cvVoIPCallHistoryReceiveDelay"] = cvvoipcallhistoryentry.Cvvoipcallhistoryreceivedelay
    leafs["cvVoIPCallHistoryVADEnable"] = cvvoipcallhistoryentry.Cvvoipcallhistoryvadenable
    leafs["cvVoIPCallHistoryCoderTypeRate"] = cvvoipcallhistoryentry.Cvvoipcallhistorycodertyperate
    leafs["cvVoIPCallHistoryIcpif"] = cvvoipcallhistoryentry.Cvvoipcallhistoryicpif
    leafs["cvVoIPCallHistoryLostPackets"] = cvvoipcallhistoryentry.Cvvoipcallhistorylostpackets
    leafs["cvVoIPCallHistoryEarlyPackets"] = cvvoipcallhistoryentry.Cvvoipcallhistoryearlypackets
    leafs["cvVoIPCallHistoryLatePackets"] = cvvoipcallhistoryentry.Cvvoipcallhistorylatepackets
    leafs["cvVoIPCallHistoryUsername"] = cvvoipcallhistoryentry.Cvvoipcallhistoryusername
    leafs["cvVoIPCallHistoryProtocolCallId"] = cvvoipcallhistoryentry.Cvvoipcallhistoryprotocolcallid
    leafs["cvVoIPCallHistoryRemSigIPAddrT"] = cvvoipcallhistoryentry.Cvvoipcallhistoryremsigipaddrt
    leafs["cvVoIPCallHistoryRemSigIPAddr"] = cvvoipcallhistoryentry.Cvvoipcallhistoryremsigipaddr
    leafs["cvVoIPCallHistoryRemSigPort"] = cvvoipcallhistoryentry.Cvvoipcallhistoryremsigport
    leafs["cvVoIPCallHistoryRemMediaIPAddrT"] = cvvoipcallhistoryentry.Cvvoipcallhistoryremmediaipaddrt
    leafs["cvVoIPCallHistoryRemMediaIPAddr"] = cvvoipcallhistoryentry.Cvvoipcallhistoryremmediaipaddr
    leafs["cvVoIPCallHistoryRemMediaPort"] = cvvoipcallhistoryentry.Cvvoipcallhistoryremmediaport
    leafs["cvVoIPCallHistorySRTPEnable"] = cvvoipcallhistoryentry.Cvvoipcallhistorysrtpenable
    leafs["cvVoIPCallHistoryFallbackIcpif"] = cvvoipcallhistoryentry.Cvvoipcallhistoryfallbackicpif
    leafs["cvVoIPCallHistoryFallbackLoss"] = cvvoipcallhistoryentry.Cvvoipcallhistoryfallbackloss
    leafs["cvVoIPCallHistoryFallbackDelay"] = cvvoipcallhistoryentry.Cvvoipcallhistoryfallbackdelay
    leafs["cvVoIPCallHistoryOctetAligned"] = cvvoipcallhistoryentry.Cvvoipcallhistoryoctetaligned
    leafs["cvVoIPCallHistoryBitRates"] = cvvoipcallhistoryentry.Cvvoipcallhistorybitrates
    leafs["cvVoIPCallHistoryModeChgPeriod"] = cvvoipcallhistoryentry.Cvvoipcallhistorymodechgperiod
    leafs["cvVoIPCallHistoryModeChgNeighbor"] = cvvoipcallhistoryentry.Cvvoipcallhistorymodechgneighbor
    leafs["cvVoIPCallHistoryMaxPtime"] = cvvoipcallhistoryentry.Cvvoipcallhistorymaxptime
    leafs["cvVoIPCallHistoryCRC"] = cvvoipcallhistoryentry.Cvvoipcallhistorycrc
    leafs["cvVoIPCallHistoryRobustSorting"] = cvvoipcallhistoryentry.Cvvoipcallhistoryrobustsorting
    leafs["cvVoIPCallHistoryEncap"] = cvvoipcallhistoryentry.Cvvoipcallhistoryencap
    leafs["cvVoIPCallHistoryInterleaving"] = cvvoipcallhistoryentry.Cvvoipcallhistoryinterleaving
    leafs["cvVoIPCallHistoryPtime"] = cvvoipcallhistoryentry.Cvvoipcallhistoryptime
    leafs["cvVoIPCallHistoryChannels"] = cvvoipcallhistoryentry.Cvvoipcallhistorychannels
    leafs["cvVoIPCallHistoryCoderMode"] = cvvoipcallhistoryentry.Cvvoipcallhistorycodermode
    leafs["cvVoIPCallHistoryCallId"] = cvvoipcallhistoryentry.Cvvoipcallhistorycallid
    leafs["cvVoIPCallHistoryCallReferenceId"] = cvvoipcallhistoryentry.Cvvoipcallhistorycallreferenceid
    leafs["cvVoIPCallHistorySessionId"] = cvvoipcallhistoryentry.Cvvoipcallhistorysessionid
    return leafs
}

func (cvvoipcallhistoryentry *CISCOVOICEDIALCONTROLMIB_Cvvoipcallhistorytable_Cvvoipcallhistoryentry) GetBundleName() string { return "cisco_ios_xe" }

func (cvvoipcallhistoryentry *CISCOVOICEDIALCONTROLMIB_Cvvoipcallhistorytable_Cvvoipcallhistoryentry) GetYangName() string { return "cvVoIPCallHistoryEntry" }

func (cvvoipcallhistoryentry *CISCOVOICEDIALCONTROLMIB_Cvvoipcallhistorytable_Cvvoipcallhistoryentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvvoipcallhistoryentry *CISCOVOICEDIALCONTROLMIB_Cvvoipcallhistorytable_Cvvoipcallhistoryentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvvoipcallhistoryentry *CISCOVOICEDIALCONTROLMIB_Cvvoipcallhistorytable_Cvvoipcallhistoryentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvvoipcallhistoryentry *CISCOVOICEDIALCONTROLMIB_Cvvoipcallhistorytable_Cvvoipcallhistoryentry) SetParent(parent types.Entity) { cvvoipcallhistoryentry.parent = parent }

func (cvvoipcallhistoryentry *CISCOVOICEDIALCONTROLMIB_Cvvoipcallhistorytable_Cvvoipcallhistoryentry) GetParent() types.Entity { return cvvoipcallhistoryentry.parent }

func (cvvoipcallhistoryentry *CISCOVOICEDIALCONTROLMIB_Cvvoipcallhistorytable_Cvvoipcallhistoryentry) GetParentYangName() string { return "cvVoIPCallHistoryTable" }

// CISCOVOICEDIALCONTROLMIB_Cvcallratestatstable
// This table represents voice call rate measurement in various
// interval lengths defined by the 
// CvCallVolumeStatsIntvlType object.
// 
// Each interval may contain one or more entries to allow for
// detailed measurement to be collected.
type CISCOVOICEDIALCONTROLMIB_Cvcallratestatstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This is a conceptual-row in cvCallRateStatsTable This entry is created at
    // the system initialization and is updated at every epoch based on
    // CvCallVolumeStatsIntvlType. The type is slice of
    // CISCOVOICEDIALCONTROLMIB_Cvcallratestatstable_Cvcallratestatsentry.
    Cvcallratestatsentry []CISCOVOICEDIALCONTROLMIB_Cvcallratestatstable_Cvcallratestatsentry
}

func (cvcallratestatstable *CISCOVOICEDIALCONTROLMIB_Cvcallratestatstable) GetFilter() yfilter.YFilter { return cvcallratestatstable.YFilter }

func (cvcallratestatstable *CISCOVOICEDIALCONTROLMIB_Cvcallratestatstable) SetFilter(yf yfilter.YFilter) { cvcallratestatstable.YFilter = yf }

func (cvcallratestatstable *CISCOVOICEDIALCONTROLMIB_Cvcallratestatstable) GetGoName(yname string) string {
    if yname == "cvCallRateStatsEntry" { return "Cvcallratestatsentry" }
    return ""
}

func (cvcallratestatstable *CISCOVOICEDIALCONTROLMIB_Cvcallratestatstable) GetSegmentPath() string {
    return "cvCallRateStatsTable"
}

func (cvcallratestatstable *CISCOVOICEDIALCONTROLMIB_Cvcallratestatstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cvCallRateStatsEntry" {
        for _, c := range cvcallratestatstable.Cvcallratestatsentry {
            if cvcallratestatstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOVOICEDIALCONTROLMIB_Cvcallratestatstable_Cvcallratestatsentry{}
        cvcallratestatstable.Cvcallratestatsentry = append(cvcallratestatstable.Cvcallratestatsentry, child)
        return &cvcallratestatstable.Cvcallratestatsentry[len(cvcallratestatstable.Cvcallratestatsentry)-1]
    }
    return nil
}

func (cvcallratestatstable *CISCOVOICEDIALCONTROLMIB_Cvcallratestatstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cvcallratestatstable.Cvcallratestatsentry {
        children[cvcallratestatstable.Cvcallratestatsentry[i].GetSegmentPath()] = &cvcallratestatstable.Cvcallratestatsentry[i]
    }
    return children
}

func (cvcallratestatstable *CISCOVOICEDIALCONTROLMIB_Cvcallratestatstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cvcallratestatstable *CISCOVOICEDIALCONTROLMIB_Cvcallratestatstable) GetBundleName() string { return "cisco_ios_xe" }

func (cvcallratestatstable *CISCOVOICEDIALCONTROLMIB_Cvcallratestatstable) GetYangName() string { return "cvCallRateStatsTable" }

func (cvcallratestatstable *CISCOVOICEDIALCONTROLMIB_Cvcallratestatstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvcallratestatstable *CISCOVOICEDIALCONTROLMIB_Cvcallratestatstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvcallratestatstable *CISCOVOICEDIALCONTROLMIB_Cvcallratestatstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvcallratestatstable *CISCOVOICEDIALCONTROLMIB_Cvcallratestatstable) SetParent(parent types.Entity) { cvcallratestatstable.parent = parent }

func (cvcallratestatstable *CISCOVOICEDIALCONTROLMIB_Cvcallratestatstable) GetParent() types.Entity { return cvcallratestatstable.parent }

func (cvcallratestatstable *CISCOVOICEDIALCONTROLMIB_Cvcallratestatstable) GetParentYangName() string { return "CISCO-VOICE-DIAL-CONTROL-MIB" }

// CISCOVOICEDIALCONTROLMIB_Cvcallratestatstable_Cvcallratestatsentry
// This is a conceptual-row in cvCallRateStatsTable
// This entry is created at the system initialization and is
// updated at every epoch based on CvCallVolumeStatsIntvlType
type CISCOVOICEDIALCONTROLMIB_Cvcallratestatstable_Cvcallratestatsentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The Object indexes in Call Rate Table to select
    // one among three interval-tables.  The different types in this table are
    // represented by  CvCallVolumeStatsIntvlType. The type is
    // CvCallVolumeStatsIntvlType.
    Cvcallratestatsintvldurunits interface{}

    // This attribute is a key. This is an index that references to the different
    // past periods in given in interval of call rate table. This range is 1-60
    // for Seconds and Minutes table  wherein 1-72 for hours table. The type is
    // interface{} with range: 1..72.
    Cvcallratestatsintvldur interface{}

    // This object indicates the maximum calls per second that occured for the
    // given period for the given interval. The type is interface{} with range:
    // 0..4294967295. Units are calls-per-second.
    Cvcallratestatsmaxval interface{}

    // This object indicates the average calls per second that occured for the
    // given period for the given interval. The type is interface{} with range:
    // 0..4294967295. Units are calls-per-second.
    Cvcallratestatsavgval interface{}
}

func (cvcallratestatsentry *CISCOVOICEDIALCONTROLMIB_Cvcallratestatstable_Cvcallratestatsentry) GetFilter() yfilter.YFilter { return cvcallratestatsentry.YFilter }

func (cvcallratestatsentry *CISCOVOICEDIALCONTROLMIB_Cvcallratestatstable_Cvcallratestatsentry) SetFilter(yf yfilter.YFilter) { cvcallratestatsentry.YFilter = yf }

func (cvcallratestatsentry *CISCOVOICEDIALCONTROLMIB_Cvcallratestatstable_Cvcallratestatsentry) GetGoName(yname string) string {
    if yname == "cvCallRateStatsIntvlDurUnits" { return "Cvcallratestatsintvldurunits" }
    if yname == "cvCallRateStatsIntvlDur" { return "Cvcallratestatsintvldur" }
    if yname == "cvCallRateStatsMaxVal" { return "Cvcallratestatsmaxval" }
    if yname == "cvCallRateStatsAvgVal" { return "Cvcallratestatsavgval" }
    return ""
}

func (cvcallratestatsentry *CISCOVOICEDIALCONTROLMIB_Cvcallratestatstable_Cvcallratestatsentry) GetSegmentPath() string {
    return "cvCallRateStatsEntry" + "[cvCallRateStatsIntvlDurUnits='" + fmt.Sprintf("%v", cvcallratestatsentry.Cvcallratestatsintvldurunits) + "']" + "[cvCallRateStatsIntvlDur='" + fmt.Sprintf("%v", cvcallratestatsentry.Cvcallratestatsintvldur) + "']"
}

func (cvcallratestatsentry *CISCOVOICEDIALCONTROLMIB_Cvcallratestatstable_Cvcallratestatsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cvcallratestatsentry *CISCOVOICEDIALCONTROLMIB_Cvcallratestatstable_Cvcallratestatsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cvcallratestatsentry *CISCOVOICEDIALCONTROLMIB_Cvcallratestatstable_Cvcallratestatsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cvCallRateStatsIntvlDurUnits"] = cvcallratestatsentry.Cvcallratestatsintvldurunits
    leafs["cvCallRateStatsIntvlDur"] = cvcallratestatsentry.Cvcallratestatsintvldur
    leafs["cvCallRateStatsMaxVal"] = cvcallratestatsentry.Cvcallratestatsmaxval
    leafs["cvCallRateStatsAvgVal"] = cvcallratestatsentry.Cvcallratestatsavgval
    return leafs
}

func (cvcallratestatsentry *CISCOVOICEDIALCONTROLMIB_Cvcallratestatstable_Cvcallratestatsentry) GetBundleName() string { return "cisco_ios_xe" }

func (cvcallratestatsentry *CISCOVOICEDIALCONTROLMIB_Cvcallratestatstable_Cvcallratestatsentry) GetYangName() string { return "cvCallRateStatsEntry" }

func (cvcallratestatsentry *CISCOVOICEDIALCONTROLMIB_Cvcallratestatstable_Cvcallratestatsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvcallratestatsentry *CISCOVOICEDIALCONTROLMIB_Cvcallratestatstable_Cvcallratestatsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvcallratestatsentry *CISCOVOICEDIALCONTROLMIB_Cvcallratestatstable_Cvcallratestatsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvcallratestatsentry *CISCOVOICEDIALCONTROLMIB_Cvcallratestatstable_Cvcallratestatsentry) SetParent(parent types.Entity) { cvcallratestatsentry.parent = parent }

func (cvcallratestatsentry *CISCOVOICEDIALCONTROLMIB_Cvcallratestatstable_Cvcallratestatsentry) GetParent() types.Entity { return cvcallratestatsentry.parent }

func (cvcallratestatsentry *CISCOVOICEDIALCONTROLMIB_Cvcallratestatstable_Cvcallratestatsentry) GetParentYangName() string { return "cvCallRateStatsTable" }

// CISCOVOICEDIALCONTROLMIB_Cvcalllegratestatstable
// cvCallLegRateStatsTable table represents voice call leg rate
// measurement in various interval lengths defined by 
// the CvCallVolumeStatsIntvlType object.
// Each interval may contain one or more entries to allow for
// detailed measurement to be collected.
type CISCOVOICEDIALCONTROLMIB_Cvcalllegratestatstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This is a conceptual-row in cvCallLegRateStatsTable This entry is created
    // at the system initialization and is updated at every epoch based on
    // CvCallVolumeStatsIntvlType. The type is slice of
    // CISCOVOICEDIALCONTROLMIB_Cvcalllegratestatstable_Cvcalllegratestatsentry.
    Cvcalllegratestatsentry []CISCOVOICEDIALCONTROLMIB_Cvcalllegratestatstable_Cvcalllegratestatsentry
}

func (cvcalllegratestatstable *CISCOVOICEDIALCONTROLMIB_Cvcalllegratestatstable) GetFilter() yfilter.YFilter { return cvcalllegratestatstable.YFilter }

func (cvcalllegratestatstable *CISCOVOICEDIALCONTROLMIB_Cvcalllegratestatstable) SetFilter(yf yfilter.YFilter) { cvcalllegratestatstable.YFilter = yf }

func (cvcalllegratestatstable *CISCOVOICEDIALCONTROLMIB_Cvcalllegratestatstable) GetGoName(yname string) string {
    if yname == "cvCallLegRateStatsEntry" { return "Cvcalllegratestatsentry" }
    return ""
}

func (cvcalllegratestatstable *CISCOVOICEDIALCONTROLMIB_Cvcalllegratestatstable) GetSegmentPath() string {
    return "cvCallLegRateStatsTable"
}

func (cvcalllegratestatstable *CISCOVOICEDIALCONTROLMIB_Cvcalllegratestatstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cvCallLegRateStatsEntry" {
        for _, c := range cvcalllegratestatstable.Cvcalllegratestatsentry {
            if cvcalllegratestatstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOVOICEDIALCONTROLMIB_Cvcalllegratestatstable_Cvcalllegratestatsentry{}
        cvcalllegratestatstable.Cvcalllegratestatsentry = append(cvcalllegratestatstable.Cvcalllegratestatsentry, child)
        return &cvcalllegratestatstable.Cvcalllegratestatsentry[len(cvcalllegratestatstable.Cvcalllegratestatsentry)-1]
    }
    return nil
}

func (cvcalllegratestatstable *CISCOVOICEDIALCONTROLMIB_Cvcalllegratestatstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cvcalllegratestatstable.Cvcalllegratestatsentry {
        children[cvcalllegratestatstable.Cvcalllegratestatsentry[i].GetSegmentPath()] = &cvcalllegratestatstable.Cvcalllegratestatsentry[i]
    }
    return children
}

func (cvcalllegratestatstable *CISCOVOICEDIALCONTROLMIB_Cvcalllegratestatstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cvcalllegratestatstable *CISCOVOICEDIALCONTROLMIB_Cvcalllegratestatstable) GetBundleName() string { return "cisco_ios_xe" }

func (cvcalllegratestatstable *CISCOVOICEDIALCONTROLMIB_Cvcalllegratestatstable) GetYangName() string { return "cvCallLegRateStatsTable" }

func (cvcalllegratestatstable *CISCOVOICEDIALCONTROLMIB_Cvcalllegratestatstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvcalllegratestatstable *CISCOVOICEDIALCONTROLMIB_Cvcalllegratestatstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvcalllegratestatstable *CISCOVOICEDIALCONTROLMIB_Cvcalllegratestatstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvcalllegratestatstable *CISCOVOICEDIALCONTROLMIB_Cvcalllegratestatstable) SetParent(parent types.Entity) { cvcalllegratestatstable.parent = parent }

func (cvcalllegratestatstable *CISCOVOICEDIALCONTROLMIB_Cvcalllegratestatstable) GetParent() types.Entity { return cvcalllegratestatstable.parent }

func (cvcalllegratestatstable *CISCOVOICEDIALCONTROLMIB_Cvcalllegratestatstable) GetParentYangName() string { return "CISCO-VOICE-DIAL-CONTROL-MIB" }

// CISCOVOICEDIALCONTROLMIB_Cvcalllegratestatstable_Cvcalllegratestatsentry
// This is a conceptual-row in cvCallLegRateStatsTable
// This entry is created at the system initialization and is
// updated at every epoch based on CvCallVolumeStatsIntvlType
type CISCOVOICEDIALCONTROLMIB_Cvcalllegratestatstable_Cvcalllegratestatsentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The Object indexes in Call Leg Rate Table to
    // select one among three interval-tables.  The different types in this table
    // are represented by  CvCallVolumeStatsIntvlType. The type is
    // CvCallVolumeStatsIntvlType.
    Cvcalllegratestatsintvldurunits interface{}

    // This attribute is a key. This is an index that references to the different
    // past periods in given in interval of call rate table. This range is 1-60
    // for Seconds and Minutes table  wherein 1-72 for hours table. The type is
    // interface{} with range: 1..72.
    Cvcalllegratestatsintvldur interface{}

    // This object indicates the maximum call-legs per second that occured for the
    // given period for the given interval. The type is interface{} with range:
    // 0..4294967295. Units are call-legs per second.
    Cvcalllegratestatsmaxval interface{}

    // This object indicates the average call-legs per second that occured for the
    // given period for the given interval. The type is interface{} with range:
    // 0..4294967295. Units are call-legs per second.
    Cvcalllegratestatsavgval interface{}
}

func (cvcalllegratestatsentry *CISCOVOICEDIALCONTROLMIB_Cvcalllegratestatstable_Cvcalllegratestatsentry) GetFilter() yfilter.YFilter { return cvcalllegratestatsentry.YFilter }

func (cvcalllegratestatsentry *CISCOVOICEDIALCONTROLMIB_Cvcalllegratestatstable_Cvcalllegratestatsentry) SetFilter(yf yfilter.YFilter) { cvcalllegratestatsentry.YFilter = yf }

func (cvcalllegratestatsentry *CISCOVOICEDIALCONTROLMIB_Cvcalllegratestatstable_Cvcalllegratestatsentry) GetGoName(yname string) string {
    if yname == "cvCallLegRateStatsIntvlDurUnits" { return "Cvcalllegratestatsintvldurunits" }
    if yname == "cvCallLegRateStatsIntvlDur" { return "Cvcalllegratestatsintvldur" }
    if yname == "cvCallLegRateStatsMaxVal" { return "Cvcalllegratestatsmaxval" }
    if yname == "cvCallLegRateStatsAvgVal" { return "Cvcalllegratestatsavgval" }
    return ""
}

func (cvcalllegratestatsentry *CISCOVOICEDIALCONTROLMIB_Cvcalllegratestatstable_Cvcalllegratestatsentry) GetSegmentPath() string {
    return "cvCallLegRateStatsEntry" + "[cvCallLegRateStatsIntvlDurUnits='" + fmt.Sprintf("%v", cvcalllegratestatsentry.Cvcalllegratestatsintvldurunits) + "']" + "[cvCallLegRateStatsIntvlDur='" + fmt.Sprintf("%v", cvcalllegratestatsentry.Cvcalllegratestatsintvldur) + "']"
}

func (cvcalllegratestatsentry *CISCOVOICEDIALCONTROLMIB_Cvcalllegratestatstable_Cvcalllegratestatsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cvcalllegratestatsentry *CISCOVOICEDIALCONTROLMIB_Cvcalllegratestatstable_Cvcalllegratestatsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cvcalllegratestatsentry *CISCOVOICEDIALCONTROLMIB_Cvcalllegratestatstable_Cvcalllegratestatsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cvCallLegRateStatsIntvlDurUnits"] = cvcalllegratestatsentry.Cvcalllegratestatsintvldurunits
    leafs["cvCallLegRateStatsIntvlDur"] = cvcalllegratestatsentry.Cvcalllegratestatsintvldur
    leafs["cvCallLegRateStatsMaxVal"] = cvcalllegratestatsentry.Cvcalllegratestatsmaxval
    leafs["cvCallLegRateStatsAvgVal"] = cvcalllegratestatsentry.Cvcalllegratestatsavgval
    return leafs
}

func (cvcalllegratestatsentry *CISCOVOICEDIALCONTROLMIB_Cvcalllegratestatstable_Cvcalllegratestatsentry) GetBundleName() string { return "cisco_ios_xe" }

func (cvcalllegratestatsentry *CISCOVOICEDIALCONTROLMIB_Cvcalllegratestatstable_Cvcalllegratestatsentry) GetYangName() string { return "cvCallLegRateStatsEntry" }

func (cvcalllegratestatsentry *CISCOVOICEDIALCONTROLMIB_Cvcalllegratestatstable_Cvcalllegratestatsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvcalllegratestatsentry *CISCOVOICEDIALCONTROLMIB_Cvcalllegratestatstable_Cvcalllegratestatsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvcalllegratestatsentry *CISCOVOICEDIALCONTROLMIB_Cvcalllegratestatstable_Cvcalllegratestatsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvcalllegratestatsentry *CISCOVOICEDIALCONTROLMIB_Cvcalllegratestatstable_Cvcalllegratestatsentry) SetParent(parent types.Entity) { cvcalllegratestatsentry.parent = parent }

func (cvcalllegratestatsentry *CISCOVOICEDIALCONTROLMIB_Cvcalllegratestatstable_Cvcalllegratestatsentry) GetParent() types.Entity { return cvcalllegratestatsentry.parent }

func (cvcalllegratestatsentry *CISCOVOICEDIALCONTROLMIB_Cvcalllegratestatstable_Cvcalllegratestatsentry) GetParentYangName() string { return "cvCallLegRateStatsTable" }

// CISCOVOICEDIALCONTROLMIB_Cvactivecallstatstable
// This table represents the active voice calls in various
// interval lengths defined by the 
// CvCallVolumeStatsIntvlType object.
// 
// Each interval may contain one or more entries to allow for
// detailed measurement to be collected.
type CISCOVOICEDIALCONTROLMIB_Cvactivecallstatstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This is a conceptual-row in cvActiveCallStatsTable This entry is created at
    // the system initialization and is updated at every epoch based on
    // CvCallVolumeStatsIntvlType. The type is slice of
    // CISCOVOICEDIALCONTROLMIB_Cvactivecallstatstable_Cvactivecallstatsentry.
    Cvactivecallstatsentry []CISCOVOICEDIALCONTROLMIB_Cvactivecallstatstable_Cvactivecallstatsentry
}

func (cvactivecallstatstable *CISCOVOICEDIALCONTROLMIB_Cvactivecallstatstable) GetFilter() yfilter.YFilter { return cvactivecallstatstable.YFilter }

func (cvactivecallstatstable *CISCOVOICEDIALCONTROLMIB_Cvactivecallstatstable) SetFilter(yf yfilter.YFilter) { cvactivecallstatstable.YFilter = yf }

func (cvactivecallstatstable *CISCOVOICEDIALCONTROLMIB_Cvactivecallstatstable) GetGoName(yname string) string {
    if yname == "cvActiveCallStatsEntry" { return "Cvactivecallstatsentry" }
    return ""
}

func (cvactivecallstatstable *CISCOVOICEDIALCONTROLMIB_Cvactivecallstatstable) GetSegmentPath() string {
    return "cvActiveCallStatsTable"
}

func (cvactivecallstatstable *CISCOVOICEDIALCONTROLMIB_Cvactivecallstatstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cvActiveCallStatsEntry" {
        for _, c := range cvactivecallstatstable.Cvactivecallstatsentry {
            if cvactivecallstatstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOVOICEDIALCONTROLMIB_Cvactivecallstatstable_Cvactivecallstatsentry{}
        cvactivecallstatstable.Cvactivecallstatsentry = append(cvactivecallstatstable.Cvactivecallstatsentry, child)
        return &cvactivecallstatstable.Cvactivecallstatsentry[len(cvactivecallstatstable.Cvactivecallstatsentry)-1]
    }
    return nil
}

func (cvactivecallstatstable *CISCOVOICEDIALCONTROLMIB_Cvactivecallstatstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cvactivecallstatstable.Cvactivecallstatsentry {
        children[cvactivecallstatstable.Cvactivecallstatsentry[i].GetSegmentPath()] = &cvactivecallstatstable.Cvactivecallstatsentry[i]
    }
    return children
}

func (cvactivecallstatstable *CISCOVOICEDIALCONTROLMIB_Cvactivecallstatstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cvactivecallstatstable *CISCOVOICEDIALCONTROLMIB_Cvactivecallstatstable) GetBundleName() string { return "cisco_ios_xe" }

func (cvactivecallstatstable *CISCOVOICEDIALCONTROLMIB_Cvactivecallstatstable) GetYangName() string { return "cvActiveCallStatsTable" }

func (cvactivecallstatstable *CISCOVOICEDIALCONTROLMIB_Cvactivecallstatstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvactivecallstatstable *CISCOVOICEDIALCONTROLMIB_Cvactivecallstatstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvactivecallstatstable *CISCOVOICEDIALCONTROLMIB_Cvactivecallstatstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvactivecallstatstable *CISCOVOICEDIALCONTROLMIB_Cvactivecallstatstable) SetParent(parent types.Entity) { cvactivecallstatstable.parent = parent }

func (cvactivecallstatstable *CISCOVOICEDIALCONTROLMIB_Cvactivecallstatstable) GetParent() types.Entity { return cvactivecallstatstable.parent }

func (cvactivecallstatstable *CISCOVOICEDIALCONTROLMIB_Cvactivecallstatstable) GetParentYangName() string { return "CISCO-VOICE-DIAL-CONTROL-MIB" }

// CISCOVOICEDIALCONTROLMIB_Cvactivecallstatstable_Cvactivecallstatsentry
// This is a conceptual-row in cvActiveCallStatsTable
// This entry is created at the system initialization and is
// updated at every epoch based on CvCallVolumeStatsIntvlType
type CISCOVOICEDIALCONTROLMIB_Cvactivecallstatstable_Cvactivecallstatsentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The Object indexes in Active Call Rate Table
    // (con-current calls table) to select one among three interval-tables.  The
    // different types in this table are represented by 
    // CvCallVolumeStatsIntvlType. The type is CvCallVolumeStatsIntvlType.
    Cvactivecallstatsintvldurunits interface{}

    // This attribute is a key. This is an index that references to the different
    // past periods in given in interval of active call table. This range is 1-60
    // for Seconds and Minutes table  wherein 1-72 for hours table. The type is
    // interface{} with range: 1..72.
    Cvactivecallstatsintvldur interface{}

    // This object indicates the maximum number of active call that occured for
    // the given period for the given interval. The type is interface{} with
    // range: 0..4294967295. Units are calls.
    Cvactivecallstatsmaxval interface{}

    // This object indicates the average number of active calls that occured for
    // the given period for the given interval. The type is interface{} with
    // range: 0..4294967295. Units are calls.
    Cvactivecallstatsavgval interface{}
}

func (cvactivecallstatsentry *CISCOVOICEDIALCONTROLMIB_Cvactivecallstatstable_Cvactivecallstatsentry) GetFilter() yfilter.YFilter { return cvactivecallstatsentry.YFilter }

func (cvactivecallstatsentry *CISCOVOICEDIALCONTROLMIB_Cvactivecallstatstable_Cvactivecallstatsentry) SetFilter(yf yfilter.YFilter) { cvactivecallstatsentry.YFilter = yf }

func (cvactivecallstatsentry *CISCOVOICEDIALCONTROLMIB_Cvactivecallstatstable_Cvactivecallstatsentry) GetGoName(yname string) string {
    if yname == "cvActiveCallStatsIntvlDurUnits" { return "Cvactivecallstatsintvldurunits" }
    if yname == "cvActiveCallStatsIntvlDur" { return "Cvactivecallstatsintvldur" }
    if yname == "cvActiveCallStatsMaxVal" { return "Cvactivecallstatsmaxval" }
    if yname == "cvActiveCallStatsAvgVal" { return "Cvactivecallstatsavgval" }
    return ""
}

func (cvactivecallstatsentry *CISCOVOICEDIALCONTROLMIB_Cvactivecallstatstable_Cvactivecallstatsentry) GetSegmentPath() string {
    return "cvActiveCallStatsEntry" + "[cvActiveCallStatsIntvlDurUnits='" + fmt.Sprintf("%v", cvactivecallstatsentry.Cvactivecallstatsintvldurunits) + "']" + "[cvActiveCallStatsIntvlDur='" + fmt.Sprintf("%v", cvactivecallstatsentry.Cvactivecallstatsintvldur) + "']"
}

func (cvactivecallstatsentry *CISCOVOICEDIALCONTROLMIB_Cvactivecallstatstable_Cvactivecallstatsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cvactivecallstatsentry *CISCOVOICEDIALCONTROLMIB_Cvactivecallstatstable_Cvactivecallstatsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cvactivecallstatsentry *CISCOVOICEDIALCONTROLMIB_Cvactivecallstatstable_Cvactivecallstatsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cvActiveCallStatsIntvlDurUnits"] = cvactivecallstatsentry.Cvactivecallstatsintvldurunits
    leafs["cvActiveCallStatsIntvlDur"] = cvactivecallstatsentry.Cvactivecallstatsintvldur
    leafs["cvActiveCallStatsMaxVal"] = cvactivecallstatsentry.Cvactivecallstatsmaxval
    leafs["cvActiveCallStatsAvgVal"] = cvactivecallstatsentry.Cvactivecallstatsavgval
    return leafs
}

func (cvactivecallstatsentry *CISCOVOICEDIALCONTROLMIB_Cvactivecallstatstable_Cvactivecallstatsentry) GetBundleName() string { return "cisco_ios_xe" }

func (cvactivecallstatsentry *CISCOVOICEDIALCONTROLMIB_Cvactivecallstatstable_Cvactivecallstatsentry) GetYangName() string { return "cvActiveCallStatsEntry" }

func (cvactivecallstatsentry *CISCOVOICEDIALCONTROLMIB_Cvactivecallstatstable_Cvactivecallstatsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvactivecallstatsentry *CISCOVOICEDIALCONTROLMIB_Cvactivecallstatstable_Cvactivecallstatsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvactivecallstatsentry *CISCOVOICEDIALCONTROLMIB_Cvactivecallstatstable_Cvactivecallstatsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvactivecallstatsentry *CISCOVOICEDIALCONTROLMIB_Cvactivecallstatstable_Cvactivecallstatsentry) SetParent(parent types.Entity) { cvactivecallstatsentry.parent = parent }

func (cvactivecallstatsentry *CISCOVOICEDIALCONTROLMIB_Cvactivecallstatstable_Cvactivecallstatsentry) GetParent() types.Entity { return cvactivecallstatsentry.parent }

func (cvactivecallstatsentry *CISCOVOICEDIALCONTROLMIB_Cvactivecallstatstable_Cvactivecallstatsentry) GetParentYangName() string { return "cvActiveCallStatsTable" }

// CISCOVOICEDIALCONTROLMIB_Cvcalldurationstatstable
// This table represents the number of calls below a specific
// duration in various interval length defined by 
// the CvCallVolumeStatsIntvlType object.  
// 
// The specific duration is configurable value of 
//  cvCallDurationStatsThreshold object.
// 
// Each interval may contain one or more entries to allow for 
// detailed measurement to be collected.
type CISCOVOICEDIALCONTROLMIB_Cvcalldurationstatstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This is a conceptual-row in cvCallDurationStatsTable This entry is created
    // at the system initialization and is updated at every epoch based on
    // CvCallVolumeStatsIntvlType. The type is slice of
    // CISCOVOICEDIALCONTROLMIB_Cvcalldurationstatstable_Cvcalldurationstatsentry.
    Cvcalldurationstatsentry []CISCOVOICEDIALCONTROLMIB_Cvcalldurationstatstable_Cvcalldurationstatsentry
}

func (cvcalldurationstatstable *CISCOVOICEDIALCONTROLMIB_Cvcalldurationstatstable) GetFilter() yfilter.YFilter { return cvcalldurationstatstable.YFilter }

func (cvcalldurationstatstable *CISCOVOICEDIALCONTROLMIB_Cvcalldurationstatstable) SetFilter(yf yfilter.YFilter) { cvcalldurationstatstable.YFilter = yf }

func (cvcalldurationstatstable *CISCOVOICEDIALCONTROLMIB_Cvcalldurationstatstable) GetGoName(yname string) string {
    if yname == "cvCallDurationStatsEntry" { return "Cvcalldurationstatsentry" }
    return ""
}

func (cvcalldurationstatstable *CISCOVOICEDIALCONTROLMIB_Cvcalldurationstatstable) GetSegmentPath() string {
    return "cvCallDurationStatsTable"
}

func (cvcalldurationstatstable *CISCOVOICEDIALCONTROLMIB_Cvcalldurationstatstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cvCallDurationStatsEntry" {
        for _, c := range cvcalldurationstatstable.Cvcalldurationstatsentry {
            if cvcalldurationstatstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOVOICEDIALCONTROLMIB_Cvcalldurationstatstable_Cvcalldurationstatsentry{}
        cvcalldurationstatstable.Cvcalldurationstatsentry = append(cvcalldurationstatstable.Cvcalldurationstatsentry, child)
        return &cvcalldurationstatstable.Cvcalldurationstatsentry[len(cvcalldurationstatstable.Cvcalldurationstatsentry)-1]
    }
    return nil
}

func (cvcalldurationstatstable *CISCOVOICEDIALCONTROLMIB_Cvcalldurationstatstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cvcalldurationstatstable.Cvcalldurationstatsentry {
        children[cvcalldurationstatstable.Cvcalldurationstatsentry[i].GetSegmentPath()] = &cvcalldurationstatstable.Cvcalldurationstatsentry[i]
    }
    return children
}

func (cvcalldurationstatstable *CISCOVOICEDIALCONTROLMIB_Cvcalldurationstatstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cvcalldurationstatstable *CISCOVOICEDIALCONTROLMIB_Cvcalldurationstatstable) GetBundleName() string { return "cisco_ios_xe" }

func (cvcalldurationstatstable *CISCOVOICEDIALCONTROLMIB_Cvcalldurationstatstable) GetYangName() string { return "cvCallDurationStatsTable" }

func (cvcalldurationstatstable *CISCOVOICEDIALCONTROLMIB_Cvcalldurationstatstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvcalldurationstatstable *CISCOVOICEDIALCONTROLMIB_Cvcalldurationstatstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvcalldurationstatstable *CISCOVOICEDIALCONTROLMIB_Cvcalldurationstatstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvcalldurationstatstable *CISCOVOICEDIALCONTROLMIB_Cvcalldurationstatstable) SetParent(parent types.Entity) { cvcalldurationstatstable.parent = parent }

func (cvcalldurationstatstable *CISCOVOICEDIALCONTROLMIB_Cvcalldurationstatstable) GetParent() types.Entity { return cvcalldurationstatstable.parent }

func (cvcalldurationstatstable *CISCOVOICEDIALCONTROLMIB_Cvcalldurationstatstable) GetParentYangName() string { return "CISCO-VOICE-DIAL-CONTROL-MIB" }

// CISCOVOICEDIALCONTROLMIB_Cvcalldurationstatstable_Cvcalldurationstatsentry
// This is a conceptual-row in cvCallDurationStatsTable
// This entry is created at the system initialization and is
// updated at every epoch based on CvCallVolumeStatsIntvlType
type CISCOVOICEDIALCONTROLMIB_Cvcalldurationstatstable_Cvcalldurationstatsentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The Object indexes in Call Duration Table to
    // select one among three interval-tables.  The different types in this table
    // are represented by  CvCallVolumeStatsIntvlType. The type is
    // CvCallVolumeStatsIntvlType.
    Cvcalldurationstatsintvldurunits interface{}

    // This attribute is a key. This is an index that references to the different
    // past periods in given in interval of call Duration table. This range is
    // 1-60 for Seconds and Minutes table  wherein 1-72 for hours table. The type
    // is interface{} with range: 1..72.
    Cvcalldurationstatsintvldur interface{}

    // This object indicates the maximum number of calls having a duration which
    // is below the threshold for the given interval. The type is interface{} with
    // range: 0..4294967295. Units are calls.
    Cvcalldurationstatsmaxval interface{}

    // This object indicates the average number of calls having a duration which
    // is below the threshold for the given interval. The type is interface{} with
    // range: 0..4294967295. Units are calls.
    Cvcalldurationstatsavgval interface{}
}

func (cvcalldurationstatsentry *CISCOVOICEDIALCONTROLMIB_Cvcalldurationstatstable_Cvcalldurationstatsentry) GetFilter() yfilter.YFilter { return cvcalldurationstatsentry.YFilter }

func (cvcalldurationstatsentry *CISCOVOICEDIALCONTROLMIB_Cvcalldurationstatstable_Cvcalldurationstatsentry) SetFilter(yf yfilter.YFilter) { cvcalldurationstatsentry.YFilter = yf }

func (cvcalldurationstatsentry *CISCOVOICEDIALCONTROLMIB_Cvcalldurationstatstable_Cvcalldurationstatsentry) GetGoName(yname string) string {
    if yname == "cvCallDurationStatsIntvlDurUnits" { return "Cvcalldurationstatsintvldurunits" }
    if yname == "cvCallDurationStatsIntvlDur" { return "Cvcalldurationstatsintvldur" }
    if yname == "cvCallDurationStatsMaxVal" { return "Cvcalldurationstatsmaxval" }
    if yname == "cvCallDurationStatsAvgVal" { return "Cvcalldurationstatsavgval" }
    return ""
}

func (cvcalldurationstatsentry *CISCOVOICEDIALCONTROLMIB_Cvcalldurationstatstable_Cvcalldurationstatsentry) GetSegmentPath() string {
    return "cvCallDurationStatsEntry" + "[cvCallDurationStatsIntvlDurUnits='" + fmt.Sprintf("%v", cvcalldurationstatsentry.Cvcalldurationstatsintvldurunits) + "']" + "[cvCallDurationStatsIntvlDur='" + fmt.Sprintf("%v", cvcalldurationstatsentry.Cvcalldurationstatsintvldur) + "']"
}

func (cvcalldurationstatsentry *CISCOVOICEDIALCONTROLMIB_Cvcalldurationstatstable_Cvcalldurationstatsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cvcalldurationstatsentry *CISCOVOICEDIALCONTROLMIB_Cvcalldurationstatstable_Cvcalldurationstatsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cvcalldurationstatsentry *CISCOVOICEDIALCONTROLMIB_Cvcalldurationstatstable_Cvcalldurationstatsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cvCallDurationStatsIntvlDurUnits"] = cvcalldurationstatsentry.Cvcalldurationstatsintvldurunits
    leafs["cvCallDurationStatsIntvlDur"] = cvcalldurationstatsentry.Cvcalldurationstatsintvldur
    leafs["cvCallDurationStatsMaxVal"] = cvcalldurationstatsentry.Cvcalldurationstatsmaxval
    leafs["cvCallDurationStatsAvgVal"] = cvcalldurationstatsentry.Cvcalldurationstatsavgval
    return leafs
}

func (cvcalldurationstatsentry *CISCOVOICEDIALCONTROLMIB_Cvcalldurationstatstable_Cvcalldurationstatsentry) GetBundleName() string { return "cisco_ios_xe" }

func (cvcalldurationstatsentry *CISCOVOICEDIALCONTROLMIB_Cvcalldurationstatstable_Cvcalldurationstatsentry) GetYangName() string { return "cvCallDurationStatsEntry" }

func (cvcalldurationstatsentry *CISCOVOICEDIALCONTROLMIB_Cvcalldurationstatstable_Cvcalldurationstatsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvcalldurationstatsentry *CISCOVOICEDIALCONTROLMIB_Cvcalldurationstatstable_Cvcalldurationstatsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvcalldurationstatsentry *CISCOVOICEDIALCONTROLMIB_Cvcalldurationstatstable_Cvcalldurationstatsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvcalldurationstatsentry *CISCOVOICEDIALCONTROLMIB_Cvcalldurationstatstable_Cvcalldurationstatsentry) SetParent(parent types.Entity) { cvcalldurationstatsentry.parent = parent }

func (cvcalldurationstatsentry *CISCOVOICEDIALCONTROLMIB_Cvcalldurationstatstable_Cvcalldurationstatsentry) GetParent() types.Entity { return cvcalldurationstatsentry.parent }

func (cvcalldurationstatsentry *CISCOVOICEDIALCONTROLMIB_Cvcalldurationstatstable_Cvcalldurationstatsentry) GetParentYangName() string { return "cvCallDurationStatsTable" }

// CISCOVOICEDIALCONTROLMIB_Cvsipmsgratestatstable
// This table represents the SIP message rate measurement in
// various interval length defined by the 
// CvCallVolumeStatsIntvlType object.
// 
// Each interval may contain one or more entries to allow for
// detailed measurement to be collected
type CISCOVOICEDIALCONTROLMIB_Cvsipmsgratestatstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This is a conceptual-row in cvSipMsgRateStatsTable This entry is created at
    // the system initialization and is updated at every epoch based on
    // CvCallVolumeStatsIntvlType. The type is slice of
    // CISCOVOICEDIALCONTROLMIB_Cvsipmsgratestatstable_Cvsipmsgratestatsentry.
    Cvsipmsgratestatsentry []CISCOVOICEDIALCONTROLMIB_Cvsipmsgratestatstable_Cvsipmsgratestatsentry
}

func (cvsipmsgratestatstable *CISCOVOICEDIALCONTROLMIB_Cvsipmsgratestatstable) GetFilter() yfilter.YFilter { return cvsipmsgratestatstable.YFilter }

func (cvsipmsgratestatstable *CISCOVOICEDIALCONTROLMIB_Cvsipmsgratestatstable) SetFilter(yf yfilter.YFilter) { cvsipmsgratestatstable.YFilter = yf }

func (cvsipmsgratestatstable *CISCOVOICEDIALCONTROLMIB_Cvsipmsgratestatstable) GetGoName(yname string) string {
    if yname == "cvSipMsgRateStatsEntry" { return "Cvsipmsgratestatsentry" }
    return ""
}

func (cvsipmsgratestatstable *CISCOVOICEDIALCONTROLMIB_Cvsipmsgratestatstable) GetSegmentPath() string {
    return "cvSipMsgRateStatsTable"
}

func (cvsipmsgratestatstable *CISCOVOICEDIALCONTROLMIB_Cvsipmsgratestatstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cvSipMsgRateStatsEntry" {
        for _, c := range cvsipmsgratestatstable.Cvsipmsgratestatsentry {
            if cvsipmsgratestatstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOVOICEDIALCONTROLMIB_Cvsipmsgratestatstable_Cvsipmsgratestatsentry{}
        cvsipmsgratestatstable.Cvsipmsgratestatsentry = append(cvsipmsgratestatstable.Cvsipmsgratestatsentry, child)
        return &cvsipmsgratestatstable.Cvsipmsgratestatsentry[len(cvsipmsgratestatstable.Cvsipmsgratestatsentry)-1]
    }
    return nil
}

func (cvsipmsgratestatstable *CISCOVOICEDIALCONTROLMIB_Cvsipmsgratestatstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cvsipmsgratestatstable.Cvsipmsgratestatsentry {
        children[cvsipmsgratestatstable.Cvsipmsgratestatsentry[i].GetSegmentPath()] = &cvsipmsgratestatstable.Cvsipmsgratestatsentry[i]
    }
    return children
}

func (cvsipmsgratestatstable *CISCOVOICEDIALCONTROLMIB_Cvsipmsgratestatstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cvsipmsgratestatstable *CISCOVOICEDIALCONTROLMIB_Cvsipmsgratestatstable) GetBundleName() string { return "cisco_ios_xe" }

func (cvsipmsgratestatstable *CISCOVOICEDIALCONTROLMIB_Cvsipmsgratestatstable) GetYangName() string { return "cvSipMsgRateStatsTable" }

func (cvsipmsgratestatstable *CISCOVOICEDIALCONTROLMIB_Cvsipmsgratestatstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvsipmsgratestatstable *CISCOVOICEDIALCONTROLMIB_Cvsipmsgratestatstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvsipmsgratestatstable *CISCOVOICEDIALCONTROLMIB_Cvsipmsgratestatstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvsipmsgratestatstable *CISCOVOICEDIALCONTROLMIB_Cvsipmsgratestatstable) SetParent(parent types.Entity) { cvsipmsgratestatstable.parent = parent }

func (cvsipmsgratestatstable *CISCOVOICEDIALCONTROLMIB_Cvsipmsgratestatstable) GetParent() types.Entity { return cvsipmsgratestatstable.parent }

func (cvsipmsgratestatstable *CISCOVOICEDIALCONTROLMIB_Cvsipmsgratestatstable) GetParentYangName() string { return "CISCO-VOICE-DIAL-CONTROL-MIB" }

// CISCOVOICEDIALCONTROLMIB_Cvsipmsgratestatstable_Cvsipmsgratestatsentry
// This is a conceptual-row in cvSipMsgRateStatsTable
// This entry is created at the system initialization and is
// updated at every epoch based on CvCallVolumeStatsIntvlType
type CISCOVOICEDIALCONTROLMIB_Cvsipmsgratestatstable_Cvsipmsgratestatsentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The Object indexes in SIP Message Rate Table to
    // select one among three interval-tables.  The different types in this table
    // are represented by  CvCallVolumeStatsIntvlType. The type is
    // CvCallVolumeStatsIntvlType.
    Cvsipmsgratestatsintvldurunits interface{}

    // This attribute is a key. This is an index that references to the different
    // past periods in given in interval of SIP message rate table. This range is
    // 1-60 for Seconds and Minutes table  wherein 1-72 for hours table. The type
    // is interface{} with range: 1..72.
    Cvsipmsgratestatsintvldur interface{}

    // This object indicates the maximum SIP messages  per second that is received
    // for the given interval. The type is interface{} with range: 0..4294967295.
    // Units are SIP messages per second.
    Cvsipmsgratestatsmaxval interface{}

    // This object indicates the average SIP messages per second that is received
    // for the given interval. The type is interface{} with range: 0..4294967295.
    // Units are SIP messages per second.
    Cvsipmsgratestatsavgval interface{}
}

func (cvsipmsgratestatsentry *CISCOVOICEDIALCONTROLMIB_Cvsipmsgratestatstable_Cvsipmsgratestatsentry) GetFilter() yfilter.YFilter { return cvsipmsgratestatsentry.YFilter }

func (cvsipmsgratestatsentry *CISCOVOICEDIALCONTROLMIB_Cvsipmsgratestatstable_Cvsipmsgratestatsentry) SetFilter(yf yfilter.YFilter) { cvsipmsgratestatsentry.YFilter = yf }

func (cvsipmsgratestatsentry *CISCOVOICEDIALCONTROLMIB_Cvsipmsgratestatstable_Cvsipmsgratestatsentry) GetGoName(yname string) string {
    if yname == "cvSipMsgRateStatsIntvlDurUnits" { return "Cvsipmsgratestatsintvldurunits" }
    if yname == "cvSipMsgRateStatsIntvlDur" { return "Cvsipmsgratestatsintvldur" }
    if yname == "cvSipMsgRateStatsMaxVal" { return "Cvsipmsgratestatsmaxval" }
    if yname == "cvSipMsgRateStatsAvgVal" { return "Cvsipmsgratestatsavgval" }
    return ""
}

func (cvsipmsgratestatsentry *CISCOVOICEDIALCONTROLMIB_Cvsipmsgratestatstable_Cvsipmsgratestatsentry) GetSegmentPath() string {
    return "cvSipMsgRateStatsEntry" + "[cvSipMsgRateStatsIntvlDurUnits='" + fmt.Sprintf("%v", cvsipmsgratestatsentry.Cvsipmsgratestatsintvldurunits) + "']" + "[cvSipMsgRateStatsIntvlDur='" + fmt.Sprintf("%v", cvsipmsgratestatsentry.Cvsipmsgratestatsintvldur) + "']"
}

func (cvsipmsgratestatsentry *CISCOVOICEDIALCONTROLMIB_Cvsipmsgratestatstable_Cvsipmsgratestatsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cvsipmsgratestatsentry *CISCOVOICEDIALCONTROLMIB_Cvsipmsgratestatstable_Cvsipmsgratestatsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cvsipmsgratestatsentry *CISCOVOICEDIALCONTROLMIB_Cvsipmsgratestatstable_Cvsipmsgratestatsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cvSipMsgRateStatsIntvlDurUnits"] = cvsipmsgratestatsentry.Cvsipmsgratestatsintvldurunits
    leafs["cvSipMsgRateStatsIntvlDur"] = cvsipmsgratestatsentry.Cvsipmsgratestatsintvldur
    leafs["cvSipMsgRateStatsMaxVal"] = cvsipmsgratestatsentry.Cvsipmsgratestatsmaxval
    leafs["cvSipMsgRateStatsAvgVal"] = cvsipmsgratestatsentry.Cvsipmsgratestatsavgval
    return leafs
}

func (cvsipmsgratestatsentry *CISCOVOICEDIALCONTROLMIB_Cvsipmsgratestatstable_Cvsipmsgratestatsentry) GetBundleName() string { return "cisco_ios_xe" }

func (cvsipmsgratestatsentry *CISCOVOICEDIALCONTROLMIB_Cvsipmsgratestatstable_Cvsipmsgratestatsentry) GetYangName() string { return "cvSipMsgRateStatsEntry" }

func (cvsipmsgratestatsentry *CISCOVOICEDIALCONTROLMIB_Cvsipmsgratestatstable_Cvsipmsgratestatsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvsipmsgratestatsentry *CISCOVOICEDIALCONTROLMIB_Cvsipmsgratestatstable_Cvsipmsgratestatsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvsipmsgratestatsentry *CISCOVOICEDIALCONTROLMIB_Cvsipmsgratestatstable_Cvsipmsgratestatsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvsipmsgratestatsentry *CISCOVOICEDIALCONTROLMIB_Cvsipmsgratestatstable_Cvsipmsgratestatsentry) SetParent(parent types.Entity) { cvsipmsgratestatsentry.parent = parent }

func (cvsipmsgratestatsentry *CISCOVOICEDIALCONTROLMIB_Cvsipmsgratestatstable_Cvsipmsgratestatsentry) GetParent() types.Entity { return cvsipmsgratestatsentry.parent }

func (cvsipmsgratestatsentry *CISCOVOICEDIALCONTROLMIB_Cvsipmsgratestatstable_Cvsipmsgratestatsentry) GetParentYangName() string { return "cvSipMsgRateStatsTable" }

// CISCOVOICEDIALCONTROLMIB_Cvcallratewmtable
// This table represents high watermarks achieved
// by call rate in various interval length defined 
// by CvCallVolumeWMIntvlType. 
// 
// Each interval may contain one or more entries to allow for 
// detailed measurement to be collected
type CISCOVOICEDIALCONTROLMIB_Cvcallratewmtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This is a conceptual-row in cvCallRateWMTable This entry is created at the
    // system initialization and is updated whenever  a) This entry is obsolete OR
    // b) A new/higher entry is available. These entries are
    // reinitialised/added/deleted  if cvCallVolumeWMTableSize is changed. The
    // type is slice of
    // CISCOVOICEDIALCONTROLMIB_Cvcallratewmtable_Cvcallratewmentry.
    Cvcallratewmentry []CISCOVOICEDIALCONTROLMIB_Cvcallratewmtable_Cvcallratewmentry
}

func (cvcallratewmtable *CISCOVOICEDIALCONTROLMIB_Cvcallratewmtable) GetFilter() yfilter.YFilter { return cvcallratewmtable.YFilter }

func (cvcallratewmtable *CISCOVOICEDIALCONTROLMIB_Cvcallratewmtable) SetFilter(yf yfilter.YFilter) { cvcallratewmtable.YFilter = yf }

func (cvcallratewmtable *CISCOVOICEDIALCONTROLMIB_Cvcallratewmtable) GetGoName(yname string) string {
    if yname == "cvCallRateWMEntry" { return "Cvcallratewmentry" }
    return ""
}

func (cvcallratewmtable *CISCOVOICEDIALCONTROLMIB_Cvcallratewmtable) GetSegmentPath() string {
    return "cvCallRateWMTable"
}

func (cvcallratewmtable *CISCOVOICEDIALCONTROLMIB_Cvcallratewmtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cvCallRateWMEntry" {
        for _, c := range cvcallratewmtable.Cvcallratewmentry {
            if cvcallratewmtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOVOICEDIALCONTROLMIB_Cvcallratewmtable_Cvcallratewmentry{}
        cvcallratewmtable.Cvcallratewmentry = append(cvcallratewmtable.Cvcallratewmentry, child)
        return &cvcallratewmtable.Cvcallratewmentry[len(cvcallratewmtable.Cvcallratewmentry)-1]
    }
    return nil
}

func (cvcallratewmtable *CISCOVOICEDIALCONTROLMIB_Cvcallratewmtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cvcallratewmtable.Cvcallratewmentry {
        children[cvcallratewmtable.Cvcallratewmentry[i].GetSegmentPath()] = &cvcallratewmtable.Cvcallratewmentry[i]
    }
    return children
}

func (cvcallratewmtable *CISCOVOICEDIALCONTROLMIB_Cvcallratewmtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cvcallratewmtable *CISCOVOICEDIALCONTROLMIB_Cvcallratewmtable) GetBundleName() string { return "cisco_ios_xe" }

func (cvcallratewmtable *CISCOVOICEDIALCONTROLMIB_Cvcallratewmtable) GetYangName() string { return "cvCallRateWMTable" }

func (cvcallratewmtable *CISCOVOICEDIALCONTROLMIB_Cvcallratewmtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvcallratewmtable *CISCOVOICEDIALCONTROLMIB_Cvcallratewmtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvcallratewmtable *CISCOVOICEDIALCONTROLMIB_Cvcallratewmtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvcallratewmtable *CISCOVOICEDIALCONTROLMIB_Cvcallratewmtable) SetParent(parent types.Entity) { cvcallratewmtable.parent = parent }

func (cvcallratewmtable *CISCOVOICEDIALCONTROLMIB_Cvcallratewmtable) GetParent() types.Entity { return cvcallratewmtable.parent }

func (cvcallratewmtable *CISCOVOICEDIALCONTROLMIB_Cvcallratewmtable) GetParentYangName() string { return "CISCO-VOICE-DIAL-CONTROL-MIB" }

// CISCOVOICEDIALCONTROLMIB_Cvcallratewmtable_Cvcallratewmentry
// This is a conceptual-row in cvCallRateWMTable
// This entry is created at the system initialization and is
// updated whenever 
// a) This entry is obsolete OR
// b) A new/higher entry is available.
// These entries are reinitialised/added/deleted  if
// cvCallVolumeWMTableSize is changed
type CISCOVOICEDIALCONTROLMIB_Cvcallratewmtable_Cvcallratewmentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The Object indexes in call rate Water mark Table
    // to select one among four interval-tables.  The different types in this
    // table are represented by  CvCallVolumeWMIntvlType. The type is
    // CvCallVolumeWMIntvlType.
    Cvcallratewmintvldurunits interface{}

    // This attribute is a key. This is an index that references to different
    // peaks in past period in call rate watermark table.  The number of
    // watermarks entries stored for each table are  based on
    // cvCallVolumeWMTableSize. The type is interface{} with range: 1..10.
    Cvcallratewmindex interface{}

    // This object indicates high watermark value achieved by the calls per second
    // for the given interval. The type is interface{} with range: 0..4294967295.
    // Units are calls per second.
    Cvcallratewmvalue interface{}

    // This object indicates date and Time when the high watermark is achieved for
    // calls per second for the given interval. The type is string.
    Cvcallratewmts interface{}
}

func (cvcallratewmentry *CISCOVOICEDIALCONTROLMIB_Cvcallratewmtable_Cvcallratewmentry) GetFilter() yfilter.YFilter { return cvcallratewmentry.YFilter }

func (cvcallratewmentry *CISCOVOICEDIALCONTROLMIB_Cvcallratewmtable_Cvcallratewmentry) SetFilter(yf yfilter.YFilter) { cvcallratewmentry.YFilter = yf }

func (cvcallratewmentry *CISCOVOICEDIALCONTROLMIB_Cvcallratewmtable_Cvcallratewmentry) GetGoName(yname string) string {
    if yname == "cvCallRateWMIntvlDurUnits" { return "Cvcallratewmintvldurunits" }
    if yname == "cvCallRateWMIndex" { return "Cvcallratewmindex" }
    if yname == "cvCallRateWMValue" { return "Cvcallratewmvalue" }
    if yname == "cvCallRateWMts" { return "Cvcallratewmts" }
    return ""
}

func (cvcallratewmentry *CISCOVOICEDIALCONTROLMIB_Cvcallratewmtable_Cvcallratewmentry) GetSegmentPath() string {
    return "cvCallRateWMEntry" + "[cvCallRateWMIntvlDurUnits='" + fmt.Sprintf("%v", cvcallratewmentry.Cvcallratewmintvldurunits) + "']" + "[cvCallRateWMIndex='" + fmt.Sprintf("%v", cvcallratewmentry.Cvcallratewmindex) + "']"
}

func (cvcallratewmentry *CISCOVOICEDIALCONTROLMIB_Cvcallratewmtable_Cvcallratewmentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cvcallratewmentry *CISCOVOICEDIALCONTROLMIB_Cvcallratewmtable_Cvcallratewmentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cvcallratewmentry *CISCOVOICEDIALCONTROLMIB_Cvcallratewmtable_Cvcallratewmentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cvCallRateWMIntvlDurUnits"] = cvcallratewmentry.Cvcallratewmintvldurunits
    leafs["cvCallRateWMIndex"] = cvcallratewmentry.Cvcallratewmindex
    leafs["cvCallRateWMValue"] = cvcallratewmentry.Cvcallratewmvalue
    leafs["cvCallRateWMts"] = cvcallratewmentry.Cvcallratewmts
    return leafs
}

func (cvcallratewmentry *CISCOVOICEDIALCONTROLMIB_Cvcallratewmtable_Cvcallratewmentry) GetBundleName() string { return "cisco_ios_xe" }

func (cvcallratewmentry *CISCOVOICEDIALCONTROLMIB_Cvcallratewmtable_Cvcallratewmentry) GetYangName() string { return "cvCallRateWMEntry" }

func (cvcallratewmentry *CISCOVOICEDIALCONTROLMIB_Cvcallratewmtable_Cvcallratewmentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvcallratewmentry *CISCOVOICEDIALCONTROLMIB_Cvcallratewmtable_Cvcallratewmentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvcallratewmentry *CISCOVOICEDIALCONTROLMIB_Cvcallratewmtable_Cvcallratewmentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvcallratewmentry *CISCOVOICEDIALCONTROLMIB_Cvcallratewmtable_Cvcallratewmentry) SetParent(parent types.Entity) { cvcallratewmentry.parent = parent }

func (cvcallratewmentry *CISCOVOICEDIALCONTROLMIB_Cvcallratewmtable_Cvcallratewmentry) GetParent() types.Entity { return cvcallratewmentry.parent }

func (cvcallratewmentry *CISCOVOICEDIALCONTROLMIB_Cvcallratewmtable_Cvcallratewmentry) GetParentYangName() string { return "cvCallRateWMTable" }

// CISCOVOICEDIALCONTROLMIB_Cvcalllegratewmtable
// cvCallLegRateWMTable table represents high watermarks achieved
// by call-leg rate in various interval length defined 
// by CvCallVolumeWMIntvlType. 
// 
// Each interval may contain one or more entries to allow for 
// detailed measurement to be collected
type CISCOVOICEDIALCONTROLMIB_Cvcalllegratewmtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This is a conceptual-row in cvCallLegRateWMTable This entry is created at
    // the system initialization and is updated whenever  a) This entry is
    // obsolete OR b) A new/higher entry is available. These entries are
    // reinitialised/added/deleted  if cvCallVolumeWMTableSize is changed. The
    // type is slice of
    // CISCOVOICEDIALCONTROLMIB_Cvcalllegratewmtable_Cvcalllegratewmentry.
    Cvcalllegratewmentry []CISCOVOICEDIALCONTROLMIB_Cvcalllegratewmtable_Cvcalllegratewmentry
}

func (cvcalllegratewmtable *CISCOVOICEDIALCONTROLMIB_Cvcalllegratewmtable) GetFilter() yfilter.YFilter { return cvcalllegratewmtable.YFilter }

func (cvcalllegratewmtable *CISCOVOICEDIALCONTROLMIB_Cvcalllegratewmtable) SetFilter(yf yfilter.YFilter) { cvcalllegratewmtable.YFilter = yf }

func (cvcalllegratewmtable *CISCOVOICEDIALCONTROLMIB_Cvcalllegratewmtable) GetGoName(yname string) string {
    if yname == "cvCallLegRateWMEntry" { return "Cvcalllegratewmentry" }
    return ""
}

func (cvcalllegratewmtable *CISCOVOICEDIALCONTROLMIB_Cvcalllegratewmtable) GetSegmentPath() string {
    return "cvCallLegRateWMTable"
}

func (cvcalllegratewmtable *CISCOVOICEDIALCONTROLMIB_Cvcalllegratewmtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cvCallLegRateWMEntry" {
        for _, c := range cvcalllegratewmtable.Cvcalllegratewmentry {
            if cvcalllegratewmtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOVOICEDIALCONTROLMIB_Cvcalllegratewmtable_Cvcalllegratewmentry{}
        cvcalllegratewmtable.Cvcalllegratewmentry = append(cvcalllegratewmtable.Cvcalllegratewmentry, child)
        return &cvcalllegratewmtable.Cvcalllegratewmentry[len(cvcalllegratewmtable.Cvcalllegratewmentry)-1]
    }
    return nil
}

func (cvcalllegratewmtable *CISCOVOICEDIALCONTROLMIB_Cvcalllegratewmtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cvcalllegratewmtable.Cvcalllegratewmentry {
        children[cvcalllegratewmtable.Cvcalllegratewmentry[i].GetSegmentPath()] = &cvcalllegratewmtable.Cvcalllegratewmentry[i]
    }
    return children
}

func (cvcalllegratewmtable *CISCOVOICEDIALCONTROLMIB_Cvcalllegratewmtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cvcalllegratewmtable *CISCOVOICEDIALCONTROLMIB_Cvcalllegratewmtable) GetBundleName() string { return "cisco_ios_xe" }

func (cvcalllegratewmtable *CISCOVOICEDIALCONTROLMIB_Cvcalllegratewmtable) GetYangName() string { return "cvCallLegRateWMTable" }

func (cvcalllegratewmtable *CISCOVOICEDIALCONTROLMIB_Cvcalllegratewmtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvcalllegratewmtable *CISCOVOICEDIALCONTROLMIB_Cvcalllegratewmtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvcalllegratewmtable *CISCOVOICEDIALCONTROLMIB_Cvcalllegratewmtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvcalllegratewmtable *CISCOVOICEDIALCONTROLMIB_Cvcalllegratewmtable) SetParent(parent types.Entity) { cvcalllegratewmtable.parent = parent }

func (cvcalllegratewmtable *CISCOVOICEDIALCONTROLMIB_Cvcalllegratewmtable) GetParent() types.Entity { return cvcalllegratewmtable.parent }

func (cvcalllegratewmtable *CISCOVOICEDIALCONTROLMIB_Cvcalllegratewmtable) GetParentYangName() string { return "CISCO-VOICE-DIAL-CONTROL-MIB" }

// CISCOVOICEDIALCONTROLMIB_Cvcalllegratewmtable_Cvcalllegratewmentry
// This is a conceptual-row in cvCallLegRateWMTable
// This entry is created at the system initialization and is
// updated whenever 
// a) This entry is obsolete OR
// b) A new/higher entry is available.
// These entries are reinitialised/added/deleted  if
// cvCallVolumeWMTableSize is changed
type CISCOVOICEDIALCONTROLMIB_Cvcalllegratewmtable_Cvcalllegratewmentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The Object indexes in call leg rate Water mark
    // Table to select one among four interval-tables.  The different types in
    // this table are represented by  CvCallVolumeWMIntvlType. The type is
    // CvCallVolumeWMIntvlType.
    Cvcalllegratewmintvldurunits interface{}

    // This attribute is a key. This is an index that references to different
    // peaks in past period in call leg rate watermark table.  The number of
    // watermarks entries stored for each table are  based on
    // cvCallVolumeWMTableSize. The type is interface{} with range: 1..10.
    Cvcalllegratewmindex interface{}

    // This object indicates high watermark value achieved by the call legs per
    // second for the given interval. The type is interface{} with range:
    // 0..4294967295. Units are call legs per second.
    Cvcalllegratewmvalue interface{}

    // This object indicates date and time when the high watermark is achieved for
    // call-legs per second for the given interval. The type is string.
    Cvcalllegratewmts interface{}
}

func (cvcalllegratewmentry *CISCOVOICEDIALCONTROLMIB_Cvcalllegratewmtable_Cvcalllegratewmentry) GetFilter() yfilter.YFilter { return cvcalllegratewmentry.YFilter }

func (cvcalllegratewmentry *CISCOVOICEDIALCONTROLMIB_Cvcalllegratewmtable_Cvcalllegratewmentry) SetFilter(yf yfilter.YFilter) { cvcalllegratewmentry.YFilter = yf }

func (cvcalllegratewmentry *CISCOVOICEDIALCONTROLMIB_Cvcalllegratewmtable_Cvcalllegratewmentry) GetGoName(yname string) string {
    if yname == "cvCallLegRateWMIntvlDurUnits" { return "Cvcalllegratewmintvldurunits" }
    if yname == "cvCallLegRateWMIndex" { return "Cvcalllegratewmindex" }
    if yname == "cvCallLegRateWMValue" { return "Cvcalllegratewmvalue" }
    if yname == "cvCallLegRateWMts" { return "Cvcalllegratewmts" }
    return ""
}

func (cvcalllegratewmentry *CISCOVOICEDIALCONTROLMIB_Cvcalllegratewmtable_Cvcalllegratewmentry) GetSegmentPath() string {
    return "cvCallLegRateWMEntry" + "[cvCallLegRateWMIntvlDurUnits='" + fmt.Sprintf("%v", cvcalllegratewmentry.Cvcalllegratewmintvldurunits) + "']" + "[cvCallLegRateWMIndex='" + fmt.Sprintf("%v", cvcalllegratewmentry.Cvcalllegratewmindex) + "']"
}

func (cvcalllegratewmentry *CISCOVOICEDIALCONTROLMIB_Cvcalllegratewmtable_Cvcalllegratewmentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cvcalllegratewmentry *CISCOVOICEDIALCONTROLMIB_Cvcalllegratewmtable_Cvcalllegratewmentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cvcalllegratewmentry *CISCOVOICEDIALCONTROLMIB_Cvcalllegratewmtable_Cvcalllegratewmentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cvCallLegRateWMIntvlDurUnits"] = cvcalllegratewmentry.Cvcalllegratewmintvldurunits
    leafs["cvCallLegRateWMIndex"] = cvcalllegratewmentry.Cvcalllegratewmindex
    leafs["cvCallLegRateWMValue"] = cvcalllegratewmentry.Cvcalllegratewmvalue
    leafs["cvCallLegRateWMts"] = cvcalllegratewmentry.Cvcalllegratewmts
    return leafs
}

func (cvcalllegratewmentry *CISCOVOICEDIALCONTROLMIB_Cvcalllegratewmtable_Cvcalllegratewmentry) GetBundleName() string { return "cisco_ios_xe" }

func (cvcalllegratewmentry *CISCOVOICEDIALCONTROLMIB_Cvcalllegratewmtable_Cvcalllegratewmentry) GetYangName() string { return "cvCallLegRateWMEntry" }

func (cvcalllegratewmentry *CISCOVOICEDIALCONTROLMIB_Cvcalllegratewmtable_Cvcalllegratewmentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvcalllegratewmentry *CISCOVOICEDIALCONTROLMIB_Cvcalllegratewmtable_Cvcalllegratewmentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvcalllegratewmentry *CISCOVOICEDIALCONTROLMIB_Cvcalllegratewmtable_Cvcalllegratewmentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvcalllegratewmentry *CISCOVOICEDIALCONTROLMIB_Cvcalllegratewmtable_Cvcalllegratewmentry) SetParent(parent types.Entity) { cvcalllegratewmentry.parent = parent }

func (cvcalllegratewmentry *CISCOVOICEDIALCONTROLMIB_Cvcalllegratewmtable_Cvcalllegratewmentry) GetParent() types.Entity { return cvcalllegratewmentry.parent }

func (cvcalllegratewmentry *CISCOVOICEDIALCONTROLMIB_Cvcalllegratewmtable_Cvcalllegratewmentry) GetParentYangName() string { return "cvCallLegRateWMTable" }

// CISCOVOICEDIALCONTROLMIB_Cvactivecallwmtable
// This table represents high watermarks achieved
// by active calls in various interval length defined 
// by CvCallVolumeWMIntvlType. 
// 
// Each interval may contain one or more entries to allow 
// for detailed measurement to be collected.
type CISCOVOICEDIALCONTROLMIB_Cvactivecallwmtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This is a conceptual-row in cvActiveCallWMTable This entry is created at
    // the system initialization and is updated whenever  a) This entry is
    // obsolete OR b) A new/higher entry is available. These entries are
    // reinitialised/added/deleted  if cvCallVolumeWMTableSize is changed. The
    // type is slice of
    // CISCOVOICEDIALCONTROLMIB_Cvactivecallwmtable_Cvactivecallwmentry.
    Cvactivecallwmentry []CISCOVOICEDIALCONTROLMIB_Cvactivecallwmtable_Cvactivecallwmentry
}

func (cvactivecallwmtable *CISCOVOICEDIALCONTROLMIB_Cvactivecallwmtable) GetFilter() yfilter.YFilter { return cvactivecallwmtable.YFilter }

func (cvactivecallwmtable *CISCOVOICEDIALCONTROLMIB_Cvactivecallwmtable) SetFilter(yf yfilter.YFilter) { cvactivecallwmtable.YFilter = yf }

func (cvactivecallwmtable *CISCOVOICEDIALCONTROLMIB_Cvactivecallwmtable) GetGoName(yname string) string {
    if yname == "cvActiveCallWMEntry" { return "Cvactivecallwmentry" }
    return ""
}

func (cvactivecallwmtable *CISCOVOICEDIALCONTROLMIB_Cvactivecallwmtable) GetSegmentPath() string {
    return "cvActiveCallWMTable"
}

func (cvactivecallwmtable *CISCOVOICEDIALCONTROLMIB_Cvactivecallwmtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cvActiveCallWMEntry" {
        for _, c := range cvactivecallwmtable.Cvactivecallwmentry {
            if cvactivecallwmtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOVOICEDIALCONTROLMIB_Cvactivecallwmtable_Cvactivecallwmentry{}
        cvactivecallwmtable.Cvactivecallwmentry = append(cvactivecallwmtable.Cvactivecallwmentry, child)
        return &cvactivecallwmtable.Cvactivecallwmentry[len(cvactivecallwmtable.Cvactivecallwmentry)-1]
    }
    return nil
}

func (cvactivecallwmtable *CISCOVOICEDIALCONTROLMIB_Cvactivecallwmtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cvactivecallwmtable.Cvactivecallwmentry {
        children[cvactivecallwmtable.Cvactivecallwmentry[i].GetSegmentPath()] = &cvactivecallwmtable.Cvactivecallwmentry[i]
    }
    return children
}

func (cvactivecallwmtable *CISCOVOICEDIALCONTROLMIB_Cvactivecallwmtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cvactivecallwmtable *CISCOVOICEDIALCONTROLMIB_Cvactivecallwmtable) GetBundleName() string { return "cisco_ios_xe" }

func (cvactivecallwmtable *CISCOVOICEDIALCONTROLMIB_Cvactivecallwmtable) GetYangName() string { return "cvActiveCallWMTable" }

func (cvactivecallwmtable *CISCOVOICEDIALCONTROLMIB_Cvactivecallwmtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvactivecallwmtable *CISCOVOICEDIALCONTROLMIB_Cvactivecallwmtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvactivecallwmtable *CISCOVOICEDIALCONTROLMIB_Cvactivecallwmtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvactivecallwmtable *CISCOVOICEDIALCONTROLMIB_Cvactivecallwmtable) SetParent(parent types.Entity) { cvactivecallwmtable.parent = parent }

func (cvactivecallwmtable *CISCOVOICEDIALCONTROLMIB_Cvactivecallwmtable) GetParent() types.Entity { return cvactivecallwmtable.parent }

func (cvactivecallwmtable *CISCOVOICEDIALCONTROLMIB_Cvactivecallwmtable) GetParentYangName() string { return "CISCO-VOICE-DIAL-CONTROL-MIB" }

// CISCOVOICEDIALCONTROLMIB_Cvactivecallwmtable_Cvactivecallwmentry
// This is a conceptual-row in cvActiveCallWMTable
// This entry is created at the system initialization and is
// updated whenever 
// a) This entry is obsolete OR
// b) A new/higher entry is available.
// These entries are reinitialised/added/deleted  if
// cvCallVolumeWMTableSize is changed
type CISCOVOICEDIALCONTROLMIB_Cvactivecallwmtable_Cvactivecallwmentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The Object indexes in active call Water mark Table
    // to select one among four interval-tables.  The different types in this
    // table are represented by  CvCallVolumeWMIntvlType. The type is
    // CvCallVolumeWMIntvlType.
    Cvactivecallwmintvldurunits interface{}

    // This attribute is a key. This is an index that references to different
    // peaks in past period in acive call watermark table.  The number of
    // watermarks entries stored for each table are  based on
    // cvCallVolumeWMTableSize. The type is interface{} with range: 1..10.
    Cvactivecallwmindex interface{}

    // This object indicates high watermark value achieved by the active calls for
    // the given interval. The type is interface{} with range: 0..4294967295.
    // Units are calls.
    Cvactivecallwmvalue interface{}

    // This object indicates date and time when the high watermark is achieved for
    // active calls for the given interval. The type is string.
    Cvactivecallwmts interface{}
}

func (cvactivecallwmentry *CISCOVOICEDIALCONTROLMIB_Cvactivecallwmtable_Cvactivecallwmentry) GetFilter() yfilter.YFilter { return cvactivecallwmentry.YFilter }

func (cvactivecallwmentry *CISCOVOICEDIALCONTROLMIB_Cvactivecallwmtable_Cvactivecallwmentry) SetFilter(yf yfilter.YFilter) { cvactivecallwmentry.YFilter = yf }

func (cvactivecallwmentry *CISCOVOICEDIALCONTROLMIB_Cvactivecallwmtable_Cvactivecallwmentry) GetGoName(yname string) string {
    if yname == "cvActiveCallWMIntvlDurUnits" { return "Cvactivecallwmintvldurunits" }
    if yname == "cvActiveCallWMIndex" { return "Cvactivecallwmindex" }
    if yname == "cvActiveCallWMValue" { return "Cvactivecallwmvalue" }
    if yname == "cvActiveCallWMts" { return "Cvactivecallwmts" }
    return ""
}

func (cvactivecallwmentry *CISCOVOICEDIALCONTROLMIB_Cvactivecallwmtable_Cvactivecallwmentry) GetSegmentPath() string {
    return "cvActiveCallWMEntry" + "[cvActiveCallWMIntvlDurUnits='" + fmt.Sprintf("%v", cvactivecallwmentry.Cvactivecallwmintvldurunits) + "']" + "[cvActiveCallWMIndex='" + fmt.Sprintf("%v", cvactivecallwmentry.Cvactivecallwmindex) + "']"
}

func (cvactivecallwmentry *CISCOVOICEDIALCONTROLMIB_Cvactivecallwmtable_Cvactivecallwmentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cvactivecallwmentry *CISCOVOICEDIALCONTROLMIB_Cvactivecallwmtable_Cvactivecallwmentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cvactivecallwmentry *CISCOVOICEDIALCONTROLMIB_Cvactivecallwmtable_Cvactivecallwmentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cvActiveCallWMIntvlDurUnits"] = cvactivecallwmentry.Cvactivecallwmintvldurunits
    leafs["cvActiveCallWMIndex"] = cvactivecallwmentry.Cvactivecallwmindex
    leafs["cvActiveCallWMValue"] = cvactivecallwmentry.Cvactivecallwmvalue
    leafs["cvActiveCallWMts"] = cvactivecallwmentry.Cvactivecallwmts
    return leafs
}

func (cvactivecallwmentry *CISCOVOICEDIALCONTROLMIB_Cvactivecallwmtable_Cvactivecallwmentry) GetBundleName() string { return "cisco_ios_xe" }

func (cvactivecallwmentry *CISCOVOICEDIALCONTROLMIB_Cvactivecallwmtable_Cvactivecallwmentry) GetYangName() string { return "cvActiveCallWMEntry" }

func (cvactivecallwmentry *CISCOVOICEDIALCONTROLMIB_Cvactivecallwmtable_Cvactivecallwmentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvactivecallwmentry *CISCOVOICEDIALCONTROLMIB_Cvactivecallwmtable_Cvactivecallwmentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvactivecallwmentry *CISCOVOICEDIALCONTROLMIB_Cvactivecallwmtable_Cvactivecallwmentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvactivecallwmentry *CISCOVOICEDIALCONTROLMIB_Cvactivecallwmtable_Cvactivecallwmentry) SetParent(parent types.Entity) { cvactivecallwmentry.parent = parent }

func (cvactivecallwmentry *CISCOVOICEDIALCONTROLMIB_Cvactivecallwmtable_Cvactivecallwmentry) GetParent() types.Entity { return cvactivecallwmentry.parent }

func (cvactivecallwmentry *CISCOVOICEDIALCONTROLMIB_Cvactivecallwmtable_Cvactivecallwmentry) GetParentYangName() string { return "cvActiveCallWMTable" }

// CISCOVOICEDIALCONTROLMIB_Cvsipmsgratewmtable
// This table represents of high watermarks achieved
// by SIP message rate in various interval length defined 
// by CvCallVolumeWMIntvlType. 
// 
// Each interval may contain one or more entries to allow for
// detailed measurement to be collected
type CISCOVOICEDIALCONTROLMIB_Cvsipmsgratewmtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This is a conceptual-row in cvSipMsgRateWMTable. This entry is created at
    // the system initialization and is updated whenever  a) This entry is
    // obsolete OR b) A new/higher entry is available. These entries are
    // reinitialised/added/deleted if cvCallVolumeWMTableSize is changed. The type
    // is slice of
    // CISCOVOICEDIALCONTROLMIB_Cvsipmsgratewmtable_Cvsipmsgratewmentry.
    Cvsipmsgratewmentry []CISCOVOICEDIALCONTROLMIB_Cvsipmsgratewmtable_Cvsipmsgratewmentry
}

func (cvsipmsgratewmtable *CISCOVOICEDIALCONTROLMIB_Cvsipmsgratewmtable) GetFilter() yfilter.YFilter { return cvsipmsgratewmtable.YFilter }

func (cvsipmsgratewmtable *CISCOVOICEDIALCONTROLMIB_Cvsipmsgratewmtable) SetFilter(yf yfilter.YFilter) { cvsipmsgratewmtable.YFilter = yf }

func (cvsipmsgratewmtable *CISCOVOICEDIALCONTROLMIB_Cvsipmsgratewmtable) GetGoName(yname string) string {
    if yname == "cvSipMsgRateWMEntry" { return "Cvsipmsgratewmentry" }
    return ""
}

func (cvsipmsgratewmtable *CISCOVOICEDIALCONTROLMIB_Cvsipmsgratewmtable) GetSegmentPath() string {
    return "cvSipMsgRateWMTable"
}

func (cvsipmsgratewmtable *CISCOVOICEDIALCONTROLMIB_Cvsipmsgratewmtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cvSipMsgRateWMEntry" {
        for _, c := range cvsipmsgratewmtable.Cvsipmsgratewmentry {
            if cvsipmsgratewmtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOVOICEDIALCONTROLMIB_Cvsipmsgratewmtable_Cvsipmsgratewmentry{}
        cvsipmsgratewmtable.Cvsipmsgratewmentry = append(cvsipmsgratewmtable.Cvsipmsgratewmentry, child)
        return &cvsipmsgratewmtable.Cvsipmsgratewmentry[len(cvsipmsgratewmtable.Cvsipmsgratewmentry)-1]
    }
    return nil
}

func (cvsipmsgratewmtable *CISCOVOICEDIALCONTROLMIB_Cvsipmsgratewmtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cvsipmsgratewmtable.Cvsipmsgratewmentry {
        children[cvsipmsgratewmtable.Cvsipmsgratewmentry[i].GetSegmentPath()] = &cvsipmsgratewmtable.Cvsipmsgratewmentry[i]
    }
    return children
}

func (cvsipmsgratewmtable *CISCOVOICEDIALCONTROLMIB_Cvsipmsgratewmtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cvsipmsgratewmtable *CISCOVOICEDIALCONTROLMIB_Cvsipmsgratewmtable) GetBundleName() string { return "cisco_ios_xe" }

func (cvsipmsgratewmtable *CISCOVOICEDIALCONTROLMIB_Cvsipmsgratewmtable) GetYangName() string { return "cvSipMsgRateWMTable" }

func (cvsipmsgratewmtable *CISCOVOICEDIALCONTROLMIB_Cvsipmsgratewmtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvsipmsgratewmtable *CISCOVOICEDIALCONTROLMIB_Cvsipmsgratewmtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvsipmsgratewmtable *CISCOVOICEDIALCONTROLMIB_Cvsipmsgratewmtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvsipmsgratewmtable *CISCOVOICEDIALCONTROLMIB_Cvsipmsgratewmtable) SetParent(parent types.Entity) { cvsipmsgratewmtable.parent = parent }

func (cvsipmsgratewmtable *CISCOVOICEDIALCONTROLMIB_Cvsipmsgratewmtable) GetParent() types.Entity { return cvsipmsgratewmtable.parent }

func (cvsipmsgratewmtable *CISCOVOICEDIALCONTROLMIB_Cvsipmsgratewmtable) GetParentYangName() string { return "CISCO-VOICE-DIAL-CONTROL-MIB" }

// CISCOVOICEDIALCONTROLMIB_Cvsipmsgratewmtable_Cvsipmsgratewmentry
// This is a conceptual-row in cvSipMsgRateWMTable.
// This entry is created at the system initialization and is
// updated whenever 
// a) This entry is obsolete OR
// b) A new/higher entry is available.
// These entries are reinitialised/added/deleted if
// cvCallVolumeWMTableSize is changed
type CISCOVOICEDIALCONTROLMIB_Cvsipmsgratewmtable_Cvsipmsgratewmentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The Object indexes in SIP Message rate Water mark
    // Table to select one among four interval-tables.  The different types in
    // this table are represented by  CvCallVolumeWMIntvlType. The type is
    // CvCallVolumeWMIntvlType.
    Cvsipmsgratewmintvldurunits interface{}

    // This attribute is a key. This is an index that references to different
    // peaks in past period in sip message rate watermark table.  The number of
    // watermarks entries stored for each table are  based on
    // cvCallVolumeWMTableSize. The type is interface{} with range: 1..10.
    Cvsipmsgratewmindex interface{}

    // This object indicates high watermark value achieved by the SIP messages per
    // second for the given interval. The type is interface{} with range:
    // 0..4294967295. Units are SIP messages per second.
    Cvsipmsgratewmvalue interface{}

    // This object indicates date and time when the high watermark is achieved for
    // SIP messages per second for the given interval. The type is string.
    Cvsipmsgratewmts interface{}
}

func (cvsipmsgratewmentry *CISCOVOICEDIALCONTROLMIB_Cvsipmsgratewmtable_Cvsipmsgratewmentry) GetFilter() yfilter.YFilter { return cvsipmsgratewmentry.YFilter }

func (cvsipmsgratewmentry *CISCOVOICEDIALCONTROLMIB_Cvsipmsgratewmtable_Cvsipmsgratewmentry) SetFilter(yf yfilter.YFilter) { cvsipmsgratewmentry.YFilter = yf }

func (cvsipmsgratewmentry *CISCOVOICEDIALCONTROLMIB_Cvsipmsgratewmtable_Cvsipmsgratewmentry) GetGoName(yname string) string {
    if yname == "cvSipMsgRateWMIntvlDurUnits" { return "Cvsipmsgratewmintvldurunits" }
    if yname == "cvSipMsgRateWMIndex" { return "Cvsipmsgratewmindex" }
    if yname == "cvSipMsgRateWMValue" { return "Cvsipmsgratewmvalue" }
    if yname == "cvSipMsgRateWMts" { return "Cvsipmsgratewmts" }
    return ""
}

func (cvsipmsgratewmentry *CISCOVOICEDIALCONTROLMIB_Cvsipmsgratewmtable_Cvsipmsgratewmentry) GetSegmentPath() string {
    return "cvSipMsgRateWMEntry" + "[cvSipMsgRateWMIntvlDurUnits='" + fmt.Sprintf("%v", cvsipmsgratewmentry.Cvsipmsgratewmintvldurunits) + "']" + "[cvSipMsgRateWMIndex='" + fmt.Sprintf("%v", cvsipmsgratewmentry.Cvsipmsgratewmindex) + "']"
}

func (cvsipmsgratewmentry *CISCOVOICEDIALCONTROLMIB_Cvsipmsgratewmtable_Cvsipmsgratewmentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cvsipmsgratewmentry *CISCOVOICEDIALCONTROLMIB_Cvsipmsgratewmtable_Cvsipmsgratewmentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cvsipmsgratewmentry *CISCOVOICEDIALCONTROLMIB_Cvsipmsgratewmtable_Cvsipmsgratewmentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cvSipMsgRateWMIntvlDurUnits"] = cvsipmsgratewmentry.Cvsipmsgratewmintvldurunits
    leafs["cvSipMsgRateWMIndex"] = cvsipmsgratewmentry.Cvsipmsgratewmindex
    leafs["cvSipMsgRateWMValue"] = cvsipmsgratewmentry.Cvsipmsgratewmvalue
    leafs["cvSipMsgRateWMts"] = cvsipmsgratewmentry.Cvsipmsgratewmts
    return leafs
}

func (cvsipmsgratewmentry *CISCOVOICEDIALCONTROLMIB_Cvsipmsgratewmtable_Cvsipmsgratewmentry) GetBundleName() string { return "cisco_ios_xe" }

func (cvsipmsgratewmentry *CISCOVOICEDIALCONTROLMIB_Cvsipmsgratewmtable_Cvsipmsgratewmentry) GetYangName() string { return "cvSipMsgRateWMEntry" }

func (cvsipmsgratewmentry *CISCOVOICEDIALCONTROLMIB_Cvsipmsgratewmtable_Cvsipmsgratewmentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvsipmsgratewmentry *CISCOVOICEDIALCONTROLMIB_Cvsipmsgratewmtable_Cvsipmsgratewmentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvsipmsgratewmentry *CISCOVOICEDIALCONTROLMIB_Cvsipmsgratewmtable_Cvsipmsgratewmentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvsipmsgratewmentry *CISCOVOICEDIALCONTROLMIB_Cvsipmsgratewmtable_Cvsipmsgratewmentry) SetParent(parent types.Entity) { cvsipmsgratewmentry.parent = parent }

func (cvsipmsgratewmentry *CISCOVOICEDIALCONTROLMIB_Cvsipmsgratewmtable_Cvsipmsgratewmentry) GetParent() types.Entity { return cvsipmsgratewmentry.parent }

func (cvsipmsgratewmentry *CISCOVOICEDIALCONTROLMIB_Cvsipmsgratewmtable_Cvsipmsgratewmentry) GetParentYangName() string { return "cvSipMsgRateWMTable" }

