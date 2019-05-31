// This module contains a collection of YANG definitions
// for NTP operational data.
// Copyright (c) 2017-2018 by Cisco Systems, Inc.
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

// PeerSelectStatus represents Selection status of peer
type PeerSelectStatus string

const (
    // The peer is a survivor but not among the first 
    // six peers
    PeerSelectStatus_ntp_peer_as_backup PeerSelectStatus = "ntp-peer-as-backup"

    // The peer was rejected due to a loop or due
    // to becoming unreachable or due to bad synchronization distance
    PeerSelectStatus_ntp_peer_rejected PeerSelectStatus = "ntp-peer-rejected"

    // The peer or server is discarded due to false tick
    // or clock errors
    PeerSelectStatus_ntp_peer_false_ticker PeerSelectStatus = "ntp-peer-false-ticker"

    // The peer is discarded as it is not among the 
    // first ten peers sorted by synchronization distance
    PeerSelectStatus_ntp_peer_excess PeerSelectStatus = "ntp-peer-excess"

    // NTP server or peer rejected as outlier
    PeerSelectStatus_ntp_peer_outlier PeerSelectStatus = "ntp-peer-outlier"

    // Possible candidate for selection as time server
    PeerSelectStatus_ntp_peer_candidate PeerSelectStatus = "ntp-peer-candidate"

    // Peer or server selected as time server
    PeerSelectStatus_ntp_peer_sys_peer PeerSelectStatus = "ntp-peer-sys-peer"

    // Peer or server selected as time server. In this 
    // case the Pulse Per Second signal is used to synchronize the client and
    // server or peer
    PeerSelectStatus_ntp_peer_pps_peer PeerSelectStatus = "ntp-peer-pps-peer"
)

// PeerAuthStatus represents Status of authenticating switch with peer or server
type PeerAuthStatus string

const (
    // The NTP client or server packet has MAC  
    // field and authentication succeded
    PeerAuthStatus_ntp_auth_ok PeerAuthStatus = "ntp-auth-ok"

    // The NTP client or server packet has MAC  
    // and decryption failed with AUTH_ERROR 
    PeerAuthStatus_ntp_auth_bad_auth PeerAuthStatus = "ntp-auth-bad-auth"

    // The NTP client or server is not configured with authenication
    // with server or client
    PeerAuthStatus_ntp_auth_auth_not_configured PeerAuthStatus = "ntp-auth-auth-not-configured"

    // The NTP client or server authentication stautus is not available
    //  as now
    PeerAuthStatus_ntp_auth_status_not_available PeerAuthStatus = "ntp-auth-status-not-available"

    // The NTP client or server packet has no MAC  
    // with server or peer
    PeerAuthStatus_ntp_auth_none PeerAuthStatus = "ntp-auth-none"

    // crypto-NAK. The MAC has four octets only and could not  
    // determine authentication status with peer
    PeerAuthStatus_ntp_auth_crypto PeerAuthStatus = "ntp-auth-crypto"
)

// PeerStatusWord represents Peer status word or crypto of ntp server or ntp peer
type PeerStatusWord string

const (
    // In autokey[public key ntp authentication protocol ], this flag is 
    // set when host certificate is signed by server.This is not implemented/supported as of now 
    PeerStatusWord_crypto_flag_sig PeerStatusWord = "crypto-flag-sig"

    // In autokey, this flag is set when leap second values
    // are received and validated
    PeerStatusWord_crypto_flag_leap PeerStatusWord = "crypto-flag-leap"

    // In autokey,this flag is set when the trusted host identity
    // credentials are  confirmed 
    PeerStatusWord_crypto_flag_vrfy PeerStatusWord = "crypto-flag-vrfy"

    // In autokey, this flag is set when the cookie is received  and validated
    // when set, keylists with  nonzero cookies are generated, 
    // when reset cookie is zero
    PeerStatusWord_crypto_flag_cook PeerStatusWord = "crypto-flag-cook"

    // In autokey, this flag is set when autokey values are received and validated,
    // when set client can validate packets without extension field 
    // according to the autokey sequence
    PeerStatusWord_crypto_flag_auto PeerStatusWord = "crypto-flag-auto"

    // In autokey, this flag is set when trusted host certificate and public key are verified
    PeerStatusWord_crypto_flag_cert PeerStatusWord = "crypto-flag-cert"
)

// PeerEvent represents Event received by switch or router and sent by peer
type PeerEvent string

const (
    // This event is used to allocate resources and 
    // initialize defaults or values when a NTP association is setup
    PeerEvent_ntp_peer_event_mobilize PeerEvent = "ntp-peer-event-mobilize"

    // This event is used to tear down the resources
    // associated with a NTP association
    PeerEvent_ntp_peer_event_demobilize PeerEvent = "ntp-peer-event-demobilize"

    // This event indicates that the NTP peer is 
    // unreachable
    PeerEvent_ntp_peer_event_unreachable PeerEvent = "ntp-peer-event-unreachable"

    // This event indicates that the peer is reachable
    PeerEvent_ntp_peer_event_reachable PeerEvent = "ntp-peer-event-reachable"

    // Event to indicate that the NTP process restart
    // is now complete
    PeerEvent_ntp_peer_event_restart PeerEvent = "ntp-peer-event-restart"

    // NTP peer or server reply event in response to a 
    // get time request from the client
    PeerEvent_ntp_peer_event_reply PeerEvent = "ntp-peer-event-reply"

    // This event is used to synchronize the client 
    // and peer or server through a flow contol mechanism
    PeerEvent_ntp_peer_event_rate PeerEvent = "ntp-peer-event-rate"

    // Event from peer that indicates denial of access to the 
    // switch or router
    PeerEvent_ntp_peer_event_deny PeerEvent = "ntp-peer-event-deny"

    // This event clears or resets the NTP flag after the
    // leap second event so that the system is now ready to receive the next
    // leap second event
    PeerEvent_ntp_peer_disarmed PeerEvent = "ntp-peer-disarmed"

    // Peer event armed means that the event for delaying
    // the clock increment by one second for a leap year will be scheduled 
    // next month
    PeerEvent_ntp_peer_armed PeerEvent = "ntp-peer-armed"

    // New peer added to association
    PeerEvent_ntp_peer_event_newpeer PeerEvent = "ntp-peer-event-newpeer"

    // This event indicates clock tick errors
    PeerEvent_ntp_peer_event_clock PeerEvent = "ntp-peer-event-clock"

    // Event indicating status of authenticating switch
    // or router with peer
    PeerEvent_ntp_peer_event_auth PeerEvent = "ntp-peer-event-auth"

    // Popcorn event indicates a delayed NTP packet due
    // to congestion in the network
    PeerEvent_ntp_peer_event_popcorn PeerEvent = "ntp-peer-event-popcorn"

    // Event for NTP peer or server leaving the 
    // association
    PeerEvent_ntp_peer_event_xleave PeerEvent = "ntp-peer-event-xleave"

    // NTP event for an error message received from 
    // peer or server
    PeerEvent_ntp_peer_event_xerr PeerEvent = "ntp-peer-event-xerr"

    // Event for incorporating correction for 
    // International Atomic Time based on offsets from UTC
    PeerEvent_ntp_peer_event_tai PeerEvent = "ntp-peer-event-tai"
)

// ServerType represents Status of remote entity whether server or peer
type ServerType string

const (
    // Remote entity is a NTP peer
    ServerType_ntp_peer ServerType = "ntp-peer"

    // Remote NTP is a NTP server
    ServerType_ntp_server ServerType = "ntp-server"

    // Status of remote entity could not be found
    ServerType_ntp_unknown_type ServerType = "ntp-unknown-type"
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
    ntpOperData.EntityData.AbsolutePath = ntpOperData.EntityData.SegmentPath
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
    // b'\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}(\\.\\d+)?(Z|[\\+\\-]\\d{2}:\\d{2})'.
    Reftime interface{}

    // Frequency or periodicity of NTP polling in seconds expressed as a power of
    // two as per model and rfc. The type is interface{} with range: 0..255.
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
    ntpStatusInfo.EntityData.AbsolutePath = "Cisco-IOS-XE-ntp-oper:ntp-oper-data/" + ntpStatusInfo.EntityData.SegmentPath
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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
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
    refid.EntityData.AbsolutePath = "Cisco-IOS-XE-ntp-oper:ntp-oper-data/ntp-status-info/" + refid.EntityData.SegmentPath
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
    kodData.EntityData.AbsolutePath = "Cisco-IOS-XE-ntp-oper:ntp-oper-data/ntp-status-info/refid/" + kodData.EntityData.SegmentPath
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
    refClkSrcData.EntityData.AbsolutePath = "Cisco-IOS-XE-ntp-oper:ntp-oper-data/ntp-status-info/refid/" + refClkSrcData.EntityData.SegmentPath
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
    YListKey string

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
    // b'\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}(\\.\\d+)?(Z|[\\+\\-]\\d{2}:\\d{2})'.
    Reftime interface{}

    // The time of  the last NTP poll or update that happened in seconds. How many
    // seconds back did the last update happen or when did the last NTP update
    // happen ?. The type is interface{} with range: 0..18446744073709551615.
    LastPollTime interface{}

    // Maximum poll time of NTP in seconds expressed as a power of two as per
    // model and RFC. The type is interface{} with range: 0..4294967295.
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

    // Count of number of error events received from peer. The type is interface{}
    // with range: 0..255.
    NumEvents interface{}

    // Last event received from peer. The type is PeerEvent.
    LastPeerEvent interface{}

    // Status of peer selection based on the NTP selection  algorithm. The type is
    // PeerSelectStatus.
    PeerSelectionStatus interface{}

    // Status of authentication of switch or router by peer. The type is
    // PeerAuthStatus.
    PeerAuthenticationStatus interface{}

    // Whether the remote NTP device is a server or peer. The type is ServerType.
    ServType interface{}

    // Peer status word of ntp server or peer when authentication configured. The
    // type is PeerStatusWord.
    PswCrypto interface{}

    // refid refers to either an IP address or a clock source or KOD type code.
    Refid NtpOperData_NtpStatusInfo_NtpAssociations_Refid

    // NTP address consists of an IP address and a VRF name.
    NtpAddress NtpOperData_NtpStatusInfo_NtpAssociations_NtpAddress
}

func (ntpAssociations *NtpOperData_NtpStatusInfo_NtpAssociations) GetEntityData() *types.CommonEntityData {
    ntpAssociations.EntityData.YFilter = ntpAssociations.YFilter
    ntpAssociations.EntityData.YangName = "ntp-associations"
    ntpAssociations.EntityData.BundleName = "cisco_ios_xe"
    ntpAssociations.EntityData.ParentYangName = "ntp-status-info"
    ntpAssociations.EntityData.SegmentPath = "ntp-associations" + types.AddKeyToken(ntpAssociations.AssocId, "assoc-id")
    ntpAssociations.EntityData.AbsolutePath = "Cisco-IOS-XE-ntp-oper:ntp-oper-data/ntp-status-info/" + ntpAssociations.EntityData.SegmentPath
    ntpAssociations.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ntpAssociations.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ntpAssociations.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ntpAssociations.EntityData.Children = types.NewOrderedMap()
    ntpAssociations.EntityData.Children.Append("refid", types.YChild{"Refid", &ntpAssociations.Refid})
    ntpAssociations.EntityData.Children.Append("ntp-address", types.YChild{"NtpAddress", &ntpAssociations.NtpAddress})
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
    ntpAssociations.EntityData.Leafs.Append("num-events", types.YLeaf{"NumEvents", ntpAssociations.NumEvents})
    ntpAssociations.EntityData.Leafs.Append("last-peer-event", types.YLeaf{"LastPeerEvent", ntpAssociations.LastPeerEvent})
    ntpAssociations.EntityData.Leafs.Append("peer-selection-status", types.YLeaf{"PeerSelectionStatus", ntpAssociations.PeerSelectionStatus})
    ntpAssociations.EntityData.Leafs.Append("peer-authentication-status", types.YLeaf{"PeerAuthenticationStatus", ntpAssociations.PeerAuthenticationStatus})
    ntpAssociations.EntityData.Leafs.Append("serv-type", types.YLeaf{"ServType", ntpAssociations.ServType})
    ntpAssociations.EntityData.Leafs.Append("psw-crypto", types.YLeaf{"PswCrypto", ntpAssociations.PswCrypto})

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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
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
    refid.EntityData.AbsolutePath = "Cisco-IOS-XE-ntp-oper:ntp-oper-data/ntp-status-info/ntp-associations/" + refid.EntityData.SegmentPath
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
    kodData.EntityData.AbsolutePath = "Cisco-IOS-XE-ntp-oper:ntp-oper-data/ntp-status-info/ntp-associations/refid/" + kodData.EntityData.SegmentPath
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
    refClkSrcData.EntityData.AbsolutePath = "Cisco-IOS-XE-ntp-oper:ntp-oper-data/ntp-status-info/ntp-associations/refid/" + refClkSrcData.EntityData.SegmentPath
    refClkSrcData.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    refClkSrcData.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    refClkSrcData.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    refClkSrcData.EntityData.Children = types.NewOrderedMap()
    refClkSrcData.EntityData.Leafs = types.NewOrderedMap()
    refClkSrcData.EntityData.Leafs.Append("ref-clk-src-type", types.YLeaf{"RefClkSrcType", refClkSrcData.RefClkSrcType})

    refClkSrcData.EntityData.YListKeys = []string {}

    return &(refClkSrcData.EntityData)
}

// NtpOperData_NtpStatusInfo_NtpAssociations_NtpAddress
// NTP address consists of an IP address and a VRF name
type NtpOperData_NtpStatusInfo_NtpAssociations_NtpAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP address is the IP address of  the NTP server or peer. The type is one of
    // the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    IpAddr interface{}

    // VRF name is the virtual routing instance through which we can find the ntp
    // server or peer. The type is string.
    VrfName interface{}
}

func (ntpAddress *NtpOperData_NtpStatusInfo_NtpAssociations_NtpAddress) GetEntityData() *types.CommonEntityData {
    ntpAddress.EntityData.YFilter = ntpAddress.YFilter
    ntpAddress.EntityData.YangName = "ntp-address"
    ntpAddress.EntityData.BundleName = "cisco_ios_xe"
    ntpAddress.EntityData.ParentYangName = "ntp-associations"
    ntpAddress.EntityData.SegmentPath = "ntp-address"
    ntpAddress.EntityData.AbsolutePath = "Cisco-IOS-XE-ntp-oper:ntp-oper-data/ntp-status-info/ntp-associations/" + ntpAddress.EntityData.SegmentPath
    ntpAddress.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ntpAddress.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ntpAddress.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ntpAddress.EntityData.Children = types.NewOrderedMap()
    ntpAddress.EntityData.Leafs = types.NewOrderedMap()
    ntpAddress.EntityData.Leafs.Append("ip-addr", types.YLeaf{"IpAddr", ntpAddress.IpAddr})
    ntpAddress.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", ntpAddress.VrfName})

    ntpAddress.EntityData.YListKeys = []string {}

    return &(ntpAddress.EntityData)
}

