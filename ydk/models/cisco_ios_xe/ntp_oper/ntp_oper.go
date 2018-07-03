// This module contains a collection of YANG definitions
// for NTP operational data.
// Copyright (c) 2017 by Cisco Systems, Inc.
// All rights reserved.
package ntp_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ntp_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-ntp-oper ntp-oper-data}", reflect.TypeOf(NtpOperData{}))
    ydk.RegisterEntity("Cisco-IOS-XE-ntp-oper:ntp-oper-data", reflect.TypeOf(NtpOperData{}))
}

// KissCodeType represents Kiss code is used for debug or maintenance purposes in devices in stratum 0 or 16
type KissCodeType string

const (
    // The association belongs to a unicast server
    KissCodeType_ntp_ref_acst KissCodeType = "ntp-ref-acst"

    // Server authentication failed
    KissCodeType_ntp_ref_auth KissCodeType = "ntp-ref-auth"

    // Autokey sequence failed
    KissCodeType_ntp_ref_auto KissCodeType = "ntp-ref-auto"

    // The association belongs to a broadcast server
    KissCodeType_ntp_ref_bcst KissCodeType = "ntp-ref-bcst"

    // Cryptographic authentication or identification failed
    KissCodeType_ntp_ref_cryp KissCodeType = "ntp-ref-cryp"

    // Access denied by remote server
    KissCodeType_ntp_ref_deny KissCodeType = "ntp-ref-deny"

    // Lost peer in symmetric mode
    KissCodeType_ntp_ref_drop KissCodeType = "ntp-ref-drop"

    // Access denied due to local policy
    KissCodeType_ntp_ref_rstr KissCodeType = "ntp-ref-rstr"

    // The association has not synchronized for the first time
    KissCodeType_ntp_ref_init KissCodeType = "ntp-ref-init"

    // The association belongs to a dynamically discovered server
    KissCodeType_ntp_ref_mcst KissCodeType = "ntp-ref-mcst"

    // No key found. Either the key was never installed or not trusted
    KissCodeType_ntp_ref_nkey KissCodeType = "ntp-ref-nkey"

    // The server has temporarily denied access because
    // the client exceeded the rate threshold
    KissCodeType_ntp_ref_rate KissCodeType = "ntp-ref-rate"

    // Alteration of association from a remote host running ntpdc
    KissCodeType_ntp_ref_rmot KissCodeType = "ntp-ref-rmot"

    // STEP means the offset is less than the panic threshold but greater than the step threshold of 125 ms
    KissCodeType_ntp_ref_step KissCodeType = "ntp-ref-step"
)

// RefClockSourceType represents Clock source type for NTP
type RefClockSourceType string

const (
    // Geosynchronous Orbit Environment Satellite
    RefClockSourceType_ntp_ref_goes RefClockSourceType = "ntp-ref-goes"

    // Global Position System
    RefClockSourceType_ntp_ref_gps RefClockSourceType = "ntp-ref-gps"

    // Galileo Positioning System
    RefClockSourceType_ntp_ref_gal RefClockSourceType = "ntp-ref-gal"

    // Generic pulse-per-second
    RefClockSourceType_ntp_ref_pps RefClockSourceType = "ntp-ref-pps"

    // Inter-Range Instrumentation Group
    RefClockSourceType_ntp_ref_irig RefClockSourceType = "ntp-ref-irig"

    // LF Radio WWVB Ft. Collins
    RefClockSourceType_ntp_ref_wwvb RefClockSourceType = "ntp-ref-wwvb"

    // LF Radio DCF77 Mainflingen
    RefClockSourceType_ntp_ref_dcf RefClockSourceType = "ntp-ref-dcf"

    // LF Radio HBG Prangins
    RefClockSourceType_ntp_ref_hbg RefClockSourceType = "ntp-ref-hbg"

    // LF Radio MSF Anthorn
    RefClockSourceType_ntp_ref_msf RefClockSourceType = "ntp-ref-msf"

    // LF Radio JJY Fukushima
    RefClockSourceType_ntp_ref_jjy RefClockSourceType = "ntp-ref-jjy"

    // MF Radio LORAN C station
    RefClockSourceType_ntp_ref_lorc RefClockSourceType = "ntp-ref-lorc"

    // MF Radio Allouis
    RefClockSourceType_ntp_ref_tdf RefClockSourceType = "ntp-ref-tdf"

    // HF Radio CHU Ottawa
    RefClockSourceType_ntp_ref_chu RefClockSourceType = "ntp-ref-chu"

    // HF Radio WWV Ft. Collins
    RefClockSourceType_ntp_ref_wwv RefClockSourceType = "ntp-ref-wwv"

    // HF Radio WWVH Kauai
    RefClockSourceType_ntp_ref_wwvh RefClockSourceType = "ntp-ref-wwvh"

    // NIST telephone modem
    RefClockSourceType_ntp_ref_nist RefClockSourceType = "ntp-ref-nist"

    // NIST telephone modem
    RefClockSourceType_ntp_ref_acts RefClockSourceType = "ntp-ref-acts"

    // USNO telephone modem
    RefClockSourceType_ntp_ref_usno RefClockSourceType = "ntp-ref-usno"

    // European telephone modem
    RefClockSourceType_ntp_ref_ptb RefClockSourceType = "ntp-ref-ptb"
)

// RefidPktTypeInfo represents The type of information stored in the refid
type RefidPktTypeInfo string

const (
    // Kiss of Death code or KOD contains debug or maintenance code. Refid is set to these codes in stratums 0 and 16 (unspec,invalid, unsync)
    RefidPktTypeInfo_ntp_ref_state_kod RefidPktTypeInfo = "ntp-ref-state-kod"

    // CLK Source type occurs for all primary time servers in stratum 1
    RefidPktTypeInfo_ntp_ref_state_resolved_with_clk_source RefidPktTypeInfo = "ntp-ref-state-resolved-with-clk-source"

    // IP address occurs for clients in stratums >= 2 and  <=15 
    RefidPktTypeInfo_ntp_ref_state_resolved_with_ip_addr RefidPktTypeInfo = "ntp-ref-state-resolved-with-ip-addr"

    // Bad state which serves as a default criterion for  a complete mismatch with all cases
    RefidPktTypeInfo_ntp_ref_state_bad_state RefidPktTypeInfo = "ntp-ref-state-bad-state"
)

// NtpOperData
// NTP operational data
type NtpOperData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Contains ntp status info for the queried switch or router  which includes
    // reference identifier, reference time, stratum  delay and other details.
    NtpStatusInfo NtpOperData_NtpStatusInfo
}

func (ntpOperData *NtpOperData) GetEntityData() *types.CommonEntityData {
    ntpOperData.EntityData.YFilter = ntpOperData.YFilter
    ntpOperData.EntityData.YangName = "ntp-oper-data"
    ntpOperData.EntityData.BundleName = "cisco_ios_xe"
    ntpOperData.EntityData.ParentYangName = "Cisco-IOS-XE-ntp-oper"
    ntpOperData.EntityData.SegmentPath = "Cisco-IOS-XE-ntp-oper:ntp-oper-data"
    ntpOperData.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ntpOperData.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ntpOperData.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ntpOperData.EntityData.Children = types.NewOrderedMap()
    ntpOperData.EntityData.Children.Append("ntp-status-info", types.YChild{"NtpStatusInfo", &ntpOperData.NtpStatusInfo})
    ntpOperData.EntityData.Leafs = types.NewOrderedMap()

    ntpOperData.EntityData.YListKeys = []string {}

    return &(ntpOperData.EntityData)
}

// NtpOperData_NtpStatusInfo
// Contains ntp status info for the queried switch or router 
// which includes reference identifier, reference time, stratum 
// delay and other details
// This type is a presence type.
type NtpOperData_NtpStatusInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Unix calendar time. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    Reftime interface{}

    // Frequency or periodicity of NTP polling in seconds. The type is interface{}
    // with range: 0..255.
    SysPoll interface{}

    // How far away the current switch is in term of hops  from the primary time
    // source of the subnet or from the root of the subnet. The type is
    // interface{} with range: 0..4294967295.
    Stratum interface{}

    // Round trip delay with respect to the primary time server. The type is
    // string with range: -92233720368547758.08..92233720368547758.07.
    RootDelay interface{}

    // Deviation of offset with respect to time. All measurements are between the
    // current switch and the root of the subnet. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    RootDisp interface{}

    // Difference in time between current switch and peer and server clock. The
    // type is string with range: -92233720368547758.08..92233720368547758.07.
    Offset interface{}

    // The second derivative of offset with time. In NTP version 4, this is always
    // 0. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    FreqDriftPpm interface{}

    // Reference id can either be a KOD code or a clock source or an IP address.
    Refid NtpOperData_NtpStatusInfo_Refid

    // Table of NTP associations with servers and peers. The type is slice of
    // NtpOperData_NtpStatusInfo_NtpAssociations.
    NtpAssociations []*NtpOperData_NtpStatusInfo_NtpAssociations
}

func (ntpStatusInfo *NtpOperData_NtpStatusInfo) GetEntityData() *types.CommonEntityData {
    ntpStatusInfo.EntityData.YFilter = ntpStatusInfo.YFilter
    ntpStatusInfo.EntityData.YangName = "ntp-status-info"
    ntpStatusInfo.EntityData.BundleName = "cisco_ios_xe"
    ntpStatusInfo.EntityData.ParentYangName = "ntp-oper-data"
    ntpStatusInfo.EntityData.SegmentPath = "ntp-status-info"
    ntpStatusInfo.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ntpStatusInfo.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ntpStatusInfo.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ntpStatusInfo.EntityData.Children = types.NewOrderedMap()
    ntpStatusInfo.EntityData.Children.Append("refid", types.YChild{"Refid", &ntpStatusInfo.Refid})
    ntpStatusInfo.EntityData.Children.Append("ntp-associations", types.YChild{"NtpAssociations", nil})
    for i := range ntpStatusInfo.NtpAssociations {
        ntpStatusInfo.EntityData.Children.Append(types.GetSegmentPath(ntpStatusInfo.NtpAssociations[i]), types.YChild{"NtpAssociations", ntpStatusInfo.NtpAssociations[i]})
    }
    ntpStatusInfo.EntityData.Leafs = types.NewOrderedMap()
    ntpStatusInfo.EntityData.Leafs.Append("reftime", types.YLeaf{"Reftime", ntpStatusInfo.Reftime})
    ntpStatusInfo.EntityData.Leafs.Append("sys-poll", types.YLeaf{"SysPoll", ntpStatusInfo.SysPoll})
    ntpStatusInfo.EntityData.Leafs.Append("stratum", types.YLeaf{"Stratum", ntpStatusInfo.Stratum})
    ntpStatusInfo.EntityData.Leafs.Append("root-delay", types.YLeaf{"RootDelay", ntpStatusInfo.RootDelay})
    ntpStatusInfo.EntityData.Leafs.Append("root-disp", types.YLeaf{"RootDisp", ntpStatusInfo.RootDisp})
    ntpStatusInfo.EntityData.Leafs.Append("offset", types.YLeaf{"Offset", ntpStatusInfo.Offset})
    ntpStatusInfo.EntityData.Leafs.Append("freq-drift-ppm", types.YLeaf{"FreqDriftPpm", ntpStatusInfo.FreqDriftPpm})

    ntpStatusInfo.EntityData.YListKeys = []string {}

    return &(ntpStatusInfo.EntityData)
}

// NtpOperData_NtpStatusInfo_Refid
// Reference id can either be a KOD code or a clock source or an IP address
type NtpOperData_NtpStatusInfo_Refid struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPV4 or IPV6 ip address. The type is one of the following types: string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddr interface{}

    // Bad stat or exception code in case the 3 criteria of ip, clock and kod
    // don't match. The type is interface{} with range: 0..4294967295.
    ExceptionCode interface{}

    // Container for KOD type eg INIT ACTS.
    KodData NtpOperData_NtpStatusInfo_Refid_KodData

    // Container for clock data. GPS is the only source supported by Cisco
    // currrently.
    RefClkSrcData NtpOperData_NtpStatusInfo_Refid_RefClkSrcData
}

func (refid *NtpOperData_NtpStatusInfo_Refid) GetEntityData() *types.CommonEntityData {
    refid.EntityData.YFilter = refid.YFilter
    refid.EntityData.YangName = "refid"
    refid.EntityData.BundleName = "cisco_ios_xe"
    refid.EntityData.ParentYangName = "ntp-status-info"
    refid.EntityData.SegmentPath = "refid"
    refid.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    refid.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    refid.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    refid.EntityData.Children = types.NewOrderedMap()
    refid.EntityData.Children.Append("kod-data", types.YChild{"KodData", &refid.KodData})
    refid.EntityData.Children.Append("ref-clk-src-data", types.YChild{"RefClkSrcData", &refid.RefClkSrcData})
    refid.EntityData.Leafs = types.NewOrderedMap()
    refid.EntityData.Leafs.Append("ip-addr", types.YLeaf{"IpAddr", refid.IpAddr})
    refid.EntityData.Leafs.Append("exception-code", types.YLeaf{"ExceptionCode", refid.ExceptionCode})

    refid.EntityData.YListKeys = []string {}

    return &(refid.EntityData)
}

// NtpOperData_NtpStatusInfo_Refid_KodData
// Container for KOD type eg INIT ACTS
type NtpOperData_NtpStatusInfo_Refid_KodData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // KOD types could be any of  the enums including INIT   ACTS etc. The type is
    // KissCodeType.
    KodType interface{}
}

func (kodData *NtpOperData_NtpStatusInfo_Refid_KodData) GetEntityData() *types.CommonEntityData {
    kodData.EntityData.YFilter = kodData.YFilter
    kodData.EntityData.YangName = "kod-data"
    kodData.EntityData.BundleName = "cisco_ios_xe"
    kodData.EntityData.ParentYangName = "refid"
    kodData.EntityData.SegmentPath = "kod-data"
    kodData.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    kodData.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    kodData.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    kodData.EntityData.Children = types.NewOrderedMap()
    kodData.EntityData.Leafs = types.NewOrderedMap()
    kodData.EntityData.Leafs.Append("kod-type", types.YLeaf{"KodType", kodData.KodType})

    kodData.EntityData.YListKeys = []string {}

    return &(kodData.EntityData)
}

// NtpOperData_NtpStatusInfo_Refid_RefClkSrcData
// Container for clock data. GPS is the only source supported by Cisco currrently
type NtpOperData_NtpStatusInfo_Refid_RefClkSrcData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Contains clock source type specifics eg GPS and container extensions. The
    // type is RefClockSourceType.
    RefClkSrcType interface{}
}

func (refClkSrcData *NtpOperData_NtpStatusInfo_Refid_RefClkSrcData) GetEntityData() *types.CommonEntityData {
    refClkSrcData.EntityData.YFilter = refClkSrcData.YFilter
    refClkSrcData.EntityData.YangName = "ref-clk-src-data"
    refClkSrcData.EntityData.BundleName = "cisco_ios_xe"
    refClkSrcData.EntityData.ParentYangName = "refid"
    refClkSrcData.EntityData.SegmentPath = "ref-clk-src-data"
    refClkSrcData.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    refClkSrcData.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    refClkSrcData.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    refClkSrcData.EntityData.Children = types.NewOrderedMap()
    refClkSrcData.EntityData.Leafs = types.NewOrderedMap()
    refClkSrcData.EntityData.Leafs.Append("ref-clk-src-type", types.YLeaf{"RefClkSrcType", refClkSrcData.RefClkSrcType})

    refClkSrcData.EntityData.YListKeys = []string {}

    return &(refClkSrcData.EntityData)
}

// NtpOperData_NtpStatusInfo_NtpAssociations
// Table of NTP associations with servers and peers
type NtpOperData_NtpStatusInfo_NtpAssociations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Association id is a descriptor which describes the
    // association between two NTP entities whether client and peer or client and
    // server. The type is interface{} with range: 0..65535.
    AssocId interface{}

    // The status of the last 8 NTP packet exchanges with peers. 1 is encoded in
    // the bitmask for a successful attempt and 0 is encoded for failure. If all
    // the last 8 transactions with peers or messages sent to peers are successful
    // the encoding becomes 0xff. The type is interface{} with range: 0..255.
    PeerReach interface{}

    // Stratum in which the peer exists. The type is interface{} with range:
    // 0..4294967295.
    PeerStratum interface{}

    // Reference UNIX calendar time. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    Reftime interface{}

    // The time of  the last NTP poll or update that happened in seconds. How many
    // seconds back did the last update happen or when did the last NTP update
    // happen ?. The type is interface{} with range: 0..18446744073709551615.
    LastPollTime interface{}

    // Maximum poll time of NTP in seconds. The type is interface{} with range:
    // 0..4294967295.
    Poll interface{}

    // Round trip delay of reaching the peer and returning. The type is string
    // with range: -92233720368547758.08..92233720368547758.07.
    Delay interface{}

    // Difference in ms between server time and local time. offset with respect to
    // the peer/server clock. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    Offset interface{}

    // Jitter in ms refers to short-term variations in frequency of components
    // greater than 10 hz. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    Jitter interface{}

    // refid refers to either an IP address or a clock source or KOD type code.
    Refid NtpOperData_NtpStatusInfo_NtpAssociations_Refid
}

func (ntpAssociations *NtpOperData_NtpStatusInfo_NtpAssociations) GetEntityData() *types.CommonEntityData {
    ntpAssociations.EntityData.YFilter = ntpAssociations.YFilter
    ntpAssociations.EntityData.YangName = "ntp-associations"
    ntpAssociations.EntityData.BundleName = "cisco_ios_xe"
    ntpAssociations.EntityData.ParentYangName = "ntp-status-info"
    ntpAssociations.EntityData.SegmentPath = "ntp-associations" + types.AddKeyToken(ntpAssociations.AssocId, "assoc-id")
    ntpAssociations.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ntpAssociations.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ntpAssociations.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ntpAssociations.EntityData.Children = types.NewOrderedMap()
    ntpAssociations.EntityData.Children.Append("refid", types.YChild{"Refid", &ntpAssociations.Refid})
    ntpAssociations.EntityData.Leafs = types.NewOrderedMap()
    ntpAssociations.EntityData.Leafs.Append("assoc-id", types.YLeaf{"AssocId", ntpAssociations.AssocId})
    ntpAssociations.EntityData.Leafs.Append("peer-reach", types.YLeaf{"PeerReach", ntpAssociations.PeerReach})
    ntpAssociations.EntityData.Leafs.Append("peer-stratum", types.YLeaf{"PeerStratum", ntpAssociations.PeerStratum})
    ntpAssociations.EntityData.Leafs.Append("reftime", types.YLeaf{"Reftime", ntpAssociations.Reftime})
    ntpAssociations.EntityData.Leafs.Append("last-poll-time", types.YLeaf{"LastPollTime", ntpAssociations.LastPollTime})
    ntpAssociations.EntityData.Leafs.Append("poll", types.YLeaf{"Poll", ntpAssociations.Poll})
    ntpAssociations.EntityData.Leafs.Append("delay", types.YLeaf{"Delay", ntpAssociations.Delay})
    ntpAssociations.EntityData.Leafs.Append("offset", types.YLeaf{"Offset", ntpAssociations.Offset})
    ntpAssociations.EntityData.Leafs.Append("jitter", types.YLeaf{"Jitter", ntpAssociations.Jitter})

    ntpAssociations.EntityData.YListKeys = []string {"AssocId"}

    return &(ntpAssociations.EntityData)
}

// NtpOperData_NtpStatusInfo_NtpAssociations_Refid
// refid refers to either an IP address or a clock source or KOD type code
type NtpOperData_NtpStatusInfo_NtpAssociations_Refid struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPV4 or IPV6 ip address. The type is one of the following types: string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddr interface{}

    // Bad stat or exception code in case the 3 criteria of ip, clock and kod
    // don't match. The type is interface{} with range: 0..4294967295.
    ExceptionCode interface{}

    // Container for KOD type eg INIT ACTS.
    KodData NtpOperData_NtpStatusInfo_NtpAssociations_Refid_KodData

    // Container for clock data. GPS is the only source supported by Cisco
    // currrently.
    RefClkSrcData NtpOperData_NtpStatusInfo_NtpAssociations_Refid_RefClkSrcData
}

func (refid *NtpOperData_NtpStatusInfo_NtpAssociations_Refid) GetEntityData() *types.CommonEntityData {
    refid.EntityData.YFilter = refid.YFilter
    refid.EntityData.YangName = "refid"
    refid.EntityData.BundleName = "cisco_ios_xe"
    refid.EntityData.ParentYangName = "ntp-associations"
    refid.EntityData.SegmentPath = "refid"
    refid.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    refid.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    refid.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    refid.EntityData.Children = types.NewOrderedMap()
    refid.EntityData.Children.Append("kod-data", types.YChild{"KodData", &refid.KodData})
    refid.EntityData.Children.Append("ref-clk-src-data", types.YChild{"RefClkSrcData", &refid.RefClkSrcData})
    refid.EntityData.Leafs = types.NewOrderedMap()
    refid.EntityData.Leafs.Append("ip-addr", types.YLeaf{"IpAddr", refid.IpAddr})
    refid.EntityData.Leafs.Append("exception-code", types.YLeaf{"ExceptionCode", refid.ExceptionCode})

    refid.EntityData.YListKeys = []string {}

    return &(refid.EntityData)
}

// NtpOperData_NtpStatusInfo_NtpAssociations_Refid_KodData
// Container for KOD type eg INIT ACTS
type NtpOperData_NtpStatusInfo_NtpAssociations_Refid_KodData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // KOD types could be any of  the enums including INIT   ACTS etc. The type is
    // KissCodeType.
    KodType interface{}
}

func (kodData *NtpOperData_NtpStatusInfo_NtpAssociations_Refid_KodData) GetEntityData() *types.CommonEntityData {
    kodData.EntityData.YFilter = kodData.YFilter
    kodData.EntityData.YangName = "kod-data"
    kodData.EntityData.BundleName = "cisco_ios_xe"
    kodData.EntityData.ParentYangName = "refid"
    kodData.EntityData.SegmentPath = "kod-data"
    kodData.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    kodData.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    kodData.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    kodData.EntityData.Children = types.NewOrderedMap()
    kodData.EntityData.Leafs = types.NewOrderedMap()
    kodData.EntityData.Leafs.Append("kod-type", types.YLeaf{"KodType", kodData.KodType})

    kodData.EntityData.YListKeys = []string {}

    return &(kodData.EntityData)
}

// NtpOperData_NtpStatusInfo_NtpAssociations_Refid_RefClkSrcData
// Container for clock data. GPS is the only source supported by Cisco currrently
type NtpOperData_NtpStatusInfo_NtpAssociations_Refid_RefClkSrcData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Contains clock source type specifics eg GPS and container extensions. The
    // type is RefClockSourceType.
    RefClkSrcType interface{}
}

func (refClkSrcData *NtpOperData_NtpStatusInfo_NtpAssociations_Refid_RefClkSrcData) GetEntityData() *types.CommonEntityData {
    refClkSrcData.EntityData.YFilter = refClkSrcData.YFilter
    refClkSrcData.EntityData.YangName = "ref-clk-src-data"
    refClkSrcData.EntityData.BundleName = "cisco_ios_xe"
    refClkSrcData.EntityData.ParentYangName = "refid"
    refClkSrcData.EntityData.SegmentPath = "ref-clk-src-data"
    refClkSrcData.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    refClkSrcData.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    refClkSrcData.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    refClkSrcData.EntityData.Children = types.NewOrderedMap()
    refClkSrcData.EntityData.Leafs = types.NewOrderedMap()
    refClkSrcData.EntityData.Leafs.Append("ref-clk-src-type", types.YLeaf{"RefClkSrcType", refClkSrcData.RefClkSrcType})

    refClkSrcData.EntityData.YListKeys = []string {}

    return &(refClkSrcData.EntityData)
}

