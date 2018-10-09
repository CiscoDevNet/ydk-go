// This module contains a collection of YANG definitions
// for Cisco IOS-XR tunnel-l2tun package operational data.
// 
// This module contains definitions
// for the following management objects:
//   l2tp: L2TP operational data
//   l2tpv2: l2tpv2
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package tunnel_l2tun_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package tunnel_l2tun_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-tunnel-l2tun-oper l2tp}", reflect.TypeOf(L2tp{}))
    ydk.RegisterEntity("Cisco-IOS-XR-tunnel-l2tun-oper:l2tp", reflect.TypeOf(L2tp{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-tunnel-l2tun-oper l2tpv2}", reflect.TypeOf(L2tpv2{}))
    ydk.RegisterEntity("Cisco-IOS-XR-tunnel-l2tun-oper:l2tpv2", reflect.TypeOf(L2tpv2{}))
}

// DigestHash represents Digest hash types
type DigestHash string

const (
    // MD5
    DigestHash_md5 DigestHash = "md5"

    // SHA1
    DigestHash_sha1 DigestHash = "sha1"
)

// L2tp
// L2TP operational data
type L2tp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2TP control messages counters.
    Counters L2tp_Counters

    // List of tunnel IDs.
    TunnelConfigurations L2tp_TunnelConfigurations

    // Failure events leading to disconnection.
    CounterHistFail L2tp_CounterHistFail

    // List of L2TP class names.
    Classes L2tp_Classes

    // List of tunnel IDs.
    Tunnels L2tp_Tunnels

    // List of session IDs.
    Sessions L2tp_Sessions

    // L2TP control messages counters.
    Session L2tp_Session
}

func (l2tp *L2tp) GetEntityData() *types.CommonEntityData {
    l2tp.EntityData.YFilter = l2tp.YFilter
    l2tp.EntityData.YangName = "l2tp"
    l2tp.EntityData.BundleName = "cisco_ios_xr"
    l2tp.EntityData.ParentYangName = "Cisco-IOS-XR-tunnel-l2tun-oper"
    l2tp.EntityData.SegmentPath = "Cisco-IOS-XR-tunnel-l2tun-oper:l2tp"
    l2tp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    l2tp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    l2tp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    l2tp.EntityData.Children = types.NewOrderedMap()
    l2tp.EntityData.Children.Append("counters", types.YChild{"Counters", &l2tp.Counters})
    l2tp.EntityData.Children.Append("tunnel-configurations", types.YChild{"TunnelConfigurations", &l2tp.TunnelConfigurations})
    l2tp.EntityData.Children.Append("counter-hist-fail", types.YChild{"CounterHistFail", &l2tp.CounterHistFail})
    l2tp.EntityData.Children.Append("classes", types.YChild{"Classes", &l2tp.Classes})
    l2tp.EntityData.Children.Append("tunnels", types.YChild{"Tunnels", &l2tp.Tunnels})
    l2tp.EntityData.Children.Append("sessions", types.YChild{"Sessions", &l2tp.Sessions})
    l2tp.EntityData.Children.Append("session", types.YChild{"Session", &l2tp.Session})
    l2tp.EntityData.Leafs = types.NewOrderedMap()

    l2tp.EntityData.YListKeys = []string {}

    return &(l2tp.EntityData)
}

// L2tp_Counters
// L2TP control messages counters
type L2tp_Counters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2TP control messages counters.
    Control L2tp_Counters_Control
}

func (counters *L2tp_Counters) GetEntityData() *types.CommonEntityData {
    counters.EntityData.YFilter = counters.YFilter
    counters.EntityData.YangName = "counters"
    counters.EntityData.BundleName = "cisco_ios_xr"
    counters.EntityData.ParentYangName = "l2tp"
    counters.EntityData.SegmentPath = "counters"
    counters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    counters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    counters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    counters.EntityData.Children = types.NewOrderedMap()
    counters.EntityData.Children.Append("control", types.YChild{"Control", &counters.Control})
    counters.EntityData.Leafs = types.NewOrderedMap()

    counters.EntityData.YListKeys = []string {}

    return &(counters.EntityData)
}

// L2tp_Counters_Control
// L2TP control messages counters
type L2tp_Counters_Control struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2TP control tunnel messages counters.
    TunnelXr L2tp_Counters_Control_TunnelXr

    // Table of tunnel IDs of control message counters.
    Tunnels L2tp_Counters_Control_Tunnels
}

func (control *L2tp_Counters_Control) GetEntityData() *types.CommonEntityData {
    control.EntityData.YFilter = control.YFilter
    control.EntityData.YangName = "control"
    control.EntityData.BundleName = "cisco_ios_xr"
    control.EntityData.ParentYangName = "counters"
    control.EntityData.SegmentPath = "control"
    control.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    control.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    control.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    control.EntityData.Children = types.NewOrderedMap()
    control.EntityData.Children.Append("tunnel-xr", types.YChild{"TunnelXr", &control.TunnelXr})
    control.EntityData.Children.Append("tunnels", types.YChild{"Tunnels", &control.Tunnels})
    control.EntityData.Leafs = types.NewOrderedMap()

    control.EntityData.YListKeys = []string {}

    return &(control.EntityData)
}

// L2tp_Counters_Control_TunnelXr
// L2TP control tunnel messages counters
type L2tp_Counters_Control_TunnelXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tunnel authentication counters.
    Authentication L2tp_Counters_Control_TunnelXr_Authentication

    // Tunnel counters.
    Global L2tp_Counters_Control_TunnelXr_Global
}

func (tunnelXr *L2tp_Counters_Control_TunnelXr) GetEntityData() *types.CommonEntityData {
    tunnelXr.EntityData.YFilter = tunnelXr.YFilter
    tunnelXr.EntityData.YangName = "tunnel-xr"
    tunnelXr.EntityData.BundleName = "cisco_ios_xr"
    tunnelXr.EntityData.ParentYangName = "control"
    tunnelXr.EntityData.SegmentPath = "tunnel-xr"
    tunnelXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnelXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnelXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnelXr.EntityData.Children = types.NewOrderedMap()
    tunnelXr.EntityData.Children.Append("authentication", types.YChild{"Authentication", &tunnelXr.Authentication})
    tunnelXr.EntityData.Children.Append("global", types.YChild{"Global", &tunnelXr.Global})
    tunnelXr.EntityData.Leafs = types.NewOrderedMap()

    tunnelXr.EntityData.YListKeys = []string {}

    return &(tunnelXr.EntityData)
}

// L2tp_Counters_Control_TunnelXr_Authentication
// Tunnel authentication counters
type L2tp_Counters_Control_TunnelXr_Authentication struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Nonce AVP statistics.
    NonceAvp L2tp_Counters_Control_TunnelXr_Authentication_NonceAvp

    // Common digest statistics.
    CommonDigest L2tp_Counters_Control_TunnelXr_Authentication_CommonDigest

    // Primary digest statistics.
    PrimaryDigest L2tp_Counters_Control_TunnelXr_Authentication_PrimaryDigest

    // Secondary digest statistics.
    SecondaryDigest L2tp_Counters_Control_TunnelXr_Authentication_SecondaryDigest

    // Integrity check statistics.
    IntegrityCheck L2tp_Counters_Control_TunnelXr_Authentication_IntegrityCheck

    // Local secret statistics.
    LocalSecret L2tp_Counters_Control_TunnelXr_Authentication_LocalSecret

    // Challenge AVP statistics.
    ChallengeAvp L2tp_Counters_Control_TunnelXr_Authentication_ChallengeAvp

    // Challenge response statistics.
    ChallengeReponse L2tp_Counters_Control_TunnelXr_Authentication_ChallengeReponse

    // Overall statistics.
    OverallStatistics L2tp_Counters_Control_TunnelXr_Authentication_OverallStatistics
}

func (authentication *L2tp_Counters_Control_TunnelXr_Authentication) GetEntityData() *types.CommonEntityData {
    authentication.EntityData.YFilter = authentication.YFilter
    authentication.EntityData.YangName = "authentication"
    authentication.EntityData.BundleName = "cisco_ios_xr"
    authentication.EntityData.ParentYangName = "tunnel-xr"
    authentication.EntityData.SegmentPath = "authentication"
    authentication.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authentication.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authentication.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authentication.EntityData.Children = types.NewOrderedMap()
    authentication.EntityData.Children.Append("nonce-avp", types.YChild{"NonceAvp", &authentication.NonceAvp})
    authentication.EntityData.Children.Append("common-digest", types.YChild{"CommonDigest", &authentication.CommonDigest})
    authentication.EntityData.Children.Append("primary-digest", types.YChild{"PrimaryDigest", &authentication.PrimaryDigest})
    authentication.EntityData.Children.Append("secondary-digest", types.YChild{"SecondaryDigest", &authentication.SecondaryDigest})
    authentication.EntityData.Children.Append("integrity-check", types.YChild{"IntegrityCheck", &authentication.IntegrityCheck})
    authentication.EntityData.Children.Append("local-secret", types.YChild{"LocalSecret", &authentication.LocalSecret})
    authentication.EntityData.Children.Append("challenge-avp", types.YChild{"ChallengeAvp", &authentication.ChallengeAvp})
    authentication.EntityData.Children.Append("challenge-reponse", types.YChild{"ChallengeReponse", &authentication.ChallengeReponse})
    authentication.EntityData.Children.Append("overall-statistics", types.YChild{"OverallStatistics", &authentication.OverallStatistics})
    authentication.EntityData.Leafs = types.NewOrderedMap()

    authentication.EntityData.YListKeys = []string {}

    return &(authentication.EntityData)
}

// L2tp_Counters_Control_TunnelXr_Authentication_NonceAvp
// Nonce AVP statistics
type L2tp_Counters_Control_TunnelXr_Authentication_NonceAvp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Validate. The type is interface{} with range: 0..4294967295.
    Validate interface{}

    // Bad hash. The type is interface{} with range: 0..4294967295.
    BadHash interface{}

    // Bad length. The type is interface{} with range: 0..4294967295.
    BadLength interface{}

    // Ignored. The type is interface{} with range: 0..4294967295.
    Ignored interface{}

    // Missing. The type is interface{} with range: 0..4294967295.
    Missing interface{}

    // Passed. The type is interface{} with range: 0..4294967295.
    Passed interface{}

    // Failed. The type is interface{} with range: 0..4294967295.
    Failed interface{}

    // Skipped. The type is interface{} with range: 0..4294967295.
    Skipped interface{}

    // Generate response fail. The type is interface{} with range: 0..4294967295.
    GenerateResponseFailures interface{}

    // Unexpected. The type is interface{} with range: 0..4294967295.
    Unexpected interface{}

    // Unexpected ZLB. The type is interface{} with range: 0..4294967295.
    UnexpectedZlb interface{}
}

func (nonceAvp *L2tp_Counters_Control_TunnelXr_Authentication_NonceAvp) GetEntityData() *types.CommonEntityData {
    nonceAvp.EntityData.YFilter = nonceAvp.YFilter
    nonceAvp.EntityData.YangName = "nonce-avp"
    nonceAvp.EntityData.BundleName = "cisco_ios_xr"
    nonceAvp.EntityData.ParentYangName = "authentication"
    nonceAvp.EntityData.SegmentPath = "nonce-avp"
    nonceAvp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nonceAvp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nonceAvp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nonceAvp.EntityData.Children = types.NewOrderedMap()
    nonceAvp.EntityData.Leafs = types.NewOrderedMap()
    nonceAvp.EntityData.Leafs.Append("validate", types.YLeaf{"Validate", nonceAvp.Validate})
    nonceAvp.EntityData.Leafs.Append("bad-hash", types.YLeaf{"BadHash", nonceAvp.BadHash})
    nonceAvp.EntityData.Leafs.Append("bad-length", types.YLeaf{"BadLength", nonceAvp.BadLength})
    nonceAvp.EntityData.Leafs.Append("ignored", types.YLeaf{"Ignored", nonceAvp.Ignored})
    nonceAvp.EntityData.Leafs.Append("missing", types.YLeaf{"Missing", nonceAvp.Missing})
    nonceAvp.EntityData.Leafs.Append("passed", types.YLeaf{"Passed", nonceAvp.Passed})
    nonceAvp.EntityData.Leafs.Append("failed", types.YLeaf{"Failed", nonceAvp.Failed})
    nonceAvp.EntityData.Leafs.Append("skipped", types.YLeaf{"Skipped", nonceAvp.Skipped})
    nonceAvp.EntityData.Leafs.Append("generate-response-failures", types.YLeaf{"GenerateResponseFailures", nonceAvp.GenerateResponseFailures})
    nonceAvp.EntityData.Leafs.Append("unexpected", types.YLeaf{"Unexpected", nonceAvp.Unexpected})
    nonceAvp.EntityData.Leafs.Append("unexpected-zlb", types.YLeaf{"UnexpectedZlb", nonceAvp.UnexpectedZlb})

    nonceAvp.EntityData.YListKeys = []string {}

    return &(nonceAvp.EntityData)
}

// L2tp_Counters_Control_TunnelXr_Authentication_CommonDigest
// Common digest statistics
type L2tp_Counters_Control_TunnelXr_Authentication_CommonDigest struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Validate. The type is interface{} with range: 0..4294967295.
    Validate interface{}

    // Bad hash. The type is interface{} with range: 0..4294967295.
    BadHash interface{}

    // Bad length. The type is interface{} with range: 0..4294967295.
    BadLength interface{}

    // Ignored. The type is interface{} with range: 0..4294967295.
    Ignored interface{}

    // Missing. The type is interface{} with range: 0..4294967295.
    Missing interface{}

    // Passed. The type is interface{} with range: 0..4294967295.
    Passed interface{}

    // Failed. The type is interface{} with range: 0..4294967295.
    Failed interface{}

    // Skipped. The type is interface{} with range: 0..4294967295.
    Skipped interface{}

    // Generate response fail. The type is interface{} with range: 0..4294967295.
    GenerateResponseFailures interface{}

    // Unexpected. The type is interface{} with range: 0..4294967295.
    Unexpected interface{}

    // Unexpected ZLB. The type is interface{} with range: 0..4294967295.
    UnexpectedZlb interface{}
}

func (commonDigest *L2tp_Counters_Control_TunnelXr_Authentication_CommonDigest) GetEntityData() *types.CommonEntityData {
    commonDigest.EntityData.YFilter = commonDigest.YFilter
    commonDigest.EntityData.YangName = "common-digest"
    commonDigest.EntityData.BundleName = "cisco_ios_xr"
    commonDigest.EntityData.ParentYangName = "authentication"
    commonDigest.EntityData.SegmentPath = "common-digest"
    commonDigest.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    commonDigest.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    commonDigest.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    commonDigest.EntityData.Children = types.NewOrderedMap()
    commonDigest.EntityData.Leafs = types.NewOrderedMap()
    commonDigest.EntityData.Leafs.Append("validate", types.YLeaf{"Validate", commonDigest.Validate})
    commonDigest.EntityData.Leafs.Append("bad-hash", types.YLeaf{"BadHash", commonDigest.BadHash})
    commonDigest.EntityData.Leafs.Append("bad-length", types.YLeaf{"BadLength", commonDigest.BadLength})
    commonDigest.EntityData.Leafs.Append("ignored", types.YLeaf{"Ignored", commonDigest.Ignored})
    commonDigest.EntityData.Leafs.Append("missing", types.YLeaf{"Missing", commonDigest.Missing})
    commonDigest.EntityData.Leafs.Append("passed", types.YLeaf{"Passed", commonDigest.Passed})
    commonDigest.EntityData.Leafs.Append("failed", types.YLeaf{"Failed", commonDigest.Failed})
    commonDigest.EntityData.Leafs.Append("skipped", types.YLeaf{"Skipped", commonDigest.Skipped})
    commonDigest.EntityData.Leafs.Append("generate-response-failures", types.YLeaf{"GenerateResponseFailures", commonDigest.GenerateResponseFailures})
    commonDigest.EntityData.Leafs.Append("unexpected", types.YLeaf{"Unexpected", commonDigest.Unexpected})
    commonDigest.EntityData.Leafs.Append("unexpected-zlb", types.YLeaf{"UnexpectedZlb", commonDigest.UnexpectedZlb})

    commonDigest.EntityData.YListKeys = []string {}

    return &(commonDigest.EntityData)
}

// L2tp_Counters_Control_TunnelXr_Authentication_PrimaryDigest
// Primary digest statistics
type L2tp_Counters_Control_TunnelXr_Authentication_PrimaryDigest struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Validate. The type is interface{} with range: 0..4294967295.
    Validate interface{}

    // Bad hash. The type is interface{} with range: 0..4294967295.
    BadHash interface{}

    // Bad length. The type is interface{} with range: 0..4294967295.
    BadLength interface{}

    // Ignored. The type is interface{} with range: 0..4294967295.
    Ignored interface{}

    // Missing. The type is interface{} with range: 0..4294967295.
    Missing interface{}

    // Passed. The type is interface{} with range: 0..4294967295.
    Passed interface{}

    // Failed. The type is interface{} with range: 0..4294967295.
    Failed interface{}

    // Skipped. The type is interface{} with range: 0..4294967295.
    Skipped interface{}

    // Generate response fail. The type is interface{} with range: 0..4294967295.
    GenerateResponseFailures interface{}

    // Unexpected. The type is interface{} with range: 0..4294967295.
    Unexpected interface{}

    // Unexpected ZLB. The type is interface{} with range: 0..4294967295.
    UnexpectedZlb interface{}
}

func (primaryDigest *L2tp_Counters_Control_TunnelXr_Authentication_PrimaryDigest) GetEntityData() *types.CommonEntityData {
    primaryDigest.EntityData.YFilter = primaryDigest.YFilter
    primaryDigest.EntityData.YangName = "primary-digest"
    primaryDigest.EntityData.BundleName = "cisco_ios_xr"
    primaryDigest.EntityData.ParentYangName = "authentication"
    primaryDigest.EntityData.SegmentPath = "primary-digest"
    primaryDigest.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    primaryDigest.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    primaryDigest.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    primaryDigest.EntityData.Children = types.NewOrderedMap()
    primaryDigest.EntityData.Leafs = types.NewOrderedMap()
    primaryDigest.EntityData.Leafs.Append("validate", types.YLeaf{"Validate", primaryDigest.Validate})
    primaryDigest.EntityData.Leafs.Append("bad-hash", types.YLeaf{"BadHash", primaryDigest.BadHash})
    primaryDigest.EntityData.Leafs.Append("bad-length", types.YLeaf{"BadLength", primaryDigest.BadLength})
    primaryDigest.EntityData.Leafs.Append("ignored", types.YLeaf{"Ignored", primaryDigest.Ignored})
    primaryDigest.EntityData.Leafs.Append("missing", types.YLeaf{"Missing", primaryDigest.Missing})
    primaryDigest.EntityData.Leafs.Append("passed", types.YLeaf{"Passed", primaryDigest.Passed})
    primaryDigest.EntityData.Leafs.Append("failed", types.YLeaf{"Failed", primaryDigest.Failed})
    primaryDigest.EntityData.Leafs.Append("skipped", types.YLeaf{"Skipped", primaryDigest.Skipped})
    primaryDigest.EntityData.Leafs.Append("generate-response-failures", types.YLeaf{"GenerateResponseFailures", primaryDigest.GenerateResponseFailures})
    primaryDigest.EntityData.Leafs.Append("unexpected", types.YLeaf{"Unexpected", primaryDigest.Unexpected})
    primaryDigest.EntityData.Leafs.Append("unexpected-zlb", types.YLeaf{"UnexpectedZlb", primaryDigest.UnexpectedZlb})

    primaryDigest.EntityData.YListKeys = []string {}

    return &(primaryDigest.EntityData)
}

// L2tp_Counters_Control_TunnelXr_Authentication_SecondaryDigest
// Secondary digest statistics
type L2tp_Counters_Control_TunnelXr_Authentication_SecondaryDigest struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Validate. The type is interface{} with range: 0..4294967295.
    Validate interface{}

    // Bad hash. The type is interface{} with range: 0..4294967295.
    BadHash interface{}

    // Bad length. The type is interface{} with range: 0..4294967295.
    BadLength interface{}

    // Ignored. The type is interface{} with range: 0..4294967295.
    Ignored interface{}

    // Missing. The type is interface{} with range: 0..4294967295.
    Missing interface{}

    // Passed. The type is interface{} with range: 0..4294967295.
    Passed interface{}

    // Failed. The type is interface{} with range: 0..4294967295.
    Failed interface{}

    // Skipped. The type is interface{} with range: 0..4294967295.
    Skipped interface{}

    // Generate response fail. The type is interface{} with range: 0..4294967295.
    GenerateResponseFailures interface{}

    // Unexpected. The type is interface{} with range: 0..4294967295.
    Unexpected interface{}

    // Unexpected ZLB. The type is interface{} with range: 0..4294967295.
    UnexpectedZlb interface{}
}

func (secondaryDigest *L2tp_Counters_Control_TunnelXr_Authentication_SecondaryDigest) GetEntityData() *types.CommonEntityData {
    secondaryDigest.EntityData.YFilter = secondaryDigest.YFilter
    secondaryDigest.EntityData.YangName = "secondary-digest"
    secondaryDigest.EntityData.BundleName = "cisco_ios_xr"
    secondaryDigest.EntityData.ParentYangName = "authentication"
    secondaryDigest.EntityData.SegmentPath = "secondary-digest"
    secondaryDigest.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    secondaryDigest.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    secondaryDigest.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    secondaryDigest.EntityData.Children = types.NewOrderedMap()
    secondaryDigest.EntityData.Leafs = types.NewOrderedMap()
    secondaryDigest.EntityData.Leafs.Append("validate", types.YLeaf{"Validate", secondaryDigest.Validate})
    secondaryDigest.EntityData.Leafs.Append("bad-hash", types.YLeaf{"BadHash", secondaryDigest.BadHash})
    secondaryDigest.EntityData.Leafs.Append("bad-length", types.YLeaf{"BadLength", secondaryDigest.BadLength})
    secondaryDigest.EntityData.Leafs.Append("ignored", types.YLeaf{"Ignored", secondaryDigest.Ignored})
    secondaryDigest.EntityData.Leafs.Append("missing", types.YLeaf{"Missing", secondaryDigest.Missing})
    secondaryDigest.EntityData.Leafs.Append("passed", types.YLeaf{"Passed", secondaryDigest.Passed})
    secondaryDigest.EntityData.Leafs.Append("failed", types.YLeaf{"Failed", secondaryDigest.Failed})
    secondaryDigest.EntityData.Leafs.Append("skipped", types.YLeaf{"Skipped", secondaryDigest.Skipped})
    secondaryDigest.EntityData.Leafs.Append("generate-response-failures", types.YLeaf{"GenerateResponseFailures", secondaryDigest.GenerateResponseFailures})
    secondaryDigest.EntityData.Leafs.Append("unexpected", types.YLeaf{"Unexpected", secondaryDigest.Unexpected})
    secondaryDigest.EntityData.Leafs.Append("unexpected-zlb", types.YLeaf{"UnexpectedZlb", secondaryDigest.UnexpectedZlb})

    secondaryDigest.EntityData.YListKeys = []string {}

    return &(secondaryDigest.EntityData)
}

// L2tp_Counters_Control_TunnelXr_Authentication_IntegrityCheck
// Integrity check statistics
type L2tp_Counters_Control_TunnelXr_Authentication_IntegrityCheck struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Validate. The type is interface{} with range: 0..4294967295.
    Validate interface{}

    // Bad hash. The type is interface{} with range: 0..4294967295.
    BadHash interface{}

    // Bad length. The type is interface{} with range: 0..4294967295.
    BadLength interface{}

    // Ignored. The type is interface{} with range: 0..4294967295.
    Ignored interface{}

    // Missing. The type is interface{} with range: 0..4294967295.
    Missing interface{}

    // Passed. The type is interface{} with range: 0..4294967295.
    Passed interface{}

    // Failed. The type is interface{} with range: 0..4294967295.
    Failed interface{}

    // Skipped. The type is interface{} with range: 0..4294967295.
    Skipped interface{}

    // Generate response fail. The type is interface{} with range: 0..4294967295.
    GenerateResponseFailures interface{}

    // Unexpected. The type is interface{} with range: 0..4294967295.
    Unexpected interface{}

    // Unexpected ZLB. The type is interface{} with range: 0..4294967295.
    UnexpectedZlb interface{}
}

func (integrityCheck *L2tp_Counters_Control_TunnelXr_Authentication_IntegrityCheck) GetEntityData() *types.CommonEntityData {
    integrityCheck.EntityData.YFilter = integrityCheck.YFilter
    integrityCheck.EntityData.YangName = "integrity-check"
    integrityCheck.EntityData.BundleName = "cisco_ios_xr"
    integrityCheck.EntityData.ParentYangName = "authentication"
    integrityCheck.EntityData.SegmentPath = "integrity-check"
    integrityCheck.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    integrityCheck.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    integrityCheck.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    integrityCheck.EntityData.Children = types.NewOrderedMap()
    integrityCheck.EntityData.Leafs = types.NewOrderedMap()
    integrityCheck.EntityData.Leafs.Append("validate", types.YLeaf{"Validate", integrityCheck.Validate})
    integrityCheck.EntityData.Leafs.Append("bad-hash", types.YLeaf{"BadHash", integrityCheck.BadHash})
    integrityCheck.EntityData.Leafs.Append("bad-length", types.YLeaf{"BadLength", integrityCheck.BadLength})
    integrityCheck.EntityData.Leafs.Append("ignored", types.YLeaf{"Ignored", integrityCheck.Ignored})
    integrityCheck.EntityData.Leafs.Append("missing", types.YLeaf{"Missing", integrityCheck.Missing})
    integrityCheck.EntityData.Leafs.Append("passed", types.YLeaf{"Passed", integrityCheck.Passed})
    integrityCheck.EntityData.Leafs.Append("failed", types.YLeaf{"Failed", integrityCheck.Failed})
    integrityCheck.EntityData.Leafs.Append("skipped", types.YLeaf{"Skipped", integrityCheck.Skipped})
    integrityCheck.EntityData.Leafs.Append("generate-response-failures", types.YLeaf{"GenerateResponseFailures", integrityCheck.GenerateResponseFailures})
    integrityCheck.EntityData.Leafs.Append("unexpected", types.YLeaf{"Unexpected", integrityCheck.Unexpected})
    integrityCheck.EntityData.Leafs.Append("unexpected-zlb", types.YLeaf{"UnexpectedZlb", integrityCheck.UnexpectedZlb})

    integrityCheck.EntityData.YListKeys = []string {}

    return &(integrityCheck.EntityData)
}

// L2tp_Counters_Control_TunnelXr_Authentication_LocalSecret
// Local secret statistics
type L2tp_Counters_Control_TunnelXr_Authentication_LocalSecret struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Validate. The type is interface{} with range: 0..4294967295.
    Validate interface{}

    // Bad hash. The type is interface{} with range: 0..4294967295.
    BadHash interface{}

    // Bad length. The type is interface{} with range: 0..4294967295.
    BadLength interface{}

    // Ignored. The type is interface{} with range: 0..4294967295.
    Ignored interface{}

    // Missing. The type is interface{} with range: 0..4294967295.
    Missing interface{}

    // Passed. The type is interface{} with range: 0..4294967295.
    Passed interface{}

    // Failed. The type is interface{} with range: 0..4294967295.
    Failed interface{}

    // Skipped. The type is interface{} with range: 0..4294967295.
    Skipped interface{}

    // Generate response fail. The type is interface{} with range: 0..4294967295.
    GenerateResponseFailures interface{}

    // Unexpected. The type is interface{} with range: 0..4294967295.
    Unexpected interface{}

    // Unexpected ZLB. The type is interface{} with range: 0..4294967295.
    UnexpectedZlb interface{}
}

func (localSecret *L2tp_Counters_Control_TunnelXr_Authentication_LocalSecret) GetEntityData() *types.CommonEntityData {
    localSecret.EntityData.YFilter = localSecret.YFilter
    localSecret.EntityData.YangName = "local-secret"
    localSecret.EntityData.BundleName = "cisco_ios_xr"
    localSecret.EntityData.ParentYangName = "authentication"
    localSecret.EntityData.SegmentPath = "local-secret"
    localSecret.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localSecret.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localSecret.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localSecret.EntityData.Children = types.NewOrderedMap()
    localSecret.EntityData.Leafs = types.NewOrderedMap()
    localSecret.EntityData.Leafs.Append("validate", types.YLeaf{"Validate", localSecret.Validate})
    localSecret.EntityData.Leafs.Append("bad-hash", types.YLeaf{"BadHash", localSecret.BadHash})
    localSecret.EntityData.Leafs.Append("bad-length", types.YLeaf{"BadLength", localSecret.BadLength})
    localSecret.EntityData.Leafs.Append("ignored", types.YLeaf{"Ignored", localSecret.Ignored})
    localSecret.EntityData.Leafs.Append("missing", types.YLeaf{"Missing", localSecret.Missing})
    localSecret.EntityData.Leafs.Append("passed", types.YLeaf{"Passed", localSecret.Passed})
    localSecret.EntityData.Leafs.Append("failed", types.YLeaf{"Failed", localSecret.Failed})
    localSecret.EntityData.Leafs.Append("skipped", types.YLeaf{"Skipped", localSecret.Skipped})
    localSecret.EntityData.Leafs.Append("generate-response-failures", types.YLeaf{"GenerateResponseFailures", localSecret.GenerateResponseFailures})
    localSecret.EntityData.Leafs.Append("unexpected", types.YLeaf{"Unexpected", localSecret.Unexpected})
    localSecret.EntityData.Leafs.Append("unexpected-zlb", types.YLeaf{"UnexpectedZlb", localSecret.UnexpectedZlb})

    localSecret.EntityData.YListKeys = []string {}

    return &(localSecret.EntityData)
}

// L2tp_Counters_Control_TunnelXr_Authentication_ChallengeAvp
// Challenge AVP statistics
type L2tp_Counters_Control_TunnelXr_Authentication_ChallengeAvp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Validate. The type is interface{} with range: 0..4294967295.
    Validate interface{}

    // Bad hash. The type is interface{} with range: 0..4294967295.
    BadHash interface{}

    // Bad length. The type is interface{} with range: 0..4294967295.
    BadLength interface{}

    // Ignored. The type is interface{} with range: 0..4294967295.
    Ignored interface{}

    // Missing. The type is interface{} with range: 0..4294967295.
    Missing interface{}

    // Passed. The type is interface{} with range: 0..4294967295.
    Passed interface{}

    // Failed. The type is interface{} with range: 0..4294967295.
    Failed interface{}

    // Skipped. The type is interface{} with range: 0..4294967295.
    Skipped interface{}

    // Generate response fail. The type is interface{} with range: 0..4294967295.
    GenerateResponseFailures interface{}

    // Unexpected. The type is interface{} with range: 0..4294967295.
    Unexpected interface{}

    // Unexpected ZLB. The type is interface{} with range: 0..4294967295.
    UnexpectedZlb interface{}
}

func (challengeAvp *L2tp_Counters_Control_TunnelXr_Authentication_ChallengeAvp) GetEntityData() *types.CommonEntityData {
    challengeAvp.EntityData.YFilter = challengeAvp.YFilter
    challengeAvp.EntityData.YangName = "challenge-avp"
    challengeAvp.EntityData.BundleName = "cisco_ios_xr"
    challengeAvp.EntityData.ParentYangName = "authentication"
    challengeAvp.EntityData.SegmentPath = "challenge-avp"
    challengeAvp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    challengeAvp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    challengeAvp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    challengeAvp.EntityData.Children = types.NewOrderedMap()
    challengeAvp.EntityData.Leafs = types.NewOrderedMap()
    challengeAvp.EntityData.Leafs.Append("validate", types.YLeaf{"Validate", challengeAvp.Validate})
    challengeAvp.EntityData.Leafs.Append("bad-hash", types.YLeaf{"BadHash", challengeAvp.BadHash})
    challengeAvp.EntityData.Leafs.Append("bad-length", types.YLeaf{"BadLength", challengeAvp.BadLength})
    challengeAvp.EntityData.Leafs.Append("ignored", types.YLeaf{"Ignored", challengeAvp.Ignored})
    challengeAvp.EntityData.Leafs.Append("missing", types.YLeaf{"Missing", challengeAvp.Missing})
    challengeAvp.EntityData.Leafs.Append("passed", types.YLeaf{"Passed", challengeAvp.Passed})
    challengeAvp.EntityData.Leafs.Append("failed", types.YLeaf{"Failed", challengeAvp.Failed})
    challengeAvp.EntityData.Leafs.Append("skipped", types.YLeaf{"Skipped", challengeAvp.Skipped})
    challengeAvp.EntityData.Leafs.Append("generate-response-failures", types.YLeaf{"GenerateResponseFailures", challengeAvp.GenerateResponseFailures})
    challengeAvp.EntityData.Leafs.Append("unexpected", types.YLeaf{"Unexpected", challengeAvp.Unexpected})
    challengeAvp.EntityData.Leafs.Append("unexpected-zlb", types.YLeaf{"UnexpectedZlb", challengeAvp.UnexpectedZlb})

    challengeAvp.EntityData.YListKeys = []string {}

    return &(challengeAvp.EntityData)
}

// L2tp_Counters_Control_TunnelXr_Authentication_ChallengeReponse
// Challenge response statistics
type L2tp_Counters_Control_TunnelXr_Authentication_ChallengeReponse struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Validate. The type is interface{} with range: 0..4294967295.
    Validate interface{}

    // Bad hash. The type is interface{} with range: 0..4294967295.
    BadHash interface{}

    // Bad length. The type is interface{} with range: 0..4294967295.
    BadLength interface{}

    // Ignored. The type is interface{} with range: 0..4294967295.
    Ignored interface{}

    // Missing. The type is interface{} with range: 0..4294967295.
    Missing interface{}

    // Passed. The type is interface{} with range: 0..4294967295.
    Passed interface{}

    // Failed. The type is interface{} with range: 0..4294967295.
    Failed interface{}

    // Skipped. The type is interface{} with range: 0..4294967295.
    Skipped interface{}

    // Generate response fail. The type is interface{} with range: 0..4294967295.
    GenerateResponseFailures interface{}

    // Unexpected. The type is interface{} with range: 0..4294967295.
    Unexpected interface{}

    // Unexpected ZLB. The type is interface{} with range: 0..4294967295.
    UnexpectedZlb interface{}
}

func (challengeReponse *L2tp_Counters_Control_TunnelXr_Authentication_ChallengeReponse) GetEntityData() *types.CommonEntityData {
    challengeReponse.EntityData.YFilter = challengeReponse.YFilter
    challengeReponse.EntityData.YangName = "challenge-reponse"
    challengeReponse.EntityData.BundleName = "cisco_ios_xr"
    challengeReponse.EntityData.ParentYangName = "authentication"
    challengeReponse.EntityData.SegmentPath = "challenge-reponse"
    challengeReponse.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    challengeReponse.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    challengeReponse.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    challengeReponse.EntityData.Children = types.NewOrderedMap()
    challengeReponse.EntityData.Leafs = types.NewOrderedMap()
    challengeReponse.EntityData.Leafs.Append("validate", types.YLeaf{"Validate", challengeReponse.Validate})
    challengeReponse.EntityData.Leafs.Append("bad-hash", types.YLeaf{"BadHash", challengeReponse.BadHash})
    challengeReponse.EntityData.Leafs.Append("bad-length", types.YLeaf{"BadLength", challengeReponse.BadLength})
    challengeReponse.EntityData.Leafs.Append("ignored", types.YLeaf{"Ignored", challengeReponse.Ignored})
    challengeReponse.EntityData.Leafs.Append("missing", types.YLeaf{"Missing", challengeReponse.Missing})
    challengeReponse.EntityData.Leafs.Append("passed", types.YLeaf{"Passed", challengeReponse.Passed})
    challengeReponse.EntityData.Leafs.Append("failed", types.YLeaf{"Failed", challengeReponse.Failed})
    challengeReponse.EntityData.Leafs.Append("skipped", types.YLeaf{"Skipped", challengeReponse.Skipped})
    challengeReponse.EntityData.Leafs.Append("generate-response-failures", types.YLeaf{"GenerateResponseFailures", challengeReponse.GenerateResponseFailures})
    challengeReponse.EntityData.Leafs.Append("unexpected", types.YLeaf{"Unexpected", challengeReponse.Unexpected})
    challengeReponse.EntityData.Leafs.Append("unexpected-zlb", types.YLeaf{"UnexpectedZlb", challengeReponse.UnexpectedZlb})

    challengeReponse.EntityData.YListKeys = []string {}

    return &(challengeReponse.EntityData)
}

// L2tp_Counters_Control_TunnelXr_Authentication_OverallStatistics
// Overall statistics
type L2tp_Counters_Control_TunnelXr_Authentication_OverallStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Validate. The type is interface{} with range: 0..4294967295.
    Validate interface{}

    // Bad hash. The type is interface{} with range: 0..4294967295.
    BadHash interface{}

    // Bad length. The type is interface{} with range: 0..4294967295.
    BadLength interface{}

    // Ignored. The type is interface{} with range: 0..4294967295.
    Ignored interface{}

    // Missing. The type is interface{} with range: 0..4294967295.
    Missing interface{}

    // Passed. The type is interface{} with range: 0..4294967295.
    Passed interface{}

    // Failed. The type is interface{} with range: 0..4294967295.
    Failed interface{}

    // Skipped. The type is interface{} with range: 0..4294967295.
    Skipped interface{}

    // Generate response fail. The type is interface{} with range: 0..4294967295.
    GenerateResponseFailures interface{}

    // Unexpected. The type is interface{} with range: 0..4294967295.
    Unexpected interface{}

    // Unexpected ZLB. The type is interface{} with range: 0..4294967295.
    UnexpectedZlb interface{}
}

func (overallStatistics *L2tp_Counters_Control_TunnelXr_Authentication_OverallStatistics) GetEntityData() *types.CommonEntityData {
    overallStatistics.EntityData.YFilter = overallStatistics.YFilter
    overallStatistics.EntityData.YangName = "overall-statistics"
    overallStatistics.EntityData.BundleName = "cisco_ios_xr"
    overallStatistics.EntityData.ParentYangName = "authentication"
    overallStatistics.EntityData.SegmentPath = "overall-statistics"
    overallStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    overallStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    overallStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    overallStatistics.EntityData.Children = types.NewOrderedMap()
    overallStatistics.EntityData.Leafs = types.NewOrderedMap()
    overallStatistics.EntityData.Leafs.Append("validate", types.YLeaf{"Validate", overallStatistics.Validate})
    overallStatistics.EntityData.Leafs.Append("bad-hash", types.YLeaf{"BadHash", overallStatistics.BadHash})
    overallStatistics.EntityData.Leafs.Append("bad-length", types.YLeaf{"BadLength", overallStatistics.BadLength})
    overallStatistics.EntityData.Leafs.Append("ignored", types.YLeaf{"Ignored", overallStatistics.Ignored})
    overallStatistics.EntityData.Leafs.Append("missing", types.YLeaf{"Missing", overallStatistics.Missing})
    overallStatistics.EntityData.Leafs.Append("passed", types.YLeaf{"Passed", overallStatistics.Passed})
    overallStatistics.EntityData.Leafs.Append("failed", types.YLeaf{"Failed", overallStatistics.Failed})
    overallStatistics.EntityData.Leafs.Append("skipped", types.YLeaf{"Skipped", overallStatistics.Skipped})
    overallStatistics.EntityData.Leafs.Append("generate-response-failures", types.YLeaf{"GenerateResponseFailures", overallStatistics.GenerateResponseFailures})
    overallStatistics.EntityData.Leafs.Append("unexpected", types.YLeaf{"Unexpected", overallStatistics.Unexpected})
    overallStatistics.EntityData.Leafs.Append("unexpected-zlb", types.YLeaf{"UnexpectedZlb", overallStatistics.UnexpectedZlb})

    overallStatistics.EntityData.YListKeys = []string {}

    return &(overallStatistics.EntityData)
}

// L2tp_Counters_Control_TunnelXr_Global
// Tunnel counters
type L2tp_Counters_Control_TunnelXr_Global struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Total transmit. The type is interface{} with range: 0..4294967295.
    TotalTransmit interface{}

    // Total retransmit. The type is interface{} with range: 0..4294967295.
    TotalRetransmit interface{}

    // Total received. The type is interface{} with range: 0..4294967295.
    TotalReceived interface{}

    // Total drop. The type is interface{} with range: 0..4294967295.
    TotalDrop interface{}

    // Transmit data.
    Transmit L2tp_Counters_Control_TunnelXr_Global_Transmit

    // Re transmit data.
    Retransmit L2tp_Counters_Control_TunnelXr_Global_Retransmit

    // Received data.
    Received L2tp_Counters_Control_TunnelXr_Global_Received

    // Drop data.
    Drop L2tp_Counters_Control_TunnelXr_Global_Drop
}

func (global *L2tp_Counters_Control_TunnelXr_Global) GetEntityData() *types.CommonEntityData {
    global.EntityData.YFilter = global.YFilter
    global.EntityData.YangName = "global"
    global.EntityData.BundleName = "cisco_ios_xr"
    global.EntityData.ParentYangName = "tunnel-xr"
    global.EntityData.SegmentPath = "global"
    global.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    global.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    global.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    global.EntityData.Children = types.NewOrderedMap()
    global.EntityData.Children.Append("transmit", types.YChild{"Transmit", &global.Transmit})
    global.EntityData.Children.Append("retransmit", types.YChild{"Retransmit", &global.Retransmit})
    global.EntityData.Children.Append("received", types.YChild{"Received", &global.Received})
    global.EntityData.Children.Append("drop", types.YChild{"Drop", &global.Drop})
    global.EntityData.Leafs = types.NewOrderedMap()
    global.EntityData.Leafs.Append("total-transmit", types.YLeaf{"TotalTransmit", global.TotalTransmit})
    global.EntityData.Leafs.Append("total-retransmit", types.YLeaf{"TotalRetransmit", global.TotalRetransmit})
    global.EntityData.Leafs.Append("total-received", types.YLeaf{"TotalReceived", global.TotalReceived})
    global.EntityData.Leafs.Append("total-drop", types.YLeaf{"TotalDrop", global.TotalDrop})

    global.EntityData.YListKeys = []string {}

    return &(global.EntityData)
}

// L2tp_Counters_Control_TunnelXr_Global_Transmit
// Transmit data
type L2tp_Counters_Control_TunnelXr_Global_Transmit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Unknown packets. The type is interface{} with range: 0..4294967295.
    UnknownPackets interface{}

    // Zero length body packets. The type is interface{} with range:
    // 0..4294967295.
    ZeroLengthBodyPackets interface{}

    // Start control connection requests. The type is interface{} with range:
    // 0..4294967295.
    StartControlConnectionRequests interface{}

    // Start control connection replies. The type is interface{} with range:
    // 0..4294967295.
    StartControlConnectionReplies interface{}

    // Start control connection notifications. The type is interface{} with range:
    // 0..4294967295.
    StartControlConnectionNotifications interface{}

    // Stop control connection notifications. The type is interface{} with range:
    // 0..4294967295.
    StopControlConnectionNotifications interface{}

    // Keep alive messages. The type is interface{} with range: 0..4294967295.
    HelloPackets interface{}

    // Outgoing call requests. The type is interface{} with range: 0..4294967295.
    OutgoingCallRequests interface{}

    // Outgoing call replies. The type is interface{} with range: 0..4294967295.
    OutgoingCallReplies interface{}

    // Outgoing call connected packets. The type is interface{} with range:
    // 0..4294967295.
    OutgoingCallConnectedPackets interface{}

    // Incoming call requests. The type is interface{} with range: 0..4294967295.
    IncomingCallRequests interface{}

    // Incoming call replies. The type is interface{} with range: 0..4294967295.
    IncomingCallReplies interface{}

    // Incoming call connected packets. The type is interface{} with range:
    // 0..4294967295.
    IncomingCallConnectedPackets interface{}

    // Call disconnect notify packets. The type is interface{} with range:
    // 0..4294967295.
    CallDisconnectNotifyPackets interface{}

    // WAN error notify packets. The type is interface{} with range:
    // 0..4294967295.
    WanErrorNotifyPackets interface{}

    // Set link info packets. The type is interface{} with range: 0..4294967295.
    SetLinkInfoPackets interface{}

    // Service relay request counts. The type is interface{} with range:
    // 0..4294967295.
    ServiceRelayRequests interface{}

    // Service relay reply counts. The type is interface{} with range:
    // 0..4294967295.
    ServiceRelayReplies interface{}

    // Packets acknowledgement. The type is interface{} with range: 0..4294967295.
    AcknowledgementPackets interface{}
}

func (transmit *L2tp_Counters_Control_TunnelXr_Global_Transmit) GetEntityData() *types.CommonEntityData {
    transmit.EntityData.YFilter = transmit.YFilter
    transmit.EntityData.YangName = "transmit"
    transmit.EntityData.BundleName = "cisco_ios_xr"
    transmit.EntityData.ParentYangName = "global"
    transmit.EntityData.SegmentPath = "transmit"
    transmit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transmit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transmit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transmit.EntityData.Children = types.NewOrderedMap()
    transmit.EntityData.Leafs = types.NewOrderedMap()
    transmit.EntityData.Leafs.Append("unknown-packets", types.YLeaf{"UnknownPackets", transmit.UnknownPackets})
    transmit.EntityData.Leafs.Append("zero-length-body-packets", types.YLeaf{"ZeroLengthBodyPackets", transmit.ZeroLengthBodyPackets})
    transmit.EntityData.Leafs.Append("start-control-connection-requests", types.YLeaf{"StartControlConnectionRequests", transmit.StartControlConnectionRequests})
    transmit.EntityData.Leafs.Append("start-control-connection-replies", types.YLeaf{"StartControlConnectionReplies", transmit.StartControlConnectionReplies})
    transmit.EntityData.Leafs.Append("start-control-connection-notifications", types.YLeaf{"StartControlConnectionNotifications", transmit.StartControlConnectionNotifications})
    transmit.EntityData.Leafs.Append("stop-control-connection-notifications", types.YLeaf{"StopControlConnectionNotifications", transmit.StopControlConnectionNotifications})
    transmit.EntityData.Leafs.Append("hello-packets", types.YLeaf{"HelloPackets", transmit.HelloPackets})
    transmit.EntityData.Leafs.Append("outgoing-call-requests", types.YLeaf{"OutgoingCallRequests", transmit.OutgoingCallRequests})
    transmit.EntityData.Leafs.Append("outgoing-call-replies", types.YLeaf{"OutgoingCallReplies", transmit.OutgoingCallReplies})
    transmit.EntityData.Leafs.Append("outgoing-call-connected-packets", types.YLeaf{"OutgoingCallConnectedPackets", transmit.OutgoingCallConnectedPackets})
    transmit.EntityData.Leafs.Append("incoming-call-requests", types.YLeaf{"IncomingCallRequests", transmit.IncomingCallRequests})
    transmit.EntityData.Leafs.Append("incoming-call-replies", types.YLeaf{"IncomingCallReplies", transmit.IncomingCallReplies})
    transmit.EntityData.Leafs.Append("incoming-call-connected-packets", types.YLeaf{"IncomingCallConnectedPackets", transmit.IncomingCallConnectedPackets})
    transmit.EntityData.Leafs.Append("call-disconnect-notify-packets", types.YLeaf{"CallDisconnectNotifyPackets", transmit.CallDisconnectNotifyPackets})
    transmit.EntityData.Leafs.Append("wan-error-notify-packets", types.YLeaf{"WanErrorNotifyPackets", transmit.WanErrorNotifyPackets})
    transmit.EntityData.Leafs.Append("set-link-info-packets", types.YLeaf{"SetLinkInfoPackets", transmit.SetLinkInfoPackets})
    transmit.EntityData.Leafs.Append("service-relay-requests", types.YLeaf{"ServiceRelayRequests", transmit.ServiceRelayRequests})
    transmit.EntityData.Leafs.Append("service-relay-replies", types.YLeaf{"ServiceRelayReplies", transmit.ServiceRelayReplies})
    transmit.EntityData.Leafs.Append("acknowledgement-packets", types.YLeaf{"AcknowledgementPackets", transmit.AcknowledgementPackets})

    transmit.EntityData.YListKeys = []string {}

    return &(transmit.EntityData)
}

// L2tp_Counters_Control_TunnelXr_Global_Retransmit
// Re transmit data
type L2tp_Counters_Control_TunnelXr_Global_Retransmit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Unknown packets. The type is interface{} with range: 0..4294967295.
    UnknownPackets interface{}

    // Zero length body packets. The type is interface{} with range:
    // 0..4294967295.
    ZeroLengthBodyPackets interface{}

    // Start control connection requests. The type is interface{} with range:
    // 0..4294967295.
    StartControlConnectionRequests interface{}

    // Start control connection replies. The type is interface{} with range:
    // 0..4294967295.
    StartControlConnectionReplies interface{}

    // Start control connection notifications. The type is interface{} with range:
    // 0..4294967295.
    StartControlConnectionNotifications interface{}

    // Stop control connection notifications. The type is interface{} with range:
    // 0..4294967295.
    StopControlConnectionNotifications interface{}

    // Keep alive messages. The type is interface{} with range: 0..4294967295.
    HelloPackets interface{}

    // Outgoing call requests. The type is interface{} with range: 0..4294967295.
    OutgoingCallRequests interface{}

    // Outgoing call replies. The type is interface{} with range: 0..4294967295.
    OutgoingCallReplies interface{}

    // Outgoing call connected packets. The type is interface{} with range:
    // 0..4294967295.
    OutgoingCallConnectedPackets interface{}

    // Incoming call requests. The type is interface{} with range: 0..4294967295.
    IncomingCallRequests interface{}

    // Incoming call replies. The type is interface{} with range: 0..4294967295.
    IncomingCallReplies interface{}

    // Incoming call connected packets. The type is interface{} with range:
    // 0..4294967295.
    IncomingCallConnectedPackets interface{}

    // Call disconnect notify packets. The type is interface{} with range:
    // 0..4294967295.
    CallDisconnectNotifyPackets interface{}

    // WAN error notify packets. The type is interface{} with range:
    // 0..4294967295.
    WanErrorNotifyPackets interface{}

    // Set link info packets. The type is interface{} with range: 0..4294967295.
    SetLinkInfoPackets interface{}

    // Service relay request counts. The type is interface{} with range:
    // 0..4294967295.
    ServiceRelayRequests interface{}

    // Service relay reply counts. The type is interface{} with range:
    // 0..4294967295.
    ServiceRelayReplies interface{}

    // Packets acknowledgement. The type is interface{} with range: 0..4294967295.
    AcknowledgementPackets interface{}
}

func (retransmit *L2tp_Counters_Control_TunnelXr_Global_Retransmit) GetEntityData() *types.CommonEntityData {
    retransmit.EntityData.YFilter = retransmit.YFilter
    retransmit.EntityData.YangName = "retransmit"
    retransmit.EntityData.BundleName = "cisco_ios_xr"
    retransmit.EntityData.ParentYangName = "global"
    retransmit.EntityData.SegmentPath = "retransmit"
    retransmit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    retransmit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    retransmit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    retransmit.EntityData.Children = types.NewOrderedMap()
    retransmit.EntityData.Leafs = types.NewOrderedMap()
    retransmit.EntityData.Leafs.Append("unknown-packets", types.YLeaf{"UnknownPackets", retransmit.UnknownPackets})
    retransmit.EntityData.Leafs.Append("zero-length-body-packets", types.YLeaf{"ZeroLengthBodyPackets", retransmit.ZeroLengthBodyPackets})
    retransmit.EntityData.Leafs.Append("start-control-connection-requests", types.YLeaf{"StartControlConnectionRequests", retransmit.StartControlConnectionRequests})
    retransmit.EntityData.Leafs.Append("start-control-connection-replies", types.YLeaf{"StartControlConnectionReplies", retransmit.StartControlConnectionReplies})
    retransmit.EntityData.Leafs.Append("start-control-connection-notifications", types.YLeaf{"StartControlConnectionNotifications", retransmit.StartControlConnectionNotifications})
    retransmit.EntityData.Leafs.Append("stop-control-connection-notifications", types.YLeaf{"StopControlConnectionNotifications", retransmit.StopControlConnectionNotifications})
    retransmit.EntityData.Leafs.Append("hello-packets", types.YLeaf{"HelloPackets", retransmit.HelloPackets})
    retransmit.EntityData.Leafs.Append("outgoing-call-requests", types.YLeaf{"OutgoingCallRequests", retransmit.OutgoingCallRequests})
    retransmit.EntityData.Leafs.Append("outgoing-call-replies", types.YLeaf{"OutgoingCallReplies", retransmit.OutgoingCallReplies})
    retransmit.EntityData.Leafs.Append("outgoing-call-connected-packets", types.YLeaf{"OutgoingCallConnectedPackets", retransmit.OutgoingCallConnectedPackets})
    retransmit.EntityData.Leafs.Append("incoming-call-requests", types.YLeaf{"IncomingCallRequests", retransmit.IncomingCallRequests})
    retransmit.EntityData.Leafs.Append("incoming-call-replies", types.YLeaf{"IncomingCallReplies", retransmit.IncomingCallReplies})
    retransmit.EntityData.Leafs.Append("incoming-call-connected-packets", types.YLeaf{"IncomingCallConnectedPackets", retransmit.IncomingCallConnectedPackets})
    retransmit.EntityData.Leafs.Append("call-disconnect-notify-packets", types.YLeaf{"CallDisconnectNotifyPackets", retransmit.CallDisconnectNotifyPackets})
    retransmit.EntityData.Leafs.Append("wan-error-notify-packets", types.YLeaf{"WanErrorNotifyPackets", retransmit.WanErrorNotifyPackets})
    retransmit.EntityData.Leafs.Append("set-link-info-packets", types.YLeaf{"SetLinkInfoPackets", retransmit.SetLinkInfoPackets})
    retransmit.EntityData.Leafs.Append("service-relay-requests", types.YLeaf{"ServiceRelayRequests", retransmit.ServiceRelayRequests})
    retransmit.EntityData.Leafs.Append("service-relay-replies", types.YLeaf{"ServiceRelayReplies", retransmit.ServiceRelayReplies})
    retransmit.EntityData.Leafs.Append("acknowledgement-packets", types.YLeaf{"AcknowledgementPackets", retransmit.AcknowledgementPackets})

    retransmit.EntityData.YListKeys = []string {}

    return &(retransmit.EntityData)
}

// L2tp_Counters_Control_TunnelXr_Global_Received
// Received data
type L2tp_Counters_Control_TunnelXr_Global_Received struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Unknown packets. The type is interface{} with range: 0..4294967295.
    UnknownPackets interface{}

    // Zero length body packets. The type is interface{} with range:
    // 0..4294967295.
    ZeroLengthBodyPackets interface{}

    // Start control connection requests. The type is interface{} with range:
    // 0..4294967295.
    StartControlConnectionRequests interface{}

    // Start control connection replies. The type is interface{} with range:
    // 0..4294967295.
    StartControlConnectionReplies interface{}

    // Start control connection notifications. The type is interface{} with range:
    // 0..4294967295.
    StartControlConnectionNotifications interface{}

    // Stop control connection notifications. The type is interface{} with range:
    // 0..4294967295.
    StopControlConnectionNotifications interface{}

    // Keep alive messages. The type is interface{} with range: 0..4294967295.
    HelloPackets interface{}

    // Outgoing call requests. The type is interface{} with range: 0..4294967295.
    OutgoingCallRequests interface{}

    // Outgoing call replies. The type is interface{} with range: 0..4294967295.
    OutgoingCallReplies interface{}

    // Outgoing call connected packets. The type is interface{} with range:
    // 0..4294967295.
    OutgoingCallConnectedPackets interface{}

    // Incoming call requests. The type is interface{} with range: 0..4294967295.
    IncomingCallRequests interface{}

    // Incoming call replies. The type is interface{} with range: 0..4294967295.
    IncomingCallReplies interface{}

    // Incoming call connected packets. The type is interface{} with range:
    // 0..4294967295.
    IncomingCallConnectedPackets interface{}

    // Call disconnect notify packets. The type is interface{} with range:
    // 0..4294967295.
    CallDisconnectNotifyPackets interface{}

    // WAN error notify packets. The type is interface{} with range:
    // 0..4294967295.
    WanErrorNotifyPackets interface{}

    // Set link info packets. The type is interface{} with range: 0..4294967295.
    SetLinkInfoPackets interface{}

    // Service relay request counts. The type is interface{} with range:
    // 0..4294967295.
    ServiceRelayRequests interface{}

    // Service relay reply counts. The type is interface{} with range:
    // 0..4294967295.
    ServiceRelayReplies interface{}

    // Packets acknowledgement. The type is interface{} with range: 0..4294967295.
    AcknowledgementPackets interface{}
}

func (received *L2tp_Counters_Control_TunnelXr_Global_Received) GetEntityData() *types.CommonEntityData {
    received.EntityData.YFilter = received.YFilter
    received.EntityData.YangName = "received"
    received.EntityData.BundleName = "cisco_ios_xr"
    received.EntityData.ParentYangName = "global"
    received.EntityData.SegmentPath = "received"
    received.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    received.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    received.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    received.EntityData.Children = types.NewOrderedMap()
    received.EntityData.Leafs = types.NewOrderedMap()
    received.EntityData.Leafs.Append("unknown-packets", types.YLeaf{"UnknownPackets", received.UnknownPackets})
    received.EntityData.Leafs.Append("zero-length-body-packets", types.YLeaf{"ZeroLengthBodyPackets", received.ZeroLengthBodyPackets})
    received.EntityData.Leafs.Append("start-control-connection-requests", types.YLeaf{"StartControlConnectionRequests", received.StartControlConnectionRequests})
    received.EntityData.Leafs.Append("start-control-connection-replies", types.YLeaf{"StartControlConnectionReplies", received.StartControlConnectionReplies})
    received.EntityData.Leafs.Append("start-control-connection-notifications", types.YLeaf{"StartControlConnectionNotifications", received.StartControlConnectionNotifications})
    received.EntityData.Leafs.Append("stop-control-connection-notifications", types.YLeaf{"StopControlConnectionNotifications", received.StopControlConnectionNotifications})
    received.EntityData.Leafs.Append("hello-packets", types.YLeaf{"HelloPackets", received.HelloPackets})
    received.EntityData.Leafs.Append("outgoing-call-requests", types.YLeaf{"OutgoingCallRequests", received.OutgoingCallRequests})
    received.EntityData.Leafs.Append("outgoing-call-replies", types.YLeaf{"OutgoingCallReplies", received.OutgoingCallReplies})
    received.EntityData.Leafs.Append("outgoing-call-connected-packets", types.YLeaf{"OutgoingCallConnectedPackets", received.OutgoingCallConnectedPackets})
    received.EntityData.Leafs.Append("incoming-call-requests", types.YLeaf{"IncomingCallRequests", received.IncomingCallRequests})
    received.EntityData.Leafs.Append("incoming-call-replies", types.YLeaf{"IncomingCallReplies", received.IncomingCallReplies})
    received.EntityData.Leafs.Append("incoming-call-connected-packets", types.YLeaf{"IncomingCallConnectedPackets", received.IncomingCallConnectedPackets})
    received.EntityData.Leafs.Append("call-disconnect-notify-packets", types.YLeaf{"CallDisconnectNotifyPackets", received.CallDisconnectNotifyPackets})
    received.EntityData.Leafs.Append("wan-error-notify-packets", types.YLeaf{"WanErrorNotifyPackets", received.WanErrorNotifyPackets})
    received.EntityData.Leafs.Append("set-link-info-packets", types.YLeaf{"SetLinkInfoPackets", received.SetLinkInfoPackets})
    received.EntityData.Leafs.Append("service-relay-requests", types.YLeaf{"ServiceRelayRequests", received.ServiceRelayRequests})
    received.EntityData.Leafs.Append("service-relay-replies", types.YLeaf{"ServiceRelayReplies", received.ServiceRelayReplies})
    received.EntityData.Leafs.Append("acknowledgement-packets", types.YLeaf{"AcknowledgementPackets", received.AcknowledgementPackets})

    received.EntityData.YListKeys = []string {}

    return &(received.EntityData)
}

// L2tp_Counters_Control_TunnelXr_Global_Drop
// Drop data
type L2tp_Counters_Control_TunnelXr_Global_Drop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Unknown packets. The type is interface{} with range: 0..4294967295.
    UnknownPackets interface{}

    // Zero length body packets. The type is interface{} with range:
    // 0..4294967295.
    ZeroLengthBodyPackets interface{}

    // Start control connection requests. The type is interface{} with range:
    // 0..4294967295.
    StartControlConnectionRequests interface{}

    // Start control connection replies. The type is interface{} with range:
    // 0..4294967295.
    StartControlConnectionReplies interface{}

    // Start control connection notifications. The type is interface{} with range:
    // 0..4294967295.
    StartControlConnectionNotifications interface{}

    // Stop control connection notifications. The type is interface{} with range:
    // 0..4294967295.
    StopControlConnectionNotifications interface{}

    // Keep alive messages. The type is interface{} with range: 0..4294967295.
    HelloPackets interface{}

    // Outgoing call requests. The type is interface{} with range: 0..4294967295.
    OutgoingCallRequests interface{}

    // Outgoing call replies. The type is interface{} with range: 0..4294967295.
    OutgoingCallReplies interface{}

    // Outgoing call connected packets. The type is interface{} with range:
    // 0..4294967295.
    OutgoingCallConnectedPackets interface{}

    // Incoming call requests. The type is interface{} with range: 0..4294967295.
    IncomingCallRequests interface{}

    // Incoming call replies. The type is interface{} with range: 0..4294967295.
    IncomingCallReplies interface{}

    // Incoming call connected packets. The type is interface{} with range:
    // 0..4294967295.
    IncomingCallConnectedPackets interface{}

    // Call disconnect notify packets. The type is interface{} with range:
    // 0..4294967295.
    CallDisconnectNotifyPackets interface{}

    // WAN error notify packets. The type is interface{} with range:
    // 0..4294967295.
    WanErrorNotifyPackets interface{}

    // Set link info packets. The type is interface{} with range: 0..4294967295.
    SetLinkInfoPackets interface{}

    // Service relay request counts. The type is interface{} with range:
    // 0..4294967295.
    ServiceRelayRequests interface{}

    // Service relay reply counts. The type is interface{} with range:
    // 0..4294967295.
    ServiceRelayReplies interface{}

    // Packets acknowledgement. The type is interface{} with range: 0..4294967295.
    AcknowledgementPackets interface{}
}

func (drop *L2tp_Counters_Control_TunnelXr_Global_Drop) GetEntityData() *types.CommonEntityData {
    drop.EntityData.YFilter = drop.YFilter
    drop.EntityData.YangName = "drop"
    drop.EntityData.BundleName = "cisco_ios_xr"
    drop.EntityData.ParentYangName = "global"
    drop.EntityData.SegmentPath = "drop"
    drop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    drop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    drop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    drop.EntityData.Children = types.NewOrderedMap()
    drop.EntityData.Leafs = types.NewOrderedMap()
    drop.EntityData.Leafs.Append("unknown-packets", types.YLeaf{"UnknownPackets", drop.UnknownPackets})
    drop.EntityData.Leafs.Append("zero-length-body-packets", types.YLeaf{"ZeroLengthBodyPackets", drop.ZeroLengthBodyPackets})
    drop.EntityData.Leafs.Append("start-control-connection-requests", types.YLeaf{"StartControlConnectionRequests", drop.StartControlConnectionRequests})
    drop.EntityData.Leafs.Append("start-control-connection-replies", types.YLeaf{"StartControlConnectionReplies", drop.StartControlConnectionReplies})
    drop.EntityData.Leafs.Append("start-control-connection-notifications", types.YLeaf{"StartControlConnectionNotifications", drop.StartControlConnectionNotifications})
    drop.EntityData.Leafs.Append("stop-control-connection-notifications", types.YLeaf{"StopControlConnectionNotifications", drop.StopControlConnectionNotifications})
    drop.EntityData.Leafs.Append("hello-packets", types.YLeaf{"HelloPackets", drop.HelloPackets})
    drop.EntityData.Leafs.Append("outgoing-call-requests", types.YLeaf{"OutgoingCallRequests", drop.OutgoingCallRequests})
    drop.EntityData.Leafs.Append("outgoing-call-replies", types.YLeaf{"OutgoingCallReplies", drop.OutgoingCallReplies})
    drop.EntityData.Leafs.Append("outgoing-call-connected-packets", types.YLeaf{"OutgoingCallConnectedPackets", drop.OutgoingCallConnectedPackets})
    drop.EntityData.Leafs.Append("incoming-call-requests", types.YLeaf{"IncomingCallRequests", drop.IncomingCallRequests})
    drop.EntityData.Leafs.Append("incoming-call-replies", types.YLeaf{"IncomingCallReplies", drop.IncomingCallReplies})
    drop.EntityData.Leafs.Append("incoming-call-connected-packets", types.YLeaf{"IncomingCallConnectedPackets", drop.IncomingCallConnectedPackets})
    drop.EntityData.Leafs.Append("call-disconnect-notify-packets", types.YLeaf{"CallDisconnectNotifyPackets", drop.CallDisconnectNotifyPackets})
    drop.EntityData.Leafs.Append("wan-error-notify-packets", types.YLeaf{"WanErrorNotifyPackets", drop.WanErrorNotifyPackets})
    drop.EntityData.Leafs.Append("set-link-info-packets", types.YLeaf{"SetLinkInfoPackets", drop.SetLinkInfoPackets})
    drop.EntityData.Leafs.Append("service-relay-requests", types.YLeaf{"ServiceRelayRequests", drop.ServiceRelayRequests})
    drop.EntityData.Leafs.Append("service-relay-replies", types.YLeaf{"ServiceRelayReplies", drop.ServiceRelayReplies})
    drop.EntityData.Leafs.Append("acknowledgement-packets", types.YLeaf{"AcknowledgementPackets", drop.AcknowledgementPackets})

    drop.EntityData.YListKeys = []string {}

    return &(drop.EntityData)
}

// L2tp_Counters_Control_Tunnels
// Table of tunnel IDs of control message counters
type L2tp_Counters_Control_Tunnels struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2TP tunnel control message counters. The type is slice of
    // L2tp_Counters_Control_Tunnels_Tunnel.
    Tunnel []*L2tp_Counters_Control_Tunnels_Tunnel
}

func (tunnels *L2tp_Counters_Control_Tunnels) GetEntityData() *types.CommonEntityData {
    tunnels.EntityData.YFilter = tunnels.YFilter
    tunnels.EntityData.YangName = "tunnels"
    tunnels.EntityData.BundleName = "cisco_ios_xr"
    tunnels.EntityData.ParentYangName = "control"
    tunnels.EntityData.SegmentPath = "tunnels"
    tunnels.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnels.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnels.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnels.EntityData.Children = types.NewOrderedMap()
    tunnels.EntityData.Children.Append("tunnel", types.YChild{"Tunnel", nil})
    for i := range tunnels.Tunnel {
        tunnels.EntityData.Children.Append(types.GetSegmentPath(tunnels.Tunnel[i]), types.YChild{"Tunnel", tunnels.Tunnel[i]})
    }
    tunnels.EntityData.Leafs = types.NewOrderedMap()

    tunnels.EntityData.YListKeys = []string {}

    return &(tunnels.EntityData)
}

// L2tp_Counters_Control_Tunnels_Tunnel
// L2TP tunnel control message counters
type L2tp_Counters_Control_Tunnels_Tunnel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. L2TP tunnel ID. The type is interface{} with
    // range: 0..4294967295.
    TunnelId interface{}

    // L2TP control message local and remote addresses.
    Brief L2tp_Counters_Control_Tunnels_Tunnel_Brief

    // Global data.
    Global L2tp_Counters_Control_Tunnels_Tunnel_Global
}

func (tunnel *L2tp_Counters_Control_Tunnels_Tunnel) GetEntityData() *types.CommonEntityData {
    tunnel.EntityData.YFilter = tunnel.YFilter
    tunnel.EntityData.YangName = "tunnel"
    tunnel.EntityData.BundleName = "cisco_ios_xr"
    tunnel.EntityData.ParentYangName = "tunnels"
    tunnel.EntityData.SegmentPath = "tunnel" + types.AddKeyToken(tunnel.TunnelId, "tunnel-id")
    tunnel.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnel.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnel.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnel.EntityData.Children = types.NewOrderedMap()
    tunnel.EntityData.Children.Append("brief", types.YChild{"Brief", &tunnel.Brief})
    tunnel.EntityData.Children.Append("global", types.YChild{"Global", &tunnel.Global})
    tunnel.EntityData.Leafs = types.NewOrderedMap()
    tunnel.EntityData.Leafs.Append("tunnel-id", types.YLeaf{"TunnelId", tunnel.TunnelId})

    tunnel.EntityData.YListKeys = []string {"TunnelId"}

    return &(tunnel.EntityData)
}

// L2tp_Counters_Control_Tunnels_Tunnel_Brief
// L2TP control message local and remote addresses
type L2tp_Counters_Control_Tunnels_Tunnel_Brief struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Remote tunnel ID. The type is interface{} with range: 0..4294967295.
    RemoteTunnelId interface{}

    // Local IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    LocalAddress interface{}

    // Remote IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RemoteAddress interface{}
}

func (brief *L2tp_Counters_Control_Tunnels_Tunnel_Brief) GetEntityData() *types.CommonEntityData {
    brief.EntityData.YFilter = brief.YFilter
    brief.EntityData.YangName = "brief"
    brief.EntityData.BundleName = "cisco_ios_xr"
    brief.EntityData.ParentYangName = "tunnel"
    brief.EntityData.SegmentPath = "brief"
    brief.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    brief.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    brief.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    brief.EntityData.Children = types.NewOrderedMap()
    brief.EntityData.Leafs = types.NewOrderedMap()
    brief.EntityData.Leafs.Append("remote-tunnel-id", types.YLeaf{"RemoteTunnelId", brief.RemoteTunnelId})
    brief.EntityData.Leafs.Append("local-address", types.YLeaf{"LocalAddress", brief.LocalAddress})
    brief.EntityData.Leafs.Append("remote-address", types.YLeaf{"RemoteAddress", brief.RemoteAddress})

    brief.EntityData.YListKeys = []string {}

    return &(brief.EntityData)
}

// L2tp_Counters_Control_Tunnels_Tunnel_Global
// Global data
type L2tp_Counters_Control_Tunnels_Tunnel_Global struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Total transmit. The type is interface{} with range: 0..4294967295.
    TotalTransmit interface{}

    // Total retransmit. The type is interface{} with range: 0..4294967295.
    TotalRetransmit interface{}

    // Total received. The type is interface{} with range: 0..4294967295.
    TotalReceived interface{}

    // Total drop. The type is interface{} with range: 0..4294967295.
    TotalDrop interface{}

    // Transmit data.
    Transmit L2tp_Counters_Control_Tunnels_Tunnel_Global_Transmit

    // Re transmit data.
    Retransmit L2tp_Counters_Control_Tunnels_Tunnel_Global_Retransmit

    // Received data.
    Received L2tp_Counters_Control_Tunnels_Tunnel_Global_Received

    // Drop data.
    Drop L2tp_Counters_Control_Tunnels_Tunnel_Global_Drop
}

func (global *L2tp_Counters_Control_Tunnels_Tunnel_Global) GetEntityData() *types.CommonEntityData {
    global.EntityData.YFilter = global.YFilter
    global.EntityData.YangName = "global"
    global.EntityData.BundleName = "cisco_ios_xr"
    global.EntityData.ParentYangName = "tunnel"
    global.EntityData.SegmentPath = "global"
    global.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    global.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    global.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    global.EntityData.Children = types.NewOrderedMap()
    global.EntityData.Children.Append("transmit", types.YChild{"Transmit", &global.Transmit})
    global.EntityData.Children.Append("retransmit", types.YChild{"Retransmit", &global.Retransmit})
    global.EntityData.Children.Append("received", types.YChild{"Received", &global.Received})
    global.EntityData.Children.Append("drop", types.YChild{"Drop", &global.Drop})
    global.EntityData.Leafs = types.NewOrderedMap()
    global.EntityData.Leafs.Append("total-transmit", types.YLeaf{"TotalTransmit", global.TotalTransmit})
    global.EntityData.Leafs.Append("total-retransmit", types.YLeaf{"TotalRetransmit", global.TotalRetransmit})
    global.EntityData.Leafs.Append("total-received", types.YLeaf{"TotalReceived", global.TotalReceived})
    global.EntityData.Leafs.Append("total-drop", types.YLeaf{"TotalDrop", global.TotalDrop})

    global.EntityData.YListKeys = []string {}

    return &(global.EntityData)
}

// L2tp_Counters_Control_Tunnels_Tunnel_Global_Transmit
// Transmit data
type L2tp_Counters_Control_Tunnels_Tunnel_Global_Transmit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Unknown packets. The type is interface{} with range: 0..4294967295.
    UnknownPackets interface{}

    // Zero length body packets. The type is interface{} with range:
    // 0..4294967295.
    ZeroLengthBodyPackets interface{}

    // Start control connection requests. The type is interface{} with range:
    // 0..4294967295.
    StartControlConnectionRequests interface{}

    // Start control connection replies. The type is interface{} with range:
    // 0..4294967295.
    StartControlConnectionReplies interface{}

    // Start control connection notifications. The type is interface{} with range:
    // 0..4294967295.
    StartControlConnectionNotifications interface{}

    // Stop control connection notifications. The type is interface{} with range:
    // 0..4294967295.
    StopControlConnectionNotifications interface{}

    // Keep alive messages. The type is interface{} with range: 0..4294967295.
    HelloPackets interface{}

    // Outgoing call requests. The type is interface{} with range: 0..4294967295.
    OutgoingCallRequests interface{}

    // Outgoing call replies. The type is interface{} with range: 0..4294967295.
    OutgoingCallReplies interface{}

    // Outgoing call connected packets. The type is interface{} with range:
    // 0..4294967295.
    OutgoingCallConnectedPackets interface{}

    // Incoming call requests. The type is interface{} with range: 0..4294967295.
    IncomingCallRequests interface{}

    // Incoming call replies. The type is interface{} with range: 0..4294967295.
    IncomingCallReplies interface{}

    // Incoming call connected packets. The type is interface{} with range:
    // 0..4294967295.
    IncomingCallConnectedPackets interface{}

    // Call disconnect notify packets. The type is interface{} with range:
    // 0..4294967295.
    CallDisconnectNotifyPackets interface{}

    // WAN error notify packets. The type is interface{} with range:
    // 0..4294967295.
    WanErrorNotifyPackets interface{}

    // Set link info packets. The type is interface{} with range: 0..4294967295.
    SetLinkInfoPackets interface{}

    // Service relay request counts. The type is interface{} with range:
    // 0..4294967295.
    ServiceRelayRequests interface{}

    // Service relay reply counts. The type is interface{} with range:
    // 0..4294967295.
    ServiceRelayReplies interface{}

    // Packets acknowledgement. The type is interface{} with range: 0..4294967295.
    AcknowledgementPackets interface{}
}

func (transmit *L2tp_Counters_Control_Tunnels_Tunnel_Global_Transmit) GetEntityData() *types.CommonEntityData {
    transmit.EntityData.YFilter = transmit.YFilter
    transmit.EntityData.YangName = "transmit"
    transmit.EntityData.BundleName = "cisco_ios_xr"
    transmit.EntityData.ParentYangName = "global"
    transmit.EntityData.SegmentPath = "transmit"
    transmit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transmit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transmit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transmit.EntityData.Children = types.NewOrderedMap()
    transmit.EntityData.Leafs = types.NewOrderedMap()
    transmit.EntityData.Leafs.Append("unknown-packets", types.YLeaf{"UnknownPackets", transmit.UnknownPackets})
    transmit.EntityData.Leafs.Append("zero-length-body-packets", types.YLeaf{"ZeroLengthBodyPackets", transmit.ZeroLengthBodyPackets})
    transmit.EntityData.Leafs.Append("start-control-connection-requests", types.YLeaf{"StartControlConnectionRequests", transmit.StartControlConnectionRequests})
    transmit.EntityData.Leafs.Append("start-control-connection-replies", types.YLeaf{"StartControlConnectionReplies", transmit.StartControlConnectionReplies})
    transmit.EntityData.Leafs.Append("start-control-connection-notifications", types.YLeaf{"StartControlConnectionNotifications", transmit.StartControlConnectionNotifications})
    transmit.EntityData.Leafs.Append("stop-control-connection-notifications", types.YLeaf{"StopControlConnectionNotifications", transmit.StopControlConnectionNotifications})
    transmit.EntityData.Leafs.Append("hello-packets", types.YLeaf{"HelloPackets", transmit.HelloPackets})
    transmit.EntityData.Leafs.Append("outgoing-call-requests", types.YLeaf{"OutgoingCallRequests", transmit.OutgoingCallRequests})
    transmit.EntityData.Leafs.Append("outgoing-call-replies", types.YLeaf{"OutgoingCallReplies", transmit.OutgoingCallReplies})
    transmit.EntityData.Leafs.Append("outgoing-call-connected-packets", types.YLeaf{"OutgoingCallConnectedPackets", transmit.OutgoingCallConnectedPackets})
    transmit.EntityData.Leafs.Append("incoming-call-requests", types.YLeaf{"IncomingCallRequests", transmit.IncomingCallRequests})
    transmit.EntityData.Leafs.Append("incoming-call-replies", types.YLeaf{"IncomingCallReplies", transmit.IncomingCallReplies})
    transmit.EntityData.Leafs.Append("incoming-call-connected-packets", types.YLeaf{"IncomingCallConnectedPackets", transmit.IncomingCallConnectedPackets})
    transmit.EntityData.Leafs.Append("call-disconnect-notify-packets", types.YLeaf{"CallDisconnectNotifyPackets", transmit.CallDisconnectNotifyPackets})
    transmit.EntityData.Leafs.Append("wan-error-notify-packets", types.YLeaf{"WanErrorNotifyPackets", transmit.WanErrorNotifyPackets})
    transmit.EntityData.Leafs.Append("set-link-info-packets", types.YLeaf{"SetLinkInfoPackets", transmit.SetLinkInfoPackets})
    transmit.EntityData.Leafs.Append("service-relay-requests", types.YLeaf{"ServiceRelayRequests", transmit.ServiceRelayRequests})
    transmit.EntityData.Leafs.Append("service-relay-replies", types.YLeaf{"ServiceRelayReplies", transmit.ServiceRelayReplies})
    transmit.EntityData.Leafs.Append("acknowledgement-packets", types.YLeaf{"AcknowledgementPackets", transmit.AcknowledgementPackets})

    transmit.EntityData.YListKeys = []string {}

    return &(transmit.EntityData)
}

// L2tp_Counters_Control_Tunnels_Tunnel_Global_Retransmit
// Re transmit data
type L2tp_Counters_Control_Tunnels_Tunnel_Global_Retransmit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Unknown packets. The type is interface{} with range: 0..4294967295.
    UnknownPackets interface{}

    // Zero length body packets. The type is interface{} with range:
    // 0..4294967295.
    ZeroLengthBodyPackets interface{}

    // Start control connection requests. The type is interface{} with range:
    // 0..4294967295.
    StartControlConnectionRequests interface{}

    // Start control connection replies. The type is interface{} with range:
    // 0..4294967295.
    StartControlConnectionReplies interface{}

    // Start control connection notifications. The type is interface{} with range:
    // 0..4294967295.
    StartControlConnectionNotifications interface{}

    // Stop control connection notifications. The type is interface{} with range:
    // 0..4294967295.
    StopControlConnectionNotifications interface{}

    // Keep alive messages. The type is interface{} with range: 0..4294967295.
    HelloPackets interface{}

    // Outgoing call requests. The type is interface{} with range: 0..4294967295.
    OutgoingCallRequests interface{}

    // Outgoing call replies. The type is interface{} with range: 0..4294967295.
    OutgoingCallReplies interface{}

    // Outgoing call connected packets. The type is interface{} with range:
    // 0..4294967295.
    OutgoingCallConnectedPackets interface{}

    // Incoming call requests. The type is interface{} with range: 0..4294967295.
    IncomingCallRequests interface{}

    // Incoming call replies. The type is interface{} with range: 0..4294967295.
    IncomingCallReplies interface{}

    // Incoming call connected packets. The type is interface{} with range:
    // 0..4294967295.
    IncomingCallConnectedPackets interface{}

    // Call disconnect notify packets. The type is interface{} with range:
    // 0..4294967295.
    CallDisconnectNotifyPackets interface{}

    // WAN error notify packets. The type is interface{} with range:
    // 0..4294967295.
    WanErrorNotifyPackets interface{}

    // Set link info packets. The type is interface{} with range: 0..4294967295.
    SetLinkInfoPackets interface{}

    // Service relay request counts. The type is interface{} with range:
    // 0..4294967295.
    ServiceRelayRequests interface{}

    // Service relay reply counts. The type is interface{} with range:
    // 0..4294967295.
    ServiceRelayReplies interface{}

    // Packets acknowledgement. The type is interface{} with range: 0..4294967295.
    AcknowledgementPackets interface{}
}

func (retransmit *L2tp_Counters_Control_Tunnels_Tunnel_Global_Retransmit) GetEntityData() *types.CommonEntityData {
    retransmit.EntityData.YFilter = retransmit.YFilter
    retransmit.EntityData.YangName = "retransmit"
    retransmit.EntityData.BundleName = "cisco_ios_xr"
    retransmit.EntityData.ParentYangName = "global"
    retransmit.EntityData.SegmentPath = "retransmit"
    retransmit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    retransmit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    retransmit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    retransmit.EntityData.Children = types.NewOrderedMap()
    retransmit.EntityData.Leafs = types.NewOrderedMap()
    retransmit.EntityData.Leafs.Append("unknown-packets", types.YLeaf{"UnknownPackets", retransmit.UnknownPackets})
    retransmit.EntityData.Leafs.Append("zero-length-body-packets", types.YLeaf{"ZeroLengthBodyPackets", retransmit.ZeroLengthBodyPackets})
    retransmit.EntityData.Leafs.Append("start-control-connection-requests", types.YLeaf{"StartControlConnectionRequests", retransmit.StartControlConnectionRequests})
    retransmit.EntityData.Leafs.Append("start-control-connection-replies", types.YLeaf{"StartControlConnectionReplies", retransmit.StartControlConnectionReplies})
    retransmit.EntityData.Leafs.Append("start-control-connection-notifications", types.YLeaf{"StartControlConnectionNotifications", retransmit.StartControlConnectionNotifications})
    retransmit.EntityData.Leafs.Append("stop-control-connection-notifications", types.YLeaf{"StopControlConnectionNotifications", retransmit.StopControlConnectionNotifications})
    retransmit.EntityData.Leafs.Append("hello-packets", types.YLeaf{"HelloPackets", retransmit.HelloPackets})
    retransmit.EntityData.Leafs.Append("outgoing-call-requests", types.YLeaf{"OutgoingCallRequests", retransmit.OutgoingCallRequests})
    retransmit.EntityData.Leafs.Append("outgoing-call-replies", types.YLeaf{"OutgoingCallReplies", retransmit.OutgoingCallReplies})
    retransmit.EntityData.Leafs.Append("outgoing-call-connected-packets", types.YLeaf{"OutgoingCallConnectedPackets", retransmit.OutgoingCallConnectedPackets})
    retransmit.EntityData.Leafs.Append("incoming-call-requests", types.YLeaf{"IncomingCallRequests", retransmit.IncomingCallRequests})
    retransmit.EntityData.Leafs.Append("incoming-call-replies", types.YLeaf{"IncomingCallReplies", retransmit.IncomingCallReplies})
    retransmit.EntityData.Leafs.Append("incoming-call-connected-packets", types.YLeaf{"IncomingCallConnectedPackets", retransmit.IncomingCallConnectedPackets})
    retransmit.EntityData.Leafs.Append("call-disconnect-notify-packets", types.YLeaf{"CallDisconnectNotifyPackets", retransmit.CallDisconnectNotifyPackets})
    retransmit.EntityData.Leafs.Append("wan-error-notify-packets", types.YLeaf{"WanErrorNotifyPackets", retransmit.WanErrorNotifyPackets})
    retransmit.EntityData.Leafs.Append("set-link-info-packets", types.YLeaf{"SetLinkInfoPackets", retransmit.SetLinkInfoPackets})
    retransmit.EntityData.Leafs.Append("service-relay-requests", types.YLeaf{"ServiceRelayRequests", retransmit.ServiceRelayRequests})
    retransmit.EntityData.Leafs.Append("service-relay-replies", types.YLeaf{"ServiceRelayReplies", retransmit.ServiceRelayReplies})
    retransmit.EntityData.Leafs.Append("acknowledgement-packets", types.YLeaf{"AcknowledgementPackets", retransmit.AcknowledgementPackets})

    retransmit.EntityData.YListKeys = []string {}

    return &(retransmit.EntityData)
}

// L2tp_Counters_Control_Tunnels_Tunnel_Global_Received
// Received data
type L2tp_Counters_Control_Tunnels_Tunnel_Global_Received struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Unknown packets. The type is interface{} with range: 0..4294967295.
    UnknownPackets interface{}

    // Zero length body packets. The type is interface{} with range:
    // 0..4294967295.
    ZeroLengthBodyPackets interface{}

    // Start control connection requests. The type is interface{} with range:
    // 0..4294967295.
    StartControlConnectionRequests interface{}

    // Start control connection replies. The type is interface{} with range:
    // 0..4294967295.
    StartControlConnectionReplies interface{}

    // Start control connection notifications. The type is interface{} with range:
    // 0..4294967295.
    StartControlConnectionNotifications interface{}

    // Stop control connection notifications. The type is interface{} with range:
    // 0..4294967295.
    StopControlConnectionNotifications interface{}

    // Keep alive messages. The type is interface{} with range: 0..4294967295.
    HelloPackets interface{}

    // Outgoing call requests. The type is interface{} with range: 0..4294967295.
    OutgoingCallRequests interface{}

    // Outgoing call replies. The type is interface{} with range: 0..4294967295.
    OutgoingCallReplies interface{}

    // Outgoing call connected packets. The type is interface{} with range:
    // 0..4294967295.
    OutgoingCallConnectedPackets interface{}

    // Incoming call requests. The type is interface{} with range: 0..4294967295.
    IncomingCallRequests interface{}

    // Incoming call replies. The type is interface{} with range: 0..4294967295.
    IncomingCallReplies interface{}

    // Incoming call connected packets. The type is interface{} with range:
    // 0..4294967295.
    IncomingCallConnectedPackets interface{}

    // Call disconnect notify packets. The type is interface{} with range:
    // 0..4294967295.
    CallDisconnectNotifyPackets interface{}

    // WAN error notify packets. The type is interface{} with range:
    // 0..4294967295.
    WanErrorNotifyPackets interface{}

    // Set link info packets. The type is interface{} with range: 0..4294967295.
    SetLinkInfoPackets interface{}

    // Service relay request counts. The type is interface{} with range:
    // 0..4294967295.
    ServiceRelayRequests interface{}

    // Service relay reply counts. The type is interface{} with range:
    // 0..4294967295.
    ServiceRelayReplies interface{}

    // Packets acknowledgement. The type is interface{} with range: 0..4294967295.
    AcknowledgementPackets interface{}
}

func (received *L2tp_Counters_Control_Tunnels_Tunnel_Global_Received) GetEntityData() *types.CommonEntityData {
    received.EntityData.YFilter = received.YFilter
    received.EntityData.YangName = "received"
    received.EntityData.BundleName = "cisco_ios_xr"
    received.EntityData.ParentYangName = "global"
    received.EntityData.SegmentPath = "received"
    received.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    received.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    received.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    received.EntityData.Children = types.NewOrderedMap()
    received.EntityData.Leafs = types.NewOrderedMap()
    received.EntityData.Leafs.Append("unknown-packets", types.YLeaf{"UnknownPackets", received.UnknownPackets})
    received.EntityData.Leafs.Append("zero-length-body-packets", types.YLeaf{"ZeroLengthBodyPackets", received.ZeroLengthBodyPackets})
    received.EntityData.Leafs.Append("start-control-connection-requests", types.YLeaf{"StartControlConnectionRequests", received.StartControlConnectionRequests})
    received.EntityData.Leafs.Append("start-control-connection-replies", types.YLeaf{"StartControlConnectionReplies", received.StartControlConnectionReplies})
    received.EntityData.Leafs.Append("start-control-connection-notifications", types.YLeaf{"StartControlConnectionNotifications", received.StartControlConnectionNotifications})
    received.EntityData.Leafs.Append("stop-control-connection-notifications", types.YLeaf{"StopControlConnectionNotifications", received.StopControlConnectionNotifications})
    received.EntityData.Leafs.Append("hello-packets", types.YLeaf{"HelloPackets", received.HelloPackets})
    received.EntityData.Leafs.Append("outgoing-call-requests", types.YLeaf{"OutgoingCallRequests", received.OutgoingCallRequests})
    received.EntityData.Leafs.Append("outgoing-call-replies", types.YLeaf{"OutgoingCallReplies", received.OutgoingCallReplies})
    received.EntityData.Leafs.Append("outgoing-call-connected-packets", types.YLeaf{"OutgoingCallConnectedPackets", received.OutgoingCallConnectedPackets})
    received.EntityData.Leafs.Append("incoming-call-requests", types.YLeaf{"IncomingCallRequests", received.IncomingCallRequests})
    received.EntityData.Leafs.Append("incoming-call-replies", types.YLeaf{"IncomingCallReplies", received.IncomingCallReplies})
    received.EntityData.Leafs.Append("incoming-call-connected-packets", types.YLeaf{"IncomingCallConnectedPackets", received.IncomingCallConnectedPackets})
    received.EntityData.Leafs.Append("call-disconnect-notify-packets", types.YLeaf{"CallDisconnectNotifyPackets", received.CallDisconnectNotifyPackets})
    received.EntityData.Leafs.Append("wan-error-notify-packets", types.YLeaf{"WanErrorNotifyPackets", received.WanErrorNotifyPackets})
    received.EntityData.Leafs.Append("set-link-info-packets", types.YLeaf{"SetLinkInfoPackets", received.SetLinkInfoPackets})
    received.EntityData.Leafs.Append("service-relay-requests", types.YLeaf{"ServiceRelayRequests", received.ServiceRelayRequests})
    received.EntityData.Leafs.Append("service-relay-replies", types.YLeaf{"ServiceRelayReplies", received.ServiceRelayReplies})
    received.EntityData.Leafs.Append("acknowledgement-packets", types.YLeaf{"AcknowledgementPackets", received.AcknowledgementPackets})

    received.EntityData.YListKeys = []string {}

    return &(received.EntityData)
}

// L2tp_Counters_Control_Tunnels_Tunnel_Global_Drop
// Drop data
type L2tp_Counters_Control_Tunnels_Tunnel_Global_Drop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Unknown packets. The type is interface{} with range: 0..4294967295.
    UnknownPackets interface{}

    // Zero length body packets. The type is interface{} with range:
    // 0..4294967295.
    ZeroLengthBodyPackets interface{}

    // Start control connection requests. The type is interface{} with range:
    // 0..4294967295.
    StartControlConnectionRequests interface{}

    // Start control connection replies. The type is interface{} with range:
    // 0..4294967295.
    StartControlConnectionReplies interface{}

    // Start control connection notifications. The type is interface{} with range:
    // 0..4294967295.
    StartControlConnectionNotifications interface{}

    // Stop control connection notifications. The type is interface{} with range:
    // 0..4294967295.
    StopControlConnectionNotifications interface{}

    // Keep alive messages. The type is interface{} with range: 0..4294967295.
    HelloPackets interface{}

    // Outgoing call requests. The type is interface{} with range: 0..4294967295.
    OutgoingCallRequests interface{}

    // Outgoing call replies. The type is interface{} with range: 0..4294967295.
    OutgoingCallReplies interface{}

    // Outgoing call connected packets. The type is interface{} with range:
    // 0..4294967295.
    OutgoingCallConnectedPackets interface{}

    // Incoming call requests. The type is interface{} with range: 0..4294967295.
    IncomingCallRequests interface{}

    // Incoming call replies. The type is interface{} with range: 0..4294967295.
    IncomingCallReplies interface{}

    // Incoming call connected packets. The type is interface{} with range:
    // 0..4294967295.
    IncomingCallConnectedPackets interface{}

    // Call disconnect notify packets. The type is interface{} with range:
    // 0..4294967295.
    CallDisconnectNotifyPackets interface{}

    // WAN error notify packets. The type is interface{} with range:
    // 0..4294967295.
    WanErrorNotifyPackets interface{}

    // Set link info packets. The type is interface{} with range: 0..4294967295.
    SetLinkInfoPackets interface{}

    // Service relay request counts. The type is interface{} with range:
    // 0..4294967295.
    ServiceRelayRequests interface{}

    // Service relay reply counts. The type is interface{} with range:
    // 0..4294967295.
    ServiceRelayReplies interface{}

    // Packets acknowledgement. The type is interface{} with range: 0..4294967295.
    AcknowledgementPackets interface{}
}

func (drop *L2tp_Counters_Control_Tunnels_Tunnel_Global_Drop) GetEntityData() *types.CommonEntityData {
    drop.EntityData.YFilter = drop.YFilter
    drop.EntityData.YangName = "drop"
    drop.EntityData.BundleName = "cisco_ios_xr"
    drop.EntityData.ParentYangName = "global"
    drop.EntityData.SegmentPath = "drop"
    drop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    drop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    drop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    drop.EntityData.Children = types.NewOrderedMap()
    drop.EntityData.Leafs = types.NewOrderedMap()
    drop.EntityData.Leafs.Append("unknown-packets", types.YLeaf{"UnknownPackets", drop.UnknownPackets})
    drop.EntityData.Leafs.Append("zero-length-body-packets", types.YLeaf{"ZeroLengthBodyPackets", drop.ZeroLengthBodyPackets})
    drop.EntityData.Leafs.Append("start-control-connection-requests", types.YLeaf{"StartControlConnectionRequests", drop.StartControlConnectionRequests})
    drop.EntityData.Leafs.Append("start-control-connection-replies", types.YLeaf{"StartControlConnectionReplies", drop.StartControlConnectionReplies})
    drop.EntityData.Leafs.Append("start-control-connection-notifications", types.YLeaf{"StartControlConnectionNotifications", drop.StartControlConnectionNotifications})
    drop.EntityData.Leafs.Append("stop-control-connection-notifications", types.YLeaf{"StopControlConnectionNotifications", drop.StopControlConnectionNotifications})
    drop.EntityData.Leafs.Append("hello-packets", types.YLeaf{"HelloPackets", drop.HelloPackets})
    drop.EntityData.Leafs.Append("outgoing-call-requests", types.YLeaf{"OutgoingCallRequests", drop.OutgoingCallRequests})
    drop.EntityData.Leafs.Append("outgoing-call-replies", types.YLeaf{"OutgoingCallReplies", drop.OutgoingCallReplies})
    drop.EntityData.Leafs.Append("outgoing-call-connected-packets", types.YLeaf{"OutgoingCallConnectedPackets", drop.OutgoingCallConnectedPackets})
    drop.EntityData.Leafs.Append("incoming-call-requests", types.YLeaf{"IncomingCallRequests", drop.IncomingCallRequests})
    drop.EntityData.Leafs.Append("incoming-call-replies", types.YLeaf{"IncomingCallReplies", drop.IncomingCallReplies})
    drop.EntityData.Leafs.Append("incoming-call-connected-packets", types.YLeaf{"IncomingCallConnectedPackets", drop.IncomingCallConnectedPackets})
    drop.EntityData.Leafs.Append("call-disconnect-notify-packets", types.YLeaf{"CallDisconnectNotifyPackets", drop.CallDisconnectNotifyPackets})
    drop.EntityData.Leafs.Append("wan-error-notify-packets", types.YLeaf{"WanErrorNotifyPackets", drop.WanErrorNotifyPackets})
    drop.EntityData.Leafs.Append("set-link-info-packets", types.YLeaf{"SetLinkInfoPackets", drop.SetLinkInfoPackets})
    drop.EntityData.Leafs.Append("service-relay-requests", types.YLeaf{"ServiceRelayRequests", drop.ServiceRelayRequests})
    drop.EntityData.Leafs.Append("service-relay-replies", types.YLeaf{"ServiceRelayReplies", drop.ServiceRelayReplies})
    drop.EntityData.Leafs.Append("acknowledgement-packets", types.YLeaf{"AcknowledgementPackets", drop.AcknowledgementPackets})

    drop.EntityData.YListKeys = []string {}

    return &(drop.EntityData)
}

// L2tp_TunnelConfigurations
// List of tunnel IDs
type L2tp_TunnelConfigurations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2TP tunnel information. The type is slice of
    // L2tp_TunnelConfigurations_TunnelConfiguration.
    TunnelConfiguration []*L2tp_TunnelConfigurations_TunnelConfiguration
}

func (tunnelConfigurations *L2tp_TunnelConfigurations) GetEntityData() *types.CommonEntityData {
    tunnelConfigurations.EntityData.YFilter = tunnelConfigurations.YFilter
    tunnelConfigurations.EntityData.YangName = "tunnel-configurations"
    tunnelConfigurations.EntityData.BundleName = "cisco_ios_xr"
    tunnelConfigurations.EntityData.ParentYangName = "l2tp"
    tunnelConfigurations.EntityData.SegmentPath = "tunnel-configurations"
    tunnelConfigurations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnelConfigurations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnelConfigurations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnelConfigurations.EntityData.Children = types.NewOrderedMap()
    tunnelConfigurations.EntityData.Children.Append("tunnel-configuration", types.YChild{"TunnelConfiguration", nil})
    for i := range tunnelConfigurations.TunnelConfiguration {
        tunnelConfigurations.EntityData.Children.Append(types.GetSegmentPath(tunnelConfigurations.TunnelConfiguration[i]), types.YChild{"TunnelConfiguration", tunnelConfigurations.TunnelConfiguration[i]})
    }
    tunnelConfigurations.EntityData.Leafs = types.NewOrderedMap()

    tunnelConfigurations.EntityData.YListKeys = []string {}

    return &(tunnelConfigurations.EntityData)
}

// L2tp_TunnelConfigurations_TunnelConfiguration
// L2TP tunnel information
type L2tp_TunnelConfigurations_TunnelConfiguration struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Local tunnel ID. The type is interface{} with
    // range: 0..4294967295.
    LocalTunnelId interface{}

    // Remote tunnel ID. The type is interface{} with range: 0..4294967295.
    RemoteTunnelId interface{}

    // L2Tp class data.
    L2tpClass L2tp_TunnelConfigurations_TunnelConfiguration_L2tpClass
}

func (tunnelConfiguration *L2tp_TunnelConfigurations_TunnelConfiguration) GetEntityData() *types.CommonEntityData {
    tunnelConfiguration.EntityData.YFilter = tunnelConfiguration.YFilter
    tunnelConfiguration.EntityData.YangName = "tunnel-configuration"
    tunnelConfiguration.EntityData.BundleName = "cisco_ios_xr"
    tunnelConfiguration.EntityData.ParentYangName = "tunnel-configurations"
    tunnelConfiguration.EntityData.SegmentPath = "tunnel-configuration" + types.AddKeyToken(tunnelConfiguration.LocalTunnelId, "local-tunnel-id")
    tunnelConfiguration.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnelConfiguration.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnelConfiguration.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnelConfiguration.EntityData.Children = types.NewOrderedMap()
    tunnelConfiguration.EntityData.Children.Append("l2tp-class", types.YChild{"L2tpClass", &tunnelConfiguration.L2tpClass})
    tunnelConfiguration.EntityData.Leafs = types.NewOrderedMap()
    tunnelConfiguration.EntityData.Leafs.Append("local-tunnel-id", types.YLeaf{"LocalTunnelId", tunnelConfiguration.LocalTunnelId})
    tunnelConfiguration.EntityData.Leafs.Append("remote-tunnel-id", types.YLeaf{"RemoteTunnelId", tunnelConfiguration.RemoteTunnelId})

    tunnelConfiguration.EntityData.YListKeys = []string {"LocalTunnelId"}

    return &(tunnelConfiguration.EntityData)
}

// L2tp_TunnelConfigurations_TunnelConfiguration_L2tpClass
// L2Tp class data
type L2tp_TunnelConfigurations_TunnelConfiguration_L2tpClass struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP TOS. The type is interface{} with range: 0..255.
    IpTos interface{}

    // VRF name. The type is string with length: 0..256.
    VrfName interface{}

    // Receive window size. The type is interface{} with range: 0..65535.
    ReceiveWindowSize interface{}

    // Class name. The type is string with length: 0..256.
    ClassNameXr interface{}

    // Hash configured as MD5 or SHA1. The type is DigestHash.
    DigestHash interface{}

    // Password. The type is string with length: 0..25.
    Password interface{}

    // Encoded password. The type is string with length: 0..256.
    EncodedPassword interface{}

    // Host name. The type is string with length: 0..256.
    HostName interface{}

    // Accounting List. The type is string with length: 0..256.
    AccountingMethodList interface{}

    // Hello timeout value in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    HelloTimeout interface{}

    // Timeout setup value in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    SetupTimeout interface{}

    // Retransmit minimum timeout in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    RetransmitMinimumTimeout interface{}

    // Retransmit maximum timeout in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    RetransmitMaximumTimeout interface{}

    // Initial timeout minimum in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    InitialRetransmitMinimumTimeout interface{}

    // Initial timeout maximum in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    InitialRetransmitMaximumTimeout interface{}

    // Timeout no user. The type is interface{} with range: 0..4294967295.
    TimeoutNoUser interface{}

    // Retransmit retries. The type is interface{} with range: 0..4294967295.
    RetransmitRetries interface{}

    // Initial retransmit retries. The type is interface{} with range:
    // 0..4294967295.
    InitialRetransmitRetries interface{}

    // True if authentication is enabled. The type is bool.
    IsAuthenticationEnabled interface{}

    // True if class is hidden. The type is bool.
    IsHidden interface{}

    // True if digest authentication is enabled. The type is bool.
    IsDigestEnabled interface{}

    // True if digest check is enabled. The type is bool.
    IsDigestCheckEnabled interface{}

    // True if congestion control is enabled. The type is bool.
    IsCongestionControlEnabled interface{}

    // True if peer address is checked. The type is bool.
    IsPeerAddressChecked interface{}
}

func (l2tpClass *L2tp_TunnelConfigurations_TunnelConfiguration_L2tpClass) GetEntityData() *types.CommonEntityData {
    l2tpClass.EntityData.YFilter = l2tpClass.YFilter
    l2tpClass.EntityData.YangName = "l2tp-class"
    l2tpClass.EntityData.BundleName = "cisco_ios_xr"
    l2tpClass.EntityData.ParentYangName = "tunnel-configuration"
    l2tpClass.EntityData.SegmentPath = "l2tp-class"
    l2tpClass.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    l2tpClass.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    l2tpClass.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    l2tpClass.EntityData.Children = types.NewOrderedMap()
    l2tpClass.EntityData.Leafs = types.NewOrderedMap()
    l2tpClass.EntityData.Leafs.Append("ip-tos", types.YLeaf{"IpTos", l2tpClass.IpTos})
    l2tpClass.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", l2tpClass.VrfName})
    l2tpClass.EntityData.Leafs.Append("receive-window-size", types.YLeaf{"ReceiveWindowSize", l2tpClass.ReceiveWindowSize})
    l2tpClass.EntityData.Leafs.Append("class-name-xr", types.YLeaf{"ClassNameXr", l2tpClass.ClassNameXr})
    l2tpClass.EntityData.Leafs.Append("digest-hash", types.YLeaf{"DigestHash", l2tpClass.DigestHash})
    l2tpClass.EntityData.Leafs.Append("password", types.YLeaf{"Password", l2tpClass.Password})
    l2tpClass.EntityData.Leafs.Append("encoded-password", types.YLeaf{"EncodedPassword", l2tpClass.EncodedPassword})
    l2tpClass.EntityData.Leafs.Append("host-name", types.YLeaf{"HostName", l2tpClass.HostName})
    l2tpClass.EntityData.Leafs.Append("accounting-method-list", types.YLeaf{"AccountingMethodList", l2tpClass.AccountingMethodList})
    l2tpClass.EntityData.Leafs.Append("hello-timeout", types.YLeaf{"HelloTimeout", l2tpClass.HelloTimeout})
    l2tpClass.EntityData.Leafs.Append("setup-timeout", types.YLeaf{"SetupTimeout", l2tpClass.SetupTimeout})
    l2tpClass.EntityData.Leafs.Append("retransmit-minimum-timeout", types.YLeaf{"RetransmitMinimumTimeout", l2tpClass.RetransmitMinimumTimeout})
    l2tpClass.EntityData.Leafs.Append("retransmit-maximum-timeout", types.YLeaf{"RetransmitMaximumTimeout", l2tpClass.RetransmitMaximumTimeout})
    l2tpClass.EntityData.Leafs.Append("initial-retransmit-minimum-timeout", types.YLeaf{"InitialRetransmitMinimumTimeout", l2tpClass.InitialRetransmitMinimumTimeout})
    l2tpClass.EntityData.Leafs.Append("initial-retransmit-maximum-timeout", types.YLeaf{"InitialRetransmitMaximumTimeout", l2tpClass.InitialRetransmitMaximumTimeout})
    l2tpClass.EntityData.Leafs.Append("timeout-no-user", types.YLeaf{"TimeoutNoUser", l2tpClass.TimeoutNoUser})
    l2tpClass.EntityData.Leafs.Append("retransmit-retries", types.YLeaf{"RetransmitRetries", l2tpClass.RetransmitRetries})
    l2tpClass.EntityData.Leafs.Append("initial-retransmit-retries", types.YLeaf{"InitialRetransmitRetries", l2tpClass.InitialRetransmitRetries})
    l2tpClass.EntityData.Leafs.Append("is-authentication-enabled", types.YLeaf{"IsAuthenticationEnabled", l2tpClass.IsAuthenticationEnabled})
    l2tpClass.EntityData.Leafs.Append("is-hidden", types.YLeaf{"IsHidden", l2tpClass.IsHidden})
    l2tpClass.EntityData.Leafs.Append("is-digest-enabled", types.YLeaf{"IsDigestEnabled", l2tpClass.IsDigestEnabled})
    l2tpClass.EntityData.Leafs.Append("is-digest-check-enabled", types.YLeaf{"IsDigestCheckEnabled", l2tpClass.IsDigestCheckEnabled})
    l2tpClass.EntityData.Leafs.Append("is-congestion-control-enabled", types.YLeaf{"IsCongestionControlEnabled", l2tpClass.IsCongestionControlEnabled})
    l2tpClass.EntityData.Leafs.Append("is-peer-address-checked", types.YLeaf{"IsPeerAddressChecked", l2tpClass.IsPeerAddressChecked})

    l2tpClass.EntityData.YListKeys = []string {}

    return &(l2tpClass.EntityData)
}

// L2tp_CounterHistFail
// Failure events leading to disconnection
type L2tp_CounterHistFail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // sesions affected due to timeout. The type is interface{} with range:
    // 0..4294967295.
    SessDownTmout interface{}

    // Send side counters. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    TxCounters interface{}

    // Receive side counters. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    RxCounters interface{}

    // timeout events by packet. The type is slice of
    // L2tp_CounterHistFail_PktTimeout.
    PktTimeout []*L2tp_CounterHistFail_PktTimeout
}

func (counterHistFail *L2tp_CounterHistFail) GetEntityData() *types.CommonEntityData {
    counterHistFail.EntityData.YFilter = counterHistFail.YFilter
    counterHistFail.EntityData.YangName = "counter-hist-fail"
    counterHistFail.EntityData.BundleName = "cisco_ios_xr"
    counterHistFail.EntityData.ParentYangName = "l2tp"
    counterHistFail.EntityData.SegmentPath = "counter-hist-fail"
    counterHistFail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    counterHistFail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    counterHistFail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    counterHistFail.EntityData.Children = types.NewOrderedMap()
    counterHistFail.EntityData.Children.Append("pkt-timeout", types.YChild{"PktTimeout", nil})
    for i := range counterHistFail.PktTimeout {
        counterHistFail.EntityData.Children.Append(types.GetSegmentPath(counterHistFail.PktTimeout[i]), types.YChild{"PktTimeout", counterHistFail.PktTimeout[i]})
    }
    counterHistFail.EntityData.Leafs = types.NewOrderedMap()
    counterHistFail.EntityData.Leafs.Append("sess-down-tmout", types.YLeaf{"SessDownTmout", counterHistFail.SessDownTmout})
    counterHistFail.EntityData.Leafs.Append("tx-counters", types.YLeaf{"TxCounters", counterHistFail.TxCounters})
    counterHistFail.EntityData.Leafs.Append("rx-counters", types.YLeaf{"RxCounters", counterHistFail.RxCounters})

    counterHistFail.EntityData.YListKeys = []string {}

    return &(counterHistFail.EntityData)
}

// L2tp_CounterHistFail_PktTimeout
// timeout events by packet
type L2tp_CounterHistFail_PktTimeout struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    Entry interface{}
}

func (pktTimeout *L2tp_CounterHistFail_PktTimeout) GetEntityData() *types.CommonEntityData {
    pktTimeout.EntityData.YFilter = pktTimeout.YFilter
    pktTimeout.EntityData.YangName = "pkt-timeout"
    pktTimeout.EntityData.BundleName = "cisco_ios_xr"
    pktTimeout.EntityData.ParentYangName = "counter-hist-fail"
    pktTimeout.EntityData.SegmentPath = "pkt-timeout"
    pktTimeout.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pktTimeout.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pktTimeout.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pktTimeout.EntityData.Children = types.NewOrderedMap()
    pktTimeout.EntityData.Leafs = types.NewOrderedMap()
    pktTimeout.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", pktTimeout.Entry})

    pktTimeout.EntityData.YListKeys = []string {}

    return &(pktTimeout.EntityData)
}

// L2tp_Classes
// List of L2TP class names
type L2tp_Classes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2TP class name. The type is slice of L2tp_Classes_Class.
    Class []*L2tp_Classes_Class
}

func (classes *L2tp_Classes) GetEntityData() *types.CommonEntityData {
    classes.EntityData.YFilter = classes.YFilter
    classes.EntityData.YangName = "classes"
    classes.EntityData.BundleName = "cisco_ios_xr"
    classes.EntityData.ParentYangName = "l2tp"
    classes.EntityData.SegmentPath = "classes"
    classes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    classes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    classes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    classes.EntityData.Children = types.NewOrderedMap()
    classes.EntityData.Children.Append("class", types.YChild{"Class", nil})
    for i := range classes.Class {
        classes.EntityData.Children.Append(types.GetSegmentPath(classes.Class[i]), types.YChild{"Class", classes.Class[i]})
    }
    classes.EntityData.Leafs = types.NewOrderedMap()

    classes.EntityData.YListKeys = []string {}

    return &(classes.EntityData)
}

// L2tp_Classes_Class
// L2TP class name
type L2tp_Classes_Class struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. L2TP class name. The type is string with length:
    // 1..31.
    ClassName interface{}

    // IP TOS. The type is interface{} with range: 0..255.
    IpTos interface{}

    // VRF name. The type is string with length: 0..256.
    VrfName interface{}

    // Receive window size. The type is interface{} with range: 0..65535.
    ReceiveWindowSize interface{}

    // Class name. The type is string with length: 0..256.
    ClassNameXr interface{}

    // Hash configured as MD5 or SHA1. The type is DigestHash.
    DigestHash interface{}

    // Password. The type is string with length: 0..25.
    Password interface{}

    // Encoded password. The type is string with length: 0..256.
    EncodedPassword interface{}

    // Host name. The type is string with length: 0..256.
    HostName interface{}

    // Accounting List. The type is string with length: 0..256.
    AccountingMethodList interface{}

    // Hello timeout value in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    HelloTimeout interface{}

    // Timeout setup value in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    SetupTimeout interface{}

    // Retransmit minimum timeout in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    RetransmitMinimumTimeout interface{}

    // Retransmit maximum timeout in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    RetransmitMaximumTimeout interface{}

    // Initial timeout minimum in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    InitialRetransmitMinimumTimeout interface{}

    // Initial timeout maximum in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    InitialRetransmitMaximumTimeout interface{}

    // Timeout no user. The type is interface{} with range: 0..4294967295.
    TimeoutNoUser interface{}

    // Retransmit retries. The type is interface{} with range: 0..4294967295.
    RetransmitRetries interface{}

    // Initial retransmit retries. The type is interface{} with range:
    // 0..4294967295.
    InitialRetransmitRetries interface{}

    // True if authentication is enabled. The type is bool.
    IsAuthenticationEnabled interface{}

    // True if class is hidden. The type is bool.
    IsHidden interface{}

    // True if digest authentication is enabled. The type is bool.
    IsDigestEnabled interface{}

    // True if digest check is enabled. The type is bool.
    IsDigestCheckEnabled interface{}

    // True if congestion control is enabled. The type is bool.
    IsCongestionControlEnabled interface{}

    // True if peer address is checked. The type is bool.
    IsPeerAddressChecked interface{}
}

func (class *L2tp_Classes_Class) GetEntityData() *types.CommonEntityData {
    class.EntityData.YFilter = class.YFilter
    class.EntityData.YangName = "class"
    class.EntityData.BundleName = "cisco_ios_xr"
    class.EntityData.ParentYangName = "classes"
    class.EntityData.SegmentPath = "class" + types.AddKeyToken(class.ClassName, "class-name")
    class.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    class.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    class.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    class.EntityData.Children = types.NewOrderedMap()
    class.EntityData.Leafs = types.NewOrderedMap()
    class.EntityData.Leafs.Append("class-name", types.YLeaf{"ClassName", class.ClassName})
    class.EntityData.Leafs.Append("ip-tos", types.YLeaf{"IpTos", class.IpTos})
    class.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", class.VrfName})
    class.EntityData.Leafs.Append("receive-window-size", types.YLeaf{"ReceiveWindowSize", class.ReceiveWindowSize})
    class.EntityData.Leafs.Append("class-name-xr", types.YLeaf{"ClassNameXr", class.ClassNameXr})
    class.EntityData.Leafs.Append("digest-hash", types.YLeaf{"DigestHash", class.DigestHash})
    class.EntityData.Leafs.Append("password", types.YLeaf{"Password", class.Password})
    class.EntityData.Leafs.Append("encoded-password", types.YLeaf{"EncodedPassword", class.EncodedPassword})
    class.EntityData.Leafs.Append("host-name", types.YLeaf{"HostName", class.HostName})
    class.EntityData.Leafs.Append("accounting-method-list", types.YLeaf{"AccountingMethodList", class.AccountingMethodList})
    class.EntityData.Leafs.Append("hello-timeout", types.YLeaf{"HelloTimeout", class.HelloTimeout})
    class.EntityData.Leafs.Append("setup-timeout", types.YLeaf{"SetupTimeout", class.SetupTimeout})
    class.EntityData.Leafs.Append("retransmit-minimum-timeout", types.YLeaf{"RetransmitMinimumTimeout", class.RetransmitMinimumTimeout})
    class.EntityData.Leafs.Append("retransmit-maximum-timeout", types.YLeaf{"RetransmitMaximumTimeout", class.RetransmitMaximumTimeout})
    class.EntityData.Leafs.Append("initial-retransmit-minimum-timeout", types.YLeaf{"InitialRetransmitMinimumTimeout", class.InitialRetransmitMinimumTimeout})
    class.EntityData.Leafs.Append("initial-retransmit-maximum-timeout", types.YLeaf{"InitialRetransmitMaximumTimeout", class.InitialRetransmitMaximumTimeout})
    class.EntityData.Leafs.Append("timeout-no-user", types.YLeaf{"TimeoutNoUser", class.TimeoutNoUser})
    class.EntityData.Leafs.Append("retransmit-retries", types.YLeaf{"RetransmitRetries", class.RetransmitRetries})
    class.EntityData.Leafs.Append("initial-retransmit-retries", types.YLeaf{"InitialRetransmitRetries", class.InitialRetransmitRetries})
    class.EntityData.Leafs.Append("is-authentication-enabled", types.YLeaf{"IsAuthenticationEnabled", class.IsAuthenticationEnabled})
    class.EntityData.Leafs.Append("is-hidden", types.YLeaf{"IsHidden", class.IsHidden})
    class.EntityData.Leafs.Append("is-digest-enabled", types.YLeaf{"IsDigestEnabled", class.IsDigestEnabled})
    class.EntityData.Leafs.Append("is-digest-check-enabled", types.YLeaf{"IsDigestCheckEnabled", class.IsDigestCheckEnabled})
    class.EntityData.Leafs.Append("is-congestion-control-enabled", types.YLeaf{"IsCongestionControlEnabled", class.IsCongestionControlEnabled})
    class.EntityData.Leafs.Append("is-peer-address-checked", types.YLeaf{"IsPeerAddressChecked", class.IsPeerAddressChecked})

    class.EntityData.YListKeys = []string {"ClassName"}

    return &(class.EntityData)
}

// L2tp_Tunnels
// List of tunnel IDs
type L2tp_Tunnels struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2TP tunnel  information. The type is slice of L2tp_Tunnels_Tunnel.
    Tunnel []*L2tp_Tunnels_Tunnel
}

func (tunnels *L2tp_Tunnels) GetEntityData() *types.CommonEntityData {
    tunnels.EntityData.YFilter = tunnels.YFilter
    tunnels.EntityData.YangName = "tunnels"
    tunnels.EntityData.BundleName = "cisco_ios_xr"
    tunnels.EntityData.ParentYangName = "l2tp"
    tunnels.EntityData.SegmentPath = "tunnels"
    tunnels.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnels.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnels.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnels.EntityData.Children = types.NewOrderedMap()
    tunnels.EntityData.Children.Append("tunnel", types.YChild{"Tunnel", nil})
    for i := range tunnels.Tunnel {
        tunnels.EntityData.Children.Append(types.GetSegmentPath(tunnels.Tunnel[i]), types.YChild{"Tunnel", tunnels.Tunnel[i]})
    }
    tunnels.EntityData.Leafs = types.NewOrderedMap()

    tunnels.EntityData.YListKeys = []string {}

    return &(tunnels.EntityData)
}

// L2tp_Tunnels_Tunnel
// L2TP tunnel  information
type L2tp_Tunnels_Tunnel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Local tunnel ID. The type is interface{} with
    // range: 0..4294967295.
    LocalTunnelId interface{}

    // Local tunnel address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    LocalAddress interface{}

    // Remote tunnel address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RemoteAddress interface{}

    // Local port. The type is interface{} with range: 0..65535.
    LocalPort interface{}

    // Remote port. The type is interface{} with range: 0..65535.
    RemotePort interface{}

    // Protocol. The type is interface{} with range: 0..255.
    Protocol interface{}

    // True if tunnel PMTU checking is enabled. The type is bool.
    IsPmtuEnabled interface{}

    // Remote tunnel ID. The type is interface{} with range: 0..4294967295.
    RemoteTunnelId interface{}

    // Local tunnel name. The type is string with length: 0..256.
    LocalTunnelName interface{}

    // Remote tunnel name. The type is string with length: 0..256.
    RemoteTunnelName interface{}

    // L2TP class name. The type is string with length: 0..256.
    ClassName interface{}

    // Number of active sessions. The type is interface{} with range:
    // 0..4294967295.
    ActiveSessions interface{}

    // Sequence NS. The type is interface{} with range: 0..65535.
    SequenceNs interface{}

    // Sequence NR. The type is interface{} with range: 0..65535.
    SequenceNr interface{}

    // Local window size. The type is interface{} with range: 0..65535.
    LocalWindowSize interface{}

    // Remote window size. The type is interface{} with range: 0..65535.
    RemoteWindowSize interface{}

    // Retransmission time in seconds. The type is interface{} with range:
    // 0..65535. Units are second.
    RetransmissionTime interface{}

    // Maximum retransmission time in seconds. The type is interface{} with range:
    // 0..65535. Units are second.
    MaximumRetransmissionTime interface{}

    // Unsent queue size. The type is interface{} with range: 0..65535.
    UnsentQueueSize interface{}

    // Unsent maximum queue size. The type is interface{} with range: 0..65535.
    UnsentMaximumQueueSize interface{}

    // Resend queue size. The type is interface{} with range: 0..65535.
    ResendQueueSize interface{}

    // Resend maximum queue size. The type is interface{} with range: 0..65535.
    ResendMaximumQueueSize interface{}

    // Order queue size. The type is interface{} with range: 0..65535.
    OrderQueueSize interface{}

    // Current number session packet queue check. The type is interface{} with
    // range: 0..65535.
    PacketQueueCheck interface{}

    // Control message authentication with digest secrets. The type is interface{}
    // with range: 0..65535.
    DigestSecrets interface{}

    // Total resends. The type is interface{} with range: 0..4294967295.
    Resends interface{}

    // Total zero length body acknowledgement. The type is interface{} with range:
    // 0..4294967295.
    ZeroLengthBodyAcknowledgementSent interface{}

    // Total out of order dropped packets. The type is interface{} with range:
    // 0..4294967295.
    TotalOutOfOrderDropPackets interface{}

    // Total out of order reorder packets. The type is interface{} with range:
    // 0..4294967295.
    TotalOutOfOrderReorderPackets interface{}

    // Number of peer authentication failures. The type is interface{} with range:
    // 0..4294967295.
    TotalPeerAuthenticationFailures interface{}

    // True if tunnel is up. The type is bool.
    IsTunnelUp interface{}

    // True if congestion control is enabled else false. The type is bool.
    IsCongestionControlEnabled interface{}

    // Retransmit time distribution in seconds. The type is slice of
    // L2tp_Tunnels_Tunnel_RetransmitTime.
    RetransmitTime []*L2tp_Tunnels_Tunnel_RetransmitTime
}

func (tunnel *L2tp_Tunnels_Tunnel) GetEntityData() *types.CommonEntityData {
    tunnel.EntityData.YFilter = tunnel.YFilter
    tunnel.EntityData.YangName = "tunnel"
    tunnel.EntityData.BundleName = "cisco_ios_xr"
    tunnel.EntityData.ParentYangName = "tunnels"
    tunnel.EntityData.SegmentPath = "tunnel" + types.AddKeyToken(tunnel.LocalTunnelId, "local-tunnel-id")
    tunnel.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnel.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnel.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnel.EntityData.Children = types.NewOrderedMap()
    tunnel.EntityData.Children.Append("retransmit-time", types.YChild{"RetransmitTime", nil})
    for i := range tunnel.RetransmitTime {
        tunnel.EntityData.Children.Append(types.GetSegmentPath(tunnel.RetransmitTime[i]), types.YChild{"RetransmitTime", tunnel.RetransmitTime[i]})
    }
    tunnel.EntityData.Leafs = types.NewOrderedMap()
    tunnel.EntityData.Leafs.Append("local-tunnel-id", types.YLeaf{"LocalTunnelId", tunnel.LocalTunnelId})
    tunnel.EntityData.Leafs.Append("local-address", types.YLeaf{"LocalAddress", tunnel.LocalAddress})
    tunnel.EntityData.Leafs.Append("remote-address", types.YLeaf{"RemoteAddress", tunnel.RemoteAddress})
    tunnel.EntityData.Leafs.Append("local-port", types.YLeaf{"LocalPort", tunnel.LocalPort})
    tunnel.EntityData.Leafs.Append("remote-port", types.YLeaf{"RemotePort", tunnel.RemotePort})
    tunnel.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", tunnel.Protocol})
    tunnel.EntityData.Leafs.Append("is-pmtu-enabled", types.YLeaf{"IsPmtuEnabled", tunnel.IsPmtuEnabled})
    tunnel.EntityData.Leafs.Append("remote-tunnel-id", types.YLeaf{"RemoteTunnelId", tunnel.RemoteTunnelId})
    tunnel.EntityData.Leafs.Append("local-tunnel-name", types.YLeaf{"LocalTunnelName", tunnel.LocalTunnelName})
    tunnel.EntityData.Leafs.Append("remote-tunnel-name", types.YLeaf{"RemoteTunnelName", tunnel.RemoteTunnelName})
    tunnel.EntityData.Leafs.Append("class-name", types.YLeaf{"ClassName", tunnel.ClassName})
    tunnel.EntityData.Leafs.Append("active-sessions", types.YLeaf{"ActiveSessions", tunnel.ActiveSessions})
    tunnel.EntityData.Leafs.Append("sequence-ns", types.YLeaf{"SequenceNs", tunnel.SequenceNs})
    tunnel.EntityData.Leafs.Append("sequence-nr", types.YLeaf{"SequenceNr", tunnel.SequenceNr})
    tunnel.EntityData.Leafs.Append("local-window-size", types.YLeaf{"LocalWindowSize", tunnel.LocalWindowSize})
    tunnel.EntityData.Leafs.Append("remote-window-size", types.YLeaf{"RemoteWindowSize", tunnel.RemoteWindowSize})
    tunnel.EntityData.Leafs.Append("retransmission-time", types.YLeaf{"RetransmissionTime", tunnel.RetransmissionTime})
    tunnel.EntityData.Leafs.Append("maximum-retransmission-time", types.YLeaf{"MaximumRetransmissionTime", tunnel.MaximumRetransmissionTime})
    tunnel.EntityData.Leafs.Append("unsent-queue-size", types.YLeaf{"UnsentQueueSize", tunnel.UnsentQueueSize})
    tunnel.EntityData.Leafs.Append("unsent-maximum-queue-size", types.YLeaf{"UnsentMaximumQueueSize", tunnel.UnsentMaximumQueueSize})
    tunnel.EntityData.Leafs.Append("resend-queue-size", types.YLeaf{"ResendQueueSize", tunnel.ResendQueueSize})
    tunnel.EntityData.Leafs.Append("resend-maximum-queue-size", types.YLeaf{"ResendMaximumQueueSize", tunnel.ResendMaximumQueueSize})
    tunnel.EntityData.Leafs.Append("order-queue-size", types.YLeaf{"OrderQueueSize", tunnel.OrderQueueSize})
    tunnel.EntityData.Leafs.Append("packet-queue-check", types.YLeaf{"PacketQueueCheck", tunnel.PacketQueueCheck})
    tunnel.EntityData.Leafs.Append("digest-secrets", types.YLeaf{"DigestSecrets", tunnel.DigestSecrets})
    tunnel.EntityData.Leafs.Append("resends", types.YLeaf{"Resends", tunnel.Resends})
    tunnel.EntityData.Leafs.Append("zero-length-body-acknowledgement-sent", types.YLeaf{"ZeroLengthBodyAcknowledgementSent", tunnel.ZeroLengthBodyAcknowledgementSent})
    tunnel.EntityData.Leafs.Append("total-out-of-order-drop-packets", types.YLeaf{"TotalOutOfOrderDropPackets", tunnel.TotalOutOfOrderDropPackets})
    tunnel.EntityData.Leafs.Append("total-out-of-order-reorder-packets", types.YLeaf{"TotalOutOfOrderReorderPackets", tunnel.TotalOutOfOrderReorderPackets})
    tunnel.EntityData.Leafs.Append("total-peer-authentication-failures", types.YLeaf{"TotalPeerAuthenticationFailures", tunnel.TotalPeerAuthenticationFailures})
    tunnel.EntityData.Leafs.Append("is-tunnel-up", types.YLeaf{"IsTunnelUp", tunnel.IsTunnelUp})
    tunnel.EntityData.Leafs.Append("is-congestion-control-enabled", types.YLeaf{"IsCongestionControlEnabled", tunnel.IsCongestionControlEnabled})

    tunnel.EntityData.YListKeys = []string {"LocalTunnelId"}

    return &(tunnel.EntityData)
}

// L2tp_Tunnels_Tunnel_RetransmitTime
// Retransmit time distribution in seconds
type L2tp_Tunnels_Tunnel_RetransmitTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..65535. Units are second.
    Entry interface{}
}

func (retransmitTime *L2tp_Tunnels_Tunnel_RetransmitTime) GetEntityData() *types.CommonEntityData {
    retransmitTime.EntityData.YFilter = retransmitTime.YFilter
    retransmitTime.EntityData.YangName = "retransmit-time"
    retransmitTime.EntityData.BundleName = "cisco_ios_xr"
    retransmitTime.EntityData.ParentYangName = "tunnel"
    retransmitTime.EntityData.SegmentPath = "retransmit-time"
    retransmitTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    retransmitTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    retransmitTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    retransmitTime.EntityData.Children = types.NewOrderedMap()
    retransmitTime.EntityData.Leafs = types.NewOrderedMap()
    retransmitTime.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", retransmitTime.Entry})

    retransmitTime.EntityData.YListKeys = []string {}

    return &(retransmitTime.EntityData)
}

// L2tp_Sessions
// List of session IDs
type L2tp_Sessions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2TP information for a particular session. The type is slice of
    // L2tp_Sessions_Session.
    Session []*L2tp_Sessions_Session
}

func (sessions *L2tp_Sessions) GetEntityData() *types.CommonEntityData {
    sessions.EntityData.YFilter = sessions.YFilter
    sessions.EntityData.YangName = "sessions"
    sessions.EntityData.BundleName = "cisco_ios_xr"
    sessions.EntityData.ParentYangName = "l2tp"
    sessions.EntityData.SegmentPath = "sessions"
    sessions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessions.EntityData.Children = types.NewOrderedMap()
    sessions.EntityData.Children.Append("session", types.YChild{"Session", nil})
    for i := range sessions.Session {
        sessions.EntityData.Children.Append(types.GetSegmentPath(sessions.Session[i]), types.YChild{"Session", sessions.Session[i]})
    }
    sessions.EntityData.Leafs = types.NewOrderedMap()

    sessions.EntityData.YListKeys = []string {}

    return &(sessions.EntityData)
}

// L2tp_Sessions_Session
// L2TP information for a particular session
type L2tp_Sessions_Session struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Local tunnel ID. The type is interface{} with
    // range: 0..4294967295.
    LocalTunnelId interface{}

    // This attribute is a key. Local session ID. The type is interface{} with
    // range: 0..4294967295.
    LocalSessionId interface{}

    // Local session IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    LocalIpAddress interface{}

    // Remote session IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RemoteIpAddress interface{}

    // l2tp sh sess udp lport. The type is interface{} with range: 0..65535.
    L2tpShSessUdpLport interface{}

    // l2tp sh sess udp rport. The type is interface{} with range: 0..65535.
    L2tpShSessUdpRport interface{}

    // Protocol. The type is interface{} with range: 0..255.
    Protocol interface{}

    // Remote tunnel ID. The type is interface{} with range: 0..4294967295.
    RemoteTunnelId interface{}

    // Call serial number. The type is interface{} with range: 0..4294967295.
    CallSerialNumber interface{}

    // Local tunnel name. The type is string with length: 0..256.
    LocalTunnelName interface{}

    // Remote tunnel name. The type is string with length: 0..256.
    RemoteTunnelName interface{}

    // Remote session ID. The type is interface{} with range: 0..4294967295.
    RemoteSessionId interface{}

    // l2tp sh sess tie breaker enabled. The type is interface{} with range:
    // 0..255.
    L2tpShSessTieBreakerEnabled interface{}

    // l2tp sh sess tie breaker. The type is interface{} with range:
    // 0..18446744073709551615.
    L2tpShSessTieBreaker interface{}

    // True if session is manual. The type is bool.
    IsSessionManual interface{}

    // True if session is up. The type is bool.
    IsSessionUp interface{}

    // True if UDP checksum enabled. The type is bool.
    IsUdpChecksumEnabled interface{}

    // True if session sequence is on. The type is bool.
    IsSequencingOn interface{}

    // True if session state is established. The type is bool.
    IsSessionStateEstablished interface{}

    // True if session initiated locally. The type is bool.
    IsSessionLocallyInitiated interface{}

    // True if conditional debugging is enabled. The type is bool.
    IsConditionalDebugEnabled interface{}

    // Unique ID. The type is interface{} with range: 0..4294967295.
    UniqueId interface{}

    // Interface name. The type is string with length: 0..256.
    InterfaceName interface{}

    // Session application data.
    SessionApplicationData L2tp_Sessions_Session_SessionApplicationData
}

func (session *L2tp_Sessions_Session) GetEntityData() *types.CommonEntityData {
    session.EntityData.YFilter = session.YFilter
    session.EntityData.YangName = "session"
    session.EntityData.BundleName = "cisco_ios_xr"
    session.EntityData.ParentYangName = "sessions"
    session.EntityData.SegmentPath = "session" + types.AddKeyToken(session.LocalTunnelId, "local-tunnel-id") + types.AddKeyToken(session.LocalSessionId, "local-session-id")
    session.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    session.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    session.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    session.EntityData.Children = types.NewOrderedMap()
    session.EntityData.Children.Append("session-application-data", types.YChild{"SessionApplicationData", &session.SessionApplicationData})
    session.EntityData.Leafs = types.NewOrderedMap()
    session.EntityData.Leafs.Append("local-tunnel-id", types.YLeaf{"LocalTunnelId", session.LocalTunnelId})
    session.EntityData.Leafs.Append("local-session-id", types.YLeaf{"LocalSessionId", session.LocalSessionId})
    session.EntityData.Leafs.Append("local-ip-address", types.YLeaf{"LocalIpAddress", session.LocalIpAddress})
    session.EntityData.Leafs.Append("remote-ip-address", types.YLeaf{"RemoteIpAddress", session.RemoteIpAddress})
    session.EntityData.Leafs.Append("l2tp-sh-sess-udp-lport", types.YLeaf{"L2tpShSessUdpLport", session.L2tpShSessUdpLport})
    session.EntityData.Leafs.Append("l2tp-sh-sess-udp-rport", types.YLeaf{"L2tpShSessUdpRport", session.L2tpShSessUdpRport})
    session.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", session.Protocol})
    session.EntityData.Leafs.Append("remote-tunnel-id", types.YLeaf{"RemoteTunnelId", session.RemoteTunnelId})
    session.EntityData.Leafs.Append("call-serial-number", types.YLeaf{"CallSerialNumber", session.CallSerialNumber})
    session.EntityData.Leafs.Append("local-tunnel-name", types.YLeaf{"LocalTunnelName", session.LocalTunnelName})
    session.EntityData.Leafs.Append("remote-tunnel-name", types.YLeaf{"RemoteTunnelName", session.RemoteTunnelName})
    session.EntityData.Leafs.Append("remote-session-id", types.YLeaf{"RemoteSessionId", session.RemoteSessionId})
    session.EntityData.Leafs.Append("l2tp-sh-sess-tie-breaker-enabled", types.YLeaf{"L2tpShSessTieBreakerEnabled", session.L2tpShSessTieBreakerEnabled})
    session.EntityData.Leafs.Append("l2tp-sh-sess-tie-breaker", types.YLeaf{"L2tpShSessTieBreaker", session.L2tpShSessTieBreaker})
    session.EntityData.Leafs.Append("is-session-manual", types.YLeaf{"IsSessionManual", session.IsSessionManual})
    session.EntityData.Leafs.Append("is-session-up", types.YLeaf{"IsSessionUp", session.IsSessionUp})
    session.EntityData.Leafs.Append("is-udp-checksum-enabled", types.YLeaf{"IsUdpChecksumEnabled", session.IsUdpChecksumEnabled})
    session.EntityData.Leafs.Append("is-sequencing-on", types.YLeaf{"IsSequencingOn", session.IsSequencingOn})
    session.EntityData.Leafs.Append("is-session-state-established", types.YLeaf{"IsSessionStateEstablished", session.IsSessionStateEstablished})
    session.EntityData.Leafs.Append("is-session-locally-initiated", types.YLeaf{"IsSessionLocallyInitiated", session.IsSessionLocallyInitiated})
    session.EntityData.Leafs.Append("is-conditional-debug-enabled", types.YLeaf{"IsConditionalDebugEnabled", session.IsConditionalDebugEnabled})
    session.EntityData.Leafs.Append("unique-id", types.YLeaf{"UniqueId", session.UniqueId})
    session.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", session.InterfaceName})

    session.EntityData.YListKeys = []string {"LocalTunnelId", "LocalSessionId"}

    return &(session.EntityData)
}

// L2tp_Sessions_Session_SessionApplicationData
// Session application data
type L2tp_Sessions_Session_SessionApplicationData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // l2tp sh sess app type. The type is interface{} with range: 0..4294967295.
    L2tpShSessAppType interface{}

    // Xconnect data.
    Xconnect L2tp_Sessions_Session_SessionApplicationData_Xconnect

    // VPDN data.
    Vpdn L2tp_Sessions_Session_SessionApplicationData_Vpdn
}

func (sessionApplicationData *L2tp_Sessions_Session_SessionApplicationData) GetEntityData() *types.CommonEntityData {
    sessionApplicationData.EntityData.YFilter = sessionApplicationData.YFilter
    sessionApplicationData.EntityData.YangName = "session-application-data"
    sessionApplicationData.EntityData.BundleName = "cisco_ios_xr"
    sessionApplicationData.EntityData.ParentYangName = "session"
    sessionApplicationData.EntityData.SegmentPath = "session-application-data"
    sessionApplicationData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionApplicationData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionApplicationData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionApplicationData.EntityData.Children = types.NewOrderedMap()
    sessionApplicationData.EntityData.Children.Append("xconnect", types.YChild{"Xconnect", &sessionApplicationData.Xconnect})
    sessionApplicationData.EntityData.Children.Append("vpdn", types.YChild{"Vpdn", &sessionApplicationData.Vpdn})
    sessionApplicationData.EntityData.Leafs = types.NewOrderedMap()
    sessionApplicationData.EntityData.Leafs.Append("l2tp-sh-sess-app-type", types.YLeaf{"L2tpShSessAppType", sessionApplicationData.L2tpShSessAppType})

    sessionApplicationData.EntityData.YListKeys = []string {}

    return &(sessionApplicationData.EntityData)
}

// L2tp_Sessions_Session_SessionApplicationData_Xconnect
// Xconnect data
type L2tp_Sessions_Session_SessionApplicationData_Xconnect struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Circuit name. The type is string.
    CircuitName interface{}

    // Session VC ID. The type is interface{} with range: 0..4294967295.
    SessionvcId interface{}

    // True if circuit state is up. The type is bool.
    IsCircuitStateUp interface{}

    // True if local circuit state is up. The type is bool.
    IsLocalCircuitStateUp interface{}

    // True if remote circuit state is up. The type is bool.
    IsRemoteCircuitStateUp interface{}

    // IPv6ProtocolTunneling. The type is bool.
    Ipv6ProtocolTunneling interface{}
}

func (xconnect *L2tp_Sessions_Session_SessionApplicationData_Xconnect) GetEntityData() *types.CommonEntityData {
    xconnect.EntityData.YFilter = xconnect.YFilter
    xconnect.EntityData.YangName = "xconnect"
    xconnect.EntityData.BundleName = "cisco_ios_xr"
    xconnect.EntityData.ParentYangName = "session-application-data"
    xconnect.EntityData.SegmentPath = "xconnect"
    xconnect.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    xconnect.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    xconnect.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    xconnect.EntityData.Children = types.NewOrderedMap()
    xconnect.EntityData.Leafs = types.NewOrderedMap()
    xconnect.EntityData.Leafs.Append("circuit-name", types.YLeaf{"CircuitName", xconnect.CircuitName})
    xconnect.EntityData.Leafs.Append("sessionvc-id", types.YLeaf{"SessionvcId", xconnect.SessionvcId})
    xconnect.EntityData.Leafs.Append("is-circuit-state-up", types.YLeaf{"IsCircuitStateUp", xconnect.IsCircuitStateUp})
    xconnect.EntityData.Leafs.Append("is-local-circuit-state-up", types.YLeaf{"IsLocalCircuitStateUp", xconnect.IsLocalCircuitStateUp})
    xconnect.EntityData.Leafs.Append("is-remote-circuit-state-up", types.YLeaf{"IsRemoteCircuitStateUp", xconnect.IsRemoteCircuitStateUp})
    xconnect.EntityData.Leafs.Append("ipv6-protocol-tunneling", types.YLeaf{"Ipv6ProtocolTunneling", xconnect.Ipv6ProtocolTunneling})

    xconnect.EntityData.YListKeys = []string {}

    return &(xconnect.EntityData)
}

// L2tp_Sessions_Session_SessionApplicationData_Vpdn
// VPDN data
type L2tp_Sessions_Session_SessionApplicationData_Vpdn struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Session username. The type is string.
    Username interface{}

    // Interface name. The type is string with pattern: [a-zA-Z0-9._/-]+.
    InterfaceName interface{}
}

func (vpdn *L2tp_Sessions_Session_SessionApplicationData_Vpdn) GetEntityData() *types.CommonEntityData {
    vpdn.EntityData.YFilter = vpdn.YFilter
    vpdn.EntityData.YangName = "vpdn"
    vpdn.EntityData.BundleName = "cisco_ios_xr"
    vpdn.EntityData.ParentYangName = "session-application-data"
    vpdn.EntityData.SegmentPath = "vpdn"
    vpdn.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vpdn.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vpdn.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vpdn.EntityData.Children = types.NewOrderedMap()
    vpdn.EntityData.Leafs = types.NewOrderedMap()
    vpdn.EntityData.Leafs.Append("username", types.YLeaf{"Username", vpdn.Username})
    vpdn.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", vpdn.InterfaceName})

    vpdn.EntityData.YListKeys = []string {}

    return &(vpdn.EntityData)
}

// L2tp_Session
// L2TP control messages counters
type L2tp_Session struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2TP session unavailable  information.
    Unavailable L2tp_Session_Unavailable
}

func (session *L2tp_Session) GetEntityData() *types.CommonEntityData {
    session.EntityData.YFilter = session.YFilter
    session.EntityData.YangName = "session"
    session.EntityData.BundleName = "cisco_ios_xr"
    session.EntityData.ParentYangName = "l2tp"
    session.EntityData.SegmentPath = "session"
    session.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    session.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    session.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    session.EntityData.Children = types.NewOrderedMap()
    session.EntityData.Children.Append("unavailable", types.YChild{"Unavailable", &session.Unavailable})
    session.EntityData.Leafs = types.NewOrderedMap()

    session.EntityData.YListKeys = []string {}

    return &(session.EntityData)
}

// L2tp_Session_Unavailable
// L2TP session unavailable  information
type L2tp_Session_Unavailable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of session ID in hold database. The type is interface{} with range:
    // 0..4294967295.
    SessionsOnHold interface{}
}

func (unavailable *L2tp_Session_Unavailable) GetEntityData() *types.CommonEntityData {
    unavailable.EntityData.YFilter = unavailable.YFilter
    unavailable.EntityData.YangName = "unavailable"
    unavailable.EntityData.BundleName = "cisco_ios_xr"
    unavailable.EntityData.ParentYangName = "session"
    unavailable.EntityData.SegmentPath = "unavailable"
    unavailable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unavailable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unavailable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unavailable.EntityData.Children = types.NewOrderedMap()
    unavailable.EntityData.Leafs = types.NewOrderedMap()
    unavailable.EntityData.Leafs.Append("sessions-on-hold", types.YLeaf{"SessionsOnHold", unavailable.SessionsOnHold})

    unavailable.EntityData.YListKeys = []string {}

    return &(unavailable.EntityData)
}

// L2tpv2
// l2tpv2
type L2tpv2 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2TP control messages counters.
    Counters L2tpv2_Counters

    // L2TP v2 statistics information.
    Statistics L2tpv2_Statistics

    // L2TPv2 tunnel .
    Tunnel L2tpv2_Tunnel

    // List of tunnel IDs.
    TunnelConfigurations L2tpv2_TunnelConfigurations

    // Failure events leading to disconnection.
    CounterHistFail L2tpv2_CounterHistFail

    // List of L2TP class names.
    Classes L2tpv2_Classes

    // List of tunnel IDs.
    Tunnels L2tpv2_Tunnels

    // List of session IDs.
    Sessions L2tpv2_Sessions

    // L2TP control messages counters.
    Session L2tpv2_Session
}

func (l2tpv2 *L2tpv2) GetEntityData() *types.CommonEntityData {
    l2tpv2.EntityData.YFilter = l2tpv2.YFilter
    l2tpv2.EntityData.YangName = "l2tpv2"
    l2tpv2.EntityData.BundleName = "cisco_ios_xr"
    l2tpv2.EntityData.ParentYangName = "Cisco-IOS-XR-tunnel-l2tun-oper"
    l2tpv2.EntityData.SegmentPath = "Cisco-IOS-XR-tunnel-l2tun-oper:l2tpv2"
    l2tpv2.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    l2tpv2.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    l2tpv2.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    l2tpv2.EntityData.Children = types.NewOrderedMap()
    l2tpv2.EntityData.Children.Append("counters", types.YChild{"Counters", &l2tpv2.Counters})
    l2tpv2.EntityData.Children.Append("statistics", types.YChild{"Statistics", &l2tpv2.Statistics})
    l2tpv2.EntityData.Children.Append("tunnel", types.YChild{"Tunnel", &l2tpv2.Tunnel})
    l2tpv2.EntityData.Children.Append("tunnel-configurations", types.YChild{"TunnelConfigurations", &l2tpv2.TunnelConfigurations})
    l2tpv2.EntityData.Children.Append("counter-hist-fail", types.YChild{"CounterHistFail", &l2tpv2.CounterHistFail})
    l2tpv2.EntityData.Children.Append("classes", types.YChild{"Classes", &l2tpv2.Classes})
    l2tpv2.EntityData.Children.Append("tunnels", types.YChild{"Tunnels", &l2tpv2.Tunnels})
    l2tpv2.EntityData.Children.Append("sessions", types.YChild{"Sessions", &l2tpv2.Sessions})
    l2tpv2.EntityData.Children.Append("session", types.YChild{"Session", &l2tpv2.Session})
    l2tpv2.EntityData.Leafs = types.NewOrderedMap()

    l2tpv2.EntityData.YListKeys = []string {}

    return &(l2tpv2.EntityData)
}

// L2tpv2_Counters
// L2TP control messages counters
type L2tpv2_Counters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2TP forwarding messages counters.
    Forwarding L2tpv2_Counters_Forwarding

    // L2TP control messages counters.
    Control L2tpv2_Counters_Control
}

func (counters *L2tpv2_Counters) GetEntityData() *types.CommonEntityData {
    counters.EntityData.YFilter = counters.YFilter
    counters.EntityData.YangName = "counters"
    counters.EntityData.BundleName = "cisco_ios_xr"
    counters.EntityData.ParentYangName = "l2tpv2"
    counters.EntityData.SegmentPath = "counters"
    counters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    counters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    counters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    counters.EntityData.Children = types.NewOrderedMap()
    counters.EntityData.Children.Append("forwarding", types.YChild{"Forwarding", &counters.Forwarding})
    counters.EntityData.Children.Append("control", types.YChild{"Control", &counters.Control})
    counters.EntityData.Leafs = types.NewOrderedMap()

    counters.EntityData.YListKeys = []string {}

    return &(counters.EntityData)
}

// L2tpv2_Counters_Forwarding
// L2TP forwarding messages counters
type L2tpv2_Counters_Forwarding struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of class and session IDs.
    Sessions L2tpv2_Counters_Forwarding_Sessions
}

func (forwarding *L2tpv2_Counters_Forwarding) GetEntityData() *types.CommonEntityData {
    forwarding.EntityData.YFilter = forwarding.YFilter
    forwarding.EntityData.YangName = "forwarding"
    forwarding.EntityData.BundleName = "cisco_ios_xr"
    forwarding.EntityData.ParentYangName = "counters"
    forwarding.EntityData.SegmentPath = "forwarding"
    forwarding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    forwarding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    forwarding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    forwarding.EntityData.Children = types.NewOrderedMap()
    forwarding.EntityData.Children.Append("sessions", types.YChild{"Sessions", &forwarding.Sessions})
    forwarding.EntityData.Leafs = types.NewOrderedMap()

    forwarding.EntityData.YListKeys = []string {}

    return &(forwarding.EntityData)
}

// L2tpv2_Counters_Forwarding_Sessions
// List of class and session IDs
type L2tpv2_Counters_Forwarding_Sessions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2TP information for a particular session. The type is slice of
    // L2tpv2_Counters_Forwarding_Sessions_Session.
    Session []*L2tpv2_Counters_Forwarding_Sessions_Session
}

func (sessions *L2tpv2_Counters_Forwarding_Sessions) GetEntityData() *types.CommonEntityData {
    sessions.EntityData.YFilter = sessions.YFilter
    sessions.EntityData.YangName = "sessions"
    sessions.EntityData.BundleName = "cisco_ios_xr"
    sessions.EntityData.ParentYangName = "forwarding"
    sessions.EntityData.SegmentPath = "sessions"
    sessions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessions.EntityData.Children = types.NewOrderedMap()
    sessions.EntityData.Children.Append("session", types.YChild{"Session", nil})
    for i := range sessions.Session {
        sessions.EntityData.Children.Append(types.GetSegmentPath(sessions.Session[i]), types.YChild{"Session", sessions.Session[i]})
    }
    sessions.EntityData.Leafs = types.NewOrderedMap()

    sessions.EntityData.YListKeys = []string {}

    return &(sessions.EntityData)
}

// L2tpv2_Counters_Forwarding_Sessions_Session
// L2TP information for a particular session
type L2tpv2_Counters_Forwarding_Sessions_Session struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Local tunnel ID. The type is interface{} with
    // range: 0..4294967295.
    TunnelId interface{}

    // This attribute is a key. Local session ID. The type is interface{} with
    // range: 0..4294967295.
    SessionId interface{}

    // Remote session ID. The type is interface{} with range: 0..4294967295.
    RemoteSessionId interface{}

    // Number of packets sent in. The type is interface{} with range:
    // 0..18446744073709551615.
    InPackets interface{}

    // Number of packets sent out. The type is interface{} with range:
    // 0..18446744073709551615.
    OutPackets interface{}

    // Number of bytes sent in. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    InBytes interface{}

    // Number of bytes sent out. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    OutBytes interface{}
}

func (session *L2tpv2_Counters_Forwarding_Sessions_Session) GetEntityData() *types.CommonEntityData {
    session.EntityData.YFilter = session.YFilter
    session.EntityData.YangName = "session"
    session.EntityData.BundleName = "cisco_ios_xr"
    session.EntityData.ParentYangName = "sessions"
    session.EntityData.SegmentPath = "session" + types.AddKeyToken(session.TunnelId, "tunnel-id") + types.AddKeyToken(session.SessionId, "session-id")
    session.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    session.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    session.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    session.EntityData.Children = types.NewOrderedMap()
    session.EntityData.Leafs = types.NewOrderedMap()
    session.EntityData.Leafs.Append("tunnel-id", types.YLeaf{"TunnelId", session.TunnelId})
    session.EntityData.Leafs.Append("session-id", types.YLeaf{"SessionId", session.SessionId})
    session.EntityData.Leafs.Append("remote-session-id", types.YLeaf{"RemoteSessionId", session.RemoteSessionId})
    session.EntityData.Leafs.Append("in-packets", types.YLeaf{"InPackets", session.InPackets})
    session.EntityData.Leafs.Append("out-packets", types.YLeaf{"OutPackets", session.OutPackets})
    session.EntityData.Leafs.Append("in-bytes", types.YLeaf{"InBytes", session.InBytes})
    session.EntityData.Leafs.Append("out-bytes", types.YLeaf{"OutBytes", session.OutBytes})

    session.EntityData.YListKeys = []string {"TunnelId", "SessionId"}

    return &(session.EntityData)
}

// L2tpv2_Counters_Control
// L2TP control messages counters
type L2tpv2_Counters_Control struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2TP control tunnel messages counters.
    TunnelXr L2tpv2_Counters_Control_TunnelXr

    // Table of tunnel IDs of control message counters.
    Tunnels L2tpv2_Counters_Control_Tunnels
}

func (control *L2tpv2_Counters_Control) GetEntityData() *types.CommonEntityData {
    control.EntityData.YFilter = control.YFilter
    control.EntityData.YangName = "control"
    control.EntityData.BundleName = "cisco_ios_xr"
    control.EntityData.ParentYangName = "counters"
    control.EntityData.SegmentPath = "control"
    control.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    control.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    control.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    control.EntityData.Children = types.NewOrderedMap()
    control.EntityData.Children.Append("tunnel-xr", types.YChild{"TunnelXr", &control.TunnelXr})
    control.EntityData.Children.Append("tunnels", types.YChild{"Tunnels", &control.Tunnels})
    control.EntityData.Leafs = types.NewOrderedMap()

    control.EntityData.YListKeys = []string {}

    return &(control.EntityData)
}

// L2tpv2_Counters_Control_TunnelXr
// L2TP control tunnel messages counters
type L2tpv2_Counters_Control_TunnelXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tunnel authentication counters.
    Authentication L2tpv2_Counters_Control_TunnelXr_Authentication

    // Tunnel counters.
    Global L2tpv2_Counters_Control_TunnelXr_Global
}

func (tunnelXr *L2tpv2_Counters_Control_TunnelXr) GetEntityData() *types.CommonEntityData {
    tunnelXr.EntityData.YFilter = tunnelXr.YFilter
    tunnelXr.EntityData.YangName = "tunnel-xr"
    tunnelXr.EntityData.BundleName = "cisco_ios_xr"
    tunnelXr.EntityData.ParentYangName = "control"
    tunnelXr.EntityData.SegmentPath = "tunnel-xr"
    tunnelXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnelXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnelXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnelXr.EntityData.Children = types.NewOrderedMap()
    tunnelXr.EntityData.Children.Append("authentication", types.YChild{"Authentication", &tunnelXr.Authentication})
    tunnelXr.EntityData.Children.Append("global", types.YChild{"Global", &tunnelXr.Global})
    tunnelXr.EntityData.Leafs = types.NewOrderedMap()

    tunnelXr.EntityData.YListKeys = []string {}

    return &(tunnelXr.EntityData)
}

// L2tpv2_Counters_Control_TunnelXr_Authentication
// Tunnel authentication counters
type L2tpv2_Counters_Control_TunnelXr_Authentication struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Nonce AVP statistics.
    NonceAvp L2tpv2_Counters_Control_TunnelXr_Authentication_NonceAvp

    // Common digest statistics.
    CommonDigest L2tpv2_Counters_Control_TunnelXr_Authentication_CommonDigest

    // Primary digest statistics.
    PrimaryDigest L2tpv2_Counters_Control_TunnelXr_Authentication_PrimaryDigest

    // Secondary digest statistics.
    SecondaryDigest L2tpv2_Counters_Control_TunnelXr_Authentication_SecondaryDigest

    // Integrity check statistics.
    IntegrityCheck L2tpv2_Counters_Control_TunnelXr_Authentication_IntegrityCheck

    // Local secret statistics.
    LocalSecret L2tpv2_Counters_Control_TunnelXr_Authentication_LocalSecret

    // Challenge AVP statistics.
    ChallengeAvp L2tpv2_Counters_Control_TunnelXr_Authentication_ChallengeAvp

    // Challenge response statistics.
    ChallengeReponse L2tpv2_Counters_Control_TunnelXr_Authentication_ChallengeReponse

    // Overall statistics.
    OverallStatistics L2tpv2_Counters_Control_TunnelXr_Authentication_OverallStatistics
}

func (authentication *L2tpv2_Counters_Control_TunnelXr_Authentication) GetEntityData() *types.CommonEntityData {
    authentication.EntityData.YFilter = authentication.YFilter
    authentication.EntityData.YangName = "authentication"
    authentication.EntityData.BundleName = "cisco_ios_xr"
    authentication.EntityData.ParentYangName = "tunnel-xr"
    authentication.EntityData.SegmentPath = "authentication"
    authentication.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authentication.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authentication.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authentication.EntityData.Children = types.NewOrderedMap()
    authentication.EntityData.Children.Append("nonce-avp", types.YChild{"NonceAvp", &authentication.NonceAvp})
    authentication.EntityData.Children.Append("common-digest", types.YChild{"CommonDigest", &authentication.CommonDigest})
    authentication.EntityData.Children.Append("primary-digest", types.YChild{"PrimaryDigest", &authentication.PrimaryDigest})
    authentication.EntityData.Children.Append("secondary-digest", types.YChild{"SecondaryDigest", &authentication.SecondaryDigest})
    authentication.EntityData.Children.Append("integrity-check", types.YChild{"IntegrityCheck", &authentication.IntegrityCheck})
    authentication.EntityData.Children.Append("local-secret", types.YChild{"LocalSecret", &authentication.LocalSecret})
    authentication.EntityData.Children.Append("challenge-avp", types.YChild{"ChallengeAvp", &authentication.ChallengeAvp})
    authentication.EntityData.Children.Append("challenge-reponse", types.YChild{"ChallengeReponse", &authentication.ChallengeReponse})
    authentication.EntityData.Children.Append("overall-statistics", types.YChild{"OverallStatistics", &authentication.OverallStatistics})
    authentication.EntityData.Leafs = types.NewOrderedMap()

    authentication.EntityData.YListKeys = []string {}

    return &(authentication.EntityData)
}

// L2tpv2_Counters_Control_TunnelXr_Authentication_NonceAvp
// Nonce AVP statistics
type L2tpv2_Counters_Control_TunnelXr_Authentication_NonceAvp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Validate. The type is interface{} with range: 0..4294967295.
    Validate interface{}

    // Bad hash. The type is interface{} with range: 0..4294967295.
    BadHash interface{}

    // Bad length. The type is interface{} with range: 0..4294967295.
    BadLength interface{}

    // Ignored. The type is interface{} with range: 0..4294967295.
    Ignored interface{}

    // Missing. The type is interface{} with range: 0..4294967295.
    Missing interface{}

    // Passed. The type is interface{} with range: 0..4294967295.
    Passed interface{}

    // Failed. The type is interface{} with range: 0..4294967295.
    Failed interface{}

    // Skipped. The type is interface{} with range: 0..4294967295.
    Skipped interface{}

    // Generate response fail. The type is interface{} with range: 0..4294967295.
    GenerateResponseFailures interface{}

    // Unexpected. The type is interface{} with range: 0..4294967295.
    Unexpected interface{}

    // Unexpected ZLB. The type is interface{} with range: 0..4294967295.
    UnexpectedZlb interface{}
}

func (nonceAvp *L2tpv2_Counters_Control_TunnelXr_Authentication_NonceAvp) GetEntityData() *types.CommonEntityData {
    nonceAvp.EntityData.YFilter = nonceAvp.YFilter
    nonceAvp.EntityData.YangName = "nonce-avp"
    nonceAvp.EntityData.BundleName = "cisco_ios_xr"
    nonceAvp.EntityData.ParentYangName = "authentication"
    nonceAvp.EntityData.SegmentPath = "nonce-avp"
    nonceAvp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nonceAvp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nonceAvp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nonceAvp.EntityData.Children = types.NewOrderedMap()
    nonceAvp.EntityData.Leafs = types.NewOrderedMap()
    nonceAvp.EntityData.Leafs.Append("validate", types.YLeaf{"Validate", nonceAvp.Validate})
    nonceAvp.EntityData.Leafs.Append("bad-hash", types.YLeaf{"BadHash", nonceAvp.BadHash})
    nonceAvp.EntityData.Leafs.Append("bad-length", types.YLeaf{"BadLength", nonceAvp.BadLength})
    nonceAvp.EntityData.Leafs.Append("ignored", types.YLeaf{"Ignored", nonceAvp.Ignored})
    nonceAvp.EntityData.Leafs.Append("missing", types.YLeaf{"Missing", nonceAvp.Missing})
    nonceAvp.EntityData.Leafs.Append("passed", types.YLeaf{"Passed", nonceAvp.Passed})
    nonceAvp.EntityData.Leafs.Append("failed", types.YLeaf{"Failed", nonceAvp.Failed})
    nonceAvp.EntityData.Leafs.Append("skipped", types.YLeaf{"Skipped", nonceAvp.Skipped})
    nonceAvp.EntityData.Leafs.Append("generate-response-failures", types.YLeaf{"GenerateResponseFailures", nonceAvp.GenerateResponseFailures})
    nonceAvp.EntityData.Leafs.Append("unexpected", types.YLeaf{"Unexpected", nonceAvp.Unexpected})
    nonceAvp.EntityData.Leafs.Append("unexpected-zlb", types.YLeaf{"UnexpectedZlb", nonceAvp.UnexpectedZlb})

    nonceAvp.EntityData.YListKeys = []string {}

    return &(nonceAvp.EntityData)
}

// L2tpv2_Counters_Control_TunnelXr_Authentication_CommonDigest
// Common digest statistics
type L2tpv2_Counters_Control_TunnelXr_Authentication_CommonDigest struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Validate. The type is interface{} with range: 0..4294967295.
    Validate interface{}

    // Bad hash. The type is interface{} with range: 0..4294967295.
    BadHash interface{}

    // Bad length. The type is interface{} with range: 0..4294967295.
    BadLength interface{}

    // Ignored. The type is interface{} with range: 0..4294967295.
    Ignored interface{}

    // Missing. The type is interface{} with range: 0..4294967295.
    Missing interface{}

    // Passed. The type is interface{} with range: 0..4294967295.
    Passed interface{}

    // Failed. The type is interface{} with range: 0..4294967295.
    Failed interface{}

    // Skipped. The type is interface{} with range: 0..4294967295.
    Skipped interface{}

    // Generate response fail. The type is interface{} with range: 0..4294967295.
    GenerateResponseFailures interface{}

    // Unexpected. The type is interface{} with range: 0..4294967295.
    Unexpected interface{}

    // Unexpected ZLB. The type is interface{} with range: 0..4294967295.
    UnexpectedZlb interface{}
}

func (commonDigest *L2tpv2_Counters_Control_TunnelXr_Authentication_CommonDigest) GetEntityData() *types.CommonEntityData {
    commonDigest.EntityData.YFilter = commonDigest.YFilter
    commonDigest.EntityData.YangName = "common-digest"
    commonDigest.EntityData.BundleName = "cisco_ios_xr"
    commonDigest.EntityData.ParentYangName = "authentication"
    commonDigest.EntityData.SegmentPath = "common-digest"
    commonDigest.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    commonDigest.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    commonDigest.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    commonDigest.EntityData.Children = types.NewOrderedMap()
    commonDigest.EntityData.Leafs = types.NewOrderedMap()
    commonDigest.EntityData.Leafs.Append("validate", types.YLeaf{"Validate", commonDigest.Validate})
    commonDigest.EntityData.Leafs.Append("bad-hash", types.YLeaf{"BadHash", commonDigest.BadHash})
    commonDigest.EntityData.Leafs.Append("bad-length", types.YLeaf{"BadLength", commonDigest.BadLength})
    commonDigest.EntityData.Leafs.Append("ignored", types.YLeaf{"Ignored", commonDigest.Ignored})
    commonDigest.EntityData.Leafs.Append("missing", types.YLeaf{"Missing", commonDigest.Missing})
    commonDigest.EntityData.Leafs.Append("passed", types.YLeaf{"Passed", commonDigest.Passed})
    commonDigest.EntityData.Leafs.Append("failed", types.YLeaf{"Failed", commonDigest.Failed})
    commonDigest.EntityData.Leafs.Append("skipped", types.YLeaf{"Skipped", commonDigest.Skipped})
    commonDigest.EntityData.Leafs.Append("generate-response-failures", types.YLeaf{"GenerateResponseFailures", commonDigest.GenerateResponseFailures})
    commonDigest.EntityData.Leafs.Append("unexpected", types.YLeaf{"Unexpected", commonDigest.Unexpected})
    commonDigest.EntityData.Leafs.Append("unexpected-zlb", types.YLeaf{"UnexpectedZlb", commonDigest.UnexpectedZlb})

    commonDigest.EntityData.YListKeys = []string {}

    return &(commonDigest.EntityData)
}

// L2tpv2_Counters_Control_TunnelXr_Authentication_PrimaryDigest
// Primary digest statistics
type L2tpv2_Counters_Control_TunnelXr_Authentication_PrimaryDigest struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Validate. The type is interface{} with range: 0..4294967295.
    Validate interface{}

    // Bad hash. The type is interface{} with range: 0..4294967295.
    BadHash interface{}

    // Bad length. The type is interface{} with range: 0..4294967295.
    BadLength interface{}

    // Ignored. The type is interface{} with range: 0..4294967295.
    Ignored interface{}

    // Missing. The type is interface{} with range: 0..4294967295.
    Missing interface{}

    // Passed. The type is interface{} with range: 0..4294967295.
    Passed interface{}

    // Failed. The type is interface{} with range: 0..4294967295.
    Failed interface{}

    // Skipped. The type is interface{} with range: 0..4294967295.
    Skipped interface{}

    // Generate response fail. The type is interface{} with range: 0..4294967295.
    GenerateResponseFailures interface{}

    // Unexpected. The type is interface{} with range: 0..4294967295.
    Unexpected interface{}

    // Unexpected ZLB. The type is interface{} with range: 0..4294967295.
    UnexpectedZlb interface{}
}

func (primaryDigest *L2tpv2_Counters_Control_TunnelXr_Authentication_PrimaryDigest) GetEntityData() *types.CommonEntityData {
    primaryDigest.EntityData.YFilter = primaryDigest.YFilter
    primaryDigest.EntityData.YangName = "primary-digest"
    primaryDigest.EntityData.BundleName = "cisco_ios_xr"
    primaryDigest.EntityData.ParentYangName = "authentication"
    primaryDigest.EntityData.SegmentPath = "primary-digest"
    primaryDigest.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    primaryDigest.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    primaryDigest.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    primaryDigest.EntityData.Children = types.NewOrderedMap()
    primaryDigest.EntityData.Leafs = types.NewOrderedMap()
    primaryDigest.EntityData.Leafs.Append("validate", types.YLeaf{"Validate", primaryDigest.Validate})
    primaryDigest.EntityData.Leafs.Append("bad-hash", types.YLeaf{"BadHash", primaryDigest.BadHash})
    primaryDigest.EntityData.Leafs.Append("bad-length", types.YLeaf{"BadLength", primaryDigest.BadLength})
    primaryDigest.EntityData.Leafs.Append("ignored", types.YLeaf{"Ignored", primaryDigest.Ignored})
    primaryDigest.EntityData.Leafs.Append("missing", types.YLeaf{"Missing", primaryDigest.Missing})
    primaryDigest.EntityData.Leafs.Append("passed", types.YLeaf{"Passed", primaryDigest.Passed})
    primaryDigest.EntityData.Leafs.Append("failed", types.YLeaf{"Failed", primaryDigest.Failed})
    primaryDigest.EntityData.Leafs.Append("skipped", types.YLeaf{"Skipped", primaryDigest.Skipped})
    primaryDigest.EntityData.Leafs.Append("generate-response-failures", types.YLeaf{"GenerateResponseFailures", primaryDigest.GenerateResponseFailures})
    primaryDigest.EntityData.Leafs.Append("unexpected", types.YLeaf{"Unexpected", primaryDigest.Unexpected})
    primaryDigest.EntityData.Leafs.Append("unexpected-zlb", types.YLeaf{"UnexpectedZlb", primaryDigest.UnexpectedZlb})

    primaryDigest.EntityData.YListKeys = []string {}

    return &(primaryDigest.EntityData)
}

// L2tpv2_Counters_Control_TunnelXr_Authentication_SecondaryDigest
// Secondary digest statistics
type L2tpv2_Counters_Control_TunnelXr_Authentication_SecondaryDigest struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Validate. The type is interface{} with range: 0..4294967295.
    Validate interface{}

    // Bad hash. The type is interface{} with range: 0..4294967295.
    BadHash interface{}

    // Bad length. The type is interface{} with range: 0..4294967295.
    BadLength interface{}

    // Ignored. The type is interface{} with range: 0..4294967295.
    Ignored interface{}

    // Missing. The type is interface{} with range: 0..4294967295.
    Missing interface{}

    // Passed. The type is interface{} with range: 0..4294967295.
    Passed interface{}

    // Failed. The type is interface{} with range: 0..4294967295.
    Failed interface{}

    // Skipped. The type is interface{} with range: 0..4294967295.
    Skipped interface{}

    // Generate response fail. The type is interface{} with range: 0..4294967295.
    GenerateResponseFailures interface{}

    // Unexpected. The type is interface{} with range: 0..4294967295.
    Unexpected interface{}

    // Unexpected ZLB. The type is interface{} with range: 0..4294967295.
    UnexpectedZlb interface{}
}

func (secondaryDigest *L2tpv2_Counters_Control_TunnelXr_Authentication_SecondaryDigest) GetEntityData() *types.CommonEntityData {
    secondaryDigest.EntityData.YFilter = secondaryDigest.YFilter
    secondaryDigest.EntityData.YangName = "secondary-digest"
    secondaryDigest.EntityData.BundleName = "cisco_ios_xr"
    secondaryDigest.EntityData.ParentYangName = "authentication"
    secondaryDigest.EntityData.SegmentPath = "secondary-digest"
    secondaryDigest.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    secondaryDigest.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    secondaryDigest.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    secondaryDigest.EntityData.Children = types.NewOrderedMap()
    secondaryDigest.EntityData.Leafs = types.NewOrderedMap()
    secondaryDigest.EntityData.Leafs.Append("validate", types.YLeaf{"Validate", secondaryDigest.Validate})
    secondaryDigest.EntityData.Leafs.Append("bad-hash", types.YLeaf{"BadHash", secondaryDigest.BadHash})
    secondaryDigest.EntityData.Leafs.Append("bad-length", types.YLeaf{"BadLength", secondaryDigest.BadLength})
    secondaryDigest.EntityData.Leafs.Append("ignored", types.YLeaf{"Ignored", secondaryDigest.Ignored})
    secondaryDigest.EntityData.Leafs.Append("missing", types.YLeaf{"Missing", secondaryDigest.Missing})
    secondaryDigest.EntityData.Leafs.Append("passed", types.YLeaf{"Passed", secondaryDigest.Passed})
    secondaryDigest.EntityData.Leafs.Append("failed", types.YLeaf{"Failed", secondaryDigest.Failed})
    secondaryDigest.EntityData.Leafs.Append("skipped", types.YLeaf{"Skipped", secondaryDigest.Skipped})
    secondaryDigest.EntityData.Leafs.Append("generate-response-failures", types.YLeaf{"GenerateResponseFailures", secondaryDigest.GenerateResponseFailures})
    secondaryDigest.EntityData.Leafs.Append("unexpected", types.YLeaf{"Unexpected", secondaryDigest.Unexpected})
    secondaryDigest.EntityData.Leafs.Append("unexpected-zlb", types.YLeaf{"UnexpectedZlb", secondaryDigest.UnexpectedZlb})

    secondaryDigest.EntityData.YListKeys = []string {}

    return &(secondaryDigest.EntityData)
}

// L2tpv2_Counters_Control_TunnelXr_Authentication_IntegrityCheck
// Integrity check statistics
type L2tpv2_Counters_Control_TunnelXr_Authentication_IntegrityCheck struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Validate. The type is interface{} with range: 0..4294967295.
    Validate interface{}

    // Bad hash. The type is interface{} with range: 0..4294967295.
    BadHash interface{}

    // Bad length. The type is interface{} with range: 0..4294967295.
    BadLength interface{}

    // Ignored. The type is interface{} with range: 0..4294967295.
    Ignored interface{}

    // Missing. The type is interface{} with range: 0..4294967295.
    Missing interface{}

    // Passed. The type is interface{} with range: 0..4294967295.
    Passed interface{}

    // Failed. The type is interface{} with range: 0..4294967295.
    Failed interface{}

    // Skipped. The type is interface{} with range: 0..4294967295.
    Skipped interface{}

    // Generate response fail. The type is interface{} with range: 0..4294967295.
    GenerateResponseFailures interface{}

    // Unexpected. The type is interface{} with range: 0..4294967295.
    Unexpected interface{}

    // Unexpected ZLB. The type is interface{} with range: 0..4294967295.
    UnexpectedZlb interface{}
}

func (integrityCheck *L2tpv2_Counters_Control_TunnelXr_Authentication_IntegrityCheck) GetEntityData() *types.CommonEntityData {
    integrityCheck.EntityData.YFilter = integrityCheck.YFilter
    integrityCheck.EntityData.YangName = "integrity-check"
    integrityCheck.EntityData.BundleName = "cisco_ios_xr"
    integrityCheck.EntityData.ParentYangName = "authentication"
    integrityCheck.EntityData.SegmentPath = "integrity-check"
    integrityCheck.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    integrityCheck.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    integrityCheck.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    integrityCheck.EntityData.Children = types.NewOrderedMap()
    integrityCheck.EntityData.Leafs = types.NewOrderedMap()
    integrityCheck.EntityData.Leafs.Append("validate", types.YLeaf{"Validate", integrityCheck.Validate})
    integrityCheck.EntityData.Leafs.Append("bad-hash", types.YLeaf{"BadHash", integrityCheck.BadHash})
    integrityCheck.EntityData.Leafs.Append("bad-length", types.YLeaf{"BadLength", integrityCheck.BadLength})
    integrityCheck.EntityData.Leafs.Append("ignored", types.YLeaf{"Ignored", integrityCheck.Ignored})
    integrityCheck.EntityData.Leafs.Append("missing", types.YLeaf{"Missing", integrityCheck.Missing})
    integrityCheck.EntityData.Leafs.Append("passed", types.YLeaf{"Passed", integrityCheck.Passed})
    integrityCheck.EntityData.Leafs.Append("failed", types.YLeaf{"Failed", integrityCheck.Failed})
    integrityCheck.EntityData.Leafs.Append("skipped", types.YLeaf{"Skipped", integrityCheck.Skipped})
    integrityCheck.EntityData.Leafs.Append("generate-response-failures", types.YLeaf{"GenerateResponseFailures", integrityCheck.GenerateResponseFailures})
    integrityCheck.EntityData.Leafs.Append("unexpected", types.YLeaf{"Unexpected", integrityCheck.Unexpected})
    integrityCheck.EntityData.Leafs.Append("unexpected-zlb", types.YLeaf{"UnexpectedZlb", integrityCheck.UnexpectedZlb})

    integrityCheck.EntityData.YListKeys = []string {}

    return &(integrityCheck.EntityData)
}

// L2tpv2_Counters_Control_TunnelXr_Authentication_LocalSecret
// Local secret statistics
type L2tpv2_Counters_Control_TunnelXr_Authentication_LocalSecret struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Validate. The type is interface{} with range: 0..4294967295.
    Validate interface{}

    // Bad hash. The type is interface{} with range: 0..4294967295.
    BadHash interface{}

    // Bad length. The type is interface{} with range: 0..4294967295.
    BadLength interface{}

    // Ignored. The type is interface{} with range: 0..4294967295.
    Ignored interface{}

    // Missing. The type is interface{} with range: 0..4294967295.
    Missing interface{}

    // Passed. The type is interface{} with range: 0..4294967295.
    Passed interface{}

    // Failed. The type is interface{} with range: 0..4294967295.
    Failed interface{}

    // Skipped. The type is interface{} with range: 0..4294967295.
    Skipped interface{}

    // Generate response fail. The type is interface{} with range: 0..4294967295.
    GenerateResponseFailures interface{}

    // Unexpected. The type is interface{} with range: 0..4294967295.
    Unexpected interface{}

    // Unexpected ZLB. The type is interface{} with range: 0..4294967295.
    UnexpectedZlb interface{}
}

func (localSecret *L2tpv2_Counters_Control_TunnelXr_Authentication_LocalSecret) GetEntityData() *types.CommonEntityData {
    localSecret.EntityData.YFilter = localSecret.YFilter
    localSecret.EntityData.YangName = "local-secret"
    localSecret.EntityData.BundleName = "cisco_ios_xr"
    localSecret.EntityData.ParentYangName = "authentication"
    localSecret.EntityData.SegmentPath = "local-secret"
    localSecret.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localSecret.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localSecret.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localSecret.EntityData.Children = types.NewOrderedMap()
    localSecret.EntityData.Leafs = types.NewOrderedMap()
    localSecret.EntityData.Leafs.Append("validate", types.YLeaf{"Validate", localSecret.Validate})
    localSecret.EntityData.Leafs.Append("bad-hash", types.YLeaf{"BadHash", localSecret.BadHash})
    localSecret.EntityData.Leafs.Append("bad-length", types.YLeaf{"BadLength", localSecret.BadLength})
    localSecret.EntityData.Leafs.Append("ignored", types.YLeaf{"Ignored", localSecret.Ignored})
    localSecret.EntityData.Leafs.Append("missing", types.YLeaf{"Missing", localSecret.Missing})
    localSecret.EntityData.Leafs.Append("passed", types.YLeaf{"Passed", localSecret.Passed})
    localSecret.EntityData.Leafs.Append("failed", types.YLeaf{"Failed", localSecret.Failed})
    localSecret.EntityData.Leafs.Append("skipped", types.YLeaf{"Skipped", localSecret.Skipped})
    localSecret.EntityData.Leafs.Append("generate-response-failures", types.YLeaf{"GenerateResponseFailures", localSecret.GenerateResponseFailures})
    localSecret.EntityData.Leafs.Append("unexpected", types.YLeaf{"Unexpected", localSecret.Unexpected})
    localSecret.EntityData.Leafs.Append("unexpected-zlb", types.YLeaf{"UnexpectedZlb", localSecret.UnexpectedZlb})

    localSecret.EntityData.YListKeys = []string {}

    return &(localSecret.EntityData)
}

// L2tpv2_Counters_Control_TunnelXr_Authentication_ChallengeAvp
// Challenge AVP statistics
type L2tpv2_Counters_Control_TunnelXr_Authentication_ChallengeAvp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Validate. The type is interface{} with range: 0..4294967295.
    Validate interface{}

    // Bad hash. The type is interface{} with range: 0..4294967295.
    BadHash interface{}

    // Bad length. The type is interface{} with range: 0..4294967295.
    BadLength interface{}

    // Ignored. The type is interface{} with range: 0..4294967295.
    Ignored interface{}

    // Missing. The type is interface{} with range: 0..4294967295.
    Missing interface{}

    // Passed. The type is interface{} with range: 0..4294967295.
    Passed interface{}

    // Failed. The type is interface{} with range: 0..4294967295.
    Failed interface{}

    // Skipped. The type is interface{} with range: 0..4294967295.
    Skipped interface{}

    // Generate response fail. The type is interface{} with range: 0..4294967295.
    GenerateResponseFailures interface{}

    // Unexpected. The type is interface{} with range: 0..4294967295.
    Unexpected interface{}

    // Unexpected ZLB. The type is interface{} with range: 0..4294967295.
    UnexpectedZlb interface{}
}

func (challengeAvp *L2tpv2_Counters_Control_TunnelXr_Authentication_ChallengeAvp) GetEntityData() *types.CommonEntityData {
    challengeAvp.EntityData.YFilter = challengeAvp.YFilter
    challengeAvp.EntityData.YangName = "challenge-avp"
    challengeAvp.EntityData.BundleName = "cisco_ios_xr"
    challengeAvp.EntityData.ParentYangName = "authentication"
    challengeAvp.EntityData.SegmentPath = "challenge-avp"
    challengeAvp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    challengeAvp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    challengeAvp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    challengeAvp.EntityData.Children = types.NewOrderedMap()
    challengeAvp.EntityData.Leafs = types.NewOrderedMap()
    challengeAvp.EntityData.Leafs.Append("validate", types.YLeaf{"Validate", challengeAvp.Validate})
    challengeAvp.EntityData.Leafs.Append("bad-hash", types.YLeaf{"BadHash", challengeAvp.BadHash})
    challengeAvp.EntityData.Leafs.Append("bad-length", types.YLeaf{"BadLength", challengeAvp.BadLength})
    challengeAvp.EntityData.Leafs.Append("ignored", types.YLeaf{"Ignored", challengeAvp.Ignored})
    challengeAvp.EntityData.Leafs.Append("missing", types.YLeaf{"Missing", challengeAvp.Missing})
    challengeAvp.EntityData.Leafs.Append("passed", types.YLeaf{"Passed", challengeAvp.Passed})
    challengeAvp.EntityData.Leafs.Append("failed", types.YLeaf{"Failed", challengeAvp.Failed})
    challengeAvp.EntityData.Leafs.Append("skipped", types.YLeaf{"Skipped", challengeAvp.Skipped})
    challengeAvp.EntityData.Leafs.Append("generate-response-failures", types.YLeaf{"GenerateResponseFailures", challengeAvp.GenerateResponseFailures})
    challengeAvp.EntityData.Leafs.Append("unexpected", types.YLeaf{"Unexpected", challengeAvp.Unexpected})
    challengeAvp.EntityData.Leafs.Append("unexpected-zlb", types.YLeaf{"UnexpectedZlb", challengeAvp.UnexpectedZlb})

    challengeAvp.EntityData.YListKeys = []string {}

    return &(challengeAvp.EntityData)
}

// L2tpv2_Counters_Control_TunnelXr_Authentication_ChallengeReponse
// Challenge response statistics
type L2tpv2_Counters_Control_TunnelXr_Authentication_ChallengeReponse struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Validate. The type is interface{} with range: 0..4294967295.
    Validate interface{}

    // Bad hash. The type is interface{} with range: 0..4294967295.
    BadHash interface{}

    // Bad length. The type is interface{} with range: 0..4294967295.
    BadLength interface{}

    // Ignored. The type is interface{} with range: 0..4294967295.
    Ignored interface{}

    // Missing. The type is interface{} with range: 0..4294967295.
    Missing interface{}

    // Passed. The type is interface{} with range: 0..4294967295.
    Passed interface{}

    // Failed. The type is interface{} with range: 0..4294967295.
    Failed interface{}

    // Skipped. The type is interface{} with range: 0..4294967295.
    Skipped interface{}

    // Generate response fail. The type is interface{} with range: 0..4294967295.
    GenerateResponseFailures interface{}

    // Unexpected. The type is interface{} with range: 0..4294967295.
    Unexpected interface{}

    // Unexpected ZLB. The type is interface{} with range: 0..4294967295.
    UnexpectedZlb interface{}
}

func (challengeReponse *L2tpv2_Counters_Control_TunnelXr_Authentication_ChallengeReponse) GetEntityData() *types.CommonEntityData {
    challengeReponse.EntityData.YFilter = challengeReponse.YFilter
    challengeReponse.EntityData.YangName = "challenge-reponse"
    challengeReponse.EntityData.BundleName = "cisco_ios_xr"
    challengeReponse.EntityData.ParentYangName = "authentication"
    challengeReponse.EntityData.SegmentPath = "challenge-reponse"
    challengeReponse.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    challengeReponse.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    challengeReponse.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    challengeReponse.EntityData.Children = types.NewOrderedMap()
    challengeReponse.EntityData.Leafs = types.NewOrderedMap()
    challengeReponse.EntityData.Leafs.Append("validate", types.YLeaf{"Validate", challengeReponse.Validate})
    challengeReponse.EntityData.Leafs.Append("bad-hash", types.YLeaf{"BadHash", challengeReponse.BadHash})
    challengeReponse.EntityData.Leafs.Append("bad-length", types.YLeaf{"BadLength", challengeReponse.BadLength})
    challengeReponse.EntityData.Leafs.Append("ignored", types.YLeaf{"Ignored", challengeReponse.Ignored})
    challengeReponse.EntityData.Leafs.Append("missing", types.YLeaf{"Missing", challengeReponse.Missing})
    challengeReponse.EntityData.Leafs.Append("passed", types.YLeaf{"Passed", challengeReponse.Passed})
    challengeReponse.EntityData.Leafs.Append("failed", types.YLeaf{"Failed", challengeReponse.Failed})
    challengeReponse.EntityData.Leafs.Append("skipped", types.YLeaf{"Skipped", challengeReponse.Skipped})
    challengeReponse.EntityData.Leafs.Append("generate-response-failures", types.YLeaf{"GenerateResponseFailures", challengeReponse.GenerateResponseFailures})
    challengeReponse.EntityData.Leafs.Append("unexpected", types.YLeaf{"Unexpected", challengeReponse.Unexpected})
    challengeReponse.EntityData.Leafs.Append("unexpected-zlb", types.YLeaf{"UnexpectedZlb", challengeReponse.UnexpectedZlb})

    challengeReponse.EntityData.YListKeys = []string {}

    return &(challengeReponse.EntityData)
}

// L2tpv2_Counters_Control_TunnelXr_Authentication_OverallStatistics
// Overall statistics
type L2tpv2_Counters_Control_TunnelXr_Authentication_OverallStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Validate. The type is interface{} with range: 0..4294967295.
    Validate interface{}

    // Bad hash. The type is interface{} with range: 0..4294967295.
    BadHash interface{}

    // Bad length. The type is interface{} with range: 0..4294967295.
    BadLength interface{}

    // Ignored. The type is interface{} with range: 0..4294967295.
    Ignored interface{}

    // Missing. The type is interface{} with range: 0..4294967295.
    Missing interface{}

    // Passed. The type is interface{} with range: 0..4294967295.
    Passed interface{}

    // Failed. The type is interface{} with range: 0..4294967295.
    Failed interface{}

    // Skipped. The type is interface{} with range: 0..4294967295.
    Skipped interface{}

    // Generate response fail. The type is interface{} with range: 0..4294967295.
    GenerateResponseFailures interface{}

    // Unexpected. The type is interface{} with range: 0..4294967295.
    Unexpected interface{}

    // Unexpected ZLB. The type is interface{} with range: 0..4294967295.
    UnexpectedZlb interface{}
}

func (overallStatistics *L2tpv2_Counters_Control_TunnelXr_Authentication_OverallStatistics) GetEntityData() *types.CommonEntityData {
    overallStatistics.EntityData.YFilter = overallStatistics.YFilter
    overallStatistics.EntityData.YangName = "overall-statistics"
    overallStatistics.EntityData.BundleName = "cisco_ios_xr"
    overallStatistics.EntityData.ParentYangName = "authentication"
    overallStatistics.EntityData.SegmentPath = "overall-statistics"
    overallStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    overallStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    overallStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    overallStatistics.EntityData.Children = types.NewOrderedMap()
    overallStatistics.EntityData.Leafs = types.NewOrderedMap()
    overallStatistics.EntityData.Leafs.Append("validate", types.YLeaf{"Validate", overallStatistics.Validate})
    overallStatistics.EntityData.Leafs.Append("bad-hash", types.YLeaf{"BadHash", overallStatistics.BadHash})
    overallStatistics.EntityData.Leafs.Append("bad-length", types.YLeaf{"BadLength", overallStatistics.BadLength})
    overallStatistics.EntityData.Leafs.Append("ignored", types.YLeaf{"Ignored", overallStatistics.Ignored})
    overallStatistics.EntityData.Leafs.Append("missing", types.YLeaf{"Missing", overallStatistics.Missing})
    overallStatistics.EntityData.Leafs.Append("passed", types.YLeaf{"Passed", overallStatistics.Passed})
    overallStatistics.EntityData.Leafs.Append("failed", types.YLeaf{"Failed", overallStatistics.Failed})
    overallStatistics.EntityData.Leafs.Append("skipped", types.YLeaf{"Skipped", overallStatistics.Skipped})
    overallStatistics.EntityData.Leafs.Append("generate-response-failures", types.YLeaf{"GenerateResponseFailures", overallStatistics.GenerateResponseFailures})
    overallStatistics.EntityData.Leafs.Append("unexpected", types.YLeaf{"Unexpected", overallStatistics.Unexpected})
    overallStatistics.EntityData.Leafs.Append("unexpected-zlb", types.YLeaf{"UnexpectedZlb", overallStatistics.UnexpectedZlb})

    overallStatistics.EntityData.YListKeys = []string {}

    return &(overallStatistics.EntityData)
}

// L2tpv2_Counters_Control_TunnelXr_Global
// Tunnel counters
type L2tpv2_Counters_Control_TunnelXr_Global struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Total transmit. The type is interface{} with range: 0..4294967295.
    TotalTransmit interface{}

    // Total retransmit. The type is interface{} with range: 0..4294967295.
    TotalRetransmit interface{}

    // Total received. The type is interface{} with range: 0..4294967295.
    TotalReceived interface{}

    // Total drop. The type is interface{} with range: 0..4294967295.
    TotalDrop interface{}

    // Transmit data.
    Transmit L2tpv2_Counters_Control_TunnelXr_Global_Transmit

    // Re transmit data.
    Retransmit L2tpv2_Counters_Control_TunnelXr_Global_Retransmit

    // Received data.
    Received L2tpv2_Counters_Control_TunnelXr_Global_Received

    // Drop data.
    Drop L2tpv2_Counters_Control_TunnelXr_Global_Drop
}

func (global *L2tpv2_Counters_Control_TunnelXr_Global) GetEntityData() *types.CommonEntityData {
    global.EntityData.YFilter = global.YFilter
    global.EntityData.YangName = "global"
    global.EntityData.BundleName = "cisco_ios_xr"
    global.EntityData.ParentYangName = "tunnel-xr"
    global.EntityData.SegmentPath = "global"
    global.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    global.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    global.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    global.EntityData.Children = types.NewOrderedMap()
    global.EntityData.Children.Append("transmit", types.YChild{"Transmit", &global.Transmit})
    global.EntityData.Children.Append("retransmit", types.YChild{"Retransmit", &global.Retransmit})
    global.EntityData.Children.Append("received", types.YChild{"Received", &global.Received})
    global.EntityData.Children.Append("drop", types.YChild{"Drop", &global.Drop})
    global.EntityData.Leafs = types.NewOrderedMap()
    global.EntityData.Leafs.Append("total-transmit", types.YLeaf{"TotalTransmit", global.TotalTransmit})
    global.EntityData.Leafs.Append("total-retransmit", types.YLeaf{"TotalRetransmit", global.TotalRetransmit})
    global.EntityData.Leafs.Append("total-received", types.YLeaf{"TotalReceived", global.TotalReceived})
    global.EntityData.Leafs.Append("total-drop", types.YLeaf{"TotalDrop", global.TotalDrop})

    global.EntityData.YListKeys = []string {}

    return &(global.EntityData)
}

// L2tpv2_Counters_Control_TunnelXr_Global_Transmit
// Transmit data
type L2tpv2_Counters_Control_TunnelXr_Global_Transmit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Unknown packets. The type is interface{} with range: 0..4294967295.
    UnknownPackets interface{}

    // Zero length body packets. The type is interface{} with range:
    // 0..4294967295.
    ZeroLengthBodyPackets interface{}

    // Start control connection requests. The type is interface{} with range:
    // 0..4294967295.
    StartControlConnectionRequests interface{}

    // Start control connection replies. The type is interface{} with range:
    // 0..4294967295.
    StartControlConnectionReplies interface{}

    // Start control connection notifications. The type is interface{} with range:
    // 0..4294967295.
    StartControlConnectionNotifications interface{}

    // Stop control connection notifications. The type is interface{} with range:
    // 0..4294967295.
    StopControlConnectionNotifications interface{}

    // Keep alive messages. The type is interface{} with range: 0..4294967295.
    HelloPackets interface{}

    // Outgoing call requests. The type is interface{} with range: 0..4294967295.
    OutgoingCallRequests interface{}

    // Outgoing call replies. The type is interface{} with range: 0..4294967295.
    OutgoingCallReplies interface{}

    // Outgoing call connected packets. The type is interface{} with range:
    // 0..4294967295.
    OutgoingCallConnectedPackets interface{}

    // Incoming call requests. The type is interface{} with range: 0..4294967295.
    IncomingCallRequests interface{}

    // Incoming call replies. The type is interface{} with range: 0..4294967295.
    IncomingCallReplies interface{}

    // Incoming call connected packets. The type is interface{} with range:
    // 0..4294967295.
    IncomingCallConnectedPackets interface{}

    // Call disconnect notify packets. The type is interface{} with range:
    // 0..4294967295.
    CallDisconnectNotifyPackets interface{}

    // WAN error notify packets. The type is interface{} with range:
    // 0..4294967295.
    WanErrorNotifyPackets interface{}

    // Set link info packets. The type is interface{} with range: 0..4294967295.
    SetLinkInfoPackets interface{}

    // Service relay request counts. The type is interface{} with range:
    // 0..4294967295.
    ServiceRelayRequests interface{}

    // Service relay reply counts. The type is interface{} with range:
    // 0..4294967295.
    ServiceRelayReplies interface{}

    // Packets acknowledgement. The type is interface{} with range: 0..4294967295.
    AcknowledgementPackets interface{}
}

func (transmit *L2tpv2_Counters_Control_TunnelXr_Global_Transmit) GetEntityData() *types.CommonEntityData {
    transmit.EntityData.YFilter = transmit.YFilter
    transmit.EntityData.YangName = "transmit"
    transmit.EntityData.BundleName = "cisco_ios_xr"
    transmit.EntityData.ParentYangName = "global"
    transmit.EntityData.SegmentPath = "transmit"
    transmit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transmit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transmit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transmit.EntityData.Children = types.NewOrderedMap()
    transmit.EntityData.Leafs = types.NewOrderedMap()
    transmit.EntityData.Leafs.Append("unknown-packets", types.YLeaf{"UnknownPackets", transmit.UnknownPackets})
    transmit.EntityData.Leafs.Append("zero-length-body-packets", types.YLeaf{"ZeroLengthBodyPackets", transmit.ZeroLengthBodyPackets})
    transmit.EntityData.Leafs.Append("start-control-connection-requests", types.YLeaf{"StartControlConnectionRequests", transmit.StartControlConnectionRequests})
    transmit.EntityData.Leafs.Append("start-control-connection-replies", types.YLeaf{"StartControlConnectionReplies", transmit.StartControlConnectionReplies})
    transmit.EntityData.Leafs.Append("start-control-connection-notifications", types.YLeaf{"StartControlConnectionNotifications", transmit.StartControlConnectionNotifications})
    transmit.EntityData.Leafs.Append("stop-control-connection-notifications", types.YLeaf{"StopControlConnectionNotifications", transmit.StopControlConnectionNotifications})
    transmit.EntityData.Leafs.Append("hello-packets", types.YLeaf{"HelloPackets", transmit.HelloPackets})
    transmit.EntityData.Leafs.Append("outgoing-call-requests", types.YLeaf{"OutgoingCallRequests", transmit.OutgoingCallRequests})
    transmit.EntityData.Leafs.Append("outgoing-call-replies", types.YLeaf{"OutgoingCallReplies", transmit.OutgoingCallReplies})
    transmit.EntityData.Leafs.Append("outgoing-call-connected-packets", types.YLeaf{"OutgoingCallConnectedPackets", transmit.OutgoingCallConnectedPackets})
    transmit.EntityData.Leafs.Append("incoming-call-requests", types.YLeaf{"IncomingCallRequests", transmit.IncomingCallRequests})
    transmit.EntityData.Leafs.Append("incoming-call-replies", types.YLeaf{"IncomingCallReplies", transmit.IncomingCallReplies})
    transmit.EntityData.Leafs.Append("incoming-call-connected-packets", types.YLeaf{"IncomingCallConnectedPackets", transmit.IncomingCallConnectedPackets})
    transmit.EntityData.Leafs.Append("call-disconnect-notify-packets", types.YLeaf{"CallDisconnectNotifyPackets", transmit.CallDisconnectNotifyPackets})
    transmit.EntityData.Leafs.Append("wan-error-notify-packets", types.YLeaf{"WanErrorNotifyPackets", transmit.WanErrorNotifyPackets})
    transmit.EntityData.Leafs.Append("set-link-info-packets", types.YLeaf{"SetLinkInfoPackets", transmit.SetLinkInfoPackets})
    transmit.EntityData.Leafs.Append("service-relay-requests", types.YLeaf{"ServiceRelayRequests", transmit.ServiceRelayRequests})
    transmit.EntityData.Leafs.Append("service-relay-replies", types.YLeaf{"ServiceRelayReplies", transmit.ServiceRelayReplies})
    transmit.EntityData.Leafs.Append("acknowledgement-packets", types.YLeaf{"AcknowledgementPackets", transmit.AcknowledgementPackets})

    transmit.EntityData.YListKeys = []string {}

    return &(transmit.EntityData)
}

// L2tpv2_Counters_Control_TunnelXr_Global_Retransmit
// Re transmit data
type L2tpv2_Counters_Control_TunnelXr_Global_Retransmit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Unknown packets. The type is interface{} with range: 0..4294967295.
    UnknownPackets interface{}

    // Zero length body packets. The type is interface{} with range:
    // 0..4294967295.
    ZeroLengthBodyPackets interface{}

    // Start control connection requests. The type is interface{} with range:
    // 0..4294967295.
    StartControlConnectionRequests interface{}

    // Start control connection replies. The type is interface{} with range:
    // 0..4294967295.
    StartControlConnectionReplies interface{}

    // Start control connection notifications. The type is interface{} with range:
    // 0..4294967295.
    StartControlConnectionNotifications interface{}

    // Stop control connection notifications. The type is interface{} with range:
    // 0..4294967295.
    StopControlConnectionNotifications interface{}

    // Keep alive messages. The type is interface{} with range: 0..4294967295.
    HelloPackets interface{}

    // Outgoing call requests. The type is interface{} with range: 0..4294967295.
    OutgoingCallRequests interface{}

    // Outgoing call replies. The type is interface{} with range: 0..4294967295.
    OutgoingCallReplies interface{}

    // Outgoing call connected packets. The type is interface{} with range:
    // 0..4294967295.
    OutgoingCallConnectedPackets interface{}

    // Incoming call requests. The type is interface{} with range: 0..4294967295.
    IncomingCallRequests interface{}

    // Incoming call replies. The type is interface{} with range: 0..4294967295.
    IncomingCallReplies interface{}

    // Incoming call connected packets. The type is interface{} with range:
    // 0..4294967295.
    IncomingCallConnectedPackets interface{}

    // Call disconnect notify packets. The type is interface{} with range:
    // 0..4294967295.
    CallDisconnectNotifyPackets interface{}

    // WAN error notify packets. The type is interface{} with range:
    // 0..4294967295.
    WanErrorNotifyPackets interface{}

    // Set link info packets. The type is interface{} with range: 0..4294967295.
    SetLinkInfoPackets interface{}

    // Service relay request counts. The type is interface{} with range:
    // 0..4294967295.
    ServiceRelayRequests interface{}

    // Service relay reply counts. The type is interface{} with range:
    // 0..4294967295.
    ServiceRelayReplies interface{}

    // Packets acknowledgement. The type is interface{} with range: 0..4294967295.
    AcknowledgementPackets interface{}
}

func (retransmit *L2tpv2_Counters_Control_TunnelXr_Global_Retransmit) GetEntityData() *types.CommonEntityData {
    retransmit.EntityData.YFilter = retransmit.YFilter
    retransmit.EntityData.YangName = "retransmit"
    retransmit.EntityData.BundleName = "cisco_ios_xr"
    retransmit.EntityData.ParentYangName = "global"
    retransmit.EntityData.SegmentPath = "retransmit"
    retransmit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    retransmit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    retransmit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    retransmit.EntityData.Children = types.NewOrderedMap()
    retransmit.EntityData.Leafs = types.NewOrderedMap()
    retransmit.EntityData.Leafs.Append("unknown-packets", types.YLeaf{"UnknownPackets", retransmit.UnknownPackets})
    retransmit.EntityData.Leafs.Append("zero-length-body-packets", types.YLeaf{"ZeroLengthBodyPackets", retransmit.ZeroLengthBodyPackets})
    retransmit.EntityData.Leafs.Append("start-control-connection-requests", types.YLeaf{"StartControlConnectionRequests", retransmit.StartControlConnectionRequests})
    retransmit.EntityData.Leafs.Append("start-control-connection-replies", types.YLeaf{"StartControlConnectionReplies", retransmit.StartControlConnectionReplies})
    retransmit.EntityData.Leafs.Append("start-control-connection-notifications", types.YLeaf{"StartControlConnectionNotifications", retransmit.StartControlConnectionNotifications})
    retransmit.EntityData.Leafs.Append("stop-control-connection-notifications", types.YLeaf{"StopControlConnectionNotifications", retransmit.StopControlConnectionNotifications})
    retransmit.EntityData.Leafs.Append("hello-packets", types.YLeaf{"HelloPackets", retransmit.HelloPackets})
    retransmit.EntityData.Leafs.Append("outgoing-call-requests", types.YLeaf{"OutgoingCallRequests", retransmit.OutgoingCallRequests})
    retransmit.EntityData.Leafs.Append("outgoing-call-replies", types.YLeaf{"OutgoingCallReplies", retransmit.OutgoingCallReplies})
    retransmit.EntityData.Leafs.Append("outgoing-call-connected-packets", types.YLeaf{"OutgoingCallConnectedPackets", retransmit.OutgoingCallConnectedPackets})
    retransmit.EntityData.Leafs.Append("incoming-call-requests", types.YLeaf{"IncomingCallRequests", retransmit.IncomingCallRequests})
    retransmit.EntityData.Leafs.Append("incoming-call-replies", types.YLeaf{"IncomingCallReplies", retransmit.IncomingCallReplies})
    retransmit.EntityData.Leafs.Append("incoming-call-connected-packets", types.YLeaf{"IncomingCallConnectedPackets", retransmit.IncomingCallConnectedPackets})
    retransmit.EntityData.Leafs.Append("call-disconnect-notify-packets", types.YLeaf{"CallDisconnectNotifyPackets", retransmit.CallDisconnectNotifyPackets})
    retransmit.EntityData.Leafs.Append("wan-error-notify-packets", types.YLeaf{"WanErrorNotifyPackets", retransmit.WanErrorNotifyPackets})
    retransmit.EntityData.Leafs.Append("set-link-info-packets", types.YLeaf{"SetLinkInfoPackets", retransmit.SetLinkInfoPackets})
    retransmit.EntityData.Leafs.Append("service-relay-requests", types.YLeaf{"ServiceRelayRequests", retransmit.ServiceRelayRequests})
    retransmit.EntityData.Leafs.Append("service-relay-replies", types.YLeaf{"ServiceRelayReplies", retransmit.ServiceRelayReplies})
    retransmit.EntityData.Leafs.Append("acknowledgement-packets", types.YLeaf{"AcknowledgementPackets", retransmit.AcknowledgementPackets})

    retransmit.EntityData.YListKeys = []string {}

    return &(retransmit.EntityData)
}

// L2tpv2_Counters_Control_TunnelXr_Global_Received
// Received data
type L2tpv2_Counters_Control_TunnelXr_Global_Received struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Unknown packets. The type is interface{} with range: 0..4294967295.
    UnknownPackets interface{}

    // Zero length body packets. The type is interface{} with range:
    // 0..4294967295.
    ZeroLengthBodyPackets interface{}

    // Start control connection requests. The type is interface{} with range:
    // 0..4294967295.
    StartControlConnectionRequests interface{}

    // Start control connection replies. The type is interface{} with range:
    // 0..4294967295.
    StartControlConnectionReplies interface{}

    // Start control connection notifications. The type is interface{} with range:
    // 0..4294967295.
    StartControlConnectionNotifications interface{}

    // Stop control connection notifications. The type is interface{} with range:
    // 0..4294967295.
    StopControlConnectionNotifications interface{}

    // Keep alive messages. The type is interface{} with range: 0..4294967295.
    HelloPackets interface{}

    // Outgoing call requests. The type is interface{} with range: 0..4294967295.
    OutgoingCallRequests interface{}

    // Outgoing call replies. The type is interface{} with range: 0..4294967295.
    OutgoingCallReplies interface{}

    // Outgoing call connected packets. The type is interface{} with range:
    // 0..4294967295.
    OutgoingCallConnectedPackets interface{}

    // Incoming call requests. The type is interface{} with range: 0..4294967295.
    IncomingCallRequests interface{}

    // Incoming call replies. The type is interface{} with range: 0..4294967295.
    IncomingCallReplies interface{}

    // Incoming call connected packets. The type is interface{} with range:
    // 0..4294967295.
    IncomingCallConnectedPackets interface{}

    // Call disconnect notify packets. The type is interface{} with range:
    // 0..4294967295.
    CallDisconnectNotifyPackets interface{}

    // WAN error notify packets. The type is interface{} with range:
    // 0..4294967295.
    WanErrorNotifyPackets interface{}

    // Set link info packets. The type is interface{} with range: 0..4294967295.
    SetLinkInfoPackets interface{}

    // Service relay request counts. The type is interface{} with range:
    // 0..4294967295.
    ServiceRelayRequests interface{}

    // Service relay reply counts. The type is interface{} with range:
    // 0..4294967295.
    ServiceRelayReplies interface{}

    // Packets acknowledgement. The type is interface{} with range: 0..4294967295.
    AcknowledgementPackets interface{}
}

func (received *L2tpv2_Counters_Control_TunnelXr_Global_Received) GetEntityData() *types.CommonEntityData {
    received.EntityData.YFilter = received.YFilter
    received.EntityData.YangName = "received"
    received.EntityData.BundleName = "cisco_ios_xr"
    received.EntityData.ParentYangName = "global"
    received.EntityData.SegmentPath = "received"
    received.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    received.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    received.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    received.EntityData.Children = types.NewOrderedMap()
    received.EntityData.Leafs = types.NewOrderedMap()
    received.EntityData.Leafs.Append("unknown-packets", types.YLeaf{"UnknownPackets", received.UnknownPackets})
    received.EntityData.Leafs.Append("zero-length-body-packets", types.YLeaf{"ZeroLengthBodyPackets", received.ZeroLengthBodyPackets})
    received.EntityData.Leafs.Append("start-control-connection-requests", types.YLeaf{"StartControlConnectionRequests", received.StartControlConnectionRequests})
    received.EntityData.Leafs.Append("start-control-connection-replies", types.YLeaf{"StartControlConnectionReplies", received.StartControlConnectionReplies})
    received.EntityData.Leafs.Append("start-control-connection-notifications", types.YLeaf{"StartControlConnectionNotifications", received.StartControlConnectionNotifications})
    received.EntityData.Leafs.Append("stop-control-connection-notifications", types.YLeaf{"StopControlConnectionNotifications", received.StopControlConnectionNotifications})
    received.EntityData.Leafs.Append("hello-packets", types.YLeaf{"HelloPackets", received.HelloPackets})
    received.EntityData.Leafs.Append("outgoing-call-requests", types.YLeaf{"OutgoingCallRequests", received.OutgoingCallRequests})
    received.EntityData.Leafs.Append("outgoing-call-replies", types.YLeaf{"OutgoingCallReplies", received.OutgoingCallReplies})
    received.EntityData.Leafs.Append("outgoing-call-connected-packets", types.YLeaf{"OutgoingCallConnectedPackets", received.OutgoingCallConnectedPackets})
    received.EntityData.Leafs.Append("incoming-call-requests", types.YLeaf{"IncomingCallRequests", received.IncomingCallRequests})
    received.EntityData.Leafs.Append("incoming-call-replies", types.YLeaf{"IncomingCallReplies", received.IncomingCallReplies})
    received.EntityData.Leafs.Append("incoming-call-connected-packets", types.YLeaf{"IncomingCallConnectedPackets", received.IncomingCallConnectedPackets})
    received.EntityData.Leafs.Append("call-disconnect-notify-packets", types.YLeaf{"CallDisconnectNotifyPackets", received.CallDisconnectNotifyPackets})
    received.EntityData.Leafs.Append("wan-error-notify-packets", types.YLeaf{"WanErrorNotifyPackets", received.WanErrorNotifyPackets})
    received.EntityData.Leafs.Append("set-link-info-packets", types.YLeaf{"SetLinkInfoPackets", received.SetLinkInfoPackets})
    received.EntityData.Leafs.Append("service-relay-requests", types.YLeaf{"ServiceRelayRequests", received.ServiceRelayRequests})
    received.EntityData.Leafs.Append("service-relay-replies", types.YLeaf{"ServiceRelayReplies", received.ServiceRelayReplies})
    received.EntityData.Leafs.Append("acknowledgement-packets", types.YLeaf{"AcknowledgementPackets", received.AcknowledgementPackets})

    received.EntityData.YListKeys = []string {}

    return &(received.EntityData)
}

// L2tpv2_Counters_Control_TunnelXr_Global_Drop
// Drop data
type L2tpv2_Counters_Control_TunnelXr_Global_Drop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Unknown packets. The type is interface{} with range: 0..4294967295.
    UnknownPackets interface{}

    // Zero length body packets. The type is interface{} with range:
    // 0..4294967295.
    ZeroLengthBodyPackets interface{}

    // Start control connection requests. The type is interface{} with range:
    // 0..4294967295.
    StartControlConnectionRequests interface{}

    // Start control connection replies. The type is interface{} with range:
    // 0..4294967295.
    StartControlConnectionReplies interface{}

    // Start control connection notifications. The type is interface{} with range:
    // 0..4294967295.
    StartControlConnectionNotifications interface{}

    // Stop control connection notifications. The type is interface{} with range:
    // 0..4294967295.
    StopControlConnectionNotifications interface{}

    // Keep alive messages. The type is interface{} with range: 0..4294967295.
    HelloPackets interface{}

    // Outgoing call requests. The type is interface{} with range: 0..4294967295.
    OutgoingCallRequests interface{}

    // Outgoing call replies. The type is interface{} with range: 0..4294967295.
    OutgoingCallReplies interface{}

    // Outgoing call connected packets. The type is interface{} with range:
    // 0..4294967295.
    OutgoingCallConnectedPackets interface{}

    // Incoming call requests. The type is interface{} with range: 0..4294967295.
    IncomingCallRequests interface{}

    // Incoming call replies. The type is interface{} with range: 0..4294967295.
    IncomingCallReplies interface{}

    // Incoming call connected packets. The type is interface{} with range:
    // 0..4294967295.
    IncomingCallConnectedPackets interface{}

    // Call disconnect notify packets. The type is interface{} with range:
    // 0..4294967295.
    CallDisconnectNotifyPackets interface{}

    // WAN error notify packets. The type is interface{} with range:
    // 0..4294967295.
    WanErrorNotifyPackets interface{}

    // Set link info packets. The type is interface{} with range: 0..4294967295.
    SetLinkInfoPackets interface{}

    // Service relay request counts. The type is interface{} with range:
    // 0..4294967295.
    ServiceRelayRequests interface{}

    // Service relay reply counts. The type is interface{} with range:
    // 0..4294967295.
    ServiceRelayReplies interface{}

    // Packets acknowledgement. The type is interface{} with range: 0..4294967295.
    AcknowledgementPackets interface{}
}

func (drop *L2tpv2_Counters_Control_TunnelXr_Global_Drop) GetEntityData() *types.CommonEntityData {
    drop.EntityData.YFilter = drop.YFilter
    drop.EntityData.YangName = "drop"
    drop.EntityData.BundleName = "cisco_ios_xr"
    drop.EntityData.ParentYangName = "global"
    drop.EntityData.SegmentPath = "drop"
    drop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    drop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    drop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    drop.EntityData.Children = types.NewOrderedMap()
    drop.EntityData.Leafs = types.NewOrderedMap()
    drop.EntityData.Leafs.Append("unknown-packets", types.YLeaf{"UnknownPackets", drop.UnknownPackets})
    drop.EntityData.Leafs.Append("zero-length-body-packets", types.YLeaf{"ZeroLengthBodyPackets", drop.ZeroLengthBodyPackets})
    drop.EntityData.Leafs.Append("start-control-connection-requests", types.YLeaf{"StartControlConnectionRequests", drop.StartControlConnectionRequests})
    drop.EntityData.Leafs.Append("start-control-connection-replies", types.YLeaf{"StartControlConnectionReplies", drop.StartControlConnectionReplies})
    drop.EntityData.Leafs.Append("start-control-connection-notifications", types.YLeaf{"StartControlConnectionNotifications", drop.StartControlConnectionNotifications})
    drop.EntityData.Leafs.Append("stop-control-connection-notifications", types.YLeaf{"StopControlConnectionNotifications", drop.StopControlConnectionNotifications})
    drop.EntityData.Leafs.Append("hello-packets", types.YLeaf{"HelloPackets", drop.HelloPackets})
    drop.EntityData.Leafs.Append("outgoing-call-requests", types.YLeaf{"OutgoingCallRequests", drop.OutgoingCallRequests})
    drop.EntityData.Leafs.Append("outgoing-call-replies", types.YLeaf{"OutgoingCallReplies", drop.OutgoingCallReplies})
    drop.EntityData.Leafs.Append("outgoing-call-connected-packets", types.YLeaf{"OutgoingCallConnectedPackets", drop.OutgoingCallConnectedPackets})
    drop.EntityData.Leafs.Append("incoming-call-requests", types.YLeaf{"IncomingCallRequests", drop.IncomingCallRequests})
    drop.EntityData.Leafs.Append("incoming-call-replies", types.YLeaf{"IncomingCallReplies", drop.IncomingCallReplies})
    drop.EntityData.Leafs.Append("incoming-call-connected-packets", types.YLeaf{"IncomingCallConnectedPackets", drop.IncomingCallConnectedPackets})
    drop.EntityData.Leafs.Append("call-disconnect-notify-packets", types.YLeaf{"CallDisconnectNotifyPackets", drop.CallDisconnectNotifyPackets})
    drop.EntityData.Leafs.Append("wan-error-notify-packets", types.YLeaf{"WanErrorNotifyPackets", drop.WanErrorNotifyPackets})
    drop.EntityData.Leafs.Append("set-link-info-packets", types.YLeaf{"SetLinkInfoPackets", drop.SetLinkInfoPackets})
    drop.EntityData.Leafs.Append("service-relay-requests", types.YLeaf{"ServiceRelayRequests", drop.ServiceRelayRequests})
    drop.EntityData.Leafs.Append("service-relay-replies", types.YLeaf{"ServiceRelayReplies", drop.ServiceRelayReplies})
    drop.EntityData.Leafs.Append("acknowledgement-packets", types.YLeaf{"AcknowledgementPackets", drop.AcknowledgementPackets})

    drop.EntityData.YListKeys = []string {}

    return &(drop.EntityData)
}

// L2tpv2_Counters_Control_Tunnels
// Table of tunnel IDs of control message counters
type L2tpv2_Counters_Control_Tunnels struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2TP tunnel control message counters. The type is slice of
    // L2tpv2_Counters_Control_Tunnels_Tunnel.
    Tunnel []*L2tpv2_Counters_Control_Tunnels_Tunnel
}

func (tunnels *L2tpv2_Counters_Control_Tunnels) GetEntityData() *types.CommonEntityData {
    tunnels.EntityData.YFilter = tunnels.YFilter
    tunnels.EntityData.YangName = "tunnels"
    tunnels.EntityData.BundleName = "cisco_ios_xr"
    tunnels.EntityData.ParentYangName = "control"
    tunnels.EntityData.SegmentPath = "tunnels"
    tunnels.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnels.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnels.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnels.EntityData.Children = types.NewOrderedMap()
    tunnels.EntityData.Children.Append("tunnel", types.YChild{"Tunnel", nil})
    for i := range tunnels.Tunnel {
        tunnels.EntityData.Children.Append(types.GetSegmentPath(tunnels.Tunnel[i]), types.YChild{"Tunnel", tunnels.Tunnel[i]})
    }
    tunnels.EntityData.Leafs = types.NewOrderedMap()

    tunnels.EntityData.YListKeys = []string {}

    return &(tunnels.EntityData)
}

// L2tpv2_Counters_Control_Tunnels_Tunnel
// L2TP tunnel control message counters
type L2tpv2_Counters_Control_Tunnels_Tunnel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. L2TP tunnel ID. The type is interface{} with
    // range: 0..4294967295.
    TunnelId interface{}

    // L2TP control message local and remote addresses.
    Brief L2tpv2_Counters_Control_Tunnels_Tunnel_Brief

    // Global data.
    Global L2tpv2_Counters_Control_Tunnels_Tunnel_Global
}

func (tunnel *L2tpv2_Counters_Control_Tunnels_Tunnel) GetEntityData() *types.CommonEntityData {
    tunnel.EntityData.YFilter = tunnel.YFilter
    tunnel.EntityData.YangName = "tunnel"
    tunnel.EntityData.BundleName = "cisco_ios_xr"
    tunnel.EntityData.ParentYangName = "tunnels"
    tunnel.EntityData.SegmentPath = "tunnel" + types.AddKeyToken(tunnel.TunnelId, "tunnel-id")
    tunnel.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnel.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnel.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnel.EntityData.Children = types.NewOrderedMap()
    tunnel.EntityData.Children.Append("brief", types.YChild{"Brief", &tunnel.Brief})
    tunnel.EntityData.Children.Append("global", types.YChild{"Global", &tunnel.Global})
    tunnel.EntityData.Leafs = types.NewOrderedMap()
    tunnel.EntityData.Leafs.Append("tunnel-id", types.YLeaf{"TunnelId", tunnel.TunnelId})

    tunnel.EntityData.YListKeys = []string {"TunnelId"}

    return &(tunnel.EntityData)
}

// L2tpv2_Counters_Control_Tunnels_Tunnel_Brief
// L2TP control message local and remote addresses
type L2tpv2_Counters_Control_Tunnels_Tunnel_Brief struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Remote tunnel ID. The type is interface{} with range: 0..4294967295.
    RemoteTunnelId interface{}

    // Local IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    LocalAddress interface{}

    // Remote IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RemoteAddress interface{}
}

func (brief *L2tpv2_Counters_Control_Tunnels_Tunnel_Brief) GetEntityData() *types.CommonEntityData {
    brief.EntityData.YFilter = brief.YFilter
    brief.EntityData.YangName = "brief"
    brief.EntityData.BundleName = "cisco_ios_xr"
    brief.EntityData.ParentYangName = "tunnel"
    brief.EntityData.SegmentPath = "brief"
    brief.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    brief.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    brief.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    brief.EntityData.Children = types.NewOrderedMap()
    brief.EntityData.Leafs = types.NewOrderedMap()
    brief.EntityData.Leafs.Append("remote-tunnel-id", types.YLeaf{"RemoteTunnelId", brief.RemoteTunnelId})
    brief.EntityData.Leafs.Append("local-address", types.YLeaf{"LocalAddress", brief.LocalAddress})
    brief.EntityData.Leafs.Append("remote-address", types.YLeaf{"RemoteAddress", brief.RemoteAddress})

    brief.EntityData.YListKeys = []string {}

    return &(brief.EntityData)
}

// L2tpv2_Counters_Control_Tunnels_Tunnel_Global
// Global data
type L2tpv2_Counters_Control_Tunnels_Tunnel_Global struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Total transmit. The type is interface{} with range: 0..4294967295.
    TotalTransmit interface{}

    // Total retransmit. The type is interface{} with range: 0..4294967295.
    TotalRetransmit interface{}

    // Total received. The type is interface{} with range: 0..4294967295.
    TotalReceived interface{}

    // Total drop. The type is interface{} with range: 0..4294967295.
    TotalDrop interface{}

    // Transmit data.
    Transmit L2tpv2_Counters_Control_Tunnels_Tunnel_Global_Transmit

    // Re transmit data.
    Retransmit L2tpv2_Counters_Control_Tunnels_Tunnel_Global_Retransmit

    // Received data.
    Received L2tpv2_Counters_Control_Tunnels_Tunnel_Global_Received

    // Drop data.
    Drop L2tpv2_Counters_Control_Tunnels_Tunnel_Global_Drop
}

func (global *L2tpv2_Counters_Control_Tunnels_Tunnel_Global) GetEntityData() *types.CommonEntityData {
    global.EntityData.YFilter = global.YFilter
    global.EntityData.YangName = "global"
    global.EntityData.BundleName = "cisco_ios_xr"
    global.EntityData.ParentYangName = "tunnel"
    global.EntityData.SegmentPath = "global"
    global.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    global.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    global.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    global.EntityData.Children = types.NewOrderedMap()
    global.EntityData.Children.Append("transmit", types.YChild{"Transmit", &global.Transmit})
    global.EntityData.Children.Append("retransmit", types.YChild{"Retransmit", &global.Retransmit})
    global.EntityData.Children.Append("received", types.YChild{"Received", &global.Received})
    global.EntityData.Children.Append("drop", types.YChild{"Drop", &global.Drop})
    global.EntityData.Leafs = types.NewOrderedMap()
    global.EntityData.Leafs.Append("total-transmit", types.YLeaf{"TotalTransmit", global.TotalTransmit})
    global.EntityData.Leafs.Append("total-retransmit", types.YLeaf{"TotalRetransmit", global.TotalRetransmit})
    global.EntityData.Leafs.Append("total-received", types.YLeaf{"TotalReceived", global.TotalReceived})
    global.EntityData.Leafs.Append("total-drop", types.YLeaf{"TotalDrop", global.TotalDrop})

    global.EntityData.YListKeys = []string {}

    return &(global.EntityData)
}

// L2tpv2_Counters_Control_Tunnels_Tunnel_Global_Transmit
// Transmit data
type L2tpv2_Counters_Control_Tunnels_Tunnel_Global_Transmit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Unknown packets. The type is interface{} with range: 0..4294967295.
    UnknownPackets interface{}

    // Zero length body packets. The type is interface{} with range:
    // 0..4294967295.
    ZeroLengthBodyPackets interface{}

    // Start control connection requests. The type is interface{} with range:
    // 0..4294967295.
    StartControlConnectionRequests interface{}

    // Start control connection replies. The type is interface{} with range:
    // 0..4294967295.
    StartControlConnectionReplies interface{}

    // Start control connection notifications. The type is interface{} with range:
    // 0..4294967295.
    StartControlConnectionNotifications interface{}

    // Stop control connection notifications. The type is interface{} with range:
    // 0..4294967295.
    StopControlConnectionNotifications interface{}

    // Keep alive messages. The type is interface{} with range: 0..4294967295.
    HelloPackets interface{}

    // Outgoing call requests. The type is interface{} with range: 0..4294967295.
    OutgoingCallRequests interface{}

    // Outgoing call replies. The type is interface{} with range: 0..4294967295.
    OutgoingCallReplies interface{}

    // Outgoing call connected packets. The type is interface{} with range:
    // 0..4294967295.
    OutgoingCallConnectedPackets interface{}

    // Incoming call requests. The type is interface{} with range: 0..4294967295.
    IncomingCallRequests interface{}

    // Incoming call replies. The type is interface{} with range: 0..4294967295.
    IncomingCallReplies interface{}

    // Incoming call connected packets. The type is interface{} with range:
    // 0..4294967295.
    IncomingCallConnectedPackets interface{}

    // Call disconnect notify packets. The type is interface{} with range:
    // 0..4294967295.
    CallDisconnectNotifyPackets interface{}

    // WAN error notify packets. The type is interface{} with range:
    // 0..4294967295.
    WanErrorNotifyPackets interface{}

    // Set link info packets. The type is interface{} with range: 0..4294967295.
    SetLinkInfoPackets interface{}

    // Service relay request counts. The type is interface{} with range:
    // 0..4294967295.
    ServiceRelayRequests interface{}

    // Service relay reply counts. The type is interface{} with range:
    // 0..4294967295.
    ServiceRelayReplies interface{}

    // Packets acknowledgement. The type is interface{} with range: 0..4294967295.
    AcknowledgementPackets interface{}
}

func (transmit *L2tpv2_Counters_Control_Tunnels_Tunnel_Global_Transmit) GetEntityData() *types.CommonEntityData {
    transmit.EntityData.YFilter = transmit.YFilter
    transmit.EntityData.YangName = "transmit"
    transmit.EntityData.BundleName = "cisco_ios_xr"
    transmit.EntityData.ParentYangName = "global"
    transmit.EntityData.SegmentPath = "transmit"
    transmit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transmit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transmit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transmit.EntityData.Children = types.NewOrderedMap()
    transmit.EntityData.Leafs = types.NewOrderedMap()
    transmit.EntityData.Leafs.Append("unknown-packets", types.YLeaf{"UnknownPackets", transmit.UnknownPackets})
    transmit.EntityData.Leafs.Append("zero-length-body-packets", types.YLeaf{"ZeroLengthBodyPackets", transmit.ZeroLengthBodyPackets})
    transmit.EntityData.Leafs.Append("start-control-connection-requests", types.YLeaf{"StartControlConnectionRequests", transmit.StartControlConnectionRequests})
    transmit.EntityData.Leafs.Append("start-control-connection-replies", types.YLeaf{"StartControlConnectionReplies", transmit.StartControlConnectionReplies})
    transmit.EntityData.Leafs.Append("start-control-connection-notifications", types.YLeaf{"StartControlConnectionNotifications", transmit.StartControlConnectionNotifications})
    transmit.EntityData.Leafs.Append("stop-control-connection-notifications", types.YLeaf{"StopControlConnectionNotifications", transmit.StopControlConnectionNotifications})
    transmit.EntityData.Leafs.Append("hello-packets", types.YLeaf{"HelloPackets", transmit.HelloPackets})
    transmit.EntityData.Leafs.Append("outgoing-call-requests", types.YLeaf{"OutgoingCallRequests", transmit.OutgoingCallRequests})
    transmit.EntityData.Leafs.Append("outgoing-call-replies", types.YLeaf{"OutgoingCallReplies", transmit.OutgoingCallReplies})
    transmit.EntityData.Leafs.Append("outgoing-call-connected-packets", types.YLeaf{"OutgoingCallConnectedPackets", transmit.OutgoingCallConnectedPackets})
    transmit.EntityData.Leafs.Append("incoming-call-requests", types.YLeaf{"IncomingCallRequests", transmit.IncomingCallRequests})
    transmit.EntityData.Leafs.Append("incoming-call-replies", types.YLeaf{"IncomingCallReplies", transmit.IncomingCallReplies})
    transmit.EntityData.Leafs.Append("incoming-call-connected-packets", types.YLeaf{"IncomingCallConnectedPackets", transmit.IncomingCallConnectedPackets})
    transmit.EntityData.Leafs.Append("call-disconnect-notify-packets", types.YLeaf{"CallDisconnectNotifyPackets", transmit.CallDisconnectNotifyPackets})
    transmit.EntityData.Leafs.Append("wan-error-notify-packets", types.YLeaf{"WanErrorNotifyPackets", transmit.WanErrorNotifyPackets})
    transmit.EntityData.Leafs.Append("set-link-info-packets", types.YLeaf{"SetLinkInfoPackets", transmit.SetLinkInfoPackets})
    transmit.EntityData.Leafs.Append("service-relay-requests", types.YLeaf{"ServiceRelayRequests", transmit.ServiceRelayRequests})
    transmit.EntityData.Leafs.Append("service-relay-replies", types.YLeaf{"ServiceRelayReplies", transmit.ServiceRelayReplies})
    transmit.EntityData.Leafs.Append("acknowledgement-packets", types.YLeaf{"AcknowledgementPackets", transmit.AcknowledgementPackets})

    transmit.EntityData.YListKeys = []string {}

    return &(transmit.EntityData)
}

// L2tpv2_Counters_Control_Tunnels_Tunnel_Global_Retransmit
// Re transmit data
type L2tpv2_Counters_Control_Tunnels_Tunnel_Global_Retransmit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Unknown packets. The type is interface{} with range: 0..4294967295.
    UnknownPackets interface{}

    // Zero length body packets. The type is interface{} with range:
    // 0..4294967295.
    ZeroLengthBodyPackets interface{}

    // Start control connection requests. The type is interface{} with range:
    // 0..4294967295.
    StartControlConnectionRequests interface{}

    // Start control connection replies. The type is interface{} with range:
    // 0..4294967295.
    StartControlConnectionReplies interface{}

    // Start control connection notifications. The type is interface{} with range:
    // 0..4294967295.
    StartControlConnectionNotifications interface{}

    // Stop control connection notifications. The type is interface{} with range:
    // 0..4294967295.
    StopControlConnectionNotifications interface{}

    // Keep alive messages. The type is interface{} with range: 0..4294967295.
    HelloPackets interface{}

    // Outgoing call requests. The type is interface{} with range: 0..4294967295.
    OutgoingCallRequests interface{}

    // Outgoing call replies. The type is interface{} with range: 0..4294967295.
    OutgoingCallReplies interface{}

    // Outgoing call connected packets. The type is interface{} with range:
    // 0..4294967295.
    OutgoingCallConnectedPackets interface{}

    // Incoming call requests. The type is interface{} with range: 0..4294967295.
    IncomingCallRequests interface{}

    // Incoming call replies. The type is interface{} with range: 0..4294967295.
    IncomingCallReplies interface{}

    // Incoming call connected packets. The type is interface{} with range:
    // 0..4294967295.
    IncomingCallConnectedPackets interface{}

    // Call disconnect notify packets. The type is interface{} with range:
    // 0..4294967295.
    CallDisconnectNotifyPackets interface{}

    // WAN error notify packets. The type is interface{} with range:
    // 0..4294967295.
    WanErrorNotifyPackets interface{}

    // Set link info packets. The type is interface{} with range: 0..4294967295.
    SetLinkInfoPackets interface{}

    // Service relay request counts. The type is interface{} with range:
    // 0..4294967295.
    ServiceRelayRequests interface{}

    // Service relay reply counts. The type is interface{} with range:
    // 0..4294967295.
    ServiceRelayReplies interface{}

    // Packets acknowledgement. The type is interface{} with range: 0..4294967295.
    AcknowledgementPackets interface{}
}

func (retransmit *L2tpv2_Counters_Control_Tunnels_Tunnel_Global_Retransmit) GetEntityData() *types.CommonEntityData {
    retransmit.EntityData.YFilter = retransmit.YFilter
    retransmit.EntityData.YangName = "retransmit"
    retransmit.EntityData.BundleName = "cisco_ios_xr"
    retransmit.EntityData.ParentYangName = "global"
    retransmit.EntityData.SegmentPath = "retransmit"
    retransmit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    retransmit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    retransmit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    retransmit.EntityData.Children = types.NewOrderedMap()
    retransmit.EntityData.Leafs = types.NewOrderedMap()
    retransmit.EntityData.Leafs.Append("unknown-packets", types.YLeaf{"UnknownPackets", retransmit.UnknownPackets})
    retransmit.EntityData.Leafs.Append("zero-length-body-packets", types.YLeaf{"ZeroLengthBodyPackets", retransmit.ZeroLengthBodyPackets})
    retransmit.EntityData.Leafs.Append("start-control-connection-requests", types.YLeaf{"StartControlConnectionRequests", retransmit.StartControlConnectionRequests})
    retransmit.EntityData.Leafs.Append("start-control-connection-replies", types.YLeaf{"StartControlConnectionReplies", retransmit.StartControlConnectionReplies})
    retransmit.EntityData.Leafs.Append("start-control-connection-notifications", types.YLeaf{"StartControlConnectionNotifications", retransmit.StartControlConnectionNotifications})
    retransmit.EntityData.Leafs.Append("stop-control-connection-notifications", types.YLeaf{"StopControlConnectionNotifications", retransmit.StopControlConnectionNotifications})
    retransmit.EntityData.Leafs.Append("hello-packets", types.YLeaf{"HelloPackets", retransmit.HelloPackets})
    retransmit.EntityData.Leafs.Append("outgoing-call-requests", types.YLeaf{"OutgoingCallRequests", retransmit.OutgoingCallRequests})
    retransmit.EntityData.Leafs.Append("outgoing-call-replies", types.YLeaf{"OutgoingCallReplies", retransmit.OutgoingCallReplies})
    retransmit.EntityData.Leafs.Append("outgoing-call-connected-packets", types.YLeaf{"OutgoingCallConnectedPackets", retransmit.OutgoingCallConnectedPackets})
    retransmit.EntityData.Leafs.Append("incoming-call-requests", types.YLeaf{"IncomingCallRequests", retransmit.IncomingCallRequests})
    retransmit.EntityData.Leafs.Append("incoming-call-replies", types.YLeaf{"IncomingCallReplies", retransmit.IncomingCallReplies})
    retransmit.EntityData.Leafs.Append("incoming-call-connected-packets", types.YLeaf{"IncomingCallConnectedPackets", retransmit.IncomingCallConnectedPackets})
    retransmit.EntityData.Leafs.Append("call-disconnect-notify-packets", types.YLeaf{"CallDisconnectNotifyPackets", retransmit.CallDisconnectNotifyPackets})
    retransmit.EntityData.Leafs.Append("wan-error-notify-packets", types.YLeaf{"WanErrorNotifyPackets", retransmit.WanErrorNotifyPackets})
    retransmit.EntityData.Leafs.Append("set-link-info-packets", types.YLeaf{"SetLinkInfoPackets", retransmit.SetLinkInfoPackets})
    retransmit.EntityData.Leafs.Append("service-relay-requests", types.YLeaf{"ServiceRelayRequests", retransmit.ServiceRelayRequests})
    retransmit.EntityData.Leafs.Append("service-relay-replies", types.YLeaf{"ServiceRelayReplies", retransmit.ServiceRelayReplies})
    retransmit.EntityData.Leafs.Append("acknowledgement-packets", types.YLeaf{"AcknowledgementPackets", retransmit.AcknowledgementPackets})

    retransmit.EntityData.YListKeys = []string {}

    return &(retransmit.EntityData)
}

// L2tpv2_Counters_Control_Tunnels_Tunnel_Global_Received
// Received data
type L2tpv2_Counters_Control_Tunnels_Tunnel_Global_Received struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Unknown packets. The type is interface{} with range: 0..4294967295.
    UnknownPackets interface{}

    // Zero length body packets. The type is interface{} with range:
    // 0..4294967295.
    ZeroLengthBodyPackets interface{}

    // Start control connection requests. The type is interface{} with range:
    // 0..4294967295.
    StartControlConnectionRequests interface{}

    // Start control connection replies. The type is interface{} with range:
    // 0..4294967295.
    StartControlConnectionReplies interface{}

    // Start control connection notifications. The type is interface{} with range:
    // 0..4294967295.
    StartControlConnectionNotifications interface{}

    // Stop control connection notifications. The type is interface{} with range:
    // 0..4294967295.
    StopControlConnectionNotifications interface{}

    // Keep alive messages. The type is interface{} with range: 0..4294967295.
    HelloPackets interface{}

    // Outgoing call requests. The type is interface{} with range: 0..4294967295.
    OutgoingCallRequests interface{}

    // Outgoing call replies. The type is interface{} with range: 0..4294967295.
    OutgoingCallReplies interface{}

    // Outgoing call connected packets. The type is interface{} with range:
    // 0..4294967295.
    OutgoingCallConnectedPackets interface{}

    // Incoming call requests. The type is interface{} with range: 0..4294967295.
    IncomingCallRequests interface{}

    // Incoming call replies. The type is interface{} with range: 0..4294967295.
    IncomingCallReplies interface{}

    // Incoming call connected packets. The type is interface{} with range:
    // 0..4294967295.
    IncomingCallConnectedPackets interface{}

    // Call disconnect notify packets. The type is interface{} with range:
    // 0..4294967295.
    CallDisconnectNotifyPackets interface{}

    // WAN error notify packets. The type is interface{} with range:
    // 0..4294967295.
    WanErrorNotifyPackets interface{}

    // Set link info packets. The type is interface{} with range: 0..4294967295.
    SetLinkInfoPackets interface{}

    // Service relay request counts. The type is interface{} with range:
    // 0..4294967295.
    ServiceRelayRequests interface{}

    // Service relay reply counts. The type is interface{} with range:
    // 0..4294967295.
    ServiceRelayReplies interface{}

    // Packets acknowledgement. The type is interface{} with range: 0..4294967295.
    AcknowledgementPackets interface{}
}

func (received *L2tpv2_Counters_Control_Tunnels_Tunnel_Global_Received) GetEntityData() *types.CommonEntityData {
    received.EntityData.YFilter = received.YFilter
    received.EntityData.YangName = "received"
    received.EntityData.BundleName = "cisco_ios_xr"
    received.EntityData.ParentYangName = "global"
    received.EntityData.SegmentPath = "received"
    received.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    received.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    received.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    received.EntityData.Children = types.NewOrderedMap()
    received.EntityData.Leafs = types.NewOrderedMap()
    received.EntityData.Leafs.Append("unknown-packets", types.YLeaf{"UnknownPackets", received.UnknownPackets})
    received.EntityData.Leafs.Append("zero-length-body-packets", types.YLeaf{"ZeroLengthBodyPackets", received.ZeroLengthBodyPackets})
    received.EntityData.Leafs.Append("start-control-connection-requests", types.YLeaf{"StartControlConnectionRequests", received.StartControlConnectionRequests})
    received.EntityData.Leafs.Append("start-control-connection-replies", types.YLeaf{"StartControlConnectionReplies", received.StartControlConnectionReplies})
    received.EntityData.Leafs.Append("start-control-connection-notifications", types.YLeaf{"StartControlConnectionNotifications", received.StartControlConnectionNotifications})
    received.EntityData.Leafs.Append("stop-control-connection-notifications", types.YLeaf{"StopControlConnectionNotifications", received.StopControlConnectionNotifications})
    received.EntityData.Leafs.Append("hello-packets", types.YLeaf{"HelloPackets", received.HelloPackets})
    received.EntityData.Leafs.Append("outgoing-call-requests", types.YLeaf{"OutgoingCallRequests", received.OutgoingCallRequests})
    received.EntityData.Leafs.Append("outgoing-call-replies", types.YLeaf{"OutgoingCallReplies", received.OutgoingCallReplies})
    received.EntityData.Leafs.Append("outgoing-call-connected-packets", types.YLeaf{"OutgoingCallConnectedPackets", received.OutgoingCallConnectedPackets})
    received.EntityData.Leafs.Append("incoming-call-requests", types.YLeaf{"IncomingCallRequests", received.IncomingCallRequests})
    received.EntityData.Leafs.Append("incoming-call-replies", types.YLeaf{"IncomingCallReplies", received.IncomingCallReplies})
    received.EntityData.Leafs.Append("incoming-call-connected-packets", types.YLeaf{"IncomingCallConnectedPackets", received.IncomingCallConnectedPackets})
    received.EntityData.Leafs.Append("call-disconnect-notify-packets", types.YLeaf{"CallDisconnectNotifyPackets", received.CallDisconnectNotifyPackets})
    received.EntityData.Leafs.Append("wan-error-notify-packets", types.YLeaf{"WanErrorNotifyPackets", received.WanErrorNotifyPackets})
    received.EntityData.Leafs.Append("set-link-info-packets", types.YLeaf{"SetLinkInfoPackets", received.SetLinkInfoPackets})
    received.EntityData.Leafs.Append("service-relay-requests", types.YLeaf{"ServiceRelayRequests", received.ServiceRelayRequests})
    received.EntityData.Leafs.Append("service-relay-replies", types.YLeaf{"ServiceRelayReplies", received.ServiceRelayReplies})
    received.EntityData.Leafs.Append("acknowledgement-packets", types.YLeaf{"AcknowledgementPackets", received.AcknowledgementPackets})

    received.EntityData.YListKeys = []string {}

    return &(received.EntityData)
}

// L2tpv2_Counters_Control_Tunnels_Tunnel_Global_Drop
// Drop data
type L2tpv2_Counters_Control_Tunnels_Tunnel_Global_Drop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Unknown packets. The type is interface{} with range: 0..4294967295.
    UnknownPackets interface{}

    // Zero length body packets. The type is interface{} with range:
    // 0..4294967295.
    ZeroLengthBodyPackets interface{}

    // Start control connection requests. The type is interface{} with range:
    // 0..4294967295.
    StartControlConnectionRequests interface{}

    // Start control connection replies. The type is interface{} with range:
    // 0..4294967295.
    StartControlConnectionReplies interface{}

    // Start control connection notifications. The type is interface{} with range:
    // 0..4294967295.
    StartControlConnectionNotifications interface{}

    // Stop control connection notifications. The type is interface{} with range:
    // 0..4294967295.
    StopControlConnectionNotifications interface{}

    // Keep alive messages. The type is interface{} with range: 0..4294967295.
    HelloPackets interface{}

    // Outgoing call requests. The type is interface{} with range: 0..4294967295.
    OutgoingCallRequests interface{}

    // Outgoing call replies. The type is interface{} with range: 0..4294967295.
    OutgoingCallReplies interface{}

    // Outgoing call connected packets. The type is interface{} with range:
    // 0..4294967295.
    OutgoingCallConnectedPackets interface{}

    // Incoming call requests. The type is interface{} with range: 0..4294967295.
    IncomingCallRequests interface{}

    // Incoming call replies. The type is interface{} with range: 0..4294967295.
    IncomingCallReplies interface{}

    // Incoming call connected packets. The type is interface{} with range:
    // 0..4294967295.
    IncomingCallConnectedPackets interface{}

    // Call disconnect notify packets. The type is interface{} with range:
    // 0..4294967295.
    CallDisconnectNotifyPackets interface{}

    // WAN error notify packets. The type is interface{} with range:
    // 0..4294967295.
    WanErrorNotifyPackets interface{}

    // Set link info packets. The type is interface{} with range: 0..4294967295.
    SetLinkInfoPackets interface{}

    // Service relay request counts. The type is interface{} with range:
    // 0..4294967295.
    ServiceRelayRequests interface{}

    // Service relay reply counts. The type is interface{} with range:
    // 0..4294967295.
    ServiceRelayReplies interface{}

    // Packets acknowledgement. The type is interface{} with range: 0..4294967295.
    AcknowledgementPackets interface{}
}

func (drop *L2tpv2_Counters_Control_Tunnels_Tunnel_Global_Drop) GetEntityData() *types.CommonEntityData {
    drop.EntityData.YFilter = drop.YFilter
    drop.EntityData.YangName = "drop"
    drop.EntityData.BundleName = "cisco_ios_xr"
    drop.EntityData.ParentYangName = "global"
    drop.EntityData.SegmentPath = "drop"
    drop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    drop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    drop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    drop.EntityData.Children = types.NewOrderedMap()
    drop.EntityData.Leafs = types.NewOrderedMap()
    drop.EntityData.Leafs.Append("unknown-packets", types.YLeaf{"UnknownPackets", drop.UnknownPackets})
    drop.EntityData.Leafs.Append("zero-length-body-packets", types.YLeaf{"ZeroLengthBodyPackets", drop.ZeroLengthBodyPackets})
    drop.EntityData.Leafs.Append("start-control-connection-requests", types.YLeaf{"StartControlConnectionRequests", drop.StartControlConnectionRequests})
    drop.EntityData.Leafs.Append("start-control-connection-replies", types.YLeaf{"StartControlConnectionReplies", drop.StartControlConnectionReplies})
    drop.EntityData.Leafs.Append("start-control-connection-notifications", types.YLeaf{"StartControlConnectionNotifications", drop.StartControlConnectionNotifications})
    drop.EntityData.Leafs.Append("stop-control-connection-notifications", types.YLeaf{"StopControlConnectionNotifications", drop.StopControlConnectionNotifications})
    drop.EntityData.Leafs.Append("hello-packets", types.YLeaf{"HelloPackets", drop.HelloPackets})
    drop.EntityData.Leafs.Append("outgoing-call-requests", types.YLeaf{"OutgoingCallRequests", drop.OutgoingCallRequests})
    drop.EntityData.Leafs.Append("outgoing-call-replies", types.YLeaf{"OutgoingCallReplies", drop.OutgoingCallReplies})
    drop.EntityData.Leafs.Append("outgoing-call-connected-packets", types.YLeaf{"OutgoingCallConnectedPackets", drop.OutgoingCallConnectedPackets})
    drop.EntityData.Leafs.Append("incoming-call-requests", types.YLeaf{"IncomingCallRequests", drop.IncomingCallRequests})
    drop.EntityData.Leafs.Append("incoming-call-replies", types.YLeaf{"IncomingCallReplies", drop.IncomingCallReplies})
    drop.EntityData.Leafs.Append("incoming-call-connected-packets", types.YLeaf{"IncomingCallConnectedPackets", drop.IncomingCallConnectedPackets})
    drop.EntityData.Leafs.Append("call-disconnect-notify-packets", types.YLeaf{"CallDisconnectNotifyPackets", drop.CallDisconnectNotifyPackets})
    drop.EntityData.Leafs.Append("wan-error-notify-packets", types.YLeaf{"WanErrorNotifyPackets", drop.WanErrorNotifyPackets})
    drop.EntityData.Leafs.Append("set-link-info-packets", types.YLeaf{"SetLinkInfoPackets", drop.SetLinkInfoPackets})
    drop.EntityData.Leafs.Append("service-relay-requests", types.YLeaf{"ServiceRelayRequests", drop.ServiceRelayRequests})
    drop.EntityData.Leafs.Append("service-relay-replies", types.YLeaf{"ServiceRelayReplies", drop.ServiceRelayReplies})
    drop.EntityData.Leafs.Append("acknowledgement-packets", types.YLeaf{"AcknowledgementPackets", drop.AcknowledgementPackets})

    drop.EntityData.YListKeys = []string {}

    return &(drop.EntityData)
}

// L2tpv2_Statistics
// L2TP v2 statistics information
type L2tpv2_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of tunnels. The type is interface{} with range: 0..4294967295.
    Tunnels interface{}

    // Number of sessions. The type is interface{} with range: 0..4294967295.
    Sessions interface{}

    // Number of packets sent. The type is interface{} with range: 0..4294967295.
    SentPackets interface{}

    // Number of packets received. The type is interface{} with range:
    // 0..4294967295.
    ReceivedPackets interface{}

    // Average processing time for received packets  (in micro seconds). The type
    // is interface{} with range: 0..4294967295. Units are microsecond.
    AveragePacketProcessingTime interface{}

    // Out of order packets received. The type is interface{} with range:
    // 0..4294967295.
    ReceivedOutOfOrderPackets interface{}

    // Re order packets. The type is interface{} with range: 0..4294967295.
    ReorderPackets interface{}

    // Re order deviation. The type is interface{} with range: 0..4294967295.
    ReorderDeviationPackets interface{}

    // In coming packets dropped. The type is interface{} with range:
    // 0..4294967295.
    IncomingDroppedPackets interface{}

    // Bufferred packets. The type is interface{} with range: 0..4294967295.
    BufferedPackets interface{}

    // Packets RX in netio. The type is interface{} with range: 0..4294967295.
    NetioPackets interface{}
}

func (statistics *L2tpv2_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "l2tpv2"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = types.NewOrderedMap()
    statistics.EntityData.Leafs = types.NewOrderedMap()
    statistics.EntityData.Leafs.Append("tunnels", types.YLeaf{"Tunnels", statistics.Tunnels})
    statistics.EntityData.Leafs.Append("sessions", types.YLeaf{"Sessions", statistics.Sessions})
    statistics.EntityData.Leafs.Append("sent-packets", types.YLeaf{"SentPackets", statistics.SentPackets})
    statistics.EntityData.Leafs.Append("received-packets", types.YLeaf{"ReceivedPackets", statistics.ReceivedPackets})
    statistics.EntityData.Leafs.Append("average-packet-processing-time", types.YLeaf{"AveragePacketProcessingTime", statistics.AveragePacketProcessingTime})
    statistics.EntityData.Leafs.Append("received-out-of-order-packets", types.YLeaf{"ReceivedOutOfOrderPackets", statistics.ReceivedOutOfOrderPackets})
    statistics.EntityData.Leafs.Append("reorder-packets", types.YLeaf{"ReorderPackets", statistics.ReorderPackets})
    statistics.EntityData.Leafs.Append("reorder-deviation-packets", types.YLeaf{"ReorderDeviationPackets", statistics.ReorderDeviationPackets})
    statistics.EntityData.Leafs.Append("incoming-dropped-packets", types.YLeaf{"IncomingDroppedPackets", statistics.IncomingDroppedPackets})
    statistics.EntityData.Leafs.Append("buffered-packets", types.YLeaf{"BufferedPackets", statistics.BufferedPackets})
    statistics.EntityData.Leafs.Append("netio-packets", types.YLeaf{"NetioPackets", statistics.NetioPackets})

    statistics.EntityData.YListKeys = []string {}

    return &(statistics.EntityData)
}

// L2tpv2_Tunnel
// L2TPv2 tunnel 
type L2tpv2_Tunnel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tunnel accounting counters.
    Accounting L2tpv2_Tunnel_Accounting
}

func (tunnel *L2tpv2_Tunnel) GetEntityData() *types.CommonEntityData {
    tunnel.EntityData.YFilter = tunnel.YFilter
    tunnel.EntityData.YangName = "tunnel"
    tunnel.EntityData.BundleName = "cisco_ios_xr"
    tunnel.EntityData.ParentYangName = "l2tpv2"
    tunnel.EntityData.SegmentPath = "tunnel"
    tunnel.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnel.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnel.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnel.EntityData.Children = types.NewOrderedMap()
    tunnel.EntityData.Children.Append("accounting", types.YChild{"Accounting", &tunnel.Accounting})
    tunnel.EntityData.Leafs = types.NewOrderedMap()

    tunnel.EntityData.YListKeys = []string {}

    return &(tunnel.EntityData)
}

// L2tpv2_Tunnel_Accounting
// Tunnel accounting counters
type L2tpv2_Tunnel_Accounting struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tunnel accounting statistics.
    Statistics L2tpv2_Tunnel_Accounting_Statistics
}

func (accounting *L2tpv2_Tunnel_Accounting) GetEntityData() *types.CommonEntityData {
    accounting.EntityData.YFilter = accounting.YFilter
    accounting.EntityData.YangName = "accounting"
    accounting.EntityData.BundleName = "cisco_ios_xr"
    accounting.EntityData.ParentYangName = "tunnel"
    accounting.EntityData.SegmentPath = "accounting"
    accounting.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accounting.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accounting.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accounting.EntityData.Children = types.NewOrderedMap()
    accounting.EntityData.Children.Append("statistics", types.YChild{"Statistics", &accounting.Statistics})
    accounting.EntityData.Leafs = types.NewOrderedMap()

    accounting.EntityData.YListKeys = []string {}

    return &(accounting.EntityData)
}

// L2tpv2_Tunnel_Accounting_Statistics
// Tunnel accounting statistics
type L2tpv2_Tunnel_Accounting_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Accounting records sent successfully. The type is interface{} with range:
    // 0..18446744073709551615.
    RecordsSentSuccessfully interface{}

    // Accounting start. The type is interface{} with range:
    // 0..18446744073709551615.
    Start interface{}

    // Accounting stop. The type is interface{} with range:
    // 0..18446744073709551615.
    Stop interface{}

    // Accounting reject. The type is interface{} with range:
    // 0..18446744073709551615.
    Reject interface{}

    // Transport failures. The type is interface{} with range:
    // 0..18446744073709551615.
    TransportFailures interface{}

    // Positive acknowledgement. The type is interface{} with range:
    // 0..18446744073709551615.
    PositiveAcknowledgement interface{}

    // Negative acknowledgement. The type is interface{} with range:
    // 0..18446744073709551615.
    NegativeAcknowledgement interface{}

    // Total records checkpointed. The type is interface{} with range:
    // 0..18446744073709551615.
    RecordsCheckpointed interface{}

    // Records fail to checkpoint. The type is interface{} with range:
    // 0..18446744073709551615.
    RecordsFailedToCheckpoint interface{}

    // Records sent from queue. The type is interface{} with range:
    // 0..18446744073709551615.
    RecordsSentFromQueue interface{}

    // Memory failures. The type is interface{} with range: 0..4294967295.
    MemoryFailures interface{}

    // Current checkpoint size. The type is interface{} with range: 0..4294967295.
    CurrentSize interface{}

    // Records recovered from checkpoint. The type is interface{} with range:
    // 0..4294967295.
    RecordsRecoveredFromCheckpoint interface{}

    // Records fail to recover. The type is interface{} with range: 0..4294967295.
    RecordsFailToRecover interface{}

    // Queue statistics size. The type is interface{} with range:
    // -2147483648..2147483647.
    QueueStatisticsSize interface{}
}

func (statistics *L2tpv2_Tunnel_Accounting_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "accounting"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = types.NewOrderedMap()
    statistics.EntityData.Leafs = types.NewOrderedMap()
    statistics.EntityData.Leafs.Append("records-sent-successfully", types.YLeaf{"RecordsSentSuccessfully", statistics.RecordsSentSuccessfully})
    statistics.EntityData.Leafs.Append("start", types.YLeaf{"Start", statistics.Start})
    statistics.EntityData.Leafs.Append("stop", types.YLeaf{"Stop", statistics.Stop})
    statistics.EntityData.Leafs.Append("reject", types.YLeaf{"Reject", statistics.Reject})
    statistics.EntityData.Leafs.Append("transport-failures", types.YLeaf{"TransportFailures", statistics.TransportFailures})
    statistics.EntityData.Leafs.Append("positive-acknowledgement", types.YLeaf{"PositiveAcknowledgement", statistics.PositiveAcknowledgement})
    statistics.EntityData.Leafs.Append("negative-acknowledgement", types.YLeaf{"NegativeAcknowledgement", statistics.NegativeAcknowledgement})
    statistics.EntityData.Leafs.Append("records-checkpointed", types.YLeaf{"RecordsCheckpointed", statistics.RecordsCheckpointed})
    statistics.EntityData.Leafs.Append("records-failed-to-checkpoint", types.YLeaf{"RecordsFailedToCheckpoint", statistics.RecordsFailedToCheckpoint})
    statistics.EntityData.Leafs.Append("records-sent-from-queue", types.YLeaf{"RecordsSentFromQueue", statistics.RecordsSentFromQueue})
    statistics.EntityData.Leafs.Append("memory-failures", types.YLeaf{"MemoryFailures", statistics.MemoryFailures})
    statistics.EntityData.Leafs.Append("current-size", types.YLeaf{"CurrentSize", statistics.CurrentSize})
    statistics.EntityData.Leafs.Append("records-recovered-from-checkpoint", types.YLeaf{"RecordsRecoveredFromCheckpoint", statistics.RecordsRecoveredFromCheckpoint})
    statistics.EntityData.Leafs.Append("records-fail-to-recover", types.YLeaf{"RecordsFailToRecover", statistics.RecordsFailToRecover})
    statistics.EntityData.Leafs.Append("queue-statistics-size", types.YLeaf{"QueueStatisticsSize", statistics.QueueStatisticsSize})

    statistics.EntityData.YListKeys = []string {}

    return &(statistics.EntityData)
}

// L2tpv2_TunnelConfigurations
// List of tunnel IDs
type L2tpv2_TunnelConfigurations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2TP tunnel information. The type is slice of
    // L2tpv2_TunnelConfigurations_TunnelConfiguration.
    TunnelConfiguration []*L2tpv2_TunnelConfigurations_TunnelConfiguration
}

func (tunnelConfigurations *L2tpv2_TunnelConfigurations) GetEntityData() *types.CommonEntityData {
    tunnelConfigurations.EntityData.YFilter = tunnelConfigurations.YFilter
    tunnelConfigurations.EntityData.YangName = "tunnel-configurations"
    tunnelConfigurations.EntityData.BundleName = "cisco_ios_xr"
    tunnelConfigurations.EntityData.ParentYangName = "l2tpv2"
    tunnelConfigurations.EntityData.SegmentPath = "tunnel-configurations"
    tunnelConfigurations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnelConfigurations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnelConfigurations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnelConfigurations.EntityData.Children = types.NewOrderedMap()
    tunnelConfigurations.EntityData.Children.Append("tunnel-configuration", types.YChild{"TunnelConfiguration", nil})
    for i := range tunnelConfigurations.TunnelConfiguration {
        tunnelConfigurations.EntityData.Children.Append(types.GetSegmentPath(tunnelConfigurations.TunnelConfiguration[i]), types.YChild{"TunnelConfiguration", tunnelConfigurations.TunnelConfiguration[i]})
    }
    tunnelConfigurations.EntityData.Leafs = types.NewOrderedMap()

    tunnelConfigurations.EntityData.YListKeys = []string {}

    return &(tunnelConfigurations.EntityData)
}

// L2tpv2_TunnelConfigurations_TunnelConfiguration
// L2TP tunnel information
type L2tpv2_TunnelConfigurations_TunnelConfiguration struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Local tunnel ID. The type is interface{} with
    // range: 0..4294967295.
    LocalTunnelId interface{}

    // Remote tunnel ID. The type is interface{} with range: 0..4294967295.
    RemoteTunnelId interface{}

    // L2Tp class data.
    L2tpClass L2tpv2_TunnelConfigurations_TunnelConfiguration_L2tpClass
}

func (tunnelConfiguration *L2tpv2_TunnelConfigurations_TunnelConfiguration) GetEntityData() *types.CommonEntityData {
    tunnelConfiguration.EntityData.YFilter = tunnelConfiguration.YFilter
    tunnelConfiguration.EntityData.YangName = "tunnel-configuration"
    tunnelConfiguration.EntityData.BundleName = "cisco_ios_xr"
    tunnelConfiguration.EntityData.ParentYangName = "tunnel-configurations"
    tunnelConfiguration.EntityData.SegmentPath = "tunnel-configuration" + types.AddKeyToken(tunnelConfiguration.LocalTunnelId, "local-tunnel-id")
    tunnelConfiguration.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnelConfiguration.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnelConfiguration.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnelConfiguration.EntityData.Children = types.NewOrderedMap()
    tunnelConfiguration.EntityData.Children.Append("l2tp-class", types.YChild{"L2tpClass", &tunnelConfiguration.L2tpClass})
    tunnelConfiguration.EntityData.Leafs = types.NewOrderedMap()
    tunnelConfiguration.EntityData.Leafs.Append("local-tunnel-id", types.YLeaf{"LocalTunnelId", tunnelConfiguration.LocalTunnelId})
    tunnelConfiguration.EntityData.Leafs.Append("remote-tunnel-id", types.YLeaf{"RemoteTunnelId", tunnelConfiguration.RemoteTunnelId})

    tunnelConfiguration.EntityData.YListKeys = []string {"LocalTunnelId"}

    return &(tunnelConfiguration.EntityData)
}

// L2tpv2_TunnelConfigurations_TunnelConfiguration_L2tpClass
// L2Tp class data
type L2tpv2_TunnelConfigurations_TunnelConfiguration_L2tpClass struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP TOS. The type is interface{} with range: 0..255.
    IpTos interface{}

    // VRF name. The type is string with length: 0..256.
    VrfName interface{}

    // Receive window size. The type is interface{} with range: 0..65535.
    ReceiveWindowSize interface{}

    // Class name. The type is string with length: 0..256.
    ClassNameXr interface{}

    // Hash configured as MD5 or SHA1. The type is DigestHash.
    DigestHash interface{}

    // Password. The type is string with length: 0..25.
    Password interface{}

    // Encoded password. The type is string with length: 0..256.
    EncodedPassword interface{}

    // Host name. The type is string with length: 0..256.
    HostName interface{}

    // Accounting List. The type is string with length: 0..256.
    AccountingMethodList interface{}

    // Hello timeout value in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    HelloTimeout interface{}

    // Timeout setup value in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    SetupTimeout interface{}

    // Retransmit minimum timeout in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    RetransmitMinimumTimeout interface{}

    // Retransmit maximum timeout in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    RetransmitMaximumTimeout interface{}

    // Initial timeout minimum in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    InitialRetransmitMinimumTimeout interface{}

    // Initial timeout maximum in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    InitialRetransmitMaximumTimeout interface{}

    // Timeout no user. The type is interface{} with range: 0..4294967295.
    TimeoutNoUser interface{}

    // Retransmit retries. The type is interface{} with range: 0..4294967295.
    RetransmitRetries interface{}

    // Initial retransmit retries. The type is interface{} with range:
    // 0..4294967295.
    InitialRetransmitRetries interface{}

    // True if authentication is enabled. The type is bool.
    IsAuthenticationEnabled interface{}

    // True if class is hidden. The type is bool.
    IsHidden interface{}

    // True if digest authentication is enabled. The type is bool.
    IsDigestEnabled interface{}

    // True if digest check is enabled. The type is bool.
    IsDigestCheckEnabled interface{}

    // True if congestion control is enabled. The type is bool.
    IsCongestionControlEnabled interface{}

    // True if peer address is checked. The type is bool.
    IsPeerAddressChecked interface{}
}

func (l2tpClass *L2tpv2_TunnelConfigurations_TunnelConfiguration_L2tpClass) GetEntityData() *types.CommonEntityData {
    l2tpClass.EntityData.YFilter = l2tpClass.YFilter
    l2tpClass.EntityData.YangName = "l2tp-class"
    l2tpClass.EntityData.BundleName = "cisco_ios_xr"
    l2tpClass.EntityData.ParentYangName = "tunnel-configuration"
    l2tpClass.EntityData.SegmentPath = "l2tp-class"
    l2tpClass.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    l2tpClass.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    l2tpClass.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    l2tpClass.EntityData.Children = types.NewOrderedMap()
    l2tpClass.EntityData.Leafs = types.NewOrderedMap()
    l2tpClass.EntityData.Leafs.Append("ip-tos", types.YLeaf{"IpTos", l2tpClass.IpTos})
    l2tpClass.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", l2tpClass.VrfName})
    l2tpClass.EntityData.Leafs.Append("receive-window-size", types.YLeaf{"ReceiveWindowSize", l2tpClass.ReceiveWindowSize})
    l2tpClass.EntityData.Leafs.Append("class-name-xr", types.YLeaf{"ClassNameXr", l2tpClass.ClassNameXr})
    l2tpClass.EntityData.Leafs.Append("digest-hash", types.YLeaf{"DigestHash", l2tpClass.DigestHash})
    l2tpClass.EntityData.Leafs.Append("password", types.YLeaf{"Password", l2tpClass.Password})
    l2tpClass.EntityData.Leafs.Append("encoded-password", types.YLeaf{"EncodedPassword", l2tpClass.EncodedPassword})
    l2tpClass.EntityData.Leafs.Append("host-name", types.YLeaf{"HostName", l2tpClass.HostName})
    l2tpClass.EntityData.Leafs.Append("accounting-method-list", types.YLeaf{"AccountingMethodList", l2tpClass.AccountingMethodList})
    l2tpClass.EntityData.Leafs.Append("hello-timeout", types.YLeaf{"HelloTimeout", l2tpClass.HelloTimeout})
    l2tpClass.EntityData.Leafs.Append("setup-timeout", types.YLeaf{"SetupTimeout", l2tpClass.SetupTimeout})
    l2tpClass.EntityData.Leafs.Append("retransmit-minimum-timeout", types.YLeaf{"RetransmitMinimumTimeout", l2tpClass.RetransmitMinimumTimeout})
    l2tpClass.EntityData.Leafs.Append("retransmit-maximum-timeout", types.YLeaf{"RetransmitMaximumTimeout", l2tpClass.RetransmitMaximumTimeout})
    l2tpClass.EntityData.Leafs.Append("initial-retransmit-minimum-timeout", types.YLeaf{"InitialRetransmitMinimumTimeout", l2tpClass.InitialRetransmitMinimumTimeout})
    l2tpClass.EntityData.Leafs.Append("initial-retransmit-maximum-timeout", types.YLeaf{"InitialRetransmitMaximumTimeout", l2tpClass.InitialRetransmitMaximumTimeout})
    l2tpClass.EntityData.Leafs.Append("timeout-no-user", types.YLeaf{"TimeoutNoUser", l2tpClass.TimeoutNoUser})
    l2tpClass.EntityData.Leafs.Append("retransmit-retries", types.YLeaf{"RetransmitRetries", l2tpClass.RetransmitRetries})
    l2tpClass.EntityData.Leafs.Append("initial-retransmit-retries", types.YLeaf{"InitialRetransmitRetries", l2tpClass.InitialRetransmitRetries})
    l2tpClass.EntityData.Leafs.Append("is-authentication-enabled", types.YLeaf{"IsAuthenticationEnabled", l2tpClass.IsAuthenticationEnabled})
    l2tpClass.EntityData.Leafs.Append("is-hidden", types.YLeaf{"IsHidden", l2tpClass.IsHidden})
    l2tpClass.EntityData.Leafs.Append("is-digest-enabled", types.YLeaf{"IsDigestEnabled", l2tpClass.IsDigestEnabled})
    l2tpClass.EntityData.Leafs.Append("is-digest-check-enabled", types.YLeaf{"IsDigestCheckEnabled", l2tpClass.IsDigestCheckEnabled})
    l2tpClass.EntityData.Leafs.Append("is-congestion-control-enabled", types.YLeaf{"IsCongestionControlEnabled", l2tpClass.IsCongestionControlEnabled})
    l2tpClass.EntityData.Leafs.Append("is-peer-address-checked", types.YLeaf{"IsPeerAddressChecked", l2tpClass.IsPeerAddressChecked})

    l2tpClass.EntityData.YListKeys = []string {}

    return &(l2tpClass.EntityData)
}

// L2tpv2_CounterHistFail
// Failure events leading to disconnection
type L2tpv2_CounterHistFail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // sesions affected due to timeout. The type is interface{} with range:
    // 0..4294967295.
    SessDownTmout interface{}

    // Send side counters. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    TxCounters interface{}

    // Receive side counters. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    RxCounters interface{}

    // timeout events by packet. The type is slice of
    // L2tpv2_CounterHistFail_PktTimeout.
    PktTimeout []*L2tpv2_CounterHistFail_PktTimeout
}

func (counterHistFail *L2tpv2_CounterHistFail) GetEntityData() *types.CommonEntityData {
    counterHistFail.EntityData.YFilter = counterHistFail.YFilter
    counterHistFail.EntityData.YangName = "counter-hist-fail"
    counterHistFail.EntityData.BundleName = "cisco_ios_xr"
    counterHistFail.EntityData.ParentYangName = "l2tpv2"
    counterHistFail.EntityData.SegmentPath = "counter-hist-fail"
    counterHistFail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    counterHistFail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    counterHistFail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    counterHistFail.EntityData.Children = types.NewOrderedMap()
    counterHistFail.EntityData.Children.Append("pkt-timeout", types.YChild{"PktTimeout", nil})
    for i := range counterHistFail.PktTimeout {
        counterHistFail.EntityData.Children.Append(types.GetSegmentPath(counterHistFail.PktTimeout[i]), types.YChild{"PktTimeout", counterHistFail.PktTimeout[i]})
    }
    counterHistFail.EntityData.Leafs = types.NewOrderedMap()
    counterHistFail.EntityData.Leafs.Append("sess-down-tmout", types.YLeaf{"SessDownTmout", counterHistFail.SessDownTmout})
    counterHistFail.EntityData.Leafs.Append("tx-counters", types.YLeaf{"TxCounters", counterHistFail.TxCounters})
    counterHistFail.EntityData.Leafs.Append("rx-counters", types.YLeaf{"RxCounters", counterHistFail.RxCounters})

    counterHistFail.EntityData.YListKeys = []string {}

    return &(counterHistFail.EntityData)
}

// L2tpv2_CounterHistFail_PktTimeout
// timeout events by packet
type L2tpv2_CounterHistFail_PktTimeout struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    Entry interface{}
}

func (pktTimeout *L2tpv2_CounterHistFail_PktTimeout) GetEntityData() *types.CommonEntityData {
    pktTimeout.EntityData.YFilter = pktTimeout.YFilter
    pktTimeout.EntityData.YangName = "pkt-timeout"
    pktTimeout.EntityData.BundleName = "cisco_ios_xr"
    pktTimeout.EntityData.ParentYangName = "counter-hist-fail"
    pktTimeout.EntityData.SegmentPath = "pkt-timeout"
    pktTimeout.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pktTimeout.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pktTimeout.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pktTimeout.EntityData.Children = types.NewOrderedMap()
    pktTimeout.EntityData.Leafs = types.NewOrderedMap()
    pktTimeout.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", pktTimeout.Entry})

    pktTimeout.EntityData.YListKeys = []string {}

    return &(pktTimeout.EntityData)
}

// L2tpv2_Classes
// List of L2TP class names
type L2tpv2_Classes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2TP class name. The type is slice of L2tpv2_Classes_Class.
    Class []*L2tpv2_Classes_Class
}

func (classes *L2tpv2_Classes) GetEntityData() *types.CommonEntityData {
    classes.EntityData.YFilter = classes.YFilter
    classes.EntityData.YangName = "classes"
    classes.EntityData.BundleName = "cisco_ios_xr"
    classes.EntityData.ParentYangName = "l2tpv2"
    classes.EntityData.SegmentPath = "classes"
    classes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    classes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    classes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    classes.EntityData.Children = types.NewOrderedMap()
    classes.EntityData.Children.Append("class", types.YChild{"Class", nil})
    for i := range classes.Class {
        classes.EntityData.Children.Append(types.GetSegmentPath(classes.Class[i]), types.YChild{"Class", classes.Class[i]})
    }
    classes.EntityData.Leafs = types.NewOrderedMap()

    classes.EntityData.YListKeys = []string {}

    return &(classes.EntityData)
}

// L2tpv2_Classes_Class
// L2TP class name
type L2tpv2_Classes_Class struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. L2TP class name. The type is string with length:
    // 1..31.
    ClassName interface{}

    // IP TOS. The type is interface{} with range: 0..255.
    IpTos interface{}

    // VRF name. The type is string with length: 0..256.
    VrfName interface{}

    // Receive window size. The type is interface{} with range: 0..65535.
    ReceiveWindowSize interface{}

    // Class name. The type is string with length: 0..256.
    ClassNameXr interface{}

    // Hash configured as MD5 or SHA1. The type is DigestHash.
    DigestHash interface{}

    // Password. The type is string with length: 0..25.
    Password interface{}

    // Encoded password. The type is string with length: 0..256.
    EncodedPassword interface{}

    // Host name. The type is string with length: 0..256.
    HostName interface{}

    // Accounting List. The type is string with length: 0..256.
    AccountingMethodList interface{}

    // Hello timeout value in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    HelloTimeout interface{}

    // Timeout setup value in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    SetupTimeout interface{}

    // Retransmit minimum timeout in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    RetransmitMinimumTimeout interface{}

    // Retransmit maximum timeout in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    RetransmitMaximumTimeout interface{}

    // Initial timeout minimum in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    InitialRetransmitMinimumTimeout interface{}

    // Initial timeout maximum in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    InitialRetransmitMaximumTimeout interface{}

    // Timeout no user. The type is interface{} with range: 0..4294967295.
    TimeoutNoUser interface{}

    // Retransmit retries. The type is interface{} with range: 0..4294967295.
    RetransmitRetries interface{}

    // Initial retransmit retries. The type is interface{} with range:
    // 0..4294967295.
    InitialRetransmitRetries interface{}

    // True if authentication is enabled. The type is bool.
    IsAuthenticationEnabled interface{}

    // True if class is hidden. The type is bool.
    IsHidden interface{}

    // True if digest authentication is enabled. The type is bool.
    IsDigestEnabled interface{}

    // True if digest check is enabled. The type is bool.
    IsDigestCheckEnabled interface{}

    // True if congestion control is enabled. The type is bool.
    IsCongestionControlEnabled interface{}

    // True if peer address is checked. The type is bool.
    IsPeerAddressChecked interface{}
}

func (class *L2tpv2_Classes_Class) GetEntityData() *types.CommonEntityData {
    class.EntityData.YFilter = class.YFilter
    class.EntityData.YangName = "class"
    class.EntityData.BundleName = "cisco_ios_xr"
    class.EntityData.ParentYangName = "classes"
    class.EntityData.SegmentPath = "class" + types.AddKeyToken(class.ClassName, "class-name")
    class.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    class.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    class.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    class.EntityData.Children = types.NewOrderedMap()
    class.EntityData.Leafs = types.NewOrderedMap()
    class.EntityData.Leafs.Append("class-name", types.YLeaf{"ClassName", class.ClassName})
    class.EntityData.Leafs.Append("ip-tos", types.YLeaf{"IpTos", class.IpTos})
    class.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", class.VrfName})
    class.EntityData.Leafs.Append("receive-window-size", types.YLeaf{"ReceiveWindowSize", class.ReceiveWindowSize})
    class.EntityData.Leafs.Append("class-name-xr", types.YLeaf{"ClassNameXr", class.ClassNameXr})
    class.EntityData.Leafs.Append("digest-hash", types.YLeaf{"DigestHash", class.DigestHash})
    class.EntityData.Leafs.Append("password", types.YLeaf{"Password", class.Password})
    class.EntityData.Leafs.Append("encoded-password", types.YLeaf{"EncodedPassword", class.EncodedPassword})
    class.EntityData.Leafs.Append("host-name", types.YLeaf{"HostName", class.HostName})
    class.EntityData.Leafs.Append("accounting-method-list", types.YLeaf{"AccountingMethodList", class.AccountingMethodList})
    class.EntityData.Leafs.Append("hello-timeout", types.YLeaf{"HelloTimeout", class.HelloTimeout})
    class.EntityData.Leafs.Append("setup-timeout", types.YLeaf{"SetupTimeout", class.SetupTimeout})
    class.EntityData.Leafs.Append("retransmit-minimum-timeout", types.YLeaf{"RetransmitMinimumTimeout", class.RetransmitMinimumTimeout})
    class.EntityData.Leafs.Append("retransmit-maximum-timeout", types.YLeaf{"RetransmitMaximumTimeout", class.RetransmitMaximumTimeout})
    class.EntityData.Leafs.Append("initial-retransmit-minimum-timeout", types.YLeaf{"InitialRetransmitMinimumTimeout", class.InitialRetransmitMinimumTimeout})
    class.EntityData.Leafs.Append("initial-retransmit-maximum-timeout", types.YLeaf{"InitialRetransmitMaximumTimeout", class.InitialRetransmitMaximumTimeout})
    class.EntityData.Leafs.Append("timeout-no-user", types.YLeaf{"TimeoutNoUser", class.TimeoutNoUser})
    class.EntityData.Leafs.Append("retransmit-retries", types.YLeaf{"RetransmitRetries", class.RetransmitRetries})
    class.EntityData.Leafs.Append("initial-retransmit-retries", types.YLeaf{"InitialRetransmitRetries", class.InitialRetransmitRetries})
    class.EntityData.Leafs.Append("is-authentication-enabled", types.YLeaf{"IsAuthenticationEnabled", class.IsAuthenticationEnabled})
    class.EntityData.Leafs.Append("is-hidden", types.YLeaf{"IsHidden", class.IsHidden})
    class.EntityData.Leafs.Append("is-digest-enabled", types.YLeaf{"IsDigestEnabled", class.IsDigestEnabled})
    class.EntityData.Leafs.Append("is-digest-check-enabled", types.YLeaf{"IsDigestCheckEnabled", class.IsDigestCheckEnabled})
    class.EntityData.Leafs.Append("is-congestion-control-enabled", types.YLeaf{"IsCongestionControlEnabled", class.IsCongestionControlEnabled})
    class.EntityData.Leafs.Append("is-peer-address-checked", types.YLeaf{"IsPeerAddressChecked", class.IsPeerAddressChecked})

    class.EntityData.YListKeys = []string {"ClassName"}

    return &(class.EntityData)
}

// L2tpv2_Tunnels
// List of tunnel IDs
type L2tpv2_Tunnels struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2TP tunnel  information. The type is slice of L2tpv2_Tunnels_Tunnel.
    Tunnel []*L2tpv2_Tunnels_Tunnel
}

func (tunnels *L2tpv2_Tunnels) GetEntityData() *types.CommonEntityData {
    tunnels.EntityData.YFilter = tunnels.YFilter
    tunnels.EntityData.YangName = "tunnels"
    tunnels.EntityData.BundleName = "cisco_ios_xr"
    tunnels.EntityData.ParentYangName = "l2tpv2"
    tunnels.EntityData.SegmentPath = "tunnels"
    tunnels.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnels.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnels.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnels.EntityData.Children = types.NewOrderedMap()
    tunnels.EntityData.Children.Append("tunnel", types.YChild{"Tunnel", nil})
    for i := range tunnels.Tunnel {
        tunnels.EntityData.Children.Append(types.GetSegmentPath(tunnels.Tunnel[i]), types.YChild{"Tunnel", tunnels.Tunnel[i]})
    }
    tunnels.EntityData.Leafs = types.NewOrderedMap()

    tunnels.EntityData.YListKeys = []string {}

    return &(tunnels.EntityData)
}

// L2tpv2_Tunnels_Tunnel
// L2TP tunnel  information
type L2tpv2_Tunnels_Tunnel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Local tunnel ID. The type is interface{} with
    // range: 0..4294967295.
    LocalTunnelId interface{}

    // Local tunnel address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    LocalAddress interface{}

    // Remote tunnel address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RemoteAddress interface{}

    // Local port. The type is interface{} with range: 0..65535.
    LocalPort interface{}

    // Remote port. The type is interface{} with range: 0..65535.
    RemotePort interface{}

    // Protocol. The type is interface{} with range: 0..255.
    Protocol interface{}

    // True if tunnel PMTU checking is enabled. The type is bool.
    IsPmtuEnabled interface{}

    // Remote tunnel ID. The type is interface{} with range: 0..4294967295.
    RemoteTunnelId interface{}

    // Local tunnel name. The type is string with length: 0..256.
    LocalTunnelName interface{}

    // Remote tunnel name. The type is string with length: 0..256.
    RemoteTunnelName interface{}

    // L2TP class name. The type is string with length: 0..256.
    ClassName interface{}

    // Number of active sessions. The type is interface{} with range:
    // 0..4294967295.
    ActiveSessions interface{}

    // Sequence NS. The type is interface{} with range: 0..65535.
    SequenceNs interface{}

    // Sequence NR. The type is interface{} with range: 0..65535.
    SequenceNr interface{}

    // Local window size. The type is interface{} with range: 0..65535.
    LocalWindowSize interface{}

    // Remote window size. The type is interface{} with range: 0..65535.
    RemoteWindowSize interface{}

    // Retransmission time in seconds. The type is interface{} with range:
    // 0..65535. Units are second.
    RetransmissionTime interface{}

    // Maximum retransmission time in seconds. The type is interface{} with range:
    // 0..65535. Units are second.
    MaximumRetransmissionTime interface{}

    // Unsent queue size. The type is interface{} with range: 0..65535.
    UnsentQueueSize interface{}

    // Unsent maximum queue size. The type is interface{} with range: 0..65535.
    UnsentMaximumQueueSize interface{}

    // Resend queue size. The type is interface{} with range: 0..65535.
    ResendQueueSize interface{}

    // Resend maximum queue size. The type is interface{} with range: 0..65535.
    ResendMaximumQueueSize interface{}

    // Order queue size. The type is interface{} with range: 0..65535.
    OrderQueueSize interface{}

    // Current number session packet queue check. The type is interface{} with
    // range: 0..65535.
    PacketQueueCheck interface{}

    // Control message authentication with digest secrets. The type is interface{}
    // with range: 0..65535.
    DigestSecrets interface{}

    // Total resends. The type is interface{} with range: 0..4294967295.
    Resends interface{}

    // Total zero length body acknowledgement. The type is interface{} with range:
    // 0..4294967295.
    ZeroLengthBodyAcknowledgementSent interface{}

    // Total out of order dropped packets. The type is interface{} with range:
    // 0..4294967295.
    TotalOutOfOrderDropPackets interface{}

    // Total out of order reorder packets. The type is interface{} with range:
    // 0..4294967295.
    TotalOutOfOrderReorderPackets interface{}

    // Number of peer authentication failures. The type is interface{} with range:
    // 0..4294967295.
    TotalPeerAuthenticationFailures interface{}

    // True if tunnel is up. The type is bool.
    IsTunnelUp interface{}

    // True if congestion control is enabled else false. The type is bool.
    IsCongestionControlEnabled interface{}

    // Retransmit time distribution in seconds. The type is slice of
    // L2tpv2_Tunnels_Tunnel_RetransmitTime.
    RetransmitTime []*L2tpv2_Tunnels_Tunnel_RetransmitTime
}

func (tunnel *L2tpv2_Tunnels_Tunnel) GetEntityData() *types.CommonEntityData {
    tunnel.EntityData.YFilter = tunnel.YFilter
    tunnel.EntityData.YangName = "tunnel"
    tunnel.EntityData.BundleName = "cisco_ios_xr"
    tunnel.EntityData.ParentYangName = "tunnels"
    tunnel.EntityData.SegmentPath = "tunnel" + types.AddKeyToken(tunnel.LocalTunnelId, "local-tunnel-id")
    tunnel.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnel.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnel.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnel.EntityData.Children = types.NewOrderedMap()
    tunnel.EntityData.Children.Append("retransmit-time", types.YChild{"RetransmitTime", nil})
    for i := range tunnel.RetransmitTime {
        tunnel.EntityData.Children.Append(types.GetSegmentPath(tunnel.RetransmitTime[i]), types.YChild{"RetransmitTime", tunnel.RetransmitTime[i]})
    }
    tunnel.EntityData.Leafs = types.NewOrderedMap()
    tunnel.EntityData.Leafs.Append("local-tunnel-id", types.YLeaf{"LocalTunnelId", tunnel.LocalTunnelId})
    tunnel.EntityData.Leafs.Append("local-address", types.YLeaf{"LocalAddress", tunnel.LocalAddress})
    tunnel.EntityData.Leafs.Append("remote-address", types.YLeaf{"RemoteAddress", tunnel.RemoteAddress})
    tunnel.EntityData.Leafs.Append("local-port", types.YLeaf{"LocalPort", tunnel.LocalPort})
    tunnel.EntityData.Leafs.Append("remote-port", types.YLeaf{"RemotePort", tunnel.RemotePort})
    tunnel.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", tunnel.Protocol})
    tunnel.EntityData.Leafs.Append("is-pmtu-enabled", types.YLeaf{"IsPmtuEnabled", tunnel.IsPmtuEnabled})
    tunnel.EntityData.Leafs.Append("remote-tunnel-id", types.YLeaf{"RemoteTunnelId", tunnel.RemoteTunnelId})
    tunnel.EntityData.Leafs.Append("local-tunnel-name", types.YLeaf{"LocalTunnelName", tunnel.LocalTunnelName})
    tunnel.EntityData.Leafs.Append("remote-tunnel-name", types.YLeaf{"RemoteTunnelName", tunnel.RemoteTunnelName})
    tunnel.EntityData.Leafs.Append("class-name", types.YLeaf{"ClassName", tunnel.ClassName})
    tunnel.EntityData.Leafs.Append("active-sessions", types.YLeaf{"ActiveSessions", tunnel.ActiveSessions})
    tunnel.EntityData.Leafs.Append("sequence-ns", types.YLeaf{"SequenceNs", tunnel.SequenceNs})
    tunnel.EntityData.Leafs.Append("sequence-nr", types.YLeaf{"SequenceNr", tunnel.SequenceNr})
    tunnel.EntityData.Leafs.Append("local-window-size", types.YLeaf{"LocalWindowSize", tunnel.LocalWindowSize})
    tunnel.EntityData.Leafs.Append("remote-window-size", types.YLeaf{"RemoteWindowSize", tunnel.RemoteWindowSize})
    tunnel.EntityData.Leafs.Append("retransmission-time", types.YLeaf{"RetransmissionTime", tunnel.RetransmissionTime})
    tunnel.EntityData.Leafs.Append("maximum-retransmission-time", types.YLeaf{"MaximumRetransmissionTime", tunnel.MaximumRetransmissionTime})
    tunnel.EntityData.Leafs.Append("unsent-queue-size", types.YLeaf{"UnsentQueueSize", tunnel.UnsentQueueSize})
    tunnel.EntityData.Leafs.Append("unsent-maximum-queue-size", types.YLeaf{"UnsentMaximumQueueSize", tunnel.UnsentMaximumQueueSize})
    tunnel.EntityData.Leafs.Append("resend-queue-size", types.YLeaf{"ResendQueueSize", tunnel.ResendQueueSize})
    tunnel.EntityData.Leafs.Append("resend-maximum-queue-size", types.YLeaf{"ResendMaximumQueueSize", tunnel.ResendMaximumQueueSize})
    tunnel.EntityData.Leafs.Append("order-queue-size", types.YLeaf{"OrderQueueSize", tunnel.OrderQueueSize})
    tunnel.EntityData.Leafs.Append("packet-queue-check", types.YLeaf{"PacketQueueCheck", tunnel.PacketQueueCheck})
    tunnel.EntityData.Leafs.Append("digest-secrets", types.YLeaf{"DigestSecrets", tunnel.DigestSecrets})
    tunnel.EntityData.Leafs.Append("resends", types.YLeaf{"Resends", tunnel.Resends})
    tunnel.EntityData.Leafs.Append("zero-length-body-acknowledgement-sent", types.YLeaf{"ZeroLengthBodyAcknowledgementSent", tunnel.ZeroLengthBodyAcknowledgementSent})
    tunnel.EntityData.Leafs.Append("total-out-of-order-drop-packets", types.YLeaf{"TotalOutOfOrderDropPackets", tunnel.TotalOutOfOrderDropPackets})
    tunnel.EntityData.Leafs.Append("total-out-of-order-reorder-packets", types.YLeaf{"TotalOutOfOrderReorderPackets", tunnel.TotalOutOfOrderReorderPackets})
    tunnel.EntityData.Leafs.Append("total-peer-authentication-failures", types.YLeaf{"TotalPeerAuthenticationFailures", tunnel.TotalPeerAuthenticationFailures})
    tunnel.EntityData.Leafs.Append("is-tunnel-up", types.YLeaf{"IsTunnelUp", tunnel.IsTunnelUp})
    tunnel.EntityData.Leafs.Append("is-congestion-control-enabled", types.YLeaf{"IsCongestionControlEnabled", tunnel.IsCongestionControlEnabled})

    tunnel.EntityData.YListKeys = []string {"LocalTunnelId"}

    return &(tunnel.EntityData)
}

// L2tpv2_Tunnels_Tunnel_RetransmitTime
// Retransmit time distribution in seconds
type L2tpv2_Tunnels_Tunnel_RetransmitTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..65535. Units are second.
    Entry interface{}
}

func (retransmitTime *L2tpv2_Tunnels_Tunnel_RetransmitTime) GetEntityData() *types.CommonEntityData {
    retransmitTime.EntityData.YFilter = retransmitTime.YFilter
    retransmitTime.EntityData.YangName = "retransmit-time"
    retransmitTime.EntityData.BundleName = "cisco_ios_xr"
    retransmitTime.EntityData.ParentYangName = "tunnel"
    retransmitTime.EntityData.SegmentPath = "retransmit-time"
    retransmitTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    retransmitTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    retransmitTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    retransmitTime.EntityData.Children = types.NewOrderedMap()
    retransmitTime.EntityData.Leafs = types.NewOrderedMap()
    retransmitTime.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", retransmitTime.Entry})

    retransmitTime.EntityData.YListKeys = []string {}

    return &(retransmitTime.EntityData)
}

// L2tpv2_Sessions
// List of session IDs
type L2tpv2_Sessions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2TP information for a particular session. The type is slice of
    // L2tpv2_Sessions_Session.
    Session []*L2tpv2_Sessions_Session
}

func (sessions *L2tpv2_Sessions) GetEntityData() *types.CommonEntityData {
    sessions.EntityData.YFilter = sessions.YFilter
    sessions.EntityData.YangName = "sessions"
    sessions.EntityData.BundleName = "cisco_ios_xr"
    sessions.EntityData.ParentYangName = "l2tpv2"
    sessions.EntityData.SegmentPath = "sessions"
    sessions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessions.EntityData.Children = types.NewOrderedMap()
    sessions.EntityData.Children.Append("session", types.YChild{"Session", nil})
    for i := range sessions.Session {
        sessions.EntityData.Children.Append(types.GetSegmentPath(sessions.Session[i]), types.YChild{"Session", sessions.Session[i]})
    }
    sessions.EntityData.Leafs = types.NewOrderedMap()

    sessions.EntityData.YListKeys = []string {}

    return &(sessions.EntityData)
}

// L2tpv2_Sessions_Session
// L2TP information for a particular session
type L2tpv2_Sessions_Session struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Local tunnel ID. The type is interface{} with
    // range: 0..4294967295.
    LocalTunnelId interface{}

    // This attribute is a key. Local session ID. The type is interface{} with
    // range: 0..4294967295.
    LocalSessionId interface{}

    // Local session IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    LocalIpAddress interface{}

    // Remote session IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RemoteIpAddress interface{}

    // l2tp sh sess udp lport. The type is interface{} with range: 0..65535.
    L2tpShSessUdpLport interface{}

    // l2tp sh sess udp rport. The type is interface{} with range: 0..65535.
    L2tpShSessUdpRport interface{}

    // Protocol. The type is interface{} with range: 0..255.
    Protocol interface{}

    // Remote tunnel ID. The type is interface{} with range: 0..4294967295.
    RemoteTunnelId interface{}

    // Call serial number. The type is interface{} with range: 0..4294967295.
    CallSerialNumber interface{}

    // Local tunnel name. The type is string with length: 0..256.
    LocalTunnelName interface{}

    // Remote tunnel name. The type is string with length: 0..256.
    RemoteTunnelName interface{}

    // Remote session ID. The type is interface{} with range: 0..4294967295.
    RemoteSessionId interface{}

    // l2tp sh sess tie breaker enabled. The type is interface{} with range:
    // 0..255.
    L2tpShSessTieBreakerEnabled interface{}

    // l2tp sh sess tie breaker. The type is interface{} with range:
    // 0..18446744073709551615.
    L2tpShSessTieBreaker interface{}

    // True if session is manual. The type is bool.
    IsSessionManual interface{}

    // True if session is up. The type is bool.
    IsSessionUp interface{}

    // True if UDP checksum enabled. The type is bool.
    IsUdpChecksumEnabled interface{}

    // True if session sequence is on. The type is bool.
    IsSequencingOn interface{}

    // True if session state is established. The type is bool.
    IsSessionStateEstablished interface{}

    // True if session initiated locally. The type is bool.
    IsSessionLocallyInitiated interface{}

    // True if conditional debugging is enabled. The type is bool.
    IsConditionalDebugEnabled interface{}

    // Unique ID. The type is interface{} with range: 0..4294967295.
    UniqueId interface{}

    // Interface name. The type is string with length: 0..256.
    InterfaceName interface{}

    // Session application data.
    SessionApplicationData L2tpv2_Sessions_Session_SessionApplicationData
}

func (session *L2tpv2_Sessions_Session) GetEntityData() *types.CommonEntityData {
    session.EntityData.YFilter = session.YFilter
    session.EntityData.YangName = "session"
    session.EntityData.BundleName = "cisco_ios_xr"
    session.EntityData.ParentYangName = "sessions"
    session.EntityData.SegmentPath = "session" + types.AddKeyToken(session.LocalTunnelId, "local-tunnel-id") + types.AddKeyToken(session.LocalSessionId, "local-session-id")
    session.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    session.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    session.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    session.EntityData.Children = types.NewOrderedMap()
    session.EntityData.Children.Append("session-application-data", types.YChild{"SessionApplicationData", &session.SessionApplicationData})
    session.EntityData.Leafs = types.NewOrderedMap()
    session.EntityData.Leafs.Append("local-tunnel-id", types.YLeaf{"LocalTunnelId", session.LocalTunnelId})
    session.EntityData.Leafs.Append("local-session-id", types.YLeaf{"LocalSessionId", session.LocalSessionId})
    session.EntityData.Leafs.Append("local-ip-address", types.YLeaf{"LocalIpAddress", session.LocalIpAddress})
    session.EntityData.Leafs.Append("remote-ip-address", types.YLeaf{"RemoteIpAddress", session.RemoteIpAddress})
    session.EntityData.Leafs.Append("l2tp-sh-sess-udp-lport", types.YLeaf{"L2tpShSessUdpLport", session.L2tpShSessUdpLport})
    session.EntityData.Leafs.Append("l2tp-sh-sess-udp-rport", types.YLeaf{"L2tpShSessUdpRport", session.L2tpShSessUdpRport})
    session.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", session.Protocol})
    session.EntityData.Leafs.Append("remote-tunnel-id", types.YLeaf{"RemoteTunnelId", session.RemoteTunnelId})
    session.EntityData.Leafs.Append("call-serial-number", types.YLeaf{"CallSerialNumber", session.CallSerialNumber})
    session.EntityData.Leafs.Append("local-tunnel-name", types.YLeaf{"LocalTunnelName", session.LocalTunnelName})
    session.EntityData.Leafs.Append("remote-tunnel-name", types.YLeaf{"RemoteTunnelName", session.RemoteTunnelName})
    session.EntityData.Leafs.Append("remote-session-id", types.YLeaf{"RemoteSessionId", session.RemoteSessionId})
    session.EntityData.Leafs.Append("l2tp-sh-sess-tie-breaker-enabled", types.YLeaf{"L2tpShSessTieBreakerEnabled", session.L2tpShSessTieBreakerEnabled})
    session.EntityData.Leafs.Append("l2tp-sh-sess-tie-breaker", types.YLeaf{"L2tpShSessTieBreaker", session.L2tpShSessTieBreaker})
    session.EntityData.Leafs.Append("is-session-manual", types.YLeaf{"IsSessionManual", session.IsSessionManual})
    session.EntityData.Leafs.Append("is-session-up", types.YLeaf{"IsSessionUp", session.IsSessionUp})
    session.EntityData.Leafs.Append("is-udp-checksum-enabled", types.YLeaf{"IsUdpChecksumEnabled", session.IsUdpChecksumEnabled})
    session.EntityData.Leafs.Append("is-sequencing-on", types.YLeaf{"IsSequencingOn", session.IsSequencingOn})
    session.EntityData.Leafs.Append("is-session-state-established", types.YLeaf{"IsSessionStateEstablished", session.IsSessionStateEstablished})
    session.EntityData.Leafs.Append("is-session-locally-initiated", types.YLeaf{"IsSessionLocallyInitiated", session.IsSessionLocallyInitiated})
    session.EntityData.Leafs.Append("is-conditional-debug-enabled", types.YLeaf{"IsConditionalDebugEnabled", session.IsConditionalDebugEnabled})
    session.EntityData.Leafs.Append("unique-id", types.YLeaf{"UniqueId", session.UniqueId})
    session.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", session.InterfaceName})

    session.EntityData.YListKeys = []string {"LocalTunnelId", "LocalSessionId"}

    return &(session.EntityData)
}

// L2tpv2_Sessions_Session_SessionApplicationData
// Session application data
type L2tpv2_Sessions_Session_SessionApplicationData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // l2tp sh sess app type. The type is interface{} with range: 0..4294967295.
    L2tpShSessAppType interface{}

    // Xconnect data.
    Xconnect L2tpv2_Sessions_Session_SessionApplicationData_Xconnect

    // VPDN data.
    Vpdn L2tpv2_Sessions_Session_SessionApplicationData_Vpdn
}

func (sessionApplicationData *L2tpv2_Sessions_Session_SessionApplicationData) GetEntityData() *types.CommonEntityData {
    sessionApplicationData.EntityData.YFilter = sessionApplicationData.YFilter
    sessionApplicationData.EntityData.YangName = "session-application-data"
    sessionApplicationData.EntityData.BundleName = "cisco_ios_xr"
    sessionApplicationData.EntityData.ParentYangName = "session"
    sessionApplicationData.EntityData.SegmentPath = "session-application-data"
    sessionApplicationData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionApplicationData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionApplicationData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionApplicationData.EntityData.Children = types.NewOrderedMap()
    sessionApplicationData.EntityData.Children.Append("xconnect", types.YChild{"Xconnect", &sessionApplicationData.Xconnect})
    sessionApplicationData.EntityData.Children.Append("vpdn", types.YChild{"Vpdn", &sessionApplicationData.Vpdn})
    sessionApplicationData.EntityData.Leafs = types.NewOrderedMap()
    sessionApplicationData.EntityData.Leafs.Append("l2tp-sh-sess-app-type", types.YLeaf{"L2tpShSessAppType", sessionApplicationData.L2tpShSessAppType})

    sessionApplicationData.EntityData.YListKeys = []string {}

    return &(sessionApplicationData.EntityData)
}

// L2tpv2_Sessions_Session_SessionApplicationData_Xconnect
// Xconnect data
type L2tpv2_Sessions_Session_SessionApplicationData_Xconnect struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Circuit name. The type is string.
    CircuitName interface{}

    // Session VC ID. The type is interface{} with range: 0..4294967295.
    SessionvcId interface{}

    // True if circuit state is up. The type is bool.
    IsCircuitStateUp interface{}

    // True if local circuit state is up. The type is bool.
    IsLocalCircuitStateUp interface{}

    // True if remote circuit state is up. The type is bool.
    IsRemoteCircuitStateUp interface{}

    // IPv6ProtocolTunneling. The type is bool.
    Ipv6ProtocolTunneling interface{}
}

func (xconnect *L2tpv2_Sessions_Session_SessionApplicationData_Xconnect) GetEntityData() *types.CommonEntityData {
    xconnect.EntityData.YFilter = xconnect.YFilter
    xconnect.EntityData.YangName = "xconnect"
    xconnect.EntityData.BundleName = "cisco_ios_xr"
    xconnect.EntityData.ParentYangName = "session-application-data"
    xconnect.EntityData.SegmentPath = "xconnect"
    xconnect.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    xconnect.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    xconnect.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    xconnect.EntityData.Children = types.NewOrderedMap()
    xconnect.EntityData.Leafs = types.NewOrderedMap()
    xconnect.EntityData.Leafs.Append("circuit-name", types.YLeaf{"CircuitName", xconnect.CircuitName})
    xconnect.EntityData.Leafs.Append("sessionvc-id", types.YLeaf{"SessionvcId", xconnect.SessionvcId})
    xconnect.EntityData.Leafs.Append("is-circuit-state-up", types.YLeaf{"IsCircuitStateUp", xconnect.IsCircuitStateUp})
    xconnect.EntityData.Leafs.Append("is-local-circuit-state-up", types.YLeaf{"IsLocalCircuitStateUp", xconnect.IsLocalCircuitStateUp})
    xconnect.EntityData.Leafs.Append("is-remote-circuit-state-up", types.YLeaf{"IsRemoteCircuitStateUp", xconnect.IsRemoteCircuitStateUp})
    xconnect.EntityData.Leafs.Append("ipv6-protocol-tunneling", types.YLeaf{"Ipv6ProtocolTunneling", xconnect.Ipv6ProtocolTunneling})

    xconnect.EntityData.YListKeys = []string {}

    return &(xconnect.EntityData)
}

// L2tpv2_Sessions_Session_SessionApplicationData_Vpdn
// VPDN data
type L2tpv2_Sessions_Session_SessionApplicationData_Vpdn struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Session username. The type is string.
    Username interface{}

    // Interface name. The type is string with pattern: [a-zA-Z0-9._/-]+.
    InterfaceName interface{}
}

func (vpdn *L2tpv2_Sessions_Session_SessionApplicationData_Vpdn) GetEntityData() *types.CommonEntityData {
    vpdn.EntityData.YFilter = vpdn.YFilter
    vpdn.EntityData.YangName = "vpdn"
    vpdn.EntityData.BundleName = "cisco_ios_xr"
    vpdn.EntityData.ParentYangName = "session-application-data"
    vpdn.EntityData.SegmentPath = "vpdn"
    vpdn.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vpdn.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vpdn.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vpdn.EntityData.Children = types.NewOrderedMap()
    vpdn.EntityData.Leafs = types.NewOrderedMap()
    vpdn.EntityData.Leafs.Append("username", types.YLeaf{"Username", vpdn.Username})
    vpdn.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", vpdn.InterfaceName})

    vpdn.EntityData.YListKeys = []string {}

    return &(vpdn.EntityData)
}

// L2tpv2_Session
// L2TP control messages counters
type L2tpv2_Session struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2TP session unavailable  information.
    Unavailable L2tpv2_Session_Unavailable
}

func (session *L2tpv2_Session) GetEntityData() *types.CommonEntityData {
    session.EntityData.YFilter = session.YFilter
    session.EntityData.YangName = "session"
    session.EntityData.BundleName = "cisco_ios_xr"
    session.EntityData.ParentYangName = "l2tpv2"
    session.EntityData.SegmentPath = "session"
    session.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    session.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    session.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    session.EntityData.Children = types.NewOrderedMap()
    session.EntityData.Children.Append("unavailable", types.YChild{"Unavailable", &session.Unavailable})
    session.EntityData.Leafs = types.NewOrderedMap()

    session.EntityData.YListKeys = []string {}

    return &(session.EntityData)
}

// L2tpv2_Session_Unavailable
// L2TP session unavailable  information
type L2tpv2_Session_Unavailable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of session ID in hold database. The type is interface{} with range:
    // 0..4294967295.
    SessionsOnHold interface{}
}

func (unavailable *L2tpv2_Session_Unavailable) GetEntityData() *types.CommonEntityData {
    unavailable.EntityData.YFilter = unavailable.YFilter
    unavailable.EntityData.YangName = "unavailable"
    unavailable.EntityData.BundleName = "cisco_ios_xr"
    unavailable.EntityData.ParentYangName = "session"
    unavailable.EntityData.SegmentPath = "unavailable"
    unavailable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unavailable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unavailable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unavailable.EntityData.Children = types.NewOrderedMap()
    unavailable.EntityData.Leafs = types.NewOrderedMap()
    unavailable.EntityData.Leafs.Append("sessions-on-hold", types.YLeaf{"SessionsOnHold", unavailable.SessionsOnHold})

    unavailable.EntityData.YListKeys = []string {}

    return &(unavailable.EntityData)
}

