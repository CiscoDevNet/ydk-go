// This MIB module contains voice related objects that
// are common across more than one network
// encapsulation i.e VoIP, VoATM and VoFR.
// 
// *** ABBREVIATIONS, ACRONYMS AND SYMBOLS ***
// 
// GSM    - Global System for Mobile Communication
// 
// AMR-NB - Adaptive Multi Rate - Narrow Band
// 
// iLBC   - internet Low Bitrate Codec
// 
// iSAC   - internet Speech Audio Codec
package cisco_voice_common_dial_control_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_voice_common_dial_control_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-VOICE-COMMON-DIAL-CONTROL-MIB CISCO-VOICE-COMMON-DIAL-CONTROL-MIB}", reflect.TypeOf(CISCOVOICECOMMONDIALCONTROLMIB{}))
    ydk.RegisterEntity("CISCO-VOICE-COMMON-DIAL-CONTROL-MIB:CISCO-VOICE-COMMON-DIAL-CONTROL-MIB", reflect.TypeOf(CISCOVOICECOMMONDIALCONTROLMIB{}))
}

// CvcSpeechCoderRate represents aacld          - AAC-LD MPEG-4 Low Delay Audio Coder
type CvcSpeechCoderRate string

const (
    CvcSpeechCoderRate_g729r8000 CvcSpeechCoderRate = "g729r8000"

    CvcSpeechCoderRate_g729Ar8000 CvcSpeechCoderRate = "g729Ar8000"

    CvcSpeechCoderRate_g726r16000 CvcSpeechCoderRate = "g726r16000"

    CvcSpeechCoderRate_g726r24000 CvcSpeechCoderRate = "g726r24000"

    CvcSpeechCoderRate_g726r32000 CvcSpeechCoderRate = "g726r32000"

    CvcSpeechCoderRate_g711ulawr64000 CvcSpeechCoderRate = "g711ulawr64000"

    CvcSpeechCoderRate_g711Alawr64000 CvcSpeechCoderRate = "g711Alawr64000"

    CvcSpeechCoderRate_g728r16000 CvcSpeechCoderRate = "g728r16000"

    CvcSpeechCoderRate_g723r6300 CvcSpeechCoderRate = "g723r6300"

    CvcSpeechCoderRate_g723r5300 CvcSpeechCoderRate = "g723r5300"

    CvcSpeechCoderRate_gsmr13200 CvcSpeechCoderRate = "gsmr13200"

    CvcSpeechCoderRate_g729Br8000 CvcSpeechCoderRate = "g729Br8000"

    CvcSpeechCoderRate_g729ABr8000 CvcSpeechCoderRate = "g729ABr8000"

    CvcSpeechCoderRate_g723Ar6300 CvcSpeechCoderRate = "g723Ar6300"

    CvcSpeechCoderRate_g723Ar5300 CvcSpeechCoderRate = "g723Ar5300"

    CvcSpeechCoderRate_g729IETFr8000 CvcSpeechCoderRate = "g729IETFr8000"

    CvcSpeechCoderRate_gsmeEr12200 CvcSpeechCoderRate = "gsmeEr12200"

    CvcSpeechCoderRate_clearChannel CvcSpeechCoderRate = "clearChannel"

    CvcSpeechCoderRate_g726r40000 CvcSpeechCoderRate = "g726r40000"

    CvcSpeechCoderRate_llcc CvcSpeechCoderRate = "llcc"

    CvcSpeechCoderRate_gsmAmrNb CvcSpeechCoderRate = "gsmAmrNb"

    CvcSpeechCoderRate_iLBC CvcSpeechCoderRate = "iLBC"

    CvcSpeechCoderRate_iLBCr15200 CvcSpeechCoderRate = "iLBCr15200"

    CvcSpeechCoderRate_iLBCr13330 CvcSpeechCoderRate = "iLBCr13330"

    CvcSpeechCoderRate_g722r4800 CvcSpeechCoderRate = "g722r4800"

    CvcSpeechCoderRate_g722r5600 CvcSpeechCoderRate = "g722r5600"

    CvcSpeechCoderRate_g722r6400 CvcSpeechCoderRate = "g722r6400"

    CvcSpeechCoderRate_iSAC CvcSpeechCoderRate = "iSAC"

    CvcSpeechCoderRate_aaclc CvcSpeechCoderRate = "aaclc"

    CvcSpeechCoderRate_aacld CvcSpeechCoderRate = "aacld"
)

// CvcFaxTransmitRate represents fax12000  - 12000 bps FAX transmit rate.
type CvcFaxTransmitRate string

const (
    CvcFaxTransmitRate_none CvcFaxTransmitRate = "none"

    CvcFaxTransmitRate_voiceRate CvcFaxTransmitRate = "voiceRate"

    CvcFaxTransmitRate_fax2400 CvcFaxTransmitRate = "fax2400"

    CvcFaxTransmitRate_fax4800 CvcFaxTransmitRate = "fax4800"

    CvcFaxTransmitRate_fax7200 CvcFaxTransmitRate = "fax7200"

    CvcFaxTransmitRate_fax9600 CvcFaxTransmitRate = "fax9600"

    CvcFaxTransmitRate_fax14400 CvcFaxTransmitRate = "fax14400"

    CvcFaxTransmitRate_fax12000 CvcFaxTransmitRate = "fax12000"
)

// CvcCoderTypeRate represents aacld          - AAC-LD MPEG-4 Low Delay Audio Coder
type CvcCoderTypeRate string

const (
    CvcCoderTypeRate_other CvcCoderTypeRate = "other"

    CvcCoderTypeRate_fax2400 CvcCoderTypeRate = "fax2400"

    CvcCoderTypeRate_fax4800 CvcCoderTypeRate = "fax4800"

    CvcCoderTypeRate_fax7200 CvcCoderTypeRate = "fax7200"

    CvcCoderTypeRate_fax9600 CvcCoderTypeRate = "fax9600"

    CvcCoderTypeRate_fax14400 CvcCoderTypeRate = "fax14400"

    CvcCoderTypeRate_fax12000 CvcCoderTypeRate = "fax12000"

    CvcCoderTypeRate_g729r8000 CvcCoderTypeRate = "g729r8000"

    CvcCoderTypeRate_g729Ar8000 CvcCoderTypeRate = "g729Ar8000"

    CvcCoderTypeRate_g726r16000 CvcCoderTypeRate = "g726r16000"

    CvcCoderTypeRate_g726r24000 CvcCoderTypeRate = "g726r24000"

    CvcCoderTypeRate_g726r32000 CvcCoderTypeRate = "g726r32000"

    CvcCoderTypeRate_g711ulawr64000 CvcCoderTypeRate = "g711ulawr64000"

    CvcCoderTypeRate_g711Alawr64000 CvcCoderTypeRate = "g711Alawr64000"

    CvcCoderTypeRate_g728r16000 CvcCoderTypeRate = "g728r16000"

    CvcCoderTypeRate_g723r6300 CvcCoderTypeRate = "g723r6300"

    CvcCoderTypeRate_g723r5300 CvcCoderTypeRate = "g723r5300"

    CvcCoderTypeRate_gsmr13200 CvcCoderTypeRate = "gsmr13200"

    CvcCoderTypeRate_g729Br8000 CvcCoderTypeRate = "g729Br8000"

    CvcCoderTypeRate_g729ABr8000 CvcCoderTypeRate = "g729ABr8000"

    CvcCoderTypeRate_g723Ar6300 CvcCoderTypeRate = "g723Ar6300"

    CvcCoderTypeRate_g723Ar5300 CvcCoderTypeRate = "g723Ar5300"

    CvcCoderTypeRate_ietfg729r8000 CvcCoderTypeRate = "ietfg729r8000"

    CvcCoderTypeRate_gsmeEr12200 CvcCoderTypeRate = "gsmeEr12200"

    CvcCoderTypeRate_clearChannel CvcCoderTypeRate = "clearChannel"

    CvcCoderTypeRate_g726r40000 CvcCoderTypeRate = "g726r40000"

    CvcCoderTypeRate_llcc CvcCoderTypeRate = "llcc"

    CvcCoderTypeRate_gsmAmrNb CvcCoderTypeRate = "gsmAmrNb"

    CvcCoderTypeRate_g722 CvcCoderTypeRate = "g722"

    CvcCoderTypeRate_iLBC CvcCoderTypeRate = "iLBC"

    CvcCoderTypeRate_iLBCr15200 CvcCoderTypeRate = "iLBCr15200"

    CvcCoderTypeRate_iLBCr13330 CvcCoderTypeRate = "iLBCr13330"

    CvcCoderTypeRate_g722r4800 CvcCoderTypeRate = "g722r4800"

    CvcCoderTypeRate_g722r5600 CvcCoderTypeRate = "g722r5600"

    CvcCoderTypeRate_g722r6400 CvcCoderTypeRate = "g722r6400"

    CvcCoderTypeRate_iSAC CvcCoderTypeRate = "iSAC"

    CvcCoderTypeRate_aaclc CvcCoderTypeRate = "aaclc"

    CvcCoderTypeRate_aacld CvcCoderTypeRate = "aacld"
)

// CvcInBandSignaling represents               GR303
type CvcInBandSignaling string

const (
    CvcInBandSignaling_cas CvcInBandSignaling = "cas"

    CvcInBandSignaling_none CvcInBandSignaling = "none"

    CvcInBandSignaling_cept CvcInBandSignaling = "cept"

    CvcInBandSignaling_transparent CvcInBandSignaling = "transparent"

    CvcInBandSignaling_gr303 CvcInBandSignaling = "gr303"
)

// CvcH320CallType represents This object specifies the H320 call type of a voice call.
type CvcH320CallType string

const (
    CvcH320CallType_none CvcH320CallType = "none"

    CvcH320CallType_primary CvcH320CallType = "primary"

    CvcH320CallType_secondary CvcH320CallType = "secondary"
)

// CvcVideoCoderRate represents the video data of the voice call.
type CvcVideoCoderRate string

const (
    CvcVideoCoderRate_none CvcVideoCoderRate = "none"

    CvcVideoCoderRate_h261 CvcVideoCoderRate = "h261"

    CvcVideoCoderRate_h263 CvcVideoCoderRate = "h263"

    CvcVideoCoderRate_h263plus CvcVideoCoderRate = "h263plus"

    CvcVideoCoderRate_h264 CvcVideoCoderRate = "h264"
)

// CISCOVOICECOMMONDIALCONTROLMIB
type CISCOVOICECOMMONDIALCONTROLMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This table is a common extension to the call active table of IETF Dial
    // Control MIB. It contains common call  leg information about a network call
    // leg.
    Cvcommondccallactivetable CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallactivetable

    // This table is the Common extension to the call history table of IETF Dial
    // Control MIB. It contains Common call  leg information about a network call
    // leg.
    Cvcommondccallhistorytable CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallhistorytable
}

func (cISCOVOICECOMMONDIALCONTROLMIB *CISCOVOICECOMMONDIALCONTROLMIB) GetEntityData() *types.CommonEntityData {
    cISCOVOICECOMMONDIALCONTROLMIB.EntityData.YFilter = cISCOVOICECOMMONDIALCONTROLMIB.YFilter
    cISCOVOICECOMMONDIALCONTROLMIB.EntityData.YangName = "CISCO-VOICE-COMMON-DIAL-CONTROL-MIB"
    cISCOVOICECOMMONDIALCONTROLMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOVOICECOMMONDIALCONTROLMIB.EntityData.ParentYangName = "CISCO-VOICE-COMMON-DIAL-CONTROL-MIB"
    cISCOVOICECOMMONDIALCONTROLMIB.EntityData.SegmentPath = "CISCO-VOICE-COMMON-DIAL-CONTROL-MIB:CISCO-VOICE-COMMON-DIAL-CONTROL-MIB"
    cISCOVOICECOMMONDIALCONTROLMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOVOICECOMMONDIALCONTROLMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOVOICECOMMONDIALCONTROLMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOVOICECOMMONDIALCONTROLMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOVOICECOMMONDIALCONTROLMIB.EntityData.Children["cvCommonDcCallActiveTable"] = types.YChild{"Cvcommondccallactivetable", &cISCOVOICECOMMONDIALCONTROLMIB.Cvcommondccallactivetable}
    cISCOVOICECOMMONDIALCONTROLMIB.EntityData.Children["cvCommonDcCallHistoryTable"] = types.YChild{"Cvcommondccallhistorytable", &cISCOVOICECOMMONDIALCONTROLMIB.Cvcommondccallhistorytable}
    cISCOVOICECOMMONDIALCONTROLMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOVOICECOMMONDIALCONTROLMIB.EntityData)
}

// CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallactivetable
// This table is a common extension to the call active
// table of IETF Dial Control MIB. It contains common call 
// leg information about a network call leg.
type CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallactivetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The common information regarding a single network call leg. The call leg
    // entry is identified by using the same  index objects that are used by Call
    // Active table of IETF  Dial Control MIB to identify the call. An entry of
    // this table is created when its associated  call active entry in the IETF
    // Dial Control MIB is created and the call active entry contains information
    // for the  call establishment to a network dialpeer.              The entry
    // is deleted when its associated call active entry  in the IETF Dial Control
    // MIB is deleted. The type is slice of
    // CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallactivetable_Cvcommondccallactiveentry.
    Cvcommondccallactiveentry []CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallactivetable_Cvcommondccallactiveentry
}

func (cvcommondccallactivetable *CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallactivetable) GetEntityData() *types.CommonEntityData {
    cvcommondccallactivetable.EntityData.YFilter = cvcommondccallactivetable.YFilter
    cvcommondccallactivetable.EntityData.YangName = "cvCommonDcCallActiveTable"
    cvcommondccallactivetable.EntityData.BundleName = "cisco_ios_xe"
    cvcommondccallactivetable.EntityData.ParentYangName = "CISCO-VOICE-COMMON-DIAL-CONTROL-MIB"
    cvcommondccallactivetable.EntityData.SegmentPath = "cvCommonDcCallActiveTable"
    cvcommondccallactivetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvcommondccallactivetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvcommondccallactivetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvcommondccallactivetable.EntityData.Children = make(map[string]types.YChild)
    cvcommondccallactivetable.EntityData.Children["cvCommonDcCallActiveEntry"] = types.YChild{"Cvcommondccallactiveentry", nil}
    for i := range cvcommondccallactivetable.Cvcommondccallactiveentry {
        cvcommondccallactivetable.EntityData.Children[types.GetSegmentPath(&cvcommondccallactivetable.Cvcommondccallactiveentry[i])] = types.YChild{"Cvcommondccallactiveentry", &cvcommondccallactivetable.Cvcommondccallactiveentry[i]}
    }
    cvcommondccallactivetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cvcommondccallactivetable.EntityData)
}

// CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallactivetable_Cvcommondccallactiveentry
// The common information regarding a single network call
// leg. The call leg entry is identified by using the same 
// index objects that are used by Call Active table of IETF 
// Dial Control MIB to identify the call.
// An entry of this table is created when its associated 
// call active entry in the IETF Dial Control MIB is created
// and the call active entry contains information for the 
// call establishment to a network dialpeer.             
// The entry is deleted when its associated call active entry 
// in the IETF Dial Control MIB is deleted.
type CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallactivetable_Cvcommondccallactiveentry struct {
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

    // The global call identifier for the network call. The type is string with
    // length: 0..16.
    Cvcommondccallactiveconnectionid interface{}

    // The object indicates whether or not the VAD (Voice Activity Detection) is
    // enabled for the voice call. The type is bool.
    Cvcommondccallactivevadenable interface{}

    // The negotiated coder rate. It specifies the transmit rate of voice/fax
    // compression to its associated call leg  for the call. This rate is
    // different from the configuration rate  because of rate negotiation during
    // the call. The type is CvcCoderTypeRate.
    Cvcommondccallactivecodertyperate interface{}

    // Specifies the payload size of the voice packet. The type is interface{}
    // with range: 10..255.
    Cvcommondccallactivecodecbytes interface{}

    // Specifies the type of in-band signaling being used on the call.  This
    // object is instantiated only for  Connection Trunk calls. The type is
    // CvcInBandSignaling.
    Cvcommondccallactiveinbandsignaling interface{}

    // The calling party name this call is connected to. If the name is not
    // available, then it will have a length of zero. The type is string with
    // length: 0..64.
    Cvcommondccallactivecallingname interface{}

    // The object indicates whether or not the caller ID feature is blocked for
    // this voice call. The type is bool.
    Cvcommondccallactivecalleridblock interface{}
}

func (cvcommondccallactiveentry *CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallactivetable_Cvcommondccallactiveentry) GetEntityData() *types.CommonEntityData {
    cvcommondccallactiveentry.EntityData.YFilter = cvcommondccallactiveentry.YFilter
    cvcommondccallactiveentry.EntityData.YangName = "cvCommonDcCallActiveEntry"
    cvcommondccallactiveentry.EntityData.BundleName = "cisco_ios_xe"
    cvcommondccallactiveentry.EntityData.ParentYangName = "cvCommonDcCallActiveTable"
    cvcommondccallactiveentry.EntityData.SegmentPath = "cvCommonDcCallActiveEntry" + "[callActiveSetupTime='" + fmt.Sprintf("%v", cvcommondccallactiveentry.Callactivesetuptime) + "']" + "[callActiveIndex='" + fmt.Sprintf("%v", cvcommondccallactiveentry.Callactiveindex) + "']"
    cvcommondccallactiveentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvcommondccallactiveentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvcommondccallactiveentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvcommondccallactiveentry.EntityData.Children = make(map[string]types.YChild)
    cvcommondccallactiveentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cvcommondccallactiveentry.EntityData.Leafs["callActiveSetupTime"] = types.YLeaf{"Callactivesetuptime", cvcommondccallactiveentry.Callactivesetuptime}
    cvcommondccallactiveentry.EntityData.Leafs["callActiveIndex"] = types.YLeaf{"Callactiveindex", cvcommondccallactiveentry.Callactiveindex}
    cvcommondccallactiveentry.EntityData.Leafs["cvCommonDcCallActiveConnectionId"] = types.YLeaf{"Cvcommondccallactiveconnectionid", cvcommondccallactiveentry.Cvcommondccallactiveconnectionid}
    cvcommondccallactiveentry.EntityData.Leafs["cvCommonDcCallActiveVADEnable"] = types.YLeaf{"Cvcommondccallactivevadenable", cvcommondccallactiveentry.Cvcommondccallactivevadenable}
    cvcommondccallactiveentry.EntityData.Leafs["cvCommonDcCallActiveCoderTypeRate"] = types.YLeaf{"Cvcommondccallactivecodertyperate", cvcommondccallactiveentry.Cvcommondccallactivecodertyperate}
    cvcommondccallactiveentry.EntityData.Leafs["cvCommonDcCallActiveCodecBytes"] = types.YLeaf{"Cvcommondccallactivecodecbytes", cvcommondccallactiveentry.Cvcommondccallactivecodecbytes}
    cvcommondccallactiveentry.EntityData.Leafs["cvCommonDcCallActiveInBandSignaling"] = types.YLeaf{"Cvcommondccallactiveinbandsignaling", cvcommondccallactiveentry.Cvcommondccallactiveinbandsignaling}
    cvcommondccallactiveentry.EntityData.Leafs["cvCommonDcCallActiveCallingName"] = types.YLeaf{"Cvcommondccallactivecallingname", cvcommondccallactiveentry.Cvcommondccallactivecallingname}
    cvcommondccallactiveentry.EntityData.Leafs["cvCommonDcCallActiveCallerIDBlock"] = types.YLeaf{"Cvcommondccallactivecalleridblock", cvcommondccallactiveentry.Cvcommondccallactivecalleridblock}
    return &(cvcommondccallactiveentry.EntityData)
}

// CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallhistorytable
// This table is the Common extension to the call history
// table of IETF Dial Control MIB. It contains Common call 
// leg information about a network call leg.
type CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallhistorytable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The common information regarding a single network call leg. The call leg
    // entry is identified by using the same  index objects that are used by Call
    // Active table of IETF  Dial Control MIB to identify the call. An entry of
    // this table is created when its associated  call history entry in the IETF
    // Dial Control MIB is  created and the call history entry contains
    // information  for the call establishment to a network dialpeer. The entry is
    // deleted when its associated call history  entry in the IETF Dial Control
    // MIB is deleted. The type is slice of
    // CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallhistorytable_Cvcommondccallhistoryentry.
    Cvcommondccallhistoryentry []CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallhistorytable_Cvcommondccallhistoryentry
}

func (cvcommondccallhistorytable *CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallhistorytable) GetEntityData() *types.CommonEntityData {
    cvcommondccallhistorytable.EntityData.YFilter = cvcommondccallhistorytable.YFilter
    cvcommondccallhistorytable.EntityData.YangName = "cvCommonDcCallHistoryTable"
    cvcommondccallhistorytable.EntityData.BundleName = "cisco_ios_xe"
    cvcommondccallhistorytable.EntityData.ParentYangName = "CISCO-VOICE-COMMON-DIAL-CONTROL-MIB"
    cvcommondccallhistorytable.EntityData.SegmentPath = "cvCommonDcCallHistoryTable"
    cvcommondccallhistorytable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvcommondccallhistorytable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvcommondccallhistorytable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvcommondccallhistorytable.EntityData.Children = make(map[string]types.YChild)
    cvcommondccallhistorytable.EntityData.Children["cvCommonDcCallHistoryEntry"] = types.YChild{"Cvcommondccallhistoryentry", nil}
    for i := range cvcommondccallhistorytable.Cvcommondccallhistoryentry {
        cvcommondccallhistorytable.EntityData.Children[types.GetSegmentPath(&cvcommondccallhistorytable.Cvcommondccallhistoryentry[i])] = types.YChild{"Cvcommondccallhistoryentry", &cvcommondccallhistorytable.Cvcommondccallhistoryentry[i]}
    }
    cvcommondccallhistorytable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cvcommondccallhistorytable.EntityData)
}

// CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallhistorytable_Cvcommondccallhistoryentry
// The common information regarding a single network call
// leg. The call leg entry is identified by using the same 
// index objects that are used by Call Active table of IETF 
// Dial Control MIB to identify the call.
// An entry of this table is created when its associated 
// call history entry in the IETF Dial Control MIB is 
// created and the call history entry contains information 
// for the call establishment to a network dialpeer.
// The entry is deleted when its associated call history 
// entry in the IETF Dial Control MIB is deleted.
type CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallhistorytable_Cvcommondccallhistoryentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // cisco_dial_control_mib.CISCODIALCONTROLMIB_Ccallhistorytable_Ccallhistoryentry_Ccallhistoryindex
    Ccallhistoryindex interface{}

    // The global call identifier for the gateway call. The type is string with
    // length: 0..16.
    Cvcommondccallhistoryconnectionid interface{}

    // The object indicates whether or not the VAD (Voice Activity Detection) was
    // enabled for the voice call. The type is bool.
    Cvcommondccallhistoryvadenable interface{}

    // The negotiated coder rate. It specifies the transmit rate of voice/fax
    // compression to its associated call leg for  the call. This rate is
    // different from the configuration rate  because of rate negotiation during
    // the call. The type is CvcCoderTypeRate.
    Cvcommondccallhistorycodertyperate interface{}

    // Specifies the payload size of the voice packet. The type is interface{}
    // with range: 10..255.
    Cvcommondccallhistorycodecbytes interface{}

    // Specifies the type of in-band signaling used on the call.  This object is
    // instantiated only for  Connection Trunk calls. The type is
    // CvcInBandSignaling.
    Cvcommondccallhistoryinbandsignaling interface{}

    // The calling party name this call is connected to. If the name is not
    // available, then it will have a length  of zero. The type is string with
    // length: 0..64.
    Cvcommondccallhistorycallingname interface{}

    // The object indicates whether or not the caller ID feature is blocked for
    // this voice call. The type is bool.
    Cvcommondccallhistorycalleridblock interface{}
}

func (cvcommondccallhistoryentry *CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallhistorytable_Cvcommondccallhistoryentry) GetEntityData() *types.CommonEntityData {
    cvcommondccallhistoryentry.EntityData.YFilter = cvcommondccallhistoryentry.YFilter
    cvcommondccallhistoryentry.EntityData.YangName = "cvCommonDcCallHistoryEntry"
    cvcommondccallhistoryentry.EntityData.BundleName = "cisco_ios_xe"
    cvcommondccallhistoryentry.EntityData.ParentYangName = "cvCommonDcCallHistoryTable"
    cvcommondccallhistoryentry.EntityData.SegmentPath = "cvCommonDcCallHistoryEntry" + "[cCallHistoryIndex='" + fmt.Sprintf("%v", cvcommondccallhistoryentry.Ccallhistoryindex) + "']"
    cvcommondccallhistoryentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvcommondccallhistoryentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvcommondccallhistoryentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvcommondccallhistoryentry.EntityData.Children = make(map[string]types.YChild)
    cvcommondccallhistoryentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cvcommondccallhistoryentry.EntityData.Leafs["cCallHistoryIndex"] = types.YLeaf{"Ccallhistoryindex", cvcommondccallhistoryentry.Ccallhistoryindex}
    cvcommondccallhistoryentry.EntityData.Leafs["cvCommonDcCallHistoryConnectionId"] = types.YLeaf{"Cvcommondccallhistoryconnectionid", cvcommondccallhistoryentry.Cvcommondccallhistoryconnectionid}
    cvcommondccallhistoryentry.EntityData.Leafs["cvCommonDcCallHistoryVADEnable"] = types.YLeaf{"Cvcommondccallhistoryvadenable", cvcommondccallhistoryentry.Cvcommondccallhistoryvadenable}
    cvcommondccallhistoryentry.EntityData.Leafs["cvCommonDcCallHistoryCoderTypeRate"] = types.YLeaf{"Cvcommondccallhistorycodertyperate", cvcommondccallhistoryentry.Cvcommondccallhistorycodertyperate}
    cvcommondccallhistoryentry.EntityData.Leafs["cvCommonDcCallHistoryCodecBytes"] = types.YLeaf{"Cvcommondccallhistorycodecbytes", cvcommondccallhistoryentry.Cvcommondccallhistorycodecbytes}
    cvcommondccallhistoryentry.EntityData.Leafs["cvCommonDcCallHistoryInBandSignaling"] = types.YLeaf{"Cvcommondccallhistoryinbandsignaling", cvcommondccallhistoryentry.Cvcommondccallhistoryinbandsignaling}
    cvcommondccallhistoryentry.EntityData.Leafs["cvCommonDcCallHistoryCallingName"] = types.YLeaf{"Cvcommondccallhistorycallingname", cvcommondccallhistoryentry.Cvcommondccallhistorycallingname}
    cvcommondccallhistoryentry.EntityData.Leafs["cvCommonDcCallHistoryCallerIDBlock"] = types.YLeaf{"Cvcommondccallhistorycalleridblock", cvcommondccallhistoryentry.Cvcommondccallhistorycalleridblock}
    return &(cvcommondccallhistoryentry.EntityData)
}

