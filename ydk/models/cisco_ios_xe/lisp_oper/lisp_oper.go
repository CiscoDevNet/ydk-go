// This module contains a collection of YANG definitions for
// LISP operational data.
// Copyright (c) 2017-2018 by Cisco Systems, Inc.
// All rights reserved.
package lisp_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package lisp_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-lisp-oper lisp-state}", reflect.TypeOf(LispState{}))
    ydk.RegisterEntity("Cisco-IOS-XE-lisp-oper:lisp-state", reflect.TypeOf(LispState{}))
}

// LispAddressFamilyType represents Address-family of a LISP address or prefix
type LispAddressFamilyType string

const (
    // IANA IPv4 address family
    LispAddressFamilyType_ipv4_afi LispAddressFamilyType = "ipv4-afi"

    // IANA IPv6 address family
    LispAddressFamilyType_ipv6_afi LispAddressFamilyType = "ipv6-afi"

    // IANA MAC address family
    LispAddressFamilyType_mac_afi LispAddressFamilyType = "mac-afi"
)

// LispRlocStateType represents Reachability state of a RLOC
type LispRlocStateType string

const (
    // Locator is down or unreachable
    LispRlocStateType_lisp_rloc_state_down LispRlocStateType = "lisp-rloc-state-down"

    // Locator is up and reachable
    LispRlocStateType_lisp_rloc_state_up LispRlocStateType = "lisp-rloc-state-up"
)

// LispMapReplyActionType represents for negative map-cache entries
type LispMapReplyActionType string

const (
    // Mapping is kept alive and no encapsulation occurs
    LispMapReplyActionType_no_action LispMapReplyActionType = "no-action"

    // Matching packets are forwarded without
    // LISP encapsulation
    LispMapReplyActionType_natively_forward LispMapReplyActionType = "natively-forward"

    // Matching packets trigger sending Map-Requests
    LispMapReplyActionType_send_map_request LispMapReplyActionType = "send-map-request"

    // Matching packets are dropped
    LispMapReplyActionType_drop LispMapReplyActionType = "drop"
)

// LispIaftypeType represents LISP instance
type LispIaftypeType string

const (
    // IPv4 instance service
    LispIaftypeType_iaf_ipv4 LispIaftypeType = "iaf-ipv4"

    // IPv6 instance service
    LispIaftypeType_iaf_ipv6 LispIaftypeType = "iaf-ipv6"

    // MAC (L2) instance service
    LispIaftypeType_iaf_mac LispIaftypeType = "iaf-mac"

    // Address Resolution (L3 address-to-MAC) instance
    // service
    LispIaftypeType_iaf_ar LispIaftypeType = "iaf-ar"

    // Reverse Address Resolution (MAC-to-L3 address)
    // instance service
    LispIaftypeType_iaf_rar LispIaftypeType = "iaf-rar"
)

// LispSessionStateType represents State of a TCP session
type LispSessionStateType string

const (
    // Session parameters are incomplete
    LispSessionStateType_lisp_session_state_incomplete LispSessionStateType = "lisp-session-state-incomplete"

    // Session represents the passively listening socket
    LispSessionStateType_lisp_session_state_listening LispSessionStateType = "lisp-session-state-listening"

    // Session is down
    LispSessionStateType_lisp_session_state_down LispSessionStateType = "lisp-session-state-down"

    // Session is up
    LispSessionStateType_lisp_session_state_up LispSessionStateType = "lisp-session-state-up"
)

// LispState
// Operational state of the LISP subsystem
type LispState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of LISP routers. The type is slice of LispState_LispRouters.
    LispRouters []*LispState_LispRouters
}

func (lispState *LispState) GetEntityData() *types.CommonEntityData {
    lispState.EntityData.YFilter = lispState.YFilter
    lispState.EntityData.YangName = "lisp-state"
    lispState.EntityData.BundleName = "cisco_ios_xe"
    lispState.EntityData.ParentYangName = "Cisco-IOS-XE-lisp-oper"
    lispState.EntityData.SegmentPath = "Cisco-IOS-XE-lisp-oper:lisp-state"
    lispState.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lispState.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lispState.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lispState.EntityData.Children = types.NewOrderedMap()
    lispState.EntityData.Children.Append("lisp-routers", types.YChild{"LispRouters", nil})
    for i := range lispState.LispRouters {
        lispState.EntityData.Children.Append(types.GetSegmentPath(lispState.LispRouters[i]), types.YChild{"LispRouters", lispState.LispRouters[i]})
    }
    lispState.EntityData.Leafs = types.NewOrderedMap()

    lispState.EntityData.YListKeys = []string {}

    return &(lispState.EntityData)
}

// LispState_LispRouters
// List of LISP routers
type LispState_LispRouters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. ID number of the LISP router. The type is
    // interface{} with range: 0..4294967295.
    TopId interface{}

    // Site-ID common to all devices attached to the same site. The type is
    // interface{} with range: 0..18446744073709551615.
    SiteId interface{}

    // xTR-ID of the device. The type is slice of interface{} with range: 0..255.
    XtrId []interface{}

    // List of LISP instances. The type is slice of
    // LispState_LispRouters_Instances.
    Instances []*LispState_LispRouters_Instances

    // List of Reliable Registration sessions. The type is slice of
    // LispState_LispRouters_Sessions.
    Sessions []*LispState_LispRouters_Sessions

    // This list represents the set of routing locators configured on this device.
    // The type is slice of LispState_LispRouters_LocalRlocs.
    LocalRlocs []*LispState_LispRouters_LocalRlocs
}

func (lispRouters *LispState_LispRouters) GetEntityData() *types.CommonEntityData {
    lispRouters.EntityData.YFilter = lispRouters.YFilter
    lispRouters.EntityData.YangName = "lisp-routers"
    lispRouters.EntityData.BundleName = "cisco_ios_xe"
    lispRouters.EntityData.ParentYangName = "lisp-state"
    lispRouters.EntityData.SegmentPath = "lisp-routers" + types.AddKeyToken(lispRouters.TopId, "top-id")
    lispRouters.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lispRouters.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lispRouters.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lispRouters.EntityData.Children = types.NewOrderedMap()
    lispRouters.EntityData.Children.Append("instances", types.YChild{"Instances", nil})
    for i := range lispRouters.Instances {
        lispRouters.EntityData.Children.Append(types.GetSegmentPath(lispRouters.Instances[i]), types.YChild{"Instances", lispRouters.Instances[i]})
    }
    lispRouters.EntityData.Children.Append("sessions", types.YChild{"Sessions", nil})
    for i := range lispRouters.Sessions {
        lispRouters.EntityData.Children.Append(types.GetSegmentPath(lispRouters.Sessions[i]), types.YChild{"Sessions", lispRouters.Sessions[i]})
    }
    lispRouters.EntityData.Children.Append("local-rlocs", types.YChild{"LocalRlocs", nil})
    for i := range lispRouters.LocalRlocs {
        lispRouters.EntityData.Children.Append(types.GetSegmentPath(lispRouters.LocalRlocs[i]), types.YChild{"LocalRlocs", lispRouters.LocalRlocs[i]})
    }
    lispRouters.EntityData.Leafs = types.NewOrderedMap()
    lispRouters.EntityData.Leafs.Append("top-id", types.YLeaf{"TopId", lispRouters.TopId})
    lispRouters.EntityData.Leafs.Append("site-id", types.YLeaf{"SiteId", lispRouters.SiteId})
    lispRouters.EntityData.Leafs.Append("xtr-id", types.YLeaf{"XtrId", lispRouters.XtrId})

    lispRouters.EntityData.YListKeys = []string {"TopId"}

    return &(lispRouters.EntityData)
}

// LispState_LispRouters_Instances
// List of LISP instances
type LispState_LispRouters_Instances struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. LISP Instance ID. The type is interface{} with
    // range: 0..16777215.
    Iid interface{}

    // Name of VRF that is mapped to the given LISP Instance ID. The type is
    // string.
    VrfName interface{}

    // Status of RLOC Probing. The type is bool.
    IsRlocProbing interface{}

    // List of Address-Families enabled in this LISP instance. The type is slice
    // of LispState_LispRouters_Instances_Af.
    Af []*LispState_LispRouters_Instances_Af

    // MS registration EID membership list (list of locators known to the MS as
    // allowed to send traffic in the instance). The type is slice of
    // LispState_LispRouters_Instances_MsEidMembership.
    MsEidMembership []*LispState_LispRouters_Instances_MsEidMembership

    // ETR EID membership list (list of locators known to the ETR as allowed to
    // send traffic in the instance). The type is slice of
    // LispState_LispRouters_Instances_EtrEidMembership.
    EtrEidMembership []*LispState_LispRouters_Instances_EtrEidMembership
}

func (instances *LispState_LispRouters_Instances) GetEntityData() *types.CommonEntityData {
    instances.EntityData.YFilter = instances.YFilter
    instances.EntityData.YangName = "instances"
    instances.EntityData.BundleName = "cisco_ios_xe"
    instances.EntityData.ParentYangName = "lisp-routers"
    instances.EntityData.SegmentPath = "instances" + types.AddKeyToken(instances.Iid, "iid")
    instances.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    instances.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    instances.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    instances.EntityData.Children = types.NewOrderedMap()
    instances.EntityData.Children.Append("af", types.YChild{"Af", nil})
    for i := range instances.Af {
        instances.EntityData.Children.Append(types.GetSegmentPath(instances.Af[i]), types.YChild{"Af", instances.Af[i]})
    }
    instances.EntityData.Children.Append("ms-eid-membership", types.YChild{"MsEidMembership", nil})
    for i := range instances.MsEidMembership {
        instances.EntityData.Children.Append(types.GetSegmentPath(instances.MsEidMembership[i]), types.YChild{"MsEidMembership", instances.MsEidMembership[i]})
    }
    instances.EntityData.Children.Append("etr-eid-membership", types.YChild{"EtrEidMembership", nil})
    for i := range instances.EtrEidMembership {
        instances.EntityData.Children.Append(types.GetSegmentPath(instances.EtrEidMembership[i]), types.YChild{"EtrEidMembership", instances.EtrEidMembership[i]})
    }
    instances.EntityData.Leafs = types.NewOrderedMap()
    instances.EntityData.Leafs.Append("iid", types.YLeaf{"Iid", instances.Iid})
    instances.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", instances.VrfName})
    instances.EntityData.Leafs.Append("is-rloc-probing", types.YLeaf{"IsRlocProbing", instances.IsRlocProbing})

    instances.EntityData.YListKeys = []string {"Iid"}

    return &(instances.EntityData)
}

// LispState_LispRouters_Instances_Af
// List of Address-Families enabled in this LISP instance
type LispState_LispRouters_Instances_Af struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Instance-specific service type. The type is
    // LispIaftypeType.
    Iaftype interface{}

    // ID of Vlan or Bridge-Domain that is mapped to the given L2 LISP Instance
    // ID. The type is interface{} with range: 0..4294967295.
    L2DomainId interface{}

    // Current size of EID-to-RLOC map-cache on this device. The type is
    // interface{} with range: 0..4294967295.
    MapCacheSize interface{}

    // Maximum permissible number of entries in EID-to-RLOC map-cache on this
    // device. The type is interface{} with range: 0..4294967295.
    MapCacheLimit interface{}

    // TTL of the EID-to-RLOC map record provided by the local device in mapping
    // records. The type is interface{} with range: 0..4294967295. Units are
    // minutes.
    EtrMapCacheTtl interface{}

    // Number of EID prefix registrations that were accepted as a result of the
    // 'accept-more-specific' configuration. The type is interface{} with range:
    // 0..4294967295.
    RegistrationMoreSpecific interface{}

    // The warning threshold for the 'accept-more-specific' registration count on
    // the Map-Server. The type is interface{} with range: 0..4294967295.
    RegistrationMoreSpecificWarningThreshold interface{}

    // Maximum number of registrations on the Map-Server which could be accepted
    // due to configuration of 'accept-more-specific'. The type is interface{}
    // with range: 0..4294967295.
    RegistrationMoreSpecificLimit interface{}

    // The map-cache utilization warning threshold on the xTR. The type is
    // interface{} with range: 0..4294967295.
    MapCacheThreshold interface{}

    // Total number of map requests received by this device for any EID-Prefix of
    // the given address family and Instance ID. The type is interface{} with
    // range: 0..18446744073709551615.
    MapRequestsIn interface{}

    // Total number of map requests sent by this device for any EID-Prefix of the
    // given address family and Instance ID. The type is interface{} with range:
    // 0..18446744073709551615.
    MapRequestsOut interface{}

    // Total number of encapsulated Map-Requests received by this device for any
    // EID-Prefix of the given address family and Instance ID. The type is
    // interface{} with range: 0..18446744073709551615.
    EncapsulatedMapRequestsIn interface{}

    // Total number of encapsulated Map-Requests sent by this device for any
    // EID-Prefix of the given address family and Instance ID. The type is
    // interface{} with range: 0..18446744073709551615.
    EncapsulatedMapRequestsOut interface{}

    // Total number of RLOC probe Map-Requests received by this device for any
    // EID-Prefix of the given address family and Instance ID. The type is
    // interface{} with range: 0..18446744073709551615.
    RlocProbeMapRequestsIn interface{}

    // Total number of RLOC probe Map-Requests sent by this device for any
    // EID-Prefix of the given address family and Instance ID. The type is
    // interface{} with range: 0..18446744073709551615.
    RlocProbeMapRequestsOut interface{}

    // Total number of Map-Requests for any EID-Prefix of the given address family
    // and Instance ID which were not sent out by this device because they timed
    // out while on the transmit queue. The type is interface{} with range:
    // 0..18446744073709551615.
    MapRequestsExpiredOnQueue interface{}

    // Total number of Map-Requests attempted by this device for any EID-Prefix of
    // the given address family and Instance ID without reciving a Map-Reply. The
    // type is interface{} with range: 0..18446744073709551615.
    MapRequestsNoReply interface{}

    // Total number of Map-Request messages for any EID-Prefix of the given
    // address family and Instance ID dropped by this device due to configuration
    // 'map-resolver allowed-locator'. The type is interface{} with range:
    // 0..18446744073709551615.
    MapRequestsFromDisallowedLocators interface{}

    // Total number of SMR Map-Requests received by this device for any EID-Prefix
    // of the given address family and Instance ID. The type is interface{} with
    // range: 0..18446744073709551615.
    SmrMapRequestsIn interface{}

    // Total number of SMR Map-Requests sent by this device for any EID-Prefix of
    // the given address family and Instance ID. The type is interface{} with
    // range: 0..18446744073709551615.
    SmrMapRequestsOut interface{}

    // Total number of ITR's Map-Request messages for any EID-Prefix of the given
    // address family and Instance ID dropped by the DDT Map-Resolver. The type is
    // interface{} with range: 0..18446744073709551615.
    DdtItrMapRequestsDropped interface{}

    // Total number of ITR's Map-Request messages for any EID-Prefix of the given
    // address family and Instance ID dropped by the DDT Map-Resolver due to nonce
    // conflict. The type is interface{} with range: 0..18446744073709551615.
    DdtItrMapRequestsNonceCollision interface{}

    // Total number of ITR's Map-Request messages for any EID-Prefix of the given
    // address family and Instance ID dropped by the DDT Map-Resolver due bad
    // nonce. The type is interface{} with range: 0..18446744073709551615.
    DdtItrMapRequestsBadXtrNonce interface{}

    // Total number of Map-Requests for any EID-Prefix of the given address family
    // and Instance ID forwarded by this device to Map-Resolver on ALT. The type
    // is interface{} with range: 0..18446744073709551615.
    MrMapRequestsForwarded interface{}

    // Total number of Map-Requests for any EID-Prefix of the given address family
    // and Instance ID forwarded by this device to ETR. The type is interface{}
    // with range: 0..18446744073709551615.
    MsMapRequestsForwarded interface{}

    // Total number of Map-Requests for any EID-Prefix of the given address family
    // and Instance ID forwarded by this device to ALT. The type is interface{}
    // with range: 0..18446744073709551615.
    ToAltMapRequestsOut interface{}

    // Total number of Map-Reply records received by this device for any
    // EID-Prefix of the given address family and Instance ID. The type is
    // interface{} with range: 0..18446744073709551615.
    MapReplyRecordsIn interface{}

    // Total number of Map-Reply records sent by this device for any EID-Prefix of
    // the given address family and Instance ID. The type is interface{} with
    // range: 0..18446744073709551615.
    MapReplyRecordsOut interface{}

    // Total number of authoritative Map-Reply records received by this device for
    // any EID-Prefix of the given address family and Instance ID. The type is
    // interface{} with range: 0..18446744073709551615.
    AuthoritativeRecordsIn interface{}

    // Total number of authoritative Map-Reply records sent by this device for any
    // EID-Prefix of the given address family and Instance ID. The type is
    // interface{} with range: 0..18446744073709551615.
    AuthoritativeRecordsOut interface{}

    // Total number of non authoritative Map-Reply records received by this device
    // for any EID-Prefix of the given address family and Instance ID. The type is
    // interface{} with range: 0..18446744073709551615.
    NonAuthoritativeRecordsIn interface{}

    // Total number of non authoritative Map-Reply records sent by this device for
    // any EID-Prefix of the given address family and Instance ID. The type is
    // interface{} with range: 0..18446744073709551615.
    NonAuthoritativeRecordsOut interface{}

    // Total number of negative Map-Reply records received by this device for any
    // EID-Prefix of the given address family and Instance ID. The type is
    // interface{} with range: 0..18446744073709551615.
    NegativeRecordsIn interface{}

    // Total number of negative Map-Reply records sent by this device for any
    // EID-Prefix of the given address family and Instance ID. The type is
    // interface{} with range: 0..18446744073709551615.
    NegativeRecordsOut interface{}

    // Total number of RLOC probe Map-Replies for any EID-Prefix of the given
    // address family and Instance ID received by this device. The type is
    // interface{} with range: 0..18446744073709551615.
    RlocProbeRecordsIn interface{}

    // Total number of RLOC probe Map-Replies for any EID-Prefix of the given
    // address family and Instance ID sent by this device. The type is interface{}
    // with range: 0..18446744073709551615.
    RlocProbeRecordsOut interface{}

    // Total number of MS proxy Map-Replies for any EID-Prefix of the given
    // address family and Instance ID sent by this device. The type is interface{}
    // with range: 0..18446744073709551615.
    MsProxyReplyRecordsOut interface{}

    // Total number of WLC Subscribe messages received by this device for the
    // given address family and Instance ID. The type is interface{} with range:
    // 0..18446744073709551615.
    WlcSubscribeIn interface{}

    // Total number of WLC Subscribe messages sent by this device for the given
    // address family and Instance ID. The type is interface{} with range:
    // 0..18446744073709551615.
    WlcSubscribeOut interface{}

    // Total number of received WLC Subscribe messages for the given address
    // family and Instance ID with incorrect formatting. The type is interface{}
    // with range: 0..18446744073709551615.
    WlcSubscribeInFailure interface{}

    // Total number of WLC Subscribe messages for the given address family and
    // Instance ID which were not sent due to internal errors. The type is
    // interface{} with range: 0..18446744073709551615.
    WlcSubscribeOutFailure interface{}

    // Total number of WLC Unsubscribe messages received by this device for the
    // given address family and Instance ID. The type is interface{} with range:
    // 0..18446744073709551615.
    WlcUnsubscribeIn interface{}

    // Total number of WLC Unsubscribe messages sent by this device for the given
    // address family and Instance ID. The type is interface{} with range:
    // 0..18446744073709551615.
    WlcUnsubscribeOut interface{}

    // Total number of received WLC Unsubscribe messages for the given address
    // family and Instance ID with incorrect formatting. The type is interface{}
    // with range: 0..18446744073709551615.
    WlcUnsubscribeInFailure interface{}

    // Total number of WLC Unsubscribe messages for the given address family and
    // Instance ID which were not sent due to internal errors. The type is
    // interface{} with range: 0..18446744073709551615.
    WlcUnsubscribeOutFailure interface{}

    // Total number of Map-Register records for any EID-Prefix of the given
    // address family and Instance ID received by this device. The type is
    // interface{} with range: 0..18446744073709551615.
    MapRegisterRecordsIn interface{}

    // Total number of Map-Registers records sent by this device for any
    // EID-Prefix of the given address family and Instance ID. The type is
    // interface{} with range: 0..18446744073709551615.
    MapRegisterRecordsOut interface{}

    // Total number of Map-Register messages dropped due to Map-Server
    // functionality not enabled for the given address family and Instance ID. The
    // type is interface{} with range: 0..18446744073709551615.
    MapRegistersMsDisabled interface{}

    // Total number of Map-Register messages for any EID-Prefix of the given
    // address family and Instance ID dropped due to authentication failure. The
    // type is interface{} with range: 0..18446744073709551615.
    MapRegistersAuthFailed interface{}

    // Total number of Map-Register messages received from disallowed locators.
    // The type is interface{} with range: 0..18446744073709551615.
    MapRegistersFromDisallowedLocators interface{}

    // Total number of WLC Map-Register messages received by this device for any
    // EID-Prefix of the given address family and Instance ID. The type is
    // interface{} with range: 0..18446744073709551615.
    WlcMapRegisterRecordsIn interface{}

    // Total number of WLC Map-Register messages sent by this device for any
    // EID-Prefix of the given address family and Instance ID. The type is
    // interface{} with range: 0..18446744073709551615.
    WlcMapRegisterRecordsOut interface{}

    // Total number of WLC Map-Register messages received by this device for AP
    // join in the given address family and Instance ID. The type is interface{}
    // with range: 0..18446744073709551615.
    WlcMapRegisterRecordsInAp interface{}

    // Total number of WLC Map-Register messages sent by this device for AP join
    // in the given address family and Instance ID. The type is interface{} with
    // range: 0..18446744073709551615.
    WlcMapRegisterRecordsOutAp interface{}

    // Total number of WLC Map-Register messages received by this device for
    // wireless client join in the given address family and Instance ID. The type
    // is interface{} with range: 0..18446744073709551615.
    WlcMapRegisterRecordsInClient interface{}

    // Total number of WLC Map-Register messages sent by this device for wireless
    // client join in the given address family and Instance ID. The type is
    // interface{} with range: 0..18446744073709551615.
    WlcMapRegisterRecordsOutClient interface{}

    // Total number of WLC Map-Register messages received in the given address
    // family and Instance ID and discarded due to parsing error. The type is
    // interface{} with range: 0..18446744073709551615.
    WlcMapRegisterRecordsInFailure interface{}

    // Total number of WLC Map-Register messages for any EID-Prefix of the given
    // address family and Instance ID which were not sent because of internal
    // error. The type is interface{} with range: 0..18446744073709551615.
    WlcMapRegisterRecordsOutFailure interface{}

    // Total number of Map-Notify records for any EID-Prefix of the given address
    // family and Instance ID received by this device. The type is interface{}
    // with range: 0..18446744073709551615.
    MapNotifyRecordsIn interface{}

    // Total number of Map-Notify records for any EID-Prefix of the given address
    // family and Instance ID sent by this device. The type is interface{} with
    // range: 0..18446744073709551615.
    MapNotifyRecordsOut interface{}

    // Total number of Map-Notify messages for any EID-Prefix of the given address
    // family and Instance ID dropped due to authentication failure. The type is
    // interface{} with range: 0..18446744073709551615.
    MapNotifyAuthFailed interface{}

    // Total number of WLC Map-Notify records for any EID-Prefix of the given
    // address family and Instance ID received by this device. The type is
    // interface{} with range: 0..18446744073709551615.
    WlcMapNotifyRecordsIn interface{}

    // Total number of WLC Map-Notify records for any EID-Prefix of the given
    // address family and Instance ID sent by this device. The type is interface{}
    // with range: 0..18446744073709551615.
    WlcMapNotifyRecordsOut interface{}

    // Total number of WLC Map-Notify records for any EID-Prefix of the given
    // address family and Instance ID received by this device for AP join. The
    // type is interface{} with range: 0..18446744073709551615.
    WlcMapNotifyRecordsInAp interface{}

    // Total number of WLC Map-Notify records for any EID-Prefix of the given
    // address family and Instance ID sent by this device for AP join. The type is
    // interface{} with range: 0..18446744073709551615.
    WlcMapNotifyRecordsOutAp interface{}

    // Total number of WLC Map-Notify records for any EID-Prefix of the given
    // address family and Instance ID received by this device for wireless client
    // join. The type is interface{} with range: 0..18446744073709551615.
    WlcMapNotifyRecordsInClient interface{}

    // Total number of WLC Map-Notify records for any EID-Prefix of the given
    // address family and Instance ID sent by this device for wireless client
    // join. The type is interface{} with range: 0..18446744073709551615.
    WlcMapNotifyRecordsOutClient interface{}

    // Total number of WLC Map-Notify messages received in the given address
    // family and Instance ID and discarded due to parsing error. The type is
    // interface{} with range: 0..18446744073709551615.
    WlcMapNotifyRecordsInFailure interface{}

    // Total number of WLC Map-Notify messages for any EID-Prefix of the given
    // address family and Instance ID which were not sent because of internal
    // error. The type is interface{} with range: 0..18446744073709551615.
    WlcMapNotifyRecordsOutFailure interface{}

    // Total number of mapping records received by this device for any EID-Prefix
    // of the given address family and Instance ID with TTL exceeding 7 days. The
    // type is interface{} with range: 0..18446744073709551615.
    MappingRecordTtlAlerts interface{}

    // Total number of remote EID map-cache entries created by this device for any
    // EID-Prefix of the given address family and Instance ID. The type is
    // interface{} with range: 0..18446744073709551615.
    RemoteEidEntriesCreated interface{}

    // Total number of remote EID map-cache entries deleted by this device for any
    // EID-Prefix of the given address family and Instance ID. The type is
    // interface{} with range: 0..18446744073709551615.
    RemoteEidEntriesDeleted interface{}

    // Total number of remote EID map-cache entries for any EID-Prefix of the
    // given address family and Instance ID recreated by this device after NSF
    // switchover. The type is interface{} with range: 0..18446744073709551615.
    RemoteEidNsfReplayEntriesCreated interface{}

    // Total number of forwarding plane data signals processed by this device for
    // any EID-Prefix of the given address family and Instance ID. The type is
    // interface{} with range: 0..18446744073709551615.
    ForwardingDataSignalsProcessed interface{}

    // Total number of forwarding plane data signals for any EID-Prefix of the
    // given address family and Instance ID dropped by this device. The type is
    // interface{} with range: 0..18446744073709551615.
    ForwardingDataSignalsDropped interface{}

    // Total number of forwarding plane reachability reports for any EID-Prefix of
    // the given address family and Instance ID processed by this device. The type
    // is interface{} with range: 0..18446744073709551615.
    ForwardingReachabilityReportsProcessed interface{}

    // Total number of forwarding plane reachability reports for any EID-Prefix of
    // the given address family and Instance ID dropped by this device. The type
    // is interface{} with range: 0..18446744073709551615.
    ForwardingReachabilityReportsDropped interface{}

    // Indicates whether the ETR accepts piggybacked mapping records received in a
    // Map-Request. The type is bool.
    IsEtrAcceptMapping interface{}

    // Indicates if ETR will try to verify accepted piggybacked mapping records
    // received in a Map-Request. The type is bool.
    IsEtrAcceptMappingVerify interface{}

    // LISP device role for this service.
    Role LispState_LispRouters_Instances_Af_Role

    // Map-cache for this service instance. The type is slice of
    // LispState_LispRouters_Instances_Af_MapCache.
    MapCache []*LispState_LispRouters_Instances_Af_MapCache

    // ETR's database of local EID prefixes. The type is slice of
    // LispState_LispRouters_Instances_Af_LocalDbase.
    LocalDbase []*LispState_LispRouters_Instances_Af_LocalDbase

    // Map-Server database of registered EID Prefixes. The type is slice of
    // LispState_LispRouters_Instances_Af_MsRegistrations.
    MsRegistrations []*LispState_LispRouters_Instances_Af_MsRegistrations

    // List of Map-Servers to which the ETR should register. The type is slice of
    // LispState_LispRouters_Instances_Af_MapServers.
    MapServers []*LispState_LispRouters_Instances_Af_MapServers

    // List of Map-Resolvers where [P]ITR should send its Map-Requests. The type
    // is slice of LispState_LispRouters_Instances_Af_MapResolvers.
    MapResolvers []*LispState_LispRouters_Instances_Af_MapResolvers

    // List of all Proxy ETRs that this device is configured to use. The type is
    // slice of LispState_LispRouters_Instances_Af_ProxyEtrs.
    ProxyEtrs []*LispState_LispRouters_Instances_Af_ProxyEtrs
}

func (af *LispState_LispRouters_Instances_Af) GetEntityData() *types.CommonEntityData {
    af.EntityData.YFilter = af.YFilter
    af.EntityData.YangName = "af"
    af.EntityData.BundleName = "cisco_ios_xe"
    af.EntityData.ParentYangName = "instances"
    af.EntityData.SegmentPath = "af" + types.AddKeyToken(af.Iaftype, "iaftype")
    af.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    af.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    af.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    af.EntityData.Children = types.NewOrderedMap()
    af.EntityData.Children.Append("role", types.YChild{"Role", &af.Role})
    af.EntityData.Children.Append("map-cache", types.YChild{"MapCache", nil})
    for i := range af.MapCache {
        af.EntityData.Children.Append(types.GetSegmentPath(af.MapCache[i]), types.YChild{"MapCache", af.MapCache[i]})
    }
    af.EntityData.Children.Append("local-dbase", types.YChild{"LocalDbase", nil})
    for i := range af.LocalDbase {
        af.EntityData.Children.Append(types.GetSegmentPath(af.LocalDbase[i]), types.YChild{"LocalDbase", af.LocalDbase[i]})
    }
    af.EntityData.Children.Append("ms-registrations", types.YChild{"MsRegistrations", nil})
    for i := range af.MsRegistrations {
        af.EntityData.Children.Append(types.GetSegmentPath(af.MsRegistrations[i]), types.YChild{"MsRegistrations", af.MsRegistrations[i]})
    }
    af.EntityData.Children.Append("map-servers", types.YChild{"MapServers", nil})
    for i := range af.MapServers {
        af.EntityData.Children.Append(types.GetSegmentPath(af.MapServers[i]), types.YChild{"MapServers", af.MapServers[i]})
    }
    af.EntityData.Children.Append("map-resolvers", types.YChild{"MapResolvers", nil})
    for i := range af.MapResolvers {
        af.EntityData.Children.Append(types.GetSegmentPath(af.MapResolvers[i]), types.YChild{"MapResolvers", af.MapResolvers[i]})
    }
    af.EntityData.Children.Append("proxy-etrs", types.YChild{"ProxyEtrs", nil})
    for i := range af.ProxyEtrs {
        af.EntityData.Children.Append(types.GetSegmentPath(af.ProxyEtrs[i]), types.YChild{"ProxyEtrs", af.ProxyEtrs[i]})
    }
    af.EntityData.Leafs = types.NewOrderedMap()
    af.EntityData.Leafs.Append("iaftype", types.YLeaf{"Iaftype", af.Iaftype})
    af.EntityData.Leafs.Append("l2-domain-id", types.YLeaf{"L2DomainId", af.L2DomainId})
    af.EntityData.Leafs.Append("map-cache-size", types.YLeaf{"MapCacheSize", af.MapCacheSize})
    af.EntityData.Leafs.Append("map-cache-limit", types.YLeaf{"MapCacheLimit", af.MapCacheLimit})
    af.EntityData.Leafs.Append("etr-map-cache-ttl", types.YLeaf{"EtrMapCacheTtl", af.EtrMapCacheTtl})
    af.EntityData.Leafs.Append("registration-more-specific", types.YLeaf{"RegistrationMoreSpecific", af.RegistrationMoreSpecific})
    af.EntityData.Leafs.Append("registration-more-specific-warning-threshold", types.YLeaf{"RegistrationMoreSpecificWarningThreshold", af.RegistrationMoreSpecificWarningThreshold})
    af.EntityData.Leafs.Append("registration-more-specific-limit", types.YLeaf{"RegistrationMoreSpecificLimit", af.RegistrationMoreSpecificLimit})
    af.EntityData.Leafs.Append("map-cache-threshold", types.YLeaf{"MapCacheThreshold", af.MapCacheThreshold})
    af.EntityData.Leafs.Append("map-requests-in", types.YLeaf{"MapRequestsIn", af.MapRequestsIn})
    af.EntityData.Leafs.Append("map-requests-out", types.YLeaf{"MapRequestsOut", af.MapRequestsOut})
    af.EntityData.Leafs.Append("encapsulated-map-requests-in", types.YLeaf{"EncapsulatedMapRequestsIn", af.EncapsulatedMapRequestsIn})
    af.EntityData.Leafs.Append("encapsulated-map-requests-out", types.YLeaf{"EncapsulatedMapRequestsOut", af.EncapsulatedMapRequestsOut})
    af.EntityData.Leafs.Append("rloc-probe-map-requests-in", types.YLeaf{"RlocProbeMapRequestsIn", af.RlocProbeMapRequestsIn})
    af.EntityData.Leafs.Append("rloc-probe-map-requests-out", types.YLeaf{"RlocProbeMapRequestsOut", af.RlocProbeMapRequestsOut})
    af.EntityData.Leafs.Append("map-requests-expired-on-queue", types.YLeaf{"MapRequestsExpiredOnQueue", af.MapRequestsExpiredOnQueue})
    af.EntityData.Leafs.Append("map-requests-no-reply", types.YLeaf{"MapRequestsNoReply", af.MapRequestsNoReply})
    af.EntityData.Leafs.Append("map-requests-from-disallowed-locators", types.YLeaf{"MapRequestsFromDisallowedLocators", af.MapRequestsFromDisallowedLocators})
    af.EntityData.Leafs.Append("smr-map-requests-in", types.YLeaf{"SmrMapRequestsIn", af.SmrMapRequestsIn})
    af.EntityData.Leafs.Append("smr-map-requests-out", types.YLeaf{"SmrMapRequestsOut", af.SmrMapRequestsOut})
    af.EntityData.Leafs.Append("ddt-itr-map-requests-dropped", types.YLeaf{"DdtItrMapRequestsDropped", af.DdtItrMapRequestsDropped})
    af.EntityData.Leafs.Append("ddt-itr-map-requests-nonce-collision", types.YLeaf{"DdtItrMapRequestsNonceCollision", af.DdtItrMapRequestsNonceCollision})
    af.EntityData.Leafs.Append("ddt-itr-map-requests-bad-xtr-nonce", types.YLeaf{"DdtItrMapRequestsBadXtrNonce", af.DdtItrMapRequestsBadXtrNonce})
    af.EntityData.Leafs.Append("mr-map-requests-forwarded", types.YLeaf{"MrMapRequestsForwarded", af.MrMapRequestsForwarded})
    af.EntityData.Leafs.Append("ms-map-requests-forwarded", types.YLeaf{"MsMapRequestsForwarded", af.MsMapRequestsForwarded})
    af.EntityData.Leafs.Append("to-alt-map-requests-out", types.YLeaf{"ToAltMapRequestsOut", af.ToAltMapRequestsOut})
    af.EntityData.Leafs.Append("map-reply-records-in", types.YLeaf{"MapReplyRecordsIn", af.MapReplyRecordsIn})
    af.EntityData.Leafs.Append("map-reply-records-out", types.YLeaf{"MapReplyRecordsOut", af.MapReplyRecordsOut})
    af.EntityData.Leafs.Append("authoritative-records-in", types.YLeaf{"AuthoritativeRecordsIn", af.AuthoritativeRecordsIn})
    af.EntityData.Leafs.Append("authoritative-records-out", types.YLeaf{"AuthoritativeRecordsOut", af.AuthoritativeRecordsOut})
    af.EntityData.Leafs.Append("non-authoritative-records-in", types.YLeaf{"NonAuthoritativeRecordsIn", af.NonAuthoritativeRecordsIn})
    af.EntityData.Leafs.Append("non-authoritative-records-out", types.YLeaf{"NonAuthoritativeRecordsOut", af.NonAuthoritativeRecordsOut})
    af.EntityData.Leafs.Append("negative-records-in", types.YLeaf{"NegativeRecordsIn", af.NegativeRecordsIn})
    af.EntityData.Leafs.Append("negative-records-out", types.YLeaf{"NegativeRecordsOut", af.NegativeRecordsOut})
    af.EntityData.Leafs.Append("rloc-probe-records-in", types.YLeaf{"RlocProbeRecordsIn", af.RlocProbeRecordsIn})
    af.EntityData.Leafs.Append("rloc-probe-records-out", types.YLeaf{"RlocProbeRecordsOut", af.RlocProbeRecordsOut})
    af.EntityData.Leafs.Append("ms-proxy-reply-records-out", types.YLeaf{"MsProxyReplyRecordsOut", af.MsProxyReplyRecordsOut})
    af.EntityData.Leafs.Append("wlc-subscribe-in", types.YLeaf{"WlcSubscribeIn", af.WlcSubscribeIn})
    af.EntityData.Leafs.Append("wlc-subscribe-out", types.YLeaf{"WlcSubscribeOut", af.WlcSubscribeOut})
    af.EntityData.Leafs.Append("wlc-subscribe-in-failure", types.YLeaf{"WlcSubscribeInFailure", af.WlcSubscribeInFailure})
    af.EntityData.Leafs.Append("wlc-subscribe-out-failure", types.YLeaf{"WlcSubscribeOutFailure", af.WlcSubscribeOutFailure})
    af.EntityData.Leafs.Append("wlc-unsubscribe-in", types.YLeaf{"WlcUnsubscribeIn", af.WlcUnsubscribeIn})
    af.EntityData.Leafs.Append("wlc-unsubscribe-out", types.YLeaf{"WlcUnsubscribeOut", af.WlcUnsubscribeOut})
    af.EntityData.Leafs.Append("wlc-unsubscribe-in-failure", types.YLeaf{"WlcUnsubscribeInFailure", af.WlcUnsubscribeInFailure})
    af.EntityData.Leafs.Append("wlc-unsubscribe-out-failure", types.YLeaf{"WlcUnsubscribeOutFailure", af.WlcUnsubscribeOutFailure})
    af.EntityData.Leafs.Append("map-register-records-in", types.YLeaf{"MapRegisterRecordsIn", af.MapRegisterRecordsIn})
    af.EntityData.Leafs.Append("map-register-records-out", types.YLeaf{"MapRegisterRecordsOut", af.MapRegisterRecordsOut})
    af.EntityData.Leafs.Append("map-registers-ms-disabled", types.YLeaf{"MapRegistersMsDisabled", af.MapRegistersMsDisabled})
    af.EntityData.Leafs.Append("map-registers-auth-failed", types.YLeaf{"MapRegistersAuthFailed", af.MapRegistersAuthFailed})
    af.EntityData.Leafs.Append("map-registers-from-disallowed-locators", types.YLeaf{"MapRegistersFromDisallowedLocators", af.MapRegistersFromDisallowedLocators})
    af.EntityData.Leafs.Append("wlc-map-register-records-in", types.YLeaf{"WlcMapRegisterRecordsIn", af.WlcMapRegisterRecordsIn})
    af.EntityData.Leafs.Append("wlc-map-register-records-out", types.YLeaf{"WlcMapRegisterRecordsOut", af.WlcMapRegisterRecordsOut})
    af.EntityData.Leafs.Append("wlc-map-register-records-in-ap", types.YLeaf{"WlcMapRegisterRecordsInAp", af.WlcMapRegisterRecordsInAp})
    af.EntityData.Leafs.Append("wlc-map-register-records-out-ap", types.YLeaf{"WlcMapRegisterRecordsOutAp", af.WlcMapRegisterRecordsOutAp})
    af.EntityData.Leafs.Append("wlc-map-register-records-in-client", types.YLeaf{"WlcMapRegisterRecordsInClient", af.WlcMapRegisterRecordsInClient})
    af.EntityData.Leafs.Append("wlc-map-register-records-out-client", types.YLeaf{"WlcMapRegisterRecordsOutClient", af.WlcMapRegisterRecordsOutClient})
    af.EntityData.Leafs.Append("wlc-map-register-records-in-failure", types.YLeaf{"WlcMapRegisterRecordsInFailure", af.WlcMapRegisterRecordsInFailure})
    af.EntityData.Leafs.Append("wlc-map-register-records-out-failure", types.YLeaf{"WlcMapRegisterRecordsOutFailure", af.WlcMapRegisterRecordsOutFailure})
    af.EntityData.Leafs.Append("map-notify-records-in", types.YLeaf{"MapNotifyRecordsIn", af.MapNotifyRecordsIn})
    af.EntityData.Leafs.Append("map-notify-records-out", types.YLeaf{"MapNotifyRecordsOut", af.MapNotifyRecordsOut})
    af.EntityData.Leafs.Append("map-notify-auth-failed", types.YLeaf{"MapNotifyAuthFailed", af.MapNotifyAuthFailed})
    af.EntityData.Leafs.Append("wlc-map-notify-records-in", types.YLeaf{"WlcMapNotifyRecordsIn", af.WlcMapNotifyRecordsIn})
    af.EntityData.Leafs.Append("wlc-map-notify-records-out", types.YLeaf{"WlcMapNotifyRecordsOut", af.WlcMapNotifyRecordsOut})
    af.EntityData.Leafs.Append("wlc-map-notify-records-in-ap", types.YLeaf{"WlcMapNotifyRecordsInAp", af.WlcMapNotifyRecordsInAp})
    af.EntityData.Leafs.Append("wlc-map-notify-records-out-ap", types.YLeaf{"WlcMapNotifyRecordsOutAp", af.WlcMapNotifyRecordsOutAp})
    af.EntityData.Leafs.Append("wlc-map-notify-records-in-client", types.YLeaf{"WlcMapNotifyRecordsInClient", af.WlcMapNotifyRecordsInClient})
    af.EntityData.Leafs.Append("wlc-map-notify-records-out-client", types.YLeaf{"WlcMapNotifyRecordsOutClient", af.WlcMapNotifyRecordsOutClient})
    af.EntityData.Leafs.Append("wlc-map-notify-records-in-failure", types.YLeaf{"WlcMapNotifyRecordsInFailure", af.WlcMapNotifyRecordsInFailure})
    af.EntityData.Leafs.Append("wlc-map-notify-records-out-failure", types.YLeaf{"WlcMapNotifyRecordsOutFailure", af.WlcMapNotifyRecordsOutFailure})
    af.EntityData.Leafs.Append("mapping-record-ttl-alerts", types.YLeaf{"MappingRecordTtlAlerts", af.MappingRecordTtlAlerts})
    af.EntityData.Leafs.Append("remote-eid-entries-created", types.YLeaf{"RemoteEidEntriesCreated", af.RemoteEidEntriesCreated})
    af.EntityData.Leafs.Append("remote-eid-entries-deleted", types.YLeaf{"RemoteEidEntriesDeleted", af.RemoteEidEntriesDeleted})
    af.EntityData.Leafs.Append("remote-eid-nsf-replay-entries-created", types.YLeaf{"RemoteEidNsfReplayEntriesCreated", af.RemoteEidNsfReplayEntriesCreated})
    af.EntityData.Leafs.Append("forwarding-data-signals-processed", types.YLeaf{"ForwardingDataSignalsProcessed", af.ForwardingDataSignalsProcessed})
    af.EntityData.Leafs.Append("forwarding-data-signals-dropped", types.YLeaf{"ForwardingDataSignalsDropped", af.ForwardingDataSignalsDropped})
    af.EntityData.Leafs.Append("forwarding-reachability-reports-processed", types.YLeaf{"ForwardingReachabilityReportsProcessed", af.ForwardingReachabilityReportsProcessed})
    af.EntityData.Leafs.Append("forwarding-reachability-reports-dropped", types.YLeaf{"ForwardingReachabilityReportsDropped", af.ForwardingReachabilityReportsDropped})
    af.EntityData.Leafs.Append("is-etr-accept-mapping", types.YLeaf{"IsEtrAcceptMapping", af.IsEtrAcceptMapping})
    af.EntityData.Leafs.Append("is-etr-accept-mapping-verify", types.YLeaf{"IsEtrAcceptMappingVerify", af.IsEtrAcceptMappingVerify})

    af.EntityData.YListKeys = []string {"Iaftype"}

    return &(af.EntityData)
}

// LispState_LispRouters_Instances_Af_Role
// LISP device role for this service
type LispState_LispRouters_Instances_Af_Role struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LISP Map-Server. The type is bool.
    IsMs interface{}

    // LISP Map-Resolver. The type is bool.
    IsMr interface{}

    // LISP ITR. The type is bool.
    IsItr interface{}

    // LISP ETR. The type is bool.
    IsEtr interface{}

    // LISP Proxy-ITR. The type is bool.
    IsPitr interface{}

    // LISP Proxy-ETR. The type is bool.
    IsPetr interface{}
}

func (role *LispState_LispRouters_Instances_Af_Role) GetEntityData() *types.CommonEntityData {
    role.EntityData.YFilter = role.YFilter
    role.EntityData.YangName = "role"
    role.EntityData.BundleName = "cisco_ios_xe"
    role.EntityData.ParentYangName = "af"
    role.EntityData.SegmentPath = "role"
    role.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    role.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    role.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    role.EntityData.Children = types.NewOrderedMap()
    role.EntityData.Leafs = types.NewOrderedMap()
    role.EntityData.Leafs.Append("is-ms", types.YLeaf{"IsMs", role.IsMs})
    role.EntityData.Leafs.Append("is-mr", types.YLeaf{"IsMr", role.IsMr})
    role.EntityData.Leafs.Append("is-itr", types.YLeaf{"IsItr", role.IsItr})
    role.EntityData.Leafs.Append("is-etr", types.YLeaf{"IsEtr", role.IsEtr})
    role.EntityData.Leafs.Append("is-pitr", types.YLeaf{"IsPitr", role.IsPitr})
    role.EntityData.Leafs.Append("is-petr", types.YLeaf{"IsPetr", role.IsPetr})

    role.EntityData.YListKeys = []string {}

    return &(role.EntityData)
}

// LispState_LispRouters_Instances_Af_MapCache
// Map-cache for this service instance
type LispState_LispRouters_Instances_Af_MapCache struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. LISP Address-Family of the prefix. The type is
    // LispAddressFamilyType.
    Afi interface{}

    // This attribute is a key. LISP prefix. Format is defined by the AF. The type
    // is string.
    Prefix interface{}

    // Time that this entry was created. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    UpTime interface{}

    // Last time that the RLOC information or the entry state were modified. The
    // type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    LastModifiedTime interface{}

    // Last time a mapping record for this entry was received, not valid if the
    // entry is static. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    LastUpdateTime interface{}

    // Mapping validity period. The type is interface{} with range: 0..4294967295.
    // Units are milliseconds.
    Ttl interface{}

    // Indication if the mapping came from an authoritative source. The type is
    // bool.
    IsAuthoritative interface{}

    // Indication if the mapping is static (i.e. configured). The type is bool.
    IsStatic interface{}

    // Indication if the mapping is negative (i.e. provides no locators for LISP
    // encapsulation). The type is bool.
    IsNegative interface{}

    // Forwarding action in case of negative entry. The type is
    // LispMapReplyActionType.
    NmrAction interface{}

    // Time when this entry will expire if not refreshed; for entries which do not
    // have an expiration time this time will be less than the entry creation
    // time. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    ExpiryTime interface{}

    // Number of packets of the transit traffic which were encapsulated because
    // they matched this map-cache entry. The type is interface{} with range:
    // 0..18446744073709551615.
    EncapsulatedPackets interface{}

    // Number of bytes of the transit traffic which were encapsulated because they
    // matched this map-cache entry. The type is interface{} with range:
    // 0..18446744073709551615.
    EncapsulatedOctets interface{}

    // Indication if the mapping is active, that is there was transit traffic
    // matching this map-cache entry seen in approximately the last minute. The
    // type is bool.
    IsActive interface{}

    // List of locators for positive mapping. The type is slice of
    // LispState_LispRouters_Instances_Af_MapCache_MapCacheRloc.
    MapCacheRloc []*LispState_LispRouters_Instances_Af_MapCache_MapCacheRloc
}

func (mapCache *LispState_LispRouters_Instances_Af_MapCache) GetEntityData() *types.CommonEntityData {
    mapCache.EntityData.YFilter = mapCache.YFilter
    mapCache.EntityData.YangName = "map-cache"
    mapCache.EntityData.BundleName = "cisco_ios_xe"
    mapCache.EntityData.ParentYangName = "af"
    mapCache.EntityData.SegmentPath = "map-cache" + types.AddKeyToken(mapCache.Afi, "afi") + types.AddKeyToken(mapCache.Prefix, "prefix")
    mapCache.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mapCache.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mapCache.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mapCache.EntityData.Children = types.NewOrderedMap()
    mapCache.EntityData.Children.Append("map-cache-rloc", types.YChild{"MapCacheRloc", nil})
    for i := range mapCache.MapCacheRloc {
        mapCache.EntityData.Children.Append(types.GetSegmentPath(mapCache.MapCacheRloc[i]), types.YChild{"MapCacheRloc", mapCache.MapCacheRloc[i]})
    }
    mapCache.EntityData.Leafs = types.NewOrderedMap()
    mapCache.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", mapCache.Afi})
    mapCache.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", mapCache.Prefix})
    mapCache.EntityData.Leafs.Append("up-time", types.YLeaf{"UpTime", mapCache.UpTime})
    mapCache.EntityData.Leafs.Append("last-modified-time", types.YLeaf{"LastModifiedTime", mapCache.LastModifiedTime})
    mapCache.EntityData.Leafs.Append("last-update-time", types.YLeaf{"LastUpdateTime", mapCache.LastUpdateTime})
    mapCache.EntityData.Leafs.Append("ttl", types.YLeaf{"Ttl", mapCache.Ttl})
    mapCache.EntityData.Leafs.Append("is-authoritative", types.YLeaf{"IsAuthoritative", mapCache.IsAuthoritative})
    mapCache.EntityData.Leafs.Append("is-static", types.YLeaf{"IsStatic", mapCache.IsStatic})
    mapCache.EntityData.Leafs.Append("is-negative", types.YLeaf{"IsNegative", mapCache.IsNegative})
    mapCache.EntityData.Leafs.Append("nmr-action", types.YLeaf{"NmrAction", mapCache.NmrAction})
    mapCache.EntityData.Leafs.Append("expiry-time", types.YLeaf{"ExpiryTime", mapCache.ExpiryTime})
    mapCache.EntityData.Leafs.Append("encapsulated-packets", types.YLeaf{"EncapsulatedPackets", mapCache.EncapsulatedPackets})
    mapCache.EntityData.Leafs.Append("encapsulated-octets", types.YLeaf{"EncapsulatedOctets", mapCache.EncapsulatedOctets})
    mapCache.EntityData.Leafs.Append("is-active", types.YLeaf{"IsActive", mapCache.IsActive})

    mapCache.EntityData.YListKeys = []string {"Afi", "Prefix"}

    return &(mapCache.EntityData)
}

// LispState_LispRouters_Instances_Af_MapCache_MapCacheRloc
// List of locators for positive mapping
type LispState_LispRouters_Instances_Af_MapCache_MapCacheRloc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. LISP Address-Family of the address. The type is
    // LispAddressFamilyType.
    Afi interface{}

    // This attribute is a key. LISP address. Format is defined by the AF. The
    // type is string.
    Address interface{}

    // Up/down state of the locator. The type is LispRlocStateType.
    State interface{}

    // Time when this RLOC entry was created. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    CreationTime interface{}

    // Time when up/down state of the RLOC for this map-cache entry last changed.
    // The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    LastStateChangeTime interface{}

    // Round-trip time of RLOC probe and corresponding reply. The type is
    // interface{} with range: 0..4294967295. Units are milliseconds.
    RlocProbeRtt interface{}

    // Properties of the locator.
    Params LispState_LispRouters_Instances_Af_MapCache_MapCacheRloc_Params
}

func (mapCacheRloc *LispState_LispRouters_Instances_Af_MapCache_MapCacheRloc) GetEntityData() *types.CommonEntityData {
    mapCacheRloc.EntityData.YFilter = mapCacheRloc.YFilter
    mapCacheRloc.EntityData.YangName = "map-cache-rloc"
    mapCacheRloc.EntityData.BundleName = "cisco_ios_xe"
    mapCacheRloc.EntityData.ParentYangName = "map-cache"
    mapCacheRloc.EntityData.SegmentPath = "map-cache-rloc" + types.AddKeyToken(mapCacheRloc.Afi, "afi") + types.AddKeyToken(mapCacheRloc.Address, "address")
    mapCacheRloc.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mapCacheRloc.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mapCacheRloc.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mapCacheRloc.EntityData.Children = types.NewOrderedMap()
    mapCacheRloc.EntityData.Children.Append("params", types.YChild{"Params", &mapCacheRloc.Params})
    mapCacheRloc.EntityData.Leafs = types.NewOrderedMap()
    mapCacheRloc.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", mapCacheRloc.Afi})
    mapCacheRloc.EntityData.Leafs.Append("address", types.YLeaf{"Address", mapCacheRloc.Address})
    mapCacheRloc.EntityData.Leafs.Append("state", types.YLeaf{"State", mapCacheRloc.State})
    mapCacheRloc.EntityData.Leafs.Append("creation-time", types.YLeaf{"CreationTime", mapCacheRloc.CreationTime})
    mapCacheRloc.EntityData.Leafs.Append("last-state-change-time", types.YLeaf{"LastStateChangeTime", mapCacheRloc.LastStateChangeTime})
    mapCacheRloc.EntityData.Leafs.Append("rloc-probe-rtt", types.YLeaf{"RlocProbeRtt", mapCacheRloc.RlocProbeRtt})

    mapCacheRloc.EntityData.YListKeys = []string {"Afi", "Address"}

    return &(mapCacheRloc.EntityData)
}

// LispState_LispRouters_Instances_Af_MapCache_MapCacheRloc_Params
// Properties of the locator
type LispState_LispRouters_Instances_Af_MapCache_MapCacheRloc_Params struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Locator priority. The type is interface{} with range: 0..255.
    Priority interface{}

    // Locator weight. The type is interface{} with range: 0..255.
    Weight interface{}

    // Locator's multicast priority. The type is interface{} with range: 0..255.
    McastPriority interface{}

    // Locator's multicast weight. The type is interface{} with range: 0..255.
    McastWeight interface{}
}

func (params *LispState_LispRouters_Instances_Af_MapCache_MapCacheRloc_Params) GetEntityData() *types.CommonEntityData {
    params.EntityData.YFilter = params.YFilter
    params.EntityData.YangName = "params"
    params.EntityData.BundleName = "cisco_ios_xe"
    params.EntityData.ParentYangName = "map-cache-rloc"
    params.EntityData.SegmentPath = "params"
    params.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    params.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    params.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    params.EntityData.Children = types.NewOrderedMap()
    params.EntityData.Leafs = types.NewOrderedMap()
    params.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", params.Priority})
    params.EntityData.Leafs.Append("weight", types.YLeaf{"Weight", params.Weight})
    params.EntityData.Leafs.Append("mcast-priority", types.YLeaf{"McastPriority", params.McastPriority})
    params.EntityData.Leafs.Append("mcast-weight", types.YLeaf{"McastWeight", params.McastWeight})

    params.EntityData.YListKeys = []string {}

    return &(params.EntityData)
}

// LispState_LispRouters_Instances_Af_LocalDbase
// ETR's database of local EID prefixes
type LispState_LispRouters_Instances_Af_LocalDbase struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. LISP Address-Family of the prefix. The type is
    // LispAddressFamilyType.
    Afi interface{}

    // This attribute is a key. LISP prefix. Format is defined by the AF. The type
    // is string.
    Prefix interface{}

    // The Locator Status Bits for this EID-Prefix. The type is interface{} with
    // range: 0..4294967295.
    Lsb interface{}

    // If the prefix is currently reachable from this device via EID interfaces.
    // The type is bool.
    IsReachable interface{}

    // List of locators. The type is slice of
    // LispState_LispRouters_Instances_Af_LocalDbase_LocalDbaseRloc.
    LocalDbaseRloc []*LispState_LispRouters_Instances_Af_LocalDbase_LocalDbaseRloc
}

func (localDbase *LispState_LispRouters_Instances_Af_LocalDbase) GetEntityData() *types.CommonEntityData {
    localDbase.EntityData.YFilter = localDbase.YFilter
    localDbase.EntityData.YangName = "local-dbase"
    localDbase.EntityData.BundleName = "cisco_ios_xe"
    localDbase.EntityData.ParentYangName = "af"
    localDbase.EntityData.SegmentPath = "local-dbase" + types.AddKeyToken(localDbase.Afi, "afi") + types.AddKeyToken(localDbase.Prefix, "prefix")
    localDbase.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    localDbase.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    localDbase.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    localDbase.EntityData.Children = types.NewOrderedMap()
    localDbase.EntityData.Children.Append("local-dbase-rloc", types.YChild{"LocalDbaseRloc", nil})
    for i := range localDbase.LocalDbaseRloc {
        localDbase.EntityData.Children.Append(types.GetSegmentPath(localDbase.LocalDbaseRloc[i]), types.YChild{"LocalDbaseRloc", localDbase.LocalDbaseRloc[i]})
    }
    localDbase.EntityData.Leafs = types.NewOrderedMap()
    localDbase.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", localDbase.Afi})
    localDbase.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", localDbase.Prefix})
    localDbase.EntityData.Leafs.Append("lsb", types.YLeaf{"Lsb", localDbase.Lsb})
    localDbase.EntityData.Leafs.Append("is-reachable", types.YLeaf{"IsReachable", localDbase.IsReachable})

    localDbase.EntityData.YListKeys = []string {"Afi", "Prefix"}

    return &(localDbase.EntityData)
}

// LispState_LispRouters_Instances_Af_LocalDbase_LocalDbaseRloc
// List of locators
type LispState_LispRouters_Instances_Af_LocalDbase_LocalDbaseRloc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. LISP Address-Family of the address. The type is
    // LispAddressFamilyType.
    Afi interface{}

    // This attribute is a key. LISP address. Format is defined by the AF. The
    // type is string.
    Address interface{}

    // Up/down state of the locator. The type is LispRlocStateType.
    State interface{}

    // Indicates if RLOC local to the device or to another xTR in the site; TRUE
    // means RLOC is local to the device. The type is bool.
    IsLocal interface{}

    // Properties of the locator.
    Params LispState_LispRouters_Instances_Af_LocalDbase_LocalDbaseRloc_Params
}

func (localDbaseRloc *LispState_LispRouters_Instances_Af_LocalDbase_LocalDbaseRloc) GetEntityData() *types.CommonEntityData {
    localDbaseRloc.EntityData.YFilter = localDbaseRloc.YFilter
    localDbaseRloc.EntityData.YangName = "local-dbase-rloc"
    localDbaseRloc.EntityData.BundleName = "cisco_ios_xe"
    localDbaseRloc.EntityData.ParentYangName = "local-dbase"
    localDbaseRloc.EntityData.SegmentPath = "local-dbase-rloc" + types.AddKeyToken(localDbaseRloc.Afi, "afi") + types.AddKeyToken(localDbaseRloc.Address, "address")
    localDbaseRloc.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    localDbaseRloc.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    localDbaseRloc.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    localDbaseRloc.EntityData.Children = types.NewOrderedMap()
    localDbaseRloc.EntityData.Children.Append("params", types.YChild{"Params", &localDbaseRloc.Params})
    localDbaseRloc.EntityData.Leafs = types.NewOrderedMap()
    localDbaseRloc.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", localDbaseRloc.Afi})
    localDbaseRloc.EntityData.Leafs.Append("address", types.YLeaf{"Address", localDbaseRloc.Address})
    localDbaseRloc.EntityData.Leafs.Append("state", types.YLeaf{"State", localDbaseRloc.State})
    localDbaseRloc.EntityData.Leafs.Append("is-local", types.YLeaf{"IsLocal", localDbaseRloc.IsLocal})

    localDbaseRloc.EntityData.YListKeys = []string {"Afi", "Address"}

    return &(localDbaseRloc.EntityData)
}

// LispState_LispRouters_Instances_Af_LocalDbase_LocalDbaseRloc_Params
// Properties of the locator
type LispState_LispRouters_Instances_Af_LocalDbase_LocalDbaseRloc_Params struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Locator priority. The type is interface{} with range: 0..255.
    Priority interface{}

    // Locator weight. The type is interface{} with range: 0..255.
    Weight interface{}

    // Locator's multicast priority. The type is interface{} with range: 0..255.
    McastPriority interface{}

    // Locator's multicast weight. The type is interface{} with range: 0..255.
    McastWeight interface{}
}

func (params *LispState_LispRouters_Instances_Af_LocalDbase_LocalDbaseRloc_Params) GetEntityData() *types.CommonEntityData {
    params.EntityData.YFilter = params.YFilter
    params.EntityData.YangName = "params"
    params.EntityData.BundleName = "cisco_ios_xe"
    params.EntityData.ParentYangName = "local-dbase-rloc"
    params.EntityData.SegmentPath = "params"
    params.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    params.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    params.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    params.EntityData.Children = types.NewOrderedMap()
    params.EntityData.Leafs = types.NewOrderedMap()
    params.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", params.Priority})
    params.EntityData.Leafs.Append("weight", types.YLeaf{"Weight", params.Weight})
    params.EntityData.Leafs.Append("mcast-priority", types.YLeaf{"McastPriority", params.McastPriority})
    params.EntityData.Leafs.Append("mcast-weight", types.YLeaf{"McastWeight", params.McastWeight})

    params.EntityData.YListKeys = []string {}

    return &(params.EntityData)
}

// LispState_LispRouters_Instances_Af_MsRegistrations
// Map-Server database of registered EID Prefixes
type LispState_LispRouters_Instances_Af_MsRegistrations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. LISP Address-Family of the prefix. The type is
    // LispAddressFamilyType.
    Afi interface{}

    // This attribute is a key. LISP prefix. Format is defined by the AF. The type
    // is string.
    Prefix interface{}

    // Source port of the last valid registration received for this EID prefix.
    // The type is interface{} with range: 0..65535.
    LastRegistrationSourcePort interface{}

    // Time when a valid registration was first received for this EID prefix. The
    // type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    CreationTime interface{}

    // Time when most recent valid registration was received for this EID prefix.
    // The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    LastRegistrationTime interface{}

    // Name of site matching this registration. The type is string.
    SiteName interface{}

    // Description given for the site matching this registration. The type is
    // string.
    SiteDescription interface{}

    // Indicates the registration status of the given EID-Prefix. If this object
    // is true, then it means the EID-Prefix is registered. The type is bool.
    IsRegistered interface{}

    // Count of registrations received for the EID prefix with authentication
    // error. The type is interface{} with range: 0..18446744073709551615.
    AuthenticationError interface{}

    // Count of received registrations for the prefix that had at least one RLOC
    // that was not in the allowed list of RLOCs. The type is interface{} with
    // range: 0..18446744073709551615.
    RlocMismatchError interface{}

    // Source address of the last valid registration received for this EID prefix.
    LastRegistrationSource LispState_LispRouters_Instances_Af_MsRegistrations_LastRegistrationSource

    // List of registrations for this EID prefix received from different ETRs. The
    // type is slice of
    // LispState_LispRouters_Instances_Af_MsRegistrations_EtrRegistrations.
    EtrRegistrations []*LispState_LispRouters_Instances_Af_MsRegistrations_EtrRegistrations
}

func (msRegistrations *LispState_LispRouters_Instances_Af_MsRegistrations) GetEntityData() *types.CommonEntityData {
    msRegistrations.EntityData.YFilter = msRegistrations.YFilter
    msRegistrations.EntityData.YangName = "ms-registrations"
    msRegistrations.EntityData.BundleName = "cisco_ios_xe"
    msRegistrations.EntityData.ParentYangName = "af"
    msRegistrations.EntityData.SegmentPath = "ms-registrations" + types.AddKeyToken(msRegistrations.Afi, "afi") + types.AddKeyToken(msRegistrations.Prefix, "prefix")
    msRegistrations.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    msRegistrations.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    msRegistrations.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    msRegistrations.EntityData.Children = types.NewOrderedMap()
    msRegistrations.EntityData.Children.Append("last-registration-source", types.YChild{"LastRegistrationSource", &msRegistrations.LastRegistrationSource})
    msRegistrations.EntityData.Children.Append("etr-registrations", types.YChild{"EtrRegistrations", nil})
    for i := range msRegistrations.EtrRegistrations {
        msRegistrations.EntityData.Children.Append(types.GetSegmentPath(msRegistrations.EtrRegistrations[i]), types.YChild{"EtrRegistrations", msRegistrations.EtrRegistrations[i]})
    }
    msRegistrations.EntityData.Leafs = types.NewOrderedMap()
    msRegistrations.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", msRegistrations.Afi})
    msRegistrations.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", msRegistrations.Prefix})
    msRegistrations.EntityData.Leafs.Append("last-registration-source-port", types.YLeaf{"LastRegistrationSourcePort", msRegistrations.LastRegistrationSourcePort})
    msRegistrations.EntityData.Leafs.Append("creation-time", types.YLeaf{"CreationTime", msRegistrations.CreationTime})
    msRegistrations.EntityData.Leafs.Append("last-registration-time", types.YLeaf{"LastRegistrationTime", msRegistrations.LastRegistrationTime})
    msRegistrations.EntityData.Leafs.Append("site-name", types.YLeaf{"SiteName", msRegistrations.SiteName})
    msRegistrations.EntityData.Leafs.Append("site-description", types.YLeaf{"SiteDescription", msRegistrations.SiteDescription})
    msRegistrations.EntityData.Leafs.Append("is-registered", types.YLeaf{"IsRegistered", msRegistrations.IsRegistered})
    msRegistrations.EntityData.Leafs.Append("authentication-error", types.YLeaf{"AuthenticationError", msRegistrations.AuthenticationError})
    msRegistrations.EntityData.Leafs.Append("rloc-mismatch-error", types.YLeaf{"RlocMismatchError", msRegistrations.RlocMismatchError})

    msRegistrations.EntityData.YListKeys = []string {"Afi", "Prefix"}

    return &(msRegistrations.EntityData)
}

// LispState_LispRouters_Instances_Af_MsRegistrations_LastRegistrationSource
// Source address of the last valid registration received
// for this EID prefix
type LispState_LispRouters_Instances_Af_MsRegistrations_LastRegistrationSource struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LISP Address-Family of the address. The type is LispAddressFamilyType.
    Afi interface{}

    // LISP address. Format is defined by the AF. The type is string.
    Address interface{}
}

func (lastRegistrationSource *LispState_LispRouters_Instances_Af_MsRegistrations_LastRegistrationSource) GetEntityData() *types.CommonEntityData {
    lastRegistrationSource.EntityData.YFilter = lastRegistrationSource.YFilter
    lastRegistrationSource.EntityData.YangName = "last-registration-source"
    lastRegistrationSource.EntityData.BundleName = "cisco_ios_xe"
    lastRegistrationSource.EntityData.ParentYangName = "ms-registrations"
    lastRegistrationSource.EntityData.SegmentPath = "last-registration-source"
    lastRegistrationSource.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lastRegistrationSource.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lastRegistrationSource.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lastRegistrationSource.EntityData.Children = types.NewOrderedMap()
    lastRegistrationSource.EntityData.Leafs = types.NewOrderedMap()
    lastRegistrationSource.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", lastRegistrationSource.Afi})
    lastRegistrationSource.EntityData.Leafs.Append("address", types.YLeaf{"Address", lastRegistrationSource.Address})

    lastRegistrationSource.EntityData.YListKeys = []string {}

    return &(lastRegistrationSource.EntityData)
}

// LispState_LispRouters_Instances_Af_MsRegistrations_EtrRegistrations
// List of registrations for this EID prefix received
// from different ETRs
type LispState_LispRouters_Instances_Af_MsRegistrations_EtrRegistrations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. RLOC address of the registration source. The type
    // is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // This attribute is a key. Port of the registration source. The type is
    // interface{} with range: 0..65535.
    SourcePort interface{}

    // Time when valid registration from the source was last received. The type is
    // string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    LastRegistrationTime interface{}

    // Registration validity period. The type is interface{} with range:
    // 0..4294967295. Units are minutes.
    Ttl interface{}

    // Indicates if the registering ETR requested proxy-replying on Map-Requests
    // by the Map-Server. The type is bool.
    ProxyReply interface{}

    // Indicates if the registering ETR wants to be informed about matching
    // registrations with Map-Notifies. The type is bool.
    WantsMapNotify interface{}

    // List of locators. The type is slice of
    // LispState_LispRouters_Instances_Af_MsRegistrations_EtrRegistrations_MsRegistrationRloc.
    MsRegistrationRloc []*LispState_LispRouters_Instances_Af_MsRegistrations_EtrRegistrations_MsRegistrationRloc
}

func (etrRegistrations *LispState_LispRouters_Instances_Af_MsRegistrations_EtrRegistrations) GetEntityData() *types.CommonEntityData {
    etrRegistrations.EntityData.YFilter = etrRegistrations.YFilter
    etrRegistrations.EntityData.YangName = "etr-registrations"
    etrRegistrations.EntityData.BundleName = "cisco_ios_xe"
    etrRegistrations.EntityData.ParentYangName = "ms-registrations"
    etrRegistrations.EntityData.SegmentPath = "etr-registrations" + types.AddKeyToken(etrRegistrations.SourceAddress, "source-address") + types.AddKeyToken(etrRegistrations.SourcePort, "source-port")
    etrRegistrations.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    etrRegistrations.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    etrRegistrations.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    etrRegistrations.EntityData.Children = types.NewOrderedMap()
    etrRegistrations.EntityData.Children.Append("ms-registration-rloc", types.YChild{"MsRegistrationRloc", nil})
    for i := range etrRegistrations.MsRegistrationRloc {
        etrRegistrations.EntityData.Children.Append(types.GetSegmentPath(etrRegistrations.MsRegistrationRloc[i]), types.YChild{"MsRegistrationRloc", etrRegistrations.MsRegistrationRloc[i]})
    }
    etrRegistrations.EntityData.Leafs = types.NewOrderedMap()
    etrRegistrations.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", etrRegistrations.SourceAddress})
    etrRegistrations.EntityData.Leafs.Append("source-port", types.YLeaf{"SourcePort", etrRegistrations.SourcePort})
    etrRegistrations.EntityData.Leafs.Append("last-registration-time", types.YLeaf{"LastRegistrationTime", etrRegistrations.LastRegistrationTime})
    etrRegistrations.EntityData.Leafs.Append("ttl", types.YLeaf{"Ttl", etrRegistrations.Ttl})
    etrRegistrations.EntityData.Leafs.Append("proxy-reply", types.YLeaf{"ProxyReply", etrRegistrations.ProxyReply})
    etrRegistrations.EntityData.Leafs.Append("wants-map-notify", types.YLeaf{"WantsMapNotify", etrRegistrations.WantsMapNotify})

    etrRegistrations.EntityData.YListKeys = []string {"SourceAddress", "SourcePort"}

    return &(etrRegistrations.EntityData)
}

// LispState_LispRouters_Instances_Af_MsRegistrations_EtrRegistrations_MsRegistrationRloc
// List of locators
type LispState_LispRouters_Instances_Af_MsRegistrations_EtrRegistrations_MsRegistrationRloc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. LISP Address-Family of the address. The type is
    // LispAddressFamilyType.
    Afi interface{}

    // This attribute is a key. LISP address. Format is defined by the AF. The
    // type is string.
    Address interface{}

    // Up/down state of the locator. The type is LispRlocStateType.
    State interface{}

    // Indicates if RLOC is local to device which sent registration or to another
    // xTR in the site; TRUE means RLOC is local to the registering device. The
    // type is bool.
    IsLocal interface{}

    // Properties of the locator.
    Params LispState_LispRouters_Instances_Af_MsRegistrations_EtrRegistrations_MsRegistrationRloc_Params
}

func (msRegistrationRloc *LispState_LispRouters_Instances_Af_MsRegistrations_EtrRegistrations_MsRegistrationRloc) GetEntityData() *types.CommonEntityData {
    msRegistrationRloc.EntityData.YFilter = msRegistrationRloc.YFilter
    msRegistrationRloc.EntityData.YangName = "ms-registration-rloc"
    msRegistrationRloc.EntityData.BundleName = "cisco_ios_xe"
    msRegistrationRloc.EntityData.ParentYangName = "etr-registrations"
    msRegistrationRloc.EntityData.SegmentPath = "ms-registration-rloc" + types.AddKeyToken(msRegistrationRloc.Afi, "afi") + types.AddKeyToken(msRegistrationRloc.Address, "address")
    msRegistrationRloc.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    msRegistrationRloc.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    msRegistrationRloc.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    msRegistrationRloc.EntityData.Children = types.NewOrderedMap()
    msRegistrationRloc.EntityData.Children.Append("params", types.YChild{"Params", &msRegistrationRloc.Params})
    msRegistrationRloc.EntityData.Leafs = types.NewOrderedMap()
    msRegistrationRloc.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", msRegistrationRloc.Afi})
    msRegistrationRloc.EntityData.Leafs.Append("address", types.YLeaf{"Address", msRegistrationRloc.Address})
    msRegistrationRloc.EntityData.Leafs.Append("state", types.YLeaf{"State", msRegistrationRloc.State})
    msRegistrationRloc.EntityData.Leafs.Append("is-local", types.YLeaf{"IsLocal", msRegistrationRloc.IsLocal})

    msRegistrationRloc.EntityData.YListKeys = []string {"Afi", "Address"}

    return &(msRegistrationRloc.EntityData)
}

// LispState_LispRouters_Instances_Af_MsRegistrations_EtrRegistrations_MsRegistrationRloc_Params
// Properties of the locator
type LispState_LispRouters_Instances_Af_MsRegistrations_EtrRegistrations_MsRegistrationRloc_Params struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Locator priority. The type is interface{} with range: 0..255.
    Priority interface{}

    // Locator weight. The type is interface{} with range: 0..255.
    Weight interface{}

    // Locator's multicast priority. The type is interface{} with range: 0..255.
    McastPriority interface{}

    // Locator's multicast weight. The type is interface{} with range: 0..255.
    McastWeight interface{}
}

func (params *LispState_LispRouters_Instances_Af_MsRegistrations_EtrRegistrations_MsRegistrationRloc_Params) GetEntityData() *types.CommonEntityData {
    params.EntityData.YFilter = params.YFilter
    params.EntityData.YangName = "params"
    params.EntityData.BundleName = "cisco_ios_xe"
    params.EntityData.ParentYangName = "ms-registration-rloc"
    params.EntityData.SegmentPath = "params"
    params.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    params.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    params.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    params.EntityData.Children = types.NewOrderedMap()
    params.EntityData.Leafs = types.NewOrderedMap()
    params.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", params.Priority})
    params.EntityData.Leafs.Append("weight", types.YLeaf{"Weight", params.Weight})
    params.EntityData.Leafs.Append("mcast-priority", types.YLeaf{"McastPriority", params.McastPriority})
    params.EntityData.Leafs.Append("mcast-weight", types.YLeaf{"McastWeight", params.McastWeight})

    params.EntityData.YListKeys = []string {}

    return &(params.EntityData)
}

// LispState_LispRouters_Instances_Af_MapServers
// List of Map-Servers to which the ETR should register
type LispState_LispRouters_Instances_Af_MapServers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. LISP Address-Family of the address. The type is
    // LispAddressFamilyType.
    Afi interface{}

    // This attribute is a key. LISP address. Format is defined by the AF. The
    // type is string.
    Address interface{}

    // Up/down state of the locator. The type is LispRlocStateType.
    State interface{}
}

func (mapServers *LispState_LispRouters_Instances_Af_MapServers) GetEntityData() *types.CommonEntityData {
    mapServers.EntityData.YFilter = mapServers.YFilter
    mapServers.EntityData.YangName = "map-servers"
    mapServers.EntityData.BundleName = "cisco_ios_xe"
    mapServers.EntityData.ParentYangName = "af"
    mapServers.EntityData.SegmentPath = "map-servers" + types.AddKeyToken(mapServers.Afi, "afi") + types.AddKeyToken(mapServers.Address, "address")
    mapServers.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mapServers.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mapServers.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mapServers.EntityData.Children = types.NewOrderedMap()
    mapServers.EntityData.Leafs = types.NewOrderedMap()
    mapServers.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", mapServers.Afi})
    mapServers.EntityData.Leafs.Append("address", types.YLeaf{"Address", mapServers.Address})
    mapServers.EntityData.Leafs.Append("state", types.YLeaf{"State", mapServers.State})

    mapServers.EntityData.YListKeys = []string {"Afi", "Address"}

    return &(mapServers.EntityData)
}

// LispState_LispRouters_Instances_Af_MapResolvers
// List of Map-Resolvers where [P]ITR should send its
// Map-Requests
type LispState_LispRouters_Instances_Af_MapResolvers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. LISP Address-Family of the address. The type is
    // LispAddressFamilyType.
    Afi interface{}

    // This attribute is a key. LISP address. Format is defined by the AF. The
    // type is string.
    Address interface{}

    // Up/down state of the locator. The type is LispRlocStateType.
    State interface{}
}

func (mapResolvers *LispState_LispRouters_Instances_Af_MapResolvers) GetEntityData() *types.CommonEntityData {
    mapResolvers.EntityData.YFilter = mapResolvers.YFilter
    mapResolvers.EntityData.YangName = "map-resolvers"
    mapResolvers.EntityData.BundleName = "cisco_ios_xe"
    mapResolvers.EntityData.ParentYangName = "af"
    mapResolvers.EntityData.SegmentPath = "map-resolvers" + types.AddKeyToken(mapResolvers.Afi, "afi") + types.AddKeyToken(mapResolvers.Address, "address")
    mapResolvers.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mapResolvers.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mapResolvers.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mapResolvers.EntityData.Children = types.NewOrderedMap()
    mapResolvers.EntityData.Leafs = types.NewOrderedMap()
    mapResolvers.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", mapResolvers.Afi})
    mapResolvers.EntityData.Leafs.Append("address", types.YLeaf{"Address", mapResolvers.Address})
    mapResolvers.EntityData.Leafs.Append("state", types.YLeaf{"State", mapResolvers.State})

    mapResolvers.EntityData.YListKeys = []string {"Afi", "Address"}

    return &(mapResolvers.EntityData)
}

// LispState_LispRouters_Instances_Af_ProxyEtrs
// List of all Proxy ETRs that this device is configured
// to use
type LispState_LispRouters_Instances_Af_ProxyEtrs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. LISP Address-Family of the address. The type is
    // LispAddressFamilyType.
    Afi interface{}

    // This attribute is a key. LISP address. Format is defined by the AF. The
    // type is string.
    Address interface{}

    // Up/down state of the locator. The type is LispRlocStateType.
    State interface{}

    // Properties of the locator.
    Params LispState_LispRouters_Instances_Af_ProxyEtrs_Params
}

func (proxyEtrs *LispState_LispRouters_Instances_Af_ProxyEtrs) GetEntityData() *types.CommonEntityData {
    proxyEtrs.EntityData.YFilter = proxyEtrs.YFilter
    proxyEtrs.EntityData.YangName = "proxy-etrs"
    proxyEtrs.EntityData.BundleName = "cisco_ios_xe"
    proxyEtrs.EntityData.ParentYangName = "af"
    proxyEtrs.EntityData.SegmentPath = "proxy-etrs" + types.AddKeyToken(proxyEtrs.Afi, "afi") + types.AddKeyToken(proxyEtrs.Address, "address")
    proxyEtrs.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    proxyEtrs.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    proxyEtrs.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    proxyEtrs.EntityData.Children = types.NewOrderedMap()
    proxyEtrs.EntityData.Children.Append("params", types.YChild{"Params", &proxyEtrs.Params})
    proxyEtrs.EntityData.Leafs = types.NewOrderedMap()
    proxyEtrs.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", proxyEtrs.Afi})
    proxyEtrs.EntityData.Leafs.Append("address", types.YLeaf{"Address", proxyEtrs.Address})
    proxyEtrs.EntityData.Leafs.Append("state", types.YLeaf{"State", proxyEtrs.State})

    proxyEtrs.EntityData.YListKeys = []string {"Afi", "Address"}

    return &(proxyEtrs.EntityData)
}

// LispState_LispRouters_Instances_Af_ProxyEtrs_Params
// Properties of the locator
type LispState_LispRouters_Instances_Af_ProxyEtrs_Params struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Locator priority. The type is interface{} with range: 0..255.
    Priority interface{}

    // Locator weight. The type is interface{} with range: 0..255.
    Weight interface{}

    // Locator's multicast priority. The type is interface{} with range: 0..255.
    McastPriority interface{}

    // Locator's multicast weight. The type is interface{} with range: 0..255.
    McastWeight interface{}
}

func (params *LispState_LispRouters_Instances_Af_ProxyEtrs_Params) GetEntityData() *types.CommonEntityData {
    params.EntityData.YFilter = params.YFilter
    params.EntityData.YangName = "params"
    params.EntityData.BundleName = "cisco_ios_xe"
    params.EntityData.ParentYangName = "proxy-etrs"
    params.EntityData.SegmentPath = "params"
    params.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    params.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    params.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    params.EntityData.Children = types.NewOrderedMap()
    params.EntityData.Leafs = types.NewOrderedMap()
    params.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", params.Priority})
    params.EntityData.Leafs.Append("weight", types.YLeaf{"Weight", params.Weight})
    params.EntityData.Leafs.Append("mcast-priority", types.YLeaf{"McastPriority", params.McastPriority})
    params.EntityData.Leafs.Append("mcast-weight", types.YLeaf{"McastWeight", params.McastWeight})

    params.EntityData.YListKeys = []string {}

    return &(params.EntityData)
}

// LispState_LispRouters_Instances_MsEidMembership
// MS registration EID membership list (list of locators
// known to the MS as allowed to send traffic in the
// instance)
type LispState_LispRouters_Instances_MsEidMembership struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. RLOC which is the allowed member. The type is one
    // of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Rloc interface{}

    // Time when this RLOC was added to the list of allowed locators. The type is
    // string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    MemberSince interface{}

    // Indicates if MS gleaned this RLOC from received EID prefix registration.
    // The type is bool.
    IsGleaned interface{}

    // Indicates if this RLOC membership was provided via configuration. The type
    // is bool.
    IsConfigured interface{}
}

func (msEidMembership *LispState_LispRouters_Instances_MsEidMembership) GetEntityData() *types.CommonEntityData {
    msEidMembership.EntityData.YFilter = msEidMembership.YFilter
    msEidMembership.EntityData.YangName = "ms-eid-membership"
    msEidMembership.EntityData.BundleName = "cisco_ios_xe"
    msEidMembership.EntityData.ParentYangName = "instances"
    msEidMembership.EntityData.SegmentPath = "ms-eid-membership" + types.AddKeyToken(msEidMembership.Rloc, "rloc")
    msEidMembership.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    msEidMembership.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    msEidMembership.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    msEidMembership.EntityData.Children = types.NewOrderedMap()
    msEidMembership.EntityData.Leafs = types.NewOrderedMap()
    msEidMembership.EntityData.Leafs.Append("rloc", types.YLeaf{"Rloc", msEidMembership.Rloc})
    msEidMembership.EntityData.Leafs.Append("member-since", types.YLeaf{"MemberSince", msEidMembership.MemberSince})
    msEidMembership.EntityData.Leafs.Append("is-gleaned", types.YLeaf{"IsGleaned", msEidMembership.IsGleaned})
    msEidMembership.EntityData.Leafs.Append("is-configured", types.YLeaf{"IsConfigured", msEidMembership.IsConfigured})

    msEidMembership.EntityData.YListKeys = []string {"Rloc"}

    return &(msEidMembership.EntityData)
}

// LispState_LispRouters_Instances_EtrEidMembership
// ETR EID membership list (list of locators known to the ETR
// as allowed to send traffic in the instance)
type LispState_LispRouters_Instances_EtrEidMembership struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. RLOC which is the allowed member. The type is one
    // of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Rloc interface{}

    // Time when this RLOC was added to the list of allowed locators. The type is
    // string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    MemberSince interface{}

    // Indicates if ETR learned about this RLOC membership via message received
    // from a Map-Server. The type is bool.
    IsLearnedFromMs interface{}

    // Indicates if this RLOC membership was provided via configuration. The type
    // is bool.
    IsConfigured interface{}
}

func (etrEidMembership *LispState_LispRouters_Instances_EtrEidMembership) GetEntityData() *types.CommonEntityData {
    etrEidMembership.EntityData.YFilter = etrEidMembership.YFilter
    etrEidMembership.EntityData.YangName = "etr-eid-membership"
    etrEidMembership.EntityData.BundleName = "cisco_ios_xe"
    etrEidMembership.EntityData.ParentYangName = "instances"
    etrEidMembership.EntityData.SegmentPath = "etr-eid-membership" + types.AddKeyToken(etrEidMembership.Rloc, "rloc")
    etrEidMembership.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    etrEidMembership.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    etrEidMembership.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    etrEidMembership.EntityData.Children = types.NewOrderedMap()
    etrEidMembership.EntityData.Leafs = types.NewOrderedMap()
    etrEidMembership.EntityData.Leafs.Append("rloc", types.YLeaf{"Rloc", etrEidMembership.Rloc})
    etrEidMembership.EntityData.Leafs.Append("member-since", types.YLeaf{"MemberSince", etrEidMembership.MemberSince})
    etrEidMembership.EntityData.Leafs.Append("is-learned-from-ms", types.YLeaf{"IsLearnedFromMs", etrEidMembership.IsLearnedFromMs})
    etrEidMembership.EntityData.Leafs.Append("is-configured", types.YLeaf{"IsConfigured", etrEidMembership.IsConfigured})

    etrEidMembership.EntityData.YListKeys = []string {"Rloc"}

    return &(etrEidMembership.EntityData)
}

// LispState_LispRouters_Sessions
// List of Reliable Registration sessions
type LispState_LispRouters_Sessions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Address of the local socket. The type is one of
    // the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    LocalAddress interface{}

    // This attribute is a key. Address of the peer. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    PeerAddress interface{}

    // This attribute is a key. Port of the local socket. The type is interface{}
    // with range: 0..65535.
    LocalPort interface{}

    // This attribute is a key. Port used by the peer. The type is interface{}
    // with range: 0..65535.
    PeerPort interface{}

    // Up/down state of the session. The type is LispSessionStateType.
    State interface{}

    // Timestamp when the session's state last changed. The type is string with
    // pattern: \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    LastStateChangeTime interface{}

    // Is session opening role Active or Passive; TRUE means session role is
    // Active. The type is bool.
    IsRoleActive interface{}

    // Is route to peer's address known. The type is bool.
    IsRoutable interface{}

    // Number of messages received over this session. The type is interface{} with
    // range: 0..18446744073709551615.
    MessagesIn interface{}

    // Number of messages sent over this session. The type is interface{} with
    // range: 0..18446744073709551615.
    MessagesOut interface{}

    // Number of bytes received over this session. The type is interface{} with
    // range: 0..18446744073709551615.
    BytesIn interface{}

    // Number of bytes sent over this session. The type is interface{} with range:
    // 0..18446744073709551615.
    BytesOut interface{}
}

func (sessions *LispState_LispRouters_Sessions) GetEntityData() *types.CommonEntityData {
    sessions.EntityData.YFilter = sessions.YFilter
    sessions.EntityData.YangName = "sessions"
    sessions.EntityData.BundleName = "cisco_ios_xe"
    sessions.EntityData.ParentYangName = "lisp-routers"
    sessions.EntityData.SegmentPath = "sessions" + types.AddKeyToken(sessions.LocalAddress, "local-address") + types.AddKeyToken(sessions.PeerAddress, "peer-address") + types.AddKeyToken(sessions.LocalPort, "local-port") + types.AddKeyToken(sessions.PeerPort, "peer-port")
    sessions.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sessions.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sessions.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sessions.EntityData.Children = types.NewOrderedMap()
    sessions.EntityData.Leafs = types.NewOrderedMap()
    sessions.EntityData.Leafs.Append("local-address", types.YLeaf{"LocalAddress", sessions.LocalAddress})
    sessions.EntityData.Leafs.Append("peer-address", types.YLeaf{"PeerAddress", sessions.PeerAddress})
    sessions.EntityData.Leafs.Append("local-port", types.YLeaf{"LocalPort", sessions.LocalPort})
    sessions.EntityData.Leafs.Append("peer-port", types.YLeaf{"PeerPort", sessions.PeerPort})
    sessions.EntityData.Leafs.Append("state", types.YLeaf{"State", sessions.State})
    sessions.EntityData.Leafs.Append("last-state-change-time", types.YLeaf{"LastStateChangeTime", sessions.LastStateChangeTime})
    sessions.EntityData.Leafs.Append("is-role-active", types.YLeaf{"IsRoleActive", sessions.IsRoleActive})
    sessions.EntityData.Leafs.Append("is-routable", types.YLeaf{"IsRoutable", sessions.IsRoutable})
    sessions.EntityData.Leafs.Append("messages-in", types.YLeaf{"MessagesIn", sessions.MessagesIn})
    sessions.EntityData.Leafs.Append("messages-out", types.YLeaf{"MessagesOut", sessions.MessagesOut})
    sessions.EntityData.Leafs.Append("bytes-in", types.YLeaf{"BytesIn", sessions.BytesIn})
    sessions.EntityData.Leafs.Append("bytes-out", types.YLeaf{"BytesOut", sessions.BytesOut})

    sessions.EntityData.YListKeys = []string {"LocalAddress", "PeerAddress", "LocalPort", "PeerPort"}

    return &(sessions.EntityData)
}

// LispState_LispRouters_LocalRlocs
// This list represents the set of routing locators
// configured on this device
type LispState_LispRouters_LocalRlocs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. LISP Address-Family of the address. The type is
    // LispAddressFamilyType.
    Afi interface{}

    // This attribute is a key. LISP address. Format is defined by the AF. The
    // type is string.
    Address interface{}

    // Up/down state of the locator. The type is LispRlocStateType.
    State interface{}

    // Indicates if RLOC is local to the device or to another xTR in the site;
    // TRUE means RLOC is local to the device. The type is bool.
    IsLocal interface{}
}

func (localRlocs *LispState_LispRouters_LocalRlocs) GetEntityData() *types.CommonEntityData {
    localRlocs.EntityData.YFilter = localRlocs.YFilter
    localRlocs.EntityData.YangName = "local-rlocs"
    localRlocs.EntityData.BundleName = "cisco_ios_xe"
    localRlocs.EntityData.ParentYangName = "lisp-routers"
    localRlocs.EntityData.SegmentPath = "local-rlocs" + types.AddKeyToken(localRlocs.Afi, "afi") + types.AddKeyToken(localRlocs.Address, "address")
    localRlocs.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    localRlocs.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    localRlocs.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    localRlocs.EntityData.Children = types.NewOrderedMap()
    localRlocs.EntityData.Leafs = types.NewOrderedMap()
    localRlocs.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", localRlocs.Afi})
    localRlocs.EntityData.Leafs.Append("address", types.YLeaf{"Address", localRlocs.Address})
    localRlocs.EntityData.Leafs.Append("state", types.YLeaf{"State", localRlocs.State})
    localRlocs.EntityData.Leafs.Append("is-local", types.YLeaf{"IsLocal", localRlocs.IsLocal})

    localRlocs.EntityData.YListKeys = []string {"Afi", "Address"}

    return &(localRlocs.EntityData)
}

