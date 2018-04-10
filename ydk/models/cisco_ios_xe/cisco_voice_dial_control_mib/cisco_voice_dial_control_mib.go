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

// CvCallVolumeStatsIntvlType represents 3 : Hours Table: Here each entry corresponds to an hour
type CvCallVolumeStatsIntvlType string

const (
    CvCallVolumeStatsIntvlType_secondStats CvCallVolumeStatsIntvlType = "secondStats"

    CvCallVolumeStatsIntvlType_minuteStats CvCallVolumeStatsIntvlType = "minuteStats"

    CvCallVolumeStatsIntvlType_hourStats CvCallVolumeStatsIntvlType = "hourStats"
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

// CvAmrNbRtpEncap represents Represents GSM AMR-NB codec RTP encapsulation type.
type CvAmrNbRtpEncap string

const (
    CvAmrNbRtpEncap_rfc3267 CvAmrNbRtpEncap = "rfc3267"
)

// CvIlbcFrameMode represents              frame is packetized in 50 bytes.
type CvIlbcFrameMode string

const (
    CvIlbcFrameMode_frameMode20 CvIlbcFrameMode = "frameMode20"

    CvIlbcFrameMode_frameMode30 CvIlbcFrameMode = "frameMode30"
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

// CISCOVOICEDIALCONTROLMIB
type CISCOVOICEDIALCONTROLMIB struct {
    EntityData types.CommonEntityData
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

func (cISCOVOICEDIALCONTROLMIB *CISCOVOICEDIALCONTROLMIB) GetEntityData() *types.CommonEntityData {
    cISCOVOICEDIALCONTROLMIB.EntityData.YFilter = cISCOVOICEDIALCONTROLMIB.YFilter
    cISCOVOICEDIALCONTROLMIB.EntityData.YangName = "CISCO-VOICE-DIAL-CONTROL-MIB"
    cISCOVOICEDIALCONTROLMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOVOICEDIALCONTROLMIB.EntityData.ParentYangName = "CISCO-VOICE-DIAL-CONTROL-MIB"
    cISCOVOICEDIALCONTROLMIB.EntityData.SegmentPath = "CISCO-VOICE-DIAL-CONTROL-MIB:CISCO-VOICE-DIAL-CONTROL-MIB"
    cISCOVOICEDIALCONTROLMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOVOICEDIALCONTROLMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOVOICEDIALCONTROLMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOVOICEDIALCONTROLMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOVOICEDIALCONTROLMIB.EntityData.Children["cvGeneralConfiguration"] = types.YChild{"Cvgeneralconfiguration", &cISCOVOICEDIALCONTROLMIB.Cvgeneralconfiguration}
    cISCOVOICEDIALCONTROLMIB.EntityData.Children["cvGatewayCallActive"] = types.YChild{"Cvgatewaycallactive", &cISCOVOICEDIALCONTROLMIB.Cvgatewaycallactive}
    cISCOVOICEDIALCONTROLMIB.EntityData.Children["cvCallVolume"] = types.YChild{"Cvcallvolume", &cISCOVOICEDIALCONTROLMIB.Cvcallvolume}
    cISCOVOICEDIALCONTROLMIB.EntityData.Children["cvCallRateMonitor"] = types.YChild{"Cvcallratemonitor", &cISCOVOICEDIALCONTROLMIB.Cvcallratemonitor}
    cISCOVOICEDIALCONTROLMIB.EntityData.Children["cvCallVolumeStatsHistory"] = types.YChild{"Cvcallvolumestatshistory", &cISCOVOICEDIALCONTROLMIB.Cvcallvolumestatshistory}
    cISCOVOICEDIALCONTROLMIB.EntityData.Children["cvPeerCfgTable"] = types.YChild{"Cvpeercfgtable", &cISCOVOICEDIALCONTROLMIB.Cvpeercfgtable}
    cISCOVOICEDIALCONTROLMIB.EntityData.Children["cvVoicePeerCfgTable"] = types.YChild{"Cvvoicepeercfgtable", &cISCOVOICEDIALCONTROLMIB.Cvvoicepeercfgtable}
    cISCOVOICEDIALCONTROLMIB.EntityData.Children["cvVoIPPeerCfgTable"] = types.YChild{"Cvvoippeercfgtable", &cISCOVOICEDIALCONTROLMIB.Cvvoippeercfgtable}
    cISCOVOICEDIALCONTROLMIB.EntityData.Children["cvPeerCommonCfgTable"] = types.YChild{"Cvpeercommoncfgtable", &cISCOVOICEDIALCONTROLMIB.Cvpeercommoncfgtable}
    cISCOVOICEDIALCONTROLMIB.EntityData.Children["cvCallActiveTable"] = types.YChild{"Cvcallactivetable", &cISCOVOICEDIALCONTROLMIB.Cvcallactivetable}
    cISCOVOICEDIALCONTROLMIB.EntityData.Children["cvVoIPCallActiveTable"] = types.YChild{"Cvvoipcallactivetable", &cISCOVOICEDIALCONTROLMIB.Cvvoipcallactivetable}
    cISCOVOICEDIALCONTROLMIB.EntityData.Children["cvCallVolConnTable"] = types.YChild{"Cvcallvolconntable", &cISCOVOICEDIALCONTROLMIB.Cvcallvolconntable}
    cISCOVOICEDIALCONTROLMIB.EntityData.Children["cvCallVolIfTable"] = types.YChild{"Cvcallvoliftable", &cISCOVOICEDIALCONTROLMIB.Cvcallvoliftable}
    cISCOVOICEDIALCONTROLMIB.EntityData.Children["cvCallHistoryTable"] = types.YChild{"Cvcallhistorytable", &cISCOVOICEDIALCONTROLMIB.Cvcallhistorytable}
    cISCOVOICEDIALCONTROLMIB.EntityData.Children["cvVoIPCallHistoryTable"] = types.YChild{"Cvvoipcallhistorytable", &cISCOVOICEDIALCONTROLMIB.Cvvoipcallhistorytable}
    cISCOVOICEDIALCONTROLMIB.EntityData.Children["cvCallRateStatsTable"] = types.YChild{"Cvcallratestatstable", &cISCOVOICEDIALCONTROLMIB.Cvcallratestatstable}
    cISCOVOICEDIALCONTROLMIB.EntityData.Children["cvCallLegRateStatsTable"] = types.YChild{"Cvcalllegratestatstable", &cISCOVOICEDIALCONTROLMIB.Cvcalllegratestatstable}
    cISCOVOICEDIALCONTROLMIB.EntityData.Children["cvActiveCallStatsTable"] = types.YChild{"Cvactivecallstatstable", &cISCOVOICEDIALCONTROLMIB.Cvactivecallstatstable}
    cISCOVOICEDIALCONTROLMIB.EntityData.Children["cvCallDurationStatsTable"] = types.YChild{"Cvcalldurationstatstable", &cISCOVOICEDIALCONTROLMIB.Cvcalldurationstatstable}
    cISCOVOICEDIALCONTROLMIB.EntityData.Children["cvSipMsgRateStatsTable"] = types.YChild{"Cvsipmsgratestatstable", &cISCOVOICEDIALCONTROLMIB.Cvsipmsgratestatstable}
    cISCOVOICEDIALCONTROLMIB.EntityData.Children["cvCallRateWMTable"] = types.YChild{"Cvcallratewmtable", &cISCOVOICEDIALCONTROLMIB.Cvcallratewmtable}
    cISCOVOICEDIALCONTROLMIB.EntityData.Children["cvCallLegRateWMTable"] = types.YChild{"Cvcalllegratewmtable", &cISCOVOICEDIALCONTROLMIB.Cvcalllegratewmtable}
    cISCOVOICEDIALCONTROLMIB.EntityData.Children["cvActiveCallWMTable"] = types.YChild{"Cvactivecallwmtable", &cISCOVOICEDIALCONTROLMIB.Cvactivecallwmtable}
    cISCOVOICEDIALCONTROLMIB.EntityData.Children["cvSipMsgRateWMTable"] = types.YChild{"Cvsipmsgratewmtable", &cISCOVOICEDIALCONTROLMIB.Cvsipmsgratewmtable}
    cISCOVOICEDIALCONTROLMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOVOICEDIALCONTROLMIB.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_Cvgeneralconfiguration
type CISCOVOICEDIALCONTROLMIB_Cvgeneralconfiguration struct {
    EntityData types.CommonEntityData
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

func (cvgeneralconfiguration *CISCOVOICEDIALCONTROLMIB_Cvgeneralconfiguration) GetEntityData() *types.CommonEntityData {
    cvgeneralconfiguration.EntityData.YFilter = cvgeneralconfiguration.YFilter
    cvgeneralconfiguration.EntityData.YangName = "cvGeneralConfiguration"
    cvgeneralconfiguration.EntityData.BundleName = "cisco_ios_xe"
    cvgeneralconfiguration.EntityData.ParentYangName = "CISCO-VOICE-DIAL-CONTROL-MIB"
    cvgeneralconfiguration.EntityData.SegmentPath = "cvGeneralConfiguration"
    cvgeneralconfiguration.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvgeneralconfiguration.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvgeneralconfiguration.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvgeneralconfiguration.EntityData.Children = make(map[string]types.YChild)
    cvgeneralconfiguration.EntityData.Leafs = make(map[string]types.YLeaf)
    cvgeneralconfiguration.EntityData.Leafs["cvGeneralPoorQoVNotificationEnable"] = types.YLeaf{"Cvgeneralpoorqovnotificationenable", cvgeneralconfiguration.Cvgeneralpoorqovnotificationenable}
    cvgeneralconfiguration.EntityData.Leafs["cvGeneralFallbackNotificationEnable"] = types.YLeaf{"Cvgeneralfallbacknotificationenable", cvgeneralconfiguration.Cvgeneralfallbacknotificationenable}
    cvgeneralconfiguration.EntityData.Leafs["cvGeneralDSCPPolicyNotificationEnable"] = types.YLeaf{"Cvgeneraldscppolicynotificationenable", cvgeneralconfiguration.Cvgeneraldscppolicynotificationenable}
    cvgeneralconfiguration.EntityData.Leafs["cvGeneralMediaPolicyNotificationEnable"] = types.YLeaf{"Cvgeneralmediapolicynotificationenable", cvgeneralconfiguration.Cvgeneralmediapolicynotificationenable}
    return &(cvgeneralconfiguration.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_Cvgatewaycallactive
type CISCOVOICEDIALCONTROLMIB_Cvgatewaycallactive struct {
    EntityData types.CommonEntityData
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

func (cvgatewaycallactive *CISCOVOICEDIALCONTROLMIB_Cvgatewaycallactive) GetEntityData() *types.CommonEntityData {
    cvgatewaycallactive.EntityData.YFilter = cvgatewaycallactive.YFilter
    cvgatewaycallactive.EntityData.YangName = "cvGatewayCallActive"
    cvgatewaycallactive.EntityData.BundleName = "cisco_ios_xe"
    cvgatewaycallactive.EntityData.ParentYangName = "CISCO-VOICE-DIAL-CONTROL-MIB"
    cvgatewaycallactive.EntityData.SegmentPath = "cvGatewayCallActive"
    cvgatewaycallactive.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvgatewaycallactive.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvgatewaycallactive.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvgatewaycallactive.EntityData.Children = make(map[string]types.YChild)
    cvgatewaycallactive.EntityData.Leafs = make(map[string]types.YLeaf)
    cvgatewaycallactive.EntityData.Leafs["cvCallActiveDS0s"] = types.YLeaf{"Cvcallactiveds0S", cvgatewaycallactive.Cvcallactiveds0S}
    cvgatewaycallactive.EntityData.Leafs["cvCallActiveDS0sHighThreshold"] = types.YLeaf{"Cvcallactiveds0Shighthreshold", cvgatewaycallactive.Cvcallactiveds0Shighthreshold}
    cvgatewaycallactive.EntityData.Leafs["cvCallActiveDS0sLowThreshold"] = types.YLeaf{"Cvcallactiveds0Slowthreshold", cvgatewaycallactive.Cvcallactiveds0Slowthreshold}
    cvgatewaycallactive.EntityData.Leafs["cvCallActiveDS0sHighNotifyEnable"] = types.YLeaf{"Cvcallactiveds0Shighnotifyenable", cvgatewaycallactive.Cvcallactiveds0Shighnotifyenable}
    cvgatewaycallactive.EntityData.Leafs["cvCallActiveDS0sLowNotifyEnable"] = types.YLeaf{"Cvcallactiveds0Slownotifyenable", cvgatewaycallactive.Cvcallactiveds0Slownotifyenable}
    return &(cvgatewaycallactive.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_Cvcallvolume
type CISCOVOICEDIALCONTROLMIB_Cvcallvolume struct {
    EntityData types.CommonEntityData
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

func (cvcallvolume *CISCOVOICEDIALCONTROLMIB_Cvcallvolume) GetEntityData() *types.CommonEntityData {
    cvcallvolume.EntityData.YFilter = cvcallvolume.YFilter
    cvcallvolume.EntityData.YangName = "cvCallVolume"
    cvcallvolume.EntityData.BundleName = "cisco_ios_xe"
    cvcallvolume.EntityData.ParentYangName = "CISCO-VOICE-DIAL-CONTROL-MIB"
    cvcallvolume.EntityData.SegmentPath = "cvCallVolume"
    cvcallvolume.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvcallvolume.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvcallvolume.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvcallvolume.EntityData.Children = make(map[string]types.YChild)
    cvcallvolume.EntityData.Leafs = make(map[string]types.YLeaf)
    cvcallvolume.EntityData.Leafs["cvCallVolConnTotalActiveConnections"] = types.YLeaf{"Cvcallvolconntotalactiveconnections", cvcallvolume.Cvcallvolconntotalactiveconnections}
    cvcallvolume.EntityData.Leafs["cvCallVolConnMaxCallConnectionLicenese"] = types.YLeaf{"Cvcallvolconnmaxcallconnectionlicenese", cvcallvolume.Cvcallvolconnmaxcallconnectionlicenese}
    return &(cvcallvolume.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_Cvcallratemonitor
type CISCOVOICEDIALCONTROLMIB_Cvcallratemonitor struct {
    EntityData types.CommonEntityData
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

func (cvcallratemonitor *CISCOVOICEDIALCONTROLMIB_Cvcallratemonitor) GetEntityData() *types.CommonEntityData {
    cvcallratemonitor.EntityData.YFilter = cvcallratemonitor.YFilter
    cvcallratemonitor.EntityData.YangName = "cvCallRateMonitor"
    cvcallratemonitor.EntityData.BundleName = "cisco_ios_xe"
    cvcallratemonitor.EntityData.ParentYangName = "CISCO-VOICE-DIAL-CONTROL-MIB"
    cvcallratemonitor.EntityData.SegmentPath = "cvCallRateMonitor"
    cvcallratemonitor.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvcallratemonitor.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvcallratemonitor.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvcallratemonitor.EntityData.Children = make(map[string]types.YChild)
    cvcallratemonitor.EntityData.Leafs = make(map[string]types.YLeaf)
    cvcallratemonitor.EntityData.Leafs["cvCallRateMonitorEnable"] = types.YLeaf{"Cvcallratemonitorenable", cvcallratemonitor.Cvcallratemonitorenable}
    cvcallratemonitor.EntityData.Leafs["cvCallRateMonitorTime"] = types.YLeaf{"Cvcallratemonitortime", cvcallratemonitor.Cvcallratemonitortime}
    cvcallratemonitor.EntityData.Leafs["cvCallRate"] = types.YLeaf{"Cvcallrate", cvcallratemonitor.Cvcallrate}
    cvcallratemonitor.EntityData.Leafs["cvCallRateHiWaterMark"] = types.YLeaf{"Cvcallratehiwatermark", cvcallratemonitor.Cvcallratehiwatermark}
    return &(cvcallratemonitor.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_Cvcallvolumestatshistory
type CISCOVOICEDIALCONTROLMIB_Cvcallvolumestatshistory struct {
    EntityData types.CommonEntityData
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

func (cvcallvolumestatshistory *CISCOVOICEDIALCONTROLMIB_Cvcallvolumestatshistory) GetEntityData() *types.CommonEntityData {
    cvcallvolumestatshistory.EntityData.YFilter = cvcallvolumestatshistory.YFilter
    cvcallvolumestatshistory.EntityData.YangName = "cvCallVolumeStatsHistory"
    cvcallvolumestatshistory.EntityData.BundleName = "cisco_ios_xe"
    cvcallvolumestatshistory.EntityData.ParentYangName = "CISCO-VOICE-DIAL-CONTROL-MIB"
    cvcallvolumestatshistory.EntityData.SegmentPath = "cvCallVolumeStatsHistory"
    cvcallvolumestatshistory.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvcallvolumestatshistory.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvcallvolumestatshistory.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvcallvolumestatshistory.EntityData.Children = make(map[string]types.YChild)
    cvcallvolumestatshistory.EntityData.Leafs = make(map[string]types.YLeaf)
    cvcallvolumestatshistory.EntityData.Leafs["cvCallDurationStatsThreshold"] = types.YLeaf{"Cvcalldurationstatsthreshold", cvcallvolumestatshistory.Cvcalldurationstatsthreshold}
    cvcallvolumestatshistory.EntityData.Leafs["cvCallVolumeWMTableSize"] = types.YLeaf{"Cvcallvolumewmtablesize", cvcallvolumestatshistory.Cvcallvolumewmtablesize}
    return &(cvcallvolumestatshistory.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_Cvpeercfgtable
// The table contains the Voice Generic Peer information that
// is used to create an ifIndexed row with an appropriate
// ifType that is associated with the cvPeerCfgType and
// cvPeerCfgPeerType objects.
type CISCOVOICEDIALCONTROLMIB_Cvpeercfgtable struct {
    EntityData types.CommonEntityData
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

func (cvpeercfgtable *CISCOVOICEDIALCONTROLMIB_Cvpeercfgtable) GetEntityData() *types.CommonEntityData {
    cvpeercfgtable.EntityData.YFilter = cvpeercfgtable.YFilter
    cvpeercfgtable.EntityData.YangName = "cvPeerCfgTable"
    cvpeercfgtable.EntityData.BundleName = "cisco_ios_xe"
    cvpeercfgtable.EntityData.ParentYangName = "CISCO-VOICE-DIAL-CONTROL-MIB"
    cvpeercfgtable.EntityData.SegmentPath = "cvPeerCfgTable"
    cvpeercfgtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvpeercfgtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvpeercfgtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvpeercfgtable.EntityData.Children = make(map[string]types.YChild)
    cvpeercfgtable.EntityData.Children["cvPeerCfgEntry"] = types.YChild{"Cvpeercfgentry", nil}
    for i := range cvpeercfgtable.Cvpeercfgentry {
        cvpeercfgtable.EntityData.Children[types.GetSegmentPath(&cvpeercfgtable.Cvpeercfgentry[i])] = types.YChild{"Cvpeercfgentry", &cvpeercfgtable.Cvpeercfgentry[i]}
    }
    cvpeercfgtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cvpeercfgtable.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (cvpeercfgentry *CISCOVOICEDIALCONTROLMIB_Cvpeercfgtable_Cvpeercfgentry) GetEntityData() *types.CommonEntityData {
    cvpeercfgentry.EntityData.YFilter = cvpeercfgentry.YFilter
    cvpeercfgentry.EntityData.YangName = "cvPeerCfgEntry"
    cvpeercfgentry.EntityData.BundleName = "cisco_ios_xe"
    cvpeercfgentry.EntityData.ParentYangName = "cvPeerCfgTable"
    cvpeercfgentry.EntityData.SegmentPath = "cvPeerCfgEntry" + "[cvPeerCfgIndex='" + fmt.Sprintf("%v", cvpeercfgentry.Cvpeercfgindex) + "']"
    cvpeercfgentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvpeercfgentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvpeercfgentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvpeercfgentry.EntityData.Children = make(map[string]types.YChild)
    cvpeercfgentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cvpeercfgentry.EntityData.Leafs["cvPeerCfgIndex"] = types.YLeaf{"Cvpeercfgindex", cvpeercfgentry.Cvpeercfgindex}
    cvpeercfgentry.EntityData.Leafs["cvPeerCfgIfIndex"] = types.YLeaf{"Cvpeercfgifindex", cvpeercfgentry.Cvpeercfgifindex}
    cvpeercfgentry.EntityData.Leafs["cvPeerCfgType"] = types.YLeaf{"Cvpeercfgtype", cvpeercfgentry.Cvpeercfgtype}
    cvpeercfgentry.EntityData.Leafs["cvPeerCfgRowStatus"] = types.YLeaf{"Cvpeercfgrowstatus", cvpeercfgentry.Cvpeercfgrowstatus}
    cvpeercfgentry.EntityData.Leafs["cvPeerCfgPeerType"] = types.YLeaf{"Cvpeercfgpeertype", cvpeercfgentry.Cvpeercfgpeertype}
    cvpeercfgentry.EntityData.Leafs["cvCallVolPeerIncomingCalls"] = types.YLeaf{"Cvcallvolpeerincomingcalls", cvpeercfgentry.Cvcallvolpeerincomingcalls}
    cvpeercfgentry.EntityData.Leafs["cvCallVolPeerOutgoingCalls"] = types.YLeaf{"Cvcallvolpeeroutgoingcalls", cvpeercfgentry.Cvcallvolpeeroutgoingcalls}
    return &(cvpeercfgentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A single Voice specific Peer. One entry per voice encapsulation. The entry
    // is created when its associated 'voiceEncap(103)' encapsulation ifEntry is
    // created. This entry is deleted when its associated ifEntry is deleted. The
    // type is slice of
    // CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable_Cvvoicepeercfgentry.
    Cvvoicepeercfgentry []CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable_Cvvoicepeercfgentry
}

func (cvvoicepeercfgtable *CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable) GetEntityData() *types.CommonEntityData {
    cvvoicepeercfgtable.EntityData.YFilter = cvvoicepeercfgtable.YFilter
    cvvoicepeercfgtable.EntityData.YangName = "cvVoicePeerCfgTable"
    cvvoicepeercfgtable.EntityData.BundleName = "cisco_ios_xe"
    cvvoicepeercfgtable.EntityData.ParentYangName = "CISCO-VOICE-DIAL-CONTROL-MIB"
    cvvoicepeercfgtable.EntityData.SegmentPath = "cvVoicePeerCfgTable"
    cvvoicepeercfgtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvvoicepeercfgtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvvoicepeercfgtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvvoicepeercfgtable.EntityData.Children = make(map[string]types.YChild)
    cvvoicepeercfgtable.EntityData.Children["cvVoicePeerCfgEntry"] = types.YChild{"Cvvoicepeercfgentry", nil}
    for i := range cvvoicepeercfgtable.Cvvoicepeercfgentry {
        cvvoicepeercfgtable.EntityData.Children[types.GetSegmentPath(&cvvoicepeercfgtable.Cvvoicepeercfgentry[i])] = types.YChild{"Cvvoicepeercfgentry", &cvvoicepeercfgtable.Cvvoicepeercfgentry[i]}
    }
    cvvoicepeercfgtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cvvoicepeercfgtable.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable_Cvvoicepeercfgentry
// A single Voice specific Peer. One entry per voice
// encapsulation.
// The entry is created when its associated 'voiceEncap(103)'
// encapsulation ifEntry is created.
// This entry is deleted when its associated ifEntry is
// deleted.
type CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable_Cvvoicepeercfgentry struct {
    EntityData types.CommonEntityData
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

func (cvvoicepeercfgentry *CISCOVOICEDIALCONTROLMIB_Cvvoicepeercfgtable_Cvvoicepeercfgentry) GetEntityData() *types.CommonEntityData {
    cvvoicepeercfgentry.EntityData.YFilter = cvvoicepeercfgentry.YFilter
    cvvoicepeercfgentry.EntityData.YangName = "cvVoicePeerCfgEntry"
    cvvoicepeercfgentry.EntityData.BundleName = "cisco_ios_xe"
    cvvoicepeercfgentry.EntityData.ParentYangName = "cvVoicePeerCfgTable"
    cvvoicepeercfgentry.EntityData.SegmentPath = "cvVoicePeerCfgEntry" + "[ifIndex='" + fmt.Sprintf("%v", cvvoicepeercfgentry.Ifindex) + "']"
    cvvoicepeercfgentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvvoicepeercfgentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvvoicepeercfgentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvvoicepeercfgentry.EntityData.Children = make(map[string]types.YChild)
    cvvoicepeercfgentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cvvoicepeercfgentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", cvvoicepeercfgentry.Ifindex}
    cvvoicepeercfgentry.EntityData.Leafs["cvVoicePeerCfgSessionTarget"] = types.YLeaf{"Cvvoicepeercfgsessiontarget", cvvoicepeercfgentry.Cvvoicepeercfgsessiontarget}
    cvvoicepeercfgentry.EntityData.Leafs["cvVoicePeerCfgDialDigitsPrefix"] = types.YLeaf{"Cvvoicepeercfgdialdigitsprefix", cvvoicepeercfgentry.Cvvoicepeercfgdialdigitsprefix}
    cvvoicepeercfgentry.EntityData.Leafs["cvVoicePeerCfgDIDCallEnable"] = types.YLeaf{"Cvvoicepeercfgdidcallenable", cvvoicepeercfgentry.Cvvoicepeercfgdidcallenable}
    cvvoicepeercfgentry.EntityData.Leafs["cvVoicePeerCfgCasGroup"] = types.YLeaf{"Cvvoicepeercfgcasgroup", cvvoicepeercfgentry.Cvvoicepeercfgcasgroup}
    cvvoicepeercfgentry.EntityData.Leafs["cvVoicePeerCfgRegisterE164"] = types.YLeaf{"Cvvoicepeercfgregistere164", cvvoicepeercfgentry.Cvvoicepeercfgregistere164}
    cvvoicepeercfgentry.EntityData.Leafs["cvVoicePeerCfgForwardDigits"] = types.YLeaf{"Cvvoicepeercfgforwarddigits", cvvoicepeercfgentry.Cvvoicepeercfgforwarddigits}
    cvvoicepeercfgentry.EntityData.Leafs["cvVoicePeerCfgEchoCancellerTest"] = types.YLeaf{"Cvvoicepeercfgechocancellertest", cvvoicepeercfgentry.Cvvoicepeercfgechocancellertest}
    return &(cvvoicepeercfgentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A single VoIP specific Peer. One entry per VoIP encapsulation. The entry is
    // created when its associated 'voiceOverIp(104)' encapsulation ifEntry is
    // created. This entry is deleted when its associated ifEntry is deleted. The
    // type is slice of
    // CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable_Cvvoippeercfgentry.
    Cvvoippeercfgentry []CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable_Cvvoippeercfgentry
}

func (cvvoippeercfgtable *CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable) GetEntityData() *types.CommonEntityData {
    cvvoippeercfgtable.EntityData.YFilter = cvvoippeercfgtable.YFilter
    cvvoippeercfgtable.EntityData.YangName = "cvVoIPPeerCfgTable"
    cvvoippeercfgtable.EntityData.BundleName = "cisco_ios_xe"
    cvvoippeercfgtable.EntityData.ParentYangName = "CISCO-VOICE-DIAL-CONTROL-MIB"
    cvvoippeercfgtable.EntityData.SegmentPath = "cvVoIPPeerCfgTable"
    cvvoippeercfgtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvvoippeercfgtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvvoippeercfgtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvvoippeercfgtable.EntityData.Children = make(map[string]types.YChild)
    cvvoippeercfgtable.EntityData.Children["cvVoIPPeerCfgEntry"] = types.YChild{"Cvvoippeercfgentry", nil}
    for i := range cvvoippeercfgtable.Cvvoippeercfgentry {
        cvvoippeercfgtable.EntityData.Children[types.GetSegmentPath(&cvvoippeercfgtable.Cvvoippeercfgentry[i])] = types.YChild{"Cvvoippeercfgentry", &cvvoippeercfgtable.Cvvoippeercfgentry[i]}
    }
    cvvoippeercfgtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cvvoippeercfgtable.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable_Cvvoippeercfgentry
// A single VoIP specific Peer. One entry per VoIP
// encapsulation.
// The entry is created when its associated 'voiceOverIp(104)'
// encapsulation ifEntry is created.
// This entry is deleted when its associated ifEntry is
// deleted.
type CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable_Cvvoippeercfgentry struct {
    EntityData types.CommonEntityData
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

func (cvvoippeercfgentry *CISCOVOICEDIALCONTROLMIB_Cvvoippeercfgtable_Cvvoippeercfgentry) GetEntityData() *types.CommonEntityData {
    cvvoippeercfgentry.EntityData.YFilter = cvvoippeercfgentry.YFilter
    cvvoippeercfgentry.EntityData.YangName = "cvVoIPPeerCfgEntry"
    cvvoippeercfgentry.EntityData.BundleName = "cisco_ios_xe"
    cvvoippeercfgentry.EntityData.ParentYangName = "cvVoIPPeerCfgTable"
    cvvoippeercfgentry.EntityData.SegmentPath = "cvVoIPPeerCfgEntry" + "[ifIndex='" + fmt.Sprintf("%v", cvvoippeercfgentry.Ifindex) + "']"
    cvvoippeercfgentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvvoippeercfgentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvvoippeercfgentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvvoippeercfgentry.EntityData.Children = make(map[string]types.YChild)
    cvvoippeercfgentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cvvoippeercfgentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", cvvoippeercfgentry.Ifindex}
    cvvoippeercfgentry.EntityData.Leafs["cvVoIPPeerCfgSessionProtocol"] = types.YLeaf{"Cvvoippeercfgsessionprotocol", cvvoippeercfgentry.Cvvoippeercfgsessionprotocol}
    cvvoippeercfgentry.EntityData.Leafs["cvVoIPPeerCfgDesiredQoS"] = types.YLeaf{"Cvvoippeercfgdesiredqos", cvvoippeercfgentry.Cvvoippeercfgdesiredqos}
    cvvoippeercfgentry.EntityData.Leafs["cvVoIPPeerCfgMinAcceptableQoS"] = types.YLeaf{"Cvvoippeercfgminacceptableqos", cvvoippeercfgentry.Cvvoippeercfgminacceptableqos}
    cvvoippeercfgentry.EntityData.Leafs["cvVoIPPeerCfgSessionTarget"] = types.YLeaf{"Cvvoippeercfgsessiontarget", cvvoippeercfgentry.Cvvoippeercfgsessiontarget}
    cvvoippeercfgentry.EntityData.Leafs["cvVoIPPeerCfgCoderRate"] = types.YLeaf{"Cvvoippeercfgcoderrate", cvvoippeercfgentry.Cvvoippeercfgcoderrate}
    cvvoippeercfgentry.EntityData.Leafs["cvVoIPPeerCfgFaxRate"] = types.YLeaf{"Cvvoippeercfgfaxrate", cvvoippeercfgentry.Cvvoippeercfgfaxrate}
    cvvoippeercfgentry.EntityData.Leafs["cvVoIPPeerCfgVADEnable"] = types.YLeaf{"Cvvoippeercfgvadenable", cvvoippeercfgentry.Cvvoippeercfgvadenable}
    cvvoippeercfgentry.EntityData.Leafs["cvVoIPPeerCfgExpectFactor"] = types.YLeaf{"Cvvoippeercfgexpectfactor", cvvoippeercfgentry.Cvvoippeercfgexpectfactor}
    cvvoippeercfgentry.EntityData.Leafs["cvVoIPPeerCfgIcpif"] = types.YLeaf{"Cvvoippeercfgicpif", cvvoippeercfgentry.Cvvoippeercfgicpif}
    cvvoippeercfgentry.EntityData.Leafs["cvVoIPPeerCfgPoorQoVNotificationEnable"] = types.YLeaf{"Cvvoippeercfgpoorqovnotificationenable", cvvoippeercfgentry.Cvvoippeercfgpoorqovnotificationenable}
    cvvoippeercfgentry.EntityData.Leafs["cvVoIPPeerCfgUDPChecksumEnable"] = types.YLeaf{"Cvvoippeercfgudpchecksumenable", cvvoippeercfgentry.Cvvoippeercfgudpchecksumenable}
    cvvoippeercfgentry.EntityData.Leafs["cvVoIPPeerCfgIPPrecedence"] = types.YLeaf{"Cvvoippeercfgipprecedence", cvvoippeercfgentry.Cvvoippeercfgipprecedence}
    cvvoippeercfgentry.EntityData.Leafs["cvVoIPPeerCfgTechPrefix"] = types.YLeaf{"Cvvoippeercfgtechprefix", cvvoippeercfgentry.Cvvoippeercfgtechprefix}
    cvvoippeercfgentry.EntityData.Leafs["cvVoIPPeerCfgDigitRelay"] = types.YLeaf{"Cvvoippeercfgdigitrelay", cvvoippeercfgentry.Cvvoippeercfgdigitrelay}
    cvvoippeercfgentry.EntityData.Leafs["cvVoIPPeerCfgCoderBytes"] = types.YLeaf{"Cvvoippeercfgcoderbytes", cvvoippeercfgentry.Cvvoippeercfgcoderbytes}
    cvvoippeercfgentry.EntityData.Leafs["cvVoIPPeerCfgFaxBytes"] = types.YLeaf{"Cvvoippeercfgfaxbytes", cvvoippeercfgentry.Cvvoippeercfgfaxbytes}
    cvvoippeercfgentry.EntityData.Leafs["cvVoIPPeerCfgInBandSignaling"] = types.YLeaf{"Cvvoippeercfginbandsignaling", cvvoippeercfgentry.Cvvoippeercfginbandsignaling}
    cvvoippeercfgentry.EntityData.Leafs["cvVoIPPeerCfgMediaSetting"] = types.YLeaf{"Cvvoippeercfgmediasetting", cvvoippeercfgentry.Cvvoippeercfgmediasetting}
    cvvoippeercfgentry.EntityData.Leafs["cvVoIPPeerCfgDesiredQoSVideo"] = types.YLeaf{"Cvvoippeercfgdesiredqosvideo", cvvoippeercfgentry.Cvvoippeercfgdesiredqosvideo}
    cvvoippeercfgentry.EntityData.Leafs["cvVoIPPeerCfgMinAcceptableQoSVideo"] = types.YLeaf{"Cvvoippeercfgminacceptableqosvideo", cvvoippeercfgentry.Cvvoippeercfgminacceptableqosvideo}
    cvvoippeercfgentry.EntityData.Leafs["cvVoIPPeerCfgRedirectip2ip"] = types.YLeaf{"Cvvoippeercfgredirectip2Ip", cvvoippeercfgentry.Cvvoippeercfgredirectip2Ip}
    cvvoippeercfgentry.EntityData.Leafs["cvVoIPPeerCfgOctetAligned"] = types.YLeaf{"Cvvoippeercfgoctetaligned", cvvoippeercfgentry.Cvvoippeercfgoctetaligned}
    cvvoippeercfgentry.EntityData.Leafs["cvVoIPPeerCfgBitRates"] = types.YLeaf{"Cvvoippeercfgbitrates", cvvoippeercfgentry.Cvvoippeercfgbitrates}
    cvvoippeercfgentry.EntityData.Leafs["cvVoIPPeerCfgCRC"] = types.YLeaf{"Cvvoippeercfgcrc", cvvoippeercfgentry.Cvvoippeercfgcrc}
    cvvoippeercfgentry.EntityData.Leafs["cvVoIPPeerCfgCoderMode"] = types.YLeaf{"Cvvoippeercfgcodermode", cvvoippeercfgentry.Cvvoippeercfgcodermode}
    cvvoippeercfgentry.EntityData.Leafs["cvVoIPPeerCfgCodingMode"] = types.YLeaf{"Cvvoippeercfgcodingmode", cvvoippeercfgentry.Cvvoippeercfgcodingmode}
    cvvoippeercfgentry.EntityData.Leafs["cvVoIPPeerCfgBitRate"] = types.YLeaf{"Cvvoippeercfgbitrate", cvvoippeercfgentry.Cvvoippeercfgbitrate}
    cvvoippeercfgentry.EntityData.Leafs["cvVoIPPeerCfgFrameSize"] = types.YLeaf{"Cvvoippeercfgframesize", cvvoippeercfgentry.Cvvoippeercfgframesize}
    cvvoippeercfgentry.EntityData.Leafs["cvVoIPPeerCfgDSCPPolicyNotificationEnable"] = types.YLeaf{"Cvvoippeercfgdscppolicynotificationenable", cvvoippeercfgentry.Cvvoippeercfgdscppolicynotificationenable}
    cvvoippeercfgentry.EntityData.Leafs["cvVoIPPeerCfgMediaPolicyNotificationEnable"] = types.YLeaf{"Cvvoippeercfgmediapolicynotificationenable", cvvoippeercfgentry.Cvvoippeercfgmediapolicynotificationenable}
    return &(cvvoippeercfgentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A single Voice specific Peer. One entry per voice related encapsulation.
    // The entry is created when a voice related encapsulation ifEntry is created.
    // This entry is deleted when its associated ifEntry is deleted. The type is
    // slice of
    // CISCOVOICEDIALCONTROLMIB_Cvpeercommoncfgtable_Cvpeercommoncfgentry.
    Cvpeercommoncfgentry []CISCOVOICEDIALCONTROLMIB_Cvpeercommoncfgtable_Cvpeercommoncfgentry
}

func (cvpeercommoncfgtable *CISCOVOICEDIALCONTROLMIB_Cvpeercommoncfgtable) GetEntityData() *types.CommonEntityData {
    cvpeercommoncfgtable.EntityData.YFilter = cvpeercommoncfgtable.YFilter
    cvpeercommoncfgtable.EntityData.YangName = "cvPeerCommonCfgTable"
    cvpeercommoncfgtable.EntityData.BundleName = "cisco_ios_xe"
    cvpeercommoncfgtable.EntityData.ParentYangName = "CISCO-VOICE-DIAL-CONTROL-MIB"
    cvpeercommoncfgtable.EntityData.SegmentPath = "cvPeerCommonCfgTable"
    cvpeercommoncfgtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvpeercommoncfgtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvpeercommoncfgtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvpeercommoncfgtable.EntityData.Children = make(map[string]types.YChild)
    cvpeercommoncfgtable.EntityData.Children["cvPeerCommonCfgEntry"] = types.YChild{"Cvpeercommoncfgentry", nil}
    for i := range cvpeercommoncfgtable.Cvpeercommoncfgentry {
        cvpeercommoncfgtable.EntityData.Children[types.GetSegmentPath(&cvpeercommoncfgtable.Cvpeercommoncfgentry[i])] = types.YChild{"Cvpeercommoncfgentry", &cvpeercommoncfgtable.Cvpeercommoncfgentry[i]}
    }
    cvpeercommoncfgtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cvpeercommoncfgtable.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_Cvpeercommoncfgtable_Cvpeercommoncfgentry
// A single Voice specific Peer. One entry per voice related
// encapsulation.
// The entry is created when a voice related encapsulation
// ifEntry is created.
// This entry is deleted when its associated ifEntry is
// deleted.
type CISCOVOICEDIALCONTROLMIB_Cvpeercommoncfgtable_Cvpeercommoncfgentry struct {
    EntityData types.CommonEntityData
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

func (cvpeercommoncfgentry *CISCOVOICEDIALCONTROLMIB_Cvpeercommoncfgtable_Cvpeercommoncfgentry) GetEntityData() *types.CommonEntityData {
    cvpeercommoncfgentry.EntityData.YFilter = cvpeercommoncfgentry.YFilter
    cvpeercommoncfgentry.EntityData.YangName = "cvPeerCommonCfgEntry"
    cvpeercommoncfgentry.EntityData.BundleName = "cisco_ios_xe"
    cvpeercommoncfgentry.EntityData.ParentYangName = "cvPeerCommonCfgTable"
    cvpeercommoncfgentry.EntityData.SegmentPath = "cvPeerCommonCfgEntry" + "[ifIndex='" + fmt.Sprintf("%v", cvpeercommoncfgentry.Ifindex) + "']"
    cvpeercommoncfgentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvpeercommoncfgentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvpeercommoncfgentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvpeercommoncfgentry.EntityData.Children = make(map[string]types.YChild)
    cvpeercommoncfgentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cvpeercommoncfgentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", cvpeercommoncfgentry.Ifindex}
    cvpeercommoncfgentry.EntityData.Leafs["cvPeerCommonCfgIncomingDnisDigits"] = types.YLeaf{"Cvpeercommoncfgincomingdnisdigits", cvpeercommoncfgentry.Cvpeercommoncfgincomingdnisdigits}
    cvpeercommoncfgentry.EntityData.Leafs["cvPeerCommonCfgMaxConnections"] = types.YLeaf{"Cvpeercommoncfgmaxconnections", cvpeercommoncfgentry.Cvpeercommoncfgmaxconnections}
    cvpeercommoncfgentry.EntityData.Leafs["cvPeerCommonCfgApplicationName"] = types.YLeaf{"Cvpeercommoncfgapplicationname", cvpeercommoncfgentry.Cvpeercommoncfgapplicationname}
    cvpeercommoncfgentry.EntityData.Leafs["cvPeerCommonCfgPreference"] = types.YLeaf{"Cvpeercommoncfgpreference", cvpeercommoncfgentry.Cvpeercommoncfgpreference}
    cvpeercommoncfgentry.EntityData.Leafs["cvPeerCommonCfgHuntStop"] = types.YLeaf{"Cvpeercommoncfghuntstop", cvpeercommoncfgentry.Cvpeercommoncfghuntstop}
    cvpeercommoncfgentry.EntityData.Leafs["cvPeerCommonCfgDnisMappingName"] = types.YLeaf{"Cvpeercommoncfgdnismappingname", cvpeercommoncfgentry.Cvpeercommoncfgdnismappingname}
    cvpeercommoncfgentry.EntityData.Leafs["cvPeerCommonCfgSourceCarrierId"] = types.YLeaf{"Cvpeercommoncfgsourcecarrierid", cvpeercommoncfgentry.Cvpeercommoncfgsourcecarrierid}
    cvpeercommoncfgentry.EntityData.Leafs["cvPeerCommonCfgTargetCarrierId"] = types.YLeaf{"Cvpeercommoncfgtargetcarrierid", cvpeercommoncfgentry.Cvpeercommoncfgtargetcarrierid}
    cvpeercommoncfgentry.EntityData.Leafs["cvPeerCommonCfgSourceTrunkGrpLabel"] = types.YLeaf{"Cvpeercommoncfgsourcetrunkgrplabel", cvpeercommoncfgentry.Cvpeercommoncfgsourcetrunkgrplabel}
    cvpeercommoncfgentry.EntityData.Leafs["cvPeerCommonCfgTargetTrunkGrpLabel"] = types.YLeaf{"Cvpeercommoncfgtargettrunkgrplabel", cvpeercommoncfgentry.Cvpeercommoncfgtargettrunkgrplabel}
    return &(cvpeercommoncfgentry.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_Cvcallactivetable
// This table is the voice extension to the call active table
// of IETF Dial Control MIB. It contains voice encapsulation
// call leg information that is derived from the statistics
// of lower layer telephony interface.
type CISCOVOICEDIALCONTROLMIB_Cvcallactivetable struct {
    EntityData types.CommonEntityData
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

func (cvcallactivetable *CISCOVOICEDIALCONTROLMIB_Cvcallactivetable) GetEntityData() *types.CommonEntityData {
    cvcallactivetable.EntityData.YFilter = cvcallactivetable.YFilter
    cvcallactivetable.EntityData.YangName = "cvCallActiveTable"
    cvcallactivetable.EntityData.BundleName = "cisco_ios_xe"
    cvcallactivetable.EntityData.ParentYangName = "CISCO-VOICE-DIAL-CONTROL-MIB"
    cvcallactivetable.EntityData.SegmentPath = "cvCallActiveTable"
    cvcallactivetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvcallactivetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvcallactivetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvcallactivetable.EntityData.Children = make(map[string]types.YChild)
    cvcallactivetable.EntityData.Children["cvCallActiveEntry"] = types.YChild{"Cvcallactiveentry", nil}
    for i := range cvcallactivetable.Cvcallactiveentry {
        cvcallactivetable.EntityData.Children[types.GetSegmentPath(&cvcallactivetable.Cvcallactiveentry[i])] = types.YChild{"Cvcallactiveentry", &cvcallactivetable.Cvcallactiveentry[i]}
    }
    cvcallactivetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cvcallactivetable.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (cvcallactiveentry *CISCOVOICEDIALCONTROLMIB_Cvcallactivetable_Cvcallactiveentry) GetEntityData() *types.CommonEntityData {
    cvcallactiveentry.EntityData.YFilter = cvcallactiveentry.YFilter
    cvcallactiveentry.EntityData.YangName = "cvCallActiveEntry"
    cvcallactiveentry.EntityData.BundleName = "cisco_ios_xe"
    cvcallactiveentry.EntityData.ParentYangName = "cvCallActiveTable"
    cvcallactiveentry.EntityData.SegmentPath = "cvCallActiveEntry" + "[callActiveSetupTime='" + fmt.Sprintf("%v", cvcallactiveentry.Callactivesetuptime) + "']" + "[callActiveIndex='" + fmt.Sprintf("%v", cvcallactiveentry.Callactiveindex) + "']"
    cvcallactiveentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvcallactiveentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvcallactiveentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvcallactiveentry.EntityData.Children = make(map[string]types.YChild)
    cvcallactiveentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cvcallactiveentry.EntityData.Leafs["callActiveSetupTime"] = types.YLeaf{"Callactivesetuptime", cvcallactiveentry.Callactivesetuptime}
    cvcallactiveentry.EntityData.Leafs["callActiveIndex"] = types.YLeaf{"Callactiveindex", cvcallactiveentry.Callactiveindex}
    cvcallactiveentry.EntityData.Leafs["cvCallActiveConnectionId"] = types.YLeaf{"Cvcallactiveconnectionid", cvcallactiveentry.Cvcallactiveconnectionid}
    cvcallactiveentry.EntityData.Leafs["cvCallActiveTxDuration"] = types.YLeaf{"Cvcallactivetxduration", cvcallactiveentry.Cvcallactivetxduration}
    cvcallactiveentry.EntityData.Leafs["cvCallActiveVoiceTxDuration"] = types.YLeaf{"Cvcallactivevoicetxduration", cvcallactiveentry.Cvcallactivevoicetxduration}
    cvcallactiveentry.EntityData.Leafs["cvCallActiveFaxTxDuration"] = types.YLeaf{"Cvcallactivefaxtxduration", cvcallactiveentry.Cvcallactivefaxtxduration}
    cvcallactiveentry.EntityData.Leafs["cvCallActiveCoderTypeRate"] = types.YLeaf{"Cvcallactivecodertyperate", cvcallactiveentry.Cvcallactivecodertyperate}
    cvcallactiveentry.EntityData.Leafs["cvCallActiveNoiseLevel"] = types.YLeaf{"Cvcallactivenoiselevel", cvcallactiveentry.Cvcallactivenoiselevel}
    cvcallactiveentry.EntityData.Leafs["cvCallActiveACOMLevel"] = types.YLeaf{"Cvcallactiveacomlevel", cvcallactiveentry.Cvcallactiveacomlevel}
    cvcallactiveentry.EntityData.Leafs["cvCallActiveOutSignalLevel"] = types.YLeaf{"Cvcallactiveoutsignallevel", cvcallactiveentry.Cvcallactiveoutsignallevel}
    cvcallactiveentry.EntityData.Leafs["cvCallActiveInSignalLevel"] = types.YLeaf{"Cvcallactiveinsignallevel", cvcallactiveentry.Cvcallactiveinsignallevel}
    cvcallactiveentry.EntityData.Leafs["cvCallActiveERLLevel"] = types.YLeaf{"Cvcallactiveerllevel", cvcallactiveentry.Cvcallactiveerllevel}
    cvcallactiveentry.EntityData.Leafs["cvCallActiveSessionTarget"] = types.YLeaf{"Cvcallactivesessiontarget", cvcallactiveentry.Cvcallactivesessiontarget}
    cvcallactiveentry.EntityData.Leafs["cvCallActiveImgPageCount"] = types.YLeaf{"Cvcallactiveimgpagecount", cvcallactiveentry.Cvcallactiveimgpagecount}
    cvcallactiveentry.EntityData.Leafs["cvCallActiveCallingName"] = types.YLeaf{"Cvcallactivecallingname", cvcallactiveentry.Cvcallactivecallingname}
    cvcallactiveentry.EntityData.Leafs["cvCallActiveCallerIDBlock"] = types.YLeaf{"Cvcallactivecalleridblock", cvcallactiveentry.Cvcallactivecalleridblock}
    cvcallactiveentry.EntityData.Leafs["cvCallActiveEcanReflectorLocation"] = types.YLeaf{"Cvcallactiveecanreflectorlocation", cvcallactiveentry.Cvcallactiveecanreflectorlocation}
    cvcallactiveentry.EntityData.Leafs["cvCallActiveAccountCode"] = types.YLeaf{"Cvcallactiveaccountcode", cvcallactiveentry.Cvcallactiveaccountcode}
    cvcallactiveentry.EntityData.Leafs["cvCallActiveERLLevelRev1"] = types.YLeaf{"Cvcallactiveerllevelrev1", cvcallactiveentry.Cvcallactiveerllevelrev1}
    cvcallactiveentry.EntityData.Leafs["cvCallActiveCallId"] = types.YLeaf{"Cvcallactivecallid", cvcallactiveentry.Cvcallactivecallid}
    return &(cvcallactiveentry.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_Cvvoipcallactivetable
// This table is the VoIP extension to the call active table of
// IETF Dial Control MIB. It contains VoIP call leg
// information about specific VoIP call destination and the
// selected QoS for the call leg.
type CISCOVOICEDIALCONTROLMIB_Cvvoipcallactivetable struct {
    EntityData types.CommonEntityData
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

func (cvvoipcallactivetable *CISCOVOICEDIALCONTROLMIB_Cvvoipcallactivetable) GetEntityData() *types.CommonEntityData {
    cvvoipcallactivetable.EntityData.YFilter = cvvoipcallactivetable.YFilter
    cvvoipcallactivetable.EntityData.YangName = "cvVoIPCallActiveTable"
    cvvoipcallactivetable.EntityData.BundleName = "cisco_ios_xe"
    cvvoipcallactivetable.EntityData.ParentYangName = "CISCO-VOICE-DIAL-CONTROL-MIB"
    cvvoipcallactivetable.EntityData.SegmentPath = "cvVoIPCallActiveTable"
    cvvoipcallactivetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvvoipcallactivetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvvoipcallactivetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvvoipcallactivetable.EntityData.Children = make(map[string]types.YChild)
    cvvoipcallactivetable.EntityData.Children["cvVoIPCallActiveEntry"] = types.YChild{"Cvvoipcallactiveentry", nil}
    for i := range cvvoipcallactivetable.Cvvoipcallactiveentry {
        cvvoipcallactivetable.EntityData.Children[types.GetSegmentPath(&cvvoipcallactivetable.Cvvoipcallactiveentry[i])] = types.YChild{"Cvvoipcallactiveentry", &cvvoipcallactivetable.Cvvoipcallactiveentry[i]}
    }
    cvvoipcallactivetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cvvoipcallactivetable.EntityData)
}

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
    EntityData types.CommonEntityData
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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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

func (cvvoipcallactiveentry *CISCOVOICEDIALCONTROLMIB_Cvvoipcallactivetable_Cvvoipcallactiveentry) GetEntityData() *types.CommonEntityData {
    cvvoipcallactiveentry.EntityData.YFilter = cvvoipcallactiveentry.YFilter
    cvvoipcallactiveentry.EntityData.YangName = "cvVoIPCallActiveEntry"
    cvvoipcallactiveentry.EntityData.BundleName = "cisco_ios_xe"
    cvvoipcallactiveentry.EntityData.ParentYangName = "cvVoIPCallActiveTable"
    cvvoipcallactiveentry.EntityData.SegmentPath = "cvVoIPCallActiveEntry" + "[callActiveSetupTime='" + fmt.Sprintf("%v", cvvoipcallactiveentry.Callactivesetuptime) + "']" + "[callActiveIndex='" + fmt.Sprintf("%v", cvvoipcallactiveentry.Callactiveindex) + "']"
    cvvoipcallactiveentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvvoipcallactiveentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvvoipcallactiveentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvvoipcallactiveentry.EntityData.Children = make(map[string]types.YChild)
    cvvoipcallactiveentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cvvoipcallactiveentry.EntityData.Leafs["callActiveSetupTime"] = types.YLeaf{"Callactivesetuptime", cvvoipcallactiveentry.Callactivesetuptime}
    cvvoipcallactiveentry.EntityData.Leafs["callActiveIndex"] = types.YLeaf{"Callactiveindex", cvvoipcallactiveentry.Callactiveindex}
    cvvoipcallactiveentry.EntityData.Leafs["cvVoIPCallActiveConnectionId"] = types.YLeaf{"Cvvoipcallactiveconnectionid", cvvoipcallactiveentry.Cvvoipcallactiveconnectionid}
    cvvoipcallactiveentry.EntityData.Leafs["cvVoIPCallActiveRemoteIPAddress"] = types.YLeaf{"Cvvoipcallactiveremoteipaddress", cvvoipcallactiveentry.Cvvoipcallactiveremoteipaddress}
    cvvoipcallactiveentry.EntityData.Leafs["cvVoIPCallActiveRemoteUDPPort"] = types.YLeaf{"Cvvoipcallactiveremoteudpport", cvvoipcallactiveentry.Cvvoipcallactiveremoteudpport}
    cvvoipcallactiveentry.EntityData.Leafs["cvVoIPCallActiveRoundTripDelay"] = types.YLeaf{"Cvvoipcallactiveroundtripdelay", cvvoipcallactiveentry.Cvvoipcallactiveroundtripdelay}
    cvvoipcallactiveentry.EntityData.Leafs["cvVoIPCallActiveSelectedQoS"] = types.YLeaf{"Cvvoipcallactiveselectedqos", cvvoipcallactiveentry.Cvvoipcallactiveselectedqos}
    cvvoipcallactiveentry.EntityData.Leafs["cvVoIPCallActiveSessionProtocol"] = types.YLeaf{"Cvvoipcallactivesessionprotocol", cvvoipcallactiveentry.Cvvoipcallactivesessionprotocol}
    cvvoipcallactiveentry.EntityData.Leafs["cvVoIPCallActiveSessionTarget"] = types.YLeaf{"Cvvoipcallactivesessiontarget", cvvoipcallactiveentry.Cvvoipcallactivesessiontarget}
    cvvoipcallactiveentry.EntityData.Leafs["cvVoIPCallActiveOnTimeRvPlayout"] = types.YLeaf{"Cvvoipcallactiveontimervplayout", cvvoipcallactiveentry.Cvvoipcallactiveontimervplayout}
    cvvoipcallactiveentry.EntityData.Leafs["cvVoIPCallActiveGapFillWithSilence"] = types.YLeaf{"Cvvoipcallactivegapfillwithsilence", cvvoipcallactiveentry.Cvvoipcallactivegapfillwithsilence}
    cvvoipcallactiveentry.EntityData.Leafs["cvVoIPCallActiveGapFillWithPrediction"] = types.YLeaf{"Cvvoipcallactivegapfillwithprediction", cvvoipcallactiveentry.Cvvoipcallactivegapfillwithprediction}
    cvvoipcallactiveentry.EntityData.Leafs["cvVoIPCallActiveGapFillWithInterpolation"] = types.YLeaf{"Cvvoipcallactivegapfillwithinterpolation", cvvoipcallactiveentry.Cvvoipcallactivegapfillwithinterpolation}
    cvvoipcallactiveentry.EntityData.Leafs["cvVoIPCallActiveGapFillWithRedundancy"] = types.YLeaf{"Cvvoipcallactivegapfillwithredundancy", cvvoipcallactiveentry.Cvvoipcallactivegapfillwithredundancy}
    cvvoipcallactiveentry.EntityData.Leafs["cvVoIPCallActiveHiWaterPlayoutDelay"] = types.YLeaf{"Cvvoipcallactivehiwaterplayoutdelay", cvvoipcallactiveentry.Cvvoipcallactivehiwaterplayoutdelay}
    cvvoipcallactiveentry.EntityData.Leafs["cvVoIPCallActiveLoWaterPlayoutDelay"] = types.YLeaf{"Cvvoipcallactivelowaterplayoutdelay", cvvoipcallactiveentry.Cvvoipcallactivelowaterplayoutdelay}
    cvvoipcallactiveentry.EntityData.Leafs["cvVoIPCallActiveReceiveDelay"] = types.YLeaf{"Cvvoipcallactivereceivedelay", cvvoipcallactiveentry.Cvvoipcallactivereceivedelay}
    cvvoipcallactiveentry.EntityData.Leafs["cvVoIPCallActiveVADEnable"] = types.YLeaf{"Cvvoipcallactivevadenable", cvvoipcallactiveentry.Cvvoipcallactivevadenable}
    cvvoipcallactiveentry.EntityData.Leafs["cvVoIPCallActiveCoderTypeRate"] = types.YLeaf{"Cvvoipcallactivecodertyperate", cvvoipcallactiveentry.Cvvoipcallactivecodertyperate}
    cvvoipcallactiveentry.EntityData.Leafs["cvVoIPCallActiveLostPackets"] = types.YLeaf{"Cvvoipcallactivelostpackets", cvvoipcallactiveentry.Cvvoipcallactivelostpackets}
    cvvoipcallactiveentry.EntityData.Leafs["cvVoIPCallActiveEarlyPackets"] = types.YLeaf{"Cvvoipcallactiveearlypackets", cvvoipcallactiveentry.Cvvoipcallactiveearlypackets}
    cvvoipcallactiveentry.EntityData.Leafs["cvVoIPCallActiveLatePackets"] = types.YLeaf{"Cvvoipcallactivelatepackets", cvvoipcallactiveentry.Cvvoipcallactivelatepackets}
    cvvoipcallactiveentry.EntityData.Leafs["cvVoIPCallActiveUsername"] = types.YLeaf{"Cvvoipcallactiveusername", cvvoipcallactiveentry.Cvvoipcallactiveusername}
    cvvoipcallactiveentry.EntityData.Leafs["cvVoIPCallActiveProtocolCallId"] = types.YLeaf{"Cvvoipcallactiveprotocolcallid", cvvoipcallactiveentry.Cvvoipcallactiveprotocolcallid}
    cvvoipcallactiveentry.EntityData.Leafs["cvVoIPCallActiveRemSigIPAddrT"] = types.YLeaf{"Cvvoipcallactiveremsigipaddrt", cvvoipcallactiveentry.Cvvoipcallactiveremsigipaddrt}
    cvvoipcallactiveentry.EntityData.Leafs["cvVoIPCallActiveRemSigIPAddr"] = types.YLeaf{"Cvvoipcallactiveremsigipaddr", cvvoipcallactiveentry.Cvvoipcallactiveremsigipaddr}
    cvvoipcallactiveentry.EntityData.Leafs["cvVoIPCallActiveRemSigPort"] = types.YLeaf{"Cvvoipcallactiveremsigport", cvvoipcallactiveentry.Cvvoipcallactiveremsigport}
    cvvoipcallactiveentry.EntityData.Leafs["cvVoIPCallActiveRemMediaIPAddrT"] = types.YLeaf{"Cvvoipcallactiveremmediaipaddrt", cvvoipcallactiveentry.Cvvoipcallactiveremmediaipaddrt}
    cvvoipcallactiveentry.EntityData.Leafs["cvVoIPCallActiveRemMediaIPAddr"] = types.YLeaf{"Cvvoipcallactiveremmediaipaddr", cvvoipcallactiveentry.Cvvoipcallactiveremmediaipaddr}
    cvvoipcallactiveentry.EntityData.Leafs["cvVoIPCallActiveRemMediaPort"] = types.YLeaf{"Cvvoipcallactiveremmediaport", cvvoipcallactiveentry.Cvvoipcallactiveremmediaport}
    cvvoipcallactiveentry.EntityData.Leafs["cvVoIPCallActiveSRTPEnable"] = types.YLeaf{"Cvvoipcallactivesrtpenable", cvvoipcallactiveentry.Cvvoipcallactivesrtpenable}
    cvvoipcallactiveentry.EntityData.Leafs["cvVoIPCallActiveOctetAligned"] = types.YLeaf{"Cvvoipcallactiveoctetaligned", cvvoipcallactiveentry.Cvvoipcallactiveoctetaligned}
    cvvoipcallactiveentry.EntityData.Leafs["cvVoIPCallActiveBitRates"] = types.YLeaf{"Cvvoipcallactivebitrates", cvvoipcallactiveentry.Cvvoipcallactivebitrates}
    cvvoipcallactiveentry.EntityData.Leafs["cvVoIPCallActiveModeChgPeriod"] = types.YLeaf{"Cvvoipcallactivemodechgperiod", cvvoipcallactiveentry.Cvvoipcallactivemodechgperiod}
    cvvoipcallactiveentry.EntityData.Leafs["cvVoIPCallActiveModeChgNeighbor"] = types.YLeaf{"Cvvoipcallactivemodechgneighbor", cvvoipcallactiveentry.Cvvoipcallactivemodechgneighbor}
    cvvoipcallactiveentry.EntityData.Leafs["cvVoIPCallActiveMaxPtime"] = types.YLeaf{"Cvvoipcallactivemaxptime", cvvoipcallactiveentry.Cvvoipcallactivemaxptime}
    cvvoipcallactiveentry.EntityData.Leafs["cvVoIPCallActiveCRC"] = types.YLeaf{"Cvvoipcallactivecrc", cvvoipcallactiveentry.Cvvoipcallactivecrc}
    cvvoipcallactiveentry.EntityData.Leafs["cvVoIPCallActiveRobustSorting"] = types.YLeaf{"Cvvoipcallactiverobustsorting", cvvoipcallactiveentry.Cvvoipcallactiverobustsorting}
    cvvoipcallactiveentry.EntityData.Leafs["cvVoIPCallActiveEncap"] = types.YLeaf{"Cvvoipcallactiveencap", cvvoipcallactiveentry.Cvvoipcallactiveencap}
    cvvoipcallactiveentry.EntityData.Leafs["cvVoIPCallActiveInterleaving"] = types.YLeaf{"Cvvoipcallactiveinterleaving", cvvoipcallactiveentry.Cvvoipcallactiveinterleaving}
    cvvoipcallactiveentry.EntityData.Leafs["cvVoIPCallActivePtime"] = types.YLeaf{"Cvvoipcallactiveptime", cvvoipcallactiveentry.Cvvoipcallactiveptime}
    cvvoipcallactiveentry.EntityData.Leafs["cvVoIPCallActiveChannels"] = types.YLeaf{"Cvvoipcallactivechannels", cvvoipcallactiveentry.Cvvoipcallactivechannels}
    cvvoipcallactiveentry.EntityData.Leafs["cvVoIPCallActiveCoderMode"] = types.YLeaf{"Cvvoipcallactivecodermode", cvvoipcallactiveentry.Cvvoipcallactivecodermode}
    cvvoipcallactiveentry.EntityData.Leafs["cvVoIPCallActiveCallId"] = types.YLeaf{"Cvvoipcallactivecallid", cvvoipcallactiveentry.Cvvoipcallactivecallid}
    cvvoipcallactiveentry.EntityData.Leafs["cvVoIPCallActiveCallReferenceId"] = types.YLeaf{"Cvvoipcallactivecallreferenceid", cvvoipcallactiveentry.Cvvoipcallactivecallreferenceid}
    cvvoipcallactiveentry.EntityData.Leafs["ccVoIPCallActivePolicyName"] = types.YLeaf{"Ccvoipcallactivepolicyname", cvvoipcallactiveentry.Ccvoipcallactivepolicyname}
    cvvoipcallactiveentry.EntityData.Leafs["cvVoIPCallActiveReversedDirectionPeerAddress"] = types.YLeaf{"Cvvoipcallactivereverseddirectionpeeraddress", cvvoipcallactiveentry.Cvvoipcallactivereverseddirectionpeeraddress}
    cvvoipcallactiveentry.EntityData.Leafs["cvVoIPCallActiveSessionId"] = types.YLeaf{"Cvvoipcallactivesessionid", cvvoipcallactiveentry.Cvvoipcallactivesessionid}
    return &(cvvoipcallactiveentry.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_Cvcallvolconntable
// This table represents the number of active
// call connections for each call connection type
// in the voice gateway.
type CISCOVOICEDIALCONTROLMIB_Cvcallvolconntable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the cvCallVolConnTable indicates number of active calls for a
    // call connection type in the voice gateway. The type is slice of
    // CISCOVOICEDIALCONTROLMIB_Cvcallvolconntable_Cvcallvolconnentry.
    Cvcallvolconnentry []CISCOVOICEDIALCONTROLMIB_Cvcallvolconntable_Cvcallvolconnentry
}

func (cvcallvolconntable *CISCOVOICEDIALCONTROLMIB_Cvcallvolconntable) GetEntityData() *types.CommonEntityData {
    cvcallvolconntable.EntityData.YFilter = cvcallvolconntable.YFilter
    cvcallvolconntable.EntityData.YangName = "cvCallVolConnTable"
    cvcallvolconntable.EntityData.BundleName = "cisco_ios_xe"
    cvcallvolconntable.EntityData.ParentYangName = "CISCO-VOICE-DIAL-CONTROL-MIB"
    cvcallvolconntable.EntityData.SegmentPath = "cvCallVolConnTable"
    cvcallvolconntable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvcallvolconntable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvcallvolconntable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvcallvolconntable.EntityData.Children = make(map[string]types.YChild)
    cvcallvolconntable.EntityData.Children["cvCallVolConnEntry"] = types.YChild{"Cvcallvolconnentry", nil}
    for i := range cvcallvolconntable.Cvcallvolconnentry {
        cvcallvolconntable.EntityData.Children[types.GetSegmentPath(&cvcallvolconntable.Cvcallvolconnentry[i])] = types.YChild{"Cvcallvolconnentry", &cvcallvolconntable.Cvcallvolconnentry[i]}
    }
    cvcallvolconntable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cvcallvolconntable.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_Cvcallvolconntable_Cvcallvolconnentry
// An entry in the cvCallVolConnTable indicates
// number of active calls for a call connection type
// in the voice gateway.
type CISCOVOICEDIALCONTROLMIB_Cvcallvolconntable_Cvcallvolconnentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This object represents index to the
    // cvCallVolConnTable. The type is CvCallConnectionType.
    Cvcallvolconnindex interface{}

    // This object represents the total number of active calls for a connection
    // type  in the voice gateway. The type is interface{} with range: 0..65535.
    Cvcallvolconnactiveconnection interface{}
}

func (cvcallvolconnentry *CISCOVOICEDIALCONTROLMIB_Cvcallvolconntable_Cvcallvolconnentry) GetEntityData() *types.CommonEntityData {
    cvcallvolconnentry.EntityData.YFilter = cvcallvolconnentry.YFilter
    cvcallvolconnentry.EntityData.YangName = "cvCallVolConnEntry"
    cvcallvolconnentry.EntityData.BundleName = "cisco_ios_xe"
    cvcallvolconnentry.EntityData.ParentYangName = "cvCallVolConnTable"
    cvcallvolconnentry.EntityData.SegmentPath = "cvCallVolConnEntry" + "[cvCallVolConnIndex='" + fmt.Sprintf("%v", cvcallvolconnentry.Cvcallvolconnindex) + "']"
    cvcallvolconnentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvcallvolconnentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvcallvolconnentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvcallvolconnentry.EntityData.Children = make(map[string]types.YChild)
    cvcallvolconnentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cvcallvolconnentry.EntityData.Leafs["cvCallVolConnIndex"] = types.YLeaf{"Cvcallvolconnindex", cvcallvolconnentry.Cvcallvolconnindex}
    cvcallvolconnentry.EntityData.Leafs["cvCallVolConnActiveConnection"] = types.YLeaf{"Cvcallvolconnactiveconnection", cvcallvolconnentry.Cvcallvolconnactiveconnection}
    return &(cvcallvolconnentry.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_Cvcallvoliftable
// This table represents the information about
// the usage of an IP interface in a voice
// gateway for voice media calls. This table
// has a sparse-dependent relationship with  
// ifTable. There exists an entry in this table, 
// for each of the  entries in ifTable where ifType 
// is one of 'ethernetCsmacd' and 'softwareLoopback'.
type CISCOVOICEDIALCONTROLMIB_Cvcallvoliftable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry represents a row in cvCallVolIfTable and corresponds to the
    // information about an IP  interface in the voice gateway. The type is slice
    // of CISCOVOICEDIALCONTROLMIB_Cvcallvoliftable_Cvcallvolifentry.
    Cvcallvolifentry []CISCOVOICEDIALCONTROLMIB_Cvcallvoliftable_Cvcallvolifentry
}

func (cvcallvoliftable *CISCOVOICEDIALCONTROLMIB_Cvcallvoliftable) GetEntityData() *types.CommonEntityData {
    cvcallvoliftable.EntityData.YFilter = cvcallvoliftable.YFilter
    cvcallvoliftable.EntityData.YangName = "cvCallVolIfTable"
    cvcallvoliftable.EntityData.BundleName = "cisco_ios_xe"
    cvcallvoliftable.EntityData.ParentYangName = "CISCO-VOICE-DIAL-CONTROL-MIB"
    cvcallvoliftable.EntityData.SegmentPath = "cvCallVolIfTable"
    cvcallvoliftable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvcallvoliftable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvcallvoliftable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvcallvoliftable.EntityData.Children = make(map[string]types.YChild)
    cvcallvoliftable.EntityData.Children["cvCallVolIfEntry"] = types.YChild{"Cvcallvolifentry", nil}
    for i := range cvcallvoliftable.Cvcallvolifentry {
        cvcallvoliftable.EntityData.Children[types.GetSegmentPath(&cvcallvoliftable.Cvcallvolifentry[i])] = types.YChild{"Cvcallvolifentry", &cvcallvoliftable.Cvcallvolifentry[i]}
    }
    cvcallvoliftable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cvcallvoliftable.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_Cvcallvoliftable_Cvcallvolifentry
// Each entry represents a row in cvCallVolIfTable
// and corresponds to the information about an IP 
// interface in the voice gateway.
type CISCOVOICEDIALCONTROLMIB_Cvcallvoliftable_Cvcallvolifentry struct {
    EntityData types.CommonEntityData
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

func (cvcallvolifentry *CISCOVOICEDIALCONTROLMIB_Cvcallvoliftable_Cvcallvolifentry) GetEntityData() *types.CommonEntityData {
    cvcallvolifentry.EntityData.YFilter = cvcallvolifentry.YFilter
    cvcallvolifentry.EntityData.YangName = "cvCallVolIfEntry"
    cvcallvolifentry.EntityData.BundleName = "cisco_ios_xe"
    cvcallvolifentry.EntityData.ParentYangName = "cvCallVolIfTable"
    cvcallvolifentry.EntityData.SegmentPath = "cvCallVolIfEntry" + "[ifIndex='" + fmt.Sprintf("%v", cvcallvolifentry.Ifindex) + "']"
    cvcallvolifentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvcallvolifentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvcallvolifentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvcallvolifentry.EntityData.Children = make(map[string]types.YChild)
    cvcallvolifentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cvcallvolifentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", cvcallvolifentry.Ifindex}
    cvcallvolifentry.EntityData.Leafs["cvCallVolMediaIncomingCalls"] = types.YLeaf{"Cvcallvolmediaincomingcalls", cvcallvolifentry.Cvcallvolmediaincomingcalls}
    cvcallvolifentry.EntityData.Leafs["cvCallVolMediaOutgoingCalls"] = types.YLeaf{"Cvcallvolmediaoutgoingcalls", cvcallvolifentry.Cvcallvolmediaoutgoingcalls}
    return &(cvcallvolifentry.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_Cvcallhistorytable
// This table is the voice extension to the call history table
// of IETF Dial Control MIB. It contains voice encapsulation
// call leg information such as voice packet statistics,
// coder usage and end to end bandwidth of the call leg.
type CISCOVOICEDIALCONTROLMIB_Cvcallhistorytable struct {
    EntityData types.CommonEntityData
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

func (cvcallhistorytable *CISCOVOICEDIALCONTROLMIB_Cvcallhistorytable) GetEntityData() *types.CommonEntityData {
    cvcallhistorytable.EntityData.YFilter = cvcallhistorytable.YFilter
    cvcallhistorytable.EntityData.YangName = "cvCallHistoryTable"
    cvcallhistorytable.EntityData.BundleName = "cisco_ios_xe"
    cvcallhistorytable.EntityData.ParentYangName = "CISCO-VOICE-DIAL-CONTROL-MIB"
    cvcallhistorytable.EntityData.SegmentPath = "cvCallHistoryTable"
    cvcallhistorytable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvcallhistorytable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvcallhistorytable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvcallhistorytable.EntityData.Children = make(map[string]types.YChild)
    cvcallhistorytable.EntityData.Children["cvCallHistoryEntry"] = types.YChild{"Cvcallhistoryentry", nil}
    for i := range cvcallhistorytable.Cvcallhistoryentry {
        cvcallhistorytable.EntityData.Children[types.GetSegmentPath(&cvcallhistorytable.Cvcallhistoryentry[i])] = types.YChild{"Cvcallhistoryentry", &cvcallhistorytable.Cvcallhistoryentry[i]}
    }
    cvcallhistorytable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cvcallhistorytable.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (cvcallhistoryentry *CISCOVOICEDIALCONTROLMIB_Cvcallhistorytable_Cvcallhistoryentry) GetEntityData() *types.CommonEntityData {
    cvcallhistoryentry.EntityData.YFilter = cvcallhistoryentry.YFilter
    cvcallhistoryentry.EntityData.YangName = "cvCallHistoryEntry"
    cvcallhistoryentry.EntityData.BundleName = "cisco_ios_xe"
    cvcallhistoryentry.EntityData.ParentYangName = "cvCallHistoryTable"
    cvcallhistoryentry.EntityData.SegmentPath = "cvCallHistoryEntry" + "[cCallHistoryIndex='" + fmt.Sprintf("%v", cvcallhistoryentry.Ccallhistoryindex) + "']"
    cvcallhistoryentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvcallhistoryentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvcallhistoryentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvcallhistoryentry.EntityData.Children = make(map[string]types.YChild)
    cvcallhistoryentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cvcallhistoryentry.EntityData.Leafs["cCallHistoryIndex"] = types.YLeaf{"Ccallhistoryindex", cvcallhistoryentry.Ccallhistoryindex}
    cvcallhistoryentry.EntityData.Leafs["cvCallHistoryConnectionId"] = types.YLeaf{"Cvcallhistoryconnectionid", cvcallhistoryentry.Cvcallhistoryconnectionid}
    cvcallhistoryentry.EntityData.Leafs["cvCallHistoryTxDuration"] = types.YLeaf{"Cvcallhistorytxduration", cvcallhistoryentry.Cvcallhistorytxduration}
    cvcallhistoryentry.EntityData.Leafs["cvCallHistoryVoiceTxDuration"] = types.YLeaf{"Cvcallhistoryvoicetxduration", cvcallhistoryentry.Cvcallhistoryvoicetxduration}
    cvcallhistoryentry.EntityData.Leafs["cvCallHistoryFaxTxDuration"] = types.YLeaf{"Cvcallhistoryfaxtxduration", cvcallhistoryentry.Cvcallhistoryfaxtxduration}
    cvcallhistoryentry.EntityData.Leafs["cvCallHistoryCoderTypeRate"] = types.YLeaf{"Cvcallhistorycodertyperate", cvcallhistoryentry.Cvcallhistorycodertyperate}
    cvcallhistoryentry.EntityData.Leafs["cvCallHistoryNoiseLevel"] = types.YLeaf{"Cvcallhistorynoiselevel", cvcallhistoryentry.Cvcallhistorynoiselevel}
    cvcallhistoryentry.EntityData.Leafs["cvCallHistoryACOMLevel"] = types.YLeaf{"Cvcallhistoryacomlevel", cvcallhistoryentry.Cvcallhistoryacomlevel}
    cvcallhistoryentry.EntityData.Leafs["cvCallHistorySessionTarget"] = types.YLeaf{"Cvcallhistorysessiontarget", cvcallhistoryentry.Cvcallhistorysessiontarget}
    cvcallhistoryentry.EntityData.Leafs["cvCallHistoryImgPageCount"] = types.YLeaf{"Cvcallhistoryimgpagecount", cvcallhistoryentry.Cvcallhistoryimgpagecount}
    cvcallhistoryentry.EntityData.Leafs["cvCallHistoryCallingName"] = types.YLeaf{"Cvcallhistorycallingname", cvcallhistoryentry.Cvcallhistorycallingname}
    cvcallhistoryentry.EntityData.Leafs["cvCallHistoryCallerIDBlock"] = types.YLeaf{"Cvcallhistorycalleridblock", cvcallhistoryentry.Cvcallhistorycalleridblock}
    cvcallhistoryentry.EntityData.Leafs["cvCallHistoryAccountCode"] = types.YLeaf{"Cvcallhistoryaccountcode", cvcallhistoryentry.Cvcallhistoryaccountcode}
    cvcallhistoryentry.EntityData.Leafs["cvCallHistoryCallId"] = types.YLeaf{"Cvcallhistorycallid", cvcallhistoryentry.Cvcallhistorycallid}
    return &(cvcallhistoryentry.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_Cvvoipcallhistorytable
// This table is the VoIP extension to the call history table
// of IETF Dial Control MIB. It contains VoIP call leg
// information about specific VoIP call destination and the
// selected QoS for the call leg.
type CISCOVOICEDIALCONTROLMIB_Cvvoipcallhistorytable struct {
    EntityData types.CommonEntityData
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

func (cvvoipcallhistorytable *CISCOVOICEDIALCONTROLMIB_Cvvoipcallhistorytable) GetEntityData() *types.CommonEntityData {
    cvvoipcallhistorytable.EntityData.YFilter = cvvoipcallhistorytable.YFilter
    cvvoipcallhistorytable.EntityData.YangName = "cvVoIPCallHistoryTable"
    cvvoipcallhistorytable.EntityData.BundleName = "cisco_ios_xe"
    cvvoipcallhistorytable.EntityData.ParentYangName = "CISCO-VOICE-DIAL-CONTROL-MIB"
    cvvoipcallhistorytable.EntityData.SegmentPath = "cvVoIPCallHistoryTable"
    cvvoipcallhistorytable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvvoipcallhistorytable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvvoipcallhistorytable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvvoipcallhistorytable.EntityData.Children = make(map[string]types.YChild)
    cvvoipcallhistorytable.EntityData.Children["cvVoIPCallHistoryEntry"] = types.YChild{"Cvvoipcallhistoryentry", nil}
    for i := range cvvoipcallhistorytable.Cvvoipcallhistoryentry {
        cvvoipcallhistorytable.EntityData.Children[types.GetSegmentPath(&cvvoipcallhistorytable.Cvvoipcallhistoryentry[i])] = types.YChild{"Cvvoipcallhistoryentry", &cvvoipcallhistorytable.Cvvoipcallhistoryentry[i]}
    }
    cvvoipcallhistorytable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cvvoipcallhistorytable.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // cisco_dial_control_mib.CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryindex
    Ccallhistoryindex interface{}

    // The global connection identifier for the VoIP leg, which was assigned to
    // the call. The type is string with length: 0..16.
    Cvvoipcallhistoryconnectionid interface{}

    // Remote system IP address for the call. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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

func (cvvoipcallhistoryentry *CISCOVOICEDIALCONTROLMIB_Cvvoipcallhistorytable_Cvvoipcallhistoryentry) GetEntityData() *types.CommonEntityData {
    cvvoipcallhistoryentry.EntityData.YFilter = cvvoipcallhistoryentry.YFilter
    cvvoipcallhistoryentry.EntityData.YangName = "cvVoIPCallHistoryEntry"
    cvvoipcallhistoryentry.EntityData.BundleName = "cisco_ios_xe"
    cvvoipcallhistoryentry.EntityData.ParentYangName = "cvVoIPCallHistoryTable"
    cvvoipcallhistoryentry.EntityData.SegmentPath = "cvVoIPCallHistoryEntry" + "[cCallHistoryIndex='" + fmt.Sprintf("%v", cvvoipcallhistoryentry.Ccallhistoryindex) + "']"
    cvvoipcallhistoryentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvvoipcallhistoryentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvvoipcallhistoryentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvvoipcallhistoryentry.EntityData.Children = make(map[string]types.YChild)
    cvvoipcallhistoryentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cvvoipcallhistoryentry.EntityData.Leafs["cCallHistoryIndex"] = types.YLeaf{"Ccallhistoryindex", cvvoipcallhistoryentry.Ccallhistoryindex}
    cvvoipcallhistoryentry.EntityData.Leafs["cvVoIPCallHistoryConnectionId"] = types.YLeaf{"Cvvoipcallhistoryconnectionid", cvvoipcallhistoryentry.Cvvoipcallhistoryconnectionid}
    cvvoipcallhistoryentry.EntityData.Leafs["cvVoIPCallHistoryRemoteIPAddress"] = types.YLeaf{"Cvvoipcallhistoryremoteipaddress", cvvoipcallhistoryentry.Cvvoipcallhistoryremoteipaddress}
    cvvoipcallhistoryentry.EntityData.Leafs["cvVoIPCallHistoryRemoteUDPPort"] = types.YLeaf{"Cvvoipcallhistoryremoteudpport", cvvoipcallhistoryentry.Cvvoipcallhistoryremoteudpport}
    cvvoipcallhistoryentry.EntityData.Leafs["cvVoIPCallHistoryRoundTripDelay"] = types.YLeaf{"Cvvoipcallhistoryroundtripdelay", cvvoipcallhistoryentry.Cvvoipcallhistoryroundtripdelay}
    cvvoipcallhistoryentry.EntityData.Leafs["cvVoIPCallHistorySelectedQoS"] = types.YLeaf{"Cvvoipcallhistoryselectedqos", cvvoipcallhistoryentry.Cvvoipcallhistoryselectedqos}
    cvvoipcallhistoryentry.EntityData.Leafs["cvVoIPCallHistorySessionProtocol"] = types.YLeaf{"Cvvoipcallhistorysessionprotocol", cvvoipcallhistoryentry.Cvvoipcallhistorysessionprotocol}
    cvvoipcallhistoryentry.EntityData.Leafs["cvVoIPCallHistorySessionTarget"] = types.YLeaf{"Cvvoipcallhistorysessiontarget", cvvoipcallhistoryentry.Cvvoipcallhistorysessiontarget}
    cvvoipcallhistoryentry.EntityData.Leafs["cvVoIPCallHistoryOnTimeRvPlayout"] = types.YLeaf{"Cvvoipcallhistoryontimervplayout", cvvoipcallhistoryentry.Cvvoipcallhistoryontimervplayout}
    cvvoipcallhistoryentry.EntityData.Leafs["cvVoIPCallHistoryGapFillWithSilence"] = types.YLeaf{"Cvvoipcallhistorygapfillwithsilence", cvvoipcallhistoryentry.Cvvoipcallhistorygapfillwithsilence}
    cvvoipcallhistoryentry.EntityData.Leafs["cvVoIPCallHistoryGapFillWithPrediction"] = types.YLeaf{"Cvvoipcallhistorygapfillwithprediction", cvvoipcallhistoryentry.Cvvoipcallhistorygapfillwithprediction}
    cvvoipcallhistoryentry.EntityData.Leafs["cvVoIPCallHistoryGapFillWithInterpolation"] = types.YLeaf{"Cvvoipcallhistorygapfillwithinterpolation", cvvoipcallhistoryentry.Cvvoipcallhistorygapfillwithinterpolation}
    cvvoipcallhistoryentry.EntityData.Leafs["cvVoIPCallHistoryGapFillWithRedundancy"] = types.YLeaf{"Cvvoipcallhistorygapfillwithredundancy", cvvoipcallhistoryentry.Cvvoipcallhistorygapfillwithredundancy}
    cvvoipcallhistoryentry.EntityData.Leafs["cvVoIPCallHistoryHiWaterPlayoutDelay"] = types.YLeaf{"Cvvoipcallhistoryhiwaterplayoutdelay", cvvoipcallhistoryentry.Cvvoipcallhistoryhiwaterplayoutdelay}
    cvvoipcallhistoryentry.EntityData.Leafs["cvVoIPCallHistoryLoWaterPlayoutDelay"] = types.YLeaf{"Cvvoipcallhistorylowaterplayoutdelay", cvvoipcallhistoryentry.Cvvoipcallhistorylowaterplayoutdelay}
    cvvoipcallhistoryentry.EntityData.Leafs["cvVoIPCallHistoryReceiveDelay"] = types.YLeaf{"Cvvoipcallhistoryreceivedelay", cvvoipcallhistoryentry.Cvvoipcallhistoryreceivedelay}
    cvvoipcallhistoryentry.EntityData.Leafs["cvVoIPCallHistoryVADEnable"] = types.YLeaf{"Cvvoipcallhistoryvadenable", cvvoipcallhistoryentry.Cvvoipcallhistoryvadenable}
    cvvoipcallhistoryentry.EntityData.Leafs["cvVoIPCallHistoryCoderTypeRate"] = types.YLeaf{"Cvvoipcallhistorycodertyperate", cvvoipcallhistoryentry.Cvvoipcallhistorycodertyperate}
    cvvoipcallhistoryentry.EntityData.Leafs["cvVoIPCallHistoryIcpif"] = types.YLeaf{"Cvvoipcallhistoryicpif", cvvoipcallhistoryentry.Cvvoipcallhistoryicpif}
    cvvoipcallhistoryentry.EntityData.Leafs["cvVoIPCallHistoryLostPackets"] = types.YLeaf{"Cvvoipcallhistorylostpackets", cvvoipcallhistoryentry.Cvvoipcallhistorylostpackets}
    cvvoipcallhistoryentry.EntityData.Leafs["cvVoIPCallHistoryEarlyPackets"] = types.YLeaf{"Cvvoipcallhistoryearlypackets", cvvoipcallhistoryentry.Cvvoipcallhistoryearlypackets}
    cvvoipcallhistoryentry.EntityData.Leafs["cvVoIPCallHistoryLatePackets"] = types.YLeaf{"Cvvoipcallhistorylatepackets", cvvoipcallhistoryentry.Cvvoipcallhistorylatepackets}
    cvvoipcallhistoryentry.EntityData.Leafs["cvVoIPCallHistoryUsername"] = types.YLeaf{"Cvvoipcallhistoryusername", cvvoipcallhistoryentry.Cvvoipcallhistoryusername}
    cvvoipcallhistoryentry.EntityData.Leafs["cvVoIPCallHistoryProtocolCallId"] = types.YLeaf{"Cvvoipcallhistoryprotocolcallid", cvvoipcallhistoryentry.Cvvoipcallhistoryprotocolcallid}
    cvvoipcallhistoryentry.EntityData.Leafs["cvVoIPCallHistoryRemSigIPAddrT"] = types.YLeaf{"Cvvoipcallhistoryremsigipaddrt", cvvoipcallhistoryentry.Cvvoipcallhistoryremsigipaddrt}
    cvvoipcallhistoryentry.EntityData.Leafs["cvVoIPCallHistoryRemSigIPAddr"] = types.YLeaf{"Cvvoipcallhistoryremsigipaddr", cvvoipcallhistoryentry.Cvvoipcallhistoryremsigipaddr}
    cvvoipcallhistoryentry.EntityData.Leafs["cvVoIPCallHistoryRemSigPort"] = types.YLeaf{"Cvvoipcallhistoryremsigport", cvvoipcallhistoryentry.Cvvoipcallhistoryremsigport}
    cvvoipcallhistoryentry.EntityData.Leafs["cvVoIPCallHistoryRemMediaIPAddrT"] = types.YLeaf{"Cvvoipcallhistoryremmediaipaddrt", cvvoipcallhistoryentry.Cvvoipcallhistoryremmediaipaddrt}
    cvvoipcallhistoryentry.EntityData.Leafs["cvVoIPCallHistoryRemMediaIPAddr"] = types.YLeaf{"Cvvoipcallhistoryremmediaipaddr", cvvoipcallhistoryentry.Cvvoipcallhistoryremmediaipaddr}
    cvvoipcallhistoryentry.EntityData.Leafs["cvVoIPCallHistoryRemMediaPort"] = types.YLeaf{"Cvvoipcallhistoryremmediaport", cvvoipcallhistoryentry.Cvvoipcallhistoryremmediaport}
    cvvoipcallhistoryentry.EntityData.Leafs["cvVoIPCallHistorySRTPEnable"] = types.YLeaf{"Cvvoipcallhistorysrtpenable", cvvoipcallhistoryentry.Cvvoipcallhistorysrtpenable}
    cvvoipcallhistoryentry.EntityData.Leafs["cvVoIPCallHistoryFallbackIcpif"] = types.YLeaf{"Cvvoipcallhistoryfallbackicpif", cvvoipcallhistoryentry.Cvvoipcallhistoryfallbackicpif}
    cvvoipcallhistoryentry.EntityData.Leafs["cvVoIPCallHistoryFallbackLoss"] = types.YLeaf{"Cvvoipcallhistoryfallbackloss", cvvoipcallhistoryentry.Cvvoipcallhistoryfallbackloss}
    cvvoipcallhistoryentry.EntityData.Leafs["cvVoIPCallHistoryFallbackDelay"] = types.YLeaf{"Cvvoipcallhistoryfallbackdelay", cvvoipcallhistoryentry.Cvvoipcallhistoryfallbackdelay}
    cvvoipcallhistoryentry.EntityData.Leafs["cvVoIPCallHistoryOctetAligned"] = types.YLeaf{"Cvvoipcallhistoryoctetaligned", cvvoipcallhistoryentry.Cvvoipcallhistoryoctetaligned}
    cvvoipcallhistoryentry.EntityData.Leafs["cvVoIPCallHistoryBitRates"] = types.YLeaf{"Cvvoipcallhistorybitrates", cvvoipcallhistoryentry.Cvvoipcallhistorybitrates}
    cvvoipcallhistoryentry.EntityData.Leafs["cvVoIPCallHistoryModeChgPeriod"] = types.YLeaf{"Cvvoipcallhistorymodechgperiod", cvvoipcallhistoryentry.Cvvoipcallhistorymodechgperiod}
    cvvoipcallhistoryentry.EntityData.Leafs["cvVoIPCallHistoryModeChgNeighbor"] = types.YLeaf{"Cvvoipcallhistorymodechgneighbor", cvvoipcallhistoryentry.Cvvoipcallhistorymodechgneighbor}
    cvvoipcallhistoryentry.EntityData.Leafs["cvVoIPCallHistoryMaxPtime"] = types.YLeaf{"Cvvoipcallhistorymaxptime", cvvoipcallhistoryentry.Cvvoipcallhistorymaxptime}
    cvvoipcallhistoryentry.EntityData.Leafs["cvVoIPCallHistoryCRC"] = types.YLeaf{"Cvvoipcallhistorycrc", cvvoipcallhistoryentry.Cvvoipcallhistorycrc}
    cvvoipcallhistoryentry.EntityData.Leafs["cvVoIPCallHistoryRobustSorting"] = types.YLeaf{"Cvvoipcallhistoryrobustsorting", cvvoipcallhistoryentry.Cvvoipcallhistoryrobustsorting}
    cvvoipcallhistoryentry.EntityData.Leafs["cvVoIPCallHistoryEncap"] = types.YLeaf{"Cvvoipcallhistoryencap", cvvoipcallhistoryentry.Cvvoipcallhistoryencap}
    cvvoipcallhistoryentry.EntityData.Leafs["cvVoIPCallHistoryInterleaving"] = types.YLeaf{"Cvvoipcallhistoryinterleaving", cvvoipcallhistoryentry.Cvvoipcallhistoryinterleaving}
    cvvoipcallhistoryentry.EntityData.Leafs["cvVoIPCallHistoryPtime"] = types.YLeaf{"Cvvoipcallhistoryptime", cvvoipcallhistoryentry.Cvvoipcallhistoryptime}
    cvvoipcallhistoryentry.EntityData.Leafs["cvVoIPCallHistoryChannels"] = types.YLeaf{"Cvvoipcallhistorychannels", cvvoipcallhistoryentry.Cvvoipcallhistorychannels}
    cvvoipcallhistoryentry.EntityData.Leafs["cvVoIPCallHistoryCoderMode"] = types.YLeaf{"Cvvoipcallhistorycodermode", cvvoipcallhistoryentry.Cvvoipcallhistorycodermode}
    cvvoipcallhistoryentry.EntityData.Leafs["cvVoIPCallHistoryCallId"] = types.YLeaf{"Cvvoipcallhistorycallid", cvvoipcallhistoryentry.Cvvoipcallhistorycallid}
    cvvoipcallhistoryentry.EntityData.Leafs["cvVoIPCallHistoryCallReferenceId"] = types.YLeaf{"Cvvoipcallhistorycallreferenceid", cvvoipcallhistoryentry.Cvvoipcallhistorycallreferenceid}
    cvvoipcallhistoryentry.EntityData.Leafs["cvVoIPCallHistorySessionId"] = types.YLeaf{"Cvvoipcallhistorysessionid", cvvoipcallhistoryentry.Cvvoipcallhistorysessionid}
    return &(cvvoipcallhistoryentry.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_Cvcallratestatstable
// This table represents voice call rate measurement in various
// interval lengths defined by the 
// CvCallVolumeStatsIntvlType object.
// 
// Each interval may contain one or more entries to allow for
// detailed measurement to be collected.
type CISCOVOICEDIALCONTROLMIB_Cvcallratestatstable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This is a conceptual-row in cvCallRateStatsTable This entry is created at
    // the system initialization and is updated at every epoch based on
    // CvCallVolumeStatsIntvlType. The type is slice of
    // CISCOVOICEDIALCONTROLMIB_Cvcallratestatstable_Cvcallratestatsentry.
    Cvcallratestatsentry []CISCOVOICEDIALCONTROLMIB_Cvcallratestatstable_Cvcallratestatsentry
}

func (cvcallratestatstable *CISCOVOICEDIALCONTROLMIB_Cvcallratestatstable) GetEntityData() *types.CommonEntityData {
    cvcallratestatstable.EntityData.YFilter = cvcallratestatstable.YFilter
    cvcallratestatstable.EntityData.YangName = "cvCallRateStatsTable"
    cvcallratestatstable.EntityData.BundleName = "cisco_ios_xe"
    cvcallratestatstable.EntityData.ParentYangName = "CISCO-VOICE-DIAL-CONTROL-MIB"
    cvcallratestatstable.EntityData.SegmentPath = "cvCallRateStatsTable"
    cvcallratestatstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvcallratestatstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvcallratestatstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvcallratestatstable.EntityData.Children = make(map[string]types.YChild)
    cvcallratestatstable.EntityData.Children["cvCallRateStatsEntry"] = types.YChild{"Cvcallratestatsentry", nil}
    for i := range cvcallratestatstable.Cvcallratestatsentry {
        cvcallratestatstable.EntityData.Children[types.GetSegmentPath(&cvcallratestatstable.Cvcallratestatsentry[i])] = types.YChild{"Cvcallratestatsentry", &cvcallratestatstable.Cvcallratestatsentry[i]}
    }
    cvcallratestatstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cvcallratestatstable.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_Cvcallratestatstable_Cvcallratestatsentry
// This is a conceptual-row in cvCallRateStatsTable
// This entry is created at the system initialization and is
// updated at every epoch based on CvCallVolumeStatsIntvlType
type CISCOVOICEDIALCONTROLMIB_Cvcallratestatstable_Cvcallratestatsentry struct {
    EntityData types.CommonEntityData
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

func (cvcallratestatsentry *CISCOVOICEDIALCONTROLMIB_Cvcallratestatstable_Cvcallratestatsentry) GetEntityData() *types.CommonEntityData {
    cvcallratestatsentry.EntityData.YFilter = cvcallratestatsentry.YFilter
    cvcallratestatsentry.EntityData.YangName = "cvCallRateStatsEntry"
    cvcallratestatsentry.EntityData.BundleName = "cisco_ios_xe"
    cvcallratestatsentry.EntityData.ParentYangName = "cvCallRateStatsTable"
    cvcallratestatsentry.EntityData.SegmentPath = "cvCallRateStatsEntry" + "[cvCallRateStatsIntvlDurUnits='" + fmt.Sprintf("%v", cvcallratestatsentry.Cvcallratestatsintvldurunits) + "']" + "[cvCallRateStatsIntvlDur='" + fmt.Sprintf("%v", cvcallratestatsentry.Cvcallratestatsintvldur) + "']"
    cvcallratestatsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvcallratestatsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvcallratestatsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvcallratestatsentry.EntityData.Children = make(map[string]types.YChild)
    cvcallratestatsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cvcallratestatsentry.EntityData.Leafs["cvCallRateStatsIntvlDurUnits"] = types.YLeaf{"Cvcallratestatsintvldurunits", cvcallratestatsentry.Cvcallratestatsintvldurunits}
    cvcallratestatsentry.EntityData.Leafs["cvCallRateStatsIntvlDur"] = types.YLeaf{"Cvcallratestatsintvldur", cvcallratestatsentry.Cvcallratestatsintvldur}
    cvcallratestatsentry.EntityData.Leafs["cvCallRateStatsMaxVal"] = types.YLeaf{"Cvcallratestatsmaxval", cvcallratestatsentry.Cvcallratestatsmaxval}
    cvcallratestatsentry.EntityData.Leafs["cvCallRateStatsAvgVal"] = types.YLeaf{"Cvcallratestatsavgval", cvcallratestatsentry.Cvcallratestatsavgval}
    return &(cvcallratestatsentry.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_Cvcalllegratestatstable
// cvCallLegRateStatsTable table represents voice call leg rate
// measurement in various interval lengths defined by 
// the CvCallVolumeStatsIntvlType object.
// Each interval may contain one or more entries to allow for
// detailed measurement to be collected.
type CISCOVOICEDIALCONTROLMIB_Cvcalllegratestatstable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This is a conceptual-row in cvCallLegRateStatsTable This entry is created
    // at the system initialization and is updated at every epoch based on
    // CvCallVolumeStatsIntvlType. The type is slice of
    // CISCOVOICEDIALCONTROLMIB_Cvcalllegratestatstable_Cvcalllegratestatsentry.
    Cvcalllegratestatsentry []CISCOVOICEDIALCONTROLMIB_Cvcalllegratestatstable_Cvcalllegratestatsentry
}

func (cvcalllegratestatstable *CISCOVOICEDIALCONTROLMIB_Cvcalllegratestatstable) GetEntityData() *types.CommonEntityData {
    cvcalllegratestatstable.EntityData.YFilter = cvcalllegratestatstable.YFilter
    cvcalllegratestatstable.EntityData.YangName = "cvCallLegRateStatsTable"
    cvcalllegratestatstable.EntityData.BundleName = "cisco_ios_xe"
    cvcalllegratestatstable.EntityData.ParentYangName = "CISCO-VOICE-DIAL-CONTROL-MIB"
    cvcalllegratestatstable.EntityData.SegmentPath = "cvCallLegRateStatsTable"
    cvcalllegratestatstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvcalllegratestatstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvcalllegratestatstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvcalllegratestatstable.EntityData.Children = make(map[string]types.YChild)
    cvcalllegratestatstable.EntityData.Children["cvCallLegRateStatsEntry"] = types.YChild{"Cvcalllegratestatsentry", nil}
    for i := range cvcalllegratestatstable.Cvcalllegratestatsentry {
        cvcalllegratestatstable.EntityData.Children[types.GetSegmentPath(&cvcalllegratestatstable.Cvcalllegratestatsentry[i])] = types.YChild{"Cvcalllegratestatsentry", &cvcalllegratestatstable.Cvcalllegratestatsentry[i]}
    }
    cvcalllegratestatstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cvcalllegratestatstable.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_Cvcalllegratestatstable_Cvcalllegratestatsentry
// This is a conceptual-row in cvCallLegRateStatsTable
// This entry is created at the system initialization and is
// updated at every epoch based on CvCallVolumeStatsIntvlType
type CISCOVOICEDIALCONTROLMIB_Cvcalllegratestatstable_Cvcalllegratestatsentry struct {
    EntityData types.CommonEntityData
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

func (cvcalllegratestatsentry *CISCOVOICEDIALCONTROLMIB_Cvcalllegratestatstable_Cvcalllegratestatsentry) GetEntityData() *types.CommonEntityData {
    cvcalllegratestatsentry.EntityData.YFilter = cvcalllegratestatsentry.YFilter
    cvcalllegratestatsentry.EntityData.YangName = "cvCallLegRateStatsEntry"
    cvcalllegratestatsentry.EntityData.BundleName = "cisco_ios_xe"
    cvcalllegratestatsentry.EntityData.ParentYangName = "cvCallLegRateStatsTable"
    cvcalllegratestatsentry.EntityData.SegmentPath = "cvCallLegRateStatsEntry" + "[cvCallLegRateStatsIntvlDurUnits='" + fmt.Sprintf("%v", cvcalllegratestatsentry.Cvcalllegratestatsintvldurunits) + "']" + "[cvCallLegRateStatsIntvlDur='" + fmt.Sprintf("%v", cvcalllegratestatsentry.Cvcalllegratestatsintvldur) + "']"
    cvcalllegratestatsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvcalllegratestatsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvcalllegratestatsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvcalllegratestatsentry.EntityData.Children = make(map[string]types.YChild)
    cvcalllegratestatsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cvcalllegratestatsentry.EntityData.Leafs["cvCallLegRateStatsIntvlDurUnits"] = types.YLeaf{"Cvcalllegratestatsintvldurunits", cvcalllegratestatsentry.Cvcalllegratestatsintvldurunits}
    cvcalllegratestatsentry.EntityData.Leafs["cvCallLegRateStatsIntvlDur"] = types.YLeaf{"Cvcalllegratestatsintvldur", cvcalllegratestatsentry.Cvcalllegratestatsintvldur}
    cvcalllegratestatsentry.EntityData.Leafs["cvCallLegRateStatsMaxVal"] = types.YLeaf{"Cvcalllegratestatsmaxval", cvcalllegratestatsentry.Cvcalllegratestatsmaxval}
    cvcalllegratestatsentry.EntityData.Leafs["cvCallLegRateStatsAvgVal"] = types.YLeaf{"Cvcalllegratestatsavgval", cvcalllegratestatsentry.Cvcalllegratestatsavgval}
    return &(cvcalllegratestatsentry.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_Cvactivecallstatstable
// This table represents the active voice calls in various
// interval lengths defined by the 
// CvCallVolumeStatsIntvlType object.
// 
// Each interval may contain one or more entries to allow for
// detailed measurement to be collected.
type CISCOVOICEDIALCONTROLMIB_Cvactivecallstatstable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This is a conceptual-row in cvActiveCallStatsTable This entry is created at
    // the system initialization and is updated at every epoch based on
    // CvCallVolumeStatsIntvlType. The type is slice of
    // CISCOVOICEDIALCONTROLMIB_Cvactivecallstatstable_Cvactivecallstatsentry.
    Cvactivecallstatsentry []CISCOVOICEDIALCONTROLMIB_Cvactivecallstatstable_Cvactivecallstatsentry
}

func (cvactivecallstatstable *CISCOVOICEDIALCONTROLMIB_Cvactivecallstatstable) GetEntityData() *types.CommonEntityData {
    cvactivecallstatstable.EntityData.YFilter = cvactivecallstatstable.YFilter
    cvactivecallstatstable.EntityData.YangName = "cvActiveCallStatsTable"
    cvactivecallstatstable.EntityData.BundleName = "cisco_ios_xe"
    cvactivecallstatstable.EntityData.ParentYangName = "CISCO-VOICE-DIAL-CONTROL-MIB"
    cvactivecallstatstable.EntityData.SegmentPath = "cvActiveCallStatsTable"
    cvactivecallstatstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvactivecallstatstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvactivecallstatstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvactivecallstatstable.EntityData.Children = make(map[string]types.YChild)
    cvactivecallstatstable.EntityData.Children["cvActiveCallStatsEntry"] = types.YChild{"Cvactivecallstatsentry", nil}
    for i := range cvactivecallstatstable.Cvactivecallstatsentry {
        cvactivecallstatstable.EntityData.Children[types.GetSegmentPath(&cvactivecallstatstable.Cvactivecallstatsentry[i])] = types.YChild{"Cvactivecallstatsentry", &cvactivecallstatstable.Cvactivecallstatsentry[i]}
    }
    cvactivecallstatstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cvactivecallstatstable.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_Cvactivecallstatstable_Cvactivecallstatsentry
// This is a conceptual-row in cvActiveCallStatsTable
// This entry is created at the system initialization and is
// updated at every epoch based on CvCallVolumeStatsIntvlType
type CISCOVOICEDIALCONTROLMIB_Cvactivecallstatstable_Cvactivecallstatsentry struct {
    EntityData types.CommonEntityData
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

func (cvactivecallstatsentry *CISCOVOICEDIALCONTROLMIB_Cvactivecallstatstable_Cvactivecallstatsentry) GetEntityData() *types.CommonEntityData {
    cvactivecallstatsentry.EntityData.YFilter = cvactivecallstatsentry.YFilter
    cvactivecallstatsentry.EntityData.YangName = "cvActiveCallStatsEntry"
    cvactivecallstatsentry.EntityData.BundleName = "cisco_ios_xe"
    cvactivecallstatsentry.EntityData.ParentYangName = "cvActiveCallStatsTable"
    cvactivecallstatsentry.EntityData.SegmentPath = "cvActiveCallStatsEntry" + "[cvActiveCallStatsIntvlDurUnits='" + fmt.Sprintf("%v", cvactivecallstatsentry.Cvactivecallstatsintvldurunits) + "']" + "[cvActiveCallStatsIntvlDur='" + fmt.Sprintf("%v", cvactivecallstatsentry.Cvactivecallstatsintvldur) + "']"
    cvactivecallstatsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvactivecallstatsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvactivecallstatsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvactivecallstatsentry.EntityData.Children = make(map[string]types.YChild)
    cvactivecallstatsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cvactivecallstatsentry.EntityData.Leafs["cvActiveCallStatsIntvlDurUnits"] = types.YLeaf{"Cvactivecallstatsintvldurunits", cvactivecallstatsentry.Cvactivecallstatsintvldurunits}
    cvactivecallstatsentry.EntityData.Leafs["cvActiveCallStatsIntvlDur"] = types.YLeaf{"Cvactivecallstatsintvldur", cvactivecallstatsentry.Cvactivecallstatsintvldur}
    cvactivecallstatsentry.EntityData.Leafs["cvActiveCallStatsMaxVal"] = types.YLeaf{"Cvactivecallstatsmaxval", cvactivecallstatsentry.Cvactivecallstatsmaxval}
    cvactivecallstatsentry.EntityData.Leafs["cvActiveCallStatsAvgVal"] = types.YLeaf{"Cvactivecallstatsavgval", cvactivecallstatsentry.Cvactivecallstatsavgval}
    return &(cvactivecallstatsentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This is a conceptual-row in cvCallDurationStatsTable This entry is created
    // at the system initialization and is updated at every epoch based on
    // CvCallVolumeStatsIntvlType. The type is slice of
    // CISCOVOICEDIALCONTROLMIB_Cvcalldurationstatstable_Cvcalldurationstatsentry.
    Cvcalldurationstatsentry []CISCOVOICEDIALCONTROLMIB_Cvcalldurationstatstable_Cvcalldurationstatsentry
}

func (cvcalldurationstatstable *CISCOVOICEDIALCONTROLMIB_Cvcalldurationstatstable) GetEntityData() *types.CommonEntityData {
    cvcalldurationstatstable.EntityData.YFilter = cvcalldurationstatstable.YFilter
    cvcalldurationstatstable.EntityData.YangName = "cvCallDurationStatsTable"
    cvcalldurationstatstable.EntityData.BundleName = "cisco_ios_xe"
    cvcalldurationstatstable.EntityData.ParentYangName = "CISCO-VOICE-DIAL-CONTROL-MIB"
    cvcalldurationstatstable.EntityData.SegmentPath = "cvCallDurationStatsTable"
    cvcalldurationstatstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvcalldurationstatstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvcalldurationstatstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvcalldurationstatstable.EntityData.Children = make(map[string]types.YChild)
    cvcalldurationstatstable.EntityData.Children["cvCallDurationStatsEntry"] = types.YChild{"Cvcalldurationstatsentry", nil}
    for i := range cvcalldurationstatstable.Cvcalldurationstatsentry {
        cvcalldurationstatstable.EntityData.Children[types.GetSegmentPath(&cvcalldurationstatstable.Cvcalldurationstatsentry[i])] = types.YChild{"Cvcalldurationstatsentry", &cvcalldurationstatstable.Cvcalldurationstatsentry[i]}
    }
    cvcalldurationstatstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cvcalldurationstatstable.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_Cvcalldurationstatstable_Cvcalldurationstatsentry
// This is a conceptual-row in cvCallDurationStatsTable
// This entry is created at the system initialization and is
// updated at every epoch based on CvCallVolumeStatsIntvlType
type CISCOVOICEDIALCONTROLMIB_Cvcalldurationstatstable_Cvcalldurationstatsentry struct {
    EntityData types.CommonEntityData
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

func (cvcalldurationstatsentry *CISCOVOICEDIALCONTROLMIB_Cvcalldurationstatstable_Cvcalldurationstatsentry) GetEntityData() *types.CommonEntityData {
    cvcalldurationstatsentry.EntityData.YFilter = cvcalldurationstatsentry.YFilter
    cvcalldurationstatsentry.EntityData.YangName = "cvCallDurationStatsEntry"
    cvcalldurationstatsentry.EntityData.BundleName = "cisco_ios_xe"
    cvcalldurationstatsentry.EntityData.ParentYangName = "cvCallDurationStatsTable"
    cvcalldurationstatsentry.EntityData.SegmentPath = "cvCallDurationStatsEntry" + "[cvCallDurationStatsIntvlDurUnits='" + fmt.Sprintf("%v", cvcalldurationstatsentry.Cvcalldurationstatsintvldurunits) + "']" + "[cvCallDurationStatsIntvlDur='" + fmt.Sprintf("%v", cvcalldurationstatsentry.Cvcalldurationstatsintvldur) + "']"
    cvcalldurationstatsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvcalldurationstatsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvcalldurationstatsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvcalldurationstatsentry.EntityData.Children = make(map[string]types.YChild)
    cvcalldurationstatsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cvcalldurationstatsentry.EntityData.Leafs["cvCallDurationStatsIntvlDurUnits"] = types.YLeaf{"Cvcalldurationstatsintvldurunits", cvcalldurationstatsentry.Cvcalldurationstatsintvldurunits}
    cvcalldurationstatsentry.EntityData.Leafs["cvCallDurationStatsIntvlDur"] = types.YLeaf{"Cvcalldurationstatsintvldur", cvcalldurationstatsentry.Cvcalldurationstatsintvldur}
    cvcalldurationstatsentry.EntityData.Leafs["cvCallDurationStatsMaxVal"] = types.YLeaf{"Cvcalldurationstatsmaxval", cvcalldurationstatsentry.Cvcalldurationstatsmaxval}
    cvcalldurationstatsentry.EntityData.Leafs["cvCallDurationStatsAvgVal"] = types.YLeaf{"Cvcalldurationstatsavgval", cvcalldurationstatsentry.Cvcalldurationstatsavgval}
    return &(cvcalldurationstatsentry.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_Cvsipmsgratestatstable
// This table represents the SIP message rate measurement in
// various interval length defined by the 
// CvCallVolumeStatsIntvlType object.
// 
// Each interval may contain one or more entries to allow for
// detailed measurement to be collected
type CISCOVOICEDIALCONTROLMIB_Cvsipmsgratestatstable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This is a conceptual-row in cvSipMsgRateStatsTable This entry is created at
    // the system initialization and is updated at every epoch based on
    // CvCallVolumeStatsIntvlType. The type is slice of
    // CISCOVOICEDIALCONTROLMIB_Cvsipmsgratestatstable_Cvsipmsgratestatsentry.
    Cvsipmsgratestatsentry []CISCOVOICEDIALCONTROLMIB_Cvsipmsgratestatstable_Cvsipmsgratestatsentry
}

func (cvsipmsgratestatstable *CISCOVOICEDIALCONTROLMIB_Cvsipmsgratestatstable) GetEntityData() *types.CommonEntityData {
    cvsipmsgratestatstable.EntityData.YFilter = cvsipmsgratestatstable.YFilter
    cvsipmsgratestatstable.EntityData.YangName = "cvSipMsgRateStatsTable"
    cvsipmsgratestatstable.EntityData.BundleName = "cisco_ios_xe"
    cvsipmsgratestatstable.EntityData.ParentYangName = "CISCO-VOICE-DIAL-CONTROL-MIB"
    cvsipmsgratestatstable.EntityData.SegmentPath = "cvSipMsgRateStatsTable"
    cvsipmsgratestatstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvsipmsgratestatstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvsipmsgratestatstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvsipmsgratestatstable.EntityData.Children = make(map[string]types.YChild)
    cvsipmsgratestatstable.EntityData.Children["cvSipMsgRateStatsEntry"] = types.YChild{"Cvsipmsgratestatsentry", nil}
    for i := range cvsipmsgratestatstable.Cvsipmsgratestatsentry {
        cvsipmsgratestatstable.EntityData.Children[types.GetSegmentPath(&cvsipmsgratestatstable.Cvsipmsgratestatsentry[i])] = types.YChild{"Cvsipmsgratestatsentry", &cvsipmsgratestatstable.Cvsipmsgratestatsentry[i]}
    }
    cvsipmsgratestatstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cvsipmsgratestatstable.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_Cvsipmsgratestatstable_Cvsipmsgratestatsentry
// This is a conceptual-row in cvSipMsgRateStatsTable
// This entry is created at the system initialization and is
// updated at every epoch based on CvCallVolumeStatsIntvlType
type CISCOVOICEDIALCONTROLMIB_Cvsipmsgratestatstable_Cvsipmsgratestatsentry struct {
    EntityData types.CommonEntityData
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

func (cvsipmsgratestatsentry *CISCOVOICEDIALCONTROLMIB_Cvsipmsgratestatstable_Cvsipmsgratestatsentry) GetEntityData() *types.CommonEntityData {
    cvsipmsgratestatsentry.EntityData.YFilter = cvsipmsgratestatsentry.YFilter
    cvsipmsgratestatsentry.EntityData.YangName = "cvSipMsgRateStatsEntry"
    cvsipmsgratestatsentry.EntityData.BundleName = "cisco_ios_xe"
    cvsipmsgratestatsentry.EntityData.ParentYangName = "cvSipMsgRateStatsTable"
    cvsipmsgratestatsentry.EntityData.SegmentPath = "cvSipMsgRateStatsEntry" + "[cvSipMsgRateStatsIntvlDurUnits='" + fmt.Sprintf("%v", cvsipmsgratestatsentry.Cvsipmsgratestatsintvldurunits) + "']" + "[cvSipMsgRateStatsIntvlDur='" + fmt.Sprintf("%v", cvsipmsgratestatsentry.Cvsipmsgratestatsintvldur) + "']"
    cvsipmsgratestatsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvsipmsgratestatsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvsipmsgratestatsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvsipmsgratestatsentry.EntityData.Children = make(map[string]types.YChild)
    cvsipmsgratestatsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cvsipmsgratestatsentry.EntityData.Leafs["cvSipMsgRateStatsIntvlDurUnits"] = types.YLeaf{"Cvsipmsgratestatsintvldurunits", cvsipmsgratestatsentry.Cvsipmsgratestatsintvldurunits}
    cvsipmsgratestatsentry.EntityData.Leafs["cvSipMsgRateStatsIntvlDur"] = types.YLeaf{"Cvsipmsgratestatsintvldur", cvsipmsgratestatsentry.Cvsipmsgratestatsintvldur}
    cvsipmsgratestatsentry.EntityData.Leafs["cvSipMsgRateStatsMaxVal"] = types.YLeaf{"Cvsipmsgratestatsmaxval", cvsipmsgratestatsentry.Cvsipmsgratestatsmaxval}
    cvsipmsgratestatsentry.EntityData.Leafs["cvSipMsgRateStatsAvgVal"] = types.YLeaf{"Cvsipmsgratestatsavgval", cvsipmsgratestatsentry.Cvsipmsgratestatsavgval}
    return &(cvsipmsgratestatsentry.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_Cvcallratewmtable
// This table represents high watermarks achieved
// by call rate in various interval length defined 
// by CvCallVolumeWMIntvlType. 
// 
// Each interval may contain one or more entries to allow for 
// detailed measurement to be collected
type CISCOVOICEDIALCONTROLMIB_Cvcallratewmtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This is a conceptual-row in cvCallRateWMTable This entry is created at the
    // system initialization and is updated whenever  a) This entry is obsolete OR
    // b) A new/higher entry is available. These entries are
    // reinitialised/added/deleted  if cvCallVolumeWMTableSize is changed. The
    // type is slice of
    // CISCOVOICEDIALCONTROLMIB_Cvcallratewmtable_Cvcallratewmentry.
    Cvcallratewmentry []CISCOVOICEDIALCONTROLMIB_Cvcallratewmtable_Cvcallratewmentry
}

func (cvcallratewmtable *CISCOVOICEDIALCONTROLMIB_Cvcallratewmtable) GetEntityData() *types.CommonEntityData {
    cvcallratewmtable.EntityData.YFilter = cvcallratewmtable.YFilter
    cvcallratewmtable.EntityData.YangName = "cvCallRateWMTable"
    cvcallratewmtable.EntityData.BundleName = "cisco_ios_xe"
    cvcallratewmtable.EntityData.ParentYangName = "CISCO-VOICE-DIAL-CONTROL-MIB"
    cvcallratewmtable.EntityData.SegmentPath = "cvCallRateWMTable"
    cvcallratewmtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvcallratewmtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvcallratewmtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvcallratewmtable.EntityData.Children = make(map[string]types.YChild)
    cvcallratewmtable.EntityData.Children["cvCallRateWMEntry"] = types.YChild{"Cvcallratewmentry", nil}
    for i := range cvcallratewmtable.Cvcallratewmentry {
        cvcallratewmtable.EntityData.Children[types.GetSegmentPath(&cvcallratewmtable.Cvcallratewmentry[i])] = types.YChild{"Cvcallratewmentry", &cvcallratewmtable.Cvcallratewmentry[i]}
    }
    cvcallratewmtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cvcallratewmtable.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_Cvcallratewmtable_Cvcallratewmentry
// This is a conceptual-row in cvCallRateWMTable
// This entry is created at the system initialization and is
// updated whenever 
// a) This entry is obsolete OR
// b) A new/higher entry is available.
// These entries are reinitialised/added/deleted  if
// cvCallVolumeWMTableSize is changed
type CISCOVOICEDIALCONTROLMIB_Cvcallratewmtable_Cvcallratewmentry struct {
    EntityData types.CommonEntityData
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

func (cvcallratewmentry *CISCOVOICEDIALCONTROLMIB_Cvcallratewmtable_Cvcallratewmentry) GetEntityData() *types.CommonEntityData {
    cvcallratewmentry.EntityData.YFilter = cvcallratewmentry.YFilter
    cvcallratewmentry.EntityData.YangName = "cvCallRateWMEntry"
    cvcallratewmentry.EntityData.BundleName = "cisco_ios_xe"
    cvcallratewmentry.EntityData.ParentYangName = "cvCallRateWMTable"
    cvcallratewmentry.EntityData.SegmentPath = "cvCallRateWMEntry" + "[cvCallRateWMIntvlDurUnits='" + fmt.Sprintf("%v", cvcallratewmentry.Cvcallratewmintvldurunits) + "']" + "[cvCallRateWMIndex='" + fmt.Sprintf("%v", cvcallratewmentry.Cvcallratewmindex) + "']"
    cvcallratewmentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvcallratewmentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvcallratewmentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvcallratewmentry.EntityData.Children = make(map[string]types.YChild)
    cvcallratewmentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cvcallratewmentry.EntityData.Leafs["cvCallRateWMIntvlDurUnits"] = types.YLeaf{"Cvcallratewmintvldurunits", cvcallratewmentry.Cvcallratewmintvldurunits}
    cvcallratewmentry.EntityData.Leafs["cvCallRateWMIndex"] = types.YLeaf{"Cvcallratewmindex", cvcallratewmentry.Cvcallratewmindex}
    cvcallratewmentry.EntityData.Leafs["cvCallRateWMValue"] = types.YLeaf{"Cvcallratewmvalue", cvcallratewmentry.Cvcallratewmvalue}
    cvcallratewmentry.EntityData.Leafs["cvCallRateWMts"] = types.YLeaf{"Cvcallratewmts", cvcallratewmentry.Cvcallratewmts}
    return &(cvcallratewmentry.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_Cvcalllegratewmtable
// cvCallLegRateWMTable table represents high watermarks achieved
// by call-leg rate in various interval length defined 
// by CvCallVolumeWMIntvlType. 
// 
// Each interval may contain one or more entries to allow for 
// detailed measurement to be collected
type CISCOVOICEDIALCONTROLMIB_Cvcalllegratewmtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This is a conceptual-row in cvCallLegRateWMTable This entry is created at
    // the system initialization and is updated whenever  a) This entry is
    // obsolete OR b) A new/higher entry is available. These entries are
    // reinitialised/added/deleted  if cvCallVolumeWMTableSize is changed. The
    // type is slice of
    // CISCOVOICEDIALCONTROLMIB_Cvcalllegratewmtable_Cvcalllegratewmentry.
    Cvcalllegratewmentry []CISCOVOICEDIALCONTROLMIB_Cvcalllegratewmtable_Cvcalllegratewmentry
}

func (cvcalllegratewmtable *CISCOVOICEDIALCONTROLMIB_Cvcalllegratewmtable) GetEntityData() *types.CommonEntityData {
    cvcalllegratewmtable.EntityData.YFilter = cvcalllegratewmtable.YFilter
    cvcalllegratewmtable.EntityData.YangName = "cvCallLegRateWMTable"
    cvcalllegratewmtable.EntityData.BundleName = "cisco_ios_xe"
    cvcalllegratewmtable.EntityData.ParentYangName = "CISCO-VOICE-DIAL-CONTROL-MIB"
    cvcalllegratewmtable.EntityData.SegmentPath = "cvCallLegRateWMTable"
    cvcalllegratewmtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvcalllegratewmtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvcalllegratewmtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvcalllegratewmtable.EntityData.Children = make(map[string]types.YChild)
    cvcalllegratewmtable.EntityData.Children["cvCallLegRateWMEntry"] = types.YChild{"Cvcalllegratewmentry", nil}
    for i := range cvcalllegratewmtable.Cvcalllegratewmentry {
        cvcalllegratewmtable.EntityData.Children[types.GetSegmentPath(&cvcalllegratewmtable.Cvcalllegratewmentry[i])] = types.YChild{"Cvcalllegratewmentry", &cvcalllegratewmtable.Cvcalllegratewmentry[i]}
    }
    cvcalllegratewmtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cvcalllegratewmtable.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_Cvcalllegratewmtable_Cvcalllegratewmentry
// This is a conceptual-row in cvCallLegRateWMTable
// This entry is created at the system initialization and is
// updated whenever 
// a) This entry is obsolete OR
// b) A new/higher entry is available.
// These entries are reinitialised/added/deleted  if
// cvCallVolumeWMTableSize is changed
type CISCOVOICEDIALCONTROLMIB_Cvcalllegratewmtable_Cvcalllegratewmentry struct {
    EntityData types.CommonEntityData
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

func (cvcalllegratewmentry *CISCOVOICEDIALCONTROLMIB_Cvcalllegratewmtable_Cvcalllegratewmentry) GetEntityData() *types.CommonEntityData {
    cvcalllegratewmentry.EntityData.YFilter = cvcalllegratewmentry.YFilter
    cvcalllegratewmentry.EntityData.YangName = "cvCallLegRateWMEntry"
    cvcalllegratewmentry.EntityData.BundleName = "cisco_ios_xe"
    cvcalllegratewmentry.EntityData.ParentYangName = "cvCallLegRateWMTable"
    cvcalllegratewmentry.EntityData.SegmentPath = "cvCallLegRateWMEntry" + "[cvCallLegRateWMIntvlDurUnits='" + fmt.Sprintf("%v", cvcalllegratewmentry.Cvcalllegratewmintvldurunits) + "']" + "[cvCallLegRateWMIndex='" + fmt.Sprintf("%v", cvcalllegratewmentry.Cvcalllegratewmindex) + "']"
    cvcalllegratewmentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvcalllegratewmentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvcalllegratewmentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvcalllegratewmentry.EntityData.Children = make(map[string]types.YChild)
    cvcalllegratewmentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cvcalllegratewmentry.EntityData.Leafs["cvCallLegRateWMIntvlDurUnits"] = types.YLeaf{"Cvcalllegratewmintvldurunits", cvcalllegratewmentry.Cvcalllegratewmintvldurunits}
    cvcalllegratewmentry.EntityData.Leafs["cvCallLegRateWMIndex"] = types.YLeaf{"Cvcalllegratewmindex", cvcalllegratewmentry.Cvcalllegratewmindex}
    cvcalllegratewmentry.EntityData.Leafs["cvCallLegRateWMValue"] = types.YLeaf{"Cvcalllegratewmvalue", cvcalllegratewmentry.Cvcalllegratewmvalue}
    cvcalllegratewmentry.EntityData.Leafs["cvCallLegRateWMts"] = types.YLeaf{"Cvcalllegratewmts", cvcalllegratewmentry.Cvcalllegratewmts}
    return &(cvcalllegratewmentry.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_Cvactivecallwmtable
// This table represents high watermarks achieved
// by active calls in various interval length defined 
// by CvCallVolumeWMIntvlType. 
// 
// Each interval may contain one or more entries to allow 
// for detailed measurement to be collected.
type CISCOVOICEDIALCONTROLMIB_Cvactivecallwmtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This is a conceptual-row in cvActiveCallWMTable This entry is created at
    // the system initialization and is updated whenever  a) This entry is
    // obsolete OR b) A new/higher entry is available. These entries are
    // reinitialised/added/deleted  if cvCallVolumeWMTableSize is changed. The
    // type is slice of
    // CISCOVOICEDIALCONTROLMIB_Cvactivecallwmtable_Cvactivecallwmentry.
    Cvactivecallwmentry []CISCOVOICEDIALCONTROLMIB_Cvactivecallwmtable_Cvactivecallwmentry
}

func (cvactivecallwmtable *CISCOVOICEDIALCONTROLMIB_Cvactivecallwmtable) GetEntityData() *types.CommonEntityData {
    cvactivecallwmtable.EntityData.YFilter = cvactivecallwmtable.YFilter
    cvactivecallwmtable.EntityData.YangName = "cvActiveCallWMTable"
    cvactivecallwmtable.EntityData.BundleName = "cisco_ios_xe"
    cvactivecallwmtable.EntityData.ParentYangName = "CISCO-VOICE-DIAL-CONTROL-MIB"
    cvactivecallwmtable.EntityData.SegmentPath = "cvActiveCallWMTable"
    cvactivecallwmtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvactivecallwmtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvactivecallwmtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvactivecallwmtable.EntityData.Children = make(map[string]types.YChild)
    cvactivecallwmtable.EntityData.Children["cvActiveCallWMEntry"] = types.YChild{"Cvactivecallwmentry", nil}
    for i := range cvactivecallwmtable.Cvactivecallwmentry {
        cvactivecallwmtable.EntityData.Children[types.GetSegmentPath(&cvactivecallwmtable.Cvactivecallwmentry[i])] = types.YChild{"Cvactivecallwmentry", &cvactivecallwmtable.Cvactivecallwmentry[i]}
    }
    cvactivecallwmtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cvactivecallwmtable.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_Cvactivecallwmtable_Cvactivecallwmentry
// This is a conceptual-row in cvActiveCallWMTable
// This entry is created at the system initialization and is
// updated whenever 
// a) This entry is obsolete OR
// b) A new/higher entry is available.
// These entries are reinitialised/added/deleted  if
// cvCallVolumeWMTableSize is changed
type CISCOVOICEDIALCONTROLMIB_Cvactivecallwmtable_Cvactivecallwmentry struct {
    EntityData types.CommonEntityData
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

func (cvactivecallwmentry *CISCOVOICEDIALCONTROLMIB_Cvactivecallwmtable_Cvactivecallwmentry) GetEntityData() *types.CommonEntityData {
    cvactivecallwmentry.EntityData.YFilter = cvactivecallwmentry.YFilter
    cvactivecallwmentry.EntityData.YangName = "cvActiveCallWMEntry"
    cvactivecallwmentry.EntityData.BundleName = "cisco_ios_xe"
    cvactivecallwmentry.EntityData.ParentYangName = "cvActiveCallWMTable"
    cvactivecallwmentry.EntityData.SegmentPath = "cvActiveCallWMEntry" + "[cvActiveCallWMIntvlDurUnits='" + fmt.Sprintf("%v", cvactivecallwmentry.Cvactivecallwmintvldurunits) + "']" + "[cvActiveCallWMIndex='" + fmt.Sprintf("%v", cvactivecallwmentry.Cvactivecallwmindex) + "']"
    cvactivecallwmentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvactivecallwmentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvactivecallwmentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvactivecallwmentry.EntityData.Children = make(map[string]types.YChild)
    cvactivecallwmentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cvactivecallwmentry.EntityData.Leafs["cvActiveCallWMIntvlDurUnits"] = types.YLeaf{"Cvactivecallwmintvldurunits", cvactivecallwmentry.Cvactivecallwmintvldurunits}
    cvactivecallwmentry.EntityData.Leafs["cvActiveCallWMIndex"] = types.YLeaf{"Cvactivecallwmindex", cvactivecallwmentry.Cvactivecallwmindex}
    cvactivecallwmentry.EntityData.Leafs["cvActiveCallWMValue"] = types.YLeaf{"Cvactivecallwmvalue", cvactivecallwmentry.Cvactivecallwmvalue}
    cvactivecallwmentry.EntityData.Leafs["cvActiveCallWMts"] = types.YLeaf{"Cvactivecallwmts", cvactivecallwmentry.Cvactivecallwmts}
    return &(cvactivecallwmentry.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_Cvsipmsgratewmtable
// This table represents of high watermarks achieved
// by SIP message rate in various interval length defined 
// by CvCallVolumeWMIntvlType. 
// 
// Each interval may contain one or more entries to allow for
// detailed measurement to be collected
type CISCOVOICEDIALCONTROLMIB_Cvsipmsgratewmtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This is a conceptual-row in cvSipMsgRateWMTable. This entry is created at
    // the system initialization and is updated whenever  a) This entry is
    // obsolete OR b) A new/higher entry is available. These entries are
    // reinitialised/added/deleted if cvCallVolumeWMTableSize is changed. The type
    // is slice of
    // CISCOVOICEDIALCONTROLMIB_Cvsipmsgratewmtable_Cvsipmsgratewmentry.
    Cvsipmsgratewmentry []CISCOVOICEDIALCONTROLMIB_Cvsipmsgratewmtable_Cvsipmsgratewmentry
}

func (cvsipmsgratewmtable *CISCOVOICEDIALCONTROLMIB_Cvsipmsgratewmtable) GetEntityData() *types.CommonEntityData {
    cvsipmsgratewmtable.EntityData.YFilter = cvsipmsgratewmtable.YFilter
    cvsipmsgratewmtable.EntityData.YangName = "cvSipMsgRateWMTable"
    cvsipmsgratewmtable.EntityData.BundleName = "cisco_ios_xe"
    cvsipmsgratewmtable.EntityData.ParentYangName = "CISCO-VOICE-DIAL-CONTROL-MIB"
    cvsipmsgratewmtable.EntityData.SegmentPath = "cvSipMsgRateWMTable"
    cvsipmsgratewmtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvsipmsgratewmtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvsipmsgratewmtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvsipmsgratewmtable.EntityData.Children = make(map[string]types.YChild)
    cvsipmsgratewmtable.EntityData.Children["cvSipMsgRateWMEntry"] = types.YChild{"Cvsipmsgratewmentry", nil}
    for i := range cvsipmsgratewmtable.Cvsipmsgratewmentry {
        cvsipmsgratewmtable.EntityData.Children[types.GetSegmentPath(&cvsipmsgratewmtable.Cvsipmsgratewmentry[i])] = types.YChild{"Cvsipmsgratewmentry", &cvsipmsgratewmtable.Cvsipmsgratewmentry[i]}
    }
    cvsipmsgratewmtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cvsipmsgratewmtable.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_Cvsipmsgratewmtable_Cvsipmsgratewmentry
// This is a conceptual-row in cvSipMsgRateWMTable.
// This entry is created at the system initialization and is
// updated whenever 
// a) This entry is obsolete OR
// b) A new/higher entry is available.
// These entries are reinitialised/added/deleted if
// cvCallVolumeWMTableSize is changed
type CISCOVOICEDIALCONTROLMIB_Cvsipmsgratewmtable_Cvsipmsgratewmentry struct {
    EntityData types.CommonEntityData
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

func (cvsipmsgratewmentry *CISCOVOICEDIALCONTROLMIB_Cvsipmsgratewmtable_Cvsipmsgratewmentry) GetEntityData() *types.CommonEntityData {
    cvsipmsgratewmentry.EntityData.YFilter = cvsipmsgratewmentry.YFilter
    cvsipmsgratewmentry.EntityData.YangName = "cvSipMsgRateWMEntry"
    cvsipmsgratewmentry.EntityData.BundleName = "cisco_ios_xe"
    cvsipmsgratewmentry.EntityData.ParentYangName = "cvSipMsgRateWMTable"
    cvsipmsgratewmentry.EntityData.SegmentPath = "cvSipMsgRateWMEntry" + "[cvSipMsgRateWMIntvlDurUnits='" + fmt.Sprintf("%v", cvsipmsgratewmentry.Cvsipmsgratewmintvldurunits) + "']" + "[cvSipMsgRateWMIndex='" + fmt.Sprintf("%v", cvsipmsgratewmentry.Cvsipmsgratewmindex) + "']"
    cvsipmsgratewmentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvsipmsgratewmentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvsipmsgratewmentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvsipmsgratewmentry.EntityData.Children = make(map[string]types.YChild)
    cvsipmsgratewmentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cvsipmsgratewmentry.EntityData.Leafs["cvSipMsgRateWMIntvlDurUnits"] = types.YLeaf{"Cvsipmsgratewmintvldurunits", cvsipmsgratewmentry.Cvsipmsgratewmintvldurunits}
    cvsipmsgratewmentry.EntityData.Leafs["cvSipMsgRateWMIndex"] = types.YLeaf{"Cvsipmsgratewmindex", cvsipmsgratewmentry.Cvsipmsgratewmindex}
    cvsipmsgratewmentry.EntityData.Leafs["cvSipMsgRateWMValue"] = types.YLeaf{"Cvsipmsgratewmvalue", cvsipmsgratewmentry.Cvsipmsgratewmvalue}
    cvsipmsgratewmentry.EntityData.Leafs["cvSipMsgRateWMts"] = types.YLeaf{"Cvsipmsgratewmts", cvsipmsgratewmentry.Cvsipmsgratewmts}
    return &(cvsipmsgratewmentry.EntityData)
}

