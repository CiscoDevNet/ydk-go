// This module contains a collection of YANG definitions
// for Cisco IOS-XR man-ems package operational data.
// 
// This module contains definitions
// for the following management objects:
//   grpc: grpc commands
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package man_ems_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package man_ems_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-man-ems-oper grpc}", reflect.TypeOf(Grpc{}))
    ydk.RegisterEntity("Cisco-IOS-XR-man-ems-oper:grpc", reflect.TypeOf(Grpc{}))
}

// Grpc
// grpc commands
type Grpc struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Grpc Statistics.
    Statistics Grpc_Statistics

    // Grpc Status.
    Status Grpc_Status
}

func (grpc *Grpc) GetFilter() yfilter.YFilter { return grpc.YFilter }

func (grpc *Grpc) SetFilter(yf yfilter.YFilter) { grpc.YFilter = yf }

func (grpc *Grpc) GetGoName(yname string) string {
    if yname == "statistics" { return "Statistics" }
    if yname == "status" { return "Status" }
    return ""
}

func (grpc *Grpc) GetSegmentPath() string {
    return "Cisco-IOS-XR-man-ems-oper:grpc"
}

func (grpc *Grpc) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "statistics" {
        return &grpc.Statistics
    }
    if childYangName == "status" {
        return &grpc.Status
    }
    return nil
}

func (grpc *Grpc) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["statistics"] = &grpc.Statistics
    children["status"] = &grpc.Status
    return children
}

func (grpc *Grpc) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (grpc *Grpc) GetBundleName() string { return "cisco_ios_xr" }

func (grpc *Grpc) GetYangName() string { return "grpc" }

func (grpc *Grpc) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (grpc *Grpc) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (grpc *Grpc) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (grpc *Grpc) SetParent(parent types.Entity) { grpc.parent = parent }

func (grpc *Grpc) GetParent() types.Entity { return grpc.parent }

func (grpc *Grpc) GetParentYangName() string { return "Cisco-IOS-XR-man-ems-oper" }

// Grpc_Statistics
// Grpc Statistics
type Grpc_Statistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // CounterShowCmdTxtReqRecv. The type is interface{} with range:
    // 0..18446744073709551615.
    CtShowCmdTxtReqRecv interface{}

    // CounterShowCmdTxtResSent. The type is interface{} with range:
    // 0..18446744073709551615.
    CtShowCmdTxtResSent interface{}

    // CounterGetConfigReqRecv. The type is interface{} with range:
    // 0..18446744073709551615.
    CtGetConfigReqRecv interface{}

    // CounterGetConfigResSent. The type is interface{} with range:
    // 0..18446744073709551615.
    CtGetConfigResSent interface{}

    // CounterCliConfigReqRecv. The type is interface{} with range:
    // 0..18446744073709551615.
    CtCliConfigReqRecv interface{}

    // CounterCliConfigResSent. The type is interface{} with range:
    // 0..18446744073709551615.
    CtCliConfigResSent interface{}

    // CounterMergeConfigReq. The type is interface{} with range:
    // 0..18446744073709551615.
    CtMergeConfigReqRecv interface{}

    // CounterMergeConfigRes. The type is interface{} with range:
    // 0..18446744073709551615.
    CtMergeConfigResSent interface{}

    // CounterCommitReplaceReq. The type is interface{} with range:
    // 0..18446744073709551615.
    CtCommitReplaceReqRecv interface{}

    // CounterCommitReplaceRes. The type is interface{} with range:
    // 0..18446744073709551615.
    CtCommitReplaceResSent interface{}

    // CounterDeleteConfigReq. The type is interface{} with range:
    // 0..18446744073709551615.
    CtDeleteConfigReqRecv interface{}

    // CounterDeleteConfigRes. The type is interface{} with range:
    // 0..18446744073709551615.
    CtDeleteConfigResSent interface{}

    // CounterReplaceConfigReq. The type is interface{} with range:
    // 0..18446744073709551615.
    CtReplaceConfigReqRecv interface{}

    // CounterReplaceConfigSent. The type is interface{} with range:
    // 0..18446744073709551615.
    CtReplaceConfigResSent interface{}

    // CounterGetOperReqRecv. The type is interface{} with range:
    // 0..18446744073709551615.
    CtGetOperReqRecv interface{}

    // CounterGetOperResSent. The type is interface{} with range:
    // 0..18446744073709551615.
    CtGetOperResSent interface{}

    // CounterGetCurrentSession. The type is interface{} with range:
    // 0..4294967295.
    CtGetCurrentSession interface{}

    // CounterForHowManyCommitConfigRequests. The type is interface{} with range:
    // 0..18446744073709551615.
    CtCommitConfigReqRecv interface{}

    // CounterForHowManyCommitConfigResponses. The type is interface{} with range:
    // 0..18446744073709551615.
    CtCommitConfigResSent interface{}

    // CounterForHowManyActionJsonRequests. The type is interface{} with range:
    // 0..18446744073709551615.
    CtActionJsonReqRecv interface{}

    // CounterForHowManyActionJsonResponses. The type is interface{} with range:
    // 0..18446744073709551615.
    CtActionJsonResSent interface{}
}

func (statistics *Grpc_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *Grpc_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *Grpc_Statistics) GetGoName(yname string) string {
    if yname == "ct-show-cmd-txt-req-recv" { return "CtShowCmdTxtReqRecv" }
    if yname == "ct-show-cmd-txt-res-sent" { return "CtShowCmdTxtResSent" }
    if yname == "ct-get-config-req-recv" { return "CtGetConfigReqRecv" }
    if yname == "ct-get-config-res-sent" { return "CtGetConfigResSent" }
    if yname == "ct-cli-config-req-recv" { return "CtCliConfigReqRecv" }
    if yname == "ct-cli-config-res-sent" { return "CtCliConfigResSent" }
    if yname == "ct-merge-config-req-recv" { return "CtMergeConfigReqRecv" }
    if yname == "ct-merge-config-res-sent" { return "CtMergeConfigResSent" }
    if yname == "ct-commit-replace-req-recv" { return "CtCommitReplaceReqRecv" }
    if yname == "ct-commit-replace-res-sent" { return "CtCommitReplaceResSent" }
    if yname == "ct-delete-config-req-recv" { return "CtDeleteConfigReqRecv" }
    if yname == "ct-delete-config-res-sent" { return "CtDeleteConfigResSent" }
    if yname == "ct-replace-config-req-recv" { return "CtReplaceConfigReqRecv" }
    if yname == "ct-replace-config-res-sent" { return "CtReplaceConfigResSent" }
    if yname == "ct-get-oper-req-recv" { return "CtGetOperReqRecv" }
    if yname == "ct-get-oper-res-sent" { return "CtGetOperResSent" }
    if yname == "ct-get-current-session" { return "CtGetCurrentSession" }
    if yname == "ct-commit-config-req-recv" { return "CtCommitConfigReqRecv" }
    if yname == "ct-commit-config-res-sent" { return "CtCommitConfigResSent" }
    if yname == "ct-action-json-req-recv" { return "CtActionJsonReqRecv" }
    if yname == "ct-action-json-res-sent" { return "CtActionJsonResSent" }
    return ""
}

func (statistics *Grpc_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *Grpc_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (statistics *Grpc_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (statistics *Grpc_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ct-show-cmd-txt-req-recv"] = statistics.CtShowCmdTxtReqRecv
    leafs["ct-show-cmd-txt-res-sent"] = statistics.CtShowCmdTxtResSent
    leafs["ct-get-config-req-recv"] = statistics.CtGetConfigReqRecv
    leafs["ct-get-config-res-sent"] = statistics.CtGetConfigResSent
    leafs["ct-cli-config-req-recv"] = statistics.CtCliConfigReqRecv
    leafs["ct-cli-config-res-sent"] = statistics.CtCliConfigResSent
    leafs["ct-merge-config-req-recv"] = statistics.CtMergeConfigReqRecv
    leafs["ct-merge-config-res-sent"] = statistics.CtMergeConfigResSent
    leafs["ct-commit-replace-req-recv"] = statistics.CtCommitReplaceReqRecv
    leafs["ct-commit-replace-res-sent"] = statistics.CtCommitReplaceResSent
    leafs["ct-delete-config-req-recv"] = statistics.CtDeleteConfigReqRecv
    leafs["ct-delete-config-res-sent"] = statistics.CtDeleteConfigResSent
    leafs["ct-replace-config-req-recv"] = statistics.CtReplaceConfigReqRecv
    leafs["ct-replace-config-res-sent"] = statistics.CtReplaceConfigResSent
    leafs["ct-get-oper-req-recv"] = statistics.CtGetOperReqRecv
    leafs["ct-get-oper-res-sent"] = statistics.CtGetOperResSent
    leafs["ct-get-current-session"] = statistics.CtGetCurrentSession
    leafs["ct-commit-config-req-recv"] = statistics.CtCommitConfigReqRecv
    leafs["ct-commit-config-res-sent"] = statistics.CtCommitConfigResSent
    leafs["ct-action-json-req-recv"] = statistics.CtActionJsonReqRecv
    leafs["ct-action-json-res-sent"] = statistics.CtActionJsonResSent
    return leafs
}

func (statistics *Grpc_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *Grpc_Statistics) GetYangName() string { return "statistics" }

func (statistics *Grpc_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *Grpc_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *Grpc_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *Grpc_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *Grpc_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *Grpc_Statistics) GetParentYangName() string { return "grpc" }

// Grpc_Status
// Grpc Status
type Grpc_Status struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // GRPCTransport. The type is string.
    Transport interface{}

    // AddressFamily. The type is string.
    AddressFamily interface{}

    // GRPCTLS. The type is string.
    Tls interface{}

    // GRPCTrustpoint. The type is string.
    Trustpoint interface{}

    // ListeningPort. The type is interface{} with range: -2147483648..2147483647.
    ListeningPort interface{}

    // VrfSocketNamespacePath. The type is string.
    VrfSocketNsPath interface{}

    // MaxReqPerUser. The type is interface{} with range: 0..4294967295.
    MaxReqPerUser interface{}

    // MaxReqTotal. The type is interface{} with range: 0..4294967295.
    MaxReqTotal interface{}
}

func (status *Grpc_Status) GetFilter() yfilter.YFilter { return status.YFilter }

func (status *Grpc_Status) SetFilter(yf yfilter.YFilter) { status.YFilter = yf }

func (status *Grpc_Status) GetGoName(yname string) string {
    if yname == "transport" { return "Transport" }
    if yname == "address-family" { return "AddressFamily" }
    if yname == "tls" { return "Tls" }
    if yname == "trustpoint" { return "Trustpoint" }
    if yname == "listening-port" { return "ListeningPort" }
    if yname == "vrf-socket-ns-path" { return "VrfSocketNsPath" }
    if yname == "max-req-per-user" { return "MaxReqPerUser" }
    if yname == "max-req-total" { return "MaxReqTotal" }
    return ""
}

func (status *Grpc_Status) GetSegmentPath() string {
    return "status"
}

func (status *Grpc_Status) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (status *Grpc_Status) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (status *Grpc_Status) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["transport"] = status.Transport
    leafs["address-family"] = status.AddressFamily
    leafs["tls"] = status.Tls
    leafs["trustpoint"] = status.Trustpoint
    leafs["listening-port"] = status.ListeningPort
    leafs["vrf-socket-ns-path"] = status.VrfSocketNsPath
    leafs["max-req-per-user"] = status.MaxReqPerUser
    leafs["max-req-total"] = status.MaxReqTotal
    return leafs
}

func (status *Grpc_Status) GetBundleName() string { return "cisco_ios_xr" }

func (status *Grpc_Status) GetYangName() string { return "status" }

func (status *Grpc_Status) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (status *Grpc_Status) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (status *Grpc_Status) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (status *Grpc_Status) SetParent(parent types.Entity) { status.parent = parent }

func (status *Grpc_Status) GetParent() types.Entity { return status.parent }

func (status *Grpc_Status) GetParentYangName() string { return "grpc" }

