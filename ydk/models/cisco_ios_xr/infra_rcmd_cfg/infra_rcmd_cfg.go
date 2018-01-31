// This module contains a collection of YANG definitions
// for Cisco IOS-XR infra-rcmd package configuration.
// 
// This module contains definitions
// for the following management objects:
//   router-convergence: Configure Router Convergence Monitoring
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package infra_rcmd_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package infra_rcmd_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-infra-rcmd-cfg router-convergence}", reflect.TypeOf(RouterConvergence{}))
    ydk.RegisterEntity("Cisco-IOS-XR-infra-rcmd-cfg:router-convergence", reflect.TypeOf(RouterConvergence{}))
}

// RcmdPriority represents Rcmd priority
type RcmdPriority string

const (
    // Critical routes
    RcmdPriority_critical RcmdPriority = "critical"

    // High priority routes
    RcmdPriority_high RcmdPriority = "high"

    // Medium priority routes
    RcmdPriority_medium RcmdPriority = "medium"

    // Low priority routes
    RcmdPriority_low RcmdPriority = "low"
)

// ProtocolName represents Protocol name
type ProtocolName string

const (
    // Configure parameters related to OSPF
    ProtocolName_ospf ProtocolName = "ospf"

    // Configure parameters related to ISIS
    ProtocolName_isis ProtocolName = "isis"
)

// RouterConvergence
// Configure Router Convergence Monitoring
type RouterConvergence struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Event buffer size for storing event traces (as number of events). The type
    // is interface{} with range: 100..500.
    EventBufferSize interface{}

    // Limits Individual Prefix Monitoring. The type is interface{} with range:
    // 0..100.
    PrefixMonitorLimit interface{}

    // Disable the monitoring of route convergence on the entire router. The type
    // is interface{}.
    Disable interface{}

    // Enable Configure Router Convergence Monitoring. Deletion of this object
    // also causes deletion of all associated objects under RouterConvergence. The
    // type is interface{}.
    Enable interface{}

    // Maximum number of events stored in the server. The type is interface{} with
    // range: 10..500.
    MaxEventsStored interface{}

    // Interval in which to collect logs (in mins). The type is interface{} with
    // range: 5..120. Units are minute.
    MonitoringInterval interface{}

    // Table of Protocol.
    Protocols RouterConvergence_Protocols

    // Absolute directory path for saving the archive files. Example /disk0:/rcmd/
    // or <tftp-location>/rcmd/.
    StorageLocation RouterConvergence_StorageLocation

    // RCMD related configuration for MPLS-LDP.
    MplsLdp RouterConvergence_MplsLdp

    // Table of CollectDiagnostics.
    CollectDiagnostics RouterConvergence_CollectDiagnostics

    // Table of Node.
    Nodes RouterConvergence_Nodes
}

func (routerConvergence *RouterConvergence) GetFilter() yfilter.YFilter { return routerConvergence.YFilter }

func (routerConvergence *RouterConvergence) SetFilter(yf yfilter.YFilter) { routerConvergence.YFilter = yf }

func (routerConvergence *RouterConvergence) GetGoName(yname string) string {
    if yname == "event-buffer-size" { return "EventBufferSize" }
    if yname == "prefix-monitor-limit" { return "PrefixMonitorLimit" }
    if yname == "disable" { return "Disable" }
    if yname == "enable" { return "Enable" }
    if yname == "max-events-stored" { return "MaxEventsStored" }
    if yname == "monitoring-interval" { return "MonitoringInterval" }
    if yname == "protocols" { return "Protocols" }
    if yname == "storage-location" { return "StorageLocation" }
    if yname == "mpls-ldp" { return "MplsLdp" }
    if yname == "collect-diagnostics" { return "CollectDiagnostics" }
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (routerConvergence *RouterConvergence) GetSegmentPath() string {
    return "Cisco-IOS-XR-infra-rcmd-cfg:router-convergence"
}

func (routerConvergence *RouterConvergence) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "protocols" {
        return &routerConvergence.Protocols
    }
    if childYangName == "storage-location" {
        return &routerConvergence.StorageLocation
    }
    if childYangName == "mpls-ldp" {
        return &routerConvergence.MplsLdp
    }
    if childYangName == "collect-diagnostics" {
        return &routerConvergence.CollectDiagnostics
    }
    if childYangName == "nodes" {
        return &routerConvergence.Nodes
    }
    return nil
}

func (routerConvergence *RouterConvergence) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["protocols"] = &routerConvergence.Protocols
    children["storage-location"] = &routerConvergence.StorageLocation
    children["mpls-ldp"] = &routerConvergence.MplsLdp
    children["collect-diagnostics"] = &routerConvergence.CollectDiagnostics
    children["nodes"] = &routerConvergence.Nodes
    return children
}

func (routerConvergence *RouterConvergence) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["event-buffer-size"] = routerConvergence.EventBufferSize
    leafs["prefix-monitor-limit"] = routerConvergence.PrefixMonitorLimit
    leafs["disable"] = routerConvergence.Disable
    leafs["enable"] = routerConvergence.Enable
    leafs["max-events-stored"] = routerConvergence.MaxEventsStored
    leafs["monitoring-interval"] = routerConvergence.MonitoringInterval
    return leafs
}

func (routerConvergence *RouterConvergence) GetBundleName() string { return "cisco_ios_xr" }

func (routerConvergence *RouterConvergence) GetYangName() string { return "router-convergence" }

func (routerConvergence *RouterConvergence) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routerConvergence *RouterConvergence) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routerConvergence *RouterConvergence) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routerConvergence *RouterConvergence) SetParent(parent types.Entity) { routerConvergence.parent = parent }

func (routerConvergence *RouterConvergence) GetParent() types.Entity { return routerConvergence.parent }

func (routerConvergence *RouterConvergence) GetParentYangName() string { return "Cisco-IOS-XR-infra-rcmd-cfg" }

// RouterConvergence_Protocols
// Table of Protocol
type RouterConvergence_Protocols struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Protocol for which to configure RCMD parameters. The type is slice of
    // RouterConvergence_Protocols_Protocol.
    Protocol []RouterConvergence_Protocols_Protocol
}

func (protocols *RouterConvergence_Protocols) GetFilter() yfilter.YFilter { return protocols.YFilter }

func (protocols *RouterConvergence_Protocols) SetFilter(yf yfilter.YFilter) { protocols.YFilter = yf }

func (protocols *RouterConvergence_Protocols) GetGoName(yname string) string {
    if yname == "protocol" { return "Protocol" }
    return ""
}

func (protocols *RouterConvergence_Protocols) GetSegmentPath() string {
    return "protocols"
}

func (protocols *RouterConvergence_Protocols) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "protocol" {
        for _, c := range protocols.Protocol {
            if protocols.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RouterConvergence_Protocols_Protocol{}
        protocols.Protocol = append(protocols.Protocol, child)
        return &protocols.Protocol[len(protocols.Protocol)-1]
    }
    return nil
}

func (protocols *RouterConvergence_Protocols) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range protocols.Protocol {
        children[protocols.Protocol[i].GetSegmentPath()] = &protocols.Protocol[i]
    }
    return children
}

func (protocols *RouterConvergence_Protocols) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (protocols *RouterConvergence_Protocols) GetBundleName() string { return "cisco_ios_xr" }

func (protocols *RouterConvergence_Protocols) GetYangName() string { return "protocols" }

func (protocols *RouterConvergence_Protocols) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (protocols *RouterConvergence_Protocols) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (protocols *RouterConvergence_Protocols) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (protocols *RouterConvergence_Protocols) SetParent(parent types.Entity) { protocols.parent = parent }

func (protocols *RouterConvergence_Protocols) GetParent() types.Entity { return protocols.parent }

func (protocols *RouterConvergence_Protocols) GetParentYangName() string { return "router-convergence" }

// RouterConvergence_Protocols_Protocol
// Protocol for which to configure RCMD parameters
type RouterConvergence_Protocols_Protocol struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Specify the protocol. The type is ProtocolName.
    ProtocolName interface{}

    // Enable Protocol for which to configure RCMD parameters. Deletion of this
    // object also causes deletion of all associated objects under Protocol. The
    // type is interface{}.
    Enable interface{}

    // Table of Priority.
    Priorities RouterConvergence_Protocols_Protocol_Priorities
}

func (protocol *RouterConvergence_Protocols_Protocol) GetFilter() yfilter.YFilter { return protocol.YFilter }

func (protocol *RouterConvergence_Protocols_Protocol) SetFilter(yf yfilter.YFilter) { protocol.YFilter = yf }

func (protocol *RouterConvergence_Protocols_Protocol) GetGoName(yname string) string {
    if yname == "protocol-name" { return "ProtocolName" }
    if yname == "enable" { return "Enable" }
    if yname == "priorities" { return "Priorities" }
    return ""
}

func (protocol *RouterConvergence_Protocols_Protocol) GetSegmentPath() string {
    return "protocol" + "[protocol-name='" + fmt.Sprintf("%v", protocol.ProtocolName) + "']"
}

func (protocol *RouterConvergence_Protocols_Protocol) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "priorities" {
        return &protocol.Priorities
    }
    return nil
}

func (protocol *RouterConvergence_Protocols_Protocol) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["priorities"] = &protocol.Priorities
    return children
}

func (protocol *RouterConvergence_Protocols_Protocol) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["protocol-name"] = protocol.ProtocolName
    leafs["enable"] = protocol.Enable
    return leafs
}

func (protocol *RouterConvergence_Protocols_Protocol) GetBundleName() string { return "cisco_ios_xr" }

func (protocol *RouterConvergence_Protocols_Protocol) GetYangName() string { return "protocol" }

func (protocol *RouterConvergence_Protocols_Protocol) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (protocol *RouterConvergence_Protocols_Protocol) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (protocol *RouterConvergence_Protocols_Protocol) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (protocol *RouterConvergence_Protocols_Protocol) SetParent(parent types.Entity) { protocol.parent = parent }

func (protocol *RouterConvergence_Protocols_Protocol) GetParent() types.Entity { return protocol.parent }

func (protocol *RouterConvergence_Protocols_Protocol) GetParentYangName() string { return "protocols" }

// RouterConvergence_Protocols_Protocol_Priorities
// Table of Priority
type RouterConvergence_Protocols_Protocol_Priorities struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Priority. The type is slice of
    // RouterConvergence_Protocols_Protocol_Priorities_Priority.
    Priority []RouterConvergence_Protocols_Protocol_Priorities_Priority
}

func (priorities *RouterConvergence_Protocols_Protocol_Priorities) GetFilter() yfilter.YFilter { return priorities.YFilter }

func (priorities *RouterConvergence_Protocols_Protocol_Priorities) SetFilter(yf yfilter.YFilter) { priorities.YFilter = yf }

func (priorities *RouterConvergence_Protocols_Protocol_Priorities) GetGoName(yname string) string {
    if yname == "priority" { return "Priority" }
    return ""
}

func (priorities *RouterConvergence_Protocols_Protocol_Priorities) GetSegmentPath() string {
    return "priorities"
}

func (priorities *RouterConvergence_Protocols_Protocol_Priorities) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "priority" {
        for _, c := range priorities.Priority {
            if priorities.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RouterConvergence_Protocols_Protocol_Priorities_Priority{}
        priorities.Priority = append(priorities.Priority, child)
        return &priorities.Priority[len(priorities.Priority)-1]
    }
    return nil
}

func (priorities *RouterConvergence_Protocols_Protocol_Priorities) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range priorities.Priority {
        children[priorities.Priority[i].GetSegmentPath()] = &priorities.Priority[i]
    }
    return children
}

func (priorities *RouterConvergence_Protocols_Protocol_Priorities) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (priorities *RouterConvergence_Protocols_Protocol_Priorities) GetBundleName() string { return "cisco_ios_xr" }

func (priorities *RouterConvergence_Protocols_Protocol_Priorities) GetYangName() string { return "priorities" }

func (priorities *RouterConvergence_Protocols_Protocol_Priorities) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (priorities *RouterConvergence_Protocols_Protocol_Priorities) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (priorities *RouterConvergence_Protocols_Protocol_Priorities) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (priorities *RouterConvergence_Protocols_Protocol_Priorities) SetParent(parent types.Entity) { priorities.parent = parent }

func (priorities *RouterConvergence_Protocols_Protocol_Priorities) GetParent() types.Entity { return priorities.parent }

func (priorities *RouterConvergence_Protocols_Protocol_Priorities) GetParentYangName() string { return "protocol" }

// RouterConvergence_Protocols_Protocol_Priorities_Priority
// Priority
type RouterConvergence_Protocols_Protocol_Priorities_Priority struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Specify the priority. The type is RcmdPriority.
    RcmdPriority interface{}

    // Threshold value for convergence (in msec). The type is interface{} with
    // range: -2147483648..2147483647.
    Threshold interface{}

    // Specify the maximum number of leaf networks monitored. The type is
    // interface{} with range: 10..100.
    LeafNetworks interface{}

    // Disables the monitoring of route convergence for specified priority. The
    // type is interface{}.
    Disable interface{}

    // Enable Priority. Deletion of this object also causes deletion of all
    // associated objects under Priority. The type is interface{}.
    Enable interface{}

    // Threshold value for Fast ReRoute Coverage (in percentage). The type is
    // interface{} with range: 1..100. Units are percentage.
    FrrThreshold interface{}
}

func (priority *RouterConvergence_Protocols_Protocol_Priorities_Priority) GetFilter() yfilter.YFilter { return priority.YFilter }

func (priority *RouterConvergence_Protocols_Protocol_Priorities_Priority) SetFilter(yf yfilter.YFilter) { priority.YFilter = yf }

func (priority *RouterConvergence_Protocols_Protocol_Priorities_Priority) GetGoName(yname string) string {
    if yname == "rcmd-priority" { return "RcmdPriority" }
    if yname == "threshold" { return "Threshold" }
    if yname == "leaf-networks" { return "LeafNetworks" }
    if yname == "disable" { return "Disable" }
    if yname == "enable" { return "Enable" }
    if yname == "frr-threshold" { return "FrrThreshold" }
    return ""
}

func (priority *RouterConvergence_Protocols_Protocol_Priorities_Priority) GetSegmentPath() string {
    return "priority" + "[rcmd-priority='" + fmt.Sprintf("%v", priority.RcmdPriority) + "']"
}

func (priority *RouterConvergence_Protocols_Protocol_Priorities_Priority) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (priority *RouterConvergence_Protocols_Protocol_Priorities_Priority) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (priority *RouterConvergence_Protocols_Protocol_Priorities_Priority) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rcmd-priority"] = priority.RcmdPriority
    leafs["threshold"] = priority.Threshold
    leafs["leaf-networks"] = priority.LeafNetworks
    leafs["disable"] = priority.Disable
    leafs["enable"] = priority.Enable
    leafs["frr-threshold"] = priority.FrrThreshold
    return leafs
}

func (priority *RouterConvergence_Protocols_Protocol_Priorities_Priority) GetBundleName() string { return "cisco_ios_xr" }

func (priority *RouterConvergence_Protocols_Protocol_Priorities_Priority) GetYangName() string { return "priority" }

func (priority *RouterConvergence_Protocols_Protocol_Priorities_Priority) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (priority *RouterConvergence_Protocols_Protocol_Priorities_Priority) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (priority *RouterConvergence_Protocols_Protocol_Priorities_Priority) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (priority *RouterConvergence_Protocols_Protocol_Priorities_Priority) SetParent(parent types.Entity) { priority.parent = parent }

func (priority *RouterConvergence_Protocols_Protocol_Priorities_Priority) GetParent() types.Entity { return priority.parent }

func (priority *RouterConvergence_Protocols_Protocol_Priorities_Priority) GetParentYangName() string { return "priorities" }

// RouterConvergence_StorageLocation
// Absolute directory path for saving the archive
// files. Example /disk0:/rcmd/ or
// <tftp-location>/rcmd/
// This type is a presence type.
type RouterConvergence_StorageLocation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Absolute directory path for storing diagnostic reports. Example
    // /disk0:/rcmd/ or <tftp-location>/rcmd/. The type is string.
    Diagnostics interface{}

    // Maximum size of diagnostics dir (5% - 80%) for local storage. The type is
    // interface{} with range: 5..80.
    DiagnosticsSize interface{}

    // Maximum size of reports dir (5% - 80%) for local storage. The type is
    // interface{} with range: 5..80.
    ReportsSize interface{}

    // Absolute directory path for storing reports. Example /disk0:/rcmd/ or
    // <tftp-location>/rcmd/. The type is string.
    Reports interface{}
}

func (storageLocation *RouterConvergence_StorageLocation) GetFilter() yfilter.YFilter { return storageLocation.YFilter }

func (storageLocation *RouterConvergence_StorageLocation) SetFilter(yf yfilter.YFilter) { storageLocation.YFilter = yf }

func (storageLocation *RouterConvergence_StorageLocation) GetGoName(yname string) string {
    if yname == "diagnostics" { return "Diagnostics" }
    if yname == "diagnostics-size" { return "DiagnosticsSize" }
    if yname == "reports-size" { return "ReportsSize" }
    if yname == "reports" { return "Reports" }
    return ""
}

func (storageLocation *RouterConvergence_StorageLocation) GetSegmentPath() string {
    return "storage-location"
}

func (storageLocation *RouterConvergence_StorageLocation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (storageLocation *RouterConvergence_StorageLocation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (storageLocation *RouterConvergence_StorageLocation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["diagnostics"] = storageLocation.Diagnostics
    leafs["diagnostics-size"] = storageLocation.DiagnosticsSize
    leafs["reports-size"] = storageLocation.ReportsSize
    leafs["reports"] = storageLocation.Reports
    return leafs
}

func (storageLocation *RouterConvergence_StorageLocation) GetBundleName() string { return "cisco_ios_xr" }

func (storageLocation *RouterConvergence_StorageLocation) GetYangName() string { return "storage-location" }

func (storageLocation *RouterConvergence_StorageLocation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (storageLocation *RouterConvergence_StorageLocation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (storageLocation *RouterConvergence_StorageLocation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (storageLocation *RouterConvergence_StorageLocation) SetParent(parent types.Entity) { storageLocation.parent = parent }

func (storageLocation *RouterConvergence_StorageLocation) GetParent() types.Entity { return storageLocation.parent }

func (storageLocation *RouterConvergence_StorageLocation) GetParentYangName() string { return "router-convergence" }

// RouterConvergence_MplsLdp
// RCMD related configuration for MPLS-LDP
// This type is a presence type.
type RouterConvergence_MplsLdp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Monitoring configuration for Remote LFA.
    RemoteLfa RouterConvergence_MplsLdp_RemoteLfa
}

func (mplsLdp *RouterConvergence_MplsLdp) GetFilter() yfilter.YFilter { return mplsLdp.YFilter }

func (mplsLdp *RouterConvergence_MplsLdp) SetFilter(yf yfilter.YFilter) { mplsLdp.YFilter = yf }

func (mplsLdp *RouterConvergence_MplsLdp) GetGoName(yname string) string {
    if yname == "remote-lfa" { return "RemoteLfa" }
    return ""
}

func (mplsLdp *RouterConvergence_MplsLdp) GetSegmentPath() string {
    return "mpls-ldp"
}

func (mplsLdp *RouterConvergence_MplsLdp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "remote-lfa" {
        return &mplsLdp.RemoteLfa
    }
    return nil
}

func (mplsLdp *RouterConvergence_MplsLdp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["remote-lfa"] = &mplsLdp.RemoteLfa
    return children
}

func (mplsLdp *RouterConvergence_MplsLdp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mplsLdp *RouterConvergence_MplsLdp) GetBundleName() string { return "cisco_ios_xr" }

func (mplsLdp *RouterConvergence_MplsLdp) GetYangName() string { return "mpls-ldp" }

func (mplsLdp *RouterConvergence_MplsLdp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mplsLdp *RouterConvergence_MplsLdp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mplsLdp *RouterConvergence_MplsLdp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mplsLdp *RouterConvergence_MplsLdp) SetParent(parent types.Entity) { mplsLdp.parent = parent }

func (mplsLdp *RouterConvergence_MplsLdp) GetParent() types.Entity { return mplsLdp.parent }

func (mplsLdp *RouterConvergence_MplsLdp) GetParentYangName() string { return "router-convergence" }

// RouterConvergence_MplsLdp_RemoteLfa
// Monitoring configuration for Remote LFA
// This type is a presence type.
type RouterConvergence_MplsLdp_RemoteLfa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Threshold value for label coverage (in percentage). The type is interface{}
    // with range: 1..100. Units are percentage.
    Threshold interface{}
}

func (remoteLfa *RouterConvergence_MplsLdp_RemoteLfa) GetFilter() yfilter.YFilter { return remoteLfa.YFilter }

func (remoteLfa *RouterConvergence_MplsLdp_RemoteLfa) SetFilter(yf yfilter.YFilter) { remoteLfa.YFilter = yf }

func (remoteLfa *RouterConvergence_MplsLdp_RemoteLfa) GetGoName(yname string) string {
    if yname == "threshold" { return "Threshold" }
    return ""
}

func (remoteLfa *RouterConvergence_MplsLdp_RemoteLfa) GetSegmentPath() string {
    return "remote-lfa"
}

func (remoteLfa *RouterConvergence_MplsLdp_RemoteLfa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (remoteLfa *RouterConvergence_MplsLdp_RemoteLfa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (remoteLfa *RouterConvergence_MplsLdp_RemoteLfa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["threshold"] = remoteLfa.Threshold
    return leafs
}

func (remoteLfa *RouterConvergence_MplsLdp_RemoteLfa) GetBundleName() string { return "cisco_ios_xr" }

func (remoteLfa *RouterConvergence_MplsLdp_RemoteLfa) GetYangName() string { return "remote-lfa" }

func (remoteLfa *RouterConvergence_MplsLdp_RemoteLfa) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (remoteLfa *RouterConvergence_MplsLdp_RemoteLfa) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (remoteLfa *RouterConvergence_MplsLdp_RemoteLfa) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (remoteLfa *RouterConvergence_MplsLdp_RemoteLfa) SetParent(parent types.Entity) { remoteLfa.parent = parent }

func (remoteLfa *RouterConvergence_MplsLdp_RemoteLfa) GetParent() types.Entity { return remoteLfa.parent }

func (remoteLfa *RouterConvergence_MplsLdp_RemoteLfa) GetParentYangName() string { return "mpls-ldp" }

// RouterConvergence_CollectDiagnostics
// Table of CollectDiagnostics
type RouterConvergence_CollectDiagnostics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Collect diagnostics on specified node. The type is slice of
    // RouterConvergence_CollectDiagnostics_CollectDiagnostic.
    CollectDiagnostic []RouterConvergence_CollectDiagnostics_CollectDiagnostic
}

func (collectDiagnostics *RouterConvergence_CollectDiagnostics) GetFilter() yfilter.YFilter { return collectDiagnostics.YFilter }

func (collectDiagnostics *RouterConvergence_CollectDiagnostics) SetFilter(yf yfilter.YFilter) { collectDiagnostics.YFilter = yf }

func (collectDiagnostics *RouterConvergence_CollectDiagnostics) GetGoName(yname string) string {
    if yname == "collect-diagnostic" { return "CollectDiagnostic" }
    return ""
}

func (collectDiagnostics *RouterConvergence_CollectDiagnostics) GetSegmentPath() string {
    return "collect-diagnostics"
}

func (collectDiagnostics *RouterConvergence_CollectDiagnostics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "collect-diagnostic" {
        for _, c := range collectDiagnostics.CollectDiagnostic {
            if collectDiagnostics.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RouterConvergence_CollectDiagnostics_CollectDiagnostic{}
        collectDiagnostics.CollectDiagnostic = append(collectDiagnostics.CollectDiagnostic, child)
        return &collectDiagnostics.CollectDiagnostic[len(collectDiagnostics.CollectDiagnostic)-1]
    }
    return nil
}

func (collectDiagnostics *RouterConvergence_CollectDiagnostics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range collectDiagnostics.CollectDiagnostic {
        children[collectDiagnostics.CollectDiagnostic[i].GetSegmentPath()] = &collectDiagnostics.CollectDiagnostic[i]
    }
    return children
}

func (collectDiagnostics *RouterConvergence_CollectDiagnostics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (collectDiagnostics *RouterConvergence_CollectDiagnostics) GetBundleName() string { return "cisco_ios_xr" }

func (collectDiagnostics *RouterConvergence_CollectDiagnostics) GetYangName() string { return "collect-diagnostics" }

func (collectDiagnostics *RouterConvergence_CollectDiagnostics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (collectDiagnostics *RouterConvergence_CollectDiagnostics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (collectDiagnostics *RouterConvergence_CollectDiagnostics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (collectDiagnostics *RouterConvergence_CollectDiagnostics) SetParent(parent types.Entity) { collectDiagnostics.parent = parent }

func (collectDiagnostics *RouterConvergence_CollectDiagnostics) GetParent() types.Entity { return collectDiagnostics.parent }

func (collectDiagnostics *RouterConvergence_CollectDiagnostics) GetParentYangName() string { return "router-convergence" }

// RouterConvergence_CollectDiagnostics_CollectDiagnostic
// Collect diagnostics on specified node
type RouterConvergence_CollectDiagnostics_CollectDiagnostic struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Specified location. The type is string with
    // pattern: ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Enables collection of diagnostics on the specified location. The type is
    // interface{}.
    Enable interface{}
}

func (collectDiagnostic *RouterConvergence_CollectDiagnostics_CollectDiagnostic) GetFilter() yfilter.YFilter { return collectDiagnostic.YFilter }

func (collectDiagnostic *RouterConvergence_CollectDiagnostics_CollectDiagnostic) SetFilter(yf yfilter.YFilter) { collectDiagnostic.YFilter = yf }

func (collectDiagnostic *RouterConvergence_CollectDiagnostics_CollectDiagnostic) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "enable" { return "Enable" }
    return ""
}

func (collectDiagnostic *RouterConvergence_CollectDiagnostics_CollectDiagnostic) GetSegmentPath() string {
    return "collect-diagnostic" + "[node-name='" + fmt.Sprintf("%v", collectDiagnostic.NodeName) + "']"
}

func (collectDiagnostic *RouterConvergence_CollectDiagnostics_CollectDiagnostic) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (collectDiagnostic *RouterConvergence_CollectDiagnostics_CollectDiagnostic) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (collectDiagnostic *RouterConvergence_CollectDiagnostics_CollectDiagnostic) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = collectDiagnostic.NodeName
    leafs["enable"] = collectDiagnostic.Enable
    return leafs
}

func (collectDiagnostic *RouterConvergence_CollectDiagnostics_CollectDiagnostic) GetBundleName() string { return "cisco_ios_xr" }

func (collectDiagnostic *RouterConvergence_CollectDiagnostics_CollectDiagnostic) GetYangName() string { return "collect-diagnostic" }

func (collectDiagnostic *RouterConvergence_CollectDiagnostics_CollectDiagnostic) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (collectDiagnostic *RouterConvergence_CollectDiagnostics_CollectDiagnostic) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (collectDiagnostic *RouterConvergence_CollectDiagnostics_CollectDiagnostic) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (collectDiagnostic *RouterConvergence_CollectDiagnostics_CollectDiagnostic) SetParent(parent types.Entity) { collectDiagnostic.parent = parent }

func (collectDiagnostic *RouterConvergence_CollectDiagnostics_CollectDiagnostic) GetParent() types.Entity { return collectDiagnostic.parent }

func (collectDiagnostic *RouterConvergence_CollectDiagnostics_CollectDiagnostic) GetParentYangName() string { return "collect-diagnostics" }

// RouterConvergence_Nodes
// Table of Node
type RouterConvergence_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure parameters for the specified node (Partially qualified location
    // allowed). The type is slice of RouterConvergence_Nodes_Node.
    Node []RouterConvergence_Nodes_Node
}

func (nodes *RouterConvergence_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *RouterConvergence_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *RouterConvergence_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *RouterConvergence_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *RouterConvergence_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RouterConvergence_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *RouterConvergence_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *RouterConvergence_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *RouterConvergence_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *RouterConvergence_Nodes) GetYangName() string { return "nodes" }

func (nodes *RouterConvergence_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *RouterConvergence_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *RouterConvergence_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *RouterConvergence_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *RouterConvergence_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *RouterConvergence_Nodes) GetParentYangName() string { return "router-convergence" }

// RouterConvergence_Nodes_Node
// Configure parameters for the specified node
// (Partially qualified location allowed)
type RouterConvergence_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Wildcard expression(eg. */*/*, R/*/*, R/S/*,
    // R/S/I). The type is string with pattern:
    // ((([a-zA-Z0-9_]*\d+)|(\*))/){2}(([a-zA-Z0-9_]*\d+)|(\*)).
    NodeName interface{}

    // Disables the monitoring of route convergence on specified location. The
    // type is interface{}.
    Disable interface{}

    // Enable Configure parameters for the specified node (Partially qualified
    // location allowed). Deletion of this object also causes deletion of all
    // associated objects under Node. The type is interface{}.
    Enable interface{}
}

func (node *RouterConvergence_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *RouterConvergence_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *RouterConvergence_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "disable" { return "Disable" }
    if yname == "enable" { return "Enable" }
    return ""
}

func (node *RouterConvergence_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *RouterConvergence_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (node *RouterConvergence_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (node *RouterConvergence_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    leafs["disable"] = node.Disable
    leafs["enable"] = node.Enable
    return leafs
}

func (node *RouterConvergence_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *RouterConvergence_Nodes_Node) GetYangName() string { return "node" }

func (node *RouterConvergence_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *RouterConvergence_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *RouterConvergence_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *RouterConvergence_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *RouterConvergence_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *RouterConvergence_Nodes_Node) GetParentYangName() string { return "nodes" }

