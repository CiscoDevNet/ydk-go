// This module contains a collection of YANG definitions
// for Cisco IOS-XR tunnel-l2tun package operational data.
// 
// This module contains definitions
// for the following management objects:
//   l2tp: L2TP operational data
//   l2tpv2: l2tpv2
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-tunnel-l2tun-oper l2tp}", reflect.TypeOf(L2Tp{}))
    ydk.RegisterEntity("Cisco-IOS-XR-tunnel-l2tun-oper:l2tp", reflect.TypeOf(L2Tp{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-tunnel-l2tun-oper l2tpv2}", reflect.TypeOf(L2Tpv2{}))
    ydk.RegisterEntity("Cisco-IOS-XR-tunnel-l2tun-oper:l2tpv2", reflect.TypeOf(L2Tpv2{}))
}

// DigestHash represents Digest hash types
type DigestHash string

const (
    // MD5
    DigestHash_md5 DigestHash = "md5"

    // SHA1
    DigestHash_sha1 DigestHash = "sha1"
)

// L2Tp
// L2TP operational data
type L2Tp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2TP control messages counters.
    Counters L2Tp_Counters

    // List of tunnel IDs.
    TunnelConfigurations L2Tp_TunnelConfigurations

    // Failure events leading to disconnection.
    CounterHistFail L2Tp_CounterHistFail

    // List of L2TP class names.
    Classes L2Tp_Classes

    // List of tunnel IDs.
    Tunnels L2Tp_Tunnels

    // List of session IDs.
    Sessions L2Tp_Sessions

    // L2TP control messages counters.
    Session L2Tp_Session
}

func (l2Tp *L2Tp) GetEntityData() *types.CommonEntityData {
    l2Tp.EntityData.YFilter = l2Tp.YFilter
    l2Tp.EntityData.YangName = "l2tp"
    l2Tp.EntityData.BundleName = "cisco_ios_xr"
    l2Tp.EntityData.ParentYangName = "Cisco-IOS-XR-tunnel-l2tun-oper"
    l2Tp.EntityData.SegmentPath = "Cisco-IOS-XR-tunnel-l2tun-oper:l2tp"
    l2Tp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    l2Tp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    l2Tp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    l2Tp.EntityData.Children = make(map[string]types.YChild)
    l2Tp.EntityData.Children["counters"] = types.YChild{"Counters", &l2Tp.Counters}
    l2Tp.EntityData.Children["tunnel-configurations"] = types.YChild{"TunnelConfigurations", &l2Tp.TunnelConfigurations}
    l2Tp.EntityData.Children["counter-hist-fail"] = types.YChild{"CounterHistFail", &l2Tp.CounterHistFail}
    l2Tp.EntityData.Children["classes"] = types.YChild{"Classes", &l2Tp.Classes}
    l2Tp.EntityData.Children["tunnels"] = types.YChild{"Tunnels", &l2Tp.Tunnels}
    l2Tp.EntityData.Children["sessions"] = types.YChild{"Sessions", &l2Tp.Sessions}
    l2Tp.EntityData.Children["session"] = types.YChild{"Session", &l2Tp.Session}
    l2Tp.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(l2Tp.EntityData)
}

// L2Tp_Counters
// L2TP control messages counters
type L2Tp_Counters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2TP control messages counters.
    Control L2Tp_Counters_Control
}

func (counters *L2Tp_Counters) GetEntityData() *types.CommonEntityData {
    counters.EntityData.YFilter = counters.YFilter
    counters.EntityData.YangName = "counters"
    counters.EntityData.BundleName = "cisco_ios_xr"
    counters.EntityData.ParentYangName = "l2tp"
    counters.EntityData.SegmentPath = "counters"
    counters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    counters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    counters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    counters.EntityData.Children = make(map[string]types.YChild)
    counters.EntityData.Children["control"] = types.YChild{"Control", &counters.Control}
    counters.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(counters.EntityData)
}

// L2Tp_Counters_Control
// L2TP control messages counters
type L2Tp_Counters_Control struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2TP control tunnel messages counters.
    TunnelXr L2Tp_Counters_Control_TunnelXr

    // Table of tunnel IDs of control message counters.
    Tunnels L2Tp_Counters_Control_Tunnels
}

func (control *L2Tp_Counters_Control) GetEntityData() *types.CommonEntityData {
    control.EntityData.YFilter = control.YFilter
    control.EntityData.YangName = "control"
    control.EntityData.BundleName = "cisco_ios_xr"
    control.EntityData.ParentYangName = "counters"
    control.EntityData.SegmentPath = "control"
    control.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    control.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    control.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    control.EntityData.Children = make(map[string]types.YChild)
    control.EntityData.Children["tunnel-xr"] = types.YChild{"TunnelXr", &control.TunnelXr}
    control.EntityData.Children["tunnels"] = types.YChild{"Tunnels", &control.Tunnels}
    control.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(control.EntityData)
}

// L2Tp_Counters_Control_TunnelXr
// L2TP control tunnel messages counters
type L2Tp_Counters_Control_TunnelXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tunnel authentication counters.
    Authentication L2Tp_Counters_Control_TunnelXr_Authentication

    // Tunnel counters.
    Global L2Tp_Counters_Control_TunnelXr_Global
}

func (tunnelXr *L2Tp_Counters_Control_TunnelXr) GetEntityData() *types.CommonEntityData {
    tunnelXr.EntityData.YFilter = tunnelXr.YFilter
    tunnelXr.EntityData.YangName = "tunnel-xr"
    tunnelXr.EntityData.BundleName = "cisco_ios_xr"
    tunnelXr.EntityData.ParentYangName = "control"
    tunnelXr.EntityData.SegmentPath = "tunnel-xr"
    tunnelXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnelXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnelXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnelXr.EntityData.Children = make(map[string]types.YChild)
    tunnelXr.EntityData.Children["authentication"] = types.YChild{"Authentication", &tunnelXr.Authentication}
    tunnelXr.EntityData.Children["global"] = types.YChild{"Global", &tunnelXr.Global}
    tunnelXr.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(tunnelXr.EntityData)
}

// L2Tp_Counters_Control_TunnelXr_Authentication
// Tunnel authentication counters
type L2Tp_Counters_Control_TunnelXr_Authentication struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Nonce AVP statistics.
    NonceAvp L2Tp_Counters_Control_TunnelXr_Authentication_NonceAvp

    // Common digest statistics.
    CommonDigest L2Tp_Counters_Control_TunnelXr_Authentication_CommonDigest

    // Primary digest statistics.
    PrimaryDigest L2Tp_Counters_Control_TunnelXr_Authentication_PrimaryDigest

    // Secondary digest statistics.
    SecondaryDigest L2Tp_Counters_Control_TunnelXr_Authentication_SecondaryDigest

    // Integrity check statistics.
    IntegrityCheck L2Tp_Counters_Control_TunnelXr_Authentication_IntegrityCheck

    // Local secret statistics.
    LocalSecret L2Tp_Counters_Control_TunnelXr_Authentication_LocalSecret

    // Challenge AVP statistics.
    ChallengeAvp L2Tp_Counters_Control_TunnelXr_Authentication_ChallengeAvp

    // Challenge response statistics.
    ChallengeReponse L2Tp_Counters_Control_TunnelXr_Authentication_ChallengeReponse

    // Overall statistics.
    OverallStatistics L2Tp_Counters_Control_TunnelXr_Authentication_OverallStatistics
}

func (authentication *L2Tp_Counters_Control_TunnelXr_Authentication) GetEntityData() *types.CommonEntityData {
    authentication.EntityData.YFilter = authentication.YFilter
    authentication.EntityData.YangName = "authentication"
    authentication.EntityData.BundleName = "cisco_ios_xr"
    authentication.EntityData.ParentYangName = "tunnel-xr"
    authentication.EntityData.SegmentPath = "authentication"
    authentication.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authentication.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authentication.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authentication.EntityData.Children = make(map[string]types.YChild)
    authentication.EntityData.Children["nonce-avp"] = types.YChild{"NonceAvp", &authentication.NonceAvp}
    authentication.EntityData.Children["common-digest"] = types.YChild{"CommonDigest", &authentication.CommonDigest}
    authentication.EntityData.Children["primary-digest"] = types.YChild{"PrimaryDigest", &authentication.PrimaryDigest}
    authentication.EntityData.Children["secondary-digest"] = types.YChild{"SecondaryDigest", &authentication.SecondaryDigest}
    authentication.EntityData.Children["integrity-check"] = types.YChild{"IntegrityCheck", &authentication.IntegrityCheck}
    authentication.EntityData.Children["local-secret"] = types.YChild{"LocalSecret", &authentication.LocalSecret}
    authentication.EntityData.Children["challenge-avp"] = types.YChild{"ChallengeAvp", &authentication.ChallengeAvp}
    authentication.EntityData.Children["challenge-reponse"] = types.YChild{"ChallengeReponse", &authentication.ChallengeReponse}
    authentication.EntityData.Children["overall-statistics"] = types.YChild{"OverallStatistics", &authentication.OverallStatistics}
    authentication.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(authentication.EntityData)
}

// L2Tp_Counters_Control_TunnelXr_Authentication_NonceAvp
// Nonce AVP statistics
type L2Tp_Counters_Control_TunnelXr_Authentication_NonceAvp struct {
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

func (nonceAvp *L2Tp_Counters_Control_TunnelXr_Authentication_NonceAvp) GetEntityData() *types.CommonEntityData {
    nonceAvp.EntityData.YFilter = nonceAvp.YFilter
    nonceAvp.EntityData.YangName = "nonce-avp"
    nonceAvp.EntityData.BundleName = "cisco_ios_xr"
    nonceAvp.EntityData.ParentYangName = "authentication"
    nonceAvp.EntityData.SegmentPath = "nonce-avp"
    nonceAvp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nonceAvp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nonceAvp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nonceAvp.EntityData.Children = make(map[string]types.YChild)
    nonceAvp.EntityData.Leafs = make(map[string]types.YLeaf)
    nonceAvp.EntityData.Leafs["validate"] = types.YLeaf{"Validate", nonceAvp.Validate}
    nonceAvp.EntityData.Leafs["bad-hash"] = types.YLeaf{"BadHash", nonceAvp.BadHash}
    nonceAvp.EntityData.Leafs["bad-length"] = types.YLeaf{"BadLength", nonceAvp.BadLength}
    nonceAvp.EntityData.Leafs["ignored"] = types.YLeaf{"Ignored", nonceAvp.Ignored}
    nonceAvp.EntityData.Leafs["missing"] = types.YLeaf{"Missing", nonceAvp.Missing}
    nonceAvp.EntityData.Leafs["passed"] = types.YLeaf{"Passed", nonceAvp.Passed}
    nonceAvp.EntityData.Leafs["failed"] = types.YLeaf{"Failed", nonceAvp.Failed}
    nonceAvp.EntityData.Leafs["skipped"] = types.YLeaf{"Skipped", nonceAvp.Skipped}
    nonceAvp.EntityData.Leafs["generate-response-failures"] = types.YLeaf{"GenerateResponseFailures", nonceAvp.GenerateResponseFailures}
    nonceAvp.EntityData.Leafs["unexpected"] = types.YLeaf{"Unexpected", nonceAvp.Unexpected}
    nonceAvp.EntityData.Leafs["unexpected-zlb"] = types.YLeaf{"UnexpectedZlb", nonceAvp.UnexpectedZlb}
    return &(nonceAvp.EntityData)
}

// L2Tp_Counters_Control_TunnelXr_Authentication_CommonDigest
// Common digest statistics
type L2Tp_Counters_Control_TunnelXr_Authentication_CommonDigest struct {
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

func (commonDigest *L2Tp_Counters_Control_TunnelXr_Authentication_CommonDigest) GetEntityData() *types.CommonEntityData {
    commonDigest.EntityData.YFilter = commonDigest.YFilter
    commonDigest.EntityData.YangName = "common-digest"
    commonDigest.EntityData.BundleName = "cisco_ios_xr"
    commonDigest.EntityData.ParentYangName = "authentication"
    commonDigest.EntityData.SegmentPath = "common-digest"
    commonDigest.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    commonDigest.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    commonDigest.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    commonDigest.EntityData.Children = make(map[string]types.YChild)
    commonDigest.EntityData.Leafs = make(map[string]types.YLeaf)
    commonDigest.EntityData.Leafs["validate"] = types.YLeaf{"Validate", commonDigest.Validate}
    commonDigest.EntityData.Leafs["bad-hash"] = types.YLeaf{"BadHash", commonDigest.BadHash}
    commonDigest.EntityData.Leafs["bad-length"] = types.YLeaf{"BadLength", commonDigest.BadLength}
    commonDigest.EntityData.Leafs["ignored"] = types.YLeaf{"Ignored", commonDigest.Ignored}
    commonDigest.EntityData.Leafs["missing"] = types.YLeaf{"Missing", commonDigest.Missing}
    commonDigest.EntityData.Leafs["passed"] = types.YLeaf{"Passed", commonDigest.Passed}
    commonDigest.EntityData.Leafs["failed"] = types.YLeaf{"Failed", commonDigest.Failed}
    commonDigest.EntityData.Leafs["skipped"] = types.YLeaf{"Skipped", commonDigest.Skipped}
    commonDigest.EntityData.Leafs["generate-response-failures"] = types.YLeaf{"GenerateResponseFailures", commonDigest.GenerateResponseFailures}
    commonDigest.EntityData.Leafs["unexpected"] = types.YLeaf{"Unexpected", commonDigest.Unexpected}
    commonDigest.EntityData.Leafs["unexpected-zlb"] = types.YLeaf{"UnexpectedZlb", commonDigest.UnexpectedZlb}
    return &(commonDigest.EntityData)
}

// L2Tp_Counters_Control_TunnelXr_Authentication_PrimaryDigest
// Primary digest statistics
type L2Tp_Counters_Control_TunnelXr_Authentication_PrimaryDigest struct {
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

func (primaryDigest *L2Tp_Counters_Control_TunnelXr_Authentication_PrimaryDigest) GetEntityData() *types.CommonEntityData {
    primaryDigest.EntityData.YFilter = primaryDigest.YFilter
    primaryDigest.EntityData.YangName = "primary-digest"
    primaryDigest.EntityData.BundleName = "cisco_ios_xr"
    primaryDigest.EntityData.ParentYangName = "authentication"
    primaryDigest.EntityData.SegmentPath = "primary-digest"
    primaryDigest.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    primaryDigest.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    primaryDigest.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    primaryDigest.EntityData.Children = make(map[string]types.YChild)
    primaryDigest.EntityData.Leafs = make(map[string]types.YLeaf)
    primaryDigest.EntityData.Leafs["validate"] = types.YLeaf{"Validate", primaryDigest.Validate}
    primaryDigest.EntityData.Leafs["bad-hash"] = types.YLeaf{"BadHash", primaryDigest.BadHash}
    primaryDigest.EntityData.Leafs["bad-length"] = types.YLeaf{"BadLength", primaryDigest.BadLength}
    primaryDigest.EntityData.Leafs["ignored"] = types.YLeaf{"Ignored", primaryDigest.Ignored}
    primaryDigest.EntityData.Leafs["missing"] = types.YLeaf{"Missing", primaryDigest.Missing}
    primaryDigest.EntityData.Leafs["passed"] = types.YLeaf{"Passed", primaryDigest.Passed}
    primaryDigest.EntityData.Leafs["failed"] = types.YLeaf{"Failed", primaryDigest.Failed}
    primaryDigest.EntityData.Leafs["skipped"] = types.YLeaf{"Skipped", primaryDigest.Skipped}
    primaryDigest.EntityData.Leafs["generate-response-failures"] = types.YLeaf{"GenerateResponseFailures", primaryDigest.GenerateResponseFailures}
    primaryDigest.EntityData.Leafs["unexpected"] = types.YLeaf{"Unexpected", primaryDigest.Unexpected}
    primaryDigest.EntityData.Leafs["unexpected-zlb"] = types.YLeaf{"UnexpectedZlb", primaryDigest.UnexpectedZlb}
    return &(primaryDigest.EntityData)
}

// L2Tp_Counters_Control_TunnelXr_Authentication_SecondaryDigest
// Secondary digest statistics
type L2Tp_Counters_Control_TunnelXr_Authentication_SecondaryDigest struct {
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

func (secondaryDigest *L2Tp_Counters_Control_TunnelXr_Authentication_SecondaryDigest) GetEntityData() *types.CommonEntityData {
    secondaryDigest.EntityData.YFilter = secondaryDigest.YFilter
    secondaryDigest.EntityData.YangName = "secondary-digest"
    secondaryDigest.EntityData.BundleName = "cisco_ios_xr"
    secondaryDigest.EntityData.ParentYangName = "authentication"
    secondaryDigest.EntityData.SegmentPath = "secondary-digest"
    secondaryDigest.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    secondaryDigest.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    secondaryDigest.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    secondaryDigest.EntityData.Children = make(map[string]types.YChild)
    secondaryDigest.EntityData.Leafs = make(map[string]types.YLeaf)
    secondaryDigest.EntityData.Leafs["validate"] = types.YLeaf{"Validate", secondaryDigest.Validate}
    secondaryDigest.EntityData.Leafs["bad-hash"] = types.YLeaf{"BadHash", secondaryDigest.BadHash}
    secondaryDigest.EntityData.Leafs["bad-length"] = types.YLeaf{"BadLength", secondaryDigest.BadLength}
    secondaryDigest.EntityData.Leafs["ignored"] = types.YLeaf{"Ignored", secondaryDigest.Ignored}
    secondaryDigest.EntityData.Leafs["missing"] = types.YLeaf{"Missing", secondaryDigest.Missing}
    secondaryDigest.EntityData.Leafs["passed"] = types.YLeaf{"Passed", secondaryDigest.Passed}
    secondaryDigest.EntityData.Leafs["failed"] = types.YLeaf{"Failed", secondaryDigest.Failed}
    secondaryDigest.EntityData.Leafs["skipped"] = types.YLeaf{"Skipped", secondaryDigest.Skipped}
    secondaryDigest.EntityData.Leafs["generate-response-failures"] = types.YLeaf{"GenerateResponseFailures", secondaryDigest.GenerateResponseFailures}
    secondaryDigest.EntityData.Leafs["unexpected"] = types.YLeaf{"Unexpected", secondaryDigest.Unexpected}
    secondaryDigest.EntityData.Leafs["unexpected-zlb"] = types.YLeaf{"UnexpectedZlb", secondaryDigest.UnexpectedZlb}
    return &(secondaryDigest.EntityData)
}

// L2Tp_Counters_Control_TunnelXr_Authentication_IntegrityCheck
// Integrity check statistics
type L2Tp_Counters_Control_TunnelXr_Authentication_IntegrityCheck struct {
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

func (integrityCheck *L2Tp_Counters_Control_TunnelXr_Authentication_IntegrityCheck) GetEntityData() *types.CommonEntityData {
    integrityCheck.EntityData.YFilter = integrityCheck.YFilter
    integrityCheck.EntityData.YangName = "integrity-check"
    integrityCheck.EntityData.BundleName = "cisco_ios_xr"
    integrityCheck.EntityData.ParentYangName = "authentication"
    integrityCheck.EntityData.SegmentPath = "integrity-check"
    integrityCheck.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    integrityCheck.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    integrityCheck.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    integrityCheck.EntityData.Children = make(map[string]types.YChild)
    integrityCheck.EntityData.Leafs = make(map[string]types.YLeaf)
    integrityCheck.EntityData.Leafs["validate"] = types.YLeaf{"Validate", integrityCheck.Validate}
    integrityCheck.EntityData.Leafs["bad-hash"] = types.YLeaf{"BadHash", integrityCheck.BadHash}
    integrityCheck.EntityData.Leafs["bad-length"] = types.YLeaf{"BadLength", integrityCheck.BadLength}
    integrityCheck.EntityData.Leafs["ignored"] = types.YLeaf{"Ignored", integrityCheck.Ignored}
    integrityCheck.EntityData.Leafs["missing"] = types.YLeaf{"Missing", integrityCheck.Missing}
    integrityCheck.EntityData.Leafs["passed"] = types.YLeaf{"Passed", integrityCheck.Passed}
    integrityCheck.EntityData.Leafs["failed"] = types.YLeaf{"Failed", integrityCheck.Failed}
    integrityCheck.EntityData.Leafs["skipped"] = types.YLeaf{"Skipped", integrityCheck.Skipped}
    integrityCheck.EntityData.Leafs["generate-response-failures"] = types.YLeaf{"GenerateResponseFailures", integrityCheck.GenerateResponseFailures}
    integrityCheck.EntityData.Leafs["unexpected"] = types.YLeaf{"Unexpected", integrityCheck.Unexpected}
    integrityCheck.EntityData.Leafs["unexpected-zlb"] = types.YLeaf{"UnexpectedZlb", integrityCheck.UnexpectedZlb}
    return &(integrityCheck.EntityData)
}

// L2Tp_Counters_Control_TunnelXr_Authentication_LocalSecret
// Local secret statistics
type L2Tp_Counters_Control_TunnelXr_Authentication_LocalSecret struct {
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

func (localSecret *L2Tp_Counters_Control_TunnelXr_Authentication_LocalSecret) GetEntityData() *types.CommonEntityData {
    localSecret.EntityData.YFilter = localSecret.YFilter
    localSecret.EntityData.YangName = "local-secret"
    localSecret.EntityData.BundleName = "cisco_ios_xr"
    localSecret.EntityData.ParentYangName = "authentication"
    localSecret.EntityData.SegmentPath = "local-secret"
    localSecret.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localSecret.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localSecret.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localSecret.EntityData.Children = make(map[string]types.YChild)
    localSecret.EntityData.Leafs = make(map[string]types.YLeaf)
    localSecret.EntityData.Leafs["validate"] = types.YLeaf{"Validate", localSecret.Validate}
    localSecret.EntityData.Leafs["bad-hash"] = types.YLeaf{"BadHash", localSecret.BadHash}
    localSecret.EntityData.Leafs["bad-length"] = types.YLeaf{"BadLength", localSecret.BadLength}
    localSecret.EntityData.Leafs["ignored"] = types.YLeaf{"Ignored", localSecret.Ignored}
    localSecret.EntityData.Leafs["missing"] = types.YLeaf{"Missing", localSecret.Missing}
    localSecret.EntityData.Leafs["passed"] = types.YLeaf{"Passed", localSecret.Passed}
    localSecret.EntityData.Leafs["failed"] = types.YLeaf{"Failed", localSecret.Failed}
    localSecret.EntityData.Leafs["skipped"] = types.YLeaf{"Skipped", localSecret.Skipped}
    localSecret.EntityData.Leafs["generate-response-failures"] = types.YLeaf{"GenerateResponseFailures", localSecret.GenerateResponseFailures}
    localSecret.EntityData.Leafs["unexpected"] = types.YLeaf{"Unexpected", localSecret.Unexpected}
    localSecret.EntityData.Leafs["unexpected-zlb"] = types.YLeaf{"UnexpectedZlb", localSecret.UnexpectedZlb}
    return &(localSecret.EntityData)
}

// L2Tp_Counters_Control_TunnelXr_Authentication_ChallengeAvp
// Challenge AVP statistics
type L2Tp_Counters_Control_TunnelXr_Authentication_ChallengeAvp struct {
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

func (challengeAvp *L2Tp_Counters_Control_TunnelXr_Authentication_ChallengeAvp) GetEntityData() *types.CommonEntityData {
    challengeAvp.EntityData.YFilter = challengeAvp.YFilter
    challengeAvp.EntityData.YangName = "challenge-avp"
    challengeAvp.EntityData.BundleName = "cisco_ios_xr"
    challengeAvp.EntityData.ParentYangName = "authentication"
    challengeAvp.EntityData.SegmentPath = "challenge-avp"
    challengeAvp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    challengeAvp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    challengeAvp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    challengeAvp.EntityData.Children = make(map[string]types.YChild)
    challengeAvp.EntityData.Leafs = make(map[string]types.YLeaf)
    challengeAvp.EntityData.Leafs["validate"] = types.YLeaf{"Validate", challengeAvp.Validate}
    challengeAvp.EntityData.Leafs["bad-hash"] = types.YLeaf{"BadHash", challengeAvp.BadHash}
    challengeAvp.EntityData.Leafs["bad-length"] = types.YLeaf{"BadLength", challengeAvp.BadLength}
    challengeAvp.EntityData.Leafs["ignored"] = types.YLeaf{"Ignored", challengeAvp.Ignored}
    challengeAvp.EntityData.Leafs["missing"] = types.YLeaf{"Missing", challengeAvp.Missing}
    challengeAvp.EntityData.Leafs["passed"] = types.YLeaf{"Passed", challengeAvp.Passed}
    challengeAvp.EntityData.Leafs["failed"] = types.YLeaf{"Failed", challengeAvp.Failed}
    challengeAvp.EntityData.Leafs["skipped"] = types.YLeaf{"Skipped", challengeAvp.Skipped}
    challengeAvp.EntityData.Leafs["generate-response-failures"] = types.YLeaf{"GenerateResponseFailures", challengeAvp.GenerateResponseFailures}
    challengeAvp.EntityData.Leafs["unexpected"] = types.YLeaf{"Unexpected", challengeAvp.Unexpected}
    challengeAvp.EntityData.Leafs["unexpected-zlb"] = types.YLeaf{"UnexpectedZlb", challengeAvp.UnexpectedZlb}
    return &(challengeAvp.EntityData)
}

// L2Tp_Counters_Control_TunnelXr_Authentication_ChallengeReponse
// Challenge response statistics
type L2Tp_Counters_Control_TunnelXr_Authentication_ChallengeReponse struct {
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

func (challengeReponse *L2Tp_Counters_Control_TunnelXr_Authentication_ChallengeReponse) GetEntityData() *types.CommonEntityData {
    challengeReponse.EntityData.YFilter = challengeReponse.YFilter
    challengeReponse.EntityData.YangName = "challenge-reponse"
    challengeReponse.EntityData.BundleName = "cisco_ios_xr"
    challengeReponse.EntityData.ParentYangName = "authentication"
    challengeReponse.EntityData.SegmentPath = "challenge-reponse"
    challengeReponse.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    challengeReponse.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    challengeReponse.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    challengeReponse.EntityData.Children = make(map[string]types.YChild)
    challengeReponse.EntityData.Leafs = make(map[string]types.YLeaf)
    challengeReponse.EntityData.Leafs["validate"] = types.YLeaf{"Validate", challengeReponse.Validate}
    challengeReponse.EntityData.Leafs["bad-hash"] = types.YLeaf{"BadHash", challengeReponse.BadHash}
    challengeReponse.EntityData.Leafs["bad-length"] = types.YLeaf{"BadLength", challengeReponse.BadLength}
    challengeReponse.EntityData.Leafs["ignored"] = types.YLeaf{"Ignored", challengeReponse.Ignored}
    challengeReponse.EntityData.Leafs["missing"] = types.YLeaf{"Missing", challengeReponse.Missing}
    challengeReponse.EntityData.Leafs["passed"] = types.YLeaf{"Passed", challengeReponse.Passed}
    challengeReponse.EntityData.Leafs["failed"] = types.YLeaf{"Failed", challengeReponse.Failed}
    challengeReponse.EntityData.Leafs["skipped"] = types.YLeaf{"Skipped", challengeReponse.Skipped}
    challengeReponse.EntityData.Leafs["generate-response-failures"] = types.YLeaf{"GenerateResponseFailures", challengeReponse.GenerateResponseFailures}
    challengeReponse.EntityData.Leafs["unexpected"] = types.YLeaf{"Unexpected", challengeReponse.Unexpected}
    challengeReponse.EntityData.Leafs["unexpected-zlb"] = types.YLeaf{"UnexpectedZlb", challengeReponse.UnexpectedZlb}
    return &(challengeReponse.EntityData)
}

// L2Tp_Counters_Control_TunnelXr_Authentication_OverallStatistics
// Overall statistics
type L2Tp_Counters_Control_TunnelXr_Authentication_OverallStatistics struct {
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

func (overallStatistics *L2Tp_Counters_Control_TunnelXr_Authentication_OverallStatistics) GetEntityData() *types.CommonEntityData {
    overallStatistics.EntityData.YFilter = overallStatistics.YFilter
    overallStatistics.EntityData.YangName = "overall-statistics"
    overallStatistics.EntityData.BundleName = "cisco_ios_xr"
    overallStatistics.EntityData.ParentYangName = "authentication"
    overallStatistics.EntityData.SegmentPath = "overall-statistics"
    overallStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    overallStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    overallStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    overallStatistics.EntityData.Children = make(map[string]types.YChild)
    overallStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    overallStatistics.EntityData.Leafs["validate"] = types.YLeaf{"Validate", overallStatistics.Validate}
    overallStatistics.EntityData.Leafs["bad-hash"] = types.YLeaf{"BadHash", overallStatistics.BadHash}
    overallStatistics.EntityData.Leafs["bad-length"] = types.YLeaf{"BadLength", overallStatistics.BadLength}
    overallStatistics.EntityData.Leafs["ignored"] = types.YLeaf{"Ignored", overallStatistics.Ignored}
    overallStatistics.EntityData.Leafs["missing"] = types.YLeaf{"Missing", overallStatistics.Missing}
    overallStatistics.EntityData.Leafs["passed"] = types.YLeaf{"Passed", overallStatistics.Passed}
    overallStatistics.EntityData.Leafs["failed"] = types.YLeaf{"Failed", overallStatistics.Failed}
    overallStatistics.EntityData.Leafs["skipped"] = types.YLeaf{"Skipped", overallStatistics.Skipped}
    overallStatistics.EntityData.Leafs["generate-response-failures"] = types.YLeaf{"GenerateResponseFailures", overallStatistics.GenerateResponseFailures}
    overallStatistics.EntityData.Leafs["unexpected"] = types.YLeaf{"Unexpected", overallStatistics.Unexpected}
    overallStatistics.EntityData.Leafs["unexpected-zlb"] = types.YLeaf{"UnexpectedZlb", overallStatistics.UnexpectedZlb}
    return &(overallStatistics.EntityData)
}

// L2Tp_Counters_Control_TunnelXr_Global
// Tunnel counters
type L2Tp_Counters_Control_TunnelXr_Global struct {
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
    Transmit L2Tp_Counters_Control_TunnelXr_Global_Transmit

    // Re transmit data.
    Retransmit L2Tp_Counters_Control_TunnelXr_Global_Retransmit

    // Received data.
    Received L2Tp_Counters_Control_TunnelXr_Global_Received

    // Drop data.
    Drop L2Tp_Counters_Control_TunnelXr_Global_Drop
}

func (global *L2Tp_Counters_Control_TunnelXr_Global) GetEntityData() *types.CommonEntityData {
    global.EntityData.YFilter = global.YFilter
    global.EntityData.YangName = "global"
    global.EntityData.BundleName = "cisco_ios_xr"
    global.EntityData.ParentYangName = "tunnel-xr"
    global.EntityData.SegmentPath = "global"
    global.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    global.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    global.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    global.EntityData.Children = make(map[string]types.YChild)
    global.EntityData.Children["transmit"] = types.YChild{"Transmit", &global.Transmit}
    global.EntityData.Children["retransmit"] = types.YChild{"Retransmit", &global.Retransmit}
    global.EntityData.Children["received"] = types.YChild{"Received", &global.Received}
    global.EntityData.Children["drop"] = types.YChild{"Drop", &global.Drop}
    global.EntityData.Leafs = make(map[string]types.YLeaf)
    global.EntityData.Leafs["total-transmit"] = types.YLeaf{"TotalTransmit", global.TotalTransmit}
    global.EntityData.Leafs["total-retransmit"] = types.YLeaf{"TotalRetransmit", global.TotalRetransmit}
    global.EntityData.Leafs["total-received"] = types.YLeaf{"TotalReceived", global.TotalReceived}
    global.EntityData.Leafs["total-drop"] = types.YLeaf{"TotalDrop", global.TotalDrop}
    return &(global.EntityData)
}

// L2Tp_Counters_Control_TunnelXr_Global_Transmit
// Transmit data
type L2Tp_Counters_Control_TunnelXr_Global_Transmit struct {
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

func (transmit *L2Tp_Counters_Control_TunnelXr_Global_Transmit) GetEntityData() *types.CommonEntityData {
    transmit.EntityData.YFilter = transmit.YFilter
    transmit.EntityData.YangName = "transmit"
    transmit.EntityData.BundleName = "cisco_ios_xr"
    transmit.EntityData.ParentYangName = "global"
    transmit.EntityData.SegmentPath = "transmit"
    transmit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transmit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transmit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transmit.EntityData.Children = make(map[string]types.YChild)
    transmit.EntityData.Leafs = make(map[string]types.YLeaf)
    transmit.EntityData.Leafs["unknown-packets"] = types.YLeaf{"UnknownPackets", transmit.UnknownPackets}
    transmit.EntityData.Leafs["zero-length-body-packets"] = types.YLeaf{"ZeroLengthBodyPackets", transmit.ZeroLengthBodyPackets}
    transmit.EntityData.Leafs["start-control-connection-requests"] = types.YLeaf{"StartControlConnectionRequests", transmit.StartControlConnectionRequests}
    transmit.EntityData.Leafs["start-control-connection-replies"] = types.YLeaf{"StartControlConnectionReplies", transmit.StartControlConnectionReplies}
    transmit.EntityData.Leafs["start-control-connection-notifications"] = types.YLeaf{"StartControlConnectionNotifications", transmit.StartControlConnectionNotifications}
    transmit.EntityData.Leafs["stop-control-connection-notifications"] = types.YLeaf{"StopControlConnectionNotifications", transmit.StopControlConnectionNotifications}
    transmit.EntityData.Leafs["hello-packets"] = types.YLeaf{"HelloPackets", transmit.HelloPackets}
    transmit.EntityData.Leafs["outgoing-call-requests"] = types.YLeaf{"OutgoingCallRequests", transmit.OutgoingCallRequests}
    transmit.EntityData.Leafs["outgoing-call-replies"] = types.YLeaf{"OutgoingCallReplies", transmit.OutgoingCallReplies}
    transmit.EntityData.Leafs["outgoing-call-connected-packets"] = types.YLeaf{"OutgoingCallConnectedPackets", transmit.OutgoingCallConnectedPackets}
    transmit.EntityData.Leafs["incoming-call-requests"] = types.YLeaf{"IncomingCallRequests", transmit.IncomingCallRequests}
    transmit.EntityData.Leafs["incoming-call-replies"] = types.YLeaf{"IncomingCallReplies", transmit.IncomingCallReplies}
    transmit.EntityData.Leafs["incoming-call-connected-packets"] = types.YLeaf{"IncomingCallConnectedPackets", transmit.IncomingCallConnectedPackets}
    transmit.EntityData.Leafs["call-disconnect-notify-packets"] = types.YLeaf{"CallDisconnectNotifyPackets", transmit.CallDisconnectNotifyPackets}
    transmit.EntityData.Leafs["wan-error-notify-packets"] = types.YLeaf{"WanErrorNotifyPackets", transmit.WanErrorNotifyPackets}
    transmit.EntityData.Leafs["set-link-info-packets"] = types.YLeaf{"SetLinkInfoPackets", transmit.SetLinkInfoPackets}
    transmit.EntityData.Leafs["service-relay-requests"] = types.YLeaf{"ServiceRelayRequests", transmit.ServiceRelayRequests}
    transmit.EntityData.Leafs["service-relay-replies"] = types.YLeaf{"ServiceRelayReplies", transmit.ServiceRelayReplies}
    transmit.EntityData.Leafs["acknowledgement-packets"] = types.YLeaf{"AcknowledgementPackets", transmit.AcknowledgementPackets}
    return &(transmit.EntityData)
}

// L2Tp_Counters_Control_TunnelXr_Global_Retransmit
// Re transmit data
type L2Tp_Counters_Control_TunnelXr_Global_Retransmit struct {
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

func (retransmit *L2Tp_Counters_Control_TunnelXr_Global_Retransmit) GetEntityData() *types.CommonEntityData {
    retransmit.EntityData.YFilter = retransmit.YFilter
    retransmit.EntityData.YangName = "retransmit"
    retransmit.EntityData.BundleName = "cisco_ios_xr"
    retransmit.EntityData.ParentYangName = "global"
    retransmit.EntityData.SegmentPath = "retransmit"
    retransmit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    retransmit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    retransmit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    retransmit.EntityData.Children = make(map[string]types.YChild)
    retransmit.EntityData.Leafs = make(map[string]types.YLeaf)
    retransmit.EntityData.Leafs["unknown-packets"] = types.YLeaf{"UnknownPackets", retransmit.UnknownPackets}
    retransmit.EntityData.Leafs["zero-length-body-packets"] = types.YLeaf{"ZeroLengthBodyPackets", retransmit.ZeroLengthBodyPackets}
    retransmit.EntityData.Leafs["start-control-connection-requests"] = types.YLeaf{"StartControlConnectionRequests", retransmit.StartControlConnectionRequests}
    retransmit.EntityData.Leafs["start-control-connection-replies"] = types.YLeaf{"StartControlConnectionReplies", retransmit.StartControlConnectionReplies}
    retransmit.EntityData.Leafs["start-control-connection-notifications"] = types.YLeaf{"StartControlConnectionNotifications", retransmit.StartControlConnectionNotifications}
    retransmit.EntityData.Leafs["stop-control-connection-notifications"] = types.YLeaf{"StopControlConnectionNotifications", retransmit.StopControlConnectionNotifications}
    retransmit.EntityData.Leafs["hello-packets"] = types.YLeaf{"HelloPackets", retransmit.HelloPackets}
    retransmit.EntityData.Leafs["outgoing-call-requests"] = types.YLeaf{"OutgoingCallRequests", retransmit.OutgoingCallRequests}
    retransmit.EntityData.Leafs["outgoing-call-replies"] = types.YLeaf{"OutgoingCallReplies", retransmit.OutgoingCallReplies}
    retransmit.EntityData.Leafs["outgoing-call-connected-packets"] = types.YLeaf{"OutgoingCallConnectedPackets", retransmit.OutgoingCallConnectedPackets}
    retransmit.EntityData.Leafs["incoming-call-requests"] = types.YLeaf{"IncomingCallRequests", retransmit.IncomingCallRequests}
    retransmit.EntityData.Leafs["incoming-call-replies"] = types.YLeaf{"IncomingCallReplies", retransmit.IncomingCallReplies}
    retransmit.EntityData.Leafs["incoming-call-connected-packets"] = types.YLeaf{"IncomingCallConnectedPackets", retransmit.IncomingCallConnectedPackets}
    retransmit.EntityData.Leafs["call-disconnect-notify-packets"] = types.YLeaf{"CallDisconnectNotifyPackets", retransmit.CallDisconnectNotifyPackets}
    retransmit.EntityData.Leafs["wan-error-notify-packets"] = types.YLeaf{"WanErrorNotifyPackets", retransmit.WanErrorNotifyPackets}
    retransmit.EntityData.Leafs["set-link-info-packets"] = types.YLeaf{"SetLinkInfoPackets", retransmit.SetLinkInfoPackets}
    retransmit.EntityData.Leafs["service-relay-requests"] = types.YLeaf{"ServiceRelayRequests", retransmit.ServiceRelayRequests}
    retransmit.EntityData.Leafs["service-relay-replies"] = types.YLeaf{"ServiceRelayReplies", retransmit.ServiceRelayReplies}
    retransmit.EntityData.Leafs["acknowledgement-packets"] = types.YLeaf{"AcknowledgementPackets", retransmit.AcknowledgementPackets}
    return &(retransmit.EntityData)
}

// L2Tp_Counters_Control_TunnelXr_Global_Received
// Received data
type L2Tp_Counters_Control_TunnelXr_Global_Received struct {
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

func (received *L2Tp_Counters_Control_TunnelXr_Global_Received) GetEntityData() *types.CommonEntityData {
    received.EntityData.YFilter = received.YFilter
    received.EntityData.YangName = "received"
    received.EntityData.BundleName = "cisco_ios_xr"
    received.EntityData.ParentYangName = "global"
    received.EntityData.SegmentPath = "received"
    received.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    received.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    received.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    received.EntityData.Children = make(map[string]types.YChild)
    received.EntityData.Leafs = make(map[string]types.YLeaf)
    received.EntityData.Leafs["unknown-packets"] = types.YLeaf{"UnknownPackets", received.UnknownPackets}
    received.EntityData.Leafs["zero-length-body-packets"] = types.YLeaf{"ZeroLengthBodyPackets", received.ZeroLengthBodyPackets}
    received.EntityData.Leafs["start-control-connection-requests"] = types.YLeaf{"StartControlConnectionRequests", received.StartControlConnectionRequests}
    received.EntityData.Leafs["start-control-connection-replies"] = types.YLeaf{"StartControlConnectionReplies", received.StartControlConnectionReplies}
    received.EntityData.Leafs["start-control-connection-notifications"] = types.YLeaf{"StartControlConnectionNotifications", received.StartControlConnectionNotifications}
    received.EntityData.Leafs["stop-control-connection-notifications"] = types.YLeaf{"StopControlConnectionNotifications", received.StopControlConnectionNotifications}
    received.EntityData.Leafs["hello-packets"] = types.YLeaf{"HelloPackets", received.HelloPackets}
    received.EntityData.Leafs["outgoing-call-requests"] = types.YLeaf{"OutgoingCallRequests", received.OutgoingCallRequests}
    received.EntityData.Leafs["outgoing-call-replies"] = types.YLeaf{"OutgoingCallReplies", received.OutgoingCallReplies}
    received.EntityData.Leafs["outgoing-call-connected-packets"] = types.YLeaf{"OutgoingCallConnectedPackets", received.OutgoingCallConnectedPackets}
    received.EntityData.Leafs["incoming-call-requests"] = types.YLeaf{"IncomingCallRequests", received.IncomingCallRequests}
    received.EntityData.Leafs["incoming-call-replies"] = types.YLeaf{"IncomingCallReplies", received.IncomingCallReplies}
    received.EntityData.Leafs["incoming-call-connected-packets"] = types.YLeaf{"IncomingCallConnectedPackets", received.IncomingCallConnectedPackets}
    received.EntityData.Leafs["call-disconnect-notify-packets"] = types.YLeaf{"CallDisconnectNotifyPackets", received.CallDisconnectNotifyPackets}
    received.EntityData.Leafs["wan-error-notify-packets"] = types.YLeaf{"WanErrorNotifyPackets", received.WanErrorNotifyPackets}
    received.EntityData.Leafs["set-link-info-packets"] = types.YLeaf{"SetLinkInfoPackets", received.SetLinkInfoPackets}
    received.EntityData.Leafs["service-relay-requests"] = types.YLeaf{"ServiceRelayRequests", received.ServiceRelayRequests}
    received.EntityData.Leafs["service-relay-replies"] = types.YLeaf{"ServiceRelayReplies", received.ServiceRelayReplies}
    received.EntityData.Leafs["acknowledgement-packets"] = types.YLeaf{"AcknowledgementPackets", received.AcknowledgementPackets}
    return &(received.EntityData)
}

// L2Tp_Counters_Control_TunnelXr_Global_Drop
// Drop data
type L2Tp_Counters_Control_TunnelXr_Global_Drop struct {
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

func (drop *L2Tp_Counters_Control_TunnelXr_Global_Drop) GetEntityData() *types.CommonEntityData {
    drop.EntityData.YFilter = drop.YFilter
    drop.EntityData.YangName = "drop"
    drop.EntityData.BundleName = "cisco_ios_xr"
    drop.EntityData.ParentYangName = "global"
    drop.EntityData.SegmentPath = "drop"
    drop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    drop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    drop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    drop.EntityData.Children = make(map[string]types.YChild)
    drop.EntityData.Leafs = make(map[string]types.YLeaf)
    drop.EntityData.Leafs["unknown-packets"] = types.YLeaf{"UnknownPackets", drop.UnknownPackets}
    drop.EntityData.Leafs["zero-length-body-packets"] = types.YLeaf{"ZeroLengthBodyPackets", drop.ZeroLengthBodyPackets}
    drop.EntityData.Leafs["start-control-connection-requests"] = types.YLeaf{"StartControlConnectionRequests", drop.StartControlConnectionRequests}
    drop.EntityData.Leafs["start-control-connection-replies"] = types.YLeaf{"StartControlConnectionReplies", drop.StartControlConnectionReplies}
    drop.EntityData.Leafs["start-control-connection-notifications"] = types.YLeaf{"StartControlConnectionNotifications", drop.StartControlConnectionNotifications}
    drop.EntityData.Leafs["stop-control-connection-notifications"] = types.YLeaf{"StopControlConnectionNotifications", drop.StopControlConnectionNotifications}
    drop.EntityData.Leafs["hello-packets"] = types.YLeaf{"HelloPackets", drop.HelloPackets}
    drop.EntityData.Leafs["outgoing-call-requests"] = types.YLeaf{"OutgoingCallRequests", drop.OutgoingCallRequests}
    drop.EntityData.Leafs["outgoing-call-replies"] = types.YLeaf{"OutgoingCallReplies", drop.OutgoingCallReplies}
    drop.EntityData.Leafs["outgoing-call-connected-packets"] = types.YLeaf{"OutgoingCallConnectedPackets", drop.OutgoingCallConnectedPackets}
    drop.EntityData.Leafs["incoming-call-requests"] = types.YLeaf{"IncomingCallRequests", drop.IncomingCallRequests}
    drop.EntityData.Leafs["incoming-call-replies"] = types.YLeaf{"IncomingCallReplies", drop.IncomingCallReplies}
    drop.EntityData.Leafs["incoming-call-connected-packets"] = types.YLeaf{"IncomingCallConnectedPackets", drop.IncomingCallConnectedPackets}
    drop.EntityData.Leafs["call-disconnect-notify-packets"] = types.YLeaf{"CallDisconnectNotifyPackets", drop.CallDisconnectNotifyPackets}
    drop.EntityData.Leafs["wan-error-notify-packets"] = types.YLeaf{"WanErrorNotifyPackets", drop.WanErrorNotifyPackets}
    drop.EntityData.Leafs["set-link-info-packets"] = types.YLeaf{"SetLinkInfoPackets", drop.SetLinkInfoPackets}
    drop.EntityData.Leafs["service-relay-requests"] = types.YLeaf{"ServiceRelayRequests", drop.ServiceRelayRequests}
    drop.EntityData.Leafs["service-relay-replies"] = types.YLeaf{"ServiceRelayReplies", drop.ServiceRelayReplies}
    drop.EntityData.Leafs["acknowledgement-packets"] = types.YLeaf{"AcknowledgementPackets", drop.AcknowledgementPackets}
    return &(drop.EntityData)
}

// L2Tp_Counters_Control_Tunnels
// Table of tunnel IDs of control message counters
type L2Tp_Counters_Control_Tunnels struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2TP tunnel control message counters. The type is slice of
    // L2Tp_Counters_Control_Tunnels_Tunnel.
    Tunnel []L2Tp_Counters_Control_Tunnels_Tunnel
}

func (tunnels *L2Tp_Counters_Control_Tunnels) GetEntityData() *types.CommonEntityData {
    tunnels.EntityData.YFilter = tunnels.YFilter
    tunnels.EntityData.YangName = "tunnels"
    tunnels.EntityData.BundleName = "cisco_ios_xr"
    tunnels.EntityData.ParentYangName = "control"
    tunnels.EntityData.SegmentPath = "tunnels"
    tunnels.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnels.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnels.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnels.EntityData.Children = make(map[string]types.YChild)
    tunnels.EntityData.Children["tunnel"] = types.YChild{"Tunnel", nil}
    for i := range tunnels.Tunnel {
        tunnels.EntityData.Children[types.GetSegmentPath(&tunnels.Tunnel[i])] = types.YChild{"Tunnel", &tunnels.Tunnel[i]}
    }
    tunnels.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(tunnels.EntityData)
}

// L2Tp_Counters_Control_Tunnels_Tunnel
// L2TP tunnel control message counters
type L2Tp_Counters_Control_Tunnels_Tunnel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. L2TP tunnel ID. The type is interface{} with
    // range: -2147483648..2147483647.
    TunnelId interface{}

    // L2TP control message local and remote addresses.
    Brief L2Tp_Counters_Control_Tunnels_Tunnel_Brief

    // Global data.
    Global L2Tp_Counters_Control_Tunnels_Tunnel_Global
}

func (tunnel *L2Tp_Counters_Control_Tunnels_Tunnel) GetEntityData() *types.CommonEntityData {
    tunnel.EntityData.YFilter = tunnel.YFilter
    tunnel.EntityData.YangName = "tunnel"
    tunnel.EntityData.BundleName = "cisco_ios_xr"
    tunnel.EntityData.ParentYangName = "tunnels"
    tunnel.EntityData.SegmentPath = "tunnel" + "[tunnel-id='" + fmt.Sprintf("%v", tunnel.TunnelId) + "']"
    tunnel.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnel.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnel.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnel.EntityData.Children = make(map[string]types.YChild)
    tunnel.EntityData.Children["brief"] = types.YChild{"Brief", &tunnel.Brief}
    tunnel.EntityData.Children["global"] = types.YChild{"Global", &tunnel.Global}
    tunnel.EntityData.Leafs = make(map[string]types.YLeaf)
    tunnel.EntityData.Leafs["tunnel-id"] = types.YLeaf{"TunnelId", tunnel.TunnelId}
    return &(tunnel.EntityData)
}

// L2Tp_Counters_Control_Tunnels_Tunnel_Brief
// L2TP control message local and remote addresses
type L2Tp_Counters_Control_Tunnels_Tunnel_Brief struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Remote tunnel ID. The type is interface{} with range: 0..4294967295.
    RemoteTunnelId interface{}

    // Local IP address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    LocalAddress interface{}

    // Remote IP address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RemoteAddress interface{}
}

func (brief *L2Tp_Counters_Control_Tunnels_Tunnel_Brief) GetEntityData() *types.CommonEntityData {
    brief.EntityData.YFilter = brief.YFilter
    brief.EntityData.YangName = "brief"
    brief.EntityData.BundleName = "cisco_ios_xr"
    brief.EntityData.ParentYangName = "tunnel"
    brief.EntityData.SegmentPath = "brief"
    brief.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    brief.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    brief.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    brief.EntityData.Children = make(map[string]types.YChild)
    brief.EntityData.Leafs = make(map[string]types.YLeaf)
    brief.EntityData.Leafs["remote-tunnel-id"] = types.YLeaf{"RemoteTunnelId", brief.RemoteTunnelId}
    brief.EntityData.Leafs["local-address"] = types.YLeaf{"LocalAddress", brief.LocalAddress}
    brief.EntityData.Leafs["remote-address"] = types.YLeaf{"RemoteAddress", brief.RemoteAddress}
    return &(brief.EntityData)
}

// L2Tp_Counters_Control_Tunnels_Tunnel_Global
// Global data
type L2Tp_Counters_Control_Tunnels_Tunnel_Global struct {
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
    Transmit L2Tp_Counters_Control_Tunnels_Tunnel_Global_Transmit

    // Re transmit data.
    Retransmit L2Tp_Counters_Control_Tunnels_Tunnel_Global_Retransmit

    // Received data.
    Received L2Tp_Counters_Control_Tunnels_Tunnel_Global_Received

    // Drop data.
    Drop L2Tp_Counters_Control_Tunnels_Tunnel_Global_Drop
}

func (global *L2Tp_Counters_Control_Tunnels_Tunnel_Global) GetEntityData() *types.CommonEntityData {
    global.EntityData.YFilter = global.YFilter
    global.EntityData.YangName = "global"
    global.EntityData.BundleName = "cisco_ios_xr"
    global.EntityData.ParentYangName = "tunnel"
    global.EntityData.SegmentPath = "global"
    global.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    global.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    global.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    global.EntityData.Children = make(map[string]types.YChild)
    global.EntityData.Children["transmit"] = types.YChild{"Transmit", &global.Transmit}
    global.EntityData.Children["retransmit"] = types.YChild{"Retransmit", &global.Retransmit}
    global.EntityData.Children["received"] = types.YChild{"Received", &global.Received}
    global.EntityData.Children["drop"] = types.YChild{"Drop", &global.Drop}
    global.EntityData.Leafs = make(map[string]types.YLeaf)
    global.EntityData.Leafs["total-transmit"] = types.YLeaf{"TotalTransmit", global.TotalTransmit}
    global.EntityData.Leafs["total-retransmit"] = types.YLeaf{"TotalRetransmit", global.TotalRetransmit}
    global.EntityData.Leafs["total-received"] = types.YLeaf{"TotalReceived", global.TotalReceived}
    global.EntityData.Leafs["total-drop"] = types.YLeaf{"TotalDrop", global.TotalDrop}
    return &(global.EntityData)
}

// L2Tp_Counters_Control_Tunnels_Tunnel_Global_Transmit
// Transmit data
type L2Tp_Counters_Control_Tunnels_Tunnel_Global_Transmit struct {
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

func (transmit *L2Tp_Counters_Control_Tunnels_Tunnel_Global_Transmit) GetEntityData() *types.CommonEntityData {
    transmit.EntityData.YFilter = transmit.YFilter
    transmit.EntityData.YangName = "transmit"
    transmit.EntityData.BundleName = "cisco_ios_xr"
    transmit.EntityData.ParentYangName = "global"
    transmit.EntityData.SegmentPath = "transmit"
    transmit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transmit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transmit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transmit.EntityData.Children = make(map[string]types.YChild)
    transmit.EntityData.Leafs = make(map[string]types.YLeaf)
    transmit.EntityData.Leafs["unknown-packets"] = types.YLeaf{"UnknownPackets", transmit.UnknownPackets}
    transmit.EntityData.Leafs["zero-length-body-packets"] = types.YLeaf{"ZeroLengthBodyPackets", transmit.ZeroLengthBodyPackets}
    transmit.EntityData.Leafs["start-control-connection-requests"] = types.YLeaf{"StartControlConnectionRequests", transmit.StartControlConnectionRequests}
    transmit.EntityData.Leafs["start-control-connection-replies"] = types.YLeaf{"StartControlConnectionReplies", transmit.StartControlConnectionReplies}
    transmit.EntityData.Leafs["start-control-connection-notifications"] = types.YLeaf{"StartControlConnectionNotifications", transmit.StartControlConnectionNotifications}
    transmit.EntityData.Leafs["stop-control-connection-notifications"] = types.YLeaf{"StopControlConnectionNotifications", transmit.StopControlConnectionNotifications}
    transmit.EntityData.Leafs["hello-packets"] = types.YLeaf{"HelloPackets", transmit.HelloPackets}
    transmit.EntityData.Leafs["outgoing-call-requests"] = types.YLeaf{"OutgoingCallRequests", transmit.OutgoingCallRequests}
    transmit.EntityData.Leafs["outgoing-call-replies"] = types.YLeaf{"OutgoingCallReplies", transmit.OutgoingCallReplies}
    transmit.EntityData.Leafs["outgoing-call-connected-packets"] = types.YLeaf{"OutgoingCallConnectedPackets", transmit.OutgoingCallConnectedPackets}
    transmit.EntityData.Leafs["incoming-call-requests"] = types.YLeaf{"IncomingCallRequests", transmit.IncomingCallRequests}
    transmit.EntityData.Leafs["incoming-call-replies"] = types.YLeaf{"IncomingCallReplies", transmit.IncomingCallReplies}
    transmit.EntityData.Leafs["incoming-call-connected-packets"] = types.YLeaf{"IncomingCallConnectedPackets", transmit.IncomingCallConnectedPackets}
    transmit.EntityData.Leafs["call-disconnect-notify-packets"] = types.YLeaf{"CallDisconnectNotifyPackets", transmit.CallDisconnectNotifyPackets}
    transmit.EntityData.Leafs["wan-error-notify-packets"] = types.YLeaf{"WanErrorNotifyPackets", transmit.WanErrorNotifyPackets}
    transmit.EntityData.Leafs["set-link-info-packets"] = types.YLeaf{"SetLinkInfoPackets", transmit.SetLinkInfoPackets}
    transmit.EntityData.Leafs["service-relay-requests"] = types.YLeaf{"ServiceRelayRequests", transmit.ServiceRelayRequests}
    transmit.EntityData.Leafs["service-relay-replies"] = types.YLeaf{"ServiceRelayReplies", transmit.ServiceRelayReplies}
    transmit.EntityData.Leafs["acknowledgement-packets"] = types.YLeaf{"AcknowledgementPackets", transmit.AcknowledgementPackets}
    return &(transmit.EntityData)
}

// L2Tp_Counters_Control_Tunnels_Tunnel_Global_Retransmit
// Re transmit data
type L2Tp_Counters_Control_Tunnels_Tunnel_Global_Retransmit struct {
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

func (retransmit *L2Tp_Counters_Control_Tunnels_Tunnel_Global_Retransmit) GetEntityData() *types.CommonEntityData {
    retransmit.EntityData.YFilter = retransmit.YFilter
    retransmit.EntityData.YangName = "retransmit"
    retransmit.EntityData.BundleName = "cisco_ios_xr"
    retransmit.EntityData.ParentYangName = "global"
    retransmit.EntityData.SegmentPath = "retransmit"
    retransmit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    retransmit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    retransmit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    retransmit.EntityData.Children = make(map[string]types.YChild)
    retransmit.EntityData.Leafs = make(map[string]types.YLeaf)
    retransmit.EntityData.Leafs["unknown-packets"] = types.YLeaf{"UnknownPackets", retransmit.UnknownPackets}
    retransmit.EntityData.Leafs["zero-length-body-packets"] = types.YLeaf{"ZeroLengthBodyPackets", retransmit.ZeroLengthBodyPackets}
    retransmit.EntityData.Leafs["start-control-connection-requests"] = types.YLeaf{"StartControlConnectionRequests", retransmit.StartControlConnectionRequests}
    retransmit.EntityData.Leafs["start-control-connection-replies"] = types.YLeaf{"StartControlConnectionReplies", retransmit.StartControlConnectionReplies}
    retransmit.EntityData.Leafs["start-control-connection-notifications"] = types.YLeaf{"StartControlConnectionNotifications", retransmit.StartControlConnectionNotifications}
    retransmit.EntityData.Leafs["stop-control-connection-notifications"] = types.YLeaf{"StopControlConnectionNotifications", retransmit.StopControlConnectionNotifications}
    retransmit.EntityData.Leafs["hello-packets"] = types.YLeaf{"HelloPackets", retransmit.HelloPackets}
    retransmit.EntityData.Leafs["outgoing-call-requests"] = types.YLeaf{"OutgoingCallRequests", retransmit.OutgoingCallRequests}
    retransmit.EntityData.Leafs["outgoing-call-replies"] = types.YLeaf{"OutgoingCallReplies", retransmit.OutgoingCallReplies}
    retransmit.EntityData.Leafs["outgoing-call-connected-packets"] = types.YLeaf{"OutgoingCallConnectedPackets", retransmit.OutgoingCallConnectedPackets}
    retransmit.EntityData.Leafs["incoming-call-requests"] = types.YLeaf{"IncomingCallRequests", retransmit.IncomingCallRequests}
    retransmit.EntityData.Leafs["incoming-call-replies"] = types.YLeaf{"IncomingCallReplies", retransmit.IncomingCallReplies}
    retransmit.EntityData.Leafs["incoming-call-connected-packets"] = types.YLeaf{"IncomingCallConnectedPackets", retransmit.IncomingCallConnectedPackets}
    retransmit.EntityData.Leafs["call-disconnect-notify-packets"] = types.YLeaf{"CallDisconnectNotifyPackets", retransmit.CallDisconnectNotifyPackets}
    retransmit.EntityData.Leafs["wan-error-notify-packets"] = types.YLeaf{"WanErrorNotifyPackets", retransmit.WanErrorNotifyPackets}
    retransmit.EntityData.Leafs["set-link-info-packets"] = types.YLeaf{"SetLinkInfoPackets", retransmit.SetLinkInfoPackets}
    retransmit.EntityData.Leafs["service-relay-requests"] = types.YLeaf{"ServiceRelayRequests", retransmit.ServiceRelayRequests}
    retransmit.EntityData.Leafs["service-relay-replies"] = types.YLeaf{"ServiceRelayReplies", retransmit.ServiceRelayReplies}
    retransmit.EntityData.Leafs["acknowledgement-packets"] = types.YLeaf{"AcknowledgementPackets", retransmit.AcknowledgementPackets}
    return &(retransmit.EntityData)
}

// L2Tp_Counters_Control_Tunnels_Tunnel_Global_Received
// Received data
type L2Tp_Counters_Control_Tunnels_Tunnel_Global_Received struct {
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

func (received *L2Tp_Counters_Control_Tunnels_Tunnel_Global_Received) GetEntityData() *types.CommonEntityData {
    received.EntityData.YFilter = received.YFilter
    received.EntityData.YangName = "received"
    received.EntityData.BundleName = "cisco_ios_xr"
    received.EntityData.ParentYangName = "global"
    received.EntityData.SegmentPath = "received"
    received.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    received.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    received.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    received.EntityData.Children = make(map[string]types.YChild)
    received.EntityData.Leafs = make(map[string]types.YLeaf)
    received.EntityData.Leafs["unknown-packets"] = types.YLeaf{"UnknownPackets", received.UnknownPackets}
    received.EntityData.Leafs["zero-length-body-packets"] = types.YLeaf{"ZeroLengthBodyPackets", received.ZeroLengthBodyPackets}
    received.EntityData.Leafs["start-control-connection-requests"] = types.YLeaf{"StartControlConnectionRequests", received.StartControlConnectionRequests}
    received.EntityData.Leafs["start-control-connection-replies"] = types.YLeaf{"StartControlConnectionReplies", received.StartControlConnectionReplies}
    received.EntityData.Leafs["start-control-connection-notifications"] = types.YLeaf{"StartControlConnectionNotifications", received.StartControlConnectionNotifications}
    received.EntityData.Leafs["stop-control-connection-notifications"] = types.YLeaf{"StopControlConnectionNotifications", received.StopControlConnectionNotifications}
    received.EntityData.Leafs["hello-packets"] = types.YLeaf{"HelloPackets", received.HelloPackets}
    received.EntityData.Leafs["outgoing-call-requests"] = types.YLeaf{"OutgoingCallRequests", received.OutgoingCallRequests}
    received.EntityData.Leafs["outgoing-call-replies"] = types.YLeaf{"OutgoingCallReplies", received.OutgoingCallReplies}
    received.EntityData.Leafs["outgoing-call-connected-packets"] = types.YLeaf{"OutgoingCallConnectedPackets", received.OutgoingCallConnectedPackets}
    received.EntityData.Leafs["incoming-call-requests"] = types.YLeaf{"IncomingCallRequests", received.IncomingCallRequests}
    received.EntityData.Leafs["incoming-call-replies"] = types.YLeaf{"IncomingCallReplies", received.IncomingCallReplies}
    received.EntityData.Leafs["incoming-call-connected-packets"] = types.YLeaf{"IncomingCallConnectedPackets", received.IncomingCallConnectedPackets}
    received.EntityData.Leafs["call-disconnect-notify-packets"] = types.YLeaf{"CallDisconnectNotifyPackets", received.CallDisconnectNotifyPackets}
    received.EntityData.Leafs["wan-error-notify-packets"] = types.YLeaf{"WanErrorNotifyPackets", received.WanErrorNotifyPackets}
    received.EntityData.Leafs["set-link-info-packets"] = types.YLeaf{"SetLinkInfoPackets", received.SetLinkInfoPackets}
    received.EntityData.Leafs["service-relay-requests"] = types.YLeaf{"ServiceRelayRequests", received.ServiceRelayRequests}
    received.EntityData.Leafs["service-relay-replies"] = types.YLeaf{"ServiceRelayReplies", received.ServiceRelayReplies}
    received.EntityData.Leafs["acknowledgement-packets"] = types.YLeaf{"AcknowledgementPackets", received.AcknowledgementPackets}
    return &(received.EntityData)
}

// L2Tp_Counters_Control_Tunnels_Tunnel_Global_Drop
// Drop data
type L2Tp_Counters_Control_Tunnels_Tunnel_Global_Drop struct {
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

func (drop *L2Tp_Counters_Control_Tunnels_Tunnel_Global_Drop) GetEntityData() *types.CommonEntityData {
    drop.EntityData.YFilter = drop.YFilter
    drop.EntityData.YangName = "drop"
    drop.EntityData.BundleName = "cisco_ios_xr"
    drop.EntityData.ParentYangName = "global"
    drop.EntityData.SegmentPath = "drop"
    drop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    drop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    drop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    drop.EntityData.Children = make(map[string]types.YChild)
    drop.EntityData.Leafs = make(map[string]types.YLeaf)
    drop.EntityData.Leafs["unknown-packets"] = types.YLeaf{"UnknownPackets", drop.UnknownPackets}
    drop.EntityData.Leafs["zero-length-body-packets"] = types.YLeaf{"ZeroLengthBodyPackets", drop.ZeroLengthBodyPackets}
    drop.EntityData.Leafs["start-control-connection-requests"] = types.YLeaf{"StartControlConnectionRequests", drop.StartControlConnectionRequests}
    drop.EntityData.Leafs["start-control-connection-replies"] = types.YLeaf{"StartControlConnectionReplies", drop.StartControlConnectionReplies}
    drop.EntityData.Leafs["start-control-connection-notifications"] = types.YLeaf{"StartControlConnectionNotifications", drop.StartControlConnectionNotifications}
    drop.EntityData.Leafs["stop-control-connection-notifications"] = types.YLeaf{"StopControlConnectionNotifications", drop.StopControlConnectionNotifications}
    drop.EntityData.Leafs["hello-packets"] = types.YLeaf{"HelloPackets", drop.HelloPackets}
    drop.EntityData.Leafs["outgoing-call-requests"] = types.YLeaf{"OutgoingCallRequests", drop.OutgoingCallRequests}
    drop.EntityData.Leafs["outgoing-call-replies"] = types.YLeaf{"OutgoingCallReplies", drop.OutgoingCallReplies}
    drop.EntityData.Leafs["outgoing-call-connected-packets"] = types.YLeaf{"OutgoingCallConnectedPackets", drop.OutgoingCallConnectedPackets}
    drop.EntityData.Leafs["incoming-call-requests"] = types.YLeaf{"IncomingCallRequests", drop.IncomingCallRequests}
    drop.EntityData.Leafs["incoming-call-replies"] = types.YLeaf{"IncomingCallReplies", drop.IncomingCallReplies}
    drop.EntityData.Leafs["incoming-call-connected-packets"] = types.YLeaf{"IncomingCallConnectedPackets", drop.IncomingCallConnectedPackets}
    drop.EntityData.Leafs["call-disconnect-notify-packets"] = types.YLeaf{"CallDisconnectNotifyPackets", drop.CallDisconnectNotifyPackets}
    drop.EntityData.Leafs["wan-error-notify-packets"] = types.YLeaf{"WanErrorNotifyPackets", drop.WanErrorNotifyPackets}
    drop.EntityData.Leafs["set-link-info-packets"] = types.YLeaf{"SetLinkInfoPackets", drop.SetLinkInfoPackets}
    drop.EntityData.Leafs["service-relay-requests"] = types.YLeaf{"ServiceRelayRequests", drop.ServiceRelayRequests}
    drop.EntityData.Leafs["service-relay-replies"] = types.YLeaf{"ServiceRelayReplies", drop.ServiceRelayReplies}
    drop.EntityData.Leafs["acknowledgement-packets"] = types.YLeaf{"AcknowledgementPackets", drop.AcknowledgementPackets}
    return &(drop.EntityData)
}

// L2Tp_TunnelConfigurations
// List of tunnel IDs
type L2Tp_TunnelConfigurations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2TP tunnel information. The type is slice of
    // L2Tp_TunnelConfigurations_TunnelConfiguration.
    TunnelConfiguration []L2Tp_TunnelConfigurations_TunnelConfiguration
}

func (tunnelConfigurations *L2Tp_TunnelConfigurations) GetEntityData() *types.CommonEntityData {
    tunnelConfigurations.EntityData.YFilter = tunnelConfigurations.YFilter
    tunnelConfigurations.EntityData.YangName = "tunnel-configurations"
    tunnelConfigurations.EntityData.BundleName = "cisco_ios_xr"
    tunnelConfigurations.EntityData.ParentYangName = "l2tp"
    tunnelConfigurations.EntityData.SegmentPath = "tunnel-configurations"
    tunnelConfigurations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnelConfigurations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnelConfigurations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnelConfigurations.EntityData.Children = make(map[string]types.YChild)
    tunnelConfigurations.EntityData.Children["tunnel-configuration"] = types.YChild{"TunnelConfiguration", nil}
    for i := range tunnelConfigurations.TunnelConfiguration {
        tunnelConfigurations.EntityData.Children[types.GetSegmentPath(&tunnelConfigurations.TunnelConfiguration[i])] = types.YChild{"TunnelConfiguration", &tunnelConfigurations.TunnelConfiguration[i]}
    }
    tunnelConfigurations.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(tunnelConfigurations.EntityData)
}

// L2Tp_TunnelConfigurations_TunnelConfiguration
// L2TP tunnel information
type L2Tp_TunnelConfigurations_TunnelConfiguration struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Local tunnel ID. The type is interface{} with
    // range: -2147483648..2147483647.
    LocalTunnelId interface{}

    // Remote tunnel ID. The type is interface{} with range: 0..4294967295.
    RemoteTunnelId interface{}

    // L2Tp class data.
    L2TpClass L2Tp_TunnelConfigurations_TunnelConfiguration_L2TpClass
}

func (tunnelConfiguration *L2Tp_TunnelConfigurations_TunnelConfiguration) GetEntityData() *types.CommonEntityData {
    tunnelConfiguration.EntityData.YFilter = tunnelConfiguration.YFilter
    tunnelConfiguration.EntityData.YangName = "tunnel-configuration"
    tunnelConfiguration.EntityData.BundleName = "cisco_ios_xr"
    tunnelConfiguration.EntityData.ParentYangName = "tunnel-configurations"
    tunnelConfiguration.EntityData.SegmentPath = "tunnel-configuration" + "[local-tunnel-id='" + fmt.Sprintf("%v", tunnelConfiguration.LocalTunnelId) + "']"
    tunnelConfiguration.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnelConfiguration.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnelConfiguration.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnelConfiguration.EntityData.Children = make(map[string]types.YChild)
    tunnelConfiguration.EntityData.Children["l2tp-class"] = types.YChild{"L2TpClass", &tunnelConfiguration.L2TpClass}
    tunnelConfiguration.EntityData.Leafs = make(map[string]types.YLeaf)
    tunnelConfiguration.EntityData.Leafs["local-tunnel-id"] = types.YLeaf{"LocalTunnelId", tunnelConfiguration.LocalTunnelId}
    tunnelConfiguration.EntityData.Leafs["remote-tunnel-id"] = types.YLeaf{"RemoteTunnelId", tunnelConfiguration.RemoteTunnelId}
    return &(tunnelConfiguration.EntityData)
}

// L2Tp_TunnelConfigurations_TunnelConfiguration_L2TpClass
// L2Tp class data
type L2Tp_TunnelConfigurations_TunnelConfiguration_L2TpClass struct {
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

func (l2TpClass *L2Tp_TunnelConfigurations_TunnelConfiguration_L2TpClass) GetEntityData() *types.CommonEntityData {
    l2TpClass.EntityData.YFilter = l2TpClass.YFilter
    l2TpClass.EntityData.YangName = "l2tp-class"
    l2TpClass.EntityData.BundleName = "cisco_ios_xr"
    l2TpClass.EntityData.ParentYangName = "tunnel-configuration"
    l2TpClass.EntityData.SegmentPath = "l2tp-class"
    l2TpClass.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    l2TpClass.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    l2TpClass.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    l2TpClass.EntityData.Children = make(map[string]types.YChild)
    l2TpClass.EntityData.Leafs = make(map[string]types.YLeaf)
    l2TpClass.EntityData.Leafs["ip-tos"] = types.YLeaf{"IpTos", l2TpClass.IpTos}
    l2TpClass.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", l2TpClass.VrfName}
    l2TpClass.EntityData.Leafs["receive-window-size"] = types.YLeaf{"ReceiveWindowSize", l2TpClass.ReceiveWindowSize}
    l2TpClass.EntityData.Leafs["class-name-xr"] = types.YLeaf{"ClassNameXr", l2TpClass.ClassNameXr}
    l2TpClass.EntityData.Leafs["digest-hash"] = types.YLeaf{"DigestHash", l2TpClass.DigestHash}
    l2TpClass.EntityData.Leafs["password"] = types.YLeaf{"Password", l2TpClass.Password}
    l2TpClass.EntityData.Leafs["encoded-password"] = types.YLeaf{"EncodedPassword", l2TpClass.EncodedPassword}
    l2TpClass.EntityData.Leafs["host-name"] = types.YLeaf{"HostName", l2TpClass.HostName}
    l2TpClass.EntityData.Leafs["accounting-method-list"] = types.YLeaf{"AccountingMethodList", l2TpClass.AccountingMethodList}
    l2TpClass.EntityData.Leafs["hello-timeout"] = types.YLeaf{"HelloTimeout", l2TpClass.HelloTimeout}
    l2TpClass.EntityData.Leafs["setup-timeout"] = types.YLeaf{"SetupTimeout", l2TpClass.SetupTimeout}
    l2TpClass.EntityData.Leafs["retransmit-minimum-timeout"] = types.YLeaf{"RetransmitMinimumTimeout", l2TpClass.RetransmitMinimumTimeout}
    l2TpClass.EntityData.Leafs["retransmit-maximum-timeout"] = types.YLeaf{"RetransmitMaximumTimeout", l2TpClass.RetransmitMaximumTimeout}
    l2TpClass.EntityData.Leafs["initial-retransmit-minimum-timeout"] = types.YLeaf{"InitialRetransmitMinimumTimeout", l2TpClass.InitialRetransmitMinimumTimeout}
    l2TpClass.EntityData.Leafs["initial-retransmit-maximum-timeout"] = types.YLeaf{"InitialRetransmitMaximumTimeout", l2TpClass.InitialRetransmitMaximumTimeout}
    l2TpClass.EntityData.Leafs["timeout-no-user"] = types.YLeaf{"TimeoutNoUser", l2TpClass.TimeoutNoUser}
    l2TpClass.EntityData.Leafs["retransmit-retries"] = types.YLeaf{"RetransmitRetries", l2TpClass.RetransmitRetries}
    l2TpClass.EntityData.Leafs["initial-retransmit-retries"] = types.YLeaf{"InitialRetransmitRetries", l2TpClass.InitialRetransmitRetries}
    l2TpClass.EntityData.Leafs["is-authentication-enabled"] = types.YLeaf{"IsAuthenticationEnabled", l2TpClass.IsAuthenticationEnabled}
    l2TpClass.EntityData.Leafs["is-hidden"] = types.YLeaf{"IsHidden", l2TpClass.IsHidden}
    l2TpClass.EntityData.Leafs["is-digest-enabled"] = types.YLeaf{"IsDigestEnabled", l2TpClass.IsDigestEnabled}
    l2TpClass.EntityData.Leafs["is-digest-check-enabled"] = types.YLeaf{"IsDigestCheckEnabled", l2TpClass.IsDigestCheckEnabled}
    l2TpClass.EntityData.Leafs["is-congestion-control-enabled"] = types.YLeaf{"IsCongestionControlEnabled", l2TpClass.IsCongestionControlEnabled}
    l2TpClass.EntityData.Leafs["is-peer-address-checked"] = types.YLeaf{"IsPeerAddressChecked", l2TpClass.IsPeerAddressChecked}
    return &(l2TpClass.EntityData)
}

// L2Tp_CounterHistFail
// Failure events leading to disconnection
type L2Tp_CounterHistFail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // sesions affected due to timeout. The type is interface{} with range:
    // 0..4294967295.
    SessDownTmout interface{}

    // Send side counters. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    TxCounters interface{}

    // Receive side counters. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    RxCounters interface{}

    // timeout events by packet. The type is slice of interface{} with range:
    // 0..4294967295.
    PktTimeout []interface{}
}

func (counterHistFail *L2Tp_CounterHistFail) GetEntityData() *types.CommonEntityData {
    counterHistFail.EntityData.YFilter = counterHistFail.YFilter
    counterHistFail.EntityData.YangName = "counter-hist-fail"
    counterHistFail.EntityData.BundleName = "cisco_ios_xr"
    counterHistFail.EntityData.ParentYangName = "l2tp"
    counterHistFail.EntityData.SegmentPath = "counter-hist-fail"
    counterHistFail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    counterHistFail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    counterHistFail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    counterHistFail.EntityData.Children = make(map[string]types.YChild)
    counterHistFail.EntityData.Leafs = make(map[string]types.YLeaf)
    counterHistFail.EntityData.Leafs["sess-down-tmout"] = types.YLeaf{"SessDownTmout", counterHistFail.SessDownTmout}
    counterHistFail.EntityData.Leafs["tx-counters"] = types.YLeaf{"TxCounters", counterHistFail.TxCounters}
    counterHistFail.EntityData.Leafs["rx-counters"] = types.YLeaf{"RxCounters", counterHistFail.RxCounters}
    counterHistFail.EntityData.Leafs["pkt-timeout"] = types.YLeaf{"PktTimeout", counterHistFail.PktTimeout}
    return &(counterHistFail.EntityData)
}

// L2Tp_Classes
// List of L2TP class names
type L2Tp_Classes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2TP class name. The type is slice of L2Tp_Classes_Class.
    Class []L2Tp_Classes_Class
}

func (classes *L2Tp_Classes) GetEntityData() *types.CommonEntityData {
    classes.EntityData.YFilter = classes.YFilter
    classes.EntityData.YangName = "classes"
    classes.EntityData.BundleName = "cisco_ios_xr"
    classes.EntityData.ParentYangName = "l2tp"
    classes.EntityData.SegmentPath = "classes"
    classes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    classes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    classes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    classes.EntityData.Children = make(map[string]types.YChild)
    classes.EntityData.Children["class"] = types.YChild{"Class", nil}
    for i := range classes.Class {
        classes.EntityData.Children[types.GetSegmentPath(&classes.Class[i])] = types.YChild{"Class", &classes.Class[i]}
    }
    classes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(classes.EntityData)
}

// L2Tp_Classes_Class
// L2TP class name
type L2Tp_Classes_Class struct {
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

func (class *L2Tp_Classes_Class) GetEntityData() *types.CommonEntityData {
    class.EntityData.YFilter = class.YFilter
    class.EntityData.YangName = "class"
    class.EntityData.BundleName = "cisco_ios_xr"
    class.EntityData.ParentYangName = "classes"
    class.EntityData.SegmentPath = "class" + "[class-name='" + fmt.Sprintf("%v", class.ClassName) + "']"
    class.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    class.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    class.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    class.EntityData.Children = make(map[string]types.YChild)
    class.EntityData.Leafs = make(map[string]types.YLeaf)
    class.EntityData.Leafs["class-name"] = types.YLeaf{"ClassName", class.ClassName}
    class.EntityData.Leafs["ip-tos"] = types.YLeaf{"IpTos", class.IpTos}
    class.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", class.VrfName}
    class.EntityData.Leafs["receive-window-size"] = types.YLeaf{"ReceiveWindowSize", class.ReceiveWindowSize}
    class.EntityData.Leafs["class-name-xr"] = types.YLeaf{"ClassNameXr", class.ClassNameXr}
    class.EntityData.Leafs["digest-hash"] = types.YLeaf{"DigestHash", class.DigestHash}
    class.EntityData.Leafs["password"] = types.YLeaf{"Password", class.Password}
    class.EntityData.Leafs["encoded-password"] = types.YLeaf{"EncodedPassword", class.EncodedPassword}
    class.EntityData.Leafs["host-name"] = types.YLeaf{"HostName", class.HostName}
    class.EntityData.Leafs["accounting-method-list"] = types.YLeaf{"AccountingMethodList", class.AccountingMethodList}
    class.EntityData.Leafs["hello-timeout"] = types.YLeaf{"HelloTimeout", class.HelloTimeout}
    class.EntityData.Leafs["setup-timeout"] = types.YLeaf{"SetupTimeout", class.SetupTimeout}
    class.EntityData.Leafs["retransmit-minimum-timeout"] = types.YLeaf{"RetransmitMinimumTimeout", class.RetransmitMinimumTimeout}
    class.EntityData.Leafs["retransmit-maximum-timeout"] = types.YLeaf{"RetransmitMaximumTimeout", class.RetransmitMaximumTimeout}
    class.EntityData.Leafs["initial-retransmit-minimum-timeout"] = types.YLeaf{"InitialRetransmitMinimumTimeout", class.InitialRetransmitMinimumTimeout}
    class.EntityData.Leafs["initial-retransmit-maximum-timeout"] = types.YLeaf{"InitialRetransmitMaximumTimeout", class.InitialRetransmitMaximumTimeout}
    class.EntityData.Leafs["timeout-no-user"] = types.YLeaf{"TimeoutNoUser", class.TimeoutNoUser}
    class.EntityData.Leafs["retransmit-retries"] = types.YLeaf{"RetransmitRetries", class.RetransmitRetries}
    class.EntityData.Leafs["initial-retransmit-retries"] = types.YLeaf{"InitialRetransmitRetries", class.InitialRetransmitRetries}
    class.EntityData.Leafs["is-authentication-enabled"] = types.YLeaf{"IsAuthenticationEnabled", class.IsAuthenticationEnabled}
    class.EntityData.Leafs["is-hidden"] = types.YLeaf{"IsHidden", class.IsHidden}
    class.EntityData.Leafs["is-digest-enabled"] = types.YLeaf{"IsDigestEnabled", class.IsDigestEnabled}
    class.EntityData.Leafs["is-digest-check-enabled"] = types.YLeaf{"IsDigestCheckEnabled", class.IsDigestCheckEnabled}
    class.EntityData.Leafs["is-congestion-control-enabled"] = types.YLeaf{"IsCongestionControlEnabled", class.IsCongestionControlEnabled}
    class.EntityData.Leafs["is-peer-address-checked"] = types.YLeaf{"IsPeerAddressChecked", class.IsPeerAddressChecked}
    return &(class.EntityData)
}

// L2Tp_Tunnels
// List of tunnel IDs
type L2Tp_Tunnels struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2TP tunnel  information. The type is slice of L2Tp_Tunnels_Tunnel.
    Tunnel []L2Tp_Tunnels_Tunnel
}

func (tunnels *L2Tp_Tunnels) GetEntityData() *types.CommonEntityData {
    tunnels.EntityData.YFilter = tunnels.YFilter
    tunnels.EntityData.YangName = "tunnels"
    tunnels.EntityData.BundleName = "cisco_ios_xr"
    tunnels.EntityData.ParentYangName = "l2tp"
    tunnels.EntityData.SegmentPath = "tunnels"
    tunnels.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnels.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnels.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnels.EntityData.Children = make(map[string]types.YChild)
    tunnels.EntityData.Children["tunnel"] = types.YChild{"Tunnel", nil}
    for i := range tunnels.Tunnel {
        tunnels.EntityData.Children[types.GetSegmentPath(&tunnels.Tunnel[i])] = types.YChild{"Tunnel", &tunnels.Tunnel[i]}
    }
    tunnels.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(tunnels.EntityData)
}

// L2Tp_Tunnels_Tunnel
// L2TP tunnel  information
type L2Tp_Tunnels_Tunnel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Local tunnel ID. The type is interface{} with
    // range: -2147483648..2147483647.
    LocalTunnelId interface{}

    // Local tunnel address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    LocalAddress interface{}

    // Remote tunnel address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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

    // Retransmit time distribution in seconds. The type is slice of interface{}
    // with range: 0..65535. Units are second.
    RetransmitTime []interface{}
}

func (tunnel *L2Tp_Tunnels_Tunnel) GetEntityData() *types.CommonEntityData {
    tunnel.EntityData.YFilter = tunnel.YFilter
    tunnel.EntityData.YangName = "tunnel"
    tunnel.EntityData.BundleName = "cisco_ios_xr"
    tunnel.EntityData.ParentYangName = "tunnels"
    tunnel.EntityData.SegmentPath = "tunnel" + "[local-tunnel-id='" + fmt.Sprintf("%v", tunnel.LocalTunnelId) + "']"
    tunnel.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnel.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnel.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnel.EntityData.Children = make(map[string]types.YChild)
    tunnel.EntityData.Leafs = make(map[string]types.YLeaf)
    tunnel.EntityData.Leafs["local-tunnel-id"] = types.YLeaf{"LocalTunnelId", tunnel.LocalTunnelId}
    tunnel.EntityData.Leafs["local-address"] = types.YLeaf{"LocalAddress", tunnel.LocalAddress}
    tunnel.EntityData.Leafs["remote-address"] = types.YLeaf{"RemoteAddress", tunnel.RemoteAddress}
    tunnel.EntityData.Leafs["local-port"] = types.YLeaf{"LocalPort", tunnel.LocalPort}
    tunnel.EntityData.Leafs["remote-port"] = types.YLeaf{"RemotePort", tunnel.RemotePort}
    tunnel.EntityData.Leafs["protocol"] = types.YLeaf{"Protocol", tunnel.Protocol}
    tunnel.EntityData.Leafs["is-pmtu-enabled"] = types.YLeaf{"IsPmtuEnabled", tunnel.IsPmtuEnabled}
    tunnel.EntityData.Leafs["remote-tunnel-id"] = types.YLeaf{"RemoteTunnelId", tunnel.RemoteTunnelId}
    tunnel.EntityData.Leafs["local-tunnel-name"] = types.YLeaf{"LocalTunnelName", tunnel.LocalTunnelName}
    tunnel.EntityData.Leafs["remote-tunnel-name"] = types.YLeaf{"RemoteTunnelName", tunnel.RemoteTunnelName}
    tunnel.EntityData.Leafs["class-name"] = types.YLeaf{"ClassName", tunnel.ClassName}
    tunnel.EntityData.Leafs["active-sessions"] = types.YLeaf{"ActiveSessions", tunnel.ActiveSessions}
    tunnel.EntityData.Leafs["sequence-ns"] = types.YLeaf{"SequenceNs", tunnel.SequenceNs}
    tunnel.EntityData.Leafs["sequence-nr"] = types.YLeaf{"SequenceNr", tunnel.SequenceNr}
    tunnel.EntityData.Leafs["local-window-size"] = types.YLeaf{"LocalWindowSize", tunnel.LocalWindowSize}
    tunnel.EntityData.Leafs["remote-window-size"] = types.YLeaf{"RemoteWindowSize", tunnel.RemoteWindowSize}
    tunnel.EntityData.Leafs["retransmission-time"] = types.YLeaf{"RetransmissionTime", tunnel.RetransmissionTime}
    tunnel.EntityData.Leafs["maximum-retransmission-time"] = types.YLeaf{"MaximumRetransmissionTime", tunnel.MaximumRetransmissionTime}
    tunnel.EntityData.Leafs["unsent-queue-size"] = types.YLeaf{"UnsentQueueSize", tunnel.UnsentQueueSize}
    tunnel.EntityData.Leafs["unsent-maximum-queue-size"] = types.YLeaf{"UnsentMaximumQueueSize", tunnel.UnsentMaximumQueueSize}
    tunnel.EntityData.Leafs["resend-queue-size"] = types.YLeaf{"ResendQueueSize", tunnel.ResendQueueSize}
    tunnel.EntityData.Leafs["resend-maximum-queue-size"] = types.YLeaf{"ResendMaximumQueueSize", tunnel.ResendMaximumQueueSize}
    tunnel.EntityData.Leafs["order-queue-size"] = types.YLeaf{"OrderQueueSize", tunnel.OrderQueueSize}
    tunnel.EntityData.Leafs["packet-queue-check"] = types.YLeaf{"PacketQueueCheck", tunnel.PacketQueueCheck}
    tunnel.EntityData.Leafs["digest-secrets"] = types.YLeaf{"DigestSecrets", tunnel.DigestSecrets}
    tunnel.EntityData.Leafs["resends"] = types.YLeaf{"Resends", tunnel.Resends}
    tunnel.EntityData.Leafs["zero-length-body-acknowledgement-sent"] = types.YLeaf{"ZeroLengthBodyAcknowledgementSent", tunnel.ZeroLengthBodyAcknowledgementSent}
    tunnel.EntityData.Leafs["total-out-of-order-drop-packets"] = types.YLeaf{"TotalOutOfOrderDropPackets", tunnel.TotalOutOfOrderDropPackets}
    tunnel.EntityData.Leafs["total-out-of-order-reorder-packets"] = types.YLeaf{"TotalOutOfOrderReorderPackets", tunnel.TotalOutOfOrderReorderPackets}
    tunnel.EntityData.Leafs["total-peer-authentication-failures"] = types.YLeaf{"TotalPeerAuthenticationFailures", tunnel.TotalPeerAuthenticationFailures}
    tunnel.EntityData.Leafs["is-tunnel-up"] = types.YLeaf{"IsTunnelUp", tunnel.IsTunnelUp}
    tunnel.EntityData.Leafs["is-congestion-control-enabled"] = types.YLeaf{"IsCongestionControlEnabled", tunnel.IsCongestionControlEnabled}
    tunnel.EntityData.Leafs["retransmit-time"] = types.YLeaf{"RetransmitTime", tunnel.RetransmitTime}
    return &(tunnel.EntityData)
}

// L2Tp_Sessions
// List of session IDs
type L2Tp_Sessions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2TP information for a particular session. The type is slice of
    // L2Tp_Sessions_Session.
    Session []L2Tp_Sessions_Session
}

func (sessions *L2Tp_Sessions) GetEntityData() *types.CommonEntityData {
    sessions.EntityData.YFilter = sessions.YFilter
    sessions.EntityData.YangName = "sessions"
    sessions.EntityData.BundleName = "cisco_ios_xr"
    sessions.EntityData.ParentYangName = "l2tp"
    sessions.EntityData.SegmentPath = "sessions"
    sessions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessions.EntityData.Children = make(map[string]types.YChild)
    sessions.EntityData.Children["session"] = types.YChild{"Session", nil}
    for i := range sessions.Session {
        sessions.EntityData.Children[types.GetSegmentPath(&sessions.Session[i])] = types.YChild{"Session", &sessions.Session[i]}
    }
    sessions.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sessions.EntityData)
}

// L2Tp_Sessions_Session
// L2TP information for a particular session
type L2Tp_Sessions_Session struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Local tunnel ID. The type is interface{} with
    // range: -2147483648..2147483647.
    LocalTunnelId interface{}

    // This attribute is a key. Local session ID. The type is interface{} with
    // range: -2147483648..2147483647.
    LocalSessionId interface{}

    // Local session IP address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    LocalIpAddress interface{}

    // Remote session IP address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RemoteIpAddress interface{}

    // l2tp sh sess udp lport. The type is interface{} with range: 0..65535.
    L2TpShSessUdpLport interface{}

    // l2tp sh sess udp rport. The type is interface{} with range: 0..65535.
    L2TpShSessUdpRport interface{}

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
    L2TpShSessTieBreakerEnabled interface{}

    // l2tp sh sess tie breaker. The type is interface{} with range:
    // 0..18446744073709551615.
    L2TpShSessTieBreaker interface{}

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
    SessionApplicationData L2Tp_Sessions_Session_SessionApplicationData
}

func (session *L2Tp_Sessions_Session) GetEntityData() *types.CommonEntityData {
    session.EntityData.YFilter = session.YFilter
    session.EntityData.YangName = "session"
    session.EntityData.BundleName = "cisco_ios_xr"
    session.EntityData.ParentYangName = "sessions"
    session.EntityData.SegmentPath = "session" + "[local-tunnel-id='" + fmt.Sprintf("%v", session.LocalTunnelId) + "']" + "[local-session-id='" + fmt.Sprintf("%v", session.LocalSessionId) + "']"
    session.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    session.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    session.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    session.EntityData.Children = make(map[string]types.YChild)
    session.EntityData.Children["session-application-data"] = types.YChild{"SessionApplicationData", &session.SessionApplicationData}
    session.EntityData.Leafs = make(map[string]types.YLeaf)
    session.EntityData.Leafs["local-tunnel-id"] = types.YLeaf{"LocalTunnelId", session.LocalTunnelId}
    session.EntityData.Leafs["local-session-id"] = types.YLeaf{"LocalSessionId", session.LocalSessionId}
    session.EntityData.Leafs["local-ip-address"] = types.YLeaf{"LocalIpAddress", session.LocalIpAddress}
    session.EntityData.Leafs["remote-ip-address"] = types.YLeaf{"RemoteIpAddress", session.RemoteIpAddress}
    session.EntityData.Leafs["l2tp-sh-sess-udp-lport"] = types.YLeaf{"L2TpShSessUdpLport", session.L2TpShSessUdpLport}
    session.EntityData.Leafs["l2tp-sh-sess-udp-rport"] = types.YLeaf{"L2TpShSessUdpRport", session.L2TpShSessUdpRport}
    session.EntityData.Leafs["protocol"] = types.YLeaf{"Protocol", session.Protocol}
    session.EntityData.Leafs["remote-tunnel-id"] = types.YLeaf{"RemoteTunnelId", session.RemoteTunnelId}
    session.EntityData.Leafs["call-serial-number"] = types.YLeaf{"CallSerialNumber", session.CallSerialNumber}
    session.EntityData.Leafs["local-tunnel-name"] = types.YLeaf{"LocalTunnelName", session.LocalTunnelName}
    session.EntityData.Leafs["remote-tunnel-name"] = types.YLeaf{"RemoteTunnelName", session.RemoteTunnelName}
    session.EntityData.Leafs["remote-session-id"] = types.YLeaf{"RemoteSessionId", session.RemoteSessionId}
    session.EntityData.Leafs["l2tp-sh-sess-tie-breaker-enabled"] = types.YLeaf{"L2TpShSessTieBreakerEnabled", session.L2TpShSessTieBreakerEnabled}
    session.EntityData.Leafs["l2tp-sh-sess-tie-breaker"] = types.YLeaf{"L2TpShSessTieBreaker", session.L2TpShSessTieBreaker}
    session.EntityData.Leafs["is-session-manual"] = types.YLeaf{"IsSessionManual", session.IsSessionManual}
    session.EntityData.Leafs["is-session-up"] = types.YLeaf{"IsSessionUp", session.IsSessionUp}
    session.EntityData.Leafs["is-udp-checksum-enabled"] = types.YLeaf{"IsUdpChecksumEnabled", session.IsUdpChecksumEnabled}
    session.EntityData.Leafs["is-sequencing-on"] = types.YLeaf{"IsSequencingOn", session.IsSequencingOn}
    session.EntityData.Leafs["is-session-state-established"] = types.YLeaf{"IsSessionStateEstablished", session.IsSessionStateEstablished}
    session.EntityData.Leafs["is-session-locally-initiated"] = types.YLeaf{"IsSessionLocallyInitiated", session.IsSessionLocallyInitiated}
    session.EntityData.Leafs["is-conditional-debug-enabled"] = types.YLeaf{"IsConditionalDebugEnabled", session.IsConditionalDebugEnabled}
    session.EntityData.Leafs["unique-id"] = types.YLeaf{"UniqueId", session.UniqueId}
    session.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", session.InterfaceName}
    return &(session.EntityData)
}

// L2Tp_Sessions_Session_SessionApplicationData
// Session application data
type L2Tp_Sessions_Session_SessionApplicationData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // l2tp sh sess app type. The type is interface{} with range: 0..4294967295.
    L2TpShSessAppType interface{}

    // Xconnect data.
    Xconnect L2Tp_Sessions_Session_SessionApplicationData_Xconnect

    // VPDN data.
    Vpdn L2Tp_Sessions_Session_SessionApplicationData_Vpdn
}

func (sessionApplicationData *L2Tp_Sessions_Session_SessionApplicationData) GetEntityData() *types.CommonEntityData {
    sessionApplicationData.EntityData.YFilter = sessionApplicationData.YFilter
    sessionApplicationData.EntityData.YangName = "session-application-data"
    sessionApplicationData.EntityData.BundleName = "cisco_ios_xr"
    sessionApplicationData.EntityData.ParentYangName = "session"
    sessionApplicationData.EntityData.SegmentPath = "session-application-data"
    sessionApplicationData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionApplicationData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionApplicationData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionApplicationData.EntityData.Children = make(map[string]types.YChild)
    sessionApplicationData.EntityData.Children["xconnect"] = types.YChild{"Xconnect", &sessionApplicationData.Xconnect}
    sessionApplicationData.EntityData.Children["vpdn"] = types.YChild{"Vpdn", &sessionApplicationData.Vpdn}
    sessionApplicationData.EntityData.Leafs = make(map[string]types.YLeaf)
    sessionApplicationData.EntityData.Leafs["l2tp-sh-sess-app-type"] = types.YLeaf{"L2TpShSessAppType", sessionApplicationData.L2TpShSessAppType}
    return &(sessionApplicationData.EntityData)
}

// L2Tp_Sessions_Session_SessionApplicationData_Xconnect
// Xconnect data
type L2Tp_Sessions_Session_SessionApplicationData_Xconnect struct {
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

func (xconnect *L2Tp_Sessions_Session_SessionApplicationData_Xconnect) GetEntityData() *types.CommonEntityData {
    xconnect.EntityData.YFilter = xconnect.YFilter
    xconnect.EntityData.YangName = "xconnect"
    xconnect.EntityData.BundleName = "cisco_ios_xr"
    xconnect.EntityData.ParentYangName = "session-application-data"
    xconnect.EntityData.SegmentPath = "xconnect"
    xconnect.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    xconnect.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    xconnect.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    xconnect.EntityData.Children = make(map[string]types.YChild)
    xconnect.EntityData.Leafs = make(map[string]types.YLeaf)
    xconnect.EntityData.Leafs["circuit-name"] = types.YLeaf{"CircuitName", xconnect.CircuitName}
    xconnect.EntityData.Leafs["sessionvc-id"] = types.YLeaf{"SessionvcId", xconnect.SessionvcId}
    xconnect.EntityData.Leafs["is-circuit-state-up"] = types.YLeaf{"IsCircuitStateUp", xconnect.IsCircuitStateUp}
    xconnect.EntityData.Leafs["is-local-circuit-state-up"] = types.YLeaf{"IsLocalCircuitStateUp", xconnect.IsLocalCircuitStateUp}
    xconnect.EntityData.Leafs["is-remote-circuit-state-up"] = types.YLeaf{"IsRemoteCircuitStateUp", xconnect.IsRemoteCircuitStateUp}
    xconnect.EntityData.Leafs["ipv6-protocol-tunneling"] = types.YLeaf{"Ipv6ProtocolTunneling", xconnect.Ipv6ProtocolTunneling}
    return &(xconnect.EntityData)
}

// L2Tp_Sessions_Session_SessionApplicationData_Vpdn
// VPDN data
type L2Tp_Sessions_Session_SessionApplicationData_Vpdn struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Session username. The type is string.
    Username interface{}

    // Interface name. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}
}

func (vpdn *L2Tp_Sessions_Session_SessionApplicationData_Vpdn) GetEntityData() *types.CommonEntityData {
    vpdn.EntityData.YFilter = vpdn.YFilter
    vpdn.EntityData.YangName = "vpdn"
    vpdn.EntityData.BundleName = "cisco_ios_xr"
    vpdn.EntityData.ParentYangName = "session-application-data"
    vpdn.EntityData.SegmentPath = "vpdn"
    vpdn.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vpdn.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vpdn.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vpdn.EntityData.Children = make(map[string]types.YChild)
    vpdn.EntityData.Leafs = make(map[string]types.YLeaf)
    vpdn.EntityData.Leafs["username"] = types.YLeaf{"Username", vpdn.Username}
    vpdn.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", vpdn.InterfaceName}
    return &(vpdn.EntityData)
}

// L2Tp_Session
// L2TP control messages counters
type L2Tp_Session struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2TP session unavailable  information.
    Unavailable L2Tp_Session_Unavailable
}

func (session *L2Tp_Session) GetEntityData() *types.CommonEntityData {
    session.EntityData.YFilter = session.YFilter
    session.EntityData.YangName = "session"
    session.EntityData.BundleName = "cisco_ios_xr"
    session.EntityData.ParentYangName = "l2tp"
    session.EntityData.SegmentPath = "session"
    session.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    session.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    session.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    session.EntityData.Children = make(map[string]types.YChild)
    session.EntityData.Children["unavailable"] = types.YChild{"Unavailable", &session.Unavailable}
    session.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(session.EntityData)
}

// L2Tp_Session_Unavailable
// L2TP session unavailable  information
type L2Tp_Session_Unavailable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of session ID in hold database. The type is interface{} with range:
    // 0..4294967295.
    SessionsOnHold interface{}
}

func (unavailable *L2Tp_Session_Unavailable) GetEntityData() *types.CommonEntityData {
    unavailable.EntityData.YFilter = unavailable.YFilter
    unavailable.EntityData.YangName = "unavailable"
    unavailable.EntityData.BundleName = "cisco_ios_xr"
    unavailable.EntityData.ParentYangName = "session"
    unavailable.EntityData.SegmentPath = "unavailable"
    unavailable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unavailable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unavailable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unavailable.EntityData.Children = make(map[string]types.YChild)
    unavailable.EntityData.Leafs = make(map[string]types.YLeaf)
    unavailable.EntityData.Leafs["sessions-on-hold"] = types.YLeaf{"SessionsOnHold", unavailable.SessionsOnHold}
    return &(unavailable.EntityData)
}

// L2Tpv2
// l2tpv2
type L2Tpv2 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2TP control messages counters.
    Counters L2Tpv2_Counters

    // L2TP v2 statistics information.
    Statistics L2Tpv2_Statistics

    // L2TPv2 tunnel .
    Tunnel L2Tpv2_Tunnel

    // List of tunnel IDs.
    TunnelConfigurations L2Tpv2_TunnelConfigurations

    // Failure events leading to disconnection.
    CounterHistFail L2Tpv2_CounterHistFail

    // List of L2TP class names.
    Classes L2Tpv2_Classes

    // List of tunnel IDs.
    Tunnels L2Tpv2_Tunnels

    // List of session IDs.
    Sessions L2Tpv2_Sessions

    // L2TP control messages counters.
    Session L2Tpv2_Session
}

func (l2Tpv2 *L2Tpv2) GetEntityData() *types.CommonEntityData {
    l2Tpv2.EntityData.YFilter = l2Tpv2.YFilter
    l2Tpv2.EntityData.YangName = "l2tpv2"
    l2Tpv2.EntityData.BundleName = "cisco_ios_xr"
    l2Tpv2.EntityData.ParentYangName = "Cisco-IOS-XR-tunnel-l2tun-oper"
    l2Tpv2.EntityData.SegmentPath = "Cisco-IOS-XR-tunnel-l2tun-oper:l2tpv2"
    l2Tpv2.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    l2Tpv2.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    l2Tpv2.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    l2Tpv2.EntityData.Children = make(map[string]types.YChild)
    l2Tpv2.EntityData.Children["counters"] = types.YChild{"Counters", &l2Tpv2.Counters}
    l2Tpv2.EntityData.Children["statistics"] = types.YChild{"Statistics", &l2Tpv2.Statistics}
    l2Tpv2.EntityData.Children["tunnel"] = types.YChild{"Tunnel", &l2Tpv2.Tunnel}
    l2Tpv2.EntityData.Children["tunnel-configurations"] = types.YChild{"TunnelConfigurations", &l2Tpv2.TunnelConfigurations}
    l2Tpv2.EntityData.Children["counter-hist-fail"] = types.YChild{"CounterHistFail", &l2Tpv2.CounterHistFail}
    l2Tpv2.EntityData.Children["classes"] = types.YChild{"Classes", &l2Tpv2.Classes}
    l2Tpv2.EntityData.Children["tunnels"] = types.YChild{"Tunnels", &l2Tpv2.Tunnels}
    l2Tpv2.EntityData.Children["sessions"] = types.YChild{"Sessions", &l2Tpv2.Sessions}
    l2Tpv2.EntityData.Children["session"] = types.YChild{"Session", &l2Tpv2.Session}
    l2Tpv2.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(l2Tpv2.EntityData)
}

// L2Tpv2_Counters
// L2TP control messages counters
type L2Tpv2_Counters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2TP forwarding messages counters.
    Forwarding L2Tpv2_Counters_Forwarding

    // L2TP control messages counters.
    Control L2Tpv2_Counters_Control
}

func (counters *L2Tpv2_Counters) GetEntityData() *types.CommonEntityData {
    counters.EntityData.YFilter = counters.YFilter
    counters.EntityData.YangName = "counters"
    counters.EntityData.BundleName = "cisco_ios_xr"
    counters.EntityData.ParentYangName = "l2tpv2"
    counters.EntityData.SegmentPath = "counters"
    counters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    counters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    counters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    counters.EntityData.Children = make(map[string]types.YChild)
    counters.EntityData.Children["forwarding"] = types.YChild{"Forwarding", &counters.Forwarding}
    counters.EntityData.Children["control"] = types.YChild{"Control", &counters.Control}
    counters.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(counters.EntityData)
}

// L2Tpv2_Counters_Forwarding
// L2TP forwarding messages counters
type L2Tpv2_Counters_Forwarding struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of class and session IDs.
    Sessions L2Tpv2_Counters_Forwarding_Sessions
}

func (forwarding *L2Tpv2_Counters_Forwarding) GetEntityData() *types.CommonEntityData {
    forwarding.EntityData.YFilter = forwarding.YFilter
    forwarding.EntityData.YangName = "forwarding"
    forwarding.EntityData.BundleName = "cisco_ios_xr"
    forwarding.EntityData.ParentYangName = "counters"
    forwarding.EntityData.SegmentPath = "forwarding"
    forwarding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    forwarding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    forwarding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    forwarding.EntityData.Children = make(map[string]types.YChild)
    forwarding.EntityData.Children["sessions"] = types.YChild{"Sessions", &forwarding.Sessions}
    forwarding.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(forwarding.EntityData)
}

// L2Tpv2_Counters_Forwarding_Sessions
// List of class and session IDs
type L2Tpv2_Counters_Forwarding_Sessions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2TP information for a particular session. The type is slice of
    // L2Tpv2_Counters_Forwarding_Sessions_Session.
    Session []L2Tpv2_Counters_Forwarding_Sessions_Session
}

func (sessions *L2Tpv2_Counters_Forwarding_Sessions) GetEntityData() *types.CommonEntityData {
    sessions.EntityData.YFilter = sessions.YFilter
    sessions.EntityData.YangName = "sessions"
    sessions.EntityData.BundleName = "cisco_ios_xr"
    sessions.EntityData.ParentYangName = "forwarding"
    sessions.EntityData.SegmentPath = "sessions"
    sessions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessions.EntityData.Children = make(map[string]types.YChild)
    sessions.EntityData.Children["session"] = types.YChild{"Session", nil}
    for i := range sessions.Session {
        sessions.EntityData.Children[types.GetSegmentPath(&sessions.Session[i])] = types.YChild{"Session", &sessions.Session[i]}
    }
    sessions.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sessions.EntityData)
}

// L2Tpv2_Counters_Forwarding_Sessions_Session
// L2TP information for a particular session
type L2Tpv2_Counters_Forwarding_Sessions_Session struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Local tunnel ID. The type is interface{} with
    // range: -2147483648..2147483647.
    TunnelId interface{}

    // This attribute is a key. Local session ID. The type is interface{} with
    // range: -2147483648..2147483647.
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

func (session *L2Tpv2_Counters_Forwarding_Sessions_Session) GetEntityData() *types.CommonEntityData {
    session.EntityData.YFilter = session.YFilter
    session.EntityData.YangName = "session"
    session.EntityData.BundleName = "cisco_ios_xr"
    session.EntityData.ParentYangName = "sessions"
    session.EntityData.SegmentPath = "session" + "[tunnel-id='" + fmt.Sprintf("%v", session.TunnelId) + "']" + "[session-id='" + fmt.Sprintf("%v", session.SessionId) + "']"
    session.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    session.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    session.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    session.EntityData.Children = make(map[string]types.YChild)
    session.EntityData.Leafs = make(map[string]types.YLeaf)
    session.EntityData.Leafs["tunnel-id"] = types.YLeaf{"TunnelId", session.TunnelId}
    session.EntityData.Leafs["session-id"] = types.YLeaf{"SessionId", session.SessionId}
    session.EntityData.Leafs["remote-session-id"] = types.YLeaf{"RemoteSessionId", session.RemoteSessionId}
    session.EntityData.Leafs["in-packets"] = types.YLeaf{"InPackets", session.InPackets}
    session.EntityData.Leafs["out-packets"] = types.YLeaf{"OutPackets", session.OutPackets}
    session.EntityData.Leafs["in-bytes"] = types.YLeaf{"InBytes", session.InBytes}
    session.EntityData.Leafs["out-bytes"] = types.YLeaf{"OutBytes", session.OutBytes}
    return &(session.EntityData)
}

// L2Tpv2_Counters_Control
// L2TP control messages counters
type L2Tpv2_Counters_Control struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2TP control tunnel messages counters.
    TunnelXr L2Tpv2_Counters_Control_TunnelXr

    // Table of tunnel IDs of control message counters.
    Tunnels L2Tpv2_Counters_Control_Tunnels
}

func (control *L2Tpv2_Counters_Control) GetEntityData() *types.CommonEntityData {
    control.EntityData.YFilter = control.YFilter
    control.EntityData.YangName = "control"
    control.EntityData.BundleName = "cisco_ios_xr"
    control.EntityData.ParentYangName = "counters"
    control.EntityData.SegmentPath = "control"
    control.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    control.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    control.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    control.EntityData.Children = make(map[string]types.YChild)
    control.EntityData.Children["tunnel-xr"] = types.YChild{"TunnelXr", &control.TunnelXr}
    control.EntityData.Children["tunnels"] = types.YChild{"Tunnels", &control.Tunnels}
    control.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(control.EntityData)
}

// L2Tpv2_Counters_Control_TunnelXr
// L2TP control tunnel messages counters
type L2Tpv2_Counters_Control_TunnelXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tunnel authentication counters.
    Authentication L2Tpv2_Counters_Control_TunnelXr_Authentication

    // Tunnel counters.
    Global L2Tpv2_Counters_Control_TunnelXr_Global
}

func (tunnelXr *L2Tpv2_Counters_Control_TunnelXr) GetEntityData() *types.CommonEntityData {
    tunnelXr.EntityData.YFilter = tunnelXr.YFilter
    tunnelXr.EntityData.YangName = "tunnel-xr"
    tunnelXr.EntityData.BundleName = "cisco_ios_xr"
    tunnelXr.EntityData.ParentYangName = "control"
    tunnelXr.EntityData.SegmentPath = "tunnel-xr"
    tunnelXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnelXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnelXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnelXr.EntityData.Children = make(map[string]types.YChild)
    tunnelXr.EntityData.Children["authentication"] = types.YChild{"Authentication", &tunnelXr.Authentication}
    tunnelXr.EntityData.Children["global"] = types.YChild{"Global", &tunnelXr.Global}
    tunnelXr.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(tunnelXr.EntityData)
}

// L2Tpv2_Counters_Control_TunnelXr_Authentication
// Tunnel authentication counters
type L2Tpv2_Counters_Control_TunnelXr_Authentication struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Nonce AVP statistics.
    NonceAvp L2Tpv2_Counters_Control_TunnelXr_Authentication_NonceAvp

    // Common digest statistics.
    CommonDigest L2Tpv2_Counters_Control_TunnelXr_Authentication_CommonDigest

    // Primary digest statistics.
    PrimaryDigest L2Tpv2_Counters_Control_TunnelXr_Authentication_PrimaryDigest

    // Secondary digest statistics.
    SecondaryDigest L2Tpv2_Counters_Control_TunnelXr_Authentication_SecondaryDigest

    // Integrity check statistics.
    IntegrityCheck L2Tpv2_Counters_Control_TunnelXr_Authentication_IntegrityCheck

    // Local secret statistics.
    LocalSecret L2Tpv2_Counters_Control_TunnelXr_Authentication_LocalSecret

    // Challenge AVP statistics.
    ChallengeAvp L2Tpv2_Counters_Control_TunnelXr_Authentication_ChallengeAvp

    // Challenge response statistics.
    ChallengeReponse L2Tpv2_Counters_Control_TunnelXr_Authentication_ChallengeReponse

    // Overall statistics.
    OverallStatistics L2Tpv2_Counters_Control_TunnelXr_Authentication_OverallStatistics
}

func (authentication *L2Tpv2_Counters_Control_TunnelXr_Authentication) GetEntityData() *types.CommonEntityData {
    authentication.EntityData.YFilter = authentication.YFilter
    authentication.EntityData.YangName = "authentication"
    authentication.EntityData.BundleName = "cisco_ios_xr"
    authentication.EntityData.ParentYangName = "tunnel-xr"
    authentication.EntityData.SegmentPath = "authentication"
    authentication.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authentication.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authentication.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authentication.EntityData.Children = make(map[string]types.YChild)
    authentication.EntityData.Children["nonce-avp"] = types.YChild{"NonceAvp", &authentication.NonceAvp}
    authentication.EntityData.Children["common-digest"] = types.YChild{"CommonDigest", &authentication.CommonDigest}
    authentication.EntityData.Children["primary-digest"] = types.YChild{"PrimaryDigest", &authentication.PrimaryDigest}
    authentication.EntityData.Children["secondary-digest"] = types.YChild{"SecondaryDigest", &authentication.SecondaryDigest}
    authentication.EntityData.Children["integrity-check"] = types.YChild{"IntegrityCheck", &authentication.IntegrityCheck}
    authentication.EntityData.Children["local-secret"] = types.YChild{"LocalSecret", &authentication.LocalSecret}
    authentication.EntityData.Children["challenge-avp"] = types.YChild{"ChallengeAvp", &authentication.ChallengeAvp}
    authentication.EntityData.Children["challenge-reponse"] = types.YChild{"ChallengeReponse", &authentication.ChallengeReponse}
    authentication.EntityData.Children["overall-statistics"] = types.YChild{"OverallStatistics", &authentication.OverallStatistics}
    authentication.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(authentication.EntityData)
}

// L2Tpv2_Counters_Control_TunnelXr_Authentication_NonceAvp
// Nonce AVP statistics
type L2Tpv2_Counters_Control_TunnelXr_Authentication_NonceAvp struct {
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

func (nonceAvp *L2Tpv2_Counters_Control_TunnelXr_Authentication_NonceAvp) GetEntityData() *types.CommonEntityData {
    nonceAvp.EntityData.YFilter = nonceAvp.YFilter
    nonceAvp.EntityData.YangName = "nonce-avp"
    nonceAvp.EntityData.BundleName = "cisco_ios_xr"
    nonceAvp.EntityData.ParentYangName = "authentication"
    nonceAvp.EntityData.SegmentPath = "nonce-avp"
    nonceAvp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nonceAvp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nonceAvp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nonceAvp.EntityData.Children = make(map[string]types.YChild)
    nonceAvp.EntityData.Leafs = make(map[string]types.YLeaf)
    nonceAvp.EntityData.Leafs["validate"] = types.YLeaf{"Validate", nonceAvp.Validate}
    nonceAvp.EntityData.Leafs["bad-hash"] = types.YLeaf{"BadHash", nonceAvp.BadHash}
    nonceAvp.EntityData.Leafs["bad-length"] = types.YLeaf{"BadLength", nonceAvp.BadLength}
    nonceAvp.EntityData.Leafs["ignored"] = types.YLeaf{"Ignored", nonceAvp.Ignored}
    nonceAvp.EntityData.Leafs["missing"] = types.YLeaf{"Missing", nonceAvp.Missing}
    nonceAvp.EntityData.Leafs["passed"] = types.YLeaf{"Passed", nonceAvp.Passed}
    nonceAvp.EntityData.Leafs["failed"] = types.YLeaf{"Failed", nonceAvp.Failed}
    nonceAvp.EntityData.Leafs["skipped"] = types.YLeaf{"Skipped", nonceAvp.Skipped}
    nonceAvp.EntityData.Leafs["generate-response-failures"] = types.YLeaf{"GenerateResponseFailures", nonceAvp.GenerateResponseFailures}
    nonceAvp.EntityData.Leafs["unexpected"] = types.YLeaf{"Unexpected", nonceAvp.Unexpected}
    nonceAvp.EntityData.Leafs["unexpected-zlb"] = types.YLeaf{"UnexpectedZlb", nonceAvp.UnexpectedZlb}
    return &(nonceAvp.EntityData)
}

// L2Tpv2_Counters_Control_TunnelXr_Authentication_CommonDigest
// Common digest statistics
type L2Tpv2_Counters_Control_TunnelXr_Authentication_CommonDigest struct {
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

func (commonDigest *L2Tpv2_Counters_Control_TunnelXr_Authentication_CommonDigest) GetEntityData() *types.CommonEntityData {
    commonDigest.EntityData.YFilter = commonDigest.YFilter
    commonDigest.EntityData.YangName = "common-digest"
    commonDigest.EntityData.BundleName = "cisco_ios_xr"
    commonDigest.EntityData.ParentYangName = "authentication"
    commonDigest.EntityData.SegmentPath = "common-digest"
    commonDigest.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    commonDigest.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    commonDigest.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    commonDigest.EntityData.Children = make(map[string]types.YChild)
    commonDigest.EntityData.Leafs = make(map[string]types.YLeaf)
    commonDigest.EntityData.Leafs["validate"] = types.YLeaf{"Validate", commonDigest.Validate}
    commonDigest.EntityData.Leafs["bad-hash"] = types.YLeaf{"BadHash", commonDigest.BadHash}
    commonDigest.EntityData.Leafs["bad-length"] = types.YLeaf{"BadLength", commonDigest.BadLength}
    commonDigest.EntityData.Leafs["ignored"] = types.YLeaf{"Ignored", commonDigest.Ignored}
    commonDigest.EntityData.Leafs["missing"] = types.YLeaf{"Missing", commonDigest.Missing}
    commonDigest.EntityData.Leafs["passed"] = types.YLeaf{"Passed", commonDigest.Passed}
    commonDigest.EntityData.Leafs["failed"] = types.YLeaf{"Failed", commonDigest.Failed}
    commonDigest.EntityData.Leafs["skipped"] = types.YLeaf{"Skipped", commonDigest.Skipped}
    commonDigest.EntityData.Leafs["generate-response-failures"] = types.YLeaf{"GenerateResponseFailures", commonDigest.GenerateResponseFailures}
    commonDigest.EntityData.Leafs["unexpected"] = types.YLeaf{"Unexpected", commonDigest.Unexpected}
    commonDigest.EntityData.Leafs["unexpected-zlb"] = types.YLeaf{"UnexpectedZlb", commonDigest.UnexpectedZlb}
    return &(commonDigest.EntityData)
}

// L2Tpv2_Counters_Control_TunnelXr_Authentication_PrimaryDigest
// Primary digest statistics
type L2Tpv2_Counters_Control_TunnelXr_Authentication_PrimaryDigest struct {
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

func (primaryDigest *L2Tpv2_Counters_Control_TunnelXr_Authentication_PrimaryDigest) GetEntityData() *types.CommonEntityData {
    primaryDigest.EntityData.YFilter = primaryDigest.YFilter
    primaryDigest.EntityData.YangName = "primary-digest"
    primaryDigest.EntityData.BundleName = "cisco_ios_xr"
    primaryDigest.EntityData.ParentYangName = "authentication"
    primaryDigest.EntityData.SegmentPath = "primary-digest"
    primaryDigest.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    primaryDigest.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    primaryDigest.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    primaryDigest.EntityData.Children = make(map[string]types.YChild)
    primaryDigest.EntityData.Leafs = make(map[string]types.YLeaf)
    primaryDigest.EntityData.Leafs["validate"] = types.YLeaf{"Validate", primaryDigest.Validate}
    primaryDigest.EntityData.Leafs["bad-hash"] = types.YLeaf{"BadHash", primaryDigest.BadHash}
    primaryDigest.EntityData.Leafs["bad-length"] = types.YLeaf{"BadLength", primaryDigest.BadLength}
    primaryDigest.EntityData.Leafs["ignored"] = types.YLeaf{"Ignored", primaryDigest.Ignored}
    primaryDigest.EntityData.Leafs["missing"] = types.YLeaf{"Missing", primaryDigest.Missing}
    primaryDigest.EntityData.Leafs["passed"] = types.YLeaf{"Passed", primaryDigest.Passed}
    primaryDigest.EntityData.Leafs["failed"] = types.YLeaf{"Failed", primaryDigest.Failed}
    primaryDigest.EntityData.Leafs["skipped"] = types.YLeaf{"Skipped", primaryDigest.Skipped}
    primaryDigest.EntityData.Leafs["generate-response-failures"] = types.YLeaf{"GenerateResponseFailures", primaryDigest.GenerateResponseFailures}
    primaryDigest.EntityData.Leafs["unexpected"] = types.YLeaf{"Unexpected", primaryDigest.Unexpected}
    primaryDigest.EntityData.Leafs["unexpected-zlb"] = types.YLeaf{"UnexpectedZlb", primaryDigest.UnexpectedZlb}
    return &(primaryDigest.EntityData)
}

// L2Tpv2_Counters_Control_TunnelXr_Authentication_SecondaryDigest
// Secondary digest statistics
type L2Tpv2_Counters_Control_TunnelXr_Authentication_SecondaryDigest struct {
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

func (secondaryDigest *L2Tpv2_Counters_Control_TunnelXr_Authentication_SecondaryDigest) GetEntityData() *types.CommonEntityData {
    secondaryDigest.EntityData.YFilter = secondaryDigest.YFilter
    secondaryDigest.EntityData.YangName = "secondary-digest"
    secondaryDigest.EntityData.BundleName = "cisco_ios_xr"
    secondaryDigest.EntityData.ParentYangName = "authentication"
    secondaryDigest.EntityData.SegmentPath = "secondary-digest"
    secondaryDigest.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    secondaryDigest.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    secondaryDigest.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    secondaryDigest.EntityData.Children = make(map[string]types.YChild)
    secondaryDigest.EntityData.Leafs = make(map[string]types.YLeaf)
    secondaryDigest.EntityData.Leafs["validate"] = types.YLeaf{"Validate", secondaryDigest.Validate}
    secondaryDigest.EntityData.Leafs["bad-hash"] = types.YLeaf{"BadHash", secondaryDigest.BadHash}
    secondaryDigest.EntityData.Leafs["bad-length"] = types.YLeaf{"BadLength", secondaryDigest.BadLength}
    secondaryDigest.EntityData.Leafs["ignored"] = types.YLeaf{"Ignored", secondaryDigest.Ignored}
    secondaryDigest.EntityData.Leafs["missing"] = types.YLeaf{"Missing", secondaryDigest.Missing}
    secondaryDigest.EntityData.Leafs["passed"] = types.YLeaf{"Passed", secondaryDigest.Passed}
    secondaryDigest.EntityData.Leafs["failed"] = types.YLeaf{"Failed", secondaryDigest.Failed}
    secondaryDigest.EntityData.Leafs["skipped"] = types.YLeaf{"Skipped", secondaryDigest.Skipped}
    secondaryDigest.EntityData.Leafs["generate-response-failures"] = types.YLeaf{"GenerateResponseFailures", secondaryDigest.GenerateResponseFailures}
    secondaryDigest.EntityData.Leafs["unexpected"] = types.YLeaf{"Unexpected", secondaryDigest.Unexpected}
    secondaryDigest.EntityData.Leafs["unexpected-zlb"] = types.YLeaf{"UnexpectedZlb", secondaryDigest.UnexpectedZlb}
    return &(secondaryDigest.EntityData)
}

// L2Tpv2_Counters_Control_TunnelXr_Authentication_IntegrityCheck
// Integrity check statistics
type L2Tpv2_Counters_Control_TunnelXr_Authentication_IntegrityCheck struct {
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

func (integrityCheck *L2Tpv2_Counters_Control_TunnelXr_Authentication_IntegrityCheck) GetEntityData() *types.CommonEntityData {
    integrityCheck.EntityData.YFilter = integrityCheck.YFilter
    integrityCheck.EntityData.YangName = "integrity-check"
    integrityCheck.EntityData.BundleName = "cisco_ios_xr"
    integrityCheck.EntityData.ParentYangName = "authentication"
    integrityCheck.EntityData.SegmentPath = "integrity-check"
    integrityCheck.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    integrityCheck.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    integrityCheck.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    integrityCheck.EntityData.Children = make(map[string]types.YChild)
    integrityCheck.EntityData.Leafs = make(map[string]types.YLeaf)
    integrityCheck.EntityData.Leafs["validate"] = types.YLeaf{"Validate", integrityCheck.Validate}
    integrityCheck.EntityData.Leafs["bad-hash"] = types.YLeaf{"BadHash", integrityCheck.BadHash}
    integrityCheck.EntityData.Leafs["bad-length"] = types.YLeaf{"BadLength", integrityCheck.BadLength}
    integrityCheck.EntityData.Leafs["ignored"] = types.YLeaf{"Ignored", integrityCheck.Ignored}
    integrityCheck.EntityData.Leafs["missing"] = types.YLeaf{"Missing", integrityCheck.Missing}
    integrityCheck.EntityData.Leafs["passed"] = types.YLeaf{"Passed", integrityCheck.Passed}
    integrityCheck.EntityData.Leafs["failed"] = types.YLeaf{"Failed", integrityCheck.Failed}
    integrityCheck.EntityData.Leafs["skipped"] = types.YLeaf{"Skipped", integrityCheck.Skipped}
    integrityCheck.EntityData.Leafs["generate-response-failures"] = types.YLeaf{"GenerateResponseFailures", integrityCheck.GenerateResponseFailures}
    integrityCheck.EntityData.Leafs["unexpected"] = types.YLeaf{"Unexpected", integrityCheck.Unexpected}
    integrityCheck.EntityData.Leafs["unexpected-zlb"] = types.YLeaf{"UnexpectedZlb", integrityCheck.UnexpectedZlb}
    return &(integrityCheck.EntityData)
}

// L2Tpv2_Counters_Control_TunnelXr_Authentication_LocalSecret
// Local secret statistics
type L2Tpv2_Counters_Control_TunnelXr_Authentication_LocalSecret struct {
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

func (localSecret *L2Tpv2_Counters_Control_TunnelXr_Authentication_LocalSecret) GetEntityData() *types.CommonEntityData {
    localSecret.EntityData.YFilter = localSecret.YFilter
    localSecret.EntityData.YangName = "local-secret"
    localSecret.EntityData.BundleName = "cisco_ios_xr"
    localSecret.EntityData.ParentYangName = "authentication"
    localSecret.EntityData.SegmentPath = "local-secret"
    localSecret.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localSecret.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localSecret.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localSecret.EntityData.Children = make(map[string]types.YChild)
    localSecret.EntityData.Leafs = make(map[string]types.YLeaf)
    localSecret.EntityData.Leafs["validate"] = types.YLeaf{"Validate", localSecret.Validate}
    localSecret.EntityData.Leafs["bad-hash"] = types.YLeaf{"BadHash", localSecret.BadHash}
    localSecret.EntityData.Leafs["bad-length"] = types.YLeaf{"BadLength", localSecret.BadLength}
    localSecret.EntityData.Leafs["ignored"] = types.YLeaf{"Ignored", localSecret.Ignored}
    localSecret.EntityData.Leafs["missing"] = types.YLeaf{"Missing", localSecret.Missing}
    localSecret.EntityData.Leafs["passed"] = types.YLeaf{"Passed", localSecret.Passed}
    localSecret.EntityData.Leafs["failed"] = types.YLeaf{"Failed", localSecret.Failed}
    localSecret.EntityData.Leafs["skipped"] = types.YLeaf{"Skipped", localSecret.Skipped}
    localSecret.EntityData.Leafs["generate-response-failures"] = types.YLeaf{"GenerateResponseFailures", localSecret.GenerateResponseFailures}
    localSecret.EntityData.Leafs["unexpected"] = types.YLeaf{"Unexpected", localSecret.Unexpected}
    localSecret.EntityData.Leafs["unexpected-zlb"] = types.YLeaf{"UnexpectedZlb", localSecret.UnexpectedZlb}
    return &(localSecret.EntityData)
}

// L2Tpv2_Counters_Control_TunnelXr_Authentication_ChallengeAvp
// Challenge AVP statistics
type L2Tpv2_Counters_Control_TunnelXr_Authentication_ChallengeAvp struct {
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

func (challengeAvp *L2Tpv2_Counters_Control_TunnelXr_Authentication_ChallengeAvp) GetEntityData() *types.CommonEntityData {
    challengeAvp.EntityData.YFilter = challengeAvp.YFilter
    challengeAvp.EntityData.YangName = "challenge-avp"
    challengeAvp.EntityData.BundleName = "cisco_ios_xr"
    challengeAvp.EntityData.ParentYangName = "authentication"
    challengeAvp.EntityData.SegmentPath = "challenge-avp"
    challengeAvp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    challengeAvp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    challengeAvp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    challengeAvp.EntityData.Children = make(map[string]types.YChild)
    challengeAvp.EntityData.Leafs = make(map[string]types.YLeaf)
    challengeAvp.EntityData.Leafs["validate"] = types.YLeaf{"Validate", challengeAvp.Validate}
    challengeAvp.EntityData.Leafs["bad-hash"] = types.YLeaf{"BadHash", challengeAvp.BadHash}
    challengeAvp.EntityData.Leafs["bad-length"] = types.YLeaf{"BadLength", challengeAvp.BadLength}
    challengeAvp.EntityData.Leafs["ignored"] = types.YLeaf{"Ignored", challengeAvp.Ignored}
    challengeAvp.EntityData.Leafs["missing"] = types.YLeaf{"Missing", challengeAvp.Missing}
    challengeAvp.EntityData.Leafs["passed"] = types.YLeaf{"Passed", challengeAvp.Passed}
    challengeAvp.EntityData.Leafs["failed"] = types.YLeaf{"Failed", challengeAvp.Failed}
    challengeAvp.EntityData.Leafs["skipped"] = types.YLeaf{"Skipped", challengeAvp.Skipped}
    challengeAvp.EntityData.Leafs["generate-response-failures"] = types.YLeaf{"GenerateResponseFailures", challengeAvp.GenerateResponseFailures}
    challengeAvp.EntityData.Leafs["unexpected"] = types.YLeaf{"Unexpected", challengeAvp.Unexpected}
    challengeAvp.EntityData.Leafs["unexpected-zlb"] = types.YLeaf{"UnexpectedZlb", challengeAvp.UnexpectedZlb}
    return &(challengeAvp.EntityData)
}

// L2Tpv2_Counters_Control_TunnelXr_Authentication_ChallengeReponse
// Challenge response statistics
type L2Tpv2_Counters_Control_TunnelXr_Authentication_ChallengeReponse struct {
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

func (challengeReponse *L2Tpv2_Counters_Control_TunnelXr_Authentication_ChallengeReponse) GetEntityData() *types.CommonEntityData {
    challengeReponse.EntityData.YFilter = challengeReponse.YFilter
    challengeReponse.EntityData.YangName = "challenge-reponse"
    challengeReponse.EntityData.BundleName = "cisco_ios_xr"
    challengeReponse.EntityData.ParentYangName = "authentication"
    challengeReponse.EntityData.SegmentPath = "challenge-reponse"
    challengeReponse.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    challengeReponse.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    challengeReponse.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    challengeReponse.EntityData.Children = make(map[string]types.YChild)
    challengeReponse.EntityData.Leafs = make(map[string]types.YLeaf)
    challengeReponse.EntityData.Leafs["validate"] = types.YLeaf{"Validate", challengeReponse.Validate}
    challengeReponse.EntityData.Leafs["bad-hash"] = types.YLeaf{"BadHash", challengeReponse.BadHash}
    challengeReponse.EntityData.Leafs["bad-length"] = types.YLeaf{"BadLength", challengeReponse.BadLength}
    challengeReponse.EntityData.Leafs["ignored"] = types.YLeaf{"Ignored", challengeReponse.Ignored}
    challengeReponse.EntityData.Leafs["missing"] = types.YLeaf{"Missing", challengeReponse.Missing}
    challengeReponse.EntityData.Leafs["passed"] = types.YLeaf{"Passed", challengeReponse.Passed}
    challengeReponse.EntityData.Leafs["failed"] = types.YLeaf{"Failed", challengeReponse.Failed}
    challengeReponse.EntityData.Leafs["skipped"] = types.YLeaf{"Skipped", challengeReponse.Skipped}
    challengeReponse.EntityData.Leafs["generate-response-failures"] = types.YLeaf{"GenerateResponseFailures", challengeReponse.GenerateResponseFailures}
    challengeReponse.EntityData.Leafs["unexpected"] = types.YLeaf{"Unexpected", challengeReponse.Unexpected}
    challengeReponse.EntityData.Leafs["unexpected-zlb"] = types.YLeaf{"UnexpectedZlb", challengeReponse.UnexpectedZlb}
    return &(challengeReponse.EntityData)
}

// L2Tpv2_Counters_Control_TunnelXr_Authentication_OverallStatistics
// Overall statistics
type L2Tpv2_Counters_Control_TunnelXr_Authentication_OverallStatistics struct {
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

func (overallStatistics *L2Tpv2_Counters_Control_TunnelXr_Authentication_OverallStatistics) GetEntityData() *types.CommonEntityData {
    overallStatistics.EntityData.YFilter = overallStatistics.YFilter
    overallStatistics.EntityData.YangName = "overall-statistics"
    overallStatistics.EntityData.BundleName = "cisco_ios_xr"
    overallStatistics.EntityData.ParentYangName = "authentication"
    overallStatistics.EntityData.SegmentPath = "overall-statistics"
    overallStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    overallStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    overallStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    overallStatistics.EntityData.Children = make(map[string]types.YChild)
    overallStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    overallStatistics.EntityData.Leafs["validate"] = types.YLeaf{"Validate", overallStatistics.Validate}
    overallStatistics.EntityData.Leafs["bad-hash"] = types.YLeaf{"BadHash", overallStatistics.BadHash}
    overallStatistics.EntityData.Leafs["bad-length"] = types.YLeaf{"BadLength", overallStatistics.BadLength}
    overallStatistics.EntityData.Leafs["ignored"] = types.YLeaf{"Ignored", overallStatistics.Ignored}
    overallStatistics.EntityData.Leafs["missing"] = types.YLeaf{"Missing", overallStatistics.Missing}
    overallStatistics.EntityData.Leafs["passed"] = types.YLeaf{"Passed", overallStatistics.Passed}
    overallStatistics.EntityData.Leafs["failed"] = types.YLeaf{"Failed", overallStatistics.Failed}
    overallStatistics.EntityData.Leafs["skipped"] = types.YLeaf{"Skipped", overallStatistics.Skipped}
    overallStatistics.EntityData.Leafs["generate-response-failures"] = types.YLeaf{"GenerateResponseFailures", overallStatistics.GenerateResponseFailures}
    overallStatistics.EntityData.Leafs["unexpected"] = types.YLeaf{"Unexpected", overallStatistics.Unexpected}
    overallStatistics.EntityData.Leafs["unexpected-zlb"] = types.YLeaf{"UnexpectedZlb", overallStatistics.UnexpectedZlb}
    return &(overallStatistics.EntityData)
}

// L2Tpv2_Counters_Control_TunnelXr_Global
// Tunnel counters
type L2Tpv2_Counters_Control_TunnelXr_Global struct {
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
    Transmit L2Tpv2_Counters_Control_TunnelXr_Global_Transmit

    // Re transmit data.
    Retransmit L2Tpv2_Counters_Control_TunnelXr_Global_Retransmit

    // Received data.
    Received L2Tpv2_Counters_Control_TunnelXr_Global_Received

    // Drop data.
    Drop L2Tpv2_Counters_Control_TunnelXr_Global_Drop
}

func (global *L2Tpv2_Counters_Control_TunnelXr_Global) GetEntityData() *types.CommonEntityData {
    global.EntityData.YFilter = global.YFilter
    global.EntityData.YangName = "global"
    global.EntityData.BundleName = "cisco_ios_xr"
    global.EntityData.ParentYangName = "tunnel-xr"
    global.EntityData.SegmentPath = "global"
    global.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    global.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    global.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    global.EntityData.Children = make(map[string]types.YChild)
    global.EntityData.Children["transmit"] = types.YChild{"Transmit", &global.Transmit}
    global.EntityData.Children["retransmit"] = types.YChild{"Retransmit", &global.Retransmit}
    global.EntityData.Children["received"] = types.YChild{"Received", &global.Received}
    global.EntityData.Children["drop"] = types.YChild{"Drop", &global.Drop}
    global.EntityData.Leafs = make(map[string]types.YLeaf)
    global.EntityData.Leafs["total-transmit"] = types.YLeaf{"TotalTransmit", global.TotalTransmit}
    global.EntityData.Leafs["total-retransmit"] = types.YLeaf{"TotalRetransmit", global.TotalRetransmit}
    global.EntityData.Leafs["total-received"] = types.YLeaf{"TotalReceived", global.TotalReceived}
    global.EntityData.Leafs["total-drop"] = types.YLeaf{"TotalDrop", global.TotalDrop}
    return &(global.EntityData)
}

// L2Tpv2_Counters_Control_TunnelXr_Global_Transmit
// Transmit data
type L2Tpv2_Counters_Control_TunnelXr_Global_Transmit struct {
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

func (transmit *L2Tpv2_Counters_Control_TunnelXr_Global_Transmit) GetEntityData() *types.CommonEntityData {
    transmit.EntityData.YFilter = transmit.YFilter
    transmit.EntityData.YangName = "transmit"
    transmit.EntityData.BundleName = "cisco_ios_xr"
    transmit.EntityData.ParentYangName = "global"
    transmit.EntityData.SegmentPath = "transmit"
    transmit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transmit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transmit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transmit.EntityData.Children = make(map[string]types.YChild)
    transmit.EntityData.Leafs = make(map[string]types.YLeaf)
    transmit.EntityData.Leafs["unknown-packets"] = types.YLeaf{"UnknownPackets", transmit.UnknownPackets}
    transmit.EntityData.Leafs["zero-length-body-packets"] = types.YLeaf{"ZeroLengthBodyPackets", transmit.ZeroLengthBodyPackets}
    transmit.EntityData.Leafs["start-control-connection-requests"] = types.YLeaf{"StartControlConnectionRequests", transmit.StartControlConnectionRequests}
    transmit.EntityData.Leafs["start-control-connection-replies"] = types.YLeaf{"StartControlConnectionReplies", transmit.StartControlConnectionReplies}
    transmit.EntityData.Leafs["start-control-connection-notifications"] = types.YLeaf{"StartControlConnectionNotifications", transmit.StartControlConnectionNotifications}
    transmit.EntityData.Leafs["stop-control-connection-notifications"] = types.YLeaf{"StopControlConnectionNotifications", transmit.StopControlConnectionNotifications}
    transmit.EntityData.Leafs["hello-packets"] = types.YLeaf{"HelloPackets", transmit.HelloPackets}
    transmit.EntityData.Leafs["outgoing-call-requests"] = types.YLeaf{"OutgoingCallRequests", transmit.OutgoingCallRequests}
    transmit.EntityData.Leafs["outgoing-call-replies"] = types.YLeaf{"OutgoingCallReplies", transmit.OutgoingCallReplies}
    transmit.EntityData.Leafs["outgoing-call-connected-packets"] = types.YLeaf{"OutgoingCallConnectedPackets", transmit.OutgoingCallConnectedPackets}
    transmit.EntityData.Leafs["incoming-call-requests"] = types.YLeaf{"IncomingCallRequests", transmit.IncomingCallRequests}
    transmit.EntityData.Leafs["incoming-call-replies"] = types.YLeaf{"IncomingCallReplies", transmit.IncomingCallReplies}
    transmit.EntityData.Leafs["incoming-call-connected-packets"] = types.YLeaf{"IncomingCallConnectedPackets", transmit.IncomingCallConnectedPackets}
    transmit.EntityData.Leafs["call-disconnect-notify-packets"] = types.YLeaf{"CallDisconnectNotifyPackets", transmit.CallDisconnectNotifyPackets}
    transmit.EntityData.Leafs["wan-error-notify-packets"] = types.YLeaf{"WanErrorNotifyPackets", transmit.WanErrorNotifyPackets}
    transmit.EntityData.Leafs["set-link-info-packets"] = types.YLeaf{"SetLinkInfoPackets", transmit.SetLinkInfoPackets}
    transmit.EntityData.Leafs["service-relay-requests"] = types.YLeaf{"ServiceRelayRequests", transmit.ServiceRelayRequests}
    transmit.EntityData.Leafs["service-relay-replies"] = types.YLeaf{"ServiceRelayReplies", transmit.ServiceRelayReplies}
    transmit.EntityData.Leafs["acknowledgement-packets"] = types.YLeaf{"AcknowledgementPackets", transmit.AcknowledgementPackets}
    return &(transmit.EntityData)
}

// L2Tpv2_Counters_Control_TunnelXr_Global_Retransmit
// Re transmit data
type L2Tpv2_Counters_Control_TunnelXr_Global_Retransmit struct {
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

func (retransmit *L2Tpv2_Counters_Control_TunnelXr_Global_Retransmit) GetEntityData() *types.CommonEntityData {
    retransmit.EntityData.YFilter = retransmit.YFilter
    retransmit.EntityData.YangName = "retransmit"
    retransmit.EntityData.BundleName = "cisco_ios_xr"
    retransmit.EntityData.ParentYangName = "global"
    retransmit.EntityData.SegmentPath = "retransmit"
    retransmit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    retransmit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    retransmit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    retransmit.EntityData.Children = make(map[string]types.YChild)
    retransmit.EntityData.Leafs = make(map[string]types.YLeaf)
    retransmit.EntityData.Leafs["unknown-packets"] = types.YLeaf{"UnknownPackets", retransmit.UnknownPackets}
    retransmit.EntityData.Leafs["zero-length-body-packets"] = types.YLeaf{"ZeroLengthBodyPackets", retransmit.ZeroLengthBodyPackets}
    retransmit.EntityData.Leafs["start-control-connection-requests"] = types.YLeaf{"StartControlConnectionRequests", retransmit.StartControlConnectionRequests}
    retransmit.EntityData.Leafs["start-control-connection-replies"] = types.YLeaf{"StartControlConnectionReplies", retransmit.StartControlConnectionReplies}
    retransmit.EntityData.Leafs["start-control-connection-notifications"] = types.YLeaf{"StartControlConnectionNotifications", retransmit.StartControlConnectionNotifications}
    retransmit.EntityData.Leafs["stop-control-connection-notifications"] = types.YLeaf{"StopControlConnectionNotifications", retransmit.StopControlConnectionNotifications}
    retransmit.EntityData.Leafs["hello-packets"] = types.YLeaf{"HelloPackets", retransmit.HelloPackets}
    retransmit.EntityData.Leafs["outgoing-call-requests"] = types.YLeaf{"OutgoingCallRequests", retransmit.OutgoingCallRequests}
    retransmit.EntityData.Leafs["outgoing-call-replies"] = types.YLeaf{"OutgoingCallReplies", retransmit.OutgoingCallReplies}
    retransmit.EntityData.Leafs["outgoing-call-connected-packets"] = types.YLeaf{"OutgoingCallConnectedPackets", retransmit.OutgoingCallConnectedPackets}
    retransmit.EntityData.Leafs["incoming-call-requests"] = types.YLeaf{"IncomingCallRequests", retransmit.IncomingCallRequests}
    retransmit.EntityData.Leafs["incoming-call-replies"] = types.YLeaf{"IncomingCallReplies", retransmit.IncomingCallReplies}
    retransmit.EntityData.Leafs["incoming-call-connected-packets"] = types.YLeaf{"IncomingCallConnectedPackets", retransmit.IncomingCallConnectedPackets}
    retransmit.EntityData.Leafs["call-disconnect-notify-packets"] = types.YLeaf{"CallDisconnectNotifyPackets", retransmit.CallDisconnectNotifyPackets}
    retransmit.EntityData.Leafs["wan-error-notify-packets"] = types.YLeaf{"WanErrorNotifyPackets", retransmit.WanErrorNotifyPackets}
    retransmit.EntityData.Leafs["set-link-info-packets"] = types.YLeaf{"SetLinkInfoPackets", retransmit.SetLinkInfoPackets}
    retransmit.EntityData.Leafs["service-relay-requests"] = types.YLeaf{"ServiceRelayRequests", retransmit.ServiceRelayRequests}
    retransmit.EntityData.Leafs["service-relay-replies"] = types.YLeaf{"ServiceRelayReplies", retransmit.ServiceRelayReplies}
    retransmit.EntityData.Leafs["acknowledgement-packets"] = types.YLeaf{"AcknowledgementPackets", retransmit.AcknowledgementPackets}
    return &(retransmit.EntityData)
}

// L2Tpv2_Counters_Control_TunnelXr_Global_Received
// Received data
type L2Tpv2_Counters_Control_TunnelXr_Global_Received struct {
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

func (received *L2Tpv2_Counters_Control_TunnelXr_Global_Received) GetEntityData() *types.CommonEntityData {
    received.EntityData.YFilter = received.YFilter
    received.EntityData.YangName = "received"
    received.EntityData.BundleName = "cisco_ios_xr"
    received.EntityData.ParentYangName = "global"
    received.EntityData.SegmentPath = "received"
    received.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    received.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    received.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    received.EntityData.Children = make(map[string]types.YChild)
    received.EntityData.Leafs = make(map[string]types.YLeaf)
    received.EntityData.Leafs["unknown-packets"] = types.YLeaf{"UnknownPackets", received.UnknownPackets}
    received.EntityData.Leafs["zero-length-body-packets"] = types.YLeaf{"ZeroLengthBodyPackets", received.ZeroLengthBodyPackets}
    received.EntityData.Leafs["start-control-connection-requests"] = types.YLeaf{"StartControlConnectionRequests", received.StartControlConnectionRequests}
    received.EntityData.Leafs["start-control-connection-replies"] = types.YLeaf{"StartControlConnectionReplies", received.StartControlConnectionReplies}
    received.EntityData.Leafs["start-control-connection-notifications"] = types.YLeaf{"StartControlConnectionNotifications", received.StartControlConnectionNotifications}
    received.EntityData.Leafs["stop-control-connection-notifications"] = types.YLeaf{"StopControlConnectionNotifications", received.StopControlConnectionNotifications}
    received.EntityData.Leafs["hello-packets"] = types.YLeaf{"HelloPackets", received.HelloPackets}
    received.EntityData.Leafs["outgoing-call-requests"] = types.YLeaf{"OutgoingCallRequests", received.OutgoingCallRequests}
    received.EntityData.Leafs["outgoing-call-replies"] = types.YLeaf{"OutgoingCallReplies", received.OutgoingCallReplies}
    received.EntityData.Leafs["outgoing-call-connected-packets"] = types.YLeaf{"OutgoingCallConnectedPackets", received.OutgoingCallConnectedPackets}
    received.EntityData.Leafs["incoming-call-requests"] = types.YLeaf{"IncomingCallRequests", received.IncomingCallRequests}
    received.EntityData.Leafs["incoming-call-replies"] = types.YLeaf{"IncomingCallReplies", received.IncomingCallReplies}
    received.EntityData.Leafs["incoming-call-connected-packets"] = types.YLeaf{"IncomingCallConnectedPackets", received.IncomingCallConnectedPackets}
    received.EntityData.Leafs["call-disconnect-notify-packets"] = types.YLeaf{"CallDisconnectNotifyPackets", received.CallDisconnectNotifyPackets}
    received.EntityData.Leafs["wan-error-notify-packets"] = types.YLeaf{"WanErrorNotifyPackets", received.WanErrorNotifyPackets}
    received.EntityData.Leafs["set-link-info-packets"] = types.YLeaf{"SetLinkInfoPackets", received.SetLinkInfoPackets}
    received.EntityData.Leafs["service-relay-requests"] = types.YLeaf{"ServiceRelayRequests", received.ServiceRelayRequests}
    received.EntityData.Leafs["service-relay-replies"] = types.YLeaf{"ServiceRelayReplies", received.ServiceRelayReplies}
    received.EntityData.Leafs["acknowledgement-packets"] = types.YLeaf{"AcknowledgementPackets", received.AcknowledgementPackets}
    return &(received.EntityData)
}

// L2Tpv2_Counters_Control_TunnelXr_Global_Drop
// Drop data
type L2Tpv2_Counters_Control_TunnelXr_Global_Drop struct {
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

func (drop *L2Tpv2_Counters_Control_TunnelXr_Global_Drop) GetEntityData() *types.CommonEntityData {
    drop.EntityData.YFilter = drop.YFilter
    drop.EntityData.YangName = "drop"
    drop.EntityData.BundleName = "cisco_ios_xr"
    drop.EntityData.ParentYangName = "global"
    drop.EntityData.SegmentPath = "drop"
    drop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    drop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    drop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    drop.EntityData.Children = make(map[string]types.YChild)
    drop.EntityData.Leafs = make(map[string]types.YLeaf)
    drop.EntityData.Leafs["unknown-packets"] = types.YLeaf{"UnknownPackets", drop.UnknownPackets}
    drop.EntityData.Leafs["zero-length-body-packets"] = types.YLeaf{"ZeroLengthBodyPackets", drop.ZeroLengthBodyPackets}
    drop.EntityData.Leafs["start-control-connection-requests"] = types.YLeaf{"StartControlConnectionRequests", drop.StartControlConnectionRequests}
    drop.EntityData.Leafs["start-control-connection-replies"] = types.YLeaf{"StartControlConnectionReplies", drop.StartControlConnectionReplies}
    drop.EntityData.Leafs["start-control-connection-notifications"] = types.YLeaf{"StartControlConnectionNotifications", drop.StartControlConnectionNotifications}
    drop.EntityData.Leafs["stop-control-connection-notifications"] = types.YLeaf{"StopControlConnectionNotifications", drop.StopControlConnectionNotifications}
    drop.EntityData.Leafs["hello-packets"] = types.YLeaf{"HelloPackets", drop.HelloPackets}
    drop.EntityData.Leafs["outgoing-call-requests"] = types.YLeaf{"OutgoingCallRequests", drop.OutgoingCallRequests}
    drop.EntityData.Leafs["outgoing-call-replies"] = types.YLeaf{"OutgoingCallReplies", drop.OutgoingCallReplies}
    drop.EntityData.Leafs["outgoing-call-connected-packets"] = types.YLeaf{"OutgoingCallConnectedPackets", drop.OutgoingCallConnectedPackets}
    drop.EntityData.Leafs["incoming-call-requests"] = types.YLeaf{"IncomingCallRequests", drop.IncomingCallRequests}
    drop.EntityData.Leafs["incoming-call-replies"] = types.YLeaf{"IncomingCallReplies", drop.IncomingCallReplies}
    drop.EntityData.Leafs["incoming-call-connected-packets"] = types.YLeaf{"IncomingCallConnectedPackets", drop.IncomingCallConnectedPackets}
    drop.EntityData.Leafs["call-disconnect-notify-packets"] = types.YLeaf{"CallDisconnectNotifyPackets", drop.CallDisconnectNotifyPackets}
    drop.EntityData.Leafs["wan-error-notify-packets"] = types.YLeaf{"WanErrorNotifyPackets", drop.WanErrorNotifyPackets}
    drop.EntityData.Leafs["set-link-info-packets"] = types.YLeaf{"SetLinkInfoPackets", drop.SetLinkInfoPackets}
    drop.EntityData.Leafs["service-relay-requests"] = types.YLeaf{"ServiceRelayRequests", drop.ServiceRelayRequests}
    drop.EntityData.Leafs["service-relay-replies"] = types.YLeaf{"ServiceRelayReplies", drop.ServiceRelayReplies}
    drop.EntityData.Leafs["acknowledgement-packets"] = types.YLeaf{"AcknowledgementPackets", drop.AcknowledgementPackets}
    return &(drop.EntityData)
}

// L2Tpv2_Counters_Control_Tunnels
// Table of tunnel IDs of control message counters
type L2Tpv2_Counters_Control_Tunnels struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2TP tunnel control message counters. The type is slice of
    // L2Tpv2_Counters_Control_Tunnels_Tunnel.
    Tunnel []L2Tpv2_Counters_Control_Tunnels_Tunnel
}

func (tunnels *L2Tpv2_Counters_Control_Tunnels) GetEntityData() *types.CommonEntityData {
    tunnels.EntityData.YFilter = tunnels.YFilter
    tunnels.EntityData.YangName = "tunnels"
    tunnels.EntityData.BundleName = "cisco_ios_xr"
    tunnels.EntityData.ParentYangName = "control"
    tunnels.EntityData.SegmentPath = "tunnels"
    tunnels.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnels.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnels.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnels.EntityData.Children = make(map[string]types.YChild)
    tunnels.EntityData.Children["tunnel"] = types.YChild{"Tunnel", nil}
    for i := range tunnels.Tunnel {
        tunnels.EntityData.Children[types.GetSegmentPath(&tunnels.Tunnel[i])] = types.YChild{"Tunnel", &tunnels.Tunnel[i]}
    }
    tunnels.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(tunnels.EntityData)
}

// L2Tpv2_Counters_Control_Tunnels_Tunnel
// L2TP tunnel control message counters
type L2Tpv2_Counters_Control_Tunnels_Tunnel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. L2TP tunnel ID. The type is interface{} with
    // range: -2147483648..2147483647.
    TunnelId interface{}

    // L2TP control message local and remote addresses.
    Brief L2Tpv2_Counters_Control_Tunnels_Tunnel_Brief

    // Global data.
    Global L2Tpv2_Counters_Control_Tunnels_Tunnel_Global
}

func (tunnel *L2Tpv2_Counters_Control_Tunnels_Tunnel) GetEntityData() *types.CommonEntityData {
    tunnel.EntityData.YFilter = tunnel.YFilter
    tunnel.EntityData.YangName = "tunnel"
    tunnel.EntityData.BundleName = "cisco_ios_xr"
    tunnel.EntityData.ParentYangName = "tunnels"
    tunnel.EntityData.SegmentPath = "tunnel" + "[tunnel-id='" + fmt.Sprintf("%v", tunnel.TunnelId) + "']"
    tunnel.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnel.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnel.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnel.EntityData.Children = make(map[string]types.YChild)
    tunnel.EntityData.Children["brief"] = types.YChild{"Brief", &tunnel.Brief}
    tunnel.EntityData.Children["global"] = types.YChild{"Global", &tunnel.Global}
    tunnel.EntityData.Leafs = make(map[string]types.YLeaf)
    tunnel.EntityData.Leafs["tunnel-id"] = types.YLeaf{"TunnelId", tunnel.TunnelId}
    return &(tunnel.EntityData)
}

// L2Tpv2_Counters_Control_Tunnels_Tunnel_Brief
// L2TP control message local and remote addresses
type L2Tpv2_Counters_Control_Tunnels_Tunnel_Brief struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Remote tunnel ID. The type is interface{} with range: 0..4294967295.
    RemoteTunnelId interface{}

    // Local IP address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    LocalAddress interface{}

    // Remote IP address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RemoteAddress interface{}
}

func (brief *L2Tpv2_Counters_Control_Tunnels_Tunnel_Brief) GetEntityData() *types.CommonEntityData {
    brief.EntityData.YFilter = brief.YFilter
    brief.EntityData.YangName = "brief"
    brief.EntityData.BundleName = "cisco_ios_xr"
    brief.EntityData.ParentYangName = "tunnel"
    brief.EntityData.SegmentPath = "brief"
    brief.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    brief.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    brief.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    brief.EntityData.Children = make(map[string]types.YChild)
    brief.EntityData.Leafs = make(map[string]types.YLeaf)
    brief.EntityData.Leafs["remote-tunnel-id"] = types.YLeaf{"RemoteTunnelId", brief.RemoteTunnelId}
    brief.EntityData.Leafs["local-address"] = types.YLeaf{"LocalAddress", brief.LocalAddress}
    brief.EntityData.Leafs["remote-address"] = types.YLeaf{"RemoteAddress", brief.RemoteAddress}
    return &(brief.EntityData)
}

// L2Tpv2_Counters_Control_Tunnels_Tunnel_Global
// Global data
type L2Tpv2_Counters_Control_Tunnels_Tunnel_Global struct {
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
    Transmit L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Transmit

    // Re transmit data.
    Retransmit L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Retransmit

    // Received data.
    Received L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Received

    // Drop data.
    Drop L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Drop
}

func (global *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global) GetEntityData() *types.CommonEntityData {
    global.EntityData.YFilter = global.YFilter
    global.EntityData.YangName = "global"
    global.EntityData.BundleName = "cisco_ios_xr"
    global.EntityData.ParentYangName = "tunnel"
    global.EntityData.SegmentPath = "global"
    global.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    global.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    global.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    global.EntityData.Children = make(map[string]types.YChild)
    global.EntityData.Children["transmit"] = types.YChild{"Transmit", &global.Transmit}
    global.EntityData.Children["retransmit"] = types.YChild{"Retransmit", &global.Retransmit}
    global.EntityData.Children["received"] = types.YChild{"Received", &global.Received}
    global.EntityData.Children["drop"] = types.YChild{"Drop", &global.Drop}
    global.EntityData.Leafs = make(map[string]types.YLeaf)
    global.EntityData.Leafs["total-transmit"] = types.YLeaf{"TotalTransmit", global.TotalTransmit}
    global.EntityData.Leafs["total-retransmit"] = types.YLeaf{"TotalRetransmit", global.TotalRetransmit}
    global.EntityData.Leafs["total-received"] = types.YLeaf{"TotalReceived", global.TotalReceived}
    global.EntityData.Leafs["total-drop"] = types.YLeaf{"TotalDrop", global.TotalDrop}
    return &(global.EntityData)
}

// L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Transmit
// Transmit data
type L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Transmit struct {
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

func (transmit *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Transmit) GetEntityData() *types.CommonEntityData {
    transmit.EntityData.YFilter = transmit.YFilter
    transmit.EntityData.YangName = "transmit"
    transmit.EntityData.BundleName = "cisco_ios_xr"
    transmit.EntityData.ParentYangName = "global"
    transmit.EntityData.SegmentPath = "transmit"
    transmit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transmit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transmit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transmit.EntityData.Children = make(map[string]types.YChild)
    transmit.EntityData.Leafs = make(map[string]types.YLeaf)
    transmit.EntityData.Leafs["unknown-packets"] = types.YLeaf{"UnknownPackets", transmit.UnknownPackets}
    transmit.EntityData.Leafs["zero-length-body-packets"] = types.YLeaf{"ZeroLengthBodyPackets", transmit.ZeroLengthBodyPackets}
    transmit.EntityData.Leafs["start-control-connection-requests"] = types.YLeaf{"StartControlConnectionRequests", transmit.StartControlConnectionRequests}
    transmit.EntityData.Leafs["start-control-connection-replies"] = types.YLeaf{"StartControlConnectionReplies", transmit.StartControlConnectionReplies}
    transmit.EntityData.Leafs["start-control-connection-notifications"] = types.YLeaf{"StartControlConnectionNotifications", transmit.StartControlConnectionNotifications}
    transmit.EntityData.Leafs["stop-control-connection-notifications"] = types.YLeaf{"StopControlConnectionNotifications", transmit.StopControlConnectionNotifications}
    transmit.EntityData.Leafs["hello-packets"] = types.YLeaf{"HelloPackets", transmit.HelloPackets}
    transmit.EntityData.Leafs["outgoing-call-requests"] = types.YLeaf{"OutgoingCallRequests", transmit.OutgoingCallRequests}
    transmit.EntityData.Leafs["outgoing-call-replies"] = types.YLeaf{"OutgoingCallReplies", transmit.OutgoingCallReplies}
    transmit.EntityData.Leafs["outgoing-call-connected-packets"] = types.YLeaf{"OutgoingCallConnectedPackets", transmit.OutgoingCallConnectedPackets}
    transmit.EntityData.Leafs["incoming-call-requests"] = types.YLeaf{"IncomingCallRequests", transmit.IncomingCallRequests}
    transmit.EntityData.Leafs["incoming-call-replies"] = types.YLeaf{"IncomingCallReplies", transmit.IncomingCallReplies}
    transmit.EntityData.Leafs["incoming-call-connected-packets"] = types.YLeaf{"IncomingCallConnectedPackets", transmit.IncomingCallConnectedPackets}
    transmit.EntityData.Leafs["call-disconnect-notify-packets"] = types.YLeaf{"CallDisconnectNotifyPackets", transmit.CallDisconnectNotifyPackets}
    transmit.EntityData.Leafs["wan-error-notify-packets"] = types.YLeaf{"WanErrorNotifyPackets", transmit.WanErrorNotifyPackets}
    transmit.EntityData.Leafs["set-link-info-packets"] = types.YLeaf{"SetLinkInfoPackets", transmit.SetLinkInfoPackets}
    transmit.EntityData.Leafs["service-relay-requests"] = types.YLeaf{"ServiceRelayRequests", transmit.ServiceRelayRequests}
    transmit.EntityData.Leafs["service-relay-replies"] = types.YLeaf{"ServiceRelayReplies", transmit.ServiceRelayReplies}
    transmit.EntityData.Leafs["acknowledgement-packets"] = types.YLeaf{"AcknowledgementPackets", transmit.AcknowledgementPackets}
    return &(transmit.EntityData)
}

// L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Retransmit
// Re transmit data
type L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Retransmit struct {
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

func (retransmit *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Retransmit) GetEntityData() *types.CommonEntityData {
    retransmit.EntityData.YFilter = retransmit.YFilter
    retransmit.EntityData.YangName = "retransmit"
    retransmit.EntityData.BundleName = "cisco_ios_xr"
    retransmit.EntityData.ParentYangName = "global"
    retransmit.EntityData.SegmentPath = "retransmit"
    retransmit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    retransmit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    retransmit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    retransmit.EntityData.Children = make(map[string]types.YChild)
    retransmit.EntityData.Leafs = make(map[string]types.YLeaf)
    retransmit.EntityData.Leafs["unknown-packets"] = types.YLeaf{"UnknownPackets", retransmit.UnknownPackets}
    retransmit.EntityData.Leafs["zero-length-body-packets"] = types.YLeaf{"ZeroLengthBodyPackets", retransmit.ZeroLengthBodyPackets}
    retransmit.EntityData.Leafs["start-control-connection-requests"] = types.YLeaf{"StartControlConnectionRequests", retransmit.StartControlConnectionRequests}
    retransmit.EntityData.Leafs["start-control-connection-replies"] = types.YLeaf{"StartControlConnectionReplies", retransmit.StartControlConnectionReplies}
    retransmit.EntityData.Leafs["start-control-connection-notifications"] = types.YLeaf{"StartControlConnectionNotifications", retransmit.StartControlConnectionNotifications}
    retransmit.EntityData.Leafs["stop-control-connection-notifications"] = types.YLeaf{"StopControlConnectionNotifications", retransmit.StopControlConnectionNotifications}
    retransmit.EntityData.Leafs["hello-packets"] = types.YLeaf{"HelloPackets", retransmit.HelloPackets}
    retransmit.EntityData.Leafs["outgoing-call-requests"] = types.YLeaf{"OutgoingCallRequests", retransmit.OutgoingCallRequests}
    retransmit.EntityData.Leafs["outgoing-call-replies"] = types.YLeaf{"OutgoingCallReplies", retransmit.OutgoingCallReplies}
    retransmit.EntityData.Leafs["outgoing-call-connected-packets"] = types.YLeaf{"OutgoingCallConnectedPackets", retransmit.OutgoingCallConnectedPackets}
    retransmit.EntityData.Leafs["incoming-call-requests"] = types.YLeaf{"IncomingCallRequests", retransmit.IncomingCallRequests}
    retransmit.EntityData.Leafs["incoming-call-replies"] = types.YLeaf{"IncomingCallReplies", retransmit.IncomingCallReplies}
    retransmit.EntityData.Leafs["incoming-call-connected-packets"] = types.YLeaf{"IncomingCallConnectedPackets", retransmit.IncomingCallConnectedPackets}
    retransmit.EntityData.Leafs["call-disconnect-notify-packets"] = types.YLeaf{"CallDisconnectNotifyPackets", retransmit.CallDisconnectNotifyPackets}
    retransmit.EntityData.Leafs["wan-error-notify-packets"] = types.YLeaf{"WanErrorNotifyPackets", retransmit.WanErrorNotifyPackets}
    retransmit.EntityData.Leafs["set-link-info-packets"] = types.YLeaf{"SetLinkInfoPackets", retransmit.SetLinkInfoPackets}
    retransmit.EntityData.Leafs["service-relay-requests"] = types.YLeaf{"ServiceRelayRequests", retransmit.ServiceRelayRequests}
    retransmit.EntityData.Leafs["service-relay-replies"] = types.YLeaf{"ServiceRelayReplies", retransmit.ServiceRelayReplies}
    retransmit.EntityData.Leafs["acknowledgement-packets"] = types.YLeaf{"AcknowledgementPackets", retransmit.AcknowledgementPackets}
    return &(retransmit.EntityData)
}

// L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Received
// Received data
type L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Received struct {
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

func (received *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Received) GetEntityData() *types.CommonEntityData {
    received.EntityData.YFilter = received.YFilter
    received.EntityData.YangName = "received"
    received.EntityData.BundleName = "cisco_ios_xr"
    received.EntityData.ParentYangName = "global"
    received.EntityData.SegmentPath = "received"
    received.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    received.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    received.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    received.EntityData.Children = make(map[string]types.YChild)
    received.EntityData.Leafs = make(map[string]types.YLeaf)
    received.EntityData.Leafs["unknown-packets"] = types.YLeaf{"UnknownPackets", received.UnknownPackets}
    received.EntityData.Leafs["zero-length-body-packets"] = types.YLeaf{"ZeroLengthBodyPackets", received.ZeroLengthBodyPackets}
    received.EntityData.Leafs["start-control-connection-requests"] = types.YLeaf{"StartControlConnectionRequests", received.StartControlConnectionRequests}
    received.EntityData.Leafs["start-control-connection-replies"] = types.YLeaf{"StartControlConnectionReplies", received.StartControlConnectionReplies}
    received.EntityData.Leafs["start-control-connection-notifications"] = types.YLeaf{"StartControlConnectionNotifications", received.StartControlConnectionNotifications}
    received.EntityData.Leafs["stop-control-connection-notifications"] = types.YLeaf{"StopControlConnectionNotifications", received.StopControlConnectionNotifications}
    received.EntityData.Leafs["hello-packets"] = types.YLeaf{"HelloPackets", received.HelloPackets}
    received.EntityData.Leafs["outgoing-call-requests"] = types.YLeaf{"OutgoingCallRequests", received.OutgoingCallRequests}
    received.EntityData.Leafs["outgoing-call-replies"] = types.YLeaf{"OutgoingCallReplies", received.OutgoingCallReplies}
    received.EntityData.Leafs["outgoing-call-connected-packets"] = types.YLeaf{"OutgoingCallConnectedPackets", received.OutgoingCallConnectedPackets}
    received.EntityData.Leafs["incoming-call-requests"] = types.YLeaf{"IncomingCallRequests", received.IncomingCallRequests}
    received.EntityData.Leafs["incoming-call-replies"] = types.YLeaf{"IncomingCallReplies", received.IncomingCallReplies}
    received.EntityData.Leafs["incoming-call-connected-packets"] = types.YLeaf{"IncomingCallConnectedPackets", received.IncomingCallConnectedPackets}
    received.EntityData.Leafs["call-disconnect-notify-packets"] = types.YLeaf{"CallDisconnectNotifyPackets", received.CallDisconnectNotifyPackets}
    received.EntityData.Leafs["wan-error-notify-packets"] = types.YLeaf{"WanErrorNotifyPackets", received.WanErrorNotifyPackets}
    received.EntityData.Leafs["set-link-info-packets"] = types.YLeaf{"SetLinkInfoPackets", received.SetLinkInfoPackets}
    received.EntityData.Leafs["service-relay-requests"] = types.YLeaf{"ServiceRelayRequests", received.ServiceRelayRequests}
    received.EntityData.Leafs["service-relay-replies"] = types.YLeaf{"ServiceRelayReplies", received.ServiceRelayReplies}
    received.EntityData.Leafs["acknowledgement-packets"] = types.YLeaf{"AcknowledgementPackets", received.AcknowledgementPackets}
    return &(received.EntityData)
}

// L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Drop
// Drop data
type L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Drop struct {
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

func (drop *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Drop) GetEntityData() *types.CommonEntityData {
    drop.EntityData.YFilter = drop.YFilter
    drop.EntityData.YangName = "drop"
    drop.EntityData.BundleName = "cisco_ios_xr"
    drop.EntityData.ParentYangName = "global"
    drop.EntityData.SegmentPath = "drop"
    drop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    drop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    drop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    drop.EntityData.Children = make(map[string]types.YChild)
    drop.EntityData.Leafs = make(map[string]types.YLeaf)
    drop.EntityData.Leafs["unknown-packets"] = types.YLeaf{"UnknownPackets", drop.UnknownPackets}
    drop.EntityData.Leafs["zero-length-body-packets"] = types.YLeaf{"ZeroLengthBodyPackets", drop.ZeroLengthBodyPackets}
    drop.EntityData.Leafs["start-control-connection-requests"] = types.YLeaf{"StartControlConnectionRequests", drop.StartControlConnectionRequests}
    drop.EntityData.Leafs["start-control-connection-replies"] = types.YLeaf{"StartControlConnectionReplies", drop.StartControlConnectionReplies}
    drop.EntityData.Leafs["start-control-connection-notifications"] = types.YLeaf{"StartControlConnectionNotifications", drop.StartControlConnectionNotifications}
    drop.EntityData.Leafs["stop-control-connection-notifications"] = types.YLeaf{"StopControlConnectionNotifications", drop.StopControlConnectionNotifications}
    drop.EntityData.Leafs["hello-packets"] = types.YLeaf{"HelloPackets", drop.HelloPackets}
    drop.EntityData.Leafs["outgoing-call-requests"] = types.YLeaf{"OutgoingCallRequests", drop.OutgoingCallRequests}
    drop.EntityData.Leafs["outgoing-call-replies"] = types.YLeaf{"OutgoingCallReplies", drop.OutgoingCallReplies}
    drop.EntityData.Leafs["outgoing-call-connected-packets"] = types.YLeaf{"OutgoingCallConnectedPackets", drop.OutgoingCallConnectedPackets}
    drop.EntityData.Leafs["incoming-call-requests"] = types.YLeaf{"IncomingCallRequests", drop.IncomingCallRequests}
    drop.EntityData.Leafs["incoming-call-replies"] = types.YLeaf{"IncomingCallReplies", drop.IncomingCallReplies}
    drop.EntityData.Leafs["incoming-call-connected-packets"] = types.YLeaf{"IncomingCallConnectedPackets", drop.IncomingCallConnectedPackets}
    drop.EntityData.Leafs["call-disconnect-notify-packets"] = types.YLeaf{"CallDisconnectNotifyPackets", drop.CallDisconnectNotifyPackets}
    drop.EntityData.Leafs["wan-error-notify-packets"] = types.YLeaf{"WanErrorNotifyPackets", drop.WanErrorNotifyPackets}
    drop.EntityData.Leafs["set-link-info-packets"] = types.YLeaf{"SetLinkInfoPackets", drop.SetLinkInfoPackets}
    drop.EntityData.Leafs["service-relay-requests"] = types.YLeaf{"ServiceRelayRequests", drop.ServiceRelayRequests}
    drop.EntityData.Leafs["service-relay-replies"] = types.YLeaf{"ServiceRelayReplies", drop.ServiceRelayReplies}
    drop.EntityData.Leafs["acknowledgement-packets"] = types.YLeaf{"AcknowledgementPackets", drop.AcknowledgementPackets}
    return &(drop.EntityData)
}

// L2Tpv2_Statistics
// L2TP v2 statistics information
type L2Tpv2_Statistics struct {
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

func (statistics *L2Tpv2_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "l2tpv2"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = make(map[string]types.YChild)
    statistics.EntityData.Leafs = make(map[string]types.YLeaf)
    statistics.EntityData.Leafs["tunnels"] = types.YLeaf{"Tunnels", statistics.Tunnels}
    statistics.EntityData.Leafs["sessions"] = types.YLeaf{"Sessions", statistics.Sessions}
    statistics.EntityData.Leafs["sent-packets"] = types.YLeaf{"SentPackets", statistics.SentPackets}
    statistics.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", statistics.ReceivedPackets}
    statistics.EntityData.Leafs["average-packet-processing-time"] = types.YLeaf{"AveragePacketProcessingTime", statistics.AveragePacketProcessingTime}
    statistics.EntityData.Leafs["received-out-of-order-packets"] = types.YLeaf{"ReceivedOutOfOrderPackets", statistics.ReceivedOutOfOrderPackets}
    statistics.EntityData.Leafs["reorder-packets"] = types.YLeaf{"ReorderPackets", statistics.ReorderPackets}
    statistics.EntityData.Leafs["reorder-deviation-packets"] = types.YLeaf{"ReorderDeviationPackets", statistics.ReorderDeviationPackets}
    statistics.EntityData.Leafs["incoming-dropped-packets"] = types.YLeaf{"IncomingDroppedPackets", statistics.IncomingDroppedPackets}
    statistics.EntityData.Leafs["buffered-packets"] = types.YLeaf{"BufferedPackets", statistics.BufferedPackets}
    statistics.EntityData.Leafs["netio-packets"] = types.YLeaf{"NetioPackets", statistics.NetioPackets}
    return &(statistics.EntityData)
}

// L2Tpv2_Tunnel
// L2TPv2 tunnel 
type L2Tpv2_Tunnel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tunnel accounting counters.
    Accounting L2Tpv2_Tunnel_Accounting
}

func (tunnel *L2Tpv2_Tunnel) GetEntityData() *types.CommonEntityData {
    tunnel.EntityData.YFilter = tunnel.YFilter
    tunnel.EntityData.YangName = "tunnel"
    tunnel.EntityData.BundleName = "cisco_ios_xr"
    tunnel.EntityData.ParentYangName = "l2tpv2"
    tunnel.EntityData.SegmentPath = "tunnel"
    tunnel.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnel.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnel.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnel.EntityData.Children = make(map[string]types.YChild)
    tunnel.EntityData.Children["accounting"] = types.YChild{"Accounting", &tunnel.Accounting}
    tunnel.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(tunnel.EntityData)
}

// L2Tpv2_Tunnel_Accounting
// Tunnel accounting counters
type L2Tpv2_Tunnel_Accounting struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tunnel accounting statistics.
    Statistics L2Tpv2_Tunnel_Accounting_Statistics
}

func (accounting *L2Tpv2_Tunnel_Accounting) GetEntityData() *types.CommonEntityData {
    accounting.EntityData.YFilter = accounting.YFilter
    accounting.EntityData.YangName = "accounting"
    accounting.EntityData.BundleName = "cisco_ios_xr"
    accounting.EntityData.ParentYangName = "tunnel"
    accounting.EntityData.SegmentPath = "accounting"
    accounting.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accounting.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accounting.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accounting.EntityData.Children = make(map[string]types.YChild)
    accounting.EntityData.Children["statistics"] = types.YChild{"Statistics", &accounting.Statistics}
    accounting.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(accounting.EntityData)
}

// L2Tpv2_Tunnel_Accounting_Statistics
// Tunnel accounting statistics
type L2Tpv2_Tunnel_Accounting_Statistics struct {
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

func (statistics *L2Tpv2_Tunnel_Accounting_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "accounting"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = make(map[string]types.YChild)
    statistics.EntityData.Leafs = make(map[string]types.YLeaf)
    statistics.EntityData.Leafs["records-sent-successfully"] = types.YLeaf{"RecordsSentSuccessfully", statistics.RecordsSentSuccessfully}
    statistics.EntityData.Leafs["start"] = types.YLeaf{"Start", statistics.Start}
    statistics.EntityData.Leafs["stop"] = types.YLeaf{"Stop", statistics.Stop}
    statistics.EntityData.Leafs["reject"] = types.YLeaf{"Reject", statistics.Reject}
    statistics.EntityData.Leafs["transport-failures"] = types.YLeaf{"TransportFailures", statistics.TransportFailures}
    statistics.EntityData.Leafs["positive-acknowledgement"] = types.YLeaf{"PositiveAcknowledgement", statistics.PositiveAcknowledgement}
    statistics.EntityData.Leafs["negative-acknowledgement"] = types.YLeaf{"NegativeAcknowledgement", statistics.NegativeAcknowledgement}
    statistics.EntityData.Leafs["records-checkpointed"] = types.YLeaf{"RecordsCheckpointed", statistics.RecordsCheckpointed}
    statistics.EntityData.Leafs["records-failed-to-checkpoint"] = types.YLeaf{"RecordsFailedToCheckpoint", statistics.RecordsFailedToCheckpoint}
    statistics.EntityData.Leafs["records-sent-from-queue"] = types.YLeaf{"RecordsSentFromQueue", statistics.RecordsSentFromQueue}
    statistics.EntityData.Leafs["memory-failures"] = types.YLeaf{"MemoryFailures", statistics.MemoryFailures}
    statistics.EntityData.Leafs["current-size"] = types.YLeaf{"CurrentSize", statistics.CurrentSize}
    statistics.EntityData.Leafs["records-recovered-from-checkpoint"] = types.YLeaf{"RecordsRecoveredFromCheckpoint", statistics.RecordsRecoveredFromCheckpoint}
    statistics.EntityData.Leafs["records-fail-to-recover"] = types.YLeaf{"RecordsFailToRecover", statistics.RecordsFailToRecover}
    statistics.EntityData.Leafs["queue-statistics-size"] = types.YLeaf{"QueueStatisticsSize", statistics.QueueStatisticsSize}
    return &(statistics.EntityData)
}

// L2Tpv2_TunnelConfigurations
// List of tunnel IDs
type L2Tpv2_TunnelConfigurations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2TP tunnel information. The type is slice of
    // L2Tpv2_TunnelConfigurations_TunnelConfiguration.
    TunnelConfiguration []L2Tpv2_TunnelConfigurations_TunnelConfiguration
}

func (tunnelConfigurations *L2Tpv2_TunnelConfigurations) GetEntityData() *types.CommonEntityData {
    tunnelConfigurations.EntityData.YFilter = tunnelConfigurations.YFilter
    tunnelConfigurations.EntityData.YangName = "tunnel-configurations"
    tunnelConfigurations.EntityData.BundleName = "cisco_ios_xr"
    tunnelConfigurations.EntityData.ParentYangName = "l2tpv2"
    tunnelConfigurations.EntityData.SegmentPath = "tunnel-configurations"
    tunnelConfigurations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnelConfigurations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnelConfigurations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnelConfigurations.EntityData.Children = make(map[string]types.YChild)
    tunnelConfigurations.EntityData.Children["tunnel-configuration"] = types.YChild{"TunnelConfiguration", nil}
    for i := range tunnelConfigurations.TunnelConfiguration {
        tunnelConfigurations.EntityData.Children[types.GetSegmentPath(&tunnelConfigurations.TunnelConfiguration[i])] = types.YChild{"TunnelConfiguration", &tunnelConfigurations.TunnelConfiguration[i]}
    }
    tunnelConfigurations.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(tunnelConfigurations.EntityData)
}

// L2Tpv2_TunnelConfigurations_TunnelConfiguration
// L2TP tunnel information
type L2Tpv2_TunnelConfigurations_TunnelConfiguration struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Local tunnel ID. The type is interface{} with
    // range: -2147483648..2147483647.
    LocalTunnelId interface{}

    // Remote tunnel ID. The type is interface{} with range: 0..4294967295.
    RemoteTunnelId interface{}

    // L2Tp class data.
    L2TpClass L2Tpv2_TunnelConfigurations_TunnelConfiguration_L2TpClass
}

func (tunnelConfiguration *L2Tpv2_TunnelConfigurations_TunnelConfiguration) GetEntityData() *types.CommonEntityData {
    tunnelConfiguration.EntityData.YFilter = tunnelConfiguration.YFilter
    tunnelConfiguration.EntityData.YangName = "tunnel-configuration"
    tunnelConfiguration.EntityData.BundleName = "cisco_ios_xr"
    tunnelConfiguration.EntityData.ParentYangName = "tunnel-configurations"
    tunnelConfiguration.EntityData.SegmentPath = "tunnel-configuration" + "[local-tunnel-id='" + fmt.Sprintf("%v", tunnelConfiguration.LocalTunnelId) + "']"
    tunnelConfiguration.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnelConfiguration.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnelConfiguration.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnelConfiguration.EntityData.Children = make(map[string]types.YChild)
    tunnelConfiguration.EntityData.Children["l2tp-class"] = types.YChild{"L2TpClass", &tunnelConfiguration.L2TpClass}
    tunnelConfiguration.EntityData.Leafs = make(map[string]types.YLeaf)
    tunnelConfiguration.EntityData.Leafs["local-tunnel-id"] = types.YLeaf{"LocalTunnelId", tunnelConfiguration.LocalTunnelId}
    tunnelConfiguration.EntityData.Leafs["remote-tunnel-id"] = types.YLeaf{"RemoteTunnelId", tunnelConfiguration.RemoteTunnelId}
    return &(tunnelConfiguration.EntityData)
}

// L2Tpv2_TunnelConfigurations_TunnelConfiguration_L2TpClass
// L2Tp class data
type L2Tpv2_TunnelConfigurations_TunnelConfiguration_L2TpClass struct {
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

func (l2TpClass *L2Tpv2_TunnelConfigurations_TunnelConfiguration_L2TpClass) GetEntityData() *types.CommonEntityData {
    l2TpClass.EntityData.YFilter = l2TpClass.YFilter
    l2TpClass.EntityData.YangName = "l2tp-class"
    l2TpClass.EntityData.BundleName = "cisco_ios_xr"
    l2TpClass.EntityData.ParentYangName = "tunnel-configuration"
    l2TpClass.EntityData.SegmentPath = "l2tp-class"
    l2TpClass.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    l2TpClass.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    l2TpClass.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    l2TpClass.EntityData.Children = make(map[string]types.YChild)
    l2TpClass.EntityData.Leafs = make(map[string]types.YLeaf)
    l2TpClass.EntityData.Leafs["ip-tos"] = types.YLeaf{"IpTos", l2TpClass.IpTos}
    l2TpClass.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", l2TpClass.VrfName}
    l2TpClass.EntityData.Leafs["receive-window-size"] = types.YLeaf{"ReceiveWindowSize", l2TpClass.ReceiveWindowSize}
    l2TpClass.EntityData.Leafs["class-name-xr"] = types.YLeaf{"ClassNameXr", l2TpClass.ClassNameXr}
    l2TpClass.EntityData.Leafs["digest-hash"] = types.YLeaf{"DigestHash", l2TpClass.DigestHash}
    l2TpClass.EntityData.Leafs["password"] = types.YLeaf{"Password", l2TpClass.Password}
    l2TpClass.EntityData.Leafs["encoded-password"] = types.YLeaf{"EncodedPassword", l2TpClass.EncodedPassword}
    l2TpClass.EntityData.Leafs["host-name"] = types.YLeaf{"HostName", l2TpClass.HostName}
    l2TpClass.EntityData.Leafs["accounting-method-list"] = types.YLeaf{"AccountingMethodList", l2TpClass.AccountingMethodList}
    l2TpClass.EntityData.Leafs["hello-timeout"] = types.YLeaf{"HelloTimeout", l2TpClass.HelloTimeout}
    l2TpClass.EntityData.Leafs["setup-timeout"] = types.YLeaf{"SetupTimeout", l2TpClass.SetupTimeout}
    l2TpClass.EntityData.Leafs["retransmit-minimum-timeout"] = types.YLeaf{"RetransmitMinimumTimeout", l2TpClass.RetransmitMinimumTimeout}
    l2TpClass.EntityData.Leafs["retransmit-maximum-timeout"] = types.YLeaf{"RetransmitMaximumTimeout", l2TpClass.RetransmitMaximumTimeout}
    l2TpClass.EntityData.Leafs["initial-retransmit-minimum-timeout"] = types.YLeaf{"InitialRetransmitMinimumTimeout", l2TpClass.InitialRetransmitMinimumTimeout}
    l2TpClass.EntityData.Leafs["initial-retransmit-maximum-timeout"] = types.YLeaf{"InitialRetransmitMaximumTimeout", l2TpClass.InitialRetransmitMaximumTimeout}
    l2TpClass.EntityData.Leafs["timeout-no-user"] = types.YLeaf{"TimeoutNoUser", l2TpClass.TimeoutNoUser}
    l2TpClass.EntityData.Leafs["retransmit-retries"] = types.YLeaf{"RetransmitRetries", l2TpClass.RetransmitRetries}
    l2TpClass.EntityData.Leafs["initial-retransmit-retries"] = types.YLeaf{"InitialRetransmitRetries", l2TpClass.InitialRetransmitRetries}
    l2TpClass.EntityData.Leafs["is-authentication-enabled"] = types.YLeaf{"IsAuthenticationEnabled", l2TpClass.IsAuthenticationEnabled}
    l2TpClass.EntityData.Leafs["is-hidden"] = types.YLeaf{"IsHidden", l2TpClass.IsHidden}
    l2TpClass.EntityData.Leafs["is-digest-enabled"] = types.YLeaf{"IsDigestEnabled", l2TpClass.IsDigestEnabled}
    l2TpClass.EntityData.Leafs["is-digest-check-enabled"] = types.YLeaf{"IsDigestCheckEnabled", l2TpClass.IsDigestCheckEnabled}
    l2TpClass.EntityData.Leafs["is-congestion-control-enabled"] = types.YLeaf{"IsCongestionControlEnabled", l2TpClass.IsCongestionControlEnabled}
    l2TpClass.EntityData.Leafs["is-peer-address-checked"] = types.YLeaf{"IsPeerAddressChecked", l2TpClass.IsPeerAddressChecked}
    return &(l2TpClass.EntityData)
}

// L2Tpv2_CounterHistFail
// Failure events leading to disconnection
type L2Tpv2_CounterHistFail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // sesions affected due to timeout. The type is interface{} with range:
    // 0..4294967295.
    SessDownTmout interface{}

    // Send side counters. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    TxCounters interface{}

    // Receive side counters. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    RxCounters interface{}

    // timeout events by packet. The type is slice of interface{} with range:
    // 0..4294967295.
    PktTimeout []interface{}
}

func (counterHistFail *L2Tpv2_CounterHistFail) GetEntityData() *types.CommonEntityData {
    counterHistFail.EntityData.YFilter = counterHistFail.YFilter
    counterHistFail.EntityData.YangName = "counter-hist-fail"
    counterHistFail.EntityData.BundleName = "cisco_ios_xr"
    counterHistFail.EntityData.ParentYangName = "l2tpv2"
    counterHistFail.EntityData.SegmentPath = "counter-hist-fail"
    counterHistFail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    counterHistFail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    counterHistFail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    counterHistFail.EntityData.Children = make(map[string]types.YChild)
    counterHistFail.EntityData.Leafs = make(map[string]types.YLeaf)
    counterHistFail.EntityData.Leafs["sess-down-tmout"] = types.YLeaf{"SessDownTmout", counterHistFail.SessDownTmout}
    counterHistFail.EntityData.Leafs["tx-counters"] = types.YLeaf{"TxCounters", counterHistFail.TxCounters}
    counterHistFail.EntityData.Leafs["rx-counters"] = types.YLeaf{"RxCounters", counterHistFail.RxCounters}
    counterHistFail.EntityData.Leafs["pkt-timeout"] = types.YLeaf{"PktTimeout", counterHistFail.PktTimeout}
    return &(counterHistFail.EntityData)
}

// L2Tpv2_Classes
// List of L2TP class names
type L2Tpv2_Classes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2TP class name. The type is slice of L2Tpv2_Classes_Class.
    Class []L2Tpv2_Classes_Class
}

func (classes *L2Tpv2_Classes) GetEntityData() *types.CommonEntityData {
    classes.EntityData.YFilter = classes.YFilter
    classes.EntityData.YangName = "classes"
    classes.EntityData.BundleName = "cisco_ios_xr"
    classes.EntityData.ParentYangName = "l2tpv2"
    classes.EntityData.SegmentPath = "classes"
    classes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    classes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    classes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    classes.EntityData.Children = make(map[string]types.YChild)
    classes.EntityData.Children["class"] = types.YChild{"Class", nil}
    for i := range classes.Class {
        classes.EntityData.Children[types.GetSegmentPath(&classes.Class[i])] = types.YChild{"Class", &classes.Class[i]}
    }
    classes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(classes.EntityData)
}

// L2Tpv2_Classes_Class
// L2TP class name
type L2Tpv2_Classes_Class struct {
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

func (class *L2Tpv2_Classes_Class) GetEntityData() *types.CommonEntityData {
    class.EntityData.YFilter = class.YFilter
    class.EntityData.YangName = "class"
    class.EntityData.BundleName = "cisco_ios_xr"
    class.EntityData.ParentYangName = "classes"
    class.EntityData.SegmentPath = "class" + "[class-name='" + fmt.Sprintf("%v", class.ClassName) + "']"
    class.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    class.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    class.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    class.EntityData.Children = make(map[string]types.YChild)
    class.EntityData.Leafs = make(map[string]types.YLeaf)
    class.EntityData.Leafs["class-name"] = types.YLeaf{"ClassName", class.ClassName}
    class.EntityData.Leafs["ip-tos"] = types.YLeaf{"IpTos", class.IpTos}
    class.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", class.VrfName}
    class.EntityData.Leafs["receive-window-size"] = types.YLeaf{"ReceiveWindowSize", class.ReceiveWindowSize}
    class.EntityData.Leafs["class-name-xr"] = types.YLeaf{"ClassNameXr", class.ClassNameXr}
    class.EntityData.Leafs["digest-hash"] = types.YLeaf{"DigestHash", class.DigestHash}
    class.EntityData.Leafs["password"] = types.YLeaf{"Password", class.Password}
    class.EntityData.Leafs["encoded-password"] = types.YLeaf{"EncodedPassword", class.EncodedPassword}
    class.EntityData.Leafs["host-name"] = types.YLeaf{"HostName", class.HostName}
    class.EntityData.Leafs["accounting-method-list"] = types.YLeaf{"AccountingMethodList", class.AccountingMethodList}
    class.EntityData.Leafs["hello-timeout"] = types.YLeaf{"HelloTimeout", class.HelloTimeout}
    class.EntityData.Leafs["setup-timeout"] = types.YLeaf{"SetupTimeout", class.SetupTimeout}
    class.EntityData.Leafs["retransmit-minimum-timeout"] = types.YLeaf{"RetransmitMinimumTimeout", class.RetransmitMinimumTimeout}
    class.EntityData.Leafs["retransmit-maximum-timeout"] = types.YLeaf{"RetransmitMaximumTimeout", class.RetransmitMaximumTimeout}
    class.EntityData.Leafs["initial-retransmit-minimum-timeout"] = types.YLeaf{"InitialRetransmitMinimumTimeout", class.InitialRetransmitMinimumTimeout}
    class.EntityData.Leafs["initial-retransmit-maximum-timeout"] = types.YLeaf{"InitialRetransmitMaximumTimeout", class.InitialRetransmitMaximumTimeout}
    class.EntityData.Leafs["timeout-no-user"] = types.YLeaf{"TimeoutNoUser", class.TimeoutNoUser}
    class.EntityData.Leafs["retransmit-retries"] = types.YLeaf{"RetransmitRetries", class.RetransmitRetries}
    class.EntityData.Leafs["initial-retransmit-retries"] = types.YLeaf{"InitialRetransmitRetries", class.InitialRetransmitRetries}
    class.EntityData.Leafs["is-authentication-enabled"] = types.YLeaf{"IsAuthenticationEnabled", class.IsAuthenticationEnabled}
    class.EntityData.Leafs["is-hidden"] = types.YLeaf{"IsHidden", class.IsHidden}
    class.EntityData.Leafs["is-digest-enabled"] = types.YLeaf{"IsDigestEnabled", class.IsDigestEnabled}
    class.EntityData.Leafs["is-digest-check-enabled"] = types.YLeaf{"IsDigestCheckEnabled", class.IsDigestCheckEnabled}
    class.EntityData.Leafs["is-congestion-control-enabled"] = types.YLeaf{"IsCongestionControlEnabled", class.IsCongestionControlEnabled}
    class.EntityData.Leafs["is-peer-address-checked"] = types.YLeaf{"IsPeerAddressChecked", class.IsPeerAddressChecked}
    return &(class.EntityData)
}

// L2Tpv2_Tunnels
// List of tunnel IDs
type L2Tpv2_Tunnels struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2TP tunnel  information. The type is slice of L2Tpv2_Tunnels_Tunnel.
    Tunnel []L2Tpv2_Tunnels_Tunnel
}

func (tunnels *L2Tpv2_Tunnels) GetEntityData() *types.CommonEntityData {
    tunnels.EntityData.YFilter = tunnels.YFilter
    tunnels.EntityData.YangName = "tunnels"
    tunnels.EntityData.BundleName = "cisco_ios_xr"
    tunnels.EntityData.ParentYangName = "l2tpv2"
    tunnels.EntityData.SegmentPath = "tunnels"
    tunnels.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnels.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnels.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnels.EntityData.Children = make(map[string]types.YChild)
    tunnels.EntityData.Children["tunnel"] = types.YChild{"Tunnel", nil}
    for i := range tunnels.Tunnel {
        tunnels.EntityData.Children[types.GetSegmentPath(&tunnels.Tunnel[i])] = types.YChild{"Tunnel", &tunnels.Tunnel[i]}
    }
    tunnels.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(tunnels.EntityData)
}

// L2Tpv2_Tunnels_Tunnel
// L2TP tunnel  information
type L2Tpv2_Tunnels_Tunnel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Local tunnel ID. The type is interface{} with
    // range: -2147483648..2147483647.
    LocalTunnelId interface{}

    // Local tunnel address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    LocalAddress interface{}

    // Remote tunnel address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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

    // Retransmit time distribution in seconds. The type is slice of interface{}
    // with range: 0..65535. Units are second.
    RetransmitTime []interface{}
}

func (tunnel *L2Tpv2_Tunnels_Tunnel) GetEntityData() *types.CommonEntityData {
    tunnel.EntityData.YFilter = tunnel.YFilter
    tunnel.EntityData.YangName = "tunnel"
    tunnel.EntityData.BundleName = "cisco_ios_xr"
    tunnel.EntityData.ParentYangName = "tunnels"
    tunnel.EntityData.SegmentPath = "tunnel" + "[local-tunnel-id='" + fmt.Sprintf("%v", tunnel.LocalTunnelId) + "']"
    tunnel.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnel.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnel.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnel.EntityData.Children = make(map[string]types.YChild)
    tunnel.EntityData.Leafs = make(map[string]types.YLeaf)
    tunnel.EntityData.Leafs["local-tunnel-id"] = types.YLeaf{"LocalTunnelId", tunnel.LocalTunnelId}
    tunnel.EntityData.Leafs["local-address"] = types.YLeaf{"LocalAddress", tunnel.LocalAddress}
    tunnel.EntityData.Leafs["remote-address"] = types.YLeaf{"RemoteAddress", tunnel.RemoteAddress}
    tunnel.EntityData.Leafs["local-port"] = types.YLeaf{"LocalPort", tunnel.LocalPort}
    tunnel.EntityData.Leafs["remote-port"] = types.YLeaf{"RemotePort", tunnel.RemotePort}
    tunnel.EntityData.Leafs["protocol"] = types.YLeaf{"Protocol", tunnel.Protocol}
    tunnel.EntityData.Leafs["is-pmtu-enabled"] = types.YLeaf{"IsPmtuEnabled", tunnel.IsPmtuEnabled}
    tunnel.EntityData.Leafs["remote-tunnel-id"] = types.YLeaf{"RemoteTunnelId", tunnel.RemoteTunnelId}
    tunnel.EntityData.Leafs["local-tunnel-name"] = types.YLeaf{"LocalTunnelName", tunnel.LocalTunnelName}
    tunnel.EntityData.Leafs["remote-tunnel-name"] = types.YLeaf{"RemoteTunnelName", tunnel.RemoteTunnelName}
    tunnel.EntityData.Leafs["class-name"] = types.YLeaf{"ClassName", tunnel.ClassName}
    tunnel.EntityData.Leafs["active-sessions"] = types.YLeaf{"ActiveSessions", tunnel.ActiveSessions}
    tunnel.EntityData.Leafs["sequence-ns"] = types.YLeaf{"SequenceNs", tunnel.SequenceNs}
    tunnel.EntityData.Leafs["sequence-nr"] = types.YLeaf{"SequenceNr", tunnel.SequenceNr}
    tunnel.EntityData.Leafs["local-window-size"] = types.YLeaf{"LocalWindowSize", tunnel.LocalWindowSize}
    tunnel.EntityData.Leafs["remote-window-size"] = types.YLeaf{"RemoteWindowSize", tunnel.RemoteWindowSize}
    tunnel.EntityData.Leafs["retransmission-time"] = types.YLeaf{"RetransmissionTime", tunnel.RetransmissionTime}
    tunnel.EntityData.Leafs["maximum-retransmission-time"] = types.YLeaf{"MaximumRetransmissionTime", tunnel.MaximumRetransmissionTime}
    tunnel.EntityData.Leafs["unsent-queue-size"] = types.YLeaf{"UnsentQueueSize", tunnel.UnsentQueueSize}
    tunnel.EntityData.Leafs["unsent-maximum-queue-size"] = types.YLeaf{"UnsentMaximumQueueSize", tunnel.UnsentMaximumQueueSize}
    tunnel.EntityData.Leafs["resend-queue-size"] = types.YLeaf{"ResendQueueSize", tunnel.ResendQueueSize}
    tunnel.EntityData.Leafs["resend-maximum-queue-size"] = types.YLeaf{"ResendMaximumQueueSize", tunnel.ResendMaximumQueueSize}
    tunnel.EntityData.Leafs["order-queue-size"] = types.YLeaf{"OrderQueueSize", tunnel.OrderQueueSize}
    tunnel.EntityData.Leafs["packet-queue-check"] = types.YLeaf{"PacketQueueCheck", tunnel.PacketQueueCheck}
    tunnel.EntityData.Leafs["digest-secrets"] = types.YLeaf{"DigestSecrets", tunnel.DigestSecrets}
    tunnel.EntityData.Leafs["resends"] = types.YLeaf{"Resends", tunnel.Resends}
    tunnel.EntityData.Leafs["zero-length-body-acknowledgement-sent"] = types.YLeaf{"ZeroLengthBodyAcknowledgementSent", tunnel.ZeroLengthBodyAcknowledgementSent}
    tunnel.EntityData.Leafs["total-out-of-order-drop-packets"] = types.YLeaf{"TotalOutOfOrderDropPackets", tunnel.TotalOutOfOrderDropPackets}
    tunnel.EntityData.Leafs["total-out-of-order-reorder-packets"] = types.YLeaf{"TotalOutOfOrderReorderPackets", tunnel.TotalOutOfOrderReorderPackets}
    tunnel.EntityData.Leafs["total-peer-authentication-failures"] = types.YLeaf{"TotalPeerAuthenticationFailures", tunnel.TotalPeerAuthenticationFailures}
    tunnel.EntityData.Leafs["is-tunnel-up"] = types.YLeaf{"IsTunnelUp", tunnel.IsTunnelUp}
    tunnel.EntityData.Leafs["is-congestion-control-enabled"] = types.YLeaf{"IsCongestionControlEnabled", tunnel.IsCongestionControlEnabled}
    tunnel.EntityData.Leafs["retransmit-time"] = types.YLeaf{"RetransmitTime", tunnel.RetransmitTime}
    return &(tunnel.EntityData)
}

// L2Tpv2_Sessions
// List of session IDs
type L2Tpv2_Sessions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2TP information for a particular session. The type is slice of
    // L2Tpv2_Sessions_Session.
    Session []L2Tpv2_Sessions_Session
}

func (sessions *L2Tpv2_Sessions) GetEntityData() *types.CommonEntityData {
    sessions.EntityData.YFilter = sessions.YFilter
    sessions.EntityData.YangName = "sessions"
    sessions.EntityData.BundleName = "cisco_ios_xr"
    sessions.EntityData.ParentYangName = "l2tpv2"
    sessions.EntityData.SegmentPath = "sessions"
    sessions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessions.EntityData.Children = make(map[string]types.YChild)
    sessions.EntityData.Children["session"] = types.YChild{"Session", nil}
    for i := range sessions.Session {
        sessions.EntityData.Children[types.GetSegmentPath(&sessions.Session[i])] = types.YChild{"Session", &sessions.Session[i]}
    }
    sessions.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sessions.EntityData)
}

// L2Tpv2_Sessions_Session
// L2TP information for a particular session
type L2Tpv2_Sessions_Session struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Local tunnel ID. The type is interface{} with
    // range: -2147483648..2147483647.
    LocalTunnelId interface{}

    // This attribute is a key. Local session ID. The type is interface{} with
    // range: -2147483648..2147483647.
    LocalSessionId interface{}

    // Local session IP address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    LocalIpAddress interface{}

    // Remote session IP address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RemoteIpAddress interface{}

    // l2tp sh sess udp lport. The type is interface{} with range: 0..65535.
    L2TpShSessUdpLport interface{}

    // l2tp sh sess udp rport. The type is interface{} with range: 0..65535.
    L2TpShSessUdpRport interface{}

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
    L2TpShSessTieBreakerEnabled interface{}

    // l2tp sh sess tie breaker. The type is interface{} with range:
    // 0..18446744073709551615.
    L2TpShSessTieBreaker interface{}

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
    SessionApplicationData L2Tpv2_Sessions_Session_SessionApplicationData
}

func (session *L2Tpv2_Sessions_Session) GetEntityData() *types.CommonEntityData {
    session.EntityData.YFilter = session.YFilter
    session.EntityData.YangName = "session"
    session.EntityData.BundleName = "cisco_ios_xr"
    session.EntityData.ParentYangName = "sessions"
    session.EntityData.SegmentPath = "session" + "[local-tunnel-id='" + fmt.Sprintf("%v", session.LocalTunnelId) + "']" + "[local-session-id='" + fmt.Sprintf("%v", session.LocalSessionId) + "']"
    session.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    session.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    session.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    session.EntityData.Children = make(map[string]types.YChild)
    session.EntityData.Children["session-application-data"] = types.YChild{"SessionApplicationData", &session.SessionApplicationData}
    session.EntityData.Leafs = make(map[string]types.YLeaf)
    session.EntityData.Leafs["local-tunnel-id"] = types.YLeaf{"LocalTunnelId", session.LocalTunnelId}
    session.EntityData.Leafs["local-session-id"] = types.YLeaf{"LocalSessionId", session.LocalSessionId}
    session.EntityData.Leafs["local-ip-address"] = types.YLeaf{"LocalIpAddress", session.LocalIpAddress}
    session.EntityData.Leafs["remote-ip-address"] = types.YLeaf{"RemoteIpAddress", session.RemoteIpAddress}
    session.EntityData.Leafs["l2tp-sh-sess-udp-lport"] = types.YLeaf{"L2TpShSessUdpLport", session.L2TpShSessUdpLport}
    session.EntityData.Leafs["l2tp-sh-sess-udp-rport"] = types.YLeaf{"L2TpShSessUdpRport", session.L2TpShSessUdpRport}
    session.EntityData.Leafs["protocol"] = types.YLeaf{"Protocol", session.Protocol}
    session.EntityData.Leafs["remote-tunnel-id"] = types.YLeaf{"RemoteTunnelId", session.RemoteTunnelId}
    session.EntityData.Leafs["call-serial-number"] = types.YLeaf{"CallSerialNumber", session.CallSerialNumber}
    session.EntityData.Leafs["local-tunnel-name"] = types.YLeaf{"LocalTunnelName", session.LocalTunnelName}
    session.EntityData.Leafs["remote-tunnel-name"] = types.YLeaf{"RemoteTunnelName", session.RemoteTunnelName}
    session.EntityData.Leafs["remote-session-id"] = types.YLeaf{"RemoteSessionId", session.RemoteSessionId}
    session.EntityData.Leafs["l2tp-sh-sess-tie-breaker-enabled"] = types.YLeaf{"L2TpShSessTieBreakerEnabled", session.L2TpShSessTieBreakerEnabled}
    session.EntityData.Leafs["l2tp-sh-sess-tie-breaker"] = types.YLeaf{"L2TpShSessTieBreaker", session.L2TpShSessTieBreaker}
    session.EntityData.Leafs["is-session-manual"] = types.YLeaf{"IsSessionManual", session.IsSessionManual}
    session.EntityData.Leafs["is-session-up"] = types.YLeaf{"IsSessionUp", session.IsSessionUp}
    session.EntityData.Leafs["is-udp-checksum-enabled"] = types.YLeaf{"IsUdpChecksumEnabled", session.IsUdpChecksumEnabled}
    session.EntityData.Leafs["is-sequencing-on"] = types.YLeaf{"IsSequencingOn", session.IsSequencingOn}
    session.EntityData.Leafs["is-session-state-established"] = types.YLeaf{"IsSessionStateEstablished", session.IsSessionStateEstablished}
    session.EntityData.Leafs["is-session-locally-initiated"] = types.YLeaf{"IsSessionLocallyInitiated", session.IsSessionLocallyInitiated}
    session.EntityData.Leafs["is-conditional-debug-enabled"] = types.YLeaf{"IsConditionalDebugEnabled", session.IsConditionalDebugEnabled}
    session.EntityData.Leafs["unique-id"] = types.YLeaf{"UniqueId", session.UniqueId}
    session.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", session.InterfaceName}
    return &(session.EntityData)
}

// L2Tpv2_Sessions_Session_SessionApplicationData
// Session application data
type L2Tpv2_Sessions_Session_SessionApplicationData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // l2tp sh sess app type. The type is interface{} with range: 0..4294967295.
    L2TpShSessAppType interface{}

    // Xconnect data.
    Xconnect L2Tpv2_Sessions_Session_SessionApplicationData_Xconnect

    // VPDN data.
    Vpdn L2Tpv2_Sessions_Session_SessionApplicationData_Vpdn
}

func (sessionApplicationData *L2Tpv2_Sessions_Session_SessionApplicationData) GetEntityData() *types.CommonEntityData {
    sessionApplicationData.EntityData.YFilter = sessionApplicationData.YFilter
    sessionApplicationData.EntityData.YangName = "session-application-data"
    sessionApplicationData.EntityData.BundleName = "cisco_ios_xr"
    sessionApplicationData.EntityData.ParentYangName = "session"
    sessionApplicationData.EntityData.SegmentPath = "session-application-data"
    sessionApplicationData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionApplicationData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionApplicationData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionApplicationData.EntityData.Children = make(map[string]types.YChild)
    sessionApplicationData.EntityData.Children["xconnect"] = types.YChild{"Xconnect", &sessionApplicationData.Xconnect}
    sessionApplicationData.EntityData.Children["vpdn"] = types.YChild{"Vpdn", &sessionApplicationData.Vpdn}
    sessionApplicationData.EntityData.Leafs = make(map[string]types.YLeaf)
    sessionApplicationData.EntityData.Leafs["l2tp-sh-sess-app-type"] = types.YLeaf{"L2TpShSessAppType", sessionApplicationData.L2TpShSessAppType}
    return &(sessionApplicationData.EntityData)
}

// L2Tpv2_Sessions_Session_SessionApplicationData_Xconnect
// Xconnect data
type L2Tpv2_Sessions_Session_SessionApplicationData_Xconnect struct {
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

func (xconnect *L2Tpv2_Sessions_Session_SessionApplicationData_Xconnect) GetEntityData() *types.CommonEntityData {
    xconnect.EntityData.YFilter = xconnect.YFilter
    xconnect.EntityData.YangName = "xconnect"
    xconnect.EntityData.BundleName = "cisco_ios_xr"
    xconnect.EntityData.ParentYangName = "session-application-data"
    xconnect.EntityData.SegmentPath = "xconnect"
    xconnect.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    xconnect.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    xconnect.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    xconnect.EntityData.Children = make(map[string]types.YChild)
    xconnect.EntityData.Leafs = make(map[string]types.YLeaf)
    xconnect.EntityData.Leafs["circuit-name"] = types.YLeaf{"CircuitName", xconnect.CircuitName}
    xconnect.EntityData.Leafs["sessionvc-id"] = types.YLeaf{"SessionvcId", xconnect.SessionvcId}
    xconnect.EntityData.Leafs["is-circuit-state-up"] = types.YLeaf{"IsCircuitStateUp", xconnect.IsCircuitStateUp}
    xconnect.EntityData.Leafs["is-local-circuit-state-up"] = types.YLeaf{"IsLocalCircuitStateUp", xconnect.IsLocalCircuitStateUp}
    xconnect.EntityData.Leafs["is-remote-circuit-state-up"] = types.YLeaf{"IsRemoteCircuitStateUp", xconnect.IsRemoteCircuitStateUp}
    xconnect.EntityData.Leafs["ipv6-protocol-tunneling"] = types.YLeaf{"Ipv6ProtocolTunneling", xconnect.Ipv6ProtocolTunneling}
    return &(xconnect.EntityData)
}

// L2Tpv2_Sessions_Session_SessionApplicationData_Vpdn
// VPDN data
type L2Tpv2_Sessions_Session_SessionApplicationData_Vpdn struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Session username. The type is string.
    Username interface{}

    // Interface name. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}
}

func (vpdn *L2Tpv2_Sessions_Session_SessionApplicationData_Vpdn) GetEntityData() *types.CommonEntityData {
    vpdn.EntityData.YFilter = vpdn.YFilter
    vpdn.EntityData.YangName = "vpdn"
    vpdn.EntityData.BundleName = "cisco_ios_xr"
    vpdn.EntityData.ParentYangName = "session-application-data"
    vpdn.EntityData.SegmentPath = "vpdn"
    vpdn.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vpdn.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vpdn.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vpdn.EntityData.Children = make(map[string]types.YChild)
    vpdn.EntityData.Leafs = make(map[string]types.YLeaf)
    vpdn.EntityData.Leafs["username"] = types.YLeaf{"Username", vpdn.Username}
    vpdn.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", vpdn.InterfaceName}
    return &(vpdn.EntityData)
}

// L2Tpv2_Session
// L2TP control messages counters
type L2Tpv2_Session struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2TP session unavailable  information.
    Unavailable L2Tpv2_Session_Unavailable
}

func (session *L2Tpv2_Session) GetEntityData() *types.CommonEntityData {
    session.EntityData.YFilter = session.YFilter
    session.EntityData.YangName = "session"
    session.EntityData.BundleName = "cisco_ios_xr"
    session.EntityData.ParentYangName = "l2tpv2"
    session.EntityData.SegmentPath = "session"
    session.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    session.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    session.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    session.EntityData.Children = make(map[string]types.YChild)
    session.EntityData.Children["unavailable"] = types.YChild{"Unavailable", &session.Unavailable}
    session.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(session.EntityData)
}

// L2Tpv2_Session_Unavailable
// L2TP session unavailable  information
type L2Tpv2_Session_Unavailable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of session ID in hold database. The type is interface{} with range:
    // 0..4294967295.
    SessionsOnHold interface{}
}

func (unavailable *L2Tpv2_Session_Unavailable) GetEntityData() *types.CommonEntityData {
    unavailable.EntityData.YFilter = unavailable.YFilter
    unavailable.EntityData.YangName = "unavailable"
    unavailable.EntityData.BundleName = "cisco_ios_xr"
    unavailable.EntityData.ParentYangName = "session"
    unavailable.EntityData.SegmentPath = "unavailable"
    unavailable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unavailable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unavailable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unavailable.EntityData.Children = make(map[string]types.YChild)
    unavailable.EntityData.Leafs = make(map[string]types.YLeaf)
    unavailable.EntityData.Leafs["sessions-on-hold"] = types.YLeaf{"SessionsOnHold", unavailable.SessionsOnHold}
    return &(unavailable.EntityData)
}

