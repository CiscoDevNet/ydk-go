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

// CvcInBandSignaling represents               GR303
type CvcInBandSignaling string

const (
    CvcInBandSignaling_cas CvcInBandSignaling = "cas"

    CvcInBandSignaling_none CvcInBandSignaling = "none"

    CvcInBandSignaling_cept CvcInBandSignaling = "cept"

    CvcInBandSignaling_transparent CvcInBandSignaling = "transparent"

    CvcInBandSignaling_gr303 CvcInBandSignaling = "gr303"
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

// CvcVideoCoderRate represents the video data of the voice call.
type CvcVideoCoderRate string

const (
    CvcVideoCoderRate_none CvcVideoCoderRate = "none"

    CvcVideoCoderRate_h261 CvcVideoCoderRate = "h261"

    CvcVideoCoderRate_h263 CvcVideoCoderRate = "h263"

    CvcVideoCoderRate_h263plus CvcVideoCoderRate = "h263plus"

    CvcVideoCoderRate_h264 CvcVideoCoderRate = "h264"
)

// CvcH320CallType represents This object specifies the H320 call type of a voice call.
type CvcH320CallType string

const (
    CvcH320CallType_none CvcH320CallType = "none"

    CvcH320CallType_primary CvcH320CallType = "primary"

    CvcH320CallType_secondary CvcH320CallType = "secondary"
)

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

// CISCOVOICECOMMONDIALCONTROLMIB
type CISCOVOICECOMMONDIALCONTROLMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This table is a common extension to the call active table of IETF Dial
    // Control MIB. It contains common call  leg information about a network call
    // leg.
    CvCommonDcCallActiveTable CISCOVOICECOMMONDIALCONTROLMIB_CvCommonDcCallActiveTable

    // This table is the Common extension to the call history table of IETF Dial
    // Control MIB. It contains Common call  leg information about a network call
    // leg.
    CvCommonDcCallHistoryTable CISCOVOICECOMMONDIALCONTROLMIB_CvCommonDcCallHistoryTable
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

    cISCOVOICECOMMONDIALCONTROLMIB.EntityData.Children = types.NewOrderedMap()
    cISCOVOICECOMMONDIALCONTROLMIB.EntityData.Children.Append("cvCommonDcCallActiveTable", types.YChild{"CvCommonDcCallActiveTable", &cISCOVOICECOMMONDIALCONTROLMIB.CvCommonDcCallActiveTable})
    cISCOVOICECOMMONDIALCONTROLMIB.EntityData.Children.Append("cvCommonDcCallHistoryTable", types.YChild{"CvCommonDcCallHistoryTable", &cISCOVOICECOMMONDIALCONTROLMIB.CvCommonDcCallHistoryTable})
    cISCOVOICECOMMONDIALCONTROLMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOVOICECOMMONDIALCONTROLMIB.EntityData.YListKeys = []string {}

    return &(cISCOVOICECOMMONDIALCONTROLMIB.EntityData)
}

// CISCOVOICECOMMONDIALCONTROLMIB_CvCommonDcCallActiveTable
// This table is a common extension to the call active
// table of IETF Dial Control MIB. It contains common call 
// leg information about a network call leg.
type CISCOVOICECOMMONDIALCONTROLMIB_CvCommonDcCallActiveTable struct {
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
    // CISCOVOICECOMMONDIALCONTROLMIB_CvCommonDcCallActiveTable_CvCommonDcCallActiveEntry.
    CvCommonDcCallActiveEntry []*CISCOVOICECOMMONDIALCONTROLMIB_CvCommonDcCallActiveTable_CvCommonDcCallActiveEntry
}

func (cvCommonDcCallActiveTable *CISCOVOICECOMMONDIALCONTROLMIB_CvCommonDcCallActiveTable) GetEntityData() *types.CommonEntityData {
    cvCommonDcCallActiveTable.EntityData.YFilter = cvCommonDcCallActiveTable.YFilter
    cvCommonDcCallActiveTable.EntityData.YangName = "cvCommonDcCallActiveTable"
    cvCommonDcCallActiveTable.EntityData.BundleName = "cisco_ios_xe"
    cvCommonDcCallActiveTable.EntityData.ParentYangName = "CISCO-VOICE-COMMON-DIAL-CONTROL-MIB"
    cvCommonDcCallActiveTable.EntityData.SegmentPath = "cvCommonDcCallActiveTable"
    cvCommonDcCallActiveTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvCommonDcCallActiveTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvCommonDcCallActiveTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvCommonDcCallActiveTable.EntityData.Children = types.NewOrderedMap()
    cvCommonDcCallActiveTable.EntityData.Children.Append("cvCommonDcCallActiveEntry", types.YChild{"CvCommonDcCallActiveEntry", nil})
    for i := range cvCommonDcCallActiveTable.CvCommonDcCallActiveEntry {
        cvCommonDcCallActiveTable.EntityData.Children.Append(types.GetSegmentPath(cvCommonDcCallActiveTable.CvCommonDcCallActiveEntry[i]), types.YChild{"CvCommonDcCallActiveEntry", cvCommonDcCallActiveTable.CvCommonDcCallActiveEntry[i]})
    }
    cvCommonDcCallActiveTable.EntityData.Leafs = types.NewOrderedMap()

    cvCommonDcCallActiveTable.EntityData.YListKeys = []string {}

    return &(cvCommonDcCallActiveTable.EntityData)
}

// CISCOVOICECOMMONDIALCONTROLMIB_CvCommonDcCallActiveTable_CvCommonDcCallActiveEntry
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
type CISCOVOICECOMMONDIALCONTROLMIB_CvCommonDcCallActiveTable_CvCommonDcCallActiveEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // dial_control_mib.DIALCONTROLMIB_CallActiveTable_CallActiveEntry_CallActiveSetupTime
    CallActiveSetupTime interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // dial_control_mib.DIALCONTROLMIB_CallActiveTable_CallActiveEntry_CallActiveIndex
    CallActiveIndex interface{}

    // The global call identifier for the network call. The type is string with
    // length: 0..16.
    CvCommonDcCallActiveConnectionId interface{}

    // The object indicates whether or not the VAD (Voice Activity Detection) is
    // enabled for the voice call. The type is bool.
    CvCommonDcCallActiveVADEnable interface{}

    // The negotiated coder rate. It specifies the transmit rate of voice/fax
    // compression to its associated call leg  for the call. This rate is
    // different from the configuration rate  because of rate negotiation during
    // the call. The type is CvcCoderTypeRate.
    CvCommonDcCallActiveCoderTypeRate interface{}

    // Specifies the payload size of the voice packet. The type is interface{}
    // with range: 10..255.
    CvCommonDcCallActiveCodecBytes interface{}

    // Specifies the type of in-band signaling being used on the call.  This
    // object is instantiated only for  Connection Trunk calls. The type is
    // CvcInBandSignaling.
    CvCommonDcCallActiveInBandSignaling interface{}

    // The calling party name this call is connected to. If the name is not
    // available, then it will have a length of zero. The type is string with
    // length: 0..64.
    CvCommonDcCallActiveCallingName interface{}

    // The object indicates whether or not the caller ID feature is blocked for
    // this voice call. The type is bool.
    CvCommonDcCallActiveCallerIDBlock interface{}
}

func (cvCommonDcCallActiveEntry *CISCOVOICECOMMONDIALCONTROLMIB_CvCommonDcCallActiveTable_CvCommonDcCallActiveEntry) GetEntityData() *types.CommonEntityData {
    cvCommonDcCallActiveEntry.EntityData.YFilter = cvCommonDcCallActiveEntry.YFilter
    cvCommonDcCallActiveEntry.EntityData.YangName = "cvCommonDcCallActiveEntry"
    cvCommonDcCallActiveEntry.EntityData.BundleName = "cisco_ios_xe"
    cvCommonDcCallActiveEntry.EntityData.ParentYangName = "cvCommonDcCallActiveTable"
    cvCommonDcCallActiveEntry.EntityData.SegmentPath = "cvCommonDcCallActiveEntry" + types.AddKeyToken(cvCommonDcCallActiveEntry.CallActiveSetupTime, "callActiveSetupTime") + types.AddKeyToken(cvCommonDcCallActiveEntry.CallActiveIndex, "callActiveIndex")
    cvCommonDcCallActiveEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvCommonDcCallActiveEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvCommonDcCallActiveEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvCommonDcCallActiveEntry.EntityData.Children = types.NewOrderedMap()
    cvCommonDcCallActiveEntry.EntityData.Leafs = types.NewOrderedMap()
    cvCommonDcCallActiveEntry.EntityData.Leafs.Append("callActiveSetupTime", types.YLeaf{"CallActiveSetupTime", cvCommonDcCallActiveEntry.CallActiveSetupTime})
    cvCommonDcCallActiveEntry.EntityData.Leafs.Append("callActiveIndex", types.YLeaf{"CallActiveIndex", cvCommonDcCallActiveEntry.CallActiveIndex})
    cvCommonDcCallActiveEntry.EntityData.Leafs.Append("cvCommonDcCallActiveConnectionId", types.YLeaf{"CvCommonDcCallActiveConnectionId", cvCommonDcCallActiveEntry.CvCommonDcCallActiveConnectionId})
    cvCommonDcCallActiveEntry.EntityData.Leafs.Append("cvCommonDcCallActiveVADEnable", types.YLeaf{"CvCommonDcCallActiveVADEnable", cvCommonDcCallActiveEntry.CvCommonDcCallActiveVADEnable})
    cvCommonDcCallActiveEntry.EntityData.Leafs.Append("cvCommonDcCallActiveCoderTypeRate", types.YLeaf{"CvCommonDcCallActiveCoderTypeRate", cvCommonDcCallActiveEntry.CvCommonDcCallActiveCoderTypeRate})
    cvCommonDcCallActiveEntry.EntityData.Leafs.Append("cvCommonDcCallActiveCodecBytes", types.YLeaf{"CvCommonDcCallActiveCodecBytes", cvCommonDcCallActiveEntry.CvCommonDcCallActiveCodecBytes})
    cvCommonDcCallActiveEntry.EntityData.Leafs.Append("cvCommonDcCallActiveInBandSignaling", types.YLeaf{"CvCommonDcCallActiveInBandSignaling", cvCommonDcCallActiveEntry.CvCommonDcCallActiveInBandSignaling})
    cvCommonDcCallActiveEntry.EntityData.Leafs.Append("cvCommonDcCallActiveCallingName", types.YLeaf{"CvCommonDcCallActiveCallingName", cvCommonDcCallActiveEntry.CvCommonDcCallActiveCallingName})
    cvCommonDcCallActiveEntry.EntityData.Leafs.Append("cvCommonDcCallActiveCallerIDBlock", types.YLeaf{"CvCommonDcCallActiveCallerIDBlock", cvCommonDcCallActiveEntry.CvCommonDcCallActiveCallerIDBlock})

    cvCommonDcCallActiveEntry.EntityData.YListKeys = []string {"CallActiveSetupTime", "CallActiveIndex"}

    return &(cvCommonDcCallActiveEntry.EntityData)
}

// CISCOVOICECOMMONDIALCONTROLMIB_CvCommonDcCallHistoryTable
// This table is the Common extension to the call history
// table of IETF Dial Control MIB. It contains Common call 
// leg information about a network call leg.
type CISCOVOICECOMMONDIALCONTROLMIB_CvCommonDcCallHistoryTable struct {
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
    // CISCOVOICECOMMONDIALCONTROLMIB_CvCommonDcCallHistoryTable_CvCommonDcCallHistoryEntry.
    CvCommonDcCallHistoryEntry []*CISCOVOICECOMMONDIALCONTROLMIB_CvCommonDcCallHistoryTable_CvCommonDcCallHistoryEntry
}

func (cvCommonDcCallHistoryTable *CISCOVOICECOMMONDIALCONTROLMIB_CvCommonDcCallHistoryTable) GetEntityData() *types.CommonEntityData {
    cvCommonDcCallHistoryTable.EntityData.YFilter = cvCommonDcCallHistoryTable.YFilter
    cvCommonDcCallHistoryTable.EntityData.YangName = "cvCommonDcCallHistoryTable"
    cvCommonDcCallHistoryTable.EntityData.BundleName = "cisco_ios_xe"
    cvCommonDcCallHistoryTable.EntityData.ParentYangName = "CISCO-VOICE-COMMON-DIAL-CONTROL-MIB"
    cvCommonDcCallHistoryTable.EntityData.SegmentPath = "cvCommonDcCallHistoryTable"
    cvCommonDcCallHistoryTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvCommonDcCallHistoryTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvCommonDcCallHistoryTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvCommonDcCallHistoryTable.EntityData.Children = types.NewOrderedMap()
    cvCommonDcCallHistoryTable.EntityData.Children.Append("cvCommonDcCallHistoryEntry", types.YChild{"CvCommonDcCallHistoryEntry", nil})
    for i := range cvCommonDcCallHistoryTable.CvCommonDcCallHistoryEntry {
        cvCommonDcCallHistoryTable.EntityData.Children.Append(types.GetSegmentPath(cvCommonDcCallHistoryTable.CvCommonDcCallHistoryEntry[i]), types.YChild{"CvCommonDcCallHistoryEntry", cvCommonDcCallHistoryTable.CvCommonDcCallHistoryEntry[i]})
    }
    cvCommonDcCallHistoryTable.EntityData.Leafs = types.NewOrderedMap()

    cvCommonDcCallHistoryTable.EntityData.YListKeys = []string {}

    return &(cvCommonDcCallHistoryTable.EntityData)
}

// CISCOVOICECOMMONDIALCONTROLMIB_CvCommonDcCallHistoryTable_CvCommonDcCallHistoryEntry
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
type CISCOVOICECOMMONDIALCONTROLMIB_CvCommonDcCallHistoryTable_CvCommonDcCallHistoryEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // cisco_dial_control_mib.CISCODIALCONTROLMIB_CCallHistoryTable_CCallHistoryEntry_CCallHistoryIndex
    CCallHistoryIndex interface{}

    // The global call identifier for the gateway call. The type is string with
    // length: 0..16.
    CvCommonDcCallHistoryConnectionId interface{}

    // The object indicates whether or not the VAD (Voice Activity Detection) was
    // enabled for the voice call. The type is bool.
    CvCommonDcCallHistoryVADEnable interface{}

    // The negotiated coder rate. It specifies the transmit rate of voice/fax
    // compression to its associated call leg for  the call. This rate is
    // different from the configuration rate  because of rate negotiation during
    // the call. The type is CvcCoderTypeRate.
    CvCommonDcCallHistoryCoderTypeRate interface{}

    // Specifies the payload size of the voice packet. The type is interface{}
    // with range: 10..255.
    CvCommonDcCallHistoryCodecBytes interface{}

    // Specifies the type of in-band signaling used on the call.  This object is
    // instantiated only for  Connection Trunk calls. The type is
    // CvcInBandSignaling.
    CvCommonDcCallHistoryInBandSignaling interface{}

    // The calling party name this call is connected to. If the name is not
    // available, then it will have a length  of zero. The type is string with
    // length: 0..64.
    CvCommonDcCallHistoryCallingName interface{}

    // The object indicates whether or not the caller ID feature is blocked for
    // this voice call. The type is bool.
    CvCommonDcCallHistoryCallerIDBlock interface{}
}

func (cvCommonDcCallHistoryEntry *CISCOVOICECOMMONDIALCONTROLMIB_CvCommonDcCallHistoryTable_CvCommonDcCallHistoryEntry) GetEntityData() *types.CommonEntityData {
    cvCommonDcCallHistoryEntry.EntityData.YFilter = cvCommonDcCallHistoryEntry.YFilter
    cvCommonDcCallHistoryEntry.EntityData.YangName = "cvCommonDcCallHistoryEntry"
    cvCommonDcCallHistoryEntry.EntityData.BundleName = "cisco_ios_xe"
    cvCommonDcCallHistoryEntry.EntityData.ParentYangName = "cvCommonDcCallHistoryTable"
    cvCommonDcCallHistoryEntry.EntityData.SegmentPath = "cvCommonDcCallHistoryEntry" + types.AddKeyToken(cvCommonDcCallHistoryEntry.CCallHistoryIndex, "cCallHistoryIndex")
    cvCommonDcCallHistoryEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvCommonDcCallHistoryEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvCommonDcCallHistoryEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvCommonDcCallHistoryEntry.EntityData.Children = types.NewOrderedMap()
    cvCommonDcCallHistoryEntry.EntityData.Leafs = types.NewOrderedMap()
    cvCommonDcCallHistoryEntry.EntityData.Leafs.Append("cCallHistoryIndex", types.YLeaf{"CCallHistoryIndex", cvCommonDcCallHistoryEntry.CCallHistoryIndex})
    cvCommonDcCallHistoryEntry.EntityData.Leafs.Append("cvCommonDcCallHistoryConnectionId", types.YLeaf{"CvCommonDcCallHistoryConnectionId", cvCommonDcCallHistoryEntry.CvCommonDcCallHistoryConnectionId})
    cvCommonDcCallHistoryEntry.EntityData.Leafs.Append("cvCommonDcCallHistoryVADEnable", types.YLeaf{"CvCommonDcCallHistoryVADEnable", cvCommonDcCallHistoryEntry.CvCommonDcCallHistoryVADEnable})
    cvCommonDcCallHistoryEntry.EntityData.Leafs.Append("cvCommonDcCallHistoryCoderTypeRate", types.YLeaf{"CvCommonDcCallHistoryCoderTypeRate", cvCommonDcCallHistoryEntry.CvCommonDcCallHistoryCoderTypeRate})
    cvCommonDcCallHistoryEntry.EntityData.Leafs.Append("cvCommonDcCallHistoryCodecBytes", types.YLeaf{"CvCommonDcCallHistoryCodecBytes", cvCommonDcCallHistoryEntry.CvCommonDcCallHistoryCodecBytes})
    cvCommonDcCallHistoryEntry.EntityData.Leafs.Append("cvCommonDcCallHistoryInBandSignaling", types.YLeaf{"CvCommonDcCallHistoryInBandSignaling", cvCommonDcCallHistoryEntry.CvCommonDcCallHistoryInBandSignaling})
    cvCommonDcCallHistoryEntry.EntityData.Leafs.Append("cvCommonDcCallHistoryCallingName", types.YLeaf{"CvCommonDcCallHistoryCallingName", cvCommonDcCallHistoryEntry.CvCommonDcCallHistoryCallingName})
    cvCommonDcCallHistoryEntry.EntityData.Leafs.Append("cvCommonDcCallHistoryCallerIDBlock", types.YLeaf{"CvCommonDcCallHistoryCallerIDBlock", cvCommonDcCallHistoryEntry.CvCommonDcCallHistoryCallerIDBlock})

    cvCommonDcCallHistoryEntry.EntityData.YListKeys = []string {"CCallHistoryIndex"}

    return &(cvCommonDcCallHistoryEntry.EntityData)
}

