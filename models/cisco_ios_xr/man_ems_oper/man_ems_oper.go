// This module contains a collection of YANG definitions
// for Cisco IOS-XR man-ems package operational data.
// 
// This module contains definitions
// for the following management objects:
//   grpc: grpc commands
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Grpc Statistics.
    Statistics Grpc_Statistics

    // Grpc Status.
    Status Grpc_Status
}

func (grpc *Grpc) GetEntityData() *types.CommonEntityData {
    grpc.EntityData.YFilter = grpc.YFilter
    grpc.EntityData.YangName = "grpc"
    grpc.EntityData.BundleName = "cisco_ios_xr"
    grpc.EntityData.ParentYangName = "Cisco-IOS-XR-man-ems-oper"
    grpc.EntityData.SegmentPath = "Cisco-IOS-XR-man-ems-oper:grpc"
    grpc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    grpc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    grpc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    grpc.EntityData.Children = types.NewOrderedMap()
    grpc.EntityData.Children.Append("statistics", types.YChild{"Statistics", &grpc.Statistics})
    grpc.EntityData.Children.Append("status", types.YChild{"Status", &grpc.Status})
    grpc.EntityData.Leafs = types.NewOrderedMap()

    grpc.EntityData.YListKeys = []string {}

    return &(grpc.EntityData)
}

// Grpc_Statistics
// Grpc Statistics
type Grpc_Statistics struct {
    EntityData types.CommonEntityData
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

func (statistics *Grpc_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "grpc"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = types.NewOrderedMap()
    statistics.EntityData.Leafs = types.NewOrderedMap()
    statistics.EntityData.Leafs.Append("ct-show-cmd-txt-req-recv", types.YLeaf{"CtShowCmdTxtReqRecv", statistics.CtShowCmdTxtReqRecv})
    statistics.EntityData.Leafs.Append("ct-show-cmd-txt-res-sent", types.YLeaf{"CtShowCmdTxtResSent", statistics.CtShowCmdTxtResSent})
    statistics.EntityData.Leafs.Append("ct-get-config-req-recv", types.YLeaf{"CtGetConfigReqRecv", statistics.CtGetConfigReqRecv})
    statistics.EntityData.Leafs.Append("ct-get-config-res-sent", types.YLeaf{"CtGetConfigResSent", statistics.CtGetConfigResSent})
    statistics.EntityData.Leafs.Append("ct-cli-config-req-recv", types.YLeaf{"CtCliConfigReqRecv", statistics.CtCliConfigReqRecv})
    statistics.EntityData.Leafs.Append("ct-cli-config-res-sent", types.YLeaf{"CtCliConfigResSent", statistics.CtCliConfigResSent})
    statistics.EntityData.Leafs.Append("ct-merge-config-req-recv", types.YLeaf{"CtMergeConfigReqRecv", statistics.CtMergeConfigReqRecv})
    statistics.EntityData.Leafs.Append("ct-merge-config-res-sent", types.YLeaf{"CtMergeConfigResSent", statistics.CtMergeConfigResSent})
    statistics.EntityData.Leafs.Append("ct-commit-replace-req-recv", types.YLeaf{"CtCommitReplaceReqRecv", statistics.CtCommitReplaceReqRecv})
    statistics.EntityData.Leafs.Append("ct-commit-replace-res-sent", types.YLeaf{"CtCommitReplaceResSent", statistics.CtCommitReplaceResSent})
    statistics.EntityData.Leafs.Append("ct-delete-config-req-recv", types.YLeaf{"CtDeleteConfigReqRecv", statistics.CtDeleteConfigReqRecv})
    statistics.EntityData.Leafs.Append("ct-delete-config-res-sent", types.YLeaf{"CtDeleteConfigResSent", statistics.CtDeleteConfigResSent})
    statistics.EntityData.Leafs.Append("ct-replace-config-req-recv", types.YLeaf{"CtReplaceConfigReqRecv", statistics.CtReplaceConfigReqRecv})
    statistics.EntityData.Leafs.Append("ct-replace-config-res-sent", types.YLeaf{"CtReplaceConfigResSent", statistics.CtReplaceConfigResSent})
    statistics.EntityData.Leafs.Append("ct-get-oper-req-recv", types.YLeaf{"CtGetOperReqRecv", statistics.CtGetOperReqRecv})
    statistics.EntityData.Leafs.Append("ct-get-oper-res-sent", types.YLeaf{"CtGetOperResSent", statistics.CtGetOperResSent})
    statistics.EntityData.Leafs.Append("ct-get-current-session", types.YLeaf{"CtGetCurrentSession", statistics.CtGetCurrentSession})
    statistics.EntityData.Leafs.Append("ct-commit-config-req-recv", types.YLeaf{"CtCommitConfigReqRecv", statistics.CtCommitConfigReqRecv})
    statistics.EntityData.Leafs.Append("ct-commit-config-res-sent", types.YLeaf{"CtCommitConfigResSent", statistics.CtCommitConfigResSent})
    statistics.EntityData.Leafs.Append("ct-action-json-req-recv", types.YLeaf{"CtActionJsonReqRecv", statistics.CtActionJsonReqRecv})
    statistics.EntityData.Leafs.Append("ct-action-json-res-sent", types.YLeaf{"CtActionJsonResSent", statistics.CtActionJsonResSent})

    statistics.EntityData.YListKeys = []string {}

    return &(statistics.EntityData)
}

// Grpc_Status
// Grpc Status
type Grpc_Status struct {
    EntityData types.CommonEntityData
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

    // Maximum number of streaming gRPCs. The type is interface{} with range:
    // 0..4294967295.
    MaxStreams interface{}

    // Maximum number of streaming gRPCs per user. The type is interface{} with
    // range: 0..4294967295.
    MaxStreamsPerUser interface{}
}

func (status *Grpc_Status) GetEntityData() *types.CommonEntityData {
    status.EntityData.YFilter = status.YFilter
    status.EntityData.YangName = "status"
    status.EntityData.BundleName = "cisco_ios_xr"
    status.EntityData.ParentYangName = "grpc"
    status.EntityData.SegmentPath = "status"
    status.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    status.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    status.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    status.EntityData.Children = types.NewOrderedMap()
    status.EntityData.Leafs = types.NewOrderedMap()
    status.EntityData.Leafs.Append("transport", types.YLeaf{"Transport", status.Transport})
    status.EntityData.Leafs.Append("address-family", types.YLeaf{"AddressFamily", status.AddressFamily})
    status.EntityData.Leafs.Append("tls", types.YLeaf{"Tls", status.Tls})
    status.EntityData.Leafs.Append("trustpoint", types.YLeaf{"Trustpoint", status.Trustpoint})
    status.EntityData.Leafs.Append("listening-port", types.YLeaf{"ListeningPort", status.ListeningPort})
    status.EntityData.Leafs.Append("vrf-socket-ns-path", types.YLeaf{"VrfSocketNsPath", status.VrfSocketNsPath})
    status.EntityData.Leafs.Append("max-req-per-user", types.YLeaf{"MaxReqPerUser", status.MaxReqPerUser})
    status.EntityData.Leafs.Append("max-req-total", types.YLeaf{"MaxReqTotal", status.MaxReqTotal})
    status.EntityData.Leafs.Append("max-streams", types.YLeaf{"MaxStreams", status.MaxStreams})
    status.EntityData.Leafs.Append("max-streams-per-user", types.YLeaf{"MaxStreamsPerUser", status.MaxStreamsPerUser})

    status.EntityData.YListKeys = []string {}

    return &(status.EntityData)
}

