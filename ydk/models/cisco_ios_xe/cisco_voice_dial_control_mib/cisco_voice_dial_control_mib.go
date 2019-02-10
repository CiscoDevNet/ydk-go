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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    CvGeneralConfiguration CISCOVOICEDIALCONTROLMIB_CvGeneralConfiguration

    
    CvGatewayCallActive CISCOVOICEDIALCONTROLMIB_CvGatewayCallActive

    
    CvCallVolume CISCOVOICEDIALCONTROLMIB_CvCallVolume

    
    CvCallRateMonitor CISCOVOICEDIALCONTROLMIB_CvCallRateMonitor

    
    CvCallVolumeStatsHistory CISCOVOICEDIALCONTROLMIB_CvCallVolumeStatsHistory

    // The table contains the Voice Generic Peer information that is used to
    // create an ifIndexed row with an appropriate ifType that is associated with
    // the cvPeerCfgType and cvPeerCfgPeerType objects.
    CvPeerCfgTable CISCOVOICEDIALCONTROLMIB_CvPeerCfgTable

    // The table contains the Voice over Telephony peer specific information that
    // is required to accept voice calls or to which it will place them or perform
    // various loopback tests via interface.
    CvVoicePeerCfgTable CISCOVOICEDIALCONTROLMIB_CvVoicePeerCfgTable

    // The table contains the Voice over IP (VoIP) peer specific information that
    // is required to accept voice calls or to which it will place them via IP
    // backbone with the specified session protocol in
    // cvVoIPPeerCfgSessionProtocol.
    CvVoIPPeerCfgTable CISCOVOICEDIALCONTROLMIB_CvVoIPPeerCfgTable

    // The table contains the Voice specific peer common configuration information
    // that is required to accept voice calls or to which it will place them or
    // process the incoming calls.
    CvPeerCommonCfgTable CISCOVOICEDIALCONTROLMIB_CvPeerCommonCfgTable

    // This table is the voice extension to the call active table of IETF Dial
    // Control MIB. It contains voice encapsulation call leg information that is
    // derived from the statistics of lower layer telephony interface.
    CvCallActiveTable CISCOVOICEDIALCONTROLMIB_CvCallActiveTable

    // This table is the VoIP extension to the call active table of IETF Dial
    // Control MIB. It contains VoIP call leg information about specific VoIP call
    // destination and the selected QoS for the call leg.
    CvVoIPCallActiveTable CISCOVOICEDIALCONTROLMIB_CvVoIPCallActiveTable

    // This table represents the number of active call connections for each call
    // connection type in the voice gateway.
    CvCallVolConnTable CISCOVOICEDIALCONTROLMIB_CvCallVolConnTable

    // This table represents the information about the usage of an IP interface in
    // a voice gateway for voice media calls. This table has a sparse-dependent
    // relationship with   ifTable. There exists an entry in this table,  for each
    // of the  entries in ifTable where ifType  is one of 'ethernetCsmacd' and
    // 'softwareLoopback'.
    CvCallVolIfTable CISCOVOICEDIALCONTROLMIB_CvCallVolIfTable

    // This table is the voice extension to the call history table of IETF Dial
    // Control MIB. It contains voice encapsulation call leg information such as
    // voice packet statistics, coder usage and end to end bandwidth of the call
    // leg.
    CvCallHistoryTable CISCOVOICEDIALCONTROLMIB_CvCallHistoryTable

    // This table is the VoIP extension to the call history table of IETF Dial
    // Control MIB. It contains VoIP call leg information about specific VoIP call
    // destination and the selected QoS for the call leg.
    CvVoIPCallHistoryTable CISCOVOICEDIALCONTROLMIB_CvVoIPCallHistoryTable

    // This table represents voice call rate measurement in various interval
    // lengths defined by the  CvCallVolumeStatsIntvlType object.  Each interval
    // may contain one or more entries to allow for detailed measurement to be
    // collected.
    CvCallRateStatsTable CISCOVOICEDIALCONTROLMIB_CvCallRateStatsTable

    // cvCallLegRateStatsTable table represents voice call leg rate measurement in
    // various interval lengths defined by  the CvCallVolumeStatsIntvlType object.
    // Each interval may contain one or more entries to allow for detailed
    // measurement to be collected.
    CvCallLegRateStatsTable CISCOVOICEDIALCONTROLMIB_CvCallLegRateStatsTable

    // This table represents the active voice calls in various interval lengths
    // defined by the  CvCallVolumeStatsIntvlType object.  Each interval may
    // contain one or more entries to allow for detailed measurement to be
    // collected.
    CvActiveCallStatsTable CISCOVOICEDIALCONTROLMIB_CvActiveCallStatsTable

    // This table represents the number of calls below a specific duration in
    // various interval length defined by  the CvCallVolumeStatsIntvlType object. 
    // The specific duration is configurable value of  
    // cvCallDurationStatsThreshold object.  Each interval may contain one or more
    // entries to allow for  detailed measurement to be collected.
    CvCallDurationStatsTable CISCOVOICEDIALCONTROLMIB_CvCallDurationStatsTable

    // This table represents the SIP message rate measurement in various interval
    // length defined by the  CvCallVolumeStatsIntvlType object.  Each interval
    // may contain one or more entries to allow for detailed measurement to be
    // collected.
    CvSipMsgRateStatsTable CISCOVOICEDIALCONTROLMIB_CvSipMsgRateStatsTable

    // This table represents high watermarks achieved by call rate in various
    // interval length defined  by CvCallVolumeWMIntvlType.   Each interval may
    // contain one or more entries to allow for  detailed measurement to be
    // collected.
    CvCallRateWMTable CISCOVOICEDIALCONTROLMIB_CvCallRateWMTable

    // cvCallLegRateWMTable table represents high watermarks achieved by call-leg
    // rate in various interval length defined  by CvCallVolumeWMIntvlType.   Each
    // interval may contain one or more entries to allow for  detailed measurement
    // to be collected.
    CvCallLegRateWMTable CISCOVOICEDIALCONTROLMIB_CvCallLegRateWMTable

    // This table represents high watermarks achieved by active calls in various
    // interval length defined  by CvCallVolumeWMIntvlType.   Each interval may
    // contain one or more entries to allow  for detailed measurement to be
    // collected.
    CvActiveCallWMTable CISCOVOICEDIALCONTROLMIB_CvActiveCallWMTable

    // This table represents of high watermarks achieved by SIP message rate in
    // various interval length defined  by CvCallVolumeWMIntvlType.   Each
    // interval may contain one or more entries to allow for detailed measurement
    // to be collected.
    CvSipMsgRateWMTable CISCOVOICEDIALCONTROLMIB_CvSipMsgRateWMTable
}

func (cISCOVOICEDIALCONTROLMIB *CISCOVOICEDIALCONTROLMIB) GetEntityData() *types.CommonEntityData {
    cISCOVOICEDIALCONTROLMIB.EntityData.YFilter = cISCOVOICEDIALCONTROLMIB.YFilter
    cISCOVOICEDIALCONTROLMIB.EntityData.YangName = "CISCO-VOICE-DIAL-CONTROL-MIB"
    cISCOVOICEDIALCONTROLMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOVOICEDIALCONTROLMIB.EntityData.ParentYangName = "CISCO-VOICE-DIAL-CONTROL-MIB"
    cISCOVOICEDIALCONTROLMIB.EntityData.SegmentPath = "CISCO-VOICE-DIAL-CONTROL-MIB:CISCO-VOICE-DIAL-CONTROL-MIB"
    cISCOVOICEDIALCONTROLMIB.EntityData.AbsolutePath = cISCOVOICEDIALCONTROLMIB.EntityData.SegmentPath
    cISCOVOICEDIALCONTROLMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOVOICEDIALCONTROLMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOVOICEDIALCONTROLMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOVOICEDIALCONTROLMIB.EntityData.Children = types.NewOrderedMap()
    cISCOVOICEDIALCONTROLMIB.EntityData.Children.Append("cvGeneralConfiguration", types.YChild{"CvGeneralConfiguration", &cISCOVOICEDIALCONTROLMIB.CvGeneralConfiguration})
    cISCOVOICEDIALCONTROLMIB.EntityData.Children.Append("cvGatewayCallActive", types.YChild{"CvGatewayCallActive", &cISCOVOICEDIALCONTROLMIB.CvGatewayCallActive})
    cISCOVOICEDIALCONTROLMIB.EntityData.Children.Append("cvCallVolume", types.YChild{"CvCallVolume", &cISCOVOICEDIALCONTROLMIB.CvCallVolume})
    cISCOVOICEDIALCONTROLMIB.EntityData.Children.Append("cvCallRateMonitor", types.YChild{"CvCallRateMonitor", &cISCOVOICEDIALCONTROLMIB.CvCallRateMonitor})
    cISCOVOICEDIALCONTROLMIB.EntityData.Children.Append("cvCallVolumeStatsHistory", types.YChild{"CvCallVolumeStatsHistory", &cISCOVOICEDIALCONTROLMIB.CvCallVolumeStatsHistory})
    cISCOVOICEDIALCONTROLMIB.EntityData.Children.Append("cvPeerCfgTable", types.YChild{"CvPeerCfgTable", &cISCOVOICEDIALCONTROLMIB.CvPeerCfgTable})
    cISCOVOICEDIALCONTROLMIB.EntityData.Children.Append("cvVoicePeerCfgTable", types.YChild{"CvVoicePeerCfgTable", &cISCOVOICEDIALCONTROLMIB.CvVoicePeerCfgTable})
    cISCOVOICEDIALCONTROLMIB.EntityData.Children.Append("cvVoIPPeerCfgTable", types.YChild{"CvVoIPPeerCfgTable", &cISCOVOICEDIALCONTROLMIB.CvVoIPPeerCfgTable})
    cISCOVOICEDIALCONTROLMIB.EntityData.Children.Append("cvPeerCommonCfgTable", types.YChild{"CvPeerCommonCfgTable", &cISCOVOICEDIALCONTROLMIB.CvPeerCommonCfgTable})
    cISCOVOICEDIALCONTROLMIB.EntityData.Children.Append("cvCallActiveTable", types.YChild{"CvCallActiveTable", &cISCOVOICEDIALCONTROLMIB.CvCallActiveTable})
    cISCOVOICEDIALCONTROLMIB.EntityData.Children.Append("cvVoIPCallActiveTable", types.YChild{"CvVoIPCallActiveTable", &cISCOVOICEDIALCONTROLMIB.CvVoIPCallActiveTable})
    cISCOVOICEDIALCONTROLMIB.EntityData.Children.Append("cvCallVolConnTable", types.YChild{"CvCallVolConnTable", &cISCOVOICEDIALCONTROLMIB.CvCallVolConnTable})
    cISCOVOICEDIALCONTROLMIB.EntityData.Children.Append("cvCallVolIfTable", types.YChild{"CvCallVolIfTable", &cISCOVOICEDIALCONTROLMIB.CvCallVolIfTable})
    cISCOVOICEDIALCONTROLMIB.EntityData.Children.Append("cvCallHistoryTable", types.YChild{"CvCallHistoryTable", &cISCOVOICEDIALCONTROLMIB.CvCallHistoryTable})
    cISCOVOICEDIALCONTROLMIB.EntityData.Children.Append("cvVoIPCallHistoryTable", types.YChild{"CvVoIPCallHistoryTable", &cISCOVOICEDIALCONTROLMIB.CvVoIPCallHistoryTable})
    cISCOVOICEDIALCONTROLMIB.EntityData.Children.Append("cvCallRateStatsTable", types.YChild{"CvCallRateStatsTable", &cISCOVOICEDIALCONTROLMIB.CvCallRateStatsTable})
    cISCOVOICEDIALCONTROLMIB.EntityData.Children.Append("cvCallLegRateStatsTable", types.YChild{"CvCallLegRateStatsTable", &cISCOVOICEDIALCONTROLMIB.CvCallLegRateStatsTable})
    cISCOVOICEDIALCONTROLMIB.EntityData.Children.Append("cvActiveCallStatsTable", types.YChild{"CvActiveCallStatsTable", &cISCOVOICEDIALCONTROLMIB.CvActiveCallStatsTable})
    cISCOVOICEDIALCONTROLMIB.EntityData.Children.Append("cvCallDurationStatsTable", types.YChild{"CvCallDurationStatsTable", &cISCOVOICEDIALCONTROLMIB.CvCallDurationStatsTable})
    cISCOVOICEDIALCONTROLMIB.EntityData.Children.Append("cvSipMsgRateStatsTable", types.YChild{"CvSipMsgRateStatsTable", &cISCOVOICEDIALCONTROLMIB.CvSipMsgRateStatsTable})
    cISCOVOICEDIALCONTROLMIB.EntityData.Children.Append("cvCallRateWMTable", types.YChild{"CvCallRateWMTable", &cISCOVOICEDIALCONTROLMIB.CvCallRateWMTable})
    cISCOVOICEDIALCONTROLMIB.EntityData.Children.Append("cvCallLegRateWMTable", types.YChild{"CvCallLegRateWMTable", &cISCOVOICEDIALCONTROLMIB.CvCallLegRateWMTable})
    cISCOVOICEDIALCONTROLMIB.EntityData.Children.Append("cvActiveCallWMTable", types.YChild{"CvActiveCallWMTable", &cISCOVOICEDIALCONTROLMIB.CvActiveCallWMTable})
    cISCOVOICEDIALCONTROLMIB.EntityData.Children.Append("cvSipMsgRateWMTable", types.YChild{"CvSipMsgRateWMTable", &cISCOVOICEDIALCONTROLMIB.CvSipMsgRateWMTable})
    cISCOVOICEDIALCONTROLMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOVOICEDIALCONTROLMIB.EntityData.YListKeys = []string {}

    return &(cISCOVOICEDIALCONTROLMIB.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_CvGeneralConfiguration
type CISCOVOICEDIALCONTROLMIB_CvGeneralConfiguration struct {
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
    CvGeneralPoorQoVNotificationEnable interface{}

    // This object indicates whether cvdcFallbackNotification traps should be
    // generated for fallback. If the value of this object is 'true',
    // cvdcFallbackNotification traps will be generated for all voice over IP
    // peers. The type is bool.
    CvGeneralFallbackNotificationEnable interface{}

    // This object indicates whether cvdcPolicyViolationNotification traps should
    // be generated for a RPH to DSCP mapping violation for SIP voice calls.  If
    // the value of this object is 'true', cvdcPolicyViolationNotification traps
    // will be generated for SIP voice over IP peers when a RPH to DSCP violation
    // condition is detected .  If the value of this object is 'false',
    // cvdcPolicyViolationNotification traps will be generated only for calls for
    // which the  cvVoIPPeerCfgDSCPPolicyNotificationEnable object of voice over
    // IP peers having set to 'true'. The type is bool.
    CvGeneralDSCPPolicyNotificationEnable interface{}

    // This object indicates whether cvdcPolicyViolationNotification traps should
    // be generated for Media violation for SIP voice calls.  If the value of this
    // object is 'true', cvdcPolicyViolationNotification traps will be generated
    // for SIP voice over IP peers when media violation condition is detected . 
    // If the value of this object is 'false', cvdcPolicyViolationNotification
    // traps will be generated only for calls for which the 
    // cvVoIPPeerCfgMediaPolicyNotificationEnable object of voice over IP peers
    // having set to 'true'. The type is bool.
    CvGeneralMediaPolicyNotificationEnable interface{}
}

func (cvGeneralConfiguration *CISCOVOICEDIALCONTROLMIB_CvGeneralConfiguration) GetEntityData() *types.CommonEntityData {
    cvGeneralConfiguration.EntityData.YFilter = cvGeneralConfiguration.YFilter
    cvGeneralConfiguration.EntityData.YangName = "cvGeneralConfiguration"
    cvGeneralConfiguration.EntityData.BundleName = "cisco_ios_xe"
    cvGeneralConfiguration.EntityData.ParentYangName = "CISCO-VOICE-DIAL-CONTROL-MIB"
    cvGeneralConfiguration.EntityData.SegmentPath = "cvGeneralConfiguration"
    cvGeneralConfiguration.EntityData.AbsolutePath = "CISCO-VOICE-DIAL-CONTROL-MIB:CISCO-VOICE-DIAL-CONTROL-MIB/" + cvGeneralConfiguration.EntityData.SegmentPath
    cvGeneralConfiguration.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvGeneralConfiguration.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvGeneralConfiguration.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvGeneralConfiguration.EntityData.Children = types.NewOrderedMap()
    cvGeneralConfiguration.EntityData.Leafs = types.NewOrderedMap()
    cvGeneralConfiguration.EntityData.Leafs.Append("cvGeneralPoorQoVNotificationEnable", types.YLeaf{"CvGeneralPoorQoVNotificationEnable", cvGeneralConfiguration.CvGeneralPoorQoVNotificationEnable})
    cvGeneralConfiguration.EntityData.Leafs.Append("cvGeneralFallbackNotificationEnable", types.YLeaf{"CvGeneralFallbackNotificationEnable", cvGeneralConfiguration.CvGeneralFallbackNotificationEnable})
    cvGeneralConfiguration.EntityData.Leafs.Append("cvGeneralDSCPPolicyNotificationEnable", types.YLeaf{"CvGeneralDSCPPolicyNotificationEnable", cvGeneralConfiguration.CvGeneralDSCPPolicyNotificationEnable})
    cvGeneralConfiguration.EntityData.Leafs.Append("cvGeneralMediaPolicyNotificationEnable", types.YLeaf{"CvGeneralMediaPolicyNotificationEnable", cvGeneralConfiguration.CvGeneralMediaPolicyNotificationEnable})

    cvGeneralConfiguration.EntityData.YListKeys = []string {}

    return &(cvGeneralConfiguration.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_CvGatewayCallActive
type CISCOVOICEDIALCONTROLMIB_CvGatewayCallActive struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The current number of DS0 interfaces used for the active calls. The type is
    // interface{} with range: 0..4294967295. Units are interfaces.
    CvCallActiveDS0s interface{}

    // A high threshold used to determine when to generate the
    // cvdcActiveDS0sHighNotification. This object  represents the percentage of
    // active DS0s in total number  of DS0s. The type is interface{} with range:
    // 0..100. Units are percent.
    CvCallActiveDS0sHighThreshold interface{}

    // A low threshold used to determine when to generate the
    // cvdcActiveDS0sLowNotification notification. This object  represents the
    // percentage of active DS0s in total number  of DS0s. The type is interface{}
    // with range: 0..100. Units are percent.
    CvCallActiveDS0sLowThreshold interface{}

    // Specifies whether or not cvdcActiveDS0sHighNotification should be
    // generated.  'true' : Indicates that the cvdcActiveDS0sHighNotification     
    // generation is enabled.  'false': Indicates that
    // cvdcActiveDS0sHighNotification          generation is disabled. The type is
    // bool.
    CvCallActiveDS0sHighNotifyEnable interface{}

    // Specifies whether or not cvdcActiveDS0sLowNotification should be generated.
    // 'true' : Indicates that the cvdcActiveDS0sLowNotification         
    // generation is enabled.  'false': Indicates that
    // cvdcActiveDS0sLowNotification          generation is disabled. The type is
    // bool.
    CvCallActiveDS0sLowNotifyEnable interface{}
}

func (cvGatewayCallActive *CISCOVOICEDIALCONTROLMIB_CvGatewayCallActive) GetEntityData() *types.CommonEntityData {
    cvGatewayCallActive.EntityData.YFilter = cvGatewayCallActive.YFilter
    cvGatewayCallActive.EntityData.YangName = "cvGatewayCallActive"
    cvGatewayCallActive.EntityData.BundleName = "cisco_ios_xe"
    cvGatewayCallActive.EntityData.ParentYangName = "CISCO-VOICE-DIAL-CONTROL-MIB"
    cvGatewayCallActive.EntityData.SegmentPath = "cvGatewayCallActive"
    cvGatewayCallActive.EntityData.AbsolutePath = "CISCO-VOICE-DIAL-CONTROL-MIB:CISCO-VOICE-DIAL-CONTROL-MIB/" + cvGatewayCallActive.EntityData.SegmentPath
    cvGatewayCallActive.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvGatewayCallActive.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvGatewayCallActive.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvGatewayCallActive.EntityData.Children = types.NewOrderedMap()
    cvGatewayCallActive.EntityData.Leafs = types.NewOrderedMap()
    cvGatewayCallActive.EntityData.Leafs.Append("cvCallActiveDS0s", types.YLeaf{"CvCallActiveDS0s", cvGatewayCallActive.CvCallActiveDS0s})
    cvGatewayCallActive.EntityData.Leafs.Append("cvCallActiveDS0sHighThreshold", types.YLeaf{"CvCallActiveDS0sHighThreshold", cvGatewayCallActive.CvCallActiveDS0sHighThreshold})
    cvGatewayCallActive.EntityData.Leafs.Append("cvCallActiveDS0sLowThreshold", types.YLeaf{"CvCallActiveDS0sLowThreshold", cvGatewayCallActive.CvCallActiveDS0sLowThreshold})
    cvGatewayCallActive.EntityData.Leafs.Append("cvCallActiveDS0sHighNotifyEnable", types.YLeaf{"CvCallActiveDS0sHighNotifyEnable", cvGatewayCallActive.CvCallActiveDS0sHighNotifyEnable})
    cvGatewayCallActive.EntityData.Leafs.Append("cvCallActiveDS0sLowNotifyEnable", types.YLeaf{"CvCallActiveDS0sLowNotifyEnable", cvGatewayCallActive.CvCallActiveDS0sLowNotifyEnable})

    cvGatewayCallActive.EntityData.YListKeys = []string {}

    return &(cvGatewayCallActive.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_CvCallVolume
type CISCOVOICEDIALCONTROLMIB_CvCallVolume struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object represents the total number of active call legs in the voice
    // gateway. The type is interface{} with range: 0..65535.
    CvCallVolConnTotalActiveConnections interface{}

    // This object represents the licensed call capacity for a voice gateway.  If
    // the value is 0, no  licensing is done and the gateway can be  accomodate as
    // many calls depending on its capability. The type is interface{} with range:
    // 0..65535.
    CvCallVolConnMaxCallConnectionLicenese interface{}
}

func (cvCallVolume *CISCOVOICEDIALCONTROLMIB_CvCallVolume) GetEntityData() *types.CommonEntityData {
    cvCallVolume.EntityData.YFilter = cvCallVolume.YFilter
    cvCallVolume.EntityData.YangName = "cvCallVolume"
    cvCallVolume.EntityData.BundleName = "cisco_ios_xe"
    cvCallVolume.EntityData.ParentYangName = "CISCO-VOICE-DIAL-CONTROL-MIB"
    cvCallVolume.EntityData.SegmentPath = "cvCallVolume"
    cvCallVolume.EntityData.AbsolutePath = "CISCO-VOICE-DIAL-CONTROL-MIB:CISCO-VOICE-DIAL-CONTROL-MIB/" + cvCallVolume.EntityData.SegmentPath
    cvCallVolume.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvCallVolume.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvCallVolume.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvCallVolume.EntityData.Children = types.NewOrderedMap()
    cvCallVolume.EntityData.Leafs = types.NewOrderedMap()
    cvCallVolume.EntityData.Leafs.Append("cvCallVolConnTotalActiveConnections", types.YLeaf{"CvCallVolConnTotalActiveConnections", cvCallVolume.CvCallVolConnTotalActiveConnections})
    cvCallVolume.EntityData.Leafs.Append("cvCallVolConnMaxCallConnectionLicenese", types.YLeaf{"CvCallVolConnMaxCallConnectionLicenese", cvCallVolume.CvCallVolConnMaxCallConnectionLicenese})

    cvCallVolume.EntityData.YListKeys = []string {}

    return &(cvCallVolume.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_CvCallRateMonitor
type CISCOVOICEDIALCONTROLMIB_CvCallRateMonitor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object represents the state of call-monitoring. A value of 'true'
    // indicates that call-monitoring  is enabled.  A value of 'false' indicates
    // that  call-monitoring is disabled. The type is bool.
    CvCallRateMonitorEnable interface{}

    // This object represents the interval for which the gateway monitors the
    // call-rate. The type is interface{} with range: 1..12. Units are five
    // seconds.
    CvCallRateMonitorTime interface{}

    // This object represents the total number of calls handled by the gateway
    // during the  monitored time. The type is interface{} with range: 0..65535.
    CvCallRate interface{}

    // This object represents the high water mark for the number of calls handled
    // by the  gateway in an unit interval of  cvCallRateMonitorTime, from the
    // time  the call-monitoring is enabled. The type is interface{} with range:
    // 0..65535.
    CvCallRateHiWaterMark interface{}
}

func (cvCallRateMonitor *CISCOVOICEDIALCONTROLMIB_CvCallRateMonitor) GetEntityData() *types.CommonEntityData {
    cvCallRateMonitor.EntityData.YFilter = cvCallRateMonitor.YFilter
    cvCallRateMonitor.EntityData.YangName = "cvCallRateMonitor"
    cvCallRateMonitor.EntityData.BundleName = "cisco_ios_xe"
    cvCallRateMonitor.EntityData.ParentYangName = "CISCO-VOICE-DIAL-CONTROL-MIB"
    cvCallRateMonitor.EntityData.SegmentPath = "cvCallRateMonitor"
    cvCallRateMonitor.EntityData.AbsolutePath = "CISCO-VOICE-DIAL-CONTROL-MIB:CISCO-VOICE-DIAL-CONTROL-MIB/" + cvCallRateMonitor.EntityData.SegmentPath
    cvCallRateMonitor.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvCallRateMonitor.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvCallRateMonitor.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvCallRateMonitor.EntityData.Children = types.NewOrderedMap()
    cvCallRateMonitor.EntityData.Leafs = types.NewOrderedMap()
    cvCallRateMonitor.EntityData.Leafs.Append("cvCallRateMonitorEnable", types.YLeaf{"CvCallRateMonitorEnable", cvCallRateMonitor.CvCallRateMonitorEnable})
    cvCallRateMonitor.EntityData.Leafs.Append("cvCallRateMonitorTime", types.YLeaf{"CvCallRateMonitorTime", cvCallRateMonitor.CvCallRateMonitorTime})
    cvCallRateMonitor.EntityData.Leafs.Append("cvCallRate", types.YLeaf{"CvCallRate", cvCallRateMonitor.CvCallRate})
    cvCallRateMonitor.EntityData.Leafs.Append("cvCallRateHiWaterMark", types.YLeaf{"CvCallRateHiWaterMark", cvCallRateMonitor.CvCallRateHiWaterMark})

    cvCallRateMonitor.EntityData.YListKeys = []string {}

    return &(cvCallRateMonitor.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_CvCallVolumeStatsHistory
type CISCOVOICEDIALCONTROLMIB_CvCallVolumeStatsHistory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This Object specifies the thresold duration in seconds.
    // cvCallDurationStatsTable will monitor all the calls below this  threshold. 
    // Decresing the value of the threshold will reset this table. The type is
    // interface{} with range: 1..3600. Units are seconds.
    CvCallDurationStatsThreshold interface{}

    // This Object specifies the number of entries the watermark table will
    // maintain.  This value will decide the number of elements in
    // cvCallRateWMTable, cvCallLegRateWMTable, cvActiveCallWMTable and
    // cvSipMsgRateWMTable. The type is interface{} with range: 3..10.
    CvCallVolumeWMTableSize interface{}
}

func (cvCallVolumeStatsHistory *CISCOVOICEDIALCONTROLMIB_CvCallVolumeStatsHistory) GetEntityData() *types.CommonEntityData {
    cvCallVolumeStatsHistory.EntityData.YFilter = cvCallVolumeStatsHistory.YFilter
    cvCallVolumeStatsHistory.EntityData.YangName = "cvCallVolumeStatsHistory"
    cvCallVolumeStatsHistory.EntityData.BundleName = "cisco_ios_xe"
    cvCallVolumeStatsHistory.EntityData.ParentYangName = "CISCO-VOICE-DIAL-CONTROL-MIB"
    cvCallVolumeStatsHistory.EntityData.SegmentPath = "cvCallVolumeStatsHistory"
    cvCallVolumeStatsHistory.EntityData.AbsolutePath = "CISCO-VOICE-DIAL-CONTROL-MIB:CISCO-VOICE-DIAL-CONTROL-MIB/" + cvCallVolumeStatsHistory.EntityData.SegmentPath
    cvCallVolumeStatsHistory.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvCallVolumeStatsHistory.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvCallVolumeStatsHistory.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvCallVolumeStatsHistory.EntityData.Children = types.NewOrderedMap()
    cvCallVolumeStatsHistory.EntityData.Leafs = types.NewOrderedMap()
    cvCallVolumeStatsHistory.EntityData.Leafs.Append("cvCallDurationStatsThreshold", types.YLeaf{"CvCallDurationStatsThreshold", cvCallVolumeStatsHistory.CvCallDurationStatsThreshold})
    cvCallVolumeStatsHistory.EntityData.Leafs.Append("cvCallVolumeWMTableSize", types.YLeaf{"CvCallVolumeWMTableSize", cvCallVolumeStatsHistory.CvCallVolumeWMTableSize})

    cvCallVolumeStatsHistory.EntityData.YListKeys = []string {}

    return &(cvCallVolumeStatsHistory.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_CvPeerCfgTable
// The table contains the Voice Generic Peer information that
// is used to create an ifIndexed row with an appropriate
// ifType that is associated with the cvPeerCfgType and
// cvPeerCfgPeerType objects.
type CISCOVOICEDIALCONTROLMIB_CvPeerCfgTable struct {
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
    // of CISCOVOICEDIALCONTROLMIB_CvPeerCfgTable_CvPeerCfgEntry.
    CvPeerCfgEntry []*CISCOVOICEDIALCONTROLMIB_CvPeerCfgTable_CvPeerCfgEntry
}

func (cvPeerCfgTable *CISCOVOICEDIALCONTROLMIB_CvPeerCfgTable) GetEntityData() *types.CommonEntityData {
    cvPeerCfgTable.EntityData.YFilter = cvPeerCfgTable.YFilter
    cvPeerCfgTable.EntityData.YangName = "cvPeerCfgTable"
    cvPeerCfgTable.EntityData.BundleName = "cisco_ios_xe"
    cvPeerCfgTable.EntityData.ParentYangName = "CISCO-VOICE-DIAL-CONTROL-MIB"
    cvPeerCfgTable.EntityData.SegmentPath = "cvPeerCfgTable"
    cvPeerCfgTable.EntityData.AbsolutePath = "CISCO-VOICE-DIAL-CONTROL-MIB:CISCO-VOICE-DIAL-CONTROL-MIB/" + cvPeerCfgTable.EntityData.SegmentPath
    cvPeerCfgTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvPeerCfgTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvPeerCfgTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvPeerCfgTable.EntityData.Children = types.NewOrderedMap()
    cvPeerCfgTable.EntityData.Children.Append("cvPeerCfgEntry", types.YChild{"CvPeerCfgEntry", nil})
    for i := range cvPeerCfgTable.CvPeerCfgEntry {
        cvPeerCfgTable.EntityData.Children.Append(types.GetSegmentPath(cvPeerCfgTable.CvPeerCfgEntry[i]), types.YChild{"CvPeerCfgEntry", cvPeerCfgTable.CvPeerCfgEntry[i]})
    }
    cvPeerCfgTable.EntityData.Leafs = types.NewOrderedMap()

    cvPeerCfgTable.EntityData.YListKeys = []string {}

    return &(cvPeerCfgTable.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_CvPeerCfgTable_CvPeerCfgEntry
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
type CISCOVOICEDIALCONTROLMIB_CvPeerCfgTable_CvPeerCfgEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. An arbitrary index that uniquely identifies a
    // generic voice peer. The type is interface{} with range: 1..2147483647.
    CvPeerCfgIndex interface{}

    // The ifIndex of the peer associated ifEntry. The ifIndex appears after the
    // associated ifEntry is created successfully. This ifIndex will be used to
    // access the objects in the Voice over Telephony or Voice over IP peer
    // specific table. In addition, the ifIndex is also used to access the
    // associated peer configuration entry of the IETF Dial Control MIB. If the
    // peer associated ifEntry had not been created, then this object has a value
    // of zero. The type is interface{} with range: 0..2147483647.
    CvPeerCfgIfIndex interface{}

    // Specifies the type of voice related encapsulation. voice - voice
    // encapsulation (voiceEncap ifType) on the         telephony network. voip  -
    // VoIP encapsulation (voiceOverIp ifType) on the IP         network. mmail -
    // Media Mail over IP encapsulation (mediaMailOverIp         ifType) on the IP
    // network. voatm - VoATM encapsulation (voiceOverATM ifType) on the        
    // ATM network. vofr  - VoFR encapsulation (voiceOverFR ifType) on the        
    // Frame Relay network. The type is CvPeerCfgType.
    CvPeerCfgType interface{}

    // This object is used to create a new row or modify or delete an existing row
    // in this table. The type is RowStatus.
    CvPeerCfgRowStatus interface{}

    // Specifies the type of a peer. voice - peer in voice type to be defined in a
    // voice         gateway for voice calls.  data  - peer in data type to be
    // defined in gateway         that supports universal ports for modem/data    
    // calls and integrated ports for data calls. The type is CvPeerCfgPeerType.
    CvPeerCfgPeerType interface{}

    // This object represents the total number of active calls that has selected
    // the dialpeer as an incoming dialpeer. The type is interface{} with range:
    // 0..65535.
    CvCallVolPeerIncomingCalls interface{}

    // This object represents the total number of active calls that has selected
    // the dialpeer as an outgoing dialpeer. The type is interface{} with range:
    // 0..65535.
    CvCallVolPeerOutgoingCalls interface{}
}

func (cvPeerCfgEntry *CISCOVOICEDIALCONTROLMIB_CvPeerCfgTable_CvPeerCfgEntry) GetEntityData() *types.CommonEntityData {
    cvPeerCfgEntry.EntityData.YFilter = cvPeerCfgEntry.YFilter
    cvPeerCfgEntry.EntityData.YangName = "cvPeerCfgEntry"
    cvPeerCfgEntry.EntityData.BundleName = "cisco_ios_xe"
    cvPeerCfgEntry.EntityData.ParentYangName = "cvPeerCfgTable"
    cvPeerCfgEntry.EntityData.SegmentPath = "cvPeerCfgEntry" + types.AddKeyToken(cvPeerCfgEntry.CvPeerCfgIndex, "cvPeerCfgIndex")
    cvPeerCfgEntry.EntityData.AbsolutePath = "CISCO-VOICE-DIAL-CONTROL-MIB:CISCO-VOICE-DIAL-CONTROL-MIB/cvPeerCfgTable/" + cvPeerCfgEntry.EntityData.SegmentPath
    cvPeerCfgEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvPeerCfgEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvPeerCfgEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvPeerCfgEntry.EntityData.Children = types.NewOrderedMap()
    cvPeerCfgEntry.EntityData.Leafs = types.NewOrderedMap()
    cvPeerCfgEntry.EntityData.Leafs.Append("cvPeerCfgIndex", types.YLeaf{"CvPeerCfgIndex", cvPeerCfgEntry.CvPeerCfgIndex})
    cvPeerCfgEntry.EntityData.Leafs.Append("cvPeerCfgIfIndex", types.YLeaf{"CvPeerCfgIfIndex", cvPeerCfgEntry.CvPeerCfgIfIndex})
    cvPeerCfgEntry.EntityData.Leafs.Append("cvPeerCfgType", types.YLeaf{"CvPeerCfgType", cvPeerCfgEntry.CvPeerCfgType})
    cvPeerCfgEntry.EntityData.Leafs.Append("cvPeerCfgRowStatus", types.YLeaf{"CvPeerCfgRowStatus", cvPeerCfgEntry.CvPeerCfgRowStatus})
    cvPeerCfgEntry.EntityData.Leafs.Append("cvPeerCfgPeerType", types.YLeaf{"CvPeerCfgPeerType", cvPeerCfgEntry.CvPeerCfgPeerType})
    cvPeerCfgEntry.EntityData.Leafs.Append("cvCallVolPeerIncomingCalls", types.YLeaf{"CvCallVolPeerIncomingCalls", cvPeerCfgEntry.CvCallVolPeerIncomingCalls})
    cvPeerCfgEntry.EntityData.Leafs.Append("cvCallVolPeerOutgoingCalls", types.YLeaf{"CvCallVolPeerOutgoingCalls", cvPeerCfgEntry.CvCallVolPeerOutgoingCalls})

    cvPeerCfgEntry.EntityData.YListKeys = []string {"CvPeerCfgIndex"}

    return &(cvPeerCfgEntry.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_CvPeerCfgTable_CvPeerCfgEntry_CvPeerCfgPeerType represents         calls and integrated ports for data calls.
type CISCOVOICEDIALCONTROLMIB_CvPeerCfgTable_CvPeerCfgEntry_CvPeerCfgPeerType string

const (
    CISCOVOICEDIALCONTROLMIB_CvPeerCfgTable_CvPeerCfgEntry_CvPeerCfgPeerType_voice CISCOVOICEDIALCONTROLMIB_CvPeerCfgTable_CvPeerCfgEntry_CvPeerCfgPeerType = "voice"

    CISCOVOICEDIALCONTROLMIB_CvPeerCfgTable_CvPeerCfgEntry_CvPeerCfgPeerType_data CISCOVOICEDIALCONTROLMIB_CvPeerCfgTable_CvPeerCfgEntry_CvPeerCfgPeerType = "data"
)

// CISCOVOICEDIALCONTROLMIB_CvPeerCfgTable_CvPeerCfgEntry_CvPeerCfgType represents         Frame Relay network.
type CISCOVOICEDIALCONTROLMIB_CvPeerCfgTable_CvPeerCfgEntry_CvPeerCfgType string

const (
    CISCOVOICEDIALCONTROLMIB_CvPeerCfgTable_CvPeerCfgEntry_CvPeerCfgType_voice CISCOVOICEDIALCONTROLMIB_CvPeerCfgTable_CvPeerCfgEntry_CvPeerCfgType = "voice"

    CISCOVOICEDIALCONTROLMIB_CvPeerCfgTable_CvPeerCfgEntry_CvPeerCfgType_voip CISCOVOICEDIALCONTROLMIB_CvPeerCfgTable_CvPeerCfgEntry_CvPeerCfgType = "voip"

    CISCOVOICEDIALCONTROLMIB_CvPeerCfgTable_CvPeerCfgEntry_CvPeerCfgType_mmail CISCOVOICEDIALCONTROLMIB_CvPeerCfgTable_CvPeerCfgEntry_CvPeerCfgType = "mmail"

    CISCOVOICEDIALCONTROLMIB_CvPeerCfgTable_CvPeerCfgEntry_CvPeerCfgType_voatm CISCOVOICEDIALCONTROLMIB_CvPeerCfgTable_CvPeerCfgEntry_CvPeerCfgType = "voatm"

    CISCOVOICEDIALCONTROLMIB_CvPeerCfgTable_CvPeerCfgEntry_CvPeerCfgType_vofr CISCOVOICEDIALCONTROLMIB_CvPeerCfgTable_CvPeerCfgEntry_CvPeerCfgType = "vofr"
)

// CISCOVOICEDIALCONTROLMIB_CvVoicePeerCfgTable
// The table contains the Voice over Telephony peer specific
// information that is required to accept voice calls or to
// which it will place them or perform various loopback tests
// via interface.
type CISCOVOICEDIALCONTROLMIB_CvVoicePeerCfgTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A single Voice specific Peer. One entry per voice encapsulation. The entry
    // is created when its associated 'voiceEncap(103)' encapsulation ifEntry is
    // created. This entry is deleted when its associated ifEntry is deleted. The
    // type is slice of
    // CISCOVOICEDIALCONTROLMIB_CvVoicePeerCfgTable_CvVoicePeerCfgEntry.
    CvVoicePeerCfgEntry []*CISCOVOICEDIALCONTROLMIB_CvVoicePeerCfgTable_CvVoicePeerCfgEntry
}

func (cvVoicePeerCfgTable *CISCOVOICEDIALCONTROLMIB_CvVoicePeerCfgTable) GetEntityData() *types.CommonEntityData {
    cvVoicePeerCfgTable.EntityData.YFilter = cvVoicePeerCfgTable.YFilter
    cvVoicePeerCfgTable.EntityData.YangName = "cvVoicePeerCfgTable"
    cvVoicePeerCfgTable.EntityData.BundleName = "cisco_ios_xe"
    cvVoicePeerCfgTable.EntityData.ParentYangName = "CISCO-VOICE-DIAL-CONTROL-MIB"
    cvVoicePeerCfgTable.EntityData.SegmentPath = "cvVoicePeerCfgTable"
    cvVoicePeerCfgTable.EntityData.AbsolutePath = "CISCO-VOICE-DIAL-CONTROL-MIB:CISCO-VOICE-DIAL-CONTROL-MIB/" + cvVoicePeerCfgTable.EntityData.SegmentPath
    cvVoicePeerCfgTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvVoicePeerCfgTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvVoicePeerCfgTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvVoicePeerCfgTable.EntityData.Children = types.NewOrderedMap()
    cvVoicePeerCfgTable.EntityData.Children.Append("cvVoicePeerCfgEntry", types.YChild{"CvVoicePeerCfgEntry", nil})
    for i := range cvVoicePeerCfgTable.CvVoicePeerCfgEntry {
        cvVoicePeerCfgTable.EntityData.Children.Append(types.GetSegmentPath(cvVoicePeerCfgTable.CvVoicePeerCfgEntry[i]), types.YChild{"CvVoicePeerCfgEntry", cvVoicePeerCfgTable.CvVoicePeerCfgEntry[i]})
    }
    cvVoicePeerCfgTable.EntityData.Leafs = types.NewOrderedMap()

    cvVoicePeerCfgTable.EntityData.YListKeys = []string {}

    return &(cvVoicePeerCfgTable.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_CvVoicePeerCfgTable_CvVoicePeerCfgEntry
// A single Voice specific Peer. One entry per voice
// encapsulation.
// The entry is created when its associated 'voiceEncap(103)'
// encapsulation ifEntry is created.
// This entry is deleted when its associated ifEntry is
// deleted.
type CISCOVOICEDIALCONTROLMIB_CvVoicePeerCfgTable_CvVoicePeerCfgEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

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
    CvVoicePeerCfgSessionTarget interface{}

    // The object specifies the prefix of the dial digits for the peer. The dial
    // digits prefix is sent to telephony interface before the real phone number
    // when the system places the outgoing call to the voice encapsulation peer
    // over telephony interface. The type is string with length: 0..32.
    CvVoicePeerCfgDialDigitsPrefix interface{}

    // The object enables/disables the DID call treatment for incoming DNIS
    // digits. true  - treat the incoming DNIS digits as if the digits         are
    // received from DID trunk. false - Disable DID call treatment for incoming
    // DNIS         digits. The type is bool.
    CvVoicePeerCfgDIDCallEnable interface{}

    // The object specifies the CAS group number of a CAS capable T1/E1  that is
    // in the dialCtlPeerCfgLowerIf object of RFC2128. This object can be set to a
    // valid CAS group number only if the dialCtlPeerCfgLowerIf contains a valid
    // ifIndex for a CAS capable T1/E1. The object must set to -1 before the value
    // of the  dialCtlPeerCfgLowerIf object of RFC2128 can be changed. The type is
    // interface{} with range: -1..30.
    CvVoicePeerCfgCasGroup interface{}

    // This object specifies that the E.164 number specified in the
    // dialCtlPeerCfgOriginateAddress field of the associated dialCtlPeerCfgTable
    // entry be registered as an extension  phone number of this gateway for H323
    // gatekeeper and/or  SIP registrar. The type is bool.
    CvVoicePeerCfgRegisterE164 interface{}

    // This object specifies the number of dialed digits to forward to the remote
    // destination in the setup message. The object can take the value:     0-32
    // number of right justified digits to forward     -1 default     -2 forward
    // extra digits i.e those over and above        those needed to match to the
    // destination pattern     -3 forward all digits. The type is interface{} with
    // range: -3..32.
    CvVoicePeerCfgForwardDigits interface{}

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
    // CvVoicePeerCfgEchoCancellerTest.
    CvVoicePeerCfgEchoCancellerTest interface{}
}

func (cvVoicePeerCfgEntry *CISCOVOICEDIALCONTROLMIB_CvVoicePeerCfgTable_CvVoicePeerCfgEntry) GetEntityData() *types.CommonEntityData {
    cvVoicePeerCfgEntry.EntityData.YFilter = cvVoicePeerCfgEntry.YFilter
    cvVoicePeerCfgEntry.EntityData.YangName = "cvVoicePeerCfgEntry"
    cvVoicePeerCfgEntry.EntityData.BundleName = "cisco_ios_xe"
    cvVoicePeerCfgEntry.EntityData.ParentYangName = "cvVoicePeerCfgTable"
    cvVoicePeerCfgEntry.EntityData.SegmentPath = "cvVoicePeerCfgEntry" + types.AddKeyToken(cvVoicePeerCfgEntry.IfIndex, "ifIndex")
    cvVoicePeerCfgEntry.EntityData.AbsolutePath = "CISCO-VOICE-DIAL-CONTROL-MIB:CISCO-VOICE-DIAL-CONTROL-MIB/cvVoicePeerCfgTable/" + cvVoicePeerCfgEntry.EntityData.SegmentPath
    cvVoicePeerCfgEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvVoicePeerCfgEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvVoicePeerCfgEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvVoicePeerCfgEntry.EntityData.Children = types.NewOrderedMap()
    cvVoicePeerCfgEntry.EntityData.Leafs = types.NewOrderedMap()
    cvVoicePeerCfgEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", cvVoicePeerCfgEntry.IfIndex})
    cvVoicePeerCfgEntry.EntityData.Leafs.Append("cvVoicePeerCfgSessionTarget", types.YLeaf{"CvVoicePeerCfgSessionTarget", cvVoicePeerCfgEntry.CvVoicePeerCfgSessionTarget})
    cvVoicePeerCfgEntry.EntityData.Leafs.Append("cvVoicePeerCfgDialDigitsPrefix", types.YLeaf{"CvVoicePeerCfgDialDigitsPrefix", cvVoicePeerCfgEntry.CvVoicePeerCfgDialDigitsPrefix})
    cvVoicePeerCfgEntry.EntityData.Leafs.Append("cvVoicePeerCfgDIDCallEnable", types.YLeaf{"CvVoicePeerCfgDIDCallEnable", cvVoicePeerCfgEntry.CvVoicePeerCfgDIDCallEnable})
    cvVoicePeerCfgEntry.EntityData.Leafs.Append("cvVoicePeerCfgCasGroup", types.YLeaf{"CvVoicePeerCfgCasGroup", cvVoicePeerCfgEntry.CvVoicePeerCfgCasGroup})
    cvVoicePeerCfgEntry.EntityData.Leafs.Append("cvVoicePeerCfgRegisterE164", types.YLeaf{"CvVoicePeerCfgRegisterE164", cvVoicePeerCfgEntry.CvVoicePeerCfgRegisterE164})
    cvVoicePeerCfgEntry.EntityData.Leafs.Append("cvVoicePeerCfgForwardDigits", types.YLeaf{"CvVoicePeerCfgForwardDigits", cvVoicePeerCfgEntry.CvVoicePeerCfgForwardDigits})
    cvVoicePeerCfgEntry.EntityData.Leafs.Append("cvVoicePeerCfgEchoCancellerTest", types.YLeaf{"CvVoicePeerCfgEchoCancellerTest", cvVoicePeerCfgEntry.CvVoicePeerCfgEchoCancellerTest})

    cvVoicePeerCfgEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(cvVoicePeerCfgEntry.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_CvVoicePeerCfgTable_CvVoicePeerCfgEntry_CvVoicePeerCfgEchoCancellerTest represents echoCancellerG168Test9   - run ITU-T G.168 Test 9.
type CISCOVOICEDIALCONTROLMIB_CvVoicePeerCfgTable_CvVoicePeerCfgEntry_CvVoicePeerCfgEchoCancellerTest string

const (
    CISCOVOICEDIALCONTROLMIB_CvVoicePeerCfgTable_CvVoicePeerCfgEntry_CvVoicePeerCfgEchoCancellerTest_echoCancellerTestNone CISCOVOICEDIALCONTROLMIB_CvVoicePeerCfgTable_CvVoicePeerCfgEntry_CvVoicePeerCfgEchoCancellerTest = "echoCancellerTestNone"

    CISCOVOICEDIALCONTROLMIB_CvVoicePeerCfgTable_CvVoicePeerCfgEntry_CvVoicePeerCfgEchoCancellerTest_echoCancellerG168Test2A CISCOVOICEDIALCONTROLMIB_CvVoicePeerCfgTable_CvVoicePeerCfgEntry_CvVoicePeerCfgEchoCancellerTest = "echoCancellerG168Test2A"

    CISCOVOICEDIALCONTROLMIB_CvVoicePeerCfgTable_CvVoicePeerCfgEntry_CvVoicePeerCfgEchoCancellerTest_echoCancellerG168Test2B CISCOVOICEDIALCONTROLMIB_CvVoicePeerCfgTable_CvVoicePeerCfgEntry_CvVoicePeerCfgEchoCancellerTest = "echoCancellerG168Test2B"

    CISCOVOICEDIALCONTROLMIB_CvVoicePeerCfgTable_CvVoicePeerCfgEntry_CvVoicePeerCfgEchoCancellerTest_echoCancellerG168Test2Ca CISCOVOICEDIALCONTROLMIB_CvVoicePeerCfgTable_CvVoicePeerCfgEntry_CvVoicePeerCfgEchoCancellerTest = "echoCancellerG168Test2Ca"

    CISCOVOICEDIALCONTROLMIB_CvVoicePeerCfgTable_CvVoicePeerCfgEntry_CvVoicePeerCfgEchoCancellerTest_echoCancellerG168Test2Cb CISCOVOICEDIALCONTROLMIB_CvVoicePeerCfgTable_CvVoicePeerCfgEntry_CvVoicePeerCfgEchoCancellerTest = "echoCancellerG168Test2Cb"

    CISCOVOICEDIALCONTROLMIB_CvVoicePeerCfgTable_CvVoicePeerCfgEntry_CvVoicePeerCfgEchoCancellerTest_echoCancellerG168Test3A CISCOVOICEDIALCONTROLMIB_CvVoicePeerCfgTable_CvVoicePeerCfgEntry_CvVoicePeerCfgEchoCancellerTest = "echoCancellerG168Test3A"

    CISCOVOICEDIALCONTROLMIB_CvVoicePeerCfgTable_CvVoicePeerCfgEntry_CvVoicePeerCfgEchoCancellerTest_echoCancellerG168Test3B CISCOVOICEDIALCONTROLMIB_CvVoicePeerCfgTable_CvVoicePeerCfgEntry_CvVoicePeerCfgEchoCancellerTest = "echoCancellerG168Test3B"

    CISCOVOICEDIALCONTROLMIB_CvVoicePeerCfgTable_CvVoicePeerCfgEntry_CvVoicePeerCfgEchoCancellerTest_echoCancellerG168Test3C CISCOVOICEDIALCONTROLMIB_CvVoicePeerCfgTable_CvVoicePeerCfgEntry_CvVoicePeerCfgEchoCancellerTest = "echoCancellerG168Test3C"

    CISCOVOICEDIALCONTROLMIB_CvVoicePeerCfgTable_CvVoicePeerCfgEntry_CvVoicePeerCfgEchoCancellerTest_echoCancellerG168Test4 CISCOVOICEDIALCONTROLMIB_CvVoicePeerCfgTable_CvVoicePeerCfgEntry_CvVoicePeerCfgEchoCancellerTest = "echoCancellerG168Test4"

    CISCOVOICEDIALCONTROLMIB_CvVoicePeerCfgTable_CvVoicePeerCfgEntry_CvVoicePeerCfgEchoCancellerTest_echoCancellerG168Test6 CISCOVOICEDIALCONTROLMIB_CvVoicePeerCfgTable_CvVoicePeerCfgEntry_CvVoicePeerCfgEchoCancellerTest = "echoCancellerG168Test6"

    CISCOVOICEDIALCONTROLMIB_CvVoicePeerCfgTable_CvVoicePeerCfgEntry_CvVoicePeerCfgEchoCancellerTest_echoCancellerG168Test9 CISCOVOICEDIALCONTROLMIB_CvVoicePeerCfgTable_CvVoicePeerCfgEntry_CvVoicePeerCfgEchoCancellerTest = "echoCancellerG168Test9"

    CISCOVOICEDIALCONTROLMIB_CvVoicePeerCfgTable_CvVoicePeerCfgEntry_CvVoicePeerCfgEchoCancellerTest_echoCancellerG168Test5 CISCOVOICEDIALCONTROLMIB_CvVoicePeerCfgTable_CvVoicePeerCfgEntry_CvVoicePeerCfgEchoCancellerTest = "echoCancellerG168Test5"

    CISCOVOICEDIALCONTROLMIB_CvVoicePeerCfgTable_CvVoicePeerCfgEntry_CvVoicePeerCfgEchoCancellerTest_echoCancellerG168Test7 CISCOVOICEDIALCONTROLMIB_CvVoicePeerCfgTable_CvVoicePeerCfgEntry_CvVoicePeerCfgEchoCancellerTest = "echoCancellerG168Test7"
)

// CISCOVOICEDIALCONTROLMIB_CvVoIPPeerCfgTable
// The table contains the Voice over IP (VoIP) peer specific
// information that is required to accept voice calls or to
// which it will place them via IP backbone with the
// specified session protocol in cvVoIPPeerCfgSessionProtocol.
type CISCOVOICEDIALCONTROLMIB_CvVoIPPeerCfgTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A single VoIP specific Peer. One entry per VoIP encapsulation. The entry is
    // created when its associated 'voiceOverIp(104)' encapsulation ifEntry is
    // created. This entry is deleted when its associated ifEntry is deleted. The
    // type is slice of
    // CISCOVOICEDIALCONTROLMIB_CvVoIPPeerCfgTable_CvVoIPPeerCfgEntry.
    CvVoIPPeerCfgEntry []*CISCOVOICEDIALCONTROLMIB_CvVoIPPeerCfgTable_CvVoIPPeerCfgEntry
}

func (cvVoIPPeerCfgTable *CISCOVOICEDIALCONTROLMIB_CvVoIPPeerCfgTable) GetEntityData() *types.CommonEntityData {
    cvVoIPPeerCfgTable.EntityData.YFilter = cvVoIPPeerCfgTable.YFilter
    cvVoIPPeerCfgTable.EntityData.YangName = "cvVoIPPeerCfgTable"
    cvVoIPPeerCfgTable.EntityData.BundleName = "cisco_ios_xe"
    cvVoIPPeerCfgTable.EntityData.ParentYangName = "CISCO-VOICE-DIAL-CONTROL-MIB"
    cvVoIPPeerCfgTable.EntityData.SegmentPath = "cvVoIPPeerCfgTable"
    cvVoIPPeerCfgTable.EntityData.AbsolutePath = "CISCO-VOICE-DIAL-CONTROL-MIB:CISCO-VOICE-DIAL-CONTROL-MIB/" + cvVoIPPeerCfgTable.EntityData.SegmentPath
    cvVoIPPeerCfgTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvVoIPPeerCfgTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvVoIPPeerCfgTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvVoIPPeerCfgTable.EntityData.Children = types.NewOrderedMap()
    cvVoIPPeerCfgTable.EntityData.Children.Append("cvVoIPPeerCfgEntry", types.YChild{"CvVoIPPeerCfgEntry", nil})
    for i := range cvVoIPPeerCfgTable.CvVoIPPeerCfgEntry {
        cvVoIPPeerCfgTable.EntityData.Children.Append(types.GetSegmentPath(cvVoIPPeerCfgTable.CvVoIPPeerCfgEntry[i]), types.YChild{"CvVoIPPeerCfgEntry", cvVoIPPeerCfgTable.CvVoIPPeerCfgEntry[i]})
    }
    cvVoIPPeerCfgTable.EntityData.Leafs = types.NewOrderedMap()

    cvVoIPPeerCfgTable.EntityData.YListKeys = []string {}

    return &(cvVoIPPeerCfgTable.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_CvVoIPPeerCfgTable_CvVoIPPeerCfgEntry
// A single VoIP specific Peer. One entry per VoIP
// encapsulation.
// The entry is created when its associated 'voiceOverIp(104)'
// encapsulation ifEntry is created.
// This entry is deleted when its associated ifEntry is
// deleted.
type CISCOVOICEDIALCONTROLMIB_CvVoIPPeerCfgTable_CvVoIPPeerCfgEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // The object specifies the session protocol to be used for Internet call
    // between local and remote router via IP backbone. The type is
    // CvSessionProtocol.
    CvVoIPPeerCfgSessionProtocol interface{}

    // The object specifies the user requested Quality of Service for the call.
    // The type is QosService.
    CvVoIPPeerCfgDesiredQoS interface{}

    // The object specifies the minimally acceptable Quality of Service for the
    // call. The type is QosService.
    CvVoIPPeerCfgMinAcceptableQoS interface{}

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
    CvVoIPPeerCfgSessionTarget interface{}

    // This object specifies the most desirable codec of speech for the VoIP peer.
    // The type is CvcSpeechCoderRate.
    CvVoIPPeerCfgCoderRate interface{}

    // This object specifies the default transmit rate of FAX the VoIP peer. If
    // the value of this object is 'none', then the Fax relay feature is disabled;
    // otherwise the Fax relay feature is enabled. The type is CvcFaxTransmitRate.
    CvVoIPPeerCfgFaxRate interface{}

    // This object specifies whether or not the VAD (Voice Activity Detection)
    // voice data is continuously transmitted to IP backbone. The type is bool.
    CvVoIPPeerCfgVADEnable interface{}

    // This object specifies the user requested Expectation Factor of voice
    // quality for the call via this peer. The type is interface{} with range:
    // 0..20. Units are equipment impairment factor (eif).
    CvVoIPPeerCfgExpectFactor interface{}

    // This object specifies the user requested Calculated Planning Impairment
    // Factor (Icpif) for the call via this peer. The type is interface{} with
    // range: 0..55. Units are equipment impairment factor (eif).
    CvVoIPPeerCfgIcpif interface{}

    // This object specifies whether cvdcPoorQoVNotification (or the newer
    // cvdcPoorQoVNotificationRev1) traps should be generated for the call that is
    // associated with this peer. The type is bool.
    CvVoIPPeerCfgPoorQoVNotificationEnable interface{}

    // This object specifies whether the outgoing voice related UDP packet
    // contains a valid checksum value. true  - enable the checksum of outgoing
    // voice UDP packets false - disable the checksum of outgoing voice UDP
    // packets. The type is bool.
    CvVoIPPeerCfgUDPChecksumEnable interface{}

    // This object specifies the value to be stored in the IP Precedence field of
    // voice packets, with values ranging from 0 (normal precedence) through 7
    // (network control), allowing the managed system to set the importance of
    // each voice packet for delivering them to the destination peer. The type is
    // interface{} with range: 0..7.
    CvVoIPPeerCfgIPPrecedence interface{}

    // This object specifies the technology prefix of the peer, The technology
    // prefix and the called party address are passed in Admission Request (ARQ)
    // to gatekeeper for the called party address resolution during call setup.
    // The type is string with length: 0..128.
    CvVoIPPeerCfgTechPrefix interface{}

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
    CvVoIPPeerCfgDigitRelay interface{}

    // This object specifies the size of the voice payload sample to be produced
    // by the coder specified in cvVoIPPeerCfgCoderRate. Each coder sample
    // produces 10 bytes of voice payload. The specified value will be rounded
    // down to the nearest valid size.  A value of 0, specifies that the coder
    // defined by cvVoIPPeerCfgCoderRate should produce its default payload size.
    // The type is interface{} with range: 0..None | 10..240. Units are bytes.
    CvVoIPPeerCfgCoderBytes interface{}

    // This object specifies the payload size to be produced by the coder when it
    // is generating fax data. A value of 0, specifies that the coder specified in
    // cvVoIPCfgPeerCoderRate should produce its default fax payload size. The
    // type is interface{} with range: 0..None | 10..255. Units are bytes.
    CvVoIPPeerCfgFaxBytes interface{}

    // This object specifies the type of in-band signaling that will be used
    // between the end points of the call. It is used by the router to determine
    // how to interpret ABCD signaling bits sent as part of voice payload data.
    // The type is CvcInBandSignaling.
    CvVoIPPeerCfgInBandSignaling interface{}

    // This object specifies how the media is to be setup on an IP-IP Gateway. Two
    // choices are valid: flow-through and flow-around. When in flow-through mode,
    // which is the default setting, the IP-IP Gateway will terminate and  then
    // re-originate the media stream. When flow-around is configured the Gateway
    // will not be involved with the media, since it will flow-around the Gateway
    // and will be established directly between the endpoints. The type is
    // CvVoIPPeerCfgMediaSetting.
    CvVoIPPeerCfgMediaSetting interface{}

    // The object specifies the user requested Quality of Service for the video
    // portion of the call. The type is QosService.
    CvVoIPPeerCfgDesiredQoSVideo interface{}

    // The object specifies the minimally acceptable Quality of Service for the
    // video portion of the call. The type is QosService.
    CvVoIPPeerCfgMinAcceptableQoSVideo interface{}

    // This object specifies the Inbound VoIP calls that match an outbound VoIP
    // dialpeer will be responded with a SIP  redirect(for inbound SIP) or H.450.3
    // call-forward(for  inbound H.323). The type is bool.
    CvVoIPPeerCfgRedirectip2ip interface{}

    // If the object has a value true(1) octet align operation is used, and if the
    // value is false(2), bandwidth efficient operation is used. This object is
    // not instantiated when the object cvVoIPPeerCfgCoderRate is not equal to
    // gsmAmrNb enum. The type is bool.
    CvVoIPPeerCfgOctetAligned interface{}

    // This object indicates modes of Bit rates. One or more upto four modes can
    // be configured at the same time as bit rates can be changed dynamically for
    // AMR-NB codec. This object is not instantiated when the object
    // cvVoIPPeerCfgCoderRate is not equal to gsmAmrNb enum. The type is
    // map[string]bool.
    CvVoIPPeerCfgBitRates interface{}

    // If the object has a value of true(1), frame CRC will be included in the
    // payload and if the value is false(2), frame CRC will not be included in the
    // payload. This object is applicable only when RTP frame type is octet
    // aligned. This object is not instantiated when the object
    // cvVoIPPeerCfgCoderRate is not equal to gsmAmrNb enum. The type is bool.
    CvVoIPPeerCfgCRC interface{}

    // This object indicates the iLBC codec mode to be used. The value of this
    // object is valid only if  cvVoIPPeerCfgCoderRate is equal to 'iLBC'. The
    // type is CvIlbcFrameMode.
    CvVoIPPeerCfgCoderMode interface{}

    // This object specifies the coding mode to be used. The object is
    // instantiated only if cvVoIPPeerCfgCoderRate is 'iSAC'. Following coding
    // modes are supported: adaptive    (1) - adaptive mode where iSAC performs
    // bandwidth                     estimation and adapts to the available
    // channel                    bandwidth. independent (2) - independent mode in
    // which no bandwidth estimation                    is performed. The type is
    // CvVoIPPeerCfgCodingMode.
    CvVoIPPeerCfgCodingMode interface{}

    // This object specifies the target bit rate. The object is instantiated only
    // if cvVoIPPeerCfgCoderRate is 'iSAC'. The type is interface{} with range:
    // 10000..32000.
    CvVoIPPeerCfgBitRate interface{}

    // This object specifies the frame size used. The object is instantiated only
    // if cvVoIPPeerCfgCoderRate is 'iSAC'. The frame size can be 30 ms or 60 ms,
    // and it can be fixed for all packets or vary depending on the configuration
    // and bandwidth estimation. Thus it can have the following values:
    // frameSize30      - initial frame size of 30 ms frameSize60      - initial
    // frame size of 60 ms frameSize30fixed - fixed frame size 30 ms
    // frameSize60fixed - fixed frame size 60 ms. The type is
    // CvVoIPPeerCfgFrameSize.
    CvVoIPPeerCfgFrameSize interface{}

    // This object specifies whether cvdcPolicyViolationNotification traps should
    // be generated for the call that is associated with this peer for RPH to DSCP
    // mapping and policing feature. The type is bool.
    CvVoIPPeerCfgDSCPPolicyNotificationEnable interface{}

    // This object specifies whether cvdcPolicyViolationNotification traps should
    // be generated for the call that is associated with this peer for Media
    // policing feature.. The type is bool.
    CvVoIPPeerCfgMediaPolicyNotificationEnable interface{}
}

func (cvVoIPPeerCfgEntry *CISCOVOICEDIALCONTROLMIB_CvVoIPPeerCfgTable_CvVoIPPeerCfgEntry) GetEntityData() *types.CommonEntityData {
    cvVoIPPeerCfgEntry.EntityData.YFilter = cvVoIPPeerCfgEntry.YFilter
    cvVoIPPeerCfgEntry.EntityData.YangName = "cvVoIPPeerCfgEntry"
    cvVoIPPeerCfgEntry.EntityData.BundleName = "cisco_ios_xe"
    cvVoIPPeerCfgEntry.EntityData.ParentYangName = "cvVoIPPeerCfgTable"
    cvVoIPPeerCfgEntry.EntityData.SegmentPath = "cvVoIPPeerCfgEntry" + types.AddKeyToken(cvVoIPPeerCfgEntry.IfIndex, "ifIndex")
    cvVoIPPeerCfgEntry.EntityData.AbsolutePath = "CISCO-VOICE-DIAL-CONTROL-MIB:CISCO-VOICE-DIAL-CONTROL-MIB/cvVoIPPeerCfgTable/" + cvVoIPPeerCfgEntry.EntityData.SegmentPath
    cvVoIPPeerCfgEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvVoIPPeerCfgEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvVoIPPeerCfgEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvVoIPPeerCfgEntry.EntityData.Children = types.NewOrderedMap()
    cvVoIPPeerCfgEntry.EntityData.Leafs = types.NewOrderedMap()
    cvVoIPPeerCfgEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", cvVoIPPeerCfgEntry.IfIndex})
    cvVoIPPeerCfgEntry.EntityData.Leafs.Append("cvVoIPPeerCfgSessionProtocol", types.YLeaf{"CvVoIPPeerCfgSessionProtocol", cvVoIPPeerCfgEntry.CvVoIPPeerCfgSessionProtocol})
    cvVoIPPeerCfgEntry.EntityData.Leafs.Append("cvVoIPPeerCfgDesiredQoS", types.YLeaf{"CvVoIPPeerCfgDesiredQoS", cvVoIPPeerCfgEntry.CvVoIPPeerCfgDesiredQoS})
    cvVoIPPeerCfgEntry.EntityData.Leafs.Append("cvVoIPPeerCfgMinAcceptableQoS", types.YLeaf{"CvVoIPPeerCfgMinAcceptableQoS", cvVoIPPeerCfgEntry.CvVoIPPeerCfgMinAcceptableQoS})
    cvVoIPPeerCfgEntry.EntityData.Leafs.Append("cvVoIPPeerCfgSessionTarget", types.YLeaf{"CvVoIPPeerCfgSessionTarget", cvVoIPPeerCfgEntry.CvVoIPPeerCfgSessionTarget})
    cvVoIPPeerCfgEntry.EntityData.Leafs.Append("cvVoIPPeerCfgCoderRate", types.YLeaf{"CvVoIPPeerCfgCoderRate", cvVoIPPeerCfgEntry.CvVoIPPeerCfgCoderRate})
    cvVoIPPeerCfgEntry.EntityData.Leafs.Append("cvVoIPPeerCfgFaxRate", types.YLeaf{"CvVoIPPeerCfgFaxRate", cvVoIPPeerCfgEntry.CvVoIPPeerCfgFaxRate})
    cvVoIPPeerCfgEntry.EntityData.Leafs.Append("cvVoIPPeerCfgVADEnable", types.YLeaf{"CvVoIPPeerCfgVADEnable", cvVoIPPeerCfgEntry.CvVoIPPeerCfgVADEnable})
    cvVoIPPeerCfgEntry.EntityData.Leafs.Append("cvVoIPPeerCfgExpectFactor", types.YLeaf{"CvVoIPPeerCfgExpectFactor", cvVoIPPeerCfgEntry.CvVoIPPeerCfgExpectFactor})
    cvVoIPPeerCfgEntry.EntityData.Leafs.Append("cvVoIPPeerCfgIcpif", types.YLeaf{"CvVoIPPeerCfgIcpif", cvVoIPPeerCfgEntry.CvVoIPPeerCfgIcpif})
    cvVoIPPeerCfgEntry.EntityData.Leafs.Append("cvVoIPPeerCfgPoorQoVNotificationEnable", types.YLeaf{"CvVoIPPeerCfgPoorQoVNotificationEnable", cvVoIPPeerCfgEntry.CvVoIPPeerCfgPoorQoVNotificationEnable})
    cvVoIPPeerCfgEntry.EntityData.Leafs.Append("cvVoIPPeerCfgUDPChecksumEnable", types.YLeaf{"CvVoIPPeerCfgUDPChecksumEnable", cvVoIPPeerCfgEntry.CvVoIPPeerCfgUDPChecksumEnable})
    cvVoIPPeerCfgEntry.EntityData.Leafs.Append("cvVoIPPeerCfgIPPrecedence", types.YLeaf{"CvVoIPPeerCfgIPPrecedence", cvVoIPPeerCfgEntry.CvVoIPPeerCfgIPPrecedence})
    cvVoIPPeerCfgEntry.EntityData.Leafs.Append("cvVoIPPeerCfgTechPrefix", types.YLeaf{"CvVoIPPeerCfgTechPrefix", cvVoIPPeerCfgEntry.CvVoIPPeerCfgTechPrefix})
    cvVoIPPeerCfgEntry.EntityData.Leafs.Append("cvVoIPPeerCfgDigitRelay", types.YLeaf{"CvVoIPPeerCfgDigitRelay", cvVoIPPeerCfgEntry.CvVoIPPeerCfgDigitRelay})
    cvVoIPPeerCfgEntry.EntityData.Leafs.Append("cvVoIPPeerCfgCoderBytes", types.YLeaf{"CvVoIPPeerCfgCoderBytes", cvVoIPPeerCfgEntry.CvVoIPPeerCfgCoderBytes})
    cvVoIPPeerCfgEntry.EntityData.Leafs.Append("cvVoIPPeerCfgFaxBytes", types.YLeaf{"CvVoIPPeerCfgFaxBytes", cvVoIPPeerCfgEntry.CvVoIPPeerCfgFaxBytes})
    cvVoIPPeerCfgEntry.EntityData.Leafs.Append("cvVoIPPeerCfgInBandSignaling", types.YLeaf{"CvVoIPPeerCfgInBandSignaling", cvVoIPPeerCfgEntry.CvVoIPPeerCfgInBandSignaling})
    cvVoIPPeerCfgEntry.EntityData.Leafs.Append("cvVoIPPeerCfgMediaSetting", types.YLeaf{"CvVoIPPeerCfgMediaSetting", cvVoIPPeerCfgEntry.CvVoIPPeerCfgMediaSetting})
    cvVoIPPeerCfgEntry.EntityData.Leafs.Append("cvVoIPPeerCfgDesiredQoSVideo", types.YLeaf{"CvVoIPPeerCfgDesiredQoSVideo", cvVoIPPeerCfgEntry.CvVoIPPeerCfgDesiredQoSVideo})
    cvVoIPPeerCfgEntry.EntityData.Leafs.Append("cvVoIPPeerCfgMinAcceptableQoSVideo", types.YLeaf{"CvVoIPPeerCfgMinAcceptableQoSVideo", cvVoIPPeerCfgEntry.CvVoIPPeerCfgMinAcceptableQoSVideo})
    cvVoIPPeerCfgEntry.EntityData.Leafs.Append("cvVoIPPeerCfgRedirectip2ip", types.YLeaf{"CvVoIPPeerCfgRedirectip2ip", cvVoIPPeerCfgEntry.CvVoIPPeerCfgRedirectip2ip})
    cvVoIPPeerCfgEntry.EntityData.Leafs.Append("cvVoIPPeerCfgOctetAligned", types.YLeaf{"CvVoIPPeerCfgOctetAligned", cvVoIPPeerCfgEntry.CvVoIPPeerCfgOctetAligned})
    cvVoIPPeerCfgEntry.EntityData.Leafs.Append("cvVoIPPeerCfgBitRates", types.YLeaf{"CvVoIPPeerCfgBitRates", cvVoIPPeerCfgEntry.CvVoIPPeerCfgBitRates})
    cvVoIPPeerCfgEntry.EntityData.Leafs.Append("cvVoIPPeerCfgCRC", types.YLeaf{"CvVoIPPeerCfgCRC", cvVoIPPeerCfgEntry.CvVoIPPeerCfgCRC})
    cvVoIPPeerCfgEntry.EntityData.Leafs.Append("cvVoIPPeerCfgCoderMode", types.YLeaf{"CvVoIPPeerCfgCoderMode", cvVoIPPeerCfgEntry.CvVoIPPeerCfgCoderMode})
    cvVoIPPeerCfgEntry.EntityData.Leafs.Append("cvVoIPPeerCfgCodingMode", types.YLeaf{"CvVoIPPeerCfgCodingMode", cvVoIPPeerCfgEntry.CvVoIPPeerCfgCodingMode})
    cvVoIPPeerCfgEntry.EntityData.Leafs.Append("cvVoIPPeerCfgBitRate", types.YLeaf{"CvVoIPPeerCfgBitRate", cvVoIPPeerCfgEntry.CvVoIPPeerCfgBitRate})
    cvVoIPPeerCfgEntry.EntityData.Leafs.Append("cvVoIPPeerCfgFrameSize", types.YLeaf{"CvVoIPPeerCfgFrameSize", cvVoIPPeerCfgEntry.CvVoIPPeerCfgFrameSize})
    cvVoIPPeerCfgEntry.EntityData.Leafs.Append("cvVoIPPeerCfgDSCPPolicyNotificationEnable", types.YLeaf{"CvVoIPPeerCfgDSCPPolicyNotificationEnable", cvVoIPPeerCfgEntry.CvVoIPPeerCfgDSCPPolicyNotificationEnable})
    cvVoIPPeerCfgEntry.EntityData.Leafs.Append("cvVoIPPeerCfgMediaPolicyNotificationEnable", types.YLeaf{"CvVoIPPeerCfgMediaPolicyNotificationEnable", cvVoIPPeerCfgEntry.CvVoIPPeerCfgMediaPolicyNotificationEnable})

    cvVoIPPeerCfgEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(cvVoIPPeerCfgEntry.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_CvVoIPPeerCfgTable_CvVoIPPeerCfgEntry_CvVoIPPeerCfgCodingMode represents                   is performed.
type CISCOVOICEDIALCONTROLMIB_CvVoIPPeerCfgTable_CvVoIPPeerCfgEntry_CvVoIPPeerCfgCodingMode string

const (
    CISCOVOICEDIALCONTROLMIB_CvVoIPPeerCfgTable_CvVoIPPeerCfgEntry_CvVoIPPeerCfgCodingMode_adaptive CISCOVOICEDIALCONTROLMIB_CvVoIPPeerCfgTable_CvVoIPPeerCfgEntry_CvVoIPPeerCfgCodingMode = "adaptive"

    CISCOVOICEDIALCONTROLMIB_CvVoIPPeerCfgTable_CvVoIPPeerCfgEntry_CvVoIPPeerCfgCodingMode_independent CISCOVOICEDIALCONTROLMIB_CvVoIPPeerCfgTable_CvVoIPPeerCfgEntry_CvVoIPPeerCfgCodingMode = "independent"
)

// CISCOVOICEDIALCONTROLMIB_CvVoIPPeerCfgTable_CvVoIPPeerCfgEntry_CvVoIPPeerCfgFrameSize represents frameSize60fixed - fixed frame size 60 ms
type CISCOVOICEDIALCONTROLMIB_CvVoIPPeerCfgTable_CvVoIPPeerCfgEntry_CvVoIPPeerCfgFrameSize string

const (
    CISCOVOICEDIALCONTROLMIB_CvVoIPPeerCfgTable_CvVoIPPeerCfgEntry_CvVoIPPeerCfgFrameSize_frameSize30 CISCOVOICEDIALCONTROLMIB_CvVoIPPeerCfgTable_CvVoIPPeerCfgEntry_CvVoIPPeerCfgFrameSize = "frameSize30"

    CISCOVOICEDIALCONTROLMIB_CvVoIPPeerCfgTable_CvVoIPPeerCfgEntry_CvVoIPPeerCfgFrameSize_frameSize60 CISCOVOICEDIALCONTROLMIB_CvVoIPPeerCfgTable_CvVoIPPeerCfgEntry_CvVoIPPeerCfgFrameSize = "frameSize60"

    CISCOVOICEDIALCONTROLMIB_CvVoIPPeerCfgTable_CvVoIPPeerCfgEntry_CvVoIPPeerCfgFrameSize_frameSize30fixed CISCOVOICEDIALCONTROLMIB_CvVoIPPeerCfgTable_CvVoIPPeerCfgEntry_CvVoIPPeerCfgFrameSize = "frameSize30fixed"

    CISCOVOICEDIALCONTROLMIB_CvVoIPPeerCfgTable_CvVoIPPeerCfgEntry_CvVoIPPeerCfgFrameSize_frameSize60fixed CISCOVOICEDIALCONTROLMIB_CvVoIPPeerCfgTable_CvVoIPPeerCfgEntry_CvVoIPPeerCfgFrameSize = "frameSize60fixed"
)

// CISCOVOICEDIALCONTROLMIB_CvVoIPPeerCfgTable_CvVoIPPeerCfgEntry_CvVoIPPeerCfgMediaSetting represents be established directly between the endpoints.
type CISCOVOICEDIALCONTROLMIB_CvVoIPPeerCfgTable_CvVoIPPeerCfgEntry_CvVoIPPeerCfgMediaSetting string

const (
    CISCOVOICEDIALCONTROLMIB_CvVoIPPeerCfgTable_CvVoIPPeerCfgEntry_CvVoIPPeerCfgMediaSetting_flowThrough CISCOVOICEDIALCONTROLMIB_CvVoIPPeerCfgTable_CvVoIPPeerCfgEntry_CvVoIPPeerCfgMediaSetting = "flowThrough"

    CISCOVOICEDIALCONTROLMIB_CvVoIPPeerCfgTable_CvVoIPPeerCfgEntry_CvVoIPPeerCfgMediaSetting_flowAround CISCOVOICEDIALCONTROLMIB_CvVoIPPeerCfgTable_CvVoIPPeerCfgEntry_CvVoIPPeerCfgMediaSetting = "flowAround"
)

// CISCOVOICEDIALCONTROLMIB_CvPeerCommonCfgTable
// The table contains the Voice specific peer common
// configuration information that is required to accept voice
// calls or to which it will place them or process the
// incoming calls.
type CISCOVOICEDIALCONTROLMIB_CvPeerCommonCfgTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A single Voice specific Peer. One entry per voice related encapsulation.
    // The entry is created when a voice related encapsulation ifEntry is created.
    // This entry is deleted when its associated ifEntry is deleted. The type is
    // slice of
    // CISCOVOICEDIALCONTROLMIB_CvPeerCommonCfgTable_CvPeerCommonCfgEntry.
    CvPeerCommonCfgEntry []*CISCOVOICEDIALCONTROLMIB_CvPeerCommonCfgTable_CvPeerCommonCfgEntry
}

func (cvPeerCommonCfgTable *CISCOVOICEDIALCONTROLMIB_CvPeerCommonCfgTable) GetEntityData() *types.CommonEntityData {
    cvPeerCommonCfgTable.EntityData.YFilter = cvPeerCommonCfgTable.YFilter
    cvPeerCommonCfgTable.EntityData.YangName = "cvPeerCommonCfgTable"
    cvPeerCommonCfgTable.EntityData.BundleName = "cisco_ios_xe"
    cvPeerCommonCfgTable.EntityData.ParentYangName = "CISCO-VOICE-DIAL-CONTROL-MIB"
    cvPeerCommonCfgTable.EntityData.SegmentPath = "cvPeerCommonCfgTable"
    cvPeerCommonCfgTable.EntityData.AbsolutePath = "CISCO-VOICE-DIAL-CONTROL-MIB:CISCO-VOICE-DIAL-CONTROL-MIB/" + cvPeerCommonCfgTable.EntityData.SegmentPath
    cvPeerCommonCfgTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvPeerCommonCfgTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvPeerCommonCfgTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvPeerCommonCfgTable.EntityData.Children = types.NewOrderedMap()
    cvPeerCommonCfgTable.EntityData.Children.Append("cvPeerCommonCfgEntry", types.YChild{"CvPeerCommonCfgEntry", nil})
    for i := range cvPeerCommonCfgTable.CvPeerCommonCfgEntry {
        cvPeerCommonCfgTable.EntityData.Children.Append(types.GetSegmentPath(cvPeerCommonCfgTable.CvPeerCommonCfgEntry[i]), types.YChild{"CvPeerCommonCfgEntry", cvPeerCommonCfgTable.CvPeerCommonCfgEntry[i]})
    }
    cvPeerCommonCfgTable.EntityData.Leafs = types.NewOrderedMap()

    cvPeerCommonCfgTable.EntityData.YListKeys = []string {}

    return &(cvPeerCommonCfgTable.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_CvPeerCommonCfgTable_CvPeerCommonCfgEntry
// A single Voice specific Peer. One entry per voice related
// encapsulation.
// The entry is created when a voice related encapsulation
// ifEntry is created.
// This entry is deleted when its associated ifEntry is
// deleted.
type CISCOVOICEDIALCONTROLMIB_CvPeerCommonCfgTable_CvPeerCommonCfgEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // The object specifies the prefix of the incoming Dialed Number
    // Identification Service (DNIS) digits for the peer. The DNIS digits prefix
    // is used to match with the incoming DNIS number for incoming call
    // discrimination. If the digits in this object are matched with incoming DNIS
    // number, the  associated dialCtlPeerCfgInfoType in RFC 2128 will be used as
    // a call discriminator for differentiating speech, data, fax, video or modem
    // calls. The type is string with length: 0..32.
    CvPeerCommonCfgIncomingDnisDigits interface{}

    // The object specifies the maximum allowed connection to/from the peer. A
    // value of -1 disables the limit of maximum connections. The type is
    // interface{} with range: -1..None | 1..2147483647. Units are connections.
    CvPeerCommonCfgMaxConnections interface{}

    // The object specifies the application to handle the incoming call after the
    // peer is selected. If no application name is specified, then the default
    // session application will take care of the incoming call. The type is
    // string.
    CvPeerCommonCfgApplicationName interface{}

    // This object specifies the selection preference of a peer when multiple
    // peers are matched to the selection criteria. The value of 0 has the lowest
    // preference for peer selection. The type is interface{} with range: 0..10.
    CvPeerCommonCfgPreference interface{}

    // This object specifies whether dialpeer hunting should stop when this peer
    // is reached. The type is bool.
    CvPeerCommonCfgHuntStop interface{}

    // The object specifies a Dialer Number Identification Service (DNIS) map name
    // for the Voice specific peer entry specified in this row. A DNIS is a called
    // party number and they can be grouped and identified by DNIS map. The type
    // is string.
    CvPeerCommonCfgDnisMappingName interface{}

    // The object specifies the Source Carrier Id for the peer. The Source Carrier
    // Id is used to match with the Source Carrier Id of a call. If the Source
    // Carrier Id in this object is matched with the Source Carrier Id of a call,
    // then the associated peer will be used to handle the call.  Only
    // alphanumeric characters, '-' and '_' are allowed in the string. The type is
    // string.
    CvPeerCommonCfgSourceCarrierId interface{}

    // The object specifies the Target Carrier Id for the peer. The Target Carrier
    // Id is used to match with the Target Carrier Id of a call. If the Target
    // Carrier Id in this object is matched with the Target Carrier Id of a call,
    // then the associated peer will be used to handle the call. Only alphanumeric
    // characters, '-' and '_' are allowed in the string. The type is string.
    CvPeerCommonCfgTargetCarrierId interface{}

    // The object specifies the Source Trunk Group Label for the peer. The Source
    // Trunk Group Label is used to match with the Source Trunk Group Label of a
    // call. If the Source Trunk Group Label in this object is matched with the
    // Source Trunk Group Label of a call, then the associated peer will be used
    // to handle the call.  Only alphanumeric characters, '-' and '_' are allowed
    // in the string. The type is string.
    CvPeerCommonCfgSourceTrunkGrpLabel interface{}

    // The object specifies the Target Trunk Group Label for the peer. The Target
    // Trunk Group Label is used to match with the Target Trunk Group Label of a
    // call. If the Target Trunk Group Label in this object is matched with the
    // Target Trunk Group Label of a call, then the associated peer will be used
    // to handle the call. Only alphanumeric characters, '-' and '_' are allowed
    // in the string. The type is string.
    CvPeerCommonCfgTargetTrunkGrpLabel interface{}
}

func (cvPeerCommonCfgEntry *CISCOVOICEDIALCONTROLMIB_CvPeerCommonCfgTable_CvPeerCommonCfgEntry) GetEntityData() *types.CommonEntityData {
    cvPeerCommonCfgEntry.EntityData.YFilter = cvPeerCommonCfgEntry.YFilter
    cvPeerCommonCfgEntry.EntityData.YangName = "cvPeerCommonCfgEntry"
    cvPeerCommonCfgEntry.EntityData.BundleName = "cisco_ios_xe"
    cvPeerCommonCfgEntry.EntityData.ParentYangName = "cvPeerCommonCfgTable"
    cvPeerCommonCfgEntry.EntityData.SegmentPath = "cvPeerCommonCfgEntry" + types.AddKeyToken(cvPeerCommonCfgEntry.IfIndex, "ifIndex")
    cvPeerCommonCfgEntry.EntityData.AbsolutePath = "CISCO-VOICE-DIAL-CONTROL-MIB:CISCO-VOICE-DIAL-CONTROL-MIB/cvPeerCommonCfgTable/" + cvPeerCommonCfgEntry.EntityData.SegmentPath
    cvPeerCommonCfgEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvPeerCommonCfgEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvPeerCommonCfgEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvPeerCommonCfgEntry.EntityData.Children = types.NewOrderedMap()
    cvPeerCommonCfgEntry.EntityData.Leafs = types.NewOrderedMap()
    cvPeerCommonCfgEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", cvPeerCommonCfgEntry.IfIndex})
    cvPeerCommonCfgEntry.EntityData.Leafs.Append("cvPeerCommonCfgIncomingDnisDigits", types.YLeaf{"CvPeerCommonCfgIncomingDnisDigits", cvPeerCommonCfgEntry.CvPeerCommonCfgIncomingDnisDigits})
    cvPeerCommonCfgEntry.EntityData.Leafs.Append("cvPeerCommonCfgMaxConnections", types.YLeaf{"CvPeerCommonCfgMaxConnections", cvPeerCommonCfgEntry.CvPeerCommonCfgMaxConnections})
    cvPeerCommonCfgEntry.EntityData.Leafs.Append("cvPeerCommonCfgApplicationName", types.YLeaf{"CvPeerCommonCfgApplicationName", cvPeerCommonCfgEntry.CvPeerCommonCfgApplicationName})
    cvPeerCommonCfgEntry.EntityData.Leafs.Append("cvPeerCommonCfgPreference", types.YLeaf{"CvPeerCommonCfgPreference", cvPeerCommonCfgEntry.CvPeerCommonCfgPreference})
    cvPeerCommonCfgEntry.EntityData.Leafs.Append("cvPeerCommonCfgHuntStop", types.YLeaf{"CvPeerCommonCfgHuntStop", cvPeerCommonCfgEntry.CvPeerCommonCfgHuntStop})
    cvPeerCommonCfgEntry.EntityData.Leafs.Append("cvPeerCommonCfgDnisMappingName", types.YLeaf{"CvPeerCommonCfgDnisMappingName", cvPeerCommonCfgEntry.CvPeerCommonCfgDnisMappingName})
    cvPeerCommonCfgEntry.EntityData.Leafs.Append("cvPeerCommonCfgSourceCarrierId", types.YLeaf{"CvPeerCommonCfgSourceCarrierId", cvPeerCommonCfgEntry.CvPeerCommonCfgSourceCarrierId})
    cvPeerCommonCfgEntry.EntityData.Leafs.Append("cvPeerCommonCfgTargetCarrierId", types.YLeaf{"CvPeerCommonCfgTargetCarrierId", cvPeerCommonCfgEntry.CvPeerCommonCfgTargetCarrierId})
    cvPeerCommonCfgEntry.EntityData.Leafs.Append("cvPeerCommonCfgSourceTrunkGrpLabel", types.YLeaf{"CvPeerCommonCfgSourceTrunkGrpLabel", cvPeerCommonCfgEntry.CvPeerCommonCfgSourceTrunkGrpLabel})
    cvPeerCommonCfgEntry.EntityData.Leafs.Append("cvPeerCommonCfgTargetTrunkGrpLabel", types.YLeaf{"CvPeerCommonCfgTargetTrunkGrpLabel", cvPeerCommonCfgEntry.CvPeerCommonCfgTargetTrunkGrpLabel})

    cvPeerCommonCfgEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(cvPeerCommonCfgEntry.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_CvCallActiveTable
// This table is the voice extension to the call active table
// of IETF Dial Control MIB. It contains voice encapsulation
// call leg information that is derived from the statistics
// of lower layer telephony interface.
type CISCOVOICEDIALCONTROLMIB_CvCallActiveTable struct {
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
    // CISCOVOICEDIALCONTROLMIB_CvCallActiveTable_CvCallActiveEntry.
    CvCallActiveEntry []*CISCOVOICEDIALCONTROLMIB_CvCallActiveTable_CvCallActiveEntry
}

func (cvCallActiveTable *CISCOVOICEDIALCONTROLMIB_CvCallActiveTable) GetEntityData() *types.CommonEntityData {
    cvCallActiveTable.EntityData.YFilter = cvCallActiveTable.YFilter
    cvCallActiveTable.EntityData.YangName = "cvCallActiveTable"
    cvCallActiveTable.EntityData.BundleName = "cisco_ios_xe"
    cvCallActiveTable.EntityData.ParentYangName = "CISCO-VOICE-DIAL-CONTROL-MIB"
    cvCallActiveTable.EntityData.SegmentPath = "cvCallActiveTable"
    cvCallActiveTable.EntityData.AbsolutePath = "CISCO-VOICE-DIAL-CONTROL-MIB:CISCO-VOICE-DIAL-CONTROL-MIB/" + cvCallActiveTable.EntityData.SegmentPath
    cvCallActiveTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvCallActiveTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvCallActiveTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvCallActiveTable.EntityData.Children = types.NewOrderedMap()
    cvCallActiveTable.EntityData.Children.Append("cvCallActiveEntry", types.YChild{"CvCallActiveEntry", nil})
    for i := range cvCallActiveTable.CvCallActiveEntry {
        cvCallActiveTable.EntityData.Children.Append(types.GetSegmentPath(cvCallActiveTable.CvCallActiveEntry[i]), types.YChild{"CvCallActiveEntry", cvCallActiveTable.CvCallActiveEntry[i]})
    }
    cvCallActiveTable.EntityData.Leafs = types.NewOrderedMap()

    cvCallActiveTable.EntityData.YListKeys = []string {}

    return &(cvCallActiveTable.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_CvCallActiveTable_CvCallActiveEntry
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
type CISCOVOICEDIALCONTROLMIB_CvCallActiveTable_CvCallActiveEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // dial_control_mib.DIALCONTROLMIB_CallActiveTable_CallActiveEntry_CallActiveSetupTime
    CallActiveSetupTime interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // dial_control_mib.DIALCONTROLMIB_CallActiveTable_CallActiveEntry_CallActiveIndex
    CallActiveIndex interface{}

    // The global connection identifier for the active telephony leg of the call.
    // The type is string with length: 0..16.
    CvCallActiveConnectionId interface{}

    // Duration of Transmit path open from this peer to the voice gateway for the
    // call leg. This counter object will lock at the maximum value which is
    // approximately two days. The type is interface{} with range: 0..4294967295.
    // Units are milliseconds.
    CvCallActiveTxDuration interface{}

    // Duration of voice transmitted from this peer to voice gateway for this call
    // leg. The Voice Utilization Rate can be obtained by dividing this by
    // cvCallActiveTXDuration object. This counter object will lock at the maximum
    // value which is approximately two days. The type is interface{} with range:
    // 0..4294967295. Units are milliseconds.
    CvCallActiveVoiceTxDuration interface{}

    // Duration of fax transmitted from this peer to voice gateway for this call
    // leg. The FAX Utilization Rate can be obtained by dividing this by
    // cvCallActiveTXDuration object. This counter object will lock at the maximum
    // value which is approximately two days. The type is interface{} with range:
    // 0..4294967295. Units are milliseconds.
    CvCallActiveFaxTxDuration interface{}

    // The negotiated coder rate. It specifies the transmit rate of voice/fax
    // compression to its associated call leg for the call. The type is
    // CvcCoderTypeRate.
    CvCallActiveCoderTypeRate interface{}

    // The object contains the active noise level for the call leg. The type is
    // interface{} with range: -128..8. Units are dBm.
    CvCallActiveNoiseLevel interface{}

    // The object contains the sum of Echo Return Loss (ERL), cancellation loss
    // (Echo Return Loss Enhancement) and nonlinear processing loss for the call
    // leg. The value -1 indicates the level can not be determined or level
    // detection is disabled. The type is interface{} with range: -1..127. Units
    // are dB.
    CvCallActiveACOMLevel interface{}

    // The object contains the active output signal level to telephony interface
    // that is used by the call leg. The type is interface{} with range: -128..8.
    // Units are dBm.
    CvCallActiveOutSignalLevel interface{}

    // The object contains the active input signal level from telephony interface
    // that is used by the call leg. The type is interface{} with range: -128..8.
    // Units are dBm.
    CvCallActiveInSignalLevel interface{}

    // The object contains the current Echo Return Loss (ERL) level for the call
    // leg. The value -1 indicates the level can not be determined or level
    // detection is disabled. The type is interface{} with range: -1..45. Units
    // are dB.
    CvCallActiveERLLevel interface{}

    // The object specifies the session target of the peer that is used for the
    // call leg. This object is set with the information in the call associated
    // cvVoicePeerCfgSessionTarget object when the call is connected. The type is
    // string with length: 0..64.
    CvCallActiveSessionTarget interface{}

    // The number of FAX related image pages are received or transmitted via the
    // peer for the call leg. The type is interface{} with range: 0..4294967295.
    // Units are pages.
    CvCallActiveImgPageCount interface{}

    // The calling party name of the call. If the name is not available, then it
    // will have a length of zero. The type is string.
    CvCallActiveCallingName interface{}

    // The object indicates whether or not the caller ID feature is blocked for
    // this call. The type is bool.
    CvCallActiveCallerIDBlock interface{}

    // The location in milliseconds of the largest amplitude reflector detected by
    // the echo canceller for this call.  The value 0 indicates there is no
    // reflector or the  information is not available. The type is interface{}
    // with range: 0..128.
    CvCallActiveEcanReflectorLocation interface{}

    // The object indicates the account code input to the call. It could be used
    // for call screen or by down stream server for billing purpose. The value of
    // empty string indicates no account code input. The type is string with
    // length: 0..50.
    CvCallActiveAccountCode interface{}

    // The object contains the current Echo Return Loss (ERL) level for the call
    // leg. The value -1 indicates the level can not be determined or level
    // detection is disabled. The type is interface{} with range: -1..200. Units
    // are dB.
    CvCallActiveERLLevelRev1 interface{}

    // This object represents the call identifier for the active telephony leg of
    // the call. The type is interface{} with range: 1..4294967295.
    CvCallActiveCallId interface{}
}

func (cvCallActiveEntry *CISCOVOICEDIALCONTROLMIB_CvCallActiveTable_CvCallActiveEntry) GetEntityData() *types.CommonEntityData {
    cvCallActiveEntry.EntityData.YFilter = cvCallActiveEntry.YFilter
    cvCallActiveEntry.EntityData.YangName = "cvCallActiveEntry"
    cvCallActiveEntry.EntityData.BundleName = "cisco_ios_xe"
    cvCallActiveEntry.EntityData.ParentYangName = "cvCallActiveTable"
    cvCallActiveEntry.EntityData.SegmentPath = "cvCallActiveEntry" + types.AddKeyToken(cvCallActiveEntry.CallActiveSetupTime, "callActiveSetupTime") + types.AddKeyToken(cvCallActiveEntry.CallActiveIndex, "callActiveIndex")
    cvCallActiveEntry.EntityData.AbsolutePath = "CISCO-VOICE-DIAL-CONTROL-MIB:CISCO-VOICE-DIAL-CONTROL-MIB/cvCallActiveTable/" + cvCallActiveEntry.EntityData.SegmentPath
    cvCallActiveEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvCallActiveEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvCallActiveEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvCallActiveEntry.EntityData.Children = types.NewOrderedMap()
    cvCallActiveEntry.EntityData.Leafs = types.NewOrderedMap()
    cvCallActiveEntry.EntityData.Leafs.Append("callActiveSetupTime", types.YLeaf{"CallActiveSetupTime", cvCallActiveEntry.CallActiveSetupTime})
    cvCallActiveEntry.EntityData.Leafs.Append("callActiveIndex", types.YLeaf{"CallActiveIndex", cvCallActiveEntry.CallActiveIndex})
    cvCallActiveEntry.EntityData.Leafs.Append("cvCallActiveConnectionId", types.YLeaf{"CvCallActiveConnectionId", cvCallActiveEntry.CvCallActiveConnectionId})
    cvCallActiveEntry.EntityData.Leafs.Append("cvCallActiveTxDuration", types.YLeaf{"CvCallActiveTxDuration", cvCallActiveEntry.CvCallActiveTxDuration})
    cvCallActiveEntry.EntityData.Leafs.Append("cvCallActiveVoiceTxDuration", types.YLeaf{"CvCallActiveVoiceTxDuration", cvCallActiveEntry.CvCallActiveVoiceTxDuration})
    cvCallActiveEntry.EntityData.Leafs.Append("cvCallActiveFaxTxDuration", types.YLeaf{"CvCallActiveFaxTxDuration", cvCallActiveEntry.CvCallActiveFaxTxDuration})
    cvCallActiveEntry.EntityData.Leafs.Append("cvCallActiveCoderTypeRate", types.YLeaf{"CvCallActiveCoderTypeRate", cvCallActiveEntry.CvCallActiveCoderTypeRate})
    cvCallActiveEntry.EntityData.Leafs.Append("cvCallActiveNoiseLevel", types.YLeaf{"CvCallActiveNoiseLevel", cvCallActiveEntry.CvCallActiveNoiseLevel})
    cvCallActiveEntry.EntityData.Leafs.Append("cvCallActiveACOMLevel", types.YLeaf{"CvCallActiveACOMLevel", cvCallActiveEntry.CvCallActiveACOMLevel})
    cvCallActiveEntry.EntityData.Leafs.Append("cvCallActiveOutSignalLevel", types.YLeaf{"CvCallActiveOutSignalLevel", cvCallActiveEntry.CvCallActiveOutSignalLevel})
    cvCallActiveEntry.EntityData.Leafs.Append("cvCallActiveInSignalLevel", types.YLeaf{"CvCallActiveInSignalLevel", cvCallActiveEntry.CvCallActiveInSignalLevel})
    cvCallActiveEntry.EntityData.Leafs.Append("cvCallActiveERLLevel", types.YLeaf{"CvCallActiveERLLevel", cvCallActiveEntry.CvCallActiveERLLevel})
    cvCallActiveEntry.EntityData.Leafs.Append("cvCallActiveSessionTarget", types.YLeaf{"CvCallActiveSessionTarget", cvCallActiveEntry.CvCallActiveSessionTarget})
    cvCallActiveEntry.EntityData.Leafs.Append("cvCallActiveImgPageCount", types.YLeaf{"CvCallActiveImgPageCount", cvCallActiveEntry.CvCallActiveImgPageCount})
    cvCallActiveEntry.EntityData.Leafs.Append("cvCallActiveCallingName", types.YLeaf{"CvCallActiveCallingName", cvCallActiveEntry.CvCallActiveCallingName})
    cvCallActiveEntry.EntityData.Leafs.Append("cvCallActiveCallerIDBlock", types.YLeaf{"CvCallActiveCallerIDBlock", cvCallActiveEntry.CvCallActiveCallerIDBlock})
    cvCallActiveEntry.EntityData.Leafs.Append("cvCallActiveEcanReflectorLocation", types.YLeaf{"CvCallActiveEcanReflectorLocation", cvCallActiveEntry.CvCallActiveEcanReflectorLocation})
    cvCallActiveEntry.EntityData.Leafs.Append("cvCallActiveAccountCode", types.YLeaf{"CvCallActiveAccountCode", cvCallActiveEntry.CvCallActiveAccountCode})
    cvCallActiveEntry.EntityData.Leafs.Append("cvCallActiveERLLevelRev1", types.YLeaf{"CvCallActiveERLLevelRev1", cvCallActiveEntry.CvCallActiveERLLevelRev1})
    cvCallActiveEntry.EntityData.Leafs.Append("cvCallActiveCallId", types.YLeaf{"CvCallActiveCallId", cvCallActiveEntry.CvCallActiveCallId})

    cvCallActiveEntry.EntityData.YListKeys = []string {"CallActiveSetupTime", "CallActiveIndex"}

    return &(cvCallActiveEntry.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_CvVoIPCallActiveTable
// This table is the VoIP extension to the call active table of
// IETF Dial Control MIB. It contains VoIP call leg
// information about specific VoIP call destination and the
// selected QoS for the call leg.
type CISCOVOICEDIALCONTROLMIB_CvVoIPCallActiveTable struct {
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
    // CISCOVOICEDIALCONTROLMIB_CvVoIPCallActiveTable_CvVoIPCallActiveEntry.
    CvVoIPCallActiveEntry []*CISCOVOICEDIALCONTROLMIB_CvVoIPCallActiveTable_CvVoIPCallActiveEntry
}

func (cvVoIPCallActiveTable *CISCOVOICEDIALCONTROLMIB_CvVoIPCallActiveTable) GetEntityData() *types.CommonEntityData {
    cvVoIPCallActiveTable.EntityData.YFilter = cvVoIPCallActiveTable.YFilter
    cvVoIPCallActiveTable.EntityData.YangName = "cvVoIPCallActiveTable"
    cvVoIPCallActiveTable.EntityData.BundleName = "cisco_ios_xe"
    cvVoIPCallActiveTable.EntityData.ParentYangName = "CISCO-VOICE-DIAL-CONTROL-MIB"
    cvVoIPCallActiveTable.EntityData.SegmentPath = "cvVoIPCallActiveTable"
    cvVoIPCallActiveTable.EntityData.AbsolutePath = "CISCO-VOICE-DIAL-CONTROL-MIB:CISCO-VOICE-DIAL-CONTROL-MIB/" + cvVoIPCallActiveTable.EntityData.SegmentPath
    cvVoIPCallActiveTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvVoIPCallActiveTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvVoIPCallActiveTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvVoIPCallActiveTable.EntityData.Children = types.NewOrderedMap()
    cvVoIPCallActiveTable.EntityData.Children.Append("cvVoIPCallActiveEntry", types.YChild{"CvVoIPCallActiveEntry", nil})
    for i := range cvVoIPCallActiveTable.CvVoIPCallActiveEntry {
        cvVoIPCallActiveTable.EntityData.Children.Append(types.GetSegmentPath(cvVoIPCallActiveTable.CvVoIPCallActiveEntry[i]), types.YChild{"CvVoIPCallActiveEntry", cvVoIPCallActiveTable.CvVoIPCallActiveEntry[i]})
    }
    cvVoIPCallActiveTable.EntityData.Leafs = types.NewOrderedMap()

    cvVoIPCallActiveTable.EntityData.YListKeys = []string {}

    return &(cvVoIPCallActiveTable.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_CvVoIPCallActiveTable_CvVoIPCallActiveEntry
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
type CISCOVOICEDIALCONTROLMIB_CvVoIPCallActiveTable_CvVoIPCallActiveEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // dial_control_mib.DIALCONTROLMIB_CallActiveTable_CallActiveEntry_CallActiveSetupTime
    CallActiveSetupTime interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // dial_control_mib.DIALCONTROLMIB_CallActiveTable_CallActiveEntry_CallActiveIndex
    CallActiveIndex interface{}

    // The global connection identifier for the active VoIP leg of the call. The
    // type is string with length: 0..16.
    CvVoIPCallActiveConnectionId interface{}

    // Remote system IP address for the VoIP call. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    CvVoIPCallActiveRemoteIPAddress interface{}

    // Remote system UDP listener port to which to transmit voice packets. The
    // type is interface{} with range: 0..65535.
    CvVoIPCallActiveRemoteUDPPort interface{}

    // The voice packet round trip delay between local and the remote system on
    // the IP backbone during the call. The type is interface{} with range:
    // 0..4294967295. Units are milliseconds.
    CvVoIPCallActiveRoundTripDelay interface{}

    // The selected RSVP QoS for the voice call. The type is QosService.
    CvVoIPCallActiveSelectedQoS interface{}

    // The object specifies the session protocol to be used for Internet call
    // between local and remote router via IP backbone. The type is
    // CvSessionProtocol.
    CvVoIPCallActiveSessionProtocol interface{}

    // The object specifies the session target of the peer that is used for the
    // call. This object is set with the information in the call associated
    // cvVoIPPeerCfgSessionTarget object when the voice over IP call is connected.
    // The type is string.
    CvVoIPCallActiveSessionTarget interface{}

    // Duration of voice playout from data received on time for this call. This
    // plus the durations for the GapFills in the following entries gives the
    // Total Voice Playout Duration for Active Voice. This does not include
    // duration for which no data was sent by the Transmit end as voice signal,
    // e.g., silence suppression and fax signal. The On Time Playout Rate can be
    // computed by dividing this entry by the Total Voice Playout Duration. This
    // counter object will lock at the maximum value which is approximately two
    // days. The type is interface{} with range: 0..4294967295. Units are
    // milliseconds.
    CvVoIPCallActiveOnTimeRvPlayout interface{}

    // Duration of voice signal replaced with signal played out during silence due
    // to voice data not received on time (or lost) from voice gateway this call.
    // This counter object will lock at the maximum value which is approximately
    // two days. The type is interface{} with range: 0..4294967295. Units are
    // milliseconds.
    CvVoIPCallActiveGapFillWithSilence interface{}

    // Duration of voice signal played out with signal synthesized from parameters
    // or samples of data preceding in time due to voice data not received on time
    // (or lost) from voice gateway for this call. An example of such playout is
    // frame-erasure or frame-concealment strategies in G.729 and G.723.1
    // compression algorithms. This counter object will lock at the maximum value
    // which is approximately two days. The type is interface{} with range:
    // 0..4294967295. Units are milliseconds.
    CvVoIPCallActiveGapFillWithPrediction interface{}

    // Duration of voice signal played out with signal synthesized from parameters
    // or samples of data preceding and following in time due to voice data not
    // received on time (or lost) from voice gateway for this call. This counter
    // object will lock at the maximum value which is approximately two days. The
    // type is interface{} with range: 0..4294967295. Units are milliseconds.
    CvVoIPCallActiveGapFillWithInterpolation interface{}

    // Duration of voice signal played out with signal synthesized from redundancy
    // parameters available due to voice data not received on time (or lost) from
    // voice gateway for this call. This counter object will lock at the maximum
    // value which is approximately two days. The type is interface{} with range:
    // 0..4294967295. Units are milliseconds.
    CvVoIPCallActiveGapFillWithRedundancy interface{}

    // The high water mark Voice Playout FIFO Delay during the voice call. This
    // counter object will lock at the maximum value which is approximately two
    // days. The type is interface{} with range: 0..4294967295. Units are
    // milliseconds.
    CvVoIPCallActiveHiWaterPlayoutDelay interface{}

    // The low water mark Voice Playout FIFO Delay during the voice call. The type
    // is interface{} with range: 0..4294967295. Units are milliseconds.
    CvVoIPCallActiveLoWaterPlayoutDelay interface{}

    // The average Playout FIFO Delay plus the decoder delay during the voice
    // call. The type is interface{} with range: 0..4294967295.
    CvVoIPCallActiveReceiveDelay interface{}

    // The object indicates whether or not the VAD (Voice Activity Detection) was
    // enabled for the voice call. The type is bool.
    CvVoIPCallActiveVADEnable interface{}

    // The negotiated coder rate. It specifies the transmit rate of voice/fax
    // compression to its associated call leg for the call. This rate is different
    // from the configuration rate because of rate negotiation during the call.
    // The type is CvcCoderTypeRate.
    CvVoIPCallActiveCoderTypeRate interface{}

    // The number of lost voice packets during the call. The type is interface{}
    // with range: 0..4294967295. Units are packets.
    CvVoIPCallActiveLostPackets interface{}

    // The number of received voice packets that arrived too early to store in
    // jitter buffer during the call. The type is interface{} with range:
    // 0..4294967295. Units are packets.
    CvVoIPCallActiveEarlyPackets interface{}

    // The number of received voice packets that arrived too late to playout with
    // CODEC (Coder/Decoder) during the call. The type is interface{} with range:
    // 0..4294967295. Units are packets.
    CvVoIPCallActiveLatePackets interface{}

    // The textual identifier of the calling party (user) of the call. If the
    // username is not available, then the value of this object will have a length
    // of zero. The type is string with length: 0..16.
    CvVoIPCallActiveUsername interface{}

    // The protocol-specific call identifier for the VoIP call. The type is string
    // with length: 0..255.
    CvVoIPCallActiveProtocolCallId interface{}

    // This object specifies the type of address contained in the associated
    // instance of cvVoIPCallActiveRemSigIPAddr. The type is InetAddressType.
    CvVoIPCallActiveRemSigIPAddrT interface{}

    // Remote signalling IP address for the VoIP call. The type is string with
    // length: 0..255.
    CvVoIPCallActiveRemSigIPAddr interface{}

    // Remote signalling listener port to which to transmit voice packets. The
    // type is interface{} with range: 0..65535.
    CvVoIPCallActiveRemSigPort interface{}

    // This object specifies the type of address contained in the associated
    // instance of cvVoIPCallActiveRemMediaIPAddr. The type is InetAddressType.
    CvVoIPCallActiveRemMediaIPAddrT interface{}

    // Remote media end point IP address for the VoIP call. The type is string
    // with length: 0..255.
    CvVoIPCallActiveRemMediaIPAddr interface{}

    // Remote media end point listener port to which to transmit voice packets.
    // The type is interface{} with range: 0..65535.
    CvVoIPCallActiveRemMediaPort interface{}

    // The object indicates whether or not the SRTP (Secured RTP) was enabled for
    // the voice call. The type is bool.
    CvVoIPCallActiveSRTPEnable interface{}

    // If the object has a value true(1) octet align operation is used, and if the
    // value is false(2), bandwidth efficient operation is used. This object is
    // not instantiated when the object cvVoIPCallActiveCoderTypeRate is not equal
    // to gsmAmrNb enum. The type is bool.
    CvVoIPCallActiveOctetAligned interface{}

    // This object indicates modes of Bit rates. This object is not instantiated
    // when the object cvVoIPCallActiveCoderTypeRate is not equal to gsmAmrNb
    // enum. The type is map[string]bool.
    CvVoIPCallActiveBitRates interface{}

    // The object indicates the interval (N frame-blocks) at which codec mode
    // changes are allowed. This object is not instantiated when the object
    // cvVoIPCallActiveCoderTypeRate is not equal to gsmAmrNb enum. The type is
    // interface{} with range: 1..100. Units are frame-blocks.
    CvVoIPCallActiveModeChgPeriod interface{}

    // If the object has a value of true(1), mode changes will be made to only
    // neighboring modes set to cvVoIPCallActiveBitRates object. If the value is
    // false(2), mode changes will be allowed to any modes set to
    // cvVoIPCallActiveBitRates object. This object is not instantiated when the
    // object cvVoIPCallActiveCoderTypeRate is not equal to gsmAmrNb enum. The
    // type is bool.
    CvVoIPCallActiveModeChgNeighbor interface{}

    // The object indicates the maximum amount of media that can be encapsulated
    // in a payload. Supported value is 20 msec. This object is not instantiated
    // when the object cvVoIPCallActiveCoderTypeRate is not equal to gsmAmrNb
    // enum. The type is interface{} with range: 20..100. Units are milliseconds.
    CvVoIPCallActiveMaxPtime interface{}

    // If the object has a value of true(1), frame CRC will be included in the
    // payload and if the value is false(2), frame CRC will not be included in the
    // payload. This object is applicable only when RTP frame type is octet
    // aligned. This object is not instantiated when the object
    // cvVoIPCallActiveCoderTypeRate is not equal to gsmAmrNb enum. The type is
    // bool.
    CvVoIPCallActiveCRC interface{}

    // If the object has a value of true(1), payload employs robust sorting and if
    // the value is false(2), payload does not employ robust sorting. This object
    // is applicable only when RTP frame type is octet aligned. This object is not
    // instantiated when the object cvVoIPCallActiveCoderTypeRate is not equal to
    // gsmAmrNb enum. The type is bool.
    CvVoIPCallActiveRobustSorting interface{}

    // The object indicates the RTP encapsulation type. Supported RTP
    // encapsulation type is RFC3267. This object is not instantiated when the
    // object cvVoIPCallActiveCoderTypeRate is not equal to gsmAmrNb enum. The
    // type is CvAmrNbRtpEncap.
    CvVoIPCallActiveEncap interface{}

    // The object indicates the maximum number of frame-blocks allowed in an
    // interleaving group. It indicates that frame-block level interleaving will
    // be used for that session. If this object is not set, interleaving is not
    // used. This object is applicable only when RTP frame type is octet aligned.
    // This object is not instantiated when the object
    // cvVoIPCallActiveCoderTypeRate is not equal to gsmAmrNb enum. The type is
    // interface{} with range: 1..50. Units are frame-blocks.
    CvVoIPCallActiveInterleaving interface{}

    // The object indicates the length of the time in milliseconds represented by
    // the media of the packet. Supported value is 20 milliseconds. This object is
    // not instantiated when the object cvVoIPCallActiveCoderTypeRate is not equal
    // to gsmAmrNb enum. The type is interface{} with range: 20..100. Units are
    // milliseconds.
    CvVoIPCallActivePtime interface{}

    // The object indicates the number of audio channels. Supported value is 1.
    // This object is not instantiated when the object
    // cvVoIPCallActiveCoderTypeRate is not equal to gsmAmrNb enum. The type is
    // interface{} with range: 1..6. Units are channels.
    CvVoIPCallActiveChannels interface{}

    // The object indicates the iLBC codec mode. The value of this object is valid
    // only if  cvVoIPCallActiveCoderTypeRate is equal to  'iLBC'. The type is
    // CvIlbcFrameMode. Units are milliseconds.
    CvVoIPCallActiveCoderMode interface{}

    // This object represents the call identifier for the active VoIP leg of the
    // call. The type is interface{} with range: 1..4294967295.
    CvVoIPCallActiveCallId interface{}

    // The call reference ID associates the video call entry and voice call entry
    // of the same endpoint.  An audio-only call may or may not have a valid call
    // reference ID (i.e. value greater than zero), but in both cases, there will
    // not be a video call entry associated with it.    For a video call, the
    // video-specific information  is stored in a call entry in
    // cVideoSessionActive of CISCO-VIDEO-SESSION-MIB, in which the call reference
    // ID is also identified. The type is interface{} with range: 0..4294967295.
    CvVoIPCallActiveCallReferenceId interface{}

    // This object holds the policy name. It could be media policy, DSCP policy
    // etc. The type is string.
    CcVoIPCallActivePolicyName interface{}

    // This object store the reversed direction peer address  If the address is
    // not available, then it will have a length of zero.  If the call is ingress
    // then it contains called number and if the call is egress then it contains
    // calling number. The type is string.
    CvVoIPCallActiveReversedDirectionPeerAddress interface{}

    // This object indicates the active session ID assigned by the call manager to
    // identify call legs that belong to the same call session. The type is
    // interface{} with range: 0..4294967295.
    CvVoIPCallActiveSessionId interface{}
}

func (cvVoIPCallActiveEntry *CISCOVOICEDIALCONTROLMIB_CvVoIPCallActiveTable_CvVoIPCallActiveEntry) GetEntityData() *types.CommonEntityData {
    cvVoIPCallActiveEntry.EntityData.YFilter = cvVoIPCallActiveEntry.YFilter
    cvVoIPCallActiveEntry.EntityData.YangName = "cvVoIPCallActiveEntry"
    cvVoIPCallActiveEntry.EntityData.BundleName = "cisco_ios_xe"
    cvVoIPCallActiveEntry.EntityData.ParentYangName = "cvVoIPCallActiveTable"
    cvVoIPCallActiveEntry.EntityData.SegmentPath = "cvVoIPCallActiveEntry" + types.AddKeyToken(cvVoIPCallActiveEntry.CallActiveSetupTime, "callActiveSetupTime") + types.AddKeyToken(cvVoIPCallActiveEntry.CallActiveIndex, "callActiveIndex")
    cvVoIPCallActiveEntry.EntityData.AbsolutePath = "CISCO-VOICE-DIAL-CONTROL-MIB:CISCO-VOICE-DIAL-CONTROL-MIB/cvVoIPCallActiveTable/" + cvVoIPCallActiveEntry.EntityData.SegmentPath
    cvVoIPCallActiveEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvVoIPCallActiveEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvVoIPCallActiveEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvVoIPCallActiveEntry.EntityData.Children = types.NewOrderedMap()
    cvVoIPCallActiveEntry.EntityData.Leafs = types.NewOrderedMap()
    cvVoIPCallActiveEntry.EntityData.Leafs.Append("callActiveSetupTime", types.YLeaf{"CallActiveSetupTime", cvVoIPCallActiveEntry.CallActiveSetupTime})
    cvVoIPCallActiveEntry.EntityData.Leafs.Append("callActiveIndex", types.YLeaf{"CallActiveIndex", cvVoIPCallActiveEntry.CallActiveIndex})
    cvVoIPCallActiveEntry.EntityData.Leafs.Append("cvVoIPCallActiveConnectionId", types.YLeaf{"CvVoIPCallActiveConnectionId", cvVoIPCallActiveEntry.CvVoIPCallActiveConnectionId})
    cvVoIPCallActiveEntry.EntityData.Leafs.Append("cvVoIPCallActiveRemoteIPAddress", types.YLeaf{"CvVoIPCallActiveRemoteIPAddress", cvVoIPCallActiveEntry.CvVoIPCallActiveRemoteIPAddress})
    cvVoIPCallActiveEntry.EntityData.Leafs.Append("cvVoIPCallActiveRemoteUDPPort", types.YLeaf{"CvVoIPCallActiveRemoteUDPPort", cvVoIPCallActiveEntry.CvVoIPCallActiveRemoteUDPPort})
    cvVoIPCallActiveEntry.EntityData.Leafs.Append("cvVoIPCallActiveRoundTripDelay", types.YLeaf{"CvVoIPCallActiveRoundTripDelay", cvVoIPCallActiveEntry.CvVoIPCallActiveRoundTripDelay})
    cvVoIPCallActiveEntry.EntityData.Leafs.Append("cvVoIPCallActiveSelectedQoS", types.YLeaf{"CvVoIPCallActiveSelectedQoS", cvVoIPCallActiveEntry.CvVoIPCallActiveSelectedQoS})
    cvVoIPCallActiveEntry.EntityData.Leafs.Append("cvVoIPCallActiveSessionProtocol", types.YLeaf{"CvVoIPCallActiveSessionProtocol", cvVoIPCallActiveEntry.CvVoIPCallActiveSessionProtocol})
    cvVoIPCallActiveEntry.EntityData.Leafs.Append("cvVoIPCallActiveSessionTarget", types.YLeaf{"CvVoIPCallActiveSessionTarget", cvVoIPCallActiveEntry.CvVoIPCallActiveSessionTarget})
    cvVoIPCallActiveEntry.EntityData.Leafs.Append("cvVoIPCallActiveOnTimeRvPlayout", types.YLeaf{"CvVoIPCallActiveOnTimeRvPlayout", cvVoIPCallActiveEntry.CvVoIPCallActiveOnTimeRvPlayout})
    cvVoIPCallActiveEntry.EntityData.Leafs.Append("cvVoIPCallActiveGapFillWithSilence", types.YLeaf{"CvVoIPCallActiveGapFillWithSilence", cvVoIPCallActiveEntry.CvVoIPCallActiveGapFillWithSilence})
    cvVoIPCallActiveEntry.EntityData.Leafs.Append("cvVoIPCallActiveGapFillWithPrediction", types.YLeaf{"CvVoIPCallActiveGapFillWithPrediction", cvVoIPCallActiveEntry.CvVoIPCallActiveGapFillWithPrediction})
    cvVoIPCallActiveEntry.EntityData.Leafs.Append("cvVoIPCallActiveGapFillWithInterpolation", types.YLeaf{"CvVoIPCallActiveGapFillWithInterpolation", cvVoIPCallActiveEntry.CvVoIPCallActiveGapFillWithInterpolation})
    cvVoIPCallActiveEntry.EntityData.Leafs.Append("cvVoIPCallActiveGapFillWithRedundancy", types.YLeaf{"CvVoIPCallActiveGapFillWithRedundancy", cvVoIPCallActiveEntry.CvVoIPCallActiveGapFillWithRedundancy})
    cvVoIPCallActiveEntry.EntityData.Leafs.Append("cvVoIPCallActiveHiWaterPlayoutDelay", types.YLeaf{"CvVoIPCallActiveHiWaterPlayoutDelay", cvVoIPCallActiveEntry.CvVoIPCallActiveHiWaterPlayoutDelay})
    cvVoIPCallActiveEntry.EntityData.Leafs.Append("cvVoIPCallActiveLoWaterPlayoutDelay", types.YLeaf{"CvVoIPCallActiveLoWaterPlayoutDelay", cvVoIPCallActiveEntry.CvVoIPCallActiveLoWaterPlayoutDelay})
    cvVoIPCallActiveEntry.EntityData.Leafs.Append("cvVoIPCallActiveReceiveDelay", types.YLeaf{"CvVoIPCallActiveReceiveDelay", cvVoIPCallActiveEntry.CvVoIPCallActiveReceiveDelay})
    cvVoIPCallActiveEntry.EntityData.Leafs.Append("cvVoIPCallActiveVADEnable", types.YLeaf{"CvVoIPCallActiveVADEnable", cvVoIPCallActiveEntry.CvVoIPCallActiveVADEnable})
    cvVoIPCallActiveEntry.EntityData.Leafs.Append("cvVoIPCallActiveCoderTypeRate", types.YLeaf{"CvVoIPCallActiveCoderTypeRate", cvVoIPCallActiveEntry.CvVoIPCallActiveCoderTypeRate})
    cvVoIPCallActiveEntry.EntityData.Leafs.Append("cvVoIPCallActiveLostPackets", types.YLeaf{"CvVoIPCallActiveLostPackets", cvVoIPCallActiveEntry.CvVoIPCallActiveLostPackets})
    cvVoIPCallActiveEntry.EntityData.Leafs.Append("cvVoIPCallActiveEarlyPackets", types.YLeaf{"CvVoIPCallActiveEarlyPackets", cvVoIPCallActiveEntry.CvVoIPCallActiveEarlyPackets})
    cvVoIPCallActiveEntry.EntityData.Leafs.Append("cvVoIPCallActiveLatePackets", types.YLeaf{"CvVoIPCallActiveLatePackets", cvVoIPCallActiveEntry.CvVoIPCallActiveLatePackets})
    cvVoIPCallActiveEntry.EntityData.Leafs.Append("cvVoIPCallActiveUsername", types.YLeaf{"CvVoIPCallActiveUsername", cvVoIPCallActiveEntry.CvVoIPCallActiveUsername})
    cvVoIPCallActiveEntry.EntityData.Leafs.Append("cvVoIPCallActiveProtocolCallId", types.YLeaf{"CvVoIPCallActiveProtocolCallId", cvVoIPCallActiveEntry.CvVoIPCallActiveProtocolCallId})
    cvVoIPCallActiveEntry.EntityData.Leafs.Append("cvVoIPCallActiveRemSigIPAddrT", types.YLeaf{"CvVoIPCallActiveRemSigIPAddrT", cvVoIPCallActiveEntry.CvVoIPCallActiveRemSigIPAddrT})
    cvVoIPCallActiveEntry.EntityData.Leafs.Append("cvVoIPCallActiveRemSigIPAddr", types.YLeaf{"CvVoIPCallActiveRemSigIPAddr", cvVoIPCallActiveEntry.CvVoIPCallActiveRemSigIPAddr})
    cvVoIPCallActiveEntry.EntityData.Leafs.Append("cvVoIPCallActiveRemSigPort", types.YLeaf{"CvVoIPCallActiveRemSigPort", cvVoIPCallActiveEntry.CvVoIPCallActiveRemSigPort})
    cvVoIPCallActiveEntry.EntityData.Leafs.Append("cvVoIPCallActiveRemMediaIPAddrT", types.YLeaf{"CvVoIPCallActiveRemMediaIPAddrT", cvVoIPCallActiveEntry.CvVoIPCallActiveRemMediaIPAddrT})
    cvVoIPCallActiveEntry.EntityData.Leafs.Append("cvVoIPCallActiveRemMediaIPAddr", types.YLeaf{"CvVoIPCallActiveRemMediaIPAddr", cvVoIPCallActiveEntry.CvVoIPCallActiveRemMediaIPAddr})
    cvVoIPCallActiveEntry.EntityData.Leafs.Append("cvVoIPCallActiveRemMediaPort", types.YLeaf{"CvVoIPCallActiveRemMediaPort", cvVoIPCallActiveEntry.CvVoIPCallActiveRemMediaPort})
    cvVoIPCallActiveEntry.EntityData.Leafs.Append("cvVoIPCallActiveSRTPEnable", types.YLeaf{"CvVoIPCallActiveSRTPEnable", cvVoIPCallActiveEntry.CvVoIPCallActiveSRTPEnable})
    cvVoIPCallActiveEntry.EntityData.Leafs.Append("cvVoIPCallActiveOctetAligned", types.YLeaf{"CvVoIPCallActiveOctetAligned", cvVoIPCallActiveEntry.CvVoIPCallActiveOctetAligned})
    cvVoIPCallActiveEntry.EntityData.Leafs.Append("cvVoIPCallActiveBitRates", types.YLeaf{"CvVoIPCallActiveBitRates", cvVoIPCallActiveEntry.CvVoIPCallActiveBitRates})
    cvVoIPCallActiveEntry.EntityData.Leafs.Append("cvVoIPCallActiveModeChgPeriod", types.YLeaf{"CvVoIPCallActiveModeChgPeriod", cvVoIPCallActiveEntry.CvVoIPCallActiveModeChgPeriod})
    cvVoIPCallActiveEntry.EntityData.Leafs.Append("cvVoIPCallActiveModeChgNeighbor", types.YLeaf{"CvVoIPCallActiveModeChgNeighbor", cvVoIPCallActiveEntry.CvVoIPCallActiveModeChgNeighbor})
    cvVoIPCallActiveEntry.EntityData.Leafs.Append("cvVoIPCallActiveMaxPtime", types.YLeaf{"CvVoIPCallActiveMaxPtime", cvVoIPCallActiveEntry.CvVoIPCallActiveMaxPtime})
    cvVoIPCallActiveEntry.EntityData.Leafs.Append("cvVoIPCallActiveCRC", types.YLeaf{"CvVoIPCallActiveCRC", cvVoIPCallActiveEntry.CvVoIPCallActiveCRC})
    cvVoIPCallActiveEntry.EntityData.Leafs.Append("cvVoIPCallActiveRobustSorting", types.YLeaf{"CvVoIPCallActiveRobustSorting", cvVoIPCallActiveEntry.CvVoIPCallActiveRobustSorting})
    cvVoIPCallActiveEntry.EntityData.Leafs.Append("cvVoIPCallActiveEncap", types.YLeaf{"CvVoIPCallActiveEncap", cvVoIPCallActiveEntry.CvVoIPCallActiveEncap})
    cvVoIPCallActiveEntry.EntityData.Leafs.Append("cvVoIPCallActiveInterleaving", types.YLeaf{"CvVoIPCallActiveInterleaving", cvVoIPCallActiveEntry.CvVoIPCallActiveInterleaving})
    cvVoIPCallActiveEntry.EntityData.Leafs.Append("cvVoIPCallActivePtime", types.YLeaf{"CvVoIPCallActivePtime", cvVoIPCallActiveEntry.CvVoIPCallActivePtime})
    cvVoIPCallActiveEntry.EntityData.Leafs.Append("cvVoIPCallActiveChannels", types.YLeaf{"CvVoIPCallActiveChannels", cvVoIPCallActiveEntry.CvVoIPCallActiveChannels})
    cvVoIPCallActiveEntry.EntityData.Leafs.Append("cvVoIPCallActiveCoderMode", types.YLeaf{"CvVoIPCallActiveCoderMode", cvVoIPCallActiveEntry.CvVoIPCallActiveCoderMode})
    cvVoIPCallActiveEntry.EntityData.Leafs.Append("cvVoIPCallActiveCallId", types.YLeaf{"CvVoIPCallActiveCallId", cvVoIPCallActiveEntry.CvVoIPCallActiveCallId})
    cvVoIPCallActiveEntry.EntityData.Leafs.Append("cvVoIPCallActiveCallReferenceId", types.YLeaf{"CvVoIPCallActiveCallReferenceId", cvVoIPCallActiveEntry.CvVoIPCallActiveCallReferenceId})
    cvVoIPCallActiveEntry.EntityData.Leafs.Append("ccVoIPCallActivePolicyName", types.YLeaf{"CcVoIPCallActivePolicyName", cvVoIPCallActiveEntry.CcVoIPCallActivePolicyName})
    cvVoIPCallActiveEntry.EntityData.Leafs.Append("cvVoIPCallActiveReversedDirectionPeerAddress", types.YLeaf{"CvVoIPCallActiveReversedDirectionPeerAddress", cvVoIPCallActiveEntry.CvVoIPCallActiveReversedDirectionPeerAddress})
    cvVoIPCallActiveEntry.EntityData.Leafs.Append("cvVoIPCallActiveSessionId", types.YLeaf{"CvVoIPCallActiveSessionId", cvVoIPCallActiveEntry.CvVoIPCallActiveSessionId})

    cvVoIPCallActiveEntry.EntityData.YListKeys = []string {"CallActiveSetupTime", "CallActiveIndex"}

    return &(cvVoIPCallActiveEntry.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_CvCallVolConnTable
// This table represents the number of active
// call connections for each call connection type
// in the voice gateway.
type CISCOVOICEDIALCONTROLMIB_CvCallVolConnTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the cvCallVolConnTable indicates number of active calls for a
    // call connection type in the voice gateway. The type is slice of
    // CISCOVOICEDIALCONTROLMIB_CvCallVolConnTable_CvCallVolConnEntry.
    CvCallVolConnEntry []*CISCOVOICEDIALCONTROLMIB_CvCallVolConnTable_CvCallVolConnEntry
}

func (cvCallVolConnTable *CISCOVOICEDIALCONTROLMIB_CvCallVolConnTable) GetEntityData() *types.CommonEntityData {
    cvCallVolConnTable.EntityData.YFilter = cvCallVolConnTable.YFilter
    cvCallVolConnTable.EntityData.YangName = "cvCallVolConnTable"
    cvCallVolConnTable.EntityData.BundleName = "cisco_ios_xe"
    cvCallVolConnTable.EntityData.ParentYangName = "CISCO-VOICE-DIAL-CONTROL-MIB"
    cvCallVolConnTable.EntityData.SegmentPath = "cvCallVolConnTable"
    cvCallVolConnTable.EntityData.AbsolutePath = "CISCO-VOICE-DIAL-CONTROL-MIB:CISCO-VOICE-DIAL-CONTROL-MIB/" + cvCallVolConnTable.EntityData.SegmentPath
    cvCallVolConnTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvCallVolConnTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvCallVolConnTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvCallVolConnTable.EntityData.Children = types.NewOrderedMap()
    cvCallVolConnTable.EntityData.Children.Append("cvCallVolConnEntry", types.YChild{"CvCallVolConnEntry", nil})
    for i := range cvCallVolConnTable.CvCallVolConnEntry {
        cvCallVolConnTable.EntityData.Children.Append(types.GetSegmentPath(cvCallVolConnTable.CvCallVolConnEntry[i]), types.YChild{"CvCallVolConnEntry", cvCallVolConnTable.CvCallVolConnEntry[i]})
    }
    cvCallVolConnTable.EntityData.Leafs = types.NewOrderedMap()

    cvCallVolConnTable.EntityData.YListKeys = []string {}

    return &(cvCallVolConnTable.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_CvCallVolConnTable_CvCallVolConnEntry
// An entry in the cvCallVolConnTable indicates
// number of active calls for a call connection type
// in the voice gateway.
type CISCOVOICEDIALCONTROLMIB_CvCallVolConnTable_CvCallVolConnEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. This object represents index to the
    // cvCallVolConnTable. The type is CvCallConnectionType.
    CvCallVolConnIndex interface{}

    // This object represents the total number of active calls for a connection
    // type  in the voice gateway. The type is interface{} with range: 0..65535.
    CvCallVolConnActiveConnection interface{}
}

func (cvCallVolConnEntry *CISCOVOICEDIALCONTROLMIB_CvCallVolConnTable_CvCallVolConnEntry) GetEntityData() *types.CommonEntityData {
    cvCallVolConnEntry.EntityData.YFilter = cvCallVolConnEntry.YFilter
    cvCallVolConnEntry.EntityData.YangName = "cvCallVolConnEntry"
    cvCallVolConnEntry.EntityData.BundleName = "cisco_ios_xe"
    cvCallVolConnEntry.EntityData.ParentYangName = "cvCallVolConnTable"
    cvCallVolConnEntry.EntityData.SegmentPath = "cvCallVolConnEntry" + types.AddKeyToken(cvCallVolConnEntry.CvCallVolConnIndex, "cvCallVolConnIndex")
    cvCallVolConnEntry.EntityData.AbsolutePath = "CISCO-VOICE-DIAL-CONTROL-MIB:CISCO-VOICE-DIAL-CONTROL-MIB/cvCallVolConnTable/" + cvCallVolConnEntry.EntityData.SegmentPath
    cvCallVolConnEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvCallVolConnEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvCallVolConnEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvCallVolConnEntry.EntityData.Children = types.NewOrderedMap()
    cvCallVolConnEntry.EntityData.Leafs = types.NewOrderedMap()
    cvCallVolConnEntry.EntityData.Leafs.Append("cvCallVolConnIndex", types.YLeaf{"CvCallVolConnIndex", cvCallVolConnEntry.CvCallVolConnIndex})
    cvCallVolConnEntry.EntityData.Leafs.Append("cvCallVolConnActiveConnection", types.YLeaf{"CvCallVolConnActiveConnection", cvCallVolConnEntry.CvCallVolConnActiveConnection})

    cvCallVolConnEntry.EntityData.YListKeys = []string {"CvCallVolConnIndex"}

    return &(cvCallVolConnEntry.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_CvCallVolIfTable
// This table represents the information about
// the usage of an IP interface in a voice
// gateway for voice media calls. This table
// has a sparse-dependent relationship with  
// ifTable. There exists an entry in this table, 
// for each of the  entries in ifTable where ifType 
// is one of 'ethernetCsmacd' and 'softwareLoopback'.
type CISCOVOICEDIALCONTROLMIB_CvCallVolIfTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry represents a row in cvCallVolIfTable and corresponds to the
    // information about an IP  interface in the voice gateway. The type is slice
    // of CISCOVOICEDIALCONTROLMIB_CvCallVolIfTable_CvCallVolIfEntry.
    CvCallVolIfEntry []*CISCOVOICEDIALCONTROLMIB_CvCallVolIfTable_CvCallVolIfEntry
}

func (cvCallVolIfTable *CISCOVOICEDIALCONTROLMIB_CvCallVolIfTable) GetEntityData() *types.CommonEntityData {
    cvCallVolIfTable.EntityData.YFilter = cvCallVolIfTable.YFilter
    cvCallVolIfTable.EntityData.YangName = "cvCallVolIfTable"
    cvCallVolIfTable.EntityData.BundleName = "cisco_ios_xe"
    cvCallVolIfTable.EntityData.ParentYangName = "CISCO-VOICE-DIAL-CONTROL-MIB"
    cvCallVolIfTable.EntityData.SegmentPath = "cvCallVolIfTable"
    cvCallVolIfTable.EntityData.AbsolutePath = "CISCO-VOICE-DIAL-CONTROL-MIB:CISCO-VOICE-DIAL-CONTROL-MIB/" + cvCallVolIfTable.EntityData.SegmentPath
    cvCallVolIfTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvCallVolIfTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvCallVolIfTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvCallVolIfTable.EntityData.Children = types.NewOrderedMap()
    cvCallVolIfTable.EntityData.Children.Append("cvCallVolIfEntry", types.YChild{"CvCallVolIfEntry", nil})
    for i := range cvCallVolIfTable.CvCallVolIfEntry {
        cvCallVolIfTable.EntityData.Children.Append(types.GetSegmentPath(cvCallVolIfTable.CvCallVolIfEntry[i]), types.YChild{"CvCallVolIfEntry", cvCallVolIfTable.CvCallVolIfEntry[i]})
    }
    cvCallVolIfTable.EntityData.Leafs = types.NewOrderedMap()

    cvCallVolIfTable.EntityData.YListKeys = []string {}

    return &(cvCallVolIfTable.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_CvCallVolIfTable_CvCallVolIfEntry
// Each entry represents a row in cvCallVolIfTable
// and corresponds to the information about an IP 
// interface in the voice gateway.
type CISCOVOICEDIALCONTROLMIB_CvCallVolIfTable_CvCallVolIfEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This object represents the total number of inbound active media calls
    // through this IP  interface. The type is interface{} with range: 0..65535.
    CvCallVolMediaIncomingCalls interface{}

    // This object represents the total number of outbound active media calls
    // through the IP  interface. The type is interface{} with range: 0..65535.
    CvCallVolMediaOutgoingCalls interface{}
}

func (cvCallVolIfEntry *CISCOVOICEDIALCONTROLMIB_CvCallVolIfTable_CvCallVolIfEntry) GetEntityData() *types.CommonEntityData {
    cvCallVolIfEntry.EntityData.YFilter = cvCallVolIfEntry.YFilter
    cvCallVolIfEntry.EntityData.YangName = "cvCallVolIfEntry"
    cvCallVolIfEntry.EntityData.BundleName = "cisco_ios_xe"
    cvCallVolIfEntry.EntityData.ParentYangName = "cvCallVolIfTable"
    cvCallVolIfEntry.EntityData.SegmentPath = "cvCallVolIfEntry" + types.AddKeyToken(cvCallVolIfEntry.IfIndex, "ifIndex")
    cvCallVolIfEntry.EntityData.AbsolutePath = "CISCO-VOICE-DIAL-CONTROL-MIB:CISCO-VOICE-DIAL-CONTROL-MIB/cvCallVolIfTable/" + cvCallVolIfEntry.EntityData.SegmentPath
    cvCallVolIfEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvCallVolIfEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvCallVolIfEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvCallVolIfEntry.EntityData.Children = types.NewOrderedMap()
    cvCallVolIfEntry.EntityData.Leafs = types.NewOrderedMap()
    cvCallVolIfEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", cvCallVolIfEntry.IfIndex})
    cvCallVolIfEntry.EntityData.Leafs.Append("cvCallVolMediaIncomingCalls", types.YLeaf{"CvCallVolMediaIncomingCalls", cvCallVolIfEntry.CvCallVolMediaIncomingCalls})
    cvCallVolIfEntry.EntityData.Leafs.Append("cvCallVolMediaOutgoingCalls", types.YLeaf{"CvCallVolMediaOutgoingCalls", cvCallVolIfEntry.CvCallVolMediaOutgoingCalls})

    cvCallVolIfEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(cvCallVolIfEntry.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_CvCallHistoryTable
// This table is the voice extension to the call history table
// of IETF Dial Control MIB. It contains voice encapsulation
// call leg information such as voice packet statistics,
// coder usage and end to end bandwidth of the call leg.
type CISCOVOICEDIALCONTROLMIB_CvCallHistoryTable struct {
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
    // CISCOVOICEDIALCONTROLMIB_CvCallHistoryTable_CvCallHistoryEntry.
    CvCallHistoryEntry []*CISCOVOICEDIALCONTROLMIB_CvCallHistoryTable_CvCallHistoryEntry
}

func (cvCallHistoryTable *CISCOVOICEDIALCONTROLMIB_CvCallHistoryTable) GetEntityData() *types.CommonEntityData {
    cvCallHistoryTable.EntityData.YFilter = cvCallHistoryTable.YFilter
    cvCallHistoryTable.EntityData.YangName = "cvCallHistoryTable"
    cvCallHistoryTable.EntityData.BundleName = "cisco_ios_xe"
    cvCallHistoryTable.EntityData.ParentYangName = "CISCO-VOICE-DIAL-CONTROL-MIB"
    cvCallHistoryTable.EntityData.SegmentPath = "cvCallHistoryTable"
    cvCallHistoryTable.EntityData.AbsolutePath = "CISCO-VOICE-DIAL-CONTROL-MIB:CISCO-VOICE-DIAL-CONTROL-MIB/" + cvCallHistoryTable.EntityData.SegmentPath
    cvCallHistoryTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvCallHistoryTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvCallHistoryTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvCallHistoryTable.EntityData.Children = types.NewOrderedMap()
    cvCallHistoryTable.EntityData.Children.Append("cvCallHistoryEntry", types.YChild{"CvCallHistoryEntry", nil})
    for i := range cvCallHistoryTable.CvCallHistoryEntry {
        cvCallHistoryTable.EntityData.Children.Append(types.GetSegmentPath(cvCallHistoryTable.CvCallHistoryEntry[i]), types.YChild{"CvCallHistoryEntry", cvCallHistoryTable.CvCallHistoryEntry[i]})
    }
    cvCallHistoryTable.EntityData.Leafs = types.NewOrderedMap()

    cvCallHistoryTable.EntityData.YListKeys = []string {}

    return &(cvCallHistoryTable.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_CvCallHistoryTable_CvCallHistoryEntry
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
type CISCOVOICEDIALCONTROLMIB_CvCallHistoryTable_CvCallHistoryEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // cisco_dial_control_mib.CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryIndex
    CCallHistoryIndex interface{}

    // The global connection identifier for the telephony leg, which was assigned
    // to the call. The type is string with length: 0..16.
    CvCallHistoryConnectionId interface{}

    // Duration of Transmit path open from this peer to the voice gateway for the
    // call leg. This counter object will lock at the maximum value which is
    // approximately two days. The type is interface{} with range: 0..4294967295.
    // Units are milliseconds.
    CvCallHistoryTxDuration interface{}

    // Duration for this call leg. The Voice Utilization Rate can be obtained by
    // dividing this by cvCallHistoryTXDuration object. This counter object will
    // lock at the maximum value which is approximately two days. The type is
    // interface{} with range: 0..4294967295. Units are milliseconds.
    CvCallHistoryVoiceTxDuration interface{}

    // Duration of fax transmitted from this peer to voice gateway for this call
    // leg. The FAX Utilization Rate can be obtained by dividing this by
    // cvCallHistoryTXDuration object. This counter object will lock at the
    // maximum value which is approximately two days. The type is interface{} with
    // range: 0..4294967295. Units are milliseconds.
    CvCallHistoryFaxTxDuration interface{}

    // The negotiated coder rate. It specifies the transmit rate of voice/fax
    // compression to its associated call leg for the call. The type is
    // CvcCoderTypeRate.
    CvCallHistoryCoderTypeRate interface{}

    // The object contains the average noise level for the call leg. The type is
    // interface{} with range: -128..8. Units are dBm.
    CvCallHistoryNoiseLevel interface{}

    // The object contains the average ACOM level for the call leg. The value -1
    // indicates the level can not be determined or level detection is disabled.
    // The type is interface{} with range: -1..127. Units are dB.
    CvCallHistoryACOMLevel interface{}

    // The object specifies the session target of the peer that is used for the
    // call leg via telephony interface. The type is string with length: 0..64.
    CvCallHistorySessionTarget interface{}

    // The number of FAX related image pages are received or transmitted via the
    // peer for the call leg. The type is interface{} with range: 0..4294967295.
    // Units are pages.
    CvCallHistoryImgPageCount interface{}

    // The calling party name of the call. If the name is not available, then it
    // will have a length of zero. The type is string.
    CvCallHistoryCallingName interface{}

    // The object indicates whether or not the caller ID feature is blocked for
    // this call. The type is bool.
    CvCallHistoryCallerIDBlock interface{}

    // The object indicates the account code input to the call. It could be used
    // by down stream billing server. The value of empty string indicates no
    // account code input. The type is string with length: 0..50.
    CvCallHistoryAccountCode interface{}

    // This object represents the call identifier for the telephony leg, which was
    // assigned to the call. The type is interface{} with range: 1..4294967295.
    CvCallHistoryCallId interface{}
}

func (cvCallHistoryEntry *CISCOVOICEDIALCONTROLMIB_CvCallHistoryTable_CvCallHistoryEntry) GetEntityData() *types.CommonEntityData {
    cvCallHistoryEntry.EntityData.YFilter = cvCallHistoryEntry.YFilter
    cvCallHistoryEntry.EntityData.YangName = "cvCallHistoryEntry"
    cvCallHistoryEntry.EntityData.BundleName = "cisco_ios_xe"
    cvCallHistoryEntry.EntityData.ParentYangName = "cvCallHistoryTable"
    cvCallHistoryEntry.EntityData.SegmentPath = "cvCallHistoryEntry" + types.AddKeyToken(cvCallHistoryEntry.CCallHistoryIndex, "cCallHistoryIndex")
    cvCallHistoryEntry.EntityData.AbsolutePath = "CISCO-VOICE-DIAL-CONTROL-MIB:CISCO-VOICE-DIAL-CONTROL-MIB/cvCallHistoryTable/" + cvCallHistoryEntry.EntityData.SegmentPath
    cvCallHistoryEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvCallHistoryEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvCallHistoryEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvCallHistoryEntry.EntityData.Children = types.NewOrderedMap()
    cvCallHistoryEntry.EntityData.Leafs = types.NewOrderedMap()
    cvCallHistoryEntry.EntityData.Leafs.Append("cCallHistoryIndex", types.YLeaf{"CCallHistoryIndex", cvCallHistoryEntry.CCallHistoryIndex})
    cvCallHistoryEntry.EntityData.Leafs.Append("cvCallHistoryConnectionId", types.YLeaf{"CvCallHistoryConnectionId", cvCallHistoryEntry.CvCallHistoryConnectionId})
    cvCallHistoryEntry.EntityData.Leafs.Append("cvCallHistoryTxDuration", types.YLeaf{"CvCallHistoryTxDuration", cvCallHistoryEntry.CvCallHistoryTxDuration})
    cvCallHistoryEntry.EntityData.Leafs.Append("cvCallHistoryVoiceTxDuration", types.YLeaf{"CvCallHistoryVoiceTxDuration", cvCallHistoryEntry.CvCallHistoryVoiceTxDuration})
    cvCallHistoryEntry.EntityData.Leafs.Append("cvCallHistoryFaxTxDuration", types.YLeaf{"CvCallHistoryFaxTxDuration", cvCallHistoryEntry.CvCallHistoryFaxTxDuration})
    cvCallHistoryEntry.EntityData.Leafs.Append("cvCallHistoryCoderTypeRate", types.YLeaf{"CvCallHistoryCoderTypeRate", cvCallHistoryEntry.CvCallHistoryCoderTypeRate})
    cvCallHistoryEntry.EntityData.Leafs.Append("cvCallHistoryNoiseLevel", types.YLeaf{"CvCallHistoryNoiseLevel", cvCallHistoryEntry.CvCallHistoryNoiseLevel})
    cvCallHistoryEntry.EntityData.Leafs.Append("cvCallHistoryACOMLevel", types.YLeaf{"CvCallHistoryACOMLevel", cvCallHistoryEntry.CvCallHistoryACOMLevel})
    cvCallHistoryEntry.EntityData.Leafs.Append("cvCallHistorySessionTarget", types.YLeaf{"CvCallHistorySessionTarget", cvCallHistoryEntry.CvCallHistorySessionTarget})
    cvCallHistoryEntry.EntityData.Leafs.Append("cvCallHistoryImgPageCount", types.YLeaf{"CvCallHistoryImgPageCount", cvCallHistoryEntry.CvCallHistoryImgPageCount})
    cvCallHistoryEntry.EntityData.Leafs.Append("cvCallHistoryCallingName", types.YLeaf{"CvCallHistoryCallingName", cvCallHistoryEntry.CvCallHistoryCallingName})
    cvCallHistoryEntry.EntityData.Leafs.Append("cvCallHistoryCallerIDBlock", types.YLeaf{"CvCallHistoryCallerIDBlock", cvCallHistoryEntry.CvCallHistoryCallerIDBlock})
    cvCallHistoryEntry.EntityData.Leafs.Append("cvCallHistoryAccountCode", types.YLeaf{"CvCallHistoryAccountCode", cvCallHistoryEntry.CvCallHistoryAccountCode})
    cvCallHistoryEntry.EntityData.Leafs.Append("cvCallHistoryCallId", types.YLeaf{"CvCallHistoryCallId", cvCallHistoryEntry.CvCallHistoryCallId})

    cvCallHistoryEntry.EntityData.YListKeys = []string {"CCallHistoryIndex"}

    return &(cvCallHistoryEntry.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_CvVoIPCallHistoryTable
// This table is the VoIP extension to the call history table
// of IETF Dial Control MIB. It contains VoIP call leg
// information about specific VoIP call destination and the
// selected QoS for the call leg.
type CISCOVOICEDIALCONTROLMIB_CvVoIPCallHistoryTable struct {
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
    // CISCOVOICEDIALCONTROLMIB_CvVoIPCallHistoryTable_CvVoIPCallHistoryEntry.
    CvVoIPCallHistoryEntry []*CISCOVOICEDIALCONTROLMIB_CvVoIPCallHistoryTable_CvVoIPCallHistoryEntry
}

func (cvVoIPCallHistoryTable *CISCOVOICEDIALCONTROLMIB_CvVoIPCallHistoryTable) GetEntityData() *types.CommonEntityData {
    cvVoIPCallHistoryTable.EntityData.YFilter = cvVoIPCallHistoryTable.YFilter
    cvVoIPCallHistoryTable.EntityData.YangName = "cvVoIPCallHistoryTable"
    cvVoIPCallHistoryTable.EntityData.BundleName = "cisco_ios_xe"
    cvVoIPCallHistoryTable.EntityData.ParentYangName = "CISCO-VOICE-DIAL-CONTROL-MIB"
    cvVoIPCallHistoryTable.EntityData.SegmentPath = "cvVoIPCallHistoryTable"
    cvVoIPCallHistoryTable.EntityData.AbsolutePath = "CISCO-VOICE-DIAL-CONTROL-MIB:CISCO-VOICE-DIAL-CONTROL-MIB/" + cvVoIPCallHistoryTable.EntityData.SegmentPath
    cvVoIPCallHistoryTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvVoIPCallHistoryTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvVoIPCallHistoryTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvVoIPCallHistoryTable.EntityData.Children = types.NewOrderedMap()
    cvVoIPCallHistoryTable.EntityData.Children.Append("cvVoIPCallHistoryEntry", types.YChild{"CvVoIPCallHistoryEntry", nil})
    for i := range cvVoIPCallHistoryTable.CvVoIPCallHistoryEntry {
        cvVoIPCallHistoryTable.EntityData.Children.Append(types.GetSegmentPath(cvVoIPCallHistoryTable.CvVoIPCallHistoryEntry[i]), types.YChild{"CvVoIPCallHistoryEntry", cvVoIPCallHistoryTable.CvVoIPCallHistoryEntry[i]})
    }
    cvVoIPCallHistoryTable.EntityData.Leafs = types.NewOrderedMap()

    cvVoIPCallHistoryTable.EntityData.YListKeys = []string {}

    return &(cvVoIPCallHistoryTable.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_CvVoIPCallHistoryTable_CvVoIPCallHistoryEntry
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
type CISCOVOICEDIALCONTROLMIB_CvVoIPCallHistoryTable_CvVoIPCallHistoryEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // cisco_dial_control_mib.CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryIndex
    CCallHistoryIndex interface{}

    // The global connection identifier for the VoIP leg, which was assigned to
    // the call. The type is string with length: 0..16.
    CvVoIPCallHistoryConnectionId interface{}

    // Remote system IP address for the call. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    CvVoIPCallHistoryRemoteIPAddress interface{}

    // Remote system UDP listener port to which to transmit voice packets for the
    // call. The type is interface{} with range: 0..65535.
    CvVoIPCallHistoryRemoteUDPPort interface{}

    // The voice packet round trip delay between local and the remote system on
    // the IP backbone during the call. The type is interface{} with range:
    // 0..4294967295. Units are milliseconds.
    CvVoIPCallHistoryRoundTripDelay interface{}

    // The selected RSVP QoS for the call. The type is QosService.
    CvVoIPCallHistorySelectedQoS interface{}

    // The object specifies the session protocol to be used for Internet call
    // between local and remote router via IP backbone. The type is
    // CvSessionProtocol.
    CvVoIPCallHistorySessionProtocol interface{}

    // The object specifies the session target of the peer that is used for the
    // Voice over IP call. The type is string.
    CvVoIPCallHistorySessionTarget interface{}

    // Duration of voice playout from data received on time for this call. This
    // plus the durations for the GapFills in the following entries gives the
    // Total Voice Playout Duration for Active Voice. This does not include
    // duration for which no data was sent by the Transmit end as voice signal,
    // e.g., silence suppression and fax signal. The On Time Playout Rate can be
    // computed by dividing this entry by the Total Voice Playout Duration. This
    // counter object will lock at the maximum value which is approximately two
    // days. The type is interface{} with range: 0..4294967295. Units are
    // milliseconds.
    CvVoIPCallHistoryOnTimeRvPlayout interface{}

    // Duration of voice signal replaced with signal played out during silence due
    // to voice data not received on time (or lost) from voice gateway this call.
    // This counter object will lock at the maximum value which is approximately
    // two days. The type is interface{} with range: 0..4294967295. Units are
    // milliseconds.
    CvVoIPCallHistoryGapFillWithSilence interface{}

    // Duration of voice signal played out with signal synthesized from parameters
    // or samples of data preceding in time due to voice data not received on time
    // (or lost) from voice gateway for this call. An example of such playout is
    // frame-erasure or  frame-concealment strategies in G.729 and G.723.1
    // compression algorithms. This counter object will lock at the maximum value
    // which is approximately two days. The type is interface{} with range:
    // 0..4294967295. Units are milliseconds.
    CvVoIPCallHistoryGapFillWithPrediction interface{}

    // Duration of voice signal played out with signal synthesized from parameters
    // or samples of data preceding and following in time due to voice data not
    // received on time (or lost) from voice gateway for this call. This counter
    // object will lock at the maximum value which is approximately two days. The
    // type is interface{} with range: 0..4294967295. Units are milliseconds.
    CvVoIPCallHistoryGapFillWithInterpolation interface{}

    // Duration of voice signal played out with signal synthesized from redundancy
    // parameters available due to voice data not received on time (or lost) from
    // voice gateway for this call. This counter object will lock at the maximum
    // value which is approximately two days. The type is interface{} with range:
    // 0..4294967295. Units are milliseconds.
    CvVoIPCallHistoryGapFillWithRedundancy interface{}

    // The high water mark Voice Playout FIFO Delay during the voice call. This
    // counter object will lock at the maximum value which is approximately two
    // days. The type is interface{} with range: 0..4294967295. Units are
    // milliseconds.
    CvVoIPCallHistoryHiWaterPlayoutDelay interface{}

    // The low water mark Voice Playout FIFO Delay during the voice call. The type
    // is interface{} with range: 0..4294967295. Units are milliseconds.
    CvVoIPCallHistoryLoWaterPlayoutDelay interface{}

    // The average Playout FIFO Delay plus the decoder delay during the voice
    // call. The type is interface{} with range: 0..4294967295.
    CvVoIPCallHistoryReceiveDelay interface{}

    // The object indicates whether or not the VAD (Voice Activity Detection) was
    // enabled for the voice call. The type is bool.
    CvVoIPCallHistoryVADEnable interface{}

    // The negotiated coder rate. It specifies the transmit rate of voice/fax
    // compression to its associated call leg for the call. This rate is different
    // from the configuration rate because of rate negotiation during the call.
    // The type is CvcCoderTypeRate.
    CvVoIPCallHistoryCoderTypeRate interface{}

    // The Calculated Planning Impairment Factor (Icpif) of the call  that is
    // associated to this call leg. The value in this object is computed by the
    // following equation. Icpif of the call = Itotal (total impairment value) of
    // the call - A (Expectation Factor) in the cvVoIPPeerCfgExpectFactor of the
    // call leg associated peer. A value of -1 implies that Icpif was not
    // calculated and is meaningless for this call. The type is interface{} with
    // range: -1..55. Units are equipment impairment factor (eif).
    CvVoIPCallHistoryIcpif interface{}

    // The number of lost voice packets during the call. The type is interface{}
    // with range: 0..4294967295. Units are packets.
    CvVoIPCallHistoryLostPackets interface{}

    // The number of received voice packets that are arrived too early to store in
    // jitter buffer during the call. The type is interface{} with range:
    // 0..4294967295. Units are packets.
    CvVoIPCallHistoryEarlyPackets interface{}

    // The number of received voice packets that are arrived too late to playout
    // with CODEC (Coder/Decoder) during the call. The type is interface{} with
    // range: 0..4294967295. Units are packets.
    CvVoIPCallHistoryLatePackets interface{}

    // The textual identifier of the calling party (user) of the call. If the
    // username is not available, then the value of this object will have a length
    // of zero. The type is string with length: 0..16.
    CvVoIPCallHistoryUsername interface{}

    // The protocol-specific call identifier for the VoIP call. The type is string
    // with length: 0..255.
    CvVoIPCallHistoryProtocolCallId interface{}

    // This object specifies the type of address contained in the associated
    // instance of cvVoIPCallHistoryRemSigIPAddr. The type is InetAddressType.
    CvVoIPCallHistoryRemSigIPAddrT interface{}

    // Remote signalling IP address for the VoIP call. The type is string with
    // length: 0..255.
    CvVoIPCallHistoryRemSigIPAddr interface{}

    // Remote signalling listener port to which to transmit voice packets. The
    // type is interface{} with range: 0..65535.
    CvVoIPCallHistoryRemSigPort interface{}

    // This object specifies the type of address contained in the associated
    // instance of cvVoIPCallHistoryRemMediaIPAddr. The type is InetAddressType.
    CvVoIPCallHistoryRemMediaIPAddrT interface{}

    // Remote media end point IP address for the VoIP call. The type is string
    // with length: 0..255.
    CvVoIPCallHistoryRemMediaIPAddr interface{}

    // Remote media end point listener port to which to transmit voice packets.
    // The type is interface{} with range: 0..65535.
    CvVoIPCallHistoryRemMediaPort interface{}

    // The object indicates whether or not the SRTP (Secured RTP) was enabled for
    // the voice call. The type is bool.
    CvVoIPCallHistorySRTPEnable interface{}

    // The Calculated Planning Impairment Factor (Icpif) of the call  that is
    // associated to this call leg. The value in this object is computed by the
    // following equation. Icpif of the fallback probe = Itotal (total impairment
    // value)  - configured fallback (Expectation Factor). A value of 0 implies
    // that Icpif was not calculated and is meaningless for this attempt. The type
    // is interface{} with range: -2147483648..2147483647.
    CvVoIPCallHistoryFallbackIcpif interface{}

    // FallbackLoss is the percentage of loss packets based on the total packets
    // sent. The type is interface{} with range: 0..4294967295.
    CvVoIPCallHistoryFallbackLoss interface{}

    // The FallbackDelay is calculated as follows - Take the sum of the round
    // trips for all the probes,  divide by the number of probes,  and divide by
    // two to get the one-way delay.   Then add in jitter_in or jiter_out, which
    // ever is higher. The type is interface{} with range: 0..4294967295.
    CvVoIPCallHistoryFallbackDelay interface{}

    // If the object has a value true(1) octet align operation is used, and if the
    // value is false(2), bandwidth efficient operation is used. This object is
    // not instantiated when the object cvVoIPCallHistoryCoderTypeRate is not
    // equal to gsmAmrNb enum. The type is bool.
    CvVoIPCallHistoryOctetAligned interface{}

    // This object indicates modes of Bit rates. This object is not instantiated
    // when the object cvVoIPCallHistoryCoderTypeRate is not equal to gsmAmrNb
    // enum. The type is map[string]bool.
    CvVoIPCallHistoryBitRates interface{}

    // The object indicates the interval (N frame-blocks) at which codec mode
    // changes are allowed. This object is not instantiated when the object
    // cvVoIPCallHistoryCoderTypeRate is not equal to gsmAmrNb enum. The type is
    // interface{} with range: 1..100. Units are frame-blocks.
    CvVoIPCallHistoryModeChgPeriod interface{}

    // If the object has a value of true(1), mode changes will be made to only
    // neighboring modes set to cvVoIPCallHistoryBitRates object. If the value is
    // false(2), mode changes will be allowed to any modes set to
    // cvVoIPCallHistoryBitRates object. This object is not instantiated when the
    // object cvVoIPCallHistoryCoderTypeRate is not equal to gsmAmrNb enum. The
    // type is bool.
    CvVoIPCallHistoryModeChgNeighbor interface{}

    // The object indicates the maximum amount of media that can be encapsulated
    // in a payload. Supported value is 20 msec. This object is not instantiated
    // when the object cvVoIPCallHistoryCoderTypeRate is not equal to gsmAmrNb
    // enum. The type is interface{} with range: 20..100. Units are milliseconds.
    CvVoIPCallHistoryMaxPtime interface{}

    // If the object has a value of true(1), frame CRC will be included in the
    // payload and if the value is false(2), frame CRC will not be included in the
    // payload. This object is applicable only when RTP frame type is octet
    // aligned. This object is not instantiated when the object
    // cvVoIPCallHistoryCoderTypeRate is not equal to gsmAmrNb enum. The type is
    // bool.
    CvVoIPCallHistoryCRC interface{}

    // If the object has a value of true(1), payload employs robust sorting and if
    // the value is false(2), payload does not employ robust sorting. This object
    // is applicable only when RTP frame type is octet aligned. This object is not
    // instantiated when the object cvVoIPCallHistoryCoderTypeRate is not equal to
    // gsmAmrNb enum. The type is bool.
    CvVoIPCallHistoryRobustSorting interface{}

    // The object indicates the RTP encapsulation type. Supported RTP
    // encapsulation type is RFC3267. This object is not instantiated when the
    // object cvVoIPCallHistoryCoderTypeRate is not equal to gsmAmrNb enum. The
    // type is CvAmrNbRtpEncap.
    CvVoIPCallHistoryEncap interface{}

    // The object indicates the maximum number of frame-blocks allowed in an
    // interleaving group. It indicates that frame-block level interleaving will
    // be used for that session. If this object is not set, interleaving is not
    // used. This object is applicable only when RTP frame type is octet aligned.
    // This object is not instantiated when the object
    // cvVoIPCallHistoryCoderTypeRate is not equal to gsmAmrNb enum. The type is
    // interface{} with range: 1..50. Units are frame-blocks.
    CvVoIPCallHistoryInterleaving interface{}

    // The object indicates the length of the time in milliseconds represented by
    // the media of the packet. Supported value is 20 milliseconds. This object is
    // not instantiated when the object cvVoIPCallHistoryCoderTypeRate is not
    // equal to gsmAmrNb enum. The type is interface{} with range: 20..100. Units
    // are milliseconds.
    CvVoIPCallHistoryPtime interface{}

    // The object indicates the number of audio channels. Supported value is 1.
    // This object is not instantiated when the object
    // cvVoIPCallHistoryCoderTypeRate is not equal to gsmAmrNb enum. The type is
    // interface{} with range: 1..6. Units are channels.
    CvVoIPCallHistoryChannels interface{}

    // The object indicates the iLBC mode. The value of this object is valid only
    // if  cvVoIPCallHistoryCoderTypeRate is equal to  'iLBC'. The type is
    // CvIlbcFrameMode. Units are milliseconds.
    CvVoIPCallHistoryCoderMode interface{}

    // This object represents the call identifier for the VoIP leg, which was
    // assigned to the call. The type is interface{} with range: 1..4294967295.
    CvVoIPCallHistoryCallId interface{}

    // The call reference ID associates the video call entry and voice call entry
    // of the same endpoint.  An audio-only call may or may not have a valid call
    // reference ID (i.e. value greater than zero), but in both cases, there will
    // not be a video call entry associated with it.   For a video call, the
    // video-specific information  is stored in a call entry in
    // cVideoSessionActive of CISCO-VIDEO-SESSION-MIB, in which the call reference
    // ID is also identified. The type is interface{} with range: 0..4294967295.
    CvVoIPCallHistoryCallReferenceId interface{}

    // This object indicates the session ID assigned by the call manager to
    // identify call legs that belong to the same call session.  This session ID
    // (history) represents a completed call session, whereas the active session
    // ID (cvVoIPCallActiveSessionId) represents an ongoing session. The type is
    // interface{} with range: 0..4294967295.
    CvVoIPCallHistorySessionId interface{}
}

func (cvVoIPCallHistoryEntry *CISCOVOICEDIALCONTROLMIB_CvVoIPCallHistoryTable_CvVoIPCallHistoryEntry) GetEntityData() *types.CommonEntityData {
    cvVoIPCallHistoryEntry.EntityData.YFilter = cvVoIPCallHistoryEntry.YFilter
    cvVoIPCallHistoryEntry.EntityData.YangName = "cvVoIPCallHistoryEntry"
    cvVoIPCallHistoryEntry.EntityData.BundleName = "cisco_ios_xe"
    cvVoIPCallHistoryEntry.EntityData.ParentYangName = "cvVoIPCallHistoryTable"
    cvVoIPCallHistoryEntry.EntityData.SegmentPath = "cvVoIPCallHistoryEntry" + types.AddKeyToken(cvVoIPCallHistoryEntry.CCallHistoryIndex, "cCallHistoryIndex")
    cvVoIPCallHistoryEntry.EntityData.AbsolutePath = "CISCO-VOICE-DIAL-CONTROL-MIB:CISCO-VOICE-DIAL-CONTROL-MIB/cvVoIPCallHistoryTable/" + cvVoIPCallHistoryEntry.EntityData.SegmentPath
    cvVoIPCallHistoryEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvVoIPCallHistoryEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvVoIPCallHistoryEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvVoIPCallHistoryEntry.EntityData.Children = types.NewOrderedMap()
    cvVoIPCallHistoryEntry.EntityData.Leafs = types.NewOrderedMap()
    cvVoIPCallHistoryEntry.EntityData.Leafs.Append("cCallHistoryIndex", types.YLeaf{"CCallHistoryIndex", cvVoIPCallHistoryEntry.CCallHistoryIndex})
    cvVoIPCallHistoryEntry.EntityData.Leafs.Append("cvVoIPCallHistoryConnectionId", types.YLeaf{"CvVoIPCallHistoryConnectionId", cvVoIPCallHistoryEntry.CvVoIPCallHistoryConnectionId})
    cvVoIPCallHistoryEntry.EntityData.Leafs.Append("cvVoIPCallHistoryRemoteIPAddress", types.YLeaf{"CvVoIPCallHistoryRemoteIPAddress", cvVoIPCallHistoryEntry.CvVoIPCallHistoryRemoteIPAddress})
    cvVoIPCallHistoryEntry.EntityData.Leafs.Append("cvVoIPCallHistoryRemoteUDPPort", types.YLeaf{"CvVoIPCallHistoryRemoteUDPPort", cvVoIPCallHistoryEntry.CvVoIPCallHistoryRemoteUDPPort})
    cvVoIPCallHistoryEntry.EntityData.Leafs.Append("cvVoIPCallHistoryRoundTripDelay", types.YLeaf{"CvVoIPCallHistoryRoundTripDelay", cvVoIPCallHistoryEntry.CvVoIPCallHistoryRoundTripDelay})
    cvVoIPCallHistoryEntry.EntityData.Leafs.Append("cvVoIPCallHistorySelectedQoS", types.YLeaf{"CvVoIPCallHistorySelectedQoS", cvVoIPCallHistoryEntry.CvVoIPCallHistorySelectedQoS})
    cvVoIPCallHistoryEntry.EntityData.Leafs.Append("cvVoIPCallHistorySessionProtocol", types.YLeaf{"CvVoIPCallHistorySessionProtocol", cvVoIPCallHistoryEntry.CvVoIPCallHistorySessionProtocol})
    cvVoIPCallHistoryEntry.EntityData.Leafs.Append("cvVoIPCallHistorySessionTarget", types.YLeaf{"CvVoIPCallHistorySessionTarget", cvVoIPCallHistoryEntry.CvVoIPCallHistorySessionTarget})
    cvVoIPCallHistoryEntry.EntityData.Leafs.Append("cvVoIPCallHistoryOnTimeRvPlayout", types.YLeaf{"CvVoIPCallHistoryOnTimeRvPlayout", cvVoIPCallHistoryEntry.CvVoIPCallHistoryOnTimeRvPlayout})
    cvVoIPCallHistoryEntry.EntityData.Leafs.Append("cvVoIPCallHistoryGapFillWithSilence", types.YLeaf{"CvVoIPCallHistoryGapFillWithSilence", cvVoIPCallHistoryEntry.CvVoIPCallHistoryGapFillWithSilence})
    cvVoIPCallHistoryEntry.EntityData.Leafs.Append("cvVoIPCallHistoryGapFillWithPrediction", types.YLeaf{"CvVoIPCallHistoryGapFillWithPrediction", cvVoIPCallHistoryEntry.CvVoIPCallHistoryGapFillWithPrediction})
    cvVoIPCallHistoryEntry.EntityData.Leafs.Append("cvVoIPCallHistoryGapFillWithInterpolation", types.YLeaf{"CvVoIPCallHistoryGapFillWithInterpolation", cvVoIPCallHistoryEntry.CvVoIPCallHistoryGapFillWithInterpolation})
    cvVoIPCallHistoryEntry.EntityData.Leafs.Append("cvVoIPCallHistoryGapFillWithRedundancy", types.YLeaf{"CvVoIPCallHistoryGapFillWithRedundancy", cvVoIPCallHistoryEntry.CvVoIPCallHistoryGapFillWithRedundancy})
    cvVoIPCallHistoryEntry.EntityData.Leafs.Append("cvVoIPCallHistoryHiWaterPlayoutDelay", types.YLeaf{"CvVoIPCallHistoryHiWaterPlayoutDelay", cvVoIPCallHistoryEntry.CvVoIPCallHistoryHiWaterPlayoutDelay})
    cvVoIPCallHistoryEntry.EntityData.Leafs.Append("cvVoIPCallHistoryLoWaterPlayoutDelay", types.YLeaf{"CvVoIPCallHistoryLoWaterPlayoutDelay", cvVoIPCallHistoryEntry.CvVoIPCallHistoryLoWaterPlayoutDelay})
    cvVoIPCallHistoryEntry.EntityData.Leafs.Append("cvVoIPCallHistoryReceiveDelay", types.YLeaf{"CvVoIPCallHistoryReceiveDelay", cvVoIPCallHistoryEntry.CvVoIPCallHistoryReceiveDelay})
    cvVoIPCallHistoryEntry.EntityData.Leafs.Append("cvVoIPCallHistoryVADEnable", types.YLeaf{"CvVoIPCallHistoryVADEnable", cvVoIPCallHistoryEntry.CvVoIPCallHistoryVADEnable})
    cvVoIPCallHistoryEntry.EntityData.Leafs.Append("cvVoIPCallHistoryCoderTypeRate", types.YLeaf{"CvVoIPCallHistoryCoderTypeRate", cvVoIPCallHistoryEntry.CvVoIPCallHistoryCoderTypeRate})
    cvVoIPCallHistoryEntry.EntityData.Leafs.Append("cvVoIPCallHistoryIcpif", types.YLeaf{"CvVoIPCallHistoryIcpif", cvVoIPCallHistoryEntry.CvVoIPCallHistoryIcpif})
    cvVoIPCallHistoryEntry.EntityData.Leafs.Append("cvVoIPCallHistoryLostPackets", types.YLeaf{"CvVoIPCallHistoryLostPackets", cvVoIPCallHistoryEntry.CvVoIPCallHistoryLostPackets})
    cvVoIPCallHistoryEntry.EntityData.Leafs.Append("cvVoIPCallHistoryEarlyPackets", types.YLeaf{"CvVoIPCallHistoryEarlyPackets", cvVoIPCallHistoryEntry.CvVoIPCallHistoryEarlyPackets})
    cvVoIPCallHistoryEntry.EntityData.Leafs.Append("cvVoIPCallHistoryLatePackets", types.YLeaf{"CvVoIPCallHistoryLatePackets", cvVoIPCallHistoryEntry.CvVoIPCallHistoryLatePackets})
    cvVoIPCallHistoryEntry.EntityData.Leafs.Append("cvVoIPCallHistoryUsername", types.YLeaf{"CvVoIPCallHistoryUsername", cvVoIPCallHistoryEntry.CvVoIPCallHistoryUsername})
    cvVoIPCallHistoryEntry.EntityData.Leafs.Append("cvVoIPCallHistoryProtocolCallId", types.YLeaf{"CvVoIPCallHistoryProtocolCallId", cvVoIPCallHistoryEntry.CvVoIPCallHistoryProtocolCallId})
    cvVoIPCallHistoryEntry.EntityData.Leafs.Append("cvVoIPCallHistoryRemSigIPAddrT", types.YLeaf{"CvVoIPCallHistoryRemSigIPAddrT", cvVoIPCallHistoryEntry.CvVoIPCallHistoryRemSigIPAddrT})
    cvVoIPCallHistoryEntry.EntityData.Leafs.Append("cvVoIPCallHistoryRemSigIPAddr", types.YLeaf{"CvVoIPCallHistoryRemSigIPAddr", cvVoIPCallHistoryEntry.CvVoIPCallHistoryRemSigIPAddr})
    cvVoIPCallHistoryEntry.EntityData.Leafs.Append("cvVoIPCallHistoryRemSigPort", types.YLeaf{"CvVoIPCallHistoryRemSigPort", cvVoIPCallHistoryEntry.CvVoIPCallHistoryRemSigPort})
    cvVoIPCallHistoryEntry.EntityData.Leafs.Append("cvVoIPCallHistoryRemMediaIPAddrT", types.YLeaf{"CvVoIPCallHistoryRemMediaIPAddrT", cvVoIPCallHistoryEntry.CvVoIPCallHistoryRemMediaIPAddrT})
    cvVoIPCallHistoryEntry.EntityData.Leafs.Append("cvVoIPCallHistoryRemMediaIPAddr", types.YLeaf{"CvVoIPCallHistoryRemMediaIPAddr", cvVoIPCallHistoryEntry.CvVoIPCallHistoryRemMediaIPAddr})
    cvVoIPCallHistoryEntry.EntityData.Leafs.Append("cvVoIPCallHistoryRemMediaPort", types.YLeaf{"CvVoIPCallHistoryRemMediaPort", cvVoIPCallHistoryEntry.CvVoIPCallHistoryRemMediaPort})
    cvVoIPCallHistoryEntry.EntityData.Leafs.Append("cvVoIPCallHistorySRTPEnable", types.YLeaf{"CvVoIPCallHistorySRTPEnable", cvVoIPCallHistoryEntry.CvVoIPCallHistorySRTPEnable})
    cvVoIPCallHistoryEntry.EntityData.Leafs.Append("cvVoIPCallHistoryFallbackIcpif", types.YLeaf{"CvVoIPCallHistoryFallbackIcpif", cvVoIPCallHistoryEntry.CvVoIPCallHistoryFallbackIcpif})
    cvVoIPCallHistoryEntry.EntityData.Leafs.Append("cvVoIPCallHistoryFallbackLoss", types.YLeaf{"CvVoIPCallHistoryFallbackLoss", cvVoIPCallHistoryEntry.CvVoIPCallHistoryFallbackLoss})
    cvVoIPCallHistoryEntry.EntityData.Leafs.Append("cvVoIPCallHistoryFallbackDelay", types.YLeaf{"CvVoIPCallHistoryFallbackDelay", cvVoIPCallHistoryEntry.CvVoIPCallHistoryFallbackDelay})
    cvVoIPCallHistoryEntry.EntityData.Leafs.Append("cvVoIPCallHistoryOctetAligned", types.YLeaf{"CvVoIPCallHistoryOctetAligned", cvVoIPCallHistoryEntry.CvVoIPCallHistoryOctetAligned})
    cvVoIPCallHistoryEntry.EntityData.Leafs.Append("cvVoIPCallHistoryBitRates", types.YLeaf{"CvVoIPCallHistoryBitRates", cvVoIPCallHistoryEntry.CvVoIPCallHistoryBitRates})
    cvVoIPCallHistoryEntry.EntityData.Leafs.Append("cvVoIPCallHistoryModeChgPeriod", types.YLeaf{"CvVoIPCallHistoryModeChgPeriod", cvVoIPCallHistoryEntry.CvVoIPCallHistoryModeChgPeriod})
    cvVoIPCallHistoryEntry.EntityData.Leafs.Append("cvVoIPCallHistoryModeChgNeighbor", types.YLeaf{"CvVoIPCallHistoryModeChgNeighbor", cvVoIPCallHistoryEntry.CvVoIPCallHistoryModeChgNeighbor})
    cvVoIPCallHistoryEntry.EntityData.Leafs.Append("cvVoIPCallHistoryMaxPtime", types.YLeaf{"CvVoIPCallHistoryMaxPtime", cvVoIPCallHistoryEntry.CvVoIPCallHistoryMaxPtime})
    cvVoIPCallHistoryEntry.EntityData.Leafs.Append("cvVoIPCallHistoryCRC", types.YLeaf{"CvVoIPCallHistoryCRC", cvVoIPCallHistoryEntry.CvVoIPCallHistoryCRC})
    cvVoIPCallHistoryEntry.EntityData.Leafs.Append("cvVoIPCallHistoryRobustSorting", types.YLeaf{"CvVoIPCallHistoryRobustSorting", cvVoIPCallHistoryEntry.CvVoIPCallHistoryRobustSorting})
    cvVoIPCallHistoryEntry.EntityData.Leafs.Append("cvVoIPCallHistoryEncap", types.YLeaf{"CvVoIPCallHistoryEncap", cvVoIPCallHistoryEntry.CvVoIPCallHistoryEncap})
    cvVoIPCallHistoryEntry.EntityData.Leafs.Append("cvVoIPCallHistoryInterleaving", types.YLeaf{"CvVoIPCallHistoryInterleaving", cvVoIPCallHistoryEntry.CvVoIPCallHistoryInterleaving})
    cvVoIPCallHistoryEntry.EntityData.Leafs.Append("cvVoIPCallHistoryPtime", types.YLeaf{"CvVoIPCallHistoryPtime", cvVoIPCallHistoryEntry.CvVoIPCallHistoryPtime})
    cvVoIPCallHistoryEntry.EntityData.Leafs.Append("cvVoIPCallHistoryChannels", types.YLeaf{"CvVoIPCallHistoryChannels", cvVoIPCallHistoryEntry.CvVoIPCallHistoryChannels})
    cvVoIPCallHistoryEntry.EntityData.Leafs.Append("cvVoIPCallHistoryCoderMode", types.YLeaf{"CvVoIPCallHistoryCoderMode", cvVoIPCallHistoryEntry.CvVoIPCallHistoryCoderMode})
    cvVoIPCallHistoryEntry.EntityData.Leafs.Append("cvVoIPCallHistoryCallId", types.YLeaf{"CvVoIPCallHistoryCallId", cvVoIPCallHistoryEntry.CvVoIPCallHistoryCallId})
    cvVoIPCallHistoryEntry.EntityData.Leafs.Append("cvVoIPCallHistoryCallReferenceId", types.YLeaf{"CvVoIPCallHistoryCallReferenceId", cvVoIPCallHistoryEntry.CvVoIPCallHistoryCallReferenceId})
    cvVoIPCallHistoryEntry.EntityData.Leafs.Append("cvVoIPCallHistorySessionId", types.YLeaf{"CvVoIPCallHistorySessionId", cvVoIPCallHistoryEntry.CvVoIPCallHistorySessionId})

    cvVoIPCallHistoryEntry.EntityData.YListKeys = []string {"CCallHistoryIndex"}

    return &(cvVoIPCallHistoryEntry.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_CvCallRateStatsTable
// This table represents voice call rate measurement in various
// interval lengths defined by the 
// CvCallVolumeStatsIntvlType object.
// 
// Each interval may contain one or more entries to allow for
// detailed measurement to be collected.
type CISCOVOICEDIALCONTROLMIB_CvCallRateStatsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This is a conceptual-row in cvCallRateStatsTable This entry is created at
    // the system initialization and is updated at every epoch based on
    // CvCallVolumeStatsIntvlType. The type is slice of
    // CISCOVOICEDIALCONTROLMIB_CvCallRateStatsTable_CvCallRateStatsEntry.
    CvCallRateStatsEntry []*CISCOVOICEDIALCONTROLMIB_CvCallRateStatsTable_CvCallRateStatsEntry
}

func (cvCallRateStatsTable *CISCOVOICEDIALCONTROLMIB_CvCallRateStatsTable) GetEntityData() *types.CommonEntityData {
    cvCallRateStatsTable.EntityData.YFilter = cvCallRateStatsTable.YFilter
    cvCallRateStatsTable.EntityData.YangName = "cvCallRateStatsTable"
    cvCallRateStatsTable.EntityData.BundleName = "cisco_ios_xe"
    cvCallRateStatsTable.EntityData.ParentYangName = "CISCO-VOICE-DIAL-CONTROL-MIB"
    cvCallRateStatsTable.EntityData.SegmentPath = "cvCallRateStatsTable"
    cvCallRateStatsTable.EntityData.AbsolutePath = "CISCO-VOICE-DIAL-CONTROL-MIB:CISCO-VOICE-DIAL-CONTROL-MIB/" + cvCallRateStatsTable.EntityData.SegmentPath
    cvCallRateStatsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvCallRateStatsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvCallRateStatsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvCallRateStatsTable.EntityData.Children = types.NewOrderedMap()
    cvCallRateStatsTable.EntityData.Children.Append("cvCallRateStatsEntry", types.YChild{"CvCallRateStatsEntry", nil})
    for i := range cvCallRateStatsTable.CvCallRateStatsEntry {
        cvCallRateStatsTable.EntityData.Children.Append(types.GetSegmentPath(cvCallRateStatsTable.CvCallRateStatsEntry[i]), types.YChild{"CvCallRateStatsEntry", cvCallRateStatsTable.CvCallRateStatsEntry[i]})
    }
    cvCallRateStatsTable.EntityData.Leafs = types.NewOrderedMap()

    cvCallRateStatsTable.EntityData.YListKeys = []string {}

    return &(cvCallRateStatsTable.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_CvCallRateStatsTable_CvCallRateStatsEntry
// This is a conceptual-row in cvCallRateStatsTable
// This entry is created at the system initialization and is
// updated at every epoch based on CvCallVolumeStatsIntvlType
type CISCOVOICEDIALCONTROLMIB_CvCallRateStatsTable_CvCallRateStatsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The Object indexes in Call Rate Table to select
    // one among three interval-tables.  The different types in this table are
    // represented by  CvCallVolumeStatsIntvlType. The type is
    // CvCallVolumeStatsIntvlType.
    CvCallRateStatsIntvlDurUnits interface{}

    // This attribute is a key. This is an index that references to the different
    // past periods in given in interval of call rate table. This range is 1-60
    // for Seconds and Minutes table  wherein 1-72 for hours table. The type is
    // interface{} with range: 1..72.
    CvCallRateStatsIntvlDur interface{}

    // This object indicates the maximum calls per second that occured for the
    // given period for the given interval. The type is interface{} with range:
    // 0..4294967295. Units are calls-per-second.
    CvCallRateStatsMaxVal interface{}

    // This object indicates the average calls per second that occured for the
    // given period for the given interval. The type is interface{} with range:
    // 0..4294967295. Units are calls-per-second.
    CvCallRateStatsAvgVal interface{}
}

func (cvCallRateStatsEntry *CISCOVOICEDIALCONTROLMIB_CvCallRateStatsTable_CvCallRateStatsEntry) GetEntityData() *types.CommonEntityData {
    cvCallRateStatsEntry.EntityData.YFilter = cvCallRateStatsEntry.YFilter
    cvCallRateStatsEntry.EntityData.YangName = "cvCallRateStatsEntry"
    cvCallRateStatsEntry.EntityData.BundleName = "cisco_ios_xe"
    cvCallRateStatsEntry.EntityData.ParentYangName = "cvCallRateStatsTable"
    cvCallRateStatsEntry.EntityData.SegmentPath = "cvCallRateStatsEntry" + types.AddKeyToken(cvCallRateStatsEntry.CvCallRateStatsIntvlDurUnits, "cvCallRateStatsIntvlDurUnits") + types.AddKeyToken(cvCallRateStatsEntry.CvCallRateStatsIntvlDur, "cvCallRateStatsIntvlDur")
    cvCallRateStatsEntry.EntityData.AbsolutePath = "CISCO-VOICE-DIAL-CONTROL-MIB:CISCO-VOICE-DIAL-CONTROL-MIB/cvCallRateStatsTable/" + cvCallRateStatsEntry.EntityData.SegmentPath
    cvCallRateStatsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvCallRateStatsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvCallRateStatsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvCallRateStatsEntry.EntityData.Children = types.NewOrderedMap()
    cvCallRateStatsEntry.EntityData.Leafs = types.NewOrderedMap()
    cvCallRateStatsEntry.EntityData.Leafs.Append("cvCallRateStatsIntvlDurUnits", types.YLeaf{"CvCallRateStatsIntvlDurUnits", cvCallRateStatsEntry.CvCallRateStatsIntvlDurUnits})
    cvCallRateStatsEntry.EntityData.Leafs.Append("cvCallRateStatsIntvlDur", types.YLeaf{"CvCallRateStatsIntvlDur", cvCallRateStatsEntry.CvCallRateStatsIntvlDur})
    cvCallRateStatsEntry.EntityData.Leafs.Append("cvCallRateStatsMaxVal", types.YLeaf{"CvCallRateStatsMaxVal", cvCallRateStatsEntry.CvCallRateStatsMaxVal})
    cvCallRateStatsEntry.EntityData.Leafs.Append("cvCallRateStatsAvgVal", types.YLeaf{"CvCallRateStatsAvgVal", cvCallRateStatsEntry.CvCallRateStatsAvgVal})

    cvCallRateStatsEntry.EntityData.YListKeys = []string {"CvCallRateStatsIntvlDurUnits", "CvCallRateStatsIntvlDur"}

    return &(cvCallRateStatsEntry.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_CvCallLegRateStatsTable
// cvCallLegRateStatsTable table represents voice call leg rate
// measurement in various interval lengths defined by 
// the CvCallVolumeStatsIntvlType object.
// Each interval may contain one or more entries to allow for
// detailed measurement to be collected.
type CISCOVOICEDIALCONTROLMIB_CvCallLegRateStatsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This is a conceptual-row in cvCallLegRateStatsTable This entry is created
    // at the system initialization and is updated at every epoch based on
    // CvCallVolumeStatsIntvlType. The type is slice of
    // CISCOVOICEDIALCONTROLMIB_CvCallLegRateStatsTable_CvCallLegRateStatsEntry.
    CvCallLegRateStatsEntry []*CISCOVOICEDIALCONTROLMIB_CvCallLegRateStatsTable_CvCallLegRateStatsEntry
}

func (cvCallLegRateStatsTable *CISCOVOICEDIALCONTROLMIB_CvCallLegRateStatsTable) GetEntityData() *types.CommonEntityData {
    cvCallLegRateStatsTable.EntityData.YFilter = cvCallLegRateStatsTable.YFilter
    cvCallLegRateStatsTable.EntityData.YangName = "cvCallLegRateStatsTable"
    cvCallLegRateStatsTable.EntityData.BundleName = "cisco_ios_xe"
    cvCallLegRateStatsTable.EntityData.ParentYangName = "CISCO-VOICE-DIAL-CONTROL-MIB"
    cvCallLegRateStatsTable.EntityData.SegmentPath = "cvCallLegRateStatsTable"
    cvCallLegRateStatsTable.EntityData.AbsolutePath = "CISCO-VOICE-DIAL-CONTROL-MIB:CISCO-VOICE-DIAL-CONTROL-MIB/" + cvCallLegRateStatsTable.EntityData.SegmentPath
    cvCallLegRateStatsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvCallLegRateStatsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvCallLegRateStatsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvCallLegRateStatsTable.EntityData.Children = types.NewOrderedMap()
    cvCallLegRateStatsTable.EntityData.Children.Append("cvCallLegRateStatsEntry", types.YChild{"CvCallLegRateStatsEntry", nil})
    for i := range cvCallLegRateStatsTable.CvCallLegRateStatsEntry {
        cvCallLegRateStatsTable.EntityData.Children.Append(types.GetSegmentPath(cvCallLegRateStatsTable.CvCallLegRateStatsEntry[i]), types.YChild{"CvCallLegRateStatsEntry", cvCallLegRateStatsTable.CvCallLegRateStatsEntry[i]})
    }
    cvCallLegRateStatsTable.EntityData.Leafs = types.NewOrderedMap()

    cvCallLegRateStatsTable.EntityData.YListKeys = []string {}

    return &(cvCallLegRateStatsTable.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_CvCallLegRateStatsTable_CvCallLegRateStatsEntry
// This is a conceptual-row in cvCallLegRateStatsTable
// This entry is created at the system initialization and is
// updated at every epoch based on CvCallVolumeStatsIntvlType
type CISCOVOICEDIALCONTROLMIB_CvCallLegRateStatsTable_CvCallLegRateStatsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The Object indexes in Call Leg Rate Table to
    // select one among three interval-tables.  The different types in this table
    // are represented by  CvCallVolumeStatsIntvlType. The type is
    // CvCallVolumeStatsIntvlType.
    CvCallLegRateStatsIntvlDurUnits interface{}

    // This attribute is a key. This is an index that references to the different
    // past periods in given in interval of call rate table. This range is 1-60
    // for Seconds and Minutes table  wherein 1-72 for hours table. The type is
    // interface{} with range: 1..72.
    CvCallLegRateStatsIntvlDur interface{}

    // This object indicates the maximum call-legs per second that occured for the
    // given period for the given interval. The type is interface{} with range:
    // 0..4294967295. Units are call-legs per second.
    CvCallLegRateStatsMaxVal interface{}

    // This object indicates the average call-legs per second that occured for the
    // given period for the given interval. The type is interface{} with range:
    // 0..4294967295. Units are call-legs per second.
    CvCallLegRateStatsAvgVal interface{}
}

func (cvCallLegRateStatsEntry *CISCOVOICEDIALCONTROLMIB_CvCallLegRateStatsTable_CvCallLegRateStatsEntry) GetEntityData() *types.CommonEntityData {
    cvCallLegRateStatsEntry.EntityData.YFilter = cvCallLegRateStatsEntry.YFilter
    cvCallLegRateStatsEntry.EntityData.YangName = "cvCallLegRateStatsEntry"
    cvCallLegRateStatsEntry.EntityData.BundleName = "cisco_ios_xe"
    cvCallLegRateStatsEntry.EntityData.ParentYangName = "cvCallLegRateStatsTable"
    cvCallLegRateStatsEntry.EntityData.SegmentPath = "cvCallLegRateStatsEntry" + types.AddKeyToken(cvCallLegRateStatsEntry.CvCallLegRateStatsIntvlDurUnits, "cvCallLegRateStatsIntvlDurUnits") + types.AddKeyToken(cvCallLegRateStatsEntry.CvCallLegRateStatsIntvlDur, "cvCallLegRateStatsIntvlDur")
    cvCallLegRateStatsEntry.EntityData.AbsolutePath = "CISCO-VOICE-DIAL-CONTROL-MIB:CISCO-VOICE-DIAL-CONTROL-MIB/cvCallLegRateStatsTable/" + cvCallLegRateStatsEntry.EntityData.SegmentPath
    cvCallLegRateStatsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvCallLegRateStatsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvCallLegRateStatsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvCallLegRateStatsEntry.EntityData.Children = types.NewOrderedMap()
    cvCallLegRateStatsEntry.EntityData.Leafs = types.NewOrderedMap()
    cvCallLegRateStatsEntry.EntityData.Leafs.Append("cvCallLegRateStatsIntvlDurUnits", types.YLeaf{"CvCallLegRateStatsIntvlDurUnits", cvCallLegRateStatsEntry.CvCallLegRateStatsIntvlDurUnits})
    cvCallLegRateStatsEntry.EntityData.Leafs.Append("cvCallLegRateStatsIntvlDur", types.YLeaf{"CvCallLegRateStatsIntvlDur", cvCallLegRateStatsEntry.CvCallLegRateStatsIntvlDur})
    cvCallLegRateStatsEntry.EntityData.Leafs.Append("cvCallLegRateStatsMaxVal", types.YLeaf{"CvCallLegRateStatsMaxVal", cvCallLegRateStatsEntry.CvCallLegRateStatsMaxVal})
    cvCallLegRateStatsEntry.EntityData.Leafs.Append("cvCallLegRateStatsAvgVal", types.YLeaf{"CvCallLegRateStatsAvgVal", cvCallLegRateStatsEntry.CvCallLegRateStatsAvgVal})

    cvCallLegRateStatsEntry.EntityData.YListKeys = []string {"CvCallLegRateStatsIntvlDurUnits", "CvCallLegRateStatsIntvlDur"}

    return &(cvCallLegRateStatsEntry.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_CvActiveCallStatsTable
// This table represents the active voice calls in various
// interval lengths defined by the 
// CvCallVolumeStatsIntvlType object.
// 
// Each interval may contain one or more entries to allow for
// detailed measurement to be collected.
type CISCOVOICEDIALCONTROLMIB_CvActiveCallStatsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This is a conceptual-row in cvActiveCallStatsTable This entry is created at
    // the system initialization and is updated at every epoch based on
    // CvCallVolumeStatsIntvlType. The type is slice of
    // CISCOVOICEDIALCONTROLMIB_CvActiveCallStatsTable_CvActiveCallStatsEntry.
    CvActiveCallStatsEntry []*CISCOVOICEDIALCONTROLMIB_CvActiveCallStatsTable_CvActiveCallStatsEntry
}

func (cvActiveCallStatsTable *CISCOVOICEDIALCONTROLMIB_CvActiveCallStatsTable) GetEntityData() *types.CommonEntityData {
    cvActiveCallStatsTable.EntityData.YFilter = cvActiveCallStatsTable.YFilter
    cvActiveCallStatsTable.EntityData.YangName = "cvActiveCallStatsTable"
    cvActiveCallStatsTable.EntityData.BundleName = "cisco_ios_xe"
    cvActiveCallStatsTable.EntityData.ParentYangName = "CISCO-VOICE-DIAL-CONTROL-MIB"
    cvActiveCallStatsTable.EntityData.SegmentPath = "cvActiveCallStatsTable"
    cvActiveCallStatsTable.EntityData.AbsolutePath = "CISCO-VOICE-DIAL-CONTROL-MIB:CISCO-VOICE-DIAL-CONTROL-MIB/" + cvActiveCallStatsTable.EntityData.SegmentPath
    cvActiveCallStatsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvActiveCallStatsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvActiveCallStatsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvActiveCallStatsTable.EntityData.Children = types.NewOrderedMap()
    cvActiveCallStatsTable.EntityData.Children.Append("cvActiveCallStatsEntry", types.YChild{"CvActiveCallStatsEntry", nil})
    for i := range cvActiveCallStatsTable.CvActiveCallStatsEntry {
        cvActiveCallStatsTable.EntityData.Children.Append(types.GetSegmentPath(cvActiveCallStatsTable.CvActiveCallStatsEntry[i]), types.YChild{"CvActiveCallStatsEntry", cvActiveCallStatsTable.CvActiveCallStatsEntry[i]})
    }
    cvActiveCallStatsTable.EntityData.Leafs = types.NewOrderedMap()

    cvActiveCallStatsTable.EntityData.YListKeys = []string {}

    return &(cvActiveCallStatsTable.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_CvActiveCallStatsTable_CvActiveCallStatsEntry
// This is a conceptual-row in cvActiveCallStatsTable
// This entry is created at the system initialization and is
// updated at every epoch based on CvCallVolumeStatsIntvlType
type CISCOVOICEDIALCONTROLMIB_CvActiveCallStatsTable_CvActiveCallStatsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The Object indexes in Active Call Rate Table
    // (con-current calls table) to select one among three interval-tables.  The
    // different types in this table are represented by 
    // CvCallVolumeStatsIntvlType. The type is CvCallVolumeStatsIntvlType.
    CvActiveCallStatsIntvlDurUnits interface{}

    // This attribute is a key. This is an index that references to the different
    // past periods in given in interval of active call table. This range is 1-60
    // for Seconds and Minutes table  wherein 1-72 for hours table. The type is
    // interface{} with range: 1..72.
    CvActiveCallStatsIntvlDur interface{}

    // This object indicates the maximum number of active call that occured for
    // the given period for the given interval. The type is interface{} with
    // range: 0..4294967295. Units are calls.
    CvActiveCallStatsMaxVal interface{}

    // This object indicates the average number of active calls that occured for
    // the given period for the given interval. The type is interface{} with
    // range: 0..4294967295. Units are calls.
    CvActiveCallStatsAvgVal interface{}
}

func (cvActiveCallStatsEntry *CISCOVOICEDIALCONTROLMIB_CvActiveCallStatsTable_CvActiveCallStatsEntry) GetEntityData() *types.CommonEntityData {
    cvActiveCallStatsEntry.EntityData.YFilter = cvActiveCallStatsEntry.YFilter
    cvActiveCallStatsEntry.EntityData.YangName = "cvActiveCallStatsEntry"
    cvActiveCallStatsEntry.EntityData.BundleName = "cisco_ios_xe"
    cvActiveCallStatsEntry.EntityData.ParentYangName = "cvActiveCallStatsTable"
    cvActiveCallStatsEntry.EntityData.SegmentPath = "cvActiveCallStatsEntry" + types.AddKeyToken(cvActiveCallStatsEntry.CvActiveCallStatsIntvlDurUnits, "cvActiveCallStatsIntvlDurUnits") + types.AddKeyToken(cvActiveCallStatsEntry.CvActiveCallStatsIntvlDur, "cvActiveCallStatsIntvlDur")
    cvActiveCallStatsEntry.EntityData.AbsolutePath = "CISCO-VOICE-DIAL-CONTROL-MIB:CISCO-VOICE-DIAL-CONTROL-MIB/cvActiveCallStatsTable/" + cvActiveCallStatsEntry.EntityData.SegmentPath
    cvActiveCallStatsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvActiveCallStatsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvActiveCallStatsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvActiveCallStatsEntry.EntityData.Children = types.NewOrderedMap()
    cvActiveCallStatsEntry.EntityData.Leafs = types.NewOrderedMap()
    cvActiveCallStatsEntry.EntityData.Leafs.Append("cvActiveCallStatsIntvlDurUnits", types.YLeaf{"CvActiveCallStatsIntvlDurUnits", cvActiveCallStatsEntry.CvActiveCallStatsIntvlDurUnits})
    cvActiveCallStatsEntry.EntityData.Leafs.Append("cvActiveCallStatsIntvlDur", types.YLeaf{"CvActiveCallStatsIntvlDur", cvActiveCallStatsEntry.CvActiveCallStatsIntvlDur})
    cvActiveCallStatsEntry.EntityData.Leafs.Append("cvActiveCallStatsMaxVal", types.YLeaf{"CvActiveCallStatsMaxVal", cvActiveCallStatsEntry.CvActiveCallStatsMaxVal})
    cvActiveCallStatsEntry.EntityData.Leafs.Append("cvActiveCallStatsAvgVal", types.YLeaf{"CvActiveCallStatsAvgVal", cvActiveCallStatsEntry.CvActiveCallStatsAvgVal})

    cvActiveCallStatsEntry.EntityData.YListKeys = []string {"CvActiveCallStatsIntvlDurUnits", "CvActiveCallStatsIntvlDur"}

    return &(cvActiveCallStatsEntry.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_CvCallDurationStatsTable
// This table represents the number of calls below a specific
// duration in various interval length defined by 
// the CvCallVolumeStatsIntvlType object.  
// 
// The specific duration is configurable value of 
//  cvCallDurationStatsThreshold object.
// 
// Each interval may contain one or more entries to allow for 
// detailed measurement to be collected.
type CISCOVOICEDIALCONTROLMIB_CvCallDurationStatsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This is a conceptual-row in cvCallDurationStatsTable This entry is created
    // at the system initialization and is updated at every epoch based on
    // CvCallVolumeStatsIntvlType. The type is slice of
    // CISCOVOICEDIALCONTROLMIB_CvCallDurationStatsTable_CvCallDurationStatsEntry.
    CvCallDurationStatsEntry []*CISCOVOICEDIALCONTROLMIB_CvCallDurationStatsTable_CvCallDurationStatsEntry
}

func (cvCallDurationStatsTable *CISCOVOICEDIALCONTROLMIB_CvCallDurationStatsTable) GetEntityData() *types.CommonEntityData {
    cvCallDurationStatsTable.EntityData.YFilter = cvCallDurationStatsTable.YFilter
    cvCallDurationStatsTable.EntityData.YangName = "cvCallDurationStatsTable"
    cvCallDurationStatsTable.EntityData.BundleName = "cisco_ios_xe"
    cvCallDurationStatsTable.EntityData.ParentYangName = "CISCO-VOICE-DIAL-CONTROL-MIB"
    cvCallDurationStatsTable.EntityData.SegmentPath = "cvCallDurationStatsTable"
    cvCallDurationStatsTable.EntityData.AbsolutePath = "CISCO-VOICE-DIAL-CONTROL-MIB:CISCO-VOICE-DIAL-CONTROL-MIB/" + cvCallDurationStatsTable.EntityData.SegmentPath
    cvCallDurationStatsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvCallDurationStatsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvCallDurationStatsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvCallDurationStatsTable.EntityData.Children = types.NewOrderedMap()
    cvCallDurationStatsTable.EntityData.Children.Append("cvCallDurationStatsEntry", types.YChild{"CvCallDurationStatsEntry", nil})
    for i := range cvCallDurationStatsTable.CvCallDurationStatsEntry {
        cvCallDurationStatsTable.EntityData.Children.Append(types.GetSegmentPath(cvCallDurationStatsTable.CvCallDurationStatsEntry[i]), types.YChild{"CvCallDurationStatsEntry", cvCallDurationStatsTable.CvCallDurationStatsEntry[i]})
    }
    cvCallDurationStatsTable.EntityData.Leafs = types.NewOrderedMap()

    cvCallDurationStatsTable.EntityData.YListKeys = []string {}

    return &(cvCallDurationStatsTable.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_CvCallDurationStatsTable_CvCallDurationStatsEntry
// This is a conceptual-row in cvCallDurationStatsTable
// This entry is created at the system initialization and is
// updated at every epoch based on CvCallVolumeStatsIntvlType
type CISCOVOICEDIALCONTROLMIB_CvCallDurationStatsTable_CvCallDurationStatsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The Object indexes in Call Duration Table to
    // select one among three interval-tables.  The different types in this table
    // are represented by  CvCallVolumeStatsIntvlType. The type is
    // CvCallVolumeStatsIntvlType.
    CvCallDurationStatsIntvlDurUnits interface{}

    // This attribute is a key. This is an index that references to the different
    // past periods in given in interval of call Duration table. This range is
    // 1-60 for Seconds and Minutes table  wherein 1-72 for hours table. The type
    // is interface{} with range: 1..72.
    CvCallDurationStatsIntvlDur interface{}

    // This object indicates the maximum number of calls having a duration which
    // is below the threshold for the given interval. The type is interface{} with
    // range: 0..4294967295. Units are calls.
    CvCallDurationStatsMaxVal interface{}

    // This object indicates the average number of calls having a duration which
    // is below the threshold for the given interval. The type is interface{} with
    // range: 0..4294967295. Units are calls.
    CvCallDurationStatsAvgVal interface{}
}

func (cvCallDurationStatsEntry *CISCOVOICEDIALCONTROLMIB_CvCallDurationStatsTable_CvCallDurationStatsEntry) GetEntityData() *types.CommonEntityData {
    cvCallDurationStatsEntry.EntityData.YFilter = cvCallDurationStatsEntry.YFilter
    cvCallDurationStatsEntry.EntityData.YangName = "cvCallDurationStatsEntry"
    cvCallDurationStatsEntry.EntityData.BundleName = "cisco_ios_xe"
    cvCallDurationStatsEntry.EntityData.ParentYangName = "cvCallDurationStatsTable"
    cvCallDurationStatsEntry.EntityData.SegmentPath = "cvCallDurationStatsEntry" + types.AddKeyToken(cvCallDurationStatsEntry.CvCallDurationStatsIntvlDurUnits, "cvCallDurationStatsIntvlDurUnits") + types.AddKeyToken(cvCallDurationStatsEntry.CvCallDurationStatsIntvlDur, "cvCallDurationStatsIntvlDur")
    cvCallDurationStatsEntry.EntityData.AbsolutePath = "CISCO-VOICE-DIAL-CONTROL-MIB:CISCO-VOICE-DIAL-CONTROL-MIB/cvCallDurationStatsTable/" + cvCallDurationStatsEntry.EntityData.SegmentPath
    cvCallDurationStatsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvCallDurationStatsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvCallDurationStatsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvCallDurationStatsEntry.EntityData.Children = types.NewOrderedMap()
    cvCallDurationStatsEntry.EntityData.Leafs = types.NewOrderedMap()
    cvCallDurationStatsEntry.EntityData.Leafs.Append("cvCallDurationStatsIntvlDurUnits", types.YLeaf{"CvCallDurationStatsIntvlDurUnits", cvCallDurationStatsEntry.CvCallDurationStatsIntvlDurUnits})
    cvCallDurationStatsEntry.EntityData.Leafs.Append("cvCallDurationStatsIntvlDur", types.YLeaf{"CvCallDurationStatsIntvlDur", cvCallDurationStatsEntry.CvCallDurationStatsIntvlDur})
    cvCallDurationStatsEntry.EntityData.Leafs.Append("cvCallDurationStatsMaxVal", types.YLeaf{"CvCallDurationStatsMaxVal", cvCallDurationStatsEntry.CvCallDurationStatsMaxVal})
    cvCallDurationStatsEntry.EntityData.Leafs.Append("cvCallDurationStatsAvgVal", types.YLeaf{"CvCallDurationStatsAvgVal", cvCallDurationStatsEntry.CvCallDurationStatsAvgVal})

    cvCallDurationStatsEntry.EntityData.YListKeys = []string {"CvCallDurationStatsIntvlDurUnits", "CvCallDurationStatsIntvlDur"}

    return &(cvCallDurationStatsEntry.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_CvSipMsgRateStatsTable
// This table represents the SIP message rate measurement in
// various interval length defined by the 
// CvCallVolumeStatsIntvlType object.
// 
// Each interval may contain one or more entries to allow for
// detailed measurement to be collected
type CISCOVOICEDIALCONTROLMIB_CvSipMsgRateStatsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This is a conceptual-row in cvSipMsgRateStatsTable This entry is created at
    // the system initialization and is updated at every epoch based on
    // CvCallVolumeStatsIntvlType. The type is slice of
    // CISCOVOICEDIALCONTROLMIB_CvSipMsgRateStatsTable_CvSipMsgRateStatsEntry.
    CvSipMsgRateStatsEntry []*CISCOVOICEDIALCONTROLMIB_CvSipMsgRateStatsTable_CvSipMsgRateStatsEntry
}

func (cvSipMsgRateStatsTable *CISCOVOICEDIALCONTROLMIB_CvSipMsgRateStatsTable) GetEntityData() *types.CommonEntityData {
    cvSipMsgRateStatsTable.EntityData.YFilter = cvSipMsgRateStatsTable.YFilter
    cvSipMsgRateStatsTable.EntityData.YangName = "cvSipMsgRateStatsTable"
    cvSipMsgRateStatsTable.EntityData.BundleName = "cisco_ios_xe"
    cvSipMsgRateStatsTable.EntityData.ParentYangName = "CISCO-VOICE-DIAL-CONTROL-MIB"
    cvSipMsgRateStatsTable.EntityData.SegmentPath = "cvSipMsgRateStatsTable"
    cvSipMsgRateStatsTable.EntityData.AbsolutePath = "CISCO-VOICE-DIAL-CONTROL-MIB:CISCO-VOICE-DIAL-CONTROL-MIB/" + cvSipMsgRateStatsTable.EntityData.SegmentPath
    cvSipMsgRateStatsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvSipMsgRateStatsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvSipMsgRateStatsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvSipMsgRateStatsTable.EntityData.Children = types.NewOrderedMap()
    cvSipMsgRateStatsTable.EntityData.Children.Append("cvSipMsgRateStatsEntry", types.YChild{"CvSipMsgRateStatsEntry", nil})
    for i := range cvSipMsgRateStatsTable.CvSipMsgRateStatsEntry {
        cvSipMsgRateStatsTable.EntityData.Children.Append(types.GetSegmentPath(cvSipMsgRateStatsTable.CvSipMsgRateStatsEntry[i]), types.YChild{"CvSipMsgRateStatsEntry", cvSipMsgRateStatsTable.CvSipMsgRateStatsEntry[i]})
    }
    cvSipMsgRateStatsTable.EntityData.Leafs = types.NewOrderedMap()

    cvSipMsgRateStatsTable.EntityData.YListKeys = []string {}

    return &(cvSipMsgRateStatsTable.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_CvSipMsgRateStatsTable_CvSipMsgRateStatsEntry
// This is a conceptual-row in cvSipMsgRateStatsTable
// This entry is created at the system initialization and is
// updated at every epoch based on CvCallVolumeStatsIntvlType
type CISCOVOICEDIALCONTROLMIB_CvSipMsgRateStatsTable_CvSipMsgRateStatsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The Object indexes in SIP Message Rate Table to
    // select one among three interval-tables.  The different types in this table
    // are represented by  CvCallVolumeStatsIntvlType. The type is
    // CvCallVolumeStatsIntvlType.
    CvSipMsgRateStatsIntvlDurUnits interface{}

    // This attribute is a key. This is an index that references to the different
    // past periods in given in interval of SIP message rate table. This range is
    // 1-60 for Seconds and Minutes table  wherein 1-72 for hours table. The type
    // is interface{} with range: 1..72.
    CvSipMsgRateStatsIntvlDur interface{}

    // This object indicates the maximum SIP messages  per second that is received
    // for the given interval. The type is interface{} with range: 0..4294967295.
    // Units are SIP messages per second.
    CvSipMsgRateStatsMaxVal interface{}

    // This object indicates the average SIP messages per second that is received
    // for the given interval. The type is interface{} with range: 0..4294967295.
    // Units are SIP messages per second.
    CvSipMsgRateStatsAvgVal interface{}
}

func (cvSipMsgRateStatsEntry *CISCOVOICEDIALCONTROLMIB_CvSipMsgRateStatsTable_CvSipMsgRateStatsEntry) GetEntityData() *types.CommonEntityData {
    cvSipMsgRateStatsEntry.EntityData.YFilter = cvSipMsgRateStatsEntry.YFilter
    cvSipMsgRateStatsEntry.EntityData.YangName = "cvSipMsgRateStatsEntry"
    cvSipMsgRateStatsEntry.EntityData.BundleName = "cisco_ios_xe"
    cvSipMsgRateStatsEntry.EntityData.ParentYangName = "cvSipMsgRateStatsTable"
    cvSipMsgRateStatsEntry.EntityData.SegmentPath = "cvSipMsgRateStatsEntry" + types.AddKeyToken(cvSipMsgRateStatsEntry.CvSipMsgRateStatsIntvlDurUnits, "cvSipMsgRateStatsIntvlDurUnits") + types.AddKeyToken(cvSipMsgRateStatsEntry.CvSipMsgRateStatsIntvlDur, "cvSipMsgRateStatsIntvlDur")
    cvSipMsgRateStatsEntry.EntityData.AbsolutePath = "CISCO-VOICE-DIAL-CONTROL-MIB:CISCO-VOICE-DIAL-CONTROL-MIB/cvSipMsgRateStatsTable/" + cvSipMsgRateStatsEntry.EntityData.SegmentPath
    cvSipMsgRateStatsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvSipMsgRateStatsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvSipMsgRateStatsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvSipMsgRateStatsEntry.EntityData.Children = types.NewOrderedMap()
    cvSipMsgRateStatsEntry.EntityData.Leafs = types.NewOrderedMap()
    cvSipMsgRateStatsEntry.EntityData.Leafs.Append("cvSipMsgRateStatsIntvlDurUnits", types.YLeaf{"CvSipMsgRateStatsIntvlDurUnits", cvSipMsgRateStatsEntry.CvSipMsgRateStatsIntvlDurUnits})
    cvSipMsgRateStatsEntry.EntityData.Leafs.Append("cvSipMsgRateStatsIntvlDur", types.YLeaf{"CvSipMsgRateStatsIntvlDur", cvSipMsgRateStatsEntry.CvSipMsgRateStatsIntvlDur})
    cvSipMsgRateStatsEntry.EntityData.Leafs.Append("cvSipMsgRateStatsMaxVal", types.YLeaf{"CvSipMsgRateStatsMaxVal", cvSipMsgRateStatsEntry.CvSipMsgRateStatsMaxVal})
    cvSipMsgRateStatsEntry.EntityData.Leafs.Append("cvSipMsgRateStatsAvgVal", types.YLeaf{"CvSipMsgRateStatsAvgVal", cvSipMsgRateStatsEntry.CvSipMsgRateStatsAvgVal})

    cvSipMsgRateStatsEntry.EntityData.YListKeys = []string {"CvSipMsgRateStatsIntvlDurUnits", "CvSipMsgRateStatsIntvlDur"}

    return &(cvSipMsgRateStatsEntry.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_CvCallRateWMTable
// This table represents high watermarks achieved
// by call rate in various interval length defined 
// by CvCallVolumeWMIntvlType. 
// 
// Each interval may contain one or more entries to allow for 
// detailed measurement to be collected
type CISCOVOICEDIALCONTROLMIB_CvCallRateWMTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This is a conceptual-row in cvCallRateWMTable This entry is created at the
    // system initialization and is updated whenever  a) This entry is obsolete OR
    // b) A new/higher entry is available. These entries are
    // reinitialised/added/deleted  if cvCallVolumeWMTableSize is changed. The
    // type is slice of
    // CISCOVOICEDIALCONTROLMIB_CvCallRateWMTable_CvCallRateWMEntry.
    CvCallRateWMEntry []*CISCOVOICEDIALCONTROLMIB_CvCallRateWMTable_CvCallRateWMEntry
}

func (cvCallRateWMTable *CISCOVOICEDIALCONTROLMIB_CvCallRateWMTable) GetEntityData() *types.CommonEntityData {
    cvCallRateWMTable.EntityData.YFilter = cvCallRateWMTable.YFilter
    cvCallRateWMTable.EntityData.YangName = "cvCallRateWMTable"
    cvCallRateWMTable.EntityData.BundleName = "cisco_ios_xe"
    cvCallRateWMTable.EntityData.ParentYangName = "CISCO-VOICE-DIAL-CONTROL-MIB"
    cvCallRateWMTable.EntityData.SegmentPath = "cvCallRateWMTable"
    cvCallRateWMTable.EntityData.AbsolutePath = "CISCO-VOICE-DIAL-CONTROL-MIB:CISCO-VOICE-DIAL-CONTROL-MIB/" + cvCallRateWMTable.EntityData.SegmentPath
    cvCallRateWMTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvCallRateWMTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvCallRateWMTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvCallRateWMTable.EntityData.Children = types.NewOrderedMap()
    cvCallRateWMTable.EntityData.Children.Append("cvCallRateWMEntry", types.YChild{"CvCallRateWMEntry", nil})
    for i := range cvCallRateWMTable.CvCallRateWMEntry {
        cvCallRateWMTable.EntityData.Children.Append(types.GetSegmentPath(cvCallRateWMTable.CvCallRateWMEntry[i]), types.YChild{"CvCallRateWMEntry", cvCallRateWMTable.CvCallRateWMEntry[i]})
    }
    cvCallRateWMTable.EntityData.Leafs = types.NewOrderedMap()

    cvCallRateWMTable.EntityData.YListKeys = []string {}

    return &(cvCallRateWMTable.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_CvCallRateWMTable_CvCallRateWMEntry
// This is a conceptual-row in cvCallRateWMTable
// This entry is created at the system initialization and is
// updated whenever 
// a) This entry is obsolete OR
// b) A new/higher entry is available.
// These entries are reinitialised/added/deleted  if
// cvCallVolumeWMTableSize is changed
type CISCOVOICEDIALCONTROLMIB_CvCallRateWMTable_CvCallRateWMEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The Object indexes in call rate Water mark Table
    // to select one among four interval-tables.  The different types in this
    // table are represented by  CvCallVolumeWMIntvlType. The type is
    // CvCallVolumeWMIntvlType.
    CvCallRateWMIntvlDurUnits interface{}

    // This attribute is a key. This is an index that references to different
    // peaks in past period in call rate watermark table.  The number of
    // watermarks entries stored for each table are  based on
    // cvCallVolumeWMTableSize. The type is interface{} with range: 1..10.
    CvCallRateWMIndex interface{}

    // This object indicates high watermark value achieved by the calls per second
    // for the given interval. The type is interface{} with range: 0..4294967295.
    // Units are calls per second.
    CvCallRateWMValue interface{}

    // This object indicates date and Time when the high watermark is achieved for
    // calls per second for the given interval. The type is string.
    CvCallRateWMts interface{}
}

func (cvCallRateWMEntry *CISCOVOICEDIALCONTROLMIB_CvCallRateWMTable_CvCallRateWMEntry) GetEntityData() *types.CommonEntityData {
    cvCallRateWMEntry.EntityData.YFilter = cvCallRateWMEntry.YFilter
    cvCallRateWMEntry.EntityData.YangName = "cvCallRateWMEntry"
    cvCallRateWMEntry.EntityData.BundleName = "cisco_ios_xe"
    cvCallRateWMEntry.EntityData.ParentYangName = "cvCallRateWMTable"
    cvCallRateWMEntry.EntityData.SegmentPath = "cvCallRateWMEntry" + types.AddKeyToken(cvCallRateWMEntry.CvCallRateWMIntvlDurUnits, "cvCallRateWMIntvlDurUnits") + types.AddKeyToken(cvCallRateWMEntry.CvCallRateWMIndex, "cvCallRateWMIndex")
    cvCallRateWMEntry.EntityData.AbsolutePath = "CISCO-VOICE-DIAL-CONTROL-MIB:CISCO-VOICE-DIAL-CONTROL-MIB/cvCallRateWMTable/" + cvCallRateWMEntry.EntityData.SegmentPath
    cvCallRateWMEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvCallRateWMEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvCallRateWMEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvCallRateWMEntry.EntityData.Children = types.NewOrderedMap()
    cvCallRateWMEntry.EntityData.Leafs = types.NewOrderedMap()
    cvCallRateWMEntry.EntityData.Leafs.Append("cvCallRateWMIntvlDurUnits", types.YLeaf{"CvCallRateWMIntvlDurUnits", cvCallRateWMEntry.CvCallRateWMIntvlDurUnits})
    cvCallRateWMEntry.EntityData.Leafs.Append("cvCallRateWMIndex", types.YLeaf{"CvCallRateWMIndex", cvCallRateWMEntry.CvCallRateWMIndex})
    cvCallRateWMEntry.EntityData.Leafs.Append("cvCallRateWMValue", types.YLeaf{"CvCallRateWMValue", cvCallRateWMEntry.CvCallRateWMValue})
    cvCallRateWMEntry.EntityData.Leafs.Append("cvCallRateWMts", types.YLeaf{"CvCallRateWMts", cvCallRateWMEntry.CvCallRateWMts})

    cvCallRateWMEntry.EntityData.YListKeys = []string {"CvCallRateWMIntvlDurUnits", "CvCallRateWMIndex"}

    return &(cvCallRateWMEntry.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_CvCallLegRateWMTable
// cvCallLegRateWMTable table represents high watermarks achieved
// by call-leg rate in various interval length defined 
// by CvCallVolumeWMIntvlType. 
// 
// Each interval may contain one or more entries to allow for 
// detailed measurement to be collected
type CISCOVOICEDIALCONTROLMIB_CvCallLegRateWMTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This is a conceptual-row in cvCallLegRateWMTable This entry is created at
    // the system initialization and is updated whenever  a) This entry is
    // obsolete OR b) A new/higher entry is available. These entries are
    // reinitialised/added/deleted  if cvCallVolumeWMTableSize is changed. The
    // type is slice of
    // CISCOVOICEDIALCONTROLMIB_CvCallLegRateWMTable_CvCallLegRateWMEntry.
    CvCallLegRateWMEntry []*CISCOVOICEDIALCONTROLMIB_CvCallLegRateWMTable_CvCallLegRateWMEntry
}

func (cvCallLegRateWMTable *CISCOVOICEDIALCONTROLMIB_CvCallLegRateWMTable) GetEntityData() *types.CommonEntityData {
    cvCallLegRateWMTable.EntityData.YFilter = cvCallLegRateWMTable.YFilter
    cvCallLegRateWMTable.EntityData.YangName = "cvCallLegRateWMTable"
    cvCallLegRateWMTable.EntityData.BundleName = "cisco_ios_xe"
    cvCallLegRateWMTable.EntityData.ParentYangName = "CISCO-VOICE-DIAL-CONTROL-MIB"
    cvCallLegRateWMTable.EntityData.SegmentPath = "cvCallLegRateWMTable"
    cvCallLegRateWMTable.EntityData.AbsolutePath = "CISCO-VOICE-DIAL-CONTROL-MIB:CISCO-VOICE-DIAL-CONTROL-MIB/" + cvCallLegRateWMTable.EntityData.SegmentPath
    cvCallLegRateWMTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvCallLegRateWMTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvCallLegRateWMTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvCallLegRateWMTable.EntityData.Children = types.NewOrderedMap()
    cvCallLegRateWMTable.EntityData.Children.Append("cvCallLegRateWMEntry", types.YChild{"CvCallLegRateWMEntry", nil})
    for i := range cvCallLegRateWMTable.CvCallLegRateWMEntry {
        cvCallLegRateWMTable.EntityData.Children.Append(types.GetSegmentPath(cvCallLegRateWMTable.CvCallLegRateWMEntry[i]), types.YChild{"CvCallLegRateWMEntry", cvCallLegRateWMTable.CvCallLegRateWMEntry[i]})
    }
    cvCallLegRateWMTable.EntityData.Leafs = types.NewOrderedMap()

    cvCallLegRateWMTable.EntityData.YListKeys = []string {}

    return &(cvCallLegRateWMTable.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_CvCallLegRateWMTable_CvCallLegRateWMEntry
// This is a conceptual-row in cvCallLegRateWMTable
// This entry is created at the system initialization and is
// updated whenever 
// a) This entry is obsolete OR
// b) A new/higher entry is available.
// These entries are reinitialised/added/deleted  if
// cvCallVolumeWMTableSize is changed
type CISCOVOICEDIALCONTROLMIB_CvCallLegRateWMTable_CvCallLegRateWMEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The Object indexes in call leg rate Water mark
    // Table to select one among four interval-tables.  The different types in
    // this table are represented by  CvCallVolumeWMIntvlType. The type is
    // CvCallVolumeWMIntvlType.
    CvCallLegRateWMIntvlDurUnits interface{}

    // This attribute is a key. This is an index that references to different
    // peaks in past period in call leg rate watermark table.  The number of
    // watermarks entries stored for each table are  based on
    // cvCallVolumeWMTableSize. The type is interface{} with range: 1..10.
    CvCallLegRateWMIndex interface{}

    // This object indicates high watermark value achieved by the call legs per
    // second for the given interval. The type is interface{} with range:
    // 0..4294967295. Units are call legs per second.
    CvCallLegRateWMValue interface{}

    // This object indicates date and time when the high watermark is achieved for
    // call-legs per second for the given interval. The type is string.
    CvCallLegRateWMts interface{}
}

func (cvCallLegRateWMEntry *CISCOVOICEDIALCONTROLMIB_CvCallLegRateWMTable_CvCallLegRateWMEntry) GetEntityData() *types.CommonEntityData {
    cvCallLegRateWMEntry.EntityData.YFilter = cvCallLegRateWMEntry.YFilter
    cvCallLegRateWMEntry.EntityData.YangName = "cvCallLegRateWMEntry"
    cvCallLegRateWMEntry.EntityData.BundleName = "cisco_ios_xe"
    cvCallLegRateWMEntry.EntityData.ParentYangName = "cvCallLegRateWMTable"
    cvCallLegRateWMEntry.EntityData.SegmentPath = "cvCallLegRateWMEntry" + types.AddKeyToken(cvCallLegRateWMEntry.CvCallLegRateWMIntvlDurUnits, "cvCallLegRateWMIntvlDurUnits") + types.AddKeyToken(cvCallLegRateWMEntry.CvCallLegRateWMIndex, "cvCallLegRateWMIndex")
    cvCallLegRateWMEntry.EntityData.AbsolutePath = "CISCO-VOICE-DIAL-CONTROL-MIB:CISCO-VOICE-DIAL-CONTROL-MIB/cvCallLegRateWMTable/" + cvCallLegRateWMEntry.EntityData.SegmentPath
    cvCallLegRateWMEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvCallLegRateWMEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvCallLegRateWMEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvCallLegRateWMEntry.EntityData.Children = types.NewOrderedMap()
    cvCallLegRateWMEntry.EntityData.Leafs = types.NewOrderedMap()
    cvCallLegRateWMEntry.EntityData.Leafs.Append("cvCallLegRateWMIntvlDurUnits", types.YLeaf{"CvCallLegRateWMIntvlDurUnits", cvCallLegRateWMEntry.CvCallLegRateWMIntvlDurUnits})
    cvCallLegRateWMEntry.EntityData.Leafs.Append("cvCallLegRateWMIndex", types.YLeaf{"CvCallLegRateWMIndex", cvCallLegRateWMEntry.CvCallLegRateWMIndex})
    cvCallLegRateWMEntry.EntityData.Leafs.Append("cvCallLegRateWMValue", types.YLeaf{"CvCallLegRateWMValue", cvCallLegRateWMEntry.CvCallLegRateWMValue})
    cvCallLegRateWMEntry.EntityData.Leafs.Append("cvCallLegRateWMts", types.YLeaf{"CvCallLegRateWMts", cvCallLegRateWMEntry.CvCallLegRateWMts})

    cvCallLegRateWMEntry.EntityData.YListKeys = []string {"CvCallLegRateWMIntvlDurUnits", "CvCallLegRateWMIndex"}

    return &(cvCallLegRateWMEntry.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_CvActiveCallWMTable
// This table represents high watermarks achieved
// by active calls in various interval length defined 
// by CvCallVolumeWMIntvlType. 
// 
// Each interval may contain one or more entries to allow 
// for detailed measurement to be collected.
type CISCOVOICEDIALCONTROLMIB_CvActiveCallWMTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This is a conceptual-row in cvActiveCallWMTable This entry is created at
    // the system initialization and is updated whenever  a) This entry is
    // obsolete OR b) A new/higher entry is available. These entries are
    // reinitialised/added/deleted  if cvCallVolumeWMTableSize is changed. The
    // type is slice of
    // CISCOVOICEDIALCONTROLMIB_CvActiveCallWMTable_CvActiveCallWMEntry.
    CvActiveCallWMEntry []*CISCOVOICEDIALCONTROLMIB_CvActiveCallWMTable_CvActiveCallWMEntry
}

func (cvActiveCallWMTable *CISCOVOICEDIALCONTROLMIB_CvActiveCallWMTable) GetEntityData() *types.CommonEntityData {
    cvActiveCallWMTable.EntityData.YFilter = cvActiveCallWMTable.YFilter
    cvActiveCallWMTable.EntityData.YangName = "cvActiveCallWMTable"
    cvActiveCallWMTable.EntityData.BundleName = "cisco_ios_xe"
    cvActiveCallWMTable.EntityData.ParentYangName = "CISCO-VOICE-DIAL-CONTROL-MIB"
    cvActiveCallWMTable.EntityData.SegmentPath = "cvActiveCallWMTable"
    cvActiveCallWMTable.EntityData.AbsolutePath = "CISCO-VOICE-DIAL-CONTROL-MIB:CISCO-VOICE-DIAL-CONTROL-MIB/" + cvActiveCallWMTable.EntityData.SegmentPath
    cvActiveCallWMTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvActiveCallWMTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvActiveCallWMTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvActiveCallWMTable.EntityData.Children = types.NewOrderedMap()
    cvActiveCallWMTable.EntityData.Children.Append("cvActiveCallWMEntry", types.YChild{"CvActiveCallWMEntry", nil})
    for i := range cvActiveCallWMTable.CvActiveCallWMEntry {
        cvActiveCallWMTable.EntityData.Children.Append(types.GetSegmentPath(cvActiveCallWMTable.CvActiveCallWMEntry[i]), types.YChild{"CvActiveCallWMEntry", cvActiveCallWMTable.CvActiveCallWMEntry[i]})
    }
    cvActiveCallWMTable.EntityData.Leafs = types.NewOrderedMap()

    cvActiveCallWMTable.EntityData.YListKeys = []string {}

    return &(cvActiveCallWMTable.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_CvActiveCallWMTable_CvActiveCallWMEntry
// This is a conceptual-row in cvActiveCallWMTable
// This entry is created at the system initialization and is
// updated whenever 
// a) This entry is obsolete OR
// b) A new/higher entry is available.
// These entries are reinitialised/added/deleted  if
// cvCallVolumeWMTableSize is changed
type CISCOVOICEDIALCONTROLMIB_CvActiveCallWMTable_CvActiveCallWMEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The Object indexes in active call Water mark Table
    // to select one among four interval-tables.  The different types in this
    // table are represented by  CvCallVolumeWMIntvlType. The type is
    // CvCallVolumeWMIntvlType.
    CvActiveCallWMIntvlDurUnits interface{}

    // This attribute is a key. This is an index that references to different
    // peaks in past period in acive call watermark table.  The number of
    // watermarks entries stored for each table are  based on
    // cvCallVolumeWMTableSize. The type is interface{} with range: 1..10.
    CvActiveCallWMIndex interface{}

    // This object indicates high watermark value achieved by the active calls for
    // the given interval. The type is interface{} with range: 0..4294967295.
    // Units are calls.
    CvActiveCallWMValue interface{}

    // This object indicates date and time when the high watermark is achieved for
    // active calls for the given interval. The type is string.
    CvActiveCallWMts interface{}
}

func (cvActiveCallWMEntry *CISCOVOICEDIALCONTROLMIB_CvActiveCallWMTable_CvActiveCallWMEntry) GetEntityData() *types.CommonEntityData {
    cvActiveCallWMEntry.EntityData.YFilter = cvActiveCallWMEntry.YFilter
    cvActiveCallWMEntry.EntityData.YangName = "cvActiveCallWMEntry"
    cvActiveCallWMEntry.EntityData.BundleName = "cisco_ios_xe"
    cvActiveCallWMEntry.EntityData.ParentYangName = "cvActiveCallWMTable"
    cvActiveCallWMEntry.EntityData.SegmentPath = "cvActiveCallWMEntry" + types.AddKeyToken(cvActiveCallWMEntry.CvActiveCallWMIntvlDurUnits, "cvActiveCallWMIntvlDurUnits") + types.AddKeyToken(cvActiveCallWMEntry.CvActiveCallWMIndex, "cvActiveCallWMIndex")
    cvActiveCallWMEntry.EntityData.AbsolutePath = "CISCO-VOICE-DIAL-CONTROL-MIB:CISCO-VOICE-DIAL-CONTROL-MIB/cvActiveCallWMTable/" + cvActiveCallWMEntry.EntityData.SegmentPath
    cvActiveCallWMEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvActiveCallWMEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvActiveCallWMEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvActiveCallWMEntry.EntityData.Children = types.NewOrderedMap()
    cvActiveCallWMEntry.EntityData.Leafs = types.NewOrderedMap()
    cvActiveCallWMEntry.EntityData.Leafs.Append("cvActiveCallWMIntvlDurUnits", types.YLeaf{"CvActiveCallWMIntvlDurUnits", cvActiveCallWMEntry.CvActiveCallWMIntvlDurUnits})
    cvActiveCallWMEntry.EntityData.Leafs.Append("cvActiveCallWMIndex", types.YLeaf{"CvActiveCallWMIndex", cvActiveCallWMEntry.CvActiveCallWMIndex})
    cvActiveCallWMEntry.EntityData.Leafs.Append("cvActiveCallWMValue", types.YLeaf{"CvActiveCallWMValue", cvActiveCallWMEntry.CvActiveCallWMValue})
    cvActiveCallWMEntry.EntityData.Leafs.Append("cvActiveCallWMts", types.YLeaf{"CvActiveCallWMts", cvActiveCallWMEntry.CvActiveCallWMts})

    cvActiveCallWMEntry.EntityData.YListKeys = []string {"CvActiveCallWMIntvlDurUnits", "CvActiveCallWMIndex"}

    return &(cvActiveCallWMEntry.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_CvSipMsgRateWMTable
// This table represents of high watermarks achieved
// by SIP message rate in various interval length defined 
// by CvCallVolumeWMIntvlType. 
// 
// Each interval may contain one or more entries to allow for
// detailed measurement to be collected
type CISCOVOICEDIALCONTROLMIB_CvSipMsgRateWMTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This is a conceptual-row in cvSipMsgRateWMTable. This entry is created at
    // the system initialization and is updated whenever  a) This entry is
    // obsolete OR b) A new/higher entry is available. These entries are
    // reinitialised/added/deleted if cvCallVolumeWMTableSize is changed. The type
    // is slice of
    // CISCOVOICEDIALCONTROLMIB_CvSipMsgRateWMTable_CvSipMsgRateWMEntry.
    CvSipMsgRateWMEntry []*CISCOVOICEDIALCONTROLMIB_CvSipMsgRateWMTable_CvSipMsgRateWMEntry
}

func (cvSipMsgRateWMTable *CISCOVOICEDIALCONTROLMIB_CvSipMsgRateWMTable) GetEntityData() *types.CommonEntityData {
    cvSipMsgRateWMTable.EntityData.YFilter = cvSipMsgRateWMTable.YFilter
    cvSipMsgRateWMTable.EntityData.YangName = "cvSipMsgRateWMTable"
    cvSipMsgRateWMTable.EntityData.BundleName = "cisco_ios_xe"
    cvSipMsgRateWMTable.EntityData.ParentYangName = "CISCO-VOICE-DIAL-CONTROL-MIB"
    cvSipMsgRateWMTable.EntityData.SegmentPath = "cvSipMsgRateWMTable"
    cvSipMsgRateWMTable.EntityData.AbsolutePath = "CISCO-VOICE-DIAL-CONTROL-MIB:CISCO-VOICE-DIAL-CONTROL-MIB/" + cvSipMsgRateWMTable.EntityData.SegmentPath
    cvSipMsgRateWMTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvSipMsgRateWMTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvSipMsgRateWMTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvSipMsgRateWMTable.EntityData.Children = types.NewOrderedMap()
    cvSipMsgRateWMTable.EntityData.Children.Append("cvSipMsgRateWMEntry", types.YChild{"CvSipMsgRateWMEntry", nil})
    for i := range cvSipMsgRateWMTable.CvSipMsgRateWMEntry {
        cvSipMsgRateWMTable.EntityData.Children.Append(types.GetSegmentPath(cvSipMsgRateWMTable.CvSipMsgRateWMEntry[i]), types.YChild{"CvSipMsgRateWMEntry", cvSipMsgRateWMTable.CvSipMsgRateWMEntry[i]})
    }
    cvSipMsgRateWMTable.EntityData.Leafs = types.NewOrderedMap()

    cvSipMsgRateWMTable.EntityData.YListKeys = []string {}

    return &(cvSipMsgRateWMTable.EntityData)
}

// CISCOVOICEDIALCONTROLMIB_CvSipMsgRateWMTable_CvSipMsgRateWMEntry
// This is a conceptual-row in cvSipMsgRateWMTable.
// This entry is created at the system initialization and is
// updated whenever 
// a) This entry is obsolete OR
// b) A new/higher entry is available.
// These entries are reinitialised/added/deleted if
// cvCallVolumeWMTableSize is changed
type CISCOVOICEDIALCONTROLMIB_CvSipMsgRateWMTable_CvSipMsgRateWMEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The Object indexes in SIP Message rate Water mark
    // Table to select one among four interval-tables.  The different types in
    // this table are represented by  CvCallVolumeWMIntvlType. The type is
    // CvCallVolumeWMIntvlType.
    CvSipMsgRateWMIntvlDurUnits interface{}

    // This attribute is a key. This is an index that references to different
    // peaks in past period in sip message rate watermark table.  The number of
    // watermarks entries stored for each table are  based on
    // cvCallVolumeWMTableSize. The type is interface{} with range: 1..10.
    CvSipMsgRateWMIndex interface{}

    // This object indicates high watermark value achieved by the SIP messages per
    // second for the given interval. The type is interface{} with range:
    // 0..4294967295. Units are SIP messages per second.
    CvSipMsgRateWMValue interface{}

    // This object indicates date and time when the high watermark is achieved for
    // SIP messages per second for the given interval. The type is string.
    CvSipMsgRateWMts interface{}
}

func (cvSipMsgRateWMEntry *CISCOVOICEDIALCONTROLMIB_CvSipMsgRateWMTable_CvSipMsgRateWMEntry) GetEntityData() *types.CommonEntityData {
    cvSipMsgRateWMEntry.EntityData.YFilter = cvSipMsgRateWMEntry.YFilter
    cvSipMsgRateWMEntry.EntityData.YangName = "cvSipMsgRateWMEntry"
    cvSipMsgRateWMEntry.EntityData.BundleName = "cisco_ios_xe"
    cvSipMsgRateWMEntry.EntityData.ParentYangName = "cvSipMsgRateWMTable"
    cvSipMsgRateWMEntry.EntityData.SegmentPath = "cvSipMsgRateWMEntry" + types.AddKeyToken(cvSipMsgRateWMEntry.CvSipMsgRateWMIntvlDurUnits, "cvSipMsgRateWMIntvlDurUnits") + types.AddKeyToken(cvSipMsgRateWMEntry.CvSipMsgRateWMIndex, "cvSipMsgRateWMIndex")
    cvSipMsgRateWMEntry.EntityData.AbsolutePath = "CISCO-VOICE-DIAL-CONTROL-MIB:CISCO-VOICE-DIAL-CONTROL-MIB/cvSipMsgRateWMTable/" + cvSipMsgRateWMEntry.EntityData.SegmentPath
    cvSipMsgRateWMEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvSipMsgRateWMEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvSipMsgRateWMEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvSipMsgRateWMEntry.EntityData.Children = types.NewOrderedMap()
    cvSipMsgRateWMEntry.EntityData.Leafs = types.NewOrderedMap()
    cvSipMsgRateWMEntry.EntityData.Leafs.Append("cvSipMsgRateWMIntvlDurUnits", types.YLeaf{"CvSipMsgRateWMIntvlDurUnits", cvSipMsgRateWMEntry.CvSipMsgRateWMIntvlDurUnits})
    cvSipMsgRateWMEntry.EntityData.Leafs.Append("cvSipMsgRateWMIndex", types.YLeaf{"CvSipMsgRateWMIndex", cvSipMsgRateWMEntry.CvSipMsgRateWMIndex})
    cvSipMsgRateWMEntry.EntityData.Leafs.Append("cvSipMsgRateWMValue", types.YLeaf{"CvSipMsgRateWMValue", cvSipMsgRateWMEntry.CvSipMsgRateWMValue})
    cvSipMsgRateWMEntry.EntityData.Leafs.Append("cvSipMsgRateWMts", types.YLeaf{"CvSipMsgRateWMts", cvSipMsgRateWMEntry.CvSipMsgRateWMts})

    cvSipMsgRateWMEntry.EntityData.YListKeys = []string {"CvSipMsgRateWMIntvlDurUnits", "CvSipMsgRateWMIndex"}

    return &(cvSipMsgRateWMEntry.EntityData)
}

