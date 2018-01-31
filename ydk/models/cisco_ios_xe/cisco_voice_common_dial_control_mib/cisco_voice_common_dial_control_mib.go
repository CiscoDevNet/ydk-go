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
    parent types.Entity
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

func (cISCOVOICECOMMONDIALCONTROLMIB *CISCOVOICECOMMONDIALCONTROLMIB) GetFilter() yfilter.YFilter { return cISCOVOICECOMMONDIALCONTROLMIB.YFilter }

func (cISCOVOICECOMMONDIALCONTROLMIB *CISCOVOICECOMMONDIALCONTROLMIB) SetFilter(yf yfilter.YFilter) { cISCOVOICECOMMONDIALCONTROLMIB.YFilter = yf }

func (cISCOVOICECOMMONDIALCONTROLMIB *CISCOVOICECOMMONDIALCONTROLMIB) GetGoName(yname string) string {
    if yname == "cvCommonDcCallActiveTable" { return "Cvcommondccallactivetable" }
    if yname == "cvCommonDcCallHistoryTable" { return "Cvcommondccallhistorytable" }
    return ""
}

func (cISCOVOICECOMMONDIALCONTROLMIB *CISCOVOICECOMMONDIALCONTROLMIB) GetSegmentPath() string {
    return "CISCO-VOICE-COMMON-DIAL-CONTROL-MIB:CISCO-VOICE-COMMON-DIAL-CONTROL-MIB"
}

func (cISCOVOICECOMMONDIALCONTROLMIB *CISCOVOICECOMMONDIALCONTROLMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cvCommonDcCallActiveTable" {
        return &cISCOVOICECOMMONDIALCONTROLMIB.Cvcommondccallactivetable
    }
    if childYangName == "cvCommonDcCallHistoryTable" {
        return &cISCOVOICECOMMONDIALCONTROLMIB.Cvcommondccallhistorytable
    }
    return nil
}

func (cISCOVOICECOMMONDIALCONTROLMIB *CISCOVOICECOMMONDIALCONTROLMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cvCommonDcCallActiveTable"] = &cISCOVOICECOMMONDIALCONTROLMIB.Cvcommondccallactivetable
    children["cvCommonDcCallHistoryTable"] = &cISCOVOICECOMMONDIALCONTROLMIB.Cvcommondccallhistorytable
    return children
}

func (cISCOVOICECOMMONDIALCONTROLMIB *CISCOVOICECOMMONDIALCONTROLMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOVOICECOMMONDIALCONTROLMIB *CISCOVOICECOMMONDIALCONTROLMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOVOICECOMMONDIALCONTROLMIB *CISCOVOICECOMMONDIALCONTROLMIB) GetYangName() string { return "CISCO-VOICE-COMMON-DIAL-CONTROL-MIB" }

func (cISCOVOICECOMMONDIALCONTROLMIB *CISCOVOICECOMMONDIALCONTROLMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOVOICECOMMONDIALCONTROLMIB *CISCOVOICECOMMONDIALCONTROLMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOVOICECOMMONDIALCONTROLMIB *CISCOVOICECOMMONDIALCONTROLMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOVOICECOMMONDIALCONTROLMIB *CISCOVOICECOMMONDIALCONTROLMIB) SetParent(parent types.Entity) { cISCOVOICECOMMONDIALCONTROLMIB.parent = parent }

func (cISCOVOICECOMMONDIALCONTROLMIB *CISCOVOICECOMMONDIALCONTROLMIB) GetParent() types.Entity { return cISCOVOICECOMMONDIALCONTROLMIB.parent }

func (cISCOVOICECOMMONDIALCONTROLMIB *CISCOVOICECOMMONDIALCONTROLMIB) GetParentYangName() string { return "CISCO-VOICE-COMMON-DIAL-CONTROL-MIB" }

// CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallactivetable
// This table is a common extension to the call active
// table of IETF Dial Control MIB. It contains common call 
// leg information about a network call leg.
type CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallactivetable struct {
    parent types.Entity
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

func (cvcommondccallactivetable *CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallactivetable) GetFilter() yfilter.YFilter { return cvcommondccallactivetable.YFilter }

func (cvcommondccallactivetable *CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallactivetable) SetFilter(yf yfilter.YFilter) { cvcommondccallactivetable.YFilter = yf }

func (cvcommondccallactivetable *CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallactivetable) GetGoName(yname string) string {
    if yname == "cvCommonDcCallActiveEntry" { return "Cvcommondccallactiveentry" }
    return ""
}

func (cvcommondccallactivetable *CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallactivetable) GetSegmentPath() string {
    return "cvCommonDcCallActiveTable"
}

func (cvcommondccallactivetable *CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallactivetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cvCommonDcCallActiveEntry" {
        for _, c := range cvcommondccallactivetable.Cvcommondccallactiveentry {
            if cvcommondccallactivetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallactivetable_Cvcommondccallactiveentry{}
        cvcommondccallactivetable.Cvcommondccallactiveentry = append(cvcommondccallactivetable.Cvcommondccallactiveentry, child)
        return &cvcommondccallactivetable.Cvcommondccallactiveentry[len(cvcommondccallactivetable.Cvcommondccallactiveentry)-1]
    }
    return nil
}

func (cvcommondccallactivetable *CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallactivetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cvcommondccallactivetable.Cvcommondccallactiveentry {
        children[cvcommondccallactivetable.Cvcommondccallactiveentry[i].GetSegmentPath()] = &cvcommondccallactivetable.Cvcommondccallactiveentry[i]
    }
    return children
}

func (cvcommondccallactivetable *CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallactivetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cvcommondccallactivetable *CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallactivetable) GetBundleName() string { return "cisco_ios_xe" }

func (cvcommondccallactivetable *CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallactivetable) GetYangName() string { return "cvCommonDcCallActiveTable" }

func (cvcommondccallactivetable *CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallactivetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvcommondccallactivetable *CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallactivetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvcommondccallactivetable *CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallactivetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvcommondccallactivetable *CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallactivetable) SetParent(parent types.Entity) { cvcommondccallactivetable.parent = parent }

func (cvcommondccallactivetable *CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallactivetable) GetParent() types.Entity { return cvcommondccallactivetable.parent }

func (cvcommondccallactivetable *CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallactivetable) GetParentYangName() string { return "CISCO-VOICE-COMMON-DIAL-CONTROL-MIB" }

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

func (cvcommondccallactiveentry *CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallactivetable_Cvcommondccallactiveentry) GetFilter() yfilter.YFilter { return cvcommondccallactiveentry.YFilter }

func (cvcommondccallactiveentry *CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallactivetable_Cvcommondccallactiveentry) SetFilter(yf yfilter.YFilter) { cvcommondccallactiveentry.YFilter = yf }

func (cvcommondccallactiveentry *CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallactivetable_Cvcommondccallactiveentry) GetGoName(yname string) string {
    if yname == "callActiveSetupTime" { return "Callactivesetuptime" }
    if yname == "callActiveIndex" { return "Callactiveindex" }
    if yname == "cvCommonDcCallActiveConnectionId" { return "Cvcommondccallactiveconnectionid" }
    if yname == "cvCommonDcCallActiveVADEnable" { return "Cvcommondccallactivevadenable" }
    if yname == "cvCommonDcCallActiveCoderTypeRate" { return "Cvcommondccallactivecodertyperate" }
    if yname == "cvCommonDcCallActiveCodecBytes" { return "Cvcommondccallactivecodecbytes" }
    if yname == "cvCommonDcCallActiveInBandSignaling" { return "Cvcommondccallactiveinbandsignaling" }
    if yname == "cvCommonDcCallActiveCallingName" { return "Cvcommondccallactivecallingname" }
    if yname == "cvCommonDcCallActiveCallerIDBlock" { return "Cvcommondccallactivecalleridblock" }
    return ""
}

func (cvcommondccallactiveentry *CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallactivetable_Cvcommondccallactiveentry) GetSegmentPath() string {
    return "cvCommonDcCallActiveEntry" + "[callActiveSetupTime='" + fmt.Sprintf("%v", cvcommondccallactiveentry.Callactivesetuptime) + "']" + "[callActiveIndex='" + fmt.Sprintf("%v", cvcommondccallactiveentry.Callactiveindex) + "']"
}

func (cvcommondccallactiveentry *CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallactivetable_Cvcommondccallactiveentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cvcommondccallactiveentry *CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallactivetable_Cvcommondccallactiveentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cvcommondccallactiveentry *CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallactivetable_Cvcommondccallactiveentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["callActiveSetupTime"] = cvcommondccallactiveentry.Callactivesetuptime
    leafs["callActiveIndex"] = cvcommondccallactiveentry.Callactiveindex
    leafs["cvCommonDcCallActiveConnectionId"] = cvcommondccallactiveentry.Cvcommondccallactiveconnectionid
    leafs["cvCommonDcCallActiveVADEnable"] = cvcommondccallactiveentry.Cvcommondccallactivevadenable
    leafs["cvCommonDcCallActiveCoderTypeRate"] = cvcommondccallactiveentry.Cvcommondccallactivecodertyperate
    leafs["cvCommonDcCallActiveCodecBytes"] = cvcommondccallactiveentry.Cvcommondccallactivecodecbytes
    leafs["cvCommonDcCallActiveInBandSignaling"] = cvcommondccallactiveentry.Cvcommondccallactiveinbandsignaling
    leafs["cvCommonDcCallActiveCallingName"] = cvcommondccallactiveentry.Cvcommondccallactivecallingname
    leafs["cvCommonDcCallActiveCallerIDBlock"] = cvcommondccallactiveentry.Cvcommondccallactivecalleridblock
    return leafs
}

func (cvcommondccallactiveentry *CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallactivetable_Cvcommondccallactiveentry) GetBundleName() string { return "cisco_ios_xe" }

func (cvcommondccallactiveentry *CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallactivetable_Cvcommondccallactiveentry) GetYangName() string { return "cvCommonDcCallActiveEntry" }

func (cvcommondccallactiveentry *CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallactivetable_Cvcommondccallactiveentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvcommondccallactiveentry *CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallactivetable_Cvcommondccallactiveentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvcommondccallactiveentry *CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallactivetable_Cvcommondccallactiveentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvcommondccallactiveentry *CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallactivetable_Cvcommondccallactiveentry) SetParent(parent types.Entity) { cvcommondccallactiveentry.parent = parent }

func (cvcommondccallactiveentry *CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallactivetable_Cvcommondccallactiveentry) GetParent() types.Entity { return cvcommondccallactiveentry.parent }

func (cvcommondccallactiveentry *CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallactivetable_Cvcommondccallactiveentry) GetParentYangName() string { return "cvCommonDcCallActiveTable" }

// CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallhistorytable
// This table is the Common extension to the call history
// table of IETF Dial Control MIB. It contains Common call 
// leg information about a network call leg.
type CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallhistorytable struct {
    parent types.Entity
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

func (cvcommondccallhistorytable *CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallhistorytable) GetFilter() yfilter.YFilter { return cvcommondccallhistorytable.YFilter }

func (cvcommondccallhistorytable *CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallhistorytable) SetFilter(yf yfilter.YFilter) { cvcommondccallhistorytable.YFilter = yf }

func (cvcommondccallhistorytable *CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallhistorytable) GetGoName(yname string) string {
    if yname == "cvCommonDcCallHistoryEntry" { return "Cvcommondccallhistoryentry" }
    return ""
}

func (cvcommondccallhistorytable *CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallhistorytable) GetSegmentPath() string {
    return "cvCommonDcCallHistoryTable"
}

func (cvcommondccallhistorytable *CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallhistorytable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cvCommonDcCallHistoryEntry" {
        for _, c := range cvcommondccallhistorytable.Cvcommondccallhistoryentry {
            if cvcommondccallhistorytable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallhistorytable_Cvcommondccallhistoryentry{}
        cvcommondccallhistorytable.Cvcommondccallhistoryentry = append(cvcommondccallhistorytable.Cvcommondccallhistoryentry, child)
        return &cvcommondccallhistorytable.Cvcommondccallhistoryentry[len(cvcommondccallhistorytable.Cvcommondccallhistoryentry)-1]
    }
    return nil
}

func (cvcommondccallhistorytable *CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallhistorytable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cvcommondccallhistorytable.Cvcommondccallhistoryentry {
        children[cvcommondccallhistorytable.Cvcommondccallhistoryentry[i].GetSegmentPath()] = &cvcommondccallhistorytable.Cvcommondccallhistoryentry[i]
    }
    return children
}

func (cvcommondccallhistorytable *CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallhistorytable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cvcommondccallhistorytable *CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallhistorytable) GetBundleName() string { return "cisco_ios_xe" }

func (cvcommondccallhistorytable *CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallhistorytable) GetYangName() string { return "cvCommonDcCallHistoryTable" }

func (cvcommondccallhistorytable *CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallhistorytable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvcommondccallhistorytable *CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallhistorytable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvcommondccallhistorytable *CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallhistorytable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvcommondccallhistorytable *CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallhistorytable) SetParent(parent types.Entity) { cvcommondccallhistorytable.parent = parent }

func (cvcommondccallhistorytable *CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallhistorytable) GetParent() types.Entity { return cvcommondccallhistorytable.parent }

func (cvcommondccallhistorytable *CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallhistorytable) GetParentYangName() string { return "CISCO-VOICE-COMMON-DIAL-CONTROL-MIB" }

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
    parent types.Entity
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

func (cvcommondccallhistoryentry *CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallhistorytable_Cvcommondccallhistoryentry) GetFilter() yfilter.YFilter { return cvcommondccallhistoryentry.YFilter }

func (cvcommondccallhistoryentry *CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallhistorytable_Cvcommondccallhistoryentry) SetFilter(yf yfilter.YFilter) { cvcommondccallhistoryentry.YFilter = yf }

func (cvcommondccallhistoryentry *CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallhistorytable_Cvcommondccallhistoryentry) GetGoName(yname string) string {
    if yname == "cCallHistoryIndex" { return "Ccallhistoryindex" }
    if yname == "cvCommonDcCallHistoryConnectionId" { return "Cvcommondccallhistoryconnectionid" }
    if yname == "cvCommonDcCallHistoryVADEnable" { return "Cvcommondccallhistoryvadenable" }
    if yname == "cvCommonDcCallHistoryCoderTypeRate" { return "Cvcommondccallhistorycodertyperate" }
    if yname == "cvCommonDcCallHistoryCodecBytes" { return "Cvcommondccallhistorycodecbytes" }
    if yname == "cvCommonDcCallHistoryInBandSignaling" { return "Cvcommondccallhistoryinbandsignaling" }
    if yname == "cvCommonDcCallHistoryCallingName" { return "Cvcommondccallhistorycallingname" }
    if yname == "cvCommonDcCallHistoryCallerIDBlock" { return "Cvcommondccallhistorycalleridblock" }
    return ""
}

func (cvcommondccallhistoryentry *CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallhistorytable_Cvcommondccallhistoryentry) GetSegmentPath() string {
    return "cvCommonDcCallHistoryEntry" + "[cCallHistoryIndex='" + fmt.Sprintf("%v", cvcommondccallhistoryentry.Ccallhistoryindex) + "']"
}

func (cvcommondccallhistoryentry *CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallhistorytable_Cvcommondccallhistoryentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cvcommondccallhistoryentry *CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallhistorytable_Cvcommondccallhistoryentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cvcommondccallhistoryentry *CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallhistorytable_Cvcommondccallhistoryentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cCallHistoryIndex"] = cvcommondccallhistoryentry.Ccallhistoryindex
    leafs["cvCommonDcCallHistoryConnectionId"] = cvcommondccallhistoryentry.Cvcommondccallhistoryconnectionid
    leafs["cvCommonDcCallHistoryVADEnable"] = cvcommondccallhistoryentry.Cvcommondccallhistoryvadenable
    leafs["cvCommonDcCallHistoryCoderTypeRate"] = cvcommondccallhistoryentry.Cvcommondccallhistorycodertyperate
    leafs["cvCommonDcCallHistoryCodecBytes"] = cvcommondccallhistoryentry.Cvcommondccallhistorycodecbytes
    leafs["cvCommonDcCallHistoryInBandSignaling"] = cvcommondccallhistoryentry.Cvcommondccallhistoryinbandsignaling
    leafs["cvCommonDcCallHistoryCallingName"] = cvcommondccallhistoryentry.Cvcommondccallhistorycallingname
    leafs["cvCommonDcCallHistoryCallerIDBlock"] = cvcommondccallhistoryentry.Cvcommondccallhistorycalleridblock
    return leafs
}

func (cvcommondccallhistoryentry *CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallhistorytable_Cvcommondccallhistoryentry) GetBundleName() string { return "cisco_ios_xe" }

func (cvcommondccallhistoryentry *CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallhistorytable_Cvcommondccallhistoryentry) GetYangName() string { return "cvCommonDcCallHistoryEntry" }

func (cvcommondccallhistoryentry *CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallhistorytable_Cvcommondccallhistoryentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvcommondccallhistoryentry *CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallhistorytable_Cvcommondccallhistoryentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvcommondccallhistoryentry *CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallhistorytable_Cvcommondccallhistoryentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvcommondccallhistoryentry *CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallhistorytable_Cvcommondccallhistoryentry) SetParent(parent types.Entity) { cvcommondccallhistoryentry.parent = parent }

func (cvcommondccallhistoryentry *CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallhistorytable_Cvcommondccallhistoryentry) GetParent() types.Entity { return cvcommondccallhistoryentry.parent }

func (cvcommondccallhistoryentry *CISCOVOICECOMMONDIALCONTROLMIB_Cvcommondccallhistorytable_Cvcommondccallhistoryentry) GetParentYangName() string { return "cvCommonDcCallHistoryTable" }

