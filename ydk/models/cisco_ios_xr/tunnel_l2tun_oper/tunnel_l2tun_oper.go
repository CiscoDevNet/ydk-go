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
    parent types.Entity
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

func (l2Tp *L2Tp) GetFilter() yfilter.YFilter { return l2Tp.YFilter }

func (l2Tp *L2Tp) SetFilter(yf yfilter.YFilter) { l2Tp.YFilter = yf }

func (l2Tp *L2Tp) GetGoName(yname string) string {
    if yname == "counters" { return "Counters" }
    if yname == "tunnel-configurations" { return "TunnelConfigurations" }
    if yname == "counter-hist-fail" { return "CounterHistFail" }
    if yname == "classes" { return "Classes" }
    if yname == "tunnels" { return "Tunnels" }
    if yname == "sessions" { return "Sessions" }
    if yname == "session" { return "Session" }
    return ""
}

func (l2Tp *L2Tp) GetSegmentPath() string {
    return "Cisco-IOS-XR-tunnel-l2tun-oper:l2tp"
}

func (l2Tp *L2Tp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "counters" {
        return &l2Tp.Counters
    }
    if childYangName == "tunnel-configurations" {
        return &l2Tp.TunnelConfigurations
    }
    if childYangName == "counter-hist-fail" {
        return &l2Tp.CounterHistFail
    }
    if childYangName == "classes" {
        return &l2Tp.Classes
    }
    if childYangName == "tunnels" {
        return &l2Tp.Tunnels
    }
    if childYangName == "sessions" {
        return &l2Tp.Sessions
    }
    if childYangName == "session" {
        return &l2Tp.Session
    }
    return nil
}

func (l2Tp *L2Tp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["counters"] = &l2Tp.Counters
    children["tunnel-configurations"] = &l2Tp.TunnelConfigurations
    children["counter-hist-fail"] = &l2Tp.CounterHistFail
    children["classes"] = &l2Tp.Classes
    children["tunnels"] = &l2Tp.Tunnels
    children["sessions"] = &l2Tp.Sessions
    children["session"] = &l2Tp.Session
    return children
}

func (l2Tp *L2Tp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (l2Tp *L2Tp) GetBundleName() string { return "cisco_ios_xr" }

func (l2Tp *L2Tp) GetYangName() string { return "l2tp" }

func (l2Tp *L2Tp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (l2Tp *L2Tp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (l2Tp *L2Tp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (l2Tp *L2Tp) SetParent(parent types.Entity) { l2Tp.parent = parent }

func (l2Tp *L2Tp) GetParent() types.Entity { return l2Tp.parent }

func (l2Tp *L2Tp) GetParentYangName() string { return "Cisco-IOS-XR-tunnel-l2tun-oper" }

// L2Tp_Counters
// L2TP control messages counters
type L2Tp_Counters struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // L2TP control messages counters.
    Control L2Tp_Counters_Control
}

func (counters *L2Tp_Counters) GetFilter() yfilter.YFilter { return counters.YFilter }

func (counters *L2Tp_Counters) SetFilter(yf yfilter.YFilter) { counters.YFilter = yf }

func (counters *L2Tp_Counters) GetGoName(yname string) string {
    if yname == "control" { return "Control" }
    return ""
}

func (counters *L2Tp_Counters) GetSegmentPath() string {
    return "counters"
}

func (counters *L2Tp_Counters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "control" {
        return &counters.Control
    }
    return nil
}

func (counters *L2Tp_Counters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["control"] = &counters.Control
    return children
}

func (counters *L2Tp_Counters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (counters *L2Tp_Counters) GetBundleName() string { return "cisco_ios_xr" }

func (counters *L2Tp_Counters) GetYangName() string { return "counters" }

func (counters *L2Tp_Counters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (counters *L2Tp_Counters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (counters *L2Tp_Counters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (counters *L2Tp_Counters) SetParent(parent types.Entity) { counters.parent = parent }

func (counters *L2Tp_Counters) GetParent() types.Entity { return counters.parent }

func (counters *L2Tp_Counters) GetParentYangName() string { return "l2tp" }

// L2Tp_Counters_Control
// L2TP control messages counters
type L2Tp_Counters_Control struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // L2TP control tunnel messages counters.
    TunnelXr L2Tp_Counters_Control_TunnelXr

    // Table of tunnel IDs of control message counters.
    Tunnels L2Tp_Counters_Control_Tunnels
}

func (control *L2Tp_Counters_Control) GetFilter() yfilter.YFilter { return control.YFilter }

func (control *L2Tp_Counters_Control) SetFilter(yf yfilter.YFilter) { control.YFilter = yf }

func (control *L2Tp_Counters_Control) GetGoName(yname string) string {
    if yname == "tunnel-xr" { return "TunnelXr" }
    if yname == "tunnels" { return "Tunnels" }
    return ""
}

func (control *L2Tp_Counters_Control) GetSegmentPath() string {
    return "control"
}

func (control *L2Tp_Counters_Control) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tunnel-xr" {
        return &control.TunnelXr
    }
    if childYangName == "tunnels" {
        return &control.Tunnels
    }
    return nil
}

func (control *L2Tp_Counters_Control) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["tunnel-xr"] = &control.TunnelXr
    children["tunnels"] = &control.Tunnels
    return children
}

func (control *L2Tp_Counters_Control) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (control *L2Tp_Counters_Control) GetBundleName() string { return "cisco_ios_xr" }

func (control *L2Tp_Counters_Control) GetYangName() string { return "control" }

func (control *L2Tp_Counters_Control) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (control *L2Tp_Counters_Control) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (control *L2Tp_Counters_Control) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (control *L2Tp_Counters_Control) SetParent(parent types.Entity) { control.parent = parent }

func (control *L2Tp_Counters_Control) GetParent() types.Entity { return control.parent }

func (control *L2Tp_Counters_Control) GetParentYangName() string { return "counters" }

// L2Tp_Counters_Control_TunnelXr
// L2TP control tunnel messages counters
type L2Tp_Counters_Control_TunnelXr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Tunnel authentication counters.
    Authentication L2Tp_Counters_Control_TunnelXr_Authentication

    // Tunnel counters.
    Global L2Tp_Counters_Control_TunnelXr_Global
}

func (tunnelXr *L2Tp_Counters_Control_TunnelXr) GetFilter() yfilter.YFilter { return tunnelXr.YFilter }

func (tunnelXr *L2Tp_Counters_Control_TunnelXr) SetFilter(yf yfilter.YFilter) { tunnelXr.YFilter = yf }

func (tunnelXr *L2Tp_Counters_Control_TunnelXr) GetGoName(yname string) string {
    if yname == "authentication" { return "Authentication" }
    if yname == "global" { return "Global" }
    return ""
}

func (tunnelXr *L2Tp_Counters_Control_TunnelXr) GetSegmentPath() string {
    return "tunnel-xr"
}

func (tunnelXr *L2Tp_Counters_Control_TunnelXr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "authentication" {
        return &tunnelXr.Authentication
    }
    if childYangName == "global" {
        return &tunnelXr.Global
    }
    return nil
}

func (tunnelXr *L2Tp_Counters_Control_TunnelXr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["authentication"] = &tunnelXr.Authentication
    children["global"] = &tunnelXr.Global
    return children
}

func (tunnelXr *L2Tp_Counters_Control_TunnelXr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (tunnelXr *L2Tp_Counters_Control_TunnelXr) GetBundleName() string { return "cisco_ios_xr" }

func (tunnelXr *L2Tp_Counters_Control_TunnelXr) GetYangName() string { return "tunnel-xr" }

func (tunnelXr *L2Tp_Counters_Control_TunnelXr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tunnelXr *L2Tp_Counters_Control_TunnelXr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tunnelXr *L2Tp_Counters_Control_TunnelXr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tunnelXr *L2Tp_Counters_Control_TunnelXr) SetParent(parent types.Entity) { tunnelXr.parent = parent }

func (tunnelXr *L2Tp_Counters_Control_TunnelXr) GetParent() types.Entity { return tunnelXr.parent }

func (tunnelXr *L2Tp_Counters_Control_TunnelXr) GetParentYangName() string { return "control" }

// L2Tp_Counters_Control_TunnelXr_Authentication
// Tunnel authentication counters
type L2Tp_Counters_Control_TunnelXr_Authentication struct {
    parent types.Entity
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

func (authentication *L2Tp_Counters_Control_TunnelXr_Authentication) GetFilter() yfilter.YFilter { return authentication.YFilter }

func (authentication *L2Tp_Counters_Control_TunnelXr_Authentication) SetFilter(yf yfilter.YFilter) { authentication.YFilter = yf }

func (authentication *L2Tp_Counters_Control_TunnelXr_Authentication) GetGoName(yname string) string {
    if yname == "nonce-avp" { return "NonceAvp" }
    if yname == "common-digest" { return "CommonDigest" }
    if yname == "primary-digest" { return "PrimaryDigest" }
    if yname == "secondary-digest" { return "SecondaryDigest" }
    if yname == "integrity-check" { return "IntegrityCheck" }
    if yname == "local-secret" { return "LocalSecret" }
    if yname == "challenge-avp" { return "ChallengeAvp" }
    if yname == "challenge-reponse" { return "ChallengeReponse" }
    if yname == "overall-statistics" { return "OverallStatistics" }
    return ""
}

func (authentication *L2Tp_Counters_Control_TunnelXr_Authentication) GetSegmentPath() string {
    return "authentication"
}

func (authentication *L2Tp_Counters_Control_TunnelXr_Authentication) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nonce-avp" {
        return &authentication.NonceAvp
    }
    if childYangName == "common-digest" {
        return &authentication.CommonDigest
    }
    if childYangName == "primary-digest" {
        return &authentication.PrimaryDigest
    }
    if childYangName == "secondary-digest" {
        return &authentication.SecondaryDigest
    }
    if childYangName == "integrity-check" {
        return &authentication.IntegrityCheck
    }
    if childYangName == "local-secret" {
        return &authentication.LocalSecret
    }
    if childYangName == "challenge-avp" {
        return &authentication.ChallengeAvp
    }
    if childYangName == "challenge-reponse" {
        return &authentication.ChallengeReponse
    }
    if childYangName == "overall-statistics" {
        return &authentication.OverallStatistics
    }
    return nil
}

func (authentication *L2Tp_Counters_Control_TunnelXr_Authentication) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nonce-avp"] = &authentication.NonceAvp
    children["common-digest"] = &authentication.CommonDigest
    children["primary-digest"] = &authentication.PrimaryDigest
    children["secondary-digest"] = &authentication.SecondaryDigest
    children["integrity-check"] = &authentication.IntegrityCheck
    children["local-secret"] = &authentication.LocalSecret
    children["challenge-avp"] = &authentication.ChallengeAvp
    children["challenge-reponse"] = &authentication.ChallengeReponse
    children["overall-statistics"] = &authentication.OverallStatistics
    return children
}

func (authentication *L2Tp_Counters_Control_TunnelXr_Authentication) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (authentication *L2Tp_Counters_Control_TunnelXr_Authentication) GetBundleName() string { return "cisco_ios_xr" }

func (authentication *L2Tp_Counters_Control_TunnelXr_Authentication) GetYangName() string { return "authentication" }

func (authentication *L2Tp_Counters_Control_TunnelXr_Authentication) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (authentication *L2Tp_Counters_Control_TunnelXr_Authentication) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (authentication *L2Tp_Counters_Control_TunnelXr_Authentication) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (authentication *L2Tp_Counters_Control_TunnelXr_Authentication) SetParent(parent types.Entity) { authentication.parent = parent }

func (authentication *L2Tp_Counters_Control_TunnelXr_Authentication) GetParent() types.Entity { return authentication.parent }

func (authentication *L2Tp_Counters_Control_TunnelXr_Authentication) GetParentYangName() string { return "tunnel-xr" }

// L2Tp_Counters_Control_TunnelXr_Authentication_NonceAvp
// Nonce AVP statistics
type L2Tp_Counters_Control_TunnelXr_Authentication_NonceAvp struct {
    parent types.Entity
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

func (nonceAvp *L2Tp_Counters_Control_TunnelXr_Authentication_NonceAvp) GetFilter() yfilter.YFilter { return nonceAvp.YFilter }

func (nonceAvp *L2Tp_Counters_Control_TunnelXr_Authentication_NonceAvp) SetFilter(yf yfilter.YFilter) { nonceAvp.YFilter = yf }

func (nonceAvp *L2Tp_Counters_Control_TunnelXr_Authentication_NonceAvp) GetGoName(yname string) string {
    if yname == "validate" { return "Validate" }
    if yname == "bad-hash" { return "BadHash" }
    if yname == "bad-length" { return "BadLength" }
    if yname == "ignored" { return "Ignored" }
    if yname == "missing" { return "Missing" }
    if yname == "passed" { return "Passed" }
    if yname == "failed" { return "Failed" }
    if yname == "skipped" { return "Skipped" }
    if yname == "generate-response-failures" { return "GenerateResponseFailures" }
    if yname == "unexpected" { return "Unexpected" }
    if yname == "unexpected-zlb" { return "UnexpectedZlb" }
    return ""
}

func (nonceAvp *L2Tp_Counters_Control_TunnelXr_Authentication_NonceAvp) GetSegmentPath() string {
    return "nonce-avp"
}

func (nonceAvp *L2Tp_Counters_Control_TunnelXr_Authentication_NonceAvp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nonceAvp *L2Tp_Counters_Control_TunnelXr_Authentication_NonceAvp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nonceAvp *L2Tp_Counters_Control_TunnelXr_Authentication_NonceAvp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["validate"] = nonceAvp.Validate
    leafs["bad-hash"] = nonceAvp.BadHash
    leafs["bad-length"] = nonceAvp.BadLength
    leafs["ignored"] = nonceAvp.Ignored
    leafs["missing"] = nonceAvp.Missing
    leafs["passed"] = nonceAvp.Passed
    leafs["failed"] = nonceAvp.Failed
    leafs["skipped"] = nonceAvp.Skipped
    leafs["generate-response-failures"] = nonceAvp.GenerateResponseFailures
    leafs["unexpected"] = nonceAvp.Unexpected
    leafs["unexpected-zlb"] = nonceAvp.UnexpectedZlb
    return leafs
}

func (nonceAvp *L2Tp_Counters_Control_TunnelXr_Authentication_NonceAvp) GetBundleName() string { return "cisco_ios_xr" }

func (nonceAvp *L2Tp_Counters_Control_TunnelXr_Authentication_NonceAvp) GetYangName() string { return "nonce-avp" }

func (nonceAvp *L2Tp_Counters_Control_TunnelXr_Authentication_NonceAvp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nonceAvp *L2Tp_Counters_Control_TunnelXr_Authentication_NonceAvp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nonceAvp *L2Tp_Counters_Control_TunnelXr_Authentication_NonceAvp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nonceAvp *L2Tp_Counters_Control_TunnelXr_Authentication_NonceAvp) SetParent(parent types.Entity) { nonceAvp.parent = parent }

func (nonceAvp *L2Tp_Counters_Control_TunnelXr_Authentication_NonceAvp) GetParent() types.Entity { return nonceAvp.parent }

func (nonceAvp *L2Tp_Counters_Control_TunnelXr_Authentication_NonceAvp) GetParentYangName() string { return "authentication" }

// L2Tp_Counters_Control_TunnelXr_Authentication_CommonDigest
// Common digest statistics
type L2Tp_Counters_Control_TunnelXr_Authentication_CommonDigest struct {
    parent types.Entity
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

func (commonDigest *L2Tp_Counters_Control_TunnelXr_Authentication_CommonDigest) GetFilter() yfilter.YFilter { return commonDigest.YFilter }

func (commonDigest *L2Tp_Counters_Control_TunnelXr_Authentication_CommonDigest) SetFilter(yf yfilter.YFilter) { commonDigest.YFilter = yf }

func (commonDigest *L2Tp_Counters_Control_TunnelXr_Authentication_CommonDigest) GetGoName(yname string) string {
    if yname == "validate" { return "Validate" }
    if yname == "bad-hash" { return "BadHash" }
    if yname == "bad-length" { return "BadLength" }
    if yname == "ignored" { return "Ignored" }
    if yname == "missing" { return "Missing" }
    if yname == "passed" { return "Passed" }
    if yname == "failed" { return "Failed" }
    if yname == "skipped" { return "Skipped" }
    if yname == "generate-response-failures" { return "GenerateResponseFailures" }
    if yname == "unexpected" { return "Unexpected" }
    if yname == "unexpected-zlb" { return "UnexpectedZlb" }
    return ""
}

func (commonDigest *L2Tp_Counters_Control_TunnelXr_Authentication_CommonDigest) GetSegmentPath() string {
    return "common-digest"
}

func (commonDigest *L2Tp_Counters_Control_TunnelXr_Authentication_CommonDigest) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (commonDigest *L2Tp_Counters_Control_TunnelXr_Authentication_CommonDigest) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (commonDigest *L2Tp_Counters_Control_TunnelXr_Authentication_CommonDigest) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["validate"] = commonDigest.Validate
    leafs["bad-hash"] = commonDigest.BadHash
    leafs["bad-length"] = commonDigest.BadLength
    leafs["ignored"] = commonDigest.Ignored
    leafs["missing"] = commonDigest.Missing
    leafs["passed"] = commonDigest.Passed
    leafs["failed"] = commonDigest.Failed
    leafs["skipped"] = commonDigest.Skipped
    leafs["generate-response-failures"] = commonDigest.GenerateResponseFailures
    leafs["unexpected"] = commonDigest.Unexpected
    leafs["unexpected-zlb"] = commonDigest.UnexpectedZlb
    return leafs
}

func (commonDigest *L2Tp_Counters_Control_TunnelXr_Authentication_CommonDigest) GetBundleName() string { return "cisco_ios_xr" }

func (commonDigest *L2Tp_Counters_Control_TunnelXr_Authentication_CommonDigest) GetYangName() string { return "common-digest" }

func (commonDigest *L2Tp_Counters_Control_TunnelXr_Authentication_CommonDigest) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (commonDigest *L2Tp_Counters_Control_TunnelXr_Authentication_CommonDigest) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (commonDigest *L2Tp_Counters_Control_TunnelXr_Authentication_CommonDigest) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (commonDigest *L2Tp_Counters_Control_TunnelXr_Authentication_CommonDigest) SetParent(parent types.Entity) { commonDigest.parent = parent }

func (commonDigest *L2Tp_Counters_Control_TunnelXr_Authentication_CommonDigest) GetParent() types.Entity { return commonDigest.parent }

func (commonDigest *L2Tp_Counters_Control_TunnelXr_Authentication_CommonDigest) GetParentYangName() string { return "authentication" }

// L2Tp_Counters_Control_TunnelXr_Authentication_PrimaryDigest
// Primary digest statistics
type L2Tp_Counters_Control_TunnelXr_Authentication_PrimaryDigest struct {
    parent types.Entity
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

func (primaryDigest *L2Tp_Counters_Control_TunnelXr_Authentication_PrimaryDigest) GetFilter() yfilter.YFilter { return primaryDigest.YFilter }

func (primaryDigest *L2Tp_Counters_Control_TunnelXr_Authentication_PrimaryDigest) SetFilter(yf yfilter.YFilter) { primaryDigest.YFilter = yf }

func (primaryDigest *L2Tp_Counters_Control_TunnelXr_Authentication_PrimaryDigest) GetGoName(yname string) string {
    if yname == "validate" { return "Validate" }
    if yname == "bad-hash" { return "BadHash" }
    if yname == "bad-length" { return "BadLength" }
    if yname == "ignored" { return "Ignored" }
    if yname == "missing" { return "Missing" }
    if yname == "passed" { return "Passed" }
    if yname == "failed" { return "Failed" }
    if yname == "skipped" { return "Skipped" }
    if yname == "generate-response-failures" { return "GenerateResponseFailures" }
    if yname == "unexpected" { return "Unexpected" }
    if yname == "unexpected-zlb" { return "UnexpectedZlb" }
    return ""
}

func (primaryDigest *L2Tp_Counters_Control_TunnelXr_Authentication_PrimaryDigest) GetSegmentPath() string {
    return "primary-digest"
}

func (primaryDigest *L2Tp_Counters_Control_TunnelXr_Authentication_PrimaryDigest) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (primaryDigest *L2Tp_Counters_Control_TunnelXr_Authentication_PrimaryDigest) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (primaryDigest *L2Tp_Counters_Control_TunnelXr_Authentication_PrimaryDigest) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["validate"] = primaryDigest.Validate
    leafs["bad-hash"] = primaryDigest.BadHash
    leafs["bad-length"] = primaryDigest.BadLength
    leafs["ignored"] = primaryDigest.Ignored
    leafs["missing"] = primaryDigest.Missing
    leafs["passed"] = primaryDigest.Passed
    leafs["failed"] = primaryDigest.Failed
    leafs["skipped"] = primaryDigest.Skipped
    leafs["generate-response-failures"] = primaryDigest.GenerateResponseFailures
    leafs["unexpected"] = primaryDigest.Unexpected
    leafs["unexpected-zlb"] = primaryDigest.UnexpectedZlb
    return leafs
}

func (primaryDigest *L2Tp_Counters_Control_TunnelXr_Authentication_PrimaryDigest) GetBundleName() string { return "cisco_ios_xr" }

func (primaryDigest *L2Tp_Counters_Control_TunnelXr_Authentication_PrimaryDigest) GetYangName() string { return "primary-digest" }

func (primaryDigest *L2Tp_Counters_Control_TunnelXr_Authentication_PrimaryDigest) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (primaryDigest *L2Tp_Counters_Control_TunnelXr_Authentication_PrimaryDigest) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (primaryDigest *L2Tp_Counters_Control_TunnelXr_Authentication_PrimaryDigest) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (primaryDigest *L2Tp_Counters_Control_TunnelXr_Authentication_PrimaryDigest) SetParent(parent types.Entity) { primaryDigest.parent = parent }

func (primaryDigest *L2Tp_Counters_Control_TunnelXr_Authentication_PrimaryDigest) GetParent() types.Entity { return primaryDigest.parent }

func (primaryDigest *L2Tp_Counters_Control_TunnelXr_Authentication_PrimaryDigest) GetParentYangName() string { return "authentication" }

// L2Tp_Counters_Control_TunnelXr_Authentication_SecondaryDigest
// Secondary digest statistics
type L2Tp_Counters_Control_TunnelXr_Authentication_SecondaryDigest struct {
    parent types.Entity
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

func (secondaryDigest *L2Tp_Counters_Control_TunnelXr_Authentication_SecondaryDigest) GetFilter() yfilter.YFilter { return secondaryDigest.YFilter }

func (secondaryDigest *L2Tp_Counters_Control_TunnelXr_Authentication_SecondaryDigest) SetFilter(yf yfilter.YFilter) { secondaryDigest.YFilter = yf }

func (secondaryDigest *L2Tp_Counters_Control_TunnelXr_Authentication_SecondaryDigest) GetGoName(yname string) string {
    if yname == "validate" { return "Validate" }
    if yname == "bad-hash" { return "BadHash" }
    if yname == "bad-length" { return "BadLength" }
    if yname == "ignored" { return "Ignored" }
    if yname == "missing" { return "Missing" }
    if yname == "passed" { return "Passed" }
    if yname == "failed" { return "Failed" }
    if yname == "skipped" { return "Skipped" }
    if yname == "generate-response-failures" { return "GenerateResponseFailures" }
    if yname == "unexpected" { return "Unexpected" }
    if yname == "unexpected-zlb" { return "UnexpectedZlb" }
    return ""
}

func (secondaryDigest *L2Tp_Counters_Control_TunnelXr_Authentication_SecondaryDigest) GetSegmentPath() string {
    return "secondary-digest"
}

func (secondaryDigest *L2Tp_Counters_Control_TunnelXr_Authentication_SecondaryDigest) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (secondaryDigest *L2Tp_Counters_Control_TunnelXr_Authentication_SecondaryDigest) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (secondaryDigest *L2Tp_Counters_Control_TunnelXr_Authentication_SecondaryDigest) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["validate"] = secondaryDigest.Validate
    leafs["bad-hash"] = secondaryDigest.BadHash
    leafs["bad-length"] = secondaryDigest.BadLength
    leafs["ignored"] = secondaryDigest.Ignored
    leafs["missing"] = secondaryDigest.Missing
    leafs["passed"] = secondaryDigest.Passed
    leafs["failed"] = secondaryDigest.Failed
    leafs["skipped"] = secondaryDigest.Skipped
    leafs["generate-response-failures"] = secondaryDigest.GenerateResponseFailures
    leafs["unexpected"] = secondaryDigest.Unexpected
    leafs["unexpected-zlb"] = secondaryDigest.UnexpectedZlb
    return leafs
}

func (secondaryDigest *L2Tp_Counters_Control_TunnelXr_Authentication_SecondaryDigest) GetBundleName() string { return "cisco_ios_xr" }

func (secondaryDigest *L2Tp_Counters_Control_TunnelXr_Authentication_SecondaryDigest) GetYangName() string { return "secondary-digest" }

func (secondaryDigest *L2Tp_Counters_Control_TunnelXr_Authentication_SecondaryDigest) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (secondaryDigest *L2Tp_Counters_Control_TunnelXr_Authentication_SecondaryDigest) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (secondaryDigest *L2Tp_Counters_Control_TunnelXr_Authentication_SecondaryDigest) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (secondaryDigest *L2Tp_Counters_Control_TunnelXr_Authentication_SecondaryDigest) SetParent(parent types.Entity) { secondaryDigest.parent = parent }

func (secondaryDigest *L2Tp_Counters_Control_TunnelXr_Authentication_SecondaryDigest) GetParent() types.Entity { return secondaryDigest.parent }

func (secondaryDigest *L2Tp_Counters_Control_TunnelXr_Authentication_SecondaryDigest) GetParentYangName() string { return "authentication" }

// L2Tp_Counters_Control_TunnelXr_Authentication_IntegrityCheck
// Integrity check statistics
type L2Tp_Counters_Control_TunnelXr_Authentication_IntegrityCheck struct {
    parent types.Entity
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

func (integrityCheck *L2Tp_Counters_Control_TunnelXr_Authentication_IntegrityCheck) GetFilter() yfilter.YFilter { return integrityCheck.YFilter }

func (integrityCheck *L2Tp_Counters_Control_TunnelXr_Authentication_IntegrityCheck) SetFilter(yf yfilter.YFilter) { integrityCheck.YFilter = yf }

func (integrityCheck *L2Tp_Counters_Control_TunnelXr_Authentication_IntegrityCheck) GetGoName(yname string) string {
    if yname == "validate" { return "Validate" }
    if yname == "bad-hash" { return "BadHash" }
    if yname == "bad-length" { return "BadLength" }
    if yname == "ignored" { return "Ignored" }
    if yname == "missing" { return "Missing" }
    if yname == "passed" { return "Passed" }
    if yname == "failed" { return "Failed" }
    if yname == "skipped" { return "Skipped" }
    if yname == "generate-response-failures" { return "GenerateResponseFailures" }
    if yname == "unexpected" { return "Unexpected" }
    if yname == "unexpected-zlb" { return "UnexpectedZlb" }
    return ""
}

func (integrityCheck *L2Tp_Counters_Control_TunnelXr_Authentication_IntegrityCheck) GetSegmentPath() string {
    return "integrity-check"
}

func (integrityCheck *L2Tp_Counters_Control_TunnelXr_Authentication_IntegrityCheck) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (integrityCheck *L2Tp_Counters_Control_TunnelXr_Authentication_IntegrityCheck) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (integrityCheck *L2Tp_Counters_Control_TunnelXr_Authentication_IntegrityCheck) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["validate"] = integrityCheck.Validate
    leafs["bad-hash"] = integrityCheck.BadHash
    leafs["bad-length"] = integrityCheck.BadLength
    leafs["ignored"] = integrityCheck.Ignored
    leafs["missing"] = integrityCheck.Missing
    leafs["passed"] = integrityCheck.Passed
    leafs["failed"] = integrityCheck.Failed
    leafs["skipped"] = integrityCheck.Skipped
    leafs["generate-response-failures"] = integrityCheck.GenerateResponseFailures
    leafs["unexpected"] = integrityCheck.Unexpected
    leafs["unexpected-zlb"] = integrityCheck.UnexpectedZlb
    return leafs
}

func (integrityCheck *L2Tp_Counters_Control_TunnelXr_Authentication_IntegrityCheck) GetBundleName() string { return "cisco_ios_xr" }

func (integrityCheck *L2Tp_Counters_Control_TunnelXr_Authentication_IntegrityCheck) GetYangName() string { return "integrity-check" }

func (integrityCheck *L2Tp_Counters_Control_TunnelXr_Authentication_IntegrityCheck) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (integrityCheck *L2Tp_Counters_Control_TunnelXr_Authentication_IntegrityCheck) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (integrityCheck *L2Tp_Counters_Control_TunnelXr_Authentication_IntegrityCheck) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (integrityCheck *L2Tp_Counters_Control_TunnelXr_Authentication_IntegrityCheck) SetParent(parent types.Entity) { integrityCheck.parent = parent }

func (integrityCheck *L2Tp_Counters_Control_TunnelXr_Authentication_IntegrityCheck) GetParent() types.Entity { return integrityCheck.parent }

func (integrityCheck *L2Tp_Counters_Control_TunnelXr_Authentication_IntegrityCheck) GetParentYangName() string { return "authentication" }

// L2Tp_Counters_Control_TunnelXr_Authentication_LocalSecret
// Local secret statistics
type L2Tp_Counters_Control_TunnelXr_Authentication_LocalSecret struct {
    parent types.Entity
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

func (localSecret *L2Tp_Counters_Control_TunnelXr_Authentication_LocalSecret) GetFilter() yfilter.YFilter { return localSecret.YFilter }

func (localSecret *L2Tp_Counters_Control_TunnelXr_Authentication_LocalSecret) SetFilter(yf yfilter.YFilter) { localSecret.YFilter = yf }

func (localSecret *L2Tp_Counters_Control_TunnelXr_Authentication_LocalSecret) GetGoName(yname string) string {
    if yname == "validate" { return "Validate" }
    if yname == "bad-hash" { return "BadHash" }
    if yname == "bad-length" { return "BadLength" }
    if yname == "ignored" { return "Ignored" }
    if yname == "missing" { return "Missing" }
    if yname == "passed" { return "Passed" }
    if yname == "failed" { return "Failed" }
    if yname == "skipped" { return "Skipped" }
    if yname == "generate-response-failures" { return "GenerateResponseFailures" }
    if yname == "unexpected" { return "Unexpected" }
    if yname == "unexpected-zlb" { return "UnexpectedZlb" }
    return ""
}

func (localSecret *L2Tp_Counters_Control_TunnelXr_Authentication_LocalSecret) GetSegmentPath() string {
    return "local-secret"
}

func (localSecret *L2Tp_Counters_Control_TunnelXr_Authentication_LocalSecret) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (localSecret *L2Tp_Counters_Control_TunnelXr_Authentication_LocalSecret) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (localSecret *L2Tp_Counters_Control_TunnelXr_Authentication_LocalSecret) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["validate"] = localSecret.Validate
    leafs["bad-hash"] = localSecret.BadHash
    leafs["bad-length"] = localSecret.BadLength
    leafs["ignored"] = localSecret.Ignored
    leafs["missing"] = localSecret.Missing
    leafs["passed"] = localSecret.Passed
    leafs["failed"] = localSecret.Failed
    leafs["skipped"] = localSecret.Skipped
    leafs["generate-response-failures"] = localSecret.GenerateResponseFailures
    leafs["unexpected"] = localSecret.Unexpected
    leafs["unexpected-zlb"] = localSecret.UnexpectedZlb
    return leafs
}

func (localSecret *L2Tp_Counters_Control_TunnelXr_Authentication_LocalSecret) GetBundleName() string { return "cisco_ios_xr" }

func (localSecret *L2Tp_Counters_Control_TunnelXr_Authentication_LocalSecret) GetYangName() string { return "local-secret" }

func (localSecret *L2Tp_Counters_Control_TunnelXr_Authentication_LocalSecret) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (localSecret *L2Tp_Counters_Control_TunnelXr_Authentication_LocalSecret) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (localSecret *L2Tp_Counters_Control_TunnelXr_Authentication_LocalSecret) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (localSecret *L2Tp_Counters_Control_TunnelXr_Authentication_LocalSecret) SetParent(parent types.Entity) { localSecret.parent = parent }

func (localSecret *L2Tp_Counters_Control_TunnelXr_Authentication_LocalSecret) GetParent() types.Entity { return localSecret.parent }

func (localSecret *L2Tp_Counters_Control_TunnelXr_Authentication_LocalSecret) GetParentYangName() string { return "authentication" }

// L2Tp_Counters_Control_TunnelXr_Authentication_ChallengeAvp
// Challenge AVP statistics
type L2Tp_Counters_Control_TunnelXr_Authentication_ChallengeAvp struct {
    parent types.Entity
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

func (challengeAvp *L2Tp_Counters_Control_TunnelXr_Authentication_ChallengeAvp) GetFilter() yfilter.YFilter { return challengeAvp.YFilter }

func (challengeAvp *L2Tp_Counters_Control_TunnelXr_Authentication_ChallengeAvp) SetFilter(yf yfilter.YFilter) { challengeAvp.YFilter = yf }

func (challengeAvp *L2Tp_Counters_Control_TunnelXr_Authentication_ChallengeAvp) GetGoName(yname string) string {
    if yname == "validate" { return "Validate" }
    if yname == "bad-hash" { return "BadHash" }
    if yname == "bad-length" { return "BadLength" }
    if yname == "ignored" { return "Ignored" }
    if yname == "missing" { return "Missing" }
    if yname == "passed" { return "Passed" }
    if yname == "failed" { return "Failed" }
    if yname == "skipped" { return "Skipped" }
    if yname == "generate-response-failures" { return "GenerateResponseFailures" }
    if yname == "unexpected" { return "Unexpected" }
    if yname == "unexpected-zlb" { return "UnexpectedZlb" }
    return ""
}

func (challengeAvp *L2Tp_Counters_Control_TunnelXr_Authentication_ChallengeAvp) GetSegmentPath() string {
    return "challenge-avp"
}

func (challengeAvp *L2Tp_Counters_Control_TunnelXr_Authentication_ChallengeAvp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (challengeAvp *L2Tp_Counters_Control_TunnelXr_Authentication_ChallengeAvp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (challengeAvp *L2Tp_Counters_Control_TunnelXr_Authentication_ChallengeAvp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["validate"] = challengeAvp.Validate
    leafs["bad-hash"] = challengeAvp.BadHash
    leafs["bad-length"] = challengeAvp.BadLength
    leafs["ignored"] = challengeAvp.Ignored
    leafs["missing"] = challengeAvp.Missing
    leafs["passed"] = challengeAvp.Passed
    leafs["failed"] = challengeAvp.Failed
    leafs["skipped"] = challengeAvp.Skipped
    leafs["generate-response-failures"] = challengeAvp.GenerateResponseFailures
    leafs["unexpected"] = challengeAvp.Unexpected
    leafs["unexpected-zlb"] = challengeAvp.UnexpectedZlb
    return leafs
}

func (challengeAvp *L2Tp_Counters_Control_TunnelXr_Authentication_ChallengeAvp) GetBundleName() string { return "cisco_ios_xr" }

func (challengeAvp *L2Tp_Counters_Control_TunnelXr_Authentication_ChallengeAvp) GetYangName() string { return "challenge-avp" }

func (challengeAvp *L2Tp_Counters_Control_TunnelXr_Authentication_ChallengeAvp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (challengeAvp *L2Tp_Counters_Control_TunnelXr_Authentication_ChallengeAvp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (challengeAvp *L2Tp_Counters_Control_TunnelXr_Authentication_ChallengeAvp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (challengeAvp *L2Tp_Counters_Control_TunnelXr_Authentication_ChallengeAvp) SetParent(parent types.Entity) { challengeAvp.parent = parent }

func (challengeAvp *L2Tp_Counters_Control_TunnelXr_Authentication_ChallengeAvp) GetParent() types.Entity { return challengeAvp.parent }

func (challengeAvp *L2Tp_Counters_Control_TunnelXr_Authentication_ChallengeAvp) GetParentYangName() string { return "authentication" }

// L2Tp_Counters_Control_TunnelXr_Authentication_ChallengeReponse
// Challenge response statistics
type L2Tp_Counters_Control_TunnelXr_Authentication_ChallengeReponse struct {
    parent types.Entity
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

func (challengeReponse *L2Tp_Counters_Control_TunnelXr_Authentication_ChallengeReponse) GetFilter() yfilter.YFilter { return challengeReponse.YFilter }

func (challengeReponse *L2Tp_Counters_Control_TunnelXr_Authentication_ChallengeReponse) SetFilter(yf yfilter.YFilter) { challengeReponse.YFilter = yf }

func (challengeReponse *L2Tp_Counters_Control_TunnelXr_Authentication_ChallengeReponse) GetGoName(yname string) string {
    if yname == "validate" { return "Validate" }
    if yname == "bad-hash" { return "BadHash" }
    if yname == "bad-length" { return "BadLength" }
    if yname == "ignored" { return "Ignored" }
    if yname == "missing" { return "Missing" }
    if yname == "passed" { return "Passed" }
    if yname == "failed" { return "Failed" }
    if yname == "skipped" { return "Skipped" }
    if yname == "generate-response-failures" { return "GenerateResponseFailures" }
    if yname == "unexpected" { return "Unexpected" }
    if yname == "unexpected-zlb" { return "UnexpectedZlb" }
    return ""
}

func (challengeReponse *L2Tp_Counters_Control_TunnelXr_Authentication_ChallengeReponse) GetSegmentPath() string {
    return "challenge-reponse"
}

func (challengeReponse *L2Tp_Counters_Control_TunnelXr_Authentication_ChallengeReponse) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (challengeReponse *L2Tp_Counters_Control_TunnelXr_Authentication_ChallengeReponse) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (challengeReponse *L2Tp_Counters_Control_TunnelXr_Authentication_ChallengeReponse) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["validate"] = challengeReponse.Validate
    leafs["bad-hash"] = challengeReponse.BadHash
    leafs["bad-length"] = challengeReponse.BadLength
    leafs["ignored"] = challengeReponse.Ignored
    leafs["missing"] = challengeReponse.Missing
    leafs["passed"] = challengeReponse.Passed
    leafs["failed"] = challengeReponse.Failed
    leafs["skipped"] = challengeReponse.Skipped
    leafs["generate-response-failures"] = challengeReponse.GenerateResponseFailures
    leafs["unexpected"] = challengeReponse.Unexpected
    leafs["unexpected-zlb"] = challengeReponse.UnexpectedZlb
    return leafs
}

func (challengeReponse *L2Tp_Counters_Control_TunnelXr_Authentication_ChallengeReponse) GetBundleName() string { return "cisco_ios_xr" }

func (challengeReponse *L2Tp_Counters_Control_TunnelXr_Authentication_ChallengeReponse) GetYangName() string { return "challenge-reponse" }

func (challengeReponse *L2Tp_Counters_Control_TunnelXr_Authentication_ChallengeReponse) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (challengeReponse *L2Tp_Counters_Control_TunnelXr_Authentication_ChallengeReponse) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (challengeReponse *L2Tp_Counters_Control_TunnelXr_Authentication_ChallengeReponse) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (challengeReponse *L2Tp_Counters_Control_TunnelXr_Authentication_ChallengeReponse) SetParent(parent types.Entity) { challengeReponse.parent = parent }

func (challengeReponse *L2Tp_Counters_Control_TunnelXr_Authentication_ChallengeReponse) GetParent() types.Entity { return challengeReponse.parent }

func (challengeReponse *L2Tp_Counters_Control_TunnelXr_Authentication_ChallengeReponse) GetParentYangName() string { return "authentication" }

// L2Tp_Counters_Control_TunnelXr_Authentication_OverallStatistics
// Overall statistics
type L2Tp_Counters_Control_TunnelXr_Authentication_OverallStatistics struct {
    parent types.Entity
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

func (overallStatistics *L2Tp_Counters_Control_TunnelXr_Authentication_OverallStatistics) GetFilter() yfilter.YFilter { return overallStatistics.YFilter }

func (overallStatistics *L2Tp_Counters_Control_TunnelXr_Authentication_OverallStatistics) SetFilter(yf yfilter.YFilter) { overallStatistics.YFilter = yf }

func (overallStatistics *L2Tp_Counters_Control_TunnelXr_Authentication_OverallStatistics) GetGoName(yname string) string {
    if yname == "validate" { return "Validate" }
    if yname == "bad-hash" { return "BadHash" }
    if yname == "bad-length" { return "BadLength" }
    if yname == "ignored" { return "Ignored" }
    if yname == "missing" { return "Missing" }
    if yname == "passed" { return "Passed" }
    if yname == "failed" { return "Failed" }
    if yname == "skipped" { return "Skipped" }
    if yname == "generate-response-failures" { return "GenerateResponseFailures" }
    if yname == "unexpected" { return "Unexpected" }
    if yname == "unexpected-zlb" { return "UnexpectedZlb" }
    return ""
}

func (overallStatistics *L2Tp_Counters_Control_TunnelXr_Authentication_OverallStatistics) GetSegmentPath() string {
    return "overall-statistics"
}

func (overallStatistics *L2Tp_Counters_Control_TunnelXr_Authentication_OverallStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (overallStatistics *L2Tp_Counters_Control_TunnelXr_Authentication_OverallStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (overallStatistics *L2Tp_Counters_Control_TunnelXr_Authentication_OverallStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["validate"] = overallStatistics.Validate
    leafs["bad-hash"] = overallStatistics.BadHash
    leafs["bad-length"] = overallStatistics.BadLength
    leafs["ignored"] = overallStatistics.Ignored
    leafs["missing"] = overallStatistics.Missing
    leafs["passed"] = overallStatistics.Passed
    leafs["failed"] = overallStatistics.Failed
    leafs["skipped"] = overallStatistics.Skipped
    leafs["generate-response-failures"] = overallStatistics.GenerateResponseFailures
    leafs["unexpected"] = overallStatistics.Unexpected
    leafs["unexpected-zlb"] = overallStatistics.UnexpectedZlb
    return leafs
}

func (overallStatistics *L2Tp_Counters_Control_TunnelXr_Authentication_OverallStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (overallStatistics *L2Tp_Counters_Control_TunnelXr_Authentication_OverallStatistics) GetYangName() string { return "overall-statistics" }

func (overallStatistics *L2Tp_Counters_Control_TunnelXr_Authentication_OverallStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (overallStatistics *L2Tp_Counters_Control_TunnelXr_Authentication_OverallStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (overallStatistics *L2Tp_Counters_Control_TunnelXr_Authentication_OverallStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (overallStatistics *L2Tp_Counters_Control_TunnelXr_Authentication_OverallStatistics) SetParent(parent types.Entity) { overallStatistics.parent = parent }

func (overallStatistics *L2Tp_Counters_Control_TunnelXr_Authentication_OverallStatistics) GetParent() types.Entity { return overallStatistics.parent }

func (overallStatistics *L2Tp_Counters_Control_TunnelXr_Authentication_OverallStatistics) GetParentYangName() string { return "authentication" }

// L2Tp_Counters_Control_TunnelXr_Global
// Tunnel counters
type L2Tp_Counters_Control_TunnelXr_Global struct {
    parent types.Entity
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

func (global *L2Tp_Counters_Control_TunnelXr_Global) GetFilter() yfilter.YFilter { return global.YFilter }

func (global *L2Tp_Counters_Control_TunnelXr_Global) SetFilter(yf yfilter.YFilter) { global.YFilter = yf }

func (global *L2Tp_Counters_Control_TunnelXr_Global) GetGoName(yname string) string {
    if yname == "total-transmit" { return "TotalTransmit" }
    if yname == "total-retransmit" { return "TotalRetransmit" }
    if yname == "total-received" { return "TotalReceived" }
    if yname == "total-drop" { return "TotalDrop" }
    if yname == "transmit" { return "Transmit" }
    if yname == "retransmit" { return "Retransmit" }
    if yname == "received" { return "Received" }
    if yname == "drop" { return "Drop" }
    return ""
}

func (global *L2Tp_Counters_Control_TunnelXr_Global) GetSegmentPath() string {
    return "global"
}

func (global *L2Tp_Counters_Control_TunnelXr_Global) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "transmit" {
        return &global.Transmit
    }
    if childYangName == "retransmit" {
        return &global.Retransmit
    }
    if childYangName == "received" {
        return &global.Received
    }
    if childYangName == "drop" {
        return &global.Drop
    }
    return nil
}

func (global *L2Tp_Counters_Control_TunnelXr_Global) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["transmit"] = &global.Transmit
    children["retransmit"] = &global.Retransmit
    children["received"] = &global.Received
    children["drop"] = &global.Drop
    return children
}

func (global *L2Tp_Counters_Control_TunnelXr_Global) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["total-transmit"] = global.TotalTransmit
    leafs["total-retransmit"] = global.TotalRetransmit
    leafs["total-received"] = global.TotalReceived
    leafs["total-drop"] = global.TotalDrop
    return leafs
}

func (global *L2Tp_Counters_Control_TunnelXr_Global) GetBundleName() string { return "cisco_ios_xr" }

func (global *L2Tp_Counters_Control_TunnelXr_Global) GetYangName() string { return "global" }

func (global *L2Tp_Counters_Control_TunnelXr_Global) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (global *L2Tp_Counters_Control_TunnelXr_Global) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (global *L2Tp_Counters_Control_TunnelXr_Global) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (global *L2Tp_Counters_Control_TunnelXr_Global) SetParent(parent types.Entity) { global.parent = parent }

func (global *L2Tp_Counters_Control_TunnelXr_Global) GetParent() types.Entity { return global.parent }

func (global *L2Tp_Counters_Control_TunnelXr_Global) GetParentYangName() string { return "tunnel-xr" }

// L2Tp_Counters_Control_TunnelXr_Global_Transmit
// Transmit data
type L2Tp_Counters_Control_TunnelXr_Global_Transmit struct {
    parent types.Entity
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

func (transmit *L2Tp_Counters_Control_TunnelXr_Global_Transmit) GetFilter() yfilter.YFilter { return transmit.YFilter }

func (transmit *L2Tp_Counters_Control_TunnelXr_Global_Transmit) SetFilter(yf yfilter.YFilter) { transmit.YFilter = yf }

func (transmit *L2Tp_Counters_Control_TunnelXr_Global_Transmit) GetGoName(yname string) string {
    if yname == "unknown-packets" { return "UnknownPackets" }
    if yname == "zero-length-body-packets" { return "ZeroLengthBodyPackets" }
    if yname == "start-control-connection-requests" { return "StartControlConnectionRequests" }
    if yname == "start-control-connection-replies" { return "StartControlConnectionReplies" }
    if yname == "start-control-connection-notifications" { return "StartControlConnectionNotifications" }
    if yname == "stop-control-connection-notifications" { return "StopControlConnectionNotifications" }
    if yname == "hello-packets" { return "HelloPackets" }
    if yname == "outgoing-call-requests" { return "OutgoingCallRequests" }
    if yname == "outgoing-call-replies" { return "OutgoingCallReplies" }
    if yname == "outgoing-call-connected-packets" { return "OutgoingCallConnectedPackets" }
    if yname == "incoming-call-requests" { return "IncomingCallRequests" }
    if yname == "incoming-call-replies" { return "IncomingCallReplies" }
    if yname == "incoming-call-connected-packets" { return "IncomingCallConnectedPackets" }
    if yname == "call-disconnect-notify-packets" { return "CallDisconnectNotifyPackets" }
    if yname == "wan-error-notify-packets" { return "WanErrorNotifyPackets" }
    if yname == "set-link-info-packets" { return "SetLinkInfoPackets" }
    if yname == "service-relay-requests" { return "ServiceRelayRequests" }
    if yname == "service-relay-replies" { return "ServiceRelayReplies" }
    if yname == "acknowledgement-packets" { return "AcknowledgementPackets" }
    return ""
}

func (transmit *L2Tp_Counters_Control_TunnelXr_Global_Transmit) GetSegmentPath() string {
    return "transmit"
}

func (transmit *L2Tp_Counters_Control_TunnelXr_Global_Transmit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (transmit *L2Tp_Counters_Control_TunnelXr_Global_Transmit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (transmit *L2Tp_Counters_Control_TunnelXr_Global_Transmit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["unknown-packets"] = transmit.UnknownPackets
    leafs["zero-length-body-packets"] = transmit.ZeroLengthBodyPackets
    leafs["start-control-connection-requests"] = transmit.StartControlConnectionRequests
    leafs["start-control-connection-replies"] = transmit.StartControlConnectionReplies
    leafs["start-control-connection-notifications"] = transmit.StartControlConnectionNotifications
    leafs["stop-control-connection-notifications"] = transmit.StopControlConnectionNotifications
    leafs["hello-packets"] = transmit.HelloPackets
    leafs["outgoing-call-requests"] = transmit.OutgoingCallRequests
    leafs["outgoing-call-replies"] = transmit.OutgoingCallReplies
    leafs["outgoing-call-connected-packets"] = transmit.OutgoingCallConnectedPackets
    leafs["incoming-call-requests"] = transmit.IncomingCallRequests
    leafs["incoming-call-replies"] = transmit.IncomingCallReplies
    leafs["incoming-call-connected-packets"] = transmit.IncomingCallConnectedPackets
    leafs["call-disconnect-notify-packets"] = transmit.CallDisconnectNotifyPackets
    leafs["wan-error-notify-packets"] = transmit.WanErrorNotifyPackets
    leafs["set-link-info-packets"] = transmit.SetLinkInfoPackets
    leafs["service-relay-requests"] = transmit.ServiceRelayRequests
    leafs["service-relay-replies"] = transmit.ServiceRelayReplies
    leafs["acknowledgement-packets"] = transmit.AcknowledgementPackets
    return leafs
}

func (transmit *L2Tp_Counters_Control_TunnelXr_Global_Transmit) GetBundleName() string { return "cisco_ios_xr" }

func (transmit *L2Tp_Counters_Control_TunnelXr_Global_Transmit) GetYangName() string { return "transmit" }

func (transmit *L2Tp_Counters_Control_TunnelXr_Global_Transmit) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (transmit *L2Tp_Counters_Control_TunnelXr_Global_Transmit) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (transmit *L2Tp_Counters_Control_TunnelXr_Global_Transmit) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (transmit *L2Tp_Counters_Control_TunnelXr_Global_Transmit) SetParent(parent types.Entity) { transmit.parent = parent }

func (transmit *L2Tp_Counters_Control_TunnelXr_Global_Transmit) GetParent() types.Entity { return transmit.parent }

func (transmit *L2Tp_Counters_Control_TunnelXr_Global_Transmit) GetParentYangName() string { return "global" }

// L2Tp_Counters_Control_TunnelXr_Global_Retransmit
// Re transmit data
type L2Tp_Counters_Control_TunnelXr_Global_Retransmit struct {
    parent types.Entity
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

func (retransmit *L2Tp_Counters_Control_TunnelXr_Global_Retransmit) GetFilter() yfilter.YFilter { return retransmit.YFilter }

func (retransmit *L2Tp_Counters_Control_TunnelXr_Global_Retransmit) SetFilter(yf yfilter.YFilter) { retransmit.YFilter = yf }

func (retransmit *L2Tp_Counters_Control_TunnelXr_Global_Retransmit) GetGoName(yname string) string {
    if yname == "unknown-packets" { return "UnknownPackets" }
    if yname == "zero-length-body-packets" { return "ZeroLengthBodyPackets" }
    if yname == "start-control-connection-requests" { return "StartControlConnectionRequests" }
    if yname == "start-control-connection-replies" { return "StartControlConnectionReplies" }
    if yname == "start-control-connection-notifications" { return "StartControlConnectionNotifications" }
    if yname == "stop-control-connection-notifications" { return "StopControlConnectionNotifications" }
    if yname == "hello-packets" { return "HelloPackets" }
    if yname == "outgoing-call-requests" { return "OutgoingCallRequests" }
    if yname == "outgoing-call-replies" { return "OutgoingCallReplies" }
    if yname == "outgoing-call-connected-packets" { return "OutgoingCallConnectedPackets" }
    if yname == "incoming-call-requests" { return "IncomingCallRequests" }
    if yname == "incoming-call-replies" { return "IncomingCallReplies" }
    if yname == "incoming-call-connected-packets" { return "IncomingCallConnectedPackets" }
    if yname == "call-disconnect-notify-packets" { return "CallDisconnectNotifyPackets" }
    if yname == "wan-error-notify-packets" { return "WanErrorNotifyPackets" }
    if yname == "set-link-info-packets" { return "SetLinkInfoPackets" }
    if yname == "service-relay-requests" { return "ServiceRelayRequests" }
    if yname == "service-relay-replies" { return "ServiceRelayReplies" }
    if yname == "acknowledgement-packets" { return "AcknowledgementPackets" }
    return ""
}

func (retransmit *L2Tp_Counters_Control_TunnelXr_Global_Retransmit) GetSegmentPath() string {
    return "retransmit"
}

func (retransmit *L2Tp_Counters_Control_TunnelXr_Global_Retransmit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (retransmit *L2Tp_Counters_Control_TunnelXr_Global_Retransmit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (retransmit *L2Tp_Counters_Control_TunnelXr_Global_Retransmit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["unknown-packets"] = retransmit.UnknownPackets
    leafs["zero-length-body-packets"] = retransmit.ZeroLengthBodyPackets
    leafs["start-control-connection-requests"] = retransmit.StartControlConnectionRequests
    leafs["start-control-connection-replies"] = retransmit.StartControlConnectionReplies
    leafs["start-control-connection-notifications"] = retransmit.StartControlConnectionNotifications
    leafs["stop-control-connection-notifications"] = retransmit.StopControlConnectionNotifications
    leafs["hello-packets"] = retransmit.HelloPackets
    leafs["outgoing-call-requests"] = retransmit.OutgoingCallRequests
    leafs["outgoing-call-replies"] = retransmit.OutgoingCallReplies
    leafs["outgoing-call-connected-packets"] = retransmit.OutgoingCallConnectedPackets
    leafs["incoming-call-requests"] = retransmit.IncomingCallRequests
    leafs["incoming-call-replies"] = retransmit.IncomingCallReplies
    leafs["incoming-call-connected-packets"] = retransmit.IncomingCallConnectedPackets
    leafs["call-disconnect-notify-packets"] = retransmit.CallDisconnectNotifyPackets
    leafs["wan-error-notify-packets"] = retransmit.WanErrorNotifyPackets
    leafs["set-link-info-packets"] = retransmit.SetLinkInfoPackets
    leafs["service-relay-requests"] = retransmit.ServiceRelayRequests
    leafs["service-relay-replies"] = retransmit.ServiceRelayReplies
    leafs["acknowledgement-packets"] = retransmit.AcknowledgementPackets
    return leafs
}

func (retransmit *L2Tp_Counters_Control_TunnelXr_Global_Retransmit) GetBundleName() string { return "cisco_ios_xr" }

func (retransmit *L2Tp_Counters_Control_TunnelXr_Global_Retransmit) GetYangName() string { return "retransmit" }

func (retransmit *L2Tp_Counters_Control_TunnelXr_Global_Retransmit) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (retransmit *L2Tp_Counters_Control_TunnelXr_Global_Retransmit) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (retransmit *L2Tp_Counters_Control_TunnelXr_Global_Retransmit) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (retransmit *L2Tp_Counters_Control_TunnelXr_Global_Retransmit) SetParent(parent types.Entity) { retransmit.parent = parent }

func (retransmit *L2Tp_Counters_Control_TunnelXr_Global_Retransmit) GetParent() types.Entity { return retransmit.parent }

func (retransmit *L2Tp_Counters_Control_TunnelXr_Global_Retransmit) GetParentYangName() string { return "global" }

// L2Tp_Counters_Control_TunnelXr_Global_Received
// Received data
type L2Tp_Counters_Control_TunnelXr_Global_Received struct {
    parent types.Entity
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

func (received *L2Tp_Counters_Control_TunnelXr_Global_Received) GetFilter() yfilter.YFilter { return received.YFilter }

func (received *L2Tp_Counters_Control_TunnelXr_Global_Received) SetFilter(yf yfilter.YFilter) { received.YFilter = yf }

func (received *L2Tp_Counters_Control_TunnelXr_Global_Received) GetGoName(yname string) string {
    if yname == "unknown-packets" { return "UnknownPackets" }
    if yname == "zero-length-body-packets" { return "ZeroLengthBodyPackets" }
    if yname == "start-control-connection-requests" { return "StartControlConnectionRequests" }
    if yname == "start-control-connection-replies" { return "StartControlConnectionReplies" }
    if yname == "start-control-connection-notifications" { return "StartControlConnectionNotifications" }
    if yname == "stop-control-connection-notifications" { return "StopControlConnectionNotifications" }
    if yname == "hello-packets" { return "HelloPackets" }
    if yname == "outgoing-call-requests" { return "OutgoingCallRequests" }
    if yname == "outgoing-call-replies" { return "OutgoingCallReplies" }
    if yname == "outgoing-call-connected-packets" { return "OutgoingCallConnectedPackets" }
    if yname == "incoming-call-requests" { return "IncomingCallRequests" }
    if yname == "incoming-call-replies" { return "IncomingCallReplies" }
    if yname == "incoming-call-connected-packets" { return "IncomingCallConnectedPackets" }
    if yname == "call-disconnect-notify-packets" { return "CallDisconnectNotifyPackets" }
    if yname == "wan-error-notify-packets" { return "WanErrorNotifyPackets" }
    if yname == "set-link-info-packets" { return "SetLinkInfoPackets" }
    if yname == "service-relay-requests" { return "ServiceRelayRequests" }
    if yname == "service-relay-replies" { return "ServiceRelayReplies" }
    if yname == "acknowledgement-packets" { return "AcknowledgementPackets" }
    return ""
}

func (received *L2Tp_Counters_Control_TunnelXr_Global_Received) GetSegmentPath() string {
    return "received"
}

func (received *L2Tp_Counters_Control_TunnelXr_Global_Received) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (received *L2Tp_Counters_Control_TunnelXr_Global_Received) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (received *L2Tp_Counters_Control_TunnelXr_Global_Received) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["unknown-packets"] = received.UnknownPackets
    leafs["zero-length-body-packets"] = received.ZeroLengthBodyPackets
    leafs["start-control-connection-requests"] = received.StartControlConnectionRequests
    leafs["start-control-connection-replies"] = received.StartControlConnectionReplies
    leafs["start-control-connection-notifications"] = received.StartControlConnectionNotifications
    leafs["stop-control-connection-notifications"] = received.StopControlConnectionNotifications
    leafs["hello-packets"] = received.HelloPackets
    leafs["outgoing-call-requests"] = received.OutgoingCallRequests
    leafs["outgoing-call-replies"] = received.OutgoingCallReplies
    leafs["outgoing-call-connected-packets"] = received.OutgoingCallConnectedPackets
    leafs["incoming-call-requests"] = received.IncomingCallRequests
    leafs["incoming-call-replies"] = received.IncomingCallReplies
    leafs["incoming-call-connected-packets"] = received.IncomingCallConnectedPackets
    leafs["call-disconnect-notify-packets"] = received.CallDisconnectNotifyPackets
    leafs["wan-error-notify-packets"] = received.WanErrorNotifyPackets
    leafs["set-link-info-packets"] = received.SetLinkInfoPackets
    leafs["service-relay-requests"] = received.ServiceRelayRequests
    leafs["service-relay-replies"] = received.ServiceRelayReplies
    leafs["acknowledgement-packets"] = received.AcknowledgementPackets
    return leafs
}

func (received *L2Tp_Counters_Control_TunnelXr_Global_Received) GetBundleName() string { return "cisco_ios_xr" }

func (received *L2Tp_Counters_Control_TunnelXr_Global_Received) GetYangName() string { return "received" }

func (received *L2Tp_Counters_Control_TunnelXr_Global_Received) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (received *L2Tp_Counters_Control_TunnelXr_Global_Received) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (received *L2Tp_Counters_Control_TunnelXr_Global_Received) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (received *L2Tp_Counters_Control_TunnelXr_Global_Received) SetParent(parent types.Entity) { received.parent = parent }

func (received *L2Tp_Counters_Control_TunnelXr_Global_Received) GetParent() types.Entity { return received.parent }

func (received *L2Tp_Counters_Control_TunnelXr_Global_Received) GetParentYangName() string { return "global" }

// L2Tp_Counters_Control_TunnelXr_Global_Drop
// Drop data
type L2Tp_Counters_Control_TunnelXr_Global_Drop struct {
    parent types.Entity
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

func (drop *L2Tp_Counters_Control_TunnelXr_Global_Drop) GetFilter() yfilter.YFilter { return drop.YFilter }

func (drop *L2Tp_Counters_Control_TunnelXr_Global_Drop) SetFilter(yf yfilter.YFilter) { drop.YFilter = yf }

func (drop *L2Tp_Counters_Control_TunnelXr_Global_Drop) GetGoName(yname string) string {
    if yname == "unknown-packets" { return "UnknownPackets" }
    if yname == "zero-length-body-packets" { return "ZeroLengthBodyPackets" }
    if yname == "start-control-connection-requests" { return "StartControlConnectionRequests" }
    if yname == "start-control-connection-replies" { return "StartControlConnectionReplies" }
    if yname == "start-control-connection-notifications" { return "StartControlConnectionNotifications" }
    if yname == "stop-control-connection-notifications" { return "StopControlConnectionNotifications" }
    if yname == "hello-packets" { return "HelloPackets" }
    if yname == "outgoing-call-requests" { return "OutgoingCallRequests" }
    if yname == "outgoing-call-replies" { return "OutgoingCallReplies" }
    if yname == "outgoing-call-connected-packets" { return "OutgoingCallConnectedPackets" }
    if yname == "incoming-call-requests" { return "IncomingCallRequests" }
    if yname == "incoming-call-replies" { return "IncomingCallReplies" }
    if yname == "incoming-call-connected-packets" { return "IncomingCallConnectedPackets" }
    if yname == "call-disconnect-notify-packets" { return "CallDisconnectNotifyPackets" }
    if yname == "wan-error-notify-packets" { return "WanErrorNotifyPackets" }
    if yname == "set-link-info-packets" { return "SetLinkInfoPackets" }
    if yname == "service-relay-requests" { return "ServiceRelayRequests" }
    if yname == "service-relay-replies" { return "ServiceRelayReplies" }
    if yname == "acknowledgement-packets" { return "AcknowledgementPackets" }
    return ""
}

func (drop *L2Tp_Counters_Control_TunnelXr_Global_Drop) GetSegmentPath() string {
    return "drop"
}

func (drop *L2Tp_Counters_Control_TunnelXr_Global_Drop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (drop *L2Tp_Counters_Control_TunnelXr_Global_Drop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (drop *L2Tp_Counters_Control_TunnelXr_Global_Drop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["unknown-packets"] = drop.UnknownPackets
    leafs["zero-length-body-packets"] = drop.ZeroLengthBodyPackets
    leafs["start-control-connection-requests"] = drop.StartControlConnectionRequests
    leafs["start-control-connection-replies"] = drop.StartControlConnectionReplies
    leafs["start-control-connection-notifications"] = drop.StartControlConnectionNotifications
    leafs["stop-control-connection-notifications"] = drop.StopControlConnectionNotifications
    leafs["hello-packets"] = drop.HelloPackets
    leafs["outgoing-call-requests"] = drop.OutgoingCallRequests
    leafs["outgoing-call-replies"] = drop.OutgoingCallReplies
    leafs["outgoing-call-connected-packets"] = drop.OutgoingCallConnectedPackets
    leafs["incoming-call-requests"] = drop.IncomingCallRequests
    leafs["incoming-call-replies"] = drop.IncomingCallReplies
    leafs["incoming-call-connected-packets"] = drop.IncomingCallConnectedPackets
    leafs["call-disconnect-notify-packets"] = drop.CallDisconnectNotifyPackets
    leafs["wan-error-notify-packets"] = drop.WanErrorNotifyPackets
    leafs["set-link-info-packets"] = drop.SetLinkInfoPackets
    leafs["service-relay-requests"] = drop.ServiceRelayRequests
    leafs["service-relay-replies"] = drop.ServiceRelayReplies
    leafs["acknowledgement-packets"] = drop.AcknowledgementPackets
    return leafs
}

func (drop *L2Tp_Counters_Control_TunnelXr_Global_Drop) GetBundleName() string { return "cisco_ios_xr" }

func (drop *L2Tp_Counters_Control_TunnelXr_Global_Drop) GetYangName() string { return "drop" }

func (drop *L2Tp_Counters_Control_TunnelXr_Global_Drop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (drop *L2Tp_Counters_Control_TunnelXr_Global_Drop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (drop *L2Tp_Counters_Control_TunnelXr_Global_Drop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (drop *L2Tp_Counters_Control_TunnelXr_Global_Drop) SetParent(parent types.Entity) { drop.parent = parent }

func (drop *L2Tp_Counters_Control_TunnelXr_Global_Drop) GetParent() types.Entity { return drop.parent }

func (drop *L2Tp_Counters_Control_TunnelXr_Global_Drop) GetParentYangName() string { return "global" }

// L2Tp_Counters_Control_Tunnels
// Table of tunnel IDs of control message counters
type L2Tp_Counters_Control_Tunnels struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // L2TP tunnel control message counters. The type is slice of
    // L2Tp_Counters_Control_Tunnels_Tunnel.
    Tunnel []L2Tp_Counters_Control_Tunnels_Tunnel
}

func (tunnels *L2Tp_Counters_Control_Tunnels) GetFilter() yfilter.YFilter { return tunnels.YFilter }

func (tunnels *L2Tp_Counters_Control_Tunnels) SetFilter(yf yfilter.YFilter) { tunnels.YFilter = yf }

func (tunnels *L2Tp_Counters_Control_Tunnels) GetGoName(yname string) string {
    if yname == "tunnel" { return "Tunnel" }
    return ""
}

func (tunnels *L2Tp_Counters_Control_Tunnels) GetSegmentPath() string {
    return "tunnels"
}

func (tunnels *L2Tp_Counters_Control_Tunnels) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tunnel" {
        for _, c := range tunnels.Tunnel {
            if tunnels.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := L2Tp_Counters_Control_Tunnels_Tunnel{}
        tunnels.Tunnel = append(tunnels.Tunnel, child)
        return &tunnels.Tunnel[len(tunnels.Tunnel)-1]
    }
    return nil
}

func (tunnels *L2Tp_Counters_Control_Tunnels) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range tunnels.Tunnel {
        children[tunnels.Tunnel[i].GetSegmentPath()] = &tunnels.Tunnel[i]
    }
    return children
}

func (tunnels *L2Tp_Counters_Control_Tunnels) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (tunnels *L2Tp_Counters_Control_Tunnels) GetBundleName() string { return "cisco_ios_xr" }

func (tunnels *L2Tp_Counters_Control_Tunnels) GetYangName() string { return "tunnels" }

func (tunnels *L2Tp_Counters_Control_Tunnels) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tunnels *L2Tp_Counters_Control_Tunnels) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tunnels *L2Tp_Counters_Control_Tunnels) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tunnels *L2Tp_Counters_Control_Tunnels) SetParent(parent types.Entity) { tunnels.parent = parent }

func (tunnels *L2Tp_Counters_Control_Tunnels) GetParent() types.Entity { return tunnels.parent }

func (tunnels *L2Tp_Counters_Control_Tunnels) GetParentYangName() string { return "control" }

// L2Tp_Counters_Control_Tunnels_Tunnel
// L2TP tunnel control message counters
type L2Tp_Counters_Control_Tunnels_Tunnel struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. L2TP tunnel ID. The type is interface{} with
    // range: -2147483648..2147483647.
    TunnelId interface{}

    // L2TP control message local and remote addresses.
    Brief L2Tp_Counters_Control_Tunnels_Tunnel_Brief

    // Global data.
    Global L2Tp_Counters_Control_Tunnels_Tunnel_Global
}

func (tunnel *L2Tp_Counters_Control_Tunnels_Tunnel) GetFilter() yfilter.YFilter { return tunnel.YFilter }

func (tunnel *L2Tp_Counters_Control_Tunnels_Tunnel) SetFilter(yf yfilter.YFilter) { tunnel.YFilter = yf }

func (tunnel *L2Tp_Counters_Control_Tunnels_Tunnel) GetGoName(yname string) string {
    if yname == "tunnel-id" { return "TunnelId" }
    if yname == "brief" { return "Brief" }
    if yname == "global" { return "Global" }
    return ""
}

func (tunnel *L2Tp_Counters_Control_Tunnels_Tunnel) GetSegmentPath() string {
    return "tunnel" + "[tunnel-id='" + fmt.Sprintf("%v", tunnel.TunnelId) + "']"
}

func (tunnel *L2Tp_Counters_Control_Tunnels_Tunnel) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "brief" {
        return &tunnel.Brief
    }
    if childYangName == "global" {
        return &tunnel.Global
    }
    return nil
}

func (tunnel *L2Tp_Counters_Control_Tunnels_Tunnel) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["brief"] = &tunnel.Brief
    children["global"] = &tunnel.Global
    return children
}

func (tunnel *L2Tp_Counters_Control_Tunnels_Tunnel) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tunnel-id"] = tunnel.TunnelId
    return leafs
}

func (tunnel *L2Tp_Counters_Control_Tunnels_Tunnel) GetBundleName() string { return "cisco_ios_xr" }

func (tunnel *L2Tp_Counters_Control_Tunnels_Tunnel) GetYangName() string { return "tunnel" }

func (tunnel *L2Tp_Counters_Control_Tunnels_Tunnel) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tunnel *L2Tp_Counters_Control_Tunnels_Tunnel) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tunnel *L2Tp_Counters_Control_Tunnels_Tunnel) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tunnel *L2Tp_Counters_Control_Tunnels_Tunnel) SetParent(parent types.Entity) { tunnel.parent = parent }

func (tunnel *L2Tp_Counters_Control_Tunnels_Tunnel) GetParent() types.Entity { return tunnel.parent }

func (tunnel *L2Tp_Counters_Control_Tunnels_Tunnel) GetParentYangName() string { return "tunnels" }

// L2Tp_Counters_Control_Tunnels_Tunnel_Brief
// L2TP control message local and remote addresses
type L2Tp_Counters_Control_Tunnels_Tunnel_Brief struct {
    parent types.Entity
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

func (brief *L2Tp_Counters_Control_Tunnels_Tunnel_Brief) GetFilter() yfilter.YFilter { return brief.YFilter }

func (brief *L2Tp_Counters_Control_Tunnels_Tunnel_Brief) SetFilter(yf yfilter.YFilter) { brief.YFilter = yf }

func (brief *L2Tp_Counters_Control_Tunnels_Tunnel_Brief) GetGoName(yname string) string {
    if yname == "remote-tunnel-id" { return "RemoteTunnelId" }
    if yname == "local-address" { return "LocalAddress" }
    if yname == "remote-address" { return "RemoteAddress" }
    return ""
}

func (brief *L2Tp_Counters_Control_Tunnels_Tunnel_Brief) GetSegmentPath() string {
    return "brief"
}

func (brief *L2Tp_Counters_Control_Tunnels_Tunnel_Brief) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (brief *L2Tp_Counters_Control_Tunnels_Tunnel_Brief) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (brief *L2Tp_Counters_Control_Tunnels_Tunnel_Brief) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["remote-tunnel-id"] = brief.RemoteTunnelId
    leafs["local-address"] = brief.LocalAddress
    leafs["remote-address"] = brief.RemoteAddress
    return leafs
}

func (brief *L2Tp_Counters_Control_Tunnels_Tunnel_Brief) GetBundleName() string { return "cisco_ios_xr" }

func (brief *L2Tp_Counters_Control_Tunnels_Tunnel_Brief) GetYangName() string { return "brief" }

func (brief *L2Tp_Counters_Control_Tunnels_Tunnel_Brief) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (brief *L2Tp_Counters_Control_Tunnels_Tunnel_Brief) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (brief *L2Tp_Counters_Control_Tunnels_Tunnel_Brief) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (brief *L2Tp_Counters_Control_Tunnels_Tunnel_Brief) SetParent(parent types.Entity) { brief.parent = parent }

func (brief *L2Tp_Counters_Control_Tunnels_Tunnel_Brief) GetParent() types.Entity { return brief.parent }

func (brief *L2Tp_Counters_Control_Tunnels_Tunnel_Brief) GetParentYangName() string { return "tunnel" }

// L2Tp_Counters_Control_Tunnels_Tunnel_Global
// Global data
type L2Tp_Counters_Control_Tunnels_Tunnel_Global struct {
    parent types.Entity
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

func (global *L2Tp_Counters_Control_Tunnels_Tunnel_Global) GetFilter() yfilter.YFilter { return global.YFilter }

func (global *L2Tp_Counters_Control_Tunnels_Tunnel_Global) SetFilter(yf yfilter.YFilter) { global.YFilter = yf }

func (global *L2Tp_Counters_Control_Tunnels_Tunnel_Global) GetGoName(yname string) string {
    if yname == "total-transmit" { return "TotalTransmit" }
    if yname == "total-retransmit" { return "TotalRetransmit" }
    if yname == "total-received" { return "TotalReceived" }
    if yname == "total-drop" { return "TotalDrop" }
    if yname == "transmit" { return "Transmit" }
    if yname == "retransmit" { return "Retransmit" }
    if yname == "received" { return "Received" }
    if yname == "drop" { return "Drop" }
    return ""
}

func (global *L2Tp_Counters_Control_Tunnels_Tunnel_Global) GetSegmentPath() string {
    return "global"
}

func (global *L2Tp_Counters_Control_Tunnels_Tunnel_Global) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "transmit" {
        return &global.Transmit
    }
    if childYangName == "retransmit" {
        return &global.Retransmit
    }
    if childYangName == "received" {
        return &global.Received
    }
    if childYangName == "drop" {
        return &global.Drop
    }
    return nil
}

func (global *L2Tp_Counters_Control_Tunnels_Tunnel_Global) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["transmit"] = &global.Transmit
    children["retransmit"] = &global.Retransmit
    children["received"] = &global.Received
    children["drop"] = &global.Drop
    return children
}

func (global *L2Tp_Counters_Control_Tunnels_Tunnel_Global) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["total-transmit"] = global.TotalTransmit
    leafs["total-retransmit"] = global.TotalRetransmit
    leafs["total-received"] = global.TotalReceived
    leafs["total-drop"] = global.TotalDrop
    return leafs
}

func (global *L2Tp_Counters_Control_Tunnels_Tunnel_Global) GetBundleName() string { return "cisco_ios_xr" }

func (global *L2Tp_Counters_Control_Tunnels_Tunnel_Global) GetYangName() string { return "global" }

func (global *L2Tp_Counters_Control_Tunnels_Tunnel_Global) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (global *L2Tp_Counters_Control_Tunnels_Tunnel_Global) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (global *L2Tp_Counters_Control_Tunnels_Tunnel_Global) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (global *L2Tp_Counters_Control_Tunnels_Tunnel_Global) SetParent(parent types.Entity) { global.parent = parent }

func (global *L2Tp_Counters_Control_Tunnels_Tunnel_Global) GetParent() types.Entity { return global.parent }

func (global *L2Tp_Counters_Control_Tunnels_Tunnel_Global) GetParentYangName() string { return "tunnel" }

// L2Tp_Counters_Control_Tunnels_Tunnel_Global_Transmit
// Transmit data
type L2Tp_Counters_Control_Tunnels_Tunnel_Global_Transmit struct {
    parent types.Entity
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

func (transmit *L2Tp_Counters_Control_Tunnels_Tunnel_Global_Transmit) GetFilter() yfilter.YFilter { return transmit.YFilter }

func (transmit *L2Tp_Counters_Control_Tunnels_Tunnel_Global_Transmit) SetFilter(yf yfilter.YFilter) { transmit.YFilter = yf }

func (transmit *L2Tp_Counters_Control_Tunnels_Tunnel_Global_Transmit) GetGoName(yname string) string {
    if yname == "unknown-packets" { return "UnknownPackets" }
    if yname == "zero-length-body-packets" { return "ZeroLengthBodyPackets" }
    if yname == "start-control-connection-requests" { return "StartControlConnectionRequests" }
    if yname == "start-control-connection-replies" { return "StartControlConnectionReplies" }
    if yname == "start-control-connection-notifications" { return "StartControlConnectionNotifications" }
    if yname == "stop-control-connection-notifications" { return "StopControlConnectionNotifications" }
    if yname == "hello-packets" { return "HelloPackets" }
    if yname == "outgoing-call-requests" { return "OutgoingCallRequests" }
    if yname == "outgoing-call-replies" { return "OutgoingCallReplies" }
    if yname == "outgoing-call-connected-packets" { return "OutgoingCallConnectedPackets" }
    if yname == "incoming-call-requests" { return "IncomingCallRequests" }
    if yname == "incoming-call-replies" { return "IncomingCallReplies" }
    if yname == "incoming-call-connected-packets" { return "IncomingCallConnectedPackets" }
    if yname == "call-disconnect-notify-packets" { return "CallDisconnectNotifyPackets" }
    if yname == "wan-error-notify-packets" { return "WanErrorNotifyPackets" }
    if yname == "set-link-info-packets" { return "SetLinkInfoPackets" }
    if yname == "service-relay-requests" { return "ServiceRelayRequests" }
    if yname == "service-relay-replies" { return "ServiceRelayReplies" }
    if yname == "acknowledgement-packets" { return "AcknowledgementPackets" }
    return ""
}

func (transmit *L2Tp_Counters_Control_Tunnels_Tunnel_Global_Transmit) GetSegmentPath() string {
    return "transmit"
}

func (transmit *L2Tp_Counters_Control_Tunnels_Tunnel_Global_Transmit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (transmit *L2Tp_Counters_Control_Tunnels_Tunnel_Global_Transmit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (transmit *L2Tp_Counters_Control_Tunnels_Tunnel_Global_Transmit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["unknown-packets"] = transmit.UnknownPackets
    leafs["zero-length-body-packets"] = transmit.ZeroLengthBodyPackets
    leafs["start-control-connection-requests"] = transmit.StartControlConnectionRequests
    leafs["start-control-connection-replies"] = transmit.StartControlConnectionReplies
    leafs["start-control-connection-notifications"] = transmit.StartControlConnectionNotifications
    leafs["stop-control-connection-notifications"] = transmit.StopControlConnectionNotifications
    leafs["hello-packets"] = transmit.HelloPackets
    leafs["outgoing-call-requests"] = transmit.OutgoingCallRequests
    leafs["outgoing-call-replies"] = transmit.OutgoingCallReplies
    leafs["outgoing-call-connected-packets"] = transmit.OutgoingCallConnectedPackets
    leafs["incoming-call-requests"] = transmit.IncomingCallRequests
    leafs["incoming-call-replies"] = transmit.IncomingCallReplies
    leafs["incoming-call-connected-packets"] = transmit.IncomingCallConnectedPackets
    leafs["call-disconnect-notify-packets"] = transmit.CallDisconnectNotifyPackets
    leafs["wan-error-notify-packets"] = transmit.WanErrorNotifyPackets
    leafs["set-link-info-packets"] = transmit.SetLinkInfoPackets
    leafs["service-relay-requests"] = transmit.ServiceRelayRequests
    leafs["service-relay-replies"] = transmit.ServiceRelayReplies
    leafs["acknowledgement-packets"] = transmit.AcknowledgementPackets
    return leafs
}

func (transmit *L2Tp_Counters_Control_Tunnels_Tunnel_Global_Transmit) GetBundleName() string { return "cisco_ios_xr" }

func (transmit *L2Tp_Counters_Control_Tunnels_Tunnel_Global_Transmit) GetYangName() string { return "transmit" }

func (transmit *L2Tp_Counters_Control_Tunnels_Tunnel_Global_Transmit) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (transmit *L2Tp_Counters_Control_Tunnels_Tunnel_Global_Transmit) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (transmit *L2Tp_Counters_Control_Tunnels_Tunnel_Global_Transmit) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (transmit *L2Tp_Counters_Control_Tunnels_Tunnel_Global_Transmit) SetParent(parent types.Entity) { transmit.parent = parent }

func (transmit *L2Tp_Counters_Control_Tunnels_Tunnel_Global_Transmit) GetParent() types.Entity { return transmit.parent }

func (transmit *L2Tp_Counters_Control_Tunnels_Tunnel_Global_Transmit) GetParentYangName() string { return "global" }

// L2Tp_Counters_Control_Tunnels_Tunnel_Global_Retransmit
// Re transmit data
type L2Tp_Counters_Control_Tunnels_Tunnel_Global_Retransmit struct {
    parent types.Entity
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

func (retransmit *L2Tp_Counters_Control_Tunnels_Tunnel_Global_Retransmit) GetFilter() yfilter.YFilter { return retransmit.YFilter }

func (retransmit *L2Tp_Counters_Control_Tunnels_Tunnel_Global_Retransmit) SetFilter(yf yfilter.YFilter) { retransmit.YFilter = yf }

func (retransmit *L2Tp_Counters_Control_Tunnels_Tunnel_Global_Retransmit) GetGoName(yname string) string {
    if yname == "unknown-packets" { return "UnknownPackets" }
    if yname == "zero-length-body-packets" { return "ZeroLengthBodyPackets" }
    if yname == "start-control-connection-requests" { return "StartControlConnectionRequests" }
    if yname == "start-control-connection-replies" { return "StartControlConnectionReplies" }
    if yname == "start-control-connection-notifications" { return "StartControlConnectionNotifications" }
    if yname == "stop-control-connection-notifications" { return "StopControlConnectionNotifications" }
    if yname == "hello-packets" { return "HelloPackets" }
    if yname == "outgoing-call-requests" { return "OutgoingCallRequests" }
    if yname == "outgoing-call-replies" { return "OutgoingCallReplies" }
    if yname == "outgoing-call-connected-packets" { return "OutgoingCallConnectedPackets" }
    if yname == "incoming-call-requests" { return "IncomingCallRequests" }
    if yname == "incoming-call-replies" { return "IncomingCallReplies" }
    if yname == "incoming-call-connected-packets" { return "IncomingCallConnectedPackets" }
    if yname == "call-disconnect-notify-packets" { return "CallDisconnectNotifyPackets" }
    if yname == "wan-error-notify-packets" { return "WanErrorNotifyPackets" }
    if yname == "set-link-info-packets" { return "SetLinkInfoPackets" }
    if yname == "service-relay-requests" { return "ServiceRelayRequests" }
    if yname == "service-relay-replies" { return "ServiceRelayReplies" }
    if yname == "acknowledgement-packets" { return "AcknowledgementPackets" }
    return ""
}

func (retransmit *L2Tp_Counters_Control_Tunnels_Tunnel_Global_Retransmit) GetSegmentPath() string {
    return "retransmit"
}

func (retransmit *L2Tp_Counters_Control_Tunnels_Tunnel_Global_Retransmit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (retransmit *L2Tp_Counters_Control_Tunnels_Tunnel_Global_Retransmit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (retransmit *L2Tp_Counters_Control_Tunnels_Tunnel_Global_Retransmit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["unknown-packets"] = retransmit.UnknownPackets
    leafs["zero-length-body-packets"] = retransmit.ZeroLengthBodyPackets
    leafs["start-control-connection-requests"] = retransmit.StartControlConnectionRequests
    leafs["start-control-connection-replies"] = retransmit.StartControlConnectionReplies
    leafs["start-control-connection-notifications"] = retransmit.StartControlConnectionNotifications
    leafs["stop-control-connection-notifications"] = retransmit.StopControlConnectionNotifications
    leafs["hello-packets"] = retransmit.HelloPackets
    leafs["outgoing-call-requests"] = retransmit.OutgoingCallRequests
    leafs["outgoing-call-replies"] = retransmit.OutgoingCallReplies
    leafs["outgoing-call-connected-packets"] = retransmit.OutgoingCallConnectedPackets
    leafs["incoming-call-requests"] = retransmit.IncomingCallRequests
    leafs["incoming-call-replies"] = retransmit.IncomingCallReplies
    leafs["incoming-call-connected-packets"] = retransmit.IncomingCallConnectedPackets
    leafs["call-disconnect-notify-packets"] = retransmit.CallDisconnectNotifyPackets
    leafs["wan-error-notify-packets"] = retransmit.WanErrorNotifyPackets
    leafs["set-link-info-packets"] = retransmit.SetLinkInfoPackets
    leafs["service-relay-requests"] = retransmit.ServiceRelayRequests
    leafs["service-relay-replies"] = retransmit.ServiceRelayReplies
    leafs["acknowledgement-packets"] = retransmit.AcknowledgementPackets
    return leafs
}

func (retransmit *L2Tp_Counters_Control_Tunnels_Tunnel_Global_Retransmit) GetBundleName() string { return "cisco_ios_xr" }

func (retransmit *L2Tp_Counters_Control_Tunnels_Tunnel_Global_Retransmit) GetYangName() string { return "retransmit" }

func (retransmit *L2Tp_Counters_Control_Tunnels_Tunnel_Global_Retransmit) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (retransmit *L2Tp_Counters_Control_Tunnels_Tunnel_Global_Retransmit) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (retransmit *L2Tp_Counters_Control_Tunnels_Tunnel_Global_Retransmit) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (retransmit *L2Tp_Counters_Control_Tunnels_Tunnel_Global_Retransmit) SetParent(parent types.Entity) { retransmit.parent = parent }

func (retransmit *L2Tp_Counters_Control_Tunnels_Tunnel_Global_Retransmit) GetParent() types.Entity { return retransmit.parent }

func (retransmit *L2Tp_Counters_Control_Tunnels_Tunnel_Global_Retransmit) GetParentYangName() string { return "global" }

// L2Tp_Counters_Control_Tunnels_Tunnel_Global_Received
// Received data
type L2Tp_Counters_Control_Tunnels_Tunnel_Global_Received struct {
    parent types.Entity
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

func (received *L2Tp_Counters_Control_Tunnels_Tunnel_Global_Received) GetFilter() yfilter.YFilter { return received.YFilter }

func (received *L2Tp_Counters_Control_Tunnels_Tunnel_Global_Received) SetFilter(yf yfilter.YFilter) { received.YFilter = yf }

func (received *L2Tp_Counters_Control_Tunnels_Tunnel_Global_Received) GetGoName(yname string) string {
    if yname == "unknown-packets" { return "UnknownPackets" }
    if yname == "zero-length-body-packets" { return "ZeroLengthBodyPackets" }
    if yname == "start-control-connection-requests" { return "StartControlConnectionRequests" }
    if yname == "start-control-connection-replies" { return "StartControlConnectionReplies" }
    if yname == "start-control-connection-notifications" { return "StartControlConnectionNotifications" }
    if yname == "stop-control-connection-notifications" { return "StopControlConnectionNotifications" }
    if yname == "hello-packets" { return "HelloPackets" }
    if yname == "outgoing-call-requests" { return "OutgoingCallRequests" }
    if yname == "outgoing-call-replies" { return "OutgoingCallReplies" }
    if yname == "outgoing-call-connected-packets" { return "OutgoingCallConnectedPackets" }
    if yname == "incoming-call-requests" { return "IncomingCallRequests" }
    if yname == "incoming-call-replies" { return "IncomingCallReplies" }
    if yname == "incoming-call-connected-packets" { return "IncomingCallConnectedPackets" }
    if yname == "call-disconnect-notify-packets" { return "CallDisconnectNotifyPackets" }
    if yname == "wan-error-notify-packets" { return "WanErrorNotifyPackets" }
    if yname == "set-link-info-packets" { return "SetLinkInfoPackets" }
    if yname == "service-relay-requests" { return "ServiceRelayRequests" }
    if yname == "service-relay-replies" { return "ServiceRelayReplies" }
    if yname == "acknowledgement-packets" { return "AcknowledgementPackets" }
    return ""
}

func (received *L2Tp_Counters_Control_Tunnels_Tunnel_Global_Received) GetSegmentPath() string {
    return "received"
}

func (received *L2Tp_Counters_Control_Tunnels_Tunnel_Global_Received) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (received *L2Tp_Counters_Control_Tunnels_Tunnel_Global_Received) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (received *L2Tp_Counters_Control_Tunnels_Tunnel_Global_Received) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["unknown-packets"] = received.UnknownPackets
    leafs["zero-length-body-packets"] = received.ZeroLengthBodyPackets
    leafs["start-control-connection-requests"] = received.StartControlConnectionRequests
    leafs["start-control-connection-replies"] = received.StartControlConnectionReplies
    leafs["start-control-connection-notifications"] = received.StartControlConnectionNotifications
    leafs["stop-control-connection-notifications"] = received.StopControlConnectionNotifications
    leafs["hello-packets"] = received.HelloPackets
    leafs["outgoing-call-requests"] = received.OutgoingCallRequests
    leafs["outgoing-call-replies"] = received.OutgoingCallReplies
    leafs["outgoing-call-connected-packets"] = received.OutgoingCallConnectedPackets
    leafs["incoming-call-requests"] = received.IncomingCallRequests
    leafs["incoming-call-replies"] = received.IncomingCallReplies
    leafs["incoming-call-connected-packets"] = received.IncomingCallConnectedPackets
    leafs["call-disconnect-notify-packets"] = received.CallDisconnectNotifyPackets
    leafs["wan-error-notify-packets"] = received.WanErrorNotifyPackets
    leafs["set-link-info-packets"] = received.SetLinkInfoPackets
    leafs["service-relay-requests"] = received.ServiceRelayRequests
    leafs["service-relay-replies"] = received.ServiceRelayReplies
    leafs["acknowledgement-packets"] = received.AcknowledgementPackets
    return leafs
}

func (received *L2Tp_Counters_Control_Tunnels_Tunnel_Global_Received) GetBundleName() string { return "cisco_ios_xr" }

func (received *L2Tp_Counters_Control_Tunnels_Tunnel_Global_Received) GetYangName() string { return "received" }

func (received *L2Tp_Counters_Control_Tunnels_Tunnel_Global_Received) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (received *L2Tp_Counters_Control_Tunnels_Tunnel_Global_Received) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (received *L2Tp_Counters_Control_Tunnels_Tunnel_Global_Received) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (received *L2Tp_Counters_Control_Tunnels_Tunnel_Global_Received) SetParent(parent types.Entity) { received.parent = parent }

func (received *L2Tp_Counters_Control_Tunnels_Tunnel_Global_Received) GetParent() types.Entity { return received.parent }

func (received *L2Tp_Counters_Control_Tunnels_Tunnel_Global_Received) GetParentYangName() string { return "global" }

// L2Tp_Counters_Control_Tunnels_Tunnel_Global_Drop
// Drop data
type L2Tp_Counters_Control_Tunnels_Tunnel_Global_Drop struct {
    parent types.Entity
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

func (drop *L2Tp_Counters_Control_Tunnels_Tunnel_Global_Drop) GetFilter() yfilter.YFilter { return drop.YFilter }

func (drop *L2Tp_Counters_Control_Tunnels_Tunnel_Global_Drop) SetFilter(yf yfilter.YFilter) { drop.YFilter = yf }

func (drop *L2Tp_Counters_Control_Tunnels_Tunnel_Global_Drop) GetGoName(yname string) string {
    if yname == "unknown-packets" { return "UnknownPackets" }
    if yname == "zero-length-body-packets" { return "ZeroLengthBodyPackets" }
    if yname == "start-control-connection-requests" { return "StartControlConnectionRequests" }
    if yname == "start-control-connection-replies" { return "StartControlConnectionReplies" }
    if yname == "start-control-connection-notifications" { return "StartControlConnectionNotifications" }
    if yname == "stop-control-connection-notifications" { return "StopControlConnectionNotifications" }
    if yname == "hello-packets" { return "HelloPackets" }
    if yname == "outgoing-call-requests" { return "OutgoingCallRequests" }
    if yname == "outgoing-call-replies" { return "OutgoingCallReplies" }
    if yname == "outgoing-call-connected-packets" { return "OutgoingCallConnectedPackets" }
    if yname == "incoming-call-requests" { return "IncomingCallRequests" }
    if yname == "incoming-call-replies" { return "IncomingCallReplies" }
    if yname == "incoming-call-connected-packets" { return "IncomingCallConnectedPackets" }
    if yname == "call-disconnect-notify-packets" { return "CallDisconnectNotifyPackets" }
    if yname == "wan-error-notify-packets" { return "WanErrorNotifyPackets" }
    if yname == "set-link-info-packets" { return "SetLinkInfoPackets" }
    if yname == "service-relay-requests" { return "ServiceRelayRequests" }
    if yname == "service-relay-replies" { return "ServiceRelayReplies" }
    if yname == "acknowledgement-packets" { return "AcknowledgementPackets" }
    return ""
}

func (drop *L2Tp_Counters_Control_Tunnels_Tunnel_Global_Drop) GetSegmentPath() string {
    return "drop"
}

func (drop *L2Tp_Counters_Control_Tunnels_Tunnel_Global_Drop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (drop *L2Tp_Counters_Control_Tunnels_Tunnel_Global_Drop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (drop *L2Tp_Counters_Control_Tunnels_Tunnel_Global_Drop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["unknown-packets"] = drop.UnknownPackets
    leafs["zero-length-body-packets"] = drop.ZeroLengthBodyPackets
    leafs["start-control-connection-requests"] = drop.StartControlConnectionRequests
    leafs["start-control-connection-replies"] = drop.StartControlConnectionReplies
    leafs["start-control-connection-notifications"] = drop.StartControlConnectionNotifications
    leafs["stop-control-connection-notifications"] = drop.StopControlConnectionNotifications
    leafs["hello-packets"] = drop.HelloPackets
    leafs["outgoing-call-requests"] = drop.OutgoingCallRequests
    leafs["outgoing-call-replies"] = drop.OutgoingCallReplies
    leafs["outgoing-call-connected-packets"] = drop.OutgoingCallConnectedPackets
    leafs["incoming-call-requests"] = drop.IncomingCallRequests
    leafs["incoming-call-replies"] = drop.IncomingCallReplies
    leafs["incoming-call-connected-packets"] = drop.IncomingCallConnectedPackets
    leafs["call-disconnect-notify-packets"] = drop.CallDisconnectNotifyPackets
    leafs["wan-error-notify-packets"] = drop.WanErrorNotifyPackets
    leafs["set-link-info-packets"] = drop.SetLinkInfoPackets
    leafs["service-relay-requests"] = drop.ServiceRelayRequests
    leafs["service-relay-replies"] = drop.ServiceRelayReplies
    leafs["acknowledgement-packets"] = drop.AcknowledgementPackets
    return leafs
}

func (drop *L2Tp_Counters_Control_Tunnels_Tunnel_Global_Drop) GetBundleName() string { return "cisco_ios_xr" }

func (drop *L2Tp_Counters_Control_Tunnels_Tunnel_Global_Drop) GetYangName() string { return "drop" }

func (drop *L2Tp_Counters_Control_Tunnels_Tunnel_Global_Drop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (drop *L2Tp_Counters_Control_Tunnels_Tunnel_Global_Drop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (drop *L2Tp_Counters_Control_Tunnels_Tunnel_Global_Drop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (drop *L2Tp_Counters_Control_Tunnels_Tunnel_Global_Drop) SetParent(parent types.Entity) { drop.parent = parent }

func (drop *L2Tp_Counters_Control_Tunnels_Tunnel_Global_Drop) GetParent() types.Entity { return drop.parent }

func (drop *L2Tp_Counters_Control_Tunnels_Tunnel_Global_Drop) GetParentYangName() string { return "global" }

// L2Tp_TunnelConfigurations
// List of tunnel IDs
type L2Tp_TunnelConfigurations struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // L2TP tunnel information. The type is slice of
    // L2Tp_TunnelConfigurations_TunnelConfiguration.
    TunnelConfiguration []L2Tp_TunnelConfigurations_TunnelConfiguration
}

func (tunnelConfigurations *L2Tp_TunnelConfigurations) GetFilter() yfilter.YFilter { return tunnelConfigurations.YFilter }

func (tunnelConfigurations *L2Tp_TunnelConfigurations) SetFilter(yf yfilter.YFilter) { tunnelConfigurations.YFilter = yf }

func (tunnelConfigurations *L2Tp_TunnelConfigurations) GetGoName(yname string) string {
    if yname == "tunnel-configuration" { return "TunnelConfiguration" }
    return ""
}

func (tunnelConfigurations *L2Tp_TunnelConfigurations) GetSegmentPath() string {
    return "tunnel-configurations"
}

func (tunnelConfigurations *L2Tp_TunnelConfigurations) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tunnel-configuration" {
        for _, c := range tunnelConfigurations.TunnelConfiguration {
            if tunnelConfigurations.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := L2Tp_TunnelConfigurations_TunnelConfiguration{}
        tunnelConfigurations.TunnelConfiguration = append(tunnelConfigurations.TunnelConfiguration, child)
        return &tunnelConfigurations.TunnelConfiguration[len(tunnelConfigurations.TunnelConfiguration)-1]
    }
    return nil
}

func (tunnelConfigurations *L2Tp_TunnelConfigurations) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range tunnelConfigurations.TunnelConfiguration {
        children[tunnelConfigurations.TunnelConfiguration[i].GetSegmentPath()] = &tunnelConfigurations.TunnelConfiguration[i]
    }
    return children
}

func (tunnelConfigurations *L2Tp_TunnelConfigurations) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (tunnelConfigurations *L2Tp_TunnelConfigurations) GetBundleName() string { return "cisco_ios_xr" }

func (tunnelConfigurations *L2Tp_TunnelConfigurations) GetYangName() string { return "tunnel-configurations" }

func (tunnelConfigurations *L2Tp_TunnelConfigurations) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tunnelConfigurations *L2Tp_TunnelConfigurations) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tunnelConfigurations *L2Tp_TunnelConfigurations) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tunnelConfigurations *L2Tp_TunnelConfigurations) SetParent(parent types.Entity) { tunnelConfigurations.parent = parent }

func (tunnelConfigurations *L2Tp_TunnelConfigurations) GetParent() types.Entity { return tunnelConfigurations.parent }

func (tunnelConfigurations *L2Tp_TunnelConfigurations) GetParentYangName() string { return "l2tp" }

// L2Tp_TunnelConfigurations_TunnelConfiguration
// L2TP tunnel information
type L2Tp_TunnelConfigurations_TunnelConfiguration struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Local tunnel ID. The type is interface{} with
    // range: -2147483648..2147483647.
    LocalTunnelId interface{}

    // Remote tunnel ID. The type is interface{} with range: 0..4294967295.
    RemoteTunnelId interface{}

    // L2Tp class data.
    L2TpClass L2Tp_TunnelConfigurations_TunnelConfiguration_L2TpClass
}

func (tunnelConfiguration *L2Tp_TunnelConfigurations_TunnelConfiguration) GetFilter() yfilter.YFilter { return tunnelConfiguration.YFilter }

func (tunnelConfiguration *L2Tp_TunnelConfigurations_TunnelConfiguration) SetFilter(yf yfilter.YFilter) { tunnelConfiguration.YFilter = yf }

func (tunnelConfiguration *L2Tp_TunnelConfigurations_TunnelConfiguration) GetGoName(yname string) string {
    if yname == "local-tunnel-id" { return "LocalTunnelId" }
    if yname == "remote-tunnel-id" { return "RemoteTunnelId" }
    if yname == "l2tp-class" { return "L2TpClass" }
    return ""
}

func (tunnelConfiguration *L2Tp_TunnelConfigurations_TunnelConfiguration) GetSegmentPath() string {
    return "tunnel-configuration" + "[local-tunnel-id='" + fmt.Sprintf("%v", tunnelConfiguration.LocalTunnelId) + "']"
}

func (tunnelConfiguration *L2Tp_TunnelConfigurations_TunnelConfiguration) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "l2tp-class" {
        return &tunnelConfiguration.L2TpClass
    }
    return nil
}

func (tunnelConfiguration *L2Tp_TunnelConfigurations_TunnelConfiguration) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["l2tp-class"] = &tunnelConfiguration.L2TpClass
    return children
}

func (tunnelConfiguration *L2Tp_TunnelConfigurations_TunnelConfiguration) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["local-tunnel-id"] = tunnelConfiguration.LocalTunnelId
    leafs["remote-tunnel-id"] = tunnelConfiguration.RemoteTunnelId
    return leafs
}

func (tunnelConfiguration *L2Tp_TunnelConfigurations_TunnelConfiguration) GetBundleName() string { return "cisco_ios_xr" }

func (tunnelConfiguration *L2Tp_TunnelConfigurations_TunnelConfiguration) GetYangName() string { return "tunnel-configuration" }

func (tunnelConfiguration *L2Tp_TunnelConfigurations_TunnelConfiguration) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tunnelConfiguration *L2Tp_TunnelConfigurations_TunnelConfiguration) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tunnelConfiguration *L2Tp_TunnelConfigurations_TunnelConfiguration) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tunnelConfiguration *L2Tp_TunnelConfigurations_TunnelConfiguration) SetParent(parent types.Entity) { tunnelConfiguration.parent = parent }

func (tunnelConfiguration *L2Tp_TunnelConfigurations_TunnelConfiguration) GetParent() types.Entity { return tunnelConfiguration.parent }

func (tunnelConfiguration *L2Tp_TunnelConfigurations_TunnelConfiguration) GetParentYangName() string { return "tunnel-configurations" }

// L2Tp_TunnelConfigurations_TunnelConfiguration_L2TpClass
// L2Tp class data
type L2Tp_TunnelConfigurations_TunnelConfiguration_L2TpClass struct {
    parent types.Entity
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

func (l2TpClass *L2Tp_TunnelConfigurations_TunnelConfiguration_L2TpClass) GetFilter() yfilter.YFilter { return l2TpClass.YFilter }

func (l2TpClass *L2Tp_TunnelConfigurations_TunnelConfiguration_L2TpClass) SetFilter(yf yfilter.YFilter) { l2TpClass.YFilter = yf }

func (l2TpClass *L2Tp_TunnelConfigurations_TunnelConfiguration_L2TpClass) GetGoName(yname string) string {
    if yname == "ip-tos" { return "IpTos" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "receive-window-size" { return "ReceiveWindowSize" }
    if yname == "class-name-xr" { return "ClassNameXr" }
    if yname == "digest-hash" { return "DigestHash" }
    if yname == "password" { return "Password" }
    if yname == "encoded-password" { return "EncodedPassword" }
    if yname == "host-name" { return "HostName" }
    if yname == "accounting-method-list" { return "AccountingMethodList" }
    if yname == "hello-timeout" { return "HelloTimeout" }
    if yname == "setup-timeout" { return "SetupTimeout" }
    if yname == "retransmit-minimum-timeout" { return "RetransmitMinimumTimeout" }
    if yname == "retransmit-maximum-timeout" { return "RetransmitMaximumTimeout" }
    if yname == "initial-retransmit-minimum-timeout" { return "InitialRetransmitMinimumTimeout" }
    if yname == "initial-retransmit-maximum-timeout" { return "InitialRetransmitMaximumTimeout" }
    if yname == "timeout-no-user" { return "TimeoutNoUser" }
    if yname == "retransmit-retries" { return "RetransmitRetries" }
    if yname == "initial-retransmit-retries" { return "InitialRetransmitRetries" }
    if yname == "is-authentication-enabled" { return "IsAuthenticationEnabled" }
    if yname == "is-hidden" { return "IsHidden" }
    if yname == "is-digest-enabled" { return "IsDigestEnabled" }
    if yname == "is-digest-check-enabled" { return "IsDigestCheckEnabled" }
    if yname == "is-congestion-control-enabled" { return "IsCongestionControlEnabled" }
    if yname == "is-peer-address-checked" { return "IsPeerAddressChecked" }
    return ""
}

func (l2TpClass *L2Tp_TunnelConfigurations_TunnelConfiguration_L2TpClass) GetSegmentPath() string {
    return "l2tp-class"
}

func (l2TpClass *L2Tp_TunnelConfigurations_TunnelConfiguration_L2TpClass) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (l2TpClass *L2Tp_TunnelConfigurations_TunnelConfiguration_L2TpClass) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (l2TpClass *L2Tp_TunnelConfigurations_TunnelConfiguration_L2TpClass) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip-tos"] = l2TpClass.IpTos
    leafs["vrf-name"] = l2TpClass.VrfName
    leafs["receive-window-size"] = l2TpClass.ReceiveWindowSize
    leafs["class-name-xr"] = l2TpClass.ClassNameXr
    leafs["digest-hash"] = l2TpClass.DigestHash
    leafs["password"] = l2TpClass.Password
    leafs["encoded-password"] = l2TpClass.EncodedPassword
    leafs["host-name"] = l2TpClass.HostName
    leafs["accounting-method-list"] = l2TpClass.AccountingMethodList
    leafs["hello-timeout"] = l2TpClass.HelloTimeout
    leafs["setup-timeout"] = l2TpClass.SetupTimeout
    leafs["retransmit-minimum-timeout"] = l2TpClass.RetransmitMinimumTimeout
    leafs["retransmit-maximum-timeout"] = l2TpClass.RetransmitMaximumTimeout
    leafs["initial-retransmit-minimum-timeout"] = l2TpClass.InitialRetransmitMinimumTimeout
    leafs["initial-retransmit-maximum-timeout"] = l2TpClass.InitialRetransmitMaximumTimeout
    leafs["timeout-no-user"] = l2TpClass.TimeoutNoUser
    leafs["retransmit-retries"] = l2TpClass.RetransmitRetries
    leafs["initial-retransmit-retries"] = l2TpClass.InitialRetransmitRetries
    leafs["is-authentication-enabled"] = l2TpClass.IsAuthenticationEnabled
    leafs["is-hidden"] = l2TpClass.IsHidden
    leafs["is-digest-enabled"] = l2TpClass.IsDigestEnabled
    leafs["is-digest-check-enabled"] = l2TpClass.IsDigestCheckEnabled
    leafs["is-congestion-control-enabled"] = l2TpClass.IsCongestionControlEnabled
    leafs["is-peer-address-checked"] = l2TpClass.IsPeerAddressChecked
    return leafs
}

func (l2TpClass *L2Tp_TunnelConfigurations_TunnelConfiguration_L2TpClass) GetBundleName() string { return "cisco_ios_xr" }

func (l2TpClass *L2Tp_TunnelConfigurations_TunnelConfiguration_L2TpClass) GetYangName() string { return "l2tp-class" }

func (l2TpClass *L2Tp_TunnelConfigurations_TunnelConfiguration_L2TpClass) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (l2TpClass *L2Tp_TunnelConfigurations_TunnelConfiguration_L2TpClass) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (l2TpClass *L2Tp_TunnelConfigurations_TunnelConfiguration_L2TpClass) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (l2TpClass *L2Tp_TunnelConfigurations_TunnelConfiguration_L2TpClass) SetParent(parent types.Entity) { l2TpClass.parent = parent }

func (l2TpClass *L2Tp_TunnelConfigurations_TunnelConfiguration_L2TpClass) GetParent() types.Entity { return l2TpClass.parent }

func (l2TpClass *L2Tp_TunnelConfigurations_TunnelConfiguration_L2TpClass) GetParentYangName() string { return "tunnel-configuration" }

// L2Tp_CounterHistFail
// Failure events leading to disconnection
type L2Tp_CounterHistFail struct {
    parent types.Entity
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

    // timeout events by packet. The type is slice of interface{} with range:
    // 0..4294967295.
    PktTimeout []interface{}
}

func (counterHistFail *L2Tp_CounterHistFail) GetFilter() yfilter.YFilter { return counterHistFail.YFilter }

func (counterHistFail *L2Tp_CounterHistFail) SetFilter(yf yfilter.YFilter) { counterHistFail.YFilter = yf }

func (counterHistFail *L2Tp_CounterHistFail) GetGoName(yname string) string {
    if yname == "sess-down-tmout" { return "SessDownTmout" }
    if yname == "tx-counters" { return "TxCounters" }
    if yname == "rx-counters" { return "RxCounters" }
    if yname == "pkt-timeout" { return "PktTimeout" }
    return ""
}

func (counterHistFail *L2Tp_CounterHistFail) GetSegmentPath() string {
    return "counter-hist-fail"
}

func (counterHistFail *L2Tp_CounterHistFail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (counterHistFail *L2Tp_CounterHistFail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (counterHistFail *L2Tp_CounterHistFail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sess-down-tmout"] = counterHistFail.SessDownTmout
    leafs["tx-counters"] = counterHistFail.TxCounters
    leafs["rx-counters"] = counterHistFail.RxCounters
    leafs["pkt-timeout"] = counterHistFail.PktTimeout
    return leafs
}

func (counterHistFail *L2Tp_CounterHistFail) GetBundleName() string { return "cisco_ios_xr" }

func (counterHistFail *L2Tp_CounterHistFail) GetYangName() string { return "counter-hist-fail" }

func (counterHistFail *L2Tp_CounterHistFail) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (counterHistFail *L2Tp_CounterHistFail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (counterHistFail *L2Tp_CounterHistFail) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (counterHistFail *L2Tp_CounterHistFail) SetParent(parent types.Entity) { counterHistFail.parent = parent }

func (counterHistFail *L2Tp_CounterHistFail) GetParent() types.Entity { return counterHistFail.parent }

func (counterHistFail *L2Tp_CounterHistFail) GetParentYangName() string { return "l2tp" }

// L2Tp_Classes
// List of L2TP class names
type L2Tp_Classes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // L2TP class name. The type is slice of L2Tp_Classes_Class.
    Class []L2Tp_Classes_Class
}

func (classes *L2Tp_Classes) GetFilter() yfilter.YFilter { return classes.YFilter }

func (classes *L2Tp_Classes) SetFilter(yf yfilter.YFilter) { classes.YFilter = yf }

func (classes *L2Tp_Classes) GetGoName(yname string) string {
    if yname == "class" { return "Class" }
    return ""
}

func (classes *L2Tp_Classes) GetSegmentPath() string {
    return "classes"
}

func (classes *L2Tp_Classes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "class" {
        for _, c := range classes.Class {
            if classes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := L2Tp_Classes_Class{}
        classes.Class = append(classes.Class, child)
        return &classes.Class[len(classes.Class)-1]
    }
    return nil
}

func (classes *L2Tp_Classes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range classes.Class {
        children[classes.Class[i].GetSegmentPath()] = &classes.Class[i]
    }
    return children
}

func (classes *L2Tp_Classes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (classes *L2Tp_Classes) GetBundleName() string { return "cisco_ios_xr" }

func (classes *L2Tp_Classes) GetYangName() string { return "classes" }

func (classes *L2Tp_Classes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (classes *L2Tp_Classes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (classes *L2Tp_Classes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (classes *L2Tp_Classes) SetParent(parent types.Entity) { classes.parent = parent }

func (classes *L2Tp_Classes) GetParent() types.Entity { return classes.parent }

func (classes *L2Tp_Classes) GetParentYangName() string { return "l2tp" }

// L2Tp_Classes_Class
// L2TP class name
type L2Tp_Classes_Class struct {
    parent types.Entity
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

func (class *L2Tp_Classes_Class) GetFilter() yfilter.YFilter { return class.YFilter }

func (class *L2Tp_Classes_Class) SetFilter(yf yfilter.YFilter) { class.YFilter = yf }

func (class *L2Tp_Classes_Class) GetGoName(yname string) string {
    if yname == "class-name" { return "ClassName" }
    if yname == "ip-tos" { return "IpTos" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "receive-window-size" { return "ReceiveWindowSize" }
    if yname == "class-name-xr" { return "ClassNameXr" }
    if yname == "digest-hash" { return "DigestHash" }
    if yname == "password" { return "Password" }
    if yname == "encoded-password" { return "EncodedPassword" }
    if yname == "host-name" { return "HostName" }
    if yname == "accounting-method-list" { return "AccountingMethodList" }
    if yname == "hello-timeout" { return "HelloTimeout" }
    if yname == "setup-timeout" { return "SetupTimeout" }
    if yname == "retransmit-minimum-timeout" { return "RetransmitMinimumTimeout" }
    if yname == "retransmit-maximum-timeout" { return "RetransmitMaximumTimeout" }
    if yname == "initial-retransmit-minimum-timeout" { return "InitialRetransmitMinimumTimeout" }
    if yname == "initial-retransmit-maximum-timeout" { return "InitialRetransmitMaximumTimeout" }
    if yname == "timeout-no-user" { return "TimeoutNoUser" }
    if yname == "retransmit-retries" { return "RetransmitRetries" }
    if yname == "initial-retransmit-retries" { return "InitialRetransmitRetries" }
    if yname == "is-authentication-enabled" { return "IsAuthenticationEnabled" }
    if yname == "is-hidden" { return "IsHidden" }
    if yname == "is-digest-enabled" { return "IsDigestEnabled" }
    if yname == "is-digest-check-enabled" { return "IsDigestCheckEnabled" }
    if yname == "is-congestion-control-enabled" { return "IsCongestionControlEnabled" }
    if yname == "is-peer-address-checked" { return "IsPeerAddressChecked" }
    return ""
}

func (class *L2Tp_Classes_Class) GetSegmentPath() string {
    return "class" + "[class-name='" + fmt.Sprintf("%v", class.ClassName) + "']"
}

func (class *L2Tp_Classes_Class) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (class *L2Tp_Classes_Class) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (class *L2Tp_Classes_Class) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["class-name"] = class.ClassName
    leafs["ip-tos"] = class.IpTos
    leafs["vrf-name"] = class.VrfName
    leafs["receive-window-size"] = class.ReceiveWindowSize
    leafs["class-name-xr"] = class.ClassNameXr
    leafs["digest-hash"] = class.DigestHash
    leafs["password"] = class.Password
    leafs["encoded-password"] = class.EncodedPassword
    leafs["host-name"] = class.HostName
    leafs["accounting-method-list"] = class.AccountingMethodList
    leafs["hello-timeout"] = class.HelloTimeout
    leafs["setup-timeout"] = class.SetupTimeout
    leafs["retransmit-minimum-timeout"] = class.RetransmitMinimumTimeout
    leafs["retransmit-maximum-timeout"] = class.RetransmitMaximumTimeout
    leafs["initial-retransmit-minimum-timeout"] = class.InitialRetransmitMinimumTimeout
    leafs["initial-retransmit-maximum-timeout"] = class.InitialRetransmitMaximumTimeout
    leafs["timeout-no-user"] = class.TimeoutNoUser
    leafs["retransmit-retries"] = class.RetransmitRetries
    leafs["initial-retransmit-retries"] = class.InitialRetransmitRetries
    leafs["is-authentication-enabled"] = class.IsAuthenticationEnabled
    leafs["is-hidden"] = class.IsHidden
    leafs["is-digest-enabled"] = class.IsDigestEnabled
    leafs["is-digest-check-enabled"] = class.IsDigestCheckEnabled
    leafs["is-congestion-control-enabled"] = class.IsCongestionControlEnabled
    leafs["is-peer-address-checked"] = class.IsPeerAddressChecked
    return leafs
}

func (class *L2Tp_Classes_Class) GetBundleName() string { return "cisco_ios_xr" }

func (class *L2Tp_Classes_Class) GetYangName() string { return "class" }

func (class *L2Tp_Classes_Class) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (class *L2Tp_Classes_Class) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (class *L2Tp_Classes_Class) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (class *L2Tp_Classes_Class) SetParent(parent types.Entity) { class.parent = parent }

func (class *L2Tp_Classes_Class) GetParent() types.Entity { return class.parent }

func (class *L2Tp_Classes_Class) GetParentYangName() string { return "classes" }

// L2Tp_Tunnels
// List of tunnel IDs
type L2Tp_Tunnels struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // L2TP tunnel  information. The type is slice of L2Tp_Tunnels_Tunnel.
    Tunnel []L2Tp_Tunnels_Tunnel
}

func (tunnels *L2Tp_Tunnels) GetFilter() yfilter.YFilter { return tunnels.YFilter }

func (tunnels *L2Tp_Tunnels) SetFilter(yf yfilter.YFilter) { tunnels.YFilter = yf }

func (tunnels *L2Tp_Tunnels) GetGoName(yname string) string {
    if yname == "tunnel" { return "Tunnel" }
    return ""
}

func (tunnels *L2Tp_Tunnels) GetSegmentPath() string {
    return "tunnels"
}

func (tunnels *L2Tp_Tunnels) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tunnel" {
        for _, c := range tunnels.Tunnel {
            if tunnels.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := L2Tp_Tunnels_Tunnel{}
        tunnels.Tunnel = append(tunnels.Tunnel, child)
        return &tunnels.Tunnel[len(tunnels.Tunnel)-1]
    }
    return nil
}

func (tunnels *L2Tp_Tunnels) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range tunnels.Tunnel {
        children[tunnels.Tunnel[i].GetSegmentPath()] = &tunnels.Tunnel[i]
    }
    return children
}

func (tunnels *L2Tp_Tunnels) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (tunnels *L2Tp_Tunnels) GetBundleName() string { return "cisco_ios_xr" }

func (tunnels *L2Tp_Tunnels) GetYangName() string { return "tunnels" }

func (tunnels *L2Tp_Tunnels) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tunnels *L2Tp_Tunnels) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tunnels *L2Tp_Tunnels) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tunnels *L2Tp_Tunnels) SetParent(parent types.Entity) { tunnels.parent = parent }

func (tunnels *L2Tp_Tunnels) GetParent() types.Entity { return tunnels.parent }

func (tunnels *L2Tp_Tunnels) GetParentYangName() string { return "l2tp" }

// L2Tp_Tunnels_Tunnel
// L2TP tunnel  information
type L2Tp_Tunnels_Tunnel struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Local tunnel ID. The type is interface{} with
    // range: -2147483648..2147483647.
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

    // Retransmit time distribution in seconds. The type is slice of interface{}
    // with range: 0..65535. Units are second.
    RetransmitTime []interface{}
}

func (tunnel *L2Tp_Tunnels_Tunnel) GetFilter() yfilter.YFilter { return tunnel.YFilter }

func (tunnel *L2Tp_Tunnels_Tunnel) SetFilter(yf yfilter.YFilter) { tunnel.YFilter = yf }

func (tunnel *L2Tp_Tunnels_Tunnel) GetGoName(yname string) string {
    if yname == "local-tunnel-id" { return "LocalTunnelId" }
    if yname == "local-address" { return "LocalAddress" }
    if yname == "remote-address" { return "RemoteAddress" }
    if yname == "local-port" { return "LocalPort" }
    if yname == "remote-port" { return "RemotePort" }
    if yname == "protocol" { return "Protocol" }
    if yname == "is-pmtu-enabled" { return "IsPmtuEnabled" }
    if yname == "remote-tunnel-id" { return "RemoteTunnelId" }
    if yname == "local-tunnel-name" { return "LocalTunnelName" }
    if yname == "remote-tunnel-name" { return "RemoteTunnelName" }
    if yname == "class-name" { return "ClassName" }
    if yname == "active-sessions" { return "ActiveSessions" }
    if yname == "sequence-ns" { return "SequenceNs" }
    if yname == "sequence-nr" { return "SequenceNr" }
    if yname == "local-window-size" { return "LocalWindowSize" }
    if yname == "remote-window-size" { return "RemoteWindowSize" }
    if yname == "retransmission-time" { return "RetransmissionTime" }
    if yname == "maximum-retransmission-time" { return "MaximumRetransmissionTime" }
    if yname == "unsent-queue-size" { return "UnsentQueueSize" }
    if yname == "unsent-maximum-queue-size" { return "UnsentMaximumQueueSize" }
    if yname == "resend-queue-size" { return "ResendQueueSize" }
    if yname == "resend-maximum-queue-size" { return "ResendMaximumQueueSize" }
    if yname == "order-queue-size" { return "OrderQueueSize" }
    if yname == "packet-queue-check" { return "PacketQueueCheck" }
    if yname == "digest-secrets" { return "DigestSecrets" }
    if yname == "resends" { return "Resends" }
    if yname == "zero-length-body-acknowledgement-sent" { return "ZeroLengthBodyAcknowledgementSent" }
    if yname == "total-out-of-order-drop-packets" { return "TotalOutOfOrderDropPackets" }
    if yname == "total-out-of-order-reorder-packets" { return "TotalOutOfOrderReorderPackets" }
    if yname == "total-peer-authentication-failures" { return "TotalPeerAuthenticationFailures" }
    if yname == "is-tunnel-up" { return "IsTunnelUp" }
    if yname == "is-congestion-control-enabled" { return "IsCongestionControlEnabled" }
    if yname == "retransmit-time" { return "RetransmitTime" }
    return ""
}

func (tunnel *L2Tp_Tunnels_Tunnel) GetSegmentPath() string {
    return "tunnel" + "[local-tunnel-id='" + fmt.Sprintf("%v", tunnel.LocalTunnelId) + "']"
}

func (tunnel *L2Tp_Tunnels_Tunnel) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (tunnel *L2Tp_Tunnels_Tunnel) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (tunnel *L2Tp_Tunnels_Tunnel) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["local-tunnel-id"] = tunnel.LocalTunnelId
    leafs["local-address"] = tunnel.LocalAddress
    leafs["remote-address"] = tunnel.RemoteAddress
    leafs["local-port"] = tunnel.LocalPort
    leafs["remote-port"] = tunnel.RemotePort
    leafs["protocol"] = tunnel.Protocol
    leafs["is-pmtu-enabled"] = tunnel.IsPmtuEnabled
    leafs["remote-tunnel-id"] = tunnel.RemoteTunnelId
    leafs["local-tunnel-name"] = tunnel.LocalTunnelName
    leafs["remote-tunnel-name"] = tunnel.RemoteTunnelName
    leafs["class-name"] = tunnel.ClassName
    leafs["active-sessions"] = tunnel.ActiveSessions
    leafs["sequence-ns"] = tunnel.SequenceNs
    leafs["sequence-nr"] = tunnel.SequenceNr
    leafs["local-window-size"] = tunnel.LocalWindowSize
    leafs["remote-window-size"] = tunnel.RemoteWindowSize
    leafs["retransmission-time"] = tunnel.RetransmissionTime
    leafs["maximum-retransmission-time"] = tunnel.MaximumRetransmissionTime
    leafs["unsent-queue-size"] = tunnel.UnsentQueueSize
    leafs["unsent-maximum-queue-size"] = tunnel.UnsentMaximumQueueSize
    leafs["resend-queue-size"] = tunnel.ResendQueueSize
    leafs["resend-maximum-queue-size"] = tunnel.ResendMaximumQueueSize
    leafs["order-queue-size"] = tunnel.OrderQueueSize
    leafs["packet-queue-check"] = tunnel.PacketQueueCheck
    leafs["digest-secrets"] = tunnel.DigestSecrets
    leafs["resends"] = tunnel.Resends
    leafs["zero-length-body-acknowledgement-sent"] = tunnel.ZeroLengthBodyAcknowledgementSent
    leafs["total-out-of-order-drop-packets"] = tunnel.TotalOutOfOrderDropPackets
    leafs["total-out-of-order-reorder-packets"] = tunnel.TotalOutOfOrderReorderPackets
    leafs["total-peer-authentication-failures"] = tunnel.TotalPeerAuthenticationFailures
    leafs["is-tunnel-up"] = tunnel.IsTunnelUp
    leafs["is-congestion-control-enabled"] = tunnel.IsCongestionControlEnabled
    leafs["retransmit-time"] = tunnel.RetransmitTime
    return leafs
}

func (tunnel *L2Tp_Tunnels_Tunnel) GetBundleName() string { return "cisco_ios_xr" }

func (tunnel *L2Tp_Tunnels_Tunnel) GetYangName() string { return "tunnel" }

func (tunnel *L2Tp_Tunnels_Tunnel) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tunnel *L2Tp_Tunnels_Tunnel) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tunnel *L2Tp_Tunnels_Tunnel) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tunnel *L2Tp_Tunnels_Tunnel) SetParent(parent types.Entity) { tunnel.parent = parent }

func (tunnel *L2Tp_Tunnels_Tunnel) GetParent() types.Entity { return tunnel.parent }

func (tunnel *L2Tp_Tunnels_Tunnel) GetParentYangName() string { return "tunnels" }

// L2Tp_Sessions
// List of session IDs
type L2Tp_Sessions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // L2TP information for a particular session. The type is slice of
    // L2Tp_Sessions_Session.
    Session []L2Tp_Sessions_Session
}

func (sessions *L2Tp_Sessions) GetFilter() yfilter.YFilter { return sessions.YFilter }

func (sessions *L2Tp_Sessions) SetFilter(yf yfilter.YFilter) { sessions.YFilter = yf }

func (sessions *L2Tp_Sessions) GetGoName(yname string) string {
    if yname == "session" { return "Session" }
    return ""
}

func (sessions *L2Tp_Sessions) GetSegmentPath() string {
    return "sessions"
}

func (sessions *L2Tp_Sessions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "session" {
        for _, c := range sessions.Session {
            if sessions.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := L2Tp_Sessions_Session{}
        sessions.Session = append(sessions.Session, child)
        return &sessions.Session[len(sessions.Session)-1]
    }
    return nil
}

func (sessions *L2Tp_Sessions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sessions.Session {
        children[sessions.Session[i].GetSegmentPath()] = &sessions.Session[i]
    }
    return children
}

func (sessions *L2Tp_Sessions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sessions *L2Tp_Sessions) GetBundleName() string { return "cisco_ios_xr" }

func (sessions *L2Tp_Sessions) GetYangName() string { return "sessions" }

func (sessions *L2Tp_Sessions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sessions *L2Tp_Sessions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sessions *L2Tp_Sessions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sessions *L2Tp_Sessions) SetParent(parent types.Entity) { sessions.parent = parent }

func (sessions *L2Tp_Sessions) GetParent() types.Entity { return sessions.parent }

func (sessions *L2Tp_Sessions) GetParentYangName() string { return "l2tp" }

// L2Tp_Sessions_Session
// L2TP information for a particular session
type L2Tp_Sessions_Session struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Local tunnel ID. The type is interface{} with
    // range: -2147483648..2147483647.
    LocalTunnelId interface{}

    // This attribute is a key. Local session ID. The type is interface{} with
    // range: -2147483648..2147483647.
    LocalSessionId interface{}

    // Local session IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    LocalIpAddress interface{}

    // Remote session IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
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

func (session *L2Tp_Sessions_Session) GetFilter() yfilter.YFilter { return session.YFilter }

func (session *L2Tp_Sessions_Session) SetFilter(yf yfilter.YFilter) { session.YFilter = yf }

func (session *L2Tp_Sessions_Session) GetGoName(yname string) string {
    if yname == "local-tunnel-id" { return "LocalTunnelId" }
    if yname == "local-session-id" { return "LocalSessionId" }
    if yname == "local-ip-address" { return "LocalIpAddress" }
    if yname == "remote-ip-address" { return "RemoteIpAddress" }
    if yname == "l2tp-sh-sess-udp-lport" { return "L2TpShSessUdpLport" }
    if yname == "l2tp-sh-sess-udp-rport" { return "L2TpShSessUdpRport" }
    if yname == "protocol" { return "Protocol" }
    if yname == "remote-tunnel-id" { return "RemoteTunnelId" }
    if yname == "call-serial-number" { return "CallSerialNumber" }
    if yname == "local-tunnel-name" { return "LocalTunnelName" }
    if yname == "remote-tunnel-name" { return "RemoteTunnelName" }
    if yname == "remote-session-id" { return "RemoteSessionId" }
    if yname == "l2tp-sh-sess-tie-breaker-enabled" { return "L2TpShSessTieBreakerEnabled" }
    if yname == "l2tp-sh-sess-tie-breaker" { return "L2TpShSessTieBreaker" }
    if yname == "is-session-manual" { return "IsSessionManual" }
    if yname == "is-session-up" { return "IsSessionUp" }
    if yname == "is-udp-checksum-enabled" { return "IsUdpChecksumEnabled" }
    if yname == "is-sequencing-on" { return "IsSequencingOn" }
    if yname == "is-session-state-established" { return "IsSessionStateEstablished" }
    if yname == "is-session-locally-initiated" { return "IsSessionLocallyInitiated" }
    if yname == "is-conditional-debug-enabled" { return "IsConditionalDebugEnabled" }
    if yname == "unique-id" { return "UniqueId" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "session-application-data" { return "SessionApplicationData" }
    return ""
}

func (session *L2Tp_Sessions_Session) GetSegmentPath() string {
    return "session" + "[local-tunnel-id='" + fmt.Sprintf("%v", session.LocalTunnelId) + "']" + "[local-session-id='" + fmt.Sprintf("%v", session.LocalSessionId) + "']"
}

func (session *L2Tp_Sessions_Session) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "session-application-data" {
        return &session.SessionApplicationData
    }
    return nil
}

func (session *L2Tp_Sessions_Session) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["session-application-data"] = &session.SessionApplicationData
    return children
}

func (session *L2Tp_Sessions_Session) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["local-tunnel-id"] = session.LocalTunnelId
    leafs["local-session-id"] = session.LocalSessionId
    leafs["local-ip-address"] = session.LocalIpAddress
    leafs["remote-ip-address"] = session.RemoteIpAddress
    leafs["l2tp-sh-sess-udp-lport"] = session.L2TpShSessUdpLport
    leafs["l2tp-sh-sess-udp-rport"] = session.L2TpShSessUdpRport
    leafs["protocol"] = session.Protocol
    leafs["remote-tunnel-id"] = session.RemoteTunnelId
    leafs["call-serial-number"] = session.CallSerialNumber
    leafs["local-tunnel-name"] = session.LocalTunnelName
    leafs["remote-tunnel-name"] = session.RemoteTunnelName
    leafs["remote-session-id"] = session.RemoteSessionId
    leafs["l2tp-sh-sess-tie-breaker-enabled"] = session.L2TpShSessTieBreakerEnabled
    leafs["l2tp-sh-sess-tie-breaker"] = session.L2TpShSessTieBreaker
    leafs["is-session-manual"] = session.IsSessionManual
    leafs["is-session-up"] = session.IsSessionUp
    leafs["is-udp-checksum-enabled"] = session.IsUdpChecksumEnabled
    leafs["is-sequencing-on"] = session.IsSequencingOn
    leafs["is-session-state-established"] = session.IsSessionStateEstablished
    leafs["is-session-locally-initiated"] = session.IsSessionLocallyInitiated
    leafs["is-conditional-debug-enabled"] = session.IsConditionalDebugEnabled
    leafs["unique-id"] = session.UniqueId
    leafs["interface-name"] = session.InterfaceName
    return leafs
}

func (session *L2Tp_Sessions_Session) GetBundleName() string { return "cisco_ios_xr" }

func (session *L2Tp_Sessions_Session) GetYangName() string { return "session" }

func (session *L2Tp_Sessions_Session) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (session *L2Tp_Sessions_Session) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (session *L2Tp_Sessions_Session) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (session *L2Tp_Sessions_Session) SetParent(parent types.Entity) { session.parent = parent }

func (session *L2Tp_Sessions_Session) GetParent() types.Entity { return session.parent }

func (session *L2Tp_Sessions_Session) GetParentYangName() string { return "sessions" }

// L2Tp_Sessions_Session_SessionApplicationData
// Session application data
type L2Tp_Sessions_Session_SessionApplicationData struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // l2tp sh sess app type. The type is interface{} with range: 0..4294967295.
    L2TpShSessAppType interface{}

    // Xconnect data.
    Xconnect L2Tp_Sessions_Session_SessionApplicationData_Xconnect

    // VPDN data.
    Vpdn L2Tp_Sessions_Session_SessionApplicationData_Vpdn
}

func (sessionApplicationData *L2Tp_Sessions_Session_SessionApplicationData) GetFilter() yfilter.YFilter { return sessionApplicationData.YFilter }

func (sessionApplicationData *L2Tp_Sessions_Session_SessionApplicationData) SetFilter(yf yfilter.YFilter) { sessionApplicationData.YFilter = yf }

func (sessionApplicationData *L2Tp_Sessions_Session_SessionApplicationData) GetGoName(yname string) string {
    if yname == "l2tp-sh-sess-app-type" { return "L2TpShSessAppType" }
    if yname == "xconnect" { return "Xconnect" }
    if yname == "vpdn" { return "Vpdn" }
    return ""
}

func (sessionApplicationData *L2Tp_Sessions_Session_SessionApplicationData) GetSegmentPath() string {
    return "session-application-data"
}

func (sessionApplicationData *L2Tp_Sessions_Session_SessionApplicationData) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "xconnect" {
        return &sessionApplicationData.Xconnect
    }
    if childYangName == "vpdn" {
        return &sessionApplicationData.Vpdn
    }
    return nil
}

func (sessionApplicationData *L2Tp_Sessions_Session_SessionApplicationData) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["xconnect"] = &sessionApplicationData.Xconnect
    children["vpdn"] = &sessionApplicationData.Vpdn
    return children
}

func (sessionApplicationData *L2Tp_Sessions_Session_SessionApplicationData) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["l2tp-sh-sess-app-type"] = sessionApplicationData.L2TpShSessAppType
    return leafs
}

func (sessionApplicationData *L2Tp_Sessions_Session_SessionApplicationData) GetBundleName() string { return "cisco_ios_xr" }

func (sessionApplicationData *L2Tp_Sessions_Session_SessionApplicationData) GetYangName() string { return "session-application-data" }

func (sessionApplicationData *L2Tp_Sessions_Session_SessionApplicationData) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sessionApplicationData *L2Tp_Sessions_Session_SessionApplicationData) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sessionApplicationData *L2Tp_Sessions_Session_SessionApplicationData) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sessionApplicationData *L2Tp_Sessions_Session_SessionApplicationData) SetParent(parent types.Entity) { sessionApplicationData.parent = parent }

func (sessionApplicationData *L2Tp_Sessions_Session_SessionApplicationData) GetParent() types.Entity { return sessionApplicationData.parent }

func (sessionApplicationData *L2Tp_Sessions_Session_SessionApplicationData) GetParentYangName() string { return "session" }

// L2Tp_Sessions_Session_SessionApplicationData_Xconnect
// Xconnect data
type L2Tp_Sessions_Session_SessionApplicationData_Xconnect struct {
    parent types.Entity
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

func (xconnect *L2Tp_Sessions_Session_SessionApplicationData_Xconnect) GetFilter() yfilter.YFilter { return xconnect.YFilter }

func (xconnect *L2Tp_Sessions_Session_SessionApplicationData_Xconnect) SetFilter(yf yfilter.YFilter) { xconnect.YFilter = yf }

func (xconnect *L2Tp_Sessions_Session_SessionApplicationData_Xconnect) GetGoName(yname string) string {
    if yname == "circuit-name" { return "CircuitName" }
    if yname == "sessionvc-id" { return "SessionvcId" }
    if yname == "is-circuit-state-up" { return "IsCircuitStateUp" }
    if yname == "is-local-circuit-state-up" { return "IsLocalCircuitStateUp" }
    if yname == "is-remote-circuit-state-up" { return "IsRemoteCircuitStateUp" }
    if yname == "ipv6-protocol-tunneling" { return "Ipv6ProtocolTunneling" }
    return ""
}

func (xconnect *L2Tp_Sessions_Session_SessionApplicationData_Xconnect) GetSegmentPath() string {
    return "xconnect"
}

func (xconnect *L2Tp_Sessions_Session_SessionApplicationData_Xconnect) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (xconnect *L2Tp_Sessions_Session_SessionApplicationData_Xconnect) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (xconnect *L2Tp_Sessions_Session_SessionApplicationData_Xconnect) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["circuit-name"] = xconnect.CircuitName
    leafs["sessionvc-id"] = xconnect.SessionvcId
    leafs["is-circuit-state-up"] = xconnect.IsCircuitStateUp
    leafs["is-local-circuit-state-up"] = xconnect.IsLocalCircuitStateUp
    leafs["is-remote-circuit-state-up"] = xconnect.IsRemoteCircuitStateUp
    leafs["ipv6-protocol-tunneling"] = xconnect.Ipv6ProtocolTunneling
    return leafs
}

func (xconnect *L2Tp_Sessions_Session_SessionApplicationData_Xconnect) GetBundleName() string { return "cisco_ios_xr" }

func (xconnect *L2Tp_Sessions_Session_SessionApplicationData_Xconnect) GetYangName() string { return "xconnect" }

func (xconnect *L2Tp_Sessions_Session_SessionApplicationData_Xconnect) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (xconnect *L2Tp_Sessions_Session_SessionApplicationData_Xconnect) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (xconnect *L2Tp_Sessions_Session_SessionApplicationData_Xconnect) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (xconnect *L2Tp_Sessions_Session_SessionApplicationData_Xconnect) SetParent(parent types.Entity) { xconnect.parent = parent }

func (xconnect *L2Tp_Sessions_Session_SessionApplicationData_Xconnect) GetParent() types.Entity { return xconnect.parent }

func (xconnect *L2Tp_Sessions_Session_SessionApplicationData_Xconnect) GetParentYangName() string { return "session-application-data" }

// L2Tp_Sessions_Session_SessionApplicationData_Vpdn
// VPDN data
type L2Tp_Sessions_Session_SessionApplicationData_Vpdn struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Session username. The type is string.
    Username interface{}

    // Interface name. The type is string with pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}
}

func (vpdn *L2Tp_Sessions_Session_SessionApplicationData_Vpdn) GetFilter() yfilter.YFilter { return vpdn.YFilter }

func (vpdn *L2Tp_Sessions_Session_SessionApplicationData_Vpdn) SetFilter(yf yfilter.YFilter) { vpdn.YFilter = yf }

func (vpdn *L2Tp_Sessions_Session_SessionApplicationData_Vpdn) GetGoName(yname string) string {
    if yname == "username" { return "Username" }
    if yname == "interface-name" { return "InterfaceName" }
    return ""
}

func (vpdn *L2Tp_Sessions_Session_SessionApplicationData_Vpdn) GetSegmentPath() string {
    return "vpdn"
}

func (vpdn *L2Tp_Sessions_Session_SessionApplicationData_Vpdn) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vpdn *L2Tp_Sessions_Session_SessionApplicationData_Vpdn) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vpdn *L2Tp_Sessions_Session_SessionApplicationData_Vpdn) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["username"] = vpdn.Username
    leafs["interface-name"] = vpdn.InterfaceName
    return leafs
}

func (vpdn *L2Tp_Sessions_Session_SessionApplicationData_Vpdn) GetBundleName() string { return "cisco_ios_xr" }

func (vpdn *L2Tp_Sessions_Session_SessionApplicationData_Vpdn) GetYangName() string { return "vpdn" }

func (vpdn *L2Tp_Sessions_Session_SessionApplicationData_Vpdn) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vpdn *L2Tp_Sessions_Session_SessionApplicationData_Vpdn) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vpdn *L2Tp_Sessions_Session_SessionApplicationData_Vpdn) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vpdn *L2Tp_Sessions_Session_SessionApplicationData_Vpdn) SetParent(parent types.Entity) { vpdn.parent = parent }

func (vpdn *L2Tp_Sessions_Session_SessionApplicationData_Vpdn) GetParent() types.Entity { return vpdn.parent }

func (vpdn *L2Tp_Sessions_Session_SessionApplicationData_Vpdn) GetParentYangName() string { return "session-application-data" }

// L2Tp_Session
// L2TP control messages counters
type L2Tp_Session struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // L2TP session unavailable  information.
    Unavailable L2Tp_Session_Unavailable
}

func (session *L2Tp_Session) GetFilter() yfilter.YFilter { return session.YFilter }

func (session *L2Tp_Session) SetFilter(yf yfilter.YFilter) { session.YFilter = yf }

func (session *L2Tp_Session) GetGoName(yname string) string {
    if yname == "unavailable" { return "Unavailable" }
    return ""
}

func (session *L2Tp_Session) GetSegmentPath() string {
    return "session"
}

func (session *L2Tp_Session) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "unavailable" {
        return &session.Unavailable
    }
    return nil
}

func (session *L2Tp_Session) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["unavailable"] = &session.Unavailable
    return children
}

func (session *L2Tp_Session) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (session *L2Tp_Session) GetBundleName() string { return "cisco_ios_xr" }

func (session *L2Tp_Session) GetYangName() string { return "session" }

func (session *L2Tp_Session) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (session *L2Tp_Session) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (session *L2Tp_Session) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (session *L2Tp_Session) SetParent(parent types.Entity) { session.parent = parent }

func (session *L2Tp_Session) GetParent() types.Entity { return session.parent }

func (session *L2Tp_Session) GetParentYangName() string { return "l2tp" }

// L2Tp_Session_Unavailable
// L2TP session unavailable  information
type L2Tp_Session_Unavailable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of session ID in hold database. The type is interface{} with range:
    // 0..4294967295.
    SessionsOnHold interface{}
}

func (unavailable *L2Tp_Session_Unavailable) GetFilter() yfilter.YFilter { return unavailable.YFilter }

func (unavailable *L2Tp_Session_Unavailable) SetFilter(yf yfilter.YFilter) { unavailable.YFilter = yf }

func (unavailable *L2Tp_Session_Unavailable) GetGoName(yname string) string {
    if yname == "sessions-on-hold" { return "SessionsOnHold" }
    return ""
}

func (unavailable *L2Tp_Session_Unavailable) GetSegmentPath() string {
    return "unavailable"
}

func (unavailable *L2Tp_Session_Unavailable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (unavailable *L2Tp_Session_Unavailable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (unavailable *L2Tp_Session_Unavailable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sessions-on-hold"] = unavailable.SessionsOnHold
    return leafs
}

func (unavailable *L2Tp_Session_Unavailable) GetBundleName() string { return "cisco_ios_xr" }

func (unavailable *L2Tp_Session_Unavailable) GetYangName() string { return "unavailable" }

func (unavailable *L2Tp_Session_Unavailable) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (unavailable *L2Tp_Session_Unavailable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (unavailable *L2Tp_Session_Unavailable) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (unavailable *L2Tp_Session_Unavailable) SetParent(parent types.Entity) { unavailable.parent = parent }

func (unavailable *L2Tp_Session_Unavailable) GetParent() types.Entity { return unavailable.parent }

func (unavailable *L2Tp_Session_Unavailable) GetParentYangName() string { return "session" }

// L2Tpv2
// l2tpv2
type L2Tpv2 struct {
    parent types.Entity
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

func (l2Tpv2 *L2Tpv2) GetFilter() yfilter.YFilter { return l2Tpv2.YFilter }

func (l2Tpv2 *L2Tpv2) SetFilter(yf yfilter.YFilter) { l2Tpv2.YFilter = yf }

func (l2Tpv2 *L2Tpv2) GetGoName(yname string) string {
    if yname == "counters" { return "Counters" }
    if yname == "statistics" { return "Statistics" }
    if yname == "tunnel" { return "Tunnel" }
    if yname == "tunnel-configurations" { return "TunnelConfigurations" }
    if yname == "counter-hist-fail" { return "CounterHistFail" }
    if yname == "classes" { return "Classes" }
    if yname == "tunnels" { return "Tunnels" }
    if yname == "sessions" { return "Sessions" }
    if yname == "session" { return "Session" }
    return ""
}

func (l2Tpv2 *L2Tpv2) GetSegmentPath() string {
    return "Cisco-IOS-XR-tunnel-l2tun-oper:l2tpv2"
}

func (l2Tpv2 *L2Tpv2) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "counters" {
        return &l2Tpv2.Counters
    }
    if childYangName == "statistics" {
        return &l2Tpv2.Statistics
    }
    if childYangName == "tunnel" {
        return &l2Tpv2.Tunnel
    }
    if childYangName == "tunnel-configurations" {
        return &l2Tpv2.TunnelConfigurations
    }
    if childYangName == "counter-hist-fail" {
        return &l2Tpv2.CounterHistFail
    }
    if childYangName == "classes" {
        return &l2Tpv2.Classes
    }
    if childYangName == "tunnels" {
        return &l2Tpv2.Tunnels
    }
    if childYangName == "sessions" {
        return &l2Tpv2.Sessions
    }
    if childYangName == "session" {
        return &l2Tpv2.Session
    }
    return nil
}

func (l2Tpv2 *L2Tpv2) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["counters"] = &l2Tpv2.Counters
    children["statistics"] = &l2Tpv2.Statistics
    children["tunnel"] = &l2Tpv2.Tunnel
    children["tunnel-configurations"] = &l2Tpv2.TunnelConfigurations
    children["counter-hist-fail"] = &l2Tpv2.CounterHistFail
    children["classes"] = &l2Tpv2.Classes
    children["tunnels"] = &l2Tpv2.Tunnels
    children["sessions"] = &l2Tpv2.Sessions
    children["session"] = &l2Tpv2.Session
    return children
}

func (l2Tpv2 *L2Tpv2) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (l2Tpv2 *L2Tpv2) GetBundleName() string { return "cisco_ios_xr" }

func (l2Tpv2 *L2Tpv2) GetYangName() string { return "l2tpv2" }

func (l2Tpv2 *L2Tpv2) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (l2Tpv2 *L2Tpv2) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (l2Tpv2 *L2Tpv2) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (l2Tpv2 *L2Tpv2) SetParent(parent types.Entity) { l2Tpv2.parent = parent }

func (l2Tpv2 *L2Tpv2) GetParent() types.Entity { return l2Tpv2.parent }

func (l2Tpv2 *L2Tpv2) GetParentYangName() string { return "Cisco-IOS-XR-tunnel-l2tun-oper" }

// L2Tpv2_Counters
// L2TP control messages counters
type L2Tpv2_Counters struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // L2TP forwarding messages counters.
    Forwarding L2Tpv2_Counters_Forwarding

    // L2TP control messages counters.
    Control L2Tpv2_Counters_Control
}

func (counters *L2Tpv2_Counters) GetFilter() yfilter.YFilter { return counters.YFilter }

func (counters *L2Tpv2_Counters) SetFilter(yf yfilter.YFilter) { counters.YFilter = yf }

func (counters *L2Tpv2_Counters) GetGoName(yname string) string {
    if yname == "forwarding" { return "Forwarding" }
    if yname == "control" { return "Control" }
    return ""
}

func (counters *L2Tpv2_Counters) GetSegmentPath() string {
    return "counters"
}

func (counters *L2Tpv2_Counters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "forwarding" {
        return &counters.Forwarding
    }
    if childYangName == "control" {
        return &counters.Control
    }
    return nil
}

func (counters *L2Tpv2_Counters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["forwarding"] = &counters.Forwarding
    children["control"] = &counters.Control
    return children
}

func (counters *L2Tpv2_Counters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (counters *L2Tpv2_Counters) GetBundleName() string { return "cisco_ios_xr" }

func (counters *L2Tpv2_Counters) GetYangName() string { return "counters" }

func (counters *L2Tpv2_Counters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (counters *L2Tpv2_Counters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (counters *L2Tpv2_Counters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (counters *L2Tpv2_Counters) SetParent(parent types.Entity) { counters.parent = parent }

func (counters *L2Tpv2_Counters) GetParent() types.Entity { return counters.parent }

func (counters *L2Tpv2_Counters) GetParentYangName() string { return "l2tpv2" }

// L2Tpv2_Counters_Forwarding
// L2TP forwarding messages counters
type L2Tpv2_Counters_Forwarding struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of class and session IDs.
    Sessions L2Tpv2_Counters_Forwarding_Sessions
}

func (forwarding *L2Tpv2_Counters_Forwarding) GetFilter() yfilter.YFilter { return forwarding.YFilter }

func (forwarding *L2Tpv2_Counters_Forwarding) SetFilter(yf yfilter.YFilter) { forwarding.YFilter = yf }

func (forwarding *L2Tpv2_Counters_Forwarding) GetGoName(yname string) string {
    if yname == "sessions" { return "Sessions" }
    return ""
}

func (forwarding *L2Tpv2_Counters_Forwarding) GetSegmentPath() string {
    return "forwarding"
}

func (forwarding *L2Tpv2_Counters_Forwarding) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sessions" {
        return &forwarding.Sessions
    }
    return nil
}

func (forwarding *L2Tpv2_Counters_Forwarding) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sessions"] = &forwarding.Sessions
    return children
}

func (forwarding *L2Tpv2_Counters_Forwarding) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (forwarding *L2Tpv2_Counters_Forwarding) GetBundleName() string { return "cisco_ios_xr" }

func (forwarding *L2Tpv2_Counters_Forwarding) GetYangName() string { return "forwarding" }

func (forwarding *L2Tpv2_Counters_Forwarding) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (forwarding *L2Tpv2_Counters_Forwarding) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (forwarding *L2Tpv2_Counters_Forwarding) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (forwarding *L2Tpv2_Counters_Forwarding) SetParent(parent types.Entity) { forwarding.parent = parent }

func (forwarding *L2Tpv2_Counters_Forwarding) GetParent() types.Entity { return forwarding.parent }

func (forwarding *L2Tpv2_Counters_Forwarding) GetParentYangName() string { return "counters" }

// L2Tpv2_Counters_Forwarding_Sessions
// List of class and session IDs
type L2Tpv2_Counters_Forwarding_Sessions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // L2TP information for a particular session. The type is slice of
    // L2Tpv2_Counters_Forwarding_Sessions_Session.
    Session []L2Tpv2_Counters_Forwarding_Sessions_Session
}

func (sessions *L2Tpv2_Counters_Forwarding_Sessions) GetFilter() yfilter.YFilter { return sessions.YFilter }

func (sessions *L2Tpv2_Counters_Forwarding_Sessions) SetFilter(yf yfilter.YFilter) { sessions.YFilter = yf }

func (sessions *L2Tpv2_Counters_Forwarding_Sessions) GetGoName(yname string) string {
    if yname == "session" { return "Session" }
    return ""
}

func (sessions *L2Tpv2_Counters_Forwarding_Sessions) GetSegmentPath() string {
    return "sessions"
}

func (sessions *L2Tpv2_Counters_Forwarding_Sessions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "session" {
        for _, c := range sessions.Session {
            if sessions.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := L2Tpv2_Counters_Forwarding_Sessions_Session{}
        sessions.Session = append(sessions.Session, child)
        return &sessions.Session[len(sessions.Session)-1]
    }
    return nil
}

func (sessions *L2Tpv2_Counters_Forwarding_Sessions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sessions.Session {
        children[sessions.Session[i].GetSegmentPath()] = &sessions.Session[i]
    }
    return children
}

func (sessions *L2Tpv2_Counters_Forwarding_Sessions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sessions *L2Tpv2_Counters_Forwarding_Sessions) GetBundleName() string { return "cisco_ios_xr" }

func (sessions *L2Tpv2_Counters_Forwarding_Sessions) GetYangName() string { return "sessions" }

func (sessions *L2Tpv2_Counters_Forwarding_Sessions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sessions *L2Tpv2_Counters_Forwarding_Sessions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sessions *L2Tpv2_Counters_Forwarding_Sessions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sessions *L2Tpv2_Counters_Forwarding_Sessions) SetParent(parent types.Entity) { sessions.parent = parent }

func (sessions *L2Tpv2_Counters_Forwarding_Sessions) GetParent() types.Entity { return sessions.parent }

func (sessions *L2Tpv2_Counters_Forwarding_Sessions) GetParentYangName() string { return "forwarding" }

// L2Tpv2_Counters_Forwarding_Sessions_Session
// L2TP information for a particular session
type L2Tpv2_Counters_Forwarding_Sessions_Session struct {
    parent types.Entity
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

func (session *L2Tpv2_Counters_Forwarding_Sessions_Session) GetFilter() yfilter.YFilter { return session.YFilter }

func (session *L2Tpv2_Counters_Forwarding_Sessions_Session) SetFilter(yf yfilter.YFilter) { session.YFilter = yf }

func (session *L2Tpv2_Counters_Forwarding_Sessions_Session) GetGoName(yname string) string {
    if yname == "tunnel-id" { return "TunnelId" }
    if yname == "session-id" { return "SessionId" }
    if yname == "remote-session-id" { return "RemoteSessionId" }
    if yname == "in-packets" { return "InPackets" }
    if yname == "out-packets" { return "OutPackets" }
    if yname == "in-bytes" { return "InBytes" }
    if yname == "out-bytes" { return "OutBytes" }
    return ""
}

func (session *L2Tpv2_Counters_Forwarding_Sessions_Session) GetSegmentPath() string {
    return "session" + "[tunnel-id='" + fmt.Sprintf("%v", session.TunnelId) + "']" + "[session-id='" + fmt.Sprintf("%v", session.SessionId) + "']"
}

func (session *L2Tpv2_Counters_Forwarding_Sessions_Session) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (session *L2Tpv2_Counters_Forwarding_Sessions_Session) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (session *L2Tpv2_Counters_Forwarding_Sessions_Session) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tunnel-id"] = session.TunnelId
    leafs["session-id"] = session.SessionId
    leafs["remote-session-id"] = session.RemoteSessionId
    leafs["in-packets"] = session.InPackets
    leafs["out-packets"] = session.OutPackets
    leafs["in-bytes"] = session.InBytes
    leafs["out-bytes"] = session.OutBytes
    return leafs
}

func (session *L2Tpv2_Counters_Forwarding_Sessions_Session) GetBundleName() string { return "cisco_ios_xr" }

func (session *L2Tpv2_Counters_Forwarding_Sessions_Session) GetYangName() string { return "session" }

func (session *L2Tpv2_Counters_Forwarding_Sessions_Session) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (session *L2Tpv2_Counters_Forwarding_Sessions_Session) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (session *L2Tpv2_Counters_Forwarding_Sessions_Session) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (session *L2Tpv2_Counters_Forwarding_Sessions_Session) SetParent(parent types.Entity) { session.parent = parent }

func (session *L2Tpv2_Counters_Forwarding_Sessions_Session) GetParent() types.Entity { return session.parent }

func (session *L2Tpv2_Counters_Forwarding_Sessions_Session) GetParentYangName() string { return "sessions" }

// L2Tpv2_Counters_Control
// L2TP control messages counters
type L2Tpv2_Counters_Control struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // L2TP control tunnel messages counters.
    TunnelXr L2Tpv2_Counters_Control_TunnelXr

    // Table of tunnel IDs of control message counters.
    Tunnels L2Tpv2_Counters_Control_Tunnels
}

func (control *L2Tpv2_Counters_Control) GetFilter() yfilter.YFilter { return control.YFilter }

func (control *L2Tpv2_Counters_Control) SetFilter(yf yfilter.YFilter) { control.YFilter = yf }

func (control *L2Tpv2_Counters_Control) GetGoName(yname string) string {
    if yname == "tunnel-xr" { return "TunnelXr" }
    if yname == "tunnels" { return "Tunnels" }
    return ""
}

func (control *L2Tpv2_Counters_Control) GetSegmentPath() string {
    return "control"
}

func (control *L2Tpv2_Counters_Control) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tunnel-xr" {
        return &control.TunnelXr
    }
    if childYangName == "tunnels" {
        return &control.Tunnels
    }
    return nil
}

func (control *L2Tpv2_Counters_Control) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["tunnel-xr"] = &control.TunnelXr
    children["tunnels"] = &control.Tunnels
    return children
}

func (control *L2Tpv2_Counters_Control) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (control *L2Tpv2_Counters_Control) GetBundleName() string { return "cisco_ios_xr" }

func (control *L2Tpv2_Counters_Control) GetYangName() string { return "control" }

func (control *L2Tpv2_Counters_Control) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (control *L2Tpv2_Counters_Control) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (control *L2Tpv2_Counters_Control) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (control *L2Tpv2_Counters_Control) SetParent(parent types.Entity) { control.parent = parent }

func (control *L2Tpv2_Counters_Control) GetParent() types.Entity { return control.parent }

func (control *L2Tpv2_Counters_Control) GetParentYangName() string { return "counters" }

// L2Tpv2_Counters_Control_TunnelXr
// L2TP control tunnel messages counters
type L2Tpv2_Counters_Control_TunnelXr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Tunnel authentication counters.
    Authentication L2Tpv2_Counters_Control_TunnelXr_Authentication

    // Tunnel counters.
    Global L2Tpv2_Counters_Control_TunnelXr_Global
}

func (tunnelXr *L2Tpv2_Counters_Control_TunnelXr) GetFilter() yfilter.YFilter { return tunnelXr.YFilter }

func (tunnelXr *L2Tpv2_Counters_Control_TunnelXr) SetFilter(yf yfilter.YFilter) { tunnelXr.YFilter = yf }

func (tunnelXr *L2Tpv2_Counters_Control_TunnelXr) GetGoName(yname string) string {
    if yname == "authentication" { return "Authentication" }
    if yname == "global" { return "Global" }
    return ""
}

func (tunnelXr *L2Tpv2_Counters_Control_TunnelXr) GetSegmentPath() string {
    return "tunnel-xr"
}

func (tunnelXr *L2Tpv2_Counters_Control_TunnelXr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "authentication" {
        return &tunnelXr.Authentication
    }
    if childYangName == "global" {
        return &tunnelXr.Global
    }
    return nil
}

func (tunnelXr *L2Tpv2_Counters_Control_TunnelXr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["authentication"] = &tunnelXr.Authentication
    children["global"] = &tunnelXr.Global
    return children
}

func (tunnelXr *L2Tpv2_Counters_Control_TunnelXr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (tunnelXr *L2Tpv2_Counters_Control_TunnelXr) GetBundleName() string { return "cisco_ios_xr" }

func (tunnelXr *L2Tpv2_Counters_Control_TunnelXr) GetYangName() string { return "tunnel-xr" }

func (tunnelXr *L2Tpv2_Counters_Control_TunnelXr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tunnelXr *L2Tpv2_Counters_Control_TunnelXr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tunnelXr *L2Tpv2_Counters_Control_TunnelXr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tunnelXr *L2Tpv2_Counters_Control_TunnelXr) SetParent(parent types.Entity) { tunnelXr.parent = parent }

func (tunnelXr *L2Tpv2_Counters_Control_TunnelXr) GetParent() types.Entity { return tunnelXr.parent }

func (tunnelXr *L2Tpv2_Counters_Control_TunnelXr) GetParentYangName() string { return "control" }

// L2Tpv2_Counters_Control_TunnelXr_Authentication
// Tunnel authentication counters
type L2Tpv2_Counters_Control_TunnelXr_Authentication struct {
    parent types.Entity
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

func (authentication *L2Tpv2_Counters_Control_TunnelXr_Authentication) GetFilter() yfilter.YFilter { return authentication.YFilter }

func (authentication *L2Tpv2_Counters_Control_TunnelXr_Authentication) SetFilter(yf yfilter.YFilter) { authentication.YFilter = yf }

func (authentication *L2Tpv2_Counters_Control_TunnelXr_Authentication) GetGoName(yname string) string {
    if yname == "nonce-avp" { return "NonceAvp" }
    if yname == "common-digest" { return "CommonDigest" }
    if yname == "primary-digest" { return "PrimaryDigest" }
    if yname == "secondary-digest" { return "SecondaryDigest" }
    if yname == "integrity-check" { return "IntegrityCheck" }
    if yname == "local-secret" { return "LocalSecret" }
    if yname == "challenge-avp" { return "ChallengeAvp" }
    if yname == "challenge-reponse" { return "ChallengeReponse" }
    if yname == "overall-statistics" { return "OverallStatistics" }
    return ""
}

func (authentication *L2Tpv2_Counters_Control_TunnelXr_Authentication) GetSegmentPath() string {
    return "authentication"
}

func (authentication *L2Tpv2_Counters_Control_TunnelXr_Authentication) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nonce-avp" {
        return &authentication.NonceAvp
    }
    if childYangName == "common-digest" {
        return &authentication.CommonDigest
    }
    if childYangName == "primary-digest" {
        return &authentication.PrimaryDigest
    }
    if childYangName == "secondary-digest" {
        return &authentication.SecondaryDigest
    }
    if childYangName == "integrity-check" {
        return &authentication.IntegrityCheck
    }
    if childYangName == "local-secret" {
        return &authentication.LocalSecret
    }
    if childYangName == "challenge-avp" {
        return &authentication.ChallengeAvp
    }
    if childYangName == "challenge-reponse" {
        return &authentication.ChallengeReponse
    }
    if childYangName == "overall-statistics" {
        return &authentication.OverallStatistics
    }
    return nil
}

func (authentication *L2Tpv2_Counters_Control_TunnelXr_Authentication) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nonce-avp"] = &authentication.NonceAvp
    children["common-digest"] = &authentication.CommonDigest
    children["primary-digest"] = &authentication.PrimaryDigest
    children["secondary-digest"] = &authentication.SecondaryDigest
    children["integrity-check"] = &authentication.IntegrityCheck
    children["local-secret"] = &authentication.LocalSecret
    children["challenge-avp"] = &authentication.ChallengeAvp
    children["challenge-reponse"] = &authentication.ChallengeReponse
    children["overall-statistics"] = &authentication.OverallStatistics
    return children
}

func (authentication *L2Tpv2_Counters_Control_TunnelXr_Authentication) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (authentication *L2Tpv2_Counters_Control_TunnelXr_Authentication) GetBundleName() string { return "cisco_ios_xr" }

func (authentication *L2Tpv2_Counters_Control_TunnelXr_Authentication) GetYangName() string { return "authentication" }

func (authentication *L2Tpv2_Counters_Control_TunnelXr_Authentication) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (authentication *L2Tpv2_Counters_Control_TunnelXr_Authentication) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (authentication *L2Tpv2_Counters_Control_TunnelXr_Authentication) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (authentication *L2Tpv2_Counters_Control_TunnelXr_Authentication) SetParent(parent types.Entity) { authentication.parent = parent }

func (authentication *L2Tpv2_Counters_Control_TunnelXr_Authentication) GetParent() types.Entity { return authentication.parent }

func (authentication *L2Tpv2_Counters_Control_TunnelXr_Authentication) GetParentYangName() string { return "tunnel-xr" }

// L2Tpv2_Counters_Control_TunnelXr_Authentication_NonceAvp
// Nonce AVP statistics
type L2Tpv2_Counters_Control_TunnelXr_Authentication_NonceAvp struct {
    parent types.Entity
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

func (nonceAvp *L2Tpv2_Counters_Control_TunnelXr_Authentication_NonceAvp) GetFilter() yfilter.YFilter { return nonceAvp.YFilter }

func (nonceAvp *L2Tpv2_Counters_Control_TunnelXr_Authentication_NonceAvp) SetFilter(yf yfilter.YFilter) { nonceAvp.YFilter = yf }

func (nonceAvp *L2Tpv2_Counters_Control_TunnelXr_Authentication_NonceAvp) GetGoName(yname string) string {
    if yname == "validate" { return "Validate" }
    if yname == "bad-hash" { return "BadHash" }
    if yname == "bad-length" { return "BadLength" }
    if yname == "ignored" { return "Ignored" }
    if yname == "missing" { return "Missing" }
    if yname == "passed" { return "Passed" }
    if yname == "failed" { return "Failed" }
    if yname == "skipped" { return "Skipped" }
    if yname == "generate-response-failures" { return "GenerateResponseFailures" }
    if yname == "unexpected" { return "Unexpected" }
    if yname == "unexpected-zlb" { return "UnexpectedZlb" }
    return ""
}

func (nonceAvp *L2Tpv2_Counters_Control_TunnelXr_Authentication_NonceAvp) GetSegmentPath() string {
    return "nonce-avp"
}

func (nonceAvp *L2Tpv2_Counters_Control_TunnelXr_Authentication_NonceAvp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nonceAvp *L2Tpv2_Counters_Control_TunnelXr_Authentication_NonceAvp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nonceAvp *L2Tpv2_Counters_Control_TunnelXr_Authentication_NonceAvp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["validate"] = nonceAvp.Validate
    leafs["bad-hash"] = nonceAvp.BadHash
    leafs["bad-length"] = nonceAvp.BadLength
    leafs["ignored"] = nonceAvp.Ignored
    leafs["missing"] = nonceAvp.Missing
    leafs["passed"] = nonceAvp.Passed
    leafs["failed"] = nonceAvp.Failed
    leafs["skipped"] = nonceAvp.Skipped
    leafs["generate-response-failures"] = nonceAvp.GenerateResponseFailures
    leafs["unexpected"] = nonceAvp.Unexpected
    leafs["unexpected-zlb"] = nonceAvp.UnexpectedZlb
    return leafs
}

func (nonceAvp *L2Tpv2_Counters_Control_TunnelXr_Authentication_NonceAvp) GetBundleName() string { return "cisco_ios_xr" }

func (nonceAvp *L2Tpv2_Counters_Control_TunnelXr_Authentication_NonceAvp) GetYangName() string { return "nonce-avp" }

func (nonceAvp *L2Tpv2_Counters_Control_TunnelXr_Authentication_NonceAvp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nonceAvp *L2Tpv2_Counters_Control_TunnelXr_Authentication_NonceAvp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nonceAvp *L2Tpv2_Counters_Control_TunnelXr_Authentication_NonceAvp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nonceAvp *L2Tpv2_Counters_Control_TunnelXr_Authentication_NonceAvp) SetParent(parent types.Entity) { nonceAvp.parent = parent }

func (nonceAvp *L2Tpv2_Counters_Control_TunnelXr_Authentication_NonceAvp) GetParent() types.Entity { return nonceAvp.parent }

func (nonceAvp *L2Tpv2_Counters_Control_TunnelXr_Authentication_NonceAvp) GetParentYangName() string { return "authentication" }

// L2Tpv2_Counters_Control_TunnelXr_Authentication_CommonDigest
// Common digest statistics
type L2Tpv2_Counters_Control_TunnelXr_Authentication_CommonDigest struct {
    parent types.Entity
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

func (commonDigest *L2Tpv2_Counters_Control_TunnelXr_Authentication_CommonDigest) GetFilter() yfilter.YFilter { return commonDigest.YFilter }

func (commonDigest *L2Tpv2_Counters_Control_TunnelXr_Authentication_CommonDigest) SetFilter(yf yfilter.YFilter) { commonDigest.YFilter = yf }

func (commonDigest *L2Tpv2_Counters_Control_TunnelXr_Authentication_CommonDigest) GetGoName(yname string) string {
    if yname == "validate" { return "Validate" }
    if yname == "bad-hash" { return "BadHash" }
    if yname == "bad-length" { return "BadLength" }
    if yname == "ignored" { return "Ignored" }
    if yname == "missing" { return "Missing" }
    if yname == "passed" { return "Passed" }
    if yname == "failed" { return "Failed" }
    if yname == "skipped" { return "Skipped" }
    if yname == "generate-response-failures" { return "GenerateResponseFailures" }
    if yname == "unexpected" { return "Unexpected" }
    if yname == "unexpected-zlb" { return "UnexpectedZlb" }
    return ""
}

func (commonDigest *L2Tpv2_Counters_Control_TunnelXr_Authentication_CommonDigest) GetSegmentPath() string {
    return "common-digest"
}

func (commonDigest *L2Tpv2_Counters_Control_TunnelXr_Authentication_CommonDigest) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (commonDigest *L2Tpv2_Counters_Control_TunnelXr_Authentication_CommonDigest) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (commonDigest *L2Tpv2_Counters_Control_TunnelXr_Authentication_CommonDigest) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["validate"] = commonDigest.Validate
    leafs["bad-hash"] = commonDigest.BadHash
    leafs["bad-length"] = commonDigest.BadLength
    leafs["ignored"] = commonDigest.Ignored
    leafs["missing"] = commonDigest.Missing
    leafs["passed"] = commonDigest.Passed
    leafs["failed"] = commonDigest.Failed
    leafs["skipped"] = commonDigest.Skipped
    leafs["generate-response-failures"] = commonDigest.GenerateResponseFailures
    leafs["unexpected"] = commonDigest.Unexpected
    leafs["unexpected-zlb"] = commonDigest.UnexpectedZlb
    return leafs
}

func (commonDigest *L2Tpv2_Counters_Control_TunnelXr_Authentication_CommonDigest) GetBundleName() string { return "cisco_ios_xr" }

func (commonDigest *L2Tpv2_Counters_Control_TunnelXr_Authentication_CommonDigest) GetYangName() string { return "common-digest" }

func (commonDigest *L2Tpv2_Counters_Control_TunnelXr_Authentication_CommonDigest) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (commonDigest *L2Tpv2_Counters_Control_TunnelXr_Authentication_CommonDigest) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (commonDigest *L2Tpv2_Counters_Control_TunnelXr_Authentication_CommonDigest) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (commonDigest *L2Tpv2_Counters_Control_TunnelXr_Authentication_CommonDigest) SetParent(parent types.Entity) { commonDigest.parent = parent }

func (commonDigest *L2Tpv2_Counters_Control_TunnelXr_Authentication_CommonDigest) GetParent() types.Entity { return commonDigest.parent }

func (commonDigest *L2Tpv2_Counters_Control_TunnelXr_Authentication_CommonDigest) GetParentYangName() string { return "authentication" }

// L2Tpv2_Counters_Control_TunnelXr_Authentication_PrimaryDigest
// Primary digest statistics
type L2Tpv2_Counters_Control_TunnelXr_Authentication_PrimaryDigest struct {
    parent types.Entity
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

func (primaryDigest *L2Tpv2_Counters_Control_TunnelXr_Authentication_PrimaryDigest) GetFilter() yfilter.YFilter { return primaryDigest.YFilter }

func (primaryDigest *L2Tpv2_Counters_Control_TunnelXr_Authentication_PrimaryDigest) SetFilter(yf yfilter.YFilter) { primaryDigest.YFilter = yf }

func (primaryDigest *L2Tpv2_Counters_Control_TunnelXr_Authentication_PrimaryDigest) GetGoName(yname string) string {
    if yname == "validate" { return "Validate" }
    if yname == "bad-hash" { return "BadHash" }
    if yname == "bad-length" { return "BadLength" }
    if yname == "ignored" { return "Ignored" }
    if yname == "missing" { return "Missing" }
    if yname == "passed" { return "Passed" }
    if yname == "failed" { return "Failed" }
    if yname == "skipped" { return "Skipped" }
    if yname == "generate-response-failures" { return "GenerateResponseFailures" }
    if yname == "unexpected" { return "Unexpected" }
    if yname == "unexpected-zlb" { return "UnexpectedZlb" }
    return ""
}

func (primaryDigest *L2Tpv2_Counters_Control_TunnelXr_Authentication_PrimaryDigest) GetSegmentPath() string {
    return "primary-digest"
}

func (primaryDigest *L2Tpv2_Counters_Control_TunnelXr_Authentication_PrimaryDigest) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (primaryDigest *L2Tpv2_Counters_Control_TunnelXr_Authentication_PrimaryDigest) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (primaryDigest *L2Tpv2_Counters_Control_TunnelXr_Authentication_PrimaryDigest) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["validate"] = primaryDigest.Validate
    leafs["bad-hash"] = primaryDigest.BadHash
    leafs["bad-length"] = primaryDigest.BadLength
    leafs["ignored"] = primaryDigest.Ignored
    leafs["missing"] = primaryDigest.Missing
    leafs["passed"] = primaryDigest.Passed
    leafs["failed"] = primaryDigest.Failed
    leafs["skipped"] = primaryDigest.Skipped
    leafs["generate-response-failures"] = primaryDigest.GenerateResponseFailures
    leafs["unexpected"] = primaryDigest.Unexpected
    leafs["unexpected-zlb"] = primaryDigest.UnexpectedZlb
    return leafs
}

func (primaryDigest *L2Tpv2_Counters_Control_TunnelXr_Authentication_PrimaryDigest) GetBundleName() string { return "cisco_ios_xr" }

func (primaryDigest *L2Tpv2_Counters_Control_TunnelXr_Authentication_PrimaryDigest) GetYangName() string { return "primary-digest" }

func (primaryDigest *L2Tpv2_Counters_Control_TunnelXr_Authentication_PrimaryDigest) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (primaryDigest *L2Tpv2_Counters_Control_TunnelXr_Authentication_PrimaryDigest) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (primaryDigest *L2Tpv2_Counters_Control_TunnelXr_Authentication_PrimaryDigest) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (primaryDigest *L2Tpv2_Counters_Control_TunnelXr_Authentication_PrimaryDigest) SetParent(parent types.Entity) { primaryDigest.parent = parent }

func (primaryDigest *L2Tpv2_Counters_Control_TunnelXr_Authentication_PrimaryDigest) GetParent() types.Entity { return primaryDigest.parent }

func (primaryDigest *L2Tpv2_Counters_Control_TunnelXr_Authentication_PrimaryDigest) GetParentYangName() string { return "authentication" }

// L2Tpv2_Counters_Control_TunnelXr_Authentication_SecondaryDigest
// Secondary digest statistics
type L2Tpv2_Counters_Control_TunnelXr_Authentication_SecondaryDigest struct {
    parent types.Entity
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

func (secondaryDigest *L2Tpv2_Counters_Control_TunnelXr_Authentication_SecondaryDigest) GetFilter() yfilter.YFilter { return secondaryDigest.YFilter }

func (secondaryDigest *L2Tpv2_Counters_Control_TunnelXr_Authentication_SecondaryDigest) SetFilter(yf yfilter.YFilter) { secondaryDigest.YFilter = yf }

func (secondaryDigest *L2Tpv2_Counters_Control_TunnelXr_Authentication_SecondaryDigest) GetGoName(yname string) string {
    if yname == "validate" { return "Validate" }
    if yname == "bad-hash" { return "BadHash" }
    if yname == "bad-length" { return "BadLength" }
    if yname == "ignored" { return "Ignored" }
    if yname == "missing" { return "Missing" }
    if yname == "passed" { return "Passed" }
    if yname == "failed" { return "Failed" }
    if yname == "skipped" { return "Skipped" }
    if yname == "generate-response-failures" { return "GenerateResponseFailures" }
    if yname == "unexpected" { return "Unexpected" }
    if yname == "unexpected-zlb" { return "UnexpectedZlb" }
    return ""
}

func (secondaryDigest *L2Tpv2_Counters_Control_TunnelXr_Authentication_SecondaryDigest) GetSegmentPath() string {
    return "secondary-digest"
}

func (secondaryDigest *L2Tpv2_Counters_Control_TunnelXr_Authentication_SecondaryDigest) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (secondaryDigest *L2Tpv2_Counters_Control_TunnelXr_Authentication_SecondaryDigest) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (secondaryDigest *L2Tpv2_Counters_Control_TunnelXr_Authentication_SecondaryDigest) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["validate"] = secondaryDigest.Validate
    leafs["bad-hash"] = secondaryDigest.BadHash
    leafs["bad-length"] = secondaryDigest.BadLength
    leafs["ignored"] = secondaryDigest.Ignored
    leafs["missing"] = secondaryDigest.Missing
    leafs["passed"] = secondaryDigest.Passed
    leafs["failed"] = secondaryDigest.Failed
    leafs["skipped"] = secondaryDigest.Skipped
    leafs["generate-response-failures"] = secondaryDigest.GenerateResponseFailures
    leafs["unexpected"] = secondaryDigest.Unexpected
    leafs["unexpected-zlb"] = secondaryDigest.UnexpectedZlb
    return leafs
}

func (secondaryDigest *L2Tpv2_Counters_Control_TunnelXr_Authentication_SecondaryDigest) GetBundleName() string { return "cisco_ios_xr" }

func (secondaryDigest *L2Tpv2_Counters_Control_TunnelXr_Authentication_SecondaryDigest) GetYangName() string { return "secondary-digest" }

func (secondaryDigest *L2Tpv2_Counters_Control_TunnelXr_Authentication_SecondaryDigest) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (secondaryDigest *L2Tpv2_Counters_Control_TunnelXr_Authentication_SecondaryDigest) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (secondaryDigest *L2Tpv2_Counters_Control_TunnelXr_Authentication_SecondaryDigest) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (secondaryDigest *L2Tpv2_Counters_Control_TunnelXr_Authentication_SecondaryDigest) SetParent(parent types.Entity) { secondaryDigest.parent = parent }

func (secondaryDigest *L2Tpv2_Counters_Control_TunnelXr_Authentication_SecondaryDigest) GetParent() types.Entity { return secondaryDigest.parent }

func (secondaryDigest *L2Tpv2_Counters_Control_TunnelXr_Authentication_SecondaryDigest) GetParentYangName() string { return "authentication" }

// L2Tpv2_Counters_Control_TunnelXr_Authentication_IntegrityCheck
// Integrity check statistics
type L2Tpv2_Counters_Control_TunnelXr_Authentication_IntegrityCheck struct {
    parent types.Entity
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

func (integrityCheck *L2Tpv2_Counters_Control_TunnelXr_Authentication_IntegrityCheck) GetFilter() yfilter.YFilter { return integrityCheck.YFilter }

func (integrityCheck *L2Tpv2_Counters_Control_TunnelXr_Authentication_IntegrityCheck) SetFilter(yf yfilter.YFilter) { integrityCheck.YFilter = yf }

func (integrityCheck *L2Tpv2_Counters_Control_TunnelXr_Authentication_IntegrityCheck) GetGoName(yname string) string {
    if yname == "validate" { return "Validate" }
    if yname == "bad-hash" { return "BadHash" }
    if yname == "bad-length" { return "BadLength" }
    if yname == "ignored" { return "Ignored" }
    if yname == "missing" { return "Missing" }
    if yname == "passed" { return "Passed" }
    if yname == "failed" { return "Failed" }
    if yname == "skipped" { return "Skipped" }
    if yname == "generate-response-failures" { return "GenerateResponseFailures" }
    if yname == "unexpected" { return "Unexpected" }
    if yname == "unexpected-zlb" { return "UnexpectedZlb" }
    return ""
}

func (integrityCheck *L2Tpv2_Counters_Control_TunnelXr_Authentication_IntegrityCheck) GetSegmentPath() string {
    return "integrity-check"
}

func (integrityCheck *L2Tpv2_Counters_Control_TunnelXr_Authentication_IntegrityCheck) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (integrityCheck *L2Tpv2_Counters_Control_TunnelXr_Authentication_IntegrityCheck) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (integrityCheck *L2Tpv2_Counters_Control_TunnelXr_Authentication_IntegrityCheck) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["validate"] = integrityCheck.Validate
    leafs["bad-hash"] = integrityCheck.BadHash
    leafs["bad-length"] = integrityCheck.BadLength
    leafs["ignored"] = integrityCheck.Ignored
    leafs["missing"] = integrityCheck.Missing
    leafs["passed"] = integrityCheck.Passed
    leafs["failed"] = integrityCheck.Failed
    leafs["skipped"] = integrityCheck.Skipped
    leafs["generate-response-failures"] = integrityCheck.GenerateResponseFailures
    leafs["unexpected"] = integrityCheck.Unexpected
    leafs["unexpected-zlb"] = integrityCheck.UnexpectedZlb
    return leafs
}

func (integrityCheck *L2Tpv2_Counters_Control_TunnelXr_Authentication_IntegrityCheck) GetBundleName() string { return "cisco_ios_xr" }

func (integrityCheck *L2Tpv2_Counters_Control_TunnelXr_Authentication_IntegrityCheck) GetYangName() string { return "integrity-check" }

func (integrityCheck *L2Tpv2_Counters_Control_TunnelXr_Authentication_IntegrityCheck) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (integrityCheck *L2Tpv2_Counters_Control_TunnelXr_Authentication_IntegrityCheck) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (integrityCheck *L2Tpv2_Counters_Control_TunnelXr_Authentication_IntegrityCheck) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (integrityCheck *L2Tpv2_Counters_Control_TunnelXr_Authentication_IntegrityCheck) SetParent(parent types.Entity) { integrityCheck.parent = parent }

func (integrityCheck *L2Tpv2_Counters_Control_TunnelXr_Authentication_IntegrityCheck) GetParent() types.Entity { return integrityCheck.parent }

func (integrityCheck *L2Tpv2_Counters_Control_TunnelXr_Authentication_IntegrityCheck) GetParentYangName() string { return "authentication" }

// L2Tpv2_Counters_Control_TunnelXr_Authentication_LocalSecret
// Local secret statistics
type L2Tpv2_Counters_Control_TunnelXr_Authentication_LocalSecret struct {
    parent types.Entity
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

func (localSecret *L2Tpv2_Counters_Control_TunnelXr_Authentication_LocalSecret) GetFilter() yfilter.YFilter { return localSecret.YFilter }

func (localSecret *L2Tpv2_Counters_Control_TunnelXr_Authentication_LocalSecret) SetFilter(yf yfilter.YFilter) { localSecret.YFilter = yf }

func (localSecret *L2Tpv2_Counters_Control_TunnelXr_Authentication_LocalSecret) GetGoName(yname string) string {
    if yname == "validate" { return "Validate" }
    if yname == "bad-hash" { return "BadHash" }
    if yname == "bad-length" { return "BadLength" }
    if yname == "ignored" { return "Ignored" }
    if yname == "missing" { return "Missing" }
    if yname == "passed" { return "Passed" }
    if yname == "failed" { return "Failed" }
    if yname == "skipped" { return "Skipped" }
    if yname == "generate-response-failures" { return "GenerateResponseFailures" }
    if yname == "unexpected" { return "Unexpected" }
    if yname == "unexpected-zlb" { return "UnexpectedZlb" }
    return ""
}

func (localSecret *L2Tpv2_Counters_Control_TunnelXr_Authentication_LocalSecret) GetSegmentPath() string {
    return "local-secret"
}

func (localSecret *L2Tpv2_Counters_Control_TunnelXr_Authentication_LocalSecret) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (localSecret *L2Tpv2_Counters_Control_TunnelXr_Authentication_LocalSecret) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (localSecret *L2Tpv2_Counters_Control_TunnelXr_Authentication_LocalSecret) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["validate"] = localSecret.Validate
    leafs["bad-hash"] = localSecret.BadHash
    leafs["bad-length"] = localSecret.BadLength
    leafs["ignored"] = localSecret.Ignored
    leafs["missing"] = localSecret.Missing
    leafs["passed"] = localSecret.Passed
    leafs["failed"] = localSecret.Failed
    leafs["skipped"] = localSecret.Skipped
    leafs["generate-response-failures"] = localSecret.GenerateResponseFailures
    leafs["unexpected"] = localSecret.Unexpected
    leafs["unexpected-zlb"] = localSecret.UnexpectedZlb
    return leafs
}

func (localSecret *L2Tpv2_Counters_Control_TunnelXr_Authentication_LocalSecret) GetBundleName() string { return "cisco_ios_xr" }

func (localSecret *L2Tpv2_Counters_Control_TunnelXr_Authentication_LocalSecret) GetYangName() string { return "local-secret" }

func (localSecret *L2Tpv2_Counters_Control_TunnelXr_Authentication_LocalSecret) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (localSecret *L2Tpv2_Counters_Control_TunnelXr_Authentication_LocalSecret) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (localSecret *L2Tpv2_Counters_Control_TunnelXr_Authentication_LocalSecret) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (localSecret *L2Tpv2_Counters_Control_TunnelXr_Authentication_LocalSecret) SetParent(parent types.Entity) { localSecret.parent = parent }

func (localSecret *L2Tpv2_Counters_Control_TunnelXr_Authentication_LocalSecret) GetParent() types.Entity { return localSecret.parent }

func (localSecret *L2Tpv2_Counters_Control_TunnelXr_Authentication_LocalSecret) GetParentYangName() string { return "authentication" }

// L2Tpv2_Counters_Control_TunnelXr_Authentication_ChallengeAvp
// Challenge AVP statistics
type L2Tpv2_Counters_Control_TunnelXr_Authentication_ChallengeAvp struct {
    parent types.Entity
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

func (challengeAvp *L2Tpv2_Counters_Control_TunnelXr_Authentication_ChallengeAvp) GetFilter() yfilter.YFilter { return challengeAvp.YFilter }

func (challengeAvp *L2Tpv2_Counters_Control_TunnelXr_Authentication_ChallengeAvp) SetFilter(yf yfilter.YFilter) { challengeAvp.YFilter = yf }

func (challengeAvp *L2Tpv2_Counters_Control_TunnelXr_Authentication_ChallengeAvp) GetGoName(yname string) string {
    if yname == "validate" { return "Validate" }
    if yname == "bad-hash" { return "BadHash" }
    if yname == "bad-length" { return "BadLength" }
    if yname == "ignored" { return "Ignored" }
    if yname == "missing" { return "Missing" }
    if yname == "passed" { return "Passed" }
    if yname == "failed" { return "Failed" }
    if yname == "skipped" { return "Skipped" }
    if yname == "generate-response-failures" { return "GenerateResponseFailures" }
    if yname == "unexpected" { return "Unexpected" }
    if yname == "unexpected-zlb" { return "UnexpectedZlb" }
    return ""
}

func (challengeAvp *L2Tpv2_Counters_Control_TunnelXr_Authentication_ChallengeAvp) GetSegmentPath() string {
    return "challenge-avp"
}

func (challengeAvp *L2Tpv2_Counters_Control_TunnelXr_Authentication_ChallengeAvp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (challengeAvp *L2Tpv2_Counters_Control_TunnelXr_Authentication_ChallengeAvp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (challengeAvp *L2Tpv2_Counters_Control_TunnelXr_Authentication_ChallengeAvp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["validate"] = challengeAvp.Validate
    leafs["bad-hash"] = challengeAvp.BadHash
    leafs["bad-length"] = challengeAvp.BadLength
    leafs["ignored"] = challengeAvp.Ignored
    leafs["missing"] = challengeAvp.Missing
    leafs["passed"] = challengeAvp.Passed
    leafs["failed"] = challengeAvp.Failed
    leafs["skipped"] = challengeAvp.Skipped
    leafs["generate-response-failures"] = challengeAvp.GenerateResponseFailures
    leafs["unexpected"] = challengeAvp.Unexpected
    leafs["unexpected-zlb"] = challengeAvp.UnexpectedZlb
    return leafs
}

func (challengeAvp *L2Tpv2_Counters_Control_TunnelXr_Authentication_ChallengeAvp) GetBundleName() string { return "cisco_ios_xr" }

func (challengeAvp *L2Tpv2_Counters_Control_TunnelXr_Authentication_ChallengeAvp) GetYangName() string { return "challenge-avp" }

func (challengeAvp *L2Tpv2_Counters_Control_TunnelXr_Authentication_ChallengeAvp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (challengeAvp *L2Tpv2_Counters_Control_TunnelXr_Authentication_ChallengeAvp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (challengeAvp *L2Tpv2_Counters_Control_TunnelXr_Authentication_ChallengeAvp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (challengeAvp *L2Tpv2_Counters_Control_TunnelXr_Authentication_ChallengeAvp) SetParent(parent types.Entity) { challengeAvp.parent = parent }

func (challengeAvp *L2Tpv2_Counters_Control_TunnelXr_Authentication_ChallengeAvp) GetParent() types.Entity { return challengeAvp.parent }

func (challengeAvp *L2Tpv2_Counters_Control_TunnelXr_Authentication_ChallengeAvp) GetParentYangName() string { return "authentication" }

// L2Tpv2_Counters_Control_TunnelXr_Authentication_ChallengeReponse
// Challenge response statistics
type L2Tpv2_Counters_Control_TunnelXr_Authentication_ChallengeReponse struct {
    parent types.Entity
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

func (challengeReponse *L2Tpv2_Counters_Control_TunnelXr_Authentication_ChallengeReponse) GetFilter() yfilter.YFilter { return challengeReponse.YFilter }

func (challengeReponse *L2Tpv2_Counters_Control_TunnelXr_Authentication_ChallengeReponse) SetFilter(yf yfilter.YFilter) { challengeReponse.YFilter = yf }

func (challengeReponse *L2Tpv2_Counters_Control_TunnelXr_Authentication_ChallengeReponse) GetGoName(yname string) string {
    if yname == "validate" { return "Validate" }
    if yname == "bad-hash" { return "BadHash" }
    if yname == "bad-length" { return "BadLength" }
    if yname == "ignored" { return "Ignored" }
    if yname == "missing" { return "Missing" }
    if yname == "passed" { return "Passed" }
    if yname == "failed" { return "Failed" }
    if yname == "skipped" { return "Skipped" }
    if yname == "generate-response-failures" { return "GenerateResponseFailures" }
    if yname == "unexpected" { return "Unexpected" }
    if yname == "unexpected-zlb" { return "UnexpectedZlb" }
    return ""
}

func (challengeReponse *L2Tpv2_Counters_Control_TunnelXr_Authentication_ChallengeReponse) GetSegmentPath() string {
    return "challenge-reponse"
}

func (challengeReponse *L2Tpv2_Counters_Control_TunnelXr_Authentication_ChallengeReponse) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (challengeReponse *L2Tpv2_Counters_Control_TunnelXr_Authentication_ChallengeReponse) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (challengeReponse *L2Tpv2_Counters_Control_TunnelXr_Authentication_ChallengeReponse) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["validate"] = challengeReponse.Validate
    leafs["bad-hash"] = challengeReponse.BadHash
    leafs["bad-length"] = challengeReponse.BadLength
    leafs["ignored"] = challengeReponse.Ignored
    leafs["missing"] = challengeReponse.Missing
    leafs["passed"] = challengeReponse.Passed
    leafs["failed"] = challengeReponse.Failed
    leafs["skipped"] = challengeReponse.Skipped
    leafs["generate-response-failures"] = challengeReponse.GenerateResponseFailures
    leafs["unexpected"] = challengeReponse.Unexpected
    leafs["unexpected-zlb"] = challengeReponse.UnexpectedZlb
    return leafs
}

func (challengeReponse *L2Tpv2_Counters_Control_TunnelXr_Authentication_ChallengeReponse) GetBundleName() string { return "cisco_ios_xr" }

func (challengeReponse *L2Tpv2_Counters_Control_TunnelXr_Authentication_ChallengeReponse) GetYangName() string { return "challenge-reponse" }

func (challengeReponse *L2Tpv2_Counters_Control_TunnelXr_Authentication_ChallengeReponse) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (challengeReponse *L2Tpv2_Counters_Control_TunnelXr_Authentication_ChallengeReponse) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (challengeReponse *L2Tpv2_Counters_Control_TunnelXr_Authentication_ChallengeReponse) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (challengeReponse *L2Tpv2_Counters_Control_TunnelXr_Authentication_ChallengeReponse) SetParent(parent types.Entity) { challengeReponse.parent = parent }

func (challengeReponse *L2Tpv2_Counters_Control_TunnelXr_Authentication_ChallengeReponse) GetParent() types.Entity { return challengeReponse.parent }

func (challengeReponse *L2Tpv2_Counters_Control_TunnelXr_Authentication_ChallengeReponse) GetParentYangName() string { return "authentication" }

// L2Tpv2_Counters_Control_TunnelXr_Authentication_OverallStatistics
// Overall statistics
type L2Tpv2_Counters_Control_TunnelXr_Authentication_OverallStatistics struct {
    parent types.Entity
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

func (overallStatistics *L2Tpv2_Counters_Control_TunnelXr_Authentication_OverallStatistics) GetFilter() yfilter.YFilter { return overallStatistics.YFilter }

func (overallStatistics *L2Tpv2_Counters_Control_TunnelXr_Authentication_OverallStatistics) SetFilter(yf yfilter.YFilter) { overallStatistics.YFilter = yf }

func (overallStatistics *L2Tpv2_Counters_Control_TunnelXr_Authentication_OverallStatistics) GetGoName(yname string) string {
    if yname == "validate" { return "Validate" }
    if yname == "bad-hash" { return "BadHash" }
    if yname == "bad-length" { return "BadLength" }
    if yname == "ignored" { return "Ignored" }
    if yname == "missing" { return "Missing" }
    if yname == "passed" { return "Passed" }
    if yname == "failed" { return "Failed" }
    if yname == "skipped" { return "Skipped" }
    if yname == "generate-response-failures" { return "GenerateResponseFailures" }
    if yname == "unexpected" { return "Unexpected" }
    if yname == "unexpected-zlb" { return "UnexpectedZlb" }
    return ""
}

func (overallStatistics *L2Tpv2_Counters_Control_TunnelXr_Authentication_OverallStatistics) GetSegmentPath() string {
    return "overall-statistics"
}

func (overallStatistics *L2Tpv2_Counters_Control_TunnelXr_Authentication_OverallStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (overallStatistics *L2Tpv2_Counters_Control_TunnelXr_Authentication_OverallStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (overallStatistics *L2Tpv2_Counters_Control_TunnelXr_Authentication_OverallStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["validate"] = overallStatistics.Validate
    leafs["bad-hash"] = overallStatistics.BadHash
    leafs["bad-length"] = overallStatistics.BadLength
    leafs["ignored"] = overallStatistics.Ignored
    leafs["missing"] = overallStatistics.Missing
    leafs["passed"] = overallStatistics.Passed
    leafs["failed"] = overallStatistics.Failed
    leafs["skipped"] = overallStatistics.Skipped
    leafs["generate-response-failures"] = overallStatistics.GenerateResponseFailures
    leafs["unexpected"] = overallStatistics.Unexpected
    leafs["unexpected-zlb"] = overallStatistics.UnexpectedZlb
    return leafs
}

func (overallStatistics *L2Tpv2_Counters_Control_TunnelXr_Authentication_OverallStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (overallStatistics *L2Tpv2_Counters_Control_TunnelXr_Authentication_OverallStatistics) GetYangName() string { return "overall-statistics" }

func (overallStatistics *L2Tpv2_Counters_Control_TunnelXr_Authentication_OverallStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (overallStatistics *L2Tpv2_Counters_Control_TunnelXr_Authentication_OverallStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (overallStatistics *L2Tpv2_Counters_Control_TunnelXr_Authentication_OverallStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (overallStatistics *L2Tpv2_Counters_Control_TunnelXr_Authentication_OverallStatistics) SetParent(parent types.Entity) { overallStatistics.parent = parent }

func (overallStatistics *L2Tpv2_Counters_Control_TunnelXr_Authentication_OverallStatistics) GetParent() types.Entity { return overallStatistics.parent }

func (overallStatistics *L2Tpv2_Counters_Control_TunnelXr_Authentication_OverallStatistics) GetParentYangName() string { return "authentication" }

// L2Tpv2_Counters_Control_TunnelXr_Global
// Tunnel counters
type L2Tpv2_Counters_Control_TunnelXr_Global struct {
    parent types.Entity
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

func (global *L2Tpv2_Counters_Control_TunnelXr_Global) GetFilter() yfilter.YFilter { return global.YFilter }

func (global *L2Tpv2_Counters_Control_TunnelXr_Global) SetFilter(yf yfilter.YFilter) { global.YFilter = yf }

func (global *L2Tpv2_Counters_Control_TunnelXr_Global) GetGoName(yname string) string {
    if yname == "total-transmit" { return "TotalTransmit" }
    if yname == "total-retransmit" { return "TotalRetransmit" }
    if yname == "total-received" { return "TotalReceived" }
    if yname == "total-drop" { return "TotalDrop" }
    if yname == "transmit" { return "Transmit" }
    if yname == "retransmit" { return "Retransmit" }
    if yname == "received" { return "Received" }
    if yname == "drop" { return "Drop" }
    return ""
}

func (global *L2Tpv2_Counters_Control_TunnelXr_Global) GetSegmentPath() string {
    return "global"
}

func (global *L2Tpv2_Counters_Control_TunnelXr_Global) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "transmit" {
        return &global.Transmit
    }
    if childYangName == "retransmit" {
        return &global.Retransmit
    }
    if childYangName == "received" {
        return &global.Received
    }
    if childYangName == "drop" {
        return &global.Drop
    }
    return nil
}

func (global *L2Tpv2_Counters_Control_TunnelXr_Global) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["transmit"] = &global.Transmit
    children["retransmit"] = &global.Retransmit
    children["received"] = &global.Received
    children["drop"] = &global.Drop
    return children
}

func (global *L2Tpv2_Counters_Control_TunnelXr_Global) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["total-transmit"] = global.TotalTransmit
    leafs["total-retransmit"] = global.TotalRetransmit
    leafs["total-received"] = global.TotalReceived
    leafs["total-drop"] = global.TotalDrop
    return leafs
}

func (global *L2Tpv2_Counters_Control_TunnelXr_Global) GetBundleName() string { return "cisco_ios_xr" }

func (global *L2Tpv2_Counters_Control_TunnelXr_Global) GetYangName() string { return "global" }

func (global *L2Tpv2_Counters_Control_TunnelXr_Global) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (global *L2Tpv2_Counters_Control_TunnelXr_Global) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (global *L2Tpv2_Counters_Control_TunnelXr_Global) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (global *L2Tpv2_Counters_Control_TunnelXr_Global) SetParent(parent types.Entity) { global.parent = parent }

func (global *L2Tpv2_Counters_Control_TunnelXr_Global) GetParent() types.Entity { return global.parent }

func (global *L2Tpv2_Counters_Control_TunnelXr_Global) GetParentYangName() string { return "tunnel-xr" }

// L2Tpv2_Counters_Control_TunnelXr_Global_Transmit
// Transmit data
type L2Tpv2_Counters_Control_TunnelXr_Global_Transmit struct {
    parent types.Entity
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

func (transmit *L2Tpv2_Counters_Control_TunnelXr_Global_Transmit) GetFilter() yfilter.YFilter { return transmit.YFilter }

func (transmit *L2Tpv2_Counters_Control_TunnelXr_Global_Transmit) SetFilter(yf yfilter.YFilter) { transmit.YFilter = yf }

func (transmit *L2Tpv2_Counters_Control_TunnelXr_Global_Transmit) GetGoName(yname string) string {
    if yname == "unknown-packets" { return "UnknownPackets" }
    if yname == "zero-length-body-packets" { return "ZeroLengthBodyPackets" }
    if yname == "start-control-connection-requests" { return "StartControlConnectionRequests" }
    if yname == "start-control-connection-replies" { return "StartControlConnectionReplies" }
    if yname == "start-control-connection-notifications" { return "StartControlConnectionNotifications" }
    if yname == "stop-control-connection-notifications" { return "StopControlConnectionNotifications" }
    if yname == "hello-packets" { return "HelloPackets" }
    if yname == "outgoing-call-requests" { return "OutgoingCallRequests" }
    if yname == "outgoing-call-replies" { return "OutgoingCallReplies" }
    if yname == "outgoing-call-connected-packets" { return "OutgoingCallConnectedPackets" }
    if yname == "incoming-call-requests" { return "IncomingCallRequests" }
    if yname == "incoming-call-replies" { return "IncomingCallReplies" }
    if yname == "incoming-call-connected-packets" { return "IncomingCallConnectedPackets" }
    if yname == "call-disconnect-notify-packets" { return "CallDisconnectNotifyPackets" }
    if yname == "wan-error-notify-packets" { return "WanErrorNotifyPackets" }
    if yname == "set-link-info-packets" { return "SetLinkInfoPackets" }
    if yname == "service-relay-requests" { return "ServiceRelayRequests" }
    if yname == "service-relay-replies" { return "ServiceRelayReplies" }
    if yname == "acknowledgement-packets" { return "AcknowledgementPackets" }
    return ""
}

func (transmit *L2Tpv2_Counters_Control_TunnelXr_Global_Transmit) GetSegmentPath() string {
    return "transmit"
}

func (transmit *L2Tpv2_Counters_Control_TunnelXr_Global_Transmit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (transmit *L2Tpv2_Counters_Control_TunnelXr_Global_Transmit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (transmit *L2Tpv2_Counters_Control_TunnelXr_Global_Transmit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["unknown-packets"] = transmit.UnknownPackets
    leafs["zero-length-body-packets"] = transmit.ZeroLengthBodyPackets
    leafs["start-control-connection-requests"] = transmit.StartControlConnectionRequests
    leafs["start-control-connection-replies"] = transmit.StartControlConnectionReplies
    leafs["start-control-connection-notifications"] = transmit.StartControlConnectionNotifications
    leafs["stop-control-connection-notifications"] = transmit.StopControlConnectionNotifications
    leafs["hello-packets"] = transmit.HelloPackets
    leafs["outgoing-call-requests"] = transmit.OutgoingCallRequests
    leafs["outgoing-call-replies"] = transmit.OutgoingCallReplies
    leafs["outgoing-call-connected-packets"] = transmit.OutgoingCallConnectedPackets
    leafs["incoming-call-requests"] = transmit.IncomingCallRequests
    leafs["incoming-call-replies"] = transmit.IncomingCallReplies
    leafs["incoming-call-connected-packets"] = transmit.IncomingCallConnectedPackets
    leafs["call-disconnect-notify-packets"] = transmit.CallDisconnectNotifyPackets
    leafs["wan-error-notify-packets"] = transmit.WanErrorNotifyPackets
    leafs["set-link-info-packets"] = transmit.SetLinkInfoPackets
    leafs["service-relay-requests"] = transmit.ServiceRelayRequests
    leafs["service-relay-replies"] = transmit.ServiceRelayReplies
    leafs["acknowledgement-packets"] = transmit.AcknowledgementPackets
    return leafs
}

func (transmit *L2Tpv2_Counters_Control_TunnelXr_Global_Transmit) GetBundleName() string { return "cisco_ios_xr" }

func (transmit *L2Tpv2_Counters_Control_TunnelXr_Global_Transmit) GetYangName() string { return "transmit" }

func (transmit *L2Tpv2_Counters_Control_TunnelXr_Global_Transmit) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (transmit *L2Tpv2_Counters_Control_TunnelXr_Global_Transmit) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (transmit *L2Tpv2_Counters_Control_TunnelXr_Global_Transmit) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (transmit *L2Tpv2_Counters_Control_TunnelXr_Global_Transmit) SetParent(parent types.Entity) { transmit.parent = parent }

func (transmit *L2Tpv2_Counters_Control_TunnelXr_Global_Transmit) GetParent() types.Entity { return transmit.parent }

func (transmit *L2Tpv2_Counters_Control_TunnelXr_Global_Transmit) GetParentYangName() string { return "global" }

// L2Tpv2_Counters_Control_TunnelXr_Global_Retransmit
// Re transmit data
type L2Tpv2_Counters_Control_TunnelXr_Global_Retransmit struct {
    parent types.Entity
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

func (retransmit *L2Tpv2_Counters_Control_TunnelXr_Global_Retransmit) GetFilter() yfilter.YFilter { return retransmit.YFilter }

func (retransmit *L2Tpv2_Counters_Control_TunnelXr_Global_Retransmit) SetFilter(yf yfilter.YFilter) { retransmit.YFilter = yf }

func (retransmit *L2Tpv2_Counters_Control_TunnelXr_Global_Retransmit) GetGoName(yname string) string {
    if yname == "unknown-packets" { return "UnknownPackets" }
    if yname == "zero-length-body-packets" { return "ZeroLengthBodyPackets" }
    if yname == "start-control-connection-requests" { return "StartControlConnectionRequests" }
    if yname == "start-control-connection-replies" { return "StartControlConnectionReplies" }
    if yname == "start-control-connection-notifications" { return "StartControlConnectionNotifications" }
    if yname == "stop-control-connection-notifications" { return "StopControlConnectionNotifications" }
    if yname == "hello-packets" { return "HelloPackets" }
    if yname == "outgoing-call-requests" { return "OutgoingCallRequests" }
    if yname == "outgoing-call-replies" { return "OutgoingCallReplies" }
    if yname == "outgoing-call-connected-packets" { return "OutgoingCallConnectedPackets" }
    if yname == "incoming-call-requests" { return "IncomingCallRequests" }
    if yname == "incoming-call-replies" { return "IncomingCallReplies" }
    if yname == "incoming-call-connected-packets" { return "IncomingCallConnectedPackets" }
    if yname == "call-disconnect-notify-packets" { return "CallDisconnectNotifyPackets" }
    if yname == "wan-error-notify-packets" { return "WanErrorNotifyPackets" }
    if yname == "set-link-info-packets" { return "SetLinkInfoPackets" }
    if yname == "service-relay-requests" { return "ServiceRelayRequests" }
    if yname == "service-relay-replies" { return "ServiceRelayReplies" }
    if yname == "acknowledgement-packets" { return "AcknowledgementPackets" }
    return ""
}

func (retransmit *L2Tpv2_Counters_Control_TunnelXr_Global_Retransmit) GetSegmentPath() string {
    return "retransmit"
}

func (retransmit *L2Tpv2_Counters_Control_TunnelXr_Global_Retransmit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (retransmit *L2Tpv2_Counters_Control_TunnelXr_Global_Retransmit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (retransmit *L2Tpv2_Counters_Control_TunnelXr_Global_Retransmit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["unknown-packets"] = retransmit.UnknownPackets
    leafs["zero-length-body-packets"] = retransmit.ZeroLengthBodyPackets
    leafs["start-control-connection-requests"] = retransmit.StartControlConnectionRequests
    leafs["start-control-connection-replies"] = retransmit.StartControlConnectionReplies
    leafs["start-control-connection-notifications"] = retransmit.StartControlConnectionNotifications
    leafs["stop-control-connection-notifications"] = retransmit.StopControlConnectionNotifications
    leafs["hello-packets"] = retransmit.HelloPackets
    leafs["outgoing-call-requests"] = retransmit.OutgoingCallRequests
    leafs["outgoing-call-replies"] = retransmit.OutgoingCallReplies
    leafs["outgoing-call-connected-packets"] = retransmit.OutgoingCallConnectedPackets
    leafs["incoming-call-requests"] = retransmit.IncomingCallRequests
    leafs["incoming-call-replies"] = retransmit.IncomingCallReplies
    leafs["incoming-call-connected-packets"] = retransmit.IncomingCallConnectedPackets
    leafs["call-disconnect-notify-packets"] = retransmit.CallDisconnectNotifyPackets
    leafs["wan-error-notify-packets"] = retransmit.WanErrorNotifyPackets
    leafs["set-link-info-packets"] = retransmit.SetLinkInfoPackets
    leafs["service-relay-requests"] = retransmit.ServiceRelayRequests
    leafs["service-relay-replies"] = retransmit.ServiceRelayReplies
    leafs["acknowledgement-packets"] = retransmit.AcknowledgementPackets
    return leafs
}

func (retransmit *L2Tpv2_Counters_Control_TunnelXr_Global_Retransmit) GetBundleName() string { return "cisco_ios_xr" }

func (retransmit *L2Tpv2_Counters_Control_TunnelXr_Global_Retransmit) GetYangName() string { return "retransmit" }

func (retransmit *L2Tpv2_Counters_Control_TunnelXr_Global_Retransmit) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (retransmit *L2Tpv2_Counters_Control_TunnelXr_Global_Retransmit) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (retransmit *L2Tpv2_Counters_Control_TunnelXr_Global_Retransmit) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (retransmit *L2Tpv2_Counters_Control_TunnelXr_Global_Retransmit) SetParent(parent types.Entity) { retransmit.parent = parent }

func (retransmit *L2Tpv2_Counters_Control_TunnelXr_Global_Retransmit) GetParent() types.Entity { return retransmit.parent }

func (retransmit *L2Tpv2_Counters_Control_TunnelXr_Global_Retransmit) GetParentYangName() string { return "global" }

// L2Tpv2_Counters_Control_TunnelXr_Global_Received
// Received data
type L2Tpv2_Counters_Control_TunnelXr_Global_Received struct {
    parent types.Entity
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

func (received *L2Tpv2_Counters_Control_TunnelXr_Global_Received) GetFilter() yfilter.YFilter { return received.YFilter }

func (received *L2Tpv2_Counters_Control_TunnelXr_Global_Received) SetFilter(yf yfilter.YFilter) { received.YFilter = yf }

func (received *L2Tpv2_Counters_Control_TunnelXr_Global_Received) GetGoName(yname string) string {
    if yname == "unknown-packets" { return "UnknownPackets" }
    if yname == "zero-length-body-packets" { return "ZeroLengthBodyPackets" }
    if yname == "start-control-connection-requests" { return "StartControlConnectionRequests" }
    if yname == "start-control-connection-replies" { return "StartControlConnectionReplies" }
    if yname == "start-control-connection-notifications" { return "StartControlConnectionNotifications" }
    if yname == "stop-control-connection-notifications" { return "StopControlConnectionNotifications" }
    if yname == "hello-packets" { return "HelloPackets" }
    if yname == "outgoing-call-requests" { return "OutgoingCallRequests" }
    if yname == "outgoing-call-replies" { return "OutgoingCallReplies" }
    if yname == "outgoing-call-connected-packets" { return "OutgoingCallConnectedPackets" }
    if yname == "incoming-call-requests" { return "IncomingCallRequests" }
    if yname == "incoming-call-replies" { return "IncomingCallReplies" }
    if yname == "incoming-call-connected-packets" { return "IncomingCallConnectedPackets" }
    if yname == "call-disconnect-notify-packets" { return "CallDisconnectNotifyPackets" }
    if yname == "wan-error-notify-packets" { return "WanErrorNotifyPackets" }
    if yname == "set-link-info-packets" { return "SetLinkInfoPackets" }
    if yname == "service-relay-requests" { return "ServiceRelayRequests" }
    if yname == "service-relay-replies" { return "ServiceRelayReplies" }
    if yname == "acknowledgement-packets" { return "AcknowledgementPackets" }
    return ""
}

func (received *L2Tpv2_Counters_Control_TunnelXr_Global_Received) GetSegmentPath() string {
    return "received"
}

func (received *L2Tpv2_Counters_Control_TunnelXr_Global_Received) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (received *L2Tpv2_Counters_Control_TunnelXr_Global_Received) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (received *L2Tpv2_Counters_Control_TunnelXr_Global_Received) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["unknown-packets"] = received.UnknownPackets
    leafs["zero-length-body-packets"] = received.ZeroLengthBodyPackets
    leafs["start-control-connection-requests"] = received.StartControlConnectionRequests
    leafs["start-control-connection-replies"] = received.StartControlConnectionReplies
    leafs["start-control-connection-notifications"] = received.StartControlConnectionNotifications
    leafs["stop-control-connection-notifications"] = received.StopControlConnectionNotifications
    leafs["hello-packets"] = received.HelloPackets
    leafs["outgoing-call-requests"] = received.OutgoingCallRequests
    leafs["outgoing-call-replies"] = received.OutgoingCallReplies
    leafs["outgoing-call-connected-packets"] = received.OutgoingCallConnectedPackets
    leafs["incoming-call-requests"] = received.IncomingCallRequests
    leafs["incoming-call-replies"] = received.IncomingCallReplies
    leafs["incoming-call-connected-packets"] = received.IncomingCallConnectedPackets
    leafs["call-disconnect-notify-packets"] = received.CallDisconnectNotifyPackets
    leafs["wan-error-notify-packets"] = received.WanErrorNotifyPackets
    leafs["set-link-info-packets"] = received.SetLinkInfoPackets
    leafs["service-relay-requests"] = received.ServiceRelayRequests
    leafs["service-relay-replies"] = received.ServiceRelayReplies
    leafs["acknowledgement-packets"] = received.AcknowledgementPackets
    return leafs
}

func (received *L2Tpv2_Counters_Control_TunnelXr_Global_Received) GetBundleName() string { return "cisco_ios_xr" }

func (received *L2Tpv2_Counters_Control_TunnelXr_Global_Received) GetYangName() string { return "received" }

func (received *L2Tpv2_Counters_Control_TunnelXr_Global_Received) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (received *L2Tpv2_Counters_Control_TunnelXr_Global_Received) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (received *L2Tpv2_Counters_Control_TunnelXr_Global_Received) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (received *L2Tpv2_Counters_Control_TunnelXr_Global_Received) SetParent(parent types.Entity) { received.parent = parent }

func (received *L2Tpv2_Counters_Control_TunnelXr_Global_Received) GetParent() types.Entity { return received.parent }

func (received *L2Tpv2_Counters_Control_TunnelXr_Global_Received) GetParentYangName() string { return "global" }

// L2Tpv2_Counters_Control_TunnelXr_Global_Drop
// Drop data
type L2Tpv2_Counters_Control_TunnelXr_Global_Drop struct {
    parent types.Entity
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

func (drop *L2Tpv2_Counters_Control_TunnelXr_Global_Drop) GetFilter() yfilter.YFilter { return drop.YFilter }

func (drop *L2Tpv2_Counters_Control_TunnelXr_Global_Drop) SetFilter(yf yfilter.YFilter) { drop.YFilter = yf }

func (drop *L2Tpv2_Counters_Control_TunnelXr_Global_Drop) GetGoName(yname string) string {
    if yname == "unknown-packets" { return "UnknownPackets" }
    if yname == "zero-length-body-packets" { return "ZeroLengthBodyPackets" }
    if yname == "start-control-connection-requests" { return "StartControlConnectionRequests" }
    if yname == "start-control-connection-replies" { return "StartControlConnectionReplies" }
    if yname == "start-control-connection-notifications" { return "StartControlConnectionNotifications" }
    if yname == "stop-control-connection-notifications" { return "StopControlConnectionNotifications" }
    if yname == "hello-packets" { return "HelloPackets" }
    if yname == "outgoing-call-requests" { return "OutgoingCallRequests" }
    if yname == "outgoing-call-replies" { return "OutgoingCallReplies" }
    if yname == "outgoing-call-connected-packets" { return "OutgoingCallConnectedPackets" }
    if yname == "incoming-call-requests" { return "IncomingCallRequests" }
    if yname == "incoming-call-replies" { return "IncomingCallReplies" }
    if yname == "incoming-call-connected-packets" { return "IncomingCallConnectedPackets" }
    if yname == "call-disconnect-notify-packets" { return "CallDisconnectNotifyPackets" }
    if yname == "wan-error-notify-packets" { return "WanErrorNotifyPackets" }
    if yname == "set-link-info-packets" { return "SetLinkInfoPackets" }
    if yname == "service-relay-requests" { return "ServiceRelayRequests" }
    if yname == "service-relay-replies" { return "ServiceRelayReplies" }
    if yname == "acknowledgement-packets" { return "AcknowledgementPackets" }
    return ""
}

func (drop *L2Tpv2_Counters_Control_TunnelXr_Global_Drop) GetSegmentPath() string {
    return "drop"
}

func (drop *L2Tpv2_Counters_Control_TunnelXr_Global_Drop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (drop *L2Tpv2_Counters_Control_TunnelXr_Global_Drop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (drop *L2Tpv2_Counters_Control_TunnelXr_Global_Drop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["unknown-packets"] = drop.UnknownPackets
    leafs["zero-length-body-packets"] = drop.ZeroLengthBodyPackets
    leafs["start-control-connection-requests"] = drop.StartControlConnectionRequests
    leafs["start-control-connection-replies"] = drop.StartControlConnectionReplies
    leafs["start-control-connection-notifications"] = drop.StartControlConnectionNotifications
    leafs["stop-control-connection-notifications"] = drop.StopControlConnectionNotifications
    leafs["hello-packets"] = drop.HelloPackets
    leafs["outgoing-call-requests"] = drop.OutgoingCallRequests
    leafs["outgoing-call-replies"] = drop.OutgoingCallReplies
    leafs["outgoing-call-connected-packets"] = drop.OutgoingCallConnectedPackets
    leafs["incoming-call-requests"] = drop.IncomingCallRequests
    leafs["incoming-call-replies"] = drop.IncomingCallReplies
    leafs["incoming-call-connected-packets"] = drop.IncomingCallConnectedPackets
    leafs["call-disconnect-notify-packets"] = drop.CallDisconnectNotifyPackets
    leafs["wan-error-notify-packets"] = drop.WanErrorNotifyPackets
    leafs["set-link-info-packets"] = drop.SetLinkInfoPackets
    leafs["service-relay-requests"] = drop.ServiceRelayRequests
    leafs["service-relay-replies"] = drop.ServiceRelayReplies
    leafs["acknowledgement-packets"] = drop.AcknowledgementPackets
    return leafs
}

func (drop *L2Tpv2_Counters_Control_TunnelXr_Global_Drop) GetBundleName() string { return "cisco_ios_xr" }

func (drop *L2Tpv2_Counters_Control_TunnelXr_Global_Drop) GetYangName() string { return "drop" }

func (drop *L2Tpv2_Counters_Control_TunnelXr_Global_Drop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (drop *L2Tpv2_Counters_Control_TunnelXr_Global_Drop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (drop *L2Tpv2_Counters_Control_TunnelXr_Global_Drop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (drop *L2Tpv2_Counters_Control_TunnelXr_Global_Drop) SetParent(parent types.Entity) { drop.parent = parent }

func (drop *L2Tpv2_Counters_Control_TunnelXr_Global_Drop) GetParent() types.Entity { return drop.parent }

func (drop *L2Tpv2_Counters_Control_TunnelXr_Global_Drop) GetParentYangName() string { return "global" }

// L2Tpv2_Counters_Control_Tunnels
// Table of tunnel IDs of control message counters
type L2Tpv2_Counters_Control_Tunnels struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // L2TP tunnel control message counters. The type is slice of
    // L2Tpv2_Counters_Control_Tunnels_Tunnel.
    Tunnel []L2Tpv2_Counters_Control_Tunnels_Tunnel
}

func (tunnels *L2Tpv2_Counters_Control_Tunnels) GetFilter() yfilter.YFilter { return tunnels.YFilter }

func (tunnels *L2Tpv2_Counters_Control_Tunnels) SetFilter(yf yfilter.YFilter) { tunnels.YFilter = yf }

func (tunnels *L2Tpv2_Counters_Control_Tunnels) GetGoName(yname string) string {
    if yname == "tunnel" { return "Tunnel" }
    return ""
}

func (tunnels *L2Tpv2_Counters_Control_Tunnels) GetSegmentPath() string {
    return "tunnels"
}

func (tunnels *L2Tpv2_Counters_Control_Tunnels) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tunnel" {
        for _, c := range tunnels.Tunnel {
            if tunnels.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := L2Tpv2_Counters_Control_Tunnels_Tunnel{}
        tunnels.Tunnel = append(tunnels.Tunnel, child)
        return &tunnels.Tunnel[len(tunnels.Tunnel)-1]
    }
    return nil
}

func (tunnels *L2Tpv2_Counters_Control_Tunnels) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range tunnels.Tunnel {
        children[tunnels.Tunnel[i].GetSegmentPath()] = &tunnels.Tunnel[i]
    }
    return children
}

func (tunnels *L2Tpv2_Counters_Control_Tunnels) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (tunnels *L2Tpv2_Counters_Control_Tunnels) GetBundleName() string { return "cisco_ios_xr" }

func (tunnels *L2Tpv2_Counters_Control_Tunnels) GetYangName() string { return "tunnels" }

func (tunnels *L2Tpv2_Counters_Control_Tunnels) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tunnels *L2Tpv2_Counters_Control_Tunnels) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tunnels *L2Tpv2_Counters_Control_Tunnels) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tunnels *L2Tpv2_Counters_Control_Tunnels) SetParent(parent types.Entity) { tunnels.parent = parent }

func (tunnels *L2Tpv2_Counters_Control_Tunnels) GetParent() types.Entity { return tunnels.parent }

func (tunnels *L2Tpv2_Counters_Control_Tunnels) GetParentYangName() string { return "control" }

// L2Tpv2_Counters_Control_Tunnels_Tunnel
// L2TP tunnel control message counters
type L2Tpv2_Counters_Control_Tunnels_Tunnel struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. L2TP tunnel ID. The type is interface{} with
    // range: -2147483648..2147483647.
    TunnelId interface{}

    // L2TP control message local and remote addresses.
    Brief L2Tpv2_Counters_Control_Tunnels_Tunnel_Brief

    // Global data.
    Global L2Tpv2_Counters_Control_Tunnels_Tunnel_Global
}

func (tunnel *L2Tpv2_Counters_Control_Tunnels_Tunnel) GetFilter() yfilter.YFilter { return tunnel.YFilter }

func (tunnel *L2Tpv2_Counters_Control_Tunnels_Tunnel) SetFilter(yf yfilter.YFilter) { tunnel.YFilter = yf }

func (tunnel *L2Tpv2_Counters_Control_Tunnels_Tunnel) GetGoName(yname string) string {
    if yname == "tunnel-id" { return "TunnelId" }
    if yname == "brief" { return "Brief" }
    if yname == "global" { return "Global" }
    return ""
}

func (tunnel *L2Tpv2_Counters_Control_Tunnels_Tunnel) GetSegmentPath() string {
    return "tunnel" + "[tunnel-id='" + fmt.Sprintf("%v", tunnel.TunnelId) + "']"
}

func (tunnel *L2Tpv2_Counters_Control_Tunnels_Tunnel) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "brief" {
        return &tunnel.Brief
    }
    if childYangName == "global" {
        return &tunnel.Global
    }
    return nil
}

func (tunnel *L2Tpv2_Counters_Control_Tunnels_Tunnel) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["brief"] = &tunnel.Brief
    children["global"] = &tunnel.Global
    return children
}

func (tunnel *L2Tpv2_Counters_Control_Tunnels_Tunnel) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tunnel-id"] = tunnel.TunnelId
    return leafs
}

func (tunnel *L2Tpv2_Counters_Control_Tunnels_Tunnel) GetBundleName() string { return "cisco_ios_xr" }

func (tunnel *L2Tpv2_Counters_Control_Tunnels_Tunnel) GetYangName() string { return "tunnel" }

func (tunnel *L2Tpv2_Counters_Control_Tunnels_Tunnel) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tunnel *L2Tpv2_Counters_Control_Tunnels_Tunnel) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tunnel *L2Tpv2_Counters_Control_Tunnels_Tunnel) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tunnel *L2Tpv2_Counters_Control_Tunnels_Tunnel) SetParent(parent types.Entity) { tunnel.parent = parent }

func (tunnel *L2Tpv2_Counters_Control_Tunnels_Tunnel) GetParent() types.Entity { return tunnel.parent }

func (tunnel *L2Tpv2_Counters_Control_Tunnels_Tunnel) GetParentYangName() string { return "tunnels" }

// L2Tpv2_Counters_Control_Tunnels_Tunnel_Brief
// L2TP control message local and remote addresses
type L2Tpv2_Counters_Control_Tunnels_Tunnel_Brief struct {
    parent types.Entity
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

func (brief *L2Tpv2_Counters_Control_Tunnels_Tunnel_Brief) GetFilter() yfilter.YFilter { return brief.YFilter }

func (brief *L2Tpv2_Counters_Control_Tunnels_Tunnel_Brief) SetFilter(yf yfilter.YFilter) { brief.YFilter = yf }

func (brief *L2Tpv2_Counters_Control_Tunnels_Tunnel_Brief) GetGoName(yname string) string {
    if yname == "remote-tunnel-id" { return "RemoteTunnelId" }
    if yname == "local-address" { return "LocalAddress" }
    if yname == "remote-address" { return "RemoteAddress" }
    return ""
}

func (brief *L2Tpv2_Counters_Control_Tunnels_Tunnel_Brief) GetSegmentPath() string {
    return "brief"
}

func (brief *L2Tpv2_Counters_Control_Tunnels_Tunnel_Brief) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (brief *L2Tpv2_Counters_Control_Tunnels_Tunnel_Brief) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (brief *L2Tpv2_Counters_Control_Tunnels_Tunnel_Brief) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["remote-tunnel-id"] = brief.RemoteTunnelId
    leafs["local-address"] = brief.LocalAddress
    leafs["remote-address"] = brief.RemoteAddress
    return leafs
}

func (brief *L2Tpv2_Counters_Control_Tunnels_Tunnel_Brief) GetBundleName() string { return "cisco_ios_xr" }

func (brief *L2Tpv2_Counters_Control_Tunnels_Tunnel_Brief) GetYangName() string { return "brief" }

func (brief *L2Tpv2_Counters_Control_Tunnels_Tunnel_Brief) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (brief *L2Tpv2_Counters_Control_Tunnels_Tunnel_Brief) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (brief *L2Tpv2_Counters_Control_Tunnels_Tunnel_Brief) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (brief *L2Tpv2_Counters_Control_Tunnels_Tunnel_Brief) SetParent(parent types.Entity) { brief.parent = parent }

func (brief *L2Tpv2_Counters_Control_Tunnels_Tunnel_Brief) GetParent() types.Entity { return brief.parent }

func (brief *L2Tpv2_Counters_Control_Tunnels_Tunnel_Brief) GetParentYangName() string { return "tunnel" }

// L2Tpv2_Counters_Control_Tunnels_Tunnel_Global
// Global data
type L2Tpv2_Counters_Control_Tunnels_Tunnel_Global struct {
    parent types.Entity
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

func (global *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global) GetFilter() yfilter.YFilter { return global.YFilter }

func (global *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global) SetFilter(yf yfilter.YFilter) { global.YFilter = yf }

func (global *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global) GetGoName(yname string) string {
    if yname == "total-transmit" { return "TotalTransmit" }
    if yname == "total-retransmit" { return "TotalRetransmit" }
    if yname == "total-received" { return "TotalReceived" }
    if yname == "total-drop" { return "TotalDrop" }
    if yname == "transmit" { return "Transmit" }
    if yname == "retransmit" { return "Retransmit" }
    if yname == "received" { return "Received" }
    if yname == "drop" { return "Drop" }
    return ""
}

func (global *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global) GetSegmentPath() string {
    return "global"
}

func (global *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "transmit" {
        return &global.Transmit
    }
    if childYangName == "retransmit" {
        return &global.Retransmit
    }
    if childYangName == "received" {
        return &global.Received
    }
    if childYangName == "drop" {
        return &global.Drop
    }
    return nil
}

func (global *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["transmit"] = &global.Transmit
    children["retransmit"] = &global.Retransmit
    children["received"] = &global.Received
    children["drop"] = &global.Drop
    return children
}

func (global *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["total-transmit"] = global.TotalTransmit
    leafs["total-retransmit"] = global.TotalRetransmit
    leafs["total-received"] = global.TotalReceived
    leafs["total-drop"] = global.TotalDrop
    return leafs
}

func (global *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global) GetBundleName() string { return "cisco_ios_xr" }

func (global *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global) GetYangName() string { return "global" }

func (global *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (global *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (global *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (global *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global) SetParent(parent types.Entity) { global.parent = parent }

func (global *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global) GetParent() types.Entity { return global.parent }

func (global *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global) GetParentYangName() string { return "tunnel" }

// L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Transmit
// Transmit data
type L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Transmit struct {
    parent types.Entity
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

func (transmit *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Transmit) GetFilter() yfilter.YFilter { return transmit.YFilter }

func (transmit *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Transmit) SetFilter(yf yfilter.YFilter) { transmit.YFilter = yf }

func (transmit *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Transmit) GetGoName(yname string) string {
    if yname == "unknown-packets" { return "UnknownPackets" }
    if yname == "zero-length-body-packets" { return "ZeroLengthBodyPackets" }
    if yname == "start-control-connection-requests" { return "StartControlConnectionRequests" }
    if yname == "start-control-connection-replies" { return "StartControlConnectionReplies" }
    if yname == "start-control-connection-notifications" { return "StartControlConnectionNotifications" }
    if yname == "stop-control-connection-notifications" { return "StopControlConnectionNotifications" }
    if yname == "hello-packets" { return "HelloPackets" }
    if yname == "outgoing-call-requests" { return "OutgoingCallRequests" }
    if yname == "outgoing-call-replies" { return "OutgoingCallReplies" }
    if yname == "outgoing-call-connected-packets" { return "OutgoingCallConnectedPackets" }
    if yname == "incoming-call-requests" { return "IncomingCallRequests" }
    if yname == "incoming-call-replies" { return "IncomingCallReplies" }
    if yname == "incoming-call-connected-packets" { return "IncomingCallConnectedPackets" }
    if yname == "call-disconnect-notify-packets" { return "CallDisconnectNotifyPackets" }
    if yname == "wan-error-notify-packets" { return "WanErrorNotifyPackets" }
    if yname == "set-link-info-packets" { return "SetLinkInfoPackets" }
    if yname == "service-relay-requests" { return "ServiceRelayRequests" }
    if yname == "service-relay-replies" { return "ServiceRelayReplies" }
    if yname == "acknowledgement-packets" { return "AcknowledgementPackets" }
    return ""
}

func (transmit *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Transmit) GetSegmentPath() string {
    return "transmit"
}

func (transmit *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Transmit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (transmit *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Transmit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (transmit *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Transmit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["unknown-packets"] = transmit.UnknownPackets
    leafs["zero-length-body-packets"] = transmit.ZeroLengthBodyPackets
    leafs["start-control-connection-requests"] = transmit.StartControlConnectionRequests
    leafs["start-control-connection-replies"] = transmit.StartControlConnectionReplies
    leafs["start-control-connection-notifications"] = transmit.StartControlConnectionNotifications
    leafs["stop-control-connection-notifications"] = transmit.StopControlConnectionNotifications
    leafs["hello-packets"] = transmit.HelloPackets
    leafs["outgoing-call-requests"] = transmit.OutgoingCallRequests
    leafs["outgoing-call-replies"] = transmit.OutgoingCallReplies
    leafs["outgoing-call-connected-packets"] = transmit.OutgoingCallConnectedPackets
    leafs["incoming-call-requests"] = transmit.IncomingCallRequests
    leafs["incoming-call-replies"] = transmit.IncomingCallReplies
    leafs["incoming-call-connected-packets"] = transmit.IncomingCallConnectedPackets
    leafs["call-disconnect-notify-packets"] = transmit.CallDisconnectNotifyPackets
    leafs["wan-error-notify-packets"] = transmit.WanErrorNotifyPackets
    leafs["set-link-info-packets"] = transmit.SetLinkInfoPackets
    leafs["service-relay-requests"] = transmit.ServiceRelayRequests
    leafs["service-relay-replies"] = transmit.ServiceRelayReplies
    leafs["acknowledgement-packets"] = transmit.AcknowledgementPackets
    return leafs
}

func (transmit *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Transmit) GetBundleName() string { return "cisco_ios_xr" }

func (transmit *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Transmit) GetYangName() string { return "transmit" }

func (transmit *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Transmit) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (transmit *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Transmit) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (transmit *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Transmit) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (transmit *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Transmit) SetParent(parent types.Entity) { transmit.parent = parent }

func (transmit *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Transmit) GetParent() types.Entity { return transmit.parent }

func (transmit *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Transmit) GetParentYangName() string { return "global" }

// L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Retransmit
// Re transmit data
type L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Retransmit struct {
    parent types.Entity
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

func (retransmit *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Retransmit) GetFilter() yfilter.YFilter { return retransmit.YFilter }

func (retransmit *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Retransmit) SetFilter(yf yfilter.YFilter) { retransmit.YFilter = yf }

func (retransmit *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Retransmit) GetGoName(yname string) string {
    if yname == "unknown-packets" { return "UnknownPackets" }
    if yname == "zero-length-body-packets" { return "ZeroLengthBodyPackets" }
    if yname == "start-control-connection-requests" { return "StartControlConnectionRequests" }
    if yname == "start-control-connection-replies" { return "StartControlConnectionReplies" }
    if yname == "start-control-connection-notifications" { return "StartControlConnectionNotifications" }
    if yname == "stop-control-connection-notifications" { return "StopControlConnectionNotifications" }
    if yname == "hello-packets" { return "HelloPackets" }
    if yname == "outgoing-call-requests" { return "OutgoingCallRequests" }
    if yname == "outgoing-call-replies" { return "OutgoingCallReplies" }
    if yname == "outgoing-call-connected-packets" { return "OutgoingCallConnectedPackets" }
    if yname == "incoming-call-requests" { return "IncomingCallRequests" }
    if yname == "incoming-call-replies" { return "IncomingCallReplies" }
    if yname == "incoming-call-connected-packets" { return "IncomingCallConnectedPackets" }
    if yname == "call-disconnect-notify-packets" { return "CallDisconnectNotifyPackets" }
    if yname == "wan-error-notify-packets" { return "WanErrorNotifyPackets" }
    if yname == "set-link-info-packets" { return "SetLinkInfoPackets" }
    if yname == "service-relay-requests" { return "ServiceRelayRequests" }
    if yname == "service-relay-replies" { return "ServiceRelayReplies" }
    if yname == "acknowledgement-packets" { return "AcknowledgementPackets" }
    return ""
}

func (retransmit *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Retransmit) GetSegmentPath() string {
    return "retransmit"
}

func (retransmit *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Retransmit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (retransmit *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Retransmit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (retransmit *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Retransmit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["unknown-packets"] = retransmit.UnknownPackets
    leafs["zero-length-body-packets"] = retransmit.ZeroLengthBodyPackets
    leafs["start-control-connection-requests"] = retransmit.StartControlConnectionRequests
    leafs["start-control-connection-replies"] = retransmit.StartControlConnectionReplies
    leafs["start-control-connection-notifications"] = retransmit.StartControlConnectionNotifications
    leafs["stop-control-connection-notifications"] = retransmit.StopControlConnectionNotifications
    leafs["hello-packets"] = retransmit.HelloPackets
    leafs["outgoing-call-requests"] = retransmit.OutgoingCallRequests
    leafs["outgoing-call-replies"] = retransmit.OutgoingCallReplies
    leafs["outgoing-call-connected-packets"] = retransmit.OutgoingCallConnectedPackets
    leafs["incoming-call-requests"] = retransmit.IncomingCallRequests
    leafs["incoming-call-replies"] = retransmit.IncomingCallReplies
    leafs["incoming-call-connected-packets"] = retransmit.IncomingCallConnectedPackets
    leafs["call-disconnect-notify-packets"] = retransmit.CallDisconnectNotifyPackets
    leafs["wan-error-notify-packets"] = retransmit.WanErrorNotifyPackets
    leafs["set-link-info-packets"] = retransmit.SetLinkInfoPackets
    leafs["service-relay-requests"] = retransmit.ServiceRelayRequests
    leafs["service-relay-replies"] = retransmit.ServiceRelayReplies
    leafs["acknowledgement-packets"] = retransmit.AcknowledgementPackets
    return leafs
}

func (retransmit *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Retransmit) GetBundleName() string { return "cisco_ios_xr" }

func (retransmit *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Retransmit) GetYangName() string { return "retransmit" }

func (retransmit *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Retransmit) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (retransmit *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Retransmit) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (retransmit *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Retransmit) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (retransmit *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Retransmit) SetParent(parent types.Entity) { retransmit.parent = parent }

func (retransmit *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Retransmit) GetParent() types.Entity { return retransmit.parent }

func (retransmit *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Retransmit) GetParentYangName() string { return "global" }

// L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Received
// Received data
type L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Received struct {
    parent types.Entity
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

func (received *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Received) GetFilter() yfilter.YFilter { return received.YFilter }

func (received *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Received) SetFilter(yf yfilter.YFilter) { received.YFilter = yf }

func (received *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Received) GetGoName(yname string) string {
    if yname == "unknown-packets" { return "UnknownPackets" }
    if yname == "zero-length-body-packets" { return "ZeroLengthBodyPackets" }
    if yname == "start-control-connection-requests" { return "StartControlConnectionRequests" }
    if yname == "start-control-connection-replies" { return "StartControlConnectionReplies" }
    if yname == "start-control-connection-notifications" { return "StartControlConnectionNotifications" }
    if yname == "stop-control-connection-notifications" { return "StopControlConnectionNotifications" }
    if yname == "hello-packets" { return "HelloPackets" }
    if yname == "outgoing-call-requests" { return "OutgoingCallRequests" }
    if yname == "outgoing-call-replies" { return "OutgoingCallReplies" }
    if yname == "outgoing-call-connected-packets" { return "OutgoingCallConnectedPackets" }
    if yname == "incoming-call-requests" { return "IncomingCallRequests" }
    if yname == "incoming-call-replies" { return "IncomingCallReplies" }
    if yname == "incoming-call-connected-packets" { return "IncomingCallConnectedPackets" }
    if yname == "call-disconnect-notify-packets" { return "CallDisconnectNotifyPackets" }
    if yname == "wan-error-notify-packets" { return "WanErrorNotifyPackets" }
    if yname == "set-link-info-packets" { return "SetLinkInfoPackets" }
    if yname == "service-relay-requests" { return "ServiceRelayRequests" }
    if yname == "service-relay-replies" { return "ServiceRelayReplies" }
    if yname == "acknowledgement-packets" { return "AcknowledgementPackets" }
    return ""
}

func (received *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Received) GetSegmentPath() string {
    return "received"
}

func (received *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Received) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (received *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Received) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (received *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Received) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["unknown-packets"] = received.UnknownPackets
    leafs["zero-length-body-packets"] = received.ZeroLengthBodyPackets
    leafs["start-control-connection-requests"] = received.StartControlConnectionRequests
    leafs["start-control-connection-replies"] = received.StartControlConnectionReplies
    leafs["start-control-connection-notifications"] = received.StartControlConnectionNotifications
    leafs["stop-control-connection-notifications"] = received.StopControlConnectionNotifications
    leafs["hello-packets"] = received.HelloPackets
    leafs["outgoing-call-requests"] = received.OutgoingCallRequests
    leafs["outgoing-call-replies"] = received.OutgoingCallReplies
    leafs["outgoing-call-connected-packets"] = received.OutgoingCallConnectedPackets
    leafs["incoming-call-requests"] = received.IncomingCallRequests
    leafs["incoming-call-replies"] = received.IncomingCallReplies
    leafs["incoming-call-connected-packets"] = received.IncomingCallConnectedPackets
    leafs["call-disconnect-notify-packets"] = received.CallDisconnectNotifyPackets
    leafs["wan-error-notify-packets"] = received.WanErrorNotifyPackets
    leafs["set-link-info-packets"] = received.SetLinkInfoPackets
    leafs["service-relay-requests"] = received.ServiceRelayRequests
    leafs["service-relay-replies"] = received.ServiceRelayReplies
    leafs["acknowledgement-packets"] = received.AcknowledgementPackets
    return leafs
}

func (received *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Received) GetBundleName() string { return "cisco_ios_xr" }

func (received *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Received) GetYangName() string { return "received" }

func (received *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Received) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (received *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Received) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (received *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Received) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (received *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Received) SetParent(parent types.Entity) { received.parent = parent }

func (received *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Received) GetParent() types.Entity { return received.parent }

func (received *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Received) GetParentYangName() string { return "global" }

// L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Drop
// Drop data
type L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Drop struct {
    parent types.Entity
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

func (drop *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Drop) GetFilter() yfilter.YFilter { return drop.YFilter }

func (drop *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Drop) SetFilter(yf yfilter.YFilter) { drop.YFilter = yf }

func (drop *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Drop) GetGoName(yname string) string {
    if yname == "unknown-packets" { return "UnknownPackets" }
    if yname == "zero-length-body-packets" { return "ZeroLengthBodyPackets" }
    if yname == "start-control-connection-requests" { return "StartControlConnectionRequests" }
    if yname == "start-control-connection-replies" { return "StartControlConnectionReplies" }
    if yname == "start-control-connection-notifications" { return "StartControlConnectionNotifications" }
    if yname == "stop-control-connection-notifications" { return "StopControlConnectionNotifications" }
    if yname == "hello-packets" { return "HelloPackets" }
    if yname == "outgoing-call-requests" { return "OutgoingCallRequests" }
    if yname == "outgoing-call-replies" { return "OutgoingCallReplies" }
    if yname == "outgoing-call-connected-packets" { return "OutgoingCallConnectedPackets" }
    if yname == "incoming-call-requests" { return "IncomingCallRequests" }
    if yname == "incoming-call-replies" { return "IncomingCallReplies" }
    if yname == "incoming-call-connected-packets" { return "IncomingCallConnectedPackets" }
    if yname == "call-disconnect-notify-packets" { return "CallDisconnectNotifyPackets" }
    if yname == "wan-error-notify-packets" { return "WanErrorNotifyPackets" }
    if yname == "set-link-info-packets" { return "SetLinkInfoPackets" }
    if yname == "service-relay-requests" { return "ServiceRelayRequests" }
    if yname == "service-relay-replies" { return "ServiceRelayReplies" }
    if yname == "acknowledgement-packets" { return "AcknowledgementPackets" }
    return ""
}

func (drop *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Drop) GetSegmentPath() string {
    return "drop"
}

func (drop *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Drop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (drop *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Drop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (drop *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Drop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["unknown-packets"] = drop.UnknownPackets
    leafs["zero-length-body-packets"] = drop.ZeroLengthBodyPackets
    leafs["start-control-connection-requests"] = drop.StartControlConnectionRequests
    leafs["start-control-connection-replies"] = drop.StartControlConnectionReplies
    leafs["start-control-connection-notifications"] = drop.StartControlConnectionNotifications
    leafs["stop-control-connection-notifications"] = drop.StopControlConnectionNotifications
    leafs["hello-packets"] = drop.HelloPackets
    leafs["outgoing-call-requests"] = drop.OutgoingCallRequests
    leafs["outgoing-call-replies"] = drop.OutgoingCallReplies
    leafs["outgoing-call-connected-packets"] = drop.OutgoingCallConnectedPackets
    leafs["incoming-call-requests"] = drop.IncomingCallRequests
    leafs["incoming-call-replies"] = drop.IncomingCallReplies
    leafs["incoming-call-connected-packets"] = drop.IncomingCallConnectedPackets
    leafs["call-disconnect-notify-packets"] = drop.CallDisconnectNotifyPackets
    leafs["wan-error-notify-packets"] = drop.WanErrorNotifyPackets
    leafs["set-link-info-packets"] = drop.SetLinkInfoPackets
    leafs["service-relay-requests"] = drop.ServiceRelayRequests
    leafs["service-relay-replies"] = drop.ServiceRelayReplies
    leafs["acknowledgement-packets"] = drop.AcknowledgementPackets
    return leafs
}

func (drop *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Drop) GetBundleName() string { return "cisco_ios_xr" }

func (drop *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Drop) GetYangName() string { return "drop" }

func (drop *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Drop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (drop *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Drop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (drop *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Drop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (drop *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Drop) SetParent(parent types.Entity) { drop.parent = parent }

func (drop *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Drop) GetParent() types.Entity { return drop.parent }

func (drop *L2Tpv2_Counters_Control_Tunnels_Tunnel_Global_Drop) GetParentYangName() string { return "global" }

// L2Tpv2_Statistics
// L2TP v2 statistics information
type L2Tpv2_Statistics struct {
    parent types.Entity
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

func (statistics *L2Tpv2_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *L2Tpv2_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *L2Tpv2_Statistics) GetGoName(yname string) string {
    if yname == "tunnels" { return "Tunnels" }
    if yname == "sessions" { return "Sessions" }
    if yname == "sent-packets" { return "SentPackets" }
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "average-packet-processing-time" { return "AveragePacketProcessingTime" }
    if yname == "received-out-of-order-packets" { return "ReceivedOutOfOrderPackets" }
    if yname == "reorder-packets" { return "ReorderPackets" }
    if yname == "reorder-deviation-packets" { return "ReorderDeviationPackets" }
    if yname == "incoming-dropped-packets" { return "IncomingDroppedPackets" }
    if yname == "buffered-packets" { return "BufferedPackets" }
    if yname == "netio-packets" { return "NetioPackets" }
    return ""
}

func (statistics *L2Tpv2_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *L2Tpv2_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (statistics *L2Tpv2_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (statistics *L2Tpv2_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tunnels"] = statistics.Tunnels
    leafs["sessions"] = statistics.Sessions
    leafs["sent-packets"] = statistics.SentPackets
    leafs["received-packets"] = statistics.ReceivedPackets
    leafs["average-packet-processing-time"] = statistics.AveragePacketProcessingTime
    leafs["received-out-of-order-packets"] = statistics.ReceivedOutOfOrderPackets
    leafs["reorder-packets"] = statistics.ReorderPackets
    leafs["reorder-deviation-packets"] = statistics.ReorderDeviationPackets
    leafs["incoming-dropped-packets"] = statistics.IncomingDroppedPackets
    leafs["buffered-packets"] = statistics.BufferedPackets
    leafs["netio-packets"] = statistics.NetioPackets
    return leafs
}

func (statistics *L2Tpv2_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *L2Tpv2_Statistics) GetYangName() string { return "statistics" }

func (statistics *L2Tpv2_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *L2Tpv2_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *L2Tpv2_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *L2Tpv2_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *L2Tpv2_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *L2Tpv2_Statistics) GetParentYangName() string { return "l2tpv2" }

// L2Tpv2_Tunnel
// L2TPv2 tunnel 
type L2Tpv2_Tunnel struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Tunnel accounting counters.
    Accounting L2Tpv2_Tunnel_Accounting
}

func (tunnel *L2Tpv2_Tunnel) GetFilter() yfilter.YFilter { return tunnel.YFilter }

func (tunnel *L2Tpv2_Tunnel) SetFilter(yf yfilter.YFilter) { tunnel.YFilter = yf }

func (tunnel *L2Tpv2_Tunnel) GetGoName(yname string) string {
    if yname == "accounting" { return "Accounting" }
    return ""
}

func (tunnel *L2Tpv2_Tunnel) GetSegmentPath() string {
    return "tunnel"
}

func (tunnel *L2Tpv2_Tunnel) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "accounting" {
        return &tunnel.Accounting
    }
    return nil
}

func (tunnel *L2Tpv2_Tunnel) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["accounting"] = &tunnel.Accounting
    return children
}

func (tunnel *L2Tpv2_Tunnel) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (tunnel *L2Tpv2_Tunnel) GetBundleName() string { return "cisco_ios_xr" }

func (tunnel *L2Tpv2_Tunnel) GetYangName() string { return "tunnel" }

func (tunnel *L2Tpv2_Tunnel) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tunnel *L2Tpv2_Tunnel) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tunnel *L2Tpv2_Tunnel) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tunnel *L2Tpv2_Tunnel) SetParent(parent types.Entity) { tunnel.parent = parent }

func (tunnel *L2Tpv2_Tunnel) GetParent() types.Entity { return tunnel.parent }

func (tunnel *L2Tpv2_Tunnel) GetParentYangName() string { return "l2tpv2" }

// L2Tpv2_Tunnel_Accounting
// Tunnel accounting counters
type L2Tpv2_Tunnel_Accounting struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Tunnel accounting statistics.
    Statistics L2Tpv2_Tunnel_Accounting_Statistics
}

func (accounting *L2Tpv2_Tunnel_Accounting) GetFilter() yfilter.YFilter { return accounting.YFilter }

func (accounting *L2Tpv2_Tunnel_Accounting) SetFilter(yf yfilter.YFilter) { accounting.YFilter = yf }

func (accounting *L2Tpv2_Tunnel_Accounting) GetGoName(yname string) string {
    if yname == "statistics" { return "Statistics" }
    return ""
}

func (accounting *L2Tpv2_Tunnel_Accounting) GetSegmentPath() string {
    return "accounting"
}

func (accounting *L2Tpv2_Tunnel_Accounting) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "statistics" {
        return &accounting.Statistics
    }
    return nil
}

func (accounting *L2Tpv2_Tunnel_Accounting) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["statistics"] = &accounting.Statistics
    return children
}

func (accounting *L2Tpv2_Tunnel_Accounting) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (accounting *L2Tpv2_Tunnel_Accounting) GetBundleName() string { return "cisco_ios_xr" }

func (accounting *L2Tpv2_Tunnel_Accounting) GetYangName() string { return "accounting" }

func (accounting *L2Tpv2_Tunnel_Accounting) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accounting *L2Tpv2_Tunnel_Accounting) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accounting *L2Tpv2_Tunnel_Accounting) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accounting *L2Tpv2_Tunnel_Accounting) SetParent(parent types.Entity) { accounting.parent = parent }

func (accounting *L2Tpv2_Tunnel_Accounting) GetParent() types.Entity { return accounting.parent }

func (accounting *L2Tpv2_Tunnel_Accounting) GetParentYangName() string { return "tunnel" }

// L2Tpv2_Tunnel_Accounting_Statistics
// Tunnel accounting statistics
type L2Tpv2_Tunnel_Accounting_Statistics struct {
    parent types.Entity
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

func (statistics *L2Tpv2_Tunnel_Accounting_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *L2Tpv2_Tunnel_Accounting_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *L2Tpv2_Tunnel_Accounting_Statistics) GetGoName(yname string) string {
    if yname == "records-sent-successfully" { return "RecordsSentSuccessfully" }
    if yname == "start" { return "Start" }
    if yname == "stop" { return "Stop" }
    if yname == "reject" { return "Reject" }
    if yname == "transport-failures" { return "TransportFailures" }
    if yname == "positive-acknowledgement" { return "PositiveAcknowledgement" }
    if yname == "negative-acknowledgement" { return "NegativeAcknowledgement" }
    if yname == "records-checkpointed" { return "RecordsCheckpointed" }
    if yname == "records-failed-to-checkpoint" { return "RecordsFailedToCheckpoint" }
    if yname == "records-sent-from-queue" { return "RecordsSentFromQueue" }
    if yname == "memory-failures" { return "MemoryFailures" }
    if yname == "current-size" { return "CurrentSize" }
    if yname == "records-recovered-from-checkpoint" { return "RecordsRecoveredFromCheckpoint" }
    if yname == "records-fail-to-recover" { return "RecordsFailToRecover" }
    if yname == "queue-statistics-size" { return "QueueStatisticsSize" }
    return ""
}

func (statistics *L2Tpv2_Tunnel_Accounting_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *L2Tpv2_Tunnel_Accounting_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (statistics *L2Tpv2_Tunnel_Accounting_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (statistics *L2Tpv2_Tunnel_Accounting_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["records-sent-successfully"] = statistics.RecordsSentSuccessfully
    leafs["start"] = statistics.Start
    leafs["stop"] = statistics.Stop
    leafs["reject"] = statistics.Reject
    leafs["transport-failures"] = statistics.TransportFailures
    leafs["positive-acknowledgement"] = statistics.PositiveAcknowledgement
    leafs["negative-acknowledgement"] = statistics.NegativeAcknowledgement
    leafs["records-checkpointed"] = statistics.RecordsCheckpointed
    leafs["records-failed-to-checkpoint"] = statistics.RecordsFailedToCheckpoint
    leafs["records-sent-from-queue"] = statistics.RecordsSentFromQueue
    leafs["memory-failures"] = statistics.MemoryFailures
    leafs["current-size"] = statistics.CurrentSize
    leafs["records-recovered-from-checkpoint"] = statistics.RecordsRecoveredFromCheckpoint
    leafs["records-fail-to-recover"] = statistics.RecordsFailToRecover
    leafs["queue-statistics-size"] = statistics.QueueStatisticsSize
    return leafs
}

func (statistics *L2Tpv2_Tunnel_Accounting_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *L2Tpv2_Tunnel_Accounting_Statistics) GetYangName() string { return "statistics" }

func (statistics *L2Tpv2_Tunnel_Accounting_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *L2Tpv2_Tunnel_Accounting_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *L2Tpv2_Tunnel_Accounting_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *L2Tpv2_Tunnel_Accounting_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *L2Tpv2_Tunnel_Accounting_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *L2Tpv2_Tunnel_Accounting_Statistics) GetParentYangName() string { return "accounting" }

// L2Tpv2_TunnelConfigurations
// List of tunnel IDs
type L2Tpv2_TunnelConfigurations struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // L2TP tunnel information. The type is slice of
    // L2Tpv2_TunnelConfigurations_TunnelConfiguration.
    TunnelConfiguration []L2Tpv2_TunnelConfigurations_TunnelConfiguration
}

func (tunnelConfigurations *L2Tpv2_TunnelConfigurations) GetFilter() yfilter.YFilter { return tunnelConfigurations.YFilter }

func (tunnelConfigurations *L2Tpv2_TunnelConfigurations) SetFilter(yf yfilter.YFilter) { tunnelConfigurations.YFilter = yf }

func (tunnelConfigurations *L2Tpv2_TunnelConfigurations) GetGoName(yname string) string {
    if yname == "tunnel-configuration" { return "TunnelConfiguration" }
    return ""
}

func (tunnelConfigurations *L2Tpv2_TunnelConfigurations) GetSegmentPath() string {
    return "tunnel-configurations"
}

func (tunnelConfigurations *L2Tpv2_TunnelConfigurations) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tunnel-configuration" {
        for _, c := range tunnelConfigurations.TunnelConfiguration {
            if tunnelConfigurations.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := L2Tpv2_TunnelConfigurations_TunnelConfiguration{}
        tunnelConfigurations.TunnelConfiguration = append(tunnelConfigurations.TunnelConfiguration, child)
        return &tunnelConfigurations.TunnelConfiguration[len(tunnelConfigurations.TunnelConfiguration)-1]
    }
    return nil
}

func (tunnelConfigurations *L2Tpv2_TunnelConfigurations) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range tunnelConfigurations.TunnelConfiguration {
        children[tunnelConfigurations.TunnelConfiguration[i].GetSegmentPath()] = &tunnelConfigurations.TunnelConfiguration[i]
    }
    return children
}

func (tunnelConfigurations *L2Tpv2_TunnelConfigurations) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (tunnelConfigurations *L2Tpv2_TunnelConfigurations) GetBundleName() string { return "cisco_ios_xr" }

func (tunnelConfigurations *L2Tpv2_TunnelConfigurations) GetYangName() string { return "tunnel-configurations" }

func (tunnelConfigurations *L2Tpv2_TunnelConfigurations) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tunnelConfigurations *L2Tpv2_TunnelConfigurations) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tunnelConfigurations *L2Tpv2_TunnelConfigurations) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tunnelConfigurations *L2Tpv2_TunnelConfigurations) SetParent(parent types.Entity) { tunnelConfigurations.parent = parent }

func (tunnelConfigurations *L2Tpv2_TunnelConfigurations) GetParent() types.Entity { return tunnelConfigurations.parent }

func (tunnelConfigurations *L2Tpv2_TunnelConfigurations) GetParentYangName() string { return "l2tpv2" }

// L2Tpv2_TunnelConfigurations_TunnelConfiguration
// L2TP tunnel information
type L2Tpv2_TunnelConfigurations_TunnelConfiguration struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Local tunnel ID. The type is interface{} with
    // range: -2147483648..2147483647.
    LocalTunnelId interface{}

    // Remote tunnel ID. The type is interface{} with range: 0..4294967295.
    RemoteTunnelId interface{}

    // L2Tp class data.
    L2TpClass L2Tpv2_TunnelConfigurations_TunnelConfiguration_L2TpClass
}

func (tunnelConfiguration *L2Tpv2_TunnelConfigurations_TunnelConfiguration) GetFilter() yfilter.YFilter { return tunnelConfiguration.YFilter }

func (tunnelConfiguration *L2Tpv2_TunnelConfigurations_TunnelConfiguration) SetFilter(yf yfilter.YFilter) { tunnelConfiguration.YFilter = yf }

func (tunnelConfiguration *L2Tpv2_TunnelConfigurations_TunnelConfiguration) GetGoName(yname string) string {
    if yname == "local-tunnel-id" { return "LocalTunnelId" }
    if yname == "remote-tunnel-id" { return "RemoteTunnelId" }
    if yname == "l2tp-class" { return "L2TpClass" }
    return ""
}

func (tunnelConfiguration *L2Tpv2_TunnelConfigurations_TunnelConfiguration) GetSegmentPath() string {
    return "tunnel-configuration" + "[local-tunnel-id='" + fmt.Sprintf("%v", tunnelConfiguration.LocalTunnelId) + "']"
}

func (tunnelConfiguration *L2Tpv2_TunnelConfigurations_TunnelConfiguration) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "l2tp-class" {
        return &tunnelConfiguration.L2TpClass
    }
    return nil
}

func (tunnelConfiguration *L2Tpv2_TunnelConfigurations_TunnelConfiguration) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["l2tp-class"] = &tunnelConfiguration.L2TpClass
    return children
}

func (tunnelConfiguration *L2Tpv2_TunnelConfigurations_TunnelConfiguration) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["local-tunnel-id"] = tunnelConfiguration.LocalTunnelId
    leafs["remote-tunnel-id"] = tunnelConfiguration.RemoteTunnelId
    return leafs
}

func (tunnelConfiguration *L2Tpv2_TunnelConfigurations_TunnelConfiguration) GetBundleName() string { return "cisco_ios_xr" }

func (tunnelConfiguration *L2Tpv2_TunnelConfigurations_TunnelConfiguration) GetYangName() string { return "tunnel-configuration" }

func (tunnelConfiguration *L2Tpv2_TunnelConfigurations_TunnelConfiguration) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tunnelConfiguration *L2Tpv2_TunnelConfigurations_TunnelConfiguration) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tunnelConfiguration *L2Tpv2_TunnelConfigurations_TunnelConfiguration) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tunnelConfiguration *L2Tpv2_TunnelConfigurations_TunnelConfiguration) SetParent(parent types.Entity) { tunnelConfiguration.parent = parent }

func (tunnelConfiguration *L2Tpv2_TunnelConfigurations_TunnelConfiguration) GetParent() types.Entity { return tunnelConfiguration.parent }

func (tunnelConfiguration *L2Tpv2_TunnelConfigurations_TunnelConfiguration) GetParentYangName() string { return "tunnel-configurations" }

// L2Tpv2_TunnelConfigurations_TunnelConfiguration_L2TpClass
// L2Tp class data
type L2Tpv2_TunnelConfigurations_TunnelConfiguration_L2TpClass struct {
    parent types.Entity
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

func (l2TpClass *L2Tpv2_TunnelConfigurations_TunnelConfiguration_L2TpClass) GetFilter() yfilter.YFilter { return l2TpClass.YFilter }

func (l2TpClass *L2Tpv2_TunnelConfigurations_TunnelConfiguration_L2TpClass) SetFilter(yf yfilter.YFilter) { l2TpClass.YFilter = yf }

func (l2TpClass *L2Tpv2_TunnelConfigurations_TunnelConfiguration_L2TpClass) GetGoName(yname string) string {
    if yname == "ip-tos" { return "IpTos" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "receive-window-size" { return "ReceiveWindowSize" }
    if yname == "class-name-xr" { return "ClassNameXr" }
    if yname == "digest-hash" { return "DigestHash" }
    if yname == "password" { return "Password" }
    if yname == "encoded-password" { return "EncodedPassword" }
    if yname == "host-name" { return "HostName" }
    if yname == "accounting-method-list" { return "AccountingMethodList" }
    if yname == "hello-timeout" { return "HelloTimeout" }
    if yname == "setup-timeout" { return "SetupTimeout" }
    if yname == "retransmit-minimum-timeout" { return "RetransmitMinimumTimeout" }
    if yname == "retransmit-maximum-timeout" { return "RetransmitMaximumTimeout" }
    if yname == "initial-retransmit-minimum-timeout" { return "InitialRetransmitMinimumTimeout" }
    if yname == "initial-retransmit-maximum-timeout" { return "InitialRetransmitMaximumTimeout" }
    if yname == "timeout-no-user" { return "TimeoutNoUser" }
    if yname == "retransmit-retries" { return "RetransmitRetries" }
    if yname == "initial-retransmit-retries" { return "InitialRetransmitRetries" }
    if yname == "is-authentication-enabled" { return "IsAuthenticationEnabled" }
    if yname == "is-hidden" { return "IsHidden" }
    if yname == "is-digest-enabled" { return "IsDigestEnabled" }
    if yname == "is-digest-check-enabled" { return "IsDigestCheckEnabled" }
    if yname == "is-congestion-control-enabled" { return "IsCongestionControlEnabled" }
    if yname == "is-peer-address-checked" { return "IsPeerAddressChecked" }
    return ""
}

func (l2TpClass *L2Tpv2_TunnelConfigurations_TunnelConfiguration_L2TpClass) GetSegmentPath() string {
    return "l2tp-class"
}

func (l2TpClass *L2Tpv2_TunnelConfigurations_TunnelConfiguration_L2TpClass) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (l2TpClass *L2Tpv2_TunnelConfigurations_TunnelConfiguration_L2TpClass) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (l2TpClass *L2Tpv2_TunnelConfigurations_TunnelConfiguration_L2TpClass) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip-tos"] = l2TpClass.IpTos
    leafs["vrf-name"] = l2TpClass.VrfName
    leafs["receive-window-size"] = l2TpClass.ReceiveWindowSize
    leafs["class-name-xr"] = l2TpClass.ClassNameXr
    leafs["digest-hash"] = l2TpClass.DigestHash
    leafs["password"] = l2TpClass.Password
    leafs["encoded-password"] = l2TpClass.EncodedPassword
    leafs["host-name"] = l2TpClass.HostName
    leafs["accounting-method-list"] = l2TpClass.AccountingMethodList
    leafs["hello-timeout"] = l2TpClass.HelloTimeout
    leafs["setup-timeout"] = l2TpClass.SetupTimeout
    leafs["retransmit-minimum-timeout"] = l2TpClass.RetransmitMinimumTimeout
    leafs["retransmit-maximum-timeout"] = l2TpClass.RetransmitMaximumTimeout
    leafs["initial-retransmit-minimum-timeout"] = l2TpClass.InitialRetransmitMinimumTimeout
    leafs["initial-retransmit-maximum-timeout"] = l2TpClass.InitialRetransmitMaximumTimeout
    leafs["timeout-no-user"] = l2TpClass.TimeoutNoUser
    leafs["retransmit-retries"] = l2TpClass.RetransmitRetries
    leafs["initial-retransmit-retries"] = l2TpClass.InitialRetransmitRetries
    leafs["is-authentication-enabled"] = l2TpClass.IsAuthenticationEnabled
    leafs["is-hidden"] = l2TpClass.IsHidden
    leafs["is-digest-enabled"] = l2TpClass.IsDigestEnabled
    leafs["is-digest-check-enabled"] = l2TpClass.IsDigestCheckEnabled
    leafs["is-congestion-control-enabled"] = l2TpClass.IsCongestionControlEnabled
    leafs["is-peer-address-checked"] = l2TpClass.IsPeerAddressChecked
    return leafs
}

func (l2TpClass *L2Tpv2_TunnelConfigurations_TunnelConfiguration_L2TpClass) GetBundleName() string { return "cisco_ios_xr" }

func (l2TpClass *L2Tpv2_TunnelConfigurations_TunnelConfiguration_L2TpClass) GetYangName() string { return "l2tp-class" }

func (l2TpClass *L2Tpv2_TunnelConfigurations_TunnelConfiguration_L2TpClass) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (l2TpClass *L2Tpv2_TunnelConfigurations_TunnelConfiguration_L2TpClass) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (l2TpClass *L2Tpv2_TunnelConfigurations_TunnelConfiguration_L2TpClass) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (l2TpClass *L2Tpv2_TunnelConfigurations_TunnelConfiguration_L2TpClass) SetParent(parent types.Entity) { l2TpClass.parent = parent }

func (l2TpClass *L2Tpv2_TunnelConfigurations_TunnelConfiguration_L2TpClass) GetParent() types.Entity { return l2TpClass.parent }

func (l2TpClass *L2Tpv2_TunnelConfigurations_TunnelConfiguration_L2TpClass) GetParentYangName() string { return "tunnel-configuration" }

// L2Tpv2_CounterHistFail
// Failure events leading to disconnection
type L2Tpv2_CounterHistFail struct {
    parent types.Entity
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

    // timeout events by packet. The type is slice of interface{} with range:
    // 0..4294967295.
    PktTimeout []interface{}
}

func (counterHistFail *L2Tpv2_CounterHistFail) GetFilter() yfilter.YFilter { return counterHistFail.YFilter }

func (counterHistFail *L2Tpv2_CounterHistFail) SetFilter(yf yfilter.YFilter) { counterHistFail.YFilter = yf }

func (counterHistFail *L2Tpv2_CounterHistFail) GetGoName(yname string) string {
    if yname == "sess-down-tmout" { return "SessDownTmout" }
    if yname == "tx-counters" { return "TxCounters" }
    if yname == "rx-counters" { return "RxCounters" }
    if yname == "pkt-timeout" { return "PktTimeout" }
    return ""
}

func (counterHistFail *L2Tpv2_CounterHistFail) GetSegmentPath() string {
    return "counter-hist-fail"
}

func (counterHistFail *L2Tpv2_CounterHistFail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (counterHistFail *L2Tpv2_CounterHistFail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (counterHistFail *L2Tpv2_CounterHistFail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sess-down-tmout"] = counterHistFail.SessDownTmout
    leafs["tx-counters"] = counterHistFail.TxCounters
    leafs["rx-counters"] = counterHistFail.RxCounters
    leafs["pkt-timeout"] = counterHistFail.PktTimeout
    return leafs
}

func (counterHistFail *L2Tpv2_CounterHistFail) GetBundleName() string { return "cisco_ios_xr" }

func (counterHistFail *L2Tpv2_CounterHistFail) GetYangName() string { return "counter-hist-fail" }

func (counterHistFail *L2Tpv2_CounterHistFail) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (counterHistFail *L2Tpv2_CounterHistFail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (counterHistFail *L2Tpv2_CounterHistFail) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (counterHistFail *L2Tpv2_CounterHistFail) SetParent(parent types.Entity) { counterHistFail.parent = parent }

func (counterHistFail *L2Tpv2_CounterHistFail) GetParent() types.Entity { return counterHistFail.parent }

func (counterHistFail *L2Tpv2_CounterHistFail) GetParentYangName() string { return "l2tpv2" }

// L2Tpv2_Classes
// List of L2TP class names
type L2Tpv2_Classes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // L2TP class name. The type is slice of L2Tpv2_Classes_Class.
    Class []L2Tpv2_Classes_Class
}

func (classes *L2Tpv2_Classes) GetFilter() yfilter.YFilter { return classes.YFilter }

func (classes *L2Tpv2_Classes) SetFilter(yf yfilter.YFilter) { classes.YFilter = yf }

func (classes *L2Tpv2_Classes) GetGoName(yname string) string {
    if yname == "class" { return "Class" }
    return ""
}

func (classes *L2Tpv2_Classes) GetSegmentPath() string {
    return "classes"
}

func (classes *L2Tpv2_Classes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "class" {
        for _, c := range classes.Class {
            if classes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := L2Tpv2_Classes_Class{}
        classes.Class = append(classes.Class, child)
        return &classes.Class[len(classes.Class)-1]
    }
    return nil
}

func (classes *L2Tpv2_Classes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range classes.Class {
        children[classes.Class[i].GetSegmentPath()] = &classes.Class[i]
    }
    return children
}

func (classes *L2Tpv2_Classes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (classes *L2Tpv2_Classes) GetBundleName() string { return "cisco_ios_xr" }

func (classes *L2Tpv2_Classes) GetYangName() string { return "classes" }

func (classes *L2Tpv2_Classes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (classes *L2Tpv2_Classes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (classes *L2Tpv2_Classes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (classes *L2Tpv2_Classes) SetParent(parent types.Entity) { classes.parent = parent }

func (classes *L2Tpv2_Classes) GetParent() types.Entity { return classes.parent }

func (classes *L2Tpv2_Classes) GetParentYangName() string { return "l2tpv2" }

// L2Tpv2_Classes_Class
// L2TP class name
type L2Tpv2_Classes_Class struct {
    parent types.Entity
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

func (class *L2Tpv2_Classes_Class) GetFilter() yfilter.YFilter { return class.YFilter }

func (class *L2Tpv2_Classes_Class) SetFilter(yf yfilter.YFilter) { class.YFilter = yf }

func (class *L2Tpv2_Classes_Class) GetGoName(yname string) string {
    if yname == "class-name" { return "ClassName" }
    if yname == "ip-tos" { return "IpTos" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "receive-window-size" { return "ReceiveWindowSize" }
    if yname == "class-name-xr" { return "ClassNameXr" }
    if yname == "digest-hash" { return "DigestHash" }
    if yname == "password" { return "Password" }
    if yname == "encoded-password" { return "EncodedPassword" }
    if yname == "host-name" { return "HostName" }
    if yname == "accounting-method-list" { return "AccountingMethodList" }
    if yname == "hello-timeout" { return "HelloTimeout" }
    if yname == "setup-timeout" { return "SetupTimeout" }
    if yname == "retransmit-minimum-timeout" { return "RetransmitMinimumTimeout" }
    if yname == "retransmit-maximum-timeout" { return "RetransmitMaximumTimeout" }
    if yname == "initial-retransmit-minimum-timeout" { return "InitialRetransmitMinimumTimeout" }
    if yname == "initial-retransmit-maximum-timeout" { return "InitialRetransmitMaximumTimeout" }
    if yname == "timeout-no-user" { return "TimeoutNoUser" }
    if yname == "retransmit-retries" { return "RetransmitRetries" }
    if yname == "initial-retransmit-retries" { return "InitialRetransmitRetries" }
    if yname == "is-authentication-enabled" { return "IsAuthenticationEnabled" }
    if yname == "is-hidden" { return "IsHidden" }
    if yname == "is-digest-enabled" { return "IsDigestEnabled" }
    if yname == "is-digest-check-enabled" { return "IsDigestCheckEnabled" }
    if yname == "is-congestion-control-enabled" { return "IsCongestionControlEnabled" }
    if yname == "is-peer-address-checked" { return "IsPeerAddressChecked" }
    return ""
}

func (class *L2Tpv2_Classes_Class) GetSegmentPath() string {
    return "class" + "[class-name='" + fmt.Sprintf("%v", class.ClassName) + "']"
}

func (class *L2Tpv2_Classes_Class) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (class *L2Tpv2_Classes_Class) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (class *L2Tpv2_Classes_Class) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["class-name"] = class.ClassName
    leafs["ip-tos"] = class.IpTos
    leafs["vrf-name"] = class.VrfName
    leafs["receive-window-size"] = class.ReceiveWindowSize
    leafs["class-name-xr"] = class.ClassNameXr
    leafs["digest-hash"] = class.DigestHash
    leafs["password"] = class.Password
    leafs["encoded-password"] = class.EncodedPassword
    leafs["host-name"] = class.HostName
    leafs["accounting-method-list"] = class.AccountingMethodList
    leafs["hello-timeout"] = class.HelloTimeout
    leafs["setup-timeout"] = class.SetupTimeout
    leafs["retransmit-minimum-timeout"] = class.RetransmitMinimumTimeout
    leafs["retransmit-maximum-timeout"] = class.RetransmitMaximumTimeout
    leafs["initial-retransmit-minimum-timeout"] = class.InitialRetransmitMinimumTimeout
    leafs["initial-retransmit-maximum-timeout"] = class.InitialRetransmitMaximumTimeout
    leafs["timeout-no-user"] = class.TimeoutNoUser
    leafs["retransmit-retries"] = class.RetransmitRetries
    leafs["initial-retransmit-retries"] = class.InitialRetransmitRetries
    leafs["is-authentication-enabled"] = class.IsAuthenticationEnabled
    leafs["is-hidden"] = class.IsHidden
    leafs["is-digest-enabled"] = class.IsDigestEnabled
    leafs["is-digest-check-enabled"] = class.IsDigestCheckEnabled
    leafs["is-congestion-control-enabled"] = class.IsCongestionControlEnabled
    leafs["is-peer-address-checked"] = class.IsPeerAddressChecked
    return leafs
}

func (class *L2Tpv2_Classes_Class) GetBundleName() string { return "cisco_ios_xr" }

func (class *L2Tpv2_Classes_Class) GetYangName() string { return "class" }

func (class *L2Tpv2_Classes_Class) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (class *L2Tpv2_Classes_Class) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (class *L2Tpv2_Classes_Class) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (class *L2Tpv2_Classes_Class) SetParent(parent types.Entity) { class.parent = parent }

func (class *L2Tpv2_Classes_Class) GetParent() types.Entity { return class.parent }

func (class *L2Tpv2_Classes_Class) GetParentYangName() string { return "classes" }

// L2Tpv2_Tunnels
// List of tunnel IDs
type L2Tpv2_Tunnels struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // L2TP tunnel  information. The type is slice of L2Tpv2_Tunnels_Tunnel.
    Tunnel []L2Tpv2_Tunnels_Tunnel
}

func (tunnels *L2Tpv2_Tunnels) GetFilter() yfilter.YFilter { return tunnels.YFilter }

func (tunnels *L2Tpv2_Tunnels) SetFilter(yf yfilter.YFilter) { tunnels.YFilter = yf }

func (tunnels *L2Tpv2_Tunnels) GetGoName(yname string) string {
    if yname == "tunnel" { return "Tunnel" }
    return ""
}

func (tunnels *L2Tpv2_Tunnels) GetSegmentPath() string {
    return "tunnels"
}

func (tunnels *L2Tpv2_Tunnels) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tunnel" {
        for _, c := range tunnels.Tunnel {
            if tunnels.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := L2Tpv2_Tunnels_Tunnel{}
        tunnels.Tunnel = append(tunnels.Tunnel, child)
        return &tunnels.Tunnel[len(tunnels.Tunnel)-1]
    }
    return nil
}

func (tunnels *L2Tpv2_Tunnels) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range tunnels.Tunnel {
        children[tunnels.Tunnel[i].GetSegmentPath()] = &tunnels.Tunnel[i]
    }
    return children
}

func (tunnels *L2Tpv2_Tunnels) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (tunnels *L2Tpv2_Tunnels) GetBundleName() string { return "cisco_ios_xr" }

func (tunnels *L2Tpv2_Tunnels) GetYangName() string { return "tunnels" }

func (tunnels *L2Tpv2_Tunnels) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tunnels *L2Tpv2_Tunnels) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tunnels *L2Tpv2_Tunnels) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tunnels *L2Tpv2_Tunnels) SetParent(parent types.Entity) { tunnels.parent = parent }

func (tunnels *L2Tpv2_Tunnels) GetParent() types.Entity { return tunnels.parent }

func (tunnels *L2Tpv2_Tunnels) GetParentYangName() string { return "l2tpv2" }

// L2Tpv2_Tunnels_Tunnel
// L2TP tunnel  information
type L2Tpv2_Tunnels_Tunnel struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Local tunnel ID. The type is interface{} with
    // range: -2147483648..2147483647.
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

    // Retransmit time distribution in seconds. The type is slice of interface{}
    // with range: 0..65535. Units are second.
    RetransmitTime []interface{}
}

func (tunnel *L2Tpv2_Tunnels_Tunnel) GetFilter() yfilter.YFilter { return tunnel.YFilter }

func (tunnel *L2Tpv2_Tunnels_Tunnel) SetFilter(yf yfilter.YFilter) { tunnel.YFilter = yf }

func (tunnel *L2Tpv2_Tunnels_Tunnel) GetGoName(yname string) string {
    if yname == "local-tunnel-id" { return "LocalTunnelId" }
    if yname == "local-address" { return "LocalAddress" }
    if yname == "remote-address" { return "RemoteAddress" }
    if yname == "local-port" { return "LocalPort" }
    if yname == "remote-port" { return "RemotePort" }
    if yname == "protocol" { return "Protocol" }
    if yname == "is-pmtu-enabled" { return "IsPmtuEnabled" }
    if yname == "remote-tunnel-id" { return "RemoteTunnelId" }
    if yname == "local-tunnel-name" { return "LocalTunnelName" }
    if yname == "remote-tunnel-name" { return "RemoteTunnelName" }
    if yname == "class-name" { return "ClassName" }
    if yname == "active-sessions" { return "ActiveSessions" }
    if yname == "sequence-ns" { return "SequenceNs" }
    if yname == "sequence-nr" { return "SequenceNr" }
    if yname == "local-window-size" { return "LocalWindowSize" }
    if yname == "remote-window-size" { return "RemoteWindowSize" }
    if yname == "retransmission-time" { return "RetransmissionTime" }
    if yname == "maximum-retransmission-time" { return "MaximumRetransmissionTime" }
    if yname == "unsent-queue-size" { return "UnsentQueueSize" }
    if yname == "unsent-maximum-queue-size" { return "UnsentMaximumQueueSize" }
    if yname == "resend-queue-size" { return "ResendQueueSize" }
    if yname == "resend-maximum-queue-size" { return "ResendMaximumQueueSize" }
    if yname == "order-queue-size" { return "OrderQueueSize" }
    if yname == "packet-queue-check" { return "PacketQueueCheck" }
    if yname == "digest-secrets" { return "DigestSecrets" }
    if yname == "resends" { return "Resends" }
    if yname == "zero-length-body-acknowledgement-sent" { return "ZeroLengthBodyAcknowledgementSent" }
    if yname == "total-out-of-order-drop-packets" { return "TotalOutOfOrderDropPackets" }
    if yname == "total-out-of-order-reorder-packets" { return "TotalOutOfOrderReorderPackets" }
    if yname == "total-peer-authentication-failures" { return "TotalPeerAuthenticationFailures" }
    if yname == "is-tunnel-up" { return "IsTunnelUp" }
    if yname == "is-congestion-control-enabled" { return "IsCongestionControlEnabled" }
    if yname == "retransmit-time" { return "RetransmitTime" }
    return ""
}

func (tunnel *L2Tpv2_Tunnels_Tunnel) GetSegmentPath() string {
    return "tunnel" + "[local-tunnel-id='" + fmt.Sprintf("%v", tunnel.LocalTunnelId) + "']"
}

func (tunnel *L2Tpv2_Tunnels_Tunnel) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (tunnel *L2Tpv2_Tunnels_Tunnel) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (tunnel *L2Tpv2_Tunnels_Tunnel) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["local-tunnel-id"] = tunnel.LocalTunnelId
    leafs["local-address"] = tunnel.LocalAddress
    leafs["remote-address"] = tunnel.RemoteAddress
    leafs["local-port"] = tunnel.LocalPort
    leafs["remote-port"] = tunnel.RemotePort
    leafs["protocol"] = tunnel.Protocol
    leafs["is-pmtu-enabled"] = tunnel.IsPmtuEnabled
    leafs["remote-tunnel-id"] = tunnel.RemoteTunnelId
    leafs["local-tunnel-name"] = tunnel.LocalTunnelName
    leafs["remote-tunnel-name"] = tunnel.RemoteTunnelName
    leafs["class-name"] = tunnel.ClassName
    leafs["active-sessions"] = tunnel.ActiveSessions
    leafs["sequence-ns"] = tunnel.SequenceNs
    leafs["sequence-nr"] = tunnel.SequenceNr
    leafs["local-window-size"] = tunnel.LocalWindowSize
    leafs["remote-window-size"] = tunnel.RemoteWindowSize
    leafs["retransmission-time"] = tunnel.RetransmissionTime
    leafs["maximum-retransmission-time"] = tunnel.MaximumRetransmissionTime
    leafs["unsent-queue-size"] = tunnel.UnsentQueueSize
    leafs["unsent-maximum-queue-size"] = tunnel.UnsentMaximumQueueSize
    leafs["resend-queue-size"] = tunnel.ResendQueueSize
    leafs["resend-maximum-queue-size"] = tunnel.ResendMaximumQueueSize
    leafs["order-queue-size"] = tunnel.OrderQueueSize
    leafs["packet-queue-check"] = tunnel.PacketQueueCheck
    leafs["digest-secrets"] = tunnel.DigestSecrets
    leafs["resends"] = tunnel.Resends
    leafs["zero-length-body-acknowledgement-sent"] = tunnel.ZeroLengthBodyAcknowledgementSent
    leafs["total-out-of-order-drop-packets"] = tunnel.TotalOutOfOrderDropPackets
    leafs["total-out-of-order-reorder-packets"] = tunnel.TotalOutOfOrderReorderPackets
    leafs["total-peer-authentication-failures"] = tunnel.TotalPeerAuthenticationFailures
    leafs["is-tunnel-up"] = tunnel.IsTunnelUp
    leafs["is-congestion-control-enabled"] = tunnel.IsCongestionControlEnabled
    leafs["retransmit-time"] = tunnel.RetransmitTime
    return leafs
}

func (tunnel *L2Tpv2_Tunnels_Tunnel) GetBundleName() string { return "cisco_ios_xr" }

func (tunnel *L2Tpv2_Tunnels_Tunnel) GetYangName() string { return "tunnel" }

func (tunnel *L2Tpv2_Tunnels_Tunnel) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tunnel *L2Tpv2_Tunnels_Tunnel) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tunnel *L2Tpv2_Tunnels_Tunnel) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tunnel *L2Tpv2_Tunnels_Tunnel) SetParent(parent types.Entity) { tunnel.parent = parent }

func (tunnel *L2Tpv2_Tunnels_Tunnel) GetParent() types.Entity { return tunnel.parent }

func (tunnel *L2Tpv2_Tunnels_Tunnel) GetParentYangName() string { return "tunnels" }

// L2Tpv2_Sessions
// List of session IDs
type L2Tpv2_Sessions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // L2TP information for a particular session. The type is slice of
    // L2Tpv2_Sessions_Session.
    Session []L2Tpv2_Sessions_Session
}

func (sessions *L2Tpv2_Sessions) GetFilter() yfilter.YFilter { return sessions.YFilter }

func (sessions *L2Tpv2_Sessions) SetFilter(yf yfilter.YFilter) { sessions.YFilter = yf }

func (sessions *L2Tpv2_Sessions) GetGoName(yname string) string {
    if yname == "session" { return "Session" }
    return ""
}

func (sessions *L2Tpv2_Sessions) GetSegmentPath() string {
    return "sessions"
}

func (sessions *L2Tpv2_Sessions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "session" {
        for _, c := range sessions.Session {
            if sessions.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := L2Tpv2_Sessions_Session{}
        sessions.Session = append(sessions.Session, child)
        return &sessions.Session[len(sessions.Session)-1]
    }
    return nil
}

func (sessions *L2Tpv2_Sessions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sessions.Session {
        children[sessions.Session[i].GetSegmentPath()] = &sessions.Session[i]
    }
    return children
}

func (sessions *L2Tpv2_Sessions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sessions *L2Tpv2_Sessions) GetBundleName() string { return "cisco_ios_xr" }

func (sessions *L2Tpv2_Sessions) GetYangName() string { return "sessions" }

func (sessions *L2Tpv2_Sessions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sessions *L2Tpv2_Sessions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sessions *L2Tpv2_Sessions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sessions *L2Tpv2_Sessions) SetParent(parent types.Entity) { sessions.parent = parent }

func (sessions *L2Tpv2_Sessions) GetParent() types.Entity { return sessions.parent }

func (sessions *L2Tpv2_Sessions) GetParentYangName() string { return "l2tpv2" }

// L2Tpv2_Sessions_Session
// L2TP information for a particular session
type L2Tpv2_Sessions_Session struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Local tunnel ID. The type is interface{} with
    // range: -2147483648..2147483647.
    LocalTunnelId interface{}

    // This attribute is a key. Local session ID. The type is interface{} with
    // range: -2147483648..2147483647.
    LocalSessionId interface{}

    // Local session IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    LocalIpAddress interface{}

    // Remote session IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
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

func (session *L2Tpv2_Sessions_Session) GetFilter() yfilter.YFilter { return session.YFilter }

func (session *L2Tpv2_Sessions_Session) SetFilter(yf yfilter.YFilter) { session.YFilter = yf }

func (session *L2Tpv2_Sessions_Session) GetGoName(yname string) string {
    if yname == "local-tunnel-id" { return "LocalTunnelId" }
    if yname == "local-session-id" { return "LocalSessionId" }
    if yname == "local-ip-address" { return "LocalIpAddress" }
    if yname == "remote-ip-address" { return "RemoteIpAddress" }
    if yname == "l2tp-sh-sess-udp-lport" { return "L2TpShSessUdpLport" }
    if yname == "l2tp-sh-sess-udp-rport" { return "L2TpShSessUdpRport" }
    if yname == "protocol" { return "Protocol" }
    if yname == "remote-tunnel-id" { return "RemoteTunnelId" }
    if yname == "call-serial-number" { return "CallSerialNumber" }
    if yname == "local-tunnel-name" { return "LocalTunnelName" }
    if yname == "remote-tunnel-name" { return "RemoteTunnelName" }
    if yname == "remote-session-id" { return "RemoteSessionId" }
    if yname == "l2tp-sh-sess-tie-breaker-enabled" { return "L2TpShSessTieBreakerEnabled" }
    if yname == "l2tp-sh-sess-tie-breaker" { return "L2TpShSessTieBreaker" }
    if yname == "is-session-manual" { return "IsSessionManual" }
    if yname == "is-session-up" { return "IsSessionUp" }
    if yname == "is-udp-checksum-enabled" { return "IsUdpChecksumEnabled" }
    if yname == "is-sequencing-on" { return "IsSequencingOn" }
    if yname == "is-session-state-established" { return "IsSessionStateEstablished" }
    if yname == "is-session-locally-initiated" { return "IsSessionLocallyInitiated" }
    if yname == "is-conditional-debug-enabled" { return "IsConditionalDebugEnabled" }
    if yname == "unique-id" { return "UniqueId" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "session-application-data" { return "SessionApplicationData" }
    return ""
}

func (session *L2Tpv2_Sessions_Session) GetSegmentPath() string {
    return "session" + "[local-tunnel-id='" + fmt.Sprintf("%v", session.LocalTunnelId) + "']" + "[local-session-id='" + fmt.Sprintf("%v", session.LocalSessionId) + "']"
}

func (session *L2Tpv2_Sessions_Session) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "session-application-data" {
        return &session.SessionApplicationData
    }
    return nil
}

func (session *L2Tpv2_Sessions_Session) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["session-application-data"] = &session.SessionApplicationData
    return children
}

func (session *L2Tpv2_Sessions_Session) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["local-tunnel-id"] = session.LocalTunnelId
    leafs["local-session-id"] = session.LocalSessionId
    leafs["local-ip-address"] = session.LocalIpAddress
    leafs["remote-ip-address"] = session.RemoteIpAddress
    leafs["l2tp-sh-sess-udp-lport"] = session.L2TpShSessUdpLport
    leafs["l2tp-sh-sess-udp-rport"] = session.L2TpShSessUdpRport
    leafs["protocol"] = session.Protocol
    leafs["remote-tunnel-id"] = session.RemoteTunnelId
    leafs["call-serial-number"] = session.CallSerialNumber
    leafs["local-tunnel-name"] = session.LocalTunnelName
    leafs["remote-tunnel-name"] = session.RemoteTunnelName
    leafs["remote-session-id"] = session.RemoteSessionId
    leafs["l2tp-sh-sess-tie-breaker-enabled"] = session.L2TpShSessTieBreakerEnabled
    leafs["l2tp-sh-sess-tie-breaker"] = session.L2TpShSessTieBreaker
    leafs["is-session-manual"] = session.IsSessionManual
    leafs["is-session-up"] = session.IsSessionUp
    leafs["is-udp-checksum-enabled"] = session.IsUdpChecksumEnabled
    leafs["is-sequencing-on"] = session.IsSequencingOn
    leafs["is-session-state-established"] = session.IsSessionStateEstablished
    leafs["is-session-locally-initiated"] = session.IsSessionLocallyInitiated
    leafs["is-conditional-debug-enabled"] = session.IsConditionalDebugEnabled
    leafs["unique-id"] = session.UniqueId
    leafs["interface-name"] = session.InterfaceName
    return leafs
}

func (session *L2Tpv2_Sessions_Session) GetBundleName() string { return "cisco_ios_xr" }

func (session *L2Tpv2_Sessions_Session) GetYangName() string { return "session" }

func (session *L2Tpv2_Sessions_Session) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (session *L2Tpv2_Sessions_Session) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (session *L2Tpv2_Sessions_Session) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (session *L2Tpv2_Sessions_Session) SetParent(parent types.Entity) { session.parent = parent }

func (session *L2Tpv2_Sessions_Session) GetParent() types.Entity { return session.parent }

func (session *L2Tpv2_Sessions_Session) GetParentYangName() string { return "sessions" }

// L2Tpv2_Sessions_Session_SessionApplicationData
// Session application data
type L2Tpv2_Sessions_Session_SessionApplicationData struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // l2tp sh sess app type. The type is interface{} with range: 0..4294967295.
    L2TpShSessAppType interface{}

    // Xconnect data.
    Xconnect L2Tpv2_Sessions_Session_SessionApplicationData_Xconnect

    // VPDN data.
    Vpdn L2Tpv2_Sessions_Session_SessionApplicationData_Vpdn
}

func (sessionApplicationData *L2Tpv2_Sessions_Session_SessionApplicationData) GetFilter() yfilter.YFilter { return sessionApplicationData.YFilter }

func (sessionApplicationData *L2Tpv2_Sessions_Session_SessionApplicationData) SetFilter(yf yfilter.YFilter) { sessionApplicationData.YFilter = yf }

func (sessionApplicationData *L2Tpv2_Sessions_Session_SessionApplicationData) GetGoName(yname string) string {
    if yname == "l2tp-sh-sess-app-type" { return "L2TpShSessAppType" }
    if yname == "xconnect" { return "Xconnect" }
    if yname == "vpdn" { return "Vpdn" }
    return ""
}

func (sessionApplicationData *L2Tpv2_Sessions_Session_SessionApplicationData) GetSegmentPath() string {
    return "session-application-data"
}

func (sessionApplicationData *L2Tpv2_Sessions_Session_SessionApplicationData) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "xconnect" {
        return &sessionApplicationData.Xconnect
    }
    if childYangName == "vpdn" {
        return &sessionApplicationData.Vpdn
    }
    return nil
}

func (sessionApplicationData *L2Tpv2_Sessions_Session_SessionApplicationData) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["xconnect"] = &sessionApplicationData.Xconnect
    children["vpdn"] = &sessionApplicationData.Vpdn
    return children
}

func (sessionApplicationData *L2Tpv2_Sessions_Session_SessionApplicationData) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["l2tp-sh-sess-app-type"] = sessionApplicationData.L2TpShSessAppType
    return leafs
}

func (sessionApplicationData *L2Tpv2_Sessions_Session_SessionApplicationData) GetBundleName() string { return "cisco_ios_xr" }

func (sessionApplicationData *L2Tpv2_Sessions_Session_SessionApplicationData) GetYangName() string { return "session-application-data" }

func (sessionApplicationData *L2Tpv2_Sessions_Session_SessionApplicationData) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sessionApplicationData *L2Tpv2_Sessions_Session_SessionApplicationData) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sessionApplicationData *L2Tpv2_Sessions_Session_SessionApplicationData) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sessionApplicationData *L2Tpv2_Sessions_Session_SessionApplicationData) SetParent(parent types.Entity) { sessionApplicationData.parent = parent }

func (sessionApplicationData *L2Tpv2_Sessions_Session_SessionApplicationData) GetParent() types.Entity { return sessionApplicationData.parent }

func (sessionApplicationData *L2Tpv2_Sessions_Session_SessionApplicationData) GetParentYangName() string { return "session" }

// L2Tpv2_Sessions_Session_SessionApplicationData_Xconnect
// Xconnect data
type L2Tpv2_Sessions_Session_SessionApplicationData_Xconnect struct {
    parent types.Entity
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

func (xconnect *L2Tpv2_Sessions_Session_SessionApplicationData_Xconnect) GetFilter() yfilter.YFilter { return xconnect.YFilter }

func (xconnect *L2Tpv2_Sessions_Session_SessionApplicationData_Xconnect) SetFilter(yf yfilter.YFilter) { xconnect.YFilter = yf }

func (xconnect *L2Tpv2_Sessions_Session_SessionApplicationData_Xconnect) GetGoName(yname string) string {
    if yname == "circuit-name" { return "CircuitName" }
    if yname == "sessionvc-id" { return "SessionvcId" }
    if yname == "is-circuit-state-up" { return "IsCircuitStateUp" }
    if yname == "is-local-circuit-state-up" { return "IsLocalCircuitStateUp" }
    if yname == "is-remote-circuit-state-up" { return "IsRemoteCircuitStateUp" }
    if yname == "ipv6-protocol-tunneling" { return "Ipv6ProtocolTunneling" }
    return ""
}

func (xconnect *L2Tpv2_Sessions_Session_SessionApplicationData_Xconnect) GetSegmentPath() string {
    return "xconnect"
}

func (xconnect *L2Tpv2_Sessions_Session_SessionApplicationData_Xconnect) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (xconnect *L2Tpv2_Sessions_Session_SessionApplicationData_Xconnect) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (xconnect *L2Tpv2_Sessions_Session_SessionApplicationData_Xconnect) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["circuit-name"] = xconnect.CircuitName
    leafs["sessionvc-id"] = xconnect.SessionvcId
    leafs["is-circuit-state-up"] = xconnect.IsCircuitStateUp
    leafs["is-local-circuit-state-up"] = xconnect.IsLocalCircuitStateUp
    leafs["is-remote-circuit-state-up"] = xconnect.IsRemoteCircuitStateUp
    leafs["ipv6-protocol-tunneling"] = xconnect.Ipv6ProtocolTunneling
    return leafs
}

func (xconnect *L2Tpv2_Sessions_Session_SessionApplicationData_Xconnect) GetBundleName() string { return "cisco_ios_xr" }

func (xconnect *L2Tpv2_Sessions_Session_SessionApplicationData_Xconnect) GetYangName() string { return "xconnect" }

func (xconnect *L2Tpv2_Sessions_Session_SessionApplicationData_Xconnect) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (xconnect *L2Tpv2_Sessions_Session_SessionApplicationData_Xconnect) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (xconnect *L2Tpv2_Sessions_Session_SessionApplicationData_Xconnect) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (xconnect *L2Tpv2_Sessions_Session_SessionApplicationData_Xconnect) SetParent(parent types.Entity) { xconnect.parent = parent }

func (xconnect *L2Tpv2_Sessions_Session_SessionApplicationData_Xconnect) GetParent() types.Entity { return xconnect.parent }

func (xconnect *L2Tpv2_Sessions_Session_SessionApplicationData_Xconnect) GetParentYangName() string { return "session-application-data" }

// L2Tpv2_Sessions_Session_SessionApplicationData_Vpdn
// VPDN data
type L2Tpv2_Sessions_Session_SessionApplicationData_Vpdn struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Session username. The type is string.
    Username interface{}

    // Interface name. The type is string with pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}
}

func (vpdn *L2Tpv2_Sessions_Session_SessionApplicationData_Vpdn) GetFilter() yfilter.YFilter { return vpdn.YFilter }

func (vpdn *L2Tpv2_Sessions_Session_SessionApplicationData_Vpdn) SetFilter(yf yfilter.YFilter) { vpdn.YFilter = yf }

func (vpdn *L2Tpv2_Sessions_Session_SessionApplicationData_Vpdn) GetGoName(yname string) string {
    if yname == "username" { return "Username" }
    if yname == "interface-name" { return "InterfaceName" }
    return ""
}

func (vpdn *L2Tpv2_Sessions_Session_SessionApplicationData_Vpdn) GetSegmentPath() string {
    return "vpdn"
}

func (vpdn *L2Tpv2_Sessions_Session_SessionApplicationData_Vpdn) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vpdn *L2Tpv2_Sessions_Session_SessionApplicationData_Vpdn) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vpdn *L2Tpv2_Sessions_Session_SessionApplicationData_Vpdn) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["username"] = vpdn.Username
    leafs["interface-name"] = vpdn.InterfaceName
    return leafs
}

func (vpdn *L2Tpv2_Sessions_Session_SessionApplicationData_Vpdn) GetBundleName() string { return "cisco_ios_xr" }

func (vpdn *L2Tpv2_Sessions_Session_SessionApplicationData_Vpdn) GetYangName() string { return "vpdn" }

func (vpdn *L2Tpv2_Sessions_Session_SessionApplicationData_Vpdn) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vpdn *L2Tpv2_Sessions_Session_SessionApplicationData_Vpdn) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vpdn *L2Tpv2_Sessions_Session_SessionApplicationData_Vpdn) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vpdn *L2Tpv2_Sessions_Session_SessionApplicationData_Vpdn) SetParent(parent types.Entity) { vpdn.parent = parent }

func (vpdn *L2Tpv2_Sessions_Session_SessionApplicationData_Vpdn) GetParent() types.Entity { return vpdn.parent }

func (vpdn *L2Tpv2_Sessions_Session_SessionApplicationData_Vpdn) GetParentYangName() string { return "session-application-data" }

// L2Tpv2_Session
// L2TP control messages counters
type L2Tpv2_Session struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // L2TP session unavailable  information.
    Unavailable L2Tpv2_Session_Unavailable
}

func (session *L2Tpv2_Session) GetFilter() yfilter.YFilter { return session.YFilter }

func (session *L2Tpv2_Session) SetFilter(yf yfilter.YFilter) { session.YFilter = yf }

func (session *L2Tpv2_Session) GetGoName(yname string) string {
    if yname == "unavailable" { return "Unavailable" }
    return ""
}

func (session *L2Tpv2_Session) GetSegmentPath() string {
    return "session"
}

func (session *L2Tpv2_Session) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "unavailable" {
        return &session.Unavailable
    }
    return nil
}

func (session *L2Tpv2_Session) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["unavailable"] = &session.Unavailable
    return children
}

func (session *L2Tpv2_Session) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (session *L2Tpv2_Session) GetBundleName() string { return "cisco_ios_xr" }

func (session *L2Tpv2_Session) GetYangName() string { return "session" }

func (session *L2Tpv2_Session) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (session *L2Tpv2_Session) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (session *L2Tpv2_Session) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (session *L2Tpv2_Session) SetParent(parent types.Entity) { session.parent = parent }

func (session *L2Tpv2_Session) GetParent() types.Entity { return session.parent }

func (session *L2Tpv2_Session) GetParentYangName() string { return "l2tpv2" }

// L2Tpv2_Session_Unavailable
// L2TP session unavailable  information
type L2Tpv2_Session_Unavailable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of session ID in hold database. The type is interface{} with range:
    // 0..4294967295.
    SessionsOnHold interface{}
}

func (unavailable *L2Tpv2_Session_Unavailable) GetFilter() yfilter.YFilter { return unavailable.YFilter }

func (unavailable *L2Tpv2_Session_Unavailable) SetFilter(yf yfilter.YFilter) { unavailable.YFilter = yf }

func (unavailable *L2Tpv2_Session_Unavailable) GetGoName(yname string) string {
    if yname == "sessions-on-hold" { return "SessionsOnHold" }
    return ""
}

func (unavailable *L2Tpv2_Session_Unavailable) GetSegmentPath() string {
    return "unavailable"
}

func (unavailable *L2Tpv2_Session_Unavailable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (unavailable *L2Tpv2_Session_Unavailable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (unavailable *L2Tpv2_Session_Unavailable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sessions-on-hold"] = unavailable.SessionsOnHold
    return leafs
}

func (unavailable *L2Tpv2_Session_Unavailable) GetBundleName() string { return "cisco_ios_xr" }

func (unavailable *L2Tpv2_Session_Unavailable) GetYangName() string { return "unavailable" }

func (unavailable *L2Tpv2_Session_Unavailable) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (unavailable *L2Tpv2_Session_Unavailable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (unavailable *L2Tpv2_Session_Unavailable) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (unavailable *L2Tpv2_Session_Unavailable) SetParent(parent types.Entity) { unavailable.parent = parent }

func (unavailable *L2Tpv2_Session_Unavailable) GetParent() types.Entity { return unavailable.parent }

func (unavailable *L2Tpv2_Session_Unavailable) GetParentYangName() string { return "session" }

